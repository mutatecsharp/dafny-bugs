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
func (_static *CompanionStruct_Default___) Fm0(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Set {
	return ((func() _dafny.Set {
		if true {
			return _dafny.SetOf(true, false)
		}
		return _dafny.SetOf(false)
	})()).Union(_dafny.SetOf(!(true), true, false))
}
func (_static *CompanionStruct_Default___) Fm3(globalState *GlobalState) bool {
	return !(!((((func() _dafny.Set {
		var _coll0 = _dafny.NewBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate((_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(true, false))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(679), _dafny.IntOfInt64(-336))).Cardinality(), !(false))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(793), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(8))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))).Cardinality(), _dafny.IntOfInt64(76)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())))).Elements()); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _0_v0 _dafny.Map
			_0_v0 = interface{}(_compr_0).(_dafny.Map)
			if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.MultiSetOf(false), _dafny.MultiSetOf(true, false))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(!(false))).Cardinality())), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("l")).Cardinality()), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(_dafny.IntOfInt64(679), _dafny.IntOfInt64(-336))).Cardinality(), !(false))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(793), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(8))).Cardinality()), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq(_dafny.SeqOf(true, false))).Cardinality(), _dafny.IntOfInt64(76)), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(true)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()))), _0_v0) {
				_coll0.Add(_0_v0)
			}
		}
		return _coll0.ToSet()
	}()).Cardinality()).Plus((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Cardinality())).Cmp((_dafny.MultiSetOf(_dafny.IntOfInt64(-574), _dafny.IntOfUint32((_dafny.SeqOf(false, !(true))).Cardinality()), _dafny.IntOfInt64(-776))).Cardinality()) == 0))
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Int, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf((_dafny.IntOfInt64(402)).Cmp((_dafny.MultiSetOf(false)).Cardinality()) > 0, true, !((_dafny.MultiSetOf(!(!(false)), true, true)).Contains(false)), true, false)
}
func (_static *CompanionStruct_Default___) Fm7(p0 _dafny.CodePoint, p1 _dafny.Int, p2 bool, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm8(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(544), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("agf")).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ufyke")).Cardinality())))).Cardinality())
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.IntOfInt64(324), _dafny.IntOfInt64(155))
}
func (_static *CompanionStruct_Default___) Fm11(p0 D3, p1 bool, p2 _dafny.Int, p3 bool, globalState *GlobalState) D2 {
	return Companion_D2_.Create_DC7_(false)
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Sequence, p2 bool, globalState *GlobalState) D3 {
	if false {
		return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_(false, _dafny.IntOfInt64(707), _dafny.SetOf(false))))))
	} else {
		return Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_(false, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-504))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg0 _dafny.Int) interface{} {
				return coer0(arg0)
			}
		}(func(_1_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('h')
		}))).Cardinality()), _dafny.SetOf(false)))
	}
}
func (_static *CompanionStruct_Default___) Fm13(globalState *GlobalState) _dafny.CodePoint {
	if (!(false)) || (false) {
		if true {
			return _dafny.CodePoint('y')
		} else {
			return _dafny.CodePoint('x')
		}
	} else if !(!(true)) {
		return _dafny.CodePoint('f')
	} else {
		return _dafny.CodePoint('o')
	}
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.Int, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("nxqveefwj")
}
func (_static *CompanionStruct_Default___) Fm15(globalState *GlobalState) _dafny.Map {
	return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqOf(!(true)))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xw")).Cardinality()), (_dafny.SetOf(false)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm16(globalState *GlobalState) _dafny.Sequence {
	return _dafny.SeqOf(Companion_D1_.Create_DC5_(_dafny.IntOfInt64(710)), Companion_D1_.Create_DC5_(_dafny.IntOfInt64(611)), Companion_D1_.Create_DC5_(_dafny.IntOfInt64(-796)))
}
func (_static *CompanionStruct_Default___) Fm17(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	if (_dafny.SetOf(false, true, !(false))).Contains(false) {
		return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('q'), false)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('k'), true))
	} else {
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('d'), true)
	}
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.Int, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, false, true, false), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(true, (Companion_D12_.Create_DC34_(!(true))).Dtor_cf65(), true), _dafny.SeqOf(false)))
}
func (_static *CompanionStruct_Default___) Fm21(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_((_dafny.IntOfInt64(687)).Plus(_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())), (_dafny.IntOfInt64(818)).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-321))).Uint32(), func(coer1 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_2_i0 _dafny.Int) _dafny.Sequence {
		return _dafny.UnicodeSeqOfUtf8Bytes("qyetl")
	}))).Cardinality())) <= 0, (func() _dafny.Sequence {
		if true {
			return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(614))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg2 _dafny.Int) interface{} {
					return coer2(arg2)
				}
			}(func(_3_i1 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(8)
			}))
		}
		return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(818))).Uint32(), func(coer3 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_4_i2 _dafny.Int) _dafny.Int {
			return _dafny.IntOfInt64(-302)
		}))
	})())
}
func (_static *CompanionStruct_Default___) Fm22(globalState *GlobalState) _dafny.Set {
	return (func() _dafny.Set {
		var _coll1 = _dafny.NewBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((func() _dafny.Map {
			var _coll2 = _dafny.NewMapBuilder()
			_ = _coll2
			for _iter2 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-951))).Uint32(), func(coer4 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg4 _dafny.Int) interface{} {
					return coer4(arg4)
				}
			}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('v')
			}))).Cardinality()), _dafny.IntOfInt64(168)))).Cardinality())).Keys().Elements()); ; {
				_compr_2, _ok2 := _iter2()
				if !_ok2 {
					break
				}
				var _6_v0 _dafny.Map
				_6_v0 = interface{}(_compr_2).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-951))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg5 _dafny.Int) interface{} {
						return coer5(arg5)
					}
				}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()), _dafny.IntOfInt64(168)))).Cardinality())).Contains(_6_v0) {
					_coll2.Add(_6_v0, !(true))
				}
			}
			return _coll2.ToMap()
		}()).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _7_v1 _dafny.Map
			_7_v1 = interface{}(_compr_1).(_dafny.Map)
			if (func() _dafny.Map {
				var _coll3 = _dafny.NewMapBuilder()
				_ = _coll3
				for _iter3 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-951))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('v')
				}))).Cardinality()), _dafny.IntOfInt64(168)))).Cardinality())).Keys().Elements()); ; {
					_compr_3, _ok3 := _iter3()
					if !_ok3 {
						break
					}
					var _6_v0 _dafny.Map
					_6_v0 = interface{}(_compr_3).(_dafny.Map)
					if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-951))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg7 _dafny.Int) interface{} {
							return coer7(arg7)
						}
					}(func(_5_i0 _dafny.Int) _dafny.CodePoint {
						return _dafny.CodePoint('v')
					}))).Cardinality()), _dafny.IntOfInt64(168)))).Cardinality())).Contains(_6_v0) {
						_coll3.Add(_6_v0, !(true))
					}
				}
				return _coll3.ToMap()
			}()).Contains(_7_v1) {
				_coll1.Add(_7_v1)
			}
		}
		return _coll1.ToSet()
	}()).Union((func() _dafny.Set {
		if true {
			return _dafny.SetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, true))
		}
		return func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true), (_dafny.Zero).Minus(_dafny.IntOfInt64(-400)))).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _8_v2 _dafny.Map
				_8_v2 = interface{}(_compr_4).(_dafny.Map)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), true), (_dafny.Zero).Minus(_dafny.IntOfInt64(-400)))).Contains(_8_v2) {
					_coll4.Add(_8_v2)
				}
			}
			return _coll4.ToSet()
		}()
	})())
}
func (_static *CompanionStruct_Default___) Fm23(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.CodePoint, p3 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if true {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)
	})()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)))
}
func (_static *CompanionStruct_Default___) Fm24(p0 _dafny.Int, p1 bool, p2 bool, p3 _dafny.Int, globalState *GlobalState) _dafny.Map {
	return ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.Zero).Minus(_dafny.IntOfInt64(-285)))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!(!(false))), _dafny.IntOfInt64(-46)))).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_dafny.MultiSetOf(_dafny.IntOfInt64(289))).Cardinality())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfInt64(-844))))
}
func (_static *CompanionStruct_Default___) Fm25(globalState *GlobalState) D1 {
	var _source0 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(381))).Uint32(), func(coer8 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
		return func(arg8 _dafny.Int) interface{} {
			return coer8(arg8)
		}
	}(func(_9_i0 _dafny.Int) D1 {
		return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(928))
	}))
	_ = _source0
	var _10___mcc_h0 _dafny.Sequence = _source0
	_ = _10___mcc_h0
	var _11_cf24 _dafny.Sequence = _10___mcc_h0
	_ = _11_cf24
	return Companion_D1_.Create_DC5_(_dafny.IntOfInt64(752))
}
func (_static *CompanionStruct_Default___) Fm26(p0 bool, p1 _dafny.Int, p2 _dafny.CodePoint, globalState *GlobalState) D1 {
	return Companion_D1_.Create_DC4_(_dafny.CodePoint('a'), (_dafny.Zero).Minus((_dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality())).Times((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bfmrgm")).Cardinality())))), _dafny.CodePoint('q'), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_dafny.SetOf(false)).Cardinality())).Cardinality()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-71))).Uint32(), func(coer9 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg9 _dafny.Int) interface{} {
			return coer9(arg9)
		}
	}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('a')
	})))).Cardinality()).Times(_dafny.IntOfInt64(961)), (_dafny.MultiSetOf(!(true), true)).IsSubsetOf(_dafny.MultiSetFromSeq(_dafny.SeqOf(true, true))))
}
func (_static *CompanionStruct_Default___) Fm27(p0 _dafny.Int, p1 bool, p2 bool, globalState *GlobalState) bool {
	return (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nxgruga")).Cardinality())).Cmp(_dafny.IntOfUint32((_dafny.SeqOf(Companion_D14_.Create_DC40_((_dafny.Zero).Minus(_dafny.IntOfInt64(-75)), true, _dafny.UnicodeSeqOfUtf8Bytes("dnsuin")))).Cardinality())) == 0
}
func (_static *CompanionStruct_Default___) Fm28(p0 _dafny.Int, p1 bool, p2 _dafny.Sequence, p3 D3, globalState *GlobalState) _dafny.Sequence {
	return _dafny.UnicodeSeqOfUtf8Bytes("pbthlure")
}
func (_static *CompanionStruct_Default___) Fm29(globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false), _dafny.SeqOf((Companion_D0_.Create_DC2_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, Companion_D12_.Create_DC33_(false, _dafny.UnicodeSeqOfUtf8Bytes("yuplewtm")))).Cardinality(), true, _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bensvahfw")).Cardinality())))).Dtor_cf5())), _dafny.SeqOf(true, true, false))
}
func (_static *CompanionStruct_Default___) Fm30(p0 bool, globalState *GlobalState) _dafny.Set {
	return ((Companion_D16_.Create_DC45_(_dafny.SetOf((_dafny.Zero).Minus((_dafny.MultiSetOf(false)).Cardinality()), _dafny.IntOfInt64(292)))).Dtor_cf80()).Difference(func() _dafny.Set {
		var _coll5 = _dafny.NewBuilder()
		_ = _coll5
		for _iter5 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(363), _dafny.IntOfInt64(186))); ; {
			_compr_5, _ok5 := _iter5()
			if !_ok5 {
				break
			}
			var _13_v0 _dafny.Int
			_13_v0 = interface{}(_compr_5).(_dafny.Int)
			if ((_dafny.IntOfInt64(363)).Cmp(_13_v0) <= 0) && ((_13_v0).Cmp(_dafny.IntOfInt64(186)) < 0) {
				_coll5.Add((_13_v0).Times(_dafny.IntOfInt64(-949)))
			}
		}
		return _coll5.ToSet()
	}())
}
func (_static *CompanionStruct_Default___) Fm31(globalState *GlobalState) D6 {
	return Companion_D6_.Create_DC17_(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(992))).Uint32(), func(coer10 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg10 _dafny.Int) interface{} {
			return coer10(arg10)
		}
	}(func(_14_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('d')
	})), _dafny.UnicodeSeqOfUtf8Bytes("eanbpr")), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(156))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg11 _dafny.Int) interface{} {
			return coer11(arg11)
		}
	}(func(_15_i1 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('s')
	})), _dafny.UnicodeSeqOfUtf8Bytes("xm")))
}
func (_static *CompanionStruct_Default___) Fm32(p0 _dafny.Sequence, p1 bool, globalState *GlobalState) _dafny.Map {
	return ((func() _dafny.Map {
		if false {
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), (_dafny.MultiSetOf((_dafny.SetOf(!(false))).Cardinality())).Cardinality())).Cardinality()), _dafny.UnicodeSeqOfUtf8Bytes("sorttx"))
		}
		return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_dafny.SetOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(599))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg12 _dafny.Int) interface{} {
				return coer12(arg12)
			}
		}(func(_16_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('v')
		})))).Cardinality(), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rkbghw")).Cardinality()))).Cardinality(), _dafny.IntOfInt64(559)), _dafny.UnicodeSeqOfUtf8Bytes("yhnnprjd"))
	})()).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf(_dafny.IntOfUint32((_dafny.SeqOf(!(true), true)).Cardinality()), _dafny.IntOfInt64(67)), _dafny.UnicodeSeqOfUtf8Bytes("ug")))
}
func (_static *CompanionStruct_Default___) Fm33(globalState *GlobalState) _dafny.MultiSet {
	return _dafny.MultiSetOf(((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, false)).Cardinality())).Minus(_dafny.IntOfInt64(77)))
}
func (_static *CompanionStruct_Default___) Fm34(p0 bool, p1 D8, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	if !(_dafny.SetOf(_dafny.IntOfInt64(919), (_dafny.SetOf(!(false))).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf(true)).Cardinality()), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sswah")).Cardinality()))))).Contains(_dafny.IntOfInt64(788)) {
		return _dafny.SeqOf(_dafny.SeqOf(true, true, false))
	} else {
		return _dafny.SeqOf(_dafny.SeqOf(true, true))
	}
}
func (_static *CompanionStruct_Default___) M0(globalState *GlobalState) {
	var _17_v0 bool
	_ = _17_v0
	_17_v0 = true
	if _17_v0 {
		var _18_v1 _dafny.Int
		_ = _18_v1
		_18_v1 = _dafny.IntOfInt64(146)
		(globalState).F1 = (_dafny.Zero).Minus(_18_v1)
		var _19_v2 _dafny.MultiSet
		_ = _19_v2
		_19_v2 = _dafny.MultiSetOf(_18_v1)
		var _20_v3 _dafny.Array
		_ = _20_v3
		var _nw0 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(19))
		_ = _nw0
		_20_v3 = _nw0
		var _21_v4 _dafny.Map
		_ = _21_v4
		_21_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_18_v1).Times(_18_v1), _20_v3)
		var _22_v5 _dafny.Sequence
		_ = _22_v5
		_22_v5 = _dafny.SeqOf(Companion_Default___.Fm13(globalState))
		var _23_v6 _dafny.CodePoint
		_ = _23_v6
		_23_v6 = _dafny.CodePoint('f')
		var _24_v7 T0
		_ = _24_v7
		var _nw1 *C4 = New_C4_()
		_ = _nw1
		_nw1.Ctor__(_dafny.IntOfUint32((_22_v5).Cardinality()), _17_v0, _17_v0, _17_v0, _23_v6, _22_v5)
		_24_v7 = _nw1
		var _25_v8 _dafny.Map
		_ = _25_v8
		_25_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_22_v5).Cardinality()), _24_v7)
		var _rhs0 _dafny.Int = _18_v1
		_ = _rhs0
		var _rhs1 _dafny.MultiSet = (_19_v2).Update(_dafny.IntOfInt64(82), Companion_Default___.Abs(_dafny.IntOfInt64(397)))
		_ = _rhs1
		var _rhs2 _dafny.Map = (_21_v4).Merge(_21_v4)
		_ = _rhs2
		var _rhs3 _dafny.Int = (func() _dafny.Int {
			if (_19_v2).Contains(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_22_v5).Cardinality()), _18_v1)) {
				return (_19_v2).Multiplicity(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_22_v5).Cardinality()), _18_v1))
			}
			return _18_v1
		})()
		_ = _rhs3
		var _rhs4 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_dafny.SeqOf((_25_v8).Cardinality())).Cardinality()), Companion_Default___.SafeModuloInt(_18_v1, _18_v1))
		_ = _rhs4
		var _lhs0 *GlobalState = globalState
		_ = _lhs0
		var _lhs1 *GlobalState = globalState
		_ = _lhs1
		_18_v1 = _rhs0
		_19_v2 = _rhs1
		_21_v4 = _rhs2
		_lhs0.F1 = _rhs3
		_lhs1.F1 = _rhs4
		if !(_17_v0) {
			(globalState).F1 = (_18_v1).Times(_18_v1)
			var _26_v9 D6
			_ = _26_v9
			_26_v9 = Companion_D6_.Create_DC16_(_19_v2)
			var _27_v10 _dafny.Sequence
			_ = _27_v10
			_27_v10 = _dafny.SeqOf(_26_v9, _26_v9, _26_v9)
			var _28_v11 _dafny.Sequence
			_ = _28_v11
			_28_v11 = _dafny.SeqOf(_17_v0, true, _17_v0)
			_27_v10 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_27_v10, (Companion_Default___.SafeIndex((_dafny.MultiSetFromSeq(_28_v11)).Cardinality(), _dafny.IntOfUint32((_27_v10).Cardinality()))).Uint32(), _26_v9), _27_v10)
			_22_v5 = (func() _dafny.Sequence {
				if _17_v0 {
					return (_24_v7).F5()
				}
				return (_24_v7).F5()
			})()
			var _29_v12 _dafny.Array
			_ = _29_v12
			var _nw2 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(18))
			_ = _nw2
			_29_v12 = _nw2
			var _30_v13 _dafny.MultiSet
			_ = _30_v13
			_30_v13 = _dafny.MultiSetOf(_17_v0)
			var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_29_v12), 0))
			_ = _index0
			(_29_v12).ArraySet1(_30_v13, (_index0).Int())
			var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(430), _dafny.ArrayLen((_29_v12), 0))
			_ = _index1
			(_29_v12).ArraySet1(_30_v13, (_index1).Int())
			var _31_v14 *C1
			_ = _31_v14
			var _nw3 *C1 = New_C1_()
			_ = _nw3
			_nw3.Ctor__(_18_v1, _17_v0, _17_v0, _17_v0, _23_v6, _22_v5)
			_31_v14 = _nw3
			var _32_v15 _dafny.Sequence
			_ = _32_v15
			_32_v15 = _dafny.SeqOf(_31_v14, _31_v14, _31_v14, _31_v14)
			var _33_v16 _dafny.Map
			_ = _33_v16
			_33_v16 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_32_v15, _18_v1)
			var _34_v17 _dafny.Map
			_ = _34_v17
			_34_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_33_v16).Contains(_dafny.SeqOf(_31_v14)) {
					return (_33_v16).Get(_dafny.SeqOf(_31_v14)).(_dafny.Int)
				}
				return _18_v1
			})(), _18_v1)
			var _35_v18 _dafny.Sequence
			_ = _35_v18
			_35_v18 = _dafny.SeqOf(_18_v1)
			_34_v17 = (_34_v17).Update(_18_v1, _dafny.IntOfUint32((_35_v18).Cardinality()))
		} else {
			var _36_v20 _dafny.Array
			_ = _36_v20
			var _len0_0 _dafny.Int = _dafny.IntOfInt64(18)
			_ = _len0_0
			var _nw4 _dafny.Array
			_ = _nw4
			if _len0_0.Cmp(_dafny.Zero) == 0 {
				_nw4 = _dafny.NewArray(_len0_0)
			} else {
				var _init0 func(_dafny.Int) _dafny.Int = func(_37_i0 _dafny.Int) _dafny.Int {
					return Companion_Default___.SafeDivisionInt(_37_i0, (func() _dafny.Set {
						var _coll6 = _dafny.NewBuilder()
						_ = _coll6
						for _iter6 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-674), _dafny.IntOfInt64(-208))); ; {
							_compr_6, _ok6 := _iter6()
							if !_ok6 {
								break
							}
							var _38_v19 _dafny.Int
							_38_v19 = interface{}(_compr_6).(_dafny.Int)
							if ((_dafny.IntOfInt64(-674)).Cmp(_38_v19) <= 0) && ((_38_v19).Cmp(_dafny.IntOfInt64(-208)) < 0) {
								_coll6.Add((_38_v19).Plus(_dafny.IntOfInt64(382)))
							}
						}
						return _coll6.ToSet()
					}()).Cardinality())
				}
				_ = _init0
				var _element0_0 = _init0(_dafny.Zero)
				_ = _element0_0
				_nw4 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
				(_nw4).ArraySet1(_element0_0, 0)
				var _nativeLen0_0 = (_len0_0).Int()
				_ = _nativeLen0_0
				for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
					(_nw4).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
				}
			}
			_36_v20 = _nw4
			var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_36_v20), 0))
			_ = _index2
			(_36_v20).ArraySet1((func() _dafny.Int {
				if _17_v0 {
					return _18_v1
				}
				return _18_v1
			})(), (_index2).Int())
			var _39_v21 _dafny.Array
			_ = _39_v21
			var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(8))
			_ = _nw5
			_39_v21 = _nw5
			var _40_v22 _dafny.Array
			_ = _40_v22
			var _nw6 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(17))
			_ = _nw6
			_40_v22 = _nw6
			var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_40_v22), 0))
			_ = _index3
			(_40_v22).ArraySet1((_19_v2).Equals(_19_v2), (_index3).Int())
			var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_36_v20), 0))
			_ = _index4
			var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_40_v22), 0))
			_ = _index5
			var _rhs5 _dafny.CodePoint = _24_v7.F4()
			_ = _rhs5
			var _rhs6 _dafny.Int = _dafny.IntOfInt64(-977)
			_ = _rhs6
			var _rhs7 _dafny.Array = (func() _dafny.Array {
				if (func() bool {
					if _17_v0 {
						return Companion_Default___.Fm7(_24_v7.F4(), _18_v1, _17_v0, globalState)
					}
					return _17_v0
				})() {
					return _39_v21
				}
				return _39_v21
			})()
			_ = _rhs7
			var _rhs8 bool = !(false) || (!(!(_17_v0) || (_17_v0)))
			_ = _rhs8
			var _lhs2 _dafny.Array = _36_v20
			_ = _lhs2
			var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(243), _dafny.ArrayLen((_36_v20), 0))
			_ = _lhs3
			var _lhs4 _dafny.Array = _40_v22
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(657), _dafny.ArrayLen((_40_v22), 0))
			_ = _lhs5
			_23_v6 = _rhs5
			(_lhs2).ArraySet1(_rhs6, (_lhs3).Int())
			_39_v21 = _rhs7
			(_lhs4).ArraySet1(_rhs8, (_lhs5).Int())
			_23_v6 = _24_v7.F4()
			_17_v0 = _17_v0
			var _41_v23 _dafny.Array
			_ = _41_v23
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(3)
			_ = _len0_1
			var _nw7 _dafny.Array
			_ = _nw7
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw7 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Map = (func(_42_v7 T0) func(_dafny.Int) _dafny.Map {
					return func(_43_i1 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("itjlfvghe"), _42_v7.F4())
					}
				})(_24_v7)
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw7 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw7).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw7).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_41_v23 = _nw7
			var _44_v24 _dafny.Map
			_ = _44_v24
			_44_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("dtgg"), _24_v7.F4())
			var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_41_v23), 0))
			_ = _index6
			(_41_v23).ArraySet1((_44_v24).Merge(_44_v24), (_index6).Int())
			var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen((_41_v23), 0))
			_ = _index7
			(_41_v23).ArraySet1(_44_v24, (_index7).Int())
			_40_v22 = _40_v22
		}
		_17_v0 = Companion_Default___.Fm3(globalState)
		(globalState).F1 = (_18_v1).Minus(_18_v1)
	} else {
		_17_v0 = _17_v0
		var _45_v25 _dafny.Array
		_ = _45_v25
		var _len0_2 _dafny.Int = _dafny.IntOfInt64(11)
		_ = _len0_2
		var _nw8 _dafny.Array
		_ = _nw8
		if _len0_2.Cmp(_dafny.Zero) == 0 {
			_nw8 = _dafny.NewArray(_len0_2)
		} else {
			var _init2 func(_dafny.Int) bool = (func(_46_v0 bool) func(_dafny.Int) bool {
				return func(_47_i2 _dafny.Int) bool {
					return _46_v0
				}
			})(_17_v0)
			_ = _init2
			var _element0_2 = _init2(_dafny.Zero)
			_ = _element0_2
			_nw8 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
			(_nw8).ArraySet1(_element0_2, 0)
			var _nativeLen0_2 = (_len0_2).Int()
			_ = _nativeLen0_2
			for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
				(_nw8).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
			}
		}
		_45_v25 = _nw8
		var _48_v26 _dafny.Array
		_ = _48_v26
		var _len0_3 _dafny.Int = _dafny.IntOfInt64(6)
		_ = _len0_3
		var _nw9 _dafny.Array
		_ = _nw9
		if _len0_3.Cmp(_dafny.Zero) == 0 {
			_nw9 = _dafny.NewArray(_len0_3)
		} else {
			var _init3 func(_dafny.Int) _dafny.Sequence = (func(_49_v0 bool) func(_dafny.Int) _dafny.Sequence {
				return func(_50_i3 _dafny.Int) _dafny.Sequence {
					return _dafny.SeqOf(_49_v0, _49_v0)
				}
			})(_17_v0)
			_ = _init3
			var _element0_3 = _init3(_dafny.Zero)
			_ = _element0_3
			_nw9 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
			(_nw9).ArraySet1(_element0_3, 0)
			var _nativeLen0_3 = (_len0_3).Int()
			_ = _nativeLen0_3
			for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
				(_nw9).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
			}
		}
		_48_v26 = _nw9
		var _51_v27 D1
		_ = _51_v27
		_51_v27 = Companion_D1_.Create_DC3_(_48_v26)
		var _52_v28 _dafny.Sequence
		_ = _52_v28
		_52_v28 = _dafny.SeqOf(_48_v26)
		var _53_v29 _dafny.Int
		_ = _53_v29
		_53_v29 = _dafny.IntOfInt64(80)
		var _pat_let_tv0 = _48_v26
		_ = _pat_let_tv0
		var _54_v30 _dafny.Array
		_ = _54_v30
		var _nwElement0_0 D1 = _51_v27
		_ = _nwElement0_0
		var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(24))
		_ = _nw10
		(_nw10).ArraySet1(_nwElement0_0, 0)
		(_nw10).ArraySet1(_51_v27, 1)
		(_nw10).ArraySet1(_51_v27, 2)
		(_nw10).ArraySet1(_51_v27, 3)
		(_nw10).ArraySet1(_51_v27, 4)
		(_nw10).ArraySet1(Companion_D1_.Create_DC3_((_52_v28).Select((Companion_Default___.SafeIndex(_53_v29, _dafny.IntOfUint32((_52_v28).Cardinality()))).Uint32()).(_dafny.Array)), 5)
		(_nw10).ArraySet1(_51_v27, 6)
		(_nw10).ArraySet1(_51_v27, 7)
		(_nw10).ArraySet1(_51_v27, 8)
		(_nw10).ArraySet1(_51_v27, 9)
		(_nw10).ArraySet1(_51_v27, 10)
		(_nw10).ArraySet1(_51_v27, 11)
		(_nw10).ArraySet1(_51_v27, 12)
		(_nw10).ArraySet1(_51_v27, 13)
		(_nw10).ArraySet1(Companion_D1_.Create_DC3_(_48_v26), 14)
		(_nw10).ArraySet1(_51_v27, 15)
		(_nw10).ArraySet1(_51_v27, 16)
		(_nw10).ArraySet1(func(_pat_let0_0 D1) D1 {
			return func(_55_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let1_0 _dafny.Array) D1 {
					return func(_56_dt__update_hcf7_h0 _dafny.Array) D1 {
						return Companion_D1_.Create_DC3_(_56_dt__update_hcf7_h0)
					}(_pat_let1_0)
				}(_pat_let_tv0)
			}(_pat_let0_0)
		}(_51_v27), 17)
		(_nw10).ArraySet1(_51_v27, 18)
		(_nw10).ArraySet1(_51_v27, 19)
		(_nw10).ArraySet1(Companion_D1_.Create_DC3_(_48_v26), 20)
		(_nw10).ArraySet1(_51_v27, 21)
		(_nw10).ArraySet1(Companion_D1_.Create_DC3_(_48_v26), 22)
		(_nw10).ArraySet1(_51_v27, 23)
		_54_v30 = _nw10
		var _57_v31 *C2
		_ = _57_v31
		var _nw11 *C2 = New_C2_()
		_ = _nw11
		_nw11.Ctor__(_45_v25, _54_v30)
		_57_v31 = _nw11
		(globalState).F1 = _53_v29
		var _58_v32 _dafny.MultiSet
		_ = _58_v32
		_58_v32 = _dafny.MultiSetOf(_53_v29, _dafny.IntOfInt64(536), _53_v29, Companion_Default___.Fm8(_53_v29, _17_v0, _17_v0, globalState))
		var _59_v33 _dafny.Array
		_ = _59_v33
		var _len0_4 _dafny.Int = _dafny.IntOfInt64(19)
		_ = _len0_4
		var _nw12 _dafny.Array
		_ = _nw12
		if _len0_4.Cmp(_dafny.Zero) == 0 {
			_nw12 = _dafny.NewArray(_len0_4)
		} else {
			var _init4 func(_dafny.Int) _dafny.Int = (func(_60_v29 _dafny.Int, _61_v0 bool) func(_dafny.Int) _dafny.Int {
				return func(_62_i4 _dafny.Int) _dafny.Int {
					return (_62_i4).Times((Companion_D14_.Create_DC40_(_60_v29, _61_v0, _dafny.UnicodeSeqOfUtf8Bytes("ywpyyikps"))).Dtor_cf73())
				}
			})(_53_v29, _17_v0)
			_ = _init4
			var _element0_4 = _init4(_dafny.Zero)
			_ = _element0_4
			_nw12 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
			(_nw12).ArraySet1(_element0_4, 0)
			var _nativeLen0_4 = (_len0_4).Int()
			_ = _nativeLen0_4
			for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
				(_nw12).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
			}
		}
		_59_v33 = _nw12
		var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_59_v33), 0))
		_ = _index8
		(_59_v33).ArraySet1(_53_v29, (_index8).Int())
		var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_59_v33), 0))
		_ = _index9
		var _rhs9 _dafny.Int = (_53_v29).Plus(_53_v29)
		_ = _rhs9
		var _rhs10 _dafny.MultiSet = ((_58_v32).Union(_58_v32)).Intersection((_58_v32).Union(Companion_Default___.Fm33(globalState)))
		_ = _rhs10
		var _rhs11 _dafny.Int = _53_v29
		_ = _rhs11
		var _lhs6 _dafny.Array = _59_v33
		_ = _lhs6
		var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(824), _dafny.ArrayLen((_59_v33), 0))
		_ = _lhs7
		_53_v29 = _rhs9
		_58_v32 = _rhs10
		(_lhs6).ArraySet1(_rhs11, (_lhs7).Int())
		var _63_v34 _dafny.Sequence
		_ = _63_v34
		_63_v34 = _dafny.UnicodeSeqOfUtf8Bytes("vejr")
		_63_v34 = _63_v34
	}
	var _64_v35 _dafny.Array
	_ = _64_v35
	var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
	_ = _nw13
	_64_v35 = _nw13
	var _65_v36 _dafny.Int
	_ = _65_v36
	_65_v36 = _dafny.IntOfInt64(-214)
	var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_64_v35), 0))
	_ = _index10
	(_64_v35).ArraySet1(_65_v36, (_index10).Int())
	var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_64_v35), 0))
	_ = _index11
	(_64_v35).ArraySet1(_65_v36, (_index11).Int())
	var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_64_v35), 0))
	_ = _index12
	(_64_v35).ArraySet1((_65_v36).Plus(_65_v36), (_index12).Int())
	_65_v36 = (_64_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_64_v35), 0))).Int()).(_dafny.Int)
	var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(91), _dafny.ArrayLen((_64_v35), 0))
	_ = _index13
	(_64_v35).ArraySet1(_65_v36, (_index13).Int())
	var _66_v37 _dafny.Sequence
	_ = _66_v37
	_66_v37 = _dafny.SeqOf(_65_v36)
	var _67_v38 _dafny.CodePoint
	_ = _67_v38
	_67_v38 = _dafny.CodePoint('x')
	var _68_v39 T1
	_ = _68_v39
	var _nw14 *C0 = New_C0_()
	_ = _nw14
	_nw14.Ctor__(_17_v0, _17_v0, _67_v38, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(407))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg13 _dafny.Int) interface{} {
			return coer13(arg13)
		}
	}((func(_69_v38 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
		return func(_70_i5 _dafny.Int) _dafny.CodePoint {
			return _69_v38
		}
	})(_67_v38))))
	_68_v39 = _nw14
	var _71_v40 _dafny.MultiSet
	_ = _71_v40
	_71_v40 = _dafny.MultiSetOf(_68_v39, _68_v39)
	var _72_v41 _dafny.Sequence
	_ = _72_v41
	_72_v41 = _dafny.SeqOf(_17_v0)
	var _73_v42 _dafny.Set
	_ = _73_v42
	_73_v42 = _dafny.SetOf((_71_v40).Cardinality(), _dafny.IntOfInt64(733), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_72_v41, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.IntOfUint32((_72_v41).Cardinality()))).Uint32(), _17_v0)).Cardinality()))
	var _74_v43 _dafny.Map
	_ = _74_v43
	_74_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_68_v39).F7())
	var _75_v44 T2
	_ = _75_v44
	var _nw15 *C1 = New_C1_()
	_ = _nw15
	_nw15.Ctor__((_73_v42).Cardinality(), (_68_v39).F6(), (_68_v39).F6(), (func() bool {
		if (_74_v43).Contains((_68_v39).F7()) {
			return (_74_v43).Get((_68_v39).F7()).(bool)
		}
		return _17_v0
	})(), _67_v38, (_68_v39).F5())
	_75_v44 = _nw15
	var _76_v45 _dafny.Map
	_ = _76_v45
	_76_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D7_.Create_DC21_(_66_v37, _75_v44), _65_v36)
	var _77_v46 D7
	_ = _77_v46
	_77_v46 = Companion_D7_.Create_DC21_(_66_v37, _75_v44)
	var _pat_let_tv1 = _66_v37
	_ = _pat_let_tv1
	var _pat_let_tv2 = _75_v44
	_ = _pat_let_tv2
	var _pat_let_tv3 = _66_v37
	_ = _pat_let_tv3
	var _pat_let_tv4 = _65_v36
	_ = _pat_let_tv4
	_76_v45 = (_76_v45).Update(func(_pat_let2_0 D7) D7 {
		return func(_78_dt__update__tmp_h1 D7) D7 {
			return func(_pat_let3_0 _dafny.Sequence) D7 {
				return func(_79_dt__update_hcf39_h0 _dafny.Sequence) D7 {
					return Companion_D7_.Create_DC21_(_79_dt__update_hcf39_h0, (_78_dt__update__tmp_h1).Dtor_cf40())
				}(_pat_let3_0)
			}(_dafny.Companion_Sequence_.Update(_pat_let_tv1, (Companion_Default___.SafeIndex((_pat_let_tv2).F8(), _dafny.IntOfUint32((_pat_let_tv3).Cardinality()))).Uint32(), _pat_let_tv4))
		}(_pat_let2_0)
	}(_77_v46), _dafny.IntOfInt64(957))
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _80_v0 bool
	_ = _80_v0
	_80_v0 = true
	var _81_v1 _dafny.Array
	_ = _81_v1
	var _nw16 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(12))
	_ = _nw16
	_81_v1 = _nw16
	var _82_v2 _dafny.Map
	_ = _82_v2
	_82_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v0, _81_v1)
	var _83_globalState *GlobalState
	_ = _83_globalState
	var _nw17 *GlobalState = New_GlobalState_()
	_ = _nw17
	_nw17.Ctor__(false, _dafny.IntOfInt64(130), false, (func() _dafny.Array {
		if (_82_v2).Contains(_80_v0) {
			return (_82_v2).Get(_80_v0).(_dafny.Array)
		}
		return _81_v1
	})())
	_83_globalState = _nw17
	var _84_v3 _dafny.Int
	_ = _84_v3
	_84_v3 = _dafny.IntOfInt64(287)
	_80_v0 = ((_dafny.Zero).Minus(_84_v3)).Cmp(_dafny.IntOfInt64(398)) < 0
	var _85_v4 _dafny.Sequence
	_ = _85_v4
	_85_v4 = _dafny.UnicodeSeqOfUtf8Bytes("ikmk")
	var _86_v5 _dafny.MultiSet
	_ = _86_v5
	_86_v5 = _dafny.MultiSetOf(Companion_Default___.Fm0(_80_v0, _dafny.IntOfUint32((_85_v4).Cardinality()), _84_v3, _80_v0, _83_globalState))
	_86_v5 = _86_v5
	_80_v0 = _80_v0
	Companion_Default___.M0(_83_globalState)
	var _87_v6 _dafny.Set
	_ = _87_v6
	_87_v6 = _dafny.SetOf(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("bmln"), _dafny.UnicodeSeqOfUtf8Bytes("bkkejxop")), _dafny.UnicodeSeqOfUtf8Bytes("vnqpsmb"), _85_v4)
	var _88_v7 _dafny.Sequence
	_ = _88_v7
	_88_v7 = _dafny.SeqOf(_85_v4)
	_87_v6 = func() _dafny.Set {
		var _coll7 = _dafny.NewBuilder()
		_ = _coll7
		for _iter7 := _dafny.Iterate(((func() _dafny.Sequence {
			if _80_v0 {
				return _88_v7
			}
			return _88_v7
		})()).Elements()); ; {
			_compr_7, _ok7 := _iter7()
			if !_ok7 {
				break
			}
			var _89_v8 _dafny.Sequence
			_89_v8 = interface{}(_compr_7).(_dafny.Sequence)
			if _dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
				if _80_v0 {
					return _88_v7
				}
				return _88_v7
			})(), _89_v8) {
				_coll7.Add(_89_v8)
			}
		}
		return _coll7.ToSet()
	}()
	var _90_i0 _dafny.Int
	_ = _90_i0
	_90_i0 = _dafny.Zero
	{
		for _80_v0 {
			{
				if (_90_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_90_i0 = (_90_i0).Plus(_dafny.One)
				var _91_v9 _dafny.Map
				_ = _91_v9
				_91_v9 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(349), _85_v4)
				var _92_v10 _dafny.CodePoint
				_ = _92_v10
				_92_v10 = _dafny.CodePoint('h')
				var _93_v11 _dafny.Map
				_ = _93_v11
				_93_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v0, Companion_Default___.Fm7(_92_v10, _84_v3, true, _83_globalState))
				var _94_v12 *C3
				_ = _94_v12
				var _nw18 *C3 = New_C3_()
				_ = _nw18
				_nw18.Ctor__(false, _80_v0, (_91_v9).Cardinality(), _80_v0, true, ((_93_v11).Cardinality()).Cmp(_84_v3) <= 0, _92_v10, _dafny.Companion_Sequence_.Concatenate(_85_v4, _85_v4))
				_94_v12 = _nw18
				var _95_v13 _dafny.Sequence
				_ = _95_v13
				_95_v13 = _dafny.SeqOf(_94_v12.F12, _94_v12.F12)
				var _96_v14 _dafny.Array
				_ = _96_v14
				var _nwElement0_1 bool = _80_v0
				_ = _nwElement0_1
				var _nw19 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(29))
				_ = _nw19
				(_nw19).ArraySet1(_nwElement0_1, 0)
				(_nw19).ArraySet1((_94_v12).F13(), 1)
				(_nw19).ArraySet1((_94_v12).F13(), 2)
				(_nw19).ArraySet1(_94_v12.F12, 3)
				(_nw19).ArraySet1((_94_v12).F13(), 4)
				(_nw19).ArraySet1(_80_v0, 5)
				(_nw19).ArraySet1((_94_v12).F13(), 6)
				(_nw19).ArraySet1(_80_v0, 7)
				(_nw19).ArraySet1((_94_v12).F13(), 8)
				(_nw19).ArraySet1((_94_v12).F13(), 9)
				(_nw19).ArraySet1(Companion_Default___.Fm3(_83_globalState), 10)
				(_nw19).ArraySet1((_94_v12).F13(), 11)
				(_nw19).ArraySet1((_95_v13).Select((Companion_Default___.SafeIndex(_84_v3, _dafny.IntOfUint32((_95_v13).Cardinality()))).Uint32()).(bool), 12)
				(_nw19).ArraySet1(!(_80_v0), 13)
				(_nw19).ArraySet1(_80_v0, 14)
				(_nw19).ArraySet1(_80_v0, 15)
				(_nw19).ArraySet1((_94_v12).F13(), 16)
				(_nw19).ArraySet1(!((_94_v12).F13()), 17)
				(_nw19).ArraySet1(_94_v12.F12, 18)
				(_nw19).ArraySet1(_80_v0, 19)
				(_nw19).ArraySet1(_94_v12.F12, 20)
				(_nw19).ArraySet1(_80_v0, 21)
				(_nw19).ArraySet1((_94_v12).F13(), 22)
				(_nw19).ArraySet1(_94_v12.F12, 23)
				(_nw19).ArraySet1(_94_v12.F12, 24)
				(_nw19).ArraySet1(!(_94_v12.F12), 25)
				(_nw19).ArraySet1(_80_v0, 26)
				(_nw19).ArraySet1(_94_v12.F12, 27)
				(_nw19).ArraySet1(!(_94_v12.F12), 28)
				_96_v14 = _nw19
				var _97_v15 _dafny.Map
				_ = _97_v15
				_97_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_96_v14, (func() bool {
					if (_93_v11).Contains((_94_v12).F13()) {
						return (_93_v11).Get((_94_v12).F13()).(bool)
					}
					return (_94_v12).F13()
				})())
				_97_v15 = (_97_v15).Update(_96_v14, (_94_v12).F13())
				var _98_v16 _dafny.Array
				_ = _98_v16
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(19)
				_ = _len0_5
				var _nw20 _dafny.Array
				_ = _nw20
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw20 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) _dafny.MultiSet = (func(_99_v3 _dafny.Int, _100_v13 _dafny.Sequence) func(_dafny.Int) _dafny.MultiSet {
						return func(_101_i1 _dafny.Int) _dafny.MultiSet {
							return (_dafny.MultiSetOf(_99_v3, _dafny.IntOfUint32((_100_v13).Cardinality()))).Difference(_dafny.MultiSetOf(_99_v3, _99_v3, _99_v3))
						}
					})(_84_v3, _95_v13)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw20 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw20).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw20).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_98_v16 = _nw20
				_98_v16 = _98_v16
				var _102_v17 _dafny.Map
				_ = _102_v17
				_102_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_94_v12).F13(), _95_v13)
				var _103_v18 _dafny.Array
				_ = _103_v18
				var _nwElement0_2 _dafny.Sequence = _95_v13
				_ = _nwElement0_2
				var _nw21 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(29))
				_ = _nw21
				(_nw21).ArraySet1(_nwElement0_2, 0)
				(_nw21).ArraySet1(_95_v13, 1)
				(_nw21).ArraySet1(_95_v13, 2)
				(_nw21).ArraySet1(_95_v13, 3)
				(_nw21).ArraySet1((func() _dafny.Sequence {
					if (_102_v17).Contains(false) {
						return (_102_v17).Get(false).(_dafny.Sequence)
					}
					return _95_v13
				})(), 4)
				(_nw21).ArraySet1(_95_v13, 5)
				(_nw21).ArraySet1(_95_v13, 6)
				(_nw21).ArraySet1(_95_v13, 7)
				(_nw21).ArraySet1(_95_v13, 8)
				(_nw21).ArraySet1(_95_v13, 9)
				(_nw21).ArraySet1(_dafny.SeqOf(true, _94_v12.F12), 10)
				(_nw21).ArraySet1(_95_v13, 11)
				(_nw21).ArraySet1(_95_v13, 12)
				(_nw21).ArraySet1(_95_v13, 13)
				(_nw21).ArraySet1(_95_v13, 14)
				(_nw21).ArraySet1(_95_v13, 15)
				(_nw21).ArraySet1(_95_v13, 16)
				(_nw21).ArraySet1(_95_v13, 17)
				(_nw21).ArraySet1(_95_v13, 18)
				(_nw21).ArraySet1(_95_v13, 19)
				(_nw21).ArraySet1(Companion_Default___.Fm18(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(640))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg14 _dafny.Int) interface{} {
						return coer14(arg14)
					}
				}((func(_104_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_105_i2 _dafny.Int) _dafny.CodePoint {
						return _104_v10
					}
				})(_92_v10))), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vqt")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(640))).Uint32(), func(coer15 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}((func(_106_v10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_107_i2 _dafny.Int) _dafny.CodePoint {
						return _106_v10
					}
				})(_92_v10)))).Cardinality()))).Uint32(), _92_v10)).Cardinality()), _94_v12.F12, _83_globalState), 20)
				(_nw21).ArraySet1(_95_v13, 21)
				(_nw21).ArraySet1(_95_v13, 22)
				(_nw21).ArraySet1(_95_v13, 23)
				(_nw21).ArraySet1(_95_v13, 24)
				(_nw21).ArraySet1(_95_v13, 25)
				(_nw21).ArraySet1(_95_v13, 26)
				(_nw21).ArraySet1(_95_v13, 27)
				(_nw21).ArraySet1(_95_v13, 28)
				_103_v18 = _nw21
				var _108_v19 D1
				_ = _108_v19
				_108_v19 = Companion_D1_.Create_DC3_(_103_v18)
				var _109_v20 _dafny.Map
				_ = _109_v20
				_109_v20 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v3, _108_v19)
				var _110_v21 D2
				_ = _110_v21
				_110_v21 = Companion_D2_.Create_DC6_(_109_v20)
				var _111_v22 _dafny.Set
				_ = _111_v22
				_111_v22 = _dafny.SetOf(_110_v21)
				var _112_v23 D13
				_ = _112_v23
				_112_v23 = Companion_D13_.Create_DC35_(_111_v22)
				_111_v22 = (_112_v23).Dtor_cf66()
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	_80_v0 = (func() bool {
		if true {
			return Companion_Default___.Fm3(_83_globalState)
		}
		return false
	})()
	_80_v0 = _80_v0
	var _113_v24 _dafny.MultiSet
	_ = _113_v24
	_113_v24 = _dafny.MultiSetOf(_84_v3)
	_113_v24 = _113_v24
	_84_v3 = _84_v3
	var _114_v25 _dafny.Sequence
	_ = _114_v25
	_114_v25 = _dafny.SeqOf(_dafny.IntOfInt64(372), _84_v3)
	var _115_v26 _dafny.Sequence
	_ = _115_v26
	_115_v26 = _dafny.SeqOf(_dafny.MultiSetFromSeq(_114_v25))
	var _116_v27 _dafny.CodePoint
	_ = _116_v27
	_116_v27 = _dafny.CodePoint('b')
	var _117_v28 *C1
	_ = _117_v28
	var _nw22 *C1 = New_C1_()
	_ = _nw22
	_nw22.Ctor__(_84_v3, _80_v0, _80_v0, _80_v0, _116_v27, _85_v4)
	_117_v28 = _nw22
	var _118_v29 _dafny.Array
	_ = _118_v29
	var _nw23 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(2))
	_ = _nw23
	_118_v29 = _nw23
	var _119_v30 *C5
	_ = _119_v30
	var _nw24 *C5 = New_C5_()
	_ = _nw24
	_nw24.Ctor__(_118_v29, _81_v1, !(_80_v0), true, (_85_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-922))).Uint32(), func(coer16 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg16 _dafny.Int) interface{} {
			return coer16(arg16)
		}
	}(func(_120_i3 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('u')
	}))).Cardinality()), _dafny.IntOfUint32((_85_v4).Cardinality()))).Uint32()).(_dafny.CodePoint), _85_v4, _84_v3, _80_v0)
	_119_v30 = _nw24
	var _121_v31 _dafny.Map
	_ = _121_v31
	_121_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v0, _119_v30)
	var _122_v32 T3
	_ = _122_v32
	var _nw25 *C4 = New_C4_()
	_ = _nw25
	_nw25.Ctor__(_84_v3, _80_v0, Companion_Default___.Fm27(_84_v3, _80_v0, Companion_Default___.Fm27(_84_v3, !(!(Companion_Default___.Fm27(_84_v3, _80_v0, _80_v0, _83_globalState))), _80_v0, _83_globalState), _83_globalState), _80_v0, _dafny.CodePoint('m'), _85_v4)
	_122_v32 = _nw25
	var _123_v33 _dafny.Map
	_ = _123_v33
	_123_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_122_v32, (_122_v32).F7())
	var _124_v34 D0
	_ = _124_v34
	_124_v34 = Companion_D0_.Create_DC2_(_84_v3, _80_v0, _114_v25)
	var _125_v35 _dafny.Array
	_ = _125_v35
	var _nwElement0_3 _dafny.MultiSet = _dafny.MultiSetFromSeq(_114_v25)
	_ = _nwElement0_3
	var _nw26 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(28))
	_ = _nw26
	(_nw26).ArraySet1(_nwElement0_3, 0)
	(_nw26).ArraySet1((_113_v24).Intersection(_dafny.MultiSetOf(_84_v3)), 1)
	(_nw26).ArraySet1(_113_v24, 2)
	(_nw26).ArraySet1(((_115_v26).Select((Companion_Default___.SafeIndex(_84_v3, _dafny.IntOfUint32((_115_v26).Cardinality()))).Uint32()).(_dafny.MultiSet)).Intersection(_113_v24), 3)
	(_nw26).ArraySet1((_113_v24).Update(_84_v3, Companion_Default___.Abs(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_117_v28, ((_113_v24).Update(_84_v3, Companion_Default___.Abs(_84_v3))).Cardinality())).Update(_117_v28, _84_v3)).Cardinality())), 4)
	(_nw26).ArraySet1(_113_v24, 5)
	(_nw26).ArraySet1(_113_v24, 6)
	(_nw26).ArraySet1((_113_v24).Difference(_113_v24), 7)
	(_nw26).ArraySet1(_113_v24, 8)
	(_nw26).ArraySet1(_113_v24, 9)
	(_nw26).ArraySet1(_113_v24, 10)
	(_nw26).ArraySet1(_113_v24, 11)
	(_nw26).ArraySet1(_113_v24, 12)
	(_nw26).ArraySet1(_113_v24, 13)
	(_nw26).ArraySet1(_113_v24, 14)
	(_nw26).ArraySet1(_113_v24, 15)
	(_nw26).ArraySet1(_113_v24, 16)
	(_nw26).ArraySet1(Companion_Default___.Fm33(_83_globalState), 17)
	(_nw26).ArraySet1(_113_v24, 18)
	(_nw26).ArraySet1((_dafny.MultiSetOf((_121_v31).Cardinality(), _84_v3)).Difference(_113_v24), 19)
	(_nw26).ArraySet1(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(608))).Uint32(), func(coer17 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg17 _dafny.Int) interface{} {
			return coer17(arg17)
		}
	}((func(_126_v3 _dafny.Int) func(_dafny.Int) _dafny.Int {
		return func(_127_i4 _dafny.Int) _dafny.Int {
			return (_dafny.Zero).Minus(_126_v3)
		}
	})(_84_v3))), _dafny.SeqOf(_84_v3, _dafny.IntOfUint32((_114_v25).Cardinality()), _84_v3))), 20)
	(_nw26).ArraySet1(_113_v24, 21)
	(_nw26).ArraySet1(_113_v24, 22)
	(_nw26).ArraySet1((_dafny.MultiSetOf(_dafny.IntOfUint32((_85_v4).Cardinality()))).Union(_113_v24), 23)
	(_nw26).ArraySet1(_113_v24, 24)
	(_nw26).ArraySet1(_113_v24, 25)
	(_nw26).ArraySet1((_dafny.MultiSetOf(_84_v3)).Intersection(_113_v24), 26)
	(_nw26).ArraySet1(_dafny.MultiSetOf((_123_v33).Cardinality(), _dafny.IntOfUint32((_dafny.SeqOf((_122_v32).F6(), (_122_v32).F7(), (_122_v32).F7(), (_124_v34).Dtor_cf5(), _80_v0)).Cardinality()), _dafny.IntOfInt64(-415), _84_v3, _84_v3), 27)
	_125_v35 = _nw26
	var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_125_v35), 0))
	_ = _index14
	(_125_v35).ArraySet1(_113_v24, (_index14).Int())
	var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_125_v35), 0))
	_ = _index15
	(_125_v35).ArraySet1((_113_v24).Intersection(_113_v24), (_index15).Int())
	var _source1 D1 = Companion_Default___.Fm26((_122_v32).F7(), _dafny.IntOfInt64(147), _122_v32.F4(), _83_globalState)
	_ = _source1
	if _source1.Is_DC4() {
		var _128___mcc_h0 _dafny.CodePoint = _source1.Get_().(D1_DC4).Cf8
		_ = _128___mcc_h0
		var _129___mcc_h1 _dafny.Int = _source1.Get_().(D1_DC4).Cf9
		_ = _129___mcc_h1
		var _130___mcc_h2 _dafny.CodePoint = _source1.Get_().(D1_DC4).Cf10
		_ = _130___mcc_h2
		var _131___mcc_h3 _dafny.Int = _source1.Get_().(D1_DC4).Cf11
		_ = _131___mcc_h3
		var _132___mcc_h4 bool = _source1.Get_().(D1_DC4).Cf12
		_ = _132___mcc_h4
		var _133_cf12 bool = _132___mcc_h4
		_ = _133_cf12
		var _134_cf11 _dafny.Int = _131___mcc_h3
		_ = _134_cf11
		var _135_cf10 _dafny.CodePoint = _130___mcc_h2
		_ = _135_cf10
		var _136_cf9 _dafny.Int = _129___mcc_h1
		_ = _136_cf9
		var _137_cf8 _dafny.CodePoint = _128___mcc_h0
		_ = _137_cf8
		var _138_v36 _dafny.MultiSet
		_ = _138_v36
		_138_v36 = _dafny.MultiSetOf(_80_v0, (_122_v32).F6())
		var _139_v37 _dafny.Int
		_ = _139_v37
		var _140_v38 _dafny.Int
		_ = _140_v38
		var _141_v39 _dafny.Array
		_ = _141_v39
		var _out0 _dafny.Int
		_ = _out0
		var _out1 _dafny.Int
		_ = _out1
		var _out2 _dafny.Array
		_ = _out2
		_out0, _out1, _out2 = (_119_v30).M1((_122_v32).F6(), (_122_v32).F6(), (_138_v36).Update((_122_v32).F6(), Companion_Default___.Abs(_dafny.IntOfInt64(71))), (_122_v32).F6(), _83_globalState)
		_139_v37 = _out0
		_140_v38 = _out1
		_141_v39 = _out2
		var _142_v40 _dafny.Sequence
		_ = _142_v40
		_142_v40 = _dafny.SeqOf((_122_v32).F9())
		var _143_v41 _dafny.Map
		_ = _143_v41
		_143_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v0, _dafny.IntOfUint32((_142_v40).Cardinality()))
		var _144_v42 D10
		_ = _144_v42
		_144_v42 = Companion_D10_.Create_DC28_(_133_cf12, (_143_v41).Cardinality(), (_122_v32).F9(), (_122_v32).F9(), _117_v28)
		var _145_v43 T2
		_ = _145_v43
		var _nw27 *C1 = New_C1_()
		_ = _nw27
		_nw27.Ctor__((_122_v32).F8(), (_122_v32).F6(), (_122_v32).F6(), (_144_v42).Dtor_cf50(), _dafny.CodePoint('t'), (_122_v32).F5())
		_145_v43 = _nw27
		var _146_v44 _dafny.Map
		_ = _146_v44
		_146_v44 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_145_v43, _133_cf12)
		var _147_v45 D5
		_ = _147_v45
		_147_v45 = Companion_D5_.Create_DC15_((_119_v30).F10(), (_122_v32).F9(), _143_v41, (_146_v44).Cardinality())
		var _148_v46 T2
		_ = _148_v46
		var _nw28 *C3 = New_C3_()
		_ = _nw28
		_nw28.Ctor__(!((_122_v32).F7()), ((_122_v32).F7()) && (_80_v0), _136_cf9, (_147_v45).Dtor_cf27(), false, false, _137_cf8, _dafny.Companion_Sequence_.Concatenate((_122_v32).F5(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(519))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg18 _dafny.Int) interface{} {
				return coer18(arg18)
			}
		}((func(_149_v32 T3) func(_dafny.Int) _dafny.CodePoint {
			return func(_150_i5 _dafny.Int) _dafny.CodePoint {
				return _149_v32.F4()
			}
		})(_122_v32)))))
		_148_v46 = _nw28
		_148_v46 = _145_v43
		_133_cf12 = _dafny.Companion_Sequence_.IsPrefixOf((_145_v43).F5(), (_145_v43).F5())
		(_148_v46).F4_set_(_135_cf10)
	} else if _source1.Is_DC5() {
		var _151___mcc_h5 _dafny.Int = _source1.Get_().(D1_DC5).Cf13
		_ = _151___mcc_h5
		var _152_cf13 _dafny.Int = _151___mcc_h5
		_ = _152_cf13
		(_83_globalState).F1 = _84_v3
		var _153_v47 _dafny.Map
		_ = _153_v47
		_153_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_84_v3).Minus(_152_cf13), (_122_v32).F6())
		_153_v47 = _153_v47
		var _154_v48 _dafny.Sequence
		_ = _154_v48
		_154_v48 = _dafny.SeqOf(false)
		var _155_v49 _dafny.Map
		_ = _155_v49
		_155_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_152_cf13, (_122_v32).F8()), _154_v48)
		_155_v49 = _155_v49
		var _156_v50 D2
		_ = _156_v50
		_156_v50 = Companion_D2_.Create_DC7_((_122_v32).F9())
		_80_v0 = (_117_v28).Fm19((_122_v32).F6(), _80_v0, _dafny.SetOf(_156_v50, _156_v50, _156_v50, _156_v50, _156_v50), _83_globalState)
	} else {
		var _157___mcc_h6 _dafny.Array = _source1.Get_().(D1_DC3).Cf7
		_ = _157___mcc_h6
		var _158_cf7 _dafny.Array = _157___mcc_h6
		_ = _158_cf7
		var _159_v51 _dafny.Array
		_ = _159_v51
		var _nw29 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(28))
		_ = _nw29
		_159_v51 = _nw29
		var _160_v52 D14
		_ = _160_v52
		_160_v52 = Companion_D14_.Create_DC39_(_159_v51)
		var _161_v53 _dafny.Map
		_ = _161_v53
		_161_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_160_v52).Dtor_cf72(), _dafny.Companion_Sequence_.IsPrefixOf((_122_v32).F5(), _85_v4))
		if !(!((func() bool {
			if (_161_v53).Contains(_159_v51) {
				return (_161_v53).Get(_159_v51).(bool)
			}
			return (_122_v32).F9()
		})())) {
			var _162_v54 T2
			_ = _162_v54
			var _nw30 *C3 = New_C3_()
			_ = _nw30
			_nw30.Ctor__((_122_v32).F9(), (_122_v32).F7(), (_122_v32).F8(), (_122_v32).F6(), (_122_v32).F6(), _80_v0, _122_v32.F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(424))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_163_v32 T3) func(_dafny.Int) _dafny.CodePoint {
				return func(_164_i6 _dafny.Int) _dafny.CodePoint {
					return _163_v32.F4()
				}
			})(_122_v32))))
			_162_v54 = _nw30
			var _165_v55 D7
			_ = _165_v55
			_165_v55 = Companion_D7_.Create_DC21_(_114_v25, _162_v54)
			_80_v0 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_165_v55).Dtor_cf39()).Cardinality()), _84_v3)).Cmp((_122_v32).F8()) <= 0
			var _166_v56 _dafny.Set
			_ = _166_v56
			_166_v56 = _dafny.SetOf(_dafny.IntOfInt64(353))
			var _167_v57 _dafny.Sequence
			_ = _167_v57
			_167_v57 = _dafny.SeqOf(_166_v56)
			var _168_v58 _dafny.Map
			_ = _168_v58
			_168_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Set {
				if (_122_v32).F7() {
					return _166_v56
				}
				return _166_v56
			})(), _dafny.IntOfUint32((_167_v57).Cardinality()))
			_168_v58 = _168_v58
			_80_v0 = (_122_v32).F7()
			var _169_v59 _dafny.Map
			_ = _169_v59
			_169_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F9(), (_122_v32).F9())
			var _170_v60 _dafny.Sequence
			_ = _170_v60
			_170_v60 = _dafny.SeqOf((_122_v32).F7())
			_169_v59 = (_169_v59).Update(!(!(_dafny.Companion_Sequence_.Equal(_dafny.SeqOf((_162_v54).F8(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_170_v60).Cardinality()))), _114_v25))), (_162_v54).F7())
			var _171_v61 D5
			_ = _171_v61
			_171_v61 = Companion_D5_.Create_DC14_((_119_v30).F11())
			var _172_v62 _dafny.Set
			_ = _172_v62
			_172_v62 = _dafny.SetOf(Companion_D5_.Create_DC14_(_81_v1), _171_v61, _171_v61, Companion_D5_.Create_DC14_(_81_v1), _171_v61)
			var _173_v63 _dafny.Map
			_ = _173_v63
			_173_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_116_v27, ((_122_v32).F8()).Times((_172_v62).Cardinality()))
			_173_v63 = (_173_v63).Update(_162_v54.F4(), _84_v3)
		} else {
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen(((_119_v30).F10()), 0))
			_ = _index16
			((_119_v30).F10()).ArraySet1((_122_v32).F7(), (_index16).Int())
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen(((_119_v30).F10()), 0))
			_ = _index17
			((_119_v30).F10()).ArraySet1(!(!((_122_v32).F7())), (_index17).Int())
			var _174_v64 _dafny.Map
			_ = _174_v64
			_174_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F8(), (_122_v32).F7())
			var _175_v65 T1
			_ = _175_v65
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__((func() bool {
				if (_174_v64).Contains((_122_v32).F8()) {
					return (_174_v64).Get((_122_v32).F8()).(bool)
				}
				return !((_122_v32).F6())
			})(), ((_119_v30).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(265), _dafny.ArrayLen(((_119_v30).F10()), 0))).Int()).(bool), _122_v32.F4(), (_122_v32).F5())
			_175_v65 = _nw31
			_175_v65 = _175_v65
			var _176_v66 _dafny.Map
			_ = _176_v66
			_176_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_119_v30, (_122_v32).F8())
			var _177_v67 _dafny.Sequence
			_ = _177_v67
			_177_v67 = _dafny.SeqOf(_119_v30)
			(_83_globalState).F1 = (_114_v25).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(((func() _dafny.Int {
				if (_176_v66).Contains((_177_v67).Select((Companion_Default___.SafeIndex(_84_v3, _dafny.IntOfUint32((_177_v67).Cardinality()))).Uint32()).(*C5)) {
					return (_176_v66).Get((_177_v67).Select((Companion_Default___.SafeIndex(_84_v3, _dafny.IntOfUint32((_177_v67).Cardinality()))).Uint32()).(*C5)).(_dafny.Int)
				}
				return _84_v3
			})()).Minus((_122_v32).F8())), _dafny.IntOfUint32((_114_v25).Cardinality()))).Uint32()).(_dafny.Int)
			var _178_v68 D12
			_ = _178_v68
			_178_v68 = Companion_D12_.Create_DC34_((_122_v32).F6())
			var _179_v69 *C4
			_ = _179_v69
			var _nw32 *C4 = New_C4_()
			_ = _nw32
			_nw32.Ctor__((_122_v32).F8(), (_122_v32).F7(), false, (_178_v68).Dtor_cf65(), _122_v32.F4(), _dafny.Companion_Sequence_.Concatenate((_122_v32).F5(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(32))).Uint32(), func(coer20 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg20 _dafny.Int) interface{} {
					return coer20(arg20)
				}
			}((func(_180_v32 T3) func(_dafny.Int) _dafny.CodePoint {
				return func(_181_i7 _dafny.Int) _dafny.CodePoint {
					return _180_v32.F4()
				}
			})(_122_v32)))))
			_179_v69 = _nw32
			var _182_v70 _dafny.Array
			_ = _182_v70
			var _len0_6 _dafny.Int = _dafny.IntOfInt64(20)
			_ = _len0_6
			var _nw33 _dafny.Array
			_ = _nw33
			if _len0_6.Cmp(_dafny.Zero) == 0 {
				_nw33 = _dafny.NewArray(_len0_6)
			} else {
				var _init6 func(_dafny.Int) _dafny.Sequence = (func(_183_v32 T3) func(_dafny.Int) _dafny.Sequence {
					return func(_184_i8 _dafny.Int) _dafny.Sequence {
						return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(971))).Uint32(), func(coer21 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg21 _dafny.Int) interface{} {
								return coer21(arg21)
							}
						}((func(_185_v32 T3) func(_dafny.Int) _dafny.CodePoint {
							return func(_186_i9 _dafny.Int) _dafny.CodePoint {
								return _185_v32.F4()
							}
						})(_183_v32)))
					}
				})(_122_v32)
				_ = _init6
				var _element0_6 = _init6(_dafny.Zero)
				_ = _element0_6
				_nw33 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
				(_nw33).ArraySet1(_element0_6, 0)
				var _nativeLen0_6 = (_len0_6).Int()
				_ = _nativeLen0_6
				for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
					(_nw33).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
				}
			}
			_182_v70 = _nw33
			var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_182_v70), 0))
			_ = _index18
			(_182_v70).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("dlgmycrna"), (_index18).Int())
			var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(134), _dafny.ArrayLen((_182_v70), 0))
			_ = _index19
			(_182_v70).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_122_v32).F5(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(211))).Uint32(), func(coer22 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg22 _dafny.Int) interface{} {
					return coer22(arg22)
				}
			}(func(_187_i10 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('b')
			}))), (_index19).Int())
		}
		var _188_v71 _dafny.Sequence
		_ = _188_v71
		_188_v71 = _dafny.SeqOf((_122_v32).F7())
		var _189_v72 _dafny.Set
		_ = _189_v72
		_189_v72 = _dafny.SetOf(_188_v71)
		var _190_v73 *C4
		_ = _190_v73
		var _nw34 *C4 = New_C4_()
		_ = _nw34
		_nw34.Ctor__((_189_v72).Cardinality(), _80_v0, (_122_v32).F6(), !(!(false)), _dafny.CodePoint('g'), _dafny.Companion_Sequence_.Concatenate((_122_v32).F5(), _dafny.UnicodeSeqOfUtf8Bytes("lctrtkorn")))
		_190_v73 = _nw34
		_190_v73 = _190_v73
		_84_v3 = (_122_v32).F8()
		var _191_v74 _dafny.Map
		_ = _191_v74
		_191_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(162), _84_v3)
		var _192_v75 _dafny.MultiSet
		_ = _192_v75
		_192_v75 = _dafny.MultiSetOf((_122_v32).F9(), true, Companion_Default___.Fm3(_83_globalState))
		var _193_v76 _dafny.Int
		_ = _193_v76
		var _194_v77 _dafny.Int
		_ = _194_v77
		var _195_v78 _dafny.Array
		_ = _195_v78
		var _out3 _dafny.Int
		_ = _out3
		var _out4 _dafny.Int
		_ = _out4
		var _out5 _dafny.Array
		_ = _out5
		_out3, _out4, _out5 = (_190_v73).M1(_80_v0, !(!(_191_v74).Contains(_dafny.IntOfUint32(((_122_v32).F5()).Cardinality()))), _192_v75, (_122_v32).F6(), _83_globalState)
		_193_v76 = _out3
		_194_v77 = _out4
		_195_v78 = _out5
	}
	var _196_v79 _dafny.Set
	_ = _196_v79
	_196_v79 = _dafny.SetOf((_122_v32).F9())
	var _197_v80 _dafny.Sequence
	_ = _197_v80
	_197_v80 = _dafny.SeqOf((_196_v79).IsSubsetOf(_196_v79), (_122_v32).F9(), (_122_v32).F9())
	_197_v80 = _dafny.Companion_Sequence_.Concatenate(_197_v80, _197_v80)
	var _198_v81 *C4
	_ = _198_v81
	var _nw35 *C4 = New_C4_()
	_ = _nw35
	_nw35.Ctor__(_84_v3, _80_v0, _80_v0, true, _116_v27, _85_v4)
	_198_v81 = _nw35
	var _199_v82 _dafny.Set
	_ = _199_v82
	_199_v82 = _dafny.SetOf(_198_v81, (func() *C4 {
		if Companion_Default___.Fm27((_dafny.Zero).Minus(_dafny.IntOfInt64(-56)), true, !((_122_v32).F7()), _83_globalState) {
			return _198_v81
		}
		return _198_v81
	})())
	var _200_v83 T0
	_ = _200_v83
	var _nw36 *C4 = New_C4_()
	_ = _nw36
	_nw36.Ctor__(_84_v3, (_122_v32).F7(), (_122_v32).F9(), _80_v0, _122_v32.F4(), _85_v4)
	_200_v83 = _nw36
	var _201_v84 _dafny.Map
	_ = _201_v84
	_201_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_80_v0, (_80_v0) || ((_122_v32).F9()))
	var _202_v85 _dafny.Array
	_ = _202_v85
	var _nw37 _dafny.Array = _dafny.NewArrayWithValue(Companion_D11_.Default(), _dafny.IntOfInt64(5))
	_ = _nw37
	_202_v85 = _nw37
	var _rhs12 _dafny.Set = _199_v82
	_ = _rhs12
	var _rhs13 bool = (_122_v32).F7()
	_ = _rhs13
	var _rhs14 T0 = _200_v83
	_ = _rhs14
	var _rhs15 _dafny.Map = _201_v84
	_ = _rhs15
	var _rhs16 _dafny.Array = _202_v85
	_ = _rhs16
	_199_v82 = _rhs12
	_80_v0 = _rhs13
	_200_v83 = _rhs14
	_201_v84 = _rhs15
	_202_v85 = _rhs16
	var _203_v86 _dafny.Map
	_ = _203_v86
	_203_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F8(), true)
	if (func() bool {
		if (_203_v86).Contains(Companion_Default___.SafeModuloInt((_122_v32).F8(), _dafny.IntOfUint32(((_122_v32).F5()).Cardinality()))) {
			return (_203_v86).Get(Companion_Default___.SafeModuloInt((_122_v32).F8(), _dafny.IntOfUint32(((_122_v32).F5()).Cardinality()))).(bool)
		}
		return (_197_v80).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_85_v4).Cardinality()), _dafny.IntOfUint32((_197_v80).Cardinality()))).Uint32()).(bool)
	})() {
		if _80_v0 {
			var _204_v87 *C0
			_ = _204_v87
			var _nw38 *C0 = New_C0_()
			_ = _nw38
			_nw38.Ctor__((func() bool {
				if (_122_v32).F9() {
					return (_122_v32).F7()
				}
				return (_122_v32).F9()
			})(), (_122_v32).F6(), ((_200_v83).F5()).Select((Companion_Default___.SafeIndex(_84_v3, _dafny.IntOfUint32(((_200_v83).F5()).Cardinality()))).Uint32()).(_dafny.CodePoint), (_122_v32).F5())
			_204_v87 = _nw38
			_204_v87 = _204_v87
			(_122_v32).F4_set_(_200_v83.F4())
			var _205_v88 _dafny.Array
			_ = _205_v88
			var _nw39 _dafny.Array = _dafny.NewArrayWithValue(Companion_D7_.Default(), _dafny.IntOfInt64(22))
			_ = _nw39
			_205_v88 = _nw39
			var _206_v89 _dafny.Map
			_ = _206_v89
			_206_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F6(), _205_v88)
			var _207_v90 _dafny.Map
			_ = _207_v90
			_207_v90 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_206_v89).Merge(_206_v89), ((_117_v28).Fm1((_122_v32).F7(), _201_v84, (_122_v32).F6(), _83_globalState)).Cmp((_122_v32).F8()) < 0)
			var _208_v91 _dafny.Sequence
			_ = _208_v91
			_208_v91 = _dafny.SeqOf((_119_v30).F10(), (_119_v30).F10())
			var _209_v92 _dafny.Sequence
			_ = _209_v92
			_209_v92 = _dafny.SeqOf(_208_v91, _208_v91)
			_207_v90 = (_207_v90).Update(_206_v89, ((_125_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_125_v35), 0))).Int()).(_dafny.MultiSet)).IsProperSubsetOf(_dafny.MultiSetOf((_122_v32).F8(), _dafny.IntOfUint32(((_209_v92).Select((Companion_Default___.SafeIndex((_122_v32).F8(), _dafny.IntOfUint32((_209_v92).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), _84_v3)))
			var _210_v93 _dafny.MultiSet
			_ = _210_v93
			_210_v93 = _dafny.MultiSetOf((_122_v32).F6())
			var _211_v94 _dafny.Int
			_ = _211_v94
			var _212_v95 _dafny.Int
			_ = _212_v95
			var _213_v96 _dafny.Array
			_ = _213_v96
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			var _out8 _dafny.Array
			_ = _out8
			_out6, _out7, _out8 = (_122_v32).M1((_122_v32).F6(), _80_v0, _210_v93, !((_122_v32).F9()), _83_globalState)
			_211_v94 = _out6
			_212_v95 = _out7
			_213_v96 = _out8
			var _214_v97 D0
			_ = _214_v97
			_214_v97 = Companion_D0_.Create_DC1_(false, _211_v94, (_122_v32).F7(), _84_v3)
			_80_v0 = !((_214_v97).Dtor_cf2())
		} else {
			var _215_v98 _dafny.Sequence
			_ = _215_v98
			_215_v98 = _dafny.SeqOf((Companion_Default___.Fm15(_83_globalState)).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F8(), _84_v3)))
			_215_v98 = _215_v98
			var _216_v99 D12
			_ = _216_v99
			_216_v99 = Companion_D12_.Create_DC34_(false)
			_80_v0 = (_216_v99).Dtor_cf65()
			_80_v0 = (_122_v32).F9()
			var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen(((_119_v30).F11()), 0))
			_ = _index20
			((_119_v30).F11()).ArraySet1((_dafny.IntOfInt64(85)).Times(_84_v3), (_index20).Int())
			var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen(((_119_v30).F11()), 0))
			_ = _index21
			((_119_v30).F11()).ArraySet1(_dafny.IntOfInt64(766), (_index21).Int())
			var _217_v100 *C3
			_ = _217_v100
			var _nw40 *C3 = New_C3_()
			_ = _nw40
			_nw40.Ctor__(false, (_122_v32).F7(), ((_119_v30).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(738), _dafny.ArrayLen(((_119_v30).F11()), 0))).Int()).(_dafny.Int), true, _80_v0, !((_122_v32).F7()) || (true), _122_v32.F4(), (_200_v83).F5())
			_217_v100 = _nw40
		}
		var _218_v101 D7
		_ = _218_v101
		_218_v101 = Companion_D7_.Create_DC20_()
		var _source2 D7 = _218_v101
		_ = _source2
		if _source2.Is_DC19() {
			var _219___mcc_h7 bool = _source2.Get_().(D7_DC19).Cf34
			_ = _219___mcc_h7
			var _220___mcc_h8 _dafny.Sequence = _source2.Get_().(D7_DC19).Cf35
			_ = _220___mcc_h8
			var _221___mcc_h9 bool = _source2.Get_().(D7_DC19).Cf36
			_ = _221___mcc_h9
			var _222___mcc_h10 bool = _source2.Get_().(D7_DC19).Cf37
			_ = _222___mcc_h10
			var _223___mcc_h11 _dafny.Int = _source2.Get_().(D7_DC19).Cf38
			_ = _223___mcc_h11
			var _224_cf38 _dafny.Int = _223___mcc_h11
			_ = _224_cf38
			var _225_cf37 bool = _222___mcc_h10
			_ = _225_cf37
			var _226_cf36 bool = _221___mcc_h9
			_ = _226_cf36
			var _227_cf35 _dafny.Sequence = _220___mcc_h8
			_ = _227_cf35
			var _228_cf34 bool = _219___mcc_h7
			_ = _228_cf34
			var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen(((_119_v30).F10()), 0))
			_ = _index22
			((_119_v30).F10()).ArraySet1(_80_v0, (_index22).Int())
			var _229_v102 _dafny.Map
			_ = _229_v102
			_229_v102 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_197_v80, (_122_v32).F9())
			var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(147), _dafny.ArrayLen(((_119_v30).F10()), 0))
			_ = _index23
			((_119_v30).F10()).ArraySet1(!(!((_122_v32).F6()) || (_228_cf34)) || (_dafny.Companion_Sequence_.Equal((_122_v32).F5(), _dafny.Companion_Sequence_.Update((_200_v83).F5(), (Companion_Default___.SafeIndex((_229_v102).Cardinality(), _dafny.IntOfUint32(((_200_v83).F5()).Cardinality()))).Uint32(), _200_v83.F4()))), (_index23).Int())
			_198_v81 = _198_v81
			var _230_v103 _dafny.Map
			_ = _230_v103
			_230_v103 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).Fm1((_122_v32).F6(), _201_v84, (_122_v32).F6(), _83_globalState), _84_v3)
			var _231_v104 _dafny.MultiSet
			_ = _231_v104
			_231_v104 = _dafny.MultiSetOf(_198_v81)
			_230_v103 = (_230_v103).Update((_dafny.Zero).Minus((_122_v32).F8()), (func() _dafny.Int {
				if (_231_v104).Contains(_198_v81) {
					return (_231_v104).Multiplicity(_198_v81)
				}
				return (_122_v32).F8()
			})())
			var _232_v105 _dafny.Map
			_ = _232_v105
			_232_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_200_v83, (_122_v32).F8())
			_232_v105 = (_232_v105).Update(_200_v83, _dafny.IntOfInt64(732))
		} else if _source2.Is_DC20() {
			var _233_v106 _dafny.Array
			_ = _233_v106
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(24))
			_ = _nw41
			_233_v106 = _nw41
			var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_233_v106), 0))
			_ = _index24
			(_233_v106).ArraySet1(_85_v4, (_index24).Int())
			var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(233), _dafny.ArrayLen((_233_v106), 0))
			_ = _index25
			(_233_v106).ArraySet1((_200_v83).F5(), (_index25).Int())
			var _234_v107 *C5
			_ = _234_v107
			var _nw42 *C5 = New_C5_()
			_ = _nw42
			_nw42.Ctor__(_118_v29, (_119_v30).F11(), (_122_v32).F7(), (_122_v32).F9(), _200_v83.F4(), _85_v4, (_84_v3).Minus((_dafny.Zero).Minus((_122_v32).F8())), (_122_v32).F6())
			_234_v107 = _nw42
			var _235_v108 _dafny.Map
			_ = _235_v108
			_235_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v3, _dafny.IntOfInt64(956))
			var _236_v109 _dafny.Map
			_ = _236_v109
			_236_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((((_235_v108).Update(_dafny.IntOfUint32((_114_v25).Cardinality()), (_122_v32).F8())).Cardinality()).Cmp(_dafny.IntOfInt64(904)) > 0, (_117_v28).Fm1(!((_122_v32).F7()), Companion_Default___.Fm23(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(991))).Uint32(), func(coer23 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg23 _dafny.Int) interface{} {
					return coer23(arg23)
				}
			}(func(_237_i11 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(827)
			}))).Cardinality()), (_dafny.Zero).Minus((_122_v32).F8()), _122_v32.F4(), (_122_v32).F9(), _83_globalState), (_122_v32).F7(), _83_globalState))
			var _rhs17 _dafny.Int = (func() _dafny.Int {
				if !((_122_v32).F9()) {
					return (_203_v86).Cardinality()
				}
				return ((_dafny.SetOf(false)).Intersection(_196_v79)).Cardinality()
			})()
			_ = _rhs17
			var _rhs18 _dafny.Map = _236_v109
			_ = _rhs18
			var _lhs8 *GlobalState = _83_globalState
			_ = _lhs8
			_lhs8.F1 = _rhs17
			_236_v109 = _rhs18
			(_119_v30).M4((_203_v86).Cardinality(), _203_v86, _83_globalState)
		} else if _source2.Is_DC21() {
			var _238___mcc_h12 _dafny.Sequence = _source2.Get_().(D7_DC21).Cf39
			_ = _238___mcc_h12
			var _239___mcc_h13 T2 = _source2.Get_().(D7_DC21).Cf40
			_ = _239___mcc_h13
			var _240_cf40 T2 = _239___mcc_h13
			_ = _240_cf40
			var _241_cf39 _dafny.Sequence = _238___mcc_h12
			_ = _241_cf39
			var _242_v110 _dafny.Set
			_ = _242_v110
			_242_v110 = _dafny.SetOf((_122_v32).F8())
			_242_v110 = _dafny.SetOf(_84_v3, _dafny.IntOfUint32(((_122_v32).F5()).Cardinality()), (_122_v32).F8(), (_240_cf40).Fm1((_122_v32).F9(), _201_v84, false, _83_globalState))
			var _243_v111 *C4
			_ = _243_v111
			var _nw43 *C4 = New_C4_()
			_ = _nw43
			_nw43.Ctor__((_113_v24).Cardinality(), !_dafny.Companion_Sequence_.Equal(_197_v80, _197_v80), !((_122_v32).F9()), Companion_Default___.Fm3(_83_globalState), _240_cf40.F4(), _dafny.Companion_Sequence_.Update((Companion_D3_.Create_DC10_((_122_v32).F5())).Dtor_cf19(), (Companion_Default___.SafeIndex(_84_v3, _dafny.IntOfUint32(((Companion_D3_.Create_DC10_((_122_v32).F5())).Dtor_cf19()).Cardinality()))).Uint32(), _116_v27))
			_243_v111 = _nw43
			_197_v80 = _dafny.Companion_Sequence_.Concatenate(_197_v80, _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_197_v80, _197_v80), (Companion_Default___.SafeIndex((_196_v79).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_197_v80, _197_v80)).Cardinality()))).Uint32(), false))
			var _244_v112 _dafny.Map
			_ = _244_v112
			var _out9 _dafny.Map
			_ = _out9
			_out9 = (_119_v30).M2((_122_v32).F8(), _83_globalState)
			_244_v112 = _out9
		} else {
			var _245___mcc_h14 _dafny.Set = _source2.Get_().(D7_DC18).Cf33
			_ = _245___mcc_h14
			var _246_cf33 _dafny.Set = _245___mcc_h14
			_ = _246_cf33
			var _247_v113 _dafny.Map
			_ = _247_v113
			_247_v113 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F8(), (_dafny.Zero).Minus((_122_v32).F8()))
			var _248_v114 _dafny.Map
			_ = _248_v114
			_248_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F7(), (_247_v113).Cardinality())
			var _249_v115 D12
			_ = _249_v115
			_249_v115 = Companion_D12_.Create_DC31_(_201_v84)
			var _250_v116 *C3
			_ = _250_v116
			var _nw44 *C3 = New_C3_()
			_ = _nw44
			_nw44.Ctor__((_122_v32).F6(), (_122_v32).F7(), (func() _dafny.Int {
				if (_248_v114).Contains((_122_v32).F9()) {
					return (_248_v114).Get((_122_v32).F9()).(_dafny.Int)
				}
				return (_122_v32).F8()
			})(), !((_122_v32).F9()), (_122_v32).F7(), !((_dafny.IntOfInt64(575)).Cmp((_198_v81).Fm1(!((_122_v32).F6()), (_249_v115).Dtor_cf60(), true, _83_globalState)) >= 0), _200_v83.F4(), (_122_v32).F5())
			_250_v116 = _nw44
			var _251_v117 D15
			_ = _251_v117
			_251_v117 = Companion_D15_.Create_DC42_(_250_v116)
			var _rhs19 _dafny.Int = _84_v3
			_ = _rhs19
			var _rhs20 _dafny.Int = (func() _dafny.Int {
				if (_122_v32).F9() {
					return ((_122_v32).F8()).Minus((_dafny.Zero).Minus((_122_v32).F8()))
				}
				return (_dafny.Zero).Minus(((_122_v32).F8()).Minus((_122_v32).F8()))
			})()
			_ = _rhs20
			var _rhs21 *C3 = (_251_v117).Dtor_cf77()
			_ = _rhs21
			var _lhs9 *GlobalState = _83_globalState
			_ = _lhs9
			var _lhs10 *GlobalState = _83_globalState
			_ = _lhs10
			_lhs9.F1 = _rhs19
			_lhs10.F1 = _rhs20
			_250_v116 = _rhs21
			var _252_v118 _dafny.MultiSet
			_ = _252_v118
			_252_v118 = _dafny.MultiSetOf((_122_v32).F9())
			var _253_v119 _dafny.Int
			_ = _253_v119
			var _254_v120 _dafny.Int
			_ = _254_v120
			var _255_v121 _dafny.Array
			_ = _255_v121
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			var _out12 _dafny.Array
			_ = _out12
			_out10, _out11, _out12 = (_250_v116).M1((_122_v32).F6(), !((_122_v32).F6()), (_252_v118).Intersection(_252_v118), true, _83_globalState)
			_253_v119 = _out10
			_254_v120 = _out11
			_255_v121 = _out12
			_254_v120 = ((_dafny.Zero).Minus(_84_v3)).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ubva")).Cardinality()))
			_84_v3 = (_122_v32).F8()
		}
		var _256_v122 _dafny.Set
		_ = _256_v122
		var _257_v123 _dafny.Array
		_ = _257_v123
		var _258_v124 _dafny.Int
		_ = _258_v124
		var _259_v125 _dafny.Array
		_ = _259_v125
		var _out13 _dafny.Set
		_ = _out13
		var _out14 _dafny.Array
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		var _out16 _dafny.Array
		_ = _out16
		_out13, _out14, _out15, _out16 = (_117_v28).M8((_122_v32).F7(), (_122_v32).F9(), _83_globalState)
		_256_v122 = _out13
		_257_v123 = _out14
		_258_v124 = _out15
		_259_v125 = _out16
		var _260_v126 _dafny.Map
		_ = _260_v126
		_260_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F6(), _dafny.IntOfInt64(-203))
		var _261_v127 *C3
		_ = _261_v127
		var _nw45 *C3 = New_C3_()
		_ = _nw45
		_nw45.Ctor__(_80_v0, (_122_v32).F9(), _258_v124, (_122_v32).F7(), (_122_v32).F6(), false, _dafny.CodePoint('u'), (_122_v32).F5())
		_261_v127 = _nw45
		var _262_v128 _dafny.Map
		_ = _262_v128
		_262_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_261_v127, _258_v124)
		var _263_v129 _dafny.Map
		_ = _263_v129
		_263_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(420), (_dafny.Zero).Minus((_122_v32).F8()))
		var _rhs22 bool = (_122_v32).F6()
		_ = _rhs22
		var _rhs23 _dafny.Int = Companion_Default___.SafeDivisionInt((_84_v3).Plus(((_260_v126).Update(_80_v0, (_262_v128).Cardinality())).Cardinality()), Companion_Default___.SafeModuloInt(_84_v3, (_dafny.Zero).Minus((func() _dafny.Int {
			if (_263_v129).Contains(_258_v124) {
				return (_263_v129).Get(_258_v124).(_dafny.Int)
			}
			return (_122_v32).F8()
		})())))
		_ = _rhs23
		_80_v0 = _rhs22
		_84_v3 = _rhs23
		var _264_v130 D3
		_ = _264_v130
		_264_v130 = Companion_D3_.Create_DC11_(!(_261_v127.F12), (_122_v32).F8(), _dafny.SetOf(Companion_Default___.Fm7(_116_v27, _258_v124, (_122_v32).F6(), _83_globalState), true))
		var _265_v131 _dafny.Map
		_ = _265_v131
		_265_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_264_v130, _84_v3)
		var _266_v133 _dafny.MultiSet
		_ = _266_v133
		_266_v133 = _dafny.MultiSetOf(Companion_D3_.Create_DC11_((_122_v32).F6(), _258_v124, _196_v79))
		var _267_v134 _dafny.Map
		_ = _267_v134
		_267_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F7(), func() _dafny.Map {
			var _coll8 = _dafny.NewMapBuilder()
			_ = _coll8
			for _iter8 := _dafny.Iterate((_266_v133).Elements()); ; {
				_compr_8, _ok8 := _iter8()
				if !_ok8 {
					break
				}
				var _268_v132 D3
				_268_v132 = interface{}(_compr_8).(D3)
				if (_266_v133).Contains(_268_v132) {
					_coll8.Add(_268_v132, _84_v3)
				}
			}
			return _coll8.ToMap()
		}())
		_265_v131 = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_264_v130, (_122_v32).F8())).Merge((func() _dafny.Map {
			if (_267_v134).Contains((_122_v32).F9()) {
				return (_267_v134).Get((_122_v32).F9()).(_dafny.Map)
			}
			return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_264_v130, (_122_v32).F8())
		})())).Merge(_265_v131)
	} else {
		if (_122_v32).F9() {
			_80_v0 = (_84_v3).Cmp((_122_v32).F8()) >= 0
			_80_v0 = (func() bool {
				if (_203_v86).Contains(_84_v3) {
					return (_203_v86).Get(_84_v3).(bool)
				}
				return !((_dafny.IntOfUint32(((_122_v32).F5()).Cardinality())).Cmp(_84_v3) == 0)
			})()
			var _269_v135 _dafny.Map
			_ = _269_v135
			_269_v135 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F6(), _196_v79)
			var _270_v136 _dafny.Map
			_ = _270_v136
			_270_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_200_v83).F5()).Cardinality()), _269_v135)
			_203_v86 = (_203_v86).Update(Companion_Default___.SafeDivisionInt(_84_v3, ((func() _dafny.Map {
				if (_270_v136).Contains(_dafny.IntOfUint32((_197_v80).Cardinality())) {
					return (_270_v136).Get(_dafny.IntOfUint32((_197_v80).Cardinality())).(_dafny.Map)
				}
				return _269_v135
			})()).Cardinality()), (_122_v32).F7())
			var _271_v137 *C3
			_ = _271_v137
			var _nw46 *C3 = New_C3_()
			_ = _nw46
			_nw46.Ctor__((_122_v32).F9(), (_122_v32).F6(), (_dafny.IntOfUint32((_197_v80).Cardinality())).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(370))).Uint32(), func(coer24 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg24 _dafny.Int) interface{} {
					return coer24(arg24)
				}
			}((func(_272_v27 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_273_i12 _dafny.Int) _dafny.CodePoint {
					return _272_v27
				}
			})(_116_v27)))).Cardinality())), !((_197_v80).Select((Companion_Default___.SafeIndex((_122_v32).F8(), _dafny.IntOfUint32((_197_v80).Cardinality()))).Uint32()).(bool)) || ((_122_v32).F6()), (_122_v32).F9(), false, _116_v27, _dafny.Companion_Sequence_.Concatenate((_200_v83).F5(), (_122_v32).F5()))
			_271_v137 = _nw46
			var _274_v138 _dafny.Map
			_ = _274_v138
			_274_v138 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(!((_122_v32).F6())), _84_v3)
			var _275_v139 D5
			_ = _275_v139
			_275_v139 = Companion_D5_.Create_DC15_((_119_v30).F10(), _80_v0, _274_v138, _dafny.IntOfUint32(((_122_v32).F5()).Cardinality()))
			var _276_v140 _dafny.Set
			_ = _276_v140
			_276_v140 = _dafny.SetOf(_122_v32.F4())
			var _277_v141 _dafny.Set
			_ = _277_v141
			_277_v141 = _dafny.SetOf((_276_v140).Cardinality())
			var _278_v142 _dafny.Array
			_ = _278_v142
			var _nwElement0_4 _dafny.Set = _277_v141
			_ = _nwElement0_4
			var _nw47 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(19))
			_ = _nw47
			(_nw47).ArraySet1(_nwElement0_4, 0)
			(_nw47).ArraySet1(_277_v141, 1)
			(_nw47).ArraySet1(_277_v141, 2)
			(_nw47).ArraySet1(_277_v141, 3)
			(_nw47).ArraySet1(_277_v141, 4)
			(_nw47).ArraySet1(_277_v141, 5)
			(_nw47).ArraySet1(_277_v141, 6)
			(_nw47).ArraySet1(_277_v141, 7)
			(_nw47).ArraySet1(_277_v141, 8)
			(_nw47).ArraySet1(_277_v141, 9)
			(_nw47).ArraySet1(_277_v141, 10)
			(_nw47).ArraySet1(_277_v141, 11)
			(_nw47).ArraySet1(_dafny.SetOf((_122_v32).F8()), 12)
			(_nw47).ArraySet1(_dafny.SetOf((_122_v32).F8()), 13)
			(_nw47).ArraySet1(_277_v141, 14)
			(_nw47).ArraySet1(_277_v141, 15)
			(_nw47).ArraySet1(_277_v141, 16)
			(_nw47).ArraySet1(_277_v141, 17)
			(_nw47).ArraySet1(_dafny.SetOf((_122_v32).F8(), _84_v3, (_122_v32).F8()), 18)
			_278_v142 = _nw47
			var _279_v143 D14
			_ = _279_v143
			_279_v143 = Companion_D14_.Create_DC39_(_278_v142)
			var _pat_let_tv5 = _278_v142
			_ = _pat_let_tv5
			var _280_v144 _dafny.Map
			_ = _280_v144
			_280_v144 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_275_v139).Dtor_cf29(), func(_pat_let4_0 D14) D14 {
				return func(_281_dt__update__tmp_h0 D14) D14 {
					return func(_pat_let5_0 _dafny.Array) D14 {
						return func(_282_dt__update_hcf72_h0 _dafny.Array) D14 {
							return Companion_D14_.Create_DC39_(_282_dt__update_hcf72_h0)
						}(_pat_let5_0)
					}(_pat_let_tv5)
				}(_pat_let4_0)
			}(_279_v143))
			_280_v144 = (_280_v144).Update((_84_v3).Minus((_122_v32).F8()), Companion_D14_.Create_DC39_(_278_v142))
		} else {
			_80_v0 = _80_v0
			var _283_v145 _dafny.Array
			_ = _283_v145
			var _nw48 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(13))
			_ = _nw48
			_283_v145 = _nw48
			var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_283_v145), 0))
			_ = _index26
			(_283_v145).ArraySet1((_200_v83).F5(), (_index26).Int())
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(39), _dafny.ArrayLen((_283_v145), 0))
			_ = _index27
			(_283_v145).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(183))).Uint32(), func(coer25 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg25 _dafny.Int) interface{} {
					return coer25(arg25)
				}
			}((func(_284_v32 T3) func(_dafny.Int) _dafny.CodePoint {
				return func(_285_i13 _dafny.Int) _dafny.CodePoint {
					return _284_v32.F4()
				}
			})(_122_v32))), _85_v4), (_index27).Int())
			var _286_v146 *C0
			_ = _286_v146
			var _nw49 *C0 = New_C0_()
			_ = _nw49
			_nw49.Ctor__((_122_v32).F6(), _80_v0, _122_v32.F4(), _dafny.UnicodeSeqOfUtf8Bytes("igmblqb"))
			_286_v146 = _nw49
			var _287_v147 _dafny.Map
			_ = _287_v147
			_287_v147 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_122_v32).F9(), _286_v146)
			(_83_globalState).F1 = (((_287_v147).Cardinality()).Plus((_122_v32).F8())).Minus((_122_v32).F8())
			_201_v84 = (_201_v84).Update(false, (_122_v32).F9())
			var _288_v148 D12
			_ = _288_v148
			_288_v148 = Companion_D12_.Create_DC33_(!(false), (_200_v83).F5())
			var _289_v149 _dafny.Map
			_ = _289_v149
			_289_v149 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_288_v148, (_122_v32).F7())
			_289_v149 = (_289_v149).Update(_288_v148, (_122_v32).F6())
		}
		var _290_v151 _dafny.Set
		_ = _290_v151
		_290_v151 = _dafny.SetOf(_dafny.IntOfInt64(271))
		var _291_v152 _dafny.Map
		_ = _291_v152
		_291_v152 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_122_v32).F9()), (_122_v32).F8())
		var _292_v153 _dafny.Map
		_ = _292_v153
		_292_v153 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_84_v3, (_291_v152).Cardinality())
		var _293_v155 *C0
		_ = _293_v155
		var _nw50 *C0 = New_C0_()
		_ = _nw50
		_nw50.Ctor__((_122_v32).F9(), (_122_v32).F7(), _116_v27, (_122_v32).F5())
		_293_v155 = _nw50
		var _294_v156 _dafny.Sequence
		_ = _294_v156
		_294_v156 = _dafny.SeqOf(_293_v155)
		var _295_v157 _dafny.Sequence
		_ = _295_v157
		_295_v157 = _dafny.SeqOf(_294_v156)
		var _296_v158 _dafny.MultiSet
		_ = _296_v158
		_296_v158 = _dafny.MultiSetOf((_122_v32).F9(), (_122_v32).F7(), _80_v0)
		var _297_v160 _dafny.Map
		_ = _297_v160
		_297_v160 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32(((_295_v157).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
			if (_296_v158).Contains(_80_v0) {
				return (_296_v158).Multiplicity(_80_v0)
			}
			return (_122_v32).F8()
		})(), _dafny.IntOfUint32((_295_v157).Cardinality()))).Uint32()).(_dafny.Sequence)).Cardinality()), func() _dafny.Set {
			var _coll9 = _dafny.NewBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate(((_125_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_125_v35), 0))).Int()).(_dafny.MultiSet)).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _298_v159 _dafny.Int
				_298_v159 = interface{}(_compr_9).(_dafny.Int)
				if ((_125_v35).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(124), _dafny.ArrayLen((_125_v35), 0))).Int()).(_dafny.MultiSet)).Contains(_298_v159) {
					_coll9.Add((_298_v159).Times((_dafny.MultiSetFromSeq(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(852))).Uint32(), func(coer26 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg26 _dafny.Int) interface{} {
							return coer26(arg26)
						}
					}(func(_299_i14 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(-752)
					})))).Cardinality()))
				}
			}
			return _coll9.ToSet()
		}())
		var _300_v161 _dafny.Array
		_ = _300_v161
		var _nwElement0_5 _dafny.Set = func() _dafny.Set {
			var _coll10 = _dafny.NewBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(381), _dafny.IntOfInt64(-403))); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _301_v150 _dafny.Int
				_301_v150 = interface{}(_compr_10).(_dafny.Int)
				if ((_dafny.IntOfInt64(381)).Cmp(_301_v150) <= 0) && ((_301_v150).Cmp(_dafny.IntOfInt64(-403)) < 0) {
					_coll10.Add(Companion_Default___.SafeModuloInt(_301_v150, (_dafny.Zero).Minus((_122_v32).F8())))
				}
			}
			return _coll10.ToSet()
		}()
		_ = _nwElement0_5
		var _nw51 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(15))
		_ = _nw51
		(_nw51).ArraySet1(_nwElement0_5, 0)
		(_nw51).ArraySet1(_290_v151, 1)
		(_nw51).ArraySet1(_290_v151, 2)
		(_nw51).ArraySet1(_290_v151, 3)
		(_nw51).ArraySet1(_290_v151, 4)
		(_nw51).ArraySet1(func() _dafny.Set {
			var _coll11 = _dafny.NewBuilder()
			_ = _coll11
			for _iter11 := _dafny.Iterate((_292_v153).Keys().Elements()); ; {
				_compr_11, _ok11 := _iter11()
				if !_ok11 {
					break
				}
				var _302_v154 _dafny.Int
				_302_v154 = interface{}(_compr_11).(_dafny.Int)
				if (_292_v153).Contains(_302_v154) {
					_coll11.Add(Companion_Default___.SafeModuloInt(_302_v154, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ucmoqf")).Cardinality()))).Cardinality()))
				}
			}
			return _coll11.ToSet()
		}(), 5)
		(_nw51).ArraySet1(_290_v151, 6)
		(_nw51).ArraySet1(_dafny.SetOf((_122_v32).F8()), 7)
		(_nw51).ArraySet1(_dafny.SetOf((_122_v32).F8()), 8)
		(_nw51).ArraySet1(_290_v151, 9)
		(_nw51).ArraySet1(_290_v151, 10)
		(_nw51).ArraySet1(_290_v151, 11)
		(_nw51).ArraySet1(_290_v151, 12)
		(_nw51).ArraySet1((func() _dafny.Set {
			if (_297_v160).Contains((_122_v32).F8()) {
				return (_297_v160).Get((_122_v32).F8()).(_dafny.Set)
			}
			return _290_v151
		})(), 13)
		(_nw51).ArraySet1(_290_v151, 14)
		_300_v161 = _nw51
		var _303_v162 D14
		_ = _303_v162
		_303_v162 = Companion_D14_.Create_DC39_(_300_v161)
		var _source3 D14 = _303_v162
		_ = _source3
		if _source3.Is_DC40() {
			var _304___mcc_h15 _dafny.Int = _source3.Get_().(D14_DC40).Cf73
			_ = _304___mcc_h15
			var _305___mcc_h16 bool = _source3.Get_().(D14_DC40).Cf74
			_ = _305___mcc_h16
			var _306___mcc_h17 _dafny.Sequence = _source3.Get_().(D14_DC40).Cf75
			_ = _306___mcc_h17
			var _307_cf75 _dafny.Sequence = _306___mcc_h17
			_ = _307_cf75
			var _308_cf74 bool = _305___mcc_h16
			_ = _308_cf74
			var _309_cf73 _dafny.Int = _304___mcc_h15
			_ = _309_cf73
			Companion_Default___.M0(_83_globalState)
			_80_v0 = ((_122_v32).F8()).Cmp(Companion_Default___.SafeModuloInt(_84_v3, _309_cf73)) != 0
			_308_cf74 = (_122_v32).F9()
			var _rhs24 bool = (_122_v32).F6()
			_ = _rhs24
			var _rhs25 _dafny.Int = (_122_v32).F8()
			_ = _rhs25
			var _rhs26 bool = (_308_cf74) || (!((_122_v32).F6()))
			_ = _rhs26
			var _rhs27 _dafny.Int = ((_122_v32).F8()).Minus(_309_cf73)
			_ = _rhs27
			var _lhs11 *GlobalState = _83_globalState
			_ = _lhs11
			_80_v0 = _rhs24
			_lhs11.F1 = _rhs25
			_308_cf74 = _rhs26
			_84_v3 = _rhs27
		} else if _source3.Is_DC39() {
			var _310___mcc_h18 _dafny.Array = _source3.Get_().(D14_DC39).Cf72
			_ = _310___mcc_h18
			var _311_cf72 _dafny.Array = _310___mcc_h18
			_ = _311_cf72
			Companion_Default___.M0(_83_globalState)
			var _312_v163 _dafny.Sequence
			_ = _312_v163
			_312_v163 = _dafny.SeqOf(_117_v28)
			var _313_v164 _dafny.Array
			_ = _313_v164
			var _nwElement0_6 *C1 = _117_v28
			_ = _nwElement0_6
			var _nw52 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(14))
			_ = _nw52
			(_nw52).ArraySet1(_nwElement0_6, 0)
			(_nw52).ArraySet1(_117_v28, 1)
			(_nw52).ArraySet1(_117_v28, 2)
			(_nw52).ArraySet1(_117_v28, 3)
			(_nw52).ArraySet1((_312_v163).Select((Companion_Default___.SafeIndex((_122_v32).F8(), _dafny.IntOfUint32((_312_v163).Cardinality()))).Uint32()).(*C1), 4)
			(_nw52).ArraySet1(_117_v28, 5)
			(_nw52).ArraySet1(_117_v28, 6)
			(_nw52).ArraySet1(_117_v28, 7)
			(_nw52).ArraySet1((func() *C1 {
				if (_122_v32).F9() {
					return _117_v28
				}
				return _117_v28
			})(), 8)
			(_nw52).ArraySet1(_117_v28, 9)
			(_nw52).ArraySet1(_117_v28, 10)
			(_nw52).ArraySet1(_117_v28, 11)
			(_nw52).ArraySet1((Companion_D10_.Create_DC28_((_122_v32).F9(), (_122_v32).F8(), _80_v0, (_122_v32).F7(), _117_v28)).Dtor_cf54(), 12)
			(_nw52).ArraySet1(_117_v28, 13)
			_313_v164 = _nw52
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_313_v164), 0))
			_ = _index28
			(_313_v164).ArraySet1(_117_v28, (_index28).Int())
			var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(259), _dafny.ArrayLen((_313_v164), 0))
			_ = _index29
			(_313_v164).ArraySet1(_117_v28, (_index29).Int())
			(_83_globalState).F1 = _84_v3
			_303_v162 = _303_v162
		} else {
			var _314___mcc_h19 D14 = _source3.Get_().(D14_DC41).Cf76
			_ = _314___mcc_h19
			var _315_cf76 D14 = _314___mcc_h19
			_ = _315_cf76
			var _316_v165 _dafny.Array
			_ = _316_v165
			var _len0_7 _dafny.Int = _dafny.IntOfInt64(12)
			_ = _len0_7
			var _nw53 _dafny.Array
			_ = _nw53
			if _len0_7.Cmp(_dafny.Zero) == 0 {
				_nw53 = _dafny.NewArray(_len0_7)
			} else {
				var _init7 func(_dafny.Int) _dafny.Sequence = (func(_317_v80 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
					return func(_318_i15 _dafny.Int) _dafny.Sequence {
						return _dafny.Companion_Sequence_.Concatenate(_317_v80, _317_v80)
					}
				})(_197_v80)
				_ = _init7
				var _element0_7 = _init7(_dafny.Zero)
				_ = _element0_7
				_nw53 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
				(_nw53).ArraySet1(_element0_7, 0)
				var _nativeLen0_7 = (_len0_7).Int()
				_ = _nativeLen0_7
				for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
					(_nw53).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
				}
			}
			_316_v165 = _nw53
			_316_v165 = _316_v165
			_80_v0 = (_122_v32).F7()
			var _319_v166 _dafny.Map
			_ = _319_v166
			_319_v166 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_122_v32).F8())).Cardinality()), _197_v80)
			_197_v80 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
				if (_319_v166).Contains((_122_v32).F8()) {
					return (_319_v166).Get((_122_v32).F8()).(_dafny.Sequence)
				}
				return _197_v80
			})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_200_v83).F5()).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_319_v166).Contains((_122_v32).F8()) {
					return (_319_v166).Get((_122_v32).F8()).(_dafny.Sequence)
				}
				return _197_v80
			})()).Cardinality()))).Uint32(), (_122_v32).F9()), _197_v80), _197_v80)
			_80_v0 = (_122_v32).F9()
		}
		var _320_v167 D2
		_ = _320_v167
		_320_v167 = Companion_D2_.Create_DC8_(_84_v3, (_122_v32).F8())
		var _321_v168 D2
		_ = _321_v168
		_321_v168 = Companion_D2_.Create_DC9_(_320_v167)
		_321_v168 = _321_v168
		var _322_v169 D5
		_ = _322_v169
		_322_v169 = Companion_D5_.Create_DC14_(_81_v1)
		var _323_v170 _dafny.MultiSet
		_ = _323_v170
		_323_v170 = _dafny.MultiSetOf(_322_v169)
		_323_v170 = _323_v170
		_80_v0 = (_122_v32).F9()
	}
	var _324_v171 D0
	_ = _324_v171
	_324_v171 = Companion_D0_.Create_DC0_()
	var _325_v172 _dafny.Map
	_ = _325_v172
	_325_v172 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v171, (_122_v32).Fm1(true, _201_v84, (_122_v32).F9(), _83_globalState))
	_325_v172 = (_325_v172).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_324_v171, _dafny.IntOfInt64(763)))
	_dafny.Print(_80_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v1).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_81_v1).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_82_v2).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_83_globalState.F1)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_83_globalState).F2())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_83_globalState).F3()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_84_v3)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_85_v4.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_86_v5).Equals(_dafny.MultiSetOf(_dafny.SetOf(true, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_87_v6).Equals(_dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("ikmk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_88_v7, _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("ikmk"))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_90_i0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_113_v24).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_114_v25, _dafny.SeqOf(_dafny.IntOfInt64(372), _dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_115_v26, _dafny.SeqOf(_dafny.MultiSetOf(_dafny.IntOfInt64(372), _dafny.IntOfInt64(287)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_116_v27)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_118_v29).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_119_v30).F10()).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_119_v30).F11()).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_119_v30).F11()).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_121_v31).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v32).F8())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v32).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v32).F6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v32).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_122_v32.F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_122_v32).F5().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_123_v33).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v34).Dtor_cf4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_124_v34).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_124_v34).Dtor_cf6(), _dafny.SeqOf(_dafny.IntOfInt64(372), _dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.Zero).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(372), _dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.One).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(8)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(9)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(10)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(11)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(12)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(13)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(14)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(15)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(16)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(17)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-78))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(18)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(19)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(20)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(-287), _dafny.IntOfInt64(287), _dafny.IntOfInt64(287), _dafny.IntOfInt64(2))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(21)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(22)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(23)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(4), _dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(24)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(25)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(26)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_125_v35).ArrayGet1((_dafny.IntOfInt64(27)).Int()).(_dafny.MultiSet)).Equals(_dafny.MultiSetOf(_dafny.One, _dafny.IntOfInt64(5), _dafny.IntOfInt64(-415), _dafny.IntOfInt64(287), _dafny.IntOfInt64(287))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_196_v79).Equals(_dafny.SetOf(false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_197_v80, _dafny.SeqOf(true, false, false, true, false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_199_v82).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_200_v83.F4())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_200_v83).F5().VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_201_v84).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_203_v86).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(287), true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_325_v172).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC0_(), _dafny.IntOfInt64(763))))
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

type D0_DC0 struct {
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_() D0 {
	return D0{D0_DC0{}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf0 bool
	Cf1 _dafny.Int
	Cf2 bool
	Cf3 _dafny.Int
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf0 bool, Cf1 _dafny.Int, Cf2 bool, Cf3 _dafny.Int) D0 {
	return D0{D0_DC1{Cf0, Cf1, Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.Int
	Cf5 bool
	Cf6 _dafny.Sequence
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.Int, Cf5 bool, Cf6 _dafny.Sequence) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_()
}

func (_this D0) Dtor_cf0() bool {
	return _this.Get_().(D0_DC1).Cf0
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

func (_this D0) Dtor_cf4() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() bool {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() _dafny.Sequence {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ", " + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D0) Equals(other D0) bool {
	switch data1 := _this.Get_().(type) {
	case D0_DC0:
		{
			_, ok := other.Get_().(D0_DC0)
			return ok
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf0 == data2.Cf0 && data1.Cf1.Cmp(data2.Cf1) == 0 && data1.Cf2 == data2.Cf2 && data1.Cf3.Cmp(data2.Cf3) == 0
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Cmp(data2.Cf4) == 0 && data1.Cf5 == data2.Cf5 && data1.Cf6.Equals(data2.Cf6)
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

type D1_DC4 struct {
	Cf8  _dafny.CodePoint
	Cf9  _dafny.Int
	Cf10 _dafny.CodePoint
	Cf11 _dafny.Int
	Cf12 bool
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.CodePoint, Cf9 _dafny.Int, Cf10 _dafny.CodePoint, Cf11 _dafny.Int, Cf12 bool) D1 {
	return D1{D1_DC4{Cf8, Cf9, Cf10, Cf11, Cf12}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

type D1_DC5 struct {
	Cf13 _dafny.Int
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf13 _dafny.Int) D1 {
	return D1{D1_DC5{Cf13}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC3 struct {
	Cf7 _dafny.Array
}

func (D1_DC3) isD1() {}

func (CompanionStruct_D1_) Create_DC3_(Cf7 _dafny.Array) D1 {
	return D1{D1_DC3{Cf7}}
}

func (_this D1) Is_DC3() bool {
	_, ok := _this.Get_().(D1_DC3)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC4_(_dafny.CodePoint('D'), _dafny.Zero, _dafny.CodePoint('D'), _dafny.Zero, false)
}

func (_this D1) Dtor_cf8() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) Dtor_cf9() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf9
}

func (_this D1) Dtor_cf10() _dafny.CodePoint {
	return _this.Get_().(D1_DC4).Cf10
}

func (_this D1) Dtor_cf11() _dafny.Int {
	return _this.Get_().(D1_DC4).Cf11
}

func (_this D1) Dtor_cf12() bool {
	return _this.Get_().(D1_DC4).Cf12
}

func (_this D1) Dtor_cf13() _dafny.Int {
	return _this.Get_().(D1_DC5).Cf13
}

func (_this D1) Dtor_cf7() _dafny.Array {
	return _this.Get_().(D1_DC3).Cf7
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC4:
		{
			return "D1.DC4" + "(" + _dafny.String(data.Cf8) + ", " + _dafny.String(data.Cf9) + ", " + _dafny.String(data.Cf10) + ", " + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ")"
		}
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf13) + ")"
		}
	case D1_DC3:
		{
			return "D1.DC3" + "(" + _dafny.String(data.Cf7) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D1) Equals(other D1) bool {
	switch data1 := _this.Get_().(type) {
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf8 == data2.Cf8 && data1.Cf9.Cmp(data2.Cf9) == 0 && data1.Cf10 == data2.Cf10 && data1.Cf11.Cmp(data2.Cf11) == 0 && data1.Cf12 == data2.Cf12
		}
	case D1_DC5:
		{
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf13.Cmp(data2.Cf13) == 0
		}
	case D1_DC3:
		{
			data2, ok := other.Get_().(D1_DC3)
			return ok && data1.Cf7 == data2.Cf7
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

type D2_DC7 struct {
	Cf15 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf15 bool) D2 {
	return D2{D2_DC7{Cf15}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC8 struct {
	Cf16 _dafny.Int
	Cf17 _dafny.Int
}

func (D2_DC8) isD2() {}

func (CompanionStruct_D2_) Create_DC8_(Cf16 _dafny.Int, Cf17 _dafny.Int) D2 {
	return D2{D2_DC8{Cf16, Cf17}}
}

func (_this D2) Is_DC8() bool {
	_, ok := _this.Get_().(D2_DC8)
	return ok
}

type D2_DC6 struct {
	Cf14 _dafny.Map
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf14 _dafny.Map) D2 {
	return D2{D2_DC6{Cf14}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

type D2_DC9 struct {
	Cf18 D2
}

func (D2_DC9) isD2() {}

func (CompanionStruct_D2_) Create_DC9_(Cf18 D2) D2 {
	return D2{D2_DC9{Cf18}}
}

func (_this D2) Is_DC9() bool {
	_, ok := _this.Get_().(D2_DC9)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false)
}

func (_this D2) Dtor_cf15() bool {
	return _this.Get_().(D2_DC7).Cf15
}

func (_this D2) Dtor_cf16() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf16
}

func (_this D2) Dtor_cf17() _dafny.Int {
	return _this.Get_().(D2_DC8).Cf17
}

func (_this D2) Dtor_cf14() _dafny.Map {
	return _this.Get_().(D2_DC6).Cf14
}

func (_this D2) Dtor_cf18() D2 {
	return _this.Get_().(D2_DC9).Cf18
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D2_DC8:
		{
			return "D2.DC8" + "(" + _dafny.String(data.Cf16) + ", " + _dafny.String(data.Cf17) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D2_DC9:
		{
			return "D2.DC9" + "(" + _dafny.String(data.Cf18) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D2) Equals(other D2) bool {
	switch data1 := _this.Get_().(type) {
	case D2_DC7:
		{
			data2, ok := other.Get_().(D2_DC7)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D2_DC8:
		{
			data2, ok := other.Get_().(D2_DC8)
			return ok && data1.Cf16.Cmp(data2.Cf16) == 0 && data1.Cf17.Cmp(data2.Cf17) == 0
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf14.Equals(data2.Cf14)
		}
	case D2_DC9:
		{
			data2, ok := other.Get_().(D2_DC9)
			return ok && data1.Cf18.Equals(data2.Cf18)
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

type D3_DC11 struct {
	Cf20 bool
	Cf21 _dafny.Int
	Cf22 _dafny.Set
}

func (D3_DC11) isD3() {}

func (CompanionStruct_D3_) Create_DC11_(Cf20 bool, Cf21 _dafny.Int, Cf22 _dafny.Set) D3 {
	return D3{D3_DC11{Cf20, Cf21, Cf22}}
}

func (_this D3) Is_DC11() bool {
	_, ok := _this.Get_().(D3_DC11)
	return ok
}

type D3_DC10 struct {
	Cf19 _dafny.Sequence
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf19 _dafny.Sequence) D3 {
	return D3{D3_DC10{Cf19}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

type D3_DC12 struct {
	Cf23 D3
}

func (D3_DC12) isD3() {}

func (CompanionStruct_D3_) Create_DC12_(Cf23 D3) D3 {
	return D3{D3_DC12{Cf23}}
}

func (_this D3) Is_DC12() bool {
	_, ok := _this.Get_().(D3_DC12)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC11_(false, _dafny.Zero, _dafny.EmptySet)
}

func (_this D3) Dtor_cf20() bool {
	return _this.Get_().(D3_DC11).Cf20
}

func (_this D3) Dtor_cf21() _dafny.Int {
	return _this.Get_().(D3_DC11).Cf21
}

func (_this D3) Dtor_cf22() _dafny.Set {
	return _this.Get_().(D3_DC11).Cf22
}

func (_this D3) Dtor_cf19() _dafny.Sequence {
	return _this.Get_().(D3_DC10).Cf19
}

func (_this D3) Dtor_cf23() D3 {
	return _this.Get_().(D3_DC12).Cf23
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC11:
		{
			return "D3.DC11" + "(" + _dafny.String(data.Cf20) + ", " + _dafny.String(data.Cf21) + ", " + _dafny.String(data.Cf22) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + data.Cf19.VerbatimString(true) + ")"
		}
	case D3_DC12:
		{
			return "D3.DC12" + "(" + _dafny.String(data.Cf23) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC11:
		{
			data2, ok := other.Get_().(D3_DC11)
			return ok && data1.Cf20 == data2.Cf20 && data1.Cf21.Cmp(data2.Cf21) == 0 && data1.Cf22.Equals(data2.Cf22)
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf19.Equals(data2.Cf19)
		}
	case D3_DC12:
		{
			data2, ok := other.Get_().(D3_DC12)
			return ok && data1.Cf23.Equals(data2.Cf23)
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
	Cf24 _dafny.Sequence
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf24 _dafny.Sequence) D4 {
	return D4{D4_DC13{Cf24}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() _dafny.Sequence {
	return _dafny.EmptySeq
}

func (_this D4) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D4_DC13).Cf24
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf24) + ")"
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
			return ok && data1.Cf24.Equals(data2.Cf24)
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

type D5_DC15 struct {
	Cf26 _dafny.Array
	Cf27 bool
	Cf28 _dafny.Map
	Cf29 _dafny.Int
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf26 _dafny.Array, Cf27 bool, Cf28 _dafny.Map, Cf29 _dafny.Int) D5 {
	return D5{D5_DC15{Cf26, Cf27, Cf28, Cf29}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC14 struct {
	Cf25 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf25 _dafny.Array) D5 {
	return D5{D5_DC14{Cf25}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), false, _dafny.EmptyMap, _dafny.Zero)
}

func (_this D5) Dtor_cf26() _dafny.Array {
	return _this.Get_().(D5_DC15).Cf26
}

func (_this D5) Dtor_cf27() bool {
	return _this.Get_().(D5_DC15).Cf27
}

func (_this D5) Dtor_cf28() _dafny.Map {
	return _this.Get_().(D5_DC15).Cf28
}

func (_this D5) Dtor_cf29() _dafny.Int {
	return _this.Get_().(D5_DC15).Cf29
}

func (_this D5) Dtor_cf25() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf25
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf26) + ", " + _dafny.String(data.Cf27) + ", " + _dafny.String(data.Cf28) + ", " + _dafny.String(data.Cf29) + ")"
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
	case D5_DC15:
		{
			data2, ok := other.Get_().(D5_DC15)
			return ok && data1.Cf26 == data2.Cf26 && data1.Cf27 == data2.Cf27 && data1.Cf28.Equals(data2.Cf28) && data1.Cf29.Cmp(data2.Cf29) == 0
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf25 == data2.Cf25
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

type D6_DC17 struct {
	Cf31 _dafny.Sequence
	Cf32 _dafny.Sequence
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf31 _dafny.Sequence, Cf32 _dafny.Sequence) D6 {
	return D6{D6_DC17{Cf31, Cf32}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

type D6_DC16 struct {
	Cf30 _dafny.MultiSet
}

func (D6_DC16) isD6() {}

func (CompanionStruct_D6_) Create_DC16_(Cf30 _dafny.MultiSet) D6 {
	return D6{D6_DC16{Cf30}}
}

func (_this D6) Is_DC16() bool {
	_, ok := _this.Get_().(D6_DC16)
	return ok
}

func (CompanionStruct_D6_) Default() D6 {
	return Companion_D6_.Create_DC17_(_dafny.EmptySeq, _dafny.EmptySeq)
}

func (_this D6) Dtor_cf31() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf31
}

func (_this D6) Dtor_cf32() _dafny.Sequence {
	return _this.Get_().(D6_DC17).Cf32
}

func (_this D6) Dtor_cf30() _dafny.MultiSet {
	return _this.Get_().(D6_DC16).Cf30
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + data.Cf31.VerbatimString(true) + ", " + data.Cf32.VerbatimString(true) + ")"
		}
	case D6_DC16:
		{
			return "D6.DC16" + "(" + _dafny.String(data.Cf30) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D6) Equals(other D6) bool {
	switch data1 := _this.Get_().(type) {
	case D6_DC17:
		{
			data2, ok := other.Get_().(D6_DC17)
			return ok && data1.Cf31.Equals(data2.Cf31) && data1.Cf32.Equals(data2.Cf32)
		}
	case D6_DC16:
		{
			data2, ok := other.Get_().(D6_DC16)
			return ok && data1.Cf30.Equals(data2.Cf30)
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

type D7_DC19 struct {
	Cf34 bool
	Cf35 _dafny.Sequence
	Cf36 bool
	Cf37 bool
	Cf38 _dafny.Int
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf34 bool, Cf35 _dafny.Sequence, Cf36 bool, Cf37 bool, Cf38 _dafny.Int) D7 {
	return D7{D7_DC19{Cf34, Cf35, Cf36, Cf37, Cf38}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC20 struct {
}

func (D7_DC20) isD7() {}

func (CompanionStruct_D7_) Create_DC20_() D7 {
	return D7{D7_DC20{}}
}

func (_this D7) Is_DC20() bool {
	_, ok := _this.Get_().(D7_DC20)
	return ok
}

type D7_DC21 struct {
	Cf39 _dafny.Sequence
	Cf40 T2
}

func (D7_DC21) isD7() {}

func (CompanionStruct_D7_) Create_DC21_(Cf39 _dafny.Sequence, Cf40 T2) D7 {
	return D7{D7_DC21{Cf39, Cf40}}
}

func (_this D7) Is_DC21() bool {
	_, ok := _this.Get_().(D7_DC21)
	return ok
}

type D7_DC18 struct {
	Cf33 _dafny.Set
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf33 _dafny.Set) D7 {
	return D7{D7_DC18{Cf33}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(false, _dafny.EmptySeq, false, false, _dafny.Zero)
}

func (_this D7) Dtor_cf34() bool {
	return _this.Get_().(D7_DC19).Cf34
}

func (_this D7) Dtor_cf35() _dafny.Sequence {
	return _this.Get_().(D7_DC19).Cf35
}

func (_this D7) Dtor_cf36() bool {
	return _this.Get_().(D7_DC19).Cf36
}

func (_this D7) Dtor_cf37() bool {
	return _this.Get_().(D7_DC19).Cf37
}

func (_this D7) Dtor_cf38() _dafny.Int {
	return _this.Get_().(D7_DC19).Cf38
}

func (_this D7) Dtor_cf39() _dafny.Sequence {
	return _this.Get_().(D7_DC21).Cf39
}

func (_this D7) Dtor_cf40() T2 {
	return _this.Get_().(D7_DC21).Cf40
}

func (_this D7) Dtor_cf33() _dafny.Set {
	return _this.Get_().(D7_DC18).Cf33
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf34) + ", " + data.Cf35.VerbatimString(true) + ", " + _dafny.String(data.Cf36) + ", " + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ")"
		}
	case D7_DC20:
		{
			return "D7.DC20"
		}
	case D7_DC21:
		{
			return "D7.DC21" + "(" + _dafny.String(data.Cf39) + ", " + _dafny.String(data.Cf40) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf33) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D7) Equals(other D7) bool {
	switch data1 := _this.Get_().(type) {
	case D7_DC19:
		{
			data2, ok := other.Get_().(D7_DC19)
			return ok && data1.Cf34 == data2.Cf34 && data1.Cf35.Equals(data2.Cf35) && data1.Cf36 == data2.Cf36 && data1.Cf37 == data2.Cf37 && data1.Cf38.Cmp(data2.Cf38) == 0
		}
	case D7_DC20:
		{
			_, ok := other.Get_().(D7_DC20)
			return ok
		}
	case D7_DC21:
		{
			data2, ok := other.Get_().(D7_DC21)
			return ok && data1.Cf39.Equals(data2.Cf39) && _dafny.AreEqual(data1.Cf40, data2.Cf40)
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf33.Equals(data2.Cf33)
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

type D8_DC23 struct {
	Cf42 _dafny.Int
	Cf43 bool
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf42 _dafny.Int, Cf43 bool) D8 {
	return D8{D8_DC23{Cf42, Cf43}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

type D8_DC22 struct {
	Cf41 _dafny.Sequence
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf41 _dafny.Sequence) D8 {
	return D8{D8_DC22{Cf41}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC23_(_dafny.Zero, false)
}

func (_this D8) Dtor_cf42() _dafny.Int {
	return _this.Get_().(D8_DC23).Cf42
}

func (_this D8) Dtor_cf43() bool {
	return _this.Get_().(D8_DC23).Cf43
}

func (_this D8) Dtor_cf41() _dafny.Sequence {
	return _this.Get_().(D8_DC22).Cf41
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf42) + ", " + _dafny.String(data.Cf43) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf41) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D8) Equals(other D8) bool {
	switch data1 := _this.Get_().(type) {
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
			return ok && data1.Cf42.Cmp(data2.Cf42) == 0 && data1.Cf43 == data2.Cf43
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf41.Equals(data2.Cf41)
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

type D9_DC25 struct {
	Cf45 bool
	Cf46 bool
	Cf47 _dafny.Int
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf45 bool, Cf46 bool, Cf47 _dafny.Int) D9 {
	return D9{D9_DC25{Cf45, Cf46, Cf47}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC24 struct {
	Cf44 _dafny.Sequence
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf44 _dafny.Sequence) D9 {
	return D9{D9_DC24{Cf44}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(false, false, _dafny.Zero)
}

func (_this D9) Dtor_cf45() bool {
	return _this.Get_().(D9_DC25).Cf45
}

func (_this D9) Dtor_cf46() bool {
	return _this.Get_().(D9_DC25).Cf46
}

func (_this D9) Dtor_cf47() _dafny.Int {
	return _this.Get_().(D9_DC25).Cf47
}

func (_this D9) Dtor_cf44() _dafny.Sequence {
	return _this.Get_().(D9_DC24).Cf44
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf45) + ", " + _dafny.String(data.Cf46) + ", " + _dafny.String(data.Cf47) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf44) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D9) Equals(other D9) bool {
	switch data1 := _this.Get_().(type) {
	case D9_DC25:
		{
			data2, ok := other.Get_().(D9_DC25)
			return ok && data1.Cf45 == data2.Cf45 && data1.Cf46 == data2.Cf46 && data1.Cf47.Cmp(data2.Cf47) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
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

type D10_DC27 struct {
	Cf49 _dafny.Int
}

func (D10_DC27) isD10() {}

func (CompanionStruct_D10_) Create_DC27_(Cf49 _dafny.Int) D10 {
	return D10{D10_DC27{Cf49}}
}

func (_this D10) Is_DC27() bool {
	_, ok := _this.Get_().(D10_DC27)
	return ok
}

type D10_DC28 struct {
	Cf50 bool
	Cf51 _dafny.Int
	Cf52 bool
	Cf53 bool
	Cf54 *C1
}

func (D10_DC28) isD10() {}

func (CompanionStruct_D10_) Create_DC28_(Cf50 bool, Cf51 _dafny.Int, Cf52 bool, Cf53 bool, Cf54 *C1) D10 {
	return D10{D10_DC28{Cf50, Cf51, Cf52, Cf53, Cf54}}
}

func (_this D10) Is_DC28() bool {
	_, ok := _this.Get_().(D10_DC28)
	return ok
}

type D10_DC26 struct {
	Cf48 _dafny.Array
}

func (D10_DC26) isD10() {}

func (CompanionStruct_D10_) Create_DC26_(Cf48 _dafny.Array) D10 {
	return D10{D10_DC26{Cf48}}
}

func (_this D10) Is_DC26() bool {
	_, ok := _this.Get_().(D10_DC26)
	return ok
}

func (CompanionStruct_D10_) Default() D10 {
	return Companion_D10_.Create_DC27_(_dafny.Zero)
}

func (_this D10) Dtor_cf49() _dafny.Int {
	return _this.Get_().(D10_DC27).Cf49
}

func (_this D10) Dtor_cf50() bool {
	return _this.Get_().(D10_DC28).Cf50
}

func (_this D10) Dtor_cf51() _dafny.Int {
	return _this.Get_().(D10_DC28).Cf51
}

func (_this D10) Dtor_cf52() bool {
	return _this.Get_().(D10_DC28).Cf52
}

func (_this D10) Dtor_cf53() bool {
	return _this.Get_().(D10_DC28).Cf53
}

func (_this D10) Dtor_cf54() *C1 {
	return _this.Get_().(D10_DC28).Cf54
}

func (_this D10) Dtor_cf48() _dafny.Array {
	return _this.Get_().(D10_DC26).Cf48
}

func (_this D10) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D10_DC27:
		{
			return "D10.DC27" + "(" + _dafny.String(data.Cf49) + ")"
		}
	case D10_DC28:
		{
			return "D10.DC28" + "(" + _dafny.String(data.Cf50) + ", " + _dafny.String(data.Cf51) + ", " + _dafny.String(data.Cf52) + ", " + _dafny.String(data.Cf53) + ", " + _dafny.String(data.Cf54) + ")"
		}
	case D10_DC26:
		{
			return "D10.DC26" + "(" + _dafny.String(data.Cf48) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D10) Equals(other D10) bool {
	switch data1 := _this.Get_().(type) {
	case D10_DC27:
		{
			data2, ok := other.Get_().(D10_DC27)
			return ok && data1.Cf49.Cmp(data2.Cf49) == 0
		}
	case D10_DC28:
		{
			data2, ok := other.Get_().(D10_DC28)
			return ok && data1.Cf50 == data2.Cf50 && data1.Cf51.Cmp(data2.Cf51) == 0 && data1.Cf52 == data2.Cf52 && data1.Cf53 == data2.Cf53 && data1.Cf54 == data2.Cf54
		}
	case D10_DC26:
		{
			data2, ok := other.Get_().(D10_DC26)
			return ok && data1.Cf48 == data2.Cf48
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

type D11_DC30 struct {
	Cf56 D7
	Cf57 bool
	Cf58 bool
	Cf59 bool
}

func (D11_DC30) isD11() {}

func (CompanionStruct_D11_) Create_DC30_(Cf56 D7, Cf57 bool, Cf58 bool, Cf59 bool) D11 {
	return D11{D11_DC30{Cf56, Cf57, Cf58, Cf59}}
}

func (_this D11) Is_DC30() bool {
	_, ok := _this.Get_().(D11_DC30)
	return ok
}

type D11_DC29 struct {
	Cf55 _dafny.MultiSet
}

func (D11_DC29) isD11() {}

func (CompanionStruct_D11_) Create_DC29_(Cf55 _dafny.MultiSet) D11 {
	return D11{D11_DC29{Cf55}}
}

func (_this D11) Is_DC29() bool {
	_, ok := _this.Get_().(D11_DC29)
	return ok
}

func (CompanionStruct_D11_) Default() D11 {
	return Companion_D11_.Create_DC30_(Companion_D7_.Default(), false, false, false)
}

func (_this D11) Dtor_cf56() D7 {
	return _this.Get_().(D11_DC30).Cf56
}

func (_this D11) Dtor_cf57() bool {
	return _this.Get_().(D11_DC30).Cf57
}

func (_this D11) Dtor_cf58() bool {
	return _this.Get_().(D11_DC30).Cf58
}

func (_this D11) Dtor_cf59() bool {
	return _this.Get_().(D11_DC30).Cf59
}

func (_this D11) Dtor_cf55() _dafny.MultiSet {
	return _this.Get_().(D11_DC29).Cf55
}

func (_this D11) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D11_DC30:
		{
			return "D11.DC30" + "(" + _dafny.String(data.Cf56) + ", " + _dafny.String(data.Cf57) + ", " + _dafny.String(data.Cf58) + ", " + _dafny.String(data.Cf59) + ")"
		}
	case D11_DC29:
		{
			return "D11.DC29" + "(" + _dafny.String(data.Cf55) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D11) Equals(other D11) bool {
	switch data1 := _this.Get_().(type) {
	case D11_DC30:
		{
			data2, ok := other.Get_().(D11_DC30)
			return ok && data1.Cf56.Equals(data2.Cf56) && data1.Cf57 == data2.Cf57 && data1.Cf58 == data2.Cf58 && data1.Cf59 == data2.Cf59
		}
	case D11_DC29:
		{
			data2, ok := other.Get_().(D11_DC29)
			return ok && data1.Cf55.Equals(data2.Cf55)
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

type D12_DC32 struct {
	Cf61 bool
	Cf62 _dafny.Int
}

func (D12_DC32) isD12() {}

func (CompanionStruct_D12_) Create_DC32_(Cf61 bool, Cf62 _dafny.Int) D12 {
	return D12{D12_DC32{Cf61, Cf62}}
}

func (_this D12) Is_DC32() bool {
	_, ok := _this.Get_().(D12_DC32)
	return ok
}

type D12_DC33 struct {
	Cf63 bool
	Cf64 _dafny.Sequence
}

func (D12_DC33) isD12() {}

func (CompanionStruct_D12_) Create_DC33_(Cf63 bool, Cf64 _dafny.Sequence) D12 {
	return D12{D12_DC33{Cf63, Cf64}}
}

func (_this D12) Is_DC33() bool {
	_, ok := _this.Get_().(D12_DC33)
	return ok
}

type D12_DC34 struct {
	Cf65 bool
}

func (D12_DC34) isD12() {}

func (CompanionStruct_D12_) Create_DC34_(Cf65 bool) D12 {
	return D12{D12_DC34{Cf65}}
}

func (_this D12) Is_DC34() bool {
	_, ok := _this.Get_().(D12_DC34)
	return ok
}

type D12_DC31 struct {
	Cf60 _dafny.Map
}

func (D12_DC31) isD12() {}

func (CompanionStruct_D12_) Create_DC31_(Cf60 _dafny.Map) D12 {
	return D12{D12_DC31{Cf60}}
}

func (_this D12) Is_DC31() bool {
	_, ok := _this.Get_().(D12_DC31)
	return ok
}

func (CompanionStruct_D12_) Default() D12 {
	return Companion_D12_.Create_DC32_(false, _dafny.Zero)
}

func (_this D12) Dtor_cf61() bool {
	return _this.Get_().(D12_DC32).Cf61
}

func (_this D12) Dtor_cf62() _dafny.Int {
	return _this.Get_().(D12_DC32).Cf62
}

func (_this D12) Dtor_cf63() bool {
	return _this.Get_().(D12_DC33).Cf63
}

func (_this D12) Dtor_cf64() _dafny.Sequence {
	return _this.Get_().(D12_DC33).Cf64
}

func (_this D12) Dtor_cf65() bool {
	return _this.Get_().(D12_DC34).Cf65
}

func (_this D12) Dtor_cf60() _dafny.Map {
	return _this.Get_().(D12_DC31).Cf60
}

func (_this D12) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D12_DC32:
		{
			return "D12.DC32" + "(" + _dafny.String(data.Cf61) + ", " + _dafny.String(data.Cf62) + ")"
		}
	case D12_DC33:
		{
			return "D12.DC33" + "(" + _dafny.String(data.Cf63) + ", " + data.Cf64.VerbatimString(true) + ")"
		}
	case D12_DC34:
		{
			return "D12.DC34" + "(" + _dafny.String(data.Cf65) + ")"
		}
	case D12_DC31:
		{
			return "D12.DC31" + "(" + _dafny.String(data.Cf60) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D12) Equals(other D12) bool {
	switch data1 := _this.Get_().(type) {
	case D12_DC32:
		{
			data2, ok := other.Get_().(D12_DC32)
			return ok && data1.Cf61 == data2.Cf61 && data1.Cf62.Cmp(data2.Cf62) == 0
		}
	case D12_DC33:
		{
			data2, ok := other.Get_().(D12_DC33)
			return ok && data1.Cf63 == data2.Cf63 && data1.Cf64.Equals(data2.Cf64)
		}
	case D12_DC34:
		{
			data2, ok := other.Get_().(D12_DC34)
			return ok && data1.Cf65 == data2.Cf65
		}
	case D12_DC31:
		{
			data2, ok := other.Get_().(D12_DC31)
			return ok && data1.Cf60.Equals(data2.Cf60)
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

type D13_DC36 struct {
	Cf67 _dafny.Int
}

func (D13_DC36) isD13() {}

func (CompanionStruct_D13_) Create_DC36_(Cf67 _dafny.Int) D13 {
	return D13{D13_DC36{Cf67}}
}

func (_this D13) Is_DC36() bool {
	_, ok := _this.Get_().(D13_DC36)
	return ok
}

type D13_DC37 struct {
	Cf68 _dafny.Map
	Cf69 _dafny.Set
	Cf70 _dafny.Array
}

func (D13_DC37) isD13() {}

func (CompanionStruct_D13_) Create_DC37_(Cf68 _dafny.Map, Cf69 _dafny.Set, Cf70 _dafny.Array) D13 {
	return D13{D13_DC37{Cf68, Cf69, Cf70}}
}

func (_this D13) Is_DC37() bool {
	_, ok := _this.Get_().(D13_DC37)
	return ok
}

type D13_DC35 struct {
	Cf66 _dafny.Set
}

func (D13_DC35) isD13() {}

func (CompanionStruct_D13_) Create_DC35_(Cf66 _dafny.Set) D13 {
	return D13{D13_DC35{Cf66}}
}

func (_this D13) Is_DC35() bool {
	_, ok := _this.Get_().(D13_DC35)
	return ok
}

type D13_DC38 struct {
	Cf71 D13
}

func (D13_DC38) isD13() {}

func (CompanionStruct_D13_) Create_DC38_(Cf71 D13) D13 {
	return D13{D13_DC38{Cf71}}
}

func (_this D13) Is_DC38() bool {
	_, ok := _this.Get_().(D13_DC38)
	return ok
}

func (CompanionStruct_D13_) Default() D13 {
	return Companion_D13_.Create_DC36_(_dafny.Zero)
}

func (_this D13) Dtor_cf67() _dafny.Int {
	return _this.Get_().(D13_DC36).Cf67
}

func (_this D13) Dtor_cf68() _dafny.Map {
	return _this.Get_().(D13_DC37).Cf68
}

func (_this D13) Dtor_cf69() _dafny.Set {
	return _this.Get_().(D13_DC37).Cf69
}

func (_this D13) Dtor_cf70() _dafny.Array {
	return _this.Get_().(D13_DC37).Cf70
}

func (_this D13) Dtor_cf66() _dafny.Set {
	return _this.Get_().(D13_DC35).Cf66
}

func (_this D13) Dtor_cf71() D13 {
	return _this.Get_().(D13_DC38).Cf71
}

func (_this D13) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D13_DC36:
		{
			return "D13.DC36" + "(" + _dafny.String(data.Cf67) + ")"
		}
	case D13_DC37:
		{
			return "D13.DC37" + "(" + _dafny.String(data.Cf68) + ", " + _dafny.String(data.Cf69) + ", " + _dafny.String(data.Cf70) + ")"
		}
	case D13_DC35:
		{
			return "D13.DC35" + "(" + _dafny.String(data.Cf66) + ")"
		}
	case D13_DC38:
		{
			return "D13.DC38" + "(" + _dafny.String(data.Cf71) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D13) Equals(other D13) bool {
	switch data1 := _this.Get_().(type) {
	case D13_DC36:
		{
			data2, ok := other.Get_().(D13_DC36)
			return ok && data1.Cf67.Cmp(data2.Cf67) == 0
		}
	case D13_DC37:
		{
			data2, ok := other.Get_().(D13_DC37)
			return ok && data1.Cf68.Equals(data2.Cf68) && data1.Cf69.Equals(data2.Cf69) && data1.Cf70 == data2.Cf70
		}
	case D13_DC35:
		{
			data2, ok := other.Get_().(D13_DC35)
			return ok && data1.Cf66.Equals(data2.Cf66)
		}
	case D13_DC38:
		{
			data2, ok := other.Get_().(D13_DC38)
			return ok && data1.Cf71.Equals(data2.Cf71)
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
	Cf73 _dafny.Int
	Cf74 bool
	Cf75 _dafny.Sequence
}

func (D14_DC40) isD14() {}

func (CompanionStruct_D14_) Create_DC40_(Cf73 _dafny.Int, Cf74 bool, Cf75 _dafny.Sequence) D14 {
	return D14{D14_DC40{Cf73, Cf74, Cf75}}
}

func (_this D14) Is_DC40() bool {
	_, ok := _this.Get_().(D14_DC40)
	return ok
}

type D14_DC39 struct {
	Cf72 _dafny.Array
}

func (D14_DC39) isD14() {}

func (CompanionStruct_D14_) Create_DC39_(Cf72 _dafny.Array) D14 {
	return D14{D14_DC39{Cf72}}
}

func (_this D14) Is_DC39() bool {
	_, ok := _this.Get_().(D14_DC39)
	return ok
}

type D14_DC41 struct {
	Cf76 D14
}

func (D14_DC41) isD14() {}

func (CompanionStruct_D14_) Create_DC41_(Cf76 D14) D14 {
	return D14{D14_DC41{Cf76}}
}

func (_this D14) Is_DC41() bool {
	_, ok := _this.Get_().(D14_DC41)
	return ok
}

func (CompanionStruct_D14_) Default() D14 {
	return Companion_D14_.Create_DC40_(_dafny.Zero, false, _dafny.EmptySeq)
}

func (_this D14) Dtor_cf73() _dafny.Int {
	return _this.Get_().(D14_DC40).Cf73
}

func (_this D14) Dtor_cf74() bool {
	return _this.Get_().(D14_DC40).Cf74
}

func (_this D14) Dtor_cf75() _dafny.Sequence {
	return _this.Get_().(D14_DC40).Cf75
}

func (_this D14) Dtor_cf72() _dafny.Array {
	return _this.Get_().(D14_DC39).Cf72
}

func (_this D14) Dtor_cf76() D14 {
	return _this.Get_().(D14_DC41).Cf76
}

func (_this D14) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D14_DC40:
		{
			return "D14.DC40" + "(" + _dafny.String(data.Cf73) + ", " + _dafny.String(data.Cf74) + ", " + data.Cf75.VerbatimString(true) + ")"
		}
	case D14_DC39:
		{
			return "D14.DC39" + "(" + _dafny.String(data.Cf72) + ")"
		}
	case D14_DC41:
		{
			return "D14.DC41" + "(" + _dafny.String(data.Cf76) + ")"
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
			return ok && data1.Cf73.Cmp(data2.Cf73) == 0 && data1.Cf74 == data2.Cf74 && data1.Cf75.Equals(data2.Cf75)
		}
	case D14_DC39:
		{
			data2, ok := other.Get_().(D14_DC39)
			return ok && data1.Cf72 == data2.Cf72
		}
	case D14_DC41:
		{
			data2, ok := other.Get_().(D14_DC41)
			return ok && data1.Cf76.Equals(data2.Cf76)
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

type D15_DC43 struct {
	Cf78 bool
}

func (D15_DC43) isD15() {}

func (CompanionStruct_D15_) Create_DC43_(Cf78 bool) D15 {
	return D15{D15_DC43{Cf78}}
}

func (_this D15) Is_DC43() bool {
	_, ok := _this.Get_().(D15_DC43)
	return ok
}

type D15_DC42 struct {
	Cf77 *C3
}

func (D15_DC42) isD15() {}

func (CompanionStruct_D15_) Create_DC42_(Cf77 *C3) D15 {
	return D15{D15_DC42{Cf77}}
}

func (_this D15) Is_DC42() bool {
	_, ok := _this.Get_().(D15_DC42)
	return ok
}

type D15_DC44 struct {
	Cf79 D15
}

func (D15_DC44) isD15() {}

func (CompanionStruct_D15_) Create_DC44_(Cf79 D15) D15 {
	return D15{D15_DC44{Cf79}}
}

func (_this D15) Is_DC44() bool {
	_, ok := _this.Get_().(D15_DC44)
	return ok
}

func (CompanionStruct_D15_) Default() D15 {
	return Companion_D15_.Create_DC43_(false)
}

func (_this D15) Dtor_cf78() bool {
	return _this.Get_().(D15_DC43).Cf78
}

func (_this D15) Dtor_cf77() *C3 {
	return _this.Get_().(D15_DC42).Cf77
}

func (_this D15) Dtor_cf79() D15 {
	return _this.Get_().(D15_DC44).Cf79
}

func (_this D15) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D15_DC43:
		{
			return "D15.DC43" + "(" + _dafny.String(data.Cf78) + ")"
		}
	case D15_DC42:
		{
			return "D15.DC42" + "(" + _dafny.String(data.Cf77) + ")"
		}
	case D15_DC44:
		{
			return "D15.DC44" + "(" + _dafny.String(data.Cf79) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D15) Equals(other D15) bool {
	switch data1 := _this.Get_().(type) {
	case D15_DC43:
		{
			data2, ok := other.Get_().(D15_DC43)
			return ok && data1.Cf78 == data2.Cf78
		}
	case D15_DC42:
		{
			data2, ok := other.Get_().(D15_DC42)
			return ok && data1.Cf77 == data2.Cf77
		}
	case D15_DC44:
		{
			data2, ok := other.Get_().(D15_DC44)
			return ok && data1.Cf79.Equals(data2.Cf79)
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

type D16_DC46 struct {
	Cf81 _dafny.Int
}

func (D16_DC46) isD16() {}

func (CompanionStruct_D16_) Create_DC46_(Cf81 _dafny.Int) D16 {
	return D16{D16_DC46{Cf81}}
}

func (_this D16) Is_DC46() bool {
	_, ok := _this.Get_().(D16_DC46)
	return ok
}

type D16_DC45 struct {
	Cf80 _dafny.Set
}

func (D16_DC45) isD16() {}

func (CompanionStruct_D16_) Create_DC45_(Cf80 _dafny.Set) D16 {
	return D16{D16_DC45{Cf80}}
}

func (_this D16) Is_DC45() bool {
	_, ok := _this.Get_().(D16_DC45)
	return ok
}

type D16_DC47 struct {
	Cf82 D16
}

func (D16_DC47) isD16() {}

func (CompanionStruct_D16_) Create_DC47_(Cf82 D16) D16 {
	return D16{D16_DC47{Cf82}}
}

func (_this D16) Is_DC47() bool {
	_, ok := _this.Get_().(D16_DC47)
	return ok
}

func (CompanionStruct_D16_) Default() D16 {
	return Companion_D16_.Create_DC46_(_dafny.Zero)
}

func (_this D16) Dtor_cf81() _dafny.Int {
	return _this.Get_().(D16_DC46).Cf81
}

func (_this D16) Dtor_cf80() _dafny.Set {
	return _this.Get_().(D16_DC45).Cf80
}

func (_this D16) Dtor_cf82() D16 {
	return _this.Get_().(D16_DC47).Cf82
}

func (_this D16) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D16_DC46:
		{
			return "D16.DC46" + "(" + _dafny.String(data.Cf81) + ")"
		}
	case D16_DC45:
		{
			return "D16.DC45" + "(" + _dafny.String(data.Cf80) + ")"
		}
	case D16_DC47:
		{
			return "D16.DC47" + "(" + _dafny.String(data.Cf82) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D16) Equals(other D16) bool {
	switch data1 := _this.Get_().(type) {
	case D16_DC46:
		{
			data2, ok := other.Get_().(D16_DC46)
			return ok && data1.Cf81.Cmp(data2.Cf81) == 0
		}
	case D16_DC45:
		{
			data2, ok := other.Get_().(D16_DC45)
			return ok && data1.Cf80.Equals(data2.Cf80)
		}
	case D16_DC47:
		{
			data2, ok := other.Get_().(D16_DC47)
			return ok && data1.Cf82.Equals(data2.Cf82)
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

// Definition of trait T0
type T0 interface {
	String() string
	F4() _dafny.CodePoint
	F4_set_(value _dafny.CodePoint)
	F5() _dafny.Sequence
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
	F4() _dafny.CodePoint
	F4_set_(value _dafny.CodePoint)
	F5() _dafny.Sequence
	F6() bool
	F7() bool
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
	F4() _dafny.CodePoint
	F4_set_(value _dafny.CodePoint)
	F6() bool
	F7() bool
	F5() _dafny.Sequence
	Fm1(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Int
	Fm2(globalState *GlobalState) _dafny.Sequence
	M1(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array)
	F8() _dafny.Int
	F9() bool
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

// Definition of trait T3
type T3 interface {
	String() string
	F4() _dafny.CodePoint
	F4_set_(value _dafny.CodePoint)
	Fm1(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Int
	Fm2(globalState *GlobalState) _dafny.Sequence
	M1(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array)
	F8() _dafny.Int
	F9() bool
	F6() bool
	F7() bool
	F5() _dafny.Sequence
	M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map
}
type CompanionStruct_T3_ struct {
	TraitID_ *_dafny.TraitID
}

var Companion_T3_ = CompanionStruct_T3_{
	TraitID_: &_dafny.TraitID{},
}

func (CompanionStruct_T3_) CastTo_(x interface{}) T3 {
	var t T3
	t, _ = x.(T3)
	return t
}

// End of trait T3

// Definition of class GlobalState
type GlobalState struct {
	F1  _dafny.Int
	_f0 bool
	_f2 bool
	_f3 _dafny.Array
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F1 = _dafny.Zero
	_this._f0 = false
	_this._f2 = false
	_this._f3 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *GlobalState) Ctor__(f0 bool, f1 _dafny.Int, f2 bool, f3 _dafny.Array) {
	{
		(_this)._f0 = f0
		(_this).F1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
	}
}
func (_this *GlobalState) F0() bool {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F2() bool {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() _dafny.Array {
	{
		return _this._f3
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f4 _dafny.CodePoint
	_f6 bool
	_f7 bool
	_f5 _dafny.Sequence
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f4 = _dafny.CodePoint('D')
	_this._f6 = false
	_this._f7 = false
	_this._f5 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C0{}
var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) F4() _dafny.CodePoint {
	return _this._f4
}
func (_this *C0) F4_set_(value _dafny.CodePoint) {
	_this._f4 = value
}
func (_this *C0) F6() bool {
	return _this._f6
}
func (_this *C0) F7() bool {
	return _this._f7
}
func (_this *C0) F5() _dafny.Sequence {
	return _this._f5
}
func (_this *C0) Ctor__(f6 bool, f7 bool, f4 _dafny.CodePoint, f5 _dafny.Sequence) {
	{
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C0) Fm9(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) _dafny.Int {
	{
		return (func() _dafny.Map {
			var _coll12 = _dafny.NewMapBuilder()
			_ = _coll12
			for _iter12 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(759))).Uint32(), func(coer27 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
				return func(arg27 _dafny.Int) interface{} {
					return coer27(arg27)
				}
			}(func(_326_i0 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf((Companion_D0_.Create_DC2_(_dafny.IntOfInt64(821), (_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer28 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg28 _dafny.Int) interface{} {
						return coer28(arg28)
					}
				}(func(_327_i1 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(318)
				})))).Dtor_cf5(), (_this).F6())
			}))).Elements()); ; {
				_compr_12, _ok12 := _iter12()
				if !_ok12 {
					break
				}
				var _328_v0 _dafny.MultiSet
				_328_v0 = interface{}(_compr_12).(_dafny.MultiSet)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(759))).Uint32(), func(coer29 func(_dafny.Int) _dafny.MultiSet) func(_dafny.Int) interface{} {
					return func(arg29 _dafny.Int) interface{} {
						return coer29(arg29)
					}
				}(func(_326_i0 _dafny.Int) _dafny.MultiSet {
					return _dafny.MultiSetOf((Companion_D0_.Create_DC2_(_dafny.IntOfInt64(821), (_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(755))).Uint32(), func(coer30 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg30 _dafny.Int) interface{} {
							return coer30(arg30)
						}
					}(func(_327_i1 _dafny.Int) _dafny.Int {
						return _dafny.IntOfInt64(318)
					})))).Dtor_cf5(), (_this).F6())
				})), _328_v0) {
					_coll12.Add(_328_v0, _dafny.SeqOf(Companion_D2_.Create_DC7_((_this).F7()), Companion_D2_.Create_DC7_((_this).F6())))
				}
			}
			return _coll12.ToMap()
		}()).Cardinality()
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f4 _dafny.CodePoint
	_f8 _dafny.Int
	_f9 bool
	_f6 bool
	_f7 bool
	_f5 _dafny.Sequence
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f4 = _dafny.CodePoint('D')
	_this._f8 = _dafny.Zero
	_this._f9 = false
	_this._f6 = false
	_this._f7 = false
	_this._f5 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C1{}
var _ T1 = &C1{}
var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) F4() _dafny.CodePoint {
	return _this._f4
}
func (_this *C1) F4_set_(value _dafny.CodePoint) {
	_this._f4 = value
}
func (_this *C1) F8() _dafny.Int {
	return _this._f8
}
func (_this *C1) F9() bool {
	return _this._f9
}
func (_this *C1) F6() bool {
	return _this._f6
}
func (_this *C1) F7() bool {
	return _this._f7
}
func (_this *C1) F5() _dafny.Sequence {
	return _this._f5
}
func (_this *C1) Ctor__(f8 _dafny.Int, f9 bool, f6 bool, f7 bool, f4 _dafny.CodePoint, f5 _dafny.Sequence) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C1) Fm1(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return _dafny.IntOfInt64(-485)
	}
}
func (_this *C1) Fm2(globalState *GlobalState) _dafny.Sequence {
	{
		var _source4 D3 = Companion_D3_.Create_DC12_(Companion_D3_.Create_DC12_(Companion_D3_.Create_DC11_((_this).F9(), (_this).F8(), _dafny.SetOf((_this).F9()))))
		_ = _source4
		if _source4.Is_DC11() {
			var _329___mcc_h0 bool = _source4.Get_().(D3_DC11).Cf20
			_ = _329___mcc_h0
			var _330___mcc_h1 _dafny.Int = _source4.Get_().(D3_DC11).Cf21
			_ = _330___mcc_h1
			var _331___mcc_h2 _dafny.Set = _source4.Get_().(D3_DC11).Cf22
			_ = _331___mcc_h2
			var _332_cf22 _dafny.Set = _331___mcc_h2
			_ = _332_cf22
			var _333_cf21 _dafny.Int = _330___mcc_h1
			_ = _333_cf21
			var _334_cf20 bool = _329___mcc_h0
			_ = _334_cf20
			return _dafny.SeqOf(_333_cf21)
		} else if _source4.Is_DC10() {
			var _335___mcc_h3 _dafny.Sequence = _source4.Get_().(D3_DC10).Cf19
			_ = _335___mcc_h3
			var _336_cf19 _dafny.Sequence = _335___mcc_h3
			_ = _336_cf19
			return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-920))).Uint32(), func(coer31 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg31 _dafny.Int) interface{} {
					return coer31(arg31)
				}
			}(func(_337_i0 _dafny.Int) _dafny.Int {
				return (_this).F8()
			})), _dafny.SeqOf(_dafny.IntOfInt64(537)))
		} else {
			var _338___mcc_h4 D3 = _source4.Get_().(D3_DC12).Cf23
			_ = _338___mcc_h4
			var _339_cf23 D3 = _338___mcc_h4
			_ = _339_cf23
			return _dafny.SeqOf((func() _dafny.Set {
				var _coll13 = _dafny.NewBuilder()
				_ = _coll13
				for _iter13 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(372), _dafny.IntOfInt64(170))); ; {
					_compr_13, _ok13 := _iter13()
					if !_ok13 {
						break
					}
					var _340_v0 _dafny.Int
					_340_v0 = interface{}(_compr_13).(_dafny.Int)
					if ((_dafny.IntOfInt64(372)).Cmp(_340_v0) <= 0) && ((_340_v0).Cmp(_dafny.IntOfInt64(170)) < 0) {
						_coll13.Add((_340_v0).Plus((_this).F8()))
					}
				}
				return _coll13.ToSet()
			}()).Cardinality(), (_dafny.MultiSetOf(_this.F4())).Cardinality())
		}
	}
}
func (_this *C1) Fm19(p0 bool, p1 bool, p2 _dafny.Set, globalState *GlobalState) bool {
	{
		return (_this).F7()
	}
}
func (_this *C1) Fm20(p0 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(639))).Uint32(), func(coer32 func(_dafny.Int) D1) func(_dafny.Int) interface{} {
			return func(arg32 _dafny.Int) interface{} {
				return coer32(arg32)
			}
		}(func(_341_i0 _dafny.Int) D1 {
			return Companion_D1_.Create_DC5_((_this).F8())
		})), _dafny.SeqOf(Companion_D1_.Create_DC5_((_this).F8()), Companion_D1_.Create_DC5_((_this).F8()), Companion_D1_.Create_DC5_((_this).F8())))
	}
}
func (_this *C1) M1(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _342_v0 D2
		_ = _342_v0
		_342_v0 = Companion_D2_.Create_DC7_(p1)
		var _343_v1 _dafny.Sequence
		_ = _343_v1
		_343_v1 = _dafny.SeqOf((_342_v0).Dtor_cf15())
		var _344_v2 _dafny.Set
		_ = _344_v2
		_344_v2 = _dafny.SetOf((_this).F7(), false)
		var _345_v3 _dafny.Sequence
		_ = _345_v3
		_345_v3 = _dafny.SeqOf((_dafny.Zero).Minus((_344_v2).Cardinality()))
		var _346_v4 _dafny.Sequence
		_ = _346_v4
		_346_v4 = _dafny.SeqOf(_dafny.IntOfUint32((_345_v3).Cardinality()))
		var _347_v5 D0
		_ = _347_v5
		_347_v5 = Companion_D0_.Create_DC2_((_this).F8(), (_343_v1).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_343_v1).Cardinality()))).Uint32()).(bool), _345_v3)
		var _348_v6 _dafny.Map
		_ = _348_v6
		_348_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_347_v5).Dtor_cf5())
		var _349_v7 _dafny.Set
		_ = _349_v7
		_349_v7 = _dafny.SetOf(func(_pat_let6_0 D2) D2 {
			return func(_350_dt__update__tmp_h0 D2) D2 {
				return func(_pat_let7_0 bool) D2 {
					return func(_351_dt__update_hcf15_h0 bool) D2 {
						return Companion_D2_.Create_DC7_(_351_dt__update_hcf15_h0)
					}(_pat_let7_0)
				}(true)
			}(_pat_let6_0)
		}(Companion_D2_.Create_DC7_(p0)), _342_v0, _342_v0)
		var _352_v8 D3
		_ = _352_v8
		_352_v8 = Companion_D3_.Create_DC12_(Companion_D3_.Create_DC10_((_this).F5()))
		var _353_v9 D3
		_ = _353_v9
		_353_v9 = Companion_D3_.Create_DC12_(_352_v8)
		var _354_v10 D3
		_ = _354_v10
		_354_v10 = Companion_D3_.Create_DC12_(_352_v8)
		var _355_v11 _dafny.MultiSet
		_ = _355_v11
		_355_v11 = _dafny.MultiSetOf(_354_v10)
		var _356_v12 _dafny.Array
		_ = _356_v12
		var _nwElement0_7 _dafny.Int = (_this).Fm1((_this).F6(), _348_v6, (_this).F6(), globalState)
		_ = _nwElement0_7
		var _nw54 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(9))
		_ = _nw54
		(_nw54).ArraySet1(_nwElement0_7, 0)
		(_nw54).ArraySet1((_this).F8(), 1)
		(_nw54).ArraySet1((func() _dafny.Int {
			if (_this).F6() {
				return (_this).F8()
			}
			return _dafny.IntOfInt64(31)
		})(), 2)
		(_nw54).ArraySet1((_this).F8(), 3)
		(_nw54).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_343_v1, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(531), _dafny.IntOfUint32((_343_v1).Cardinality()))).Uint32(), (_this).F7())).Cardinality()), 4)
		(_nw54).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if (_this).Fm19(false, p0, _349_v7, globalState) {
				return _dafny.IntOfInt64(593)
			}
			return (_this).F8()
		})()), 5)
		(_nw54).ArraySet1((_this).F8(), 6)
		(_nw54).ArraySet1(((_355_v11).Update(_354_v10, Companion_Default___.Abs((_this).F8()))).Cardinality(), 7)
		(_nw54).ArraySet1((_this).F8(), 8)
		_356_v12 = _nw54
		var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_356_v12), 0))
		_ = _index30
		(_356_v12).ArraySet1(((_this).F8()).Times((_this).F8()), (_index30).Int())
		var _357_v13 _dafny.Array
		_ = _357_v13
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(16))
		_ = _nw55
		_357_v13 = _nw55
		var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_357_v13), 0))
		_ = _index31
		(_357_v13).ArraySet1((func() bool {
			if p0 {
				return (_this).F7()
			}
			return p3
		})(), (_index31).Int())
		var _pat_let_tv6 = p3
		_ = _pat_let_tv6
		var _pat_let_tv7 = _346_v4
		_ = _pat_let_tv7
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_356_v12), 0))
		_ = _index32
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_357_v13), 0))
		_ = _index33
		var _rhs28 _dafny.Int = func(_source5 D0) _dafny.Int {
			if _source5.Is_DC0() {
				return (_this).F8()
			} else if _source5.Is_DC1() {
				var _358___mcc_h0 bool = _source5.Get_().(D0_DC1).Cf0
				_ = _358___mcc_h0
				var _359___mcc_h1 _dafny.Int = _source5.Get_().(D0_DC1).Cf1
				_ = _359___mcc_h1
				var _360___mcc_h2 bool = _source5.Get_().(D0_DC1).Cf2
				_ = _360___mcc_h2
				var _361___mcc_h3 _dafny.Int = _source5.Get_().(D0_DC1).Cf3
				_ = _361___mcc_h3
				var _362_cf3 _dafny.Int = _361___mcc_h3
				_ = _362_cf3
				var _363_cf2 bool = _360___mcc_h2
				_ = _363_cf2
				var _364_cf1 _dafny.Int = _359___mcc_h1
				_ = _364_cf1
				var _365_cf0 bool = _358___mcc_h0
				_ = _365_cf0
				return _dafny.IntOfUint32(((_this).F5()).Cardinality())
			} else {
				var _366___mcc_h4 _dafny.Int = _source5.Get_().(D0_DC2).Cf4
				_ = _366___mcc_h4
				var _367___mcc_h5 bool = _source5.Get_().(D0_DC2).Cf5
				_ = _367___mcc_h5
				var _368___mcc_h6 _dafny.Sequence = _source5.Get_().(D0_DC2).Cf6
				_ = _368___mcc_h6
				var _369_cf6 _dafny.Sequence = _368___mcc_h6
				_ = _369_cf6
				var _370_cf5 bool = _367___mcc_h5
				_ = _370_cf5
				var _371_cf4 _dafny.Int = _366___mcc_h4
				_ = _371_cf4
				return (_this).F8()
			}
		}(func(_pat_let8_0 D0) D0 {
			return func(_372_dt__update__tmp_h1 D0) D0 {
				return func(_pat_let9_0 bool) D0 {
					return func(_373_dt__update_hcf5_h0 bool) D0 {
						return func(_pat_let10_0 _dafny.Sequence) D0 {
							return func(_374_dt__update_hcf6_h0 _dafny.Sequence) D0 {
								return Companion_D0_.Create_DC2_((_372_dt__update__tmp_h1).Dtor_cf4(), _373_dt__update_hcf5_h0, _374_dt__update_hcf6_h0)
							}(_pat_let10_0)
						}(_pat_let_tv7)
					}(_pat_let9_0)
				}(_pat_let_tv6)
			}(_pat_let8_0)
		}(Companion_Default___.Fm21((_this).F6(), true, (_this).F9(), p0, globalState)))
		_ = _rhs28
		var _rhs29 bool = true
		_ = _rhs29
		var _rhs30 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F8(), (_this).F8())
		_ = _rhs30
		var _lhs12 _dafny.Array = _356_v12
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_356_v12), 0))
		_ = _lhs13
		var _lhs14 _dafny.Array = _357_v13
		_ = _lhs14
		var _lhs15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_357_v13), 0))
		_ = _lhs15
		(_lhs12).ArraySet1(_rhs28, (_lhs13).Int())
		(_lhs14).ArraySet1(_rhs29, (_lhs15).Int())
		r1 = _rhs30
		var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_357_v13), 0))
		_ = _index34
		(_357_v13).ArraySet1(false, (_index34).Int())
		var _375_v14 *C0
		_ = _375_v14
		var _nw56 *C0 = New_C0_()
		_ = _nw56
		_nw56.Ctor__(!(p3), p0, _this.F4(), (func() _dafny.Sequence {
			if (_357_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(378), _dafny.ArrayLen((_357_v13), 0))).Int()).(bool) {
				return (_this).F5()
			}
			return (_this).F5()
		})())
		_375_v14 = _nw56
		var _376_v15 _dafny.Set
		_ = _376_v15
		_376_v15 = _dafny.SetOf((_this).F5(), (_this).F5())
		var _377_v16 _dafny.Set
		_ = _377_v16
		_377_v16 = _dafny.SetOf((_356_v12).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_356_v12), 0))).Int()).(_dafny.Int), (_this).F8())
		var _378_v17 _dafny.Map
		_ = _378_v17
		_378_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf((_this).F5(), _dafny.Companion_Sequence_.Update((_this).F5(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(668))).Uint32(), func(coer33 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg33 _dafny.Int) interface{} {
				return coer33(arg33)
			}
		}(func(_379_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		})), (Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F5()).Cardinality()), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(668))).Uint32(), func(coer34 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg34 _dafny.Int) interface{} {
				return coer34(arg34)
			}
		}(func(_380_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('f')
		}))).Cardinality()))).Uint32(), _this.F4())).Cardinality()), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32(), _this.F4()), _dafny.SeqOf(_this.F4()), (_this).F5(), (_this).F5())).Intersection(_376_v15), (_377_v16).IsDisjointFrom(_dafny.SetOf((_this).F8(), (_this).F8())))
		_378_v17 = (_378_v17).Update((_376_v15).Difference(_dafny.SetOf((_this).F5())), (_this).F7())
		var _381_v18 _dafny.Set
		_ = _381_v18
		_381_v18 = _dafny.SetOf(_357_v13, _357_v13)
		var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.ArrayLen((_356_v12), 0))
		_ = _index35
		(_356_v12).ArraySet1((_381_v18).Cardinality(), (_index35).Int())
		var _382_v19 *C0
		_ = _382_v19
		var _nw57 *C0 = New_C0_()
		_ = _nw57
		_nw57.Ctor__(p0, (_this).F7(), _this.F4(), (_this).F5())
		_382_v19 = _nw57
		r0 = (_this).F8()
		r1 = (_dafny.Zero).Minus(_dafny.IntOfUint32((_345_v3).Cardinality()))
		var _383_v20 _dafny.Array
		_ = _383_v20
		var _nw58 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(14))
		_ = _nw58
		_383_v20 = _nw58
		r2 = _383_v20
		return r0, r1, r2
	}
}
func (_this *C1) M8(p0 bool, p1 bool, globalState *GlobalState) (_dafny.Set, _dafny.Array, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Set = _dafny.EmptySet
		_ = r0
		var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var r3 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r3
		(globalState).F1 = (_this).F8()
		var _384_v0 _dafny.Array
		_ = _384_v0
		var _nw59 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(7))
		_ = _nw59
		_384_v0 = _nw59
		for _iter14 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_384_v0), 0))); ; {
			_guard_loop_0, _ok14 := _iter14()
			if !_ok14 {
				break
			}
			var _385_i0 _dafny.Int
			_385_i0 = interface{}(_guard_loop_0).(_dafny.Int)
			if (true) && (((_385_i0).Sign() != -1) && ((_385_i0).Cmp(_dafny.ArrayLen((_384_v0), 0)) < 0)) {
				(_384_v0).ArraySet1((_dafny.MultiSetOf((_this).F8(), (_this).F8(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32(((_this).F5()).Cardinality()))), _dafny.IntOfInt64(-67))).Difference(_dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(p1, (_this).F9(), (_this).F6())).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(173))))), (_385_i0).Int())
			}
		}
		var _386_i1 _dafny.Int
		_ = _386_i1
		_386_i1 = _dafny.Zero
		{
			for false {
				{
					if (_386_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L1
					}
					_386_i1 = (_386_i1).Plus(_dafny.One)
					var _387_v1 D1
					_ = _387_v1
					_387_v1 = Companion_D1_.Create_DC4_(_this.F4(), (_this).F8(), _this.F4(), (_this).F8(), true)
					var _388_v2 _dafny.Sequence
					_ = _388_v2
					_388_v2 = _dafny.SeqOf(_this.F4(), (_387_v1).Dtor_cf10())
					_388_v2 = _dafny.Companion_Sequence_.Update((_this).F5(), (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F8()), (_this).F8()), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32(), _this.F4())
					var _389_v3 bool
					_ = _389_v3
					_389_v3 = false
					_389_v3 = (_this).F6()
					(globalState).F1 = _dafny.IntOfInt64(76)
					var _390_v4 D3
					_ = _390_v4
					_390_v4 = Companion_D3_.Create_DC10_(_388_v2)
					var _391_v5 D3
					_ = _391_v5
					_391_v5 = Companion_D3_.Create_DC12_(_390_v4)
					_391_v5 = Companion_D3_.Create_DC12_(_390_v4)
					goto C1
				}
			C1:
			}
			goto L1
		}
	L1:
		Companion_Default___.M0(globalState)
		var _392_v6 bool
		_ = _392_v6
		_392_v6 = false
		_392_v6 = (_this).F7()
		var _393_v7 _dafny.Sequence
		_ = _393_v7
		_393_v7 = _dafny.SeqOf((_this).F8())
		var _394_v8 _dafny.Map
		_ = _394_v8
		_394_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).Fm2(globalState))
		var _395_v9 _dafny.Set
		_ = _395_v9
		_395_v9 = _dafny.SetOf(p0, (_this).F7(), (_this).F9(), _392_v6)
		var _396_v10 _dafny.Sequence
		_ = _396_v10
		_396_v10 = _dafny.SeqOf(_dafny.SeqOf((_395_v9).Cardinality()), _393_v7, (_this).Fm2(globalState))
		var _397_v11 _dafny.Array
		_ = _397_v11
		var _nwElement0_8 _dafny.Sequence = _dafny.SeqOf((_this).F8())
		_ = _nwElement0_8
		var _nw60 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(23))
		_ = _nw60
		(_nw60).ArraySet1(_nwElement0_8, 0)
		(_nw60).ArraySet1(_393_v7, 1)
		(_nw60).ArraySet1(_393_v7, 2)
		(_nw60).ArraySet1(_dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_393_v7, _393_v7), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_393_v7, _393_v7)).Cardinality()))).Uint32(), (_this).F8()), 3)
		(_nw60).ArraySet1(_393_v7, 4)
		(_nw60).ArraySet1(_393_v7, 5)
		(_nw60).ArraySet1(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if false {
				return (func() _dafny.Sequence {
					if (_394_v8).Contains(p1) {
						return (_394_v8).Get(p1).(_dafny.Sequence)
					}
					return (_396_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F5()).Cardinality()), _dafny.IntOfUint32((_396_v10).Cardinality()))).Uint32()).(_dafny.Sequence)
				})()
			}
			return _393_v7
		})(), (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_this).F8()), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if false {
				return (func() _dafny.Sequence {
					if (_394_v8).Contains(p1) {
						return (_394_v8).Get(p1).(_dafny.Sequence)
					}
					return (_396_v10).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_this).F5()).Cardinality()), _dafny.IntOfUint32((_396_v10).Cardinality()))).Uint32()).(_dafny.Sequence)
				})()
			}
			return _393_v7
		})()).Cardinality()))).Uint32(), (_this).F8()), 6)
		(_nw60).ArraySet1(_dafny.SeqOf((_this).F8(), (_this).F8(), (_this).F8(), (_this).F8(), (_this).F8()), 7)
		(_nw60).ArraySet1(_393_v7, 8)
		(_nw60).ArraySet1(_393_v7, 9)
		(_nw60).ArraySet1(_393_v7, 10)
		(_nw60).ArraySet1(_393_v7, 11)
		(_nw60).ArraySet1(_dafny.SeqOf((_this).F8(), (_this).F8(), (_this).F8(), (_this).F8()), 12)
		(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_393_v7, _393_v7), 13)
		(_nw60).ArraySet1(_393_v7, 14)
		(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_393_v7, _393_v7), 15)
		(_nw60).ArraySet1(_393_v7, 16)
		(_nw60).ArraySet1(_393_v7, 17)
		(_nw60).ArraySet1((_this).Fm2(globalState), 18)
		(_nw60).ArraySet1(_393_v7, 19)
		(_nw60).ArraySet1(_393_v7, 20)
		(_nw60).ArraySet1(_393_v7, 21)
		(_nw60).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_393_v7, _393_v7), 22)
		_397_v11 = _nw60
		var _398_v12 _dafny.Sequence
		_ = _398_v12
		_398_v12 = _dafny.SeqOf(_392_v6)
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_397_v11), 0))
		_ = _index36
		(_397_v11).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).Fm2(globalState), _dafny.Companion_Sequence_.Update(_393_v7, (Companion_Default___.SafeIndex((_393_v7).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_393_v7).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_393_v7).Cardinality()))).Uint32(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_398_v12, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_393_v7).Cardinality()), _dafny.IntOfUint32((_398_v12).Cardinality()))).Uint32(), p1)).Cardinality()))), (_index36).Int())
		var _399_v13 D0
		_ = _399_v13
		_399_v13 = Companion_D0_.Create_DC2_(_dafny.IntOfInt64(981), (_this).F7(), _393_v7)
		var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(491), _dafny.ArrayLen((_397_v11), 0))
		_ = _index37
		(_397_v11).ArraySet1((_399_v13).Dtor_cf6(), (_index37).Int())
		var _400_v14 _dafny.Map
		_ = _400_v14
		_400_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_392_v6, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(697))).Uint32(), func(coer35 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg35 _dafny.Int) interface{} {
				return coer35(arg35)
			}
		}(func(_401_i2 _dafny.Int) _dafny.CodePoint {
			return _this.F4()
		})))
		var _402_v15 _dafny.Map
		_ = _402_v15
		_402_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
		var _403_v16 _dafny.Set
		_ = _403_v16
		_403_v16 = _dafny.SetOf(_402_v15)
		var _404_v17 _dafny.Sequence
		_ = _404_v17
		_404_v17 = _dafny.SeqOf(_403_v16)
		var _405_v18 _dafny.Sequence
		_ = _405_v18
		_405_v18 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7()), Companion_Default___.Fm23((_this).F8(), (_this).F8(), _this.F4(), (_this).F6(), globalState))
		r0 = (func() _dafny.Set {
			if (_400_v14).Contains(_392_v6) {
				return (func() _dafny.Set {
					if (_this).F6() {
						return (_404_v17).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_404_v17).Cardinality()))).Uint32()).(_dafny.Set)
					}
					return Companion_Default___.Fm22(globalState)
				})()
			}
			return (_403_v16).Intersection(func() _dafny.Set {
				var _coll14 = _dafny.NewBuilder()
				_ = _coll14
				for _iter15 := _dafny.Iterate((_405_v18).Elements()); ; {
					_compr_14, _ok15 := _iter15()
					if !_ok15 {
						break
					}
					var _406_v19 _dafny.Map
					_406_v19 = interface{}(_compr_14).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_405_v18, _406_v19) {
						_coll14.Add(_406_v19)
					}
				}
				return _coll14.ToSet()
			}())
		})()
		var _407_v20 _dafny.Array
		_ = _407_v20
		var _len0_8 _dafny.Int = _dafny.IntOfInt64(2)
		_ = _len0_8
		var _nw61 _dafny.Array
		_ = _nw61
		if _len0_8.Cmp(_dafny.Zero) == 0 {
			_nw61 = _dafny.NewArray(_len0_8)
		} else {
			var _init8 func(_dafny.Int) bool = func(_408_i3 _dafny.Int) bool {
				return (_this).F7()
			}
			_ = _init8
			var _element0_8 = _init8(_dafny.Zero)
			_ = _element0_8
			_nw61 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
			(_nw61).ArraySet1(_element0_8, 0)
			var _nativeLen0_8 = (_len0_8).Int()
			_ = _nativeLen0_8
			for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
				(_nw61).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
			}
		}
		_407_v20 = _nw61
		r1 = _407_v20
		r2 = (_dafny.Zero).Minus(((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqOf(_this.F4(), _this.F4(), _this.F4())).Cardinality()))).Minus((_this).F8()))
		r3 = _407_v20
		return r0, r1, r2, r3
	}
}

// End of class C1

// Definition of class C2
type C2 struct {
	F15  _dafny.Array
	_f14 _dafny.Array
}

func New_C2_() *C2 {
	_this := C2{}

	_this.F15 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f14 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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

func (_this *C2) Ctor__(f14 _dafny.Array, f15 _dafny.Array) {
	{
		(_this)._f14 = f14
		(_this).F15 = f15
	}
}
func (_this *C2) Fm6(globalState *GlobalState) _dafny.Sequence {
	{
		if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(true), false)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(false), true)) {
			return _dafny.SeqOf(_dafny.SetOf(true))
		} else {
			return _dafny.SeqOf(_dafny.SetOf(false), _dafny.SetOf(false, true, true))
		}
	}
}
func (_this *C2) M6(globalState *GlobalState) (_dafny.Int, D1, _dafny.Array, D1) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 D1 = Companion_D1_.Default()
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var r3 D1 = Companion_D1_.Default()
		_ = r3
		var _409_v0 _dafny.Sequence
		_ = _409_v0
		_409_v0 = _dafny.UnicodeSeqOfUtf8Bytes("svbxxdv")
		var _410_i0 _dafny.Int
		_ = _410_i0
		_410_i0 = _dafny.Zero
		{
			for _dafny.Companion_Sequence_.Equal(_dafny.UnicodeSeqOfUtf8Bytes("hfrnqnere"), _409_v0) {
				{
					if (_410_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L2
					}
					_410_i0 = (_410_i0).Plus(_dafny.One)
					var _411_v1 bool
					_ = _411_v1
					_411_v1 = true
					var _412_v2 _dafny.Int
					_ = _412_v2
					_412_v2 = _dafny.IntOfInt64(-884)
					var _413_v3 _dafny.Array
					_ = _413_v3
					var _len0_9 _dafny.Int = _dafny.IntOfInt64(2)
					_ = _len0_9
					var _nw62 _dafny.Array
					_ = _nw62
					if _len0_9.Cmp(_dafny.Zero) == 0 {
						_nw62 = _dafny.NewArray(_len0_9)
					} else {
						var _init9 func(_dafny.Int) _dafny.Sequence = (func(_414_v1 bool) func(_dafny.Int) _dafny.Sequence {
							return func(_415_i1 _dafny.Int) _dafny.Sequence {
								return _dafny.SeqOf(!(_414_v1))
							}
						})(_411_v1)
						_ = _init9
						var _element0_9 = _init9(_dafny.Zero)
						_ = _element0_9
						_nw62 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
						(_nw62).ArraySet1(_element0_9, 0)
						var _nativeLen0_9 = (_len0_9).Int()
						_ = _nativeLen0_9
						for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
							(_nw62).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
						}
					}
					_413_v3 = _nw62
					var _416_v4 _dafny.Map
					_ = _416_v4
					_416_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_412_v2, Companion_D1_.Create_DC3_(_413_v3))
					var _417_v5 _dafny.Map
					_ = _417_v5
					_417_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v1, _416_v4)
					var _418_v6 D2
					_ = _418_v6
					_418_v6 = Companion_D2_.Create_DC6_(_416_v4)
					_417_v5 = (_417_v5).Update((_411_v1) == (_411_v1), (_416_v4).Merge((_418_v6).Dtor_cf14()))
					var _419_v7 _dafny.CodePoint
					_ = _419_v7
					_419_v7 = _dafny.CodePoint('y')
					var _420_v8 _dafny.Map
					_ = _420_v8
					_420_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_411_v1, _419_v7)
					var _421_v9 D1
					_ = _421_v9
					_421_v9 = Companion_D1_.Create_DC3_(_413_v3)
					var _422_v10 _dafny.MultiSet
					_ = _422_v10
					_422_v10 = _dafny.MultiSetOf(_421_v9)
					var _423_v11 _dafny.Map
					_ = _423_v11
					_423_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_419_v7, (_422_v10).Cardinality(), _411_v1, globalState), _411_v1)
					var _424_v12 _dafny.Map
					_ = _424_v12
					_424_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_420_v8, _423_v11)
					_424_v12 = _424_v12
					_409_v0 = _dafny.UnicodeSeqOfUtf8Bytes("d")
					var _425_v13 _dafny.Sequence
					_ = _425_v13
					_425_v13 = _dafny.SeqOf(!(_411_v1))
					var _426_v14 _dafny.Set
					_ = _426_v14
					_426_v14 = _dafny.SetOf(_411_v1, _411_v1)
					var _427_v15 D0
					_ = _427_v15
					_427_v15 = Companion_D0_.Create_DC2_(_412_v2, _411_v1, _dafny.SeqOf((_426_v14).Cardinality(), _dafny.IntOfUint32((_409_v0).Cardinality())))
					_409_v0 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("aetllib"), _dafny.Companion_Sequence_.Update(_409_v0, (Companion_Default___.SafeIndex(Companion_Default___.Fm8(_dafny.IntOfUint32((_dafny.SeqOf(_419_v7)).Cardinality()), _411_v1, _411_v1, globalState), _dafny.IntOfUint32((_409_v0).Cardinality()))).Uint32(), _dafny.CodePoint('o'))), _dafny.Companion_Sequence_.Update(_409_v0, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_425_v13, (Companion_Default___.SafeIndex(_412_v2, _dafny.IntOfUint32((_425_v13).Cardinality()))).Uint32(), (_427_v15).Dtor_cf5())).Cardinality()), _dafny.IntOfUint32((_409_v0).Cardinality()))).Uint32(), _419_v7))
					goto C2
				}
			C2:
			}
			goto L2
		}
	L2:
		if true {
			var _428_v16 bool
			_ = _428_v16
			_428_v16 = false
			var _429_v17 _dafny.CodePoint
			_ = _429_v17
			_429_v17 = _dafny.CodePoint('j')
			var _430_v18 *C0
			_ = _430_v18
			var _nw63 *C0 = New_C0_()
			_ = _nw63
			_nw63.Ctor__(_428_v16, _428_v16, _429_v17, _409_v0)
			_430_v18 = _nw63
			var _431_v19 _dafny.Map
			_ = _431_v19
			_431_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("gmxt"), _428_v16)
			_431_v19 = (_431_v19).Merge((_431_v19).Update(_409_v0, _428_v16))
			var _432_v20 _dafny.Int
			_ = _432_v20
			_432_v20 = _dafny.IntOfInt64(475)
			var _433_v21 _dafny.Set
			_ = _433_v21
			_433_v21 = _dafny.SetOf(_432_v20, _432_v20)
			(globalState).F1 = ((func() _dafny.Set {
				if (_433_v21).IsDisjointFrom(_433_v21) {
					return _433_v21
				}
				return Companion_Default___.Fm10(true, _432_v20, Companion_Default___.Fm8(_432_v20, _428_v16, _428_v16, globalState), globalState)
			})()).Cardinality()
			var _434_v22 D1
			_ = _434_v22
			_434_v22 = Companion_D1_.Create_DC5_(_432_v20)
			_434_v22 = (func() D1 {
				if _428_v16 {
					return _434_v22
				}
				return _434_v22
			})()
			var _435_v23 _dafny.Sequence
			_ = _435_v23
			_435_v23 = _dafny.SeqOf(_432_v20, (_dafny.Zero).Minus(_432_v20))
			var _436_v24 _dafny.Map
			_ = _436_v24
			_436_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_409_v0, _432_v20)
			var _437_v25 _dafny.Array
			_ = _437_v25
			var _nwElement0_9 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_435_v23, _435_v23)
			_ = _nwElement0_9
			var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(4))
			_ = _nw64
			(_nw64).ArraySet1(_nwElement0_9, 0)
			(_nw64).ArraySet1(_435_v23, 1)
			(_nw64).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_435_v23, _dafny.SeqOf(_432_v20, (_436_v24).Cardinality(), _432_v20, (_dafny.MultiSetOf(_428_v16, _428_v16)).Cardinality(), _dafny.IntOfInt64(782))), 2)
			(_nw64).ArraySet1(_dafny.Companion_Sequence_.Update(_435_v23, (Companion_Default___.SafeIndex(_432_v20, _dafny.IntOfUint32((_435_v23).Cardinality()))).Uint32(), _432_v20), 3)
			_437_v25 = _nw64
			_437_v25 = _437_v25
		} else {
			var _438_v26 _dafny.Int
			_ = _438_v26
			_438_v26 = _dafny.IntOfInt64(650)
			(globalState).F1 = (_438_v26).Minus(_438_v26)
			var _439_v27 bool
			_ = _439_v27
			_439_v27 = false
			_439_v27 = _439_v27
			var _440_v28 _dafny.Sequence
			_ = _440_v28
			_440_v28 = _dafny.SeqOf(_438_v26, _438_v26)
			var _441_v29 D0
			_ = _441_v29
			_441_v29 = Companion_D0_.Create_DC2_(_438_v26, false, _440_v28)
			if !((_441_v29).Dtor_cf5()) {
				var _442_v30 _dafny.Array
				_ = _442_v30
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw65
				_442_v30 = _nw65
				var _443_v31 _dafny.Sequence
				_ = _443_v31
				_443_v31 = _dafny.SeqOf(_439_v27, _439_v27, _439_v27)
				var _444_v32 _dafny.MultiSet
				_ = _444_v32
				_444_v32 = _dafny.MultiSetOf(_dafny.SeqOf(_439_v27), _443_v31)
				var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_442_v30), 0))
				_ = _index38
				(_442_v30).ArraySet1(((_444_v32).Cardinality()).Times(_438_v26), (_index38).Int())
				var _445_v33 _dafny.Array
				_ = _445_v33
				var _nw66 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(26))
				_ = _nw66
				_445_v33 = _nw66
				var _446_v34 D1
				_ = _446_v34
				_446_v34 = Companion_D1_.Create_DC3_(_445_v33)
				var _447_v35 _dafny.Map
				_ = _447_v35
				_447_v35 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_438_v26, _446_v34)
				var _448_v36 D2
				_ = _448_v36
				_448_v36 = Companion_D2_.Create_DC6_(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_438_v26, _446_v34))
				var _449_v37 _dafny.MultiSet
				_ = _449_v37
				_449_v37 = _dafny.MultiSetOf((func() D2 {
					if _439_v27 {
						return Companion_D2_.Create_DC6_(_447_v35)
					}
					return _448_v36
				})(), _448_v36)
				var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_442_v30), 0))
				_ = _index39
				(_442_v30).ArraySet1((_449_v37).Cardinality(), (_index39).Int())
				var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index40
				((_this).F14()).ArraySet1(true, (_index40).Int())
				var _450_v38 _dafny.MultiSet
				_ = _450_v38
				_450_v38 = _dafny.MultiSetOf(_439_v27, (_441_v29).Dtor_cf5(), _439_v27)
				var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index41
				((_this).F14()).ArraySet1((_dafny.IntOfInt64(181)).Cmp((func() _dafny.Int {
					if (_450_v38).Contains(_439_v27) {
						return (_450_v38).Multiplicity(_439_v27)
					}
					return (_442_v30).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_442_v30), 0))).Int()).(_dafny.Int)
				})()) > 0, (_index41).Int())
				var _451_v39 _dafny.CodePoint
				_ = _451_v39
				_451_v39 = _dafny.CodePoint('h')
				var _452_v40 *C0
				_ = _452_v40
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__(((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), (func() bool {
					if _439_v27 {
						return _439_v27
					}
					return ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				})(), _451_v39, _409_v0)
				_452_v40 = _nw67
				_452_v40 = _452_v40
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_442_v30), 0))
				_ = _index42
				var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index43
				var _rhs31 bool = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				_ = _rhs31
				var _rhs32 _dafny.Int = _438_v26
				_ = _rhs32
				var _rhs33 bool = ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool)
				_ = _rhs33
				var _rhs34 _dafny.CodePoint = _451_v39
				_ = _rhs34
				var _lhs16 _dafny.Array = _442_v30
				_ = _lhs16
				var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(333), _dafny.ArrayLen((_442_v30), 0))
				_ = _lhs17
				var _lhs18 _dafny.Array = (_this).F14()
				_ = _lhs18
				var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(310), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _lhs19
				_439_v27 = _rhs31
				(_lhs16).ArraySet1(_rhs32, (_lhs17).Int())
				(_lhs18).ArraySet1(_rhs33, (_lhs19).Int())
				_451_v39 = _rhs34
				r0 = _438_v26
			} else {
				var _453_v41 _dafny.Set
				_ = _453_v41
				_453_v41 = _dafny.SetOf(_439_v27)
				var _454_v42 D2
				_ = _454_v42
				_454_v42 = Companion_D2_.Create_DC8_((_438_v26).Plus(_438_v26), _438_v26)
				var _455_v43 _dafny.Sequence
				_ = _455_v43
				_455_v43 = _dafny.SeqOf(_439_v27, _439_v27, _439_v27, false)
				var _456_v44 _dafny.MultiSet
				_ = _456_v44
				_456_v44 = _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
					if !(_439_v27) {
						return _455_v43
					}
					return _455_v43
				})(), (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ibbeoyt")).Cardinality()), _dafny.IntOfUint32(((func() _dafny.Sequence {
					if !(_439_v27) {
						return _455_v43
					}
					return _455_v43
				})()).Cardinality()))).Uint32(), _439_v27)).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("vjhcwwr")).Cardinality()), _438_v26)
				var _rhs35 _dafny.Int = (func() _dafny.Int {
					if (_456_v44).Contains(_438_v26) {
						return (_456_v44).Multiplicity(_438_v26)
					}
					return (_dafny.Zero).Minus(_438_v26)
				})()
				_ = _rhs35
				var _rhs36 _dafny.Set = (_453_v41).Union(_453_v41)
				_ = _rhs36
				var _rhs37 D2 = _454_v42
				_ = _rhs37
				var _rhs38 _dafny.Sequence = _dafny.SeqCreate((Companion_Default___.Abs(_dafny.Zero)).Uint32(), func(coer36 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg36 _dafny.Int) interface{} {
						return coer36(arg36)
					}
				}((func(_457_v26 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_458_i2 _dafny.Int) _dafny.Int {
						return _457_v26
					}
				})(_438_v26)))
				_ = _rhs38
				var _rhs39 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_455_v43, _455_v43)
				_ = _rhs39
				var _lhs20 *GlobalState = globalState
				_ = _lhs20
				_lhs20.F1 = _rhs35
				_453_v41 = _rhs36
				_454_v42 = _rhs37
				_440_v28 = _rhs38
				_455_v43 = _rhs39
				var _459_v45 _dafny.Array
				_ = _459_v45
				var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(15))
				_ = _nw68
				_459_v45 = _nw68
				_459_v45 = _459_v45
				var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index44
				((_this).F14()).ArraySet1(_439_v27, (_index44).Int())
				var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(550), _dafny.ArrayLen(((_this).F14()), 0))
				_ = _index45
				((_this).F14()).ArraySet1(_439_v27, (_index45).Int())
				(globalState).F1 = _438_v26
				var _460_v46 _dafny.Set
				_ = _460_v46
				_460_v46 = _dafny.SetOf(_438_v26, _dafny.IntOfInt64(666))
				var _461_v47 _dafny.Sequence
				_ = _461_v47
				_461_v47 = _dafny.SeqOf(_460_v46, _dafny.SetOf(_dafny.IntOfUint32((_409_v0).Cardinality()), Companion_Default___.Fm8(_dafny.IntOfInt64(688), _439_v27, !(_439_v27), globalState)))
				_439_v27 = ((_461_v47).Select((Companion_Default___.SafeIndex(_438_v26, _dafny.IntOfUint32((_461_v47).Cardinality()))).Uint32()).(_dafny.Set)).IsDisjointFrom(_dafny.SetOf((_456_v44).Cardinality()))
			}
			var _462_v48 _dafny.Array
			_ = _462_v48
			var _nw69 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
			_ = _nw69
			_462_v48 = _nw69
			var _463_v49 _dafny.Sequence
			_ = _463_v49
			_463_v49 = _dafny.SeqOf(_462_v48, _462_v48, _462_v48, _462_v48, _462_v48)
			var _rhs40 bool = _439_v27
			_ = _rhs40
			var _rhs41 _dafny.Int = _438_v26
			_ = _rhs41
			var _rhs42 bool = !(_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_462_v48, _462_v48), _463_v49), _462_v48))
			_ = _rhs42
			var _rhs43 bool = _439_v27
			_ = _rhs43
			var _rhs44 bool = _439_v27
			_ = _rhs44
			_439_v27 = _rhs40
			_438_v26 = _rhs41
			_439_v27 = _rhs42
			_439_v27 = _rhs43
			_439_v27 = _rhs44
			_439_v27 = !(!(_439_v27))
		}
		var _464_v50 _dafny.Int
		_ = _464_v50
		_464_v50 = _dafny.IntOfInt64(-178)
		(globalState).F1 = _464_v50
		var _465_v51 _dafny.Array
		_ = _465_v51
		var _nw70 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
		_ = _nw70
		_465_v51 = _nw70
		for _iter16 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_465_v51), 0))); ; {
			_guard_loop_1, _ok16 := _iter16()
			if !_ok16 {
				break
			}
			var _466_i3 _dafny.Int
			_466_i3 = interface{}(_guard_loop_1).(_dafny.Int)
			if (true) && (((_466_i3).Sign() != -1) && ((_466_i3).Cmp(_dafny.ArrayLen((_465_v51), 0)) < 0)) {
				(_465_v51).ArraySet1(_dafny.Companion_Sequence_.Equal(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_464_v50), _dafny.SeqOf(_464_v50)), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_464_v50, _464_v50, (_dafny.SetOf(_464_v50, _464_v50)).Cardinality(), _dafny.IntOfInt64(150)), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(397))).Uint32(), func(coer37 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg37 _dafny.Int) interface{} {
						return coer37(arg37)
					}
				}((func(_467_v50 _dafny.Int) func(_dafny.Int) _dafny.Int {
					return func(_468_i4 _dafny.Int) _dafny.Int {
						return _467_v50
					}
				})(_464_v50))))), (_466_i3).Int())
			}
		}
		var _469_v52 _dafny.Array
		_ = _469_v52
		var _nw71 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(25))
		_ = _nw71
		_469_v52 = _nw71
		for _iter17 := _dafny.Iterate(_dafny.IntegerRange(_dafny.Zero, _dafny.ArrayLen((_469_v52), 0))); ; {
			_guard_loop_2, _ok17 := _iter17()
			if !_ok17 {
				break
			}
			var _470_i5 _dafny.Int
			_470_i5 = interface{}(_guard_loop_2).(_dafny.Int)
			if (true) && (((_470_i5).Sign() != -1) && ((_470_i5).Cmp(_dafny.ArrayLen((_469_v52), 0)) < 0)) {
				(_469_v52).ArraySet1(_dafny.Companion_Sequence_.Update(_409_v0, (Companion_Default___.SafeIndex(_464_v50, _dafny.IntOfUint32((_409_v0).Cardinality()))).Uint32(), _dafny.CodePoint('r')), (_470_i5).Int())
			}
		}
		var _471_v53 bool
		_ = _471_v53
		_471_v53 = false
		if (func() bool {
			if _471_v53 {
				return _471_v53
			}
			return _dafny.Companion_Sequence_.IsPrefixOf(_409_v0, _409_v0)
		})() {
			var _472_v54 _dafny.Sequence
			_ = _472_v54
			_472_v54 = _dafny.SeqOf(_471_v53, _471_v53)
			var _473_v55 _dafny.Set
			_ = _473_v55
			_473_v55 = _dafny.SetOf(_471_v53, (_472_v54).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_472_v54).Cardinality()), _dafny.IntOfUint32((_472_v54).Cardinality()))).Uint32()).(bool))
			var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_469_v52), 0))
			_ = _index46
			(_469_v52).ArraySet1(_dafny.UnicodeSeqOfUtf8Bytes("ptockx"), (_index46).Int())
			var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_469_v52), 0))
			_ = _index47
			var _rhs45 _dafny.Set = _473_v55
			_ = _rhs45
			var _rhs46 _dafny.Sequence = _409_v0
			_ = _rhs46
			var _lhs21 _dafny.Array = _469_v52
			_ = _lhs21
			var _lhs22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_469_v52), 0))
			_ = _lhs22
			_473_v55 = _rhs45
			(_lhs21).ArraySet1(_rhs46, (_lhs22).Int())
			var _474_v56 _dafny.Array
			_ = _474_v56
			var _nw72 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(13))
			_ = _nw72
			_474_v56 = _nw72
			var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))
			_ = _index48
			(_474_v56).ArraySet1((_464_v50).Plus(_dafny.IntOfUint32(((_469_v52).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(86), _dafny.ArrayLen((_469_v52), 0))).Int()).(_dafny.Sequence)).Cardinality())), (_index48).Int())
			var _475_v57 _dafny.Map
			_ = _475_v57
			_475_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_471_v53, (_473_v55).Cardinality())
			var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))
			_ = _index49
			(_474_v56).ArraySet1((func() _dafny.Int {
				if _471_v53 {
					return Companion_Default___.SafeModuloInt((_475_v57).Cardinality(), _464_v50)
				}
				return (_464_v50).Times(_464_v50)
			})(), (_index49).Int())
			var _476_v58 _dafny.Map
			_ = _476_v58
			_476_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_474_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))).Int()).(_dafny.Int), !(_471_v53))
			var _477_v59 _dafny.Sequence
			_ = _477_v59
			_477_v59 = _dafny.SeqOf((_476_v58).Update(_464_v50, _471_v53))
			var _478_v60 _dafny.Sequence
			_ = _478_v60
			_478_v60 = _dafny.SeqOf(_477_v59)
			var _479_v61 _dafny.Sequence
			_ = _479_v61
			_479_v61 = _dafny.SeqOf(_dafny.IntOfInt64(-869), (_474_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))).Int()).(_dafny.Int))
			var _480_v62 _dafny.CodePoint
			_ = _480_v62
			_480_v62 = _dafny.CodePoint('n')
			var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))
			_ = _index50
			var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))
			_ = _index51
			var _rhs47 bool = _471_v53
			_ = _rhs47
			var _rhs48 _dafny.Int = (_474_v56).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))).Int()).(_dafny.Int)
			_ = _rhs48
			var _rhs49 _dafny.Sequence = (_478_v60).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.UnicodeSeqOfUtf8Bytes("jmbaqsn"), (Companion_Default___.SafeIndex((_479_v61).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(576), _dafny.IntOfUint32((_479_v61).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("jmbaqsn")).Cardinality()))).Uint32(), _480_v62)).Cardinality()), _dafny.IntOfUint32((_478_v60).Cardinality()))).Uint32()).(_dafny.Sequence)
			_ = _rhs49
			var _rhs50 bool = _471_v53
			_ = _rhs50
			var _rhs51 _dafny.Int = _dafny.IntOfInt64(-502)
			_ = _rhs51
			var _lhs23 _dafny.Array = _474_v56
			_ = _lhs23
			var _lhs24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))
			_ = _lhs24
			var _lhs25 _dafny.Array = _474_v56
			_ = _lhs25
			var _lhs26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(760), _dafny.ArrayLen((_474_v56), 0))
			_ = _lhs26
			_471_v53 = _rhs47
			(_lhs23).ArraySet1(_rhs48, (_lhs24).Int())
			_477_v59 = _rhs49
			_471_v53 = _rhs50
			(_lhs25).ArraySet1(_rhs51, (_lhs26).Int())
			var _481_v63 *C0
			_ = _481_v63
			var _nw73 *C0 = New_C0_()
			_ = _nw73
			_nw73.Ctor__(_471_v53, (func() bool {
				if _471_v53 {
					return _471_v53
				}
				return false
			})(), _dafny.CodePoint('c'), _dafny.Companion_Sequence_.Concatenate(_409_v0, _409_v0))
			_481_v63 = _nw73
			_480_v62 = _480_v62
		} else {
			var _482_v64 _dafny.Array
			_ = _482_v64
			var _nwElement0_10 _dafny.Array = _465_v51
			_ = _nwElement0_10
			var _nw74 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(2))
			_ = _nw74
			(_nw74).ArraySet1(_nwElement0_10, 0)
			(_nw74).ArraySet1((_this).F14(), 1)
			_482_v64 = _nw74
			var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_482_v64), 0))
			_ = _index52
			(_482_v64).ArraySet1((_this).F14(), (_index52).Int())
			var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(590), _dafny.ArrayLen((_482_v64), 0))
			_ = _index53
			(_482_v64).ArraySet1(_465_v51, (_index53).Int())
			var _483_v65 _dafny.CodePoint
			_ = _483_v65
			_483_v65 = _dafny.CodePoint('p')
			var _484_v66 *C0
			_ = _484_v66
			var _nw75 *C0 = New_C0_()
			_ = _nw75
			_nw75.Ctor__(!(true), _471_v53, _483_v65, _409_v0)
			_484_v66 = _nw75
			var _485_v67 _dafny.Map
			_ = _485_v67
			_485_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_471_v53) && (_471_v53), _471_v53)
			_485_v67 = (_485_v67).Merge(_485_v67)
			_471_v53 = (func() bool {
				if !(_471_v53) {
					return Companion_Default___.Fm7(_483_v65, _464_v50, _471_v53, globalState)
				}
				return _471_v53
			})()
			_483_v65 = _483_v65
		}
		r0 = _464_v50
		var _486_v68 _dafny.Array
		_ = _486_v68
		var _nw76 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(8))
		_ = _nw76
		_486_v68 = _nw76
		var _487_v69 D1
		_ = _487_v69
		_487_v69 = Companion_D1_.Create_DC3_(_486_v68)
		r1 = _487_v69
		var _488_v71 D3
		_ = _488_v71
		_488_v71 = Companion_D3_.Create_DC10_(_409_v0)
		var _489_v72 _dafny.Map
		_ = _489_v72
		_489_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_488_v71).Dtor_cf19(), _471_v53)
		var _490_v73 _dafny.Sequence
		_ = _490_v73
		_490_v73 = _dafny.SeqOf(_464_v50)
		var _491_v74 _dafny.Array
		_ = _491_v74
		var _nwElement0_11 _dafny.Int = (_464_v50).Minus(_464_v50)
		_ = _nwElement0_11
		var _nw77 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(16))
		_ = _nw77
		(_nw77).ArraySet1(_nwElement0_11, 0)
		(_nw77).ArraySet1(_464_v50, 1)
		(_nw77).ArraySet1(_464_v50, 2)
		(_nw77).ArraySet1((func() _dafny.Map {
			var _coll15 = _dafny.NewMapBuilder()
			_ = _coll15
			for _iter18 := _dafny.Iterate((_489_v72).Keys().Elements()); ; {
				_compr_15, _ok18 := _iter18()
				if !_ok18 {
					break
				}
				var _492_v70 _dafny.Sequence
				_492_v70 = interface{}(_compr_15).(_dafny.Sequence)
				if (_489_v72).Contains(_492_v70) {
					_coll15.Add(_492_v70, _464_v50)
				}
			}
			return _coll15.ToMap()
		}()).Cardinality(), 3)
		(_nw77).ArraySet1(Companion_Default___.SafeDivisionInt(_464_v50, _dafny.IntOfInt64(719)), 4)
		(_nw77).ArraySet1(_464_v50, 5)
		(_nw77).ArraySet1((func() _dafny.Int {
			if _471_v53 {
				return _dafny.IntOfInt64(233)
			}
			return _464_v50
		})(), 6)
		(_nw77).ArraySet1(_464_v50, 7)
		(_nw77).ArraySet1((_490_v73).Select((Companion_Default___.SafeIndex(_464_v50, _dafny.IntOfUint32((_490_v73).Cardinality()))).Uint32()).(_dafny.Int), 8)
		(_nw77).ArraySet1(_464_v50, 9)
		(_nw77).ArraySet1((func() _dafny.Int {
			if _471_v53 {
				return _464_v50
			}
			return _464_v50
		})(), 10)
		(_nw77).ArraySet1(_dafny.IntOfUint32((_409_v0).Cardinality()), 11)
		(_nw77).ArraySet1(_464_v50, 12)
		(_nw77).ArraySet1((_490_v73).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(854), _dafny.IntOfUint32((_490_v73).Cardinality()))).Uint32()).(_dafny.Int), 13)
		(_nw77).ArraySet1(_464_v50, 14)
		(_nw77).ArraySet1(_dafny.IntOfUint32((_409_v0).Cardinality()), 15)
		_491_v74 = _nw77
		r2 = _491_v74
		var _pat_let_tv8 = _486_v68
		_ = _pat_let_tv8
		r3 = func(_pat_let11_0 D1) D1 {
			return func(_493_dt__update__tmp_h0 D1) D1 {
				return func(_pat_let12_0 _dafny.Array) D1 {
					return func(_494_dt__update_hcf7_h0 _dafny.Array) D1 {
						return Companion_D1_.Create_DC3_(_494_dt__update_hcf7_h0)
					}(_pat_let12_0)
				}(_pat_let_tv8)
			}(_pat_let11_0)
		}(_487_v69)
		return r0, r1, r2, r3
	}
}
func (_this *C2) M7(p0 D1, p1 _dafny.Map, p2 _dafny.Map, globalState *GlobalState) (_dafny.Int, bool, _dafny.Int) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 bool = false
		_ = r1
		var r2 _dafny.Int = _dafny.Zero
		_ = r2
		var _495_v0 bool
		_ = _495_v0
		_495_v0 = false
		var _496_v1 _dafny.Int
		_ = _496_v1
		_496_v1 = _dafny.IntOfInt64(64)
		var _497_v2 D0
		_ = _497_v2
		_497_v2 = Companion_D0_.Create_DC1_(_495_v0, (func() _dafny.Int {
			if _495_v0 {
				return _496_v1
			}
			return _496_v1
		})(), _495_v0, _496_v1)
		var _source6 D0 = _497_v2
		_ = _source6
		if _source6.Is_DC0() {
			if (_496_v1).Cmp(_496_v1) == 0 {
				var _498_v3 _dafny.MultiSet
				_ = _498_v3
				_498_v3 = _dafny.MultiSetOf(_496_v1)
				var _499_v4 _dafny.MultiSet
				_ = _499_v4
				_499_v4 = _dafny.MultiSetOf(_496_v1, (_498_v3).Cardinality(), Companion_Default___.SafeModuloInt(_496_v1, Companion_Default___.Fm8(_dafny.IntOfInt64(-351), _495_v0, _495_v0, globalState)))
				_498_v3 = _498_v3
				r2 = _496_v1
				var _500_v5 _dafny.Sequence
				_ = _500_v5
				_500_v5 = _dafny.UnicodeSeqOfUtf8Bytes("vcmororjs")
				var _501_v6 *C0
				_ = _501_v6
				var _nw78 *C0 = New_C0_()
				_ = _nw78
				_nw78.Ctor__(_495_v0, _495_v0, _dafny.CodePoint('q'), _500_v5)
				_501_v6 = _nw78
				r1 = (func() bool {
					if true {
						return _495_v0
					}
					return true
				})()
				_501_v6 = _501_v6
			} else {
				r2 = Companion_Default___.SafeModuloInt(_496_v1, Companion_Default___.Fm8(_496_v1, _495_v0, _495_v0, globalState))
				var _502_v7 _dafny.Array
				_ = _502_v7
				var _nw79 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(2))
				_ = _nw79
				_502_v7 = _nw79
				var _len0_10 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_10
				var _nw80 _dafny.Array
				_ = _nw80
				if _len0_10.Cmp(_dafny.Zero) == 0 {
					_nw80 = _dafny.NewArray(_len0_10)
				} else {
					var _init10 func(_dafny.Int) _dafny.Int = func(_503_i0 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_503_i0, _dafny.IntOfInt64(696))
					}
					_ = _init10
					var _element0_10 = _init10(_dafny.Zero)
					_ = _element0_10
					_nw80 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
					(_nw80).ArraySet1(_element0_10, 0)
					var _nativeLen0_10 = (_len0_10).Int()
					_ = _nativeLen0_10
					for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
						(_nw80).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
					}
				}
				_502_v7 = _nw80
				var _504_v8 _dafny.MultiSet
				_ = _504_v8
				_504_v8 = _dafny.MultiSetOf(!(_495_v0) || (_495_v0))
				_504_v8 = _504_v8
				var _505_v9 _dafny.Int
				_ = _505_v9
				var _506_v10 D1
				_ = _506_v10
				var _507_v11 _dafny.Array
				_ = _507_v11
				var _508_v12 D1
				_ = _508_v12
				var _out17 _dafny.Int
				_ = _out17
				var _out18 D1
				_ = _out18
				var _out19 _dafny.Array
				_ = _out19
				var _out20 D1
				_ = _out20
				_out17, _out18, _out19, _out20 = (_this).M6(globalState)
				_505_v9 = _out17
				_506_v10 = _out18
				_507_v11 = _out19
				_508_v12 = _out20
				_495_v0 = (_505_v9).Cmp(_496_v1) == 0
			}
			var _509_v13 _dafny.Sequence
			_ = _509_v13
			_509_v13 = _dafny.SeqOf((p0).Dtor_cf8())
			var _510_v14 _dafny.Sequence
			_ = _510_v14
			_510_v14 = _dafny.SeqOf(_495_v0, _dafny.Companion_Sequence_.Contains(_509_v13, (_509_v13).Select((Companion_Default___.SafeIndex(_496_v1, _dafny.IntOfUint32((_509_v13).Cardinality()))).Uint32()).(_dafny.CodePoint)))
			_495_v0 = (_510_v14).Select((Companion_Default___.SafeIndex(_496_v1, _dafny.IntOfUint32((_510_v14).Cardinality()))).Uint32()).(bool)
			(globalState).F1 = _496_v1
			var _511_v15 D3
			_ = _511_v15
			_511_v15 = Companion_D3_.Create_DC10_(_dafny.Companion_Sequence_.Concatenate(_509_v13, _509_v13))
			_511_v15 = Companion_D3_.Create_DC10_(_509_v13)
		} else if _source6.Is_DC1() {
			var _512___mcc_h0 bool = _source6.Get_().(D0_DC1).Cf0
			_ = _512___mcc_h0
			var _513___mcc_h1 _dafny.Int = _source6.Get_().(D0_DC1).Cf1
			_ = _513___mcc_h1
			var _514___mcc_h2 bool = _source6.Get_().(D0_DC1).Cf2
			_ = _514___mcc_h2
			var _515___mcc_h3 _dafny.Int = _source6.Get_().(D0_DC1).Cf3
			_ = _515___mcc_h3
			var _516_cf3 _dafny.Int = _515___mcc_h3
			_ = _516_cf3
			var _517_cf2 bool = _514___mcc_h2
			_ = _517_cf2
			var _518_cf1 _dafny.Int = _513___mcc_h1
			_ = _518_cf1
			var _519_cf0 bool = _512___mcc_h0
			_ = _519_cf0
			var _520_v16 _dafny.Set
			_ = _520_v16
			_520_v16 = _dafny.SetOf(_495_v0)
			var _521_v17 _dafny.CodePoint
			_ = _521_v17
			_521_v17 = _dafny.CodePoint('i')
			var _522_v18 _dafny.Map
			_ = _522_v18
			_522_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v1, _516_cf3)
			var _523_v19 _dafny.Set
			_ = _523_v19
			_523_v19 = _dafny.SetOf(_518_cf1, (_522_v18).Cardinality())
			r0 = (((func() _dafny.Set {
				if _517_cf2 {
					return _520_v16
				}
				return _520_v16
			})()).Intersection((_dafny.SetOf(Companion_Default___.Fm7(_521_v17, (_523_v19).Cardinality(), _519_cf0, globalState), _517_cf2)).Union(_520_v16))).Cardinality()
			r1 = (_518_cf1).Cmp((_dafny.Zero).Minus(_496_v1)) == 0
			var _524_v20 _dafny.Array
			_ = _524_v20
			var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(29))
			_ = _nw81
			_524_v20 = _nw81
			var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))
			_ = _index54
			(_524_v20).ArraySet1(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(961))).Uint32(), func(coer38 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg38 _dafny.Int) interface{} {
					return coer38(arg38)
				}
			}((func(_525_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_526_i1 _dafny.Int) _dafny.CodePoint {
					return _525_v17
				}
			})(_521_v17))), (_index54).Int())
			var _527_v21 _dafny.Sequence
			_ = _527_v21
			_527_v21 = _dafny.UnicodeSeqOfUtf8Bytes("elp")
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))
			_ = _index55
			(_524_v20).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_527_v21, _527_v21), _527_v21), (_index55).Int())
			if _517_cf2 {
				var _528_v22 *C0
				_ = _528_v22
				var _nw82 *C0 = New_C0_()
				_ = _nw82
				_nw82.Ctor__(_519_cf0, _495_v0, _521_v17, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(792))).Uint32(), func(coer39 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg39 _dafny.Int) interface{} {
						return coer39(arg39)
					}
				}((func(_529_v17 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_530_i2 _dafny.Int) _dafny.CodePoint {
						return _529_v17
					}
				})(_521_v17))))
				_528_v22 = _nw82
				var _531_v23 D2
				_ = _531_v23
				_531_v23 = Companion_D2_.Create_DC7_(_517_cf2)
				var _532_v24 _dafny.Map
				_ = _532_v24
				_532_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_528_v22, (_this).F14())
				var _533_v25 _dafny.Array
				_ = _533_v25
				var _nw83 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw83
				_533_v25 = _nw83
				var _534_v26 _dafny.Sequence
				_ = _534_v26
				_534_v26 = _dafny.SeqOf(_519_cf0)
				var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_533_v25), 0))
				_ = _index56
				(_533_v25).ArraySet1(_534_v26, (_index56).Int())
				var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_533_v25), 0))
				_ = _index57
				var _rhs52 *C0 = (func() *C0 {
					if (_dafny.IntOfInt64(966)).Cmp(_496_v1) == 0 {
						return _528_v22
					}
					return _528_v22
				})()
				_ = _rhs52
				var _rhs53 D2 = Companion_Default___.Fm11(Companion_Default___.Fm12(_519_cf0, _527_v21, _517_cf2, globalState), _517_cf2, ((_dafny.Zero).Minus(_496_v1)).Minus(_dafny.IntOfInt64(810)), _495_v0, globalState)
				_ = _rhs53
				var _rhs54 _dafny.Map = _532_v24
				_ = _rhs54
				var _rhs55 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_534_v26, (Companion_Default___.SafeIndex(_516_cf3, _dafny.IntOfUint32((_534_v26).Cardinality()))).Uint32(), _519_cf0)
				_ = _rhs55
				var _lhs27 _dafny.Array = _533_v25
				_ = _lhs27
				var _lhs28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(410), _dafny.ArrayLen((_533_v25), 0))
				_ = _lhs28
				_528_v22 = _rhs52
				_531_v23 = _rhs53
				_532_v24 = _rhs54
				(_lhs27).ArraySet1(_rhs55, (_lhs28).Int())
				_519_cf0 = _495_v0
				var _535_v27 _dafny.Sequence
				_ = _535_v27
				_535_v27 = _dafny.SeqOf(_518_cf1)
				var _536_v28 _dafny.Map
				_ = _536_v28
				_536_v28 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_495_v0, _519_cf0)
				var _537_v29 _dafny.Set
				_ = _537_v29
				_537_v29 = _dafny.SetOf((_524_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))).Int()).(_dafny.Sequence))
				var _538_v30 _dafny.Array
				_ = _538_v30
				var _nwElement0_12 _dafny.Int = _518_cf1
				_ = _nwElement0_12
				var _nw84 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(21))
				_ = _nw84
				(_nw84).ArraySet1(_nwElement0_12, 0)
				(_nw84).ArraySet1(((_535_v27).Select((Companion_Default___.SafeIndex(_516_cf3, _dafny.IntOfUint32((_535_v27).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_516_cf3), 1)
				(_nw84).ArraySet1(_518_cf1, 2)
				(_nw84).ArraySet1(_516_cf3, 3)
				(_nw84).ArraySet1(_516_cf3, 4)
				(_nw84).ArraySet1(_dafny.IntOfUint32((_527_v21).Cardinality()), 5)
				(_nw84).ArraySet1(_496_v1, 6)
				(_nw84).ArraySet1((func() _dafny.Int {
					if (func() bool {
						if (_536_v28).Contains(_495_v0) {
							return (_536_v28).Get(_495_v0).(bool)
						}
						return _519_cf0
					})() {
						return _496_v1
					}
					return _518_cf1
				})(), 7)
				(_nw84).ArraySet1((_537_v29).Cardinality(), 8)
				(_nw84).ArraySet1((_518_cf1).Times(_518_cf1), 9)
				(_nw84).ArraySet1(Companion_Default___.SafeDivisionInt(_516_cf3, _516_cf3), 10)
				(_nw84).ArraySet1(_518_cf1, 11)
				(_nw84).ArraySet1((func() _dafny.Int {
					if _517_cf2 {
						return _496_v1
					}
					return (_520_v16).Cardinality()
				})(), 12)
				(_nw84).ArraySet1(_518_cf1, 13)
				(_nw84).ArraySet1(_516_cf3, 14)
				(_nw84).ArraySet1((_518_cf1).Minus(_496_v1), 15)
				(_nw84).ArraySet1((_dafny.Zero).Minus(_518_cf1), 16)
				(_nw84).ArraySet1(Companion_Default___.Fm8((_dafny.Zero).Minus(_518_cf1), _495_v0, _495_v0, globalState), 17)
				(_nw84).ArraySet1((_dafny.MultiSetOf(false)).Cardinality(), 18)
				(_nw84).ArraySet1(_518_cf1, 19)
				(_nw84).ArraySet1(_516_cf3, 20)
				_538_v30 = _nw84
				var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_538_v30), 0))
				_ = _index58
				(_538_v30).ArraySet1(_516_cf3, (_index58).Int())
				var _539_v31 _dafny.MultiSet
				_ = _539_v31
				_539_v31 = _dafny.MultiSetOf(_521_v17)
				var _540_v32 _dafny.Map
				_ = _540_v32
				_540_v32 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_521_v17, _517_cf2)
				var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_538_v30), 0))
				_ = _index59
				var _rhs56 bool = Companion_Default___.Fm7(_521_v17, (func() _dafny.Int {
					if (_539_v31).Contains((Companion_D1_.Create_DC4_(_dafny.CodePoint('j'), _496_v1, _521_v17, _516_cf3, true)).Dtor_cf10()) {
						return (_539_v31).Multiplicity((Companion_D1_.Create_DC4_(_dafny.CodePoint('j'), _496_v1, _521_v17, _516_cf3, true)).Dtor_cf10())
					}
					return _dafny.IntOfInt64(204)
				})(), !((_496_v1).Cmp(_518_cf1) > 0), globalState)
				_ = _rhs56
				var _rhs57 bool = (((_540_v32).Cardinality()).Plus(_516_cf3)).Cmp(_516_cf3) >= 0
				_ = _rhs57
				var _rhs58 _dafny.Int = (_dafny.Zero).Minus((_535_v27).Select((Companion_Default___.SafeIndex(_496_v1, _dafny.IntOfUint32((_535_v27).Cardinality()))).Uint32()).(_dafny.Int))
				_ = _rhs58
				var _rhs59 _dafny.Int = _516_cf3
				_ = _rhs59
				var _lhs29 _dafny.Array = _538_v30
				_ = _lhs29
				var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_538_v30), 0))
				_ = _lhs30
				r1 = _rhs56
				_495_v0 = _rhs57
				(_lhs29).ArraySet1(_rhs58, (_lhs30).Int())
				r2 = _rhs59
				var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(193), _dafny.ArrayLen((_538_v30), 0))
				_ = _index60
				(_538_v30).ArraySet1(_518_cf1, (_index60).Int())
				_496_v1 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_524_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))).Int()).(_dafny.Sequence), _527_v21)).Cardinality()), _496_v1)
			} else {
				var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))
				_ = _index61
				(_524_v20).ArraySet1((_524_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))).Int()).(_dafny.Sequence), (_index61).Int())
				r1 = _495_v0
				_527_v21 = _dafny.Companion_Sequence_.Concatenate(_527_v21, _dafny.Companion_Sequence_.Update(_527_v21, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_dafny.SeqOf(_516_cf3)).Cardinality()), _dafny.IntOfUint32((_527_v21).Cardinality()))).Uint32(), _521_v17))
				var _541_v33 *C0
				_ = _541_v33
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__(_519_cf0, (_516_cf3).Cmp(Companion_Default___.Fm8(_516_cf3, _517_cf2, !(_517_cf2), globalState)) == 0, _521_v17, (_524_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(35), _dafny.ArrayLen((_524_v20), 0))).Int()).(_dafny.Sequence))
				_541_v33 = _nw85
				var _nw86 *C0 = New_C0_()
				_ = _nw86
				_nw86.Ctor__(_495_v0, _495_v0, Companion_Default___.Fm13(globalState), Companion_Default___.Fm14(_496_v1, _516_cf3, _519_cf0, true, globalState))
				_541_v33 = _nw86
				var _542_v34 D3
				_ = _542_v34
				_542_v34 = Companion_D3_.Create_DC10_(_527_v21)
				var _543_v35 D3
				_ = _543_v35
				_543_v35 = Companion_D3_.Create_DC12_(_542_v34)
				var _544_v36 _dafny.MultiSet
				_ = _544_v36
				_544_v36 = _dafny.MultiSetOf(_516_cf3)
				var _545_v37 _dafny.Sequence
				_ = _545_v37
				_545_v37 = _dafny.SeqOf(_517_cf2)
				var _546_v38 _dafny.Map
				_ = _546_v38
				_546_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_544_v36, _545_v37)
				var _547_v39 _dafny.Map
				_ = _547_v39
				_547_v39 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_543_v35, _546_v38)
				_547_v39 = (_547_v39).Update(Companion_D3_.Create_DC12_(_542_v34), _546_v38)
			}
		} else {
			var _548___mcc_h4 _dafny.Int = _source6.Get_().(D0_DC2).Cf4
			_ = _548___mcc_h4
			var _549___mcc_h5 bool = _source6.Get_().(D0_DC2).Cf5
			_ = _549___mcc_h5
			var _550___mcc_h6 _dafny.Sequence = _source6.Get_().(D0_DC2).Cf6
			_ = _550___mcc_h6
			var _551_cf6 _dafny.Sequence = _550___mcc_h6
			_ = _551_cf6
			var _552_cf5 bool = _549___mcc_h5
			_ = _552_cf5
			var _553_cf4 _dafny.Int = _548___mcc_h4
			_ = _553_cf4
			_497_v2 = _497_v2
			_495_v0 = (_dafny.IntOfInt64(619)).Cmp(_553_cf4) > 0
			(globalState).F1 = (_dafny.Zero).Minus(_553_cf4)
			Companion_Default___.M0(globalState)
		}
		r1 = _495_v0
		var _554_v40 _dafny.CodePoint
		_ = _554_v40
		_554_v40 = _dafny.CodePoint('f')
		var _555_v41 _dafny.Map
		_ = _555_v41
		_555_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_495_v0, _554_v40)
		var _556_v42 *C0
		_ = _556_v42
		var _nw87 *C0 = New_C0_()
		_ = _nw87
		_nw87.Ctor__(true, !(_495_v0), (func() _dafny.CodePoint {
			if (_555_v41).Contains(_495_v0) {
				return (_555_v41).Get(_495_v0).(_dafny.CodePoint)
			}
			return _554_v40
		})(), Companion_Default___.Fm14(_dafny.IntOfInt64(-567), _496_v1, Companion_Default___.Fm7(_554_v40, _496_v1, _495_v0, globalState), _495_v0, globalState))
		_556_v42 = _nw87
		_496_v1 = _496_v1
		if (_495_v0) && (_495_v0) {
			var _557_v43 _dafny.Map
			_ = _557_v43
			_557_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_495_v0), Companion_Default___.Fm15(globalState))
			_557_v43 = (_557_v43).Merge(_557_v43)
			if _495_v0 {
				var _558_v44 _dafny.Set
				_ = _558_v44
				_558_v44 = _dafny.SetOf(true, true)
				r2 = (_496_v1).Times(Companion_Default___.SafeModuloInt((_558_v44).Cardinality(), _496_v1))
				r0 = _496_v1
				(globalState).F1 = ((_dafny.Zero).Minus(_496_v1)).Minus(_496_v1)
				var _559_v45 D1
				_ = _559_v45
				_559_v45 = Companion_D1_.Create_DC5_(_496_v1)
				var _pat_let_tv9 = _496_v1
				_ = _pat_let_tv9
				var _pat_let_tv10 = _496_v1
				_ = _pat_let_tv10
				var _560_v46 _dafny.Sequence
				_ = _560_v46
				_560_v46 = _dafny.SeqOf(_559_v45, func(_pat_let13_0 D1) D1 {
					return func(_561_dt__update__tmp_h0 D1) D1 {
						return func(_pat_let14_0 _dafny.Int) D1 {
							return func(_562_dt__update_hcf13_h0 _dafny.Int) D1 {
								return Companion_D1_.Create_DC5_(_562_dt__update_hcf13_h0)
							}(_pat_let14_0)
						}(_pat_let_tv9)
					}(_pat_let13_0)
				}(_559_v45), func(_pat_let15_0 D1) D1 {
					return func(_563_dt__update__tmp_h1 D1) D1 {
						return func(_pat_let16_0 _dafny.Int) D1 {
							return func(_564_dt__update_hcf13_h1 _dafny.Int) D1 {
								return Companion_D1_.Create_DC5_(_564_dt__update_hcf13_h1)
							}(_pat_let16_0)
						}(_pat_let_tv10)
					}(_pat_let15_0)
				}(_559_v45))
				var _565_v47 _dafny.Sequence
				_ = _565_v47
				_565_v47 = Companion_Default___.Fm16(globalState)
				var _566_v48 _dafny.Map
				_ = _566_v48
				_566_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v1, (_565_v47))
				var _567_v49 _dafny.Sequence
				_ = _567_v49
				_567_v49 = _dafny.SeqOf(Companion_Default___.Fm16(globalState), _560_v46)
				var _568_v50 _dafny.Sequence
				_ = _568_v50
				_568_v50 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_560_v46, _dafny.SeqOf(_559_v45)), (func() _dafny.Sequence {
					if (_566_v48).Contains(_496_v1) {
						return (_566_v48).Get(_496_v1).(_dafny.Sequence)
					}
					return _560_v46
				})(), (_567_v49).Select((Companion_Default___.SafeIndex(_496_v1, _dafny.IntOfUint32((_567_v49).Cardinality()))).Uint32()).(_dafny.Sequence))
				_568_v50 = _567_v49
				var _569_v51 _dafny.Array
				_ = _569_v51
				var _nw88 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw88
				_569_v51 = _nw88
				var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_569_v51), 0))
				_ = _index62
				(_569_v51).ArraySet1(_496_v1, (_index62).Int())
				var _570_v52 _dafny.Map
				_ = _570_v52
				_570_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_554_v40, _495_v0)
				var _571_v53 _dafny.Map
				_ = _571_v53
				_571_v53 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_495_v0, _495_v0)
				var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_569_v51), 0))
				_ = _index63
				var _rhs60 _dafny.Int = _496_v1
				_ = _rhs60
				var _rhs61 _dafny.Map = Companion_Default___.Fm17(_496_v1, globalState)
				_ = _rhs61
				var _rhs62 _dafny.CodePoint = _dafny.CodePoint('j')
				_ = _rhs62
				var _rhs63 bool = ((_496_v1).Minus((_571_v53).Cardinality())).Cmp(_496_v1) == 0
				_ = _rhs63
				var _rhs64 _dafny.Int = _496_v1
				_ = _rhs64
				var _lhs31 _dafny.Array = _569_v51
				_ = _lhs31
				var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(742), _dafny.ArrayLen((_569_v51), 0))
				_ = _lhs32
				var _lhs33 *GlobalState = globalState
				_ = _lhs33
				(_lhs31).ArraySet1(_rhs60, (_lhs32).Int())
				_570_v52 = _rhs61
				_554_v40 = _rhs62
				_495_v0 = _rhs63
				_lhs33.F1 = _rhs64
			} else {
				r1 = _495_v0
				var _rhs65 _dafny.Int = _496_v1
				_ = _rhs65
				var _rhs66 _dafny.Int = _496_v1
				_ = _rhs66
				var _lhs34 *GlobalState = globalState
				_ = _lhs34
				var _lhs35 *GlobalState = globalState
				_ = _lhs35
				_lhs34.F1 = _rhs65
				_lhs35.F1 = _rhs66
				var _572_v54 _dafny.Sequence
				_ = _572_v54
				_572_v54 = _dafny.SeqOf(_495_v0)
				var _573_v55 _dafny.Sequence
				_ = _573_v55
				_573_v55 = _dafny.UnicodeSeqOfUtf8Bytes("kossprq")
				var _574_v56 _dafny.Map
				_ = _574_v56
				_574_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm7(_554_v40, _dafny.IntOfUint32((_572_v54).Cardinality()), !(_495_v0), globalState), _573_v55)
				_574_v56 = (_574_v56).Update(_495_v0, _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("fhcrvh"), _573_v55))
				r2 = _496_v1
				var _575_v57 _dafny.Array
				_ = _575_v57
				var _nwElement0_13 _dafny.Int = _dafny.IntOfInt64(-814)
				_ = _nwElement0_13
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.IntOfInt64(5))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_13, 0)
				(_nw89).ArraySet1((_496_v1).Minus(_496_v1), 1)
				(_nw89).ArraySet1(_496_v1, 2)
				(_nw89).ArraySet1(_496_v1, 3)
				(_nw89).ArraySet1(_dafny.IntOfInt64(208), 4)
				_575_v57 = _nw89
				var _576_v58 D5
				_ = _576_v58
				_576_v58 = Companion_D5_.Create_DC14_(_575_v57)
				_575_v57 = (_576_v58).Dtor_cf25()
			}
			r1 = !(true)
			r1 = _495_v0
			var _577_v59 _dafny.Array
			_ = _577_v59
			var _nw90 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(19))
			_ = _nw90
			_577_v59 = _nw90
			var _578_v60 D1
			_ = _578_v60
			_578_v60 = Companion_D1_.Create_DC3_(_577_v59)
			var _source7 D1 = _578_v60
			_ = _source7
			if _source7.Is_DC4() {
				var _579___mcc_h7 _dafny.CodePoint = _source7.Get_().(D1_DC4).Cf8
				_ = _579___mcc_h7
				var _580___mcc_h8 _dafny.Int = _source7.Get_().(D1_DC4).Cf9
				_ = _580___mcc_h8
				var _581___mcc_h9 _dafny.CodePoint = _source7.Get_().(D1_DC4).Cf10
				_ = _581___mcc_h9
				var _582___mcc_h10 _dafny.Int = _source7.Get_().(D1_DC4).Cf11
				_ = _582___mcc_h10
				var _583___mcc_h11 bool = _source7.Get_().(D1_DC4).Cf12
				_ = _583___mcc_h11
				var _584_cf12 bool = _583___mcc_h11
				_ = _584_cf12
				var _585_cf11 _dafny.Int = _582___mcc_h10
				_ = _585_cf11
				var _586_cf10 _dafny.CodePoint = _581___mcc_h9
				_ = _586_cf10
				var _587_cf9 _dafny.Int = _580___mcc_h8
				_ = _587_cf9
				var _588_cf8 _dafny.CodePoint = _579___mcc_h7
				_ = _588_cf8
				var _589_v61 _dafny.Array
				_ = _589_v61
				var _nw91 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
				_ = _nw91
				_589_v61 = _nw91
				var _590_v62 _dafny.Sequence
				_ = _590_v62
				_590_v62 = _dafny.UnicodeSeqOfUtf8Bytes("q")
				var _591_v63 _dafny.Map
				_ = _591_v63
				_591_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_590_v62, _584_cf12)
				var _592_v64 _dafny.Map
				_ = _592_v64
				_592_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_590_v62).Cardinality()), _495_v0)
				var _593_v65 _dafny.Map
				_ = _593_v65
				_593_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_495_v0, (_592_v64).Cardinality())
				var _594_v66 _dafny.MultiSet
				_ = _594_v66
				_594_v66 = _dafny.MultiSetOf(_556_v42)
				var _595_v67 _dafny.Sequence
				_ = _595_v67
				_595_v67 = _dafny.SeqOf(_585_cf11, _587_cf9, _496_v1, (_dafny.Zero).Minus(_496_v1))
				var _596_v68 _dafny.Array
				_ = _596_v68
				var _nwElement0_14 _dafny.Int = _587_cf9
				_ = _nwElement0_14
				var _nw92 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(22))
				_ = _nw92
				(_nw92).ArraySet1(_nwElement0_14, 0)
				(_nw92).ArraySet1(_587_cf9, 1)
				(_nw92).ArraySet1(_496_v1, 2)
				(_nw92).ArraySet1(_496_v1, 3)
				(_nw92).ArraySet1(_585_cf11, 4)
				(_nw92).ArraySet1(_585_cf11, 5)
				(_nw92).ArraySet1(_585_cf11, 6)
				(_nw92).ArraySet1(_585_cf11, 7)
				(_nw92).ArraySet1(_585_cf11, 8)
				(_nw92).ArraySet1(_587_cf9, 9)
				(_nw92).ArraySet1(((_591_v63).Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(562))).Uint32(), func(coer40 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg40 _dafny.Int) interface{} {
						return coer40(arg40)
					}
				}((func(_597_cf10 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
					return func(_598_i3 _dafny.Int) _dafny.CodePoint {
						return _597_cf10
					}
				})(_586_cf10))), _584_cf12)).Cardinality(), 10)
				(_nw92).ArraySet1(_587_cf9, 11)
				(_nw92).ArraySet1((_dafny.Zero).Minus(_dafny.IntOfUint32((_590_v62).Cardinality())), 12)
				(_nw92).ArraySet1(_496_v1, 13)
				(_nw92).ArraySet1(_585_cf11, 14)
				(_nw92).ArraySet1(_dafny.IntOfUint32((_590_v62).Cardinality()), 15)
				(_nw92).ArraySet1(_496_v1, 16)
				(_nw92).ArraySet1(_496_v1, 17)
				(_nw92).ArraySet1((_593_v65).Cardinality(), 18)
				(_nw92).ArraySet1((_594_v66).Cardinality(), 19)
				(_nw92).ArraySet1(_585_cf11, 20)
				(_nw92).ArraySet1(_dafny.IntOfUint32((_595_v67).Cardinality()), 21)
				_596_v68 = _nw92
				var _599_v69 _dafny.Map
				_ = _599_v69
				_599_v69 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_589_v61, _596_v68)
				_599_v69 = (_599_v69).Update(_589_v61, _596_v68)
				var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_596_v68), 0))
				_ = _index64
				(_596_v68).ArraySet1(_496_v1, (_index64).Int())
				var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(349), _dafny.ArrayLen((_596_v68), 0))
				_ = _index65
				(_596_v68).ArraySet1(_496_v1, (_index65).Int())
				_586_cf10 = _586_cf10
				var _600_v70 _dafny.Map
				_ = _600_v70
				_600_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v1, (func() _dafny.Int {
					if _495_v0 {
						return _496_v1
					}
					return _dafny.IntOfInt64(659)
				})())
				_600_v70 = (_600_v70).Update(_585_cf11, (_496_v1).Times(_587_cf9))
			} else if _source7.Is_DC5() {
				var _601___mcc_h12 _dafny.Int = _source7.Get_().(D1_DC5).Cf13
				_ = _601___mcc_h12
				var _602_cf13 _dafny.Int = _601___mcc_h12
				_ = _602_cf13
				var _603_v71 _dafny.Array
				_ = _603_v71
				var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(28))
				_ = _nw93
				_603_v71 = _nw93
				_603_v71 = _603_v71
				var _604_v72 _dafny.Map
				_ = _604_v72
				_604_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.SafeDivisionInt(_602_cf13, _602_cf13), (_this).F14())
				var _605_v73 _dafny.Sequence
				_ = _605_v73
				_605_v73 = _dafny.SeqOf(!(_495_v0), !(_495_v0))
				var _606_v74 _dafny.Sequence
				_ = _606_v74
				_606_v74 = _dafny.SeqOf(_dafny.IntOfInt64(954), _dafny.IntOfUint32((_605_v73).Cardinality()))
				var _607_v75 _dafny.Map
				_ = _607_v75
				_607_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_554_v40, _495_v0)
				var _608_v76 D3
				_ = _608_v76
				_608_v76 = Companion_D3_.Create_DC11_(_495_v0, _dafny.IntOfUint32((_606_v74).Cardinality()), _dafny.SetOf(_495_v0, (func() bool {
					if (_607_v75).Contains(_554_v40) {
						return (_607_v75).Get(_554_v40).(bool)
					}
					return _495_v0
				})()))
				_604_v72 = (_604_v72).Update((_608_v76).Dtor_cf21(), (_this).F14())
				var _609_v77 _dafny.Array
				_ = _609_v77
				var _len0_11 _dafny.Int = _dafny.IntOfInt64(4)
				_ = _len0_11
				var _nw94 _dafny.Array
				_ = _nw94
				if _len0_11.Cmp(_dafny.Zero) == 0 {
					_nw94 = _dafny.NewArray(_len0_11)
				} else {
					var _init11 func(_dafny.Int) _dafny.Int = (func(_610_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.Int {
						return func(_611_i4 _dafny.Int) _dafny.Int {
							return Companion_Default___.SafeModuloInt(_611_i4, (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(169))).Uint32(), func(coer41 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
								return func(arg41 _dafny.Int) interface{} {
									return coer41(arg41)
								}
							}((func(_612_v40 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
								return func(_613_i5 _dafny.Int) _dafny.CodePoint {
									return _612_v40
								}
							})(_610_v40)))).Cardinality())))
						}
					})(_554_v40)
					_ = _init11
					var _element0_11 = _init11(_dafny.Zero)
					_ = _element0_11
					_nw94 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
					(_nw94).ArraySet1(_element0_11, 0)
					var _nativeLen0_11 = (_len0_11).Int()
					_ = _nativeLen0_11
					for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
						(_nw94).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
					}
				}
				_609_v77 = _nw94
				var _614_v78 _dafny.Map
				_ = _614_v78
				_614_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_609_v77, _495_v0)
				_614_v78 = (_614_v78).Update(_609_v77, _495_v0)
				var _615_v79 _dafny.Sequence
				_ = _615_v79
				_615_v79 = _dafny.UnicodeSeqOfUtf8Bytes("l")
				_615_v79 = _615_v79
			} else {
				var _616___mcc_h13 _dafny.Array = _source7.Get_().(D1_DC3).Cf7
				_ = _616___mcc_h13
				var _617_cf7 _dafny.Array = _616___mcc_h13
				_ = _617_cf7
				var _618_v80 _dafny.Sequence
				_ = _618_v80
				_618_v80 = _dafny.UnicodeSeqOfUtf8Bytes("egghuib")
				_618_v80 = _618_v80
				var _619_v81 _dafny.Map
				_ = _619_v81
				_619_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_618_v80).Cardinality()), _495_v0)
				var _620_v82 _dafny.Sequence
				_ = _620_v82
				_620_v82 = _dafny.SeqOf(_496_v1, (_619_v81).Cardinality())
				_620_v82 = _620_v82
				var _621_v83 *C0
				_ = _621_v83
				var _nw95 *C0 = New_C0_()
				_ = _nw95
				_nw95.Ctor__(_495_v0, false, _554_v40, _618_v80)
				_621_v83 = _nw95
				_621_v83 = _621_v83
				_495_v0 = (_496_v1).Cmp(_496_v1) == 0
			}
		} else {
			var _622_v84 _dafny.Map
			_ = _622_v84
			_622_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _dafny.IntOfUint32((_dafny.SeqOf(_496_v1)).Cardinality()))
			var _623_v85 _dafny.Sequence
			_ = _623_v85
			_623_v85 = _dafny.SeqOf(_495_v0, _495_v0)
			var _624_v86 _dafny.Map
			_ = _624_v86
			_624_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_495_v0, _495_v0)
			var _625_v87 _dafny.Sequence
			_ = _625_v87
			_625_v87 = _dafny.UnicodeSeqOfUtf8Bytes("jtb")
			var _626_v88 _dafny.MultiSet
			_ = _626_v88
			_626_v88 = _dafny.MultiSetOf(_623_v85, Companion_Default___.Fm18(_496_v1, (func() bool {
				if (_624_v86).Contains(_495_v0) {
					return (_624_v86).Get(_495_v0).(bool)
				}
				return true
			})(), globalState), _dafny.Companion_Sequence_.Update(_623_v85, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((Companion_Default___.Fm18(_dafny.IntOfUint32((_625_v87).Cardinality()), _495_v0, globalState)).Cardinality()), _dafny.IntOfUint32((_623_v85).Cardinality()))).Uint32(), false), _623_v85, _623_v85)
			var _627_v89 D5
			_ = _627_v89
			_627_v89 = Companion_D5_.Create_DC15_((func() _dafny.Array {
				if _495_v0 {
					return (_this).F14()
				}
				return (_this).F14()
			})(), _495_v0, _622_v84, (_626_v88).Cardinality())
			_627_v89 = _627_v89
			var _628_v90 T2
			_ = _628_v90
			var _nw96 *C1 = New_C1_()
			_ = _nw96
			_nw96.Ctor__(_496_v1, true, _495_v0, true, _554_v40, _625_v87)
			_628_v90 = _nw96
			var _629_v91 _dafny.Map
			_ = _629_v91
			_629_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_628_v90, (_628_v90).F7())
			var _630_v92 _dafny.MultiSet
			_ = _630_v92
			_630_v92 = _dafny.MultiSetOf((_628_v90).F7())
			var _631_v93 _dafny.Array
			_ = _631_v93
			var _nwElement0_15 bool = (func() bool {
				if (_629_v91).Contains(_628_v90) {
					return (_629_v91).Get(_628_v90).(bool)
				}
				return true
			})()
			_ = _nwElement0_15
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(20))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_15, 0)
			(_nw97).ArraySet1(_495_v0, 1)
			(_nw97).ArraySet1((_628_v90).F6(), 2)
			(_nw97).ArraySet1(!(Companion_Default___.Fm7(_554_v40, _496_v1, !((_628_v90).F9()), globalState)), 3)
			(_nw97).ArraySet1(true, 4)
			(_nw97).ArraySet1((_628_v90).F7(), 5)
			(_nw97).ArraySet1((_628_v90).F9(), 6)
			(_nw97).ArraySet1((_628_v90).F9(), 7)
			(_nw97).ArraySet1(true, 8)
			(_nw97).ArraySet1((_628_v90).F6(), 9)
			(_nw97).ArraySet1(!_dafny.Companion_Sequence_.Equal((_628_v90).F5(), _625_v87), 10)
			(_nw97).ArraySet1(!((_628_v90).F7()), 11)
			(_nw97).ArraySet1((_dafny.MultiSetOf((_628_v90).F9())).IsProperSubsetOf(_630_v92), 12)
			(_nw97).ArraySet1(((_628_v90).F9()) == ((_628_v90).F9()), 13)
			(_nw97).ArraySet1(_495_v0, 14)
			(_nw97).ArraySet1((_628_v90).F6(), 15)
			(_nw97).ArraySet1((_628_v90).F7(), 16)
			(_nw97).ArraySet1(_495_v0, 17)
			(_nw97).ArraySet1(false, 18)
			(_nw97).ArraySet1(!((_628_v90).F9()) || ((_628_v90).F7()), 19)
			_631_v93 = _nw97
			_631_v93 = (_this).F14()
			_625_v87 = _625_v87
			var _632_v94 _dafny.Map
			_ = _632_v94
			_632_v94 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_628_v90).F8(), (_624_v86).Cardinality())
			var _633_v95 _dafny.Set
			_ = _633_v95
			_633_v95 = _dafny.SetOf((_628_v90).F9(), (_628_v90).F7(), (_628_v90).F7())
			var _634_v96 _dafny.Sequence
			_ = _634_v96
			_634_v96 = _dafny.SeqOf((func() _dafny.Int {
				if (_632_v94).Contains(_496_v1) {
					return (_632_v94).Get(_496_v1).(_dafny.Int)
				}
				return (_628_v90).F8()
			})(), Companion_Default___.SafeModuloInt((_628_v90).F8(), (_628_v90).F8()), Companion_Default___.SafeModuloInt((_628_v90).F8(), (_633_v95).Cardinality()), (_dafny.Zero).Minus(_496_v1), (_628_v90).F8())
			_634_v96 = _dafny.SeqOf(Companion_Default___.SafeModuloInt(_496_v1, _dafny.IntOfInt64(977)))
			if ((_628_v90).F8()).Cmp((_628_v90).Fm1(_495_v0, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_628_v90).F7(), (_628_v90).F6()), _495_v0, globalState)) == 0 {
				r1 = (_628_v90).F9()
				var _635_v97 _dafny.Sequence
				_ = _635_v97
				_635_v97 = _dafny.SeqOf((_628_v90).F5(), _dafny.UnicodeSeqOfUtf8Bytes("nravm"), _dafny.UnicodeSeqOfUtf8Bytes("uqwbaqax"))
				var _636_v98 *C0
				_ = _636_v98
				var _nw98 *C0 = New_C0_()
				_ = _nw98
				_nw98.Ctor__(Companion_Default___.Fm7(_628_v90.F4(), (_628_v90).F8(), false, globalState), true, _628_v90.F4(), (_635_v97).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.IntOfUint32((_635_v97).Cardinality()))).Uint32()).(_dafny.Sequence))
				_636_v98 = _nw98
				var _637_v99 D1
				_ = _637_v99
				_637_v99 = Companion_D1_.Create_DC4_(_dafny.CodePoint('n'), (_628_v90).F8(), _554_v40, (_dafny.Zero).Minus((_628_v90).F8()), !((_628_v90).F9()) || (true))
				var _638_v100 _dafny.Array
				_ = _638_v100
				var _nw99 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(17))
				_ = _nw99
				_638_v100 = _nw99
				var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_638_v100), 0))
				_ = _index66
				(_638_v100).ArraySet1((_628_v90).F5(), (_index66).Int())
				var _639_v101 _dafny.Map
				_ = _639_v101
				_639_v101 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_628_v90).F8(), (_628_v90).F6())
				var _640_v102 _dafny.Set
				_ = _640_v102
				_640_v102 = _dafny.SetOf(_639_v101)
				var _pat_let_tv11 = _496_v1
				_ = _pat_let_tv11
				var _pat_let_tv12 = _628_v90
				_ = _pat_let_tv12
				var _pat_let_tv13 = _628_v90
				_ = _pat_let_tv13
				var _pat_let_tv14 = _628_v90
				_ = _pat_let_tv14
				var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_638_v100), 0))
				_ = _index67
				var _rhs67 bool = !(Companion_Default___.Fm7(_554_v40, (_dafny.SetOf((_628_v90).F8(), (_628_v90).F8(), (_640_v102).Cardinality())).Cardinality(), _495_v0, globalState)) || ((_628_v90).F7())
				_ = _rhs67
				var _rhs68 _dafny.Int = _496_v1
				_ = _rhs68
				var _rhs69 D1 = func(_pat_let17_0 D1) D1 {
					return func(_641_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let18_0 _dafny.Int) D1 {
							return func(_642_dt__update_hcf9_h0 _dafny.Int) D1 {
								return func(_pat_let19_0 _dafny.CodePoint) D1 {
									return func(_643_dt__update_hcf8_h0 _dafny.CodePoint) D1 {
										return func(_pat_let20_0 _dafny.CodePoint) D1 {
											return func(_644_dt__update_hcf10_h0 _dafny.CodePoint) D1 {
												return func(_pat_let21_0 _dafny.Int) D1 {
													return func(_645_dt__update_hcf11_h0 _dafny.Int) D1 {
														return Companion_D1_.Create_DC4_(_643_dt__update_hcf8_h0, _642_dt__update_hcf9_h0, _644_dt__update_hcf10_h0, _645_dt__update_hcf11_h0, (_641_dt__update__tmp_h2).Dtor_cf12())
													}(_pat_let21_0)
												}((_pat_let_tv14).F8())
											}(_pat_let20_0)
										}(_pat_let_tv13.F4())
									}(_pat_let19_0)
								}(_pat_let_tv12.F4())
							}(_pat_let18_0)
						}(_pat_let_tv11)
					}(_pat_let17_0)
				}(p0)
				_ = _rhs69
				var _rhs70 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_625_v87, _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_625_v87, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_628_v90).F8())), _dafny.IntOfUint32((_625_v87).Cardinality()))).Uint32(), _554_v40), _dafny.UnicodeSeqOfUtf8Bytes("agc")))
				_ = _rhs70
				var _rhs71 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_623_v85, _dafny.SeqOf(true, (_628_v90).F9()))).Cardinality()))
				_ = _rhs71
				var _lhs36 _dafny.Array = _638_v100
				_ = _lhs36
				var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(343), _dafny.ArrayLen((_638_v100), 0))
				_ = _lhs37
				var _lhs38 *GlobalState = globalState
				_ = _lhs38
				_495_v0 = _rhs67
				_496_v1 = _rhs68
				_637_v99 = _rhs69
				(_lhs36).ArraySet1(_rhs70, (_lhs37).Int())
				_lhs38.F1 = _rhs71
				var _646_v103 *C0
				_ = _646_v103
				var _nw100 *C0 = New_C0_()
				_ = _nw100
				_nw100.Ctor__((_628_v90).F9(), (_628_v90).F9(), _628_v90.F4(), (_628_v90).F5())
				_646_v103 = _nw100
				(globalState).F1 = (_628_v90).F8()
			} else {
				var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_631_v93), 0))
				_ = _index68
				(_631_v93).ArraySet1(!(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_628_v90).F7(), (_628_v90).F8())).Contains(true), (_index68).Int())
				var _647_v104 _dafny.Array
				_ = _647_v104
				var _nw101 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(11))
				_ = _nw101
				_647_v104 = _nw101
				var _648_v105 *C1
				_ = _648_v105
				var _nw102 *C1 = New_C1_()
				_ = _nw102
				_nw102.Ctor__((_628_v90).F8(), (_628_v90).F6(), (_628_v90).F7(), (_628_v90).F9(), _628_v90.F4(), (_628_v90).F5())
				_648_v105 = _nw102
				var _649_v106 _dafny.Array
				_ = _649_v106
				var _nwElement0_16 *C1 = _648_v105
				_ = _nwElement0_16
				var _nw103 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(7))
				_ = _nw103
				(_nw103).ArraySet1(_nwElement0_16, 0)
				(_nw103).ArraySet1(_648_v105, 1)
				(_nw103).ArraySet1(_648_v105, 2)
				(_nw103).ArraySet1(_648_v105, 3)
				(_nw103).ArraySet1(_648_v105, 4)
				(_nw103).ArraySet1(_648_v105, 5)
				(_nw103).ArraySet1(_648_v105, 6)
				_649_v106 = _nw103
				var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_647_v104), 0))
				_ = _index69
				(_647_v104).ArraySet1(_649_v106, (_index69).Int())
				var _650_v107 _dafny.MultiSet
				_ = _650_v107
				_650_v107 = _dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("sd"), (_628_v90).F5())
				var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_631_v93), 0))
				_ = _index70
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_647_v104), 0))
				_ = _index71
				var _rhs72 bool = !(!((_650_v107).IsDisjointFrom(_650_v107)))
				_ = _rhs72
				var _rhs73 bool = (_628_v90).F9()
				_ = _rhs73
				var _rhs74 _dafny.Array = _649_v106
				_ = _rhs74
				var _lhs39 _dafny.Array = _631_v93
				_ = _lhs39
				var _lhs40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(634), _dafny.ArrayLen((_631_v93), 0))
				_ = _lhs40
				var _lhs41 _dafny.Array = _647_v104
				_ = _lhs41
				var _lhs42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(842), _dafny.ArrayLen((_647_v104), 0))
				_ = _lhs42
				(_lhs39).ArraySet1(_rhs72, (_lhs40).Int())
				_495_v0 = _rhs73
				(_lhs41).ArraySet1(_rhs74, (_lhs42).Int())
				r1 = !((_628_v90).F6())
				_495_v0 = !(!((func() bool {
					if (_628_v90).F7() {
						return true
					}
					return (_628_v90).F6()
				})()) || ((_628_v90).F7()))
				var _651_v108 _dafny.Map
				_ = _651_v108
				_651_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_633_v95).Union(_633_v95), _628_v90.F4())
				_651_v108 = _651_v108
				r2 = (_628_v90).F8()
			}
		}
		var _652_i6 _dafny.Int
		_ = _652_i6
		_652_i6 = _dafny.Zero
		{
			for !(_495_v0) {
				{
					if (_652_i6).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L3
					}
					_652_i6 = (_652_i6).Plus(_dafny.One)
					_495_v0 = !(_495_v0)
					var _653_v109 _dafny.Map
					_ = _653_v109
					_653_v109 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_496_v1, _495_v0)
					var _654_v110 _dafny.Sequence
					_ = _654_v110
					_654_v110 = _dafny.UnicodeSeqOfUtf8Bytes("dvhbgf")
					var _655_v111 *C0
					_ = _655_v111
					var _nw104 *C0 = New_C0_()
					_ = _nw104
					_nw104.Ctor__(_495_v0, !((func() bool {
						if (_653_v109).Contains(_496_v1) {
							return (_653_v109).Get(_496_v1).(bool)
						}
						return !(_495_v0)
					})()) || (_495_v0), Companion_Default___.Fm13(globalState), _654_v110)
					_655_v111 = _nw104
					var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index72
					((_this).F14()).ArraySet1(!(_495_v0), (_index72).Int())
					var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen(((_this).F14()), 0))
					_ = _index73
					((_this).F14()).ArraySet1(_495_v0, (_index73).Int())
					var _656_v112 _dafny.Map
					_ = _656_v112
					_656_v112 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool))
					var _657_v113 *C1
					_ = _657_v113
					var _nw105 *C1 = New_C1_()
					_ = _nw105
					_nw105.Ctor__(_dafny.IntOfInt64(873), (func() bool {
						if ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool) {
							return !((func() bool {
								if (_656_v112).Contains(_495_v0) {
									return (_656_v112).Get(_495_v0).(bool)
								}
								return _495_v0
							})())
						}
						return _495_v0
					})(), _495_v0, ((_this).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(975), _dafny.ArrayLen(((_this).F14()), 0))).Int()).(bool), _554_v40, _654_v110)
					_657_v113 = _nw105
					goto C3
				}
			C3:
			}
			goto L3
		}
	L3:
		r0 = _496_v1
		r1 = _495_v0
		r2 = _496_v1
		return r0, r1, r2
	}
}
func (_this *C2) F14() _dafny.Array {
	{
		return _this._f14
	}
}

// End of class C2

// Definition of class C3
type C3 struct {
	_f4  _dafny.CodePoint
	_f8  _dafny.Int
	_f9  bool
	_f6  bool
	_f7  bool
	_f5  _dafny.Sequence
	F12  bool
	_f13 bool
}

func New_C3_() *C3 {
	_this := C3{}

	_this._f4 = _dafny.CodePoint('D')
	_this._f8 = _dafny.Zero
	_this._f9 = false
	_this._f6 = false
	_this._f7 = false
	_this._f5 = _dafny.EmptySeq
	_this.F12 = false
	_this._f13 = false
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T1_.TraitID_, Companion_T0_.TraitID_}
}

var _ T2 = &C3{}
var _ T1 = &C3{}
var _ T0 = &C3{}
var _ _dafny.TraitOffspring = &C3{}

func (_this *C3) F4() _dafny.CodePoint {
	return _this._f4
}
func (_this *C3) F4_set_(value _dafny.CodePoint) {
	_this._f4 = value
}
func (_this *C3) F8() _dafny.Int {
	return _this._f8
}
func (_this *C3) F9() bool {
	return _this._f9
}
func (_this *C3) F6() bool {
	return _this._f6
}
func (_this *C3) F7() bool {
	return _this._f7
}
func (_this *C3) F5() _dafny.Sequence {
	return _this._f5
}
func (_this *C3) Ctor__(f12 bool, f13 bool, f8 _dafny.Int, f9 bool, f6 bool, f7 bool, f4 _dafny.CodePoint, f5 _dafny.Sequence) {
	{
		(_this).F12 = f12
		(_this)._f13 = f13
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C3) Fm1(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C3) Fm2(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(559))).Uint32(), func(coer42 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg42 _dafny.Int) interface{} {
				return coer42(arg42)
			}
		}(func(_658_i0 _dafny.Int) _dafny.Int {
			return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(329))).Uint32(), func(coer43 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg43 _dafny.Int) interface{} {
					return coer43(arg43)
				}
			}(func(_659_i1 _dafny.Int) _dafny.Int {
				return (_this).F8()
			}))).Cardinality())
		})), (Companion_D0_.Create_DC2_((_this).F8(), !(true), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer44 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg44 _dafny.Int) interface{} {
				return coer44(arg44)
			}
		}(func(_660_i2 _dafny.Int) _dafny.Int {
			return (_this).F8()
		})))).Dtor_cf6())
	}
}
func (_this *C3) Fm5(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.MultiSet, p3 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C3) M1(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _661_v0 _dafny.Map
		_ = _661_v0
		_661_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F12)
		var _662_v1 *C0
		_ = _662_v1
		var _nw106 *C0 = New_C0_()
		_ = _nw106
		_nw106.Ctor__(!(_661_v0).Contains(_this.F12), (_this).F7(), _this.F4(), (_this).F5())
		_662_v1 = _nw106
		var _663_v2 _dafny.Array
		_ = _663_v2
		var _nw107 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw107
		_663_v2 = _nw107
		var _664_v3 _dafny.Sequence
		_ = _664_v3
		_664_v3 = _dafny.SeqOf((_this).F9())
		var _665_v4 _dafny.Sequence
		_ = _665_v4
		_665_v4 = _dafny.SeqOf(_664_v3)
		var _666_v5 _dafny.Array
		_ = _666_v5
		var _nwElement0_17 _dafny.Sequence = _664_v3
		_ = _nwElement0_17
		var _nw108 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(16))
		_ = _nw108
		(_nw108).ArraySet1(_nwElement0_17, 0)
		(_nw108).ArraySet1(_dafny.SeqOf(p0), 1)
		(_nw108).ArraySet1(_dafny.SeqOf(false, p3, (_this).F9()), 2)
		(_nw108).ArraySet1(_664_v3, 3)
		(_nw108).ArraySet1(_664_v3, 4)
		(_nw108).ArraySet1((_665_v4).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_665_v4).Cardinality()))).Uint32()).(_dafny.Sequence), 5)
		(_nw108).ArraySet1(_dafny.SeqOf((_this).F7()), 6)
		(_nw108).ArraySet1(_664_v3, 7)
		(_nw108).ArraySet1(_664_v3, 8)
		(_nw108).ArraySet1(_664_v3, 9)
		(_nw108).ArraySet1(_664_v3, 10)
		(_nw108).ArraySet1(_664_v3, 11)
		(_nw108).ArraySet1(_664_v3, 12)
		(_nw108).ArraySet1(_664_v3, 13)
		(_nw108).ArraySet1(_664_v3, 14)
		(_nw108).ArraySet1(_664_v3, 15)
		_666_v5 = _nw108
		var _667_v6 D1
		_ = _667_v6
		_667_v6 = Companion_D1_.Create_DC3_(_666_v5)
		var _668_v7 _dafny.Array
		_ = _668_v7
		var _nwElement0_18 D1 = _667_v6
		_ = _nwElement0_18
		var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(25))
		_ = _nw109
		(_nw109).ArraySet1(_nwElement0_18, 0)
		(_nw109).ArraySet1(_667_v6, 1)
		(_nw109).ArraySet1(_667_v6, 2)
		(_nw109).ArraySet1(_667_v6, 3)
		(_nw109).ArraySet1(_667_v6, 4)
		(_nw109).ArraySet1(_667_v6, 5)
		(_nw109).ArraySet1(_667_v6, 6)
		(_nw109).ArraySet1(_667_v6, 7)
		(_nw109).ArraySet1(_667_v6, 8)
		(_nw109).ArraySet1(_667_v6, 9)
		(_nw109).ArraySet1(_667_v6, 10)
		(_nw109).ArraySet1(_667_v6, 11)
		(_nw109).ArraySet1(Companion_D1_.Create_DC3_(_666_v5), 12)
		(_nw109).ArraySet1(_667_v6, 13)
		(_nw109).ArraySet1(Companion_D1_.Create_DC3_(_666_v5), 14)
		(_nw109).ArraySet1(_667_v6, 15)
		(_nw109).ArraySet1(_667_v6, 16)
		(_nw109).ArraySet1(_667_v6, 17)
		(_nw109).ArraySet1(Companion_D1_.Create_DC3_(_666_v5), 18)
		(_nw109).ArraySet1(_667_v6, 19)
		(_nw109).ArraySet1(_667_v6, 20)
		(_nw109).ArraySet1(_667_v6, 21)
		(_nw109).ArraySet1(_667_v6, 22)
		(_nw109).ArraySet1(_667_v6, 23)
		(_nw109).ArraySet1(Companion_D1_.Create_DC3_(_666_v5), 24)
		_668_v7 = _nw109
		var _669_v8 *C2
		_ = _669_v8
		var _nw110 *C2 = New_C2_()
		_ = _nw110
		_nw110.Ctor__(_663_v2, _668_v7)
		_669_v8 = _nw110
		(_this).F4_set_(_this.F4())
		var _670_v9 _dafny.Array
		_ = _670_v9
		var _len0_12 _dafny.Int = _dafny.IntOfInt64(7)
		_ = _len0_12
		var _nw111 _dafny.Array
		_ = _nw111
		if _len0_12.Cmp(_dafny.Zero) == 0 {
			_nw111 = _dafny.NewArray(_len0_12)
		} else {
			var _init12 func(_dafny.Int) _dafny.Int = (func(_671_v3 _dafny.Sequence) func(_dafny.Int) _dafny.Int {
				return func(_672_i0 _dafny.Int) _dafny.Int {
					return (_672_i0).Minus(_dafny.IntOfUint32((_671_v3).Cardinality()))
				}
			})(_664_v3)
			_ = _init12
			var _element0_12 = _init12(_dafny.Zero)
			_ = _element0_12
			_nw111 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
			(_nw111).ArraySet1(_element0_12, 0)
			var _nativeLen0_12 = (_len0_12).Int()
			_ = _nativeLen0_12
			for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
				(_nw111).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
			}
		}
		_670_v9 = _nw111
		var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))
		_ = _index74
		(_670_v9).ArraySet1((_dafny.Zero).Minus((_this).F8()), (_index74).Int())
		var _673_v10 _dafny.MultiSet
		_ = _673_v10
		_673_v10 = _dafny.MultiSetOf((p2).Cardinality())
		var _674_v11 _dafny.Map
		_ = _674_v11
		_674_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _673_v10)
		var _675_v12 _dafny.MultiSet
		_ = _675_v12
		_675_v12 = _dafny.MultiSetOf((func() _dafny.MultiSet {
			if (_674_v11).Contains((_this).F6()) {
				return (_674_v11).Get((_this).F6()).(_dafny.MultiSet)
			}
			return (_673_v10).Update((_this).F8(), Companion_Default___.Abs((_this).F8()))
		})(), _673_v10)
		var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))
		_ = _index75
		(_670_v9).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
			if true {
				return ((_this).F8()).Plus((_this).F8())
			}
			return (_675_v12).Cardinality()
		})()), (_index75).Int())
		(_this).F12 = (_this).F9()
		var _hi0 _dafny.Int = (_dafny.MultiSetOf(p1)).Cardinality()
		_ = _hi0
		for _676_i1 := (_670_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))).Int()).(_dafny.Int); _676_i1.Cmp(_hi0) < 0; _676_i1 = _676_i1.Plus(_dafny.One) {
			var _677_v13 *C0
			_ = _677_v13
			var _nw112 *C0 = New_C0_()
			_ = _nw112
			_nw112.Ctor__(p3, (_this).F13(), _this.F4(), (_this).F5())
			_677_v13 = _nw112
			_673_v10 = (_dafny.MultiSetOf(_dafny.IntOfInt64(490), (_670_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))).Int()).(_dafny.Int))).Intersection(_673_v10)
			r0 = _676_i1
			(_this).F12 = p1
		}
		var _678_v14 _dafny.Set
		_ = _678_v14
		_678_v14 = _dafny.SetOf((_670_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))).Int()).(_dafny.Int), (_this).F8())
		r0 = ((_dafny.IntOfInt64(362)).Times((_dafny.Zero).Minus((_678_v14).Cardinality()))).Plus(_dafny.IntOfInt64(-419))
		var _679_v15 _dafny.Map
		_ = _679_v15
		_679_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, (_this).F8())
		r1 = (func() _dafny.Int {
			if !(_dafny.Companion_Sequence_.IsPrefixOf((_this).F5(), _dafny.UnicodeSeqOfUtf8Bytes("n"))) {
				return _dafny.IntOfUint32(((_this).F5()).Cardinality())
			}
			return Companion_Default___.SafeModuloInt((_670_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))).Int()).(_dafny.Int), (func() _dafny.Int {
				if (_679_v15).Contains(p0) {
					return (_679_v15).Get(p0).(_dafny.Int)
				}
				return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfInt64(704))).Cardinality()
			})())
		})()
		var _680_v16 _dafny.Sequence
		_ = _680_v16
		_680_v16 = _dafny.SeqOf(_670_v9, _670_v9)
		var _681_v17 _dafny.Array
		_ = _681_v17
		var _nwElement0_19 _dafny.Array = _670_v9
		_ = _nwElement0_19
		var _nw113 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(20))
		_ = _nw113
		(_nw113).ArraySet1(_nwElement0_19, 0)
		(_nw113).ArraySet1(_670_v9, 1)
		(_nw113).ArraySet1((_680_v16).Select((Companion_Default___.SafeIndex((_670_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_680_v16).Cardinality()))).Uint32()).(_dafny.Array), 2)
		(_nw113).ArraySet1((_680_v16).Select((Companion_Default___.SafeIndex((_670_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(912), _dafny.ArrayLen((_670_v9), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_680_v16).Cardinality()))).Uint32()).(_dafny.Array), 3)
		(_nw113).ArraySet1(_670_v9, 4)
		(_nw113).ArraySet1(_670_v9, 5)
		(_nw113).ArraySet1(_670_v9, 6)
		(_nw113).ArraySet1(_670_v9, 7)
		(_nw113).ArraySet1(_670_v9, 8)
		(_nw113).ArraySet1(_670_v9, 9)
		(_nw113).ArraySet1(_670_v9, 10)
		(_nw113).ArraySet1(_670_v9, 11)
		(_nw113).ArraySet1(_670_v9, 12)
		(_nw113).ArraySet1(_670_v9, 13)
		(_nw113).ArraySet1(_670_v9, 14)
		(_nw113).ArraySet1(_670_v9, 15)
		(_nw113).ArraySet1((func() _dafny.Array {
			if (_this).F9() {
				return _670_v9
			}
			return _670_v9
		})(), 16)
		(_nw113).ArraySet1(_670_v9, 17)
		(_nw113).ArraySet1(_670_v9, 18)
		(_nw113).ArraySet1(_670_v9, 19)
		_681_v17 = _nw113
		r2 = _681_v17
		return r0, r1, r2
	}
}
func (_this *C3) M5(p0 bool, p1 _dafny.Map, p2 _dafny.Array, p3 _dafny.CodePoint, globalState *GlobalState) _dafny.Array {
	{
		var r0 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r0
		var _682_v0 _dafny.Map
		_ = _682_v0
		_682_v0 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), ((_this).F8()).Plus(_dafny.IntOfInt64(678)))
		_682_v0 = (_682_v0).Update((_this).F8(), Companion_Default___.SafeModuloInt((_dafny.Zero).Minus((_this).F8()), (_this).F8()))
		var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
		_ = _index76
		(p2).ArraySet1(p0, (_index76).Int())
		var _683_v1 _dafny.Array
		_ = _683_v1
		var _nw114 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(2))
		_ = _nw114
		_683_v1 = _nw114
		var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
		_ = _index77
		(_683_v1).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_this).F5(), (_this).F5()), (_index77).Int())
		var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
		_ = _index78
		var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
		_ = _index79
		var _rhs75 _dafny.Int = (_this).F8()
		_ = _rhs75
		var _rhs76 bool = p0
		_ = _rhs76
		var _rhs77 _dafny.Sequence = (_this).F5()
		_ = _rhs77
		var _lhs43 *GlobalState = globalState
		_ = _lhs43
		var _lhs44 _dafny.Array = p2
		_ = _lhs44
		var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
		_ = _lhs45
		var _lhs46 _dafny.Array = _683_v1
		_ = _lhs46
		var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
		_ = _lhs47
		_lhs43.F1 = _rhs75
		(_lhs44).ArraySet1(_rhs76, (_lhs45).Int())
		(_lhs46).ArraySet1(_rhs77, (_lhs47).Int())
		var _684_v2 _dafny.Sequence
		_ = _684_v2
		_684_v2 = _dafny.SeqOf((_this).F8(), (_this).F8())
		var _685_v3 _dafny.Sequence
		_ = _685_v3
		_685_v3 = _dafny.SeqOf(!(!(p0)) || ((_this).F6()), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.Companion_Sequence_.Update(_684_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(101), _dafny.IntOfUint32((_684_v2).Cardinality()))).Uint32(), (_this).F8()), _684_v2), (_this).F13())
		var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
		_ = _index80
		(p2).ArraySet1((_685_v3).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_685_v3).Cardinality()))).Uint32()).(bool), (_index80).Int())
		var _hi1 _dafny.Int = ((_this).F8()).Plus((_this).F8())
		_ = _hi1
		for _686_i0 := (_this).F8(); _686_i0.Cmp(_hi1) < 0; _686_i0 = _686_i0.Plus(_dafny.One) {
			_685_v3 = _685_v3
			var _687_v4 _dafny.Array
			_ = _687_v4
			var _nw115 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(24))
			_ = _nw115
			_687_v4 = _nw115
			var _688_v5 D0
			_ = _688_v5
			_688_v5 = Companion_D0_.Create_DC2_((_this).F8(), Companion_Default___.Fm7(p3, _686_i0, false, globalState), _dafny.SeqOf(_686_i0))
			var _689_v6 _dafny.Map
			_ = _689_v6
			_689_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_688_v5, (_this).F8())
			var _690_v7 *C1
			_ = _690_v7
			var _nw116 *C1 = New_C1_()
			_ = _nw116
			_nw116.Ctor__((_689_v6).Cardinality(), false, (_this).F13(), _this.F12, p3, (_this).F5())
			_690_v7 = _nw116
			var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_687_v4), 0))
			_ = _index81
			(_687_v4).ArraySet1(_690_v7, (_index81).Int())
			var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_687_v4), 0))
			_ = _index82
			(_687_v4).ArraySet1(_690_v7, (_index82).Int())
			(globalState).F1 = _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update((_this).F5(), (Companion_Default___.SafeIndex((func() _dafny.Int {
				if true {
					return (_this).F8()
				}
				return _dafny.IntOfUint32(((_683_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))).Int()).(_dafny.Sequence)).Cardinality())
			})(), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32(), p3)).Cardinality())
			var _691_v8 _dafny.Set
			_ = _691_v8
			_691_v8 = _dafny.SetOf(_685_v3)
			var _692_v9 _dafny.Set
			_ = _692_v9
			_692_v9 = _dafny.SetOf((_this).F6(), _this.F12, (_this).F9())
			var _693_v10 D3
			_ = _693_v10
			_693_v10 = Companion_D3_.Create_DC11_((_this).F9(), (_691_v8).Cardinality(), _692_v9)
			if !((_693_v10).Dtor_cf20()) {
				(globalState).F1 = _686_i0
				(globalState).F1 = _dafny.IntOfInt64(275)
				(globalState).F1 = (_686_i0).Minus(_dafny.IntOfUint32((_684_v2).Cardinality()))
				var _694_v11 _dafny.Array
				_ = _694_v11
				var _nw117 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
				_ = _nw117
				_694_v11 = _nw117
				var _695_v12 _dafny.Array
				_ = _695_v12
				var _nw118 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(20))
				_ = _nw118
				_695_v12 = _nw118
				var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_694_v11), 0))
				_ = _index83
				(_694_v11).ArraySet1(_695_v12, (_index83).Int())
				var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(601), _dafny.ArrayLen((_694_v11), 0))
				_ = _index84
				(_694_v11).ArraySet1(_695_v12, (_index84).Int())
				(globalState).F1 = (func() _dafny.Int {
					if (_686_i0).Cmp(_686_i0) == 0 {
						return (_this).F8()
					}
					return (_this).F8()
				})()
			} else {
				var _696_v13 _dafny.Map
				_ = _696_v13
				_696_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F8()).Minus((_this).F8()), (_this).F5())
				var _697_v14 _dafny.Array
				_ = _697_v14
				var _nw119 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(23))
				_ = _nw119
				_697_v14 = _nw119
				var _698_v15 _dafny.MultiSet
				_ = _698_v15
				_698_v15 = _dafny.MultiSetOf(_697_v14, _697_v14, _697_v14)
				(globalState).F1 = _dafny.IntOfUint32(((func() _dafny.Sequence {
					if (_696_v13).Contains((func() _dafny.Int {
						if (_698_v15).Contains(_697_v14) {
							return (_698_v15).Multiplicity(_697_v14)
						}
						return (_this).F8()
					})()) {
						return (_696_v13).Get((func() _dafny.Int {
							if (_698_v15).Contains(_697_v14) {
								return (_698_v15).Multiplicity(_697_v14)
							}
							return (_this).F8()
						})()).(_dafny.Sequence)
					}
					return _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(291))).Uint32(), func(coer45 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg45 _dafny.Int) interface{} {
							return coer45(arg45)
						}
					}(func(_699_i1 _dafny.Int) _dafny.CodePoint {
						return _this.F4()
					}))
				})()).Cardinality())
				var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_697_v14), 0))
				_ = _index85
				(_697_v14).ArraySet1(Companion_Default___.SafeDivisionInt((_693_v10).Dtor_cf21(), Companion_Default___.Fm8(_dafny.IntOfUint32(((_this).F5()).Cardinality()), (_this).F9(), true, globalState)), (_index85).Int())
				var _700_v16 _dafny.MultiSet
				_ = _700_v16
				_700_v16 = _dafny.MultiSetOf((_this).F13())
				var _701_v17 D5
				_ = _701_v17
				_701_v17 = Companion_D5_.Create_DC15_(p2, p0, Companion_Default___.Fm24(_686_i0, true, p0, (_700_v16).Cardinality(), globalState), _dafny.IntOfInt64(69))
				var _702_v18 _dafny.Map
				_ = _702_v18
				_702_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_701_v17).Dtor_cf29())
				var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_697_v14), 0))
				_ = _index86
				var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
				_ = _index87
				var _rhs78 _dafny.Int = (_this).F8()
				_ = _rhs78
				var _rhs79 _dafny.Int = (_this).F8()
				_ = _rhs79
				var _rhs80 _dafny.Map = (_702_v18).Merge((_702_v18).Merge(_702_v18))
				_ = _rhs80
				var _rhs81 _dafny.Int = (_dafny.Zero).Minus(((_684_v2).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_683_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))).Int()).(_dafny.Sequence)).Cardinality()), _dafny.IntOfUint32((_684_v2).Cardinality()))).Uint32()).(_dafny.Int)).Plus(_686_i0))
				_ = _rhs81
				var _rhs82 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate((_683_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))).Int()).(_dafny.Sequence), (_this).F5()), (Companion_Default___.SafeIndex(_686_i0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((_683_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))).Int()).(_dafny.Sequence), (_this).F5())).Cardinality()))).Uint32(), p3)
				_ = _rhs82
				var _lhs48 *GlobalState = globalState
				_ = _lhs48
				var _lhs49 _dafny.Array = _697_v14
				_ = _lhs49
				var _lhs50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_697_v14), 0))
				_ = _lhs50
				var _lhs51 *GlobalState = globalState
				_ = _lhs51
				var _lhs52 _dafny.Array = _683_v1
				_ = _lhs52
				var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
				_ = _lhs53
				_lhs48.F1 = _rhs78
				(_lhs49).ArraySet1(_rhs79, (_lhs50).Int())
				_702_v18 = _rhs80
				_lhs51.F1 = _rhs81
				(_lhs52).ArraySet1(_rhs82, (_lhs53).Int())
				var _703_v19 _dafny.Map
				_ = _703_v19
				_703_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_686_i0, p3)
				_703_v19 = (_703_v19).Update(_dafny.IntOfInt64(-494), (func() _dafny.CodePoint {
					if p0 {
						return _this.F4()
					}
					return p3
				})())
				var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
				_ = _index88
				(p2).ArraySet1(false, (_index88).Int())
				(globalState).F1 = ((_dafny.Zero).Minus((_this).F8())).Times((_697_v14).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(220), _dafny.ArrayLen((_697_v14), 0))).Int()).(_dafny.Int))
			}
		}
		var _704_v20 _dafny.Array
		_ = _704_v20
		var _nw120 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw120
		_704_v20 = _nw120
		var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_704_v20), 0))
		_ = _index89
		(_704_v20).ArraySet1((_this).F8(), (_index89).Int())
		var _705_v21 _dafny.Array
		_ = _705_v21
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(24)
		_ = _len0_13
		var _nw121 _dafny.Array
		_ = _nw121
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw121 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) _dafny.MultiSet = func(_706_i2 _dafny.Int) _dafny.MultiSet {
				return _dafny.MultiSetOf(true, _this.F12, true, false, !((_this).F9()))
			}
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw121 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw121).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw121).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_705_v21 = _nw121
		var _707_v22 _dafny.MultiSet
		_ = _707_v22
		_707_v22 = _dafny.MultiSetOf((_this).F8())
		var _708_v23 D0
		_ = _708_v23
		_708_v23 = Companion_D0_.Create_DC1_((_this).F9(), (_707_v22).Cardinality(), (_this).F7(), _dafny.IntOfInt64(-704))
		var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
		_ = _index90
		var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
		_ = _index91
		var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_704_v20), 0))
		_ = _index92
		var _rhs83 _dafny.Sequence = (func() _dafny.Sequence {
			if ((_708_v23).Dtor_cf1()).Cmp((_this).F8()) == 0 {
				return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("a"), _dafny.UnicodeSeqOfUtf8Bytes("dbviiq"))
			}
			return (_this).F5()
		})()
		_ = _rhs83
		var _rhs84 bool = (_this).F13()
		_ = _rhs84
		var _rhs85 _dafny.Int = Companion_Default___.SafeModuloInt((_this).F8(), Companion_Default___.SafeModuloInt((_this).F8(), (_this).F8()))
		_ = _rhs85
		var _rhs86 _dafny.Array = _705_v21
		_ = _rhs86
		var _lhs54 _dafny.Array = _683_v1
		_ = _lhs54
		var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))
		_ = _lhs55
		var _lhs56 _dafny.Array = p2
		_ = _lhs56
		var _lhs57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))
		_ = _lhs57
		var _lhs58 _dafny.Array = _704_v20
		_ = _lhs58
		var _lhs59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(470), _dafny.ArrayLen((_704_v20), 0))
		_ = _lhs59
		(_lhs54).ArraySet1(_rhs83, (_lhs55).Int())
		(_lhs56).ArraySet1(_rhs84, (_lhs57).Int())
		(_lhs58).ArraySet1(_rhs85, (_lhs59).Int())
		_705_v21 = _rhs86
		var _709_v24 T2
		_ = _709_v24
		var _nw122 *C1 = New_C1_()
		_ = _nw122
		_nw122.Ctor__((func() _dafny.Int {
			if (_682_v0).Contains((_this).F8()) {
				return (_682_v0).Get((_this).F8()).(_dafny.Int)
			}
			return _dafny.IntOfInt64(444)
		})(), (_this).F13(), true, (_this).F6(), _this.F4(), (_683_v1).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(685), _dafny.ArrayLen((_683_v1), 0))).Int()).(_dafny.Sequence))
		_709_v24 = _nw122
		var _710_v25 _dafny.Map
		_ = _710_v25
		_710_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))).Int()).(bool), _709_v24)
		var _711_v26 _dafny.Array
		_ = _711_v26
		var _nwElement0_20 _dafny.Map = _710_v25
		_ = _nwElement0_20
		var _nw123 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(14))
		_ = _nw123
		(_nw123).ArraySet1(_nwElement0_20, 0)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F12, _709_v24), 1)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_709_v24).F6(), _709_v24), 2)
		(_nw123).ArraySet1(_710_v25, 3)
		(_nw123).ArraySet1(_710_v25, 4)
		(_nw123).ArraySet1(_710_v25, 5)
		(_nw123).ArraySet1(_710_v25, 6)
		(_nw123).ArraySet1(_710_v25, 7)
		(_nw123).ArraySet1(_710_v25, 8)
		(_nw123).ArraySet1(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_709_v24).F7(), _709_v24), 9)
		(_nw123).ArraySet1((_710_v25).Update((_this).F9(), _709_v24), 10)
		(_nw123).ArraySet1(_710_v25, 11)
		(_nw123).ArraySet1(_710_v25, 12)
		(_nw123).ArraySet1((_710_v25).Update((p2).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(570), _dafny.ArrayLen((p2), 0))).Int()).(bool), _709_v24), 13)
		_711_v26 = _nw123
		var _712_v27 _dafny.Map
		_ = _712_v27
		_712_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_711_v26, (_685_v3).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_685_v3).Cardinality()))).Uint32()).(bool))
		(_this).F12 = (func() bool {
			if (_712_v27).Contains(_711_v26) {
				return (_712_v27).Get(_711_v26).(bool)
			}
			return !(p0)
		})()
		r0 = _704_v20
		return r0
	}
}
func (_this *C3) F13() bool {
	{
		return _this._f13
	}
}

// End of class C3

// Definition of class C4
type C4 struct {
	_f4 _dafny.CodePoint
	_f8 _dafny.Int
	_f9 bool
	_f6 bool
	_f7 bool
	_f5 _dafny.Sequence
}

func New_C4_() *C4 {
	_this := C4{}

	_this._f4 = _dafny.CodePoint('D')
	_this._f8 = _dafny.Zero
	_this._f9 = false
	_this._f6 = false
	_this._f7 = false
	_this._f5 = _dafny.EmptySeq
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
	return [](*_dafny.TraitID){Companion_T2_.TraitID_, Companion_T0_.TraitID_, Companion_T3_.TraitID_, Companion_T1_.TraitID_}
}

var _ T2 = &C4{}
var _ T0 = &C4{}
var _ T3 = &C4{}
var _ T1 = &C4{}
var _ _dafny.TraitOffspring = &C4{}

func (_this *C4) F4() _dafny.CodePoint {
	return _this._f4
}
func (_this *C4) F4_set_(value _dafny.CodePoint) {
	_this._f4 = value
}
func (_this *C4) F8() _dafny.Int {
	return _this._f8
}
func (_this *C4) F9() bool {
	return _this._f9
}
func (_this *C4) F6() bool {
	return _this._f6
}
func (_this *C4) F7() bool {
	return _this._f7
}
func (_this *C4) F5() _dafny.Sequence {
	return _this._f5
}
func (_this *C4) Ctor__(f8 _dafny.Int, f9 bool, f6 bool, f7 bool, f4 _dafny.CodePoint, f5 _dafny.Sequence) {
	{
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f4 = f4
		(_this)._f5 = f5
	}
}
func (_this *C4) Fm1(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_this).F8()
	}
}
func (_this *C4) Fm2(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_dafny.MultiSetOf((_this).F7())).Cardinality()), (func() _dafny.Sequence {
			if !(!((_this).F9())) {
				return _dafny.SeqOf((_this).F8(), (_this).F8(), (_this).F8(), (_this).F8(), (_this).F8())
			}
			return _dafny.SeqOf(_dafny.IntOfUint32(((_this).F5()).Cardinality()))
		})())
	}
}
func (_this *C4) M1(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _713_v0 _dafny.Sequence
		_ = _713_v0
		_713_v0 = _dafny.SeqOf((_this).F8(), (_this).F8())
		var _714_v1 _dafny.MultiSet
		_ = _714_v1
		_714_v1 = _dafny.MultiSetOf(_dafny.IntOfUint32(((_this).F5()).Cardinality()))
		var _715_v2 _dafny.Sequence
		_ = _715_v2
		_715_v2 = _dafny.SeqOf(p3, (_dafny.MultiSetFromSeq(_713_v0)).Equals(_714_v1))
		_715_v2 = _715_v2
		var _716_v3 *C1
		_ = _716_v3
		var _nw124 *C1 = New_C1_()
		_ = _nw124
		_nw124.Ctor__((_this).F8(), false, p1, (_this).F6(), _this.F4(), (_this).F5())
		_716_v3 = _nw124
		var _717_v4 _dafny.Map
		_ = _717_v4
		_717_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_716_v3, (_this).F6())
		var _718_v5 _dafny.Map
		_ = _718_v5
		_718_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p3, (_this).F8())
		var _719_v6 _dafny.Map
		_ = _719_v6
		_719_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (func() _dafny.Int {
			if (_718_v5).Contains(p0) {
				return (_718_v5).Get(p0).(_dafny.Int)
			}
			return (_this).F8()
		})())
		var _720_v7 _dafny.MultiSet
		_ = _720_v7
		_720_v7 = _dafny.MultiSetOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_717_v4).Cardinality()), _719_v6)
		r0 = (func() _dafny.Int {
			if (_720_v7).Contains((func() _dafny.Map {
				var _coll16 = _dafny.NewMapBuilder()
				_ = _coll16
				for _iter19 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(908), _dafny.IntOfInt64(663))); ; {
					_compr_16, _ok19 := _iter19()
					if !_ok19 {
						break
					}
					var _721_v8 _dafny.Int
					_721_v8 = interface{}(_compr_16).(_dafny.Int)
					if ((_dafny.IntOfInt64(908)).Cmp(_721_v8) <= 0) && ((_721_v8).Cmp(_dafny.IntOfInt64(663)) < 0) {
						_coll16.Add((_721_v8).Plus((_this).F8()), (_this).F8())
					}
				}
				return _coll16.ToMap()
			}()).Merge(_719_v6)) {
				return (_720_v7).Multiplicity((func() _dafny.Map {
					var _coll17 = _dafny.NewMapBuilder()
					_ = _coll17
					for _iter20 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(908), _dafny.IntOfInt64(663))); ; {
						_compr_17, _ok20 := _iter20()
						if !_ok20 {
							break
						}
						var _722_v8 _dafny.Int
						_722_v8 = interface{}(_compr_17).(_dafny.Int)
						if ((_dafny.IntOfInt64(908)).Cmp(_722_v8) <= 0) && ((_722_v8).Cmp(_dafny.IntOfInt64(663)) < 0) {
							_coll17.Add((_722_v8).Plus((_this).F8()), (_this).F8())
						}
					}
					return _coll17.ToMap()
				}()).Merge(_719_v6))
			}
			return (_this).F8()
		})()
		var _hi2 _dafny.Int = (_this).F8()
		_ = _hi2
		for _723_i0 := ((_this).F8()).Plus((_this).F8()); _723_i0.Cmp(_hi2) < 0; _723_i0 = _723_i0.Plus(_dafny.One) {
			r1 = (_dafny.Zero).Minus(((func() _dafny.Int {
				if p0 {
					return (_this).F8()
				}
				return _dafny.IntOfUint32(((_this).F5()).Cardinality())
			})()).Minus((_723_i0).Times(_723_i0)))
			var _724_v9 _dafny.Array
			_ = _724_v9
			var _nw125 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(5))
			_ = _nw125
			_724_v9 = _nw125
			var _725_v10 _dafny.Array
			_ = _725_v10
			var _nw126 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.One)
			_ = _nw126
			_725_v10 = _nw126
			var _726_v11 D5
			_ = _726_v11
			_726_v11 = Companion_D5_.Create_DC15_(_725_v10, true, _718_v5, _dafny.IntOfInt64(-268))
			var _727_v12 _dafny.Array
			_ = _727_v12
			var _nwElement0_21 bool = (_this).F6()
			_ = _nwElement0_21
			var _nw127 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(4))
			_ = _nw127
			(_nw127).ArraySet1(_nwElement0_21, 0)
			(_nw127).ArraySet1(p1, 1)
			(_nw127).ArraySet1((_726_v11).Dtor_cf27(), 2)
			(_nw127).ArraySet1(true, 3)
			_727_v12 = _nw127
			var _728_v13 _dafny.Map
			_ = _728_v13
			_728_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _725_v10)
			var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_724_v9), 0))
			_ = _index93
			(_724_v9).ArraySet1(_728_v13, (_index93).Int())
			var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(900), _dafny.ArrayLen((_724_v9), 0))
			_ = _index94
			(_724_v9).ArraySet1((_728_v13).Update(_dafny.IntOfInt64(-894), _727_v12), (_index94).Int())
			var _729_v14 _dafny.Array
			_ = _729_v14
			var _nw128 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(5))
			_ = _nw128
			_729_v14 = _nw128
			var _730_v15 _dafny.Array
			_ = _730_v15
			var _nwElement0_22 _dafny.Int = (_this).F8()
			_ = _nwElement0_22
			var _nw129 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(6))
			_ = _nw129
			(_nw129).ArraySet1(_nwElement0_22, 0)
			(_nw129).ArraySet1(_723_i0, 1)
			(_nw129).ArraySet1((_this).F8(), 2)
			(_nw129).ArraySet1(_dafny.IntOfInt64(716), 3)
			(_nw129).ArraySet1(_723_i0, 4)
			(_nw129).ArraySet1((_this).F8(), 5)
			_730_v15 = _nw129
			var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_729_v14), 0))
			_ = _index95
			(_729_v14).ArraySet1(_730_v15, (_index95).Int())
			var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_729_v14), 0))
			_ = _index96
			var _rhs87 _dafny.Array = _730_v15
			_ = _rhs87
			var _rhs88 _dafny.Array = _730_v15
			_ = _rhs88
			var _lhs60 _dafny.Array = _729_v14
			_ = _lhs60
			var _lhs61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(244), _dafny.ArrayLen((_729_v14), 0))
			_ = _lhs61
			(_lhs60).ArraySet1(_rhs87, (_lhs61).Int())
			_730_v15 = _rhs88
			(globalState).F1 = (func() _dafny.Int {
				if (_this).F6() {
					return _723_i0
				}
				return _723_i0
			})()
		}
		_714_v1 = ((_714_v1).Intersection((_714_v1).Update(_dafny.IntOfInt64(94), Companion_Default___.Abs((_this).F8())))).Difference(_714_v1)
		(_this).F4_set_(_this.F4())
		r0 = (_this).F8()
		r0 = (_this).F8()
		r1 = (_this).F8()
		var _731_v16 _dafny.Array
		_ = _731_v16
		var _nw130 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(3))
		_ = _nw130
		_731_v16 = _nw130
		r2 = _731_v16
		return r0, r1, r2
	}
}
func (_this *C4) M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _732_v0 _dafny.MultiSet
		_ = _732_v0
		_732_v0 = _dafny.MultiSetOf((_this).F9(), !(true))
		var _733_v1 _dafny.Sequence
		_ = _733_v1
		_733_v1 = _dafny.SeqOf((_this).F7())
		if ((_732_v0).Difference(_dafny.MultiSetOf((_this).F9()))).IsDisjointFrom(_dafny.MultiSetFromSeq(_733_v1)) {
			Companion_Default___.M0(globalState)
			var _734_v2 bool
			_ = _734_v2
			_734_v2 = false
			_734_v2 = (_this).F9()
			var _735_v3 _dafny.Array
			_ = _735_v3
			var _nw131 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(25))
			_ = _nw131
			_735_v3 = _nw131
			var _736_v4 _dafny.Sequence
			_ = _736_v4
			_736_v4 = _dafny.SeqOf((_this).F8())
			var _737_v5 _dafny.Map
			_ = _737_v5
			_737_v5 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), Companion_Default___.Fm27((_this).F8(), _734_v2, _734_v2, globalState))
			var _738_v6 *C1
			_ = _738_v6
			var _nw132 *C1 = New_C1_()
			_ = _nw132
			_nw132.Ctor__(_dafny.IntOfUint32((_736_v4).Cardinality()), (_this).F9(), (func() bool {
				if (_737_v5).Contains(_this.F4()) {
					return (_737_v5).Get(_this.F4()).(bool)
				}
				return (_this).F9()
			})(), !(_734_v2), _dafny.CodePoint('i'), (_this).F5())
			_738_v6 = _nw132
			var _739_v7 _dafny.Map
			_ = _739_v7
			_739_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_738_v6, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _734_v2))
			var _740_v8 _dafny.Map
			_ = _740_v8
			_740_v8 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _734_v2)
			var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_735_v3), 0))
			_ = _index97
			(_735_v3).ArraySet1((_dafny.Zero).Minus(((func() _dafny.Map {
				if (_739_v7).Contains(_738_v6) {
					return (_739_v7).Get(_738_v6).(_dafny.Map)
				}
				return _740_v8
			})()).Cardinality()), (_index97).Int())
			var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_735_v3), 0))
			_ = _index98
			var _rhs89 _dafny.Int = (_dafny.Zero).Minus((_this).F8())
			_ = _rhs89
			var _rhs90 _dafny.Int = (_this).F8()
			_ = _rhs90
			var _rhs91 bool = !((func() bool {
				if ((_dafny.Zero).Minus(p0)).Cmp(p0) < 0 {
					return _734_v2
				}
				return (_this).F9()
			})())
			_ = _rhs91
			var _lhs62 _dafny.Array = _735_v3
			_ = _lhs62
			var _lhs63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_735_v3), 0))
			_ = _lhs63
			var _lhs64 *GlobalState = globalState
			_ = _lhs64
			(_lhs62).ArraySet1(_rhs89, (_lhs63).Int())
			_lhs64.F1 = _rhs90
			_734_v2 = _rhs91
			var _741_v9 *C1
			_ = _741_v9
			var _nw133 *C1 = New_C1_()
			_ = _nw133
			_nw133.Ctor__(Companion_Default___.SafeModuloInt((_735_v3).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(688), _dafny.ArrayLen((_735_v3), 0))).Int()).(_dafny.Int), (_this).F8()), _734_v2, _734_v2, (_this).F7(), _this.F4(), (_this).F5())
			_741_v9 = _nw133
			(_this).F4_set_(_this.F4())
		} else {
			var _742_v10 bool
			_ = _742_v10
			_742_v10 = false
			_742_v10 = Companion_Default___.Fm27(p0, (_this).F9(), ((_this).F7()) || ((_this).F6()), globalState)
			var _743_v11 D1
			_ = _743_v11
			_743_v11 = Companion_D1_.Create_DC5_((_this).F8())
			var _744_v12 _dafny.Sequence
			_ = _744_v12
			_744_v12 = _dafny.SeqOf(_743_v11, _743_v11)
			var _745_v13 _dafny.Sequence
			_ = _745_v13
			_745_v13 = _744_v12
			var _source8 _dafny.Sequence = _745_v13
			_ = _source8
			var _746___mcc_h0 _dafny.Sequence = _source8
			_ = _746___mcc_h0
			var _747_cf24 _dafny.Sequence = _746___mcc_h0
			_ = _747_cf24
			var _748_v14 _dafny.Sequence
			_ = _748_v14
			_748_v14 = _dafny.SeqOf(_dafny.IntOfInt64(272))
			var _749_v15 _dafny.Set
			_ = _749_v15
			_749_v15 = _dafny.SetOf(_748_v14)
			var _750_v16 *C1
			_ = _750_v16
			var _nw134 *C1 = New_C1_()
			_ = _nw134
			_nw134.Ctor__((_dafny.Zero).Minus(p0), !((_this).F6()), (_dafny.IntOfInt64(653)).Cmp((_this).F8()) > 0, Companion_Default___.Fm27((_749_v15).Cardinality(), false, _742_v10, globalState), _this.F4(), _dafny.Companion_Sequence_.Concatenate((_this).F5(), (_this).F5()))
			_750_v16 = _nw134
			var _751_v17 D1
			_ = _751_v17
			_751_v17 = Companion_D1_.Create_DC4_(_this.F4(), (_this).F8(), _this.F4(), p0, (_this).F6())
			_742_v10 = (((_this).F8()).Plus((_751_v17).Dtor_cf11())).Cmp(Companion_Default___.SafeModuloInt((_this).F8(), p0)) < 0
			var _752_v18 _dafny.Sequence
			_ = _752_v18
			_752_v18 = _dafny.UnicodeSeqOfUtf8Bytes("euaygar")
			var _753_v19 _dafny.Array
			_ = _753_v19
			var _nw135 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(5))
			_ = _nw135
			_753_v19 = _nw135
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_753_v19), 0))
			_ = _index99
			(_753_v19).ArraySet1(_dafny.MultiSetOf(_dafny.UnicodeSeqOfUtf8Bytes("pas")), (_index99).Int())
			var _754_v20 _dafny.Array
			_ = _754_v20
			var _nw136 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(4))
			_ = _nw136
			_754_v20 = _nw136
			var _755_v21 _dafny.Array
			_ = _755_v21
			var _nwElement0_23 _dafny.Array = _754_v20
			_ = _nwElement0_23
			var _nw137 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_23, nil, _dafny.IntOfInt64(18))
			_ = _nw137
			(_nw137).ArraySet1(_nwElement0_23, 0)
			(_nw137).ArraySet1(_754_v20, 1)
			(_nw137).ArraySet1(_754_v20, 2)
			(_nw137).ArraySet1(_754_v20, 3)
			(_nw137).ArraySet1(_754_v20, 4)
			(_nw137).ArraySet1(_754_v20, 5)
			(_nw137).ArraySet1(_754_v20, 6)
			(_nw137).ArraySet1(_754_v20, 7)
			(_nw137).ArraySet1(_754_v20, 8)
			(_nw137).ArraySet1(_754_v20, 9)
			(_nw137).ArraySet1(_754_v20, 10)
			(_nw137).ArraySet1(_754_v20, 11)
			(_nw137).ArraySet1(_754_v20, 12)
			(_nw137).ArraySet1(_754_v20, 13)
			(_nw137).ArraySet1(_754_v20, 14)
			(_nw137).ArraySet1(_754_v20, 15)
			(_nw137).ArraySet1((func() _dafny.Array {
				if Companion_Default___.Fm27(p0, false, Companion_Default___.Fm27(p0, (_733_v1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.IntOfUint32((_733_v1).Cardinality()))).Uint32()).(bool), (_this).F6(), globalState), globalState) {
					return _754_v20
				}
				return _754_v20
			})(), 16)
			(_nw137).ArraySet1(_754_v20, 17)
			_755_v21 = _nw137
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_755_v21), 0))
			_ = _index100
			(_755_v21).ArraySet1(_754_v20, (_index100).Int())
			var _756_v22 _dafny.Set
			_ = _756_v22
			_756_v22 = _dafny.SetOf((_this).F7(), (_this).F7())
			var _757_v23 D3
			_ = _757_v23
			_757_v23 = Companion_D3_.Create_DC11_((_this).F7(), p0, _756_v22)
			var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_753_v19), 0))
			_ = _index101
			var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_755_v21), 0))
			_ = _index102
			var _rhs92 bool = (false) == ((_this).F9())
			_ = _rhs92
			var _rhs93 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("wxeov"), Companion_Default___.Fm28(_dafny.IntOfUint32((_748_v14).Cardinality()), _742_v10, _733_v1, Companion_D3_.Create_DC11_(false, (_this).F8(), _756_v22), globalState))
			_ = _rhs93
			var _rhs94 _dafny.MultiSet = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm28((_this).F8(), (_this).F9(), _733_v1, _757_v23, globalState), _dafny.UnicodeSeqOfUtf8Bytes("nservh")))
			_ = _rhs94
			var _rhs95 _dafny.Array = _754_v20
			_ = _rhs95
			var _lhs65 _dafny.Array = _753_v19
			_ = _lhs65
			var _lhs66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(227), _dafny.ArrayLen((_753_v19), 0))
			_ = _lhs66
			var _lhs67 _dafny.Array = _755_v21
			_ = _lhs67
			var _lhs68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.Zero, _dafny.ArrayLen((_755_v21), 0))
			_ = _lhs68
			_742_v10 = _rhs92
			_752_v18 = _rhs93
			(_lhs65).ArraySet1(_rhs94, (_lhs66).Int())
			(_lhs67).ArraySet1(_rhs95, (_lhs68).Int())
			var _758_v24 _dafny.Array
			_ = _758_v24
			var _nw138 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(9))
			_ = _nw138
			_758_v24 = _nw138
			var _759_v25 _dafny.Map
			_ = _759_v25
			_759_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_754_v20, _758_v24)
			var _760_v26 _dafny.MultiSet
			_ = _760_v26
			_760_v26 = _dafny.MultiSetOf((_this).F8())
			var _761_v27 _dafny.Map
			_ = _761_v27
			_761_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Array {
				if (_759_v25).Contains(_754_v20) {
					return (_759_v25).Get(_754_v20).(_dafny.Array)
				}
				return _758_v24
			})(), _760_v26)
			var _762_v28 *C3
			_ = _762_v28
			var _nw139 *C3 = New_C3_()
			_ = _nw139
			_nw139.Ctor__(false, _742_v10, (_dafny.Zero).Minus((_761_v27).Cardinality()), (_this).F7(), (_this).F7(), (_this).F7(), _this.F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(657))).Uint32(), func(coer46 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg46 _dafny.Int) interface{} {
					return coer46(arg46)
				}
			}(func(_763_i0 _dafny.Int) _dafny.CodePoint {
				return _this.F4()
			})))
			_762_v28 = _nw139
			var _764_v29 _dafny.Array
			_ = _764_v29
			var _len0_14 _dafny.Int = _dafny.IntOfInt64(14)
			_ = _len0_14
			var _nw140 _dafny.Array
			_ = _nw140
			if _len0_14.Cmp(_dafny.Zero) == 0 {
				_nw140 = _dafny.NewArray(_len0_14)
			} else {
				var _init14 func(_dafny.Int) bool = func(_765_i1 _dafny.Int) bool {
					return false
				}
				_ = _init14
				var _element0_14 = _init14(_dafny.Zero)
				_ = _element0_14
				_nw140 = _dafny.NewArrayFromExample(_element0_14, nil, _len0_14)
				(_nw140).ArraySet1(_element0_14, 0)
				var _nativeLen0_14 = (_len0_14).Int()
				_ = _nativeLen0_14
				for _i0_14 := 1; _i0_14 < _nativeLen0_14; _i0_14++ {
					(_nw140).ArraySet1(_init14(_dafny.IntOf(_i0_14)), _i0_14)
				}
			}
			_764_v29 = _nw140
			var _766_v30 _dafny.Set
			_ = _766_v30
			_766_v30 = _dafny.SetOf(_764_v29)
			var _767_v31 D7
			_ = _767_v31
			_767_v31 = Companion_D7_.Create_DC18_(_766_v30)
			var _rhs96 bool = (_this).F7()
			_ = _rhs96
			var _rhs97 bool = !(_742_v10) || ((_this).F7())
			_ = _rhs97
			var _rhs98 _dafny.Set = (_767_v31).Dtor_cf33()
			_ = _rhs98
			var _rhs99 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs99
			var _lhs69 *GlobalState = globalState
			_ = _lhs69
			_742_v10 = _rhs96
			_742_v10 = _rhs97
			_766_v30 = _rhs98
			_lhs69.F1 = _rhs99
			var _768_v32 _dafny.Array
			_ = _768_v32
			var _nw141 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
			_ = _nw141
			_768_v32 = _nw141
			var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_768_v32), 0))
			_ = _index103
			(_768_v32).ArraySet1((_dafny.Zero).Minus(p0), (_index103).Int())
			var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_768_v32), 0))
			_ = _index104
			(_768_v32).ArraySet1((_dafny.Zero).Minus(p0), (_index104).Int())
			var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(971), _dafny.ArrayLen((_768_v32), 0))
			_ = _index105
			(_768_v32).ArraySet1((_this).F8(), (_index105).Int())
		}
		var _hi3 _dafny.Int = Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus(p0), p0)
		_ = _hi3
		for _769_i2 := (_dafny.Zero).Minus((_this).F8()); _769_i2.Cmp(_hi3) < 0; _769_i2 = _769_i2.Plus(_dafny.One) {
			Companion_Default___.M0(globalState)
			var _770_v33 _dafny.Sequence
			_ = _770_v33
			_770_v33 = _dafny.UnicodeSeqOfUtf8Bytes("mm")
			_770_v33 = _dafny.Companion_Sequence_.Update(_770_v33, (Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p0, p0), _dafny.IntOfUint32((_770_v33).Cardinality()))).Uint32(), _this.F4())
			var _771_v34 _dafny.Array
			_ = _771_v34
			var _nwElement0_24 bool = (_this).F7()
			_ = _nwElement0_24
			var _nw142 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_24, nil, _dafny.IntOfInt64(7))
			_ = _nw142
			(_nw142).ArraySet1(_nwElement0_24, 0)
			(_nw142).ArraySet1(true, 1)
			(_nw142).ArraySet1((_this).F6(), 2)
			(_nw142).ArraySet1(Companion_Default___.Fm27(_dafny.IntOfUint32(((_this).F5()).Cardinality()), (_this).F9(), (_this).F6(), globalState), 3)
			(_nw142).ArraySet1((_this).F9(), 4)
			(_nw142).ArraySet1((_this).F6(), 5)
			(_nw142).ArraySet1((_this).F9(), 6)
			_771_v34 = _nw142
			var _772_v35 _dafny.Array
			_ = _772_v35
			var _nwElement0_25 _dafny.Sequence = _dafny.SeqOf((_this).F7(), false)
			_ = _nwElement0_25
			var _nw143 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_25, nil, _dafny.IntOfInt64(26))
			_ = _nw143
			(_nw143).ArraySet1(_nwElement0_25, 0)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F7()), 1)
			(_nw143).ArraySet1(_733_v1, 2)
			(_nw143).ArraySet1(_733_v1, 3)
			(_nw143).ArraySet1(_733_v1, 4)
			(_nw143).ArraySet1(_733_v1, 5)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F9()), 6)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F9()), 7)
			(_nw143).ArraySet1(_733_v1, 8)
			(_nw143).ArraySet1(_733_v1, 9)
			(_nw143).ArraySet1(_733_v1, 10)
			(_nw143).ArraySet1(_733_v1, 11)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F7()), 12)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F7()), 13)
			(_nw143).ArraySet1(_733_v1, 14)
			(_nw143).ArraySet1(_733_v1, 15)
			(_nw143).ArraySet1(_733_v1, 16)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F6(), (_this).F7()), 17)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F9()), 18)
			(_nw143).ArraySet1(Companion_Default___.Fm29(globalState), 19)
			(_nw143).ArraySet1(_733_v1, 20)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F9()), 21)
			(_nw143).ArraySet1(_733_v1, 22)
			(_nw143).ArraySet1(_733_v1, 23)
			(_nw143).ArraySet1(_dafny.SeqOf((_this).F6()), 24)
			(_nw143).ArraySet1(_733_v1, 25)
			_772_v35 = _nw143
			var _773_v36 D1
			_ = _773_v36
			_773_v36 = Companion_D1_.Create_DC3_(_772_v35)
			var _774_v37 _dafny.Array
			_ = _774_v37
			var _nwElement0_26 D1 = _773_v36
			_ = _nwElement0_26
			var _nw144 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_26, nil, _dafny.IntOfInt64(26))
			_ = _nw144
			(_nw144).ArraySet1(_nwElement0_26, 0)
			(_nw144).ArraySet1(_773_v36, 1)
			(_nw144).ArraySet1(Companion_D1_.Create_DC3_(_772_v35), 2)
			(_nw144).ArraySet1(_773_v36, 3)
			(_nw144).ArraySet1(_773_v36, 4)
			(_nw144).ArraySet1(_773_v36, 5)
			(_nw144).ArraySet1(Companion_D1_.Create_DC3_(_772_v35), 6)
			(_nw144).ArraySet1(_773_v36, 7)
			(_nw144).ArraySet1(_773_v36, 8)
			(_nw144).ArraySet1(Companion_D1_.Create_DC3_(_772_v35), 9)
			(_nw144).ArraySet1(_773_v36, 10)
			(_nw144).ArraySet1(_773_v36, 11)
			(_nw144).ArraySet1(_773_v36, 12)
			(_nw144).ArraySet1(_773_v36, 13)
			(_nw144).ArraySet1(_773_v36, 14)
			(_nw144).ArraySet1(Companion_D1_.Create_DC3_(_772_v35), 15)
			(_nw144).ArraySet1(Companion_D1_.Create_DC3_(_772_v35), 16)
			(_nw144).ArraySet1(_773_v36, 17)
			(_nw144).ArraySet1(_773_v36, 18)
			(_nw144).ArraySet1(_773_v36, 19)
			(_nw144).ArraySet1(_773_v36, 20)
			(_nw144).ArraySet1(Companion_D1_.Create_DC3_(_772_v35), 21)
			(_nw144).ArraySet1(_773_v36, 22)
			(_nw144).ArraySet1(_773_v36, 23)
			(_nw144).ArraySet1(_773_v36, 24)
			(_nw144).ArraySet1(_773_v36, 25)
			_774_v37 = _nw144
			var _775_v38 *C2
			_ = _775_v38
			var _nw145 *C2 = New_C2_()
			_ = _nw145
			_nw145.Ctor__(_771_v34, _774_v37)
			_775_v38 = _nw145
			var _776_v39 _dafny.Sequence
			_ = _776_v39
			_776_v39 = _dafny.SeqOf(_775_v38, _775_v38)
			_775_v38 = (_776_v39).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(p0, p0), _dafny.IntOfUint32((_776_v39).Cardinality()))).Uint32()).(*C2)
			if (_this).F7() {
				var _777_v40 _dafny.Map
				_ = _777_v40
				_777_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm30((_this).F9(), globalState)).Cardinality(), (_732_v0).Cardinality())
				var _778_v41 _dafny.Map
				_ = _778_v41
				_778_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_777_v40, (_this).F8())
				var _779_v42 _dafny.Map
				_ = _779_v42
				_779_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), false)
				var _780_v43 _dafny.Sequence
				_ = _780_v43
				_780_v43 = _dafny.SeqOf(p0, (_this).F8())
				_778_v41 = (_778_v41).Update((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_732_v0).Cardinality(), p0)).Update((_dafny.Zero).Minus(_769_i2), ((_779_v42).Update(!((_this).F7()), (_this).F7())).Cardinality()), (_769_i2).Minus((_780_v43).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_780_v43).Cardinality()))).Uint32()).(_dafny.Int)))
				var _781_v44 bool
				_ = _781_v44
				_781_v44 = true
				var _782_v45 _dafny.Set
				_ = _782_v45
				_782_v45 = _dafny.SetOf((_this).F7())
				var _783_v46 D3
				_ = _783_v46
				_783_v46 = Companion_D3_.Create_DC11_(false, p0, _782_v45)
				var _784_v47 _dafny.MultiSet
				_ = _784_v47
				_784_v47 = _dafny.MultiSetOf((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_780_v43).Cardinality()), _dafny.IntOfInt64(646))))
				var _rhs100 bool = (func(_pat_let22_0 D3) D3 {
					return func(_785_dt__update__tmp_h0 D3) D3 {
						return func(_pat_let23_0 _dafny.Set) D3 {
							return func(_786_dt__update_hcf22_h0 _dafny.Set) D3 {
								return Companion_D3_.Create_DC11_((_785_dt__update__tmp_h0).Dtor_cf20(), (_785_dt__update__tmp_h0).Dtor_cf21(), _786_dt__update_hcf22_h0)
							}(_pat_let23_0)
						}(_dafny.SetOf((_this).F9()))
					}(_pat_let22_0)
				}(_783_v46)).Dtor_cf20()
				_ = _rhs100
				var _rhs101 _dafny.Int = (_784_v47).Cardinality()
				_ = _rhs101
				var _lhs70 *GlobalState = globalState
				_ = _lhs70
				_781_v44 = _rhs100
				_lhs70.F1 = _rhs101
				var _787_v48 _dafny.Map
				_ = _787_v48
				_787_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((true) || ((_this).F9()), _this.F4())
				_787_v48 = _787_v48
				_781_v44 = (_this).F9()
				var _788_v49 _dafny.Set
				_ = _788_v49
				_788_v49 = _dafny.SetOf(p0)
				var _789_v50 _dafny.Map
				_ = _789_v50
				_789_v50 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.SeqOf(_dafny.IntOfUint32(((_this).F5()).Cardinality())))
				var _790_v51 _dafny.Map
				_ = _790_v51
				_790_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_788_v49).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_733_v1).Cardinality()), (func() _dafny.Sequence {
					if (_789_v50).Contains((_this).F8()) {
						return (_789_v50).Get((_this).F8()).(_dafny.Sequence)
					}
					return _780_v43
				})())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(264), _780_v43)))
				_790_v51 = (_790_v51).Update((_this).F8(), _789_v50)
			} else {
				var _791_v52 bool
				_ = _791_v52
				_791_v52 = false
				_791_v52 = !((func() bool {
					if true {
						return _791_v52
					}
					return (_this).F7()
				})())
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_775_v38).F14()), 0))
				_ = _index106
				((_775_v38).F14()).ArraySet1((p0).Cmp(_769_i2) == 0, (_index106).Int())
				var _792_v53 D1
				_ = _792_v53
				_792_v53 = Companion_D1_.Create_DC4_(_this.F4(), _769_i2, _this.F4(), p0, (_this).F6())
				var _793_v54 _dafny.Set
				_ = _793_v54
				_793_v54 = _dafny.SetOf(_792_v53)
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_775_v38).F14()), 0))
				_ = _index107
				((_775_v38).F14()).ArraySet1(!(_793_v54).Equals(_793_v54), (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_775_v38).F14()), 0))
				_ = _index108
				((_775_v38).F14()).ArraySet1(_dafny.Companion_Sequence_.Contains((func() _dafny.Sequence {
					if !(_791_v52) {
						return (_this).F5()
					}
					return (_this).F5()
				})(), _this.F4()), (_index108).Int())
				var _794_v55 *C0
				_ = _794_v55
				var _nw146 *C0 = New_C0_()
				_ = _nw146
				_nw146.Ctor__(((_775_v38).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_775_v38).F14()), 0))).Int()).(bool), ((_775_v38).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(59), _dafny.ArrayLen(((_775_v38).F14()), 0))).Int()).(bool), _this.F4(), (_this).F5())
				_794_v55 = _nw146
				var _795_v56 _dafny.Map
				_ = _795_v56
				_795_v56 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_794_v55, _dafny.CodePoint('m'))
				(globalState).F1 = Companion_Default___.SafeModuloInt((_this).F8(), ((_795_v56).Merge(_795_v56)).Cardinality())
				var _796_v57 _dafny.Map
				_ = _796_v57
				_796_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _791_v52)
				var _797_v58 *C3
				_ = _797_v58
				var _nw147 *C3 = New_C3_()
				_ = _nw147
				_nw147.Ctor__(_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p0), (_this).F8()), (func() bool {
					if (_796_v57).Contains(_769_i2) {
						return (_796_v57).Get(_769_i2).(bool)
					}
					return (_this).F6()
				})(), (_dafny.Zero).Minus((p0).Minus((_this).F8())), (_this).F7(), (_this).F7(), (_this).F7(), _this.F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(665))).Uint32(), func(coer47 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg47 _dafny.Int) interface{} {
						return coer47(arg47)
					}
				}(func(_798_i3 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('t')
				})))
				_797_v58 = _nw147
			}
		}
		var _799_v59 _dafny.Array
		_ = _799_v59
		var _nw148 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.IntOfInt64(19))
		_ = _nw148
		_799_v59 = _nw148
		var _800_v60 D6
		_ = _800_v60
		_800_v60 = Companion_D6_.Create_DC17_((_this).F5(), (_this).F5())
		var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_799_v59), 0))
		_ = _index109
		(_799_v59).ArraySet1((func() D6 {
			if (_this).F7() {
				return _800_v60
			}
			return Companion_D6_.Create_DC17_((_this).F5(), (_this).F5())
		})(), (_index109).Int())
		var _801_v61 _dafny.Map
		_ = _801_v61
		_801_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _dafny.Companion_Sequence_.Update(_733_v1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_733_v1).Cardinality()))).Uint32(), (_this).F9()))
		var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_799_v59), 0))
		_ = _index110
		var _rhs102 _dafny.Int = (_this).F8()
		_ = _rhs102
		var _rhs103 D6 = _800_v60
		_ = _rhs103
		var _rhs104 _dafny.Int = _dafny.IntOfInt64(26)
		_ = _rhs104
		var _rhs105 _dafny.MultiSet = _dafny.MultiSetFromSeq(_dafny.Companion_Sequence_.Concatenate(_733_v1, (func() _dafny.Sequence {
			if (_801_v61).Contains(p0) {
				return (_801_v61).Get(p0).(_dafny.Sequence)
			}
			return _733_v1
		})()))
		_ = _rhs105
		var _lhs71 *GlobalState = globalState
		_ = _lhs71
		var _lhs72 _dafny.Array = _799_v59
		_ = _lhs72
		var _lhs73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(553), _dafny.ArrayLen((_799_v59), 0))
		_ = _lhs73
		var _lhs74 *GlobalState = globalState
		_ = _lhs74
		_lhs71.F1 = _rhs102
		(_lhs72).ArraySet1(_rhs103, (_lhs73).Int())
		_lhs74.F1 = _rhs104
		_732_v0 = _rhs105
		var _hi4 _dafny.Int = (_dafny.Zero).Minus((_this).F8())
		_ = _hi4
		for _802_i4 := (func() _dafny.Int {
			if (_this).F7() {
				return p0
			}
			return (_this).F8()
		})(); _802_i4.Cmp(_hi4) < 0; _802_i4 = _802_i4.Plus(_dafny.One) {
			if (_733_v1).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_802_i4), _dafny.IntOfUint32((_733_v1).Cardinality()))).Uint32()).(bool) {
				var _803_v62 _dafny.Array
				_ = _803_v62
				var _nw149 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(28))
				_ = _nw149
				_803_v62 = _nw149
				var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_803_v62), 0))
				_ = _index111
				(_803_v62).ArraySet1((_dafny.Zero).Minus(_802_i4), (_index111).Int())
				var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_803_v62), 0))
				_ = _index112
				(_803_v62).ArraySet1(p0, (_index112).Int())
				var _804_v63 _dafny.Map
				_ = _804_v63
				_804_v63 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _this.F4())
				(globalState).F1 = (func() _dafny.Int {
					if (_this).F6() {
						return _802_i4
					}
					return ((_803_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_803_v62), 0))).Int()).(_dafny.Int)).Times((_804_v63).Cardinality())
				})()
				var _805_v64 _dafny.Array
				_ = _805_v64
				var _nw150 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(28))
				_ = _nw150
				_805_v64 = _nw150
				var _806_v65 _dafny.Map
				_ = _806_v65
				_806_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F8())
				var _807_v66 _dafny.MultiSet
				_ = _807_v66
				_807_v66 = _dafny.MultiSetOf((_this).F8())
				var _808_v67 D5
				_ = _808_v67
				_808_v67 = Companion_D5_.Create_DC15_(_805_v64, (_this).F9(), _806_v65, (_807_v66).Cardinality())
				var _809_v68 _dafny.Map
				_ = _809_v68
				_809_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_808_v67).Dtor_cf27())
				var _810_v69 _dafny.Sequence
				_ = _810_v69
				_810_v69 = _dafny.SeqOf(_809_v68, _809_v68)
				var _811_v70 _dafny.Map
				_ = _811_v70
				_811_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), _dafny.IntOfUint32((_810_v69).Cardinality()))
				_806_v65 = (_806_v65).Update((_this).F6(), (_this).Fm1((_this).F6(), _809_v68, (_this).F9(), globalState))
				(globalState).F1 = _802_i4
				var _812_v71 bool
				_ = _812_v71
				_812_v71 = true
				var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_803_v62), 0))
				_ = _index113
				var _rhs106 bool = (p0).Cmp((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _805_v64)).Cardinality()) > 0
				_ = _rhs106
				var _rhs107 _dafny.Int = Companion_Default___.SafeDivisionInt(_802_i4, (_this).F8())
				_ = _rhs107
				var _lhs75 _dafny.Array = _803_v62
				_ = _lhs75
				var _lhs76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(883), _dafny.ArrayLen((_803_v62), 0))
				_ = _lhs76
				_812_v71 = _rhs106
				(_lhs75).ArraySet1(_rhs107, (_lhs76).Int())
			} else {
				(globalState).F1 = _dafny.IntOfInt64(382)
				var _813_v72 bool
				_ = _813_v72
				_813_v72 = false
				_813_v72 = (_this).F9()
				var _814_v73 _dafny.Set
				_ = _814_v73
				_814_v73 = _dafny.SetOf(_this.F4(), _this.F4())
				var _815_v74 _dafny.Array
				_ = _815_v74
				var _nwElement0_27 bool = _813_v72
				_ = _nwElement0_27
				var _nw151 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_27, nil, _dafny.IntOfInt64(11))
				_ = _nw151
				(_nw151).ArraySet1(_nwElement0_27, 0)
				(_nw151).ArraySet1((_this).F6(), 1)
				(_nw151).ArraySet1((_this).F6(), 2)
				(_nw151).ArraySet1((_this).F7(), 3)
				(_nw151).ArraySet1((_this).F9(), 4)
				(_nw151).ArraySet1(((_this).F8()).Cmp(p0) >= 0, 5)
				(_nw151).ArraySet1((_dafny.SetOf(_this.F4())).IsSubsetOf(_814_v73), 6)
				(_nw151).ArraySet1((_813_v72) && ((_this).F7()), 7)
				(_nw151).ArraySet1(_813_v72, 8)
				(_nw151).ArraySet1((_this).F7(), 9)
				(_nw151).ArraySet1(true, 10)
				_815_v74 = _nw151
				var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_815_v74), 0))
				_ = _index114
				(_815_v74).ArraySet1(_813_v72, (_index114).Int())
				var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_815_v74), 0))
				_ = _index115
				(_815_v74).ArraySet1((_this).F7(), (_index115).Int())
				var _816_v75 D2
				_ = _816_v75
				_816_v75 = Companion_D2_.Create_DC8_(p0, p0)
				var _817_v76 D2
				_ = _817_v76
				_817_v76 = Companion_D2_.Create_DC9_(_816_v75)
				_817_v76 = (func() D2 {
					if true {
						return _817_v76
					}
					return _817_v76
				})()
				var _len0_15 _dafny.Int = _dafny.IntOfInt64(20)
				_ = _len0_15
				var _nw152 _dafny.Array
				_ = _nw152
				if _len0_15.Cmp(_dafny.Zero) == 0 {
					_nw152 = _dafny.NewArray(_len0_15)
				} else {
					var _init15 func(_dafny.Int) bool = (func(_818_v74 _dafny.Array) func(_dafny.Int) bool {
						return func(_819_i5 _dafny.Int) bool {
							return (_818_v74).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(174), _dafny.ArrayLen((_818_v74), 0))).Int()).(bool)
						}
					})(_815_v74)
					_ = _init15
					var _element0_15 = _init15(_dafny.Zero)
					_ = _element0_15
					_nw152 = _dafny.NewArrayFromExample(_element0_15, nil, _len0_15)
					(_nw152).ArraySet1(_element0_15, 0)
					var _nativeLen0_15 = (_len0_15).Int()
					_ = _nativeLen0_15
					for _i0_15 := 1; _i0_15 < _nativeLen0_15; _i0_15++ {
						(_nw152).ArraySet1(_init15(_dafny.IntOf(_i0_15)), _i0_15)
					}
				}
				_815_v74 = _nw152
			}
			var _820_v77 _dafny.Array
			_ = _820_v77
			var _nw153 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(27))
			_ = _nw153
			_820_v77 = _nw153
			var _821_v78 _dafny.Set
			_ = _821_v78
			_821_v78 = _dafny.SetOf(_820_v77)
			var _822_v79 D7
			_ = _822_v79
			_822_v79 = Companion_D7_.Create_DC18_(_821_v78)
			var _pat_let_tv15 = _820_v77
			_ = _pat_let_tv15
			_822_v79 = func(_pat_let24_0 D7) D7 {
				return func(_823_dt__update__tmp_h1 D7) D7 {
					return func(_pat_let25_0 _dafny.Set) D7 {
						return func(_824_dt__update_hcf33_h0 _dafny.Set) D7 {
							return Companion_D7_.Create_DC18_(_824_dt__update_hcf33_h0)
						}(_pat_let25_0)
					}(_dafny.SetOf(_pat_let_tv15))
				}(_pat_let24_0)
			}(Companion_D7_.Create_DC18_(_821_v78))
			var _825_v80 bool
			_ = _825_v80
			_825_v80 = false
			_825_v80 = false
			var _826_v81 _dafny.Map
			_ = _826_v81
			_826_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F8())
			var _827_v82 _dafny.Map
			_ = _827_v82
			_827_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F6())
			var _828_v83 D3
			_ = _828_v83
			_828_v83 = Companion_D3_.Create_DC11_((_this).F6(), p0, _dafny.SetOf((_this).F6()))
			var _829_v84 _dafny.Array
			_ = _829_v84
			var _len0_16 _dafny.Int = _dafny.IntOfInt64(28)
			_ = _len0_16
			var _nw154 _dafny.Array
			_ = _nw154
			if _len0_16.Cmp(_dafny.Zero) == 0 {
				_nw154 = _dafny.NewArray(_len0_16)
			} else {
				var _init16 func(_dafny.Int) _dafny.Int = (func(_830_v82 _dafny.Map) func(_dafny.Int) _dafny.Int {
					return func(_831_i6 _dafny.Int) _dafny.Int {
						return Companion_Default___.SafeDivisionInt(_831_i6, (_830_v82).Cardinality())
					}
				})(_827_v82)
				_ = _init16
				var _element0_16 = _init16(_dafny.Zero)
				_ = _element0_16
				_nw154 = _dafny.NewArrayFromExample(_element0_16, nil, _len0_16)
				(_nw154).ArraySet1(_element0_16, 0)
				var _nativeLen0_16 = (_len0_16).Int()
				_ = _nativeLen0_16
				for _i0_16 := 1; _i0_16 < _nativeLen0_16; _i0_16++ {
					(_nw154).ArraySet1(_init16(_dafny.IntOf(_i0_16)), _i0_16)
				}
			}
			_829_v84 = _nw154
			var _832_v85 _dafny.MultiSet
			_ = _832_v85
			_832_v85 = _dafny.MultiSetOf(_829_v84, _829_v84, _829_v84, _829_v84, _829_v84)
			var _833_v86 _dafny.Array
			_ = _833_v86
			var _nwElement0_28 _dafny.Int = _802_i4
			_ = _nwElement0_28
			var _nw155 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_28, nil, _dafny.IntOfInt64(24))
			_ = _nw155
			(_nw155).ArraySet1(_nwElement0_28, 0)
			(_nw155).ArraySet1(((func() _dafny.Int {
				if (_826_v81).Contains((_this).F6()) {
					return (_826_v81).Get((_this).F6()).(_dafny.Int)
				}
				return _dafny.IntOfUint32(((_this).F5()).Cardinality())
			})()).Minus(_dafny.IntOfInt64(456)), 1)
			(_nw155).ArraySet1(_dafny.IntOfInt64(280), 2)
			(_nw155).ArraySet1(((_this).Fm1((_this).F7(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F9()), (_this).F7(), globalState)).Times(p0), 3)
			(_nw155).ArraySet1((_this).F8(), 4)
			(_nw155).ArraySet1((_this).Fm1(_825_v80, (_827_v82).Update((_this).F6(), (_this).F7()), (_this).F7(), globalState), 5)
			(_nw155).ArraySet1((_this).F8(), 6)
			(_nw155).ArraySet1(_802_i4, 7)
			(_nw155).ArraySet1((_828_v83).Dtor_cf21(), 8)
			(_nw155).ArraySet1((_dafny.Zero).Minus(_802_i4), 9)
			(_nw155).ArraySet1((_this).F8(), 10)
			(_nw155).ArraySet1((func() _dafny.Int {
				if (_this).F6() {
					return _dafny.IntOfUint32(((_this).F5()).Cardinality())
				}
				return _802_i4
			})(), 11)
			(_nw155).ArraySet1((_dafny.Zero).Minus(p0), 12)
			(_nw155).ArraySet1(((_832_v85).Cardinality()).Times((_this).F8()), 13)
			(_nw155).ArraySet1(_802_i4, 14)
			(_nw155).ArraySet1((_802_i4).Plus(p0), 15)
			(_nw155).ArraySet1((_dafny.Zero).Minus(p0), 16)
			(_nw155).ArraySet1(_802_i4, 17)
			(_nw155).ArraySet1(_dafny.IntOfUint32(((_this).F5()).Cardinality()), 18)
			(_nw155).ArraySet1((_this).F8(), 19)
			(_nw155).ArraySet1((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_802_i4, _dafny.IntOfUint32(((_this).F5()).Cardinality()))), 20)
			(_nw155).ArraySet1(_802_i4, 21)
			(_nw155).ArraySet1((_this).F8(), 22)
			(_nw155).ArraySet1(_dafny.IntOfInt64(-869), 23)
			_833_v86 = _nw155
			var _rhs108 _dafny.Int = (_dafny.IntOfInt64(777)).Minus((_this).F8())
			_ = _rhs108
			var _rhs109 _dafny.Array = _829_v84
			_ = _rhs109
			var _rhs110 _dafny.Int = Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(578), _802_i4)
			_ = _rhs110
			var _rhs111 _dafny.Int = p0
			_ = _rhs111
			var _lhs77 *GlobalState = globalState
			_ = _lhs77
			var _lhs78 *GlobalState = globalState
			_ = _lhs78
			var _lhs79 *GlobalState = globalState
			_ = _lhs79
			_lhs77.F1 = _rhs108
			_833_v86 = _rhs109
			_lhs78.F1 = _rhs110
			_lhs79.F1 = _rhs111
		}
		(globalState).F1 = (_this).F8()
		var _834_i7 _dafny.Int
		_ = _834_i7
		_834_i7 = _dafny.Zero
		{
			for (p0).Cmp(_dafny.IntOfUint32(((_this).F5()).Cardinality())) < 0 {
				{
					if (_834_i7).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L4
					}
					_834_i7 = (_834_i7).Plus(_dafny.One)
					var _835_v87 bool
					_ = _835_v87
					_835_v87 = false
					_835_v87 = _dafny.Companion_Sequence_.IsPrefixOf((_this).F5(), (_this).F5())
					var _836_v88 _dafny.Map
					_ = _836_v88
					_836_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_835_v87, (_this).F9())
					var _837_v89 _dafny.Map
					_ = _837_v89
					_837_v89 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_835_v87, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(58))).Uint32(), func(coer48 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
						return func(arg48 _dafny.Int) interface{} {
							return coer48(arg48)
						}
					}((func(_838_p0 _dafny.Int) func(_dafny.Int) _dafny.Int {
						return func(_839_i8 _dafny.Int) _dafny.Int {
							return _838_p0
						}
					})(p0))))
					(globalState).F1 = (_this).Fm1(true, (_836_v88).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_835_v87, true)), !(!(_837_v89).Equals(_837_v89)), globalState)
					var _840_v90 _dafny.Array
					_ = _840_v90
					var _nw156 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(11))
					_ = _nw156
					_840_v90 = _nw156
					var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_840_v90), 0))
					_ = _index116
					(_840_v90).ArraySet1((_this).F8(), (_index116).Int())
					var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_840_v90), 0))
					_ = _index117
					(_840_v90).ArraySet1(p0, (_index117).Int())
					var _841_v91 _dafny.Map
					_ = _841_v91
					_841_v91 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(97), false)
					var _842_v92 _dafny.Set
					_ = _842_v92
					_842_v92 = _dafny.SetOf(true)
					var _843_v93 D3
					_ = _843_v93
					_843_v93 = Companion_D3_.Create_DC11_((func() bool {
						if (_841_v91).Contains((_dafny.Zero).Minus(p0)) {
							return (_841_v91).Get((_dafny.Zero).Minus(p0)).(bool)
						}
						return (_this).F7()
					})(), _dafny.IntOfInt64(26), _842_v92)
					var _844_v94 T1
					_ = _844_v94
					var _nw157 *C0 = New_C0_()
					_ = _nw157
					_nw157.Ctor__(_835_v87, (_this).F7(), _this.F4(), Companion_Default___.Fm28((_840_v90).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(296), _dafny.ArrayLen((_840_v90), 0))).Int()).(_dafny.Int), (_this).F6(), _733_v1, _843_v93, globalState))
					_844_v94 = _nw157
					var _845_v95 _dafny.Array
					_ = _845_v95
					var _nw158 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
					_ = _nw158
					_845_v95 = _nw158
					var _846_v96 _dafny.Array
					_ = _846_v96
					var _nw159 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(12))
					_ = _nw159
					_846_v96 = _nw159
					var _847_v97 _dafny.Array
					_ = _847_v97
					var _nwElement0_29 _dafny.Array = _846_v96
					_ = _nwElement0_29
					var _nw160 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_29, nil, _dafny.IntOfInt64(3))
					_ = _nw160
					(_nw160).ArraySet1(_nwElement0_29, 0)
					(_nw160).ArraySet1(_846_v96, 1)
					(_nw160).ArraySet1(_846_v96, 2)
					_847_v97 = _nw160
					var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_845_v95), 0))
					_ = _index118
					(_845_v95).ArraySet1(_847_v97, (_index118).Int())
					var _848_v98 _dafny.Array
					_ = _848_v98
					var _len0_17 _dafny.Int = _dafny.IntOfInt64(4)
					_ = _len0_17
					var _nw161 _dafny.Array
					_ = _nw161
					if _len0_17.Cmp(_dafny.Zero) == 0 {
						_nw161 = _dafny.NewArray(_len0_17)
					} else {
						var _init17 func(_dafny.Int) D0 = (func(_849_v94 T1, _850_p0 _dafny.Int, _851_v87 bool) func(_dafny.Int) D0 {
							return func(_852_i9 _dafny.Int) D0 {
								return Companion_D0_.Create_DC1_((_849_v94).F6(), _850_p0, _851_v87, (_this).F8())
							}
						})(_844_v94, p0, _835_v87)
						_ = _init17
						var _element0_17 = _init17(_dafny.Zero)
						_ = _element0_17
						_nw161 = _dafny.NewArrayFromExample(_element0_17, nil, _len0_17)
						(_nw161).ArraySet1(_element0_17, 0)
						var _nativeLen0_17 = (_len0_17).Int()
						_ = _nativeLen0_17
						for _i0_17 := 1; _i0_17 < _nativeLen0_17; _i0_17++ {
							(_nw161).ArraySet1(_init17(_dafny.IntOf(_i0_17)), _i0_17)
						}
					}
					_848_v98 = _nw161
					var _853_v99 D0
					_ = _853_v99
					_853_v99 = Companion_D0_.Create_DC1_((_this).F9(), (_dafny.Zero).Minus((_this).F8()), (_844_v94).F7(), _dafny.IntOfUint32(((_844_v94).F5()).Cardinality()))
					var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_848_v98), 0))
					_ = _index119
					(_848_v98).ArraySet1(_853_v99, (_index119).Int())
					var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_845_v95), 0))
					_ = _index120
					var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_848_v98), 0))
					_ = _index121
					var _rhs112 T1 = (func() T1 {
						if (_this).F6() {
							return _844_v94
						}
						return _844_v94
					})()
					_ = _rhs112
					var _rhs113 _dafny.Array = _847_v97
					_ = _rhs113
					var _rhs114 D0 = _853_v99
					_ = _rhs114
					var _lhs80 _dafny.Array = _845_v95
					_ = _lhs80
					var _lhs81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(422), _dafny.ArrayLen((_845_v95), 0))
					_ = _lhs81
					var _lhs82 _dafny.Array = _848_v98
					_ = _lhs82
					var _lhs83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(25), _dafny.ArrayLen((_848_v98), 0))
					_ = _lhs83
					_844_v94 = _rhs112
					(_lhs80).ArraySet1(_rhs113, (_lhs81).Int())
					(_lhs82).ArraySet1(_rhs114, (_lhs83).Int())
					goto C4
				}
			C4:
			}
			goto L4
		}
	L4:
		var _854_v100 _dafny.Map
		_ = _854_v100
		_854_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, ((_this).F8()).Plus(p0))
		r0 = _854_v100
		return r0
	}
}

// End of class C4

// Definition of class C5
type C5 struct {
	_f4  _dafny.CodePoint
	_f8  _dafny.Int
	_f9  bool
	_f6  bool
	_f7  bool
	_f5  _dafny.Sequence
	_f10 _dafny.Array
	_f11 _dafny.Array
}

func New_C5_() *C5 {
	_this := C5{}

	_this._f4 = _dafny.CodePoint('D')
	_this._f8 = _dafny.Zero
	_this._f9 = false
	_this._f6 = false
	_this._f7 = false
	_this._f5 = _dafny.EmptySeq
	_this._f10 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f11 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
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
	return [](*_dafny.TraitID){Companion_T1_.TraitID_, Companion_T2_.TraitID_, Companion_T3_.TraitID_, Companion_T0_.TraitID_}
}

var _ T1 = &C5{}
var _ T2 = &C5{}
var _ T3 = &C5{}
var _ T0 = &C5{}
var _ _dafny.TraitOffspring = &C5{}

func (_this *C5) F4() _dafny.CodePoint {
	return _this._f4
}
func (_this *C5) F4_set_(value _dafny.CodePoint) {
	_this._f4 = value
}
func (_this *C5) F8() _dafny.Int {
	return _this._f8
}
func (_this *C5) F9() bool {
	return _this._f9
}
func (_this *C5) F6() bool {
	return _this._f6
}
func (_this *C5) F7() bool {
	return _this._f7
}
func (_this *C5) F5() _dafny.Sequence {
	return _this._f5
}
func (_this *C5) Ctor__(f10 _dafny.Array, f11 _dafny.Array, f6 bool, f7 bool, f4 _dafny.CodePoint, f5 _dafny.Sequence, f8 _dafny.Int, f9 bool) {
	{
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this)._f6 = f6
		(_this)._f7 = f7
		(_this)._f4 = f4
		(_this)._f5 = f5
		(_this)._f8 = f8
		(_this)._f9 = f9
	}
}
func (_this *C5) Fm1(p0 bool, p1 _dafny.Map, p2 bool, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus((_this).F8())
	}
}
func (_this *C5) Fm2(globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf((_this).F6())).Cardinality()), (_this).F8()), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F8()), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(342))).Uint32(), func(coer49 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
			return func(arg49 _dafny.Int) interface{} {
				return coer49(arg49)
			}
		}(func(_855_i0 _dafny.Int) _dafny.Int {
			return (_this).F8()
		}))))
	}
}
func (_this *C5) M1(p0 bool, p1 bool, p2 _dafny.MultiSet, p3 bool, globalState *GlobalState) (_dafny.Int, _dafny.Int, _dafny.Array) {
	{
		var r0 _dafny.Int = _dafny.Zero
		_ = r0
		var r1 _dafny.Int = _dafny.Zero
		_ = r1
		var r2 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
		_ = r2
		var _856_v0 D0
		_ = _856_v0
		_856_v0 = Companion_D0_.Create_DC0_()
		var _source9 D0 = func(_source10 D0) D0 {
			if _source10.Is_DC0() {
				return Companion_D0_.Create_DC1_((_this).F6(), (_this).F8(), (_this).F9(), (_this).F8())
			} else if _source10.Is_DC1() {
				var _857___mcc_h7 bool = _source10.Get_().(D0_DC1).Cf0
				_ = _857___mcc_h7
				var _858___mcc_h8 _dafny.Int = _source10.Get_().(D0_DC1).Cf1
				_ = _858___mcc_h8
				var _859___mcc_h9 bool = _source10.Get_().(D0_DC1).Cf2
				_ = _859___mcc_h9
				var _860___mcc_h10 _dafny.Int = _source10.Get_().(D0_DC1).Cf3
				_ = _860___mcc_h10
				var _861_cf3 _dafny.Int = _860___mcc_h10
				_ = _861_cf3
				var _862_cf2 bool = _859___mcc_h9
				_ = _862_cf2
				var _863_cf1 _dafny.Int = _858___mcc_h8
				_ = _863_cf1
				var _864_cf0 bool = _857___mcc_h7
				_ = _864_cf0
				return Companion_D0_.Create_DC1_((_this).F9(), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(528))).Uint32(), func(coer50 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg50 _dafny.Int) interface{} {
						return coer50(arg50)
					}
				}(func(_865_i0 _dafny.Int) _dafny.Int {
					return (_this).F8()
				}))).Cardinality()), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Cardinality()), _864_cf0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, (_this).F9())).Cardinality())
			} else {
				var _866___mcc_h11 _dafny.Int = _source10.Get_().(D0_DC2).Cf4
				_ = _866___mcc_h11
				var _867___mcc_h12 bool = _source10.Get_().(D0_DC2).Cf5
				_ = _867___mcc_h12
				var _868___mcc_h13 _dafny.Sequence = _source10.Get_().(D0_DC2).Cf6
				_ = _868___mcc_h13
				var _869_cf6 _dafny.Sequence = _868___mcc_h13
				_ = _869_cf6
				var _870_cf5 bool = _867___mcc_h12
				_ = _870_cf5
				var _871_cf4 _dafny.Int = _866___mcc_h11
				_ = _871_cf4
				return Companion_D0_.Create_DC1_(true, _dafny.IntOfInt64(556), (_this).F7(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_871_cf4, _870_cf5)).Cardinality())
			}
		}(_856_v0)
		_ = _source9
		if _source9.Is_DC0() {
			var _872_v1 bool
			_ = _872_v1
			_872_v1 = false
			_872_v1 = (_this).F9()
			_872_v1 = _872_v1
			var _873_v2 _dafny.Sequence
			_ = _873_v2
			_873_v2 = _dafny.SeqOf((_this).F6())
			var _874_v3 _dafny.Sequence
			_ = _874_v3
			_874_v3 = _dafny.SeqOf(_873_v2, _873_v2)
			var _875_v4 _dafny.Map
			_ = _875_v4
			_875_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(539), (_this).F8())
			var _876_v5 _dafny.Sequence
			_ = _876_v5
			_876_v5 = _dafny.SeqOf(_875_v4, _875_v4)
			var _877_v6 _dafny.Set
			_ = _877_v6
			_877_v6 = _dafny.SetOf((_this).F8())
			var _878_v7 _dafny.Map
			_ = _878_v7
			_878_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(727), (_this).F7())
			var _879_v8 _dafny.Array
			_ = _879_v8
			var _nwElement0_30 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_874_v3).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_876_v5).Cardinality()), _dafny.IntOfUint32((_874_v3).Cardinality()))).Uint32()).(_dafny.Sequence), _873_v2)
			_ = _nwElement0_30
			var _nw162 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_30, nil, _dafny.IntOfInt64(22))
			_ = _nw162
			(_nw162).ArraySet1(_nwElement0_30, 0)
			(_nw162).ArraySet1(_873_v2, 1)
			(_nw162).ArraySet1(_873_v2, 2)
			(_nw162).ArraySet1(_873_v2, 3)
			(_nw162).ArraySet1(_873_v2, 4)
			(_nw162).ArraySet1(_873_v2, 5)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_874_v3).Select((Companion_Default___.SafeIndex((_877_v6).Cardinality(), _dafny.IntOfUint32((_874_v3).Cardinality()))).Uint32()).(_dafny.Sequence), _873_v2), 6)
			(_nw162).ArraySet1(_873_v2, 7)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F7()), _dafny.SeqOf(true)), 8)
			(_nw162).ArraySet1(_873_v2, 9)
			(_nw162).ArraySet1(_873_v2, 10)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_873_v2, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(102), _dafny.IntOfUint32((_873_v2).Cardinality()))).Uint32(), _872_v1), _dafny.SeqOf(p0, p0)), 11)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_873_v2, _873_v2), 12)
			(_nw162).ArraySet1(_873_v2, 13)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F9(), !(false)), _873_v2), 14)
			(_nw162).ArraySet1(_873_v2, 15)
			(_nw162).ArraySet1(_dafny.SeqOf(p1, (func() bool {
				if (_878_v7).Contains((_this).F8()) {
					return (_878_v7).Get((_this).F8()).(bool)
				}
				return p1
			})(), Companion_Default___.Fm3(globalState)), 16)
			(_nw162).ArraySet1(_873_v2, 17)
			(_nw162).ArraySet1(_dafny.SeqOf((_this).F6()), 18)
			(_nw162).ArraySet1(_873_v2, 19)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_873_v2, _873_v2), 20)
			(_nw162).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F6(), p0), _873_v2), 21)
			_879_v8 = _nw162
			var _880_v9 _dafny.Sequence
			_ = _880_v9
			_880_v9 = _dafny.SeqOf((_this).F11())
			var _881_v10 _dafny.Array
			_ = _881_v10
			var _nwElement0_31 _dafny.Array = (_this).F11()
			_ = _nwElement0_31
			var _nw163 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_31, nil, _dafny.IntOfInt64(26))
			_ = _nw163
			(_nw163).ArraySet1(_nwElement0_31, 0)
			(_nw163).ArraySet1((_this).F11(), 1)
			(_nw163).ArraySet1((_this).F11(), 2)
			(_nw163).ArraySet1((_this).F11(), 3)
			(_nw163).ArraySet1((_this).F11(), 4)
			(_nw163).ArraySet1((_this).F11(), 5)
			(_nw163).ArraySet1((_this).F11(), 6)
			(_nw163).ArraySet1((_this).F11(), 7)
			(_nw163).ArraySet1((_this).F11(), 8)
			(_nw163).ArraySet1((_880_v9).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_880_v9).Cardinality()))).Uint32()).(_dafny.Array), 9)
			(_nw163).ArraySet1((_this).F11(), 10)
			(_nw163).ArraySet1((_this).F11(), 11)
			(_nw163).ArraySet1((_this).F11(), 12)
			(_nw163).ArraySet1((_this).F11(), 13)
			(_nw163).ArraySet1((_this).F11(), 14)
			(_nw163).ArraySet1((_this).F11(), 15)
			(_nw163).ArraySet1((_this).F11(), 16)
			(_nw163).ArraySet1((_880_v9).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_880_v9).Cardinality()))).Uint32()).(_dafny.Array), 17)
			(_nw163).ArraySet1((_this).F11(), 18)
			(_nw163).ArraySet1((_880_v9).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_880_v9).Cardinality()))).Uint32()).(_dafny.Array), 19)
			(_nw163).ArraySet1((_this).F11(), 20)
			(_nw163).ArraySet1((_this).F11(), 21)
			(_nw163).ArraySet1((_this).F11(), 22)
			(_nw163).ArraySet1((_this).F11(), 23)
			(_nw163).ArraySet1((_this).F11(), 24)
			(_nw163).ArraySet1((_this).F11(), 25)
			_881_v10 = _nw163
			var _882_v11 _dafny.Map
			_ = _882_v11
			_882_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F11())
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_881_v10), 0))
			_ = _index122
			(_881_v10).ArraySet1((func() _dafny.Array {
				if (_882_v11).Contains((Companion_Default___.Fm4(p0, (_this).F8(), (_this).F7(), (_877_v6).Cardinality(), globalState)).Cardinality()) {
					return (_882_v11).Get((Companion_Default___.Fm4(p0, (_this).F8(), (_this).F7(), (_877_v6).Cardinality(), globalState)).Cardinality()).(_dafny.Array)
				}
				return (_this).F11()
			})(), (_index122).Int())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_881_v10), 0))
			_ = _index123
			var _rhs115 _dafny.Array = (Companion_D1_.Create_DC3_(_879_v8)).Dtor_cf7()
			_ = _rhs115
			var _rhs116 _dafny.Array = (_this).F11()
			_ = _rhs116
			var _lhs84 _dafny.Array = _881_v10
			_ = _lhs84
			var _lhs85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(395), _dafny.ArrayLen((_881_v10), 0))
			_ = _lhs85
			_879_v8 = _rhs115
			(_lhs84).ArraySet1(_rhs116, (_lhs85).Int())
			_872_v1 = !(!(_872_v1))
		} else if _source9.Is_DC1() {
			var _883___mcc_h0 bool = _source9.Get_().(D0_DC1).Cf0
			_ = _883___mcc_h0
			var _884___mcc_h1 _dafny.Int = _source9.Get_().(D0_DC1).Cf1
			_ = _884___mcc_h1
			var _885___mcc_h2 bool = _source9.Get_().(D0_DC1).Cf2
			_ = _885___mcc_h2
			var _886___mcc_h3 _dafny.Int = _source9.Get_().(D0_DC1).Cf3
			_ = _886___mcc_h3
			var _887_cf3 _dafny.Int = _886___mcc_h3
			_ = _887_cf3
			var _888_cf2 bool = _885___mcc_h2
			_ = _888_cf2
			var _889_cf1 _dafny.Int = _884___mcc_h1
			_ = _889_cf1
			var _890_cf0 bool = _883___mcc_h0
			_ = _890_cf0
			_888_cf2 = !(_888_cf2)
			var _891_v12 _dafny.Sequence
			_ = _891_v12
			_891_v12 = _dafny.SeqOf(p1)
			var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index124
			((_this).F10()).ArraySet1((_891_v12).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_891_v12).Cardinality()))).Uint32()).(bool), (_index124).Int())
			var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index125
			((_this).F10()).ArraySet1(p0, (_index125).Int())
			var _892_v13 _dafny.Set
			_ = _892_v13
			_892_v13 = _dafny.SetOf(!(_890_cf0))
			var _893_v14 _dafny.Array
			_ = _893_v14
			var _nwElement0_32 bool = p1
			_ = _nwElement0_32
			var _nw164 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_32, nil, _dafny.IntOfInt64(19))
			_ = _nw164
			(_nw164).ArraySet1(_nwElement0_32, 0)
			(_nw164).ArraySet1(Companion_Default___.Fm7(_this.F4(), (_892_v13).Cardinality(), p0, globalState), 1)
			(_nw164).ArraySet1(p3, 2)
			(_nw164).ArraySet1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), 3)
			(_nw164).ArraySet1(!(Companion_Default___.Fm3(globalState)), 4)
			(_nw164).ArraySet1((_this).F9(), 5)
			(_nw164).ArraySet1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), 6)
			(_nw164).ArraySet1(p1, 7)
			(_nw164).ArraySet1(_890_cf0, 8)
			(_nw164).ArraySet1((_this).F9(), 9)
			(_nw164).ArraySet1((_this).F9(), 10)
			(_nw164).ArraySet1(!(_888_cf2), 11)
			(_nw164).ArraySet1((_891_v12).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer51 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg51 _dafny.Int) interface{} {
					return coer51(arg51)
				}
			}(func(_894_i1 _dafny.Int) _dafny.CodePoint {
				return _this.F4()
			}))).Cardinality())), _dafny.IntOfUint32((_891_v12).Cardinality()))).Uint32()).(bool), 12)
			(_nw164).ArraySet1(_888_cf2, 13)
			(_nw164).ArraySet1(_890_cf0, 14)
			(_nw164).ArraySet1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), 15)
			(_nw164).ArraySet1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), 16)
			(_nw164).ArraySet1(Companion_Default___.Fm3(globalState), 17)
			(_nw164).ArraySet1((_this).F7(), 18)
			_893_v14 = _nw164
			var _895_v15 _dafny.Array
			_ = _895_v15
			var _nw165 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
			_ = _nw165
			_895_v15 = _nw165
			var _896_v16 D1
			_ = _896_v16
			_896_v16 = Companion_D1_.Create_DC3_(_895_v15)
			var _897_v17 _dafny.Array
			_ = _897_v17
			var _nwElement0_33 D1 = _896_v16
			_ = _nwElement0_33
			var _nw166 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_33, nil, _dafny.IntOfInt64(5))
			_ = _nw166
			(_nw166).ArraySet1(_nwElement0_33, 0)
			(_nw166).ArraySet1(_896_v16, 1)
			(_nw166).ArraySet1(Companion_D1_.Create_DC3_(_895_v15), 2)
			(_nw166).ArraySet1(_896_v16, 3)
			(_nw166).ArraySet1(Companion_D1_.Create_DC3_(_895_v15), 4)
			_897_v17 = _nw166
			var _898_v18 *C2
			_ = _898_v18
			var _nw167 *C2 = New_C2_()
			_ = _nw167
			_nw167.Ctor__(_893_v14, _897_v17)
			_898_v18 = _nw167
			if ((_dafny.IntOfUint32(((_this).F5()).Cardinality())).Minus(_dafny.IntOfInt64(-466))).Cmp((_dafny.Zero).Minus((_this).F8())) > 0 {
				var _899_v19 *C0
				_ = _899_v19
				var _nw168 *C0 = New_C0_()
				_ = _nw168
				_nw168.Ctor__(!(Companion_Default___.Fm7(_this.F4(), _887_cf3, !(true), globalState)), p3, _this.F4(), (_this).F5())
				_899_v19 = _nw168
				var _900_v20 _dafny.Array
				_ = _900_v20
				var _nwElement0_34 *C0 = (func() *C0 {
					if (_this).F6() {
						return _899_v19
					}
					return _899_v19
				})()
				_ = _nwElement0_34
				var _nw169 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_34, nil, _dafny.IntOfInt64(18))
				_ = _nw169
				(_nw169).ArraySet1(_nwElement0_34, 0)
				(_nw169).ArraySet1(_899_v19, 1)
				(_nw169).ArraySet1(_899_v19, 2)
				(_nw169).ArraySet1(_899_v19, 3)
				(_nw169).ArraySet1(_899_v19, 4)
				(_nw169).ArraySet1(_899_v19, 5)
				(_nw169).ArraySet1((func() *C0 {
					if !(_890_cf0) {
						return _899_v19
					}
					return _899_v19
				})(), 6)
				(_nw169).ArraySet1(_899_v19, 7)
				(_nw169).ArraySet1(_899_v19, 8)
				(_nw169).ArraySet1(_899_v19, 9)
				(_nw169).ArraySet1(_899_v19, 10)
				(_nw169).ArraySet1(_899_v19, 11)
				(_nw169).ArraySet1(_899_v19, 12)
				(_nw169).ArraySet1(_899_v19, 13)
				(_nw169).ArraySet1(_899_v19, 14)
				(_nw169).ArraySet1(_899_v19, 15)
				(_nw169).ArraySet1(_899_v19, 16)
				(_nw169).ArraySet1(_899_v19, 17)
				_900_v20 = _nw169
				var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_900_v20), 0))
				_ = _index126
				(_900_v20).ArraySet1(_899_v19, (_index126).Int())
				var _901_v21 _dafny.Sequence
				_ = _901_v21
				_901_v21 = _dafny.SeqOf(_899_v19)
				var _902_v22 _dafny.Sequence
				_ = _902_v22
				_902_v22 = _dafny.SeqOf(_dafny.SeqOf(true, false), _891_v12)
				var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_900_v20), 0))
				_ = _index127
				var _index128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index128
				var _rhs117 _dafny.Int = ((_this).F8()).Minus(_887_cf3)
				_ = _rhs117
				var _rhs118 *C0 = (_901_v21).Select((Companion_Default___.SafeIndex(((_this).F8()).Times((_899_v19).Fm9(_889_cf1, (_902_v22).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_902_v22).Cardinality()))).Uint32()).(_dafny.Sequence), globalState)), _dafny.IntOfUint32((_901_v21).Cardinality()))).Uint32()).(*C0)
				_ = _rhs118
				var _rhs119 bool = !_dafny.Companion_Sequence_.Contains(_dafny.Companion_Sequence_.Concatenate(_891_v12, _891_v12), (_this).F7())
				_ = _rhs119
				var _lhs86 _dafny.Array = _900_v20
				_ = _lhs86
				var _lhs87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_900_v20), 0))
				_ = _lhs87
				var _lhs88 _dafny.Array = (_this).F10()
				_ = _lhs88
				var _lhs89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs89
				_887_cf3 = _rhs117
				(_lhs86).ArraySet1(_rhs118, (_lhs87).Int())
				(_lhs88).ArraySet1(_rhs119, (_lhs89).Int())
				var _903_v23 _dafny.Sequence
				_ = _903_v23
				_903_v23 = _dafny.UnicodeSeqOfUtf8Bytes("muet")
				var _904_v24 _dafny.Map
				_ = _904_v24
				_904_v24 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_890_cf0), false)
				_903_v23 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate((_this).F5(), _903_v23), _dafny.Companion_Sequence_.Concatenate((_this).F5(), Companion_Default___.Fm14((_this).Fm1(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), _904_v24, !(_890_cf0), globalState), (_this).F8(), false, p1, globalState)))
				var _index129 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_900_v20), 0))
				_ = _index129
				(_900_v20).ArraySet1((_900_v20).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(792), _dafny.ArrayLen((_900_v20), 0))).Int()).(*C0), (_index129).Int())
				var _index130 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index130
				((_this).F10()).ArraySet1(p1, (_index130).Int())
				(globalState).F1 = _889_cf1
			} else {
				var _905_v25 _dafny.Array
				_ = _905_v25
				var _nw170 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(7))
				_ = _nw170
				_905_v25 = _nw170
				var _906_v26 _dafny.Map
				_ = _906_v26
				_906_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F9())
				var _index131 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index131
				var _rhs120 bool = _890_cf0
				_ = _rhs120
				var _rhs121 _dafny.Array = _905_v25
				_ = _rhs121
				var _rhs122 _dafny.Map = ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), p0)).Merge(_906_v26)).Merge((_906_v26).Merge((_906_v26).Update((_this).F7(), (_this).F7())))
				_ = _rhs122
				var _rhs123 _dafny.Int = Companion_Default___.SafeDivisionInt(_889_cf1, _dafny.IntOfUint32((_dafny.SeqOf(_887_cf3)).Cardinality()))
				_ = _rhs123
				var _rhs124 _dafny.Int = (_887_cf3).Minus(_889_cf1)
				_ = _rhs124
				var _lhs90 _dafny.Array = (_this).F10()
				_ = _lhs90
				var _lhs91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs91
				var _lhs92 *GlobalState = globalState
				_ = _lhs92
				var _lhs93 *GlobalState = globalState
				_ = _lhs93
				(_lhs90).ArraySet1(_rhs120, (_lhs91).Int())
				_905_v25 = _rhs121
				_906_v26 = _rhs122
				_lhs92.F1 = _rhs123
				_lhs93.F1 = _rhs124
				var _index132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(385), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index132
				((_this).F10()).ArraySet1(_888_cf2, (_index132).Int())
				var _907_v27 _dafny.Sequence
				_ = _907_v27
				_907_v27 = _dafny.UnicodeSeqOfUtf8Bytes("vsc")
				_907_v27 = _dafny.Companion_Sequence_.Update(_907_v27, (Companion_Default___.SafeIndex(_889_cf1, _dafny.IntOfUint32((_907_v27).Cardinality()))).Uint32(), _this.F4())
				_887_cf3 = _887_cf3
				var _908_v28 D1
				_ = _908_v28
				_908_v28 = Companion_D1_.Create_DC5_(_889_cf1)
				var _909_v29 _dafny.Sequence
				_ = _909_v29
				_909_v29 = _dafny.SeqOf(_908_v28, _908_v28)
				var _910_v30 _dafny.Sequence
				_ = _910_v30
				_910_v30 = _dafny.Companion_Sequence_.Concatenate(_909_v29, _909_v29)
				_910_v30 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Update(_909_v29, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.IntOfUint32((_909_v29).Cardinality()))).Uint32(), _908_v28), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_909_v29, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(635), _dafny.IntOfUint32((_909_v29).Cardinality()))).Uint32(), _908_v28)).Cardinality()))).Uint32(), Companion_Default___.Fm25(globalState))
			}
		} else {
			var _911___mcc_h4 _dafny.Int = _source9.Get_().(D0_DC2).Cf4
			_ = _911___mcc_h4
			var _912___mcc_h5 bool = _source9.Get_().(D0_DC2).Cf5
			_ = _912___mcc_h5
			var _913___mcc_h6 _dafny.Sequence = _source9.Get_().(D0_DC2).Cf6
			_ = _913___mcc_h6
			var _914_cf6 _dafny.Sequence = _913___mcc_h6
			_ = _914_cf6
			var _915_cf5 bool = _912___mcc_h5
			_ = _915_cf5
			var _916_cf4 _dafny.Int = _911___mcc_h4
			_ = _916_cf4
			var _917_v31 _dafny.Sequence
			_ = _917_v31
			_917_v31 = _dafny.UnicodeSeqOfUtf8Bytes("eyth")
			var _918_v32 _dafny.MultiSet
			_ = _918_v32
			_918_v32 = _dafny.MultiSetOf(_dafny.IntOfInt64(555), (_this).F8())
			_917_v31 = (func() _dafny.Sequence {
				if p0 {
					return _917_v31
				}
				return Companion_Default___.Fm14((_this).F8(), (func() _dafny.Int {
					if (_918_v32).Contains((_this).F8()) {
						return (_918_v32).Multiplicity((_this).F8())
					}
					return (_this).F8()
				})(), _915_cf5, true, globalState)
			})()
			var _919_v33 _dafny.Array
			_ = _919_v33
			var _nw171 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.One)
			_ = _nw171
			_919_v33 = _nw171
			var _index133 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_919_v33), 0))
			_ = _index133
			(_919_v33).ArraySet1(Companion_Default___.Fm14(_916_cf4, (_this).F8(), (_this).F7(), p0, globalState), (_index133).Int())
			var _index134 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_919_v33), 0))
			_ = _index134
			(_919_v33).ArraySet1(_dafny.Companion_Sequence_.Update((_this).F5(), (Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_917_v31).Cardinality()), (p2).Cardinality()), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32(), _this.F4()), (_index134).Int())
			if (((func() _dafny.Map {
				var _coll18 = _dafny.NewMapBuilder()
				_ = _coll18
				for _iter21 := _dafny.Iterate((_918_v32).Elements()); ; {
					_compr_18, _ok21 := _iter21()
					if !_ok21 {
						break
					}
					var _920_v34 _dafny.Int
					_920_v34 = interface{}(_compr_18).(_dafny.Int)
					if (_918_v32).Contains(_920_v34) {
						_coll18.Add(Companion_Default___.SafeDivisionInt(_920_v34, _dafny.IntOfUint32((_dafny.SeqOf((_this).F9(), true)).Cardinality())), _this.F4())
					}
				}
				return _coll18.ToMap()
			}()).Cardinality()).Times((_this).F8())).Cmp((_this).F8()) != 0 {
				_915_cf5 = p0
				r1 = (_this).F8()
				var _921_v35 D1
				_ = _921_v35
				_921_v35 = Companion_D1_.Create_DC5_(_916_cf4)
				var _922_v36 _dafny.Sequence
				_ = _922_v36
				_922_v36 = _dafny.SeqOf(_921_v35)
				_922_v36 = _dafny.SeqOf(Companion_D1_.Create_DC5_(_916_cf4), _921_v35, _921_v35)
				_915_cf5 = ((func() bool {
					if (_this).F6() {
						return !(p0)
					}
					return (_this).F7()
				})()) || (!(_915_cf5))
				var _923_v37 _dafny.Array
				_ = _923_v37
				var _nw172 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(25))
				_ = _nw172
				_923_v37 = _nw172
				var _index135 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_923_v37), 0))
				_ = _index135
				(_923_v37).ArraySet1(_918_v32, (_index135).Int())
				var _index136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(695), _dafny.ArrayLen((_923_v37), 0))
				_ = _index136
				(_923_v37).ArraySet1(_918_v32, (_index136).Int())
			} else {
				var _924_v38 _dafny.Map
				_ = _924_v38
				_924_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.SetOf(_915_cf5)).Cardinality(), _dafny.CodePoint('k'))
				(_this).F4_set_((func() _dafny.CodePoint {
					if (_924_v38).Contains(Companion_Default___.SafeModuloInt(_916_cf4, _dafny.IntOfUint32(((_this).F5()).Cardinality()))) {
						return (_924_v38).Get(Companion_Default___.SafeModuloInt(_916_cf4, _dafny.IntOfUint32(((_this).F5()).Cardinality()))).(_dafny.CodePoint)
					}
					return _this.F4()
				})())
				(globalState).F1 = (_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_dafny.Zero).Minus((_916_cf4).Minus((_this).F8())))))
				var _925_v39 *C1
				_ = _925_v39
				var _nw173 *C1 = New_C1_()
				_ = _nw173
				_nw173.Ctor__(_dafny.IntOfUint32((_dafny.SeqOf(p1)).Cardinality()), !(p1), Companion_Default___.Fm3(globalState), p1, Companion_Default___.Fm13(globalState), (_919_v33).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(891), _dafny.ArrayLen((_919_v33), 0))).Int()).(_dafny.Sequence))
				_925_v39 = _nw173
				var _index137 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index137
				((_this).F10()).ArraySet1(p1, (_index137).Int())
				var _index138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index138
				((_this).F10()).ArraySet1(p3, (_index138).Int())
				var _926_v40 _dafny.Array
				_ = _926_v40
				var _nw174 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(8))
				_ = _nw174
				_926_v40 = _nw174
				var _927_v41 _dafny.Set
				_ = _927_v41
				_927_v41 = _dafny.SetOf(_926_v40)
				_915_cf5 = !((_927_v41).IsDisjointFrom(_927_v41))
			}
			var _928_v42 _dafny.Map
			_ = _928_v42
			_928_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() _dafny.Int {
				if (_this).F7() {
					return _916_cf4
				}
				return _916_cf4
			})(), (_this).F8())
			var _929_v43 *C3
			_ = _929_v43
			var _nw175 *C3 = New_C3_()
			_ = _nw175
			_nw175.Ctor__(p3, p3, _dafny.IntOfInt64(33), p0, p3, p0, _this.F4(), _dafny.UnicodeSeqOfUtf8Bytes("qkhm"))
			_929_v43 = _nw175
			var _930_v44 _dafny.MultiSet
			_ = _930_v44
			_930_v44 = _dafny.MultiSetOf(_929_v43)
			_916_cf4 = (_dafny.Zero).Minus((func() _dafny.Int {
				if (_928_v42).Contains((_dafny.IntOfInt64(837)).Plus((_930_v44).Cardinality())) {
					return (_928_v42).Get((_dafny.IntOfInt64(837)).Plus((_930_v44).Cardinality())).(_dafny.Int)
				}
				return _916_cf4
			})())
		}
		if false {
			var _931_v45 bool
			_ = _931_v45
			_931_v45 = true
			_931_v45 = !((_this).F7())
			var _932_v46 _dafny.Map
			_ = _932_v46
			_932_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_931_v45, true)
			var _933_v47 _dafny.Map
			_ = _933_v47
			_933_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm1(p3, _932_v46, (_this).F6(), globalState))
			_931_v45 = ((_this).F8()).Cmp((func() _dafny.Int {
				if (_933_v47).Contains(_931_v45) {
					return (_933_v47).Get(_931_v45).(_dafny.Int)
				}
				return _dafny.IntOfUint32(((_this).F5()).Cardinality())
			})()) == 0
			(globalState).F1 = ((func() _dafny.Int {
				if p0 {
					return (_this).F8()
				}
				return (_this).F8()
			})()).Times((_this).F8())
			_931_v45 = !((_this).F6())
			var _934_v48 *C3
			_ = _934_v48
			var _nw176 *C3 = New_C3_()
			_ = _nw176
			_nw176.Ctor__((_this).F7(), p0, (_this).F8(), p0, _931_v45, p0, _this.F4(), _dafny.Companion_Sequence_.Concatenate((_this).F5(), (_this).F5()))
			_934_v48 = _nw176
		} else {
			if Companion_Default___.Fm3(globalState) {
				var _935_v49 _dafny.Set
				_ = _935_v49
				_935_v49 = _dafny.SetOf((_this).F8(), (func() _dafny.Int {
					if (_this).F7() {
						return (_this).F8()
					}
					return (_this).F8()
				})())
				_935_v49 = func() _dafny.Set {
					var _coll19 = _dafny.NewBuilder()
					_ = _coll19
					for _iter22 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(190), _dafny.IntOfInt64(516))); ; {
						_compr_19, _ok22 := _iter22()
						if !_ok22 {
							break
						}
						var _936_v50 _dafny.Int
						_936_v50 = interface{}(_compr_19).(_dafny.Int)
						if ((_dafny.IntOfInt64(190)).Cmp(_936_v50) <= 0) && ((_936_v50).Cmp(_dafny.IntOfInt64(516)) < 0) {
							_coll19.Add((_936_v50).Plus((p2).Cardinality()))
						}
					}
					return _coll19.ToSet()
				}()
				var _937_v51 bool
				_ = _937_v51
				_937_v51 = true
				var _rhs125 _dafny.CodePoint = _this.F4()
				_ = _rhs125
				var _rhs126 bool = (_this).F7()
				_ = _rhs126
				var _lhs94 *C5 = _this
				_ = _lhs94
				_lhs94.F4_set_(_rhs125)
				_937_v51 = _rhs126
				r1 = (_this).F8()
				var _index139 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index139
				((_this).F11()).ArraySet1((_this).F8(), (_index139).Int())
				var _index140 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(107), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index140
				((_this).F11()).ArraySet1((_this).F8(), (_index140).Int())
				var _938_v52 D3
				_ = _938_v52
				_938_v52 = Companion_D3_.Create_DC10_((_this).F5())
				var _939_v53 _dafny.MultiSet
				_ = _939_v53
				_939_v53 = _dafny.MultiSetOf((_this).F5(), _dafny.Companion_Sequence_.Concatenate((_938_v52).Dtor_cf19(), (_this).F5()), (_this).F5())
				_939_v53 = _939_v53
			} else {
				r1 = (Companion_Default___.SafeModuloInt(_dafny.IntOfUint32(((_this).F5()).Cardinality()), _dafny.IntOfInt64(134))).Plus(Companion_Default___.SafeModuloInt((_this).F8(), (_this).F8()))
				var _index141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index141
				((_this).F11()).ArraySet1((_this).F8(), (_index141).Int())
				var _index142 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index142
				((_this).F11()).ArraySet1(((_this).F8()).Plus(((_this).F8()).Plus((_this).F8())), (_index142).Int())
				var _940_v54 bool
				_ = _940_v54
				_940_v54 = false
				var _941_v55 _dafny.Set
				_ = _941_v55
				_941_v55 = _dafny.SetOf(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int))
				var _942_v56 _dafny.Sequence
				_ = _942_v56
				_942_v56 = _dafny.SeqOf(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), (_this).F8(), (_941_v55).Cardinality())
				var _943_v57 _dafny.Sequence
				_ = _943_v57
				_943_v57 = _dafny.SeqOf(Companion_Default___.Fm3(globalState))
				var _944_v58 _dafny.Set
				_ = _944_v58
				_944_v58 = _dafny.SetOf(_942_v56)
				var _index143 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index143
				var _rhs127 bool = !((func() bool {
					if !(Companion_Default___.Fm7(_dafny.CodePoint('u'), ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), p0, globalState)) {
						return (_this).F9()
					}
					return ((_942_v56).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_942_v56).Cardinality()))).Uint32()).(_dafny.Int)).Cmp(_dafny.IntOfInt64(-550)) < 0
				})())
				_ = _rhs127
				var _rhs128 bool = Companion_Default___.Fm7(_this.F4(), (_dafny.Zero).Minus((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_943_v57, _dafny.Companion_Sequence_.Update(_943_v57, (Companion_Default___.SafeIndex(_dafny.IntOfUint32((_943_v57).Cardinality()), _dafny.IntOfUint32((_943_v57).Cardinality()))).Uint32(), !((_this).F6())))).Cardinality()))), !((_this).F6()), globalState)
				_ = _rhs128
				var _rhs129 _dafny.Int = (_this).F8()
				_ = _rhs129
				var _rhs130 _dafny.Int = ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int)
				_ = _rhs130
				var _rhs131 bool = (_944_v58).Equals(_944_v58)
				_ = _rhs131
				var _lhs95 *GlobalState = globalState
				_ = _lhs95
				var _lhs96 _dafny.Array = (_this).F11()
				_ = _lhs96
				var _lhs97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(694), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _lhs97
				_940_v54 = _rhs127
				_940_v54 = _rhs128
				_lhs95.F1 = _rhs129
				(_lhs96).ArraySet1(_rhs130, (_lhs97).Int())
				_940_v54 = _rhs131
				_940_v54 = (_941_v55).IsSubsetOf(_941_v55)
				_940_v54 = p0
			}
			var _945_v59 _dafny.Sequence
			_ = _945_v59
			_945_v59 = _dafny.SeqOf((_this).F9(), (_this).F9(), (_this).F6(), (_this).F6())
			if (_945_v59).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_945_v59).Cardinality()))).Uint32()).(bool) {
				Companion_Default___.M0(globalState)
				var _946_v60 _dafny.Map
				_ = _946_v60
				_946_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (_this).F7())
				var _947_v61 _dafny.Map
				_ = _947_v61
				_947_v61 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), p1)
				var _948_v62 T2
				_ = _948_v62
				var _nw177 *C3 = New_C3_()
				_ = _nw177
				_nw177.Ctor__((_this).F7(), (_this).F7(), (_this).F8(), p1, false, p3, _this.F4(), (_this).F5())
				_948_v62 = _nw177
				var _949_v63 _dafny.MultiSet
				_ = _949_v63
				_949_v63 = _dafny.MultiSetOf((_this).Fm1((_this).F7(), _946_v60, (_this).F7(), globalState), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, _948_v62)).Cardinality())
				var _950_v64 _dafny.Sequence
				_ = _950_v64
				_950_v64 = _dafny.SeqOf((_this).F8(), (_949_v63).Cardinality())
				var _951_v65 _dafny.Map
				_ = _951_v65
				_951_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p1, _dafny.IntOfUint32((_950_v64).Cardinality()))
				var _952_v66 _dafny.Map
				_ = _952_v66
				_952_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (p2).Cardinality())
				var _953_v67 _dafny.Sequence
				_ = _953_v67
				_953_v67 = _dafny.SeqOf(_946_v60)
				var _954_v68 _dafny.Map
				_ = _954_v68
				_954_v68 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(237), (_953_v67).Select((Companion_Default___.SafeIndex((_948_v62).F8(), _dafny.IntOfUint32((_953_v67).Cardinality()))).Uint32()).(_dafny.Map))
				var _955_v69 _dafny.Array
				_ = _955_v69
				var _nwElement0_35 _dafny.Int = (_this).F8()
				_ = _nwElement0_35
				var _nw178 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_35, nil, _dafny.IntOfInt64(29))
				_ = _nw178
				(_nw178).ArraySet1(_nwElement0_35, 0)
				(_nw178).ArraySet1(((_this).Fm1(p0, _946_v60, p3, globalState)).Plus((_this).F8()), 1)
				(_nw178).ArraySet1((_this).F8(), 2)
				(_nw178).ArraySet1((_this).F8(), 3)
				(_nw178).ArraySet1((_this).F8(), 4)
				(_nw178).ArraySet1((_this).F8(), 5)
				(_nw178).ArraySet1((_this).F8(), 6)
				(_nw178).ArraySet1((_this).F8(), 7)
				(_nw178).ArraySet1(Companion_Default___.SafeDivisionInt((_this).F8(), (_947_v61).Cardinality()), 8)
				(_nw178).ArraySet1((_this).F8(), 9)
				(_nw178).ArraySet1(((_this).F8()).Plus((_951_v65).Cardinality()), 10)
				(_nw178).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((_this).F5(), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32(), _948_v62.F4()), (_this).F5())).Cardinality()), 11)
				(_nw178).ArraySet1((_this).F8(), 12)
				(_nw178).ArraySet1((_dafny.Zero).Minus((_948_v62).F8()), 13)
				(_nw178).ArraySet1((_948_v62).F8(), 14)
				(_nw178).ArraySet1((func() _dafny.Int {
					if (_952_v66).Contains((_948_v62).F8()) {
						return (_952_v66).Get((_948_v62).F8()).(_dafny.Int)
					}
					return _dafny.IntOfInt64(337)
				})(), 15)
				(_nw178).ArraySet1((_this).Fm1((_this).F6(), _946_v60, (_this).F7(), globalState), 16)
				(_nw178).ArraySet1((_this).F8(), 17)
				(_nw178).ArraySet1(((_this).F8()).Times((_this).F8()), 18)
				(_nw178).ArraySet1((((func() _dafny.Map {
					if (_954_v68).Contains((_this).F8()) {
						return (_954_v68).Get((_this).F8()).(_dafny.Map)
					}
					return _946_v60
				})()).Update(!(!((_this).F6())), (_948_v62).F7())).Cardinality(), 19)
				(_nw178).ArraySet1((_948_v62).F8(), 20)
				(_nw178).ArraySet1(((p2).Difference(_dafny.MultiSetFromSeq(_945_v59))).Cardinality(), 21)
				(_nw178).ArraySet1((_948_v62).F8(), 22)
				(_nw178).ArraySet1((_this).F8(), 23)
				(_nw178).ArraySet1(((_this).F8()).Times((_948_v62).F8()), 24)
				(_nw178).ArraySet1(((_951_v65).Merge(_951_v65)).Cardinality(), 25)
				(_nw178).ArraySet1((_dafny.Zero).Minus((_this).F8()), 26)
				(_nw178).ArraySet1((_this).F8(), 27)
				(_nw178).ArraySet1((_948_v62).F8(), 28)
				_955_v69 = _nw178
				_955_v69 = _955_v69
				var _956_v70 _dafny.Map
				_ = _956_v70
				_956_v70 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_948_v62).F8(), Companion_Default___.Fm26(p3, (_948_v62).F8(), _948_v62.F4(), globalState))
				var _957_v71 _dafny.Map
				_ = _957_v71
				_957_v71 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), _956_v70)
				var _958_v72 _dafny.Array
				_ = _958_v72
				var _nwElement0_36 _dafny.Map = (_957_v71).Update(_948_v62.F4(), _956_v70)
				_ = _nwElement0_36
				var _nw179 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_36, nil, _dafny.IntOfInt64(7))
				_ = _nw179
				(_nw179).ArraySet1(_nwElement0_36, 0)
				(_nw179).ArraySet1(_957_v71, 1)
				(_nw179).ArraySet1((_957_v71).Update(_this.F4(), _956_v70), 2)
				(_nw179).ArraySet1(_957_v71, 3)
				(_nw179).ArraySet1(_957_v71, 4)
				(_nw179).ArraySet1(_957_v71, 5)
				(_nw179).ArraySet1(_957_v71, 6)
				_958_v72 = _nw179
				var _index144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_958_v72), 0))
				_ = _index144
				(_958_v72).ArraySet1(_957_v71, (_index144).Int())
				var _index145 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(398), _dafny.ArrayLen((_958_v72), 0))
				_ = _index145
				(_958_v72).ArraySet1((_957_v71).Merge(_957_v71), (_index145).Int())
				var _959_v73 _dafny.Array
				_ = _959_v73
				var _nw180 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMultiSet, _dafny.IntOfInt64(29))
				_ = _nw180
				_959_v73 = _nw180
				_959_v73 = _959_v73
				var _960_v74 _dafny.Map
				_ = _960_v74
				_960_v74 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F9())
				var _961_v75 _dafny.Map
				_ = _961_v75
				_961_v75 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_960_v74, p1)
				var _index146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index146
				((_this).F11()).ArraySet1((_948_v62).F8(), (_index146).Int())
				var _index147 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index147
				var _rhs132 _dafny.Int = Companion_Default___.Fm8(Companion_Default___.SafeDivisionInt((_this).F8(), (_this).F8()), false, ((_948_v62).F8()).Cmp((_948_v62).F8()) >= 0, globalState)
				_ = _rhs132
				var _rhs133 _dafny.Int = (_948_v62).F8()
				_ = _rhs133
				var _rhs134 _dafny.Map = _961_v75
				_ = _rhs134
				var _rhs135 _dafny.Int = Companion_Default___.SafeModuloInt((_this).Fm1(p3, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_948_v62).F9(), (_948_v62).F6()), (_this).F9(), globalState), (_948_v62).F8())
				_ = _rhs135
				var _rhs136 _dafny.Int = Companion_Default___.SafeModuloInt(((_this).F8()).Times((_948_v62).F8()), (_this).F8())
				_ = _rhs136
				var _lhs98 *GlobalState = globalState
				_ = _lhs98
				var _lhs99 _dafny.Array = (_this).F11()
				_ = _lhs99
				var _lhs100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(223), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _lhs100
				r0 = _rhs132
				_lhs98.F1 = _rhs133
				_961_v75 = _rhs134
				(_lhs99).ArraySet1(_rhs135, (_lhs100).Int())
				r0 = _rhs136
			} else {
				var _962_v76 bool
				_ = _962_v76
				_962_v76 = true
				_962_v76 = true
				_962_v76 = !(_dafny.Companion_Sequence_.IsProperPrefixOf((_this).F5(), (_this).F5()))
				var _index148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index148
				((_this).F10()).ArraySet1(p1, (_index148).Int())
				var _index149 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index149
				((_this).F10()).ArraySet1(false, (_index149).Int())
				var _963_v77 *C0
				_ = _963_v77
				var _nw181 *C0 = New_C0_()
				_ = _nw181
				_nw181.Ctor__(_962_v76, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), _this.F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(38))).Uint32(), func(coer52 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg52 _dafny.Int) interface{} {
						return coer52(arg52)
					}
				}(func(_964_i2 _dafny.Int) _dafny.CodePoint {
					return _this.F4()
				})))
				_963_v77 = _nw181
				var _965_v78 _dafny.Map
				_ = _965_v78
				_965_v78 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_963_v77, _962_v76)
				var _966_v79 _dafny.Map
				_ = _966_v79
				_966_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (func() bool {
					if (_965_v78).Contains(_963_v77) {
						return (_965_v78).Get(_963_v77).(bool)
					}
					return !((_this).F7())
				})())
				(_this).M4((_this).F8(), _966_v79, globalState)
				var _index150 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(463), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index150
				((_this).F10()).ArraySet1(!((_945_v59).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_945_v59).Cardinality()))).Uint32()).(bool)), (_index150).Int())
			}
			(_this).F4_set_(_this.F4())
			var _967_v80 bool
			_ = _967_v80
			_967_v80 = false
			_967_v80 = p0
			var _968_v81 _dafny.Sequence
			_ = _968_v81
			_968_v81 = _dafny.SeqOf((_this).F8(), (_dafny.Zero).Minus((_this).F8()))
			var _969_v82 _dafny.Set
			_ = _969_v82
			_969_v82 = _dafny.SetOf((_this).F6(), (_this).F7())
			var _970_v83 D3
			_ = _970_v83
			_970_v83 = Companion_D3_.Create_DC11_(false, _dafny.IntOfUint32((_968_v81).Cardinality()), _969_v82)
			var _971_v84 _dafny.Sequence
			_ = _971_v84
			_971_v84 = _dafny.SeqOf((_970_v83).Dtor_cf22())
			var _972_v85 D0
			_ = _972_v85
			_972_v85 = Companion_D0_.Create_DC1_(_967_v80, _dafny.IntOfInt64(-224), (_this).F6(), (_this).F8())
			(globalState).F1 = (_dafny.IntOfUint32((_971_v84).Cardinality())).Times(Companion_Default___.Fm8((_972_v85).Dtor_cf3(), !(p3), p1, globalState))
		}
		var _973_v86 _dafny.MultiSet
		_ = _973_v86
		_973_v86 = _dafny.MultiSetOf((_this).F8(), (_this).F8())
		var _974_v87 D6
		_ = _974_v87
		_974_v87 = Companion_D6_.Create_DC16_(_973_v86)
		r0 = (_dafny.Zero).Minus(((_974_v87).Dtor_cf30()).Cardinality())
		r1 = (_dafny.Zero).Minus((_this).F8())
		if (p2).IsDisjointFrom(p2) {
			var _975_v88 _dafny.Array
			_ = _975_v88
			var _nw182 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(11))
			_ = _nw182
			_975_v88 = _nw182
			var _976_v89 _dafny.Set
			_ = _976_v89
			_976_v89 = _dafny.SetOf(p0, true)
			var _index151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_975_v88), 0))
			_ = _index151
			(_975_v88).ArraySet1((func() _dafny.Set {
				if (_this).F6() {
					return Companion_Default___.Fm0((_this).F9(), _dafny.IntOfInt64(924), (p2).Cardinality(), p3, globalState)
				}
				return _976_v89
			})(), (_index151).Int())
			var _977_v90 bool
			_ = _977_v90
			_977_v90 = false
			var _978_v91 T3
			_ = _978_v91
			var _nw183 *C4 = New_C4_()
			_ = _nw183
			_nw183.Ctor__(_dafny.IntOfInt64(644), p3, !((_this).F6()), false, _this.F4(), (_this).F5())
			_978_v91 = _nw183
			var _979_v92 _dafny.Map
			_ = _979_v92
			_979_v92 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_978_v91, (_this).F11())
			var _980_v93 _dafny.Array
			_ = _980_v93
			var _nwElement0_37 _dafny.Array = (_this).F11()
			_ = _nwElement0_37
			var _nw184 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_37, nil, _dafny.IntOfInt64(27))
			_ = _nw184
			(_nw184).ArraySet1(_nwElement0_37, 0)
			(_nw184).ArraySet1((_this).F11(), 1)
			(_nw184).ArraySet1((_this).F11(), 2)
			(_nw184).ArraySet1((_this).F11(), 3)
			(_nw184).ArraySet1((_this).F11(), 4)
			(_nw184).ArraySet1((_this).F11(), 5)
			(_nw184).ArraySet1((_this).F11(), 6)
			(_nw184).ArraySet1((_this).F11(), 7)
			(_nw184).ArraySet1((_this).F11(), 8)
			(_nw184).ArraySet1((_this).F11(), 9)
			(_nw184).ArraySet1((_this).F11(), 10)
			(_nw184).ArraySet1((_this).F11(), 11)
			(_nw184).ArraySet1((_this).F11(), 12)
			(_nw184).ArraySet1((_this).F11(), 13)
			(_nw184).ArraySet1((_this).F11(), 14)
			(_nw184).ArraySet1((_this).F11(), 15)
			(_nw184).ArraySet1((_this).F11(), 16)
			(_nw184).ArraySet1((_this).F11(), 17)
			(_nw184).ArraySet1((_this).F11(), 18)
			(_nw184).ArraySet1((_this).F11(), 19)
			(_nw184).ArraySet1((_this).F11(), 20)
			(_nw184).ArraySet1((_this).F11(), 21)
			(_nw184).ArraySet1((func() _dafny.Array {
				if p0 {
					return (_this).F11()
				}
				return (_this).F11()
			})(), 22)
			(_nw184).ArraySet1((_this).F11(), 23)
			(_nw184).ArraySet1((_this).F11(), 24)
			(_nw184).ArraySet1((func() _dafny.Array {
				if (_979_v92).Contains(_978_v91) {
					return (_979_v92).Get(_978_v91).(_dafny.Array)
				}
				return (_this).F11()
			})(), 25)
			(_nw184).ArraySet1((_this).F11(), 26)
			_980_v93 = _nw184
			var _index152 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))
			_ = _index152
			(_980_v93).ArraySet1((_this).F11(), (_index152).Int())
			var _981_v94 *C0
			_ = _981_v94
			var _nw185 *C0 = New_C0_()
			_ = _nw185
			_nw185.Ctor__(p1, p1, _978_v91.F4(), _dafny.Companion_Sequence_.Concatenate((_978_v91).F5(), (_this).F5()))
			_981_v94 = _nw185
			var _index153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_975_v88), 0))
			_ = _index153
			var _index154 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))
			_ = _index154
			var _rhs137 _dafny.Set = _976_v89
			_ = _rhs137
			var _rhs138 _dafny.CodePoint = _978_v91.F4()
			_ = _rhs138
			var _rhs139 bool = !((_this).F7())
			_ = _rhs139
			var _rhs140 _dafny.Array = (func() _dafny.Array {
				if ((_this).F8()).Cmp((_this).F8()) >= 0 {
					return (_this).F11()
				}
				return (_this).F11()
			})()
			_ = _rhs140
			var _rhs141 *C0 = _981_v94
			_ = _rhs141
			var _lhs101 _dafny.Array = _975_v88
			_ = _lhs101
			var _lhs102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(828), _dafny.ArrayLen((_975_v88), 0))
			_ = _lhs102
			var _lhs103 *C5 = _this
			_ = _lhs103
			var _lhs104 _dafny.Array = _980_v93
			_ = _lhs104
			var _lhs105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))
			_ = _lhs105
			(_lhs101).ArraySet1(_rhs137, (_lhs102).Int())
			_lhs103.F4_set_(_rhs138)
			_977_v90 = _rhs139
			(_lhs104).ArraySet1(_rhs140, (_lhs105).Int())
			_981_v94 = _rhs141
			var _arr0 _dafny.Array = _dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))
			_ = _arr0
			var _index155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))
			_ = _index155
			(_arr0).ArraySet1((_this).F8(), (_index155).Int())
			var _982_v95 _dafny.Map
			_ = _982_v95
			_982_v95 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_973_v86, (_978_v91).F9())
			var _arr1 _dafny.Array = _dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))
			_ = _arr1
			var _index156 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))
			_ = _index156
			(_arr1).ArraySet1(((_982_v95).Merge(_982_v95)).Cardinality(), (_index156).Int())
			(globalState).F1 = ((_this).F8()).Times((_978_v91).F8())
			if (_978_v91).F9() {
				(_978_v91).F4_set_(_this.F4())
				var _983_v96 D1
				_ = _983_v96
				_983_v96 = Companion_D1_.Create_DC4_(_978_v91.F4(), (_this).F8(), _978_v91.F4(), (_this).F8(), (_978_v91).F6())
				var _984_v97 _dafny.Sequence
				_ = _984_v97
				_984_v97 = _dafny.SeqOf((_this).F8())
				var _985_v98 *C4
				_ = _985_v98
				var _nw186 *C4 = New_C4_()
				_ = _nw186
				_nw186.Ctor__((_dafny.Zero).Minus((_978_v91).F8()), ((_978_v91).F9()) && (false), (_983_v96).Dtor_cf12(), (_973_v86).IsSubsetOf(_dafny.MultiSetOf((_978_v91).F8(), (_this).F8())), ((_978_v91).F5()).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_984_v97).Cardinality()), _dafny.IntOfUint32(((_978_v91).F5()).Cardinality()))).Uint32()).(_dafny.CodePoint), (_this).F5())
				_985_v98 = _nw186
				var _986_v99 _dafny.Sequence
				_ = _986_v99
				_986_v99 = _dafny.SeqOf((_this).F10())
				var _index157 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index157
				((_this).F10()).ArraySet1(((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))).Int()).(_dafny.Int)).Cmp((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))).Int()).(_dafny.Int)) > 0, (_index157).Int())
				var _987_v100 _dafny.Sequence
				_ = _987_v100
				_987_v100 = _dafny.SeqOf((_this).F9())
				var _988_v101 T2
				_ = _988_v101
				var _nw187 *C1 = New_C1_()
				_ = _nw187
				_nw187.Ctor__((_this).F8(), (_987_v100).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_987_v100).Cardinality()))).Uint32()).(bool), (_this).F6(), p3, _978_v91.F4(), (_this).F5())
				_988_v101 = _nw187
				var _989_v102 D0
				_ = _989_v102
				_989_v102 = Companion_D0_.Create_DC1_(!(p1), (_978_v91).F8(), p0, (_978_v91).F8())
				var _990_v103 _dafny.Set
				_ = _990_v103
				_990_v103 = _dafny.SetOf(_989_v102, _989_v102)
				var _991_v104 _dafny.Set
				_ = _991_v104
				_991_v104 = _dafny.SetOf((_this).F8())
				var _992_v105 _dafny.Map
				_ = _992_v105
				_992_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T2 {
					if p0 {
						return _988_v101
					}
					return _988_v101
				})(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm14((_990_v103).Cardinality(), _dafny.IntOfUint32(((_988_v101).F5()).Cardinality()), (_978_v91).F6(), (_988_v101).F9(), globalState), _dafny.Companion_Sequence_.Update((_988_v101).F5(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(718), _dafny.IntOfUint32(((_988_v101).F5()).Cardinality()))).Uint32(), ((_this).F5()).Select((Companion_Default___.SafeIndex((_991_v104).Cardinality(), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32()).(_dafny.CodePoint)))).Cardinality()))
				var _index158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index158
				var _index159 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))
				_ = _index159
				var _rhs142 _dafny.Sequence = _986_v99
				_ = _rhs142
				var _rhs143 bool = (_988_v101).F6()
				_ = _rhs143
				var _rhs144 _dafny.Map = (_992_v105).Merge(_992_v105)
				_ = _rhs144
				var _rhs145 _dafny.Array = _dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))
				_ = _rhs145
				var _rhs146 _dafny.Int = (_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_987_v100, _987_v100)).Cardinality())).Times((_978_v91).F8())
				_ = _rhs146
				var _lhs106 _dafny.Array = (_this).F10()
				_ = _lhs106
				var _lhs107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs107
				var _lhs108 _dafny.Array = _980_v93
				_ = _lhs108
				var _lhs109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))
				_ = _lhs109
				var _lhs110 *GlobalState = globalState
				_ = _lhs110
				_986_v99 = _rhs142
				(_lhs106).ArraySet1(_rhs143, (_lhs107).Int())
				_992_v105 = _rhs144
				(_lhs108).ArraySet1(_rhs145, (_lhs109).Int())
				_lhs110.F1 = _rhs146
				(_this).F4_set_(_988_v101.F4())
				var _993_v106 _dafny.Array
				_ = _993_v106
				var _nwElement0_38 bool = p3
				_ = _nwElement0_38
				var _nw188 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_38, nil, _dafny.IntOfInt64(21))
				_ = _nw188
				(_nw188).ArraySet1(_nwElement0_38, 0)
				(_nw188).ArraySet1(Companion_Default___.Fm3(globalState), 1)
				(_nw188).ArraySet1(p1, 2)
				(_nw188).ArraySet1(p1, 3)
				(_nw188).ArraySet1((_978_v91).F7(), 4)
				(_nw188).ArraySet1(!(!(p1)), 5)
				(_nw188).ArraySet1((_this).F6(), 6)
				(_nw188).ArraySet1(p3, 7)
				(_nw188).ArraySet1(false, 8)
				(_nw188).ArraySet1(Companion_Default___.Fm3(globalState), 9)
				(_nw188).ArraySet1((_this).F7(), 10)
				(_nw188).ArraySet1((_this).F9(), 11)
				(_nw188).ArraySet1(false, 12)
				(_nw188).ArraySet1(p3, 13)
				(_nw188).ArraySet1(false, 14)
				(_nw188).ArraySet1((_978_v91).F6(), 15)
				(_nw188).ArraySet1((_978_v91).F9(), 16)
				(_nw188).ArraySet1(Companion_Default___.Fm7(_978_v91.F4(), _dafny.IntOfInt64(-833), p1, globalState), 17)
				(_nw188).ArraySet1((_978_v91).F9(), 18)
				(_nw188).ArraySet1((_978_v91).F7(), 19)
				(_nw188).ArraySet1((_978_v91).F6(), 20)
				_993_v106 = _nw188
				var _994_v107 _dafny.Array
				_ = _994_v107
				var _len0_18 _dafny.Int = _dafny.IntOfInt64(6)
				_ = _len0_18
				var _nw189 _dafny.Array
				_ = _nw189
				if _len0_18.Cmp(_dafny.Zero) == 0 {
					_nw189 = _dafny.NewArray(_len0_18)
				} else {
					var _init18 func(_dafny.Int) _dafny.Sequence = (func(_995_v100 _dafny.Sequence) func(_dafny.Int) _dafny.Sequence {
						return func(_996_i3 _dafny.Int) _dafny.Sequence {
							return (Companion_D8_.Create_DC22_(_995_v100)).Dtor_cf41()
						}
					})(_987_v100)
					_ = _init18
					var _element0_18 = _init18(_dafny.Zero)
					_ = _element0_18
					_nw189 = _dafny.NewArrayFromExample(_element0_18, nil, _len0_18)
					(_nw189).ArraySet1(_element0_18, 0)
					var _nativeLen0_18 = (_len0_18).Int()
					_ = _nativeLen0_18
					for _i0_18 := 1; _i0_18 < _nativeLen0_18; _i0_18++ {
						(_nw189).ArraySet1(_init18(_dafny.IntOf(_i0_18)), _i0_18)
					}
				}
				_994_v107 = _nw189
				var _997_v108 D1
				_ = _997_v108
				_997_v108 = Companion_D1_.Create_DC3_(_994_v107)
				var _pat_let_tv16 = _994_v107
				_ = _pat_let_tv16
				var _998_v109 _dafny.Array
				_ = _998_v109
				var _nwElement0_39 D1 = _997_v108
				_ = _nwElement0_39
				var _nw190 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_39, nil, _dafny.IntOfInt64(25))
				_ = _nw190
				(_nw190).ArraySet1(_nwElement0_39, 0)
				(_nw190).ArraySet1(_997_v108, 1)
				(_nw190).ArraySet1(_997_v108, 2)
				(_nw190).ArraySet1(Companion_D1_.Create_DC3_(_994_v107), 3)
				(_nw190).ArraySet1(_997_v108, 4)
				(_nw190).ArraySet1(_997_v108, 5)
				(_nw190).ArraySet1(_997_v108, 6)
				(_nw190).ArraySet1(_997_v108, 7)
				(_nw190).ArraySet1(_997_v108, 8)
				(_nw190).ArraySet1(_997_v108, 9)
				(_nw190).ArraySet1(_997_v108, 10)
				(_nw190).ArraySet1(Companion_D1_.Create_DC3_(_994_v107), 11)
				(_nw190).ArraySet1(_997_v108, 12)
				(_nw190).ArraySet1(_997_v108, 13)
				(_nw190).ArraySet1(_997_v108, 14)
				(_nw190).ArraySet1(func(_pat_let26_0 D1) D1 {
					return func(_999_dt__update__tmp_h2 D1) D1 {
						return func(_pat_let27_0 _dafny.Array) D1 {
							return func(_1000_dt__update_hcf7_h0 _dafny.Array) D1 {
								return Companion_D1_.Create_DC3_(_1000_dt__update_hcf7_h0)
							}(_pat_let27_0)
						}(_pat_let_tv16)
					}(_pat_let26_0)
				}(Companion_D1_.Create_DC3_(_994_v107)), 15)
				(_nw190).ArraySet1(Companion_D1_.Create_DC3_(_994_v107), 16)
				(_nw190).ArraySet1(Companion_D1_.Create_DC3_(_994_v107), 17)
				(_nw190).ArraySet1(_997_v108, 18)
				(_nw190).ArraySet1(Companion_D1_.Create_DC3_(_994_v107), 19)
				(_nw190).ArraySet1(_997_v108, 20)
				(_nw190).ArraySet1(_997_v108, 21)
				(_nw190).ArraySet1(_997_v108, 22)
				(_nw190).ArraySet1(Companion_D1_.Create_DC3_(_994_v107), 23)
				(_nw190).ArraySet1(_997_v108, 24)
				_998_v109 = _nw190
				var _1001_v110 *C2
				_ = _1001_v110
				var _nw191 *C2 = New_C2_()
				_ = _nw191
				_nw191.Ctor__(_993_v106, _998_v109)
				_1001_v110 = _nw191
				var _1002_v111 _dafny.Set
				_ = _1002_v111
				_1002_v111 = _dafny.SetOf((_978_v91).F5())
				var _index160 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index160
				var _rhs147 bool = (_1002_v111).IsSubsetOf((func() _dafny.Set {
					if true {
						return _dafny.SetOf(_dafny.UnicodeSeqOfUtf8Bytes("lhuqmt"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(737))).Uint32(), func(coer53 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
							return func(arg53 _dafny.Int) interface{} {
								return coer53(arg53)
							}
						}((func(_1003_v91 T3) func(_dafny.Int) _dafny.CodePoint {
							return func(_1004_i4 _dafny.Int) _dafny.CodePoint {
								return _1003_v91.F4()
							}
						})(_978_v91))))
					}
					return _1002_v111
				})())
				_ = _rhs147
				var _rhs148 *C2 = _1001_v110
				_ = _rhs148
				var _lhs111 _dafny.Array = (_this).F10()
				_ = _lhs111
				var _lhs112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(492), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs112
				(_lhs111).ArraySet1(_rhs147, (_lhs112).Int())
				_1001_v110 = _rhs148
			} else {
				var _1005_v112 _dafny.Sequence
				_ = _1005_v112
				_1005_v112 = _dafny.SeqOf(_dafny.IntOfInt64(910))
				var _1006_v113 _dafny.Sequence
				_ = _1006_v113
				_1006_v113 = _dafny.SeqOf(_dafny.IntOfUint32((_1005_v112).Cardinality()))
				var _1007_v114 _dafny.Map
				_ = _1007_v114
				_1007_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F10(), (_973_v86).IsDisjointFrom(_dafny.MultiSetFromSeq(_1006_v113)))
				var _1008_v115 _dafny.Map
				_ = _1008_v115
				_1008_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_978_v91).F6(), (_this).F6())
				_1007_v114 = (_1007_v114).Update((_this).F10(), (func() bool {
					if (_1008_v115).Contains(false) {
						return (_1008_v115).Get(false).(bool)
					}
					return (_978_v91).F9()
				})())
				(globalState).F1 = (_978_v91).F8()
				var _1009_v116 _dafny.Array
				_ = _1009_v116
				var _len0_19 _dafny.Int = _dafny.One
				_ = _len0_19
				var _nw192 _dafny.Array
				_ = _nw192
				if _len0_19.Cmp(_dafny.Zero) == 0 {
					_nw192 = _dafny.NewArray(_len0_19)
				} else {
					var _init19 func(_dafny.Int) _dafny.Map = (func(_1010_v91 T3, _1011_v93 _dafny.Array) func(_dafny.Int) _dafny.Map {
						return func(_1012_i5 _dafny.Int) _dafny.Map {
							return (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_1010_v91.F4(), (_1010_v91).F8(), _1010_v91.F4(), (_this).F8(), (_1010_v91).F7()), (_1010_v91).F8())).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D1_.Create_DC4_(_1010_v91.F4(), (_dafny.ArrayCastTo((_1011_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_1011_v93), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_1011_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_1011_v93), 0))).Int()))), 0))).Int()).(_dafny.Int), _this.F4(), (_1010_v91).F8(), (_1010_v91).F6()), (_dafny.ArrayCastTo((_1011_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_1011_v93), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_1011_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_1011_v93), 0))).Int()))), 0))).Int()).(_dafny.Int)))
						}
					})(_978_v91, _980_v93)
					_ = _init19
					var _element0_19 = _init19(_dafny.Zero)
					_ = _element0_19
					_nw192 = _dafny.NewArrayFromExample(_element0_19, nil, _len0_19)
					(_nw192).ArraySet1(_element0_19, 0)
					var _nativeLen0_19 = (_len0_19).Int()
					_ = _nativeLen0_19
					for _i0_19 := 1; _i0_19 < _nativeLen0_19; _i0_19++ {
						(_nw192).ArraySet1(_init19(_dafny.IntOf(_i0_19)), _i0_19)
					}
				}
				_1009_v116 = _nw192
				var _1013_v117 D1
				_ = _1013_v117
				_1013_v117 = Companion_D1_.Create_DC4_(_978_v91.F4(), (_this).F8(), _978_v91.F4(), (_this).F8(), false)
				var _1014_v118 _dafny.Map
				_ = _1014_v118
				_1014_v118 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1013_v117, (_978_v91).F8())
				var _index161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1009_v116), 0))
				_ = _index161
				(_1009_v116).ArraySet1(_1014_v118, (_index161).Int())
				var _1015_v119 _dafny.Sequence
				_ = _1015_v119
				_1015_v119 = _dafny.UnicodeSeqOfUtf8Bytes("ciyeo")
				var _1016_v121 T1
				_ = _1016_v121
				var _nw193 *C0 = New_C0_()
				_ = _nw193
				_nw193.Ctor__(Companion_Default___.Fm3(globalState), _977_v90, _dafny.CodePoint('f'), (_978_v91).F5())
				_1016_v121 = _nw193
				var _1017_v122 _dafny.MultiSet
				_ = _1017_v122
				_1017_v122 = _dafny.MultiSetOf(_1016_v121)
				var _1018_v123 _dafny.Map
				_ = _1018_v123
				_1018_v123 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_978_v91).F8()).Times((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))).Int()).(_dafny.Int)), (_this).F5())
				var _index162 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1009_v116), 0))
				_ = _index162
				var _rhs149 _dafny.Map = (func() _dafny.Map {
					if p0 {
						return func() _dafny.Map {
							var _coll20 = _dafny.NewMapBuilder()
							_ = _coll20
							for _iter23 := _dafny.Iterate((_dafny.MultiSetOf(Companion_D1_.Create_DC4_(_this.F4(), (_1017_v122).Cardinality(), _1016_v121.F4(), _dafny.IntOfUint32(((_978_v91).F5()).Cardinality()), p0))).Elements()); ; {
								_compr_20, _ok23 := _iter23()
								if !_ok23 {
									break
								}
								var _1019_v120 D1
								_1019_v120 = interface{}(_compr_20).(D1)
								if (_dafny.MultiSetOf(Companion_D1_.Create_DC4_(_this.F4(), (_1017_v122).Cardinality(), _1016_v121.F4(), _dafny.IntOfUint32(((_978_v91).F5()).Cardinality()), p0))).Contains(_1019_v120) {
									_coll20.Add(_1019_v120, (_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))).Int()).(_dafny.Int))
								}
							}
							return _coll20.ToMap()
						}()
					}
					return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1013_v117, (_this).F8())
				})()
				_ = _rhs149
				var _rhs150 bool = (((_this).F8()).Plus((_978_v91).F8())).Cmp((_dafny.IntOfUint32(((_978_v91).F5()).Cardinality())).Times((_978_v91).F8())) >= 0
				_ = _rhs150
				var _rhs151 _dafny.Sequence = (func() _dafny.Sequence {
					if (_1018_v123).Contains((_978_v91).F8()) {
						return (_1018_v123).Get((_978_v91).F8()).(_dafny.Sequence)
					}
					return (_1016_v121).F5()
				})()
				_ = _rhs151
				var _lhs113 _dafny.Array = _1009_v116
				_ = _lhs113
				var _lhs114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(893), _dafny.ArrayLen((_1009_v116), 0))
				_ = _lhs114
				(_lhs113).ArraySet1(_rhs149, (_lhs114).Int())
				_977_v90 = _rhs150
				_1015_v119 = _rhs151
				var _index163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index163
				((_this).F10()).ArraySet1(!((func() bool {
					if (_this).F9() {
						return (_978_v91).F7()
					}
					return (_978_v91).F9()
				})()), (_index163).Int())
				var _index164 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index164
				((_this).F10()).ArraySet1((_978_v91).F6(), (_index164).Int())
				var _1020_v124 _dafny.Map
				_ = _1020_v124
				_1020_v124 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int())), (_dafny.Zero).Minus((_978_v91).F8()))
				var _index165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index165
				var _arr2 _dafny.Array = _dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))
				_ = _arr2
				var _index166 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))
				_ = _index166
				var _rhs152 bool = (_1016_v121).F7()
				_ = _rhs152
				var _rhs153 bool = (_dafny.IntOfUint32(((_978_v91).F5()).Cardinality())).Cmp((_1020_v124).Cardinality()) > 0
				_ = _rhs153
				var _rhs154 _dafny.Sequence = _1015_v119
				_ = _rhs154
				var _rhs155 _dafny.Int = _dafny.IntOfInt64(339)
				_ = _rhs155
				var _lhs115 _dafny.Array = (_this).F10()
				_ = _lhs115
				var _lhs116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(653), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _lhs116
				var _lhs117 _dafny.Array = _dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))
				_ = _lhs117
				var _lhs118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))
				_ = _lhs118
				(_lhs115).ArraySet1(_rhs152, (_lhs116).Int())
				_977_v90 = _rhs153
				_1015_v119 = _rhs154
				(_lhs117).ArraySet1(_rhs155, (_lhs118).Int())
			}
			var _arr3 _dafny.Array = _dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))
			_ = _arr3
			var _index167 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(341), _dafny.ArrayLen((_dafny.ArrayCastTo((_980_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(451), _dafny.ArrayLen((_980_v93), 0))).Int()))), 0))
			_ = _index167
			(_arr3).ArraySet1((_978_v91).F8(), (_index167).Int())
		} else {
			var _1021_v125 _dafny.MultiSet
			_ = _1021_v125
			_1021_v125 = _dafny.MultiSetOf((_this).F9())
			_1021_v125 = _dafny.MultiSetOf((func(_pat_let28_0 D8) D8 {
				return func(_1022_dt__update__tmp_h3 D8) D8 {
					return func(_pat_let29_0 bool) D8 {
						return func(_1023_dt__update_hcf43_h0 bool) D8 {
							return Companion_D8_.Create_DC23_((_1022_dt__update__tmp_h3).Dtor_cf42(), _1023_dt__update_hcf43_h0)
						}(_pat_let29_0)
					}(true)
				}(_pat_let28_0)
			}(Companion_D8_.Create_DC23_(_dafny.IntOfInt64(937), p3))).Dtor_cf43())
			var _1024_v126 bool
			_ = _1024_v126
			_1024_v126 = true
			_1024_v126 = p1
			var _index168 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index168
			((_this).F10()).ArraySet1(_dafny.Companion_Sequence_.Equal((_this).F5(), (_this).F5()), (_index168).Int())
			var _1025_v127 _dafny.Array
			_ = _1025_v127
			var _nwElement0_40 _dafny.Array = (_this).F10()
			_ = _nwElement0_40
			var _nw194 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_40, nil, _dafny.IntOfInt64(9))
			_ = _nw194
			(_nw194).ArraySet1(_nwElement0_40, 0)
			(_nw194).ArraySet1((_this).F10(), 1)
			(_nw194).ArraySet1((_this).F10(), 2)
			(_nw194).ArraySet1((_this).F10(), 3)
			(_nw194).ArraySet1((_this).F10(), 4)
			(_nw194).ArraySet1((_this).F10(), 5)
			(_nw194).ArraySet1((_this).F10(), 6)
			(_nw194).ArraySet1((_this).F10(), 7)
			(_nw194).ArraySet1((_this).F10(), 8)
			_1025_v127 = _nw194
			var _1026_v128 _dafny.Array
			_ = _1026_v128
			var _nw195 _dafny.Array = _dafny.NewArrayWithValue(Companion_D6_.Default(), _dafny.One)
			_ = _nw195
			_1026_v128 = _nw195
			var _index169 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_1026_v128), 0))
			_ = _index169
			(_1026_v128).ArraySet1(_974_v87, (_index169).Int())
			var _1027_v129 _dafny.Sequence
			_ = _1027_v129
			_1027_v129 = _dafny.SeqOf(Companion_Default___.SafeDivisionInt((_dafny.Zero).Minus((_this).F8()), (_this).F8()))
			var _index170 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index170
			var _index171 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_1026_v128), 0))
			_ = _index171
			var _rhs156 bool = !((_dafny.IntOfInt64(718)).Cmp((_this).F8()) == 0)
			_ = _rhs156
			var _rhs157 _dafny.Array = _1025_v127
			_ = _rhs157
			var _rhs158 D6 = _974_v87
			_ = _rhs158
			var _rhs159 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1027_v129, _1027_v129)
			_ = _rhs159
			var _lhs119 _dafny.Array = (_this).F10()
			_ = _lhs119
			var _lhs120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs120
			var _lhs121 _dafny.Array = _1026_v128
			_ = _lhs121
			var _lhs122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(432), _dafny.ArrayLen((_1026_v128), 0))
			_ = _lhs122
			(_lhs119).ArraySet1(_rhs156, (_lhs120).Int())
			_1025_v127 = _rhs157
			(_lhs121).ArraySet1(_rhs158, (_lhs122).Int())
			_1027_v129 = _rhs159
			var _1028_v130 _dafny.Map
			_ = _1028_v130
			_1028_v130 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (func() bool {
				if true {
					return !(!(_1024_v126))
				}
				return true
			})())
			var _1029_v131 _dafny.Array
			_ = _1029_v131
			var _len0_20 _dafny.Int = _dafny.IntOfInt64(27)
			_ = _len0_20
			var _nw196 _dafny.Array
			_ = _nw196
			if _len0_20.Cmp(_dafny.Zero) == 0 {
				_nw196 = _dafny.NewArray(_len0_20)
			} else {
				var _init20 func(_dafny.Int) bool = (func(_1030_v126 bool) func(_dafny.Int) bool {
					return func(_1031_i6 _dafny.Int) bool {
						return _1030_v126
					}
				})(_1024_v126)
				_ = _init20
				var _element0_20 = _init20(_dafny.Zero)
				_ = _element0_20
				_nw196 = _dafny.NewArrayFromExample(_element0_20, nil, _len0_20)
				(_nw196).ArraySet1(_element0_20, 0)
				var _nativeLen0_20 = (_len0_20).Int()
				_ = _nativeLen0_20
				for _i0_20 := 1; _i0_20 < _nativeLen0_20; _i0_20++ {
					(_nw196).ArraySet1(_init20(_dafny.IntOf(_i0_20)), _i0_20)
				}
			}
			_1029_v131 = _nw196
			var _1032_v132 _dafny.Map
			_ = _1032_v132
			_1032_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F8())
			var _1033_v133 D5
			_ = _1033_v133
			_1033_v133 = Companion_D5_.Create_DC15_(_1029_v131, p1, _1032_v132, (_this).F8())
			var _index172 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index172
			((_this).F10()).ArraySet1((func() bool {
				if (_1028_v130).Contains(((_this).F8()).Minus((_this).F8())) {
					return (_1028_v130).Get(((_this).F8()).Minus((_this).F8())).(bool)
				}
				return (_1033_v133).Dtor_cf27()
			})(), (_index172).Int())
			var _1034_v134 _dafny.Map
			_ = _1034_v134
			_1034_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.IntOfInt64(-229))
			var _1035_v135 *C3
			_ = _1035_v135
			var _nw197 *C3 = New_C3_()
			_ = _nw197
			_nw197.Ctor__(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(757), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), p1, (func() _dafny.Int {
				if (_1034_v134).Contains((_this).F8()) {
					return (_1034_v134).Get((_this).F8()).(_dafny.Int)
				}
				return (_973_v86).Cardinality()
			})(), (_this).F6(), _1024_v126, (_this).F7(), _this.F4(), (_this).F5())
			_1035_v135 = _nw197
			var _1036_v136 _dafny.Map
			_ = _1036_v136
			_1036_v136 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-343), _1035_v135)
			var _1037_v137 _dafny.Map
			_ = _1037_v137
			_1037_v137 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), (func() bool {
				if (_1028_v130).Contains((_this).F8()) {
					return (_1028_v130).Get((_this).F8()).(bool)
				}
				return (_1035_v135).F13()
			})())
			_1036_v136 = (_1036_v136).Update((_this).Fm1(p0, _1037_v137, p0, globalState), (func() *C3 {
				if (_this).F6() {
					return _1035_v135
				}
				return _1035_v135
			})())
		}
		r1 = Companion_Default___.SafeDivisionInt((_this).F8(), (_this).F8())
		var _1038_v138 _dafny.MultiSet
		_ = _1038_v138
		_1038_v138 = _dafny.MultiSetOf((_this).F11(), (_this).F11(), (_this).F11(), (_this).F11())
		var _1039_v139 _dafny.Sequence
		_ = _1039_v139
		_1039_v139 = _dafny.SeqOf(_1038_v138)
		r0 = (_dafny.Zero).Minus(((_1038_v138).Intersection(((_1039_v139).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1039_v139).Cardinality()))).Uint32()).(_dafny.MultiSet)).Union(_1038_v138))).Cardinality())
		r1 = (_this).F8()
		var _1040_v140 _dafny.Array
		_ = _1040_v140
		var _nw198 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
		_ = _nw198
		_1040_v140 = _nw198
		r2 = _1040_v140
		return r0, r1, r2
	}
}
func (_this *C5) M2(p0 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		var r0 _dafny.Map = _dafny.EmptyMap
		_ = r0
		var _1041_v0 _dafny.Sequence
		_ = _1041_v0
		_1041_v0 = _dafny.SeqOf((_this).F7())
		var _1042_v1 _dafny.Sequence
		_ = _1042_v1
		_1042_v1 = _dafny.SeqOf(false, (_this).F6(), (_1041_v0).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1041_v0).Cardinality()))).Uint32()).(bool), true)
		var _1043_v2 _dafny.MultiSet
		_ = _1043_v2
		_1043_v2 = _dafny.MultiSetOf((_this).F7(), (_this).F9(), (_this).F6(), (_this).F9(), true)
		if Companion_Default___.Fm27(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_1042_v1).Cardinality()), (_1043_v2).Cardinality()), (_this).F6(), (_this).F7(), globalState) {
			var _1044_v3 _dafny.Sequence
			_ = _1044_v3
			_1044_v3 = _dafny.UnicodeSeqOfUtf8Bytes("bgo")
			var _1045_v4 _dafny.Set
			_ = _1045_v4
			_1045_v4 = _dafny.SetOf((_this).F9())
			var _1046_v5 D3
			_ = _1046_v5
			_1046_v5 = Companion_D3_.Create_DC11_((_this).F6(), (_this).F8(), _1045_v4)
			var _1047_v6 _dafny.Map
			_ = _1047_v6
			_1047_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(281))).Uint32(), func(coer54 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg54 _dafny.Int) interface{} {
					return coer54(arg54)
				}
			}(func(_1048_i0 _dafny.Int) _dafny.CodePoint {
				return _this.F4()
			})))
			_1044_v3 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm28((_this).F8(), (_this).F7(), Companion_Default___.Fm29(globalState), _1046_v5, globalState), (func() _dafny.Sequence {
				if (_1047_v6).Contains((_this).F9()) {
					return (_1047_v6).Get((_this).F9()).(_dafny.Sequence)
				}
				return _1044_v3
			})()), (_this).F5())
			_1044_v3 = _1044_v3
			_1044_v3 = _dafny.UnicodeSeqOfUtf8Bytes("xolxenqc")
			var _1049_v7 _dafny.Array
			_ = _1049_v7
			var _nw199 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(22))
			_ = _nw199
			_1049_v7 = _nw199
			var _1050_v8 *C2
			_ = _1050_v8
			var _nw200 *C2 = New_C2_()
			_ = _nw200
			_nw200.Ctor__((_this).F10(), _1049_v7)
			_1050_v8 = _nw200
			var _index173 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen(((_1050_v8).F14()), 0))
			_ = _index173
			((_1050_v8).F14()).ArraySet1((_this).F7(), (_index173).Int())
			var _1051_v9 _dafny.Array
			_ = _1051_v9
			var _len0_21 _dafny.Int = _dafny.IntOfInt64(6)
			_ = _len0_21
			var _nw201 _dafny.Array
			_ = _nw201
			if _len0_21.Cmp(_dafny.Zero) == 0 {
				_nw201 = _dafny.NewArray(_len0_21)
			} else {
				var _init21 func(_dafny.Int) D6 = func(_1052_i1 _dafny.Int) D6 {
					return Companion_D6_.Create_DC17_(_dafny.UnicodeSeqOfUtf8Bytes("xy"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-640))).Uint32(), func(coer55 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
						return func(arg55 _dafny.Int) interface{} {
							return coer55(arg55)
						}
					}(func(_1053_i2 _dafny.Int) _dafny.CodePoint {
						return _this.F4()
					})))
				}
				_ = _init21
				var _element0_21 = _init21(_dafny.Zero)
				_ = _element0_21
				_nw201 = _dafny.NewArrayFromExample(_element0_21, nil, _len0_21)
				(_nw201).ArraySet1(_element0_21, 0)
				var _nativeLen0_21 = (_len0_21).Int()
				_ = _nativeLen0_21
				for _i0_21 := 1; _i0_21 < _nativeLen0_21; _i0_21++ {
					(_nw201).ArraySet1(_init21(_dafny.IntOf(_i0_21)), _i0_21)
				}
			}
			_1051_v9 = _nw201
			var _1054_v10 _dafny.Map
			_ = _1054_v10
			_1054_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1051_v9, _1050_v8)
			var _1055_v11 _dafny.Sequence
			_ = _1055_v11
			_1055_v11 = _dafny.SeqOf(_1051_v9)
			var _1056_v12 _dafny.Map
			_ = _1056_v12
			_1056_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), _dafny.IntOfUint32(((_this).F5()).Cardinality()))
			var _index174 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen(((_1050_v8).F14()), 0))
			_ = _index174
			var _rhs160 *C2 = (func() *C2 {
				if (_1054_v10).Contains((_1055_v11).Select((Companion_Default___.SafeIndex((_1056_v12).Cardinality(), _dafny.IntOfUint32((_1055_v11).Cardinality()))).Uint32()).(_dafny.Array)) {
					return (_1054_v10).Get((_1055_v11).Select((Companion_Default___.SafeIndex((_1056_v12).Cardinality(), _dafny.IntOfUint32((_1055_v11).Cardinality()))).Uint32()).(_dafny.Array)).(*C2)
				}
				return _1050_v8
			})()
			_ = _rhs160
			var _rhs161 bool = (_this).F9()
			_ = _rhs161
			var _lhs123 _dafny.Array = (_1050_v8).F14()
			_ = _lhs123
			var _lhs124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen(((_1050_v8).F14()), 0))
			_ = _lhs124
			_1050_v8 = _rhs160
			(_lhs123).ArraySet1(_rhs161, (_lhs124).Int())
			var _1057_v13 _dafny.Map
			_ = _1057_v13
			_1057_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1044_v3, (_this).F6())
			var _index175 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(111), _dafny.ArrayLen(((_1050_v8).F14()), 0))
			_ = _index175
			((_1050_v8).F14()).ArraySet1((func() bool {
				if (_1057_v13).Contains(_dafny.Companion_Sequence_.Concatenate(_1044_v3, _1044_v3)) {
					return (_1057_v13).Get(_dafny.Companion_Sequence_.Concatenate(_1044_v3, _1044_v3)).(bool)
				}
				return (_this).F9()
			})(), (_index175).Int())
		} else {
			var _1058_v14 bool
			_ = _1058_v14
			_1058_v14 = false
			_1058_v14 = _1058_v14
			var _1059_v15 _dafny.MultiSet
			_ = _1059_v15
			_1059_v15 = _dafny.MultiSetOf((_this).F8())
			var _1060_v16 _dafny.Sequence
			_ = _1060_v16
			_1060_v16 = _dafny.SeqOf(p0)
			var _1061_v17 _dafny.Map
			_ = _1061_v17
			_1061_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1060_v16, p0)
			var _1062_v18 _dafny.Map
			_ = _1062_v18
			_1062_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1058_v14, (func() _dafny.Int {
				if (_1059_v15).Contains((_1061_v17).Cardinality()) {
					return (_1059_v15).Multiplicity((_1061_v17).Cardinality())
				}
				return p0
			})())
			_1062_v18 = _1062_v18
			var _1063_v19 D3
			_ = _1063_v19
			_1063_v19 = Companion_D3_.Create_DC10_((_this).F5())
			var _1064_v20 D3
			_ = _1064_v20
			_1064_v20 = Companion_D3_.Create_DC12_(_1063_v19)
			var _source11 D3 = _1064_v20
			_ = _source11
			if _source11.Is_DC11() {
				var _1065___mcc_h0 bool = _source11.Get_().(D3_DC11).Cf20
				_ = _1065___mcc_h0
				var _1066___mcc_h1 _dafny.Int = _source11.Get_().(D3_DC11).Cf21
				_ = _1066___mcc_h1
				var _1067___mcc_h2 _dafny.Set = _source11.Get_().(D3_DC11).Cf22
				_ = _1067___mcc_h2
				var _1068_cf22 _dafny.Set = _1067___mcc_h2
				_ = _1068_cf22
				var _1069_cf21 _dafny.Int = _1066___mcc_h1
				_ = _1069_cf21
				var _1070_cf20 bool = _1065___mcc_h0
				_ = _1070_cf20
				var _1071_v21 *C0
				_ = _1071_v21
				var _nw202 *C0 = New_C0_()
				_ = _nw202
				_nw202.Ctor__((_this).F6(), _1070_cf20, (func() _dafny.CodePoint {
					if false {
						return _this.F4()
					}
					return _dafny.CodePoint('v')
				})(), _dafny.Companion_Sequence_.Concatenate((_this).F5(), (_this).F5()))
				_1071_v21 = _nw202
				var _1072_v22 _dafny.Array
				_ = _1072_v22
				var _nw203 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(16))
				_ = _nw203
				_1072_v22 = _nw203
				var _1073_v23 *C2
				_ = _1073_v23
				var _nw204 *C2 = New_C2_()
				_ = _nw204
				_nw204.Ctor__((_this).F10(), _1072_v22)
				_1073_v23 = _nw204
				_1069_cf21 = (_this).F8()
				(globalState).F1 = (_this).F8()
			} else if _source11.Is_DC10() {
				var _1074___mcc_h3 _dafny.Sequence = _source11.Get_().(D3_DC10).Cf19
				_ = _1074___mcc_h3
				var _1075_cf19 _dafny.Sequence = _1074___mcc_h3
				_ = _1075_cf19
				var _1076_v24 _dafny.Set
				_ = _1076_v24
				_1076_v24 = _dafny.SetOf((_this).F10())
				var _1077_v25 _dafny.Map
				_ = _1077_v25
				_1077_v25 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), false)
				var _1078_v26 D0
				_ = _1078_v26
				_1078_v26 = Companion_D0_.Create_DC1_((_this).F7(), ((_1059_v15).Cardinality()).Minus((_this).F8()), Companion_Default___.Fm7(_this.F4(), (_1076_v24).Cardinality(), (_this).F9(), globalState), (_1077_v25).Cardinality())
				var _rhs162 D0 = _1078_v26
				_ = _rhs162
				var _rhs163 _dafny.MultiSet = _1059_v15
				_ = _rhs163
				var _rhs164 bool = false
				_ = _rhs164
				_1078_v26 = _rhs162
				_1059_v15 = _rhs163
				_1058_v14 = _rhs164
				var _index176 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index176
				((_this).F10()).ArraySet1((_this).F9(), (_index176).Int())
				var _index177 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index177
				((_this).F10()).ArraySet1(((_1059_v15).Update((_this).F8(), Companion_Default___.Abs((_this).F8()))).IsSubsetOf(_1059_v15), (_index177).Int())
				var _1079_v27 *C4
				_ = _1079_v27
				var _nw205 *C4 = New_C4_()
				_ = _nw205
				_nw205.Ctor__(_dafny.IntOfInt64(-331), (_this).F6(), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(273), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), ((_this).F8()).Cmp(_dafny.IntOfInt64(629)) == 0, _this.F4(), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(112))).Uint32(), func(coer56 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg56 _dafny.Int) interface{} {
						return coer56(arg56)
					}
				}(func(_1080_i3 _dafny.Int) _dafny.CodePoint {
					return _this.F4()
				})))
				_1079_v27 = _nw205
				var _1081_v28 D0
				_ = _1081_v28
				_1081_v28 = Companion_D0_.Create_DC2_(_dafny.IntOfUint32((_1060_v16).Cardinality()), (_this).F6(), _dafny.SeqOf((_this).F8(), p0))
				var _1082_v29 _dafny.Map
				_ = _1082_v29
				_1082_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1081_v28).Dtor_cf6(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1059_v15).Cardinality(), p0))
				var _index178 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index178
				((_this).F11()).ArraySet1(((_1082_v29).Cardinality()).Times((_this).F8()), (_index178).Int())
				var _1083_v30 _dafny.Set
				_ = _1083_v30
				_1083_v30 = _dafny.SetOf(_dafny.MultiSetFromSeq(_1041_v0), _1043_v2, _1043_v2)
				var _index179 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(662), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index179
				((_this).F11()).ArraySet1(((_1083_v30).Cardinality()).Minus((_this).F8()), (_index179).Int())
			} else {
				var _1084___mcc_h4 D3 = _source11.Get_().(D3_DC12).Cf23
				_ = _1084___mcc_h4
				var _1085_cf23 D3 = _1084___mcc_h4
				_ = _1085_cf23
				var _1086_v31 _dafny.Array
				_ = _1086_v31
				var _nw206 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(22))
				_ = _nw206
				_1086_v31 = _nw206
				_1086_v31 = (_this).F11()
				var _1087_v32 D8
				_ = _1087_v32
				_1087_v32 = Companion_D8_.Create_DC22_(_dafny.SeqOf((_this).F6(), (_this).F7()))
				var _1088_v33 _dafny.Array
				_ = _1088_v33
				var _nwElement0_41 _dafny.Sequence = _1042_v1
				_ = _nwElement0_41
				var _nw207 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_41, nil, _dafny.IntOfInt64(29))
				_ = _nw207
				(_nw207).ArraySet1(_nwElement0_41, 0)
				(_nw207).ArraySet1(_1041_v0, 1)
				(_nw207).ArraySet1(_1042_v1, 2)
				(_nw207).ArraySet1((_1087_v32).Dtor_cf41(), 3)
				(_nw207).ArraySet1(Companion_Default___.Fm18((_this).F8(), (_this).F6(), globalState), 4)
				(_nw207).ArraySet1(_1042_v1, 5)
				(_nw207).ArraySet1(_1041_v0, 6)
				(_nw207).ArraySet1(_1041_v0, 7)
				(_nw207).ArraySet1(_1042_v1, 8)
				(_nw207).ArraySet1(_1041_v0, 9)
				(_nw207).ArraySet1(_1041_v0, 10)
				(_nw207).ArraySet1(_dafny.SeqOf(_1058_v14), 11)
				(_nw207).ArraySet1(_1041_v0, 12)
				(_nw207).ArraySet1(_1041_v0, 13)
				(_nw207).ArraySet1(_dafny.SeqOf(true), 14)
				(_nw207).ArraySet1(_1042_v1, 15)
				(_nw207).ArraySet1(_1042_v1, 16)
				(_nw207).ArraySet1(_1041_v0, 17)
				(_nw207).ArraySet1(_1042_v1, 18)
				(_nw207).ArraySet1(_1041_v0, 19)
				(_nw207).ArraySet1(_dafny.SeqOf(_1058_v14), 20)
				(_nw207).ArraySet1(_1041_v0, 21)
				(_nw207).ArraySet1(_1042_v1, 22)
				(_nw207).ArraySet1(_1041_v0, 23)
				(_nw207).ArraySet1(_1042_v1, 24)
				(_nw207).ArraySet1(_1041_v0, 25)
				(_nw207).ArraySet1(_1042_v1, 26)
				(_nw207).ArraySet1(_1041_v0, 27)
				(_nw207).ArraySet1(_1041_v0, 28)
				_1088_v33 = _nw207
				var _1089_v34 D1
				_ = _1089_v34
				_1089_v34 = Companion_D1_.Create_DC3_(_1088_v33)
				var _1090_v35 _dafny.Sequence
				_ = _1090_v35
				_1090_v35 = _dafny.SeqOf(_1089_v34)
				var _1091_v36 D9
				_ = _1091_v36
				_1091_v36 = Companion_D9_.Create_DC24_(_1090_v35)
				var _1092_v37 _dafny.Array
				_ = _1092_v37
				var _nwElement0_42 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1090_v35, _1090_v35), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1090_v35, _1090_v35)).Cardinality()))).Uint32(), _1089_v34)
				_ = _nwElement0_42
				var _nw208 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_42, nil, _dafny.IntOfInt64(23))
				_ = _nw208
				(_nw208).ArraySet1(_nwElement0_42, 0)
				(_nw208).ArraySet1(_1090_v35, 1)
				(_nw208).ArraySet1(_1090_v35, 2)
				(_nw208).ArraySet1(_1090_v35, 3)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Update(_1090_v35, (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1090_v35).Cardinality()))).Uint32(), _1089_v34), 4)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1090_v35, _1090_v35), 5)
				(_nw208).ArraySet1(_1090_v35, 6)
				(_nw208).ArraySet1(_1090_v35, 7)
				(_nw208).ArraySet1(_1090_v35, 8)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update(_dafny.SeqOf(_1089_v34, Companion_D1_.Create_DC3_(_1088_v33), _1089_v34, _1089_v34), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.SeqOf(_1089_v34, Companion_D1_.Create_DC3_(_1088_v33), _1089_v34, _1089_v34)).Cardinality()))).Uint32(), Companion_D1_.Create_DC3_(_1088_v33)), _1090_v35), 9)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1090_v35, _1090_v35), 10)
				(_nw208).ArraySet1(_dafny.SeqOf(_1089_v34), 11)
				(_nw208).ArraySet1(_dafny.SeqOf(_1089_v34), 12)
				(_nw208).ArraySet1(_1090_v35, 13)
				(_nw208).ArraySet1(_1090_v35, 14)
				(_nw208).ArraySet1(_1090_v35, 15)
				(_nw208).ArraySet1(_1090_v35, 16)
				(_nw208).ArraySet1(_1090_v35, 17)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1090_v35, _1090_v35), 18)
				(_nw208).ArraySet1(_1090_v35, 19)
				(_nw208).ArraySet1(_1090_v35, 20)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_1090_v35, _1090_v35), 21)
				(_nw208).ArraySet1(_dafny.Companion_Sequence_.Concatenate((_1091_v36).Dtor_cf44(), _1090_v35), 22)
				_1092_v37 = _nw208
				_1092_v37 = _1092_v37
				_1058_v14 = !(false)
				var _index180 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index180
				((_this).F10()).ArraySet1((_this).F6(), (_index180).Int())
				var _index181 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(771), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index181
				((_this).F10()).ArraySet1((_this).F9(), (_index181).Int())
			}
			(globalState).F1 = (_dafny.Zero).Minus(p0)
			var _1093_v38 D6
			_ = _1093_v38
			_1093_v38 = Companion_D6_.Create_DC17_((_this).F5(), (_this).F5())
			_1093_v38 = Companion_Default___.Fm31(globalState)
		}
		if (_this).F7() {
			var _1094_v39 _dafny.Array
			_ = _1094_v39
			var _nw209 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(3))
			_ = _nw209
			_1094_v39 = _nw209
			var _1095_v40 D1
			_ = _1095_v40
			_1095_v40 = Companion_D1_.Create_DC3_(_1094_v39)
			var _pat_let_tv17 = _1094_v39
			_ = _pat_let_tv17
			var _1096_v41 _dafny.Array
			_ = _1096_v41
			var _nwElement0_43 D1 = _1095_v40
			_ = _nwElement0_43
			var _nw210 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_43, nil, _dafny.IntOfInt64(21))
			_ = _nw210
			(_nw210).ArraySet1(_nwElement0_43, 0)
			(_nw210).ArraySet1(_1095_v40, 1)
			(_nw210).ArraySet1(Companion_D1_.Create_DC3_(_1094_v39), 2)
			(_nw210).ArraySet1(_1095_v40, 3)
			(_nw210).ArraySet1(Companion_D1_.Create_DC3_(_1094_v39), 4)
			(_nw210).ArraySet1(_1095_v40, 5)
			(_nw210).ArraySet1(_1095_v40, 6)
			(_nw210).ArraySet1(_1095_v40, 7)
			(_nw210).ArraySet1(_1095_v40, 8)
			(_nw210).ArraySet1(_1095_v40, 9)
			(_nw210).ArraySet1(_1095_v40, 10)
			(_nw210).ArraySet1(_1095_v40, 11)
			(_nw210).ArraySet1(Companion_D1_.Create_DC3_(_1094_v39), 12)
			(_nw210).ArraySet1(func(_pat_let30_0 D1) D1 {
				return func(_1097_dt__update__tmp_h0 D1) D1 {
					return func(_pat_let31_0 _dafny.Array) D1 {
						return func(_1098_dt__update_hcf7_h0 _dafny.Array) D1 {
							return Companion_D1_.Create_DC3_(_1098_dt__update_hcf7_h0)
						}(_pat_let31_0)
					}(_pat_let_tv17)
				}(_pat_let30_0)
			}(Companion_D1_.Create_DC3_(_1094_v39)), 13)
			(_nw210).ArraySet1(_1095_v40, 14)
			(_nw210).ArraySet1(_1095_v40, 15)
			(_nw210).ArraySet1(_1095_v40, 16)
			(_nw210).ArraySet1(_1095_v40, 17)
			(_nw210).ArraySet1(_1095_v40, 18)
			(_nw210).ArraySet1(_1095_v40, 19)
			(_nw210).ArraySet1(_1095_v40, 20)
			_1096_v41 = _nw210
			var _1099_v42 *C2
			_ = _1099_v42
			var _nw211 *C2 = New_C2_()
			_ = _nw211
			_nw211.Ctor__((_this).F10(), _1096_v41)
			_1099_v42 = _nw211
			var _1100_v43 _dafny.Map
			_ = _1100_v43
			_1100_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.IntOfInt64(612))
			r0 = _1100_v43
			var _source12 D1 = _1095_v40
			_ = _source12
			if _source12.Is_DC4() {
				var _1101___mcc_h5 _dafny.CodePoint = _source12.Get_().(D1_DC4).Cf8
				_ = _1101___mcc_h5
				var _1102___mcc_h6 _dafny.Int = _source12.Get_().(D1_DC4).Cf9
				_ = _1102___mcc_h6
				var _1103___mcc_h7 _dafny.CodePoint = _source12.Get_().(D1_DC4).Cf10
				_ = _1103___mcc_h7
				var _1104___mcc_h8 _dafny.Int = _source12.Get_().(D1_DC4).Cf11
				_ = _1104___mcc_h8
				var _1105___mcc_h9 bool = _source12.Get_().(D1_DC4).Cf12
				_ = _1105___mcc_h9
				var _1106_cf12 bool = _1105___mcc_h9
				_ = _1106_cf12
				var _1107_cf11 _dafny.Int = _1104___mcc_h8
				_ = _1107_cf11
				var _1108_cf10 _dafny.CodePoint = _1103___mcc_h7
				_ = _1108_cf10
				var _1109_cf9 _dafny.Int = _1102___mcc_h6
				_ = _1109_cf9
				var _1110_cf8 _dafny.CodePoint = _1101___mcc_h5
				_ = _1110_cf8
				_1109_cf9 = (_1109_cf9).Plus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1041_v0, _1042_v1)).Cardinality()))
				var _1111_v44 _dafny.MultiSet
				_ = _1111_v44
				_1111_v44 = _dafny.MultiSetOf((_this).F11(), (_this).F11())
				var _index182 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen(((_1099_v42).F14()), 0))
				_ = _index182
				((_1099_v42).F14()).ArraySet1((_1111_v44).IsDisjointFrom(_1111_v44), (_index182).Int())
				var _index183 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen(((_1099_v42).F14()), 0))
				_ = _index183
				((_1099_v42).F14()).ArraySet1((_this).F7(), (_index183).Int())
				var _1112_v45 *C4
				_ = _1112_v45
				var _nw212 *C4 = New_C4_()
				_ = _nw212
				_nw212.Ctor__((_this).F8(), (_this).F6(), (_this).F9(), !(((_1099_v42).F14()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(769), _dafny.ArrayLen(((_1099_v42).F14()), 0))).Int()).(bool)), _1110_cf8, (_this).F5())
				_1112_v45 = _nw212
				var _1113_v46 _dafny.Map
				_ = _1113_v46
				_1113_v46 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1112_v45, (_this).F6())
				var _1114_v47 _dafny.Map
				_ = _1114_v47
				_1114_v47 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1108_cf10, (_this).F9())
				var _1115_v48 _dafny.Map
				_ = _1115_v48
				_1115_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1113_v46, _1114_v47)
				var _1116_v49 _dafny.Map
				_ = _1116_v49
				_1116_v49 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7())
				var _rhs165 _dafny.Map = _1115_v48
				_ = _rhs165
				var _rhs166 _dafny.Int = (_this).Fm1(Companion_Default___.Fm3(globalState), _1116_v49, !((_this).F7()), globalState)
				_ = _rhs166
				var _lhs125 *GlobalState = globalState
				_ = _lhs125
				_1115_v48 = _rhs165
				_lhs125.F1 = _rhs166
				var _1117_v50 _dafny.Sequence
				_ = _1117_v50
				_1117_v50 = _dafny.SeqOf((_this).F8())
				(globalState).F1 = Companion_Default___.SafeDivisionInt(Companion_Default___.SafeDivisionInt(_1109_cf9, _1109_cf9), _dafny.IntOfUint32((_1117_v50).Cardinality()))
			} else if _source12.Is_DC5() {
				var _1118___mcc_h10 _dafny.Int = _source12.Get_().(D1_DC5).Cf13
				_ = _1118___mcc_h10
				var _1119_cf13 _dafny.Int = _1118___mcc_h10
				_ = _1119_cf13
				var _1120_v51 _dafny.Sequence
				_ = _1120_v51
				_1120_v51 = _dafny.SeqOf((_this).F8())
				var _1121_v52 _dafny.Sequence
				_ = _1121_v52
				_1121_v52 = _dafny.SeqOf(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _dafny.IntOfUint32((_1041_v0).Cardinality())), _1100_v43)
				var _1122_v53 D6
				_ = _1122_v53
				_1122_v53 = Companion_D6_.Create_DC17_(_dafny.Companion_Sequence_.Update((_this).F5(), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(555), _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32(), _this.F4()), (_this).F5())
				var _1123_v54 _dafny.Sequence
				_ = _1123_v54
				_1123_v54 = _dafny.SeqOf((_1120_v51).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1120_v51).Cardinality()))).Uint32()).(_dafny.Int), ((_1121_v52).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32(((_1122_v53).Dtor_cf31()).Cardinality()), _dafny.IntOfUint32((_1121_v52).Cardinality()))).Uint32()).(_dafny.Map)).Cardinality())
				var _1124_v55 _dafny.Array
				_ = _1124_v55
				var _nwElement0_44 bool = _dafny.Companion_Sequence_.Contains(_1120_v51, ((_1100_v43).Update(p0, _1119_cf13)).Cardinality())
				_ = _nwElement0_44
				var _nw213 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_44, nil, _dafny.IntOfInt64(5))
				_ = _nw213
				(_nw213).ArraySet1(_nwElement0_44, 0)
				(_nw213).ArraySet1((_this).F7(), 1)
				(_nw213).ArraySet1((_this).F9(), 2)
				(_nw213).ArraySet1((p0).Cmp((_this).F8()) < 0, 3)
				(_nw213).ArraySet1((func() bool {
					if (_this).F9() {
						return (_this).F7()
					}
					return (_this).F9()
				})(), 4)
				_1124_v55 = _nw213
				_1124_v55 = (_1099_v42).F14()
				var _1125_v56 bool
				_ = _1125_v56
				_1125_v56 = false
				_1125_v56 = (_this).F7()
				var _1126_v57 _dafny.Array
				_ = _1126_v57
				var _nw214 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(10))
				_ = _nw214
				_1126_v57 = _nw214
				var _1127_v58 D2
				_ = _1127_v58
				_1127_v58 = Companion_D2_.Create_DC8_((_this).F8(), _dafny.IntOfInt64(-507))
				var _1128_v59 D0
				_ = _1128_v59
				_1128_v59 = Companion_D0_.Create_DC0_()
				var _1129_v60 _dafny.Map
				_ = _1129_v60
				_1129_v60 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1128_v59, _1125_v56)
				var _1130_v61 D0
				_ = _1130_v61
				_1130_v61 = Companion_D0_.Create_DC1_((_this).F9(), p0, (_this).F7(), (_1129_v60).Cardinality())
				var _1131_v62 _dafny.Array
				_ = _1131_v62
				var _nwElement0_45 D2 = _1127_v58
				_ = _nwElement0_45
				var _nw215 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_45, nil, _dafny.IntOfInt64(26))
				_ = _nw215
				(_nw215).ArraySet1(_nwElement0_45, 0)
				(_nw215).ArraySet1(_1127_v58, 1)
				(_nw215).ArraySet1(_1127_v58, 2)
				(_nw215).ArraySet1(_1127_v58, 3)
				(_nw215).ArraySet1(Companion_D2_.Create_DC8_((_this).F8(), (_1130_v61).Dtor_cf3()), 4)
				(_nw215).ArraySet1(_1127_v58, 5)
				(_nw215).ArraySet1(_1127_v58, 6)
				(_nw215).ArraySet1(_1127_v58, 7)
				(_nw215).ArraySet1(_1127_v58, 8)
				(_nw215).ArraySet1(_1127_v58, 9)
				(_nw215).ArraySet1(_1127_v58, 10)
				(_nw215).ArraySet1(_1127_v58, 11)
				(_nw215).ArraySet1(_1127_v58, 12)
				(_nw215).ArraySet1(_1127_v58, 13)
				(_nw215).ArraySet1(_1127_v58, 14)
				(_nw215).ArraySet1(_1127_v58, 15)
				(_nw215).ArraySet1(_1127_v58, 16)
				(_nw215).ArraySet1(_1127_v58, 17)
				(_nw215).ArraySet1(_1127_v58, 18)
				(_nw215).ArraySet1(_1127_v58, 19)
				(_nw215).ArraySet1(_1127_v58, 20)
				(_nw215).ArraySet1(Companion_D2_.Create_DC8_(p0, (_this).F8()), 21)
				(_nw215).ArraySet1(_1127_v58, 22)
				(_nw215).ArraySet1(_1127_v58, 23)
				(_nw215).ArraySet1(_1127_v58, 24)
				(_nw215).ArraySet1(_1127_v58, 25)
				_1131_v62 = _nw215
				var _1132_v63 D10
				_ = _1132_v63
				_1132_v63 = Companion_D10_.Create_DC26_(_1131_v62)
				var _1133_v64 _dafny.Array
				_ = _1133_v64
				var _nwElement0_46 _dafny.Array = _1131_v62
				_ = _nwElement0_46
				var _nw216 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_46, nil, _dafny.IntOfInt64(8))
				_ = _nw216
				(_nw216).ArraySet1(_nwElement0_46, 0)
				(_nw216).ArraySet1((_1132_v63).Dtor_cf48(), 1)
				(_nw216).ArraySet1(_1131_v62, 2)
				(_nw216).ArraySet1(_1131_v62, 3)
				(_nw216).ArraySet1(_1131_v62, 4)
				(_nw216).ArraySet1(_1131_v62, 5)
				(_nw216).ArraySet1(_1131_v62, 6)
				(_nw216).ArraySet1(_1131_v62, 7)
				_1133_v64 = _nw216
				var _index184 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1126_v57), 0))
				_ = _index184
				(_1126_v57).ArraySet1(_1133_v64, (_index184).Int())
				var _index185 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(347), _dafny.ArrayLen((_1126_v57), 0))
				_ = _index185
				(_1126_v57).ArraySet1(_1133_v64, (_index185).Int())
				var _index186 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index186
				((_this).F11()).ArraySet1(_dafny.IntOfUint32(((_this).F5()).Cardinality()), (_index186).Int())
				var _1134_v65 _dafny.Map
				_ = _1134_v65
				_1134_v65 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), !(_1125_v56))
				var _index187 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index187
				var _rhs167 _dafny.CodePoint = _this.F4()
				_ = _rhs167
				var _rhs168 _dafny.Int = _dafny.IntOfInt64(223)
				_ = _rhs168
				var _rhs169 _dafny.Int = Companion_Default___.SafeDivisionInt((_this).F8(), Companion_Default___.SafeModuloInt((_this).F8(), (_this).F8()))
				_ = _rhs169
				var _rhs170 _dafny.CodePoint = _this.F4()
				_ = _rhs170
				var _rhs171 bool = Companion_Default___.Fm7(_dafny.CodePoint('v'), (_1134_v65).Cardinality(), (_dafny.SetOf((Companion_Default___.Fm24(_1119_cf13, (_this).F9(), (_this).F7(), p0, globalState)).Cardinality())).IsDisjointFrom(Companion_Default___.Fm30((_this).F6(), globalState)), globalState)
				_ = _rhs171
				var _lhs126 *C5 = _this
				_ = _lhs126
				var _lhs127 _dafny.Array = (_this).F11()
				_ = _lhs127
				var _lhs128 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(54), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _lhs128
				var _lhs129 *C5 = _this
				_ = _lhs129
				_lhs126.F4_set_(_rhs167)
				(_lhs127).ArraySet1(_rhs168, (_lhs128).Int())
				_1119_cf13 = _rhs169
				_lhs129.F4_set_(_rhs170)
				_1125_v56 = _rhs171
			} else {
				var _1135___mcc_h11 _dafny.Array = _source12.Get_().(D1_DC3).Cf7
				_ = _1135___mcc_h11
				var _1136_cf7 _dafny.Array = _1135___mcc_h11
				_ = _1136_cf7
				(globalState).F1 = p0
				(globalState).F1 = (_dafny.Zero).Minus((_this).F8())
				var _1137_v66 _dafny.Array
				_ = _1137_v66
				var _len0_22 _dafny.Int = _dafny.IntOfInt64(17)
				_ = _len0_22
				var _nw217 _dafny.Array
				_ = _nw217
				if _len0_22.Cmp(_dafny.Zero) == 0 {
					_nw217 = _dafny.NewArray(_len0_22)
				} else {
					var _init22 func(_dafny.Int) D2 = (func(_1138_p0 _dafny.Int) func(_dafny.Int) D2 {
						return func(_1139_i4 _dafny.Int) D2 {
							return Companion_D2_.Create_DC8_(_1138_p0, _dafny.IntOfInt64(730))
						}
					})(p0)
					_ = _init22
					var _element0_22 = _init22(_dafny.Zero)
					_ = _element0_22
					_nw217 = _dafny.NewArrayFromExample(_element0_22, nil, _len0_22)
					(_nw217).ArraySet1(_element0_22, 0)
					var _nativeLen0_22 = (_len0_22).Int()
					_ = _nativeLen0_22
					for _i0_22 := 1; _i0_22 < _nativeLen0_22; _i0_22++ {
						(_nw217).ArraySet1(_init22(_dafny.IntOf(_i0_22)), _i0_22)
					}
				}
				_1137_v66 = _nw217
				var _1140_v67 D10
				_ = _1140_v67
				_1140_v67 = Companion_D10_.Create_DC26_(_1137_v66)
				var _pat_let_tv18 = _1137_v66
				_ = _pat_let_tv18
				_1140_v67 = func(_pat_let32_0 D10) D10 {
					return func(_1141_dt__update__tmp_h1 D10) D10 {
						return func(_pat_let33_0 _dafny.Array) D10 {
							return func(_1142_dt__update_hcf48_h0 _dafny.Array) D10 {
								return Companion_D10_.Create_DC26_(_1142_dt__update_hcf48_h0)
							}(_pat_let33_0)
						}(_pat_let_tv18)
					}(_pat_let32_0)
				}(_1140_v67)
				var _1143_v68 bool
				_ = _1143_v68
				_1143_v68 = true
				_1143_v68 = true
			}
			(globalState).F1 = p0
			var _1144_v69 bool
			_ = _1144_v69
			_1144_v69 = true
			_1144_v69 = false
		} else {
			var _1145_v70 bool
			_ = _1145_v70
			_1145_v70 = true
			var _1146_v71 _dafny.Set
			_ = _1146_v71
			_1146_v71 = _dafny.SetOf(_dafny.IntOfUint32(((_this).F5()).Cardinality()), (_this).F8())
			_1145_v70 = ((func() _dafny.Set {
				if (_this).F6() {
					return _1146_v71
				}
				return _1146_v71
			})()).IsSubsetOf(_1146_v71)
			_1145_v70 = !(false)
			var _1147_v72 _dafny.Array
			_ = _1147_v72
			var _nw218 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw218
			_1147_v72 = _nw218
			_1147_v72 = _1147_v72
			var _index188 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index188
			((_this).F10()).ArraySet1((_this).F7(), (_index188).Int())
			var _index189 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(136), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index189
			((_this).F10()).ArraySet1(!(_1145_v70), (_index189).Int())
			var _1148_v73 _dafny.Map
			_ = _1148_v73
			_1148_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F8())
			r0 = ((_1148_v73).Merge(_1148_v73)).Merge(_1148_v73)
		}
		(globalState).F1 = (_dafny.Zero).Minus((_this).F8())
		if func(_source13 D2) bool {
			if _source13.Is_DC7() {
				var _1149___mcc_h12 bool = _source13.Get_().(D2_DC7).Cf15
				_ = _1149___mcc_h12
				var _1150_cf15 bool = _1149___mcc_h12
				_ = _1150_cf15
				return (_this).F6()
			} else if _source13.Is_DC8() {
				var _1151___mcc_h13 _dafny.Int = _source13.Get_().(D2_DC8).Cf16
				_ = _1151___mcc_h13
				var _1152___mcc_h14 _dafny.Int = _source13.Get_().(D2_DC8).Cf17
				_ = _1152___mcc_h14
				var _1153_cf17 _dafny.Int = _1152___mcc_h14
				_ = _1153_cf17
				var _1154_cf16 _dafny.Int = _1151___mcc_h13
				_ = _1154_cf16
				return (_this).F6()
			} else if _source13.Is_DC6() {
				var _1155___mcc_h15 _dafny.Map = _source13.Get_().(D2_DC6).Cf14
				_ = _1155___mcc_h15
				var _1156_cf14 _dafny.Map = _1155___mcc_h15
				_ = _1156_cf14
				return (_this).F9()
			} else {
				var _1157___mcc_h16 D2 = _source13.Get_().(D2_DC9).Cf18
				_ = _1157___mcc_h16
				var _1158_cf18 D2 = _1157___mcc_h16
				_ = _1158_cf18
				return _dafny.Companion_Sequence_.Equal(_dafny.SeqOf(_dafny.SetOf((_this).F7())), _dafny.SeqOf(_dafny.SetOf((_this).F6(), (_this).F9(), (_this).F9(), (_this).F7(), (_this).F6()), _dafny.SetOf((_this).F6(), (_this).F7(), (_this).F7(), (_this).F9())))
			}
		}(func(_pat_let34_0 D2) D2 {
			return func(_1159_dt__update__tmp_h2 D2) D2 {
				return func(_pat_let35_0 bool) D2 {
					return func(_1160_dt__update_hcf15_h0 bool) D2 {
						return Companion_D2_.Create_DC7_(_1160_dt__update_hcf15_h0)
					}(_pat_let35_0)
				}((_this).F9())
			}(_pat_let34_0)
		}(Companion_D2_.Create_DC7_((_this).F6()))) {
			(globalState).F1 = p0
			var _1161_v74 _dafny.Set
			_ = _1161_v74
			_1161_v74 = _dafny.SetOf((_this).F9(), (_this).F6())
			if ((_this).F8()).Cmp((_1161_v74).Cardinality()) == 0 {
				var _1162_v75 bool
				_ = _1162_v75
				_1162_v75 = false
				var _1163_v76 _dafny.Map
				_ = _1163_v76
				_1163_v76 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).F9())
				var _rhs172 _dafny.Int = Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(-556), p0)
				_ = _rhs172
				var _rhs173 bool = (_1162_v75) || ((_this).F7())
				_ = _rhs173
				var _rhs174 bool = _1162_v75
				_ = _rhs174
				var _rhs175 _dafny.Map = _1163_v76
				_ = _rhs175
				var _lhs130 *GlobalState = globalState
				_ = _lhs130
				_lhs130.F1 = _rhs172
				_1162_v75 = _rhs173
				_1162_v75 = _rhs174
				_1163_v76 = _rhs175
				(globalState).F1 = (_this).F8()
				var _1164_v77 _dafny.Sequence
				_ = _1164_v77
				_1164_v77 = _dafny.UnicodeSeqOfUtf8Bytes("rc")
				_1164_v77 = _dafny.UnicodeSeqOfUtf8Bytes("jdhcdkvjy")
				var _1165_v78 *C3
				_ = _1165_v78
				var _nw219 *C3 = New_C3_()
				_ = _nw219
				_nw219.Ctor__(false, true, _dafny.IntOfInt64(317), true, !(((_this).F8()).Cmp((_dafny.Zero).Minus(p0)) >= 0), (_this).F9(), _this.F4(), Companion_Default___.Fm14((_dafny.Zero).Minus((_this).F8()), _dafny.IntOfUint32((_1164_v77).Cardinality()), (_this).F7(), (_this).F6(), globalState))
				_1165_v78 = _nw219
				var _1166_v79 _dafny.Sequence
				_ = _1166_v79
				_1166_v79 = _dafny.SeqOf((_this).F8(), p0)
				var _1167_v80 _dafny.MultiSet
				_ = _1167_v80
				_1167_v80 = _dafny.MultiSetOf(_1166_v79, _1166_v79)
				var _1168_v81 D11
				_ = _1168_v81
				_1168_v81 = Companion_D11_.Create_DC29_(_1167_v80)
				var _1169_v82 *C4
				_ = _1169_v82
				var _nw220 *C4 = New_C4_()
				_ = _nw220
				_nw220.Ctor__(_dafny.IntOfInt64(738), ((_1168_v81).Dtor_cf55()).Contains(_1166_v79), (_this).F6(), (_this).F9(), _this.F4(), _1164_v77)
				_1169_v82 = _nw220
			} else {
				var _1170_v83 _dafny.Map
				_ = _1170_v83
				_1170_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				_1170_v83 = (_1170_v83).Update(p0, p0)
				var _1171_v84 _dafny.Array
				_ = _1171_v84
				var _nw221 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(16))
				_ = _nw221
				_1171_v84 = _nw221
				var _index190 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1171_v84), 0))
				_ = _index190
				(_1171_v84).ArraySet1(_1161_v74, (_index190).Int())
				var _1172_v85 bool
				_ = _1172_v85
				_1172_v85 = true
				var _index191 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1171_v84), 0))
				_ = _index191
				var _rhs176 _dafny.Set = _1161_v74
				_ = _rhs176
				var _rhs177 _dafny.Int = ((p0).Times((_this).F8())).Plus((_this).F8())
				_ = _rhs177
				var _rhs178 bool = !((func() bool {
					if (_this).F6() {
						return _1172_v85
					}
					return !((_this).F7())
				})())
				_ = _rhs178
				var _rhs179 _dafny.Int = p0
				_ = _rhs179
				var _lhs131 _dafny.Array = _1171_v84
				_ = _lhs131
				var _lhs132 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(915), _dafny.ArrayLen((_1171_v84), 0))
				_ = _lhs132
				var _lhs133 *GlobalState = globalState
				_ = _lhs133
				var _lhs134 *GlobalState = globalState
				_ = _lhs134
				(_lhs131).ArraySet1(_rhs176, (_lhs132).Int())
				_lhs133.F1 = _rhs177
				_1172_v85 = _rhs178
				_lhs134.F1 = _rhs179
				var _1173_v86 _dafny.Sequence
				_ = _1173_v86
				_1173_v86 = _dafny.SeqOf(_1041_v0)
				var _1174_v87 _dafny.Array
				_ = _1174_v87
				var _nwElement0_47 _dafny.Sequence = _1042_v1
				_ = _nwElement0_47
				var _nw222 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_47, nil, _dafny.IntOfInt64(11))
				_ = _nw222
				(_nw222).ArraySet1(_nwElement0_47, 0)
				(_nw222).ArraySet1(_1042_v1, 1)
				(_nw222).ArraySet1((_1173_v86).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1173_v86).Cardinality()))).Uint32()).(_dafny.Sequence), 2)
				(_nw222).ArraySet1(_dafny.SeqOf((_this).F9()), 3)
				(_nw222).ArraySet1(_1041_v0, 4)
				(_nw222).ArraySet1(_1042_v1, 5)
				(_nw222).ArraySet1(_1041_v0, 6)
				(_nw222).ArraySet1(_1041_v0, 7)
				(_nw222).ArraySet1(_1042_v1, 8)
				(_nw222).ArraySet1(_1042_v1, 9)
				(_nw222).ArraySet1(_1042_v1, 10)
				_1174_v87 = _nw222
				var _1175_v88 D1
				_ = _1175_v88
				_1175_v88 = Companion_D1_.Create_DC3_(_1174_v87)
				var _1176_v89 _dafny.Array
				_ = _1176_v89
				var _nwElement0_48 D1 = Companion_D1_.Create_DC3_(_1174_v87)
				_ = _nwElement0_48
				var _nw223 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_48, nil, _dafny.IntOfInt64(19))
				_ = _nw223
				(_nw223).ArraySet1(_nwElement0_48, 0)
				(_nw223).ArraySet1(_1175_v88, 1)
				(_nw223).ArraySet1(_1175_v88, 2)
				(_nw223).ArraySet1(_1175_v88, 3)
				(_nw223).ArraySet1(_1175_v88, 4)
				(_nw223).ArraySet1(Companion_D1_.Create_DC3_(_1174_v87), 5)
				(_nw223).ArraySet1(_1175_v88, 6)
				(_nw223).ArraySet1(_1175_v88, 7)
				(_nw223).ArraySet1(Companion_D1_.Create_DC3_(_1174_v87), 8)
				(_nw223).ArraySet1(_1175_v88, 9)
				(_nw223).ArraySet1(_1175_v88, 10)
				(_nw223).ArraySet1(_1175_v88, 11)
				(_nw223).ArraySet1(_1175_v88, 12)
				(_nw223).ArraySet1(Companion_D1_.Create_DC3_(_1174_v87), 13)
				(_nw223).ArraySet1(_1175_v88, 14)
				(_nw223).ArraySet1(_1175_v88, 15)
				(_nw223).ArraySet1(_1175_v88, 16)
				(_nw223).ArraySet1(Companion_D1_.Create_DC3_(_1174_v87), 17)
				(_nw223).ArraySet1(_1175_v88, 18)
				_1176_v89 = _nw223
				var _1177_v90 *C2
				_ = _1177_v90
				var _nw224 *C2 = New_C2_()
				_ = _nw224
				_nw224.Ctor__((_this).F10(), _1176_v89)
				_1177_v90 = _nw224
				var _1178_v91 *C0
				_ = _1178_v91
				var _nw225 *C0 = New_C0_()
				_ = _nw225
				_nw225.Ctor__((func() bool {
					if _1172_v85 {
						return (_this).F7()
					}
					return (_this).F7()
				})(), (_this).F6(), _this.F4(), _dafny.Companion_Sequence_.Concatenate((_this).F5(), _dafny.UnicodeSeqOfUtf8Bytes("mum")))
				_1178_v91 = _nw225
				var _1179_v92 _dafny.Sequence
				_ = _1179_v92
				_1179_v92 = _dafny.SeqOf(p0, (_dafny.Zero).Minus((_dafny.SetOf(_dafny.IntOfInt64(-577))).Cardinality()), _dafny.IntOfInt64(258))
				var _1180_v93 _dafny.Sequence
				_ = _1180_v93
				_1180_v93 = _dafny.SeqOf(_1179_v92, _dafny.SeqOf(p0, _dafny.IntOfUint32(((_this).F5()).Cardinality())), _1179_v92, _1179_v92)
				_1180_v93 = _1180_v93
			}
			var _1181_v94 bool
			_ = _1181_v94
			_1181_v94 = false
			_1181_v94 = (p0).Cmp(p0) == 0
			var _index192 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index192
			((_this).F10()).ArraySet1(true, (_index192).Int())
			var _index193 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(782), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index193
			((_this).F10()).ArraySet1(_1181_v94, (_index193).Int())
			var _1182_v95 _dafny.Array
			_ = _1182_v95
			var _len0_23 _dafny.Int = _dafny.IntOfInt64(16)
			_ = _len0_23
			var _nw226 _dafny.Array
			_ = _nw226
			if _len0_23.Cmp(_dafny.Zero) == 0 {
				_nw226 = _dafny.NewArray(_len0_23)
			} else {
				var _init23 func(_dafny.Int) bool = func(_1183_i5 _dafny.Int) bool {
					return (_this).F6()
				}
				_ = _init23
				var _element0_23 = _init23(_dafny.Zero)
				_ = _element0_23
				_nw226 = _dafny.NewArrayFromExample(_element0_23, nil, _len0_23)
				(_nw226).ArraySet1(_element0_23, 0)
				var _nativeLen0_23 = (_len0_23).Int()
				_ = _nativeLen0_23
				for _i0_23 := 1; _i0_23 < _nativeLen0_23; _i0_23++ {
					(_nw226).ArraySet1(_init23(_dafny.IntOf(_i0_23)), _i0_23)
				}
			}
			_1182_v95 = _nw226
			var _1184_v96 _dafny.Array
			_ = _1184_v96
			var _nw227 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(20))
			_ = _nw227
			_1184_v96 = _nw227
			var _1185_v97 *C2
			_ = _1185_v97
			var _nw228 *C2 = New_C2_()
			_ = _nw228
			_nw228.Ctor__(_1182_v95, _1184_v96)
			_1185_v97 = _nw228
		} else {
			var _1186_v98 _dafny.Sequence
			_ = _1186_v98
			_1186_v98 = _dafny.SeqOf(_1042_v1)
			var _1187_v99 _dafny.Array
			_ = _1187_v99
			var _nwElement0_49 _dafny.Sequence = _1041_v0
			_ = _nwElement0_49
			var _nw229 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_49, nil, _dafny.IntOfInt64(23))
			_ = _nw229
			(_nw229).ArraySet1(_nwElement0_49, 0)
			(_nw229).ArraySet1(_1042_v1, 1)
			(_nw229).ArraySet1(_1042_v1, 2)
			(_nw229).ArraySet1(_1042_v1, 3)
			(_nw229).ArraySet1(_1042_v1, 4)
			(_nw229).ArraySet1(_1042_v1, 5)
			(_nw229).ArraySet1(_1041_v0, 6)
			(_nw229).ArraySet1(_1042_v1, 7)
			(_nw229).ArraySet1(_dafny.SeqOf((_this).F7(), (_this).F9()), 8)
			(_nw229).ArraySet1(_1041_v0, 9)
			(_nw229).ArraySet1(_1042_v1, 10)
			(_nw229).ArraySet1(_1041_v0, 11)
			(_nw229).ArraySet1(_1041_v0, 12)
			(_nw229).ArraySet1(_1042_v1, 13)
			(_nw229).ArraySet1(_1042_v1, 14)
			(_nw229).ArraySet1(_1041_v0, 15)
			(_nw229).ArraySet1(_1042_v1, 16)
			(_nw229).ArraySet1(_1041_v0, 17)
			(_nw229).ArraySet1(_dafny.SeqOf((_this).F9()), 18)
			(_nw229).ArraySet1(_dafny.SeqOf((_this).F6(), (_this).F7()), 19)
			(_nw229).ArraySet1(_1041_v0, 20)
			(_nw229).ArraySet1(_1041_v0, 21)
			(_nw229).ArraySet1((_1186_v98).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1186_v98).Cardinality()))).Uint32()).(_dafny.Sequence), 22)
			_1187_v99 = _nw229
			var _1188_v100 _dafny.Map
			_ = _1188_v100
			_1188_v100 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetFromSeq((func() _dafny.Sequence {
				if false {
					return _1042_v1
				}
				return _1041_v0
			})())).Cardinality(), Companion_D1_.Create_DC3_(_1187_v99))
			var _1189_v101 D1
			_ = _1189_v101
			_1189_v101 = Companion_D1_.Create_DC3_(_1187_v99)
			_1188_v100 = (_1188_v100).Update((func() _dafny.Int {
				if (_this).F7() {
					return p0
				}
				return p0
			})(), _1189_v101)
			var _1190_v102 bool
			_ = _1190_v102
			_1190_v102 = false
			_1190_v102 = (_this).F6()
			var _1191_v103 _dafny.Set
			_ = _1191_v103
			_1191_v103 = _dafny.SetOf((_this).F9(), false, _1190_v102)
			var _1192_v104 D3
			_ = _1192_v104
			_1192_v104 = Companion_D3_.Create_DC11_(!((_this).F6()), (_this).F8(), _1191_v103)
			var _1193_v105 _dafny.Sequence
			_ = _1193_v105
			_1193_v105 = _dafny.SeqOf(_dafny.IntOfUint32((Companion_Default___.Fm28(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(66))).Uint32(), func(coer57 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg57 _dafny.Int) interface{} {
					return coer57(arg57)
				}
			}(func(_1194_i6 _dafny.Int) _dafny.CodePoint {
				return _this.F4()
			}))).Cardinality()), (_this).F6(), _1042_v1, _1192_v104, globalState)).Cardinality()))
			var _1195_v106 _dafny.Set
			_ = _1195_v106
			_1195_v106 = _dafny.SetOf((_this).F8(), (_1193_v105).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1193_v105).Cardinality()))).Uint32()).(_dafny.Int))
			var _1196_v107 _dafny.Map
			_ = _1196_v107
			_1196_v107 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_this).F6()), _dafny.Companion_Sequence_.Update(_1193_v105, (Companion_Default___.SafeIndex((_1193_v105).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1195_v106).Cardinality()), _dafny.IntOfUint32((_1193_v105).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_1193_v105).Cardinality()))).Uint32(), (_this).F8()))
			_1196_v107 = (_1196_v107).Merge(_1196_v107)
			_1190_v102 = (_1191_v103).Contains((_this).F6())
			var _1197_v108 _dafny.Map
			_ = _1197_v108
			_1197_v108 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F6())
			var _1198_v109 D1
			_ = _1198_v109
			_1198_v109 = Companion_D1_.Create_DC5_(p0)
			var _1199_v110 _dafny.MultiSet
			_ = _1199_v110
			_1199_v110 = _dafny.MultiSetOf(p0, p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, (_this).Fm1((_this).F7(), _1197_v108, Companion_Default___.Fm27((_this).F8(), _1190_v102, (_this).F6(), globalState), globalState)), (_this).F5())).Cardinality(), (_dafny.IntOfUint32((_dafny.SeqOf(_1198_v109, Companion_D1_.Create_DC5_(_dafny.IntOfUint32(((_this).F5()).Cardinality())), _1198_v109, Companion_D1_.Create_DC5_(_dafny.IntOfInt64(528)), _1198_v109)).Cardinality())).Times(_dafny.IntOfInt64(896)))
			(globalState).F1 = (func() _dafny.Int {
				if (_1199_v110).Contains(p0) {
					return (_1199_v110).Multiplicity(p0)
				}
				return p0
			})()
		}
		var _1200_v111 D1
		_ = _1200_v111
		_1200_v111 = Companion_D1_.Create_DC5_(_dafny.IntOfInt64(-740))
		var _1201_v112 _dafny.Sequence
		_ = _1201_v112
		_1201_v112 = _dafny.SeqOf(_1200_v111, _1200_v111)
		var _1202_v113 _dafny.Sequence
		_ = _1202_v113
		_1202_v113 = _1201_v112
		if func(_source14 _dafny.Sequence) bool {
			var _1203___mcc_h17 _dafny.Sequence = _source14
			_ = _1203___mcc_h17
			var _1204_cf24 _dafny.Sequence = _1203___mcc_h17
			_ = _1204_cf24
			return (_this).F7()
		}(_1202_v113) {
			(globalState).F1 = p0
			var _1205_v114 _dafny.Map
			_ = _1205_v114
			_1205_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_this.F4(), (_this).F6())
			var _1206_v115 _dafny.Map
			_ = _1206_v115
			_1206_v115 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F7())
			(_this).M4(p0, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F9())).Update((_this).F8(), (func() bool {
				if (_1205_v114).Contains(_this.F4()) {
					return (_1205_v114).Get(_this.F4()).(bool)
				}
				return (func() bool {
					if (_1206_v115).Contains((_this).F7()) {
						return (_1206_v115).Get((_this).F7()).(bool)
					}
					return (_this).F7()
				})()
			})()), globalState)
			(globalState).F1 = _dafny.IntOfUint32((_1041_v0).Cardinality())
			var _1207_v116 *C0
			_ = _1207_v116
			var _nw230 *C0 = New_C0_()
			_ = _nw230
			_nw230.Ctor__((_1043_v2).IsDisjointFrom(_1043_v2), (_this).F9(), _dafny.CodePoint('q'), _dafny.Companion_Sequence_.Concatenate((_this).F5(), (_this).F5()))
			_1207_v116 = _nw230
			var _1208_v118 _dafny.Sequence
			_ = _1208_v118
			_1208_v118 = _dafny.SeqOf(Companion_Default___.Fm8(p0, (_this).F9(), (_this).F6(), globalState))
			var _1209_v119 _dafny.Map
			_ = _1209_v119
			_1209_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), p0)
			var _1210_v120 _dafny.Map
			_ = _1210_v120
			_1210_v120 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_1208_v118).Cardinality()), _1209_v119)
			var _1211_v121 _dafny.Sequence
			_ = _1211_v121
			_1211_v121 = _dafny.SeqOf((func() _dafny.Map {
				if (_1210_v120).Contains((_this).F8()) {
					return (_1210_v120).Get((_this).F8()).(_dafny.Map)
				}
				return _1209_v119
			})())
			var _1212_v122 _dafny.Sequence
			_ = _1212_v122
			_1212_v122 = _dafny.SeqOf((_this).F8(), (func() _dafny.Map {
				var _coll21 = _dafny.NewMapBuilder()
				_ = _coll21
				for _iter24 := _dafny.Iterate((_1211_v121).Elements()); ; {
					_compr_21, _ok24 := _iter24()
					if !_ok24 {
						break
					}
					var _1213_v117 _dafny.Map
					_1213_v117 = interface{}(_compr_21).(_dafny.Map)
					if _dafny.Companion_Sequence_.Contains(_1211_v121, _1213_v117) {
						_coll21.Add(_1213_v117, (_this).F6())
					}
				}
				return _coll21.ToMap()
			}()).Cardinality())
			(globalState).F1 = ((_1212_v122).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1212_v122).Cardinality()))).Uint32()).(_dafny.Int)).Minus((p0).Minus(p0))
		} else {
			var _1214_v123 _dafny.MultiSet
			_ = _1214_v123
			_1214_v123 = _dafny.MultiSetOf((_this).F5(), _dafny.UnicodeSeqOfUtf8Bytes("lbxa"))
			var _1215_v124 _dafny.Set
			_ = _1215_v124
			_1215_v124 = _dafny.SetOf((_this).F7(), Companion_Default___.Fm27(p0, (_this).F9(), (_this).F6(), globalState))
			var _1216_v125 D3
			_ = _1216_v125
			_1216_v125 = Companion_D3_.Create_DC11_((_this).F9(), (_1214_v123).Cardinality(), _1215_v124)
			var _1217_v126 _dafny.Map
			_ = _1217_v126
			_1217_v126 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), Companion_Default___.Fm28(p0, false, _dafny.SeqOf((_this).F7()), _1216_v125, globalState))
			_1217_v126 = (func() _dafny.Map {
				if (_this).F9() {
					return _1217_v126
				}
				return _1217_v126
			})()
			var _1218_v127 bool
			_ = _1218_v127
			var _out21 bool
			_ = _out21
			_out21 = (_this).M3(p0, (_this).F8(), p0, globalState)
			_1218_v127 = _out21
			var _1219_v128 _dafny.Map
			_ = _1219_v128
			_1219_v128 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(424))).Uint32(), func(coer58 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
				return func(arg58 _dafny.Int) interface{} {
					return coer58(arg58)
				}
			}(func(_1220_i7 _dafny.Int) _dafny.Int {
				return _dafny.IntOfInt64(462)
			}))).Cardinality()))
			var _1221_v129 _dafny.Map
			_ = _1221_v129
			_1221_v129 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
			(globalState).F1 = Companion_Default___.SafeDivisionInt((func() _dafny.Int {
				if (_1219_v128).Contains(p0) {
					return (_1219_v128).Get(p0).(_dafny.Int)
				}
				return _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-830))).Uint32(), func(coer59 func(_dafny.Int) _dafny.Sequence) func(_dafny.Int) interface{} {
					return func(arg59 _dafny.Int) interface{} {
						return coer59(arg59)
					}
				}(func(_1222_i8 _dafny.Int) _dafny.Sequence {
					return _dafny.UnicodeSeqOfUtf8Bytes("tfpkek")
				}))).Cardinality())
			})(), (_this).Fm1(Companion_Default___.Fm3(globalState), _1221_v129, (_this).F7(), globalState))
			var _1223_v130 *C0
			_ = _1223_v130
			var _nw231 *C0 = New_C0_()
			_ = _nw231
			_nw231.Ctor__(_1218_v127, (_this).F9(), _this.F4(), (_this).F5())
			_1223_v130 = _nw231
			var _index194 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index194
			((_this).F10()).ArraySet1((_this).F9(), (_index194).Int())
			var _index195 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index195
			((_this).F10()).ArraySet1((_1043_v2).IsProperSubsetOf(_1043_v2), (_index195).Int())
			var _index196 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index196
			var _index197 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index197
			var _rhs180 bool = !((!((_this).F9()) || ((_this).F7())) || ((_this).F6()))
			_ = _rhs180
			var _rhs181 bool = ((_this).F8()).Cmp((_this).F8()) > 0
			_ = _rhs181
			var _rhs182 bool = (_this).F9()
			_ = _rhs182
			var _rhs183 bool = (_this).F7()
			_ = _rhs183
			var _rhs184 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_1041_v0, _dafny.Companion_Sequence_.Concatenate(_1041_v0, _1042_v1))
			_ = _rhs184
			var _lhs135 _dafny.Array = (_this).F10()
			_ = _lhs135
			var _lhs136 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(21), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs136
			var _lhs137 _dafny.Array = (_this).F10()
			_ = _lhs137
			var _lhs138 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(739), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs138
			(_lhs135).ArraySet1(_rhs180, (_lhs136).Int())
			(_lhs137).ArraySet1(_rhs181, (_lhs138).Int())
			_1218_v127 = _rhs182
			_1218_v127 = _rhs183
			_1042_v1 = _rhs184
		}
		if (_this).F9() {
			var _1224_v131 _dafny.Set
			_ = _1224_v131
			_1224_v131 = _dafny.SetOf(_dafny.Companion_Sequence_.IsPrefixOf(_1041_v0, _dafny.SeqOf((_this).F9(), (_this).F7())), !((_this).F9()) || ((_this).F9()), true)
			_1224_v131 = _1224_v131
			var _1225_v132 _dafny.Array
			_ = _1225_v132
			var _nw232 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(28))
			_ = _nw232
			_1225_v132 = _nw232
			var _index198 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_1225_v132), 0))
			_ = _index198
			(_1225_v132).ArraySet1((_this).F10(), (_index198).Int())
			var _index199 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(348), _dafny.ArrayLen((_1225_v132), 0))
			_ = _index199
			(_1225_v132).ArraySet1((_this).F10(), (_index199).Int())
			var _1226_v133 _dafny.Map
			_ = _1226_v133
			_1226_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_this).F9())
			var _1227_v134 D3
			_ = _1227_v134
			_1227_v134 = Companion_D3_.Create_DC11_((_this).F7(), _dafny.IntOfUint32(((_this).F5()).Cardinality()), _1224_v131)
			var _pat_let_tv19 = _1224_v131
			_ = _pat_let_tv19
			var _1228_v135 *C3
			_ = _1228_v135
			var _nw233 *C3 = New_C3_()
			_ = _nw233
			_nw233.Ctor__(!_dafny.Companion_Sequence_.Contains(_dafny.SeqOf(p0), _dafny.IntOfInt64(428)), !(_1224_v131).Contains((_this).F6()), (_this).Fm1((_this).F6(), _1226_v133, (_this).F6(), globalState), (false) && (Companion_Default___.Fm7(Companion_Default___.Fm13(globalState), Companion_Default___.Fm8((_this).F8(), (_this).F6(), (_1042_v1).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1042_v1).Cardinality()))).Uint32()).(bool), globalState), true, globalState)), (_this).F6(), (_this).F9(), _this.F4(), Companion_Default___.Fm28(p0, (_this).F6(), _1041_v0, func(_pat_let36_0 D3) D3 {
				return func(_1229_dt__update__tmp_h3 D3) D3 {
					return func(_pat_let37_0 _dafny.Set) D3 {
						return func(_1230_dt__update_hcf22_h0 _dafny.Set) D3 {
							return func(_pat_let38_0 bool) D3 {
								return func(_1231_dt__update_hcf20_h0 bool) D3 {
									return Companion_D3_.Create_DC11_(_1231_dt__update_hcf20_h0, (_1229_dt__update__tmp_h3).Dtor_cf21(), _1230_dt__update_hcf22_h0)
								}(_pat_let38_0)
							}((_this).F7())
						}(_pat_let37_0)
					}(_pat_let_tv19)
				}(_pat_let36_0)
			}(_1227_v134), globalState))
			_1228_v135 = _nw233
			_1228_v135 = _1228_v135
			var _1232_v136 _dafny.Sequence
			_ = _1232_v136
			_1232_v136 = _dafny.SeqOf(p0, (_dafny.Zero).Minus(Companion_Default___.Fm8(p0, (_1228_v135).F13(), !((_this).F9()), globalState)), (_this).F8())
			_1232_v136 = _1232_v136
			var _1233_v137 D1
			_ = _1233_v137
			_1233_v137 = Companion_D1_.Create_DC4_(_this.F4(), (_dafny.Zero).Minus(p0), _this.F4(), p0, _1228_v135.F12)
			var _1234_v138 _dafny.MultiSet
			_ = _1234_v138
			_1234_v138 = _dafny.MultiSetOf(_1233_v137)
			(_1228_v135).F12 = (!((_1224_v131).IsDisjointFrom(_1224_v131))) == ((_1234_v138).IsProperSubsetOf(_1234_v138))
		} else {
			var _1235_v139 bool
			_ = _1235_v139
			_1235_v139 = true
			var _1236_v140 _dafny.Array
			_ = _1236_v140
			var _nw234 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(15))
			_ = _nw234
			_1236_v140 = _nw234
			var _index200 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))
			_ = _index200
			(_1236_v140).ArraySet1((_this).F10(), (_index200).Int())
			var _index201 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index201
			((_this).F10()).ArraySet1(_dafny.Companion_Sequence_.Contains((_this).F5(), _this.F4()), (_index201).Int())
			var _index202 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))
			_ = _index202
			var _index203 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index203
			var _rhs185 bool = !(Companion_Default___.Fm3(globalState)) || ((_this).F6())
			_ = _rhs185
			var _rhs186 _dafny.CodePoint = _this.F4()
			_ = _rhs186
			var _rhs187 _dafny.Array = (_this).F10()
			_ = _rhs187
			var _rhs188 _dafny.Int = (_this).F8()
			_ = _rhs188
			var _rhs189 bool = (_this).F9()
			_ = _rhs189
			var _lhs139 *C5 = _this
			_ = _lhs139
			var _lhs140 _dafny.Array = _1236_v140
			_ = _lhs140
			var _lhs141 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))
			_ = _lhs141
			var _lhs142 *GlobalState = globalState
			_ = _lhs142
			var _lhs143 _dafny.Array = (_this).F10()
			_ = _lhs143
			var _lhs144 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs144
			_1235_v139 = _rhs185
			_lhs139.F4_set_(_rhs186)
			(_lhs140).ArraySet1(_rhs187, (_lhs141).Int())
			_lhs142.F1 = _rhs188
			(_lhs143).ArraySet1(_rhs189, (_lhs144).Int())
			(globalState).F1 = (_this).F8()
			var _1237_v141 _dafny.Array
			_ = _1237_v141
			var _len0_24 _dafny.Int = _dafny.IntOfInt64(29)
			_ = _len0_24
			var _nw235 _dafny.Array
			_ = _nw235
			if _len0_24.Cmp(_dafny.Zero) == 0 {
				_nw235 = _dafny.NewArray(_len0_24)
			} else {
				var _init24 func(_dafny.Int) _dafny.Int = func(_1238_i9 _dafny.Int) _dafny.Int {
					return (_1238_i9).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F9())).Cardinality())
				}
				_ = _init24
				var _element0_24 = _init24(_dafny.Zero)
				_ = _element0_24
				_nw235 = _dafny.NewArrayFromExample(_element0_24, nil, _len0_24)
				(_nw235).ArraySet1(_element0_24, 0)
				var _nativeLen0_24 = (_len0_24).Int()
				_ = _nativeLen0_24
				for _i0_24 := 1; _i0_24 < _nativeLen0_24; _i0_24++ {
					(_nw235).ArraySet1(_init24(_dafny.IntOf(_i0_24)), _i0_24)
				}
			}
			_1237_v141 = _nw235
			var _index204 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index204
			((_this).F11()).ArraySet1((_dafny.Zero).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F11(), (_this).F6())).Cardinality()), (_index204).Int())
			var _index205 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index205
			var _rhs190 _dafny.Array = _1237_v141
			_ = _rhs190
			var _rhs191 bool = _1235_v139
			_ = _rhs191
			var _rhs192 bool = (_this).F6()
			_ = _rhs192
			var _rhs193 bool = (_this).F6()
			_ = _rhs193
			var _rhs194 _dafny.Int = _dafny.IntOfUint32((_1041_v0).Cardinality())
			_ = _rhs194
			var _lhs145 _dafny.Array = (_this).F11()
			_ = _lhs145
			var _lhs146 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _lhs146
			_1237_v141 = _rhs190
			_1235_v139 = _rhs191
			_1235_v139 = _rhs192
			_1235_v139 = _rhs193
			(_lhs145).ArraySet1(_rhs194, (_lhs146).Int())
			var _1239_v142 _dafny.Sequence
			_ = _1239_v142
			_1239_v142 = _dafny.SeqOf((_this).F8())
			if _dafny.Companion_Sequence_.IsPrefixOf(_1239_v142, (func() _dafny.Sequence {
				if (_this).F6() {
					return _dafny.SeqOf(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), p0)
				}
				return _1239_v142
			})()) {
				var _1240_v143 _dafny.Map
				_ = _1240_v143
				_1240_v143 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_this).F9())
				(_this).M4((_dafny.Zero).Minus(_dafny.IntOfUint32((Companion_Default___.Fm18(p0, !(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)), globalState)).Cardinality())), _1240_v143, globalState)
				var _1241_v144 _dafny.Array
				_ = _1241_v144
				var _nw236 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw236
				_1241_v144 = _nw236
				var _1242_v145 _dafny.Array
				_ = _1242_v145
				var _nw237 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(17))
				_ = _nw237
				_1242_v145 = _nw237
				var _index206 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1241_v144), 0))
				_ = _index206
				(_1241_v144).ArraySet1(_1242_v145, (_index206).Int())
				var _1243_v146 _dafny.Array
				_ = _1243_v146
				var _nw238 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(22))
				_ = _nw238
				_1243_v146 = _nw238
				var _1244_v147 *C2
				_ = _1244_v147
				var _nw239 *C2 = New_C2_()
				_ = _nw239
				_nw239.Ctor__(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), _1243_v146)
				_1244_v147 = _nw239
				var _1245_v148 _dafny.MultiSet
				_ = _1245_v148
				_1245_v148 = _dafny.MultiSetOf(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int))
				var _index207 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1241_v144), 0))
				_ = _index207
				var _index208 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index208
				var _rhs195 _dafny.Array = _1242_v145
				_ = _rhs195
				var _rhs196 _dafny.Int = (_dafny.Zero).Minus((_this).F8())
				_ = _rhs196
				var _rhs197 _dafny.Int = p0
				_ = _rhs197
				var _rhs198 *C2 = (func() *C2 {
					if (p0).Cmp((_1245_v148).Cardinality()) <= 0 {
						return _1244_v147
					}
					return _1244_v147
				})()
				_ = _rhs198
				var _lhs147 _dafny.Array = _1241_v144
				_ = _lhs147
				var _lhs148 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(20), _dafny.ArrayLen((_1241_v144), 0))
				_ = _lhs148
				var _lhs149 *GlobalState = globalState
				_ = _lhs149
				var _lhs150 _dafny.Array = (_this).F11()
				_ = _lhs150
				var _lhs151 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _lhs151
				(_lhs147).ArraySet1(_rhs195, (_lhs148).Int())
				_lhs149.F1 = _rhs196
				(_lhs150).ArraySet1(_rhs197, (_lhs151).Int())
				_1244_v147 = _rhs198
				var _index209 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))
				_ = _index209
				((_this).F11()).ArraySet1(p0, (_index209).Int())
				_1042_v1 = _1042_v1
				var _1246_v149 D1
				_ = _1246_v149
				_1246_v149 = Companion_D1_.Create_DC4_(Companion_Default___.Fm13(globalState), p0, _this.F4(), (_this).F8(), (_this).F9())
				var _1247_v150 _dafny.Array
				_ = _1247_v150
				var _nwElement0_50 _dafny.Sequence = _1041_v0
				_ = _nwElement0_50
				var _nw240 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_50, nil, _dafny.IntOfInt64(21))
				_ = _nw240
				(_nw240).ArraySet1(_nwElement0_50, 0)
				(_nw240).ArraySet1(_1041_v0, 1)
				(_nw240).ArraySet1(_1041_v0, 2)
				(_nw240).ArraySet1(_dafny.Companion_Sequence_.Update(_1042_v1, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1042_v1).Cardinality()))).Uint32(), (_this).F7()), 3)
				(_nw240).ArraySet1(_1041_v0, 4)
				(_nw240).ArraySet1(_dafny.SeqOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), (_this).F6(), (_this).F7(), _1235_v139), 5)
				(_nw240).ArraySet1(_dafny.SeqOf(_1235_v139, (_this).F7()), 6)
				(_nw240).ArraySet1(_1042_v1, 7)
				(_nw240).ArraySet1(_1042_v1, 8)
				(_nw240).ArraySet1(_1041_v0, 9)
				(_nw240).ArraySet1(_dafny.SeqOf((_this).F7()), 10)
				(_nw240).ArraySet1(_dafny.SeqOf(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)), 11)
				(_nw240).ArraySet1(_1042_v1, 12)
				(_nw240).ArraySet1(_1042_v1, 13)
				(_nw240).ArraySet1(_dafny.SeqOf((_this).F9()), 14)
				(_nw240).ArraySet1(_1041_v0, 15)
				(_nw240).ArraySet1(_1042_v1, 16)
				(_nw240).ArraySet1(_1041_v0, 17)
				(_nw240).ArraySet1(_1042_v1, 18)
				(_nw240).ArraySet1(_1041_v0, 19)
				(_nw240).ArraySet1(_1041_v0, 20)
				_1247_v150 = _nw240
				var _1248_v151 _dafny.Map
				_ = _1248_v151
				_1248_v151 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v139, Companion_D1_.Create_DC3_(_1247_v150))
				var _1249_v152 _dafny.Int
				_ = _1249_v152
				var _1250_v153 bool
				_ = _1250_v153
				var _1251_v154 _dafny.Int
				_ = _1251_v154
				var _out22 _dafny.Int
				_ = _out22
				var _out23 bool
				_ = _out23
				var _out24 _dafny.Int
				_ = _out24
				_out22, _out23, _out24 = (_1244_v147).M7(_1246_v149, Companion_Default___.Fm32(_dafny.Companion_Sequence_.Update(_1041_v0, (Companion_Default___.SafeIndex(_dafny.IntOfInt64(767), _dafny.IntOfUint32((_1041_v0).Cardinality()))).Uint32(), _1235_v139), true, globalState), _1248_v151, globalState)
				_1249_v152 = _out22
				_1250_v153 = _out23
				_1251_v154 = _out24
			} else {
				(_this).F4_set_(_this.F4())
				_1235_v139 = (_this).F9()
				var _1252_v155 _dafny.Sequence
				_ = _1252_v155
				_1252_v155 = _dafny.UnicodeSeqOfUtf8Bytes("cklypmb")
				_1252_v155 = (_this).F5()
				var _1253_v156 _dafny.Map
				_ = _1253_v156
				_1253_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _1235_v139)
				var _1254_v157 _dafny.Map
				_ = _1254_v157
				_1254_v157 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(316), _1253_v156)
				var _1255_v158 _dafny.Sequence
				_ = _1255_v158
				_1255_v158 = _dafny.SeqOf((_this).F5())
				var _1256_v159 _dafny.Map
				_ = _1256_v159
				_1256_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus((_1254_v157).Cardinality()), (_1255_v158).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1255_v158).Cardinality()))).Uint32()).(_dafny.Sequence))
				_1256_v159 = (_1256_v159).Update(_dafny.IntOfUint32((_1041_v0).Cardinality()), _1252_v155)
				var _1257_v160 D9
				_ = _1257_v160
				_1257_v160 = Companion_D9_.Create_DC25_(Companion_Default___.Fm27((_this).F8(), (_this).F6(), !(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)), globalState), (_this).F7(), (Companion_Default___.Fm33(globalState)).Cardinality())
				var _index210 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))
				_ = _index210
				((_this).F10()).ArraySet1(((_dafny.Zero).Minus(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int))).Cmp(((_1257_v160).Dtor_cf47()).Times(p0)) <= 0, (_index210).Int())
			}
			var _1258_v161 _dafny.Map
			_ = _1258_v161
			_1258_v161 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v139, ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int))
			var _1259_v162 D5
			_ = _1259_v162
			_1259_v162 = Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), false, _1258_v161, ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int))
			var _1260_v163 _dafny.Map
			_ = _1260_v163
			_1260_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1235_v139, (_1042_v1).Select((Companion_Default___.SafeIndex(((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_1042_v1).Cardinality()))).Uint32()).(bool))
			var _1261_v164 _dafny.Array
			_ = _1261_v164
			var _nwElement0_51 D5 = _1259_v162
			_ = _nwElement0_51
			var _nw241 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_51, nil, _dafny.IntOfInt64(22))
			_ = _nw241
			(_nw241).ArraySet1(_nwElement0_51, 0)
			(_nw241).ArraySet1(_1259_v162, 1)
			(_nw241).ArraySet1(_1259_v162, 2)
			(_nw241).ArraySet1(_1259_v162, 3)
			(_nw241).ArraySet1(Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), (_this).F7(), _1258_v161, ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int)), 4)
			(_nw241).ArraySet1(Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), (_this).F7(), _1258_v161, (_this).F8()), 5)
			(_nw241).ArraySet1(_1259_v162, 6)
			(_nw241).ArraySet1(_1259_v162, 7)
			(_nw241).ArraySet1(_1259_v162, 8)
			(_nw241).ArraySet1(Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), _1235_v139, _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1042_v1).Select((Companion_Default___.SafeIndex((_this).Fm1((_this).F9(), _1260_v163, ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), globalState), _dafny.IntOfUint32((_1042_v1).Cardinality()))).Uint32()).(bool), p0), ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int)), 9)
			(_nw241).ArraySet1(Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), (_this).F7(), _1258_v161, _dafny.IntOfInt64(831)), 10)
			(_nw241).ArraySet1(_1259_v162, 11)
			(_nw241).ArraySet1(_1259_v162, 12)
			(_nw241).ArraySet1(_1259_v162, 13)
			(_nw241).ArraySet1(_1259_v162, 14)
			(_nw241).ArraySet1(Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), ((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), _1258_v161, _dafny.IntOfUint32(((_this).F5()).Cardinality())), 15)
			(_nw241).ArraySet1(Companion_D5_.Create_DC15_(_dafny.ArrayCastTo((_1236_v140).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(856), _dafny.ArrayLen((_1236_v140), 0))).Int())), Companion_Default___.Fm3(globalState), _1258_v161, _dafny.IntOfInt64(576)), 16)
			(_nw241).ArraySet1(_1259_v162, 17)
			(_nw241).ArraySet1(_1259_v162, 18)
			(_nw241).ArraySet1(_1259_v162, 19)
			(_nw241).ArraySet1(_1259_v162, 20)
			(_nw241).ArraySet1(_1259_v162, 21)
			_1261_v164 = _nw241
			var _index211 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_1261_v164), 0))
			_ = _index211
			(_1261_v164).ArraySet1(_1259_v162, (_index211).Int())
			var _1262_v165 D3
			_ = _1262_v165
			_1262_v165 = Companion_D3_.Create_DC11_(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool), ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int), _dafny.SetOf((_this).F7()))
			var _pat_let_tv20 = _1258_v161
			_ = _pat_let_tv20
			var _pat_let_tv21 = _1262_v165
			_ = _pat_let_tv21
			var _pat_let_tv22 = p0
			_ = _pat_let_tv22
			var _index212 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(457), _dafny.ArrayLen((_1261_v164), 0))
			_ = _index212
			(_1261_v164).ArraySet1(func(_pat_let39_0 D5) D5 {
				return func(_1265_dt__update__tmp_h5 D5) D5 {
					return func(_pat_let42_0 bool) D5 {
						return func(_1266_dt__update_hcf27_h0 bool) D5 {
							return func(_pat_let43_0 _dafny.Map) D5 {
								return func(_1267_dt__update_hcf28_h1 _dafny.Map) D5 {
									return Companion_D5_.Create_DC15_((_1265_dt__update__tmp_h5).Dtor_cf26(), _1266_dt__update_hcf27_h0, _1267_dt__update_hcf28_h1, (_1265_dt__update__tmp_h5).Dtor_cf29())
								}(_pat_let43_0)
							}((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_pat_let_tv21).Dtor_cf20(), ((_this).F11()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(162), _dafny.ArrayLen(((_this).F11()), 0))).Int()).(_dafny.Int))).Update(!(((_this).F10()).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(461), _dafny.ArrayLen(((_this).F10()), 0))).Int()).(bool)), _pat_let_tv22))
						}(_pat_let42_0)
					}(!((_this).F6()))
				}(_pat_let39_0)
			}(func(_pat_let40_0 D5) D5 {
				return func(_1263_dt__update__tmp_h4 D5) D5 {
					return func(_pat_let41_0 _dafny.Map) D5 {
						return func(_1264_dt__update_hcf28_h0 _dafny.Map) D5 {
							return Companion_D5_.Create_DC15_((_1263_dt__update__tmp_h4).Dtor_cf26(), (_1263_dt__update__tmp_h4).Dtor_cf27(), _1264_dt__update_hcf28_h0, (_1263_dt__update__tmp_h4).Dtor_cf29())
						}(_pat_let41_0)
					}(_pat_let_tv20)
				}(_pat_let40_0)
			}(_1259_v162)), (_index212).Int())
		}
		var _1268_v166 _dafny.MultiSet
		_ = _1268_v166
		_1268_v166 = _dafny.MultiSetOf(_dafny.IntOfInt64(537), p0)
		var _1269_v167 _dafny.Map
		_ = _1269_v167
		_1269_v167 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), (_1268_v166).Cardinality())
		r0 = (_1269_v167).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F8(), _dafny.IntOfInt64(66))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.MultiSetOf(!((_this).F7()))).Cardinality(), p0)))
		return r0
	}
}
func (_this *C5) M3(p0 _dafny.Int, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) bool {
	{
		var r0 bool = false
		_ = r0
		var _1270_v0 _dafny.MultiSet
		_ = _1270_v0
		_1270_v0 = _dafny.MultiSetOf((_this).F5(), (_this).F5(), (_this).F5())
		var _1271_v1 _dafny.Sequence
		_ = _1271_v1
		_1271_v1 = _dafny.SeqOf(_1270_v0, _1270_v0)
		var _1272_v2 _dafny.Map
		_ = _1272_v2
		_1272_v2 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1270_v0).IsProperSubsetOf((_1271_v1).Select((Companion_Default___.SafeIndex(p1, _dafny.IntOfUint32((_1271_v1).Cardinality()))).Uint32()).(_dafny.MultiSet)), _this.F4())
		_1272_v2 = (_1272_v2).Update(!((_this).F9()) || ((_this).F7()), _this.F4())
		var _hi5 _dafny.Int = p0
		_ = _hi5
		for _1273_i0 := p2; _1273_i0.Cmp(_hi5) < 0; _1273_i0 = _1273_i0.Plus(_dafny.One) {
			var _1274_v3 _dafny.Set
			_ = _1274_v3
			_1274_v3 = _dafny.SetOf((_this).F9())
			var _1275_v4 _dafny.Sequence
			_ = _1275_v4
			_1275_v4 = _dafny.SeqOf(_1274_v3, _1274_v3)
			var _1276_v5 _dafny.Sequence
			_ = _1276_v5
			_1276_v5 = _dafny.SeqOf((_this).F7())
			var _1277_v6 _dafny.Sequence
			_ = _1277_v6
			_1277_v6 = _dafny.SeqOf(!((_1274_v3).IsProperSubsetOf((_1275_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_1276_v5).Cardinality()), _dafny.IntOfUint32((_1275_v4).Cardinality()))).Uint32()).(_dafny.Set))))
			_1276_v5 = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1277_v6, (func() _dafny.Sequence {
				if (_this).F7() {
					return _1277_v6
				}
				return _dafny.SeqOf((_this).F6())
			})()), (Companion_Default___.SafeIndex(_dafny.IntOfInt64(637), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1277_v6, (func() _dafny.Sequence {
				if (_this).F7() {
					return _1277_v6
				}
				return _dafny.SeqOf((_this).F6())
			})())).Cardinality()))).Uint32(), (_this).F6())
			r0 = Companion_Default___.Fm7(_this.F4(), _1273_i0, (_this).F9(), globalState)
			var _1278_v7 D1
			_ = _1278_v7
			_1278_v7 = Companion_D1_.Create_DC4_(_this.F4(), _dafny.IntOfInt64(501), _this.F4(), p2, (_this).F6())
			var _1279_v8 *C0
			_ = _1279_v8
			var _nw242 *C0 = New_C0_()
			_ = _nw242
			_nw242.Ctor__((_1278_v7).Dtor_cf12(), (_1277_v6).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_1277_v6).Cardinality()))).Uint32()).(bool), _this.F4(), (_this).F5())
			_1279_v8 = _nw242
			var _1280_v9 _dafny.Sequence
			_ = _1280_v9
			_1280_v9 = _dafny.UnicodeSeqOfUtf8Bytes("jne")
			var _index213 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index213
			((_this).F11()).ArraySet1(p1, (_index213).Int())
			var _1281_v10 _dafny.MultiSet
			_ = _1281_v10
			_1281_v10 = _dafny.MultiSetOf((_this).F9())
			var _index214 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index214
			var _rhs199 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate((_this).F5(), (_this).F5())
			_ = _rhs199
			var _rhs200 bool = true
			_ = _rhs200
			var _rhs201 _dafny.Int = (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((Companion_Default___.Fm4((_this).F9(), _dafny.IntOfUint32((_1280_v9).Cardinality()), (_this).F9(), p0, globalState)).IsProperSubsetOf(_1281_v10), _dafny.Companion_Sequence_.Concatenate(_1276_v5, _dafny.SeqOf((_this).F7(), (_this).F9())))).Cardinality()
			_ = _rhs201
			var _rhs202 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("qxofl"), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(951))).Uint32(), func(coer60 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg60 _dafny.Int) interface{} {
					return coer60(arg60)
				}
			}(func(_1282_i1 _dafny.Int) _dafny.CodePoint {
				return _this.F4()
			})))
			_ = _rhs202
			var _lhs152 _dafny.Array = (_this).F11()
			_ = _lhs152
			var _lhs153 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(596), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _lhs153
			_1280_v9 = _rhs199
			r0 = _rhs200
			(_lhs152).ArraySet1(_rhs201, (_lhs153).Int())
			_1280_v9 = _rhs202
		}
		r0 = _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(414))).Uint32(), func(coer61 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg61 _dafny.Int) interface{} {
				return coer61(arg61)
			}
		}(func(_1283_i2 _dafny.Int) _dafny.CodePoint {
			return _this.F4()
		})), (_this).F5())
		var _hi6 _dafny.Int = _dafny.IntOfUint32((_dafny.SeqOf((_this).F8())).Cardinality())
		_ = _hi6
		for _1284_i3 := p1; _1284_i3.Cmp(_hi6) < 0; _1284_i3 = _1284_i3.Plus(_dafny.One) {
			var _1285_v11 _dafny.Array
			_ = _1285_v11
			var _nw243 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(6))
			_ = _nw243
			_1285_v11 = _nw243
			_1285_v11 = _1285_v11
			var _1286_v12 _dafny.Set
			_ = _1286_v12
			_1286_v12 = _dafny.SetOf(_dafny.IntOfInt64(77))
			var _1287_v13 _dafny.Map
			_ = _1287_v13
			_1287_v13 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_1286_v12).IsProperSubsetOf(_1286_v12)), (_this).F5())
			_1287_v13 = (_1287_v13).Update((_this).F7(), (_this).F5())
			var _1288_v14 _dafny.Set
			_ = _1288_v14
			_1288_v14 = _dafny.SetOf(_this.F4())
			var _1289_v15 _dafny.Sequence
			_ = _1289_v15
			_1289_v15 = _dafny.SeqOf((_1288_v14).Difference(_1288_v14), _1288_v14, _1288_v14, (_1288_v14).Union(_1288_v14), _1288_v14)
			(globalState).F1 = ((_1289_v15).Select((Companion_Default___.SafeIndex(_1284_i3, _dafny.IntOfUint32((_1289_v15).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality()
			var _1290_v16 _dafny.Sequence
			_ = _1290_v16
			_1290_v16 = _dafny.SeqOf((_this).F9(), _dafny.Companion_Sequence_.IsProperPrefixOf((_this).F5(), (_this).F5()))
			_1290_v16 = _dafny.Companion_Sequence_.Update(_1290_v16, (Companion_Default___.SafeIndex((_1284_i3).Minus((_this).F8()), _dafny.IntOfUint32((_1290_v16).Cardinality()))).Uint32(), Companion_Default___.Fm27(p2, !((_this).F7()), (_this).F9(), globalState))
		}
		var _1291_v17 *C1
		_ = _1291_v17
		var _nw244 *C1 = New_C1_()
		_ = _nw244
		_nw244.Ctor__(p2, false, Companion_Default___.Fm27((_this).F8(), (_this).F6(), (_this).F9(), globalState), (_this).F9(), _this.F4(), (_this).F5())
		_1291_v17 = _nw244
		var _1292_v18 D10
		_ = _1292_v18
		_1292_v18 = Companion_D10_.Create_DC28_((_this).F7(), p2, (_this).F6(), (_this).F9(), _1291_v17)
		var _1293_v19 _dafny.MultiSet
		_ = _1293_v19
		_1293_v19 = _dafny.MultiSetOf((_1292_v18).Dtor_cf51())
		(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeDivisionInt((_1293_v19).Cardinality(), p1))
		var _1294_i4 _dafny.Int
		_ = _1294_i4
		_1294_i4 = _dafny.Zero
		{
			for (_this).F7() {
				{
					if (_1294_i4).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_1294_i4 = (_1294_i4).Plus(_dafny.One)
					var _1295_v20 _dafny.Array
					_ = _1295_v20
					var _nw245 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(9))
					_ = _nw245
					_1295_v20 = _nw245
					var _1296_v21 _dafny.Sequence
					_ = _1296_v21
					_1296_v21 = _dafny.SeqOf(false)
					var _1297_v22 _dafny.Sequence
					_ = _1297_v22
					_1297_v22 = _dafny.SeqOf(_1296_v21)
					var _index215 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1295_v20), 0))
					_ = _index215
					(_1295_v20).ArraySet1(_1297_v22, (_index215).Int())
					var _1298_v23 _dafny.Set
					_ = _1298_v23
					_1298_v23 = _dafny.SetOf((_this).F9(), (_this).F6(), (_this).F7(), (_this).F9(), (_this).F9())
					var _1299_v24 D8
					_ = _1299_v24
					_1299_v24 = Companion_D8_.Create_DC22_(_1296_v21)
					var _index216 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1295_v20), 0))
					_ = _index216
					var _rhs203 _dafny.Sequence = Companion_Default___.Fm34(!((_this).F9()), _1299_v24, Companion_Default___.SafeDivisionInt((_this).F8(), Companion_Default___.Fm8(p2, (_this).F7(), (_this).F6(), globalState)), globalState)
					_ = _rhs203
					var _rhs204 _dafny.Int = (_dafny.IntOfUint32((_1296_v21).Cardinality())).Plus(p2)
					_ = _rhs204
					var _rhs205 _dafny.Set = _1298_v23
					_ = _rhs205
					var _rhs206 bool = (_this).F6()
					_ = _rhs206
					var _lhs154 _dafny.Array = _1295_v20
					_ = _lhs154
					var _lhs155 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(67), _dafny.ArrayLen((_1295_v20), 0))
					_ = _lhs155
					var _lhs156 *GlobalState = globalState
					_ = _lhs156
					(_lhs154).ArraySet1(_rhs203, (_lhs155).Int())
					_lhs156.F1 = _rhs204
					_1298_v23 = _rhs205
					r0 = _rhs206
					var _1300_v25 _dafny.Sequence
					_ = _1300_v25
					_1300_v25 = _dafny.UnicodeSeqOfUtf8Bytes("etoajigf")
					_1300_v25 = _dafny.UnicodeSeqOfUtf8Bytes("a")
					(_this).F4_set_(_this.F4())
					var _1301_v26 _dafny.Sequence
					_ = _1301_v26
					_1301_v26 = _dafny.SeqOf((_dafny.MultiSetOf((_this).F9())).Update((_this).F9(), Companion_Default___.Abs(p0)), _dafny.MultiSetOf((_this).F6(), (_this).F7(), (_this).F6()))
					var _1302_v27 D3
					_ = _1302_v27
					_1302_v27 = Companion_D3_.Create_DC11_((_this).F9(), p0, _1298_v23)
					var _1303_v28 _dafny.Array
					_ = _1303_v28
					var _len0_25 _dafny.Int = _dafny.IntOfInt64(28)
					_ = _len0_25
					var _nw246 _dafny.Array
					_ = _nw246
					if _len0_25.Cmp(_dafny.Zero) == 0 {
						_nw246 = _dafny.NewArray(_len0_25)
					} else {
						var _init25 func(_dafny.Int) _dafny.CodePoint = func(_1304_i5 _dafny.Int) _dafny.CodePoint {
							return _this.F4()
						}
						_ = _init25
						var _element0_25 = _init25(_dafny.Zero)
						_ = _element0_25
						_nw246 = _dafny.NewArrayFromExample(_element0_25, nil, _len0_25)
						(_nw246).ArraySet1CodePoint(_element0_25, 0)
						var _nativeLen0_25 = (_len0_25).Int()
						_ = _nativeLen0_25
						for _i0_25 := 1; _i0_25 < _nativeLen0_25; _i0_25++ {
							(_nw246).ArraySet1CodePoint(_init25(_dafny.IntOf(_i0_25)), _i0_25)
						}
					}
					_1303_v28 = _nw246
					var _index217 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1303_v28), 0))
					_ = _index217
					(_1303_v28).ArraySet1CodePoint(((_this).F5()).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32(((_this).F5()).Cardinality()))).Uint32()).(_dafny.CodePoint), (_index217).Int())
					var _1305_v29 _dafny.Array
					_ = _1305_v29
					var _nw247 _dafny.Array = _dafny.NewArrayWithValue(Companion_D1_.Default(), _dafny.IntOfInt64(6))
					_ = _nw247
					_1305_v29 = _nw247
					var _1306_v30 *C2
					_ = _1306_v30
					var _nw248 *C2 = New_C2_()
					_ = _nw248
					_nw248.Ctor__((_this).F10(), _1305_v29)
					_1306_v30 = _nw248
					var _1307_v31 _dafny.Map
					_ = _1307_v31
					_1307_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F7(), (_this).F9())
					var _1308_v32 _dafny.MultiSet
					_ = _1308_v32
					_1308_v32 = _dafny.MultiSetOf((func() bool {
						if (_1307_v31).Contains((_this).F7()) {
							return (_1307_v31).Get((_this).F7()).(bool)
						}
						return true
					})(), (_this).F6(), (_this).F6(), !((_this).F9()), (_this).F7())
					var _1309_v33 _dafny.Map
					_ = _1309_v33
					_1309_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, p2)
					var _pat_let_tv23 = _1298_v23
					_ = _pat_let_tv23
					var _pat_let_tv24 = globalState
					_ = _pat_let_tv24
					var _index218 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1303_v28), 0))
					_ = _index218
					var _rhs207 bool = !((_this).F6())
					_ = _rhs207
					var _rhs208 _dafny.Sequence = _dafny.Companion_Sequence_.Update(_dafny.Companion_Sequence_.Concatenate(_1301_v26, _dafny.SeqOf(_dafny.MultiSetFromSeq(_1296_v21), _1308_v32)), (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_1301_v26, _dafny.SeqOf(_dafny.MultiSetFromSeq(_1296_v21), _1308_v32))).Cardinality()))).Uint32(), _1308_v32)
					_ = _rhs208
					var _rhs209 D3 = func(_pat_let44_0 D3) D3 {
						return func(_1312_dt__update__tmp_h1 D3) D3 {
							return func(_pat_let47_0 bool) D3 {
								return func(_1313_dt__update_hcf20_h0 bool) D3 {
									return Companion_D3_.Create_DC11_(_1313_dt__update_hcf20_h0, (_1312_dt__update__tmp_h1).Dtor_cf21(), (_1312_dt__update__tmp_h1).Dtor_cf22())
								}(_pat_let47_0)
							}(Companion_Default___.Fm3(_pat_let_tv24))
						}(_pat_let44_0)
					}(func(_pat_let45_0 D3) D3 {
						return func(_1310_dt__update__tmp_h0 D3) D3 {
							return func(_pat_let46_0 _dafny.Set) D3 {
								return func(_1311_dt__update_hcf22_h0 _dafny.Set) D3 {
									return Companion_D3_.Create_DC11_((_1310_dt__update__tmp_h0).Dtor_cf20(), (_1310_dt__update__tmp_h0).Dtor_cf21(), _1311_dt__update_hcf22_h0)
								}(_pat_let46_0)
							}(_pat_let_tv23)
						}(_pat_let45_0)
					}(Companion_D3_.Create_DC11_((_this).F9(), (_1309_v33).Cardinality(), _1298_v23)))
					_ = _rhs209
					var _rhs210 _dafny.CodePoint = _this.F4()
					_ = _rhs210
					var _rhs211 *C2 = _1306_v30
					_ = _rhs211
					var _lhs157 _dafny.Array = _1303_v28
					_ = _lhs157
					var _lhs158 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(133), _dafny.ArrayLen((_1303_v28), 0))
					_ = _lhs158
					r0 = _rhs207
					_1301_v26 = _rhs208
					_1302_v27 = _rhs209
					(_lhs157).ArraySet1CodePoint(_rhs210, (_lhs158).Int())
					_1306_v30 = _rhs211
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		r0 = !((func() bool {
			if false {
				return (_this).F7()
			}
			return (_this).F9()
		})())
		return r0
	}
}
func (_this *C5) M4(p0 _dafny.Int, p1 _dafny.Map, globalState *GlobalState) {
	{
		var _1314_v0 T2
		_ = _1314_v0
		var _nw249 *C4 = New_C4_()
		_ = _nw249
		_nw249.Ctor__((_this).F8(), (_this).F6(), (_this).F9(), (_this).F7(), _this.F4(), _dafny.UnicodeSeqOfUtf8Bytes("vknfksjs"))
		_1314_v0 = _nw249
		var _1315_v1 _dafny.Map
		_ = _1315_v1
		_1315_v1 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1314_v0, (_this).F8())
		var _1316_v2 D2
		_ = _1316_v2
		_1316_v2 = Companion_D2_.Create_DC8_(p0, (_1315_v1).Cardinality())
		var _pat_let_tv25 = _1314_v0
		_ = _pat_let_tv25
		var _pat_let_tv26 = _1314_v0
		_ = _pat_let_tv26
		var _pat_let_tv27 = _1314_v0
		_ = _pat_let_tv27
		(globalState).F1 = func(_source15 D2) _dafny.Int {
			if _source15.Is_DC7() {
				var _1317___mcc_h0 bool = _source15.Get_().(D2_DC7).Cf15
				_ = _1317___mcc_h0
				var _1318_cf15 bool = _1317___mcc_h0
				_ = _1318_cf15
				return (_this).F8()
			} else if _source15.Is_DC8() {
				var _1319___mcc_h1 _dafny.Int = _source15.Get_().(D2_DC8).Cf16
				_ = _1319___mcc_h1
				var _1320___mcc_h2 _dafny.Int = _source15.Get_().(D2_DC8).Cf17
				_ = _1320___mcc_h2
				var _1321_cf17 _dafny.Int = _1320___mcc_h2
				_ = _1321_cf17
				var _1322_cf16 _dafny.Int = _1319___mcc_h1
				_ = _1322_cf16
				return _1321_cf17
			} else if _source15.Is_DC6() {
				var _1323___mcc_h3 _dafny.Map = _source15.Get_().(D2_DC6).Cf14
				_ = _1323___mcc_h3
				var _1324_cf14 _dafny.Map = _1323___mcc_h3
				_ = _1324_cf14
				return (_dafny.Zero).Minus(_dafny.IntOfUint32(((_pat_let_tv25).F5()).Cardinality()))
			} else {
				var _1325___mcc_h4 D2 = _source15.Get_().(D2_DC9).Cf18
				_ = _1325___mcc_h4
				var _1326_cf18 D2 = _1325___mcc_h4
				_ = _1326_cf18
				return ((_dafny.MultiSetOf((_pat_let_tv26).F7())).Difference(_dafny.MultiSetOf((_pat_let_tv27).F6()))).Cardinality()
			}
		}(_1316_v2)
		var _1327_v3 _dafny.MultiSet
		_ = _1327_v3
		_1327_v3 = _dafny.MultiSetOf(false, (_1314_v0).F7())
		var _1328_i0 _dafny.Int
		_ = _1328_i0
		_1328_i0 = _dafny.Zero
		{
			for (_1327_v3).IsDisjointFrom(Companion_Default___.Fm4((_1314_v0).F6(), (_this).F8(), (_1314_v0).F6(), p0, globalState)) {
				{
					if (_1328_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L6
					}
					_1328_i0 = (_1328_i0).Plus(_dafny.One)
					var _1329_v4 _dafny.Map
					_ = _1329_v4
					_1329_v4 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F9(), (_1314_v0).F9())
					var _1330_v5 D12
					_ = _1330_v5
					_1330_v5 = Companion_D12_.Create_DC31_((_1329_v4).Update((_this).F6(), (_1314_v0).F7()))
					(globalState).F1 = Companion_Default___.SafeModuloInt((_this).Fm1((_this).F6(), (_1330_v5).Dtor_cf60(), (_1314_v0).F6(), globalState), Companion_Default___.SafeModuloInt((_dafny.SetOf(true)).Cardinality(), (_this).F8()))
					var _1331_v6 _dafny.Sequence
					_ = _1331_v6
					_1331_v6 = _dafny.SeqOf((_this).F8(), (_1314_v0).Fm1((_1314_v0).F9(), _1329_v4, (_this).F7(), globalState), Companion_Default___.SafeModuloInt((_1329_v4).Cardinality(), (_1314_v0).F8()), ((_1314_v0).F8()).Times(p0), (_this).F8())
					_1331_v6 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_1331_v6, _1331_v6), _1331_v6)
					var _index219 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen(((_this).F10()), 0))
					_ = _index219
					((_this).F10()).ArraySet1((_1314_v0).F7(), (_index219).Int())
					var _index220 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen(((_this).F10()), 0))
					_ = _index220
					((_this).F10()).ArraySet1((_1314_v0).F9(), (_index220).Int())
					var _1332_v7 _dafny.Map
					_ = _1332_v7
					_1332_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1331_v6, true)
					var _index221 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(374), _dafny.ArrayLen(((_this).F10()), 0))
					_ = _index221
					((_this).F10()).ArraySet1(!(_1332_v7).Equals((func() _dafny.Map {
						if Companion_Default___.Fm27((_1314_v0).F8(), (_this).F9(), false, globalState) {
							return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1331_v6, (_1314_v0).F7())
						}
						return _1332_v7
					})()), (_index221).Int())
					goto C6
				}
			C6:
			}
			goto L6
		}
	L6:
		var _1333_v8 _dafny.Array
		_ = _1333_v8
		var _nw250 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(22))
		_ = _nw250
		_1333_v8 = _nw250
		var _nw251 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(27))
		_ = _nw251
		_1333_v8 = _nw251
		var _1334_v9 *C1
		_ = _1334_v9
		var _nw252 *C1 = New_C1_()
		_ = _nw252
		_nw252.Ctor__((_1314_v0).F8(), (_this).F9(), (_this).F9(), (_1314_v0).F7(), _1314_v0.F4(), _dafny.UnicodeSeqOfUtf8Bytes("gedhynvk"))
		_1334_v9 = _nw252
		var _1335_v10 _dafny.MultiSet
		_ = _1335_v10
		_1335_v10 = _dafny.MultiSetOf((Companion_D10_.Create_DC28_((_this).F9(), (_this).F8(), (_1314_v0).F7(), !((_this).F7()), _1334_v9)).Dtor_cf51())
		var _hi7 _dafny.Int = _dafny.IntOfInt64(104)
		_ = _hi7
		for _1336_i1 := (_1335_v10).Cardinality(); _1336_i1.Cmp(_hi7) < 0; _1336_i1 = _1336_i1.Plus(_dafny.One) {
			var _1337_v11 _dafny.Map
			_ = _1337_v11
			_1337_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1314_v0).F8(), (_1314_v0).F8())
			_1337_v11 = (_1337_v11).Update(p0, (_dafny.Zero).Minus((_1314_v0).F8()))
			if (_1314_v0).F6() {
				var _1338_v12 bool
				_ = _1338_v12
				_1338_v12 = true
				var _1339_v13 D12
				_ = _1339_v13
				_1339_v13 = Companion_D12_.Create_DC33_(_1338_v12, _dafny.UnicodeSeqOfUtf8Bytes("o"))
				var _1340_v14 _dafny.Map
				_ = _1340_v14
				_1340_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, !((_1314_v0).F7()))
				var _pat_let_tv28 = _1314_v0
				_ = _pat_let_tv28
				var _pat_let_tv29 = globalState
				_ = _pat_let_tv29
				var _pat_let_tv30 = _1340_v14
				_ = _pat_let_tv30
				var _pat_let_tv31 = _1314_v0
				_ = _pat_let_tv31
				var _pat_let_tv32 = globalState
				_ = _pat_let_tv32
				var _pat_let_tv33 = _1314_v0
				_ = _pat_let_tv33
				var _pat_let_tv34 = globalState
				_ = _pat_let_tv34
				_1338_v12 = (func(_pat_let48_0 D12) D12 {
					return func(_1341_dt__update__tmp_h0 D12) D12 {
						return func(_pat_let49_0 bool) D12 {
							return func(_1342_dt__update_hcf63_h0 bool) D12 {
								return Companion_D12_.Create_DC33_(_1342_dt__update_hcf63_h0, (_1341_dt__update__tmp_h0).Dtor_cf64())
							}(_pat_let49_0)
						}(Companion_Default___.Fm7(_pat_let_tv28.F4(), (_this).Fm1(Companion_Default___.Fm3(_pat_let_tv29), _pat_let_tv30, (_pat_let_tv31).F6(), _pat_let_tv32), (_pat_let_tv33).F7(), _pat_let_tv34))
					}(_pat_let48_0)
				}(_1339_v13)).Dtor_cf63()
				var _1343_v15 _dafny.Set
				_ = _1343_v15
				var _1344_v16 _dafny.Array
				_ = _1344_v16
				var _1345_v17 _dafny.Int
				_ = _1345_v17
				var _1346_v18 _dafny.Array
				_ = _1346_v18
				var _out25 _dafny.Set
				_ = _out25
				var _out26 _dafny.Array
				_ = _out26
				var _out27 _dafny.Int
				_ = _out27
				var _out28 _dafny.Array
				_ = _out28
				_out25, _out26, _out27, _out28 = (_1334_v9).M8((func() bool {
					if (_this).F7() {
						return (_1314_v0).F6()
					}
					return (_this).F6()
				})(), (_1314_v0).F7(), globalState)
				_1343_v15 = _out25
				_1344_v16 = _out26
				_1345_v17 = _out27
				_1346_v18 = _out28
				var _1347_v19 _dafny.Array
				_ = _1347_v19
				var _nw253 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(17))
				_ = _nw253
				_1347_v19 = _nw253
				var _nw254 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(16))
				_ = _nw254
				_1347_v19 = _nw254
				_1338_v12 = _1338_v12
				var _index222 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1346_v18), 0))
				_ = _index222
				(_1346_v18).ArraySet1((_this).F7(), (_index222).Int())
				var _index223 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(737), _dafny.ArrayLen((_1346_v18), 0))
				_ = _index223
				(_1346_v18).ArraySet1((_1314_v0).F7(), (_index223).Int())
			} else {
				var _1348_v20 bool
				_ = _1348_v20
				_1348_v20 = true
				var _rhs212 bool = !((func() bool {
					if (p1).Contains((_this).F8()) {
						return (p1).Get((_this).F8()).(bool)
					}
					return (_1314_v0).F6()
				})()) || ((_1314_v0).F7())
				_ = _rhs212
				var _rhs213 bool = !((_1314_v0).F7())
				_ = _rhs213
				_1348_v20 = _rhs212
				_1348_v20 = _rhs213
				var _rhs214 _dafny.Int = (_this).F8()
				_ = _rhs214
				var _rhs215 T2 = _1314_v0
				_ = _rhs215
				var _lhs159 *GlobalState = globalState
				_ = _lhs159
				_lhs159.F1 = _rhs214
				_1314_v0 = _rhs215
				_1348_v20 = true
				var _1349_v21 _dafny.Map
				_ = _1349_v21
				_1349_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1348_v20, (_this).F7())
				(globalState).F1 = (_1334_v9).Fm1(true, _1349_v21, (_this).F6(), globalState)
				var _1350_v22 _dafny.Map
				_ = _1350_v22
				_1350_v22 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1314_v0).F8(), true)
				var _1351_v23 _dafny.Sequence
				_ = _1351_v23
				_1351_v23 = _dafny.SeqOf((_1314_v0).F9())
				_1350_v22 = (_1350_v22).Update(((_1337_v11).Merge(_1337_v11)).Cardinality(), ((_1351_v23).Select((Companion_Default___.SafeIndex((_dafny.Zero).Minus((_1337_v11).Cardinality()), _dafny.IntOfUint32((_1351_v23).Cardinality()))).Uint32()).(bool)) || (!((_1314_v0).F7())))
			}
			(globalState).F1 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(((_1335_v10).Difference(_1335_v10)).Cardinality(), _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer62 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg62 _dafny.Int) interface{} {
					return coer62(arg62)
				}
			}((func(_1352_v0 T2) func(_dafny.Int) _dafny.CodePoint {
				return func(_1353_i2 _dafny.Int) _dafny.CodePoint {
					return _1352_v0.F4()
				}
			})(_1314_v0))), (Companion_Default___.SafeIndex(Companion_Default___.Fm8((_this).F8(), (_this).F9(), (_1314_v0).F6(), globalState), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(603))).Uint32(), func(coer63 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg63 _dafny.Int) interface{} {
					return coer63(arg63)
				}
			}((func(_1354_v0 T2) func(_dafny.Int) _dafny.CodePoint {
				return func(_1355_i2 _dafny.Int) _dafny.CodePoint {
					return _1354_v0.F4()
				}
			})(_1314_v0)))).Cardinality()))).Uint32(), _1314_v0.F4())).Cardinality())))
			var _index224 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index224
			((_this).F10()).ArraySet1((_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bhpixis")).Cardinality())).Cmp(_1336_i1) != 0, (_index224).Int())
			var _1356_v24 bool
			_ = _1356_v24
			_1356_v24 = true
			var _1357_v25 _dafny.Array
			_ = _1357_v25
			var _nw255 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(6))
			_ = _nw255
			_1357_v25 = _nw255
			var _1358_v26 _dafny.Sequence
			_ = _1358_v26
			_1358_v26 = _dafny.SeqOf((_1314_v0).F7(), (_this).F7())
			var _1359_v27 _dafny.Map
			_ = _1359_v27
			_1359_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1314_v0).F6(), _dafny.IntOfUint32((_1358_v26).Cardinality()))
			var _1360_v28 D5
			_ = _1360_v28
			_1360_v28 = Companion_D5_.Create_DC15_((_this).F10(), (_1314_v0).F9(), _1359_v27, _dafny.IntOfInt64(368))
			var _1361_v29 _dafny.Sequence
			_ = _1361_v29
			_1361_v29 = _dafny.SeqOf((_this).F10(), (_this).F10(), (_this).F10(), (_this).F10())
			var _1362_v30 _dafny.Array
			_ = _1362_v30
			var _nwElement0_52 _dafny.Array = (_this).F10()
			_ = _nwElement0_52
			var _nw256 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_52, nil, _dafny.IntOfInt64(29))
			_ = _nw256
			(_nw256).ArraySet1(_nwElement0_52, 0)
			(_nw256).ArraySet1((_this).F10(), 1)
			(_nw256).ArraySet1((_this).F10(), 2)
			(_nw256).ArraySet1((_this).F10(), 3)
			(_nw256).ArraySet1((_1360_v28).Dtor_cf26(), 4)
			(_nw256).ArraySet1((_this).F10(), 5)
			(_nw256).ArraySet1((_this).F10(), 6)
			(_nw256).ArraySet1((_this).F10(), 7)
			(_nw256).ArraySet1((_this).F10(), 8)
			(_nw256).ArraySet1((_1361_v29).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1361_v29).Cardinality()))).Uint32()).(_dafny.Array), 9)
			(_nw256).ArraySet1((_this).F10(), 10)
			(_nw256).ArraySet1((_this).F10(), 11)
			(_nw256).ArraySet1((_this).F10(), 12)
			(_nw256).ArraySet1((_this).F10(), 13)
			(_nw256).ArraySet1((_this).F10(), 14)
			(_nw256).ArraySet1((_this).F10(), 15)
			(_nw256).ArraySet1((_this).F10(), 16)
			(_nw256).ArraySet1((_1361_v29).Select((Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_1361_v29).Cardinality()))).Uint32()).(_dafny.Array), 17)
			(_nw256).ArraySet1((_this).F10(), 18)
			(_nw256).ArraySet1((_this).F10(), 19)
			(_nw256).ArraySet1((_this).F10(), 20)
			(_nw256).ArraySet1((_this).F10(), 21)
			(_nw256).ArraySet1((_this).F10(), 22)
			(_nw256).ArraySet1((_this).F10(), 23)
			(_nw256).ArraySet1((_this).F10(), 24)
			(_nw256).ArraySet1((_this).F10(), 25)
			(_nw256).ArraySet1((_this).F10(), 26)
			(_nw256).ArraySet1((_this).F10(), 27)
			(_nw256).ArraySet1((_this).F10(), 28)
			_1362_v30 = _nw256
			var _index225 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_1357_v25), 0))
			_ = _index225
			(_1357_v25).ArraySet1(_1362_v30, (_index225).Int())
			var _1363_v31 _dafny.Sequence
			_ = _1363_v31
			_1363_v31 = _dafny.SeqOf((_this).F8())
			var _1364_v32 _dafny.Sequence
			_ = _1364_v32
			_1364_v32 = _dafny.SeqOf(_1363_v31)
			var _index226 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index226
			((_this).F11()).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_dafny.SeqOf(p0, (_this).F8()), (Companion_Default___.SafeIndex((_this).F8(), _dafny.IntOfUint32((_dafny.SeqOf(p0, (_this).F8())).Cardinality()))).Uint32(), _dafny.IntOfUint32((_1364_v32).Cardinality()))).Cardinality()), (_index226).Int())
			var _1365_v33 _dafny.Map
			_ = _1365_v33
			_1365_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F6(), false)
			var _1366_v34 D8
			_ = _1366_v34
			_1366_v34 = Companion_D8_.Create_DC23_((_1334_v9).Fm1((_1314_v0).F9(), _1365_v33, (_1314_v0).F6(), globalState), (_this).F9())
			var _1367_v35 _dafny.Set
			_ = _1367_v35
			_1367_v35 = _dafny.SetOf((_this).F8())
			var _index227 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _index227
			var _index228 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_1357_v25), 0))
			_ = _index228
			var _index229 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _index229
			var _rhs216 bool = (func() bool {
				if false {
					return (_1314_v0).F9()
				}
				return (_1366_v34).Dtor_cf43()
			})()
			_ = _rhs216
			var _rhs217 bool = Companion_Default___.Fm27((_dafny.Zero).Minus((_1335_v10).Cardinality()), (_1367_v35).IsSubsetOf(_1367_v35), (_this).F7(), globalState)
			_ = _rhs217
			var _rhs218 _dafny.Array = _1362_v30
			_ = _rhs218
			var _rhs219 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfInt64(-822))
			_ = _rhs219
			var _lhs160 _dafny.Array = (_this).F10()
			_ = _lhs160
			var _lhs161 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(446), _dafny.ArrayLen(((_this).F10()), 0))
			_ = _lhs161
			var _lhs162 _dafny.Array = _1357_v25
			_ = _lhs162
			var _lhs163 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(901), _dafny.ArrayLen((_1357_v25), 0))
			_ = _lhs163
			var _lhs164 _dafny.Array = (_this).F11()
			_ = _lhs164
			var _lhs165 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(779), _dafny.ArrayLen(((_this).F11()), 0))
			_ = _lhs165
			(_lhs160).ArraySet1(_rhs216, (_lhs161).Int())
			_1356_v24 = _rhs217
			(_lhs162).ArraySet1(_rhs218, (_lhs163).Int())
			(_lhs164).ArraySet1(_rhs219, (_lhs165).Int())
		}
		var _index230 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen(((_this).F11()), 0))
		_ = _index230
		((_this).F11()).ArraySet1(((_dafny.Zero).Minus(p0)).Plus((_this).F8()), (_index230).Int())
		var _1368_v36 _dafny.Set
		_ = _1368_v36
		_1368_v36 = _dafny.SetOf((_1314_v0).F8(), (_1314_v0).F8(), (_1314_v0).F8(), p0, _dafny.IntOfInt64(852))
		var _1369_v37 _dafny.Map
		_ = _1369_v37
		_1369_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_1314_v0).F8(), _1335_v10)
		var _1370_v38 _dafny.Map
		_ = _1370_v38
		_1370_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_1368_v36, (Companion_D9_.Create_DC25_((_this).F9(), (_this).F6(), (_1369_v37).Cardinality())).Dtor_cf47())
		var _index231 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(340), _dafny.ArrayLen(((_this).F11()), 0))
		_ = _index231
		((_this).F11()).ArraySet1((func() _dafny.Int {
			if (_1370_v38).Contains(_dafny.SetOf((_this).F8(), (_1314_v0).F8(), (_this).F8(), _dafny.IntOfUint32(((_this).F5()).Cardinality()), p0)) {
				return (_1370_v38).Get(_dafny.SetOf((_this).F8(), (_1314_v0).F8(), (_this).F8(), _dafny.IntOfUint32(((_this).F5()).Cardinality()), p0)).(_dafny.Int)
			}
			return Companion_Default___.SafeModuloInt(p0, (_1314_v0).F8())
		})(), (_index231).Int())
		var _1371_v39 bool
		_ = _1371_v39
		_1371_v39 = false
		_1371_v39 = !((_this).F6())
	}
}
func (_this *C5) F10() _dafny.Array {
	{
		return _this._f10
	}
}
func (_this *C5) F11() _dafny.Array {
	{
		return _this._f11
	}
}

// End of class C5
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
