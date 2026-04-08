# Shared Params Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared#OrderParam">OrderParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared#Order">Order</a>

# Pets

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#CategoryParam">CategoryParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetParam">PetParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#TagParam">TagParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Category">Category</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Pet">Pet</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Tag">Tag</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetUploadImageResponse">PetUploadImageResponse</a>

Methods:

- <code title="post /pet">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetNewParams">PetNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pet">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetUpdateParams">PetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /pet/findByStatus">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.FindByStatus">FindByStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetFindByStatusParams">PetFindByStatusParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pet/findByTags">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.FindByTags">FindByTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetFindByTagsParams">PetFindByTagsParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#Pet">Pet</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /pet/{petId}">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.UpdateByID">UpdateByID</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetUpdateByIDParams">PetUpdateByIDParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /pet/{petId}/uploadImage">client.Pets.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetService.UploadImage">UploadImage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, petID <a href="https://pkg.go.dev/builtin#int64">int64</a>, image <a href="https://pkg.go.dev/builtin#io.Reader">io.Reader</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetUploadImageParams">PetUploadImageParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#PetUploadImageResponse">PetUploadImageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Store

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreListInventoryResponse">StoreListInventoryResponse</a>

Methods:

- <code title="get /store/inventory">client.Store.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreService.ListInventory">ListInventory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreListInventoryResponse">StoreListInventoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orders

Methods:

- <code title="post /store/order">client.Store.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreOrderService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreOrderNewParams">StoreOrderNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /store/order/{orderId}">client.Store.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreOrderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go/shared#Order">Order</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /store/order/{orderId}">client.Store.Orders.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#StoreOrderService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orderID <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserParam">UserParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#User">User</a>

Methods:

- <code title="post /user">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserNewParams">UserNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, existingUsername <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserUpdateParams">UserUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /user/{username}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, username <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /user/createWithList">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.NewWithList">NewWithList</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserNewWithListParams">UserNewWithListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/login">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserLoginParams">UserLoginParams</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/logout">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/luma-agents-go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
