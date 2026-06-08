# Generations

Params Types:

- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#AdvancedControlsParam">AdvancedControlsParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#DepthControlParam">DepthControlParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#FaceControlParam">FaceControlParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#ImageRefParam">ImageRefParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#NormalsControlParam">NormalsControlParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#PoseControlParam">PoseControlParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#PoseControlStrength">PoseControlStrength</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#SourcePositionParam">SourcePositionParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#TrajectoryControlParam">TrajectoryControlParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#VideoDuration">VideoDuration</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#VideoEditOptionsParam">VideoEditOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#VideoEditStrength">VideoEditStrength</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#VideoOptionsParam">VideoOptionsParam</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#VideoResolution">VideoResolution</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#Generation">Generation</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#GenerationFailureCode">GenerationFailureCode</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#GenerationOutput">GenerationOutput</a>
- <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#Model">Model</a>

Methods:

- <code title="post /generations">client.Generations.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#GenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#GenerationNewParams">GenerationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#Generation">Generation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /generations/{generation_id}">client.Generations.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#GenerationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, generationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go">lumaagents</a>.<a href="https://pkg.go.dev/github.com/lumalabs/luma-agents-go#Generation">Generation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
