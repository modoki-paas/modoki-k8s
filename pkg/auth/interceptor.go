package auth

type AuthorizerInterceptor struct {
	tokens []string
}

// validateTokenFromHeader validates tokens in env config
func (ai *AuthorizerInterceptor) validateTokenFromHeader(md metadata.MD) bool {
	headers := md.Get("Authorization")

	if len(headers) == 0 {
		return false
	}

	key := strings.TrimPrefix(headers[0], "Bearer ")

	for _, k := range ai.tokens {
		if k == key {
			return true
		}
	}

	return false
}

func (ai *AuthorizerInterceptor) getUserID(md metadata.MD) (id string, ok bool) {
	arr := md.Get(UserIDHeader)

	if len(arr) == 0 {
		return "", false
	}

	return arr[0], true
}

func (ai *AuthorizerInterceptor) getTargetID(md metadata.MD) (id string, ok bool) {
	arr := md.Get(TargetIDHeader)

	if len(arr) == 0 {
		return "", false
	}

	return arr[0], true
}

func (ai *AuthorizerInterceptor) getPermissions(md metadata.MD, targetID string) (permMap map[string]struct{}) {
	arr := md.Get(RolesHeader)

	if len(arr) == 0 {
		return nil
	}

	var rb RoleBindings
	if err := json.Unmarshal([]byte(arr[0]), &rb); err != nil {
		return nil
	}

	perms := []string{}

	systemRoleName, ok := rb["*"]

	if ok {
		systemRole := roles.FindSystemRole(systemRoleName)

		if systemRole != nil {
			for i := range systemRole.Permissions {
				perms = append(perms, systemRole.Permissions[i].Name)
			}
		}
	}

	roleName, ok := rb[targetID]

	if ok {
		role := roles.FindRole(roleName)

		if role != nil {
			for i := range role.Permissions {
				perms = append(perms, role.Permissions[i].Name)
			}
		}
	}

	permMap = map[string]struct{}{}

	for i := range perms {
		permMap[perms[i]] = struct{}{}
	}

	return permMap
}

// UnaryServerInterceptor handles authentication for each call
func UnaryServerInterceptor(tokens []string) grpc.UnaryServerInterceptor {
	ai := &AuthorizerInterceptor{
		tokens: tokens,
	}

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, status.Error(codes.PermissionDenied, "unauthorized: metadata missing")
		}

		if !ai.validateTokenFromHeader(md) {
			return nil, status.Error(codes.PermissionDenied, "unauthorized")
		}

		id, ok := ai.getUserID(md)

		if !ok {
			return nil, status.Error(codes.PermissionDenied, "unknown user")
		}

		ctx = ai.addUserIDContext(ctx, id)

		targetID, ok := ai.getTargetID(md)

		if !ok {
			targetID = id
		}

		perms := ai.getPermissions(md, targetID)

		ctx = ai.addPermissionsContext(ctx, perms)

		return handler(ctx, req)
	}
}