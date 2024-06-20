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
    def fm0(p0, globalState):
        if True:
            return 772
        elif True:
            return (914) * (333)

    @staticmethod
    def fm1(p0, globalState):
        return ((D0_DC0(len(_dafny.Set({_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([137 for d_0_i0_ in range(default__.abs(918))]))}), _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))})})), _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tq")): -356}), _dafny.MultiSet([-192, len(_dafny.Map({True: True}))]))).cf0) <= (((0) - ((0) - ((0) - (-610)))) - (len(_dafny.Map({True: True}))))

    @staticmethod
    def fm2(p0, globalState):
        return _dafny.MultiSet([392])

    @staticmethod
    def fm3(globalState):
        return D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))), (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnjhhmdfk")): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tslm")))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lhmcdhe")): 798})), (_dafny.MultiSet([(0) - ((_dafny.MultiSet([False])).cardinality), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoksdycy")))])) | (_dafny.MultiSet([298])))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        return (_dafny.Set({not(not(False)), False, False, True})) - ((_dafny.Set({False, True, False})) - (_dafny.Set({not(True), False})))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1_i0_ in range(default__.abs(780))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('o')]))

    @staticmethod
    def fm6(p0, globalState):
        return _dafny.CodePoint('c')

    @staticmethod
    def fm7(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([True, False])) - (_dafny.MultiSet([False, True, False]))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False, True, True, not(True)]))).intersection(_dafny.MultiSet([True])))

    @staticmethod
    def fm8(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({204, 26}))) for d_2_i0_ in range(default__.abs(649))])).Elements:
                d_3_v0_: int = compr_0_
                if (d_3_v0_) in (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({204, 26}))) for d_2_i0_ in range(default__.abs(649))])):
                    coll0_[default__.safeDivisionInt(d_3_v0_, 308)] = True
            return _dafny.Map(coll0_)
        def iife1_():
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))})).keys.Elements:
                    d_4_v2_: int = compr_3_
                    if (d_4_v2_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))})):
                        coll3_[(d_4_v2_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qc"))))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrjysrs")))
                return _dafny.Map(coll3_)
            coll1_ = _dafny.Map()
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))})).keys.Elements:
                    d_4_v2_: int = compr_2_
                    if (d_4_v2_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))): len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))})):
                        coll2_[(d_4_v2_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qc"))))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrjysrs")))
                return _dafny.Map(coll2_)
            compr_1_: int
            for compr_1_ in (iife2_()
            ).keys.Elements:
                d_5_v1_: int = compr_1_
                if (d_5_v1_) in (iife3_()
                ):
                    coll1_[(d_5_v1_) - (432)] = True
            return _dafny.Map(coll1_)
        return ((D4_DC11(iife0_()
) if False else D4_DC11(iife1_()
))).cf19

    @staticmethod
    def m0(p0, p1, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.Set = _dafny.Set({})
        r3: bool = False
        d_6_v0_: _dafny.Seq
        d_6_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hkx"))
        d_7_v1_: _dafny.Seq
        d_7_v1_ = _dafny.SeqWithoutIsStrInference([d_6_v0_])
        d_8_v2_: bool
        d_8_v2_ = True
        d_9_v3_: _dafny.Map
        d_9_v3_ = _dafny.Map({d_8_v2_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_10_i0_ in range(default__.abs(902))]))})
        d_11_v4_: _dafny.Map
        d_11_v4_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p1])): d_8_v2_})
        d_12_v5_: _dafny.Seq
        d_12_v5_ = _dafny.SeqWithoutIsStrInference([d_8_v2_, d_8_v2_])
        d_13_v6_: _dafny.Map
        d_13_v6_ = _dafny.Map({d_6_v0_: p1})
        d_14_v7_: str
        d_14_v7_ = _dafny.CodePoint('i')
        d_15_v8_: _dafny.Map
        d_15_v8_ = _dafny.Map({d_14_v7_: True})
        d_16_v9_: _dafny.MultiSet
        d_16_v9_ = _dafny.MultiSet([len(d_15_v8_), -545, p1, p1])
        d_17_v10_: D0
        d_17_v10_ = D0_DC0(p1, d_13_v6_, d_16_v9_)
        d_18_v11_: _dafny.Map
        d_18_v11_ = _dafny.Map({default__.fm0(d_17_v10_, globalState): d_6_v0_})
        d_19_v12_: _dafny.Map
        d_19_v12_ = _dafny.Map({((d_11_v4_)[len(d_12_v5_)] if (len(d_12_v5_)) in (d_11_v4_) else d_8_v2_): ((d_18_v11_)[default__.fm0(d_17_v10_, globalState)] if (default__.fm0(d_17_v10_, globalState)) in (d_18_v11_) else d_6_v0_)})
        d_20_v13_: D1
        d_20_v13_ = D1_DC4(d_6_v0_, d_8_v2_, d_7_v1_, d_6_v0_, d_14_v7_)
        d_21_v14_: _dafny.Array
        nw0_ = _dafny.Array(None, 28)
        nw0_[int(0)] = ((d_7_v1_)[default__.safeIndex(p1, len(d_7_v1_))]) + (d_6_v0_)
        nw0_[int(1)] = d_6_v0_
        nw0_[int(2)] = default__.fm5(len(d_6_v0_), p1, p1, len(d_9_v3_), globalState)
        nw0_[int(3)] = d_6_v0_
        nw0_[int(4)] = d_6_v0_
        nw0_[int(5)] = (d_6_v0_) + (d_6_v0_)
        nw0_[int(6)] = _dafny.SeqWithoutIsStrInference([(d_6_v0_)[default__.safeIndex(629, len(d_6_v0_))] for d_22_i1_ in range(default__.abs(713))])
        nw0_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrekd"))
        nw0_[int(8)] = d_6_v0_
        nw0_[int(9)] = d_6_v0_
        nw0_[int(10)] = d_6_v0_
        nw0_[int(11)] = d_6_v0_
        nw0_[int(12)] = ((d_19_v12_)[d_8_v2_] if (d_8_v2_) in (d_19_v12_) else (d_6_v0_).set(default__.safeIndex(385, len(d_6_v0_)), (d_20_v13_).cf10))
        nw0_[int(13)] = d_6_v0_
        nw0_[int(14)] = d_6_v0_
        nw0_[int(15)] = (d_6_v0_) + (d_6_v0_)
        nw0_[int(16)] = d_6_v0_
        nw0_[int(17)] = d_6_v0_
        nw0_[int(18)] = d_6_v0_
        nw0_[int(19)] = d_6_v0_
        nw0_[int(20)] = d_6_v0_
        nw0_[int(21)] = d_6_v0_
        nw0_[int(22)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxsg"))
        nw0_[int(23)] = (d_6_v0_) + (d_6_v0_)
        nw0_[int(24)] = d_6_v0_
        nw0_[int(25)] = d_6_v0_
        nw0_[int(26)] = d_6_v0_
        nw0_[int(27)] = d_6_v0_
        d_21_v14_ = nw0_
        d_21_v14_ = d_21_v14_
        d_23_v15_: C0
        nw1_ = C0()
        nw1_.ctor__((p1) + (p1))
        d_23_v15_ = nw1_
        d_24_v16_: _dafny.Set
        d_24_v16_ = _dafny.Set({d_8_v2_, d_8_v2_})
        d_25_v17_: D1
        d_25_v17_ = D1_DC3(d_24_v16_)
        pat_let_tv0_ = d_6_v0_
        pat_let_tv1_ = d_6_v0_
        pat_let_tv2_ = d_8_v2_
        pat_let_tv3_ = d_8_v2_
        pat_let_tv4_ = d_8_v2_
        def lambda0_(source0_):
            if source0_.is_DC3:
                d_26___mcc_h0_ = source0_.cf5
                d_27_cf5_ = d_26___mcc_h0_
                return (pat_let_tv0_) < (pat_let_tv1_)
            elif source0_.is_DC4:
                d_28___mcc_h1_ = source0_.cf6
                d_29___mcc_h2_ = source0_.cf7
                d_30___mcc_h3_ = source0_.cf8
                d_31___mcc_h4_ = source0_.cf9
                d_32___mcc_h5_ = source0_.cf10
                d_33_cf10_ = d_32___mcc_h5_
                d_34_cf9_ = d_31___mcc_h4_
                d_35_cf8_ = d_30___mcc_h3_
                d_36_cf7_ = d_29___mcc_h2_
                d_37_cf6_ = d_28___mcc_h1_
                return d_36_cf7_
            elif source0_.is_DC2:
                d_38___mcc_h6_ = source0_.cf4
                d_39_cf4_ = d_38___mcc_h6_
                return pat_let_tv2_
            elif True:
                d_40___mcc_h7_ = source0_.cf11
                d_41_cf11_ = d_40___mcc_h7_
                return (pat_let_tv3_) or (pat_let_tv4_)

        r3 = lambda0_(d_25_v17_)
        d_42_v18_: _dafny.Array
        def lambda1_(d_43_v15_):
            def lambda2_(d_44_i2_):
                return (d_44_i2_) * ((d_43_v15_).f18)

            return lambda2_

        init0_ = lambda1_(d_23_v15_)
        nw2_ = _dafny.Array(None, 19)
        for i0_0_ in range(nw2_.length(0)):
            nw2_[i0_0_] = init0_(i0_0_)
        d_42_v18_ = nw2_
        def lambda3_(d_45_i3_):
            return (d_45_i3_) * (-986)

        init1_ = lambda3_
        nw3_ = _dafny.Array(None, 17)
        for i0_1_ in range(nw3_.length(0)):
            nw3_[i0_1_] = init1_(i0_1_)
        d_42_v18_ = nw3_
        hi0_ = (d_23_v15_).f18
        for d_46_i4_ in range((d_23_v15_).f18, hi0_):
            index0_ = default__.safeIndex(580, (d_42_v18_).length(0))
            (d_42_v18_)[index0_] = (d_23_v15_).f18
            index1_ = default__.safeIndex(580, (d_42_v18_).length(0))
            (d_42_v18_)[index1_] = ((d_23_v15_).f18 if d_8_v2_ else (p1) * (p1))
            d_6_v0_ = (d_6_v0_) + (d_6_v0_)
            d_47_v19_: _dafny.Seq
            d_47_v19_ = _dafny.SeqWithoutIsStrInference([95, p1])
            d_48_v20_: _dafny.Array
            nw4_ = _dafny.Array(None, 21)
            nw4_[int(0)] = d_8_v2_
            nw4_[int(1)] = d_8_v2_
            nw4_[int(2)] = d_8_v2_
            nw4_[int(3)] = d_8_v2_
            nw4_[int(4)] = d_8_v2_
            nw4_[int(5)] = d_8_v2_
            nw4_[int(6)] = d_8_v2_
            nw4_[int(7)] = d_8_v2_
            nw4_[int(8)] = d_8_v2_
            nw4_[int(9)] = d_8_v2_
            nw4_[int(10)] = d_8_v2_
            nw4_[int(11)] = True
            nw4_[int(12)] = d_8_v2_
            nw4_[int(13)] = d_8_v2_
            nw4_[int(14)] = d_8_v2_
            nw4_[int(15)] = d_8_v2_
            nw4_[int(16)] = d_8_v2_
            nw4_[int(17)] = d_8_v2_
            nw4_[int(18)] = d_8_v2_
            nw4_[int(19)] = d_8_v2_
            nw4_[int(20)] = d_8_v2_
            d_48_v20_ = nw4_
            d_49_v21_: _dafny.Seq
            d_49_v21_ = _dafny.SeqWithoutIsStrInference([d_48_v20_, d_48_v20_, d_48_v20_, d_48_v20_])
            d_12_v5_ = ((d_12_v5_) + ((d_12_v5_) + (d_12_v5_))).set(default__.safeIndex((d_17_v10_).cf0, len((d_12_v5_) + ((d_12_v5_) + (d_12_v5_)))), (len(d_47_v19_)) < ((0) - (len(d_49_v21_))))
            source1_ = d_20_v13_
            if source1_.is_DC3:
                d_50___mcc_h8_ = source1_.cf5
                d_51_cf5_ = d_50___mcc_h8_
                (globalState).f13 = (721) > (d_46_i4_)
                (globalState).f13 = d_8_v2_
                (globalState).f10 = (d_23_v15_).f18
                r3 = d_8_v2_
            elif source1_.is_DC4:
                d_52___mcc_h9_ = source1_.cf6
                d_53___mcc_h10_ = source1_.cf7
                d_54___mcc_h11_ = source1_.cf8
                d_55___mcc_h12_ = source1_.cf9
                d_56___mcc_h13_ = source1_.cf10
                d_57_cf10_ = d_56___mcc_h13_
                d_58_cf9_ = d_55___mcc_h12_
                d_59_cf8_ = d_54___mcc_h11_
                d_60_cf7_ = d_53___mcc_h10_
                d_61_cf6_ = d_52___mcc_h9_
                d_62_v22_: _dafny.Seq
                d_62_v22_ = _dafny.SeqWithoutIsStrInference([d_23_v15_, d_23_v15_, d_23_v15_])
                d_62_v22_ = (d_62_v22_) + (d_62_v22_)
                d_63_v23_: _dafny.Map
                d_63_v23_ = _dafny.Map({d_47_v19_: (d_23_v15_).f18})
                index2_ = default__.safeIndex(580, (d_42_v18_).length(0))
                rhs0_ = default__.safeDivisionInt((d_23_v15_).f18, (0) - (default__.fm0(d_17_v10_, globalState)))
                rhs1_ = d_63_v23_
                rhs2_ = len(((_dafny.SeqWithoutIsStrInference([len(_dafny.Set({((d_16_v9_)[(d_23_v15_).f18] if ((d_23_v15_).f18) in (d_16_v9_) else len(d_6_v0_))})) for d_64_i5_ in range(default__.abs(-273))])) + (d_47_v19_)) + ((d_47_v19_) + (d_47_v19_)))
                lhs0_ = d_42_v18_
                lhs1_ = default__.safeIndex(580, (d_42_v18_).length(0))
                lhs2_ = globalState
                lhs0_[lhs1_] = rhs0_
                d_63_v23_ = rhs1_
                lhs2_.f4 = rhs2_
                index3_ = default__.safeIndex(580, (d_42_v18_).length(0))
                (d_42_v18_)[index3_] = d_46_i4_
                d_65_v24_: _dafny.Array
                def lambda4_(d_66_v0_):
                    def lambda5_(d_67_i6_):
                        return (d_66_v0_)[default__.safeIndex(-327, len(d_66_v0_))]

                    return lambda5_

                init2_ = lambda4_(d_6_v0_)
                nw5_ = _dafny.Array(None, 25)
                for i0_2_ in range(nw5_.length(0)):
                    nw5_[i0_2_] = init2_(i0_2_)
                d_65_v24_ = nw5_
                d_68_v25_: _dafny.MultiSet
                d_68_v25_ = _dafny.MultiSet([d_65_v24_])
                d_68_v25_ = d_68_v25_
            elif source1_.is_DC2:
                d_69___mcc_h14_ = source1_.cf4
                d_70_cf4_ = d_69___mcc_h14_
                d_71_v26_: C0
                nw6_ = C0()
                nw6_.ctor__((d_46_i4_) - (default__.fm0(d_17_v10_, globalState)))
                d_71_v26_ = nw6_
                r3 = ((d_23_v15_).f18) > (p1)
                d_72_v27_: _dafny.Map
                d_72_v27_ = _dafny.Map({d_71_v26_: d_8_v2_})
                d_73_v28_: _dafny.Array
                nw7_ = _dafny.Array(None, 11)
                nw7_[int(0)] = d_72_v27_
                nw7_[int(1)] = d_72_v27_
                nw7_[int(2)] = _dafny.Map({d_71_v26_: not(d_8_v2_)})
                nw7_[int(3)] = d_72_v27_
                nw7_[int(4)] = (d_72_v27_).set(d_71_v26_, d_8_v2_)
                nw7_[int(5)] = (d_72_v27_) | (d_72_v27_)
                nw7_[int(6)] = d_72_v27_
                nw7_[int(7)] = d_72_v27_
                nw7_[int(8)] = d_72_v27_
                nw7_[int(9)] = d_72_v27_
                nw7_[int(10)] = d_72_v27_
                d_73_v28_ = nw7_
                d_74_v29_: _dafny.MultiSet
                d_74_v29_ = _dafny.MultiSet([d_8_v2_])
                index4_ = default__.safeIndex(487, (d_73_v28_).length(0))
                (d_73_v28_)[index4_] = (_dafny.Map({d_71_v26_: (d_12_v5_)[default__.safeIndex((d_74_v29_).cardinality, len(d_12_v5_))]})) | (d_72_v27_)
                index5_ = default__.safeIndex(487, (d_73_v28_).length(0))
                (d_73_v28_)[index5_] = d_72_v27_
                d_75_v30_: _dafny.Map
                d_75_v30_ = _dafny.Map({(d_23_v15_).f18: (d_23_v15_).f18})
                d_76_v31_: _dafny.Array
                nw8_ = _dafny.Array(None, 14)
                nw8_[int(0)] = d_46_i4_
                nw8_[int(1)] = (d_42_v18_)[default__.safeIndex(580, (d_42_v18_).length(0))]
                nw8_[int(2)] = (d_71_v26_).f18
                nw8_[int(3)] = (d_23_v15_).f18
                nw8_[int(4)] = (d_71_v26_).f18
                nw8_[int(5)] = (d_42_v18_)[default__.safeIndex(580, (d_42_v18_).length(0))]
                nw8_[int(6)] = (0) - ((d_23_v15_).f18)
                nw8_[int(7)] = (d_42_v18_)[default__.safeIndex(580, (d_42_v18_).length(0))]
                nw8_[int(8)] = (0) - ((d_23_v15_).f18)
                nw8_[int(9)] = (d_23_v15_).f18
                nw8_[int(10)] = (d_23_v15_).f18
                nw8_[int(11)] = len((d_75_v30_).set(p1, len(d_12_v5_)))
                nw8_[int(12)] = (d_42_v18_)[default__.safeIndex(580, (d_42_v18_).length(0))]
                nw8_[int(13)] = (d_23_v15_).f18
                d_76_v31_ = nw8_
                d_77_v32_: D3
                d_77_v32_ = D3_DC10((True if d_8_v2_ else d_8_v2_), d_76_v31_, (d_6_v0_) + (d_6_v0_))
                d_77_v32_ = d_77_v32_
            elif True:
                d_78___mcc_h15_ = source1_.cf11
                d_79_cf11_ = d_78___mcc_h15_
                (globalState).f0 = d_46_i4_
                index6_ = default__.safeIndex(960, (d_48_v20_).length(0))
                (d_48_v20_)[index6_] = d_8_v2_
                index7_ = default__.safeIndex(960, (d_48_v20_).length(0))
                (d_48_v20_)[index7_] = d_8_v2_
                d_80_v33_: _dafny.Array
                def lambda6_(d_81_v7_):
                    def lambda7_(d_82_i7_):
                        return d_81_v7_

                    return lambda7_

                init3_ = lambda6_(d_14_v7_)
                nw9_ = _dafny.Array(None, 18)
                for i0_3_ in range(nw9_.length(0)):
                    nw9_[i0_3_] = init3_(i0_3_)
                d_80_v33_ = nw9_
                d_80_v33_ = d_80_v33_
                index8_ = default__.safeIndex(960, (d_48_v20_).length(0))
                (d_48_v20_)[index8_] = not ((d_48_v20_)[default__.safeIndex(960, (d_48_v20_).length(0))]) or ((d_48_v20_)[default__.safeIndex(960, (d_48_v20_).length(0))])
        pat_let_tv5_ = d_7_v1_
        pat_let_tv6_ = d_14_v7_
        def iife4_(_pat_let0_0):
            def iife5_(d_83_dt__update__tmp_h0_):
                def iife6_(_pat_let1_0):
                    def iife7_(d_84_dt__update_hcf8_h0_):
                        def iife8_(_pat_let2_0):
                            def iife9_(d_86_dt__update_hcf6_h0_):
                                return D1_DC4(d_86_dt__update_hcf6_h0_, (d_83_dt__update__tmp_h0_).cf7, d_84_dt__update_hcf8_h0_, (d_83_dt__update__tmp_h0_).cf9, (d_83_dt__update__tmp_h0_).cf10)
                            return iife9_(_pat_let2_0)
                        return iife8_(_dafny.SeqWithoutIsStrInference([pat_let_tv6_ for d_85_i8_ in range(default__.abs(335))]))
                    return iife7_(_pat_let1_0)
                return iife6_(pat_let_tv5_)
            return iife5_(_pat_let0_0)
        source2_ = iife4_(d_20_v13_)
        if source2_.is_DC3:
            d_87___mcc_h16_ = source2_.cf5
            d_88_cf5_ = d_87___mcc_h16_
            d_89_v34_: _dafny.Array
            nw10_ = _dafny.Array(False, 5)
            d_89_v34_ = nw10_
            d_90_v35_: D2
            d_90_v35_ = D2_DC7(d_89_v34_)
            d_91_v36_: D2
            d_91_v36_ = D2_DC8(d_90_v35_)
            d_91_v36_ = d_91_v36_
            index9_ = default__.safeIndex(898, (d_89_v34_).length(0))
            (d_89_v34_)[index9_] = not (d_8_v2_) or (d_8_v2_)
            index10_ = default__.safeIndex(898, (d_89_v34_).length(0))
            (d_89_v34_)[index10_] = d_8_v2_
            d_89_v34_ = (d_89_v34_ if d_8_v2_ else d_89_v34_)
            d_92_v37_: D0
            d_92_v37_ = D0_DC1(default__.fm3(globalState))
            d_92_v37_ = p0
        elif source2_.is_DC4:
            d_93___mcc_h17_ = source2_.cf6
            d_94___mcc_h18_ = source2_.cf7
            d_95___mcc_h19_ = source2_.cf8
            d_96___mcc_h20_ = source2_.cf9
            d_97___mcc_h21_ = source2_.cf10
            d_98_cf10_ = d_97___mcc_h21_
            d_99_cf9_ = d_96___mcc_h20_
            d_100_cf8_ = d_95___mcc_h19_
            d_101_cf7_ = d_94___mcc_h18_
            d_102_cf6_ = d_93___mcc_h17_
            d_103_v38_: _dafny.Array
            def lambda8_(d_104_v5_, d_105_v15_, d_106_v3_, d_107_p1_):
                def lambda9_(d_108_i9_):
                    return ((d_104_v5_)[default__.safeIndex(len(_dafny.Map({len(_dafny.Set({(d_105_v15_).f18, (d_105_v15_).f18, len(d_106_v3_)})): d_107_p1_})), len(d_104_v5_))]) not in (d_106_v3_)

                return lambda9_

            init4_ = lambda8_(d_12_v5_, d_23_v15_, d_9_v3_, p1)
            nw11_ = _dafny.Array(None, 27)
            for i0_4_ in range(nw11_.length(0)):
                nw11_[i0_4_] = init4_(i0_4_)
            d_103_v38_ = nw11_
            rhs3_ = d_23_v15_
            rhs4_ = d_103_v38_
            d_23_v15_ = rhs3_
            d_103_v38_ = rhs4_
            d_109_v39_: _dafny.MultiSet
            d_109_v39_ = _dafny.MultiSet([d_101_cf7_, d_101_cf7_])
            index11_ = default__.safeIndex(62, (d_42_v18_).length(0))
            (d_42_v18_)[index11_] = ((d_109_v39_ if d_8_v2_ else default__.fm7(p1, d_12_v5_, d_8_v2_, globalState))).cardinality
            index12_ = default__.safeIndex(62, (d_42_v18_).length(0))
            (d_42_v18_)[index12_] = p1
            d_110_v40_: _dafny.Map
            d_110_v40_ = _dafny.Map({d_23_v15_: (d_23_v15_).f18})
            d_110_v40_ = ((d_110_v40_) | (d_110_v40_)) | (d_110_v40_)
            d_14_v7_ = d_14_v7_
        elif source2_.is_DC2:
            d_111___mcc_h22_ = source2_.cf4
            d_112_cf4_ = d_111___mcc_h22_
            d_113_v41_: _dafny.Map
            d_113_v41_ = _dafny.Map({d_8_v2_: True})
            d_113_v41_ = (d_113_v41_) | (d_113_v41_)
            if d_8_v2_:
                d_114_v42_: _dafny.Array
                nw12_ = _dafny.Array(False, 15)
                d_114_v42_ = nw12_
                d_115_v43_: _dafny.Seq
                d_115_v43_ = _dafny.SeqWithoutIsStrInference([d_114_v42_])
                d_116_v44_: _dafny.Map
                d_116_v44_ = _dafny.Map({d_23_v15_: d_115_v43_})
                d_117_v45_: _dafny.Map
                d_117_v45_ = _dafny.Map({d_8_v2_: d_42_v18_})
                d_118_v46_: _dafny.Map
                d_118_v46_ = _dafny.Map({len(d_117_v45_): d_115_v43_})
                d_119_v47_: _dafny.Array
                nw13_ = _dafny.Array(None, 9)
                nw13_[int(0)] = d_115_v43_
                nw13_[int(1)] = (d_115_v43_) + (((d_116_v44_)[d_23_v15_] if (d_23_v15_) in (d_116_v44_) else d_115_v43_))
                nw13_[int(2)] = ((d_115_v43_).set(default__.safeIndex(p1, len(d_115_v43_)), d_114_v42_)) + (d_115_v43_)
                nw13_[int(3)] = d_115_v43_
                nw13_[int(4)] = ((d_118_v46_)[p1] if (p1) in (d_118_v46_) else d_115_v43_)
                nw13_[int(5)] = d_115_v43_
                nw13_[int(6)] = d_115_v43_
                nw13_[int(7)] = d_115_v43_
                nw13_[int(8)] = (d_115_v43_) + (d_115_v43_)
                d_119_v47_ = nw13_
                index13_ = default__.safeIndex(565, (d_119_v47_).length(0))
                (d_119_v47_)[index13_] = (_dafny.SeqWithoutIsStrInference([d_114_v42_, d_114_v42_])) + (d_115_v43_)
                index14_ = default__.safeIndex(368, (d_42_v18_).length(0))
                (d_42_v18_)[index14_] = (0) - ((0) - (len(d_24_v16_)))
                d_120_v48_: D3
                d_120_v48_ = D3_DC10(not (d_8_v2_) or (False), d_42_v18_, ((_dafny.SeqWithoutIsStrInference([d_14_v7_ for d_121_i10_ in range(default__.abs(-28))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i")))).set(default__.safeIndex((d_23_v15_).f18, len((_dafny.SeqWithoutIsStrInference([d_14_v7_ for d_122_i10_ in range(default__.abs(-28))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))))), d_14_v7_))
                d_123_v49_: _dafny.Map
                d_123_v49_ = _dafny.Map({d_8_v2_: d_23_v15_})
                index15_ = default__.safeIndex(565, (d_119_v47_).length(0))
                index16_ = default__.safeIndex(368, (d_42_v18_).length(0))
                rhs5_ = (d_115_v43_) + (d_115_v43_)
                rhs6_ = d_6_v0_
                rhs7_ = len((d_123_v49_) | (d_123_v49_))
                rhs8_ = d_120_v48_
                lhs3_ = d_119_v47_
                lhs4_ = default__.safeIndex(565, (d_119_v47_).length(0))
                lhs5_ = d_42_v18_
                lhs6_ = default__.safeIndex(368, (d_42_v18_).length(0))
                lhs3_[lhs4_] = rhs5_
                d_6_v0_ = rhs6_
                lhs5_[lhs6_] = rhs7_
                d_120_v48_ = rhs8_
                d_124_v50_: _dafny.Map
                d_124_v50_ = _dafny.Map({(d_42_v18_)[default__.safeIndex(368, (d_42_v18_).length(0))]: (d_23_v15_).f18})
                d_124_v50_ = (d_124_v50_).set((d_42_v18_)[default__.safeIndex(368, (d_42_v18_).length(0))], (d_23_v15_).f18)
                (globalState).f0 = p1
                d_114_v42_ = d_114_v42_
                d_125_v51_: _dafny.MultiSet
                d_125_v51_ = _dafny.MultiSet([not(d_8_v2_), True])
                d_126_v52_: C0
                nw14_ = C0()
                nw14_.ctor__(((d_125_v51_).intersection(d_125_v51_)).cardinality)
                d_126_v52_ = nw14_
            elif True:
                d_127_v53_: _dafny.Array
                nw15_ = _dafny.Array(_dafny.Set({}), 20)
                d_127_v53_ = nw15_
                d_127_v53_ = d_127_v53_
                d_128_v54_: D1
                d_128_v54_ = D1_DC2(d_9_v3_)
                d_129_v55_: _dafny.Map
                d_129_v55_ = _dafny.Map({d_128_v54_: (d_23_v15_).f18})
                d_129_v55_ = d_129_v55_
                d_130_v56_: _dafny.Seq
                d_130_v56_ = _dafny.SeqWithoutIsStrInference([(d_23_v15_).f18, (d_23_v15_).f18, (d_23_v15_).f18])
                d_131_v57_: _dafny.Map
                d_131_v57_ = _dafny.Map({d_8_v2_: (d_130_v56_) + (d_130_v56_)})
                r1 = ((d_131_v57_)[d_8_v2_] if (d_8_v2_) in (d_131_v57_) else (d_130_v56_ if d_8_v2_ else _dafny.SeqWithoutIsStrInference([(0) - (len(d_112_cf4_))])))
                d_132_v58_: _dafny.Array
                def lambda10_(d_133_v2_):
                    def lambda11_(d_134_i11_):
                        return d_133_v2_

                    return lambda11_

                init5_ = lambda10_(d_8_v2_)
                nw16_ = _dafny.Array(None, 10)
                for i0_5_ in range(nw16_.length(0)):
                    nw16_[i0_5_] = init5_(i0_5_)
                d_132_v58_ = nw16_
                index17_ = default__.safeIndex(781, (d_132_v58_).length(0))
                (d_132_v58_)[index17_] = True
                index18_ = default__.safeIndex(781, (d_132_v58_).length(0))
                (d_132_v58_)[index18_] = d_8_v2_
                (globalState).f15 = (d_23_v15_).f18
            d_135_v59_: _dafny.MultiSet
            d_135_v59_ = _dafny.MultiSet([False])
            r3 = ((_dafny.MultiSet([d_8_v2_])).set(not(d_8_v2_), default__.abs(p1))).issubset((d_135_v59_) - (_dafny.MultiSet([not(default__.fm1(d_8_v2_, globalState)), d_8_v2_])))
            d_136_v60_: _dafny.Array
            nw17_ = _dafny.Array(_dafny.Seq({}), 24)
            d_136_v60_ = nw17_
            nw18_ = _dafny.Array(_dafny.Seq({}), 29)
            d_136_v60_ = nw18_
        elif True:
            d_137___mcc_h23_ = source2_.cf11
            d_138_cf11_ = d_137___mcc_h23_
            d_139_v61_: _dafny.Seq
            d_139_v61_ = _dafny.SeqWithoutIsStrInference([d_12_v5_, d_12_v5_, d_12_v5_])
            d_140_v62_: _dafny.Map
            d_140_v62_ = _dafny.Map({-890: d_14_v7_})
            d_141_v63_: _dafny.Seq
            d_141_v63_ = _dafny.SeqWithoutIsStrInference([len(default__.fm8(globalState)), (d_23_v15_).f18])
            d_142_v64_: _dafny.Seq
            d_142_v64_ = _dafny.SeqWithoutIsStrInference([len(d_141_v63_)])
            d_143_v65_: D3
            d_143_v65_ = D3_DC10(d_8_v2_, d_42_v18_, d_6_v0_)
            d_144_v66_: _dafny.Array
            nw19_ = _dafny.Array(None, 22)
            nw19_[int(0)] = True
            nw19_[int(1)] = (d_20_v13_).cf7
            nw19_[int(2)] = d_8_v2_
            nw19_[int(3)] = default__.fm1(d_8_v2_, globalState)
            nw19_[int(4)] = (d_139_v61_) <= (_dafny.SeqWithoutIsStrInference([d_12_v5_]))
            nw19_[int(5)] = not(True)
            nw19_[int(6)] = d_8_v2_
            nw19_[int(7)] = d_8_v2_
            nw19_[int(8)] = ((d_23_v15_).f18) != ((d_23_v15_).f18)
            nw19_[int(9)] = default__.fm1(d_8_v2_, globalState)
            nw19_[int(10)] = d_8_v2_
            nw19_[int(11)] = (d_140_v62_) == (_dafny.Map({(d_142_v64_)[default__.safeIndex(len(d_6_v0_), len(d_142_v64_))]: d_14_v7_}))
            nw19_[int(12)] = not(d_8_v2_)
            nw19_[int(13)] = True
            nw19_[int(14)] = (d_12_v5_)[default__.safeIndex(p1, len(d_12_v5_))]
            nw19_[int(15)] = d_8_v2_
            nw19_[int(16)] = ((d_23_v15_).f18) > (p1)
            nw19_[int(17)] = not(d_8_v2_)
            nw19_[int(18)] = (d_143_v65_).cf16
            nw19_[int(19)] = d_8_v2_
            nw19_[int(20)] = d_8_v2_
            nw19_[int(21)] = default__.fm1(not(False), globalState)
            d_144_v66_ = nw19_
            d_144_v66_ = d_144_v66_
            d_145_v67_: _dafny.Array
            nw20_ = _dafny.Array(_dafny.Seq({}), 21)
            d_145_v67_ = nw20_
            d_145_v67_ = d_145_v67_
            r3 = d_8_v2_
            d_146_v68_: _dafny.Map
            d_146_v68_ = _dafny.Map({d_144_v66_: d_23_v15_})
            d_146_v68_ = d_146_v68_
        pat_let_tv7_ = d_23_v15_
        def lambda12_(source3_):
            if source3_.is_DC0:
                d_147___mcc_h24_ = source3_.cf0
                d_148___mcc_h25_ = source3_.cf1
                d_149___mcc_h26_ = source3_.cf2
                d_150_cf2_ = d_149___mcc_h26_
                d_151_cf1_ = d_148___mcc_h25_
                d_152_cf0_ = d_147___mcc_h24_
                return (0) - ((d_152_cf0_) * (483))
            elif True:
                d_153___mcc_h27_ = source3_.cf3
                d_154_cf3_ = d_153___mcc_h27_
                return (pat_let_tv7_).f18

        r0 = (0) - (lambda12_(p0))
        d_155_v69_: _dafny.Seq
        d_155_v69_ = _dafny.SeqWithoutIsStrInference([(d_23_v15_).f18, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxu")))])
        r1 = d_155_v69_
        d_156_v70_: _dafny.Set
        d_156_v70_ = _dafny.Set({p1, default__.fm0(d_17_v10_, globalState)})
        r2 = (d_156_v70_) | (d_156_v70_)
        r3 = d_8_v2_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_157_v0_: bool
        d_157_v0_ = True
        d_158_v1_: _dafny.Set
        d_158_v1_ = _dafny.Set({d_157_v0_})
        d_159_v2_: _dafny.Array
        nw21_ = _dafny.Array(_dafny.Map({}), 15)
        d_159_v2_ = nw21_
        d_160_v3_: int
        d_160_v3_ = 639
        d_161_v4_: _dafny.MultiSet
        d_161_v4_ = _dafny.MultiSet([d_160_v3_, -408])
        d_162_v5_: _dafny.MultiSet
        d_162_v5_ = _dafny.MultiSet([d_157_v0_, d_157_v0_])
        d_163_v6_: _dafny.Seq
        d_163_v6_ = _dafny.SeqWithoutIsStrInference([(d_161_v4_).cardinality, (d_162_v5_).cardinality])
        d_164_v7_: _dafny.Array
        def lambda13_(d_165_v0_):
            def lambda14_(d_166_i0_):
                return d_165_v0_

            return lambda14_

        init6_ = lambda13_(d_157_v0_)
        nw22_ = _dafny.Array(None, 11)
        for i0_6_ in range(nw22_.length(0)):
            nw22_[i0_6_] = init6_(i0_6_)
        d_164_v7_ = nw22_
        d_167_v8_: _dafny.Seq
        d_167_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgx"))
        d_168_globalState_: GlobalState
        nw23_ = GlobalState()
        nw23_.ctor__(-206, (d_158_v1_).intersection(d_158_v1_), False, False, 633, 769, True, d_159_v2_, ((d_163_v6_).set(default__.safeIndex(d_160_v3_, len(d_163_v6_)), d_160_v3_)) + (d_163_v6_), d_161_v4_, -738, 714, _dafny.CodePoint('r'), False, 760, 654, d_164_v7_, d_167_v8_)
        d_168_globalState_ = nw23_
        d_169_v9_: _dafny.Map
        d_169_v9_ = _dafny.Map({d_160_v3_: d_160_v3_})
        d_170_v10_: _dafny.Map
        d_170_v10_ = _dafny.Map({d_157_v0_: len(d_169_v9_)})
        d_171_v11_: _dafny.Map
        d_171_v11_ = _dafny.Map({d_167_v8_: ((d_170_v10_)[d_157_v0_] if (d_157_v0_) in (d_170_v10_) else d_160_v3_)})
        d_172_v12_: D0
        d_172_v12_ = D0_DC0(d_160_v3_, d_171_v11_, d_161_v4_)
        source4_ = d_172_v12_
        if source4_.is_DC0:
            d_173___mcc_h0_ = source4_.cf0
            d_174___mcc_h1_ = source4_.cf1
            d_175___mcc_h2_ = source4_.cf2
            d_176_cf2_ = d_175___mcc_h2_
            d_177_cf1_ = d_174___mcc_h1_
            d_178_cf0_ = d_173___mcc_h0_
            d_179_v13_: D0
            d_179_v13_ = D0_DC1(D0_DC1(D0_DC0(len(d_163_v6_), d_171_v11_, d_176_cf2_)))
            d_180_v14_: D0
            d_180_v14_ = D0_DC1(d_179_v13_)
            d_181_v15_: int
            d_182_v16_: _dafny.Seq
            d_183_v17_: _dafny.Set
            d_184_v18_: bool
            out0_: int
            out1_: _dafny.Seq
            out2_: _dafny.Set
            out3_: bool
            out0_, out1_, out2_, out3_ = default__.m0(D0_DC1(d_180_v14_), default__.fm0(d_172_v12_, d_168_globalState_), d_168_globalState_)
            d_181_v15_ = out0_
            d_182_v16_ = out1_
            d_183_v17_ = out2_
            d_184_v18_ = out3_
            d_157_v0_ = not((default__.fm0(d_172_v12_, d_168_globalState_)) == (((0) - (d_160_v3_)) - (d_178_cf0_)))
            if default__.fm1(d_157_v0_, d_168_globalState_):
                d_185_v19_: D0
                d_185_v19_ = D0_DC1(d_179_v13_)
                d_186_v20_: int
                d_187_v21_: _dafny.Seq
                d_188_v22_: _dafny.Set
                d_189_v23_: bool
                out4_: int
                out5_: _dafny.Seq
                out6_: _dafny.Set
                out7_: bool
                out4_, out5_, out6_, out7_ = default__.m0((d_185_v19_ if d_157_v0_ else d_185_v19_), default__.safeModuloInt(len(d_167_v8_), d_178_cf0_), d_168_globalState_)
                d_186_v20_ = out4_
                d_187_v21_ = out5_
                d_188_v22_ = out6_
                d_189_v23_ = out7_
                (d_168_globalState_).f13 = default__.fm1(d_189_v23_, d_168_globalState_)
                (d_168_globalState_).f13 = d_184_v18_
                d_170_v10_ = (d_170_v10_).set((d_167_v8_) < (d_167_v8_), d_160_v3_)
                (d_168_globalState_).f0 = d_160_v3_
            elif True:
                d_172_v12_ = d_172_v12_
                d_190_v24_: D0
                d_190_v24_ = D0_DC1(d_179_v13_)
                d_191_v25_: int
                d_192_v26_: _dafny.Seq
                d_193_v27_: _dafny.Set
                d_194_v28_: bool
                out8_: int
                out9_: _dafny.Seq
                out10_: _dafny.Set
                out11_: bool
                out8_, out9_, out10_, out11_ = default__.m0(d_190_v24_, (len(d_167_v8_)) * (d_160_v3_), d_168_globalState_)
                d_191_v25_ = out8_
                d_192_v26_ = out9_
                d_193_v27_ = out10_
                d_194_v28_ = out11_
                (d_168_globalState_).f13 = True
                d_160_v3_ = default__.safeModuloInt(472, len(d_167_v8_))
                d_195_v29_: C0
                nw24_ = C0()
                nw24_.ctor__(d_191_v25_)
                d_195_v29_ = nw24_
            d_196_v30_: _dafny.Array
            nw25_ = _dafny.Array(int(0), 6)
            d_196_v30_ = nw25_
            index19_ = default__.safeIndex(841, (d_196_v30_).length(0))
            (d_196_v30_)[index19_] = (d_160_v3_) - (d_178_cf0_)
            index20_ = default__.safeIndex(841, (d_196_v30_).length(0))
            rhs9_ = d_178_cf0_
            rhs10_ = d_160_v3_
            rhs11_ = not(((d_178_cf0_) * (d_181_v15_)) > (len(d_169_v9_)))
            lhs7_ = d_196_v30_
            lhs8_ = default__.safeIndex(841, (d_196_v30_).length(0))
            lhs7_[lhs8_] = rhs9_
            d_160_v3_ = rhs10_
            d_157_v0_ = rhs11_
        elif True:
            d_197___mcc_h3_ = source4_.cf3
            d_198_cf3_ = d_197___mcc_h3_
            source5_ = (d_172_v12_ if False else d_172_v12_)
            if source5_.is_DC0:
                d_199___mcc_h4_ = source5_.cf0
                d_200___mcc_h5_ = source5_.cf1
                d_201___mcc_h6_ = source5_.cf2
                d_202_cf2_ = d_201___mcc_h6_
                d_203_cf1_ = d_200___mcc_h5_
                d_204_cf0_ = d_199___mcc_h4_
                d_205_v31_: _dafny.Map
                d_205_v31_ = _dafny.Map({((d_162_v5_)[d_157_v0_] if (d_157_v0_) in (d_162_v5_) else d_160_v3_): d_157_v0_})
                d_157_v0_ = default__.fm1(((d_205_v31_)[d_204_cf0_] if (d_204_cf0_) in (d_205_v31_) else d_157_v0_), d_168_globalState_)
                (d_168_globalState_).f15 = ((len(_dafny.Map({d_157_v0_: d_157_v0_}))) + (len(d_167_v8_)) if (d_158_v1_).isdisjoint(d_158_v1_) else (0) - (d_160_v3_))
                d_206_v33_: _dafny.Map
                d_206_v33_ = _dafny.Map({d_167_v8_: _dafny.SeqWithoutIsStrInference([not(default__.fm1(False, d_168_globalState_))])})
                d_207_v34_: D0
                def iife10_():
                    coll4_ = _dafny.Map()
                    compr_4_: _dafny.Seq
                    for compr_4_ in (d_206_v33_).keys.Elements:
                        d_208_v32_: _dafny.Seq = compr_4_
                        if (d_208_v32_) in (d_206_v33_):
                            coll4_[d_208_v32_] = (d_172_v12_).cf0
                    return _dafny.Map(coll4_)
                d_207_v34_ = D0_DC0(d_204_cf0_, iife10_()
, (d_161_v4_).set(((d_161_v4_)[d_204_cf0_] if (d_204_cf0_) in (d_161_v4_) else d_160_v3_), default__.abs(d_160_v3_)))
                d_209_v35_: D0
                d_209_v35_ = D0_DC1(d_207_v34_)
                d_210_v36_: D0
                d_210_v36_ = D0_DC1(d_209_v35_)
                d_211_v37_: D0
                d_211_v37_ = D0_DC1(d_210_v36_)
                d_212_v38_: _dafny.Set
                d_212_v38_ = _dafny.Set({d_167_v8_})
                d_213_v39_: int
                d_214_v40_: _dafny.Seq
                d_215_v41_: _dafny.Set
                d_216_v42_: bool
                out12_: int
                out13_: _dafny.Seq
                out14_: _dafny.Set
                out15_: bool
                out12_, out13_, out14_, out15_ = default__.m0(d_211_v37_, len(_dafny.SeqWithoutIsStrInference([(0) - ((0) - ((0) - (d_204_cf0_))), d_160_v3_, (0) - (len(_dafny.SeqWithoutIsStrInference([(d_167_v8_)[default__.safeIndex(d_160_v3_, len(d_167_v8_))] for d_217_i1_ in range(default__.abs(830))]))), d_160_v3_, (0) - (len(d_212_v38_))])), d_168_globalState_)
                d_213_v39_ = out12_
                d_214_v40_ = out13_
                d_215_v41_ = out14_
                d_216_v42_ = out15_
                d_218_v43_: C0
                nw26_ = C0()
                nw26_.ctor__(d_204_cf0_)
                d_218_v43_ = nw26_
                d_218_v43_ = d_218_v43_
            elif True:
                d_219___mcc_h7_ = source5_.cf3
                d_220_cf3_ = d_219___mcc_h7_
                d_221_v44_: C0
                nw27_ = C0()
                nw27_.ctor__((d_160_v3_) - (d_160_v3_))
                d_221_v44_ = nw27_
                d_221_v44_ = d_221_v44_
                d_222_v45_: D0
                d_222_v45_ = D0_DC0(d_160_v3_, d_171_v11_, d_161_v4_)
                d_223_v46_: int
                d_224_v47_: _dafny.Seq
                d_225_v48_: _dafny.Set
                d_226_v49_: bool
                out16_: int
                out17_: _dafny.Seq
                out18_: _dafny.Set
                out19_: bool
                out16_, out17_, out18_, out19_ = default__.m0(D0_DC1(d_222_v45_), (D0_DC0(len(d_163_v6_), d_171_v11_, default__.fm2(d_157_v0_, d_168_globalState_))).cf0, d_168_globalState_)
                d_223_v46_ = out16_
                d_224_v47_ = out17_
                d_225_v48_ = out18_
                d_226_v49_ = out19_
                d_227_v50_: D0
                d_227_v50_ = D0_DC1(d_222_v45_)
                d_228_v51_: int
                d_229_v52_: _dafny.Seq
                d_230_v53_: _dafny.Set
                d_231_v54_: bool
                out20_: int
                out21_: _dafny.Seq
                out22_: _dafny.Set
                out23_: bool
                out20_, out21_, out22_, out23_ = default__.m0(d_227_v50_, (d_162_v5_).cardinality, d_168_globalState_)
                d_228_v51_ = out20_
                d_229_v52_ = out21_
                d_230_v53_ = out22_
                d_231_v54_ = out23_
                pat_let_tv8_ = d_222_v45_
                d_232_v55_: int
                d_233_v56_: _dafny.Seq
                d_234_v57_: _dafny.Set
                d_235_v58_: bool
                out24_: int
                out25_: _dafny.Seq
                out26_: _dafny.Set
                out27_: bool
                def iife11_(_pat_let3_0):
                    def iife12_(d_236_dt__update__tmp_h0_):
                        def iife13_(_pat_let4_0):
                            def iife14_(d_237_dt__update_hcf3_h0_):
                                return D0_DC1(d_237_dt__update_hcf3_h0_)
                            return iife14_(_pat_let4_0)
                        return iife13_(pat_let_tv8_)
                    return iife12_(_pat_let3_0)
                out24_, out25_, out26_, out27_ = default__.m0(iife11_(d_227_v50_), len((_dafny.Set({(d_221_v44_).f18, d_223_v46_})) - (d_225_v48_)), d_168_globalState_)
                d_232_v55_ = out24_
                d_233_v56_ = out25_
                d_234_v57_ = out26_
                d_235_v58_ = out27_
            d_238_v59_: _dafny.MultiSet
            d_238_v59_ = _dafny.MultiSet([_dafny.MultiSet([False]), _dafny.MultiSet([d_157_v0_]), d_162_v5_, d_162_v5_])
            d_239_v60_: _dafny.Seq
            d_239_v60_ = _dafny.SeqWithoutIsStrInference([d_162_v5_])
            d_238_v59_ = _dafny.MultiSet((d_239_v60_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_157_v0_, False, d_157_v0_])) for d_240_i2_ in range(default__.abs(-680))])))
            (d_168_globalState_).f0 = 12
            (d_168_globalState_).f10 = d_160_v3_
        d_241_v62_: _dafny.Seq
        d_241_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_163_v6_), d_161_v4_, d_161_v4_, d_161_v4_])
        d_242_v63_: C0
        nw28_ = C0()
        nw28_.ctor__((d_161_v4_).cardinality)
        d_242_v63_ = nw28_
        d_243_v64_: _dafny.Map
        d_243_v64_ = _dafny.Map({d_157_v0_: d_242_v63_})
        d_244_v65_: D1
        d_244_v65_ = D1_DC2(_dafny.Map({not(default__.fm1(False, d_168_globalState_)): d_160_v3_}))
        d_245_v66_: _dafny.Array
        nw29_ = _dafny.Array(None, 14)
        nw29_[int(0)] = default__.fm0(default__.fm3(d_168_globalState_), d_168_globalState_)
        nw29_[int(1)] = d_160_v3_
        nw29_[int(2)] = (d_160_v3_) + (d_160_v3_)
        nw29_[int(3)] = d_160_v3_
        nw29_[int(4)] = d_160_v3_
        nw29_[int(5)] = -437
        nw29_[int(6)] = (d_163_v6_)[default__.safeIndex(default__.fm0(D0_DC0(d_160_v3_, _dafny.Map({d_167_v8_: d_160_v3_}), d_161_v4_), d_168_globalState_), len(d_163_v6_))]
        def iife15_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.MultiSet
            for compr_5_ in (d_241_v62_).Elements:
                d_246_v61_: _dafny.MultiSet = compr_5_
                if (d_246_v61_) in (d_241_v62_):
                    coll5_[d_246_v61_] = d_160_v3_
            return _dafny.Map(coll5_)
        nw29_[int(7)] = (0) - (len(iife15_()
        ))
        nw29_[int(8)] = d_160_v3_
        nw29_[int(9)] = len((d_167_v8_) + (d_167_v8_))
        nw29_[int(10)] = len(d_243_v64_)
        nw29_[int(11)] = len((d_244_v65_).cf4)
        nw29_[int(12)] = (len(d_170_v10_)) + ((d_242_v63_).f18)
        nw29_[int(13)] = d_160_v3_
        d_245_v66_ = nw29_
        index21_ = default__.safeIndex(580, (d_245_v66_).length(0))
        (d_245_v66_)[index21_] = ((d_242_v63_).f18) * (len(d_167_v8_))
        index22_ = default__.safeIndex(831, (d_245_v66_).length(0))
        (d_245_v66_)[index22_] = ((d_242_v63_).f18) * (d_160_v3_)
        d_247_v67_: _dafny.Set
        d_247_v67_ = _dafny.Set({(d_242_v63_).f18, (d_242_v63_).f18})
        d_248_v68_: _dafny.Seq
        d_248_v68_ = _dafny.SeqWithoutIsStrInference([d_157_v0_])
        index23_ = default__.safeIndex(580, (d_245_v66_).length(0))
        index24_ = default__.safeIndex(831, (d_245_v66_).length(0))
        rhs12_ = (d_160_v3_ if d_157_v0_ else len(d_247_v67_))
        rhs13_ = default__.safeModuloInt(((d_242_v63_).f18) - (len(d_167_v8_)), len(d_248_v68_))
        rhs14_ = (d_242_v63_).f18
        lhs9_ = d_168_globalState_
        lhs10_ = d_245_v66_
        lhs11_ = default__.safeIndex(580, (d_245_v66_).length(0))
        lhs12_ = d_245_v66_
        lhs13_ = default__.safeIndex(831, (d_245_v66_).length(0))
        lhs9_.f10 = rhs12_
        lhs10_[lhs11_] = rhs13_
        lhs12_[lhs13_] = rhs14_
        d_249_v69_: C0
        nw30_ = C0()
        nw30_.ctor__(d_160_v3_)
        d_249_v69_ = nw30_
        d_161_v4_ = d_161_v4_
        d_250_v70_: _dafny.Map
        d_250_v70_ = _dafny.Map({(d_242_v63_).f18: not(d_157_v0_)})
        if ((d_250_v70_)[default__.safeModuloInt((d_242_v63_).f18, -905)] if (default__.safeModuloInt((d_242_v63_).f18, -905)) in (d_250_v70_) else not(default__.fm1(d_157_v0_, d_168_globalState_))):
            d_251_v71_: _dafny.Array
            def lambda15_(d_252_v0_, d_253_v10_):
                def lambda16_(d_254_i3_):
                    return (d_253_v10_ if d_252_v0_ else d_253_v10_)

                return lambda16_

            init7_ = lambda15_(d_157_v0_, d_170_v10_)
            nw31_ = _dafny.Array(None, 4)
            for i0_7_ in range(nw31_.length(0)):
                nw31_[i0_7_] = init7_(i0_7_)
            d_251_v71_ = nw31_
            index25_ = default__.safeIndex(584, (d_251_v71_).length(0))
            (d_251_v71_)[index25_] = d_170_v10_
            index26_ = default__.safeIndex(584, (d_251_v71_).length(0))
            (d_251_v71_)[index26_] = d_170_v10_
            d_255_v72_: _dafny.Array
            nw32_ = _dafny.Array(_dafny.MultiSet({}), 13)
            d_255_v72_ = nw32_
            d_256_v73_: _dafny.Array
            nw33_ = _dafny.Array(None, 2)
            nw33_[int(0)] = d_244_v65_
            nw33_[int(1)] = d_244_v65_
            d_256_v73_ = nw33_
            index27_ = default__.safeIndex(458, (d_255_v72_).length(0))
            (d_255_v72_)[index27_] = _dafny.MultiSet([d_256_v73_, d_256_v73_, d_256_v73_, d_256_v73_, d_256_v73_])
            d_257_v74_: _dafny.MultiSet
            d_257_v74_ = _dafny.MultiSet([d_256_v73_, d_256_v73_])
            index28_ = default__.safeIndex(458, (d_255_v72_).length(0))
            (d_255_v72_)[index28_] = ((d_257_v74_) - (_dafny.MultiSet([d_256_v73_, d_256_v73_]))).intersection((d_257_v74_).intersection(d_257_v74_))
            d_258_v75_: _dafny.Array
            nw34_ = _dafny.Array(None, 25)
            d_258_v75_ = nw34_
            index29_ = default__.safeIndex(299, (d_258_v75_).length(0))
            (d_258_v75_)[index29_] = d_242_v63_
            index30_ = default__.safeIndex(299, (d_258_v75_).length(0))
            (d_258_v75_)[index30_] = d_242_v63_
            d_259_v76_: _dafny.Array
            nw35_ = _dafny.Array(_dafny.Map({}), 17)
            d_259_v76_ = nw35_
            d_259_v76_ = d_259_v76_
            d_260_v77_: D1
            d_260_v77_ = D1_DC3(d_158_v1_)
            d_261_v78_: _dafny.Array
            nw36_ = _dafny.Array(None, 16)
            nw36_[int(0)] = D1_DC3(d_158_v1_)
            nw36_[int(1)] = d_260_v77_
            nw36_[int(2)] = d_260_v77_
            nw36_[int(3)] = d_260_v77_
            nw36_[int(4)] = d_260_v77_
            nw36_[int(5)] = d_260_v77_
            nw36_[int(6)] = D1_DC3(d_158_v1_)
            nw36_[int(7)] = D1_DC3(d_158_v1_)
            nw36_[int(8)] = d_260_v77_
            nw36_[int(9)] = d_260_v77_
            nw36_[int(10)] = (d_260_v77_ if True else d_260_v77_)
            nw36_[int(11)] = d_260_v77_
            nw36_[int(12)] = d_260_v77_
            nw36_[int(13)] = D1_DC3(default__.fm4((d_242_v63_).f18, (d_249_v69_).f18, (d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))], (d_242_v63_).f18, d_168_globalState_))
            nw36_[int(14)] = D1_DC3(d_158_v1_)
            nw36_[int(15)] = d_260_v77_
            d_261_v78_ = nw36_
            index31_ = default__.safeIndex(53, (d_261_v78_).length(0))
            (d_261_v78_)[index31_] = d_260_v77_
            index32_ = default__.safeIndex(53, (d_261_v78_).length(0))
            (d_261_v78_)[index32_] = d_260_v77_
        elif True:
            index33_ = default__.safeIndex(826, (d_164_v7_).length(0))
            (d_164_v7_)[index33_] = d_157_v0_
            index34_ = default__.safeIndex(826, (d_164_v7_).length(0))
            (d_164_v7_)[index34_] = d_157_v0_
            d_262_v79_: D0
            d_262_v79_ = D0_DC0((d_242_v63_).f18, ((d_171_v11_).set(d_167_v8_, (_dafny.MultiSet(d_248_v68_)).cardinality)).set(d_167_v8_, (d_242_v63_).f18), d_161_v4_)
            d_263_v80_: D0
            d_263_v80_ = D0_DC1(d_262_v79_)
            d_264_v81_: D0
            d_264_v81_ = D0_DC1(d_262_v79_)
            d_265_v82_: int
            d_266_v83_: _dafny.Seq
            d_267_v84_: _dafny.Set
            d_268_v85_: bool
            out28_: int
            out29_: _dafny.Seq
            out30_: _dafny.Set
            out31_: bool
            out28_, out29_, out30_, out31_ = default__.m0(d_264_v81_, ((d_242_v63_).f18 if d_157_v0_ else -94), d_168_globalState_)
            d_265_v82_ = out28_
            d_266_v83_ = out29_
            d_267_v84_ = out30_
            d_268_v85_ = out31_
            d_269_v86_: C0
            nw37_ = C0()
            nw37_.ctor__(d_160_v3_)
            d_269_v86_ = nw37_
            d_270_v87_: _dafny.Map
            d_270_v87_ = _dafny.Map({(d_164_v7_)[default__.safeIndex(826, (d_164_v7_).length(0))]: d_264_v81_})
            d_270_v87_ = (d_270_v87_).set(d_268_v85_, D0_DC1(d_262_v79_))
            (d_168_globalState_).f13 = d_157_v0_
        d_169_v9_ = (d_169_v9_).set((d_242_v63_).f18, (d_242_v63_).f18)
        d_271_v88_: str
        d_271_v88_ = _dafny.CodePoint('j')
        d_272_v89_: _dafny.Map
        d_272_v89_ = _dafny.Map({_dafny.MultiSet([(d_249_v69_).f18, (d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))], d_160_v3_, d_160_v3_, -242]): d_271_v88_})
        d_273_v90_: _dafny.Map
        d_273_v90_ = _dafny.Map({d_157_v0_: d_272_v89_})
        d_272_v89_ = ((d_273_v90_)[True] if (True) in (d_273_v90_) else d_272_v89_)
        (d_168_globalState_).f4 = (0) - (d_160_v3_)
        d_157_v0_ = not(d_157_v0_)
        d_274_v91_: _dafny.Array
        nw38_ = _dafny.Array(_dafny.Seq({}), 18)
        d_274_v91_ = nw38_
        index35_ = default__.safeIndex(747, (d_274_v91_).length(0))
        (d_274_v91_)[index35_] = d_163_v6_
        index36_ = default__.safeIndex(747, (d_274_v91_).length(0))
        (d_274_v91_)[index36_] = (d_163_v6_) + (d_163_v6_)
        d_164_v7_ = d_164_v7_
        d_157_v0_ = d_157_v0_
        if ((d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))]) == (len(d_167_v8_)):
            d_275_v92_: _dafny.Seq
            d_275_v92_ = _dafny.SeqWithoutIsStrInference([d_249_v69_, d_249_v69_, d_242_v63_, d_249_v69_])
            d_276_v93_: _dafny.Map
            d_276_v93_ = _dafny.Map({d_271_v88_: (d_275_v92_) < (_dafny.SeqWithoutIsStrInference([d_242_v63_, d_242_v63_]))})
            d_276_v93_ = (d_276_v93_).set(d_271_v88_, d_157_v0_)
            if d_157_v0_:
                d_277_v94_: _dafny.Map
                d_277_v94_ = _dafny.Map({54: d_172_v12_})
                d_278_v96_: _dafny.Seq
                def iife16_():
                    coll6_ = _dafny.Set()
                    compr_6_: int
                    for compr_6_ in (d_277_v94_).keys.Elements:
                        d_279_v95_: int = compr_6_
                        if (d_279_v95_) in (d_277_v94_):
                            coll6_ = coll6_.union(_dafny.Set([(d_279_v95_) + (514)]))
                    return _dafny.Set(coll6_)
                d_278_v96_ = _dafny.SeqWithoutIsStrInference([d_247_v67_, d_247_v67_, iife16_()
                ])
                d_278_v96_ = (d_278_v96_) + (d_278_v96_)
                d_280_v97_: D1
                d_280_v97_ = D1_DC3(d_158_v1_)
                d_281_v98_: _dafny.Map
                d_281_v98_ = _dafny.Map({d_158_v1_: d_167_v8_})
                pat_let_tv9_ = d_158_v1_
                def iife17_(_pat_let5_0):
                    def iife18_(d_282_dt__update__tmp_h1_):
                        def iife19_(_pat_let6_0):
                            def iife20_(d_283_dt__update_hcf5_h0_):
                                return D1_DC3(d_283_dt__update_hcf5_h0_)
                            return iife20_(_pat_let6_0)
                        return iife19_(pat_let_tv9_)
                    return iife18_(_pat_let5_0)
                d_157_v0_ = ((iife17_(d_280_v97_)).cf5) not in (d_281_v98_)
                d_284_v99_: D0
                d_284_v99_ = D0_DC0((d_249_v69_).f18, d_171_v11_, d_161_v4_)
                d_285_v100_: D0
                d_285_v100_ = D0_DC1(d_284_v99_)
                d_286_v101_: int
                d_287_v102_: _dafny.Seq
                d_288_v103_: _dafny.Set
                d_289_v104_: bool
                out32_: int
                out33_: _dafny.Seq
                out34_: _dafny.Set
                out35_: bool
                out32_, out33_, out34_, out35_ = default__.m0(d_285_v100_, default__.fm0(d_172_v12_, d_168_globalState_), d_168_globalState_)
                d_286_v101_ = out32_
                d_287_v102_ = out33_
                d_288_v103_ = out34_
                d_289_v104_ = out35_
                d_290_v105_: _dafny.Map
                d_290_v105_ = _dafny.Map({d_286_v101_: d_275_v92_})
                d_291_v106_: _dafny.Map
                d_291_v106_ = _dafny.Map({d_157_v0_: d_275_v92_})
                d_292_v107_: _dafny.MultiSet
                d_292_v107_ = _dafny.MultiSet([((d_290_v105_)[d_160_v3_] if (d_160_v3_) in (d_290_v105_) else ((d_291_v106_)[d_157_v0_] if (d_157_v0_) in (d_291_v106_) else d_275_v92_)), d_275_v92_])
                index37_ = default__.safeIndex(580, (d_245_v66_).length(0))
                (d_245_v66_)[index37_] = (((d_292_v107_).intersection(d_292_v107_)).cardinality) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tprjqefl"))))
                d_275_v92_ = d_275_v92_
            elif True:
                (d_168_globalState_).f13 = (len(d_275_v92_)) < ((d_249_v69_).f18)
                d_293_v108_: _dafny.Seq
                d_293_v108_ = _dafny.SeqWithoutIsStrInference([d_167_v8_, d_167_v8_])
                d_294_v109_: D1
                d_294_v109_ = D1_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cs")), d_157_v0_, d_293_v108_, default__.fm5((d_161_v4_).cardinality, ((d_274_v91_)[default__.safeIndex(747, (d_274_v91_).length(0))])[default__.safeIndex(default__.fm0(d_172_v12_, d_168_globalState_), len((d_274_v91_)[default__.safeIndex(747, (d_274_v91_).length(0))]))], (d_249_v69_).f18, (d_242_v63_).f18, d_168_globalState_), d_271_v88_)
                index38_ = default__.safeIndex(580, (d_245_v66_).length(0))
                rhs15_ = len((d_167_v8_) + ((_dafny.SeqWithoutIsStrInference([d_271_v88_ for d_295_i4_ in range(default__.abs(650))])) + (d_167_v8_)))
                rhs16_ = d_249_v69_
                rhs17_ = d_250_v70_
                rhs18_ = (d_294_v109_).cf7
                rhs19_ = (d_242_v63_).f18
                lhs14_ = d_168_globalState_
                lhs15_ = d_168_globalState_
                lhs16_ = d_245_v66_
                lhs17_ = default__.safeIndex(580, (d_245_v66_).length(0))
                lhs14_.f15 = rhs15_
                d_242_v63_ = rhs16_
                d_250_v70_ = rhs17_
                lhs15_.f13 = rhs18_
                lhs16_[lhs17_] = rhs19_
                d_296_v110_: D2
                d_296_v110_ = D2_DC6(d_249_v69_)
                d_297_v111_: _dafny.Array
                nw39_ = _dafny.Array(None, 29)
                nw39_[int(0)] = d_249_v69_
                nw39_[int(1)] = (d_249_v69_ if not(d_157_v0_) else (d_296_v110_).cf12)
                nw39_[int(2)] = d_249_v69_
                nw39_[int(3)] = d_242_v63_
                nw39_[int(4)] = d_242_v63_
                nw39_[int(5)] = d_242_v63_
                nw39_[int(6)] = d_242_v63_
                nw39_[int(7)] = d_242_v63_
                nw39_[int(8)] = d_242_v63_
                nw39_[int(9)] = d_242_v63_
                nw39_[int(10)] = d_249_v69_
                nw39_[int(11)] = (d_275_v92_)[default__.safeIndex((d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))], len(d_275_v92_))]
                nw39_[int(12)] = d_242_v63_
                nw39_[int(13)] = d_242_v63_
                nw39_[int(14)] = d_249_v69_
                nw39_[int(15)] = d_249_v69_
                nw39_[int(16)] = d_242_v63_
                nw39_[int(17)] = d_242_v63_
                nw39_[int(18)] = (d_275_v92_)[default__.safeIndex((d_249_v69_).f18, len(d_275_v92_))]
                nw39_[int(19)] = d_242_v63_
                nw39_[int(20)] = (d_275_v92_)[default__.safeIndex(917, len(d_275_v92_))]
                nw39_[int(21)] = d_249_v69_
                nw39_[int(22)] = d_249_v69_
                nw39_[int(23)] = d_242_v63_
                nw39_[int(24)] = d_249_v69_
                nw39_[int(25)] = d_242_v63_
                nw39_[int(26)] = d_242_v63_
                nw39_[int(27)] = d_249_v69_
                nw39_[int(28)] = d_249_v69_
                d_297_v111_ = nw39_
                index39_ = default__.safeIndex(892, (d_297_v111_).length(0))
                (d_297_v111_)[index39_] = d_249_v69_
                index40_ = default__.safeIndex(892, (d_297_v111_).length(0))
                (d_297_v111_)[index40_] = d_249_v69_
                d_172_v12_ = (d_172_v12_ if d_157_v0_ else d_172_v12_)
                index41_ = default__.safeIndex(51, (d_164_v7_).length(0))
                (d_164_v7_)[index41_] = not (d_157_v0_) or (d_157_v0_)
                index42_ = default__.safeIndex(51, (d_164_v7_).length(0))
                (d_164_v7_)[index42_] = d_157_v0_
            d_298_v112_: D0
            d_298_v112_ = D0_DC0((d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))], (d_171_v11_).set(d_167_v8_, (d_242_v63_).f18), _dafny.MultiSet([(0) - (d_160_v3_), (d_249_v69_).f18]))
            d_299_v113_: D0
            d_299_v113_ = D0_DC1(d_298_v112_)
            d_300_v114_: int
            d_301_v115_: _dafny.Seq
            d_302_v116_: _dafny.Set
            d_303_v117_: bool
            out36_: int
            out37_: _dafny.Seq
            out38_: _dafny.Set
            out39_: bool
            out36_, out37_, out38_, out39_ = default__.m0(D0_DC1(d_299_v113_), ((d_249_v69_).f18) - ((d_249_v69_).f18), d_168_globalState_)
            d_300_v114_ = out36_
            d_301_v115_ = out37_
            d_302_v116_ = out38_
            d_303_v117_ = out39_
            d_157_v0_ = d_157_v0_
            index43_ = default__.safeIndex(708, (d_164_v7_).length(0))
            (d_164_v7_)[index43_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqmqr"))) <= (d_167_v8_)
            index44_ = default__.safeIndex(708, (d_164_v7_).length(0))
            (d_164_v7_)[index44_] = d_303_v117_
        elif True:
            d_304_v118_: C0
            nw40_ = C0()
            nw40_.ctor__((d_242_v63_).f18)
            d_304_v118_ = nw40_
            (d_168_globalState_).f4 = (d_242_v63_).f18
            rhs20_ = (d_242_v63_).f18
            rhs21_ = d_249_v69_
            rhs22_ = not(default__.fm1(d_157_v0_, d_168_globalState_))
            lhs18_ = d_168_globalState_
            lhs19_ = d_168_globalState_
            lhs18_.f4 = rhs20_
            d_249_v69_ = rhs21_
            lhs19_.f13 = rhs22_
            d_305_v119_: _dafny.Map
            d_305_v119_ = _dafny.Map({d_271_v88_: default__.fm1(d_157_v0_, d_168_globalState_)})
            rhs23_ = not(d_157_v0_)
            rhs24_ = d_157_v0_
            rhs25_ = d_305_v119_
            lhs20_ = d_168_globalState_
            d_157_v0_ = rhs23_
            lhs20_.f13 = rhs24_
            d_305_v119_ = rhs25_
            d_306_v120_: _dafny.Map
            d_306_v120_ = _dafny.Map({(d_249_v69_).f18: _dafny.SeqWithoutIsStrInference([d_271_v88_ for d_307_i5_ in range(default__.abs(-148))])})
            d_308_v121_: D3
            d_308_v121_ = D3_DC9(d_306_v120_)
            d_309_v123_: _dafny.Array
            nw41_ = _dafny.Array(None, 7)
            nw41_[int(0)] = d_306_v120_
            nw41_[int(1)] = (_dafny.Map({696: d_167_v8_}) if d_157_v0_ else d_306_v120_)
            nw41_[int(2)] = (d_308_v121_).cf15
            def iife21_():
                coll7_ = _dafny.Map()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(905, 714):
                    d_310_v122_: int = compr_7_
                    if ((905) <= (d_310_v122_)) and ((d_310_v122_) < (714)):
                        coll7_[default__.safeDivisionInt(d_310_v122_, len(d_248_v68_))] = d_167_v8_
                return _dafny.Map(coll7_)
            nw41_[int(3)] = iife21_()
            
            nw41_[int(4)] = _dafny.Map({(d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))]: d_167_v8_})
            nw41_[int(5)] = d_306_v120_
            nw41_[int(6)] = d_306_v120_
            d_309_v123_ = nw41_
            index45_ = default__.safeIndex(856, (d_309_v123_).length(0))
            (d_309_v123_)[index45_] = d_306_v120_
            index46_ = default__.safeIndex(856, (d_309_v123_).length(0))
            (d_309_v123_)[index46_] = (_dafny.Map({(d_249_v69_).f18: d_167_v8_}) if d_157_v0_ else _dafny.Map({(0) - (len(d_167_v8_)): _dafny.SeqWithoutIsStrInference([d_271_v88_ for d_311_i6_ in range(default__.abs(-266))])}))
        d_271_v88_ = default__.fm6((d_247_v67_).ispropersubset(d_247_v67_), d_168_globalState_)
        d_170_v10_ = (d_170_v10_).set((d_157_v0_) == (not(d_157_v0_)), (d_242_v63_).f18)
        if d_157_v0_:
            index47_ = default__.safeIndex(580, (d_245_v66_).length(0))
            (d_245_v66_)[index47_] = default__.safeModuloInt(((d_249_v69_).f18 if d_157_v0_ else 706), (d_242_v63_).f18)
            rhs26_ = (d_242_v63_).f18
            rhs27_ = default__.fm0(d_172_v12_, d_168_globalState_)
            lhs21_ = d_168_globalState_
            lhs22_ = d_168_globalState_
            lhs21_.f4 = rhs26_
            lhs22_.f4 = rhs27_
            d_312_v124_: C0
            nw42_ = C0()
            nw42_.ctor__(len((_dafny.SeqWithoutIsStrInference([d_248_v68_ for d_313_i7_ in range(default__.abs(51))])).set(default__.safeIndex((d_245_v66_)[default__.safeIndex(580, (d_245_v66_).length(0))], len(_dafny.SeqWithoutIsStrInference([d_248_v68_ for d_314_i7_ in range(default__.abs(51))]))), d_248_v68_)))
            d_312_v124_ = nw42_
            d_171_v11_ = (d_171_v11_).set(d_167_v8_, 759)
            (d_168_globalState_).f4 = default__.safeModuloInt(default__.safeDivisionInt((d_249_v69_).f18, (d_163_v6_)[default__.safeIndex((d_312_v124_).f18, len(d_163_v6_))]), default__.fm0(d_172_v12_, d_168_globalState_))
        elif True:
            d_315_v125_: _dafny.Map
            d_315_v125_ = _dafny.Map({(default__.fm7(len(_dafny.SeqWithoutIsStrInference([((d_250_v70_)[(d_249_v69_).f18] if ((d_249_v69_).f18) in (d_250_v70_) else d_157_v0_), not(d_157_v0_), d_157_v0_])), d_248_v68_, d_157_v0_, d_168_globalState_)).cardinality: d_167_v8_})
            d_316_v126_: _dafny.Array
            nw43_ = _dafny.Array(D2.default()(), 19)
            d_316_v126_ = nw43_
            d_317_v127_: _dafny.Map
            d_317_v127_ = _dafny.Map({d_315_v125_: d_316_v126_})
            d_317_v127_ = (d_317_v127_).set(d_315_v125_, d_316_v126_)
            d_318_v128_: C0
            nw44_ = C0()
            nw44_.ctor__((868) + (823))
            d_318_v128_ = nw44_
            (d_168_globalState_).f13 = (d_157_v0_ if True else not(d_157_v0_))
            d_319_v129_: _dafny.Map
            d_319_v129_ = _dafny.Map({d_157_v0_: d_245_v66_})
            index48_ = default__.safeIndex(580, (d_245_v66_).length(0))
            (d_245_v66_)[index48_] = (((0) - (d_160_v3_)) + ((d_242_v63_).f18)) - (len(d_319_v129_))
            (d_168_globalState_).f0 = (d_160_v3_) + ((326) + ((d_161_v4_).cardinality))
        _dafny.print(_dafny.string_of(d_157_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v1_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_161_v4_) == (_dafny.MultiSet([639, -408]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v5_) == (_dafny.MultiSet([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v6_) == (_dafny.SeqWithoutIsStrInference([2, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v7_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_167_v8_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_168_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f1) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_168_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f8) == (_dafny.SeqWithoutIsStrInference([2, 639, 2, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f9) == (_dafny.MultiSet([639, -408]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_168_globalState_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_168_globalState_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_168_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_168_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_168_globalState_).f16)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_168_globalState_).f17).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_169_v9_) == (_dafny.Map({639: 639, 2: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v10_) == (_dafny.Map({True: 639, False: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v11_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgx")): 759}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v12_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v12_).cf1) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgx")): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v12_).cf2) == (_dafny.MultiSet([639, -408]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v62_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([2, 2]), _dafny.MultiSet([639, -408]), _dafny.MultiSet([639, -408]), _dafny.MultiSet([639, -408])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v63_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_243_v64_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_244_v65_).cf4) == (_dafny.Map({False: 639}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v66_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v67_) == (_dafny.Set({2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v68_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_249_v69_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_250_v70_) == (_dafny.Map({2: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_v88_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v89_) == (_dafny.Map({_dafny.MultiSet([639, 639, 639, 0, -242]): _dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_273_v90_) == (_dafny.Map({True: _dafny.Map({_dafny.MultiSet([639, 639, 639, 0, -242]): _dafny.CodePoint('j')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_274_v91_)[9]) == (_dafny.SeqWithoutIsStrInference([2, 2, 2, 2]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), _dafny.Map({}), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC1(D0, NamedTuple('DC1', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC3(_dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({self.cf6.VerbatimString(True)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {self.cf9.VerbatimString(True)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, _dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {self.cf18.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf15 == __o.cf15
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
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC12(D4, NamedTuple('DC12', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12)
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
        self.f0: int = int(0)
        self.f4: int = int(0)
        self.f10: int = int(0)
        self.f13: bool = False
        self.f15: int = int(0)
        self._f1: _dafny.Set = _dafny.Set({})
        self._f2: bool = False
        self._f3: bool = False
        self._f5: int = int(0)
        self._f6: bool = False
        self._f7: _dafny.Array = _dafny.Array(None, 0)
        self._f8: _dafny.Seq = _dafny.Seq({})
        self._f9: _dafny.MultiSet = _dafny.MultiSet({})
        self._f11: int = int(0)
        self._f12: str = _dafny.CodePoint('D')
        self._f14: int = int(0)
        self._f16: _dafny.Array = _dafny.Array(None, 0)
        self._f17: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self)._f16 = f16
        (self)._f17 = f17

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
    def f5(self):
        return self._f5
    @property
    def f6(self):
        return self._f6
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
    def f14(self):
        return self._f14
    @property
    def f16(self):
        return self._f16
    @property
    def f17(self):
        return self._f17

class C0:
    def  __init__(self):
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f18):
        (self)._f18 = f18

    @property
    def f18(self):
        return self._f18
