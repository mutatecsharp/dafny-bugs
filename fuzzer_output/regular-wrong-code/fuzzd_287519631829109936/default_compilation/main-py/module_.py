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
    def fm0(p0, p1, globalState):
        return ((448) * (len(_dafny.SeqWithoutIsStrInference([False, True])))) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_0_i0_ in range(default__.abs(393))])), len(_dafny.Map({False: len(_dafny.Map({False: not(False)}))}))))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return True

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1_i0_ in range(default__.abs(-484))])

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False, not(True)]))) + ((_dafny.SeqWithoutIsStrInference([(D7_DC22(True, 629, 716)).cf35])) + (_dafny.SeqWithoutIsStrInference([False, True])))

    @staticmethod
    def fm4(p0, globalState):
        source0_ = D7_DC23(True, len(_dafny.Map({(0) - (-437): D7_DC21(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([D2_DC7(_dafny.CodePoint('i')), D2_DC7(_dafny.CodePoint('f')), D2_DC7(_dafny.CodePoint('o')), D2_DC7(_dafny.CodePoint('r')), D2_DC7(_dafny.CodePoint('m'))])))})))
        if source0_.is_DC22:
            d_2___mcc_h0_ = source0_.cf35
            d_3___mcc_h1_ = source0_.cf36
            d_4___mcc_h2_ = source0_.cf37
            d_5_cf37_ = d_4___mcc_h2_
            d_6_cf36_ = d_3___mcc_h1_
            d_7_cf35_ = d_2___mcc_h0_
            return D0_DC1(D0_DC0(d_6_cf36_, _dafny.SeqWithoutIsStrInference([True])))
        elif source0_.is_DC23:
            d_8___mcc_h3_ = source0_.cf38
            d_9___mcc_h4_ = source0_.cf39
            d_10_cf39_ = d_9___mcc_h4_
            d_11_cf38_ = d_8___mcc_h3_
            return D0_DC1((D0_DC1(D0_DC0(d_10_cf39_, _dafny.SeqWithoutIsStrInference([d_11_cf38_])))).cf2)
        elif True:
            d_12___mcc_h5_ = source0_.cf34
            d_13_cf34_ = d_12___mcc_h5_
            return D0_DC1(D0_DC0(len(_dafny.Map({_dafny.Map({224: True}): len(_dafny.SeqWithoutIsStrInference([186]))})), _dafny.SeqWithoutIsStrInference([True, False, not(not(False))])))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return D1_DC2(_dafny.CodePoint('e'))

    @staticmethod
    def fm9(p0, p1, p2, p3, globalState):
        return _dafny.Set({(True if True else True), True})

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        if True:
            def iife0_():
                coll0_ = _dafny.Set()
                compr_0_: D0
                for compr_0_ in (_dafny.MultiSet([D0_DC0(len(_dafny.Map({_dafny.CodePoint('b'): True})), _dafny.SeqWithoutIsStrInference([True, False]))])).Elements:
                    d_14_v0_: D0 = compr_0_
                    if (d_14_v0_) in (_dafny.MultiSet([D0_DC0(len(_dafny.Map({_dafny.CodePoint('b'): True})), _dafny.SeqWithoutIsStrInference([True, False]))])):
                        coll0_ = coll0_.union(_dafny.Set([d_14_v0_]))
                return _dafny.Set(coll0_)
            return iife0_()
            
        elif True:
            return _dafny.Set({D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngymbxll"))), _dafny.SeqWithoutIsStrInference([False])), D0_DC0(531, _dafny.SeqWithoutIsStrInference([not(not(True)), not(False), False]))})

    @staticmethod
    def fm11(globalState):
        source1_ = D0_DC0(len(_dafny.Set({934, 600})), _dafny.SeqWithoutIsStrInference([True, not(False)]))
        if source1_.is_DC0:
            d_15___mcc_h0_ = source1_.cf0
            d_16___mcc_h1_ = source1_.cf1
            d_17_cf1_ = d_16___mcc_h1_
            d_18_cf0_ = d_15___mcc_h0_
            return (D1_DC2(_dafny.CodePoint('y'))).cf3
        elif True:
            d_19___mcc_h2_ = source1_.cf2
            d_20_cf2_ = d_19___mcc_h2_
            return _dafny.CodePoint('y')

    @staticmethod
    def fm12(p0, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.Map({len(_dafny.Set({False, False})): True})).keys.Elements:
                d_21_v0_: int = compr_1_
                if (d_21_v0_) in (_dafny.Map({len(_dafny.Set({False, False})): True})):
                    coll1_[(d_21_v0_) + (-936)] = False
            return _dafny.Map(coll1_)
        return ((_dafny.MultiSet([490])) - (_dafny.MultiSet([969]))) - ((D8_DC24(_dafny.MultiSet([246, len(iife1_()
)]))).cf40)

    @staticmethod
    def fm13(p0, globalState):
        return _dafny.Map({default__.safeModuloInt(451, 803): False})

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        r0: bool = False
        d_22_v0_: _dafny.Seq
        d_22_v0_ = _dafny.SeqWithoutIsStrInference([p1, p0, p1, p2, p3])
        d_23_v1_: int
        d_23_v1_ = 761
        d_24_v2_: _dafny.Seq
        d_24_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeoyscw"))
        d_25_v3_: _dafny.Array
        nw0_ = _dafny.Array(None, 21)
        nw0_[int(0)] = d_22_v0_
        nw0_[int(1)] = d_22_v0_
        nw0_[int(2)] = default__.fm3(d_23_v1_, d_23_v1_, True, globalState)
        nw0_[int(3)] = d_22_v0_
        nw0_[int(4)] = d_22_v0_
        nw0_[int(5)] = d_22_v0_
        nw0_[int(6)] = (d_22_v0_ if p0 else d_22_v0_)
        nw0_[int(7)] = _dafny.SeqWithoutIsStrInference([True])
        nw0_[int(8)] = d_22_v0_
        nw0_[int(9)] = (_dafny.SeqWithoutIsStrInference([not(p0)]) if False else d_22_v0_)
        nw0_[int(10)] = _dafny.SeqWithoutIsStrInference([p3])
        nw0_[int(11)] = (d_22_v0_) + (default__.fm3(669, d_23_v1_, p0, globalState))
        nw0_[int(12)] = d_22_v0_
        nw0_[int(13)] = d_22_v0_
        nw0_[int(14)] = (d_22_v0_).set(default__.safeIndex(default__.fm0(not(p2), d_24_v2_, globalState), len(d_22_v0_)), default__.fm1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpcnkcvm")), p3, d_23_v1_, globalState))
        nw0_[int(15)] = d_22_v0_
        nw0_[int(16)] = d_22_v0_
        nw0_[int(17)] = (_dafny.SeqWithoutIsStrInference([p3, p1])) + (d_22_v0_)
        nw0_[int(18)] = d_22_v0_
        nw0_[int(19)] = (d_22_v0_) + (d_22_v0_)
        nw0_[int(20)] = d_22_v0_
        d_25_v3_ = nw0_
        index0_ = default__.safeIndex(91, (d_25_v3_).length(0))
        (d_25_v3_)[index0_] = _dafny.SeqWithoutIsStrInference([p2])
        index1_ = default__.safeIndex(91, (d_25_v3_).length(0))
        (d_25_v3_)[index1_] = d_22_v0_
        d_26_v4_: str
        d_26_v4_ = _dafny.CodePoint('j')
        d_27_v5_: _dafny.Array
        nw1_ = _dafny.Array(int(0), 9)
        d_27_v5_ = nw1_
        d_28_v6_: _dafny.Map
        d_28_v6_ = _dafny.Map({d_26_v4_: d_27_v5_})
        d_29_v7_: _dafny.Map
        d_29_v7_ = _dafny.Map({p2: p1})
        d_30_v8_: _dafny.Seq
        d_30_v8_ = _dafny.SeqWithoutIsStrInference([d_29_v7_])
        d_31_v9_: _dafny.Map
        d_31_v9_ = _dafny.Map({d_30_v8_: d_27_v5_})
        d_28_v6_ = (d_28_v6_).set(d_26_v4_, ((d_31_v9_)[d_30_v8_] if (d_30_v8_) in (d_31_v9_) else d_27_v5_))
        d_32_v10_: _dafny.MultiSet
        d_32_v10_ = _dafny.MultiSet([p0, not(not(p0))])
        rhs0_ = ((d_24_v2_ if p2 else (d_24_v2_).set(default__.safeIndex(497, len(d_24_v2_)), (d_24_v2_)[default__.safeIndex(d_23_v1_, len(d_24_v2_))]))) <= ((d_24_v2_) + (d_24_v2_))
        rhs1_ = p2
        rhs2_ = d_32_v10_
        lhs0_ = globalState
        lhs1_ = globalState
        lhs0_.f7 = rhs0_
        lhs1_.f6 = rhs1_
        d_32_v10_ = rhs2_
        (globalState).f7 = (d_23_v1_) != ((d_23_v1_) - (d_23_v1_))
        (globalState).f7 = ((p2) not in (default__.fm3((0) - (d_23_v1_), d_23_v1_, p0, globalState)) if p0 else p2)
        (globalState).f7 = (d_23_v1_) != (default__.safeDivisionInt(len(default__.fm13(p0, globalState)), d_23_v1_))
        r0 = not (p1) or (p0)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_33_v0_: str
        d_33_v0_ = _dafny.CodePoint('n')
        d_34_v1_: _dafny.MultiSet
        d_34_v1_ = _dafny.MultiSet([d_33_v0_, d_33_v0_])
        d_35_v2_: int
        d_35_v2_ = 863
        d_36_v3_: _dafny.Seq
        d_36_v3_ = _dafny.SeqWithoutIsStrInference([((d_34_v1_)[d_33_v0_] if (d_33_v0_) in (d_34_v1_) else (0) - (d_35_v2_))])
        d_37_v4_: _dafny.Seq
        d_37_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yw"))
        d_38_v8_: bool
        d_38_v8_ = False
        d_39_v9_: _dafny.Seq
        d_39_v9_ = _dafny.SeqWithoutIsStrInference([d_38_v8_, False])
        d_40_v10_: _dafny.Seq
        d_40_v10_ = _dafny.SeqWithoutIsStrInference([d_39_v9_, d_39_v9_])
        d_41_globalState_: GlobalState
        nw2_ = GlobalState()
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([d_35_v2_ for d_42_i0_ in range(default__.abs(10))])).Elements:
                d_43_v5_: int = compr_2_
                if (d_43_v5_) in (_dafny.SeqWithoutIsStrInference([d_35_v2_ for d_42_i0_ in range(default__.abs(10))])):
                    coll2_[default__.safeModuloInt(d_43_v5_, len(d_37_v4_))] = len(_dafny.Map({not(False): d_35_v2_}))
            return _dafny.Map(coll2_)
        def iife3_():
            def iife5_():
                coll5_ = _dafny.Map()
                compr_5_: int
                for compr_5_ in _dafny.IntegerRange(827, 741):
                    d_44_v7_: int = compr_5_
                    if ((827) <= (d_44_v7_)) and ((d_44_v7_) < (741)):
                        coll5_[(d_44_v7_) + (d_35_v2_)] = False
                return _dafny.Map(coll5_)
            coll3_ = _dafny.Map()
            def iife4_():
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(827, 741):
                    d_44_v7_: int = compr_4_
                    if ((827) <= (d_44_v7_)) and ((d_44_v7_) < (741)):
                        coll4_[(d_44_v7_) + (d_35_v2_)] = False
                return _dafny.Map(coll4_)
            compr_3_: _dafny.Map
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([iife4_()
 for d_45_i1_ in range(default__.abs(-414))])).Elements:
                d_46_v6_: _dafny.Map = compr_3_
                if (d_46_v6_) in (_dafny.SeqWithoutIsStrInference([iife5_()
 for d_45_i1_ in range(default__.abs(-414))])):
                    coll3_[d_46_v6_] = d_35_v2_
            return _dafny.Map(coll3_)
        nw2_.ctor__(-585, d_36_v3_, 991, False, -718, False, True, False, 30, True, True, d_37_v4_, iife2_()
        , True, iife3_()
        , True, 897, d_40_v10_, True)
        d_41_globalState_ = nw2_
        d_47_v11_: D0
        d_47_v11_ = D0_DC0(d_35_v2_, (_dafny.SeqWithoutIsStrInference([d_38_v8_, d_38_v8_])) + (d_39_v9_))
        source2_ = d_47_v11_
        if source2_.is_DC0:
            d_48___mcc_h0_ = source2_.cf0
            d_49___mcc_h1_ = source2_.cf1
            d_50_cf1_ = d_49___mcc_h1_
            d_51_cf0_ = d_48___mcc_h0_
            (d_41_globalState_).f0 = (default__.safeDivisionInt(len(d_37_v4_), d_35_v2_)) - (d_35_v2_)
            d_33_v0_ = d_33_v0_
            (d_41_globalState_).f7 = not(not (d_38_v8_) or (not(d_38_v8_)))
            d_47_v11_ = d_47_v11_
        elif True:
            d_52___mcc_h2_ = source2_.cf2
            d_53_cf2_ = d_52___mcc_h2_
            d_54_v12_: _dafny.Array
            nw3_ = _dafny.Array(int(0), 10)
            d_54_v12_ = nw3_
            index2_ = default__.safeIndex(328, (d_54_v12_).length(0))
            (d_54_v12_)[index2_] = d_35_v2_
            d_55_v13_: _dafny.Array
            nw4_ = _dafny.Array(False, 16)
            d_55_v13_ = nw4_
            index3_ = default__.safeIndex(21, (d_55_v13_).length(0))
            (d_55_v13_)[index3_] = True
            d_56_v14_: _dafny.Map
            d_56_v14_ = _dafny.Map({d_54_v12_: (d_35_v2_) + (d_35_v2_)})
            d_57_v15_: _dafny.MultiSet
            d_57_v15_ = _dafny.MultiSet([d_38_v8_, d_38_v8_])
            index4_ = default__.safeIndex(328, (d_54_v12_).length(0))
            index5_ = default__.safeIndex(21, (d_55_v13_).length(0))
            rhs3_ = d_35_v2_
            rhs4_ = ((d_56_v14_)[d_54_v12_] if (d_54_v12_) in (d_56_v14_) else ((d_57_v15_)[d_38_v8_] if (d_38_v8_) in (d_57_v15_) else d_35_v2_))
            rhs5_ = d_38_v8_
            rhs6_ = (d_35_v2_) * (d_35_v2_)
            lhs2_ = d_41_globalState_
            lhs3_ = d_54_v12_
            lhs4_ = default__.safeIndex(328, (d_54_v12_).length(0))
            lhs5_ = d_55_v13_
            lhs6_ = default__.safeIndex(21, (d_55_v13_).length(0))
            lhs7_ = d_41_globalState_
            lhs2_.f16 = rhs3_
            lhs3_[lhs4_] = rhs4_
            lhs5_[lhs6_] = rhs5_
            lhs7_.f16 = rhs6_
            d_58_v16_: D0
            d_58_v16_ = D0_DC0(len(d_37_v4_), d_39_v9_)
            d_59_v17_: D0
            d_59_v17_ = D0_DC1(d_58_v16_)
            source3_ = d_59_v17_
            if source3_.is_DC0:
                d_60___mcc_h3_ = source3_.cf0
                d_61___mcc_h4_ = source3_.cf1
                d_62_cf1_ = d_61___mcc_h4_
                d_63_cf0_ = d_60___mcc_h3_
                d_64_v18_: _dafny.Array
                nw5_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_64_v18_ = nw5_
                d_64_v18_ = d_64_v18_
                d_65_v19_: _dafny.Array
                def lambda0_(d_66_v4_):
                    def lambda1_(d_67_i2_):
                        return d_66_v4_

                    return lambda1_

                init0_ = lambda0_(d_37_v4_)
                nw6_ = _dafny.Array(None, 8)
                for i0_0_ in range(nw6_.length(0)):
                    nw6_[i0_0_] = init0_(i0_0_)
                d_65_v19_ = nw6_
                index6_ = default__.safeIndex(779, (d_65_v19_).length(0))
                (d_65_v19_)[index6_] = d_37_v4_
                index7_ = default__.safeIndex(21, (d_55_v13_).length(0))
                index8_ = default__.safeIndex(779, (d_65_v19_).length(0))
                index9_ = default__.safeIndex(21, (d_55_v13_).length(0))
                index10_ = default__.safeIndex(328, (d_54_v12_).length(0))
                rhs7_ = default__.fm0((d_62_cf1_) <= (d_62_cf1_), d_37_v4_, d_41_globalState_)
                rhs8_ = ((d_55_v13_)[default__.safeIndex(21, (d_55_v13_).length(0))]) or (d_38_v8_)
                rhs9_ = (d_37_v4_).set(default__.safeIndex((d_57_v15_).cardinality, len(d_37_v4_)), (d_37_v4_)[default__.safeIndex((_dafny.MultiSet(d_39_v9_)).cardinality, len(d_37_v4_))])
                rhs10_ = (d_55_v13_)[default__.safeIndex(21, (d_55_v13_).length(0))]
                rhs11_ = (0) - ((d_47_v11_).cf0)
                lhs8_ = d_55_v13_
                lhs9_ = default__.safeIndex(21, (d_55_v13_).length(0))
                lhs10_ = d_65_v19_
                lhs11_ = default__.safeIndex(779, (d_65_v19_).length(0))
                lhs12_ = d_55_v13_
                lhs13_ = default__.safeIndex(21, (d_55_v13_).length(0))
                lhs14_ = d_54_v12_
                lhs15_ = default__.safeIndex(328, (d_54_v12_).length(0))
                d_63_cf0_ = rhs7_
                lhs8_[lhs9_] = rhs8_
                lhs10_[lhs11_] = rhs9_
                lhs12_[lhs13_] = rhs10_
                lhs14_[lhs15_] = rhs11_
                d_33_v0_ = d_33_v0_
                index11_ = default__.safeIndex(21, (d_55_v13_).length(0))
                (d_55_v13_)[index11_] = (True) in (d_62_cf1_)
            elif True:
                d_68___mcc_h5_ = source3_.cf2
                d_69_cf2_ = d_68___mcc_h5_
                d_38_v8_ = default__.fm1(d_37_v4_, (d_38_v8_) == (False), (d_54_v12_)[default__.safeIndex(328, (d_54_v12_).length(0))], d_41_globalState_)
                (d_41_globalState_).f0 = 273
                index12_ = default__.safeIndex(328, (d_54_v12_).length(0))
                (d_54_v12_)[index12_] = (0) - ((d_54_v12_)[default__.safeIndex(328, (d_54_v12_).length(0))])
                index13_ = default__.safeIndex(21, (d_55_v13_).length(0))
                (d_55_v13_)[index13_] = d_38_v8_
            d_70_v20_: _dafny.Map
            d_70_v20_ = _dafny.Map({((d_54_v12_)[default__.safeIndex(328, (d_54_v12_).length(0))]) + ((d_54_v12_)[default__.safeIndex(328, (d_54_v12_).length(0))]): (d_55_v13_)[default__.safeIndex(21, (d_55_v13_).length(0))]})
            d_70_v20_ = (d_70_v20_).set(default__.fm0((d_55_v13_)[default__.safeIndex(21, (d_55_v13_).length(0))], d_37_v4_, d_41_globalState_), False)
            (d_41_globalState_).f6 = default__.fm1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjtqkbv")), d_38_v8_, (d_54_v12_)[default__.safeIndex(328, (d_54_v12_).length(0))], d_41_globalState_)
        d_35_v2_ = (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyv"))) + (d_37_v4_))) + (default__.safeModuloInt(len(d_37_v4_), len(d_37_v4_)))
        hi0_ = (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_38_v8_]) if d_38_v8_ else d_39_v9_))).cardinality
        for d_71_i3_ in range(d_35_v2_, hi0_):
            d_72_v21_: _dafny.Set
            d_72_v21_ = _dafny.Set({d_71_i3_})
            d_73_v22_: _dafny.Map
            d_73_v22_ = _dafny.Map({d_72_v21_: d_38_v8_})
            d_35_v2_ = len(((d_73_v22_) | (d_73_v22_) if d_38_v8_ else d_73_v22_))
            d_74_v23_: _dafny.Map
            d_74_v23_ = _dafny.Map({d_38_v8_: 109})
            d_75_v24_: _dafny.Map
            d_75_v24_ = _dafny.Map({d_33_v0_: (d_74_v23_).set(d_38_v8_, d_35_v2_)})
            d_76_v25_: _dafny.Array
            nw7_ = _dafny.Array(int(0), 3)
            d_76_v25_ = nw7_
            d_77_v26_: _dafny.Array
            nw8_ = _dafny.Array(False, 15)
            d_77_v26_ = nw8_
            index14_ = default__.safeIndex(922, (d_77_v26_).length(0))
            (d_77_v26_)[index14_] = d_38_v8_
            index15_ = default__.safeIndex(922, (d_77_v26_).length(0))
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: str
                for compr_6_ in (d_34_v1_).Elements:
                    d_78_v27_: str = compr_6_
                    if (d_78_v27_) in (d_34_v1_):
                        coll6_[d_78_v27_] = d_74_v23_
                return _dafny.Map(coll6_)
            rhs12_ = ((iife6_()
             if d_38_v8_ else d_75_v24_)) | (d_75_v24_)
            rhs13_ = d_76_v25_
            rhs14_ = (d_71_i3_) - ((d_71_i3_) + ((0) - (d_71_i3_)))
            rhs15_ = d_38_v8_
            rhs16_ = not(default__.fm1(d_37_v4_, d_38_v8_, d_35_v2_, d_41_globalState_))
            lhs16_ = d_41_globalState_
            lhs17_ = d_41_globalState_
            lhs18_ = d_77_v26_
            lhs19_ = default__.safeIndex(922, (d_77_v26_).length(0))
            d_75_v24_ = rhs12_
            d_76_v25_ = rhs13_
            lhs16_.f16 = rhs14_
            lhs17_.f6 = rhs15_
            lhs18_[lhs19_] = rhs16_
            (d_41_globalState_).f0 = d_71_i3_
            if not((d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))]):
                d_79_v28_: bool
                out0_: bool
                out0_ = default__.m0(not(d_38_v8_), (d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))], not(d_38_v8_), default__.fm1(d_37_v4_, (d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))], d_71_i3_, d_41_globalState_), d_41_globalState_)
                d_79_v28_ = out0_
                d_76_v25_ = d_76_v25_
                index16_ = default__.safeIndex(922, (d_77_v26_).length(0))
                (d_77_v26_)[index16_] = d_79_v28_
                index17_ = default__.safeIndex(657, (d_76_v25_).length(0))
                (d_76_v25_)[index17_] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcuajj")))
                d_80_v29_: _dafny.Array
                nw9_ = _dafny.Array(_dafny.Seq({}), 3)
                d_80_v29_ = nw9_
                d_81_v30_: _dafny.Set
                d_81_v30_ = _dafny.Set({d_79_v28_})
                index18_ = default__.safeIndex(657, (d_76_v25_).length(0))
                rhs17_ = d_71_i3_
                rhs18_ = not(default__.fm1((default__.fm2(d_38_v8_, d_38_v8_, d_35_v2_, d_81_v30_, d_41_globalState_)) + (d_37_v4_), (d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))], d_35_v2_, d_41_globalState_))
                rhs19_ = d_80_v29_
                lhs20_ = d_76_v25_
                lhs21_ = default__.safeIndex(657, (d_76_v25_).length(0))
                lhs22_ = d_41_globalState_
                lhs20_[lhs21_] = rhs17_
                lhs22_.f6 = rhs18_
                d_80_v29_ = rhs19_
                d_82_v31_: _dafny.Array
                nw10_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                d_82_v31_ = nw10_
                index19_ = default__.safeIndex(674, (d_82_v31_).length(0))
                (d_82_v31_)[index19_] = d_37_v4_
                d_83_v32_: _dafny.MultiSet
                d_83_v32_ = _dafny.MultiSet([False])
                index20_ = default__.safeIndex(657, (d_76_v25_).length(0))
                index21_ = default__.safeIndex(922, (d_77_v26_).length(0))
                index22_ = default__.safeIndex(674, (d_82_v31_).length(0))
                rhs20_ = (0) - ((default__.safeDivisionInt(d_71_i3_, (d_76_v25_)[default__.safeIndex(657, (d_76_v25_).length(0))])) + (d_71_i3_))
                rhs21_ = (d_83_v32_).isdisjoint(d_83_v32_)
                rhs22_ = _dafny.SeqWithoutIsStrInference([d_33_v0_ for d_84_i4_ in range(default__.abs(-715))])
                lhs23_ = d_76_v25_
                lhs24_ = default__.safeIndex(657, (d_76_v25_).length(0))
                lhs25_ = d_77_v26_
                lhs26_ = default__.safeIndex(922, (d_77_v26_).length(0))
                lhs27_ = d_82_v31_
                lhs28_ = default__.safeIndex(674, (d_82_v31_).length(0))
                lhs23_[lhs24_] = rhs20_
                lhs25_[lhs26_] = rhs21_
                lhs27_[lhs28_] = rhs22_
            elif True:
                d_74_v23_ = (d_74_v23_) | (d_74_v23_)
                d_85_v33_: D0
                d_85_v33_ = D0_DC0(d_35_v2_, d_39_v9_)
                d_86_v34_: D0
                d_86_v34_ = D0_DC1(d_85_v33_)
                d_87_v35_: D0
                d_87_v35_ = D0_DC1(d_86_v34_)
                d_88_v36_: D0
                d_88_v36_ = D0_DC1(d_86_v34_)
                d_89_v37_: D0
                d_89_v37_ = D0_DC1(d_87_v35_)
                d_90_v38_: D0
                d_90_v38_ = D0_DC1(d_88_v36_)
                d_91_v39_: _dafny.MultiSet
                d_91_v39_ = _dafny.MultiSet([d_90_v38_, D0_DC1(d_88_v36_)])
                d_92_v40_: _dafny.Seq
                d_92_v40_ = _dafny.SeqWithoutIsStrInference([d_90_v38_, d_90_v38_, d_90_v38_])
                d_38_v8_ = not(((d_91_v39_) | (_dafny.MultiSet(d_92_v40_))).isdisjoint((d_91_v39_).set(D0_DC1(d_89_v37_), default__.abs(646))))
                d_93_v41_: _dafny.MultiSet
                d_93_v41_ = _dafny.MultiSet([d_37_v4_])
                index23_ = default__.safeIndex(922, (d_77_v26_).length(0))
                (d_77_v26_)[index23_] = (d_93_v41_).issubset((_dafny.MultiSet([d_37_v4_])).set(d_37_v4_, default__.abs(len(d_72_v21_))))
                rhs23_ = d_35_v2_
                rhs24_ = d_90_v38_
                lhs29_ = d_41_globalState_
                lhs29_.f16 = rhs23_
                d_90_v38_ = rhs24_
                d_94_v42_: _dafny.Map
                d_94_v42_ = _dafny.Map({(0) - (d_35_v2_): d_74_v23_})
                d_95_v43_: _dafny.Set
                d_95_v43_ = _dafny.Set({(d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))], (d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))]})
                d_96_v44_: _dafny.Map
                d_96_v44_ = _dafny.Map({(d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))]: d_94_v42_})
                d_97_v46_: _dafny.Map
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in _dafny.IntegerRange(542, 652):
                        d_98_v45_: int = compr_7_
                        if ((542) <= (d_98_v45_)) and ((d_98_v45_) < (652)):
                            coll7_[(d_98_v45_) * (d_71_i3_)] = d_74_v23_
                    return _dafny.Map(coll7_)
                d_97_v46_ = _dafny.Map({d_39_v9_: iife7_()
                })
                rhs25_ = (d_95_v43_) != (d_95_v43_)
                rhs26_ = ((d_94_v42_) | (((d_96_v44_)[(d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))]] if ((d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))]) in (d_96_v44_) else ((d_97_v46_)[(d_39_v9_).set(default__.safeIndex(d_35_v2_, len(d_39_v9_)), (d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))])] if ((d_39_v9_).set(default__.safeIndex(d_35_v2_, len(d_39_v9_)), (d_77_v26_)[default__.safeIndex(922, (d_77_v26_).length(0))])) in (d_97_v46_) else d_94_v42_)))) | (_dafny.Map({d_35_v2_: d_74_v23_}))
                lhs30_ = d_41_globalState_
                lhs30_.f7 = rhs25_
                d_94_v42_ = rhs26_
        d_38_v8_ = (979) < (283)
        d_99_v47_: _dafny.Array
        nw11_ = _dafny.Array(False, 27)
        d_99_v47_ = nw11_
        d_100_v48_: D0
        d_100_v48_ = D0_DC0(636, default__.fm3(569, len(d_36_v3_), d_38_v8_, d_41_globalState_))
        d_101_v49_: D0
        d_101_v49_ = D0_DC1(d_100_v48_)
        d_102_v50_: D0
        d_102_v50_ = D0_DC1(d_101_v49_)
        d_103_v51_: _dafny.Map
        d_103_v51_ = _dafny.Map({d_33_v0_: d_35_v2_})
        d_104_v52_: _dafny.MultiSet
        d_104_v52_ = _dafny.MultiSet([d_99_v47_])
        rhs27_ = d_99_v47_
        rhs28_ = default__.fm4(d_103_v51_, d_41_globalState_)
        rhs29_ = d_38_v8_
        rhs30_ = (_dafny.MultiSet([d_99_v47_, d_99_v47_, d_99_v47_])).isdisjoint((d_104_v52_).intersection(d_104_v52_))
        lhs31_ = d_41_globalState_
        d_99_v47_ = rhs27_
        d_102_v50_ = rhs28_
        lhs31_.f7 = rhs29_
        d_38_v8_ = rhs30_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_99_v47_).length(0)):
            d_105_i5_: int = guard_loop_0_
            if (True) and (((0) <= (d_105_i5_)) and ((d_105_i5_) < ((d_99_v47_).length(0)))):
                (d_99_v47_)[(d_105_i5_)] = d_38_v8_
        d_106_v53_: _dafny.Array
        nw12_ = _dafny.Array(None, 8)
        nw12_[int(0)] = d_35_v2_
        nw12_[int(1)] = d_35_v2_
        nw12_[int(2)] = (721) * (d_35_v2_)
        nw12_[int(3)] = 1
        nw12_[int(4)] = (0) - (d_35_v2_)
        nw12_[int(5)] = d_35_v2_
        nw12_[int(6)] = default__.fm0(d_38_v8_, d_37_v4_, d_41_globalState_)
        nw12_[int(7)] = d_35_v2_
        d_106_v53_ = nw12_
        index24_ = default__.safeIndex(741, (d_106_v53_).length(0))
        (d_106_v53_)[index24_] = (d_35_v2_) * (d_35_v2_)
        index25_ = default__.safeIndex(539, (d_106_v53_).length(0))
        (d_106_v53_)[index25_] = len(d_37_v4_)
        d_107_v54_: _dafny.MultiSet
        d_107_v54_ = _dafny.MultiSet([d_38_v8_])
        d_108_v55_: _dafny.Map
        d_108_v55_ = _dafny.Map({d_107_v54_: d_99_v47_})
        index26_ = default__.safeIndex(741, (d_106_v53_).length(0))
        index27_ = default__.safeIndex(539, (d_106_v53_).length(0))
        rhs31_ = default__.fm0(d_38_v8_, d_37_v4_, d_41_globalState_)
        rhs32_ = (default__.fm5(d_38_v8_, d_33_v0_, d_35_v2_, d_38_v8_, d_41_globalState_)).cf3
        rhs33_ = (0) - ((len(d_108_v55_)) - (-525))
        rhs34_ = not(d_38_v8_)
        lhs32_ = d_106_v53_
        lhs33_ = default__.safeIndex(741, (d_106_v53_).length(0))
        lhs34_ = d_106_v53_
        lhs35_ = default__.safeIndex(539, (d_106_v53_).length(0))
        lhs36_ = d_41_globalState_
        lhs32_[lhs33_] = rhs31_
        d_33_v0_ = rhs32_
        lhs34_[lhs35_] = rhs33_
        lhs36_.f6 = rhs34_
        if ((d_38_v8_) and (d_38_v8_)) in (d_107_v54_):
            d_109_v56_: C1
            nw13_ = C1()
            nw13_.ctor__(_dafny.Map({not(True): len(default__.fm9((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], d_38_v8_, False, len(_dafny.SeqWithoutIsStrInference([(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))] for d_110_i6_ in range(default__.abs(282))])), d_41_globalState_))}), d_102_v50_)
            d_109_v56_ = nw13_
            d_111_v57_: D2
            d_111_v57_ = D2_DC8(d_38_v8_)
            d_112_v58_: _dafny.Map
            d_112_v58_ = _dafny.Map({d_111_v57_: default__.fm1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uteyy")), d_38_v8_, (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], d_41_globalState_)})
            d_113_v59_: _dafny.MultiSet
            d_113_v59_ = _dafny.MultiSet([d_35_v2_])
            (d_41_globalState_).f6 = not((default__.fm10(d_112_v58_, d_38_v8_, d_113_v59_, d_41_globalState_)).isdisjoint(_dafny.Set({d_47_v11_})))
            if (True) and (d_38_v8_):
                (d_41_globalState_).f7 = ((d_107_v54_).cardinality) <= ((d_35_v2_) + (d_35_v2_))
                rhs35_ = ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) == (len((d_37_v4_).set(default__.safeIndex((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], len(d_37_v4_)), d_33_v0_)))
                rhs36_ = (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]
                d_38_v8_ = rhs35_
                d_35_v2_ = rhs36_
                d_37_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsbt"))
                d_38_v8_ = False
                d_114_v60_: C0
                nw14_ = C0()
                nw14_.ctor__()
                d_114_v60_ = nw14_
            elif True:
                d_115_v61_: _dafny.Array
                nw15_ = _dafny.Array(_dafny.Map({}), 6)
                d_115_v61_ = nw15_
                rhs37_ = not(d_38_v8_)
                rhs38_ = default__.safeModuloInt((d_35_v2_ if d_38_v8_ else d_35_v2_), (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])
                rhs39_ = (d_109_v56_).fm6((d_37_v4_)[default__.safeIndex(d_35_v2_, len(d_37_v4_))], d_38_v8_, not(d_38_v8_), d_38_v8_, d_41_globalState_)
                rhs40_ = d_115_v61_
                lhs37_ = d_41_globalState_
                lhs38_ = d_41_globalState_
                lhs39_ = d_41_globalState_
                lhs37_.f6 = rhs37_
                lhs38_.f16 = rhs38_
                lhs39_.f16 = rhs39_
                d_115_v61_ = rhs40_
                d_116_v62_: _dafny.Seq
                d_116_v62_ = _dafny.SeqWithoutIsStrInference([d_37_v4_])
                (d_41_globalState_).f6 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "loukb"))) in (d_116_v62_)
                d_117_v63_: _dafny.Set
                d_117_v63_ = _dafny.Set({811, len(_dafny.Set({d_38_v8_, d_38_v8_})), (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], ((d_109_v56_.f19)[d_38_v8_] if (d_38_v8_) in (d_109_v56_.f19) else d_35_v2_), d_35_v2_})
                d_118_v64_: _dafny.Array
                d_119_v65_: _dafny.Map
                out1_: _dafny.Array
                out2_: _dafny.Map
                out1_, out2_ = (d_109_v56_).m1(2, (_dafny.Set({d_35_v2_})).intersection(d_117_v63_), d_35_v2_, d_41_globalState_)
                d_118_v64_ = out1_
                d_119_v65_ = out2_
                d_120_v66_: _dafny.Seq
                d_120_v66_ = _dafny.SeqWithoutIsStrInference([d_117_v63_, d_117_v63_])
                d_121_v68_: _dafny.Array
                d_122_v69_: _dafny.Map
                out3_: _dafny.Array
                out4_: _dafny.Map
                def iife8_():
                    coll8_ = _dafny.Set()
                    compr_8_: int
                    for compr_8_ in _dafny.IntegerRange(913, -188):
                        d_123_v67_: int = compr_8_
                        if ((913) <= (d_123_v67_)) and ((d_123_v67_) < (-188)):
                            coll8_ = coll8_.union(_dafny.Set([(d_123_v67_) + ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])]))
                    return _dafny.Set(coll8_)
                out3_, out4_ = (d_109_v56_).m1(default__.safeModuloInt(d_35_v2_, 789), ((d_120_v66_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]])), len(d_120_v66_))]) | (_dafny.Set({-395, len(iife8_()
                )})), (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], d_41_globalState_)
                d_121_v68_ = out3_
                d_122_v69_ = out4_
                index28_ = default__.safeIndex(322, (d_121_v68_).length(0))
                (d_121_v68_)[index28_] = (d_38_v8_) == (d_38_v8_)
                index29_ = default__.safeIndex(322, (d_121_v68_).length(0))
                (d_121_v68_)[index29_] = d_38_v8_
            (d_41_globalState_).f16 = (0) - ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])
            d_124_v70_: C1
            nw16_ = C1()
            nw16_.ctor__(d_109_v56_.f19, D0_DC1(D0_DC0(len(_dafny.SeqWithoutIsStrInference([d_33_v0_ for d_125_i7_ in range(default__.abs(352))])), d_39_v9_)))
            d_124_v70_ = nw16_
            d_124_v70_ = d_124_v70_
        elif True:
            (d_41_globalState_).f7 = d_38_v8_
            d_126_v71_: D2
            d_126_v71_ = D2_DC7(d_33_v0_)
            pat_let_tv0_ = d_33_v0_
            def iife9_(_pat_let0_0):
                def iife10_(d_127_dt__update__tmp_h0_):
                    def iife11_(_pat_let1_0):
                        def iife12_(d_128_dt__update_hcf12_h0_):
                            return D2_DC7(d_128_dt__update_hcf12_h0_)
                        return iife12_(_pat_let1_0)
                    return iife11_(pat_let_tv0_)
                return iife10_(_pat_let0_0)
            d_126_v71_ = iife9_(D2_DC7(_dafny.CodePoint('q')))
            d_129_v72_: _dafny.Array
            nw17_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 18)
            d_129_v72_ = nw17_
            d_130_v73_: _dafny.Seq
            d_130_v73_ = _dafny.SeqWithoutIsStrInference([d_129_v72_, d_129_v72_])
            d_129_v72_ = (d_130_v73_)[default__.safeIndex(d_35_v2_, len(d_130_v73_))]
            source4_ = D2_DC7(d_33_v0_)
            if source4_.is_DC7:
                d_131___mcc_h6_ = source4_.cf12
                d_132_cf12_ = d_131___mcc_h6_
                index30_ = default__.safeIndex(741, (d_106_v53_).length(0))
                (d_106_v53_)[index30_] = (0) - ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])
                d_133_v74_: _dafny.Map
                d_133_v74_ = _dafny.Map({d_35_v2_: d_36_v3_})
                d_134_v75_: _dafny.Map
                d_134_v75_ = _dafny.Map({(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]: (d_36_v3_) + (((d_133_v74_)[(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]] if ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) in (d_133_v74_) else d_36_v3_))})
                d_134_v75_ = (d_134_v75_).set(d_35_v2_, _dafny.SeqWithoutIsStrInference([(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))] for d_135_i8_ in range(default__.abs(999))]))
                d_136_v76_: C0
                nw18_ = C0()
                nw18_.ctor__()
                d_136_v76_ = nw18_
                d_137_v77_: _dafny.MultiSet
                d_137_v77_ = _dafny.MultiSet([d_136_v76_])
                d_138_v78_: bool
                out5_: bool
                out5_ = default__.m0(d_38_v8_, (d_35_v2_) != ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]), (d_137_v77_).isdisjoint(d_137_v77_), d_38_v8_, d_41_globalState_)
                d_138_v78_ = out5_
                d_139_v79_: _dafny.Map
                d_139_v79_ = _dafny.Map({d_138_v78_: len(_dafny.SeqWithoutIsStrInference([d_33_v0_ for d_140_i9_ in range(default__.abs(8))]))})
                d_141_v80_: C1
                nw19_ = C1()
                nw19_.ctor__(d_139_v79_, d_102_v50_)
                d_141_v80_ = nw19_
                d_142_v81_: _dafny.MultiSet
                d_142_v81_ = _dafny.MultiSet([d_141_v80_, d_141_v80_, d_141_v80_])
                d_143_v82_: _dafny.Set
                d_143_v82_ = _dafny.Set({(d_142_v81_).cardinality})
                d_144_v83_: _dafny.Map
                d_144_v83_ = _dafny.Map({(d_38_v8_ if d_138_v78_ else not(d_138_v78_)): d_143_v82_})
                d_144_v83_ = d_144_v83_
            elif source4_.is_DC8:
                d_145___mcc_h7_ = source4_.cf13
                d_146_cf13_ = d_145___mcc_h7_
                d_147_v84_: _dafny.Map
                d_147_v84_ = _dafny.Map({d_38_v8_: d_35_v2_})
                d_148_v85_: C1
                nw20_ = C1()
                nw20_.ctor__(d_147_v84_, D0_DC1(D0_DC1(d_100_v48_)))
                d_148_v85_ = nw20_
                d_149_v86_: _dafny.Seq
                d_149_v86_ = _dafny.SeqWithoutIsStrInference([d_148_v85_])
                d_150_v87_: bool
                out6_: bool
                out6_ = default__.m0(d_38_v8_, d_38_v8_, d_146_cf13_, default__.fm1(d_37_v4_, d_38_v8_, len((D4_DC11(d_149_v86_)).cf16), d_41_globalState_), d_41_globalState_)
                d_150_v87_ = out6_
                (d_41_globalState_).f16 = (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]
                d_33_v0_ = d_33_v0_
                d_151_v88_: C0
                nw21_ = C0()
                nw21_.ctor__()
                d_151_v88_ = nw21_
            elif True:
                d_152___mcc_h8_ = source4_.cf11
                d_153_cf11_ = d_152___mcc_h8_
                d_154_v89_: D1
                d_154_v89_ = D1_DC3(d_35_v2_)
                d_154_v89_ = D1_DC3((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])
                d_38_v8_ = (d_38_v8_) or (d_38_v8_)
                d_155_v90_: _dafny.Seq
                d_155_v90_ = _dafny.SeqWithoutIsStrInference([d_106_v53_, d_106_v53_])
                d_156_v91_: _dafny.Map
                d_156_v91_ = _dafny.Map({d_38_v8_: -63})
                d_157_v92_: C1
                nw22_ = C1()
                nw22_.ctor__(d_156_v91_, d_102_v50_)
                d_157_v92_ = nw22_
                d_158_v93_: _dafny.Seq
                d_158_v93_ = _dafny.SeqWithoutIsStrInference([d_106_v53_, d_106_v53_, (d_155_v90_)[default__.safeIndex(d_35_v2_, len(d_155_v90_))], (D5_DC17(d_107_v54_, d_106_v53_, d_157_v92_, len(d_37_v4_), (d_37_v4_).set(default__.safeIndex(d_35_v2_, len(d_37_v4_)), d_33_v0_))).cf25, d_106_v53_])
                rhs41_ = d_35_v2_
                rhs42_ = default__.fm11(d_41_globalState_)
                rhs43_ = (d_158_v93_)[default__.safeIndex(d_35_v2_, len(d_158_v93_))]
                lhs40_ = d_41_globalState_
                lhs40_.f0 = rhs41_
                d_33_v0_ = rhs42_
                d_106_v53_ = rhs43_
                d_159_v94_: _dafny.Map
                d_159_v94_ = _dafny.Map({d_38_v8_: D0_DC1(d_100_v48_)})
                d_159_v94_ = (d_159_v94_).set(d_38_v8_, (d_102_v50_ if d_38_v8_ else d_102_v50_))
            d_160_v95_: bool
            out7_: bool
            out7_ = default__.m0(d_38_v8_, d_38_v8_, d_38_v8_, (d_38_v8_) or (not(d_38_v8_)), d_41_globalState_)
            d_160_v95_ = out7_
        index31_ = default__.safeIndex(938, (d_99_v47_).length(0))
        (d_99_v47_)[index31_] = d_38_v8_
        d_161_v96_: _dafny.Map
        d_161_v96_ = _dafny.Map({d_38_v8_: 676})
        d_162_v97_: C1
        nw23_ = C1()
        nw23_.ctor__(d_161_v96_, D0_DC1(d_100_v48_))
        d_162_v97_ = nw23_
        d_163_v98_: _dafny.Seq
        d_163_v98_ = _dafny.SeqWithoutIsStrInference([d_162_v97_])
        d_164_v99_: D4
        d_164_v99_ = D4_DC11(d_163_v98_)
        d_165_v100_: D4
        d_165_v100_ = D4_DC14(d_164_v99_)
        pat_let_tv1_ = d_164_v99_
        d_166_v101_: _dafny.Map
        def iife13_(_pat_let2_0):
            def iife14_(d_167_dt__update__tmp_h1_):
                def iife15_(_pat_let3_0):
                    def iife16_(d_168_dt__update_hcf21_h0_):
                        return D4_DC14(d_168_dt__update_hcf21_h0_)
                    return iife16_(_pat_let3_0)
                return iife15_(pat_let_tv1_)
            return iife14_(_pat_let2_0)
        d_166_v101_ = _dafny.Map({False: iife13_(d_165_v100_)})
        d_169_v102_: _dafny.Seq
        d_169_v102_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_38_v8_: D4_DC14(d_164_v99_)}), _dafny.Map({d_38_v8_: d_165_v100_}), (d_166_v101_).set(True, D4_DC14(D4_DC12()))])
        d_170_v103_: _dafny.MultiSet
        d_170_v103_ = _dafny.MultiSet([(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]])
        d_171_v104_: _dafny.Map
        d_171_v104_ = _dafny.Map({d_38_v8_: d_170_v103_})
        index32_ = default__.safeIndex(938, (d_99_v47_).length(0))
        index33_ = default__.safeIndex(741, (d_106_v53_).length(0))
        index34_ = default__.safeIndex(741, (d_106_v53_).length(0))
        rhs44_ = not (not((d_170_v103_).ispropersubset(((d_171_v104_)[d_38_v8_] if (d_38_v8_) in (d_171_v104_) else default__.fm12(238, d_41_globalState_))))) or (d_38_v8_)
        rhs45_ = (d_169_v102_) + (d_169_v102_)
        rhs46_ = d_38_v8_
        rhs47_ = default__.fm0((d_38_v8_ if d_38_v8_ else d_38_v8_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrcu")), d_41_globalState_)
        rhs48_ = ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) + (default__.safeModuloInt((d_162_v97_).fm6(d_33_v0_, d_38_v8_, d_38_v8_, d_38_v8_, d_41_globalState_), (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]))
        lhs41_ = d_99_v47_
        lhs42_ = default__.safeIndex(938, (d_99_v47_).length(0))
        lhs43_ = d_41_globalState_
        lhs44_ = d_106_v53_
        lhs45_ = default__.safeIndex(741, (d_106_v53_).length(0))
        lhs46_ = d_106_v53_
        lhs47_ = default__.safeIndex(741, (d_106_v53_).length(0))
        lhs41_[lhs42_] = rhs44_
        d_169_v102_ = rhs45_
        lhs43_.f6 = rhs46_
        lhs44_[lhs45_] = rhs47_
        lhs46_[lhs47_] = rhs48_
        d_172_v105_: _dafny.Seq
        d_172_v105_ = _dafny.SeqWithoutIsStrInference([d_99_v47_])
        d_173_v106_: _dafny.Set
        d_173_v106_ = _dafny.Set({True, (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]})
        rhs49_ = default__.safeDivisionInt(d_35_v2_, (d_35_v2_) * (d_35_v2_))
        rhs50_ = (d_36_v3_) <= (d_36_v3_)
        rhs51_ = (d_35_v2_) <= (default__.safeModuloInt((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], ((d_162_v97_.f19)[not((d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))])] if (not((d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))])) in (d_162_v97_.f19) else len(d_161_v96_))))
        rhs52_ = default__.fm2((d_172_v105_) < (d_172_v105_), not(False), (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], (d_173_v106_) | (d_173_v106_), d_41_globalState_)
        rhs53_ = (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]
        lhs48_ = d_41_globalState_
        lhs49_ = d_41_globalState_
        lhs50_ = d_41_globalState_
        lhs48_.f0 = rhs49_
        d_38_v8_ = rhs50_
        lhs49_.f7 = rhs51_
        d_37_v4_ = rhs52_
        lhs50_.f6 = rhs53_
        index35_ = default__.safeIndex(741, (d_106_v53_).length(0))
        (d_106_v53_)[index35_] = ((len(d_173_v106_)) - ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) if (True) and (d_38_v8_) else d_35_v2_)
        d_174_v107_: _dafny.Map
        d_174_v107_ = _dafny.Map({(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]: (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]})
        d_175_v108_: _dafny.Seq
        d_175_v108_ = _dafny.SeqWithoutIsStrInference([d_106_v53_])
        d_176_v109_: C0
        nw24_ = C0()
        nw24_.ctor__()
        d_176_v109_ = nw24_
        d_177_v110_: _dafny.Map
        d_177_v110_ = _dafny.Map({d_176_v109_: (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]})
        d_178_v111_: _dafny.Map
        d_178_v111_ = _dafny.Map({d_177_v110_: d_38_v8_})
        d_179_v112_: D5
        d_179_v112_ = D5_DC17(_dafny.MultiSet([d_38_v8_]), d_106_v53_, d_162_v97_, len(d_39_v9_), d_37_v4_)
        d_180_v113_: _dafny.Array
        nw25_ = _dafny.Array(None, 17)
        nw25_[int(0)] = d_106_v53_
        nw25_[int(1)] = d_106_v53_
        nw25_[int(2)] = d_106_v53_
        nw25_[int(3)] = ((d_175_v108_)[default__.safeIndex((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], len(d_175_v108_))] if (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))] else d_106_v53_)
        nw25_[int(4)] = d_106_v53_
        nw25_[int(5)] = d_106_v53_
        nw25_[int(6)] = d_106_v53_
        nw25_[int(7)] = d_106_v53_
        nw25_[int(8)] = d_106_v53_
        nw25_[int(9)] = d_106_v53_
        nw25_[int(10)] = d_106_v53_
        nw25_[int(11)] = d_106_v53_
        nw25_[int(12)] = (d_175_v108_)[default__.safeIndex(len(d_178_v111_), len(d_175_v108_))]
        nw25_[int(13)] = (d_179_v112_).cf25
        nw25_[int(14)] = d_106_v53_
        nw25_[int(15)] = d_106_v53_
        nw25_[int(16)] = d_106_v53_
        d_180_v113_ = nw25_
        index36_ = default__.safeIndex(352, (d_180_v113_).length(0))
        (d_180_v113_)[index36_] = d_106_v53_
        d_181_v114_: _dafny.Map
        d_181_v114_ = _dafny.Map({d_99_v47_: d_35_v2_})
        index37_ = default__.safeIndex(352, (d_180_v113_).length(0))
        rhs54_ = ((d_35_v2_) - (d_35_v2_)) != ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])
        rhs55_ = _dafny.Map({len(d_37_v4_): (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]})
        rhs56_ = d_106_v53_
        rhs57_ = (d_181_v114_).set(d_99_v47_, d_35_v2_)
        rhs58_ = (d_33_v0_) in (d_37_v4_)
        lhs51_ = d_41_globalState_
        lhs52_ = d_180_v113_
        lhs53_ = default__.safeIndex(352, (d_180_v113_).length(0))
        lhs54_ = d_41_globalState_
        lhs51_.f7 = rhs54_
        d_174_v107_ = rhs55_
        lhs52_[lhs53_] = rhs56_
        d_181_v114_ = rhs57_
        lhs54_.f6 = rhs58_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, ((d_180_v113_)[default__.safeIndex(352, (d_180_v113_).length(0))]).length(0)):
            d_182_i10_: int = guard_loop_1_
            if (True) and (((0) <= (d_182_i10_)) and ((d_182_i10_) < (((d_180_v113_)[default__.safeIndex(352, (d_180_v113_).length(0))]).length(0)))):
                _ingredientsd_0.append(((d_180_v113_)[default__.safeIndex(352, (d_180_v113_).length(0))], int(d_182_i10_), (d_182_i10_) - ((0) - (d_35_v2_))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        index38_ = default__.safeIndex(352, (d_180_v113_).length(0))
        rhs59_ = (d_180_v113_)[default__.safeIndex(352, (d_180_v113_).length(0))]
        rhs60_ = d_38_v8_
        lhs55_ = d_180_v113_
        lhs56_ = default__.safeIndex(352, (d_180_v113_).length(0))
        lhs55_[lhs56_] = rhs59_
        d_38_v8_ = rhs60_
        if (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]:
            d_183_v115_: _dafny.Map
            d_183_v115_ = _dafny.Map({d_37_v4_: (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]})
            d_184_v116_: _dafny.Seq
            d_184_v116_ = _dafny.SeqWithoutIsStrInference([d_37_v4_, (d_37_v4_).set(default__.safeIndex(((d_183_v115_)[d_37_v4_] if (d_37_v4_) in (d_183_v115_) else d_35_v2_), len(d_37_v4_)), d_33_v0_)])
            index39_ = default__.safeIndex(741, (d_106_v53_).length(0))
            rhs61_ = d_184_v116_
            rhs62_ = d_36_v3_
            rhs63_ = (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]
            lhs57_ = d_106_v53_
            lhs58_ = default__.safeIndex(741, (d_106_v53_).length(0))
            d_184_v116_ = rhs61_
            d_36_v3_ = rhs62_
            lhs57_[lhs58_] = rhs63_
            d_185_v117_: D2
            d_185_v117_ = D2_DC7(d_33_v0_)
            source5_ = d_185_v117_
            if source5_.is_DC7:
                d_186___mcc_h9_ = source5_.cf12
                d_187_cf12_ = d_186___mcc_h9_
                index40_ = default__.safeIndex(938, (d_99_v47_).length(0))
                (d_99_v47_)[index40_] = (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]
                (d_41_globalState_).f0 = d_35_v2_
                d_188_v118_: _dafny.Array
                def lambda2_(d_189_v4_):
                    def lambda3_(d_190_i11_):
                        return (d_189_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymqrueuc")))

                    return lambda3_

                init1_ = lambda2_(d_37_v4_)
                nw26_ = _dafny.Array(None, 19)
                for i0_1_ in range(nw26_.length(0)):
                    nw26_[i0_1_] = init1_(i0_1_)
                d_188_v118_ = nw26_
                index41_ = default__.safeIndex(944, (d_188_v118_).length(0))
                (d_188_v118_)[index41_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgxpoeh"))
                index42_ = default__.safeIndex(944, (d_188_v118_).length(0))
                (d_188_v118_)[index42_] = _dafny.SeqWithoutIsStrInference([d_187_cf12_ for d_191_i12_ in range(default__.abs(-478))])
                rhs64_ = d_176_v109_
                rhs65_ = (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]
                lhs59_ = d_41_globalState_
                d_176_v109_ = rhs64_
                lhs59_.f0 = rhs65_
            elif source5_.is_DC8:
                d_192___mcc_h10_ = source5_.cf13
                d_193_cf13_ = d_192___mcc_h10_
                d_194_v119_: _dafny.Array
                nw27_ = _dafny.Array(None, 10)
                nw27_[int(0)] = d_37_v4_
                nw27_[int(1)] = _dafny.SeqWithoutIsStrInference([d_33_v0_ for d_195_i13_ in range(default__.abs(997))])
                nw27_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "levgfjrvj"))
                nw27_[int(3)] = d_37_v4_
                nw27_[int(4)] = d_37_v4_
                nw27_[int(5)] = (d_37_v4_).set(default__.safeIndex(d_35_v2_, len(d_37_v4_)), _dafny.CodePoint('v'))
                nw27_[int(6)] = d_37_v4_
                nw27_[int(7)] = d_37_v4_
                nw27_[int(8)] = d_37_v4_
                nw27_[int(9)] = ((d_184_v116_)[default__.safeIndex(d_35_v2_, len(d_184_v116_))]) + (d_37_v4_)
                d_194_v119_ = nw27_
                index43_ = default__.safeIndex(938, (d_194_v119_).length(0))
                (d_194_v119_)[index43_] = d_37_v4_
                index44_ = default__.safeIndex(938, (d_194_v119_).length(0))
                index45_ = default__.safeIndex(741, (d_106_v53_).length(0))
                rhs66_ = (d_37_v4_) + (d_37_v4_)
                rhs67_ = d_176_v109_
                rhs68_ = ((0) - (((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) + ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]))) * (len(d_39_v9_))
                lhs60_ = d_194_v119_
                lhs61_ = default__.safeIndex(938, (d_194_v119_).length(0))
                lhs62_ = d_106_v53_
                lhs63_ = default__.safeIndex(741, (d_106_v53_).length(0))
                lhs60_[lhs61_] = rhs66_
                d_176_v109_ = rhs67_
                lhs62_[lhs63_] = rhs68_
                d_196_v120_: C1
                nw28_ = C1()
                nw28_.ctor__((d_162_v97_.f19) | (_dafny.Map({(d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]: len(_dafny.SeqWithoutIsStrInference([d_33_v0_ for d_197_i14_ in range(default__.abs(916))]))})), (d_162_v97_).f20)
                d_196_v120_ = nw28_
                d_198_v121_: C1
                nw29_ = C1()
                nw29_.ctor__(d_196_v120_.f19, (d_162_v97_).f20)
                d_198_v121_ = nw29_
                d_199_v122_: C1
                nw30_ = C1()
                nw30_.ctor__(d_161_v96_, (d_162_v97_).f20)
                d_199_v122_ = nw30_
            elif True:
                d_200___mcc_h11_ = source5_.cf11
                d_201_cf11_ = d_200___mcc_h11_
                d_201_cf11_ = d_201_cf11_
                index46_ = default__.safeIndex(938, (d_99_v47_).length(0))
                (d_99_v47_)[index46_] = default__.fm1((d_37_v4_) + (d_37_v4_), (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))], (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], d_41_globalState_)
                d_202_v123_: _dafny.Map
                d_202_v123_ = _dafny.Map({(d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]: d_38_v8_})
                d_203_v124_: _dafny.Map
                d_203_v124_ = _dafny.Map({d_202_v123_: len(d_174_v107_)})
                d_204_v125_: bool
                out8_: bool
                out8_ = default__.m0(d_38_v8_, default__.fm1(d_37_v4_, (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))], len(d_201_cf11_), d_41_globalState_), (len(d_36_v3_)) not in (_dafny.Set({((d_203_v124_)[_dafny.Map({d_38_v8_: (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]})] if (_dafny.Map({d_38_v8_: (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]})) in (d_203_v124_) else d_35_v2_), (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]})), (d_39_v9_)[default__.safeIndex(d_35_v2_, len(d_39_v9_))], d_41_globalState_)
                d_204_v125_ = out8_
                (d_41_globalState_).f7 = (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]
            d_205_v126_: _dafny.Map
            d_205_v126_ = _dafny.Map({d_35_v2_: d_173_v106_})
            rhs69_ = (((d_205_v126_)[(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]] if ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) in (d_205_v126_) else _dafny.Set({d_38_v8_, (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))], d_38_v8_, d_38_v8_}))).intersection((d_173_v106_).intersection(d_173_v106_))
            rhs70_ = ((0) - (d_35_v2_)) - (d_35_v2_)
            d_173_v106_ = rhs69_
            d_35_v2_ = rhs70_
            d_206_v127_: D5
            d_206_v127_ = D5_DC15(d_106_v53_)
            d_206_v127_ = d_206_v127_
            index47_ = default__.safeIndex(352, (d_180_v113_).length(0))
            (d_180_v113_)[index47_] = d_106_v53_
        elif True:
            (d_41_globalState_).f7 = d_38_v8_
            d_207_v128_: _dafny.Array
            def lambda4_(d_208_v47_, d_209_v8_, d_210_v106_):
                def lambda5_(d_211_i15_):
                    return (_dafny.Map({D2_DC6(_dafny.Set({(d_208_v47_)[default__.safeIndex(938, (d_208_v47_).length(0))], d_209_v8_})): d_209_v8_})) | (_dafny.Map({D2_DC6(d_210_v106_): (d_208_v47_)[default__.safeIndex(938, (d_208_v47_).length(0))]}))

                return lambda5_

            init2_ = lambda4_(d_99_v47_, d_38_v8_, d_173_v106_)
            nw31_ = _dafny.Array(None, 19)
            for i0_2_ in range(nw31_.length(0)):
                nw31_[i0_2_] = init2_(i0_2_)
            d_207_v128_ = nw31_
            def lambda6_(d_212_v106_, d_213_v8_):
                def lambda7_(d_214_i16_):
                    def iife17_():
                        coll9_ = _dafny.Map()
                        compr_9_: D2
                        for compr_9_ in (_dafny.SeqWithoutIsStrInference([D2_DC6(d_212_v106_), D2_DC6(d_212_v106_), D2_DC6(d_212_v106_), D2_DC6(d_212_v106_)])).Elements:
                            d_215_v129_: D2 = compr_9_
                            if (d_215_v129_) in (_dafny.SeqWithoutIsStrInference([D2_DC6(d_212_v106_), D2_DC6(d_212_v106_), D2_DC6(d_212_v106_), D2_DC6(d_212_v106_)])):
                                coll9_[d_215_v129_] = d_213_v8_
                        return _dafny.Map(coll9_)
                    return iife17_()
                    

                return lambda7_

            init3_ = lambda6_(d_173_v106_, d_38_v8_)
            nw32_ = _dafny.Array(None, 14)
            for i0_3_ in range(nw32_.length(0)):
                nw32_[i0_3_] = init3_(i0_3_)
            d_207_v128_ = nw32_
            d_216_v130_: C0
            nw33_ = C0()
            nw33_.ctor__()
            d_216_v130_ = nw33_
            (d_41_globalState_).f16 = (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]
            d_217_v131_: D1
            d_217_v131_ = D1_DC4((d_170_v103_).cardinality, d_103_v51_, d_35_v2_, (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))])
            (d_41_globalState_).f7 = (d_217_v131_).cf8
        d_218_v132_: D3
        d_218_v132_ = D3_DC9(d_176_v109_)
        source6_ = d_218_v132_
        if source6_.is_DC10:
            d_219___mcc_h12_ = source6_.cf15
            d_220_cf15_ = d_219___mcc_h12_
            d_221_v133_: _dafny.Set
            d_221_v133_ = _dafny.Set({867})
            index48_ = default__.safeIndex(938, (d_99_v47_).length(0))
            index49_ = default__.safeIndex(938, (d_99_v47_).length(0))
            def iife18_():
                coll10_ = _dafny.Set()
                compr_10_: int
                for compr_10_ in _dafny.IntegerRange(669, 51):
                    d_222_v134_: int = compr_10_
                    if ((669) <= (d_222_v134_)) and ((d_222_v134_) < (51)):
                        coll10_ = coll10_.union(_dafny.Set([(d_222_v134_) + (d_220_cf15_)]))
                return _dafny.Set(coll10_)
            rhs71_ = d_38_v8_
            rhs72_ = ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]) * ((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))])
            rhs73_ = d_220_cf15_
            rhs74_ = (d_221_v133_) == ((iife18_()
            ).intersection(d_221_v133_))
            rhs75_ = (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))]
            lhs64_ = d_99_v47_
            lhs65_ = default__.safeIndex(938, (d_99_v47_).length(0))
            lhs66_ = d_99_v47_
            lhs67_ = default__.safeIndex(938, (d_99_v47_).length(0))
            lhs68_ = d_41_globalState_
            lhs64_[lhs65_] = rhs71_
            d_35_v2_ = rhs72_
            d_220_cf15_ = rhs73_
            lhs66_[lhs67_] = rhs74_
            lhs68_.f7 = rhs75_
            (d_41_globalState_).f6 = d_38_v8_
            d_223_v135_: _dafny.Map
            d_223_v135_ = _dafny.Map({d_162_v97_: d_99_v47_})
            d_224_v136_: _dafny.Array
            nw34_ = _dafny.Array(None, 2)
            nw34_[int(0)] = ((d_223_v135_)[d_162_v97_] if (d_162_v97_) in (d_223_v135_) else d_99_v47_)
            nw34_[int(1)] = d_99_v47_
            d_224_v136_ = nw34_
            index50_ = default__.safeIndex(471, (d_224_v136_).length(0))
            (d_224_v136_)[index50_] = d_99_v47_
            d_225_v137_: D2
            d_225_v137_ = D2_DC7(d_33_v0_)
            d_226_v138_: _dafny.MultiSet
            d_226_v138_ = _dafny.MultiSet([d_225_v137_])
            d_227_v139_: _dafny.Set
            d_227_v139_ = _dafny.Set({d_162_v97_})
            d_228_v140_: _dafny.Map
            d_228_v140_ = _dafny.Map({d_38_v8_: d_227_v139_})
            d_229_v141_: D6
            d_229_v141_ = D6_DC19(d_227_v139_)
            d_230_v142_: _dafny.Array
            nw35_ = _dafny.Array(None, 25)
            nw35_[int(0)] = d_227_v139_
            nw35_[int(1)] = d_227_v139_
            nw35_[int(2)] = ((d_228_v140_)[True] if (True) in (d_228_v140_) else _dafny.Set({d_162_v97_}))
            nw35_[int(3)] = (d_227_v139_) | (d_227_v139_)
            nw35_[int(4)] = d_227_v139_
            nw35_[int(5)] = (d_227_v139_) | (d_227_v139_)
            nw35_[int(6)] = d_227_v139_
            nw35_[int(7)] = d_227_v139_
            nw35_[int(8)] = (_dafny.Set({d_162_v97_, d_162_v97_})) | (d_227_v139_)
            nw35_[int(9)] = d_227_v139_
            nw35_[int(10)] = ((d_229_v141_).cf30).intersection(_dafny.Set({d_162_v97_, d_162_v97_, d_162_v97_}))
            nw35_[int(11)] = (d_227_v139_).intersection(d_227_v139_)
            nw35_[int(12)] = d_227_v139_
            nw35_[int(13)] = d_227_v139_
            nw35_[int(14)] = d_227_v139_
            nw35_[int(15)] = d_227_v139_
            nw35_[int(16)] = _dafny.Set({d_162_v97_, d_162_v97_, d_162_v97_})
            nw35_[int(17)] = d_227_v139_
            nw35_[int(18)] = _dafny.Set({d_162_v97_})
            nw35_[int(19)] = (d_227_v139_) | (d_227_v139_)
            nw35_[int(20)] = d_227_v139_
            nw35_[int(21)] = (d_227_v139_).intersection(d_227_v139_)
            nw35_[int(22)] = _dafny.Set({d_162_v97_, d_162_v97_, d_162_v97_, d_162_v97_, (d_163_v98_)[default__.safeIndex((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], len(d_163_v98_))]})
            nw35_[int(23)] = (d_227_v139_) - (d_227_v139_)
            nw35_[int(24)] = d_227_v139_
            d_230_v142_ = nw35_
            index51_ = default__.safeIndex(604, (d_230_v142_).length(0))
            (d_230_v142_)[index51_] = _dafny.Set({d_162_v97_})
            d_231_v143_: _dafny.Seq
            d_231_v143_ = _dafny.SeqWithoutIsStrInference([d_225_v137_, d_225_v137_])
            pat_let_tv2_ = d_33_v0_
            index52_ = default__.safeIndex(471, (d_224_v136_).length(0))
            index53_ = default__.safeIndex(604, (d_230_v142_).length(0))
            def iife19_(_pat_let4_0):
                def iife20_(d_232_dt__update__tmp_h2_):
                    def iife21_(_pat_let5_0):
                        def iife22_(d_233_dt__update_hcf12_h1_):
                            return D2_DC7(d_233_dt__update_hcf12_h1_)
                        return iife22_(_pat_let5_0)
                    return iife21_(pat_let_tv2_)
                return iife20_(_pat_let4_0)
            rhs76_ = (d_172_v105_)[default__.safeIndex(d_35_v2_, len(d_172_v105_))]
            rhs77_ = ((D7_DC21(_dafny.MultiSet(d_231_v143_))).cf34 if default__.fm1(d_37_v4_, d_38_v8_, d_35_v2_, d_41_globalState_) else _dafny.MultiSet((d_231_v143_).set(default__.safeIndex(d_220_cf15_, len(d_231_v143_)), iife19_(d_225_v137_))))
            rhs78_ = d_38_v8_
            rhs79_ = d_38_v8_
            rhs80_ = _dafny.Set({d_162_v97_, d_162_v97_})
            lhs69_ = d_224_v136_
            lhs70_ = default__.safeIndex(471, (d_224_v136_).length(0))
            lhs71_ = d_41_globalState_
            lhs72_ = d_41_globalState_
            lhs73_ = d_230_v142_
            lhs74_ = default__.safeIndex(604, (d_230_v142_).length(0))
            lhs69_[lhs70_] = rhs76_
            d_226_v138_ = rhs77_
            lhs71_.f6 = rhs78_
            lhs72_.f7 = rhs79_
            lhs73_[lhs74_] = rhs80_
            d_220_cf15_ = d_220_cf15_
        elif True:
            d_234___mcc_h13_ = source6_.cf14
            d_235_cf14_ = d_234___mcc_h13_
            d_236_v144_: _dafny.Set
            d_236_v144_ = _dafny.Set({(d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]})
            d_237_v145_: _dafny.Array
            d_238_v146_: _dafny.Map
            out9_: _dafny.Array
            out10_: _dafny.Map
            out9_, out10_ = (d_162_v97_).m1((d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))], d_236_v144_, len(d_174_v107_), d_41_globalState_)
            d_237_v145_ = out9_
            d_238_v146_ = out10_
            (d_41_globalState_).f0 = (d_106_v53_)[default__.safeIndex(741, (d_106_v53_).length(0))]
            d_239_v147_: bool
            out11_: bool
            out11_ = default__.m0(False, (d_99_v47_)[default__.safeIndex(938, (d_99_v47_).length(0))], d_38_v8_, d_38_v8_, d_41_globalState_)
            d_239_v147_ = out11_
            (d_41_globalState_).f0 = d_35_v2_
        _dafny.print(_dafny.string_of(d_33_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_34_v1_) == (_dafny.MultiSet([_dafny.CodePoint('n'), _dafny.CodePoint('n')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_35_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_36_v3_) == (_dafny.SeqWithoutIsStrInference([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_37_v4_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_38_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_39_v9_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_40_v10_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_41_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_41_globalState_).f1) == (_dafny.SeqWithoutIsStrInference([2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_41_globalState_.f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_41_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_41_globalState_.f11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_41_globalState_).f12) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_41_globalState_).f14) == (_dafny.Map({_dafny.Map({}): 863}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_41_globalState_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_41_globalState_).f17) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_41_globalState_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_47_v11_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_47_v11_).cf1) == (_dafny.SeqWithoutIsStrInference([False, False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_99_v47_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v48_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_100_v48_).cf1) == (_dafny.SeqWithoutIsStrInference([True, True, False, False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_101_v49_).cf2).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_101_v49_).cf2).cf1) == (_dafny.SeqWithoutIsStrInference([True, True, False, False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_102_v50_).cf2).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_102_v50_).cf2).cf1) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v51_) == (_dafny.Map({_dafny.CodePoint('n'): 5}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v52_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_v53_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_107_v54_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_108_v55_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v96_) == (_dafny.Map({False: 676}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v97_.f19) == (_dafny.Map({False: 676}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_162_v97_).f20).cf2).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_162_v97_).f20).cf2).cf1) == (_dafny.SeqWithoutIsStrInference([True, True, False, False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_163_v98_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_164_v99_).cf16)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(((d_165_v100_).cf21).cf16)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_166_v101_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_169_v102_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v103_) == (_dafny.MultiSet([896]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v104_) == (_dafny.Map({False: _dafny.MultiSet([896])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_172_v105_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v106_) == (_dafny.Set({True, False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v107_) == (_dafny.Map({484: -1362}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_175_v108_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_177_v110_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_178_v111_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf24) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf25)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf26.f19) == (_dafny.Map({False: 676}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((((d_179_v112_).cf26).f20).cf2).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((((d_179_v112_).cf26).f20).cf2).cf1) == (_dafny.SeqWithoutIsStrInference([True, True, False, False, True, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_179_v112_).cf27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_179_v112_).cf28) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x'), _dafny.CodePoint('x')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[1])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[2])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[3])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[4])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[5])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[6])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[7])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[8])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[9])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[10])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[11])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[12])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[13])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[14])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[15])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_180_v113_)[16])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_181_v114_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC12()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)

class D4_DC12(D4, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf17', Any), ('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)

class D5_DC16(D5, NamedTuple('DC16', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC17(D5, NamedTuple('DC17', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {self.cf28.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC20(False, D5.default()(), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC20(D6, NamedTuple('DC20', [('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC22(False, int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D7_DC23)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC22(D7, NamedTuple('DC22', [('cf35', Any), ('cf36', Any), ('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC23(D7, NamedTuple('DC23', [('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC23({_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC23) and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC21(D7, NamedTuple('DC21', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC25(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D8_DC26)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)

class D8_DC25(D8, NamedTuple('DC25', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC26(D8, NamedTuple('DC26', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC26({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC26) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC24(D8, NamedTuple('DC24', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f6: bool = False
        self.f7: bool = False
        self.f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f16: int = int(0)
        self._f1: _dafny.Seq = _dafny.Seq({})
        self._f2: int = int(0)
        self._f3: bool = False
        self._f4: int = int(0)
        self._f5: bool = False
        self._f8: int = int(0)
        self._f9: bool = False
        self._f10: bool = False
        self._f12: _dafny.Map = _dafny.Map({})
        self._f13: bool = False
        self._f14: _dafny.Map = _dafny.Map({})
        self._f15: bool = False
        self._f17: _dafny.Seq = _dafny.Seq({})
        self._f18: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18

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
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14
    @property
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm7(self, p0, p1, globalState):
        if not (False) or (not(True)):
            return (_dafny.MultiSet([len(_dafny.Map({538: D0_DC0(len(_dafny.SeqWithoutIsStrInference([-515])), _dafny.SeqWithoutIsStrInference([True, True]))}))])).intersection(_dafny.MultiSet([490, len(_dafny.Map({-631: False}))]))
        elif True:
            return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([666]))

    def fm8(self, p0, globalState):
        source7_ = D0_DC0(856, _dafny.SeqWithoutIsStrInference([not(True)]))
        if source7_.is_DC0:
            d_240___mcc_h0_ = source7_.cf0
            d_241___mcc_h1_ = source7_.cf1
            d_242_cf1_ = d_241___mcc_h1_
            d_243_cf0_ = d_240___mcc_h0_
            return (_dafny.Map({not(not(False)): False})) | (_dafny.Map({True: not(False)}))
        elif True:
            d_244___mcc_h2_ = source7_.cf2
            d_245_cf2_ = d_244___mcc_h2_
            if False:
                return _dafny.Map({True: True})
            elif True:
                return _dafny.Map({False: True})


class C1:
    def  __init__(self):
        self.f19: _dafny.Map = _dafny.Map({})
        self._f20: D0 = D0.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f19, f20):
        (self).f19 = f19
        (self)._f20 = f20

    def fm6(self, p0, p1, p2, p3, globalState):
        return (0) - (((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vc")))) + (len(_dafny.Map({True: True})))) - (471))

    def m1(self, p0, p1, p2, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Map = _dafny.Map({})
        d_246_v0_: _dafny.Seq
        d_246_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iuyubyqfn"))
        d_247_v1_: str
        d_247_v1_ = _dafny.CodePoint('y')
        hi1_ = p0
        for d_248_i0_ in range((len((d_246_v0_).set(default__.safeIndex(p2, len(d_246_v0_)), d_247_v1_))) * (p0), hi1_):
            (globalState).f0 = d_248_i0_
            d_249_v2_: bool
            d_249_v2_ = False
            d_250_v3_: _dafny.Seq
            d_250_v3_ = _dafny.SeqWithoutIsStrInference([d_249_v2_])
            (globalState).f7 = not ((d_250_v3_)[default__.safeIndex(p2, len(d_250_v3_))]) or (d_249_v2_)
            (globalState).f6 = (p0) in (_dafny.Set({p0}))
            d_251_v4_: C0
            nw36_ = C0()
            nw36_.ctor__()
            d_251_v4_ = nw36_
        d_252_v5_: _dafny.Array
        nw37_ = _dafny.Array(False, 29)
        d_252_v5_ = nw37_
        d_253_v6_: bool
        d_253_v6_ = True
        index54_ = default__.safeIndex(35, (d_252_v5_).length(0))
        (d_252_v5_)[index54_] = d_253_v6_
        d_254_v7_: _dafny.Map
        d_254_v7_ = _dafny.Map({d_247_v1_: p0})
        d_255_v8_: D1
        d_255_v8_ = D1_DC4(p2, d_254_v7_, p2, d_253_v6_)
        d_256_v9_: _dafny.Seq
        d_256_v9_ = _dafny.SeqWithoutIsStrInference([d_253_v6_, d_253_v6_, d_253_v6_, (d_255_v8_).cf8])
        index55_ = default__.safeIndex(35, (d_252_v5_).length(0))
        (d_252_v5_)[index55_] = not ((d_256_v9_)[default__.safeIndex(p2, len(d_256_v9_))]) or (d_253_v6_)
        d_257_v10_: _dafny.Array
        def lambda8_(d_258_p0_, d_259_v5_):
            def lambda9_(d_260_i2_):
                return (d_260_i2_) + (len(_dafny.Map({d_258_p0_: not((d_259_v5_)[default__.safeIndex(35, (d_259_v5_).length(0))])})))

            return lambda9_

        init4_ = lambda8_(p0, d_252_v5_)
        nw38_ = _dafny.Array(None, 6)
        for i0_4_ in range(nw38_.length(0)):
            nw38_[i0_4_] = init4_(i0_4_)
        d_257_v10_ = nw38_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_257_v10_).length(0)):
            d_261_i1_: int = guard_loop_2_
            if (True) and (((0) <= (d_261_i1_)) and ((d_261_i1_) < ((d_257_v10_).length(0)))):
                (d_257_v10_)[(d_261_i1_)] = (d_261_i1_) - (len(d_246_v0_))
        d_262_v11_: _dafny.Array
        nw39_ = _dafny.Array(None, 4)
        nw39_[int(0)] = d_257_v10_
        nw39_[int(1)] = d_257_v10_
        nw39_[int(2)] = d_257_v10_
        nw39_[int(3)] = d_257_v10_
        d_262_v11_ = nw39_
        d_263_v12_: _dafny.Map
        d_263_v12_ = _dafny.Map({(d_252_v5_)[default__.safeIndex(35, (d_252_v5_).length(0))]: d_262_v11_})
        d_263_v12_ = (d_263_v12_).set(d_253_v6_, d_262_v11_)
        d_264_v13_: D1
        d_264_v13_ = D1_DC5(d_252_v5_, p2)
        d_265_v14_: D2
        d_265_v14_ = D2_DC6(_dafny.Set({d_253_v6_}))
        d_266_v15_: D1
        d_266_v15_ = D1_DC5((d_264_v13_).cf9, len((d_265_v14_).cf11))
        d_267_v16_: _dafny.Seq
        d_267_v16_ = _dafny.SeqWithoutIsStrInference([d_252_v5_, d_252_v5_, d_252_v5_])
        d_268_v17_: _dafny.Array
        nw40_ = _dafny.Array(None, 28)
        nw40_[int(0)] = d_252_v5_
        nw40_[int(1)] = d_252_v5_
        nw40_[int(2)] = d_252_v5_
        nw40_[int(3)] = d_252_v5_
        nw40_[int(4)] = d_252_v5_
        nw40_[int(5)] = d_252_v5_
        nw40_[int(6)] = (d_266_v15_).cf9
        nw40_[int(7)] = d_252_v5_
        nw40_[int(8)] = d_252_v5_
        nw40_[int(9)] = d_252_v5_
        nw40_[int(10)] = d_252_v5_
        nw40_[int(11)] = d_252_v5_
        nw40_[int(12)] = d_252_v5_
        nw40_[int(13)] = d_252_v5_
        nw40_[int(14)] = d_252_v5_
        nw40_[int(15)] = d_252_v5_
        nw40_[int(16)] = d_252_v5_
        nw40_[int(17)] = d_252_v5_
        nw40_[int(18)] = d_252_v5_
        nw40_[int(19)] = (d_267_v16_)[default__.safeIndex(784, len(d_267_v16_))]
        nw40_[int(20)] = d_252_v5_
        nw40_[int(21)] = d_252_v5_
        nw40_[int(22)] = d_252_v5_
        nw40_[int(23)] = d_252_v5_
        nw40_[int(24)] = d_252_v5_
        nw40_[int(25)] = d_252_v5_
        nw40_[int(26)] = d_252_v5_
        nw40_[int(27)] = d_252_v5_
        d_268_v17_ = nw40_
        index56_ = default__.safeIndex(920, (d_268_v17_).length(0))
        (d_268_v17_)[index56_] = d_252_v5_
        index57_ = default__.safeIndex(920, (d_268_v17_).length(0))
        (d_268_v17_)[index57_] = d_252_v5_
        d_269_v18_: _dafny.Seq
        d_269_v18_ = _dafny.SeqWithoutIsStrInference([d_246_v0_, d_246_v0_])
        d_270_v19_: C0
        nw41_ = C0()
        nw41_.ctor__()
        d_270_v19_ = nw41_
        d_271_v20_: _dafny.Seq
        d_271_v20_ = _dafny.SeqWithoutIsStrInference([p2])
        d_272_v21_: _dafny.MultiSet
        d_272_v21_ = _dafny.MultiSet([p2, len(d_271_v20_)])
        d_273_v22_: _dafny.Map
        d_273_v22_ = _dafny.Map({(d_272_v21_).cardinality: d_270_v19_})
        d_274_v23_: _dafny.Seq
        d_274_v23_ = _dafny.SeqWithoutIsStrInference([d_270_v19_, d_270_v19_])
        d_275_v24_: _dafny.Array
        nw42_ = _dafny.Array(None, 28)
        nw42_[int(0)] = d_270_v19_
        nw42_[int(1)] = ((d_273_v22_)[p2] if (p2) in (d_273_v22_) else d_270_v19_)
        nw42_[int(2)] = d_270_v19_
        nw42_[int(3)] = d_270_v19_
        nw42_[int(4)] = d_270_v19_
        nw42_[int(5)] = d_270_v19_
        nw42_[int(6)] = d_270_v19_
        nw42_[int(7)] = d_270_v19_
        nw42_[int(8)] = (D3_DC9(d_270_v19_)).cf14
        nw42_[int(9)] = d_270_v19_
        nw42_[int(10)] = d_270_v19_
        nw42_[int(11)] = (d_274_v23_)[default__.safeIndex(p0, len(d_274_v23_))]
        nw42_[int(12)] = d_270_v19_
        nw42_[int(13)] = d_270_v19_
        nw42_[int(14)] = d_270_v19_
        nw42_[int(15)] = d_270_v19_
        nw42_[int(16)] = d_270_v19_
        nw42_[int(17)] = d_270_v19_
        nw42_[int(18)] = d_270_v19_
        nw42_[int(19)] = d_270_v19_
        nw42_[int(20)] = d_270_v19_
        nw42_[int(21)] = d_270_v19_
        nw42_[int(22)] = d_270_v19_
        nw42_[int(23)] = d_270_v19_
        nw42_[int(24)] = d_270_v19_
        nw42_[int(25)] = d_270_v19_
        nw42_[int(26)] = d_270_v19_
        nw42_[int(27)] = d_270_v19_
        d_275_v24_ = nw42_
        d_276_v25_: _dafny.Array
        nw43_ = _dafny.Array(None, 3)
        nw43_[int(0)] = d_275_v24_
        nw43_[int(1)] = d_275_v24_
        nw43_[int(2)] = d_275_v24_
        d_276_v25_ = nw43_
        index58_ = default__.safeIndex(196, (d_276_v25_).length(0))
        (d_276_v25_)[index58_] = d_275_v24_
        index59_ = default__.safeIndex(196, (d_276_v25_).length(0))
        rhs81_ = d_269_v18_
        rhs82_ = p0
        rhs83_ = d_275_v24_
        rhs84_ = (d_252_v5_)[default__.safeIndex(35, (d_252_v5_).length(0))]
        lhs75_ = globalState
        lhs76_ = d_276_v25_
        lhs77_ = default__.safeIndex(196, (d_276_v25_).length(0))
        d_269_v18_ = rhs81_
        lhs75_.f0 = rhs82_
        lhs76_[lhs77_] = rhs83_
        d_253_v6_ = rhs84_
        def lambda10_(d_277_v6_):
            def lambda11_(d_278_i3_):
                return d_277_v6_

            return lambda11_

        init5_ = lambda10_(d_253_v6_)
        nw44_ = _dafny.Array(None, 17)
        for i0_5_ in range(nw44_.length(0)):
            nw44_[i0_5_] = init5_(i0_5_)
        r0 = nw44_
        d_279_v26_: _dafny.Seq
        d_279_v26_ = _dafny.SeqWithoutIsStrInference([d_271_v20_, d_271_v20_, d_271_v20_, d_271_v20_, d_271_v20_])
        d_280_v27_: _dafny.Map
        d_280_v27_ = _dafny.Map({(d_279_v26_)[default__.safeIndex(p2, len(d_279_v26_))]: d_253_v6_})
        d_281_v28_: _dafny.Map
        d_281_v28_ = _dafny.Map({(d_252_v5_)[default__.safeIndex(35, (d_252_v5_).length(0))]: d_271_v20_})
        d_282_v29_: _dafny.Map
        d_282_v29_ = _dafny.Map({-977: p2})
        r1 = ((d_280_v27_).set(((d_281_v28_)[(d_252_v5_)[default__.safeIndex(35, (d_252_v5_).length(0))]] if ((d_252_v5_)[default__.safeIndex(35, (d_252_v5_).length(0))]) in (d_281_v28_) else d_271_v20_), (d_252_v5_)[default__.safeIndex(35, (d_252_v5_).length(0))])).set((_dafny.SeqWithoutIsStrInference([p2 for d_283_i4_ in range(default__.abs(722))])).set(default__.safeIndex((0) - (p2), len(_dafny.SeqWithoutIsStrInference([p2 for d_284_i4_ in range(default__.abs(722))]))), len(d_282_v29_)), d_253_v6_)
        return r0, r1

    @property
    def f20(self):
        return self._f20
