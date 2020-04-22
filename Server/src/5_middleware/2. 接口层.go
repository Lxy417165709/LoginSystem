package middleware

import "4_controls"

var Test = NewHandler(controls.Test, []middleWareType{TimeCounter, TimesCounter}).Format()
var Login = NewHandler(controls.Login, []middleWareType{TimeCounter, TimesCounter}).Format()
var Register = NewHandler(controls.Register, []middleWareType{TimeCounter, TimesCounter}).Format()
var UpdateUpi = NewHandler(controls.UpdateUpi, []middleWareType{TimeCounter, TimesCounter}).Format()
var GetUpi = NewHandler(controls.GetUpi, []middleWareType{TimeCounter, TimesCounter}).Format()
var SendRegisterVrc = NewHandler(controls.SendRegisterVrc, []middleWareType{TimeCounter, TimesCounter}).Format()
var GetPhoto = NewHandler(controls.GetPhoto, []middleWareType{TimeCounter, TimesCounter}).Format()
var UpdatePhoto = NewHandler(controls.UpdatePhoto, []middleWareType{TimeCounter, TimesCounter}).Format()
var GetUai = NewHandler(controls.GetUai, []middleWareType{TimeCounter, TimesCounter}).Format()
