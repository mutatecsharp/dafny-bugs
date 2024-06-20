// Dafny program main.dfy compiled into Go
package main

import (
	_System "System_"
	_dafny "dafny"
	os "os"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ _System.Dummy__

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_module.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) Abs(x _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return (_dafny.IntOfInt64(-1)).Times(x)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeIndex(x _dafny.Int, length _dafny.Int) _dafny.Int {
	if (x).Sign() == -1 {
		return _dafny.Zero
	} else if (x).Cmp(length) >= 0 {
		return (x).Modulo(length)
	} else {
		return x
	}
}
func (_static *CompanionStruct_Default___) SafeDivisionInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).DivBy(x2)
	}
}
func (_static *CompanionStruct_Default___) SafeModuloInt(x1 _dafny.Int, x2 _dafny.Int) _dafny.Int {
	if (x2).Sign() == 0 {
		return x1
	} else {
		return (x1).Modulo(x2)
	}
}
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 bool, p2 D0, globalState *GlobalState) _dafny.Map {
	var _source0 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
	_ = _source0
	var _0___mcc_h0 _dafny.Map = _source0
	_ = _0___mcc_h0
	var _1_cf18 _dafny.Map = _0___mcc_h0
	_ = _1_cf18
	return func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(538), _dafny.IntOfInt64(46))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _2_v0 _dafny.Int
			_2_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(538)).Cmp(_2_v0) <= 0) && ((_2_v0).Cmp(_dafny.IntOfInt64(46)) < 0) {
				_coll0.Add((_2_v0).Plus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kdylahw")).Cardinality())), _dafny.SetOf(true))
			}
		}
		return _coll0.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm1(p0 D0, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) bool {
	if false {
		return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, !(false)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).IsProperSubsetOf(_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	} else {
		return !(true) || (false)
	}
}
func (_static *CompanionStruct_Default___) Fm2(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("jkyhcsar")
}
func (_static *CompanionStruct_Default___) Fm4(p0 _dafny.Map, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source1 D17 = Companion_D17_.Create_DC41_(_dafny.IntOfInt64(-267), !(true))
	_ = _source1
	if _source1.Is_DC41() {
		var _3___mcc_h0 _dafny.Int = _source1.Get_().(D17_DC41).Cf77
		_ = _3___mcc_h0
		var _4___mcc_h1 bool = _source1.Get_().(D17_DC41).Cf78
		_ = _4___mcc_h1
		var _5_cf78 bool = _4___mcc_h1
		_ = _5_cf78
		var _6_cf77 _dafny.Int = _3___mcc_h0
		_ = _6_cf77
		return _dafny.CodePoint('b')
	} else if _source1.Is_DC40() {
		var _7___mcc_h2 _dafny.Map = _source1.Get_().(D17_DC40).Cf76
		_ = _7___mcc_h2
		var _8_cf76 _dafny.Map = _7___mcc_h2
		_ = _8_cf76
		return _dafny.CodePoint('u')
	} else {
		var _9___mcc_h3 D17 = _source1.Get_().(D17_DC42).Cf79
		_ = _9___mcc_h3
		var _10_cf79 D17 = _9___mcc_h3
		_ = _10_cf79
		return _dafny.CodePoint('y')
	}
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(548), _dafny.IntOfInt64(-508))
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.CodePoint, p1 bool, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.CodePoint {
	var _source2 D14 = Companion_D14_.Create_DC36_(true, _dafny.IntOfInt64(110), (_dafny.Zero).Minus((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.SetOf(true, !(!(false)))).Cardinality())).Cardinality())).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(888), true))
	_ = _source2
	if _source2.Is_DC36() {
		var _11___mcc_h0 bool = _source2.Get_().(D14_DC36).Cf68
		_ = _11___mcc_h0
		var _12___mcc_h1 _dafny.Int = _source2.Get_().(D14_DC36).Cf69
		_ = _12___mcc_h1
		var _13___mcc_h2 _dafny.Int = _source2.Get_().(D14_DC36).Cf70
		_ = _13___mcc_h2
		var _14___mcc_h3 _dafny.Map = _source2.Get_().(D14_DC36).Cf71
		_ = _14___mcc_h3
		var _15_cf71 _dafny.Map = _14___mcc_h3
		_ = _15_cf71
		var _16_cf70 _dafny.Int = _13___mcc_h2
		_ = _16_cf70
		var _17_cf69 _dafny.Int = _12___mcc_h1
		_ = _17_cf69
		var _18_cf68 bool = _11___mcc_h0
		_ = _18_cf68
		return _dafny.CodePoint('q')
	} else {
		var _19___mcc_h4 _dafny.Array = _source2.Get_().(D14_DC35).Cf67
		_ = _19___mcc_h4
		var _20_cf67 _dafny.Array = _19___mcc_h4
		_ = _20_cf67
		return _dafny.CodePoint('h')
	}
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Set {
	var _source3 D1 = Companion_D1_.Create_DC5_()
	_ = _source3
	if _source3.Is_DC5() {
		return (_dafny.SetOf(_dafny.MultiSetOf(!(true), !(false)), _dafny.MultiSetOf(true, false))).Difference(func() _dafny.Set {
			var _coll1 = _dafny.NewBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _dafny.IntOfInt64(213))).Keys().Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _21_v0 _dafny.MultiSet
				_21_v0 = interface{}(_compr_1).(_dafny.MultiSet)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(true), _dafny.IntOfInt64(213))).Contains(_21_v0) {
					_coll1.Add(_21_v0)
				}
			}
			return _coll1.ToSet()
		}())
	} else {
		var _22___mcc_h0 _dafny.Sequence = _source3.Get_().(D1_DC4).Cf11
		_ = _22___mcc_h0
		var _23_cf11 _dafny.Sequence = _22___mcc_h0
		_ = _23_cf11
		return (_dafny.SetOf(_dafny.MultiSetOf(true, true, false, true, false))).Difference(_dafny.SetOf(_dafny.MultiSetOf(true), _dafny.MultiSetOf(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm10(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return ((_dafny.SetOf(_dafny.IntOfInt64(-989))).Union(_dafny.SetOf(_dafny.IntOfInt64(922)))).Difference((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality())))).Cardinality())).Union(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(17), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bbpgrxd")).Cardinality()))).Cardinality()), false)).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(791), !(!(true)))).Cardinality(), _dafny.IntOfInt64(854))).Cardinality(), _dafny.IntOfInt64(439))))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.CodePoint {
	var _source4 D1 = Companion_D1_.Create_DC4_(_dafny.SeqOf((_dafny.SetOf(false)).Cardinality()))
	_ = _source4
	if _source4.Is_DC5() {
		if true {
			return _dafny.CodePoint('u')
		} else {
			return _dafny.CodePoint('u')
		}
	} else {
		var _24___mcc_h0 _dafny.Sequence = _source4.Get_().(D1_DC4).Cf11
		_ = _24___mcc_h0
		var _25_cf11 _dafny.Sequence = _24___mcc_h0
		_ = _25_cf11
		if !(!(false)) {
			return _dafny.CodePoint('r')
		} else {
			return _dafny.CodePoint('b')
		}
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.CodePoint, p1 _dafny.Sequence, globalState *GlobalState) D0 {
	var _source5 D5 = Companion_D5_.Create_DC13_(true, !(!(false)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-59))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_26_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('i')
	})), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(921))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_27_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	}))).Cardinality()))
	_ = _source5
	if _source5.Is_DC12() {
		var _28___mcc_h0 _dafny.Int = _source5.Get_().(D5_DC12).Cf20
		_ = _28___mcc_h0
		var _29_cf20 _dafny.Int = _28___mcc_h0
		_ = _29_cf20
		return Companion_D0_.Create_DC1_(_29_cf20, false, _29_cf20, !(true))
	} else if _source5.Is_DC13() {
		var _30___mcc_h1 bool = _source5.Get_().(D5_DC13).Cf21
		_ = _30___mcc_h1
		var _31___mcc_h2 bool = _source5.Get_().(D5_DC13).Cf22
		_ = _31___mcc_h2
		var _32___mcc_h3 _dafny.Sequence = _source5.Get_().(D5_DC13).Cf23
		_ = _32___mcc_h3
		var _33___mcc_h4 _dafny.Int = _source5.Get_().(D5_DC13).Cf24
		_ = _33___mcc_h4
		var _34_cf24 _dafny.Int = _33___mcc_h4
		_ = _34_cf24
		var _35_cf23 _dafny.Sequence = _32___mcc_h3
		_ = _35_cf23
		var _36_cf22 bool = _31___mcc_h2
		_ = _36_cf22
		var _37_cf21 bool = _30___mcc_h1
		_ = _37_cf21
		return Companion_D0_.Create_DC1_(_34_cf24, _36_cf22, _dafny.IntOfInt64(842), _36_cf22)
	} else if _source5.Is_DC11() {
		var _38___mcc_h5 _dafny.Set = _source5.Get_().(D5_DC11).Cf19
		_ = _38___mcc_h5
		var _39_cf19 _dafny.Set = _38___mcc_h5
		_ = _39_cf19
		return Companion_D0_.Create_DC1_(_dafny.IntOfInt64(651), !(false), _dafny.IntOfInt64(-746), false)
	} else {
		var _40___mcc_h6 D5 = _source5.Get_().(D5_DC14).Cf25
		_ = _40___mcc_h6
		var _41_cf25 D5 = _40___mcc_h6
		_ = _41_cf25
		return Companion_D0_.Create_DC1_(_dafny.IntOfInt64(114), true, _dafny.IntOfInt64(-419), !(true))
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if !(!((func() bool {
		if true {
			return true
		}
		return true
	})())) {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_D17_.Create_DC41_(_dafny.IntOfInt64(377), true)).Dtor_cf78(), true)
	} else {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
	}
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Int {
	return (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(false, false)).Cardinality()), _dafny.IntOfInt64(183))).Minus((_dafny.IntOfInt64(879)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("raqgm")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm19(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Map, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(((_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(-412))).Cardinality())).Intersection(func() _dafny.Set {
		var _coll2 = _dafny.NewBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(180), _dafny.IntOfInt64(64))); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _42_v0 _dafny.Int
			_42_v0 = interface{}(_compr_2).(_dafny.Int)
			if ((_dafny.IntOfInt64(180)).Cmp(_42_v0) <= 0) && ((_42_v0).Cmp(_dafny.IntOfInt64(64)) < 0) {
				_coll2.Add((_42_v0).Times(_dafny.IntOfInt64(170)))
			}
		}
		return _coll2.ToSet()
	}())).Cardinality(), _dafny.CodePoint('f'), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("li")).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("mhhv")).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfInt64(938)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(288))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_43_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(false, false, true, true)).Cardinality()))).Cardinality())
	})))).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm21(p0 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(!(true))), _dafny.SeqOf(!(!(false))))
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Set {
	return func() _dafny.Set {
		var _coll3 = _dafny.NewBuilder()
		_ = _coll3
		for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-362), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ksdltj")).Cardinality())))).Keys().Elements()); ; {
			_compr_3, _ok3 := _iter3()
			if !_ok3 {
				break
			}
			var _44_v0 _dafny.Int
			_44_v0 = interface{}(_compr_3).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-362), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ksdltj")).Cardinality())))).Contains(_44_v0) {
				_coll3.Add((_44_v0).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(428), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bdnraex"), _dafny.UnicodeSeqOfUtf8Bytes("lktqteah"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(589))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg3 _dafny.Int) interface{} {
						return coer3(arg3)
					}
				}(func(_45_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				})))).Cardinality()))))
			}
		}
		return _coll3.ToSet()
	}()
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	if (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality())).Cmp(_dafny.IntOfInt64(814)) > 0 {
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-666), (Companion_D17_.Create_DC41_(_dafny.IntOfInt64(94), true)).Dtor_cf77())
	} else {
		return Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("hhew")).Cardinality())), _dafny.IntOfInt64(885))
	}
}
func (_static *CompanionStruct_Default___) Fm25(p0 _dafny.Int, p1 _dafny.Map, p2 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.IntOfInt64(-408), Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality()), _dafny.IntOfInt64(572)), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false))).Cardinality())), _dafny.IntOfInt64(547))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false)).Intersection(_dafny.SetOf(false, true, false, true))
}
func (_static *CompanionStruct_Default___) Fm27(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(_dafny.SeqOf(!(false), false), _dafny.SeqOf(true, false), _dafny.SeqOf(false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Int, p2 bool, globalState *GlobalState) _dafny.CodePoint {
	if !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC19_(_dafny.MultiSetOf(false, !(!(false))), _dafny.IntOfInt64(330)), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !(!(true)))).Cardinality())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC19_(_dafny.MultiSetOf(!(true)), (func() _dafny.Set {
		var _coll4 = _dafny.NewBuilder()
		_ = _coll4
		for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-713), _dafny.IntOfInt64(-163))); ; {
			_compr_4, _ok4 := _iter4()
			if !_ok4 {
				break
			}
			var _46_v0 _dafny.Int
			_46_v0 = interface{}(_compr_4).(_dafny.Int)
			if ((_dafny.IntOfInt64(-713)).Cmp(_46_v0) <= 0) && ((_46_v0).Cmp(_dafny.IntOfInt64(-163)) < 0) {
				_coll4.Add(Companion_Default___.SafeDivisionInt(_46_v0, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
			}
		}
		return _coll4.ToSet()
	}()).Cardinality()), (_dafny.MultiSetOf(_dafny.IntOfInt64(26), _dafny.IntOfInt64(942))).Cardinality())) {
		if false {
			return _dafny.CodePoint('x')
		} else {
			return _dafny.CodePoint('r')
		}
	} else if true {
		return _dafny.CodePoint('q')
	} else {
		return _dafny.CodePoint('p')
	}
}
func (_static *CompanionStruct_Default___) Fm29(p0 bool, p1 _dafny.Int, globalState *GlobalState) D6 {
	var _source6 D13 = Companion_D13_.Create_DC31_(func() _dafny.Map {
		var _coll5 = _dafny.NewMapBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality())).Elements()); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _47_v0 _dafny.Int
			_47_v0 = interface{}(_compr_5).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(false))).Cardinality()), _47_v0) {
				_coll5.Add((_47_v0).Times(_dafny.IntOfInt64(-611)), false)
			}
		}
		return _coll5.ToMap()
	}())
	_ = _source6
	if _source6.Is_DC32() {
		var _48___mcc_h0 _dafny.Int = _source6.Get_().(D13_DC32).Cf56
		_ = _48___mcc_h0
		var _49___mcc_h1 bool = _source6.Get_().(D13_DC32).Cf57
		_ = _49___mcc_h1
		var _50___mcc_h2 bool = _source6.Get_().(D13_DC32).Cf58
		_ = _50___mcc_h2
		var _51___mcc_h3 bool = _source6.Get_().(D13_DC32).Cf59
		_ = _51___mcc_h3
		var _52___mcc_h4 bool = _source6.Get_().(D13_DC32).Cf60
		_ = _52___mcc_h4
		var _53_cf60 bool = _52___mcc_h4
		_ = _53_cf60
		var _54_cf59 bool = _51___mcc_h3
		_ = _54_cf59
		var _55_cf58 bool = _50___mcc_h2
		_ = _55_cf58
		var _56_cf57 bool = _49___mcc_h1
		_ = _56_cf57
		var _57_cf56 _dafny.Int = _48___mcc_h0
		_ = _57_cf56
		return Companion_D6_.Create_DC16_((_dafny.Zero).Minus(_57_cf56))
	} else if _source6.Is_DC33() {
		var _58___mcc_h5 bool = _source6.Get_().(D13_DC33).Cf61
		_ = _58___mcc_h5
		var _59___mcc_h6 _dafny.Int = _source6.Get_().(D13_DC33).Cf62
		_ = _59___mcc_h6
		var _60___mcc_h7 _dafny.Int = _source6.Get_().(D13_DC33).Cf63
		_ = _60___mcc_h7
		var _61___mcc_h8 bool = _source6.Get_().(D13_DC33).Cf64
		_ = _61___mcc_h8
		var _62___mcc_h9 *C3 = _source6.Get_().(D13_DC33).Cf65
		_ = _62___mcc_h9
		var _63_cf65 *C3 = _62___mcc_h9
		_ = _63_cf65
		var _64_cf64 bool = _61___mcc_h8
		_ = _64_cf64
		var _65_cf63 _dafny.Int = _60___mcc_h7
		_ = _65_cf63
		var _66_cf62 _dafny.Int = _59___mcc_h6
		_ = _66_cf62
		var _67_cf61 bool = _58___mcc_h5
		_ = _67_cf61
		return Companion_D6_.Create_DC16_((_63_cf65).F9())
	} else if _source6.Is_DC31() {
		var _68___mcc_h10 _dafny.Map = _source6.Get_().(D13_DC31).Cf55
		_ = _68___mcc_h10
		var _69_cf55 _dafny.Map = _68___mcc_h10
		_ = _69_cf55
		return Companion_D6_.Create_DC16_((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-504))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_70_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality())))
	} else {
		var _71___mcc_h11 D13 = _source6.Get_().(D13_DC34).Cf66
		_ = _71___mcc_h11
		var _72_cf66 D13 = _71___mcc_h11
		_ = _72_cf66
		if false {
			return Companion_D6_.Create_DC16_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(33), false)).Cardinality())
		} else {
			return Companion_D6_.Create_DC16_(_dafny.IntOfInt64(746))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm30(p0 _dafny.Sequence, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(true)
}
func (_static *CompanionStruct_Default___) Fm31(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.MultiSet, globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC15_((_dafny.SetOf(true)).Union(_dafny.SetOf(true)))
}
func (_static *CompanionStruct_Default___) Fm32(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC3_(Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(-345))).Cardinality(), _dafny.SetOf(true))))
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) D1 {
	var _source7 D17 = Companion_D17_.Create_DC41_((_dafny.MultiSetOf(true)).Cardinality(), false)
	_ = _source7
	if _source7.Is_DC41() {
		var _73___mcc_h0 _dafny.Int = _source7.Get_().(D17_DC41).Cf77
		_ = _73___mcc_h0
		var _74___mcc_h1 bool = _source7.Get_().(D17_DC41).Cf78
		_ = _74___mcc_h1
		var _75_cf78 bool = _74___mcc_h1
		_ = _75_cf78
		var _76_cf77 _dafny.Int = _73___mcc_h0
		_ = _76_cf77
		if _75_cf78 {
			return Companion_D1_.Create_DC4_(_dafny.SeqOf(_76_cf77, _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))
		} else {
			return Companion_D1_.Create_DC4_(_dafny.SeqOf(_76_cf77, _76_cf77, _76_cf77))
		}
	} else if _source7.Is_DC40() {
		var _77___mcc_h2 _dafny.Map = _source7.Get_().(D17_DC40).Cf76
		_ = _77___mcc_h2
		var _78_cf76 _dafny.Map = _77___mcc_h2
		_ = _78_cf76
		return Companion_D1_.Create_DC4_(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(458))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg5 _dafny.Int) interface{} {
				return coer5(arg5)
			}
		}(func(_79_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}))).Cardinality()), (_dafny.Zero).Minus((_dafny.MultiSetOf(false, true, true, false, false)).Cardinality())))
	} else {
		var _80___mcc_h3 D17 = _source7.Get_().(D17_DC42).Cf79
		_ = _80___mcc_h3
		var _81_cf79 D17 = _80___mcc_h3
		_ = _81_cf79
		return Companion_D1_.Create_DC4_(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(128), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), _dafny.IntOfInt64(321))).Cardinality(), _dafny.IntOfInt64(658), _dafny.IntOfInt64(195), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vhqc")).Cardinality()), _dafny.IntOfInt64(888)))
	}
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(func() _dafny.Map {
		var _coll6 = _dafny.NewMapBuilder()
		_ = _coll6
		for _iter6 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(394), true), func() _dafny.Map {
			var _coll7 = _dafny.NewMapBuilder()
			_ = _coll7
			for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(246), _dafny.IntOfInt64(118))); ; {
				_compr_7, _ok7 := _iter7()
				if !_ok7 {
					break
				}
				var _82_v1 _dafny.Int
				_82_v1 = interface{}(_compr_7).(_dafny.Int)
				if ((_dafny.IntOfInt64(246)).Cmp(_82_v1) <= 0) && ((_82_v1).Cmp(_dafny.IntOfInt64(118)) < 0) {
					_coll7.Add((_82_v1).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), !(true))
				}
			}
			return _coll7.ToMap()
		}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(70), !(false)))).Elements()); ; {
			_compr_6, _ok6 := _iter6()
			if !_ok6 {
				break
			}
			var _83_v0 _dafny.Map
			_83_v0 = interface{}(_compr_6).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(394), true), func() _dafny.Map {
				var _coll8 = _dafny.NewMapBuilder()
				_ = _coll8
				for _iter8 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(246), _dafny.IntOfInt64(118))); ; {
					_compr_8, _ok8 := _iter8()
					if !_ok8 {
						break
					}
					var _82_v1 _dafny.Int
					_82_v1 = interface{}(_compr_8).(_dafny.Int)
					if ((_dafny.IntOfInt64(246)).Cmp(_82_v1) <= 0) && ((_82_v1).Cmp(_dafny.IntOfInt64(118)) < 0) {
						_coll8.Add((_82_v1).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), !(true))
					}
				}
				return _coll8.ToMap()
			}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(70), !(false))), _83_v0) {
				_coll6.Add(_83_v0, _dafny.IntOfInt64(999))
			}
		}
		return _coll6.ToMap()
	}()).Contains((Companion_D13_.Create_DC31_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.SeqOf(false), _dafny.SeqOf(true, false))).Cardinality()), true))).Dtor_cf55()), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-291))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_84_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(720)
	}))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(-82), (func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(133), _dafny.IntOfInt64(439))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _85_v2 _dafny.Int
			_85_v2 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(133)).Cmp(_85_v2) <= 0) && ((_85_v2).Cmp(_dafny.IntOfInt64(439)) < 0) {
				_coll9.Add((_85_v2).Minus(_dafny.IntOfInt64(214)), (_dafny.SetOf(false)).Cardinality())
			}
		}
		return _coll9.ToMap()
	}()).Cardinality())).Cardinality(), !(true))).Cardinality())).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(735), _dafny.IntOfInt64(-829))); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _86_v0 _dafny.Int
			_86_v0 = interface{}(_compr_10).(_dafny.Int)
			if ((_dafny.IntOfInt64(735)).Cmp(_86_v0) <= 0) && ((_86_v0).Cmp(_dafny.IntOfInt64(-829)) < 0) {
				_coll10.Add(Companion_Default___.SafeModuloInt(_86_v0, _dafny.IntOfInt64(520)), false)
			}
		}
		return _coll10.ToMap()
	}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality()), false))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(679))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_87_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('c')
	}))).Cardinality())), true)).Merge(func() _dafny.Map {
		var _coll11 = _dafny.NewMapBuilder()
		_ = _coll11
		for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(120), _dafny.IntOfInt64(-129))); ; {
			_compr_11, _ok11 := _iter11()
			if !_ok11 {
				break
			}
			var _88_v1 _dafny.Int
			_88_v1 = interface{}(_compr_11).(_dafny.Int)
			if ((_dafny.IntOfInt64(120)).Cmp(_88_v1) <= 0) && ((_88_v1).Cmp(_dafny.IntOfInt64(-129)) < 0) {
				_coll11.Add((_88_v1).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())), false)
			}
		}
		return _coll11.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm36(p0 bool, p1 bool, p2 bool, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return func() _dafny.Map {
		var _coll12 = _dafny.NewMapBuilder()
		_ = _coll12
		for _iter12 := _dafny.Iterate((_dafny.SetOf(_dafny.CodePoint('u'), _dafny.CodePoint('g'), _dafny.CodePoint('k'))).Elements()); ; {
			_compr_12, _ok12 := _iter12()
			if !_ok12 {
				break
			}
			var _89_v0 _dafny.CodePoint
			_89_v0 = interface{}(_compr_12).(_dafny.CodePoint)
			if (_dafny.SetOf(_dafny.CodePoint('u'), _dafny.CodePoint('g'), _dafny.CodePoint('k'))).Contains(_89_v0) {
				_coll12.Add(_89_v0, (_dafny.SetOf(true, true)).Contains(!(true)))
			}
		}
		return _coll12.ToMap()
	}()
}
func (_static *CompanionStruct_Default___) Fm37(p0 bool, p1 _dafny.Int, p2 _dafny.MultiSet, globalState *GlobalState) D12 {
	return Companion_D12_.Create_DC28_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-806)))
}
func (_static *CompanionStruct_Default___) Fm38(p0 bool, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.SeqOf(false))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm39(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (Companion_D16_.Create_DC39_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ok")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(57))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_90_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('y')
	}))).Cardinality()), (_dafny.MultiSetOf(true)).Cardinality(), _dafny.IntOfInt64(-821))).Cardinality()), _dafny.IntOfInt64(181)))).Dtor_cf75()
}
func (_static *CompanionStruct_Default___) Fm40(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-31), _dafny.IntOfInt64(889))), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true))).Cardinality())), _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(677)), false)).Cardinality())))
}
func (_static *CompanionStruct_Default___) M7(p0 _dafny.Int, globalState *GlobalState) {
	var _91_v0 _dafny.Sequence
	_ = _91_v0
	_91_v0 = _dafny.SeqOf(p0)
	var _92_v1 *C3
	_ = _92_v1
	var _nw0 *C3 = New_C3_()
	_ = _nw0
	_nw0.Ctor__(_dafny.IntOfInt64(823), _91_v0, (_dafny.Zero).Minus(p0))
	_92_v1 = _nw0
	var _93_v2 bool
	_ = _93_v2
	_93_v2 = false
	var _94_v3 _dafny.Map
	_ = _94_v3
	_94_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_93_v2, (_92_v1).F9())
	var _95_v4 _dafny.Map
	_ = _95_v4
	_95_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _94_v3)
	var _96_v5 _dafny.Sequence
	_ = _96_v5
	_96_v5 = _dafny.SeqOf(_93_v2, _93_v2)
	var _97_v6 _dafny.Sequence
	_ = _97_v6
	_97_v6 = _dafny.SeqOf(!((_96_v5).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_96_v5).Cardinality()))).Uint32()).(bool)))
	var _98_v7 D13
	_ = _98_v7
	_98_v7 = Companion_D13_.Create_DC33_(_93_v2, (_95_v4).Cardinality(), _dafny.IntOfUint32((_96_v5).Cardinality()), _93_v2, _92_v1)
	var _99_v8 _dafny.Array
	_ = _99_v8
	var _nwElement0_0 *C3 = _92_v1
	_ = _nwElement0_0
	var _nw1 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(10))
	_ = _nw1
	(_nw1).ArraySet1(_nwElement0_0, 0)
	(_nw1).ArraySet1(_92_v1, 1)
	(_nw1).ArraySet1((func() *C3 {
		if _93_v2 {
			return _92_v1
		}
		return _92_v1
	})(), 2)
	(_nw1).ArraySet1((_98_v7).Dtor_cf65(), 3)
	(_nw1).ArraySet1(_92_v1, 4)
	(_nw1).ArraySet1(_92_v1, 5)
	(_nw1).ArraySet1(_92_v1, 6)
	(_nw1).ArraySet1(_92_v1, 7)
	(_nw1).ArraySet1(_92_v1, 8)
	(_nw1).ArraySet1(_92_v1, 9)
	_99_v8 = _nw1
	var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_99_v8), 0))
	_ = _index0
	(_99_v8).ArraySet1(_92_v1, (_index0).Int())
	var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(551), _dafny.ArrayLen((_99_v8), 0))
	_ = _index1
	(_99_v8).ArraySet1(_92_v1, (_index1).Int())
	_97_v6 = _97_v6
	var _100_v9 _dafny.Array
	_ = _100_v9
	var _nw2 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(26))
	_ = _nw2
	_100_v9 = _nw2
	var _101_v10 *C0
	_ = _101_v10
	var _nw3 *C0 = New_C0_()
	_ = _nw3
	_nw3.Ctor__(!(!(_93_v2)))
	_101_v10 = _nw3
	var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_100_v9), 0))
	_ = _index2
	(_100_v9).ArraySet1(_101_v10, (_index2).Int())
	var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(796), _dafny.ArrayLen((_100_v9), 0))
	_ = _index3
	var _nw4 *C0 = New_C0_()
	_ = _nw4
	_nw4.Ctor__((_101_v10).F13())
	(_100_v9).ArraySet1(_nw4, (_index3).Int())
	var _102_v11 *C0
	_ = _102_v11
	var _nw5 *C0 = New_C0_()
	_ = _nw5
	_nw5.Ctor__((_101_v10).F13())
	_102_v11 = _nw5
	var _103_v12 _dafny.Array
	_ = _103_v12
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(5)
	_ = _len0_0
	var _nw6 _dafny.Array
	_ = _nw6
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw6 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) bool = (func(_104_v10 *C0) func(_dafny.Int) bool {
			return func(_105_i0 _dafny.Int) bool {
				return (_104_v10).F13()
			}
		})(_101_v10)
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw6 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw6).ArraySet1(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw6).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_103_v12 = _nw6
	var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
	_ = _index4
	(_103_v12).ArraySet1(_93_v2, (_index4).Int())
	var _106_v13 D5
	_ = _106_v13
	_106_v13 = Companion_D5_.Create_DC12_(p0)
	var _pat_let_tv0 = _92_v1
	_ = _pat_let_tv0
	var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
	_ = _index5
	var _rhs0 bool = (_102_v11).F13()
	_ = _rhs0
	var _rhs1 bool = !(_93_v2) || (_93_v2)
	_ = _rhs1
	var _rhs2 D5 = func(_pat_let0_0 D5) D5 {
		return func(_107_dt__update__tmp_h0 D5) D5 {
			return func(_pat_let1_0 _dafny.Int) D5 {
				return func(_108_dt__update_hcf20_h0 _dafny.Int) D5 {
					return Companion_D5_.Create_DC12_(_108_dt__update_hcf20_h0)
				}(_pat_let1_0)
			}((_pat_let_tv0).F9())
		}(_pat_let0_0)
	}(_106_v13)
	_ = _rhs2
	var _lhs0 _dafny.Array = _103_v12
	_ = _lhs0
	var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
	_ = _lhs1
	var _lhs2 *GlobalState = globalState
	_ = _lhs2
	(_lhs0).ArraySet1(_rhs0, (_lhs1).Int())
	_lhs2.F5 = _rhs1
	_106_v13 = _rhs2
	var _hi0 _dafny.Int = (_92_v1).F9()
	_ = _hi0
	for _109_i1 := p0; _109_i1.Cmp(_hi0) < 0; _109_i1 = _109_i1.Plus(_dafny.One) {
		var _110_v14 _dafny.Sequence
		_ = _110_v14
		_110_v14 = _dafny.UnicodeSeqOfUtf8Bytes("ode")
		var _111_v15 _dafny.Map
		_ = _111_v15
		_111_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_92_v1).F9(), (func() _dafny.Sequence {
			if (_103_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))).Int()).(bool) {
				return _110_v14
			}
			return _110_v14
		})())
		_111_v15 = (_111_v15).Update(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(479))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg9 _dafny.Int) interface{} {
				return coer9(arg9)
			}
		}((func(_112_v1 *C3) func(_dafny.Int) _dafny.Int {
			return func(_113_i2 _dafny.Int) _dafny.Int {
				return (_112_v1).F9()
			}
		})(_92_v1)))).Cardinality()), (_92_v1).F9()), _110_v14)
		var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
		_ = _index6
		(_103_v12).ArraySet1((_96_v5).Select((Companion_Default___.SafeIndex(_109_i1, _dafny.IntOfUint32((_96_v5).Cardinality()))).Uint32()).(bool), (_index6).Int())
		if !(_93_v2) {
			var _114_v16 _dafny.Map
			_ = _114_v16
			_114_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_103_v12, (_92_v1).Fm16(_dafny.MultiSetOf((_101_v10).F13()), globalState))
			_114_v16 = _114_v16
			_93_v2 = !(_93_v2)
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
			_ = _index7
			(_103_v12).ArraySet1((_101_v10).F13(), (_index7).Int())
			var _115_v17 *C1
			_ = _115_v17
			var _nw7 *C1 = New_C1_()
			_ = _nw7
			_nw7.Ctor__()
			_115_v17 = _nw7
			var _116_v18 _dafny.MultiSet
			_ = _116_v18
			_116_v18 = _dafny.MultiSetOf(_dafny.IntOfInt64(442), (_dafny.Zero).Minus((_dafny.Zero).Minus((_92_v1).F9())), (_dafny.Zero).Minus(p0))
			var _117_v19 _dafny.Sequence
			_ = _117_v19
			_117_v19 = _dafny.SeqOf(_116_v18, _dafny.MultiSetFromSeq((_92_v1).F10()))
			var _118_v20 _dafny.Set
			_ = _118_v20
			_118_v20 = _dafny.SetOf((_102_v11).F13(), (_102_v11).F13())
			var _119_v21 _dafny.Map
			_ = _119_v21
			_119_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_117_v19).Select((Companion_Default___.SafeIndex((_118_v20).Cardinality(), _dafny.IntOfUint32((_117_v19).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsProperSubsetOf(_116_v18), (_101_v10).F13())
			var _120_v22 _dafny.Map
			_ = _120_v22
			_120_v22 = _119_v21
			_119_v21 = (_119_v21).Merge((_120_v22))
		} else {
			var _121_v23 _dafny.Array
			_ = _121_v23
			var _nw8 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
			_ = _nw8
			_121_v23 = _nw8
			var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_121_v23), 0))
			_ = _index8
			(_121_v23).ArraySet1(p0, (_index8).Int())
			var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_121_v23), 0))
			_ = _index9
			(_121_v23).ArraySet1(_109_i1, (_index9).Int())
			var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
			_ = _index10
			(_103_v12).ArraySet1((_101_v10).F13(), (_index10).Int())
			var _122_v24 _dafny.CodePoint
			_ = _122_v24
			_122_v24 = _dafny.CodePoint('w')
			var _123_v25 D0
			_ = _123_v25
			_123_v25 = Companion_D0_.Create_DC2_((_dafny.Zero).Minus((_dafny.Zero).Minus(_109_i1)), _122_v24, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_97_v6, (Companion_Default___.SafeIndex((_121_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_121_v23), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_97_v6).Cardinality()))).Uint32(), (_101_v10).F13())).Cardinality())), _109_i1, (_92_v1).F9())
			_122_v24 = (_123_v25).Dtor_cf6()
			var _124_v26 _dafny.Set
			_ = _124_v26
			_124_v26 = _dafny.SetOf(!((_101_v10).F13()) || ((_101_v10).F13()), _dafny.Companion_Sequence_.IsPrefixOf(_110_v14, _110_v14), (_102_v11).F13())
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
			_ = _index11
			var _rhs3 _dafny.Array = _103_v12
			_ = _rhs3
			var _rhs4 bool = (_102_v11).F13()
			_ = _rhs4
			var _rhs5 _dafny.Set = (_124_v26).Difference((_124_v26).Intersection(_124_v26))
			_ = _rhs5
			var _lhs3 _dafny.Array = _103_v12
			_ = _lhs3
			var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(177), _dafny.ArrayLen((_103_v12), 0))
			_ = _lhs4
			_103_v12 = _rhs3
			(_lhs3).ArraySet1(_rhs4, (_lhs4).Int())
			_124_v26 = _rhs5
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(532), _dafny.ArrayLen((_121_v23), 0))
			_ = _index12
			(_121_v23).ArraySet1((_92_v1).F9(), (_index12).Int())
		}
		var _125_v27 _dafny.MultiSet
		_ = _125_v27
		_125_v27 = _dafny.MultiSetOf(_103_v12, _103_v12, _103_v12, _103_v12, _103_v12)
		(globalState).F1 = !(((_125_v27).Union(_125_v27)).IsSubsetOf(_125_v27))
	}
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _126_v0 _dafny.Sequence
	_ = _126_v0
	_126_v0 = _dafny.UnicodeSeqOfUtf8Bytes("oeg")
	var _127_globalState *GlobalState
	_ = _127_globalState
	var _nw9 *GlobalState = New_GlobalState_()
	_ = _nw9
	_nw9.Ctor__(_dafny.IntOfInt64(-971), false, _dafny.IntOfInt64(560), true, _126_v0, false)
	_127_globalState = _nw9
	var _128_v1 bool
	_ = _128_v1
	_128_v1 = false
	var _129_v2 _dafny.Set
	_ = _129_v2
	_129_v2 = _dafny.SetOf(_128_v1)
	var _130_v3 _dafny.Map
	_ = _130_v3
	_130_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _129_v2)
	var _131_v5 _dafny.Int
	_ = _131_v5
	_131_v5 = _dafny.IntOfInt64(371)
	var _132_v7 _dafny.CodePoint
	_ = _132_v7
	_132_v7 = _dafny.CodePoint('g')
	var _133_v8 _dafny.MultiSet
	_ = _133_v8
	_133_v8 = _dafny.MultiSetOf(_131_v5)
	var _134_v9 D0
	_ = _134_v9
	_134_v9 = Companion_D0_.Create_DC2_(_131_v5, _132_v7, _131_v5, _131_v5, (_133_v8).Cardinality())
	var _pat_let_tv1 = _131_v5
	_ = _pat_let_tv1
	var _135_v10 _dafny.Array
	_ = _135_v10
	var _nwElement0_1 _dafny.Map = (func() _dafny.Map {
		if true {
			return _130_v3
		}
		return _130_v3
	})()
	_ = _nwElement0_1
	var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(17))
	_ = _nw10
	(_nw10).ArraySet1(_nwElement0_1, 0)
	(_nw10).ArraySet1(_130_v3, 1)
	(_nw10).ArraySet1(_130_v3, 2)
	(_nw10).ArraySet1(_130_v3, 3)
	(_nw10).ArraySet1(_130_v3, 4)
	(_nw10).ArraySet1((_130_v3).Merge(_130_v3), 5)
	(_nw10).ArraySet1(_130_v3, 6)
	(_nw10).ArraySet1(_130_v3, 7)
	(_nw10).ArraySet1(_130_v3, 8)
	(_nw10).ArraySet1(_130_v3, 9)
	(_nw10).ArraySet1(func() _dafny.Map {
		var _coll13 = _dafny.NewMapBuilder()
		_ = _coll13
		for _iter13 := _dafny.Iterate((_dafny.MultiSetOf(_131_v5)).Elements()); ; {
			_compr_13, _ok13 := _iter13()
			if !_ok13 {
				break
			}
			var _136_v4 _dafny.Int
			_136_v4 = interface{}(_compr_13).(_dafny.Int)
			if (_dafny.MultiSetOf(_131_v5)).Contains(_136_v4) {
				_coll13.Add((_136_v4).Plus(_131_v5), _129_v2)
			}
		}
		return _coll13.ToMap()
	}(), 10)
	(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _129_v2), 11)
	(_nw10).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _dafny.SetOf(_128_v1, _128_v1)), 12)
	(_nw10).ArraySet1((_130_v3).Merge(_130_v3), 13)
	(_nw10).ArraySet1((Companion_D0_.Create_DC0_(func() _dafny.Map {
		var _coll14 = _dafny.NewMapBuilder()
		_ = _coll14
		for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(181), _dafny.IntOfInt64(229))); ; {
			_compr_14, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _137_v6 _dafny.Int
			_137_v6 = interface{}(_compr_14).(_dafny.Int)
			if ((_dafny.IntOfInt64(181)).Cmp(_137_v6) <= 0) && ((_137_v6).Cmp(_dafny.IntOfInt64(229)) < 0) {
				_coll14.Add(Companion_Default___.SafeDivisionInt(_137_v6, _131_v5), _129_v2)
			}
		}
		return _coll14.ToMap()
	}())).Dtor_cf0(), 14)
	(_nw10).ArraySet1(_130_v3, 15)
	(_nw10).ArraySet1(Companion_Default___.Fm0(_128_v1, Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_130_v3), _128_v1, _131_v5, _128_v1, _127_globalState), func(_pat_let2_0 D0) D0 {
		return func(_138_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let3_0 _dafny.Int) D0 {
				return func(_139_dt__update_hcf9_h0 _dafny.Int) D0 {
					return Companion_D0_.Create_DC2_((_138_dt__update__tmp_h0).Dtor_cf5(), (_138_dt__update__tmp_h0).Dtor_cf6(), (_138_dt__update__tmp_h0).Dtor_cf7(), (_138_dt__update__tmp_h0).Dtor_cf8(), _139_dt__update_hcf9_h0)
				}(_pat_let3_0)
			}(_pat_let_tv1)
		}(_pat_let2_0)
	}(_134_v9), _127_globalState), 16)
	_135_v10 = _nw10
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_135_v10), 0))
	_ = _index13
	(_135_v10).ArraySet1(_130_v3, (_index13).Int())
	var _140_v11 _dafny.Map
	_ = _140_v11
	_140_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _130_v3)
	var _141_v12 D0
	_ = _141_v12
	_141_v12 = Companion_D0_.Create_DC0_(_130_v3)
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(207), _dafny.ArrayLen((_135_v10), 0))
	_ = _index14
	(_135_v10).ArraySet1((func() _dafny.Map {
		if (_140_v11).Contains(_131_v5) {
			return (_140_v11).Get(_131_v5).(_dafny.Map)
		}
		return Companion_Default___.Fm0(Companion_Default___.Fm1(_141_v12, _128_v1, (_dafny.Zero).Minus((_129_v2).Cardinality()), _128_v1, _127_globalState), _128_v1, _134_v9, _127_globalState)
	})(), (_index14).Int())
	if _128_v1 {
		var _142_v13 _dafny.Map
		_ = _142_v13
		_142_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_131_v5, _131_v5), _131_v5)
		_142_v13 = _142_v13
		var _143_v14 _dafny.Array
		_ = _143_v14
		var _nw11 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw11
		_143_v14 = _nw11
		var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))
		_ = _index15
		(_143_v14).ArraySet1(_128_v1, (_index15).Int())
		var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))
		_ = _index16
		(_143_v14).ArraySet1(_128_v1, (_index16).Int())
		var _144_v15 _dafny.Map
		_ = _144_v15
		_144_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), _131_v5)
		var _145_v16 _dafny.Map
		_ = _145_v16
		_145_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_126_v0, (_144_v15).Cardinality())
		var _146_v17 _dafny.Map
		_ = _146_v17
		_146_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _dafny.UnicodeSeqOfUtf8Bytes("e"))
		var _147_v19 _dafny.Sequence
		_ = _147_v19
		_147_v19 = _dafny.SeqOf((func() _dafny.Set {
			var _coll15 = _dafny.NewBuilder()
			_ = _coll15
			for _iter15 := _dafny.Iterate(((_146_v17).Update(_131_v5, _dafny.UnicodeSeqOfUtf8Bytes("k"))).Keys().Elements()); ; {
				_compr_15, _ok15 := _iter15()
				if !_ok15 {
					break
				}
				var _148_v18 _dafny.Int
				_148_v18 = interface{}(_compr_15).(_dafny.Int)
				if ((_146_v17).Update(_131_v5, _dafny.UnicodeSeqOfUtf8Bytes("k"))).Contains(_148_v18) {
					_coll15.Add(Companion_Default___.SafeDivisionInt(_148_v18, (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(511), _dafny.IntOfInt64(38))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("edejin")).Cardinality()))).Cardinality()))
				}
			}
			return _coll15.ToSet()
		}()).Cardinality())
		_145_v16 = (_145_v16).Update(_dafny.Companion_Sequence_.Concatenate(_126_v0, _126_v0), _dafny.IntOfUint32((_147_v19).Cardinality()))
		_126_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-612))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg10 _dafny.Int) interface{} {
				return coer10(arg10)
			}
		}(func(_149_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})), Companion_Default___.Fm2(_126_v0, _127_globalState))
		var _150_v20 D0
		_ = _150_v20
		_150_v20 = Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_126_v0).Cardinality()), (_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), _131_v5, (_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool))
		var _151_v21 _dafny.Sequence
		_ = _151_v21
		_151_v21 = _dafny.SeqOf(_128_v1)
		var _152_v22 _dafny.MultiSet
		_ = _152_v22
		_152_v22 = _dafny.MultiSetOf((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), _128_v1, false)
		var _nwElement0_2 bool = !((_150_v20).Dtor_cf2())
		_ = _nwElement0_2
		var _nw12 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(15))
		_ = _nw12
		(_nw12).ArraySet1(_nwElement0_2, 0)
		(_nw12).ArraySet1((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), 1)
		(_nw12).ArraySet1((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), 2)
		(_nw12).ArraySet1(Companion_Default___.Fm1(_141_v12, _128_v1, _131_v5, (_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), _127_globalState), 3)
		(_nw12).ArraySet1(_128_v1, 4)
		(_nw12).ArraySet1((_131_v5).Cmp(_131_v5) < 0, 5)
		(_nw12).ArraySet1(!((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool)) || ((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool)), 6)
		(_nw12).ArraySet1(_128_v1, 7)
		(_nw12).ArraySet1((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), 8)
		(_nw12).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_151_v21, _151_v21), 9)
		(_nw12).ArraySet1((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), 10)
		(_nw12).ArraySet1(_128_v1, 11)
		(_nw12).ArraySet1(_128_v1, 12)
		(_nw12).ArraySet1((_143_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(28), _dafny.ArrayLen((_143_v14), 0))).Int()).(bool), 13)
		(_nw12).ArraySet1((_152_v22).IsSubsetOf(_152_v22), 14)
		_143_v14 = _nw12
	} else {
		(_127_globalState).F5 = Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_126_v0).Cardinality()), _129_v2)), _128_v1, (func() _dafny.Int {
			if _128_v1 {
				return _131_v5
			}
			return _131_v5
		})(), _128_v1, _127_globalState)
		if _128_v1 {
			var _153_v23 *C0
			_ = _153_v23
			var _nw13 *C0 = New_C0_()
			_ = _nw13
			_nw13.Ctor__((func() bool {
				if true {
					return _128_v1
				}
				return _128_v1
			})())
			_153_v23 = _nw13
			var _154_v24 _dafny.Map
			_ = _154_v24
			_154_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _131_v5)
			var _rhs6 _dafny.Int = (_131_v5).Minus((func() _dafny.Int {
				if (_154_v24).Contains(_131_v5) {
					return (_154_v24).Get(_131_v5).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_126_v0).Cardinality())
			})())
			_ = _rhs6
			var _rhs7 bool = (_153_v23).F13()
			_ = _rhs7
			var _rhs8 _dafny.Int = _131_v5
			_ = _rhs8
			var _lhs5 *GlobalState = _127_globalState
			_ = _lhs5
			var _lhs6 *GlobalState = _127_globalState
			_ = _lhs6
			var _lhs7 *GlobalState = _127_globalState
			_ = _lhs7
			_lhs5.F2 = _rhs6
			_lhs6.F1 = _rhs7
			_lhs7.F2 = _rhs8
			var _pat_let_tv2 = _131_v5
			_ = _pat_let_tv2
			var _pat_let_tv3 = _131_v5
			_ = _pat_let_tv3
			var _pat_let_tv4 = _129_v2
			_ = _pat_let_tv4
			var _pat_let_tv5 = _131_v5
			_ = _pat_let_tv5
			var _pat_let_tv6 = _131_v5
			_ = _pat_let_tv6
			var _pat_let_tv7 = _129_v2
			_ = _pat_let_tv7
			var _155_v27 D13
			_ = _155_v27
			_155_v27 = Companion_D13_.Create_DC32_(_131_v5, true, Companion_Default___.Fm1(func(_pat_let4_0 D0) D0 {
				return func(_156_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let5_0 _dafny.Map) D0 {
						return func(_160_dt__update_hcf0_h0 _dafny.Map) D0 {
							return Companion_D0_.Create_DC0_(_160_dt__update_hcf0_h0)
						}(_pat_let5_0)
					}((func() _dafny.Map {
						var _coll16 = _dafny.NewMapBuilder()
						_ = _coll16
						for _iter16 := _dafny.Iterate((func() _dafny.Set {
							var _coll17 = _dafny.NewBuilder()
							_ = _coll17
							for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(598), _dafny.IntOfInt64(884))); ; {
								_compr_17, _ok17 := _iter17()
								if !_ok17 {
									break
								}
								var _157_v26 _dafny.Int
								_157_v26 = interface{}(_compr_17).(_dafny.Int)
								if ((_dafny.IntOfInt64(598)).Cmp(_157_v26) <= 0) && ((_157_v26).Cmp(_dafny.IntOfInt64(884)) < 0) {
									_coll17.Add((_157_v26).Plus(_pat_let_tv2))
								}
							}
							return _coll17.ToSet()
						}()).Elements()); ; {
							_compr_16, _ok16 := _iter16()
							if !_ok16 {
								break
							}
							var _158_v25 _dafny.Int
							_158_v25 = interface{}(_compr_16).(_dafny.Int)
							if (func() _dafny.Set {
								var _coll18 = _dafny.NewBuilder()
								_ = _coll18
								for _iter18 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(598), _dafny.IntOfInt64(884))); ; {
									_compr_18, _ok18 := _iter18()
									if !_ok18 {
										break
									}
									var _159_v26 _dafny.Int
									_159_v26 = interface{}(_compr_18).(_dafny.Int)
									if ((_dafny.IntOfInt64(598)).Cmp(_159_v26) <= 0) && ((_159_v26).Cmp(_dafny.IntOfInt64(884)) < 0) {
										_coll18.Add((_159_v26).Plus(_pat_let_tv3))
									}
								}
								return _coll18.ToSet()
							}()).Contains(_158_v25) {
								_coll16.Add(Companion_Default___.SafeModuloInt(_158_v25, _pat_let_tv5), _pat_let_tv4)
							}
						}
						return _coll16.ToMap()
					}()).Update(_pat_let_tv6, _pat_let_tv7))
				}(_pat_let4_0)
			}(_141_v12), _128_v1, _131_v5, _128_v1, _127_globalState), (_153_v23).F13(), _128_v1)
			var _pat_let_tv8 = _128_v1
			_ = _pat_let_tv8
			var _pat_let_tv9 = _153_v23
			_ = _pat_let_tv9
			var _pat_let_tv10 = _153_v23
			_ = _pat_let_tv10
			var _pat_let_tv11 = _128_v1
			_ = _pat_let_tv11
			var _161_v28 _dafny.Map
			_ = _161_v28
			_161_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func(_pat_let6_0 D13) D13 {
				return func(_166_dt__update__tmp_h3 D13) D13 {
					return func(_pat_let11_0 bool) D13 {
						return func(_167_dt__update_hcf58_h1 bool) D13 {
							return Companion_D13_.Create_DC32_((_166_dt__update__tmp_h3).Dtor_cf56(), (_166_dt__update__tmp_h3).Dtor_cf57(), _167_dt__update_hcf58_h1, (_166_dt__update__tmp_h3).Dtor_cf59(), (_166_dt__update__tmp_h3).Dtor_cf60())
						}(_pat_let11_0)
					}(_pat_let_tv11)
				}(_pat_let6_0)
			}(func(_pat_let7_0 D13) D13 {
				return func(_162_dt__update__tmp_h2 D13) D13 {
					return func(_pat_let8_0 bool) D13 {
						return func(_163_dt__update_hcf59_h0 bool) D13 {
							return func(_pat_let9_0 bool) D13 {
								return func(_164_dt__update_hcf58_h0 bool) D13 {
									return func(_pat_let10_0 bool) D13 {
										return func(_165_dt__update_hcf57_h0 bool) D13 {
											return Companion_D13_.Create_DC32_((_162_dt__update__tmp_h2).Dtor_cf56(), _165_dt__update_hcf57_h0, _164_dt__update_hcf58_h0, _163_dt__update_hcf59_h0, (_162_dt__update__tmp_h2).Dtor_cf60())
										}(_pat_let10_0)
									}((_pat_let_tv10).F13())
								}(_pat_let9_0)
							}((_pat_let_tv9).F13())
						}(_pat_let8_0)
					}(_pat_let_tv8)
				}(_pat_let7_0)
			}(_155_v27))).Dtor_cf58(), (_dafny.SetOf((_153_v23).F13())).Cardinality())
			_161_v28 = (_161_v28).Merge((_161_v28).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_128_v1), _131_v5)))
			_126_v0 = _dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_dafny.UnicodeSeqOfUtf8Bytes("ioebdfr"), _127_globalState), _126_v0)
			_131_v5 = (_134_v9).Dtor_cf5()
		} else {
			Companion_Default___.M7(_131_v5, _127_globalState)
			var _168_v29 _dafny.Array
			_ = _168_v29
			var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(17))
			_ = _nw14
			_168_v29 = _nw14
			_168_v29 = (func() _dafny.Array {
				if (func() bool {
					if !(_128_v1) {
						return !(_128_v1)
					}
					return _128_v1
				})() {
					return _168_v29
				}
				return _168_v29
			})()
			var _169_v30 _dafny.Sequence
			_ = _169_v30
			_169_v30 = _dafny.SeqOf(_128_v1, _128_v1)
			var _170_v31 _dafny.Sequence
			_ = _170_v31
			_170_v31 = _dafny.SeqOf(_131_v5, _131_v5)
			var _171_v32 *C4
			_ = _171_v32
			var _nw15 *C4 = New_C4_()
			_ = _nw15
			_nw15.Ctor__(_dafny.IntOfUint32((_170_v31).Cardinality()))
			_171_v32 = _nw15
			var _172_v33 _dafny.Sequence
			_ = _172_v33
			_172_v33 = _dafny.SeqOf(_171_v32, _171_v32)
			var _173_v34 _dafny.Array
			_ = _173_v34
			var _nwElement0_3 _dafny.Int = _131_v5
			_ = _nwElement0_3
			var _nw16 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(24))
			_ = _nw16
			(_nw16).ArraySet1(_nwElement0_3, 0)
			(_nw16).ArraySet1(_131_v5, 1)
			(_nw16).ArraySet1(_131_v5, 2)
			(_nw16).ArraySet1(_131_v5, 3)
			(_nw16).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_169_v30).Cardinality()), _131_v5), 4)
			(_nw16).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(806), _131_v5), 5)
			(_nw16).ArraySet1(_131_v5, 6)
			(_nw16).ArraySet1((_131_v5).Times(_dafny.IntOfInt64(442)), 7)
			(_nw16).ArraySet1((_dafny.IntOfInt64(694)).Times(_dafny.IntOfUint32((_172_v33).Cardinality())), 8)
			(_nw16).ArraySet1(_131_v5, 9)
			(_nw16).ArraySet1(_131_v5, 10)
			(_nw16).ArraySet1(_dafny.IntOfInt64(660), 11)
			(_nw16).ArraySet1(_131_v5, 12)
			(_nw16).ArraySet1((_131_v5).Times(_dafny.IntOfUint32((_169_v30).Cardinality())), 13)
			(_nw16).ArraySet1(_131_v5, 14)
			(_nw16).ArraySet1(_131_v5, 15)
			(_nw16).ArraySet1(_131_v5, 16)
			(_nw16).ArraySet1(Companion_Default___.Fm24(_131_v5, _127_globalState), 17)
			(_nw16).ArraySet1((_dafny.Zero).Minus(_131_v5), 18)
			(_nw16).ArraySet1((_131_v5).Times(_131_v5), 19)
			(_nw16).ArraySet1(_131_v5, 20)
			(_nw16).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_128_v1, _128_v1)).Cardinality()), 21)
			(_nw16).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfInt64(-329)), 22)
			(_nw16).ArraySet1(_131_v5, 23)
			_173_v34 = _nw16
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_173_v34), 0))
			_ = _index17
			(_173_v34).ArraySet1(Companion_Default___.SafeDivisionInt(_131_v5, _131_v5), (_index17).Int())
			var _174_v35 _dafny.MultiSet
			_ = _174_v35
			_174_v35 = _dafny.MultiSetOf(_173_v34, _173_v34)
			var _175_v36 _dafny.Map
			_ = _175_v36
			_175_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v1, _131_v5)
			var _176_v37 _dafny.MultiSet
			_ = _176_v37
			_176_v37 = _dafny.MultiSetOf((_131_v5).Cmp((func() _dafny.Int {
				if (_174_v35).Contains(_173_v34) {
					return (_174_v35).Multiplicity(_173_v34)
				}
				return (_dafny.Zero).Minus((_175_v36).Cardinality())
			})()) != 0, Companion_Default___.Fm1(_141_v12, _128_v1, (_dafny.Zero).Minus(_131_v5), _128_v1, _127_globalState), false)
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_173_v34), 0))
			_ = _index18
			(_173_v34).ArraySet1((func() _dafny.Int {
				if (_176_v37).Contains((_128_v1) || (_128_v1)) {
					return (_176_v37).Multiplicity((_128_v1) || (_128_v1))
				}
				return _dafny.IntOfInt64(83)
			})(), (_index18).Int())
			(_127_globalState).F2 = (func() _dafny.Int {
				if (_129_v2).IsSubsetOf(_129_v2) {
					return _131_v5
				}
				return _131_v5
			})()
			(_127_globalState).F0 = (func() _dafny.Int {
				if !(((_129_v2).Cardinality()).Cmp((_173_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_173_v34), 0))).Int()).(_dafny.Int)) == 0) {
					return (_173_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(597), _dafny.ArrayLen((_173_v34), 0))).Int()).(_dafny.Int)
				}
				return Companion_Default___.Fm15(_131_v5, _128_v1, _127_globalState)
			})()
		}
		var _177_v38 _dafny.MultiSet
		_ = _177_v38
		_177_v38 = _dafny.MultiSetOf(_128_v1, _128_v1)
		(_127_globalState).F2 = (_131_v5).Plus((_177_v38).Cardinality())
		(_127_globalState).F0 = ((_131_v5).Plus((_dafny.Zero).Minus(_131_v5))).Times(_131_v5)
		if _128_v1 {
			Companion_Default___.M7(_131_v5, _127_globalState)
			var _178_v39 _dafny.Array
			_ = _178_v39
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_1
			var _nw17 _dafny.Array
			_ = _nw17
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw17 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) bool = (func(_179_v1 bool) func(_dafny.Int) bool {
					return func(_180_i1 _dafny.Int) bool {
						return _179_v1
					}
				})(_128_v1)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw17 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw17).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw17).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_178_v39 = _nw17
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_178_v39), 0))
			_ = _index19
			(_178_v39).ArraySet1(!(_128_v1), (_index19).Int())
			var _181_v40 _dafny.Sequence
			_ = _181_v40
			_181_v40 = _dafny.SeqOf(_131_v5, _131_v5, _131_v5)
			var _182_v41 _dafny.Sequence
			_ = _182_v41
			_182_v41 = _dafny.SeqOf(!(true), _128_v1, Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_130_v3), _128_v1, (_181_v40).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.IntOfUint32((_181_v40).Cardinality()))).Uint32()).(_dafny.Int), _128_v1, _127_globalState), _128_v1, _128_v1)
			var _183_v42 _dafny.Map
			_ = _183_v42
			_183_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_126_v0).Cardinality()), _128_v1)
			var _184_v43 D14
			_ = _184_v43
			_184_v43 = Companion_D14_.Create_DC36_(_128_v1, (_dafny.Zero).Minus(_131_v5), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ncd")).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-972), (func() bool {
				if (_183_v42).Contains(_131_v5) {
					return (_183_v42).Get(_131_v5).(bool)
				}
				return _128_v1
			})()))
			var _185_v44 _dafny.Map
			_ = _185_v44
			_185_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v7, _184_v43)
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_178_v39), 0))
			_ = _index20
			var _rhs9 _dafny.Int = _131_v5
			_ = _rhs9
			var _rhs10 bool = _128_v1
			_ = _rhs10
			var _rhs11 bool = (_182_v41).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_185_v44).Cardinality()), _dafny.IntOfUint32((_182_v41).Cardinality()))).Uint32()).(bool)
			_ = _rhs11
			var _lhs8 *GlobalState = _127_globalState
			_ = _lhs8
			var _lhs9 _dafny.Array = _178_v39
			_ = _lhs9
			var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_178_v39), 0))
			_ = _lhs10
			var _lhs11 *GlobalState = _127_globalState
			_ = _lhs11
			_lhs8.F2 = _rhs9
			(_lhs9).ArraySet1(_rhs10, (_lhs10).Int())
			_lhs11.F1 = _rhs11
			(_127_globalState).F5 = !((_131_v5).Cmp(_131_v5) < 0)
			_181_v40 = _181_v40
			var _186_v45 _dafny.Map
			_ = _186_v45
			_186_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v1, _dafny.SeqOf(false, (_178_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_178_v39), 0))).Int()).(bool), (_178_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_178_v39), 0))).Int()).(bool)))
			(_127_globalState).F1 = !((_186_v45).Equals((Companion_Default___.Fm38((_178_v39).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(957), _dafny.ArrayLen((_178_v39), 0))).Int()).(bool), _127_globalState)).Merge(_186_v45)))
		} else {
			var _187_v46 _dafny.Map
			_ = _187_v46
			_187_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, (func() bool {
				if true {
					return _128_v1
				}
				return _128_v1
			})())
			var _188_v47 _dafny.Map
			_ = _188_v47
			_188_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v7, _131_v5)
			_187_v46 = (_187_v46).Update(_131_v5, (_133_v8).IsSubsetOf(_dafny.MultiSetFromSeq(Companion_Default___.Fm25(_131_v5, _188_v47, _132_v7, _127_globalState))))
			_126_v0 = _dafny.UnicodeSeqOfUtf8Bytes("dcti")
			var _189_v48 T0
			_ = _189_v48
			var _nw18 *C6 = New_C6_()
			_ = _nw18
			_nw18.Ctor__(_131_v5)
			_189_v48 = _nw18
			var _190_v49 T1
			_ = _190_v49
			var _nw19 *C1 = New_C1_()
			_ = _nw19
			_nw19.Ctor__()
			_190_v49 = _nw19
			var _191_v50 _dafny.MultiSet
			_ = _191_v50
			_191_v50 = _dafny.MultiSetOf(_190_v49, _190_v49, _190_v49, _190_v49)
			var _192_v51 D11
			_ = _192_v51
			_192_v51 = Companion_D11_.Create_DC26_(_189_v48, _191_v50, _131_v5, _126_v0)
			var _193_v52 _dafny.Set
			_ = _193_v52
			_193_v52 = _dafny.SetOf(_189_v48)
			var _194_v53 _dafny.Array
			_ = _194_v53
			var _nwElement0_4 bool = _128_v1
			_ = _nwElement0_4
			var _nw20 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(5))
			_ = _nw20
			(_nw20).ArraySet1(_nwElement0_4, 0)
			(_nw20).ArraySet1((_131_v5).Cmp(_131_v5) != 0, 1)
			(_nw20).ArraySet1((_dafny.SetOf((_192_v51).Dtor_cf45(), _189_v48)).IsDisjointFrom(_193_v52), 2)
			(_nw20).ArraySet1(_128_v1, 3)
			(_nw20).ArraySet1(_128_v1, 4)
			_194_v53 = _nw20
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_194_v53), 0))
			_ = _index21
			(_194_v53).ArraySet1(_128_v1, (_index21).Int())
			var _195_v54 *C0
			_ = _195_v54
			var _nw21 *C0 = New_C0_()
			_ = _nw21
			_nw21.Ctor__(_128_v1)
			_195_v54 = _nw21
			var _196_v55 _dafny.Map
			_ = _196_v55
			_196_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(516), _195_v54)
			var _197_v56 _dafny.Set
			_ = _197_v56
			_197_v56 = _dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf((_196_v55).Cardinality(), _131_v5)).Cardinality()))
			var _198_v57 _dafny.Sequence
			_ = _198_v57
			_198_v57 = _dafny.SeqOf((_189_v48).F6(), _131_v5)
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_194_v53), 0))
			_ = _index22
			(_194_v53).ArraySet1((Companion_Default___.Fm10(_dafny.IntOfUint32((_198_v57).Cardinality()), (_195_v54).F13(), _128_v1, _127_globalState)).IsSubsetOf(_197_v56), (_index22).Int())
			_128_v1 = _128_v1
			var _199_v58 _dafny.Map
			_ = _199_v58
			_199_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_195_v54).F13(), _dafny.IntOfUint32((_198_v57).Cardinality()))
			(_127_globalState).F2 = (func() _dafny.Int {
				if _128_v1 {
					return (_dafny.Zero).Minus(((_199_v58).Cardinality()).Times(_dafny.IntOfInt64(826)))
				}
				return Companion_Default___.SafeModuloInt(_131_v5, _dafny.IntOfInt64(964))
			})()
		}
	}
	var _200_v59 *C0
	_ = _200_v59
	var _nw22 *C0 = New_C0_()
	_ = _nw22
	_nw22.Ctor__(_128_v1)
	_200_v59 = _nw22
	var _source8 *C0 = _200_v59
	_ = _source8
	var _201___mcc_h0 *C0 = _source8
	_ = _201___mcc_h0
	var _202_cf72 *C0 = _201___mcc_h0
	_ = _202_cf72
	var _203_v60 _dafny.Array
	_ = _203_v60
	var _nw23 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(6))
	_ = _nw23
	_203_v60 = _nw23
	var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_203_v60), 0))
	_ = _index23
	(_203_v60).ArraySet1(_128_v1, (_index23).Int())
	var _204_v61 _dafny.Set
	_ = _204_v61
	_204_v61 = _dafny.SetOf(_131_v5)
	var _205_v62 _dafny.MultiSet
	_ = _205_v62
	_205_v62 = _dafny.MultiSetOf(_204_v61)
	var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((_203_v60), 0))
	_ = _index24
	(_203_v60).ArraySet1((_205_v62).Contains(_204_v61), (_index24).Int())
	var _206_v63 _dafny.Sequence
	_ = _206_v63
	_206_v63 = _dafny.SeqOf(_131_v5)
	_133_v8 = _dafny.MultiSetOf((_dafny.Zero).Minus((func() _dafny.Int {
		if _128_v1 {
			return _dafny.IntOfUint32((_206_v63).Cardinality())
		}
		return Companion_Default___.Fm15(_131_v5, _128_v1, _127_globalState)
	})()), (_131_v5).Plus((_204_v61).Cardinality()), _131_v5, _131_v5)
	(_127_globalState).F1 = !((_200_v59).F13())
	if (func() bool {
		if (_200_v59).F13() {
			return (_200_v59).F13()
		}
		return !((_202_cf72).F13()) || (_128_v1)
	})() {
		var _207_v64 _dafny.Sequence
		_ = _207_v64
		_207_v64 = _dafny.SeqOf((_202_cf72).F13())
		_207_v64 = _dafny.Companion_Sequence_.Concatenate(_207_v64, _dafny.SeqOf((_200_v59).F13()))
		var _208_v65 *C2
		_ = _208_v65
		var _nw24 *C2 = New_C2_()
		_ = _nw24
		_nw24.Ctor__(_126_v0, (_202_cf72).F13())
		_208_v65 = _nw24
		var _209_v66 _dafny.Map
		_ = _209_v66
		_209_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_204_v61, _207_v64)
		(_208_v65).F12 = !(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_204_v61).Cardinality()), Companion_Default___.Fm21((_208_v65).F11(), _127_globalState))).Equals(_209_v66)
		var _210_v67 _dafny.Array
		_ = _210_v67
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(16)
		_ = _len0_2
		var _nw25 _dafny.Array
		_ = _nw25
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw25 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) _dafny.Int = (func(_211_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_212_i2 _dafny.Int) _dafny.Int {
					return (_212_i2).Times(_211_v5)
				}
			})(_131_v5)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw25 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw25).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw25).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_210_v67 = _nw25
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _len0_3
		var _nw26 _dafny.Array
		_ = _nw26
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw26 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Int = (func(_213_v1 bool) func(_dafny.Int) _dafny.Int {
				return func(_214_i3 _dafny.Int) _dafny.Int {
					return (_214_i3).Times((_dafny.Zero).Minus((_dafny.MultiSetOf(_213_v1)).Cardinality()))
				}
			})(_128_v1)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw26 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw26).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw26).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_210_v67 = _nw26
		var _215_v68 _dafny.Map
		_ = _215_v68
		_215_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_132_v7, Companion_Default___.Fm1(_141_v12, _208_v65.F12, (_dafny.Zero).Minus(_131_v5), (_202_cf72).F13(), _127_globalState))
		var _216_v69 _dafny.Sequence
		_ = _216_v69
		_216_v69 = _dafny.SeqOf(_215_v68, _215_v68, (_215_v68).Merge(_215_v68))
		_216_v69 = _216_v69
	} else {
		Companion_Default___.M7(_131_v5, _127_globalState)
		(_127_globalState).F5 = (_200_v59).F13()
		_126_v0 = _126_v0
		(_127_globalState).F5 = true
		var _217_v70 _dafny.Sequence
		_ = _217_v70
		_217_v70 = _dafny.SeqOf((_200_v59).F13(), (_200_v59).F13())
		var _218_v71 _dafny.Map
		_ = _218_v71
		_218_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_dafny.IntOfUint32((_217_v70).Cardinality())).Cmp(_131_v5) == 0)), (func() bool {
			if (_202_cf72).F13() {
				return (_200_v59).F13()
			}
			return (_202_cf72).F13()
		})())
		_218_v71 = (_218_v71).Update((func() bool {
			if (_218_v71).Contains(true) {
				return (_218_v71).Get(true).(bool)
			}
			return _128_v1
		})(), (_133_v8).IsProperSubsetOf(_133_v8))
	}
	var _219_v72 _dafny.Map
	_ = _219_v72
	_219_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v1, _131_v5)
	var _220_v73 _dafny.Map
	_ = _220_v73
	_220_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, (_200_v59).F13())
	var _221_v74 _dafny.Set
	_ = _221_v74
	_221_v74 = _dafny.SetOf(Companion_Default___.SafeModuloInt(_131_v5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_219_v72, _219_v72), (Companion_Default___.SafeIndex(_131_v5, _dafny.IntOfUint32((_dafny.SeqOf(_219_v72, _219_v72)).Cardinality()))).Uint32(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_128_v1, (_220_v73).Cardinality()))).Cardinality())))
	_221_v74 = _221_v74
	_128_v1 = false
	(_127_globalState).F0 = _131_v5
	var _rhs12 _dafny.Int = _131_v5
	_ = _rhs12
	var _rhs13 bool = _128_v1
	_ = _rhs13
	var _lhs12 *GlobalState = _127_globalState
	_ = _lhs12
	var _lhs13 *GlobalState = _127_globalState
	_ = _lhs13
	_lhs12.F0 = _rhs12
	_lhs13.F5 = _rhs13
	var _222_v75 _dafny.Map
	_ = _222_v75
	_222_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm39(_131_v5, _127_globalState)).Cardinality(), _131_v5)
	var _223_v76 D16
	_ = _223_v76
	_223_v76 = Companion_D16_.Create_DC39_((_dafny.Zero).Minus(_dafny.IntOfUint32((_126_v0).Cardinality())), _222_v75)
	var _pat_let_tv12 = _222_v75
	_ = _pat_let_tv12
	var _source9 D16 = func(_pat_let12_0 D16) D16 {
		return func(_224_dt__update__tmp_h4 D16) D16 {
			return func(_pat_let13_0 _dafny.Map) D16 {
				return func(_225_dt__update_hcf75_h0 _dafny.Map) D16 {
					return Companion_D16_.Create_DC39_((_224_dt__update__tmp_h4).Dtor_cf74(), _225_dt__update_hcf75_h0)
				}(_pat_let13_0)
			}(_pat_let_tv12)
		}(_pat_let12_0)
	}(_223_v76)
	_ = _source9
	if _source9.Is_DC39() {
		var _226___mcc_h1 _dafny.Int = _source9.Get_().(D16_DC39).Cf74
		_ = _226___mcc_h1
		var _227___mcc_h2 _dafny.Map = _source9.Get_().(D16_DC39).Cf75
		_ = _227___mcc_h2
		var _228_cf75 _dafny.Map = _227___mcc_h2
		_ = _228_cf75
		var _229_cf74 _dafny.Int = _226___mcc_h1
		_ = _229_cf74
		var _230_v77 D1
		_ = _230_v77
		_230_v77 = Companion_D1_.Create_DC5_()
		var _source10 D1 = _230_v77
		_ = _source10
		if _source10.Is_DC5() {
			var _231_v78 _dafny.MultiSet
			_ = _231_v78
			_231_v78 = _dafny.MultiSetOf((_200_v59).F13())
			var _232_v79 _dafny.Sequence
			_ = _232_v79
			_232_v79 = _dafny.SeqOf(_229_cf74)
			var _233_v80 *C3
			_ = _233_v80
			var _nw27 *C3 = New_C3_()
			_ = _nw27
			_nw27.Ctor__((_231_v78).Cardinality(), _232_v79, _229_cf74)
			_233_v80 = _nw27
			var _234_v81 D13
			_ = _234_v81
			_234_v81 = Companion_D13_.Create_DC33_((_200_v59).F13(), (_129_v2).Cardinality(), _dafny.IntOfInt64(11), (_200_v59).F13(), _233_v80)
			var _235_v82 _dafny.Array
			_ = _235_v82
			var _nwElement0_5 _dafny.Int = _229_cf74
			_ = _nwElement0_5
			var _nw28 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(6))
			_ = _nw28
			(_nw28).ArraySet1(_nwElement0_5, 0)
			(_nw28).ArraySet1((_222_v75).Cardinality(), 1)
			(_nw28).ArraySet1(Companion_Default___.SafeDivisionInt(_131_v5, _229_cf74), 2)
			(_nw28).ArraySet1(_131_v5, 3)
			(_nw28).ArraySet1((func() _dafny.Int {
				if !(Companion_Default___.Fm1(_141_v12, (_234_v81).Dtor_cf64(), _131_v5, !((_200_v59).F13()), _127_globalState)) {
					return (_221_v74).Cardinality()
				}
				return _131_v5
			})(), 4)
			(_nw28).ArraySet1(_131_v5, 5)
			_235_v82 = _nw28
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_235_v82), 0))
			_ = _index25
			(_235_v82).ArraySet1((_dafny.Zero).Minus((_dafny.Zero).Minus((_233_v80).F9())), (_index25).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(649), _dafny.ArrayLen((_235_v82), 0))
			_ = _index26
			(_235_v82).ArraySet1((_233_v80).F9(), (_index26).Int())
			var _236_v83 *C5
			_ = _236_v83
			var _nw29 *C5 = New_C5_()
			_ = _nw29
			_nw29.Ctor__((_200_v59).F13(), _132_v7)
			_236_v83 = _nw29
			_236_v83 = _236_v83
			var _237_v84 _dafny.Array
			_ = _237_v84
			var _nw30 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(12))
			_ = _nw30
			_237_v84 = _nw30
			var _238_v85 _dafny.Array
			_ = _238_v85
			var _len0_4 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_4
			var _nw31 _dafny.Array
			_ = _nw31
			if _len0_4.Cmp(_dafny.Zero) == 0 {
				_nw31 = _dafny.NewArray(_len0_4)
			} else {
				var _init4 func(_dafny.Int) bool = (func(_239_v83 *C5) func(_dafny.Int) bool {
					return func(_240_i4 _dafny.Int) bool {
						return (_239_v83).F7()
					}
				})(_236_v83)
				_ = _init4
				var _element0_4 = _init4(_dafny.Zero)
				_ = _element0_4
				_nw31 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
				(_nw31).ArraySet1(_element0_4, 0)
				var _nativeLen0_4 = (_len0_4).Int()
				_ = _nativeLen0_4
				for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
					(_nw31).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
				}
			}
			_238_v85 = _nw31
			var _241_v86 _dafny.Sequence
			_ = _241_v86
			_241_v86 = _dafny.SeqOf(_238_v85, _238_v85)
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_237_v84), 0))
			_ = _index27
			(_237_v84).ArraySet1(_241_v86, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_237_v84), 0))
			_ = _index28
			(_237_v84).ArraySet1(_241_v86, (_index28).Int())
			(_127_globalState).F2 = (_231_v78).Cardinality()
		} else {
			var _242___mcc_h4 _dafny.Sequence = _source10.Get_().(D1_DC4).Cf11
			_ = _242___mcc_h4
			var _243_cf11 _dafny.Sequence = _242___mcc_h4
			_ = _243_cf11
			var _244_v87 _dafny.Array
			_ = _244_v87
			var _nw32 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw32
			_244_v87 = _nw32
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_244_v87), 0))
			_ = _index29
			(_244_v87).ArraySet1((_200_v59).F13(), (_index29).Int())
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(671), _dafny.ArrayLen((_244_v87), 0))
			_ = _index30
			(_244_v87).ArraySet1(_128_v1, (_index30).Int())
			(_127_globalState).F1 = (_200_v59).F13()
			var _245_v88 _dafny.Array
			_ = _245_v88
			var _len0_5 _dafny.Int = _dafny.IntOfInt64(2)
			_ = _len0_5
			var _nw33 _dafny.Array
			_ = _nw33
			if _len0_5.Cmp(_dafny.Zero) == 0 {
				_nw33 = _dafny.NewArray(_len0_5)
			} else {
				var _init5 func(_dafny.Int) _dafny.Int = (func(_246_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_247_i5 _dafny.Int) _dafny.Int {
						return (_247_i5).Plus(_246_v5)
					}
				})(_131_v5)
				_ = _init5
				var _element0_5 = _init5(_dafny.Zero)
				_ = _element0_5
				_nw33 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
				(_nw33).ArraySet1(_element0_5, 0)
				var _nativeLen0_5 = (_len0_5).Int()
				_ = _nativeLen0_5
				for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
					(_nw33).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
				}
			}
			_245_v88 = _nw33
			_245_v88 = _245_v88
			Companion_Default___.M7(_229_cf74, _127_globalState)
		}
		if (_200_v59).F13() {
			var _248_v89 _dafny.Array
			_ = _248_v89
			var _nw34 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw34
			_248_v89 = _nw34
			_248_v89 = _248_v89
			var _249_v90 _dafny.Map
			_ = _249_v90
			_249_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_248_v89, _128_v1)
			(_127_globalState).F5 = (func() bool {
				if (_249_v90).Contains(_248_v89) {
					return false
				}
				return (_200_v59).F13()
			})()
			var _250_v91 D0
			_ = _250_v91
			_250_v91 = Companion_D0_.Create_DC1_(_131_v5, _128_v1, _229_cf74, (_200_v59).F13())
			var _251_v92 _dafny.Map
			_ = _251_v92
			_251_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_141_v12, !((_200_v59).F13()), _dafny.IntOfUint32((_126_v0).Cardinality()), Companion_Default___.Fm1(_141_v12, (_200_v59).F13(), _dafny.IntOfInt64(-926), (_250_v91).Dtor_cf4(), _127_globalState), _127_globalState), _126_v0)
			_251_v92 = (_251_v92).Update((_200_v59).F13(), _126_v0)
			var _252_v93 _dafny.Array
			_ = _252_v93
			var _nwElement0_6 _dafny.Array = _248_v89
			_ = _nwElement0_6
			var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(19))
			_ = _nw35
			(_nw35).ArraySet1(_nwElement0_6, 0)
			(_nw35).ArraySet1(_248_v89, 1)
			(_nw35).ArraySet1(_248_v89, 2)
			(_nw35).ArraySet1(_248_v89, 3)
			(_nw35).ArraySet1(_248_v89, 4)
			(_nw35).ArraySet1(_248_v89, 5)
			(_nw35).ArraySet1(_248_v89, 6)
			(_nw35).ArraySet1(_248_v89, 7)
			(_nw35).ArraySet1(_248_v89, 8)
			(_nw35).ArraySet1(_248_v89, 9)
			(_nw35).ArraySet1(_248_v89, 10)
			(_nw35).ArraySet1(_248_v89, 11)
			(_nw35).ArraySet1(_248_v89, 12)
			(_nw35).ArraySet1(_248_v89, 13)
			(_nw35).ArraySet1(_248_v89, 14)
			(_nw35).ArraySet1(_248_v89, 15)
			(_nw35).ArraySet1(_248_v89, 16)
			(_nw35).ArraySet1(_248_v89, 17)
			(_nw35).ArraySet1(_248_v89, 18)
			_252_v93 = _nw35
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_252_v93), 0))
			_ = _index31
			(_252_v93).ArraySet1(_248_v89, (_index31).Int())
			var _253_v94 _dafny.Array
			_ = _253_v94
			_253_v94 = _248_v89
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(944), _dafny.ArrayLen((_252_v93), 0))
			_ = _index32
			(_252_v93).ArraySet1((_253_v94), (_index32).Int())
			var _254_v95 *C0
			_ = _254_v95
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__((_229_cf74).Cmp(_229_cf74) == 0)
			_254_v95 = _nw36
		} else {
			var _255_v96 T0
			_ = _255_v96
			var _nw37 *C4 = New_C4_()
			_ = _nw37
			_nw37.Ctor__(_229_cf74)
			_255_v96 = _nw37
			var _256_v97 _dafny.Sequence
			_ = _256_v97
			_256_v97 = _dafny.SeqOf(_255_v96, _255_v96)
			var _257_v98 _dafny.Map
			_ = _257_v98
			_257_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v59).F13(), _dafny.MultiSetOf(_256_v97))
			var _258_v99 _dafny.Sequence
			_ = _258_v99
			_258_v99 = _dafny.SeqOf(_256_v97)
			_257_v98 = (_257_v98).Update((_200_v59).F13(), _dafny.MultiSetFromSeq(_258_v99))
			var _259_v100 _dafny.Sequence
			_ = _259_v100
			_259_v100 = _dafny.SeqOf(_229_cf74, (_255_v96).F6())
			_259_v100 = _259_v100
			(_127_globalState).F5 = ((_200_v59).F13()) == ((_200_v59).F13())
			var _260_v101 D8
			_ = _260_v101
			_260_v101 = Companion_D8_.Create_DC21_((_200_v59).F13(), Companion_Default___.Fm1(_141_v12, (_200_v59).F13(), (_255_v96).F6(), false, _127_globalState), _131_v5)
			(_127_globalState).F5 = !((_260_v101).Dtor_cf37())
			var _261_v102 _dafny.Sequence
			_ = _261_v102
			_261_v102 = _dafny.SeqOf((_200_v59).F13(), _128_v1, false)
			(_127_globalState).F0 = (_dafny.IntOfUint32((_261_v102).Cardinality())).Minus(_131_v5)
		}
		var _262_v103 _dafny.Array
		_ = _262_v103
		var _nw38 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(20))
		_ = _nw38
		_262_v103 = _nw38
		var _263_v104 D14
		_ = _263_v104
		_263_v104 = Companion_D14_.Create_DC35_(_262_v103)
		var _source11 D14 = _263_v104
		_ = _source11
		if _source11.Is_DC36() {
			var _264___mcc_h5 bool = _source11.Get_().(D14_DC36).Cf68
			_ = _264___mcc_h5
			var _265___mcc_h6 _dafny.Int = _source11.Get_().(D14_DC36).Cf69
			_ = _265___mcc_h6
			var _266___mcc_h7 _dafny.Int = _source11.Get_().(D14_DC36).Cf70
			_ = _266___mcc_h7
			var _267___mcc_h8 _dafny.Map = _source11.Get_().(D14_DC36).Cf71
			_ = _267___mcc_h8
			var _268_cf71 _dafny.Map = _267___mcc_h8
			_ = _268_cf71
			var _269_cf70 _dafny.Int = _266___mcc_h7
			_ = _269_cf70
			var _270_cf69 _dafny.Int = _265___mcc_h6
			_ = _270_cf69
			var _271_cf68 bool = _264___mcc_h5
			_ = _271_cf68
			(_127_globalState).F1 = Companion_Default___.Fm1(_141_v12, !(_271_cf68), _269_cf70, (_200_v59).F13(), _127_globalState)
			_128_v1 = ((_269_cf70).Times(_dafny.IntOfUint32((_126_v0).Cardinality()))).Cmp((_dafny.IntOfUint32((_126_v0).Cardinality())).Minus(_229_cf74)) != 0
			var _272_v105 _dafny.Sequence
			_ = _272_v105
			_272_v105 = _dafny.SeqOf(Companion_D1_.Create_DC5_(), _230_v77, (func() D1 {
				if _271_cf68 {
					return Companion_D1_.Create_DC5_()
				}
				return _230_v77
			})(), Companion_D1_.Create_DC5_(), _230_v77)
			_272_v105 = (func() _dafny.Sequence {
				if _128_v1 {
					return _272_v105
				}
				return _272_v105
			})()
			Companion_Default___.M7(_131_v5, _127_globalState)
		} else {
			var _273___mcc_h9 _dafny.Array = _source11.Get_().(D14_DC35).Cf67
			_ = _273___mcc_h9
			var _274_cf67 _dafny.Array = _273___mcc_h9
			_ = _274_cf67
			var _275_v106 _dafny.Array
			_ = _275_v106
			var _nwElement0_7 bool = (_200_v59).F13()
			_ = _nwElement0_7
			var _nw39 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(3))
			_ = _nw39
			(_nw39).ArraySet1(_nwElement0_7, 0)
			(_nw39).ArraySet1((_200_v59).F13(), 1)
			(_nw39).ArraySet1((_200_v59).F13(), 2)
			_275_v106 = _nw39
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_275_v106), 0))
			_ = _index33
			(_275_v106).ArraySet1((_200_v59).F13(), (_index33).Int())
			var _276_v107 _dafny.Set
			_ = _276_v107
			_276_v107 = _dafny.SetOf(_dafny.Companion_Sequence_.Update(_126_v0, (Companion_Default___.SafeIndex(_131_v5, _dafny.IntOfUint32((_126_v0).Cardinality()))).Uint32(), _132_v7))
			var _277_v108 _dafny.Map
			_ = _277_v108
			_277_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, Companion_Default___.Fm8(_132_v7, _128_v1, _dafny.IntOfInt64(-237), Companion_Default___.Fm24(_131_v5, _127_globalState), _127_globalState))
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_275_v106), 0))
			_ = _index34
			var _rhs14 bool = !((((_129_v2).Difference(_129_v2)).Cardinality()).Cmp(_dafny.IntOfInt64(285)) < 0)
			_ = _rhs14
			var _rhs15 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_277_v108).Contains(_131_v5) {
					return (_277_v108).Get(_131_v5).(_dafny.CodePoint)
				}
				return _132_v7
			})()
			_ = _rhs15
			var _rhs16 bool = (_200_v59).F13()
			_ = _rhs16
			var _rhs17 _dafny.Set = _dafny.SetOf(_126_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(569))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg11 _dafny.Int) interface{} {
					return coer11(arg11)
				}
			}((func(_278_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_279_i6 _dafny.Int) _dafny.CodePoint {
					return _278_v7
				}
			})(_132_v7))), _dafny.Companion_Sequence_.Concatenate(_126_v0, _126_v0))
			_ = _rhs17
			var _lhs14 *GlobalState = _127_globalState
			_ = _lhs14
			var _lhs15 _dafny.Array = _275_v106
			_ = _lhs15
			var _lhs16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_275_v106), 0))
			_ = _lhs16
			_lhs14.F1 = _rhs14
			_132_v7 = _rhs15
			(_lhs15).ArraySet1(_rhs16, (_lhs16).Int())
			_276_v107 = _rhs17
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_275_v106), 0))
			_ = _index35
			(_275_v106).ArraySet1(Companion_Default___.Fm1(_141_v12, (_131_v5).Cmp(_dafny.IntOfInt64(695)) >= 0, ((func() _dafny.Int {
				if (_133_v8).Contains(_229_cf74) {
					return (_133_v8).Multiplicity(_229_cf74)
				}
				return _131_v5
			})()).Times(_dafny.IntOfInt64(-56)), (_200_v59).F13(), _127_globalState), (_index35).Int())
			var _280_v109 _dafny.Array
			_ = _280_v109
			var _nw40 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw40
			_280_v109 = _nw40
			var _rhs18 _dafny.Array = _280_v109
			_ = _rhs18
			var _rhs19 _dafny.Int = Companion_Default___.Fm24(_131_v5, _127_globalState)
			_ = _rhs19
			var _rhs20 bool = (_131_v5).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_126_v0, _126_v0)).Cardinality())) > 0
			_ = _rhs20
			_280_v109 = _rhs18
			_131_v5 = _rhs19
			_128_v1 = _rhs20
			var _281_v110 D8
			_ = _281_v110
			_281_v110 = Companion_D8_.Create_DC21_((_275_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(482), _dafny.ArrayLen((_275_v106), 0))).Int()).(bool), (_200_v59).F13(), _229_cf74)
			(_127_globalState).F2 = (_281_v110).Dtor_cf38()
		}
		var _282_v111 _dafny.Map
		_ = _282_v111
		_282_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v77, (_200_v59).F13())
		var _rhs21 _dafny.Int = Companion_Default___.SafeModuloInt(_131_v5, _229_cf74)
		_ = _rhs21
		var _rhs22 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_230_v77, false)
		_ = _rhs22
		_131_v5 = _rhs21
		_282_v111 = _rhs22
	} else {
		var _283___mcc_h3 _dafny.MultiSet = _source9.Get_().(D16_DC38).Cf73
		_ = _283___mcc_h3
		var _284_cf73 _dafny.MultiSet = _283___mcc_h3
		_ = _284_cf73
		_126_v0 = _126_v0
		var _285_v112 _dafny.Array
		_ = _285_v112
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(22)
		_ = _len0_6
		var _nw41 _dafny.Array
		_ = _nw41
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw41 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) bool = (func(_286_v59 *C0) func(_dafny.Int) bool {
				return func(_287_i7 _dafny.Int) bool {
					return (_286_v59).F13()
				}
			})(_200_v59)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw41 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw41).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw41).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_285_v112 = _nw41
		var _288_v113 _dafny.Sequence
		_ = _288_v113
		_288_v113 = _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(-204), _131_v5, _131_v5, _131_v5, _131_v5))
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_285_v112), 0))
		_ = _index36
		(_285_v112).ArraySet1((_200_v59).F13(), (_index36).Int())
		var _289_v114 _dafny.Sequence
		_ = _289_v114
		_289_v114 = _dafny.SeqOf(_129_v2, _129_v2, _129_v2)
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_285_v112), 0))
		_ = _index37
		var _rhs23 _dafny.Int = (_dafny.Zero).Minus((((_289_v114).Select((Companion_Default___.SafeIndex(_131_v5, _dafny.IntOfUint32((_289_v114).Cardinality()))).Uint32()).(_dafny.Set)).Difference(_dafny.SetOf((_200_v59).F13()))).Cardinality())
		_ = _rhs23
		var _rhs24 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_126_v0).Cardinality()))).Merge(_219_v72)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v59).F13(), _dafny.IntOfUint32((_126_v0).Cardinality())))
		_ = _rhs24
		var _rhs25 _dafny.Array = _285_v112
		_ = _rhs25
		var _rhs26 _dafny.Sequence = Companion_Default___.Fm40(_127_globalState)
		_ = _rhs26
		var _rhs27 bool = _128_v1
		_ = _rhs27
		var _lhs17 *GlobalState = _127_globalState
		_ = _lhs17
		var _lhs18 _dafny.Array = _285_v112
		_ = _lhs18
		var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_285_v112), 0))
		_ = _lhs19
		_lhs17.F2 = _rhs23
		_219_v72 = _rhs24
		_285_v112 = _rhs25
		_288_v113 = _rhs26
		(_lhs18).ArraySet1(_rhs27, (_lhs19).Int())
		var _290_v115 T1
		_ = _290_v115
		var _nw42 *C1 = New_C1_()
		_ = _nw42
		_nw42.Ctor__()
		_290_v115 = _nw42
		var _291_v116 _dafny.Map
		_ = _291_v116
		_291_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_290_v115, (_285_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_285_v112), 0))).Int()).(bool))
		var _rhs28 _dafny.Map = _291_v116
		_ = _rhs28
		var _rhs29 _dafny.Set = _129_v2
		_ = _rhs29
		var _rhs30 *C0 = _200_v59
		_ = _rhs30
		var _rhs31 _dafny.Int = (_dafny.Zero).Minus((_131_v5).Plus((_dafny.Zero).Minus((_219_v72).Cardinality())))
		_ = _rhs31
		var _rhs32 _dafny.Set = (func() _dafny.Set {
			if (_285_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(421), _dafny.ArrayLen((_285_v112), 0))).Int()).(bool) {
				return (func() _dafny.Set {
					var _coll19 = _dafny.NewBuilder()
					_ = _coll19
					for _iter19 := _dafny.Iterate((_222_v75).Keys().Elements()); ; {
						_compr_19, _ok19 := _iter19()
						if !_ok19 {
							break
						}
						var _292_v117 _dafny.Int
						_292_v117 = interface{}(_compr_19).(_dafny.Int)
						if (_222_v75).Contains(_292_v117) {
							_coll19.Add((_292_v117).Plus(_dafny.IntOfInt64(734)))
						}
					}
					return _coll19.ToSet()
				}()).Difference(_221_v74)
			}
			return func() _dafny.Set {
				var _coll20 = _dafny.NewBuilder()
				_ = _coll20
				for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-261), _dafny.IntOfInt64(593))); ; {
					_compr_20, _ok20 := _iter20()
					if !_ok20 {
						break
					}
					var _293_v118 _dafny.Int
					_293_v118 = interface{}(_compr_20).(_dafny.Int)
					if ((_dafny.IntOfInt64(-261)).Cmp(_293_v118) <= 0) && ((_293_v118).Cmp(_dafny.IntOfInt64(593)) < 0) {
						_coll20.Add((_293_v118).Times(_dafny.IntOfUint32((_126_v0).Cardinality())))
					}
				}
				return _coll20.ToSet()
			}()
		})()
		_ = _rhs32
		var _lhs20 *GlobalState = _127_globalState
		_ = _lhs20
		_291_v116 = _rhs28
		_129_v2 = _rhs29
		_200_v59 = _rhs30
		_lhs20.F0 = _rhs31
		_221_v74 = _rhs32
		var _294_v119 *C2
		_ = _294_v119
		var _nw43 *C2 = New_C2_()
		_ = _nw43
		_nw43.Ctor__(_126_v0, true)
		_294_v119 = _nw43
		_294_v119 = _294_v119
	}
	var _295_v120 _dafny.Set
	_ = _295_v120
	_295_v120 = _dafny.SetOf(_200_v59)
	var _296_v121 *C0
	_ = _296_v121
	_296_v121 = _200_v59
	_295_v120 = _dafny.SetOf(_200_v59, (_200_v59), _200_v59, _200_v59)
	var _297_v122 _dafny.MultiSet
	_ = _297_v122
	_297_v122 = _dafny.MultiSetOf(_128_v1, _128_v1, _128_v1, _128_v1, false)
	var _298_v123 _dafny.Map
	_ = _298_v123
	_298_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_297_v122).Cardinality()).Cmp((_129_v2).Cardinality()) != 0, _129_v2)
	_298_v123 = (_298_v123).Merge((_298_v123).Merge(_298_v123))
	var _hi1 _dafny.Int = (_131_v5).Minus(_131_v5)
	_ = _hi1
	for _299_i8 := _dafny.IntOfUint32((_126_v0).Cardinality()); _299_i8.Cmp(_hi1) < 0; _299_i8 = _299_i8.Plus(_dafny.One) {
		var _rhs33 _dafny.CodePoint = _132_v7
		_ = _rhs33
		var _rhs34 _dafny.Int = (_dafny.Zero).Minus(_299_i8)
		_ = _rhs34
		var _lhs21 *GlobalState = _127_globalState
		_ = _lhs21
		_132_v7 = _rhs33
		_lhs21.F2 = _rhs34
		var _300_v124 _dafny.Array
		_ = _300_v124
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_7
		var _nw44 _dafny.Array
		_ = _nw44
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw44 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) bool = (func(_301_v59 *C0) func(_dafny.Int) bool {
				return func(_302_i9 _dafny.Int) bool {
					return (_301_v59).F13()
				}
			})(_200_v59)
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw44 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw44).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw44).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_300_v124 = _nw44
		var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_300_v124), 0))
		_ = _index38
		(_300_v124).ArraySet1(_128_v1, (_index38).Int())
		var _303_v125 _dafny.Sequence
		_ = _303_v125
		_303_v125 = _dafny.SeqOf(_299_i8, _131_v5)
		var _304_v126 _dafny.Map
		_ = _304_v126
		_304_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v59).F13(), Companion_Default___.Fm8(_132_v7, (_200_v59).F13(), _131_v5, _131_v5, _127_globalState))
		var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_300_v124), 0))
		_ = _index39
		(_300_v124).ArraySet1(Companion_Default___.Fm1(_141_v12, true, (_303_v125).Select((Companion_Default___.SafeIndex(_131_v5, _dafny.IntOfUint32((_303_v125).Cardinality()))).Uint32()).(_dafny.Int), ((_304_v126).Cardinality()).Cmp(_299_i8) < 0, _127_globalState), (_index39).Int())
		(_127_globalState).F1 = true
		var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(752), _dafny.ArrayLen((_300_v124), 0))
		_ = _index40
		(_300_v124).ArraySet1(_128_v1, (_index40).Int())
	}
	Companion_Default___.M7(_131_v5, _127_globalState)
	var _305_v127 _dafny.Sequence
	_ = _305_v127
	_305_v127 = _dafny.SeqOf(_200_v59)
	var _306_v128 _dafny.Array
	_ = _306_v128
	var _nwElement0_8 _dafny.Int = (_dafny.IntOfInt64(773)).Plus(_131_v5)
	_ = _nwElement0_8
	var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(13))
	_ = _nw45
	(_nw45).ArraySet1(_nwElement0_8, 0)
	(_nw45).ArraySet1(Companion_Default___.Fm15((_dafny.Zero).Minus(_131_v5), _128_v1, _127_globalState), 1)
	(_nw45).ArraySet1((_131_v5).Times((_134_v9).Dtor_cf9()), 2)
	(_nw45).ArraySet1(_131_v5, 3)
	(_nw45).ArraySet1(_131_v5, 4)
	(_nw45).ArraySet1((_131_v5).Plus(_dafny.IntOfUint32((_305_v127).Cardinality())), 5)
	(_nw45).ArraySet1((_223_v76).Dtor_cf74(), 6)
	(_nw45).ArraySet1(_131_v5, 7)
	(_nw45).ArraySet1(_131_v5, 8)
	(_nw45).ArraySet1(_131_v5, 9)
	(_nw45).ArraySet1(_131_v5, 10)
	(_nw45).ArraySet1((_131_v5).Minus(_131_v5), 11)
	(_nw45).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
		if (_200_v59).F13() {
			return _131_v5
		}
		return _dafny.IntOfUint32((_dafny.SeqOf(!((_200_v59).F13()), _128_v1, _128_v1)).Cardinality())
	})()), 12)
	_306_v128 = _nw45
	for _iter21 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_306_v128), 0))); ; {
		_guard_loop_0, _ok21 := _iter21()
		if !_ok21 {
			break
		}
		var _307_i10 _dafny.Int
		_307_i10 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_307_i10).Sign() != -1) && ((_307_i10).Cmp(_dafny.ArrayLen((_306_v128), 0)) < 0)) {
			(_306_v128).ArraySet1(Companion_Default___.SafeModuloInt(_307_i10, _131_v5), (_307_i10).Int())
		}
	}
	var _hi2 _dafny.Int = (_131_v5).Plus(_131_v5)
	_ = _hi2
	for _308_i11 := _131_v5; _308_i11.Cmp(_hi2) < 0; _308_i11 = _308_i11.Plus(_dafny.One) {
		var _309_v129 T1
		_ = _309_v129
		var _nw46 *C1 = New_C1_()
		_ = _nw46
		_nw46.Ctor__()
		_309_v129 = _nw46
		var _310_v130 *C6
		_ = _310_v130
		var _nw47 *C6 = New_C6_()
		_ = _nw47
		_nw47.Ctor__(_131_v5)
		_310_v130 = _nw47
		var _311_v131 _dafny.Map
		_ = _311_v131
		_311_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v129, _310_v130)
		_311_v131 = _311_v131
		(_127_globalState).F5 = (_200_v59).F13()
		var _312_v132 _dafny.Array
		_ = _312_v132
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(25)
		_ = _len0_8
		var _nw48 _dafny.Array
		_ = _nw48
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw48 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = (func(_313_v59 *C0) func(_dafny.Int) bool {
				return func(_314_i12 _dafny.Int) bool {
					return (_313_v59).F13()
				}
			})(_200_v59)
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw48 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw48).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw48).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_312_v132 = _nw48
		var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_312_v132), 0))
		_ = _index41
		(_312_v132).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_126_v0, _126_v0), (_index41).Int())
		var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_312_v132), 0))
		_ = _index42
		(_312_v132).ArraySet1((_200_v59).F13(), (_index42).Int())
		_126_v0 = _126_v0
	}
	if (_200_v59).F13() {
		if (_200_v59).F13() {
			var _315_v133 _dafny.Map
			_ = _315_v133
			_315_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v59).Fm23(_131_v5, _131_v5, _132_v7, (_200_v59).F13(), _127_globalState), _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-126))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_316_i13 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vky")).Cardinality())
			})), _dafny.SeqOf(_131_v5)), (Companion_Default___.SafeIndex(_131_v5, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-126))).Uint32(), func(coer13 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg13 _dafny.Int) interface{} {
					return coer13(arg13)
				}
			}(func(_317_i13 _dafny.Int) _dafny.Int {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vky")).Cardinality())
			})), _dafny.SeqOf(_131_v5))).Cardinality()))).Uint32(), _dafny.IntOfInt64(224)))
			var _318_v134 _dafny.Sequence
			_ = _318_v134
			_318_v134 = _dafny.SeqOf(_131_v5)
			_315_v133 = (_315_v133).Update(_131_v5, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(860))).Uint32(), func(coer14 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg14 _dafny.Int) interface{} {
					return coer14(arg14)
				}
			}((func(_319_v0 _dafny.Sequence, _320_v59 *C0) func(_dafny.Int) _dafny.Int {
				return func(_321_i14 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32(((Companion_D12_.Create_DC29_(_319_v0, (_320_v59).F13())).Dtor_cf52()).Cardinality())
				}
			})(_126_v0, _200_v59))), _318_v134))
			var _322_v135 _dafny.Array
			_ = _322_v135
			var _nw49 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw49
			_322_v135 = _nw49
			_322_v135 = _322_v135
			(_127_globalState).F0 = Companion_Default___.SafeDivisionInt(_131_v5, _131_v5)
			(_127_globalState).F0 = ((_318_v134).Select((Companion_Default___.SafeIndex((_200_v59).Fm23(_131_v5, (_222_v75).Cardinality(), _132_v7, _128_v1, _127_globalState), _dafny.IntOfUint32((_318_v134).Cardinality()))).Uint32()).(_dafny.Int)).Minus(((_129_v2).Intersection(_dafny.SetOf(!((_200_v59).F13())))).Cardinality())
			var _323_v136 _dafny.Sequence
			_ = _323_v136
			_323_v136 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-864))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg15 _dafny.Int) interface{} {
					return coer15(arg15)
				}
			}((func(_324_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_325_i15 _dafny.Int) _dafny.CodePoint {
					return _324_v7
				}
			})(_132_v7))), _dafny.UnicodeSeqOfUtf8Bytes("c"), _126_v0, _126_v0)
			Companion_Default___.M7(Companion_Default___.SafeModuloInt(_131_v5, _dafny.IntOfUint32((_323_v136).Cardinality())), _127_globalState)
		} else {
			var _326_v137 _dafny.Array
			_ = _326_v137
			var _len0_9 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_9
			var _nw50 _dafny.Array
			_ = _nw50
			if _len0_9.Cmp(_dafny.Zero) == 0 {
				_nw50 = _dafny.NewArray(_len0_9)
			} else {
				var _init9 func(_dafny.Int) bool = (func(_327_v1 bool) func(_dafny.Int) bool {
					return func(_328_i16 _dafny.Int) bool {
						return _327_v1
					}
				})(_128_v1)
				_ = _init9
				var _element0_9 = _init9(_dafny.Zero)
				_ = _element0_9
				_nw50 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
				(_nw50).ArraySet1(_element0_9, 0)
				var _nativeLen0_9 = (_len0_9).Int()
				_ = _nativeLen0_9
				for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
					(_nw50).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
				}
			}
			_326_v137 = _nw50
			var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_326_v137), 0))
			_ = _index43
			(_326_v137).ArraySet1(!(_128_v1), (_index43).Int())
			var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_306_v128), 0))
			_ = _index44
			(_306_v128).ArraySet1(_131_v5, (_index44).Int())
			var _329_v138 _dafny.Sequence
			_ = _329_v138
			_329_v138 = _dafny.SeqOf(_131_v5)
			var _330_v139 _dafny.Map
			_ = _330_v139
			_330_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_329_v138, _306_v128)
			var _331_v140 D8
			_ = _331_v140
			_331_v140 = Companion_D8_.Create_DC20_(_330_v139)
			var _332_v141 _dafny.Set
			_ = _332_v141
			_332_v141 = _dafny.SetOf(_331_v140, _331_v140)
			var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_326_v137), 0))
			_ = _index45
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_306_v128), 0))
			_ = _index46
			var _rhs35 _dafny.Int = (_dafny.Zero).Minus((_329_v138).Select((Companion_Default___.SafeIndex((_131_v5).Times(_131_v5), _dafny.IntOfUint32((_329_v138).Cardinality()))).Uint32()).(_dafny.Int))
			_ = _rhs35
			var _rhs36 bool = (_133_v8).IsSubsetOf(_133_v8)
			_ = _rhs36
			var _rhs37 bool = !(((_332_v141).Intersection(_332_v141)).IsSubsetOf(_332_v141))
			_ = _rhs37
			var _rhs38 bool = (_200_v59).F13()
			_ = _rhs38
			var _rhs39 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_126_v0).Cardinality()), Companion_Default___.SafeModuloInt(_131_v5, _131_v5))
			_ = _rhs39
			var _lhs22 *GlobalState = _127_globalState
			_ = _lhs22
			var _lhs23 _dafny.Array = _326_v137
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_326_v137), 0))
			_ = _lhs24
			var _lhs25 *GlobalState = _127_globalState
			_ = _lhs25
			var _lhs26 *GlobalState = _127_globalState
			_ = _lhs26
			var _lhs27 _dafny.Array = _306_v128
			_ = _lhs27
			var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_306_v128), 0))
			_ = _lhs28
			_lhs22.F0 = _rhs35
			(_lhs23).ArraySet1(_rhs36, (_lhs24).Int())
			_lhs25.F1 = _rhs37
			_lhs26.F5 = _rhs38
			(_lhs27).ArraySet1(_rhs39, (_lhs28).Int())
			var _333_v142 D5
			_ = _333_v142
			_333_v142 = Companion_D5_.Create_DC13_(_128_v1, _128_v1, _126_v0, _dafny.IntOfInt64(-978))
			var _334_v143 D3
			_ = _334_v143
			_334_v143 = Companion_D3_.Create_DC7_((_333_v142).Dtor_cf23())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_326_v137), 0))
			_ = _index47
			var _rhs40 D3 = _334_v143
			_ = _rhs40
			var _rhs41 _dafny.Int = (_dafny.IntOfUint32((_126_v0).Cardinality())).Plus((func() _dafny.Int {
				if true {
					return (_306_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_306_v128), 0))).Int()).(_dafny.Int)
				}
				return _131_v5
			})())
			_ = _rhs41
			var _rhs42 bool = (_200_v59).F13()
			_ = _rhs42
			var _lhs29 *GlobalState = _127_globalState
			_ = _lhs29
			var _lhs30 _dafny.Array = _326_v137
			_ = _lhs30
			var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(331), _dafny.ArrayLen((_326_v137), 0))
			_ = _lhs31
			_334_v143 = _rhs40
			_lhs29.F2 = _rhs41
			(_lhs30).ArraySet1(_rhs42, (_lhs31).Int())
			var _335_v144 D8
			_ = _335_v144
			_335_v144 = Companion_D8_.Create_DC21_((_200_v59).F13(), false, (_dafny.IntOfUint32((_126_v0).Cardinality())).Times((_dafny.MultiSetOf(_128_v1, _128_v1, (_200_v59).F13(), _128_v1)).Cardinality()))
			_335_v144 = _335_v144
			var _336_v145 *C4
			_ = _336_v145
			var _nw51 *C4 = New_C4_()
			_ = _nw51
			_nw51.Ctor__((_306_v128).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(785), _dafny.ArrayLen((_306_v128), 0))).Int()).(_dafny.Int))
			_336_v145 = _nw51
			var _337_v146 _dafny.Array
			_ = _337_v146
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(15)
			_ = _len0_10
			var _nw52 _dafny.Array
			_ = _nw52
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw52 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Int = (func(_338_v5 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_339_i17 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_339_i17, _338_v5)
					}
				})(_131_v5)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw52 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw52).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw52).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_337_v146 = _nw52
			var _340_v147 _dafny.Map
			_ = _340_v147
			_340_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_337_v146, (_dafny.Zero).Minus(_131_v5))
			(_127_globalState).F5 = (_340_v147).Contains(_337_v146)
		}
		(_127_globalState).F5 = !((_200_v59).F13())
		var _341_v148 _dafny.Array
		_ = _341_v148
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
		_ = _nw53
		_341_v148 = _nw53
		var _342_v149 _dafny.Array
		_ = _342_v149
		var _nw54 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
		_ = _nw54
		_342_v149 = _nw54
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_341_v148), 0))
		_ = _index48
		(_341_v148).ArraySet1(_342_v149, (_index48).Int())
		var _343_v150 T1
		_ = _343_v150
		var _nw55 *C1 = New_C1_()
		_ = _nw55
		_nw55.Ctor__()
		_343_v150 = _nw55
		var _344_v151 _dafny.Array
		_ = _344_v151
		var _nwElement0_9 bool = (_200_v59).F13()
		_ = _nwElement0_9
		var _nw56 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(22))
		_ = _nw56
		(_nw56).ArraySet1(_nwElement0_9, 0)
		(_nw56).ArraySet1((_200_v59).F13(), 1)
		(_nw56).ArraySet1((_200_v59).F13(), 2)
		(_nw56).ArraySet1(_128_v1, 3)
		(_nw56).ArraySet1((_200_v59).F13(), 4)
		(_nw56).ArraySet1((_200_v59).F13(), 5)
		(_nw56).ArraySet1((_200_v59).F13(), 6)
		(_nw56).ArraySet1(true, 7)
		(_nw56).ArraySet1((_200_v59).F13(), 8)
		(_nw56).ArraySet1(!(_128_v1), 9)
		(_nw56).ArraySet1(_128_v1, 10)
		(_nw56).ArraySet1(!((_200_v59).F13()), 11)
		(_nw56).ArraySet1((_200_v59).F13(), 12)
		(_nw56).ArraySet1(false, 13)
		(_nw56).ArraySet1((_200_v59).F13(), 14)
		(_nw56).ArraySet1((_200_v59).F13(), 15)
		(_nw56).ArraySet1((_200_v59).F13(), 16)
		(_nw56).ArraySet1((_200_v59).F13(), 17)
		(_nw56).ArraySet1(!((_200_v59).F13()), 18)
		(_nw56).ArraySet1(!((_200_v59).F13()), 19)
		(_nw56).ArraySet1(!(true), 20)
		(_nw56).ArraySet1(_128_v1, 21)
		_344_v151 = _nw56
		var _345_v152 D11
		_ = _345_v152
		_345_v152 = Companion_D11_.Create_DC25_(_344_v151, _343_v150, (_221_v74).Cardinality())
		var _346_v153 _dafny.Array
		_ = _346_v153
		var _nwElement0_10 T1 = _343_v150
		_ = _nwElement0_10
		var _nw57 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(27))
		_ = _nw57
		(_nw57).ArraySet1(_nwElement0_10, 0)
		(_nw57).ArraySet1(_343_v150, 1)
		(_nw57).ArraySet1(_343_v150, 2)
		(_nw57).ArraySet1((func() T1 {
			if (_200_v59).F13() {
				return _343_v150
			}
			return _343_v150
		})(), 3)
		(_nw57).ArraySet1(_343_v150, 4)
		(_nw57).ArraySet1(_343_v150, 5)
		(_nw57).ArraySet1(_343_v150, 6)
		(_nw57).ArraySet1(_343_v150, 7)
		(_nw57).ArraySet1(_343_v150, 8)
		(_nw57).ArraySet1(_343_v150, 9)
		(_nw57).ArraySet1((_345_v152).Dtor_cf43(), 10)
		(_nw57).ArraySet1(_343_v150, 11)
		(_nw57).ArraySet1(_343_v150, 12)
		(_nw57).ArraySet1(_343_v150, 13)
		(_nw57).ArraySet1(_343_v150, 14)
		(_nw57).ArraySet1(_343_v150, 15)
		(_nw57).ArraySet1(_343_v150, 16)
		(_nw57).ArraySet1(_343_v150, 17)
		(_nw57).ArraySet1(_343_v150, 18)
		(_nw57).ArraySet1(_343_v150, 19)
		(_nw57).ArraySet1(_343_v150, 20)
		(_nw57).ArraySet1(_343_v150, 21)
		(_nw57).ArraySet1(_343_v150, 22)
		(_nw57).ArraySet1(_343_v150, 23)
		(_nw57).ArraySet1(_343_v150, 24)
		(_nw57).ArraySet1(_343_v150, 25)
		(_nw57).ArraySet1(_343_v150, 26)
		_346_v153 = _nw57
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_346_v153), 0))
		_ = _index49
		(_346_v153).ArraySet1(_343_v150, (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_341_v148), 0))
		_ = _index50
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_346_v153), 0))
		_ = _index51
		var _rhs43 _dafny.Array = _342_v149
		_ = _rhs43
		var _rhs44 T1 = _343_v150
		_ = _rhs44
		var _lhs32 _dafny.Array = _341_v148
		_ = _lhs32
		var _lhs33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(896), _dafny.ArrayLen((_341_v148), 0))
		_ = _lhs33
		var _lhs34 _dafny.Array = _346_v153
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_346_v153), 0))
		_ = _lhs35
		(_lhs32).ArraySet1(_rhs43, (_lhs33).Int())
		(_lhs34).ArraySet1(_rhs44, (_lhs35).Int())
		_221_v74 = func() _dafny.Set {
			var _coll21 = _dafny.NewBuilder()
			_ = _coll21
			for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(306), _dafny.IntOfInt64(-263))); ; {
				_compr_21, _ok22 := _iter22()
				if !_ok22 {
					break
				}
				var _347_v154 _dafny.Int
				_347_v154 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(306)).Cmp(_347_v154) <= 0) && ((_347_v154).Cmp(_dafny.IntOfInt64(-263)) < 0) {
					_coll21.Add((_347_v154).Minus(_131_v5))
				}
			}
			return _coll21.ToSet()
		}()
		if !(_128_v1) || ((func() bool {
			if (_200_v59).F13() {
				return Companion_Default___.Fm1(_141_v12, (_200_v59).F13(), _131_v5, (_200_v59).F13(), _127_globalState)
			}
			return (_200_v59).F13()
		})()) {
			var _348_v155 *C5
			_ = _348_v155
			var _nw58 *C5 = New_C5_()
			_ = _nw58
			_nw58.Ctor__(!(!(_128_v1)), _132_v7)
			_348_v155 = _nw58
			var _349_v156 *C1
			_ = _349_v156
			var _nw59 *C1 = New_C1_()
			_ = _nw59
			_nw59.Ctor__()
			_349_v156 = _nw59
			var _350_v157 D18
			_ = _350_v157
			_350_v157 = Companion_D18_.Create_DC43_(_349_v156)
			var _351_v158 _dafny.Map
			_ = _351_v158
			_351_v158 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _349_v156)
			var _352_v159 _dafny.Sequence
			_ = _352_v159
			_352_v159 = _dafny.SeqOf(_131_v5, _131_v5)
			var _353_v160 T0
			_ = _353_v160
			var _nw60 *C3 = New_C3_()
			_ = _nw60
			_nw60.Ctor__(_131_v5, _352_v159, _dafny.IntOfInt64(507))
			_353_v160 = _nw60
			var _354_v161 _dafny.Map
			_ = _354_v161
			_354_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_353_v160, _dafny.CodePoint('k'))
			var _355_v162 _dafny.MultiSet
			_ = _355_v162
			_355_v162 = _dafny.MultiSetOf(_354_v161, _354_v161)
			var _356_v163 _dafny.Array
			_ = _356_v163
			var _nwElement0_11 *C1 = _349_v156
			_ = _nwElement0_11
			var _nw61 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(24))
			_ = _nw61
			(_nw61).ArraySet1(_nwElement0_11, 0)
			(_nw61).ArraySet1(_349_v156, 1)
			(_nw61).ArraySet1(_349_v156, 2)
			(_nw61).ArraySet1(_349_v156, 3)
			(_nw61).ArraySet1(_349_v156, 4)
			(_nw61).ArraySet1(_349_v156, 5)
			(_nw61).ArraySet1(_349_v156, 6)
			(_nw61).ArraySet1(_349_v156, 7)
			(_nw61).ArraySet1(_349_v156, 8)
			(_nw61).ArraySet1(_349_v156, 9)
			(_nw61).ArraySet1(_349_v156, 10)
			(_nw61).ArraySet1(_349_v156, 11)
			(_nw61).ArraySet1(_349_v156, 12)
			(_nw61).ArraySet1((_350_v157).Dtor_cf80(), 13)
			(_nw61).ArraySet1(_349_v156, 14)
			(_nw61).ArraySet1(_349_v156, 15)
			(_nw61).ArraySet1(_349_v156, 16)
			(_nw61).ArraySet1(_349_v156, 17)
			(_nw61).ArraySet1(_349_v156, 18)
			(_nw61).ArraySet1(_349_v156, 19)
			(_nw61).ArraySet1((func() *C1 {
				if (_351_v158).Contains((_355_v162).Cardinality()) {
					return (_351_v158).Get((_355_v162).Cardinality()).(*C1)
				}
				return _349_v156
			})(), 20)
			(_nw61).ArraySet1(_349_v156, 21)
			(_nw61).ArraySet1(_349_v156, 22)
			(_nw61).ArraySet1(_349_v156, 23)
			_356_v163 = _nw61
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_356_v163), 0))
			_ = _index52
			(_356_v163).ArraySet1(_349_v156, (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(305), _dafny.ArrayLen((_356_v163), 0))
			_ = _index53
			(_356_v163).ArraySet1(_349_v156, (_index53).Int())
			(_127_globalState).F5 = (_200_v59).F13()
			_132_v7 = _132_v7
			var _357_v164 _dafny.Array
			_ = _357_v164
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
			_ = _nw62
			_357_v164 = _nw62
			var _358_v165 D19
			_ = _358_v165
			_358_v165 = Companion_D19_.Create_DC45_(_357_v164)
			_357_v164 = (_358_v165).Dtor_cf86()
		} else {
			var _359_v166 _dafny.Map
			_ = _359_v166
			_359_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_T1_.CastTo_((_346_v153).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_346_v153), 0))).Int())), _128_v1)
			var _360_v167 D13
			_ = _360_v167
			_360_v167 = Companion_D13_.Create_DC32_(_131_v5, false, (_200_v59).F13(), (_200_v59).F13(), (_200_v59).F13())
			_359_v166 = (_359_v166).Update(Companion_T1_.CastTo_((_346_v153).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(892), _dafny.ArrayLen((_346_v153), 0))).Int())), (_360_v167).Dtor_cf57())
			var _361_v168 _dafny.Sequence
			_ = _361_v168
			_361_v168 = _dafny.SeqOf(_128_v1)
			(_127_globalState).F1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_361_v168, _361_v168), _dafny.Companion_Sequence_.Concatenate(_361_v168, _361_v168))
			var _362_v169 _dafny.Array
			_ = _362_v169
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_11
			var _nw63 _dafny.Array
			_ = _nw63
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw63 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.CodePoint = (func(_363_v7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_364_i18 _dafny.Int) _dafny.CodePoint {
						return _363_v7
					}
				})(_132_v7)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw63 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw63).ArraySet1CodePoint(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw63).ArraySet1CodePoint(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_362_v169 = _nw63
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_362_v169), 0))
			_ = _index54
			(_362_v169).ArraySet1CodePoint(_132_v7, (_index54).Int())
			var _365_v170 _dafny.Map
			_ = _365_v170
			_365_v170 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_200_v59).F13(), (_200_v59).F13())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_344_v151), 0))
			_ = _index55
			(_344_v151).ArraySet1((func() bool {
				if (_365_v170).Contains(_128_v1) {
					return (_365_v170).Get(_128_v1).(bool)
				}
				return _128_v1
			})(), (_index55).Int())
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_362_v169), 0))
			_ = _index56
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_344_v151), 0))
			_ = _index57
			var _rhs45 _dafny.CodePoint = (Companion_Default___.Fm19((_200_v59).F13(), _132_v7, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_219_v72).Contains(_128_v1) {
					return (_219_v72).Get(_128_v1).(_dafny.Int)
				}
				return _131_v5
			})(), true), _128_v1, _127_globalState)).Dtor_cf6()
			_ = _rhs45
			var _rhs46 bool = !((_200_v59).F13()) || (!(_dafny.Companion_Sequence_.IsProperPrefixOf(_126_v0, _dafny.Companion_Sequence_.Update(_126_v0, (Companion_Default___.SafeIndex(_131_v5, _dafny.IntOfUint32((_126_v0).Cardinality()))).Uint32(), _132_v7))))
			_ = _rhs46
			var _lhs36 _dafny.Array = _362_v169
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen((_362_v169), 0))
			_ = _lhs37
			var _lhs38 _dafny.Array = _344_v151
			_ = _lhs38
			var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(904), _dafny.ArrayLen((_344_v151), 0))
			_ = _lhs39
			(_lhs36).ArraySet1CodePoint(_rhs45, (_lhs37).Int())
			(_lhs38).ArraySet1(_rhs46, (_lhs39).Int())
			Companion_Default___.M7(_131_v5, _127_globalState)
			(_127_globalState).F0 = (_dafny.Zero).Minus((Companion_Default___.SafeDivisionInt(_131_v5, _dafny.IntOfUint32((_126_v0).Cardinality()))).Times((_200_v59).Fm23(_131_v5, _131_v5, _dafny.CodePoint('n'), (Companion_D8_.Create_DC21_(_128_v1, (_200_v59).F13(), _131_v5)).Dtor_cf37(), _127_globalState)))
		}
	} else {
		var _366_v171 *C4
		_ = _366_v171
		var _nw64 *C4 = New_C4_()
		_ = _nw64
		_nw64.Ctor__(_131_v5)
		_366_v171 = _nw64
		var _367_v172 _dafny.Sequence
		_ = _367_v172
		_367_v172 = _dafny.SeqOf(_dafny.IntOfInt64(771), _131_v5, _131_v5)
		var _368_v173 _dafny.Map
		_ = _368_v173
		_368_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _126_v0)
		var _369_v174 _dafny.MultiSet
		_ = _369_v174
		_369_v174 = _dafny.MultiSetOf(_367_v172, _dafny.SeqOf(_dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_368_v173).Contains(_131_v5) {
				return (_368_v173).Get(_131_v5).(_dafny.Sequence)
			}
			return _126_v0
		})()).Cardinality()), _dafny.IntOfUint32((Companion_Default___.Fm21(_126_v0, _127_globalState)).Cardinality())), _367_v172)
		var _370_v175 _dafny.Sequence
		_ = _370_v175
		_370_v175 = _dafny.SeqOf(_306_v128, _306_v128)
		var _371_v176 _dafny.Map
		_ = _371_v176
		_371_v176 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.MultiSet {
			if (_200_v59).F13() {
				return _dafny.MultiSetFromSeq(Companion_Default___.Fm40(_127_globalState))
			}
			return _369_v174
		})(), Companion_Default___.Fm15(Companion_Default___.Fm15((_dafny.Zero).Minus(_dafny.IntOfUint32((_370_v175).Cardinality())), (_200_v59).F13(), _127_globalState), (_200_v59).F13(), _127_globalState))
		_371_v176 = (_371_v176).Update(_369_v174, _131_v5)
		var _372_v177 _dafny.Array
		_ = _372_v177
		var _nw65 _dafny.Array = _dafny.NewArrayWithValue(Companion_D14_.Default(), _dafny.IntOfInt64(21))
		_ = _nw65
		_372_v177 = _nw65
		_372_v177 = _372_v177
		var _373_v178 _dafny.Map
		_ = _373_v178
		_373_v178 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_131_v5, _132_v7)
		var _374_v179 *C3
		_ = _374_v179
		var _nw66 *C3 = New_C3_()
		_ = _nw66
		_nw66.Ctor__(Companion_Default___.SafeModuloInt(_131_v5, _dafny.IntOfUint32((_dafny.SeqOf((_200_v59).F13(), (_200_v59).F13(), _128_v1, _128_v1)).Cardinality())), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_373_v178).Cardinality()), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm2(_126_v0, _127_globalState)).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf((_373_v178).Cardinality())).Cardinality()))).Uint32(), _131_v5), _131_v5)
		_374_v179 = _nw66
		_128_v1 = !(_128_v1) || (false)
	}
	(_127_globalState).F0 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_126_v0, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(788))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_375_i19 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('l')
	}))), _dafny.UnicodeSeqOfUtf8Bytes("oobul"))).Cardinality())
	_dafny.Print(_126_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_127_globalState).F4().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_127_globalState.F5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_128_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_129_v2).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_130_v3).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_131_v5)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_132_v7)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_133_v8).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(872), _dafny.IntOfInt64(372), _dafny.IntOfInt64(371), _dafny.IntOfInt64(371))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v9).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v9).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v9).Dtor_cf7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v9).Dtor_cf8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_134_v9).Dtor_cf9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(742), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(371), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(371), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Zero, _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_135_v10).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_140_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(371), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_141_v12).Dtor_cf0()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(277), _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_200_v59).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_219_v72).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(371))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_220_v73).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(371), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_221_v74).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_222_v75).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(371))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_223_v76).Dtor_cf74())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_223_v76).Dtor_cf75()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.One, _dafny.IntOfInt64(371))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_295_v120).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_296_v121).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_297_v122).Equals(_dafny.MultiSetOf(true, true, true, true, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_298_v123).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SetOf(false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_305_v127).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_306_v128).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
}

// End of class Default__

// Definition of datatype D0
type D0 struct {
	Data_D0_
}

func (_this D0) Get_() Data_D0_ {
	return _this.Data_D0_
}

type Data_D0_ interface {
	isD0()
}

type CompanionStruct_D0_ struct {
}

var Companion_D0_ = CompanionStruct_D0_{}

type D0_DC1 struct {
	Cf1 _dafny.Int
	Cf2 bool
	Cf3 _dafny.Int
	Cf4 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf1 _dafny.Int, Cf2 bool, Cf3 _dafny.Int, Cf4 bool) D0 {
	return D0{D0_DC1{Cf1, Cf2, Cf3, Cf4}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf5 _dafny.Int
	Cf6 _dafny.CodePoint
	Cf7 _dafny.Int
	Cf8 _dafny.Int
	Cf9 _dafny.Int
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf5 _dafny.Int, Cf6 _dafny.CodePoint, Cf7 _dafny.Int, Cf8 _dafny.Int, Cf9 _dafny.Int) D0 {
	return D0{D0_DC2{Cf5, Cf6, Cf7, Cf8, Cf9}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Map
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Map) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC3 struct {
	Cf10 D0
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf10 D0) D0 {
	return D0{D0_DC3{Cf10}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_(_dafny.Zero, false, _dafny.Zero, false)
}

func (_this D0) Dtor_cf1() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf1
}

func (_this D0) Dtor_cf2() bool {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() bool {
	return _this.Get_().(D0_DC1).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf7
}

func (_this D0) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf8
}

func (_this D0) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf9
}

func (_this D0) Dtor_cf0() _dafny.Map {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf10() D0 {
	return _this.Get_().(D0_DC3).Cf10
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ", " + _dafny.String(data.Cf4) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ")"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0 && data1.Cf4 == data2.Cf4
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6 && data1.Cf7.Cmp(data2.Cf7) == 0 && data1.Cf8.Cmp(data2.Cf8) == 0 && data1.Cf9.Cmp(data2.Cf9) == 0
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Equals(data2.Cf0)
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf10.Equals(data2.Cf10)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D0) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D0)
	return ok && _this.Equals(typed)
}

func Type_D0_() _dafny.TypeDescriptor {
	return type_D0_{}
}

type type_D0_ struct {
}

func (_this type_D0_) Default() interface{} {
	return Companion_D0_.Default()
}

func (_this type_D0_) String() string {
	return "main.D0"
}
func (_this D0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D0{}

// End of datatype D0

// Definition of datatype D1
type D1 struct {
	Data_D1_
}

func (_this D1) Get_() Data_D1_ {
	return _this.Data_D1_
}

type Data_D1_ interface {
	isD1()
}

type CompanionStruct_D1_ struct {
}

var Companion_D1_ = CompanionStruct_D1_{}

type D1_DC5 struct {
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_() D1 {
	return D1{D1_DC5{}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf11 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf11 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf11}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_()
}

func (_this D1) Dtor_cf11() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf11) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC5:
		{
			_, ok := other.Get_().(D1_DC5)
			return ok
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf11.Equals(data2.Cf11)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D1) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D1)
	return ok && _this.Equals(typed)
}

func Type_D1_() _dafny.TypeDescriptor {
	return type_D1_{}
}

type type_D1_ struct {
}

func (_this type_D1_) Default() interface{} {
	return Companion_D1_.Default()
}

func (_this type_D1_) String() string {
	return "main.D1"
}
func (_this D1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D1{}

// End of datatype D1

// Definition of datatype D2
type D2 struct {
	Data_D2_
}

func (_this D2) Get_() Data_D2_ {
	return _this.Data_D2_
}

type Data_D2_ interface {
	isD2()
}

type CompanionStruct_D2_ struct {
}

var Companion_D2_ = CompanionStruct_D2_{}

type D2_DC6 struct {
	Cf12 _dafny.Array
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf12 _dafny.Array) D2 {
	return D2{D2_DC6{Cf12}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() _dafny.Array {
	return _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
}

func (_this D2) Dtor_cf12() _dafny.Array {
	return _this.Get_().(D2_DC6).Cf12
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf12) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf12 == data2.Cf12
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D2) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D2)
	return ok && _this.Equals(typed)
}

func Type_D2_() _dafny.TypeDescriptor {
	return type_D2_{}
}

type type_D2_ struct {
}

func (_this type_D2_) Default() interface{} {
	return Companion_D2_.Default()
}

func (_this type_D2_) String() string {
	return "main.D2"
}
func (_this D2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D2{}

// End of datatype D2

// Definition of datatype D3
type D3 struct {
	Data_D3_
}

func (_this D3) Get_() Data_D3_ {
	return _this.Data_D3_
}

type Data_D3_ interface {
	isD3()
}

type CompanionStruct_D3_ struct {
}

var Companion_D3_ = CompanionStruct_D3_{}

type D3_DC8 struct {
	Cf14 D1
	Cf15 _dafny.Array
	Cf16 _dafny.Int
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 D1, Cf15 _dafny.Array, Cf16 _dafny.Int) D3 {
	return D3{D3_DC8{Cf14, Cf15, Cf16}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC7 struct {
	Cf13 _dafny.Sequence
}

func (D3_DC7) isD3() {}

func (CompanionStruct_D3_) Create_DC7_(Cf13 _dafny.Sequence) D3 {
	return D3{D3_DC7{Cf13}}
}

func (_this D3) Is_DC7() bool {
	_, ok := _this.Get_().(D3_DC7)
	return ok
}

type D3_DC9 struct {
	Cf17 D3
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf17 D3) D3 {
	return D3{D3_DC9{Cf17}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC8_(Companion_D1_.Default(), _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.Zero)
}

func (_this D3) Dtor_cf14() D1 {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf15() _dafny.Array {
	return _this.Get_().(D3_DC8).Cf15
}

func (_this D3) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D3_DC8).Cf16
}

func (_this D3) Dtor_cf13() _dafny.Sequence {
	return _this.Get_().(D3_DC7).Cf13
}

func (_this D3) Dtor_cf17() D3 {
	return _this.Get_().(D3_DC9).Cf17
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ", " + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D3_DC7:
		{
			return "D3.DC7" + "(" + data.Cf13.VerbatimString(true) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf17) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf14.Equals(data2.Cf14) && data1.Cf15 == data2.Cf15 && data1.Cf16.Cmp(data2.Cf16) == 0
		}
	case D3_DC7:
		{
			data2, ok := other.Get_().(D3_DC7)
			return ok && data1.Cf13.Equals(data2.Cf13)
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D3) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D3)
	return ok && _this.Equals(typed)
}

func Type_D3_() _dafny.TypeDescriptor {
	return type_D3_{}
}

type type_D3_ struct {
}

func (_this type_D3_) Default() interface{} {
	return Companion_D3_.Default()
}

func (_this type_D3_) String() string {
	return "main.D3"
}
func (_this D3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D3{}

// End of datatype D3

// Definition of datatype D4
type D4 struct {
	Data_D4_
}

func (_this D4) Get_() Data_D4_ {
	return _this.Data_D4_
}

type Data_D4_ interface {
	isD4()
}

type CompanionStruct_D4_ struct {
}

var Companion_D4_ = CompanionStruct_D4_{}

type D4_DC10 struct {
	Cf18 _dafny.Map
}

func (D4_DC10) isD4() {}

func (CompanionStruct_D4_) Create_DC10_(Cf18 _dafny.Map) D4 {
	return D4{D4_DC10{Cf18}}
}

func (_this D4) Is_DC10() bool {
	_, ok := _this.Get_().(D4_DC10)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D4) Dtor_cf18() _dafny.Map {
	return _this.Get_().(D4_DC10).Cf18
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC10:
		{
			return "D4.DC10" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC10:
		{
			data2, ok := other.Get_().(D4_DC10)
			return ok && data1.Cf18.Equals(data2.Cf18)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D4) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D4)
	return ok && _this.Equals(typed)
}

func Type_D4_() _dafny.TypeDescriptor {
	return type_D4_{}
}

type type_D4_ struct {
}

func (_this type_D4_) Default() interface{} {
	return Companion_D4_.Default()
}

func (_this type_D4_) String() string {
	return "main.D4"
}
func (_this D4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D4{}

// End of datatype D4

// Definition of datatype D5
type D5 struct {
	Data_D5_
}

func (_this D5) Get_() Data_D5_ {
	return _this.Data_D5_
}

type Data_D5_ interface {
	isD5()
}

type CompanionStruct_D5_ struct {
}

var Companion_D5_ = CompanionStruct_D5_{}

type D5_DC12 struct {
	Cf20 _dafny.Int
}

func (D5_DC12) isD5() {}

func (CompanionStruct_D5_) Create_DC12_(Cf20 _dafny.Int) D5 {
	return D5{D5_DC12{Cf20}}
}

func (_this D5) Is_DC12() bool {
	_, ok := _this.Get_().(D5_DC12)
	return ok
}

type D5_DC13 struct {
	Cf21 bool
	Cf22 bool
	Cf23 _dafny.Sequence
	Cf24 _dafny.Int
}

func (D5_DC13) isD5() {}

func (CompanionStruct_D5_) Create_DC13_(Cf21 bool, Cf22 bool, Cf23 _dafny.Sequence, Cf24 _dafny.Int) D5 {
	return D5{D5_DC13{Cf21, Cf22, Cf23, Cf24}}
}

func (_this D5) Is_DC13() bool {
	_, ok := _this.Get_().(D5_DC13)
	return ok
}

type D5_DC11 struct {
	Cf19 _dafny.Set
}

func (D5_DC11) isD5() {}

func (CompanionStruct_D5_) Create_DC11_(Cf19 _dafny.Set) D5 {
	return D5{D5_DC11{Cf19}}
}

func (_this D5) Is_DC11() bool {
	_, ok := _this.Get_().(D5_DC11)
	return ok
}

type D5_DC14 struct {
	Cf25 D5
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf25 D5) D5 {
	return D5{D5_DC14{Cf25}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC12_(_dafny.Zero)
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC12).Cf20
}

func (_this D5) Dtor_cf21() bool {
	return _this.Get_().(D5_DC13).Cf21
}

func (_this D5) Dtor_cf22() bool {
	return _this.Get_().(D5_DC13).Cf22
}

func (_this D5) Dtor_cf23() _dafny.Sequence {
	return _this.Get_().(D5_DC13).Cf23
}

func (_this D5) Dtor_cf24() _dafny.Int {
	return _this.Get_().(D5_DC13).Cf24
}

func (_this D5) Dtor_cf19() _dafny.Set {
	return _this.Get_().(D5_DC11).Cf19
}

func (_this D5) Dtor_cf25() D5 {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC12:
		{
			return "D5.DC12" + "(" + _dafny.String(data.Cf20) + ")"
		}
	case D5_DC13:
		{
			return "D5.DC13" + "(" + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + data.Cf23.VerbatimString(true) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC11:
		{
			return "D5.DC11" + "(" + _dafny.String(data.Cf19) + ")"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf25) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC12:
		{
			data2, ok := other.Get_().(D5_DC12)
			return ok && data1.Cf20.Cmp(data2.Cf20) == 0
		}
	case D5_DC13:
		{
			data2, ok := other.Get_().(D5_DC13)
			return ok && data1.Cf21 == data2.Cf21 && data1.Cf22 == data2.Cf22 && data1.Cf23.Equals(data2.Cf23) && data1.Cf24.Cmp(data2.Cf24) == 0
		}
	case D5_DC11:
		{
			data2, ok := other.Get_().(D5_DC11)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf25.Equals(data2.Cf25)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D5) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D5)
	return ok && _this.Equals(typed)
}

func Type_D5_() _dafny.TypeDescriptor {
	return type_D5_{}
}

type type_D5_ struct {
}

func (_this type_D5_) Default() interface{} {
	return Companion_D5_.Default()
}

func (_this type_D5_) String() string {
	return "main.D5"
}
func (_this D5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D5{}

// End of datatype D5

// Definition of datatype D6
type D6 struct {
	Data_D6_
}

func (_this D6) Get_() Data_D6_ {
	return _this.Data_D6_
}

type Data_D6_ interface {
	isD6()
}

type CompanionStruct_D6_ struct {
}

var Companion_D6_ = CompanionStruct_D6_{}

type D6_DC16 struct {
	Cf27 _dafny.Int
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf27 _dafny.Int) D6 {
	return D6{D6_DC16{Cf27}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

type D6_DC15 struct {
	Cf26 _dafny.Set
}

func (D6_DC15) isD6() {}

func (CompanionStruct_D6_) Create_DC15_(Cf26 _dafny.Set) D6 {
	return D6{D6_DC15{Cf26}}
}

func (_this D6) Is_DC15() bool {
	_, ok := _this.Get_().(D6_DC15)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC16_(_dafny.Zero)
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC16).Cf27
}

func (_this D6) Dtor_cf26() _dafny.Set {
	return _this.Get_().(D6_DC15).Cf26
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC15:
		{
			return "D6.DC15" + "(" + _dafny.String(data.Cf26) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D6_DC15:
		{
			data2, ok := other.Get_().(D6_DC15)
			return ok && data1.Cf26.Equals(data2.Cf26)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D6) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D6)
	return ok && _this.Equals(typed)
}

func Type_D6_() _dafny.TypeDescriptor {
	return type_D6_{}
}

type type_D6_ struct {
}

func (_this type_D6_) Default() interface{} {
	return Companion_D6_.Default()
}

func (_this type_D6_) String() string {
	return "main.D6"
}
func (_this D6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D6{}

// End of datatype D6

// Definition of datatype D7
type D7 struct {
	Data_D7_
}

func (_this D7) Get_() Data_D7_ {
	return _this.Data_D7_
}

type Data_D7_ interface {
	isD7()
}

type CompanionStruct_D7_ struct {
}

var Companion_D7_ = CompanionStruct_D7_{}

type D7_DC18 struct {
	Cf29 bool
	Cf30 bool
	Cf31 _dafny.Int
	Cf32 bool
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf29 bool, Cf30 bool, Cf31 _dafny.Int, Cf32 bool) D7 {
	return D7{D7_DC18{Cf29, Cf30, Cf31, Cf32}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

type D7_DC19 struct {
	Cf33 _dafny.MultiSet
	Cf34 _dafny.Int
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf33 _dafny.MultiSet, Cf34 _dafny.Int) D7 {
	return D7{D7_DC19{Cf33, Cf34}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC17 struct {
	Cf28 _dafny.Sequence
}

func (D7_DC17) isD7() {}

func (CompanionStruct_D7_) Create_DC17_(Cf28 _dafny.Sequence) D7 {
	return D7{D7_DC17{Cf28}}
}

func (_this D7) Is_DC17() bool {
	_, ok := _this.Get_().(D7_DC17)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC18_(false, false, _dafny.Zero, false)
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC18).Cf29
}

func (_this D7) Dtor_cf30() bool {
	return _this.Get_().(D7_DC18).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D7_DC18).Cf31
}

func (_this D7) Dtor_cf32() bool {
	return _this.Get_().(D7_DC18).Cf32
}

func (_this D7) Dtor_cf33() _dafny.MultiSet {
	return _this.Get_().(D7_DC19).Cf33
}

func (_this D7) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf34
}

func (_this D7) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D7_DC17).Cf28
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ", " + _dafny.String(data.Cf32) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D7_DC17:
		{
			return "D7.DC17" + "(" + _dafny.String(data.Cf28) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf29 == data2.Cf29 && data1.Cf30 == data2.Cf30 && data1.Cf31.Cmp(data2.Cf31) == 0 && data1.Cf32 == data2.Cf32
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf33.Equals(data2.Cf33) && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D7_DC17:
		{
			data2, ok := other.Get_().(D7_DC17)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D7) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D7)
	return ok && _this.Equals(typed)
}

func Type_D7_() _dafny.TypeDescriptor {
	return type_D7_{}
}

type type_D7_ struct {
}

func (_this type_D7_) Default() interface{} {
	return Companion_D7_.Default()
}

func (_this type_D7_) String() string {
	return "main.D7"
}
func (_this D7) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D7{}

// End of datatype D7

// Definition of datatype D8
type D8 struct {
	Data_D8_
}

func (_this D8) Get_() Data_D8_ {
	return _this.Data_D8_
}

type Data_D8_ interface {
	isD8()
}

type CompanionStruct_D8_ struct {
}

var Companion_D8_ = CompanionStruct_D8_{}

type D8_DC21 struct {
	Cf36 bool
	Cf37 bool
	Cf38 _dafny.Int
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf36 bool, Cf37 bool, Cf38 _dafny.Int) D8 {
	return D8{D8_DC21{Cf36, Cf37, Cf38}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC20 struct {
	Cf35 _dafny.Map
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf35 _dafny.Map) D8 {
	return D8{D8_DC20{Cf35}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(false, false, _dafny.Zero)
}

func (_this D8) Dtor_cf36() bool {
	return _this.Get_().(D8_DC21).Cf36
}

func (_this D8) Dtor_cf37() bool {
	return _this.Get_().(D8_DC21).Cf37
}

func (_this D8) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D8_DC21).Cf38
}

func (_this D8) Dtor_cf35() _dafny.Map {
	return _this.Get_().(D8_DC20).Cf35
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf35) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC21:
		{
			data2, ok := other.Get_().(D8_DC21)
			return ok && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf35.Equals(data2.Cf35)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D8) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D8)
	return ok && _this.Equals(typed)
}

func Type_D8_() _dafny.TypeDescriptor {
	return type_D8_{}
}

type type_D8_ struct {
}

func (_this type_D8_) Default() interface{} {
	return Companion_D8_.Default()
}

func (_this type_D8_) String() string {
	return "main.D8"
}
func (_this D8) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D8{}

// End of datatype D8

// Definition of datatype D9
type D9 struct {
	Data_D9_
}

func (_this D9) Get_() Data_D9_ {
	return _this.Data_D9_
}

type Data_D9_ interface {
	isD9()
}

type CompanionStruct_D9_ struct {
}

var Companion_D9_ = CompanionStruct_D9_{}

type D9_DC22 struct {
	Cf39 _dafny.MultiSet
}

func (D9_DC22) isD9() {}

func (CompanionStruct_D9_) Create_DC22_(Cf39 _dafny.MultiSet) D9 {
	return D9{D9_DC22{Cf39}}
}

func (_this D9) Is_DC22() bool {
	_, ok := _this.Get_().(D9_DC22)
	return ok
}

func (CompanionStruct_D9_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D9) Dtor_cf39() _dafny.MultiSet {
	return _this.Get_().(D9_DC22).Cf39
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC22:
		{
			return "D9.DC22" + "(" + _dafny.String(data.Cf39) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC22:
		{
			data2, ok := other.Get_().(D9_DC22)
			return ok && data1.Cf39.Equals(data2.Cf39)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D9) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D9)
	return ok && _this.Equals(typed)
}

func Type_D9_() _dafny.TypeDescriptor {
	return type_D9_{}
}

type type_D9_ struct {
}

func (_this type_D9_) Default() interface{} {
	return Companion_D9_.Default()
}

func (_this type_D9_) String() string {
	return "main.D9"
}
func (_this D9) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D9{}

// End of datatype D9

// Definition of datatype D10
type D10 struct {
	Data_D10_
}

func (_this D10) Get_() Data_D10_ {
	return _this.Data_D10_
}

type Data_D10_ interface {
	isD10()
}

type CompanionStruct_D10_ struct {
}

var Companion_D10_ = CompanionStruct_D10_{}

type D10_DC23 struct {
	Cf40 T0
}

func (D10_DC23) isD10() {}

func (CompanionStruct_D10_) Create_DC23_(Cf40 T0) D10 {
	return D10{D10_DC23{Cf40}}
}

func (_this D10) Is_DC23() bool {
	_, ok := _this.Get_().(D10_DC23)
	return ok
}

func (CompanionStruct_D10_) Default() T0 {
	return (T0)(nil)
}

func (_this D10) Dtor_cf40() T0 {
	return _this.Get_().(D10_DC23).Cf40
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC23:
		{
			return "D10.DC23" + "(" + _dafny.String(data.Cf40) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC23:
		{
			data2, ok := other.Get_().(D10_DC23)
			return ok && _dafny.AreEqual(data1.Cf40, data2.Cf40)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D10) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D10)
	return ok && _this.Equals(typed)
}

func Type_D10_() _dafny.TypeDescriptor {
	return type_D10_{}
}

type type_D10_ struct {
}

func (_this type_D10_) Default() interface{} {
	return Companion_D10_.Default()
}

func (_this type_D10_) String() string {
	return "main.D10"
}
func (_this D10) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D10{}

// End of datatype D10

// Definition of datatype D11
type D11 struct {
	Data_D11_
}

func (_this D11) Get_() Data_D11_ {
	return _this.Data_D11_
}

type Data_D11_ interface {
	isD11()
}

type CompanionStruct_D11_ struct {
}

var Companion_D11_ = CompanionStruct_D11_{}

type D11_DC25 struct {
	Cf42 _dafny.Array
	Cf43 T1
	Cf44 _dafny.Int
}

func (D11_DC25) isD11() {}

func (CompanionStruct_D11_) Create_DC25_(Cf42 _dafny.Array, Cf43 T1, Cf44 _dafny.Int) D11 {
	return D11{D11_DC25{Cf42, Cf43, Cf44}}
}

func (_this D11) Is_DC25() bool {
	_, ok := _this.Get_().(D11_DC25)
	return ok
}

type D11_DC26 struct {
	Cf45 T0
	Cf46 _dafny.MultiSet
	Cf47 _dafny.Int
	Cf48 _dafny.Sequence
}

func (D11_DC26) isD11() {}

func (CompanionStruct_D11_) Create_DC26_(Cf45 T0, Cf46 _dafny.MultiSet, Cf47 _dafny.Int, Cf48 _dafny.Sequence) D11 {
	return D11{D11_DC26{Cf45, Cf46, Cf47, Cf48}}
}

func (_this D11) Is_DC26() bool {
	_, ok := _this.Get_().(D11_DC26)
	return ok
}

type D11_DC27 struct {
	Cf49 _dafny.Set
	Cf50 _dafny.Int
}

func (D11_DC27) isD11() {}

func (CompanionStruct_D11_) Create_DC27_(Cf49 _dafny.Set, Cf50 _dafny.Int) D11 {
	return D11{D11_DC27{Cf49, Cf50}}
}

func (_this D11) Is_DC27() bool {
	_, ok := _this.Get_().(D11_DC27)
	return ok
}

type D11_DC24 struct {
	Cf41 _dafny.Map
}

func (D11_DC24) isD11() {}

func (CompanionStruct_D11_) Create_DC24_(Cf41 _dafny.Map) D11 {
	return D11{D11_DC24{Cf41}}
}

func (_this D11) Is_DC24() bool {
	_, ok := _this.Get_().(D11_DC24)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC25_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), (T1)(nil), _dafny.Zero)
}

func (_this D11) Dtor_cf42() _dafny.Array {
	return _this.Get_().(D11_DC25).Cf42
}

func (_this D11) Dtor_cf43() T1 {
	return _this.Get_().(D11_DC25).Cf43
}

func (_this D11) Dtor_cf44() _dafny.Int {
	return _this.Get_().(D11_DC25).Cf44
}

func (_this D11) Dtor_cf45() T0 {
	return _this.Get_().(D11_DC26).Cf45
}

func (_this D11) Dtor_cf46() _dafny.MultiSet {
	return _this.Get_().(D11_DC26).Cf46
}

func (_this D11) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D11_DC26).Cf47
}

func (_this D11) Dtor_cf48() _dafny.Sequence {
	return _this.Get_().(D11_DC26).Cf48
}

func (_this D11) Dtor_cf49() _dafny.Set {
	return _this.Get_().(D11_DC27).Cf49
}

func (_this D11) Dtor_cf50() _dafny.Int {
	return _this.Get_().(D11_DC27).Cf50
}

func (_this D11) Dtor_cf41() _dafny.Map {
	return _this.Get_().(D11_DC24).Cf41
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC25:
		{
			return "D11.DC25" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ", " + _dafny.String(data.Cf44) + ")"
		}
	case D11_DC26:
		{
			return "D11.DC26" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + data.Cf48.VerbatimString(true) + ")"
		}
	case D11_DC27:
		{
			return "D11.DC27" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D11_DC24:
		{
			return "D11.DC24" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC25:
		{
			data2, ok := other.Get_().(D11_DC25)
			return ok && data1.Cf42 == data2.Cf42 && _dafny.AreEqual(data1.Cf43, data2.Cf43) && data1.Cf44.Cmp(data2.Cf44) == 0
		}
	case D11_DC26:
		{
			data2, ok := other.Get_().(D11_DC26)
			return ok && _dafny.AreEqual(data1.Cf45, data2.Cf45) && data1.Cf46.Equals(data2.Cf46) && data1.Cf47.Cmp(data2.Cf47) == 0 && data1.Cf48.Equals(data2.Cf48)
		}
	case D11_DC27:
		{
			data2, ok := other.Get_().(D11_DC27)
			return ok && data1.Cf49.Equals(data2.Cf49) && data1.Cf50.Cmp(data2.Cf50) == 0
		}
	case D11_DC24:
		{
			data2, ok := other.Get_().(D11_DC24)
			return ok && data1.Cf41.Equals(data2.Cf41)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D11) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D11)
	return ok && _this.Equals(typed)
}

func Type_D11_() _dafny.TypeDescriptor {
	return type_D11_{}
}

type type_D11_ struct {
}

func (_this type_D11_) Default() interface{} {
	return Companion_D11_.Default()
}

func (_this type_D11_) String() string {
	return "main.D11"
}
func (_this D11) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D11{}

// End of datatype D11

// Definition of datatype D12
type D12 struct {
	Data_D12_
}

func (_this D12) Get_() Data_D12_ {
	return _this.Data_D12_
}

type Data_D12_ interface {
	isD12()
}

type CompanionStruct_D12_ struct {
}

var Companion_D12_ = CompanionStruct_D12_{}

type D12_DC29 struct {
	Cf52 _dafny.Sequence
	Cf53 bool
}

func (D12_DC29) isD12() {}

func (CompanionStruct_D12_) Create_DC29_(Cf52 _dafny.Sequence, Cf53 bool) D12 {
	return D12{D12_DC29{Cf52, Cf53}}
}

func (_this D12) Is_DC29() bool {
	_, ok := _this.Get_().(D12_DC29)
	return ok
}

type D12_DC28 struct {
	Cf51 _dafny.Map
}

func (D12_DC28) isD12() {}

func (CompanionStruct_D12_) Create_DC28_(Cf51 _dafny.Map) D12 {
	return D12{D12_DC28{Cf51}}
}

func (_this D12) Is_DC28() bool {
	_, ok := _this.Get_().(D12_DC28)
	return ok
}

type D12_DC30 struct {
	Cf54 D12
}

func (D12_DC30) isD12() {}

func (CompanionStruct_D12_) Create_DC30_(Cf54 D12) D12 {
	return D12{D12_DC30{Cf54}}
}

func (_this D12) Is_DC30() bool {
	_, ok := _this.Get_().(D12_DC30)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC29_(_dafny.EmptySeq, false)
}

func (_this D12) Dtor_cf52() _dafny.Sequence {
	return _this.Get_().(D12_DC29).Cf52
}

func (_this D12) Dtor_cf53() bool {
	return _this.Get_().(D12_DC29).Cf53
}

func (_this D12) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D12_DC28).Cf51
}

func (_this D12) Dtor_cf54() D12 {
	return _this.Get_().(D12_DC30).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC29:
		{
			return "D12.DC29" + "(" + data.Cf52.VerbatimString(true) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D12_DC28:
		{
			return "D12.DC28" + "(" + _dafny.String(data.Cf51) + ")"
		}
	case D12_DC30:
		{
			return "D12.DC30" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC29:
		{
			data2, ok := other.Get_().(D12_DC29)
			return ok && data1.Cf52.Equals(data2.Cf52) && data1.Cf53 == data2.Cf53
		}
	case D12_DC28:
		{
			data2, ok := other.Get_().(D12_DC28)
			return ok && data1.Cf51.Equals(data2.Cf51)
		}
	case D12_DC30:
		{
			data2, ok := other.Get_().(D12_DC30)
			return ok && data1.Cf54.Equals(data2.Cf54)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D12) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D12)
	return ok && _this.Equals(typed)
}

func Type_D12_() _dafny.TypeDescriptor {
	return type_D12_{}
}

type type_D12_ struct {
}

func (_this type_D12_) Default() interface{} {
	return Companion_D12_.Default()
}

func (_this type_D12_) String() string {
	return "main.D12"
}
func (_this D12) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D12{}

// End of datatype D12

// Definition of datatype D13
type D13 struct {
	Data_D13_
}

func (_this D13) Get_() Data_D13_ {
	return _this.Data_D13_
}

type Data_D13_ interface {
	isD13()
}

type CompanionStruct_D13_ struct {
}

var Companion_D13_ = CompanionStruct_D13_{}

type D13_DC32 struct {
	Cf56 _dafny.Int
	Cf57 bool
	Cf58 bool
	Cf59 bool
	Cf60 bool
}

func (D13_DC32) isD13() {}

func (CompanionStruct_D13_) Create_DC32_(Cf56 _dafny.Int, Cf57 bool, Cf58 bool, Cf59 bool, Cf60 bool) D13 {
	return D13{D13_DC32{Cf56, Cf57, Cf58, Cf59, Cf60}}
}

func (_this D13) Is_DC32() bool {
	_, ok := _this.Get_().(D13_DC32)
	return ok
}

type D13_DC33 struct {
	Cf61 bool
	Cf62 _dafny.Int
	Cf63 _dafny.Int
	Cf64 bool
	Cf65 *C3
}

func (D13_DC33) isD13() {}

func (CompanionStruct_D13_) Create_DC33_(Cf61 bool, Cf62 _dafny.Int, Cf63 _dafny.Int, Cf64 bool, Cf65 *C3) D13 {
	return D13{D13_DC33{Cf61, Cf62, Cf63, Cf64, Cf65}}
}

func (_this D13) Is_DC33() bool {
	_, ok := _this.Get_().(D13_DC33)
	return ok
}

type D13_DC31 struct {
	Cf55 _dafny.Map
}

func (D13_DC31) isD13() {}

func (CompanionStruct_D13_) Create_DC31_(Cf55 _dafny.Map) D13 {
	return D13{D13_DC31{Cf55}}
}

func (_this D13) Is_DC31() bool {
	_, ok := _this.Get_().(D13_DC31)
	return ok
}

type D13_DC34 struct {
	Cf66 D13
}

func (D13_DC34) isD13() {}

func (CompanionStruct_D13_) Create_DC34_(Cf66 D13) D13 {
	return D13{D13_DC34{Cf66}}
}

func (_this D13) Is_DC34() bool {
	_, ok := _this.Get_().(D13_DC34)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC32_(_dafny.Zero, false, false, false, false)
}

func (_this D13) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D13_DC32).Cf56
}

func (_this D13) Dtor_cf57() bool {
	return _this.Get_().(D13_DC32).Cf57
}

func (_this D13) Dtor_cf58() bool {
	return _this.Get_().(D13_DC32).Cf58
}

func (_this D13) Dtor_cf59() bool {
	return _this.Get_().(D13_DC32).Cf59
}

func (_this D13) Dtor_cf60() bool {
	return _this.Get_().(D13_DC32).Cf60
}

func (_this D13) Dtor_cf61() bool {
	return _this.Get_().(D13_DC33).Cf61
}

func (_this D13) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D13_DC33).Cf62
}

func (_this D13) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D13_DC33).Cf63
}

func (_this D13) Dtor_cf64() bool {
	return _this.Get_().(D13_DC33).Cf64
}

func (_this D13) Dtor_cf65() *C3 {
	return _this.Get_().(D13_DC33).Cf65
}

func (_this D13) Dtor_cf55() _dafny.Map {
	return _this.Get_().(D13_DC31).Cf55
}

func (_this D13) Dtor_cf66() D13 {
	return _this.Get_().(D13_DC34).Cf66
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC32:
		{
			return "D13.DC32" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ", " + _dafny.String(data.Cf60) + ")"
		}
	case D13_DC33:
		{
			return "D13.DC33" + "(" + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ", " + _dafny.String(data.Cf65) + ")"
		}
	case D13_DC31:
		{
			return "D13.DC31" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D13_DC34:
		{
			return "D13.DC34" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC32:
		{
			data2, ok := other.Get_().(D13_DC32)
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0 && data1.Cf57 == data2.Cf57 && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59 && data1.Cf60 == data2.Cf60
		}
	case D13_DC33:
		{
			data2, ok := other.Get_().(D13_DC33)
			return ok && data1.Cf61 == data2.Cf61 && data1.Cf62.Cmp(data2.Cf62) == 0 && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64 == data2.Cf64 && data1.Cf65 == data2.Cf65
		}
	case D13_DC31:
		{
			data2, ok := other.Get_().(D13_DC31)
			return ok && data1.Cf55.Equals(data2.Cf55)
		}
	case D13_DC34:
		{
			data2, ok := other.Get_().(D13_DC34)
			return ok && data1.Cf66.Equals(data2.Cf66)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D13) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D13)
	return ok && _this.Equals(typed)
}

func Type_D13_() _dafny.TypeDescriptor {
	return type_D13_{}
}

type type_D13_ struct {
}

func (_this type_D13_) Default() interface{} {
	return Companion_D13_.Default()
}

func (_this type_D13_) String() string {
	return "main.D13"
}
func (_this D13) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D13{}

// End of datatype D13

// Definition of datatype D14
type D14 struct {
	Data_D14_
}

func (_this D14) Get_() Data_D14_ {
	return _this.Data_D14_
}

type Data_D14_ interface {
	isD14()
}

type CompanionStruct_D14_ struct {
}

var Companion_D14_ = CompanionStruct_D14_{}

type D14_DC36 struct {
	Cf68 bool
	Cf69 _dafny.Int
	Cf70 _dafny.Int
	Cf71 _dafny.Map
}

func (D14_DC36) isD14() {}

func (CompanionStruct_D14_) Create_DC36_(Cf68 bool, Cf69 _dafny.Int, Cf70 _dafny.Int, Cf71 _dafny.Map) D14 {
	return D14{D14_DC36{Cf68, Cf69, Cf70, Cf71}}
}

func (_this D14) Is_DC36() bool {
	_, ok := _this.Get_().(D14_DC36)
	return ok
}

type D14_DC35 struct {
	Cf67 _dafny.Array
}

func (D14_DC35) isD14() {}

func (CompanionStruct_D14_) Create_DC35_(Cf67 _dafny.Array) D14 {
	return D14{D14_DC35{Cf67}}
}

func (_this D14) Is_DC35() bool {
	_, ok := _this.Get_().(D14_DC35)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC36_(false, _dafny.Zero, _dafny.Zero, _dafny.EmptyMap)
}

func (_this D14) Dtor_cf68() bool {
	return _this.Get_().(D14_DC36).Cf68
}

func (_this D14) Dtor_cf69() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf69
}

func (_this D14) Dtor_cf70() _dafny.Int {
	return _this.Get_().(D14_DC36).Cf70
}

func (_this D14) Dtor_cf71() _dafny.Map {
	return _this.Get_().(D14_DC36).Cf71
}

func (_this D14) Dtor_cf67() _dafny.Array {
	return _this.Get_().(D14_DC35).Cf67
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC36:
		{
			return "D14.DC36" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D14_DC35:
		{
			return "D14.DC35" + "(" + _dafny.String(data.Cf67) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC36:
		{
			data2, ok := other.Get_().(D14_DC36)
			return ok && data1.Cf68 == data2.Cf68 && data1.Cf69.Cmp(data2.Cf69) == 0 && data1.Cf70.Cmp(data2.Cf70) == 0 && data1.Cf71.Equals(data2.Cf71)
		}
	case D14_DC35:
		{
			data2, ok := other.Get_().(D14_DC35)
			return ok && data1.Cf67 == data2.Cf67
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D14) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D14)
	return ok && _this.Equals(typed)
}

func Type_D14_() _dafny.TypeDescriptor {
	return type_D14_{}
}

type type_D14_ struct {
}

func (_this type_D14_) Default() interface{} {
	return Companion_D14_.Default()
}

func (_this type_D14_) String() string {
	return "main.D14"
}
func (_this D14) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D14{}

// End of datatype D14

// Definition of datatype D15
type D15 struct {
	Data_D15_
}

func (_this D15) Get_() Data_D15_ {
	return _this.Data_D15_
}

type Data_D15_ interface {
	isD15()
}

type CompanionStruct_D15_ struct {
}

var Companion_D15_ = CompanionStruct_D15_{}

type D15_DC37 struct {
	Cf72 *C0
}

func (D15_DC37) isD15() {}

func (CompanionStruct_D15_) Create_DC37_(Cf72 *C0) D15 {
	return D15{D15_DC37{Cf72}}
}

func (_this D15) Is_DC37() bool {
	_, ok := _this.Get_().(D15_DC37)
	return ok
}

func (CompanionStruct_D15_) Default() *C0 {
	return (*C0)(nil)
}

func (_this D15) Dtor_cf72() *C0 {
	return _this.Get_().(D15_DC37).Cf72
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC37:
		{
			return "D15.DC37" + "(" + _dafny.String(data.Cf72) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC37:
		{
			data2, ok := other.Get_().(D15_DC37)
			return ok && data1.Cf72 == data2.Cf72
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D15) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D15)
	return ok && _this.Equals(typed)
}

func Type_D15_() _dafny.TypeDescriptor {
	return type_D15_{}
}

type type_D15_ struct {
}

func (_this type_D15_) Default() interface{} {
	return Companion_D15_.Default()
}

func (_this type_D15_) String() string {
	return "main.D15"
}
func (_this D15) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D15{}

// End of datatype D15

// Definition of datatype D16
type D16 struct {
	Data_D16_
}

func (_this D16) Get_() Data_D16_ {
	return _this.Data_D16_
}

type Data_D16_ interface {
	isD16()
}

type CompanionStruct_D16_ struct {
}

var Companion_D16_ = CompanionStruct_D16_{}

type D16_DC39 struct {
	Cf74 _dafny.Int
	Cf75 _dafny.Map
}

func (D16_DC39) isD16() {}

func (CompanionStruct_D16_) Create_DC39_(Cf74 _dafny.Int, Cf75 _dafny.Map) D16 {
	return D16{D16_DC39{Cf74, Cf75}}
}

func (_this D16) Is_DC39() bool {
	_, ok := _this.Get_().(D16_DC39)
	return ok
}

type D16_DC38 struct {
	Cf73 _dafny.MultiSet
}

func (D16_DC38) isD16() {}

func (CompanionStruct_D16_) Create_DC38_(Cf73 _dafny.MultiSet) D16 {
	return D16{D16_DC38{Cf73}}
}

func (_this D16) Is_DC38() bool {
	_, ok := _this.Get_().(D16_DC38)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC39_(_dafny.Zero, _dafny.EmptyMap)
}

func (_this D16) Dtor_cf74() _dafny.Int {
	return _this.Get_().(D16_DC39).Cf74
}

func (_this D16) Dtor_cf75() _dafny.Map {
	return _this.Get_().(D16_DC39).Cf75
}

func (_this D16) Dtor_cf73() _dafny.MultiSet {
	return _this.Get_().(D16_DC38).Cf73
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC39:
		{
			return "D16.DC39" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ")"
		}
	case D16_DC38:
		{
			return "D16.DC38" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC39:
		{
			data2, ok := other.Get_().(D16_DC39)
			return ok && data1.Cf74.Cmp(data2.Cf74) == 0 && data1.Cf75.Equals(data2.Cf75)
		}
	case D16_DC38:
		{
			data2, ok := other.Get_().(D16_DC38)
			return ok && data1.Cf73.Equals(data2.Cf73)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D16) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D16)
	return ok && _this.Equals(typed)
}

func Type_D16_() _dafny.TypeDescriptor {
	return type_D16_{}
}

type type_D16_ struct {
}

func (_this type_D16_) Default() interface{} {
	return Companion_D16_.Default()
}

func (_this type_D16_) String() string {
	return "main.D16"
}
func (_this D16) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D16{}

// End of datatype D16

// Definition of datatype D17
type D17 struct {
	Data_D17_
}

func (_this D17) Get_() Data_D17_ {
	return _this.Data_D17_
}

type Data_D17_ interface {
	isD17()
}

type CompanionStruct_D17_ struct {
}

var Companion_D17_ = CompanionStruct_D17_{}

type D17_DC41 struct {
	Cf77 _dafny.Int
	Cf78 bool
}

func (D17_DC41) isD17() {}

func (CompanionStruct_D17_) Create_DC41_(Cf77 _dafny.Int, Cf78 bool) D17 {
	return D17{D17_DC41{Cf77, Cf78}}
}

func (_this D17) Is_DC41() bool {
	_, ok := _this.Get_().(D17_DC41)
	return ok
}

type D17_DC40 struct {
	Cf76 _dafny.Map
}

func (D17_DC40) isD17() {}

func (CompanionStruct_D17_) Create_DC40_(Cf76 _dafny.Map) D17 {
	return D17{D17_DC40{Cf76}}
}

func (_this D17) Is_DC40() bool {
	_, ok := _this.Get_().(D17_DC40)
	return ok
}

type D17_DC42 struct {
	Cf79 D17
}

func (D17_DC42) isD17() {}

func (CompanionStruct_D17_) Create_DC42_(Cf79 D17) D17 {
	return D17{D17_DC42{Cf79}}
}

func (_this D17) Is_DC42() bool {
	_, ok := _this.Get_().(D17_DC42)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC41_(_dafny.Zero, false)
}

func (_this D17) Dtor_cf77() _dafny.Int {
	return _this.Get_().(D17_DC41).Cf77
}

func (_this D17) Dtor_cf78() bool {
	return _this.Get_().(D17_DC41).Cf78
}

func (_this D17) Dtor_cf76() _dafny.Map {
	return _this.Get_().(D17_DC40).Cf76
}

func (_this D17) Dtor_cf79() D17 {
	return _this.Get_().(D17_DC42).Cf79
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC41:
		{
			return "D17.DC41" + "(" + _dafny.String(data.Cf77) + ", " + _dafny.String(data.Cf78) + ")"
		}
	case D17_DC40:
		{
			return "D17.DC40" + "(" + _dafny.String(data.Cf76) + ")"
		}
	case D17_DC42:
		{
			return "D17.DC42" + "(" + _dafny.String(data.Cf79) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC41:
		{
			data2, ok := other.Get_().(D17_DC41)
			return ok && data1.Cf77.Cmp(data2.Cf77) == 0 && data1.Cf78 == data2.Cf78
		}
	case D17_DC40:
		{
			data2, ok := other.Get_().(D17_DC40)
			return ok && data1.Cf76.Equals(data2.Cf76)
		}
	case D17_DC42:
		{
			data2, ok := other.Get_().(D17_DC42)
			return ok && data1.Cf79.Equals(data2.Cf79)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D17) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D17)
	return ok && _this.Equals(typed)
}

func Type_D17_() _dafny.TypeDescriptor {
	return type_D17_{}
}

type type_D17_ struct {
}

func (_this type_D17_) Default() interface{} {
	return Companion_D17_.Default()
}

func (_this type_D17_) String() string {
	return "main.D17"
}
func (_this D17) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D17{}

// End of datatype D17

// Definition of datatype D18
type D18 struct {
	Data_D18_
}

func (_this D18) Get_() Data_D18_ {
	return _this.Data_D18_
}

type Data_D18_ interface {
	isD18()
}

type CompanionStruct_D18_ struct {
}

var Companion_D18_ = CompanionStruct_D18_{}

type D18_DC44 struct {
	Cf81 _dafny.Set
	Cf82 _dafny.Int
	Cf83 _dafny.Int
	Cf84 bool
	Cf85 bool
}

func (D18_DC44) isD18() {}

func (CompanionStruct_D18_) Create_DC44_(Cf81 _dafny.Set, Cf82 _dafny.Int, Cf83 _dafny.Int, Cf84 bool, Cf85 bool) D18 {
	return D18{D18_DC44{Cf81, Cf82, Cf83, Cf84, Cf85}}
}

func (_this D18) Is_DC44() bool {
	_, ok := _this.Get_().(D18_DC44)
	return ok
}

type D18_DC43 struct {
	Cf80 *C1
}

func (D18_DC43) isD18() {}

func (CompanionStruct_D18_) Create_DC43_(Cf80 *C1) D18 {
	return D18{D18_DC43{Cf80}}
}

func (_this D18) Is_DC43() bool {
	_, ok := _this.Get_().(D18_DC43)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC44_(_dafny.EmptySet, _dafny.Zero, _dafny.Zero, false, false)
}

func (_this D18) Dtor_cf81() _dafny.Set {
	return _this.Get_().(D18_DC44).Cf81
}

func (_this D18) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D18_DC44).Cf82
}

func (_this D18) Dtor_cf83() _dafny.Int {
	return _this.Get_().(D18_DC44).Cf83
}

func (_this D18) Dtor_cf84() bool {
	return _this.Get_().(D18_DC44).Cf84
}

func (_this D18) Dtor_cf85() bool {
	return _this.Get_().(D18_DC44).Cf85
}

func (_this D18) Dtor_cf80() *C1 {
	return _this.Get_().(D18_DC43).Cf80
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC44:
		{
			return "D18.DC44" + "(" + _dafny.String(data.Cf81) + ", " + _dafny.String(data.Cf82) + ", " + _dafny.String(data.Cf83) + ", " + _dafny.String(data.Cf84) + ", " + _dafny.String(data.Cf85) + ")"
		}
	case D18_DC43:
		{
			return "D18.DC43" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC44:
		{
			data2, ok := other.Get_().(D18_DC44)
			return ok && data1.Cf81.Equals(data2.Cf81) && data1.Cf82.Cmp(data2.Cf82) == 0 && data1.Cf83.Cmp(data2.Cf83) == 0 && data1.Cf84 == data2.Cf84 && data1.Cf85 == data2.Cf85
		}
	case D18_DC43:
		{
			data2, ok := other.Get_().(D18_DC43)
			return ok && data1.Cf80 == data2.Cf80
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D18) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D18)
	return ok && _this.Equals(typed)
}

func Type_D18_() _dafny.TypeDescriptor {
	return type_D18_{}
}

type type_D18_ struct {
}

func (_this type_D18_) Default() interface{} {
	return Companion_D18_.Default()
}

func (_this type_D18_) String() string {
	return "main.D18"
}
func (_this D18) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D18{}

// End of datatype D18

// Definition of datatype D19
type D19 struct {
	Data_D19_
}

func (_this D19) Get_() Data_D19_ {
	return _this.Data_D19_
}

type Data_D19_ interface {
	isD19()
}

type CompanionStruct_D19_ struct {
}

var Companion_D19_ = CompanionStruct_D19_{}

type D19_DC46 struct {
	Cf87 _dafny.Sequence
	Cf88 _dafny.Int
	Cf89 _dafny.Int
	Cf90 _dafny.Int
	Cf91 _dafny.Set
}

func (D19_DC46) isD19() {}

func (CompanionStruct_D19_) Create_DC46_(Cf87 _dafny.Sequence, Cf88 _dafny.Int, Cf89 _dafny.Int, Cf90 _dafny.Int, Cf91 _dafny.Set) D19 {
	return D19{D19_DC46{Cf87, Cf88, Cf89, Cf90, Cf91}}
}

func (_this D19) Is_DC46() bool {
	_, ok := _this.Get_().(D19_DC46)
	return ok
}

type D19_DC45 struct {
	Cf86 _dafny.Array
}

func (D19_DC45) isD19() {}

func (CompanionStruct_D19_) Create_DC45_(Cf86 _dafny.Array) D19 {
	return D19{D19_DC45{Cf86}}
}

func (_this D19) Is_DC45() bool {
	_, ok := _this.Get_().(D19_DC45)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC46_(_dafny.EmptySeq, _dafny.Zero, _dafny.Zero, _dafny.Zero, _dafny.EmptySet)
}

func (_this D19) Dtor_cf87() _dafny.Sequence {
	return _this.Get_().(D19_DC46).Cf87
}

func (_this D19) Dtor_cf88() _dafny.Int {
	return _this.Get_().(D19_DC46).Cf88
}

func (_this D19) Dtor_cf89() _dafny.Int {
	return _this.Get_().(D19_DC46).Cf89
}

func (_this D19) Dtor_cf90() _dafny.Int {
	return _this.Get_().(D19_DC46).Cf90
}

func (_this D19) Dtor_cf91() _dafny.Set {
	return _this.Get_().(D19_DC46).Cf91
}

func (_this D19) Dtor_cf86() _dafny.Array {
	return _this.Get_().(D19_DC45).Cf86
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC46:
		{
			return "D19.DC46" + "(" + data.Cf87.VerbatimString(true) + ", " + _dafny.String(data.Cf88) + ", " + _dafny.String(data.Cf89) + ", " + _dafny.String(data.Cf90) + ", " + _dafny.String(data.Cf91) + ")"
		}
	case D19_DC45:
		{
			return "D19.DC45" + "(" + _dafny.String(data.Cf86) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC46:
		{
			data2, ok := other.Get_().(D19_DC46)
			return ok && data1.Cf87.Equals(data2.Cf87) && data1.Cf88.Cmp(data2.Cf88) == 0 && data1.Cf89.Cmp(data2.Cf89) == 0 && data1.Cf90.Cmp(data2.Cf90) == 0 && data1.Cf91.Equals(data2.Cf91)
		}
	case D19_DC45:
		{
			data2, ok := other.Get_().(D19_DC45)
			return ok && data1.Cf86 == data2.Cf86
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D19) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D19)
	return ok && _this.Equals(typed)
}

func Type_D19_() _dafny.TypeDescriptor {
	return type_D19_{}
}

type type_D19_ struct {
}

func (_this type_D19_) Default() interface{} {
	return Companion_D19_.Default()
}

func (_this type_D19_) String() string {
	return "main.D19"
}
func (_this D19) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D19{}

// End of datatype D19

// Definition of trait T0
type T0 interface {
	String() string
	M0(globalState *GlobalState) (_dafny.Int, bool, bool)
	F6() _dafny.Int
}
type CompanionStruct_T0_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T0_ = CompanionStruct_T0_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T0_) CastTo_(x interface{}) T0 {
	var t T0
	t, _ = x.(T0)
	return t
}

// End of trait T0

// Definition of trait T1
type T1 interface {
	String() string
	Fm16(p0 _dafny.MultiSet, globalState *GlobalState) bool
	Fm17(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int
}
type CompanionStruct_T1_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T1_ = CompanionStruct_T1_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T1_) CastTo_(x interface{}) T1 {
	var t T1
	t, _ = x.(T1)
	return t
}

// End of trait T1

// Definition of class GlobalState
type GlobalState struct {
	F0  _dafny.Int
	F1  bool
	F2  _dafny.Int
	F5  bool
	_f3 bool
	_f4 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F1 = false
	_this.F2 = _dafny.Zero
	_this.F5 = false
	_this._f3 = false
	_this._f4 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_GlobalState_ struct {
}

var Companion_GlobalState_ = CompanionStruct_GlobalState_{}

func (_this *GlobalState) Equals(other *GlobalState) bool {
	return _this == other
}

func (_this *GlobalState) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*GlobalState)
	return ok && _this.Equals(other)
}

func (*GlobalState) String() string {
	return "_module.GlobalState"
}

func Type_GlobalState_() _dafny.TypeDescriptor {
	return type_GlobalState_{}
}

type type_GlobalState_ struct {
}

func (_this type_GlobalState_) Default() interface{} {
	return (*GlobalState)(nil)
}

func (_this type_GlobalState_) String() string {
	return "main.GlobalState"
}
func (_this *GlobalState) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &GlobalState{}

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 bool, f2 _dafny.Int, f3 bool, f4 _dafny.Sequence, f5 bool) {
	{
		(_this).F0 = f0
		(_this).F1 = f1
		(_this).F2 = f2
		(_this)._f3 = f3
		(_this)._f4 = f4
		(_this).F5 = f5
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F4() _dafny.Sequence {
	{
		return _this._f4
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f13 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f13 = false
	return &_this
}

type CompanionStruct_C0_ struct {
}

var Companion_C0_ = CompanionStruct_C0_{}

func (_this *C0) Equals(other *C0) bool {
	return _this == other
}

func (_this *C0) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C0)
	return ok && _this.Equals(other)
}

func (*C0) String() string {
	return "_module.C0"
}

func Type_C0_() _dafny.TypeDescriptor {
	return type_C0_{}
}

type type_C0_ struct {
}

func (_this type_C0_) Default() interface{} {
	return (*C0)(nil)
}

func (_this type_C0_) String() string {
	return "main.C0"
}
func (_this *C0) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__(f13 bool) {
	{
		(_this)._f13 = f13
	}
}
func (_this *C0) Fm23(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(-986), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("sicrkla"), _dafny.UnicodeSeqOfUtf8Bytes("nyetvlcd"))).Cardinality()))
	}
}
func (_this *C0) F13() bool {
	{
		return _this._f13
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	dummy byte
}

func New_C1_() *C1 {
	_this := C1{}

	return &_this
}

type CompanionStruct_C1_ struct {
}

var Companion_C1_ = CompanionStruct_C1_{}

func (_this *C1) Equals(other *C1) bool {
	return _this == other
}

func (_this *C1) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C1)
	return ok && _this.Equals(other)
}

func (*C1) String() string {
	return "_module.C1"
}

func Type_C1_() _dafny.TypeDescriptor {
	return type_C1_{}
}

type type_C1_ struct {
}

func (_this type_C1_) Default() interface{} {
	return (*C1)(nil)
}

func (_this type_C1_) String() string {
	return "main.C1"
}
func (_this *C1) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T1_.TraitID_}
}

var _ T1 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__() {
	{
	}
}
func (_this *C1) Fm16(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return !(!(_dafny.Companion_Sequence_.Contains(_dafny.UnicodeSeqOfUtf8Bytes("lfuq"), _dafny.CodePoint('d'))) || (_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.CodePoint('k'), _dafny.CodePoint('g'), _dafny.CodePoint('x'), _dafny.CodePoint('j'), _dafny.CodePoint('l')), _dafny.SeqOf(_dafny.CodePoint('j'), _dafny.CodePoint('i')))))
	}
}
func (_this *C1) Fm17(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return ((_dafny.Zero).Minus((func() _dafny.Int {
			if true {
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("igmm")).Cardinality())
			}
			return _dafny.IntOfInt64(913)
		})())).Plus(_dafny.IntOfInt64(828))
	}
}
func (_this *C1) Fm20(p0 D5, p1 _dafny.Int, p2 D1, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(699))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg17 _dafny.Int) interface{} {
				return coer17(arg17)
			}
		}(func(_376_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		})), _dafny.UnicodeSeqOfUtf8Bytes("qodmkq"))
	}
}
func (_this *C1) M6(globalState *GlobalState) (_dafny.Int, _dafny.Int, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _377_v0 _dafny.Int
		_ = _377_v0
		_377_v0 = _dafny.IntOfInt64(477)
		(globalState).F2 = _377_v0
		var _hi3 _dafny.Int = _377_v0
		_ = _hi3
		for _378_i0 := _377_v0; _378_i0.Cmp(_hi3) < 0; _378_i0 = _378_i0.Plus(_dafny.One) {
			var _379_v1 _dafny.Map
			_ = _379_v1
			_379_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(65), _378_i0)
			var _380_v2 bool
			_ = _380_v2
			_380_v2 = true
			var _381_v4 _dafny.Sequence
			_ = _381_v4
			_381_v4 = _dafny.SeqOf(_dafny.SetOf(_377_v0))
			var _382_v5 _dafny.Sequence
			_ = _382_v5
			_382_v5 = _dafny.UnicodeSeqOfUtf8Bytes("a")
			var _383_v6 _dafny.CodePoint
			_ = _383_v6
			_383_v6 = _dafny.CodePoint('t')
			var _384_v7 _dafny.Set
			_ = _384_v7
			_384_v7 = _dafny.SetOf((_this).Fm17(_dafny.IntOfInt64(-88), _dafny.SetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_382_v5, (Companion_Default___.SafeIndex(_377_v0, _dafny.IntOfUint32((_382_v5).Cardinality()))).Uint32(), _383_v6)).Cardinality())), globalState))
			var _385_v9 _dafny.Map
			_ = _385_v9
			_385_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(17), _380_v2)
			var _386_v10 _dafny.MultiSet
			_ = _386_v10
			_386_v10 = _dafny.MultiSetOf((_385_v9).Cardinality(), _377_v0)
			var _387_v12 _dafny.Array
			_ = _387_v12
			var _nwElement0_12 _dafny.Map = _379_v1
			_ = _nwElement0_12
			var _nw67 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(14))
			_ = _nw67
			(_nw67).ArraySet1(_nwElement0_12, 0)
			(_nw67).ArraySet1((func() _dafny.Map {
				if _380_v2 {
					return _379_v1
				}
				return func() _dafny.Map {
					var _coll22 = _dafny.NewMapBuilder()
					_ = _coll22
					for _iter23 := _dafny.Iterate(((_381_v4).Select((Companion_Default___.SafeIndex(_377_v0, _dafny.IntOfUint32((_381_v4).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
						_compr_22, _ok23 := _iter23()
						if !_ok23 {
							break
						}
						var _388_v3 _dafny.Int
						_388_v3 = interface{}(_compr_22).(_dafny.Int)
						if ((_381_v4).Select((Companion_Default___.SafeIndex(_377_v0, _dafny.IntOfUint32((_381_v4).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_388_v3) {
							_coll22.Add(Companion_Default___.SafeModuloInt(_388_v3, _378_i0), _378_i0)
						}
					}
					return _coll22.ToMap()
				}()
			})(), 1)
			(_nw67).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _377_v0), 2)
			(_nw67).ArraySet1(_379_v1, 3)
			(_nw67).ArraySet1((_379_v1).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, (_384_v7).Cardinality())), 4)
			(_nw67).ArraySet1(_379_v1, 5)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll23 = _dafny.NewMapBuilder()
				_ = _coll23
				for _iter24 := _dafny.Iterate((_386_v10).Elements()); ; {
					_compr_23, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _389_v8 _dafny.Int
					_389_v8 = interface{}(_compr_23).(_dafny.Int)
					if (_386_v10).Contains(_389_v8) {
						_coll23.Add((_389_v8).Times(_378_i0), _dafny.IntOfUint32((_dafny.SeqOf(_378_i0)).Cardinality()))
					}
				}
				return _coll23.ToMap()
			}(), 6)
			(_nw67).ArraySet1(_379_v1, 7)
			(_nw67).ArraySet1((_379_v1).Update(_377_v0, _377_v0), 8)
			(_nw67).ArraySet1(_379_v1, 9)
			(_nw67).ArraySet1(_379_v1, 10)
			(_nw67).ArraySet1(_379_v1, 11)
			(_nw67).ArraySet1(_379_v1, 12)
			(_nw67).ArraySet1(func() _dafny.Map {
				var _coll24 = _dafny.NewMapBuilder()
				_ = _coll24
				for _iter25 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(318), _dafny.IntOfInt64(205))); ; {
					_compr_24, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _390_v11 _dafny.Int
					_390_v11 = interface{}(_compr_24).(_dafny.Int)
					if ((_dafny.IntOfInt64(318)).Cmp(_390_v11) <= 0) && ((_390_v11).Cmp(_dafny.IntOfInt64(205)) < 0) {
						_coll24.Add((_390_v11).Plus(_378_i0), _377_v0)
					}
				}
				return _coll24.ToMap()
			}(), 13)
			_387_v12 = _nw67
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_387_v12), 0))
			_ = _index58
			(_387_v12).ArraySet1(_379_v1, (_index58).Int())
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_387_v12), 0))
			_ = _index59
			(_387_v12).ArraySet1(_379_v1, (_index59).Int())
			var _391_v13 _dafny.Map
			_ = _391_v13
			_391_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(754))
			_391_v13 = (_391_v13).Update(!(_380_v2), _377_v0)
			var _392_v14 _dafny.Array
			_ = _392_v14
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
			_ = _nw68
			_392_v14 = _nw68
			var _393_v15 _dafny.Map
			_ = _393_v15
			_393_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_380_v2, _392_v14)
			var _394_v16 D0
			_ = _394_v16
			_394_v16 = Companion_D0_.Create_DC2_(_377_v0, _383_v6, (_393_v15).Cardinality(), _378_i0, (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm21(_382_v5, globalState)).Cardinality())))
			_383_v6 = (_394_v16).Dtor_cf6()
			_384_v7 = ((_384_v7).Difference(Companion_Default___.Fm22(globalState))).Difference(func() _dafny.Set {
				var _coll25 = _dafny.NewBuilder()
				_ = _coll25
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(873), _dafny.IntOfInt64(595))); ; {
					_compr_25, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _395_v17 _dafny.Int
					_395_v17 = interface{}(_compr_25).(_dafny.Int)
					if ((_dafny.IntOfInt64(873)).Cmp(_395_v17) <= 0) && ((_395_v17).Cmp(_dafny.IntOfInt64(595)) < 0) {
						_coll25.Add(Companion_Default___.SafeDivisionInt(_395_v17, _377_v0))
					}
				}
				return _coll25.ToSet()
			}())
		}
		var _396_v18 _dafny.Array
		_ = _396_v18
		var _nw69 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(25))
		_ = _nw69
		_396_v18 = _nw69
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
		_ = _nw70
		_396_v18 = _nw70
		var _397_i1 _dafny.Int
		_ = _397_i1
		_397_i1 = _dafny.Zero
		{
			for !(false) {
				{
					if (_397_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L0
					}
					_397_i1 = (_397_i1).Plus(_dafny.One)
					var _398_v19 bool
					_ = _398_v19
					_398_v19 = true
					var _399_v20 _dafny.MultiSet
					_ = _399_v20
					_399_v20 = _dafny.MultiSetOf(_398_v19, _398_v19)
					var _400_v21 _dafny.Sequence
					_ = _400_v21
					_400_v21 = _dafny.SeqOf(_398_v19, _398_v19, (_this).Fm16(_399_v20, globalState))
					var _401_v22 _dafny.CodePoint
					_ = _401_v22
					_401_v22 = _dafny.CodePoint('b')
					var _402_v23 _dafny.Map
					_ = _402_v23
					_402_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_401_v22, _401_v22)
					(globalState).F5 = _dafny.Companion_Sequence_.Equal(_400_v21, _dafny.Companion_Sequence_.Update(_400_v21, (Companion_Default___.SafeIndex((_402_v23).Cardinality(), _dafny.IntOfUint32((_400_v21).Cardinality()))).Uint32(), true))
					var _403_v24 _dafny.Sequence
					_ = _403_v24
					_403_v24 = _dafny.UnicodeSeqOfUtf8Bytes("whpm")
					_403_v24 = _403_v24
					var _404_v25 _dafny.MultiSet
					_ = _404_v25
					_404_v25 = _dafny.MultiSetOf(Companion_Default___.SafeModuloInt(_377_v0, _377_v0))
					_404_v25 = _404_v25
					var _405_v26 D0
					_ = _405_v26
					_405_v26 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(955), _dafny.CodePoint('y'), _377_v0, _377_v0, _377_v0)
					var _406_v27 _dafny.Map
					_ = _406_v27
					_406_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v0, _405_v26)
					var _407_v28 _dafny.Map
					_ = _407_v28
					_407_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_398_v19, (_406_v27).Cardinality())
					var _408_v29 _dafny.Array
					_ = _408_v29
					var _len0_12 _dafny.Int = _dafny.IntOfInt64(20)
					_ = _len0_12
					var _nw71 _dafny.Array
					_ = _nw71
					if _len0_12.Cmp(_dafny.Zero) == 0 {
						_nw71 = _dafny.NewArray(_len0_12)
					} else {
						var _init12 func(_dafny.Int) _dafny.Int = (func(_409_v0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_410_i2 _dafny.Int) _dafny.Int {
								return (_410_i2).Plus(_409_v0)
							}
						})(_377_v0)
						_ = _init12
						var _element0_12 = _init12(_dafny.Zero)
						_ = _element0_12
						_nw71 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
						(_nw71).ArraySet1(_element0_12, 0)
						var _nativeLen0_12 = (_len0_12).Int()
						_ = _nativeLen0_12
						for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
							(_nw71).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
						}
					}
					_408_v29 = _nw71
					var _411_v30 _dafny.Map
					_ = _411_v30
					_411_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_dafny.Zero).Minus((_407_v28).Cardinality())), _408_v29)
					var _412_v31 *C0
					_ = _412_v31
					var _nw72 *C0 = New_C0_()
					_ = _nw72
					_nw72.Ctor__((_411_v30).Equals(_411_v30))
					_412_v31 = _nw72
					goto C0
				}
			C0:
			}
			goto L0
		}
	L0:
		var _413_v32 bool
		_ = _413_v32
		_413_v32 = false
		var _414_v33 _dafny.Set
		_ = _414_v33
		_414_v33 = _dafny.SetOf(_413_v32)
		var _415_v34 D6
		_ = _415_v34
		_415_v34 = Companion_D6_.Create_DC15_(_414_v33)
		var _416_v35 *C0
		_ = _416_v35
		var _nw73 *C0 = New_C0_()
		_ = _nw73
		_nw73.Ctor__(((_415_v34).Dtor_cf26()).IsSubsetOf(_414_v33))
		_416_v35 = _nw73
		var _417_v36 _dafny.Sequence
		_ = _417_v36
		_417_v36 = _dafny.UnicodeSeqOfUtf8Bytes("le")
		var _418_v37 _dafny.Map
		_ = _418_v37
		_418_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_413_v32, _dafny.Companion_Sequence_.Concatenate(_417_v36, Companion_Default___.Fm2(_417_v36, globalState)))
		_418_v37 = (_418_v37).Update(!(_413_v32), _417_v36)
		r0 = _377_v0
		var _419_v38 _dafny.CodePoint
		_ = _419_v38
		_419_v38 = _dafny.CodePoint('q')
		var _420_v39 _dafny.Map
		_ = _420_v39
		_420_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(237), _dafny.SetOf(true))
		var _421_v40 D0
		_ = _421_v40
		_421_v40 = Companion_D0_.Create_DC0_(_420_v39)
		var _422_v41 _dafny.Sequence
		_ = _422_v41
		_422_v41 = _dafny.SeqOf(_413_v32, (_416_v35).F13())
		var _423_v43 _dafny.Array
		_ = _423_v43
		var _nwElement0_13 _dafny.Int = _377_v0
		_ = _nwElement0_13
		var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(19))
		_ = _nw74
		(_nw74).ArraySet1(_nwElement0_13, 0)
		(_nw74).ArraySet1(_377_v0, 1)
		(_nw74).ArraySet1((_this).Fm17(_dafny.IntOfUint32((_417_v36).Cardinality()), func() _dafny.Set {
			var _coll26 = _dafny.NewBuilder()
			_ = _coll26
			for _iter27 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(786))); ; {
				_compr_26, _ok27 := _iter27()
				if !_ok27 {
					break
				}
				var _424_v42 _dafny.Int
				_424_v42 = interface{}(_compr_26).(_dafny.Int)
				if ((_dafny.IntOfInt64(381)).Cmp(_424_v42) <= 0) && ((_424_v42).Cmp(_dafny.IntOfInt64(786)) < 0) {
					_coll26.Add((_424_v42).Plus(_dafny.IntOfInt64(330)))
				}
			}
			return _coll26.ToSet()
		}(), globalState), 2)
		(_nw74).ArraySet1(_377_v0, 3)
		(_nw74).ArraySet1(_377_v0, 4)
		(_nw74).ArraySet1(_377_v0, 5)
		(_nw74).ArraySet1(_377_v0, 6)
		(_nw74).ArraySet1(_377_v0, 7)
		(_nw74).ArraySet1(_377_v0, 8)
		(_nw74).ArraySet1(_377_v0, 9)
		(_nw74).ArraySet1(_377_v0, 10)
		(_nw74).ArraySet1(_377_v0, 11)
		(_nw74).ArraySet1(_377_v0, 12)
		(_nw74).ArraySet1(_377_v0, 13)
		(_nw74).ArraySet1(_377_v0, 14)
		(_nw74).ArraySet1((_414_v33).Cardinality(), 15)
		(_nw74).ArraySet1((_414_v33).Cardinality(), 16)
		(_nw74).ArraySet1(_377_v0, 17)
		(_nw74).ArraySet1(_dafny.IntOfUint32((_417_v36).Cardinality()), 18)
		_423_v43 = _nw74
		var _425_v44 _dafny.Sequence
		_ = _425_v44
		_425_v44 = _dafny.SeqOf(_423_v43, _423_v43, _423_v43, _423_v43)
		r1 = ((_dafny.Zero).Minus((func() _dafny.Int {
			if (_416_v35).F13() {
				return _dafny.IntOfUint32((_417_v36).Cardinality())
			}
			return _dafny.IntOfInt64(231)
		})())).Plus((_416_v35).Fm23(_377_v0, _377_v0, _419_v38, Companion_Default___.Fm1(_421_v40, (_416_v35).F13(), _377_v0, (_422_v41).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_425_v44).Cardinality()), _dafny.IntOfUint32((_422_v41).Cardinality()))).Uint32()).(bool), globalState), globalState))
		r2 = !((_416_v35).F13()) || ((_422_v41).Select((Companion_Default___.SafeIndex(_377_v0, _dafny.IntOfUint32((_422_v41).Cardinality()))).Uint32()).(bool))
		return r0, r1, r2
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F12  bool
	_f11 _dafny.Sequence
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F12 = false
	_this._f11 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C2_ struct {
}

var Companion_C2_ = CompanionStruct_C2_{}

func (_this *C2) Equals(other *C2) bool {
	return _this == other
}

func (_this *C2) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C2)
	return ok && _this.Equals(other)
}

func (*C2) String() string {
	return "_module.C2"
}

func Type_C2_() _dafny.TypeDescriptor {
	return type_C2_{}
}

type type_C2_ struct {
}

func (_this type_C2_) Default() interface{} {
	return (*C2)(nil)
}

func (_this type_C2_) String() string {
	return "main.C2"
}
func (_this *C2) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) Ctor__(f11 _dafny.Sequence, f12 bool) {
	{
		(_this)._f11 = f11
		(_this).F12 = f12
	}
}
func (_this *C2) M5(p0 bool, p1 bool, p2 _dafny.Array, p3 D0, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		if (p1) && ((func() bool {
			if false {
				return !(p1)
			}
			return p1
		})()) {
			var _426_v0 _dafny.Int
			_ = _426_v0
			_426_v0 = _dafny.IntOfInt64(-405)
			var _427_v1 _dafny.MultiSet
			_ = _427_v1
			_427_v1 = _dafny.MultiSetOf((_dafny.IntOfInt64(803)).Minus(_426_v0))
			var _428_v2 _dafny.Sequence
			_ = _428_v2
			_428_v2 = _dafny.SeqOf(_426_v0, _dafny.IntOfInt64(432))
			var _429_v3 _dafny.Sequence
			_ = _429_v3
			_429_v3 = _dafny.SeqOf(_427_v1)
			_427_v1 = ((_dafny.MultiSetFromSeq(_428_v2)).Intersection(_427_v1)).Union((_429_v3).Select((Companion_Default___.SafeIndex(_426_v0, _dafny.IntOfUint32((_429_v3).Cardinality()))).Uint32()).(_dafny.MultiSet))
			var _430_v4 _dafny.MultiSet
			_ = _430_v4
			_430_v4 = _dafny.MultiSetOf(_this.F12, p0)
			var _431_v5 _dafny.Set
			_ = _431_v5
			_431_v5 = _dafny.SetOf(_430_v4, _430_v4)
			if ((_431_v5).Cardinality()).Cmp(_426_v0) >= 0 {
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((p2), 0))
				_ = _index60
				(p2).ArraySet1(p1, (_index60).Int())
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((p2), 0))
				_ = _index61
				var _rhs47 bool = p0
				_ = _rhs47
				var _rhs48 bool = p0
				_ = _rhs48
				var _rhs49 bool = (func() bool {
					if (func() bool {
						if _this.F12 {
							return p1
						}
						return _this.F12
					})() {
						return true
					}
					return true
				})()
				_ = _rhs49
				var _lhs40 *C2 = _this
				_ = _lhs40
				var _lhs41 _dafny.Array = p2
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((p2), 0))
				_ = _lhs42
				var _lhs43 *C2 = _this
				_ = _lhs43
				_lhs40.F12 = _rhs47
				(_lhs41).ArraySet1(_rhs48, (_lhs42).Int())
				_lhs43.F12 = _rhs49
				var _432_v6 _dafny.Set
				_ = _432_v6
				_432_v6 = _dafny.SetOf(_426_v0, (_427_v1).Cardinality())
				(globalState).F2 = ((_432_v6).Union((_432_v6).Intersection(_432_v6))).Cardinality()
				(globalState).F0 = _dafny.IntOfInt64(409)
				var _433_v7 _dafny.Map
				_ = _433_v7
				_433_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _426_v0)
				var _434_v8 _dafny.Map
				_ = _434_v8
				_434_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen((p2), 0))).Int()).(bool))
				var _rhs50 _dafny.Map = ((_433_v7).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _426_v0))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_434_v8).Cardinality())).Merge(_433_v7))
				_ = _rhs50
				var _rhs51 bool = p0
				_ = _rhs51
				_433_v7 = _rhs50
				r0 = _rhs51
				_426_v0 = _426_v0
			} else {
				var _435_v9 *C0
				_ = _435_v9
				var _nw75 *C0 = New_C0_()
				_ = _nw75
				_nw75.Ctor__(!(p0))
				_435_v9 = _nw75
				var _436_v10 D5
				_ = _436_v10
				_436_v10 = Companion_D5_.Create_DC12_(_426_v0)
				(globalState).F2 = (_436_v10).Dtor_cf20()
				(globalState).F0 = _426_v0
				var _437_v11 _dafny.Sequence
				_ = _437_v11
				_437_v11 = _dafny.UnicodeSeqOfUtf8Bytes("ts")
				_437_v11 = (_this).F11()
				(globalState).F0 = _426_v0
			}
			if _this.F12 {
				var _438_v12 _dafny.Array
				_ = _438_v12
				var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
				_ = _nw76
				_438_v12 = _nw76
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))
				_ = _index62
				(_438_v12).ArraySet1(Companion_Default___.Fm24(_426_v0, globalState), (_index62).Int())
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))
				_ = _index63
				(_438_v12).ArraySet1(_dafny.IntOfInt64(-795), (_index63).Int())
				(globalState).F0 = Companion_Default___.SafeDivisionInt((_426_v0).Minus(_426_v0), _426_v0)
				(globalState).F0 = _426_v0
				var _439_v13 _dafny.Set
				_ = _439_v13
				_439_v13 = _dafny.SetOf((func() _dafny.Int {
					if (_430_v4).Contains(p1) {
						return (_430_v4).Multiplicity(p1)
					}
					return _426_v0
				})(), (_438_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))).Int()).(_dafny.Int), (_438_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))).Int()).(_dafny.Int))
				var _440_v14 _dafny.Map
				_ = _440_v14
				_440_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_430_v4, _426_v0)
				var _441_v15 _dafny.CodePoint
				_ = _441_v15
				_441_v15 = _dafny.CodePoint('m')
				var _442_v16 _dafny.Map
				_ = _442_v16
				_442_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v15, (_438_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))).Int()).(_dafny.Int))
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))
				_ = _index64
				(_438_v12).ArraySet1((((_439_v13).Union(_439_v13)).Cardinality()).Times(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if !(p0) {
						return Companion_Default___.Fm25((_440_v14).Cardinality(), (_442_v16).Update(_441_v15, _426_v0), _441_v15, globalState)
					}
					return _428_v2
				})()).Cardinality())), (_index64).Int())
				(globalState).F1 = ((_438_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))).Int()).(_dafny.Int)).Cmp((_dafny.Zero).Minus((_438_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(714), _dafny.ArrayLen((_438_v12), 0))).Int()).(_dafny.Int))) != 0
			} else {
				var _443_v17 D1
				_ = _443_v17
				_443_v17 = Companion_D1_.Create_DC5_()
				var _444_v18 _dafny.CodePoint
				_ = _444_v18
				_444_v18 = _dafny.CodePoint('o')
				var _445_v19 _dafny.Array
				_ = _445_v19
				var _nwElement0_14 _dafny.CodePoint = _444_v18
				_ = _nwElement0_14
				var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(10))
				_ = _nw77
				(_nw77).ArraySet1CodePoint(_nwElement0_14, 0)
				(_nw77).ArraySet1CodePoint(((_this).F11()).Select((Companion_Default___.SafeIndex(_426_v0, _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
				(_nw77).ArraySet1CodePoint(_444_v18, 2)
				(_nw77).ArraySet1CodePoint(_444_v18, 3)
				(_nw77).ArraySet1CodePoint(_444_v18, 4)
				(_nw77).ArraySet1CodePoint(_444_v18, 5)
				(_nw77).ArraySet1CodePoint(_444_v18, 6)
				(_nw77).ArraySet1CodePoint(_444_v18, 7)
				(_nw77).ArraySet1CodePoint(_444_v18, 8)
				(_nw77).ArraySet1CodePoint(_444_v18, 9)
				_445_v19 = _nw77
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_445_v19), 0))
				_ = _index65
				(_445_v19).ArraySet1CodePoint(_444_v18, (_index65).Int())
				var _446_v20 _dafny.Sequence
				_ = _446_v20
				_446_v20 = _dafny.SeqOf((_430_v4).Update(false, Companion_Default___.Abs(_426_v0)))
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_445_v19), 0))
				_ = _index66
				var _rhs52 D1 = Companion_D1_.Create_DC5_()
				_ = _rhs52
				var _rhs53 _dafny.CodePoint = _444_v18
				_ = _rhs53
				var _rhs54 _dafny.CodePoint = ((_this).F11()).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_426_v0), _dafny.IntOfUint32(((_this).F11()).Cardinality()))).Uint32()).(_dafny.CodePoint)
				_ = _rhs54
				var _rhs55 _dafny.Int = _426_v0
				_ = _rhs55
				var _rhs56 bool = (_426_v0).Cmp(_dafny.IntOfUint32((_446_v20).Cardinality())) > 0
				_ = _rhs56
				var _lhs44 _dafny.Array = _445_v19
				_ = _lhs44
				var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(458), _dafny.ArrayLen((_445_v19), 0))
				_ = _lhs45
				var _lhs46 *GlobalState = globalState
				_ = _lhs46
				var _lhs47 *C2 = _this
				_ = _lhs47
				_443_v17 = _rhs52
				_444_v18 = _rhs53
				(_lhs44).ArraySet1CodePoint(_rhs54, (_lhs45).Int())
				_lhs46.F2 = _rhs55
				_lhs47.F12 = _rhs56
				var _447_v21 _dafny.Sequence
				_ = _447_v21
				_447_v21 = _dafny.SeqOf(p1)
				var _rhs57 _dafny.Int = _426_v0
				_ = _rhs57
				var _rhs58 _dafny.Sequence = _447_v21
				_ = _rhs58
				var _lhs48 *GlobalState = globalState
				_ = _lhs48
				_lhs48.F2 = _rhs57
				_447_v21 = _rhs58
				var _448_v22 _dafny.Array
				_ = _448_v22
				var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw78
				_448_v22 = _nw78
				_448_v22 = _448_v22
				(globalState).F1 = false
				_428_v2 = _428_v2
			}
			var _449_v23 _dafny.CodePoint
			_ = _449_v23
			_449_v23 = _dafny.CodePoint('s')
			var _450_v24 _dafny.Set
			_ = _450_v24
			_450_v24 = _dafny.SetOf(_426_v0)
			var _451_v25 D0
			_ = _451_v25
			_451_v25 = Companion_D0_.Create_DC2_((_dafny.SetOf((_430_v4).Cardinality())).Cardinality(), _449_v23, _426_v0, _426_v0, (_450_v24).Cardinality())
			var _pat_let_tv13 = _426_v0
			_ = _pat_let_tv13
			var _452_v26 D0
			_ = _452_v26
			_452_v26 = Companion_D0_.Create_DC0_(Companion_Default___.Fm0(_this.F12, p1, func(_pat_let14_0 D0) D0 {
				return func(_453_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let15_0 _dafny.Int) D0 {
						return func(_454_dt__update_hcf5_h0 _dafny.Int) D0 {
							return Companion_D0_.Create_DC2_(_454_dt__update_hcf5_h0, (_453_dt__update__tmp_h0).Dtor_cf6(), (_453_dt__update__tmp_h0).Dtor_cf7(), (_453_dt__update__tmp_h0).Dtor_cf8(), (_453_dt__update__tmp_h0).Dtor_cf9())
						}(_pat_let15_0)
					}(_pat_let_tv13)
				}(_pat_let14_0)
			}(_451_v25), globalState))
			var _455_v27 _dafny.Set
			_ = _455_v27
			_455_v27 = _dafny.SetOf(_this.F12, true, p0)
			var _456_v28 _dafny.Map
			_ = _456_v28
			_456_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_426_v0, _dafny.IntOfInt64(168))).Cardinality(), _455_v27)
			_452_v26 = Companion_D0_.Create_DC0_((_456_v28).Merge(_456_v28))
			var _457_v29 _dafny.Map
			_ = _457_v29
			_457_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_428_v2).Cardinality()))
			_457_v29 = (_457_v29).Update(!(p0), _426_v0)
		} else {
			var _458_v30 _dafny.Int
			_ = _458_v30
			_458_v30 = _dafny.IntOfInt64(12)
			var _459_v31 _dafny.Map
			_ = _459_v31
			_459_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, p0)).Cardinality(), _458_v30)
			var _460_v32 _dafny.Array
			_ = _460_v32
			var _nwElement0_15 _dafny.Int = _dafny.IntOfInt64(914)
			_ = _nwElement0_15
			var _nw79 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(7))
			_ = _nw79
			(_nw79).ArraySet1(_nwElement0_15, 0)
			(_nw79).ArraySet1((_459_v31).Cardinality(), 1)
			(_nw79).ArraySet1(_458_v30, 2)
			(_nw79).ArraySet1(_dafny.IntOfUint32(((_this).F11()).Cardinality()), 3)
			(_nw79).ArraySet1(_458_v30, 4)
			(_nw79).ArraySet1(_458_v30, 5)
			(_nw79).ArraySet1(_458_v30, 6)
			_460_v32 = _nw79
			var _461_v33 _dafny.Set
			_ = _461_v33
			_461_v33 = _dafny.SetOf(_460_v32)
			var _462_v34 D5
			_ = _462_v34
			_462_v34 = Companion_D5_.Create_DC11_(_461_v33)
			var _463_v35 _dafny.Sequence
			_ = _463_v35
			_463_v35 = _dafny.UnicodeSeqOfUtf8Bytes("lwclh")
			var _464_v36 _dafny.Sequence
			_ = _464_v36
			_464_v36 = _dafny.SeqOf(_this.F12)
			var _465_v37 _dafny.Map
			_ = _465_v37
			_465_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_458_v30, Companion_Default___.Fm26(p1, globalState))
			var _466_v38 D0
			_ = _466_v38
			_466_v38 = Companion_D0_.Create_DC0_(_465_v37)
			var _rhs59 D5 = _462_v34
			_ = _rhs59
			var _rhs60 bool = Companion_Default___.Fm1((func() D0 {
				if (_464_v36).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("losi")).Cardinality()), _dafny.IntOfUint32((_464_v36).Cardinality()))).Uint32()).(bool) {
					return _466_v38
				}
				return _466_v38
			})(), (false) || (p0), _458_v30, p1, globalState)
			_ = _rhs60
			var _rhs61 _dafny.Sequence = Companion_Default___.Fm2((_this).F11(), globalState)
			_ = _rhs61
			var _rhs62 bool = p0
			_ = _rhs62
			var _rhs63 bool = !((p3).Dtor_cf2())
			_ = _rhs63
			var _lhs49 *GlobalState = globalState
			_ = _lhs49
			var _lhs50 *GlobalState = globalState
			_ = _lhs50
			_462_v34 = _rhs59
			_lhs49.F1 = _rhs60
			_463_v35 = _rhs61
			r0 = _rhs62
			_lhs50.F1 = _rhs63
			(globalState).F2 = Companion_Default___.SafeModuloInt(_458_v30, _458_v30)
			var _467_v39 _dafny.Sequence
			_ = _467_v39
			_467_v39 = _dafny.SeqOf(_dafny.IntOfUint32((_463_v35).Cardinality()))
			var _468_v40 D3
			_ = _468_v40
			_468_v40 = Companion_D3_.Create_DC8_(Companion_D1_.Create_DC4_(_467_v39), _460_v32, _dafny.IntOfUint32(((_this).F11()).Cardinality()))
			var _469_v41 D3
			_ = _469_v41
			_469_v41 = Companion_D3_.Create_DC9_((func() D3 {
				if p0 {
					return _468_v40
				}
				return _468_v40
			})())
			var _470_v42 _dafny.Map
			_ = _470_v42
			_470_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm27(globalState), _467_v39)
			var _471_v43 _dafny.Sequence
			_ = _471_v43
			_471_v43 = _dafny.SeqOf(_464_v36, _464_v36, _dafny.SeqOf(Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_465_v37), false, _dafny.IntOfInt64(842), p1, globalState), p1))
			var _472_v44 _dafny.Map
			_ = _472_v44
			_472_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F12)
			var _rhs64 D3 = _469_v41
			_ = _rhs64
			var _rhs65 _dafny.Sequence = (func() _dafny.Sequence {
				if (_470_v42).Contains(_471_v43) {
					return (_470_v42).Get(_471_v43).(_dafny.Sequence)
				}
				return _dafny.SeqOf((_472_v44).Cardinality())
			})()
			_ = _rhs65
			var _rhs66 _dafny.Int = (_458_v30).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((func() _dafny.Int {
				if _this.F12 {
					return _458_v30
				}
				return _458_v30
			})())))
			_ = _rhs66
			_469_v41 = _rhs64
			_467_v39 = _rhs65
			_458_v30 = _rhs66
			r0 = _this.F12
			_459_v31 = (_459_v31).Update(_458_v30, _458_v30)
		}
		var _473_v45 _dafny.Int
		_ = _473_v45
		_473_v45 = _dafny.IntOfInt64(997)
		var _474_v46 _dafny.Sequence
		_ = _474_v46
		_474_v46 = _dafny.SeqOf(_473_v45)
		_474_v46 = _474_v46
		if p1 {
			var _475_v47 _dafny.Sequence
			_ = _475_v47
			_475_v47 = _dafny.SeqOf(!(_this.F12))
			(globalState).F0 = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_473_v45).Minus(_dafny.IntOfUint32((_475_v47).Cardinality()))), _dafny.IntOfInt64(921))
			r0 = _this.F12
			var _476_v48 _dafny.Array
			_ = _476_v48
			var _nw80 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(23))
			_ = _nw80
			_476_v48 = _nw80
			var _477_v49 _dafny.Sequence
			_ = _477_v49
			_477_v49 = _dafny.SeqOf(_476_v48, _476_v48)
			(globalState).F2 = _dafny.IntOfUint32((_477_v49).Cardinality())
			(globalState).F1 = !(p0)
			var _478_v50 *C1
			_ = _478_v50
			var _nw81 *C1 = New_C1_()
			_ = _nw81
			_nw81.Ctor__()
			_478_v50 = _nw81
		} else {
			var _479_v51 *C0
			_ = _479_v51
			var _nw82 *C0 = New_C0_()
			_ = _nw82
			_nw82.Ctor__(p0)
			_479_v51 = _nw82
			if !_dafny.Companion_Sequence_.Equal((_this).F11(), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("cyofsjs"), (_this).F11())) {
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p2), 0))
				_ = _index67
				(p2).ArraySet1((func() bool {
					if (_479_v51).F13() {
						return p1
					}
					return p1
				})(), (_index67).Int())
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p2), 0))
				_ = _index68
				(p2).ArraySet1(p0, (_index68).Int())
				var _480_v52 _dafny.MultiSet
				_ = _480_v52
				_480_v52 = _dafny.MultiSetOf(p0)
				var _481_v53 _dafny.Set
				_ = _481_v53
				_481_v53 = _dafny.SetOf(p1)
				var _482_v54 D0
				_ = _482_v54
				_482_v54 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
					if (_480_v52).Contains((_479_v51).F13()) {
						return (_480_v52).Multiplicity((_479_v51).F13())
					}
					return _dafny.IntOfUint32((_474_v46).Cardinality())
				})(), _481_v53))
				var _483_v55 _dafny.Map
				_ = _483_v55
				_483_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_479_v51).F13(), _dafny.IntOfInt64(-292))
				(globalState).F5 = Companion_Default___.Fm1(_482_v54, (_479_v51).F13(), ((_483_v55).Cardinality()).Times(_473_v45), (_479_v51).F13(), globalState)
				var _484_v56 _dafny.Array
				_ = _484_v56
				var _nwElement0_16 bool = p1
				_ = _nwElement0_16
				var _nw83 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(6))
				_ = _nw83
				(_nw83).ArraySet1(_nwElement0_16, 0)
				(_nw83).ArraySet1((_480_v52).Contains((_479_v51).F13()), 1)
				(_nw83).ArraySet1((_479_v51).F13(), 2)
				(_nw83).ArraySet1((_479_v51).F13(), 3)
				(_nw83).ArraySet1((func() bool {
					if (_479_v51).F13() {
						return p0
					}
					return Companion_Default___.Fm1(_482_v54, _this.F12, _473_v45, p1, globalState)
				})(), 4)
				(_nw83).ArraySet1(_this.F12, 5)
				_484_v56 = _nw83
				var _485_v57 _dafny.Array
				_ = _485_v57
				_485_v57 = _484_v56
				var _486_v58 _dafny.Sequence
				_ = _486_v58
				_486_v58 = _dafny.SeqOf((_485_v57), _484_v56)
				_484_v56 = (_486_v58).Select((Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32((_486_v58).Cardinality()))).Uint32()).(_dafny.Array)
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p2), 0))
				_ = _index69
				(p2).ArraySet1((func() bool {
					if _this.F12 {
						return (func() bool {
							if (_479_v51).F13() {
								return _this.F12
							}
							return (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p2), 0))).Int()).(bool)
						})()
					}
					return (_479_v51).F13()
				})(), (_index69).Int())
				var _487_v59 _dafny.Map
				_ = _487_v59
				_487_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_473_v45, _dafny.SetOf(p0, (p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(889), _dafny.ArrayLen((p2), 0))).Int()).(bool), p1, p0))
				(globalState).F1 = (func() bool {
					if Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_487_v59), _this.F12, (_481_v53).Cardinality(), (_479_v51).F13(), globalState) {
						return p0
					}
					return (_479_v51).F13()
				})()
			} else {
				var _488_v61 _dafny.Map
				_ = _488_v61
				_488_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _473_v45)
				var _489_v62 _dafny.Array
				_ = _489_v62
				var _nwElement0_17 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_473_v45), _474_v46)
				_ = _nwElement0_17
				var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(20))
				_ = _nw84
				(_nw84).ArraySet1(_nwElement0_17, 0)
				(_nw84).ArraySet1(_474_v46, 1)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(637))).Uint32(), func(coer18 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg18 _dafny.Int) interface{} {
						return coer18(arg18)
					}
				}((func(_490_v45 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_491_i0 _dafny.Int) _dafny.Int {
						return _490_v45
					}
				})(_473_v45))), _474_v46), 2)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(906))).Uint32(), func(coer19 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}((func(_492_v45 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_493_i1 _dafny.Int) _dafny.Int {
						return _492_v45
					}
				})(_473_v45))), _474_v46), 3)
				(_nw84).ArraySet1(_474_v46, 4)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_474_v46, _474_v46), 5)
				(_nw84).ArraySet1((Companion_D1_.Create_DC4_(_474_v46)).Dtor_cf11(), 6)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Update(_474_v46, (Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32((_474_v46).Cardinality()))).Uint32(), (_474_v46).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(func() _dafny.Set {
					var _coll27 = _dafny.NewBuilder()
					_ = _coll27
					for _iter28 := _dafny.Iterate((_474_v46).Elements()); ; {
						_compr_27, _ok28 := _iter28()
						if !_ok28 {
							break
						}
						var _494_v60 _dafny.Int
						_494_v60 = interface{}(_compr_27).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_474_v46, _494_v60) {
							_coll27.Add((_494_v60).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-190))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg20 _dafny.Int) interface{} {
									return coer20(arg20)
								}
							}(func(_495_i2 _dafny.Int) _dafny.CodePoint {
								return _dafny.CodePoint('g')
							}))).Cardinality()))))
						}
					}
					return _coll27.ToSet()
				}())).Cardinality()), _dafny.IntOfUint32((_474_v46).Cardinality()))).Uint32()).(_dafny.Int)), 7)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_474_v46, _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_488_v61, (_479_v51).F13())).Cardinality(), _473_v45, _473_v45)), 8)
				(_nw84).ArraySet1(_474_v46, 9)
				(_nw84).ArraySet1(_474_v46, 10)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_474_v46, _474_v46), 11)
				(_nw84).ArraySet1(_474_v46, 12)
				(_nw84).ArraySet1(_dafny.Companion_Sequence_.Update(_474_v46, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("usfonr")).Cardinality()), _dafny.IntOfUint32((_474_v46).Cardinality()))).Uint32(), _473_v45), 13)
				(_nw84).ArraySet1((func() _dafny.Sequence {
					if (_479_v51).F13() {
						return _474_v46
					}
					return _474_v46
				})(), 14)
				(_nw84).ArraySet1(_dafny.SeqOf(_473_v45), 15)
				(_nw84).ArraySet1(_474_v46, 16)
				(_nw84).ArraySet1(_474_v46, 17)
				(_nw84).ArraySet1(_474_v46, 18)
				(_nw84).ArraySet1((func() _dafny.Sequence {
					if (_479_v51).F13() {
						return _474_v46
					}
					return _474_v46
				})(), 19)
				_489_v62 = _nw84
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_489_v62), 0))
				_ = _index70
				(_489_v62).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_474_v46, _474_v46), (_index70).Int())
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(178), _dafny.ArrayLen((_489_v62), 0))
				_ = _index71
				(_489_v62).ArraySet1(_474_v46, (_index71).Int())
				var _496_v63 D3
				_ = _496_v63
				_496_v63 = Companion_D3_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("x"))
				var _497_v64 _dafny.Array
				_ = _497_v64
				var _nwElement0_18 _dafny.Sequence = (_this).F11()
				_ = _nwElement0_18
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(15))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_18, 0)
				(_nw85).ArraySet1(Companion_Default___.Fm2((_this).F11(), globalState), 1)
				(_nw85).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(728))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}(func(_498_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('r')
				})), 2)
				(_nw85).ArraySet1((_this).F11(), 3)
				(_nw85).ArraySet1((_this).F11(), 4)
				(_nw85).ArraySet1((_this).F11(), 5)
				(_nw85).ArraySet1((_this).F11(), 6)
				(_nw85).ArraySet1(Companion_Default___.Fm2((_this).F11(), globalState), 7)
				(_nw85).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(382))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg22 _dafny.Int) interface{} {
						return coer22(arg22)
					}
				}(func(_499_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})), 8)
				(_nw85).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F11(), (_this).F11()), 9)
				(_nw85).ArraySet1((_this).F11(), 10)
				(_nw85).ArraySet1((_this).F11(), 11)
				(_nw85).ArraySet1((_496_v63).Dtor_cf13(), 12)
				(_nw85).ArraySet1((_this).F11(), 13)
				(_nw85).ArraySet1((_this).F11(), 14)
				_497_v64 = _nw85
				_497_v64 = _497_v64
				var _500_v67 _dafny.Sequence
				_ = _500_v67
				_500_v67 = _dafny.SeqOf(true)
				var _501_v68 _dafny.Set
				_ = _501_v68
				_501_v68 = _dafny.SetOf(_473_v45, _473_v45, _dafny.IntOfInt64(162), _dafny.IntOfUint32((_500_v67).Cardinality()))
				var _502_v69 _dafny.Sequence
				_ = _502_v69
				_502_v69 = _dafny.SeqOf(_501_v68, _501_v68, _501_v68, _501_v68, _501_v68)
				var _503_v70 _dafny.CodePoint
				_ = _503_v70
				_503_v70 = _dafny.CodePoint('p')
				var _504_v71 _dafny.Map
				_ = _504_v71
				_504_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_503_v70, _473_v45)
				var _505_v72 D5
				_ = _505_v72
				_505_v72 = Companion_D5_.Create_DC12_(_473_v45)
				var _506_v73 _dafny.Map
				_ = _506_v73
				_506_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_473_v45, _dafny.IntOfUint32(((_this).F11()).Cardinality()), (_504_v71).Cardinality(), _473_v45, _473_v45), _505_v72)
				var _507_v74 _dafny.Map
				_ = _507_v74
				_507_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
				var _508_v76 _dafny.Map
				_ = _508_v76
				_508_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v68, _473_v45)
				var _509_v77 _dafny.Sequence
				_ = _509_v77
				_509_v77 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_473_v45, (_507_v74).Cardinality()), _505_v72), func() _dafny.Map {
					var _coll28 = _dafny.NewMapBuilder()
					_ = _coll28
					for _iter29 := _dafny.Iterate((_508_v76).Keys().Elements()); ; {
						_compr_28, _ok29 := _iter29()
						if !_ok29 {
							break
						}
						var _510_v75 _dafny.Set
						_510_v75 = interface{}(_compr_28).(_dafny.Set)
						if (_508_v76).Contains(_510_v75) {
							_coll28.Add(_510_v75, _505_v72)
						}
					}
					return _coll28.ToMap()
				}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_502_v69).Select((Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32((_502_v69).Cardinality()))).Uint32()).(_dafny.Set), _505_v72), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v68, _505_v72))
				var _511_v79 _dafny.Set
				_ = _511_v79
				_511_v79 = _dafny.SetOf(_501_v68, _501_v68)
				var _pat_let_tv14 = _473_v45
				_ = _pat_let_tv14
				var _512_v80 _dafny.Array
				_ = _512_v80
				var _nwElement0_19 _dafny.Map = func() _dafny.Map {
					var _coll29 = _dafny.NewMapBuilder()
					_ = _coll29
					for _iter30 := _dafny.Iterate((func() _dafny.Map {
						var _coll30 = _dafny.NewMapBuilder()
						_ = _coll30
						for _iter31 := _dafny.Iterate((_502_v69).Elements()); ; {
							_compr_30, _ok31 := _iter31()
							if !_ok31 {
								break
							}
							var _513_v66 _dafny.Set
							_513_v66 = interface{}(_compr_30).(_dafny.Set)
							if _dafny.Companion_Sequence_.Contains(_502_v69, _513_v66) {
								_coll30.Add(_513_v66, p0)
							}
						}
						return _coll30.ToMap()
					}()).Keys().Elements()); ; {
						_compr_29, _ok30 := _iter30()
						if !_ok30 {
							break
						}
						var _514_v65 _dafny.Set
						_514_v65 = interface{}(_compr_29).(_dafny.Set)
						if (func() _dafny.Map {
							var _coll31 = _dafny.NewMapBuilder()
							_ = _coll31
							for _iter32 := _dafny.Iterate((_502_v69).Elements()); ; {
								_compr_31, _ok32 := _iter32()
								if !_ok32 {
									break
								}
								var _513_v66 _dafny.Set
								_513_v66 = interface{}(_compr_31).(_dafny.Set)
								if _dafny.Companion_Sequence_.Contains(_502_v69, _513_v66) {
									_coll31.Add(_513_v66, p0)
								}
							}
							return _coll31.ToMap()
						}()).Contains(_514_v65) {
							_coll29.Add(_514_v65, Companion_D5_.Create_DC12_(_473_v45))
						}
					}
					return _coll29.ToMap()
				}()
				_ = _nwElement0_19
				var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(17))
				_ = _nw86
				(_nw86).ArraySet1(_nwElement0_19, 0)
				(_nw86).ArraySet1(_506_v73, 1)
				(_nw86).ArraySet1((_509_v77).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_473_v45), _dafny.IntOfUint32((_509_v77).Cardinality()))).Uint32()).(_dafny.Map), 2)
				(_nw86).ArraySet1(func() _dafny.Map {
					var _coll32 = _dafny.NewMapBuilder()
					_ = _coll32
					for _iter33 := _dafny.Iterate((_511_v79).Elements()); ; {
						_compr_32, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _515_v78 _dafny.Set
						_515_v78 = interface{}(_compr_32).(_dafny.Set)
						if (_511_v79).Contains(_515_v78) {
							_coll32.Add(_515_v78, _505_v72)
						}
					}
					return _coll32.ToMap()
				}(), 3)
				(_nw86).ArraySet1(_506_v73, 4)
				(_nw86).ArraySet1(_506_v73, 5)
				(_nw86).ArraySet1(_506_v73, 6)
				(_nw86).ArraySet1(_506_v73, 7)
				(_nw86).ArraySet1(_506_v73, 8)
				(_nw86).ArraySet1(_506_v73, 9)
				(_nw86).ArraySet1(_506_v73, 10)
				(_nw86).ArraySet1(_506_v73, 11)
				(_nw86).ArraySet1(_506_v73, 12)
				(_nw86).ArraySet1(_506_v73, 13)
				(_nw86).ArraySet1(_506_v73, 14)
				(_nw86).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v68, func(_pat_let16_0 D5) D5 {
					return func(_516_dt__update__tmp_h1 D5) D5 {
						return func(_pat_let17_0 _dafny.Int) D5 {
							return func(_517_dt__update_hcf20_h0 _dafny.Int) D5 {
								return Companion_D5_.Create_DC12_(_517_dt__update_hcf20_h0)
							}(_pat_let17_0)
						}(_pat_let_tv14)
					}(_pat_let16_0)
				}(_505_v72)), 15)
				(_nw86).ArraySet1(_506_v73, 16)
				_512_v80 = _nw86
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_512_v80), 0))
				_ = _index72
				(_512_v80).ArraySet1(_506_v73, (_index72).Int())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(511), _dafny.ArrayLen((_512_v80), 0))
				_ = _index73
				(_512_v80).ArraySet1((_506_v73).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_501_v68, _505_v72)), (_index73).Int())
				var _518_v81 _dafny.Map
				_ = _518_v81
				_518_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_473_v45, _473_v45)
				var _519_v82 D1
				_ = _519_v82
				_519_v82 = Companion_D1_.Create_DC4_(_dafny.SeqOf((_518_v81).Cardinality(), _473_v45, _dafny.IntOfUint32((_500_v67).Cardinality())))
				var _520_v83 _dafny.Array
				_ = _520_v83
				var _len0_13 _dafny.Int = _dafny.IntOfInt64(8)
				_ = _len0_13
				var _nw87 _dafny.Array
				_ = _nw87
				if _len0_13.Cmp(_dafny.Zero) == 0 {
					_nw87 = _dafny.NewArray(_len0_13)
				} else {
					var _init13 func(_dafny.Int) _dafny.Int = (func(_521_v45 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_522_i5 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_522_i5, _521_v45)
						}
					})(_473_v45)
					_ = _init13
					var _element0_13 = _init13(_dafny.Zero)
					_ = _element0_13
					_nw87 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
					(_nw87).ArraySet1(_element0_13, 0)
					var _nativeLen0_13 = (_len0_13).Int()
					_ = _nativeLen0_13
					for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
						(_nw87).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
					}
				}
				_520_v83 = _nw87
				var _523_v84 D3
				_ = _523_v84
				_523_v84 = Companion_D3_.Create_DC8_(_519_v82, _520_v83, _dafny.IntOfUint32(((_this).F11()).Cardinality()))
				(globalState).F2 = ((func() _dafny.Int {
					if p0 {
						return Companion_Default___.Fm24((_523_v84).Dtor_cf16(), globalState)
					}
					return _473_v45
				})()).Times(_473_v45)
				var _524_v85 _dafny.Sequence
				_ = _524_v85
				_524_v85 = _dafny.SeqOf((_this).F11())
				_524_v85 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_524_v85, _524_v85), (Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_524_v85, _524_v85)).Cardinality()))).Uint32(), (_this).F11()), _524_v85)
			}
			var _525_v86 _dafny.CodePoint
			_ = _525_v86
			_525_v86 = _dafny.CodePoint('p')
			_525_v86 = _525_v86
			var _526_v87 _dafny.Sequence
			_ = _526_v87
			_526_v87 = _dafny.UnicodeSeqOfUtf8Bytes("bdx")
			var _527_v88 D3
			_ = _527_v88
			_527_v88 = Companion_D3_.Create_DC7_(_526_v87)
			_526_v87 = (_527_v88).Dtor_cf13()
			var _528_v90 D0
			_ = _528_v90
			_528_v90 = Companion_D0_.Create_DC0_(func() _dafny.Map {
				var _coll33 = _dafny.NewMapBuilder()
				_ = _coll33
				for _iter34 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-760), _dafny.IntOfInt64(-629))); ; {
					_compr_33, _ok34 := _iter34()
					if !_ok34 {
						break
					}
					var _529_v89 _dafny.Int
					_529_v89 = interface{}(_compr_33).(_dafny.Int)
					if ((_dafny.IntOfInt64(-760)).Cmp(_529_v89) <= 0) && ((_529_v89).Cmp(_dafny.IntOfInt64(-629)) < 0) {
						_coll33.Add(Companion_Default___.SafeModuloInt(_529_v89, (_dafny.SetOf(false, (_479_v51).F13())).Cardinality()), _dafny.SetOf(false))
					}
				}
				return _coll33.ToMap()
			}())
			var _530_v91 _dafny.MultiSet
			_ = _530_v91
			_530_v91 = _dafny.MultiSetOf(Companion_Default___.Fm1(_528_v90, !((_479_v51).F13()), _473_v45, !((_479_v51).F13()), globalState))
			var _531_v92 *C1
			_ = _531_v92
			var _nw88 *C1 = New_C1_()
			_ = _nw88
			_nw88.Ctor__()
			_531_v92 = _nw88
			var _532_v93 _dafny.Array
			_ = _532_v93
			var _nwElement0_20 _dafny.Int = _473_v45
			_ = _nwElement0_20
			var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(10))
			_ = _nw89
			(_nw89).ArraySet1(_nwElement0_20, 0)
			(_nw89).ArraySet1(_473_v45, 1)
			(_nw89).ArraySet1(((_dafny.MultiSetOf(p0)).Difference(_530_v91)).Cardinality(), 2)
			(_nw89).ArraySet1(Companion_Default___.SafeDivisionInt(_473_v45, _dafny.IntOfUint32(((_this).F11()).Cardinality())), 3)
			(_nw89).ArraySet1(_473_v45, 4)
			(_nw89).ArraySet1(_473_v45, 5)
			(_nw89).ArraySet1(_dafny.IntOfUint32((_526_v87).Cardinality()), 6)
			(_nw89).ArraySet1(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_dafny.Zero).Minus(_473_v45)), _473_v45), 7)
			(_nw89).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.SetOf(_531_v92, _531_v92, _531_v92, _531_v92, _531_v92)).Cardinality(), _473_v45), 8)
			(_nw89).ArraySet1(_dafny.IntOfUint32(((_527_v88).Dtor_cf13()).Cardinality()), 9)
			_532_v93 = _nw89
			var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_532_v93), 0))
			_ = _index74
			(_532_v93).ArraySet1(_473_v45, (_index74).Int())
			var _533_v94 _dafny.MultiSet
			_ = _533_v94
			_533_v94 = _dafny.MultiSetOf((_dafny.Zero).Minus(_473_v45))
			var _534_v95 D3
			_ = _534_v95
			_534_v95 = Companion_D3_.Create_DC8_(Companion_D1_.Create_DC4_(_474_v46), _532_v93, (_533_v94).Cardinality())
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_532_v93), 0))
			_ = _index75
			var _rhs67 _dafny.Int = (_534_v95).Dtor_cf16()
			_ = _rhs67
			var _rhs68 bool = p1
			_ = _rhs68
			var _lhs51 _dafny.Array = _532_v93
			_ = _lhs51
			var _lhs52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_532_v93), 0))
			_ = _lhs52
			var _lhs53 *GlobalState = globalState
			_ = _lhs53
			(_lhs51).ArraySet1(_rhs67, (_lhs52).Int())
			_lhs53.F5 = _rhs68
		}
		var _535_v96 _dafny.Array
		_ = _535_v96
		var _len0_14 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_14
		var _nw90 _dafny.Array
		_ = _nw90
		if _len0_14.Cmp(_dafny.Zero) == 0 {
			_nw90 = _dafny.NewArray(_len0_14)
		} else {
			var _init14 func(_dafny.Int) _dafny.Int = (func(_536_p0 bool) func(_dafny.Int) _dafny.Int {
				return func(_537_i7 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_537_i7, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _536_p0)).Cardinality())
				}
			})(p0)
			_ = _init14
			var _element0_14 = _init14(_dafny.Zero)
			_ = _element0_14
			_nw90 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
			(_nw90).ArraySet1(_element0_14, 0)
			var _nativeLen0_14 = (_len0_14).Int()
			_ = _nativeLen0_14
			for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
				(_nw90).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
			}
		}
		_535_v96 = _nw90
		for _iter35 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_535_v96), 0))); ; {
			_guard_loop_1, _ok35 := _iter35()
			if !_ok35 {
				break
			}
			var _538_i6 _dafny.Int
			_538_i6 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_538_i6).Sign() != -1) && ((_538_i6).Cmp(_dafny.ArrayLen((_535_v96), 0)) < 0)) {
				(_535_v96).ArraySet1((_538_i6).Plus((_dafny.SetOf(_473_v45, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(249))).Uint32(), func(coer23 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_539_i8 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('o')
				}))).Cardinality()))).Cardinality()), (_538_i6).Int())
			}
		}
		var _540_v97 _dafny.Sequence
		_ = _540_v97
		_540_v97 = _dafny.SeqOf(_this.F12, p0, _this.F12, false)
		var _541_v98 _dafny.Set
		_ = _541_v98
		_541_v98 = _dafny.SetOf(_540_v97, _540_v97)
		if (_541_v98).IsSubsetOf(_541_v98) {
			var _542_v99 _dafny.CodePoint
			_ = _542_v99
			_542_v99 = _dafny.CodePoint('t')
			var _543_v100 _dafny.Map
			_ = _543_v100
			_543_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _473_v45)
			var _544_v101 D0
			_ = _544_v101
			_544_v101 = Companion_D0_.Create_DC2_(_473_v45, _542_v99, _473_v45, _473_v45, (func() _dafny.Int {
				if (_543_v100).Contains(p0) {
					return (_543_v100).Get(p0).(_dafny.Int)
				}
				return _473_v45
			})())
			var _545_v102 D0
			_ = _545_v102
			_545_v102 = Companion_D0_.Create_DC3_(_544_v101)
			var _546_v103 _dafny.Set
			_ = _546_v103
			_546_v103 = _dafny.SetOf(!(!(_this.F12)), p1)
			_545_v102 = (func() D0 {
				if (_546_v103).IsSubsetOf(_546_v103) {
					return _545_v102
				}
				return _545_v102
			})()
			var _547_v104 *C0
			_ = _547_v104
			var _nw91 *C0 = New_C0_()
			_ = _nw91
			_nw91.Ctor__(_this.F12)
			_547_v104 = _nw91
			var _548_v105 *C1
			_ = _548_v105
			var _nw92 *C1 = New_C1_()
			_ = _nw92
			_nw92.Ctor__()
			_548_v105 = _nw92
			(_this).F12 = _this.F12
			_542_v99 = _542_v99
		} else {
			var _549_v106 _dafny.Set
			_ = _549_v106
			_549_v106 = _dafny.SetOf(_473_v45)
			var _550_v107 D7
			_ = _550_v107
			_550_v107 = Companion_D7_.Create_DC17_(_540_v97)
			var _551_v108 _dafny.Map
			_ = _551_v108
			_551_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_473_v45, _549_v106)
			var _552_v112 _dafny.Array
			_ = _552_v112
			var _nwElement0_21 _dafny.Set = _549_v106
			_ = _nwElement0_21
			var _nw93 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(27))
			_ = _nw93
			(_nw93).ArraySet1(_nwElement0_21, 0)
			(_nw93).ArraySet1((_549_v106).Union(_549_v106), 1)
			(_nw93).ArraySet1(_549_v106, 2)
			(_nw93).ArraySet1(_dafny.SetOf(_dafny.IntOfUint32(((_550_v107).Dtor_cf28()).Cardinality())), 3)
			(_nw93).ArraySet1((func() _dafny.Set {
				if (_551_v108).Contains(_473_v45) {
					return (_551_v108).Get(_473_v45).(_dafny.Set)
				}
				return func() _dafny.Set {
					var _coll34 = _dafny.NewBuilder()
					_ = _coll34
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(170), _dafny.IntOfInt64(858))); ; {
						_compr_34, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _553_v109 _dafny.Int
						_553_v109 = interface{}(_compr_34).(_dafny.Int)
						if ((_dafny.IntOfInt64(170)).Cmp(_553_v109) <= 0) && ((_553_v109).Cmp(_dafny.IntOfInt64(858)) < 0) {
							_coll34.Add(Companion_Default___.SafeModuloInt(_553_v109, _473_v45))
						}
					}
					return _coll34.ToSet()
				}()
			})(), 4)
			(_nw93).ArraySet1(_549_v106, 5)
			(_nw93).ArraySet1((_549_v106).Union(_549_v106), 6)
			(_nw93).ArraySet1(_dafny.SetOf(_473_v45), 7)
			(_nw93).ArraySet1((_549_v106).Intersection(_549_v106), 8)
			(_nw93).ArraySet1(_549_v106, 9)
			(_nw93).ArraySet1((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_473_v45, p0)).Cardinality())).Intersection(_dafny.SetOf(_dafny.IntOfInt64(-132))), 10)
			(_nw93).ArraySet1(func() _dafny.Set {
				var _coll35 = _dafny.NewBuilder()
				_ = _coll35
				for _iter37 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(790), _dafny.IntOfInt64(697))); ; {
					_compr_35, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _554_v110 _dafny.Int
					_554_v110 = interface{}(_compr_35).(_dafny.Int)
					if ((_dafny.IntOfInt64(790)).Cmp(_554_v110) <= 0) && ((_554_v110).Cmp(_dafny.IntOfInt64(697)) < 0) {
						_coll35.Add((_554_v110).Plus((_474_v46).Select((Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32((_474_v46).Cardinality()))).Uint32()).(_dafny.Int)))
					}
				}
				return _coll35.ToSet()
			}(), 11)
			(_nw93).ArraySet1((_549_v106).Union(_549_v106), 12)
			(_nw93).ArraySet1(_549_v106, 13)
			(_nw93).ArraySet1(_549_v106, 14)
			(_nw93).ArraySet1(_549_v106, 15)
			(_nw93).ArraySet1(_549_v106, 16)
			(_nw93).ArraySet1((func() _dafny.Set {
				if p0 {
					return func() _dafny.Set {
						var _coll36 = _dafny.NewBuilder()
						_ = _coll36
						for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-491), _dafny.IntOfInt64(281))); ; {
							_compr_36, _ok38 := _iter38()
							if !_ok38 {
								break
							}
							var _555_v111 _dafny.Int
							_555_v111 = interface{}(_compr_36).(_dafny.Int)
							if ((_dafny.IntOfInt64(-491)).Cmp(_555_v111) <= 0) && ((_555_v111).Cmp(_dafny.IntOfInt64(281)) < 0) {
								_coll36.Add(Companion_Default___.SafeDivisionInt(_555_v111, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_474_v46).Cardinality()), _473_v45)).Cardinality()))
							}
						}
						return _coll36.ToSet()
					}()
				}
				return _549_v106
			})(), 17)
			(_nw93).ArraySet1((Companion_Default___.Fm22(globalState)).Union(_549_v106), 18)
			(_nw93).ArraySet1(_549_v106, 19)
			(_nw93).ArraySet1(_549_v106, 20)
			(_nw93).ArraySet1(_549_v106, 21)
			(_nw93).ArraySet1(_549_v106, 22)
			(_nw93).ArraySet1(_549_v106, 23)
			(_nw93).ArraySet1(_549_v106, 24)
			(_nw93).ArraySet1(_dafny.SetOf(_473_v45, _dafny.IntOfUint32((_540_v97).Cardinality()), _473_v45, _473_v45), 25)
			(_nw93).ArraySet1(_549_v106, 26)
			_552_v112 = _nw93
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_552_v112), 0))
			_ = _index76
			(_552_v112).ArraySet1(_549_v106, (_index76).Int())
			var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(740), _dafny.ArrayLen((_552_v112), 0))
			_ = _index77
			(_552_v112).ArraySet1((_549_v106).Union(func() _dafny.Set {
				var _coll37 = _dafny.NewBuilder()
				_ = _coll37
				for _iter39 := _dafny.Iterate((Companion_Default___.Fm22(globalState)).Elements()); ; {
					_compr_37, _ok39 := _iter39()
					if !_ok39 {
						break
					}
					var _556_v113 _dafny.Int
					_556_v113 = interface{}(_compr_37).(_dafny.Int)
					if (Companion_Default___.Fm22(globalState)).Contains(_556_v113) {
						_coll37.Add((_556_v113).Plus((Companion_D0_.Create_DC1_(_dafny.IntOfInt64(-491), !(true), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false)).Dtor_cf3()))
					}
				}
				return _coll37.ToSet()
			}()), (_index77).Int())
			(globalState).F1 = p0
			var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_535_v96), 0))
			_ = _index78
			(_535_v96).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_540_v97, _540_v97)).Cardinality()), (_index78).Int())
			var _557_v114 _dafny.Sequence
			_ = _557_v114
			_557_v114 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate((_this).F11(), (_this).F11()))
			var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_535_v96), 0))
			_ = _index79
			(_535_v96).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_557_v114).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_473_v45, _473_v45), _dafny.IntOfUint32((_557_v114).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32(((_557_v114).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_473_v45, _473_v45), _dafny.IntOfUint32((_557_v114).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), Companion_Default___.Fm28((_474_v46).Select((Companion_Default___.SafeIndex(_473_v45, _dafny.IntOfUint32((_474_v46).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(Companion_Default___.Fm24(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xvcevpcui")).Cardinality()), globalState)), _this.F12, globalState))).Cardinality()), (_index79).Int())
			var _558_v115 *C0
			_ = _558_v115
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__(false)
			_558_v115 = _nw94
			var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((p2), 0))
			_ = _index80
			(p2).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_540_v97, _540_v97), (_index80).Int())
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(251), _dafny.ArrayLen((p2), 0))
			_ = _index81
			(p2).ArraySet1((_558_v115).F13(), (_index81).Int())
		}
		var _559_v116 D6
		_ = _559_v116
		_559_v116 = Companion_D6_.Create_DC16_(_473_v45)
		(globalState).F2 = ((func() D6 {
			if _this.F12 {
				return Companion_Default___.Fm29(p0, _473_v45, globalState)
			}
			return _559_v116
		})()).Dtor_cf27()
		r0 = p1
		return r0
	}
}
func (_this *C2) F11() _dafny.Sequence {
	{
		return _this._f11
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f6  _dafny.Int
	_f9  _dafny.Int
	_f10 _dafny.Sequence
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f6 = _dafny.Zero
	_this._f9 = _dafny.Zero
	_this._f10 = _dafny.EmptySeq
	return &_this
}

type CompanionStruct_C3_ struct {
}

var Companion_C3_ = CompanionStruct_C3_{}

func (_this *C3) Equals(other *C3) bool {
	return _this == other
}

func (_this *C3) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C3)
	return ok && _this.Equals(other)
}

func (*C3) String() string {
	return "_module.C3"
}

func Type_C3_() _dafny.TypeDescriptor {
	return type_C3_{}
}

type type_C3_ struct {
}

func (_this type_C3_) Default() interface{} {
	return (*C3)(nil)
}

func (_this type_C3_) String() string {
	return "main.C3"
}
func (_this *C3) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_, Companion_T1_.TraitID_}
}

var _ T0 = &C3{}
var _ T1 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F6() _dafny.Int {
	return _this._f6
}
func (_this *C3) Ctor__(f9 _dafny.Int, f10 _dafny.Sequence, f6 _dafny.Int) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f6 = f6
	}
}
func (_this *C3) Fm16(p0 _dafny.MultiSet, globalState *GlobalState) bool {
	{
		return ((_this).F6()).Cmp((_this).F6()) != 0
	}
}
func (_this *C3) Fm17(p0 _dafny.Int, p1 _dafny.Set, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeModuloInt((_this).F9(), ((_dafny.SetOf(!(true))).Union(_dafny.SetOf(false))).Cardinality())
	}
}
func (_this *C3) Fm18(p0 bool, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true))
	}
}
func (_this *C3) M0(globalState *GlobalState) (_dafny.Int, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _560_v0 _dafny.Array
		_ = _560_v0
		var _nw95 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw95
		_560_v0 = _nw95
		for _iter40 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_560_v0), 0))); ; {
			_guard_loop_2, _ok40 := _iter40()
			if !_ok40 {
				break
			}
			var _561_i0 _dafny.Int
			_561_i0 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_561_i0).Sign() != -1) && ((_561_i0).Cmp(_dafny.ArrayLen((_560_v0), 0)) < 0)) {
				(_560_v0).ArraySet1((_561_i0).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-943), (_this).F6())), (_561_i0).Int())
			}
		}
		var _562_v1 D5
		_ = _562_v1
		_562_v1 = Companion_D5_.Create_DC11_(_dafny.SetOf(_560_v0, _560_v0, _560_v0))
		var _563_v2 _dafny.Set
		_ = _563_v2
		_563_v2 = _dafny.SetOf(_560_v0, _560_v0, _560_v0)
		var _564_v3 _dafny.Sequence
		_ = _564_v3
		_564_v3 = _dafny.SeqOf(true, ((_this).F9()).Cmp((_this).F6()) > 0, (_563_v2).IsProperSubsetOf((_562_v1).Dtor_cf19()))
		_564_v3 = _564_v3
		var _565_v4 bool
		_ = _565_v4
		_565_v4 = true
		var _566_v5 _dafny.CodePoint
		_ = _566_v5
		_566_v5 = _dafny.CodePoint('v')
		var _567_v6 _dafny.Map
		_ = _567_v6
		_567_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(27), _566_v5)
		var _568_v7 _dafny.Map
		_ = _568_v7
		_568_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _565_v4)
		var _569_i1 _dafny.Int
		_ = _569_i1
		_569_i1 = _dafny.Zero
		{
			var _pat_let_tv15 = _565_v4
			_ = _pat_let_tv15
			var _pat_let_tv16 = _564_v3
			_ = _pat_let_tv16
			var _pat_let_tv17 = _564_v3
			_ = _pat_let_tv17
			var _pat_let_tv18 = _565_v4
			_ = _pat_let_tv18
			var _pat_let_tv19 = _565_v4
			_ = _pat_let_tv19
			var _pat_let_tv20 = _565_v4
			_ = _pat_let_tv20
			for func(_source12 D0) bool {
				if _source12.Is_DC1() {
					var _575___mcc_h0 _dafny.Int = _source12.Get_().(D0_DC1).Cf1
					_ = _575___mcc_h0
					var _576___mcc_h1 bool = _source12.Get_().(D0_DC1).Cf2
					_ = _576___mcc_h1
					var _577___mcc_h2 _dafny.Int = _source12.Get_().(D0_DC1).Cf3
					_ = _577___mcc_h2
					var _578___mcc_h3 bool = _source12.Get_().(D0_DC1).Cf4
					_ = _578___mcc_h3
					var _579_cf4 bool = _578___mcc_h3
					_ = _579_cf4
					var _580_cf3 _dafny.Int = _577___mcc_h2
					_ = _580_cf3
					var _581_cf2 bool = _576___mcc_h1
					_ = _581_cf2
					var _582_cf1 _dafny.Int = _575___mcc_h0
					_ = _582_cf1
					return _581_cf2
				} else if _source12.Is_DC2() {
					var _583___mcc_h4 _dafny.Int = _source12.Get_().(D0_DC2).Cf5
					_ = _583___mcc_h4
					var _584___mcc_h5 _dafny.CodePoint = _source12.Get_().(D0_DC2).Cf6
					_ = _584___mcc_h5
					var _585___mcc_h6 _dafny.Int = _source12.Get_().(D0_DC2).Cf7
					_ = _585___mcc_h6
					var _586___mcc_h7 _dafny.Int = _source12.Get_().(D0_DC2).Cf8
					_ = _586___mcc_h7
					var _587___mcc_h8 _dafny.Int = _source12.Get_().(D0_DC2).Cf9
					_ = _587___mcc_h8
					var _588_cf9 _dafny.Int = _587___mcc_h8
					_ = _588_cf9
					var _589_cf8 _dafny.Int = _586___mcc_h7
					_ = _589_cf8
					var _590_cf7 _dafny.Int = _585___mcc_h6
					_ = _590_cf7
					var _591_cf6 _dafny.CodePoint = _584___mcc_h5
					_ = _591_cf6
					var _592_cf5 _dafny.Int = _583___mcc_h4
					_ = _592_cf5
					return _pat_let_tv15
				} else if _source12.Is_DC0() {
					var _593___mcc_h9 _dafny.Map = _source12.Get_().(D0_DC0).Cf0
					_ = _593___mcc_h9
					var _594_cf0 _dafny.Map = _593___mcc_h9
					_ = _594_cf0
					return !_dafny.Companion_Sequence_.Equal(_pat_let_tv16, _pat_let_tv17)
				} else {
					var _595___mcc_h10 D0 = _source12.Get_().(D0_DC3).Cf10
					_ = _595___mcc_h10
					var _596_cf10 D0 = _595___mcc_h10
					_ = _596_cf10
					if _pat_let_tv18 {
						return _pat_let_tv19
					} else {
						return _pat_let_tv20
					}
				}
			}(Companion_Default___.Fm19(_565_v4, (func() _dafny.CodePoint {
				if (_567_v6).Contains((_this).F6()) {
					return (_567_v6).Get((_this).F6()).(_dafny.CodePoint)
				}
				return _566_v5
			})(), _568_v7, true, globalState)) {
				{
					if (_569_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_569_i1 = (_569_i1).Plus(_dafny.One)
					var _570_v8 _dafny.Map
					_ = _570_v8
					_570_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_565_v4, _dafny.IntOfInt64(969))
					var _571_v9 _dafny.Sequence
					_ = _571_v9
					_571_v9 = _dafny.SeqOf(_570_v8)
					_571_v9 = _571_v9
					var _572_v10 *C1
					_ = _572_v10
					var _nw96 *C1 = New_C1_()
					_ = _nw96
					_nw96.Ctor__()
					_572_v10 = _nw96
					var _573_v11 D7
					_ = _573_v11
					_573_v11 = Companion_D7_.Create_DC18_(!(_565_v4), !(_565_v4), (_this).F6(), true)
					(globalState).F2 = (_dafny.Zero).Minus((_573_v11).Dtor_cf31())
					var _574_v12 *C1
					_ = _574_v12
					var _nw97 *C1 = New_C1_()
					_ = _nw97
					_nw97.Ctor__()
					_574_v12 = _nw97
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		var _597_v13 _dafny.Array
		_ = _597_v13
		var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(14))
		_ = _nw98
		_597_v13 = _nw98
		var _598_v14 _dafny.Sequence
		_ = _598_v14
		_598_v14 = _dafny.UnicodeSeqOfUtf8Bytes("ftx")
		var _599_v15 _dafny.Sequence
		_ = _599_v15
		_599_v15 = _dafny.SeqOf(_598_v14)
		var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_597_v13), 0))
		_ = _index82
		(_597_v13).ArraySet1(_dafny.Companion_Sequence_.Update(_599_v15, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_599_v15).Cardinality()))).Uint32(), _598_v14), (_index82).Int())
		var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(672), _dafny.ArrayLen((_597_v13), 0))
		_ = _index83
		(_597_v13).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(359))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}((func(_600_v14 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
			return func(_601_i2 _dafny.Int) _dafny.Sequence {
				return _600_v14
			}
		})(_598_v14))), (_index83).Int())
		var _602_v16 *C0
		_ = _602_v16
		var _nw99 *C0 = New_C0_()
		_ = _nw99
		_nw99.Ctor__(_565_v4)
		_602_v16 = _nw99
		var _hi4 _dafny.Int = ((_this).F9()).Plus((_this).F9())
		_ = _hi4
		for _603_i3 := (_this).F9(); _603_i3.Cmp(_hi4) < 0; _603_i3 = _603_i3.Plus(_dafny.One) {
			_598_v14 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ubvcc"), _598_v14)
			r0 = (_this).F6()
			var _604_v17 _dafny.Array
			_ = _604_v17
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_15
			var _nw100 _dafny.Array
			_ = _nw100
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw100 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) _dafny.Sequence = (func(_605_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
					return func(_606_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-871))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg25 _dafny.Int) interface{} {
								return coer25(arg25)
							}
						}((func(_607_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_608_i5 _dafny.Int) _dafny.CodePoint {
								return _607_v5
							}
						})(_605_v5)))
					}
				})(_566_v5)
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw100 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw100).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw100).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_604_v17 = _nw100
			var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_604_v17), 0))
			_ = _index84
			(_604_v17).ArraySet1(_598_v14, (_index84).Int())
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_604_v17), 0))
			_ = _index85
			(_604_v17).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_598_v14, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(190))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg26 _dafny.Int) interface{} {
					return coer26(arg26)
				}
			}((func(_609_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_610_i6 _dafny.Int) _dafny.CodePoint {
					return _609_v5
				}
			})(_566_v5)))), (_index85).Int())
			var _611_v18 _dafny.Set
			_ = _611_v18
			_611_v18 = _dafny.SetOf(_603_i3, _603_i3, Companion_Default___.SafeDivisionInt(_603_i3, _603_i3))
			var _rhs69 bool = (func() bool {
				if (_602_v16).F13() {
					return !((_602_v16).F13())
				}
				return (_602_v16).F13()
			})()
			_ = _rhs69
			var _rhs70 _dafny.Set = _611_v18
			_ = _rhs70
			var _rhs71 _dafny.Array = _560_v0
			_ = _rhs71
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			_lhs54.F1 = _rhs69
			_611_v18 = _rhs70
			_560_v0 = _rhs71
		}
		r0 = (func() _dafny.Int {
			if _565_v4 {
				return (_this).F6()
			}
			return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jlx")).Cardinality()))
		})()
		var _612_v20 _dafny.Set
		_ = _612_v20
		_612_v20 = _dafny.SetOf(_dafny.CodePoint('g'), _566_v5)
		var _613_v21 _dafny.Sequence
		_ = _613_v21
		_613_v21 = _dafny.SeqOf(_612_v20, _612_v20)
		r1 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(712))).Uint32(), func(coer27 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_614_v5 _dafny.CodePoint) func(_dafny.Int) _dafny.Set {
			return func(_615_i7 _dafny.Int) _dafny.Set {
				return func() _dafny.Set {
					var _coll38 = _dafny.NewBuilder()
					_ = _coll38
					for _iter41 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v5, (_this).F6())).Keys().Elements()); ; {
						_compr_38, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _616_v19 _dafny.CodePoint
						_616_v19 = interface{}(_compr_38).(_dafny.CodePoint)
						if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_614_v5, (_this).F6())).Contains(_616_v19) {
							_coll38.Add(_616_v19)
						}
					}
					return _coll38.ToSet()
				}()
			}
		})(_566_v5))), _613_v21)
		var _617_v22 _dafny.Set
		_ = _617_v22
		_617_v22 = _dafny.SetOf(true)
		var _618_v23 _dafny.Map
		_ = _618_v23
		_618_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _617_v22)
		r2 = (((_this).F9()).Cmp((_567_v6).Cardinality()) <= 0) == (Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_618_v23), (_602_v16).F13(), (_this).F9(), (_602_v16).F13(), globalState))
		return r0, r1, r2
	}
}
func (_this *C3) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *C3) F10() _dafny.Sequence {
	{
		return _this._f10
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f6 _dafny.Int
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f6 = _dafny.Zero
	return &_this
}

type CompanionStruct_C4_ struct {
}

var Companion_C4_ = CompanionStruct_C4_{}

func (_this *C4) Equals(other *C4) bool {
	return _this == other
}

func (_this *C4) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C4)
	return ok && _this.Equals(other)
}

func (*C4) String() string {
	return "_module.C4"
}

func Type_C4_() _dafny.TypeDescriptor {
	return type_C4_{}
}

type type_C4_ struct {
}

func (_this type_C4_) Default() interface{} {
	return (*C4)(nil)
}

func (_this type_C4_) String() string {
	return "main.C4"
}
func (_this *C4) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F6() _dafny.Int {
	return _this._f6
}
func (_this *C4) Ctor__(f6 _dafny.Int) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C4) Fm11(globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("tp"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-574))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}(func(_619_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('w')
		})))).Intersection(_dafny.SetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(622))).Uint32(), func(coer29 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg29 _dafny.Int) interface{} {
				return coer29(arg29)
			}
		}(func(_620_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('i')
		}))))).IsDisjointFrom((func() _dafny.Set {
			var _coll39 = _dafny.NewBuilder()
			_ = _coll39
			for _iter42 := _dafny.Iterate((_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("wcshiaqcw"))).Elements()); ; {
				_compr_39, _ok42 := _iter42()
				if !_ok42 {
					break
				}
				var _621_v0 _dafny.Sequence
				_621_v0 = interface{}(_compr_39).(_dafny.Sequence)
				if (_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("wcshiaqcw"))).Contains(_621_v0) {
					_coll39.Add(_621_v0)
				}
			}
			return _coll39.ToSet()
		}()).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("qhbibqr"), _dafny.UnicodeSeqOfUtf8Bytes("jiutwlf"), _dafny.UnicodeSeqOfUtf8Bytes("tmcd"))))
	}
}
func (_this *C4) M0(globalState *GlobalState) (_dafny.Int, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _622_v0 _dafny.Array
		_ = _622_v0
		var _nw101 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
		_ = _nw101
		_622_v0 = _nw101
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
		_ = _index86
		(_622_v0).ArraySet1(false, (_index86).Int())
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
		_ = _index87
		(_622_v0).ArraySet1(false, (_index87).Int())
		var _623_v1 _dafny.CodePoint
		_ = _623_v1
		_623_v1 = _dafny.CodePoint('c')
		var _624_v2 _dafny.Sequence
		_ = _624_v2
		_624_v2 = _dafny.SeqOf(_623_v1, _dafny.CodePoint('a'))
		var _625_v3 _dafny.Array
		_ = _625_v3
		var _nwElement0_22 _dafny.CodePoint = (func() _dafny.CodePoint {
			if (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool) {
				return _623_v1
			}
			return _623_v1
		})()
		_ = _nwElement0_22
		var _nw102 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(28))
		_ = _nw102
		(_nw102).ArraySet1CodePoint(_nwElement0_22, 0)
		(_nw102).ArraySet1CodePoint((_624_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-956), _dafny.IntOfUint32((_624_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), 1)
		(_nw102).ArraySet1CodePoint(_623_v1, 2)
		(_nw102).ArraySet1CodePoint(_623_v1, 3)
		(_nw102).ArraySet1CodePoint(_623_v1, 4)
		(_nw102).ArraySet1CodePoint(_623_v1, 5)
		(_nw102).ArraySet1CodePoint(_623_v1, 6)
		(_nw102).ArraySet1CodePoint(_623_v1, 7)
		(_nw102).ArraySet1CodePoint(_623_v1, 8)
		(_nw102).ArraySet1CodePoint((_624_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(953), _dafny.IntOfUint32((_624_v2).Cardinality()))).Uint32()).(_dafny.CodePoint), 9)
		(_nw102).ArraySet1CodePoint(_623_v1, 10)
		(_nw102).ArraySet1CodePoint(_623_v1, 11)
		(_nw102).ArraySet1CodePoint(_623_v1, 12)
		(_nw102).ArraySet1CodePoint(_623_v1, 13)
		(_nw102).ArraySet1CodePoint(_623_v1, 14)
		(_nw102).ArraySet1CodePoint(_623_v1, 15)
		(_nw102).ArraySet1CodePoint(_623_v1, 16)
		(_nw102).ArraySet1CodePoint(_623_v1, 17)
		(_nw102).ArraySet1CodePoint(_623_v1, 18)
		(_nw102).ArraySet1CodePoint(_623_v1, 19)
		(_nw102).ArraySet1CodePoint(_623_v1, 20)
		(_nw102).ArraySet1CodePoint((func() _dafny.CodePoint {
			if (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool) {
				return _dafny.CodePoint('y')
			}
			return _623_v1
		})(), 21)
		(_nw102).ArraySet1CodePoint(_623_v1, 22)
		(_nw102).ArraySet1CodePoint(_623_v1, 23)
		(_nw102).ArraySet1CodePoint(Companion_Default___.Fm12(globalState), 24)
		(_nw102).ArraySet1CodePoint(_623_v1, 25)
		(_nw102).ArraySet1CodePoint(_dafny.CodePoint('s'), 26)
		(_nw102).ArraySet1CodePoint(_623_v1, 27)
		_625_v3 = _nw102
		var _626_v4 _dafny.Set
		_ = _626_v4
		_626_v4 = _dafny.SetOf((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
		var _627_v5 D0
		_ = _627_v5
		_627_v5 = Companion_D0_.Create_DC1_((_this).F6(), ((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)) && ((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)), (_this).F6(), (_626_v4).IsProperSubsetOf(_626_v4))
		var _628_v6 D3
		_ = _628_v6
		_628_v6 = Companion_D3_.Create_DC7_(_624_v2)
		var _pat_let_tv21 = _624_v2
		_ = _pat_let_tv21
		var _629_v7 D0
		_ = _629_v7
		_629_v7 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(542), _623_v1, (_this).F6(), (_this).F6(), _dafny.IntOfUint32(((func(_pat_let18_0 D3) D3 {
			return func(_630_dt__update__tmp_h0 D3) D3 {
				return func(_pat_let19_0 _dafny.Sequence) D3 {
					return func(_631_dt__update_hcf13_h0 _dafny.Sequence) D3 {
						return Companion_D3_.Create_DC7_(_631_dt__update_hcf13_h0)
					}(_pat_let19_0)
				}(_pat_let_tv21)
			}(_pat_let18_0)
		}(_628_v6)).Dtor_cf13()).Cardinality()))
		var _rhs72 _dafny.Array = (func() _dafny.Array {
			if (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool) {
				return _625_v3
			}
			return _625_v3
		})()
		_ = _rhs72
		var _rhs73 _dafny.Int = (_629_v7).Dtor_cf8()
		_ = _rhs73
		var _rhs74 D0 = Companion_Default___.Fm13(_623_v1, _624_v2, globalState)
		_ = _rhs74
		var _lhs55 *GlobalState = globalState
		_ = _lhs55
		_625_v3 = _rhs72
		_lhs55.F2 = _rhs73
		_627_v5 = _rhs74
		var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
		_ = _index88
		(_622_v0).ArraySet1((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_index88).Int())
		var _hi5 _dafny.Int = (_this).F6()
		_ = _hi5
		for _632_i0 := (_this).F6(); _632_i0.Cmp(_hi5) < 0; _632_i0 = _632_i0.Plus(_dafny.One) {
			if (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool) {
				var _633_v8 D0
				_ = _633_v8
				_633_v8 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_632_i0, _626_v4))
				r2 = !(!(Companion_Default___.Fm1(_633_v8, (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), _632_i0, !((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)), globalState)) || ((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))) || (!((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)))
				var _634_v9 _dafny.Array
				_ = _634_v9
				var _len0_16 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_16
				var _nw103 _dafny.Array
				_ = _nw103
				if _len0_16.Cmp(_dafny.Zero) == 0 {
					_nw103 = _dafny.NewArray(_len0_16)
				} else {
					var _init16 func(_dafny.Int) _dafny.Int = (func(_635_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_636_i1 _dafny.Int) _dafny.Int {
							return (_636_i1).Plus(_635_i0)
						}
					})(_632_i0)
					_ = _init16
					var _element0_16 = _init16(_dafny.Zero)
					_ = _element0_16
					_nw103 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
					(_nw103).ArraySet1(_element0_16, 0)
					var _nativeLen0_16 = (_len0_16).Int()
					_ = _nativeLen0_16
					for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
						(_nw103).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
					}
				}
				_634_v9 = _nw103
				var _len0_17 _dafny.Int = _dafny.IntOfInt64(14)
				_ = _len0_17
				var _nw104 _dafny.Array
				_ = _nw104
				if _len0_17.Cmp(_dafny.Zero) == 0 {
					_nw104 = _dafny.NewArray(_len0_17)
				} else {
					var _init17 func(_dafny.Int) _dafny.Int = func(_637_i2 _dafny.Int) _dafny.Int {
						return (_637_i2).Minus((_this).F6())
					}
					_ = _init17
					var _element0_17 = _init17(_dafny.Zero)
					_ = _element0_17
					_nw104 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
					(_nw104).ArraySet1(_element0_17, 0)
					var _nativeLen0_17 = (_len0_17).Int()
					_ = _nativeLen0_17
					for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
						(_nw104).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
					}
				}
				_634_v9 = _nw104
				_623_v1 = (_629_v7).Dtor_cf6()
				var _638_v11 _dafny.Array
				_ = _638_v11
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_18
				var _nw105 _dafny.Array
				_ = _nw105
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw105 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Set = (func(_639_v0 _dafny.Array) func(_dafny.Int) _dafny.Set {
						return func(_640_i3 _dafny.Int) _dafny.Set {
							return func() _dafny.Set {
								var _coll40 = _dafny.NewBuilder()
								_ = _coll40
								for _iter43 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_639_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_639_v0), 0))).Int()).(bool), (_639_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_639_v0), 0))).Int()).(bool)))).Elements()); ; {
									_compr_40, _ok43 := _iter43()
									if !_ok43 {
										break
									}
									var _641_v10 _dafny.Map
									_641_v10 = interface{}(_compr_40).(_dafny.Map)
									if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_639_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_639_v0), 0))).Int()).(bool), (_639_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_639_v0), 0))).Int()).(bool))), _641_v10) {
										_coll40.Add(_641_v10)
									}
								}
								return _coll40.ToSet()
							}()
						}
					})(_622_v0)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw105 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw105).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw105).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_638_v11 = _nw105
				var _642_v12 _dafny.Map
				_ = _642_v12
				_642_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
				var _643_v13 _dafny.Sequence
				_ = _643_v13
				_643_v13 = _dafny.SeqOf(_dafny.SetOf(_642_v12, _642_v12), _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), true)))
				var _644_v14 _dafny.MultiSet
				_ = _644_v14
				_644_v14 = _dafny.MultiSetOf((_this).F6())
				var _645_v15 _dafny.Sequence
				_ = _645_v15
				_645_v15 = _dafny.SeqOf((_644_v14).Cardinality(), _632_i0)
				var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_638_v11), 0))
				_ = _index89
				(_638_v11).ArraySet1(((_643_v13).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_643_v13).Cardinality()))).Uint32()).(_dafny.Set)).Union(_dafny.SetOf(Companion_Default___.Fm14(_dafny.IntOfInt64(702), _632_i0, _dafny.IntOfUint32((_645_v15).Cardinality()), globalState))), (_index89).Int())
				var _646_v16 _dafny.Map
				_ = _646_v16
				_646_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), _642_v12)
				var _647_v17 _dafny.Map
				_ = _647_v17
				_647_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Map {
					if (_646_v16).Contains((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)) {
						return (_646_v16).Get((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)).(_dafny.Map)
					}
					return _642_v12
				})(), _632_i0)
				var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_638_v11), 0))
				_ = _index90
				var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
				_ = _index91
				var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
				_ = _index92
				var _rhs75 _dafny.Set = func() _dafny.Set {
					var _coll41 = _dafny.NewBuilder()
					_ = _coll41
					for _iter44 := _dafny.Iterate(((_647_v17).Merge(_647_v17)).Keys().Elements()); ; {
						_compr_41, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _648_v18 _dafny.Map
						_648_v18 = interface{}(_compr_41).(_dafny.Map)
						if ((_647_v17).Merge(_647_v17)).Contains(_648_v18) {
							_coll41.Add(_648_v18)
						}
					}
					return _coll41.ToSet()
				}()
				_ = _rhs75
				var _rhs76 bool = (_this).Fm11(globalState)
				_ = _rhs76
				var _rhs77 bool = (_627_v5).Dtor_cf2()
				_ = _rhs77
				var _rhs78 bool = (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)
				_ = _rhs78
				var _lhs56 _dafny.Array = _638_v11
				_ = _lhs56
				var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(825), _dafny.ArrayLen((_638_v11), 0))
				_ = _lhs57
				var _lhs58 _dafny.Array = _622_v0
				_ = _lhs58
				var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
				_ = _lhs59
				var _lhs60 _dafny.Array = _622_v0
				_ = _lhs60
				var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))
				_ = _lhs61
				var _lhs62 *GlobalState = globalState
				_ = _lhs62
				(_lhs56).ArraySet1(_rhs75, (_lhs57).Int())
				(_lhs58).ArraySet1(_rhs76, (_lhs59).Int())
				(_lhs60).ArraySet1(_rhs77, (_lhs61).Int())
				_lhs62.F1 = _rhs78
				var _649_v19 _dafny.Array
				_ = _649_v19
				var _nw106 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
				_ = _nw106
				_649_v19 = _nw106
				var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_649_v19), 0))
				_ = _index93
				(_649_v19).ArraySet1(_624_v2, (_index93).Int())
				var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(860), _dafny.ArrayLen((_649_v19), 0))
				_ = _index94
				(_649_v19).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_624_v2, _624_v2), (_index94).Int())
			} else {
				var _650_v20 _dafny.Sequence
				_ = _650_v20
				_650_v20 = _dafny.SeqOf(false, (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
				var _651_v21 _dafny.Sequence
				_ = _651_v21
				_651_v21 = _dafny.SeqOf(Companion_Default___.Fm15(_632_i0, true, globalState), _dafny.IntOfUint32((_650_v20).Cardinality()), (_this).F6(), _632_i0, (_this).F6())
				var _652_v22 _dafny.Int
				_ = _652_v22
				var _653_v23 D1
				_ = _653_v23
				var _654_v24 _dafny.Int
				_ = _654_v24
				var _655_v25 _dafny.Set
				_ = _655_v25
				var _out0 _dafny.Int
				_ = _out0
				var _out1 D1
				_ = _out1
				var _out2 _dafny.Int
				_ = _out2
				var _out3 _dafny.Set
				_ = _out3
				_out0, _out1, _out2, _out3 = (_this).M3(_dafny.IntOfUint32((_651_v21).Cardinality()), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), _632_i0, (_632_i0).Cmp(_dafny.IntOfInt64(231)) == 0, globalState)
				_652_v22 = _out0
				_653_v23 = _out1
				_654_v24 = _out2
				_655_v25 = _out3
				var _656_v26 _dafny.Map
				_ = _656_v26
				_656_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v24, (_this).F6())
				var _657_v27 _dafny.Set
				_ = _657_v27
				_657_v27 = _dafny.SetOf((func() _dafny.Int {
					if (_656_v26).Contains(_dafny.IntOfInt64(-238)) {
						return (_656_v26).Get(_dafny.IntOfInt64(-238)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_651_v21).Cardinality())
				})())
				var _rhs79 _dafny.Int = ((func() _dafny.Set {
					if (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool) {
						return (_657_v27).Union(_dafny.SetOf(_dafny.IntOfInt64(562)))
					}
					return _dafny.SetOf((_this).F6(), _654_v24)
				})()).Cardinality()
				_ = _rhs79
				var _rhs80 bool = (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)
				_ = _rhs80
				var _rhs81 bool = (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)
				_ = _rhs81
				var _rhs82 _dafny.Int = _dafny.IntOfInt64(370)
				_ = _rhs82
				var _lhs63 *GlobalState = globalState
				_ = _lhs63
				var _lhs64 *GlobalState = globalState
				_ = _lhs64
				var _lhs65 *GlobalState = globalState
				_ = _lhs65
				r0 = _rhs79
				_lhs63.F1 = _rhs80
				_lhs64.F1 = _rhs81
				_lhs65.F2 = _rhs82
				var _658_v28 _dafny.MultiSet
				_ = _658_v28
				_658_v28 = _dafny.MultiSetOf((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
				var _659_v29 _dafny.Array
				_ = _659_v29
				var _nwElement0_23 _dafny.Int = (_dafny.Zero).Minus(_654_v24)
				_ = _nwElement0_23
				var _nw107 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(12))
				_ = _nw107
				(_nw107).ArraySet1(_nwElement0_23, 0)
				(_nw107).ArraySet1(Companion_Default___.Fm15(_654_v24, (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), globalState), 1)
				(_nw107).ArraySet1((_657_v27).Cardinality(), 2)
				(_nw107).ArraySet1((_this).F6(), 3)
				(_nw107).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_651_v21, _651_v21)).Cardinality()), 4)
				(_nw107).ArraySet1((func() _dafny.Int {
					if (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool) {
						return _652_v22
					}
					return (_this).F6()
				})(), 5)
				(_nw107).ArraySet1(Companion_Default___.SafeModuloInt(_654_v24, _632_i0), 6)
				(_nw107).ArraySet1((func() _dafny.Int {
					if (_658_v28).Contains((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)) {
						return (_658_v28).Multiplicity((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
					}
					return _dafny.IntOfInt64(-667)
				})(), 7)
				(_nw107).ArraySet1(_dafny.IntOfInt64(962), 8)
				(_nw107).ArraySet1(_654_v24, 9)
				(_nw107).ArraySet1(_652_v22, 10)
				(_nw107).ArraySet1((_dafny.Zero).Minus(_632_i0), 11)
				_659_v29 = _nw107
				var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_659_v29), 0))
				_ = _index95
				(_659_v29).ArraySet1(Companion_Default___.SafeModuloInt((_this).F6(), _632_i0), (_index95).Int())
				var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_659_v29), 0))
				_ = _index96
				(_659_v29).ArraySet1(Companion_Default___.SafeDivisionInt((_652_v22).Minus((_this).F6()), (_dafny.Zero).Minus((_627_v5).Dtor_cf1())), (_index96).Int())
				_623_v1 = (_624_v2).Select((Companion_Default___.SafeIndex(_632_i0, _dafny.IntOfUint32((_624_v2).Cardinality()))).Uint32()).(_dafny.CodePoint)
				(globalState).F5 = (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)
			}
			var _660_v30 _dafny.Map
			_ = _660_v30
			_660_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_this).F6())
			var _661_v31 _dafny.Map
			_ = _661_v31
			_661_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
			var _662_v32 _dafny.Map
			_ = _662_v32
			_662_v32 = _661_v31
			_660_v30 = (_660_v30).Update((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_dafny.Zero).Minus((_661_v31).Cardinality()))
			_623_v1 = _623_v1
			var _663_v33 _dafny.Map
			_ = _663_v33
			_663_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_627_v5, _dafny.IntOfInt64(-138))
			_663_v33 = (_663_v33).Update(_627_v5, (_629_v7).Dtor_cf7())
		}
		var _664_v34 _dafny.Map
		_ = _664_v34
		_664_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
		var _665_v35 _dafny.Map
		_ = _665_v35
		_665_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_664_v34, (_this).Fm11(globalState))
		_665_v35 = (_665_v35).Update((_664_v34).Update((_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool)), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
		var _666_v36 *C2
		_ = _666_v36
		var _nw108 *C2 = New_C2_()
		_ = _nw108
		_nw108.Ctor__(_dafny.UnicodeSeqOfUtf8Bytes("iyanqn"), (_622_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(880), _dafny.ArrayLen((_622_v0), 0))).Int()).(bool))
		_666_v36 = _nw108
		r0 = (_this).F6()
		r1 = true
		var _667_v37 _dafny.MultiSet
		_ = _667_v37
		_667_v37 = _dafny.MultiSetOf(_666_v36.F12)
		var _668_v38 D7
		_ = _668_v38
		_668_v38 = Companion_D7_.Create_DC19_(_667_v37, (_this).F6())
		r2 = (_667_v37).Equals((_668_v38).Dtor_cf33())
		return r0, r1, r2
	}
}
func (_this *C4) M3(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) (_dafny.Int, D1, _dafny.Int, _dafny.Set) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 D1 = Companion_D1_.Default()
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Set = _dafny.EmptySet
		_ = r3
		(globalState).F5 = p3
		var _hi6 _dafny.Int = (_this).F6()
		_ = _hi6
		for _669_i0 := p0; _669_i0.Cmp(_hi6) < 0; _669_i0 = _669_i0.Plus(_dafny.One) {
			var _670_v0 _dafny.Array
			_ = _670_v0
			var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(14))
			_ = _nw109
			_670_v0 = _nw109
			var _671_v1 _dafny.Sequence
			_ = _671_v1
			_671_v1 = _dafny.UnicodeSeqOfUtf8Bytes("gpng")
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_670_v0), 0))
			_ = _index97
			(_670_v0).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_671_v1).Cardinality())), (_index97).Int())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_670_v0), 0))
			_ = _index98
			(_670_v0).ArraySet1((_dafny.Zero).Minus(_669_i0), (_index98).Int())
			var _672_v2 _dafny.CodePoint
			_ = _672_v2
			_672_v2 = _dafny.CodePoint('p')
			var _673_v3 _dafny.Map
			_ = _673_v3
			_673_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_672_v2, (_dafny.SetOf(p3)).Cardinality())
			var _674_v4 _dafny.Sequence
			_ = _674_v4
			_674_v4 = _dafny.SeqOf(p3, false)
			var _675_v5 _dafny.Map
			_ = _675_v5
			_675_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_674_v4).Cardinality()), _dafny.IntOfInt64(-529))
			var _676_v6 D7
			_ = _676_v6
			_676_v6 = Companion_D7_.Create_DC18_(p3, p1, (_675_v5).Cardinality(), true)
			var _677_v7 D3
			_ = _677_v7
			_677_v7 = Companion_D3_.Create_DC7_(_671_v1)
			var _678_v8 D5
			_ = _678_v8
			_678_v8 = Companion_D5_.Create_DC13_((_676_v6).Dtor_cf30(), p3, (_677_v7).Dtor_cf13(), p0)
			var _679_v9 *C3
			_ = _679_v9
			var _nw110 *C3 = New_C3_()
			_ = _nw110
			_nw110.Ctor__(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("poee")).Cardinality()), Companion_Default___.Fm25(p2, _673_v3, _672_v2, globalState), (_678_v8).Dtor_cf24())
			_679_v9 = _nw110
			var _680_v10 _dafny.Map
			_ = _680_v10
			_680_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(688), _674_v4)
			_680_v10 = (_680_v10).Update((_670_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(959), _dafny.ArrayLen((_670_v0), 0))).Int()).(_dafny.Int), _674_v4)
			var _681_v11 _dafny.MultiSet
			_ = _681_v11
			_681_v11 = _dafny.MultiSetOf(p3, p3, p1, p3)
			var _682_v12 _dafny.Sequence
			_ = _682_v12
			_682_v12 = _dafny.SeqOf(_681_v11)
			var _683_v13 _dafny.Array
			_ = _683_v13
			var _nwElement0_24 bool = true
			_ = _nwElement0_24
			var _nw111 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(11))
			_ = _nw111
			(_nw111).ArraySet1(_nwElement0_24, 0)
			(_nw111).ArraySet1(((_682_v12).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_682_v12).Cardinality()))).Uint32()).(_dafny.MultiSet)).IsDisjointFrom(_dafny.MultiSetFromSeq(_674_v4)), 1)
			(_nw111).ArraySet1(p3, 2)
			(_nw111).ArraySet1(p1, 3)
			(_nw111).ArraySet1(p1, 4)
			(_nw111).ArraySet1((p1) || (p3), 5)
			(_nw111).ArraySet1(p1, 6)
			(_nw111).ArraySet1(p3, 7)
			(_nw111).ArraySet1(p3, 8)
			(_nw111).ArraySet1((func() bool {
				if p1 {
					return !(false)
				}
				return p3
			})(), 9)
			(_nw111).ArraySet1(_dafny.Companion_Sequence_.Equal(_671_v1, _671_v1), 10)
			_683_v13 = _nw111
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_683_v13), 0))
			_ = _index99
			(_683_v13).ArraySet1(p1, (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_683_v13), 0))
			_ = _index100
			(_683_v13).ArraySet1(p1, (_index100).Int())
		}
		if true {
			(globalState).F5 = p1
			var _684_v14 _dafny.Array
			_ = _684_v14
			var _nw112 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(22))
			_ = _nw112
			_684_v14 = _nw112
			var _685_v15 _dafny.Sequence
			_ = _685_v15
			_685_v15 = _dafny.SeqOf(_684_v14, _684_v14, _684_v14)
			_685_v15 = _dafny.Companion_Sequence_.Concatenate(_685_v15, _685_v15)
			var _686_v16 D0
			_ = _686_v16
			_686_v16 = Companion_D0_.Create_DC1_(Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(251)), p1, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("fr")).Cardinality()), p1)
			var _pat_let_tv22 = p0
			_ = _pat_let_tv22
			_686_v16 = func(_pat_let20_0 D0) D0 {
				return func(_687_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let21_0 _dafny.Int) D0 {
						return func(_688_dt__update_hcf1_h0 _dafny.Int) D0 {
							return func(_pat_let22_0 bool) D0 {
								return func(_689_dt__update_hcf4_h0 bool) D0 {
									return func(_pat_let23_0 _dafny.Int) D0 {
										return func(_690_dt__update_hcf3_h0 _dafny.Int) D0 {
											return Companion_D0_.Create_DC1_(_688_dt__update_hcf1_h0, (_687_dt__update__tmp_h0).Dtor_cf2(), _690_dt__update_hcf3_h0, _689_dt__update_hcf4_h0)
										}(_pat_let23_0)
									}((_this).F6())
								}(_pat_let22_0)
							}(true)
						}(_pat_let21_0)
					}(_pat_let_tv22)
				}(_pat_let20_0)
			}(_686_v16)
			var _691_v17 _dafny.Array
			_ = _691_v17
			var _nw113 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw113
			_691_v17 = _nw113
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_691_v17), 0))
			_ = _index101
			(_691_v17).ArraySet1((_dafny.IntOfInt64(-761)).Times((_this).F6()), (_index101).Int())
			var _692_v18 _dafny.Map
			_ = _692_v18
			_692_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p2), _684_v14)
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_691_v17), 0))
			_ = _index102
			var _rhs83 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(362), (p2).Minus(_dafny.IntOfInt64(667)))
			_ = _rhs83
			var _rhs84 _dafny.Array = (func() _dafny.Array {
				if (_692_v18).Contains((_this).F6()) {
					return (_692_v18).Get((_this).F6()).(_dafny.Array)
				}
				return _684_v14
			})()
			_ = _rhs84
			var _rhs85 _dafny.Array = _684_v14
			_ = _rhs85
			var _rhs86 _dafny.Int = (_this).F6()
			_ = _rhs86
			var _lhs66 _dafny.Array = _691_v17
			_ = _lhs66
			var _lhs67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_691_v17), 0))
			_ = _lhs67
			var _lhs68 *GlobalState = globalState
			_ = _lhs68
			(_lhs66).ArraySet1(_rhs83, (_lhs67).Int())
			_684_v14 = _rhs84
			_684_v14 = _rhs85
			_lhs68.F2 = _rhs86
			var _693_v19 _dafny.Sequence
			_ = _693_v19
			_693_v19 = _dafny.SeqOf(p0)
			var _694_v20 _dafny.Map
			_ = _694_v20
			_694_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, p3)
			var _695_v21 *C3
			_ = _695_v21
			var _nw114 *C3 = New_C3_()
			_ = _nw114
			_nw114.Ctor__(p0, _693_v19, (_dafny.IntOfUint32(((Companion_D1_.Create_DC4_(_693_v19)).Dtor_cf11()).Cardinality())).Plus((_694_v20).Cardinality()))
			_695_v21 = _nw114
		} else {
			var _696_v22 _dafny.Map
			_ = _696_v22
			_696_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			if !((_696_v22).Merge(_696_v22)).Equals(_696_v22) {
				var _697_v23 _dafny.Sequence
				_ = _697_v23
				_697_v23 = _dafny.SeqOf(p0)
				var _698_v24 _dafny.Sequence
				_ = _698_v24
				_698_v24 = _dafny.UnicodeSeqOfUtf8Bytes("aofqlqta")
				var _699_v25 T0
				_ = _699_v25
				var _nw115 *C3 = New_C3_()
				_ = _nw115
				_nw115.Ctor__((_this).F6(), _697_v23, _dafny.IntOfUint32((_698_v24).Cardinality()))
				_699_v25 = _nw115
				_699_v25 = _699_v25
				var _700_v26 _dafny.Map
				_ = _700_v26
				_700_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf((_696_v22).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(256))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg30 _dafny.Int) interface{} {
						return coer30(arg30)
					}
				}((func(_701_p0 _dafny.Int, _702_p3 bool) func(_dafny.Int) _dafny.Map {
					return func(_703_i1 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_701_p0, _702_p3)
					}
				})(p0, p3)))).Cardinality()))).Cardinality()), (_699_v25).F6()), p0)
				_700_v26 = (_700_v26).Merge(_700_v26)
				var _704_v27 _dafny.Map
				_ = _704_v27
				_704_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(p0, (_699_v25).F6()), !(p3) || (p3))
				(globalState).F5 = (func() bool {
					if (_704_v27).Contains((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
						if !(p3) {
							return _698_v24
						}
						return _698_v24
					})()).Cardinality()))) {
						return (_704_v27).Get((_dafny.Zero).Minus(_dafny.IntOfUint32(((func() _dafny.Sequence {
							if !(p3) {
								return _698_v24
							}
							return _698_v24
						})()).Cardinality()))).(bool)
					}
					return (_dafny.IntOfUint32((_698_v24).Cardinality())).Cmp(p2) < 0
				})()
				var _705_v28 _dafny.Set
				_ = _705_v28
				_705_v28 = _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("cdiwerm"))
				var _706_v29 _dafny.Sequence
				_ = _706_v29
				_706_v29 = _dafny.SeqOf(_698_v24, _698_v24)
				var _707_v30 _dafny.Set
				_ = _707_v30
				_707_v30 = _dafny.SetOf(_dafny.IntOfInt64(-403), (_this).F6(), ((_705_v28).Cardinality()).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_706_v29).Cardinality()), Companion_Default___.Fm24((_699_v25).F6(), globalState))).Cardinality())))
				_707_v30 = _707_v30
				(globalState).F5 = p1
			} else {
				var _708_v31 _dafny.Array
				_ = _708_v31
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw116
				_708_v31 = _nw116
				var _709_v32 _dafny.Map
				_ = _709_v32
				_709_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(Companion_Default___.Fm24(_dafny.IntOfInt64(811), globalState)), _708_v31)
				var _710_v33 _dafny.Sequence
				_ = _710_v33
				_710_v33 = _dafny.UnicodeSeqOfUtf8Bytes("dmvccnxx")
				var _711_v34 _dafny.Sequence
				_ = _711_v34
				_711_v34 = _dafny.SeqOf((_this).F6(), _dafny.IntOfUint32((_710_v33).Cardinality()))
				var _712_v35 D8
				_ = _712_v35
				_712_v35 = Companion_D8_.Create_DC20_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_711_v34, _708_v31))
				var _713_v36 D1
				_ = _713_v36
				_713_v36 = Companion_D1_.Create_DC4_(_711_v34)
				var _714_v37 _dafny.Array
				_ = _714_v37
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw117
				_714_v37 = _nw117
				var _715_v38 _dafny.Map
				_ = _715_v38
				_715_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_714_v37, _709_v32)
				var _pat_let_tv23 = p2
				_ = _pat_let_tv23
				var _pat_let_tv24 = _708_v31
				_ = _pat_let_tv24
				var _716_v39 _dafny.Array
				_ = _716_v39
				var _nwElement0_25 _dafny.Map = _709_v32
				_ = _nwElement0_25
				var _nw118 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(24))
				_ = _nw118
				(_nw118).ArraySet1(_nwElement0_25, 0)
				(_nw118).ArraySet1(_709_v32, 1)
				(_nw118).ArraySet1((_709_v32).Merge(_709_v32), 2)
				(_nw118).ArraySet1(_709_v32, 3)
				(_nw118).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F6()), _708_v31), 4)
				(_nw118).ArraySet1((func(_pat_let24_0 D8) D8 {
					return func(_717_dt__update__tmp_h1 D8) D8 {
						return func(_pat_let25_0 _dafny.Map) D8 {
							return func(_720_dt__update_hcf35_h0 _dafny.Map) D8 {
								return Companion_D8_.Create_DC20_(_720_dt__update_hcf35_h0)
							}(_pat_let25_0)
						}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-35))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg31 _dafny.Int) interface{} {
								return coer31(arg31)
							}
						}((func(_718_p2 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_719_i2 _dafny.Int) _dafny.Int {
								return _718_p2
							}
						})(_pat_let_tv23))), _pat_let_tv24))
					}(_pat_let24_0)
				}(_712_v35)).Dtor_cf35(), 5)
				(_nw118).ArraySet1(_709_v32, 6)
				(_nw118).ArraySet1(_709_v32, 7)
				(_nw118).ArraySet1(_709_v32, 8)
				(_nw118).ArraySet1(_709_v32, 9)
				(_nw118).ArraySet1(_709_v32, 10)
				(_nw118).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((_713_v36).Dtor_cf11(), (Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32(((_713_v36).Dtor_cf11()).Cardinality()))).Uint32(), p0), _708_v31), 11)
				(_nw118).ArraySet1((func() _dafny.Map {
					if (_715_v38).Contains(_714_v37) {
						return (_715_v38).Get(_714_v37).(_dafny.Map)
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_711_v34, _708_v31)
				})(), 12)
				(_nw118).ArraySet1(_709_v32, 13)
				(_nw118).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_this).F6()), _708_v31), 14)
				(_nw118).ArraySet1(_709_v32, 15)
				(_nw118).ArraySet1(_709_v32, 16)
				(_nw118).ArraySet1(_709_v32, 17)
				(_nw118).ArraySet1((_709_v32).Merge(_709_v32), 18)
				(_nw118).ArraySet1(_709_v32, 19)
				(_nw118).ArraySet1(_709_v32, 20)
				(_nw118).ArraySet1((_712_v35).Dtor_cf35(), 21)
				(_nw118).ArraySet1(_709_v32, 22)
				(_nw118).ArraySet1(_709_v32, 23)
				_716_v39 = _nw118
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_716_v39), 0))
				_ = _index103
				(_716_v39).ArraySet1(_709_v32, (_index103).Int())
				var _721_v40 _dafny.Map
				_ = _721_v40
				_721_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm12(globalState), _709_v32)
				var _722_v41 _dafny.CodePoint
				_ = _722_v41
				_722_v41 = _dafny.CodePoint('v')
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(778), _dafny.ArrayLen((_716_v39), 0))
				_ = _index104
				(_716_v39).ArraySet1(((func() _dafny.Map {
					if (_721_v40).Contains(_722_v41) {
						return (_721_v40).Get(_722_v41).(_dafny.Map)
					}
					return _709_v32
				})()).Update(_dafny.SeqOf((_this).F6()), _708_v31), (_index104).Int())
				var _723_v42 _dafny.Sequence
				_ = _723_v42
				_723_v42 = _dafny.SeqOf(p3)
				(globalState).F0 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if true {
						return _723_v42
					}
					return _dafny.SeqOf(p1)
				})()).Cardinality()), _dafny.IntOfInt64(435))
				(globalState).F1 = p3
				(globalState).F0 = p0
				r0 = p2
			}
			r2 = p0
			var _724_v43 _dafny.MultiSet
			_ = _724_v43
			_724_v43 = _dafny.MultiSetOf(p1, p3)
			var _725_v44 _dafny.Sequence
			_ = _725_v44
			_725_v44 = _dafny.UnicodeSeqOfUtf8Bytes("prilxjot")
			var _726_v45 _dafny.Map
			_ = _726_v45
			_726_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_725_v44, p2)
			_724_v43 = Companion_Default___.Fm30(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xadt"), _725_v44), _726_v45, p3, globalState)
			var _727_v46 _dafny.Array
			_ = _727_v46
			var _len0_19 _dafny.Int = _dafny.IntOfInt64(13)
			_ = _len0_19
			var _nw119 _dafny.Array
			_ = _nw119
			if _len0_19.Cmp(_dafny.Zero) == 0 {
				_nw119 = _dafny.NewArray(_len0_19)
			} else {
				var _init19 func(_dafny.Int) _dafny.Int = func(_728_i3 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_728_i3, (_this).F6())
				}
				_ = _init19
				var _element0_19 = _init19(_dafny.Zero)
				_ = _element0_19
				_nw119 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
				(_nw119).ArraySet1(_element0_19, 0)
				var _nativeLen0_19 = (_len0_19).Int()
				_ = _nativeLen0_19
				for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
					(_nw119).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
				}
			}
			_727_v46 = _nw119
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))
			_ = _index105
			(_727_v46).ArraySet1((_this).F6(), (_index105).Int())
			var _729_v47 _dafny.MultiSet
			_ = _729_v47
			_729_v47 = _dafny.MultiSetOf((_this).F6())
			var _730_v48 _dafny.Map
			_ = _730_v48
			_730_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p3)
			var _731_v49 _dafny.Map
			_ = _731_v49
			_731_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_729_v47).Contains((_730_v48).Cardinality()) {
					return (_729_v47).Multiplicity((_730_v48).Cardinality())
				}
				return (_dafny.Zero).Minus(p0)
			})(), (_this).F6())
			var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))
			_ = _index106
			(_727_v46).ArraySet1((func() _dafny.Int {
				if (_731_v49).Contains((_this).F6()) {
					return (_731_v49).Get((_this).F6()).(_dafny.Int)
				}
				return p2
			})(), (_index106).Int())
			if !(!(p1) || (p3)) {
				var _732_v50 T1
				_ = _732_v50
				var _nw120 *C1 = New_C1_()
				_ = _nw120
				_nw120.Ctor__()
				_732_v50 = _nw120
				var _733_v51 _dafny.MultiSet
				_ = _733_v51
				_733_v51 = _dafny.MultiSetOf(_732_v50)
				var _734_v52 _dafny.Map
				_ = _734_v52
				_734_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_733_v51).Cardinality(), _730_v48)
				var _735_v53 D5
				_ = _735_v53
				_735_v53 = Companion_D5_.Create_DC13_(p1, !(p3), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(478))).Uint32(), func(coer32 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_736_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('s')
				})), (_727_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))).Int()).(_dafny.Int))
				var _737_v54 _dafny.Sequence
				_ = _737_v54
				_737_v54 = _dafny.SeqOf(_735_v53)
				var _738_v56 _dafny.Sequence
				_ = _738_v56
				_738_v56 = _dafny.SeqOf(p2)
				var _739_v57 _dafny.Sequence
				_ = _739_v57
				_739_v57 = _dafny.SeqOf(_dafny.IntOfUint32((_738_v56).Cardinality()), _dafny.IntOfInt64(567))
				_734_v52 = (_734_v52).Update((_dafny.SetOf(_dafny.SeqOf(_735_v53, Companion_D5_.Create_DC13_(p1, p1, _725_v44, (_this).F6())), _737_v54)).Cardinality(), (func() _dafny.Map {
					if p1 {
						return func() _dafny.Map {
							var _coll42 = _dafny.NewMapBuilder()
							_ = _coll42
							for _iter45 := _dafny.Iterate((_739_v57).Elements()); ; {
								_compr_42, _ok45 := _iter45()
								if !_ok45 {
									break
								}
								var _740_v55 _dafny.Int
								_740_v55 = interface{}(_compr_42).(_dafny.Int)
								if _dafny.Companion_Sequence_.Contains(_739_v57, _740_v55) {
									_coll42.Add((_740_v55).Times(p2), false)
								}
							}
							return _coll42.ToMap()
						}()
					}
					return func() _dafny.Map {
						var _coll43 = _dafny.NewMapBuilder()
						_ = _coll43
						for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(476), _dafny.IntOfInt64(297))); ; {
							_compr_43, _ok46 := _iter46()
							if !_ok46 {
								break
							}
							var _741_v58 _dafny.Int
							_741_v58 = interface{}(_compr_43).(_dafny.Int)
							if ((_dafny.IntOfInt64(476)).Cmp(_741_v58) <= 0) && ((_741_v58).Cmp(_dafny.IntOfInt64(297)) < 0) {
								_coll43.Add((_741_v58).Minus(p0), p1)
							}
						}
						return _coll43.ToMap()
					}()
				})())
				var _742_v59 _dafny.Set
				_ = _742_v59
				_742_v59 = _dafny.SetOf(p1)
				_731_v49 = (_731_v49).Update((_742_v59).Cardinality(), p0)
				var _743_v60 D6
				_ = _743_v60
				_743_v60 = Companion_D6_.Create_DC15_((func() _dafny.Set {
					if p3 {
						return _742_v59
					}
					return _dafny.SetOf(true, p3)
				})())
				var _744_v61 _dafny.Sequence
				_ = _744_v61
				_744_v61 = _dafny.SeqOf(p3)
				_743_v60 = Companion_Default___.Fm31((_dafny.Zero).Minus((_727_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))).Int()).(_dafny.Int)), p1, p3, _dafny.MultiSetOf(_744_v61, _744_v61, _744_v61, _744_v61), globalState)
				var _745_v62 _dafny.Array
				_ = _745_v62
				var _nw121 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw121
				_745_v62 = _nw121
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_745_v62), 0))
				_ = _index107
				(_745_v62).ArraySet1(p1, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_745_v62), 0))
				_ = _index108
				(_745_v62).ArraySet1(p1, (_index108).Int())
				var _746_v63 _dafny.Array
				_ = _746_v63
				var _len0_20 _dafny.Int = _dafny.IntOfInt64(9)
				_ = _len0_20
				var _nw122 _dafny.Array
				_ = _nw122
				if _len0_20.Cmp(_dafny.Zero) == 0 {
					_nw122 = _dafny.NewArray(_len0_20)
				} else {
					var _init20 func(_dafny.Int) D5 = (func(_747_v53 D5) func(_dafny.Int) D5 {
						return func(_748_i5 _dafny.Int) D5 {
							return _747_v53
						}
					})(_735_v53)
					_ = _init20
					var _element0_20 = _init20(_dafny.Zero)
					_ = _element0_20
					_nw122 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
					(_nw122).ArraySet1(_element0_20, 0)
					var _nativeLen0_20 = (_len0_20).Int()
					_ = _nativeLen0_20
					for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
						(_nw122).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
					}
				}
				_746_v63 = _nw122
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_746_v63), 0))
				_ = _index109
				(_746_v63).ArraySet1(_735_v53, (_index109).Int())
				var _749_v64 _dafny.CodePoint
				_ = _749_v64
				_749_v64 = _dafny.CodePoint('r')
				var _pat_let_tv25 = _725_v44
				_ = _pat_let_tv25
				var _pat_let_tv26 = _725_v44
				_ = _pat_let_tv26
				var _pat_let_tv27 = _749_v64
				_ = _pat_let_tv27
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_746_v63), 0))
				_ = _index110
				var _rhs87 _dafny.Int = (func() _dafny.Int {
					if p1 {
						return p2
					}
					return (_dafny.SetOf((_745_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_745_v62), 0))).Int()).(bool), !(p1))).Cardinality()
				})()
				_ = _rhs87
				var _rhs88 bool = (_745_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(166), _dafny.ArrayLen((_745_v62), 0))).Int()).(bool)
				_ = _rhs88
				var _rhs89 D5 = func(_pat_let26_0 D5) D5 {
					return func(_750_dt__update__tmp_h2 D5) D5 {
						return func(_pat_let27_0 _dafny.Sequence) D5 {
							return func(_751_dt__update_hcf23_h0 _dafny.Sequence) D5 {
								return Companion_D5_.Create_DC13_((_750_dt__update__tmp_h2).Dtor_cf21(), (_750_dt__update__tmp_h2).Dtor_cf22(), _751_dt__update_hcf23_h0, (_750_dt__update__tmp_h2).Dtor_cf24())
							}(_pat_let27_0)
						}(_dafny.Companion_Sequence_.Update(_pat_let_tv25, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(211), _dafny.IntOfUint32((_pat_let_tv26).Cardinality()))).Uint32(), _pat_let_tv27))
					}(_pat_let26_0)
				}(_735_v53)
				_ = _rhs89
				var _lhs69 *GlobalState = globalState
				_ = _lhs69
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				var _lhs71 _dafny.Array = _746_v63
				_ = _lhs71
				var _lhs72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_746_v63), 0))
				_ = _lhs72
				_lhs69.F0 = _rhs87
				_lhs70.F5 = _rhs88
				(_lhs71).ArraySet1(_rhs89, (_lhs72).Int())
			} else {
				_725_v44 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(533))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg33 _dafny.Int) interface{} {
						return coer33(arg33)
					}
				}(func(_752_i6 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('m')
				}))
				var _753_v65 _dafny.Map
				_ = _753_v65
				_753_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm15(p0, p1, globalState), _dafny.SetOf(p3, !(p3), p1, p3))
				var _754_v66 D0
				_ = _754_v66
				_754_v66 = Companion_D0_.Create_DC0_(_753_v65)
				(globalState).F5 = Companion_Default___.Fm1(_754_v66, (func() bool {
					if p3 {
						return p1
					}
					return !(p1)
				})(), p0, (func() bool {
					if false {
						return p1
					}
					return !(p3)
				})(), globalState)
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))
				_ = _index111
				(_727_v46).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F6(), (_dafny.IntOfInt64(352)).Minus((_727_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))).Int()).(_dafny.Int))), (_index111).Int())
				var _755_v67 _dafny.CodePoint
				_ = _755_v67
				_755_v67 = _dafny.CodePoint('w')
				var _756_v68 D0
				_ = _756_v68
				_756_v68 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(5), _755_v67, (_this).F6(), p0, p2)
				var _757_v69 D0
				_ = _757_v69
				_757_v69 = Companion_D0_.Create_DC3_(_756_v68)
				_757_v69 = Companion_Default___.Fm32(globalState)
				var _758_v70 _dafny.Sequence
				_ = _758_v70
				_758_v70 = _dafny.SeqOf((_727_v46).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_727_v46), 0))).Int()).(_dafny.Int))
				_758_v70 = (Companion_Default___.Fm33(globalState)).Dtor_cf11()
			}
		}
		var _759_v71 _dafny.Sequence
		_ = _759_v71
		_759_v71 = _dafny.UnicodeSeqOfUtf8Bytes("tvo")
		_759_v71 = _759_v71
		var _hi7 _dafny.Int = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(421))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_760_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		}))).Cardinality())).Plus(p0)
		_ = _hi7
		for _761_i7 := Companion_Default___.Fm24((_this).F6(), globalState); _761_i7.Cmp(_hi7) < 0; _761_i7 = _761_i7.Plus(_dafny.One) {
			var _762_v72 _dafny.Sequence
			_ = _762_v72
			_762_v72 = _dafny.SeqOf(p2)
			var _763_v73 _dafny.Array
			_ = _763_v73
			var _nwElement0_26 _dafny.Int = _761_i7
			_ = _nwElement0_26
			var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(14))
			_ = _nw123
			(_nw123).ArraySet1(_nwElement0_26, 0)
			(_nw123).ArraySet1((_dafny.Zero).Minus(((_dafny.Zero).Minus(p0)).Times((_762_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_762_v72).Cardinality()))).Uint32()).(_dafny.Int))), 1)
			(_nw123).ArraySet1(_761_i7, 2)
			(_nw123).ArraySet1((_this).F6(), 3)
			(_nw123).ArraySet1(p0, 4)
			(_nw123).ArraySet1((_762_v72).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_762_v72).Cardinality()))).Uint32()).(_dafny.Int), 5)
			(_nw123).ArraySet1((_this).F6(), 6)
			(_nw123).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm2(_759_v71, globalState)).Cardinality())), 7)
			(_nw123).ArraySet1(p2, 8)
			(_nw123).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qqajrefo")).Cardinality()), 9)
			(_nw123).ArraySet1(_dafny.IntOfUint32((_759_v71).Cardinality()), 10)
			(_nw123).ArraySet1(_dafny.IntOfUint32((_759_v71).Cardinality()), 11)
			(_nw123).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.Fm15(p2, p1, globalState), (_dafny.Zero).Minus(_761_i7)), 12)
			(_nw123).ArraySet1(_dafny.IntOfInt64(898), 13)
			_763_v73 = _nw123
			var _rhs90 _dafny.Int = p2
			_ = _rhs90
			var _rhs91 _dafny.Int = p2
			_ = _rhs91
			var _rhs92 _dafny.Array = _763_v73
			_ = _rhs92
			var _lhs73 *GlobalState = globalState
			_ = _lhs73
			var _lhs74 *GlobalState = globalState
			_ = _lhs74
			_lhs73.F2 = _rhs90
			_lhs74.F2 = _rhs91
			_763_v73 = _rhs92
			var _764_v74 _dafny.Array
			_ = _764_v74
			var _nw124 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(3))
			_ = _nw124
			_764_v74 = _nw124
			var _765_v75 _dafny.Map
			_ = _765_v75
			_765_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(p0, _761_i7, _761_i7, _761_i7, (_dafny.Zero).Minus(p0)), p1)
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_764_v74), 0))
			_ = _index112
			(_764_v74).ArraySet1(_765_v75, (_index112).Int())
			var _766_v76 _dafny.Array
			_ = _766_v76
			var _nw125 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
			_ = _nw125
			_766_v76 = _nw125
			var _767_v77 _dafny.Set
			_ = _767_v77
			_767_v77 = _dafny.SetOf(_766_v76)
			var _768_v78 D0
			_ = _768_v78
			_768_v78 = Companion_D0_.Create_DC1_(_dafny.IntOfInt64(117), p3, (_this).F6(), p1)
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_764_v74), 0))
			_ = _index113
			var _rhs93 bool = (_767_v77).IsSubsetOf(_767_v77)
			_ = _rhs93
			var _rhs94 _dafny.Map = _765_v75
			_ = _rhs94
			var _rhs95 bool = (_768_v78).Dtor_cf2()
			_ = _rhs95
			var _rhs96 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs96
			var _rhs97 bool = p1
			_ = _rhs97
			var _lhs75 *GlobalState = globalState
			_ = _lhs75
			var _lhs76 _dafny.Array = _764_v74
			_ = _lhs76
			var _lhs77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_764_v74), 0))
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			var _lhs79 *GlobalState = globalState
			_ = _lhs79
			_lhs75.F5 = _rhs93
			(_lhs76).ArraySet1(_rhs94, (_lhs77).Int())
			_lhs78.F5 = _rhs95
			r0 = _rhs96
			_lhs79.F1 = _rhs97
			var _769_v79 T0
			_ = _769_v79
			var _nw126 *C3 = New_C3_()
			_ = _nw126
			_nw126.Ctor__(p2, _762_v72, (_this).F6())
			_769_v79 = _nw126
			var _770_v80 _dafny.MultiSet
			_ = _770_v80
			_770_v80 = _dafny.MultiSetOf(_769_v79)
			var _rhs98 _dafny.Int = ((_770_v80).Update(_769_v79, Companion_Default___.Abs((_dafny.Zero).Minus(_761_i7)))).Cardinality()
			_ = _rhs98
			var _rhs99 _dafny.Int = (Companion_Default___.Fm24((_this).F6(), globalState)).Minus(p2)
			_ = _rhs99
			var _rhs100 _dafny.Int = _dafny.IntOfInt64(600)
			_ = _rhs100
			var _rhs101 _dafny.Int = (_769_v79).F6()
			_ = _rhs101
			var _lhs80 *GlobalState = globalState
			_ = _lhs80
			var _lhs81 *GlobalState = globalState
			_ = _lhs81
			var _lhs82 *GlobalState = globalState
			_ = _lhs82
			_lhs80.F2 = _rhs98
			_lhs81.F0 = _rhs99
			_lhs82.F2 = _rhs100
			r2 = _rhs101
			var _771_v81 _dafny.CodePoint
			_ = _771_v81
			_771_v81 = _dafny.CodePoint('o')
			var _772_v82 _dafny.Set
			_ = _772_v82
			_772_v82 = _dafny.SetOf(_771_v81, _771_v81)
			var _773_v83 _dafny.Map
			_ = _773_v83
			_773_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p0)
			var _774_v84 _dafny.Set
			_ = _774_v84
			_774_v84 = _dafny.SetOf((_772_v82).Cardinality(), _dafny.IntOfUint32((_759_v71).Cardinality()), p0, (_773_v83).Cardinality())
			if (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("l"), _dafny.UnicodeSeqOfUtf8Bytes("ngk"))).Cardinality())).Cmp((_774_v84).Cardinality()) <= 0 {
				var _775_v85 _dafny.Map
				_ = _775_v85
				_775_v85 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p1)
				var _776_v86 _dafny.Map
				_ = _776_v86
				_776_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_775_v85, _766_v76)
				var _777_v87 _dafny.Map
				_ = _777_v87
				_777_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_769_v79).F6(), (_769_v79).F6())
				var _778_v88 *C3
				_ = _778_v88
				var _nw127 *C3 = New_C3_()
				_ = _nw127
				_nw127.Ctor__((_dafny.Zero).Minus((_769_v79).F6()), _dafny.SeqOf(p0, _761_i7, (_769_v79).F6()), (func() _dafny.Int {
					if (_777_v87).Contains(p2) {
						return (_777_v87).Get(p2).(_dafny.Int)
					}
					return p0
				})())
				_778_v88 = _nw127
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_763_v73), 0))
				_ = _index114
				(_763_v73).ArraySet1((_773_v83).Cardinality(), (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_763_v73), 0))
				_ = _index115
				var _rhs102 _dafny.Map = (_776_v86).Merge(_776_v86)
				_ = _rhs102
				var _rhs103 _dafny.Array = _763_v73
				_ = _rhs103
				var _rhs104 *C3 = _778_v88
				_ = _rhs104
				var _rhs105 _dafny.Int = Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus((_775_v85).Cardinality())).Times((_769_v79).F6()), ((_this).F6()).Minus(p0))
				_ = _rhs105
				var _rhs106 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(p2), ((_778_v88).F10()).Select((Companion_Default___.SafeIndex((_769_v79).F6(), _dafny.IntOfUint32(((_778_v88).F10()).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs106
				var _lhs83 _dafny.Array = _763_v73
				_ = _lhs83
				var _lhs84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(895), _dafny.ArrayLen((_763_v73), 0))
				_ = _lhs84
				_776_v86 = _rhs102
				_763_v73 = _rhs103
				_778_v88 = _rhs104
				(_lhs83).ArraySet1(_rhs105, (_lhs84).Int())
				r2 = _rhs106
				var _779_v89 _dafny.MultiSet
				_ = _779_v89
				_779_v89 = _dafny.MultiSetOf(p3, p3)
				(globalState).F2 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_761_i7), ((_779_v89).Cardinality()).Minus(_761_i7))
				var _780_v90 D8
				_ = _780_v90
				_780_v90 = Companion_D8_.Create_DC21_(true, p3, (_769_v79).F6())
				var _781_v91 *C0
				_ = _781_v91
				var _nw128 *C0 = New_C0_()
				_ = _nw128
				_nw128.Ctor__((_780_v90).Dtor_cf37())
				_781_v91 = _nw128
				_781_v91 = (func() *C0 {
					if (func() bool {
						if p1 {
							return p1
						}
						return p1
					})() {
						return _781_v91
					}
					return _781_v91
				})()
				var _782_v92 *C0
				_ = _782_v92
				var _nw129 *C0 = New_C0_()
				_ = _nw129
				_nw129.Ctor__(p1)
				_782_v92 = _nw129
				var _783_v93 *C3
				_ = _783_v93
				var _nw130 *C3 = New_C3_()
				_ = _nw130
				_nw130.Ctor__(_dafny.IntOfInt64(-724), _762_v72, (_this).F6())
				_783_v93 = _nw130
			} else {
				_763_v73 = _763_v73
				var _784_v94 _dafny.Array
				_ = _784_v94
				var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(27))
				_ = _nw131
				_784_v94 = _nw131
				var _785_v95 _dafny.MultiSet
				_ = _785_v95
				_785_v95 = _dafny.MultiSetOf((_dafny.Zero).Minus((_769_v79).F6()))
				var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_784_v94), 0))
				_ = _index116
				(_784_v94).ArraySet1(_785_v95, (_index116).Int())
				var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(857), _dafny.ArrayLen((_784_v94), 0))
				_ = _index117
				(_784_v94).ArraySet1(_785_v95, (_index117).Int())
				var _786_v96 _dafny.Set
				_ = _786_v96
				_786_v96 = _dafny.SetOf((_769_v79))
				var _787_v97 _dafny.MultiSet
				_ = _787_v97
				_787_v97 = _dafny.MultiSetOf(p1)
				r3 = _dafny.SetOf((_786_v96).IsProperSubsetOf(_786_v96), (_787_v97).IsDisjointFrom(_787_v97))
				var _788_v98 D7
				_ = _788_v98
				_788_v98 = Companion_D7_.Create_DC18_(!(p1), p1, _761_i7, p3)
				var _789_v99 _dafny.Map
				_ = _789_v99
				_789_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_788_v98, _766_v76)
				var _790_v100 D11
				_ = _790_v100
				_790_v100 = Companion_D11_.Create_DC24_(_789_v99)
				_789_v99 = ((_790_v100).Dtor_cf41()).Merge((_789_v99).Merge(_789_v99))
				var _791_v101 _dafny.Set
				_ = _791_v101
				_791_v101 = _dafny.SetOf(p1)
				var _792_v102 D0
				_ = _792_v102
				_792_v102 = Companion_D0_.Create_DC0_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _791_v101))
				var _793_v103 _dafny.Sequence
				_ = _793_v103
				_793_v103 = _dafny.SeqOf(p3)
				var _794_v104 D3
				_ = _794_v104
				_794_v104 = Companion_D3_.Create_DC7_(_759_v71)
				var _795_v105 D12
				_ = _795_v105
				_795_v105 = Companion_D12_.Create_DC28_(Companion_Default___.Fm34(globalState))
				var _rhs107 bool = Companion_Default___.Fm1(_792_v102, p1, Companion_Default___.Fm15(_dafny.IntOfInt64(388), (_793_v103).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_793_v103).Cardinality()))).Uint32()).(bool), globalState), (!(p3)) || (true), globalState)
				_ = _rhs107
				var _rhs108 bool = !_dafny.Companion_Sequence_.Contains((_794_v104).Dtor_cf13(), _771_v81)
				_ = _rhs108
				var _rhs109 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(601))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg35 _dafny.Int) interface{} {
						return coer35(arg35)
					}
				}((func(_796_v81 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_797_i9 _dafny.Int) _dafny.CodePoint {
						return _796_v81
					}
				})(_771_v81)))
				_ = _rhs109
				var _rhs110 bool = (_793_v103).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_793_v103).Cardinality()))).Uint32()).(bool)
				_ = _rhs110
				var _rhs111 _dafny.Map = (_795_v105).Dtor_cf51()
				_ = _rhs111
				var _lhs85 *GlobalState = globalState
				_ = _lhs85
				var _lhs86 *GlobalState = globalState
				_ = _lhs86
				var _lhs87 *GlobalState = globalState
				_ = _lhs87
				_lhs85.F5 = _rhs107
				_lhs86.F1 = _rhs108
				_759_v71 = _rhs109
				_lhs87.F5 = _rhs110
				_773_v83 = _rhs111
			}
		}
		var _798_v106 _dafny.Set
		_ = _798_v106
		_798_v106 = _dafny.SetOf(p1, p3)
		var _799_v107 _dafny.Map
		_ = _799_v107
		_799_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _798_v106)
		var _800_v108 D0
		_ = _800_v108
		_800_v108 = Companion_D0_.Create_DC0_(_799_v107)
		var _801_v109 _dafny.Set
		_ = _801_v109
		_801_v109 = _dafny.SetOf(Companion_Default___.Fm2(_dafny.UnicodeSeqOfUtf8Bytes("et"), globalState), _759_v71)
		var _802_v110 _dafny.Sequence
		_ = _802_v110
		_802_v110 = _dafny.SeqOf(p1)
		var _803_v111 _dafny.Sequence
		_ = _803_v111
		_803_v111 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_802_v110, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_802_v110).Cardinality()))).Uint32(), false), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_802_v110, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_802_v110).Cardinality()))).Uint32(), false)).Cardinality()))).Uint32(), p3))
		var _rhs112 D0 = _800_v108
		_ = _rhs112
		var _rhs113 bool = _dafny.Companion_Sequence_.IsProperPrefixOf(_803_v111, _dafny.SeqOf(_802_v110))
		_ = _rhs113
		var _rhs114 _dafny.Set = (func() _dafny.Set {
			var _coll44 = _dafny.NewBuilder()
			_ = _coll44
			for _iter47 := _dafny.Iterate((_801_v109).Elements()); ; {
				_compr_44, _ok47 := _iter47()
				if !_ok47 {
					break
				}
				var _804_v112 _dafny.Sequence
				_804_v112 = interface{}(_compr_44).(_dafny.Sequence)
				if (_801_v109).Contains(_804_v112) {
					_coll44.Add(_804_v112)
				}
			}
			return _coll44.ToSet()
		}()).Difference((_801_v109).Difference(_801_v109))
		_ = _rhs114
		var _rhs115 bool = p3
		_ = _rhs115
		var _lhs88 *GlobalState = globalState
		_ = _lhs88
		var _lhs89 *GlobalState = globalState
		_ = _lhs89
		_800_v108 = _rhs112
		_lhs88.F5 = _rhs113
		_801_v109 = _rhs114
		_lhs89.F1 = _rhs115
		r0 = (_this).F6()
		r1 = Companion_D1_.Create_DC5_()
		var _805_v113 _dafny.Map
		_ = _805_v113
		_805_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		r2 = (_805_v113).Cardinality()
		r3 = _798_v106
		return r0, r1, r2, r3
	}
}
func (_this *C4) M4(globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Int, _dafny.CodePoint) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.CodePoint = _dafny.CodePoint('D')
		_ = r3
		var _806_i0 _dafny.Int
		_ = _806_i0
		_806_i0 = _dafny.Zero
		{
			for false {
				{
					if (_806_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_806_i0 = (_806_i0).Plus(_dafny.One)
					r2 = (_dafny.Zero).Minus((_this).F6())
					var _807_v0 bool
					_ = _807_v0
					_807_v0 = false
					var _808_v1 _dafny.Sequence
					_ = _808_v1
					_808_v1 = _dafny.SeqOf(true, _807_v0, _807_v0, _807_v0, _807_v0)
					var _809_v2 D7
					_ = _809_v2
					_809_v2 = Companion_D7_.Create_DC19_((func() _dafny.MultiSet {
						if _807_v0 {
							return _dafny.MultiSetOf(false, _807_v0)
						}
						return (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_808_v1, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_808_v1).Cardinality()))).Uint32(), _807_v0))).Update(_807_v0, Companion_Default___.Abs((_this).F6()))
					})(), _dafny.IntOfInt64(415))
					var _source13 D7 = _809_v2
					_ = _source13
					if _source13.Is_DC18() {
						var _810___mcc_h0 bool = _source13.Get_().(D7_DC18).Cf29
						_ = _810___mcc_h0
						var _811___mcc_h1 bool = _source13.Get_().(D7_DC18).Cf30
						_ = _811___mcc_h1
						var _812___mcc_h2 _dafny.Int = _source13.Get_().(D7_DC18).Cf31
						_ = _812___mcc_h2
						var _813___mcc_h3 bool = _source13.Get_().(D7_DC18).Cf32
						_ = _813___mcc_h3
						var _814_cf32 bool = _813___mcc_h3
						_ = _814_cf32
						var _815_cf31 _dafny.Int = _812___mcc_h2
						_ = _815_cf31
						var _816_cf30 bool = _811___mcc_h1
						_ = _816_cf30
						var _817_cf29 bool = _810___mcc_h0
						_ = _817_cf29
						var _818_v3 _dafny.Set
						_ = _818_v3
						_818_v3 = _dafny.SetOf(_814_cf32)
						r1 = (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(454))).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg36 _dafny.Int) interface{} {
								return coer36(arg36)
							}
						}((func(_819_cf31 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_820_i1 _dafny.Int) _dafny.Int {
								return _819_cf31
							}
						})(_815_cf31)))).Cardinality())).Minus(((_818_v3).Cardinality()).Minus((_dafny.Zero).Minus((_this).F6())))
						var _821_v4 _dafny.Sequence
						_ = _821_v4
						_821_v4 = _dafny.UnicodeSeqOfUtf8Bytes("mqmrskldp")
						var _822_v5 *C2
						_ = _822_v5
						var _nw132 *C2 = New_C2_()
						_ = _nw132
						_nw132.Ctor__(_821_v4, (func() bool {
							if _814_cf32 {
								return _816_cf30
							}
							return _816_cf30
						})())
						_822_v5 = _nw132
						(globalState).F2 = _dafny.IntOfUint32((_808_v1).Cardinality())
						var _823_v6 D3
						_ = _823_v6
						_823_v6 = Companion_D3_.Create_DC7_(_dafny.Companion_Sequence_.Concatenate((_822_v5).F11(), (_822_v5).F11()))
						var _824_v7 _dafny.Array
						_ = _824_v7
						var _nw133 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
						_ = _nw133
						_824_v7 = _nw133
						var _825_v8 _dafny.Array
						_ = _825_v8
						var _len0_21 _dafny.Int = _dafny.IntOfInt64(2)
						_ = _len0_21
						var _nw134 _dafny.Array
						_ = _nw134
						if _len0_21.Cmp(_dafny.Zero) == 0 {
							_nw134 = _dafny.NewArray(_len0_21)
						} else {
							var _init21 func(_dafny.Int) _dafny.CodePoint = (func(_826_v5 *C2) func(_dafny.Int) _dafny.CodePoint {
								return func(_827_i2 _dafny.Int) _dafny.CodePoint {
									return (func() _dafny.CodePoint {
										if _826_v5.F12 {
											return _dafny.CodePoint('g')
										}
										return _dafny.CodePoint('p')
									})()
								}
							})(_822_v5)
							_ = _init21
							var _element0_21 = _init21(_dafny.Zero)
							_ = _element0_21
							_nw134 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
							(_nw134).ArraySet1CodePoint(_element0_21, 0)
							var _nativeLen0_21 = (_len0_21).Int()
							_ = _nativeLen0_21
							for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
								(_nw134).ArraySet1CodePoint(_init21(_dafny.IntOf(_i0_21)), _i0_21)
							}
						}
						_825_v8 = _nw134
						var _828_v9 _dafny.CodePoint
						_ = _828_v9
						_828_v9 = _dafny.CodePoint('b')
						var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_825_v8), 0))
						_ = _index118
						(_825_v8).ArraySet1CodePoint(_828_v9, (_index118).Int())
						var _829_v10 _dafny.MultiSet
						_ = _829_v10
						_829_v10 = _dafny.MultiSetOf(_815_cf31)
						var _pat_let_tv28 = _821_v4
						_ = _pat_let_tv28
						var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_825_v8), 0))
						_ = _index119
						var _rhs116 D3 = func(_pat_let28_0 D3) D3 {
							return func(_832_dt__update__tmp_h1 D3) D3 {
								return func(_pat_let31_0 _dafny.Sequence) D3 {
									return func(_833_dt__update_hcf13_h1 _dafny.Sequence) D3 {
										return Companion_D3_.Create_DC7_(_833_dt__update_hcf13_h1)
									}(_pat_let31_0)
								}(_dafny.UnicodeSeqOfUtf8Bytes("cwj"))
							}(_pat_let28_0)
						}(func(_pat_let29_0 D3) D3 {
							return func(_830_dt__update__tmp_h0 D3) D3 {
								return func(_pat_let30_0 _dafny.Sequence) D3 {
									return func(_831_dt__update_hcf13_h0 _dafny.Sequence) D3 {
										return Companion_D3_.Create_DC7_(_831_dt__update_hcf13_h0)
									}(_pat_let30_0)
								}(_pat_let_tv28)
							}(_pat_let29_0)
						}(_823_v6))
						_ = _rhs116
						var _rhs117 _dafny.Array = _824_v7
						_ = _rhs117
						var _rhs118 _dafny.CodePoint = _828_v9
						_ = _rhs118
						var _rhs119 _dafny.MultiSet = _829_v10
						_ = _rhs119
						var _rhs120 bool = !(_dafny.Companion_Sequence_.Contains((_822_v5).F11(), _828_v9))
						_ = _rhs120
						var _lhs90 _dafny.Array = _825_v8
						_ = _lhs90
						var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(628), _dafny.ArrayLen((_825_v8), 0))
						_ = _lhs91
						_823_v6 = _rhs116
						_824_v7 = _rhs117
						(_lhs90).ArraySet1CodePoint(_rhs118, (_lhs91).Int())
						_829_v10 = _rhs119
						_814_cf32 = _rhs120
					} else if _source13.Is_DC19() {
						var _834___mcc_h4 _dafny.MultiSet = _source13.Get_().(D7_DC19).Cf33
						_ = _834___mcc_h4
						var _835___mcc_h5 _dafny.Int = _source13.Get_().(D7_DC19).Cf34
						_ = _835___mcc_h5
						var _836_cf34 _dafny.Int = _835___mcc_h5
						_ = _836_cf34
						var _837_cf33 _dafny.MultiSet = _834___mcc_h4
						_ = _837_cf33
						var _838_v11 _dafny.CodePoint
						_ = _838_v11
						_838_v11 = _dafny.CodePoint('i')
						var _839_v12 _dafny.MultiSet
						_ = _839_v12
						_839_v12 = _dafny.MultiSetOf(_838_v11)
						var _840_v13 _dafny.Map
						_ = _840_v13
						_840_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_836_cf34, _807_v0)
						var _841_v14 _dafny.Map
						_ = _841_v14
						_841_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm15((func() _dafny.Int {
							if (_839_v12).Contains(_838_v11) {
								return (_839_v12).Multiplicity(_838_v11)
							}
							return _836_cf34
						})(), (func() bool {
							if (_840_v13).Contains((_this).F6()) {
								return (_840_v13).Get((_this).F6()).(bool)
							}
							return true
						})(), globalState), _807_v0)
						var _842_v15 D13
						_ = _842_v15
						_842_v15 = Companion_D13_.Create_DC32_((_this).F6(), _807_v0, _807_v0, _807_v0, (_808_v1).Select((Companion_Default___.SafeIndex(_836_cf34, _dafny.IntOfUint32((_808_v1).Cardinality()))).Uint32()).(bool))
						var _843_v16 _dafny.Sequence
						_ = _843_v16
						_843_v16 = _dafny.UnicodeSeqOfUtf8Bytes("ohdkp")
						var _844_v17 _dafny.Sequence
						_ = _844_v17
						_844_v17 = _dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_843_v16).Cardinality())))
						var _845_v18 _dafny.Set
						_ = _845_v18
						_845_v18 = _dafny.SetOf(_807_v0)
						var _846_v19 _dafny.Array
						_ = _846_v19
						var _len0_22 _dafny.Int = _dafny.IntOfInt64(4)
						_ = _len0_22
						var _nw135 _dafny.Array
						_ = _nw135
						if _len0_22.Cmp(_dafny.Zero) == 0 {
							_nw135 = _dafny.NewArray(_len0_22)
						} else {
							var _init22 func(_dafny.Int) bool = (func(_847_v0 bool) func(_dafny.Int) bool {
								return func(_848_i3 _dafny.Int) bool {
									return _847_v0
								}
							})(_807_v0)
							_ = _init22
							var _element0_22 = _init22(_dafny.Zero)
							_ = _element0_22
							_nw135 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
							(_nw135).ArraySet1(_element0_22, 0)
							var _nativeLen0_22 = (_len0_22).Int()
							_ = _nativeLen0_22
							for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
								(_nw135).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
							}
						}
						_846_v19 = _nw135
						var _849_v20 _dafny.Map
						_ = _849_v20
						_849_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_846_v19, _841_v14)
						var _850_v21 _dafny.MultiSet
						_ = _850_v21
						_850_v21 = _dafny.MultiSetOf(_836_cf34)
						var _851_v22 _dafny.Array
						_ = _851_v22
						var _nwElement0_27 bool = _807_v0
						_ = _nwElement0_27
						var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(28))
						_ = _nw136
						(_nw136).ArraySet1(_nwElement0_27, 0)
						(_nw136).ArraySet1(_807_v0, 1)
						(_nw136).ArraySet1(true, 2)
						(_nw136).ArraySet1(!(_807_v0), 3)
						(_nw136).ArraySet1(!(_807_v0), 4)
						(_nw136).ArraySet1(false, 5)
						(_nw136).ArraySet1(_807_v0, 6)
						(_nw136).ArraySet1((_840_v13).Equals(((Companion_D13_.Create_DC31_(_841_v14)).Dtor_cf55()).Update(_836_cf34, false)), 7)
						(_nw136).ArraySet1(_807_v0, 8)
						(_nw136).ArraySet1((_842_v15).Dtor_cf59(), 9)
						(_nw136).ArraySet1((_dafny.IntOfUint32((_844_v17).Cardinality())).Cmp((_this).F6()) > 0, 10)
						(_nw136).ArraySet1(_807_v0, 11)
						(_nw136).ArraySet1(((_dafny.Zero).Minus((_845_v18).Cardinality())).Cmp((_this).F6()) > 0, 12)
						(_nw136).ArraySet1(!(_807_v0) || (_807_v0), 13)
						(_nw136).ArraySet1((_807_v0) && (_807_v0), 14)
						(_nw136).ArraySet1(_807_v0, 15)
						(_nw136).ArraySet1(_807_v0, 16)
						(_nw136).ArraySet1((_dafny.MultiSetOf(true, _807_v0, _807_v0, _807_v0, (func() bool {
							if (_840_v13).Contains((_849_v20).Cardinality()) {
								return (_840_v13).Get((_849_v20).Cardinality()).(bool)
							}
							return _807_v0
						})())).IsSubsetOf(_837_cf33), 17)
						(_nw136).ArraySet1(_807_v0, 18)
						(_nw136).ArraySet1(_807_v0, 19)
						(_nw136).ArraySet1(!_dafny.Companion_Sequence_.Equal(_844_v17, _844_v17), 20)
						(_nw136).ArraySet1(_807_v0, 21)
						(_nw136).ArraySet1(_807_v0, 22)
						(_nw136).ArraySet1(_807_v0, 23)
						(_nw136).ArraySet1(_807_v0, 24)
						(_nw136).ArraySet1((_836_cf34).Cmp((_this).F6()) == 0, 25)
						(_nw136).ArraySet1(!((_dafny.MultiSetFromSeq(_844_v17)).IsProperSubsetOf(_850_v21)), 26)
						(_nw136).ArraySet1(!(false) || (_807_v0), 27)
						_851_v22 = _nw136
						_851_v22 = _846_v19
						var _852_v23 _dafny.Array
						_ = _852_v23
						var _nwElement0_28 _dafny.Sequence = _843_v16
						_ = _nwElement0_28
						var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(14))
						_ = _nw137
						(_nw137).ArraySet1(_nwElement0_28, 0)
						(_nw137).ArraySet1(_843_v16, 1)
						(_nw137).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(896))).Uint32(), func(coer37 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg37 _dafny.Int) interface{} {
								return coer37(arg37)
							}
						}((func(_853_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_854_i4 _dafny.Int) _dafny.CodePoint {
								return _853_v11
							}
						})(_838_v11))), 2)
						(_nw137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("khfcae"), 3)
						(_nw137).ArraySet1(_843_v16, 4)
						(_nw137).ArraySet1(_843_v16, 5)
						(_nw137).ArraySet1(_843_v16, 6)
						(_nw137).ArraySet1(_843_v16, 7)
						(_nw137).ArraySet1(_843_v16, 8)
						(_nw137).ArraySet1(_843_v16, 9)
						(_nw137).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_843_v16, _843_v16), 10)
						(_nw137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("yareb"), 11)
						(_nw137).ArraySet1(_843_v16, 12)
						(_nw137).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("aqkxf"), 13)
						_852_v23 = _nw137
						var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_852_v23), 0))
						_ = _index120
						(_852_v23).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_843_v16, _843_v16), (_index120).Int())
						var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(152), _dafny.ArrayLen((_852_v23), 0))
						_ = _index121
						(_852_v23).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_843_v16, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-243))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg38 _dafny.Int) interface{} {
								return coer38(arg38)
							}
						}((func(_855_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_856_i5 _dafny.Int) _dafny.CodePoint {
								return _855_v11
							}
						})(_838_v11))), _843_v16)), (Companion_Default___.SafeIndex(_836_cf34, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_843_v16, _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-243))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg39 _dafny.Int) interface{} {
								return coer39(arg39)
							}
						}((func(_857_v11 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
							return func(_858_i5 _dafny.Int) _dafny.CodePoint {
								return _857_v11
							}
						})(_838_v11))), _843_v16))).Cardinality()))).Uint32(), _838_v11), (_index121).Int())
						(globalState).F5 = _807_v0
						_844_v17 = _844_v17
					} else {
						var _859___mcc_h6 _dafny.Sequence = _source13.Get_().(D7_DC17).Cf28
						_ = _859___mcc_h6
						var _860_cf28 _dafny.Sequence = _859___mcc_h6
						_ = _860_cf28
						var _861_v24 _dafny.Array
						_ = _861_v24
						var _len0_23 _dafny.Int = _dafny.IntOfInt64(8)
						_ = _len0_23
						var _nw138 _dafny.Array
						_ = _nw138
						if _len0_23.Cmp(_dafny.Zero) == 0 {
							_nw138 = _dafny.NewArray(_len0_23)
						} else {
							var _init23 func(_dafny.Int) bool = (func(_862_v0 bool) func(_dafny.Int) bool {
								return func(_863_i6 _dafny.Int) bool {
									return _862_v0
								}
							})(_807_v0)
							_ = _init23
							var _element0_23 = _init23(_dafny.Zero)
							_ = _element0_23
							_nw138 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
							(_nw138).ArraySet1(_element0_23, 0)
							var _nativeLen0_23 = (_len0_23).Int()
							_ = _nativeLen0_23
							for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
								(_nw138).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
							}
						}
						_861_v24 = _nw138
						var _864_v25 _dafny.Sequence
						_ = _864_v25
						_864_v25 = _dafny.SeqOf(_861_v24, _861_v24)
						_864_v25 = _864_v25
						(globalState).F1 = _807_v0
						var _865_v26 _dafny.CodePoint
						_ = _865_v26
						_865_v26 = _dafny.CodePoint('o')
						r3 = _865_v26
						var _866_v27 _dafny.Set
						_ = _866_v27
						_866_v27 = _dafny.SetOf(_807_v0, _807_v0, _807_v0)
						var _867_v28 _dafny.Map
						_ = _867_v28
						_867_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_866_v27).Equals(_866_v27), (_this).F6())
						_867_v28 = (_867_v28).Update(_807_v0, (_this).F6())
					}
					var _868_v29 _dafny.Sequence
					_ = _868_v29
					_868_v29 = _dafny.UnicodeSeqOfUtf8Bytes("lropjewiy")
					var _869_v30 _dafny.Set
					_ = _869_v30
					_869_v30 = _dafny.SetOf(_807_v0)
					var _870_v31 _dafny.Set
					_ = _870_v31
					_870_v31 = _dafny.SetOf((_869_v30).Cardinality(), (Companion_Default___.Fm15((_this).F6(), _807_v0, globalState)).Plus((_this).F6()), (_this).F6(), (_dafny.Zero).Minus((_this).F6()))
					var _871_v32 _dafny.Array
					_ = _871_v32
					var _len0_24 _dafny.Int = _dafny.IntOfInt64(7)
					_ = _len0_24
					var _nw139 _dafny.Array
					_ = _nw139
					if _len0_24.Cmp(_dafny.Zero) == 0 {
						_nw139 = _dafny.NewArray(_len0_24)
					} else {
						var _init24 func(_dafny.Int) bool = (func(_872_v0 bool) func(_dafny.Int) bool {
							return func(_873_i7 _dafny.Int) bool {
								return _872_v0
							}
						})(_807_v0)
						_ = _init24
						var _element0_24 = _init24(_dafny.Zero)
						_ = _element0_24
						_nw139 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
						(_nw139).ArraySet1(_element0_24, 0)
						var _nativeLen0_24 = (_len0_24).Int()
						_ = _nativeLen0_24
						for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
							(_nw139).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
						}
					}
					_871_v32 = _nw139
					var _874_v33 _dafny.Array
					_ = _874_v33
					var _nw140 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
					_ = _nw140
					_874_v33 = _nw140
					var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))
					_ = _index122
					(_874_v33).ArraySet1((_this).F6(), (_index122).Int())
					var _875_v34 D7
					_ = _875_v34
					_875_v34 = Companion_D7_.Create_DC18_(_807_v0, _807_v0, (_dafny.Zero).Minus((_this).F6()), _807_v0)
					var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))
					_ = _index123
					var _rhs121 _dafny.Sequence = (Companion_D3_.Create_DC7_(_868_v29)).Dtor_cf13()
					_ = _rhs121
					var _rhs122 _dafny.Set = _dafny.SetOf((_this).F6())
					_ = _rhs122
					var _rhs123 _dafny.Array = _871_v32
					_ = _rhs123
					var _rhs124 _dafny.Int = (_this).F6()
					_ = _rhs124
					var _rhs125 _dafny.Int = (_875_v34).Dtor_cf31()
					_ = _rhs125
					var _lhs92 _dafny.Array = _874_v33
					_ = _lhs92
					var _lhs93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))
					_ = _lhs93
					var _lhs94 *GlobalState = globalState
					_ = _lhs94
					_868_v29 = _rhs121
					_870_v31 = _rhs122
					_871_v32 = _rhs123
					(_lhs92).ArraySet1(_rhs124, (_lhs93).Int())
					_lhs94.F0 = _rhs125
					var _876_v35 _dafny.Map
					_ = _876_v35
					_876_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (func() _dafny.Int {
						if _807_v0 {
							return _dafny.IntOfUint32((_808_v1).Cardinality())
						}
						return (_874_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))).Int()).(_dafny.Int)
					})())
					var _877_v36 _dafny.Sequence
					_ = _877_v36
					_877_v36 = _dafny.SeqOf((_874_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))).Int()).(_dafny.Int), (_this).F6())
					var _878_v37 _dafny.CodePoint
					_ = _878_v37
					_878_v37 = _dafny.CodePoint('p')
					var _879_v38 _dafny.Map
					_ = _879_v38
					_879_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeModuloInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_877_v36, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_808_v1).Cardinality())), _dafny.IntOfUint32((_877_v36).Cardinality()))).Uint32(), (_874_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))).Int()).(_dafny.Int))).Cardinality()), _dafny.IntOfInt64(-615))).Cardinality(), Companion_Default___.Fm24((_874_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))).Int()).(_dafny.Int), globalState)), _878_v37)
					var _rhs126 _dafny.Int = Companion_Default___.Fm15((_874_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(381), _dafny.ArrayLen((_874_v33), 0))).Int()).(_dafny.Int), _807_v0, globalState)
					_ = _rhs126
					var _rhs127 _dafny.CodePoint = (func() _dafny.CodePoint {
						if (_879_v38).Contains((func() _dafny.Map {
							var _coll45 = _dafny.NewMapBuilder()
							_ = _coll45
							for _iter48 := _dafny.Iterate((_868_v29).Elements()); ; {
								_compr_45, _ok48 := _iter48()
								if !_ok48 {
									break
								}
								var _880_v39 _dafny.CodePoint
								_880_v39 = interface{}(_compr_45).(_dafny.CodePoint)
								if _dafny.Companion_Sequence_.Contains(_868_v29, _880_v39) {
									_coll45.Add(_880_v39, _807_v0)
								}
							}
							return _coll45.ToMap()
						}()).Cardinality()) {
							return (_879_v38).Get((func() _dafny.Map {
								var _coll46 = _dafny.NewMapBuilder()
								_ = _coll46
								for _iter49 := _dafny.Iterate((_868_v29).Elements()); ; {
									_compr_46, _ok49 := _iter49()
									if !_ok49 {
										break
									}
									var _881_v39 _dafny.CodePoint
									_881_v39 = interface{}(_compr_46).(_dafny.CodePoint)
									if _dafny.Companion_Sequence_.Contains(_868_v29, _881_v39) {
										_coll46.Add(_881_v39, _807_v0)
									}
								}
								return _coll46.ToMap()
							}()).Cardinality()).(_dafny.CodePoint)
						}
						return _878_v37
					})()
					_ = _rhs127
					var _rhs128 _dafny.Map = _876_v35
					_ = _rhs128
					r1 = _rhs126
					r3 = _rhs127
					_876_v35 = _rhs128
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _882_v40 bool
		_ = _882_v40
		_882_v40 = true
		if _882_v40 {
			(globalState).F5 = _882_v40
			var _883_v41 _dafny.Array
			_ = _883_v41
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(24)
			_ = _len0_25
			var _nw141 _dafny.Array
			_ = _nw141
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw141 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) bool = (func(_884_v40 bool) func(_dafny.Int) bool {
					return func(_885_i8 _dafny.Int) bool {
						return _884_v40
					}
				})(_882_v40)
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw141 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw141).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw141).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_883_v41 = _nw141
			var _886_v42 _dafny.Map
			_ = _886_v42
			_886_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _883_v41)
			var _rhs129 _dafny.Map = _886_v42
			_ = _rhs129
			var _rhs130 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F6()), (_this).F6())
			_ = _rhs130
			_886_v42 = _rhs129
			r1 = _rhs130
			var _887_v43 _dafny.CodePoint
			_ = _887_v43
			_887_v43 = _dafny.CodePoint('j')
			var _888_v44 _dafny.Map
			_ = _888_v44
			_888_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_887_v43, (_this).F6())
			var _889_v45 _dafny.Sequence
			_ = _889_v45
			_889_v45 = _dafny.SeqOf((func() _dafny.Int {
				if (_888_v44).Contains(_887_v43) {
					return (_888_v44).Get(_887_v43).(_dafny.Int)
				}
				return (_this).F6()
			})())
			_889_v45 = (Companion_D1_.Create_DC4_(_889_v45)).Dtor_cf11()
			var _rhs131 _dafny.Array = (func() _dafny.Array {
				if (_886_v42).Contains(_882_v40) {
					return (_886_v42).Get(_882_v40).(_dafny.Array)
				}
				return _883_v41
			})()
			_ = _rhs131
			var _rhs132 _dafny.Int = Companion_Default___.Fm24((_this).F6(), globalState)
			_ = _rhs132
			var _lhs95 *GlobalState = globalState
			_ = _lhs95
			_883_v41 = _rhs131
			_lhs95.F0 = _rhs132
			var _890_v46 _dafny.Sequence
			_ = _890_v46
			_890_v46 = _dafny.UnicodeSeqOfUtf8Bytes("nemkv")
			var _891_v47 _dafny.Map
			_ = _891_v47
			_891_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _890_v46)
			var _892_v48 _dafny.MultiSet
			_ = _892_v48
			_892_v48 = _dafny.MultiSetOf(_882_v40)
			var _893_v49 T0
			_ = _893_v49
			var _nw142 *C3 = New_C3_()
			_ = _nw142
			_nw142.Ctor__(_dafny.IntOfInt64(186), _889_v45, (_892_v48).Cardinality())
			_893_v49 = _nw142
			var _894_v50 T1
			_ = _894_v50
			var _nw143 *C3 = New_C3_()
			_ = _nw143
			_nw143.Ctor__((_this).F6(), _889_v45, _dafny.IntOfUint32((Companion_Default___.Fm2(_890_v46, globalState)).Cardinality()))
			_894_v50 = _nw143
			var _895_v51 _dafny.MultiSet
			_ = _895_v51
			_895_v51 = _dafny.MultiSetOf(_894_v50)
			var _896_v52 _dafny.MultiSet
			_ = _896_v52
			_896_v52 = _895_v51
			var _897_v53 D11
			_ = _897_v53
			_897_v53 = Companion_D11_.Create_DC26_(_893_v49, _896_v52, (_888_v44).Cardinality(), _890_v46)
			var _898_v54 _dafny.Array
			_ = _898_v54
			var _nwElement0_29 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_890_v46, (Companion_Default___.SafeIndex((_889_v45).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_889_v45).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_890_v46).Cardinality()))).Uint32(), _887_v43)
			_ = _nwElement0_29
			var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(17))
			_ = _nw144
			(_nw144).ArraySet1(_nwElement0_29, 0)
			(_nw144).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(199))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg40 _dafny.Int) interface{} {
					return coer40(arg40)
				}
			}((func(_899_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_900_i9 _dafny.Int) _dafny.CodePoint {
					return _899_v43
				}
			})(_887_v43))), 1)
			(_nw144).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(947))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg41 _dafny.Int) interface{} {
					return coer41(arg41)
				}
			}((func(_901_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_902_i10 _dafny.Int) _dafny.CodePoint {
					return _901_v43
				}
			})(_887_v43))), 2)
			(_nw144).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("npv"), 3)
			(_nw144).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ev"), _890_v46), 4)
			(_nw144).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_890_v46, _890_v46), 5)
			(_nw144).ArraySet1(_890_v46, 6)
			(_nw144).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_890_v46, _890_v46), 7)
			(_nw144).ArraySet1(_890_v46, 8)
			(_nw144).ArraySet1(_890_v46, 9)
			(_nw144).ArraySet1((func() _dafny.Sequence {
				if _882_v40 {
					return _890_v46
				}
				return _dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(489))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_903_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_904_i11 _dafny.Int) _dafny.CodePoint {
						return _903_v43
					}
				})(_887_v43))), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(489))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_905_v43 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_906_i11 _dafny.Int) _dafny.CodePoint {
						return _905_v43
					}
				})(_887_v43)))).Cardinality()))).Uint32(), _dafny.CodePoint('p'))
			})(), 10)
			(_nw144).ArraySet1(_890_v46, 11)
			(_nw144).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_890_v46, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_890_v46).Cardinality()))).Uint32(), _887_v43), _890_v46), 12)
			(_nw144).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_890_v46, (func() _dafny.Sequence {
				if (_891_v47).Contains(_882_v40) {
					return (_891_v47).Get(_882_v40).(_dafny.Sequence)
				}
				return _890_v46
			})()), 13)
			(_nw144).ArraySet1(_890_v46, 14)
			(_nw144).ArraySet1((func() _dafny.Sequence {
				if _882_v40 {
					return _890_v46
				}
				return _890_v46
			})(), 15)
			(_nw144).ArraySet1(_dafny.Companion_Sequence_.Update((_897_v53).Dtor_cf48(), (Companion_Default___.SafeIndex((_893_v49).F6(), _dafny.IntOfUint32(((_897_v53).Dtor_cf48()).Cardinality()))).Uint32(), _887_v43), 16)
			_898_v54 = _nw144
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_898_v54), 0))
			_ = _index124
			(_898_v54).ArraySet1(_890_v46, (_index124).Int())
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(946), _dafny.ArrayLen((_898_v54), 0))
			_ = _index125
			(_898_v54).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("nqeesu"), (_index125).Int())
		} else {
			var _907_v55 _dafny.Sequence
			_ = _907_v55
			_907_v55 = _dafny.UnicodeSeqOfUtf8Bytes("rgltcl")
			_907_v55 = Companion_Default___.Fm2(_907_v55, globalState)
			var _908_v56 _dafny.Array
			_ = _908_v56
			var _nw145 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
			_ = _nw145
			_908_v56 = _nw145
			var _909_v57 _dafny.Map
			_ = _909_v57
			_909_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), (_this).F6())).Cardinality()))
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_908_v56), 0))
			_ = _index126
			(_908_v56).ArraySet1(_909_v57, (_index126).Int())
			var _910_v58 _dafny.Map
			_ = _910_v58
			_910_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _dafny.CodePoint('o'))
			var _911_v59 _dafny.CodePoint
			_ = _911_v59
			_911_v59 = _dafny.CodePoint('t')
			var _912_v60 _dafny.Sequence
			_ = _912_v60
			_912_v60 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _dafny.IntOfUint32((_dafny.SeqOf((func() _dafny.CodePoint {
				if (_910_v58).Contains(!(_882_v40)) {
					return (_910_v58).Get(!(_882_v40)).(_dafny.CodePoint)
				}
				return _911_v59
			})())).Cardinality()))).Merge(_909_v57), _909_v57)
			var _913_v61 _dafny.MultiSet
			_ = _913_v61
			_913_v61 = _dafny.MultiSetOf(_882_v40)
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(84), _dafny.ArrayLen((_908_v56), 0))
			_ = _index127
			(_908_v56).ArraySet1((_912_v60).Select((Companion_Default___.SafeIndex(((_913_v61).Cardinality()).Times((_this).F6()), _dafny.IntOfUint32((_912_v60).Cardinality()))).Uint32()).(_dafny.Map), (_index127).Int())
			r0 = _dafny.IntOfInt64(575)
			var _914_v62 _dafny.Map
			_ = _914_v62
			_914_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), _dafny.SetOf(_882_v40))
			var _915_v63 D0
			_ = _915_v63
			_915_v63 = Companion_D0_.Create_DC0_(_914_v62)
			var _916_v64 *C1
			_ = _916_v64
			var _nw146 *C1 = New_C1_()
			_ = _nw146
			_nw146.Ctor__()
			_916_v64 = _nw146
			var _917_v65 _dafny.Sequence
			_ = _917_v65
			_917_v65 = _dafny.SeqOf(_882_v40)
			var _918_v66 _dafny.Array
			_ = _918_v66
			var _nwElement0_30 bool = Companion_Default___.Fm1(_915_v63, _882_v40, (_this).F6(), _882_v40, globalState)
			_ = _nwElement0_30
			var _nw147 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(20))
			_ = _nw147
			(_nw147).ArraySet1(_nwElement0_30, 0)
			(_nw147).ArraySet1(_882_v40, 1)
			(_nw147).ArraySet1((_917_v65).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.IntOfUint32((_917_v65).Cardinality()))).Uint32()).(bool), 2)
			(_nw147).ArraySet1(_882_v40, 3)
			(_nw147).ArraySet1(_882_v40, 4)
			(_nw147).ArraySet1(_882_v40, 5)
			(_nw147).ArraySet1(_882_v40, 6)
			(_nw147).ArraySet1(_882_v40, 7)
			(_nw147).ArraySet1(true, 8)
			(_nw147).ArraySet1(_882_v40, 9)
			(_nw147).ArraySet1(_882_v40, 10)
			(_nw147).ArraySet1(true, 11)
			(_nw147).ArraySet1(_882_v40, 12)
			(_nw147).ArraySet1(!(_882_v40), 13)
			(_nw147).ArraySet1(!(_882_v40), 14)
			(_nw147).ArraySet1(_882_v40, 15)
			(_nw147).ArraySet1(false, 16)
			(_nw147).ArraySet1(_882_v40, 17)
			(_nw147).ArraySet1(_882_v40, 18)
			(_nw147).ArraySet1(_882_v40, 19)
			_918_v66 = _nw147
			var _919_v67 _dafny.MultiSet
			_ = _919_v67
			_919_v67 = _dafny.MultiSetOf(_918_v66)
			var _920_v68 _dafny.Map
			_ = _920_v68
			_920_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _916_v64)
			var _921_v69 _dafny.Sequence
			_ = _921_v69
			_921_v69 = _dafny.SeqOf(_916_v64)
			r2 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
				if Companion_Default___.Fm1(_915_v63, true, Companion_Default___.Fm24(_dafny.IntOfInt64(731), globalState), _882_v40, globalState) {
					return _dafny.SeqOf(_916_v64)
				}
				return _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_916_v64, _916_v64), (Companion_Default___.SafeIndex((_919_v67).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_916_v64, _916_v64)).Cardinality()))).Uint32(), (func() *C1 {
					if (_920_v68).Contains(_dafny.IntOfInt64(72)) {
						return (_920_v68).Get(_dafny.IntOfInt64(72)).(*C1)
					}
					return _916_v64
				})())
			})(), _dafny.Companion_Sequence_.Concatenate(_921_v69, _921_v69))).Cardinality())
			var _922_v70 _dafny.Map
			_ = _922_v70
			_922_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
			r1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.SeqOf(_882_v40, _882_v40, (_917_v65).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_917_v65).Cardinality()))).Uint32()).(bool), _882_v40)).Cardinality()), (_this).F6()), ((func() _dafny.Int {
				if (_922_v70).Contains((_this).F6()) {
					return (_922_v70).Get((_this).F6()).(_dafny.Int)
				}
				return (_dafny.Zero).Minus(_dafny.IntOfUint32((_907_v55).Cardinality()))
			})()).Times((_this).F6()))
		}
		var _923_v71 _dafny.Array
		_ = _923_v71
		var _len0_26 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_26
		var _nw148 _dafny.Array
		_ = _nw148
		if _len0_26.Cmp(_dafny.Zero) == 0 {
			_nw148 = _dafny.NewArray(_len0_26)
		} else {
			var _init26 func(_dafny.Int) bool = func(_924_i13 _dafny.Int) bool {
				return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("igm"), _dafny.UnicodeSeqOfUtf8Bytes("sgxxap"))
			}
			_ = _init26
			var _element0_26 = _init26(_dafny.Zero)
			_ = _element0_26
			_nw148 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
			(_nw148).ArraySet1(_element0_26, 0)
			var _nativeLen0_26 = (_len0_26).Int()
			_ = _nativeLen0_26
			for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
				(_nw148).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
			}
		}
		_923_v71 = _nw148
		for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_923_v71), 0))); ; {
			_guard_loop_3, _ok50 := _iter50()
			if !_ok50 {
				break
			}
			var _925_i12 _dafny.Int
			_925_i12 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_925_i12).Sign() != -1) && ((_925_i12).Cmp(_dafny.ArrayLen((_923_v71), 0)) < 0)) {
				(_923_v71).ArraySet1(!((_dafny.SetOf(Companion_D7_.Create_DC17_(_dafny.SeqOf(_882_v40, _882_v40)), Companion_D7_.Create_DC17_(_dafny.SeqOf(_882_v40, _882_v40)))).Union(_dafny.SetOf(Companion_D7_.Create_DC17_(_dafny.SeqOf(_882_v40))))).Equals(_dafny.SetOf(Companion_D7_.Create_DC17_(_dafny.SeqOf(_882_v40)))), (_925_i12).Int())
			}
		}
		var _926_v72 _dafny.Sequence
		_ = _926_v72
		_926_v72 = _dafny.UnicodeSeqOfUtf8Bytes("gvljalgu")
		var _927_v73 D3
		_ = _927_v73
		_927_v73 = Companion_D3_.Create_DC7_(_926_v72)
		var _source14 D3 = Companion_D3_.Create_DC9_((func() D3 {
			if _882_v40 {
				return Companion_D3_.Create_DC9_(_927_v73)
			}
			return _927_v73
		})())
		_ = _source14
		if _source14.Is_DC8() {
			var _928___mcc_h7 D1 = _source14.Get_().(D3_DC8).Cf14
			_ = _928___mcc_h7
			var _929___mcc_h8 _dafny.Array = _source14.Get_().(D3_DC8).Cf15
			_ = _929___mcc_h8
			var _930___mcc_h9 _dafny.Int = _source14.Get_().(D3_DC8).Cf16
			_ = _930___mcc_h9
			var _931_cf16 _dafny.Int = _930___mcc_h9
			_ = _931_cf16
			var _932_cf15 _dafny.Array = _929___mcc_h8
			_ = _932_cf15
			var _933_cf14 D1 = _928___mcc_h7
			_ = _933_cf14
			r1 = Companion_Default___.SafeDivisionInt(_931_cf16, _931_cf16)
			var _934_v74 _dafny.Array
			_ = _934_v74
			_934_v74 = _923_v71
			var _rhs133 _dafny.Array = (_934_v74)
			_ = _rhs133
			var _rhs134 _dafny.Array = _932_cf15
			_ = _rhs134
			_923_v71 = _rhs133
			_932_cf15 = _rhs134
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_923_v71), 0))
			_ = _index128
			(_923_v71).ArraySet1(_882_v40, (_index128).Int())
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(981), _dafny.ArrayLen((_923_v71), 0))
			_ = _index129
			(_923_v71).ArraySet1(_882_v40, (_index129).Int())
			var _935_v75 *C1
			_ = _935_v75
			var _nw149 *C1 = New_C1_()
			_ = _nw149
			_nw149.Ctor__()
			_935_v75 = _nw149
		} else if _source14.Is_DC7() {
			var _936___mcc_h10 _dafny.Sequence = _source14.Get_().(D3_DC7).Cf13
			_ = _936___mcc_h10
			var _937_cf13 _dafny.Sequence = _936___mcc_h10
			_ = _937_cf13
			var _938_v76 _dafny.Array
			_ = _938_v76
			var _nw150 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw150
			_938_v76 = _nw150
			var _939_v77 D3
			_ = _939_v77
			_939_v77 = Companion_D3_.Create_DC7_(_dafny.UnicodeSeqOfUtf8Bytes("dfgoq"))
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_938_v76), 0))
			_ = _index130
			(_938_v76).ArraySet1(_dafny.IntOfUint32(((_939_v77).Dtor_cf13()).Cardinality()), (_index130).Int())
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_938_v76), 0))
			_ = _index131
			(_938_v76).ArraySet1((_this).F6(), (_index131).Int())
			(globalState).F5 = _882_v40
			(globalState).F5 = (_882_v40) || (_882_v40)
			var _940_v78 _dafny.Array
			_ = _940_v78
			var _len0_27 _dafny.Int = _dafny.IntOfInt64(11)
			_ = _len0_27
			var _nw151 _dafny.Array
			_ = _nw151
			if _len0_27.Cmp(_dafny.Zero) == 0 {
				_nw151 = _dafny.NewArray(_len0_27)
			} else {
				var _init27 func(_dafny.Int) _dafny.Map = func(_941_i14 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), false)
				}
				_ = _init27
				var _element0_27 = _init27(_dafny.Zero)
				_ = _element0_27
				_nw151 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
				(_nw151).ArraySet1(_element0_27, 0)
				var _nativeLen0_27 = (_len0_27).Int()
				_ = _nativeLen0_27
				for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
					(_nw151).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
				}
			}
			_940_v78 = _nw151
			var _942_v79 D14
			_ = _942_v79
			_942_v79 = Companion_D14_.Create_DC35_(_940_v78)
			var _943_v80 _dafny.Array
			_ = _943_v80
			var _nwElement0_31 _dafny.Array = _940_v78
			_ = _nwElement0_31
			var _nw152 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(16))
			_ = _nw152
			(_nw152).ArraySet1(_nwElement0_31, 0)
			(_nw152).ArraySet1(_940_v78, 1)
			(_nw152).ArraySet1(_940_v78, 2)
			(_nw152).ArraySet1(_940_v78, 3)
			(_nw152).ArraySet1(_940_v78, 4)
			(_nw152).ArraySet1((func() _dafny.Array {
				if _882_v40 {
					return _940_v78
				}
				return _940_v78
			})(), 5)
			(_nw152).ArraySet1(_940_v78, 6)
			(_nw152).ArraySet1(_940_v78, 7)
			(_nw152).ArraySet1((_942_v79).Dtor_cf67(), 8)
			(_nw152).ArraySet1(_940_v78, 9)
			(_nw152).ArraySet1(_940_v78, 10)
			(_nw152).ArraySet1(_940_v78, 11)
			(_nw152).ArraySet1(_940_v78, 12)
			(_nw152).ArraySet1(_940_v78, 13)
			(_nw152).ArraySet1(_940_v78, 14)
			(_nw152).ArraySet1((func() _dafny.Array {
				if _882_v40 {
					return _940_v78
				}
				return _940_v78
			})(), 15)
			_943_v80 = _nw152
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_943_v80), 0))
			_ = _index132
			(_943_v80).ArraySet1(_940_v78, (_index132).Int())
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_943_v80), 0))
			_ = _index133
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_938_v76), 0))
			_ = _index134
			var _rhs135 bool = !(_882_v40)
			_ = _rhs135
			var _rhs136 _dafny.Int = (_938_v76).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_938_v76), 0))).Int()).(_dafny.Int)
			_ = _rhs136
			var _rhs137 _dafny.Array = _940_v78
			_ = _rhs137
			var _rhs138 _dafny.Int = (_this).F6()
			_ = _rhs138
			var _lhs96 *GlobalState = globalState
			_ = _lhs96
			var _lhs97 *GlobalState = globalState
			_ = _lhs97
			var _lhs98 _dafny.Array = _943_v80
			_ = _lhs98
			var _lhs99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(201), _dafny.ArrayLen((_943_v80), 0))
			_ = _lhs99
			var _lhs100 _dafny.Array = _938_v76
			_ = _lhs100
			var _lhs101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(412), _dafny.ArrayLen((_938_v76), 0))
			_ = _lhs101
			_lhs96.F1 = _rhs135
			_lhs97.F2 = _rhs136
			(_lhs98).ArraySet1(_rhs137, (_lhs99).Int())
			(_lhs100).ArraySet1(_rhs138, (_lhs101).Int())
		} else {
			var _944___mcc_h11 D3 = _source14.Get_().(D3_DC9).Cf17
			_ = _944___mcc_h11
			var _945_cf17 D3 = _944___mcc_h11
			_ = _945_cf17
			var _946_v81 _dafny.Sequence
			_ = _946_v81
			_946_v81 = _dafny.SeqOf(_882_v40, !(_882_v40))
			var _947_v82 D12
			_ = _947_v82
			_947_v82 = Companion_D12_.Create_DC28_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _dafny.IntOfUint32((_946_v81).Cardinality())))
			_947_v82 = _947_v82
			var _948_v83 _dafny.Set
			_ = _948_v83
			_948_v83 = _dafny.SetOf(_882_v40, _882_v40)
			var _949_v84 *C0
			_ = _949_v84
			var _nw153 *C0 = New_C0_()
			_ = _nw153
			_nw153.Ctor__(_882_v40)
			_949_v84 = _nw153
			var _950_v85 *C0
			_ = _950_v85
			_950_v85 = _949_v84
			var _951_v86 _dafny.Array
			_ = _951_v86
			var _nwElement0_32 *C0 = _949_v84
			_ = _nwElement0_32
			var _nw154 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(18))
			_ = _nw154
			(_nw154).ArraySet1(_nwElement0_32, 0)
			(_nw154).ArraySet1(_949_v84, 1)
			(_nw154).ArraySet1(_949_v84, 2)
			(_nw154).ArraySet1(_949_v84, 3)
			(_nw154).ArraySet1(_949_v84, 4)
			(_nw154).ArraySet1(_949_v84, 5)
			(_nw154).ArraySet1(_949_v84, 6)
			(_nw154).ArraySet1(_949_v84, 7)
			(_nw154).ArraySet1(_949_v84, 8)
			(_nw154).ArraySet1(_949_v84, 9)
			(_nw154).ArraySet1(_949_v84, 10)
			(_nw154).ArraySet1(_949_v84, 11)
			(_nw154).ArraySet1(_949_v84, 12)
			(_nw154).ArraySet1(_949_v84, 13)
			(_nw154).ArraySet1((_950_v85), 14)
			(_nw154).ArraySet1(_949_v84, 15)
			(_nw154).ArraySet1(_949_v84, 16)
			(_nw154).ArraySet1(_949_v84, 17)
			_951_v86 = _nw154
			var _952_v87 _dafny.Map
			_ = _952_v87
			_952_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_948_v83).Cardinality(), _951_v86)
			_952_v87 = (_952_v87).Update((_this).F6(), _951_v86)
			var _953_v88 _dafny.Set
			_ = _953_v88
			_953_v88 = _dafny.SetOf((_this).F6())
			var _954_v89 _dafny.Sequence
			_ = _954_v89
			_954_v89 = _dafny.SeqOf((_this).F6(), (_dafny.Zero).Minus((_this).F6()), (_this).F6(), (_this).F6(), (_this).F6())
			var _955_v90 *C3
			_ = _955_v90
			var _nw155 *C3 = New_C3_()
			_ = _nw155
			_nw155.Ctor__(Companion_Default___.SafeModuloInt((_this).F6(), (_953_v88).Cardinality()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F6()), _954_v89), _dafny.IntOfInt64(918))
			_955_v90 = _nw155
			var _956_v91 *C0
			_ = _956_v91
			var _nw156 *C0 = New_C0_()
			_ = _nw156
			_nw156.Ctor__((func() bool {
				if _882_v40 {
					return _882_v40
				}
				return _882_v40
			})())
			_956_v91 = _nw156
		}
		var _957_v92 _dafny.Set
		_ = _957_v92
		_957_v92 = _dafny.SetOf(_882_v40)
		var _958_v93 D6
		_ = _958_v93
		_958_v93 = Companion_D6_.Create_DC15_(_957_v92)
		var _959_i15 _dafny.Int
		_ = _959_i15
		_959_i15 = _dafny.Zero
		{
			var _pat_let_tv29 = _882_v40
			_ = _pat_let_tv29
			var _pat_let_tv30 = _882_v40
			_ = _pat_let_tv30
			for func(_source15 D6) bool {
				if _source15.Is_DC16() {
					var _969___mcc_h12 _dafny.Int = _source15.Get_().(D6_DC16).Cf27
					_ = _969___mcc_h12
					var _970_cf27 _dafny.Int = _969___mcc_h12
					_ = _970_cf27
					return ((_this).F6()).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv29, _dafny.SeqOf(_dafny.IntOfInt64(13)))).Cardinality()) > 0
				} else {
					var _971___mcc_h13 _dafny.Set = _source15.Get_().(D6_DC15).Cf26
					_ = _971___mcc_h13
					var _972_cf26 _dafny.Set = _971___mcc_h13
					_ = _972_cf26
					return _pat_let_tv30
				}
			}(_958_v93) {
				{
					if (_959_i15).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_959_i15 = (_959_i15).Plus(_dafny.One)
					var _960_v94 _dafny.Sequence
					_ = _960_v94
					_960_v94 = _dafny.SeqOf((_this).F6())
					var _961_v95 D13
					_ = _961_v95
					_961_v95 = Companion_D13_.Create_DC32_(_dafny.IntOfUint32((_960_v94).Cardinality()), _882_v40, _882_v40, _882_v40, _882_v40)
					var _962_v96 _dafny.Sequence
					_ = _962_v96
					_962_v96 = _dafny.SeqOf((_961_v95).Dtor_cf56())
					var _963_v97 *C3
					_ = _963_v97
					var _nw157 *C3 = New_C3_()
					_ = _nw157
					_nw157.Ctor__((_this).F6(), _960_v94, (_this).F6())
					_963_v97 = _nw157
					(globalState).F0 = Companion_Default___.SafeDivisionInt((_this).F6(), (_dafny.IntOfUint32((_926_v72).Cardinality())).Times((_963_v97).F9()))
					var _964_v98 _dafny.Array
					_ = _964_v98
					var _len0_28 _dafny.Int = _dafny.IntOfInt64(23)
					_ = _len0_28
					var _nw158 _dafny.Array
					_ = _nw158
					if _len0_28.Cmp(_dafny.Zero) == 0 {
						_nw158 = _dafny.NewArray(_len0_28)
					} else {
						var _init28 func(_dafny.Int) _dafny.Int = (func(_965_v92 _dafny.Set) func(_dafny.Int) _dafny.Int {
							return func(_966_i16 _dafny.Int) _dafny.Int {
								return Companion_Default___.SafeDivisionInt(_966_i16, (_965_v92).Cardinality())
							}
						})(_957_v92)
						_ = _init28
						var _element0_28 = _init28(_dafny.Zero)
						_ = _element0_28
						_nw158 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
						(_nw158).ArraySet1(_element0_28, 0)
						var _nativeLen0_28 = (_len0_28).Int()
						_ = _nativeLen0_28
						for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
							(_nw158).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
						}
					}
					_964_v98 = _nw158
					var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_964_v98), 0))
					_ = _index135
					(_964_v98).ArraySet1((_this).F6(), (_index135).Int())
					var _967_v99 _dafny.Map
					_ = _967_v99
					_967_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, (_this).F6())
					var _968_v100 D12
					_ = _968_v100
					_968_v100 = Companion_D12_.Create_DC28_(_967_v99)
					var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(473), _dafny.ArrayLen((_964_v98), 0))
					_ = _index136
					(_964_v98).ArraySet1(((_968_v100).Dtor_cf51()).Cardinality(), (_index136).Int())
					(globalState).F1 = _882_v40
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		var _973_i17 _dafny.Int
		_ = _973_i17
		_973_i17 = _dafny.Zero
		{
			for false {
				{
					if (_973_i17).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_973_i17 = (_973_i17).Plus(_dafny.One)
					var _974_v101 _dafny.CodePoint
					_ = _974_v101
					_974_v101 = _dafny.CodePoint('j')
					var _975_v102 _dafny.Array
					_ = _975_v102
					var _nwElement0_33 _dafny.CodePoint = _974_v101
					_ = _nwElement0_33
					var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(19))
					_ = _nw159
					(_nw159).ArraySet1CodePoint(_nwElement0_33, 0)
					(_nw159).ArraySet1CodePoint(_dafny.CodePoint('l'), 1)
					(_nw159).ArraySet1CodePoint(_974_v101, 2)
					(_nw159).ArraySet1CodePoint(_974_v101, 3)
					(_nw159).ArraySet1CodePoint(_974_v101, 4)
					(_nw159).ArraySet1CodePoint(_974_v101, 5)
					(_nw159).ArraySet1CodePoint(_974_v101, 6)
					(_nw159).ArraySet1CodePoint(_974_v101, 7)
					(_nw159).ArraySet1CodePoint(_974_v101, 8)
					(_nw159).ArraySet1CodePoint(_974_v101, 9)
					(_nw159).ArraySet1CodePoint(_974_v101, 10)
					(_nw159).ArraySet1CodePoint(_974_v101, 11)
					(_nw159).ArraySet1CodePoint((func() _dafny.CodePoint {
						if false {
							return _974_v101
						}
						return _974_v101
					})(), 12)
					(_nw159).ArraySet1CodePoint(_974_v101, 13)
					(_nw159).ArraySet1CodePoint(_974_v101, 14)
					(_nw159).ArraySet1CodePoint(_974_v101, 15)
					(_nw159).ArraySet1CodePoint(_974_v101, 16)
					(_nw159).ArraySet1CodePoint(_974_v101, 17)
					(_nw159).ArraySet1CodePoint(_974_v101, 18)
					_975_v102 = _nw159
					var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_975_v102), 0))
					_ = _index137
					(_975_v102).ArraySet1CodePoint(_dafny.CodePoint('m'), (_index137).Int())
					var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_975_v102), 0))
					_ = _index138
					(_975_v102).ArraySet1CodePoint(_974_v101, (_index138).Int())
					var _976_v103 _dafny.Sequence
					_ = _976_v103
					_976_v103 = _dafny.SeqOf(_882_v40, _882_v40)
					var _977_v104 _dafny.MultiSet
					_ = _977_v104
					_977_v104 = _dafny.MultiSetOf(_882_v40)
					var _978_v105 _dafny.Sequence
					_ = _978_v105
					_978_v105 = _dafny.SeqOf((_this).F6())
					var _979_v106 D1
					_ = _979_v106
					_979_v106 = Companion_D1_.Create_DC4_(_978_v105)
					var _980_v107 _dafny.Array
					_ = _980_v107
					var _nw160 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
					_ = _nw160
					_980_v107 = _nw160
					var _981_v108 D3
					_ = _981_v108
					_981_v108 = Companion_D3_.Create_DC8_(_979_v106, _980_v107, (_this).F6())
					var _982_v109 _dafny.Map
					_ = _982_v109
					_982_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_882_v40, _882_v40)
					var _983_v110 _dafny.Map
					_ = _983_v110
					_983_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_982_v109).Cardinality())
					var _984_v111 _dafny.MultiSet
					_ = _984_v111
					_984_v111 = _dafny.MultiSetOf(_dafny.IntOfInt64(375), (_982_v109).Cardinality())
					var _pat_let_tv31 = _984_v111
					_ = _pat_let_tv31
					var _985_v112 _dafny.Array
					_ = _985_v112
					var _nwElement0_34 _dafny.Int = (_this).F6()
					_ = _nwElement0_34
					var _nw161 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(23))
					_ = _nw161
					(_nw161).ArraySet1(_nwElement0_34, 0)
					(_nw161).ArraySet1(_dafny.IntOfUint32((_976_v103).Cardinality()), 1)
					(_nw161).ArraySet1((_this).F6(), 2)
					(_nw161).ArraySet1((_this).F6(), 3)
					(_nw161).ArraySet1(_dafny.IntOfInt64(364), 4)
					(_nw161).ArraySet1((_977_v104).Cardinality(), 5)
					(_nw161).ArraySet1((_this).F6(), 6)
					(_nw161).ArraySet1((_dafny.MultiSetFromSeq(_978_v105)).Cardinality(), 7)
					(_nw161).ArraySet1((_this).F6(), 8)
					(_nw161).ArraySet1((_this).F6(), 9)
					(_nw161).ArraySet1((_this).F6(), 10)
					(_nw161).ArraySet1(((_this).F6()).Times((_this).F6()), 11)
					(_nw161).ArraySet1(Companion_Default___.Fm15((_981_v108).Dtor_cf16(), _882_v40, globalState), 12)
					(_nw161).ArraySet1((_this).F6(), 13)
					(_nw161).ArraySet1(((_983_v110).Cardinality()).Minus((_977_v104).Cardinality()), 14)
					(_nw161).ArraySet1(_dafny.IntOfUint32((_976_v103).Cardinality()), 15)
					(_nw161).ArraySet1(Companion_Default___.SafeModuloInt((_this).F6(), ((func(_pat_let32_0 D16) D16 {
						return func(_986_dt__update__tmp_h2 D16) D16 {
							return func(_pat_let33_0 _dafny.MultiSet) D16 {
								return func(_987_dt__update_hcf73_h0 _dafny.MultiSet) D16 {
									return Companion_D16_.Create_DC38_(_987_dt__update_hcf73_h0)
								}(_pat_let33_0)
							}(_pat_let_tv31)
						}(_pat_let32_0)
					}(Companion_D16_.Create_DC38_(_dafny.MultiSetFromSeq(_dafny.SeqOf((_this).F6()))))).Dtor_cf73()).Cardinality()), 16)
					(_nw161).ArraySet1((_this).F6(), 17)
					(_nw161).ArraySet1((_this).F6(), 18)
					(_nw161).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_978_v105, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(235))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg44 _dafny.Int) interface{} {
							return coer44(arg44)
						}
					}(func(_988_i18 _dafny.Int) _dafny.Int {
						return (_this).F6()
					})))).Cardinality()), 19)
					(_nw161).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F6(), (_this).F6()), 20)
					(_nw161).ArraySet1((_this).F6(), 21)
					(_nw161).ArraySet1((_dafny.Zero).Minus((_this).F6()), 22)
					_985_v112 = _nw161
					var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))
					_ = _index139
					(_985_v112).ArraySet1(((_this).F6()).Minus((_dafny.Zero).Minus((_this).F6())), (_index139).Int())
					var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))
					_ = _index140
					(_985_v112).ArraySet1((_977_v104).Cardinality(), (_index140).Int())
					var _989_v113 _dafny.Array
					_ = _989_v113
					var _len0_29 _dafny.Int = _dafny.IntOfInt64(27)
					_ = _len0_29
					var _nw162 _dafny.Array
					_ = _nw162
					if _len0_29.Cmp(_dafny.Zero) == 0 {
						_nw162 = _dafny.NewArray(_len0_29)
					} else {
						var _init29 func(_dafny.Int) _dafny.Map = (func(_990_v102 _dafny.Array, _991_v40 bool) func(_dafny.Int) _dafny.Map {
							return func(_992_i19 _dafny.Int) _dafny.Map {
								return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_990_v102).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_990_v102), 0))).Int()), _991_v40)
							}
						})(_975_v102, _882_v40)
						_ = _init29
						var _element0_29 = _init29(_dafny.Zero)
						_ = _element0_29
						_nw162 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
						(_nw162).ArraySet1(_element0_29, 0)
						var _nativeLen0_29 = (_len0_29).Int()
						_ = _nativeLen0_29
						for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
							(_nw162).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
						}
					}
					_989_v113 = _nw162
					var _993_v114 D14
					_ = _993_v114
					_993_v114 = Companion_D14_.Create_DC35_(_989_v113)
					var _994_v115 _dafny.Sequence
					_ = _994_v115
					_994_v115 = _dafny.SeqOf(_993_v114)
					if !_dafny.Companion_Sequence_.Equal((func() _dafny.Sequence {
						if false {
							return _994_v115
						}
						return _994_v115
					})(), _dafny.SeqOf(_993_v114)) {
						var _995_v117 _dafny.Map
						_ = _995_v117
						_995_v117 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_975_v102).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_975_v102), 0))).Int()), (_975_v102).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(37), _dafny.ArrayLen((_975_v102), 0))).Int()))
						(globalState).F1 = (_dafny.IntOfInt64(37)).Cmp((_dafny.Zero).Minus(((func() _dafny.Map {
							var _coll47 = _dafny.NewMapBuilder()
							_ = _coll47
							for _iter51 := _dafny.Iterate((_995_v117).Keys().Elements()); ; {
								_compr_47, _ok51 := _iter51()
								if !_ok51 {
									break
								}
								var _996_v116 _dafny.CodePoint
								_996_v116 = interface{}(_compr_47).(_dafny.CodePoint)
								if (_995_v117).Contains(_996_v116) {
									_coll47.Add(_996_v116, _882_v40)
								}
							}
							return _coll47.ToMap()
						}()).Cardinality()).Minus(_dafny.IntOfInt64(788)))) != 0
						_926_v72 = _926_v72
						var _997_v118 _dafny.Map
						_ = _997_v118
						_997_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_985_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))).Int()).(_dafny.Int), _976_v103)
						_997_v118 = _997_v118
						var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_923_v71), 0))
						_ = _index141
						(_923_v71).ArraySet1(_882_v40, (_index141).Int())
						var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_923_v71), 0))
						_ = _index142
						(_923_v71).ArraySet1(_882_v40, (_index142).Int())
						var _998_v120 D0
						_ = _998_v120
						_998_v120 = Companion_D0_.Create_DC0_(func() _dafny.Map {
							var _coll48 = _dafny.NewMapBuilder()
							_ = _coll48
							for _iter52 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(934), _dafny.IntOfInt64(861))); ; {
								_compr_48, _ok52 := _iter52()
								if !_ok52 {
									break
								}
								var _999_v119 _dafny.Int
								_999_v119 = interface{}(_compr_48).(_dafny.Int)
								if ((_dafny.IntOfInt64(934)).Cmp(_999_v119) <= 0) && ((_999_v119).Cmp(_dafny.IntOfInt64(861)) < 0) {
									_coll48.Add((_999_v119).Plus((_985_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))).Int()).(_dafny.Int)), _957_v92)
								}
							}
							return _coll48.ToMap()
						}())
						var _1000_v121 _dafny.Array
						_ = _1000_v121
						var _len0_30 _dafny.Int = _dafny.IntOfInt64(14)
						_ = _len0_30
						var _nw163 _dafny.Array
						_ = _nw163
						if _len0_30.Cmp(_dafny.Zero) == 0 {
							_nw163 = _dafny.NewArray(_len0_30)
						} else {
							var _init30 func(_dafny.Int) bool = (func(_1001_v71 _dafny.Array) func(_dafny.Int) bool {
								return func(_1002_i20 _dafny.Int) bool {
									return (_1001_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_1001_v71), 0))).Int()).(bool)
								}
							})(_923_v71)
							_ = _init30
							var _element0_30 = _init30(_dafny.Zero)
							_ = _element0_30
							_nw163 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
							(_nw163).ArraySet1(_element0_30, 0)
							var _nativeLen0_30 = (_len0_30).Int()
							_ = _nativeLen0_30
							for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
								(_nw163).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
							}
						}
						_1000_v121 = _nw163
						var _1003_v122 _dafny.Sequence
						_ = _1003_v122
						_1003_v122 = _dafny.SeqOf(_1000_v121)
						var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))
						_ = _index143
						(_985_v112).ArraySet1(((_dafny.Zero).Minus((_985_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))).Int()).(_dafny.Int))).Minus((_dafny.Zero).Minus(((Companion_Default___.Fm35((_985_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))).Int()).(_dafny.Int), globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_985_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))).Int()).(_dafny.Int), Companion_Default___.Fm1(_998_v120, (_923_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_923_v71), 0))).Int()).(bool), _dafny.IntOfUint32((_1003_v122).Cardinality()), (_923_v71).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(529), _dafny.ArrayLen((_923_v71), 0))).Int()).(bool), globalState)))).Cardinality())), (_index143).Int())
					} else {
						(globalState).F5 = _882_v40
						_978_v105 = _978_v105
						var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))
						_ = _index144
						var _rhs139 bool = _882_v40
						_ = _rhs139
						var _rhs140 _dafny.Int = (func() _dafny.Int {
							if ((_this).F6()).Cmp(_dafny.IntOfUint32((_926_v72).Cardinality())) != 0 {
								return (_985_v112).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))).Int()).(_dafny.Int)
							}
							return ((_this).F6()).Minus((_this).F6())
						})()
						_ = _rhs140
						var _rhs141 bool = (_957_v92).IsProperSubsetOf((_957_v92).Union(_957_v92))
						_ = _rhs141
						var _rhs142 _dafny.Int = ((_957_v92).Intersection(_957_v92)).Cardinality()
						_ = _rhs142
						var _rhs143 _dafny.Int = ((_dafny.Zero).Minus((_this).F6())).Plus((_this).F6())
						_ = _rhs143
						var _lhs102 _dafny.Array = _985_v112
						_ = _lhs102
						var _lhs103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))
						_ = _lhs103
						var _lhs104 *GlobalState = globalState
						_ = _lhs104
						_882_v40 = _rhs139
						(_lhs102).ArraySet1(_rhs140, (_lhs103).Int())
						_882_v40 = _rhs141
						r2 = _rhs142
						_lhs104.F0 = _rhs143
						(globalState).F2 = (_this).F6()
						var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(478), _dafny.ArrayLen((_985_v112), 0))
						_ = _index145
						(_985_v112).ArraySet1((_this).F6(), (_index145).Int())
					}
					var _1004_v123 _dafny.Sequence
					_ = _1004_v123
					_1004_v123 = _dafny.SeqOf((_979_v106).Dtor_cf11())
					_1004_v123 = _1004_v123
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		r0 = (_this).F6()
		r1 = (_dafny.Zero).Minus((_this).F6())
		r2 = _dafny.IntOfInt64(-773)
		var _1005_v124 _dafny.CodePoint
		_ = _1005_v124
		_1005_v124 = _dafny.CodePoint('a')
		r3 = _1005_v124
		return r0, r1, r2, r3
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f7 bool
	_f8 _dafny.CodePoint
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f7 = false
	_this._f8 = _dafny.CodePoint('D')
	return &_this
}

type CompanionStruct_C5_ struct {
}

var Companion_C5_ = CompanionStruct_C5_{}

func (_this *C5) Equals(other *C5) bool {
	return _this == other
}

func (_this *C5) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C5)
	return ok && _this.Equals(other)
}

func (*C5) String() string {
	return "_module.C5"
}

func Type_C5_() _dafny.TypeDescriptor {
	return type_C5_{}
}

type type_C5_ struct {
}

func (_this type_C5_) Default() interface{} {
	return (*C5)(nil)
}

func (_this type_C5_) String() string {
	return "main.C5"
}
func (_this *C5) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) Ctor__(f7 bool, f8 _dafny.CodePoint) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
	}
}
func (_this *C5) Fm5(globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((func(_source16 D0) _dafny.Sequence {
			if _source16.Is_DC1() {
				var _1006___mcc_h0 _dafny.Int = _source16.Get_().(D0_DC1).Cf1
				_ = _1006___mcc_h0
				var _1007___mcc_h1 bool = _source16.Get_().(D0_DC1).Cf2
				_ = _1007___mcc_h1
				var _1008___mcc_h2 _dafny.Int = _source16.Get_().(D0_DC1).Cf3
				_ = _1008___mcc_h2
				var _1009___mcc_h3 bool = _source16.Get_().(D0_DC1).Cf4
				_ = _1009___mcc_h3
				var _1010_cf4 bool = _1009___mcc_h3
				_ = _1010_cf4
				var _1011_cf3 _dafny.Int = _1008___mcc_h2
				_ = _1011_cf3
				var _1012_cf2 bool = _1007___mcc_h1
				_ = _1012_cf2
				var _1013_cf1 _dafny.Int = _1006___mcc_h0
				_ = _1013_cf1
				return _dafny.Companion_Sequence_.Concatenate((Companion_D1_.Create_DC4_(_dafny.SeqOf(_1011_cf3, _1013_cf1))).Dtor_cf11(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(581))).Uint32(), func(coer45 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg45 _dafny.Int) interface{} {
						return coer45(arg45)
					}
				}((func(_1014_cf3 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1015_i0 _dafny.Int) _dafny.Int {
						return (_dafny.Zero).Minus(_1014_cf3)
					}
				})(_1011_cf3))))
			} else if _source16.Is_DC2() {
				var _1016___mcc_h4 _dafny.Int = _source16.Get_().(D0_DC2).Cf5
				_ = _1016___mcc_h4
				var _1017___mcc_h5 _dafny.CodePoint = _source16.Get_().(D0_DC2).Cf6
				_ = _1017___mcc_h5
				var _1018___mcc_h6 _dafny.Int = _source16.Get_().(D0_DC2).Cf7
				_ = _1018___mcc_h6
				var _1019___mcc_h7 _dafny.Int = _source16.Get_().(D0_DC2).Cf8
				_ = _1019___mcc_h7
				var _1020___mcc_h8 _dafny.Int = _source16.Get_().(D0_DC2).Cf9
				_ = _1020___mcc_h8
				var _1021_cf9 _dafny.Int = _1020___mcc_h8
				_ = _1021_cf9
				var _1022_cf8 _dafny.Int = _1019___mcc_h7
				_ = _1022_cf8
				var _1023_cf7 _dafny.Int = _1018___mcc_h6
				_ = _1023_cf7
				var _1024_cf6 _dafny.CodePoint = _1017___mcc_h5
				_ = _1024_cf6
				var _1025_cf5 _dafny.Int = _1016___mcc_h4
				_ = _1025_cf5
				return _dafny.SeqOf((_dafny.SetOf((_this).F7())).Cardinality(), _1025_cf5)
			} else if _source16.Is_DC0() {
				var _1026___mcc_h9 _dafny.Map = _source16.Get_().(D0_DC0).Cf0
				_ = _1026___mcc_h9
				var _1027_cf0 _dafny.Map = _1026___mcc_h9
				_ = _1027_cf0
				return (Companion_D1_.Create_DC4_(_dafny.SeqOf((_dafny.SetOf((_this).F7())).Cardinality()))).Dtor_cf11()
			} else {
				var _1028___mcc_h10 D0 = _source16.Get_().(D0_DC3).Cf10
				_ = _1028___mcc_h10
				var _1029_cf10 D0 = _1028___mcc_h10
				_ = _1029_cf10
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F7())).Cardinality())), _dafny.IntOfInt64(943), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ejdijcrpi")).Cardinality()), (_this).F7(), _dafny.IntOfInt64(757), (_this).F7())).Dtor_cf3())).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(-20), _dafny.IntOfUint32((_dafny.SeqOf((_this).F7())).Cardinality())))
			}
		}(Companion_D0_.Create_DC0_(func() _dafny.Map {
			var _coll49 = _dafny.NewMapBuilder()
			_ = _coll49
			for _iter53 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(714))).Elements()); ; {
				_compr_49, _ok53 := _iter53()
				if !_ok53 {
					break
				}
				var _1030_v0 _dafny.Int
				_1030_v0 = interface{}(_compr_49).(_dafny.Int)
				if (_dafny.SetOf(_dafny.IntOfInt64(714))).Contains(_1030_v0) {
					_coll49.Add((_1030_v0).Times(_dafny.IntOfInt64(322)), _dafny.SetOf((_this).F7()))
				}
			}
			return _coll49.ToMap()
		}()))).Cardinality())
	}
}
func (_this *C5) Fm6(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-724))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg46 _dafny.Int) interface{} {
				return coer46(arg46)
			}
		}(func(_1031_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rj"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(846))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg47 _dafny.Int) interface{} {
					return coer47(arg47)
				}
			}(func(_1032_i1 _dafny.Int) _dafny.CodePoint {
				return (_this).F8()
			})))
		}))).Cardinality())
	}
}
func (_this *C5) M1(p0 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 _dafny.Int = _dafny.Zero
		_ = r3
		var _1033_v0 _dafny.Array
		_ = _1033_v0
		var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
		_ = _nw164
		_1033_v0 = _nw164
		_1033_v0 = _1033_v0
		var _1034_v1 _dafny.Array
		_ = _1034_v1
		var _nw165 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
		_ = _nw165
		_1034_v1 = _nw165
		var _1035_v2 _dafny.Int
		_ = _1035_v2
		_1035_v2 = _dafny.IntOfInt64(441)
		var _1036_v3 _dafny.Map
		_ = _1036_v3
		_1036_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1035_v2, _1034_v1)
		var _1037_v4 _dafny.Set
		_ = _1037_v4
		_1037_v4 = _dafny.SetOf(_1035_v2, _1035_v2)
		var _1038_v5 _dafny.Sequence
		_ = _1038_v5
		_1038_v5 = _dafny.SeqOf(_1035_v2, _1035_v2, (_1037_v4).Cardinality())
		var _1039_v6 _dafny.Array
		_ = _1039_v6
		var _nwElement0_35 _dafny.Array = _1034_v1
		_ = _nwElement0_35
		var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(18))
		_ = _nw166
		(_nw166).ArraySet1(_nwElement0_35, 0)
		(_nw166).ArraySet1(_1034_v1, 1)
		(_nw166).ArraySet1(_1034_v1, 2)
		(_nw166).ArraySet1(_1034_v1, 3)
		(_nw166).ArraySet1(_1034_v1, 4)
		(_nw166).ArraySet1(_1034_v1, 5)
		(_nw166).ArraySet1(_1034_v1, 6)
		(_nw166).ArraySet1(_1034_v1, 7)
		(_nw166).ArraySet1(_1034_v1, 8)
		(_nw166).ArraySet1(_1034_v1, 9)
		(_nw166).ArraySet1(_1034_v1, 10)
		(_nw166).ArraySet1(_1034_v1, 11)
		(_nw166).ArraySet1(_1034_v1, 12)
		(_nw166).ArraySet1(_1034_v1, 13)
		(_nw166).ArraySet1(_1034_v1, 14)
		(_nw166).ArraySet1((func() _dafny.Array {
			if (_1036_v3).Contains((_1038_v5).Select((Companion_Default___.SafeIndex(_1035_v2, _dafny.IntOfUint32((_1038_v5).Cardinality()))).Uint32()).(_dafny.Int)) {
				return (_1036_v3).Get((_1038_v5).Select((Companion_Default___.SafeIndex(_1035_v2, _dafny.IntOfUint32((_1038_v5).Cardinality()))).Uint32()).(_dafny.Int)).(_dafny.Array)
			}
			return _1034_v1
		})(), 15)
		(_nw166).ArraySet1(_1034_v1, 16)
		(_nw166).ArraySet1(_1034_v1, 17)
		_1039_v6 = _nw166
		var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1039_v6), 0))
		_ = _index146
		(_1039_v6).ArraySet1(_1034_v1, (_index146).Int())
		var _1040_v7 D1
		_ = _1040_v7
		_1040_v7 = Companion_D1_.Create_DC4_(Companion_Default___.Fm7(globalState))
		var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))
		_ = _index147
		(_1033_v0).ArraySet1(_1035_v2, (_index147).Int())
		var _1041_v8 _dafny.Sequence
		_ = _1041_v8
		_1041_v8 = _dafny.SeqOf(false)
		var _1042_v9 _dafny.Array
		_ = _1042_v9
		_1042_v9 = _1034_v1
		var _1043_v10 D1
		_ = _1043_v10
		_1043_v10 = Companion_D1_.Create_DC5_()
		var _pat_let_tv32 = _1035_v2
		_ = _pat_let_tv32
		var _pat_let_tv33 = _1035_v2
		_ = _pat_let_tv33
		var _pat_let_tv34 = _1035_v2
		_ = _pat_let_tv34
		var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1039_v6), 0))
		_ = _index148
		var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))
		_ = _index149
		var _rhs144 _dafny.Int = (_1035_v2).Plus(_1035_v2)
		_ = _rhs144
		var _rhs145 _dafny.Int = ((_this).Fm6(_1035_v2, p0, _dafny.IntOfInt64(816), globalState)).Plus((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_1041_v8).Cardinality()))).Cardinality()))
		_ = _rhs145
		var _rhs146 _dafny.Array = (_1042_v9)
		_ = _rhs146
		var _rhs147 D1 = Companion_D1_.Create_DC4_(_1038_v5)
		_ = _rhs147
		var _rhs148 _dafny.Int = func(_source17 D1) _dafny.Int {
			if _source17.Is_DC5() {
				return (_pat_let_tv32).Minus(_pat_let_tv33)
			} else {
				var _1044___mcc_h0 _dafny.Sequence = _source17.Get_().(D1_DC4).Cf11
				_ = _1044___mcc_h0
				var _1045_cf11 _dafny.Sequence = _1044___mcc_h0
				_ = _1045_cf11
				return _pat_let_tv34
			}
		}(_1043_v10)
		_ = _rhs148
		var _lhs105 *GlobalState = globalState
		_ = _lhs105
		var _lhs106 _dafny.Array = _1039_v6
		_ = _lhs106
		var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1039_v6), 0))
		_ = _lhs107
		var _lhs108 _dafny.Array = _1033_v0
		_ = _lhs108
		var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))
		_ = _lhs109
		r3 = _rhs144
		_lhs105.F0 = _rhs145
		(_lhs106).ArraySet1(_rhs146, (_lhs107).Int())
		_1040_v7 = _rhs147
		(_lhs108).ArraySet1(_rhs148, (_lhs109).Int())
		var _1046_v11 _dafny.Array
		_ = _1046_v11
		var _nw167 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(14))
		_ = _nw167
		_1046_v11 = _nw167
		var _1047_v12 _dafny.Map
		_ = _1047_v12
		_1047_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1046_v11, _1033_v0)
		_1047_v12 = (_1047_v12).Update(_1046_v11, _1033_v0)
		var _1048_v13 _dafny.Set
		_ = _1048_v13
		_1048_v13 = _dafny.SetOf(p0)
		_1048_v13 = (func() _dafny.Set {
			if false {
				return _1048_v13
			}
			return _1048_v13
		})()
		var _1049_v14 _dafny.Array
		_ = _1049_v14
		var _nw168 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
		_ = _nw168
		_1049_v14 = _nw168
		var _1050_v15 D0
		_ = _1050_v15
		_1050_v15 = Companion_D0_.Create_DC1_((_this).Fm5(globalState), p0, (_1033_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))).Int()).(_dafny.Int), p0)
		var _1051_v16 D0
		_ = _1051_v16
		_1051_v16 = Companion_D0_.Create_DC3_(_1050_v15)
		var _1052_v17 D0
		_ = _1052_v17
		_1052_v17 = Companion_D0_.Create_DC3_(_1051_v16)
		var _1053_v18 _dafny.Map
		_ = _1053_v18
		_1053_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1033_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))).Int()).(_dafny.Int), (_this).F7()), _1052_v17)
		var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1049_v14), 0))
		_ = _index150
		(_1049_v14).ArraySet1(_1053_v18, (_index150).Int())
		var _1054_v19 _dafny.CodePoint
		_ = _1054_v19
		_1054_v19 = _dafny.CodePoint('d')
		var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1049_v14), 0))
		_ = _index151
		var _rhs149 _dafny.Map = _1053_v18
		_ = _rhs149
		var _rhs150 _dafny.CodePoint = Companion_Default___.Fm8((_this).F8(), (func() bool {
			if p0 {
				return p0
			}
			return (_this).F7()
		})(), (_1033_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))).Int()).(_dafny.Int), (_1033_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(425), _dafny.ArrayLen((_1033_v0), 0))).Int()).(_dafny.Int), globalState)
		_ = _rhs150
		var _lhs110 _dafny.Array = _1049_v14
		_ = _lhs110
		var _lhs111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(711), _dafny.ArrayLen((_1049_v14), 0))
		_ = _lhs111
		(_lhs110).ArraySet1(_rhs149, (_lhs111).Int())
		_1054_v19 = _rhs150
		var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1039_v6), 0))
		_ = _index152
		(_1039_v6).ArraySet1(_dafny.ArrayCastTo((_1039_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_1039_v6), 0))).Int())), (_index152).Int())
		r0 = _1035_v2
		var _1055_v21 _dafny.MultiSet
		_ = _1055_v21
		_1055_v21 = _dafny.MultiSetOf((_this).F7())
		var _1056_v22 _dafny.Map
		_ = _1056_v22
		_1056_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1055_v21, _1035_v2)
		var _1057_v24 _dafny.Sequence
		_ = _1057_v24
		_1057_v24 = _dafny.SeqOf((func() _dafny.Set {
			if (_this).F7() {
				return Companion_Default___.Fm9(globalState)
			}
			return func() _dafny.Set {
				var _coll50 = _dafny.NewBuilder()
				_ = _coll50
				for _iter54 := _dafny.Iterate((func() _dafny.Map {
					var _coll51 = _dafny.NewMapBuilder()
					_ = _coll51
					for _iter55 := _dafny.Iterate((_1056_v22).Keys().Elements()); ; {
						_compr_51, _ok55 := _iter55()
						if !_ok55 {
							break
						}
						var _1058_v20 _dafny.MultiSet
						_1058_v20 = interface{}(_compr_51).(_dafny.MultiSet)
						if (_1056_v22).Contains(_1058_v20) {
							_coll51.Add(_1058_v20, _1041_v8)
						}
					}
					return _coll51.ToMap()
				}()).Keys().Elements()); ; {
					_compr_50, _ok54 := _iter54()
					if !_ok54 {
						break
					}
					var _1059_v23 _dafny.MultiSet
					_1059_v23 = interface{}(_compr_50).(_dafny.MultiSet)
					if (func() _dafny.Map {
						var _coll52 = _dafny.NewMapBuilder()
						_ = _coll52
						for _iter56 := _dafny.Iterate((_1056_v22).Keys().Elements()); ; {
							_compr_52, _ok56 := _iter56()
							if !_ok56 {
								break
							}
							var _1058_v20 _dafny.MultiSet
							_1058_v20 = interface{}(_compr_52).(_dafny.MultiSet)
							if (_1056_v22).Contains(_1058_v20) {
								_coll52.Add(_1058_v20, _1041_v8)
							}
						}
						return _coll52.ToMap()
					}()).Contains(_1059_v23) {
						_coll50.Add(_1059_v23)
					}
				}
				return _coll50.ToSet()
			}()
		})())
		r1 = ((_1057_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.IntOfUint32((_1057_v24).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
		r2 = p0
		r3 = _1035_v2
		return r0, r1, r2, r3
	}
}
func (_this *C5) M2(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Sequence, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1060_v0 _dafny.Map
		_ = _1060_v0
		_1060_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SetOf((_this).F7()))
		var _1061_v1 _dafny.Map
		_ = _1061_v1
		_1061_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _1060_v0)
		var _1062_v2 D0
		_ = _1062_v2
		_1062_v2 = Companion_D0_.Create_DC0_((func() _dafny.Map {
			if (_1061_v1).Contains(p2) {
				return (_1061_v1).Get(p2).(_dafny.Map)
			}
			return _1060_v0
		})())
		var _pat_let_tv35 = p0
		_ = _pat_let_tv35
		var _pat_let_tv36 = p3
		_ = _pat_let_tv36
		var _pat_let_tv37 = p0
		_ = _pat_let_tv37
		var _pat_let_tv38 = p2
		_ = _pat_let_tv38
		var _pat_let_tv39 = p0
		_ = _pat_let_tv39
		var _pat_let_tv40 = p0
		_ = _pat_let_tv40
		var _pat_let_tv41 = p1
		_ = _pat_let_tv41
		var _pat_let_tv42 = p0
		_ = _pat_let_tv42
		var _pat_let_tv43 = p0
		_ = _pat_let_tv43
		var _source18 D0 = func(_source19 D0) D0 {
			if _source19.Is_DC1() {
				var _1063___mcc_h11 _dafny.Int = _source19.Get_().(D0_DC1).Cf1
				_ = _1063___mcc_h11
				var _1064___mcc_h12 bool = _source19.Get_().(D0_DC1).Cf2
				_ = _1064___mcc_h12
				var _1065___mcc_h13 _dafny.Int = _source19.Get_().(D0_DC1).Cf3
				_ = _1065___mcc_h13
				var _1066___mcc_h14 bool = _source19.Get_().(D0_DC1).Cf4
				_ = _1066___mcc_h14
				var _1067_cf4 bool = _1066___mcc_h14
				_ = _1067_cf4
				var _1068_cf3 _dafny.Int = _1065___mcc_h13
				_ = _1068_cf3
				var _1069_cf2 bool = _1064___mcc_h12
				_ = _1069_cf2
				var _1070_cf1 _dafny.Int = _1063___mcc_h11
				_ = _1070_cf1
				return Companion_D0_.Create_DC1_(_pat_let_tv35, true, _dafny.IntOfUint32((_pat_let_tv36).Cardinality()), _1069_cf2)
			} else if _source19.Is_DC2() {
				var _1071___mcc_h15 _dafny.Int = _source19.Get_().(D0_DC2).Cf5
				_ = _1071___mcc_h15
				var _1072___mcc_h16 _dafny.CodePoint = _source19.Get_().(D0_DC2).Cf6
				_ = _1072___mcc_h16
				var _1073___mcc_h17 _dafny.Int = _source19.Get_().(D0_DC2).Cf7
				_ = _1073___mcc_h17
				var _1074___mcc_h18 _dafny.Int = _source19.Get_().(D0_DC2).Cf8
				_ = _1074___mcc_h18
				var _1075___mcc_h19 _dafny.Int = _source19.Get_().(D0_DC2).Cf9
				_ = _1075___mcc_h19
				var _1076_cf9 _dafny.Int = _1075___mcc_h19
				_ = _1076_cf9
				var _1077_cf8 _dafny.Int = _1074___mcc_h18
				_ = _1077_cf8
				var _1078_cf7 _dafny.Int = _1073___mcc_h17
				_ = _1078_cf7
				var _1079_cf6 _dafny.CodePoint = _1072___mcc_h16
				_ = _1079_cf6
				var _1080_cf5 _dafny.Int = _1071___mcc_h15
				_ = _1080_cf5
				return Companion_D0_.Create_DC1_(_1077_cf8, (_this).F7(), _1077_cf8, (_this).F7())
			} else if _source19.Is_DC0() {
				var _1081___mcc_h20 _dafny.Map = _source19.Get_().(D0_DC0).Cf0
				_ = _1081___mcc_h20
				var _1082_cf0 _dafny.Map = _1081___mcc_h20
				_ = _1082_cf0
				return Companion_D0_.Create_DC1_(_dafny.IntOfUint32((_dafny.SeqOf(_pat_let_tv37)).Cardinality()), _pat_let_tv38, (func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter57 := _dafny.Iterate((_dafny.SeqOf(_pat_let_tv39)).Elements()); ; {
						_compr_53, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1083_v3 _dafny.Int
						_1083_v3 = interface{}(_compr_53).(_dafny.Int)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_pat_let_tv40), _1083_v3) {
							_coll53.Add((_1083_v3).Times((_dafny.SetOf((_this).F7(), (_this).F7())).Cardinality()), _pat_let_tv41)
						}
					}
					return _coll53.ToMap()
				}()).Cardinality(), (_this).F7())
			} else {
				var _1084___mcc_h21 D0 = _source19.Get_().(D0_DC3).Cf10
				_ = _1084___mcc_h21
				var _1085_cf10 D0 = _1084___mcc_h21
				_ = _1085_cf10
				return Companion_D0_.Create_DC1_(_pat_let_tv42, (_this).F7(), _pat_let_tv43, (_this).F7())
			}
		}(_1062_v2)
		_ = _source18
		if _source18.Is_DC1() {
			var _1086___mcc_h0 _dafny.Int = _source18.Get_().(D0_DC1).Cf1
			_ = _1086___mcc_h0
			var _1087___mcc_h1 bool = _source18.Get_().(D0_DC1).Cf2
			_ = _1087___mcc_h1
			var _1088___mcc_h2 _dafny.Int = _source18.Get_().(D0_DC1).Cf3
			_ = _1088___mcc_h2
			var _1089___mcc_h3 bool = _source18.Get_().(D0_DC1).Cf4
			_ = _1089___mcc_h3
			var _1090_cf4 bool = _1089___mcc_h3
			_ = _1090_cf4
			var _1091_cf3 _dafny.Int = _1088___mcc_h2
			_ = _1091_cf3
			var _1092_cf2 bool = _1087___mcc_h1
			_ = _1092_cf2
			var _1093_cf1 _dafny.Int = _1086___mcc_h0
			_ = _1093_cf1
			var _1094_v4 _dafny.MultiSet
			_ = _1094_v4
			_1094_v4 = _dafny.MultiSetOf(_1090_cf4)
			(globalState).F2 = ((_1091_cf3).Minus((_1094_v4).Cardinality())).Minus(p0)
			var _1095_v5 _dafny.Map
			_ = _1095_v5
			_1095_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
				if (_this).F7() {
					return !(p2)
				}
				return _1092_cf2
			})(), p1)
			_1095_v5 = _1095_v5
			var _1096_v6 _dafny.Sequence
			_ = _1096_v6
			_1096_v6 = _dafny.UnicodeSeqOfUtf8Bytes("ja")
			_1096_v6 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(129))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg48 _dafny.Int) interface{} {
					return coer48(arg48)
				}
			}(func(_1097_i0 _dafny.Int) _dafny.CodePoint {
				return (_this).F8()
			})), _dafny.Companion_Sequence_.Concatenate(_1096_v6, _1096_v6))
			var _1098_v7 _dafny.Set
			_ = _1098_v7
			_1098_v7 = _dafny.SetOf(_1093_cf1)
			_1095_v5 = (_1095_v5).Update(_1090_cf4, (_1098_v7).IsDisjointFrom(_dafny.SetOf(_1091_cf3)))
		} else if _source18.Is_DC2() {
			var _1099___mcc_h4 _dafny.Int = _source18.Get_().(D0_DC2).Cf5
			_ = _1099___mcc_h4
			var _1100___mcc_h5 _dafny.CodePoint = _source18.Get_().(D0_DC2).Cf6
			_ = _1100___mcc_h5
			var _1101___mcc_h6 _dafny.Int = _source18.Get_().(D0_DC2).Cf7
			_ = _1101___mcc_h6
			var _1102___mcc_h7 _dafny.Int = _source18.Get_().(D0_DC2).Cf8
			_ = _1102___mcc_h7
			var _1103___mcc_h8 _dafny.Int = _source18.Get_().(D0_DC2).Cf9
			_ = _1103___mcc_h8
			var _1104_cf9 _dafny.Int = _1103___mcc_h8
			_ = _1104_cf9
			var _1105_cf8 _dafny.Int = _1102___mcc_h7
			_ = _1105_cf8
			var _1106_cf7 _dafny.Int = _1101___mcc_h6
			_ = _1106_cf7
			var _1107_cf6 _dafny.CodePoint = _1100___mcc_h5
			_ = _1107_cf6
			var _1108_cf5 _dafny.Int = _1099___mcc_h4
			_ = _1108_cf5
			var _1109_v8 _dafny.Array
			_ = _1109_v8
			var _nw169 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw169
			_1109_v8 = _nw169
			var _1110_v9 _dafny.Array
			_ = _1110_v9
			_1110_v9 = _1109_v8
			var _source20 _dafny.Array = (func() _dafny.Array {
				if (_this).F7() {
					return _1110_v9
				}
				return _1109_v8
			})()
			_ = _source20
			var _1111___mcc_h22 _dafny.Array = _source20
			_ = _1111___mcc_h22
			var _1112_cf12 _dafny.Array = _1111___mcc_h22
			_ = _1112_cf12
			var _1113_v10 _dafny.Map
			_ = _1113_v10
			_1113_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(687), p2)
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1109_v8), 0))
			_ = _index153
			(_1109_v8).ArraySet1((func() bool {
				if (_1113_v10).Contains(_1104_cf9) {
					return (_1113_v10).Get(_1104_cf9).(bool)
				}
				return !(p1)
			})(), (_index153).Int())
			var _1114_v11 _dafny.Array
			_ = _1114_v11
			var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw170
			_1114_v11 = _nw170
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1114_v11), 0))
			_ = _index154
			(_1114_v11).ArraySet1(_1112_cf12, (_index154).Int())
			var _1115_v12 _dafny.Map
			_ = _1115_v12
			_1115_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), p1)
			var _1116_v13 _dafny.Map
			_ = _1116_v13
			_1116_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm5(globalState), _1104_cf9)
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1109_v8), 0))
			_ = _index155
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1114_v11), 0))
			_ = _index156
			var _rhs151 _dafny.Int = (_1115_v12).Cardinality()
			_ = _rhs151
			var _rhs152 bool = p2
			_ = _rhs152
			var _rhs153 bool = p1
			_ = _rhs153
			var _rhs154 _dafny.Array = _1112_cf12
			_ = _rhs154
			var _rhs155 bool = (func() bool {
				if ((func() _dafny.Int {
					if (_1116_v13).Contains(_dafny.IntOfInt64(82)) {
						return (_1116_v13).Get(_dafny.IntOfInt64(82)).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qng")).Cardinality())
				})()).Cmp(_1106_cf7) < 0 {
					return p1
				}
				return (Companion_D0_.Create_DC1_(_1106_cf7, Companion_Default___.Fm1(_1062_v2, !(p2), (_this).Fm5(globalState), p2, globalState), _1105_cf8, (_this).F7())).Dtor_cf4()
			})()
			_ = _rhs155
			var _lhs112 *GlobalState = globalState
			_ = _lhs112
			var _lhs113 _dafny.Array = _1109_v8
			_ = _lhs113
			var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(161), _dafny.ArrayLen((_1109_v8), 0))
			_ = _lhs114
			var _lhs115 *GlobalState = globalState
			_ = _lhs115
			var _lhs116 _dafny.Array = _1114_v11
			_ = _lhs116
			var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(205), _dafny.ArrayLen((_1114_v11), 0))
			_ = _lhs117
			var _lhs118 *GlobalState = globalState
			_ = _lhs118
			_lhs112.F0 = _rhs151
			(_lhs113).ArraySet1(_rhs152, (_lhs114).Int())
			_lhs115.F1 = _rhs153
			(_lhs116).ArraySet1(_rhs154, (_lhs117).Int())
			_lhs118.F1 = _rhs155
			var _1117_v14 _dafny.Int
			_ = _1117_v14
			var _1118_v15 _dafny.Int
			_ = _1118_v15
			var _1119_v16 bool
			_ = _1119_v16
			var _1120_v17 _dafny.Int
			_ = _1120_v17
			var _out4 _dafny.Int
			_ = _out4
			var _out5 _dafny.Int
			_ = _out5
			var _out6 bool
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out4, _out5, _out6, _out7 = (_this).M1(p1, globalState)
			_1117_v14 = _out4
			_1118_v15 = _out5
			_1119_v16 = _out6
			_1120_v17 = _out7
			_1118_v15 = _1108_cf5
			var _1121_v18 _dafny.Array
			_ = _1121_v18
			var _len0_31 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_31
			var _nw171 _dafny.Array
			_ = _nw171
			if _len0_31.Cmp(_dafny.Zero) == 0 {
				_nw171 = _dafny.NewArray(_len0_31)
			} else {
				var _init31 func(_dafny.Int) _dafny.Int = (func(_1122_cf7 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1123_i1 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_1123_i1, _1122_cf7)
					}
				})(_1106_cf7)
				_ = _init31
				var _element0_31 = _init31(_dafny.Zero)
				_ = _element0_31
				_nw171 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
				(_nw171).ArraySet1(_element0_31, 0)
				var _nativeLen0_31 = (_len0_31).Int()
				_ = _nativeLen0_31
				for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
					(_nw171).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
				}
			}
			_1121_v18 = _nw171
			var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1121_v18), 0))
			_ = _index157
			(_1121_v18).ArraySet1(p0, (_index157).Int())
			var _1124_v19 _dafny.Sequence
			_ = _1124_v19
			_1124_v19 = _dafny.UnicodeSeqOfUtf8Bytes("hjhrdvw")
			var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1121_v18), 0))
			_ = _index158
			var _rhs156 bool = _1119_v16
			_ = _rhs156
			var _rhs157 _dafny.Int = _1117_v14
			_ = _rhs157
			var _rhs158 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_1124_v19).Cardinality()))
			_ = _rhs158
			var _rhs159 bool = (_this).F7()
			_ = _rhs159
			var _rhs160 _dafny.Int = _1104_cf9
			_ = _rhs160
			var _lhs119 *GlobalState = globalState
			_ = _lhs119
			var _lhs120 _dafny.Array = _1121_v18
			_ = _lhs120
			var _lhs121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1121_v18), 0))
			_ = _lhs121
			var _lhs122 *GlobalState = globalState
			_ = _lhs122
			var _lhs123 *GlobalState = globalState
			_ = _lhs123
			_lhs119.F1 = _rhs156
			_1105_cf8 = _rhs157
			(_lhs120).ArraySet1(_rhs158, (_lhs121).Int())
			_lhs122.F5 = _rhs159
			_lhs123.F0 = _rhs160
			_1105_cf8 = (_1108_cf5).Times(_dafny.IntOfInt64(892))
			var _1125_v20 _dafny.Sequence
			_ = _1125_v20
			_1125_v20 = _dafny.UnicodeSeqOfUtf8Bytes("qtau")
			var _1126_v21 _dafny.Map
			_ = _1126_v21
			_1126_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1125_v20)
			var _1127_v22 _dafny.MultiSet
			_ = _1127_v22
			_1127_v22 = _dafny.MultiSetOf((func() _dafny.Sequence {
				if (_1126_v21).Contains((_this).F7()) {
					return (_1126_v21).Get((_this).F7()).(_dafny.Sequence)
				}
				return _1125_v20
			})(), _1125_v20)
			_1127_v22 = (_1127_v22).Union(_dafny.MultiSetOf(_1125_v20))
			var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_1109_v8), 0))
			_ = _index159
			(_1109_v8).ArraySet1(p2, (_index159).Int())
			var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(499), _dafny.ArrayLen((_1109_v8), 0))
			_ = _index160
			(_1109_v8).ArraySet1((_this).F7(), (_index160).Int())
		} else if _source18.Is_DC0() {
			var _1128___mcc_h9 _dafny.Map = _source18.Get_().(D0_DC0).Cf0
			_ = _1128___mcc_h9
			var _1129_cf0 _dafny.Map = _1128___mcc_h9
			_ = _1129_cf0
			var _1130_v23 _dafny.Array
			_ = _1130_v23
			var _nwElement0_36 _dafny.Int = p0
			_ = _nwElement0_36
			var _nw172 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(7))
			_ = _nw172
			(_nw172).ArraySet1(_nwElement0_36, 0)
			(_nw172).ArraySet1(p0, 1)
			(_nw172).ArraySet1(p0, 2)
			(_nw172).ArraySet1(p0, 3)
			(_nw172).ArraySet1(_dafny.IntOfInt64(813), 4)
			(_nw172).ArraySet1(p0, 5)
			(_nw172).ArraySet1(p0, 6)
			_1130_v23 = _nw172
			var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))
			_ = _index161
			(_1130_v23).ArraySet1((_dafny.Zero).Minus(p0), (_index161).Int())
			var _1131_v24 D0
			_ = _1131_v24
			_1131_v24 = Companion_D0_.Create_DC1_(p0, p2, p0, (_this).F7())
			var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))
			_ = _index162
			var _rhs161 bool = (_1131_v24).Dtor_cf2()
			_ = _rhs161
			var _rhs162 _dafny.Int = p0
			_ = _rhs162
			var _lhs124 *GlobalState = globalState
			_ = _lhs124
			var _lhs125 _dafny.Array = _1130_v23
			_ = _lhs125
			var _lhs126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))
			_ = _lhs126
			_lhs124.F5 = _rhs161
			(_lhs125).ArraySet1(_rhs162, (_lhs126).Int())
			var _1132_v25 _dafny.Array
			_ = _1132_v25
			var _nw173 _dafny.Array = _dafny.NewArray(_dafny.One)
			_ = _nw173
			_1132_v25 = _nw173
			var _1133_v26 _dafny.Sequence
			_ = _1133_v26
			_1133_v26 = _dafny.SeqOf(Companion_Default___.Fm1(Companion_D0_.Create_DC0_(_1060_v0), p1, _dafny.IntOfInt64(-909), p1, globalState), (func() bool {
				if !(p2) {
					return p1
				}
				return false
			})())
			var _rhs163 _dafny.Array = _1132_v25
			_ = _rhs163
			var _rhs164 _dafny.Sequence = _dafny.SeqOf(p2)
			_ = _rhs164
			var _rhs165 _dafny.Int = p0
			_ = _rhs165
			var _lhs127 *GlobalState = globalState
			_ = _lhs127
			_1132_v25 = _rhs163
			_1133_v26 = _rhs164
			_lhs127.F2 = _rhs165
			var _1134_v27 _dafny.Set
			_ = _1134_v27
			_1134_v27 = _dafny.SetOf(Companion_Default___.SafeModuloInt(p0, (_1130_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))).Int()).(_dafny.Int)))
			_1134_v27 = _1134_v27
			var _1135_v28 _dafny.MultiSet
			_ = _1135_v28
			_1135_v28 = _dafny.MultiSetOf((_1130_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))).Int()).(_dafny.Int), (_1130_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if !(p2) {
					return (_1130_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))).Int()).(_dafny.Int)
				}
				return (_dafny.SetOf(_1134_v27, _1134_v27, _1134_v27, _1134_v27, Companion_Default___.Fm10(_dafny.IntOfInt64(802), p1, p2, globalState))).Cardinality()
			})(), (_this).Fm5(globalState))
			var _1136_v29 _dafny.Array
			_ = _1136_v29
			var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw174
			_1136_v29 = _nw174
			var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))
			_ = _index163
			var _rhs166 _dafny.Int = (_1130_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))).Int()).(_dafny.Int)
			_ = _rhs166
			var _rhs167 bool = (Companion_Default___.Fm1(_1062_v2, false, p0, (_this).F7(), globalState)) || ((_this).F7())
			_ = _rhs167
			var _rhs168 _dafny.MultiSet = _dafny.MultiSetOf(p0, (_1130_v23).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))).Int()).(_dafny.Int))
			_ = _rhs168
			var _rhs169 _dafny.Array = _1136_v29
			_ = _rhs169
			var _lhs128 _dafny.Array = _1130_v23
			_ = _lhs128
			var _lhs129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_1130_v23), 0))
			_ = _lhs129
			var _lhs130 *GlobalState = globalState
			_ = _lhs130
			(_lhs128).ArraySet1(_rhs166, (_lhs129).Int())
			_lhs130.F5 = _rhs167
			_1135_v28 = _rhs168
			_1136_v29 = _rhs169
		} else {
			var _1137___mcc_h10 D0 = _source18.Get_().(D0_DC3).Cf10
			_ = _1137___mcc_h10
			var _1138_cf10 D0 = _1137___mcc_h10
			_ = _1138_cf10
			var _1139_v30 _dafny.Map
			_ = _1139_v30
			_1139_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(575))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}(func(_1140_i2 _dafny.Int) _dafny.CodePoint {
				return (_this).F8()
			}))).Cardinality())).Minus(p0))
			_1139_v30 = (_1139_v30).Update(!(p1), (_this).Fm5(globalState))
			var _1141_v31 _dafny.Sequence
			_ = _1141_v31
			_1141_v31 = _dafny.UnicodeSeqOfUtf8Bytes("xp")
			var _1142_v32 _dafny.Map
			_ = _1142_v32
			_1142_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfUint32((_1141_v31).Cardinality())).Times((_1139_v30).Cardinality()), !_dafny.Companion_Sequence_.Equal(p3, p3))
			_1142_v32 = (_1142_v32).Update(p0, (p0).Cmp(_dafny.IntOfInt64(723)) > 0)
			var _1143_v33 _dafny.Map
			_ = _1143_v33
			_1143_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, false)
			_1143_v33 = (_1143_v33).Update((p0).Cmp(_dafny.IntOfUint32((_1141_v31).Cardinality())) == 0, p1)
			var _1144_v34 _dafny.Array
			_ = _1144_v34
			var _nw175 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(23))
			_ = _nw175
			_1144_v34 = _nw175
			var _1145_v35 _dafny.Array
			_ = _1145_v35
			var _nw176 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw176
			_1145_v35 = _nw176
			var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_1144_v34), 0))
			_ = _index164
			(_1144_v34).ArraySet1(_1145_v35, (_index164).Int())
			var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(93), _dafny.ArrayLen((_1144_v34), 0))
			_ = _index165
			(_1144_v34).ArraySet1(_1145_v35, (_index165).Int())
		}
		if (_this).F7() {
			(globalState).F0 = (_dafny.IntOfInt64(701)).Plus(p0)
			var _1146_v36 _dafny.Sequence
			_ = _1146_v36
			_1146_v36 = _dafny.SeqOf(p3)
			(globalState).F5 = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_1146_v36, _1146_v36), _1146_v36)
			_1146_v36 = _dafny.Companion_Sequence_.Concatenate(_1146_v36, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(966))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_1147_p3 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_1148_i3 _dafny.Int) _dafny.Sequence {
					return _1147_p3
				}
			})(p3))))
			var _1149_v37 _dafny.Array
			_ = _1149_v37
			var _nw177 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw177
			_1149_v37 = _nw177
			var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1149_v37), 0))
			_ = _index166
			(_1149_v37).ArraySet1(p0, (_index166).Int())
			var _1150_v38 _dafny.Map
			_ = _1150_v38
			_1150_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p2)
			var _1151_v39 D0
			_ = _1151_v39
			_1151_v39 = Companion_D0_.Create_DC1_(p0, p1, _dafny.IntOfInt64(883), (func() bool {
				if (_1150_v38).Contains(p2) {
					return (_1150_v38).Get(p2).(bool)
				}
				return !((_this).F7())
			})())
			var _pat_let_tv44 = p0
			_ = _pat_let_tv44
			var _pat_let_tv45 = p0
			_ = _pat_let_tv45
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1149_v37), 0))
			_ = _index167
			var _rhs170 _dafny.Int = _dafny.IntOfInt64(721)
			_ = _rhs170
			var _rhs171 bool = ((func() bool {
				if p1 {
					return (func() bool {
						if (_1150_v38).Contains((_this).F7()) {
							return (_1150_v38).Get((_this).F7()).(bool)
						}
						return p1
					})()
				}
				return (func(_pat_let34_0 D0) D0 {
					return func(_1152_dt__update__tmp_h0 D0) D0 {
						return func(_pat_let35_0 _dafny.Int) D0 {
							return func(_1153_dt__update_hcf3_h0 _dafny.Int) D0 {
								return func(_pat_let36_0 _dafny.Int) D0 {
									return func(_1154_dt__update_hcf1_h0 _dafny.Int) D0 {
										return func(_pat_let37_0 bool) D0 {
											return func(_1155_dt__update_hcf2_h0 bool) D0 {
												return Companion_D0_.Create_DC1_(_1154_dt__update_hcf1_h0, _1155_dt__update_hcf2_h0, _1153_dt__update_hcf3_h0, (_1152_dt__update__tmp_h0).Dtor_cf4())
											}(_pat_let37_0)
										}((_this).F7())
									}(_pat_let36_0)
								}(_pat_let_tv45)
							}(_pat_let35_0)
						}(_pat_let_tv44)
					}(_pat_let34_0)
				}(_1151_v39)).Dtor_cf2()
			})()) && ((_this).F7())
			_ = _rhs171
			var _lhs131 _dafny.Array = _1149_v37
			_ = _lhs131
			var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1149_v37), 0))
			_ = _lhs132
			var _lhs133 *GlobalState = globalState
			_ = _lhs133
			(_lhs131).ArraySet1(_rhs170, (_lhs132).Int())
			_lhs133.F5 = _rhs171
			var _1156_v40 _dafny.Sequence
			_ = _1156_v40
			_1156_v40 = _dafny.UnicodeSeqOfUtf8Bytes("t")
			var _1157_v41 _dafny.Sequence
			_ = _1157_v41
			_1157_v41 = _dafny.SeqOf((_dafny.SetOf(_dafny.IntOfUint32((_1156_v40).Cardinality()))).Cardinality())
			var _1158_v42 _dafny.Set
			_ = _1158_v42
			_1158_v42 = _dafny.SetOf(p1, false)
			var _1159_v43 _dafny.Map
			_ = _1159_v43
			_1159_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, ((_dafny.MultiSetOf(_dafny.IntOfUint32((_1157_v41).Cardinality()), (_1158_v42).Cardinality())).Update(_dafny.IntOfInt64(261), Companion_Default___.Abs((_1149_v37).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(612), _dafny.ArrayLen((_1149_v37), 0))).Int()).(_dafny.Int)))).Cardinality())
			var _1160_v44 T0
			_ = _1160_v44
			var _nw178 *C4 = New_C4_()
			_ = _nw178
			_nw178.Ctor__((func() _dafny.Int {
				if (_1159_v43).Contains(p2) {
					return (_1159_v43).Get(p2).(_dafny.Int)
				}
				return p0
			})())
			_1160_v44 = _nw178
			var _1161_v45 _dafny.Map
			_ = _1161_v45
			_1161_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F8())
			var _1162_v46 _dafny.Map
			_ = _1162_v46
			_1162_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1160_v44, _1161_v45)
			(globalState).F2 = ((_1162_v46).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1160_v44, _1161_v45)).Merge(_1162_v46))).Cardinality()
		} else {
			var _1163_v47 D13
			_ = _1163_v47
			_1163_v47 = Companion_D13_.Create_DC32_(p0, (_this).F7(), p1, p1, p2)
			var _1164_v48 D13
			_ = _1164_v48
			_1164_v48 = Companion_D13_.Create_DC34_(_1163_v47)
			var _1165_v49 _dafny.Map
			_ = _1165_v49
			_1165_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _1164_v48)
			_1165_v49 = _1165_v49
			var _1166_v50 _dafny.Array
			_ = _1166_v50
			var _len0_32 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_32
			var _nw179 _dafny.Array
			_ = _nw179
			if _len0_32.Cmp(_dafny.Zero) == 0 {
				_nw179 = _dafny.NewArray(_len0_32)
			} else {
				var _init32 func(_dafny.Int) bool = (func(_1167_p1 bool) func(_dafny.Int) bool {
					return func(_1168_i4 _dafny.Int) bool {
						return _1167_p1
					}
				})(p1)
				_ = _init32
				var _element0_32 = _init32(_dafny.Zero)
				_ = _element0_32
				_nw179 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
				(_nw179).ArraySet1(_element0_32, 0)
				var _nativeLen0_32 = (_len0_32).Int()
				_ = _nativeLen0_32
				for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
					(_nw179).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
				}
			}
			_1166_v50 = _nw179
			var _1169_v51 _dafny.Map
			_ = _1169_v51
			_1169_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
			var _1170_v52 _dafny.Sequence
			_ = _1170_v52
			_1170_v52 = _dafny.UnicodeSeqOfUtf8Bytes("gomydnwq")
			var _1171_v53 T1
			_ = _1171_v53
			var _nw180 *C3 = New_C3_()
			_ = _nw180
			_nw180.Ctor__(p0, _dafny.SeqOf((_1169_v51).Cardinality(), p0, _dafny.IntOfUint32((_1170_v52).Cardinality()), p0, p0), _dafny.IntOfUint32((p3).Cardinality()))
			_1171_v53 = _nw180
			var _1172_v54 D11
			_ = _1172_v54
			_1172_v54 = Companion_D11_.Create_DC25_(_1166_v50, _1171_v53, (p0).Times(p0))
			_1172_v54 = _1172_v54
			var _1173_v55 _dafny.Set
			_ = _1173_v55
			_1173_v55 = _dafny.SetOf(p1, false, (_this).F7(), p1, false)
			var _1174_v56 _dafny.Map
			_ = _1174_v56
			_1174_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(351), !((_this).F7()))
			(globalState).F5 = (Companion_D14_.Create_DC36_((_this).F7(), (_1173_v55).Cardinality(), _dafny.IntOfUint32((_1170_v52).Cardinality()), _1174_v56)).Dtor_cf68()
			(globalState).F0 = _dafny.IntOfInt64(909)
			var _1175_v57 _dafny.Sequence
			_ = _1175_v57
			_1175_v57 = _dafny.SeqOf(p0)
			var _1176_v58 _dafny.Array
			_ = _1176_v58
			var _nw181 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw181
			_1176_v58 = _nw181
			var _1177_v59 _dafny.Map
			_ = _1177_v59
			_1177_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1175_v57, (Companion_D3_.Create_DC8_(Companion_D1_.Create_DC4_(_1175_v57), _1176_v58, p0)).Dtor_cf15())
			_1177_v59 = _1177_v59
		}
		var _1178_v60 _dafny.Array
		_ = _1178_v60
		var _nw182 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(4))
		_ = _nw182
		_1178_v60 = _nw182
		for _iter58 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1178_v60), 0))); ; {
			_guard_loop_4, _ok58 := _iter58()
			if !_ok58 {
				break
			}
			var _1179_i5 _dafny.Int
			_1179_i5 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1179_i5).Sign() != -1) && ((_1179_i5).Cmp(_dafny.ArrayLen((_1178_v60), 0)) < 0)) {
				(_1178_v60).ArraySet1(Companion_D6_.Create_DC15_(_dafny.SetOf(p2)), (_1179_i5).Int())
			}
		}
		var _1180_v61 _dafny.Map
		_ = _1180_v61
		_1180_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F8())
		_1180_v61 = (_1180_v61).Update(p0, (_this).F8())
		var _1181_v62 _dafny.Array
		_ = _1181_v62
		var _len0_33 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_33
		var _nw183 _dafny.Array
		_ = _nw183
		if _len0_33.Cmp(_dafny.Zero) == 0 {
			_nw183 = _dafny.NewArray(_len0_33)
		} else {
			var _init33 func(_dafny.Int) _dafny.Sequence = func(_1182_i7 _dafny.Int) _dafny.Sequence {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("xbox"), _dafny.UnicodeSeqOfUtf8Bytes("rvl"))
			}
			_ = _init33
			var _element0_33 = _init33(_dafny.Zero)
			_ = _element0_33
			_nw183 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
			(_nw183).ArraySet1(_element0_33, 0)
			var _nativeLen0_33 = (_len0_33).Int()
			_ = _nativeLen0_33
			for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
				(_nw183).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
			}
		}
		_1181_v62 = _nw183
		for _iter59 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1181_v62), 0))); ; {
			_guard_loop_5, _ok59 := _iter59()
			if !_ok59 {
				break
			}
			var _1183_i6 _dafny.Int
			_1183_i6 = interface{}(_guard_loop_5).(_dafny.Int)
			if (true) && (((_1183_i6).Sign() != -1) && ((_1183_i6).Cmp(_dafny.ArrayLen((_1181_v62), 0)) < 0)) {
				(_1181_v62).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jkqigycsa"), _dafny.UnicodeSeqOfUtf8Bytes("ysjgmd")), (_1183_i6).Int())
			}
		}
		var _1184_v63 D13
		_ = _1184_v63
		_1184_v63 = Companion_D13_.Create_DC32_(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(975), p0), p1, (_this).F7(), p2, p1)
		var _source21 D13 = _1184_v63
		_ = _source21
		if _source21.Is_DC32() {
			var _1185___mcc_h23 _dafny.Int = _source21.Get_().(D13_DC32).Cf56
			_ = _1185___mcc_h23
			var _1186___mcc_h24 bool = _source21.Get_().(D13_DC32).Cf57
			_ = _1186___mcc_h24
			var _1187___mcc_h25 bool = _source21.Get_().(D13_DC32).Cf58
			_ = _1187___mcc_h25
			var _1188___mcc_h26 bool = _source21.Get_().(D13_DC32).Cf59
			_ = _1188___mcc_h26
			var _1189___mcc_h27 bool = _source21.Get_().(D13_DC32).Cf60
			_ = _1189___mcc_h27
			var _1190_cf60 bool = _1189___mcc_h27
			_ = _1190_cf60
			var _1191_cf59 bool = _1188___mcc_h26
			_ = _1191_cf59
			var _1192_cf58 bool = _1187___mcc_h25
			_ = _1192_cf58
			var _1193_cf57 bool = _1186___mcc_h24
			_ = _1193_cf57
			var _1194_cf56 _dafny.Int = _1185___mcc_h23
			_ = _1194_cf56
			var _1195_v64 D8
			_ = _1195_v64
			_1195_v64 = Companion_D8_.Create_DC21_(_1191_cf59, _1192_cf58, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p3, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p3).Cardinality()))).Uint32(), _1192_cf58)).Cardinality()))
			(globalState).F5 = (_1195_v64).Dtor_cf36()
			var _1196_v65 _dafny.Array
			_ = _1196_v65
			var _len0_34 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_34
			var _nw184 _dafny.Array
			_ = _nw184
			if _len0_34.Cmp(_dafny.Zero) == 0 {
				_nw184 = _dafny.NewArray(_len0_34)
			} else {
				var _init34 func(_dafny.Int) _dafny.Int = func(_1197_i8 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_1197_i8, _dafny.IntOfInt64(630))
				}
				_ = _init34
				var _element0_34 = _init34(_dafny.Zero)
				_ = _element0_34
				_nw184 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
				(_nw184).ArraySet1(_element0_34, 0)
				var _nativeLen0_34 = (_len0_34).Int()
				_ = _nativeLen0_34
				for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
					(_nw184).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
				}
			}
			_1196_v65 = _nw184
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1196_v65), 0))
			_ = _index168
			(_1196_v65).ArraySet1(_dafny.IntOfInt64(-844), (_index168).Int())
			var _1198_v67 _dafny.Array
			_ = _1198_v67
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_35
			var _nw185 _dafny.Array
			_ = _nw185
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw185 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Sequence = (func(_1199_p2 bool, _1200_cf56 _dafny.Int) func(_dafny.Int) _dafny.Sequence {
					return func(_1201_i9 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kxvatj")).Cardinality()), (func() _dafny.Set {
							var _coll54 = _dafny.NewBuilder()
							_ = _coll54
							for _iter60 := _dafny.Iterate((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1199_p2), _1200_cf56)).Cardinality())).Elements()); ; {
								_compr_54, _ok60 := _iter60()
								if !_ok60 {
									break
								}
								var _1202_v66 _dafny.Int
								_1202_v66 = interface{}(_compr_54).(_dafny.Int)
								if (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1199_p2), _1200_cf56)).Cardinality())).Contains(_1202_v66) {
									_coll54.Add((_1202_v66).Plus(_dafny.IntOfInt64(372)))
								}
							}
							return _coll54.ToSet()
						}()).Cardinality()), _dafny.SeqOf(_1200_cf56))
					}
				})(p2, _1194_cf56)
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw185 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw185).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw185).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1198_v67 = _nw185
			var _1203_v68 _dafny.Sequence
			_ = _1203_v68
			_1203_v68 = _dafny.UnicodeSeqOfUtf8Bytes("gaq")
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1196_v65), 0))
			_ = _index169
			var _rhs172 _dafny.Int = _1194_cf56
			_ = _rhs172
			var _rhs173 bool = !(true)
			_ = _rhs173
			var _rhs174 _dafny.Array = _1198_v67
			_ = _rhs174
			var _rhs175 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("eakxl"), _dafny.Companion_Sequence_.Concatenate(_1203_v68, _1203_v68))
			_ = _rhs175
			var _rhs176 bool = (_this).F7()
			_ = _rhs176
			var _lhs134 _dafny.Array = _1196_v65
			_ = _lhs134
			var _lhs135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(335), _dafny.ArrayLen((_1196_v65), 0))
			_ = _lhs135
			var _lhs136 *GlobalState = globalState
			_ = _lhs136
			(_lhs134).ArraySet1(_rhs172, (_lhs135).Int())
			r0 = _rhs173
			_1198_v67 = _rhs174
			_1203_v68 = _rhs175
			_lhs136.F5 = _rhs176
			var _1204_v69 _dafny.Map
			_ = _1204_v69
			_1204_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1196_v65, true)
			var _1205_v70 _dafny.Int
			_ = _1205_v70
			var _1206_v71 _dafny.Int
			_ = _1206_v71
			var _1207_v72 bool
			_ = _1207_v72
			var _1208_v73 _dafny.Int
			_ = _1208_v73
			var _out8 _dafny.Int
			_ = _out8
			var _out9 _dafny.Int
			_ = _out9
			var _out10 bool
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = (_this).M1(!(Companion_Default___.Fm1(_1062_v2, (p3).Select((Companion_Default___.SafeIndex(_1194_cf56, _dafny.IntOfUint32((p3).Cardinality()))).Uint32()).(bool), (_dafny.Zero).Minus((_1204_v69).Cardinality()), false, globalState)), globalState)
			_1205_v70 = _out8
			_1206_v71 = _out9
			_1207_v72 = _out10
			_1208_v73 = _out11
			_1196_v65 = _1196_v65
		} else if _source21.Is_DC33() {
			var _1209___mcc_h28 bool = _source21.Get_().(D13_DC33).Cf61
			_ = _1209___mcc_h28
			var _1210___mcc_h29 _dafny.Int = _source21.Get_().(D13_DC33).Cf62
			_ = _1210___mcc_h29
			var _1211___mcc_h30 _dafny.Int = _source21.Get_().(D13_DC33).Cf63
			_ = _1211___mcc_h30
			var _1212___mcc_h31 bool = _source21.Get_().(D13_DC33).Cf64
			_ = _1212___mcc_h31
			var _1213___mcc_h32 *C3 = _source21.Get_().(D13_DC33).Cf65
			_ = _1213___mcc_h32
			var _1214_cf65 *C3 = _1213___mcc_h32
			_ = _1214_cf65
			var _1215_cf64 bool = _1212___mcc_h31
			_ = _1215_cf64
			var _1216_cf63 _dafny.Int = _1211___mcc_h30
			_ = _1216_cf63
			var _1217_cf62 _dafny.Int = _1210___mcc_h29
			_ = _1217_cf62
			var _1218_cf61 bool = _1209___mcc_h28
			_ = _1218_cf61
			var _1219_v74 _dafny.CodePoint
			_ = _1219_v74
			_1219_v74 = _dafny.CodePoint('u')
			_1219_v74 = _1219_v74
			_1216_cf63 = Companion_Default___.SafeDivisionInt(_1217_cf62, _dafny.IntOfInt64(566))
			var _1220_v75 _dafny.Array
			_ = _1220_v75
			var _nw186 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
			_ = _nw186
			_1220_v75 = _nw186
			var _1221_v76 _dafny.Sequence
			_ = _1221_v76
			_1221_v76 = _dafny.UnicodeSeqOfUtf8Bytes("n")
			var _1222_v77 *C2
			_ = _1222_v77
			var _nw187 *C2 = New_C2_()
			_ = _nw187
			_nw187.Ctor__(_1221_v76, _1218_cf61)
			_1222_v77 = _nw187
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_1220_v75), 0))
			_ = _index170
			(_1220_v75).ArraySet1(_1222_v77, (_index170).Int())
			var _1223_v78 _dafny.MultiSet
			_ = _1223_v78
			_1223_v78 = _dafny.MultiSetOf((_this).F7())
			var _1224_v79 _dafny.Sequence
			_ = _1224_v79
			_1224_v79 = _dafny.SeqOf(_1222_v77, _1222_v77)
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_1220_v75), 0))
			_ = _index171
			var _rhs177 _dafny.CodePoint = (_this).F8()
			_ = _rhs177
			var _rhs178 *C2 = _1222_v77
			_ = _rhs178
			var _rhs179 *C2 = (_1224_v79).Select((Companion_Default___.SafeIndex(((_1214_cf65).F9()).Plus((_dafny.SetOf((_1214_cf65).F9(), p0)).Cardinality()), _dafny.IntOfUint32((_1224_v79).Cardinality()))).Uint32()).(*C2)
			_ = _rhs179
			var _rhs180 _dafny.MultiSet = (_1223_v78).Intersection(_1223_v78)
			_ = _rhs180
			var _lhs137 _dafny.Array = _1220_v75
			_ = _lhs137
			var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(356), _dafny.ArrayLen((_1220_v75), 0))
			_ = _lhs138
			_1219_v74 = _rhs177
			(_lhs137).ArraySet1(_rhs178, (_lhs138).Int())
			_1222_v77 = _rhs179
			_1223_v78 = _rhs180
			if !(_1215_cf64) {
				var _1225_v80 *C4
				_ = _1225_v80
				var _nw188 *C4 = New_C4_()
				_ = _nw188
				_nw188.Ctor__((_1214_cf65).F9())
				_1225_v80 = _nw188
				var _1226_v81 _dafny.Set
				_ = _1226_v81
				_1226_v81 = _dafny.SetOf(_1225_v80, _1225_v80)
				(globalState).F5 = (_1226_v81).IsSubsetOf(_1226_v81)
				var _1227_v82 _dafny.Map
				_ = _1227_v82
				_1227_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1222_v77.F12, (_1222_v77).F11())
				var _1228_v83 _dafny.MultiSet
				_ = _1228_v83
				_1228_v83 = _dafny.MultiSetOf(_1214_cf65)
				var _1229_v84 _dafny.MultiSet
				_ = _1229_v84
				_1229_v84 = _dafny.MultiSetOf((_1228_v83).Cardinality())
				var _1230_v85 _dafny.Set
				_ = _1230_v85
				_1230_v85 = _dafny.SetOf((_1229_v84).Cardinality(), (_this).Fm6(_1216_cf63, _1222_v77.F12, _1216_cf63, globalState), _1216_cf63)
				var _1231_v86 _dafny.Array
				_ = _1231_v86
				var _nwElement0_37 _dafny.Int = p0
				_ = _nwElement0_37
				var _nw189 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(13))
				_ = _nw189
				(_nw189).ArraySet1(_nwElement0_37, 0)
				(_nw189).ArraySet1(_1216_cf63, 1)
				(_nw189).ArraySet1(p0, 2)
				(_nw189).ArraySet1((_1214_cf65).F9(), 3)
				(_nw189).ArraySet1(_dafny.IntOfUint32((p3).Cardinality()), 4)
				(_nw189).ArraySet1(((_this).Fm5(globalState)).Minus(_1216_cf63), 5)
				(_nw189).ArraySet1(_dafny.IntOfInt64(588), 6)
				(_nw189).ArraySet1(_dafny.IntOfInt64(882), 7)
				(_nw189).ArraySet1(p0, 8)
				(_nw189).ArraySet1((_1216_cf63).Plus(_1217_cf62), 9)
				(_nw189).ArraySet1((_1214_cf65).F9(), 10)
				(_nw189).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_1227_v82).Contains(_1222_v77.F12) {
						return (_1227_v82).Get(_1222_v77.F12).(_dafny.Sequence)
					}
					return (func() _dafny.Sequence {
						if (_1227_v82).Contains(p2) {
							return (_1227_v82).Get(p2).(_dafny.Sequence)
						}
						return (_1222_v77).F11()
					})()
				})()).Cardinality()), _1217_cf62), 11)
				(_nw189).ArraySet1((_1214_cf65).Fm17(((_1214_cf65).F10()).Select((Companion_Default___.SafeIndex(_1216_cf63, _dafny.IntOfUint32(((_1214_cf65).F10()).Cardinality()))).Uint32()).(_dafny.Int), _1230_v85, globalState), 12)
				_1231_v86 = _nw189
				var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_1231_v86), 0))
				_ = _index172
				(_1231_v86).ArraySet1((_1214_cf65).F9(), (_index172).Int())
				var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_1231_v86), 0))
				_ = _index173
				(_1231_v86).ArraySet1(_dafny.IntOfUint32((_1221_v76).Cardinality()), (_index173).Int())
				var _1232_v87 _dafny.Map
				_ = _1232_v87
				_1232_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_1217_cf62), (_1222_v77).F11())
				_1221_v76 = _dafny.Companion_Sequence_.Concatenate((_1222_v77).F11(), (func() _dafny.Sequence {
					if (_1232_v87).Contains((_dafny.Zero).Minus((_1214_cf65).F9())) {
						return (_1232_v87).Get((_dafny.Zero).Minus((_1214_cf65).F9())).(_dafny.Sequence)
					}
					return (_1222_v77).F11()
				})())
				var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_1231_v86), 0))
				_ = _index174
				(_1231_v86).ArraySet1((p0).Times((_1231_v86).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(300), _dafny.ArrayLen((_1231_v86), 0))).Int()).(_dafny.Int)), (_index174).Int())
				(globalState).F5 = _1218_cf61
			} else {
				var _1233_v88 _dafny.Array
				_ = _1233_v88
				var _len0_36 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_36
				var _nw190 _dafny.Array
				_ = _nw190
				if _len0_36.Cmp(_dafny.Zero) == 0 {
					_nw190 = _dafny.NewArray(_len0_36)
				} else {
					var _init36 func(_dafny.Int) _dafny.Int = (func(_1234_cf65 *C3) func(_dafny.Int) _dafny.Int {
						return func(_1235_i10 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_1235_i10, (_1234_cf65).F9())
						}
					})(_1214_cf65)
					_ = _init36
					var _element0_36 = _init36(_dafny.Zero)
					_ = _element0_36
					_nw190 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
					(_nw190).ArraySet1(_element0_36, 0)
					var _nativeLen0_36 = (_len0_36).Int()
					_ = _nativeLen0_36
					for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
						(_nw190).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
					}
				}
				_1233_v88 = _nw190
				_1233_v88 = _1233_v88
				(globalState).F5 = _1222_v77.F12
				var _1236_v89 _dafny.Set
				_ = _1236_v89
				_1236_v89 = _dafny.SetOf((p0).Plus((_1214_cf65).F9()), p0, Companion_Default___.Fm15(_dafny.IntOfInt64(-984), _1218_cf61, globalState))
				_1236_v89 = _1236_v89
				var _1237_v90 *C1
				_ = _1237_v90
				var _nw191 *C1 = New_C1_()
				_ = _nw191
				_nw191.Ctor__()
				_1237_v90 = _nw191
				var _1238_v91 *C3
				_ = _1238_v91
				var _nw192 *C3 = New_C3_()
				_ = _nw192
				_nw192.Ctor__(_1216_cf63, (_1214_cf65).F10(), (_1214_cf65).F9())
				_1238_v91 = _nw192
			}
		} else if _source21.Is_DC31() {
			var _1239___mcc_h33 _dafny.Map = _source21.Get_().(D13_DC31).Cf55
			_ = _1239___mcc_h33
			var _1240_cf55 _dafny.Map = _1239___mcc_h33
			_ = _1240_cf55
			var _1241_v92 _dafny.Sequence
			_ = _1241_v92
			_1241_v92 = _dafny.SeqOf(Companion_Default___.Fm36((_this).F7(), (_this).F7(), (_this).F7(), _dafny.CodePoint('x'), globalState))
			(globalState).F1 = !(_dafny.Companion_Sequence_.Equal(_1241_v92, _1241_v92))
			var _1242_v93 _dafny.Array
			_ = _1242_v93
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(17)
			_ = _len0_37
			var _nw193 _dafny.Array
			_ = _nw193
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw193 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Int = func(_1243_i12 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_1243_i12, _dafny.IntOfInt64(-109))
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw193 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw193).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw193).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1242_v93 = _nw193
			var _1244_v94 _dafny.Map
			_ = _1244_v94
			_1244_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(910))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}((func(_1245_p1 bool) func(_dafny.Int) _dafny.Int {
				return func(_1246_i11 _dafny.Int) _dafny.Int {
					return (_dafny.SetOf(_1245_p1, (_this).F7())).Cardinality()
				}
			})(p1))), _1242_v93)
			var _1247_v95 D8
			_ = _1247_v95
			_1247_v95 = Companion_D8_.Create_DC20_((_1244_v94).Merge(_1244_v94))
			var _source22 D8 = _1247_v95
			_ = _source22
			if _source22.Is_DC21() {
				var _1248___mcc_h35 bool = _source22.Get_().(D8_DC21).Cf36
				_ = _1248___mcc_h35
				var _1249___mcc_h36 bool = _source22.Get_().(D8_DC21).Cf37
				_ = _1249___mcc_h36
				var _1250___mcc_h37 _dafny.Int = _source22.Get_().(D8_DC21).Cf38
				_ = _1250___mcc_h37
				var _1251_cf38 _dafny.Int = _1250___mcc_h37
				_ = _1251_cf38
				var _1252_cf37 bool = _1249___mcc_h36
				_ = _1252_cf37
				var _1253_cf36 bool = _1248___mcc_h35
				_ = _1253_cf36
				var _1254_v96 _dafny.MultiSet
				_ = _1254_v96
				_1254_v96 = _dafny.MultiSetOf(_1253_cf36)
				var _1255_v97 _dafny.Sequence
				_ = _1255_v97
				_1255_v97 = _dafny.UnicodeSeqOfUtf8Bytes("ngnbjrkg")
				var _1256_v98 _dafny.Map
				_ = _1256_v98
				_1256_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1255_v97, _1251_cf38)
				var _1257_v99 _dafny.Map
				_ = _1257_v99
				_1257_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1253_cf36, _1254_v96)
				var _1258_v100 _dafny.Map
				_ = _1258_v100
				_1258_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, p1)
				var _1259_v101 _dafny.Sequence
				_ = _1259_v101
				_1259_v101 = _dafny.SeqOf(_1254_v96)
				var _1260_v102 _dafny.Array
				_ = _1260_v102
				var _nwElement0_38 _dafny.MultiSet = _1254_v96
				_ = _nwElement0_38
				var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(26))
				_ = _nw194
				(_nw194).ArraySet1(_nwElement0_38, 0)
				(_nw194).ArraySet1(_1254_v96, 1)
				(_nw194).ArraySet1((_dafny.MultiSetFromSeq(p3)).Union(_1254_v96), 2)
				(_nw194).ArraySet1(_1254_v96, 3)
				(_nw194).ArraySet1(_1254_v96, 4)
				(_nw194).ArraySet1((_1254_v96).Union(_1254_v96), 5)
				(_nw194).ArraySet1(_dafny.MultiSetOf(_1253_cf36, _1252_cf37, _1252_cf37), 6)
				(_nw194).ArraySet1(_1254_v96, 7)
				(_nw194).ArraySet1(_1254_v96, 8)
				(_nw194).ArraySet1(Companion_Default___.Fm30(_dafny.UnicodeSeqOfUtf8Bytes("ae"), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1255_v97, _1251_cf38)).Update(_dafny.UnicodeSeqOfUtf8Bytes("apq"), p0), (_this).F7(), globalState), 9)
				(_nw194).ArraySet1(Companion_Default___.Fm30(_dafny.Companion_Sequence_.Update(_1255_v97, (Companion_Default___.SafeIndex(_1251_cf38, _dafny.IntOfUint32((_1255_v97).Cardinality()))).Uint32(), _dafny.CodePoint('k')), _1256_v98, p2, globalState), 10)
				(_nw194).ArraySet1(_1254_v96, 11)
				(_nw194).ArraySet1(_1254_v96, 12)
				(_nw194).ArraySet1(_1254_v96, 13)
				(_nw194).ArraySet1(_1254_v96, 14)
				(_nw194).ArraySet1((_1254_v96).Union((func() _dafny.MultiSet {
					if (_1257_v99).Contains((func() bool {
						if (_1258_v100).Contains((_this).F7()) {
							return (_1258_v100).Get((_this).F7()).(bool)
						}
						return false
					})()) {
						return (_1257_v99).Get((func() bool {
							if (_1258_v100).Contains((_this).F7()) {
								return (_1258_v100).Get((_this).F7()).(bool)
							}
							return false
						})()).(_dafny.MultiSet)
					}
					return _1254_v96
				})()), 15)
				(_nw194).ArraySet1((_1254_v96).Intersection(_dafny.MultiSetFromSeq(p3)), 16)
				(_nw194).ArraySet1(_1254_v96, 17)
				(_nw194).ArraySet1(_1254_v96, 18)
				(_nw194).ArraySet1(_dafny.MultiSetFromSeq(p3), 19)
				(_nw194).ArraySet1(_1254_v96, 20)
				(_nw194).ArraySet1(_1254_v96, 21)
				(_nw194).ArraySet1(_1254_v96, 22)
				(_nw194).ArraySet1(_1254_v96, 23)
				(_nw194).ArraySet1((_1259_v101).Select((Companion_Default___.SafeIndex(_1251_cf38, _dafny.IntOfUint32((_1259_v101).Cardinality()))).Uint32()).(_dafny.MultiSet), 24)
				(_nw194).ArraySet1(_1254_v96, 25)
				_1260_v102 = _nw194
				var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1260_v102), 0))
				_ = _index175
				(_1260_v102).ArraySet1(_1254_v96, (_index175).Int())
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(165), _dafny.ArrayLen((_1260_v102), 0))
				_ = _index176
				(_1260_v102).ArraySet1(_1254_v96, (_index176).Int())
				var _1261_v103 *C2
				_ = _1261_v103
				var _nw195 *C2 = New_C2_()
				_ = _nw195
				_nw195.Ctor__(_dafny.Companion_Sequence_.Update(_1255_v97, (Companion_Default___.SafeIndex(_1251_cf38, _dafny.IntOfUint32((_1255_v97).Cardinality()))).Uint32(), (_this).F8()), _1253_cf36)
				_1261_v103 = _nw195
				var _1262_v104 _dafny.Map
				_ = _1262_v104
				_1262_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1261_v103)
				(globalState).F0 = (_1262_v104).Cardinality()
				var _1263_v105 *C1
				_ = _1263_v105
				var _nw196 *C1 = New_C1_()
				_ = _nw196
				_nw196.Ctor__()
				_1263_v105 = _nw196
				var _1264_v106 _dafny.Sequence
				_ = _1264_v106
				_1264_v106 = _dafny.SeqOf((func() *C1 {
					if p2 {
						return _1263_v105
					}
					return _1263_v105
				})())
				_1263_v105 = (_1264_v106).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(909), _dafny.IntOfUint32((_1264_v106).Cardinality()))).Uint32()).(*C1)
				var _1265_v107 _dafny.Array
				_ = _1265_v107
				var _len0_38 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_38
				var _nw197 _dafny.Array
				_ = _nw197
				if _len0_38.Cmp(_dafny.Zero) == 0 {
					_nw197 = _dafny.NewArray(_len0_38)
				} else {
					var _init38 func(_dafny.Int) bool = (func(_1266_v103 *C2) func(_dafny.Int) bool {
						return func(_1267_i13 _dafny.Int) bool {
							return _1266_v103.F12
						}
					})(_1261_v103)
					_ = _init38
					var _element0_38 = _init38(_dafny.Zero)
					_ = _element0_38
					_nw197 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
					(_nw197).ArraySet1(_element0_38, 0)
					var _nativeLen0_38 = (_len0_38).Int()
					_ = _nativeLen0_38
					for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
						(_nw197).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
					}
				}
				_1265_v107 = _nw197
				_1265_v107 = _1265_v107
			} else {
				var _1268___mcc_h38 _dafny.Map = _source22.Get_().(D8_DC20).Cf35
				_ = _1268___mcc_h38
				var _1269_cf35 _dafny.Map = _1268___mcc_h38
				_ = _1269_cf35
				var _1270_v108 D1
				_ = _1270_v108
				_1270_v108 = Companion_D1_.Create_DC4_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(986))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1271_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1272_i14 _dafny.Int) _dafny.Int {
						return _1271_p0
					}
				})(p0))))
				var _1273_v109 *C3
				_ = _1273_v109
				var _nw198 *C3 = New_C3_()
				_ = _nw198
				_nw198.Ctor__(p0, (_1270_v108).Dtor_cf11(), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(412), (_dafny.Zero).Minus(_dafny.IntOfInt64(-262))))
				_1273_v109 = _nw198
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_1242_v93), 0))
				_ = _index177
				(_1242_v93).ArraySet1((_1273_v109).F9(), (_index177).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_1242_v93), 0))
				_ = _index178
				(_1242_v93).ArraySet1((_dafny.SetOf(_dafny.IntOfInt64(366), (_1273_v109).F9(), ((_1273_v109).F10()).Select((Companion_Default___.SafeIndex((_1273_v109).F9(), _dafny.IntOfUint32(((_1273_v109).F10()).Cardinality()))).Uint32()).(_dafny.Int), p0, (_1273_v109).F9())).Cardinality(), (_index178).Int())
				var _1274_v110 *C3
				_ = _1274_v110
				var _nw199 *C3 = New_C3_()
				_ = _nw199
				_nw199.Ctor__((_1273_v109).F9(), _dafny.SeqOf((_1273_v109).F9()), (_1242_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_1242_v93), 0))).Int()).(_dafny.Int))
				_1274_v110 = _nw199
				var _1275_v111 _dafny.Map
				_ = _1275_v111
				_1275_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1242_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(179), _dafny.ArrayLen((_1242_v93), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((p3).Cardinality()))
				_1275_v111 = _1275_v111
			}
			_1242_v93 = (func() _dafny.Array {
				if p1 {
					return _1242_v93
				}
				return _1242_v93
			})()
			(globalState).F5 = p2
		} else {
			var _1276___mcc_h34 D13 = _source21.Get_().(D13_DC34).Cf66
			_ = _1276___mcc_h34
			var _1277_cf66 D13 = _1276___mcc_h34
			_ = _1277_cf66
			var _1278_v112 _dafny.Sequence
			_ = _1278_v112
			_1278_v112 = _dafny.SeqOf(_dafny.IntOfInt64(77))
			var _1279_v113 D1
			_ = _1279_v113
			_1279_v113 = Companion_D1_.Create_DC4_(_1278_v112)
			var _1280_v114 _dafny.Array
			_ = _1280_v114
			var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(27))
			_ = _nw200
			_1280_v114 = _nw200
			var _1281_v115 D3
			_ = _1281_v115
			_1281_v115 = Companion_D3_.Create_DC8_(_1279_v113, _1280_v114, p0)
			var _1282_v116 _dafny.MultiSet
			_ = _1282_v116
			_1282_v116 = _dafny.MultiSetOf(p0, _dafny.IntOfUint32((_dafny.SeqOf(p2, !(true), p2)).Cardinality()))
			var _1283_v117 *C3
			_ = _1283_v117
			var _nw201 *C3 = New_C3_()
			_ = _nw201
			_nw201.Ctor__((_1281_v115).Dtor_cf16(), _dafny.SeqOf(p0, (_dafny.Zero).Minus((_1282_v116).Cardinality()), p0), p0)
			_1283_v117 = _nw201
			(globalState).F5 = ((_dafny.Zero).Minus((_1283_v117).F9())).Cmp(p0) != 0
			var _1284_v118 _dafny.Sequence
			_ = _1284_v118
			_1284_v118 = _dafny.UnicodeSeqOfUtf8Bytes("lhbgnik")
			var _rhs181 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1278_v112, (_1283_v117).F10()), _dafny.SeqOf((_1283_v117).F9()))
			_ = _rhs181
			var _rhs182 _dafny.Int = p0
			_ = _rhs182
			var _rhs183 bool = true
			_ = _rhs183
			var _rhs184 bool = ((_1283_v117).F9()).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg53 _dafny.Int) interface{} {
					return coer53(arg53)
				}
			}(func(_1285_i15 _dafny.Int) _dafny.CodePoint {
				return (_this).F8()
			}))).Cardinality())) != 0
			_ = _rhs184
			var _rhs185 bool = !_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm2(_dafny.UnicodeSeqOfUtf8Bytes("f"), globalState), _dafny.Companion_Sequence_.Concatenate(_1284_v118, _1284_v118))
			_ = _rhs185
			var _lhs139 *GlobalState = globalState
			_ = _lhs139
			var _lhs140 *GlobalState = globalState
			_ = _lhs140
			var _lhs141 *GlobalState = globalState
			_ = _lhs141
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			_1278_v112 = _rhs181
			_lhs139.F2 = _rhs182
			_lhs140.F1 = _rhs183
			_lhs141.F5 = _rhs184
			_lhs142.F5 = _rhs185
			(globalState).F0 = (((_dafny.Zero).Minus(_dafny.IntOfUint32((_1284_v118).Cardinality()))).Plus((_1283_v117).F9())).Plus(((_1282_v116).Cardinality()).Plus(_dafny.IntOfInt64(108)))
		}
		r0 = p2
		return r0
	}
}
func (_this *C5) F7() bool {
	{
		return _this._f7
	}
}
func (_this *C5) F8() _dafny.CodePoint {
	{
		return _this._f8
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	_f6 _dafny.Int
}

func New_C6_() *C6 {
	_this := C6{}

	_this._f6 = _dafny.Zero
	return &_this
}

type CompanionStruct_C6_ struct {
}

var Companion_C6_ = CompanionStruct_C6_{}

func (_this *C6) Equals(other *C6) bool {
	return _this == other
}

func (_this *C6) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*C6)
	return ok && _this.Equals(other)
}

func (*C6) String() string {
	return "_module.C6"
}

func Type_C6_() _dafny.TypeDescriptor {
	return type_C6_{}
}

type type_C6_ struct {
}

func (_this type_C6_) Default() interface{} {
	return (*C6)(nil)
}

func (_this type_C6_) String() string {
	return "main.C6"
}
func (_this *C6) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C6{}
var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) F6() _dafny.Int {
	return _this._f6
}
func (_this *C6) Ctor__(f6 _dafny.Int) {
	{
		(_this)._f6 = f6
	}
}
func (_this *C6) Fm3(p0 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt((_this).F6(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))).Cardinality())), ((_this).F6()).Times(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(166))).Cardinality())))
	}
}
func (_this *C6) M0(globalState *GlobalState) (_dafny.Int, bool, bool) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 bool = false
		_ = r2
		var _1286_v0 _dafny.CodePoint
		_ = _1286_v0
		_1286_v0 = _dafny.CodePoint('a')
		var _1287_v1 _dafny.Sequence
		_ = _1287_v1
		_1287_v1 = _dafny.SeqOf(_1286_v0, _1286_v0)
		var _1288_v2 _dafny.Map
		_ = _1288_v2
		_1288_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6(), _dafny.IntOfUint32((_1287_v1).Cardinality()))).Cardinality())))
		var _1289_v3 bool
		_ = _1289_v3
		_1289_v3 = true
		var _1290_v4 _dafny.Sequence
		_ = _1290_v4
		_1290_v4 = _dafny.SeqOf(!(_1288_v2).Equals(_1288_v2), _1289_v3, _1289_v3, _1289_v3)
		_1290_v4 = _dafny.Companion_Sequence_.Update(_1290_v4, (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1290_v4).Cardinality()))).Uint32(), !((_1290_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_1289_v3)).Cardinality()), _dafny.IntOfUint32((_1290_v4).Cardinality()))).Uint32()).(bool)))
		var _hi8 _dafny.Int = (_this).F6()
		_ = _hi8
		for _1291_i0 := (_this).F6(); _1291_i0.Cmp(_hi8) < 0; _1291_i0 = _1291_i0.Plus(_dafny.One) {
			var _1292_v5 D0
			_ = _1292_v5
			_1292_v5 = Companion_D0_.Create_DC2_(_1291_i0, _1286_v0, _dafny.IntOfInt64(618), (_this).F6(), _1291_i0)
			var _1293_v6 _dafny.Set
			_ = _1293_v6
			_1293_v6 = _dafny.SetOf(_1289_v3, _1289_v3)
			var _1294_v7 _dafny.Array
			_ = _1294_v7
			var _nwElement0_39 _dafny.CodePoint = _1286_v0
			_ = _nwElement0_39
			var _nw202 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(14))
			_ = _nw202
			(_nw202).ArraySet1CodePoint(_nwElement0_39, 0)
			(_nw202).ArraySet1CodePoint(_dafny.CodePoint('o'), 1)
			(_nw202).ArraySet1CodePoint(_1286_v0, 2)
			(_nw202).ArraySet1CodePoint(_1286_v0, 3)
			(_nw202).ArraySet1CodePoint(_1286_v0, 4)
			(_nw202).ArraySet1CodePoint(_1286_v0, 5)
			(_nw202).ArraySet1CodePoint(_1286_v0, 6)
			(_nw202).ArraySet1CodePoint(_1286_v0, 7)
			(_nw202).ArraySet1CodePoint(_1286_v0, 8)
			(_nw202).ArraySet1CodePoint(_1286_v0, 9)
			(_nw202).ArraySet1CodePoint((func() _dafny.CodePoint {
				if _1289_v3 {
					return _1286_v0
				}
				return _1286_v0
			})(), 10)
			(_nw202).ArraySet1CodePoint(_1286_v0, 11)
			(_nw202).ArraySet1CodePoint(_1286_v0, 12)
			(_nw202).ArraySet1CodePoint(Companion_Default___.Fm4(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf((_dafny.Zero).Minus(_1291_i0)), _1292_v5), (_dafny.Zero).Minus(_1291_i0), (_1293_v6).Cardinality(), globalState), 13)
			_1294_v7 = _nw202
			var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1294_v7), 0))
			_ = _index179
			(_1294_v7).ArraySet1CodePoint(_1286_v0, (_index179).Int())
			var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1294_v7), 0))
			_ = _index180
			(_1294_v7).ArraySet1CodePoint(_1286_v0, (_index180).Int())
			(globalState).F5 = false
			if true {
				(globalState).F0 = (func() _dafny.Int {
					if _1289_v3 {
						return (_this).F6()
					}
					return (_this).F6()
				})()
				var _1295_v8 _dafny.Set
				_ = _1295_v8
				_1295_v8 = _dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(664), _1291_i0, _1291_i0, (_this).F6()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(144))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1296_i0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1297_i1 _dafny.Int) _dafny.Int {
						return _1296_i0
					}
				})(_1291_i0))))
				var _1298_v9 _dafny.Map
				_ = _1298_v9
				_1298_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1295_v8).IsSubsetOf(_1295_v8)), _1293_v6)
				_1298_v9 = (_1298_v9).Merge(_1298_v9)
				var _1299_v10 _dafny.Map
				_ = _1299_v10
				_1299_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1291_i0, _1289_v3)
				var _1300_v11 _dafny.Map
				_ = _1300_v11
				_1300_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(187)).Plus((_1299_v10).Cardinality()), _1289_v3)
				_1300_v11 = (_1300_v11).Update(Companion_Default___.SafeDivisionInt(_1291_i0, (_this).F6()), !(_1289_v3))
				_1299_v10 = (_1299_v10).Update((_1292_v5).Dtor_cf9(), ((_this).F6()).Cmp((_this).F6()) == 0)
				var _1301_v12 _dafny.Array
				_ = _1301_v12
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw203
				_1301_v12 = _nw203
				_1301_v12 = _1301_v12
			} else {
				var _1302_v13 *C1
				_ = _1302_v13
				var _nw204 *C1 = New_C1_()
				_ = _nw204
				_nw204.Ctor__()
				_1302_v13 = _nw204
				var _1303_v14 *C0
				_ = _1303_v14
				var _nw205 *C0 = New_C0_()
				_ = _nw205
				_nw205.Ctor__(true)
				_1303_v14 = _nw205
				var _1304_v15 _dafny.Map
				_ = _1304_v15
				_1304_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _1303_v14)
				var _1305_v16 _dafny.Map
				_ = _1305_v16
				_1305_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1304_v15, (_dafny.Zero).Minus(_1291_i0))
				var _1306_v17 _dafny.Map
				_ = _1306_v17
				_1306_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1305_v16).Cardinality(), (_1303_v14).F13())
				_1306_v17 = (_1306_v17).Update((_this).F6(), (_1303_v14).F13())
				(globalState).F2 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_1289_v3), _dafny.CodePoint('y'))).Update((_1303_v14).F13(), (_1294_v7).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(387), _dafny.ArrayLen((_1294_v7), 0))).Int()))).Cardinality()
				var _1307_v18 _dafny.Set
				_ = _1307_v18
				_1307_v18 = _dafny.SetOf(_1291_i0)
				var _1308_v19 _dafny.Sequence
				_ = _1308_v19
				_1308_v19 = _dafny.SeqOf(_1291_i0, (_1307_v18).Cardinality())
				var _1309_v20 *C3
				_ = _1309_v20
				var _nw206 *C3 = New_C3_()
				_ = _nw206
				_nw206.Ctor__(_1291_i0, _1308_v19, (_this).F6())
				_1309_v20 = _nw206
				var _rhs186 bool = (Companion_D13_.Create_DC33_((_1303_v14).F13(), (_dafny.Zero).Minus((_1288_v2).Cardinality()), _1291_i0, (_1303_v14).F13(), _1309_v20)).Dtor_cf64()
				_ = _rhs186
				var _rhs187 _dafny.Int = (_this).F6()
				_ = _rhs187
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 *GlobalState = globalState
				_ = _lhs144
				_lhs143.F5 = _rhs186
				_lhs144.F0 = _rhs187
				(globalState).F1 = !((_1303_v14).F13())
			}
			_1287_v1 = _1287_v1
		}
		var _1310_v21 _dafny.Map
		_ = _1310_v21
		_1310_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(707))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg55 _dafny.Int) interface{} {
				return coer55(arg55)
			}
		}((func(_1311_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1312_i2 _dafny.Int) _dafny.CodePoint {
				return _1311_v0
			}
		})(_1286_v0))))
		var _1313_v22 _dafny.MultiSet
		_ = _1313_v22
		_1313_v22 = _dafny.MultiSetOf((func() _dafny.Sequence {
			if (_1310_v21).Contains(_1289_v3) {
				return (_1310_v21).Get(_1289_v3).(_dafny.Sequence)
			}
			return _1287_v1
		})(), _dafny.Companion_Sequence_.Concatenate(_1287_v1, _dafny.UnicodeSeqOfUtf8Bytes("gfyelbune")))
		_1313_v22 = ((_1313_v22).Difference(_dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(471))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg56 _dafny.Int) interface{} {
				return coer56(arg56)
			}
		}((func(_1314_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_1315_i3 _dafny.Int) _dafny.CodePoint {
				return _1314_v0
			}
		})(_1286_v0)))))).Update(_1287_v1, Companion_Default___.Abs((_this).F6()))
		var _1316_v23 _dafny.MultiSet
		_ = _1316_v23
		_1316_v23 = _dafny.MultiSetOf(_1289_v3, _1289_v3, !(_1289_v3), _1289_v3)
		if (_1316_v23).IsDisjointFrom((_1316_v23).Union(_1316_v23)) {
			var _1317_v24 _dafny.Map
			_ = _1317_v24
			_1317_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _1289_v3)
			var _1318_v25 _dafny.Sequence
			_ = _1318_v25
			_1318_v25 = _dafny.SeqOf((_this).F6(), (_1317_v24).Cardinality())
			var _1319_v26 _dafny.Set
			_ = _1319_v26
			_1319_v26 = _dafny.SetOf(_dafny.IntOfUint32((_1318_v25).Cardinality()), (_this).F6(), (_this).F6())
			var _1320_v27 *C1
			_ = _1320_v27
			var _nw207 *C1 = New_C1_()
			_ = _nw207
			_nw207.Ctor__()
			_1320_v27 = _nw207
			var _1321_v28 _dafny.Map
			_ = _1321_v28
			_1321_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_1318_v25).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1318_v25).Cardinality()))).Uint32()).(_dafny.Int))).IsProperSubsetOf(_1319_v26), _1320_v27)
			_1321_v28 = (_1321_v28).Update(_1289_v3, _1320_v27)
			var _1322_v29 _dafny.Map
			_ = _1322_v29
			_1322_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, (_this).F6())
			var _1323_v30 _dafny.Map
			_ = _1323_v30
			_1323_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1322_v29, _1289_v3)
			var _rhs188 _dafny.Int = ((_this).F6()).Times((_1323_v30).Cardinality())
			_ = _rhs188
			var _rhs189 _dafny.Int = (_this).F6()
			_ = _rhs189
			var _rhs190 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("npsdnb"), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("npsdnb")).Cardinality()))).Uint32(), _1286_v0)
			_ = _rhs190
			var _lhs145 *GlobalState = globalState
			_ = _lhs145
			var _lhs146 *GlobalState = globalState
			_ = _lhs146
			_lhs145.F0 = _rhs188
			_lhs146.F0 = _rhs189
			_1287_v1 = _rhs190
			var _1324_v32 D0
			_ = _1324_v32
			_1324_v32 = Companion_D0_.Create_DC2_((_this).F6(), _1286_v0, (_this).F6(), _dafny.IntOfInt64(101), (func() _dafny.Int {
				if (_1288_v2).Contains(_dafny.IntOfInt64(-905)) {
					return (_1288_v2).Get(_dafny.IntOfInt64(-905)).(_dafny.Int)
				}
				return _dafny.IntOfInt64(-20)
			})())
			var _1325_v33 _dafny.Set
			_ = _1325_v33
			_1325_v33 = _dafny.SetOf(_1286_v0, (_1324_v32).Dtor_cf6())
			(globalState).F5 = !((func() _dafny.Map {
				var _coll55 = _dafny.NewMapBuilder()
				_ = _coll55
				for _iter61 := _dafny.Iterate((_1325_v33).Elements()); ; {
					_compr_55, _ok61 := _iter61()
					if !_ok61 {
						break
					}
					var _1326_v31 _dafny.CodePoint
					_1326_v31 = interface{}(_compr_55).(_dafny.CodePoint)
					if (_1325_v33).Contains(_1326_v31) {
						_coll55.Add(_1326_v31, (_this).F6())
					}
				}
				return _coll55.ToMap()
			}()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1286_v0, (_this).F6()))).Contains((func() _dafny.CodePoint {
				if !(!(!(_1289_v3))) {
					return _1286_v0
				}
				return _1286_v0
			})())
			_1317_v24 = (_1317_v24).Update(_1289_v3, _1289_v3)
			var _1327_v34 _dafny.Int
			_ = _1327_v34
			var _1328_v35 _dafny.Int
			_ = _1328_v35
			var _1329_v36 bool
			_ = _1329_v36
			var _out12 _dafny.Int
			_ = _out12
			var _out13 _dafny.Int
			_ = _out13
			var _out14 bool
			_ = _out14
			_out12, _out13, _out14 = (_1320_v27).M6(globalState)
			_1327_v34 = _out12
			_1328_v35 = _out13
			_1329_v36 = _out14
		} else {
			r0 = _dafny.IntOfInt64(569)
			if (_1289_v3) == (_1289_v3) {
				var _1330_v37 _dafny.Map
				_ = _1330_v37
				_1330_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, (_this).F6())
				var _1331_v38 _dafny.Map
				_ = _1331_v38
				_1331_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1330_v37).Cardinality(), _1289_v3)
				var _1332_v39 _dafny.MultiSet
				_ = _1332_v39
				_1332_v39 = _dafny.MultiSetOf(_1331_v38, _1331_v38)
				var _1333_v40 _dafny.Sequence
				_ = _1333_v40
				_1333_v40 = _dafny.SeqOf(_1332_v39)
				(globalState).F5 = ((_1333_v40).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1333_v40).Cardinality()))).Uint32()).(_dafny.MultiSet)).Contains(_1331_v38)
				_1330_v37 = (_1330_v37).Update(_1289_v3, (_this).F6())
				(globalState).F0 = Companion_Default___.Fm24(_dafny.IntOfUint32((_1287_v1).Cardinality()), globalState)
				var _1334_v41 D12
				_ = _1334_v41
				_1334_v41 = Companion_D12_.Create_DC30_(Companion_Default___.Fm37(_1289_v3, (_this).F6(), _1316_v23, globalState))
				var _1335_v42 D12
				_ = _1335_v42
				_1335_v42 = Companion_D12_.Create_DC28_(_1330_v37)
				_1334_v41 = Companion_D12_.Create_DC30_(_1335_v42)
				var _1336_v43 _dafny.Array
				_ = _1336_v43
				var _nw208 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw208
				_1336_v43 = _nw208
				var _1337_v44 _dafny.Array
				_ = _1337_v44
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
				_ = _nw209
				_1337_v44 = _nw209
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_1336_v43), 0))
				_ = _index181
				(_1336_v43).ArraySet1(_1337_v44, (_index181).Int())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(324), _dafny.ArrayLen((_1336_v43), 0))
				_ = _index182
				var _nw210 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(23))
				_ = _nw210
				(_1336_v43).ArraySet1(_nw210, (_index182).Int())
			} else {
				var _1338_v45 _dafny.Set
				_ = _1338_v45
				_1338_v45 = _dafny.SetOf(_1289_v3, _1289_v3)
				var _1339_v46 _dafny.Map
				_ = _1339_v46
				_1339_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1338_v45)
				var _1340_v47 D0
				_ = _1340_v47
				_1340_v47 = Companion_D0_.Create_DC0_(_1339_v46)
				var _1341_v48 *C1
				_ = _1341_v48
				var _nw211 *C1 = New_C1_()
				_ = _nw211
				_nw211.Ctor__()
				_1341_v48 = _nw211
				var _1342_v49 _dafny.Set
				_ = _1342_v49
				_1342_v49 = _dafny.SetOf(_1286_v0)
				var _1343_v50 _dafny.Map
				_ = _1343_v50
				_1343_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1341_v48, (_1342_v49).Cardinality())
				var _1344_v51 _dafny.MultiSet
				_ = _1344_v51
				_1344_v51 = _dafny.MultiSetOf(_dafny.SetOf(!(_1289_v3), _1289_v3))
				var _1345_v52 _dafny.Map
				_ = _1345_v52
				_1345_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm1(_1340_v47, _1289_v3, (_1343_v50).Cardinality(), _1289_v3, globalState), (_1344_v51).Union(_1344_v51))
				_1345_v52 = (_1345_v52).Update(_1289_v3, _1344_v51)
				var _1346_v53 _dafny.Set
				_ = _1346_v53
				_1346_v53 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm2(_dafny.UnicodeSeqOfUtf8Bytes("bq"), globalState), _1287_v1))
				var _1347_v54 _dafny.Map
				_ = _1347_v54
				_1347_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(930))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1348_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1349_i4 _dafny.Int) _dafny.CodePoint {
						return _1348_v0
					}
				})(_1286_v0))), _dafny.IntOfInt64(-965))
				_1346_v53 = (func() _dafny.Set {
					if _1289_v3 {
						return func() _dafny.Set {
							var _coll56 = _dafny.NewBuilder()
							_ = _coll56
							for _iter62 := _dafny.Iterate((_1347_v54).Keys().Elements()); ; {
								_compr_56, _ok62 := _iter62()
								if !_ok62 {
									break
								}
								var _1350_v55 _dafny.Sequence
								_1350_v55 = interface{}(_compr_56).(_dafny.Sequence)
								if (_1347_v54).Contains(_1350_v55) {
									_coll56.Add(_1350_v55)
								}
							}
							return _coll56.ToSet()
						}()
					}
					return _1346_v53
				})()
				(globalState).F1 = _1289_v3
				var _1351_v56 _dafny.Sequence
				_ = _1351_v56
				_1351_v56 = _dafny.SeqOf((_this).F6(), (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt((_this).F6(), (_this).F6())))
				var _1352_v57 _dafny.Array
				_ = _1352_v57
				var _nwElement0_40 _dafny.CodePoint = _1286_v0
				_ = _nwElement0_40
				var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(6))
				_ = _nw212
				(_nw212).ArraySet1CodePoint(_nwElement0_40, 0)
				(_nw212).ArraySet1CodePoint(_1286_v0, 1)
				(_nw212).ArraySet1CodePoint(_1286_v0, 2)
				(_nw212).ArraySet1CodePoint(_1286_v0, 3)
				(_nw212).ArraySet1CodePoint(_1286_v0, 4)
				(_nw212).ArraySet1CodePoint(_1286_v0, 5)
				_1352_v57 = _nw212
				var _1353_v58 _dafny.MultiSet
				_ = _1353_v58
				_1353_v58 = _dafny.MultiSetOf(_1352_v57, _1352_v57)
				var _1354_v59 _dafny.Map
				_ = _1354_v59
				_1354_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1290_v4).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1290_v4).Cardinality()))).Uint32()).(bool), _1289_v3)
				var _rhs191 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_1351_v56, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_1287_v1, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1353_v58).Cardinality()), _dafny.IntOfUint32((_1287_v1).Cardinality()))).Uint32(), _1286_v0), _dafny.Companion_Sequence_.Update(_1287_v1, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1351_v56).Cardinality()), _dafny.IntOfUint32((_1287_v1).Cardinality()))).Uint32(), _1286_v0))).Cardinality()), _dafny.IntOfUint32((_1351_v56).Cardinality()))).Uint32(), (_1354_v59).Cardinality())
				_ = _rhs191
				var _rhs192 _dafny.Int = (func() _dafny.Int {
					if _1289_v3 {
						return (_this).F6()
					}
					return (_this).F6()
				})()
				_ = _rhs192
				var _lhs147 *GlobalState = globalState
				_ = _lhs147
				_1351_v56 = _rhs191
				_lhs147.F2 = _rhs192
				var _1355_v60 _dafny.Array
				_ = _1355_v60
				var _nw213 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw213
				_1355_v60 = _nw213
				var _1356_v61 _dafny.Array
				_ = _1356_v61
				var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
				_ = _nw214
				_1356_v61 = _nw214
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1355_v60), 0))
				_ = _index183
				(_1355_v60).ArraySet1(_1356_v61, (_index183).Int())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(426), _dafny.ArrayLen((_1355_v60), 0))
				_ = _index184
				var _nw215 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw215
				(_1355_v60).ArraySet1(_nw215, (_index184).Int())
			}
			var _1357_v62 _dafny.Map
			_ = _1357_v62
			_1357_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _dafny.SeqOf(_1289_v3))
			var _1358_v63 _dafny.Map
			_ = _1358_v63
			_1358_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_1357_v62).Contains(_1289_v3) {
					return (_1357_v62).Get(_1289_v3).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_1289_v3)
			})(), (Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_1357_v62).Contains(_1289_v3) {
					return (_1357_v62).Get(_1289_v3).(_dafny.Sequence)
				}
				return _dafny.SeqOf(_1289_v3)
			})()).Cardinality()))).Uint32(), _1289_v3), _1289_v3)
			_1358_v63 = (_1358_v63).Update(_dafny.SeqOf(_1289_v3, (_1290_v4).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1290_v4).Cardinality()))).Uint32()).(bool), _1289_v3, _1289_v3, _1289_v3), _1289_v3)
			if _1289_v3 {
				(globalState).F1 = _1289_v3
				r2 = _1289_v3
				var _1359_v64 _dafny.MultiSet
				_ = _1359_v64
				_1359_v64 = _dafny.MultiSetOf((_this).F6(), _dafny.IntOfInt64(201))
				r1 = ((Companion_Default___.Fm15((_this).F6(), _1289_v3, globalState)).Times(_dafny.IntOfInt64(-22))).Cmp((func() _dafny.Int {
					if (_1359_v64).Contains((_dafny.Zero).Minus((_this).F6())) {
						return (_1359_v64).Multiplicity((_dafny.Zero).Minus((_this).F6()))
					}
					return (_this).F6()
				})()) > 0
				var _1360_v65 T1
				_ = _1360_v65
				var _nw216 *C1 = New_C1_()
				_ = _nw216
				_nw216.Ctor__()
				_1360_v65 = _nw216
				_1360_v65 = (func() T1 {
					if !(_1289_v3) {
						return _1360_v65
					}
					return _1360_v65
				})()
				var _1361_v66 *C2
				_ = _1361_v66
				var _nw217 *C2 = New_C2_()
				_ = _nw217
				_nw217.Ctor__(_1287_v1, _1289_v3)
				_1361_v66 = _nw217
			} else {
				(globalState).F5 = _1289_v3
				var _1362_v67 _dafny.Map
				_ = _1362_v67
				_1362_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1286_v0, (_this).F6())
				var _1363_v68 D17
				_ = _1363_v68
				_1363_v68 = Companion_D17_.Create_DC40_(_1362_v67)
				var _1364_v69 _dafny.MultiSet
				_ = _1364_v69
				_1364_v69 = _dafny.MultiSetOf(((_1363_v68).Dtor_cf76()).Merge(_1362_v67), _1362_v67)
				_1364_v69 = (_1364_v69).Update(_1362_v67, Companion_Default___.Abs(_dafny.IntOfInt64(152)))
				var _1365_v70 _dafny.MultiSet
				_ = _1365_v70
				_1365_v70 = _dafny.MultiSetOf((_this).F6())
				var _1366_v71 _dafny.Sequence
				_ = _1366_v71
				_1366_v71 = _dafny.SeqOf(_1365_v70)
				(globalState).F1 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_1365_v70), _1366_v71)
				var _1367_v72 _dafny.Map
				_ = _1367_v72
				_1367_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_this).F6()), true)
				_1367_v72 = _1367_v72
				var _1368_v73 _dafny.Set
				_ = _1368_v73
				_1368_v73 = _dafny.SetOf((_this).F6(), _dafny.IntOfUint32((_1287_v1).Cardinality()), (_this).F6())
				var _1369_v74 _dafny.Sequence
				_ = _1369_v74
				_1369_v74 = _dafny.SeqOf((_this).F6())
				var _1370_v75 T0
				_ = _1370_v75
				var _nw218 *C3 = New_C3_()
				_ = _nw218
				_nw218.Ctor__((_1368_v73).Cardinality(), _1369_v74, (_this).F6())
				_1370_v75 = _nw218
				var _1371_v76 _dafny.Map
				_ = _1371_v76
				_1371_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _1370_v75)
				var _1372_v77 _dafny.Map
				_ = _1372_v77
				_1372_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1371_v76, (_1370_v75).F6())
				var _1373_v78 _dafny.Map
				_ = _1373_v78
				_1373_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _1289_v3)
				(globalState).F1 = (((_1372_v77).Cardinality()).Plus((_1370_v75).F6())).Cmp((_1373_v78).Cardinality()) < 0
			}
			_1286_v0 = _1286_v0
		}
		var _1374_v79 _dafny.Sequence
		_ = _1374_v79
		_1374_v79 = _dafny.SeqOf((_this).F6())
		var _1375_v80 _dafny.Set
		_ = _1375_v80
		_1375_v80 = _dafny.SetOf((_this).F6(), (_this).F6(), _dafny.IntOfUint32(((Companion_D1_.Create_DC4_(_1374_v79)).Dtor_cf11()).Cardinality()), _dafny.IntOfInt64(621), _dafny.IntOfInt64(-963))
		_1289_v3 = !((true) || ((_1375_v80).Equals(_dafny.SetOf((_this).F6(), (_dafny.Zero).Minus((_this).F6())))))
		var _1376_v81 _dafny.Map
		_ = _1376_v81
		_1376_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _1374_v79)
		var _1377_v82 *C3
		_ = _1377_v82
		var _nw219 *C3 = New_C3_()
		_ = _nw219
		_nw219.Ctor__(Companion_Default___.Fm15((_this).F6(), _1289_v3, globalState), (func() _dafny.Sequence {
			if (_1376_v81).Contains(_1289_v3) {
				return (_1376_v81).Get(_1289_v3).(_dafny.Sequence)
			}
			return _dafny.SeqOf((_this).F6(), (_this).F6())
		})(), (_this).F6())
		_1377_v82 = _nw219
		var _1378_v83 _dafny.Set
		_ = _1378_v83
		_1378_v83 = _dafny.SetOf(_1289_v3)
		var _1379_v84 _dafny.Map
		_ = _1379_v84
		_1379_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _1378_v83)
		var _1380_v85 D0
		_ = _1380_v85
		_1380_v85 = Companion_D0_.Create_DC0_(_1379_v84)
		var _1381_v86 _dafny.Map
		_ = _1381_v86
		_1381_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1289_v3, _1289_v3)
		var _1382_v87 _dafny.Array
		_ = _1382_v87
		var _nwElement0_41 bool = _1289_v3
		_ = _nwElement0_41
		var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(26))
		_ = _nw220
		(_nw220).ArraySet1(_nwElement0_41, 0)
		(_nw220).ArraySet1(_1289_v3, 1)
		(_nw220).ArraySet1((func() bool {
			if _1289_v3 {
				return _1289_v3
			}
			return !(_1289_v3)
		})(), 2)
		(_nw220).ArraySet1(false, 3)
		(_nw220).ArraySet1((Companion_D13_.Create_DC33_(true, (_1288_v2).Cardinality(), (_this).F6(), _1289_v3, _1377_v82)).Dtor_cf61(), 4)
		(_nw220).ArraySet1(_1289_v3, 5)
		(_nw220).ArraySet1((_1316_v23).Equals(_1316_v23), 6)
		(_nw220).ArraySet1(_1289_v3, 7)
		(_nw220).ArraySet1(_1289_v3, 8)
		(_nw220).ArraySet1((_1375_v80).IsDisjointFrom(_1375_v80), 9)
		(_nw220).ArraySet1((_1289_v3) == (_1289_v3), 10)
		(_nw220).ArraySet1(_1289_v3, 11)
		(_nw220).ArraySet1(_1289_v3, 12)
		(_nw220).ArraySet1((_1290_v4).Select((Companion_Default___.SafeIndex(Companion_Default___.Fm24((_this).F6(), globalState), _dafny.IntOfUint32((_1290_v4).Cardinality()))).Uint32()).(bool), 13)
		(_nw220).ArraySet1((_1377_v82).Fm16(_1316_v23, globalState), 14)
		(_nw220).ArraySet1((func() bool {
			if !(_1289_v3) {
				return _1289_v3
			}
			return _1289_v3
		})(), 15)
		(_nw220).ArraySet1(_1289_v3, 16)
		(_nw220).ArraySet1(_1289_v3, 17)
		(_nw220).ArraySet1(_1289_v3, 18)
		(_nw220).ArraySet1(_1289_v3, 19)
		(_nw220).ArraySet1(Companion_Default___.Fm1(_1380_v85, _1289_v3, (_1377_v82).F9(), _1289_v3, globalState), 20)
		(_nw220).ArraySet1(!(_1289_v3) || (!(_1289_v3)), 21)
		(_nw220).ArraySet1((func() bool {
			if Companion_Default___.Fm1(_1380_v85, _1289_v3, (_1377_v82).F9(), true, globalState) {
				return _1289_v3
			}
			return true
		})(), 22)
		(_nw220).ArraySet1((_1290_v4).Select((Companion_Default___.SafeIndex((_this).F6(), _dafny.IntOfUint32((_1290_v4).Cardinality()))).Uint32()).(bool), 23)
		(_nw220).ArraySet1(_1289_v3, 24)
		(_nw220).ArraySet1((func() bool {
			if _1289_v3 {
				return (func() bool {
					if (_1381_v86).Contains(_1289_v3) {
						return (_1381_v86).Get(_1289_v3).(bool)
					}
					return false
				})()
			}
			return _1289_v3
		})(), 25)
		_1382_v87 = _nw220
		var _1383_v88 _dafny.Array
		_ = _1383_v88
		var _nw221 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
		_ = _nw221
		_1383_v88 = _nw221
		var _1384_v89 _dafny.MultiSet
		_ = _1384_v89
		_1384_v89 = _dafny.MultiSetOf(_1383_v88, _1383_v88, _1383_v88)
		var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1382_v87), 0))
		_ = _index185
		(_1382_v87).ArraySet1(((func() _dafny.Int {
			if (_1384_v89).Contains(_1383_v88) {
				return (_1384_v89).Multiplicity(_1383_v88)
			}
			return (_1377_v82).F9()
		})()).Cmp(_dafny.IntOfUint32((_1287_v1).Cardinality())) != 0, (_index185).Int())
		var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(235), _dafny.ArrayLen((_1382_v87), 0))
		_ = _index186
		(_1382_v87).ArraySet1(_1289_v3, (_index186).Int())
		var _1385_v90 _dafny.MultiSet
		_ = _1385_v90
		_1385_v90 = _dafny.MultiSetOf((_1377_v82).Fm17((_1377_v82).F9(), _1375_v80, globalState))
		var _1386_v91 _dafny.Map
		_ = _1386_v91
		_1386_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D16_.Create_DC38_(_1385_v90), _1375_v80)
		r0 = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_1377_v82).F9()), (_1386_v91).Cardinality())
		r1 = _1289_v3
		r2 = (_1316_v23).IsProperSubsetOf(_1316_v23)
		return r0, r1, r2
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
