// Code generated by mockery. DO NOT EDIT.

package console

import (
	console "github.com/goravel/framework/contracts/console"
	mock "github.com/stretchr/testify/mock"
)

// Context is an autogenerated mock type for the Context type
type Context struct {
	mock.Mock
}

type Context_Expecter struct {
	mock *mock.Mock
}

func (_m *Context) EXPECT() *Context_Expecter {
	return &Context_Expecter{mock: &_m.Mock}
}

// Argument provides a mock function with given fields: index
func (_m *Context) Argument(index int) string {
	ret := _m.Called(index)

	if len(ret) == 0 {
		panic("no return value specified for Argument")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(index)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Context_Argument_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Argument'
type Context_Argument_Call struct {
	*mock.Call
}

// Argument is a helper method to define mock.On call
//   - index int
func (_e *Context_Expecter) Argument(index interface{}) *Context_Argument_Call {
	return &Context_Argument_Call{Call: _e.mock.On("Argument", index)}
}

func (_c *Context_Argument_Call) Run(run func(index int)) *Context_Argument_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Context_Argument_Call) Return(_a0 string) *Context_Argument_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_Argument_Call) RunAndReturn(run func(int) string) *Context_Argument_Call {
	_c.Call.Return(run)
	return _c
}

// Arguments provides a mock function with given fields:
func (_m *Context) Arguments() []string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Arguments")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Context_Arguments_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Arguments'
type Context_Arguments_Call struct {
	*mock.Call
}

// Arguments is a helper method to define mock.On call
func (_e *Context_Expecter) Arguments() *Context_Arguments_Call {
	return &Context_Arguments_Call{Call: _e.mock.On("Arguments")}
}

func (_c *Context_Arguments_Call) Run(run func()) *Context_Arguments_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Context_Arguments_Call) Return(_a0 []string) *Context_Arguments_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_Arguments_Call) RunAndReturn(run func() []string) *Context_Arguments_Call {
	_c.Call.Return(run)
	return _c
}

// Ask provides a mock function with given fields: question, option
func (_m *Context) Ask(question string, option ...console.AskOption) (string, error) {
	_va := make([]interface{}, len(option))
	for _i := range option {
		_va[_i] = option[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, question)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Ask")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...console.AskOption) (string, error)); ok {
		return rf(question, option...)
	}
	if rf, ok := ret.Get(0).(func(string, ...console.AskOption) string); ok {
		r0 = rf(question, option...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, ...console.AskOption) error); ok {
		r1 = rf(question, option...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context_Ask_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ask'
type Context_Ask_Call struct {
	*mock.Call
}

// Ask is a helper method to define mock.On call
//   - question string
//   - option ...console.AskOption
func (_e *Context_Expecter) Ask(question interface{}, option ...interface{}) *Context_Ask_Call {
	return &Context_Ask_Call{Call: _e.mock.On("Ask",
		append([]interface{}{question}, option...)...)}
}

func (_c *Context_Ask_Call) Run(run func(question string, option ...console.AskOption)) *Context_Ask_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]console.AskOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(console.AskOption)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Context_Ask_Call) Return(_a0 string, _a1 error) *Context_Ask_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Context_Ask_Call) RunAndReturn(run func(string, ...console.AskOption) (string, error)) *Context_Ask_Call {
	_c.Call.Return(run)
	return _c
}

// Choice provides a mock function with given fields: question, options, option
func (_m *Context) Choice(question string, options []console.Choice, option ...console.ChoiceOption) (string, error) {
	_va := make([]interface{}, len(option))
	for _i := range option {
		_va[_i] = option[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, question, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Choice")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []console.Choice, ...console.ChoiceOption) (string, error)); ok {
		return rf(question, options, option...)
	}
	if rf, ok := ret.Get(0).(func(string, []console.Choice, ...console.ChoiceOption) string); ok {
		r0 = rf(question, options, option...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, []console.Choice, ...console.ChoiceOption) error); ok {
		r1 = rf(question, options, option...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context_Choice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Choice'
type Context_Choice_Call struct {
	*mock.Call
}

// Choice is a helper method to define mock.On call
//   - question string
//   - options []console.Choice
//   - option ...console.ChoiceOption
func (_e *Context_Expecter) Choice(question interface{}, options interface{}, option ...interface{}) *Context_Choice_Call {
	return &Context_Choice_Call{Call: _e.mock.On("Choice",
		append([]interface{}{question, options}, option...)...)}
}

func (_c *Context_Choice_Call) Run(run func(question string, options []console.Choice, option ...console.ChoiceOption)) *Context_Choice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]console.ChoiceOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(console.ChoiceOption)
			}
		}
		run(args[0].(string), args[1].([]console.Choice), variadicArgs...)
	})
	return _c
}

func (_c *Context_Choice_Call) Return(_a0 string, _a1 error) *Context_Choice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Context_Choice_Call) RunAndReturn(run func(string, []console.Choice, ...console.ChoiceOption) (string, error)) *Context_Choice_Call {
	_c.Call.Return(run)
	return _c
}

// Comment provides a mock function with given fields: message
func (_m *Context) Comment(message string) {
	_m.Called(message)
}

// Context_Comment_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Comment'
type Context_Comment_Call struct {
	*mock.Call
}

// Comment is a helper method to define mock.On call
//   - message string
func (_e *Context_Expecter) Comment(message interface{}) *Context_Comment_Call {
	return &Context_Comment_Call{Call: _e.mock.On("Comment", message)}
}

func (_c *Context_Comment_Call) Run(run func(message string)) *Context_Comment_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_Comment_Call) Return() *Context_Comment_Call {
	_c.Call.Return()
	return _c
}

func (_c *Context_Comment_Call) RunAndReturn(run func(string)) *Context_Comment_Call {
	_c.Call.Return(run)
	return _c
}

// Confirm provides a mock function with given fields: question, option
func (_m *Context) Confirm(question string, option ...console.ConfirmOption) (bool, error) {
	_va := make([]interface{}, len(option))
	for _i := range option {
		_va[_i] = option[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, question)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Confirm")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...console.ConfirmOption) (bool, error)); ok {
		return rf(question, option...)
	}
	if rf, ok := ret.Get(0).(func(string, ...console.ConfirmOption) bool); ok {
		r0 = rf(question, option...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, ...console.ConfirmOption) error); ok {
		r1 = rf(question, option...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context_Confirm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Confirm'
type Context_Confirm_Call struct {
	*mock.Call
}

// Confirm is a helper method to define mock.On call
//   - question string
//   - option ...console.ConfirmOption
func (_e *Context_Expecter) Confirm(question interface{}, option ...interface{}) *Context_Confirm_Call {
	return &Context_Confirm_Call{Call: _e.mock.On("Confirm",
		append([]interface{}{question}, option...)...)}
}

func (_c *Context_Confirm_Call) Run(run func(question string, option ...console.ConfirmOption)) *Context_Confirm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]console.ConfirmOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(console.ConfirmOption)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Context_Confirm_Call) Return(_a0 bool, _a1 error) *Context_Confirm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Context_Confirm_Call) RunAndReturn(run func(string, ...console.ConfirmOption) (bool, error)) *Context_Confirm_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProgressBar provides a mock function with given fields: total
func (_m *Context) CreateProgressBar(total int) console.Progress {
	ret := _m.Called(total)

	if len(ret) == 0 {
		panic("no return value specified for CreateProgressBar")
	}

	var r0 console.Progress
	if rf, ok := ret.Get(0).(func(int) console.Progress); ok {
		r0 = rf(total)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(console.Progress)
		}
	}

	return r0
}

// Context_CreateProgressBar_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProgressBar'
type Context_CreateProgressBar_Call struct {
	*mock.Call
}

// CreateProgressBar is a helper method to define mock.On call
//   - total int
func (_e *Context_Expecter) CreateProgressBar(total interface{}) *Context_CreateProgressBar_Call {
	return &Context_CreateProgressBar_Call{Call: _e.mock.On("CreateProgressBar", total)}
}

func (_c *Context_CreateProgressBar_Call) Run(run func(total int)) *Context_CreateProgressBar_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Context_CreateProgressBar_Call) Return(_a0 console.Progress) *Context_CreateProgressBar_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_CreateProgressBar_Call) RunAndReturn(run func(int) console.Progress) *Context_CreateProgressBar_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields: message
func (_m *Context) Error(message string) {
	_m.Called(message)
}

// Context_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Context_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
//   - message string
func (_e *Context_Expecter) Error(message interface{}) *Context_Error_Call {
	return &Context_Error_Call{Call: _e.mock.On("Error", message)}
}

func (_c *Context_Error_Call) Run(run func(message string)) *Context_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_Error_Call) Return() *Context_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *Context_Error_Call) RunAndReturn(run func(string)) *Context_Error_Call {
	_c.Call.Return(run)
	return _c
}

// Info provides a mock function with given fields: message
func (_m *Context) Info(message string) {
	_m.Called(message)
}

// Context_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type Context_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
//   - message string
func (_e *Context_Expecter) Info(message interface{}) *Context_Info_Call {
	return &Context_Info_Call{Call: _e.mock.On("Info", message)}
}

func (_c *Context_Info_Call) Run(run func(message string)) *Context_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_Info_Call) Return() *Context_Info_Call {
	_c.Call.Return()
	return _c
}

func (_c *Context_Info_Call) RunAndReturn(run func(string)) *Context_Info_Call {
	_c.Call.Return(run)
	return _c
}

// Line provides a mock function with given fields: message
func (_m *Context) Line(message string) {
	_m.Called(message)
}

// Context_Line_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Line'
type Context_Line_Call struct {
	*mock.Call
}

// Line is a helper method to define mock.On call
//   - message string
func (_e *Context_Expecter) Line(message interface{}) *Context_Line_Call {
	return &Context_Line_Call{Call: _e.mock.On("Line", message)}
}

func (_c *Context_Line_Call) Run(run func(message string)) *Context_Line_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_Line_Call) Return() *Context_Line_Call {
	_c.Call.Return()
	return _c
}

func (_c *Context_Line_Call) RunAndReturn(run func(string)) *Context_Line_Call {
	_c.Call.Return(run)
	return _c
}

// MultiSelect provides a mock function with given fields: question, options, option
func (_m *Context) MultiSelect(question string, options []console.Choice, option ...console.MultiSelectOption) ([]string, error) {
	_va := make([]interface{}, len(option))
	for _i := range option {
		_va[_i] = option[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, question, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for MultiSelect")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, []console.Choice, ...console.MultiSelectOption) ([]string, error)); ok {
		return rf(question, options, option...)
	}
	if rf, ok := ret.Get(0).(func(string, []console.Choice, ...console.MultiSelectOption) []string); ok {
		r0 = rf(question, options, option...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(string, []console.Choice, ...console.MultiSelectOption) error); ok {
		r1 = rf(question, options, option...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context_MultiSelect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MultiSelect'
type Context_MultiSelect_Call struct {
	*mock.Call
}

// MultiSelect is a helper method to define mock.On call
//   - question string
//   - options []console.Choice
//   - option ...console.MultiSelectOption
func (_e *Context_Expecter) MultiSelect(question interface{}, options interface{}, option ...interface{}) *Context_MultiSelect_Call {
	return &Context_MultiSelect_Call{Call: _e.mock.On("MultiSelect",
		append([]interface{}{question, options}, option...)...)}
}

func (_c *Context_MultiSelect_Call) Run(run func(question string, options []console.Choice, option ...console.MultiSelectOption)) *Context_MultiSelect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]console.MultiSelectOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(console.MultiSelectOption)
			}
		}
		run(args[0].(string), args[1].([]console.Choice), variadicArgs...)
	})
	return _c
}

func (_c *Context_MultiSelect_Call) Return(_a0 []string, _a1 error) *Context_MultiSelect_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Context_MultiSelect_Call) RunAndReturn(run func(string, []console.Choice, ...console.MultiSelectOption) ([]string, error)) *Context_MultiSelect_Call {
	_c.Call.Return(run)
	return _c
}

// NewLine provides a mock function with given fields: times
func (_m *Context) NewLine(times ...int) {
	_va := make([]interface{}, len(times))
	for _i := range times {
		_va[_i] = times[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Context_NewLine_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewLine'
type Context_NewLine_Call struct {
	*mock.Call
}

// NewLine is a helper method to define mock.On call
//   - times ...int
func (_e *Context_Expecter) NewLine(times ...interface{}) *Context_NewLine_Call {
	return &Context_NewLine_Call{Call: _e.mock.On("NewLine",
		append([]interface{}{}, times...)...)}
}

func (_c *Context_NewLine_Call) Run(run func(times ...int)) *Context_NewLine_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(int)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Context_NewLine_Call) Return() *Context_NewLine_Call {
	_c.Call.Return()
	return _c
}

func (_c *Context_NewLine_Call) RunAndReturn(run func(...int)) *Context_NewLine_Call {
	_c.Call.Return(run)
	return _c
}

// Option provides a mock function with given fields: key
func (_m *Context) Option(key string) string {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for Option")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Context_Option_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Option'
type Context_Option_Call struct {
	*mock.Call
}

// Option is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) Option(key interface{}) *Context_Option_Call {
	return &Context_Option_Call{Call: _e.mock.On("Option", key)}
}

func (_c *Context_Option_Call) Run(run func(key string)) *Context_Option_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_Option_Call) Return(_a0 string) *Context_Option_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_Option_Call) RunAndReturn(run func(string) string) *Context_Option_Call {
	_c.Call.Return(run)
	return _c
}

// OptionBool provides a mock function with given fields: key
func (_m *Context) OptionBool(key string) bool {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionBool")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Context_OptionBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionBool'
type Context_OptionBool_Call struct {
	*mock.Call
}

// OptionBool is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionBool(key interface{}) *Context_OptionBool_Call {
	return &Context_OptionBool_Call{Call: _e.mock.On("OptionBool", key)}
}

func (_c *Context_OptionBool_Call) Run(run func(key string)) *Context_OptionBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionBool_Call) Return(_a0 bool) *Context_OptionBool_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionBool_Call) RunAndReturn(run func(string) bool) *Context_OptionBool_Call {
	_c.Call.Return(run)
	return _c
}

// OptionFloat64 provides a mock function with given fields: key
func (_m *Context) OptionFloat64(key string) float64 {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionFloat64")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Context_OptionFloat64_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionFloat64'
type Context_OptionFloat64_Call struct {
	*mock.Call
}

// OptionFloat64 is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionFloat64(key interface{}) *Context_OptionFloat64_Call {
	return &Context_OptionFloat64_Call{Call: _e.mock.On("OptionFloat64", key)}
}

func (_c *Context_OptionFloat64_Call) Run(run func(key string)) *Context_OptionFloat64_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionFloat64_Call) Return(_a0 float64) *Context_OptionFloat64_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionFloat64_Call) RunAndReturn(run func(string) float64) *Context_OptionFloat64_Call {
	_c.Call.Return(run)
	return _c
}

// OptionFloat64Slice provides a mock function with given fields: key
func (_m *Context) OptionFloat64Slice(key string) []float64 {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionFloat64Slice")
	}

	var r0 []float64
	if rf, ok := ret.Get(0).(func(string) []float64); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]float64)
		}
	}

	return r0
}

// Context_OptionFloat64Slice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionFloat64Slice'
type Context_OptionFloat64Slice_Call struct {
	*mock.Call
}

// OptionFloat64Slice is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionFloat64Slice(key interface{}) *Context_OptionFloat64Slice_Call {
	return &Context_OptionFloat64Slice_Call{Call: _e.mock.On("OptionFloat64Slice", key)}
}

func (_c *Context_OptionFloat64Slice_Call) Run(run func(key string)) *Context_OptionFloat64Slice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionFloat64Slice_Call) Return(_a0 []float64) *Context_OptionFloat64Slice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionFloat64Slice_Call) RunAndReturn(run func(string) []float64) *Context_OptionFloat64Slice_Call {
	_c.Call.Return(run)
	return _c
}

// OptionInt provides a mock function with given fields: key
func (_m *Context) OptionInt(key string) int {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionInt")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Context_OptionInt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionInt'
type Context_OptionInt_Call struct {
	*mock.Call
}

// OptionInt is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionInt(key interface{}) *Context_OptionInt_Call {
	return &Context_OptionInt_Call{Call: _e.mock.On("OptionInt", key)}
}

func (_c *Context_OptionInt_Call) Run(run func(key string)) *Context_OptionInt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionInt_Call) Return(_a0 int) *Context_OptionInt_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionInt_Call) RunAndReturn(run func(string) int) *Context_OptionInt_Call {
	_c.Call.Return(run)
	return _c
}

// OptionInt64 provides a mock function with given fields: key
func (_m *Context) OptionInt64(key string) int64 {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionInt64")
	}

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Context_OptionInt64_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionInt64'
type Context_OptionInt64_Call struct {
	*mock.Call
}

// OptionInt64 is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionInt64(key interface{}) *Context_OptionInt64_Call {
	return &Context_OptionInt64_Call{Call: _e.mock.On("OptionInt64", key)}
}

func (_c *Context_OptionInt64_Call) Run(run func(key string)) *Context_OptionInt64_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionInt64_Call) Return(_a0 int64) *Context_OptionInt64_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionInt64_Call) RunAndReturn(run func(string) int64) *Context_OptionInt64_Call {
	_c.Call.Return(run)
	return _c
}

// OptionInt64Slice provides a mock function with given fields: key
func (_m *Context) OptionInt64Slice(key string) []int64 {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionInt64Slice")
	}

	var r0 []int64
	if rf, ok := ret.Get(0).(func(string) []int64); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int64)
		}
	}

	return r0
}

// Context_OptionInt64Slice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionInt64Slice'
type Context_OptionInt64Slice_Call struct {
	*mock.Call
}

// OptionInt64Slice is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionInt64Slice(key interface{}) *Context_OptionInt64Slice_Call {
	return &Context_OptionInt64Slice_Call{Call: _e.mock.On("OptionInt64Slice", key)}
}

func (_c *Context_OptionInt64Slice_Call) Run(run func(key string)) *Context_OptionInt64Slice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionInt64Slice_Call) Return(_a0 []int64) *Context_OptionInt64Slice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionInt64Slice_Call) RunAndReturn(run func(string) []int64) *Context_OptionInt64Slice_Call {
	_c.Call.Return(run)
	return _c
}

// OptionIntSlice provides a mock function with given fields: key
func (_m *Context) OptionIntSlice(key string) []int {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionIntSlice")
	}

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	return r0
}

// Context_OptionIntSlice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionIntSlice'
type Context_OptionIntSlice_Call struct {
	*mock.Call
}

// OptionIntSlice is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionIntSlice(key interface{}) *Context_OptionIntSlice_Call {
	return &Context_OptionIntSlice_Call{Call: _e.mock.On("OptionIntSlice", key)}
}

func (_c *Context_OptionIntSlice_Call) Run(run func(key string)) *Context_OptionIntSlice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionIntSlice_Call) Return(_a0 []int) *Context_OptionIntSlice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionIntSlice_Call) RunAndReturn(run func(string) []int) *Context_OptionIntSlice_Call {
	_c.Call.Return(run)
	return _c
}

// OptionSlice provides a mock function with given fields: key
func (_m *Context) OptionSlice(key string) []string {
	ret := _m.Called(key)

	if len(ret) == 0 {
		panic("no return value specified for OptionSlice")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// Context_OptionSlice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OptionSlice'
type Context_OptionSlice_Call struct {
	*mock.Call
}

// OptionSlice is a helper method to define mock.On call
//   - key string
func (_e *Context_Expecter) OptionSlice(key interface{}) *Context_OptionSlice_Call {
	return &Context_OptionSlice_Call{Call: _e.mock.On("OptionSlice", key)}
}

func (_c *Context_OptionSlice_Call) Run(run func(key string)) *Context_OptionSlice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_OptionSlice_Call) Return(_a0 []string) *Context_OptionSlice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_OptionSlice_Call) RunAndReturn(run func(string) []string) *Context_OptionSlice_Call {
	_c.Call.Return(run)
	return _c
}

// Secret provides a mock function with given fields: question, option
func (_m *Context) Secret(question string, option ...console.SecretOption) (string, error) {
	_va := make([]interface{}, len(option))
	for _i := range option {
		_va[_i] = option[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, question)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Secret")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string, ...console.SecretOption) (string, error)); ok {
		return rf(question, option...)
	}
	if rf, ok := ret.Get(0).(func(string, ...console.SecretOption) string); ok {
		r0 = rf(question, option...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, ...console.SecretOption) error); ok {
		r1 = rf(question, option...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context_Secret_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Secret'
type Context_Secret_Call struct {
	*mock.Call
}

// Secret is a helper method to define mock.On call
//   - question string
//   - option ...console.SecretOption
func (_e *Context_Expecter) Secret(question interface{}, option ...interface{}) *Context_Secret_Call {
	return &Context_Secret_Call{Call: _e.mock.On("Secret",
		append([]interface{}{question}, option...)...)}
}

func (_c *Context_Secret_Call) Run(run func(question string, option ...console.SecretOption)) *Context_Secret_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]console.SecretOption, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(console.SecretOption)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Context_Secret_Call) Return(_a0 string, _a1 error) *Context_Secret_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Context_Secret_Call) RunAndReturn(run func(string, ...console.SecretOption) (string, error)) *Context_Secret_Call {
	_c.Call.Return(run)
	return _c
}

// Spinner provides a mock function with given fields: message, option
func (_m *Context) Spinner(message string, option console.SpinnerOption) error {
	ret := _m.Called(message, option)

	if len(ret) == 0 {
		panic("no return value specified for Spinner")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, console.SpinnerOption) error); ok {
		r0 = rf(message, option)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Context_Spinner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Spinner'
type Context_Spinner_Call struct {
	*mock.Call
}

// Spinner is a helper method to define mock.On call
//   - message string
//   - option console.SpinnerOption
func (_e *Context_Expecter) Spinner(message interface{}, option interface{}) *Context_Spinner_Call {
	return &Context_Spinner_Call{Call: _e.mock.On("Spinner", message, option)}
}

func (_c *Context_Spinner_Call) Run(run func(message string, option console.SpinnerOption)) *Context_Spinner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(console.SpinnerOption))
	})
	return _c
}

func (_c *Context_Spinner_Call) Return(_a0 error) *Context_Spinner_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Context_Spinner_Call) RunAndReturn(run func(string, console.SpinnerOption) error) *Context_Spinner_Call {
	_c.Call.Return(run)
	return _c
}

// Warning provides a mock function with given fields: message
func (_m *Context) Warning(message string) {
	_m.Called(message)
}

// Context_Warning_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Warning'
type Context_Warning_Call struct {
	*mock.Call
}

// Warning is a helper method to define mock.On call
//   - message string
func (_e *Context_Expecter) Warning(message interface{}) *Context_Warning_Call {
	return &Context_Warning_Call{Call: _e.mock.On("Warning", message)}
}

func (_c *Context_Warning_Call) Run(run func(message string)) *Context_Warning_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Context_Warning_Call) Return() *Context_Warning_Call {
	_c.Call.Return()
	return _c
}

func (_c *Context_Warning_Call) RunAndReturn(run func(string)) *Context_Warning_Call {
	_c.Call.Return(run)
	return _c
}

// WithProgressBar provides a mock function with given fields: items, callback
func (_m *Context) WithProgressBar(items []interface{}, callback func(interface{}) error) ([]interface{}, error) {
	ret := _m.Called(items, callback)

	if len(ret) == 0 {
		panic("no return value specified for WithProgressBar")
	}

	var r0 []interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func([]interface{}, func(interface{}) error) ([]interface{}, error)); ok {
		return rf(items, callback)
	}
	if rf, ok := ret.Get(0).(func([]interface{}, func(interface{}) error) []interface{}); ok {
		r0 = rf(items, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func([]interface{}, func(interface{}) error) error); ok {
		r1 = rf(items, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context_WithProgressBar_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithProgressBar'
type Context_WithProgressBar_Call struct {
	*mock.Call
}

// WithProgressBar is a helper method to define mock.On call
//   - items []interface{}
//   - callback func(interface{}) error
func (_e *Context_Expecter) WithProgressBar(items interface{}, callback interface{}) *Context_WithProgressBar_Call {
	return &Context_WithProgressBar_Call{Call: _e.mock.On("WithProgressBar", items, callback)}
}

func (_c *Context_WithProgressBar_Call) Run(run func(items []interface{}, callback func(interface{}) error)) *Context_WithProgressBar_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]interface{}), args[1].(func(interface{}) error))
	})
	return _c
}

func (_c *Context_WithProgressBar_Call) Return(_a0 []interface{}, _a1 error) *Context_WithProgressBar_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Context_WithProgressBar_Call) RunAndReturn(run func([]interface{}, func(interface{}) error) ([]interface{}, error)) *Context_WithProgressBar_Call {
	_c.Call.Return(run)
	return _c
}

// NewContext creates a new instance of Context. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContext(t interface {
	mock.TestingT
	Cleanup(func())
}) *Context {
	mock := &Context{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
