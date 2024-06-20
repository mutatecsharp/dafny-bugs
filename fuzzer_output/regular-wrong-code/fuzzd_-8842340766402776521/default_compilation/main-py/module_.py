import sys
from typing import Callable, Any, TypeVar, NamedTuple
from math import floor
from itertools import count

import module_
import _dafny
import System_

# Module: module_

class default__:
    def  __init__(self):
        pass

    @staticmethod
    def abs(x):
        if (x) < (0):
            return (-1) * (x)
        elif True:
            return x

    @staticmethod
    def safeIndex(x, length):
        if (x) < (0):
            return 0
        elif (x) >= (length):
            return _dafny.euclidian_modulus(x, length)
        elif True:
            return x

    @staticmethod
    def safeDivisionInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_division(x1, x2)

    @staticmethod
    def safeModuloInt(x1, x2):
        if (x2) == (0):
            return x1
        elif True:
            return _dafny.euclidian_modulus(x1, x2)

    @staticmethod
    def fm0(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(194, 317):
                d_0_v0_: int = compr_0_
                if ((194) <= (d_0_v0_)) and ((d_0_v0_) < (317)):
                    coll0_[(d_0_v0_) - (-949)] = _dafny.Set({False})
            return _dafny.Map(coll0_)
        return (D1_DC2(len(_dafny.Map({False: 898})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mcftfnqg"))), -391, len(iife0_()
))).cf3

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return (_dafny.Set({False})).isdisjoint((_dafny.Set({True})) - (_dafny.Set({True, False, False})))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))): False})) | (_dafny.Map({len(_dafny.Set({False, True, not(not(False)), not(False)})): False}))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        if (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1_i0_ in range(default__.abs(807))]))) == (len(_dafny.Set({not(not(not(False))), not(not(not(not(True)))), not(True), False}))):
            return _dafny.CodePoint('d')
        elif True:
            return _dafny.CodePoint('g')

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-994 for d_2_i0_ in range(default__.abs(-685))])): True})).keys.Elements:
                d_3_v0_: int = compr_1_
                if (d_3_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-994 for d_2_i0_ in range(default__.abs(-685))])): True})):
                    coll1_ = coll1_.union(_dafny.Set([(d_3_v0_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m'), _dafny.CodePoint('w'), _dafny.CodePoint('k'), _dafny.CodePoint('b'), _dafny.CodePoint('g')]) for d_4_i1_ in range(default__.abs(79))])))]))
            return _dafny.Set(coll1_)
        return (_dafny.Map({_dafny.Set({759}): len(_dafny.Set({True, False}))})) | (_dafny.Map({iife1_()
        : 930}))

    @staticmethod
    def fm7(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iptpuc")))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_5_i0_ in range(default__.abs(838))]))

    @staticmethod
    def fm8(p0, globalState):
        source0_ = D1_DC2(243, len(_dafny.Map({len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adlori"))})): D3_DC9(True)})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omdgnphda"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gtoykysed"))))
        if source0_.is_DC2:
            d_6___mcc_h0_ = source0_.cf2
            d_7___mcc_h1_ = source0_.cf3
            d_8___mcc_h2_ = source0_.cf4
            d_9___mcc_h3_ = source0_.cf5
            d_10_cf5_ = d_9___mcc_h3_
            d_11_cf4_ = d_8___mcc_h2_
            d_12_cf3_ = d_7___mcc_h1_
            d_13_cf2_ = d_6___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([not(False), True])) + (_dafny.SeqWithoutIsStrInference([not(True)]))
        elif source0_.is_DC1:
            d_14___mcc_h4_ = source0_.cf1
            d_15_cf1_ = d_14___mcc_h4_
            return (_dafny.SeqWithoutIsStrInference([False, False])) + (_dafny.SeqWithoutIsStrInference([not(True), True]))
        elif True:
            d_16___mcc_h5_ = source0_.cf6
            d_17_cf6_ = d_16___mcc_h5_
            return _dafny.SeqWithoutIsStrInference([True, False, False])

    @staticmethod
    def m0(p0, p1, globalState):
        (globalState).f10 = (p1) * ((p1) * (p1))
        d_18_v0_: C0
        nw0_ = C0()
        nw0_.ctor__(p1)
        d_18_v0_ = nw0_
        d_19_v1_: _dafny.Map
        d_19_v1_ = _dafny.Map({d_18_v0_: (d_18_v0_).f14})
        d_20_v2_: _dafny.Seq
        d_20_v2_ = _dafny.SeqWithoutIsStrInference([p0])
        d_21_v3_: _dafny.Seq
        d_21_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iaqmfjp"))
        d_22_v4_: _dafny.MultiSet
        d_22_v4_ = _dafny.MultiSet([(d_20_v2_)[default__.safeIndex((d_18_v0_).f14, len(d_20_v2_))], p0, p0])
        d_23_v5_: _dafny.Array
        nw1_ = _dafny.Array(None, 12)
        nw1_[int(0)] = (d_18_v0_) in (d_19_v1_)
        nw1_[int(1)] = (d_20_v2_)[default__.safeIndex(len(d_21_v3_), len(d_20_v2_))]
        nw1_[int(2)] = p0
        nw1_[int(3)] = not((p0) or (p0))
        nw1_[int(4)] = p0
        nw1_[int(5)] = p0
        nw1_[int(6)] = p0
        nw1_[int(7)] = p0
        nw1_[int(8)] = ((0) - (len(d_21_v3_))) != (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpjgtvwk"))))
        nw1_[int(9)] = p0
        nw1_[int(10)] = (d_22_v4_).ispropersubset(d_22_v4_)
        nw1_[int(11)] = (d_18_v0_).fm4(p0, globalState)
        d_23_v5_ = nw1_
        index0_ = default__.safeIndex(212, (d_23_v5_).length(0))
        (d_23_v5_)[index0_] = True
        index1_ = default__.safeIndex(212, (d_23_v5_).length(0))
        (d_23_v5_)[index1_] = p0
        d_24_v6_: C0
        nw2_ = C0()
        nw2_.ctor__(default__.safeDivisionInt(p1, p1))
        d_24_v6_ = nw2_
        d_25_v7_: _dafny.Map
        d_25_v7_ = _dafny.Map({p0: len(d_20_v2_)})
        (globalState).f10 = (((d_25_v7_)[p0] if (p0) in (d_25_v7_) else p1)) - ((d_18_v0_).f14)
        d_26_v8_: _dafny.Seq
        d_26_v8_ = _dafny.SeqWithoutIsStrInference([d_24_v6_, d_18_v0_])
        d_26_v8_ = d_26_v8_
        if not((d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]):
            (globalState).f6 = (d_24_v6_).f14
            (globalState).f10 = p1
            d_25_v7_ = (d_25_v7_).set(p0, 800)
            d_27_v9_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 5)
            d_27_v9_ = nw3_
            index2_ = default__.safeIndex(881, (d_27_v9_).length(0))
            (d_27_v9_)[index2_] = p1
            d_28_v10_: _dafny.Map
            d_28_v10_ = _dafny.Map({p0: d_23_v5_})
            index3_ = default__.safeIndex(881, (d_27_v9_).length(0))
            index4_ = default__.safeIndex(212, (d_23_v5_).length(0))
            rhs0_ = d_21_v3_
            rhs1_ = (d_24_v6_).f14
            rhs2_ = len(d_28_v10_)
            rhs3_ = not(p0)
            lhs0_ = globalState
            lhs1_ = d_27_v9_
            lhs2_ = default__.safeIndex(881, (d_27_v9_).length(0))
            lhs3_ = d_23_v5_
            lhs4_ = default__.safeIndex(212, (d_23_v5_).length(0))
            d_21_v3_ = rhs0_
            lhs0_.f10 = rhs1_
            lhs1_[lhs2_] = rhs2_
            lhs3_[lhs4_] = rhs3_
            d_29_v11_: _dafny.Map
            d_29_v11_ = _dafny.Map({(d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]: _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]: (d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]}) for d_30_i0_ in range(default__.abs(823))])})
            d_31_v12_: _dafny.Map
            d_31_v12_ = _dafny.Map({(d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]: p0})
            d_32_v13_: _dafny.Seq
            d_32_v13_ = _dafny.SeqWithoutIsStrInference([d_31_v12_, d_31_v12_])
            d_29_v11_ = (d_29_v11_).set((d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))], d_32_v13_)
        elif True:
            d_33_v14_: _dafny.Map
            d_33_v14_ = _dafny.Map({713: len(_dafny.Map({(d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]: (d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]}))})
            d_34_v15_: D1
            d_34_v15_ = D1_DC1(d_33_v14_)
            pat_let_tv0_ = d_33_v14_
            def iife2_(_pat_let0_0):
                def iife3_(d_35_dt__update__tmp_h0_):
                    def iife4_(_pat_let1_0):
                        def iife5_(d_36_dt__update_hcf1_h0_):
                            return D1_DC1(d_36_dt__update_hcf1_h0_)
                        return iife5_(_pat_let1_0)
                    return iife4_(pat_let_tv0_)
                return iife3_(_pat_let0_0)
            source1_ = iife2_(d_34_v15_)
            if source1_.is_DC2:
                d_37___mcc_h0_ = source1_.cf2
                d_38___mcc_h1_ = source1_.cf3
                d_39___mcc_h2_ = source1_.cf4
                d_40___mcc_h3_ = source1_.cf5
                d_41_cf5_ = d_40___mcc_h3_
                d_42_cf4_ = d_39___mcc_h2_
                d_43_cf3_ = d_38___mcc_h1_
                d_44_cf2_ = d_37___mcc_h0_
                d_45_v16_: _dafny.MultiSet
                d_45_v16_ = _dafny.MultiSet([d_21_v3_])
                d_46_v17_: _dafny.Map
                d_46_v17_ = _dafny.Map({(d_45_v16_).cardinality: d_21_v3_})
                d_21_v3_ = ((d_46_v17_)[(d_18_v0_).f14] if ((d_18_v0_).f14) in (d_46_v17_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_47_i1_ in range(default__.abs(166))]))
                d_48_v18_: C0
                nw4_ = C0()
                nw4_.ctor__((d_18_v0_).f14)
                d_48_v18_ = nw4_
                index5_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index5_] = (d_42_cf4_) > (d_42_cf4_)
                d_43_cf3_ = (d_24_v6_).f14
            elif source1_.is_DC1:
                d_49___mcc_h4_ = source1_.cf1
                d_50_cf1_ = d_49___mcc_h4_
                index6_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index6_] = (((0) - ((d_24_v6_).f14)) + ((0) - (len(d_21_v3_)))) <= (((d_25_v7_)[p0] if (p0) in (d_25_v7_) else (d_24_v6_).f14))
                d_51_v19_: _dafny.Array
                nw5_ = _dafny.Array(int(0), 6)
                d_51_v19_ = nw5_
                index7_ = default__.safeIndex(236, (d_51_v19_).length(0))
                (d_51_v19_)[index7_] = (d_18_v0_).f14
                index8_ = default__.safeIndex(236, (d_51_v19_).length(0))
                rhs4_ = default__.safeModuloInt((0) - (len(d_21_v3_)), (d_24_v6_).f14)
                rhs5_ = (74) - (p1)
                rhs6_ = default__.safeModuloInt((d_18_v0_).f14, p1)
                lhs5_ = globalState
                lhs6_ = d_51_v19_
                lhs7_ = default__.safeIndex(236, (d_51_v19_).length(0))
                lhs8_ = globalState
                lhs5_.f10 = rhs4_
                lhs6_[lhs7_] = rhs5_
                lhs8_.f10 = rhs6_
                d_52_v20_: str
                d_52_v20_ = _dafny.CodePoint('c')
                d_52_v20_ = d_52_v20_
                d_51_v19_ = d_51_v19_
            elif True:
                d_53___mcc_h5_ = source1_.cf6
                d_54_cf6_ = d_53___mcc_h5_
                (globalState).f10 = 568
                index9_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index9_] = (d_24_v6_).fm5(globalState)
                d_55_v21_: C0
                nw6_ = C0()
                nw6_.ctor__((d_18_v0_).f14)
                d_55_v21_ = nw6_
                d_56_v22_: _dafny.Array
                nw7_ = _dafny.Array(_dafny.Map({}), 16)
                d_56_v22_ = nw7_
                d_57_v23_: _dafny.Map
                d_57_v23_ = _dafny.Map({not(p0): (d_20_v2_).set(default__.safeIndex((d_55_v21_).f14, len(d_20_v2_)), (d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))])})
                index10_ = default__.safeIndex(241, (d_56_v22_).length(0))
                (d_56_v22_)[index10_] = d_57_v23_
                d_58_v24_: str
                d_58_v24_ = _dafny.CodePoint('d')
                d_59_v25_: _dafny.Map
                d_59_v25_ = _dafny.Map({_dafny.Map({p0: len(_dafny.SeqWithoutIsStrInference([p0]))}): d_55_v21_})
                d_60_v26_: _dafny.Map
                d_60_v26_ = _dafny.Map({d_58_v24_: ((d_59_v25_)[d_25_v7_] if (d_25_v7_) in (d_59_v25_) else d_55_v21_)})
                index11_ = default__.safeIndex(241, (d_56_v22_).length(0))
                rhs7_ = (d_24_v6_).f14
                rhs8_ = d_57_v23_
                rhs9_ = (d_60_v26_) | (_dafny.Map({d_58_v24_: d_18_v0_}))
                lhs9_ = globalState
                lhs10_ = d_56_v22_
                lhs11_ = default__.safeIndex(241, (d_56_v22_).length(0))
                lhs9_.f10 = rhs7_
                lhs10_[lhs11_] = rhs8_
                d_60_v26_ = rhs9_
            d_61_v27_: _dafny.Map
            d_61_v27_ = _dafny.Map({p1: (d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]})
            (globalState).f10 = (len(d_61_v27_)) - ((0) - (p1))
            if not((d_21_v3_) == (d_21_v3_)):
                index12_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index12_] = False
                d_62_v28_: _dafny.Seq
                d_62_v28_ = _dafny.SeqWithoutIsStrInference([p1, (d_18_v0_).f14])
                d_63_v29_: D2
                d_63_v29_ = D2_DC5(p0, p0)
                d_64_v30_: _dafny.MultiSet
                d_64_v30_ = _dafny.MultiSet([d_63_v29_, d_63_v29_])
                d_65_v31_: _dafny.MultiSet
                d_65_v31_ = _dafny.MultiSet([568, (d_64_v30_).cardinality, (d_24_v6_).f14])
                d_66_v32_: _dafny.Map
                d_66_v32_ = _dafny.Map({p1: (d_62_v28_).set(default__.safeIndex((d_18_v0_).f14, len(d_62_v28_)), (d_65_v31_).cardinality)})
                d_67_v33_: _dafny.Set
                d_67_v33_ = _dafny.Set({(d_18_v0_).f14, 831})
                index13_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index13_] = (len(d_66_v32_)) in ((d_67_v33_ if p0 else d_67_v33_))
                d_68_v34_: str
                d_68_v34_ = _dafny.CodePoint('l')
                d_69_v35_: _dafny.Map
                d_69_v35_ = _dafny.Map({_dafny.Map({p0: (d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))]}): d_68_v34_})
                (globalState).f10 = (d_62_v28_)[default__.safeIndex((len(d_69_v35_)) - ((d_24_v6_).f14), len(d_62_v28_))]
                d_70_v36_: D1
                d_70_v36_ = D1_DC1(d_33_v14_)
                d_71_v37_: D1
                d_71_v37_ = D1_DC3(d_70_v36_)
                d_71_v37_ = (D1_DC3(d_70_v36_) if p0 else d_71_v37_)
                (globalState).f10 = p1
            elif True:
                (globalState).f6 = (d_18_v0_).f14
                d_72_v38_: str
                d_72_v38_ = _dafny.CodePoint('g')
                index14_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index14_] = ((0) - (len(((d_21_v3_) + (d_21_v3_)).set(default__.safeIndex((d_18_v0_).f14, len((d_21_v3_) + (d_21_v3_))), d_72_v38_)))) != ((d_18_v0_).f14)
                d_73_v39_: D2
                d_73_v39_ = D2_DC6(not((d_24_v6_).fm5(globalState)), ((d_24_v6_).f14) < (len(d_21_v3_)), default__.fm3((d_18_v0_).f14, (d_24_v6_).f14, p0, globalState))
                pat_let_tv1_ = p0
                pat_let_tv2_ = d_20_v2_
                def iife6_(_pat_let2_0):
                    def iife7_(d_74_dt__update__tmp_h1_):
                        def iife8_(_pat_let3_0):
                            def iife9_(d_75_dt__update_hcf11_h0_):
                                return D2_DC6((d_74_dt__update__tmp_h1_).cf10, d_75_dt__update_hcf11_h0_, (d_74_dt__update__tmp_h1_).cf12)
                            return iife9_(_pat_let3_0)
                        return iife8_((pat_let_tv1_) not in (pat_let_tv2_))
                    return iife7_(_pat_let2_0)
                d_73_v39_ = iife6_(d_73_v39_)
                d_76_v40_: D1
                d_76_v40_ = D1_DC2((d_18_v0_).f14, (d_18_v0_).f14, (d_24_v6_).f14, default__.fm0(globalState))
                d_77_v41_: D1
                d_77_v41_ = D1_DC3(d_76_v40_)
                d_78_v42_: _dafny.Array
                nw8_ = _dafny.Array(_dafny.Seq({}), 1)
                d_78_v42_ = nw8_
                d_79_v43_: _dafny.Seq
                d_79_v43_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({d_25_v7_: (d_24_v6_).f14})), (d_24_v6_).f14])
                index15_ = default__.safeIndex(163, (d_78_v42_).length(0))
                (d_78_v42_)[index15_] = (d_79_v43_) + (d_79_v43_)
                index16_ = default__.safeIndex(212, (d_23_v5_).length(0))
                index17_ = default__.safeIndex(163, (d_78_v42_).length(0))
                rhs10_ = d_77_v41_
                rhs11_ = (p0) == ((d_23_v5_)[default__.safeIndex(212, (d_23_v5_).length(0))])
                rhs12_ = (d_79_v43_) + (d_79_v43_)
                lhs12_ = d_23_v5_
                lhs13_ = default__.safeIndex(212, (d_23_v5_).length(0))
                lhs14_ = d_78_v42_
                lhs15_ = default__.safeIndex(163, (d_78_v42_).length(0))
                d_77_v41_ = rhs10_
                lhs12_[lhs13_] = rhs11_
                lhs14_[lhs15_] = rhs12_
                index18_ = default__.safeIndex(212, (d_23_v5_).length(0))
                (d_23_v5_)[index18_] = False
            d_80_v44_: bool
            d_80_v44_ = False
            d_80_v44_ = d_80_v44_
            d_81_v45_: D1
            d_81_v45_ = D1_DC2(-506, len(d_33_v14_), (d_24_v6_).f14, (d_24_v6_).f14)
            d_82_v46_: _dafny.Seq
            d_82_v46_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_24_v6_).f14), (d_81_v45_).cf4])
            index19_ = default__.safeIndex(212, (d_23_v5_).length(0))
            (d_23_v5_)[index19_] = (((d_24_v6_).f14) - ((0) - ((d_82_v46_)[default__.safeIndex((d_18_v0_).f14, len(d_82_v46_))]))) > (len(d_82_v46_))

    @staticmethod
    def Main(noArgsParameter__):
        d_83_v0_: int
        d_83_v0_ = 457
        d_84_v1_: bool
        d_84_v1_ = False
        d_85_v2_: _dafny.Map
        d_85_v2_ = _dafny.Map({d_83_v0_: d_84_v1_})
        d_86_v3_: _dafny.Seq
        d_86_v3_ = _dafny.SeqWithoutIsStrInference([((d_85_v2_)[len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koq")))] if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "koq")))) in (d_85_v2_) else d_84_v1_)])
        d_87_v4_: _dafny.Map
        d_87_v4_ = _dafny.Map({924: d_86_v3_})
        d_88_v5_: _dafny.Seq
        d_88_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rlnyx"))
        d_89_v6_: _dafny.Array
        nw9_ = _dafny.Array(False, 14)
        d_89_v6_ = nw9_
        d_90_globalState_: GlobalState
        nw10_ = GlobalState()
        nw10_.ctor__(True, True, True, True, (d_87_v4_) | ((d_87_v4_).set(d_83_v0_, d_86_v3_)), -8, 256, False, d_88_v5_, d_88_v5_, 133, d_89_v6_, True, 191)
        d_90_globalState_ = nw10_
        if not (((0) - ((0) - (d_83_v0_))) > (default__.fm0(d_90_globalState_))) or (d_84_v1_):
            d_91_v7_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Map({}), 21)
            d_91_v7_ = nw11_
            d_92_v8_: _dafny.Map
            d_92_v8_ = _dafny.Map({d_91_v7_: d_84_v1_})
            d_92_v8_ = (d_92_v8_).set(d_91_v7_, d_84_v1_)
            d_83_v0_ = d_83_v0_
            (d_90_globalState_).f6 = -930
            d_93_v9_: _dafny.Array
            nw12_ = _dafny.Array(int(0), 24)
            d_93_v9_ = nw12_
            d_93_v9_ = d_93_v9_
            d_94_v10_: _dafny.Seq
            d_94_v10_ = _dafny.SeqWithoutIsStrInference([d_83_v0_])
            d_95_v11_: _dafny.MultiSet
            d_95_v11_ = _dafny.MultiSet([d_83_v0_])
            d_96_v12_: _dafny.Set
            d_96_v12_ = _dafny.Set({d_95_v11_, d_95_v11_})
            d_97_v13_: _dafny.Map
            d_97_v13_ = _dafny.Map({len(d_96_v12_): d_83_v0_})
            d_98_v14_: _dafny.Set
            d_98_v14_ = _dafny.Set({320, len(d_88_v5_), len(d_94_v10_), d_83_v0_, len(d_97_v13_)})
            d_84_v1_ = ((d_98_v14_) | (_dafny.Set({d_83_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryggvk")))}))).ispropersubset(_dafny.Set({d_83_v0_, -755, d_83_v0_, d_83_v0_}))
        elif True:
            (d_90_globalState_).f10 = (0) - (d_83_v0_)
            d_99_v15_: _dafny.Map
            d_99_v15_ = _dafny.Map({d_84_v1_: d_83_v0_})
            d_100_v16_: _dafny.Seq
            d_100_v16_ = _dafny.SeqWithoutIsStrInference([len(d_85_v2_), len(d_99_v15_)])
            d_84_v1_ = default__.fm1((d_100_v16_)[default__.safeIndex(139, len(d_100_v16_))], False, d_84_v1_, (_dafny.SeqWithoutIsStrInference([d_85_v2_, default__.fm2(d_83_v0_, d_84_v1_, d_84_v1_, d_90_globalState_)])) + (_dafny.SeqWithoutIsStrInference([d_85_v2_])), d_90_globalState_)
            d_101_v17_: _dafny.Map
            d_101_v17_ = _dafny.Map({d_83_v0_: default__.fm0(d_90_globalState_)})
            d_101_v17_ = (d_101_v17_).set(len(((d_88_v5_).set(default__.safeIndex(d_83_v0_, len(d_88_v5_)), default__.fm3(d_83_v0_, d_83_v0_, d_84_v1_, d_90_globalState_))) + (d_88_v5_)), d_83_v0_)
            (d_90_globalState_).f6 = default__.safeDivisionInt(d_83_v0_, d_83_v0_)
            d_84_v1_ = d_84_v1_
        d_102_v18_: _dafny.Map
        d_102_v18_ = _dafny.Map({d_83_v0_: d_83_v0_})
        d_103_v20_: _dafny.Seq
        def iife10_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(920, -802):
                d_104_v19_: int = compr_2_
                if ((920) <= (d_104_v19_)) and ((d_104_v19_) < (-802)):
                    coll2_[(d_104_v19_) + (len(_dafny.Map({d_84_v1_: d_84_v1_})))] = d_84_v1_
            return _dafny.Map(coll2_)
        d_103_v20_ = _dafny.SeqWithoutIsStrInference([iife10_()
        ])
        d_105_v21_: _dafny.Map
        d_105_v21_ = _dafny.Map({len(d_102_v18_): d_103_v20_})
        index20_ = default__.safeIndex(506, (d_89_v6_).length(0))
        (d_89_v6_)[index20_] = default__.fm1(d_83_v0_, True, False, ((d_105_v21_)[d_83_v0_] if (d_83_v0_) in (d_105_v21_) else d_103_v20_), d_90_globalState_)
        index21_ = default__.safeIndex(506, (d_89_v6_).length(0))
        (d_89_v6_)[index21_] = d_84_v1_
        d_106_v22_: str
        d_106_v22_ = _dafny.CodePoint('i')
        d_107_v23_: _dafny.Map
        d_107_v23_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))).set(default__.safeIndex(d_83_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))), d_106_v22_): (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]})
        d_84_v1_ = not ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]) or (((d_107_v23_)[d_88_v5_] if (d_88_v5_) in (d_107_v23_) else True))
        d_108_v24_: _dafny.Array
        def lambda0_(d_109_v0_):
            def lambda1_(d_110_i0_):
                return default__.safeDivisionInt(d_110_i0_, d_109_v0_)

            return lambda1_

        init0_ = lambda0_(d_83_v0_)
        nw13_ = _dafny.Array(None, 4)
        for i0_0_ in range(nw13_.length(0)):
            nw13_[i0_0_] = init0_(i0_0_)
        d_108_v24_ = nw13_
        index22_ = default__.safeIndex(76, (d_108_v24_).length(0))
        (d_108_v24_)[index22_] = d_83_v0_
        index23_ = default__.safeIndex(597, (d_108_v24_).length(0))
        (d_108_v24_)[index23_] = (_dafny.MultiSet([(d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_84_v1_])).cardinality
        d_111_v25_: _dafny.Seq
        d_111_v25_ = _dafny.SeqWithoutIsStrInference([d_83_v0_, d_83_v0_, d_83_v0_])
        index24_ = default__.safeIndex(76, (d_108_v24_).length(0))
        index25_ = default__.safeIndex(597, (d_108_v24_).length(0))
        rhs13_ = default__.safeModuloInt(d_83_v0_, (0) - (default__.fm0(d_90_globalState_)))
        rhs14_ = ((d_83_v0_) + (len(d_102_v18_))) * (d_83_v0_)
        rhs15_ = (d_88_v5_)[default__.safeIndex(d_83_v0_, len(d_88_v5_))]
        rhs16_ = (0) - (d_83_v0_)
        rhs17_ = d_111_v25_
        lhs16_ = d_108_v24_
        lhs17_ = default__.safeIndex(76, (d_108_v24_).length(0))
        lhs18_ = d_90_globalState_
        lhs19_ = d_108_v24_
        lhs20_ = default__.safeIndex(597, (d_108_v24_).length(0))
        lhs16_[lhs17_] = rhs13_
        lhs18_.f6 = rhs14_
        d_106_v22_ = rhs15_
        lhs19_[lhs20_] = rhs16_
        d_111_v25_ = rhs17_
        hi0_ = (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]
        for d_112_i1_ in range((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], hi0_):
            d_113_v26_: _dafny.Set
            d_113_v26_ = _dafny.Set({d_84_v1_})
            d_114_v27_: _dafny.Map
            d_114_v27_ = _dafny.Map({d_84_v1_: len(d_113_v26_)})
            d_114_v27_ = (d_114_v27_).set(True, 482)
            default__.m0((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], (0) - ((0) - (len((d_111_v25_) + (d_111_v25_)))), d_90_globalState_)
            d_86_v3_ = (d_86_v3_) + (d_86_v3_)
            d_115_v28_: _dafny.Map
            d_115_v28_ = _dafny.Map({d_86_v3_: d_108_v24_})
            d_116_v29_: _dafny.Seq
            d_116_v29_ = _dafny.SeqWithoutIsStrInference([d_108_v24_, d_108_v24_, ((d_115_v28_)[_dafny.SeqWithoutIsStrInference([d_84_v1_])] if (_dafny.SeqWithoutIsStrInference([d_84_v1_])) in (d_115_v28_) else d_108_v24_)])
            d_108_v24_ = (d_116_v29_)[default__.safeIndex(d_83_v0_, len(d_116_v29_))]
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_89_v6_).length(0)):
            d_117_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_117_i2_)) and ((d_117_i2_) < ((d_89_v6_).length(0)))):
                (d_89_v6_)[(d_117_i2_)] = d_84_v1_
        d_118_v30_: _dafny.Map
        d_118_v30_ = _dafny.Map({(d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]: d_106_v22_})
        d_118_v30_ = (d_118_v30_).set(d_83_v0_, d_106_v22_)
        if d_84_v1_:
            (d_90_globalState_).f6 = ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) * ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])
            if (not(False)) or (default__.fm1(len(d_88_v5_), (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_103_v20_, d_90_globalState_)):
                d_84_v1_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
                d_89_v6_ = d_89_v6_
                d_88_v5_ = d_88_v5_
                d_119_v31_: C0
                nw14_ = C0()
                nw14_.ctor__((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])
                d_119_v31_ = nw14_
                d_120_v32_: _dafny.Map
                d_120_v32_ = _dafny.Map({(d_86_v3_) + (d_86_v3_): not ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]) or ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))])})
                d_120_v32_ = (d_120_v32_).set((d_86_v3_).set(default__.safeIndex(d_83_v0_, len(d_86_v3_)), (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]), not(((d_86_v3_)[default__.safeIndex(d_83_v0_, len(d_86_v3_))]) == ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))])))
            elif True:
                d_121_v33_: _dafny.Seq
                d_121_v33_ = _dafny.SeqWithoutIsStrInference([d_89_v6_])
                default__.m0((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], len(d_121_v33_), d_90_globalState_)
                index26_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index26_] = default__.fm0(d_90_globalState_)
                d_84_v1_ = d_84_v1_
                d_122_v34_: _dafny.Seq
                d_122_v34_ = _dafny.SeqWithoutIsStrInference([d_88_v5_, d_88_v5_])
                d_123_v35_: _dafny.Map
                d_123_v35_ = _dafny.Map({(d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]: _dafny.SeqWithoutIsStrInference([d_83_v0_, (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]])})
                index27_ = default__.safeIndex(506, (d_89_v6_).length(0))
                rhs18_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
                rhs19_ = 602
                rhs20_ = not(not((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]))
                rhs21_ = default__.safeDivisionInt(d_83_v0_, len((d_122_v34_).set(default__.safeIndex(790, len(d_122_v34_)), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")))))
                rhs22_ = (((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) * (len(d_86_v3_))) in (d_123_v35_)
                lhs21_ = d_90_globalState_
                lhs22_ = d_90_globalState_
                lhs23_ = d_89_v6_
                lhs24_ = default__.safeIndex(506, (d_89_v6_).length(0))
                d_84_v1_ = rhs18_
                lhs21_.f10 = rhs19_
                d_84_v1_ = rhs20_
                lhs22_.f6 = rhs21_
                lhs23_[lhs24_] = rhs22_
                d_124_v36_: C0
                nw15_ = C0()
                nw15_.ctor__(478)
                d_124_v36_ = nw15_
            index28_ = default__.safeIndex(506, (d_89_v6_).length(0))
            (d_89_v6_)[index28_] = d_84_v1_
            d_125_v37_: _dafny.Map
            d_125_v37_ = _dafny.Map({d_88_v5_: d_89_v6_})
            d_89_v6_ = ((d_125_v37_)[((_dafny.SeqWithoutIsStrInference([d_106_v22_ for d_126_i3_ in range(default__.abs(76))])) + (d_88_v5_)).set(default__.safeIndex(d_83_v0_, len((_dafny.SeqWithoutIsStrInference([d_106_v22_ for d_127_i3_ in range(default__.abs(76))])) + (d_88_v5_))), _dafny.CodePoint('c'))] if (((_dafny.SeqWithoutIsStrInference([d_106_v22_ for d_128_i3_ in range(default__.abs(76))])) + (d_88_v5_)).set(default__.safeIndex(d_83_v0_, len((_dafny.SeqWithoutIsStrInference([d_106_v22_ for d_129_i3_ in range(default__.abs(76))])) + (d_88_v5_))), _dafny.CodePoint('c'))) in (d_125_v37_) else d_89_v6_)
            d_130_v38_: _dafny.Set
            d_130_v38_ = _dafny.Set({d_111_v25_})
            default__.m0((d_130_v38_).ispropersubset(_dafny.Set({_dafny.SeqWithoutIsStrInference([-144])})), (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], d_90_globalState_)
        elif True:
            d_106_v22_ = _dafny.CodePoint('y')
            d_88_v5_ = d_88_v5_
            if (len(d_88_v5_)) == (d_83_v0_):
                d_131_v39_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.MultiSet({}), 25)
                d_131_v39_ = nw16_
                d_132_v40_: C0
                nw17_ = C0()
                nw17_.ctor__((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])
                d_132_v40_ = nw17_
                d_133_v41_: _dafny.Seq
                d_133_v41_ = _dafny.SeqWithoutIsStrInference([d_132_v40_])
                d_134_v42_: _dafny.MultiSet
                d_134_v42_ = _dafny.MultiSet([d_133_v41_, d_133_v41_])
                d_135_v43_: _dafny.Map
                d_135_v43_ = _dafny.Map({d_108_v24_: d_134_v42_})
                d_136_v44_: _dafny.Seq
                d_136_v44_ = _dafny.SeqWithoutIsStrInference([d_133_v41_])
                d_137_v45_: _dafny.Map
                d_137_v45_ = _dafny.Map({(d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]: d_84_v1_})
                d_138_v46_: _dafny.Map
                d_138_v46_ = _dafny.Map({(0) - (d_83_v0_): d_137_v45_})
                index29_ = default__.safeIndex(793, (d_131_v39_).length(0))
                (d_131_v39_)[index29_] = (((d_135_v43_)[d_108_v24_] if (d_108_v24_) in (d_135_v43_) else _dafny.MultiSet(d_136_v44_))).set(d_133_v41_, default__.abs(len(d_138_v46_)))
                index30_ = default__.safeIndex(793, (d_131_v39_).length(0))
                (d_131_v39_)[index30_] = d_134_v42_
                default__.m0((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], d_90_globalState_)
                d_111_v25_ = (d_111_v25_) + ((d_111_v25_) + (d_111_v25_))
                default__.m0(d_84_v1_, d_83_v0_, d_90_globalState_)
                d_102_v18_ = (D1_DC1(d_102_v18_)).cf1
            elif True:
                index31_ = default__.safeIndex(506, (d_89_v6_).length(0))
                (d_89_v6_)[index31_] = False
                d_139_v47_: _dafny.Array
                nw18_ = _dafny.Array(_dafny.Map({}), 14)
                d_139_v47_ = nw18_
                d_140_v48_: _dafny.Map
                d_140_v48_ = _dafny.Map({d_139_v47_: len(d_111_v25_)})
                d_140_v48_ = (d_140_v48_).set(d_139_v47_, (len(d_86_v3_)) - ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]))
                index32_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index32_] = 878
                d_87_v4_ = d_87_v4_
                d_141_v49_: _dafny.Set
                d_141_v49_ = _dafny.Set({(d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]})
                d_142_v50_: _dafny.Map
                d_142_v50_ = _dafny.Map({(d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]: (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]})
                d_143_v51_: _dafny.Map
                d_143_v51_ = _dafny.Map({d_141_v49_: len((d_142_v50_).set((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]))})
                d_144_v52_: C0
                nw19_ = C0()
                nw19_.ctor__((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])
                d_144_v52_ = nw19_
                d_145_v53_: _dafny.Seq
                d_145_v53_ = _dafny.SeqWithoutIsStrInference([d_144_v52_, d_144_v52_])
                d_146_v54_: _dafny.Seq
                d_146_v54_ = _dafny.SeqWithoutIsStrInference([d_143_v51_, d_143_v51_])
                d_147_v55_: _dafny.Array
                nw20_ = _dafny.Array(None, 19)
                nw20_[int(0)] = _dafny.Map({_dafny.Set({d_83_v0_}): d_83_v0_})
                nw20_[int(1)] = (d_143_v51_).set(d_141_v49_, len(d_145_v53_))
                nw20_[int(2)] = (d_143_v51_) | (d_143_v51_)
                nw20_[int(3)] = d_143_v51_
                nw20_[int(4)] = (d_146_v54_)[default__.safeIndex(d_83_v0_, len(d_146_v54_))]
                nw20_[int(5)] = (d_143_v51_) | (d_143_v51_)
                nw20_[int(6)] = d_143_v51_
                nw20_[int(7)] = d_143_v51_
                nw20_[int(8)] = default__.fm6((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_84_v1_, (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_84_v1_, d_90_globalState_)
                nw20_[int(9)] = d_143_v51_
                nw20_[int(10)] = (d_143_v51_) | (d_143_v51_)
                nw20_[int(11)] = _dafny.Map({d_141_v49_: 735})
                nw20_[int(12)] = d_143_v51_
                nw20_[int(13)] = d_143_v51_
                nw20_[int(14)] = (d_143_v51_) | (d_143_v51_)
                nw20_[int(15)] = d_143_v51_
                nw20_[int(16)] = (d_143_v51_).set(d_141_v49_, (0) - ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]))
                nw20_[int(17)] = d_143_v51_
                nw20_[int(18)] = d_143_v51_
                d_147_v55_ = nw20_
                index33_ = default__.safeIndex(728, (d_147_v55_).length(0))
                (d_147_v55_)[index33_] = d_143_v51_
                d_148_v56_: _dafny.Seq
                d_148_v56_ = _dafny.SeqWithoutIsStrInference([d_88_v5_])
                index34_ = default__.safeIndex(76, (d_108_v24_).length(0))
                index35_ = default__.safeIndex(728, (d_147_v55_).length(0))
                rhs23_ = len(d_148_v56_)
                rhs24_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
                rhs25_ = default__.fm0(d_90_globalState_)
                rhs26_ = (d_143_v51_) | (d_143_v51_)
                lhs25_ = d_108_v24_
                lhs26_ = default__.safeIndex(76, (d_108_v24_).length(0))
                lhs27_ = d_90_globalState_
                lhs28_ = d_147_v55_
                lhs29_ = default__.safeIndex(728, (d_147_v55_).length(0))
                lhs25_[lhs26_] = rhs23_
                d_84_v1_ = rhs24_
                lhs27_.f6 = rhs25_
                lhs28_[lhs29_] = rhs26_
            d_149_v57_: _dafny.Map
            d_149_v57_ = _dafny.Map({d_84_v1_: d_83_v0_})
            d_149_v57_ = d_149_v57_
            index36_ = default__.safeIndex(506, (d_89_v6_).length(0))
            (d_89_v6_)[index36_] = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
        hi1_ = 998
        for d_150_i4_ in range((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], hi1_):
            d_151_v58_: _dafny.Array
            nw21_ = _dafny.Array(None, 6)
            d_151_v58_ = nw21_
            d_152_v59_: C0
            nw22_ = C0()
            nw22_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ook"))))
            d_152_v59_ = nw22_
            index37_ = default__.safeIndex(728, (d_151_v58_).length(0))
            (d_151_v58_)[index37_] = d_152_v59_
            index38_ = default__.safeIndex(728, (d_151_v58_).length(0))
            (d_151_v58_)[index38_] = d_152_v59_
            d_153_v60_: _dafny.MultiSet
            d_153_v60_ = _dafny.MultiSet([d_84_v1_])
            d_154_v61_: D1
            d_154_v61_ = D1_DC2((d_152_v59_).f14, (d_153_v60_).cardinality, 250, d_150_i4_)
            d_155_v62_: D1
            d_155_v62_ = D1_DC3(d_154_v61_)
            d_156_v63_: D1
            d_156_v63_ = D1_DC3(d_154_v61_)
            pat_let_tv3_ = d_154_v61_
            def iife11_(_pat_let4_0):
                def iife12_(d_157_dt__update__tmp_h0_):
                    def iife13_(_pat_let5_0):
                        def iife14_(d_158_dt__update_hcf6_h0_):
                            return D1_DC3(d_158_dt__update_hcf6_h0_)
                        return iife14_(_pat_let5_0)
                    return iife13_(pat_let_tv3_)
                return iife12_(_pat_let4_0)
            d_156_v63_ = iife11_(d_156_v63_)
            d_159_v64_: _dafny.Map
            d_159_v64_ = _dafny.Map({(d_151_v58_)[default__.safeIndex(728, (d_151_v58_).length(0))]: (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]})
            if ((d_159_v64_)[(d_151_v58_)[default__.safeIndex(728, (d_151_v58_).length(0))]] if ((d_151_v58_)[default__.safeIndex(728, (d_151_v58_).length(0))]) in (d_159_v64_) else (d_84_v1_) and ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))])):
                d_160_v65_: _dafny.Set
                d_160_v65_ = _dafny.Set({(d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]})
                d_161_v66_: _dafny.MultiSet
                d_161_v66_ = _dafny.MultiSet([d_89_v6_])
                nw23_ = _dafny.Array(None, 14)
                nw23_[int(0)] = d_150_i4_
                nw23_[int(1)] = (d_152_v59_).f14
                nw23_[int(2)] = d_150_i4_
                nw23_[int(3)] = default__.safeDivisionInt((d_153_v60_).cardinality, d_150_i4_)
                nw23_[int(4)] = -490
                nw23_[int(5)] = default__.safeModuloInt((d_152_v59_).f14, 461)
                nw23_[int(6)] = d_83_v0_
                nw23_[int(7)] = (d_152_v59_).f14
                nw23_[int(8)] = (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]
                nw23_[int(9)] = len(d_160_v65_)
                nw23_[int(10)] = (d_152_v59_).f14
                nw23_[int(11)] = (d_152_v59_).f14
                nw23_[int(12)] = (0) - ((d_152_v59_).f14)
                nw23_[int(13)] = ((d_161_v66_).cardinality) * (d_83_v0_)
                d_108_v24_ = nw23_
                default__.m0(not(d_84_v1_), len(_dafny.SeqWithoutIsStrInference([d_106_v22_ for d_162_i5_ in range(default__.abs(556))])), d_90_globalState_)
                default__.m0((not(d_84_v1_)) and (d_84_v1_), (d_152_v59_).f14, d_90_globalState_)
                d_152_v59_ = d_152_v59_
                d_163_v67_: _dafny.MultiSet
                d_163_v67_ = _dafny.MultiSet([d_106_v22_, d_106_v22_])
                d_164_v68_: C0
                nw24_ = C0()
                nw24_.ctor__(((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) * (((d_163_v67_)[d_106_v22_] if (d_106_v22_) in (d_163_v67_) else d_150_i4_)))
                d_164_v68_ = nw24_
            elif True:
                d_165_v69_: D1
                d_165_v69_ = D1_DC1((d_102_v18_) | (d_102_v18_))
                d_166_v70_: _dafny.Set
                d_166_v70_ = _dafny.Set({(d_151_v58_)[default__.safeIndex(728, (d_151_v58_).length(0))]})
                d_167_v71_: D2
                d_167_v71_ = D2_DC4(d_166_v70_)
                index39_ = default__.safeIndex(506, (d_89_v6_).length(0))
                index40_ = default__.safeIndex(506, (d_89_v6_).length(0))
                index41_ = default__.safeIndex(506, (d_89_v6_).length(0))
                rhs27_ = not(default__.fm1((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], _dafny.SeqWithoutIsStrInference([d_85_v2_]), d_90_globalState_))
                rhs28_ = ((d_166_v70_ if (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))] else (d_167_v71_).cf7)).isdisjoint(d_166_v70_)
                rhs29_ = d_84_v1_
                rhs30_ = d_165_v69_
                rhs31_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
                lhs30_ = d_89_v6_
                lhs31_ = default__.safeIndex(506, (d_89_v6_).length(0))
                lhs32_ = d_89_v6_
                lhs33_ = default__.safeIndex(506, (d_89_v6_).length(0))
                lhs34_ = d_89_v6_
                lhs35_ = default__.safeIndex(506, (d_89_v6_).length(0))
                lhs30_[lhs31_] = rhs27_
                d_84_v1_ = rhs28_
                lhs32_[lhs33_] = rhs29_
                d_165_v69_ = rhs30_
                lhs34_[lhs35_] = rhs31_
                index42_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index42_] = (0) - (default__.safeModuloInt(d_83_v0_, 371))
                def lambda2_(d_168_v6_, d_169_v1_):
                    def lambda3_(d_170_i6_):
                        return default__.safeDivisionInt(d_170_i6_, (0) - (len(_dafny.SeqWithoutIsStrInference([(d_168_v6_)[default__.safeIndex(506, (d_168_v6_).length(0))], d_169_v1_, (d_168_v6_)[default__.safeIndex(506, (d_168_v6_).length(0))], True]))))

                    return lambda3_

                init1_ = lambda2_(d_89_v6_, d_84_v1_)
                nw25_ = _dafny.Array(None, 7)
                for i0_1_ in range(nw25_.length(0)):
                    nw25_[i0_1_] = init1_(i0_1_)
                d_108_v24_ = nw25_
                d_171_v72_: C0
                nw26_ = C0()
                nw26_.ctor__((0) - ((d_83_v0_) - ((d_152_v59_).f14)))
                d_171_v72_ = nw26_
                default__.m0(d_84_v1_, default__.safeDivisionInt((d_152_v59_).f14, (d_152_v59_).f14), d_90_globalState_)
            d_84_v1_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
        d_172_v73_: _dafny.Array
        nw27_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
        d_172_v73_ = nw27_
        index43_ = default__.safeIndex(506, (d_89_v6_).length(0))
        rhs32_ = d_172_v73_
        rhs33_ = d_84_v1_
        lhs36_ = d_89_v6_
        lhs37_ = default__.safeIndex(506, (d_89_v6_).length(0))
        d_172_v73_ = rhs32_
        lhs36_[lhs37_] = rhs33_
        index44_ = default__.safeIndex(76, (d_108_v24_).length(0))
        (d_108_v24_)[index44_] = (0) - ((0) - ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]))
        d_173_v74_: C0
        nw28_ = C0()
        nw28_.ctor__(d_83_v0_)
        d_173_v74_ = nw28_
        d_174_v75_: _dafny.Seq
        d_174_v75_ = _dafny.SeqWithoutIsStrInference([d_173_v74_])
        d_175_v76_: _dafny.MultiSet
        d_175_v76_ = _dafny.MultiSet([d_173_v74_])
        d_176_v77_: _dafny.Set
        d_176_v77_ = _dafny.Set({len(d_174_v75_), d_83_v0_, (0) - ((d_173_v74_).f14), len(d_86_v3_), (d_175_v76_).cardinality})
        d_84_v1_ = ((d_176_v77_) | (d_176_v77_)).isdisjoint((d_176_v77_) | (d_176_v77_))
        if d_84_v1_:
            d_177_v78_: _dafny.Map
            d_177_v78_ = _dafny.Map({d_84_v1_: (d_173_v74_).f14})
            d_178_v79_: _dafny.Seq
            d_178_v79_ = _dafny.SeqWithoutIsStrInference([d_177_v78_])
            d_179_v80_: _dafny.Array
            def lambda4_(d_180_v3_):
                def lambda5_(d_181_i7_):
                    return _dafny.SeqWithoutIsStrInference([d_180_v3_ for d_182_i8_ in range(default__.abs(700))])

                return lambda5_

            init2_ = lambda4_(d_86_v3_)
            nw29_ = _dafny.Array(None, 10)
            for i0_2_ in range(nw29_.length(0)):
                nw29_[i0_2_] = init2_(i0_2_)
            d_179_v80_ = nw29_
            d_183_v81_: _dafny.Seq
            d_183_v81_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_84_v1_, d_84_v1_])])
            index45_ = default__.safeIndex(593, (d_179_v80_).length(0))
            (d_179_v80_)[index45_] = d_183_v81_
            d_184_v82_: _dafny.Map
            d_184_v82_ = _dafny.Map({d_106_v22_: d_106_v22_})
            d_185_v83_: _dafny.Set
            d_185_v83_ = _dafny.Set({d_184_v82_})
            d_186_v84_: D2
            d_186_v84_ = D2_DC6(not((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]), d_84_v1_, d_106_v22_)
            d_187_v86_: _dafny.Set
            d_187_v86_ = _dafny.Set({d_106_v22_})
            d_188_v87_: D3
            d_188_v87_ = D3_DC8(d_187_v86_)
            index46_ = default__.safeIndex(593, (d_179_v80_).length(0))
            def iife15_():
                def iife17_():
                    coll5_ = _dafny.Map()
                    compr_5_: str
                    for compr_5_ in ((d_188_v87_).cf14).Elements:
                        d_189_v85_: str = compr_5_
                        if (d_189_v85_) in ((d_188_v87_).cf14):
                            coll5_[d_189_v85_] = d_106_v22_
                    return _dafny.Map(coll5_)
                coll3_ = _dafny.Set()
                def iife16_():
                    coll4_ = _dafny.Map()
                    compr_4_: str
                    for compr_4_ in ((d_188_v87_).cf14).Elements:
                        d_189_v85_: str = compr_4_
                        if (d_189_v85_) in ((d_188_v87_).cf14):
                            coll4_[d_189_v85_] = d_106_v22_
                    return _dafny.Map(coll4_)
                compr_3_: _dafny.Map
                for compr_3_ in (_dafny.MultiSet([d_184_v82_, iife16_()
                ])).Elements:
                    d_190_v88_: _dafny.Map = compr_3_
                    if (d_190_v88_) in (_dafny.MultiSet([d_184_v82_, iife17_()
                    ])):
                        coll3_ = coll3_.union(_dafny.Set([d_190_v88_]))
                return _dafny.Set(coll3_)
            rhs34_ = (d_178_v79_).set(default__.safeIndex((d_173_v74_).f14, len(d_178_v79_)), d_177_v78_)
            rhs35_ = d_183_v81_
            rhs36_ = (d_185_v83_).intersection(iife15_()
            )
            rhs37_ = d_186_v84_
            lhs38_ = d_179_v80_
            lhs39_ = default__.safeIndex(593, (d_179_v80_).length(0))
            d_178_v79_ = rhs34_
            lhs38_[lhs39_] = rhs35_
            d_185_v83_ = rhs36_
            d_186_v84_ = rhs37_
            default__.m0((D2_DC5(d_84_v1_, (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))])).cf9, default__.safeDivisionInt((d_173_v74_).f14, (d_173_v74_).f14), d_90_globalState_)
            index47_ = default__.safeIndex(76, (d_108_v24_).length(0))
            (d_108_v24_)[index47_] = (d_173_v74_).f14
            index48_ = default__.safeIndex(76, (d_108_v24_).length(0))
            def iife18_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(-729, 42):
                    d_191_v89_: int = compr_6_
                    if ((-729) <= (d_191_v89_)) and ((d_191_v89_) < (42)):
                        coll6_[default__.safeModuloInt(d_191_v89_, (d_173_v74_).f14)] = d_84_v1_
                return _dafny.Map(coll6_)
            rhs38_ = (len(d_88_v5_)) != (len(d_102_v18_))
            rhs39_ = iife18_()

            rhs40_ = 516
            rhs41_ = (len(d_88_v5_)) * ((d_173_v74_).f14)
            rhs42_ = (0) - (d_83_v0_)
            lhs40_ = d_108_v24_
            lhs41_ = default__.safeIndex(76, (d_108_v24_).length(0))
            lhs42_ = d_90_globalState_
            lhs43_ = d_90_globalState_
            d_84_v1_ = rhs38_
            d_85_v2_ = rhs39_
            lhs40_[lhs41_] = rhs40_
            lhs42_.f10 = rhs41_
            lhs43_.f6 = rhs42_
            d_85_v2_ = (d_85_v2_).set((d_173_v74_).f14, d_84_v1_)
        elif True:
            if d_84_v1_:
                d_192_v90_: _dafny.Seq
                d_192_v90_ = _dafny.SeqWithoutIsStrInference([d_176_v77_])
                index49_ = default__.safeIndex(506, (d_89_v6_).length(0))
                rhs43_ = (d_173_v74_).fm4(d_84_v1_, d_90_globalState_)
                rhs44_ = d_84_v1_
                rhs45_ = (d_192_v90_)[default__.safeIndex(((d_173_v74_).f14) + ((d_173_v74_).f14), len(d_192_v90_))]
                rhs46_ = d_106_v22_
                lhs44_ = d_89_v6_
                lhs45_ = default__.safeIndex(506, (d_89_v6_).length(0))
                lhs44_[lhs45_] = rhs43_
                d_84_v1_ = rhs44_
                d_176_v77_ = rhs45_
                d_106_v22_ = rhs46_
                d_193_v91_: _dafny.Seq
                d_193_v91_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_106_v22_})])
                d_194_v92_: _dafny.Map
                d_194_v92_ = _dafny.Map({d_173_v74_: (d_193_v91_)[default__.safeIndex((d_173_v74_).f14, len(d_193_v91_))]})
                d_195_v93_: D3
                d_195_v93_ = D3_DC8(((d_194_v92_)[d_173_v74_] if (d_173_v74_) in (d_194_v92_) else _dafny.Set({d_106_v22_})))
                d_196_v94_: _dafny.MultiSet
                d_196_v94_ = _dafny.MultiSet([d_195_v93_])
                d_197_v95_: _dafny.MultiSet
                d_197_v95_ = _dafny.MultiSet([d_84_v1_, (d_84_v1_) or (d_84_v1_), (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_84_v1_, (d_196_v94_) == (d_196_v94_)])
                d_197_v95_ = ((d_197_v95_).intersection(d_197_v95_)) - (d_197_v95_)
                default__.m0(d_84_v1_, (d_173_v74_).f14, d_90_globalState_)
                d_198_v96_: _dafny.Array
                nw30_ = _dafny.Array(None, 28)
                d_198_v96_ = nw30_
                d_198_v96_ = d_198_v96_
                d_199_v97_: _dafny.MultiSet
                d_199_v97_ = _dafny.MultiSet([235])
                (d_90_globalState_).f10 = (d_111_v25_)[default__.safeIndex(((d_199_v97_)[(d_173_v74_).f14] if ((d_173_v74_).f14) in (d_199_v97_) else (d_173_v74_).f14), len(d_111_v25_))]
            elif True:
                nw31_ = _dafny.Array(int(0), 18)
                d_108_v24_ = nw31_
                d_200_v98_: _dafny.MultiSet
                d_200_v98_ = _dafny.MultiSet([d_106_v22_])
                index50_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index50_] = (((d_200_v98_) | (_dafny.MultiSet(d_88_v5_))).cardinality) - ((d_173_v74_).f14)
                d_201_v99_: _dafny.Set
                d_201_v99_ = _dafny.Set({(d_174_v75_)[default__.safeIndex(d_83_v0_, len(d_174_v75_))]})
                d_202_v100_: D2
                d_202_v100_ = D2_DC4((d_201_v99_).intersection(d_201_v99_))
                d_202_v100_ = d_202_v100_
                d_203_v101_: D2
                d_203_v101_ = D2_DC5(not((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]), d_84_v1_)
                d_204_v102_: _dafny.Map
                d_204_v102_ = _dafny.Map({_dafny.Set({d_108_v24_, d_108_v24_, d_108_v24_}): (d_203_v101_ if (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))] else d_203_v101_)})
                d_204_v102_ = (d_204_v102_).set(_dafny.Set({d_108_v24_, d_108_v24_}), d_203_v101_)
                default__.m0((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], (d_173_v74_).f14, d_90_globalState_)
            d_108_v24_ = d_108_v24_
            index51_ = default__.safeIndex(76, (d_108_v24_).length(0))
            (d_108_v24_)[index51_] = (0) - ((d_173_v74_).f14)
            d_205_v103_: _dafny.MultiSet
            d_205_v103_ = _dafny.MultiSet([d_84_v1_])
            d_206_v104_: _dafny.Map
            d_206_v104_ = _dafny.Map({d_85_v2_: not((d_86_v3_)[default__.safeIndex(((d_205_v103_)[(d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]] if ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]) in (d_205_v103_) else (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]), len(d_86_v3_))])})
            index52_ = default__.safeIndex(506, (d_89_v6_).length(0))
            rhs47_ = (d_205_v103_).ispropersubset(d_205_v103_)
            rhs48_ = default__.safeModuloInt(d_83_v0_, (d_173_v74_).f14)
            rhs49_ = (d_206_v104_) | (d_206_v104_)
            lhs46_ = d_89_v6_
            lhs47_ = default__.safeIndex(506, (d_89_v6_).length(0))
            lhs48_ = d_90_globalState_
            lhs46_[lhs47_] = rhs47_
            lhs48_.f6 = rhs48_
            d_206_v104_ = rhs49_
            if d_84_v1_:
                index53_ = default__.safeIndex(506, (d_89_v6_).length(0))
                (d_89_v6_)[index53_] = (d_83_v0_) != ((d_173_v74_).f14)
                index54_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index54_] = (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]
                d_84_v1_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
                d_207_v105_: _dafny.Seq
                d_207_v105_ = _dafny.SeqWithoutIsStrInference([d_111_v25_])
                d_207_v105_ = d_207_v105_
                d_84_v1_ = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
            elif True:
                d_84_v1_ = (_dafny.SeqWithoutIsStrInference([611, d_83_v0_])) < (d_111_v25_)
                d_84_v1_ = not ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]) or (d_84_v1_)
                d_208_v106_: _dafny.Array
                nw32_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_208_v106_ = nw32_
                index55_ = default__.safeIndex(243, (d_208_v106_).length(0))
                (d_208_v106_)[index55_] = d_175_v76_
                index56_ = default__.safeIndex(243, (d_208_v106_).length(0))
                (d_208_v106_)[index56_] = ((_dafny.MultiSet([d_173_v74_])) | (_dafny.MultiSet([d_173_v74_]))).intersection((d_175_v76_).set(d_173_v74_, default__.abs((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])))
                (d_90_globalState_).f10 = ((d_173_v74_).f14) - ((d_173_v74_).f14)
                index57_ = default__.safeIndex(506, (d_89_v6_).length(0))
                (d_89_v6_)[index57_] = False
        index58_ = default__.safeIndex(506, (d_89_v6_).length(0))
        (d_89_v6_)[index58_] = (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]
        d_209_v107_: D2
        d_209_v107_ = D2_DC5(d_84_v1_, (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))])
        d_210_v108_: D2
        d_210_v108_ = D2_DC7(d_209_v107_)
        source2_ = (d_210_v108_ if d_84_v1_ else d_210_v108_)
        if source2_.is_DC5:
            d_211___mcc_h0_ = source2_.cf8
            d_212___mcc_h1_ = source2_.cf9
            d_213_cf9_ = d_212___mcc_h1_
            d_214_cf8_ = d_211___mcc_h0_
            d_215_v109_: _dafny.Map
            d_215_v109_ = _dafny.Map({d_106_v22_: d_173_v74_})
            d_215_v109_ = (_dafny.Map({_dafny.CodePoint('l'): d_173_v74_})).set(d_106_v22_, d_173_v74_)
            d_216_v110_: _dafny.Map
            d_216_v110_ = _dafny.Map({d_83_v0_: D3_DC10(D2_DC5(d_213_cf9_, d_214_cf8_), d_106_v22_, len(d_88_v5_))})
            d_217_v111_: D2
            d_217_v111_ = D2_DC5(d_213_cf9_, d_214_cf8_)
            d_218_v112_: D3
            d_218_v112_ = D3_DC10(d_217_v111_, d_106_v22_, (d_173_v74_).f14)
            d_216_v110_ = (d_216_v110_).set((d_173_v74_).f14, d_218_v112_)
            d_213_cf9_ = not((d_88_v5_) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fglj"))))
            d_84_v1_ = (d_173_v74_).fm5(d_90_globalState_)
        elif source2_.is_DC6:
            d_219___mcc_h2_ = source2_.cf10
            d_220___mcc_h3_ = source2_.cf11
            d_221___mcc_h4_ = source2_.cf12
            d_222_cf12_ = d_221___mcc_h4_
            d_223_cf11_ = d_220___mcc_h3_
            d_224_cf10_ = d_219___mcc_h2_
            def lambda6_(d_225_v25_):
                def lambda7_(d_226_i9_):
                    return default__.safeDivisionInt(d_226_i9_, len(d_225_v25_))

                return lambda7_

            init3_ = lambda6_(d_111_v25_)
            nw33_ = _dafny.Array(None, 29)
            for i0_3_ in range(nw33_.length(0)):
                nw33_[i0_3_] = init3_(i0_3_)
            d_108_v24_ = nw33_
            d_227_v113_: C0
            nw34_ = C0()
            nw34_.ctor__((d_173_v74_).f14)
            d_227_v113_ = nw34_
            d_228_v114_: _dafny.Map
            d_228_v114_ = _dafny.Map({d_84_v1_: d_223_cf11_})
            index59_ = default__.safeIndex(506, (d_89_v6_).length(0))
            (d_89_v6_)[index59_] = ((d_228_v114_)[(d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]] if ((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]) in (d_228_v114_) else False)
            default__.m0(d_84_v1_, 644, d_90_globalState_)
        elif source2_.is_DC4:
            d_229___mcc_h5_ = source2_.cf7
            d_230_cf7_ = d_229___mcc_h5_
            if d_84_v1_:
                d_231_v115_: D3
                d_231_v115_ = D3_DC9((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))])
                d_232_v116_: _dafny.Map
                d_232_v116_ = _dafny.Map({(0) - (((d_173_v74_).f14) + (d_83_v0_)): d_231_v115_})
                pat_let_tv4_ = d_84_v1_
                def iife19_(_pat_let6_0):
                    def iife20_(d_233_dt__update__tmp_h1_):
                        def iife21_(_pat_let7_0):
                            def iife22_(d_234_dt__update_hcf15_h0_):
                                return D3_DC9(d_234_dt__update_hcf15_h0_)
                            return iife22_(_pat_let7_0)
                        return iife21_(not(pat_let_tv4_))
                    return iife20_(_pat_let6_0)
                d_232_v116_ = (d_232_v116_).set((d_173_v74_).f14, iife19_(d_231_v115_))
                d_235_v117_: C0
                nw35_ = C0()
                nw35_.ctor__((d_173_v74_).f14)
                d_235_v117_ = nw35_
                d_84_v1_ = d_84_v1_
                d_236_v118_: _dafny.Map
                d_236_v118_ = _dafny.Map({d_83_v0_: d_108_v24_})
                d_236_v118_ = ((d_236_v118_).set((d_235_v117_).f14, d_108_v24_)) | ((d_236_v118_) | (d_236_v118_))
                d_173_v74_ = d_173_v74_
            elif True:
                d_102_v18_ = (d_102_v18_).set(len(_dafny.SeqWithoutIsStrInference([d_84_v1_, (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))]])), (d_173_v74_).f14)
                d_237_v119_: _dafny.Array
                nw36_ = _dafny.Array(None, 2)
                nw36_[int(0)] = d_172_v73_
                nw36_[int(1)] = d_172_v73_
                d_237_v119_ = nw36_
                index60_ = default__.safeIndex(43, (d_237_v119_).length(0))
                (d_237_v119_)[index60_] = (d_172_v73_ if (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))] else d_172_v73_)
                d_238_v120_: D2
                d_238_v120_ = D2_DC4(d_230_cf7_)
                d_239_v121_: _dafny.Set
                d_239_v121_ = _dafny.Set({d_238_v120_, d_238_v120_, d_238_v120_, d_238_v120_, D2_DC4(d_230_cf7_)})
                index61_ = default__.safeIndex(43, (d_237_v119_).length(0))
                index62_ = default__.safeIndex(76, (d_108_v24_).length(0))
                rhs50_ = d_172_v73_
                rhs51_ = (d_173_v74_).f14
                rhs52_ = ((d_173_v74_).f14) == (d_83_v0_)
                rhs53_ = len((d_239_v121_) | (d_239_v121_))
                rhs54_ = default__.safeDivisionInt((d_173_v74_).f14, len(d_88_v5_))
                lhs49_ = d_237_v119_
                lhs50_ = default__.safeIndex(43, (d_237_v119_).length(0))
                lhs51_ = d_108_v24_
                lhs52_ = default__.safeIndex(76, (d_108_v24_).length(0))
                lhs53_ = d_90_globalState_
                lhs54_ = d_90_globalState_
                lhs49_[lhs50_] = rhs50_
                lhs51_[lhs52_] = rhs51_
                d_84_v1_ = rhs52_
                lhs53_.f6 = rhs53_
                lhs54_.f6 = rhs54_
                d_83_v0_ = default__.fm0(d_90_globalState_)
                d_106_v22_ = d_106_v22_
                d_240_v122_: C0
                nw37_ = C0()
                nw37_.ctor__((d_173_v74_).f14)
                d_240_v122_ = nw37_
            d_241_v123_: C0
            nw38_ = C0()
            nw38_.ctor__((0) - (default__.safeModuloInt((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], 806)))
            d_241_v123_ = nw38_
            if ((_dafny.MultiSet([d_173_v74_])).intersection(d_175_v76_)).ispropersubset(d_175_v76_):
                index63_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index63_] = (len(d_88_v5_)) - ((d_173_v74_).f14)
                (d_90_globalState_).f6 = (d_173_v74_).f14
                (d_90_globalState_).f6 = (d_173_v74_).f14
                d_242_v124_: _dafny.Array
                nw39_ = _dafny.Array(D3.default()(), 28)
                d_242_v124_ = nw39_
                d_243_v125_: _dafny.MultiSet
                d_243_v125_ = _dafny.MultiSet([d_242_v124_, d_242_v124_, d_242_v124_, d_242_v124_, d_242_v124_])
                rhs55_ = d_89_v6_
                rhs56_ = ((d_243_v125_)[(D4_DC11(d_242_v124_)).cf19] if ((D4_DC11(d_242_v124_)).cf19) in (d_243_v125_) else ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) + ((d_173_v74_).f14))
                rhs57_ = (d_241_v123_).f14
                rhs58_ = (d_173_v74_).f14
                lhs55_ = d_90_globalState_
                lhs56_ = d_90_globalState_
                lhs57_ = d_90_globalState_
                d_89_v6_ = rhs55_
                lhs55_.f6 = rhs56_
                lhs56_.f6 = rhs57_
                lhs57_.f6 = rhs58_
                d_244_v126_: _dafny.MultiSet
                d_244_v126_ = _dafny.MultiSet([d_89_v6_])
                (d_90_globalState_).f6 = default__.safeDivisionInt(((d_173_v74_).f14) + ((d_244_v126_).cardinality), (0) - ((0) - (((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) * (d_83_v0_))))
            elif True:
                default__.m0(d_84_v1_, (0) - (((d_241_v123_).f14) + (254)), d_90_globalState_)
                d_88_v5_ = d_88_v5_
                d_84_v1_ = (len(d_88_v5_)) <= ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])
                d_245_v127_: _dafny.MultiSet
                d_245_v127_ = _dafny.MultiSet([(d_173_v74_).f14])
                rhs59_ = (((d_245_v127_)[(d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]] if ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) in (d_245_v127_) else (d_241_v123_).f14)) + (((d_173_v74_).f14) * (len(default__.fm7((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_106_v22_, d_90_globalState_))))
                rhs60_ = (0) - ((d_241_v123_).f14)
                lhs58_ = d_90_globalState_
                lhs59_ = d_90_globalState_
                lhs58_.f6 = rhs59_
                lhs59_.f6 = rhs60_
                nw40_ = _dafny.Array(int(0), 23)
                d_108_v24_ = nw40_
            d_246_v128_: _dafny.Map
            d_246_v128_ = _dafny.Map({d_173_v74_: d_111_v25_})
            d_247_v129_: _dafny.Seq
            d_247_v129_ = _dafny.SeqWithoutIsStrInference([(d_246_v128_) | (d_246_v128_)])
            d_246_v128_ = (d_247_v129_)[default__.safeIndex((d_173_v74_).f14, len(d_247_v129_))]
        elif True:
            d_248___mcc_h6_ = source2_.cf13
            d_249_cf13_ = d_248___mcc_h6_
            nw41_ = C0()
            nw41_.ctor__(default__.safeDivisionInt((d_173_v74_).f14, d_83_v0_))
            d_173_v74_ = nw41_
            index64_ = default__.safeIndex(506, (d_89_v6_).length(0))
            rhs61_ = ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]) >= ((d_173_v74_).f14)
            rhs62_ = (d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))]
            lhs60_ = d_89_v6_
            lhs61_ = default__.safeIndex(506, (d_89_v6_).length(0))
            lhs62_ = d_90_globalState_
            lhs60_[lhs61_] = rhs61_
            lhs62_.f6 = rhs62_
            d_84_v1_ = (_dafny.SeqWithoutIsStrInference([d_84_v1_])) == (default__.fm8((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_90_globalState_))
            if (d_86_v3_)[default__.safeIndex((d_173_v74_).f14, len(d_86_v3_))]:
                d_250_v130_: D3
                d_250_v130_ = D3_DC8(_dafny.Set({d_106_v22_, d_106_v22_, d_106_v22_, d_106_v22_, _dafny.CodePoint('v')}))
                d_251_v131_: _dafny.Set
                d_251_v131_ = _dafny.Set({d_106_v22_, d_106_v22_})
                pat_let_tv5_ = d_251_v131_
                def iife23_(_pat_let8_0):
                    def iife24_(d_252_dt__update__tmp_h2_):
                        def iife25_(_pat_let9_0):
                            def iife26_(d_253_dt__update_hcf14_h0_):
                                return D3_DC8(d_253_dt__update_hcf14_h0_)
                            return iife26_(_pat_let9_0)
                        return iife25_(pat_let_tv5_)
                    return iife24_(_pat_let8_0)
                d_250_v130_ = iife23_(d_250_v130_)
                d_250_v130_ = (d_250_v130_ if (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))] else d_250_v130_)
                index65_ = default__.safeIndex(437, (d_172_v73_).length(0))
                (d_172_v73_)[index65_] = d_88_v5_
                index66_ = default__.safeIndex(437, (d_172_v73_).length(0))
                (d_172_v73_)[index66_] = (d_88_v5_) + (d_88_v5_)
                d_254_v132_: _dafny.Array
                nw42_ = _dafny.Array(None, 15)
                d_254_v132_ = nw42_
                d_254_v132_ = d_254_v132_
                (d_90_globalState_).f6 = (d_173_v74_).f14
            elif True:
                (d_90_globalState_).f6 = (0) - ((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))])
                d_255_v133_: _dafny.Seq
                d_255_v133_ = _dafny.SeqWithoutIsStrInference([d_108_v24_, d_108_v24_, d_108_v24_, d_108_v24_])
                d_256_v134_: _dafny.MultiSet
                d_256_v134_ = _dafny.MultiSet([(d_255_v133_)[default__.safeIndex((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], len(d_255_v133_))]])
                index67_ = default__.safeIndex(506, (d_89_v6_).length(0))
                (d_89_v6_)[index67_] = (_dafny.MultiSet([d_108_v24_])).ispropersubset(d_256_v134_)
                index68_ = default__.safeIndex(76, (d_108_v24_).length(0))
                (d_108_v24_)[index68_] = len((d_88_v5_) + (default__.fm7((d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], default__.fm3((d_108_v24_)[default__.safeIndex(76, (d_108_v24_).length(0))], -30, (d_89_v6_)[default__.safeIndex(506, (d_89_v6_).length(0))], d_90_globalState_), d_90_globalState_)))
                (d_90_globalState_).f10 = (0) - (d_83_v0_)
                d_257_v135_: _dafny.Map
                d_257_v135_ = _dafny.Map({d_176_v77_: default__.fm0(d_90_globalState_)})
                d_257_v135_ = (d_257_v135_).set(d_176_v77_, (d_173_v74_).f14)
        default__.m0(False, (d_173_v74_).f14, d_90_globalState_)
        _dafny.print(_dafny.string_of(d_83_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_84_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_85_v2_) == (_dafny.Map({457: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_86_v3_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v4_) == (_dafny.Map({924: _dafny.SeqWithoutIsStrInference([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_88_v5_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_89_v6_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f4) == (_dafny.Map({924: _dafny.SeqWithoutIsStrInference([False]), 457: _dafny.SeqWithoutIsStrInference([False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_90_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_90_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_90_globalState_).f9).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_90_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_90_globalState_).f11)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v18_) == (_dafny.Map({457: 457}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v20_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_v21_) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference([_dafny.Map({})])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_106_v22_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_107_v23_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_108_v24_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_v25_) == (_dafny.SeqWithoutIsStrInference([457, 457, 457]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_118_v30_) == (_dafny.Map({1: _dafny.CodePoint('n'), 457: _dafny.CodePoint('n')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v74_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_174_v75_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v76_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v77_) == (_dafny.Set({1, 457, -457}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v107_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_209_v107_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_210_v108_).cf13).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_210_v108_).cf13).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: False
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(int(0), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC5(D2, NamedTuple('DC5', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12(False, False, False, _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f6: int = int(0)
        self.f10: int = int(0)
        self._f0: bool = False
        self._f1: bool = False
        self._f2: bool = False
        self._f3: bool = False
        self._f4: _dafny.Map = _dafny.Map({})
        self._f5: int = int(0)
        self._f7: bool = False
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self._f12: bool = False
        self._f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13

class C0:
    def  __init__(self):
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f14):
        (self)._f14 = f14

    def fm4(self, p0, globalState):
        return (True)

    def fm5(self, globalState):
        return (_dafny.MultiSet([not(True), False, not(False), True])).isdisjoint((_dafny.MultiSet([True, False])) | (_dafny.MultiSet([True])))

    @property
    def f14(self):
        return self._f14
