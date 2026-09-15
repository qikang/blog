// Package assets 把 templates 和 static 通过 go:embed 编译进二进制,
//blog 最终以单一可执行文件形式分发,运行时不再依赖外部 templates/static 目录。
package assets

import "embed"

//go:embed templates/*.html
var TemplatesFS embed.FS

//go:embed static
var StaticFS embed.FS
