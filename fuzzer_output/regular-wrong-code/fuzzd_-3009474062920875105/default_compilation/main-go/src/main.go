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
func (_static *CompanionStruct_Default___) Fm0(p0 _dafny.Map, p1 bool, p2 bool, globalState *GlobalState) _dafny.Int {
	return _dafny.IntOfInt64(-184)
}
func (_static *CompanionStruct_Default___) Fm3(p0 _dafny.Int, globalState *GlobalState) D1 {
	var _source0 D2 = Companion_D2_.Create_DC7_(false, false, true)
	_ = _source0
	if _source0.Is_DC7() {
		var _0___mcc_h0 bool = _source0.Get_().(D2_DC7).Cf11
		_ = _0___mcc_h0
		var _1___mcc_h1 bool = _source0.Get_().(D2_DC7).Cf12
		_ = _1___mcc_h1
		var _2___mcc_h2 bool = _source0.Get_().(D2_DC7).Cf13
		_ = _2___mcc_h2
		var _3_cf13 bool = _2___mcc_h2
		_ = _3_cf13
		var _4_cf12 bool = _1___mcc_h1
		_ = _4_cf12
		var _5_cf11 bool = _0___mcc_h0
		_ = _5_cf11
		return Companion_D1_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("ljgfacm"))
	} else {
		var _6___mcc_h3 _dafny.Sequence = _source0.Get_().(D2_DC6).Cf10
		_ = _6___mcc_h3
		var _7_cf10 _dafny.Sequence = _6___mcc_h3
		_ = _7_cf10
		if !(false) {
			return Companion_D1_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("g"))
		} else {
			return Companion_D1_.Create_DC4_(_dafny.UnicodeSeqOfUtf8Bytes("dewv"))
		}
	}
}
func (_static *CompanionStruct_Default___) Fm4(p0 bool, p1 _dafny.Set, p2 _dafny.Int, globalState *GlobalState) bool {
	return true
}
func (_static *CompanionStruct_Default___) Fm6(p0 bool, p1 _dafny.Int, globalState *GlobalState) bool {
	return _dafny.Companion_Sequence_.Contains(_dafny.SeqOf((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _dafny.IntOfInt64(-430))).Cardinality()), ((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, func() _dafny.Map {
		var _coll0 = _dafny.NewMapBuilder()
		_ = _coll0
		for _iter0 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-548), _dafny.IntOfInt64(187))); ; {
			_compr_0, _ok0 := _iter0()
			if !_ok0 {
				break
			}
			var _8_v0 _dafny.Int
			_8_v0 = interface{}(_compr_0).(_dafny.Int)
			if ((_dafny.IntOfInt64(-548)).Cmp(_8_v0) <= 0) && ((_8_v0).Cmp(_dafny.IntOfInt64(187)) < 0) {
				_coll0.Add((_8_v0).Times(_dafny.IntOfInt64(953)), _dafny.IntOfInt64(-372))
			}
		}
		return _coll0.ToMap()
	}())).Cardinality())).Cardinality()).Minus((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)).Cardinality()))
}
func (_static *CompanionStruct_Default___) Fm7(globalState *GlobalState) D0 {
	return Companion_D0_.Create_DC2_(_dafny.SetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sfekhaf")).Cardinality()), _dafny.IntOfInt64(774), _dafny.IntOfInt64(-567), _dafny.IntOfInt64(402)), ((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-165)))).Union(_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('d'))).Cardinality()), (_dafny.MultiSetOf(false)).Cardinality(), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-3))).Uint32(), func(coer0 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg0 _dafny.Int) interface{} {
			return coer0(arg0)
		}
	}(func(_9_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('b')
	}))).Cardinality()), _dafny.IntOfInt64(740)))).Cardinality(), false)
}
func (_static *CompanionStruct_Default___) Fm8(p0 bool, p1 bool, p2 bool, p3 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(853)), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ks")).Cardinality()))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.IntOfInt64(121), _dafny.IntOfInt64(889), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("xikkpo")).Cardinality()), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(258))).Cardinality()), (_dafny.MultiSetOf(!(!(true)), true, true, !(true))).Cardinality()), _dafny.SeqOf(_dafny.IntOfInt64(379))))
}
func (_static *CompanionStruct_Default___) Fm9(p0 _dafny.CodePoint, p1 _dafny.Int, p2 _dafny.Int, p3 _dafny.Map, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqOf(_dafny.IntOfInt64(832), _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(242))).Uint32(), func(coer1 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
		return func(arg1 _dafny.Int) interface{} {
			return coer1(arg1)
		}
	}(func(_10_i0 _dafny.Int) _dafny.CodePoint {
		return _dafny.CodePoint('p')
	}))).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfInt64(-61))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(510))).Uint32(), func(coer2 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
		return func(arg2 _dafny.Int) interface{} {
			return coer2(arg2)
		}
	}(func(_11_i1 _dafny.Int) _dafny.Int {
		return _dafny.IntOfInt64(258)
	}))))
}
func (_static *CompanionStruct_Default___) Fm10(p0 bool, p1 _dafny.CodePoint, p2 _dafny.Int, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(false, !(true)), _dafny.SeqOf((Companion_D0_.Create_DC1_(_dafny.SetOf(_dafny.SetOf(true, !(true))), true)).Dtor_cf3())), _dafny.SeqOf((Companion_D0_.Create_DC2_(_dafny.SetOf((_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ajxt")).Cardinality())), _dafny.IntOfInt64(-490)), _dafny.IntOfInt64(836), true)).Dtor_cf6()))
}
func (_static *CompanionStruct_Default___) Fm11(p0 _dafny.Sequence, globalState *GlobalState) _dafny.MultiSet {
	return ((_dafny.MultiSetOf(_dafny.CodePoint('o'), _dafny.CodePoint('m'), _dafny.CodePoint('f'))).Difference(_dafny.MultiSetOf(_dafny.CodePoint('f')))).Intersection(((Companion_D7_.Create_DC18_(_dafny.MultiSetOf(_dafny.CodePoint('c')))).Dtor_cf26()).Intersection(_dafny.MultiSetOf(_dafny.CodePoint('j'), _dafny.CodePoint('y'), _dafny.CodePoint('j'))))
}
func (_static *CompanionStruct_Default___) Fm12(p0 bool, p1 _dafny.Int, globalState *GlobalState) D4 {
	if !(false) {
		return Companion_D4_.Create_DC13_(Companion_D4_.Create_DC12_(_dafny.IntOfInt64(60), _dafny.CodePoint('e'), _dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(115))).Uint32(), func(coer3 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg3 _dafny.Int) interface{} {
				return coer3(arg3)
			}
		}(func(_12_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('t')
		}))).Cardinality()))))
	} else {
		return Companion_D4_.Create_DC13_(Companion_D4_.Create_DC12_((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-739), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(380))).Cardinality()), _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("bntttie")).Cardinality()))).Cardinality()))).Cardinality(), _dafny.CodePoint('g'), _dafny.MultiSetOf((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gix")).Cardinality()))).Cardinality())))
	}
}
func (_static *CompanionStruct_Default___) Fm13(p0 D0, p1 bool, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("kem")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("cqlomrr"), _dafny.UnicodeSeqOfUtf8Bytes("hxb"))), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("dugneonhq")), _dafny.SeqOf(_dafny.UnicodeSeqOfUtf8Bytes("awqmcmgf"))))
}
func (_static *CompanionStruct_Default___) Fm14(p0 _dafny.Int, p1 _dafny.CodePoint, p2 _dafny.Int, p3 _dafny.Int, globalState *GlobalState) _dafny.Set {
	return _dafny.SetOf(_dafny.CodePoint('d'), _dafny.CodePoint('y'), _dafny.CodePoint('y'), (func() _dafny.CodePoint {
		if false {
			return _dafny.CodePoint('d')
		}
		return _dafny.CodePoint('g')
	})(), _dafny.CodePoint('j'))
}
func (_static *CompanionStruct_Default___) Fm15(p0 _dafny.CodePoint, p1 _dafny.Map, globalState *GlobalState) _dafny.CodePoint {
	return _dafny.CodePoint('b')
}
func (_static *CompanionStruct_Default___) Fm16(p0 bool, p1 bool, globalState *GlobalState) _dafny.MultiSet {
	return (_dafny.MultiSetOf((_dafny.SetOf(!(false), false)).Cardinality())).Difference((_dafny.MultiSetOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.CodePoint('g'))).Cardinality()))).Difference(_dafny.MultiSetFromSeq(_dafny.SeqOf((_dafny.SetOf(_dafny.IntOfInt64(951))).Cardinality(), _dafny.IntOfInt64(849)))))
}
func (_static *CompanionStruct_Default___) Fm17(p0 bool, p1 _dafny.MultiSet, p2 _dafny.Int, p3 _dafny.Sequence, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate((Companion_D8_.Create_DC20_(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(true, false)).Cardinality()), (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("aapfmbeu")).Cardinality()))), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lytyvmod")).Cardinality()))))).Dtor_cf28(), _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("rdnvcfxx")).Cardinality()))), _dafny.SeqOf(_dafny.SeqOf((func() _dafny.Map {
		var _coll1 = _dafny.NewMapBuilder()
		_ = _coll1
		for _iter1 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(897), false)).Keys().Elements()); ; {
			_compr_1, _ok1 := _iter1()
			if !_ok1 {
				break
			}
			var _13_v0 _dafny.Int
			_13_v0 = interface{}(_compr_1).(_dafny.Int)
			if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(897), false)).Contains(_13_v0) {
				_coll1.Add((_13_v0).Plus(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(216))).Uint32(), func(coer4 func(_dafny.Int) _dafny.Int) func(_dafny.Int) interface{} {
					return func(arg4 _dafny.Int) interface{} {
						return coer4(arg4)
					}
				}(func(_14_i0 _dafny.Int) _dafny.Int {
					return _dafny.IntOfInt64(303)
				}))).Cardinality())), _dafny.CodePoint('t'))
			}
		}
		return _coll1.ToMap()
	}()).Cardinality()), _dafny.SeqOf(_dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("gwpfdjks")).Cardinality()))).Cardinality())), _dafny.SeqOf(_dafny.IntOfInt64(764)), (Companion_D9_.Create_DC24_(_dafny.SeqOf(_dafny.IntOfInt64(369), _dafny.IntOfInt64(-732)))).Dtor_cf36(), _dafny.SeqOf(_dafny.IntOfInt64(961), (_dafny.SetOf(false)).Cardinality(), _dafny.IntOfInt64(812), _dafny.IntOfInt64(50)))))
}
func (_static *CompanionStruct_Default___) Fm18(p0 _dafny.CodePoint, globalState *GlobalState) _dafny.Sequence {
	return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("dh"), _dafny.UnicodeSeqOfUtf8Bytes("nrki"))
}
func (_static *CompanionStruct_Default___) M0(p0 _dafny.Int, p1 _dafny.Sequence, globalState *GlobalState) (bool, _dafny.Array, _dafny.Int, _dafny.Int) {
	var r0 bool = false
	_ = r0
	var r1 _dafny.Array = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_ = r1
	var r2 _dafny.Int = _dafny.Zero
	_ = r2
	var r3 _dafny.Int = _dafny.Zero
	_ = r3
	r3 = p0
	r2 = p0
	var _15_v0 bool
	_ = _15_v0
	_15_v0 = true
	var _16_i0 _dafny.Int
	_ = _16_i0
	_16_i0 = _dafny.Zero
	{
		for _15_v0 {
			{
				if (_16_i0).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L0
				}
				_16_i0 = (_16_i0).Plus(_dafny.One)
				if _15_v0 {
					r3 = (p0).Minus((p0).Times(p0))
					var _17_v1 _dafny.Sequence
					_ = _17_v1
					_17_v1 = _dafny.UnicodeSeqOfUtf8Bytes("ym")
					var _18_v2 _dafny.Set
					_ = _18_v2
					_18_v2 = _dafny.SetOf(_dafny.IntOfUint32((_17_v1).Cardinality()))
					var _19_v3 _dafny.Map
					_ = _19_v3
					_19_v3 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_15_v0) == (_15_v0), _18_v2)
					_19_v3 = (_19_v3).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_15_v0), func() _dafny.Set {
						var _coll2 = _dafny.NewBuilder()
						_ = _coll2
						for _iter2 := _dafny.Iterate(_dafny.IntegerRange(_dafny.IntOfInt64(-838), _dafny.IntOfInt64(345))); ; {
							_compr_2, _ok2 := _iter2()
							if !_ok2 {
								break
							}
							var _20_v4 _dafny.Int
							_20_v4 = interface{}(_compr_2).(_dafny.Int)
							if ((_dafny.IntOfInt64(-838)).Cmp(_20_v4) <= 0) && ((_20_v4).Cmp(_dafny.IntOfInt64(345)) < 0) {
								_coll2.Add((_20_v4).Plus(p0))
							}
						}
						return _coll2.ToSet()
					}()))
					var _21_v5 _dafny.Array
					_ = _21_v5
					var _nw0 _dafny.Array = _dafny.NewArrayWithValue(Companion_D4_.Default(), _dafny.IntOfInt64(5))
					_ = _nw0
					_21_v5 = _nw0
					var _22_v6 _dafny.CodePoint
					_ = _22_v6
					_22_v6 = _dafny.CodePoint('t')
					var _23_v7 _dafny.Map
					_ = _23_v7
					_23_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(10), _dafny.IntOfInt64(-844))
					var _24_v8 _dafny.MultiSet
					_ = _24_v8
					_24_v8 = _dafny.MultiSetOf(p0)
					var _25_v9 D4
					_ = _25_v9
					_25_v9 = Companion_D4_.Create_DC12_(p0, Companion_Default___.Fm15(_22_v6, _23_v7, globalState), _24_v8)
					var _index0 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_21_v5), 0))
					_ = _index0
					(_21_v5).ArraySet1(_25_v9, (_index0).Int())
					var _index1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(163), _dafny.ArrayLen((_21_v5), 0))
					_ = _index1
					(_21_v5).ArraySet1(_25_v9, (_index1).Int())
					var _26_v10 _dafny.Map
					_ = _26_v10
					_26_v10 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(768), _17_v1)
					var _27_v11 D0
					_ = _27_v11
					_27_v11 = Companion_D0_.Create_DC3_(_22_v6)
					_26_v10 = (_26_v10).Update(p0, Companion_Default___.Fm18((_27_v11).Dtor_cf7(), globalState))
					var _28_v12 *C1
					_ = _28_v12
					var _nw1 *C1 = New_C1_()
					_ = _nw1
					_nw1.Ctor__((func() bool {
						if _15_v0 {
							return _15_v0
						}
						return _15_v0
					})(), _dafny.IntOfInt64(-333))
					_28_v12 = _nw1
					var _nw2 *C1 = New_C1_()
					_ = _nw2
					_nw2.Ctor__((p0).Cmp(_dafny.IntOfInt64(667)) <= 0, p0)
					_28_v12 = _nw2
				} else {
					var _29_v13 _dafny.CodePoint
					_ = _29_v13
					_29_v13 = _dafny.CodePoint('u')
					(globalState).F4 = _29_v13
					var _30_v14 _dafny.Array
					_ = _30_v14
					var _nw3 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(7))
					_ = _nw3
					_30_v14 = _nw3
					var _31_v15 _dafny.Map
					_ = _31_v15
					_31_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, _30_v14)
					_31_v15 = (_31_v15).Update((_15_v0) || (_15_v0), _30_v14)
					r2 = p0
					var _32_v16 _dafny.Array
					_ = _32_v16
					var _nw4 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
					_ = _nw4
					_32_v16 = _nw4
					_32_v16 = _32_v16
					var _33_v17 _dafny.Sequence
					_ = _33_v17
					_33_v17 = _dafny.SeqOf(p0)
					r0 = _dafny.Companion_Sequence_.Contains(_33_v17, p0)
				}
				var _34_v18 _dafny.Array
				_ = _34_v18
				var _nw5 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(8))
				_ = _nw5
				_34_v18 = _nw5
				var _35_v19 D0
				_ = _35_v19
				_35_v19 = Companion_D0_.Create_DC0_(p0, _15_v0)
				var _36_v20 _dafny.MultiSet
				_ = _36_v20
				_36_v20 = _dafny.MultiSetOf((_35_v19).Dtor_cf0(), p0)
				var _37_v21 _dafny.Map
				_ = _37_v21
				_37_v21 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_36_v20).Cardinality(), _34_v18)
				var _38_v22 _dafny.Sequence
				_ = _38_v22
				_38_v22 = _dafny.SeqOf(_34_v18)
				var _39_v23 _dafny.Map
				_ = _39_v23
				_39_v23 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, _34_v18)
				var _40_v24 _dafny.Array
				_ = _40_v24
				var _nwElement0_0 _dafny.Array = _34_v18
				_ = _nwElement0_0
				var _nw6 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_0, nil, _dafny.IntOfInt64(24))
				_ = _nw6
				(_nw6).ArraySet1(_nwElement0_0, 0)
				(_nw6).ArraySet1(_34_v18, 1)
				(_nw6).ArraySet1((func() _dafny.Array {
					if (_37_v21).Contains(p0) {
						return (_37_v21).Get(p0).(_dafny.Array)
					}
					return _34_v18
				})(), 2)
				(_nw6).ArraySet1(_34_v18, 3)
				(_nw6).ArraySet1(_34_v18, 4)
				(_nw6).ArraySet1(_34_v18, 5)
				(_nw6).ArraySet1((_38_v22).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_38_v22).Cardinality()))).Uint32()).(_dafny.Array), 6)
				(_nw6).ArraySet1(_34_v18, 7)
				(_nw6).ArraySet1(_34_v18, 8)
				(_nw6).ArraySet1(_34_v18, 9)
				(_nw6).ArraySet1(_34_v18, 10)
				(_nw6).ArraySet1(_34_v18, 11)
				(_nw6).ArraySet1(_34_v18, 12)
				(_nw6).ArraySet1((func() _dafny.Array {
					if (_39_v23).Contains(_15_v0) {
						return (_39_v23).Get(_15_v0).(_dafny.Array)
					}
					return _34_v18
				})(), 13)
				(_nw6).ArraySet1(_34_v18, 14)
				(_nw6).ArraySet1((func() _dafny.Array {
					if (_37_v21).Contains(p0) {
						return (_37_v21).Get(p0).(_dafny.Array)
					}
					return _34_v18
				})(), 15)
				(_nw6).ArraySet1(_34_v18, 16)
				(_nw6).ArraySet1(_34_v18, 17)
				(_nw6).ArraySet1(_34_v18, 18)
				(_nw6).ArraySet1(_34_v18, 19)
				(_nw6).ArraySet1(_34_v18, 20)
				(_nw6).ArraySet1(_34_v18, 21)
				(_nw6).ArraySet1(_34_v18, 22)
				(_nw6).ArraySet1(_34_v18, 23)
				_40_v24 = _nw6
				var _41_v25 D5
				_ = _41_v25
				_41_v25 = Companion_D5_.Create_DC14_(_40_v24)
				_41_v25 = _41_v25
				if false {
					var _42_v26 _dafny.Sequence
					_ = _42_v26
					_42_v26 = _dafny.UnicodeSeqOfUtf8Bytes("esfxlsqkn")
					r3 = Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_42_v26).Cardinality()), p0)
					var _43_v27 _dafny.Array
					_ = _43_v27
					var _len0_0 _dafny.Int = _dafny.IntOfInt64(12)
					_ = _len0_0
					var _nw7 _dafny.Array
					_ = _nw7
					if _len0_0.Cmp(_dafny.Zero) == 0 {
						_nw7 = _dafny.NewArray(_len0_0)
					} else {
						var _init0 func(_dafny.Int) bool = (func(_44_v0 bool) func(_dafny.Int) bool {
							return func(_45_i1 _dafny.Int) bool {
								return _44_v0
							}
						})(_15_v0)
						_ = _init0
						var _element0_0 = _init0(_dafny.Zero)
						_ = _element0_0
						_nw7 = _dafny.NewArrayFromExample(_element0_0, nil, _len0_0)
						(_nw7).ArraySet1(_element0_0, 0)
						var _nativeLen0_0 = (_len0_0).Int()
						_ = _nativeLen0_0
						for _i0_0 := 1; _i0_0 < _nativeLen0_0; _i0_0++ {
							(_nw7).ArraySet1(_init0(_dafny.IntOf(_i0_0)), _i0_0)
						}
					}
					_43_v27 = _nw7
					var _index2 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_43_v27), 0))
					_ = _index2
					(_43_v27).ArraySet1((_15_v0) && (Companion_Default___.Fm6((p1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), p0, globalState)), (_index2).Int())
					var _index3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_43_v27), 0))
					_ = _index3
					(_43_v27).ArraySet1((false) || (false), (_index3).Int())
					r3 = p0
					var _46_v28 _dafny.CodePoint
					_ = _46_v28
					_46_v28 = _dafny.CodePoint('b')
					var _47_v29 _dafny.Map
					_ = _47_v29
					_47_v29 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _46_v28)
					var _48_v30 *C0
					_ = _48_v30
					var _nw8 *C0 = New_C0_()
					_ = _nw8
					_nw8.Ctor__((_47_v29).Equals(_47_v29))
					_48_v30 = _nw8
					var _49_v31 D0
					_ = _49_v31
					_49_v31 = Companion_D0_.Create_DC3_(_46_v28)
					var _50_v32 _dafny.Array
					_ = _50_v32
					var _nwElement0_1 D0 = Companion_D0_.Create_DC3_(_46_v28)
					_ = _nwElement0_1
					var _nw9 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_1, nil, _dafny.IntOfInt64(16))
					_ = _nw9
					(_nw9).ArraySet1(_nwElement0_1, 0)
					(_nw9).ArraySet1(Companion_D0_.Create_DC3_(_46_v28), 1)
					(_nw9).ArraySet1(_49_v31, 2)
					(_nw9).ArraySet1(Companion_D0_.Create_DC3_(_46_v28), 3)
					(_nw9).ArraySet1(_49_v31, 4)
					(_nw9).ArraySet1(Companion_D0_.Create_DC3_(_46_v28), 5)
					(_nw9).ArraySet1(Companion_D0_.Create_DC3_(_dafny.CodePoint('c')), 6)
					(_nw9).ArraySet1(_49_v31, 7)
					(_nw9).ArraySet1(_49_v31, 8)
					(_nw9).ArraySet1(_49_v31, 9)
					(_nw9).ArraySet1(_49_v31, 10)
					(_nw9).ArraySet1(_49_v31, 11)
					(_nw9).ArraySet1(_49_v31, 12)
					(_nw9).ArraySet1(Companion_D0_.Create_DC3_(_46_v28), 13)
					(_nw9).ArraySet1(_49_v31, 14)
					(_nw9).ArraySet1(_49_v31, 15)
					_50_v32 = _nw9
					var _51_v33 _dafny.Array
					_ = _51_v33
					var _nwElement0_2 _dafny.Array = _50_v32
					_ = _nwElement0_2
					var _nw10 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_2, nil, _dafny.IntOfInt64(9))
					_ = _nw10
					(_nw10).ArraySet1(_nwElement0_2, 0)
					(_nw10).ArraySet1(_50_v32, 1)
					(_nw10).ArraySet1(_50_v32, 2)
					(_nw10).ArraySet1(_50_v32, 3)
					(_nw10).ArraySet1(_50_v32, 4)
					(_nw10).ArraySet1(_50_v32, 5)
					(_nw10).ArraySet1(_50_v32, 6)
					(_nw10).ArraySet1(_50_v32, 7)
					(_nw10).ArraySet1(_50_v32, 8)
					_51_v33 = _nw10
					var _index4 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_51_v33), 0))
					_ = _index4
					(_51_v33).ArraySet1(_50_v32, (_index4).Int())
					var _index5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_51_v33), 0))
					_ = _index5
					var _rhs0 *C0 = _48_v30
					_ = _rhs0
					var _rhs1 _dafny.Array = _50_v32
					_ = _rhs1
					var _lhs0 _dafny.Array = _51_v33
					_ = _lhs0
					var _lhs1 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(929), _dafny.ArrayLen((_51_v33), 0))
					_ = _lhs1
					_48_v30 = _rhs0
					(_lhs0).ArraySet1(_rhs1, (_lhs1).Int())
					var _index6 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_43_v27), 0))
					_ = _index6
					(_43_v27).ArraySet1(!((_43_v27).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(92), _dafny.ArrayLen((_43_v27), 0))).Int()).(bool)), (_index6).Int())
				} else {
					r0 = _15_v0
					var _index7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))
					_ = _index7
					(_34_v18).ArraySet1(p0, (_index7).Int())
					var _index8 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))
					_ = _index8
					(_34_v18).ArraySet1(p0, (_index8).Int())
					var _52_v34 _dafny.Sequence
					_ = _52_v34
					_52_v34 = _dafny.SeqOf(_dafny.IntOfInt64(479), Companion_Default___.Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(true, _15_v0), _15_v0, _15_v0, globalState))
					var _53_v35 _dafny.Set
					_ = _53_v35
					_53_v35 = _dafny.SetOf(_15_v0, _15_v0)
					var _54_v36 _dafny.Map
					_ = _54_v36
					_54_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, Companion_Default___.Fm4(_15_v0, _53_v35, p0, globalState))
					var _55_v37 _dafny.Sequence
					_ = _55_v37
					_55_v37 = _dafny.SeqOf((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int), (_dafny.IntOfUint32((_52_v34).Cardinality())).Minus(Companion_Default___.Fm0(_54_v36, _15_v0, _15_v0, globalState)), Companion_Default___.SafeModuloInt((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int), p0))
					_55_v37 = _dafny.Companion_Sequence_.Concatenate(_55_v37, _52_v34)
					var _56_v38 _dafny.CodePoint
					_ = _56_v38
					_56_v38 = _dafny.CodePoint('g')
					var _57_v39 _dafny.Sequence
					_ = _57_v39
					_57_v39 = _dafny.UnicodeSeqOfUtf8Bytes("sir")
					var _58_v40 _dafny.Set
					_ = _58_v40
					_58_v40 = _dafny.SetOf(_36_v20, _36_v20)
					var _59_v41 _dafny.Array
					_ = _59_v41
					var _nwElement0_3 bool = _15_v0
					_ = _nwElement0_3
					var _nw11 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_3, nil, _dafny.IntOfInt64(17))
					_ = _nw11
					(_nw11).ArraySet1(_nwElement0_3, 0)
					(_nw11).ArraySet1(_15_v0, 1)
					(_nw11).ArraySet1(_15_v0, 2)
					(_nw11).ArraySet1(!(!(!_dafny.Companion_Sequence_.Contains(_57_v39, _56_v38))), 3)
					(_nw11).ArraySet1(_15_v0, 4)
					(_nw11).ArraySet1(_15_v0, 5)
					(_nw11).ArraySet1((Companion_Default___.Fm0(_54_v36, _15_v0, _15_v0, globalState)).Cmp(p0) > 0, 6)
					(_nw11).ArraySet1(_15_v0, 7)
					(_nw11).ArraySet1((p1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), 8)
					(_nw11).ArraySet1(((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int)).Cmp((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int)) >= 0, 9)
					(_nw11).ArraySet1(_15_v0, 10)
					(_nw11).ArraySet1(_15_v0, 11)
					(_nw11).ArraySet1((p0).Cmp(_dafny.IntOfInt64(77)) <= 0, 12)
					(_nw11).ArraySet1((_58_v40).Contains(_36_v20), 13)
					(_nw11).ArraySet1(_15_v0, 14)
					(_nw11).ArraySet1((_15_v0) || (!((p1).Select((Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool))), 15)
					(_nw11).ArraySet1(!(_15_v0), 16)
					_59_v41 = _nw11
					var _index9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))
					_ = _index9
					var _rhs2 _dafny.Int = p0
					_ = _rhs2
					var _rhs3 _dafny.Int = Companion_Default___.Fm0(_54_v36, !(!(_15_v0)), _15_v0, globalState)
					_ = _rhs3
					var _rhs4 _dafny.Array = _59_v41
					_ = _rhs4
					var _rhs5 bool = (func() bool {
						if (func() bool {
							if _15_v0 {
								return _15_v0
							}
							return (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(696), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool)
						})() {
							return Companion_Default___.Fm6(_15_v0, (_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int), globalState)
						}
						return ((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int)).Cmp((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int)) > 0
					})()
					_ = _rhs5
					var _lhs2 _dafny.Array = _34_v18
					_ = _lhs2
					var _lhs3 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))
					_ = _lhs3
					(_lhs2).ArraySet1(_rhs2, (_lhs3).Int())
					r3 = _rhs3
					_59_v41 = _rhs4
					_15_v0 = _rhs5
					var _index10 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))
					_ = _index10
					(_34_v18).ArraySet1((_34_v18).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(438), _dafny.ArrayLen((_34_v18), 0))).Int()).(_dafny.Int), (_index10).Int())
				}
				var _60_v42 _dafny.CodePoint
				_ = _60_v42
				_60_v42 = _dafny.CodePoint('x')
				var _61_v43 _dafny.Map
				_ = _61_v43
				_61_v43 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, p0)
				var _62_v44 _dafny.Sequence
				_ = _62_v44
				_62_v44 = _dafny.SeqOf(Companion_Default___.Fm15(_60_v42, _61_v43, globalState))
				var _63_v45 _dafny.Map
				_ = _63_v45
				_63_v45 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, _62_v44)
				var _64_v46 _dafny.Set
				_ = _64_v46
				_64_v46 = _dafny.SetOf(!(_63_v45).Contains(false), (p0).Cmp(p0) < 0, _15_v0)
				_64_v46 = ((_64_v46).Union(_64_v46)).Union(_64_v46)
				goto C0
			}
		C0:
		}
		goto L0
	}
L0:
	var _65_v47 _dafny.Array
	_ = _65_v47
	var _nw12 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(3))
	_ = _nw12
	_65_v47 = _nw12
	_65_v47 = _65_v47
	var _66_v48 _dafny.Sequence
	_ = _66_v48
	_66_v48 = _dafny.UnicodeSeqOfUtf8Bytes("mns")
	var _67_v49 D1
	_ = _67_v49
	_67_v49 = Companion_D1_.Create_DC4_(_66_v48)
	var _source1 D1 = _67_v49
	_ = _source1
	if _source1.Is_DC5() {
		var _68___mcc_h0 bool = _source1.Get_().(D1_DC5).Cf9
		_ = _68___mcc_h0
		var _69_cf9 bool = _68___mcc_h0
		_ = _69_cf9
		var _70_v50 _dafny.Set
		_ = _70_v50
		_70_v50 = _dafny.SetOf(_69_cf9)
		var _71_v51 _dafny.Map
		_ = _71_v51
		_71_v51 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _69_cf9)
		var _72_v52 _dafny.Set
		_ = _72_v52
		_72_v52 = _dafny.SetOf(Companion_Default___.Fm4(_15_v0, _70_v50, (_dafny.MultiSetOf((_71_v51).Cardinality(), p0, p0, p0, p0)).Cardinality(), globalState))
		if (func() bool {
			if _15_v0 {
				return (_72_v52).IsSubsetOf(_70_v50)
			}
			return _15_v0
		})() {
			var _73_v53 _dafny.MultiSet
			_ = _73_v53
			_73_v53 = _dafny.MultiSetOf(_dafny.SeqOf(_15_v0, _69_cf9), p1, _dafny.SeqOf(_15_v0, _69_cf9), _dafny.SeqOf(_69_cf9, _69_cf9, _15_v0, _69_cf9, !(_69_cf9)))
			var _74_v54 _dafny.MultiSet
			_ = _74_v54
			_74_v54 = _dafny.MultiSetOf(_73_v53, _dafny.MultiSetOf(p1))
			var _75_v55 _dafny.MultiSet
			_ = _75_v55
			_75_v55 = _dafny.MultiSetOf(true)
			var _76_v56 _dafny.MultiSet
			_ = _76_v56
			_76_v56 = _dafny.MultiSetOf((_dafny.Zero).Minus(p0), (func() _dafny.Int {
				if (_74_v54).Contains(_73_v53) {
					return (_74_v54).Multiplicity(_73_v53)
				}
				return (func() _dafny.Int {
					if (_75_v55).Contains(_15_v0) {
						return (_75_v55).Multiplicity(_15_v0)
					}
					return p0
				})()
			})(), p0)
			_76_v56 = _76_v56
			var _77_v57 _dafny.CodePoint
			_ = _77_v57
			_77_v57 = _dafny.CodePoint('c')
			var _78_v58 _dafny.Map
			_ = _78_v58
			_78_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Update(_66_v48, (Companion_Default___.SafeIndex(p0, _dafny.IntOfUint32((_66_v48).Cardinality()))).Uint32(), _77_v57), p0)
			var _79_v59 _dafny.Map
			_ = _79_v59
			_79_v59 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.Companion_Sequence_.Concatenate(Companion_Default___.Fm18(_77_v57, globalState), _66_v48), (_78_v58).Cardinality())
			var _80_v60 _dafny.Set
			_ = _80_v60
			_80_v60 = _dafny.SetOf(p0)
			var _81_v61 _dafny.Set
			_ = _81_v61
			_81_v61 = _dafny.SetOf(p0, (_dafny.MultiSetOf(_80_v60, _dafny.SetOf(p0))).Cardinality())
			_79_v59 = (_79_v59).Update((func() _dafny.Sequence {
				if _69_cf9 {
					return _66_v48
				}
				return _dafny.UnicodeSeqOfUtf8Bytes("bmdwy")
			})(), (_80_v60).Cardinality())
			var _82_v62 _dafny.Array
			_ = _82_v62
			var _nw13 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(24))
			_ = _nw13
			_82_v62 = _nw13
			var _index11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_82_v62), 0))
			_ = _index11
			(_82_v62).ArraySet1(p0, (_index11).Int())
			var _index12 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_82_v62), 0))
			_ = _index12
			(_82_v62).ArraySet1((p0).Minus(Companion_Default___.SafeDivisionInt(_dafny.IntOfUint32((_66_v48).Cardinality()), (_dafny.Zero).Minus((_75_v55).Cardinality()))), (_index12).Int())
			r2 = p0
			var _83_v63 *C0
			_ = _83_v63
			var _nw14 *C0 = New_C0_()
			_ = _nw14
			_nw14.Ctor__(_15_v0)
			_83_v63 = _nw14
			var _84_v64 _dafny.Map
			_ = _84_v64
			_84_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_dafny.Zero).Minus(p0), _83_v63)
			_84_v64 = (_84_v64).Update((_82_v62).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(948), _dafny.ArrayLen((_82_v62), 0))).Int()).(_dafny.Int), _83_v63)
		} else {
			var _85_v65 _dafny.Array
			_ = _85_v65
			var _len0_1 _dafny.Int = _dafny.IntOfInt64(5)
			_ = _len0_1
			var _nw15 _dafny.Array
			_ = _nw15
			if _len0_1.Cmp(_dafny.Zero) == 0 {
				_nw15 = _dafny.NewArray(_len0_1)
			} else {
				var _init1 func(_dafny.Int) _dafny.Int = func(_86_i2 _dafny.Int) _dafny.Int {
					return (_86_i2).Plus(_dafny.IntOfInt64(310))
				}
				_ = _init1
				var _element0_1 = _init1(_dafny.Zero)
				_ = _element0_1
				_nw15 = _dafny.NewArrayFromExample(_element0_1, nil, _len0_1)
				(_nw15).ArraySet1(_element0_1, 0)
				var _nativeLen0_1 = (_len0_1).Int()
				_ = _nativeLen0_1
				for _i0_1 := 1; _i0_1 < _nativeLen0_1; _i0_1++ {
					(_nw15).ArraySet1(_init1(_dafny.IntOf(_i0_1)), _i0_1)
				}
			}
			_85_v65 = _nw15
			var _index13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_85_v65), 0))
			_ = _index13
			(_85_v65).ArraySet1(p0, (_index13).Int())
			var _87_v66 _dafny.Map
			_ = _87_v66
			_87_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, _69_cf9)
			var _88_v67 _dafny.Map
			_ = _88_v67
			_88_v67 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_66_v48, _87_v66)
			var _index14 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_85_v65), 0))
			_ = _index14
			var _rhs6 bool = Companion_Default___.Fm6(_15_v0, (_dafny.Zero).Minus(Companion_Default___.Fm0((func() _dafny.Map {
				if (_88_v67).Contains(_66_v48) {
					return (_88_v67).Get(_66_v48).(_dafny.Map)
				}
				return _87_v66
			})(), (p1).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_66_v48).Cardinality()), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), _15_v0, globalState)), globalState)
			_ = _rhs6
			var _rhs7 _dafny.Int = (_dafny.Zero).Minus(p0)
			_ = _rhs7
			var _rhs8 bool = !(_69_cf9)
			_ = _rhs8
			var _rhs9 _dafny.Int = ((p0).Plus(Companion_Default___.Fm0(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_15_v0), !(true)), (func() bool {
				if (_87_v66).Contains(_15_v0) {
					return (_87_v66).Get(_15_v0).(bool)
				}
				return _69_cf9
			})(), _15_v0, globalState))).Minus(p0)
			_ = _rhs9
			var _lhs4 _dafny.Array = _85_v65
			_ = _lhs4
			var _lhs5 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(910), _dafny.ArrayLen((_85_v65), 0))
			_ = _lhs5
			r0 = _rhs6
			(_lhs4).ArraySet1(_rhs7, (_lhs5).Int())
			r0 = _rhs8
			r2 = _rhs9
			var _89_v68 T0
			_ = _89_v68
			var _nw16 *C0 = New_C0_()
			_ = _nw16
			_nw16.Ctor__(_15_v0)
			_89_v68 = _nw16
			var _rhs10 _dafny.Int = Companion_Default___.SafeModuloInt(p0, _dafny.IntOfInt64(-220))
			_ = _rhs10
			var _rhs11 bool = _15_v0
			_ = _rhs11
			var _rhs12 T0 = _89_v68
			_ = _rhs12
			r2 = _rhs10
			_15_v0 = _rhs11
			_89_v68 = _rhs12
			var _90_v69 _dafny.Array
			_ = _90_v69
			var _len0_2 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_2
			var _nw17 _dafny.Array
			_ = _nw17
			if _len0_2.Cmp(_dafny.Zero) == 0 {
				_nw17 = _dafny.NewArray(_len0_2)
			} else {
				var _init2 func(_dafny.Int) bool = (func(_91_v0 bool) func(_dafny.Int) bool {
					return func(_92_i3 _dafny.Int) bool {
						return _91_v0
					}
				})(_15_v0)
				_ = _init2
				var _element0_2 = _init2(_dafny.Zero)
				_ = _element0_2
				_nw17 = _dafny.NewArrayFromExample(_element0_2, nil, _len0_2)
				(_nw17).ArraySet1(_element0_2, 0)
				var _nativeLen0_2 = (_len0_2).Int()
				_ = _nativeLen0_2
				for _i0_2 := 1; _i0_2 < _nativeLen0_2; _i0_2++ {
					(_nw17).ArraySet1(_init2(_dafny.IntOf(_i0_2)), _i0_2)
				}
			}
			_90_v69 = _nw17
			var _index15 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_90_v69), 0))
			_ = _index15
			(_90_v69).ArraySet1(_69_cf9, (_index15).Int())
			var _index16 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_90_v69), 0))
			_ = _index16
			var _rhs13 bool = _69_cf9
			_ = _rhs13
			var _rhs14 bool = _15_v0
			_ = _rhs14
			var _lhs6 _dafny.Array = _90_v69
			_ = _lhs6
			var _lhs7 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_90_v69), 0))
			_ = _lhs7
			_69_cf9 = _rhs13
			(_lhs6).ArraySet1(_rhs14, (_lhs7).Int())
			r3 = (p0).Minus(Companion_Default___.SafeDivisionInt(p0, p0))
			var _index17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(515), _dafny.ArrayLen((_90_v69), 0))
			_ = _index17
			(_90_v69).ArraySet1(!(_69_cf9), (_index17).Int())
		}
		var _93_v70 _dafny.Array
		_ = _93_v70
		var _nw18 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
		_ = _nw18
		_93_v70 = _nw18
		var _index18 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_93_v70), 0))
		_ = _index18
		(_93_v70).ArraySet1(_69_cf9, (_index18).Int())
		var _index19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_93_v70), 0))
		_ = _index19
		(_93_v70).ArraySet1((p1).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeModuloInt(_dafny.IntOfInt64(609), _dafny.IntOfUint32((_66_v48).Cardinality())), _dafny.IntOfUint32((p1).Cardinality()))).Uint32()).(bool), (_index19).Int())
		var _94_v71 *C1
		_ = _94_v71
		var _nw19 *C1 = New_C1_()
		_ = _nw19
		_nw19.Ctor__((_93_v70).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(116), _dafny.ArrayLen((_93_v70), 0))).Int()).(bool), p0)
		_94_v71 = _nw19
		var _95_v72 _dafny.Map
		_ = _95_v72
		_95_v72 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_94_v71, p0)
		var _96_v73 _dafny.Map
		_ = _96_v73
		_96_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_69_cf9, (_95_v72).Cardinality())
		var _97_v74 _dafny.Sequence
		_ = _97_v74
		_97_v74 = _dafny.SeqOf((_94_v71).F15())
		var _98_v75 *C1
		_ = _98_v75
		var _nw20 *C1 = New_C1_()
		_ = _nw20
		_nw20.Ctor__(Companion_Default___.Fm6(_15_v0, (_96_v73).Cardinality(), globalState), _dafny.IntOfUint32((_97_v74).Cardinality()))
		_98_v75 = _nw20
		var _99_v76 _dafny.Set
		_ = _99_v76
		_99_v76 = _dafny.SetOf(_93_v70)
		var _100_v77 _dafny.Map
		_ = _100_v77
		_100_v77 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_99_v76, _96_v73)
		_100_v77 = (_100_v77).Update(_99_v76, _96_v73)
	} else {
		var _101___mcc_h1 _dafny.Sequence = _source1.Get_().(D1_DC4).Cf8
		_ = _101___mcc_h1
		var _102_cf8 _dafny.Sequence = _101___mcc_h1
		_ = _102_cf8
		var _103_v78 _dafny.Sequence
		_ = _103_v78
		_103_v78 = _dafny.SeqOf(p0)
		var _104_v79 _dafny.Map
		_ = _104_v79
		_104_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(p0, _103_v78)
		var _105_v80 _dafny.Map
		_ = _105_v80
		_105_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, _15_v0)
		_103_v78 = _dafny.Companion_Sequence_.Concatenate(_dafny.Companion_Sequence_.Update((func() _dafny.Sequence {
			if (_104_v79).Contains(p0) {
				return (_104_v79).Get(p0).(_dafny.Sequence)
			}
			return _103_v78
		})(), (Companion_Default___.SafeIndex(Companion_Default___.Fm0(_105_v80, _15_v0, _15_v0, globalState), _dafny.IntOfUint32(((func() _dafny.Sequence {
			if (_104_v79).Contains(p0) {
				return (_104_v79).Get(p0).(_dafny.Sequence)
			}
			return _103_v78
		})()).Cardinality()))).Uint32(), p0), _dafny.Companion_Sequence_.Concatenate(_103_v78, _103_v78))
		var _106_v81 *C1
		_ = _106_v81
		var _nw21 *C1 = New_C1_()
		_ = _nw21
		_nw21.Ctor__((p0).Cmp(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("sm")).Cardinality())) > 0, p0)
		_106_v81 = _nw21
		var _107_v82 _dafny.Map
		_ = _107_v82
		_107_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_106_v81).F15(), _15_v0)
		var _108_v83 _dafny.Map
		_ = _108_v83
		_108_v83 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_106_v81).F15(), _107_v82)
		var _109_v84 _dafny.Map
		_ = _109_v84
		_109_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_106_v81).F15()).Minus(((func() _dafny.Map {
			if (_108_v83).Contains(p0) {
				return (_108_v83).Get(p0).(_dafny.Map)
			}
			return _107_v82
		})()).Cardinality()), (_106_v81).F14())
		_107_v82 = (_107_v82).Update(p0, _15_v0)
		var _110_v85 _dafny.CodePoint
		_ = _110_v85
		_110_v85 = _dafny.CodePoint('y')
		var _111_v86 T0
		_ = _111_v86
		var _nw22 *C0 = New_C0_()
		_ = _nw22
		_nw22.Ctor__((_106_v81).F14())
		_111_v86 = _nw22
		var _112_v87 _dafny.Map
		_ = _112_v87
		_112_v87 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_15_v0, _110_v85)
		var _113_v88 _dafny.Map
		_ = _113_v88
		_113_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_111_v86, _112_v87)
		var _114_v89 _dafny.MultiSet
		_ = _114_v89
		_114_v89 = _dafny.MultiSetOf((_dafny.Zero).Minus(((func() _dafny.Map {
			if (_113_v88).Contains(_111_v86) {
				return (_113_v88).Get(_111_v86).(_dafny.Map)
			}
			return _112_v87
		})()).Cardinality()))
		var _115_v90 D4
		_ = _115_v90
		_115_v90 = Companion_D4_.Create_DC12_((_106_v81).F15(), _110_v85, _114_v89)
		r3 = (((_115_v90).Dtor_cf20()).Cardinality()).Plus((_dafny.Zero).Minus((_dafny.Zero).Minus((_103_v78).Select((Companion_Default___.SafeIndex((_106_v81).F15(), _dafny.IntOfUint32((_103_v78).Cardinality()))).Uint32()).(_dafny.Int))))
	}
	r0 = false
	var _116_v91 _dafny.CodePoint
	_ = _116_v91
	_116_v91 = _dafny.CodePoint('l')
	var _117_v92 _dafny.MultiSet
	_ = _117_v92
	_117_v92 = _dafny.MultiSetOf(_dafny.IntOfUint32((p1).Cardinality()))
	var _118_v93 D4
	_ = _118_v93
	_118_v93 = Companion_D4_.Create_DC12_(p0, _116_v91, _117_v92)
	var _pat_let_tv0 = _15_v0
	_ = _pat_let_tv0
	var _pat_let_tv1 = p1
	_ = _pat_let_tv1
	var _pat_let_tv2 = p1
	_ = _pat_let_tv2
	var _pat_let_tv3 = _15_v0
	_ = _pat_let_tv3
	var _pat_let_tv4 = _15_v0
	_ = _pat_let_tv4
	r0 = !(func(_source2 D4) bool {
		if _source2.Is_DC12() {
			var _119___mcc_h2 _dafny.Int = _source2.Get_().(D4_DC12).Cf18
			_ = _119___mcc_h2
			var _120___mcc_h3 _dafny.CodePoint = _source2.Get_().(D4_DC12).Cf19
			_ = _120___mcc_h3
			var _121___mcc_h4 _dafny.MultiSet = _source2.Get_().(D4_DC12).Cf20
			_ = _121___mcc_h4
			var _122_cf20 _dafny.MultiSet = _121___mcc_h4
			_ = _122_cf20
			var _123_cf19 _dafny.CodePoint = _120___mcc_h3
			_ = _123_cf19
			var _124_cf18 _dafny.Int = _119___mcc_h2
			_ = _124_cf18
			return !(_pat_let_tv0) || (!((_pat_let_tv1).Select((Companion_Default___.SafeIndex(_124_cf18, _dafny.IntOfUint32((_pat_let_tv2).Cardinality()))).Uint32()).(bool)))
		} else if _source2.Is_DC11() {
			var _125___mcc_h5 _dafny.Set = _source2.Get_().(D4_DC11).Cf17
			_ = _125___mcc_h5
			var _126_cf17 _dafny.Set = _125___mcc_h5
			_ = _126_cf17
			return _pat_let_tv3
		} else {
			var _127___mcc_h6 D4 = _source2.Get_().(D4_DC13).Cf21
			_ = _127___mcc_h6
			var _128_cf21 D4 = _127___mcc_h6
			_ = _128_cf21
			return _pat_let_tv4
		}
	}(_118_v93))
	var _129_v94 _dafny.Sequence
	_ = _129_v94
	_129_v94 = _dafny.SeqOf(p0, p0, p0, p0, p0)
	var _130_v95 D0
	_ = _130_v95
	_130_v95 = Companion_D0_.Create_DC0_(p0, _15_v0)
	var _pat_let_tv5 = p0
	_ = _pat_let_tv5
	var _131_v96 _dafny.Array
	_ = _131_v96
	var _nwElement0_4 D0 = _130_v95
	_ = _nwElement0_4
	var _nw23 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_4, nil, _dafny.IntOfInt64(14))
	_ = _nw23
	(_nw23).ArraySet1(_nwElement0_4, 0)
	(_nw23).ArraySet1(Companion_D0_.Create_DC0_(p0, _15_v0), 1)
	(_nw23).ArraySet1(Companion_D0_.Create_DC0_(p0, Companion_Default___.Fm6(_15_v0, p0, globalState)), 2)
	(_nw23).ArraySet1(_130_v95, 3)
	(_nw23).ArraySet1(_130_v95, 4)
	(_nw23).ArraySet1(_130_v95, 5)
	(_nw23).ArraySet1(Companion_D0_.Create_DC0_(p0, Companion_Default___.Fm6(_15_v0, p0, globalState)), 6)
	(_nw23).ArraySet1(_130_v95, 7)
	(_nw23).ArraySet1(_130_v95, 8)
	(_nw23).ArraySet1(func(_pat_let0_0 D0) D0 {
		return func(_132_dt__update__tmp_h0 D0) D0 {
			return func(_pat_let1_0 _dafny.Int) D0 {
				return func(_133_dt__update_hcf0_h0 _dafny.Int) D0 {
					return Companion_D0_.Create_DC0_(_133_dt__update_hcf0_h0, (_132_dt__update__tmp_h0).Dtor_cf1())
				}(_pat_let1_0)
			}(_pat_let_tv5)
		}(_pat_let0_0)
	}(_130_v95), 9)
	(_nw23).ArraySet1(_130_v95, 10)
	(_nw23).ArraySet1(_130_v95, 11)
	(_nw23).ArraySet1(Companion_D0_.Create_DC0_(p0, _15_v0), 12)
	(_nw23).ArraySet1(_130_v95, 13)
	_131_v96 = _nw23
	r1 = (func() _dafny.Array {
		if (func() bool {
			if !(_15_v0) {
				return Companion_Default___.Fm4(_15_v0, _dafny.SetOf(_15_v0, _15_v0), _dafny.IntOfUint32((_129_v94).Cardinality()), globalState)
			}
			return _15_v0
		})() {
			return (func() _dafny.Array {
				if _15_v0 {
					return _131_v96
				}
				return _131_v96
			})()
		}
		return _131_v96
	})()
	r2 = (func() _dafny.Int {
		if _15_v0 {
			return (p0).Minus((_117_v92).Cardinality())
		}
		return p0
	})()
	r3 = p0
	return r0, r1, r2, r3
}
func (_static *CompanionStruct_Default___) Main(__noArgsParameter _dafny.Sequence) {
	var _134_v0 _dafny.Int
	_ = _134_v0
	_134_v0 = _dafny.IntOfInt64(-250)
	var _135_v1 _dafny.MultiSet
	_ = _135_v1
	_135_v1 = _dafny.MultiSetOf(_dafny.MultiSetOf(_134_v0))
	var _136_v2 _dafny.Array
	_ = _136_v2
	var _nw24 _dafny.Array = _dafny.NewArrayWithValue(_dafny.CodePoint('D'), _dafny.IntOfInt64(29))
	_ = _nw24
	_136_v2 = _nw24
	var _137_v3 _dafny.Set
	_ = _137_v3
	_137_v3 = _dafny.SetOf(_136_v2)
	var _138_v4 _dafny.Sequence
	_ = _138_v4
	_138_v4 = _dafny.SeqOf(_134_v0)
	var _139_v5 _dafny.Set
	_ = _139_v5
	_139_v5 = _dafny.SetOf(_138_v4)
	var _140_v6 bool
	_ = _140_v6
	_140_v6 = false
	var _141_v7 _dafny.MultiSet
	_ = _141_v7
	_141_v7 = _dafny.MultiSetOf(_140_v6, _140_v6)
	var _142_v8 _dafny.MultiSet
	_ = _142_v8
	_142_v8 = _dafny.MultiSetOf(_141_v7)
	var _143_v9 _dafny.Sequence
	_ = _143_v9
	_143_v9 = _dafny.SeqOf(!(_140_v6))
	var _144_v10 _dafny.Sequence
	_ = _144_v10
	_144_v10 = _dafny.UnicodeSeqOfUtf8Bytes("jxxycbwg")
	var _145_v11 _dafny.Map
	_ = _145_v11
	_145_v11 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v6, true)
	var _146_v12 _dafny.Map
	_ = _146_v12
	_146_v12 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("ue")).Cardinality()))
	var _147_v13 _dafny.Array
	_ = _147_v13
	var _nwElement0_5 _dafny.Int = _134_v0
	_ = _nwElement0_5
	var _nw25 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_5, nil, _dafny.IntOfInt64(8))
	_ = _nw25
	(_nw25).ArraySet1(_nwElement0_5, 0)
	(_nw25).ArraySet1(_dafny.IntOfUint32((_144_v10).Cardinality()), 1)
	(_nw25).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v6, _145_v11)).Cardinality(), 2)
	(_nw25).ArraySet1(_134_v0, 3)
	(_nw25).ArraySet1(_dafny.IntOfUint32((_138_v4).Cardinality()), 4)
	(_nw25).ArraySet1(_dafny.IntOfUint32((_144_v10).Cardinality()), 5)
	(_nw25).ArraySet1((_146_v12).Cardinality(), 6)
	(_nw25).ArraySet1(_134_v0, 7)
	_147_v13 = _nw25
	var _148_globalState *GlobalState
	_ = _148_globalState
	var _nw26 *GlobalState = New_GlobalState_()
	_ = _nw26
	_nw26.Ctor__(_dafny.IntOfInt64(870), _dafny.IntOfInt64(-430), _135_v1, false, _dafny.CodePoint('x'), _137_v3, (_139_v5).Union(_139_v5), _dafny.CodePoint('u'), _142_v8, false, _dafny.IntOfInt64(270), _143_v9, _147_v13, true)
	_148_globalState = _nw26
	var _index20 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
	_ = _index20
	(_147_v13).ArraySet1(_134_v0, (_index20).Int())
	var _index21 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
	_ = _index21
	(_147_v13).ArraySet1(Companion_Default___.SafeModuloInt(((_dafny.Zero).Minus(_dafny.IntOfUint32((_144_v10).Cardinality()))).Times(_dafny.IntOfInt64(427)), (_dafny.Zero).Minus(_134_v0)), (_index21).Int())
	var _149_v14 _dafny.Map
	_ = _149_v14
	_149_v14 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), _140_v6)
	var _150_v15 _dafny.Set
	_ = _150_v15
	_150_v15 = _dafny.SetOf((_149_v14).Cardinality())
	var _151_v16 D0
	_ = _151_v16
	_151_v16 = Companion_D0_.Create_DC2_(_150_v15, _134_v0, _140_v6)
	var _source3 D0 = (func() D0 {
		if (_151_v16).Dtor_cf6() {
			return _151_v16
		}
		return _151_v16
	})()
	_ = _source3
	if _source3.Is_DC0() {
		var _152___mcc_h0 _dafny.Int = _source3.Get_().(D0_DC0).Cf0
		_ = _152___mcc_h0
		var _153___mcc_h1 bool = _source3.Get_().(D0_DC0).Cf1
		_ = _153___mcc_h1
		var _154_cf1 bool = _153___mcc_h1
		_ = _154_cf1
		var _155_cf0 _dafny.Int = _152___mcc_h0
		_ = _155_cf0
		var _156_v17 _dafny.Array
		_ = _156_v17
		var _nw27 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(5))
		_ = _nw27
		_156_v17 = _nw27
		var _index22 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_156_v17), 0))
		_ = _index22
		(_156_v17).ArraySet1(false, (_index22).Int())
		var _157_v18 _dafny.MultiSet
		_ = _157_v18
		_157_v18 = _dafny.MultiSetOf(_dafny.Companion_Sequence_.Concatenate(_144_v10, _144_v10))
		var _index23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_156_v17), 0))
		_ = _index23
		(_156_v17).ArraySet1(_140_v6, (_index23).Int())
		var _index24 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_156_v17), 0))
		_ = _index24
		var _index25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index25
		var _index26 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_156_v17), 0))
		_ = _index26
		var _rhs15 bool = _154_cf1
		_ = _rhs15
		var _rhs16 _dafny.Int = Companion_Default___.SafeModuloInt(_155_cf0, _134_v0)
		_ = _rhs16
		var _rhs17 _dafny.Int = (_dafny.IntOfUint32((_144_v10).Cardinality())).Times((_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_134_v0, _155_cf0)))
		_ = _rhs17
		var _rhs18 _dafny.MultiSet = _157_v18
		_ = _rhs18
		var _rhs19 bool = (_134_v0).Cmp(Companion_Default___.SafeModuloInt(_134_v0, (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))) == 0
		_ = _rhs19
		var _lhs8 _dafny.Array = _156_v17
		_ = _lhs8
		var _lhs9 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_156_v17), 0))
		_ = _lhs9
		var _lhs10 _dafny.Array = _147_v13
		_ = _lhs10
		var _lhs11 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _lhs11
		var _lhs12 _dafny.Array = _156_v17
		_ = _lhs12
		var _lhs13 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(50), _dafny.ArrayLen((_156_v17), 0))
		_ = _lhs13
		(_lhs8).ArraySet1(_rhs15, (_lhs9).Int())
		(_lhs10).ArraySet1(_rhs16, (_lhs11).Int())
		_155_cf0 = _rhs17
		_157_v18 = _rhs18
		(_lhs12).ArraySet1(_rhs19, (_lhs13).Int())
		if !(_154_cf1) {
			var _158_v19 bool
			_ = _158_v19
			var _159_v20 _dafny.Array
			_ = _159_v20
			var _160_v21 _dafny.Int
			_ = _160_v21
			var _161_v22 _dafny.Int
			_ = _161_v22
			var _out0 bool
			_ = _out0
			var _out1 _dafny.Array
			_ = _out1
			var _out2 _dafny.Int
			_ = _out2
			var _out3 _dafny.Int
			_ = _out3
			_out0, _out1, _out2, _out3 = Companion_Default___.M0(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_144_v10).Cardinality()), _dafny.IntOfInt64(-6)), _143_v9, _148_globalState)
			_158_v19 = _out0
			_159_v20 = _out1
			_160_v21 = _out2
			_161_v22 = _out3
			var _162_v23 bool
			_ = _162_v23
			var _163_v24 _dafny.Array
			_ = _163_v24
			var _164_v25 _dafny.Int
			_ = _164_v25
			var _165_v26 _dafny.Int
			_ = _165_v26
			var _out4 bool
			_ = _out4
			var _out5 _dafny.Array
			_ = _out5
			var _out6 _dafny.Int
			_ = _out6
			var _out7 _dafny.Int
			_ = _out7
			_out4, _out5, _out6, _out7 = Companion_Default___.M0((Companion_Default___.Fm0(_145_v11, _154_cf1, false, _148_globalState)).Times(_dafny.IntOfInt64(-280)), _143_v9, _148_globalState)
			_162_v23 = _out4
			_163_v24 = _out5
			_164_v25 = _out6
			_165_v26 = _out7
			var _166_v27 _dafny.MultiSet
			_ = _166_v27
			_166_v27 = _dafny.MultiSetOf(_146_v12, _146_v12, _146_v12)
			_166_v27 = _166_v27
			_161_v22 = _164_v25
			var _index27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index27
			(_147_v13).ArraySet1(_161_v22, (_index27).Int())
		} else {
			var _167_v28 _dafny.CodePoint
			_ = _167_v28
			_167_v28 = _dafny.CodePoint('l')
			var _168_v29 _dafny.MultiSet
			_ = _168_v29
			_168_v29 = _dafny.MultiSetOf(_167_v28)
			var _169_v30 _dafny.Map
			_ = _169_v30
			_169_v30 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_168_v29, _144_v10)
			var _index28 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index28
			(_147_v13).ArraySet1((((_149_v14).Update((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), (_156_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_156_v17), 0))).Int()).(bool))).Cardinality()).Plus(((func() _dafny.Map {
				if _154_cf1 {
					return _169_v30
				}
				return _169_v30
			})()).Cardinality()), (_index28).Int())
			var _170_v31 bool
			_ = _170_v31
			var _171_v32 _dafny.Array
			_ = _171_v32
			var _172_v33 _dafny.Int
			_ = _172_v33
			var _173_v34 _dafny.Int
			_ = _173_v34
			var _out8 bool
			_ = _out8
			var _out9 _dafny.Array
			_ = _out9
			var _out10 _dafny.Int
			_ = _out10
			var _out11 _dafny.Int
			_ = _out11
			_out8, _out9, _out10, _out11 = Companion_Default___.M0((_dafny.Zero).Minus((_139_v5).Cardinality()), _143_v9, _148_globalState)
			_170_v31 = _out8
			_171_v32 = _out9
			_172_v33 = _out10
			_173_v34 = _out11
			var _174_v35 *C0
			_ = _174_v35
			var _nw28 *C0 = New_C0_()
			_ = _nw28
			_nw28.Ctor__(_154_cf1)
			_174_v35 = _nw28
			var _175_v36 _dafny.Map
			_ = _175_v36
			_175_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!(_140_v6), _155_cf0)
			var _176_v37 _dafny.Map
			_ = _176_v37
			_176_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_175_v36, _144_v10)
			var _rhs20 bool = _170_v31
			_ = _rhs20
			var _rhs21 _dafny.Int = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(_173_v34, _dafny.IntOfUint32(((func() _dafny.Sequence {
				if (_176_v37).Contains(_175_v36) {
					return (_176_v37).Get(_175_v36).(_dafny.Sequence)
				}
				return _144_v10
			})()).Cardinality())))
			_ = _rhs21
			_154_cf1 = _rhs20
			_155_cf0 = _rhs21
			var _177_v38 _dafny.Set
			_ = _177_v38
			_177_v38 = _dafny.SetOf(_145_v11)
			var _178_v39 *C0
			_ = _178_v39
			var _nw29 *C0 = New_C0_()
			_ = _nw29
			_nw29.Ctor__((_177_v38).IsSubsetOf(_177_v38))
			_178_v39 = _nw29
		}
		var _index29 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index29
		(_147_v13).ArraySet1(_dafny.IntOfInt64(534), (_index29).Int())
		var _179_v40 D2
		_ = _179_v40
		_179_v40 = Companion_D2_.Create_DC6_(_dafny.SeqOf((_156_v17).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(151), _dafny.ArrayLen((_156_v17), 0))).Int()).(bool)))
		var _180_v41 _dafny.Map
		_ = _180_v41
		_180_v41 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_155_cf0).Times(_dafny.IntOfUint32((_138_v4).Cardinality())), _179_v40)
		var _181_v42 _dafny.CodePoint
		_ = _181_v42
		_181_v42 = _dafny.CodePoint('r')
		var _rhs22 _dafny.Int = (_180_v41).Cardinality()
		_ = _rhs22
		var _rhs23 _dafny.CodePoint = _181_v42
		_ = _rhs23
		var _rhs24 _dafny.Int = Companion_Default___.SafeDivisionInt(_155_cf0, (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
		_ = _rhs24
		var _lhs14 *GlobalState = _148_globalState
		_ = _lhs14
		_155_cf0 = _rhs22
		_lhs14.F4 = _rhs23
		_155_cf0 = _rhs24
	} else if _source3.Is_DC1() {
		var _182___mcc_h2 _dafny.Set = _source3.Get_().(D0_DC1).Cf2
		_ = _182___mcc_h2
		var _183___mcc_h3 bool = _source3.Get_().(D0_DC1).Cf3
		_ = _183___mcc_h3
		var _184_cf3 bool = _183___mcc_h3
		_ = _184_cf3
		var _185_cf2 _dafny.Set = _182___mcc_h2
		_ = _185_cf2
		var _186_v43 _dafny.Set
		_ = _186_v43
		_186_v43 = _dafny.SetOf(!(_140_v6))
		if !(((_186_v43).Difference(_186_v43)).IsDisjointFrom(_dafny.SetOf(!(_140_v6), _184_cf3))) {
			var _pat_let_tv6 = _150_v15
			_ = _pat_let_tv6
			var _187_v44 _dafny.Array
			_ = _187_v44
			var _nwElement0_6 D0 = (func() D0 {
				if _184_cf3 {
					return _151_v16
				}
				return _151_v16
			})()
			_ = _nwElement0_6
			var _nw30 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_6, nil, _dafny.IntOfInt64(12))
			_ = _nw30
			(_nw30).ArraySet1(_nwElement0_6, 0)
			(_nw30).ArraySet1(_151_v16, 1)
			(_nw30).ArraySet1(_151_v16, 2)
			(_nw30).ArraySet1(_151_v16, 3)
			(_nw30).ArraySet1(_151_v16, 4)
			(_nw30).ArraySet1(_151_v16, 5)
			(_nw30).ArraySet1(_151_v16, 6)
			(_nw30).ArraySet1(_151_v16, 7)
			(_nw30).ArraySet1(Companion_D0_.Create_DC2_(_150_v15, (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), _140_v6), 8)
			(_nw30).ArraySet1(Companion_D0_.Create_DC2_(_150_v15, Companion_Default___.Fm0(_145_v11, _184_cf3, _140_v6, _148_globalState), false), 9)
			(_nw30).ArraySet1(_151_v16, 10)
			(_nw30).ArraySet1(func(_pat_let2_0 D0) D0 {
				return func(_188_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let3_0 _dafny.Set) D0 {
						return func(_189_dt__update_hcf4_h0 _dafny.Set) D0 {
							return Companion_D0_.Create_DC2_(_189_dt__update_hcf4_h0, (_188_dt__update__tmp_h0).Dtor_cf5(), (_188_dt__update__tmp_h0).Dtor_cf6())
						}(_pat_let3_0)
					}(_pat_let_tv6)
				}(_pat_let2_0)
			}(Companion_D0_.Create_DC2_(_150_v15, _134_v0, _184_cf3)), 11)
			_187_v44 = _nw30
			_187_v44 = _187_v44
			var _index30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index30
			(_147_v13).ArraySet1((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), (_index30).Int())
			_134_v0 = Companion_Default___.SafeModuloInt((_134_v0).Minus(_134_v0), (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
			_144_v10 = _144_v10
			var _190_v45 *C0
			_ = _190_v45
			var _nw31 *C0 = New_C0_()
			_ = _nw31
			_nw31.Ctor__(Companion_Default___.Fm4(_140_v6, _186_v43, (_dafny.MultiSetOf((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))).Cardinality(), _148_globalState))
			_190_v45 = _nw31
		} else {
			var _191_v46 D1
			_ = _191_v46
			_191_v46 = Companion_D1_.Create_DC5_(((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Cmp((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)) <= 0)
			_191_v46 = _191_v46
			_134_v0 = (_dafny.Zero).Minus(Companion_Default___.SafeModuloInt(((_149_v14).Update(_134_v0, _140_v6)).Cardinality(), _134_v0))
			var _192_v47 *C1
			_ = _192_v47
			var _nw32 *C1 = New_C1_()
			_ = _nw32
			_nw32.Ctor__(_184_cf3, (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
			_192_v47 = _nw32
			var _193_v48 _dafny.Map
			_ = _193_v48
			_193_v48 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_192_v47, (_144_v10).Select((Companion_Default___.SafeIndex(_134_v0, _dafny.IntOfUint32((_144_v10).Cardinality()))).Uint32()).(_dafny.CodePoint))
			var _194_v49 D3
			_ = _194_v49
			_194_v49 = Companion_D3_.Create_DC8_(_192_v47)
			var _195_v50 _dafny.CodePoint
			_ = _195_v50
			_195_v50 = _dafny.CodePoint('w')
			var _index31 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index31
			var _rhs25 _dafny.CodePoint = (func() _dafny.CodePoint {
				if (_193_v48).Contains((_194_v49).Dtor_cf14()) {
					return (_193_v48).Get((_194_v49).Dtor_cf14()).(_dafny.CodePoint)
				}
				return _195_v50
			})()
			_ = _rhs25
			var _rhs26 _dafny.Int = (_192_v47).F15()
			_ = _rhs26
			var _rhs27 bool = Companion_Default___.Fm6((func() bool {
				if Companion_Default___.Fm6(Companion_Default___.Fm4(!(_184_cf3), _186_v43, _dafny.IntOfUint32((_138_v4).Cardinality()), _148_globalState), (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), _148_globalState) {
					return (_192_v47).F14()
				}
				return (_192_v47).F14()
			})(), _134_v0, _148_globalState)
			_ = _rhs27
			var _rhs28 _dafny.Int = Companion_Default___.SafeDivisionInt((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), (_192_v47).F15())
			_ = _rhs28
			var _lhs15 *GlobalState = _148_globalState
			_ = _lhs15
			var _lhs16 _dafny.Array = _147_v13
			_ = _lhs16
			var _lhs17 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _lhs17
			_lhs15.F4 = _rhs25
			_134_v0 = _rhs26
			_140_v6 = _rhs27
			(_lhs16).ArraySet1(_rhs28, (_lhs17).Int())
			(_192_v47).M1(_148_globalState)
			_134_v0 = (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)
		}
		var _196_v51 D1
		_ = _196_v51
		_196_v51 = Companion_D1_.Create_DC4_(_144_v10)
		var _pat_let_tv7 = _144_v10
		_ = _pat_let_tv7
		var _pat_let_tv8 = _144_v10
		_ = _pat_let_tv8
		var _197_v52 _dafny.Array
		_ = _197_v52
		var _nwElement0_7 D1 = Companion_D1_.Create_DC4_(_144_v10)
		_ = _nwElement0_7
		var _nw33 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_7, nil, _dafny.IntOfInt64(7))
		_ = _nw33
		(_nw33).ArraySet1(_nwElement0_7, 0)
		(_nw33).ArraySet1(_196_v51, 1)
		(_nw33).ArraySet1(_196_v51, 2)
		(_nw33).ArraySet1(_196_v51, 3)
		(_nw33).ArraySet1(func(_pat_let4_0 D1) D1 {
			return func(_198_dt__update__tmp_h1 D1) D1 {
				return func(_pat_let5_0 _dafny.Sequence) D1 {
					return func(_199_dt__update_hcf8_h0 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC4_(_199_dt__update_hcf8_h0)
					}(_pat_let5_0)
				}(_pat_let_tv7)
			}(_pat_let4_0)
		}(_196_v51), 4)
		(_nw33).ArraySet1(_196_v51, 5)
		(_nw33).ArraySet1(func(_pat_let6_0 D1) D1 {
			return func(_200_dt__update__tmp_h2 D1) D1 {
				return func(_pat_let7_0 _dafny.Sequence) D1 {
					return func(_201_dt__update_hcf8_h1 _dafny.Sequence) D1 {
						return Companion_D1_.Create_DC4_(_201_dt__update_hcf8_h1)
					}(_pat_let7_0)
				}(_pat_let_tv8)
			}(_pat_let6_0)
		}(_196_v51), 6)
		_197_v52 = _nw33
		var _202_v53 _dafny.Array
		_ = _202_v53
		var _nw34 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(29))
		_ = _nw34
		_202_v53 = _nw34
		var _rhs29 _dafny.Array = _197_v52
		_ = _rhs29
		var _rhs30 _dafny.Array = _147_v13
		_ = _rhs30
		var _rhs31 _dafny.Array = _202_v53
		_ = _rhs31
		_197_v52 = _rhs29
		_147_v13 = _rhs30
		_202_v53 = _rhs31
		var _203_v54 bool
		_ = _203_v54
		var _204_v55 _dafny.Array
		_ = _204_v55
		var _205_v56 _dafny.Int
		_ = _205_v56
		var _206_v57 _dafny.Int
		_ = _206_v57
		var _out12 bool
		_ = _out12
		var _out13 _dafny.Array
		_ = _out13
		var _out14 _dafny.Int
		_ = _out14
		var _out15 _dafny.Int
		_ = _out15
		_out12, _out13, _out14, _out15 = Companion_Default___.M0((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), _143_v9, _148_globalState)
		_203_v54 = _out12
		_204_v55 = _out13
		_205_v56 = _out14
		_206_v57 = _out15
		var _207_v58 _dafny.Sequence
		_ = _207_v58
		_207_v58 = _dafny.SeqOf(_143_v9, _143_v9)
		var _208_v59 bool
		_ = _208_v59
		var _209_v60 _dafny.Array
		_ = _209_v60
		var _210_v61 _dafny.Int
		_ = _210_v61
		var _211_v62 _dafny.Int
		_ = _211_v62
		var _out16 bool
		_ = _out16
		var _out17 _dafny.Array
		_ = _out17
		var _out18 _dafny.Int
		_ = _out18
		var _out19 _dafny.Int
		_ = _out19
		_out16, _out17, _out18, _out19 = Companion_Default___.M0(_dafny.IntOfInt64(-643), (_207_v58).Select((Companion_Default___.SafeIndex(_206_v57, _dafny.IntOfUint32((_207_v58).Cardinality()))).Uint32()).(_dafny.Sequence), _148_globalState)
		_208_v59 = _out16
		_209_v60 = _out17
		_210_v61 = _out18
		_211_v62 = _out19
	} else if _source3.Is_DC2() {
		var _212___mcc_h4 _dafny.Set = _source3.Get_().(D0_DC2).Cf4
		_ = _212___mcc_h4
		var _213___mcc_h5 _dafny.Int = _source3.Get_().(D0_DC2).Cf5
		_ = _213___mcc_h5
		var _214___mcc_h6 bool = _source3.Get_().(D0_DC2).Cf6
		_ = _214___mcc_h6
		var _215_cf6 bool = _214___mcc_h6
		_ = _215_cf6
		var _216_cf5 _dafny.Int = _213___mcc_h5
		_ = _216_cf5
		var _217_cf4 _dafny.Set = _212___mcc_h4
		_ = _217_cf4
		var _218_v63 _dafny.MultiSet
		_ = _218_v63
		_218_v63 = _dafny.MultiSetOf(_143_v9, _143_v9, _143_v9)
		var _219_v64 D0
		_ = _219_v64
		_219_v64 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_143_v9).Cardinality()), _140_v6)
		var _220_v65 _dafny.Sequence
		_ = _220_v65
		_220_v65 = _dafny.SeqOf(_143_v9)
		var _221_v66 _dafny.Map
		_ = _221_v66
		_221_v66 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_219_v64).Dtor_cf1(), _220_v65)
		var _222_v67 _dafny.Array
		_ = _222_v67
		var _nwElement0_8 _dafny.MultiSet = _218_v63
		_ = _nwElement0_8
		var _nw35 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_8, nil, _dafny.IntOfInt64(8))
		_ = _nw35
		(_nw35).ArraySet1(_nwElement0_8, 0)
		(_nw35).ArraySet1((_218_v63).Union((_218_v63).Update(_143_v9, Companion_Default___.Abs(_dafny.IntOfInt64(596)))), 1)
		(_nw35).ArraySet1((func() _dafny.MultiSet {
			if _215_cf6 {
				return _218_v63
			}
			return _218_v63
		})(), 2)
		(_nw35).ArraySet1(_218_v63, 3)
		(_nw35).ArraySet1(_dafny.MultiSetOf(_143_v9, _143_v9), 4)
		(_nw35).ArraySet1(_218_v63, 5)
		(_nw35).ArraySet1((func() _dafny.MultiSet {
			if _215_cf6 {
				return _dafny.MultiSetFromSeq((func() _dafny.Sequence {
					if (_221_v66).Contains(_215_cf6) {
						return (_221_v66).Get(_215_cf6).(_dafny.Sequence)
					}
					return _220_v65
				})())
			}
			return _218_v63
		})(), 6)
		(_nw35).ArraySet1(_218_v63, 7)
		_222_v67 = _nw35
		var _index32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_222_v67), 0))
		_ = _index32
		(_222_v67).ArraySet1(_dafny.MultiSetOf(_143_v9, _143_v9, _143_v9), (_index32).Int())
		var _223_v68 _dafny.Sequence
		_ = _223_v68
		_223_v68 = _dafny.SeqOf(_218_v63, (_218_v63).Union(_218_v63), (_218_v63).Union(_218_v63), _218_v63, _dafny.MultiSetOf(_143_v9, _143_v9))
		var _index33 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(276), _dafny.ArrayLen((_222_v67), 0))
		_ = _index33
		(_222_v67).ArraySet1((_223_v68).Select((Companion_Default___.SafeIndex(Companion_Default___.SafeDivisionInt((_138_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(504), _dafny.IntOfUint32((_138_v4).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf(false)).Cardinality())), _dafny.IntOfUint32((_223_v68).Cardinality()))).Uint32()).(_dafny.MultiSet), (_index33).Int())
		if _140_v6 {
			var _224_v69 *C0
			_ = _224_v69
			var _nw36 *C0 = New_C0_()
			_ = _nw36
			_nw36.Ctor__(Companion_Default___.Fm6(_140_v6, _dafny.IntOfInt64(-901), _148_globalState))
			_224_v69 = _nw36
			_144_v10 = (_224_v69).Fm5(_215_cf6, _216_cf5, Companion_Default___.SafeDivisionInt((_145_v11).Cardinality(), _216_cf5), (_224_v69).F16(), _148_globalState)
			var _225_v70 D4
			_ = _225_v70
			_225_v70 = Companion_D4_.Create_DC11_(_139_v5)
			var _226_v71 *C1
			_ = _226_v71
			var _nw37 *C1 = New_C1_()
			_ = _nw37
			_nw37.Ctor__(Companion_Default___.Fm6(!(true), _dafny.IntOfInt64(999), _148_globalState), ((_225_v70).Dtor_cf17()).Cardinality())
			_226_v71 = _nw37
			(_224_v69).M1(_148_globalState)
			var _227_v72 *C1
			_ = _227_v72
			var _nw38 *C1 = New_C1_()
			_ = _nw38
			_nw38.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_144_v10, _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(436))).Uint32(), func(coer5 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg5 _dafny.Int) interface{} {
					return coer5(arg5)
				}
			}(func(_228_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('w')
			}))), (_145_v11).Cardinality())
			_227_v72 = _nw38
		} else {
			var _229_v73 *C0
			_ = _229_v73
			var _nw39 *C0 = New_C0_()
			_ = _nw39
			_nw39.Ctor__(_215_cf6)
			_229_v73 = _nw39
			var _230_v74 _dafny.Sequence
			_ = _230_v74
			_230_v74 = _dafny.SeqOf(_229_v73, _229_v73)
			var _index34 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index34
			var _rhs32 _dafny.Int = ((_151_v16).Dtor_cf5()).Plus((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
			_ = _rhs32
			var _rhs33 *C0 = _229_v73
			_ = _rhs33
			var _rhs34 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_138_v4, _dafny.Companion_Sequence_.Concatenate(_138_v4, _138_v4))
			_ = _rhs34
			var _rhs35 _dafny.Int = _216_cf5
			_ = _rhs35
			var _rhs36 _dafny.Sequence = _dafny.SeqOf(_229_v73, _229_v73, _229_v73)
			_ = _rhs36
			var _lhs18 _dafny.Array = _147_v13
			_ = _lhs18
			var _lhs19 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _lhs19
			_216_cf5 = _rhs32
			_229_v73 = _rhs33
			_138_v4 = _rhs34
			(_lhs18).ArraySet1(_rhs35, (_lhs19).Int())
			_230_v74 = _rhs36
			var _index35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index35
			(_147_v13).ArraySet1(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_216_cf5), _138_v4)).Cardinality()), (_index35).Int())
			_215_cf6 = _215_cf6
			_215_cf6 = !(!(_140_v6)) || ((_229_v73).F16())
			_134_v0 = Companion_Default___.SafeModuloInt(_134_v0, _134_v0)
		}
		var _index36 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index36
		(_147_v13).ArraySet1(_134_v0, (_index36).Int())
		if _215_cf6 {
			_134_v0 = (_219_v64).Dtor_cf0()
			var _231_v75 _dafny.Array
			_ = _231_v75
			var _len0_3 _dafny.Int = _dafny.IntOfInt64(23)
			_ = _len0_3
			var _nw40 _dafny.Array
			_ = _nw40
			if _len0_3.Cmp(_dafny.Zero) == 0 {
				_nw40 = _dafny.NewArray(_len0_3)
			} else {
				var _init3 func(_dafny.Int) bool = (func(_232_v6 bool) func(_dafny.Int) bool {
					return func(_233_i1 _dafny.Int) bool {
						return !(_232_v6)
					}
				})(_140_v6)
				_ = _init3
				var _element0_3 = _init3(_dafny.Zero)
				_ = _element0_3
				_nw40 = _dafny.NewArrayFromExample(_element0_3, nil, _len0_3)
				(_nw40).ArraySet1(_element0_3, 0)
				var _nativeLen0_3 = (_len0_3).Int()
				_ = _nativeLen0_3
				for _i0_3 := 1; _i0_3 < _nativeLen0_3; _i0_3++ {
					(_nw40).ArraySet1(_init3(_dafny.IntOf(_i0_3)), _i0_3)
				}
			}
			_231_v75 = _nw40
			var _index37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_231_v75), 0))
			_ = _index37
			(_231_v75).ArraySet1(_215_cf6, (_index37).Int())
			var _index38 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_231_v75), 0))
			_ = _index38
			(_231_v75).ArraySet1(!(_140_v6), (_index38).Int())
			_143_v9 = _dafny.Companion_Sequence_.Update(_143_v9, (Companion_Default___.SafeIndex(_216_cf5, _dafny.IntOfUint32((_143_v9).Cardinality()))).Uint32(), false)
			_140_v6 = (func() bool {
				if (_149_v14).Contains(_dafny.IntOfInt64(157)) {
					return (_149_v14).Get(_dafny.IntOfInt64(157)).(bool)
				}
				return (_231_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_231_v75), 0))).Int()).(bool)
			})()
			var _234_v76 _dafny.CodePoint
			_ = _234_v76
			_234_v76 = _dafny.CodePoint('j')
			var _235_v77 _dafny.Sequence
			_ = _235_v77
			_235_v77 = _dafny.SeqOf(_144_v10)
			var _236_v78 _dafny.Array
			_ = _236_v78
			var _nw41 _dafny.Array = _dafny.NewArrayWithValue(Companion_D0_.Default(), _dafny.IntOfInt64(10))
			_ = _nw41
			_236_v78 = _nw41
			var _237_v79 _dafny.Map
			_ = _237_v79
			_237_v79 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_134_v0, _141_v7)
			var _238_v80 _dafny.Map
			_ = _238_v80
			_238_v80 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_236_v78, _237_v79)
			var _rhs37 _dafny.CodePoint = _234_v76
			_ = _rhs37
			var _rhs38 _dafny.Array = _231_v75
			_ = _rhs38
			var _rhs39 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_144_v10, (_235_v77).Select((Companion_Default___.SafeIndex(_216_cf5, _dafny.IntOfUint32((_235_v77).Cardinality()))).Uint32()).(_dafny.Sequence))
			_ = _rhs39
			var _rhs40 bool = ((func() _dafny.Map {
				if (_238_v80).Contains(_236_v78) {
					return (_238_v80).Get(_236_v78).(_dafny.Map)
				}
				return _237_v79
			})()).Contains(_dafny.IntOfUint32((_dafny.SeqOf((_231_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_231_v75), 0))).Int()).(bool))).Cardinality()))
			_ = _rhs40
			var _rhs41 bool = (_231_v75).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(380), _dafny.ArrayLen((_231_v75), 0))).Int()).(bool)
			_ = _rhs41
			var _lhs20 *GlobalState = _148_globalState
			_ = _lhs20
			_lhs20.F4 = _rhs37
			_231_v75 = _rhs38
			_144_v10 = _rhs39
			_140_v6 = _rhs40
			_140_v6 = _rhs41
		} else {
			var _239_v81 _dafny.CodePoint
			_ = _239_v81
			_239_v81 = _dafny.CodePoint('h')
			var _index39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_136_v2), 0))
			_ = _index39
			(_136_v2).ArraySet1CodePoint(_239_v81, (_index39).Int())
			var _index40 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_136_v2), 0))
			_ = _index40
			(_136_v2).ArraySet1CodePoint(_239_v81, (_index40).Int())
			_216_cf5 = (_dafny.Zero).Minus((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
			var _240_v82 _dafny.MultiSet
			_ = _240_v82
			_240_v82 = _dafny.MultiSetOf((_149_v14).Cardinality(), (_dafny.Zero).Minus(_216_cf5))
			var _index41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index41
			var _rhs42 _dafny.Array = _147_v13
			_ = _rhs42
			var _rhs43 bool = !(false)
			_ = _rhs43
			var _rhs44 _dafny.Int = _216_cf5
			_ = _rhs44
			var _rhs45 _dafny.MultiSet = _240_v82
			_ = _rhs45
			var _lhs21 *GlobalState = _148_globalState
			_ = _lhs21
			var _lhs22 _dafny.Array = _147_v13
			_ = _lhs22
			var _lhs23 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _lhs23
			_lhs21.F12 = _rhs42
			_140_v6 = _rhs43
			(_lhs22).ArraySet1(_rhs44, (_lhs23).Int())
			_240_v82 = _rhs45
			var _241_v83 bool
			_ = _241_v83
			var _242_v84 _dafny.Array
			_ = _242_v84
			var _243_v85 _dafny.Int
			_ = _243_v85
			var _244_v86 _dafny.Int
			_ = _244_v86
			var _out20 bool
			_ = _out20
			var _out21 _dafny.Array
			_ = _out21
			var _out22 _dafny.Int
			_ = _out22
			var _out23 _dafny.Int
			_ = _out23
			_out20, _out21, _out22, _out23 = Companion_Default___.M0((_dafny.Zero).Minus((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)), Companion_Default___.Fm10(!(_215_cf6), (_136_v2).ArrayGet1CodePoint((Companion_Default___.SafeIndex(_dafny.IntOfInt64(172), _dafny.ArrayLen((_136_v2), 0))).Int()), _134_v0, _148_globalState), _148_globalState)
			_241_v83 = _out20
			_242_v84 = _out21
			_243_v85 = _out22
			_244_v86 = _out23
			_140_v6 = (_140_v6) == (true)
		}
	} else {
		var _245___mcc_h7 _dafny.CodePoint = _source3.Get_().(D0_DC3).Cf7
		_ = _245___mcc_h7
		var _246_cf7 _dafny.CodePoint = _245___mcc_h7
		_ = _246_cf7
		var _247_v87 _dafny.Array
		_ = _247_v87
		var _nw42 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(8))
		_ = _nw42
		_247_v87 = _nw42
		_247_v87 = _247_v87
		_140_v6 = (_134_v0).Cmp(_134_v0) > 0
		var _248_v88 _dafny.Map
		_ = _248_v88
		_248_v88 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_143_v9).Select((Companion_Default___.SafeIndex((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_143_v9).Cardinality()))).Uint32()).(bool), (func() _dafny.Map {
			if _140_v6 {
				return _149_v14
			}
			return _149_v14
		})())
		_248_v88 = (_248_v88).Merge(_248_v88)
		_140_v6 = !(((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Cmp(Companion_Default___.SafeDivisionInt(_134_v0, _134_v0)) > 0)
	}
	_140_v6 = (_140_v6) || (((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Cmp((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)) != 0)
	var _249_i2 _dafny.Int
	_ = _249_i2
	_249_i2 = _dafny.Zero
	{
		for _140_v6 {
			{
				if (_249_i2).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L1
				}
				_249_i2 = (_249_i2).Plus(_dafny.One)
				_146_v12 = (_146_v12).Update((_134_v0).Times(_134_v0), ((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Minus((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)))
				_140_v6 = ((func() _dafny.Map {
					if _140_v6 {
						return _149_v14
					}
					return _149_v14
				})()).Contains(_134_v0)
				var _250_v89 _dafny.Array
				_ = _250_v89
				var _len0_4 _dafny.Int = _dafny.IntOfInt64(11)
				_ = _len0_4
				var _nw43 _dafny.Array
				_ = _nw43
				if _len0_4.Cmp(_dafny.Zero) == 0 {
					_nw43 = _dafny.NewArray(_len0_4)
				} else {
					var _init4 func(_dafny.Int) bool = (func(_251_v0 _dafny.Int, _252_v13 _dafny.Array) func(_dafny.Int) bool {
						return func(_253_i3 _dafny.Int) bool {
							return (_251_v0).Cmp((_252_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_252_v13), 0))).Int()).(_dafny.Int)) < 0
						}
					})(_134_v0, _147_v13)
					_ = _init4
					var _element0_4 = _init4(_dafny.Zero)
					_ = _element0_4
					_nw43 = _dafny.NewArrayFromExample(_element0_4, nil, _len0_4)
					(_nw43).ArraySet1(_element0_4, 0)
					var _nativeLen0_4 = (_len0_4).Int()
					_ = _nativeLen0_4
					for _i0_4 := 1; _i0_4 < _nativeLen0_4; _i0_4++ {
						(_nw43).ArraySet1(_init4(_dafny.IntOf(_i0_4)), _i0_4)
					}
				}
				_250_v89 = _nw43
				var _len0_5 _dafny.Int = _dafny.IntOfInt64(16)
				_ = _len0_5
				var _nw44 _dafny.Array
				_ = _nw44
				if _len0_5.Cmp(_dafny.Zero) == 0 {
					_nw44 = _dafny.NewArray(_len0_5)
				} else {
					var _init5 func(_dafny.Int) bool = (func(_254_v6 bool) func(_dafny.Int) bool {
						return func(_255_i4 _dafny.Int) bool {
							return _254_v6
						}
					})(_140_v6)
					_ = _init5
					var _element0_5 = _init5(_dafny.Zero)
					_ = _element0_5
					_nw44 = _dafny.NewArrayFromExample(_element0_5, nil, _len0_5)
					(_nw44).ArraySet1(_element0_5, 0)
					var _nativeLen0_5 = (_len0_5).Int()
					_ = _nativeLen0_5
					for _i0_5 := 1; _i0_5 < _nativeLen0_5; _i0_5++ {
						(_nw44).ArraySet1(_init5(_dafny.IntOf(_i0_5)), _i0_5)
					}
				}
				_250_v89 = _nw44
				var _256_v90 *C1
				_ = _256_v90
				var _nw45 *C1 = New_C1_()
				_ = _nw45
				_nw45.Ctor__(_140_v6, _dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(561))).Uint32(), func(coer6 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg6 _dafny.Int) interface{} {
						return coer6(arg6)
					}
				}(func(_257_i5 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('d')
				}))).Cardinality()))
				_256_v90 = _nw45
				var _index42 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
				_ = _index42
				var _rhs46 *C1 = _256_v90
				_ = _rhs46
				var _rhs47 _dafny.Int = _dafny.IntOfInt64(509)
				_ = _rhs47
				var _lhs24 _dafny.Array = _147_v13
				_ = _lhs24
				var _lhs25 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
				_ = _lhs25
				_256_v90 = _rhs46
				(_lhs24).ArraySet1(_rhs47, (_lhs25).Int())
				goto C1
			}
		C1:
		}
		goto L1
	}
L1:
	var _258_v91 _dafny.Array
	_ = _258_v91
	var _len0_6 _dafny.Int = _dafny.IntOfInt64(8)
	_ = _len0_6
	var _nw46 _dafny.Array
	_ = _nw46
	if _len0_6.Cmp(_dafny.Zero) == 0 {
		_nw46 = _dafny.NewArray(_len0_6)
	} else {
		var _init6 func(_dafny.Int) bool = (func(_259_v6 bool) func(_dafny.Int) bool {
			return func(_260_i6 _dafny.Int) bool {
				return _259_v6
			}
		})(_140_v6)
		_ = _init6
		var _element0_6 = _init6(_dafny.Zero)
		_ = _element0_6
		_nw46 = _dafny.NewArrayFromExample(_element0_6, nil, _len0_6)
		(_nw46).ArraySet1(_element0_6, 0)
		var _nativeLen0_6 = (_len0_6).Int()
		_ = _nativeLen0_6
		for _i0_6 := 1; _i0_6 < _nativeLen0_6; _i0_6++ {
			(_nw46).ArraySet1(_init6(_dafny.IntOf(_i0_6)), _i0_6)
		}
	}
	_258_v91 = _nw46
	var _261_v92 *C0
	_ = _261_v92
	var _nw47 *C0 = New_C0_()
	_ = _nw47
	_nw47.Ctor__(false)
	_261_v92 = _nw47
	var _262_v93 _dafny.Array
	_ = _262_v93
	var _nwElement0_9 *C0 = (func() *C0 {
		if _140_v6 {
			return _261_v92
		}
		return _261_v92
	})()
	_ = _nwElement0_9
	var _nw48 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_9, nil, _dafny.IntOfInt64(2))
	_ = _nw48
	(_nw48).ArraySet1(_nwElement0_9, 0)
	(_nw48).ArraySet1(_261_v92, 1)
	_262_v93 = _nw48
	var _index43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))
	_ = _index43
	(_262_v93).ArraySet1(_261_v92, (_index43).Int())
	var _263_v94 *C1
	_ = _263_v94
	var _nw49 *C1 = New_C1_()
	_ = _nw49
	_nw49.Ctor__(_140_v6, _134_v0)
	_263_v94 = _nw49
	var _264_v95 _dafny.Sequence
	_ = _264_v95
	_264_v95 = _dafny.SeqOf(_263_v94, _263_v94)
	var _265_v96 _dafny.CodePoint
	_ = _265_v96
	_265_v96 = _dafny.CodePoint('c')
	var _266_v97 _dafny.Map
	_ = _266_v97
	_266_v97 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_264_v95).Cardinality()), _265_v96)
	var _index44 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))
	_ = _index44
	var _rhs48 _dafny.Array = _258_v91
	_ = _rhs48
	var _rhs49 *C0 = _261_v92
	_ = _rhs49
	var _rhs50 _dafny.CodePoint = (func() _dafny.CodePoint {
		if (_266_v97).Contains(_134_v0) {
			return (_266_v97).Get(_134_v0).(_dafny.CodePoint)
		}
		return (_144_v10).Select((Companion_Default___.SafeIndex((_138_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-258), _dafny.IntOfUint32((_138_v4).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_144_v10).Cardinality()))).Uint32()).(_dafny.CodePoint)
	})()
	_ = _rhs50
	var _lhs26 _dafny.Array = _262_v93
	_ = _lhs26
	var _lhs27 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))
	_ = _lhs27
	var _lhs28 *GlobalState = _148_globalState
	_ = _lhs28
	_258_v91 = _rhs48
	(_lhs26).ArraySet1(_rhs49, (_lhs27).Int())
	_lhs28.F4 = _rhs50
	(_261_v92).M1(_148_globalState)
	var _source4 D0 = _151_v16
	_ = _source4
	if _source4.Is_DC0() {
		var _267___mcc_h8 _dafny.Int = _source4.Get_().(D0_DC0).Cf0
		_ = _267___mcc_h8
		var _268___mcc_h9 bool = _source4.Get_().(D0_DC0).Cf1
		_ = _268___mcc_h9
		var _269_cf1 bool = _268___mcc_h9
		_ = _269_cf1
		var _270_cf0 _dafny.Int = _267___mcc_h8
		_ = _270_cf0
		_134_v0 = _270_cf0
		var _271_v98 bool
		_ = _271_v98
		var _272_v99 _dafny.Array
		_ = _272_v99
		var _273_v100 _dafny.Int
		_ = _273_v100
		var _274_v101 _dafny.Int
		_ = _274_v101
		var _out24 bool
		_ = _out24
		var _out25 _dafny.Array
		_ = _out25
		var _out26 _dafny.Int
		_ = _out26
		var _out27 _dafny.Int
		_ = _out27
		_out24, _out25, _out26, _out27 = Companion_Default___.M0((_263_v94).F15(), _dafny.Companion_Sequence_.Concatenate(_143_v9, _dafny.SeqOf(_269_cf1, !(_269_cf1), _140_v6, (_263_v94).F14(), (_263_v94).F14())), _148_globalState)
		_271_v98 = _out24
		_272_v99 = _out25
		_273_v100 = _out26
		_274_v101 = _out27
		var _index45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index45
		(_147_v13).ArraySet1(Companion_Default___.SafeDivisionInt(_dafny.IntOfInt64(391), _270_cf0), (_index45).Int())
		_144_v10 = _dafny.UnicodeSeqOfUtf8Bytes("fxrlkl")
	} else if _source4.Is_DC1() {
		var _275___mcc_h10 _dafny.Set = _source4.Get_().(D0_DC1).Cf2
		_ = _275___mcc_h10
		var _276___mcc_h11 bool = _source4.Get_().(D0_DC1).Cf3
		_ = _276___mcc_h11
		var _277_cf3 bool = _276___mcc_h11
		_ = _277_cf3
		var _278_cf2 _dafny.Set = _275___mcc_h10
		_ = _278_cf2
		var _279_v102 _dafny.Array
		_ = _279_v102
		var _nw50 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySet, _dafny.IntOfInt64(2))
		_ = _nw50
		_279_v102 = _nw50
		var _280_v103 _dafny.Set
		_ = _280_v103
		_280_v103 = _dafny.SetOf(false, (_263_v94).F14())
		var _index46 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_279_v102), 0))
		_ = _index46
		(_279_v102).ArraySet1((_dafny.SetOf(_277_cf3, (_261_v92).F16())).Difference(_280_v103), (_index46).Int())
		var _index47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(723), _dafny.ArrayLen((_279_v102), 0))
		_ = _index47
		(_279_v102).ArraySet1((func() _dafny.Set {
			if (_263_v94).F14() {
				return _280_v103
			}
			return _280_v103
		})(), (_index47).Int())
		var _index48 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index48
		(_147_v13).ArraySet1((_263_v94).F15(), (_index48).Int())
		var _281_v104 D1
		_ = _281_v104
		_281_v104 = Companion_D1_.Create_DC5_(true)
		var _282_v105 _dafny.Map
		_ = _282_v105
		_282_v105 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_D0_.Create_DC3_(_265_v96), (_263_v94).F15())
		var _283_v106 D0
		_ = _283_v106
		_283_v106 = Companion_D0_.Create_DC3_(_265_v96)
		var _284_v107 *C1
		_ = _284_v107
		var _nw51 *C1 = New_C1_()
		_ = _nw51
		_nw51.Ctor__((_281_v104).Dtor_cf9(), (func() _dafny.Int {
			if (_282_v105).Contains(_283_v106) {
				return (_282_v105).Get(_283_v106).(_dafny.Int)
			}
			return (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)
		})())
		_284_v107 = _nw51
		var _285_v108 _dafny.Array
		_ = _285_v108
		var _nw52 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(11))
		_ = _nw52
		_285_v108 = _nw52
		var _286_v109 T0
		_ = _286_v109
		var _nw53 *C0 = New_C0_()
		_ = _nw53
		_nw53.Ctor__((_261_v92).F16())
		_286_v109 = _nw53
		var _287_v110 _dafny.Map
		_ = _287_v110
		_287_v110 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.SetOf((_263_v94).F15()), _286_v109)
		var _index49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_285_v108), 0))
		_ = _index49
		(_285_v108).ArraySet1(_287_v110, (_index49).Int())
		var _index50 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_285_v108), 0))
		_ = _index50
		var _index51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index51
		var _rhs51 _dafny.Int = (_263_v94).F15()
		_ = _rhs51
		var _rhs52 _dafny.Int = (func() _dafny.Map {
			var _coll3 = _dafny.NewMapBuilder()
			_ = _coll3
			for _iter3 := _dafny.Iterate((_dafny.SeqOf(_144_v10, _144_v10)).Elements()); ; {
				_compr_3, _ok3 := _iter3()
				if !_ok3 {
					break
				}
				var _288_v111 _dafny.Sequence
				_288_v111 = interface{}(_compr_3).(_dafny.Sequence)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_144_v10, _144_v10), _288_v111) {
					_coll3.Add(_288_v111, Companion_D0_.Create_DC1_(_278_cf2, (_263_v94).F14()))
				}
			}
			return _coll3.ToMap()
		}()).Cardinality()
		_ = _rhs52
		var _rhs53 _dafny.Map = ((_287_v110).Merge(_287_v110)).Merge(_287_v110)
		_ = _rhs53
		var _rhs54 _dafny.Int = (_141_v7).Cardinality()
		_ = _rhs54
		var _lhs29 _dafny.Array = _285_v108
		_ = _lhs29
		var _lhs30 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(337), _dafny.ArrayLen((_285_v108), 0))
		_ = _lhs30
		var _lhs31 _dafny.Array = _147_v13
		_ = _lhs31
		var _lhs32 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _lhs32
		_134_v0 = _rhs51
		_134_v0 = _rhs52
		(_lhs29).ArraySet1(_rhs53, (_lhs30).Int())
		(_lhs31).ArraySet1(_rhs54, (_lhs32).Int())
	} else if _source4.Is_DC2() {
		var _289___mcc_h12 _dafny.Set = _source4.Get_().(D0_DC2).Cf4
		_ = _289___mcc_h12
		var _290___mcc_h13 _dafny.Int = _source4.Get_().(D0_DC2).Cf5
		_ = _290___mcc_h13
		var _291___mcc_h14 bool = _source4.Get_().(D0_DC2).Cf6
		_ = _291___mcc_h14
		var _292_cf6 bool = _291___mcc_h14
		_ = _292_cf6
		var _293_cf5 _dafny.Int = _290___mcc_h13
		_ = _293_cf5
		var _294_cf4 _dafny.Set = _289___mcc_h12
		_ = _294_cf4
		var _295_v112 *C0
		_ = _295_v112
		var _nw54 *C0 = New_C0_()
		_ = _nw54
		_nw54.Ctor__(_140_v6)
		_295_v112 = _nw54
		_140_v6 = !(((_295_v112).F16()) && ((_261_v92).F16()))
		_140_v6 = ((Companion_Default___.Fm9(_265_v96, _134_v0, _dafny.IntOfInt64(904), _dafny.NewMapBuilder().ToMap().UpdateUnsafe(!((_263_v94).F14()), _293_cf5), _148_globalState)).Cardinality()).Cmp((_263_v94).F15()) >= 0
		var _296_v113 _dafny.Array
		_ = _296_v113
		var _nw55 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(13))
		_ = _nw55
		_296_v113 = _nw55
		var _297_v114 _dafny.Map
		_ = _297_v114
		_297_v114 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v10, _147_v13)
		var _index52 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_296_v113), 0))
		_ = _index52
		(_296_v113).ArraySet1(_297_v114, (_index52).Int())
		var _index53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(703), _dafny.ArrayLen((_296_v113), 0))
		_ = _index53
		(_296_v113).ArraySet1((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_144_v10, _147_v13)).Merge((_297_v114).Merge(_297_v114)), (_index53).Int())
	} else {
		var _298___mcc_h15 _dafny.CodePoint = _source4.Get_().(D0_DC3).Cf7
		_ = _298___mcc_h15
		var _299_cf7 _dafny.CodePoint = _298___mcc_h15
		_ = _299_cf7
		var _index54 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index54
		(_147_v13).ArraySet1((_263_v94).F15(), (_index54).Int())
		if (_263_v94).F14() {
			var _300_v115 D3
			_ = _300_v115
			_300_v115 = Companion_D3_.Create_DC10_(Companion_D3_.Create_DC9_(_299_cf7))
			var _pat_let_tv9 = _263_v94
			_ = _pat_let_tv9
			var _301_v116 _dafny.Map
			_ = _301_v116
			_301_v116 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.SeqOf((_261_v92).F16(), (_143_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_264_v95).Cardinality()), _dafny.IntOfUint32((_143_v9).Cardinality()))).Uint32()).(bool))).Cardinality()), func(_pat_let8_0 D3) D3 {
				return func(_302_dt__update__tmp_h3 D3) D3 {
					return func(_pat_let9_0 D3) D3 {
						return func(_303_dt__update_hcf16_h0 D3) D3 {
							return Companion_D3_.Create_DC10_(_303_dt__update_hcf16_h0)
						}(_pat_let9_0)
					}(Companion_D3_.Create_DC8_(_pat_let_tv9))
				}(_pat_let8_0)
			}(_300_v115))
			_301_v116 = (_301_v116).Update((_263_v94).F15(), Companion_D3_.Create_DC10_(Companion_D3_.Create_DC9_(_265_v96)))
			var _304_v117 T0
			_ = _304_v117
			var _nw56 *C1 = New_C1_()
			_ = _nw56
			_nw56.Ctor__((_261_v92).F16(), (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
			_304_v117 = _nw56
			_304_v117 = _304_v117
			var _305_v118 D3
			_ = _305_v118
			_305_v118 = Companion_D3_.Create_DC8_(_263_v94)
			_305_v118 = _305_v118
			_144_v10 = _144_v10
			var _306_v119 _dafny.Map
			_ = _306_v119
			_306_v119 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm6((_261_v92).F16(), (_263_v94).F15(), _148_globalState), _143_v9)
			_306_v119 = _306_v119
		} else {
			var _307_v120 _dafny.MultiSet
			_ = _307_v120
			_307_v120 = _dafny.MultiSetOf(_299_cf7)
			var _308_v121 _dafny.Map
			_ = _308_v121
			_308_v121 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_140_v6, _307_v120)
			var _309_v122 _dafny.Set
			_ = _309_v122
			_309_v122 = _dafny.SetOf(false, !((_263_v94).F14()))
			var _310_v123 _dafny.Sequence
			_ = _310_v123
			_310_v123 = _dafny.SeqOf((func() _dafny.MultiSet {
				if (_308_v121).Contains((_261_v92).F16()) {
					return (_308_v121).Get((_261_v92).F16()).(_dafny.MultiSet)
				}
				return (_307_v120).Update(_dafny.CodePoint('j'), Companion_Default___.Abs((_309_v122).Cardinality()))
			})())
			var _index55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index55
			(_147_v13).ArraySet1(Companion_Default___.SafeModuloInt(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_310_v123, _310_v123)).Cardinality()), (_263_v94).F15()), (_index55).Int())
			_134_v0 = (_263_v94).F15()
			var _311_v124 D1
			_ = _311_v124
			_311_v124 = Companion_D1_.Create_DC5_(_140_v6)
			var _312_v125 _dafny.Map
			_ = _312_v125
			_312_v125 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v94).F15(), _262_v93)
			var _index56 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))
			_ = _index56
			var _rhs55 _dafny.Array = _147_v13
			_ = _rhs55
			var _rhs56 D1 = _311_v124
			_ = _rhs56
			var _rhs57 bool = (func() bool {
				if (_145_v11).Contains((func() bool {
					if (_145_v11).Contains((_261_v92).F16()) {
						return (_145_v11).Get((_261_v92).F16()).(bool)
					}
					return (_261_v92).F16()
				})()) {
					return (_145_v11).Get((func() bool {
						if (_145_v11).Contains((_261_v92).F16()) {
							return (_145_v11).Get((_261_v92).F16()).(bool)
						}
						return (_261_v92).F16()
					})()).(bool)
				}
				return (_261_v92).F16()
			})()
			_ = _rhs57
			var _rhs58 _dafny.Map = (_312_v125).Merge((_312_v125).Update(_134_v0, _262_v93))
			_ = _rhs58
			var _rhs59 *C0 = _261_v92
			_ = _rhs59
			var _lhs33 *GlobalState = _148_globalState
			_ = _lhs33
			var _lhs34 _dafny.Array = _262_v93
			_ = _lhs34
			var _lhs35 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))
			_ = _lhs35
			_lhs33.F12 = _rhs55
			_311_v124 = _rhs56
			_140_v6 = _rhs57
			_312_v125 = _rhs58
			(_lhs34).ArraySet1(_rhs59, (_lhs35).Int())
			_134_v0 = _134_v0
			var _index57 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index57
			var _rhs60 _dafny.Int = (Companion_Default___.SafeDivisionInt((_263_v94).F15(), _134_v0)).Plus(Companion_Default___.SafeDivisionInt((_263_v94).F15(), (_263_v94).F15()))
			_ = _rhs60
			var _rhs61 _dafny.Map = _145_v11
			_ = _rhs61
			var _rhs62 _dafny.CodePoint = _299_cf7
			_ = _rhs62
			var _rhs63 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_144_v10, _144_v10)
			_ = _rhs63
			var _lhs36 _dafny.Array = _147_v13
			_ = _lhs36
			var _lhs37 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _lhs37
			(_lhs36).ArraySet1(_rhs60, (_lhs37).Int())
			_145_v11 = _rhs61
			_299_cf7 = _rhs62
			_144_v10 = _rhs63
		}
		_147_v13 = _147_v13
		_140_v6 = _dafny.Companion_Sequence_.IsProperPrefixOf(_138_v4, _138_v4)
	}
	var _313_v126 _dafny.MultiSet
	_ = _313_v126
	_313_v126 = _dafny.MultiSetOf(_265_v96)
	_313_v126 = ((_313_v126).Intersection(Companion_Default___.Fm11(_dafny.UnicodeSeqOfUtf8Bytes("avyxie"), _148_globalState))).Update(_265_v96, Companion_Default___.Abs(((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Times((_263_v94).F15())))
	_140_v6 = ((_263_v94).F15()).Cmp(_dafny.IntOfUint32((_144_v10).Cardinality())) > 0
	var _pat_let_tv10 = _151_v16
	_ = _pat_let_tv10
	var _pat_let_tv11 = _150_v15
	_ = _pat_let_tv11
	var _pat_let_tv12 = _263_v94
	_ = _pat_let_tv12
	var _pat_let_tv13 = _151_v16
	_ = _pat_let_tv13
	var _source5 D0 = func(_source6 D4) D0 {
		if _source6.Is_DC12() {
			var _314___mcc_h24 _dafny.Int = _source6.Get_().(D4_DC12).Cf18
			_ = _314___mcc_h24
			var _315___mcc_h25 _dafny.CodePoint = _source6.Get_().(D4_DC12).Cf19
			_ = _315___mcc_h25
			var _316___mcc_h26 _dafny.MultiSet = _source6.Get_().(D4_DC12).Cf20
			_ = _316___mcc_h26
			var _317_cf20 _dafny.MultiSet = _316___mcc_h26
			_ = _317_cf20
			var _318_cf19 _dafny.CodePoint = _315___mcc_h25
			_ = _318_cf19
			var _319_cf18 _dafny.Int = _314___mcc_h24
			_ = _319_cf18
			return _pat_let_tv10
		} else if _source6.Is_DC11() {
			var _320___mcc_h27 _dafny.Set = _source6.Get_().(D4_DC11).Cf17
			_ = _320___mcc_h27
			var _321_cf17 _dafny.Set = _320___mcc_h27
			_ = _321_cf17
			return Companion_D0_.Create_DC2_(_pat_let_tv11, (_pat_let_tv12).F15(), false)
		} else {
			var _322___mcc_h28 D4 = _source6.Get_().(D4_DC13).Cf21
			_ = _322___mcc_h28
			var _323_cf21 D4 = _322___mcc_h28
			_ = _323_cf21
			return _pat_let_tv13
		}
	}(Companion_Default___.Fm12(_140_v6, (_263_v94).F15(), _148_globalState))
	_ = _source5
	if _source5.Is_DC0() {
		var _324___mcc_h16 _dafny.Int = _source5.Get_().(D0_DC0).Cf0
		_ = _324___mcc_h16
		var _325___mcc_h17 bool = _source5.Get_().(D0_DC0).Cf1
		_ = _325___mcc_h17
		var _326_cf1 bool = _325___mcc_h17
		_ = _326_cf1
		var _327_cf0 _dafny.Int = _324___mcc_h16
		_ = _327_cf0
		var _328_v127 *C1
		_ = _328_v127
		var _nw57 *C1 = New_C1_()
		_ = _nw57
		_nw57.Ctor__((_261_v92).F16(), (_263_v94).F15())
		_328_v127 = _nw57
		_140_v6 = _140_v6
		if (_328_v127).F14() {
			var _329_v128 T0
			_ = _329_v128
			var _nw58 *C1 = New_C1_()
			_ = _nw58
			_nw58.Ctor__(true, (_263_v94).F15())
			_329_v128 = _nw58
			var _330_v129 _dafny.MultiSet
			_ = _330_v129
			_330_v129 = _dafny.MultiSetOf((_263_v94).F15(), (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))
			var _331_v130 D4
			_ = _331_v130
			_331_v130 = Companion_D4_.Create_DC12_((_263_v94).F15(), _265_v96, _330_v129)
			var _332_v131 _dafny.Map
			_ = _332_v131
			_332_v131 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_138_v4, _331_v130)
			var _333_v132 _dafny.Map
			_ = _333_v132
			_333_v132 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_329_v128, (_332_v131).Cardinality())
			_326_cf1 = (_333_v132).Equals(_333_v132)
			var _334_v133 _dafny.Map
			_ = _334_v133
			_334_v133 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_328_v127).F14(), _143_v9)
			var _335_v134 _dafny.Map
			_ = _335_v134
			_335_v134 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_334_v133, _147_v13)
			_147_v13 = (func() _dafny.Array {
				if (_335_v134).Contains((_334_v133).Update((_263_v94).F14(), _143_v9)) {
					return (_335_v134).Get((_334_v133).Update((_263_v94).F14(), _143_v9)).(_dafny.Array)
				}
				return (func() _dafny.Array {
					if !(Companion_Default___.Fm6((_328_v127).F14(), (_263_v94).F15(), _148_globalState)) {
						return _147_v13
					}
					return _147_v13
				})()
			})()
			var _index58 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index58
			(_147_v13).ArraySet1(((_dafny.Zero).Minus((_263_v94).F15())).Times((_263_v94).F15()), (_index58).Int())
			var _336_v135 _dafny.Array
			_ = _336_v135
			var _nw59 _dafny.Array = _dafny.NewArrayWithValue(Companion_D3_.Default(), _dafny.IntOfInt64(28))
			_ = _nw59
			_336_v135 = _nw59
			var _337_v136 D3
			_ = _337_v136
			_337_v136 = Companion_D3_.Create_DC8_(_263_v94)
			var _index59 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_336_v135), 0))
			_ = _index59
			(_336_v135).ArraySet1(_337_v136, (_index59).Int())
			var _index60 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(844), _dafny.ArrayLen((_336_v135), 0))
			_ = _index60
			(_336_v135).ArraySet1(Companion_D3_.Create_DC8_(_328_v127), (_index60).Int())
			(_148_globalState).F4 = _dafny.CodePoint('w')
		} else {
			var _index61 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
			_ = _index61
			(_147_v13).ArraySet1((_263_v94).Fm1(_144_v10, (_328_v127).F15(), (_263_v94).F15(), _148_globalState), (_index61).Int())
			(_263_v94).M1(_148_globalState)
			_140_v6 = _326_cf1
			var _index62 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_258_v91), 0))
			_ = _index62
			(_258_v91).ArraySet1(!(true), (_index62).Int())
			var _338_v137 _dafny.MultiSet
			_ = _338_v137
			_338_v137 = _dafny.MultiSetOf((_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0), (_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0), _261_v92, _261_v92, (_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0))
			var _index63 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_258_v91), 0))
			_ = _index63
			(_258_v91).ArraySet1((_338_v137).IsDisjointFrom(_338_v137), (_index63).Int())
			var _339_v138 _dafny.Set
			_ = _339_v138
			_339_v138 = _dafny.SetOf((_258_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(321), _dafny.ArrayLen((_258_v91), 0))).Int()).(bool), (_261_v92).F16())
			_140_v6 = Companion_Default___.Fm4((_263_v94).F14(), _339_v138, (_138_v4).Select((Companion_Default___.SafeIndex((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), _dafny.IntOfUint32((_138_v4).Cardinality()))).Uint32()).(_dafny.Int), _148_globalState)
		}
		var _340_v139 _dafny.Map
		_ = _340_v139
		_340_v139 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v7, (func() *C0 {
			if (_263_v94).F14() {
				return (_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0)
			}
			return _261_v92
		})())
		var _341_v140 _dafny.Sequence
		_ = _341_v140
		_341_v140 = _dafny.SeqOf(_dafny.Companion_Sequence_.Concatenate(_144_v10, _144_v10), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(854))).Uint32(), func(coer7 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg7 _dafny.Int) interface{} {
				return coer7(arg7)
			}
		}((func(_342_v96 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_343_i7 _dafny.Int) _dafny.CodePoint {
				return _342_v96
			}
		})(_265_v96))), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(240))).Uint32(), func(coer8 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg8 _dafny.Int) interface{} {
				return coer8(arg8)
			}
		}((func(_344_v96 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
			return func(_345_i8 _dafny.Int) _dafny.CodePoint {
				return _344_v96
			}
		})(_265_v96))))
		var _346_v141 _dafny.Set
		_ = _346_v141
		_346_v141 = _dafny.SetOf(_140_v6, _326_cf1, _326_cf1, (_261_v92).F16())
		var _347_v142 _dafny.Map
		_ = _347_v142
		_347_v142 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_261_v92).F16(), (_263_v94).F15())
		var _348_v143 _dafny.Set
		_ = _348_v143
		_348_v143 = _dafny.SetOf(_346_v141, _346_v141, _346_v141, Companion_Default___.Fm9(_265_v96, _327_cf0, (_263_v94).F15(), (_347_v142).Update(_140_v6, _327_cf0), _148_globalState), _346_v141)
		var _rhs64 bool = (_150_v15).IsSubsetOf(_dafny.SetOf((_263_v94).F15()))
		_ = _rhs64
		var _rhs65 bool = Companion_Default___.Fm6((false) && ((_328_v127).F14()), _327_cf0, _148_globalState)
		_ = _rhs65
		var _rhs66 _dafny.Map = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_141_v7, (_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0))
		_ = _rhs66
		var _rhs67 *C0 = (_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0)
		_ = _rhs67
		var _rhs68 _dafny.Sequence = Companion_Default___.Fm13(Companion_D0_.Create_DC1_(_348_v143, true), (_143_v9).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(-196), _dafny.IntOfUint32((_143_v9).Cardinality()))).Uint32()).(bool), _148_globalState)
		_ = _rhs68
		_326_cf1 = _rhs64
		_140_v6 = _rhs65
		_340_v139 = _rhs66
		_261_v92 = _rhs67
		_341_v140 = _rhs68
	} else if _source5.Is_DC1() {
		var _349___mcc_h18 _dafny.Set = _source5.Get_().(D0_DC1).Cf2
		_ = _349___mcc_h18
		var _350___mcc_h19 bool = _source5.Get_().(D0_DC1).Cf3
		_ = _350___mcc_h19
		var _351_cf3 bool = _350___mcc_h19
		_ = _351_cf3
		var _352_cf2 _dafny.Set = _349___mcc_h18
		_ = _352_cf2
		var _353_v144 _dafny.Array
		_ = _353_v144
		var _nw60 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(5))
		_ = _nw60
		_353_v144 = _nw60
		var _index64 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_353_v144), 0))
		_ = _index64
		(_353_v144).ArraySet1(_144_v10, (_index64).Int())
		var _index65 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(140), _dafny.ArrayLen((_353_v144), 0))
		_ = _index65
		(_353_v144).ArraySet1(_144_v10, (_index65).Int())
		var _index66 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
		_ = _index66
		(_147_v13).ArraySet1((_263_v94).F15(), (_index66).Int())
		(_263_v94).M1(_148_globalState)
		var _354_v145 _dafny.Map
		_ = _354_v145
		_354_v145 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_258_v91, _dafny.IntOfUint32((_144_v10).Cardinality()))
		_354_v145 = (_354_v145).Update(_258_v91, (_263_v94).F15())
	} else if _source5.Is_DC2() {
		var _355___mcc_h20 _dafny.Set = _source5.Get_().(D0_DC2).Cf4
		_ = _355___mcc_h20
		var _356___mcc_h21 _dafny.Int = _source5.Get_().(D0_DC2).Cf5
		_ = _356___mcc_h21
		var _357___mcc_h22 bool = _source5.Get_().(D0_DC2).Cf6
		_ = _357___mcc_h22
		var _358_cf6 bool = _357___mcc_h22
		_ = _358_cf6
		var _359_cf5 _dafny.Int = _356___mcc_h21
		_ = _359_cf5
		var _360_cf4 _dafny.Set = _355___mcc_h20
		_ = _360_cf4
		_140_v6 = (_263_v94).F14()
		_140_v6 = false
		var _361_v146 *C1
		_ = _361_v146
		var _nw61 *C1 = New_C1_()
		_ = _nw61
		_nw61.Ctor__(!(((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Cmp(_359_cf5) < 0), Companion_Default___.SafeDivisionInt(_134_v0, (_263_v94).F15()))
		_361_v146 = _nw61
		var _362_v147 bool
		_ = _362_v147
		var _363_v148 _dafny.Array
		_ = _363_v148
		var _364_v149 _dafny.Int
		_ = _364_v149
		var _365_v150 _dafny.Int
		_ = _365_v150
		var _out28 bool
		_ = _out28
		var _out29 _dafny.Array
		_ = _out29
		var _out30 _dafny.Int
		_ = _out30
		var _out31 _dafny.Int
		_ = _out31
		_out28, _out29, _out30, _out31 = Companion_Default___.M0((func() _dafny.Int {
			if (_263_v94).F14() {
				return _134_v0
			}
			return _dafny.IntOfInt64(524)
		})(), _143_v9, _148_globalState)
		_362_v147 = _out28
		_363_v148 = _out29
		_364_v149 = _out30
		_365_v150 = _out31
	} else {
		var _366___mcc_h23 _dafny.CodePoint = _source5.Get_().(D0_DC3).Cf7
		_ = _366___mcc_h23
		var _367_cf7 _dafny.CodePoint = _366___mcc_h23
		_ = _367_cf7
		_263_v94 = _263_v94
		var _index67 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_258_v91), 0))
		_ = _index67
		(_258_v91).ArraySet1(true, (_index67).Int())
		var _368_v151 _dafny.Set
		_ = _368_v151
		_368_v151 = _dafny.SetOf(_367_cf7)
		var _369_v153 _dafny.Sequence
		_ = _369_v153
		_369_v153 = _dafny.SeqOf(_368_v151, _368_v151)
		var _370_v155 _dafny.Array
		_ = _370_v155
		var _nwElement0_10 _dafny.Set = _368_v151
		_ = _nwElement0_10
		var _nw62 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_10, nil, _dafny.IntOfInt64(20))
		_ = _nw62
		(_nw62).ArraySet1(_nwElement0_10, 0)
		(_nw62).ArraySet1(_368_v151, 1)
		(_nw62).ArraySet1(_368_v151, 2)
		(_nw62).ArraySet1(_368_v151, 3)
		(_nw62).ArraySet1(_dafny.SetOf(_367_cf7), 4)
		(_nw62).ArraySet1(Companion_Default___.Fm14(_dafny.IntOfInt64(456), _367_cf7, (_263_v94).F15(), (_263_v94).F15(), _148_globalState), 5)
		(_nw62).ArraySet1(_368_v151, 6)
		(_nw62).ArraySet1(func() _dafny.Set {
			var _coll4 = _dafny.NewBuilder()
			_ = _coll4
			for _iter4 := _dafny.Iterate((func() _dafny.Map {
				var _coll5 = _dafny.NewMapBuilder()
				_ = _coll5
				for _iter5 := _dafny.Iterate(((_369_v153).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_369_v153).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
					_compr_5, _ok5 := _iter5()
					if !_ok5 {
						break
					}
					var _371_v152 _dafny.CodePoint
					_371_v152 = interface{}(_compr_5).(_dafny.CodePoint)
					if ((_369_v153).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_369_v153).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_371_v152) {
						_coll5.Add(_371_v152, _367_cf7)
					}
				}
				return _coll5.ToMap()
			}()).Keys().Elements()); ; {
				_compr_4, _ok4 := _iter4()
				if !_ok4 {
					break
				}
				var _372_v154 _dafny.CodePoint
				_372_v154 = interface{}(_compr_4).(_dafny.CodePoint)
				if (func() _dafny.Map {
					var _coll6 = _dafny.NewMapBuilder()
					_ = _coll6
					for _iter6 := _dafny.Iterate(((_369_v153).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_369_v153).Cardinality()))).Uint32()).(_dafny.Set)).Elements()); ; {
						_compr_6, _ok6 := _iter6()
						if !_ok6 {
							break
						}
						var _371_v152 _dafny.CodePoint
						_371_v152 = interface{}(_compr_6).(_dafny.CodePoint)
						if ((_369_v153).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_369_v153).Cardinality()))).Uint32()).(_dafny.Set)).Contains(_371_v152) {
							_coll6.Add(_371_v152, _367_cf7)
						}
					}
					return _coll6.ToMap()
				}()).Contains(_372_v154) {
					_coll4.Add(_372_v154)
				}
			}
			return _coll4.ToSet()
		}(), 7)
		(_nw62).ArraySet1(_368_v151, 8)
		(_nw62).ArraySet1(_dafny.SetOf(Companion_Default___.Fm15(_dafny.CodePoint('k'), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v94).F15(), _134_v0), _148_globalState)), 9)
		(_nw62).ArraySet1((func() _dafny.Set {
			if (_261_v92).F16() {
				return _368_v151
			}
			return _368_v151
		})(), 10)
		(_nw62).ArraySet1((func() _dafny.Set {
			if !(false) {
				return _368_v151
			}
			return _368_v151
		})(), 11)
		(_nw62).ArraySet1(_368_v151, 12)
		(_nw62).ArraySet1(_368_v151, 13)
		(_nw62).ArraySet1(_368_v151, 14)
		(_nw62).ArraySet1((_368_v151).Union(_368_v151), 15)
		(_nw62).ArraySet1(_368_v151, 16)
		(_nw62).ArraySet1(_dafny.SetOf(_265_v96, _367_cf7), 17)
		(_nw62).ArraySet1(Companion_Default___.Fm14(_134_v0, _367_cf7, (_141_v7).Cardinality(), _dafny.IntOfInt64(42), _148_globalState), 18)
		(_nw62).ArraySet1((_368_v151).Union(_368_v151), 19)
		_370_v155 = _nw62
		var _373_v156 _dafny.Map
		_ = _373_v156
		_373_v156 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_143_v9, true)
		var _374_v157 _dafny.Sequence
		_ = _374_v157
		_374_v157 = _dafny.SeqOf((_262_v93).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(538), _dafny.ArrayLen((_262_v93), 0))).Int()).(*C0), _261_v92)
		var _375_v158 T0
		_ = _375_v158
		var _nw63 *C1 = New_C1_()
		_ = _nw63
		_nw63.Ctor__(Companion_Default___.Fm6((_263_v94).F14(), (_263_v94).F15(), _148_globalState), _dafny.IntOfUint32((_dafny.SeqOf(_dafny.IntOfInt64(819), (_373_v156).Cardinality(), _dafny.IntOfUint32((_374_v157).Cardinality()), Companion_Default___.Fm0(_145_v11, !(false), (_263_v94).F14(), _148_globalState))).Cardinality()))
		_375_v158 = _nw63
		var _376_v159 _dafny.Map
		_ = _376_v159
		_376_v159 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() *C1 {
			if (_263_v94).F14() {
				return _263_v94
			}
			return _263_v94
		})(), _375_v158)
		var _377_v160 _dafny.Sequence
		_ = _377_v160
		_377_v160 = _dafny.SeqOf(_370_v155)
		var _index68 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_258_v91), 0))
		_ = _index68
		var _rhs69 _dafny.Sequence = _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_263_v94).F14()), _143_v9)
		_ = _rhs69
		var _rhs70 bool = (_143_v9).Select((Companion_Default___.SafeIndex(_134_v0, _dafny.IntOfUint32((_143_v9).Cardinality()))).Uint32()).(bool)
		_ = _rhs70
		var _rhs71 bool = (_261_v92).F16()
		_ = _rhs71
		var _rhs72 _dafny.Int = (_376_v159).Cardinality()
		_ = _rhs72
		var _rhs73 _dafny.Array = (_377_v160).Select((Companion_Default___.SafeIndex(_134_v0, _dafny.IntOfUint32((_377_v160).Cardinality()))).Uint32()).(_dafny.Array)
		_ = _rhs73
		var _lhs38 _dafny.Array = _258_v91
		_ = _lhs38
		var _lhs39 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_258_v91), 0))
		_ = _lhs39
		_143_v9 = _rhs69
		(_lhs38).ArraySet1(_rhs70, (_lhs39).Int())
		_140_v6 = _rhs71
		_134_v0 = _rhs72
		_370_v155 = _rhs73
		var _378_v161 _dafny.Set
		_ = _378_v161
		_378_v161 = _dafny.SetOf((_261_v92).F16(), _140_v6, (_258_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_258_v91), 0))).Int()).(bool), (_261_v92).F16(), _140_v6)
		_140_v6 = Companion_Default___.Fm4((_140_v6) || ((_261_v92).F16()), _378_v161, Companion_Default___.SafeModuloInt(_134_v0, _dafny.IntOfInt64(-752)), _148_globalState)
		var _379_v162 _dafny.MultiSet
		_ = _379_v162
		_379_v162 = _dafny.MultiSetOf((_dafny.SetOf((_263_v94).F14(), !((_258_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_258_v91), 0))).Int()).(bool)), (_263_v94).F14(), (_261_v92).F16(), (_258_v91).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(615), _dafny.ArrayLen((_258_v91), 0))).Int()).(bool))).Cardinality(), _134_v0)
		var _380_v163 _dafny.Map
		_ = _380_v163
		_380_v163 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_367_cf7, (_263_v94).F14())
		var _381_v164 _dafny.Sequence
		_ = _381_v164
		_381_v164 = _dafny.SeqOf(_379_v162)
		var _382_v165 _dafny.Array
		_ = _382_v165
		var _nwElement0_11 _dafny.MultiSet = _379_v162
		_ = _nwElement0_11
		var _nw64 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_11, nil, _dafny.IntOfInt64(13))
		_ = _nw64
		(_nw64).ArraySet1(_nwElement0_11, 0)
		(_nw64).ArraySet1((Companion_Default___.Fm16(_140_v6, (func() bool {
			if (_149_v14).Contains((_380_v163).Cardinality()) {
				return (_149_v14).Get((_380_v163).Cardinality()).(bool)
			}
			return (func() bool {
				if (_145_v11).Contains((_261_v92).F16()) {
					return (_145_v11).Get((_261_v92).F16()).(bool)
				}
				return (_263_v94).F14()
			})()
		})(), _148_globalState)).Union(_379_v162), 1)
		(_nw64).ArraySet1((_379_v162).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(-260))), 2)
		(_nw64).ArraySet1(Companion_Default___.Fm16((_263_v94).F14(), (_263_v94).F14(), _148_globalState), 3)
		(_nw64).ArraySet1((_379_v162).Difference(_dafny.MultiSetOf(_dafny.IntOfInt64(-871))), 4)
		(_nw64).ArraySet1((_381_v164).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_381_v164).Cardinality()))).Uint32()).(_dafny.MultiSet), 5)
		(_nw64).ArraySet1(_dafny.MultiSetOf(Companion_Default___.Fm0(_145_v11, (_261_v92).F16(), (_261_v92).F16(), _148_globalState)), 6)
		(_nw64).ArraySet1((_379_v162).Difference(_379_v162), 7)
		(_nw64).ArraySet1((_dafny.MultiSetOf((_263_v94).F15())).Intersection(_379_v162), 8)
		(_nw64).ArraySet1((_381_v164).Select((Companion_Default___.SafeIndex(_134_v0, _dafny.IntOfUint32((_381_v164).Cardinality()))).Uint32()).(_dafny.MultiSet), 9)
		(_nw64).ArraySet1(_dafny.MultiSetOf(_134_v0), 10)
		(_nw64).ArraySet1((_379_v162).Union(_379_v162), 11)
		(_nw64).ArraySet1(_dafny.MultiSetOf(_dafny.IntOfUint32((_143_v9).Cardinality()), _134_v0), 12)
		_382_v165 = _nw64
		var _index69 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_382_v165), 0))
		_ = _index69
		(_382_v165).ArraySet1(_379_v162, (_index69).Int())
		var _index70 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(787), _dafny.ArrayLen((_382_v165), 0))
		_ = _index70
		(_382_v165).ArraySet1((_379_v162).Intersection(_379_v162), (_index70).Int())
	}
	var _383_i9 _dafny.Int
	_ = _383_i9
	_383_i9 = _dafny.Zero
	{
		for !(_140_v6) {
			{
				if (_383_i9).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L2
				}
				_383_i9 = (_383_i9).Plus(_dafny.One)
				var _384_v166 D3
				_ = _384_v166
				_384_v166 = Companion_D3_.Create_DC9_(_265_v96)
				var _385_v167 D3
				_ = _385_v167
				_385_v167 = Companion_D3_.Create_DC10_(_384_v166)
				var _386_v168 _dafny.Set
				_ = _386_v168
				_386_v168 = _dafny.SetOf(Companion_D3_.Create_DC10_(_384_v166), _385_v167)
				var _387_v169 _dafny.Sequence
				_ = _387_v169
				_387_v169 = _dafny.SeqOf(_386_v168, _386_v168, _386_v168)
				var _index71 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
				_ = _index71
				(_147_v13).ArraySet1(((_387_v169).Select((Companion_Default___.SafeIndex((func() _dafny.Int {
					if true {
						return (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)
					}
					return (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)
				})(), _dafny.IntOfUint32((_387_v169).Cardinality()))).Uint32()).(_dafny.Set)).Cardinality(), (_index71).Int())
				var _388_v170 _dafny.Array
				_ = _388_v170
				var _nw65 _dafny.Array = _dafny.NewArrayWithValue(_dafny.NewArrayWithValue(nil, _dafny.IntOf(0)), _dafny.IntOfInt64(19))
				_ = _nw65
				_388_v170 = _nw65
				var _389_v171 D5
				_ = _389_v171
				_389_v171 = Companion_D5_.Create_DC14_(_388_v170)
				_388_v170 = ((func() D5 {
					if _140_v6 {
						return _389_v171
					}
					return _389_v171
				})()).Dtor_cf22()
				var _390_v172 *C1
				_ = _390_v172
				var _nw66 *C1 = New_C1_()
				_ = _nw66
				_nw66.Ctor__(_140_v6, (_263_v94).F15())
				_390_v172 = _nw66
				var _391_v173 _dafny.Map
				_ = _391_v173
				_391_v173 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm17(_140_v6, _141_v7, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Update(_144_v10, (Companion_Default___.SafeIndex((_390_v172).F15(), _dafny.IntOfUint32((_144_v10).Cardinality()))).Uint32(), _265_v96)).Cardinality()), _144_v10, _148_globalState), _258_v91)
				var _392_v174 _dafny.Sequence
				_ = _392_v174
				_392_v174 = _dafny.SeqOf(_138_v4, _dafny.SeqOf((_263_v94).F15(), (_dafny.Zero).Minus(_dafny.IntOfUint32((_144_v10).Cardinality()))))
				_391_v173 = (_391_v173).Update(_392_v174, _258_v91)
				goto C2
			}
		C2:
		}
		goto L2
	}
L2:
	var _393_i10 _dafny.Int
	_ = _393_i10
	_393_i10 = _dafny.Zero
	{
		for (_140_v6) && (false) {
			{
				if (_393_i10).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L3
				}
				_393_i10 = (_393_i10).Plus(_dafny.One)
				var _394_v175 _dafny.Map
				_ = _394_v175
				_394_v175 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v94).F14(), _134_v0)
				_394_v175 = ((_394_v175).Merge((_394_v175).Update(true, (_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)))).Merge(_394_v175)
				(_261_v92).M1(_148_globalState)
				var _395_v176 *C0
				_ = _395_v176
				var _nw67 *C0 = New_C0_()
				_ = _nw67
				_nw67.Ctor__((_261_v92).F16())
				_395_v176 = _nw67
				var _396_v178 D0
				_ = _396_v178
				_396_v178 = Companion_D0_.Create_DC1_(func() _dafny.Set {
					var _coll7 = _dafny.NewBuilder()
					_ = _coll7
					for _iter7 := _dafny.Iterate((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-916))).Uint32(), func(coer9 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
						return func(arg9 _dafny.Int) interface{} {
							return coer9(arg9)
						}
					}((func(_397_v92 *C0) func(_dafny.Int) _dafny.Set {
						return func(_398_i11 _dafny.Int) _dafny.Set {
							return _dafny.SetOf((_397_v92).F16(), (_397_v92).F16(), !((_397_v92).F16()))
						}
					})(_261_v92)))).Elements()); ; {
						_compr_7, _ok7 := _iter7()
						if !_ok7 {
							break
						}
						var _399_v177 _dafny.Set
						_399_v177 = interface{}(_compr_7).(_dafny.Set)
						if _dafny.Companion_Sequence_.Contains(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-916))).Uint32(), func(coer10 func(_dafny.Int) _dafny.Set) func(_dafny.Int) interface{} {
							return func(arg10 _dafny.Int) interface{} {
								return coer10(arg10)
							}
						}((func(_400_v92 *C0) func(_dafny.Int) _dafny.Set {
							return func(_398_i11 _dafny.Int) _dafny.Set {
								return _dafny.SetOf((_400_v92).F16(), (_400_v92).F16(), !((_400_v92).F16()))
							}
						})(_261_v92))), _399_v177) {
							_coll7.Add(_399_v177)
						}
					}
					return _coll7.ToSet()
				}(), (_263_v94).F14())
				var _index72 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_258_v91), 0))
				_ = _index72
				(_258_v91).ArraySet1((_396_v178).Dtor_cf3(), (_index72).Int())
				var _index73 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(360), _dafny.ArrayLen((_258_v91), 0))
				_ = _index73
				(_258_v91).ArraySet1(!(_394_v175).Contains((_263_v94).F14()), (_index73).Int())
				goto C3
			}
		C3:
		}
		goto L3
	}
L3:
	var _401_v179 _dafny.Array
	_ = _401_v179
	var _nw68 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(28))
	_ = _nw68
	_401_v179 = _nw68
	var _index74 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_401_v179), 0))
	_ = _index74
	(_401_v179).ArraySet1(func() _dafny.Map {
		var _coll8 = _dafny.NewMapBuilder()
		_ = _coll8
		for _iter8 := _dafny.Iterate((_138_v4).Elements()); ; {
			_compr_8, _ok8 := _iter8()
			if !_ok8 {
				break
			}
			var _402_v180 _dafny.Int
			_402_v180 = interface{}(_compr_8).(_dafny.Int)
			if _dafny.Companion_Sequence_.Contains(_138_v4, _402_v180) {
				_coll8.Add((_402_v180).Plus((_263_v94).F15()), (_263_v94).F14())
			}
		}
		return _coll8.ToMap()
	}(), (_index74).Int())
	var _index75 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_401_v179), 0))
	_ = _index75
	(_401_v179).ArraySet1(_149_v14, (_index75).Int())
	(_261_v92).M1(_148_globalState)
	var _403_i12 _dafny.Int
	_ = _403_i12
	_403_i12 = _dafny.Zero
	{
		for ((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Cmp(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_144_v10, _dafny.Companion_Sequence_.Update(_144_v10, (Companion_Default___.SafeIndex((_dafny.Zero).Minus((_dafny.Zero).Minus((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int))), _dafny.IntOfUint32((_144_v10).Cardinality()))).Uint32(), _265_v96))).Cardinality())) != 0 {
			{
				if (_403_i12).Cmp(_dafny.IntOfInt64(100)) >= 0 {
					goto L4
				}
				_403_i12 = (_403_i12).Plus(_dafny.One)
				var _404_v181 T0
				_ = _404_v181
				var _nw69 *C1 = New_C1_()
				_ = _nw69
				_nw69.Ctor__(_140_v6, (_263_v94).F15())
				_404_v181 = _nw69
				var _405_v182 _dafny.Map
				_ = _405_v182
				_405_v182 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(784), _404_v181)
				var _406_v183 D3
				_ = _406_v183
				_406_v183 = Companion_D3_.Create_DC8_(_263_v94)
				var _407_v184 _dafny.Array
				_ = _407_v184
				var _nwElement0_12 D3 = _406_v183
				_ = _nwElement0_12
				var _nw70 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_12, nil, _dafny.IntOfInt64(3))
				_ = _nw70
				(_nw70).ArraySet1(_nwElement0_12, 0)
				(_nw70).ArraySet1(_406_v183, 1)
				(_nw70).ArraySet1(_406_v183, 2)
				_407_v184 = _nw70
				var _408_v185 _dafny.Map
				_ = _408_v185
				_408_v185 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((func() T0 {
					if (_405_v182).Contains(_134_v0) {
						return (_405_v182).Get(_134_v0).(T0)
					}
					return _404_v181
				})(), _407_v184)
				var _409_v186 _dafny.Map
				_ = _409_v186
				_409_v186 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_407_v184, _147_v13)
				_140_v6 = ((_409_v186).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_407_v184, _147_v13))).Contains((func() _dafny.Array {
					if (_408_v185).Contains(_404_v181) {
						return (_408_v185).Get(_404_v181).(_dafny.Array)
					}
					return _407_v184
				})())
				if (_263_v94).F14() {
					_140_v6 = (_263_v94).F14()
					_140_v6 = (_261_v92).F16()
					var _410_v187 *C1
					_ = _410_v187
					var _nw71 *C1 = New_C1_()
					_ = _nw71
					_nw71.Ctor__((_261_v92).F16(), (_263_v94).F15())
					_410_v187 = _nw71
					_140_v6 = Companion_Default___.Fm6((_263_v94).F14(), (_263_v94).F15(), _148_globalState)
					var _411_v188 _dafny.Set
					_ = _411_v188
					_411_v188 = _dafny.SetOf((_263_v94).F14())
					var _412_v189 _dafny.MultiSet
					_ = _412_v189
					_412_v189 = _dafny.MultiSetOf(_411_v188)
					var _index76 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_258_v91), 0))
					_ = _index76
					(_258_v91).ArraySet1((_412_v189).IsProperSubsetOf(_412_v189), (_index76).Int())
					var _index77 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(171), _dafny.ArrayLen((_258_v91), 0))
					_ = _index77
					(_258_v91).ArraySet1(_140_v6, (_index77).Int())
				} else {
					(_148_globalState).F4 = _265_v96
					_147_v13 = _147_v13
					var _413_v190 _dafny.Sequence
					_ = _413_v190
					_413_v190 = _dafny.SeqOf(_dafny.Companion_Sequence_.Update(_dafny.SeqOf((_263_v94).F14()), (Companion_Default___.SafeIndex((_138_v4).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_138_v4).Cardinality()))).Uint32()).(_dafny.Int), _dafny.IntOfUint32((_dafny.SeqOf((_263_v94).F14())).Cardinality()))).Uint32(), _140_v6), _143_v9, _dafny.Companion_Sequence_.Concatenate(_143_v9, _143_v9), _143_v9)
					_413_v190 = _dafny.Companion_Sequence_.Concatenate(_413_v190, _413_v190)
					var _index78 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_258_v91), 0))
					_ = _index78
					(_258_v91).ArraySet1((_263_v94).F14(), (_index78).Int())
					var _414_v191 _dafny.Set
					_ = _414_v191
					_414_v191 = _dafny.SetOf(_140_v6, (_263_v94).F14())
					var _index79 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(864), _dafny.ArrayLen((_258_v91), 0))
					_ = _index79
					(_258_v91).ArraySet1((func() bool {
						if _140_v6 {
							return ((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)).Cmp((_263_v94).F15()) > 0
						}
						return Companion_Default___.Fm4((_263_v94).F14(), _414_v191, _134_v0, _148_globalState)
					})(), (_index79).Int())
					_405_v182 = (_405_v182).Update((_134_v0).Plus((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int)), _404_v181)
				}
				if (_261_v92).F16() {
					var _415_v192 _dafny.Sequence
					_ = _415_v192
					_415_v192 = _dafny.SeqOf(_404_v181, _404_v181, _404_v181, _404_v181, _404_v181)
					var _416_v193 _dafny.Array
					_ = _416_v193
					var _nwElement0_13 _dafny.Sequence = _415_v192
					_ = _nwElement0_13
					var _nw72 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_13, nil, _dafny.One)
					_ = _nw72
					(_nw72).ArraySet1(_nwElement0_13, 0)
					_416_v193 = _nw72
					var _index80 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_416_v193), 0))
					_ = _index80
					(_416_v193).ArraySet1(_415_v192, (_index80).Int())
					var _index81 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(941), _dafny.ArrayLen((_416_v193), 0))
					_ = _index81
					(_416_v193).ArraySet1(_415_v192, (_index81).Int())
					_144_v10 = _dafny.Companion_Sequence_.Concatenate(_144_v10, _144_v10)
					var _index82 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
					_ = _index82
					(_147_v13).ArraySet1(((_263_v94).F15()).Times(((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("nadw")).Cardinality()), (_263_v94).F14())).Merge((_401_v179).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(534), _dafny.ArrayLen((_401_v179), 0))).Int()).(_dafny.Map))).Cardinality()), (_index82).Int())
					var _417_v194 _dafny.Array
					_ = _417_v194
					var _nwElement0_14 T0 = _404_v181
					_ = _nwElement0_14
					var _nw73 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_14, nil, _dafny.IntOfInt64(9))
					_ = _nw73
					(_nw73).ArraySet1(_nwElement0_14, 0)
					(_nw73).ArraySet1(_404_v181, 1)
					(_nw73).ArraySet1(_404_v181, 2)
					(_nw73).ArraySet1(_404_v181, 3)
					(_nw73).ArraySet1(_404_v181, 4)
					(_nw73).ArraySet1(_404_v181, 5)
					(_nw73).ArraySet1(_404_v181, 6)
					(_nw73).ArraySet1(_404_v181, 7)
					(_nw73).ArraySet1(_404_v181, 8)
					_417_v194 = _nw73
					_417_v194 = _417_v194
					var _418_v195 *C1
					_ = _418_v195
					var _nw74 *C1 = New_C1_()
					_ = _nw74
					_nw74.Ctor__((_261_v92).F16(), (func() _dafny.Int {
						if false {
							return (_263_v94).F15()
						}
						return (_dafny.Zero).Minus((_404_v181).Fm1(_144_v10, _dafny.IntOfInt64(-329), _134_v0, _148_globalState))
					})())
					_418_v195 = _nw74
				} else {
					var _419_v196 _dafny.Set
					_ = _419_v196
					_419_v196 = _dafny.SetOf(_404_v181, _404_v181)
					var _420_v197 _dafny.Set
					_ = _420_v197
					_420_v197 = _dafny.SetOf((_419_v196))
					var _421_v198 _dafny.Set
					_ = _421_v198
					_421_v198 = _dafny.SetOf(_404_v181)
					var _422_v199 _dafny.Sequence
					_ = _422_v199
					_422_v199 = _dafny.SeqOf(_420_v197)
					var _423_v200 _dafny.Array
					_ = _423_v200
					var _nwElement0_15 _dafny.Set = _420_v197
					_ = _nwElement0_15
					var _nw75 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_15, nil, _dafny.IntOfInt64(25))
					_ = _nw75
					(_nw75).ArraySet1(_nwElement0_15, 0)
					(_nw75).ArraySet1(_420_v197, 1)
					(_nw75).ArraySet1(_420_v197, 2)
					(_nw75).ArraySet1(_420_v197, 3)
					(_nw75).ArraySet1(_420_v197, 4)
					(_nw75).ArraySet1(_420_v197, 5)
					(_nw75).ArraySet1((_420_v197).Intersection(_420_v197), 6)
					(_nw75).ArraySet1(_420_v197, 7)
					(_nw75).ArraySet1(_420_v197, 8)
					(_nw75).ArraySet1((_420_v197).Intersection(_420_v197), 9)
					(_nw75).ArraySet1((_420_v197).Difference(_dafny.SetOf(_421_v198)), 10)
					(_nw75).ArraySet1(_420_v197, 11)
					(_nw75).ArraySet1((_420_v197).Difference(_420_v197), 12)
					(_nw75).ArraySet1(_420_v197, 13)
					(_nw75).ArraySet1(_420_v197, 14)
					(_nw75).ArraySet1(_420_v197, 15)
					(_nw75).ArraySet1((_420_v197).Difference(_420_v197), 16)
					(_nw75).ArraySet1(_420_v197, 17)
					(_nw75).ArraySet1(_420_v197, 18)
					(_nw75).ArraySet1(_420_v197, 19)
					(_nw75).ArraySet1(_420_v197, 20)
					(_nw75).ArraySet1(_420_v197, 21)
					(_nw75).ArraySet1(_420_v197, 22)
					(_nw75).ArraySet1((_422_v199).Select((Companion_Default___.SafeIndex((_263_v94).F15(), _dafny.IntOfUint32((_422_v199).Cardinality()))).Uint32()).(_dafny.Set), 23)
					(_nw75).ArraySet1(_420_v197, 24)
					_423_v200 = _nw75
					var _index83 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_423_v200), 0))
					_ = _index83
					(_423_v200).ArraySet1(_420_v197, (_index83).Int())
					var _index84 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(982), _dafny.ArrayLen((_423_v200), 0))
					_ = _index84
					(_423_v200).ArraySet1(_420_v197, (_index84).Int())
					_140_v6 = (_261_v92).F16()
					(_261_v92).M1(_148_globalState)
					var _424_v201 D0
					_ = _424_v201
					_424_v201 = Companion_D0_.Create_DC0_((_263_v94).F15(), (_263_v94).F14())
					var _425_v202 _dafny.Map
					_ = _425_v202
					_425_v202 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v94).F14(), _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_263_v94).F15(), (_261_v92).F16()))
					_424_v201 = Companion_D0_.Create_DC0_((((_425_v202).Update(false, _149_v14)).Merge(_425_v202)).Cardinality(), !((_261_v92).F16()) || (_140_v6))
					var _index85 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))
					_ = _index85
					(_147_v13).ArraySet1((_147_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(749), _dafny.ArrayLen((_147_v13), 0))).Int()).(_dafny.Int), (_index85).Int())
				}
				var _426_v203 _dafny.Set
				_ = _426_v203
				_426_v203 = _dafny.SetOf(_258_v91, _258_v91)
				_426_v203 = (_426_v203).Intersection(_426_v203)
				goto C4
			}
		C4:
		}
		goto L4
	}
L4:
	var _427_v204 D1
	_ = _427_v204
	_427_v204 = Companion_D1_.Create_DC5_(_140_v6)
	var _pat_let_tv14 = _145_v11
	_ = _pat_let_tv14
	var _pat_let_tv15 = _263_v94
	_ = _pat_let_tv15
	var _pat_let_tv16 = _145_v11
	_ = _pat_let_tv16
	var _pat_let_tv17 = _263_v94
	_ = _pat_let_tv17
	var _pat_let_tv18 = _263_v94
	_ = _pat_let_tv18
	_140_v6 = func(_source7 D1) bool {
		if _source7.Is_DC5() {
			var _428___mcc_h29 bool = _source7.Get_().(D1_DC5).Cf9
			_ = _428___mcc_h29
			var _429_cf9 bool = _428___mcc_h29
			_ = _429_cf9
			return _429_cf9
		} else {
			var _430___mcc_h30 _dafny.Sequence = _source7.Get_().(D1_DC4).Cf8
			_ = _430___mcc_h30
			var _431_cf8 _dafny.Sequence = _430___mcc_h30
			_ = _431_cf8
			if (_pat_let_tv14).Contains((_pat_let_tv15).F14()) {
				return (_pat_let_tv16).Get((_pat_let_tv17).F14()).(bool)
			} else {
				return (_pat_let_tv18).F14()
			}
		}
	}(_427_v204)
	_dafny.Print(_134_v0)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_135_v1).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-250)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_136_v2).ArrayGet1CodePoint((_dafny.IntOfInt64(27)).Int()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_137_v3).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_138_v4, _dafny.SeqOf(_dafny.IntOfInt64(-250), _dafny.IntOfInt64(-250), _dafny.IntOfInt64(-250))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_139_v5).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-250)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_140_v6)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_141_v7).Equals(_dafny.MultiSetOf(false, false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_142_v8).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal(_143_v9, _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_144_v10.VerbatimString(false))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_145_v11).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(false, true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_146_v12).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-250), _dafny.IntOfInt64(2)).UpdateUnsafe(_dafny.Zero, _dafny.Zero)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_147_v13).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F0())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F1())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_148_globalState).F2()).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(_dafny.IntOfInt64(-250)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F3())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_148_globalState.F4)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F5).Cardinality())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F6).Equals(_dafny.SetOf(_dafny.SeqOf(_dafny.IntOfInt64(-250)))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F7())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_148_globalState).F8()).Equals(_dafny.MultiSetOf(_dafny.MultiSetOf(false, false))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F9())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F10())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.Companion_Sequence_.Equal((_148_globalState).F11(), _dafny.SeqOf(true)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.Zero).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.One).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState.F12).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(_dafny.Int))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_148_globalState).F13())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_149_v14).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(84), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_150_v15).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_151_v16).Dtor_cf4()).Equals(_dafny.SetOf(_dafny.One)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v16).Dtor_cf5())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_151_v16).Dtor_cf6())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_249_i2)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.Zero).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.One).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.IntOfInt64(3)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.IntOfInt64(4)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.IntOfInt64(5)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.IntOfInt64(6)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_258_v91).ArrayGet1((_dafny.IntOfInt64(7)).Int()).(bool))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_261_v92).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_262_v93).ArrayGet1((_dafny.Zero).Int()).(*C0)).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_262_v93).ArrayGet1((_dafny.One).Int()).(*C0)).F16())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v94).F14())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_263_v94).F15())
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_dafny.IntOfUint32((_264_v95).Cardinality()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_265_v96)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_266_v97).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(2), _dafny.CodePoint('c'))))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_313_v126).Equals(_dafny.MultiSetOf()))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_383_i9)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_393_i10)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(((_401_v179).ArrayGet1((_dafny.IntOfInt64(2)).Int()).(_dafny.Map)).Equals(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(84), false)))
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print(_403_i12)
	_dafny.Print(_dafny.UnicodeSeqOfUtf8Bytes("\n").VerbatimString(false))
	_dafny.Print((_427_v204).Dtor_cf9())
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
	Cf0 _dafny.Int
	Cf1 bool
}

func (D0_DC0) isD0() {}

func (CompanionStruct_D0_) Create_DC0_(Cf0 _dafny.Int, Cf1 bool) D0 {
	return D0{D0_DC0{Cf0, Cf1}}
}

func (_this D0) Is_DC0() bool {
	_, ok := _this.Get_().(D0_DC0)
	return ok
}

type D0_DC1 struct {
	Cf2 _dafny.Set
	Cf3 bool
}

func (D0_DC1) isD0() {}

func (CompanionStruct_D0_) Create_DC1_(Cf2 _dafny.Set, Cf3 bool) D0 {
	return D0{D0_DC1{Cf2, Cf3}}
}

func (_this D0) Is_DC1() bool {
	_, ok := _this.Get_().(D0_DC1)
	return ok
}

type D0_DC2 struct {
	Cf4 _dafny.Set
	Cf5 _dafny.Int
	Cf6 bool
}

func (D0_DC2) isD0() {}

func (CompanionStruct_D0_) Create_DC2_(Cf4 _dafny.Set, Cf5 _dafny.Int, Cf6 bool) D0 {
	return D0{D0_DC2{Cf4, Cf5, Cf6}}
}

func (_this D0) Is_DC2() bool {
	_, ok := _this.Get_().(D0_DC2)
	return ok
}

type D0_DC3 struct {
	Cf7 _dafny.CodePoint
}

func (D0_DC3) isD0() {}

func (CompanionStruct_D0_) Create_DC3_(Cf7 _dafny.CodePoint) D0 {
	return D0{D0_DC3{Cf7}}
}

func (_this D0) Is_DC3() bool {
	_, ok := _this.Get_().(D0_DC3)
	return ok
}

func (CompanionStruct_D0_) Default() D0 {
	return Companion_D0_.Create_DC0_(_dafny.Zero, false)
}

func (_this D0) Dtor_cf0() _dafny.Int {
	return _this.Get_().(D0_DC0).Cf0
}

func (_this D0) Dtor_cf1() bool {
	return _this.Get_().(D0_DC0).Cf1
}

func (_this D0) Dtor_cf2() _dafny.Set {
	return _this.Get_().(D0_DC1).Cf2
}

func (_this D0) Dtor_cf3() bool {
	return _this.Get_().(D0_DC1).Cf3
}

func (_this D0) Dtor_cf4() _dafny.Set {
	return _this.Get_().(D0_DC2).Cf4
}

func (_this D0) Dtor_cf5() _dafny.Int {
	return _this.Get_().(D0_DC2).Cf5
}

func (_this D0) Dtor_cf6() bool {
	return _this.Get_().(D0_DC2).Cf6
}

func (_this D0) Dtor_cf7() _dafny.CodePoint {
	return _this.Get_().(D0_DC3).Cf7
}

func (_this D0) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D0_DC0:
		{
			return "D0.DC0" + "(" + _dafny.String(data.Cf0) + ", " + _dafny.String(data.Cf1) + ")"
		}
	case D0_DC1:
		{
			return "D0.DC1" + "(" + _dafny.String(data.Cf2) + ", " + _dafny.String(data.Cf3) + ")"
		}
	case D0_DC2:
		{
			return "D0.DC2" + "(" + _dafny.String(data.Cf4) + ", " + _dafny.String(data.Cf5) + ", " + _dafny.String(data.Cf6) + ")"
		}
	case D0_DC3:
		{
			return "D0.DC3" + "(" + _dafny.String(data.Cf7) + ")"
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
			data2, ok := other.Get_().(D0_DC0)
			return ok && data1.Cf0.Cmp(data2.Cf0) == 0 && data1.Cf1 == data2.Cf1
		}
	case D0_DC1:
		{
			data2, ok := other.Get_().(D0_DC1)
			return ok && data1.Cf2.Equals(data2.Cf2) && data1.Cf3 == data2.Cf3
		}
	case D0_DC2:
		{
			data2, ok := other.Get_().(D0_DC2)
			return ok && data1.Cf4.Equals(data2.Cf4) && data1.Cf5.Cmp(data2.Cf5) == 0 && data1.Cf6 == data2.Cf6
		}
	case D0_DC3:
		{
			data2, ok := other.Get_().(D0_DC3)
			return ok && data1.Cf7 == data2.Cf7
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
	Cf9 bool
}

func (D1_DC5) isD1() {}

func (CompanionStruct_D1_) Create_DC5_(Cf9 bool) D1 {
	return D1{D1_DC5{Cf9}}
}

func (_this D1) Is_DC5() bool {
	_, ok := _this.Get_().(D1_DC5)
	return ok
}

type D1_DC4 struct {
	Cf8 _dafny.Sequence
}

func (D1_DC4) isD1() {}

func (CompanionStruct_D1_) Create_DC4_(Cf8 _dafny.Sequence) D1 {
	return D1{D1_DC4{Cf8}}
}

func (_this D1) Is_DC4() bool {
	_, ok := _this.Get_().(D1_DC4)
	return ok
}

func (CompanionStruct_D1_) Default() D1 {
	return Companion_D1_.Create_DC5_(false)
}

func (_this D1) Dtor_cf9() bool {
	return _this.Get_().(D1_DC5).Cf9
}

func (_this D1) Dtor_cf8() _dafny.Sequence {
	return _this.Get_().(D1_DC4).Cf8
}

func (_this D1) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D1_DC5:
		{
			return "D1.DC5" + "(" + _dafny.String(data.Cf9) + ")"
		}
	case D1_DC4:
		{
			return "D1.DC4" + "(" + data.Cf8.VerbatimString(true) + ")"
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
			data2, ok := other.Get_().(D1_DC5)
			return ok && data1.Cf9 == data2.Cf9
		}
	case D1_DC4:
		{
			data2, ok := other.Get_().(D1_DC4)
			return ok && data1.Cf8.Equals(data2.Cf8)
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
	Cf11 bool
	Cf12 bool
	Cf13 bool
}

func (D2_DC7) isD2() {}

func (CompanionStruct_D2_) Create_DC7_(Cf11 bool, Cf12 bool, Cf13 bool) D2 {
	return D2{D2_DC7{Cf11, Cf12, Cf13}}
}

func (_this D2) Is_DC7() bool {
	_, ok := _this.Get_().(D2_DC7)
	return ok
}

type D2_DC6 struct {
	Cf10 _dafny.Sequence
}

func (D2_DC6) isD2() {}

func (CompanionStruct_D2_) Create_DC6_(Cf10 _dafny.Sequence) D2 {
	return D2{D2_DC6{Cf10}}
}

func (_this D2) Is_DC6() bool {
	_, ok := _this.Get_().(D2_DC6)
	return ok
}

func (CompanionStruct_D2_) Default() D2 {
	return Companion_D2_.Create_DC7_(false, false, false)
}

func (_this D2) Dtor_cf11() bool {
	return _this.Get_().(D2_DC7).Cf11
}

func (_this D2) Dtor_cf12() bool {
	return _this.Get_().(D2_DC7).Cf12
}

func (_this D2) Dtor_cf13() bool {
	return _this.Get_().(D2_DC7).Cf13
}

func (_this D2) Dtor_cf10() _dafny.Sequence {
	return _this.Get_().(D2_DC6).Cf10
}

func (_this D2) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D2_DC7:
		{
			return "D2.DC7" + "(" + _dafny.String(data.Cf11) + ", " + _dafny.String(data.Cf12) + ", " + _dafny.String(data.Cf13) + ")"
		}
	case D2_DC6:
		{
			return "D2.DC6" + "(" + _dafny.String(data.Cf10) + ")"
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
			return ok && data1.Cf11 == data2.Cf11 && data1.Cf12 == data2.Cf12 && data1.Cf13 == data2.Cf13
		}
	case D2_DC6:
		{
			data2, ok := other.Get_().(D2_DC6)
			return ok && data1.Cf10.Equals(data2.Cf10)
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

type D3_DC9 struct {
	Cf15 _dafny.CodePoint
}

func (D3_DC9) isD3() {}

func (CompanionStruct_D3_) Create_DC9_(Cf15 _dafny.CodePoint) D3 {
	return D3{D3_DC9{Cf15}}
}

func (_this D3) Is_DC9() bool {
	_, ok := _this.Get_().(D3_DC9)
	return ok
}

type D3_DC8 struct {
	Cf14 *C1
}

func (D3_DC8) isD3() {}

func (CompanionStruct_D3_) Create_DC8_(Cf14 *C1) D3 {
	return D3{D3_DC8{Cf14}}
}

func (_this D3) Is_DC8() bool {
	_, ok := _this.Get_().(D3_DC8)
	return ok
}

type D3_DC10 struct {
	Cf16 D3
}

func (D3_DC10) isD3() {}

func (CompanionStruct_D3_) Create_DC10_(Cf16 D3) D3 {
	return D3{D3_DC10{Cf16}}
}

func (_this D3) Is_DC10() bool {
	_, ok := _this.Get_().(D3_DC10)
	return ok
}

func (CompanionStruct_D3_) Default() D3 {
	return Companion_D3_.Create_DC9_(_dafny.CodePoint('D'))
}

func (_this D3) Dtor_cf15() _dafny.CodePoint {
	return _this.Get_().(D3_DC9).Cf15
}

func (_this D3) Dtor_cf14() *C1 {
	return _this.Get_().(D3_DC8).Cf14
}

func (_this D3) Dtor_cf16() D3 {
	return _this.Get_().(D3_DC10).Cf16
}

func (_this D3) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D3_DC9:
		{
			return "D3.DC9" + "(" + _dafny.String(data.Cf15) + ")"
		}
	case D3_DC8:
		{
			return "D3.DC8" + "(" + _dafny.String(data.Cf14) + ")"
		}
	case D3_DC10:
		{
			return "D3.DC10" + "(" + _dafny.String(data.Cf16) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D3) Equals(other D3) bool {
	switch data1 := _this.Get_().(type) {
	case D3_DC9:
		{
			data2, ok := other.Get_().(D3_DC9)
			return ok && data1.Cf15 == data2.Cf15
		}
	case D3_DC8:
		{
			data2, ok := other.Get_().(D3_DC8)
			return ok && data1.Cf14 == data2.Cf14
		}
	case D3_DC10:
		{
			data2, ok := other.Get_().(D3_DC10)
			return ok && data1.Cf16.Equals(data2.Cf16)
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

type D4_DC12 struct {
	Cf18 _dafny.Int
	Cf19 _dafny.CodePoint
	Cf20 _dafny.MultiSet
}

func (D4_DC12) isD4() {}

func (CompanionStruct_D4_) Create_DC12_(Cf18 _dafny.Int, Cf19 _dafny.CodePoint, Cf20 _dafny.MultiSet) D4 {
	return D4{D4_DC12{Cf18, Cf19, Cf20}}
}

func (_this D4) Is_DC12() bool {
	_, ok := _this.Get_().(D4_DC12)
	return ok
}

type D4_DC11 struct {
	Cf17 _dafny.Set
}

func (D4_DC11) isD4() {}

func (CompanionStruct_D4_) Create_DC11_(Cf17 _dafny.Set) D4 {
	return D4{D4_DC11{Cf17}}
}

func (_this D4) Is_DC11() bool {
	_, ok := _this.Get_().(D4_DC11)
	return ok
}

type D4_DC13 struct {
	Cf21 D4
}

func (D4_DC13) isD4() {}

func (CompanionStruct_D4_) Create_DC13_(Cf21 D4) D4 {
	return D4{D4_DC13{Cf21}}
}

func (_this D4) Is_DC13() bool {
	_, ok := _this.Get_().(D4_DC13)
	return ok
}

func (CompanionStruct_D4_) Default() D4 {
	return Companion_D4_.Create_DC12_(_dafny.Zero, _dafny.CodePoint('D'), _dafny.EmptyMultiSet)
}

func (_this D4) Dtor_cf18() _dafny.Int {
	return _this.Get_().(D4_DC12).Cf18
}

func (_this D4) Dtor_cf19() _dafny.CodePoint {
	return _this.Get_().(D4_DC12).Cf19
}

func (_this D4) Dtor_cf20() _dafny.MultiSet {
	return _this.Get_().(D4_DC12).Cf20
}

func (_this D4) Dtor_cf17() _dafny.Set {
	return _this.Get_().(D4_DC11).Cf17
}

func (_this D4) Dtor_cf21() D4 {
	return _this.Get_().(D4_DC13).Cf21
}

func (_this D4) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D4_DC12:
		{
			return "D4.DC12" + "(" + _dafny.String(data.Cf18) + ", " + _dafny.String(data.Cf19) + ", " + _dafny.String(data.Cf20) + ")"
		}
	case D4_DC11:
		{
			return "D4.DC11" + "(" + _dafny.String(data.Cf17) + ")"
		}
	case D4_DC13:
		{
			return "D4.DC13" + "(" + _dafny.String(data.Cf21) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this D4) Equals(other D4) bool {
	switch data1 := _this.Get_().(type) {
	case D4_DC12:
		{
			data2, ok := other.Get_().(D4_DC12)
			return ok && data1.Cf18.Cmp(data2.Cf18) == 0 && data1.Cf19 == data2.Cf19 && data1.Cf20.Equals(data2.Cf20)
		}
	case D4_DC11:
		{
			data2, ok := other.Get_().(D4_DC11)
			return ok && data1.Cf17.Equals(data2.Cf17)
		}
	case D4_DC13:
		{
			data2, ok := other.Get_().(D4_DC13)
			return ok && data1.Cf21.Equals(data2.Cf21)
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
	Cf23 _dafny.Set
	Cf24 _dafny.Sequence
}

func (D5_DC15) isD5() {}

func (CompanionStruct_D5_) Create_DC15_(Cf23 _dafny.Set, Cf24 _dafny.Sequence) D5 {
	return D5{D5_DC15{Cf23, Cf24}}
}

func (_this D5) Is_DC15() bool {
	_, ok := _this.Get_().(D5_DC15)
	return ok
}

type D5_DC16 struct {
}

func (D5_DC16) isD5() {}

func (CompanionStruct_D5_) Create_DC16_() D5 {
	return D5{D5_DC16{}}
}

func (_this D5) Is_DC16() bool {
	_, ok := _this.Get_().(D5_DC16)
	return ok
}

type D5_DC14 struct {
	Cf22 _dafny.Array
}

func (D5_DC14) isD5() {}

func (CompanionStruct_D5_) Create_DC14_(Cf22 _dafny.Array) D5 {
	return D5{D5_DC14{Cf22}}
}

func (_this D5) Is_DC14() bool {
	_, ok := _this.Get_().(D5_DC14)
	return ok
}

func (CompanionStruct_D5_) Default() D5 {
	return Companion_D5_.Create_DC15_(_dafny.EmptySet, _dafny.EmptySeq)
}

func (_this D5) Dtor_cf23() _dafny.Set {
	return _this.Get_().(D5_DC15).Cf23
}

func (_this D5) Dtor_cf24() _dafny.Sequence {
	return _this.Get_().(D5_DC15).Cf24
}

func (_this D5) Dtor_cf22() _dafny.Array {
	return _this.Get_().(D5_DC14).Cf22
}

func (_this D5) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D5_DC15:
		{
			return "D5.DC15" + "(" + _dafny.String(data.Cf23) + ", " + _dafny.String(data.Cf24) + ")"
		}
	case D5_DC16:
		{
			return "D5.DC16"
		}
	case D5_DC14:
		{
			return "D5.DC14" + "(" + _dafny.String(data.Cf22) + ")"
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
			return ok && data1.Cf23.Equals(data2.Cf23) && data1.Cf24.Equals(data2.Cf24)
		}
	case D5_DC16:
		{
			_, ok := other.Get_().(D5_DC16)
			return ok
		}
	case D5_DC14:
		{
			data2, ok := other.Get_().(D5_DC14)
			return ok && data1.Cf22 == data2.Cf22
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
	Cf25 _dafny.Set
}

func (D6_DC17) isD6() {}

func (CompanionStruct_D6_) Create_DC17_(Cf25 _dafny.Set) D6 {
	return D6{D6_DC17{Cf25}}
}

func (_this D6) Is_DC17() bool {
	_, ok := _this.Get_().(D6_DC17)
	return ok
}

func (CompanionStruct_D6_) Default() _dafny.Set {
	return _dafny.EmptySet
}

func (_this D6) Dtor_cf25() _dafny.Set {
	return _this.Get_().(D6_DC17).Cf25
}

func (_this D6) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D6_DC17:
		{
			return "D6.DC17" + "(" + _dafny.String(data.Cf25) + ")"
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
			return ok && data1.Cf25.Equals(data2.Cf25)
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
	Cf27 bool
}

func (D7_DC19) isD7() {}

func (CompanionStruct_D7_) Create_DC19_(Cf27 bool) D7 {
	return D7{D7_DC19{Cf27}}
}

func (_this D7) Is_DC19() bool {
	_, ok := _this.Get_().(D7_DC19)
	return ok
}

type D7_DC18 struct {
	Cf26 _dafny.MultiSet
}

func (D7_DC18) isD7() {}

func (CompanionStruct_D7_) Create_DC18_(Cf26 _dafny.MultiSet) D7 {
	return D7{D7_DC18{Cf26}}
}

func (_this D7) Is_DC18() bool {
	_, ok := _this.Get_().(D7_DC18)
	return ok
}

func (CompanionStruct_D7_) Default() D7 {
	return Companion_D7_.Create_DC19_(false)
}

func (_this D7) Dtor_cf27() bool {
	return _this.Get_().(D7_DC19).Cf27
}

func (_this D7) Dtor_cf26() _dafny.MultiSet {
	return _this.Get_().(D7_DC18).Cf26
}

func (_this D7) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D7_DC19:
		{
			return "D7.DC19" + "(" + _dafny.String(data.Cf27) + ")"
		}
	case D7_DC18:
		{
			return "D7.DC18" + "(" + _dafny.String(data.Cf26) + ")"
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
			return ok && data1.Cf27 == data2.Cf27
		}
	case D7_DC18:
		{
			data2, ok := other.Get_().(D7_DC18)
			return ok && data1.Cf26.Equals(data2.Cf26)
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
	Cf29 _dafny.Sequence
	Cf30 _dafny.CodePoint
	Cf31 _dafny.Set
}

func (D8_DC21) isD8() {}

func (CompanionStruct_D8_) Create_DC21_(Cf29 _dafny.Sequence, Cf30 _dafny.CodePoint, Cf31 _dafny.Set) D8 {
	return D8{D8_DC21{Cf29, Cf30, Cf31}}
}

func (_this D8) Is_DC21() bool {
	_, ok := _this.Get_().(D8_DC21)
	return ok
}

type D8_DC22 struct {
	Cf32 bool
	Cf33 _dafny.Map
	Cf34 _dafny.Int
}

func (D8_DC22) isD8() {}

func (CompanionStruct_D8_) Create_DC22_(Cf32 bool, Cf33 _dafny.Map, Cf34 _dafny.Int) D8 {
	return D8{D8_DC22{Cf32, Cf33, Cf34}}
}

func (_this D8) Is_DC22() bool {
	_, ok := _this.Get_().(D8_DC22)
	return ok
}

type D8_DC20 struct {
	Cf28 _dafny.Sequence
}

func (D8_DC20) isD8() {}

func (CompanionStruct_D8_) Create_DC20_(Cf28 _dafny.Sequence) D8 {
	return D8{D8_DC20{Cf28}}
}

func (_this D8) Is_DC20() bool {
	_, ok := _this.Get_().(D8_DC20)
	return ok
}

type D8_DC23 struct {
	Cf35 D8
}

func (D8_DC23) isD8() {}

func (CompanionStruct_D8_) Create_DC23_(Cf35 D8) D8 {
	return D8{D8_DC23{Cf35}}
}

func (_this D8) Is_DC23() bool {
	_, ok := _this.Get_().(D8_DC23)
	return ok
}

func (CompanionStruct_D8_) Default() D8 {
	return Companion_D8_.Create_DC21_(_dafny.EmptySeq, _dafny.CodePoint('D'), _dafny.EmptySet)
}

func (_this D8) Dtor_cf29() _dafny.Sequence {
	return _this.Get_().(D8_DC21).Cf29
}

func (_this D8) Dtor_cf30() _dafny.CodePoint {
	return _this.Get_().(D8_DC21).Cf30
}

func (_this D8) Dtor_cf31() _dafny.Set {
	return _this.Get_().(D8_DC21).Cf31
}

func (_this D8) Dtor_cf32() bool {
	return _this.Get_().(D8_DC22).Cf32
}

func (_this D8) Dtor_cf33() _dafny.Map {
	return _this.Get_().(D8_DC22).Cf33
}

func (_this D8) Dtor_cf34() _dafny.Int {
	return _this.Get_().(D8_DC22).Cf34
}

func (_this D8) Dtor_cf28() _dafny.Sequence {
	return _this.Get_().(D8_DC20).Cf28
}

func (_this D8) Dtor_cf35() D8 {
	return _this.Get_().(D8_DC23).Cf35
}

func (_this D8) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D8_DC21:
		{
			return "D8.DC21" + "(" + _dafny.String(data.Cf29) + ", " + _dafny.String(data.Cf30) + ", " + _dafny.String(data.Cf31) + ")"
		}
	case D8_DC22:
		{
			return "D8.DC22" + "(" + _dafny.String(data.Cf32) + ", " + _dafny.String(data.Cf33) + ", " + _dafny.String(data.Cf34) + ")"
		}
	case D8_DC20:
		{
			return "D8.DC20" + "(" + _dafny.String(data.Cf28) + ")"
		}
	case D8_DC23:
		{
			return "D8.DC23" + "(" + _dafny.String(data.Cf35) + ")"
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
			return ok && data1.Cf29.Equals(data2.Cf29) && data1.Cf30 == data2.Cf30 && data1.Cf31.Equals(data2.Cf31)
		}
	case D8_DC22:
		{
			data2, ok := other.Get_().(D8_DC22)
			return ok && data1.Cf32 == data2.Cf32 && data1.Cf33.Equals(data2.Cf33) && data1.Cf34.Cmp(data2.Cf34) == 0
		}
	case D8_DC20:
		{
			data2, ok := other.Get_().(D8_DC20)
			return ok && data1.Cf28.Equals(data2.Cf28)
		}
	case D8_DC23:
		{
			data2, ok := other.Get_().(D8_DC23)
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

type D9_DC25 struct {
	Cf37 bool
	Cf38 bool
	Cf39 bool
}

func (D9_DC25) isD9() {}

func (CompanionStruct_D9_) Create_DC25_(Cf37 bool, Cf38 bool, Cf39 bool) D9 {
	return D9{D9_DC25{Cf37, Cf38, Cf39}}
}

func (_this D9) Is_DC25() bool {
	_, ok := _this.Get_().(D9_DC25)
	return ok
}

type D9_DC26 struct {
	Cf40 bool
	Cf41 _dafny.Int
}

func (D9_DC26) isD9() {}

func (CompanionStruct_D9_) Create_DC26_(Cf40 bool, Cf41 _dafny.Int) D9 {
	return D9{D9_DC26{Cf40, Cf41}}
}

func (_this D9) Is_DC26() bool {
	_, ok := _this.Get_().(D9_DC26)
	return ok
}

type D9_DC24 struct {
	Cf36 _dafny.Sequence
}

func (D9_DC24) isD9() {}

func (CompanionStruct_D9_) Create_DC24_(Cf36 _dafny.Sequence) D9 {
	return D9{D9_DC24{Cf36}}
}

func (_this D9) Is_DC24() bool {
	_, ok := _this.Get_().(D9_DC24)
	return ok
}

type D9_DC27 struct {
	Cf42 D9
}

func (D9_DC27) isD9() {}

func (CompanionStruct_D9_) Create_DC27_(Cf42 D9) D9 {
	return D9{D9_DC27{Cf42}}
}

func (_this D9) Is_DC27() bool {
	_, ok := _this.Get_().(D9_DC27)
	return ok
}

func (CompanionStruct_D9_) Default() D9 {
	return Companion_D9_.Create_DC25_(false, false, false)
}

func (_this D9) Dtor_cf37() bool {
	return _this.Get_().(D9_DC25).Cf37
}

func (_this D9) Dtor_cf38() bool {
	return _this.Get_().(D9_DC25).Cf38
}

func (_this D9) Dtor_cf39() bool {
	return _this.Get_().(D9_DC25).Cf39
}

func (_this D9) Dtor_cf40() bool {
	return _this.Get_().(D9_DC26).Cf40
}

func (_this D9) Dtor_cf41() _dafny.Int {
	return _this.Get_().(D9_DC26).Cf41
}

func (_this D9) Dtor_cf36() _dafny.Sequence {
	return _this.Get_().(D9_DC24).Cf36
}

func (_this D9) Dtor_cf42() D9 {
	return _this.Get_().(D9_DC27).Cf42
}

func (_this D9) String() string {
	switch data := _this.Get_().(type) {
	case nil:
		return "null"
	case D9_DC25:
		{
			return "D9.DC25" + "(" + _dafny.String(data.Cf37) + ", " + _dafny.String(data.Cf38) + ", " + _dafny.String(data.Cf39) + ")"
		}
	case D9_DC26:
		{
			return "D9.DC26" + "(" + _dafny.String(data.Cf40) + ", " + _dafny.String(data.Cf41) + ")"
		}
	case D9_DC24:
		{
			return "D9.DC24" + "(" + _dafny.String(data.Cf36) + ")"
		}
	case D9_DC27:
		{
			return "D9.DC27" + "(" + _dafny.String(data.Cf42) + ")"
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
			return ok && data1.Cf37 == data2.Cf37 && data1.Cf38 == data2.Cf38 && data1.Cf39 == data2.Cf39
		}
	case D9_DC26:
		{
			data2, ok := other.Get_().(D9_DC26)
			return ok && data1.Cf40 == data2.Cf40 && data1.Cf41.Cmp(data2.Cf41) == 0
		}
	case D9_DC24:
		{
			data2, ok := other.Get_().(D9_DC24)
			return ok && data1.Cf36.Equals(data2.Cf36)
		}
	case D9_DC27:
		{
			data2, ok := other.Get_().(D9_DC27)
			return ok && data1.Cf42.Equals(data2.Cf42)
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

// Definition of trait T0
type T0 interface {
	String() string
	Fm1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int
	Fm2(p0 _dafny.Set, p1 D0, p2 _dafny.Int, globalState *GlobalState) _dafny.Map
	M1(globalState *GlobalState)
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

// Definition of class GlobalState
type GlobalState struct {
	F4   _dafny.CodePoint
	F5   _dafny.Set
	F6   _dafny.Set
	F12  _dafny.Array
	_f0  _dafny.Int
	_f1  _dafny.Int
	_f2  _dafny.MultiSet
	_f3  bool
	_f7  _dafny.CodePoint
	_f8  _dafny.MultiSet
	_f9  bool
	_f10 _dafny.Int
	_f11 _dafny.Sequence
	_f13 bool
}

func New_GlobalState_() *GlobalState {
	_this := GlobalState{}

	_this.F4 = _dafny.CodePoint('D')
	_this.F5 = _dafny.EmptySet
	_this.F6 = _dafny.EmptySet
	_this.F12 = _dafny.NewArrayWithValue(nil, _dafny.IntOf(0))
	_this._f0 = _dafny.Zero
	_this._f1 = _dafny.Zero
	_this._f2 = _dafny.EmptyMultiSet
	_this._f3 = false
	_this._f7 = _dafny.CodePoint('D')
	_this._f8 = _dafny.EmptyMultiSet
	_this._f9 = false
	_this._f10 = _dafny.Zero
	_this._f11 = _dafny.EmptySeq
	_this._f13 = false
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

func (_this *GlobalState) Ctor__(f0 _dafny.Int, f1 _dafny.Int, f2 _dafny.MultiSet, f3 bool, f4 _dafny.CodePoint, f5 _dafny.Set, f6 _dafny.Set, f7 _dafny.CodePoint, f8 _dafny.MultiSet, f9 bool, f10 _dafny.Int, f11 _dafny.Sequence, f12 _dafny.Array, f13 bool) {
	{
		(_this)._f0 = f0
		(_this)._f1 = f1
		(_this)._f2 = f2
		(_this)._f3 = f3
		(_this).F4 = f4
		(_this).F5 = f5
		(_this).F6 = f6
		(_this)._f7 = f7
		(_this)._f8 = f8
		(_this)._f9 = f9
		(_this)._f10 = f10
		(_this)._f11 = f11
		(_this).F12 = f12
		(_this)._f13 = f13
	}
}
func (_this *GlobalState) F0() _dafny.Int {
	{
		return _this._f0
	}
}
func (_this *GlobalState) F1() _dafny.Int {
	{
		return _this._f1
	}
}
func (_this *GlobalState) F2() _dafny.MultiSet {
	{
		return _this._f2
	}
}
func (_this *GlobalState) F3() bool {
	{
		return _this._f3
	}
}
func (_this *GlobalState) F7() _dafny.CodePoint {
	{
		return _this._f7
	}
}
func (_this *GlobalState) F8() _dafny.MultiSet {
	{
		return _this._f8
	}
}
func (_this *GlobalState) F9() bool {
	{
		return _this._f9
	}
}
func (_this *GlobalState) F10() _dafny.Int {
	{
		return _this._f10
	}
}
func (_this *GlobalState) F11() _dafny.Sequence {
	{
		return _this._f11
	}
}
func (_this *GlobalState) F13() bool {
	{
		return _this._f13
	}
}

// End of class GlobalState

// Definition of class C0
type C0 struct {
	_f16 bool
}

func New_C0_() *C0 {
	_this := C0{}

	_this._f16 = false
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C0{}
var _ _dafny.TraitOffspring = &C0{}

func (_this *C0) Ctor__(f16 bool) {
	{
		(_this)._f16 = f16
	}
}
func (_this *C0) Fm1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.MultiSetOf((_this).F16(), !((_this).F16()))).Cardinality()
	}
}
func (_this *C0) Fm2(p0 _dafny.Set, p1 D0, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (func() _dafny.Map {
			var _coll9 = _dafny.NewMapBuilder()
			_ = _coll9
			for _iter9 := _dafny.Iterate((_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('m'), _dafny.CodePoint('q'), _dafny.CodePoint('g'))).Elements()); ; {
				_compr_9, _ok9 := _iter9()
				if !_ok9 {
					break
				}
				var _432_v0 _dafny.CodePoint
				_432_v0 = interface{}(_compr_9).(_dafny.CodePoint)
				if _dafny.Companion_Sequence_.Contains(_dafny.SeqOf(_dafny.CodePoint('p'), _dafny.CodePoint('m'), _dafny.CodePoint('q'), _dafny.CodePoint('g')), _432_v0) {
					_coll9.Add(_432_v0, _dafny.UnicodeSeqOfUtf8Bytes("tmqlcma"))
				}
			}
			return _coll9.ToMap()
		}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('r'), _dafny.UnicodeSeqOfUtf8Bytes("qky"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('l'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(467))).Uint32(), func(coer11 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg11 _dafny.Int) interface{} {
				return coer11(arg11)
			}
		}(func(_433_i0 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('r')
		})))))
	}
}
func (_this *C0) Fm5(p0 bool, p1 _dafny.Int, p2 _dafny.Int, p3 bool, globalState *GlobalState) _dafny.Sequence {
	{
		return _dafny.Companion_Sequence_.Concatenate(_dafny.UnicodeSeqOfUtf8Bytes("tv"), (func() _dafny.Sequence {
			if (_this).F16() {
				return _dafny.UnicodeSeqOfUtf8Bytes("usgh")
			}
			return _dafny.UnicodeSeqOfUtf8Bytes("xjrsoedx")
		})())
	}
}
func (_this *C0) M1(globalState *GlobalState) {
	{
		var _434_v0 _dafny.Array
		_ = _434_v0
		var _len0_7 _dafny.Int = _dafny.IntOfInt64(4)
		_ = _len0_7
		var _nw76 _dafny.Array
		_ = _nw76
		if _len0_7.Cmp(_dafny.Zero) == 0 {
			_nw76 = _dafny.NewArray(_len0_7)
		} else {
			var _init7 func(_dafny.Int) bool = func(_435_i0 _dafny.Int) bool {
				return (_this).F16()
			}
			_ = _init7
			var _element0_7 = _init7(_dafny.Zero)
			_ = _element0_7
			_nw76 = _dafny.NewArrayFromExample(_element0_7, nil, _len0_7)
			(_nw76).ArraySet1(_element0_7, 0)
			var _nativeLen0_7 = (_len0_7).Int()
			_ = _nativeLen0_7
			for _i0_7 := 1; _i0_7 < _nativeLen0_7; _i0_7++ {
				(_nw76).ArraySet1(_init7(_dafny.IntOf(_i0_7)), _i0_7)
			}
		}
		_434_v0 = _nw76
		var _index86 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))
		_ = _index86
		(_434_v0).ArraySet1((_this).F16(), (_index86).Int())
		var _436_v1 D1
		_ = _436_v1
		_436_v1 = Companion_D1_.Create_DC5_((_this).F16())
		var _index87 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))
		_ = _index87
		(_434_v0).ArraySet1(func(_source8 D1) bool {
			if _source8.Is_DC5() {
				var _437___mcc_h0 bool = _source8.Get_().(D1_DC5).Cf9
				_ = _437___mcc_h0
				var _438_cf9 bool = _437___mcc_h0
				_ = _438_cf9
				return ((_dafny.SetOf((_this).F16())).Cardinality()).Cmp(_dafny.IntOfInt64(444)) == 0
			} else {
				var _439___mcc_h1 _dafny.Sequence = _source8.Get_().(D1_DC4).Cf8
				_ = _439___mcc_h1
				var _440_cf8 _dafny.Sequence = _439___mcc_h1
				_ = _440_cf8
				return (_this).F16()
			}
		}(_436_v1), (_index87).Int())
		var _441_v2 _dafny.Int
		_ = _441_v2
		_441_v2 = _dafny.IntOfInt64(96)
		var _442_v3 _dafny.Set
		_ = _442_v3
		_442_v3 = _dafny.SetOf(_441_v2, _441_v2)
		var _443_v4 _dafny.Sequence
		_ = _443_v4
		_443_v4 = _dafny.SeqOf(_442_v3)
		var _444_v5 _dafny.Sequence
		_ = _444_v5
		_444_v5 = _dafny.UnicodeSeqOfUtf8Bytes("oobutohe")
		var _445_v6 _dafny.Sequence
		_ = _445_v6
		_445_v6 = _dafny.SeqOf((_dafny.Zero).Minus((_dafny.Zero).Minus(_441_v2)))
		var _446_i1 _dafny.Int
		_ = _446_i1
		_446_i1 = _dafny.Zero
		{
			for ((((_443_v4).Select((Companion_Default___.SafeIndex(_dafny.IntOfUint32((_444_v5).Cardinality()), _dafny.IntOfUint32((_443_v4).Cardinality()))).Uint32()).(_dafny.Set)).Union(_442_v3)).Cardinality()).Cmp(_dafny.IntOfUint32((_445_v6).Cardinality())) <= 0 {
				{
					if (_446_i1).Cmp(_dafny.IntOfInt64(100)) >= 0 {
						goto L5
					}
					_446_i1 = (_446_i1).Plus(_dafny.One)
					var _447_v7 _dafny.Array
					_ = _447_v7
					var _len0_8 _dafny.Int = _dafny.IntOfInt64(14)
					_ = _len0_8
					var _nw77 _dafny.Array
					_ = _nw77
					if _len0_8.Cmp(_dafny.Zero) == 0 {
						_nw77 = _dafny.NewArray(_len0_8)
					} else {
						var _init8 func(_dafny.Int) _dafny.Sequence = (func(_448_v0 _dafny.Array) func(_dafny.Int) _dafny.Sequence {
							return func(_449_i2 _dafny.Int) _dafny.Sequence {
								return _dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf((_this).F16(), (_this).F16()), _dafny.SeqOf((_448_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_448_v0), 0))).Int()).(bool)))
							}
						})(_434_v0)
						_ = _init8
						var _element0_8 = _init8(_dafny.Zero)
						_ = _element0_8
						_nw77 = _dafny.NewArrayFromExample(_element0_8, nil, _len0_8)
						(_nw77).ArraySet1(_element0_8, 0)
						var _nativeLen0_8 = (_len0_8).Int()
						_ = _nativeLen0_8
						for _i0_8 := 1; _i0_8 < _nativeLen0_8; _i0_8++ {
							(_nw77).ArraySet1(_init8(_dafny.IntOf(_i0_8)), _i0_8)
						}
					}
					_447_v7 = _nw77
					var _450_v8 _dafny.Sequence
					_ = _450_v8
					_450_v8 = _dafny.SeqOf((_this).F16(), false, (_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool), !((_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool)), Companion_Default___.Fm6((_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool), _441_v2, globalState))
					var _index88 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_447_v7), 0))
					_ = _index88
					(_447_v7).ArraySet1(_dafny.Companion_Sequence_.Concatenate(_450_v8, _450_v8), (_index88).Int())
					var _index89 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_447_v7), 0))
					_ = _index89
					(_447_v7).ArraySet1(_dafny.Companion_Sequence_.Update(_450_v8, (Companion_Default___.SafeIndex(_441_v2, _dafny.IntOfUint32((_450_v8).Cardinality()))).Uint32(), (_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool)), (_index89).Int())
					var _451_v9 bool
					_ = _451_v9
					var _452_v10 _dafny.Array
					_ = _452_v10
					var _453_v11 _dafny.Int
					_ = _453_v11
					var _454_v12 _dafny.Int
					_ = _454_v12
					var _out32 bool
					_ = _out32
					var _out33 _dafny.Array
					_ = _out33
					var _out34 _dafny.Int
					_ = _out34
					var _out35 _dafny.Int
					_ = _out35
					_out32, _out33, _out34, _out35 = Companion_Default___.M0((_dafny.Zero).Minus((_441_v2).Times(_441_v2)), (_447_v7).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(22), _dafny.ArrayLen((_447_v7), 0))).Int()).(_dafny.Sequence), globalState)
					_451_v9 = _out32
					_452_v10 = _out33
					_453_v11 = _out34
					_454_v12 = _out35
					var _455_v13 _dafny.Array
					_ = _455_v13
					var _nw78 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
					_ = _nw78
					_455_v13 = _nw78
					var _456_v14 _dafny.CodePoint
					_ = _456_v14
					_456_v14 = _dafny.CodePoint('f')
					var _457_v15 _dafny.Map
					_ = _457_v15
					_457_v15 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.UnicodeSeqOfUtf8Bytes("wluaxmjb"), _456_v14)
					var _index90 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_455_v13), 0))
					_ = _index90
					(_455_v13).ArraySet1((_441_v2).Times((_457_v15).Cardinality()), (_index90).Int())
					var _index91 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_455_v13), 0))
					_ = _index91
					(_455_v13).ArraySet1(Companion_Default___.SafeModuloInt(_441_v2, _dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_445_v6, _445_v6)).Cardinality())), (_index91).Int())
					var _458_v16 _dafny.MultiSet
					_ = _458_v16
					_458_v16 = _dafny.MultiSetOf(_dafny.IntOfUint32((_444_v5).Cardinality()), (_455_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_455_v13), 0))).Int()).(_dafny.Int), _dafny.IntOfInt64(876))
					var _pat_let_tv19 = _458_v16
					_ = _pat_let_tv19
					var _source9 D0 = func(_pat_let10_0 D0) D0 {
						return func(_459_dt__update__tmp_h0 D0) D0 {
							return func(_pat_let11_0 _dafny.Int) D0 {
								return func(_460_dt__update_hcf5_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC2_((_459_dt__update__tmp_h0).Dtor_cf4(), _460_dt__update_hcf5_h0, (_459_dt__update__tmp_h0).Dtor_cf6())
								}(_pat_let11_0)
							}((_pat_let_tv19).Cardinality())
						}(_pat_let10_0)
					}(Companion_Default___.Fm7(globalState))
					_ = _source9
					if _source9.Is_DC0() {
						var _461___mcc_h2 _dafny.Int = _source9.Get_().(D0_DC0).Cf0
						_ = _461___mcc_h2
						var _462___mcc_h3 bool = _source9.Get_().(D0_DC0).Cf1
						_ = _462___mcc_h3
						var _463_cf1 bool = _462___mcc_h3
						_ = _463_cf1
						var _464_cf0 _dafny.Int = _461___mcc_h2
						_ = _464_cf0
						var _index92 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))
						_ = _index92
						(_434_v0).ArraySet1(_451_v9, (_index92).Int())
						var _465_v17 _dafny.Map
						_ = _465_v17
						_465_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_455_v13, !(_451_v9))
						_465_v17 = (_465_v17).Update(_455_v13, Companion_Default___.Fm6(_451_v9, (_442_v3).Cardinality(), globalState))
						var _466_v18 D0
						_ = _466_v18
						_466_v18 = Companion_D0_.Create_DC0_((_dafny.SetOf(_441_v2, _dafny.IntOfInt64(682))).Cardinality(), _451_v9)
						_463_cf1 = !(!(!(!((_466_v18).Dtor_cf1()))))
						var _467_v19 bool
						_ = _467_v19
						var _468_v20 _dafny.Array
						_ = _468_v20
						var _469_v21 _dafny.Int
						_ = _469_v21
						var _470_v22 _dafny.Int
						_ = _470_v22
						var _out36 bool
						_ = _out36
						var _out37 _dafny.Array
						_ = _out37
						var _out38 _dafny.Int
						_ = _out38
						var _out39 _dafny.Int
						_ = _out39
						_out36, _out37, _out38, _out39 = Companion_Default___.M0(((_455_v13).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_455_v13), 0))).Int()).(_dafny.Int)).Minus(_464_cf0), _dafny.SeqOf((_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool)), globalState)
						_467_v19 = _out36
						_468_v20 = _out37
						_469_v21 = _out38
						_470_v22 = _out39
					} else if _source9.Is_DC1() {
						var _471___mcc_h4 _dafny.Set = _source9.Get_().(D0_DC1).Cf2
						_ = _471___mcc_h4
						var _472___mcc_h5 bool = _source9.Get_().(D0_DC1).Cf3
						_ = _472___mcc_h5
						var _473_cf3 bool = _472___mcc_h5
						_ = _473_cf3
						var _474_cf2 _dafny.Set = _471___mcc_h4
						_ = _474_cf2
						_441_v2 = _dafny.IntOfInt64(868)
						var _475_v23 _dafny.Array
						_ = _475_v23
						var _len0_9 _dafny.Int = _dafny.IntOfInt64(7)
						_ = _len0_9
						var _nw79 _dafny.Array
						_ = _nw79
						if _len0_9.Cmp(_dafny.Zero) == 0 {
							_nw79 = _dafny.NewArray(_len0_9)
						} else {
							var _init9 func(_dafny.Int) _dafny.Map = (func(_476_v14 _dafny.CodePoint, _477_v0 _dafny.Array) func(_dafny.Int) _dafny.Map {
								return func(_478_i3 _dafny.Int) _dafny.Map {
									return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_476_v14, (_477_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_477_v0), 0))).Int()).(bool))
								}
							})(_456_v14, _434_v0)
							_ = _init9
							var _element0_9 = _init9(_dafny.Zero)
							_ = _element0_9
							_nw79 = _dafny.NewArrayFromExample(_element0_9, nil, _len0_9)
							(_nw79).ArraySet1(_element0_9, 0)
							var _nativeLen0_9 = (_len0_9).Int()
							_ = _nativeLen0_9
							for _i0_9 := 1; _i0_9 < _nativeLen0_9; _i0_9++ {
								(_nw79).ArraySet1(_init9(_dafny.IntOf(_i0_9)), _i0_9)
							}
						}
						_475_v23 = _nw79
						_475_v23 = _475_v23
						var _479_v24 D0
						_ = _479_v24
						_479_v24 = Companion_D0_.Create_DC3_(_456_v14)
						_479_v24 = _479_v24
						(globalState).F12 = _455_v13
					} else if _source9.Is_DC2() {
						var _480___mcc_h6 _dafny.Set = _source9.Get_().(D0_DC2).Cf4
						_ = _480___mcc_h6
						var _481___mcc_h7 _dafny.Int = _source9.Get_().(D0_DC2).Cf5
						_ = _481___mcc_h7
						var _482___mcc_h8 bool = _source9.Get_().(D0_DC2).Cf6
						_ = _482___mcc_h8
						var _483_cf6 bool = _482___mcc_h8
						_ = _483_cf6
						var _484_cf5 _dafny.Int = _481___mcc_h7
						_ = _484_cf5
						var _485_cf4 _dafny.Set = _480___mcc_h6
						_ = _485_cf4
						_444_v5 = _dafny.UnicodeSeqOfUtf8Bytes("qf")
						var _index93 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(462), _dafny.ArrayLen((_455_v13), 0))
						_ = _index93
						(_455_v13).ArraySet1(_453_v11, (_index93).Int())
						var _486_v25 _dafny.Array
						_ = _486_v25
						var _nwElement0_16 _dafny.CodePoint = _456_v14
						_ = _nwElement0_16
						var _nw80 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_16, nil, _dafny.IntOfInt64(2))
						_ = _nw80
						(_nw80).ArraySet1CodePoint(_nwElement0_16, 0)
						(_nw80).ArraySet1CodePoint(_456_v14, 1)
						_486_v25 = _nw80
						var _index94 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_486_v25), 0))
						_ = _index94
						(_486_v25).ArraySet1CodePoint(_456_v14, (_index94).Int())
						var _index95 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(885), _dafny.ArrayLen((_486_v25), 0))
						_ = _index95
						(_486_v25).ArraySet1CodePoint(_456_v14, (_index95).Int())
						var _487_v26 _dafny.Map
						_ = _487_v26
						_487_v26 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _dafny.SetOf(_dafny.IntOfUint32((_444_v5).Cardinality()), _441_v2, _484_cf5))
						var _488_v27 D0
						_ = _488_v27
						_488_v27 = Companion_D0_.Create_DC2_(_485_cf4, _dafny.IntOfUint32((_450_v8).Cardinality()), _451_v9)
						var _489_v28 D0
						_ = _489_v28
						_489_v28 = Companion_D0_.Create_DC2_((func() _dafny.Set {
							if (_487_v26).Contains(_451_v9) {
								return (_487_v26).Get(_451_v9).(_dafny.Set)
							}
							return _442_v3
						})(), (_488_v27).Dtor_cf5(), true)
						_489_v28 = Companion_Default___.Fm7(globalState)
					} else {
						var _490___mcc_h9 _dafny.CodePoint = _source9.Get_().(D0_DC3).Cf7
						_ = _490___mcc_h9
						var _491_cf7 _dafny.CodePoint = _490___mcc_h9
						_ = _491_cf7
						var _492_v29 bool
						_ = _492_v29
						var _493_v30 _dafny.Array
						_ = _493_v30
						var _494_v31 _dafny.Int
						_ = _494_v31
						var _495_v32 _dafny.Int
						_ = _495_v32
						var _out40 bool
						_ = _out40
						var _out41 _dafny.Array
						_ = _out41
						var _out42 _dafny.Int
						_ = _out42
						var _out43 _dafny.Int
						_ = _out43
						_out40, _out41, _out42, _out43 = Companion_Default___.M0(_453_v11, _dafny.SeqOf((_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool)), globalState)
						_492_v29 = _out40
						_493_v30 = _out41
						_494_v31 = _out42
						_495_v32 = _out43
						_491_cf7 = (_444_v5).Select((Companion_Default___.SafeIndex(_dafny.IntOfInt64(104), _dafny.IntOfUint32((_444_v5).Cardinality()))).Uint32()).(_dafny.CodePoint)
						var _496_v33 _dafny.Map
						_ = _496_v33
						_496_v33 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfInt64(-902), (func() bool {
							if !(_451_v9) {
								return (_this).F16()
							}
							return (_this).F16()
						})())
						_496_v33 = (_496_v33).Merge(_496_v33)
						_434_v0 = _434_v0
					}
					goto C5
				}
			C5:
			}
			goto L5
		}
	L5:
		var _497_v34 _dafny.Array
		_ = _497_v34
		var _nw81 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(26))
		_ = _nw81
		_497_v34 = _nw81
		var _index96 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))
		_ = _index96
		(_497_v34).ArraySet1(_441_v2, (_index96).Int())
		var _index97 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))
		_ = _index97
		(_497_v34).ArraySet1((_this).Fm1(_444_v5, _441_v2, (func() _dafny.Int {
			if (_this).F16() {
				return _dafny.IntOfUint32((_444_v5).Cardinality())
			}
			return _441_v2
		})(), globalState), (_index97).Int())
		var _498_v35 _dafny.CodePoint
		_ = _498_v35
		_498_v35 = _dafny.CodePoint('f')
		var _499_v36 _dafny.Map
		_ = _499_v36
		_499_v36 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F16(), _498_v35)
		(globalState).F4 = (func() _dafny.CodePoint {
			if (_499_v36).Contains((_this).F16()) {
				return (_499_v36).Get((_this).F16()).(_dafny.CodePoint)
			}
			return _dafny.CodePoint('v')
		})()
		var _500_v37 _dafny.Map
		_ = _500_v37
		_500_v37 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(((_497_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))).Int()).(_dafny.Int)).Cmp((_497_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))).Int()).(_dafny.Int)) > 0, _441_v2)
		var _501_v38 _dafny.Map
		_ = _501_v38
		_501_v38 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_442_v3, _dafny.IntOfInt64(-717))
		var _502_v39 D0
		_ = _502_v39
		_502_v39 = Companion_D0_.Create_DC2_(_442_v3, (_497_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))).Int()).(_dafny.Int), (_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool))
		var _503_v40 _dafny.Map
		_ = _503_v40
		_503_v40 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_441_v2, (_434_v0).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(848), _dafny.ArrayLen((_434_v0), 0))).Int()).(bool))
		_500_v37 = (_500_v37).Update((_this).F16(), (func() _dafny.Int {
			if (_this).F16() {
				return (_497_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))).Int()).(_dafny.Int)
			}
			return (func() _dafny.Int {
				if (_501_v38).Contains((_502_v39).Dtor_cf4()) {
					return (_501_v38).Get((_502_v39).Dtor_cf4()).(_dafny.Int)
				}
				return (_503_v40).Cardinality()
			})()
		})())
		var _index98 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))
		_ = _index98
		(_497_v34).ArraySet1((_497_v34).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(332), _dafny.ArrayLen((_497_v34), 0))).Int()).(_dafny.Int), (_index98).Int())
	}
}
func (_this *C0) F16() bool {
	{
		return _this._f16
	}
}

// End of class C0

// Definition of class C1
type C1 struct {
	_f14 bool
	_f15 _dafny.Int
}

func New_C1_() *C1 {
	_this := C1{}

	_this._f14 = false
	_this._f15 = _dafny.Zero
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
	return [](*_dafny.TraitID){Companion_T0_.TraitID_}
}

var _ T0 = &C1{}
var _ _dafny.TraitOffspring = &C1{}

func (_this *C1) Ctor__(f14 bool, f15 _dafny.Int) {
	{
		(_this)._f14 = f14
		(_this)._f15 = f15
	}
}
func (_this *C1) Fm1(p0 _dafny.Sequence, p1 _dafny.Int, p2 _dafny.Int, globalState *GlobalState) _dafny.Int {
	{
		return (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate((func() _dafny.Sequence {
			if (_this).F14() {
				return _dafny.SeqOf(_dafny.SeqOf((_this).F15(), (_this).F15(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F14())).Cardinality(), (_this).F15()))
			}
			return _dafny.SeqOf(_dafny.SeqOf((_this).F15(), (_this).F15()))
		})(), _dafny.SeqOf(_dafny.SeqOf(_dafny.IntOfInt64(512)), _dafny.SeqOf((_dafny.SetOf((_this).F14())).Cardinality(), (_dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F15())).Cardinality())))).Cardinality()))
	}
}
func (_this *C1) Fm2(p0 _dafny.Set, p1 D0, p2 _dafny.Int, globalState *GlobalState) _dafny.Map {
	{
		return (func() _dafny.Map {
			var _coll10 = _dafny.NewMapBuilder()
			_ = _coll10
			for _iter10 := _dafny.Iterate((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(861))).Uint32(), func(coer12 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg12 _dafny.Int) interface{} {
					return coer12(arg12)
				}
			}(func(_504_i0 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('j')
			})))).Keys().Elements()); ; {
				_compr_10, _ok10 := _iter10()
				if !_ok10 {
					break
				}
				var _505_v0 _dafny.CodePoint
				_505_v0 = interface{}(_compr_10).(_dafny.CodePoint)
				if (_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('x'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(861))).Uint32(), func(coer13 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
					return func(arg13 _dafny.Int) interface{} {
						return coer13(arg13)
					}
				}(func(_504_i0 _dafny.Int) _dafny.CodePoint {
					return _dafny.CodePoint('j')
				})))).Contains(_505_v0) {
					_coll10.Add(_505_v0, _dafny.UnicodeSeqOfUtf8Bytes("evyrhbeq"))
				}
			}
			return _coll10.ToMap()
		}()).Merge((_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('v'), _dafny.UnicodeSeqOfUtf8Bytes("ycuwefvef"))).Merge(_dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.CodePoint('h'), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(581))).Uint32(), func(coer14 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
			return func(arg14 _dafny.Int) interface{} {
				return coer14(arg14)
			}
		}(func(_506_i1 _dafny.Int) _dafny.CodePoint {
			return _dafny.CodePoint('x')
		})))))
	}
}
func (_this *C1) M1(globalState *GlobalState) {
	{
		var _507_v0 _dafny.Sequence
		_ = _507_v0
		_507_v0 = _dafny.UnicodeSeqOfUtf8Bytes("qqptj")
		var _508_v1 _dafny.MultiSet
		_ = _508_v1
		_508_v1 = _dafny.MultiSetOf(_507_v0, _dafny.Companion_Sequence_.Concatenate(_507_v0, _507_v0), _507_v0)
		var _509_v2 _dafny.Sequence
		_ = _509_v2
		_509_v2 = _dafny.SeqOf(_507_v0, _507_v0, (Companion_Default___.Fm3((_this).F15(), globalState)).Dtor_cf8())
		_508_v1 = ((_dafny.MultiSetFromSeq(_509_v2)).Intersection(_508_v1)).Union(_508_v1)
		_509_v2 = _509_v2
		var _hi0 _dafny.Int = (_this).F15()
		_ = _hi0
		for _510_i0 := (_this).F15(); _510_i0.Cmp(_hi0) < 0; _510_i0 = _510_i0.Plus(_dafny.One) {
			var _511_v3 _dafny.Set
			_ = _511_v3
			_511_v3 = _dafny.SetOf((_this).F14(), (_this).F14(), (_this).F14())
			var _512_v4 _dafny.MultiSet
			_ = _512_v4
			_512_v4 = _dafny.MultiSetOf(_dafny.IntOfUint32((_507_v0).Cardinality()))
			var _513_v5 _dafny.CodePoint
			_ = _513_v5
			_513_v5 = _dafny.CodePoint('r')
			var _514_v6 _dafny.Map
			_ = _514_v6
			_514_v6 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _513_v5)
			var _515_v7 _dafny.Map
			_ = _515_v7
			_515_v7 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F14())
			var _516_v8 _dafny.Sequence
			_ = _516_v8
			_516_v8 = _dafny.SeqOf(true, (_this).F14())
			var _517_v9 _dafny.Array
			_ = _517_v9
			var _nwElement0_17 bool = (_510_i0).Cmp(_510_i0) < 0
			_ = _nwElement0_17
			var _nw82 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_17, nil, _dafny.IntOfInt64(28))
			_ = _nw82
			(_nw82).ArraySet1(_nwElement0_17, 0)
			(_nw82).ArraySet1((_this).F14(), 1)
			(_nw82).ArraySet1((_this).F14(), 2)
			(_nw82).ArraySet1(true, 3)
			(_nw82).ArraySet1((Companion_D0_.Create_DC2_(_dafny.SetOf((_this).F15()), _dafny.IntOfInt64(180), true)).Dtor_cf6(), 4)
			(_nw82).ArraySet1((_this).F14(), 5)
			(_nw82).ArraySet1((_this).F14(), 6)
			(_nw82).ArraySet1((_this).F14(), 7)
			(_nw82).ArraySet1((_this).F14(), 8)
			(_nw82).ArraySet1((_511_v3).IsSubsetOf(_511_v3), 9)
			(_nw82).ArraySet1((_this).F14(), 10)
			(_nw82).ArraySet1((_this).F14(), 11)
			(_nw82).ArraySet1((_this).F14(), 12)
			(_nw82).ArraySet1((_this).F14(), 13)
			(_nw82).ArraySet1((_510_i0).Cmp((func() _dafny.Int {
				if (_512_v4).Contains(_dafny.IntOfInt64(488)) {
					return (_512_v4).Multiplicity(_dafny.IntOfInt64(488))
				}
				return (_514_v6).Cardinality()
			})()) > 0, 14)
			(_nw82).ArraySet1(!(!((_this).F14())) || ((_this).F14()), 15)
			(_nw82).ArraySet1((_this).F14(), 16)
			(_nw82).ArraySet1((_this).F14(), 17)
			(_nw82).ArraySet1((_510_i0).Cmp(_dafny.IntOfUint32((_507_v0).Cardinality())) < 0, 18)
			(_nw82).ArraySet1(!(((_515_v7).Cardinality()).Cmp(_dafny.IntOfInt64(759)) != 0), 19)
			(_nw82).ArraySet1((_this).F14(), 20)
			(_nw82).ArraySet1((_512_v4).IsProperSubsetOf((_512_v4).Update(_510_i0, Companion_Default___.Abs((_this).F15()))), 21)
			(_nw82).ArraySet1(!((func() bool {
				if (_this).F14() {
					return (_this).F14()
				}
				return (_this).F14()
			})()), 22)
			(_nw82).ArraySet1(!_dafny.Companion_Sequence_.Equal(_516_v8, _516_v8), 23)
			(_nw82).ArraySet1(!((_this).F14()), 24)
			(_nw82).ArraySet1(_dafny.Companion_Sequence_.Contains(_507_v0, _513_v5), 25)
			(_nw82).ArraySet1(_dafny.Companion_Sequence_.IsPrefixOf(_516_v8, _516_v8), 26)
			(_nw82).ArraySet1((_this).F14(), 27)
			_517_v9 = _nw82
			var _index99 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
			_ = _index99
			(_517_v9).ArraySet1(((_this).F15()).Cmp(_510_i0) >= 0, (_index99).Int())
			var _index100 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
			_ = _index100
			(_517_v9).ArraySet1((func() bool {
				if (_515_v7).Contains((_510_i0).Cmp((_this).F15()) == 0) {
					return (_515_v7).Get((_510_i0).Cmp((_this).F15()) == 0).(bool)
				}
				return ((_this).F14()) || ((_this).F14())
			})(), (_index100).Int())
			var _518_v11 _dafny.Array
			_ = _518_v11
			var _len0_10 _dafny.Int = _dafny.IntOfInt64(22)
			_ = _len0_10
			var _nw83 _dafny.Array
			_ = _nw83
			if _len0_10.Cmp(_dafny.Zero) == 0 {
				_nw83 = _dafny.NewArray(_len0_10)
			} else {
				var _init10 func(_dafny.Int) _dafny.Map = (func(_519_i0 _dafny.Int, _520_v3 _dafny.Set) func(_dafny.Int) _dafny.Map {
					return func(_521_i1 _dafny.Int) _dafny.Map {
						return func() _dafny.Map {
							var _coll11 = _dafny.NewMapBuilder()
							_ = _coll11
							for _iter11 := _dafny.Iterate((_dafny.SetOf((_this).F15(), _519_i0, (_dafny.Zero).Minus((_520_v3).Cardinality()), _519_i0, (_this).F15())).Elements()); ; {
								_compr_11, _ok11 := _iter11()
								if !_ok11 {
									break
								}
								var _522_v10 _dafny.Int
								_522_v10 = interface{}(_compr_11).(_dafny.Int)
								if (_dafny.SetOf((_this).F15(), _519_i0, (_dafny.Zero).Minus((_520_v3).Cardinality()), _519_i0, (_this).F15())).Contains(_522_v10) {
									_coll11.Add(Companion_Default___.SafeModuloInt(_522_v10, (_dafny.MultiSetOf((_this).F14(), !((_this).F14()), (_this).F14(), (_this).F14(), (_this).F14())).Cardinality()), _dafny.IntOfInt64(-420))
								}
							}
							return _coll11.ToMap()
						}()
					}
				})(_510_i0, _511_v3)
				_ = _init10
				var _element0_10 = _init10(_dafny.Zero)
				_ = _element0_10
				_nw83 = _dafny.NewArrayFromExample(_element0_10, nil, _len0_10)
				(_nw83).ArraySet1(_element0_10, 0)
				var _nativeLen0_10 = (_len0_10).Int()
				_ = _nativeLen0_10
				for _i0_10 := 1; _i0_10 < _nativeLen0_10; _i0_10++ {
					(_nw83).ArraySet1(_init10(_dafny.IntOf(_i0_10)), _i0_10)
				}
			}
			_518_v11 = _nw83
			var _len0_11 _dafny.Int = _dafny.IntOfInt64(19)
			_ = _len0_11
			var _nw84 _dafny.Array
			_ = _nw84
			if _len0_11.Cmp(_dafny.Zero) == 0 {
				_nw84 = _dafny.NewArray(_len0_11)
			} else {
				var _init11 func(_dafny.Int) _dafny.Map = (func(_523_i0 _dafny.Int) func(_dafny.Int) _dafny.Map {
					return func(_524_i2 _dafny.Int) _dafny.Map {
						return _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_523_i0, (Companion_D0_.Create_DC0_((_this).F15(), (_this).F14())).Dtor_cf0())
					}
				})(_510_i0)
				_ = _init11
				var _element0_11 = _init11(_dafny.Zero)
				_ = _element0_11
				_nw84 = _dafny.NewArrayFromExample(_element0_11, nil, _len0_11)
				(_nw84).ArraySet1(_element0_11, 0)
				var _nativeLen0_11 = (_len0_11).Int()
				_ = _nativeLen0_11
				for _i0_11 := 1; _i0_11 < _nativeLen0_11; _i0_11++ {
					(_nw84).ArraySet1(_init11(_dafny.IntOf(_i0_11)), _i0_11)
				}
			}
			_518_v11 = _nw84
			var _525_v12 _dafny.Sequence
			_ = _525_v12
			_525_v12 = _dafny.SeqOf((_this).F15())
			var _526_v13 D0
			_ = _526_v13
			_526_v13 = Companion_D0_.Create_DC0_(_dafny.IntOfUint32((_525_v12).Cardinality()), !((_this).F14()))
			var _source10 D0 = func(_pat_let12_0 D0) D0 {
				return func(_527_dt__update__tmp_h0 D0) D0 {
					return func(_pat_let13_0 bool) D0 {
						return func(_528_dt__update_hcf1_h0 bool) D0 {
							return Companion_D0_.Create_DC0_((_527_dt__update__tmp_h0).Dtor_cf0(), _528_dt__update_hcf1_h0)
						}(_pat_let13_0)
					}((_this).F14())
				}(_pat_let12_0)
			}(_526_v13)
			_ = _source10
			if _source10.Is_DC0() {
				var _529___mcc_h0 _dafny.Int = _source10.Get_().(D0_DC0).Cf0
				_ = _529___mcc_h0
				var _530___mcc_h1 bool = _source10.Get_().(D0_DC0).Cf1
				_ = _530___mcc_h1
				var _531_cf1 bool = _530___mcc_h1
				_ = _531_cf1
				var _532_cf0 _dafny.Int = _529___mcc_h0
				_ = _532_cf0
				_531_cf1 = (func() bool {
					if (_517_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))).Int()).(bool) {
						return !(Companion_Default___.Fm4((_this).F14(), _511_v3, (_this).F15(), globalState))
					}
					return (_this).F14()
				})()
				(globalState).F4 = _513_v5
				var _533_v14 *C0
				_ = _533_v14
				var _nw85 *C0 = New_C0_()
				_ = _nw85
				_nw85.Ctor__((_511_v3).IsProperSubsetOf(_511_v3))
				_533_v14 = _nw85
				var _index101 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
				_ = _index101
				(_517_v9).ArraySet1((_510_i0).Cmp(_510_i0) <= 0, (_index101).Int())
			} else if _source10.Is_DC1() {
				var _534___mcc_h2 _dafny.Set = _source10.Get_().(D0_DC1).Cf2
				_ = _534___mcc_h2
				var _535___mcc_h3 bool = _source10.Get_().(D0_DC1).Cf3
				_ = _535___mcc_h3
				var _536_cf3 bool = _535___mcc_h3
				_ = _536_cf3
				var _537_cf2 _dafny.Set = _534___mcc_h2
				_ = _537_cf2
				var _538_v15 _dafny.Int
				_ = _538_v15
				_538_v15 = _dafny.IntOfInt64(931)
				var _539_v16 _dafny.Array
				_ = _539_v16
				var _nw86 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptyMap, _dafny.IntOfInt64(19))
				_ = _nw86
				_539_v16 = _nw86
				var _index102 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_539_v16), 0))
				_ = _index102
				(_539_v16).ArraySet1(_515_v7, (_index102).Int())
				var _index103 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_539_v16), 0))
				_ = _index103
				var _index104 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
				_ = _index104
				var _rhs74 _dafny.Int = (_525_v12).Select((Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_525_v12).Cardinality()))).Uint32()).(_dafny.Int)
				_ = _rhs74
				var _rhs75 _dafny.Map = (_515_v7).Merge((_515_v7).Merge(_515_v7))
				_ = _rhs75
				var _rhs76 _dafny.Int = (Companion_Default___.SafeModuloInt(_538_v15, _510_i0)).Times(_510_i0)
				_ = _rhs76
				var _rhs77 bool = (func() bool {
					if (_this).F14() {
						return (_517_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))).Int()).(bool)
					}
					return !((_this).F14())
				})()
				_ = _rhs77
				var _lhs40 _dafny.Array = _539_v16
				_ = _lhs40
				var _lhs41 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(865), _dafny.ArrayLen((_539_v16), 0))
				_ = _lhs41
				var _lhs42 _dafny.Array = _517_v9
				_ = _lhs42
				var _lhs43 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
				_ = _lhs43
				_538_v15 = _rhs74
				(_lhs40).ArraySet1(_rhs75, (_lhs41).Int())
				_538_v15 = _rhs76
				(_lhs42).ArraySet1(_rhs77, (_lhs43).Int())
				var _540_v17 _dafny.Map
				_ = _540_v17
				_540_v17 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_510_i0, _dafny.SetOf(false))
				var _541_v18 _dafny.Map
				_ = _541_v18
				_541_v18 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_540_v17, (_this).F15())
				_541_v18 = (_541_v18).Update(_540_v17, _dafny.IntOfInt64(-424))
				var _542_v19 _dafny.Map
				_ = _542_v19
				_542_v19 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.MultiSetOf(_510_i0), (_this).F15())
				_538_v15 = Companion_Default___.SafeDivisionInt(((_this).F15()).Times((_this).F15()), (_542_v19).Cardinality())
				_536_cf3 = !(_536_cf3)
			} else if _source10.Is_DC2() {
				var _543___mcc_h4 _dafny.Set = _source10.Get_().(D0_DC2).Cf4
				_ = _543___mcc_h4
				var _544___mcc_h5 _dafny.Int = _source10.Get_().(D0_DC2).Cf5
				_ = _544___mcc_h5
				var _545___mcc_h6 bool = _source10.Get_().(D0_DC2).Cf6
				_ = _545___mcc_h6
				var _546_cf6 bool = _545___mcc_h6
				_ = _546_cf6
				var _547_cf5 _dafny.Int = _544___mcc_h5
				_ = _547_cf5
				var _548_cf4 _dafny.Set = _543___mcc_h4
				_ = _548_cf4
				_546_cf6 = !((_548_cf4).Equals(_548_cf4))
				(globalState).F4 = _513_v5
				var _index105 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
				_ = _index105
				(_517_v9).ArraySet1(((_this).F15()).Cmp((_this).Fm1(_507_v0, (_this).F15(), _dafny.IntOfInt64(602), globalState)) != 0, (_index105).Int())
				var _549_v20 *C0
				_ = _549_v20
				var _nw87 *C0 = New_C0_()
				_ = _nw87
				_nw87.Ctor__((_517_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))).Int()).(bool))
				_549_v20 = _nw87
				var _550_v21 _dafny.Sequence
				_ = _550_v21
				_550_v21 = _dafny.SeqOf(_549_v20, _549_v20, _549_v20)
				_546_cf6 = (func() bool {
					if ((_this).F15()).Cmp(_dafny.IntOfUint32((_550_v21).Cardinality())) != 0 {
						return (_549_v20).F16()
					}
					return false
				})()
			} else {
				var _551___mcc_h7 _dafny.CodePoint = _source10.Get_().(D0_DC3).Cf7
				_ = _551___mcc_h7
				var _552_cf7 _dafny.CodePoint = _551___mcc_h7
				_ = _552_cf7
				_507_v0 = _507_v0
				var _index106 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))
				_ = _index106
				(_517_v9).ArraySet1((_this).F14(), (_index106).Int())
				var _553_v22 _dafny.MultiSet
				_ = _553_v22
				_553_v22 = _dafny.MultiSetOf((func() bool {
					if (_517_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))).Int()).(bool) {
						return (_517_v9).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(466), _dafny.ArrayLen((_517_v9), 0))).Int()).(bool)
					}
					return true
				})())
				_553_v22 = _553_v22
				_517_v9 = _517_v9
			}
			var _554_v23 _dafny.Set
			_ = _554_v23
			_554_v23 = _dafny.SetOf(_511_v3)
			var _555_v24 D0
			_ = _555_v24
			_555_v24 = Companion_D0_.Create_DC1_((_554_v23).Difference(_554_v23), (_this).F14())
			_555_v24 = _555_v24
		}
		var _556_v25 _dafny.Int
		_ = _556_v25
		_556_v25 = _dafny.IntOfInt64(253)
		_556_v25 = _556_v25
		var _557_v26 _dafny.Set
		_ = _557_v26
		_557_v26 = _dafny.SetOf(_556_v25, _556_v25)
		var _558_v27 _dafny.Map
		_ = _558_v27
		_558_v27 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v25, _dafny.SetOf((_this).F14(), (_this).F14()))
		var _559_v28 _dafny.Set
		_ = _559_v28
		_559_v28 = _dafny.SetOf((_this).F14(), (_this).F14(), (_this).F14())
		var _560_v29 D0
		_ = _560_v29
		_560_v29 = Companion_D0_.Create_DC2_(_557_v26, ((_dafny.SetOf((_this).F14())).Union((func() _dafny.Set {
			if (_558_v27).Contains((_this).F15()) {
				return (_558_v27).Get((_this).F15()).(_dafny.Set)
			}
			return _559_v28
		})())).Cardinality(), (_this).F14())
		var _source11 D0 = _560_v29
		_ = _source11
		if _source11.Is_DC0() {
			var _561___mcc_h8 _dafny.Int = _source11.Get_().(D0_DC0).Cf0
			_ = _561___mcc_h8
			var _562___mcc_h9 bool = _source11.Get_().(D0_DC0).Cf1
			_ = _562___mcc_h9
			var _563_cf1 bool = _562___mcc_h9
			_ = _563_cf1
			var _564_cf0 _dafny.Int = _561___mcc_h8
			_ = _564_cf0
			var _565_v30 _dafny.MultiSet
			_ = _565_v30
			_565_v30 = _dafny.MultiSetOf((_this).F15(), _556_v25)
			var _566_v31 _dafny.Map
			_ = _566_v31
			_566_v31 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_565_v30).Difference(_565_v30), _556_v25)
			_566_v31 = (_566_v31).Merge((_566_v31).Merge(_566_v31))
			if (_563_cf1) || ((func() bool {
				if _563_cf1 {
					return _563_cf1
				}
				return _563_cf1
			})()) {
				_556_v25 = (_this).F15()
				var _567_v32 _dafny.MultiSet
				_ = _567_v32
				_567_v32 = _dafny.MultiSetOf(_563_cf1, false)
				_564_cf0 = ((_567_v32).Intersection(_567_v32)).Cardinality()
				var _568_v33 _dafny.Sequence
				_ = _568_v33
				_568_v33 = _dafny.SeqOf(_563_cf1)
				var _569_v34 _dafny.Map
				_ = _569_v34
				_569_v34 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_dafny.IntOfUint32((_507_v0).Cardinality()), (_this).F14())
				var _570_v35 _dafny.Array
				_ = _570_v35
				var _nwElement0_18 bool = !(_563_cf1)
				_ = _nwElement0_18
				var _nw88 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_18, nil, _dafny.IntOfInt64(25))
				_ = _nw88
				(_nw88).ArraySet1(_nwElement0_18, 0)
				(_nw88).ArraySet1(_563_cf1, 1)
				(_nw88).ArraySet1(_563_cf1, 2)
				(_nw88).ArraySet1((_this).F14(), 3)
				(_nw88).ArraySet1((func() bool {
					if _563_cf1 {
						return (_568_v33).Select((Companion_Default___.SafeIndex(_564_cf0, _dafny.IntOfUint32((_568_v33).Cardinality()))).Uint32()).(bool)
					}
					return false
				})(), 4)
				(_nw88).ArraySet1((_this).F14(), 5)
				(_nw88).ArraySet1(_563_cf1, 6)
				(_nw88).ArraySet1(false, 7)
				(_nw88).ArraySet1(!(((_this).F15()).Cmp((_dafny.Zero).Minus((_this).F15())) != 0), 8)
				(_nw88).ArraySet1(_dafny.Companion_Sequence_.Equal(_507_v0, _dafny.UnicodeSeqOfUtf8Bytes("wmjh")), 9)
				(_nw88).ArraySet1(false, 10)
				(_nw88).ArraySet1(!(_563_cf1), 11)
				(_nw88).ArraySet1((_this).F14(), 12)
				(_nw88).ArraySet1(!(_563_cf1) || ((_this).F14()), 13)
				(_nw88).ArraySet1((_565_v30).IsProperSubsetOf(_565_v30), 14)
				(_nw88).ArraySet1(false, 15)
				(_nw88).ArraySet1((func() bool {
					if (_569_v34).Contains(_564_cf0) {
						return (_569_v34).Get(_564_cf0).(bool)
					}
					return true
				})(), 16)
				(_nw88).ArraySet1((_567_v32).IsSubsetOf(_567_v32), 17)
				(_nw88).ArraySet1(_563_cf1, 18)
				(_nw88).ArraySet1((_this).F14(), 19)
				(_nw88).ArraySet1((_this).F14(), 20)
				(_nw88).ArraySet1(true, 21)
				(_nw88).ArraySet1((func() bool {
					if (_this).F14() {
						return _563_cf1
					}
					return true
				})(), 22)
				(_nw88).ArraySet1(!(_563_cf1), 23)
				(_nw88).ArraySet1(_563_cf1, 24)
				_570_v35 = _nw88
				var _571_v36 _dafny.CodePoint
				_ = _571_v36
				_571_v36 = _dafny.CodePoint('s')
				var _572_v37 _dafny.Set
				_ = _572_v37
				_572_v37 = _dafny.SetOf(_571_v36)
				var _nwElement0_19 bool = _563_cf1
				_ = _nwElement0_19
				var _nw89 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_19, nil, _dafny.IntOfInt64(15))
				_ = _nw89
				(_nw89).ArraySet1(_nwElement0_19, 0)
				(_nw89).ArraySet1(!(!((_this).F14())), 1)
				(_nw89).ArraySet1((_this).F14(), 2)
				(_nw89).ArraySet1(_563_cf1, 3)
				(_nw89).ArraySet1(false, 4)
				(_nw89).ArraySet1((_563_cf1) && ((_this).F14()), 5)
				(_nw89).ArraySet1(!((_this).F14()), 6)
				(_nw89).ArraySet1((true) == ((_this).F14()), 7)
				(_nw89).ArraySet1(_563_cf1, 8)
				(_nw89).ArraySet1(!(_563_cf1), 9)
				(_nw89).ArraySet1(true, 10)
				(_nw89).ArraySet1((_this).F14(), 11)
				(_nw89).ArraySet1(_563_cf1, 12)
				(_nw89).ArraySet1((_556_v25).Cmp((_572_v37).Cardinality()) < 0, 13)
				(_nw89).ArraySet1(false, 14)
				_570_v35 = _nw89
				var _573_v38 *C0
				_ = _573_v38
				var _nw90 *C0 = New_C0_()
				_ = _nw90
				_nw90.Ctor__(_563_cf1)
				_573_v38 = _nw90
				var _index107 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_570_v35), 0))
				_ = _index107
				(_570_v35).ArraySet1(_563_cf1, (_index107).Int())
				var _index108 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(394), _dafny.ArrayLen((_570_v35), 0))
				_ = _index108
				(_570_v35).ArraySet1(!(_563_cf1), (_index108).Int())
			} else {
				var _574_v39 _dafny.MultiSet
				_ = _574_v39
				_574_v39 = _dafny.MultiSetOf(_560_v29)
				var _575_v40 _dafny.Sequence
				_ = _575_v40
				_575_v40 = _dafny.SeqOf(Companion_D0_.Create_DC2_(_557_v26, _564_cf0, _563_cf1), _560_v29)
				var _576_v41 _dafny.Set
				_ = _576_v41
				_576_v41 = _dafny.SetOf(_574_v39, _574_v39, (_574_v39).Union(_dafny.MultiSetFromSeq(_575_v40)), _dafny.MultiSetOf(_560_v29), _dafny.MultiSetOf(_560_v29))
				var _577_v42 _dafny.Map
				_ = _577_v42
				_577_v42 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_563_cf1, _563_cf1)
				var _rhs78 _dafny.Set = _dafny.SetOf((_574_v39).Difference(_574_v39))
				_ = _rhs78
				var _rhs79 _dafny.Sequence = _507_v0
				_ = _rhs79
				var _rhs80 _dafny.Int = Companion_Default___.Fm0((_577_v42).Merge(_577_v42), (_563_cf1) || ((_this).F14()), _dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(158))).Uint32(), func(coer15 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg15 _dafny.Int) interface{} {
						return coer15(arg15)
					}
				}(func(_578_i3 _dafny.Int) D0 {
					return Companion_D0_.Create_DC3_(_dafny.CodePoint('k'))
				})), _dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-891))).Uint32(), func(coer16 func(_dafny.Int) D0) func(_dafny.Int) interface{} {
					return func(arg16 _dafny.Int) interface{} {
						return coer16(arg16)
					}
				}(func(_579_i4 _dafny.Int) D0 {
					return Companion_D0_.Create_DC3_(_dafny.CodePoint('m'))
				}))), globalState)
				_ = _rhs80
				var _rhs81 _dafny.Set = _559_v28
				_ = _rhs81
				var _rhs82 bool = (_this).F14()
				_ = _rhs82
				_576_v41 = _rhs78
				_507_v0 = _rhs79
				_564_cf0 = _rhs80
				_559_v28 = _rhs81
				_563_cf1 = _rhs82
				_563_cf1 = !((Companion_Default___.Fm0(_577_v42, !(!((_this).F14())), _563_cf1, globalState)).Cmp(Companion_Default___.Fm0(_577_v42, (_this).F14(), (_this).F14(), globalState)) == 0)
				_564_cf0 = Companion_Default___.SafeModuloInt((_this).F15(), _dafny.IntOfUint32((Companion_Default___.Fm8((_this).F14(), (_this).F14(), (_this).F14(), _563_cf1, globalState)).Cardinality()))
				_563_cf1 = (_this).F14()
				_564_cf0 = (_this).F15()
			}
			var _580_v43 _dafny.Array
			_ = _580_v43
			var _nw91 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(26))
			_ = _nw91
			_580_v43 = _nw91
			var _len0_12 _dafny.Int = _dafny.IntOfInt64(7)
			_ = _len0_12
			var _nw92 _dafny.Array
			_ = _nw92
			if _len0_12.Cmp(_dafny.Zero) == 0 {
				_nw92 = _dafny.NewArray(_len0_12)
			} else {
				var _init12 func(_dafny.Int) bool = func(_581_i5 _dafny.Int) bool {
					return (_this).F14()
				}
				_ = _init12
				var _element0_12 = _init12(_dafny.Zero)
				_ = _element0_12
				_nw92 = _dafny.NewArrayFromExample(_element0_12, nil, _len0_12)
				(_nw92).ArraySet1(_element0_12, 0)
				var _nativeLen0_12 = (_len0_12).Int()
				_ = _nativeLen0_12
				for _i0_12 := 1; _i0_12 < _nativeLen0_12; _i0_12++ {
					(_nw92).ArraySet1(_init12(_dafny.IntOf(_i0_12)), _i0_12)
				}
			}
			_580_v43 = _nw92
			var _582_v44 _dafny.Array
			_ = _582_v44
			var _nw93 _dafny.Array = _dafny.NewArrayWithValue(_dafny.Zero, _dafny.IntOfInt64(4))
			_ = _nw93
			_582_v44 = _nw93
			var _index109 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_582_v44), 0))
			_ = _index109
			(_582_v44).ArraySet1(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(-152))).Uint32(), func(coer17 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg17 _dafny.Int) interface{} {
					return coer17(arg17)
				}
			}((func(_583_v0 _dafny.Sequence, _584_v25 _dafny.Int) func(_dafny.Int) _dafny.CodePoint {
				return func(_585_i6 _dafny.Int) _dafny.CodePoint {
					return (_583_v0).Select((Companion_Default___.SafeIndex(_584_v25, _dafny.IntOfUint32((_583_v0).Cardinality()))).Uint32()).(_dafny.CodePoint)
				}
			})(_507_v0, _556_v25)))).Cardinality()), (_index109).Int())
			var _index110 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(24), _dafny.ArrayLen((_582_v44), 0))
			_ = _index110
			(_582_v44).ArraySet1((_dafny.Zero).Minus((func() _dafny.Int {
				if (_this).F14() {
					return _564_cf0
				}
				return (func() _dafny.Int {
					if _563_cf1 {
						return _dafny.IntOfUint32((_dafny.SeqOf((_this).F14())).Cardinality())
					}
					return _564_cf0
				})()
			})()), (_index110).Int())
		} else if _source11.Is_DC1() {
			var _586___mcc_h10 _dafny.Set = _source11.Get_().(D0_DC1).Cf2
			_ = _586___mcc_h10
			var _587___mcc_h11 bool = _source11.Get_().(D0_DC1).Cf3
			_ = _587___mcc_h11
			var _588_cf3 bool = _587___mcc_h11
			_ = _588_cf3
			var _589_cf2 _dafny.Set = _586___mcc_h10
			_ = _589_cf2
			var _590_v45 _dafny.CodePoint
			_ = _590_v45
			_590_v45 = _dafny.CodePoint('u')
			var _591_v46 *C0
			_ = _591_v46
			var _nw94 *C0 = New_C0_()
			_ = _nw94
			_nw94.Ctor__(_dafny.Companion_Sequence_.IsProperPrefixOf(_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(153))).Uint32(), func(coer18 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg18 _dafny.Int) interface{} {
					return coer18(arg18)
				}
			}(func(_592_i7 _dafny.Int) _dafny.CodePoint {
				return _dafny.CodePoint('x')
			})), _dafny.Companion_Sequence_.Update(_507_v0, (Companion_Default___.SafeIndex((_this).F15(), _dafny.IntOfUint32((_507_v0).Cardinality()))).Uint32(), _590_v45)))
			_591_v46 = _nw94
			var _593_v47 _dafny.Array
			_ = _593_v47
			var _nw95 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(10))
			_ = _nw95
			_593_v47 = _nw95
			_593_v47 = _593_v47
			var _index111 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_593_v47), 0))
			_ = _index111
			(_593_v47).ArraySet1((_591_v46).F16(), (_index111).Int())
			var _index112 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(591), _dafny.ArrayLen((_593_v47), 0))
			_ = _index112
			(_593_v47).ArraySet1((_591_v46).F16(), (_index112).Int())
			var _594_v48 T0
			_ = _594_v48
			var _nw96 *C0 = New_C0_()
			_ = _nw96
			_nw96.Ctor__((_591_v46).F16())
			_594_v48 = _nw96
			var _595_v49 _dafny.Sequence
			_ = _595_v49
			_595_v49 = _dafny.SeqOf(_594_v48)
			_595_v49 = _595_v49
		} else if _source11.Is_DC2() {
			var _596___mcc_h12 _dafny.Set = _source11.Get_().(D0_DC2).Cf4
			_ = _596___mcc_h12
			var _597___mcc_h13 _dafny.Int = _source11.Get_().(D0_DC2).Cf5
			_ = _597___mcc_h13
			var _598___mcc_h14 bool = _source11.Get_().(D0_DC2).Cf6
			_ = _598___mcc_h14
			var _599_cf6 bool = _598___mcc_h14
			_ = _599_cf6
			var _600_cf5 _dafny.Int = _597___mcc_h13
			_ = _600_cf5
			var _601_cf4 _dafny.Set = _596___mcc_h12
			_ = _601_cf4
			var _602_v50 _dafny.Sequence
			_ = _602_v50
			_602_v50 = _dafny.SeqOf(_556_v25, (_557_v26).Cardinality(), _600_cf5, _600_cf5)
			var _603_v51 _dafny.Sequence
			_ = _603_v51
			_603_v51 = _dafny.SeqOf(false, Companion_Default___.Fm4(!((_this).F14()), _559_v28, (_this).F15(), globalState))
			var _604_v52 _dafny.Map
			_ = _604_v52
			_604_v52 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _603_v51)
			var _605_v53 _dafny.Array
			_ = _605_v53
			var _nwElement0_20 _dafny.Int = (_this).F15()
			_ = _nwElement0_20
			var _nw97 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_20, nil, _dafny.IntOfInt64(15))
			_ = _nw97
			(_nw97).ArraySet1(_nwElement0_20, 0)
			(_nw97).ArraySet1(_600_cf5, 1)
			(_nw97).ArraySet1(_600_cf5, 2)
			(_nw97).ArraySet1((func() _dafny.Int {
				if (_this).F14() {
					return _556_v25
				}
				return (_this).F15()
			})(), 3)
			(_nw97).ArraySet1(_556_v25, 4)
			(_nw97).ArraySet1(_dafny.IntOfInt64(918), 5)
			(_nw97).ArraySet1((_this).F15(), 6)
			(_nw97).ArraySet1(_dafny.IntOfInt64(718), 7)
			(_nw97).ArraySet1((_556_v25).Minus(_600_cf5), 8)
			(_nw97).ArraySet1(_dafny.IntOfUint32((_602_v50).Cardinality()), 9)
			(_nw97).ArraySet1((_this).F15(), 10)
			(_nw97).ArraySet1((_dafny.Zero).Minus(_556_v25), 11)
			(_nw97).ArraySet1(_dafny.IntOfInt64(67), 12)
			(_nw97).ArraySet1(((_604_v52).Cardinality()).Minus(_556_v25), 13)
			(_nw97).ArraySet1(_556_v25, 14)
			_605_v53 = _nw97
			var _index113 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_605_v53), 0))
			_ = _index113
			(_605_v53).ArraySet1((_this).F15(), (_index113).Int())
			var _606_v54 _dafny.Array
			_ = _606_v54
			var _nw98 _dafny.Array = _dafny.NewArrayWithValue(_dafny.EmptySeq, _dafny.IntOfInt64(15))
			_ = _nw98
			_606_v54 = _nw98
			var _index114 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))
			_ = _index114
			(_606_v54).ArraySet1(_507_v0, (_index114).Int())
			var _607_v55 _dafny.CodePoint
			_ = _607_v55
			_607_v55 = _dafny.CodePoint('d')
			var _608_v56 _dafny.MultiSet
			_ = _608_v56
			_608_v56 = _dafny.MultiSetOf(_607_v55)
			var _609_v57 _dafny.Map
			_ = _609_v57
			_609_v57 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), _556_v25)
			var _index115 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_605_v53), 0))
			_ = _index115
			var _index116 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))
			_ = _index116
			var _rhs83 bool = (_608_v56).IsDisjointFrom((_608_v56).Update(_607_v55, Companion_Default___.Abs(_600_cf5)))
			_ = _rhs83
			var _rhs84 _dafny.Set = (Companion_Default___.Fm9(_607_v55, (_this).F15(), _dafny.IntOfUint32((_dafny.SeqOf(_599_cf6)).Cardinality()), _609_v57, globalState)).Difference(_559_v28)
			_ = _rhs84
			var _rhs85 _dafny.Int = (_dafny.Zero).Minus(_dafny.IntOfUint32((_dafny.Companion_Sequence_.Concatenate(_dafny.SeqOf(_507_v0, _507_v0), _dafny.Companion_Sequence_.Concatenate(_509_v2, _509_v2))).Cardinality()))
			_ = _rhs85
			var _rhs86 _dafny.Sequence = _507_v0
			_ = _rhs86
			var _lhs44 _dafny.Array = _605_v53
			_ = _lhs44
			var _lhs45 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_605_v53), 0))
			_ = _lhs45
			var _lhs46 _dafny.Array = _606_v54
			_ = _lhs46
			var _lhs47 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))
			_ = _lhs47
			_599_cf6 = _rhs83
			_559_v28 = _rhs84
			(_lhs44).ArraySet1(_rhs85, (_lhs45).Int())
			(_lhs46).ArraySet1(_rhs86, (_lhs47).Int())
			var _610_v58 _dafny.Map
			_ = _610_v58
			_610_v58 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_605_v53, (_this).F14())
			var _611_v59 _dafny.Array
			_ = _611_v59
			var _nwElement0_21 bool = _599_cf6
			_ = _nwElement0_21
			var _nw99 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_21, nil, _dafny.IntOfInt64(23))
			_ = _nw99
			(_nw99).ArraySet1(_nwElement0_21, 0)
			(_nw99).ArraySet1((_this).F14(), 1)
			(_nw99).ArraySet1((_this).F14(), 2)
			(_nw99).ArraySet1((func() bool {
				if (_610_v58).Contains(_605_v53) {
					return (_610_v58).Get(_605_v53).(bool)
				}
				return !(_599_cf6)
			})(), 3)
			(_nw99).ArraySet1(!(true), 4)
			(_nw99).ArraySet1(_599_cf6, 5)
			(_nw99).ArraySet1(Companion_Default___.Fm4(true, _559_v28, _dafny.IntOfUint32(((_606_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))).Int()).(_dafny.Sequence)).Cardinality()), globalState), 6)
			(_nw99).ArraySet1(!(!(!_dafny.Companion_Sequence_.Contains(_602_v50, _600_cf5))), 7)
			(_nw99).ArraySet1(_599_cf6, 8)
			(_nw99).ArraySet1((_this).F14(), 9)
			(_nw99).ArraySet1(false, 10)
			(_nw99).ArraySet1(_599_cf6, 11)
			(_nw99).ArraySet1((_this).F14(), 12)
			(_nw99).ArraySet1((_600_cf5).Cmp((_605_v53).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_605_v53), 0))).Int()).(_dafny.Int)) == 0, 13)
			(_nw99).ArraySet1(!_dafny.Companion_Sequence_.Equal((_606_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))).Int()).(_dafny.Sequence), (_606_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))).Int()).(_dafny.Sequence)), 14)
			(_nw99).ArraySet1((_this).F14(), 15)
			(_nw99).ArraySet1((_this).F14(), 16)
			(_nw99).ArraySet1((func() bool {
				if _599_cf6 {
					return _599_cf6
				}
				return (_this).F14()
			})(), 17)
			(_nw99).ArraySet1(!(_599_cf6) || ((_this).F14()), 18)
			(_nw99).ArraySet1((_this).F14(), 19)
			(_nw99).ArraySet1(false, 20)
			(_nw99).ArraySet1(_599_cf6, 21)
			(_nw99).ArraySet1(_599_cf6, 22)
			_611_v59 = _nw99
			var _index117 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_611_v59), 0))
			_ = _index117
			(_611_v59).ArraySet1((_this).F14(), (_index117).Int())
			var _index118 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_611_v59), 0))
			_ = _index118
			(_611_v59).ArraySet1(Companion_Default___.Fm4((_this).F14(), _559_v28, (_this).F15(), globalState), (_index118).Int())
			var _612_v60 D2
			_ = _612_v60
			_612_v60 = Companion_D2_.Create_DC6_(_dafny.SeqOf(_599_cf6))
			var _613_v61 _dafny.MultiSet
			_ = _613_v61
			_613_v61 = _dafny.MultiSetOf(!((_this).F14()))
			var _index119 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_605_v53), 0))
			_ = _index119
			var _index120 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_611_v59), 0))
			_ = _index120
			var _index121 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_611_v59), 0))
			_ = _index121
			var _rhs87 D0 = func(_pat_let14_0 D0) D0 {
				return func(_614_dt__update__tmp_h1 D0) D0 {
					return func(_pat_let15_0 bool) D0 {
						return func(_615_dt__update_hcf6_h0 bool) D0 {
							return func(_pat_let16_0 _dafny.Int) D0 {
								return func(_616_dt__update_hcf5_h0 _dafny.Int) D0 {
									return Companion_D0_.Create_DC2_((_614_dt__update__tmp_h1).Dtor_cf4(), _616_dt__update_hcf5_h0, _615_dt__update_hcf6_h0)
								}(_pat_let16_0)
							}(_dafny.IntOfInt64(360))
						}(_pat_let15_0)
					}(!((_this).F14()) || ((_this).F14()))
				}(_pat_let14_0)
			}(_560_v29)
			_ = _rhs87
			var _rhs88 _dafny.Int = (_this).F15()
			_ = _rhs88
			var _rhs89 bool = false
			_ = _rhs89
			var _rhs90 bool = true
			_ = _rhs90
			var _rhs91 bool = (func() bool {
				if (_dafny.MultiSetFromSeq((_612_v60).Dtor_cf10())).IsProperSubsetOf(_613_v61) {
					return (_613_v61).IsProperSubsetOf(_613_v61)
				}
				return !(false)
			})()
			_ = _rhs91
			var _lhs48 _dafny.Array = _605_v53
			_ = _lhs48
			var _lhs49 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(241), _dafny.ArrayLen((_605_v53), 0))
			_ = _lhs49
			var _lhs50 _dafny.Array = _611_v59
			_ = _lhs50
			var _lhs51 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_611_v59), 0))
			_ = _lhs51
			var _lhs52 _dafny.Array = _611_v59
			_ = _lhs52
			var _lhs53 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(445), _dafny.ArrayLen((_611_v59), 0))
			_ = _lhs53
			_560_v29 = _rhs87
			(_lhs48).ArraySet1(_rhs88, (_lhs49).Int())
			_599_cf6 = _rhs89
			(_lhs50).ArraySet1(_rhs90, (_lhs51).Int())
			(_lhs52).ArraySet1(_rhs91, (_lhs53).Int())
			var _617_v62 _dafny.Array
			_ = _617_v62
			var _nw100 _dafny.Array = _dafny.NewArray(_dafny.IntOfInt64(16))
			_ = _nw100
			_617_v62 = _nw100
			var _618_v63 *C0
			_ = _618_v63
			var _nw101 *C0 = New_C0_()
			_ = _nw101
			_nw101.Ctor__(!(true))
			_618_v63 = _nw101
			var _index122 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_617_v62), 0))
			_ = _index122
			(_617_v62).ArraySet1(_618_v63, (_index122).Int())
			var _index123 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(851), _dafny.ArrayLen((_617_v62), 0))
			_ = _index123
			(_617_v62).ArraySet1(_618_v63, (_index123).Int())
			var _619_v64 _dafny.Map
			_ = _619_v64
			_619_v64 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_611_v59).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(274), _dafny.ArrayLen((_611_v59), 0))).Int()).(bool), (_606_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))).Int()).(_dafny.Sequence))
			_600_cf5 = (((_619_v64).Update(_599_cf6, (_606_v54).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(678), _dafny.ArrayLen((_606_v54), 0))).Int()).(_dafny.Sequence))).Merge((_619_v64).Merge(_619_v64))).Cardinality()
		} else {
			var _620___mcc_h15 _dafny.CodePoint = _source11.Get_().(D0_DC3).Cf7
			_ = _620___mcc_h15
			var _621_cf7 _dafny.CodePoint = _620___mcc_h15
			_ = _621_cf7
			var _622_v65 D2
			_ = _622_v65
			_622_v65 = Companion_D2_.Create_DC7_(Companion_Default___.Fm6((_this).F14(), _556_v25, globalState), (_this).F14(), (_dafny.IntOfUint32((_dafny.UnicodeSeqOfUtf8Bytes("lawxj")).Cardinality())).Cmp(_dafny.IntOfInt64(819)) != 0)
			_622_v65 = _622_v65
			if (Companion_Default___.SafeDivisionInt(_556_v25, _556_v25)).Cmp((_this).F15()) != 0 {
				var _623_v66 *C0
				_ = _623_v66
				var _nw102 *C0 = New_C0_()
				_ = _nw102
				_nw102.Ctor__((_this).F14())
				_623_v66 = _nw102
				var _624_v67 _dafny.Sequence
				_ = _624_v67
				_624_v67 = _dafny.SeqOf((_this).F14(), !(!((_this).F14()) || ((_623_v66).F16())), (_623_v66).F16(), ((_dafny.Zero).Minus((_this).F15())).Cmp((_this).F15()) != 0)
				var _625_v68 _dafny.Array
				_ = _625_v68
				var _nw103 _dafny.Array = _dafny.NewArrayWithValue(false, _dafny.IntOfInt64(20))
				_ = _nw103
				_625_v68 = _nw103
				var _index124 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_625_v68), 0))
				_ = _index124
				(_625_v68).ArraySet1((_this).F14(), (_index124).Int())
				var _index125 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_625_v68), 0))
				_ = _index125
				var _rhs92 _dafny.Sequence = _624_v67
				_ = _rhs92
				var _rhs93 bool = ((_this).F15()).Cmp(_556_v25) <= 0
				_ = _rhs93
				var _rhs94 *C0 = _623_v66
				_ = _rhs94
				var _rhs95 _dafny.Array = _625_v68
				_ = _rhs95
				var _lhs54 _dafny.Array = _625_v68
				_ = _lhs54
				var _lhs55 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_625_v68), 0))
				_ = _lhs55
				_624_v67 = _rhs92
				(_lhs54).ArraySet1(_rhs93, (_lhs55).Int())
				_623_v66 = _rhs94
				_625_v68 = _rhs95
				var _626_v69 *C0
				_ = _626_v69
				var _nw104 *C0 = New_C0_()
				_ = _nw104
				_nw104.Ctor__((_625_v68).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(777), _dafny.ArrayLen((_625_v68), 0))).Int()).(bool))
				_626_v69 = _nw104
				var _627_v70 *C0
				_ = _627_v70
				var _nw105 *C0 = New_C0_()
				_ = _nw105
				_nw105.Ctor__(_dafny.Companion_Sequence_.IsPrefixOf((_626_v69).Fm5(false, _556_v25, _dafny.IntOfInt64(43), (_this).F14(), globalState), _507_v0))
				_627_v70 = _nw105
				var _628_v71 D0
				_ = _628_v71
				_628_v71 = Companion_D0_.Create_DC0_((_this).F15(), (_626_v69).F16())
				var _629_v72 *C0
				_ = _629_v72
				var _nw106 *C0 = New_C0_()
				_ = _nw106
				_nw106.Ctor__(((_628_v71).Dtor_cf1()) || (Companion_Default___.Fm6((_627_v70).F16(), _556_v25, globalState)))
				_629_v72 = _nw106
			} else {
				_556_v25 = (func() _dafny.Int {
					if (_this).F14() {
						return (_this).F15()
					}
					return Companion_Default___.SafeModuloInt(_556_v25, _556_v25)
				})()
				var _630_v73 _dafny.Map
				_ = _630_v73
				_630_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v25, _556_v25)
				_630_v73 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F15(), (_this).F15())
				var _631_v74 *C0
				_ = _631_v74
				var _nw107 *C0 = New_C0_()
				_ = _nw107
				_nw107.Ctor__(false)
				_631_v74 = _nw107
				var _632_v75 _dafny.Sequence
				_ = _632_v75
				_632_v75 = _dafny.SeqOf((_this).F14(), Companion_Default___.Fm4((_631_v74).F16(), _559_v28, _dafny.IntOfInt64(231), globalState))
				var _633_v76 bool
				_ = _633_v76
				var _634_v77 _dafny.Array
				_ = _634_v77
				var _635_v78 _dafny.Int
				_ = _635_v78
				var _636_v79 _dafny.Int
				_ = _636_v79
				var _out44 bool
				_ = _out44
				var _out45 _dafny.Array
				_ = _out45
				var _out46 _dafny.Int
				_ = _out46
				var _out47 _dafny.Int
				_ = _out47
				_out44, _out45, _out46, _out47 = Companion_Default___.M0(_556_v25, _632_v75, globalState)
				_633_v76 = _out44
				_634_v77 = _out45
				_635_v78 = _out46
				_636_v79 = _out47
				var _637_v80 D0
				_ = _637_v80
				_637_v80 = Companion_D0_.Create_DC3_(_621_cf7)
				_637_v80 = _637_v80
			}
			var _638_v81 _dafny.Map
			_ = _638_v81
			_638_v81 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F15())
			var _639_v82 _dafny.Map
			_ = _639_v82
			_639_v82 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe((_this).F14(), (_this).F14())
			var _640_v83 *C0
			_ = _640_v83
			var _nw108 *C0 = New_C0_()
			_ = _nw108
			_nw108.Ctor__((_this).F14())
			_640_v83 = _nw108
			var _641_v84 _dafny.Map
			_ = _641_v84
			_641_v84 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(_556_v25, _640_v83)
			var _642_v85 _dafny.Array
			_ = _642_v85
			var _nwElement0_22 bool = !((_this).F14()) || ((_this).F14())
			_ = _nwElement0_22
			var _nw109 _dafny.Array = _dafny.NewArrayFromExample(_nwElement0_22, nil, _dafny.IntOfInt64(16))
			_ = _nw109
			(_nw109).ArraySet1(_nwElement0_22, 0)
			(_nw109).ArraySet1((_this).F14(), 1)
			(_nw109).ArraySet1((_this).F14(), 2)
			(_nw109).ArraySet1(!((_this).F14()) || (Companion_Default___.Fm4((_this).F14(), _559_v28, (_638_v81).Cardinality(), globalState)), 3)
			(_nw109).ArraySet1((_this).F14(), 4)
			(_nw109).ArraySet1(!((_this).F14()), 5)
			(_nw109).ArraySet1((func() bool {
				if (_639_v82).Contains((_this).F14()) {
					return (_639_v82).Get((_this).F14()).(bool)
				}
				return (_this).F14()
			})(), 6)
			(_nw109).ArraySet1((_this).F14(), 7)
			(_nw109).ArraySet1((_this).F14(), 8)
			(_nw109).ArraySet1((_this).F14(), 9)
			(_nw109).ArraySet1((_this).F14(), 10)
			(_nw109).ArraySet1((_this).F14(), 11)
			(_nw109).ArraySet1(((_641_v84).Cardinality()).Cmp(_dafny.IntOfInt64(-202)) <= 0, 12)
			(_nw109).ArraySet1((_this).F14(), 13)
			(_nw109).ArraySet1(((_dafny.Zero).Minus((_this).F15())).Cmp(_dafny.IntOfUint32((_dafny.SeqCreate((Companion_Default___.Abs(_dafny.IntOfInt64(438))).Uint32(), func(coer19 func(_dafny.Int) _dafny.CodePoint) func(_dafny.Int) interface{} {
				return func(arg19 _dafny.Int) interface{} {
					return coer19(arg19)
				}
			}((func(_643_cf7 _dafny.CodePoint) func(_dafny.Int) _dafny.CodePoint {
				return func(_644_i8 _dafny.Int) _dafny.CodePoint {
					return _643_cf7
				}
			})(_621_cf7)))).Cardinality())) > 0, 14)
			(_nw109).ArraySet1(!((_640_v83).F16()), 15)
			_642_v85 = _nw109
			var _index126 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_642_v85), 0))
			_ = _index126
			(_642_v85).ArraySet1((_this).F14(), (_index126).Int())
			var _index127 _dafny.Int = Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_642_v85), 0))
			_ = _index127
			(_642_v85).ArraySet1((_this).F14(), (_index127).Int())
			var _645_v86 _dafny.Map
			_ = _645_v86
			_645_v86 = _dafny.NewMapBuilder().ToMap().UpdateUnsafe(Companion_Default___.Fm4(true, _559_v28, _dafny.IntOfInt64(-291), globalState), (func() _dafny.Sequence {
				if (_642_v85).ArrayGet1((Companion_Default___.SafeIndex(_dafny.IntOfInt64(51), _dafny.ArrayLen((_642_v85), 0))).Int()).(bool) {
					return _dafny.UnicodeSeqOfUtf8Bytes("kyo")
				}
				return _507_v0
			})())
			_645_v86 = (_645_v86).Update((_this).F14(), _dafny.Companion_Sequence_.Concatenate(_507_v0, _507_v0))
		}
		var _646_v87 bool
		_ = _646_v87
		_646_v87 = false
		var _647_v88 _dafny.Array
		_ = _647_v88
		var _len0_13 _dafny.Int = _dafny.IntOfInt64(23)
		_ = _len0_13
		var _nw110 _dafny.Array
		_ = _nw110
		if _len0_13.Cmp(_dafny.Zero) == 0 {
			_nw110 = _dafny.NewArray(_len0_13)
		} else {
			var _init13 func(_dafny.Int) bool = (func(_648_v87 bool) func(_dafny.Int) bool {
				return func(_649_i9 _dafny.Int) bool {
					return _648_v87
				}
			})(_646_v87)
			_ = _init13
			var _element0_13 = _init13(_dafny.Zero)
			_ = _element0_13
			_nw110 = _dafny.NewArrayFromExample(_element0_13, nil, _len0_13)
			(_nw110).ArraySet1(_element0_13, 0)
			var _nativeLen0_13 = (_len0_13).Int()
			_ = _nativeLen0_13
			for _i0_13 := 1; _i0_13 < _nativeLen0_13; _i0_13++ {
				(_nw110).ArraySet1(_init13(_dafny.IntOf(_i0_13)), _i0_13)
			}
		}
		_647_v88 = _nw110
		var _650_v89 _dafny.MultiSet
		_ = _650_v89
		_650_v89 = _dafny.MultiSetOf(_647_v88, _647_v88)
		var _651_v90 _dafny.MultiSet
		_ = _651_v90
		_651_v90 = _dafny.MultiSetOf((_this).Fm1(_507_v0, (_this).F15(), _556_v25, globalState))
		_646_v87 = ((_650_v89).Update(_647_v88, Companion_Default___.Abs((_dafny.Zero).Minus((_651_v90).Cardinality())))).IsSubsetOf(_650_v89)
	}
}
func (_this *C1) F14() bool {
	{
		return _this._f14
	}
}
func (_this *C1) F15() _dafny.Int {
	{
		return _this._f15
	}
}

// End of class C1
func main() {
	defer _dafny.CatchHalt()
	Companion_Default___.Main(_dafny.UnicodeFromMainArguments(os.Args))
}
