package api

import "github.com/dop251/goja"

// Page is the interface of a single browser tab.
type Page interface {
	AddInitScript(script goja.Value, arg goja.Value) *goja.Promise
	AddScriptTag(opts goja.Value) *goja.Promise
	AddStyleTag(opts goja.Value) *goja.Promise
	BringToFront()
	Check(selector string, opts goja.Value)
	Click(selector string, opts goja.Value) error
	Close(opts goja.Value)
	Content() string
	Context() BrowserContext
	Dblclick(selector string, opts goja.Value)
	DispatchEvent(selector string, typ string, eventInit goja.Value, opts goja.Value)
	DragAndDrop(source string, target string, opts goja.Value) *goja.Promise
	EmulateMedia(opts goja.Value)
	EmulateVisionDeficiency(typ string)
	Evaluate(pageFunc goja.Value, arg ...goja.Value) any
	EvaluateHandle(pageFunc goja.Value, arg ...goja.Value) JSHandle
	ExposeBinding(name string, callback goja.Callable, opts goja.Value) *goja.Promise
	ExposeFunction(name string, callback goja.Callable) *goja.Promise
	Fill(selector string, value string, opts goja.Value)
	Focus(selector string, opts goja.Value)
	Frame(frameSelector goja.Value) *goja.Promise
	Frames() []Frame
	GetAttribute(selector string, name string, opts goja.Value) goja.Value
	GoBack(opts goja.Value) *goja.Promise
	GoForward(opts goja.Value) *goja.Promise
	Goto(url string, opts goja.Value) (Response, error)
	Hover(selector string, opts goja.Value)
	InnerHTML(selector string, opts goja.Value) string
	InnerText(selector string, opts goja.Value) string
	InputValue(selector string, opts goja.Value) string
	IsChecked(selector string, opts goja.Value) bool
	IsClosed() bool
	IsDisabled(selector string, opts goja.Value) bool
	IsEditable(selector string, opts goja.Value) bool
	IsEnabled(selector string, opts goja.Value) bool
	IsHidden(selector string, opts goja.Value) bool
	IsVisible(selector string, opts goja.Value) bool
	// Locator creates and returns a new locator for this page (main frame).
	Locator(selector string, opts goja.Value) Locator
	MainFrame() Frame
	Opener() Page
	Pause() *goja.Promise
	Pdf(opts goja.Value) *goja.Promise
	Press(selector string, key string, opts goja.Value)
	Query(selector string) ElementHandle
	QueryAll(selector string) []ElementHandle
	Reload(opts goja.Value) Response
	Route(url goja.Value, handler goja.Callable) *goja.Promise
	Screenshot(opts goja.Value) goja.ArrayBuffer
	SelectOption(selector string, values goja.Value, opts goja.Value) []string
	SetContent(html string, opts goja.Value)
	SetDefaultNavigationTimeout(timeout int64)
	SetDefaultTimeout(timeout int64)
	SetExtraHTTPHeaders(headers map[string]string)
	SetInputFiles(selector string, files goja.Value, opts goja.Value) *goja.Promise
	SetViewportSize(viewportSize goja.Value)
	Tap(selector string, opts goja.Value)
	TextContent(selector string, opts goja.Value) string
	Title() string
	Type(selector string, text string, opts goja.Value)
	Uncheck(selector string, opts goja.Value)
	Unroute(url goja.Value, handler goja.Callable) *goja.Promise
	URL() string
	Video() *goja.Promise
	ViewportSize() map[string]float64
	WaitForEvent(event string, optsOrPredicate goja.Value) *goja.Promise
	WaitForFunction(fn, opts goja.Value, args ...goja.Value) (any, error)
	WaitForLoadState(state string, opts goja.Value)
	WaitForNavigation(opts goja.Value) (Response, error)
	WaitForRequest(urlOrPredicate, opts goja.Value) *goja.Promise
	WaitForResponse(urlOrPredicate, opts goja.Value) *goja.Promise
	WaitForSelector(selector string, opts goja.Value) ElementHandle
	WaitForTimeout(timeout int64)
	Workers() []Worker
}
