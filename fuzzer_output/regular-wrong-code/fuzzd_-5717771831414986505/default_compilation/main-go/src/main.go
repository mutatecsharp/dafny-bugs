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
func (_static *CompanionStruct_Default___) Fm5(globalState *GlobalState) _dafny.Sequence {
	if !_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(true), _dafny.SeqOf(false)) {
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(988))).Uint32(), func(coer0 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_0_i0 _dafny.Int) _dafny.Sequence {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(811))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg1 _dafny.Int) interface{} {
					return coer1(arg1)
				}
			}(func(_1_i1 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('k')
			}))
		}))
	} else {
		return _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("tuqlhd"), _dafny.UnicodeSeqOfUtf8Bytes("wcvbklxx"))
	}
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(297))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Int {
		return Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(411), _dafny.IntOfInt64(749))
	}))
}
func (_static *CompanionStruct_Default___) Fm9(globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), !(true))))
}
func (_static *CompanionStruct_Default___) Fm11(p0 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(307))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg3 _dafny.Int) interface{} {
			return coer3(arg3)
		}
	}(func(_3_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('o')
	}))
}
func (_static *CompanionStruct_Default___) Fm12(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(Companion_Default___.SafeDivisionInt((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-509))).Cardinality(), _dafny.IntOfInt64(214)))
}
func (_static *CompanionStruct_Default___) Fm13(p0 _dafny.Int, p1 D2, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gyflnn"), true)).Merge(func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg4 _dafny.Int) interface{} {
				return coer4(arg4)
			}
		}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('e')
		})), !(true))).Keys().Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _5_v0 _dafny.Sequence
			_5_v0 = interface{}(_compr_0).(_dafny.Sequence)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(375))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_4_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('e')
			})), !(true))).Contains(_5_v0) {
				_coll0.Add(_5_v0, false)
			}
		}
		return _coll0.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Map, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(true)).Difference(_dafny.MultiSetOf(false, !(true)))).Intersection(_dafny.MultiSetOf(true, true, true))
}
func (_static *CompanionStruct_Default___) Fm15(p0 bool, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.IntOfInt64(193), (_dafny.SetOf(_dafny.CodePoint('l'))).Cardinality(), _dafny.IntOfInt64(546))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(545), _dafny.IntOfInt64(965)))).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(808)))
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 _dafny.Map, p2 _dafny.Sequence, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('k')
}
func (_static *CompanionStruct_Default___) Fm17(globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("qrabdg")
}
func (_static *CompanionStruct_Default___) Fm18(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(189))).Uint32(), func(coer6 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg6 _dafny.Int) interface{} {
			return coer6(arg6)
		}
	}(func(_6_i0 _dafny.Int) _dafny.Int {
		return _dafny.IntOfUint32((_dafny.SeqOf(!(true), !(true))).Cardinality())
	})), _dafny.SeqOf((_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(-533), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm19(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if true {
			return (Companion_D7_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf((_dafny.MultiSetFromSeq(_dafny.SeqOf(true))).Cardinality())).Cardinality())))).Dtor_cf28()
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('u'))).Cardinality()))
	})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(true)), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("n")).Cardinality())))
}
func (_static *CompanionStruct_Default___) Fm20(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(false, !(false)), false)).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), !(true))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(true), true)))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, globalState *GlobalState) D3 {
	var _source0 D5 = Companion_D5_.Create_DC15_(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(445))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg7 _dafny.Int) interface{} {
			return coer7(arg7)
		}
	}(func(_7_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('n')
	})))
	_ = _source0
	if _source0.Is_DC16() {
		var _8___mcc_h0 D3 = _source0.Get_().(D5_DC16).Cf19
		_ = _8___mcc_h0
		var _9___mcc_h1 _dafny.Int = _source0.Get_().(D5_DC16).Cf20
		_ = _9___mcc_h1
		var _10___mcc_h2 _dafny.Int = _source0.Get_().(D5_DC16).Cf21
		_ = _10___mcc_h2
		var _11___mcc_h3 _dafny.Int = _source0.Get_().(D5_DC16).Cf22
		_ = _11___mcc_h3
		var _12___mcc_h4 *C0 = _source0.Get_().(D5_DC16).Cf23
		_ = _12___mcc_h4
		var _13_cf23 *C0 = _12___mcc_h4
		_ = _13_cf23
		var _14_cf22 _dafny.Int = _11___mcc_h3
		_ = _14_cf22
		var _15_cf21 _dafny.Int = _10___mcc_h2
		_ = _15_cf21
		var _16_cf20 _dafny.Int = _9___mcc_h1
		_ = _16_cf20
		var _17_cf19 D3 = _8___mcc_h0
		_ = _17_cf19
		return Companion_D3_.Create_DC10_(true, _16_cf20)
	} else {
		var _18___mcc_h5 _dafny.Sequence = _source0.Get_().(D5_DC15).Cf18
		_ = _18___mcc_h5
		var _19_cf18 _dafny.Sequence = _18___mcc_h5
		_ = _19_cf18
		return Companion_D3_.Create_DC10_(true, _dafny.IntOfInt64(713))
	}
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) D1 {
	if (func() bool {
		if true {
			return false
		}
		return false
	})() {
		return Companion_D1_.Create_DC3_(Companion_D0_.Create_DC1_(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(false))).Cardinality())
	} else {
		return Companion_D1_.Create_DC3_(Companion_D0_.Create_DC1_(), (func() _dafny.Map {
			var _coll1 = _dafny.NewMapBuilder()
			_ = _coll1
			for _iter1 := _dafny.Iterate((_dafny.SeqOf(_dafny.IntOfInt64(470), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(334))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg8 _dafny.Int) interface{} {
					return coer8(arg8)
				}
			}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('n')
			}))).Cardinality()), _dafny.IntOfInt64(178))).Elements()); ; {
				_compr_1, _ok1 := _iter1()
				if !_ok1 {
					break
				}
				var _21_v0 _dafny.Int
				_21_v0 = interface{}(_compr_1).(_dafny.Int)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.IntOfInt64(470), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(334))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg9 _dafny.Int) interface{} {
						return coer9(arg9)
					}
				}(func(_20_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('n')
				}))).Cardinality()), _dafny.IntOfInt64(178)), _21_v0) {
					_coll1.Add((_21_v0).Plus(_dafny.IntOfInt64(493)), _dafny.IntOfInt64(650))
				}
			}
			return _coll1.ToMap()
		}()).Cardinality())
	}
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(false)
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.IntOfInt64(670), (_dafny.Zero).Minus((_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfInt64(900), _dafny.IntOfInt64(919), _dafny.IntOfInt64(804), _dafny.IntOfInt64(901))).Cardinality(), _dafny.IntOfInt64(-743))).Cardinality()))).Difference((_dafny.SetOf(_dafny.IntOfInt64(-722), _dafny.IntOfInt64(152), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, true)).Cardinality()), _dafny.IntOfInt64(-562))).Cardinality()), _dafny.IntOfInt64(865), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(648))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_22_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('q')
	}))).Cardinality()), _dafny.IntOfInt64(132))).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(800), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(331))).Cardinality())), _dafny.IntOfInt64(-891))).Cardinality())).Difference(_dafny.SetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality(), _dafny.IntOfInt64(-476), _dafny.IntOfInt64(959), _dafny.IntOfUint32((_dafny.SeqOf(false, true)).Cardinality()))))
}
func (_static *CompanionStruct_Default___) Fm25(p0 bool, p1 _dafny.Int, p2 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_dafny.SeqOf(!(true), false, false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm26(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(212), _dafny.IntOfInt64(741))).Cardinality()), _dafny.IntOfInt64(393)), true)).Merge(func() _dafny.Map {
		var _coll2 = _dafny.NewMapBuilder()
		_ = _coll2
		for _iter2 := _dafny.Iterate((_dafny.SeqOf(func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-93), _dafny.IntOfInt64(105))); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _23_v1 _dafny.Int
				_23_v1 = interface{}(_compr_3).(_dafny.Int)
				if ((_dafny.IntOfInt64(-93)).Cmp(_23_v1) <= 0) && ((_23_v1).Cmp(_dafny.IntOfInt64(105)) < 0) {
					_coll3.Add((_23_v1).Minus(_dafny.IntOfInt64(511)), _dafny.IntOfInt64(96))
				}
			}
			return _coll3.ToMap()
		}())).Elements()); ; {
			_compr_2, _ok2 := _iter2()
			if !_ok2 {
				break
			}
			var _24_v0 _dafny.Map
			_24_v0 = interface{}(_compr_2).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(func() _dafny.Map {
				var _coll4 = _dafny.NewMapBuilder()
				_ = _coll4
				for _iter4 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-93), _dafny.IntOfInt64(105))); ; {
					_compr_4, _ok4 := _iter4()
					if !_ok4 {
						break
					}
					var _23_v1 _dafny.Int
					_23_v1 = interface{}(_compr_4).(_dafny.Int)
					if ((_dafny.IntOfInt64(-93)).Cmp(_23_v1) <= 0) && ((_23_v1).Cmp(_dafny.IntOfInt64(105)) < 0) {
						_coll4.Add((_23_v1).Minus(_dafny.IntOfInt64(511)), _dafny.IntOfInt64(96))
					}
				}
				return _coll4.ToMap()
			}()), _24_v0) {
				_coll2.Add(_24_v0, false)
			}
		}
		return _coll2.ToMap()
	}())
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 bool, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Map {
	return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(790), (((Companion_D2_.Create_DC4_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-544), !(true))).Cardinality(), (_dafny.MultiSetOf((_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfInt64(-406)), !(true))).Cardinality())).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sqqfutbxm")).Cardinality()))).Cardinality()))).Dtor_cf4()).Difference(_dafny.SetOf(_dafny.IntOfInt64(854)))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	if false {
		return (_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-750), (_dafny.SetOf(_dafny.IntOfInt64(-201), _dafny.IntOfInt64(521))).Cardinality()), func() _dafny.Map {
			var _coll5 = _dafny.NewMapBuilder()
			_ = _coll5
			for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-999), _dafny.IntOfInt64(240))); ; {
				_compr_5, _ok5 := _iter5()
				if !_ok5 {
					break
				}
				var _25_v0 _dafny.Int
				_25_v0 = interface{}(_compr_5).(_dafny.Int)
				if ((_dafny.IntOfInt64(-999)).Cmp(_25_v0) <= 0) && ((_25_v0).Cmp(_dafny.IntOfInt64(240)) < 0) {
					_coll5.Add((_25_v0).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(-219))).Cardinality())), _dafny.IntOfInt64(-167))
				}
			}
			return _coll5.ToMap()
		}())).Difference(_dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(245), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('a'), _dafny.CodePoint('p'))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-699), (_dafny.SetOf(true)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(999), _dafny.IntOfInt64(629)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true)).Cardinality())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), _dafny.IntOfInt64(52))).Cardinality(), true)).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gsppdqr")).Cardinality())), _dafny.IntOfInt64(272))))
	} else if true {
		return _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(553), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(994))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_26_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('c')
		}))).Cardinality()))).Cardinality())).Cardinality(), (_dafny.MultiSetFromSeq(_dafny.SeqOf(Companion_D18_.Create_DC49_(_dafny.SeqOf(true), (_dafny.MultiSetOf(_dafny.CodePoint('c'))).Cardinality())))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_dafny.IntOfInt64(-594))).Cardinality(), _dafny.IntOfInt64(913)))
	} else {
		return _dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(78), _dafny.IntOfInt64(765)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(665), (_dafny.SetOf(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-553))).Uint32(), func(coer12 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_27_i1 _dafny.Int) _dafny.Map {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				var _coll6 = _dafny.NewBuilder()
				_ = _coll6
				for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(414), _dafny.IntOfInt64(809))); ; {
					_compr_6, _ok6 := _iter6()
					if !_ok6 {
						break
					}
					var _28_v1 _dafny.Int
					_28_v1 = interface{}(_compr_6).(_dafny.Int)
					if ((_dafny.IntOfInt64(414)).Cmp(_28_v1) <= 0) && ((_28_v1).Cmp(_dafny.IntOfInt64(809)) < 0) {
						_coll6.Add((_28_v1).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(232))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg13 _dafny.Int) interface{} {
								return coer13(arg13)
							}
						}(func(_29_i2 _dafny.Int) _dafny.CodePoint {
							return _dafny.CodePoint('y')
						}))).Cardinality())))
					}
				}
				return _coll6.ToSet()
			}()).Cardinality(), _dafny.IntOfInt64(906))
		}))).Cardinality())))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(458), _dafny.IntOfInt64(240))))
	}
}
func (_static *CompanionStruct_Default___) Fm29(p0 _dafny.Sequence, p1 _dafny.Sequence, p2 _dafny.MultiSet, globalState *GlobalState) D7 {
	return Companion_D7_.Create_DC21_(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf((_dafny.MultiSetOf(_dafny.CodePoint('u'))).Cardinality())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))).Cardinality()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qyyg")).Cardinality())), _dafny.IntOfInt64(29))).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm30(globalState *GlobalState) D11 {
	return Companion_D11_.Create_DC32_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false))
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(700), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-206), true), func() _dafny.Map {
		var _coll7 = _dafny.NewMapBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(137), _dafny.IntOfInt64(-918))); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _30_v0 _dafny.Int
			_30_v0 = interface{}(_compr_7).(_dafny.Int)
			if ((_dafny.IntOfInt64(137)).Cmp(_30_v0) <= 0) && ((_30_v0).Cmp(_dafny.IntOfInt64(-918)) < 0) {
				_coll7.Add((_30_v0).Times(_dafny.IntOfInt64(-89)), false)
			}
		}
		return _coll7.ToMap()
	}(), func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(432), _dafny.IntOfInt64(305))).Keys().Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _31_v1 _dafny.Int
			_31_v1 = interface{}(_compr_8).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(432), _dafny.IntOfInt64(305))).Contains(_31_v1) {
				_coll8.Add(Companion_Default___.SafeModuloInt(_31_v1, _dafny.IntOfInt64(45)), true)
			}
		}
		return _coll8.ToMap()
	}(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(836), !(false)))).Intersection(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(false)).Cardinality(), !(true))), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(-814))).Cardinality()), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("tlws")).Cardinality()))).Cardinality(), false), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(292))).Cardinality()))).Cardinality(), true), func() _dafny.Map {
		var _coll9 = _dafny.NewMapBuilder()
		_ = _coll9
		for _iter9 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(561), _dafny.IntOfInt64(572))); ; {
			_compr_9, _ok9 := _iter9()
			if !_ok9 {
				break
			}
			var _32_v2 _dafny.Int
			_32_v2 = interface{}(_compr_9).(_dafny.Int)
			if ((_dafny.IntOfInt64(561)).Cmp(_32_v2) <= 0) && ((_32_v2).Cmp(_dafny.IntOfInt64(572)) < 0) {
				_coll9.Add(Companion_Default___.SafeDivisionInt(_32_v2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gswko")).Cardinality())), false)
			}
		}
		return _coll9.ToMap()
	}()))
}
func (_static *CompanionStruct_Default___) Fm32(p0 bool, p1 bool, p2 bool, globalState *GlobalState) _dafny.Set {
	return (_dafny.SetOf(false, false)).Difference(_dafny.SetOf(true))
}
func (_static *CompanionStruct_Default___) Fm33(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((func() _dafny.CodePoint {
		if !(true) {
			return _dafny.CodePoint('h')
		}
		return _dafny.CodePoint('t')
	})())
}
func (_static *CompanionStruct_Default___) Fm34(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer14 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
		return func(arg14 _dafny.Int) interface{} {
			return coer14(arg14)
		}
	}(func(_33_i0 _dafny.Int) _dafny.MultiSet {
		return _dafny.MultiSetOf(_dafny.CodePoint('y'), _dafny.CodePoint('q'))
	}))
}
func (_static *CompanionStruct_Default___) Fm35(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) bool {
	return !(func(_source1 D2) bool {
		if _source1.Is_DC5() {
			var _34___mcc_h0 bool = _source1.Get_().(D2_DC5).Cf5
			_ = _34___mcc_h0
			var _35_cf5 bool = _34___mcc_h0
			_ = _35_cf5
			return (_dafny.SetOf(_35_cf5, _35_cf5, _35_cf5, _35_cf5)).IsProperSubsetOf(_dafny.SetOf(_35_cf5))
		} else if _source1.Is_DC6() {
			return (false) == (!(!(false)))
		} else if _source1.Is_DC7() {
			var _36___mcc_h1 _dafny.CodePoint = _source1.Get_().(D2_DC7).Cf6
			_ = _36___mcc_h1
			var _37___mcc_h2 _dafny.Map = _source1.Get_().(D2_DC7).Cf7
			_ = _37___mcc_h2
			var _38___mcc_h3 _dafny.Int = _source1.Get_().(D2_DC7).Cf8
			_ = _38___mcc_h3
			var _39_cf8 _dafny.Int = _38___mcc_h3
			_ = _39_cf8
			var _40_cf7 _dafny.Map = _37___mcc_h2
			_ = _40_cf7
			var _41_cf6 _dafny.CodePoint = _36___mcc_h1
			_ = _41_cf6
			return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("fucxg"), _dafny.UnicodeSeqOfUtf8Bytes("gwvrtg"))
		} else if _source1.Is_DC4() {
			var _42___mcc_h4 _dafny.Set = _source1.Get_().(D2_DC4).Cf4
			_ = _42___mcc_h4
			var _43_cf4 _dafny.Set = _42___mcc_h4
			_ = _43_cf4
			return !(false)
		} else {
			var _44___mcc_h5 D2 = _source1.Get_().(D2_DC8).Cf9
			_ = _44___mcc_h5
			var _45_cf9 D2 = _44___mcc_h5
			_ = _45_cf9
			return true
		}
	}(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC4_(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.SetOf(_dafny.CodePoint('s'), _dafny.CodePoint('t'))).Cardinality())).Cardinality(), _dafny.IntOfInt64(870), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-713), false)).Cardinality()))))))
}
func (_static *CompanionStruct_Default___) Fm36(p0 _dafny.Map, p1 bool, p2 _dafny.Set, globalState *GlobalState) D8 {
	return Companion_D8_.Create_DC23_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(!(false)), _dafny.IntOfInt64(379)))
}
func (_static *CompanionStruct_Default___) Fm37(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('s'), true)).Merge(func() _dafny.Map {
		var _coll10 = _dafny.NewMapBuilder()
		_ = _coll10
		for _iter10 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg15 _dafny.Int) interface{} {
				return coer15(arg15)
			}
		}(func(_46_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))).Elements()); ; {
			_compr_10, _ok10 := _iter10()
			if !_ok10 {
				break
			}
			var _47_v0 _dafny.CodePoint
			_47_v0 = interface{}(_compr_10).(_dafny.CodePoint)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(193))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg16 _dafny.Int) interface{} {
					return coer16(arg16)
				}
			}(func(_46_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('t')
			})), _47_v0) {
				_coll10.Add(_47_v0, !(false))
			}
		}
		return _coll10.ToMap()
	}())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('g'), false))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _48_v0 _dafny.Sequence
	_ = _48_v0
	_48_v0 = _dafny.UnicodeSeqOfUtf8Bytes("rwdkj")
	var _49_v1 _dafny.Int
	_ = _49_v1
	_49_v1 = _dafny.IntOfInt64(143)
	var _50_v2 _dafny.CodePoint
	_ = _50_v2
	_50_v2 = _dafny.CodePoint('w')
	var _51_v3 _dafny.Array
	_ = _51_v3
	var _len0_0 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_0
	var _nw0 _dafny.Array
	_ = _nw0
	if _len0_0.Cmp(_dafny.Zero) == 0 {
		_nw0 = _dafny.NewArray(_len0_0)
	} else {
		var _init0 func(_dafny.Int) _dafny.CodePoint = func(_52_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('n')
		}
		_ = _init0
		var _element0_0 = _init0(_dafny.Zero)
		_ = _element0_0
		_nw0 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
		(_nw0).ArraySet1CodePoint(_element0_0, 0)
		var _nativeLen0_0 = (_len0_0).Int()
		_ = _nativeLen0_0
		for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
			(_nw0).ArraySet1CodePoint(_init0(_dafny.IntOf(_i0_0)), _i0_0)
		}
	}
	_51_v3 = _nw0
	var _53_globalState *GlobalState
	_ = _53_globalState
	var _nw1 *GlobalState = New_GlobalState_()
	_ = _nw1
	_nw1.Ctor__(_dafny.IntOfInt64(287), _dafny.Companion_Sequence_.Concatenate(_48_v0, _dafny.Companion_Sequence_.Update(_48_v0, (Companion_Default___.SafeIndex((_dafny.Zero).Minus(_49_v1), _dafny.IntOfUint32((_48_v0).Cardinality()))).Uint32(), _50_v2)), _51_v3)
	_53_globalState = _nw1
	var _54_v4 bool
	_ = _54_v4
	_54_v4 = true
	var _55_v5 _dafny.Map
	_ = _55_v5
	_55_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, _54_v4)
	var _56_v6 _dafny.Array
	_ = _56_v6
	var _nw2 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
	_ = _nw2
	_56_v6 = _nw2
	var _57_v7 _dafny.Set
	_ = _57_v7
	_57_v7 = _dafny.SetOf(_dafny.CodePoint('o'), _50_v2)
	var _58_v8 _dafny.Sequence
	_ = _58_v8
	_58_v8 = _dafny.SeqOf(_48_v0, _dafny.Companion_Sequence_.Update(_48_v0, (Companion_Default___.SafeIndex((_57_v7).Cardinality(), _dafny.IntOfUint32((_48_v0).Cardinality()))).Uint32(), _50_v2))
	var _59_v9 _dafny.MultiSet
	_ = _59_v9
	_59_v9 = _dafny.MultiSetOf(false)
	var _60_v10 *C5
	_ = _60_v10
	var _nw3 *C5 = New_C5_()
	_ = _nw3
	_nw3.Ctor__(_49_v1, (func() bool {
		if (_55_v5).Contains(_50_v2) {
			return (_55_v5).Get(_50_v2).(bool)
		}
		return _54_v4
	})(), _56_v6, (func() bool {
		if _54_v4 {
			return _54_v4
		}
		return _54_v4
	})(), (func() bool {
		if Companion_Default___.Fm35(_58_v8, _54_v4, _53_globalState) {
			return _54_v4
		}
		return false
	})(), _59_v9)
	_60_v10 = _nw3
	_54_v4 = (_60_v10).F10()
	var _61_i1 _dafny.Int
	_ = _61_i1
	_61_i1 = _dafny.Zero
	{
		for (_60_v10).F10() {
			{
				if (_61_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_61_i1 = (_61_i1).Plus(_dafny.One)
				var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
				_ = _index0
				(_56_v6).ArraySet1((_60_v10).F10(), (_index0).Int())
				var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
				_ = _index1
				(_56_v6).ArraySet1(_54_v4, (_index1).Int())
				if (_60_v10).F10() {
					var _62_v11 _dafny.Map
					_ = _62_v11
					_62_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("cfhadr"), (_60_v10).F10())
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
					_ = _index2
					var _rhs0 _dafny.Int = ((_60_v10).Fm4(_49_v1, _dafny.IntOfInt64(364), _62_v11, (_60_v10).F9(), _53_globalState)).Minus(_dafny.IntOfInt64(589))
					_ = _rhs0
					var _rhs1 bool = (func() bool {
						if (_60_v10).F10() {
							return !(!(!((_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool))))
						}
						return _54_v4
					})()
					_ = _rhs1
					var _lhs0 *GlobalState = _53_globalState
					_ = _lhs0
					var _lhs1 _dafny.Array = _56_v6
					_ = _lhs1
					var _lhs2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
					_ = _lhs2
					_lhs0.F0 = _rhs0
					(_lhs1).ArraySet1(_rhs1, (_lhs2).Int())
					var _63_v12 _dafny.Map
					_ = _63_v12
					_63_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F9(), (_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool))
					(_60_v10).M1((_60_v10).F9(), (_60_v10).F10(), (func() bool {
						if (_63_v12).Contains((_60_v10).F9()) {
							return (_63_v12).Get((_60_v10).F9()).(bool)
						}
						return (_60_v10).F10()
					})(), _53_globalState)
					var _64_v13 _dafny.Sequence
					_ = _64_v13
					_64_v13 = _dafny.SeqOf(_49_v1, (_dafny.Zero).Minus((_60_v10).F9()))
					var _65_v14 _dafny.Sequence
					_ = _65_v14
					_65_v14 = _dafny.SeqOf((_64_v13).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.IntOfUint32((_64_v13).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_48_v0).Cardinality()))
					(_53_globalState).F0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_64_v13).Cardinality()), (_60_v10).F9()))
					(_53_globalState).F0 = (func() _dafny.Int {
						if true {
							return (_49_v1).Minus((_60_v10).F9())
						}
						return _49_v1
					})()
					_54_v4 = _dafny.Companion_Sequence_.Contains(_48_v0, _50_v2)
				} else {
					(_53_globalState).F0 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("qr")).Cardinality())
					var _66_v15 _dafny.Set
					_ = _66_v15
					_66_v15 = _dafny.SetOf(_49_v1)
					var _67_v16 _dafny.Map
					_ = _67_v16
					_67_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F9(), (_60_v10).F9())
					var _68_v17 _dafny.Map
					_ = _68_v17
					_68_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm16((_60_v10).F10(), _67_v16, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(970))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg17 _dafny.Int) interface{} {
							return coer17(arg17)
						}
					}((func(_69_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_70_i2 _dafny.Int) _dafny.CodePoint {
							return _69_v2
						}
					})(_50_v2))), _53_globalState))
					var _71_v18 D19
					_ = _71_v18
					_71_v18 = Companion_D19_.Create_DC51_(_68_v17)
					var _72_v19 _dafny.Map
					_ = _72_v19
					_72_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v0, (_60_v10).F10())
					var _73_v21 _dafny.Sequence
					_ = _73_v21
					_73_v21 = _dafny.SeqOf((_60_v10).F9())
					var _74_v22 _dafny.Sequence
					_ = _74_v22
					_74_v22 = _dafny.SeqOf(_54_v4)
					var _75_v23 _dafny.Array
					_ = _75_v23
					var _nwElement0_0 _dafny.Int = (_60_v10).F9()
					_ = _nwElement0_0
					var _nw4 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(27))
					_ = _nw4
					(_nw4).ArraySet1(_nwElement0_0, 0)
					(_nw4).ArraySet1((_60_v10).F9(), 1)
					(_nw4).ArraySet1((_60_v10).F9(), 2)
					(_nw4).ArraySet1((_60_v10).F9(), 3)
					(_nw4).ArraySet1((_66_v15).Cardinality(), 4)
					(_nw4).ArraySet1((_60_v10).F9(), 5)
					(_nw4).ArraySet1((_60_v10).Fm4(_dafny.IntOfInt64(825), (_dafny.Zero).Minus(((_71_v18).Dtor_cf81()).Cardinality()), _72_v19, (_60_v10).F9(), _53_globalState), 6)
					(_nw4).ArraySet1(_49_v1, 7)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_48_v0).Cardinality()), 8)
					(_nw4).ArraySet1((func() _dafny.Map {
						var _coll11 = _dafny.NewMapBuilder()
						_ = _coll11
						for _iter11 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(473), _dafny.IntOfInt64(327))); ; {
							_compr_11, _ok11 := _iter11()
							if !_ok11 {
								break
							}
							var _76_v20 _dafny.Int
							_76_v20 = interface{}(_compr_11).(_dafny.Int)
							if ((_dafny.IntOfInt64(473)).Cmp(_76_v20) <= 0) && ((_76_v20).Cmp(_dafny.IntOfInt64(327)) < 0) {
								_coll11.Add(Companion_Default___.SafeDivisionInt(_76_v20, (_60_v10).F9()), _50_v2)
							}
						}
						return _coll11.ToMap()
					}()).Cardinality(), 9)
					(_nw4).ArraySet1((_60_v10).F9(), 10)
					(_nw4).ArraySet1((_60_v10).F9(), 11)
					(_nw4).ArraySet1(_49_v1, 12)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-561))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg18 _dafny.Int) interface{} {
							return coer18(arg18)
						}
					}((func(_77_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_78_i3 _dafny.Int) _dafny.CodePoint {
							return _77_v2
						}
					})(_50_v2)))).Cardinality()), 13)
					(_nw4).ArraySet1((_60_v10).F9(), 14)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_73_v21).Cardinality()), 15)
					(_nw4).ArraySet1((_60_v10).F9(), 16)
					(_nw4).ArraySet1((_60_v10).F9(), 17)
					(_nw4).ArraySet1((_60_v10).F9(), 18)
					(_nw4).ArraySet1(_dafny.IntOfInt64(307), 19)
					(_nw4).ArraySet1((_60_v10).F9(), 20)
					(_nw4).ArraySet1((_60_v10).F9(), 21)
					(_nw4).ArraySet1((_60_v10).F9(), 22)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_48_v0).Cardinality()), 23)
					(_nw4).ArraySet1((_60_v10).Fm4((_60_v10).F9(), _dafny.IntOfUint32((_74_v22).Cardinality()), _72_v19, (_60_v10).F9(), _53_globalState), 24)
					(_nw4).ArraySet1((_60_v10).F9(), 25)
					(_nw4).ArraySet1(_dafny.IntOfUint32((_74_v22).Cardinality()), 26)
					_75_v23 = _nw4
					var _79_v24 _dafny.Sequence
					_ = _79_v24
					_79_v24 = _dafny.SeqOf(_75_v23)
					var _80_v25 _dafny.Set
					_ = _80_v25
					_80_v25 = _dafny.SetOf((_60_v10).F10())
					var _81_v26 _dafny.MultiSet
					_ = _81_v26
					_81_v26 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool), (_60_v10).F10())).Cardinality(), (_dafny.Zero).Minus((_60_v10).F9()), (_60_v10).F9())
					var _82_v28 _dafny.Set
					_ = _82_v28
					_82_v28 = _dafny.SetOf(_48_v0)
					var _83_v29 _dafny.MultiSet
					_ = _83_v29
					_83_v29 = _dafny.MultiSetOf(_dafny.CodePoint('d'), _50_v2, _50_v2)
					var _84_v30 _dafny.Array
					_ = _84_v30
					var _nwElement0_1 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_79_v24, _79_v24)).Cardinality())
					_ = _nwElement0_1
					var _nw5 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(21))
					_ = _nw5
					(_nw5).ArraySet1(_nwElement0_1, 0)
					(_nw5).ArraySet1(Companion_Default___.SafeModuloInt((_80_v25).Cardinality(), (_60_v10).F9()), 1)
					(_nw5).ArraySet1((_60_v10).F9(), 2)
					(_nw5).ArraySet1(_dafny.IntOfInt64(733), 3)
					(_nw5).ArraySet1(_49_v1, 4)
					(_nw5).ArraySet1((_81_v26).Cardinality(), 5)
					(_nw5).ArraySet1((_60_v10).F9(), 6)
					(_nw5).ArraySet1(_49_v1, 7)
					(_nw5).ArraySet1(_49_v1, 8)
					(_nw5).ArraySet1(((_60_v10).F9()).Plus((_66_v15).Cardinality()), 9)
					(_nw5).ArraySet1((_60_v10).Fm4((_60_v10).F9(), _dafny.IntOfInt64(520), (func() _dafny.Map {
						var _coll12 = _dafny.NewMapBuilder()
						_ = _coll12
						for _iter12 := _dafny.Iterate((_82_v28).Elements()); ; {
							_compr_12, _ok12 := _iter12()
							if !_ok12 {
								break
							}
							var _85_v27 _dafny.Sequence
							_85_v27 = interface{}(_compr_12).(_dafny.Sequence)
							if (_82_v28).Contains(_85_v27) {
								_coll12.Add(_85_v27, (_60_v10).F10())
							}
						}
						return _coll12.ToMap()
					}()).Update(_48_v0, _54_v4), (_83_v29).Cardinality(), _53_globalState), 10)
					(_nw5).ArraySet1(_dafny.IntOfInt64(588), 11)
					(_nw5).ArraySet1(_49_v1, 12)
					(_nw5).ArraySet1((_60_v10).F9(), 13)
					(_nw5).ArraySet1((func() _dafny.Int {
						if (_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool) {
							return (_60_v10).F9()
						}
						return (_60_v10).Fm4((_68_v17).Cardinality(), (_dafny.Zero).Minus((_60_v10).F9()), _72_v19, _49_v1, _53_globalState)
					})(), 14)
					(_nw5).ArraySet1(((_60_v10).F9()).Plus((_60_v10).F9()), 15)
					(_nw5).ArraySet1((_60_v10).F9(), 16)
					(_nw5).ArraySet1((_60_v10).F9(), 17)
					(_nw5).ArraySet1(((_60_v10).F9()).Plus(_dafny.IntOfUint32((_73_v21).Cardinality())), 18)
					(_nw5).ArraySet1(_49_v1, 19)
					(_nw5).ArraySet1(_dafny.IntOfUint32((_48_v0).Cardinality()), 20)
					_84_v30 = _nw5
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_84_v30), 0))
					_ = _index3
					(_84_v30).ArraySet1((_60_v10).F9(), (_index3).Int())
					var _86_v31 D19
					_ = _86_v31
					_86_v31 = Companion_D19_.Create_DC52_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_60_v10).F10()), _49_v1)).Cardinality())
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_84_v30), 0))
					_ = _index4
					(_84_v30).ArraySet1(Companion_Default___.SafeDivisionInt((_86_v31).Dtor_cf82(), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F9(), true)).Merge(func() _dafny.Map {
						var _coll13 = _dafny.NewMapBuilder()
						_ = _coll13
						for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(779), _dafny.IntOfInt64(783))); ; {
							_compr_13, _ok13 := _iter13()
							if !_ok13 {
								break
							}
							var _87_v32 _dafny.Int
							_87_v32 = interface{}(_compr_13).(_dafny.Int)
							if ((_dafny.IntOfInt64(779)).Cmp(_87_v32) <= 0) && ((_87_v32).Cmp(_dafny.IntOfInt64(783)) < 0) {
								_coll13.Add((_87_v32).Times(_dafny.IntOfInt64(742)), _54_v4)
							}
						}
						return _coll13.ToMap()
					}())).Cardinality()), (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(668), _dafny.ArrayLen((_84_v30), 0))
					_ = _index5
					(_84_v30).ArraySet1(_49_v1, (_index5).Int())
					_54_v4 = (_60_v10).F10()
					var _88_v33 _dafny.Array
					_ = _88_v33
					var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(3))
					_ = _nw6
					_88_v33 = _nw6
					var _89_v34 *C4
					_ = _89_v34
					var _nw7 *C4 = New_C4_()
					_ = _nw7
					_nw7.Ctor__(_88_v33, _54_v4, false, _59_v9)
					_89_v34 = _nw7
					var _90_v35 _dafny.Map
					_ = _90_v35
					_90_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v4, _89_v34)
					(_53_globalState).F0 = (_60_v10).Fm4((_90_v35).Cardinality(), (_60_v10).F9(), _72_v19, _dafny.IntOfUint32((_48_v0).Cardinality()), _53_globalState)
				}
				var _91_v36 _dafny.Map
				_ = _91_v36
				_91_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_59_v9, _48_v0)
				var _92_v37 D16
				_ = _92_v37
				_92_v37 = Companion_D16_.Create_DC44_(_54_v4, (_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool), (_91_v36).Cardinality())
				var _93_v38 _dafny.Map
				_ = _93_v38
				_93_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_92_v37, _49_v1)
				var _94_v39 _dafny.Map
				_ = _94_v39
				_94_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v1, (_60_v10).F10())
				var _95_v40 _dafny.Map
				_ = _95_v40
				_95_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("rwq"), _54_v4)
				var _96_v41 _dafny.Sequence
				_ = _96_v41
				_96_v41 = _dafny.SeqOf((func() bool {
					if (_94_v39).Contains((_60_v10).F9()) {
						return (_94_v39).Get((_60_v10).F9()).(bool)
					}
					return (_60_v10).F10()
				})(), (_60_v10).F10(), (func() bool {
					if (_94_v39).Contains((_60_v10).Fm4((_60_v10).F9(), (_60_v10).F9(), _95_v40, (_60_v10).F9(), _53_globalState)) {
						return (_94_v39).Get((_60_v10).Fm4((_60_v10).F9(), (_60_v10).F9(), _95_v40, (_60_v10).F9(), _53_globalState)).(bool)
					}
					return !(true)
				})(), (_60_v10).F10())
				(_60_v10).M1((func() _dafny.Int {
					if (_93_v38).Contains(_92_v37) {
						return (_93_v38).Get(_92_v37).(_dafny.Int)
					}
					return (_60_v10).F9()
				})(), (_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool), (_96_v41).Select((Companion_Default___.SafeIndex((_60_v10).F9(), _dafny.IntOfUint32((_96_v41).Cardinality()))).Uint32()).(bool), _53_globalState)
				var _97_v42 _dafny.Sequence
				_ = _97_v42
				_97_v42 = _dafny.SeqOf((_60_v10).F9())
				var _98_v43 T1
				_ = _98_v43
				var _nw8 *C1 = New_C1_()
				_ = _nw8
				_nw8.Ctor__((_60_v10).F10(), _48_v0, _54_v4, _59_v9)
				_98_v43 = _nw8
				var _99_v44 _dafny.Map
				_ = _99_v44
				_99_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_98_v43, (_60_v10).F10())
				var _100_v45 _dafny.Sequence
				_ = _100_v45
				_100_v45 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_96_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(-771), _dafny.IntOfUint32((_96_v41).Cardinality()))).Uint32(), !((_98_v43).F6())))
				var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
				_ = _index6
				var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
				_ = _index7
				var _rhs2 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-143))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg19 _dafny.Int) interface{} {
						return coer19(arg19)
					}
				}(func(_101_i4 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('e')
				})), _48_v0)
				_ = _rhs2
				var _rhs3 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_97_v42, _dafny.SeqOf((_60_v10).F9(), (_99_v44).Cardinality(), _49_v1, (_dafny.Zero).Minus((_60_v10).F9()))), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_96_v41).Cardinality()), (_60_v10).F9()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_97_v42, _dafny.SeqOf((_60_v10).F9(), (_99_v44).Cardinality(), _49_v1, (_dafny.Zero).Minus((_60_v10).F9())))).Cardinality()))).Uint32(), (_dafny.Zero).Minus(_dafny.IntOfInt64(-259)))
				_ = _rhs3
				var _rhs4 bool = (_54_v4) || ((_60_v10).F10())
				_ = _rhs4
				var _rhs5 bool = !_dafny.Companion_Sequence_.Contains(_100_v45, _dafny.Companion_Sequence_.Concatenate(_96_v41, _96_v41))
				_ = _rhs5
				var _rhs6 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_96_v41, _dafny.SeqOf((_98_v43).F6(), (_60_v10).F10())), _dafny.Companion_Sequence_.Concatenate(_96_v41, _96_v41))
				_ = _rhs6
				var _lhs3 _dafny.Array = _56_v6
				_ = _lhs3
				var _lhs4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
				_ = _lhs4
				var _lhs5 _dafny.Array = _56_v6
				_ = _lhs5
				var _lhs6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(609), _dafny.ArrayLen((_56_v6), 0))
				_ = _lhs6
				_48_v0 = _rhs2
				_97_v42 = _rhs3
				(_lhs3).ArraySet1(_rhs4, (_lhs4).Int())
				(_lhs5).ArraySet1(_rhs5, (_lhs6).Int())
				_96_v41 = _rhs6
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_49_v1 = _49_v1
	var _102_v46 _dafny.Map
	_ = _102_v46
	_102_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _48_v0)
	_102_v46 = _102_v46
	var _103_v47 _dafny.Map
	_ = _103_v47
	_103_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-434), (_60_v10).F9()), _54_v4)
	var _104_v48 *C0
	_ = _104_v48
	var _nw9 *C0 = New_C0_()
	_ = _nw9
	_nw9.Ctor__(_103_v47)
	_104_v48 = _nw9
	var _105_v49 _dafny.Array
	_ = _105_v49
	var _len0_1 _dafny.Int = _dafny.IntOfInt64(12)
	_ = _len0_1
	var _nw10 _dafny.Array
	_ = _nw10
	if _len0_1.Cmp(_dafny.Zero) == 0 {
		_nw10 = _dafny.NewArray(_len0_1)
	} else {
		var _init1 func(_dafny.Int) _dafny.Int = (func(_106_v1 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_107_i5 _dafny.Int) _dafny.Int {
				return (_107_i5).Plus(_106_v1)
			}
		})(_49_v1)
		_ = _init1
		var _element0_1 = _init1(_dafny.Zero)
		_ = _element0_1
		_nw10 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
		(_nw10).ArraySet1(_element0_1, 0)
		var _nativeLen0_1 = (_len0_1).Int()
		_ = _nativeLen0_1
		for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
			(_nw10).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
		}
	}
	_105_v49 = _nw10
	var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
	_ = _index8
	(_105_v49).ArraySet1((_60_v10).F9(), (_index8).Int())
	var _108_v51 _dafny.Array
	_ = _108_v51
	var _len0_2 _dafny.Int = _dafny.IntOfInt64(20)
	_ = _len0_2
	var _nw11 _dafny.Array
	_ = _nw11
	if _len0_2.Cmp(_dafny.Zero) == 0 {
		_nw11 = _dafny.NewArray(_len0_2)
	} else {
		var _init2 func(_dafny.Int) _dafny.Map = (func(_109_v10 *C5, _110_v4 bool) func(_dafny.Int) _dafny.Map {
			return func(_111_i6 _dafny.Int) _dafny.Map {
				return func() _dafny.Map {
					var _coll14 = _dafny.NewMapBuilder()
					_ = _coll14
					for _iter14 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D8_.Create_DC24_((_109_v10).F9(), _110_v4, _dafny.SeqOf((_109_v10).F10()), _dafny.IntOfInt64(959), (_109_v10).F9()))).Elements()); ; {
						_compr_14, _ok14 := _iter14()
						if !_ok14 {
							break
						}
						var _112_v50 D8
						_112_v50 = interface{}(_compr_14).(D8)
						if (_dafny.MultiSetOf(Companion_D8_.Create_DC24_((_109_v10).F9(), _110_v4, _dafny.SeqOf((_109_v10).F10()), _dafny.IntOfInt64(959), (_109_v10).F9()))).Contains(_112_v50) {
							_coll14.Add(_112_v50, _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))
						}
					}
					return _coll14.ToMap()
				}()
			}
		})(_60_v10, _54_v4)
		_ = _init2
		var _element0_2 = _init2(_dafny.Zero)
		_ = _element0_2
		_nw11 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
		(_nw11).ArraySet1(_element0_2, 0)
		var _nativeLen0_2 = (_len0_2).Int()
		_ = _nativeLen0_2
		for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
			(_nw11).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
		}
	}
	_108_v51 = _nw11
	var _113_v52 _dafny.Sequence
	_ = _113_v52
	_113_v52 = _dafny.SeqOf((_60_v10).F10(), (_60_v10).F10(), (_60_v10).Fm1((_60_v10).F10(), _dafny.IntOfInt64(-253), (_60_v10).F9(), _53_globalState))
	var _114_v53 D8
	_ = _114_v53
	_114_v53 = Companion_D8_.Create_DC24_((_dafny.MultiSetFromSeq(_113_v52)).Cardinality(), (_60_v10).F10(), _113_v52, (_60_v10).F9(), _49_v1)
	var _115_v54 _dafny.Map
	_ = _115_v54
	_115_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_49_v1, (_60_v10).F9())
	var _116_v55 _dafny.Map
	_ = _116_v55
	_116_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_114_v53, _dafny.SeqOf(_50_v2, _50_v2, Companion_Default___.Fm16(_54_v4, _115_v54, _48_v0, _53_globalState), _50_v2))
	var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_108_v51), 0))
	_ = _index9
	(_108_v51).ArraySet1(_116_v55, (_index9).Int())
	var _117_v57 _dafny.Sequence
	_ = _117_v57
	_117_v57 = _dafny.SeqOf(_114_v53, _114_v53)
	var _118_v58 _dafny.Sequence
	_ = _118_v58
	_118_v58 = _dafny.SeqOf(func() _dafny.Map {
		var _coll15 = _dafny.NewMapBuilder()
		_ = _coll15
		for _iter15 := _dafny.Iterate((_117_v57).Elements()); ; {
			_compr_15, _ok15 := _iter15()
			if !_ok15 {
				break
			}
			var _119_v56 D8
			_119_v56 = interface{}(_compr_15).(D8)
			if _dafny.Companion_Sequence_.Contains(_117_v57, _119_v56) {
				_coll15.Add(_119_v56, _48_v0)
			}
		}
		return _coll15.ToMap()
	}())
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
	_ = _index10
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_108_v51), 0))
	_ = _index11
	var _rhs7 *C0 = _104_v48
	_ = _rhs7
	var _rhs8 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_60_v10).F9()), (_60_v10).F9())
	_ = _rhs8
	var _rhs9 _dafny.Map = ((_118_v58).Select((Companion_Default___.SafeIndex(_49_v1, _dafny.IntOfUint32((_118_v58).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_116_v55)
	_ = _rhs9
	var _rhs10 _dafny.Int = (_60_v10).F9()
	_ = _rhs10
	var _rhs11 bool = (func() bool {
		if (_60_v10).F10() {
			return _54_v4
		}
		return _54_v4
	})()
	_ = _rhs11
	var _lhs7 _dafny.Array = _105_v49
	_ = _lhs7
	var _lhs8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
	_ = _lhs8
	var _lhs9 _dafny.Array = _108_v51
	_ = _lhs9
	var _lhs10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_108_v51), 0))
	_ = _lhs10
	var _lhs11 *GlobalState = _53_globalState
	_ = _lhs11
	_104_v48 = _rhs7
	(_lhs7).ArraySet1(_rhs8, (_lhs8).Int())
	(_lhs9).ArraySet1(_rhs9, (_lhs10).Int())
	_lhs11.F0 = _rhs10
	_54_v4 = _rhs11
	var _120_v60 _dafny.Set
	_ = _120_v60
	_120_v60 = _dafny.SetOf((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int))
	var _source2 D8 = Companion_Default___.Fm36((_115_v54).Merge(func() _dafny.Map {
		var _coll16 = _dafny.NewMapBuilder()
		_ = _coll16
		for _iter16 := _dafny.Iterate((_115_v54).Keys().Elements()); ; {
			_compr_16, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _121_v59 _dafny.Int
			_121_v59 = interface{}(_compr_16).(_dafny.Int)
			if (_115_v54).Contains(_121_v59) {
				_coll16.Add((_121_v59).Times(_dafny.IntOfInt64(-411)), (_60_v10).F9())
			}
		}
		return _coll16.ToMap()
	}()), _54_v4, _120_v60, _53_globalState)
	_ = _source2
	if _source2.Is_DC24() {
		var _122___mcc_h0 _dafny.Int = _source2.Get_().(D8_DC24).Cf34
		_ = _122___mcc_h0
		var _123___mcc_h1 bool = _source2.Get_().(D8_DC24).Cf35
		_ = _123___mcc_h1
		var _124___mcc_h2 _dafny.Sequence = _source2.Get_().(D8_DC24).Cf36
		_ = _124___mcc_h2
		var _125___mcc_h3 _dafny.Int = _source2.Get_().(D8_DC24).Cf37
		_ = _125___mcc_h3
		var _126___mcc_h4 _dafny.Int = _source2.Get_().(D8_DC24).Cf38
		_ = _126___mcc_h4
		var _127_cf38 _dafny.Int = _126___mcc_h4
		_ = _127_cf38
		var _128_cf37 _dafny.Int = _125___mcc_h3
		_ = _128_cf37
		var _129_cf36 _dafny.Sequence = _124___mcc_h2
		_ = _129_cf36
		var _130_cf35 bool = _123___mcc_h1
		_ = _130_cf35
		var _131_cf34 _dafny.Int = _122___mcc_h0
		_ = _131_cf34
		var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
		_ = _index12
		(_105_v49).ArraySet1(((_dafny.Zero).Minus((_127_cf38).Times(_128_cf37))).Minus((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int)), (_index12).Int())
		if (_60_v10).F10() {
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _index13
			(_105_v49).ArraySet1(_dafny.IntOfInt64(59), (_index13).Int())
			var _132_v61 _dafny.Map
			_ = _132_v61
			_132_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F9(), _130_cf35)
			_54_v4 = (func() bool {
				if (_132_v61).Contains(_128_cf37) {
					return (_132_v61).Get(_128_cf37).(bool)
				}
				return (_60_v10).F10()
			})()
			_127_cf38 = (_60_v10).F9()
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_56_v6), 0))
			_ = _index14
			(_56_v6).ArraySet1(_54_v4, (_index14).Int())
			var _133_v62 _dafny.Map
			_ = _133_v62
			_133_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v0, !((_60_v10).F10()))
			var _134_v63 _dafny.Map
			_ = _134_v63
			_134_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_130_cf35, _54_v4)
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(621), _dafny.ArrayLen((_56_v6), 0))
			_ = _index15
			(_56_v6).ArraySet1((_60_v10).Fm1(!((_60_v10).F10()) || (_54_v4), Companion_Default___.SafeDivisionInt((_60_v10).Fm4((_60_v10).F9(), (_60_v10).F9(), _133_v62, (_60_v10).F9(), _53_globalState), ((_134_v63).Update((_60_v10).Fm2((_60_v10).F10(), _53_globalState), (_60_v10).F10())).Cardinality()), (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), _53_globalState), (_index15).Int())
			_48_v0 = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(187))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_135_v2 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_136_i7 _dafny.Int) _dafny.CodePoint {
					return _135_v2
				}
			})(_50_v2)))
		} else {
			(_60_v10).M2(_50_v2, _53_globalState)
			_54_v4 = !(_130_cf35)
			_48_v0 = _48_v0
			var _137_v64 _dafny.Array
			_ = _137_v64
			var _nw12 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw12
			_137_v64 = _nw12
			var _138_v65 *C2
			_ = _138_v65
			var _nw13 *C2 = New_C2_()
			_ = _nw13
			_nw13.Ctor__((_59_v9))
			_138_v65 = _nw13
			var _139_v66 D9
			_ = _139_v66
			_139_v66 = Companion_D9_.Create_DC26_(_138_v65)
			var _140_v67 _dafny.MultiSet
			_ = _140_v67
			_140_v67 = _dafny.MultiSetOf((_60_v10).F9())
			var _141_v68 _dafny.Set
			_ = _141_v68
			_141_v68 = _dafny.SetOf((_60_v10).F10(), (_60_v10).F10())
			var _rhs12 bool = _dafny.Companion_Sequence_.IsPrefixOf(_129_cf36, Companion_Default___.Fm23((func() _dafny.Int {
				if (_140_v67).Contains((_dafny.Zero).Minus((_141_v68).Cardinality())) {
					return (_140_v67).Multiplicity((_dafny.Zero).Minus((_141_v68).Cardinality()))
				}
				return _131_cf34
			})(), _53_globalState))
			_ = _rhs12
			var _rhs13 _dafny.Array = _137_v64
			_ = _rhs13
			var _rhs14 _dafny.CodePoint = _50_v2
			_ = _rhs14
			var _rhs15 D9 = _139_v66
			_ = _rhs15
			_130_cf35 = _rhs12
			_137_v64 = _rhs13
			_50_v2 = _rhs14
			_139_v66 = _rhs15
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_56_v6), 0))
			_ = _index16
			(_56_v6).ArraySet1(_54_v4, (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_56_v6), 0))
			_ = _index17
			(_56_v6).ArraySet1((_60_v10).F10(), (_index17).Int())
		}
		var _142_v69 _dafny.Map
		_ = _142_v69
		_142_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v4, _dafny.SetOf(_131_cf34))
		var _143_v70 _dafny.Set
		_ = _143_v70
		_143_v70 = _dafny.SetOf((func() _dafny.Set {
			if (_142_v69).Contains(_54_v4) {
				return (_142_v69).Get(_54_v4).(_dafny.Set)
			}
			return _120_v60
		})(), Companion_Default___.Fm24(_dafny.IntOfUint32((_48_v0).Cardinality()), true, _49_v1, _53_globalState))
		_143_v70 = func() _dafny.Set {
			var _coll17 = _dafny.NewBuilder()
			_ = _coll17
			for _iter17 := _dafny.Iterate((_143_v70).Elements()); ; {
				_compr_17, _ok17 := _iter17()
				if !_ok17 {
					break
				}
				var _144_v71 _dafny.Set
				_144_v71 = interface{}(_compr_17).(_dafny.Set)
				if (_143_v70).Contains(_144_v71) {
					_coll17.Add(_144_v71)
				}
			}
			return _coll17.ToSet()
		}()
		var _145_v72 _dafny.Array
		_ = _145_v72
		var _nw14 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(24))
		_ = _nw14
		_145_v72 = _nw14
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_145_v72), 0))
		_ = _index18
		(_145_v72).ArraySet1(_51_v3, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(99), _dafny.ArrayLen((_145_v72), 0))
		_ = _index19
		(_145_v72).ArraySet1(_51_v3, (_index19).Int())
	} else if _source2.Is_DC25() {
		var _146___mcc_h5 bool = _source2.Get_().(D8_DC25).Cf39
		_ = _146___mcc_h5
		var _147___mcc_h6 bool = _source2.Get_().(D8_DC25).Cf40
		_ = _147___mcc_h6
		var _148_cf40 bool = _147___mcc_h6
		_ = _148_cf40
		var _149_cf39 bool = _146___mcc_h5
		_ = _149_cf39
		var _150_v73 _dafny.Map
		_ = _150_v73
		_150_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F10(), _105_v49)
		_150_v73 = _150_v73
		var _rhs16 _dafny.Int = (_60_v10).F9()
		_ = _rhs16
		var _rhs17 bool = _148_cf40
		_ = _rhs17
		var _rhs18 bool = !((_60_v10).F10()) || (((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F10(), _105_v49)).Update(false, _105_v49)).Contains(_148_cf40))
		_ = _rhs18
		var _rhs19 _dafny.Sequence = _48_v0
		_ = _rhs19
		var _rhs20 bool = true
		_ = _rhs20
		var _lhs12 *GlobalState = _53_globalState
		_ = _lhs12
		_lhs12.F0 = _rhs16
		_54_v4 = _rhs17
		_149_cf39 = _rhs18
		_48_v0 = _rhs19
		_149_cf39 = _rhs20
		_50_v2 = _dafny.CodePoint('d')
		(_53_globalState).F0 = _49_v1
	} else {
		var _151___mcc_h7 _dafny.Map = _source2.Get_().(D8_DC23).Cf33
		_ = _151___mcc_h7
		var _152_cf33 _dafny.Map = _151___mcc_h7
		_ = _152_cf33
		if (_60_v10).F10() {
			(_53_globalState).F0 = (_60_v10).F9()
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_56_v6), 0))
			_ = _index20
			(_56_v6).ArraySet1(!((_60_v10).F10()), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_56_v6), 0))
			_ = _index21
			(_56_v6).ArraySet1(!((_60_v10).F10()), (_index21).Int())
			var _153_v74 _dafny.Map
			_ = _153_v74
			_153_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_54_v4, (_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(75), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool))
			var _154_v75 _dafny.Map
			_ = _154_v75
			_154_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_54_v4, (_60_v10).F10()), _153_v74)
			var _155_v77 _dafny.Map
			_ = _155_v77
			_155_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_54_v4, !(_54_v4), (_60_v10).Fm1((_60_v10).F10(), _dafny.IntOfUint32((_48_v0).Cardinality()), _dafny.IntOfUint32((_113_v52).Cardinality()), _53_globalState), (_60_v10).F10()), (_153_v74).Cardinality())
			var _156_v79 _dafny.Set
			_ = _156_v79
			_156_v79 = _dafny.SetOf(_59_v9)
			_154_v75 = (_154_v75).Merge((func() _dafny.Map {
				var _coll18 = _dafny.NewMapBuilder()
				_ = _coll18
				for _iter18 := _dafny.Iterate((_155_v77).Keys().Elements()); ; {
					_compr_18, _ok18 := _iter18()
					if !_ok18 {
						break
					}
					var _157_v76 _dafny.MultiSet
					_157_v76 = interface{}(_compr_18).(_dafny.MultiSet)
					if (_155_v77).Contains(_157_v76) {
						_coll18.Add(_157_v76, _153_v74)
					}
				}
				return _coll18.ToMap()
			}()).Merge(func() _dafny.Map {
				var _coll19 = _dafny.NewMapBuilder()
				_ = _coll19
				for _iter19 := _dafny.Iterate((_156_v79).Elements()); ; {
					_compr_19, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _158_v78 _dafny.MultiSet
					_158_v78 = interface{}(_compr_19).(_dafny.MultiSet)
					if (_156_v79).Contains(_158_v78) {
						_coll19.Add(_158_v78, _153_v74)
					}
				}
				return _coll19.ToMap()
			}()))
			(_53_globalState).F0 = (_60_v10).F9()
			var _159_v80 _dafny.Array
			_ = _159_v80
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_3
			var _nw15 _dafny.Array
			_ = _nw15
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw15 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) _dafny.Map = (func(_160_v10 *C5) func(_dafny.Int) _dafny.Map {
					return func(_161_i8 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_160_v10).F10(), _dafny.IntOfInt64(-120))
					}
				})(_60_v10)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw15 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw15).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw15).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_159_v80 = _nw15
			var _162_v81 bool
			_ = _162_v81
			var _163_v82 _dafny.Int
			_ = _163_v82
			var _164_v83 bool
			_ = _164_v83
			var _out0 bool
			_ = _out0
			var _out1 _dafny.Int
			_ = _out1
			var _out2 bool
			_ = _out2
			_out0, _out1, _out2 = (_60_v10).M5(_dafny.IntOfInt64(-675), _59_v9, _159_v80, _54_v4, _53_globalState)
			_162_v81 = _out0
			_163_v82 = _out1
			_164_v83 = _out2
		} else {
			var _165_v84 _dafny.Sequence
			_ = _165_v84
			_165_v84 = _dafny.SeqOf(_49_v1)
			(_60_v10).M4((_60_v10).Fm2((_60_v10).F10(), _53_globalState), _165_v84, _53_globalState)
			var _166_v85 _dafny.Array
			_ = _166_v85
			var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(12))
			_ = _nw16
			_166_v85 = _nw16
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_166_v85), 0))
			_ = _index22
			(_166_v85).ArraySet1(_56_v6, (_index22).Int())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(677), _dafny.ArrayLen((_166_v85), 0))
			_ = _index23
			(_166_v85).ArraySet1(_56_v6, (_index23).Int())
			var _167_v86 _dafny.Sequence
			_ = _167_v86
			_167_v86 = _dafny.SeqOf(_104_v48, _104_v48)
			var _168_v87 D2
			_ = _168_v87
			_168_v87 = Companion_D2_.Create_DC7_(_50_v2, _115_v54, (_60_v10).F9())
			var _169_v88 _dafny.Map
			_ = _169_v88
			_169_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v0, _54_v4)
			var _170_v89 _dafny.Array
			_ = _170_v89
			var _nwElement0_2 *C0 = _104_v48
			_ = _nwElement0_2
			var _nw17 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(28))
			_ = _nw17
			(_nw17).ArraySet1(_nwElement0_2, 0)
			(_nw17).ArraySet1(_104_v48, 1)
			(_nw17).ArraySet1(_104_v48, 2)
			(_nw17).ArraySet1(_104_v48, 3)
			(_nw17).ArraySet1(_104_v48, 4)
			(_nw17).ArraySet1(_104_v48, 5)
			(_nw17).ArraySet1(_104_v48, 6)
			(_nw17).ArraySet1(_104_v48, 7)
			(_nw17).ArraySet1((_167_v86).Select((Companion_Default___.SafeIndex((_60_v10).Fm4(_49_v1, (_168_v87).Dtor_cf8(), _169_v88, _49_v1, _53_globalState), _dafny.IntOfUint32((_167_v86).Cardinality()))).Uint32()).(*C0), 8)
			(_nw17).ArraySet1(_104_v48, 9)
			(_nw17).ArraySet1(_104_v48, 10)
			(_nw17).ArraySet1(_104_v48, 11)
			(_nw17).ArraySet1(_104_v48, 12)
			(_nw17).ArraySet1(_104_v48, 13)
			(_nw17).ArraySet1(_104_v48, 14)
			(_nw17).ArraySet1(_104_v48, 15)
			(_nw17).ArraySet1(_104_v48, 16)
			(_nw17).ArraySet1(_104_v48, 17)
			(_nw17).ArraySet1(_104_v48, 18)
			(_nw17).ArraySet1(_104_v48, 19)
			(_nw17).ArraySet1(_104_v48, 20)
			(_nw17).ArraySet1(_104_v48, 21)
			(_nw17).ArraySet1(_104_v48, 22)
			(_nw17).ArraySet1(_104_v48, 23)
			(_nw17).ArraySet1(_104_v48, 24)
			(_nw17).ArraySet1(_104_v48, 25)
			(_nw17).ArraySet1(_104_v48, 26)
			(_nw17).ArraySet1(_104_v48, 27)
			_170_v89 = _nw17
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_170_v89), 0))
			_ = _index24
			(_170_v89).ArraySet1(_104_v48, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_170_v89), 0))
			_ = _index25
			var _rhs21 bool = !((_dafny.IntOfInt64(381)).Cmp((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int)) < 0)
			_ = _rhs21
			var _rhs22 _dafny.Int = (_60_v10).F9()
			_ = _rhs22
			var _rhs23 *C0 = _104_v48
			_ = _rhs23
			var _lhs13 *GlobalState = _53_globalState
			_ = _lhs13
			var _lhs14 _dafny.Array = _170_v89
			_ = _lhs14
			var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(708), _dafny.ArrayLen((_170_v89), 0))
			_ = _lhs15
			_54_v4 = _rhs21
			_lhs13.F0 = _rhs22
			(_lhs14).ArraySet1(_rhs23, (_lhs15).Int())
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _index26
			(_105_v49).ArraySet1((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_index26).Int())
			var _171_v90 T0
			_ = _171_v90
			var _nw18 *C2 = New_C2_()
			_ = _nw18
			_nw18.Ctor__(_59_v9)
			_171_v90 = _nw18
			var _rhs24 bool = (_60_v10).F10()
			_ = _rhs24
			var _rhs25 T0 = _171_v90
			_ = _rhs25
			_54_v4 = _rhs24
			_171_v90 = _rhs25
		}
		var _172_v91 _dafny.Sequence
		_ = _172_v91
		_172_v91 = _dafny.SeqOf(_49_v1)
		var _173_v92 D15
		_ = _173_v92
		_173_v92 = Companion_D15_.Create_DC42_(_dafny.MultiSetFromSeq(_172_v91))
		var _174_v93 _dafny.Sequence
		_ = _174_v93
		_174_v93 = _dafny.SeqOf(Companion_D15_.Create_DC42_(_dafny.MultiSetFromSeq(_172_v91)), _173_v92)
		var _175_v94 D21
		_ = _175_v94
		_175_v94 = Companion_D21_.Create_DC54_(_174_v93)
		var _176_v95 _dafny.MultiSet
		_ = _176_v95
		_176_v95 = _dafny.MultiSetOf((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_60_v10).F9())
		if (_49_v1).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_175_v94).Dtor_cf84(), _dafny.SeqOf(Companion_D15_.Create_DC42_(_176_v95), Companion_D15_.Create_DC42_(_176_v95), _173_v92))).Cardinality())) > 0 {
			var _177_v96 *C2
			_ = _177_v96
			var _nw19 *C2 = New_C2_()
			_ = _nw19
			_nw19.Ctor__(_59_v9)
			_177_v96 = _nw19
			var _178_v97 _dafny.Map
			_ = _178_v97
			_178_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_177_v96, (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int))
			var _179_v98 _dafny.Sequence
			_ = _179_v98
			_179_v98 = _dafny.SeqOf(_56_v6)
			var _180_v99 _dafny.Map
			_ = _180_v99
			_180_v99 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_179_v98, (_60_v10).F10())
			(_60_v10).M1((_178_v97).Cardinality(), (func() bool {
				if (_180_v99).Contains(_dafny.SeqOf(_56_v6, _56_v6)) {
					return (_180_v99).Get(_dafny.SeqOf(_56_v6, _56_v6)).(bool)
				}
				return (_60_v10).F10()
			})(), (_60_v10).F10(), _53_globalState)
			var _181_v100 _dafny.Array
			_ = _181_v100
			var _nw20 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(11))
			_ = _nw20
			_181_v100 = _nw20
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_181_v100), 0))
			_ = _index27
			(_181_v100).ArraySet1(_60_v10, (_index27).Int())
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen((_181_v100), 0))
			_ = _index28
			(_181_v100).ArraySet1(_60_v10, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _index29
			(_105_v49).ArraySet1((_dafny.IntOfInt64(569)).Plus(_dafny.IntOfUint32((_113_v52).Cardinality())), (_index29).Int())
			var _182_v101 _dafny.Map
			_ = _182_v101
			_182_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F10(), (_60_v10).F10())
			var _183_v102 D13
			_ = _183_v102
			_183_v102 = Companion_D13_.Create_DC39_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_50_v2, _54_v4)).Merge(_55_v5), _56_v6, _113_v52, (_182_v101).Cardinality(), _dafny.Companion_Sequence_.Concatenate(_172_v91, _172_v91))
			_183_v102 = Companion_D13_.Create_DC39_(Companion_Default___.Fm37(_50_v2, _53_globalState), _56_v6, _113_v52, ((_60_v10).F9()).Minus(_dafny.IntOfInt64(134)), (func() _dafny.Sequence {
				if true {
					return _172_v91
				}
				return Companion_Default___.Fm18(_53_globalState)
			})())
			_54_v4 = (_dafny.MultiSetFromSeq(_113_v52)).Contains(!((_60_v10).F10()) || (_54_v4))
		} else {
			_50_v2 = _50_v2
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _index30
			(_105_v49).ArraySet1(_49_v1, (_index30).Int())
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_56_v6), 0))
			_ = _index31
			(_56_v6).ArraySet1((_60_v10).F10(), (_index31).Int())
			var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_56_v6), 0))
			_ = _index32
			(_56_v6).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_48_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("x"), _dafny.UnicodeSeqOfUtf8Bytes("btussn"))), (_index32).Int())
			var _184_v103 *C6
			_ = _184_v103
			var _nw21 *C6 = New_C6_()
			_ = _nw21
			_nw21.Ctor__(_120_v60, _50_v2)
			_184_v103 = _nw21
			var _185_v104 _dafny.Map
			_ = _185_v104
			_185_v104 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_184_v103, (_56_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_56_v6), 0))).Int()).(bool))
			_185_v104 = (_185_v104).Update((func() *C6 {
				if _54_v4 {
					return _184_v103
				}
				return _184_v103
			})(), (_60_v10).F10())
			var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _index33
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_56_v6), 0))
			_ = _index34
			var _rhs26 _dafny.Int = (_60_v10).F9()
			_ = _rhs26
			var _rhs27 bool = (_60_v10).F10()
			_ = _rhs27
			var _lhs16 _dafny.Array = _105_v49
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _lhs17
			var _lhs18 _dafny.Array = _56_v6
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(8), _dafny.ArrayLen((_56_v6), 0))
			_ = _lhs19
			(_lhs16).ArraySet1(_rhs26, (_lhs17).Int())
			(_lhs18).ArraySet1(_rhs27, (_lhs19).Int())
		}
		var _186_v105 *C3
		_ = _186_v105
		var _nw22 *C3 = New_C3_()
		_ = _nw22
		_nw22.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf(_113_v52, _113_v52), _59_v9)
		_186_v105 = _nw22
		_54_v4 = (_60_v10).F10()
	}
	var _187_v106 _dafny.Array
	_ = _187_v106
	var _len0_4 _dafny.Int = _dafny.IntOfInt64(3)
	_ = _len0_4
	var _nw23 _dafny.Array
	_ = _nw23
	if _len0_4.Cmp(_dafny.Zero) == 0 {
		_nw23 = _dafny.NewArray(_len0_4)
	} else {
		var _init4 func(_dafny.Int) D6 = (func(_188_v52 _dafny.Sequence) func(_dafny.Int) D6 {
			return func(_189_i9 _dafny.Int) D6 {
				return Companion_D6_.Create_DC17_(_188_v52)
			}
		})(_113_v52)
		_ = _init4
		var _element0_4 = _init4(_dafny.Zero)
		_ = _element0_4
		_nw23 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
		(_nw23).ArraySet1(_element0_4, 0)
		var _nativeLen0_4 = (_len0_4).Int()
		_ = _nativeLen0_4
		for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
			(_nw23).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
		}
	}
	_187_v106 = _nw23
	var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_187_v106), 0))
	_ = _index35
	(_187_v106).ArraySet1(Companion_D6_.Create_DC17_(_113_v52), (_index35).Int())
	var _190_v107 D6
	_ = _190_v107
	_190_v107 = Companion_D6_.Create_DC17_(_dafny.SeqOf(false))
	var _pat_let_tv0 = _113_v52
	_ = _pat_let_tv0
	var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(622), _dafny.ArrayLen((_187_v106), 0))
	_ = _index36
	(_187_v106).ArraySet1(func(_pat_let0_0 D6) D6 {
		return func(_191_dt__update__tmp_h0 D6) D6 {
			return func(_pat_let1_0 _dafny.Sequence) D6 {
				return func(_192_dt__update_hcf24_h0 _dafny.Sequence) D6 {
					return Companion_D6_.Create_DC17_(_192_dt__update_hcf24_h0)
				}(_pat_let1_0)
			}(_pat_let_tv0)
		}(_pat_let0_0)
	}(_190_v107), (_index36).Int())
	_54_v4 = _54_v4
	var _193_v108 *C3
	_ = _193_v108
	var _nw24 *C3 = New_C3_()
	_ = _nw24
	_nw24.Ctor__((_60_v10).F10(), _59_v9)
	_193_v108 = _nw24
	var _194_v109 _dafny.Set
	_ = _194_v109
	_194_v109 = _dafny.SetOf(_193_v108, _193_v108, _193_v108)
	_194_v109 = _194_v109
	var _195_v110 _dafny.Set
	_ = _195_v110
	_195_v110 = _dafny.SetOf((_60_v10).F10(), (_60_v10).Fm1(_54_v4, _49_v1, (_120_v60).Cardinality(), _53_globalState), _54_v4)
	var _196_v111 _dafny.Map
	_ = _196_v111
	_196_v111 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_48_v0, false)
	var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
	_ = _index37
	var _rhs28 _dafny.Int = ((_60_v10).Fm4(_dafny.IntOfUint32((_48_v0).Cardinality()), _49_v1, _196_v111, _dafny.IntOfUint32((_48_v0).Cardinality()), _53_globalState)).Plus(_dafny.IntOfInt64(-538))
	_ = _rhs28
	var _rhs29 bool = (_60_v10).F10()
	_ = _rhs29
	var _rhs30 _dafny.Set = _dafny.SetOf(!((_60_v10).F10()))
	_ = _rhs30
	var _rhs31 bool = (_113_v52).Select((Companion_Default___.SafeIndex((_60_v10).F9(), _dafny.IntOfUint32((_113_v52).Cardinality()))).Uint32()).(bool)
	_ = _rhs31
	var _lhs20 _dafny.Array = _105_v49
	_ = _lhs20
	var _lhs21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
	_ = _lhs21
	(_lhs20).ArraySet1(_rhs28, (_lhs21).Int())
	_54_v4 = _rhs29
	_195_v110 = _rhs30
	_54_v4 = _rhs31
	var _197_v112 D17
	_ = _197_v112
	_197_v112 = Companion_D17_.Create_DC47_(_54_v4, (_60_v10).Fm4((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_60_v10).F10(), (_60_v10).Fm4((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_48_v0).Cardinality()), _196_v111, (_60_v10).F9(), _53_globalState))).Cardinality(), _196_v111, (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), _53_globalState), (_60_v10).F9())
	if (_197_v112).Dtor_cf74() {
		(_193_v108).M1(_49_v1, (_60_v10).F10(), (_60_v10).F10(), _53_globalState)
		var _198_v113 _dafny.Map
		_ = _198_v113
		_198_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wfd"), _195_v110)
		_198_v113 = (_198_v113).Update(_dafny.Companion_Sequence_.Concatenate(_48_v0, _48_v0), _195_v110)
		(_60_v10).M1((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_60_v10).F10(), true, _53_globalState)
		if (_113_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(776), _dafny.IntOfUint32((_113_v52).Cardinality()))).Uint32()).(bool) {
			_54_v4 = ((_54_v4) || (_54_v4)) || (((_115_v54).Update(_dafny.IntOfInt64(8), (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int))).Contains(_49_v1))
			_54_v4 = (func() bool {
				if (_60_v10).F10() {
					return (_60_v10).F10()
				}
				return (_54_v4) || ((_60_v10).F10())
			})()
			var _199_v114 _dafny.MultiSet
			_ = _199_v114
			_199_v114 = _dafny.MultiSetOf(_195_v110, _dafny.SetOf(_54_v4), _195_v110, _195_v110)
			var _200_v115 *C3
			_ = _200_v115
			var _nw25 *C3 = New_C3_()
			_ = _nw25
			_nw25.Ctor__((_60_v10).Fm1((_60_v10).F10(), _49_v1, (_199_v114).Cardinality(), _53_globalState), (_59_v9).Union(_59_v9))
			_200_v115 = _nw25
			var _rhs32 _dafny.Int = (_60_v10).F9()
			_ = _rhs32
			var _rhs33 _dafny.Int = Companion_Default___.SafeDivisionInt(_49_v1, (_dafny.Zero).Minus((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int)))
			_ = _rhs33
			var _rhs34 bool = true
			_ = _rhs34
			var _lhs22 *GlobalState = _53_globalState
			_ = _lhs22
			var _lhs23 *GlobalState = _53_globalState
			_ = _lhs23
			_lhs22.F0 = _rhs32
			_lhs23.F0 = _rhs33
			_54_v4 = _rhs34
			var _201_v116 _dafny.Array
			_ = _201_v116
			var _nw26 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw26
			_201_v116 = _nw26
			var _202_v117 T2
			_ = _202_v117
			var _nw27 *C5 = New_C5_()
			_ = _nw27
			_nw27.Ctor__((_60_v10).F9(), (_60_v10).F10(), _56_v6, (_60_v10).F10(), _54_v4, _59_v9)
			_202_v117 = _nw27
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_201_v116), 0))
			_ = _index38
			(_201_v116).ArraySet1(_202_v117, (_index38).Int())
			var _203_v118 _dafny.Array
			_ = _203_v118
			var _nw28 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
			_ = _nw28
			_203_v118 = _nw28
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _index39
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_201_v116), 0))
			_ = _index40
			var _rhs35 _dafny.Int = (_60_v10).F9()
			_ = _rhs35
			var _rhs36 T2 = _202_v117
			_ = _rhs36
			var _rhs37 _dafny.Array = _203_v118
			_ = _rhs37
			var _rhs38 _dafny.Sequence = _48_v0
			_ = _rhs38
			var _lhs24 _dafny.Array = _105_v49
			_ = _lhs24
			var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
			_ = _lhs25
			var _lhs26 _dafny.Array = _201_v116
			_ = _lhs26
			var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_201_v116), 0))
			_ = _lhs27
			(_lhs24).ArraySet1(_rhs35, (_lhs25).Int())
			(_lhs26).ArraySet1(_rhs36, (_lhs27).Int())
			_203_v118 = _rhs37
			_48_v0 = _rhs38
		} else {
			var _204_v119 _dafny.Array
			_ = _204_v119
			var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(20))
			_ = _nw29
			_204_v119 = _nw29
			var _205_v120 _dafny.Sequence
			_ = _205_v120
			_205_v120 = _dafny.SeqOf(_105_v49)
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_204_v119), 0))
			_ = _index41
			(_204_v119).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_205_v120, _dafny.SeqOf(_105_v49)), (_index41).Int())
			var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(154), _dafny.ArrayLen((_204_v119), 0))
			_ = _index42
			(_204_v119).ArraySet1(_dafny.SeqOf(_105_v49, _105_v49), (_index42).Int())
			var _206_v121 _dafny.MultiSet
			_ = _206_v121
			_206_v121 = _dafny.MultiSetOf(_49_v1)
			_206_v121 = _206_v121
			var _207_v122 *C6
			_ = _207_v122
			var _nw30 *C6 = New_C6_()
			_ = _nw30
			_nw30.Ctor__(_dafny.SetOf((_60_v10).F9()), _50_v2)
			_207_v122 = _nw30
			_54_v4 = (_60_v10).F10()
			(_53_globalState).F0 = (_60_v10).F9()
		}
		var _208_v123 _dafny.Map
		_ = _208_v123
		_208_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_54_v4) || (!((_60_v10).F10())), (false) == ((_60_v10).F10()))
		var _209_v124 D12
		_ = _209_v124
		_209_v124 = Companion_D12_.Create_DC35_(_54_v4)
		_208_v123 = (_208_v123).Update(!_dafny.Companion_Sequence_.Equal(_dafny.SeqOf(!((_209_v124).Dtor_cf55())), _113_v52), _54_v4)
	} else {
		var _210_v125 D3
		_ = _210_v125
		_210_v125 = Companion_D3_.Create_DC9_(_105_v49)
		var _211_v126 D5
		_ = _211_v126
		_211_v126 = Companion_D5_.Create_DC16_(_210_v125, _dafny.IntOfInt64(941), (_60_v10).F9(), _dafny.IntOfInt64(-528), _104_v48)
		var _212_v127 _dafny.Map
		_ = _212_v127
		_212_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_211_v126, _105_v49)
		_105_v49 = (func() _dafny.Array {
			if (_212_v127).Contains(_211_v126) {
				return (_212_v127).Get(_211_v126).(_dafny.Array)
			}
			return _105_v49
		})()
		var _213_v128 *C3
		_ = _213_v128
		var _nw31 *C3 = New_C3_()
		_ = _nw31
		_nw31.Ctor__((_60_v10).F10(), (_dafny.MultiSetOf(_54_v4, !((_60_v10).F10()), (_60_v10).F10(), (_193_v108).Fm3(_53_globalState))).Union(_59_v9))
		_213_v128 = _nw31
		var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_56_v6), 0))
		_ = _index43
		(_56_v6).ArraySet1(_54_v4, (_index43).Int())
		var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_56_v6), 0))
		_ = _index44
		(_56_v6).ArraySet1(_54_v4, (_index44).Int())
		var _214_v129 T0
		_ = _214_v129
		var _nw32 *C2 = New_C2_()
		_ = _nw32
		_nw32.Ctor__(_dafny.MultiSetFromSeq(_dafny.SeqOf(_54_v4)))
		_214_v129 = _nw32
		var _215_v130 _dafny.MultiSet
		_ = _215_v130
		_215_v130 = _dafny.MultiSetOf(_214_v129, _214_v129)
		var _216_v131 D22
		_ = _216_v131
		_216_v131 = Companion_D22_.Create_DC57_(_214_v129)
		var _217_v132 _dafny.Sequence
		_ = _217_v132
		_217_v132 = _dafny.SeqOf(_214_v129)
		var _218_v133 _dafny.Array
		_ = _218_v133
		var _len0_5 _dafny.Int = _dafny.IntOfInt64(10)
		_ = _len0_5
		var _nw33 _dafny.Array
		_ = _nw33
		if _len0_5.Cmp(_dafny.Zero) == 0 {
			_nw33 = _dafny.NewArray(_len0_5)
		} else {
			var _init5 func(_dafny.Int) bool = (func(_219_v6 _dafny.Array) func(_dafny.Int) bool {
				return func(_220_i10 _dafny.Int) bool {
					return (_219_v6).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(989), _dafny.ArrayLen((_219_v6), 0))).Int()).(bool)
				}
			})(_56_v6)
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
		_218_v133 = _nw33
		var _221_v134 *C4
		_ = _221_v134
		var _nw34 *C4 = New_C4_()
		_ = _nw34
		_nw34.Ctor__(_218_v133, (_60_v10).F10(), _54_v4, _59_v9)
		_221_v134 = _nw34
		var _222_v135 _dafny.Map
		_ = _222_v135
		_222_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_48_v0).Cardinality()), _221_v134)
		var _223_v136 _dafny.Map
		_ = _223_v136
		_223_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C4 {
			if (_222_v135).Contains(_49_v1) {
				return (_222_v135).Get(_49_v1).(*C4)
			}
			return _221_v134
		})(), _215_v130)
		var _224_v137 _dafny.Array
		_ = _224_v137
		var _nwElement0_3 _dafny.MultiSet = (_215_v130).Union(_dafny.MultiSetOf(_214_v129))
		_ = _nwElement0_3
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(23))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_3, 0)
		(_nw35).ArraySet1((func() _dafny.MultiSet {
			if _54_v4 {
				return _215_v130
			}
			return _dafny.MultiSetOf(_214_v129, _214_v129, (_216_v131).Dtor_cf88())
		})(), 1)
		(_nw35).ArraySet1(_215_v130, 2)
		(_nw35).ArraySet1((_215_v130).Intersection(_215_v130), 3)
		(_nw35).ArraySet1(_215_v130, 4)
		(_nw35).ArraySet1(_215_v130, 5)
		(_nw35).ArraySet1((_dafny.MultiSetFromSeq(_217_v132)).Update(_214_v129, Companion_Default___.Abs((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int))), 6)
		(_nw35).ArraySet1(_215_v130, 7)
		(_nw35).ArraySet1(_215_v130, 8)
		(_nw35).ArraySet1(_215_v130, 9)
		(_nw35).ArraySet1(_215_v130, 10)
		(_nw35).ArraySet1((_215_v130).Union(_dafny.MultiSetOf(_214_v129, _214_v129)), 11)
		(_nw35).ArraySet1(_dafny.MultiSetOf(_214_v129), 12)
		(_nw35).ArraySet1((_215_v130).Difference(_215_v130), 13)
		(_nw35).ArraySet1((_dafny.MultiSetFromSeq(_217_v132)).Union(_215_v130), 14)
		(_nw35).ArraySet1((func() _dafny.MultiSet {
			if (_223_v136).Contains(_221_v134) {
				return (_223_v136).Get(_221_v134).(_dafny.MultiSet)
			}
			return _215_v130
		})(), 15)
		(_nw35).ArraySet1((_215_v130).Difference(_215_v130), 16)
		(_nw35).ArraySet1((_dafny.MultiSetFromSeq(_dafny.SeqOf(_214_v129))).Intersection((_215_v130).Update(_214_v129, Companion_Default___.Abs((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int)))), 17)
		(_nw35).ArraySet1(_dafny.MultiSetOf(_214_v129), 18)
		(_nw35).ArraySet1((_215_v130).Intersection(_215_v130), 19)
		(_nw35).ArraySet1(_215_v130, 20)
		(_nw35).ArraySet1(_215_v130, 21)
		(_nw35).ArraySet1(_215_v130, 22)
		_224_v137 = _nw35
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_224_v137), 0))
		_ = _index45
		(_224_v137).ArraySet1(_dafny.MultiSetFromSeq(_217_v132), (_index45).Int())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_224_v137), 0))
		_ = _index46
		(_224_v137).ArraySet1(_215_v130, (_index46).Int())
		var _225_v138 _dafny.Map
		_ = _225_v138
		_225_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!_dafny.Companion_Sequence_.Equal(Companion_Default___.Fm11((_221_v134).Fm1(true, _dafny.IntOfInt64(357), _49_v1, _53_globalState), _53_globalState), _48_v0), (_221_v134).Fm6(_53_globalState))
		_225_v138 = (_225_v138).Update((_60_v10).F10(), _dafny.IntOfInt64(36))
	}
	(_53_globalState).F0 = _49_v1
	var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))
	_ = _index47
	(_105_v49).ArraySet1((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int), (_index47).Int())
	for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_56_v6), 0))); ; {
		_guard_loop_0, _ok20 := _iter20()
		if !_ok20 {
			break
		}
		var _226_i11 _dafny.Int
		_226_i11 = interface{}(_guard_loop_0).(_dafny.Int)
		if (true) && (((_226_i11).Sign() != -1) && ((_226_i11).Cmp(_dafny.ArrayLen((_56_v6), 0)) < 0)) {
			(_56_v6).ArraySet1((_60_v10).F10(), (_226_i11).Int())
		}
	}
	var _227_i12 _dafny.Int
	_ = _227_i12
	_227_i12 = _dafny.Zero
	{
		for _54_v4 {
			{
				if (_227_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_227_i12 = (_227_i12).Plus(_dafny.One)
				var _228_v139 _dafny.Sequence
				_ = _228_v139
				_228_v139 = _dafny.SeqOf(_dafny.IntOfUint32((_48_v0).Cardinality()))
				var _229_v140 _dafny.MultiSet
				_ = _229_v140
				_229_v140 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_228_v139, _dafny.SeqOf((_60_v10).F9()))).Cardinality()), (_60_v10).F9(), (_193_v108).Fm4(_49_v1, _49_v1, _196_v111, (_60_v10).F9(), _53_globalState), (_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int))
				var _230_v141 _dafny.Sequence
				_ = _230_v141
				_230_v141 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(279))).Uint32(), func(coer21 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg21 _dafny.Int) interface{} {
						return coer21(arg21)
					}
				}((func(_231_v10 *C5) func(_dafny.Int) _dafny.Int {
					return func(_232_i13 _dafny.Int) _dafny.Int {
						return (_231_v10).F9()
					}
				})(_60_v10))))
				_229_v140 = _dafny.MultiSetFromSeq((_230_v141).Select((Companion_Default___.SafeIndex((_dafny.IntOfUint32((_48_v0).Cardinality())).Plus((_105_v49).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(431), _dafny.ArrayLen((_105_v49), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_230_v141).Cardinality()))).Uint32()).(_dafny.Sequence))
				_49_v1 = (_60_v10).F9()
				var _233_v142 _dafny.Set
				_ = _233_v142
				_233_v142 = _dafny.SetOf(_56_v6, _56_v6, _56_v6)
				_54_v4 = ((_233_v142).Intersection(_dafny.SetOf(_56_v6, _56_v6))).IsDisjointFrom(_233_v142)
				var _234_v143 _dafny.Array
				_ = _234_v143
				var _nw36 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.One)
				_ = _nw36
				_234_v143 = _nw36
				_234_v143 = _234_v143
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	_dafny.Print(_48_v0.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_49_v1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_50_v2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_51_v3).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_53_globalState.F0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState).F1().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.Zero).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.One).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.IntOfInt64(2)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.IntOfInt64(3)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.IntOfInt64(4)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.IntOfInt64(5)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.IntOfInt64(6)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_53_globalState.F2).ArrayGet1CodePoint((_dafny.IntOfInt64(7)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_54_v4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_55_v5).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('w'), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_56_v6).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_57_v7).Equals(_dafny.SetOf(_dafny.CodePoint('o'), _dafny.CodePoint('w'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_58_v8, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("rwdkj"), _dafny.UnicodeSeqOfUtf8Bytes("rwwkj"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_59_v9).Equals(_dafny.MultiSetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v10).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_60_v10).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_61_i1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_102_v46).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.UnicodeSeqOfUtf8Bytes("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_103_v47).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-434), _dafny.IntOfInt64(143)), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_104_v48).F11()).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-434), _dafny.IntOfInt64(143)), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_105_v49).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.Zero).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.One).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(3), true, _dafny.SeqOf(true, true, true), _dafny.IntOfInt64(143), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('w')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_108_v51).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(143), true, _dafny.SeqOf(true), _dafny.IntOfInt64(959), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('v'), _dafny.CodePoint('i')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_113_v52, _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v53).Dtor_cf34())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v53).Dtor_cf35())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_114_v53).Dtor_cf36(), _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v53).Dtor_cf37())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_114_v53).Dtor_cf38())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_115_v54).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(143), _dafny.IntOfInt64(143))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_116_v55).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(3), true, _dafny.SeqOf(true, true, true), _dafny.IntOfInt64(143), _dafny.IntOfInt64(143)), _dafny.SeqOf(_dafny.CodePoint('w'), _dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('w')))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_117_v57, _dafny.SeqOf(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(3), true, _dafny.SeqOf(true, true, true), _dafny.IntOfInt64(143), _dafny.IntOfInt64(143)), Companion_D8_.Create_DC24_(_dafny.IntOfInt64(3), true, _dafny.SeqOf(true, true, true), _dafny.IntOfInt64(143), _dafny.IntOfInt64(143)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_118_v58, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D8_.Create_DC24_(_dafny.IntOfInt64(3), true, _dafny.SeqOf(true, true, true), _dafny.IntOfInt64(143), _dafny.IntOfInt64(143)), _dafny.UnicodeSeqOfUtf8Bytes("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj")))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_120_v60).Equals(_dafny.SetOf(_dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_187_v106).ArrayGet1((_dafny.Zero).Int()).(D6)).Dtor_cf24(), _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_187_v106).ArrayGet1((_dafny.One).Int()).(D6)).Dtor_cf24(), _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(((_187_v106).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(D6)).Dtor_cf24(), _dafny.SeqOf(true, true, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_190_v107).Dtor_cf24(), _dafny.SeqOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_194_v109).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_195_v110).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v111).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeerwdkj"), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v112).Dtor_cf74())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v112).Dtor_cf75())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_197_v112).Dtor_cf76())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_227_i12)
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
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_() D0 {
	return D0{D0_DC1{}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC0 struct {
	Cf0 _dafny.Int
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int) D0 {
	return D0{D0_DC0{Cf0}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC1_()
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC1:
		{
			return "D0.DC1"
		}
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ")"
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
			_, ok := other.Get_().(D0_DC1)
			return ok
		}
	case D0_DC0:
		{
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0
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

type D1_DC3 struct {
	Cf2 D0
	Cf3 _dafny.Int
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf2 D0, Cf3 _dafny.Int) D1 {
	return D1{D1_DC3{Cf2, Cf3}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

type D1_DC2 struct {
	Cf1 bool
}

func (D1_DC2) isD1() {}

func (CompanionStruct_D1_) Create_DC2_(Cf1 bool) D1 {
	return D1{D1_DC2{Cf1}}
}

func (_this D1) Is_DC2() bool {
	_, ok := _this.Get_().(D1_DC2)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC3_(Companion_D0_.Default(), _dafny.Zero)
}

func (_this D1) Dtor_cf2() D0 {
	return _this.Get_().(D1_DC3).Cf2
}

func (_this D1) Dtor_cf3() _dafny.Int {
	return _this.Get_().(D1_DC3).Cf3
}

func (_this D1) Dtor_cf1() bool {
	return _this.Get_().(D1_DC2).Cf1
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D1_DC2:
		{
			return "D1.DC2" + "(" + _dafny.String(data.Cf1) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf2.Equals(data2.Cf2) && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D1_DC2:
		{
			data2, ok := other.Get_().(D1_DC2)
			return ok && data1.Cf1 == data2.Cf1
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

type D2_DC5 struct {
	Cf5 bool
}

func (D2_DC5) isD2() {}

func (CompanionStruct_D2_) Create_DC5_(Cf5 bool) D2 {
	return D2{D2_DC5{Cf5}}
}

func (_this D2) Is_DC5() bool {
	_, ok := _this.Get_().(D2_DC5)
	return ok
}

type D2_DC6 struct {
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_() D2 {
	return D2{D2_DC6{}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC7 struct {
	Cf6 _dafny.CodePoint
	Cf7 _dafny.Map
	Cf8 _dafny.Int
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf6 _dafny.CodePoint, Cf7 _dafny.Map, Cf8 _dafny.Int) D2 {
	return D2{D2_DC7{Cf6, Cf7, Cf8}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC4 struct {
	Cf4 _dafny.Set
}

func (D2_DC4) isD2() {}

func (CompanionStruct_D2_) Create_DC4_(Cf4 _dafny.Set) D2 {
	return D2{D2_DC4{Cf4}}
}

func (_this D2) Is_DC4() bool {
	_, ok := _this.Get_().(D2_DC4)
	return ok
}

type D2_DC8 struct {
	Cf9 D2
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf9 D2) D2 {
	return D2{D2_DC8{Cf9}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC5_(false)
}

func (_this D2) Dtor_cf5() bool {
	return _this.Get_().(D2_DC5).Cf5
}

func (_this D2) Dtor_cf6() _dafny.CodePoint {
	return _this.Get_().(D2_DC7).Cf6
}

func (_this D2) Dtor_cf7() _dafny.Map {
	return _this.Get_().(D2_DC7).Cf7
}

func (_this D2) Dtor_cf8() _dafny.Int {
	return _this.Get_().(D2_DC7).Cf8
}

func (_this D2) Dtor_cf4() _dafny.Set {
	return _this.Get_().(D2_DC4).Cf4
}

func (_this D2) Dtor_cf9() D2 {
	return _this.Get_().(D2_DC8).Cf9
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC5:
		{
			return "D2.DC5" + "(" + _dafny.String(data.Cf5) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6"
		}
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf6) + ", " + _dafny.String(data.Cf7) + ", " + _dafny.String(data.Cf8) + ")"
		}
	case D2_DC4:
		{
			return "D2.DC4" + "(" + _dafny.String(data.Cf4) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf9) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC5:
		{
			data2, ok := other.Get_().(D2_DC5)
			return ok && data1.Cf5 == data2.Cf5
		}
	case D2_DC6:
		{
			_, ok := other.Get_().(D2_DC6)
			return ok
		}
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf6 == data2.Cf6 && data1.Cf7.Equals(data2.Cf7) && data1.Cf8.Cmp(data2.Cf8) == 0
		}
	case D2_DC4:
		{
			data2, ok := other.Get_().(D2_DC4)
			return ok && data1.Cf4.Equals(data2.Cf4)
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf9.Equals(data2.Cf9)
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

type D3_DC10 struct {
	Cf11 bool
	Cf12 _dafny.Int
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf11 bool, Cf12 _dafny.Int) D3 {
	return D3{D3_DC10{Cf11, Cf12}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC11 struct {
	Cf13 _dafny.Int
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf13 _dafny.Int) D3 {
	return D3{D3_DC11{Cf13}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC9 struct {
	Cf10 _dafny.Array
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf10 _dafny.Array) D3 {
	return D3{D3_DC9{Cf10}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC10_(false, _dafny.Zero)
}

func (_this D3) Dtor_cf11() bool {
	return _this.Get_().(D3_DC10).Cf11
}

func (_this D3) Dtor_cf12() _dafny.Int {
	return _this.Get_().(D3_DC10).Cf12
}

func (_this D3) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf13
}

func (_this D3) Dtor_cf10() _dafny.Array {
	return _this.Get_().(D3_DC9).Cf10
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf10) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12.Cmp(data2.Cf12) == 0
		}
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf10 == data2.Cf10
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

type D4_DC13 struct {
	Cf15 bool
	Cf16 _dafny.Map
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf15 bool, Cf16 _dafny.Map) D4 {
	return D4{D4_DC13{Cf15, Cf16}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

type D4_DC12 struct {
	Cf14 _dafny.Array
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf14 _dafny.Array) D4 {
	return D4{D4_DC12{Cf14}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC14 struct {
	Cf17 D4
}

func (D4_DC14) isD4() {}

func (CompanionStruct_D4_) Create_DC14_(Cf17 D4) D4 {
	return D4{D4_DC14{Cf17}}
}

func (_this D4) Is_DC14() bool {
	_, ok := _this.Get_().(D4_DC14)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC13_(false, _dafny.EmptyMap)
}

func (_this D4) Dtor_cf15() bool {
	return _this.Get_().(D4_DC13).Cf15
}

func (_this D4) Dtor_cf16() _dafny.Map {
	return _this.Get_().(D4_DC13).Cf16
}

func (_this D4) Dtor_cf14() _dafny.Array {
	return _this.Get_().(D4_DC12).Cf14
}

func (_this D4) Dtor_cf17() D4 {
	return _this.Get_().(D4_DC14).Cf17
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf15) + ", " + _dafny.String(data.Cf16) + ")"
		}
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D4_DC14:
		{
			return "D4.DC14" + "(" + _dafny.String(data.Cf17) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf15 == data2.Cf15 && data1.Cf16.Equals(data2.Cf16)
		}
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf14 == data2.Cf14
		}
	case D4_DC14:
		{
			data2, ok := other.Get_().(D4_DC14)
			return ok && data1.Cf17.Equals(data2.Cf17)
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

type D5_DC16 struct {
	Cf19 D3
	Cf20 _dafny.Int
	Cf21 _dafny.Int
	Cf22 _dafny.Int
	Cf23 *C0
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_(Cf19 D3, Cf20 _dafny.Int, Cf21 _dafny.Int, Cf22 _dafny.Int, Cf23 *C0) D5 {
	return D5{D5_DC16{Cf19, Cf20, Cf21, Cf22, Cf23}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC15 struct {
	Cf18 _dafny.Sequence
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf18 _dafny.Sequence) D5 {
	return D5{D5_DC15{Cf18}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC16_(Companion_D3_.Default(), _dafny.Zero, _dafny.Zero, _dafny.Zero, (*C0)(nil))
}

func (_this D5) Dtor_cf19() D3 {
	return _this.Get_().(D5_DC16).Cf19
}

func (_this D5) Dtor_cf20() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf20
}

func (_this D5) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf21
}

func (_this D5) Dtor_cf22() _dafny.Int {
	return _this.Get_().(D5_DC16).Cf22
}

func (_this D5) Dtor_cf23() *C0 {
	return _this.Get_().(D5_DC16).Cf23
}

func (_this D5) Dtor_cf18() _dafny.Sequence {
	return _this.Get_().(D5_DC15).Cf18
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC16:
		{
			return "D5.DC16" + "(" + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ", " + _dafny.String(data.Cf23) + ")"
		}
	case D5_DC15:
		{
			return "D5.DC15" + "(" + data.Cf18.VerbatimString(true) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D5) Equals(other D5) bool {
	switch data1 := _this.Get_().(type) {
	case D5_DC16:
		{
			data2, ok := other.Get_().(D5_DC16)
			return ok && data1.Cf19.Equals(data2.Cf19) && data1.Cf20.Cmp(data2.Cf20) == 0 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Cmp(data2.Cf22) == 0 && data1.Cf23 == data2.Cf23
		}
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D6_DC18 struct {
	Cf25 *C0
	Cf26 _dafny.Sequence
	Cf27 _dafny.Int
}

func (D6_DC18) isD6() {}

func (CompanionStruct_D6_) Create_DC18_(Cf25 *C0, Cf26 _dafny.Sequence, Cf27 _dafny.Int) D6 {
	return D6{D6_DC18{Cf25, Cf26, Cf27}}
}

func (_this D6) Is_DC18() bool {
	_, ok := _this.Get_().(D6_DC18)
	return ok
}

type D6_DC17 struct {
	Cf24 _dafny.Sequence
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf24 _dafny.Sequence) D6 {
	return D6{D6_DC17{Cf24}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC18_((*C0)(nil), _dafny.EmptySeq, _dafny.Zero)
}

func (_this D6) Dtor_cf25() *C0 {
	return _this.Get_().(D6_DC18).Cf25
}

func (_this D6) Dtor_cf26() _dafny.Sequence {
	return _this.Get_().(D6_DC18).Cf26
}

func (_this D6) Dtor_cf27() _dafny.Int {
	return _this.Get_().(D6_DC18).Cf27
}

func (_this D6) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf24
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC18:
		{
			return "D6.DC18" + "(" + _dafny.String(data.Cf25) + ", " + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ")"
		}
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf24) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC18:
		{
			data2, ok := other.Get_().(D6_DC18)
			return ok && data1.Cf25 == data2.Cf25 && data1.Cf26.Equals(data2.Cf26) && data1.Cf27.Cmp(data2.Cf27) == 0
		}
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D7_DC20 struct {
	Cf29 bool
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_(Cf29 bool) D7 {
	return D7{D7_DC20{Cf29}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC21 struct {
	Cf30 _dafny.Int
	Cf31 _dafny.Int
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf30 _dafny.Int, Cf31 _dafny.Int) D7 {
	return D7{D7_DC21{Cf30, Cf31}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC19 struct {
	Cf28 _dafny.Map
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf28 _dafny.Map) D7 {
	return D7{D7_DC19{Cf28}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC22 struct {
	Cf32 D7
}

func (D7_DC22) isD7() {}

func (CompanionStruct_D7_) Create_DC22_(Cf32 D7) D7 {
	return D7{D7_DC22{Cf32}}
}

func (_this D7) Is_DC22() bool {
	_, ok := _this.Get_().(D7_DC22)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC20_(false)
}

func (_this D7) Dtor_cf29() bool {
	return _this.Get_().(D7_DC20).Cf29
}

func (_this D7) Dtor_cf30() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf30
}

func (_this D7) Dtor_cf31() _dafny.Int {
	return _this.Get_().(D7_DC21).Cf31
}

func (_this D7) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D7_DC19).Cf28
}

func (_this D7) Dtor_cf32() D7 {
	return _this.Get_().(D7_DC22).Cf32
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC20:
		{
			return "D7.DC20" + "(" + _dafny.String(data.Cf29) + ")"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D7_DC22:
		{
			return "D7.DC22" + "(" + _dafny.String(data.Cf32) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC20:
		{
			data2, ok := other.Get_().(D7_DC20)
			return ok && data1.Cf29 == data2.Cf29
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf30.Cmp(data2.Cf30) == 0 && data1.Cf31.Cmp(data2.Cf31) == 0
		}
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D7_DC22:
		{
			data2, ok := other.Get_().(D7_DC22)
			return ok && data1.Cf32.Equals(data2.Cf32)
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

type D8_DC24 struct {
	Cf34 _dafny.Int
	Cf35 bool
	Cf36 _dafny.Sequence
	Cf37 _dafny.Int
	Cf38 _dafny.Int
}

func (D8_DC24) isD8() {}

func (CompanionStruct_D8_) Create_DC24_(Cf34 _dafny.Int, Cf35 bool, Cf36 _dafny.Sequence, Cf37 _dafny.Int, Cf38 _dafny.Int) D8 {
	return D8{D8_DC24{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D8) Is_DC24() bool {
	_, ok := _this.Get_().(D8_DC24)
	return ok
}

type D8_DC25 struct {
	Cf39 bool
	Cf40 bool
}

func (D8_DC25) isD8() {}

func (CompanionStruct_D8_) Create_DC25_(Cf39 bool, Cf40 bool) D8 {
	return D8{D8_DC25{Cf39, Cf40}}
}

func (_this D8) Is_DC25() bool {
	_, ok := _this.Get_().(D8_DC25)
	return ok
}

type D8_DC23 struct {
	Cf33 _dafny.Map
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf33 _dafny.Map) D8 {
	return D8{D8_DC23{Cf33}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC24_(_dafny.Zero, false, _dafny.EmptySeq, _dafny.Zero, _dafny.Zero)
}

func (_this D8) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf34
}

func (_this D8) Dtor_cf35() bool {
	return _this.Get_().(D8_DC24).Cf35
}

func (_this D8) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D8_DC24).Cf36
}

func (_this D8) Dtor_cf37() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf37
}

func (_this D8) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D8_DC24).Cf38
}

func (_this D8) Dtor_cf39() bool {
	return _this.Get_().(D8_DC25).Cf39
}

func (_this D8) Dtor_cf40() bool {
	return _this.Get_().(D8_DC25).Cf40
}

func (_this D8) Dtor_cf33() _dafny.Map {
	return _this.Get_().(D8_DC23).Cf33
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC24:
		{
			return "D8.DC24" + "(" + _dafny.String(data.Cf34) + ", " + _dafny.String(data.Cf35) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D8_DC25:
		{
			return "D8.DC25" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC24:
		{
			data2, ok := other.Get_().(D8_DC24)
			return ok && data1.Cf34.Cmp(data2.Cf34) == 0 && data1.Cf35 == data2.Cf35 && data1.Cf36.Equals(data2.Cf36) && data1.Cf37.Cmp(data2.Cf37) == 0 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D8_DC25:
		{
			data2, ok := other.Get_().(D8_DC25)
			return ok && data1.Cf39 == data2.Cf39 && data1.Cf40 == data2.Cf40
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf33.Equals(data2.Cf33)
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

type D9_DC27 struct {
	Cf42 _dafny.Map
	Cf43 bool
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf42 _dafny.Map, Cf43 bool) D9 {
	return D9{D9_DC27{Cf42, Cf43}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

type D9_DC26 struct {
	Cf41 *C2
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf41 *C2) D9 {
	return D9{D9_DC26{Cf41}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC28 struct {
	Cf44 D9
}

func (D9_DC28) isD9() {}

func (CompanionStruct_D9_) Create_DC28_(Cf44 D9) D9 {
	return D9{D9_DC28{Cf44}}
}

func (_this D9) Is_DC28() bool {
	_, ok := _this.Get_().(D9_DC28)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC27_(_dafny.EmptyMap, false)
}

func (_this D9) Dtor_cf42() _dafny.Map {
	return _this.Get_().(D9_DC27).Cf42
}

func (_this D9) Dtor_cf43() bool {
	return _this.Get_().(D9_DC27).Cf43
}

func (_this D9) Dtor_cf41() *C2 {
	return _this.Get_().(D9_DC26).Cf41
}

func (_this D9) Dtor_cf44() D9 {
	return _this.Get_().(D9_DC28).Cf44
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC28:
		{
			return "D9.DC28" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf42.Equals(data2.Cf42) && data1.Cf43 == data2.Cf43
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf41 == data2.Cf41
		}
	case D9_DC28:
		{
			data2, ok := other.Get_().(D9_DC28)
			return ok && data1.Cf44.Equals(data2.Cf44)
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

type D10_DC30 struct {
	Cf46 bool
	Cf47 bool
	Cf48 bool
}

func (D10_DC30) isD10() {}

func (CompanionStruct_D10_) Create_DC30_(Cf46 bool, Cf47 bool, Cf48 bool) D10 {
	return D10{D10_DC30{Cf46, Cf47, Cf48}}
}

func (_this D10) Is_DC30() bool {
	_, ok := _this.Get_().(D10_DC30)
	return ok
}

type D10_DC31 struct {
	Cf49 _dafny.Sequence
	Cf50 _dafny.CodePoint
}

func (D10_DC31) isD10() {}

func (CompanionStruct_D10_) Create_DC31_(Cf49 _dafny.Sequence, Cf50 _dafny.CodePoint) D10 {
	return D10{D10_DC31{Cf49, Cf50}}
}

func (_this D10) Is_DC31() bool {
	_, ok := _this.Get_().(D10_DC31)
	return ok
}

type D10_DC29 struct {
	Cf45 _dafny.MultiSet
}

func (D10_DC29) isD10() {}

func (CompanionStruct_D10_) Create_DC29_(Cf45 _dafny.MultiSet) D10 {
	return D10{D10_DC29{Cf45}}
}

func (_this D10) Is_DC29() bool {
	_, ok := _this.Get_().(D10_DC29)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC30_(false, false, false)
}

func (_this D10) Dtor_cf46() bool {
	return _this.Get_().(D10_DC30).Cf46
}

func (_this D10) Dtor_cf47() bool {
	return _this.Get_().(D10_DC30).Cf47
}

func (_this D10) Dtor_cf48() bool {
	return _this.Get_().(D10_DC30).Cf48
}

func (_this D10) Dtor_cf49() _dafny.Sequence {
	return _this.Get_().(D10_DC31).Cf49
}

func (_this D10) Dtor_cf50() _dafny.CodePoint {
	return _this.Get_().(D10_DC31).Cf50
}

func (_this D10) Dtor_cf45() _dafny.MultiSet {
	return _this.Get_().(D10_DC29).Cf45
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC30:
		{
			return "D10.DC30" + "(" + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ", " + _dafny.String(data.Cf48) + ")"
		}
	case D10_DC31:
		{
			return "D10.DC31" + "(" + _dafny.String(data.Cf49) + ", " + _dafny.String(data.Cf50) + ")"
		}
	case D10_DC29:
		{
			return "D10.DC29" + "(" + _dafny.String(data.Cf45) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC30:
		{
			data2, ok := other.Get_().(D10_DC30)
			return ok && data1.Cf46 == data2.Cf46 && data1.Cf47 == data2.Cf47 && data1.Cf48 == data2.Cf48
		}
	case D10_DC31:
		{
			data2, ok := other.Get_().(D10_DC31)
			return ok && data1.Cf49.Equals(data2.Cf49) && data1.Cf50 == data2.Cf50
		}
	case D10_DC29:
		{
			data2, ok := other.Get_().(D10_DC29)
			return ok && data1.Cf45.Equals(data2.Cf45)
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

type D11_DC33 struct {
	Cf52 bool
	Cf53 bool
}

func (D11_DC33) isD11() {}

func (CompanionStruct_D11_) Create_DC33_(Cf52 bool, Cf53 bool) D11 {
	return D11{D11_DC33{Cf52, Cf53}}
}

func (_this D11) Is_DC33() bool {
	_, ok := _this.Get_().(D11_DC33)
	return ok
}

type D11_DC32 struct {
	Cf51 _dafny.Map
}

func (D11_DC32) isD11() {}

func (CompanionStruct_D11_) Create_DC32_(Cf51 _dafny.Map) D11 {
	return D11{D11_DC32{Cf51}}
}

func (_this D11) Is_DC32() bool {
	_, ok := _this.Get_().(D11_DC32)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC33_(false, false)
}

func (_this D11) Dtor_cf52() bool {
	return _this.Get_().(D11_DC33).Cf52
}

func (_this D11) Dtor_cf53() bool {
	return _this.Get_().(D11_DC33).Cf53
}

func (_this D11) Dtor_cf51() _dafny.Map {
	return _this.Get_().(D11_DC32).Cf51
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC33:
		{
			return "D11.DC33" + "(" + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ")"
		}
	case D11_DC32:
		{
			return "D11.DC32" + "(" + _dafny.String(data.Cf51) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC33:
		{
			data2, ok := other.Get_().(D11_DC33)
			return ok && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53
		}
	case D11_DC32:
		{
			data2, ok := other.Get_().(D11_DC32)
			return ok && data1.Cf51.Equals(data2.Cf51)
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

type D12_DC35 struct {
	Cf55 bool
}

func (D12_DC35) isD12() {}

func (CompanionStruct_D12_) Create_DC35_(Cf55 bool) D12 {
	return D12{D12_DC35{Cf55}}
}

func (_this D12) Is_DC35() bool {
	_, ok := _this.Get_().(D12_DC35)
	return ok
}

type D12_DC36 struct {
	Cf56 _dafny.Int
	Cf57 _dafny.Int
}

func (D12_DC36) isD12() {}

func (CompanionStruct_D12_) Create_DC36_(Cf56 _dafny.Int, Cf57 _dafny.Int) D12 {
	return D12{D12_DC36{Cf56, Cf57}}
}

func (_this D12) Is_DC36() bool {
	_, ok := _this.Get_().(D12_DC36)
	return ok
}

type D12_DC37 struct {
	Cf58 _dafny.CodePoint
}

func (D12_DC37) isD12() {}

func (CompanionStruct_D12_) Create_DC37_(Cf58 _dafny.CodePoint) D12 {
	return D12{D12_DC37{Cf58}}
}

func (_this D12) Is_DC37() bool {
	_, ok := _this.Get_().(D12_DC37)
	return ok
}

type D12_DC34 struct {
	Cf54 _dafny.Map
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf54 _dafny.Map) D12 {
	return D12{D12_DC34{Cf54}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC35_(false)
}

func (_this D12) Dtor_cf55() bool {
	return _this.Get_().(D12_DC35).Cf55
}

func (_this D12) Dtor_cf56() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf56
}

func (_this D12) Dtor_cf57() _dafny.Int {
	return _this.Get_().(D12_DC36).Cf57
}

func (_this D12) Dtor_cf58() _dafny.CodePoint {
	return _this.Get_().(D12_DC37).Cf58
}

func (_this D12) Dtor_cf54() _dafny.Map {
	return _this.Get_().(D12_DC34).Cf54
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC35:
		{
			return "D12.DC35" + "(" + _dafny.String(data.Cf55) + ")"
		}
	case D12_DC36:
		{
			return "D12.DC36" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ")"
		}
	case D12_DC37:
		{
			return "D12.DC37" + "(" + _dafny.String(data.Cf58) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf54) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC35:
		{
			data2, ok := other.Get_().(D12_DC35)
			return ok && data1.Cf55 == data2.Cf55
		}
	case D12_DC36:
		{
			data2, ok := other.Get_().(D12_DC36)
			return ok && data1.Cf56.Cmp(data2.Cf56) == 0 && data1.Cf57.Cmp(data2.Cf57) == 0
		}
	case D12_DC37:
		{
			data2, ok := other.Get_().(D12_DC37)
			return ok && data1.Cf58 == data2.Cf58
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
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

type D13_DC39 struct {
	Cf60 _dafny.Map
	Cf61 _dafny.Array
	Cf62 _dafny.Sequence
	Cf63 _dafny.Int
	Cf64 _dafny.Sequence
}

func (D13_DC39) isD13() {}

func (CompanionStruct_D13_) Create_DC39_(Cf60 _dafny.Map, Cf61 _dafny.Array, Cf62 _dafny.Sequence, Cf63 _dafny.Int, Cf64 _dafny.Sequence) D13 {
	return D13{D13_DC39{Cf60, Cf61, Cf62, Cf63, Cf64}}
}

func (_this D13) Is_DC39() bool {
	_, ok := _this.Get_().(D13_DC39)
	return ok
}

type D13_DC38 struct {
	Cf59 _dafny.Map
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf59 _dafny.Map) D13 {
	return D13{D13_DC38{Cf59}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC39_(_dafny.EmptyMap, _dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.EmptySeq, _dafny.Zero, _dafny.EmptySeq)
}

func (_this D13) Dtor_cf60() _dafny.Map {
	return _this.Get_().(D13_DC39).Cf60
}

func (_this D13) Dtor_cf61() _dafny.Array {
	return _this.Get_().(D13_DC39).Cf61
}

func (_this D13) Dtor_cf62() _dafny.Sequence {
	return _this.Get_().(D13_DC39).Cf62
}

func (_this D13) Dtor_cf63() _dafny.Int {
	return _this.Get_().(D13_DC39).Cf63
}

func (_this D13) Dtor_cf64() _dafny.Sequence {
	return _this.Get_().(D13_DC39).Cf64
}

func (_this D13) Dtor_cf59() _dafny.Map {
	return _this.Get_().(D13_DC38).Cf59
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC39:
		{
			return "D13.DC39" + "(" + _dafny.String(data.Cf60) + ", " + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ", " + _dafny.String(data.Cf63) + ", " + _dafny.String(data.Cf64) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf59) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC39:
		{
			data2, ok := other.Get_().(D13_DC39)
			return ok && data1.Cf60.Equals(data2.Cf60) && data1.Cf61 == data2.Cf61 && data1.Cf62.Equals(data2.Cf62) && data1.Cf63.Cmp(data2.Cf63) == 0 && data1.Cf64.Equals(data2.Cf64)
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf59.Equals(data2.Cf59)
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

type D14_DC40 struct {
	Cf65 _dafny.Map
}

func (D14_DC40) isD14() {}

func (CompanionStruct_D14_) Create_DC40_(Cf65 _dafny.Map) D14 {
	return D14{D14_DC40{Cf65}}
}

func (_this D14) Is_DC40() bool {
	_, ok := _this.Get_().(D14_DC40)
	return ok
}

func (CompanionStruct_D14_) Default() _dafny.Map {
	return _dafny.EmptyMap
}

func (_this D14) Dtor_cf65() _dafny.Map {
	return _this.Get_().(D14_DC40).Cf65
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC40:
		{
			return "D14.DC40" + "(" + _dafny.String(data.Cf65) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D14) Equals(other D14) bool {
	switch data1 := _this.Get_().(type) {
	case D14_DC40:
		{
			data2, ok := other.Get_().(D14_DC40)
			return ok && data1.Cf65.Equals(data2.Cf65)
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

type D15_DC42 struct {
	Cf67 _dafny.MultiSet
}

func (D15_DC42) isD15() {}

func (CompanionStruct_D15_) Create_DC42_(Cf67 _dafny.MultiSet) D15 {
	return D15{D15_DC42{Cf67}}
}

func (_this D15) Is_DC42() bool {
	_, ok := _this.Get_().(D15_DC42)
	return ok
}

type D15_DC41 struct {
	Cf66 _dafny.Map
}

func (D15_DC41) isD15() {}

func (CompanionStruct_D15_) Create_DC41_(Cf66 _dafny.Map) D15 {
	return D15{D15_DC41{Cf66}}
}

func (_this D15) Is_DC41() bool {
	_, ok := _this.Get_().(D15_DC41)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC42_(_dafny.EmptyMultiSet)
}

func (_this D15) Dtor_cf67() _dafny.MultiSet {
	return _this.Get_().(D15_DC42).Cf67
}

func (_this D15) Dtor_cf66() _dafny.Map {
	return _this.Get_().(D15_DC41).Cf66
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC42:
		{
			return "D15.DC42" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D15_DC41:
		{
			return "D15.DC41" + "(" + _dafny.String(data.Cf66) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC42:
		{
			data2, ok := other.Get_().(D15_DC42)
			return ok && data1.Cf67.Equals(data2.Cf67)
		}
	case D15_DC41:
		{
			data2, ok := other.Get_().(D15_DC41)
			return ok && data1.Cf66.Equals(data2.Cf66)
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

type D16_DC44 struct {
	Cf69 bool
	Cf70 bool
	Cf71 _dafny.Int
}

func (D16_DC44) isD16() {}

func (CompanionStruct_D16_) Create_DC44_(Cf69 bool, Cf70 bool, Cf71 _dafny.Int) D16 {
	return D16{D16_DC44{Cf69, Cf70, Cf71}}
}

func (_this D16) Is_DC44() bool {
	_, ok := _this.Get_().(D16_DC44)
	return ok
}

type D16_DC43 struct {
	Cf68 _dafny.MultiSet
}

func (D16_DC43) isD16() {}

func (CompanionStruct_D16_) Create_DC43_(Cf68 _dafny.MultiSet) D16 {
	return D16{D16_DC43{Cf68}}
}

func (_this D16) Is_DC43() bool {
	_, ok := _this.Get_().(D16_DC43)
	return ok
}

type D16_DC45 struct {
	Cf72 D16
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf72 D16) D16 {
	return D16{D16_DC45{Cf72}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC44_(false, false, _dafny.Zero)
}

func (_this D16) Dtor_cf69() bool {
	return _this.Get_().(D16_DC44).Cf69
}

func (_this D16) Dtor_cf70() bool {
	return _this.Get_().(D16_DC44).Cf70
}

func (_this D16) Dtor_cf71() _dafny.Int {
	return _this.Get_().(D16_DC44).Cf71
}

func (_this D16) Dtor_cf68() _dafny.MultiSet {
	return _this.Get_().(D16_DC43).Cf68
}

func (_this D16) Dtor_cf72() D16 {
	return _this.Get_().(D16_DC45).Cf72
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC44:
		{
			return "D16.DC44" + "(" + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ", " + _dafny.String(data.Cf71) + ")"
		}
	case D16_DC43:
		{
			return "D16.DC43" + "(" + _dafny.String(data.Cf68) + ")"
		}
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf72) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC44:
		{
			data2, ok := other.Get_().(D16_DC44)
			return ok && data1.Cf69 == data2.Cf69 && data1.Cf70 == data2.Cf70 && data1.Cf71.Cmp(data2.Cf71) == 0
		}
	case D16_DC43:
		{
			data2, ok := other.Get_().(D16_DC43)
			return ok && data1.Cf68.Equals(data2.Cf68)
		}
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf72.Equals(data2.Cf72)
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

type D17_DC47 struct {
	Cf74 bool
	Cf75 _dafny.Int
	Cf76 _dafny.Int
}

func (D17_DC47) isD17() {}

func (CompanionStruct_D17_) Create_DC47_(Cf74 bool, Cf75 _dafny.Int, Cf76 _dafny.Int) D17 {
	return D17{D17_DC47{Cf74, Cf75, Cf76}}
}

func (_this D17) Is_DC47() bool {
	_, ok := _this.Get_().(D17_DC47)
	return ok
}

type D17_DC46 struct {
	Cf73 _dafny.MultiSet
}

func (D17_DC46) isD17() {}

func (CompanionStruct_D17_) Create_DC46_(Cf73 _dafny.MultiSet) D17 {
	return D17{D17_DC46{Cf73}}
}

func (_this D17) Is_DC46() bool {
	_, ok := _this.Get_().(D17_DC46)
	return ok
}

func (CompanionStruct_D17_) Default() D17 {
	return Companion_D17_.Create_DC47_(false, _dafny.Zero, _dafny.Zero)
}

func (_this D17) Dtor_cf74() bool {
	return _this.Get_().(D17_DC47).Cf74
}

func (_this D17) Dtor_cf75() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf75
}

func (_this D17) Dtor_cf76() _dafny.Int {
	return _this.Get_().(D17_DC47).Cf76
}

func (_this D17) Dtor_cf73() _dafny.MultiSet {
	return _this.Get_().(D17_DC46).Cf73
}

func (_this D17) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D17_DC47:
		{
			return "D17.DC47" + "(" + _dafny.String(data.Cf74) + ", " + _dafny.String(data.Cf75) + ", " + _dafny.String(data.Cf76) + ")"
		}
	case D17_DC46:
		{
			return "D17.DC46" + "(" + _dafny.String(data.Cf73) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D17) Equals(other D17) bool {
	switch data1 := _this.Get_().(type) {
	case D17_DC47:
		{
			data2, ok := other.Get_().(D17_DC47)
			return ok && data1.Cf74 == data2.Cf74 && data1.Cf75.Cmp(data2.Cf75) == 0 && data1.Cf76.Cmp(data2.Cf76) == 0
		}
	case D17_DC46:
		{
			data2, ok := other.Get_().(D17_DC46)
			return ok && data1.Cf73.Equals(data2.Cf73)
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

type D18_DC49 struct {
	Cf78 _dafny.Sequence
	Cf79 _dafny.Int
}

func (D18_DC49) isD18() {}

func (CompanionStruct_D18_) Create_DC49_(Cf78 _dafny.Sequence, Cf79 _dafny.Int) D18 {
	return D18{D18_DC49{Cf78, Cf79}}
}

func (_this D18) Is_DC49() bool {
	_, ok := _this.Get_().(D18_DC49)
	return ok
}

type D18_DC48 struct {
	Cf77 _dafny.MultiSet
}

func (D18_DC48) isD18() {}

func (CompanionStruct_D18_) Create_DC48_(Cf77 _dafny.MultiSet) D18 {
	return D18{D18_DC48{Cf77}}
}

func (_this D18) Is_DC48() bool {
	_, ok := _this.Get_().(D18_DC48)
	return ok
}

type D18_DC50 struct {
	Cf80 D18
}

func (D18_DC50) isD18() {}

func (CompanionStruct_D18_) Create_DC50_(Cf80 D18) D18 {
	return D18{D18_DC50{Cf80}}
}

func (_this D18) Is_DC50() bool {
	_, ok := _this.Get_().(D18_DC50)
	return ok
}

func (CompanionStruct_D18_) Default() D18 {
	return Companion_D18_.Create_DC49_(_dafny.EmptySeq, _dafny.Zero)
}

func (_this D18) Dtor_cf78() _dafny.Sequence {
	return _this.Get_().(D18_DC49).Cf78
}

func (_this D18) Dtor_cf79() _dafny.Int {
	return _this.Get_().(D18_DC49).Cf79
}

func (_this D18) Dtor_cf77() _dafny.MultiSet {
	return _this.Get_().(D18_DC48).Cf77
}

func (_this D18) Dtor_cf80() D18 {
	return _this.Get_().(D18_DC50).Cf80
}

func (_this D18) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D18_DC49:
		{
			return "D18.DC49" + "(" + _dafny.String(data.Cf78) + ", " + _dafny.String(data.Cf79) + ")"
		}
	case D18_DC48:
		{
			return "D18.DC48" + "(" + _dafny.String(data.Cf77) + ")"
		}
	case D18_DC50:
		{
			return "D18.DC50" + "(" + _dafny.String(data.Cf80) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D18) Equals(other D18) bool {
	switch data1 := _this.Get_().(type) {
	case D18_DC49:
		{
			data2, ok := other.Get_().(D18_DC49)
			return ok && data1.Cf78.Equals(data2.Cf78) && data1.Cf79.Cmp(data2.Cf79) == 0
		}
	case D18_DC48:
		{
			data2, ok := other.Get_().(D18_DC48)
			return ok && data1.Cf77.Equals(data2.Cf77)
		}
	case D18_DC50:
		{
			data2, ok := other.Get_().(D18_DC50)
			return ok && data1.Cf80.Equals(data2.Cf80)
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

type D19_DC52 struct {
	Cf82 _dafny.Int
}

func (D19_DC52) isD19() {}

func (CompanionStruct_D19_) Create_DC52_(Cf82 _dafny.Int) D19 {
	return D19{D19_DC52{Cf82}}
}

func (_this D19) Is_DC52() bool {
	_, ok := _this.Get_().(D19_DC52)
	return ok
}

type D19_DC51 struct {
	Cf81 _dafny.Map
}

func (D19_DC51) isD19() {}

func (CompanionStruct_D19_) Create_DC51_(Cf81 _dafny.Map) D19 {
	return D19{D19_DC51{Cf81}}
}

func (_this D19) Is_DC51() bool {
	_, ok := _this.Get_().(D19_DC51)
	return ok
}

func (CompanionStruct_D19_) Default() D19 {
	return Companion_D19_.Create_DC52_(_dafny.Zero)
}

func (_this D19) Dtor_cf82() _dafny.Int {
	return _this.Get_().(D19_DC52).Cf82
}

func (_this D19) Dtor_cf81() _dafny.Map {
	return _this.Get_().(D19_DC51).Cf81
}

func (_this D19) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D19_DC52:
		{
			return "D19.DC52" + "(" + _dafny.String(data.Cf82) + ")"
		}
	case D19_DC51:
		{
			return "D19.DC51" + "(" + _dafny.String(data.Cf81) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D19) Equals(other D19) bool {
	switch data1 := _this.Get_().(type) {
	case D19_DC52:
		{
			data2, ok := other.Get_().(D19_DC52)
			return ok && data1.Cf82.Cmp(data2.Cf82) == 0
		}
	case D19_DC51:
		{
			data2, ok := other.Get_().(D19_DC51)
			return ok && data1.Cf81.Equals(data2.Cf81)
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

// Definition of datatype D20
type D20 struct {
	Data_D20_
}

func (_this D20) Get_() Data_D20_ {
	return _this.Data_D20_
}

type Data_D20_ interface {
	isD20()
}

type CompanionStruct_D20_ struct {
}

var Companion_D20_ = CompanionStruct_D20_{}

type D20_DC53 struct {
	Cf83 _dafny.MultiSet
}

func (D20_DC53) isD20() {}

func (CompanionStruct_D20_) Create_DC53_(Cf83 _dafny.MultiSet) D20 {
	return D20{D20_DC53{Cf83}}
}

func (_this D20) Is_DC53() bool {
	_, ok := _this.Get_().(D20_DC53)
	return ok
}

func (CompanionStruct_D20_) Default() _dafny.MultiSet {
	return _dafny.EmptyMultiSet
}

func (_this D20) Dtor_cf83() _dafny.MultiSet {
	return _this.Get_().(D20_DC53).Cf83
}

func (_this D20) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D20_DC53:
		{
			return "D20.DC53" + "(" + _dafny.String(data.Cf83) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D20) Equals(other D20) bool {
	switch data1 := _this.Get_().(type) {
	case D20_DC53:
		{
			data2, ok := other.Get_().(D20_DC53)
			return ok && data1.Cf83.Equals(data2.Cf83)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D20) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D20)
	return ok && _this.Equals(typed)
}

func Type_D20_() _dafny.TypeDescriptor {
	return type_D20_{}
}

type type_D20_ struct {
}

func (_this type_D20_) Default() interface{} {
	return Companion_D20_.Default()
}

func (_this type_D20_) String() string {
	return "main.D20"
}
func (_this D20) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D20{}

// End of datatype D20

// Definition of datatype D21
type D21 struct {
	Data_D21_
}

func (_this D21) Get_() Data_D21_ {
	return _this.Data_D21_
}

type Data_D21_ interface {
	isD21()
}

type CompanionStruct_D21_ struct {
}

var Companion_D21_ = CompanionStruct_D21_{}

type D21_DC55 struct {
	Cf85 bool
}

func (D21_DC55) isD21() {}

func (CompanionStruct_D21_) Create_DC55_(Cf85 bool) D21 {
	return D21{D21_DC55{Cf85}}
}

func (_this D21) Is_DC55() bool {
	_, ok := _this.Get_().(D21_DC55)
	return ok
}

type D21_DC56 struct {
	Cf86 _dafny.Int
	Cf87 _dafny.Int
}

func (D21_DC56) isD21() {}

func (CompanionStruct_D21_) Create_DC56_(Cf86 _dafny.Int, Cf87 _dafny.Int) D21 {
	return D21{D21_DC56{Cf86, Cf87}}
}

func (_this D21) Is_DC56() bool {
	_, ok := _this.Get_().(D21_DC56)
	return ok
}

type D21_DC54 struct {
	Cf84 _dafny.Sequence
}

func (D21_DC54) isD21() {}

func (CompanionStruct_D21_) Create_DC54_(Cf84 _dafny.Sequence) D21 {
	return D21{D21_DC54{Cf84}}
}

func (_this D21) Is_DC54() bool {
	_, ok := _this.Get_().(D21_DC54)
	return ok
}

func (CompanionStruct_D21_) Default() D21 {
	return Companion_D21_.Create_DC55_(false)
}

func (_this D21) Dtor_cf85() bool {
	return _this.Get_().(D21_DC55).Cf85
}

func (_this D21) Dtor_cf86() _dafny.Int {
	return _this.Get_().(D21_DC56).Cf86
}

func (_this D21) Dtor_cf87() _dafny.Int {
	return _this.Get_().(D21_DC56).Cf87
}

func (_this D21) Dtor_cf84() _dafny.Sequence {
	return _this.Get_().(D21_DC54).Cf84
}

func (_this D21) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D21_DC55:
		{
			return "D21.DC55" + "(" + _dafny.String(data.Cf85) + ")"
		}
	case D21_DC56:
		{
			return "D21.DC56" + "(" + _dafny.String(data.Cf86) + ", " + _dafny.String(data.Cf87) + ")"
		}
	case D21_DC54:
		{
			return "D21.DC54" + "(" + _dafny.String(data.Cf84) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D21) Equals(other D21) bool {
	switch data1 := _this.Get_().(type) {
	case D21_DC55:
		{
			data2, ok := other.Get_().(D21_DC55)
			return ok && data1.Cf85 == data2.Cf85
		}
	case D21_DC56:
		{
			data2, ok := other.Get_().(D21_DC56)
			return ok && data1.Cf86.Cmp(data2.Cf86) == 0 && data1.Cf87.Cmp(data2.Cf87) == 0
		}
	case D21_DC54:
		{
			data2, ok := other.Get_().(D21_DC54)
			return ok && data1.Cf84.Equals(data2.Cf84)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D21) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D21)
	return ok && _this.Equals(typed)
}

func Type_D21_() _dafny.TypeDescriptor {
	return type_D21_{}
}

type type_D21_ struct {
}

func (_this type_D21_) Default() interface{} {
	return Companion_D21_.Default()
}

func (_this type_D21_) String() string {
	return "main.D21"
}
func (_this D21) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D21{}

// End of datatype D21

// Definition of datatype D22
type D22 struct {
	Data_D22_
}

func (_this D22) Get_() Data_D22_ {
	return _this.Data_D22_
}

type Data_D22_ interface {
	isD22()
}

type CompanionStruct_D22_ struct {
}

var Companion_D22_ = CompanionStruct_D22_{}

type D22_DC58 struct {
}

func (D22_DC58) isD22() {}

func (CompanionStruct_D22_) Create_DC58_() D22 {
	return D22{D22_DC58{}}
}

func (_this D22) Is_DC58() bool {
	_, ok := _this.Get_().(D22_DC58)
	return ok
}

type D22_DC57 struct {
	Cf88 T0
}

func (D22_DC57) isD22() {}

func (CompanionStruct_D22_) Create_DC57_(Cf88 T0) D22 {
	return D22{D22_DC57{Cf88}}
}

func (_this D22) Is_DC57() bool {
	_, ok := _this.Get_().(D22_DC57)
	return ok
}

func (CompanionStruct_D22_) Default() D22 {
	return Companion_D22_.Create_DC58_()
}

func (_this D22) Dtor_cf88() T0 {
	return _this.Get_().(D22_DC57).Cf88
}

func (_this D22) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D22_DC58:
		{
			return "D22.DC58"
		}
	case D22_DC57:
		{
			return "D22.DC57" + "(" + _dafny.String(data.Cf88) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D22) Equals(other D22) bool {
	switch data1 := _this.Get_().(type) {
	case D22_DC58:
		{
			_, ok := other.Get_().(D22_DC58)
			return ok
		}
	case D22_DC57:
		{
			data2, ok := other.Get_().(D22_DC57)
			return ok && _dafny.AreEqual(data1.Cf88, data2.Cf88)
		}
	default:
		{
			return false // unexpected
		}
	}
}

func (_this D22) EqualsGeneric(other interface{}) bool {
	typed, ok := other.(D22)
	return ok && _this.Equals(typed)
}

func Type_D22_() _dafny.TypeDescriptor {
	return type_D22_{}
}

type type_D22_ struct {
}

func (_this type_D22_) Default() interface{} {
	return Companion_D22_.Default()
}

func (_this type_D22_) String() string {
	return "main.D22"
}
func (_this D22) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = D22{}

// End of datatype D22

// Definition of trait T0
type T0 interface {
	String() string
	Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool
	Fm2(p0 bool, globalState *GlobalState) bool
	F5() _dafny.MultiSet
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
	Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool
	Fm2(p0 bool, globalState *GlobalState) bool
	F5() _dafny.MultiSet
	Fm3(globalState *GlobalState) bool
	Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState)
	F6() bool
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

// Definition of trait T2
type T2 interface {
	String() string
	Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool
	Fm2(p0 bool, globalState *GlobalState) bool
	Fm3(globalState *GlobalState) bool
	Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int
	M1(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState)
	F6() bool
	F5() _dafny.MultiSet
	F7() _dafny.Array
	F7_set_(value _dafny.Array)
	F8() bool
	F8_set_(value bool)
	M2(p0 _dafny.CodePoint, globalState *GlobalState)
	M3(globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.Int)
}
type CompanionStruct_T2_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T2_ = CompanionStruct_T2_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T2_) CastTo_(x interface{}) T2 {
	var t T2
	t, _ = x.(T2)
	return t
}

// End of trait T2

// Definition of class GlobalState
type GlobalState struct {
	F0  _dafny.Int
	F2  _dafny.Array
	_f1 _dafny.Sequence
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F0 = _dafny.Zero
	_this.F2 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f1 = _dafny.EmptySeq
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Sequence, f2 _dafny.Array) {
	{
		(_this).F0 = f0
		(_this)._f1 = f1
		(_this).F2 = f2
	}
}
func (_this *GlobalState) F1() _dafny.Sequence {
	{
		return _this._f1
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f11 _dafny.Map
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f11 = _dafny.EmptyMap
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

func (_this *C0) Ctor__(f11 _dafny.Map) {
	{
		(_this)._f11 = f11
	}
}
func (_this *C0) Fm10(p0 _dafny.Int, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("abyn"))).Difference(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("tnbomydo")))).IsDisjointFrom((func() _dafny.Set {
			var _coll20 = _dafny.NewBuilder()
			_ = _coll20
			for _iter21 := _dafny.Iterate((_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bfuu"), _dafny.UnicodeSeqOfUtf8Bytes("qgnlm"))).Elements()); ; {
				_compr_20, _ok21 := _iter21()
				if !_ok21 {
					break
				}
				var _235_v0 _dafny.Sequence
				_235_v0 = interface{}(_compr_20).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("bfuu"), _dafny.UnicodeSeqOfUtf8Bytes("qgnlm")), _235_v0) {
					_coll20.Add(_235_v0)
				}
			}
			return _coll20.ToSet()
		}()).Union(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("hr"))))
	}
}
func (_this *C0) F11() _dafny.Map {
	{
		return _this._f11
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f6  bool
	_f5  _dafny.MultiSet
	_f12 bool
	_f13 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f6 = false
	_this._f5 = _dafny.EmptyMultiSet
	_this._f12 = false
	_this._f13 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F6() bool {
	return _this._f6
}
func (_this *C1) F5() _dafny.MultiSet {
	return _this._f5
}
func (_this *C1) Ctor__(f12 bool, f13 _dafny.Sequence, f6 bool, f5 _dafny.MultiSet) {
	{
		(_this)._f12 = f12
		(_this)._f13 = f13
		(_this)._f6 = f6
		(_this)._f5 = f5
	}
}
func (_this *C1) Fm3(globalState *GlobalState) bool {
	{
		return (_dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32(((_this).F13()).Cardinality()))).Cardinality())).IsProperSubsetOf(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(!((_this).F6()), (_this).F12())).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(323))).Uint32(), func(coer22 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg22 _dafny.Int) interface{} {
				return coer22(arg22)
			}
		}(func(_236_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(579)
		}))).Cardinality()), _dafny.IntOfInt64(821), (_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F12())).Cardinality()), _dafny.IntOfInt64(562))).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("sejx"), (Companion_Default___.SafeIndex((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), true)).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sejx")).Cardinality()))).Uint32(), _dafny.CodePoint('i'))).Cardinality())))
	}
}
func (_this *C1) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(((func() _dafny.MultiSet {
			if ((_dafny.Zero).Minus(_dafny.IntOfInt64(-463))).Cmp((_dafny.SetOf((_this).F12(), (_this).F6())).Cardinality()) >= 0 {
				return (_this).F5()
			}
			return ((_this).F5()).Update((_this).F12(), Companion_Default___.Abs(_dafny.IntOfInt64(-762)))
		})()).Cardinality())
	}
}
func (_this *C1) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C1) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.SeqOf((_this).F6()))).Union(_dafny.SetOf(_dafny.SeqOf((_this).F6()), _dafny.SeqOf((_this).F6(), (_this).F6(), (_this).F6())))).IsDisjointFrom(_dafny.SetOf(_dafny.SeqOf((_this).F6()), _dafny.SeqOf((_this).F12(), true), _dafny.SeqOf(!((_this).F12())), _dafny.SeqOf((_this).F12())))
	}
}
func (_this *C1) M1(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) {
	{
		var _237_v0 _dafny.Array
		_ = _237_v0
		var _nwElement0_4 bool = p2
		_ = _nwElement0_4
		var _nw37 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(6))
		_ = _nw37
		(_nw37).ArraySet1(_nwElement0_4, 0)
		(_nw37).ArraySet1(p2, 1)
		(_nw37).ArraySet1((_this).F6(), 2)
		(_nw37).ArraySet1(p2, 3)
		(_nw37).ArraySet1((_this).F12(), 4)
		(_nw37).ArraySet1((_this).F6(), 5)
		_237_v0 = _nw37
		var _238_v1 _dafny.MultiSet
		_ = _238_v1
		_238_v1 = _dafny.MultiSetOf(_237_v0)
		var _239_i0 _dafny.Int
		_ = _239_i0
		_239_i0 = _dafny.Zero
		{
			for ((_238_v1).Difference(_238_v1)).IsSubsetOf(_238_v1) {
				{
					if (_239_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_239_i0 = (_239_i0).Plus(_dafny.One)
					var _240_v2 _dafny.Set
					_ = _240_v2
					_240_v2 = _dafny.SetOf(p0)
					var _241_v3 D2
					_ = _241_v3
					_241_v3 = Companion_D2_.Create_DC4_(_240_v2)
					var _242_v4 _dafny.Sequence
					_ = _242_v4
					_242_v4 = _dafny.SeqOf(_241_v3)
					var _243_v5 _dafny.Sequence
					_ = _243_v5
					_243_v5 = _dafny.SeqOf(_242_v4)
					var _244_v6 _dafny.Map
					_ = _244_v6
					_244_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _242_v4)
					_243_v5 = _dafny.Companion_Sequence_.Update(_dafny.SeqOf(_242_v4), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf(_242_v4)).Cardinality()))).Uint32(), (func() _dafny.Sequence {
						if (_244_v6).Contains(p0) {
							return (_244_v6).Get(p0).(_dafny.Sequence)
						}
						return _242_v4
					})())
					var _245_v7 _dafny.Array
					_ = _245_v7
					var _nw38 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(17))
					_ = _nw38
					_245_v7 = _nw38
					var _246_v8 _dafny.CodePoint
					_ = _246_v8
					_246_v8 = _dafny.CodePoint('a')
					var _247_v9 _dafny.Map
					_ = _247_v9
					_247_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfInt64(-852))
					var _248_v10 D2
					_ = _248_v10
					_248_v10 = Companion_D2_.Create_DC7_(_246_v8, _247_v9, (_dafny.Zero).Minus((_dafny.Zero).Minus(p0)))
					var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_245_v7), 0))
					_ = _index48
					(_245_v7).ArraySet1(_248_v10, (_index48).Int())
					var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_245_v7), 0))
					_ = _index49
					(_245_v7).ArraySet1(_248_v10, (_index49).Int())
					var _249_v11 bool
					_ = _249_v11
					_249_v11 = true
					_249_v11 = (_this).F12()
					(globalState).F0 = p0
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		var _250_v12 _dafny.CodePoint
		_ = _250_v12
		_250_v12 = _dafny.CodePoint('n')
		var _251_v13 _dafny.Map
		_ = _251_v13
		_251_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F5()).Cardinality(), p0)
		var _252_v14 D2
		_ = _252_v14
		_252_v14 = Companion_D2_.Create_DC7_(_250_v12, _251_v13, p0)
		var _hi0 _dafny.Int = p0
		_ = _hi0
		for _253_i1 := Companion_Default___.SafeModuloInt(p0, (_252_v14).Dtor_cf8()); _253_i1.Cmp(_hi0) < 0; _253_i1 = _253_i1.Plus(_dafny.One) {
			var _254_v15 _dafny.Sequence
			_ = _254_v15
			_254_v15 = _dafny.UnicodeSeqOfUtf8Bytes("gmpwl")
			_254_v15 = _dafny.UnicodeSeqOfUtf8Bytes("ahtrhyr")
			var _255_v16 _dafny.MultiSet
			_ = _255_v16
			_255_v16 = _dafny.MultiSetOf(p0, p0)
			(globalState).F0 = (func() _dafny.Int {
				if p1 {
					return _253_i1
				}
				return (_255_v16).Cardinality()
			})()
			var _256_v17 _dafny.Sequence
			_ = _256_v17
			_256_v17 = _dafny.SeqOf(p2)
			var _257_v18 _dafny.Set
			_ = _257_v18
			_257_v18 = _dafny.SetOf(p1)
			_256_v17 = _dafny.SeqOf(((_this).F12()) && ((_this).Fm1(p2, p0, _253_i1, globalState)), (_this).F6(), !(_257_v18).Contains(p2))
			var _258_v19 bool
			_ = _258_v19
			_258_v19 = true
			var _259_v20 _dafny.Sequence
			_ = _259_v20
			_259_v20 = _dafny.SeqOf(_251_v13, (_252_v14).Dtor_cf7())
			_258_v19 = !_dafny.Companion_Sequence_.Contains(_259_v20, _251_v13)
		}
		(globalState).F0 = p0
		var _260_v21 _dafny.Array
		_ = _260_v21
		var _len0_6 _dafny.Int = _dafny.IntOfInt64(13)
		_ = _len0_6
		var _nw39 _dafny.Array
		_ = _nw39
		if _len0_6.Cmp(_dafny.Zero) == 0 {
			_nw39 = _dafny.NewArray(_len0_6)
		} else {
			var _init6 func(_dafny.Int) _dafny.Int = (func(_261_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
				return func(_262_i3 _dafny.Int) _dafny.Int {
					return (_262_i3).Minus((_dafny.Zero).Minus(_261_p0))
				}
			})(p0)
			_ = _init6
			var _element0_6 = _init6(_dafny.Zero)
			_ = _element0_6
			_nw39 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
			(_nw39).ArraySet1(_element0_6, 0)
			var _nativeLen0_6 = (_len0_6).Int()
			_ = _nativeLen0_6
			for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
				(_nw39).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
			}
		}
		_260_v21 = _nw39
		for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_260_v21), 0))); ; {
			_guard_loop_1, _ok22 := _iter22()
			if !_ok22 {
				break
			}
			var _263_i2 _dafny.Int
			_263_i2 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_263_i2).Sign() != -1) && ((_263_i2).Cmp(_dafny.ArrayLen((_260_v21), 0)) < 0)) {
				(_260_v21).ArraySet1(Companion_Default___.SafeModuloInt(_263_i2, p0), (_263_i2).Int())
			}
		}
		var _264_v22 D5
		_ = _264_v22
		_264_v22 = Companion_D5_.Create_DC15_((_this).F13())
		var _265_v23 _dafny.MultiSet
		_ = _265_v23
		_265_v23 = _dafny.MultiSetOf(_250_v12, _dafny.CodePoint('q'))
		var _266_v24 _dafny.Map
		_ = _266_v24
		_266_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_264_v22).Dtor_cf18(), _265_v23)
		_266_v24 = (_266_v24).Update(_dafny.Companion_Sequence_.Concatenate((_this).F13(), (_this).F13()), _265_v23)
		for _iter23 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_260_v21), 0))); ; {
			_guard_loop_2, _ok23 := _iter23()
			if !_ok23 {
				break
			}
			var _267_i4 _dafny.Int
			_267_i4 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_267_i4).Sign() != -1) && ((_267_i4).Cmp(_dafny.ArrayLen((_260_v21), 0)) < 0)) {
				(_260_v21).ArraySet1((_267_i4).Minus(p0), (_267_i4).Int())
			}
		}
	}
}
func (_this *C1) M7(p0 _dafny.Int, p1 _dafny.MultiSet, p2 _dafny.Int, globalState *GlobalState) {
	{
		var _268_v0 _dafny.Map
		_ = _268_v0
		_268_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_this).F12())).Cardinality()), (_dafny.Zero).Minus(p2))
		var _269_v1 _dafny.Sequence
		_ = _269_v1
		_269_v1 = _dafny.SeqOf(p2)
		var _270_v2 _dafny.Set
		_ = _270_v2
		_270_v2 = _dafny.SetOf((_this).F6(), !((_this).F12()))
		var _271_v3 _dafny.Map
		_ = _271_v3
		_271_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm14(_268_v0, (_this).F12(), (_269_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F13()).Cardinality()), _dafny.IntOfUint32((_269_v1).Cardinality()))).Uint32()).(_dafny.Int), (_this).F12(), globalState)).Cardinality(), (_270_v2).IsSubsetOf(_270_v2))
		var _272_v4 _dafny.Map
		_ = _272_v4
		_272_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm2((_this).F12(), globalState)), (_this).F12())
		_271_v3 = (_271_v3).Update((_dafny.Zero).Minus(((_272_v4).Cardinality()).Times(p2)), (_this).F6())
		var _273_v5 _dafny.Map
		_ = _273_v5
		_273_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F13(), (func() bool {
			if (_271_v3).Contains(p2) {
				return (_271_v3).Get(p2).(bool)
			}
			return (_this).F6()
		})())
		var _274_i0 _dafny.Int
		_ = _274_i0
		_274_i0 = _dafny.Zero
		{
			for ((_dafny.IntOfUint32(((_this).F13()).Cardinality())).Minus((_this).Fm4(_dafny.IntOfInt64(19), _dafny.IntOfUint32(((_this).F13()).Cardinality()), _273_v5, p0, globalState))).Cmp(p0) > 0 {
				{
					if (_274_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_274_i0 = (_274_i0).Plus(_dafny.One)
					var _275_v6 _dafny.Map
					_ = _275_v6
					_275_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_268_v0, !((_this).F12()))
					var _276_v7 *C0
					_ = _276_v7
					var _nw40 *C0 = New_C0_()
					_ = _nw40
					_nw40.Ctor__(_275_v6)
					_276_v7 = _nw40
					var _277_v8 D2
					_ = _277_v8
					_277_v8 = Companion_D2_.Create_DC7_(_dafny.CodePoint('m'), _268_v0, p2)
					var _278_v9 _dafny.MultiSet
					_ = _278_v9
					_278_v9 = _dafny.MultiSetOf(_277_v8)
					var _279_v10 _dafny.CodePoint
					_ = _279_v10
					_279_v10 = _dafny.CodePoint('i')
					var _pat_let_tv1 = _270_v2
					_ = _pat_let_tv1
					var _pat_let_tv2 = p0
					_ = _pat_let_tv2
					var _pat_let_tv3 = p1
					_ = _pat_let_tv3
					var _pat_let_tv4 = p2
					_ = _pat_let_tv4
					var _pat_let_tv5 = _268_v0
					_ = _pat_let_tv5
					var _pat_let_tv6 = _268_v0
					_ = _pat_let_tv6
					var _pat_let_tv7 = _270_v2
					_ = _pat_let_tv7
					var _pat_let_tv8 = p0
					_ = _pat_let_tv8
					var _pat_let_tv9 = p1
					_ = _pat_let_tv9
					var _pat_let_tv10 = p2
					_ = _pat_let_tv10
					var _pat_let_tv11 = _268_v0
					_ = _pat_let_tv11
					var _pat_let_tv12 = _268_v0
					_ = _pat_let_tv12
					(globalState).F0 = (func() _dafny.Int {
						if (_278_v9).Contains(func(_pat_let2_0 D2) D2 {
							return func(_286_dt__update__tmp_h2 D2) D2 {
								return func(_pat_let9_0 _dafny.Map) D2 {
									return func(_287_dt__update_hcf7_h2 _dafny.Map) D2 {
										return Companion_D2_.Create_DC7_((_286_dt__update__tmp_h2).Dtor_cf6(), _287_dt__update_hcf7_h2, (_286_dt__update__tmp_h2).Dtor_cf8())
									}(_pat_let9_0)
								}(_pat_let_tv6)
							}(_pat_let2_0)
						}(func(_pat_let3_0 D2) D2 {
							return func(_283_dt__update__tmp_h1 D2) D2 {
								return func(_pat_let7_0 _dafny.Int) D2 {
									return func(_284_dt__update_hcf8_h1 _dafny.Int) D2 {
										return func(_pat_let8_0 _dafny.Map) D2 {
											return func(_285_dt__update_hcf7_h1 _dafny.Map) D2 {
												return Companion_D2_.Create_DC7_((_283_dt__update__tmp_h1).Dtor_cf6(), _285_dt__update_hcf7_h1, _284_dt__update_hcf8_h1)
											}(_pat_let8_0)
										}(_pat_let_tv5)
									}(_pat_let7_0)
								}(_pat_let_tv4)
							}(_pat_let3_0)
						}(func(_pat_let4_0 D2) D2 {
							return func(_280_dt__update__tmp_h0 D2) D2 {
								return func(_pat_let5_0 _dafny.Int) D2 {
									return func(_281_dt__update_hcf8_h0 _dafny.Int) D2 {
										return func(_pat_let6_0 _dafny.Map) D2 {
											return func(_282_dt__update_hcf7_h0 _dafny.Map) D2 {
												return Companion_D2_.Create_DC7_((_280_dt__update__tmp_h0).Dtor_cf6(), _282_dt__update_hcf7_h0, _281_dt__update_hcf8_h0)
											}(_pat_let6_0)
										}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_pat_let_tv2), (_pat_let_tv3).Cardinality()))
									}(_pat_let5_0)
								}((_pat_let_tv1).Cardinality())
							}(_pat_let4_0)
						}(Companion_D2_.Create_DC7_(_279_v10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), p2))))) {
							return (_278_v9).Multiplicity(func(_pat_let10_0 D2) D2 {
								return func(_294_dt__update__tmp_h5 D2) D2 {
									return func(_pat_let17_0 _dafny.Map) D2 {
										return func(_295_dt__update_hcf7_h5 _dafny.Map) D2 {
											return Companion_D2_.Create_DC7_((_294_dt__update__tmp_h5).Dtor_cf6(), _295_dt__update_hcf7_h5, (_294_dt__update__tmp_h5).Dtor_cf8())
										}(_pat_let17_0)
									}(_pat_let_tv12)
								}(_pat_let10_0)
							}(func(_pat_let11_0 D2) D2 {
								return func(_291_dt__update__tmp_h4 D2) D2 {
									return func(_pat_let15_0 _dafny.Int) D2 {
										return func(_292_dt__update_hcf8_h3 _dafny.Int) D2 {
											return func(_pat_let16_0 _dafny.Map) D2 {
												return func(_293_dt__update_hcf7_h4 _dafny.Map) D2 {
													return Companion_D2_.Create_DC7_((_291_dt__update__tmp_h4).Dtor_cf6(), _293_dt__update_hcf7_h4, _292_dt__update_hcf8_h3)
												}(_pat_let16_0)
											}(_pat_let_tv11)
										}(_pat_let15_0)
									}(_pat_let_tv10)
								}(_pat_let11_0)
							}(func(_pat_let12_0 D2) D2 {
								return func(_288_dt__update__tmp_h3 D2) D2 {
									return func(_pat_let13_0 _dafny.Int) D2 {
										return func(_289_dt__update_hcf8_h2 _dafny.Int) D2 {
											return func(_pat_let14_0 _dafny.Map) D2 {
												return func(_290_dt__update_hcf7_h3 _dafny.Map) D2 {
													return Companion_D2_.Create_DC7_((_288_dt__update__tmp_h3).Dtor_cf6(), _290_dt__update_hcf7_h3, _289_dt__update_hcf8_h2)
												}(_pat_let14_0)
											}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_pat_let_tv8), (_pat_let_tv9).Cardinality()))
										}(_pat_let13_0)
									}((_pat_let_tv7).Cardinality())
								}(_pat_let12_0)
							}(Companion_D2_.Create_DC7_(_279_v10, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2), p2)))))
						}
						return (_dafny.Zero).Minus(p0)
					})()
					var _296_v11 bool
					_ = _296_v11
					_296_v11 = false
					_296_v11 = (p2).Cmp(p0) < 0
					var _297_v12 D3
					_ = _297_v12
					_297_v12 = Companion_D3_.Create_DC10_((_this).F6(), p0)
					(globalState).F0 = (_297_v12).Dtor_cf12()
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		(globalState).F0 = (_dafny.Zero).Minus(p0)
		var _298_v13 _dafny.Sequence
		_ = _298_v13
		_298_v13 = _dafny.SeqOf(p1, Companion_Default___.Fm15((_this).F6(), globalState), (_dafny.MultiSetFromSeq(_269_v1)).Union(p1))
		var _299_v14 _dafny.Array
		_ = _299_v14
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(20)
		_ = _len0_7
		var _nw41 _dafny.Array
		_ = _nw41
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw41 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) _dafny.Sequence = func(_300_i1 _dafny.Int) _dafny.Sequence {
				return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-706))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg23 _dafny.Int) interface{} {
						return coer23(arg23)
					}
				}(func(_301_i2 _dafny.Int) _dafny.Sequence {
					return (_this).F13()
				}))
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw41 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw41).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw41).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_299_v14 = _nw41
		var _302_v15 _dafny.Sequence
		_ = _302_v15
		_302_v15 = _dafny.SeqOf((_this).F13())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_299_v14), 0))
		_ = _index50
		(_299_v14).ArraySet1(_302_v15, (_index50).Int())
		var _303_v16 _dafny.Array
		_ = _303_v16
		var _nw42 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(18))
		_ = _nw42
		_303_v16 = _nw42
		var _304_v17 _dafny.Sequence
		_ = _304_v17
		_304_v17 = _dafny.SeqOf((_this).F12(), (_this).F12())
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_299_v14), 0))
		_ = _index51
		var _rhs39 _dafny.Sequence = _298_v13
		_ = _rhs39
		var _rhs40 _dafny.Int = p0
		_ = _rhs40
		var _rhs41 _dafny.Array = _303_v16
		_ = _rhs41
		var _rhs42 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(362))).Uint32(), func(coer24 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
			return func(arg24 _dafny.Int) interface{} {
				return coer24(arg24)
			}
		}(func(_305_i3 _dafny.Int) _dafny.Sequence {
			return (_this).F13()
		}))
		_ = _rhs42
		var _rhs43 _dafny.Int = ((func() _dafny.Int {
			if (_this).Fm2((_this).F6(), globalState) {
				return p2
			}
			return p0
		})()).Minus((func() _dafny.Int {
			if (_this).F12() {
				return _dafny.IntOfUint32((_304_v17).Cardinality())
			}
			return p0
		})())
		_ = _rhs43
		var _lhs28 *GlobalState = globalState
		_ = _lhs28
		var _lhs29 *GlobalState = globalState
		_ = _lhs29
		var _lhs30 _dafny.Array = _299_v14
		_ = _lhs30
		var _lhs31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(639), _dafny.ArrayLen((_299_v14), 0))
		_ = _lhs31
		var _lhs32 *GlobalState = globalState
		_ = _lhs32
		_298_v13 = _rhs39
		_lhs28.F0 = _rhs40
		_lhs29.F2 = _rhs41
		(_lhs30).ArraySet1(_rhs42, (_lhs31).Int())
		_lhs32.F0 = _rhs43
		var _306_v18 bool
		_ = _306_v18
		_306_v18 = false
		var _307_v19 _dafny.Sequence
		_ = _307_v19
		_307_v19 = _dafny.UnicodeSeqOfUtf8Bytes("gvyhp")
		var _rhs44 bool = ((_dafny.Zero).Minus((func() _dafny.Int {
			if (_this).F12() {
				return _dafny.IntOfUint32((_269_v1).Cardinality())
			}
			return (p1).Cardinality()
		})())).Cmp(_dafny.IntOfInt64(60)) < 0
		_ = _rhs44
		var _rhs45 _dafny.Sequence = _307_v19
		_ = _rhs45
		_306_v18 = _rhs44
		_307_v19 = _rhs45
		(globalState).F0 = (_dafny.Zero).Minus(p2)
	}
}
func (_this *C1) M8(p0 bool, globalState *GlobalState) (bool, _dafny.Int, bool, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var r3 bool = false
		_ = r3
		var _308_v0 _dafny.Int
		_ = _308_v0
		_308_v0 = _dafny.IntOfInt64(375)
		var _309_v1 _dafny.Map
		_ = _309_v1
		_309_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_308_v0, _dafny.IntOfInt64(110))
		var _310_v2 _dafny.Map
		_ = _310_v2
		_310_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_309_v1, (_this).F6())
		var _311_v3 *C0
		_ = _311_v3
		var _nw43 *C0 = New_C0_()
		_ = _nw43
		_nw43.Ctor__(_310_v2)
		_311_v3 = _nw43
		var _312_v4 _dafny.Sequence
		_ = _312_v4
		_312_v4 = _dafny.SeqOf(!((_this).F6()), (_this).F6())
		var _313_v5 _dafny.Map
		_ = _313_v5
		_313_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_311_v3, _dafny.IntOfUint32((_312_v4).Cardinality()))
		_313_v5 = (_313_v5).Update(_311_v3, Companion_Default___.SafeModuloInt(_308_v0, _308_v0))
		var _314_v6 _dafny.Sequence
		_ = _314_v6
		_314_v6 = _dafny.SeqOf(_308_v0)
		var _315_v7 _dafny.Map
		_ = _315_v7
		_315_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, !((_this).F12()))
		var _316_v8 _dafny.Map
		_ = _316_v8
		_316_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_314_v6).Select((Companion_Default___.SafeIndex(_308_v0, _dafny.IntOfUint32((_314_v6).Cardinality()))).Uint32()).(_dafny.Int), (_315_v7).Merge(_315_v7))
		_316_v8 = (_316_v8).Update((_308_v0).Minus((_dafny.Zero).Minus(_308_v0)), _315_v7)
		r2 = !(p0) || ((_this).Fm3(globalState))
		var _317_i0 _dafny.Int
		_ = _317_i0
		_317_i0 = _dafny.Zero
		{
			for (_this).F6() {
				{
					if (_317_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_317_i0 = (_317_i0).Plus(_dafny.One)
					_308_v0 = _308_v0
					var _318_v9 _dafny.Array
					_ = _318_v9
					var _len0_8 _dafny.Int = _dafny.IntOfInt64(25)
					_ = _len0_8
					var _nw44 _dafny.Array
					_ = _nw44
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw44 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) bool = func(_319_i1 _dafny.Int) bool {
							return (_this).F12()
						}
						_ = _init8
						var _element0_8 = _init8(_dafny.Zero)
						_ = _element0_8
						_nw44 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
						(_nw44).ArraySet1(_element0_8, 0)
						var _nativeLen0_8 = (_len0_8).Int()
						_ = _nativeLen0_8
						for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
							(_nw44).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
						}
					}
					_318_v9 = _nw44
					_318_v9 = _318_v9
					_311_v3 = _311_v3
					var _320_v10 _dafny.Set
					_ = _320_v10
					_320_v10 = _dafny.SetOf(_308_v0)
					var _321_v11 D6
					_ = _321_v11
					_321_v11 = Companion_D6_.Create_DC18_(_311_v3, _312_v4, (_320_v10).Cardinality())
					_315_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F12(), (_dafny.MultiSetFromSeq(_312_v4)).Equals(_dafny.MultiSetFromSeq((_321_v11).Dtor_cf26())))
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _322_v12 _dafny.Array
		_ = _322_v12
		var _nwElement0_5 _dafny.Int = _308_v0
		_ = _nwElement0_5
		var _nw45 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(7))
		_ = _nw45
		(_nw45).ArraySet1(_nwElement0_5, 0)
		(_nw45).ArraySet1(_308_v0, 1)
		(_nw45).ArraySet1(_308_v0, 2)
		(_nw45).ArraySet1(_dafny.IntOfInt64(686), 3)
		(_nw45).ArraySet1((func() _dafny.Int {
			if (_309_v1).Contains(_308_v0) {
				return (_309_v1).Get(_308_v0).(_dafny.Int)
			}
			return _308_v0
		})(), 4)
		(_nw45).ArraySet1(_308_v0, 5)
		(_nw45).ArraySet1(_308_v0, 6)
		_322_v12 = _nw45
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_322_v12), 0))
		_ = _index52
		(_322_v12).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_314_v6, (Companion_Default___.SafeIndex(_308_v0, _dafny.IntOfUint32((_314_v6).Cardinality()))).Uint32(), _308_v0)).Cardinality()), (_index52).Int())
		var _323_v13 _dafny.Map
		_ = _323_v13
		_323_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_312_v4).Select((Companion_Default___.SafeIndex(_308_v0, _dafny.IntOfUint32((_312_v4).Cardinality()))).Uint32()).(bool), _dafny.IntOfUint32((_314_v6).Cardinality()))
		var _324_v14 _dafny.Set
		_ = _324_v14
		_324_v14 = _dafny.SetOf(!((_this).F6()), (_this).F6(), (_this).F12(), (_312_v4).Select((Companion_Default___.SafeIndex(((_323_v13).Update((_this).F6(), _308_v0)).Cardinality(), _dafny.IntOfUint32((_312_v4).Cardinality()))).Uint32()).(bool))
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(643), _dafny.ArrayLen((_322_v12), 0))
		_ = _index53
		(_322_v12).ArraySet1((func() _dafny.Int {
			if p0 {
				return (_308_v0).Times((_324_v14).Cardinality())
			}
			return _308_v0
		})(), (_index53).Int())
		var _325_i2 _dafny.Int
		_ = _325_i2
		_325_i2 = _dafny.Zero
		{
			for p0 {
				{
					if (_325_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_325_i2 = (_325_i2).Plus(_dafny.One)
					var _326_v15 *C0
					_ = _326_v15
					var _nw46 *C0 = New_C0_()
					_ = _nw46
					_nw46.Ctor__(((_311_v3).F11()).Merge(((_311_v3).F11()).Update(_309_v1, (_this).F12())))
					_326_v15 = _nw46
					r2 = (_this).F12()
					var _327_v16 *C0
					_ = _327_v16
					var _nw47 *C0 = New_C0_()
					_ = _nw47
					_nw47.Ctor__((_311_v3).F11())
					_327_v16 = _nw47
					var _328_v17 _dafny.Array
					_ = _328_v17
					var _nw48 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
					_ = _nw48
					_328_v17 = _nw48
					var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_328_v17), 0))
					_ = _index54
					(_328_v17).ArraySet1((func() bool {
						if (_315_v7).Contains(p0) {
							return (_315_v7).Get(p0).(bool)
						}
						return (_this).F12()
					})(), (_index54).Int())
					var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_328_v17), 0))
					_ = _index55
					(_328_v17).ArraySet1((_this).F12(), (_index55).Int())
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		r0 = !((_this).F12()) || ((_this).F12())
		r1 = (_308_v0).Times(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(_308_v0), _dafny.IntOfInt64(585)))
		r2 = true
		r3 = (_this).Fm2(p0, globalState)
		return r0, r1, r2, r3
	}
}
func (_this *C1) F12() bool {
	{
		return _this._f12
	}
}
func (_this *C1) F13() _dafny.Sequence {
	{
		return _this._f13
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	_f5 _dafny.MultiSet
}

func New_C2_() *C2 {
	_this := C2{}

	_this._f5 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C2{}
var _ _dafny.TraitOffspring = &C2{}

func (_this *C2) F5() _dafny.MultiSet {
	return _this._f5
}
func (_this *C2) Ctor__(f5 _dafny.MultiSet) {
	{
		(_this)._f5 = f5
	}
}
func (_this *C2) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return !(true) || (true)
	}
}
func (_this *C2) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return ((_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("uvl"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(81))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg25 _dafny.Int) interface{} {
				return coer25(arg25)
			}
		}(func(_329_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})), _dafny.UnicodeSeqOfUtf8Bytes("hxubpnhri"))).Cardinality()).Cmp(_dafny.IntOfInt64(211)) > 0
	}
}
func (_this *C2) M9(p0 bool, p1 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _330_v0 _dafny.CodePoint
		_ = _330_v0
		_330_v0 = _dafny.CodePoint('g')
		var _331_v1 _dafny.Sequence
		_ = _331_v1
		_331_v1 = _dafny.SeqOf(p0)
		var _332_v2 _dafny.Map
		_ = _332_v2
		_332_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_331_v1).Cardinality()))
		var _333_v3 _dafny.Map
		_ = _333_v3
		_333_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _332_v2)
		_330_v0 = Companion_Default___.Fm16(p0, (func() _dafny.Map {
			if (_333_v3).Contains(p0) {
				return (_333_v3).Get(p0).(_dafny.Map)
			}
			return _332_v2
		})(), Companion_Default___.Fm17(globalState), globalState)
		var _334_v4 _dafny.Sequence
		_ = _334_v4
		_334_v4 = _dafny.SeqOf(p1)
		_334_v4 = Companion_Default___.Fm18(globalState)
		var _335_v5 _dafny.Map
		_ = _335_v5
		_335_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)
		var _336_v6 _dafny.Map
		_ = _336_v6
		_336_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p0), p0)
		var _337_v7 _dafny.Sequence
		_ = _337_v7
		_337_v7 = _dafny.UnicodeSeqOfUtf8Bytes("ibtnyjf")
		var _338_v8 D7
		_ = _338_v8
		_338_v8 = Companion_D7_.Create_DC19_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1))
		var _339_v9 _dafny.Array
		_ = _339_v9
		var _nwElement0_6 _dafny.Map = (func() _dafny.Map {
			if p0 {
				return _335_v5
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1)
		})()
		_ = _nwElement0_6
		var _nw49 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(23))
		_ = _nw49
		(_nw49).ArraySet1(_nwElement0_6, 0)
		(_nw49).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1), 1)
		(_nw49).ArraySet1(_335_v5, 2)
		(_nw49).ArraySet1(_335_v5, 3)
		(_nw49).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1), 4)
		(_nw49).ArraySet1((_335_v5).Merge(_335_v5), 5)
		(_nw49).ArraySet1((_335_v5).Merge(_335_v5), 6)
		(_nw49).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p1), 7)
		(_nw49).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() bool {
			if (_336_v6).Contains(p0) {
				return (_336_v6).Get(p0).(bool)
			}
			return p0
		})(), p1), 8)
		(_nw49).ArraySet1(Companion_Default___.Fm19((Companion_D3_.Create_DC10_(false, _dafny.IntOfUint32((_337_v7).Cardinality()))).Dtor_cf12(), globalState), 9)
		(_nw49).ArraySet1((_335_v5).Merge(_335_v5), 10)
		(_nw49).ArraySet1((func() _dafny.Map {
			if p0 {
				return (_335_v5).Update(true, _dafny.IntOfUint32((_337_v7).Cardinality()))
			}
			return _335_v5
		})(), 11)
		(_nw49).ArraySet1(_335_v5, 12)
		(_nw49).ArraySet1(_335_v5, 13)
		(_nw49).ArraySet1((func() _dafny.Map {
			if p0 {
				return _335_v5
			}
			return _335_v5
		})(), 14)
		(_nw49).ArraySet1(((_338_v8).Dtor_cf28()).Merge(_335_v5), 15)
		(_nw49).ArraySet1(_335_v5, 16)
		(_nw49).ArraySet1(_335_v5, 17)
		(_nw49).ArraySet1((func() _dafny.Map {
			if !(true) {
				return (_335_v5).Update((_this).Fm2(false, globalState), p1)
			}
			return _335_v5
		})(), 18)
		(_nw49).ArraySet1(_335_v5, 19)
		(_nw49).ArraySet1(Companion_Default___.Fm19((_dafny.Zero).Minus(p1), globalState), 20)
		(_nw49).ArraySet1((_335_v5).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p1)), 21)
		(_nw49).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p1)).Merge(_335_v5), 22)
		_339_v9 = _nw49
		var _340_v10 _dafny.Array
		_ = _340_v10
		var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(18))
		_ = _nw50
		_340_v10 = _nw50
		var _341_v11 bool
		_ = _341_v11
		_341_v11 = true
		var _342_v12 _dafny.Set
		_ = _342_v12
		_342_v12 = _dafny.SetOf(p1, (_336_v6).Cardinality())
		var _343_v13 _dafny.Map
		_ = _343_v13
		_343_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _342_v12)
		var _344_v14 D5
		_ = _344_v14
		_344_v14 = Companion_D5_.Create_DC15_(_337_v7)
		var _345_v15 _dafny.MultiSet
		_ = _345_v15
		_345_v15 = _dafny.MultiSetOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(674))).Uint32(), func(coer26 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg26 _dafny.Int) interface{} {
				return coer26(arg26)
			}
		}((func(_346_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_347_i0 _dafny.Int) _dafny.CodePoint {
				return _346_v0
			}
		})(_330_v0))), _337_v7)
		var _rhs46 _dafny.Array = _339_v9
		_ = _rhs46
		var _rhs47 _dafny.Array = _340_v10
		_ = _rhs47
		var _rhs48 bool = (func() bool {
			if (_336_v6).Contains((_342_v12).IsDisjointFrom((func() _dafny.Set {
				if (_343_v13).Contains(p1) {
					return (_343_v13).Get(p1).(_dafny.Set)
				}
				return _342_v12
			})())) {
				return (_336_v6).Get((_342_v12).IsDisjointFrom((func() _dafny.Set {
					if (_343_v13).Contains(p1) {
						return (_343_v13).Get(p1).(_dafny.Set)
					}
					return _342_v12
				})())).(bool)
			}
			return _dafny.Companion_Sequence_.IsPrefixOf(_337_v7, (_344_v14).Dtor_cf18())
		})()
		_ = _rhs48
		var _rhs49 bool = ((_345_v15).IsProperSubsetOf(_345_v15)) && (false)
		_ = _rhs49
		_339_v9 = _rhs46
		_340_v10 = _rhs47
		_341_v11 = _rhs48
		_341_v11 = _rhs49
		var _348_v16 _dafny.Sequence
		_ = _348_v16
		_348_v16 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(165))).Uint32(), func(coer27 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg27 _dafny.Int) interface{} {
				return coer27(arg27)
			}
		}((func(_349_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_350_i1 _dafny.Int) _dafny.CodePoint {
				return _349_v0
			}
		})(_330_v0))))
		var _351_v17 _dafny.Array
		_ = _351_v17
		var _nw51 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
		_ = _nw51
		_351_v17 = _nw51
		var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))
		_ = _index56
		(_351_v17).ArraySet1(p1, (_index56).Int())
		var _352_v19 D8
		_ = _352_v19
		_352_v19 = Companion_D8_.Create_DC23_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqOf(_341_v11), _dafny.IntOfUint32((_334_v4).Cardinality())))
		var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))
		_ = _index57
		var _rhs50 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_348_v16, _dafny.Companion_Sequence_.Concatenate(_348_v16, _348_v16))
		_ = _rhs50
		var _rhs51 _dafny.Int = (p1).Minus(p1)
		_ = _rhs51
		var _rhs52 _dafny.CodePoint = Companion_Default___.Fm16(true, (func() _dafny.Map {
			var _coll21 = _dafny.NewMapBuilder()
			_ = _coll21
			for _iter24 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(366), _dafny.IntOfInt64(185))); ; {
				_compr_21, _ok24 := _iter24()
				if !_ok24 {
					break
				}
				var _353_v18 _dafny.Int
				_353_v18 = interface{}(_compr_21).(_dafny.Int)
				if ((_dafny.IntOfInt64(366)).Cmp(_353_v18) <= 0) && ((_353_v18).Cmp(_dafny.IntOfInt64(185)) < 0) {
					_coll21.Add((_353_v18).Plus((_dafny.MultiSetOf(_330_v0)).Cardinality()), p1)
				}
			}
			return _coll21.ToMap()
		}()).Update(((_352_v19).Dtor_cf33()).Cardinality(), p1), _dafny.Companion_Sequence_.Concatenate(_337_v7, _337_v7), globalState)
		_ = _rhs52
		var _rhs53 _dafny.Int = p1
		_ = _rhs53
		var _rhs54 _dafny.Int = p1
		_ = _rhs54
		var _lhs33 *GlobalState = globalState
		_ = _lhs33
		var _lhs34 _dafny.Array = _351_v17
		_ = _lhs34
		var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))
		_ = _lhs35
		var _lhs36 *GlobalState = globalState
		_ = _lhs36
		_348_v16 = _rhs50
		_lhs33.F0 = _rhs51
		_330_v0 = _rhs52
		(_lhs34).ArraySet1(_rhs53, (_lhs35).Int())
		_lhs36.F0 = _rhs54
		var _354_v20 _dafny.Array
		_ = _354_v20
		var _len0_9 _dafny.Int = _dafny.IntOfInt64(8)
		_ = _len0_9
		var _nw52 _dafny.Array
		_ = _nw52
		if _len0_9.Cmp(_dafny.Zero) == 0 {
			_nw52 = _dafny.NewArray(_len0_9)
		} else {
			var _init9 func(_dafny.Int) _dafny.Sequence = (func(_355_v7 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
				return func(_356_i2 _dafny.Int) _dafny.Sequence {
					return _355_v7
				}
			})(_337_v7)
			_ = _init9
			var _element0_9 = _init9(_dafny.Zero)
			_ = _element0_9
			_nw52 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
			(_nw52).ArraySet1(_element0_9, 0)
			var _nativeLen0_9 = (_len0_9).Int()
			_ = _nativeLen0_9
			for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
				(_nw52).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
			}
		}
		_354_v20 = _nw52
		var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_354_v20), 0))
		_ = _index58
		(_354_v20).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(333))).Uint32(), func(coer28 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg28 _dafny.Int) interface{} {
				return coer28(arg28)
			}
		}((func(_357_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_358_i3 _dafny.Int) _dafny.CodePoint {
				return _357_v0
			}
		})(_330_v0))), (_index58).Int())
		var _359_v21 D1
		_ = _359_v21
		_359_v21 = Companion_D1_.Create_DC2_(_341_v11)
		var _360_v22 _dafny.Map
		_ = _360_v22
		_360_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_330_v0, _351_v17)
		var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))
		_ = _index59
		var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_354_v20), 0))
		_ = _index60
		var _rhs55 bool = p0
		_ = _rhs55
		var _rhs56 _dafny.Int = (_351_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))).Int()).(_dafny.Int)
		_ = _rhs56
		var _rhs57 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_337_v7, (Companion_Default___.SafeIndex((_360_v22).Cardinality(), _dafny.IntOfUint32((_337_v7).Cardinality()))).Uint32(), Companion_Default___.Fm16(!(p0), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_351_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))).Int()).(_dafny.Int), (_351_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))).Int()).(_dafny.Int)), (_348_v16).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-325), _dafny.IntOfUint32((_348_v16).Cardinality()))).Uint32()).(_dafny.Sequence), globalState))
		_ = _rhs57
		var _rhs58 D1 = _359_v21
		_ = _rhs58
		var _lhs37 _dafny.Array = _351_v17
		_ = _lhs37
		var _lhs38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(736), _dafny.ArrayLen((_351_v17), 0))
		_ = _lhs38
		var _lhs39 _dafny.Array = _354_v20
		_ = _lhs39
		var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(150), _dafny.ArrayLen((_354_v20), 0))
		_ = _lhs40
		_341_v11 = _rhs55
		(_lhs37).ArraySet1(_rhs56, (_lhs38).Int())
		(_lhs39).ArraySet1(_rhs57, (_lhs40).Int())
		_359_v21 = _rhs58
		(globalState).F0 = p1
		var _361_v23 _dafny.Array
		_ = _361_v23
		var _nw53 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(19))
		_ = _nw53
		_361_v23 = _nw53
		var _362_v24 _dafny.Map
		_ = _362_v24
		_362_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_361_v23, p0)
		var _363_v25 _dafny.Sequence
		_ = _363_v25
		_363_v25 = _dafny.SeqOf(_362_v24, (_362_v24).Merge(_362_v24), _362_v24)
		r0 = (_363_v25).Select((Companion_Default___.SafeIndex((_342_v12).Cardinality(), _dafny.IntOfUint32((_363_v25).Cardinality()))).Uint32()).(_dafny.Map)
		return r0
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f6 bool
	_f5 _dafny.MultiSet
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f6 = false
	_this._f5 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F6() bool {
	return _this._f6
}
func (_this *C3) F5() _dafny.MultiSet {
	return _this._f5
}
func (_this *C3) Ctor__(f6 bool, f5 _dafny.MultiSet) {
	{
		(_this)._f6 = f6
		(_this)._f5 = f5
	}
}
func (_this *C3) Fm3(globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C3) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F6()), _dafny.SeqOf((_this).F6()))).Cardinality())).Times((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rassq")).Cardinality())).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("cax")).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(!((_this).F6()))).Cardinality())))).Cardinality()))
	}
}
func (_this *C3) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		var _source3 D0 = Companion_D0_.Create_DC0_(_dafny.IntOfInt64(-247))
		_ = _source3
		if _source3.Is_DC1() {
			return (_this).F6()
		} else {
			var _364___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
			_ = _364___mcc_h0
			var _365_cf0 _dafny.Int = _364___mcc_h0
			_ = _365_cf0
			return !((_this).F6())
		}
	}
}
func (_this *C3) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return (Companion_D1_.Create_DC2_((_this).F6())).Dtor_cf1()
	}
}
func (_this *C3) M1(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) {
	{
		var _366_v0 D0
		_ = _366_v0
		_366_v0 = Companion_D0_.Create_DC1_()
		_366_v0 = _366_v0
		var _367_v1 _dafny.Sequence
		_ = _367_v1
		_367_v1 = _dafny.UnicodeSeqOfUtf8Bytes("vekcmyumj")
		var _368_v2 D0
		_ = _368_v2
		_368_v2 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_367_v1).Cardinality()))
		var _pat_let_tv13 = p0
		_ = _pat_let_tv13
		var _source4 D0 = func(_pat_let18_0 D0) D0 {
			return func(_369_dt__update__tmp_h0 D0) D0 {
				return func(_pat_let19_0 _dafny.Int) D0 {
					return func(_370_dt__update_hcf0_h0 _dafny.Int) D0 {
						return Companion_D0_.Create_DC0_(_370_dt__update_hcf0_h0)
					}(_pat_let19_0)
				}(_pat_let_tv13)
			}(_pat_let18_0)
		}(_368_v2)
		_ = _source4
		if _source4.Is_DC1() {
			if p1 {
				var _371_v3 _dafny.Map
				_ = _371_v3
				_371_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
				var _372_v4 _dafny.CodePoint
				_ = _372_v4
				_372_v4 = _dafny.CodePoint('m')
				var _373_v5 _dafny.Map
				_ = _373_v5
				_373_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.IntOfInt64(-887)).Cmp((_371_v3).Cardinality()) <= 0, _372_v4)
				_373_v5 = (_373_v5).Update((p0).Cmp(p0) != 0, _372_v4)
				var _374_v6 D1
				_ = _374_v6
				_374_v6 = Companion_D1_.Create_DC2_((_this).F6())
				var _375_v7 _dafny.Map
				_ = _375_v7
				_375_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p0)
				var _376_v8 _dafny.Map
				_ = _376_v8
				_376_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_374_v6, _375_v7)
				var _377_v9 _dafny.Map
				_ = _377_v9
				_377_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(p0))
				(globalState).F0 = (_dafny.Zero).Minus((func() _dafny.Int {
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_376_v8).Cardinality(), p0)).Equals(_377_v9) {
						return p0
					}
					return p0
				})())
				var _378_v10 _dafny.Map
				_ = _378_v10
				_378_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_377_v9, p1)
				var _379_v11 *C0
				_ = _379_v11
				var _nw54 *C0 = New_C0_()
				_ = _nw54
				_nw54.Ctor__(_378_v10)
				_379_v11 = _nw54
				var _380_v12 _dafny.Map
				_ = _380_v12
				_380_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_379_v11, p0)
				_380_v12 = (_380_v12).Update(_379_v11, p0)
				var _381_v13 *C0
				_ = _381_v13
				var _nw55 *C0 = New_C0_()
				_ = _nw55
				_nw55.Ctor__(((_379_v11).F11()).Merge(_378_v10))
				_381_v13 = _nw55
				(globalState).F0 = Companion_Default___.SafeModuloInt(p0, p0)
			} else {
				var _382_v14 bool
				_ = _382_v14
				_382_v14 = true
				_382_v14 = !(!((_this).F6()))
				var _383_v15 _dafny.Array
				_ = _383_v15
				var _nw56 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(22))
				_ = _nw56
				_383_v15 = _nw56
				var _384_v16 _dafny.Map
				_ = _384_v16
				_384_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_383_v15), 0))
				_ = _index61
				(_383_v15).ArraySet1((_384_v16).Merge(_384_v16), (_index61).Int())
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(137), _dafny.ArrayLen((_383_v15), 0))
				_ = _index62
				(_383_v15).ArraySet1(_384_v16, (_index62).Int())
				var _385_v17 _dafny.Array
				_ = _385_v17
				var _nw57 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
				_ = _nw57
				_385_v17 = _nw57
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_385_v17), 0))
				_ = _index63
				(_385_v17).ArraySet1(_382_v14, (_index63).Int())
				var _386_v18 _dafny.Array
				_ = _386_v18
				var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(15))
				_ = _nw58
				_386_v18 = _nw58
				var _387_v19 _dafny.Map
				_ = _387_v19
				_387_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_386_v18, p2)
				var _388_v20 _dafny.Sequence
				_ = _388_v20
				_388_v20 = _dafny.SeqOf((_this).F6())
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_385_v17), 0))
				_ = _index64
				var _rhs59 bool = (_this).F6()
				_ = _rhs59
				var _rhs60 bool = !(!(_387_v19).Equals((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_386_v18, p1)).Merge(_387_v19)))
				_ = _rhs60
				var _rhs61 bool = !((_388_v20).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_388_v20).Cardinality()))).Uint32()).(bool))
				_ = _rhs61
				var _lhs41 _dafny.Array = _385_v17
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_385_v17), 0))
				_ = _lhs42
				(_lhs41).ArraySet1(_rhs59, (_lhs42).Int())
				_382_v14 = _rhs60
				_382_v14 = _rhs61
				var _389_v21 _dafny.Map
				_ = _389_v21
				_389_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, ((_385_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_385_v17), 0))).Int()).(bool)) || ((_this).F6()))
				_389_v21 = (_389_v21).Update(false, (_385_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_385_v17), 0))).Int()).(bool))
				(globalState).F0 = ((p0).Minus(p0)).Times((p0).Plus(_dafny.IntOfUint32((_388_v20).Cardinality())))
			}
			(globalState).F0 = p0
			(globalState).F0 = (_dafny.Zero).Minus(_dafny.IntOfInt64(-971))
			var _390_v22 bool
			_ = _390_v22
			_390_v22 = true
			var _391_v23 _dafny.Set
			_ = _391_v23
			_391_v23 = _dafny.SetOf(p0)
			var _392_v24 _dafny.Map
			_ = _392_v24
			_392_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _366_v0)
			var _393_v26 _dafny.Map
			_ = _393_v26
			_393_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
			var _rhs62 bool = !((_391_v23).IsDisjointFrom(func() _dafny.Set {
				var _coll22 = _dafny.NewBuilder()
				_ = _coll22
				for _iter25 := _dafny.Iterate((_392_v24).Keys().Elements()); ; {
					_compr_22, _ok25 := _iter25()
					if !_ok25 {
						break
					}
					var _394_v25 _dafny.Int
					_394_v25 = interface{}(_compr_22).(_dafny.Int)
					if (_392_v24).Contains(_394_v25) {
						_coll22.Add(Companion_Default___.SafeDivisionInt(_394_v25, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("i")).Cardinality())))
					}
				}
				return _coll22.ToSet()
			}()))
			_ = _rhs62
			var _rhs63 _dafny.Int = (Companion_Default___.SafeModuloInt((func() _dafny.Int {
				if (_393_v26).Contains(false) {
					return (_393_v26).Get(false).(_dafny.Int)
				}
				return p0
			})(), _dafny.IntOfInt64(-562))).Plus(p0)
			_ = _rhs63
			var _lhs43 *GlobalState = globalState
			_ = _lhs43
			_390_v22 = _rhs62
			_lhs43.F0 = _rhs63
		} else {
			var _395___mcc_h0 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
			_ = _395___mcc_h0
			var _396_cf0 _dafny.Int = _395___mcc_h0
			_ = _396_cf0
			var _397_v27 _dafny.MultiSet
			_ = _397_v27
			_397_v27 = _dafny.MultiSetOf(_396_cf0)
			_397_v27 = _397_v27
			var _398_v28 D1
			_ = _398_v28
			_398_v28 = Companion_D1_.Create_DC3_(_366_v0, Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(655), _396_cf0))
			var _399_v29 bool
			_ = _399_v29
			_399_v29 = true
			var _400_v30 _dafny.Array
			_ = _400_v30
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_10
			var _nw59 _dafny.Array
			_ = _nw59
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw59 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) bool = (func(_401_v29 bool) func(_dafny.Int) bool {
					return func(_402_i0 _dafny.Int) bool {
						return _401_v29
					}
				})(_399_v29)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw59 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw59).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw59).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_400_v30 = _nw59
			var _403_v31 _dafny.Map
			_ = _403_v31
			_403_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("ypkiarto"), (_this).F6())
			var _404_v32 _dafny.Sequence
			_ = _404_v32
			_404_v32 = _dafny.SeqOf(p0, _396_cf0)
			var _rhs64 D1 = _398_v28
			_ = _rhs64
			var _rhs65 _dafny.Int = (_this).Fm4(_396_cf0, p0, (_403_v31).Merge(_403_v31), (_404_v32).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_396_cf0), _dafny.IntOfUint32((_404_v32).Cardinality()))).Uint32()).(_dafny.Int), globalState)
			_ = _rhs65
			var _rhs66 _dafny.Int = _dafny.IntOfUint32((_367_v1).Cardinality())
			_ = _rhs66
			var _rhs67 bool = p2
			_ = _rhs67
			var _rhs68 _dafny.Array = _400_v30
			_ = _rhs68
			var _lhs44 *GlobalState = globalState
			_ = _lhs44
			var _lhs45 *GlobalState = globalState
			_ = _lhs45
			_398_v28 = _rhs64
			_lhs44.F0 = _rhs65
			_lhs45.F0 = _rhs66
			_399_v29 = _rhs67
			_400_v30 = _rhs68
			var _405_v33 _dafny.Map
			_ = _405_v33
			_405_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_404_v32, _404_v32)).Cardinality()), _396_cf0)
			var _406_v34 _dafny.Set
			_ = _406_v34
			_406_v34 = _dafny.SetOf((_dafny.Zero).Minus(p0))
			var _407_v35 _dafny.Sequence
			_ = _407_v35
			_407_v35 = _dafny.SeqOf(_406_v34)
			var _408_v40 D2
			_ = _408_v40
			_408_v40 = Companion_D2_.Create_DC4_(_dafny.SetOf(_dafny.IntOfUint32((_404_v32).Cardinality()), _dafny.IntOfInt64(-705)))
			var _409_v42 _dafny.Array
			_ = _409_v42
			var _nwElement0_7 _dafny.Set = _406_v34
			_ = _nwElement0_7
			var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(22))
			_ = _nw60
			(_nw60).ArraySet1(_nwElement0_7, 0)
			(_nw60).ArraySet1((_406_v34).Difference(_406_v34), 1)
			(_nw60).ArraySet1(_406_v34, 2)
			(_nw60).ArraySet1((_407_v35).Select((Companion_Default___.SafeIndex(_396_cf0, _dafny.IntOfUint32((_407_v35).Cardinality()))).Uint32()).(_dafny.Set), 3)
			(_nw60).ArraySet1(_406_v34, 4)
			(_nw60).ArraySet1(_406_v34, 5)
			(_nw60).ArraySet1(_dafny.SetOf(p0), 6)
			(_nw60).ArraySet1(func() _dafny.Set {
				var _coll23 = _dafny.NewBuilder()
				_ = _coll23
				for _iter26 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(753), _dafny.IntOfInt64(464))); ; {
					_compr_23, _ok26 := _iter26()
					if !_ok26 {
						break
					}
					var _410_v36 _dafny.Int
					_410_v36 = interface{}(_compr_23).(_dafny.Int)
					if ((_dafny.IntOfInt64(753)).Cmp(_410_v36) <= 0) && ((_410_v36).Cmp(_dafny.IntOfInt64(464)) < 0) {
						_coll23.Add((_410_v36).Plus(p0))
					}
				}
				return _coll23.ToSet()
			}(), 7)
			(_nw60).ArraySet1(_406_v34, 8)
			(_nw60).ArraySet1(_406_v34, 9)
			(_nw60).ArraySet1(func() _dafny.Set {
				var _coll24 = _dafny.NewBuilder()
				_ = _coll24
				for _iter27 := _dafny.Iterate((_397_v27).Elements()); ; {
					_compr_24, _ok27 := _iter27()
					if !_ok27 {
						break
					}
					var _411_v37 _dafny.Int
					_411_v37 = interface{}(_compr_24).(_dafny.Int)
					if (_397_v27).Contains(_411_v37) {
						_coll24.Add(Companion_Default___.SafeDivisionInt(_411_v37, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()), (_dafny.MultiSetOf(false)).Cardinality())).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(711), (func() _dafny.Map {
							var _coll25 = _dafny.NewMapBuilder()
							_ = _coll25
							for _iter28 := _dafny.Iterate((_dafny.SetOf(_dafny.IntOfInt64(113))).Elements()); ; {
								_compr_25, _ok28 := _iter28()
								if !_ok28 {
									break
								}
								var _412_v38 _dafny.Int
								_412_v38 = interface{}(_compr_25).(_dafny.Int)
								if (_dafny.SetOf(_dafny.IntOfInt64(113))).Contains(_412_v38) {
									_coll25.Add((_412_v38).Times(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("y")).Cardinality())), !(false))
								}
							}
							return _coll25.ToMap()
						}()).Cardinality()))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), false))).Cardinality()))).Cardinality()))
					}
				}
				return _coll24.ToSet()
			}(), 10)
			(_nw60).ArraySet1((_406_v34).Intersection(_406_v34), 11)
			(_nw60).ArraySet1(func() _dafny.Set {
				var _coll26 = _dafny.NewBuilder()
				_ = _coll26
				for _iter29 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(636), _dafny.IntOfInt64(285))); ; {
					_compr_26, _ok29 := _iter29()
					if !_ok29 {
						break
					}
					var _413_v39 _dafny.Int
					_413_v39 = interface{}(_compr_26).(_dafny.Int)
					if ((_dafny.IntOfInt64(636)).Cmp(_413_v39) <= 0) && ((_413_v39).Cmp(_dafny.IntOfInt64(285)) < 0) {
						_coll26.Add(Companion_Default___.SafeDivisionInt(_413_v39, _396_cf0))
					}
				}
				return _coll26.ToSet()
			}(), 12)
			(_nw60).ArraySet1((_406_v34).Intersection(_dafny.SetOf(p0)), 13)
			(_nw60).ArraySet1(_406_v34, 14)
			(_nw60).ArraySet1(_dafny.SetOf(p0), 15)
			(_nw60).ArraySet1(_406_v34, 16)
			(_nw60).ArraySet1((_408_v40).Dtor_cf4(), 17)
			(_nw60).ArraySet1(_406_v34, 18)
			(_nw60).ArraySet1(func() _dafny.Set {
				var _coll27 = _dafny.NewBuilder()
				_ = _coll27
				for _iter30 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-506), _dafny.IntOfInt64(339))); ; {
					_compr_27, _ok30 := _iter30()
					if !_ok30 {
						break
					}
					var _414_v41 _dafny.Int
					_414_v41 = interface{}(_compr_27).(_dafny.Int)
					if ((_dafny.IntOfInt64(-506)).Cmp(_414_v41) <= 0) && ((_414_v41).Cmp(_dafny.IntOfInt64(339)) < 0) {
						_coll27.Add((_414_v41).Times(_dafny.IntOfInt64(25)))
					}
				}
				return _coll27.ToSet()
			}(), 19)
			(_nw60).ArraySet1(_406_v34, 20)
			(_nw60).ArraySet1(_406_v34, 21)
			_409_v42 = _nw60
			var _rhs69 _dafny.Map = _405_v33
			_ = _rhs69
			var _rhs70 _dafny.Int = (_dafny.IntOfInt64(105)).Times((((_this).F5()).Union((_this).F5())).Cardinality())
			_ = _rhs70
			var _rhs71 _dafny.Array = _409_v42
			_ = _rhs71
			_405_v33 = _rhs69
			_396_cf0 = _rhs70
			_409_v42 = _rhs71
			var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_400_v30), 0))
			_ = _index65
			(_400_v30).ArraySet1(!((_this).Fm2(p2, globalState)), (_index65).Int())
			var _415_v43 _dafny.Array
			_ = _415_v43
			var _nw61 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(25))
			_ = _nw61
			_415_v43 = _nw61
			var _416_v44 _dafny.Array
			_ = _416_v44
			var _nw62 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(3))
			_ = _nw62
			_416_v44 = _nw62
			var _417_v45 D3
			_ = _417_v45
			_417_v45 = Companion_D3_.Create_DC9_(_416_v44)
			var _418_v46 _dafny.Array
			_ = _418_v46
			var _nwElement0_8 _dafny.Array = _416_v44
			_ = _nwElement0_8
			var _nw63 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(11))
			_ = _nw63
			(_nw63).ArraySet1(_nwElement0_8, 0)
			(_nw63).ArraySet1(_416_v44, 1)
			(_nw63).ArraySet1(_416_v44, 2)
			(_nw63).ArraySet1((_417_v45).Dtor_cf10(), 3)
			(_nw63).ArraySet1(_416_v44, 4)
			(_nw63).ArraySet1(_416_v44, 5)
			(_nw63).ArraySet1(_416_v44, 6)
			(_nw63).ArraySet1(_416_v44, 7)
			(_nw63).ArraySet1(_416_v44, 8)
			(_nw63).ArraySet1(_416_v44, 9)
			(_nw63).ArraySet1(_416_v44, 10)
			_418_v46 = _nw63
			var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_415_v43), 0))
			_ = _index66
			(_415_v43).ArraySet1(_418_v46, (_index66).Int())
			var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_416_v44), 0))
			_ = _index67
			(_416_v44).ArraySet1(p0, (_index67).Int())
			var _419_v47 _dafny.Array
			_ = _419_v47
			var _nw64 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw64
			_419_v47 = _nw64
			var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_419_v47), 0))
			_ = _index68
			(_419_v47).ArraySet1(_367_v1, (_index68).Int())
			var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_400_v30), 0))
			_ = _index69
			var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_415_v43), 0))
			_ = _index70
			var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_416_v44), 0))
			_ = _index71
			var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_419_v47), 0))
			_ = _index72
			var _rhs72 bool = _399_v29
			_ = _rhs72
			var _rhs73 _dafny.Array = _418_v46
			_ = _rhs73
			var _rhs74 _dafny.Int = (_396_cf0).Times(p0)
			_ = _rhs74
			var _rhs75 _dafny.Sequence = _367_v1
			_ = _rhs75
			var _rhs76 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs76
			var _lhs46 _dafny.Array = _400_v30
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(160), _dafny.ArrayLen((_400_v30), 0))
			_ = _lhs47
			var _lhs48 _dafny.Array = _415_v43
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(823), _dafny.ArrayLen((_415_v43), 0))
			_ = _lhs49
			var _lhs50 _dafny.Array = _416_v44
			_ = _lhs50
			var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(26), _dafny.ArrayLen((_416_v44), 0))
			_ = _lhs51
			var _lhs52 _dafny.Array = _419_v47
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(717), _dafny.ArrayLen((_419_v47), 0))
			_ = _lhs53
			var _lhs54 *GlobalState = globalState
			_ = _lhs54
			(_lhs46).ArraySet1(_rhs72, (_lhs47).Int())
			(_lhs48).ArraySet1(_rhs73, (_lhs49).Int())
			(_lhs50).ArraySet1(_rhs74, (_lhs51).Int())
			(_lhs52).ArraySet1(_rhs75, (_lhs53).Int())
			_lhs54.F0 = _rhs76
		}
		var _420_v48 _dafny.Array
		_ = _420_v48
		var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
		_ = _nw65
		_420_v48 = _nw65
		var _421_v49 _dafny.Map
		_ = _421_v49
		_421_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_367_v1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_367_v1).Cardinality()))).Uint32(), _dafny.CodePoint('r')), _367_v1), _420_v48)
		_421_v49 = (_421_v49).Update(_dafny.Companion_Sequence_.Concatenate(_367_v1, _367_v1), _420_v48)
		var _422_i1 _dafny.Int
		_ = _422_i1
		_422_i1 = _dafny.Zero
		{
			for true {
				{
					if (_422_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_422_i1 = (_422_i1).Plus(_dafny.One)
					var _423_v50 bool
					_ = _423_v50
					_423_v50 = false
					_423_v50 = !(((_this).F6()) || (_423_v50))
					var _424_v51 _dafny.Set
					_ = _424_v51
					_424_v51 = _dafny.SetOf((Companion_D2_.Create_DC5_((_this).F6())).Dtor_cf5())
					var _425_v52 _dafny.CodePoint
					_ = _425_v52
					_425_v52 = _dafny.CodePoint('m')
					(globalState).F0 = (func() _dafny.Int {
						if p1 {
							return ((_424_v51).Cardinality()).Minus((_dafny.SetOf(_425_v52, _425_v52, _425_v52)).Cardinality())
						}
						return Companion_Default___.SafeDivisionInt(p0, p0)
					})()
					var _426_v53 _dafny.Map
					_ = _426_v53
					_426_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_423_v50, (_this).Fm3(globalState))
					var _427_v54 _dafny.Set
					_ = _427_v54
					_427_v54 = _dafny.SetOf(_425_v52)
					var _428_v55 _dafny.Map
					_ = _428_v55
					_428_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_423_v50, (_this).F6()), _427_v54)
					_423_v50 = !(_428_v55).Contains(_426_v53)
					var _429_v56 _dafny.Map
					_ = _429_v56
					_429_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
					var _430_v57 _dafny.Map
					_ = _430_v57
					_430_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_429_v56, !(p2))
					var _431_v58 *C0
					_ = _431_v58
					var _nw66 *C0 = New_C0_()
					_ = _nw66
					_nw66.Ctor__(_430_v57)
					_431_v58 = _nw66
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _432_v59 _dafny.Array
		_ = _432_v59
		var _nw67 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
		_ = _nw67
		_432_v59 = _nw67
		var _433_v60 _dafny.Sequence
		_ = _433_v60
		_433_v60 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("f"))
		var _434_v61 _dafny.CodePoint
		_ = _434_v61
		_434_v61 = _dafny.CodePoint('b')
		var _435_v62 _dafny.Map
		_ = _435_v62
		_435_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_dafny.Zero).Minus(p0))
		var _436_v63 D2
		_ = _436_v63
		_436_v63 = Companion_D2_.Create_DC7_(_434_v61, _435_v62, p0)
		var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_432_v59), 0))
		_ = _index73
		(_432_v59).ArraySet1((_433_v60).Select((Companion_Default___.SafeIndex((_this).Fm4((_436_v63).Dtor_cf8(), p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v1, !((_this).F6())), p0, globalState), _dafny.IntOfUint32((_433_v60).Cardinality()))).Uint32()).(_dafny.Sequence), (_index73).Int())
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_432_v59), 0))
		_ = _index74
		(_432_v59).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm11((_this).F6(), globalState), _367_v1), (_index74).Int())
		var _437_v64 _dafny.Map
		_ = _437_v64
		_437_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_v1, p2)
		var _438_v65 _dafny.Map
		_ = _438_v65
		_438_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F6())
		var _hi1 _dafny.Int = (_this).Fm4(p0, p0, _437_v64, (_438_v65).Cardinality(), globalState)
		_ = _hi1
		for _439_i2 := ((Companion_Default___.Fm12(globalState)).Intersection(_dafny.MultiSetFromSeq(_dafny.SeqOf(p0)))).Cardinality(); _439_i2.Cmp(_hi1) < 0; _439_i2 = _439_i2.Plus(_dafny.One) {
			var _440_v66 bool
			_ = _440_v66
			_440_v66 = false
			var _441_v67 _dafny.Array
			_ = _441_v67
			var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(22))
			_ = _nw68
			_441_v67 = _nw68
			var _442_v68 _dafny.Sequence
			_ = _442_v68
			_442_v68 = _dafny.SeqOf(_441_v67)
			var _443_v69 _dafny.Array
			_ = _443_v69
			var _nwElement0_9 _dafny.Array = _441_v67
			_ = _nwElement0_9
			var _nw69 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(12))
			_ = _nw69
			(_nw69).ArraySet1(_nwElement0_9, 0)
			(_nw69).ArraySet1(_441_v67, 1)
			(_nw69).ArraySet1(_441_v67, 2)
			(_nw69).ArraySet1(_441_v67, 3)
			(_nw69).ArraySet1(_441_v67, 4)
			(_nw69).ArraySet1((_442_v68).Select((Companion_Default___.SafeIndex(_439_i2, _dafny.IntOfUint32((_442_v68).Cardinality()))).Uint32()).(_dafny.Array), 5)
			(_nw69).ArraySet1(_441_v67, 6)
			(_nw69).ArraySet1(_441_v67, 7)
			(_nw69).ArraySet1(_441_v67, 8)
			(_nw69).ArraySet1(_441_v67, 9)
			(_nw69).ArraySet1(_441_v67, 10)
			(_nw69).ArraySet1(_441_v67, 11)
			_443_v69 = _nw69
			var _444_v70 _dafny.Sequence
			_ = _444_v70
			_444_v70 = _dafny.SeqOf((_this).F6())
			var _445_v71 _dafny.Sequence
			_ = _445_v71
			_445_v71 = _dafny.SeqOf((_439_i2).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32(((_432_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(599), _dafny.ArrayLen((_432_v59), 0))).Int()).(_dafny.Sequence)).Cardinality()))).Cardinality()), p0, (_this).Fm4(_dafny.IntOfInt64(459), _439_i2, Companion_Default___.Fm13(_dafny.IntOfUint32((_444_v70).Cardinality()), _436_v63, p2, (_dafny.SetOf(p1, (_this).F6(), (_this).F6(), p1)).Cardinality(), globalState), _dafny.IntOfInt64(49), globalState))
			var _446_v72 _dafny.Map
			_ = _446_v72
			_446_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_434_v61, _443_v69)
			var _447_v73 D4
			_ = _447_v73
			_447_v73 = Companion_D4_.Create_DC12_((func() _dafny.Array {
				if (_446_v72).Contains(_434_v61) {
					return (_446_v72).Get(_434_v61).(_dafny.Array)
				}
				return _443_v69
			})())
			var _448_v74 _dafny.Set
			_ = _448_v74
			_448_v74 = _dafny.SetOf(p0, (_dafny.Zero).Minus(p0))
			var _449_v75 _dafny.Sequence
			_ = _449_v75
			_449_v75 = _dafny.SeqOf(_445_v71, _445_v71)
			var _rhs77 bool = !(p1)
			_ = _rhs77
			var _rhs78 _dafny.Array = (_447_v73).Dtor_cf14()
			_ = _rhs78
			var _rhs79 _dafny.Int = _439_i2
			_ = _rhs79
			var _rhs80 _dafny.Int = (Companion_Default___.SafeModuloInt((_448_v74).Cardinality(), _439_i2)).Times(Companion_Default___.SafeDivisionInt(_439_i2, _439_i2))
			_ = _rhs80
			var _rhs81 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_449_v75).Select((Companion_Default___.SafeIndex(_439_i2, _dafny.IntOfUint32((_449_v75).Cardinality()))).Uint32()).(_dafny.Sequence), _445_v71)
			_ = _rhs81
			var _lhs55 *GlobalState = globalState
			_ = _lhs55
			var _lhs56 *GlobalState = globalState
			_ = _lhs56
			_440_v66 = _rhs77
			_443_v69 = _rhs78
			_lhs55.F0 = _rhs79
			_lhs56.F0 = _rhs80
			_445_v71 = _rhs81
			_434_v61 = _434_v61
			_438_v65 = (_438_v65).Update((_this).F6(), _440_v66)
			var _450_v76 _dafny.Map
			_ = _450_v76
			_450_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
			var _451_v77 T1
			_ = _451_v77
			var _nw70 *C1 = New_C1_()
			_ = _nw70
			_nw70.Ctor__((_this).F6(), _367_v1, true, _dafny.MultiSetOf(false, _440_v66))
			_451_v77 = _nw70
			var _452_v78 _dafny.Array
			_ = _452_v78
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_11
			var _nw71 _dafny.Array
			_ = _nw71
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw71 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) bool = (func(_453_v77 T1) func(_dafny.Int) bool {
					return func(_454_i3 _dafny.Int) bool {
						return !((_453_v77).F6())
					}
				})(_451_v77)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw71 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw71).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw71).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_452_v78 = _nw71
			var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_452_v78), 0))
			_ = _index75
			(_452_v78).ArraySet1(!((_this).F6()), (_index75).Int())
			var _455_v80 _dafny.MultiSet
			_ = _455_v80
			_455_v80 = _dafny.MultiSetOf(_368_v2)
			var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_452_v78), 0))
			_ = _index76
			var _rhs82 _dafny.Map = (func() _dafny.Map {
				if !((p0).Cmp(_439_i2) <= 0) {
					return _450_v76
				}
				return (func() _dafny.Map {
					if (_451_v77).F6() {
						return func() _dafny.Map {
							var _coll28 = _dafny.NewMapBuilder()
							_ = _coll28
							for _iter31 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(938), _dafny.IntOfInt64(-71))); ; {
								_compr_28, _ok31 := _iter31()
								if !_ok31 {
									break
								}
								var _456_v79 _dafny.Int
								_456_v79 = interface{}(_compr_28).(_dafny.Int)
								if ((_dafny.IntOfInt64(938)).Cmp(_456_v79) <= 0) && ((_456_v79).Cmp(_dafny.IntOfInt64(-71)) < 0) {
									_coll28.Add((_456_v79).Times(_439_i2), (_451_v77).F6())
								}
							}
							return _coll28.ToMap()
						}()
					}
					return _450_v76
				})()
			})()
			_ = _rhs82
			var _rhs83 T1 = _451_v77
			_ = _rhs83
			var _rhs84 _dafny.CodePoint = _434_v61
			_ = _rhs84
			var _rhs85 bool = ((_455_v80).Update(_368_v2, Companion_Default___.Abs(_439_i2))).IsDisjointFrom(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(522))).Uint32(), func(coer29 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
				return func(arg29 _dafny.Int) interface{} {
					return coer29(arg29)
				}
			}((func(_457_v2 D0) func(_dafny.Int) D0 {
				return func(_458_i4 _dafny.Int) D0 {
					return _457_v2
				}
			})(_368_v2)))))
			_ = _rhs85
			var _lhs57 _dafny.Array = _452_v78
			_ = _lhs57
			var _lhs58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(71), _dafny.ArrayLen((_452_v78), 0))
			_ = _lhs58
			_450_v76 = _rhs82
			_451_v77 = _rhs83
			_434_v61 = _rhs84
			(_lhs57).ArraySet1(_rhs85, (_lhs58).Int())
		}
	}
}
func (_this *C3) M6(p0 _dafny.Int, p1 _dafny.Map, p2 bool, p3 D1, globalState *GlobalState) (D1, _dafny.Sequence, _dafny.Sequence, bool) {
	{
		var r0 D1 = Companion_D1_.Default()
		_ = r0
		var r1 _dafny.Sequence = _dafny.EmptySeq
		_ = r1
		var r2 _dafny.Sequence = _dafny.EmptySeq
		_ = r2
		var r3 bool = false
		_ = r3
		var _459_v0 _dafny.Array
		_ = _459_v0
		var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.One)
		_ = _nw72
		_459_v0 = _nw72
		var _nw73 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
		_ = _nw73
		_459_v0 = _nw73
		r3 = p2
		if p2 {
			if (_this).F6() {
				var _460_v1 _dafny.Array
				_ = _460_v1
				var _nw74 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(14))
				_ = _nw74
				_460_v1 = _nw74
				var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_460_v1), 0))
				_ = _index77
				(_460_v1).ArraySet1((p0).Cmp(p0) == 0, (_index77).Int())
				var _461_v2 _dafny.Sequence
				_ = _461_v2
				_461_v2 = _dafny.SeqOf((_this).Fm1((_this).Fm3(globalState), p0, p0, globalState))
				var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_460_v1), 0))
				_ = _index78
				(_460_v1).ArraySet1((p2) && ((p0).Cmp(_dafny.IntOfUint32((_461_v2).Cardinality())) != 0), (_index78).Int())
				(globalState).F0 = (_dafny.Zero).Minus(p0)
				var _462_v3 _dafny.Array
				_ = _462_v3
				var _nw75 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
				_ = _nw75
				_462_v3 = _nw75
				var _463_v4 *C1
				_ = _463_v4
				var _nw76 *C1 = New_C1_()
				_ = _nw76
				_nw76.Ctor__(false, _dafny.UnicodeSeqOfUtf8Bytes("xy"), p2, _dafny.MultiSetFromSeq(_461_v2))
				_463_v4 = _nw76
				var _464_v5 _dafny.Array
				_ = _464_v5
				var _nwElement0_10 *C1 = _463_v4
				_ = _nwElement0_10
				var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(29))
				_ = _nw77
				(_nw77).ArraySet1(_nwElement0_10, 0)
				(_nw77).ArraySet1(_463_v4, 1)
				(_nw77).ArraySet1(_463_v4, 2)
				(_nw77).ArraySet1(_463_v4, 3)
				(_nw77).ArraySet1(_463_v4, 4)
				(_nw77).ArraySet1(_463_v4, 5)
				(_nw77).ArraySet1(_463_v4, 6)
				(_nw77).ArraySet1(_463_v4, 7)
				(_nw77).ArraySet1(_463_v4, 8)
				(_nw77).ArraySet1(_463_v4, 9)
				(_nw77).ArraySet1(_463_v4, 10)
				(_nw77).ArraySet1(_463_v4, 11)
				(_nw77).ArraySet1(_463_v4, 12)
				(_nw77).ArraySet1(_463_v4, 13)
				(_nw77).ArraySet1(_463_v4, 14)
				(_nw77).ArraySet1(_463_v4, 15)
				(_nw77).ArraySet1(_463_v4, 16)
				(_nw77).ArraySet1(_463_v4, 17)
				(_nw77).ArraySet1(_463_v4, 18)
				(_nw77).ArraySet1(_463_v4, 19)
				(_nw77).ArraySet1(_463_v4, 20)
				(_nw77).ArraySet1(_463_v4, 21)
				(_nw77).ArraySet1(_463_v4, 22)
				(_nw77).ArraySet1(_463_v4, 23)
				(_nw77).ArraySet1(_463_v4, 24)
				(_nw77).ArraySet1(_463_v4, 25)
				(_nw77).ArraySet1(_463_v4, 26)
				(_nw77).ArraySet1(_463_v4, 27)
				(_nw77).ArraySet1(_463_v4, 28)
				_464_v5 = _nw77
				var _465_v6 _dafny.Map
				_ = _465_v6
				_465_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_460_v1, _464_v5)
				var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_462_v3), 0))
				_ = _index79
				(_462_v3).ArraySet1(_465_v6, (_index79).Int())
				var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_462_v3), 0))
				_ = _index80
				(_462_v3).ArraySet1(_465_v6, (_index80).Int())
				var _466_v7 _dafny.CodePoint
				_ = _466_v7
				_466_v7 = _dafny.CodePoint('r')
				var _467_v8 bool
				_ = _467_v8
				var _468_v9 _dafny.Int
				_ = _468_v9
				var _469_v10 bool
				_ = _469_v10
				var _470_v11 bool
				_ = _470_v11
				var _out3 bool
				_ = _out3
				var _out4 _dafny.Int
				_ = _out4
				var _out5 bool
				_ = _out5
				var _out6 bool
				_ = _out6
				_out3, _out4, _out5, _out6 = (_463_v4).M8(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update((_463_v4).F13(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_463_v4).F13()).Cardinality()))).Uint32(), _466_v7), (_463_v4).F13()), globalState)
				_467_v8 = _out3
				_468_v9 = _out4
				_469_v10 = _out5
				_470_v11 = _out6
				var _471_v12 _dafny.Map
				_ = _471_v12
				_471_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_468_v9, (func() bool {
					if (_463_v4).F12() {
						return false
					}
					return (_460_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(765), _dafny.ArrayLen((_460_v1), 0))).Int()).(bool)
				})())
				_471_v12 = _471_v12
			} else {
				(globalState).F0 = p0
				var _472_v13 _dafny.Map
				_ = _472_v13
				_472_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
				var _473_v14 _dafny.Sequence
				_ = _473_v14
				_473_v14 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F6()), _472_v13, _472_v13, _472_v13)
				var _474_v15 _dafny.Map
				_ = _474_v15
				_474_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
				var _475_v16 _dafny.Sequence
				_ = _475_v16
				_475_v16 = _dafny.SeqOf(p2, (_this).F6())
				var _476_v17 _dafny.Map
				_ = _476_v17
				_476_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_472_v13).Equals((_473_v14).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_473_v14).Cardinality()))).Uint32()).(_dafny.Map)), (func() _dafny.Int {
					if (_474_v15).Contains((_this).F6()) {
						return (_474_v15).Get((_this).F6()).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_475_v16).Cardinality())
				})())
				_476_v17 = _474_v15
				r3 = p2
				(globalState).F0 = p0
				var _477_v18 _dafny.Array
				_ = _477_v18
				var _nw78 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(15))
				_ = _nw78
				_477_v18 = _nw78
				var _478_v19 T0
				_ = _478_v19
				var _nw79 *C2 = New_C2_()
				_ = _nw79
				_nw79.Ctor__((_this).F5())
				_478_v19 = _nw79
				var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_477_v18), 0))
				_ = _index81
				(_477_v18).ArraySet1(_478_v19, (_index81).Int())
				var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(949), _dafny.ArrayLen((_477_v18), 0))
				_ = _index82
				(_477_v18).ArraySet1((func() T0 {
					if false {
						return _478_v19
					}
					return _478_v19
				})(), (_index82).Int())
			}
			if !(p2) {
				var _479_v20 _dafny.Sequence
				_ = _479_v20
				_479_v20 = _dafny.SeqOf((_this).F6())
				var _480_v21 *C2
				_ = _480_v21
				var _nw80 *C2 = New_C2_()
				_ = _nw80
				_nw80.Ctor__(_dafny.MultiSetFromSeq(_479_v20))
				_480_v21 = _nw80
				var _481_v22 _dafny.Map
				_ = _481_v22
				var _out7 _dafny.Map
				_ = _out7
				_out7 = (_480_v21).M9((p2) && ((_this).F6()), p0, globalState)
				_481_v22 = _out7
				var _482_v23 _dafny.Sequence
				_ = _482_v23
				_482_v23 = _dafny.UnicodeSeqOfUtf8Bytes("edhl")
				var _483_v24 _dafny.CodePoint
				_ = _483_v24
				_483_v24 = _dafny.CodePoint('q')
				var _484_v25 _dafny.Map
				_ = _484_v25
				_484_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_482_v23, (_this).F6())
				(globalState).F0 = (_this).Fm4(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_482_v23).Cardinality()), _dafny.IntOfInt64(180)), (p0).Minus(p0), (func() _dafny.Map {
					if (_this).Fm3(globalState) {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_482_v23, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_482_v23).Cardinality()))).Uint32(), _483_v24), false)
					}
					return _484_v25
				})(), (func() _dafny.Int {
					if ((_this).F5()).Contains(p2) {
						return ((_this).F5()).Multiplicity(p2)
					}
					return p0
				})(), globalState)
				var _485_v26 D1
				_ = _485_v26
				_485_v26 = Companion_D1_.Create_DC2_((_this).F6())
				_485_v26 = p3
				var _486_v27 *C2
				_ = _486_v27
				var _nw81 *C2 = New_C2_()
				_ = _nw81
				_nw81.Ctor__(_dafny.MultiSetOf(p2, (_this).Fm3(globalState), (_this).Fm2((_this).F6(), globalState)))
				_486_v27 = _nw81
			} else {
				var _487_v28 _dafny.Sequence
				_ = _487_v28
				_487_v28 = _dafny.SeqOf((_this).F6())
				var _488_v29 _dafny.Map
				_ = _488_v29
				_488_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_487_v28, p2)
				var _489_v30 _dafny.Sequence
				_ = _489_v30
				_489_v30 = _dafny.UnicodeSeqOfUtf8Bytes("mflyhpk")
				var _490_v31 _dafny.Map
				_ = _490_v31
				_490_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).Fm1(p2, p0, _dafny.IntOfInt64(46), globalState)), (_this).F6())
				var _491_v32 *C1
				_ = _491_v32
				var _nw82 *C1 = New_C1_()
				_ = _nw82
				_nw82.Ctor__((_this).F6(), _489_v30, !(_490_v31).Contains((_this).F6()), (_this).F5())
				_491_v32 = _nw82
				var _rhs86 _dafny.Map = (Companion_Default___.Fm20(p0, globalState)).Update(_487_v28, p2)
				_ = _rhs86
				var _rhs87 *C1 = _491_v32
				_ = _rhs87
				_488_v29 = _rhs86
				_491_v32 = _rhs87
				var _492_v33 _dafny.Array
				_ = _492_v33
				var _len0_12 _dafny.Int = _dafny.IntOfInt64(22)
				_ = _len0_12
				var _nw83 _dafny.Array
				_ = _nw83
				if _len0_12.Cmp(_dafny.Zero) == 0 {
					_nw83 = _dafny.NewArray(_len0_12)
				} else {
					var _init12 func(_dafny.Int) bool = func(_493_i0 _dafny.Int) bool {
						return (_this).F6()
					}
					_ = _init12
					var _element0_12 = _init12(_dafny.Zero)
					_ = _element0_12
					_nw83 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
					(_nw83).ArraySet1(_element0_12, 0)
					var _nativeLen0_12 = (_len0_12).Int()
					_ = _nativeLen0_12
					for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
						(_nw83).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
					}
				}
				_492_v33 = _nw83
				var _494_v34 _dafny.Set
				_ = _494_v34
				_494_v34 = _dafny.SetOf(_492_v33)
				_494_v34 = _494_v34
				var _495_v35 _dafny.Map
				_ = _495_v35
				_495_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_491_v32).F13(), (p0).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_491_v32).F13()).Cardinality()))))
				var _496_v37 _dafny.Map
				_ = _496_v37
				_496_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_495_v35 = (_495_v35).Update(_dafny.Companion_Sequence_.Concatenate((_491_v32).F13(), Companion_Default___.Fm11((_491_v32).F12(), globalState)), (_dafny.Zero).Minus((func() _dafny.Int {
					if p2 {
						return (_dafny.SetOf(func() _dafny.Map {
							var _coll29 = _dafny.NewMapBuilder()
							_ = _coll29
							for _iter32 := _dafny.Iterate((_496_v37).Keys().Elements()); ; {
								_compr_29, _ok32 := _iter32()
								if !_ok32 {
									break
								}
								var _497_v36 _dafny.Int
								_497_v36 = interface{}(_compr_29).(_dafny.Int)
								if (_496_v37).Contains(_497_v36) {
									_coll29.Add(Companion_Default___.SafeModuloInt(_497_v36, p0), (_491_v32).F12())
								}
							}
							return _coll29.ToMap()
						}())).Cardinality()
					}
					return (_dafny.Zero).Minus(p0)
				})()))
				var _498_v38 _dafny.Array
				_ = _498_v38
				var _nw84 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
				_ = _nw84
				_498_v38 = _nw84
				var _499_v39 D3
				_ = _499_v39
				_499_v39 = Companion_D3_.Create_DC10_(!(false), _dafny.IntOfInt64(-542))
				var _500_v40 D0
				_ = _500_v40
				_500_v40 = Companion_D0_.Create_DC0_(p0)
				var _pat_let_tv14 = p0
				_ = _pat_let_tv14
				var _pat_let_tv15 = p0
				_ = _pat_let_tv15
				var _501_v41 _dafny.Array
				_ = _501_v41
				var _nwElement0_11 D3 = _499_v39
				_ = _nwElement0_11
				var _nw85 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(21))
				_ = _nw85
				(_nw85).ArraySet1(_nwElement0_11, 0)
				(_nw85).ArraySet1(_499_v39, 1)
				(_nw85).ArraySet1(_499_v39, 2)
				(_nw85).ArraySet1(Companion_D3_.Create_DC10_(p2, _dafny.IntOfInt64(816)), 3)
				(_nw85).ArraySet1(_499_v39, 4)
				(_nw85).ArraySet1(_499_v39, 5)
				(_nw85).ArraySet1(_499_v39, 6)
				(_nw85).ArraySet1(func(_pat_let20_0 D3) D3 {
					return func(_502_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let21_0 _dafny.Int) D3 {
							return func(_503_dt__update_hcf12_h0 _dafny.Int) D3 {
								return Companion_D3_.Create_DC10_((_502_dt__update__tmp_h0).Dtor_cf11(), _503_dt__update_hcf12_h0)
							}(_pat_let21_0)
						}(_pat_let_tv14)
					}(_pat_let20_0)
				}(Companion_Default___.Fm21((_491_v32).F12(), globalState)), 7)
				(_nw85).ArraySet1(_499_v39, 8)
				(_nw85).ArraySet1(Companion_D3_.Create_DC10_((_491_v32).F12(), p0), 9)
				(_nw85).ArraySet1(_499_v39, 10)
				(_nw85).ArraySet1(_499_v39, 11)
				(_nw85).ArraySet1(_499_v39, 12)
				(_nw85).ArraySet1(_499_v39, 13)
				(_nw85).ArraySet1(_499_v39, 14)
				(_nw85).ArraySet1(_499_v39, 15)
				(_nw85).ArraySet1(Companion_D3_.Create_DC10_(p2, p0), 16)
				(_nw85).ArraySet1(_499_v39, 17)
				(_nw85).ArraySet1(_499_v39, 18)
				(_nw85).ArraySet1(func(_pat_let22_0 D3) D3 {
					return func(_504_dt__update__tmp_h1 D3) D3 {
						return func(_pat_let23_0 _dafny.Int) D3 {
							return func(_505_dt__update_hcf12_h1 _dafny.Int) D3 {
								return Companion_D3_.Create_DC10_((_504_dt__update__tmp_h1).Dtor_cf11(), _505_dt__update_hcf12_h1)
							}(_pat_let23_0)
						}(_pat_let_tv15)
					}(_pat_let22_0)
				}(_499_v39), 19)
				(_nw85).ArraySet1(Companion_D3_.Create_DC10_(p2, (_500_v40).Dtor_cf0()), 20)
				_501_v41 = _nw85
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_498_v38), 0))
				_ = _index83
				(_498_v38).ArraySet1(_501_v41, (_index83).Int())
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(606), _dafny.ArrayLen((_498_v38), 0))
				_ = _index84
				(_498_v38).ArraySet1((func() _dafny.Array {
					if !(false) {
						return _501_v41
					}
					return _501_v41
				})(), (_index84).Int())
				_492_v33 = _492_v33
			}
			var _506_v42 _dafny.Sequence
			_ = _506_v42
			_506_v42 = _dafny.SeqOf((_this).F6())
			var _507_v43 _dafny.Set
			_ = _507_v43
			_507_v43 = _dafny.SetOf((_dafny.SetOf(p2, (_this).F6())).Cardinality(), (_dafny.MultiSetFromSeq(_506_v42)).Cardinality(), p0, p0, p0)
			var _508_v44 _dafny.Map
			_ = _508_v44
			_508_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
			var _509_v45 _dafny.Set
			_ = _509_v45
			_509_v45 = _dafny.SetOf(_459_v0)
			var _510_v46 _dafny.CodePoint
			_ = _510_v46
			_510_v46 = _dafny.CodePoint('n')
			var _511_v47 _dafny.Sequence
			_ = _511_v47
			_511_v47 = _dafny.SeqOf(_510_v46, _510_v46)
			var _512_v48 _dafny.Map
			_ = _512_v48
			_512_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_511_v47).Cardinality()), (_this).F6())
			var _513_v49 _dafny.Sequence
			_ = _513_v49
			_513_v49 = _dafny.SeqOf((_dafny.Zero).Minus(p0), p0)
			var _514_v50 _dafny.Array
			_ = _514_v50
			var _nwElement0_12 bool = !(true)
			_ = _nwElement0_12
			var _nw86 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(22))
			_ = _nw86
			(_nw86).ArraySet1(_nwElement0_12, 0)
			(_nw86).ArraySet1(p2, 1)
			(_nw86).ArraySet1(!((_this).F6()) || ((_this).F6()), 2)
			(_nw86).ArraySet1(((_this).F6()) == (p2), 3)
			(_nw86).ArraySet1((_this).F6(), 4)
			(_nw86).ArraySet1((_this).F6(), 5)
			(_nw86).ArraySet1(p2, 6)
			(_nw86).ArraySet1(p2, 7)
			(_nw86).ArraySet1((_507_v43).IsProperSubsetOf(_dafny.SetOf(_dafny.IntOfInt64(-874))), 8)
			(_nw86).ArraySet1((_this).F6(), 9)
			(_nw86).ArraySet1(p2, 10)
			(_nw86).ArraySet1(p2, 11)
			(_nw86).ArraySet1(p2, 12)
			(_nw86).ArraySet1(p2, 13)
			(_nw86).ArraySet1((Companion_Default___.Fm21(!((_this).F6()), globalState)).Dtor_cf11(), 14)
			(_nw86).ArraySet1(true, 15)
			(_nw86).ArraySet1((func() bool {
				if (_508_v44).Contains(!(false)) {
					return (_508_v44).Get(!(false)).(bool)
				}
				return (_this).F6()
			})(), 16)
			(_nw86).ArraySet1((_dafny.SetOf(_459_v0, _459_v0)).IsSubsetOf(_509_v45), 17)
			(_nw86).ArraySet1((_this).Fm3(globalState), 18)
			(_nw86).ArraySet1(((_512_v48).Cardinality()).Cmp((_513_v49).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_513_v49).Cardinality()))).Uint32()).(_dafny.Int)) <= 0, 19)
			(_nw86).ArraySet1(p2, 20)
			(_nw86).ArraySet1(p2, 21)
			_514_v50 = _nw86
			var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))
			_ = _index85
			(_514_v50).ArraySet1(p2, (_index85).Int())
			var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))
			_ = _index86
			(_514_v50).ArraySet1(false, (_index86).Int())
			var _515_v51 _dafny.Array
			_ = _515_v51
			var _len0_13 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_13
			var _nw87 _dafny.Array
			_ = _nw87
			if _len0_13.Cmp(_dafny.Zero) == 0 {
				_nw87 = _dafny.NewArray(_len0_13)
			} else {
				var _init13 func(_dafny.Int) _dafny.Map = (func(_516_v47 _dafny.Sequence, _517_v50 _dafny.Array, _518_p0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_519_i1 _dafny.Int) _dafny.Map {
						return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfUint32((_516_v47).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_517_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_517_v50), 0))).Int()).(bool), _518_p0))
					}
				})(_511_v47, _514_v50, p0)
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
			_515_v51 = _nw87
			var _520_v52 _dafny.Map
			_ = _520_v52
			_520_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_514_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))).Int()).(bool)), p0)
			var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_515_v51), 0))
			_ = _index87
			(_515_v51).ArraySet1(_520_v52, (_index87).Int())
			var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_515_v51), 0))
			_ = _index88
			var _rhs88 _dafny.Set = (func() _dafny.Set {
				if (_this).F6() {
					return _507_v43
				}
				return (_507_v43).Difference(func() _dafny.Set {
					var _coll30 = _dafny.NewBuilder()
					_ = _coll30
					for _iter33 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-695), _dafny.IntOfInt64(274))); ; {
						_compr_30, _ok33 := _iter33()
						if !_ok33 {
							break
						}
						var _521_v53 _dafny.Int
						_521_v53 = interface{}(_compr_30).(_dafny.Int)
						if ((_dafny.IntOfInt64(-695)).Cmp(_521_v53) <= 0) && ((_521_v53).Cmp(_dafny.IntOfInt64(274)) < 0) {
							_coll30.Add((_521_v53).Minus(p0))
						}
					}
					return _coll30.ToSet()
				}())
			})()
			_ = _rhs88
			var _rhs89 bool = (_this).Fm3(globalState)
			_ = _rhs89
			var _rhs90 bool = (func() bool {
				if (p0).Cmp(p0) > 0 {
					return (_514_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))).Int()).(bool)
				}
				return p2
			})()
			_ = _rhs90
			var _rhs91 _dafny.Map = ((_520_v52).Merge(Companion_Default___.Fm19(p0, globalState))).Update(((_514_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))).Int()).(bool)) && ((_this).F6()), (_dafny.IntOfInt64(438)).Times((((_this).F5()).Update((_514_v50).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))).Int()).(bool), Companion_Default___.Abs(p0))).Cardinality()))
			_ = _rhs91
			var _lhs59 _dafny.Array = _515_v51
			_ = _lhs59
			var _lhs60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_515_v51), 0))
			_ = _lhs60
			_507_v43 = _rhs88
			r3 = _rhs89
			r3 = _rhs90
			(_lhs59).ArraySet1(_rhs91, (_lhs60).Int())
			var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_514_v50), 0))
			_ = _index89
			(_514_v50).ArraySet1((_this).Fm1(false, p0, Companion_Default___.SafeDivisionInt(p0, p0), globalState), (_index89).Int())
		} else {
			var _522_v54 _dafny.Array
			_ = _522_v54
			var _nwElement0_13 bool = !(!(p2))
			_ = _nwElement0_13
			var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(8))
			_ = _nw88
			(_nw88).ArraySet1(_nwElement0_13, 0)
			(_nw88).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(824))).Uint32(), func(coer30 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg30 _dafny.Int) interface{} {
					return coer30(arg30)
				}
			}(func(_523_i2 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('f')
			})), _dafny.UnicodeSeqOfUtf8Bytes("yaxq")), 1)
			(_nw88).ArraySet1((_this).F6(), 2)
			(_nw88).ArraySet1((p0).Cmp(p0) != 0, 3)
			(_nw88).ArraySet1(p2, 4)
			(_nw88).ArraySet1((p0).Cmp(p0) != 0, 5)
			(_nw88).ArraySet1(p2, 6)
			(_nw88).ArraySet1((_this).F6(), 7)
			_522_v54 = _nw88
			var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_522_v54), 0))
			_ = _index90
			(_522_v54).ArraySet1((_this).F6(), (_index90).Int())
			var _524_v55 _dafny.Sequence
			_ = _524_v55
			_524_v55 = _dafny.SeqOf(p0, p0, _dafny.IntOfInt64(-775), p0)
			var _525_v56 _dafny.Sequence
			_ = _525_v56
			_525_v56 = _dafny.UnicodeSeqOfUtf8Bytes("jrojrlvjv")
			var _526_v57 _dafny.Map
			_ = _526_v57
			_526_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_525_v56, (_this).F6())
			var _527_v58 _dafny.Set
			_ = _527_v58
			_527_v58 = _dafny.SetOf(p2, (_this).F6())
			var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_522_v54), 0))
			_ = _index91
			var _rhs92 _dafny.Int = (_this).Fm4(Companion_Default___.SafeModuloInt((_524_v55).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_524_v55).Cardinality()))).Uint32()).(_dafny.Int), (_dafny.Zero).Minus(p0)), _dafny.IntOfInt64(512), (_526_v57).Merge(_526_v57), _dafny.IntOfInt64(108), globalState)
			_ = _rhs92
			var _rhs93 bool = p2
			_ = _rhs93
			var _rhs94 bool = !((_527_v58).IsDisjointFrom(_527_v58))
			_ = _rhs94
			var _lhs61 *GlobalState = globalState
			_ = _lhs61
			var _lhs62 _dafny.Array = _522_v54
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_522_v54), 0))
			_ = _lhs63
			_lhs61.F0 = _rhs92
			(_lhs62).ArraySet1(_rhs93, (_lhs63).Int())
			r3 = _rhs94
			(globalState).F0 = p0
			var _rhs95 _dafny.Int = p0
			_ = _rhs95
			var _rhs96 _dafny.Int = (((_527_v58).Cardinality()).Plus(p0)).Times(p0)
			_ = _rhs96
			var _lhs64 *GlobalState = globalState
			_ = _lhs64
			var _lhs65 *GlobalState = globalState
			_ = _lhs65
			_lhs64.F0 = _rhs95
			_lhs65.F0 = _rhs96
			var _528_v59 _dafny.CodePoint
			_ = _528_v59
			_528_v59 = _dafny.CodePoint('w')
			_528_v59 = _528_v59
			var _529_v60 _dafny.Map
			_ = _529_v60
			_529_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(p0, p0, p0, _dafny.IntOfInt64(966)), _dafny.SeqOf((_522_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_522_v54), 0))).Int()).(bool), true))
			var _530_v61 _dafny.Map
			_ = _530_v61
			_530_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_529_v60, ((_522_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(725), _dafny.ArrayLen((_522_v54), 0))).Int()).(bool)) && (p2))
			_530_v61 = (_530_v61).Update(_529_v60, !(!(!((_this).F6())) || (false)))
		}
		if (_this).F6() {
			var _531_v62 _dafny.MultiSet
			_ = _531_v62
			_531_v62 = _dafny.MultiSetOf((_this).F6())
			_531_v62 = ((_this).F5()).Difference((_this).F5())
			var _532_v63 _dafny.CodePoint
			_ = _532_v63
			_532_v63 = _dafny.CodePoint('t')
			var _533_v64 D2
			_ = _533_v64
			_533_v64 = Companion_D2_.Create_DC7_(_532_v63, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0), p0)
			var _534_v65 D2
			_ = _534_v65
			_534_v65 = Companion_D2_.Create_DC8_(_533_v64)
			var _535_v66 _dafny.Map
			_ = _535_v66
			_535_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_534_v65, p0)
			var _536_v67 _dafny.Map
			_ = _536_v67
			_536_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_535_v66).Update(_534_v65, p0), p0)
			(globalState).F0 = (func() _dafny.Int {
				if (_536_v67).Contains(_535_v66) {
					return (_536_v67).Get(_535_v66).(_dafny.Int)
				}
				return p0
			})()
			r2 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ysdg"), Companion_Default___.Fm11(p2, globalState))
			r3 = (_this).F6()
			r3 = (_this).F6()
		} else {
			var _537_v68 D3
			_ = _537_v68
			_537_v68 = Companion_D3_.Create_DC10_(p2, p0)
			var _538_v69 _dafny.Sequence
			_ = _538_v69
			_538_v69 = _dafny.SeqOf((_537_v68).Dtor_cf12())
			_538_v69 = _538_v69
			var _539_v70 *C2
			_ = _539_v70
			var _nw89 *C2 = New_C2_()
			_ = _nw89
			_nw89.Ctor__((_this).F5())
			_539_v70 = _nw89
			r3 = (_this).F6()
			var _540_v71 _dafny.Map
			_ = _540_v71
			_540_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(622), ((_this).F5()).Cardinality())
			r3 = ((_dafny.MultiSetOf((_539_v70).Fm1((_this).F6(), (_540_v71).Cardinality(), p0, globalState), p2)).IsSubsetOf((_this).F5())) == (!((_this).F6()))
			var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_459_v0), 0))
			_ = _index92
			(_459_v0).ArraySet1(p0, (_index92).Int())
			var _541_v72 _dafny.Map
			_ = _541_v72
			_541_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), p2)
			var _542_v73 _dafny.Sequence
			_ = _542_v73
			_542_v73 = _dafny.UnicodeSeqOfUtf8Bytes("qgee")
			var _543_v74 _dafny.Map
			_ = _543_v74
			_543_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(157), _542_v73)
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_459_v0), 0))
			_ = _index93
			var _rhs97 _dafny.Int = (Companion_Default___.SafeModuloInt(p0, (_541_v72).Cardinality())).Plus(p0)
			_ = _rhs97
			var _rhs98 _dafny.Int = _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_543_v74).Contains((func() _dafny.Int {
					if (_this).F6() {
						return p0
					}
					return p0
				})()) {
					return (_543_v74).Get((func() _dafny.Int {
						if (_this).F6() {
							return p0
						}
						return p0
					})()).(_dafny.Sequence)
				}
				return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(666))).Uint32(), func(coer31 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg31 _dafny.Int) interface{} {
						return coer31(arg31)
					}
				}(func(_544_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), _542_v73)
			})()).Cardinality())
			_ = _rhs98
			var _rhs99 bool = (_this).Fm2(p2, globalState)
			_ = _rhs99
			var _lhs66 *GlobalState = globalState
			_ = _lhs66
			var _lhs67 _dafny.Array = _459_v0
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_459_v0), 0))
			_ = _lhs68
			_lhs66.F0 = _rhs97
			(_lhs67).ArraySet1(_rhs98, (_lhs68).Int())
			r3 = _rhs99
		}
		r3 = (_this).F6()
		var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_459_v0), 0))
		_ = _index94
		(_459_v0).ArraySet1(p0, (_index94).Int())
		var _545_v75 _dafny.Sequence
		_ = _545_v75
		_545_v75 = _dafny.UnicodeSeqOfUtf8Bytes("frah")
		var _546_v76 _dafny.Map
		_ = _546_v76
		_546_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_545_v75, !(p2))
		var _547_v77 _dafny.CodePoint
		_ = _547_v77
		_547_v77 = _dafny.CodePoint('t')
		var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_459_v0), 0))
		_ = _index95
		var _rhs100 _dafny.Int = (_this).Fm4(p0, (_dafny.Zero).Minus(p0), _546_v76, p0, globalState)
		_ = _rhs100
		var _rhs101 _dafny.Int = (func() _dafny.Int {
			if ((_this).F5()).Contains((_this).F6()) {
				return ((_this).F5()).Multiplicity((_this).F6())
			}
			return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_545_v75, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_545_v75).Cardinality()), _dafny.IntOfUint32((_545_v75).Cardinality()))).Uint32(), _547_v77)).Cardinality())
		})()
		_ = _rhs101
		var _lhs69 *GlobalState = globalState
		_ = _lhs69
		var _lhs70 _dafny.Array = _459_v0
		_ = _lhs70
		var _lhs71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(85), _dafny.ArrayLen((_459_v0), 0))
		_ = _lhs71
		_lhs69.F0 = _rhs100
		(_lhs70).ArraySet1(_rhs101, (_lhs71).Int())
		r0 = Companion_Default___.Fm22(globalState)
		r1 = _dafny.Companion_Sequence_.Concatenate(_545_v75, _545_v75)
		r2 = _dafny.Companion_Sequence_.Concatenate(_545_v75, _545_v75)
		r3 = (_this).F6()
		return r0, r1, r2, r3
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f7 _dafny.Array
	_f8 bool
	_f6 bool
	_f5 _dafny.MultiSet
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = false
	_this._f6 = false
	_this._f5 = _dafny.EmptyMultiSet
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C4{}
var _ T1 = &C4{}
var _ T0 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F7() _dafny.Array {
	return _this._f7
}
func (_this *C4) F7_set_(value _dafny.Array) {
	_this._f7 = value
}
func (_this *C4) F8() bool {
	return _this._f8
}
func (_this *C4) F8_set_(value bool) {
	_this._f8 = value
}
func (_this *C4) F6() bool {
	return _this._f6
}
func (_this *C4) F5() _dafny.MultiSet {
	return _this._f5
}
func (_this *C4) Ctor__(f7 _dafny.Array, f8 bool, f6 bool, f5 _dafny.MultiSet) {
	{
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f6 = f6
		(_this)._f5 = f5
	}
}
func (_this *C4) Fm3(globalState *GlobalState) bool {
	{
		return _dafny.Companion_Sequence_.IsPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf(_this.F8(), (_this).F6(), _this.F8(), _this.F8(), _this.F8())), _dafny.SeqOf(_this.F8(), (_this).F6(), (_this).F6(), (_this).F6(), _this.F8()))
	}
}
func (_this *C4) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-256)
	}
}
func (_this *C4) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return _this.F8()
	}
}
func (_this *C4) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return (_this).F6()
	}
}
func (_this *C4) Fm6(globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.IntOfInt64(32)).Minus((_dafny.Zero).Minus((func() _dafny.Int {
			if ((_this).F5()).Contains(_this.F8()) {
				return ((_this).F5()).Multiplicity(_this.F8())
			}
			return _dafny.IntOfInt64(-44)
		})()))
	}
}
func (_this *C4) Fm7(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("lsi"), _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jxsfieg"), _dafny.UnicodeSeqOfUtf8Bytes("isrelfj")))
	}
}
func (_this *C4) M2(p0 _dafny.CodePoint, globalState *GlobalState) {
	{
		var _548_v0 _dafny.CodePoint
		_ = _548_v0
		_548_v0 = _dafny.CodePoint('o')
		var _rhs102 bool = false
		_ = _rhs102
		var _rhs103 _dafny.Int = (_this).Fm6(globalState)
		_ = _rhs103
		var _rhs104 _dafny.CodePoint = _548_v0
		_ = _rhs104
		var _lhs72 *C4 = _this
		_ = _lhs72
		var _lhs73 *GlobalState = globalState
		_ = _lhs73
		_lhs72.F8_set_(_rhs102)
		_lhs73.F0 = _rhs103
		_548_v0 = _rhs104
		var _549_v1 _dafny.Map
		_ = _549_v1
		_549_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F8())
		_549_v1 = (_549_v1).Update(((_this).F6()) && ((_this).F6()), (_this).F6())
		var _550_v2 _dafny.Int
		_ = _550_v2
		_550_v2 = _dafny.IntOfInt64(725)
		var _551_v3 D0
		_ = _551_v3
		_551_v3 = Companion_D0_.Create_DC1_()
		var _552_v4 _dafny.Sequence
		_ = _552_v4
		_552_v4 = _dafny.UnicodeSeqOfUtf8Bytes("elhkbkopj")
		var _553_v5 _dafny.Sequence
		_ = _553_v5
		_553_v5 = _dafny.SeqOf(_552_v4, _552_v4)
		var _554_v6 _dafny.Array
		_ = _554_v6
		var _nwElement0_14 _dafny.Int = _550_v2
		_ = _nwElement0_14
		var _nw90 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(18))
		_ = _nw90
		(_nw90).ArraySet1(_nwElement0_14, 0)
		(_nw90).ArraySet1(_dafny.IntOfUint32((_552_v4).Cardinality()), 1)
		(_nw90).ArraySet1(_550_v2, 2)
		(_nw90).ArraySet1(_550_v2, 3)
		(_nw90).ArraySet1(_dafny.IntOfInt64(365), 4)
		(_nw90).ArraySet1(_550_v2, 5)
		(_nw90).ArraySet1((_this).Fm6(globalState), 6)
		(_nw90).ArraySet1(_550_v2, 7)
		(_nw90).ArraySet1(_dafny.IntOfUint32((_553_v5).Cardinality()), 8)
		(_nw90).ArraySet1(_550_v2, 9)
		(_nw90).ArraySet1(_550_v2, 10)
		(_nw90).ArraySet1(_550_v2, 11)
		(_nw90).ArraySet1(_dafny.IntOfUint32((_552_v4).Cardinality()), 12)
		(_nw90).ArraySet1(_dafny.IntOfUint32((_dafny.SeqOf(_550_v2, _550_v2)).Cardinality()), 13)
		(_nw90).ArraySet1(_550_v2, 14)
		(_nw90).ArraySet1(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bjayaggc")).Cardinality()), 15)
		(_nw90).ArraySet1(_550_v2, 16)
		(_nw90).ArraySet1(_dafny.IntOfUint32((Companion_Default___.Fm8(_550_v2, false, _552_v4, globalState)).Cardinality()), 17)
		_554_v6 = _nw90
		var _555_v7 _dafny.Map
		_ = _555_v7
		_555_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_551_v3, _554_v6)
		var _556_v8 _dafny.Map
		_ = _556_v8
		_556_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_550_v2, (_555_v7).Cardinality())
		var _557_v9 _dafny.Set
		_ = _557_v9
		_557_v9 = _dafny.SetOf((_this).Fm4(_550_v2, _550_v2, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_552_v4, _this.F8()), _550_v2, globalState), _550_v2)
		var _558_v10 _dafny.Array
		_ = _558_v10
		var _nwElement0_15 _dafny.Int = _550_v2
		_ = _nwElement0_15
		var _nw91 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(27))
		_ = _nw91
		(_nw91).ArraySet1(_nwElement0_15, 0)
		(_nw91).ArraySet1((_550_v2).Times((_dafny.Zero).Minus(_550_v2)), 1)
		(_nw91).ArraySet1((_556_v8).Cardinality(), 2)
		(_nw91).ArraySet1(_550_v2, 3)
		(_nw91).ArraySet1(_dafny.IntOfInt64(737), 4)
		(_nw91).ArraySet1(_550_v2, 5)
		(_nw91).ArraySet1(_dafny.IntOfInt64(115), 6)
		(_nw91).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_552_v4).Cardinality()), _550_v2), 7)
		(_nw91).ArraySet1(_550_v2, 8)
		(_nw91).ArraySet1(_550_v2, 9)
		(_nw91).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_557_v9).Cardinality(), _550_v2)).Cardinality(), 10)
		(_nw91).ArraySet1(_dafny.IntOfInt64(-640), 11)
		(_nw91).ArraySet1((_dafny.Zero).Minus((_550_v2).Plus(_550_v2)), 12)
		(_nw91).ArraySet1(_550_v2, 13)
		(_nw91).ArraySet1(_550_v2, 14)
		(_nw91).ArraySet1(Companion_Default___.SafeDivisionInt(_550_v2, _550_v2), 15)
		(_nw91).ArraySet1(_550_v2, 16)
		(_nw91).ArraySet1(((_this).F5()).Cardinality(), 17)
		(_nw91).ArraySet1(_550_v2, 18)
		(_nw91).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), _550_v2)).Cardinality(), 19)
		(_nw91).ArraySet1(_550_v2, 20)
		(_nw91).ArraySet1((Companion_Default___.Fm9(globalState)).Cardinality(), 21)
		(_nw91).ArraySet1(_550_v2, 22)
		(_nw91).ArraySet1(_550_v2, 23)
		(_nw91).ArraySet1((_dafny.Zero).Minus(_550_v2), 24)
		(_nw91).ArraySet1(_dafny.IntOfInt64(518), 25)
		(_nw91).ArraySet1(_550_v2, 26)
		_558_v10 = _nw91
		_558_v10 = _554_v6
		var _559_i0 _dafny.Int
		_ = _559_i0
		_559_i0 = _dafny.Zero
		{
			for (_550_v2).Cmp((_dafny.IntOfUint32((_552_v4).Cardinality())).Minus((_dafny.Zero).Minus(_550_v2))) != 0 {
				{
					if (_559_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L7
					}
					_559_i0 = (_559_i0).Plus(_dafny.One)
					if (_this).F6() {
						var _560_v12 _dafny.Map
						_ = _560_v12
						_560_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_551_v3, p0)
						var _561_v13 _dafny.MultiSet
						_ = _561_v13
						_561_v13 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_550_v2, (_560_v12).Cardinality()), _556_v8)
						var _562_v15 _dafny.Sequence
						_ = _562_v15
						_562_v15 = _dafny.SeqOf(_556_v8, _556_v8)
						var _563_v16 *C0
						_ = _563_v16
						var _nw92 *C0 = New_C0_()
						_ = _nw92
						_nw92.Ctor__((func() _dafny.Map {
							if !(!(_this.F8())) {
								return func() _dafny.Map {
									var _coll31 = _dafny.NewMapBuilder()
									_ = _coll31
									for _iter34 := _dafny.Iterate((_561_v13).Elements()); ; {
										_compr_31, _ok34 := _iter34()
										if !_ok34 {
											break
										}
										var _564_v11 _dafny.Map
										_564_v11 = interface{}(_compr_31).(_dafny.Map)
										if (_561_v13).Contains(_564_v11) {
											_coll31.Add(_564_v11, false)
										}
									}
									return _coll31.ToMap()
								}()
							}
							return func() _dafny.Map {
								var _coll32 = _dafny.NewMapBuilder()
								_ = _coll32
								for _iter35 := _dafny.Iterate((_562_v15).Elements()); ; {
									_compr_32, _ok35 := _iter35()
									if !_ok35 {
										break
									}
									var _565_v14 _dafny.Map
									_565_v14 = interface{}(_compr_32).(_dafny.Map)
									if _dafny.Companion_Sequence_.Contains(_562_v15, _565_v14) {
										_coll32.Add(_565_v14, _this.F8())
									}
								}
								return _coll32.ToMap()
							}()
						})())
						_563_v16 = _nw92
						var _566_v17 T1
						_ = _566_v17
						var _nw93 *C3 = New_C3_()
						_ = _nw93
						_nw93.Ctor__((_this).F6(), _dafny.MultiSetOf((_this).F6()))
						_566_v17 = _nw93
						var _567_v18 _dafny.Set
						_ = _567_v18
						_567_v18 = _dafny.SetOf(_566_v17, _566_v17)
						var _arr0 _dafny.Array = _this.F7()
						_ = _arr0
						var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_this.F7()), 0))
						_ = _index96
						(_arr0).ArraySet1(_this.F8(), (_index96).Int())
						var _568_v19 _dafny.Map
						_ = _568_v19
						_568_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(913), !((func() bool {
							if false {
								return (_566_v17).F6()
							}
							return (_this).F6()
						})()))
						var _569_v20 _dafny.Sequence
						_ = _569_v20
						_569_v20 = _dafny.SeqOf(_550_v2)
						var _arr1 _dafny.Array = _this.F7()
						_ = _arr1
						var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_this.F7()), 0))
						_ = _index97
						var _rhs105 _dafny.Set = _dafny.SetOf(_566_v17, _566_v17, _566_v17)
						_ = _rhs105
						var _rhs106 bool = (func() bool {
							if (_568_v19).Contains(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_552_v4, _dafny.Companion_Sequence_.Update(_552_v4, (Companion_Default___.SafeIndex(_550_v2, _dafny.IntOfUint32((_552_v4).Cardinality()))).Uint32(), _548_v0))).Cardinality())) {
								return (_568_v19).Get(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_552_v4, _dafny.Companion_Sequence_.Update(_552_v4, (Companion_Default___.SafeIndex(_550_v2, _dafny.IntOfUint32((_552_v4).Cardinality()))).Uint32(), _548_v0))).Cardinality())).(bool)
							}
							return (_550_v2).Cmp(_dafny.IntOfUint32((_569_v20).Cardinality())) <= 0
						})()
						_ = _rhs106
						var _rhs107 bool = (_566_v17).F6()
						_ = _rhs107
						var _rhs108 D0 = _551_v3
						_ = _rhs108
						var _lhs74 *C4 = _this
						_ = _lhs74
						var _lhs75 _dafny.Array = _this.F7()
						_ = _lhs75
						var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_this.F7()), 0))
						_ = _lhs76
						_567_v18 = _rhs105
						_lhs74.F8_set_(_rhs106)
						(_lhs75).ArraySet1(_rhs107, (_lhs76).Int())
						_551_v3 = _rhs108
						var _arr2 _dafny.Array = _this.F7()
						_ = _arr2
						var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_this.F7()), 0))
						_ = _index98
						(_arr2).ArraySet1((_566_v17).F6(), (_index98).Int())
						_550_v2 = _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ulmddp")).Cardinality())
						var _570_v21 _dafny.Map
						_ = _570_v21
						_570_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_552_v4, _this.F8())
						(globalState).F0 = (_this).Fm4(_550_v2, _550_v2, _570_v21, _550_v2, globalState)
					} else {
						var _571_v22 _dafny.MultiSet
						_ = _571_v22
						_571_v22 = _dafny.MultiSetOf(_dafny.IntOfUint32((_552_v4).Cardinality()))
						(_this).F8_set_((_571_v22).IsSubsetOf(_571_v22))
						var _arr3 _dafny.Array = _this.F7()
						_ = _arr3
						var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_this.F7()), 0))
						_ = _index99
						(_arr3).ArraySet1((_this).F6(), (_index99).Int())
						var _572_v23 _dafny.Sequence
						_ = _572_v23
						_572_v23 = _dafny.SeqOf(_dafny.IntOfUint32((_552_v4).Cardinality()))
						var _arr4 _dafny.Array = _this.F7()
						_ = _arr4
						var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_this.F7()), 0))
						_ = _index100
						var _rhs109 _dafny.Int = (_550_v2).Minus(_550_v2)
						_ = _rhs109
						var _rhs110 _dafny.CodePoint = _dafny.CodePoint('g')
						_ = _rhs110
						var _rhs111 bool = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_572_v23).Cardinality()), ((_this).F5()).Cardinality())).Cmp(_550_v2) == 0
						_ = _rhs111
						var _rhs112 _dafny.Array = _this.F7()
						_ = _rhs112
						var _rhs113 _dafny.Int = (_dafny.Zero).Minus((Companion_Default___.Fm19((_dafny.Zero).Minus(_550_v2), globalState)).Cardinality())
						_ = _rhs113
						var _lhs77 *GlobalState = globalState
						_ = _lhs77
						var _lhs78 _dafny.Array = _this.F7()
						_ = _lhs78
						var _lhs79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(393), _dafny.ArrayLen((_this.F7()), 0))
						_ = _lhs79
						var _lhs80 *C4 = _this
						_ = _lhs80
						_lhs77.F0 = _rhs109
						_548_v0 = _rhs110
						(_lhs78).ArraySet1(_rhs111, (_lhs79).Int())
						_lhs80.F7_set_(_rhs112)
						_550_v2 = _rhs113
						var _573_v24 _dafny.Map
						_ = _573_v24
						_573_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_550_v2, p0)
						_548_v0 = (func() _dafny.CodePoint {
							if (_573_v24).Contains((_dafny.Zero).Minus(_550_v2)) {
								return (_573_v24).Get((_dafny.Zero).Minus(_550_v2)).(_dafny.CodePoint)
							}
							return p0
						})()
						var _574_v25 _dafny.Map
						_ = _574_v25
						_574_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm23(_550_v2, globalState), (_this).Fm3(globalState))
						_574_v25 = _574_v25
						var _rhs114 _dafny.MultiSet = (_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_572_v23, _572_v23))).Union(_571_v22)
						_ = _rhs114
						var _rhs115 _dafny.Int = (_dafny.Zero).Minus((_550_v2).Times(_550_v2))
						_ = _rhs115
						var _lhs81 *GlobalState = globalState
						_ = _lhs81
						_571_v22 = _rhs114
						_lhs81.F0 = _rhs115
					}
					(globalState).F0 = (((func() _dafny.MultiSet {
						if (_this).F6() {
							return _dafny.MultiSetOf((_this).F6())
						}
						return (_this).F5()
					})()).Cardinality()).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ckkg")).Cardinality()))
					var _575_v26 _dafny.Map
					_ = _575_v26
					_575_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), _550_v2)
					var _576_v27 D7
					_ = _576_v27
					_576_v27 = Companion_D7_.Create_DC19_(_575_v26)
					_576_v27 = Companion_D7_.Create_DC19_(_575_v26)
					var _577_v28 *C2
					_ = _577_v28
					var _nw94 *C2 = New_C2_()
					_ = _nw94
					_nw94.Ctor__((_this).F5())
					_577_v28 = _nw94
					var _578_v29 D9
					_ = _578_v29
					_578_v29 = Companion_D9_.Create_DC26_(_577_v28)
					_577_v28 = (_578_v29).Dtor_cf41()
					goto C7
				}
			C7:
			}
			goto L7
		}
	L7:
		var _579_v30 D3
		_ = _579_v30
		_579_v30 = Companion_D3_.Create_DC10_(false, _550_v2)
		var _580_v31 *C1
		_ = _580_v31
		var _nw95 *C1 = New_C1_()
		_ = _nw95
		_nw95.Ctor__(_this.F8(), Companion_Default___.Fm17(globalState), (_this).Fm1((_579_v30).Dtor_cf11(), _550_v2, _dafny.IntOfUint32((_552_v4).Cardinality()), globalState), ((_this).F5()).Union(_dafny.MultiSetOf(!(_this.F8()))))
		_580_v31 = _nw95
		_558_v10 = _558_v10
	}
}
func (_this *C4) M3(globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		(_this).F8_set_((_this).F6())
		if true {
			var _arr5 _dafny.Array = _this.F7()
			_ = _arr5
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index101
			(_arr5).ArraySet1(true, (_index101).Int())
			var _581_v0 _dafny.Array
			_ = _581_v0
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(4)
			_ = _len0_14
			var _nw96 _dafny.Array
			_ = _nw96
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw96 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) _dafny.Int = func(_582_i0 _dafny.Int) _dafny.Int {
					return (_582_i0).Plus(_dafny.IntOfInt64(-680))
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw96 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw96).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw96).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_581_v0 = _nw96
			var _583_v1 _dafny.Int
			_ = _583_v1
			_583_v1 = _dafny.IntOfInt64(345)
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_581_v0), 0))
			_ = _index102
			(_581_v0).ArraySet1((_583_v1).Times(_583_v1), (_index102).Int())
			var _584_v2 _dafny.Sequence
			_ = _584_v2
			_584_v2 = _dafny.SeqOf((_this).Fm2((_this).F6(), globalState), (_this).F6(), (_this).F6(), _this.F8(), _this.F8())
			var _arr6 _dafny.Array = _this.F7()
			_ = _arr6
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index103
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_581_v0), 0))
			_ = _index104
			var _rhs116 _dafny.Array = _this.F7()
			_ = _rhs116
			var _rhs117 bool = (_this).Fm3(globalState)
			_ = _rhs117
			var _rhs118 bool = (_dafny.IntOfUint32((_584_v2).Cardinality())).Cmp((_this).Fm6(globalState)) == 0
			_ = _rhs118
			var _rhs119 _dafny.Int = _583_v1
			_ = _rhs119
			var _lhs82 *C4 = _this
			_ = _lhs82
			var _lhs83 *C4 = _this
			_ = _lhs83
			var _lhs84 _dafny.Array = _this.F7()
			_ = _lhs84
			var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_this.F7()), 0))
			_ = _lhs85
			var _lhs86 _dafny.Array = _581_v0
			_ = _lhs86
			var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_581_v0), 0))
			_ = _lhs87
			_lhs82.F7_set_(_rhs116)
			_lhs83.F8_set_(_rhs117)
			(_lhs84).ArraySet1(_rhs118, (_lhs85).Int())
			(_lhs86).ArraySet1(_rhs119, (_lhs87).Int())
			_584_v2 = _584_v2
			var _585_v3 _dafny.Array
			_ = _585_v3
			var _len0_15 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_15
			var _nw97 _dafny.Array
			_ = _nw97
			if _len0_15.Cmp(_dafny.Zero) == 0 {
				_nw97 = _dafny.NewArray(_len0_15)
			} else {
				var _init15 func(_dafny.Int) bool = func(_586_i1 _dafny.Int) bool {
					return false
				}
				_ = _init15
				var _element0_15 = _init15(_dafny.Zero)
				_ = _element0_15
				_nw97 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
				(_nw97).ArraySet1(_element0_15, 0)
				var _nativeLen0_15 = (_len0_15).Int()
				_ = _nativeLen0_15
				for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
					(_nw97).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
				}
			}
			_585_v3 = _nw97
			var _587_v4 _dafny.MultiSet
			_ = _587_v4
			_587_v4 = _dafny.MultiSetOf(_585_v3)
			var _588_v5 _dafny.Map
			_ = _588_v5
			_588_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(232), _this.F8())
			var _589_v6 _dafny.Map
			_ = _589_v6
			_589_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_587_v4, _588_v5)
			_589_v6 = (_589_v6).Update(_587_v4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_583_v1, (_this).F6()))
			var _590_v7 _dafny.Array
			_ = _590_v7
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_16
			var _nw98 _dafny.Array
			_ = _nw98
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw98 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.CodePoint = func(_591_i2 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('c')
				}
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw98 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw98).ArraySet1CodePoint(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw98).ArraySet1CodePoint(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_590_v7 = _nw98
			(globalState).F2 = _590_v7
			var _592_v8 _dafny.Sequence
			_ = _592_v8
			_592_v8 = _dafny.SeqOf((_581_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_581_v0), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_584_v2).Cardinality()), _583_v1, _dafny.IntOfInt64(778))
			var _arr7 _dafny.Array = _this.F7()
			_ = _arr7
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index105
			var _rhs120 _dafny.Array = _581_v0
			_ = _rhs120
			var _rhs121 _dafny.Int = (_581_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_581_v0), 0))).Int()).(_dafny.Int)
			_ = _rhs121
			var _rhs122 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_592_v8, _592_v8), _592_v8)
			_ = _rhs122
			var _rhs123 bool = (_this).F6()
			_ = _rhs123
			var _lhs88 *GlobalState = globalState
			_ = _lhs88
			var _lhs89 _dafny.Array = _this.F7()
			_ = _lhs89
			var _lhs90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(189), _dafny.ArrayLen((_this.F7()), 0))
			_ = _lhs90
			_581_v0 = _rhs120
			_lhs88.F0 = _rhs121
			_592_v8 = _rhs122
			(_lhs89).ArraySet1(_rhs123, (_lhs90).Int())
		} else {
			var _593_v9 *C2
			_ = _593_v9
			var _nw99 *C2 = New_C2_()
			_ = _nw99
			_nw99.Ctor__((_this).F5())
			_593_v9 = _nw99
			(_this).F8_set_(false)
			var _594_v10 _dafny.CodePoint
			_ = _594_v10
			_594_v10 = _dafny.CodePoint('u')
			_594_v10 = _594_v10
			var _595_v11 _dafny.Int
			_ = _595_v11
			_595_v11 = _dafny.IntOfInt64(442)
			(globalState).F0 = _595_v11
			if !(false) {
				var _596_v12 _dafny.Sequence
				_ = _596_v12
				_596_v12 = _dafny.UnicodeSeqOfUtf8Bytes("lh")
				var _597_v13 _dafny.Sequence
				_ = _597_v13
				_597_v13 = _dafny.SeqOf(_596_v12)
				_596_v12 = _dafny.Companion_Sequence_.Update((_597_v13).Select((Companion_Default___.SafeIndex(_595_v11, _dafny.IntOfUint32((_597_v13).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_this).Fm6(globalState), _dafny.IntOfUint32(((_597_v13).Select((Companion_Default___.SafeIndex(_595_v11, _dafny.IntOfUint32((_597_v13).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _594_v10)
				_596_v12 = _596_v12
				var _598_v14 _dafny.Array
				_ = _598_v14
				var _nw100 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
				_ = _nw100
				_598_v14 = _nw100
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_598_v14), 0))
				_ = _index106
				(_598_v14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_596_v12, _596_v12), (_index106).Int())
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(284), _dafny.ArrayLen((_598_v14), 0))
				_ = _index107
				(_598_v14).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("px"), _596_v12), (Companion_Default___.SafeIndex(_595_v11, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("px"), _596_v12)).Cardinality()))).Uint32(), _594_v10), _596_v12), (_index107).Int())
				_594_v10 = _594_v10
				var _599_v15 _dafny.Sequence
				_ = _599_v15
				_599_v15 = _dafny.SeqOf((_this).F6())
				var _600_v16 _dafny.MultiSet
				_ = _600_v16
				_600_v16 = _dafny.MultiSetOf(_599_v15)
				var _601_v17 D10
				_ = _601_v17
				_601_v17 = Companion_D10_.Create_DC29_(_600_v16)
				var _602_v18 _dafny.Array
				_ = _602_v18
				var _nwElement0_16 _dafny.MultiSet = (_600_v16).Difference(_600_v16)
				_ = _nwElement0_16
				var _nw101 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(17))
				_ = _nw101
				(_nw101).ArraySet1(_nwElement0_16, 0)
				(_nw101).ArraySet1(_dafny.MultiSetOf(_599_v15, _dafny.SeqOf((_this).F6())), 1)
				(_nw101).ArraySet1(_600_v16, 2)
				(_nw101).ArraySet1(_dafny.MultiSetOf(_599_v15), 3)
				(_nw101).ArraySet1(_600_v16, 4)
				(_nw101).ArraySet1(_600_v16, 5)
				(_nw101).ArraySet1(_600_v16, 6)
				(_nw101).ArraySet1((_dafny.MultiSetOf(_599_v15)).Union(_600_v16), 7)
				(_nw101).ArraySet1((_600_v16).Update(_599_v15, Companion_Default___.Abs(_595_v11)), 8)
				(_nw101).ArraySet1((_600_v16).Union((_601_v17).Dtor_cf45()), 9)
				(_nw101).ArraySet1(_600_v16, 10)
				(_nw101).ArraySet1((_600_v16).Intersection(_dafny.MultiSetOf(_599_v15, _599_v15, _599_v15, _599_v15)), 11)
				(_nw101).ArraySet1((_600_v16).Difference(_600_v16), 12)
				(_nw101).ArraySet1(_600_v16, 13)
				(_nw101).ArraySet1(_600_v16, 14)
				(_nw101).ArraySet1(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(203))).Uint32(), func(coer32 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg32 _dafny.Int) interface{} {
						return coer32(arg32)
					}
				}(func(_603_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_this.F8(), (_this).F6(), true)
				}))), 15)
				(_nw101).ArraySet1(_600_v16, 16)
				_602_v18 = _nw101
				var _604_v19 _dafny.Sequence
				_ = _604_v19
				_604_v19 = _dafny.SeqOf(_599_v15)
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_602_v18), 0))
				_ = _index108
				(_602_v18).ArraySet1(_dafny.MultiSetFromSeq(_604_v19), (_index108).Int())
				var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(423), _dafny.ArrayLen((_602_v18), 0))
				_ = _index109
				(_602_v18).ArraySet1(_dafny.MultiSetOf(_599_v15), (_index109).Int())
			} else {
				var _arr8 _dafny.Array = _this.F7()
				_ = _arr8
				var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index110
				(_arr8).ArraySet1(_this.F8(), (_index110).Int())
				var _605_v21 D2
				_ = _605_v21
				_605_v21 = Companion_D2_.Create_DC4_(func() _dafny.Set {
					var _coll33 = _dafny.NewBuilder()
					_ = _coll33
					for _iter36 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(851), _dafny.IntOfInt64(-711))); ; {
						_compr_33, _ok36 := _iter36()
						if !_ok36 {
							break
						}
						var _606_v20 _dafny.Int
						_606_v20 = interface{}(_compr_33).(_dafny.Int)
						if ((_dafny.IntOfInt64(851)).Cmp(_606_v20) <= 0) && ((_606_v20).Cmp(_dafny.IntOfInt64(-711)) < 0) {
							_coll33.Add((_606_v20).Times(_595_v11))
						}
					}
					return _coll33.ToSet()
				}())
				var _607_v22 _dafny.Sequence
				_ = _607_v22
				_607_v22 = _dafny.SeqOf(_595_v11, _595_v11)
				var _608_v23 _dafny.Set
				_ = _608_v23
				_608_v23 = _dafny.SetOf(_dafny.IntOfUint32((_607_v22).Cardinality()), _595_v11)
				var _arr9 _dafny.Array = _this.F7()
				_ = _arr9
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index111
				(_arr9).ArraySet1((_608_v23).IsSubsetOf((_605_v21).Dtor_cf4()), (_index111).Int())
				var _609_v24 _dafny.Sequence
				_ = _609_v24
				_609_v24 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("kypgb"))
				var _610_v25 _dafny.Sequence
				_ = _610_v25
				_610_v25 = _dafny.SeqOf(true)
				var _611_v26 _dafny.Map
				_ = _611_v26
				_611_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_610_v25, (Companion_Default___.SafeIndex(_595_v11, _dafny.IntOfUint32((_610_v25).Cardinality()))).Uint32(), _this.F8()), _dafny.Companion_Sequence_.Concatenate(_609_v24, _609_v24))
				_609_v24 = (func() _dafny.Sequence {
					if (_611_v26).Contains(_610_v25) {
						return (_611_v26).Get(_610_v25).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(701))).Uint32(), func(coer33 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
						return func(arg33 _dafny.Int) interface{} {
							return coer33(arg33)
						}
					}(func(_612_i4 _dafny.Int) _dafny.Sequence {
						return _dafny.UnicodeSeqOfUtf8Bytes("ykalvw")
					}))
				})()
				_609_v24 = _609_v24
				var _613_v27 _dafny.Map
				_ = _613_v27
				_613_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(240), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)), _this.F8())
				_613_v27 = (_613_v27).Update((_this).F5(), ((_dafny.Zero).Minus(_595_v11)).Cmp(_595_v11) < 0)
				var _614_v28 _dafny.Map
				_ = _614_v28
				var _out8 _dafny.Map
				_ = _out8
				_out8 = (_593_v9).M9(!((_this).F6()), _595_v11, globalState)
				_614_v28 = _out8
			}
		}
		var _615_v29 _dafny.Sequence
		_ = _615_v29
		_615_v29 = _dafny.SeqOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(250))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_616_i5 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('m')
		})))
		var _arr10 _dafny.Array = _this.F7()
		_ = _arr10
		var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index112
		(_arr10).ArraySet1(_dafny.Companion_Sequence_.Equal(_615_v29, _615_v29), (_index112).Int())
		var _arr11 _dafny.Array = _this.F7()
		_ = _arr11
		var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index113
		(_arr11).ArraySet1(!((_this).Fm3(globalState)), (_index113).Int())
		var _617_v30 _dafny.Int
		_ = _617_v30
		_617_v30 = _dafny.IntOfInt64(967)
		var _hi2 _dafny.Int = _617_v30
		_ = _hi2
		for _618_i6 := _617_v30; _618_i6.Cmp(_hi2) < 0; _618_i6 = _618_i6.Plus(_dafny.One) {
			var _619_v31 _dafny.CodePoint
			_ = _619_v31
			_619_v31 = _dafny.CodePoint('g')
			var _620_v32 _dafny.Map
			_ = _620_v32
			_620_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_617_v30, ((_this).F5()).Cardinality())
			var _arr12 _dafny.Array = _this.F7()
			_ = _arr12
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index114
			(_arr12).ArraySet1(((_618_i6).Plus((Companion_D2_.Create_DC7_(_619_v31, _620_v32, _dafny.IntOfInt64(-433))).Dtor_cf8())).Cmp(_617_v30) >= 0, (_index114).Int())
			var _621_v33 _dafny.Array
			_ = _621_v33
			var _len0_17 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_17
			var _nw102 _dafny.Array
			_ = _nw102
			if _len0_17.Cmp(_dafny.Zero) == 0 {
				_nw102 = _dafny.NewArray(_len0_17)
			} else {
				var _init17 func(_dafny.Int) _dafny.Int = (func(_622_i6 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_623_i7 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeModuloInt(_623_i7, _622_i6)
					}
				})(_618_i6)
				_ = _init17
				var _element0_17 = _init17(_dafny.Zero)
				_ = _element0_17
				_nw102 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
				(_nw102).ArraySet1(_element0_17, 0)
				var _nativeLen0_17 = (_len0_17).Int()
				_ = _nativeLen0_17
				for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
					(_nw102).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
				}
			}
			_621_v33 = _nw102
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_621_v33), 0))
			_ = _index115
			(_621_v33).ArraySet1(_dafny.IntOfInt64(728), (_index115).Int())
			var _arr13 _dafny.Array = _this.F7()
			_ = _arr13
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index116
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_621_v33), 0))
			_ = _index117
			var _rhs124 bool = _this.F8()
			_ = _rhs124
			var _rhs125 _dafny.Int = (_618_i6).Times(_618_i6)
			_ = _rhs125
			var _lhs91 _dafny.Array = _this.F7()
			_ = _lhs91
			var _lhs92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
			_ = _lhs92
			var _lhs93 _dafny.Array = _621_v33
			_ = _lhs93
			var _lhs94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(198), _dafny.ArrayLen((_621_v33), 0))
			_ = _lhs94
			(_lhs91).ArraySet1(_rhs124, (_lhs92).Int())
			(_lhs93).ArraySet1(_rhs125, (_lhs94).Int())
			var _arr14 _dafny.Array = _this.F7()
			_ = _arr14
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index118
			var _rhs126 bool = (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)
			_ = _rhs126
			var _rhs127 bool = (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)
			_ = _rhs127
			var _rhs128 _dafny.Sequence = (func() _dafny.Sequence {
				if _this.F8() {
					return _615_v29
				}
				return _615_v29
			})()
			_ = _rhs128
			var _lhs95 *C4 = _this
			_ = _lhs95
			var _lhs96 _dafny.Array = _this.F7()
			_ = _lhs96
			var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
			_ = _lhs97
			_lhs95.F8_set_(_rhs126)
			(_lhs96).ArraySet1(_rhs127, (_lhs97).Int())
			_615_v29 = _rhs128
			var _624_v34 _dafny.Sequence
			_ = _624_v34
			_624_v34 = _dafny.SeqOf(_617_v30)
			var _625_v35 _dafny.Sequence
			_ = _625_v35
			_625_v35 = _dafny.SeqOf(_624_v34)
			var _626_v36 _dafny.Sequence
			_ = _626_v36
			_626_v36 = _dafny.UnicodeSeqOfUtf8Bytes("hj")
			var _627_v37 _dafny.Map
			_ = _627_v37
			_627_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_625_v35).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_626_v36).Cardinality()), _dafny.IntOfUint32((_625_v35).Cardinality()))).Uint32()).(_dafny.Sequence), false)
			_627_v37 = _627_v37
		}
		if false {
			(globalState).F0 = _617_v30
			var _628_v38 _dafny.Array
			_ = _628_v38
			var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
			_ = _nw103
			_628_v38 = _nw103
			var _629_v39 _dafny.Map
			_ = _629_v39
			_629_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_628_v38, (_this).F6())
			var _630_v40 D9
			_ = _630_v40
			_630_v40 = Companion_D9_.Create_DC28_(Companion_D9_.Create_DC27_(_629_v39, (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)))
			_630_v40 = _630_v40
			(_this).F8_set_(_this.F8())
			(globalState).F0 = (func() _dafny.Int {
				if (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool) {
					return _dafny.IntOfInt64(560)
				}
				return _dafny.IntOfInt64(789)
			})()
			var _631_v41 _dafny.MultiSet
			_ = _631_v41
			_631_v41 = _dafny.MultiSetOf(_617_v30, (_this).Fm6(globalState), (_this).Fm6(globalState))
			var _632_v42 _dafny.Set
			_ = _632_v42
			_632_v42 = _dafny.SetOf(_617_v30)
			var _633_v43 _dafny.Sequence
			_ = _633_v43
			_633_v43 = _dafny.UnicodeSeqOfUtf8Bytes("gtark")
			var _634_v44 _dafny.Map
			_ = _634_v44
			_634_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_633_v43, (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))
			var _635_v45 _dafny.Sequence
			_ = _635_v45
			_635_v45 = _dafny.SeqOf(_617_v30, (_this).Fm4(_617_v30, _617_v30, _634_v44, _617_v30, globalState))
			var _636_v46 _dafny.Map
			_ = _636_v46
			_636_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(_617_v30), _this.F8())
			var _637_v47 _dafny.Map
			_ = _637_v47
			_637_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_636_v46).Cardinality(), _617_v30)
			var _638_v48 _dafny.Array
			_ = _638_v48
			var _nwElement0_17 _dafny.Int = _617_v30
			_ = _nwElement0_17
			var _nw104 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(28))
			_ = _nw104
			(_nw104).ArraySet1(_nwElement0_17, 0)
			(_nw104).ArraySet1(_617_v30, 1)
			(_nw104).ArraySet1(_617_v30, 2)
			(_nw104).ArraySet1(_617_v30, 3)
			(_nw104).ArraySet1(_617_v30, 4)
			(_nw104).ArraySet1((_631_v41).Cardinality(), 5)
			(_nw104).ArraySet1(_617_v30, 6)
			(_nw104).ArraySet1(_dafny.IntOfInt64(433), 7)
			(_nw104).ArraySet1(_617_v30, 8)
			(_nw104).ArraySet1((Companion_Default___.Fm15((_this).F6(), globalState)).Cardinality(), 9)
			(_nw104).ArraySet1(_617_v30, 10)
			(_nw104).ArraySet1((_632_v42).Cardinality(), 11)
			(_nw104).ArraySet1(_617_v30, 12)
			(_nw104).ArraySet1(_dafny.IntOfUint32((_635_v45).Cardinality()), 13)
			(_nw104).ArraySet1(_617_v30, 14)
			(_nw104).ArraySet1(_617_v30, 15)
			(_nw104).ArraySet1(_617_v30, 16)
			(_nw104).ArraySet1(_617_v30, 17)
			(_nw104).ArraySet1(_617_v30, 18)
			(_nw104).ArraySet1(_617_v30, 19)
			(_nw104).ArraySet1(_617_v30, 20)
			(_nw104).ArraySet1((((_this).F5()).Update(false, Companion_Default___.Abs((_636_v46).Cardinality()))).Cardinality(), 21)
			(_nw104).ArraySet1(_617_v30, 22)
			(_nw104).ArraySet1((_637_v47).Cardinality(), 23)
			(_nw104).ArraySet1(_617_v30, 24)
			(_nw104).ArraySet1(_617_v30, 25)
			(_nw104).ArraySet1(_617_v30, 26)
			(_nw104).ArraySet1(_617_v30, 27)
			_638_v48 = _nw104
			var _639_v49 _dafny.Sequence
			_ = _639_v49
			_639_v49 = _dafny.SeqOf(_638_v48)
			_639_v49 = _dafny.SeqOf(_638_v48, _638_v48, _638_v48)
		} else {
			var _640_v50 _dafny.Map
			_ = _640_v50
			_640_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_617_v30).Cmp(_617_v30) > 0, (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))
			_640_v50 = _640_v50
			var _641_v51 _dafny.Sequence
			_ = _641_v51
			_641_v51 = _dafny.UnicodeSeqOfUtf8Bytes("ccqgh")
			var _642_v52 _dafny.Set
			_ = _642_v52
			_642_v52 = _dafny.SetOf(_641_v51)
			var _arr15 _dafny.Array = _this.F7()
			_ = _arr15
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index119
			(_arr15).ArraySet1((func() _dafny.Set {
				var _coll34 = _dafny.NewBuilder()
				_ = _coll34
				for _iter37 := _dafny.Iterate((_615_v29).Elements()); ; {
					_compr_34, _ok37 := _iter37()
					if !_ok37 {
						break
					}
					var _643_v53 _dafny.Sequence
					_643_v53 = interface{}(_compr_34).(_dafny.Sequence)
					if _dafny.Companion_Sequence_.Contains(_615_v29, _643_v53) {
						_coll34.Add(_643_v53)
					}
				}
				return _coll34.ToSet()
			}()).IsSubsetOf(_642_v52), (_index119).Int())
			(globalState).F0 = _617_v30
			var _644_v54 _dafny.Map
			_ = _644_v54
			_644_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!((_this).Fm3(globalState)))), _dafny.UnicodeSeqOfUtf8Bytes("pa"))
			_644_v54 = (_644_v54).Update((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), _dafny.UnicodeSeqOfUtf8Bytes("qc"))
			_617_v30 = Companion_Default___.SafeModuloInt(_617_v30, (_617_v30).Times(_dafny.IntOfInt64(167)))
		}
		var _645_v55 _dafny.Sequence
		_ = _645_v55
		_645_v55 = _dafny.UnicodeSeqOfUtf8Bytes("djxkfjewt")
		var _646_v56 _dafny.Sequence
		_ = _646_v56
		_646_v56 = _dafny.SeqOf(_617_v30, _dafny.IntOfUint32((_645_v55).Cardinality()))
		var _647_v57 _dafny.Map
		_ = _647_v57
		_647_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_646_v56, _617_v30)
		_647_v57 = (_647_v57).Update(_dafny.SeqOf(_617_v30, _617_v30, _dafny.IntOfUint32((_645_v55).Cardinality())), (_617_v30).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))).Cardinality()))))
		var _648_v58 _dafny.Set
		_ = _648_v58
		_648_v58 = _dafny.SetOf(_617_v30, _617_v30, _617_v30, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(537))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_649_i8 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('a')
		}))).Cardinality()))
		var _650_v59 _dafny.Map
		_ = _650_v59
		_650_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_this.F8()), _this.F8())
		var _651_v60 _dafny.Map
		_ = _651_v60
		_651_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("kdtsapr"), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(183), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))
		r0 = (_648_v58).Intersection((_dafny.SetOf(_617_v30)).Intersection(Companion_Default___.Fm24((_this).Fm4(_617_v30, (_650_v59).Cardinality(), _651_v60, _617_v30, globalState), (_this).F6(), _dafny.IntOfInt64(111), globalState)))
		r1 = (_617_v30).Minus(_617_v30)
		var _652_v61 _dafny.CodePoint
		_ = _652_v61
		_652_v61 = _dafny.CodePoint('r')
		var _653_v62 _dafny.Map
		_ = _653_v62
		_653_v62 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_652_v61, _this.F8())
		r2 = Companion_Default___.SafeModuloInt(_617_v30, ((_653_v62).Update(_652_v61, (_this).Fm2(_this.F8(), globalState))).Cardinality())
		return r0, r1, r2
	}
}
func (_this *C4) M1(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) {
	{
		(_this).F8_set_(true)
		var _654_v0 _dafny.CodePoint
		_ = _654_v0
		_654_v0 = _dafny.CodePoint('t')
		var _655_v1 D7
		_ = _655_v1
		_655_v1 = Companion_D7_.Create_DC21_(p0, p0)
		var _656_v2 _dafny.Map
		_ = _656_v2
		_656_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_655_v1).Dtor_cf30(), p0)
		var _657_v3 D2
		_ = _657_v3
		_657_v3 = Companion_D2_.Create_DC7_(_654_v0, _656_v2, p0)
		var _658_i0 _dafny.Int
		_ = _658_i0
		_658_i0 = _dafny.Zero
		{
			var _pat_let_tv16 = p0
			_ = _pat_let_tv16
			var _pat_let_tv17 = p0
			_ = _pat_let_tv17
			var _pat_let_tv18 = p0
			_ = _pat_let_tv18
			for func(_source5 D2) bool {
				if _source5.Is_DC5() {
					var _664___mcc_h0 bool = _source5.Get_().(D2_DC5).Cf5
					_ = _664___mcc_h0
					var _665_cf5 bool = _664___mcc_h0
					_ = _665_cf5
					return _665_cf5
				} else if _source5.Is_DC6() {
					return (_this).F6()
				} else if _source5.Is_DC7() {
					var _666___mcc_h1 _dafny.CodePoint = _source5.Get_().(D2_DC7).Cf6
					_ = _666___mcc_h1
					var _667___mcc_h2 _dafny.Map = _source5.Get_().(D2_DC7).Cf7
					_ = _667___mcc_h2
					var _668___mcc_h3 _dafny.Int = _source5.Get_().(D2_DC7).Cf8
					_ = _668___mcc_h3
					var _669_cf8 _dafny.Int = _668___mcc_h3
					_ = _669_cf8
					var _670_cf7 _dafny.Map = _667___mcc_h2
					_ = _670_cf7
					var _671_cf6 _dafny.CodePoint = _666___mcc_h1
					_ = _671_cf6
					return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv16, (_this).F6())).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_669_cf8, (_this).F6()))
				} else if _source5.Is_DC4() {
					var _672___mcc_h4 _dafny.Set = _source5.Get_().(D2_DC4).Cf4
					_ = _672___mcc_h4
					var _673_cf4 _dafny.Set = _672___mcc_h4
					_ = _673_cf4
					return (_pat_let_tv17).Cmp(_pat_let_tv18) <= 0
				} else {
					var _674___mcc_h5 D2 = _source5.Get_().(D2_DC8).Cf9
					_ = _674___mcc_h5
					var _675_cf9 D2 = _674___mcc_h5
					_ = _675_cf9
					return _this.F8()
				}
			}(Companion_D2_.Create_DC8_(_657_v3)) {
				{
					if (_658_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L8
					}
					_658_i0 = (_658_i0).Plus(_dafny.One)
					(_this).F7_set_(_this.F7())
					var _659_v4 _dafny.Array
					_ = _659_v4
					var _len0_18 _dafny.Int = _dafny.IntOfInt64(27)
					_ = _len0_18
					var _nw105 _dafny.Array
					_ = _nw105
					if _len0_18.Cmp(_dafny.Zero) == 0 {
						_nw105 = _dafny.NewArray(_len0_18)
					} else {
						var _init18 func(_dafny.Int) D10 = func(_660_i1 _dafny.Int) D10 {
							return Companion_D10_.Create_DC29_(_dafny.MultiSetOf(_dafny.SeqOf(_this.F8(), (_this).F6()), _dafny.SeqOf((_this).F6())))
						}
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
					_659_v4 = _nw105
					_659_v4 = _659_v4
					var _661_v5 _dafny.Sequence
					_ = _661_v5
					_661_v5 = _dafny.UnicodeSeqOfUtf8Bytes("mvhgcimfk")
					var _662_v6 D5
					_ = _662_v6
					_662_v6 = Companion_D5_.Create_DC15_(_661_v5)
					var _663_v7 D3
					_ = _663_v7
					_663_v7 = Companion_D3_.Create_DC10_((_this).F6(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_662_v6, true)).Cardinality())
					(_this).F8_set_((_663_v7).Dtor_cf11())
					(globalState).F0 = p0
					goto C8
				}
			C8:
			}
			goto L8
		}
	L8:
		var _676_v8 _dafny.MultiSet
		_ = _676_v8
		_676_v8 = _dafny.MultiSetOf(p0)
		var _677_v9 T0
		_ = _677_v9
		var _nw106 *C2 = New_C2_()
		_ = _nw106
		_nw106.Ctor__((_this).F5())
		_677_v9 = _nw106
		var _rhs129 _dafny.MultiSet = _676_v8
		_ = _rhs129
		var _rhs130 _dafny.Int = p0
		_ = _rhs130
		var _rhs131 bool = (p2) == (p2)
		_ = _rhs131
		var _rhs132 bool = p1
		_ = _rhs132
		var _rhs133 T0 = (func() T0 {
			if p1 {
				return (func() T0 {
					if (_this).F6() {
						return _677_v9
					}
					return _677_v9
				})()
			}
			return _677_v9
		})()
		_ = _rhs133
		var _lhs98 *GlobalState = globalState
		_ = _lhs98
		var _lhs99 *C4 = _this
		_ = _lhs99
		var _lhs100 *C4 = _this
		_ = _lhs100
		_676_v8 = _rhs129
		_lhs98.F0 = _rhs130
		_lhs99.F8_set_(_rhs131)
		_lhs100.F8_set_(_rhs132)
		_677_v9 = _rhs133
		var _678_v10 _dafny.Sequence
		_ = _678_v10
		_678_v10 = _dafny.SeqOf(p0, p0, p0, p0, p0)
		var _679_v11 _dafny.Sequence
		_ = _679_v11
		_679_v11 = _dafny.UnicodeSeqOfUtf8Bytes("lpaexls")
		var _680_v12 _dafny.Map
		_ = _680_v12
		_680_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_678_v10, _678_v10), _dafny.Companion_Sequence_.Contains(_679_v11, _654_v0))
		if (func() bool {
			if (_680_v12).Contains(_678_v10) {
				return (_680_v12).Get(_678_v10).(bool)
			}
			return _this.F8()
		})() {
			(globalState).F0 = p0
			var _681_v13 D2
			_ = _681_v13
			_681_v13 = Companion_D2_.Create_DC8_(_657_v3)
			var _source6 D2 = _681_v13
			_ = _source6
			if _source6.Is_DC5() {
				var _682___mcc_h6 bool = _source6.Get_().(D2_DC5).Cf5
				_ = _682___mcc_h6
				var _683_cf5 bool = _682___mcc_h6
				_ = _683_cf5
				_654_v0 = _654_v0
				var _684_v14 _dafny.Array
				_ = _684_v14
				var _nw107 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(17))
				_ = _nw107
				_684_v14 = _nw107
				_684_v14 = _684_v14
				_683_cf5 = p2
				var _685_v15 D7
				_ = _685_v15
				_685_v15 = Companion_D7_.Create_DC20_(_683_cf5)
				(_this).F8_set_((_685_v15).Dtor_cf29())
			} else if _source6.Is_DC6() {
				(_this).F8_set_(p1)
				var _686_v17 D0
				_ = _686_v17
				_686_v17 = Companion_D0_.Create_DC1_()
				var _687_v18 D1
				_ = _687_v18
				_687_v18 = Companion_D1_.Create_DC3_(_686_v17, p0)
				var _rhs134 bool = !(!(func() _dafny.Map {
					var _coll35 = _dafny.NewMapBuilder()
					_ = _coll35
					for _iter38 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(687), _dafny.IntOfInt64(872))); ; {
						_compr_35, _ok38 := _iter38()
						if !_ok38 {
							break
						}
						var _688_v16 _dafny.Int
						_688_v16 = interface{}(_compr_35).(_dafny.Int)
						if ((_dafny.IntOfInt64(687)).Cmp(_688_v16) <= 0) && ((_688_v16).Cmp(_dafny.IntOfInt64(872)) < 0) {
							_coll35.Add(Companion_Default___.SafeDivisionInt(_688_v16, p0), (_this).F6())
						}
					}
					return _coll35.ToMap()
				}()).Contains((p0).Plus(p0)))
				_ = _rhs134
				var _rhs135 _dafny.Int = (_687_v18).Dtor_cf3()
				_ = _rhs135
				var _rhs136 bool = _this.F8()
				_ = _rhs136
				var _lhs101 *C4 = _this
				_ = _lhs101
				var _lhs102 *GlobalState = globalState
				_ = _lhs102
				var _lhs103 *C4 = _this
				_ = _lhs103
				_lhs101.F8_set_(_rhs134)
				_lhs102.F0 = _rhs135
				_lhs103.F8_set_(_rhs136)
				var _689_v19 _dafny.Map
				_ = _689_v19
				_689_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _this.F8())
				_689_v19 = (_689_v19).Update((_this).F6(), (_this).F6())
				(globalState).F0 = (_this).Fm6(globalState)
			} else if _source6.Is_DC7() {
				var _690___mcc_h7 _dafny.CodePoint = _source6.Get_().(D2_DC7).Cf6
				_ = _690___mcc_h7
				var _691___mcc_h8 _dafny.Map = _source6.Get_().(D2_DC7).Cf7
				_ = _691___mcc_h8
				var _692___mcc_h9 _dafny.Int = _source6.Get_().(D2_DC7).Cf8
				_ = _692___mcc_h9
				var _693_cf8 _dafny.Int = _692___mcc_h9
				_ = _693_cf8
				var _694_cf7 _dafny.Map = _691___mcc_h8
				_ = _694_cf7
				var _695_cf6 _dafny.CodePoint = _690___mcc_h7
				_ = _695_cf6
				(globalState).F0 = _693_cf8
				var _696_v20 _dafny.Sequence
				_ = _696_v20
				_696_v20 = _dafny.SeqOf(true, p1)
				var _697_v21 _dafny.Map
				_ = _697_v21
				_697_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC2_(!((_this).F6())), _696_v20)
				var _698_v22 D1
				_ = _698_v22
				_698_v22 = Companion_D1_.Create_DC2_((_this).F6())
				_697_v21 = (_697_v21).Update(_698_v22, Companion_Default___.Fm23(p0, globalState))
				var _arr16 _dafny.Array = _this.F7()
				_ = _arr16
				var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index120
				(_arr16).ArraySet1(p1, (_index120).Int())
				var _699_v23 _dafny.MultiSet
				_ = _699_v23
				_699_v23 = _dafny.MultiSetOf(_678_v10)
				var _700_v24 D0
				_ = _700_v24
				_700_v24 = Companion_D0_.Create_DC1_()
				var _701_v25 D1
				_ = _701_v25
				_701_v25 = Companion_D1_.Create_DC3_(_700_v24, p0)
				var _702_v26 _dafny.Map
				_ = _702_v26
				_702_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F8())
				var _arr17 _dafny.Array = _this.F7()
				_ = _arr17
				var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index121
				var _rhs137 bool = (_677_v9).Fm1((_699_v23).Equals(_699_v23), (p0).Minus(_693_cf8), (_701_v25).Dtor_cf3(), globalState)
				_ = _rhs137
				var _rhs138 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_678_v10, _678_v10)
				_ = _rhs138
				var _rhs139 bool = (_this).Fm2((_702_v26).Equals(_702_v26), globalState)
				_ = _rhs139
				var _rhs140 _dafny.Sequence = _678_v10
				_ = _rhs140
				var _lhs104 _dafny.Array = _this.F7()
				_ = _lhs104
				var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(682), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs105
				var _lhs106 *C4 = _this
				_ = _lhs106
				(_lhs104).ArraySet1(_rhs137, (_lhs105).Int())
				_678_v10 = _rhs138
				_lhs106.F8_set_(_rhs139)
				_678_v10 = _rhs140
				(globalState).F0 = (_dafny.Zero).Minus(((_this).Fm6(globalState)).Plus(_693_cf8))
			} else if _source6.Is_DC4() {
				var _703___mcc_h10 _dafny.Set = _source6.Get_().(D2_DC4).Cf4
				_ = _703___mcc_h10
				var _704_cf4 _dafny.Set = _703___mcc_h10
				_ = _704_cf4
				(_this).F8_set_(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("vurhyya"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(844))).Uint32(), func(coer36 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}((func(_705_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_706_i2 _dafny.Int) _dafny.CodePoint {
						return _705_v0
					}
				})(_654_v0)))), _654_v0))
				var _707_v27 _dafny.Array
				_ = _707_v27
				var _nw108 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(12))
				_ = _nw108
				_707_v27 = _nw108
				(_this).F8_set_(_dafny.Companion_Sequence_.Contains(_679_v11, (_679_v11).Select((Companion_Default___.SafeIndex((_dafny.MultiSetOf(_707_v27)).Cardinality(), _dafny.IntOfUint32((_679_v11).Cardinality()))).Uint32()).(_dafny.CodePoint)))
				(_this).F8_set_(!(p2))
				var _708_v28 _dafny.Map
				_ = _708_v28
				_708_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p1) && (_this.F8()), (_this).Fm6(globalState))
				var _709_v29 _dafny.Sequence
				_ = _709_v29
				_709_v29 = _dafny.SeqOf((_this).F6())
				_708_v28 = (_708_v28).Update((_this).F6(), (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_709_v29, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_709_v29).Cardinality()))).Uint32(), (_this).F6())).Cardinality())).Plus(p0))
			} else {
				var _710___mcc_h11 D2 = _source6.Get_().(D2_DC8).Cf9
				_ = _710___mcc_h11
				var _711_cf9 D2 = _710___mcc_h11
				_ = _711_cf9
				var _712_v30 _dafny.Sequence
				_ = _712_v30
				_712_v30 = _dafny.SeqOf(_this.F8())
				_712_v30 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_712_v30, _712_v30), _dafny.SeqOf(false))
				(_this).F8_set_(_dafny.Companion_Sequence_.IsPrefixOf((_this).Fm7(globalState), _679_v11))
				var _713_v31 _dafny.Array
				_ = _713_v31
				var _nw109 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
				_ = _nw109
				_713_v31 = _nw109
				var _714_v32 _dafny.MultiSet
				_ = _714_v32
				_714_v32 = _dafny.MultiSetOf(_713_v31)
				(_this).F8_set_((_714_v32).Equals(_714_v32))
				(_this).F7_set_(_this.F7())
			}
			var _715_v33 _dafny.Map
			_ = _715_v33
			_715_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F8())
			_715_v33 = _715_v33
			if p2 {
				var _len0_19 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_19
				var _nw110 _dafny.Array
				_ = _nw110
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw110 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) bool = func(_716_i3 _dafny.Int) bool {
						return true
					}
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw110 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw110).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw110).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				(_this).F7_set_(_nw110)
				var _717_v34 D0
				_ = _717_v34
				_717_v34 = Companion_D0_.Create_DC0_(p0)
				var _718_v35 _dafny.Sequence
				_ = _718_v35
				_718_v35 = _dafny.SeqOf(Companion_D0_.Create_DC0_(p0), _717_v34, _717_v34, (func() D0 {
					if p2 {
						return _717_v34
					}
					return _717_v34
				})(), Companion_Default___.Fm25((_this).F6(), p0, p2, globalState))
				(globalState).F0 = _dafny.IntOfUint32((_718_v35).Cardinality())
				var _719_v36 _dafny.Map
				_ = _719_v36
				_719_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), p2)
				var _720_v37 D9
				_ = _720_v37
				_720_v37 = Companion_D9_.Create_DC27_(_719_v36, (func() bool {
					if p2 {
						return _this.F8()
					}
					return p1
				})())
				_720_v37 = _720_v37
				(_this).F8_set_((_this).F6())
				var _721_v38 _dafny.Array
				_ = _721_v38
				var _nw111 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw111
				_721_v38 = _nw111
				var _722_v40 _dafny.Sequence
				_ = _722_v40
				_722_v40 = _dafny.SeqOf((_this).F5())
				var _723_v41 _dafny.Map
				_ = _723_v41
				_723_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), (_this).F6())
				var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_721_v38), 0))
				_ = _index122
				(_721_v38).ArraySet1((func() _dafny.Map {
					var _coll36 = _dafny.NewMapBuilder()
					_ = _coll36
					for _iter39 := _dafny.Iterate((_722_v40).Elements()); ; {
						_compr_36, _ok39 := _iter39()
						if !_ok39 {
							break
						}
						var _724_v39 _dafny.MultiSet
						_724_v39 = interface{}(_compr_36).(_dafny.MultiSet)
						if _dafny.Companion_Sequence_.Contains(_722_v40, _724_v39) {
							_coll36.Add(_724_v39, _this.F8())
						}
					}
					return _coll36.ToMap()
				}()).Merge(_723_v41), (_index122).Int())
				var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_721_v38), 0))
				_ = _index123
				(_721_v38).ArraySet1((_723_v41).Merge(_723_v41), (_index123).Int())
			} else {
				_654_v0 = _654_v0
				var _725_v42 *C2
				_ = _725_v42
				var _nw112 *C2 = New_C2_()
				_ = _nw112
				_nw112.Ctor__((_677_v9).F5())
				_725_v42 = _nw112
				var _726_v43 D2
				_ = _726_v43
				_726_v43 = Companion_D2_.Create_DC7_(_654_v0, _656_v2, p0)
				(globalState).F0 = (_726_v43).Dtor_cf8()
				(_this).F7_set_(_this.F7())
				(globalState).F0 = (_dafny.Zero).Minus(p0)
			}
			var _727_v44 _dafny.Map
			_ = _727_v44
			_727_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F6())
			var _728_v45 *C1
			_ = _728_v45
			var _nw113 *C1 = New_C1_()
			_ = _nw113
			_nw113.Ctor__(p2, _dafny.Companion_Sequence_.Concatenate(_679_v11, _679_v11), (func() bool {
				if (_727_v44).Contains(p0) {
					return (_727_v44).Get(p0).(bool)
				}
				return false
			})(), (_this).F5())
			_728_v45 = _nw113
		} else {
			var _729_v46 _dafny.Array
			_ = _729_v46
			var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
			_ = _nw114
			_729_v46 = _nw114
			var _730_v47 _dafny.Map
			_ = _730_v47
			_730_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_this).F6())
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_729_v46), 0))
			_ = _index124
			(_729_v46).ArraySet1((_730_v47).Merge(_730_v47), (_index124).Int())
			var _731_v48 D11
			_ = _731_v48
			_731_v48 = Companion_D11_.Create_DC32_(_730_v47)
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_729_v46), 0))
			_ = _index125
			(_729_v46).ArraySet1((_730_v47).Merge(((_731_v48).Dtor_cf51()).Merge(_730_v47)), (_index125).Int())
			var _732_v49 D1
			_ = _732_v49
			_732_v49 = Companion_D1_.Create_DC2_(p1)
			var _source7 D1 = _732_v49
			_ = _source7
			if _source7.Is_DC3() {
				var _733___mcc_h12 D0 = _source7.Get_().(D1_DC3).Cf2
				_ = _733___mcc_h12
				var _734___mcc_h13 _dafny.Int = _source7.Get_().(D1_DC3).Cf3
				_ = _734___mcc_h13
				var _735_cf3 _dafny.Int = _734___mcc_h13
				_ = _735_cf3
				var _736_cf2 D0 = _733___mcc_h12
				_ = _736_cf2
				_679_v11 = _679_v11
				(globalState).F0 = (_735_cf3).Minus(p0)
				var _737_v50 _dafny.Sequence
				_ = _737_v50
				_737_v50 = _dafny.SeqOf(p2)
				var _738_v51 *C1
				_ = _738_v51
				var _nw115 *C1 = New_C1_()
				_ = _nw115
				_nw115.Ctor__(p2, _679_v11, _dafny.Companion_Sequence_.IsProperPrefixOf(Companion_Default___.Fm23(_735_cf3, globalState), _737_v50), (_677_v9).F5())
				_738_v51 = _nw115
				var _739_v52 D3
				_ = _739_v52
				_739_v52 = Companion_D3_.Create_DC11_(_dafny.IntOfInt64(58))
				var _740_v53 _dafny.Map
				_ = _740_v53
				_740_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _739_v52)
				_735_cf3 = (_735_cf3).Minus(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), (_740_v53).Cardinality()))
			} else {
				var _741___mcc_h14 bool = _source7.Get_().(D1_DC2).Cf1
				_ = _741___mcc_h14
				var _742_cf1 bool = _741___mcc_h14
				_ = _742_cf1
				var _743_v54 _dafny.Set
				_ = _743_v54
				_743_v54 = _dafny.SetOf(p2)
				_743_v54 = _743_v54
				var _744_v55 _dafny.Array
				_ = _744_v55
				var _nw116 _dafny.Array = _dafny.NewArrayWithValue(Companion_D2_.Default(), _dafny.IntOfInt64(28))
				_ = _nw116
				_744_v55 = _nw116
				_744_v55 = _744_v55
				(globalState).F0 = (Companion_Default___.SafeDivisionInt(p0, p0)).Plus(p0)
				var _745_v56 _dafny.Array
				_ = _745_v56
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(9))
				_ = _nw117
				_745_v56 = _nw117
				var _746_v57 _dafny.Map
				_ = _746_v57
				_746_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v0, false)
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_745_v56), 0))
				_ = _index126
				(_745_v56).ArraySet1(_746_v57, (_index126).Int())
				var _747_v58 D12
				_ = _747_v58
				_747_v58 = Companion_D12_.Create_DC34_(_746_v57)
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(290), _dafny.ArrayLen((_745_v56), 0))
				_ = _index127
				(_745_v56).ArraySet1(((_747_v58).Dtor_cf54()).Merge(_746_v57), (_index127).Int())
			}
			var _748_v59 _dafny.Map
			_ = _748_v59
			_748_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _dafny.MultiSetOf(p2))
			var _749_v60 _dafny.Array
			_ = _749_v60
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_20
			var _nw118 _dafny.Array
			_ = _nw118
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw118 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) _dafny.Int = func(_750_i4 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeModuloInt(_750_i4, _dafny.IntOfInt64(-705))
				}
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw118 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw118).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw118).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_749_v60 = _nw118
			var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_749_v60), 0))
			_ = _index128
			(_749_v60).ArraySet1(p0, (_index128).Int())
			var _751_v61 D13
			_ = _751_v61
			_751_v61 = Companion_D13_.Create_DC38_(_748_v59)
			var _752_v62 _dafny.Set
			_ = _752_v62
			_752_v62 = _dafny.SetOf((_this).F6(), p2)
			var _753_v63 _dafny.Sequence
			_ = _753_v63
			_753_v63 = _dafny.SeqOf(_this.F8())
			var _754_v64 _dafny.Set
			_ = _754_v64
			_754_v64 = _dafny.SetOf(_dafny.IntOfUint32((_753_v63).Cardinality()))
			var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_749_v60), 0))
			_ = _index129
			var _rhs141 _dafny.Map = (_751_v61).Dtor_cf59()
			_ = _rhs141
			var _rhs142 _dafny.Int = ((_752_v62).Intersection((_752_v62).Intersection(_dafny.SetOf((_753_v63).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_753_v63).Cardinality()))).Uint32()).(bool))))).Cardinality()
			_ = _rhs142
			var _rhs143 bool = !(p1) || (p2)
			_ = _rhs143
			var _rhs144 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (p0).Plus((_754_v64).Cardinality()))
			_ = _rhs144
			var _lhs107 _dafny.Array = _749_v60
			_ = _lhs107
			var _lhs108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_749_v60), 0))
			_ = _lhs108
			var _lhs109 *C4 = _this
			_ = _lhs109
			var _lhs110 *GlobalState = globalState
			_ = _lhs110
			_748_v59 = _rhs141
			(_lhs107).ArraySet1(_rhs142, (_lhs108).Int())
			_lhs109.F8_set_(_rhs143)
			_lhs110.F0 = _rhs144
			(globalState).F0 = ((_749_v60).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(48), _dafny.ArrayLen((_749_v60), 0))).Int()).(_dafny.Int)).Plus(p0)
			var _755_v65 *C2
			_ = _755_v65
			var _nw119 *C2 = New_C2_()
			_ = _nw119
			_nw119.Ctor__((_677_v9).F5())
			_755_v65 = _nw119
		}
		if _this.F8() {
			(_this).F8_set_(p2)
			var _756_v66 D5
			_ = _756_v66
			_756_v66 = Companion_D5_.Create_DC15_(_dafny.UnicodeSeqOfUtf8Bytes("uipsyskp"))
			var _pat_let_tv19 = _679_v11
			_ = _pat_let_tv19
			var _757_v67 _dafny.Map
			_ = _757_v67
			_757_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_678_v10, func(_pat_let24_0 D5) D5 {
				return func(_758_dt__update__tmp_h0 D5) D5 {
					return func(_pat_let25_0 _dafny.Sequence) D5 {
						return func(_759_dt__update_hcf18_h0 _dafny.Sequence) D5 {
							return Companion_D5_.Create_DC15_(_759_dt__update_hcf18_h0)
						}(_pat_let25_0)
					}(_pat_let_tv19)
				}(_pat_let24_0)
			}(_756_v66))
			_757_v67 = _757_v67
			var _760_v68 _dafny.Sequence
			_ = _760_v68
			_760_v68 = _dafny.SeqOf(!(p2))
			_656_v2 = (_656_v2).Update(p0, (func() _dafny.Int {
				if (_676_v8).Contains(p0) {
					return (_676_v8).Multiplicity(p0)
				}
				return _dafny.IntOfUint32((_760_v68).Cardinality())
			})())
			var _761_v69 D12
			_ = _761_v69
			_761_v69 = Companion_D12_.Create_DC36_((func() _dafny.Int {
				if _this.F8() {
					return (func() _dafny.Int {
						if ((_this).F5()).Contains(_this.F8()) {
							return ((_this).F5()).Multiplicity(_this.F8())
						}
						return p0
					})()
				}
				return p0
			})(), p0)
			_761_v69 = _761_v69
			(_this).F8_set_(true)
		} else {
			var _762_v70 _dafny.Map
			_ = _762_v70
			_762_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _this.F7())
			_762_v70 = (_762_v70).Merge(_762_v70)
			_654_v0 = _654_v0
			var _763_v71 D3
			_ = _763_v71
			_763_v71 = Companion_D3_.Create_DC11_(p0)
			var _764_v72 _dafny.Map
			_ = _764_v72
			_764_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_763_v71).Dtor_cf13()), _679_v11)
			var _765_v73 _dafny.Sequence
			_ = _765_v73
			_765_v73 = _dafny.SeqOf(p2)
			_764_v72 = (_764_v72).Update(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_765_v73).Cardinality()), p0), _679_v11)
			var _766_v74 _dafny.Array
			_ = _766_v74
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_21
			var _nw120 _dafny.Array
			_ = _nw120
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw120 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) _dafny.Int = (func(_767_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_768_i5 _dafny.Int) _dafny.Int {
						return (_768_i5).Times(_767_p0)
					}
				})(p0)
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw120 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw120).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw120).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_766_v74 = _nw120
			var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_766_v74), 0))
			_ = _index130
			(_766_v74).ArraySet1(p0, (_index130).Int())
			var _769_v75 _dafny.Sequence
			_ = _769_v75
			_769_v75 = _dafny.SeqOf(_765_v73, _765_v73)
			var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_766_v74), 0))
			_ = _index131
			var _rhs145 _dafny.Int = (_678_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_769_v75, _769_v75)).Cardinality()), _dafny.IntOfUint32((_678_v10).Cardinality()))).Uint32()).(_dafny.Int)
			_ = _rhs145
			var _rhs146 _dafny.Int = (Companion_Default___.SafeDivisionInt(p0, p0)).Minus(_dafny.IntOfUint32((_765_v73).Cardinality()))
			_ = _rhs146
			var _rhs147 bool = p1
			_ = _rhs147
			var _rhs148 _dafny.Sequence = _679_v11
			_ = _rhs148
			var _rhs149 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_679_v11, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_679_v11).Cardinality()))).Uint32(), _654_v0), _dafny.Companion_Sequence_.Concatenate(_679_v11, _dafny.UnicodeSeqOfUtf8Bytes("uayxfpbbt")))
			_ = _rhs149
			var _lhs111 _dafny.Array = _766_v74
			_ = _lhs111
			var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(558), _dafny.ArrayLen((_766_v74), 0))
			_ = _lhs112
			var _lhs113 *GlobalState = globalState
			_ = _lhs113
			var _lhs114 *C4 = _this
			_ = _lhs114
			(_lhs111).ArraySet1(_rhs145, (_lhs112).Int())
			_lhs113.F0 = _rhs146
			_lhs114.F8_set_(_rhs147)
			_679_v11 = _rhs148
			_679_v11 = _rhs149
			var _770_v76 *C3
			_ = _770_v76
			var _nw121 *C3 = New_C3_()
			_ = _nw121
			_nw121.Ctor__(_this.F8(), (_677_v9).F5())
			_770_v76 = _nw121
		}
		var _771_v77 D12
		_ = _771_v77
		_771_v77 = Companion_D12_.Create_DC36_(p0, p0)
		var _source8 D12 = _771_v77
		_ = _source8
		if _source8.Is_DC35() {
			var _772___mcc_h15 bool = _source8.Get_().(D12_DC35).Cf55
			_ = _772___mcc_h15
			var _773_cf55 bool = _772___mcc_h15
			_ = _773_cf55
			_654_v0 = _654_v0
			var _774_v78 _dafny.Array
			_ = _774_v78
			var _nw122 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
			_ = _nw122
			_774_v78 = _nw122
			var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))
			_ = _index132
			(_774_v78).ArraySet1(_679_v11, (_index132).Int())
			var _775_v79 _dafny.Array
			_ = _775_v79
			var _nw123 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
			_ = _nw123
			_775_v79 = _nw123
			var _776_v80 _dafny.Sequence
			_ = _776_v80
			_776_v80 = _dafny.SeqOf(_775_v79)
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))
			_ = _index133
			(_774_v78).ArraySet1(_dafny.Companion_Sequence_.Update(_679_v11, (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_776_v80).Cardinality()), p0), _dafny.IntOfUint32((_679_v11).Cardinality()))).Uint32(), _654_v0), (_index133).Int())
			var _777_v81 _dafny.Map
			_ = _777_v81
			_777_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(547), p2)
			var _778_v82 _dafny.Map
			_ = _778_v82
			_778_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_777_v81).Cardinality(), _dafny.IntOfUint32((_678_v10).Cardinality())), _678_v10), p0)
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))
			_ = _index134
			(_775_v79).ArraySet1((p0).Minus(p0), (_index134).Int())
			var _arr18 _dafny.Array = _this.F7()
			_ = _arr18
			var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index135
			(_arr18).ArraySet1(!(_773_cf55) || (!(p1)), (_index135).Int())
			var _779_v83 _dafny.Sequence
			_ = _779_v83
			_779_v83 = _dafny.SeqOf(p1)
			var _780_v84 _dafny.Set
			_ = _780_v84
			_780_v84 = _dafny.SetOf(p0, p0)
			var _781_v86 _dafny.Sequence
			_ = _781_v86
			_781_v86 = _dafny.SeqOf(_656_v2, _656_v2, _656_v2, _656_v2, _656_v2)
			var _782_v87 *C0
			_ = _782_v87
			var _nw124 *C0 = New_C0_()
			_ = _nw124
			_nw124.Ctor__(func() _dafny.Map {
				var _coll37 = _dafny.NewMapBuilder()
				_ = _coll37
				for _iter40 := _dafny.Iterate((_781_v86).Elements()); ; {
					_compr_37, _ok40 := _iter40()
					if !_ok40 {
						break
					}
					var _783_v85 _dafny.Map
					_783_v85 = interface{}(_compr_37).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_781_v86, _783_v85) {
						_coll37.Add(_783_v85, _this.F8())
					}
				}
				return _coll37.ToMap()
			}())
			_782_v87 = _nw124
			var _784_v88 _dafny.Map
			_ = _784_v88
			_784_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(((_677_v9).F5()).Cardinality()), _782_v87)
			var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))
			_ = _index136
			var _arr19 _dafny.Array = _this.F7()
			_ = _arr19
			var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index137
			var _rhs150 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_678_v10, _dafny.IntOfUint32((_779_v83).Cardinality()))
			_ = _rhs150
			var _rhs151 bool = (_677_v9).Fm1(p2, p0, p0, globalState)
			_ = _rhs151
			var _rhs152 _dafny.Int = (p0).Times(((_678_v10).Select((Companion_Default___.SafeIndex((_780_v84).Cardinality(), _dafny.IntOfUint32((_678_v10).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_784_v88).Cardinality()))
			_ = _rhs152
			var _rhs153 bool = !((_this).F6())
			_ = _rhs153
			var _rhs154 bool = (p2) && (_dafny.Companion_Sequence_.IsPrefixOf(_679_v11, _dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("p"), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("p")).Cardinality()))).Uint32(), Companion_Default___.Fm16(p2, _656_v2, (_774_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))).Int()).(_dafny.Sequence), globalState))))
			_ = _rhs154
			var _lhs115 *C4 = _this
			_ = _lhs115
			var _lhs116 _dafny.Array = _775_v79
			_ = _lhs116
			var _lhs117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))
			_ = _lhs117
			var _lhs118 _dafny.Array = _this.F7()
			_ = _lhs118
			var _lhs119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
			_ = _lhs119
			_778_v82 = _rhs150
			_lhs115.F8_set_(_rhs151)
			(_lhs116).ArraySet1(_rhs152, (_lhs117).Int())
			(_lhs118).ArraySet1(_rhs153, (_lhs119).Int())
			_773_cf55 = _rhs154
			if p1 {
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))
				_ = _index138
				(_774_v78).ArraySet1(_679_v11, (_index138).Int())
				(_this).F8_set_(_this.F8())
				var _785_v89 *C1
				_ = _785_v89
				var _nw125 *C1 = New_C1_()
				_ = _nw125
				_nw125.Ctor__(!(!(p2)) || (_this.F8()), _679_v11, p1, (_677_v9).F5())
				_785_v89 = _nw125
				var _786_v90 _dafny.MultiSet
				_ = _786_v90
				_786_v90 = _dafny.MultiSetOf(Companion_Default___.Fm18(globalState))
				var _787_v91 _dafny.Array
				_ = _787_v91
				var _nw126 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(14))
				_ = _nw126
				_787_v91 = _nw126
				var _788_v92 _dafny.Map
				_ = _788_v92
				_788_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
					if _773_cf55 {
						return _787_v91
					}
					return _787_v91
				})(), p0)
				var _arr20 _dafny.Array = _this.F7()
				_ = _arr20
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index139
				var _rhs155 _dafny.Int = (Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(p0))).Times((_775_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))).Int()).(_dafny.Int))
				_ = _rhs155
				var _rhs156 bool = !(_780_v84).Contains((func() _dafny.Int {
					if (_786_v90).Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(810))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg37 _dafny.Int) interface{} {
							return coer37(arg37)
						}
					}((func(_789_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_790_i6 _dafny.Int) _dafny.Int {
							return _789_p0
						}
					})(p0)))) {
						return (_786_v90).Multiplicity(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(810))).Uint32(), func(coer38 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
							return func(arg38 _dafny.Int) interface{} {
								return coer38(arg38)
							}
						}((func(_791_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
							return func(_792_i6 _dafny.Int) _dafny.Int {
								return _791_p0
							}
						})(p0))))
					}
					return p0
				})())
				_ = _rhs156
				var _rhs157 _dafny.Int = (func() _dafny.Int {
					if (_788_v92).Contains(_787_v91) {
						return (_788_v92).Get(_787_v91).(_dafny.Int)
					}
					return ((_775_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))).Int()).(_dafny.Int)).Minus((_775_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))).Int()).(_dafny.Int))
				})()
				_ = _rhs157
				var _lhs120 *GlobalState = globalState
				_ = _lhs120
				var _lhs121 _dafny.Array = _this.F7()
				_ = _lhs121
				var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs122
				var _lhs123 *GlobalState = globalState
				_ = _lhs123
				_lhs120.F0 = _rhs155
				(_lhs121).ArraySet1(_rhs156, (_lhs122).Int())
				_lhs123.F0 = _rhs157
				var _793_v93 _dafny.Map
				_ = _793_v93
				_793_v93 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("yc"), (_785_v89).F12())
				var _794_v94 _dafny.Sequence
				_ = _794_v94
				_794_v94 = _dafny.SeqOf((_785_v89).F13(), (_774_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))).Int()).(_dafny.Sequence), (_774_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))).Int()).(_dafny.Sequence))
				var _795_v95 _dafny.Map
				_ = _795_v95
				_795_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_794_v94).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(280))).Uint32(), func(coer39 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_796_v79 _dafny.Array) func(_dafny.Int) _dafny.Int {
					return func(_797_i7 _dafny.Int) _dafny.Int {
						return (_796_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_796_v79), 0))).Int()).(_dafny.Int)
					}
				})(_775_v79)))).Cardinality()), _dafny.IntOfUint32((_794_v94).Cardinality()))).Uint32()).(_dafny.Sequence), _654_v0)
				(globalState).F0 = (_785_v89).Fm4(Companion_Default___.SafeDivisionInt((_775_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))).Int()).(_dafny.Int), (_775_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))).Int()).(_dafny.Int)), (_775_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(386), _dafny.ArrayLen((_775_v79), 0))).Int()).(_dafny.Int), _793_v93, (_795_v95).Cardinality(), globalState)
			} else {
				var _798_v96 *C0
				_ = _798_v96
				var _nw127 *C0 = New_C0_()
				_ = _nw127
				_nw127.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_656_v2, (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)))
				_798_v96 = _nw127
				var _799_v97 _dafny.Map
				_ = _799_v97
				_799_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_773_cf55, true)
				var _arr21 _dafny.Array = _this.F7()
				_ = _arr21
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index140
				var _rhs158 bool = p1
				_ = _rhs158
				var _rhs159 _dafny.Int = (((_799_v97).Merge(_799_v97)).Cardinality()).Times((_dafny.Zero).Minus(p0))
				_ = _rhs159
				var _rhs160 _dafny.Map = _799_v97
				_ = _rhs160
				var _lhs124 _dafny.Array = _this.F7()
				_ = _lhs124
				var _lhs125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(748), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs125
				var _lhs126 *GlobalState = globalState
				_ = _lhs126
				(_lhs124).ArraySet1(_rhs158, (_lhs125).Int())
				_lhs126.F0 = _rhs159
				_799_v97 = _rhs160
				var _800_v98 _dafny.Map
				_ = _800_v98
				_800_v98 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_654_v0, _773_cf55)
				var _801_v99 _dafny.MultiSet
				_ = _801_v99
				_801_v99 = _dafny.MultiSetOf(_800_v98)
				var _802_v100 *C1
				_ = _802_v100
				var _nw128 *C1 = New_C1_()
				_ = _nw128
				_nw128.Ctor__((_801_v99).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(564))).Uint32(), func(coer40 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_803_v98 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_804_i8 _dafny.Int) _dafny.Map {
						return _803_v98
					}
				})(_800_v98))))), (_774_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))).Int()).(_dafny.Sequence), _dafny.Companion_Sequence_.IsPrefixOf(_679_v11, (_774_v78).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(443), _dafny.ArrayLen((_774_v78), 0))).Int()).(_dafny.Sequence)), (_677_v9).F5())
				_802_v100 = _nw128
				_773_cf55 = p2
				_775_v79 = _775_v79
			}
		} else if _source8.Is_DC36() {
			var _805___mcc_h16 _dafny.Int = _source8.Get_().(D12_DC36).Cf56
			_ = _805___mcc_h16
			var _806___mcc_h17 _dafny.Int = _source8.Get_().(D12_DC36).Cf57
			_ = _806___mcc_h17
			var _807_cf57 _dafny.Int = _806___mcc_h17
			_ = _807_cf57
			var _808_cf56 _dafny.Int = _805___mcc_h16
			_ = _808_cf56
			(_this).F8_set_(!((_807_cf57).Cmp(_807_cf57) > 0) || ((func() bool {
				if p1 {
					return _this.F8()
				}
				return p1
			})()))
			_678_v10 = _678_v10
			_807_cf57 = ((_678_v10).Select((Companion_Default___.SafeIndex(_808_cf56, _dafny.IntOfUint32((_678_v10).Cardinality()))).Uint32()).(_dafny.Int)).Minus((_676_v8).Cardinality())
			(globalState).F0 = _808_cf56
		} else if _source8.Is_DC37() {
			var _809___mcc_h18 _dafny.CodePoint = _source8.Get_().(D12_DC37).Cf58
			_ = _809___mcc_h18
			var _810_cf58 _dafny.CodePoint = _809___mcc_h18
			_ = _810_cf58
			var _811_v101 *C1
			_ = _811_v101
			var _nw129 *C1 = New_C1_()
			_ = _nw129
			_nw129.Ctor__(p2, _679_v11, (_this).F6(), _dafny.MultiSetOf(p1))
			_811_v101 = _nw129
			var _812_v102 _dafny.Set
			_ = _812_v102
			_812_v102 = _dafny.SetOf(_811_v101)
			if (_812_v102).IsProperSubsetOf(_dafny.SetOf(_811_v101, _811_v101)) {
				(_this).F8_set_(_this.F8())
				var _813_v103 *C0
				_ = _813_v103
				var _nw130 *C0 = New_C0_()
				_ = _nw130
				_nw130.Ctor__(Companion_Default___.Fm26(p0, globalState))
				_813_v103 = _nw130
				(globalState).F0 = ((Companion_D6_.Create_DC18_(_813_v103, Companion_Default___.Fm23((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_813_v103, _dafny.IntOfInt64(179))).Cardinality(), globalState), _dafny.IntOfUint32((_679_v11).Cardinality()))).Dtor_cf27()).Minus(p0)
				var _814_v104 *C1
				_ = _814_v104
				var _nw131 *C1 = New_C1_()
				_ = _nw131
				_nw131.Ctor__(((_677_v9).F5()).IsDisjointFrom((_this).F5()), _dafny.UnicodeSeqOfUtf8Bytes("j"), _this.F8(), (_this).F5())
				_814_v104 = _nw131
				_656_v2 = (_656_v2).Update(p0, Companion_Default___.SafeModuloInt(((_677_v9).F5()).Cardinality(), p0))
				var _815_v105 *C3
				_ = _815_v105
				var _nw132 *C3 = New_C3_()
				_ = _nw132
				_nw132.Ctor__(p1, _dafny.MultiSetOf((_811_v101).F12()))
				_815_v105 = _nw132
			} else {
				(_this).F8_set_((p0).Cmp((_dafny.Zero).Minus(p0)) < 0)
				(_this).F8_set_((p0).Cmp(Companion_Default___.SafeModuloInt(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(509), (_811_v101).F12())).Cardinality())) == 0)
				var _816_v106 _dafny.Map
				_ = _816_v106
				_816_v106 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-876))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg41 _dafny.Int) interface{} {
						return coer41(arg41)
					}
				}(func(_817_i9 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('k')
				})), (_811_v101).F12())
				var _818_v107 _dafny.Map
				_ = _818_v107
				_818_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_811_v101).Fm4(p0, _dafny.IntOfInt64(-167), _816_v106, (_811_v101).Fm4((_dafny.Zero).Minus(p0), p0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_679_v11, p1), p0, globalState), globalState), _810_cf58)
				_818_v107 = (_818_v107).Update((_dafny.Zero).Minus(p0), Companion_Default___.Fm16(_this.F8(), _656_v2, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-403))).Uint32(), func(coer42 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg42 _dafny.Int) interface{} {
						return coer42(arg42)
					}
				}((func(_819_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_820_i10 _dafny.Int) _dafny.CodePoint {
						return _819_v0
					}
				})(_654_v0))), globalState))
				var _821_v108 _dafny.Sequence
				_ = _821_v108
				_821_v108 = _dafny.SeqOf((_811_v101).F12(), p1)
				var _822_v109 _dafny.Map
				_ = _822_v109
				_822_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(999), !(_this.F8()))
				var _823_v110 _dafny.Map
				_ = _823_v110
				_823_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(p1), (_811_v101).Fm3(globalState))
				var _824_v111 D8
				_ = _824_v111
				_824_v111 = Companion_D8_.Create_DC24_(p0, (_811_v101).F12(), _821_v108, (_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm8(p0, (func() bool {
					if (_822_v109).Contains((_823_v110).Cardinality()) {
						return (_822_v109).Get((_823_v110).Cardinality()).(bool)
					}
					return p2
				})(), _679_v11, globalState)).Cardinality())), p0)
				(globalState).F0 = _dafny.IntOfUint32(((_824_v111).Dtor_cf36()).Cardinality())
				var _825_v113 _dafny.Set
				_ = _825_v113
				_825_v113 = _dafny.SetOf(_dafny.Companion_Sequence_.Update((_811_v101).F13(), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_811_v101).F13()).Cardinality()))).Uint32(), _654_v0))
				var _826_v114 D15
				_ = _826_v114
				_826_v114 = Companion_D15_.Create_DC41_(_816_v106)
				(globalState).F0 = (_this).Fm4((_this).Fm4(p0, p0, func() _dafny.Map {
					var _coll38 = _dafny.NewMapBuilder()
					_ = _coll38
					for _iter41 := _dafny.Iterate((_825_v113).Elements()); ; {
						_compr_38, _ok41 := _iter41()
						if !_ok41 {
							break
						}
						var _827_v112 _dafny.Sequence
						_827_v112 = interface{}(_compr_38).(_dafny.Sequence)
						if (_825_v113).Contains(_827_v112) {
							_coll38.Add(_827_v112, true)
						}
					}
					return _coll38.ToMap()
				}(), (_dafny.Zero).Minus(p0), globalState), _dafny.IntOfInt64(184), (_826_v114).Dtor_cf66(), (p0).Minus((_818_v107).Cardinality()), globalState)
			}
			var _rhs161 bool = (_this).Fm3(globalState)
			_ = _rhs161
			var _rhs162 bool = p2
			_ = _rhs162
			var _lhs127 *C4 = _this
			_ = _lhs127
			var _lhs128 *C4 = _this
			_ = _lhs128
			_lhs127.F8_set_(_rhs161)
			_lhs128.F8_set_(_rhs162)
			_654_v0 = _810_cf58
			_676_v8 = _676_v8
		} else {
			var _828___mcc_h19 _dafny.Map = _source8.Get_().(D12_DC34).Cf54
			_ = _828___mcc_h19
			var _829_cf54 _dafny.Map = _828___mcc_h19
			_ = _829_cf54
			_678_v10 = _678_v10
			var _830_v115 *C2
			_ = _830_v115
			var _nw133 *C2 = New_C2_()
			_ = _nw133
			_nw133.Ctor__((_this).F5())
			_830_v115 = _nw133
			var _831_v116 D9
			_ = _831_v116
			_831_v116 = Companion_D9_.Create_DC26_(_830_v115)
			var _source9 D9 = _831_v116
			_ = _source9
			if _source9.Is_DC27() {
				var _832___mcc_h20 _dafny.Map = _source9.Get_().(D9_DC27).Cf42
				_ = _832___mcc_h20
				var _833___mcc_h21 bool = _source9.Get_().(D9_DC27).Cf43
				_ = _833___mcc_h21
				var _834_cf43 bool = _833___mcc_h21
				_ = _834_cf43
				var _835_cf42 _dafny.Map = _832___mcc_h20
				_ = _835_cf42
				_679_v11 = _679_v11
				var _836_v117 _dafny.Array
				_ = _836_v117
				var _nwElement0_18 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("sspyxy")
				_ = _nwElement0_18
				var _nw134 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(14))
				_ = _nw134
				(_nw134).ArraySet1(_nwElement0_18, 0)
				(_nw134).ArraySet1(_679_v11, 1)
				(_nw134).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("enh"), 2)
				(_nw134).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(128))).Uint32(), func(coer43 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg43 _dafny.Int) interface{} {
						return coer43(arg43)
					}
				}((func(_837_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_838_i11 _dafny.Int) _dafny.CodePoint {
						return _837_v0
					}
				})(_654_v0))), 3)
				(_nw134).ArraySet1(_679_v11, 4)
				(_nw134).ArraySet1(_679_v11, 5)
				(_nw134).ArraySet1(_679_v11, 6)
				(_nw134).ArraySet1(_679_v11, 7)
				(_nw134).ArraySet1(_679_v11, 8)
				(_nw134).ArraySet1(_679_v11, 9)
				(_nw134).ArraySet1(_679_v11, 10)
				(_nw134).ArraySet1(_679_v11, 11)
				(_nw134).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("uu"), 12)
				(_nw134).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(651))).Uint32(), func(coer44 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg44 _dafny.Int) interface{} {
						return coer44(arg44)
					}
				}((func(_839_v0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_840_i12 _dafny.Int) _dafny.CodePoint {
						return _839_v0
					}
				})(_654_v0))), 13)
				_836_v117 = _nw134
				var _841_v118 _dafny.Map
				_ = _841_v118
				_841_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_836_v117, (_this).F5())
				_841_v118 = (_841_v118).Update(_836_v117, _dafny.MultiSetOf((_this).F6()))
				var _842_v119 _dafny.Sequence
				_ = _842_v119
				_842_v119 = _dafny.SeqOf(false)
				_842_v119 = _842_v119
				(globalState).F0 = (p0).Minus((_dafny.Zero).Minus((func() _dafny.Int {
					if ((_677_v9).F5()).Contains(p1) {
						return ((_677_v9).F5()).Multiplicity(p1)
					}
					return p0
				})()))
			} else if _source9.Is_DC26() {
				var _843___mcc_h22 *C2 = _source9.Get_().(D9_DC26).Cf41
				_ = _843___mcc_h22
				var _844_cf41 *C2 = _843___mcc_h22
				_ = _844_cf41
				var _845_v120 *C1
				_ = _845_v120
				var _nw135 *C1 = New_C1_()
				_ = _nw135
				_nw135.Ctor__((_this).F6(), _679_v11, p1, (_677_v9).F5())
				_845_v120 = _nw135
				_654_v0 = (func() _dafny.CodePoint {
					if (_845_v120).F12() {
						return _654_v0
					}
					return _654_v0
				})()
				var _arr22 _dafny.Array = _this.F7()
				_ = _arr22
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index141
				(_arr22).ArraySet1((_845_v120).F12(), (_index141).Int())
				var _846_v121 _dafny.Map
				_ = _846_v121
				_846_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_845_v120).F13(), p2)
				var _847_v122 _dafny.Map
				_ = _847_v122
				_847_v122 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(634), _654_v0)
				var _848_v123 _dafny.Array
				_ = _848_v123
				var _nwElement0_19 _dafny.Int = p0
				_ = _nwElement0_19
				var _nw136 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(13))
				_ = _nw136
				(_nw136).ArraySet1(_nwElement0_19, 0)
				(_nw136).ArraySet1(p0, 1)
				(_nw136).ArraySet1(p0, 2)
				(_nw136).ArraySet1((_this).Fm6(globalState), 3)
				(_nw136).ArraySet1((_845_v120).Fm4(_dafny.IntOfInt64(109), p0, _846_v121, p0, globalState), 4)
				(_nw136).ArraySet1(p0, 5)
				(_nw136).ArraySet1(p0, 6)
				(_nw136).ArraySet1(((_847_v122).Cardinality()).Minus((_dafny.Zero).Minus(p0)), 7)
				(_nw136).ArraySet1(p0, 8)
				(_nw136).ArraySet1(p0, 9)
				(_nw136).ArraySet1(p0, 10)
				(_nw136).ArraySet1(p0, 11)
				(_nw136).ArraySet1((func() _dafny.Int {
					if (_656_v2).Contains(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nqbu")).Cardinality())) {
						return (_656_v2).Get(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nqbu")).Cardinality())).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_679_v11).Cardinality())
				})(), 12)
				_848_v123 = _nw136
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_848_v123), 0))
				_ = _index142
				(_848_v123).ArraySet1((_this).Fm4(p0, p0, _846_v121, (_678_v10).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_678_v10).Cardinality()))).Uint32()).(_dafny.Int), globalState), (_index142).Int())
				var _849_v124 _dafny.Set
				_ = _849_v124
				_849_v124 = _dafny.SetOf(_dafny.IntOfInt64(9))
				var _arr23 _dafny.Array = _this.F7()
				_ = _arr23
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index143
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_848_v123), 0))
				_ = _index144
				var _rhs163 _dafny.Int = p0
				_ = _rhs163
				var _rhs164 bool = p2
				_ = _rhs164
				var _rhs165 _dafny.Int = Companion_Default___.SafeModuloInt(((_849_v124).Difference(_849_v124)).Cardinality(), ((_677_v9).F5()).Cardinality())
				_ = _rhs165
				var _rhs166 bool = (_this).F6()
				_ = _rhs166
				var _lhs129 *GlobalState = globalState
				_ = _lhs129
				var _lhs130 _dafny.Array = _this.F7()
				_ = _lhs130
				var _lhs131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs131
				var _lhs132 _dafny.Array = _848_v123
				_ = _lhs132
				var _lhs133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen((_848_v123), 0))
				_ = _lhs133
				var _lhs134 *C4 = _this
				_ = _lhs134
				_lhs129.F0 = _rhs163
				(_lhs130).ArraySet1(_rhs164, (_lhs131).Int())
				(_lhs132).ArraySet1(_rhs165, (_lhs133).Int())
				_lhs134.F8_set_(_rhs166)
				var _850_v125 _dafny.Set
				_ = _850_v125
				_850_v125 = _dafny.SetOf((_this).F6())
				var _851_v126 _dafny.Set
				_ = _851_v126
				_851_v126 = _dafny.SetOf(_dafny.SetOf(p2, true), _850_v125)
				var _852_v127 _dafny.Map
				_ = _852_v127
				_852_v127 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_851_v126).Cardinality(), (_845_v120).F12())
				_852_v127 = (_852_v127).Update(p0, (_dafny.MultiSetOf((_this).Fm2((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(408), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), globalState), p1)).IsDisjointFrom((_677_v9).F5()))
			} else {
				var _853___mcc_h23 D9 = _source9.Get_().(D9_DC28).Cf44
				_ = _853___mcc_h23
				var _854_cf44 D9 = _853___mcc_h23
				_ = _854_cf44
				var _arr24 _dafny.Array = _this.F7()
				_ = _arr24
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index145
				(_arr24).ArraySet1(true, (_index145).Int())
				var _855_v128 *C3
				_ = _855_v128
				var _nw137 *C3 = New_C3_()
				_ = _nw137
				_nw137.Ctor__((_this).F6(), (_this).F5())
				_855_v128 = _nw137
				var _856_v129 _dafny.Map
				_ = _856_v129
				_856_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_855_v128, _this.F7())
				var _857_v130 _dafny.Sequence
				_ = _857_v130
				_857_v130 = _dafny.SeqOf(_this.F7(), (func() _dafny.Array {
					if (_856_v129).Contains(_855_v128) {
						return (_856_v129).Get(_855_v128).(_dafny.Array)
					}
					return _this.F7()
				})())
				var _858_v131 _dafny.Array
				_ = _858_v131
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_22
				var _nw138 _dafny.Array
				_ = _nw138
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw138 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) _dafny.Sequence = (func(_859_v10 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_860_i13 _dafny.Int) _dafny.Sequence {
							return _859_v10
						}
					})(_678_v10)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw138 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw138).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw138).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_858_v131 = _nw138
				var _861_v132 _dafny.Array
				_ = _861_v132
				var _nw139 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(10))
				_ = _nw139
				_861_v132 = _nw139
				var _862_v133 _dafny.Map
				_ = _862_v133
				_862_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_858_v131, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _861_v132)).Cardinality())
				var _arr25 _dafny.Array = _this.F7()
				_ = _arr25
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index146
				var _rhs167 bool = (_this).F6()
				_ = _rhs167
				var _rhs168 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_857_v130, _857_v130)).Cardinality())
				_ = _rhs168
				var _rhs169 _dafny.CodePoint = _654_v0
				_ = _rhs169
				var _rhs170 _dafny.Int = (func() _dafny.Int {
					if (_862_v133).Contains(_858_v131) {
						return (_862_v133).Get(_858_v131).(_dafny.Int)
					}
					return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ibbdsuctt")).Cardinality())
				})()
				_ = _rhs170
				var _rhs171 _dafny.Int = _dafny.IntOfInt64(635)
				_ = _rhs171
				var _lhs135 _dafny.Array = _this.F7()
				_ = _lhs135
				var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs136
				var _lhs137 *GlobalState = globalState
				_ = _lhs137
				var _lhs138 *GlobalState = globalState
				_ = _lhs138
				var _lhs139 *GlobalState = globalState
				_ = _lhs139
				(_lhs135).ArraySet1(_rhs167, (_lhs136).Int())
				_lhs137.F0 = _rhs168
				_654_v0 = _rhs169
				_lhs138.F0 = _rhs170
				_lhs139.F0 = _rhs171
				var _863_v134 _dafny.Array
				_ = _863_v134
				var _nw140 _dafny.Array = _dafny.NewArrayWithValue(Companion_D13_.Default(), _dafny.IntOfInt64(21))
				_ = _nw140
				_863_v134 = _nw140
				var _864_v135 _dafny.Sequence
				_ = _864_v135
				_864_v135 = _dafny.SeqOf(_863_v134, _863_v134, _863_v134)
				_863_v134 = (_864_v135).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(799), _dafny.IntOfUint32((_864_v135).Cardinality()))).Uint32()).(_dafny.Array)
				var _865_v136 _dafny.Map
				_ = _865_v136
				_865_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _654_v0)
				var _866_v137 _dafny.Map
				_ = _866_v137
				_866_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_865_v136, p0)
				var _867_v138 *C1
				_ = _867_v138
				var _nw141 *C1 = New_C1_()
				_ = _nw141
				_nw141.Ctor__(false, _679_v11, _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(p0, p0, p0, p0), _678_v10), (_677_v9).F5())
				_867_v138 = _nw141
				var _868_v139 D12
				_ = _868_v139
				_868_v139 = Companion_D12_.Create_DC35_(((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(7), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)) == ((_867_v138).F12()))
				var _869_v141 _dafny.Map
				_ = _869_v141
				_869_v141 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_865_v136, p1)
				var _870_v143 _dafny.Map
				_ = _870_v143
				_870_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll39 = _dafny.NewMapBuilder()
					_ = _coll39
					for _iter42 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(586), _dafny.IntOfInt64(-57))); ; {
						_compr_39, _ok42 := _iter42()
						if !_ok42 {
							break
						}
						var _871_v142 _dafny.Int
						_871_v142 = interface{}(_compr_39).(_dafny.Int)
						if ((_dafny.IntOfInt64(586)).Cmp(_871_v142) <= 0) && ((_871_v142).Cmp(_dafny.IntOfInt64(-57)) < 0) {
							_coll39.Add(Companion_Default___.SafeModuloInt(_871_v142, p0), _dafny.IntOfInt64(854))
						}
					}
					return _coll39.ToMap()
				}(), true)
				var _872_v144 *C0
				_ = _872_v144
				var _nw142 *C0 = New_C0_()
				_ = _nw142
				_nw142.Ctor__(_870_v143)
				_872_v144 = _nw142
				var _873_v145 _dafny.Set
				_ = _873_v145
				_873_v145 = _dafny.SetOf(_872_v144, _872_v144, _872_v144, _872_v144)
				var _rhs172 _dafny.Map = (_866_v137).Merge((func() _dafny.Map {
					var _coll40 = _dafny.NewMapBuilder()
					_ = _coll40
					for _iter43 := _dafny.Iterate((_869_v141).Keys().Elements()); ; {
						_compr_40, _ok43 := _iter43()
						if !_ok43 {
							break
						}
						var _874_v140 _dafny.Map
						_874_v140 = interface{}(_compr_40).(_dafny.Map)
						if (_869_v141).Contains(_874_v140) {
							_coll40.Add(_874_v140, p0)
						}
					}
					return _coll40.ToMap()
				}()).Merge(_866_v137))
				_ = _rhs172
				var _rhs173 _dafny.Int = Companion_Default___.SafeDivisionInt(p0, (p0).Plus((_dafny.Zero).Minus(p0)))
				_ = _rhs173
				var _rhs174 *C1 = _867_v138
				_ = _rhs174
				var _rhs175 _dafny.Int = ((func() _dafny.Set {
					if ((_867_v138).F12()) || (_this.F8()) {
						return _873_v145
					}
					return _873_v145
				})()).Cardinality()
				_ = _rhs175
				var _rhs176 D12 = func(_pat_let26_0 D12) D12 {
					return func(_875_dt__update__tmp_h1 D12) D12 {
						return func(_pat_let27_0 bool) D12 {
							return func(_876_dt__update_hcf55_h0 bool) D12 {
								return Companion_D12_.Create_DC35_(_876_dt__update_hcf55_h0)
							}(_pat_let27_0)
						}(_this.F8())
					}(_pat_let26_0)
				}(_868_v139)
				_ = _rhs176
				var _lhs140 *GlobalState = globalState
				_ = _lhs140
				var _lhs141 *GlobalState = globalState
				_ = _lhs141
				_866_v137 = _rhs172
				_lhs140.F0 = _rhs173
				_867_v138 = _rhs174
				_lhs141.F0 = _rhs175
				_868_v139 = _rhs176
				(globalState).F0 = p0
			}
			var _877_v146 _dafny.Map
			_ = _877_v146
			_877_v146 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("u"), _679_v11), p0)
			var _878_v147 _dafny.Sequence
			_ = _878_v147
			_878_v147 = _dafny.SeqOf(_877_v146)
			_877_v146 = (((_878_v147).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_678_v10).Cardinality()), _dafny.IntOfUint32((_878_v147).Cardinality()))).Uint32()).(_dafny.Map)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_679_v11, _dafny.IntOfInt64(777)))).Merge(_877_v146)
			if (func() bool {
				if !(_656_v2).Equals(func() _dafny.Map {
					var _coll41 = _dafny.NewMapBuilder()
					_ = _coll41
					for _iter44 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(849), _dafny.IntOfInt64(589))); ; {
						_compr_41, _ok44 := _iter44()
						if !_ok44 {
							break
						}
						var _879_v148 _dafny.Int
						_879_v148 = interface{}(_compr_41).(_dafny.Int)
						if ((_dafny.IntOfInt64(849)).Cmp(_879_v148) <= 0) && ((_879_v148).Cmp(_dafny.IntOfInt64(589)) < 0) {
							_coll41.Add(Companion_Default___.SafeDivisionInt(_879_v148, p0), p0)
						}
					}
					return _coll41.ToMap()
				}()) {
					return p2
				}
				return ((_this).F6()) && ((_this).F6())
			})() {
				var _880_v149 _dafny.Set
				_ = _880_v149
				_880_v149 = _dafny.SetOf(_679_v11, _679_v11, _679_v11, (_this).Fm7(globalState))
				_880_v149 = _880_v149
				var _881_v150 _dafny.Set
				_ = _881_v150
				_881_v150 = _dafny.SetOf(p2, _this.F8())
				var _882_v151 _dafny.Map
				_ = _882_v151
				_882_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(p1, p2, p1, _this.F8())).IsSubsetOf(_881_v150), p0)
				_882_v151 = (_882_v151).Update((func() bool {
					if false {
						return p2
					}
					return _this.F8()
				})(), p0)
				var _nw143 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
				_ = _nw143
				(_this).F7_set_(_nw143)
				var _883_v152 _dafny.Map
				_ = _883_v152
				_883_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _this.F8())
				var _884_v153 D9
				_ = _884_v153
				_884_v153 = Companion_D9_.Create_DC27_(_883_v152, p1)
				var _885_v154 _dafny.MultiSet
				_ = _885_v154
				_885_v154 = _dafny.MultiSetOf(Companion_D9_.Create_DC27_(_883_v152, _this.F8()), _884_v153)
				var _886_v155 D16
				_ = _886_v155
				_886_v155 = Companion_D16_.Create_DC43_(_885_v154)
				var _887_v156 _dafny.Sequence
				_ = _887_v156
				_887_v156 = _dafny.SeqOf((_dafny.MultiSetOf(_884_v153, _884_v153)).Difference(_885_v154), (_886_v155).Dtor_cf68(), (_dafny.MultiSetOf(_884_v153, _884_v153, _884_v153)).Difference(_885_v154))
				(globalState).F0 = ((_887_v156).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_887_v156).Cardinality()))).Uint32()).(_dafny.MultiSet)).Cardinality()
				var _888_v157 _dafny.Map
				_ = _888_v157
				_888_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_656_v2, _this.F8())
				var _889_v158 *C0
				_ = _889_v158
				var _nw144 *C0 = New_C0_()
				_ = _nw144
				_nw144.Ctor__(_888_v157)
				_889_v158 = _nw144
				_889_v158 = _889_v158
			} else {
				(_this).F8_set_((_this).F6())
				var _arr26 _dafny.Array = _this.F7()
				_ = _arr26
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index147
				(_arr26).ArraySet1((_dafny.SetOf(p0)).Contains(p0), (_index147).Int())
				var _arr27 _dafny.Array = _this.F7()
				_ = _arr27
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index148
				var _rhs177 _dafny.Int = (_this).Fm6(globalState)
				_ = _rhs177
				var _rhs178 _dafny.Int = (_dafny.IntOfInt64(122)).Plus(p0)
				_ = _rhs178
				var _rhs179 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_679_v11, (Companion_Default___.SafeIndex((_this).Fm6(globalState), _dafny.IntOfUint32((_679_v11).Cardinality()))).Uint32(), _654_v0)
				_ = _rhs179
				var _rhs180 bool = _this.F8()
				_ = _rhs180
				var _rhs181 bool = _this.F8()
				_ = _rhs181
				var _lhs142 *GlobalState = globalState
				_ = _lhs142
				var _lhs143 *GlobalState = globalState
				_ = _lhs143
				var _lhs144 *C4 = _this
				_ = _lhs144
				var _lhs145 _dafny.Array = _this.F7()
				_ = _lhs145
				var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs146
				_lhs142.F0 = _rhs177
				_lhs143.F0 = _rhs178
				_679_v11 = _rhs179
				_lhs144.F8_set_(_rhs180)
				(_lhs145).ArraySet1(_rhs181, (_lhs146).Int())
				var _890_v159 _dafny.Map
				_ = _890_v159
				var _out9 _dafny.Map
				_ = _out9
				_out9 = (_830_v115).M9(false, p0, globalState)
				_890_v159 = _out9
				(globalState).F0 = (_this).Fm6(globalState)
				var _891_v160 _dafny.Array
				_ = _891_v160
				var _len0_23 _dafny.Int = _dafny.IntOfInt64(10)
				_ = _len0_23
				var _nw145 _dafny.Array
				_ = _nw145
				if _len0_23.Cmp(_dafny.Zero) == 0 {
					_nw145 = _dafny.NewArray(_len0_23)
				} else {
					var _init23 func(_dafny.Int) _dafny.Int = (func(_892_p1 bool, _893_v11 _dafny.Sequence, _894_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_895_i14 _dafny.Int) _dafny.Int {
							return (_895_i14).Times((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_892_p1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_893_v11, _dafny.SetOf(_894_p0, (_dafny.Zero).Minus(_894_p0)))).Cardinality())).Cardinality())
						}
					})(p1, _679_v11, p0)
					_ = _init23
					var _element0_23 = _init23(_dafny.Zero)
					_ = _element0_23
					_nw145 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
					(_nw145).ArraySet1(_element0_23, 0)
					var _nativeLen0_23 = (_len0_23).Int()
					_ = _nativeLen0_23
					for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
						(_nw145).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
					}
				}
				_891_v160 = _nw145
				_891_v160 = _891_v160
			}
		}
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f7  _dafny.Array
	_f8  bool
	_f6  bool
	_f5  _dafny.MultiSet
	_f9  _dafny.Int
	_f10 bool
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f7 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f8 = false
	_this._f6 = false
	_this._f5 = _dafny.EmptyMultiSet
	_this._f9 = _dafny.Zero
	_this._f10 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C5{}
var _ T1 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F7() _dafny.Array {
	return _this._f7
}
func (_this *C5) F7_set_(value _dafny.Array) {
	_this._f7 = value
}
func (_this *C5) F8() bool {
	return _this._f8
}
func (_this *C5) F8_set_(value bool) {
	_this._f8 = value
}
func (_this *C5) F6() bool {
	return _this._f6
}
func (_this *C5) F5() _dafny.MultiSet {
	return _this._f5
}
func (_this *C5) Ctor__(f9 _dafny.Int, f10 bool, f7 _dafny.Array, f8 bool, f6 bool, f5 _dafny.MultiSet) {
	{
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f6 = f6
		(_this)._f5 = f5
	}
}
func (_this *C5) Fm3(globalState *GlobalState) bool {
	{
		return ((_this).F9()).Cmp((_this).F9()) >= 0
	}
}
func (_this *C5) Fm4(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Map, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("rnpbenu"), _dafny.UnicodeSeqOfUtf8Bytes("ygnckedh")), _dafny.UnicodeSeqOfUtf8Bytes("klfmjbsc"))).Cardinality())
	}
}
func (_this *C5) Fm1(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		return _this.F8()
	}
}
func (_this *C5) Fm2(p0 bool, globalState *GlobalState) bool {
	{
		return (Companion_Default___.SafeDivisionInt((_this).F9(), (_this).F9())).Cmp((_this).F9()) < 0
	}
}
func (_this *C5) M2(p0 _dafny.CodePoint, globalState *GlobalState) {
	{
		var _896_v0 _dafny.CodePoint
		_ = _896_v0
		_896_v0 = _dafny.CodePoint('r')
		_896_v0 = _896_v0
		var _897_v1 _dafny.Array
		_ = _897_v1
		var _len0_24 _dafny.Int = _dafny.IntOfInt64(29)
		_ = _len0_24
		var _nw146 _dafny.Array
		_ = _nw146
		if _len0_24.Cmp(_dafny.Zero) == 0 {
			_nw146 = _dafny.NewArray(_len0_24)
		} else {
			var _init24 func(_dafny.Int) _dafny.Sequence = (func(_898_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.Sequence {
				return func(_899_i0 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-489))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}((func(_900_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
						return func(_901_i1 _dafny.Int) _dafny.CodePoint {
							return _900_p0
						}
					})(_898_p0)))
				}
			})(p0)
			_ = _init24
			var _element0_24 = _init24(_dafny.Zero)
			_ = _element0_24
			_nw146 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
			(_nw146).ArraySet1(_element0_24, 0)
			var _nativeLen0_24 = (_len0_24).Int()
			_ = _nativeLen0_24
			for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
				(_nw146).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
			}
		}
		_897_v1 = _nw146
		_897_v1 = _897_v1
		var _902_v2 _dafny.Map
		_ = _902_v2
		_902_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F10())
		var _903_v3 _dafny.Sequence
		_ = _903_v3
		_903_v3 = _dafny.SeqOf(!(_this.F8()), (_this).Fm1((_this).F6(), (_this).F9(), (_902_v2).Cardinality(), globalState), (_this).F6(), true)
		var _904_v4 _dafny.Sequence
		_ = _904_v4
		_904_v4 = _dafny.UnicodeSeqOfUtf8Bytes("nhj")
		(_this).F8_set_((func() bool {
			if _this.F8() {
				return (_dafny.IntOfUint32((_903_v3).Cardinality())).Cmp((_this).F9()) < 0
			}
			return !_dafny.Companion_Sequence_.Equal(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer46 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_905_i2 _dafny.Int) _dafny.Sequence {
				return _dafny.UnicodeSeqOfUtf8Bytes("khisohk")
			})), _dafny.SeqOf(_904_v4))
		})())
		var _906_v5 _dafny.Map
		_ = _906_v5
		_906_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this.F8()) || (!((_this).F10())), _this.F7())
		_906_v5 = _906_v5
		if !(_this.F8()) {
			var _907_v6 _dafny.Sequence
			_ = _907_v6
			_907_v6 = _dafny.SeqOf(_this.F7(), _this.F7())
			var _908_v7 _dafny.Map
			_ = _908_v7
			_908_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_907_v6).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_this.F8())).Cardinality()), _dafny.IntOfUint32((_907_v6).Cardinality()))).Uint32()).(_dafny.Array), _896_v0)
			_896_v0 = (func() _dafny.CodePoint {
				if (_908_v7).Contains(_this.F7()) {
					return (_908_v7).Get(_this.F7()).(_dafny.CodePoint)
				}
				return p0
			})()
			var _909_v8 _dafny.Array
			_ = _909_v8
			var _nw147 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(23))
			_ = _nw147
			_909_v8 = _nw147
			var _910_v9 _dafny.Set
			_ = _910_v9
			_910_v9 = _dafny.SetOf(_dafny.IntOfUint32((_903_v3).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rkohfvwpq")).Cardinality()))
			var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_909_v8), 0))
			_ = _index149
			(_909_v8).ArraySet1(_910_v9, (_index149).Int())
			var _911_v10 _dafny.Map
			_ = _911_v10
			_911_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.SetOf((_this).F9(), (_this).F9(), (_this).F9()))
			var _912_v11 _dafny.Map
			_ = _912_v11
			_912_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (func() _dafny.Set {
				if (_911_v10).Contains((_this).F9()) {
					return (_911_v10).Get((_this).F9()).(_dafny.Set)
				}
				return _dafny.SetOf((_this).F9())
			})())
			var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(58), _dafny.ArrayLen((_909_v8), 0))
			_ = _index150
			(_909_v8).ArraySet1((_910_v9).Intersection((func() _dafny.Set {
				if (_912_v11).Contains((_this).F6()) {
					return (_912_v11).Get((_this).F6()).(_dafny.Set)
				}
				return _910_v9
			})()), (_index150).Int())
			if _this.F8() {
				(globalState).F0 = (_this).F9()
				var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_897_v1), 0))
				_ = _index151
				(_897_v1).ArraySet1(_904_v4, (_index151).Int())
				var _913_v12 _dafny.Map
				_ = _913_v12
				_913_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfInt64(-811))
				var _914_v13 _dafny.Sequence
				_ = _914_v13
				_914_v13 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jxhfdgy"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}((func(_915_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_916_i3 _dafny.Int) _dafny.CodePoint {
						return _915_p0
					}
				})(p0)))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_904_v4).Cardinality()), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("jxhfdgy"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer48 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg48 _dafny.Int) interface{} {
						return coer48(arg48)
					}
				}((func(_917_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_918_i3 _dafny.Int) _dafny.CodePoint {
						return _917_p0
					}
				})(p0))))).Cardinality()))).Uint32(), p0))
				var _919_v15 _dafny.Map
				_ = _919_v15
				_919_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_904_v4, _this.F8())
				var _920_v16 D0
				_ = _920_v16
				_920_v16 = Companion_D0_.Create_DC1_()
				var _921_v17 _dafny.Set
				_ = _921_v17
				_921_v17 = _dafny.SetOf(_920_v16)
				var _922_v18 _dafny.Set
				_ = _922_v18
				_922_v18 = _dafny.SetOf((func() bool {
					if (_902_v2).Contains((_this).F10()) {
						return (_902_v2).Get((_this).F10()).(bool)
					}
					return (func() bool {
						if (_902_v2).Contains((_this).F10()) {
							return (_902_v2).Get((_this).F10()).(bool)
						}
						return _this.F8()
					})()
				})(), (_this).F10(), _this.F8(), _this.F8())
				var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_897_v1), 0))
				_ = _index152
				var _rhs182 _dafny.Sequence = _904_v4
				_ = _rhs182
				var _rhs183 _dafny.Map = func() _dafny.Map {
					var _coll42 = _dafny.NewMapBuilder()
					_ = _coll42
					for _iter45 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-603), _dafny.IntOfInt64(543))); ; {
						_compr_42, _ok45 := _iter45()
						if !_ok45 {
							break
						}
						var _923_v14 _dafny.Int
						_923_v14 = interface{}(_compr_42).(_dafny.Int)
						if ((_dafny.IntOfInt64(-603)).Cmp(_923_v14) <= 0) && ((_923_v14).Cmp(_dafny.IntOfInt64(543)) < 0) {
							_coll42.Add(Companion_Default___.SafeDivisionInt(_923_v14, (_this).F9()), (Companion_D0_.Create_DC0_(_dafny.IntOfInt64(22))).Dtor_cf0())
						}
					}
					return _coll42.ToMap()
				}()
				_ = _rhs183
				var _rhs184 bool = (_this.F8()) == ((_dafny.SetOf((_902_v2).Cardinality(), _dafny.IntOfInt64(-810), (_this).F9())).IsDisjointFrom(_dafny.SetOf((_this).Fm4(_dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), (_this).F9(), _dafny.IntOfInt64(-857), (_this).F9(), (_this).F9())).Cardinality()), (_this).F9(), _919_v15, (_921_v17).Cardinality(), globalState))))
				_ = _rhs184
				var _rhs185 bool = (_922_v18).IsProperSubsetOf(_922_v18)
				_ = _rhs185
				var _rhs186 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_914_v13, Companion_Default___.Fm5(globalState))
				_ = _rhs186
				var _lhs147 _dafny.Array = _897_v1
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(206), _dafny.ArrayLen((_897_v1), 0))
				_ = _lhs148
				var _lhs149 *C5 = _this
				_ = _lhs149
				var _lhs150 *C5 = _this
				_ = _lhs150
				(_lhs147).ArraySet1(_rhs182, (_lhs148).Int())
				_913_v12 = _rhs183
				_lhs149.F8_set_(_rhs184)
				_lhs150.F8_set_(_rhs185)
				_914_v13 = _rhs186
				var _924_v19 _dafny.Map
				_ = _924_v19
				_924_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(172), (_this).F9()))
				var _925_v20 _dafny.Map
				_ = _925_v20
				_925_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_913_v12).Cardinality(), (func() _dafny.Map {
					if (_924_v19).Contains((_this).F6()) {
						return (_924_v19).Get((_this).F6()).(_dafny.Map)
					}
					return _913_v12
				})())
				var _926_v21 _dafny.Sequence
				_ = _926_v21
				_926_v21 = _dafny.SeqOf((func() _dafny.Map {
					if (_925_v20).Contains((_this).F9()) {
						return (_925_v20).Get((_this).F9()).(_dafny.Map)
					}
					return _913_v12
				})())
				var _927_v22 _dafny.MultiSet
				_ = _927_v22
				_927_v22 = _dafny.MultiSetOf((_this).F9())
				var _928_v23 _dafny.Sequence
				_ = _928_v23
				_928_v23 = _dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _896_v0)).Cardinality())
				var _929_v24 _dafny.Sequence
				_ = _929_v24
				_929_v24 = _dafny.SeqOf(_928_v23)
				var _930_v25 _dafny.Map
				_ = _930_v25
				_930_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.Companion_Sequence_.Update(_928_v23, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_928_v23).Cardinality()))).Uint32(), (_this).F9()))
				var _931_v26 _dafny.Array
				_ = _931_v26
				var _nwElement0_20 _dafny.Int = (_dafny.Zero).Minus((_this).F9())
				_ = _nwElement0_20
				var _nw148 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(24))
				_ = _nw148
				(_nw148).ArraySet1(_nwElement0_20, 0)
				(_nw148).ArraySet1((_this).F9(), 1)
				(_nw148).ArraySet1(((_this).F9()).Times((_this).F9()), 2)
				(_nw148).ArraySet1((_this).F9(), 3)
				(_nw148).ArraySet1((_this).F9(), 4)
				(_nw148).ArraySet1((_this).F9(), 5)
				(_nw148).ArraySet1(((_this).F9()).Minus((_this).F9()), 6)
				(_nw148).ArraySet1(((_dafny.Zero).Minus((_this).F9())).Times((_this).F9()), 7)
				(_nw148).ArraySet1((_dafny.Zero).Minus(((_926_v21).Select((Companion_Default___.SafeIndex((_927_v22).Cardinality(), _dafny.IntOfUint32((_926_v21).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality()), 8)
				(_nw148).ArraySet1((_dafny.Zero).Minus((_this).F9()), 9)
				(_nw148).ArraySet1((_this).F9(), 10)
				(_nw148).ArraySet1(_dafny.IntOfInt64(-585), 11)
				(_nw148).ArraySet1((_this).F9(), 12)
				(_nw148).ArraySet1((_this).F9(), 13)
				(_nw148).ArraySet1((_this).F9(), 14)
				(_nw148).ArraySet1(_dafny.IntOfUint32(((_929_v24).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_904_v4, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_904_v4).Cardinality()))).Uint32(), p0)).Cardinality()), _dafny.IntOfUint32((_929_v24).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), 15)
				(_nw148).ArraySet1(((_this).F9()).Plus((_930_v25).Cardinality()), 16)
				(_nw148).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(470), _dafny.IntOfInt64(97)), 17)
				(_nw148).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_904_v4, _dafny.UnicodeSeqOfUtf8Bytes("vhffqpp")), (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_904_v4, _dafny.UnicodeSeqOfUtf8Bytes("vhffqpp"))).Cardinality()))).Uint32(), p0)).Cardinality()), 18)
				(_nw148).ArraySet1((_this).F9(), 19)
				(_nw148).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(98), (_this).F9()), 20)
				(_nw148).ArraySet1((_this).F9(), 21)
				(_nw148).ArraySet1((_this).F9(), 22)
				(_nw148).ArraySet1(_dafny.IntOfInt64(-674), 23)
				_931_v26 = _nw148
				var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_931_v26), 0))
				_ = _index153
				(_931_v26).ArraySet1((_dafny.Zero).Minus((_this).F9()), (_index153).Int())
				var _932_v27 D0
				_ = _932_v27
				_932_v27 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_903_v3).Cardinality()))
				var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_931_v26), 0))
				_ = _index154
				(_931_v26).ArraySet1((_932_v27).Dtor_cf0(), (_index154).Int())
				(globalState).F0 = (_this).F9()
				var _933_v28 T2
				_ = _933_v28
				var _nw149 *C4 = New_C4_()
				_ = _nw149
				_nw149.Ctor__(_this.F7(), (_this).F10(), !(_this.F8()), (_this).F5())
				_933_v28 = _nw149
				var _934_v29 _dafny.Array
				_ = _934_v29
				var _nwElement0_21 T2 = _933_v28
				_ = _nwElement0_21
				var _nw150 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(10))
				_ = _nw150
				(_nw150).ArraySet1(_nwElement0_21, 0)
				(_nw150).ArraySet1(_933_v28, 1)
				(_nw150).ArraySet1(_933_v28, 2)
				(_nw150).ArraySet1(_933_v28, 3)
				(_nw150).ArraySet1(_933_v28, 4)
				(_nw150).ArraySet1(_933_v28, 5)
				(_nw150).ArraySet1(_933_v28, 6)
				(_nw150).ArraySet1(_933_v28, 7)
				(_nw150).ArraySet1(_933_v28, 8)
				(_nw150).ArraySet1(_933_v28, 9)
				_934_v29 = _nw150
				var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_934_v29), 0))
				_ = _index155
				(_934_v29).ArraySet1(_933_v28, (_index155).Int())
				var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(808), _dafny.ArrayLen((_934_v29), 0))
				_ = _index156
				var _nw151 *C4 = New_C4_()
				_ = _nw151
				_nw151.Ctor__(_933_v28.F7(), ((_933_v28).F5()).IsProperSubsetOf(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_903_v3, (Companion_Default___.SafeIndex((_931_v26).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(882), _dafny.ArrayLen((_931_v26), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_903_v3).Cardinality()))).Uint32(), (_933_v28).F6()))), _933_v28.F8(), (_dafny.MultiSetOf(_this.F8(), (_this).F6())).Difference((_this).F5()))
				(_934_v29).ArraySet1(_nw151, (_index156).Int())
			} else {
				var _935_v30 _dafny.Array
				_ = _935_v30
				var _nw152 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
				_ = _nw152
				_935_v30 = _nw152
				_935_v30 = _935_v30
				(_this).F8_set_(false)
				(_this).F8_set_(((_this).F5()).IsDisjointFrom((_this).F5()))
				var _arr28 _dafny.Array = _this.F7()
				_ = _arr28
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index157
				(_arr28).ArraySet1((_this).F6(), (_index157).Int())
				var _arr29 _dafny.Array = _this.F7()
				_ = _arr29
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index158
				var _rhs187 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("ooxmqc")
				_ = _rhs187
				var _rhs188 bool = (_this).F6()
				_ = _rhs188
				var _rhs189 bool = _this.F8()
				_ = _rhs189
				var _rhs190 _dafny.Int = (_this).F9()
				_ = _rhs190
				var _rhs191 _dafny.CodePoint = _896_v0
				_ = _rhs191
				var _lhs151 *C5 = _this
				_ = _lhs151
				var _lhs152 _dafny.Array = _this.F7()
				_ = _lhs152
				var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs153
				var _lhs154 *GlobalState = globalState
				_ = _lhs154
				_904_v4 = _rhs187
				_lhs151.F8_set_(_rhs188)
				(_lhs152).ArraySet1(_rhs189, (_lhs153).Int())
				_lhs154.F0 = _rhs190
				_896_v0 = _rhs191
				var _arr30 _dafny.Array = _this.F7()
				_ = _arr30
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(916), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index159
				(_arr30).ArraySet1(((_this).F5()).IsProperSubsetOf(((_this).F5()).Update(_this.F8(), Companion_Default___.Abs((_dafny.Zero).Minus((_this).F9())))), (_index159).Int())
			}
			(globalState).F0 = (_this).F9()
			var _936_v31 _dafny.Array
			_ = _936_v31
			var _len0_25 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_25
			var _nw153 _dafny.Array
			_ = _nw153
			if _len0_25.Cmp(_dafny.Zero) == 0 {
				_nw153 = _dafny.NewArray(_len0_25)
			} else {
				var _init25 func(_dafny.Int) _dafny.Map = func(_937_i4 _dafny.Int) _dafny.Map {
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F10())
				}
				_ = _init25
				var _element0_25 = _init25(_dafny.Zero)
				_ = _element0_25
				_nw153 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
				(_nw153).ArraySet1(_element0_25, 0)
				var _nativeLen0_25 = (_len0_25).Int()
				_ = _nativeLen0_25
				for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
					(_nw153).ArraySet1(_init25(_dafny.IntOf(_i0_25)), _i0_25)
				}
			}
			_936_v31 = _nw153
			var _rhs192 _dafny.Int = (_this).F9()
			_ = _rhs192
			var _rhs193 _dafny.Array = _936_v31
			_ = _rhs193
			var _lhs155 *GlobalState = globalState
			_ = _lhs155
			_lhs155.F0 = _rhs192
			_936_v31 = _rhs193
		} else {
			(_this).F8_set_(false)
			var _938_v32 *C3
			_ = _938_v32
			var _nw154 *C3 = New_C3_()
			_ = _nw154
			_nw154.Ctor__(((func() bool {
				if (_902_v2).Contains((_this).F6()) {
					return (_902_v2).Get((_this).F6()).(bool)
				}
				return (_this).F10()
			})()) == (_this.F8()), (_this).F5())
			_938_v32 = _nw154
			(_this).F8_set_(!(_this.F8()) || ((_this).F10()))
			var _939_v33 _dafny.Array
			_ = _939_v33
			var _nw155 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
			_ = _nw155
			_939_v33 = _nw155
			var _940_v34 _dafny.Sequence
			_ = _940_v34
			_940_v34 = _dafny.SeqOf(_this.F7())
			var _941_v35 _dafny.Set
			_ = _941_v35
			_941_v35 = _dafny.SetOf((_this).F9())
			var _942_v36 _dafny.MultiSet
			_ = _942_v36
			_942_v36 = _dafny.MultiSetOf(_this.F7(), _this.F7(), (_940_v34).Select((Companion_Default___.SafeIndex((_941_v35).Cardinality(), _dafny.IntOfUint32((_940_v34).Cardinality()))).Uint32()).(_dafny.Array), _this.F7(), _this.F7())
			var _943_v37 _dafny.Map
			_ = _943_v37
			_943_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_939_v33, (_dafny.Zero).Minus(((_this).F9()).Times((func() _dafny.Int {
				if (_942_v36).Contains(_this.F7()) {
					return (_942_v36).Multiplicity(_this.F7())
				}
				return (_this).F9()
			})())))
			var _944_v38 _dafny.Map
			_ = _944_v38
			_944_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfInt64(289))
			var _945_v39 *C1
			_ = _945_v39
			var _nw156 *C1 = New_C1_()
			_ = _nw156
			_nw156.Ctor__(_this.F8(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(500))).Uint32(), func(coer49 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg49 _dafny.Int) interface{} {
					return coer49(arg49)
				}
			}((func(_946_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_947_i5 _dafny.Int) _dafny.CodePoint {
					return _946_p0
				}
			})(p0))), (_938_v32).Fm3(globalState), Companion_Default___.Fm14(_944_v38, _this.F8(), (_this).F9(), (_this).F10(), globalState))
			_945_v39 = _nw156
			var _948_v40 _dafny.MultiSet
			_ = _948_v40
			_948_v40 = _dafny.MultiSetOf(_945_v39)
			_943_v37 = (_943_v37).Update(_939_v33, (_948_v40).Cardinality())
			var _949_v41 _dafny.MultiSet
			_ = _949_v41
			_949_v41 = _dafny.MultiSetOf((_this).F9())
			var _950_v42 D2
			_ = _950_v42
			_950_v42 = Companion_D2_.Create_DC7_(p0, _944_v38, (_this).F9())
			var _951_v43 _dafny.Map
			_ = _951_v43
			_951_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-980))).Uint32(), func(coer50 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg50 _dafny.Int) interface{} {
					return coer50(arg50)
				}
			}((func(_952_p0 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_953_i6 _dafny.Int) _dafny.CodePoint {
					return _952_p0
				}
			})(p0))), _this.F8())
			(globalState).F0 = (_945_v39).Fm4((func() _dafny.Int {
				if (_949_v41).Contains(_dafny.IntOfInt64(608)) {
					return (_949_v41).Multiplicity(_dafny.IntOfInt64(608))
				}
				return (_this).F9()
			})(), (_this).F9(), (Companion_Default___.Fm13((_this).F9(), _950_v42, (_this).F10(), _dafny.IntOfUint32((_903_v3).Cardinality()), globalState)).Merge(_951_v43), ((_this).F9()).Times((_dafny.Zero).Minus((_this).F9())), globalState)
		}
		var _954_v44 _dafny.Map
		_ = _954_v44
		_954_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(588), (_this).F6())).Cardinality()), (_this).F9())
		var _955_v45 D2
		_ = _955_v45
		_955_v45 = Companion_D2_.Create_DC7_(_896_v0, _954_v44, (_this).F9())
		var _956_v46 D12
		_ = _956_v46
		_956_v46 = Companion_D12_.Create_DC36_(_dafny.IntOfInt64(321), (_this).F9())
		var _957_v47 D15
		_ = _957_v47
		_957_v47 = Companion_D15_.Create_DC41_(Companion_Default___.Fm13((_this).F9(), _955_v45, (_this).F10(), (func(_pat_let28_0 D12) D12 {
			return func(_958_dt__update__tmp_h0 D12) D12 {
				return func(_pat_let29_0 _dafny.Int) D12 {
					return func(_959_dt__update_hcf57_h0 _dafny.Int) D12 {
						return Companion_D12_.Create_DC36_((_958_dt__update__tmp_h0).Dtor_cf56(), _959_dt__update_hcf57_h0)
					}(_pat_let29_0)
				}((_this).F9())
			}(_pat_let28_0)
		}(_956_v46)).Dtor_cf57(), globalState))
		var _960_i7 _dafny.Int
		_ = _960_i7
		_960_i7 = _dafny.Zero
		{
			for func(_source10 D15) bool {
				if _source10.Is_DC42() {
					var _963___mcc_h0 _dafny.MultiSet = _source10.Get_().(D15_DC42).Cf67
					_ = _963___mcc_h0
					var _964_cf67 _dafny.MultiSet = _963___mcc_h0
					_ = _964_cf67
					return ((_dafny.SetOf((_this).F10(), _this.F8())).Cardinality()).Cmp((_dafny.SetOf((_this).F10())).Cardinality()) == 0
				} else {
					var _965___mcc_h1 _dafny.Map = _source10.Get_().(D15_DC41).Cf66
					_ = _965___mcc_h1
					var _966_cf66 _dafny.Map = _965___mcc_h1
					_ = _966_cf66
					return (_this).F6()
				}
			}(_957_v47) {
				{
					if (_960_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L9
					}
					_960_i7 = (_960_i7).Plus(_dafny.One)
					var _961_v48 *C1
					_ = _961_v48
					var _nw157 *C1 = New_C1_()
					_ = _nw157
					_nw157.Ctor__((func() bool {
						if _this.F8() {
							return _this.F8()
						}
						return _this.F8()
					})(), _dafny.Companion_Sequence_.Concatenate(_904_v4, _dafny.UnicodeSeqOfUtf8Bytes("ceqq")), (_this).F6(), (_this).F5())
					_961_v48 = _nw157
					var _962_v49 *C4
					_ = _962_v49
					var _nw158 *C4 = New_C4_()
					_ = _nw158
					_nw158.Ctor__(_this.F7(), _this.F8(), (_this).F10(), ((_this).F5()).Union(_dafny.MultiSetOf(!((_this).F6()))))
					_962_v49 = _nw158
					(_this).F8_set_((_this).F10())
					var _arr31 _dafny.Array = _this.F7()
					_ = _arr31
					var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_this.F7()), 0))
					_ = _index160
					(_arr31).ArraySet1(!((_962_v49).Fm3(globalState)), (_index160).Int())
					var _arr32 _dafny.Array = _this.F7()
					_ = _arr32
					var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(379), _dafny.ArrayLen((_this.F7()), 0))
					_ = _index161
					(_arr32).ArraySet1(!_dafny.Companion_Sequence_.Contains(_903_v3, (func() bool {
						if (_this).F10() {
							return (_961_v48).F12()
						}
						return (_this).F6()
					})()), (_index161).Int())
					goto C9
				}
			C9:
			}
			goto L9
		}
	L9:
	}
}
func (_this *C5) M3(globalState *GlobalState) (_dafny.Set, _dafny.Int, _dafny.Int) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _arr33 _dafny.Array = _this.F7()
		_ = _arr33
		var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index162
		(_arr33).ArraySet1(!((_this).F10()), (_index162).Int())
		var _967_v0 _dafny.Sequence
		_ = _967_v0
		_967_v0 = _dafny.SeqOf((_this).F6(), (_this).F10())
		var _968_v1 D10
		_ = _968_v1
		_968_v1 = Companion_D10_.Create_DC29_(_dafny.MultiSetFromSeq(_dafny.SeqOf(_967_v0)))
		var _969_v2 _dafny.MultiSet
		_ = _969_v2
		_969_v2 = _dafny.MultiSetOf(_968_v1, _968_v1)
		var _arr34 _dafny.Array = _this.F7()
		_ = _arr34
		var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index163
		(_arr34).ArraySet1(!(((_969_v2).Intersection(_969_v2)).IsSubsetOf(_969_v2)), (_index163).Int())
		var _970_v3 D3
		_ = _970_v3
		_970_v3 = Companion_D3_.Create_DC10_(!(_this.F8()), (_this).F9())
		var _arr35 _dafny.Array = _this.F7()
		_ = _arr35
		var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index164
		(_arr35).ArraySet1(!(!(func(_source11 D3) bool {
			if _source11.Is_DC10() {
				var _971___mcc_h0 bool = _source11.Get_().(D3_DC10).Cf11
				_ = _971___mcc_h0
				var _972___mcc_h1 _dafny.Int = _source11.Get_().(D3_DC10).Cf12
				_ = _972___mcc_h1
				var _973_cf12 _dafny.Int = _972___mcc_h1
				_ = _973_cf12
				var _974_cf11 bool = _971___mcc_h0
				_ = _974_cf11
				return (_this).F10()
			} else if _source11.Is_DC11() {
				var _975___mcc_h2 _dafny.Int = _source11.Get_().(D3_DC11).Cf13
				_ = _975___mcc_h2
				var _976_cf13 _dafny.Int = _975___mcc_h2
				_ = _976_cf13
				return (_this).F10()
			} else {
				var _977___mcc_h3 _dafny.Array = _source11.Get_().(D3_DC9).Cf10
				_ = _977___mcc_h3
				var _978_cf10 _dafny.Array = _977___mcc_h3
				_ = _978_cf10
				return (_this).F6()
			}
		}(_970_v3))), (_index164).Int())
		var _979_v4 _dafny.Sequence
		_ = _979_v4
		_979_v4 = _dafny.UnicodeSeqOfUtf8Bytes("yie")
		var _980_v5 _dafny.Map
		_ = _980_v5
		_980_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(_979_v4, _dafny.UnicodeSeqOfUtf8Bytes("kynauahc")), (_this).F6())
		_980_v5 = (_980_v5).Update(_979_v4, !(_this.F8()))
		var _981_v6 _dafny.MultiSet
		_ = _981_v6
		_981_v6 = _dafny.MultiSetOf((_this).F9())
		var _982_v7 _dafny.Array
		_ = _982_v7
		var _nwElement0_22 _dafny.Sequence = _967_v0
		_ = _nwElement0_22
		var _nw159 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(17))
		_ = _nw159
		(_nw159).ArraySet1(_nwElement0_22, 0)
		(_nw159).ArraySet1(_967_v0, 1)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_967_v0, _967_v0), 2)
		(_nw159).ArraySet1(_967_v0, 3)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(!(true), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)), _967_v0), 4)
		(_nw159).ArraySet1(Companion_Default___.Fm23((func() _dafny.Int {
			if (_981_v6).Contains((_this).F9()) {
				return (_981_v6).Multiplicity((_this).F9())
			}
			return (_this).F9()
		})(), globalState), 5)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_967_v0, _967_v0), 6)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Update(_967_v0, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_967_v0).Cardinality()))).Uint32(), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)), 7)
		(_nw159).ArraySet1(_967_v0, 8)
		(_nw159).ArraySet1(_dafny.SeqOf((_this).F6()), 9)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_967_v0, _967_v0), 10)
		(_nw159).ArraySet1(_967_v0, 11)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_967_v0, _967_v0), 12)
		(_nw159).ArraySet1(Companion_Default___.Fm23((_this).F9(), globalState), 13)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_967_v0, _dafny.Companion_Sequence_.Update(_967_v0, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_967_v0).Cardinality()))).Uint32(), false)), 14)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_this.F8()), _dafny.Companion_Sequence_.Update(_dafny.SeqOf((_this).F10(), !((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))), (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_dafny.SeqOf((_this).F10(), !((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)))).Cardinality()))).Uint32(), (_this).F10())), 15)
		(_nw159).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm23(_dafny.IntOfInt64(712), globalState), _967_v0), 16)
		_982_v7 = _nw159
		for _iter46 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_982_v7), 0))); ; {
			_guard_loop_3, _ok46 := _iter46()
			if !_ok46 {
				break
			}
			var _983_i0 _dafny.Int
			_983_i0 = interface{}(_guard_loop_3).(_dafny.Int)
			if (true) && (((_983_i0).Sign() != -1) && ((_983_i0).Cmp(_dafny.ArrayLen((_982_v7), 0)) < 0)) {
				(_982_v7).ArraySet1(_967_v0, (_983_i0).Int())
			}
		}
		var _984_v8 D0
		_ = _984_v8
		_984_v8 = Companion_D0_.Create_DC1_()
		_984_v8 = _984_v8
		var _985_v9 _dafny.CodePoint
		_ = _985_v9
		_985_v9 = _dafny.CodePoint('b')
		var _986_v10 D10
		_ = _986_v10
		_986_v10 = Companion_D10_.Create_DC31_(_dafny.SeqOf((_this).F5()), _985_v9)
		(_this).F8_set_(func(_source12 D10) bool {
			if _source12.Is_DC30() {
				var _987___mcc_h4 bool = _source12.Get_().(D10_DC30).Cf46
				_ = _987___mcc_h4
				var _988___mcc_h5 bool = _source12.Get_().(D10_DC30).Cf47
				_ = _988___mcc_h5
				var _989___mcc_h6 bool = _source12.Get_().(D10_DC30).Cf48
				_ = _989___mcc_h6
				var _990_cf48 bool = _989___mcc_h6
				_ = _990_cf48
				var _991_cf47 bool = _988___mcc_h5
				_ = _991_cf47
				var _992_cf46 bool = _987___mcc_h4
				_ = _992_cf46
				return (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(238), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)
			} else if _source12.Is_DC31() {
				var _993___mcc_h7 _dafny.Sequence = _source12.Get_().(D10_DC31).Cf49
				_ = _993___mcc_h7
				var _994___mcc_h8 _dafny.CodePoint = _source12.Get_().(D10_DC31).Cf50
				_ = _994___mcc_h8
				var _995_cf50 _dafny.CodePoint = _994___mcc_h8
				_ = _995_cf50
				var _996_cf49 _dafny.Sequence = _993___mcc_h7
				_ = _996_cf49
				return (_this).F10()
			} else {
				var _997___mcc_h9 _dafny.MultiSet = _source12.Get_().(D10_DC29).Cf45
				_ = _997___mcc_h9
				var _998_cf45 _dafny.MultiSet = _997___mcc_h9
				_ = _998_cf45
				return (_this).F6()
			}
		}(_986_v10))
		var _999_v11 D12
		_ = _999_v11
		_999_v11 = Companion_D12_.Create_DC34_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_985_v9, _this.F8()))
		var _1000_v12 _dafny.Set
		_ = _1000_v12
		_1000_v12 = _dafny.SetOf((_dafny.Zero).Minus((_this).F9()), _dafny.IntOfInt64(32), Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(39), (_this).F9()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _999_v11)).Cardinality(), _dafny.IntOfUint32((_979_v4).Cardinality()))
		r0 = _1000_v12
		var _1001_v13 _dafny.Array
		_ = _1001_v13
		var _len0_26 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_26
		var _nw160 _dafny.Array
		_ = _nw160
		if _len0_26.Cmp(_dafny.Zero) == 0 {
			_nw160 = _dafny.NewArray(_len0_26)
		} else {
			var _init26 func(_dafny.Int) _dafny.Int = func(_1002_i1 _dafny.Int) _dafny.Int {
				return Companion_Default___.SafeModuloInt(_1002_i1, (_this).F9())
			}
			_ = _init26
			var _element0_26 = _init26(_dafny.Zero)
			_ = _element0_26
			_nw160 = _dafny.NewArrayFromExample(_element0_26, nil, _len0_26)
			(_nw160).ArraySet1(_element0_26, 0)
			var _nativeLen0_26 = (_len0_26).Int()
			_ = _nativeLen0_26
			for _i0_26 := 1; _i0_26 < _nativeLen0_26; _i0_26++ {
				(_nw160).ArraySet1(_init26(_dafny.IntOf(_i0_26)), _i0_26)
			}
		}
		_1001_v13 = _nw160
		var _1003_v14 _dafny.MultiSet
		_ = _1003_v14
		_1003_v14 = _dafny.MultiSetOf(_1001_v13, _1001_v13, _1001_v13)
		var _1004_v15 _dafny.Map
		_ = _1004_v15
		_1004_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_1003_v14).Update(_1001_v13, Companion_Default___.Abs((_this).F9())))
		r1 = (_this).Fm4(_dafny.IntOfInt64(217), (_this).F9(), _980_v5, (((func() _dafny.MultiSet {
			if (_1004_v15).Contains((_this).F9()) {
				return (_1004_v15).Get((_this).F9()).(_dafny.MultiSet)
			}
			return _1003_v14
		})()).Union(_1003_v14)).Cardinality(), globalState)
		r2 = (_this).F9()
		return r0, r1, r2
	}
}
func (_this *C5) M1(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) {
	{
		var _1005_i0 _dafny.Int
		_ = _1005_i0
		_1005_i0 = _dafny.Zero
		{
			for (_this).F10() {
				{
					if (_1005_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L10
					}
					_1005_i0 = (_1005_i0).Plus(_dafny.One)
					var _1006_v0 _dafny.Array
					_ = _1006_v0
					var _len0_27 _dafny.Int = _dafny.IntOfInt64(21)
					_ = _len0_27
					var _nw161 _dafny.Array
					_ = _nw161
					if _len0_27.Cmp(_dafny.Zero) == 0 {
						_nw161 = _dafny.NewArray(_len0_27)
					} else {
						var _init27 func(_dafny.Int) _dafny.Int = func(_1007_i1 _dafny.Int) _dafny.Int {
							return (_1007_i1).Times((_this).F9())
						}
						_ = _init27
						var _element0_27 = _init27(_dafny.Zero)
						_ = _element0_27
						_nw161 = _dafny.NewArrayFromExample(_element0_27, nil, _len0_27)
						(_nw161).ArraySet1(_element0_27, 0)
						var _nativeLen0_27 = (_len0_27).Int()
						_ = _nativeLen0_27
						for _i0_27 := 1; _i0_27 < _nativeLen0_27; _i0_27++ {
							(_nw161).ArraySet1(_init27(_dafny.IntOf(_i0_27)), _i0_27)
						}
					}
					_1006_v0 = _nw161
					var _1008_v1 D3
					_ = _1008_v1
					_1008_v1 = Companion_D3_.Create_DC9_(_1006_v0)
					var _1009_v2 _dafny.Sequence
					_ = _1009_v2
					_1009_v2 = _dafny.SeqOf(p0)
					var _1010_v3 _dafny.Sequence
					_ = _1010_v3
					_1010_v3 = _dafny.SeqOf(!((_this).F6()))
					var _1011_v4 _dafny.Map
					_ = _1011_v4
					_1011_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(487), _dafny.IntOfInt64(828))
					var _1012_v5 _dafny.Map
					_ = _1012_v5
					_1012_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1011_v4, (_this).F6())
					var _1013_v6 *C0
					_ = _1013_v6
					var _nw162 *C0 = New_C0_()
					_ = _nw162
					_nw162.Ctor__(_1012_v5)
					_1013_v6 = _nw162
					var _1014_v7 _dafny.Set
					_ = _1014_v7
					_1014_v7 = _dafny.SetOf((_this).F6(), _this.F8(), (_this).F6(), (_this).F6(), p2)
					var _rhs194 _dafny.Int = (p0).Times((Companion_D5_.Create_DC16_(_1008_v1, _dafny.IntOfUint32((_1009_v2).Cardinality()), _dafny.IntOfInt64(40), _dafny.IntOfUint32((_1010_v3).Cardinality()), _1013_v6)).Dtor_cf20())
					_ = _rhs194
					var _rhs195 bool = _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p0, _dafny.IntOfInt64(-625)), (func() _dafny.Int {
						if (_1011_v4).Contains(p0) {
							return (_1011_v4).Get(p0).(_dafny.Int)
						}
						return (_1014_v7).Cardinality()
					})())
					_ = _rhs195
					var _rhs196 _dafny.Int = (_1009_v2).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1009_v2).Cardinality()))).Uint32()).(_dafny.Int)
					_ = _rhs196
					var _rhs197 bool = (p0).Cmp((func() _dafny.Int {
						if ((_this).F5()).Contains(p1) {
							return ((_this).F5()).Multiplicity(p1)
						}
						return (_this).F9()
					})()) > 0
					_ = _rhs197
					var _lhs156 *GlobalState = globalState
					_ = _lhs156
					var _lhs157 *C5 = _this
					_ = _lhs157
					var _lhs158 *GlobalState = globalState
					_ = _lhs158
					var _lhs159 *C5 = _this
					_ = _lhs159
					_lhs156.F0 = _rhs194
					_lhs157.F8_set_(_rhs195)
					_lhs158.F0 = _rhs196
					_lhs159.F8_set_(_rhs197)
					(globalState).F0 = (_this).F9()
					var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_1006_v0), 0))
					_ = _index165
					(_1006_v0).ArraySet1((p0).Minus(p0), (_index165).Int())
					var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(417), _dafny.ArrayLen((_1006_v0), 0))
					_ = _index166
					(_1006_v0).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(p0, Companion_Default___.SafeModuloInt(p0, (_dafny.Zero).Minus(p0)))), (_index166).Int())
					var _1015_v8 _dafny.Sequence
					_ = _1015_v8
					_1015_v8 = _dafny.UnicodeSeqOfUtf8Bytes("bobaat")
					_1015_v8 = (func() _dafny.Sequence {
						if false {
							return _1015_v8
						}
						return _dafny.Companion_Sequence_.Concatenate(_1015_v8, _1015_v8)
					})()
					goto C10
				}
			C10:
			}
			goto L10
		}
	L10:
		var _1016_v9 _dafny.Set
		_ = _1016_v9
		_1016_v9 = _dafny.SetOf((_this).F6())
		var _1017_v10 _dafny.Sequence
		_ = _1017_v10
		_1017_v10 = _dafny.UnicodeSeqOfUtf8Bytes("rhu")
		var _1018_v11 D5
		_ = _1018_v11
		_1018_v11 = Companion_D5_.Create_DC15_(_1017_v10)
		var _1019_v12 _dafny.Map
		_ = _1019_v12
		_1019_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1018_v11).Dtor_cf18(), (_this).F10())
		(globalState).F0 = ((_this).Fm4((_1016_v9).Cardinality(), (_this).F9(), _1019_v12, (_dafny.Zero).Minus((_this).F9()), globalState)).Times(p0)
		var _pat_let_tv20 = p0
		_ = _pat_let_tv20
		var _pat_let_tv21 = p0
		_ = _pat_let_tv21
		var _pat_let_tv22 = p1
		_ = _pat_let_tv22
		var _pat_let_tv23 = p1
		_ = _pat_let_tv23
		var _pat_let_tv24 = p0
		_ = _pat_let_tv24
		(globalState).F0 = func(_source13 D7) _dafny.Int {
			if _source13.Is_DC20() {
				var _1020___mcc_h0 bool = _source13.Get_().(D7_DC20).Cf29
				_ = _1020___mcc_h0
				var _1021_cf29 bool = _1020___mcc_h0
				_ = _1021_cf29
				return Companion_Default___.SafeModuloInt(_pat_let_tv20, _pat_let_tv21)
			} else if _source13.Is_DC21() {
				var _1022___mcc_h1 _dafny.Int = _source13.Get_().(D7_DC21).Cf30
				_ = _1022___mcc_h1
				var _1023___mcc_h2 _dafny.Int = _source13.Get_().(D7_DC21).Cf31
				_ = _1023___mcc_h2
				var _1024_cf31 _dafny.Int = _1023___mcc_h2
				_ = _1024_cf31
				var _1025_cf30 _dafny.Int = _1022___mcc_h1
				_ = _1025_cf30
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _pat_let_tv22)).Cardinality()
			} else if _source13.Is_DC19() {
				var _1026___mcc_h3 _dafny.Map = _source13.Get_().(D7_DC19).Cf28
				_ = _1026___mcc_h3
				var _1027_cf28 _dafny.Map = _1026___mcc_h3
				_ = _1027_cf28
				return (_this).F9()
			} else {
				var _1028___mcc_h4 D7 = _source13.Get_().(D7_DC22).Cf32
				_ = _1028___mcc_h4
				var _1029_cf32 D7 = _1028___mcc_h4
				_ = _1029_cf32
				return Companion_Default___.SafeModuloInt((_dafny.MultiSetOf(_pat_let_tv23, false)).Cardinality(), _pat_let_tv24)
			}
		}(Companion_D7_.Create_DC22_(Companion_D7_.Create_DC21_(p0, _dafny.IntOfInt64(734))))
		var _arr36 _dafny.Array = _this.F7()
		_ = _arr36
		var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index167
		(_arr36).ArraySet1(p1, (_index167).Int())
		var _arr37 _dafny.Array = _this.F7()
		_ = _arr37
		var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index168
		(_arr37).ArraySet1(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("ibfacfd"), _1017_v10), (_index168).Int())
		var _1030_v13 D7
		_ = _1030_v13
		_1030_v13 = Companion_D7_.Create_DC21_(p0, _dafny.IntOfInt64(383))
		var _1031_v14 _dafny.MultiSet
		_ = _1031_v14
		_1031_v14 = _dafny.MultiSetOf(_1030_v13)
		var _1032_v15 _dafny.Sequence
		_ = _1032_v15
		_1032_v15 = _dafny.SeqOf((_this).F9(), (_1031_v14).Cardinality())
		var _1033_v16 _dafny.Map
		_ = _1033_v16
		_1033_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('j'), (false) || (!((_this).F10())))
		var _arr38 _dafny.Array = _this.F7()
		_ = _arr38
		var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))
		_ = _index169
		var _rhs198 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1032_v15, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-447))).Uint32(), func(coer51 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg51 _dafny.Int) interface{} {
				return coer51(arg51)
			}
		}((func(_1034_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
			return func(_1035_i2 _dafny.Int) _dafny.Int {
				return _1034_p0
			}
		})(p0))))
		_ = _rhs198
		var _rhs199 bool = !((func() bool {
			if (_1033_v16).Contains((_1017_v10).Select((Companion_Default___.SafeIndex((_this).Fm4(p0, (_dafny.Zero).Minus((_this).F9()), _1019_v12, _dafny.IntOfInt64(464), globalState), _dafny.IntOfUint32((_1017_v10).Cardinality()))).Uint32()).(_dafny.CodePoint)) {
				return (_1033_v16).Get((_1017_v10).Select((Companion_Default___.SafeIndex((_this).Fm4(p0, (_dafny.Zero).Minus((_this).F9()), _1019_v12, _dafny.IntOfInt64(464), globalState), _dafny.IntOfUint32((_1017_v10).Cardinality()))).Uint32()).(_dafny.CodePoint)).(bool)
			}
			return (p0).Cmp(p0) == 0
		})())
		_ = _rhs199
		var _lhs160 _dafny.Array = _this.F7()
		_ = _lhs160
		var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))
		_ = _lhs161
		_1032_v15 = _rhs198
		(_lhs160).ArraySet1(_rhs199, (_lhs161).Int())
		var _1036_v17 _dafny.Map
		_ = _1036_v17
		_1036_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.IntOfUint32((_1017_v10).Cardinality()))
		var _1037_v18 _dafny.Map
		_ = _1037_v18
		_1037_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (func() _dafny.Int {
			if (_1036_v17).Contains((_this).F9()) {
				return (_1036_v17).Get((_this).F9()).(_dafny.Int)
			}
			return (_this).F9()
		})())
		var _1038_i3 _dafny.Int
		_ = _1038_i3
		_1038_i3 = _dafny.Zero
		{
			for ((_dafny.IntOfUint32((_1032_v15).Cardinality())).Plus(_dafny.IntOfUint32((_1017_v10).Cardinality()))).Cmp(((Companion_Default___.Fm14(_1037_v18, false, _dafny.IntOfInt64(692), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), globalState)).Update(p1, Companion_Default___.Abs((_1032_v15).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F9()), _dafny.IntOfUint32((_1032_v15).Cardinality()))).Uint32()).(_dafny.Int)))).Cardinality()) >= 0 {
				{
					if (_1038_i3).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L11
					}
					_1038_i3 = (_1038_i3).Plus(_dafny.One)
					_1016_v9 = _dafny.SetOf((_this).F6(), p2)
					var _arr39 _dafny.Array = _this.F7()
					_ = _arr39
					var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))
					_ = _index170
					(_arr39).ArraySet1((_this).F6(), (_index170).Int())
					var _1039_v19 T0
					_ = _1039_v19
					var _nw163 *C2 = New_C2_()
					_ = _nw163
					_nw163.Ctor__(_dafny.MultiSetOf((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), true))
					_1039_v19 = _nw163
					var _1040_v20 _dafny.MultiSet
					_ = _1040_v20
					_1040_v20 = _dafny.MultiSetOf(_1039_v19, _1039_v19, _1039_v19)
					var _1041_v21 _dafny.Map
					_ = _1041_v21
					_1041_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _this.F8())
					var _1042_v22 _dafny.Sequence
					_ = _1042_v22
					_1042_v22 = _dafny.SeqOf(_1041_v21)
					(_this).F8_set_((_this).Fm1((_this).F6(), (func() _dafny.Int {
						if (_1040_v20).Contains(_1039_v19) {
							return (_1040_v20).Multiplicity(_1039_v19)
						}
						return _dafny.IntOfUint32((_1042_v22).Cardinality())
					})(), (_this).F9(), globalState))
					var _arr40 _dafny.Array = _this.F7()
					_ = _arr40
					var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(617), _dafny.ArrayLen((_this.F7()), 0))
					_ = _index171
					(_arr40).ArraySet1(true, (_index171).Int())
					goto C11
				}
			C11:
			}
			goto L11
		}
	L11:
	}
}
func (_this *C5) M4(p0 bool, p1 _dafny.Sequence, globalState *GlobalState) {
	{
		var _1043_v0 _dafny.Map
		_ = _1043_v0
		_1043_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F9())
		_1043_v0 = (_1043_v0).Update((_this).Fm2((_this).F10(), globalState), (_this).F9())
		var _1044_v1 _dafny.Sequence
		_ = _1044_v1
		_1044_v1 = _dafny.UnicodeSeqOfUtf8Bytes("nttgnqle")
		var _1045_v2 _dafny.Map
		_ = _1045_v2
		_1045_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1044_v1, true)
		var _1046_v3 D16
		_ = _1046_v3
		_1046_v3 = Companion_D16_.Create_DC44_((_this).Fm2((_this).F6(), globalState), (_this).Fm3(globalState), (_this).Fm4((_this).F9(), (_this).F9(), _1045_v2, (_this).F9(), globalState))
		var _source14 D6 = func(_source15 D16) D6 {
			if _source15.Is_DC44() {
				var _1047___mcc_h4 bool = _source15.Get_().(D16_DC44).Cf69
				_ = _1047___mcc_h4
				var _1048___mcc_h5 bool = _source15.Get_().(D16_DC44).Cf70
				_ = _1048___mcc_h5
				var _1049___mcc_h6 _dafny.Int = _source15.Get_().(D16_DC44).Cf71
				_ = _1049___mcc_h6
				var _1050_cf71 _dafny.Int = _1049___mcc_h6
				_ = _1050_cf71
				var _1051_cf70 bool = _1048___mcc_h5
				_ = _1051_cf70
				var _1052_cf69 bool = _1047___mcc_h4
				_ = _1052_cf69
				return Companion_D6_.Create_DC17_(_dafny.SeqOf(!(_1051_cf70), _1051_cf70, true))
			} else if _source15.Is_DC43() {
				var _1053___mcc_h7 _dafny.MultiSet = _source15.Get_().(D16_DC43).Cf68
				_ = _1053___mcc_h7
				var _1054_cf68 _dafny.MultiSet = _1053___mcc_h7
				_ = _1054_cf68
				return Companion_D6_.Create_DC17_(_dafny.SeqOf(_this.F8()))
			} else {
				var _1055___mcc_h8 D16 = _source15.Get_().(D16_DC45).Cf72
				_ = _1055___mcc_h8
				var _1056_cf72 D16 = _1055___mcc_h8
				_ = _1056_cf72
				return Companion_D6_.Create_DC17_(_dafny.SeqOf(_this.F8()))
			}
		}(_1046_v3)
		_ = _source14
		if _source14.Is_DC18() {
			var _1057___mcc_h0 *C0 = _source14.Get_().(D6_DC18).Cf25
			_ = _1057___mcc_h0
			var _1058___mcc_h1 _dafny.Sequence = _source14.Get_().(D6_DC18).Cf26
			_ = _1058___mcc_h1
			var _1059___mcc_h2 _dafny.Int = _source14.Get_().(D6_DC18).Cf27
			_ = _1059___mcc_h2
			var _1060_cf27 _dafny.Int = _1059___mcc_h2
			_ = _1060_cf27
			var _1061_cf26 _dafny.Sequence = _1058___mcc_h1
			_ = _1061_cf26
			var _1062_cf25 *C0 = _1057___mcc_h0
			_ = _1062_cf25
			var _1063_v4 _dafny.Array
			_ = _1063_v4
			var _nw164 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw164
			_1063_v4 = _nw164
			var _1064_v5 _dafny.Map
			_ = _1064_v5
			_1064_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1063_v4)
			_1064_v5 = (_1064_v5).Update(false, _1063_v4)
			var _1065_v6 _dafny.CodePoint
			_ = _1065_v6
			_1065_v6 = _dafny.CodePoint('l')
			var _1066_v7 _dafny.Array
			_ = _1066_v7
			var _nwElement0_23 _dafny.CodePoint = _1065_v6
			_ = _nwElement0_23
			var _nw165 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(8))
			_ = _nw165
			(_nw165).ArraySet1CodePoint(_nwElement0_23, 0)
			(_nw165).ArraySet1CodePoint(_1065_v6, 1)
			(_nw165).ArraySet1CodePoint(_1065_v6, 2)
			(_nw165).ArraySet1CodePoint(_1065_v6, 3)
			(_nw165).ArraySet1CodePoint(_1065_v6, 4)
			(_nw165).ArraySet1CodePoint(_1065_v6, 5)
			(_nw165).ArraySet1CodePoint(_1065_v6, 6)
			(_nw165).ArraySet1CodePoint(_1065_v6, 7)
			_1066_v7 = _nw165
			var _1067_v8 _dafny.Array
			_ = _1067_v8
			var _nwElement0_24 _dafny.Array = _1066_v7
			_ = _nwElement0_24
			var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(13))
			_ = _nw166
			(_nw166).ArraySet1(_nwElement0_24, 0)
			(_nw166).ArraySet1(_1066_v7, 1)
			(_nw166).ArraySet1(_1066_v7, 2)
			(_nw166).ArraySet1(_1066_v7, 3)
			(_nw166).ArraySet1(_1066_v7, 4)
			(_nw166).ArraySet1(_1066_v7, 5)
			(_nw166).ArraySet1(_1066_v7, 6)
			(_nw166).ArraySet1(_1066_v7, 7)
			(_nw166).ArraySet1(_1066_v7, 8)
			(_nw166).ArraySet1(_1066_v7, 9)
			(_nw166).ArraySet1(_1066_v7, 10)
			(_nw166).ArraySet1(_1066_v7, 11)
			(_nw166).ArraySet1(_1066_v7, 12)
			_1067_v8 = _nw166
			var _1068_v9 D4
			_ = _1068_v9
			_1068_v9 = Companion_D4_.Create_DC12_(_1067_v8)
			var _1069_v10 D4
			_ = _1069_v10
			_1069_v10 = Companion_D4_.Create_DC14_(_1068_v9)
			var _1070_v11 D4
			_ = _1070_v11
			_1070_v11 = Companion_D4_.Create_DC14_(_1069_v10)
			var _pat_let_tv25 = _1069_v10
			_ = _pat_let_tv25
			var _pat_let_tv26 = _1069_v10
			_ = _pat_let_tv26
			var _1071_v12 _dafny.Array
			_ = _1071_v12
			var _nwElement0_25 D4 = _1070_v11
			_ = _nwElement0_25
			var _nw167 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(20))
			_ = _nw167
			(_nw167).ArraySet1(_nwElement0_25, 0)
			(_nw167).ArraySet1(_1070_v11, 1)
			(_nw167).ArraySet1(_1070_v11, 2)
			(_nw167).ArraySet1(_1070_v11, 3)
			(_nw167).ArraySet1(_1070_v11, 4)
			(_nw167).ArraySet1(_1070_v11, 5)
			(_nw167).ArraySet1(_1070_v11, 6)
			(_nw167).ArraySet1(_1070_v11, 7)
			(_nw167).ArraySet1(_1070_v11, 8)
			(_nw167).ArraySet1(Companion_D4_.Create_DC14_(_1069_v10), 9)
			(_nw167).ArraySet1(_1070_v11, 10)
			(_nw167).ArraySet1(_1070_v11, 11)
			(_nw167).ArraySet1(_1070_v11, 12)
			(_nw167).ArraySet1(_1070_v11, 13)
			(_nw167).ArraySet1(Companion_D4_.Create_DC14_((_1070_v11).Dtor_cf17()), 14)
			(_nw167).ArraySet1(func(_pat_let30_0 D4) D4 {
				return func(_1072_dt__update__tmp_h0 D4) D4 {
					return func(_pat_let31_0 D4) D4 {
						return func(_1073_dt__update_hcf17_h0 D4) D4 {
							return Companion_D4_.Create_DC14_(_1073_dt__update_hcf17_h0)
						}(_pat_let31_0)
					}(_pat_let_tv25)
				}(_pat_let30_0)
			}(Companion_D4_.Create_DC14_(_1069_v10)), 15)
			(_nw167).ArraySet1(_1070_v11, 16)
			(_nw167).ArraySet1(_1070_v11, 17)
			(_nw167).ArraySet1(_1070_v11, 18)
			(_nw167).ArraySet1(func(_pat_let32_0 D4) D4 {
				return func(_1074_dt__update__tmp_h1 D4) D4 {
					return func(_pat_let33_0 D4) D4 {
						return func(_1075_dt__update_hcf17_h1 D4) D4 {
							return Companion_D4_.Create_DC14_(_1075_dt__update_hcf17_h1)
						}(_pat_let33_0)
					}(_pat_let_tv26)
				}(_pat_let32_0)
			}(_1070_v11), 19)
			_1071_v12 = _nw167
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_1071_v12), 0))
			_ = _index172
			(_1071_v12).ArraySet1(_1070_v11, (_index172).Int())
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(539), _dafny.ArrayLen((_1071_v12), 0))
			_ = _index173
			(_1071_v12).ArraySet1(_1070_v11, (_index173).Int())
			var _1076_v13 *C3
			_ = _1076_v13
			var _nw168 *C3 = New_C3_()
			_ = _nw168
			_nw168.Ctor__((_dafny.IntOfInt64(984)).Cmp((_this).F9()) > 0, (_this).F5())
			_1076_v13 = _nw168
			var _1077_v14 *C3
			_ = _1077_v14
			var _nw169 *C3 = New_C3_()
			_ = _nw169
			_nw169.Ctor__((_this).F6(), (_this).F5())
			_1077_v14 = _nw169
		} else {
			var _1078___mcc_h3 _dafny.Sequence = _source14.Get_().(D6_DC17).Cf24
			_ = _1078___mcc_h3
			var _1079_cf24 _dafny.Sequence = _1078___mcc_h3
			_ = _1079_cf24
			var _1080_v15 _dafny.Map
			_ = _1080_v15
			_1080_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _this.F8())
			var _1081_v16 _dafny.Array
			_ = _1081_v16
			var _nwElement0_26 _dafny.Int = (_1080_v15).Cardinality()
			_ = _nwElement0_26
			var _nw170 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(14))
			_ = _nw170
			(_nw170).ArraySet1(_nwElement0_26, 0)
			(_nw170).ArraySet1(((_this).F9()).Times((_this).F9()), 1)
			(_nw170).ArraySet1((_this).F9(), 2)
			(_nw170).ArraySet1((_this).F9(), 3)
			(_nw170).ArraySet1((_this).F9(), 4)
			(_nw170).ArraySet1((_this).F9(), 5)
			(_nw170).ArraySet1(_dafny.IntOfUint32((p1).Cardinality()), 6)
			(_nw170).ArraySet1((_dafny.Zero).Minus(((_this).F9()).Times((_dafny.Zero).Minus((_this).F9()))), 7)
			(_nw170).ArraySet1((_this).F9(), 8)
			(_nw170).ArraySet1((_this).F9(), 9)
			(_nw170).ArraySet1((_this).F9(), 10)
			(_nw170).ArraySet1((_this).F9(), 11)
			(_nw170).ArraySet1(_dafny.IntOfInt64(635), 12)
			(_nw170).ArraySet1((_dafny.Zero).Minus((_this).F9()), 13)
			_1081_v16 = _nw170
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
			_ = _index174
			(_1081_v16).ArraySet1(_dafny.IntOfInt64(600), (_index174).Int())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
			_ = _index175
			(_1081_v16).ArraySet1((_this).F9(), (_index175).Int())
			var _1082_v17 _dafny.Set
			_ = _1082_v17
			_1082_v17 = _dafny.SetOf((_this).F10())
			var _1083_v18 D0
			_ = _1083_v18
			_1083_v18 = Companion_D0_.Create_DC0_((_1082_v17).Cardinality())
			var _1084_v19 _dafny.Map
			_ = _1084_v19
			_1084_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1083_v18, (_1044_v1).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1044_v1).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _1085_v20 _dafny.CodePoint
			_ = _1085_v20
			_1085_v20 = _dafny.CodePoint('k')
			_1084_v19 = (_1084_v19).Update((func() D0 {
				if _this.F8() {
					return Companion_D0_.Create_DC0_((_this).F9())
				}
				return _1083_v18
			})(), _1085_v20)
			var _1086_v21 _dafny.Map
			_ = _1086_v21
			_1086_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), !(p0))
			if !((func() bool {
				if (func() bool {
					if (_1086_v21).Contains((_this).F9()) {
						return (_1086_v21).Get((_this).F9()).(bool)
					}
					return _this.F8()
				})() {
					return true
				}
				return _this.F8()
			})()) {
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
				_ = _index176
				(_1081_v16).ArraySet1((_this).F9(), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
				_ = _index177
				var _rhs200 bool = (true) && ((_this).F10())
				_ = _rhs200
				var _rhs201 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1044_v1, _1044_v1), _dafny.Companion_Sequence_.Concatenate(_1044_v1, _1044_v1))
				_ = _rhs201
				var _rhs202 bool = (_this).F6()
				_ = _rhs202
				var _rhs203 bool = (_this).F6()
				_ = _rhs203
				var _rhs204 _dafny.Int = (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int)
				_ = _rhs204
				var _lhs162 *C5 = _this
				_ = _lhs162
				var _lhs163 *C5 = _this
				_ = _lhs163
				var _lhs164 *C5 = _this
				_ = _lhs164
				var _lhs165 _dafny.Array = _1081_v16
				_ = _lhs165
				var _lhs166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
				_ = _lhs166
				_lhs162.F8_set_(_rhs200)
				_1044_v1 = _rhs201
				_lhs163.F8_set_(_rhs202)
				_lhs164.F8_set_(_rhs203)
				(_lhs165).ArraySet1(_rhs204, (_lhs166).Int())
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
				_ = _index178
				var _rhs205 _dafny.Sequence = _1044_v1
				_ = _rhs205
				var _rhs206 _dafny.Int = (_this).Fm4((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _1045_v2, (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), globalState)
				_ = _rhs206
				var _rhs207 _dafny.Int = (func() _dafny.Int {
					if (_1079_cf24).Select((Companion_Default___.SafeIndex((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1079_cf24).Cardinality()))).Uint32()).(bool) {
						return Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(p1, (Companion_Default___.SafeIndex((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((p1).Cardinality()))).Uint32(), _dafny.IntOfInt64(-872))).Cardinality()), (_this).Fm4((_1080_v15).Cardinality(), _dafny.IntOfUint32((_1044_v1).Cardinality()), _1045_v2, _dafny.IntOfUint32((_1079_cf24).Cardinality()), globalState))
					}
					return (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int)
				})()
				_ = _rhs207
				var _rhs208 bool = (!(_this.F8()) || ((_this).F6())) || (true)
				_ = _rhs208
				var _rhs209 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(452))).Uint32(), func(coer52 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}((func(_1087_v16 _dafny.Array, _1088_v0 _dafny.Map) func(_dafny.Int) _dafny.Set {
					return func(_1089_i0 _dafny.Int) _dafny.Set {
						return (_dafny.SetOf((_1087_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1087_v16), 0))).Int()).(_dafny.Int))).Union(func() _dafny.Set {
							var _coll43 = _dafny.NewBuilder()
							_ = _coll43
							for _iter47 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(314), _dafny.IntOfInt64(504))); ; {
								_compr_43, _ok47 := _iter47()
								if !_ok47 {
									break
								}
								var _1090_v22 _dafny.Int
								_1090_v22 = interface{}(_compr_43).(_dafny.Int)
								if ((_dafny.IntOfInt64(314)).Cmp(_1090_v22) <= 0) && ((_1090_v22).Cmp(_dafny.IntOfInt64(504)) < 0) {
									_coll43.Add((_1090_v22).Times((func() _dafny.Int {
										if (_1088_v0).Contains(false) {
											return (_1088_v0).Get(false).(_dafny.Int)
										}
										return (_1087_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1087_v16), 0))).Int()).(_dafny.Int)
									})()))
								}
							}
							return _coll43.ToSet()
						}())
					}
				})(_1081_v16, _1043_v0)))).Cardinality())
				_ = _rhs209
				var _lhs167 *GlobalState = globalState
				_ = _lhs167
				var _lhs168 *GlobalState = globalState
				_ = _lhs168
				var _lhs169 *C5 = _this
				_ = _lhs169
				var _lhs170 _dafny.Array = _1081_v16
				_ = _lhs170
				var _lhs171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
				_ = _lhs171
				_1044_v1 = _rhs205
				_lhs167.F0 = _rhs206
				_lhs168.F0 = _rhs207
				_lhs169.F8_set_(_rhs208)
				(_lhs170).ArraySet1(_rhs209, (_lhs171).Int())
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))
				_ = _index179
				(_1081_v16).ArraySet1((func() _dafny.Int {
					if p0 {
						return (_this).F9()
					}
					return (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int)
				})(), (_index179).Int())
				var _1091_v23 *C2
				_ = _1091_v23
				var _nw171 *C2 = New_C2_()
				_ = _nw171
				_nw171.Ctor__((_this).F5())
				_1091_v23 = _nw171
				var _1092_v24 _dafny.Set
				_ = _1092_v24
				_1092_v24 = _dafny.SetOf((_dafny.Zero).Minus((_this).F9()))
				var _rhs210 *C2 = _1091_v23
				_ = _rhs210
				var _rhs211 _dafny.Int = Companion_Default___.SafeDivisionInt(((_this).F9()).Times((_this).Fm4((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _1045_v2, (_1092_v24).Cardinality(), globalState)), (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int))
				_ = _rhs211
				var _lhs172 *GlobalState = globalState
				_ = _lhs172
				_1091_v23 = _rhs210
				_lhs172.F0 = _rhs211
			} else {
				var _1093_v25 _dafny.Sequence
				_ = _1093_v25
				_1093_v25 = _dafny.SeqOf((_1080_v15).Update(_this.F8(), _this.F8()))
				(_this).F8_set_(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(652))).Uint32(), func(coer53 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg53 _dafny.Int) interface{} {
						return coer53(arg53)
					}
				}((func(_1094_v15 _dafny.Map) func(_dafny.Int) _dafny.Map {
					return func(_1095_i1 _dafny.Int) _dafny.Map {
						return _1094_v15
					}
				})(_1080_v15))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(110))).Uint32(), func(coer54 func(_dafny.Int) _dafny.Map) func(_dafny.Int) interface{} {
					return func(arg54 _dafny.Int) interface{} {
						return coer54(arg54)
					}
				}((func(_1096_cf24 _dafny.Sequence, _1097_v16 _dafny.Array) func(_dafny.Int) _dafny.Map {
					return func(_1098_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), (_1096_cf24).Select((Companion_Default___.SafeIndex((_1097_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1097_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1096_cf24).Cardinality()))).Uint32()).(bool))
					}
				})(_1079_cf24, _1081_v16)))), _1093_v25))
				(_this).F8_set_((_this).F6())
				var _arr41 _dafny.Array = _this.F7()
				_ = _arr41
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index180
				(_arr41).ArraySet1((_this).F6(), (_index180).Int())
				var _arr42 _dafny.Array = _this.F7()
				_ = _arr42
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index181
				(_arr42).ArraySet1(!(((_1082_v17).Cardinality()).Cmp(_dafny.IntOfUint32((_1079_cf24).Cardinality())) != 0), (_index181).Int())
				var _arr43 _dafny.Array = _this.F7()
				_ = _arr43
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index182
				(_arr43).ArraySet1((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool), (_index182).Int())
				var _1099_v26 *C0
				_ = _1099_v26
				var _nw172 *C0 = New_C0_()
				_ = _nw172
				_nw172.Ctor__(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_1079_cf24, (Companion_Default___.SafeIndex((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1079_cf24).Cardinality()))).Uint32(), (_this).F10()), (Companion_Default___.SafeIndex((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_1079_cf24, (Companion_Default___.SafeIndex((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1079_cf24).Cardinality()))).Uint32(), (_this).F10())).Cardinality()))).Uint32(), p0)).Cardinality()), (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int)), (_this).Fm2((_this).F10(), globalState)))
				_1099_v26 = _nw172
			}
			var _1100_v27 _dafny.MultiSet
			_ = _1100_v27
			_1100_v27 = _dafny.MultiSetOf((_dafny.Zero).Minus((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int)))
			var _1101_v28 _dafny.Set
			_ = _1101_v28
			_1101_v28 = _dafny.SetOf(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-541), (_this).Fm4((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), (_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int), _1045_v2, (_dafny.Zero).Minus((_1081_v16).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(232), _dafny.ArrayLen((_1081_v16), 0))).Int()).(_dafny.Int)), globalState)))
			var _rhs212 bool = (_1100_v27).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(713))).Uint32(), func(coer55 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg55 _dafny.Int) interface{} {
					return coer55(arg55)
				}
			}(func(_1102_i3 _dafny.Int) _dafny.Int {
				return (_this).F9()
			}))))
			_ = _rhs212
			var _rhs213 bool = false
			_ = _rhs213
			var _rhs214 _dafny.Int = (_1101_v28).Cardinality()
			_ = _rhs214
			var _lhs173 *C5 = _this
			_ = _lhs173
			var _lhs174 *C5 = _this
			_ = _lhs174
			var _lhs175 *GlobalState = globalState
			_ = _lhs175
			_lhs173.F8_set_(_rhs212)
			_lhs174.F8_set_(_rhs213)
			_lhs175.F0 = _rhs214
		}
		if (_this).F6() {
			var _1103_v29 _dafny.CodePoint
			_ = _1103_v29
			_1103_v29 = _dafny.CodePoint('e')
			var _1104_v30 _dafny.Map
			_ = _1104_v30
			_1104_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
			var _1105_v31 D2
			_ = _1105_v31
			_1105_v31 = Companion_D2_.Create_DC7_(_1103_v29, _1104_v30, (_this).F9())
			var _1106_v32 D2
			_ = _1106_v32
			_1106_v32 = Companion_D2_.Create_DC8_(_1105_v31)
			var _source16 D2 = _1106_v32
			_ = _source16
			if _source16.Is_DC5() {
				var _1107___mcc_h9 bool = _source16.Get_().(D2_DC5).Cf5
				_ = _1107___mcc_h9
				var _1108_cf5 bool = _1107___mcc_h9
				_ = _1108_cf5
				var _1109_v33 _dafny.Array
				_ = _1109_v33
				var _nw173 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
				_ = _nw173
				_1109_v33 = _nw173
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1109_v33), 0))
				_ = _index183
				(_1109_v33).ArraySet1((_this).F9(), (_index183).Int())
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_1109_v33), 0))
				_ = _index184
				(_1109_v33).ArraySet1(((_this).F9()).Plus((_this).F9()), (_index184).Int())
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1109_v33), 0))
				_ = _index185
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_1109_v33), 0))
				_ = _index186
				var _rhs215 _dafny.Int = (_this).F9()
				_ = _rhs215
				var _rhs216 _dafny.Int = ((_this).F9()).Plus(_dafny.IntOfInt64(721))
				_ = _rhs216
				var _rhs217 _dafny.Int = (_dafny.IntOfUint32((_1044_v1).Cardinality())).Minus(_dafny.IntOfInt64(-353))
				_ = _rhs217
				var _rhs218 bool = !((func() bool {
					if (_this).F10() {
						return p0
					}
					return (_this).F10()
				})())
				_ = _rhs218
				var _lhs176 _dafny.Array = _1109_v33
				_ = _lhs176
				var _lhs177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1109_v33), 0))
				_ = _lhs177
				var _lhs178 _dafny.Array = _1109_v33
				_ = _lhs178
				var _lhs179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(837), _dafny.ArrayLen((_1109_v33), 0))
				_ = _lhs179
				var _lhs180 *GlobalState = globalState
				_ = _lhs180
				var _lhs181 *C5 = _this
				_ = _lhs181
				(_lhs176).ArraySet1(_rhs215, (_lhs177).Int())
				(_lhs178).ArraySet1(_rhs216, (_lhs179).Int())
				_lhs180.F0 = _rhs217
				_lhs181.F8_set_(_rhs218)
				(globalState).F0 = (_this).F9()
				var _1110_v34 _dafny.Map
				_ = _1110_v34
				_1110_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10())
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1109_v33), 0))
				_ = _index187
				(_1109_v33).ArraySet1((((_1110_v34).Cardinality()).Minus(_dafny.IntOfInt64(301))).Times((_1109_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1109_v33), 0))).Int()).(_dafny.Int)), (_index187).Int())
				var _1111_v35 _dafny.Map
				_ = _1111_v35
				_1111_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _1103_v29)
				var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(690), _dafny.ArrayLen((_1109_v33), 0))
				_ = _index188
				(_1109_v33).ArraySet1((_dafny.Zero).Minus(((_1111_v35).Cardinality()).Plus((_this).F9())), (_index188).Int())
			} else if _source16.Is_DC6() {
				var _1112_v36 T0
				_ = _1112_v36
				var _nw174 *C2 = New_C2_()
				_ = _nw174
				_nw174.Ctor__((_this).F5())
				_1112_v36 = _nw174
				var _1113_v37 _dafny.Map
				_ = _1113_v37
				_1113_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1112_v36, !(true))
				_1113_v37 = (_1113_v37).Merge(_1113_v37)
				var _1114_v40 _dafny.Map
				_ = _1114_v40
				_1114_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1104_v30, p0)
				var _1115_v41 *C0
				_ = _1115_v41
				var _nw175 *C0 = New_C0_()
				_ = _nw175
				_nw175.Ctor__((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(func() _dafny.Map {
					var _coll44 = _dafny.NewMapBuilder()
					_ = _coll44
					for _iter48 := _dafny.Iterate((func() _dafny.Map {
						var _coll45 = _dafny.NewMapBuilder()
						_ = _coll45
						for _iter49 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(50), _dafny.IntOfInt64(-133))); ; {
							_compr_45, _ok49 := _iter49()
							if !_ok49 {
								break
							}
							var _1116_v39 _dafny.Int
							_1116_v39 = interface{}(_compr_45).(_dafny.Int)
							if ((_dafny.IntOfInt64(50)).Cmp(_1116_v39) <= 0) && ((_1116_v39).Cmp(_dafny.IntOfInt64(-133)) < 0) {
								_coll45.Add(Companion_Default___.SafeDivisionInt(_1116_v39, (_this).F9()), _this.F8())
							}
						}
						return _coll45.ToMap()
					}()).Keys().Elements()); ; {
						_compr_44, _ok48 := _iter48()
						if !_ok48 {
							break
						}
						var _1117_v38 _dafny.Int
						_1117_v38 = interface{}(_compr_44).(_dafny.Int)
						if (func() _dafny.Map {
							var _coll46 = _dafny.NewMapBuilder()
							_ = _coll46
							for _iter50 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(50), _dafny.IntOfInt64(-133))); ; {
								_compr_46, _ok50 := _iter50()
								if !_ok50 {
									break
								}
								var _1116_v39 _dafny.Int
								_1116_v39 = interface{}(_compr_46).(_dafny.Int)
								if ((_dafny.IntOfInt64(50)).Cmp(_1116_v39) <= 0) && ((_1116_v39).Cmp(_dafny.IntOfInt64(-133)) < 0) {
									_coll46.Add(Companion_Default___.SafeDivisionInt(_1116_v39, (_this).F9()), _this.F8())
								}
							}
							return _coll46.ToMap()
						}()).Contains(_1117_v38) {
							_coll44.Add((_1117_v38).Plus((_this).F9()), _dafny.IntOfInt64(175))
						}
					}
					return _coll44.ToMap()
				}(), _this.F8())).Merge(_1114_v40))
				_1115_v41 = _nw175
				var _1118_v42 _dafny.Sequence
				_ = _1118_v42
				_1118_v42 = _dafny.SeqOf(_dafny.IntOfUint32((p1).Cardinality()), (_this).F9())
				var _1119_v43 _dafny.Array
				_ = _1119_v43
				var _nw176 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw176
				_1119_v43 = _nw176
				var _1120_v44 _dafny.Sequence
				_ = _1120_v44
				_1120_v44 = _dafny.SeqOf(p1, _1118_v42, (func() _dafny.Sequence {
					if p0 {
						return _1118_v42
					}
					return _dafny.Companion_Sequence_.Update(_1118_v42, (Companion_Default___.SafeIndex((_1104_v30).Cardinality(), _dafny.IntOfUint32((_1118_v42).Cardinality()))).Uint32(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm4(_dafny.IntOfUint32((_1044_v1).Cardinality()), (_this).F9(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1044_v1, (_this).F6()), (_this).F9(), globalState))).Cardinality())
				})(), p1, p1)
				var _1121_v45 D3
				_ = _1121_v45
				_1121_v45 = Companion_D3_.Create_DC11_((_this).F9())
				var _1122_v46 _dafny.Set
				_ = _1122_v46
				_1122_v46 = _dafny.SetOf(_1121_v45)
				var _1123_v47 _dafny.MultiSet
				_ = _1123_v47
				_1123_v47 = _dafny.MultiSetOf((_this).F9())
				var _rhs219 _dafny.Sequence = (_1120_v44).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1120_v44).Cardinality()))).Uint32()).(_dafny.Sequence)
				_ = _rhs219
				var _rhs220 _dafny.Array = _1119_v43
				_ = _rhs220
				var _rhs221 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(Companion_D3_.Create_DC11_((_this).F9()), _1121_v45)).IsSubsetOf(_1122_v46), p0)).Cardinality()
				_ = _rhs221
				var _rhs222 _dafny.Int = Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F9()), (_1123_v47).Cardinality())
				_ = _rhs222
				var _lhs182 *GlobalState = globalState
				_ = _lhs182
				var _lhs183 *GlobalState = globalState
				_ = _lhs183
				_1118_v42 = _rhs219
				_1119_v43 = _rhs220
				_lhs182.F0 = _rhs221
				_lhs183.F0 = _rhs222
				var _1124_v48 _dafny.Sequence
				_ = _1124_v48
				_1124_v48 = _dafny.SeqOf(_1046_v3, _1046_v3)
				var _1125_v49 *C4
				_ = _1125_v49
				var _nw177 *C4 = New_C4_()
				_ = _nw177
				_nw177.Ctor__(_this.F7(), _this.F8(), p0, (_dafny.MultiSetOf(p0)).Update(_this.F8(), Companion_Default___.Abs(_dafny.IntOfUint32((_1124_v48).Cardinality()))))
				_1125_v49 = _nw177
				var _1126_v50 _dafny.Set
				_ = _1126_v50
				_1126_v50 = _dafny.SetOf(_1125_v49)
				var _1127_v51 _dafny.Sequence
				_ = _1127_v51
				_1127_v51 = _dafny.SeqOf((_1126_v50).IsSubsetOf(_dafny.SetOf(_1125_v49, _1125_v49)))
				var _1128_v52 _dafny.Map
				_ = _1128_v52
				_1128_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1103_v29, (_this).F9())
				(_this).F8_set_((_1127_v51).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if (_1128_v52).Contains(_1103_v29) {
						return (_1128_v52).Get(_1103_v29).(_dafny.Int)
					}
					return (_this).F9()
				})(), _dafny.IntOfUint32((_1127_v51).Cardinality()))).Uint32()).(bool))
			} else if _source16.Is_DC7() {
				var _1129___mcc_h10 _dafny.CodePoint = _source16.Get_().(D2_DC7).Cf6
				_ = _1129___mcc_h10
				var _1130___mcc_h11 _dafny.Map = _source16.Get_().(D2_DC7).Cf7
				_ = _1130___mcc_h11
				var _1131___mcc_h12 _dafny.Int = _source16.Get_().(D2_DC7).Cf8
				_ = _1131___mcc_h12
				var _1132_cf8 _dafny.Int = _1131___mcc_h12
				_ = _1132_cf8
				var _1133_cf7 _dafny.Map = _1130___mcc_h11
				_ = _1133_cf7
				var _1134_cf6 _dafny.CodePoint = _1129___mcc_h10
				_ = _1134_cf6
				var _1135_v53 _dafny.Array
				_ = _1135_v53
				var _len0_28 _dafny.Int = _dafny.IntOfInt64(18)
				_ = _len0_28
				var _nw178 _dafny.Array
				_ = _nw178
				if _len0_28.Cmp(_dafny.Zero) == 0 {
					_nw178 = _dafny.NewArray(_len0_28)
				} else {
					var _init28 func(_dafny.Int) _dafny.Map = (func(_1136_p0 bool) func(_dafny.Int) _dafny.Map {
						return func(_1137_i4 _dafny.Int) _dafny.Map {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1136_p0, _1136_p0)
						}
					})(p0)
					_ = _init28
					var _element0_28 = _init28(_dafny.Zero)
					_ = _element0_28
					_nw178 = _dafny.NewArrayFromExample(_element0_28, nil, _len0_28)
					(_nw178).ArraySet1(_element0_28, 0)
					var _nativeLen0_28 = (_len0_28).Int()
					_ = _nativeLen0_28
					for _i0_28 := 1; _i0_28 < _nativeLen0_28; _i0_28++ {
						(_nw178).ArraySet1(_init28(_dafny.IntOf(_i0_28)), _i0_28)
					}
				}
				_1135_v53 = _nw178
				var _len0_29 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_29
				var _nw179 _dafny.Array
				_ = _nw179
				if _len0_29.Cmp(_dafny.Zero) == 0 {
					_nw179 = _dafny.NewArray(_len0_29)
				} else {
					var _init29 func(_dafny.Int) _dafny.Map = func(_1138_i5 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F8(), (_this).F6())
					}
					_ = _init29
					var _element0_29 = _init29(_dafny.Zero)
					_ = _element0_29
					_nw179 = _dafny.NewArrayFromExample(_element0_29, nil, _len0_29)
					(_nw179).ArraySet1(_element0_29, 0)
					var _nativeLen0_29 = (_len0_29).Int()
					_ = _nativeLen0_29
					for _i0_29 := 1; _i0_29 < _nativeLen0_29; _i0_29++ {
						(_nw179).ArraySet1(_init29(_dafny.IntOf(_i0_29)), _i0_29)
					}
				}
				_1135_v53 = _nw179
				(_this).F8_set_(_this.F8())
				(_this).F8_set_((_this).F10())
				var _rhs223 _dafny.Int = ((_this).F9()).Minus(_1132_cf8)
				_ = _rhs223
				var _rhs224 _dafny.Int = (_this).F9()
				_ = _rhs224
				var _rhs225 _dafny.Int = ((func() _dafny.Int {
					if (_1043_v0).Contains((_this).F6()) {
						return (_1043_v0).Get((_this).F6()).(_dafny.Int)
					}
					return (_this).F9()
				})()).Plus((_this).F9())
				_ = _rhs225
				var _lhs184 *GlobalState = globalState
				_ = _lhs184
				var _lhs185 *GlobalState = globalState
				_ = _lhs185
				_1132_cf8 = _rhs223
				_lhs184.F0 = _rhs224
				_lhs185.F0 = _rhs225
			} else if _source16.Is_DC4() {
				var _1139___mcc_h13 _dafny.Set = _source16.Get_().(D2_DC4).Cf4
				_ = _1139___mcc_h13
				var _1140_cf4 _dafny.Set = _1139___mcc_h13
				_ = _1140_cf4
				_1103_v29 = _1103_v29
				(globalState).F0 = (_this).F9()
				var _1141_v54 _dafny.Map
				_ = _1141_v54
				_1141_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1104_v30, _this.F8())
				var _1142_v55 *C0
				_ = _1142_v55
				var _nw180 *C0 = New_C0_()
				_ = _nw180
				_nw180.Ctor__(_1141_v54)
				_1142_v55 = _nw180
				var _1143_v56 _dafny.Array
				_ = _1143_v56
				var _len0_30 _dafny.Int = _dafny.IntOfInt64(5)
				_ = _len0_30
				var _nw181 _dafny.Array
				_ = _nw181
				if _len0_30.Cmp(_dafny.Zero) == 0 {
					_nw181 = _dafny.NewArray(_len0_30)
				} else {
					var _init30 func(_dafny.Int) _dafny.Sequence = (func(_1144_p0 bool, _1145_v1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_1146_i6 _dafny.Int) _dafny.Sequence {
							return (func() _dafny.Sequence {
								if _1144_p0 {
									return _1145_v1
								}
								return _dafny.UnicodeSeqOfUtf8Bytes("esgdfxat")
							})()
						}
					})(p0, _1044_v1)
					_ = _init30
					var _element0_30 = _init30(_dafny.Zero)
					_ = _element0_30
					_nw181 = _dafny.NewArrayFromExample(_element0_30, nil, _len0_30)
					(_nw181).ArraySet1(_element0_30, 0)
					var _nativeLen0_30 = (_len0_30).Int()
					_ = _nativeLen0_30
					for _i0_30 := 1; _i0_30 < _nativeLen0_30; _i0_30++ {
						(_nw181).ArraySet1(_init30(_dafny.IntOf(_i0_30)), _i0_30)
					}
				}
				_1143_v56 = _nw181
				var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1143_v56), 0))
				_ = _index189
				(_1143_v56).ArraySet1(_1044_v1, (_index189).Int())
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(733), _dafny.ArrayLen((_1143_v56), 0))
				_ = _index190
				(_1143_v56).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1044_v1, _1044_v1), (_index190).Int())
			} else {
				var _1147___mcc_h14 D2 = _source16.Get_().(D2_DC8).Cf9
				_ = _1147___mcc_h14
				var _1148_cf9 D2 = _1147___mcc_h14
				_ = _1148_cf9
				(globalState).F0 = ((_this).F9()).Plus(_dafny.IntOfUint32((p1).Cardinality()))
				var _1149_v57 D5
				_ = _1149_v57
				_1149_v57 = Companion_D5_.Create_DC15_(_1044_v1)
				var _1150_v58 _dafny.Map
				_ = _1150_v58
				_1150_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1149_v57, _1044_v1)
				var _1151_v59 *C1
				_ = _1151_v59
				var _nw182 *C1 = New_C1_()
				_ = _nw182
				_nw182.Ctor__(p0, (func() _dafny.Sequence {
					if (_1150_v58).Contains(_1149_v57) {
						return (_1150_v58).Get(_1149_v57).(_dafny.Sequence)
					}
					return _1044_v1
				})(), (_this).F6(), (_dafny.MultiSetOf((_this).F10())).Union((_this).F5()))
				_1151_v59 = _nw182
				var _1152_v60 _dafny.Sequence
				_ = _1152_v60
				_1152_v60 = _dafny.SeqOf((_this).F10(), false)
				(_this).F8_set_((_1152_v60).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((_this).F9()).Minus((_this).F9())), _dafny.IntOfUint32((_1152_v60).Cardinality()))).Uint32()).(bool))
				(globalState).F0 = (_dafny.Zero).Minus((_this).F9())
			}
			var _1153_v61 _dafny.Set
			_ = _1153_v61
			_1153_v61 = _dafny.SetOf(_1046_v3)
			var _arr44 _dafny.Array = _this.F7()
			_ = _arr44
			var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index191
			(_arr44).ArraySet1(p0, (_index191).Int())
			var _arr45 _dafny.Array = _this.F7()
			_ = _arr45
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index192
			var _rhs226 _dafny.Set = _1153_v61
			_ = _rhs226
			var _rhs227 _dafny.Int = (_this).F9()
			_ = _rhs227
			var _rhs228 bool = _dafny.Companion_Sequence_.IsPrefixOf(_dafny.UnicodeSeqOfUtf8Bytes("uqurwber"), Companion_Default___.Fm17(globalState))
			_ = _rhs228
			var _rhs229 bool = !(_this.F8())
			_ = _rhs229
			var _rhs230 _dafny.Int = (_this).F9()
			_ = _rhs230
			var _lhs186 *GlobalState = globalState
			_ = _lhs186
			var _lhs187 _dafny.Array = _this.F7()
			_ = _lhs187
			var _lhs188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
			_ = _lhs188
			var _lhs189 *C5 = _this
			_ = _lhs189
			var _lhs190 *GlobalState = globalState
			_ = _lhs190
			_1153_v61 = _rhs226
			_lhs186.F0 = _rhs227
			(_lhs187).ArraySet1(_rhs228, (_lhs188).Int())
			_lhs189.F8_set_(_rhs229)
			_lhs190.F0 = _rhs230
			if (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool) {
				_1044_v1 = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ism"), _1044_v1)
				var _1154_v62 _dafny.Array
				_ = _1154_v62
				var _len0_31 _dafny.Int = _dafny.IntOfInt64(27)
				_ = _len0_31
				var _nw183 _dafny.Array
				_ = _nw183
				if _len0_31.Cmp(_dafny.Zero) == 0 {
					_nw183 = _dafny.NewArray(_len0_31)
				} else {
					var _init31 func(_dafny.Int) D1 = func(_1155_i7 _dafny.Int) D1 {
						return Companion_D1_.Create_DC2_((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))
					}
					_ = _init31
					var _element0_31 = _init31(_dafny.Zero)
					_ = _element0_31
					_nw183 = _dafny.NewArrayFromExample(_element0_31, nil, _len0_31)
					(_nw183).ArraySet1(_element0_31, 0)
					var _nativeLen0_31 = (_len0_31).Int()
					_ = _nativeLen0_31
					for _i0_31 := 1; _i0_31 < _nativeLen0_31; _i0_31++ {
						(_nw183).ArraySet1(_init31(_dafny.IntOf(_i0_31)), _i0_31)
					}
				}
				_1154_v62 = _nw183
				var _1156_v63 D1
				_ = _1156_v63
				_1156_v63 = Companion_D1_.Create_DC2_((_this).F10())
				var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_1154_v62), 0))
				_ = _index193
				(_1154_v62).ArraySet1(_1156_v63, (_index193).Int())
				var _1157_v64 _dafny.Array
				_ = _1157_v64
				var _len0_32 _dafny.Int = _dafny.IntOfInt64(15)
				_ = _len0_32
				var _nw184 _dafny.Array
				_ = _nw184
				if _len0_32.Cmp(_dafny.Zero) == 0 {
					_nw184 = _dafny.NewArray(_len0_32)
				} else {
					var _init32 func(_dafny.Int) _dafny.Int = func(_1158_i8 _dafny.Int) _dafny.Int {
						return (_1158_i8).Times((_dafny.Zero).Minus((_this).F9()))
					}
					_ = _init32
					var _element0_32 = _init32(_dafny.Zero)
					_ = _element0_32
					_nw184 = _dafny.NewArrayFromExample(_element0_32, nil, _len0_32)
					(_nw184).ArraySet1(_element0_32, 0)
					var _nativeLen0_32 = (_len0_32).Int()
					_ = _nativeLen0_32
					for _i0_32 := 1; _i0_32 < _nativeLen0_32; _i0_32++ {
						(_nw184).ArraySet1(_init32(_dafny.IntOf(_i0_32)), _i0_32)
					}
				}
				_1157_v64 = _nw184
				var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1157_v64), 0))
				_ = _index194
				(_1157_v64).ArraySet1((_this).F9(), (_index194).Int())
				var _arr46 _dafny.Array = _this.F7()
				_ = _arr46
				var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index195
				var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_1154_v62), 0))
				_ = _index196
				var _arr47 _dafny.Array = _this.F7()
				_ = _arr47
				var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index197
				var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1157_v64), 0))
				_ = _index198
				var _rhs231 bool = (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool)
				_ = _rhs231
				var _rhs232 _dafny.Int = ((_this).F9()).Plus(((_this).F9()).Minus((_this).F9()))
				_ = _rhs232
				var _rhs233 D1 = Companion_D1_.Create_DC2_(false)
				_ = _rhs233
				var _rhs234 bool = !(_dafny.Companion_Sequence_.IsPrefixOf(_1044_v1, _1044_v1)) || (p0)
				_ = _rhs234
				var _rhs235 _dafny.Int = (_this).F9()
				_ = _rhs235
				var _lhs191 _dafny.Array = _this.F7()
				_ = _lhs191
				var _lhs192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs192
				var _lhs193 *GlobalState = globalState
				_ = _lhs193
				var _lhs194 _dafny.Array = _1154_v62
				_ = _lhs194
				var _lhs195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(32), _dafny.ArrayLen((_1154_v62), 0))
				_ = _lhs195
				var _lhs196 _dafny.Array = _this.F7()
				_ = _lhs196
				var _lhs197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs197
				var _lhs198 _dafny.Array = _1157_v64
				_ = _lhs198
				var _lhs199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1157_v64), 0))
				_ = _lhs199
				(_lhs191).ArraySet1(_rhs231, (_lhs192).Int())
				_lhs193.F0 = _rhs232
				(_lhs194).ArraySet1(_rhs233, (_lhs195).Int())
				(_lhs196).ArraySet1(_rhs234, (_lhs197).Int())
				(_lhs198).ArraySet1(_rhs235, (_lhs199).Int())
				var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1157_v64), 0))
				_ = _index199
				(_1157_v64).ArraySet1(Companion_Default___.SafeModuloInt(Companion_Default___.SafeDivisionInt((_this).F9(), (_this).F9()), (_dafny.Zero).Minus((_this).F9())), (_index199).Int())
				var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(476), _dafny.ArrayLen((_1157_v64), 0))
				_ = _index200
				(_1157_v64).ArraySet1(Companion_Default___.SafeDivisionInt((func() _dafny.Int {
					if (_this).F10() {
						return (_this).F9()
					}
					return _dafny.IntOfInt64(569)
				})(), (func() _dafny.Int {
					if (_1104_v30).Contains((_this).F9()) {
						return (_1104_v30).Get((_this).F9()).(_dafny.Int)
					}
					return (_dafny.SetOf(_1046_v3)).Cardinality()
				})()), (_index200).Int())
				var _1159_v65 _dafny.Array
				_ = _1159_v65
				var _len0_33 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_33
				var _nw185 _dafny.Array
				_ = _nw185
				if _len0_33.Cmp(_dafny.Zero) == 0 {
					_nw185 = _dafny.NewArray(_len0_33)
				} else {
					var _init33 func(_dafny.Int) bool = func(_1160_i9 _dafny.Int) bool {
						return (_this).F6()
					}
					_ = _init33
					var _element0_33 = _init33(_dafny.Zero)
					_ = _element0_33
					_nw185 = _dafny.NewArrayFromExample(_element0_33, nil, _len0_33)
					(_nw185).ArraySet1(_element0_33, 0)
					var _nativeLen0_33 = (_len0_33).Int()
					_ = _nativeLen0_33
					for _i0_33 := 1; _i0_33 < _nativeLen0_33; _i0_33++ {
						(_nw185).ArraySet1(_init33(_dafny.IntOf(_i0_33)), _i0_33)
					}
				}
				_1159_v65 = _nw185
				var _1161_v66 _dafny.Map
				_ = _1161_v66
				_1161_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F5(), _1159_v65)
				(globalState).F0 = (_this).Fm4((_this).F9(), (_1161_v66).Cardinality(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(405))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_1162_i10 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('l')
				})), (_this).F10()), (_dafny.SetOf(p0, (_this).F6(), p0)).Cardinality(), globalState)
			} else {
				(globalState).F0 = ((_this).F9()).Minus(((_this).F9()).Plus((_this).F9()))
				(_this).F8_set_(p0)
				var _1163_v67 _dafny.Array
				_ = _1163_v67
				var _nw186 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(7))
				_ = _nw186
				_1163_v67 = _nw186
				var _1164_v68 bool
				_ = _1164_v68
				var _1165_v69 _dafny.Int
				_ = _1165_v69
				var _1166_v70 bool
				_ = _1166_v70
				var _out10 bool
				_ = _out10
				var _out11 _dafny.Int
				_ = _out11
				var _out12 bool
				_ = _out12
				_out10, _out11, _out12 = (_this).M5((_this).F9(), _dafny.MultiSetOf((_this).F10()), _1163_v67, p0, globalState)
				_1164_v68 = _out10
				_1165_v69 = _out11
				_1166_v70 = _out12
				var _1167_v71 _dafny.Array
				_ = _1167_v71
				var _len0_34 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_34
				var _nw187 _dafny.Array
				_ = _nw187
				if _len0_34.Cmp(_dafny.Zero) == 0 {
					_nw187 = _dafny.NewArray(_len0_34)
				} else {
					var _init34 func(_dafny.Int) bool = func(_1168_i11 _dafny.Int) bool {
						return (_this).F10()
					}
					_ = _init34
					var _element0_34 = _init34(_dafny.Zero)
					_ = _element0_34
					_nw187 = _dafny.NewArrayFromExample(_element0_34, nil, _len0_34)
					(_nw187).ArraySet1(_element0_34, 0)
					var _nativeLen0_34 = (_len0_34).Int()
					_ = _nativeLen0_34
					for _i0_34 := 1; _i0_34 < _nativeLen0_34; _i0_34++ {
						(_nw187).ArraySet1(_init34(_dafny.IntOf(_i0_34)), _i0_34)
					}
				}
				_1167_v71 = _nw187
				_1167_v71 = _1167_v71
				var _1169_v72 *C3
				_ = _1169_v72
				var _nw188 *C3 = New_C3_()
				_ = _nw188
				_nw188.Ctor__(_1166_v70, (_this).F5())
				_1169_v72 = _nw188
			}
			var _1170_v73 _dafny.Array
			_ = _1170_v73
			var _nw189 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(11))
			_ = _nw189
			_1170_v73 = _nw189
			var _1171_v74 _dafny.Map
			_ = _1171_v74
			_1171_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1170_v73, (_this).F10())
			var _1172_v75 D9
			_ = _1172_v75
			_1172_v75 = Companion_D9_.Create_DC27_(_1171_v74, _this.F8())
			var _1173_v76 _dafny.Sequence
			_ = _1173_v76
			_1173_v76 = _dafny.SeqOf(true)
			var _1174_v77 D8
			_ = _1174_v77
			_1174_v77 = Companion_D8_.Create_DC24_((func() _dafny.Int {
				if (_1043_v0).Contains((_this).F6()) {
					return (_1043_v0).Get((_this).F6()).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bwkuqlte")).Cardinality())
			})(), (_1172_v75).Dtor_cf43(), _1173_v76, (func() _dafny.Int {
				if (_this).F6() {
					return (_this).F9()
				}
				return (_this).F9()
			})(), (_this).F9())
			var _source17 D8 = _1174_v77
			_ = _source17
			if _source17.Is_DC24() {
				var _1175___mcc_h15 _dafny.Int = _source17.Get_().(D8_DC24).Cf34
				_ = _1175___mcc_h15
				var _1176___mcc_h16 bool = _source17.Get_().(D8_DC24).Cf35
				_ = _1176___mcc_h16
				var _1177___mcc_h17 _dafny.Sequence = _source17.Get_().(D8_DC24).Cf36
				_ = _1177___mcc_h17
				var _1178___mcc_h18 _dafny.Int = _source17.Get_().(D8_DC24).Cf37
				_ = _1178___mcc_h18
				var _1179___mcc_h19 _dafny.Int = _source17.Get_().(D8_DC24).Cf38
				_ = _1179___mcc_h19
				var _1180_cf38 _dafny.Int = _1179___mcc_h19
				_ = _1180_cf38
				var _1181_cf37 _dafny.Int = _1178___mcc_h18
				_ = _1181_cf37
				var _1182_cf36 _dafny.Sequence = _1177___mcc_h17
				_ = _1182_cf36
				var _1183_cf35 bool = _1176___mcc_h16
				_ = _1183_cf35
				var _1184_cf34 _dafny.Int = _1175___mcc_h15
				_ = _1184_cf34
				var _1185_v79 _dafny.Map
				_ = _1185_v79
				_1185_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1184_cf34, _1183_cf35)
				_1103_v29 = Companion_Default___.Fm16(p0, func() _dafny.Map {
					var _coll47 = _dafny.NewMapBuilder()
					_ = _coll47
					for _iter51 := _dafny.Iterate((_1185_v79).Keys().Elements()); ; {
						_compr_47, _ok51 := _iter51()
						if !_ok51 {
							break
						}
						var _1186_v78 _dafny.Int
						_1186_v78 = interface{}(_compr_47).(_dafny.Int)
						if (_1185_v79).Contains(_1186_v78) {
							_coll47.Add((_1186_v78).Minus((Companion_D0_.Create_DC0_(_1184_cf34)).Dtor_cf0()), _1180_cf38)
						}
					}
					return _coll47.ToMap()
				}(), _1044_v1, globalState)
				_1173_v76 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1173_v76, _1182_cf36), _dafny.Companion_Sequence_.Concatenate(_1182_cf36, _1173_v76))
				var _1187_v80 D9
				_ = _1187_v80
				_1187_v80 = Companion_D9_.Create_DC27_(_1171_v74, _this.F8())
				var _1188_v81 D9
				_ = _1188_v81
				_1188_v81 = Companion_D9_.Create_DC28_(_1187_v80)
				var _1189_v82 *C4
				_ = _1189_v82
				var _nw190 *C4 = New_C4_()
				_ = _nw190
				_nw190.Ctor__(_1170_v73, _1183_cf35, false, (_this).F5())
				_1189_v82 = _nw190
				var _1190_v83 _dafny.Map
				_ = _1190_v83
				_1190_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1188_v81, _1189_v82)
				var _1191_v84 _dafny.Map
				_ = _1191_v84
				_1191_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _1190_v83)
				var _1192_v85 *C2
				_ = _1192_v85
				var _nw191 *C2 = New_C2_()
				_ = _nw191
				_nw191.Ctor__((_this).F5())
				_1192_v85 = _nw191
				var _1193_v86 _dafny.Set
				_ = _1193_v86
				_1193_v86 = _dafny.SetOf(((func() _dafny.Map {
					if (_1191_v84).Contains((_this).F10()) {
						return (_1191_v84).Get((_this).F10()).(_dafny.Map)
					}
					return _1190_v83
				})()).Update(Companion_D9_.Create_DC28_(Companion_D9_.Create_DC26_(_1192_v85)), _1189_v82), _1190_v83)
				var _1194_v87 _dafny.Sequence
				_ = _1194_v87
				_1194_v87 = _dafny.SeqOf(_1193_v86)
				var _1195_v88 _dafny.MultiSet
				_ = _1195_v88
				_1195_v88 = _dafny.MultiSetOf(_1180_cf38)
				(globalState).F0 = (_dafny.Zero).Minus(((_1194_v87).Select((Companion_Default___.SafeIndex(((_1195_v88).Intersection(_dafny.MultiSetFromSeq(p1))).Cardinality(), _dafny.IntOfUint32((_1194_v87).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality())
				(globalState).F0 = (_this).F9()
			} else if _source17.Is_DC25() {
				var _1196___mcc_h20 bool = _source17.Get_().(D8_DC25).Cf39
				_ = _1196___mcc_h20
				var _1197___mcc_h21 bool = _source17.Get_().(D8_DC25).Cf40
				_ = _1197___mcc_h21
				var _1198_cf40 bool = _1197___mcc_h21
				_ = _1198_cf40
				var _1199_cf39 bool = _1196___mcc_h20
				_ = _1199_cf39
				_1104_v30 = (_1104_v30).Merge(_1104_v30)
				_1044_v1 = _1044_v1
				var _1200_v89 _dafny.Sequence
				_ = _1200_v89
				_1200_v89 = _dafny.SeqOf(_1173_v76)
				var _1201_v90 _dafny.Map
				_ = _1201_v90
				_1201_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _1200_v89)
				_1201_v90 = (_1201_v90).Update(((_this).Fm2(_1199_cf39, globalState)) == ((_this).F10()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(810))).Uint32(), func(coer57 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg57 _dafny.Int) interface{} {
						return coer57(arg57)
					}
				}((func(_1202_v76 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1203_i12 _dafny.Int) _dafny.Sequence {
						return _1202_v76
					}
				})(_1173_v76))))
				(globalState).F0 = (_this).F9()
			} else {
				var _1204___mcc_h22 _dafny.Map = _source17.Get_().(D8_DC23).Cf33
				_ = _1204___mcc_h22
				var _1205_cf33 _dafny.Map = _1204___mcc_h22
				_ = _1205_cf33
				(_this).F8_set_((_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))
				var _arr48 _dafny.Array = _this.F7()
				_ = _arr48
				var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
				_ = _index201
				var _rhs236 bool = !(_this.F8())
				_ = _rhs236
				var _rhs237 _dafny.Sequence = _1173_v76
				_ = _rhs237
				var _lhs200 _dafny.Array = _this.F7()
				_ = _lhs200
				var _lhs201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))
				_ = _lhs201
				(_lhs200).ArraySet1(_rhs236, (_lhs201).Int())
				_1173_v76 = _rhs237
				var _1206_v91 _dafny.Array
				_ = _1206_v91
				var _nw192 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(27))
				_ = _nw192
				_1206_v91 = _nw192
				var _1207_v92 bool
				_ = _1207_v92
				var _1208_v93 _dafny.Int
				_ = _1208_v93
				var _1209_v94 bool
				_ = _1209_v94
				var _out13 bool
				_ = _out13
				var _out14 _dafny.Int
				_ = _out14
				var _out15 bool
				_ = _out15
				_out13, _out14, _out15 = (_this).M5(Companion_Default___.SafeModuloInt((_this).F9(), (_this).F9()), _dafny.MultiSetOf(true, !((_this).F10()), !(p0)), _1206_v91, p0, globalState)
				_1207_v92 = _out13
				_1208_v93 = _out14
				_1209_v94 = _out15
				var _1210_v95 _dafny.Sequence
				_ = _1210_v95
				_1210_v95 = _dafny.SeqOf(_1173_v76)
				var _1211_v96 _dafny.Map
				_ = _1211_v96
				_1211_v96 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(669))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg58 _dafny.Int) interface{} {
						return coer58(arg58)
					}
				}((func(_1212_v93 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1213_i13 _dafny.Int) _dafny.Int {
						return (_dafny.SetOf(_1212_v93)).Cardinality()
					}
				})(_1208_v93))))
				var _1214_v97 _dafny.Map
				_ = _1214_v97
				_1214_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(954), _dafny.ArrayLen((_this.F7()), 0))).Int()).(bool))
				var _1215_v98 _dafny.Array
				_ = _1215_v98
				var _nwElement0_27 _dafny.Int = (_dafny.IntOfInt64(566)).Times((_this).F9())
				_ = _nwElement0_27
				var _nw193 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(24))
				_ = _nw193
				(_nw193).ArraySet1(_nwElement0_27, 0)
				(_nw193).ArraySet1(_dafny.IntOfUint32((_1210_v95).Cardinality()), 1)
				(_nw193).ArraySet1(_1208_v93, 2)
				(_nw193).ArraySet1((_this).F9(), 3)
				(_nw193).ArraySet1((_this).F9(), 4)
				(_nw193).ArraySet1((_1208_v93).Minus((_this).F9()), 5)
				(_nw193).ArraySet1((func() _dafny.Int {
					if p0 {
						return _1208_v93
					}
					return _1208_v93
				})(), 6)
				(_nw193).ArraySet1((_dafny.Zero).Minus(_1208_v93), 7)
				(_nw193).ArraySet1(_1208_v93, 8)
				(_nw193).ArraySet1(_dafny.IntOfInt64(-602), 9)
				(_nw193).ArraySet1(Companion_Default___.SafeDivisionInt(_1208_v93, _1208_v93), 10)
				(_nw193).ArraySet1((_this).F9(), 11)
				(_nw193).ArraySet1((_dafny.Zero).Minus(((_this).F9()).Times((_this).F9())), 12)
				(_nw193).ArraySet1(((_this).F9()).Minus((_this).F9()), 13)
				(_nw193).ArraySet1((_1211_v96).Cardinality(), 14)
				(_nw193).ArraySet1(Companion_Default___.SafeModuloInt(_1208_v93, _dafny.IntOfInt64(-348)), 15)
				(_nw193).ArraySet1((func() _dafny.Int {
					if false {
						return (_this).F9()
					}
					return (_this).F9()
				})(), 16)
				(_nw193).ArraySet1(((_this).F9()).Times(_1208_v93), 17)
				(_nw193).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
					if (_this).F6() {
						return _1208_v93
					}
					return _dafny.IntOfInt64(-714)
				})()), 18)
				(_nw193).ArraySet1((_1043_v0).Cardinality(), 19)
				(_nw193).ArraySet1(_1208_v93, 20)
				(_nw193).ArraySet1(_1208_v93, 21)
				(_nw193).ArraySet1((_this).F9(), 22)
				(_nw193).ArraySet1((_this).Fm4((_this).F9(), (_this).F9(), _1045_v2, (_1214_v97).Cardinality(), globalState), 23)
				_1215_v98 = _nw193
				var _1216_v99 _dafny.Sequence
				_ = _1216_v99
				_1216_v99 = _dafny.SeqOf(_1215_v98, (func() _dafny.Array {
					if (_this).F6() {
						return _1215_v98
					}
					return _1215_v98
				})(), _1215_v98)
				_1215_v98 = (_1216_v99).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1216_v99).Cardinality()))).Uint32()).(_dafny.Array)
			}
			(globalState).F0 = (_this).F9()
		} else {
			(_this).F7_set_(_this.F7())
			var _1217_v100 _dafny.Map
			_ = _1217_v100
			_1217_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(335), (_this).F6())
			var _1218_v101 _dafny.Map
			_ = _1218_v101
			_1218_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
			var _1219_v102 _dafny.MultiSet
			_ = _1219_v102
			_1219_v102 = _dafny.MultiSetOf(Companion_Default___.Fm27(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ulvsylkks")).Cardinality()), (func() bool {
				if (_1217_v100).Contains((_this).F9()) {
					return (_1217_v100).Get((_this).F9()).(bool)
				}
				return !(_this.F8())
			})(), (_this).F9(), _1217_v100, globalState), _1218_v101)
			var _1220_v103 _dafny.Array
			_ = _1220_v103
			var _nw194 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(19))
			_ = _nw194
			_1220_v103 = _nw194
			var _1221_v104 D7
			_ = _1221_v104
			_1221_v104 = Companion_D7_.Create_DC21_((_this).F9(), (_dafny.Zero).Minus((_this).F9()))
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1220_v103), 0))
			_ = _index202
			(_1220_v103).ArraySet1(_1221_v104, (_index202).Int())
			var _1222_v105 _dafny.Array
			_ = _1222_v105
			var _nw195 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(20))
			_ = _nw195
			_1222_v105 = _nw195
			var _1223_v106 _dafny.Array
			_ = _1223_v106
			var _len0_35 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_35
			var _nw196 _dafny.Array
			_ = _nw196
			if _len0_35.Cmp(_dafny.Zero) == 0 {
				_nw196 = _dafny.NewArray(_len0_35)
			} else {
				var _init35 func(_dafny.Int) _dafny.Int = func(_1224_i14 _dafny.Int) _dafny.Int {
					return (_1224_i14).Times(_dafny.IntOfInt64(153))
				}
				_ = _init35
				var _element0_35 = _init35(_dafny.Zero)
				_ = _element0_35
				_nw196 = _dafny.NewArrayFromExample(_element0_35, nil, _len0_35)
				(_nw196).ArraySet1(_element0_35, 0)
				var _nativeLen0_35 = (_len0_35).Int()
				_ = _nativeLen0_35
				for _i0_35 := 1; _i0_35 < _nativeLen0_35; _i0_35++ {
					(_nw196).ArraySet1(_init35(_dafny.IntOf(_i0_35)), _i0_35)
				}
			}
			_1223_v106 = _nw196
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))
			_ = _index203
			(_1223_v106).ArraySet1((_this).F9(), (_index203).Int())
			var _1225_v107 _dafny.Map
			_ = _1225_v107
			_1225_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).Fm4((_this).F9(), _dafny.IntOfUint32((_1044_v1).Cardinality()), _1045_v2, (_this).Fm4((_this).F9(), _dafny.IntOfInt64(46), _1045_v2, (_this).F9(), globalState), globalState), _1223_v106)
			var _1226_v108 _dafny.MultiSet
			_ = _1226_v108
			_1226_v108 = _dafny.MultiSetOf((_1225_v107).Cardinality())
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1220_v103), 0))
			_ = _index204
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))
			_ = _index205
			var _rhs238 _dafny.MultiSet = (Companion_Default___.Fm28((_this).F9(), _1218_v101, (_this).F10(), (_this).F9(), globalState)).Intersection(_1219_v102)
			_ = _rhs238
			var _rhs239 D7 = Companion_Default___.Fm29(_1044_v1, p1, _1226_v108, globalState)
			_ = _rhs239
			var _rhs240 _dafny.Array = _1222_v105
			_ = _rhs240
			var _rhs241 _dafny.Int = (_this).F9()
			_ = _rhs241
			var _rhs242 _dafny.Int = (_1043_v0).Cardinality()
			_ = _rhs242
			var _lhs202 _dafny.Array = _1220_v103
			_ = _lhs202
			var _lhs203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(103), _dafny.ArrayLen((_1220_v103), 0))
			_ = _lhs203
			var _lhs204 _dafny.Array = _1223_v106
			_ = _lhs204
			var _lhs205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))
			_ = _lhs205
			var _lhs206 *GlobalState = globalState
			_ = _lhs206
			_1219_v102 = _rhs238
			(_lhs202).ArraySet1(_rhs239, (_lhs203).Int())
			_1222_v105 = _rhs240
			(_lhs204).ArraySet1(_rhs241, (_lhs205).Int())
			_lhs206.F0 = _rhs242
			if !((_this).F6()) {
				var _1227_v109 T2
				_ = _1227_v109
				var _nw197 *C4 = New_C4_()
				_ = _nw197
				_nw197.Ctor__(_this.F7(), (_this).F10(), true, (_this).F5())
				_1227_v109 = _nw197
				_1227_v109 = _1227_v109
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))
				_ = _index206
				(_1223_v106).ArraySet1(_dafny.IntOfInt64(-90), (_index206).Int())
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))
				_ = _index207
				(_1223_v106).ArraySet1(Companion_Default___.SafeModuloInt((p1).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.Int), (_this).F9()), (_index207).Int())
				(_this).F8_set_(_this.F8())
				(_1227_v109).F8_set_((_this).F6())
			} else {
				var _1228_v110 _dafny.Array
				_ = _1228_v110
				var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(27))
				_ = _nw198
				_1228_v110 = _nw198
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1228_v110), 0))
				_ = _index208
				(_1228_v110).ArraySet1(_this.F7(), (_index208).Int())
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(315), _dafny.ArrayLen((_1228_v110), 0))
				_ = _index209
				(_1228_v110).ArraySet1((func() _dafny.Array {
					if ((_this).F5()).IsProperSubsetOf((_this).F5()) {
						return _this.F7()
					}
					return _this.F7()
				})(), (_index209).Int())
				(globalState).F0 = (_this).F9()
				(_this).F8_set_(p0)
				(globalState).F0 = (_1223_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))).Int()).(_dafny.Int)
				var _1229_v111 _dafny.Sequence
				_ = _1229_v111
				_1229_v111 = _dafny.SeqOf(p1, p1, p1, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(172))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}(func(_1230_i15 _dafny.Int) _dafny.Int {
					return _dafny.IntOfUint32((_dafny.SeqOf(false, _this.F8())).Cardinality())
				})), Companion_Default___.Fm8((_this).F9(), _this.F8(), _1044_v1, globalState))
				_1229_v111 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-870))).Uint32(), func(coer60 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg60 _dafny.Int) interface{} {
						return coer60(arg60)
					}
				}((func(_1231_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1232_i16 _dafny.Int) _dafny.Sequence {
						return _1231_p1
					}
				})(p1))), _1229_v111)
			}
			var _1233_v112 _dafny.Sequence
			_ = _1233_v112
			_1233_v112 = _dafny.SeqOf(p1, _dafny.Companion_Sequence_.Concatenate(p1, p1), _dafny.Companion_Sequence_.Concatenate(p1, p1))
			_1233_v112 = _1233_v112
			var _rhs243 _dafny.Int = (_dafny.Zero).Minus(((_1223_v106).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(31), _dafny.ArrayLen((_1223_v106), 0))).Int()).(_dafny.Int)).Times((func() _dafny.Int {
				if (_1226_v108).Contains((_1217_v100).Cardinality()) {
					return (_1226_v108).Multiplicity((_1217_v100).Cardinality())
				}
				return (_this).F9()
			})()))
			_ = _rhs243
			var _rhs244 _dafny.Sequence = _1044_v1
			_ = _rhs244
			var _rhs245 _dafny.Array = _1223_v106
			_ = _rhs245
			var _lhs207 *GlobalState = globalState
			_ = _lhs207
			_lhs207.F0 = _rhs243
			_1044_v1 = _rhs244
			_1223_v106 = _rhs245
		}
		var _1234_v113 _dafny.Map
		_ = _1234_v113
		_1234_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), Companion_Default___.SafeModuloInt((_this).F9(), (_this).F9()))
		_1234_v113 = (_1234_v113).Update((_this).F9(), (_this).F9())
		var _1235_v114 D11
		_ = _1235_v114
		_1235_v114 = Companion_D11_.Create_DC32_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), false))
		var _1236_v115 _dafny.Map
		_ = _1236_v115
		_1236_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), false)
		var _pat_let_tv27 = p0
		_ = _pat_let_tv27
		var _1237_v116 _dafny.Array
		_ = _1237_v116
		var _nwElement0_28 D11 = _1235_v114
		_ = _nwElement0_28
		var _nw199 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(29))
		_ = _nw199
		(_nw199).ArraySet1(_nwElement0_28, 0)
		(_nw199).ArraySet1(Companion_D11_.Create_DC32_(Companion_Default___.Fm9(globalState)), 1)
		(_nw199).ArraySet1(_1235_v114, 2)
		(_nw199).ArraySet1(_1235_v114, 3)
		(_nw199).ArraySet1(_1235_v114, 4)
		(_nw199).ArraySet1(_1235_v114, 5)
		(_nw199).ArraySet1(_1235_v114, 6)
		(_nw199).ArraySet1(Companion_D11_.Create_DC32_(_1236_v115), 7)
		(_nw199).ArraySet1(func(_pat_let34_0 D11) D11 {
			return func(_1238_dt__update__tmp_h2 D11) D11 {
				return func(_pat_let35_0 _dafny.Map) D11 {
					return func(_1239_dt__update_hcf51_h0 _dafny.Map) D11 {
						return Companion_D11_.Create_DC32_(_1239_dt__update_hcf51_h0)
					}(_pat_let35_0)
				}(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_pat_let_tv27, true))
			}(_pat_let34_0)
		}(_1235_v114), 8)
		(_nw199).ArraySet1(_1235_v114, 9)
		(_nw199).ArraySet1(_1235_v114, 10)
		(_nw199).ArraySet1(Companion_Default___.Fm30(globalState), 11)
		(_nw199).ArraySet1(_1235_v114, 12)
		(_nw199).ArraySet1(_1235_v114, 13)
		(_nw199).ArraySet1(_1235_v114, 14)
		(_nw199).ArraySet1(_1235_v114, 15)
		(_nw199).ArraySet1(_1235_v114, 16)
		(_nw199).ArraySet1(_1235_v114, 17)
		(_nw199).ArraySet1(_1235_v114, 18)
		(_nw199).ArraySet1(_1235_v114, 19)
		(_nw199).ArraySet1(_1235_v114, 20)
		(_nw199).ArraySet1(Companion_D11_.Create_DC32_(_1236_v115), 21)
		(_nw199).ArraySet1(_1235_v114, 22)
		(_nw199).ArraySet1(_1235_v114, 23)
		(_nw199).ArraySet1(_1235_v114, 24)
		(_nw199).ArraySet1(_1235_v114, 25)
		(_nw199).ArraySet1(_1235_v114, 26)
		(_nw199).ArraySet1(Companion_D11_.Create_DC32_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F8())), 27)
		(_nw199).ArraySet1(_1235_v114, 28)
		_1237_v116 = _nw199
		var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_1237_v116), 0))
		_ = _index210
		(_1237_v116).ArraySet1(_1235_v114, (_index210).Int())
		var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(382), _dafny.ArrayLen((_1237_v116), 0))
		_ = _index211
		(_1237_v116).ArraySet1(_1235_v114, (_index211).Int())
		var _1240_v117 _dafny.Array
		_ = _1240_v117
		var _nw200 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(5))
		_ = _nw200
		_1240_v117 = _nw200
		var _1241_v118 _dafny.CodePoint
		_ = _1241_v118
		_1241_v118 = _dafny.CodePoint('q')
		var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_1240_v117), 0))
		_ = _index212
		(_1240_v117).ArraySet1CodePoint(_1241_v118, (_index212).Int())
		var _1242_v119 _dafny.Array
		_ = _1242_v119
		var _nw201 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
		_ = _nw201
		_1242_v119 = _nw201
		var _1243_v120 _dafny.Map
		_ = _1243_v120
		_1243_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1234_v113, _this.F8())
		var _1244_v121 *C0
		_ = _1244_v121
		var _nw202 *C0 = New_C0_()
		_ = _nw202
		_nw202.Ctor__(_1243_v120)
		_1244_v121 = _nw202
		var _1245_v122 _dafny.Set
		_ = _1245_v122
		_1245_v122 = _dafny.SetOf(_1244_v121, _1244_v121, _1244_v121)
		var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1242_v119), 0))
		_ = _index213
		(_1242_v119).ArraySet1(_1245_v122, (_index213).Int())
		var _1246_v123 D2
		_ = _1246_v123
		_1246_v123 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC8_(Companion_D2_.Create_DC7_(_dafny.CodePoint('e'), _1234_v113, (_this).F9())))
		var _1247_v124 _dafny.Set
		_ = _1247_v124
		_1247_v124 = _dafny.SetOf((_this).F9(), (_this).F9())
		var _1248_v125 _dafny.Sequence
		_ = _1248_v125
		_1248_v125 = _dafny.SeqOf((_this).F5(), (_this).F5())
		var _1249_v126 D10
		_ = _1249_v126
		_1249_v126 = Companion_D10_.Create_DC31_(_1248_v125, _1241_v118)
		var _pat_let_tv28 = _1241_v118
		_ = _pat_let_tv28
		var _pat_let_tv29 = _1241_v118
		_ = _pat_let_tv29
		var _pat_let_tv30 = _1241_v118
		_ = _pat_let_tv30
		var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_1240_v117), 0))
		_ = _index214
		var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1242_v119), 0))
		_ = _index215
		var _rhs246 _dafny.CodePoint = func(_source18 D10) _dafny.CodePoint {
			if _source18.Is_DC30() {
				var _1250___mcc_h23 bool = _source18.Get_().(D10_DC30).Cf46
				_ = _1250___mcc_h23
				var _1251___mcc_h24 bool = _source18.Get_().(D10_DC30).Cf47
				_ = _1251___mcc_h24
				var _1252___mcc_h25 bool = _source18.Get_().(D10_DC30).Cf48
				_ = _1252___mcc_h25
				var _1253_cf48 bool = _1252___mcc_h25
				_ = _1253_cf48
				var _1254_cf47 bool = _1251___mcc_h24
				_ = _1254_cf47
				var _1255_cf46 bool = _1250___mcc_h23
				_ = _1255_cf46
				return _pat_let_tv28
			} else if _source18.Is_DC31() {
				var _1256___mcc_h26 _dafny.Sequence = _source18.Get_().(D10_DC31).Cf49
				_ = _1256___mcc_h26
				var _1257___mcc_h27 _dafny.CodePoint = _source18.Get_().(D10_DC31).Cf50
				_ = _1257___mcc_h27
				var _1258_cf50 _dafny.CodePoint = _1257___mcc_h27
				_ = _1258_cf50
				var _1259_cf49 _dafny.Sequence = _1256___mcc_h26
				_ = _1259_cf49
				return _pat_let_tv29
			} else {
				var _1260___mcc_h28 _dafny.MultiSet = _source18.Get_().(D10_DC29).Cf45
				_ = _1260___mcc_h28
				var _1261_cf45 _dafny.MultiSet = _1260___mcc_h28
				_ = _1261_cf45
				return _pat_let_tv30
			}
		}(_1249_v126)
		_ = _rhs246
		var _rhs247 _dafny.Set = _1245_v122
		_ = _rhs247
		var _rhs248 D2 = _1246_v123
		_ = _rhs248
		var _rhs249 _dafny.Set = _1247_v124
		_ = _rhs249
		var _rhs250 _dafny.Map = (_1043_v0).Merge((func() _dafny.Map {
			if (_this).F6() {
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.IntOfInt64(-252))
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F9())
		})())
		_ = _rhs250
		var _lhs208 _dafny.Array = _1240_v117
		_ = _lhs208
		var _lhs209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(33), _dafny.ArrayLen((_1240_v117), 0))
		_ = _lhs209
		var _lhs210 _dafny.Array = _1242_v119
		_ = _lhs210
		var _lhs211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(510), _dafny.ArrayLen((_1242_v119), 0))
		_ = _lhs211
		(_lhs208).ArraySet1CodePoint(_rhs246, (_lhs209).Int())
		(_lhs210).ArraySet1(_rhs247, (_lhs211).Int())
		_1246_v123 = _rhs248
		_1247_v124 = _rhs249
		_1043_v0 = _rhs250
	}
}
func (_this *C5) M5(p0 _dafny.Int, p1 _dafny.MultiSet, p2 _dafny.Array, p3 bool, globalState *GlobalState) (bool, _dafny.Int, bool) {
	{
		var r0 bool = false
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 bool = false
		_ = r2
		var _1262_v2 _dafny.Map
		_ = _1262_v2
		_1262_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F9())
		var _1263_v3 _dafny.Set
		_ = _1263_v3
		_1263_v3 = _dafny.SetOf(_1262_v2, _1262_v2)
		var _1264_v4 *C0
		_ = _1264_v4
		var _nw203 *C0 = New_C0_()
		_ = _nw203
		_nw203.Ctor__(func() _dafny.Map {
			var _coll48 = _dafny.NewMapBuilder()
			_ = _coll48
			for _iter52 := _dafny.Iterate((func() _dafny.Map {
				var _coll49 = _dafny.NewMapBuilder()
				_ = _coll49
				for _iter53 := _dafny.Iterate((_1263_v3).Elements()); ; {
					_compr_49, _ok53 := _iter53()
					if !_ok53 {
						break
					}
					var _1265_v1 _dafny.Map
					_1265_v1 = interface{}(_compr_49).(_dafny.Map)
					if (_1263_v3).Contains(_1265_v1) {
						_coll49.Add(_1265_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F6())).Cardinality())
					}
				}
				return _coll49.ToMap()
			}()).Keys().Elements()); ; {
				_compr_48, _ok52 := _iter52()
				if !_ok52 {
					break
				}
				var _1266_v0 _dafny.Map
				_1266_v0 = interface{}(_compr_48).(_dafny.Map)
				if (func() _dafny.Map {
					var _coll50 = _dafny.NewMapBuilder()
					_ = _coll50
					for _iter54 := _dafny.Iterate((_1263_v3).Elements()); ; {
						_compr_50, _ok54 := _iter54()
						if !_ok54 {
							break
						}
						var _1265_v1 _dafny.Map
						_1265_v1 = interface{}(_compr_50).(_dafny.Map)
						if (_1263_v3).Contains(_1265_v1) {
							_coll50.Add(_1265_v1, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F6())).Cardinality())
						}
					}
					return _coll50.ToMap()
				}()).Contains(_1266_v0) {
					_coll48.Add(_1266_v0, (_this).F6())
				}
			}
			return _coll48.ToMap()
		}())
		_1264_v4 = _nw203
		var _1267_v6 D2
		_ = _1267_v6
		_1267_v6 = Companion_D2_.Create_DC5_((_this).F6())
		var _1268_v7 _dafny.Map
		_ = _1268_v7
		_1268_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1267_v6, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(502))).Uint32(), func(coer61 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}((func(_1269_p1 _dafny.MultiSet) func(_dafny.Int) _dafny.Int {
			return func(_1270_i0 _dafny.Int) _dafny.Int {
				return (_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(814))).Uint32(), func(coer62 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg62 _dafny.Int) interface{} {
						return coer62(arg62)
					}
				}((func(_1271_p1 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
					return func(_1272_i1 _dafny.Int) _dafny.MultiSet {
						return _1271_p1
					}
				})(_1269_p1))))).Cardinality()
			}
		})(p1)))).Cardinality()))
		var _1273_v8 _dafny.Map
		_ = _1273_v8
		_1273_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1268_v7, (_this).F10())
		var _1274_v9 D4
		_ = _1274_v9
		_1274_v9 = Companion_D4_.Create_DC13_(_this.F8(), func() _dafny.Map {
			var _coll51 = _dafny.NewMapBuilder()
			_ = _coll51
			for _iter55 := _dafny.Iterate((_1273_v8).Keys().Elements()); ; {
				_compr_51, _ok55 := _iter55()
				if !_ok55 {
					break
				}
				var _1275_v5 _dafny.Map
				_1275_v5 = interface{}(_compr_51).(_dafny.Map)
				if (_1273_v8).Contains(_1275_v5) {
					_coll51.Add(_1275_v5, false)
				}
			}
			return _coll51.ToMap()
		}())
		var _1276_v10 T2
		_ = _1276_v10
		var _nw204 *C4 = New_C4_()
		_ = _nw204
		_nw204.Ctor__(_this.F7(), (_this).F6(), (!(p3)) && ((_1274_v9).Dtor_cf15()), p1)
		_1276_v10 = _nw204
		var _1277_v11 _dafny.Map
		_ = _1277_v11
		_1277_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F10())
		var _1278_v12 _dafny.Set
		_ = _1278_v12
		_1278_v12 = _dafny.SetOf(_1277_v11, _1277_v11, (_1277_v11).Merge(_1277_v11))
		var _1279_v13 _dafny.MultiSet
		_ = _1279_v13
		_1279_v13 = _dafny.MultiSetOf(_1277_v11, _1277_v11)
		var _1280_v15 _dafny.Map
		_ = _1280_v15
		_1280_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F9()).Cmp(p0) >= 0, func() _dafny.Set {
			var _coll52 = _dafny.NewBuilder()
			_ = _coll52
			for _iter56 := _dafny.Iterate((_1279_v13).Elements()); ; {
				_compr_52, _ok56 := _iter56()
				if !_ok56 {
					break
				}
				var _1281_v14 _dafny.Map
				_1281_v14 = interface{}(_compr_52).(_dafny.Map)
				if (_1279_v13).Contains(_1281_v14) {
					_coll52.Add(_1281_v14)
				}
			}
			return _coll52.ToSet()
		}())
		var _1282_v16 _dafny.MultiSet
		_ = _1282_v16
		_1282_v16 = _dafny.MultiSetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_this).F10())).Cardinality())
		var _1283_v17 _dafny.Map
		_ = _1283_v17
		_1283_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1282_v16).Cardinality(), _1262_v2)
		var _rhs251 T2 = (func() T2 {
			if false {
				return _1276_v10
			}
			return _1276_v10
		})()
		_ = _rhs251
		var _rhs252 _dafny.Set = (func() _dafny.Set {
			if (_1280_v15).Contains((_1283_v17).Equals((_1283_v17).Update((_this).F9(), _1262_v2))) {
				return (_1280_v15).Get((_1283_v17).Equals((_1283_v17).Update((_this).F9(), _1262_v2))).(_dafny.Set)
			}
			return Companion_Default___.Fm31(globalState)
		})()
		_ = _rhs252
		var _rhs253 _dafny.Int = p0
		_ = _rhs253
		var _lhs212 *GlobalState = globalState
		_ = _lhs212
		_1276_v10 = _rhs251
		_1278_v12 = _rhs252
		_lhs212.F0 = _rhs253
		var _1284_v18 _dafny.Map
		_ = _1284_v18
		_1284_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1276_v10.F7(), p3)
		var _1285_v19 D9
		_ = _1285_v19
		_1285_v19 = Companion_D9_.Create_DC27_(_1284_v18, (_1276_v10).F6())
		if !(((func() D9 {
			if (_this).F10() {
				return Companion_D9_.Create_DC27_(_1284_v18, (_this).F10())
			}
			return _1285_v19
		})()).Dtor_cf43()) {
			_1262_v2 = (_1262_v2).Update(p0, ((_this).F9()).Minus((_this).F9()))
			var _1286_v20 _dafny.Map
			_ = _1286_v20
			_1286_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1276_v10).F6(), (_this).F6())
			var _1287_v21 _dafny.Map
			_ = _1287_v21
			_1287_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_1286_v20).Cardinality())
			var _1288_v22 D7
			_ = _1288_v22
			_1288_v22 = Companion_D7_.Create_DC19_(_1287_v21)
			var _1289_v23 _dafny.Set
			_ = _1289_v23
			_1289_v23 = _dafny.SetOf(_1288_v22, _1288_v22, _1288_v22, _1288_v22)
			(_1276_v10).F8_set_(((_1289_v23).Union(_1289_v23)).IsDisjointFrom((_1289_v23).Difference(_dafny.SetOf(_1288_v22))))
			(globalState).F0 = p0
			var _1290_v24 _dafny.Set
			_ = _1290_v24
			_1290_v24 = _dafny.SetOf(_this.F8(), _1276_v10.F8(), (_this).F10())
			r0 = (_1290_v24).Equals(Companion_Default___.Fm32((_1276_v10).F6(), (_this).F10(), !(true), globalState))
			var _1291_v25 _dafny.Sequence
			_ = _1291_v25
			_1291_v25 = _dafny.SeqOf(_1276_v10.F8())
			var _1292_v26 _dafny.Array
			_ = _1292_v26
			var _nwElement0_29 _dafny.Sequence = _1291_v25
			_ = _nwElement0_29
			var _nw205 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(6))
			_ = _nw205
			(_nw205).ArraySet1(_nwElement0_29, 0)
			(_nw205).ArraySet1(_1291_v25, 1)
			(_nw205).ArraySet1(_1291_v25, 2)
			(_nw205).ArraySet1(_1291_v25, 3)
			(_nw205).ArraySet1(_dafny.Companion_Sequence_.Update(_1291_v25, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1291_v25).Cardinality()))).Uint32(), false), 4)
			(_nw205).ArraySet1(_1291_v25, 5)
			_1292_v26 = _nw205
			var _1293_v27 _dafny.Map
			_ = _1293_v27
			_1293_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1292_v26, (_this).F9())
			_1293_v27 = (_1293_v27).Update(_1292_v26, _dafny.IntOfInt64(854))
		} else {
			var _1294_v28 _dafny.CodePoint
			_ = _1294_v28
			_1294_v28 = _dafny.CodePoint('x')
			_1294_v28 = _dafny.CodePoint('d')
			r1 = (_this).F9()
			var _1295_v29 _dafny.Array
			_ = _1295_v29
			var _len0_36 _dafny.Int = _dafny.IntOfInt64(8)
			_ = _len0_36
			var _nw206 _dafny.Array
			_ = _nw206
			if _len0_36.Cmp(_dafny.Zero) == 0 {
				_nw206 = _dafny.NewArray(_len0_36)
			} else {
				var _init36 func(_dafny.Int) _dafny.Int = (func(_1296_v10 T2, _1297_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_1298_i2 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_1298_i2, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1296_v10.F8(), _1297_p0)).Cardinality())
					}
				})(_1276_v10, p0)
				_ = _init36
				var _element0_36 = _init36(_dafny.Zero)
				_ = _element0_36
				_nw206 = _dafny.NewArrayFromExample(_element0_36, nil, _len0_36)
				(_nw206).ArraySet1(_element0_36, 0)
				var _nativeLen0_36 = (_len0_36).Int()
				_ = _nativeLen0_36
				for _i0_36 := 1; _i0_36 < _nativeLen0_36; _i0_36++ {
					(_nw206).ArraySet1(_init36(_dafny.IntOf(_i0_36)), _i0_36)
				}
			}
			_1295_v29 = _nw206
			var _1299_v30 _dafny.Map
			_ = _1299_v30
			_1299_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1295_v29, (_this).F10())
			var _arr49 _dafny.Array = _this.F7()
			_ = _arr49
			var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index216
			(_arr49).ArraySet1((func() bool {
				if (_1299_v30).Contains(_1295_v29) {
					return (_1299_v30).Get(_1295_v29).(bool)
				}
				return (_1276_v10).F6()
			})(), (_index216).Int())
			var _1300_v31 _dafny.Sequence
			_ = _1300_v31
			_1300_v31 = _dafny.UnicodeSeqOfUtf8Bytes("w")
			var _arr50 _dafny.Array = _this.F7()
			_ = _arr50
			var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(132), _dafny.ArrayLen((_this.F7()), 0))
			_ = _index217
			(_arr50).ArraySet1(!_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_1300_v31, _1300_v31), _1294_v28), (_index217).Int())
			(globalState).F0 = (_this).F9()
			(_1276_v10).M2((_1300_v31).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1300_v31).Cardinality()))).Uint32()).(_dafny.CodePoint), globalState)
		}
		if true {
			var _1301_v32 _dafny.CodePoint
			_ = _1301_v32
			_1301_v32 = _dafny.CodePoint('e')
			(_1276_v10).M2(_1301_v32, globalState)
			if (p1).Equals(_dafny.MultiSetOf((_1276_v10).Fm3(globalState))) {
				r1 = (_this).F9()
				r0 = (_1276_v10).F6()
				r2 = true
				var _1302_v33 _dafny.Sequence
				_ = _1302_v33
				_1302_v33 = _dafny.UnicodeSeqOfUtf8Bytes("nfwac")
				var _1303_v34 _dafny.Array
				_ = _1303_v34
				var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1302_v33, _1302_v33)
				_ = _nwElement0_30
				var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(12))
				_ = _nw207
				(_nw207).ArraySet1(_nwElement0_30, 0)
				(_nw207).ArraySet1(_1302_v33, 1)
				(_nw207).ArraySet1(_1302_v33, 2)
				(_nw207).ArraySet1(_1302_v33, 3)
				(_nw207).ArraySet1(_1302_v33, 4)
				(_nw207).ArraySet1(_1302_v33, 5)
				(_nw207).ArraySet1(_1302_v33, 6)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("grcliofu"), _1302_v33), 7)
				(_nw207).ArraySet1(_dafny.Companion_Sequence_.Update(_1302_v33, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1302_v33).Cardinality()))).Uint32(), _dafny.CodePoint('g')), 8)
				(_nw207).ArraySet1(_1302_v33, 9)
				(_nw207).ArraySet1(_1302_v33, 10)
				(_nw207).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(115))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg63 _dafny.Int) interface{} {
						return coer63(arg63)
					}
				}((func(_1304_v32 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_1305_i3 _dafny.Int) _dafny.CodePoint {
						return _1304_v32
					}
				})(_1301_v32))), 11)
				_1303_v34 = _nw207
				var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1303_v34), 0))
				_ = _index218
				(_1303_v34).ArraySet1(_1302_v33, (_index218).Int())
				var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_1303_v34), 0))
				_ = _index219
				(_1303_v34).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1302_v33, _1302_v33), _dafny.Companion_Sequence_.Concatenate(_1302_v33, _dafny.UnicodeSeqOfUtf8Bytes("qo"))), (_index219).Int())
				var _1306_v35 *C4
				_ = _1306_v35
				var _nw208 *C4 = New_C4_()
				_ = _nw208
				_nw208.Ctor__(_1276_v10.F7(), _this.F8(), _1276_v10.F8(), Companion_Default___.Fm14(_1262_v2, _this.F8(), _dafny.IntOfUint32((Companion_Default___.Fm17(globalState)).Cardinality()), true, globalState))
				_1306_v35 = _nw208
			} else {
				var _1307_v36 _dafny.Sequence
				_ = _1307_v36
				_1307_v36 = _dafny.SeqOf((_this).F10(), (_this).F6(), p3, (_1276_v10).F6())
				(_1276_v10).M1(p0, _this.F8(), !_dafny.Companion_Sequence_.Contains(_1307_v36, _1276_v10.F8()), globalState)
				(_this).F8_set_((Companion_Default___.Fm14(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(622), p0), _1276_v10.F8(), p0, (_1276_v10).F6(), globalState)).IsSubsetOf((_this).F5()))
				(globalState).F0 = (p0).Minus((_this).Fm4(_dafny.IntOfUint32((_dafny.SeqOf(_this.F8(), (_1276_v10).F6())).Cardinality()), (_this).F9(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wrmdqy"), (_this).F6()), p0, globalState))
				(_1276_v10).M2(_1301_v32, globalState)
				var _arr51 _dafny.Array = _1276_v10.F7()
				_ = _arr51
				var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1276_v10.F7()), 0))
				_ = _index220
				(_arr51).ArraySet1(true, (_index220).Int())
				var _1308_v37 _dafny.Array
				_ = _1308_v37
				var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
				_ = _nw209
				_1308_v37 = _nw209
				var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1308_v37), 0))
				_ = _index221
				(_1308_v37).ArraySet1((_this).F9(), (_index221).Int())
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1308_v37), 0))
				_ = _index222
				(_1308_v37).ArraySet1((_dafny.Zero).Minus((_this).F9()), (_index222).Int())
				var _1309_v38 _dafny.Sequence
				_ = _1309_v38
				_1309_v38 = _dafny.UnicodeSeqOfUtf8Bytes("r")
				var _1310_v39 _dafny.Map
				_ = _1310_v39
				_1310_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1309_v38, (_1276_v10).F6())
				var _1311_v40 D3
				_ = _1311_v40
				_1311_v40 = Companion_D3_.Create_DC9_(_1308_v37)
				var _1312_v41 D5
				_ = _1312_v41
				_1312_v41 = Companion_D5_.Create_DC16_(_1311_v40, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("kd")).Cardinality()), p0, (_this).F9(), _1264_v4)
				var _arr52 _dafny.Array = _1276_v10.F7()
				_ = _arr52
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1276_v10.F7()), 0))
				_ = _index223
				var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1308_v37), 0))
				_ = _index224
				var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1308_v37), 0))
				_ = _index225
				var _rhs254 bool = (p1).IsProperSubsetOf((_1276_v10).F5())
				_ = _rhs254
				var _rhs255 _dafny.Int = ((_1276_v10).Fm4(p0, (_this).F9(), _1310_v39, (_this).F9(), globalState)).Minus(((_1312_v41).Dtor_cf22()).Minus((_this).F9()))
				_ = _rhs255
				var _rhs256 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).Fm4((_this).F9(), (_dafny.Zero).Minus(p0), _1310_v39, (_this).F9(), globalState), (_this).F9())
				_ = _rhs256
				var _lhs213 _dafny.Array = _1276_v10.F7()
				_ = _lhs213
				var _lhs214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1276_v10.F7()), 0))
				_ = _lhs214
				var _lhs215 _dafny.Array = _1308_v37
				_ = _lhs215
				var _lhs216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(468), _dafny.ArrayLen((_1308_v37), 0))
				_ = _lhs216
				var _lhs217 _dafny.Array = _1308_v37
				_ = _lhs217
				var _lhs218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_1308_v37), 0))
				_ = _lhs218
				(_lhs213).ArraySet1(_rhs254, (_lhs214).Int())
				(_lhs215).ArraySet1(_rhs255, (_lhs216).Int())
				(_lhs217).ArraySet1(_rhs256, (_lhs218).Int())
			}
			var _1313_v42 _dafny.Array
			_ = _1313_v42
			var _len0_37 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_37
			var _nw210 _dafny.Array
			_ = _nw210
			if _len0_37.Cmp(_dafny.Zero) == 0 {
				_nw210 = _dafny.NewArray(_len0_37)
			} else {
				var _init37 func(_dafny.Int) _dafny.Sequence = func(_1314_i4 _dafny.Int) _dafny.Sequence {
					return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("ljs"), _dafny.UnicodeSeqOfUtf8Bytes("ow"))
				}
				_ = _init37
				var _element0_37 = _init37(_dafny.Zero)
				_ = _element0_37
				_nw210 = _dafny.NewArrayFromExample(_element0_37, nil, _len0_37)
				(_nw210).ArraySet1(_element0_37, 0)
				var _nativeLen0_37 = (_len0_37).Int()
				_ = _nativeLen0_37
				for _i0_37 := 1; _i0_37 < _nativeLen0_37; _i0_37++ {
					(_nw210).ArraySet1(_init37(_dafny.IntOf(_i0_37)), _i0_37)
				}
			}
			_1313_v42 = _nw210
			var _1315_v43 _dafny.Sequence
			_ = _1315_v43
			_1315_v43 = _dafny.UnicodeSeqOfUtf8Bytes("qwwlgc")
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1313_v42), 0))
			_ = _index226
			(_1313_v42).ArraySet1(_1315_v43, (_index226).Int())
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1313_v42), 0))
			_ = _index227
			(_1313_v42).ArraySet1((func() _dafny.Sequence {
				if (_this).F6() {
					return _1315_v43
				}
				return _1315_v43
			})(), (_index227).Int())
			var _1316_v44 _dafny.Array
			_ = _1316_v44
			var _nw211 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
			_ = _nw211
			_1316_v44 = _nw211
			var _1317_v45 _dafny.Sequence
			_ = _1317_v45
			_1317_v45 = _dafny.SeqOf(_dafny.IntOfUint32((_1315_v43).Cardinality()), _dafny.IntOfInt64(776), p0, _dafny.IntOfInt64(-193), _dafny.IntOfInt64(933))
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1316_v44), 0))
			_ = _index228
			(_1316_v44).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1317_v45, _1317_v45), (_index228).Int())
			var _1318_v46 _dafny.Sequence
			_ = _1318_v46
			_1318_v46 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_1277_v11).Cardinality(), p0), _1317_v45))
			var _1319_v47 _dafny.Array
			_ = _1319_v47
			var _nwElement0_31 _dafny.Int = _dafny.IntOfInt64(929)
			_ = _nwElement0_31
			var _nw212 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(8))
			_ = _nw212
			(_nw212).ArraySet1(_nwElement0_31, 0)
			(_nw212).ArraySet1(p0, 1)
			(_nw212).ArraySet1(p0, 2)
			(_nw212).ArraySet1(p0, 3)
			(_nw212).ArraySet1(p0, 4)
			(_nw212).ArraySet1((_this).F9(), 5)
			(_nw212).ArraySet1(p0, 6)
			(_nw212).ArraySet1(_dafny.IntOfUint32((_1315_v43).Cardinality()), 7)
			_1319_v47 = _nw212
			var _1320_v48 _dafny.Sequence
			_ = _1320_v48
			_1320_v48 = _dafny.SeqOf(_1319_v47, _1319_v47)
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(592), _dafny.ArrayLen((_1316_v44), 0))
			_ = _index229
			(_1316_v44).ArraySet1(_dafny.Companion_Sequence_.Update((_1318_v46).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1318_v46).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_dafny.IntOfUint32((_1320_v48).Cardinality())).Plus((_dafny.Zero).Minus((_this).F9())), _dafny.IntOfUint32(((_1318_v46).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1318_v46).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(-440))), (_index229).Int())
			var _1321_v49 _dafny.MultiSet
			_ = _1321_v49
			_1321_v49 = _dafny.MultiSetOf(_1301_v32)
			var _1322_v50 _dafny.Sequence
			_ = _1322_v50
			_1322_v50 = _dafny.SeqOf(_1321_v49, _1321_v49)
			var _1323_v51 D17
			_ = _1323_v51
			_1323_v51 = Companion_D17_.Create_DC46_(_1321_v49)
			var _1324_v52 _dafny.Array
			_ = _1324_v52
			var _nwElement0_32 _dafny.Sequence = _dafny.SeqOf(_dafny.MultiSetOf(_1301_v32))
			_ = _nwElement0_32
			var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(23))
			_ = _nw213
			(_nw213).ArraySet1(_nwElement0_32, 0)
			(_nw213).ArraySet1(_1322_v50, 1)
			(_nw213).ArraySet1(_1322_v50, 2)
			(_nw213).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(395))).Uint32(), func(coer64 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg64 _dafny.Int) interface{} {
					return coer64(arg64)
				}
			}((func(_1325_v42 _dafny.Array) func(_dafny.Int) _dafny.MultiSet {
				return func(_1326_i5 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetFromSeq((_1325_v42).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(563), _dafny.ArrayLen((_1325_v42), 0))).Int()).(_dafny.Sequence))
				}
			})(_1313_v42))), 3)
			(_nw213).ArraySet1(_dafny.SeqOf(_1321_v49, _1321_v49), 4)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Update(_1322_v50, (Companion_Default___.SafeIndex((_dafny.MultiSetOf(p3, _1276_v10.F8())).Cardinality(), _dafny.IntOfUint32((_1322_v50).Cardinality()))).Uint32(), _1321_v49), 5)
			(_nw213).ArraySet1(_1322_v50, 6)
			(_nw213).ArraySet1(_dafny.SeqOf(_dafny.MultiSetOf(_1301_v32, _1301_v32), _1321_v49, _1321_v49, _1321_v49, _1321_v49), 7)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1322_v50, _1322_v50), 8)
			(_nw213).ArraySet1(_1322_v50, 9)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1322_v50, _dafny.SeqOf(Companion_Default___.Fm33(false, (_this).F9(), globalState), _1321_v49)), 10)
			(_nw213).ArraySet1(_1322_v50, 11)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm34(globalState), _dafny.SeqOf(_1321_v49)), 12)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1321_v49, _1321_v49), _dafny.Companion_Sequence_.Update(_1322_v50, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1322_v50).Cardinality()))).Uint32(), _dafny.MultiSetOf(_1301_v32))), 13)
			(_nw213).ArraySet1((func() _dafny.Sequence {
				if (_this).F10() {
					return _dafny.SeqOf(_1321_v49)
				}
				return _1322_v50
			})(), 14)
			(_nw213).ArraySet1(_1322_v50, 15)
			(_nw213).ArraySet1(_1322_v50, 16)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm34(globalState), _1322_v50), 17)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1321_v49, _1321_v49, _dafny.MultiSetOf(_1301_v32)), _1322_v50), 18)
			(_nw213).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_1321_v49, (_1323_v51).Dtor_cf73(), _dafny.MultiSetOf(_dafny.CodePoint('g'), _1301_v32, _1301_v32), _1321_v49, _dafny.MultiSetOf(_1301_v32)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(130))).Uint32(), func(coer65 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg65 _dafny.Int) interface{} {
					return coer65(arg65)
				}
			}((func(_1327_v49 _dafny.MultiSet) func(_dafny.Int) _dafny.MultiSet {
				return func(_1328_i6 _dafny.Int) _dafny.MultiSet {
					return _1327_v49
				}
			})(_1321_v49)))), 19)
			(_nw213).ArraySet1(Companion_Default___.Fm34(globalState), 20)
			(_nw213).ArraySet1(_1322_v50, 21)
			(_nw213).ArraySet1((func() _dafny.Sequence {
				if (_1276_v10).F6() {
					return _1322_v50
				}
				return _1322_v50
			})(), 22)
			_1324_v52 = _nw213
			var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1324_v52), 0))
			_ = _index230
			(_1324_v52).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_1321_v49).Update(_1301_v32, Companion_Default___.Abs((_1317_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-144), _dafny.IntOfUint32((_1317_v45).Cardinality()))).Uint32()).(_dafny.Int))), _1321_v49, _1321_v49, _1321_v49), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.SeqOf((_1321_v49).Update(_1301_v32, Companion_Default___.Abs((_1317_v45).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-144), _dafny.IntOfUint32((_1317_v45).Cardinality()))).Uint32()).(_dafny.Int))), _1321_v49, _1321_v49, _1321_v49)).Cardinality()))).Uint32(), _1321_v49), (_index230).Int())
			var _1329_v53 _dafny.Sequence
			_ = _1329_v53
			_1329_v53 = _dafny.SeqOf(_1322_v50)
			var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1324_v52), 0))
			_ = _index231
			var _rhs257 _dafny.Int = (_this).F9()
			_ = _rhs257
			var _rhs258 _dafny.Sequence = (_1329_v53).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1329_v53).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs258
			var _rhs259 _dafny.Int = _dafny.IntOfUint32((_1317_v45).Cardinality())
			_ = _rhs259
			var _lhs219 *GlobalState = globalState
			_ = _lhs219
			var _lhs220 _dafny.Array = _1324_v52
			_ = _lhs220
			var _lhs221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(413), _dafny.ArrayLen((_1324_v52), 0))
			_ = _lhs221
			var _lhs222 *GlobalState = globalState
			_ = _lhs222
			_lhs219.F0 = _rhs257
			(_lhs220).ArraySet1(_rhs258, (_lhs221).Int())
			_lhs222.F0 = _rhs259
		} else {
			var _1330_v54 _dafny.Array
			_ = _1330_v54
			var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
			_ = _nw214
			_1330_v54 = _nw214
			var _index232 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
			_ = _index232
			(_1330_v54).ArraySet1(p0, (_index232).Int())
			var _index233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
			_ = _index233
			var _rhs260 bool = _this.F8()
			_ = _rhs260
			var _rhs261 bool = !((_this).F10())
			_ = _rhs261
			var _rhs262 _dafny.Int = p0
			_ = _rhs262
			var _lhs223 *C5 = _this
			_ = _lhs223
			var _lhs224 _dafny.Array = _1330_v54
			_ = _lhs224
			var _lhs225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
			_ = _lhs225
			r0 = _rhs260
			_lhs223.F8_set_(_rhs261)
			(_lhs224).ArraySet1(_rhs262, (_lhs225).Int())
			var _source19 D0 = Companion_D0_.Create_DC1_()
			_ = _source19
			if _source19.Is_DC1() {
				(_this).F8_set_((_1276_v10).F6())
				var _1331_v55 _dafny.Sequence
				_ = _1331_v55
				_1331_v55 = _dafny.UnicodeSeqOfUtf8Bytes("rq")
				var _1332_v56 _dafny.Map
				_ = _1332_v56
				_1332_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F7(), _1331_v55)
				var _1333_v57 *C1
				_ = _1333_v57
				var _nw215 *C1 = New_C1_()
				_ = _nw215
				_nw215.Ctor__(_1276_v10.F8(), (func() _dafny.Sequence {
					if (_1332_v56).Contains(_1276_v10.F7()) {
						return (_1332_v56).Get(_1276_v10.F7()).(_dafny.Sequence)
					}
					return _1331_v55
				})(), _this.F8(), ((_1276_v10).F5()).Intersection(p1))
				_1333_v57 = _nw215
				var _1334_v58 _dafny.Sequence
				_ = _1334_v58
				_1334_v58 = _dafny.SeqOf(!(p3))
				var _index234 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
				_ = _index234
				(_1330_v54).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_1334_v58).Cardinality()), (_this).F9()), (_index234).Int())
				var _arr53 _dafny.Array = _1276_v10.F7()
				_ = _arr53
				var _index235 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_1276_v10.F7()), 0))
				_ = _index235
				(_arr53).ArraySet1(!(!(_1276_v10.F8())), (_index235).Int())
				var _1335_v59 _dafny.CodePoint
				_ = _1335_v59
				_1335_v59 = _dafny.CodePoint('j')
				var _1336_v60 _dafny.Sequence
				_ = _1336_v60
				_1336_v60 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update((_1333_v57).F13(), (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32(((_1333_v57).F13()).Cardinality()))).Uint32(), _1335_v59))
				var _1337_v61 _dafny.Sequence
				_ = _1337_v61
				_1337_v61 = _dafny.SeqOf(_1277_v11, _1277_v11, _1277_v11, _1277_v11)
				var _1338_v62 _dafny.Sequence
				_ = _1338_v62
				_1338_v62 = _dafny.SeqOf((_this).F9())
				var _arr54 _dafny.Array = _1276_v10.F7()
				_ = _arr54
				var _index236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_1276_v10.F7()), 0))
				_ = _index236
				var _rhs263 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_1336_v60).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1337_v61, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1337_v61).Cardinality()))).Uint32(), _1277_v11))).Cardinality(), _dafny.IntOfUint32((_1336_v60).Cardinality()))).Uint32()).(_dafny.Sequence), (Companion_Default___.SafeIndex((_1330_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32(((_1336_v60).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Update(_1337_v61, (Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1337_v61).Cardinality()))).Uint32(), _1277_v11))).Cardinality(), _dafny.IntOfUint32((_1336_v60).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()))).Uint32(), _1335_v59), (_1333_v57).F13())).Cardinality())
				_ = _rhs263
				var _rhs264 bool = p3
				_ = _rhs264
				var _rhs265 bool = !((((_this).F9()).Plus((_this).F9())).Cmp((_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_1338_v62).Cardinality())))) != 0)
				_ = _rhs265
				var _lhs226 *GlobalState = globalState
				_ = _lhs226
				var _lhs227 _dafny.Array = _1276_v10.F7()
				_ = _lhs227
				var _lhs228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_1276_v10.F7()), 0))
				_ = _lhs228
				_lhs226.F0 = _rhs263
				r2 = _rhs264
				(_lhs227).ArraySet1(_rhs265, (_lhs228).Int())
			} else {
				var _1339___mcc_h0 _dafny.Int = _source19.Get_().(D0_DC0).Cf0
				_ = _1339___mcc_h0
				var _1340_cf0 _dafny.Int = _1339___mcc_h0
				_ = _1340_cf0
				(_this).F8_set_((func() bool {
					if (_1276_v10).F6() {
						return (_this).F10()
					}
					return ((_this).F6()) == ((_this).F10())
				})())
				var _1341_v63 _dafny.Sequence
				_ = _1341_v63
				_1341_v63 = _dafny.SeqOf(_1330_v54)
				var _index237 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
				_ = _index237
				var _index238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
				_ = _index238
				var _rhs266 bool = _this.F8()
				_ = _rhs266
				var _rhs267 _dafny.Int = (_this).F9()
				_ = _rhs267
				var _rhs268 _dafny.Int = ((_this).F9()).Minus((_dafny.Zero).Minus((_1330_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))).Int()).(_dafny.Int)))
				_ = _rhs268
				var _rhs269 _dafny.Int = _1340_cf0
				_ = _rhs269
				var _rhs270 _dafny.Array = (_1341_v63).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_1330_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))).Int()).(_dafny.Int), (_1330_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))).Int()).(_dafny.Int)), _dafny.IntOfUint32((_1341_v63).Cardinality()))).Uint32()).(_dafny.Array)
				_ = _rhs270
				var _lhs229 *C5 = _this
				_ = _lhs229
				var _lhs230 _dafny.Array = _1330_v54
				_ = _lhs230
				var _lhs231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
				_ = _lhs231
				var _lhs232 _dafny.Array = _1330_v54
				_ = _lhs232
				var _lhs233 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(640), _dafny.ArrayLen((_1330_v54), 0))
				_ = _lhs233
				_lhs229.F8_set_(_rhs266)
				(_lhs230).ArraySet1(_rhs267, (_lhs231).Int())
				r1 = _rhs268
				(_lhs232).ArraySet1(_rhs269, (_lhs233).Int())
				_1330_v54 = _rhs270
				var _1342_v64 D0
				_ = _1342_v64
				_1342_v64 = Companion_D0_.Create_DC1_()
				var _1343_v65 _dafny.Sequence
				_ = _1343_v65
				_1343_v65 = _dafny.SeqOf((_1276_v10).F6(), (p0).Cmp(_dafny.IntOfInt64(364)) >= 0, ((_dafny.Zero).Minus(p0)).Cmp(_1340_cf0) > 0)
				var _rhs271 D0 = _1342_v64
				_ = _rhs271
				var _rhs272 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1343_v65, _1343_v65)
				_ = _rhs272
				var _rhs273 _dafny.Int = (_this).F9()
				_ = _rhs273
				var _lhs234 *GlobalState = globalState
				_ = _lhs234
				_1342_v64 = _rhs271
				_1343_v65 = _rhs272
				_lhs234.F0 = _rhs273
				r1 = Companion_Default___.SafeDivisionInt(_1340_cf0, p0)
			}
			var _1344_v66 _dafny.Sequence
			_ = _1344_v66
			_1344_v66 = _dafny.SeqOf(_dafny.SeqOf(p1), _dafny.SeqOf((_1276_v10).F5()))
			var _1345_v67 _dafny.CodePoint
			_ = _1345_v67
			_1345_v67 = _dafny.CodePoint('u')
			var _1346_v68 D10
			_ = _1346_v68
			_1346_v68 = Companion_D10_.Create_DC31_((_1344_v66).Select((Companion_Default___.SafeIndex((_this).F9(), _dafny.IntOfUint32((_1344_v66).Cardinality()))).Uint32()).(_dafny.Sequence), _1345_v67)
			_1346_v68 = _1346_v68
			var _1347_v69 _dafny.Sequence
			_ = _1347_v69
			_1347_v69 = _dafny.UnicodeSeqOfUtf8Bytes("mu")
			var _1348_v70 _dafny.Map
			_ = _1348_v70
			_1348_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), _dafny.IntOfUint32((_1347_v69).Cardinality()))
			var _1349_v71 _dafny.Set
			_ = _1349_v71
			_1349_v71 = _dafny.SetOf((_1276_v10).F6())
			var _1350_v72 T1
			_ = _1350_v72
			var _nw216 *C3 = New_C3_()
			_ = _nw216
			_nw216.Ctor__((_1276_v10).Fm1((_this).Fm3(globalState), (_1348_v70).Cardinality(), (_1349_v71).Cardinality(), globalState), _dafny.MultiSetOf((_this).F6()))
			_1350_v72 = _nw216
			_1350_v72 = _1350_v72
			var _1351_v73 _dafny.Map
			_ = _1351_v73
			_1351_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_1264_v4).F11())
			var _1352_v74 *C0
			_ = _1352_v74
			var _nw217 *C0 = New_C0_()
			_ = _nw217
			_nw217.Ctor__((func() _dafny.Map {
				if (_1351_v73).Contains((_1276_v10).F6()) {
					return (_1351_v73).Get((_1276_v10).F6()).(_dafny.Map)
				}
				return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1262_v2, (_1276_v10).F6())
			})())
			_1352_v74 = _nw217
		}
		var _1353_v75 _dafny.Sequence
		_ = _1353_v75
		_1353_v75 = _dafny.SeqOf(_1276_v10.F8())
		var _1354_v76 D6
		_ = _1354_v76
		_1354_v76 = Companion_D6_.Create_DC17_(_1353_v75)
		var _1355_v77 D11
		_ = _1355_v77
		_1355_v77 = Companion_D11_.Create_DC33_(_1276_v10.F8(), false)
		var _pat_let_tv31 = _1276_v10
		_ = _pat_let_tv31
		var _pat_let_tv32 = p3
		_ = _pat_let_tv32
		var _1356_v78 _dafny.Map
		_ = _1356_v78
		_1356_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(func(_pat_let36_0 D6) D6 {
			return func(_1357_dt__update__tmp_h0 D6) D6 {
				return func(_pat_let37_0 _dafny.Sequence) D6 {
					return func(_1358_dt__update_hcf24_h0 _dafny.Sequence) D6 {
						return Companion_D6_.Create_DC17_(_1358_dt__update_hcf24_h0)
					}(_pat_let37_0)
				}(_dafny.SeqOf(_pat_let_tv31.F8(), _pat_let_tv32))
			}(_pat_let36_0)
		}(_1354_v76), _1355_v77)
		_1356_v78 = (_1356_v78).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D6_.Create_DC17_(_1353_v75), _1355_v77))
		var _1359_i7 _dafny.Int
		_ = _1359_i7
		_1359_i7 = _dafny.Zero
		{
			for (_this).F6() {
				{
					if (_1359_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L12
					}
					_1359_i7 = (_1359_i7).Plus(_dafny.One)
					var _1360_v79 _dafny.Array
					_ = _1360_v79
					var _len0_38 _dafny.Int = _dafny.IntOfInt64(15)
					_ = _len0_38
					var _nw218 _dafny.Array
					_ = _nw218
					if _len0_38.Cmp(_dafny.Zero) == 0 {
						_nw218 = _dafny.NewArray(_len0_38)
					} else {
						var _init38 func(_dafny.Int) _dafny.Int = func(_1361_i8 _dafny.Int) _dafny.Int {
							return (_1361_i8).Minus((_this).F9())
						}
						_ = _init38
						var _element0_38 = _init38(_dafny.Zero)
						_ = _element0_38
						_nw218 = _dafny.NewArrayFromExample(_element0_38, nil, _len0_38)
						(_nw218).ArraySet1(_element0_38, 0)
						var _nativeLen0_38 = (_len0_38).Int()
						_ = _nativeLen0_38
						for _i0_38 := 1; _i0_38 < _nativeLen0_38; _i0_38++ {
							(_nw218).ArraySet1(_init38(_dafny.IntOf(_i0_38)), _i0_38)
						}
					}
					_1360_v79 = _nw218
					var _index239 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1360_v79), 0))
					_ = _index239
					(_1360_v79).ArraySet1(p0, (_index239).Int())
					var _index240 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1360_v79), 0))
					_ = _index240
					(_1360_v79).ArraySet1(Companion_Default___.SafeModuloInt(p0, p0), (_index240).Int())
					var _1362_v80 _dafny.Sequence
					_ = _1362_v80
					_1362_v80 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_1353_v75))
					var _1363_v81 D10
					_ = _1363_v81
					_1363_v81 = Companion_D10_.Create_DC31_(_1362_v80, _dafny.CodePoint('q'))
					var _1364_v82 _dafny.Sequence
					_ = _1364_v82
					_1364_v82 = _dafny.UnicodeSeqOfUtf8Bytes("hm")
					if !(!_dafny.Companion_Sequence_.Contains(_1364_v82, (_1363_v81).Dtor_cf50())) {
						var _1365_v83 _dafny.Map
						_ = _1365_v83
						_1365_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("uexbyodl"), (_this).F6())
						r1 = (_dafny.Zero).Minus((_this).Fm4((_this).F9(), (_1360_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1360_v79), 0))).Int()).(_dafny.Int), _1365_v83, (_dafny.Zero).Minus(_dafny.IntOfUint32((_1364_v82).Cardinality())), globalState))
						(_1276_v10).F8_set_(p3)
						(globalState).F0 = p0
						var _arr55 _dafny.Array = _1276_v10.F7()
						_ = _arr55
						var _index241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1276_v10.F7()), 0))
						_ = _index241
						(_arr55).ArraySet1(p3, (_index241).Int())
						var _arr56 _dafny.Array = _1276_v10.F7()
						_ = _arr56
						var _index242 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(665), _dafny.ArrayLen((_1276_v10.F7()), 0))
						_ = _index242
						(_arr56).ArraySet1(true, (_index242).Int())
						(_1276_v10).M2(Companion_Default___.Fm16((_1276_v10).F6(), _1262_v2, _1364_v82, globalState), globalState)
					} else {
						(globalState).F0 = p0
						r2 = (_this).F10()
						var _1366_v84 _dafny.MultiSet
						_ = _1366_v84
						_1366_v84 = _dafny.MultiSetOf(_1360_v79)
						var _1367_v85 D18
						_ = _1367_v85
						_1367_v85 = Companion_D18_.Create_DC48_(_1366_v84)
						_1366_v84 = (_1366_v84).Intersection((_1366_v84).Intersection((_1367_v85).Dtor_cf77()))
						(_1276_v10).F8_set_(!((_this).F10()))
						var _1368_v86 _dafny.Array
						_ = _1368_v86
						var _nw219 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(24))
						_ = _nw219
						_1368_v86 = _nw219
						var _1369_v87 _dafny.CodePoint
						_ = _1369_v87
						_1369_v87 = _dafny.CodePoint('c')
						var _1370_v88 _dafny.Array
						_ = _1370_v88
						var _nwElement0_33 _dafny.CodePoint = _1369_v87
						_ = _nwElement0_33
						var _nw220 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(22))
						_ = _nw220
						(_nw220).ArraySet1CodePoint(_nwElement0_33, 0)
						(_nw220).ArraySet1CodePoint(_1369_v87, 1)
						(_nw220).ArraySet1CodePoint(_1369_v87, 2)
						(_nw220).ArraySet1CodePoint(_1369_v87, 3)
						(_nw220).ArraySet1CodePoint(_1369_v87, 4)
						(_nw220).ArraySet1CodePoint(_dafny.CodePoint('p'), 5)
						(_nw220).ArraySet1CodePoint(_1369_v87, 6)
						(_nw220).ArraySet1CodePoint(_dafny.CodePoint('a'), 7)
						(_nw220).ArraySet1CodePoint(_1369_v87, 8)
						(_nw220).ArraySet1CodePoint(_1369_v87, 9)
						(_nw220).ArraySet1CodePoint(_1369_v87, 10)
						(_nw220).ArraySet1CodePoint((_1363_v81).Dtor_cf50(), 11)
						(_nw220).ArraySet1CodePoint(_1369_v87, 12)
						(_nw220).ArraySet1CodePoint(_1369_v87, 13)
						(_nw220).ArraySet1CodePoint(_1369_v87, 14)
						(_nw220).ArraySet1CodePoint(_1369_v87, 15)
						(_nw220).ArraySet1CodePoint(_1369_v87, 16)
						(_nw220).ArraySet1CodePoint(_1369_v87, 17)
						(_nw220).ArraySet1CodePoint(_1369_v87, 18)
						(_nw220).ArraySet1CodePoint(_1369_v87, 19)
						(_nw220).ArraySet1CodePoint(_1369_v87, 20)
						(_nw220).ArraySet1CodePoint(_1369_v87, 21)
						_1370_v88 = _nw220
						var _1371_v89 _dafny.Array
						_ = _1371_v89
						var _nwElement0_34 _dafny.Array = _1370_v88
						_ = _nwElement0_34
						var _nw221 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(25))
						_ = _nw221
						(_nw221).ArraySet1(_nwElement0_34, 0)
						(_nw221).ArraySet1(_1370_v88, 1)
						(_nw221).ArraySet1(_1370_v88, 2)
						(_nw221).ArraySet1(_1370_v88, 3)
						(_nw221).ArraySet1(_1370_v88, 4)
						(_nw221).ArraySet1(_1370_v88, 5)
						(_nw221).ArraySet1(_1370_v88, 6)
						(_nw221).ArraySet1(_1370_v88, 7)
						(_nw221).ArraySet1(_1370_v88, 8)
						(_nw221).ArraySet1(_1370_v88, 9)
						(_nw221).ArraySet1(_1370_v88, 10)
						(_nw221).ArraySet1(_1370_v88, 11)
						(_nw221).ArraySet1(_1370_v88, 12)
						(_nw221).ArraySet1(_1370_v88, 13)
						(_nw221).ArraySet1(_1370_v88, 14)
						(_nw221).ArraySet1(_1370_v88, 15)
						(_nw221).ArraySet1(_1370_v88, 16)
						(_nw221).ArraySet1(_1370_v88, 17)
						(_nw221).ArraySet1(_1370_v88, 18)
						(_nw221).ArraySet1(_1370_v88, 19)
						(_nw221).ArraySet1(_1370_v88, 20)
						(_nw221).ArraySet1(_1370_v88, 21)
						(_nw221).ArraySet1(_1370_v88, 22)
						(_nw221).ArraySet1(_1370_v88, 23)
						(_nw221).ArraySet1(_1370_v88, 24)
						_1371_v89 = _nw221
						var _1372_v90 D4
						_ = _1372_v90
						_1372_v90 = Companion_D4_.Create_DC12_(_1371_v89)
						var _1373_v91 D4
						_ = _1373_v91
						_1373_v91 = Companion_D4_.Create_DC14_(_1372_v90)
						var _index243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1368_v86), 0))
						_ = _index243
						(_1368_v86).ArraySet1(_1373_v91, (_index243).Int())
						var _index244 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1368_v86), 0))
						_ = _index244
						var _rhs274 D4 = _1373_v91
						_ = _rhs274
						var _rhs275 _dafny.Map = _1262_v2
						_ = _rhs275
						var _rhs276 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("lsl"), (Companion_Default___.SafeIndex((_1360_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1360_v79), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lsl")).Cardinality()))).Uint32(), _1369_v87)).Cardinality())).Times((_1360_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1360_v79), 0))).Int()).(_dafny.Int))
						_ = _rhs276
						var _lhs235 _dafny.Array = _1368_v86
						_ = _lhs235
						var _lhs236 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(858), _dafny.ArrayLen((_1368_v86), 0))
						_ = _lhs236
						(_lhs235).ArraySet1(_rhs274, (_lhs236).Int())
						_1262_v2 = _rhs275
						r1 = _rhs276
					}
					var _1374_v92 *C4
					_ = _1374_v92
					var _nw222 *C4 = New_C4_()
					_ = _nw222
					_nw222.Ctor__(_this.F7(), _dafny.Companion_Sequence_.Contains(_1353_v75, (_this).F6()), (_this).F10(), (_this).F5())
					_1374_v92 = _nw222
					var _1375_v93 _dafny.Array
					_ = _1375_v93
					var _nw223 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(21))
					_ = _nw223
					_1375_v93 = _nw223
					var _index245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1375_v93), 0))
					_ = _index245
					(_1375_v93).ArraySet1(_dafny.SeqOf(_dafny.IntOfInt64(412)), (_index245).Int())
					var _1376_v94 _dafny.Sequence
					_ = _1376_v94
					_1376_v94 = _dafny.SeqOf((_this).F9())
					var _1377_v95 _dafny.Map
					_ = _1377_v95
					_1377_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1364_v82, p3)
					var _index246 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1375_v93), 0))
					_ = _index246
					var _rhs277 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1376_v94, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(708))).Uint32(), func(coer66 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg66 _dafny.Int) interface{} {
							return coer66(arg66)
						}
					}(func(_1378_i9 _dafny.Int) _dafny.Int {
						return (_this).F9()
					}))), _1376_v94)
					_ = _rhs277
					var _rhs278 bool = (_dafny.IntOfInt64(752)).Cmp(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(194), (_1276_v10).Fm4((_1360_v79).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(424), _dafny.ArrayLen((_1360_v79), 0))).Int()).(_dafny.Int), (_this).F9(), _1377_v95, p0, globalState))) > 0
					_ = _rhs278
					var _rhs279 _dafny.Sequence = _1364_v82
					_ = _rhs279
					var _rhs280 _dafny.Array = _1276_v10.F7()
					_ = _rhs280
					var _lhs237 _dafny.Array = _1375_v93
					_ = _lhs237
					var _lhs238 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(317), _dafny.ArrayLen((_1375_v93), 0))
					_ = _lhs238
					var _lhs239 T2 = _1276_v10
					_ = _lhs239
					(_lhs237).ArraySet1(_rhs277, (_lhs238).Int())
					r0 = _rhs278
					_1364_v82 = _rhs279
					_lhs239.F7_set_(_rhs280)
					goto C12
				}
			C12:
			}
			goto L12
		}
	L12:
		r0 = (_1276_v10).F6()
		r1 = (_this).F9()
		r2 = (_1276_v10).F6()
		return r0, r1, r2
	}
}
func (_this *C5) F9() _dafny.Int {
	{
		return _this._f9
	}
}
func (_this *C5) F10() bool {
	{
		return _this._f10
	}
}

// End of class C5

// Definition of class C6
type C6 struct {
	F3 _dafny.Set
	F4 _dafny.CodePoint
}

func New_C6_() *C6 {
	_this := C6{}

	_this.F3 = _dafny.EmptySet
	_this.F4 = _dafny.CodePoint('D')
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
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &C6{}

func (_this *C6) Ctor__(f3 _dafny.Set, f4 _dafny.CodePoint) {
	{
		(_this).F3 = f3
		(_this).F4 = f4
	}
}
func (_this *C6) Fm0(p0 _dafny.Sequence, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Int {
	{
		if !_dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("adv"), _dafny.UnicodeSeqOfUtf8Bytes("iy")) {
			if false {
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(731))).Uint32(), func(coer67 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg67 _dafny.Int) interface{} {
						return coer67(arg67)
					}
				}(func(_1379_i0 _dafny.Int) _dafny.CodePoint {
					return _this.F4
				}))).Cardinality())))).Cardinality()
			} else {
				return _dafny.IntOfInt64(737)
			}
		} else {
			return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sayxmsb")).Cardinality())).Plus((_dafny.MultiSetOf(true)).Cardinality())
		}
	}
}
func (_this *C6) M0(p0 _dafny.Int, p1 _dafny.Sequence, p2 _dafny.Int, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _1380_v0 bool
		_ = _1380_v0
		_1380_v0 = false
		var _1381_i0 _dafny.Int
		_ = _1381_i0
		_1381_i0 = _dafny.Zero
		{
			for _1380_v0 {
				{
					if (_1381_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L13
					}
					_1381_i0 = (_1381_i0).Plus(_dafny.One)
					var _1382_v1 _dafny.Sequence
					_ = _1382_v1
					_1382_v1 = _dafny.SeqOf((p0).Minus(p2))
					r2 = (_1382_v1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1382_v1).Cardinality()))).Uint32()).(_dafny.Int)
					var _1383_v2 _dafny.Array
					_ = _1383_v2
					var _nw224 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(18))
					_ = _nw224
					_1383_v2 = _nw224
					var _index247 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1383_v2), 0))
					_ = _index247
					(_1383_v2).ArraySet1(p0, (_index247).Int())
					var _1384_v3 _dafny.Array
					_ = _1384_v3
					var _len0_39 _dafny.Int = _dafny.IntOfInt64(5)
					_ = _len0_39
					var _nw225 _dafny.Array
					_ = _nw225
					if _len0_39.Cmp(_dafny.Zero) == 0 {
						_nw225 = _dafny.NewArray(_len0_39)
					} else {
						var _init39 func(_dafny.Int) bool = (func(_1385_v0 bool) func(_dafny.Int) bool {
							return func(_1386_i1 _dafny.Int) bool {
								return _1385_v0
							}
						})(_1380_v0)
						_ = _init39
						var _element0_39 = _init39(_dafny.Zero)
						_ = _element0_39
						_nw225 = _dafny.NewArrayFromExample(_element0_39, nil, _len0_39)
						(_nw225).ArraySet1(_element0_39, 0)
						var _nativeLen0_39 = (_len0_39).Int()
						_ = _nativeLen0_39
						for _i0_39 := 1; _i0_39 < _nativeLen0_39; _i0_39++ {
							(_nw225).ArraySet1(_init39(_dafny.IntOf(_i0_39)), _i0_39)
						}
					}
					_1384_v3 = _nw225
					var _index248 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1383_v2), 0))
					_ = _index248
					(_1383_v2).ArraySet1(p2, (_index248).Int())
					var _1387_v4 _dafny.Set
					_ = _1387_v4
					_1387_v4 = _dafny.SetOf(_1384_v3, _1384_v3)
					var _index249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1383_v2), 0))
					_ = _index249
					var _index250 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1383_v2), 0))
					_ = _index250
					var _rhs281 _dafny.Array = _1383_v2
					_ = _rhs281
					var _rhs282 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("gutgmh"))).Cardinality()), p0)
					_ = _rhs282
					var _rhs283 _dafny.Array = _1384_v3
					_ = _rhs283
					var _rhs284 _dafny.Int = Companion_Default___.SafeModuloInt(p0, ((_1387_v4).Intersection(_1387_v4)).Cardinality())
					_ = _rhs284
					var _lhs240 _dafny.Array = _1383_v2
					_ = _lhs240
					var _lhs241 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(440), _dafny.ArrayLen((_1383_v2), 0))
					_ = _lhs241
					var _lhs242 _dafny.Array = _1383_v2
					_ = _lhs242
					var _lhs243 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(684), _dafny.ArrayLen((_1383_v2), 0))
					_ = _lhs243
					_1383_v2 = _rhs281
					(_lhs240).ArraySet1(_rhs282, (_lhs241).Int())
					_1384_v3 = _rhs283
					(_lhs242).ArraySet1(_rhs284, (_lhs243).Int())
					var _1388_v5 _dafny.Sequence
					_ = _1388_v5
					_1388_v5 = _dafny.SeqOf(p1)
					var _1389_v6 *C4
					_ = _1389_v6
					var _nw226 *C4 = New_C4_()
					_ = _nw226
					_nw226.Ctor__(_1384_v3, Companion_Default___.Fm35(_1388_v5, _1380_v0, globalState), _1380_v0, _dafny.MultiSetOf(_1380_v0))
					_1389_v6 = _nw226
					var _1390_v7 _dafny.Set
					_ = _1390_v7
					_1390_v7 = _dafny.SetOf(_1380_v0, (_1389_v6).Fm3(globalState))
					var _1391_v8 _dafny.Map
					_ = _1391_v8
					_1391_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1390_v7, !(!(true)))
					_1391_v8 = (_1391_v8).Update(_1390_v7, _1380_v0)
					goto C13
				}
			C13:
			}
			goto L13
		}
	L13:
		if Companion_Default___.Fm35(_dafny.SeqOf(p1), !(_1380_v0), globalState) {
			var _1392_v9 _dafny.Map
			_ = _1392_v9
			_1392_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, false)
			var _1393_v10 _dafny.Map
			_ = _1393_v10
			_1393_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1392_v9, _1380_v0)
			var _1394_v11 _dafny.Sequence
			_ = _1394_v11
			_1394_v11 = _dafny.SeqOf(_1393_v10)
			_1380_v0 = (_this.F3).Contains(((_1394_v11).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((p1).Cardinality()), _dafny.IntOfUint32((_1394_v11).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
			var _1395_v12 _dafny.Array
			_ = _1395_v12
			var _nw227 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(15))
			_ = _nw227
			_1395_v12 = _nw227
			var _1396_v13 _dafny.Map
			_ = _1396_v13
			_1396_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1395_v12, _1380_v0)
			_1396_v13 = _1396_v13
			var _1397_v14 _dafny.Sequence
			_ = _1397_v14
			_1397_v14 = _dafny.SeqOf(true)
			var _1398_v15 D18
			_ = _1398_v15
			_1398_v15 = Companion_D18_.Create_DC49_(_dafny.Companion_Sequence_.Update(_1397_v14, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1397_v14).Cardinality()))).Uint32(), _1380_v0), _dafny.IntOfInt64(460))
			var _1399_v16 _dafny.Map
			_ = _1399_v16
			_1399_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1398_v15, false)
			if !((func() bool {
				if (_1399_v16).Contains(_1398_v15) {
					return (_1399_v16).Get(_1398_v15).(bool)
				}
				return _1380_v0
			})()) {
				(globalState).F0 = p0
				var _1400_v17 _dafny.Array
				_ = _1400_v17
				var _nw228 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw228
				_1400_v17 = _nw228
				var _1401_v18 _dafny.Array
				_ = _1401_v18
				var _nw229 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(2))
				_ = _nw229
				_1401_v18 = _nw229
				var _index251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1400_v17), 0))
				_ = _index251
				(_1400_v17).ArraySet1(_1401_v18, (_index251).Int())
				var _1402_v19 _dafny.Map
				_ = _1402_v19
				_1402_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _1403_v20 _dafny.Map
				_ = _1403_v20
				_1403_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1402_v19, _1380_v0)
				var _1404_v21 *C0
				_ = _1404_v21
				var _nw230 *C0 = New_C0_()
				_ = _nw230
				_nw230.Ctor__(_1403_v20)
				_1404_v21 = _nw230
				var _1405_v22 D6
				_ = _1405_v22
				_1405_v22 = Companion_D6_.Create_DC18_(_1404_v21, _1397_v14, p0)
				var _1406_v23 _dafny.Sequence
				_ = _1406_v23
				_1406_v23 = _dafny.SeqOf(_dafny.IntOfInt64(849), _dafny.IntOfInt64(655), (_1405_v22).Dtor_cf27(), p2)
				var _index252 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1400_v17), 0))
				_ = _index252
				var _rhs285 _dafny.Int = (_dafny.SetOf(_1395_v12, _1395_v12)).Cardinality()
				_ = _rhs285
				var _rhs286 _dafny.Array = _1401_v18
				_ = _rhs286
				var _rhs287 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1406_v23, _dafny.SeqOf(p2))).Cardinality())).Times(_dafny.IntOfUint32((_1397_v14).Cardinality()))
				_ = _rhs287
				var _lhs244 _dafny.Array = _1400_v17
				_ = _lhs244
				var _lhs245 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(327), _dafny.ArrayLen((_1400_v17), 0))
				_ = _lhs245
				r2 = _rhs285
				(_lhs244).ArraySet1(_rhs286, (_lhs245).Int())
				r2 = _rhs287
				r1 = _1380_v0
				r1 = _1380_v0
				var _1407_v25 D2
				_ = _1407_v25
				_1407_v25 = Companion_D2_.Create_DC7_(_this.F4, _1402_v19, (_dafny.Zero).Minus(p2))
				var _1408_v26 _dafny.Map
				_ = _1408_v26
				_1408_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1398_v15)
				var _1409_v27 _dafny.Sequence
				_ = _1409_v27
				_1409_v27 = _dafny.SeqOf(func() _dafny.Map {
					var _coll53 = _dafny.NewMapBuilder()
					_ = _coll53
					for _iter57 := _dafny.Iterate(((_1407_v25).Dtor_cf7()).Keys().Elements()); ; {
						_compr_53, _ok57 := _iter57()
						if !_ok57 {
							break
						}
						var _1410_v24 _dafny.Int
						_1410_v24 = interface{}(_compr_53).(_dafny.Int)
						if ((_1407_v25).Dtor_cf7()).Contains(_1410_v24) {
							_coll53.Add((_1410_v24).Times(p0), Companion_D18_.Create_DC49_(_1397_v14, _dafny.IntOfUint32((_1397_v14).Cardinality())))
						}
					}
					return _coll53.ToMap()
				}(), _1408_v26)
				var _pat_let_tv33 = p0
				_ = _pat_let_tv33
				var _1411_v28 _dafny.Map
				_ = _1411_v28
				_1411_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, func(_pat_let38_0 D18) D18 {
					return func(_1412_dt__update__tmp_h0 D18) D18 {
						return func(_pat_let39_0 _dafny.Int) D18 {
							return func(_1413_dt__update_hcf79_h0 _dafny.Int) D18 {
								return Companion_D18_.Create_DC49_((_1412_dt__update__tmp_h0).Dtor_cf78(), _1413_dt__update_hcf79_h0)
							}(_pat_let39_0)
						}(_pat_let_tv33)
					}(_pat_let38_0)
				}(_1398_v15)))
				(_this).F4 = (p1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1409_v27, _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(829), _1398_v15), (func() _dafny.Map {
					if (_1411_v28).Contains(_this.F3) {
						return (_1411_v28).Get(_this.F3).(_dafny.Map)
					}
					return _1408_v26
				})()))).Cardinality())), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(_dafny.CodePoint)
			} else {
				var _1414_v29 _dafny.Sequence
				_ = _1414_v29
				_1414_v29 = _dafny.UnicodeSeqOfUtf8Bytes("sn")
				var _1415_v30 _dafny.Sequence
				_ = _1415_v30
				_1415_v30 = _dafny.SeqOf(p1, p1, p1, _1414_v29)
				_1414_v29 = _dafny.Companion_Sequence_.Concatenate((_1415_v30).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_dafny.SeqOf(_dafny.IntOfUint32((_1414_v29).Cardinality()), p0, p2, p0))).Cardinality(), _dafny.IntOfUint32((_1415_v30).Cardinality()))).Uint32()).(_dafny.Sequence), p1)
				(globalState).F0 = p0
				var _1416_v32 _dafny.Map
				_ = _1416_v32
				_1416_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("dtk")).Cardinality()))
				var _1417_v33 _dafny.Map
				_ = _1417_v33
				_1417_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1416_v32, p2)
				var _1418_v34 *C0
				_ = _1418_v34
				var _nw231 *C0 = New_C0_()
				_ = _nw231
				_nw231.Ctor__(func() _dafny.Map {
					var _coll54 = _dafny.NewMapBuilder()
					_ = _coll54
					for _iter58 := _dafny.Iterate((_1417_v33).Keys().Elements()); ; {
						_compr_54, _ok58 := _iter58()
						if !_ok58 {
							break
						}
						var _1419_v31 _dafny.Map
						_1419_v31 = interface{}(_compr_54).(_dafny.Map)
						if (_1417_v33).Contains(_1419_v31) {
							_coll54.Add(_1419_v31, !(_1380_v0))
						}
					}
					return _coll54.ToMap()
				}())
				_1418_v34 = _nw231
				var _index253 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1395_v12), 0))
				_ = _index253
				(_1395_v12).ArraySet1(_1380_v0, (_index253).Int())
				var _index254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1395_v12), 0))
				_ = _index254
				(_1395_v12).ArraySet1(false, (_index254).Int())
				var _rhs288 bool = (_1395_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1395_v12), 0))).Int()).(bool)
				_ = _rhs288
				var _rhs289 bool = _1380_v0
				_ = _rhs289
				var _rhs290 bool = (_1395_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(485), _dafny.ArrayLen((_1395_v12), 0))).Int()).(bool)
				_ = _rhs290
				var _rhs291 _dafny.CodePoint = _this.F4
				_ = _rhs291
				var _lhs246 *C6 = _this
				_ = _lhs246
				_1380_v0 = _rhs288
				_1380_v0 = _rhs289
				r1 = _rhs290
				_lhs246.F4 = _rhs291
			}
			var _1420_v35 _dafny.Map
			_ = _1420_v35
			_1420_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p2)
			var _1421_v36 _dafny.Map
			_ = _1421_v36
			_1421_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1420_v35, _1380_v0)
			var _1422_v37 *C0
			_ = _1422_v37
			var _nw232 *C0 = New_C0_()
			_ = _nw232
			_nw232.Ctor__(_1421_v36)
			_1422_v37 = _nw232
			var _1423_v38 _dafny.Set
			_ = _1423_v38
			_1423_v38 = _dafny.SetOf(_1422_v37)
			var _1424_v39 _dafny.Map
			_ = _1424_v39
			_1424_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).Cmp((_this.F3).Cardinality()) != 0, (_1423_v38).Cardinality())
			_1424_v39 = (_1424_v39).Update((_1422_v37).Fm10(_dafny.IntOfUint32((p1).Cardinality()), globalState), p0)
			var _1425_v40 _dafny.Array
			_ = _1425_v40
			var _nw233 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(17))
			_ = _nw233
			_1425_v40 = _nw233
			var _1426_v41 D1
			_ = _1426_v41
			_1426_v41 = Companion_D1_.Create_DC2_(_1380_v0)
			var _1427_v42 _dafny.MultiSet
			_ = _1427_v42
			_1427_v42 = _dafny.MultiSetOf(_1380_v0, _1380_v0, false)
			var _1428_v43 T1
			_ = _1428_v43
			var _nw234 *C3 = New_C3_()
			_ = _nw234
			_nw234.Ctor__((_1426_v41).Dtor_cf1(), _1427_v42)
			_1428_v43 = _nw234
			var _index255 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1425_v40), 0))
			_ = _index255
			(_1425_v40).ArraySet1(_1428_v43, (_index255).Int())
			var _index256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(758), _dafny.ArrayLen((_1425_v40), 0))
			_ = _index256
			(_1425_v40).ArraySet1(_1428_v43, (_index256).Int())
		} else {
			var _1429_v44 _dafny.Array
			_ = _1429_v44
			var _len0_40 _dafny.Int = _dafny.IntOfInt64(26)
			_ = _len0_40
			var _nw235 _dafny.Array
			_ = _nw235
			if _len0_40.Cmp(_dafny.Zero) == 0 {
				_nw235 = _dafny.NewArray(_len0_40)
			} else {
				var _init40 func(_dafny.Int) bool = (func(_1430_v0 bool) func(_dafny.Int) bool {
					return func(_1431_i2 _dafny.Int) bool {
						return _1430_v0
					}
				})(_1380_v0)
				_ = _init40
				var _element0_40 = _init40(_dafny.Zero)
				_ = _element0_40
				_nw235 = _dafny.NewArrayFromExample(_element0_40, nil, _len0_40)
				(_nw235).ArraySet1(_element0_40, 0)
				var _nativeLen0_40 = (_len0_40).Int()
				_ = _nativeLen0_40
				for _i0_40 := 1; _i0_40 < _nativeLen0_40; _i0_40++ {
					(_nw235).ArraySet1(_init40(_dafny.IntOf(_i0_40)), _i0_40)
				}
			}
			_1429_v44 = _nw235
			var _1432_v45 _dafny.Sequence
			_ = _1432_v45
			_1432_v45 = _dafny.SeqOf(p1, p1, p1)
			var _1433_v46 _dafny.Map
			_ = _1433_v46
			_1433_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, Companion_Default___.Fm35(_1432_v45, _1380_v0, globalState))
			var _1434_v47 _dafny.MultiSet
			_ = _1434_v47
			_1434_v47 = _dafny.MultiSetOf(true)
			var _1435_v48 *C5
			_ = _1435_v48
			var _nw236 *C5 = New_C5_()
			_ = _nw236
			_nw236.Ctor__((_dafny.Zero).Minus(p2), _1380_v0, _1429_v44, _1380_v0, !((func() bool {
				if (_1433_v46).Contains(_1380_v0) {
					return (_1433_v46).Get(_1380_v0).(bool)
				}
				return _1380_v0
			})()), (_1434_v47).Intersection(_1434_v47))
			_1435_v48 = _nw236
			var _rhs292 bool = (Companion_D8_.Create_DC25_(_1380_v0, (_1435_v48).F10())).Dtor_cf39()
			_ = _rhs292
			var _rhs293 _dafny.Int = p0
			_ = _rhs293
			_1380_v0 = _rhs292
			r2 = _rhs293
			if (_dafny.SetOf(_1429_v44, _1429_v44, _1429_v44)).IsDisjointFrom(_dafny.SetOf(_1429_v44, _1429_v44, _1429_v44)) {
				var _1436_v49 _dafny.Set
				_ = _1436_v49
				_1436_v49 = _dafny.SetOf((_1435_v48).F10(), _1380_v0)
				var _1437_v50 D2
				_ = _1437_v50
				_1437_v50 = Companion_D2_.Create_DC7_(_this.F4, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1436_v49).Cardinality(), p2), (_1435_v48).F9())
				var _1438_v51 _dafny.Map
				_ = _1438_v51
				_1438_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, p0)
				var _1439_v52 _dafny.Array
				_ = _1439_v52
				var _nwElement0_35 _dafny.Int = _dafny.IntOfInt64(563)
				_ = _nwElement0_35
				var _nw237 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(29))
				_ = _nw237
				(_nw237).ArraySet1(_nwElement0_35, 0)
				(_nw237).ArraySet1(_dafny.IntOfInt64(288), 1)
				(_nw237).ArraySet1(_dafny.IntOfInt64(712), 2)
				(_nw237).ArraySet1((_dafny.Zero).Minus((_1435_v48).F9()), 3)
				(_nw237).ArraySet1(Companion_Default___.SafeDivisionInt(p0, _dafny.IntOfInt64(523)), 4)
				(_nw237).ArraySet1(Companion_Default___.SafeModuloInt((_1435_v48).F9(), (_1435_v48).F9()), 5)
				(_nw237).ArraySet1(p2, 6)
				(_nw237).ArraySet1((_1437_v50).Dtor_cf8(), 7)
				(_nw237).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((Companion_Default___.Fm17(globalState)).Cardinality()), p2), 8)
				(_nw237).ArraySet1((_dafny.Zero).Minus((_1435_v48).F9()), 9)
				(_nw237).ArraySet1(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("pp")).Cardinality())), p2), 10)
				(_nw237).ArraySet1((func() _dafny.Int {
					if _1380_v0 {
						return p0
					}
					return (_1438_v51).Cardinality()
				})(), 11)
				(_nw237).ArraySet1((_1435_v48).F9(), 12)
				(_nw237).ArraySet1((_1435_v48).F9(), 13)
				(_nw237).ArraySet1((_1435_v48).F9(), 14)
				(_nw237).ArraySet1(p0, 15)
				(_nw237).ArraySet1(_dafny.IntOfInt64(-644), 16)
				(_nw237).ArraySet1(p2, 17)
				(_nw237).ArraySet1((_1435_v48).F9(), 18)
				(_nw237).ArraySet1((_1435_v48).F9(), 19)
				(_nw237).ArraySet1(_dafny.IntOfInt64(-83), 20)
				(_nw237).ArraySet1(((_1435_v48).F9()).Plus((_dafny.SetOf((_1435_v48).F9())).Cardinality()), 21)
				(_nw237).ArraySet1(p0, 22)
				(_nw237).ArraySet1(p2, 23)
				(_nw237).ArraySet1(p0, 24)
				(_nw237).ArraySet1((_1435_v48).F9(), 25)
				(_nw237).ArraySet1((_1435_v48).F9(), 26)
				(_nw237).ArraySet1((p0).Plus(p2), 27)
				(_nw237).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(77), p0), 28)
				_1439_v52 = _nw237
				var _index257 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))
				_ = _index257
				(_1439_v52).ArraySet1(p2, (_index257).Int())
				var _1440_v53 _dafny.Sequence
				_ = _1440_v53
				_1440_v53 = _dafny.SeqOf((_1435_v48).F10())
				var _1441_v54 _dafny.Map
				_ = _1441_v54
				_1441_v54 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1435_v48).F10(), _dafny.Companion_Sequence_.Concatenate(_1440_v53, _1440_v53))
				var _index258 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))
				_ = _index258
				var _rhs294 _dafny.Int = p0
				_ = _rhs294
				var _rhs295 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).Fm0(_1432_v45, (_1435_v48).F9(), _1380_v0, !(!((_1435_v48).F10())), globalState), _dafny.IntOfInt64(845))
				_ = _rhs295
				var _rhs296 _dafny.Map = (_1441_v54).Merge(_1441_v54)
				_ = _rhs296
				var _rhs297 bool = Companion_Default___.Fm35(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(p1, p1), _1432_v45), _1380_v0, globalState)
				_ = _rhs297
				var _lhs247 *GlobalState = globalState
				_ = _lhs247
				var _lhs248 _dafny.Array = _1439_v52
				_ = _lhs248
				var _lhs249 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))
				_ = _lhs249
				_lhs247.F0 = _rhs294
				(_lhs248).ArraySet1(_rhs295, (_lhs249).Int())
				_1441_v54 = _rhs296
				_1380_v0 = _rhs297
				var _1442_v55 _dafny.Map
				_ = _1442_v55
				_1442_v55 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("cbep"), _1380_v0)
				var _rhs298 _dafny.Int = (_1435_v48).Fm4((_1439_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))).Int()).(_dafny.Int), p2, _1442_v55, _dafny.IntOfUint32((p1).Cardinality()), globalState)
				_ = _rhs298
				var _rhs299 _dafny.Int = p2
				_ = _rhs299
				r2 = _rhs298
				r2 = _rhs299
				var _1443_v56 _dafny.Sequence
				_ = _1443_v56
				_1443_v56 = _dafny.UnicodeSeqOfUtf8Bytes("qwdwvuxp")
				var _1444_v57 T2
				_ = _1444_v57
				var _nw238 *C4 = New_C4_()
				_ = _nw238
				_nw238.Ctor__(_1429_v44, _1380_v0, _1380_v0, _1434_v47)
				_1444_v57 = _nw238
				var _1445_v58 _dafny.Map
				_ = _1445_v58
				_1445_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1444_v57, Companion_D7_.Create_DC21_(p2, (_1439_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))).Int()).(_dafny.Int)))
				var _1446_v59 _dafny.Map
				_ = _1446_v59
				_1446_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1445_v58, (_1439_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))).Int()).(_dafny.Int))
				var _1447_v60 _dafny.Sequence
				_ = _1447_v60
				_1447_v60 = _dafny.SeqOf((_1435_v48).F9())
				var _rhs300 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-717))).Uint32(), func(coer68 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg68 _dafny.Int) interface{} {
						return coer68(arg68)
					}
				}(func(_1448_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F4
				}))
				_ = _rhs300
				var _rhs301 bool = (_1440_v53).Select((Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_1447_v60)).Cardinality(), _dafny.IntOfUint32((_1440_v53).Cardinality()))).Uint32()).(bool)
				_ = _rhs301
				var _rhs302 _dafny.Map = (_1446_v59).Merge((_1446_v59).Merge(_1446_v59))
				_ = _rhs302
				_1443_v56 = _rhs300
				_1380_v0 = _rhs301
				_1446_v59 = _rhs302
				var _arr57 _dafny.Array = _1444_v57.F7()
				_ = _arr57
				var _index259 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))
				_ = _index259
				(_arr57).ArraySet1((func() bool {
					if !(_1444_v57.F8()) {
						return (_1435_v48).F10()
					}
					return (_1444_v57).F6()
				})(), (_index259).Int())
				var _arr58 _dafny.Array = _1444_v57.F7()
				_ = _arr58
				var _index260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))
				_ = _index260
				var _rhs303 bool = !((_1435_v48).F10())
				_ = _rhs303
				var _rhs304 _dafny.Int = (_1439_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(699), _dafny.ArrayLen((_1439_v52), 0))).Int()).(_dafny.Int)
				_ = _rhs304
				var _rhs305 bool = (_this.F3).IsSubsetOf((func() _dafny.Set {
					if false {
						return _dafny.SetOf(p0, (_1435_v48).F9())
					}
					return func() _dafny.Set {
						var _coll55 = _dafny.NewBuilder()
						_ = _coll55
						for _iter59 := _dafny.Iterate((_1447_v60).Elements()); ; {
							_compr_55, _ok59 := _iter59()
							if !_ok59 {
								break
							}
							var _1449_v61 _dafny.Int
							_1449_v61 = interface{}(_compr_55).(_dafny.Int)
							if _dafny.Companion_Sequence_.Contains(_1447_v60, _1449_v61) {
								_coll55.Add((_1449_v61).Times(_dafny.IntOfInt64(133)))
							}
						}
						return _coll55.ToSet()
					}()
				})())
				_ = _rhs305
				var _lhs250 _dafny.Array = _1444_v57.F7()
				_ = _lhs250
				var _lhs251 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))
				_ = _lhs251
				var _lhs252 T2 = _1444_v57
				_ = _lhs252
				(_lhs250).ArraySet1(_rhs303, (_lhs251).Int())
				r2 = _rhs304
				_lhs252.F8_set_(_rhs305)
				var _1450_v62 *C4
				_ = _1450_v62
				var _nw239 *C4 = New_C4_()
				_ = _nw239
				_nw239.Ctor__(_1429_v44, (_1444_v57.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))).Int()).(bool), (_1444_v57.F7()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))).Int()).(bool), _1434_v47)
				_1450_v62 = _nw239
				var _1451_v63 _dafny.Sequence
				_ = _1451_v63
				_1451_v63 = _dafny.SeqOf(_1450_v62, _1450_v62)
				var _1452_v64 _dafny.Map
				_ = _1452_v64
				_1452_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1444_v57).F6()), (_1451_v63).Select((Companion_Default___.SafeIndex(p2, _dafny.IntOfUint32((_1451_v63).Cardinality()))).Uint32()).(*C4))
				var _arr59 _dafny.Array = _1444_v57.F7()
				_ = _arr59
				var _index261 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))
				_ = _index261
				var _rhs306 bool = ((_1436_v49).Union(_1436_v49)).IsSubsetOf(_dafny.SetOf(_1380_v0, (_1444_v57).F6(), _1444_v57.F8(), !((_1435_v48).F10()), _1380_v0))
				_ = _rhs306
				var _rhs307 _dafny.Map = _1452_v64
				_ = _rhs307
				var _rhs308 _dafny.Int = (_1447_v60).Select((Companion_Default___.SafeIndex((_dafny.SetOf(_1439_v52)).Cardinality(), _dafny.IntOfUint32((_1447_v60).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs308
				var _lhs253 _dafny.Array = _1444_v57.F7()
				_ = _lhs253
				var _lhs254 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(697), _dafny.ArrayLen((_1444_v57.F7()), 0))
				_ = _lhs254
				(_lhs253).ArraySet1(_rhs306, (_lhs254).Int())
				_1452_v64 = _rhs307
				r0 = _rhs308
			} else {
				_1433_v46 = (_1433_v46).Merge(_1433_v46)
				var _index262 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1429_v44), 0))
				_ = _index262
				(_1429_v44).ArraySet1((_1435_v48).F10(), (_index262).Int())
				var _1453_v65 _dafny.Map
				_ = _1453_v65
				_1453_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, (_1435_v48).F9())
				var _1454_v66 D7
				_ = _1454_v66
				_1454_v66 = Companion_D7_.Create_DC21_((_1435_v48).F9(), (p0).Times((_1453_v65).Cardinality()))
				var _index263 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1429_v44), 0))
				_ = _index263
				var _rhs309 bool = (p2).Cmp((_1435_v48).F9()) >= 0
				_ = _rhs309
				var _rhs310 _dafny.Int = (_1435_v48).F9()
				_ = _rhs310
				var _rhs311 D7 = _1454_v66
				_ = _rhs311
				var _rhs312 _dafny.Int = (_dafny.Zero).Minus((_dafny.Zero).Minus(p2))
				_ = _rhs312
				var _rhs313 _dafny.Int = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("jv")), p1), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(228), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(p1, _dafny.UnicodeSeqOfUtf8Bytes("jv")), p1)).Cardinality()))).Uint32(), _this.F4)).Cardinality())
				_ = _rhs313
				var _lhs255 _dafny.Array = _1429_v44
				_ = _lhs255
				var _lhs256 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(620), _dafny.ArrayLen((_1429_v44), 0))
				_ = _lhs256
				var _lhs257 *GlobalState = globalState
				_ = _lhs257
				var _lhs258 *GlobalState = globalState
				_ = _lhs258
				(_lhs255).ArraySet1(_rhs309, (_lhs256).Int())
				r0 = _rhs310
				_1454_v66 = _rhs311
				_lhs257.F0 = _rhs312
				_lhs258.F0 = _rhs313
				var _1455_v67 _dafny.Array
				_ = _1455_v67
				var _len0_41 _dafny.Int = _dafny.IntOfInt64(26)
				_ = _len0_41
				var _nw240 _dafny.Array
				_ = _nw240
				if _len0_41.Cmp(_dafny.Zero) == 0 {
					_nw240 = _dafny.NewArray(_len0_41)
				} else {
					var _init41 func(_dafny.Int) bool = (func(_1456_v48 *C5) func(_dafny.Int) bool {
						return func(_1457_i4 _dafny.Int) bool {
							return (_1456_v48).F10()
						}
					})(_1435_v48)
					_ = _init41
					var _element0_41 = _init41(_dafny.Zero)
					_ = _element0_41
					_nw240 = _dafny.NewArrayFromExample(_element0_41, nil, _len0_41)
					(_nw240).ArraySet1(_element0_41, 0)
					var _nativeLen0_41 = (_len0_41).Int()
					_ = _nativeLen0_41
					for _i0_41 := 1; _i0_41 < _nativeLen0_41; _i0_41++ {
						(_nw240).ArraySet1(_init41(_dafny.IntOf(_i0_41)), _i0_41)
					}
				}
				_1455_v67 = _nw240
				var _1458_v68 _dafny.Array
				_ = _1458_v68
				var _nw241 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
				_ = _nw241
				_1458_v68 = _nw241
				var _1459_v69 _dafny.Map
				_ = _1459_v69
				_1459_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1458_v68, _1455_v67)
				var _1460_v70 _dafny.Sequence
				_ = _1460_v70
				_1460_v70 = _dafny.SeqOf(_1455_v67, (func() _dafny.Array {
					if (_1459_v69).Contains(_1458_v68) {
						return (_1459_v69).Get(_1458_v68).(_dafny.Array)
					}
					return _1455_v67
				})(), _1455_v67)
				var _1461_v71 _dafny.MultiSet
				_ = _1461_v71
				_1461_v71 = _dafny.MultiSetOf(_dafny.IntOfInt64(391))
				var _rhs314 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p2), (_1435_v48).F9())
				_ = _rhs314
				var _rhs315 bool = !((Companion_Default___.Fm12(globalState)).IsDisjointFrom((_1461_v71).Update(_dafny.IntOfInt64(336), Companion_Default___.Abs(p0))))
				_ = _rhs315
				var _rhs316 _dafny.Sequence = _1460_v70
				_ = _rhs316
				var _rhs317 bool = ((_1435_v48).F10()) || (!((_1435_v48).F10()))
				_ = _rhs317
				var _rhs318 _dafny.Int = (_1435_v48).F9()
				_ = _rhs318
				r2 = _rhs314
				_1380_v0 = _rhs315
				_1460_v70 = _rhs316
				_1380_v0 = _rhs317
				r2 = _rhs318
				var _1462_v72 D1
				_ = _1462_v72
				_1462_v72 = Companion_D1_.Create_DC2_((_1435_v48).F10())
				r1 = (_1462_v72).Dtor_cf1()
				r0 = Companion_Default___.SafeModuloInt((p0).Plus(p2), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(p1, p1)).Cardinality()))
			}
			var _1463_v73 _dafny.Sequence
			_ = _1463_v73
			_1463_v73 = _dafny.UnicodeSeqOfUtf8Bytes("fieugd")
			_1463_v73 = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(225))).Uint32(), func(coer69 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg69 _dafny.Int) interface{} {
					return coer69(arg69)
				}
			}(func(_1464_i5 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('d')
			})), Companion_Default___.Fm11(Companion_Default___.Fm35(_1432_v45, _1380_v0, globalState), globalState))
			(globalState).F0 = (_1435_v48).F9()
		}
		var _1465_v74 _dafny.Array
		_ = _1465_v74
		var _nw242 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw242
		_1465_v74 = _nw242
		for _iter60 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_1465_v74), 0))); ; {
			_guard_loop_4, _ok60 := _iter60()
			if !_ok60 {
				break
			}
			var _1466_i6 _dafny.Int
			_1466_i6 = interface{}(_guard_loop_4).(_dafny.Int)
			if (true) && (((_1466_i6).Sign() != -1) && ((_1466_i6).Cmp(_dafny.ArrayLen((_1465_v74), 0)) < 0)) {
				(_1465_v74).ArraySet1((Companion_D12_.Create_DC35_(_1380_v0)).Dtor_cf55(), (_1466_i6).Int())
			}
		}
		var _1467_i7 _dafny.Int
		_ = _1467_i7
		_1467_i7 = _dafny.Zero
		{
			for (p2).Cmp(Companion_Default___.SafeModuloInt(p2, p0)) >= 0 {
				{
					if (_1467_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L14
					}
					_1467_i7 = (_1467_i7).Plus(_dafny.One)
					var _1468_v75 _dafny.Sequence
					_ = _1468_v75
					_1468_v75 = _dafny.SeqOf(_1380_v0, _1380_v0, _1380_v0)
					var _1469_v76 _dafny.Set
					_ = _1469_v76
					_1469_v76 = _dafny.SetOf(_1380_v0, false)
					var _1470_v77 _dafny.Map
					_ = _1470_v77
					_1470_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1469_v76).Contains((_1468_v75).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1468_v75).Cardinality()))).Uint32()).(bool)), _1380_v0)
					r1 = !((func() bool {
						if (_1470_v77).Contains(_1380_v0) {
							return (_1470_v77).Get(_1380_v0).(bool)
						}
						return _1380_v0
					})())
					var _1471_v78 _dafny.Map
					_ = _1471_v78
					_1471_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p2)
					var _1472_v79 _dafny.Map
					_ = _1472_v79
					_1472_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1380_v0, p2)
					_1380_v0 = ((func() _dafny.Int {
						if (_1471_v78).Contains(p2) {
							return (_1471_v78).Get(p2).(_dafny.Int)
						}
						return p0
					})()).Cmp((p0).Minus((func() _dafny.Int {
						if (_1472_v79).Contains(_1380_v0) {
							return (_1472_v79).Get(_1380_v0).(_dafny.Int)
						}
						return p2
					})())) >= 0
					var _1473_v80 D2
					_ = _1473_v80
					_1473_v80 = Companion_D2_.Create_DC5_(_1380_v0)
					var _1474_v81 D2
					_ = _1474_v81
					_1474_v81 = Companion_D2_.Create_DC8_(_1473_v80)
					var _1475_v82 _dafny.Array
					_ = _1475_v82
					var _nwElement0_36 _dafny.Sequence = _1468_v75
					_ = _nwElement0_36
					var _nw243 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(5))
					_ = _nw243
					(_nw243).ArraySet1(_nwElement0_36, 0)
					(_nw243).ArraySet1(_dafny.SeqOf(_1380_v0), 1)
					(_nw243).ArraySet1(_1468_v75, 2)
					(_nw243).ArraySet1(_1468_v75, 3)
					(_nw243).ArraySet1(_1468_v75, 4)
					_1475_v82 = _nw243
					var _index264 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_1475_v82), 0))
					_ = _index264
					(_1475_v82).ArraySet1(_1468_v75, (_index264).Int())
					var _1476_v83 _dafny.Sequence
					_ = _1476_v83
					_1476_v83 = _dafny.UnicodeSeqOfUtf8Bytes("pgaotoy")
					var _index265 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_1475_v82), 0))
					_ = _index265
					var _rhs319 D2 = Companion_D2_.Create_DC8_(Companion_D2_.Create_DC5_(_1380_v0))
					_ = _rhs319
					var _rhs320 _dafny.Sequence = _1468_v75
					_ = _rhs320
					var _rhs321 _dafny.Sequence = _1476_v83
					_ = _rhs321
					var _lhs259 _dafny.Array = _1475_v82
					_ = _lhs259
					var _lhs260 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(126), _dafny.ArrayLen((_1475_v82), 0))
					_ = _lhs260
					_1474_v81 = _rhs319
					(_lhs259).ArraySet1(_rhs320, (_lhs260).Int())
					_1476_v83 = _rhs321
					var _1477_v84 _dafny.MultiSet
					_ = _1477_v84
					_1477_v84 = _dafny.MultiSetOf(_1380_v0, _1380_v0, _1380_v0)
					var _1478_v85 *C5
					_ = _1478_v85
					var _nw244 *C5 = New_C5_()
					_ = _nw244
					_nw244.Ctor__(p2, _1380_v0, _1465_v74, _1380_v0, _1380_v0, _1477_v84)
					_1478_v85 = _nw244
					var _1479_v86 _dafny.Map
					_ = _1479_v86
					_1479_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1478_v85, (_1478_v85).F9())
					(globalState).F0 = (p0).Plus((func() _dafny.Int {
						if (_1479_v86).Contains(_1478_v85) {
							return (_1479_v86).Get(_1478_v85).(_dafny.Int)
						}
						return _dafny.IntOfUint32((Companion_Default___.Fm23(p2, globalState)).Cardinality())
					})())
					goto C14
				}
			C14:
			}
			goto L14
		}
	L14:
		var _1480_v87 *C2
		_ = _1480_v87
		var _nw245 *C2 = New_C2_()
		_ = _nw245
		_nw245.Ctor__(_dafny.MultiSetOf(_1380_v0))
		_1480_v87 = _nw245
		var _1481_v88 _dafny.Sequence
		_ = _1481_v88
		_1481_v88 = _dafny.SeqOf(_1380_v0)
		var _1482_v89 _dafny.MultiSet
		_ = _1482_v89
		_1482_v89 = _dafny.MultiSetOf(_1380_v0, _1380_v0, _1380_v0, false)
		if (_1480_v87).Fm1(_1380_v0, (func() _dafny.Int {
			if _1380_v0 {
				return p2
			}
			return _dafny.IntOfUint32((_1481_v88).Cardinality())
		})(), (func() _dafny.Int {
			if (_1482_v89).Contains(_1380_v0) {
				return (_1482_v89).Multiplicity(_1380_v0)
			}
			return _dafny.IntOfInt64(979)
		})(), globalState) {
			var _1483_v90 _dafny.Map
			_ = _1483_v90
			_1483_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1380_v0, _1380_v0)
			var _1484_v91 _dafny.Sequence
			_ = _1484_v91
			_1484_v91 = _dafny.SeqOf((p2).Times((_1483_v90).Cardinality()), (_dafny.Zero).Minus(p0), p2)
			_1484_v91 = _dafny.Companion_Sequence_.Concatenate(_1484_v91, _1484_v91)
			var _1485_v92 _dafny.Map
			_ = _1485_v92
			_1485_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), p2)
			var _1486_v93 _dafny.MultiSet
			_ = _1486_v93
			_1486_v93 = _dafny.MultiSetOf(p2)
			_1485_v92 = (_1485_v92).Update(_1380_v0, (func() _dafny.Int {
				if (_1486_v93).Contains(p0) {
					return (_1486_v93).Multiplicity(p0)
				}
				return _dafny.IntOfInt64(369)
			})())
			var _1487_v94 _dafny.Sequence
			_ = _1487_v94
			_1487_v94 = _dafny.SeqOf(p1)
			var _1488_v95 _dafny.Map
			_ = _1488_v95
			_1488_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1465_v74, p0)
			var _1489_v96 _dafny.Array
			_ = _1489_v96
			var _nwElement0_37 _dafny.Int = _dafny.IntOfUint32((Companion_Default___.Fm18(globalState)).Cardinality())
			_ = _nwElement0_37
			var _nw246 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(21))
			_ = _nw246
			(_nw246).ArraySet1(_nwElement0_37, 0)
			(_nw246).ArraySet1((_1484_v91).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1484_v91).Cardinality()))).Uint32()).(_dafny.Int), 1)
			(_nw246).ArraySet1(p0, 2)
			(_nw246).ArraySet1(p2, 3)
			(_nw246).ArraySet1(p0, 4)
			(_nw246).ArraySet1(p0, 5)
			(_nw246).ArraySet1((func() _dafny.Int {
				if false {
					return p2
				}
				return p0
			})(), 6)
			(_nw246).ArraySet1((_this).Fm0(_1487_v94, (_this.F3).Cardinality(), _1380_v0, _1380_v0, globalState), 7)
			(_nw246).ArraySet1(p0, 8)
			(_nw246).ArraySet1(p2, 9)
			(_nw246).ArraySet1(p2, 10)
			(_nw246).ArraySet1((func() _dafny.Int {
				if _1380_v0 {
					return p0
				}
				return p2
			})(), 11)
			(_nw246).ArraySet1(p2, 12)
			(_nw246).ArraySet1(p0, 13)
			(_nw246).ArraySet1(p0, 14)
			(_nw246).ArraySet1((_1488_v95).Cardinality(), 15)
			(_nw246).ArraySet1(Companion_Default___.SafeModuloInt(p2, p2), 16)
			(_nw246).ArraySet1(p2, 17)
			(_nw246).ArraySet1(p2, 18)
			(_nw246).ArraySet1((func() _dafny.Int {
				if (_1482_v89).Contains(_1380_v0) {
					return (_1482_v89).Multiplicity(_1380_v0)
				}
				return p2
			})(), 19)
			(_nw246).ArraySet1((_this).Fm0(_dafny.SeqOf(p1), _dafny.IntOfInt64(-966), _1380_v0, _1380_v0, globalState), 20)
			_1489_v96 = _nw246
			_1489_v96 = _1489_v96
			var _1490_v97 _dafny.Map
			_ = _1490_v97
			_1490_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, p0)
			var _1491_v98 _dafny.Set
			_ = _1491_v98
			_1491_v98 = _dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1483_v90).Cardinality(), _dafny.IntOfInt64(-866))).Merge((_1490_v97).Update(p2, p0)))
			var _1492_v99 _dafny.Sequence
			_ = _1492_v99
			_1492_v99 = _dafny.UnicodeSeqOfUtf8Bytes("wxp")
			var _1493_v100 _dafny.Map
			_ = _1493_v100
			_1493_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1380_v0, _1490_v97)
			var _rhs322 _dafny.Set = (_1491_v98).Union((func() _dafny.Set {
				if !(Companion_Default___.Fm35(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(132))).Uint32(), func(coer70 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg70 _dafny.Int) interface{} {
						return coer70(arg70)
					}
				}((func(_1494_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1495_i8 _dafny.Int) _dafny.Sequence {
						return _1494_p1
					}
				})(p1))), !(_1380_v0), globalState)) {
					return _1491_v98
				}
				return _dafny.SetOf(_1490_v97, (func() _dafny.Map {
					if (_1493_v100).Contains(_1380_v0) {
						return (_1493_v100).Get(_1380_v0).(_dafny.Map)
					}
					return _1490_v97
				})())
			})())
			_ = _rhs322
			var _rhs323 _dafny.Sequence = _dafny.UnicodeSeqOfUtf8Bytes("yohtecotu")
			_ = _rhs323
			_1491_v98 = _rhs322
			_1492_v99 = _rhs323
			var _1496_v101 *C4
			_ = _1496_v101
			var _nw247 *C4 = New_C4_()
			_ = _nw247
			_nw247.Ctor__(_1465_v74, false, _1380_v0, _1482_v89)
			_1496_v101 = _nw247
			var _1497_v102 _dafny.Map
			_ = _1497_v102
			_1497_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1496_v101, _1486_v93)
			var _1498_v103 T2
			_ = _1498_v103
			var _nw248 *C4 = New_C4_()
			_ = _nw248
			_nw248.Ctor__(_1465_v74, _1380_v0, (_dafny.IntOfInt64(875)).Cmp((_1497_v102).Cardinality()) < 0, _1482_v89)
			_1498_v103 = _nw248
			var _rhs324 bool = _1498_v103.F8()
			_ = _rhs324
			var _rhs325 _dafny.Int = p0
			_ = _rhs325
			var _rhs326 T2 = _1498_v103
			_ = _rhs326
			var _rhs327 _dafny.Int = p0
			_ = _rhs327
			_1380_v0 = _rhs324
			r0 = _rhs325
			_1498_v103 = _rhs326
			r2 = _rhs327
		} else {
			var _1499_v104 _dafny.Array
			_ = _1499_v104
			var _len0_42 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_42
			var _nw249 _dafny.Array
			_ = _nw249
			if _len0_42.Cmp(_dafny.Zero) == 0 {
				_nw249 = _dafny.NewArray(_len0_42)
			} else {
				var _init42 func(_dafny.Int) _dafny.Sequence = (func(_1500_p1 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_1501_i9 _dafny.Int) _dafny.Sequence {
						return _1500_p1
					}
				})(p1)
				_ = _init42
				var _element0_42 = _init42(_dafny.Zero)
				_ = _element0_42
				_nw249 = _dafny.NewArrayFromExample(_element0_42, nil, _len0_42)
				(_nw249).ArraySet1(_element0_42, 0)
				var _nativeLen0_42 = (_len0_42).Int()
				_ = _nativeLen0_42
				for _i0_42 := 1; _i0_42 < _nativeLen0_42; _i0_42++ {
					(_nw249).ArraySet1(_init42(_dafny.IntOf(_i0_42)), _i0_42)
				}
			}
			_1499_v104 = _nw249
			var _index266 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1499_v104), 0))
			_ = _index266
			(_1499_v104).ArraySet1(p1, (_index266).Int())
			var _1502_v105 D5
			_ = _1502_v105
			_1502_v105 = Companion_D5_.Create_DC15_(p1)
			var _index267 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(507), _dafny.ArrayLen((_1499_v104), 0))
			_ = _index267
			(_1499_v104).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_1502_v105).Dtor_cf18(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(72))).Uint32(), func(coer71 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg71 _dafny.Int) interface{} {
					return coer71(arg71)
				}
			}(func(_1503_i10 _dafny.Int) _dafny.CodePoint {
				return _this.F4
			}))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(373))).Uint32(), func(coer72 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg72 _dafny.Int) interface{} {
					return coer72(arg72)
				}
			}(func(_1504_i11 _dafny.Int) _dafny.CodePoint {
				return _this.F4
			}))), (_index267).Int())
			_1380_v0 = _1380_v0
			_1380_v0 = _1380_v0
			var _1505_v106 _dafny.Array
			_ = _1505_v106
			var _len0_43 _dafny.Int = _dafny.IntOfInt64(25)
			_ = _len0_43
			var _nw250 _dafny.Array
			_ = _nw250
			if _len0_43.Cmp(_dafny.Zero) == 0 {
				_nw250 = _dafny.NewArray(_len0_43)
			} else {
				var _init43 func(_dafny.Int) _dafny.Int = (func(_1506_v0 bool) func(_dafny.Int) _dafny.Int {
					return func(_1507_i12 _dafny.Int) _dafny.Int {
						return (_1507_i12).Times((_dafny.SetOf(_1506_v0, _1506_v0)).Cardinality())
					}
				})(_1380_v0)
				_ = _init43
				var _element0_43 = _init43(_dafny.Zero)
				_ = _element0_43
				_nw250 = _dafny.NewArrayFromExample(_element0_43, nil, _len0_43)
				(_nw250).ArraySet1(_element0_43, 0)
				var _nativeLen0_43 = (_len0_43).Int()
				_ = _nativeLen0_43
				for _i0_43 := 1; _i0_43 < _nativeLen0_43; _i0_43++ {
					(_nw250).ArraySet1(_init43(_dafny.IntOf(_i0_43)), _i0_43)
				}
			}
			_1505_v106 = _nw250
			var _index268 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1505_v106), 0))
			_ = _index268
			(_1505_v106).ArraySet1(p2, (_index268).Int())
			var _1508_v107 _dafny.Sequence
			_ = _1508_v107
			_1508_v107 = _dafny.SeqOf(_1481_v88, _dafny.SeqOf(true), _1481_v88, _1481_v88, _1481_v88)
			var _1509_v108 *C5
			_ = _1509_v108
			var _nw251 *C5 = New_C5_()
			_ = _nw251
			_nw251.Ctor__(p0, _1380_v0, _1465_v74, false, _1380_v0, (_1482_v89).Update(_1380_v0, Companion_Default___.Abs(p2)))
			_1509_v108 = _nw251
			var _1510_v109 _dafny.Map
			_ = _1510_v109
			_1510_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1509_v108, (_1509_v108).F10())
			var _1511_v110 _dafny.MultiSet
			_ = _1511_v110
			_1511_v110 = _dafny.MultiSetOf(_dafny.SeqOf((func() bool {
				if (_1510_v109).Contains(_1509_v108) {
					return (_1510_v109).Get(_1509_v108).(bool)
				}
				return !(!(!((_1509_v108).F10())))
			})(), _1380_v0), _1481_v88)
			var _index269 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.ArrayLen((_1505_v106), 0))
			_ = _index269
			(_1505_v106).ArraySet1(((_dafny.MultiSetFromSeq(_1508_v107)).Difference(_1511_v110)).Cardinality(), (_index269).Int())
			r1 = !(!(false) || (!(_1380_v0)))
		}
		var _1512_v111 D17
		_ = _1512_v111
		_1512_v111 = Companion_D17_.Create_DC47_(_1380_v0, p0, p2)
		var _pat_let_tv34 = p0
		_ = _pat_let_tv34
		r0 = (func(_pat_let40_0 D17) D17 {
			return func(_1513_dt__update__tmp_h1 D17) D17 {
				return func(_pat_let41_0 _dafny.Int) D17 {
					return func(_1514_dt__update_hcf75_h0 _dafny.Int) D17 {
						return Companion_D17_.Create_DC47_((_1513_dt__update__tmp_h1).Dtor_cf74(), _1514_dt__update_hcf75_h0, (_1513_dt__update__tmp_h1).Dtor_cf76())
					}(_pat_let41_0)
				}(_pat_let_tv34)
			}(_pat_let40_0)
		}(_1512_v111)).Dtor_cf75()
		var _1515_v112 _dafny.MultiSet
		_ = _1515_v112
		_1515_v112 = _dafny.MultiSetOf(p0)
		var _1516_v113 _dafny.Map
		_ = _1516_v113
		_1516_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p2, _1380_v0)
		var _1517_v114 _dafny.Sequence
		_ = _1517_v114
		_1517_v114 = _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("uksqkb"), p1)
		r1 = !((func() bool {
			if (_1516_v113).Contains((_this).Fm0(_1517_v114, _dafny.IntOfUint32((_dafny.SeqOf(p0, p0, p0)).Cardinality()), _1380_v0, _1380_v0, globalState)) {
				return (_1516_v113).Get((_this).Fm0(_1517_v114, _dafny.IntOfUint32((_dafny.SeqOf(p0, p0, p0)).Cardinality()), _1380_v0, _1380_v0, globalState)).(bool)
			}
			return _1380_v0
		})()) || ((_dafny.MultiSetOf(p2)).IsDisjointFrom(_1515_v112))
		r2 = (_dafny.Zero).Minus(p2)
		return r0, r1, r2
	}
}

// End of class C6
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
