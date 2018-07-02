// Generated from "Event reference" by Mozilla Contributors,
// https://developer.mozilla.org/en-US/docs/Web/Events, licensed under
// CC-BY-SA 2.5.
package markup

var eventMap = map[string]string{
	"afterprint":               "AfterPrint",
	"animationend":             "AnimationEnd",
	"animationiteration":       "AnimationIteration",
	"animationstart":           "AnimationStart",
	"audioprocess":             "AudioProcess",
	"audioend":                 "AudioEnd",
	"audiostart":               "AudioStart",
	"beforeprint":              "BeforePrint",
	"beforeunload":             "BeforeUnload",
	"canplay":                  "CanPlay",
	"canplaythrough":           "CanPlayThrough",
	"chargingchange":           "ChargingChange",
	"chargingtimechange":       "ChargingTimeChange",
	"compassneedscalibration":  "CompassNeedsCalibration",
	"compositionend":           "CompositionEnd",
	"compositionstart":         "CompositionStart",
	"compositionupdate":        "CompositionUpdate",
	"contextmenu":              "ContextMenu",
	"dblclick":                 "DoubleClick",
	"devicelight":              "DeviceLight",
	"devicemotion":             "DeviceMotion",
	"deviceorientation":        "DeviceOrientation",
	"deviceproximity":          "DeviceProximity",
	"dischargingtimechange":    "DischargingTimeChange",
	"dragend":                  "DragEnd",
	"dragenter":                "DragEnter",
	"dragleave":                "DragLeave",
	"dragover":                 "DragOver",
	"dragstart":                "DragStart",
	"durationchange":           "DurationChange",
	"focusin":                  "FocusIn",
	"focusout":                 "FocusOut",
	"fullscreenchange":         "FullScreenChange",
	"fullscreenerror":          "FullScreenError",
	"gamepadconnected":         "GamepadConnected",
	"gamepaddisconnected":      "GamepadDisconnected",
	"gotpointercapture":        "GotPointerCapture",
	"hashchange":               "HashChange",
	"keydown":                  "KeyDown",
	"keypress":                 "KeyPress",
	"keyup":                    "KeyUp",
	"languagechange":           "LanguageChange",
	"levelchange":              "LevelChange",
	"loadeddata":               "LoadedData",
	"loadedmetadata":           "LoadedMetadata",
	"loadend":                  "LoadEnd",
	"loadstart":                "LoadStart",
	"lostpointercapture":       "LostPointerCapture",
	"mousedown":                "MouseDown",
	"mouseenter":               "MouseEnter",
	"mouseleave":               "MouseLeave",
	"mousemove":                "MouseMove",
	"mouseout":                 "MouseOut",
	"mouseover":                "MouseOver",
	"mouseup":                  "MouseUp",
	"noupdate":                 "NoUpdate",
	"nomatch":                  "NoMatch",
	"notificationclick":        "NotificationClick",
	"orientationchange":        "OrientationChange",
	"pagehide":                 "PageHide",
	"pageshow":                 "PageShow",
	"pointercancel":            "PointerCancel",
	"pointerdown":              "PointerDown",
	"pointerenter":             "PointerEnter",
	"pointerleave":             "PointerLeave",
	"pointerlockchange":        "PointerLockChange",
	"pointerlockerror":         "PointerLockError",
	"pointermove":              "PointerMove",
	"pointerout":               "PointerOut",
	"pointerover":              "PointerOver",
	"pointerup":                "PointerUp",
	"popstate":                 "PopState",
	"pushsubscriptionchange":   "PushSubscriptionChange",
	"ratechange":               "RateChange",
	"readystatechange":         "ReadyStateChange",
	"resourcetimingbufferfull": "ResourceTimingBufferFull",
	"selectstart":              "SelectStart",
	"selectionchange":          "SelectionChange",
	"soundend":                 "SoundEnd",
	"soundstart":               "SoundStart",
	"speechend":                "SpeechEnd",
	"speechstart":              "SpeechStart",
	"timeupdate":               "TimeUpdate",
	"touchcancel":              "TouchCancel",
	"touchend":                 "TouchEnd",
	"touchenter":               "TouchEnter",
	"touchleave":               "TouchLeave",
	"touchmove":                "TouchMove",
	"touchstart":               "TouchStart",
	"transitionend":            "TransitionEnd",
	"updateready":              "UpdateReady",
	"upgradeneeded":            "UpgradeNeeded",
	"userproximity":            "UserProximity",
	"versionchange":            "VersionChange",
	"visibilitychange":         "VisibilityChange",
	"voiceschanged":            "VoicesChanged",
	"volumechange":             "VolumeChange",
	"vrdisplayconnected":       "VRDisplayConnected",
	"vrdisplaydisconnected":    "VRDisplayDisconnected",
	"vrdisplaypresentchange":   "VRDisplayPresentChange",
}

// Abort is an event fired when a transaction has been aborted.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/abort_indexedDB
func Abort(listener func(*Event)) *EventListener {
	return &EventListener{Name: "abort", Listener: listener}
}

// AfterPrint is an event fired when the associated document has started
// printing or the print preview has been closed.
//
// https://developer.mozilla.org/docs/Web/Events/afterprint
func AfterPrint(listener func(*Event)) *EventListener {
	return &EventListener{Name: "afterprint", Listener: listener}
}

// AnimationEnd is an event fired when a CSS animation has completed.
//
// https://developer.mozilla.org/docs/Web/Events/animationend
func AnimationEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "animationend", Listener: listener}
}

// AnimationIteration is an event fired when a CSS animation is repeated.
//
// https://developer.mozilla.org/docs/Web/Events/animationiteration
func AnimationIteration(listener func(*Event)) *EventListener {
	return &EventListener{Name: "animationiteration", Listener: listener}
}

// AnimationStart is an event fired when a CSS animation has started.
//
// https://developer.mozilla.org/docs/Web/Events/animationstart
func AnimationStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "animationstart", Listener: listener}
}

// AudioEnd is an event fired when the user agent has finished capturing audio
// for speech recognition.
//
// https://developer.mozilla.org/docs/Web/Events/audioend
func AudioEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "audioend", Listener: listener}
}

// AudioProcess is an event fired when the input buffer of a
// ScriptProcessorNode is ready to be processed.
//
// https://developer.mozilla.org/docs/Web/Events/audioprocess
func AudioProcess(listener func(*Event)) *EventListener {
	return &EventListener{Name: "audioprocess", Listener: listener}
}

// AudioStart is an event fired when the user agent has started to capture
// audio for speech recognition.
//
// https://developer.mozilla.org/docs/Web/Events/audiostart
func AudioStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "audiostart", Listener: listener}
}

// BeforePrint is an event fired when the associated document is about to be
// printed or previewed for printing.
//
// https://developer.mozilla.org/docs/Web/Events/beforeprint
func BeforePrint(listener func(*Event)) *EventListener {
	return &EventListener{Name: "beforeprint", Listener: listener}
}

// BeforeUnload is an event fired when the window, the document and its
// resources are about to be unloaded.
//
// https://developer.mozilla.org/docs/Web/Events/beforeunload
func BeforeUnload(listener func(*Event)) *EventListener {
	return &EventListener{Name: "beforeunload", Listener: listener}
}

// BeginEvent is an event fired when a SMIL animation element begins.
//
// https://developer.mozilla.org/docs/Web/Events/beginEvent
func BeginEvent(listener func(*Event)) *EventListener {
	return &EventListener{Name: "beginEvent", Listener: listener}
}

// Blocked is an event fired when an open connection to a database is blocking
// a versionchange transaction on the same database.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/blocked_indexedDB
func Blocked(listener func(*Event)) *EventListener {
	return &EventListener{Name: "blocked", Listener: listener}
}

// Blur is an event fired when an element has lost focus (does not bubble).
//
// https://developer.mozilla.org/docs/Web/Events/blur
func Blur(listener func(*Event)) *EventListener {
	return &EventListener{Name: "blur", Listener: listener}
}

// Boundary is an event fired when the spoken utterance reaches a word or
// sentence boundary
//
// https://developer.mozilla.org/docs/Web/Events/boundary
func Boundary(listener func(*Event)) *EventListener {
	return &EventListener{Name: "boundary", Listener: listener}
}

// Cached is an event fired when the resources listed in the manifest have been
// downloaded, and the application is now cached.
//
// https://developer.mozilla.org/docs/Web/Events/cached
func Cached(listener func(*Event)) *EventListener {
	return &EventListener{Name: "cached", Listener: listener}
}

// CanPlay is an event fired when the user agent can play the media, but
// estimates that not enough data has been loaded to play the media up to its
// end without having to stop for further buffering of content.
//
// https://developer.mozilla.org/docs/Web/Events/canplay
func CanPlay(listener func(*Event)) *EventListener {
	return &EventListener{Name: "canplay", Listener: listener}
}

// CanPlayThrough is an event fired when the user agent can play the media, and
// estimates that enough data has been loaded to play the media up to its end
// without having to stop for further buffering of content.
//
// https://developer.mozilla.org/docs/Web/Events/canplaythrough
func CanPlayThrough(listener func(*Event)) *EventListener {
	return &EventListener{Name: "canplaythrough", Listener: listener}
}

// Change is an event fired when the change event is fired for <input>,
// <select>, and <textarea> elements when a change to the element's value is
// committed by the user.
//
// https://developer.mozilla.org/docs/Web/Events/change
func Change(listener func(*Event)) *EventListener {
	return &EventListener{Name: "change", Listener: listener}
}

// ChargingChange is an event fired when the battery begins or stops charging.
//
// https://developer.mozilla.org/docs/Web/Events/chargingchange
func ChargingChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "chargingchange", Listener: listener}
}

// ChargingTimeChange is an event fired when the chargingTime attribute has
// been updated.
//
// https://developer.mozilla.org/docs/Web/Events/chargingtimechange
func ChargingTimeChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "chargingtimechange", Listener: listener}
}

// Checking is an event fired when the user agent is checking for an update, or
// attempting to download the cache manifest for the first time.
//
// https://developer.mozilla.org/docs/Web/Events/checking
func Checking(listener func(*Event)) *EventListener {
	return &EventListener{Name: "checking", Listener: listener}
}

// Click is an event fired when a pointing device button has been pressed and
// released on an element.
//
// https://developer.mozilla.org/docs/Web/Events/click
func Click(listener func(*Event)) *EventListener {
	return &EventListener{Name: "click", Listener: listener}
}

// Close is an event fired when a WebSocket connection has been closed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/close_websocket
func Close(listener func(*Event)) *EventListener {
	return &EventListener{Name: "close", Listener: listener}
}

// Complete is an event fired when the rendering of an OfflineAudioContext is
// terminated.
//
// https://developer.mozilla.org/docs/Web/Events/complete
func Complete(listener func(*Event)) *EventListener {
	return &EventListener{Name: "complete", Listener: listener}
}

// CompositionEnd is an event fired when the composition of a passage of text
// has been completed or canceled.
//
// https://developer.mozilla.org/docs/Web/Events/compositionend
func CompositionEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "compositionend", Listener: listener}
}

// CompositionStart is an event fired when the composition of a passage of text
// is prepared (similar to keydown for a keyboard input, but works with other
// inputs such as speech recognition).
//
// https://developer.mozilla.org/docs/Web/Events/compositionstart
func CompositionStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "compositionstart", Listener: listener}
}

// CompositionUpdate is an event fired when a character is added to a passage
// of text being composed.
//
// https://developer.mozilla.org/docs/Web/Events/compositionupdate
func CompositionUpdate(listener func(*Event)) *EventListener {
	return &EventListener{Name: "compositionupdate", Listener: listener}
}

// ContextMenu is an event fired when the right button of the mouse is clicked
// (before the context menu is displayed).
//
// https://developer.mozilla.org/docs/Web/Events/contextmenu
func ContextMenu(listener func(*Event)) *EventListener {
	return &EventListener{Name: "contextmenu", Listener: listener}
}

// Copy is an event fired when the text selection has been added to the
// clipboard.
//
// https://developer.mozilla.org/docs/Web/Events/copy
func Copy(listener func(*Event)) *EventListener {
	return &EventListener{Name: "copy", Listener: listener}
}

// Cut is an event fired when the text selection has been removed from the
// document and added to the clipboard.
//
// https://developer.mozilla.org/docs/Web/Events/cut
func Cut(listener func(*Event)) *EventListener {
	return &EventListener{Name: "cut", Listener: listener}
}

// DOMContentLoaded is an event fired when the document has finished loading
// (but not its dependent resources).
//
// https://developer.mozilla.org/docs/Web/Events/DOMContentLoaded
func DOMContentLoaded(listener func(*Event)) *EventListener {
	return &EventListener{Name: "DOMContentLoaded", Listener: listener}
}

// DeviceLight is an event fired when fresh data is available from a light
// sensor.
//
// https://developer.mozilla.org/docs/Web/Events/devicelight
func DeviceLight(listener func(*Event)) *EventListener {
	return &EventListener{Name: "devicelight", Listener: listener}
}

// DeviceMotion is an event fired when fresh data is available from a motion
// sensor.
//
// https://developer.mozilla.org/docs/Web/Events/devicemotion
func DeviceMotion(listener func(*Event)) *EventListener {
	return &EventListener{Name: "devicemotion", Listener: listener}
}

// DeviceOrientation is an event fired when fresh data is available from an
// orientation sensor.
//
// https://developer.mozilla.org/docs/Web/Events/deviceorientation
func DeviceOrientation(listener func(*Event)) *EventListener {
	return &EventListener{Name: "deviceorientation", Listener: listener}
}

// DeviceProximity is an event fired when fresh data is available from a
// proximity sensor (indicates an approximated distance between the device and
// a nearby object).
//
// https://developer.mozilla.org/docs/Web/Events/deviceproximity
func DeviceProximity(listener func(*Event)) *EventListener {
	return &EventListener{Name: "deviceproximity", Listener: listener}
}

// DischargingTimeChange is an event fired when the dischargingTime attribute
// has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/dischargingtimechange
func DischargingTimeChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dischargingtimechange", Listener: listener}
}

// DoubleClick is an event fired when a pointing device button is clicked twice
// on an element.
//
// https://developer.mozilla.org/docs/Web/Events/dblclick
func DoubleClick(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dblclick", Listener: listener}
}

// Downloading is an event fired when the user agent has found an update and is
// fetching it, or is downloading the resources listed by the cache manifest
// for the first time.
//
// https://developer.mozilla.org/docs/Web/Events/downloading
func Downloading(listener func(*Event)) *EventListener {
	return &EventListener{Name: "downloading", Listener: listener}
}

// Drag is an event fired when an element or text selection is being dragged
// (every 350ms).
//
// https://developer.mozilla.org/docs/Web/Events/drag
func Drag(listener func(*Event)) *EventListener {
	return &EventListener{Name: "drag", Listener: listener}
}

// DragEnd is an event fired when a drag operation is being ended (by releasing
// a mouse button or hitting the escape key).
//
// https://developer.mozilla.org/docs/Web/Events/dragend
func DragEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dragend", Listener: listener}
}

// DragEnter is an event fired when a dragged element or text selection enters
// a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/dragenter
func DragEnter(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dragenter", Listener: listener}
}

// DragLeave is an event fired when a dragged element or text selection leaves
// a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/dragleave
func DragLeave(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dragleave", Listener: listener}
}

// DragOver is an event fired when an element or text selection is being
// dragged over a valid drop target (every 350ms).
//
// https://developer.mozilla.org/docs/Web/Events/dragover
func DragOver(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dragover", Listener: listener}
}

// DragStart is an event fired when the user starts dragging an element or text
// selection.
//
// https://developer.mozilla.org/docs/Web/Events/dragstart
func DragStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "dragstart", Listener: listener}
}

// Drop is an event fired when an element is dropped on a valid drop target.
//
// https://developer.mozilla.org/docs/Web/Events/drop
func Drop(listener func(*Event)) *EventListener {
	return &EventListener{Name: "drop", Listener: listener}
}

// DurationChange is an event fired when the duration attribute has been
// updated.
//
// https://developer.mozilla.org/docs/Web/Events/durationchange
func DurationChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "durationchange", Listener: listener}
}

// Emptied is an event fired when the media has become empty; for example, this
// event is sent if the media has already been loaded (or partially loaded),
// and the load() method is called to reload it.
//
// https://developer.mozilla.org/docs/Web/Events/emptied
func Emptied(listener func(*Event)) *EventListener {
	return &EventListener{Name: "emptied", Listener: listener}
}

// End is an event fired when the utterance has finished being spoken.
//
// https://developer.mozilla.org/docs/Web/Events/end_(SpeechSynthesis)
func End(listener func(*Event)) *EventListener {
	return &EventListener{Name: "end", Listener: listener}
}

// EndEvent is an event fired when a SMIL animation element ends.
//
// https://developer.mozilla.org/docs/Web/Events/endEvent
func EndEvent(listener func(*Event)) *EventListener {
	return &EventListener{Name: "endEvent", Listener: listener}
}

// Ended is an event fired when playback has stopped because the end of the
// media was reached.
//
// https://developer.mozilla.org/docs/Web/Events/ended_(Web_Audio)
func Ended(listener func(*Event)) *EventListener {
	return &EventListener{Name: "ended", Listener: listener}
}

// Error is an event fired when an error occurs that prevents the utterance
// from being successfully spoken.
//
// https://developer.mozilla.org/docs/Web/Events/error_(SpeechSynthesisError)
func Error(listener func(*Event)) *EventListener {
	return &EventListener{Name: "error", Listener: listener}
}

// Focus is an event fired when an element has received focus (does not
// bubble).
//
// https://developer.mozilla.org/docs/Web/Events/focus
func Focus(listener func(*Event)) *EventListener {
	return &EventListener{Name: "focus", Listener: listener}
}

// FocusIn is an event fired when an element is about to receive focus
// (bubbles).
//
// https://developer.mozilla.org/docs/Web/Events/focusin
func FocusIn(listener func(*Event)) *EventListener {
	return &EventListener{Name: "focusin", Listener: listener}
}

// FocusOut is an event fired when an element is about to lose focus (bubbles).
//
// https://developer.mozilla.org/docs/Web/Events/focusout
func FocusOut(listener func(*Event)) *EventListener {
	return &EventListener{Name: "focusout", Listener: listener}
}

// FullScreenChange is an event fired when an element was turned to fullscreen
// mode or back to normal mode.
//
// https://developer.mozilla.org/docs/Web/Events/fullscreenchange
func FullScreenChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "fullscreenchange", Listener: listener}
}

// FullScreenError is an event fired when it was impossible to switch to
// fullscreen mode for technical reasons or because the permission was denied.
//
// https://developer.mozilla.org/docs/Web/Events/fullscreenerror
func FullScreenError(listener func(*Event)) *EventListener {
	return &EventListener{Name: "fullscreenerror", Listener: listener}
}

// GamepadConnected is an event fired when a gamepad has been connected.
//
// https://developer.mozilla.org/docs/Web/Events/gamepadconnected
func GamepadConnected(listener func(*Event)) *EventListener {
	return &EventListener{Name: "gamepadconnected", Listener: listener}
}

// GamepadDisconnected is an event fired when a gamepad has been disconnected.
//
// https://developer.mozilla.org/docs/Web/Events/gamepaddisconnected
func GamepadDisconnected(listener func(*Event)) *EventListener {
	return &EventListener{Name: "gamepaddisconnected", Listener: listener}
}

// GotPointerCapture is an event fired when element receives pointer capture.
//
// https://developer.mozilla.org/docs/Web/Events/gotpointercapture
func GotPointerCapture(listener func(*Event)) *EventListener {
	return &EventListener{Name: "gotpointercapture", Listener: listener}
}

// HashChange is an event fired when the fragment identifier of the URL has
// changed (the part of the URL after the #).
//
// https://developer.mozilla.org/docs/Web/Events/hashchange
func HashChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "hashchange", Listener: listener}
}

// Input is an event fired when the value of an element changes or the content
// of an element with the attribute contenteditable is modified.
//
// https://developer.mozilla.org/docs/Web/Events/input
func Input(listener func(*Event)) *EventListener {
	return &EventListener{Name: "input", Listener: listener}
}

// Invalid is an event fired when a submittable element has been checked and
// doesn't satisfy its constraints.
//
// https://developer.mozilla.org/docs/Web/Events/invalid
func Invalid(listener func(*Event)) *EventListener {
	return &EventListener{Name: "invalid", Listener: listener}
}

// KeyDown is an event fired when a key is pressed down.
//
// https://developer.mozilla.org/docs/Web/Events/keydown
func KeyDown(listener func(*Event)) *EventListener {
	return &EventListener{Name: "keydown", Listener: listener}
}

// KeyPress is an event fired when a key is pressed down and that key normally
// produces a character value (use input instead).
//
// https://developer.mozilla.org/docs/Web/Events/keypress
func KeyPress(listener func(*Event)) *EventListener {
	return &EventListener{Name: "keypress", Listener: listener}
}

// KeyUp is an event fired when a key is released.
//
// https://developer.mozilla.org/docs/Web/Events/keyup
func KeyUp(listener func(*Event)) *EventListener {
	return &EventListener{Name: "keyup", Listener: listener}
}

// LanguageChange is an event fired when the user's preferred languages have
// changed.
//
// https://developer.mozilla.org/docs/Web/Events/languagechange
func LanguageChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "languagechange", Listener: listener}
}

// LevelChange is an event fired when the level attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/levelchange
func LevelChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "levelchange", Listener: listener}
}

// Load is an event fired when progression has been successful.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/load_(ProgressEvent)
func Load(listener func(*Event)) *EventListener {
	return &EventListener{Name: "load", Listener: listener}
}

// LoadEnd is an event fired when progress has stopped (after "error", "abort"
// or "load" have been dispatched).
//
// https://developer.mozilla.org/docs/Web/Events/loadend
func LoadEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "loadend", Listener: listener}
}

// LoadStart is an event fired when progress has begun.
//
// https://developer.mozilla.org/docs/Web/Events/loadstart
func LoadStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "loadstart", Listener: listener}
}

// LoadedData is an event fired when the first frame of the media has finished
// loading.
//
// https://developer.mozilla.org/docs/Web/Events/loadeddata
func LoadedData(listener func(*Event)) *EventListener {
	return &EventListener{Name: "loadeddata", Listener: listener}
}

// LoadedMetadata is an event fired when the metadata has been loaded.
//
// https://developer.mozilla.org/docs/Web/Events/loadedmetadata
func LoadedMetadata(listener func(*Event)) *EventListener {
	return &EventListener{Name: "loadedmetadata", Listener: listener}
}

// LostPointerCapture is an event fired when element lost pointer capture.
//
// https://developer.mozilla.org/docs/Web/Events/lostpointercapture
func LostPointerCapture(listener func(*Event)) *EventListener {
	return &EventListener{Name: "lostpointercapture", Listener: listener}
}

// Mark is an event fired when the spoken utterance reaches a named SSML "mark"
// tag.
//
// https://developer.mozilla.org/docs/Web/Events/mark
func Mark(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mark", Listener: listener}
}

// Message is an event fired when a message is received from a service worker,
// or a message is received in a service worker from another context.
//
// https://developer.mozilla.org/docs/Web/Events/message_(ServiceWorker)
func Message(listener func(*Event)) *EventListener {
	return &EventListener{Name: "message", Listener: listener}
}

// MouseDown is an event fired when a pointing device button (usually a mouse)
// is pressed on an element.
//
// https://developer.mozilla.org/docs/Web/Events/mousedown
func MouseDown(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mousedown", Listener: listener}
}

// MouseEnter is an event fired when a pointing device is moved onto the
// element that has the listener attached.
//
// https://developer.mozilla.org/docs/Web/Events/mouseenter
func MouseEnter(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mouseenter", Listener: listener}
}

// MouseLeave is an event fired when a pointing device is moved off the element
// that has the listener attached.
//
// https://developer.mozilla.org/docs/Web/Events/mouseleave
func MouseLeave(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mouseleave", Listener: listener}
}

// MouseMove is an event fired when a pointing device is moved over an element.
//
// https://developer.mozilla.org/docs/Web/Events/mousemove
func MouseMove(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mousemove", Listener: listener}
}

// MouseOut is an event fired when a pointing device is moved off the element
// that has the listener attached or off one of its children.
//
// https://developer.mozilla.org/docs/Web/Events/mouseout
func MouseOut(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mouseout", Listener: listener}
}

// MouseOver is an event fired when a pointing device is moved onto the element
// that has the listener attached or onto one of its children.
//
// https://developer.mozilla.org/docs/Web/Events/mouseover
func MouseOver(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mouseover", Listener: listener}
}

// MouseUp is an event fired when a pointing device button is released over an
// element.
//
// https://developer.mozilla.org/docs/Web/Events/mouseup
func MouseUp(listener func(*Event)) *EventListener {
	return &EventListener{Name: "mouseup", Listener: listener}
}

// NoMatch is an event fired when the speech recognition service returns a
// final result with no significant recognition.
//
// https://developer.mozilla.org/docs/Web/Events/nomatch
func NoMatch(listener func(*Event)) *EventListener {
	return &EventListener{Name: "nomatch", Listener: listener}
}

// NoUpdate is an event fired when the manifest hadn't changed.
//
// https://developer.mozilla.org/docs/Web/Events/noupdate
func NoUpdate(listener func(*Event)) *EventListener {
	return &EventListener{Name: "noupdate", Listener: listener}
}

// NotificationClick is an event fired when a system notification spawned by
// ServiceWorkerRegistration.showNotification() has been clicked.
//
// https://developer.mozilla.org/docs/Web/Events/notificationclick
func NotificationClick(listener func(*Event)) *EventListener {
	return &EventListener{Name: "notificationclick", Listener: listener}
}

// Obsolete is an event fired when the manifest was found to have become a 404
// or 410 page, so the application cache is being deleted.
//
// https://developer.mozilla.org/docs/Web/Events/obsolete
func Obsolete(listener func(*Event)) *EventListener {
	return &EventListener{Name: "obsolete", Listener: listener}
}

// Offline is an event fired when the browser has lost access to the network.
//
// https://developer.mozilla.org/docs/Web/Events/offline
func Offline(listener func(*Event)) *EventListener {
	return &EventListener{Name: "offline", Listener: listener}
}

// Online is an event fired when the browser has gained access to the network
// (but particular websites might be unreachable).
//
// https://developer.mozilla.org/docs/Web/Events/online
func Online(listener func(*Event)) *EventListener {
	return &EventListener{Name: "online", Listener: listener}
}

// Open is an event fired when an event source connection has been established.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/open_serversentevents
func Open(listener func(*Event)) *EventListener {
	return &EventListener{Name: "open", Listener: listener}
}

// OrientationChange is an event fired when the orientation of the device
// (portrait/landscape) has changed
//
// https://developer.mozilla.org/docs/Web/Events/orientationchange
func OrientationChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "orientationchange", Listener: listener}
}

// PageHide is an event fired when a session history entry is being traversed
// from.
//
// https://developer.mozilla.org/docs/Web/Events/pagehide
func PageHide(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pagehide", Listener: listener}
}

// PageShow is an event fired when a session history entry is being traversed
// to.
//
// https://developer.mozilla.org/docs/Web/Events/pageshow
func PageShow(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pageshow", Listener: listener}
}

// Paste is an event fired when data has been transferred from the system
// clipboard to the document.
//
// https://developer.mozilla.org/docs/Web/Events/paste
func Paste(listener func(*Event)) *EventListener {
	return &EventListener{Name: "paste", Listener: listener}
}

// Pause is an event fired when the utterance is paused part way through.
//
// https://developer.mozilla.org/docs/Web/Events/pause_(SpeechSynthesis)
func Pause(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pause", Listener: listener}
}

// Play is an event fired when playback has begun.
//
// https://developer.mozilla.org/docs/Web/Events/play
func Play(listener func(*Event)) *EventListener {
	return &EventListener{Name: "play", Listener: listener}
}

// Playing is an event fired when playback is ready to start after having been
// paused or delayed due to lack of data.
//
// https://developer.mozilla.org/docs/Web/Events/playing
func Playing(listener func(*Event)) *EventListener {
	return &EventListener{Name: "playing", Listener: listener}
}

// PointerCancel is an event fired when the pointer is unlikely to produce any
// more events.
//
// https://developer.mozilla.org/docs/Web/Events/pointercancel
func PointerCancel(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointercancel", Listener: listener}
}

// PointerDown is an event fired when the pointer enters the active buttons
// state.
//
// https://developer.mozilla.org/docs/Web/Events/pointerdown
func PointerDown(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerdown", Listener: listener}
}

// PointerEnter is an event fired when pointing device is moved inside the
// hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerenter
func PointerEnter(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerenter", Listener: listener}
}

// PointerLeave is an event fired when pointing device is moved out of the
// hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerleave
func PointerLeave(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerleave", Listener: listener}
}

// PointerLockChange is an event fired when the pointer was locked or released.
//
// https://developer.mozilla.org/docs/Web/Events/pointerlockchange
func PointerLockChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerlockchange", Listener: listener}
}

// PointerLockError is an event fired when it was impossible to lock the
// pointer for technical reasons or because the permission was denied.
//
// https://developer.mozilla.org/docs/Web/Events/pointerlockerror
func PointerLockError(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerlockerror", Listener: listener}
}

// PointerMove is an event fired when the pointer changed coordinates.
//
// https://developer.mozilla.org/docs/Web/Events/pointermove
func PointerMove(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointermove", Listener: listener}
}

// PointerOut is an event fired when the pointing device moved out of
// hit-testing boundary or leaves detectable hover range.
//
// https://developer.mozilla.org/docs/Web/Events/pointerout
func PointerOut(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerout", Listener: listener}
}

// PointerOver is an event fired when the pointing device is moved into the
// hit-testing boundary.
//
// https://developer.mozilla.org/docs/Web/Events/pointerover
func PointerOver(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerover", Listener: listener}
}

// PointerUp is an event fired when the pointer leaves the active buttons
// state.
//
// https://developer.mozilla.org/docs/Web/Events/pointerup
func PointerUp(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pointerup", Listener: listener}
}

// PopState is an event fired when a session history entry is being navigated
// to (in certain cases).
//
// https://developer.mozilla.org/docs/Web/Events/popstate
func PopState(listener func(*Event)) *EventListener {
	return &EventListener{Name: "popstate", Listener: listener}
}

// Progress is an event fired when the user agent is downloading resources
// listed by the manifest.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/progress_(appcache_event)
func Progress(listener func(*Event)) *EventListener {
	return &EventListener{Name: "progress", Listener: listener}
}

// Push is an event fired when a Service Worker has received a push message.
//
// https://developer.mozilla.org/docs/Web/Events/push
func Push(listener func(*Event)) *EventListener {
	return &EventListener{Name: "push", Listener: listener}
}

// PushSubscriptionChange is an event fired when a PushSubscription has
// expired.
//
// https://developer.mozilla.org/docs/Web/Events/pushsubscriptionchange
func PushSubscriptionChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "pushsubscriptionchange", Listener: listener}
}

// RateChange is an event fired when the playback rate has changed.
//
// https://developer.mozilla.org/docs/Web/Events/ratechange
func RateChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "ratechange", Listener: listener}
}

// ReadyStateChange is an event fired when the readyState attribute of a
// document has changed.
//
// https://developer.mozilla.org/docs/Web/Events/readystatechange
func ReadyStateChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "readystatechange", Listener: listener}
}

// RepeatEvent is an event fired when a SMIL animation element is repeated.
//
// https://developer.mozilla.org/docs/Web/Events/repeatEvent
func RepeatEvent(listener func(*Event)) *EventListener {
	return &EventListener{Name: "repeatEvent", Listener: listener}
}

// Reset is an event fired when a form is reset.
//
// https://developer.mozilla.org/docs/Web/Events/reset
func Reset(listener func(*Event)) *EventListener {
	return &EventListener{Name: "reset", Listener: listener}
}

// Resize is an event fired when the document view has been resized.
//
// https://developer.mozilla.org/docs/Web/Events/resize
func Resize(listener func(*Event)) *EventListener {
	return &EventListener{Name: "resize", Listener: listener}
}

// ResourceTimingBufferFull is an event fired when the browser's resource
// timing buffer is full.
//
// https://developer.mozilla.org/docs/Web/Events/resourcetimingbufferfull
func ResourceTimingBufferFull(listener func(*Event)) *EventListener {
	return &EventListener{Name: "resourcetimingbufferfull", Listener: listener}
}

// Result is an event fired when the speech recognition service returns a
// result — a word or phrase has been positively recognized and this has been
// communicated back to the app.
//
// https://developer.mozilla.org/docs/Web/Events/result
func Result(listener func(*Event)) *EventListener {
	return &EventListener{Name: "result", Listener: listener}
}

// Resume is an event fired when a paused utterance is resumed.
//
// https://developer.mozilla.org/docs/Web/Events/resume
func Resume(listener func(*Event)) *EventListener {
	return &EventListener{Name: "resume", Listener: listener}
}

// SVGAbort is an event fired when page loading has been stopped before the SVG
// was loaded.
//
// https://developer.mozilla.org/docs/Web/Events/SVGAbort
func SVGAbort(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGAbort", Listener: listener}
}

// SVGError is an event fired when an error has occurred before the SVG was
// loaded.
//
// https://developer.mozilla.org/docs/Web/Events/SVGError
func SVGError(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGError", Listener: listener}
}

// SVGLoad is an event fired when an SVG document has been loaded and parsed.
//
// https://developer.mozilla.org/docs/Web/Events/SVGLoad
func SVGLoad(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGLoad", Listener: listener}
}

// SVGResize is an event fired when an SVG document is being resized.
//
// https://developer.mozilla.org/docs/Web/Events/SVGResize
func SVGResize(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGResize", Listener: listener}
}

// SVGScroll is an event fired when an SVG document is being scrolled.
//
// https://developer.mozilla.org/docs/Web/Events/SVGScroll
func SVGScroll(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGScroll", Listener: listener}
}

// SVGUnload is an event fired when an SVG document has been removed from a
// window or frame.
//
// https://developer.mozilla.org/docs/Web/Events/SVGUnload
func SVGUnload(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGUnload", Listener: listener}
}

// SVGZoom is an event fired when an SVG document is being zoomed.
//
// https://developer.mozilla.org/docs/Web/Events/SVGZoom
func SVGZoom(listener func(*Event)) *EventListener {
	return &EventListener{Name: "SVGZoom", Listener: listener}
}

// Scroll is an event fired when the document view or an element has been
// scrolled.
//
// https://developer.mozilla.org/docs/Web/Events/scroll
func Scroll(listener func(*Event)) *EventListener {
	return &EventListener{Name: "scroll", Listener: listener}
}

// Seeked is an event fired when a seek operation completed.
//
// https://developer.mozilla.org/docs/Web/Events/seeked
func Seeked(listener func(*Event)) *EventListener {
	return &EventListener{Name: "seeked", Listener: listener}
}

// Seeking is an event fired when a seek operation began.
//
// https://developer.mozilla.org/docs/Web/Events/seeking
func Seeking(listener func(*Event)) *EventListener {
	return &EventListener{Name: "seeking", Listener: listener}
}

// Select is an event fired when some text is being selected.
//
// https://developer.mozilla.org/docs/Web/Events/select
func Select(listener func(*Event)) *EventListener {
	return &EventListener{Name: "select", Listener: listener}
}

// SelectStart is an event fired when a selection just started.
//
// https://developer.mozilla.org/docs/Web/Events/selectstart
func SelectStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "selectstart", Listener: listener}
}

// SelectionChange is an event fired when the selection in the document has
// been changed.
//
// https://developer.mozilla.org/docs/Web/Events/selectionchange
func SelectionChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "selectionchange", Listener: listener}
}

// Show is an event fired when a contextmenu event was fired on/bubbled to an
// element that has a contextmenu attribute
//
// https://developer.mozilla.org/docs/Web/Events/show
func Show(listener func(*Event)) *EventListener {
	return &EventListener{Name: "show", Listener: listener}
}

// SoundEnd is an event fired when any sound — recognisable speech or not —
// has stopped being detected.
//
// https://developer.mozilla.org/docs/Web/Events/soundend
func SoundEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "soundend", Listener: listener}
}

// SoundStart is an event fired when any sound — recognisable speech or not
// — has been detected.
//
// https://developer.mozilla.org/docs/Web/Events/soundstart
func SoundStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "soundstart", Listener: listener}
}

// SpeechEnd is an event fired when speech recognised by the speech recognition
// service has stopped being detected.
//
// https://developer.mozilla.org/docs/Web/Events/speechend
func SpeechEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "speechend", Listener: listener}
}

// SpeechStart is an event fired when sound that is recognised by the speech
// recognition service as speech has been detected.
//
// https://developer.mozilla.org/docs/Web/Events/speechstart
func SpeechStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "speechstart", Listener: listener}
}

// Stalled is an event fired when the user agent is trying to fetch media data,
// but data is unexpectedly not forthcoming.
//
// https://developer.mozilla.org/docs/Web/Events/stalled
func Stalled(listener func(*Event)) *EventListener {
	return &EventListener{Name: "stalled", Listener: listener}
}

// Start is an event fired when the utterance has begun to be spoken.
//
// https://developer.mozilla.org/docs/Web/Events/start_(SpeechSynthesis)
func Start(listener func(*Event)) *EventListener {
	return &EventListener{Name: "start", Listener: listener}
}

// Storage is an event fired when a storage area (localStorage or
// sessionStorage) has changed.
//
// https://developer.mozilla.org/docs/Web/Events/storage
func Storage(listener func(*Event)) *EventListener {
	return &EventListener{Name: "storage", Listener: listener}
}

// Submit is an event fired when a form is submitted.
//
// https://developer.mozilla.org/docs/Web/Events/submit
func Submit(listener func(*Event)) *EventListener {
	return &EventListener{Name: "submit", Listener: listener}
}

// Success is an event fired when a request successfully completed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/success_indexedDB
func Success(listener func(*Event)) *EventListener {
	return &EventListener{Name: "success", Listener: listener}
}

// Suspend is an event fired when media data loading has been suspended.
//
// https://developer.mozilla.org/docs/Web/Events/suspend
func Suspend(listener func(*Event)) *EventListener {
	return &EventListener{Name: "suspend", Listener: listener}
}

// TimeUpdate is an event fired when the time indicated by the currentTime
// attribute has been updated.
//
// https://developer.mozilla.org/docs/Web/Events/timeupdate
func TimeUpdate(listener func(*Event)) *EventListener {
	return &EventListener{Name: "timeupdate", Listener: listener}
}

// (no documentation)
//
// https://developer.mozilla.org/docs/Web/Events/timeout
func Timeout(listener func(*Event)) *EventListener {
	return &EventListener{Name: "timeout", Listener: listener}
}

// TouchCancel is an event fired when a touch point has been disrupted in an
// implementation-specific manners (too many touch points for example).
//
// https://developer.mozilla.org/docs/Web/Events/touchcancel
func TouchCancel(listener func(*Event)) *EventListener {
	return &EventListener{Name: "touchcancel", Listener: listener}
}

// TouchEnd is an event fired when a touch point is removed from the touch
// surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchend
func TouchEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "touchend", Listener: listener}
}

// TouchMove is an event fired when a touch point is moved along the touch
// surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchmove
func TouchMove(listener func(*Event)) *EventListener {
	return &EventListener{Name: "touchmove", Listener: listener}
}

// TouchStart is an event fired when a touch point is placed on the touch
// surface.
//
// https://developer.mozilla.org/docs/Web/Events/touchstart
func TouchStart(listener func(*Event)) *EventListener {
	return &EventListener{Name: "touchstart", Listener: listener}
}

// TransitionEnd is an event fired when a CSS transition has completed.
//
// https://developer.mozilla.org/docs/Web/Events/transitionend
func TransitionEnd(listener func(*Event)) *EventListener {
	return &EventListener{Name: "transitionend", Listener: listener}
}

// Unload is an event fired when the document or a dependent resource is being
// unloaded.
//
// https://developer.mozilla.org/docs/Web/Events/unload
func Unload(listener func(*Event)) *EventListener {
	return &EventListener{Name: "unload", Listener: listener}
}

// UpdateReady is an event fired when the resources listed in the manifest have
// been newly redownloaded, and the script can use swapCache() to switch to the
// new cache.
//
// https://developer.mozilla.org/docs/Web/Events/updateready
func UpdateReady(listener func(*Event)) *EventListener {
	return &EventListener{Name: "updateready", Listener: listener}
}

// UpgradeNeeded is an event fired when an attempt was made to open a database
// with a version number higher than its current version. A versionchange
// transaction has been created.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/upgradeneeded_indexedDB
func UpgradeNeeded(listener func(*Event)) *EventListener {
	return &EventListener{Name: "upgradeneeded", Listener: listener}
}

// UserProximity is an event fired when fresh data is available from a
// proximity sensor (indicates whether the nearby object is near the device or
// not).
//
// https://developer.mozilla.org/docs/Web/Events/userproximity
func UserProximity(listener func(*Event)) *EventListener {
	return &EventListener{Name: "userproximity", Listener: listener}
}

// VersionChange is an event fired when a versionchange transaction completed.
//
// https://developer.mozilla.org/docs/Web/Reference/Events/versionchange_indexedDB
func VersionChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "versionchange", Listener: listener}
}

// VisibilityChange is an event fired when the content of a tab has become
// visible or has been hidden.
//
// https://developer.mozilla.org/docs/Web/Events/visibilitychange
func VisibilityChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "visibilitychange", Listener: listener}
}

// VoicesChanged is an event fired when the list of SpeechSynthesisVoice
// objects that would be returned by the SpeechSynthesis.getVoices() method has
// changed (when the voiceschanged event fires.)
//
// https://developer.mozilla.org/docs/Web/Events/voiceschanged
func VoicesChanged(listener func(*Event)) *EventListener {
	return &EventListener{Name: "voiceschanged", Listener: listener}
}

// VolumeChange is an event fired when the volume has changed.
//
// https://developer.mozilla.org/docs/Web/Events/volumechange
func VolumeChange(listener func(*Event)) *EventListener {
	return &EventListener{Name: "volumechange", Listener: listener}
}

// Waiting is an event fired when playback has stopped because of a temporary
// lack of data.
//
// https://developer.mozilla.org/docs/Web/Events/waiting
func Waiting(listener func(*Event)) *EventListener {
	return &EventListener{Name: "waiting", Listener: listener}
}

// Wheel is an event fired when a wheel button of a pointing device is rotated
// in any direction.
//
// https://developer.mozilla.org/docs/Web/Events/wheel
func Wheel(listener func(*Event)) *EventListener {
	return &EventListener{Name: "wheel", Listener: listener}
}
