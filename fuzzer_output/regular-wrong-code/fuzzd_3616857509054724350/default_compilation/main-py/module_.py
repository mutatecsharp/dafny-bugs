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
        def iife0_():
            def iife2_():
                coll2_ = _dafny.Set()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(315, 987):
                    d_2_v1_: int = compr_2_
                    if ((315) <= (d_2_v1_)) and ((d_2_v1_) < (987)):
                        coll2_ = coll2_.union(_dafny.Set([(d_2_v1_) + ((0) - (len(_dafny.Set({False}))))]))
                return _dafny.Set(coll2_)
            coll0_ = _dafny.Map()
            def iife1_():
                coll1_ = _dafny.Set()
                compr_1_: int
                for compr_1_ in _dafny.IntegerRange(315, 987):
                    d_0_v1_: int = compr_1_
                    if ((315) <= (d_0_v1_)) and ((d_0_v1_) < (987)):
                        coll1_ = coll1_.union(_dafny.Set([(d_0_v1_) + ((0) - (len(_dafny.Set({False}))))]))
                return _dafny.Set(coll1_)
            compr_0_: int
            for compr_0_ in (iife1_()
            ).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (iife2_()
                ):
                    def iife3_():
                        coll3_ = _dafny.Map()
                        compr_3_: int
                        for compr_3_ in (_dafny.Map({865: True})).keys.Elements:
                            d_3_v2_: int = compr_3_
                            if (d_3_v2_) in (_dafny.Map({865: True})):
                                coll3_[(d_3_v2_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eukgmrlmj"))))] = True
                        return _dafny.Map(coll3_)
                    coll0_[(d_1_v0_) + (171)] = iife3_()

            return _dafny.Map(coll0_)
        return (0) - ((len(iife0_()
        )) + ((_dafny.MultiSet([_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(False): True}))])), 286, 258, len(_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pfru")): False})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kabkgac")))]), _dafny.MultiSet([-446, 261, len(_dafny.SeqWithoutIsStrInference([452, 889, (_dafny.MultiSet([True])).cardinality]))]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iy"))), -558])])).cardinality))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ej"))})).Elements:
                d_4_v0_: _dafny.Seq = compr_4_
                if (d_4_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhc")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ej"))})):
                    coll4_[d_4_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tlx"))
            return _dafny.Map(coll4_)
        return (iife4_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtjwu")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aykia"))}))

    @staticmethod
    def fm9(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([D2_DC3(False) for d_5_i0_ in range(default__.abs(878))])) + (_dafny.SeqWithoutIsStrInference([D2_DC3(not(True))]))

    @staticmethod
    def fm10(p0, p1, globalState):
        return _dafny.Map({(len(_dafny.Map({not(True): 23}))) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aeuryux")))): 212})

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        if (-998) <= (765):
            return D2_DC3(False)
        elif True:
            return D2_DC3(not(True))

    @staticmethod
    def fm12(p0, globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdliekdl"))) for d_6_i0_ in range(default__.abs(988))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: _dafny.CodePoint('d')})) for d_7_i1_ in range(default__.abs(108))]))])

    @staticmethod
    def fm13(p0, globalState):
        return _dafny.CodePoint('v')

    @staticmethod
    def fm14(p0, p1, p2, globalState):
        if (_dafny.Set({True, True, False})).ispropersubset(_dafny.Set({True})):
            return _dafny.Map({False: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_8_i0_ in range(default__.abs(634))])})
        elif True:
            return (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))})) | (_dafny.Map({not(not(True)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkbipqqh"))}))

    @staticmethod
    def fm15(p0, p1, p2, p3, globalState):
        source0_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([319, 101])) for d_9_i0_ in range(default__.abs(187))]))
        d_10___mcc_h0_ = source0_
        d_11_cf14_ = d_10___mcc_h0_
        return _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_12_i1_ in range(default__.abs(317))]): 665})

    @staticmethod
    def fm16(p0, p1, p2, p3, globalState):
        if (471) <= (-851):
            return _dafny.SeqWithoutIsStrInference([937, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xuck")))])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False])), len(_dafny.Map({-132: len(_dafny.Map({_dafny.CodePoint('a'): True}))}))])) + (_dafny.SeqWithoutIsStrInference([613 for d_13_i0_ in range(default__.abs(229))]))

    @staticmethod
    def fm19(p0, p1, globalState):
        return ((_dafny.MultiSet([False])) - (_dafny.MultiSet([not(True)]))) - ((D6_DC13(_dafny.MultiSet([True]))).cf24)

    @staticmethod
    def fm20(p0, globalState):
        def iife5_():
            def iife7_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(-44, 498):
                    d_16_v0_: int = compr_7_
                    if ((-44) <= (d_16_v0_)) and ((d_16_v0_) < (498)):
                        coll7_ = coll7_.union(_dafny.Set([default__.safeDivisionInt(d_16_v0_, len(_dafny.Map({True: 884})))]))
                return _dafny.Set(coll7_)
            coll5_ = _dafny.Set()
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(-44, 498):
                    d_14_v0_: int = compr_6_
                    if ((-44) <= (d_14_v0_)) and ((d_14_v0_) < (498)):
                        coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_14_v0_, len(_dafny.Map({True: 884})))]))
                return _dafny.Set(coll6_)
            compr_5_: int
            for compr_5_ in (iife6_()
            ).Elements:
                d_15_v1_: int = compr_5_
                if (d_15_v1_) in (iife7_()
                ):
                    coll5_ = coll5_.union(_dafny.Set([default__.safeDivisionInt(d_15_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dtejhablw"))))]))
            return _dafny.Set(coll5_)
        return ((_dafny.Set({259})) | (_dafny.Set({-865, 677, len(_dafny.Map({658: not(not(False))})), 778, 557}))).intersection(iife5_()
        )

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwb"))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return True

    @staticmethod
    def fm23(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(True), False, True])) + (_dafny.SeqWithoutIsStrInference([True]))

    @staticmethod
    def fm24(globalState):
        def iife8_():
            coll8_ = _dafny.Set()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(123, -43):
                d_19_v0_: int = compr_8_
                if ((123) <= (d_19_v0_)) and ((d_19_v0_) < (-43)):
                    coll8_ = coll8_.union(_dafny.Set([(d_19_v0_) - (675)]))
            return _dafny.Set(coll8_)
        return _dafny.Set({(True if False else not(False)), (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_17_i0_ in range(default__.abs(198))])) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_18_i1_ in range(default__.abs(598))])), (iife8_()
        ) == (_dafny.Set({-714}))})

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: int = int(0)
        d_20_v0_: _dafny.Seq
        d_20_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upn"))
        d_21_v1_: C4
        nw0_ = C4()
        nw0_.ctor__(d_20_v0_)
        d_21_v1_ = nw0_
        d_22_v2_: _dafny.MultiSet
        d_22_v2_ = _dafny.MultiSet([d_21_v1_, d_21_v1_])
        (globalState).f5 = default__.safeModuloInt((d_22_v2_).cardinality, 73)
        d_23_v3_: _dafny.Array
        def lambda0_(d_24_p0_):
            def lambda1_(d_25_i0_):
                return (_dafny.SeqWithoutIsStrInference([222])) == (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({d_24_p0_: d_24_p0_})))]))

            return lambda1_

        init0_ = lambda0_(p0)
        nw1_ = _dafny.Array(None, 10)
        for i0_0_ in range(nw1_.length(0)):
            nw1_[i0_0_] = init0_(i0_0_)
        d_23_v3_ = nw1_
        index0_ = default__.safeIndex(410, (d_23_v3_).length(0))
        (d_23_v3_)[index0_] = (d_21_v1_).fm1(p2, p0, globalState)
        d_26_v4_: _dafny.Array
        def lambda2_(d_27_i1_):
            return _dafny.Set({122})

        init1_ = lambda2_
        nw2_ = _dafny.Array(None, 8)
        for i0_1_ in range(nw2_.length(0)):
            nw2_[i0_1_] = init1_(i0_1_)
        d_26_v4_ = nw2_
        d_28_v5_: _dafny.Set
        d_28_v5_ = _dafny.Set({-565})
        index1_ = default__.safeIndex(72, (d_26_v4_).length(0))
        (d_26_v4_)[index1_] = d_28_v5_
        d_29_v6_: _dafny.Array
        def lambda3_(d_30_i2_):
            return default__.safeDivisionInt(d_30_i2_, 842)

        init2_ = lambda3_
        nw3_ = _dafny.Array(None, 14)
        for i0_2_ in range(nw3_.length(0)):
            nw3_[i0_2_] = init2_(i0_2_)
        d_29_v6_ = nw3_
        d_31_v7_: int
        d_31_v7_ = -144
        index2_ = default__.safeIndex(237, (d_29_v6_).length(0))
        (d_29_v6_)[index2_] = d_31_v7_
        d_32_v8_: bool
        d_32_v8_ = False
        d_33_v9_: _dafny.MultiSet
        d_33_v9_ = _dafny.MultiSet([d_32_v8_, p0, p2])
        d_34_v10_: _dafny.MultiSet
        d_34_v10_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([707 for d_35_i3_ in range(default__.abs(107))])])
        d_36_v11_: D5
        d_36_v11_ = D5_DC10(d_34_v10_)
        d_37_v12_: _dafny.Array
        nw4_ = _dafny.Array(None, 18)
        d_37_v12_ = nw4_
        d_38_v13_: _dafny.Seq
        d_38_v13_ = _dafny.SeqWithoutIsStrInference([d_37_v12_])
        d_39_v14_: _dafny.Map
        d_39_v14_ = _dafny.Map({d_36_v11_: len(d_38_v13_)})
        index3_ = default__.safeIndex(410, (d_23_v3_).length(0))
        index4_ = default__.safeIndex(72, (d_26_v4_).length(0))
        index5_ = default__.safeIndex(237, (d_29_v6_).length(0))
        rhs0_ = p2
        rhs1_ = default__.safeDivisionInt(((0) - (d_31_v7_)) * (d_31_v7_), d_31_v7_)
        rhs2_ = d_28_v5_
        rhs3_ = (((d_33_v9_)[d_32_v8_] if (d_32_v8_) in (d_33_v9_) else len(d_39_v14_))) + (d_31_v7_)
        rhs4_ = p2
        lhs0_ = d_23_v3_
        lhs1_ = default__.safeIndex(410, (d_23_v3_).length(0))
        lhs2_ = globalState
        lhs3_ = d_26_v4_
        lhs4_ = default__.safeIndex(72, (d_26_v4_).length(0))
        lhs5_ = d_29_v6_
        lhs6_ = default__.safeIndex(237, (d_29_v6_).length(0))
        lhs0_[lhs1_] = rhs0_
        lhs2_.f5 = rhs1_
        lhs3_[lhs4_] = rhs2_
        lhs5_[lhs6_] = rhs3_
        d_32_v8_ = rhs4_
        if (d_21_v1_).fm1(p2, not ((d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]) or (d_32_v8_), globalState):
            d_40_v15_: _dafny.Map
            d_40_v15_ = _dafny.Map({(p0 if p2 else (d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]): d_32_v8_})
            d_40_v15_ = (d_40_v15_).set(p2, p2)
            (globalState).f5 = (d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))]
            index6_ = default__.safeIndex(410, (d_23_v3_).length(0))
            (d_23_v3_)[index6_] = d_32_v8_
            r0 = (d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))]
            (globalState).f5 = (d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))]
        elif True:
            d_41_v16_: _dafny.Set
            d_41_v16_ = _dafny.Set({p0, d_32_v8_, (d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))], (d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]})
            d_42_v17_: _dafny.Map
            d_42_v17_ = _dafny.Map({(d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]: d_41_v16_})
            d_41_v16_ = (default__.fm24(globalState)).intersection(((d_42_v17_)[p2] if (p2) in (d_42_v17_) else d_41_v16_))
            index7_ = default__.safeIndex(410, (d_23_v3_).length(0))
            (d_23_v3_)[index7_] = p0
            d_43_v18_: _dafny.Array
            nw5_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_43_v18_ = nw5_
            index8_ = default__.safeIndex(491, (d_43_v18_).length(0))
            (d_43_v18_)[index8_] = d_23_v3_
            index9_ = default__.safeIndex(491, (d_43_v18_).length(0))
            (d_43_v18_)[index9_] = d_23_v3_
            if not ((d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]) or ((d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]):
                d_44_v19_: _dafny.Array
                def lambda4_(d_45_v7_, d_46_v3_, d_47_p2_):
                    def lambda5_(d_48_i4_):
                        return _dafny.Map({len(_dafny.Map({d_45_v7_: (d_46_v3_)[default__.safeIndex(410, (d_46_v3_).length(0))]})): d_47_p2_})

                    return lambda5_

                init3_ = lambda4_(d_31_v7_, d_23_v3_, p2)
                nw6_ = _dafny.Array(None, 13)
                for i0_3_ in range(nw6_.length(0)):
                    nw6_[i0_3_] = init3_(i0_3_)
                d_44_v19_ = nw6_
                d_49_v20_: _dafny.Map
                d_49_v20_ = _dafny.Map({(d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))]: p0})
                index10_ = default__.safeIndex(726, (d_44_v19_).length(0))
                (d_44_v19_)[index10_] = d_49_v20_
                index11_ = default__.safeIndex(726, (d_44_v19_).length(0))
                (d_44_v19_)[index11_] = (((d_49_v20_).set((d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))], p2)).set((d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))], not((d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]))) | (d_49_v20_)
                d_50_v21_: str
                d_50_v21_ = _dafny.CodePoint('r')
                d_51_v22_: C0
                nw7_ = C0()
                nw7_.ctor__((d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))], d_50_v21_, p0, (d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))])
                d_51_v22_ = nw7_
                d_52_v23_: _dafny.Seq
                d_52_v23_ = _dafny.SeqWithoutIsStrInference([d_51_v22_, d_51_v22_, d_51_v22_, d_51_v22_])
                d_53_v24_: _dafny.Map
                d_53_v24_ = _dafny.Map({(d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))]: p0})
                d_54_v25_: _dafny.Seq
                d_54_v25_ = _dafny.SeqWithoutIsStrInference([(d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))], (d_51_v22_) not in (d_52_v23_), ((d_53_v24_)[p2] if (p2) in (d_53_v24_) else (d_51_v22_).fm5(d_32_v8_, globalState))])
                d_54_v25_ = d_54_v25_
                d_55_v27_: D5
                d_55_v27_ = D5_DC11()
                d_56_v28_: _dafny.MultiSet
                d_56_v28_ = _dafny.MultiSet([d_55_v27_])
                d_57_v29_: _dafny.Map
                d_57_v29_ = _dafny.Map({p0: d_56_v28_})
                index12_ = default__.safeIndex(410, (d_23_v3_).length(0))
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in ((_dafny.SeqWithoutIsStrInference([d_51_v22_.f13 for d_58_i5_ in range(default__.abs(946))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dy"))), len(_dafny.SeqWithoutIsStrInference([d_51_v22_.f13 for d_59_i5_ in range(default__.abs(946))]))), d_51_v22_.f13)).Elements:
                        d_60_v26_: int = compr_9_
                        if (d_60_v26_) in ((_dafny.SeqWithoutIsStrInference([d_51_v22_.f13 for d_58_i5_ in range(default__.abs(946))])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dy"))), len(_dafny.SeqWithoutIsStrInference([d_51_v22_.f13 for d_59_i5_ in range(default__.abs(946))]))), d_51_v22_.f13)):
                            coll9_[(d_60_v26_) * (d_31_v7_)] = _dafny.MultiSet([D5_DC11()])
                    return _dafny.Map(coll9_)
                rhs5_ = len(((iife9_()
                ).set(832, ((d_57_v29_)[d_32_v8_] if (d_32_v8_) in (d_57_v29_) else d_56_v28_))).set((0) - ((d_51_v22_.f13) * ((0) - (d_31_v7_))), d_56_v28_))
                rhs6_ = d_32_v8_
                lhs7_ = globalState
                lhs8_ = d_23_v3_
                lhs9_ = default__.safeIndex(410, (d_23_v3_).length(0))
                lhs7_.f5 = rhs5_
                lhs8_[lhs9_] = rhs6_
                d_32_v8_ = p0
                d_61_v30_: D1
                d_61_v30_ = D1_DC1(((d_21_v1_).f7).set(default__.safeIndex(d_31_v7_, len((d_21_v1_).f7)), d_50_v21_))
                (globalState).f5 = len(_dafny.SeqWithoutIsStrInference([d_61_v30_, d_61_v30_]))
            elif True:
                d_20_v0_ = (d_21_v1_).f7
                d_20_v0_ = (d_21_v1_).f7
                d_62_v31_: _dafny.Array
                def lambda6_(d_63_v1_):
                    def lambda7_(d_64_i6_):
                        return ((d_63_v1_).f7) + ((d_63_v1_).f7)

                    return lambda7_

                init4_ = lambda6_(d_21_v1_)
                nw8_ = _dafny.Array(None, 1)
                for i0_4_ in range(nw8_.length(0)):
                    nw8_[i0_4_] = init4_(i0_4_)
                d_62_v31_ = nw8_
                d_65_v32_: _dafny.Seq
                d_65_v32_ = _dafny.SeqWithoutIsStrInference([d_62_v31_, d_62_v31_, d_62_v31_, d_62_v31_])
                d_62_v31_ = (d_65_v32_)[default__.safeIndex(d_31_v7_, len(d_65_v32_))]
                index13_ = default__.safeIndex(410, (d_23_v3_).length(0))
                (d_23_v3_)[index13_] = p2
                index14_ = default__.safeIndex(410, (d_23_v3_).length(0))
                (d_23_v3_)[index14_] = d_32_v8_
            (globalState).f5 = d_31_v7_
        index15_ = default__.safeIndex(410, (d_23_v3_).length(0))
        (d_23_v3_)[index15_] = (d_31_v7_) > (d_31_v7_)
        d_66_i7_: int
        d_66_i7_ = 0
        with _dafny.label("0"):
            while p2:
                with _dafny.c_label("0"):
                    if (d_66_i7_) >= (100):
                        raise _dafny.Break("0")
                    d_66_i7_ = (d_66_i7_) + (1)
                    index16_ = default__.safeIndex(237, (d_29_v6_).length(0))
                    (d_29_v6_)[index16_] = default__.fm0((p0 if p0 else p0), (d_21_v1_).f7, globalState)
                    d_67_v33_: str
                    d_67_v33_ = _dafny.CodePoint('r')
                    d_68_v34_: _dafny.Map
                    d_68_v34_ = _dafny.Map({default__.fm23(p0, d_32_v8_, d_31_v7_, globalState): d_67_v33_})
                    index17_ = default__.safeIndex(237, (d_29_v6_).length(0))
                    (d_29_v6_)[index17_] = (len(d_68_v34_)) - ((d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))])
                    (globalState).f5 = default__.safeModuloInt((d_31_v7_) + ((d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))]), (d_29_v6_)[default__.safeIndex(237, (d_29_v6_).length(0))])
                    d_69_v35_: C3
                    nw9_ = C3()
                    nw9_.ctor__()
                    d_69_v35_ = nw9_
                    pass
            pass
        d_70_v36_: _dafny.Seq
        d_70_v36_ = _dafny.SeqWithoutIsStrInference([d_23_v3_, d_23_v3_])
        d_71_v37_: str
        d_71_v37_ = _dafny.CodePoint('e')
        d_72_v38_: _dafny.Seq
        d_72_v38_ = _dafny.SeqWithoutIsStrInference([(d_21_v1_).f7, d_20_v0_, d_20_v0_, ((d_21_v1_).f7).set(default__.safeIndex(d_31_v7_, len((d_21_v1_).f7)), d_71_v37_), (d_21_v1_).f7])
        d_73_v39_: D4
        d_73_v39_ = D4_DC9(p0, d_32_v8_)
        d_74_v40_: _dafny.Seq
        d_74_v40_ = _dafny.SeqWithoutIsStrInference([(d_70_v36_)[default__.safeIndex(len(default__.fm16(len(d_72_v38_), (d_23_v3_)[default__.safeIndex(410, (d_23_v3_).length(0))], (d_73_v39_).cf21, True, globalState)), len(d_70_v36_))], d_23_v3_, d_23_v3_])
        d_74_v40_ = d_74_v40_
        r0 = d_31_v7_
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_75_v0_: int
        d_75_v0_ = 717
        d_76_v1_: _dafny.Seq
        d_76_v1_ = _dafny.SeqWithoutIsStrInference([d_75_v0_])
        d_77_v2_: _dafny.Map
        d_77_v2_ = _dafny.Map({d_75_v0_: d_76_v1_})
        d_78_v3_: _dafny.Array
        nw10_ = _dafny.Array(False, 10)
        d_78_v3_ = nw10_
        d_79_globalState_: GlobalState
        nw11_ = GlobalState()
        nw11_.ctor__(d_77_v2_, False, False, d_78_v3_, True, 341, False)
        d_79_globalState_ = nw11_
        d_80_v4_: _dafny.Seq
        d_80_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbqusrq"))
        d_80_v4_ = d_80_v4_
        d_81_v5_: _dafny.Map
        d_81_v5_ = _dafny.Map({d_80_v4_: d_80_v4_})
        d_82_v6_: _dafny.Array
        def lambda8_(d_83_v0_):
            def lambda9_(d_84_i1_):
                return default__.safeDivisionInt(d_84_i1_, d_83_v0_)

            return lambda9_

        init5_ = lambda8_(d_75_v0_)
        nw12_ = _dafny.Array(None, 11)
        for i0_5_ in range(nw12_.length(0)):
            nw12_[i0_5_] = init5_(i0_5_)
        d_82_v6_ = nw12_
        d_85_v7_: bool
        d_85_v7_ = True
        d_86_v8_: _dafny.Map
        d_86_v8_ = _dafny.Map({d_85_v7_: d_75_v0_})
        d_87_v9_: _dafny.Array
        def lambda10_(d_88_i2_):
            return _dafny.CodePoint('t')

        init6_ = lambda10_
        nw13_ = _dafny.Array(None, 22)
        for i0_6_ in range(nw13_.length(0)):
            nw13_[i0_6_] = init6_(i0_6_)
        d_87_v9_ = nw13_
        d_89_v10_: D0
        d_89_v10_ = D0_DC0(_dafny.Map({d_80_v4_: d_80_v4_}), d_82_v6_, d_86_v8_, d_87_v9_, d_82_v6_)
        d_90_v11_: _dafny.Map
        d_90_v11_ = _dafny.Map({d_85_v7_: (0) - (len(_dafny.Set({D0_DC0(d_81_v5_, d_82_v6_, d_86_v8_, d_87_v9_, d_82_v6_), d_89_v10_})))})
        d_91_v12_: D0
        d_91_v12_ = D0_DC0((d_81_v5_).set(d_80_v4_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_92_i0_ in range(default__.abs(-417))])), d_82_v6_, (d_86_v8_ if d_85_v7_ else d_90_v11_), d_87_v9_, d_82_v6_)
        source1_ = d_91_v12_
        d_93___mcc_h0_ = source1_.cf0
        d_94___mcc_h1_ = source1_.cf1
        d_95___mcc_h2_ = source1_.cf2
        d_96___mcc_h3_ = source1_.cf3
        d_97___mcc_h4_ = source1_.cf4
        d_98_cf4_ = d_97___mcc_h4_
        d_99_cf3_ = d_96___mcc_h3_
        d_100_cf2_ = d_95___mcc_h2_
        d_101_cf1_ = d_94___mcc_h1_
        d_102_cf0_ = d_93___mcc_h0_
        d_103_v13_: _dafny.Set
        d_103_v13_ = _dafny.Set({d_100_cf2_, d_90_v11_})
        d_104_v14_: int
        out0_: int
        out0_ = default__.m0(d_85_v7_, d_103_v13_, (d_75_v0_) == (d_75_v0_), d_79_globalState_)
        d_104_v14_ = out0_
        d_80_v4_ = d_80_v4_
        d_105_v15_: _dafny.MultiSet
        d_105_v15_ = _dafny.MultiSet([not(d_85_v7_)])
        source2_ = D0_DC0(d_102_cf0_, d_82_v6_, _dafny.Map({False: (d_105_v15_).cardinality}), d_99_cf3_, d_98_cf4_)
        d_106___mcc_h5_ = source2_.cf0
        d_107___mcc_h6_ = source2_.cf1
        d_108___mcc_h7_ = source2_.cf2
        d_109___mcc_h8_ = source2_.cf3
        d_110___mcc_h9_ = source2_.cf4
        d_111_cf4_ = d_110___mcc_h9_
        d_112_cf3_ = d_109___mcc_h8_
        d_113_cf2_ = d_108___mcc_h7_
        d_114_cf1_ = d_107___mcc_h6_
        d_115_cf0_ = d_106___mcc_h5_
        index18_ = default__.safeIndex(597, (d_114_cf1_).length(0))
        (d_114_cf1_)[index18_] = d_104_v14_
        index19_ = default__.safeIndex(597, (d_114_cf1_).length(0))
        (d_114_cf1_)[index19_] = ((len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvgx")) for d_116_i3_ in range(default__.abs(687))]))) - (default__.fm0(d_85_v7_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_117_i4_ in range(default__.abs(-139))]), d_79_globalState_))) * (d_75_v0_)
        index20_ = default__.safeIndex(597, (d_114_cf1_).length(0))
        (d_114_cf1_)[index20_] = (d_114_cf1_)[default__.safeIndex(597, (d_114_cf1_).length(0))]
        d_118_v16_: int
        out1_: int
        out1_ = default__.m0(not(d_85_v7_), d_103_v13_, (d_85_v7_) or (d_85_v7_), d_79_globalState_)
        d_118_v16_ = out1_
        index21_ = default__.safeIndex(597, (d_114_cf1_).length(0))
        (d_114_cf1_)[index21_] = d_75_v0_
        d_119_v17_: _dafny.Map
        d_119_v17_ = _dafny.Map({d_100_cf2_: len(d_80_v4_)})
        d_120_v19_: int
        out2_: int
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: _dafny.Map
            for compr_10_ in (d_119_v17_).keys.Elements:
                d_121_v18_: _dafny.Map = compr_10_
                if (d_121_v18_) in (d_119_v17_):
                    coll10_ = coll10_.union(_dafny.Set([d_121_v18_]))
            return _dafny.Set(coll10_)
        out2_ = default__.m0(d_85_v7_, (d_103_v13_ if d_85_v7_ else iife10_()
        ), not(d_85_v7_), d_79_globalState_)
        d_120_v19_ = out2_
        d_85_v7_ = d_85_v7_
        d_75_v0_ = len((d_80_v4_) + (d_80_v4_))
        if d_85_v7_:
            d_85_v7_ = d_85_v7_
            d_122_v20_: _dafny.Set
            d_122_v20_ = _dafny.Set({d_90_v11_})
            d_123_v21_: int
            out3_: int
            out3_ = default__.m0(not (d_85_v7_) or (d_85_v7_), d_122_v20_, d_85_v7_, d_79_globalState_)
            d_123_v21_ = out3_
            d_124_v22_: str
            d_124_v22_ = _dafny.CodePoint('l')
            d_125_v23_: _dafny.Set
            d_125_v23_ = _dafny.Set({d_124_v22_, d_124_v22_})
            d_126_v24_: _dafny.MultiSet
            d_126_v24_ = _dafny.MultiSet([d_76_v1_])
            d_127_v25_: C1
            nw14_ = C1()
            nw14_.ctor__(d_126_v24_, d_85_v7_, d_124_v22_, d_85_v7_, d_85_v7_)
            d_127_v25_ = nw14_
            d_128_v26_: D4
            d_128_v26_ = D4_DC8(len(d_125_v23_), not(d_85_v7_), d_123_v21_, d_127_v25_)
            d_129_v27_: _dafny.Map
            d_129_v27_ = _dafny.Map({d_75_v0_: d_85_v7_})
            d_130_v28_: D1
            d_130_v28_ = D1_DC2(d_75_v0_, len(d_129_v27_), d_124_v22_)
            d_131_v29_: C0
            nw15_ = C0()
            nw15_.ctor__(default__.fm0((d_128_v26_).cf17, d_80_v4_, d_79_globalState_), (d_130_v28_).cf8, (d_127_v25_).f12, (d_127_v25_).f12)
            d_131_v29_ = nw15_
            d_132_v30_: _dafny.Array
            nw16_ = _dafny.Array(D2.default()(), 22)
            d_132_v30_ = nw16_
            d_133_v31_: _dafny.Seq
            d_133_v31_ = _dafny.SeqWithoutIsStrInference([(d_127_v25_).f12])
            d_134_v32_: T1
            nw17_ = C1()
            nw17_.ctor__((d_127_v25_).f11, d_85_v7_, d_124_v22_, (d_133_v31_)[default__.safeIndex(d_75_v0_, len(d_133_v31_))], d_85_v7_)
            d_134_v32_ = nw17_
            d_135_v33_: _dafny.Seq
            d_135_v33_ = _dafny.SeqWithoutIsStrInference([d_134_v32_])
            d_136_v34_: _dafny.Map
            d_136_v34_ = _dafny.Map({d_135_v33_: d_131_v29_.f13})
            d_137_v35_: _dafny.Map
            d_137_v35_ = _dafny.Map({441: d_75_v0_})
            d_138_v36_: _dafny.Set
            d_138_v36_ = _dafny.Set({d_134_v32_.f9, (d_127_v25_).f12})
            d_139_v37_: D2
            d_139_v37_ = D2_DC4(d_80_v4_, d_133_v31_, ((d_136_v34_)[d_135_v33_] if (d_135_v33_) in (d_136_v34_) else ((d_137_v35_)[d_131_v29_.f13] if (d_131_v29_.f13) in (d_137_v35_) else len(d_138_v36_))))
            d_140_v38_: D2
            d_140_v38_ = D2_DC5(d_139_v37_)
            d_141_v39_: D2
            d_141_v39_ = D2_DC5(d_139_v37_)
            index22_ = default__.safeIndex(32, (d_132_v30_).length(0))
            (d_132_v30_)[index22_] = d_141_v39_
            index23_ = default__.safeIndex(32, (d_132_v30_).length(0))
            (d_132_v30_)[index23_] = d_141_v39_
            d_142_v40_: _dafny.Map
            d_142_v40_ = _dafny.Map({d_85_v7_: (d_127_v25_).f12})
            d_143_v41_: C2
            nw18_ = C2()
            nw18_.ctor__(len(d_142_v40_), not((d_127_v25_).f12), d_134_v32_.f9)
            d_143_v41_ = nw18_
            d_144_v42_: _dafny.Array
            nw19_ = _dafny.Array(_dafny.Set({}), 22)
            d_144_v42_ = nw19_
            rhs7_ = d_80_v4_
            rhs8_ = (d_127_v25_).f12
            rhs9_ = default__.fm21((d_123_v21_ if (d_127_v25_).f12 else -1), (d_80_v4_) + (d_80_v4_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_145_i5_ in range(default__.abs(109))]), (d_143_v41_).f14, d_79_globalState_)
            rhs10_ = (d_143_v41_ if True else d_143_v41_)
            rhs11_ = d_144_v42_
            lhs10_ = d_134_v32_
            d_80_v4_ = rhs7_
            lhs10_.f9 = rhs8_
            d_80_v4_ = rhs9_
            d_143_v41_ = rhs10_
            d_144_v42_ = rhs11_
        elif True:
            d_146_v43_: _dafny.MultiSet
            d_146_v43_ = _dafny.MultiSet([d_85_v7_, d_85_v7_, d_85_v7_])
            d_147_v44_: D2
            d_147_v44_ = D2_DC3(False)
            d_148_v45_: int
            out4_: int
            out4_ = default__.m0(not((d_75_v0_) >= (d_75_v0_)), _dafny.Set({_dafny.Map({d_85_v7_: d_75_v0_})}), default__.fm22((d_146_v43_).cardinality, d_85_v7_, not(True), (d_147_v44_).cf9, d_79_globalState_), d_79_globalState_)
            d_148_v45_ = out4_
            d_149_v46_: _dafny.MultiSet
            d_149_v46_ = _dafny.MultiSet([d_76_v1_])
            d_150_v47_: str
            d_150_v47_ = _dafny.CodePoint('q')
            d_151_v48_: C1
            nw20_ = C1()
            nw20_.ctor__(d_149_v46_, d_85_v7_, d_150_v47_, d_85_v7_, d_85_v7_)
            d_151_v48_ = nw20_
            d_152_v49_: D4
            d_152_v49_ = D4_DC8(984, d_85_v7_, d_148_v45_, d_151_v48_)
            d_153_v50_: _dafny.Seq
            d_153_v50_ = _dafny.SeqWithoutIsStrInference([default__.fm22((d_152_v49_).cf18, (d_151_v48_).f12, d_85_v7_, True, d_79_globalState_), (d_151_v48_).f12, (d_151_v48_).f12, (d_151_v48_).f12, (d_151_v48_).f12])
            d_154_v51_: _dafny.Map
            d_154_v51_ = _dafny.Map({d_85_v7_: ((d_153_v50_) + (d_153_v50_)).set(default__.safeIndex(d_75_v0_, len((d_153_v50_) + (d_153_v50_))), (d_151_v48_).f12)})
            d_154_v51_ = d_154_v51_
            d_85_v7_ = False
            d_151_v48_ = (d_151_v48_ if d_85_v7_ else d_151_v48_)
            d_85_v7_ = (d_151_v48_).f12
        d_155_v52_: str
        d_155_v52_ = _dafny.CodePoint('y')
        d_155_v52_ = default__.fm13(d_75_v0_, d_79_globalState_)
        d_85_v7_ = d_85_v7_
        d_85_v7_ = d_85_v7_
        d_156_v53_: C2
        nw21_ = C2()
        nw21_.ctor__(d_75_v0_, d_85_v7_, d_85_v7_)
        d_156_v53_ = nw21_
        d_157_v54_: _dafny.Seq
        d_157_v54_ = _dafny.SeqWithoutIsStrInference([D1_DC1(d_80_v4_)])
        d_158_v55_: _dafny.Seq
        d_158_v55_ = _dafny.SeqWithoutIsStrInference([d_157_v54_, d_157_v54_])
        rhs12_ = d_85_v7_
        rhs13_ = ((d_157_v54_) + ((d_158_v55_)[default__.safeIndex((d_156_v53_).f14, len(d_158_v55_))])) + (d_157_v54_)
        rhs14_ = d_85_v7_
        d_85_v7_ = rhs12_
        d_157_v54_ = rhs13_
        d_85_v7_ = rhs14_
        d_159_v56_: _dafny.MultiSet
        d_159_v56_ = _dafny.MultiSet([d_76_v1_])
        d_160_v57_: T1
        nw22_ = C1()
        nw22_.ctor__(d_159_v56_, d_85_v7_, d_155_v52_, d_85_v7_, d_85_v7_)
        d_160_v57_ = nw22_
        d_161_v58_: _dafny.Seq
        d_161_v58_ = _dafny.SeqWithoutIsStrInference([d_160_v57_, d_160_v57_, d_160_v57_, d_160_v57_, d_160_v57_])
        d_162_v59_: _dafny.Array
        nw23_ = _dafny.Array(None, 10)
        nw23_[int(0)] = (d_160_v57_ if d_85_v7_ else d_160_v57_)
        nw23_[int(1)] = d_160_v57_
        nw23_[int(2)] = d_160_v57_
        nw23_[int(3)] = (d_160_v57_ if d_85_v7_ else d_160_v57_)
        nw23_[int(4)] = d_160_v57_
        nw23_[int(5)] = (d_161_v58_)[default__.safeIndex(len(d_80_v4_), len(d_161_v58_))]
        nw23_[int(6)] = d_160_v57_
        nw23_[int(7)] = d_160_v57_
        nw23_[int(8)] = d_160_v57_
        nw23_[int(9)] = d_160_v57_
        d_162_v59_ = nw23_
        index24_ = default__.safeIndex(979, (d_162_v59_).length(0))
        (d_162_v59_)[index24_] = d_160_v57_
        index25_ = default__.safeIndex(979, (d_162_v59_).length(0))
        (d_162_v59_)[index25_] = d_160_v57_
        d_163_v60_: _dafny.MultiSet
        d_163_v60_ = _dafny.MultiSet([d_75_v0_, d_75_v0_])
        d_164_v61_: _dafny.Set
        d_164_v61_ = _dafny.Set({(d_156_v53_).f14})
        if (((d_163_v60_).cardinality) != (d_75_v0_)) or ((d_164_v61_).isdisjoint(d_164_v61_)):
            index26_ = default__.safeIndex(927, (d_82_v6_).length(0))
            (d_82_v6_)[index26_] = (0) - (default__.safeModuloInt(d_75_v0_, (d_156_v53_).f14))
            index27_ = default__.safeIndex(927, (d_82_v6_).length(0))
            (d_82_v6_)[index27_] = default__.safeModuloInt((d_156_v53_).f14, (d_75_v0_) + (71))
            d_165_v62_: _dafny.Array
            nw24_ = _dafny.Array(int(0), 26)
            d_165_v62_ = nw24_
            d_165_v62_ = d_165_v62_
            d_85_v7_ = d_85_v7_
            (d_79_globalState_).f5 = (D1_DC2((0) - (len(_dafny.Map({d_165_v62_: d_160_v57_.f9}))), (d_82_v6_)[default__.safeIndex(927, (d_82_v6_).length(0))], d_160_v57_.f10)).cf7
            d_166_v63_: _dafny.Seq
            d_166_v63_ = _dafny.SeqWithoutIsStrInference([d_85_v7_])
            if (d_166_v63_)[default__.safeIndex(80, len(d_166_v63_))]:
                d_80_v4_ = d_80_v4_
                d_85_v7_ = (d_166_v63_)[default__.safeIndex((d_82_v6_)[default__.safeIndex(927, (d_82_v6_).length(0))], len(d_166_v63_))]
                d_167_v64_: _dafny.Array
                nw25_ = _dafny.Array(None, 3)
                d_167_v64_ = nw25_
                d_168_v65_: C1
                nw26_ = C1()
                nw26_.ctor__(d_159_v56_, not(not(d_160_v57_.f9)), d_160_v57_.f10, not((d_160_v57_).f8), d_160_v57_.f9)
                d_168_v65_ = nw26_
                index28_ = default__.safeIndex(489, (d_167_v64_).length(0))
                (d_167_v64_)[index28_] = d_168_v65_
                index29_ = default__.safeIndex(489, (d_167_v64_).length(0))
                rhs15_ = d_165_v62_
                rhs16_ = d_168_v65_
                lhs11_ = d_167_v64_
                lhs12_ = default__.safeIndex(489, (d_167_v64_).length(0))
                d_165_v62_ = rhs15_
                lhs11_[lhs12_] = rhs16_
                d_169_v66_: _dafny.Map
                d_169_v66_ = _dafny.Map({(d_80_v4_) + (d_80_v4_): len(d_166_v63_)})
                d_169_v66_ = (d_169_v66_).set(((_dafny.SeqWithoutIsStrInference([d_155_v52_ for d_170_i6_ in range(default__.abs(582))])).set(default__.safeIndex((d_156_v53_).f14, len(_dafny.SeqWithoutIsStrInference([d_155_v52_ for d_171_i6_ in range(default__.abs(582))]))), d_160_v57_.f10)) + (_dafny.SeqWithoutIsStrInference([d_160_v57_.f10 for d_172_i7_ in range(default__.abs(996))])), (len(_dafny.SeqWithoutIsStrInference([d_160_v57_.f10 for d_173_i8_ in range(default__.abs(983))]))) * ((d_156_v53_).f14))
                index30_ = default__.safeIndex(999, (d_78_v3_).length(0))
                (d_78_v3_)[index30_] = not((d_160_v57_).f8)
                d_174_v67_: _dafny.Set
                d_174_v67_ = _dafny.Set({(d_160_v57_).f8, (d_168_v65_).f12})
                index31_ = default__.safeIndex(999, (d_78_v3_).length(0))
                (d_78_v3_)[index31_] = (d_160_v57_).fm7((len(d_166_v63_)) > ((d_156_v53_).f14), d_174_v67_, d_80_v4_, d_79_globalState_)
            elif True:
                d_175_v68_: _dafny.Set
                d_175_v68_ = _dafny.Set({d_90_v11_})
                d_176_v69_: int
                out5_: int
                out5_ = default__.m0(not((d_160_v57_).fm5((d_160_v57_).f8, d_79_globalState_)), d_175_v68_, True, d_79_globalState_)
                d_176_v69_ = out5_
                (d_160_v57_).f9 = ((d_160_v57_).f8 if not((d_160_v57_).f8) else not (d_160_v57_.f9) or ((d_160_v57_).f8))
                d_177_v70_: T0
                nw27_ = C2()
                nw27_.ctor__((d_156_v53_).f14, d_160_v57_.f9, False)
                d_177_v70_ = nw27_
                d_178_v71_: _dafny.Seq
                d_178_v71_ = _dafny.SeqWithoutIsStrInference([d_177_v70_, d_177_v70_])
                (d_79_globalState_).f5 = ((d_163_v60_)[(0) - ((0) - (((d_163_v60_)[len(d_178_v71_)] if (len(d_178_v71_)) in (d_163_v60_) else (d_156_v53_).f14)))] if ((0) - ((0) - (((d_163_v60_)[len(d_178_v71_)] if (len(d_178_v71_)) in (d_163_v60_) else (d_156_v53_).f14)))) in (d_163_v60_) else default__.safeModuloInt(d_176_v69_, (d_82_v6_)[default__.safeIndex(927, (d_82_v6_).length(0))]))
                (d_160_v57_).f10 = (d_80_v4_)[default__.safeIndex(d_176_v69_, len(d_80_v4_))]
                d_164_v61_ = (default__.fm20((d_177_v70_).f8, d_79_globalState_)) | (d_164_v61_)
        elif True:
            index32_ = default__.safeIndex(298, (d_82_v6_).length(0))
            (d_82_v6_)[index32_] = (d_156_v53_).f14
            index33_ = default__.safeIndex(298, (d_82_v6_).length(0))
            (d_82_v6_)[index33_] = (d_160_v57_).fm6(d_160_v57_.f9, d_79_globalState_)
            index34_ = default__.safeIndex(680, (d_78_v3_).length(0))
            (d_78_v3_)[index34_] = d_85_v7_
            d_179_v72_: D4
            d_179_v72_ = D4_DC9((d_160_v57_).f8, d_160_v57_.f9)
            index35_ = default__.safeIndex(680, (d_78_v3_).length(0))
            (d_78_v3_)[index35_] = not((False) and ((d_179_v72_).cf21))
            d_180_v73_: _dafny.Seq
            d_180_v73_ = _dafny.SeqWithoutIsStrInference([default__.fm22((d_156_v53_).f14, (d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))], d_160_v57_.f9, d_85_v7_, d_79_globalState_)])
            if (d_163_v60_).isdisjoint(_dafny.MultiSet([len(d_180_v73_), (d_156_v53_).f14])):
                d_181_v74_: _dafny.Map
                d_181_v74_ = _dafny.Map({(d_156_v53_).f14: len((_dafny.Map({(d_160_v57_).fm6((d_160_v57_).f8, d_79_globalState_): (d_156_v53_).f14})).set((d_82_v6_)[default__.safeIndex(298, (d_82_v6_).length(0))], (d_156_v53_).f14))})
                (d_160_v57_).f9 = (default__.safeDivisionInt((d_156_v53_).f14, len(d_181_v74_))) < (((d_156_v53_).f14) + ((d_156_v53_).f14))
                d_182_v75_: C1
                nw28_ = C1()
                nw28_.ctor__((d_159_v56_) - (d_159_v56_), (d_160_v57_).f8, d_160_v57_.f10, (d_160_v57_).f8, (d_156_v53_).fm5((d_160_v57_).f8, d_79_globalState_))
                d_182_v75_ = nw28_
                d_183_v76_: D4
                d_183_v76_ = D4_DC8((d_82_v6_)[default__.safeIndex(298, (d_82_v6_).length(0))], (d_160_v57_).f8, (d_156_v53_).f14, d_182_v75_)
                pat_let_tv0_ = d_82_v6_
                pat_let_tv1_ = d_82_v6_
                pat_let_tv2_ = d_160_v57_
                nw29_ = C1()
                def iife11_(_pat_let0_0):
                    def iife12_(d_184_dt__update__tmp_h0_):
                        def iife13_(_pat_let1_0):
                            def iife14_(d_185_dt__update_hcf16_h0_):
                                def iife15_(_pat_let2_0):
                                    def iife16_(d_186_dt__update_hcf17_h0_):
                                        return D4_DC8(d_185_dt__update_hcf16_h0_, d_186_dt__update_hcf17_h0_, (d_184_dt__update__tmp_h0_).cf18, (d_184_dt__update__tmp_h0_).cf19)
                                    return iife16_(_pat_let2_0)
                                return iife15_(pat_let_tv2_.f9)
                            return iife14_(_pat_let1_0)
                        return iife13_((pat_let_tv1_)[default__.safeIndex(298, (pat_let_tv0_).length(0))])
                    return iife12_(_pat_let0_0)
                nw29_.ctor__((d_182_v75_).f11, d_85_v7_, d_155_v52_, (iife11_(d_183_v76_)).cf17, (d_160_v57_).f8)
                d_182_v75_ = nw29_
                d_163_v60_ = d_163_v60_
                d_187_v77_: _dafny.Set
                d_187_v77_ = _dafny.Set({(d_182_v75_).f12, d_160_v57_.f9, (d_160_v57_).f8, d_160_v57_.f9, not(d_160_v57_.f9)})
                d_188_v78_: _dafny.Array
                def lambda11_(d_189_v3_):
                    def lambda12_(d_190_i9_):
                        return (d_189_v3_)[default__.safeIndex(680, (d_189_v3_).length(0))]

                    return lambda12_

                init7_ = lambda11_(d_78_v3_)
                nw30_ = _dafny.Array(None, 24)
                for i0_7_ in range(nw30_.length(0)):
                    nw30_[i0_7_] = init7_(i0_7_)
                d_188_v78_ = nw30_
                d_191_v79_: _dafny.Map
                d_191_v79_ = _dafny.Map({d_80_v4_: (d_82_v6_)[default__.safeIndex(298, (d_82_v6_).length(0))]})
                index36_ = default__.safeIndex(298, (d_82_v6_).length(0))
                rhs17_ = ((d_156_v53_).f14) + ((d_163_v60_).cardinality)
                rhs18_ = _dafny.Set({(d_80_v4_) == (_dafny.SeqWithoutIsStrInference([d_160_v57_.f10, _dafny.CodePoint('h')])), not((d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))]), not (False) or (d_160_v57_.f9)})
                rhs19_ = d_188_v78_
                rhs20_ = len(d_80_v4_)
                rhs21_ = ((d_191_v79_)[(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgw"))) + (d_80_v4_)] if ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgw"))) + (d_80_v4_)) in (d_191_v79_) else (d_156_v53_).f14)
                lhs13_ = d_82_v6_
                lhs14_ = default__.safeIndex(298, (d_82_v6_).length(0))
                lhs15_ = d_79_globalState_
                lhs13_[lhs14_] = rhs17_
                d_187_v77_ = rhs18_
                d_188_v78_ = rhs19_
                d_75_v0_ = rhs20_
                lhs15_.f5 = rhs21_
                d_75_v0_ = default__.safeDivisionInt(len(d_164_v61_), 59)
            elif True:
                d_192_v80_: D2
                d_192_v80_ = D2_DC4(d_80_v4_, d_180_v73_, d_75_v0_)
                d_193_v81_: _dafny.Map
                d_193_v81_ = _dafny.Map({len(d_80_v4_): (d_160_v57_).f8})
                d_194_v82_: C1
                nw31_ = C1()
                nw31_.ctor__((D5_DC10(d_159_v56_)).cf22, True, (d_80_v4_)[default__.safeIndex((d_159_v56_).cardinality, len(d_80_v4_))], d_85_v7_, d_85_v7_)
                d_194_v82_ = nw31_
                d_195_v83_: D4
                d_195_v83_ = D4_DC8((d_156_v53_).fm18((d_160_v57_).f8, (d_82_v6_)[default__.safeIndex(298, (d_82_v6_).length(0))], (d_156_v53_).fm18(d_160_v57_.f9, len(d_180_v73_), (d_156_v53_).f14, (d_156_v53_).f14, d_79_globalState_), (d_156_v53_).f14, d_79_globalState_), (d_160_v57_).f8, len(d_193_v81_), d_194_v82_)
                d_196_v84_: _dafny.Array
                nw32_ = _dafny.Array(None, 27)
                nw32_[int(0)] = default__.fm23(not(d_160_v57_.f9), d_160_v57_.f9, len(d_180_v73_), d_79_globalState_)
                nw32_[int(1)] = d_180_v73_
                nw32_[int(2)] = d_180_v73_
                nw32_[int(3)] = d_180_v73_
                nw32_[int(4)] = _dafny.SeqWithoutIsStrInference([(d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))]])
                nw32_[int(5)] = d_180_v73_
                nw32_[int(6)] = (d_180_v73_) + (d_180_v73_)
                nw32_[int(7)] = d_180_v73_
                nw32_[int(8)] = d_180_v73_
                nw32_[int(9)] = default__.fm23(not((d_160_v57_).f8), not(d_85_v7_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yowvpprj"))), d_79_globalState_)
                nw32_[int(10)] = (d_192_v80_).cf11
                nw32_[int(11)] = (d_180_v73_) + ((d_180_v73_).set(default__.safeIndex(d_75_v0_, len(d_180_v73_)), d_85_v7_))
                nw32_[int(12)] = d_180_v73_
                nw32_[int(13)] = (d_180_v73_ if (d_195_v83_).cf17 else d_180_v73_)
                nw32_[int(14)] = d_180_v73_
                nw32_[int(15)] = d_180_v73_
                nw32_[int(16)] = ((d_180_v73_) + (_dafny.SeqWithoutIsStrInference([False, d_85_v7_]))).set(default__.safeIndex((d_156_v53_).f14, len((d_180_v73_) + (_dafny.SeqWithoutIsStrInference([False, d_85_v7_])))), d_85_v7_)
                nw32_[int(17)] = d_180_v73_
                nw32_[int(18)] = (d_180_v73_) + (d_180_v73_)
                nw32_[int(19)] = d_180_v73_
                nw32_[int(20)] = _dafny.SeqWithoutIsStrInference([(d_160_v57_).f8, True])
                nw32_[int(21)] = (d_180_v73_) + (d_180_v73_)
                nw32_[int(22)] = (d_180_v73_) + ((d_180_v73_).set(default__.safeIndex((d_82_v6_)[default__.safeIndex(298, (d_82_v6_).length(0))], len(d_180_v73_)), (d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))]))
                nw32_[int(23)] = d_180_v73_
                nw32_[int(24)] = (d_180_v73_) + (d_180_v73_)
                nw32_[int(25)] = (d_180_v73_) + (_dafny.SeqWithoutIsStrInference([(d_194_v82_).f12]))
                nw32_[int(26)] = _dafny.SeqWithoutIsStrInference([(d_160_v57_).f8, (d_194_v82_).f12, (d_194_v82_).f12, not((d_194_v82_).f12), d_160_v57_.f9])
                d_196_v84_ = nw32_
                index37_ = default__.safeIndex(2, (d_196_v84_).length(0))
                (d_196_v84_)[index37_] = default__.fm23((d_160_v57_).f8, (d_160_v57_).f8, len(_dafny.Set({d_160_v57_.f10, _dafny.CodePoint('m')})), d_79_globalState_)
                index38_ = default__.safeIndex(2, (d_196_v84_).length(0))
                (d_196_v84_)[index38_] = d_180_v73_
                index39_ = default__.safeIndex(680, (d_78_v3_).length(0))
                (d_78_v3_)[index39_] = (d_75_v0_) >= ((d_156_v53_).f14)
                d_197_v85_: _dafny.Map
                d_197_v85_ = _dafny.Map({not(d_85_v7_): _dafny.SeqWithoutIsStrInference([(d_160_v57_).f8])})
                (d_79_globalState_).f5 = len((d_197_v85_) | (d_197_v85_))
                d_198_v86_: _dafny.Map
                d_198_v86_ = _dafny.Map({(d_156_v53_).f14: d_90_v11_})
                d_199_v87_: _dafny.Seq
                d_199_v87_ = _dafny.SeqWithoutIsStrInference([(((d_198_v86_)[566] if (566) in (d_198_v86_) else d_90_v11_)).set((d_160_v57_).f8, (d_82_v6_)[default__.safeIndex(298, (d_82_v6_).length(0))]), d_86_v8_, d_90_v11_])
                d_200_v88_: _dafny.Set
                d_200_v88_ = _dafny.Set({True, (d_160_v57_).f8})
                d_201_v89_: _dafny.Map
                d_201_v89_ = _dafny.Map({d_85_v7_: d_86_v8_})
                d_202_v90_: _dafny.Set
                d_202_v90_ = _dafny.Set({(d_199_v87_)[default__.safeIndex(len(d_200_v88_), len(d_199_v87_))], d_90_v11_, ((d_201_v89_)[d_160_v57_.f9] if (d_160_v57_.f9) in (d_201_v89_) else d_90_v11_)})
                d_203_v91_: _dafny.Seq
                d_203_v91_ = _dafny.SeqWithoutIsStrInference([d_76_v1_, d_76_v1_])
                d_204_v92_: _dafny.Set
                d_204_v92_ = _dafny.Set({d_156_v53_, d_156_v53_})
                d_205_v93_: _dafny.Map
                d_205_v93_ = _dafny.Map({(d_76_v1_)[default__.safeIndex((d_156_v53_).fm17(d_160_v57_.f9, False, d_203_v91_, len(d_80_v4_), d_79_globalState_), len(d_76_v1_))]: d_204_v92_})
                d_206_v94_: int
                out6_: int
                out6_ = default__.m0((d_160_v57_).f8, d_202_v90_, (((d_205_v93_)[(d_156_v53_).f14] if ((d_156_v53_).f14) in (d_205_v93_) else d_204_v92_)).ispropersubset(_dafny.Set({d_156_v53_})), d_79_globalState_)
                d_206_v94_ = out6_
                d_207_v95_: C0
                nw33_ = C0()
                nw33_.ctor__(285, _dafny.CodePoint('y'), (d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))], (d_160_v57_).f8)
                d_207_v95_ = nw33_
                d_208_v96_: _dafny.Seq
                d_208_v96_ = _dafny.SeqWithoutIsStrInference([d_207_v95_])
                d_209_v97_: _dafny.Set
                d_209_v97_ = _dafny.Set({d_208_v96_})
                d_209_v97_ = _dafny.Set({d_208_v96_, (d_208_v96_) + (d_208_v96_), d_208_v96_})
            d_210_v98_: _dafny.Set
            d_210_v98_ = _dafny.Set({d_85_v7_, d_160_v57_.f9})
            d_211_v99_: C0
            nw34_ = C0()
            nw34_.ctor__((len(d_210_v98_)) * ((d_160_v57_).fm6((d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))], d_79_globalState_)), d_155_v52_, ((d_78_v3_)[default__.safeIndex(680, (d_78_v3_).length(0))] if not(d_160_v57_.f9) else d_85_v7_), d_85_v7_)
            d_211_v99_ = nw34_
            d_75_v0_ = d_211_v99_.f13
        d_85_v7_ = not(d_85_v7_)
        d_212_v100_: _dafny.Map
        d_212_v100_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_160_v57_.f10 for d_213_i10_ in range(default__.abs(956))]): d_160_v57_.f9})
        d_214_v101_: C0
        nw35_ = C0()
        nw35_.ctor__((d_156_v53_).f14, d_155_v52_, default__.fm22((d_156_v53_).f14, d_160_v57_.f9, (d_160_v57_).f8, d_160_v57_.f9, d_79_globalState_), False)
        d_214_v101_ = nw35_
        d_215_v102_: _dafny.Map
        d_215_v102_ = _dafny.Map({d_75_v0_: d_214_v101_})
        d_212_v100_ = (d_212_v100_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))).set(default__.safeIndex((_dafny.MultiSet([d_214_v101_, d_214_v101_, ((d_215_v102_)[d_214_v101_.f13] if (d_214_v101_.f13) in (d_215_v102_) else d_214_v101_), d_214_v101_])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")))), d_155_v52_), d_160_v57_.f9)
        d_216_v103_: _dafny.Seq
        d_216_v103_ = _dafny.SeqWithoutIsStrInference([(d_160_v57_.f9) == (d_160_v57_.f9), (d_160_v57_).f8])
        d_85_v7_ = (d_216_v103_)[default__.safeIndex((d_214_v101_.f13) * (d_214_v101_.f13), len(d_216_v103_))]
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_78_v3_).length(0)):
            d_217_i11_: int = guard_loop_0_
            if (True) and (((0) <= (d_217_i11_)) and ((d_217_i11_) < ((d_78_v3_).length(0)))):
                (d_78_v3_)[(d_217_i11_)] = (d_164_v61_) == ((d_164_v61_) | (d_164_v61_))
        _dafny.print(_dafny.string_of(d_75_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_v1_) == (_dafny.SeqWithoutIsStrInference([717]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_77_v2_) == (_dafny.Map({717: _dafny.SeqWithoutIsStrInference([717])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_78_v3_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_globalState_.f0) == (_dafny.Map({717: _dafny.SeqWithoutIsStrInference([717])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_79_globalState_).f3)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_79_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_79_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_80_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_81_v5_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbqusrq")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbqusrq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_85_v7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_86_v8_) == (_dafny.Map({True: 717}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_87_v9_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf0) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbqusrq")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbqusrq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf2) == (_dafny.Map({True: 717}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf3)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_89_v10_).cf4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_90_v11_) == (_dafny.Map({True: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf0) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhbqusrq")): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf2) == (_dafny.Map({True: 717}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf3)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_91_v12_).cf4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_155_v52_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_156_v53_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v54_) == (_dafny.SeqWithoutIsStrInference([D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwb"))), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwb"))), D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwb")))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v55_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwb")))]), _dafny.SeqWithoutIsStrInference([D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lwb")))])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v56_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([717])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_v57_.f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v57_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_160_v57_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_161_v58_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[0].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[0]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[0].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[1].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[1]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[1].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[2].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[2]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[2].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[3].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[3]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[3].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[4].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[4]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[4].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[5].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[5]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[5].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[6].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[6]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[6].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[7].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[7]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[7].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[8].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[8]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[8].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[9].f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_162_v59_)[9]).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_162_v59_)[9].f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_163_v60_) == (_dafny.MultiSet([16, 16]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_164_v61_) == (_dafny.Set({16}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_212_v100_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v'), _dafny.CodePoint('v')]): True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_214_v101_.f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_215_v102_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v103_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(_dafny.Map({}), _dafny.Array(None, 0), _dafny.Map({}), _dafny.Array(None, 0), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(int(0), int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({self.cf5.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC4(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC4(D2, NamedTuple('DC4', [('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({self.cf10.VerbatimString(True)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC3(D2, NamedTuple('DC3', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC6(D3, NamedTuple('DC6', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC8(int(0), False, int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)

class D4_DC8(D4, NamedTuple('DC8', [('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC7(D4, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC11()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D5_DC10)

class D5_DC11(D5, NamedTuple('DC11', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC10(D5, NamedTuple('DC10', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC10({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC10) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC14(int(0), int(0), int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)

class D6_DC14(D6, NamedTuple('DC14', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value

class T1:
    pass
    @property
    def f10(self):
        return self._f10
    @f10.setter
    def f10(self, value):
        self._f10 = value

class GlobalState:
    def  __init__(self):
        self.f0: _dafny.Map = _dafny.Map({})
        self.f5: int = int(0)
        self._f1: bool = False
        self._f2: bool = False
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: bool = False
        self._f6: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6

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
    def f6(self):
        return self._f6

class C0(T1, T0):
    def  __init__(self):
        self._f10: str = _dafny.CodePoint('D')
        self._f9: bool = False
        self._f8: bool = False
        self.f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f10(self):
        return self._f10
    @f10.setter
    def f10(self, value):
        self._f10 = value
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f13, f10, f8, f9):
        (self).f13 = f13
        (self).f10 = f10
        (self)._f8 = f8
        (self).f9 = f9

    def fm6(self, p0, globalState):
        return len(_dafny.Set({(self).f8, (self).f8}))

    def fm7(self, p0, p1, p2, globalState):
        return (self).f8

    def fm5(self, p0, globalState):
        return (self).f8


class C1(T1, T0):
    def  __init__(self):
        self._f10: str = _dafny.CodePoint('D')
        self._f9: bool = False
        self._f8: bool = False
        self._f11: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f10(self):
        return self._f10
    @f10.setter
    def f10(self, value):
        self._f10 = value
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f11, f12, f10, f8, f9):
        (self)._f11 = f11
        (self)._f12 = f12
        (self).f10 = f10
        (self)._f8 = f8
        (self).f9 = f9

    def fm6(self, p0, globalState):
        return len((_dafny.Map({len(_dafny.Set({(self).f12})): (self).f8})) | (_dafny.Map({18: (self).f12})))

    def fm7(self, p0, p1, p2, globalState):
        return (self).f8

    def fm5(self, p0, globalState):
        return not((_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f12])).cardinality])) <= (_dafny.SeqWithoutIsStrInference([42 for d_218_i0_ in range(default__.abs(495))])))

    def fm8(self, globalState):
        return (0) - (-303)

    def m4(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: bool = False
        r3: bool = False
        r1 = (p0) == (346)
        d_219_v0_: T1
        nw36_ = C0()
        nw36_.ctor__(303, self.f10, p3, (self).fm7(False, _dafny.Set({True}), p1, globalState))
        d_219_v0_ = nw36_
        d_220_v1_: _dafny.Map
        d_220_v1_ = _dafny.Map({(0) - (p0): d_219_v0_})
        d_220_v1_ = d_220_v1_
        d_221_v2_: _dafny.Seq
        d_221_v2_ = _dafny.SeqWithoutIsStrInference([(self).f8, (self).f12])
        d_222_v3_: D2
        d_222_v3_ = D2_DC3((d_221_v2_)[default__.safeIndex(p0, len(d_221_v2_))])
        source3_ = d_222_v3_
        if source3_.is_DC4:
            d_223___mcc_h0_ = source3_.cf10
            d_224___mcc_h1_ = source3_.cf11
            d_225___mcc_h2_ = source3_.cf12
            d_226_cf12_ = d_225___mcc_h2_
            d_227_cf11_ = d_224___mcc_h1_
            d_228_cf10_ = d_223___mcc_h0_
            d_229_v4_: _dafny.Map
            d_229_v4_ = _dafny.Map({p1: d_226_cf12_})
            r0 = ((d_229_v4_)[p1] if (p1) in (d_229_v4_) else default__.fm0(p3, d_228_cf10_, globalState))
            d_230_v5_: _dafny.MultiSet
            d_230_v5_ = _dafny.MultiSet([True, not((d_222_v3_).cf9), False, True, (d_219_v0_).f8])
            d_231_v6_: _dafny.Map
            d_231_v6_ = _dafny.Map({default__.fm9(d_226_cf12_, (d_219_v0_).f8, globalState): d_230_v5_})
            d_232_v7_: _dafny.Set
            d_232_v7_ = _dafny.Set({d_219_v0_.f9})
            rhs22_ = True
            rhs23_ = D2_DC3((d_232_v7_).issubset(d_232_v7_))
            rhs24_ = (p3) == ((d_232_v7_).issubset(d_232_v7_))
            rhs25_ = d_231_v6_
            rhs26_ = (d_230_v5_).issubset(d_230_v5_)
            lhs16_ = self
            lhs16_.f9 = rhs22_
            d_222_v3_ = rhs23_
            r3 = rhs24_
            d_231_v6_ = rhs25_
            r3 = rhs26_
            d_233_v8_: _dafny.Map
            d_233_v8_ = _dafny.Map({p0: d_226_cf12_})
            d_234_v9_: _dafny.Map
            d_234_v9_ = _dafny.Map({(d_222_v3_).cf9: d_228_cf10_})
            d_233_v8_ = default__.fm10((d_226_cf12_) + (p0), d_234_v9_, globalState)
            d_235_v10_: _dafny.Map
            d_235_v10_ = _dafny.Map({((d_233_v8_)[d_226_cf12_] if (d_226_cf12_) in (d_233_v8_) else p0): _dafny.CodePoint('j')})
            d_236_v11_: D1
            d_236_v11_ = D1_DC2(len(d_235_v10_), len(d_227_cf11_), d_219_v0_.f10)
            (globalState).f5 = (d_236_v11_).cf7
        elif source3_.is_DC3:
            d_237___mcc_h3_ = source3_.cf9
            d_238_cf9_ = d_237___mcc_h3_
            d_239_v12_: _dafny.Array
            def lambda13_(d_240_p0_):
                def lambda14_(d_241_i0_):
                    return _dafny.SeqWithoutIsStrInference([d_240_p0_ for d_242_i1_ in range(default__.abs(56))])

                return lambda14_

            init8_ = lambda13_(p0)
            nw37_ = _dafny.Array(None, 8)
            for i0_8_ in range(nw37_.length(0)):
                nw37_[i0_8_] = init8_(i0_8_)
            d_239_v12_ = nw37_
            d_243_v13_: _dafny.Seq
            d_243_v13_ = _dafny.SeqWithoutIsStrInference([392])
            index40_ = default__.safeIndex(707, (d_239_v12_).length(0))
            (d_239_v12_)[index40_] = d_243_v13_
            d_244_v14_: _dafny.Seq
            d_244_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            index41_ = default__.safeIndex(707, (d_239_v12_).length(0))
            rhs27_ = d_243_v13_
            rhs28_ = d_238_cf9_
            rhs29_ = d_244_v14_
            rhs30_ = d_244_v14_
            rhs31_ = (self).f8
            lhs17_ = d_239_v12_
            lhs18_ = default__.safeIndex(707, (d_239_v12_).length(0))
            lhs19_ = d_219_v0_
            lhs17_[lhs18_] = rhs27_
            lhs19_.f9 = rhs28_
            d_244_v14_ = rhs29_
            d_244_v14_ = rhs30_
            d_238_cf9_ = rhs31_
            d_245_v15_: _dafny.Set
            d_245_v15_ = _dafny.Set({(0) - ((p0) - (p0)), p0, (0) - (p0), p0})
            d_245_v15_ = d_245_v15_
            r2 = ((self.f10 if d_219_v0_.f9 else p2)) not in (p1)
            d_246_v16_: _dafny.Map
            d_246_v16_ = _dafny.Map({self.f9: p0})
            d_247_v17_: int
            out7_: int
            out7_ = default__.m0(d_219_v0_.f9, _dafny.Set({d_246_v16_}), (d_222_v3_).cf9, globalState)
            d_247_v17_ = out7_
        elif True:
            d_248___mcc_h4_ = source3_.cf13
            d_249_cf13_ = d_248___mcc_h4_
            d_250_v18_: _dafny.Seq
            d_250_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eahskv"))])
            d_251_v19_: _dafny.Map
            d_251_v19_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_219_v0_.f10 for d_252_i2_ in range(default__.abs(740))]): (d_250_v18_)[default__.safeIndex(p0, len(d_250_v18_))]})
            d_253_v20_: _dafny.Array
            def lambda15_(d_254_i3_):
                return default__.safeModuloInt(d_254_i3_, len(_dafny.Set({not(True), False})))

            init9_ = lambda15_
            nw38_ = _dafny.Array(None, 21)
            for i0_9_ in range(nw38_.length(0)):
                nw38_[i0_9_] = init9_(i0_9_)
            d_253_v20_ = nw38_
            d_255_v21_: _dafny.Map
            d_255_v21_ = _dafny.Map({d_219_v0_.f9: p0})
            d_256_v22_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_256_v22_ = nw39_
            d_257_v23_: D0
            d_257_v23_ = D0_DC0(d_251_v19_, d_253_v20_, (d_255_v21_).set((d_219_v0_).fm5((d_219_v0_).f8, globalState), p0), d_256_v22_, d_253_v20_)
            d_257_v23_ = d_257_v23_
            (globalState).f5 = p0
            d_258_v24_: _dafny.Set
            d_258_v24_ = _dafny.Set({(p1) + (p1), p1})
            d_258_v24_ = (d_258_v24_) - (_dafny.Set({(p1).set(default__.safeIndex(p0, len(p1)), _dafny.CodePoint('n'))}))
            d_259_v25_: D2
            d_259_v25_ = D2_DC4(p1, d_221_v2_, p0)
            d_260_v26_: _dafny.Seq
            d_260_v26_ = _dafny.SeqWithoutIsStrInference([d_259_v25_])
            d_261_v27_: D2
            d_261_v27_ = D2_DC5((d_260_v26_)[default__.safeIndex(p0, len(d_260_v26_))])
            d_262_v28_: D2
            d_262_v28_ = D2_DC5(d_261_v27_)
            d_263_v29_: D2
            d_263_v29_ = D2_DC5(d_259_v25_)
            d_264_v30_: D2
            d_264_v30_ = D2_DC5(d_261_v27_)
            d_265_v31_: _dafny.Map
            d_265_v31_ = _dafny.Map({p0: _dafny.Set({p0, p0})})
            d_266_v32_: _dafny.Set
            d_266_v32_ = _dafny.Set({p0})
            rhs32_ = (939) - (default__.safeDivisionInt(-269, len(d_221_v2_)))
            rhs33_ = self.f9
            rhs34_ = not((((d_265_v31_)[p0] if (p0) in (d_265_v31_) else d_266_v32_)) == (d_266_v32_))
            rhs35_ = len(_dafny.SeqWithoutIsStrInference([d_255_v21_, d_255_v21_, d_255_v21_]))
            rhs36_ = d_264_v30_
            lhs20_ = globalState
            lhs21_ = d_219_v0_
            lhs22_ = globalState
            lhs20_.f5 = rhs32_
            lhs21_.f9 = rhs33_
            r2 = rhs34_
            lhs22_.f5 = rhs35_
            d_264_v30_ = rhs36_
        d_267_v33_: _dafny.Map
        d_267_v33_ = _dafny.Map({(d_219_v0_).fm5((self).f12, globalState): p0})
        d_268_v34_: _dafny.Set
        d_268_v34_ = _dafny.Set({d_267_v33_, d_267_v33_})
        d_269_v35_: int
        out8_: int
        out8_ = default__.m0((True if (self).f8 else True), d_268_v34_, not((p0) >= (p0)), globalState)
        d_269_v35_ = out8_
        d_270_v36_: _dafny.Set
        d_270_v36_ = _dafny.Set({self.f10, p2})
        d_271_v37_: C0
        nw40_ = C0()
        nw40_.ctor__(p0, d_219_v0_.f10, not (not((self).f8)) or (self.f9), (d_219_v0_.f10) not in (d_270_v36_))
        d_271_v37_ = nw40_
        d_271_v37_ = d_271_v37_
        (self).f10 = (p1)[default__.safeIndex((0) - (d_269_v35_), len(p1))]
        r0 = p0
        d_272_v38_: _dafny.Set
        d_272_v38_ = _dafny.Set({(d_219_v0_).f8})
        r1 = (d_219_v0_).fm7((self).f8, d_272_v38_, p1, globalState)
        r2 = (_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(p0, len(p1))] for d_273_i4_ in range(default__.abs(150))])) <= ((D1_DC1(_dafny.SeqWithoutIsStrInference([d_219_v0_.f10 for d_274_i5_ in range(default__.abs(197))]))).cf5)
        r3 = self.f9
        return r0, r1, r2, r3

    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12

class C2(T0):
    def  __init__(self):
        self._f9: bool = False
        self._f8: bool = False
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f9(self):
        return self._f9
    @f9.setter
    def f9(self, value):
        self._f9 = value
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f14, f8, f9):
        (self)._f14 = f14
        (self)._f8 = f8
        (self).f9 = f9

    def fm5(self, p0, globalState):
        return (self).f8

    def fm17(self, p0, p1, p2, p3, globalState):
        source4_ = D2_DC3((self).f8)
        if source4_.is_DC4:
            d_275___mcc_h0_ = source4_.cf10
            d_276___mcc_h1_ = source4_.cf11
            d_277___mcc_h2_ = source4_.cf12
            d_278_cf12_ = d_277___mcc_h2_
            d_279_cf11_ = d_276___mcc_h1_
            d_280_cf10_ = d_275___mcc_h0_
            return len((d_280_cf10_) + (d_280_cf10_))
        elif source4_.is_DC3:
            d_281___mcc_h3_ = source4_.cf9
            d_282_cf9_ = d_281___mcc_h3_
            return (self).f14
        elif True:
            d_283___mcc_h4_ = source4_.cf13
            d_284_cf13_ = d_283___mcc_h4_
            return ((self).f14) * ((0) - ((self).f14))

    def fm18(self, p0, p1, p2, p3, globalState):
        return (0) - (((self).f14) - (len(_dafny.Map({(self).f14: (self).f14}))))

    @property
    def f14(self):
        return self._f14

class C3:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm3(self, globalState):
        return (_dafny.MultiSet([len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwqewtt"))), 318]))})), -220])).isdisjoint((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([128]))))

    def fm4(self, globalState):
        return D2_DC4((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmapr"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "he"))), _dafny.SeqWithoutIsStrInference([False, False]), default__.safeDivisionInt((0) - (-872), 738))

    def m2(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: str = _dafny.CodePoint('D')
        r2: bool = False
        d_285_v0_: str
        d_285_v0_ = _dafny.CodePoint('t')
        d_286_v1_: C0
        nw41_ = C0()
        nw41_.ctor__(p2, d_285_v0_, p0, not(p0))
        d_286_v1_ = nw41_
        d_287_v2_: _dafny.Array
        nw42_ = _dafny.Array(False, 26)
        d_287_v2_ = nw42_
        d_288_v3_: _dafny.Seq
        d_288_v3_ = _dafny.SeqWithoutIsStrInference([d_287_v2_])
        d_289_v4_: _dafny.Map
        d_289_v4_ = _dafny.Map({p0: d_287_v2_})
        d_290_v5_: _dafny.Array
        nw43_ = _dafny.Array(None, 8)
        nw43_[int(0)] = (d_288_v3_)[default__.safeIndex(p2, len(d_288_v3_))]
        nw43_[int(1)] = d_287_v2_
        nw43_[int(2)] = d_287_v2_
        nw43_[int(3)] = d_287_v2_
        nw43_[int(4)] = d_287_v2_
        nw43_[int(5)] = d_287_v2_
        nw43_[int(6)] = d_287_v2_
        nw43_[int(7)] = ((d_289_v4_)[p0] if (p0) in (d_289_v4_) else d_287_v2_)
        d_290_v5_ = nw43_
        index42_ = default__.safeIndex(280, (d_290_v5_).length(0))
        (d_290_v5_)[index42_] = d_287_v2_
        d_291_v6_: _dafny.Array
        def lambda16_(d_292_i0_):
            return (d_292_i0_) * (968)

        init10_ = lambda16_
        nw44_ = _dafny.Array(None, 25)
        for i0_10_ in range(nw44_.length(0)):
            nw44_[i0_10_] = init10_(i0_10_)
        d_291_v6_ = nw44_
        d_293_v7_: _dafny.Map
        d_293_v7_ = _dafny.Map({d_291_v6_: not(not(p0))})
        d_294_v8_: _dafny.Set
        d_294_v8_ = _dafny.Set({d_285_v0_})
        d_295_v9_: _dafny.Map
        d_295_v9_ = _dafny.Map({-717: not(p0)})
        index43_ = default__.safeIndex(280, (d_290_v5_).length(0))
        nw45_ = _dafny.Array(None, 7)
        nw45_[int(0)] = (p0 if p0 else p0)
        nw45_[int(1)] = p0
        nw45_[int(2)] = p0
        nw45_[int(3)] = (d_291_v6_) not in (d_293_v7_)
        nw45_[int(4)] = (_dafny.Set({d_285_v0_})).issubset(d_294_v8_)
        nw45_[int(5)] = ((D2_DC3(not(p0))).cf9) == (p0)
        nw45_[int(6)] = (d_295_v9_) == (d_295_v9_)
        (d_290_v5_)[index43_] = nw45_
        hi0_ = d_286_v1_.f13
        for d_296_i1_ in range(d_286_v1_.f13, hi0_):
            r2 = p0
            d_297_v10_: _dafny.Set
            d_297_v10_ = _dafny.Set({not(p0)})
            index44_ = default__.safeIndex(449, (d_291_v6_).length(0))
            (d_291_v6_)[index44_] = len(d_297_v10_)
            d_298_v11_: _dafny.Seq
            d_298_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nurt"))
            index45_ = default__.safeIndex(449, (d_291_v6_).length(0))
            (d_291_v6_)[index45_] = default__.safeDivisionInt(len((d_298_v11_) + (d_298_v11_)), d_296_i1_)
            arr0_ = (d_290_v5_)[default__.safeIndex(280, (d_290_v5_).length(0))]
            index46_ = default__.safeIndex(638, ((d_290_v5_)[default__.safeIndex(280, (d_290_v5_).length(0))]).length(0))
            arr0_[index46_] = p0
            d_299_v12_: _dafny.Map
            d_299_v12_ = _dafny.Map({p0: d_285_v0_})
            d_300_v13_: _dafny.Seq
            d_300_v13_ = _dafny.SeqWithoutIsStrInference([len((d_299_v12_).set(p0, d_285_v0_)), d_296_i1_])
            arr1_ = (d_290_v5_)[default__.safeIndex(280, (d_290_v5_).length(0))]
            index47_ = default__.safeIndex(638, ((d_290_v5_)[default__.safeIndex(280, (d_290_v5_).length(0))]).length(0))
            arr1_[index47_] = (d_296_i1_) not in (d_300_v13_)
            index48_ = default__.safeIndex(449, (d_291_v6_).length(0))
            (d_291_v6_)[index48_] = ((355) * (p2)) - (d_296_i1_)
        d_301_v14_: _dafny.Map
        d_301_v14_ = _dafny.Map({(d_288_v3_)[default__.safeIndex(p2, len(d_288_v3_))]: d_285_v0_})
        d_301_v14_ = d_301_v14_
        d_302_v15_: _dafny.Seq
        d_302_v15_ = _dafny.SeqWithoutIsStrInference([p2, d_286_v1_.f13, d_286_v1_.f13, p2])
        d_303_v17_: _dafny.Map
        def iife17_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in (d_302_v15_).Elements:
                d_304_v16_: int = compr_11_
                if (d_304_v16_) in (d_302_v15_):
                    coll11_ = coll11_.union(_dafny.Set([(d_304_v16_) * (679)]))
            return _dafny.Set(coll11_)
        d_303_v17_ = _dafny.Map({(p2) + ((0) - (p2)): iife17_()
        })
        d_305_v18_: _dafny.Map
        d_305_v18_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference([p0])})
        d_306_v19_: _dafny.Seq
        d_306_v19_ = _dafny.SeqWithoutIsStrInference([p0])
        d_307_v20_: _dafny.Set
        d_307_v20_ = _dafny.Set({len(((d_305_v18_)[p0] if (p0) in (d_305_v18_) else d_306_v19_))})
        d_303_v17_ = (_dafny.Map({d_286_v1_.f13: d_307_v20_})).set(p2, d_307_v20_)
        d_308_v21_: D2
        d_308_v21_ = D2_DC3(p0)
        pat_let_tv3_ = p0
        def iife18_(_pat_let3_0):
            def iife19_(d_309_dt__update__tmp_h0_):
                def iife20_(_pat_let4_0):
                    def iife21_(d_310_dt__update_hcf9_h0_):
                        return D2_DC3(d_310_dt__update_hcf9_h0_)
                    return iife21_(_pat_let4_0)
                return iife20_(pat_let_tv3_)
            return iife19_(_pat_let3_0)
        d_308_v21_ = (default__.fm11(p2, d_286_v1_.f13, iife18_(d_308_v21_), p0, globalState) if p0 else d_308_v21_)
        d_311_v22_: _dafny.Seq
        d_311_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jw"))
        d_312_v23_: D1
        d_312_v23_ = D1_DC1(d_311_v22_)
        d_313_v24_: _dafny.Set
        d_313_v24_ = _dafny.Set({((d_312_v23_).cf5) + (d_311_v22_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nujlcfog")), d_311_v22_})
        r0 = len(d_313_v24_)
        r1 = d_285_v0_
        d_314_v25_: _dafny.MultiSet
        d_314_v25_ = _dafny.MultiSet([p0])
        r2 = not((_dafny.MultiSet([p0, p0])) == (d_314_v25_))
        return r0, r1, r2

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: int = int(0)
        if False:
            d_315_v0_: _dafny.MultiSet
            d_315_v0_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([p0]), _dafny.SeqWithoutIsStrInference([p0, p0])])
            d_316_v1_: C1
            nw46_ = C1()
            nw46_.ctor__(d_315_v0_, p3, _dafny.CodePoint('a'), p3, p3)
            d_316_v1_ = nw46_
            d_317_v2_: _dafny.Seq
            d_317_v2_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), len(_dafny.SeqWithoutIsStrInference([d_316_v1_, d_316_v1_, d_316_v1_, d_316_v1_])), 444, p2])
            d_318_v3_: _dafny.Map
            d_318_v3_ = _dafny.Map({p2: (d_316_v1_).f12})
            d_319_v4_: _dafny.Seq
            d_319_v4_ = _dafny.SeqWithoutIsStrInference([len(d_317_v2_), len(d_318_v3_), (0) - (p0)])
            d_320_v5_: _dafny.MultiSet
            d_320_v5_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([len(d_317_v2_), 767]), d_317_v2_, d_319_v4_, d_319_v4_])
            d_321_v6_: str
            d_321_v6_ = _dafny.CodePoint('q')
            d_322_v7_: _dafny.Seq
            d_322_v7_ = _dafny.SeqWithoutIsStrInference([p3])
            d_323_v8_: C1
            nw47_ = C1()
            nw47_.ctor__((default__.fm12(p3, globalState)) - (d_320_v5_), True, d_321_v6_, (((d_318_v3_)[p0] if (p0) in (d_318_v3_) else not((d_322_v7_)[default__.safeIndex(p0, len(d_322_v7_))])) if False else not(p3)), (d_316_v1_).f12)
            d_323_v8_ = nw47_
            d_324_v9_: _dafny.MultiSet
            d_324_v9_ = _dafny.MultiSet([p3])
            d_325_v10_: _dafny.MultiSet
            d_325_v10_ = _dafny.MultiSet([(_dafny.MultiSet([(d_316_v1_).f12]) if False else d_324_v9_)])
            d_326_v11_: _dafny.Seq
            d_326_v11_ = _dafny.SeqWithoutIsStrInference([d_325_v10_])
            d_325_v10_ = (d_326_v11_)[default__.safeIndex(p2, len(d_326_v11_))]
            d_327_v12_: bool
            d_327_v12_ = False
            d_328_v13_: _dafny.Set
            d_328_v13_ = _dafny.Set({d_327_v12_})
            d_329_v14_: _dafny.Seq
            d_329_v14_ = _dafny.SeqWithoutIsStrInference([d_328_v13_])
            rhs37_ = not(p3)
            rhs38_ = len((_dafny.Set({(d_316_v1_).f12}) if (d_323_v8_).f12 else (d_329_v14_)[default__.safeIndex(689, len(d_329_v14_))]))
            d_327_v12_ = rhs37_
            r0 = rhs38_
            d_330_v15_: _dafny.Array
            nw48_ = _dafny.Array(int(0), 15)
            d_330_v15_ = nw48_
            index49_ = default__.safeIndex(690, (d_330_v15_).length(0))
            (d_330_v15_)[index49_] = p0
            index50_ = default__.safeIndex(690, (d_330_v15_).length(0))
            (d_330_v15_)[index50_] = p0
            d_328_v13_ = (d_328_v13_) - ((d_328_v13_) - (_dafny.Set({False})))
        elif True:
            d_331_v16_: _dafny.Array
            def lambda17_(d_332_i0_):
                return False

            init11_ = lambda17_
            nw49_ = _dafny.Array(None, 4)
            for i0_11_ in range(nw49_.length(0)):
                nw49_[i0_11_] = init11_(i0_11_)
            d_331_v16_ = nw49_
            index51_ = default__.safeIndex(609, (d_331_v16_).length(0))
            (d_331_v16_)[index51_] = True
            index52_ = default__.safeIndex(609, (d_331_v16_).length(0))
            (d_331_v16_)[index52_] = not(p3)
            d_333_v17_: _dafny.Set
            d_333_v17_ = _dafny.Set({-311})
            d_334_v18_: _dafny.MultiSet
            d_334_v18_ = _dafny.MultiSet([p0, len(d_333_v17_)])
            d_335_v19_: _dafny.Seq
            d_335_v19_ = _dafny.SeqWithoutIsStrInference([(d_331_v16_)[default__.safeIndex(609, (d_331_v16_).length(0))], (p2) <= ((d_334_v18_).cardinality), (d_331_v16_)[default__.safeIndex(609, (d_331_v16_).length(0))]])
            (globalState).f5 = len(d_335_v19_)
            d_336_v20_: _dafny.Map
            d_336_v20_ = _dafny.Map({p0: 34})
            d_337_v21_: _dafny.Map
            d_337_v21_ = _dafny.Map({p2: d_336_v20_})
            d_338_v23_: _dafny.MultiSet
            def iife22_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(139, 19):
                    d_339_v22_: int = compr_12_
                    if ((139) <= (d_339_v22_)) and ((d_339_v22_) < (19)):
                        coll12_[(d_339_v22_) + (p2)] = p0
                return _dafny.Map(coll12_)
            d_338_v23_ = _dafny.MultiSet([((d_337_v21_)[p0] if (p0) in (d_337_v21_) else iife22_()
            ), _dafny.Map({p2: p2}), d_336_v20_])
            d_340_v24_: _dafny.Map
            d_340_v24_ = _dafny.Map({default__.fm13(p2, globalState): (d_338_v23_).cardinality})
            d_341_v25_: str
            d_341_v25_ = _dafny.CodePoint('f')
            d_340_v24_ = (d_340_v24_).set((d_341_v25_ if p3 else d_341_v25_), p0)
            (globalState).f5 = (default__.safeModuloInt(p2, p0)) + (p0)
            d_342_v26_: _dafny.Seq
            d_342_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb"))
            index53_ = default__.safeIndex(609, (d_331_v16_).length(0))
            (d_331_v16_)[index53_] = (_dafny.SeqWithoutIsStrInference([d_341_v25_ for d_343_i1_ in range(default__.abs(726))])) < (d_342_v26_)
        d_344_v27_: bool
        d_344_v27_ = False
        d_344_v27_ = (d_344_v27_ if p3 else p3)
        r0 = ((p2) - (792)) * (p2)
        d_345_i2_: int
        d_345_i2_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_345_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_345_i2_ = (d_345_i2_) + (1)
                    d_346_v28_: _dafny.Seq
                    d_346_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rfvu"))
                    if (default__.safeDivisionInt(p0, len(d_346_v28_))) <= (p0):
                        d_347_v29_: _dafny.Seq
                        d_347_v29_ = _dafny.SeqWithoutIsStrInference([990, len(d_346_v28_), p2])
                        d_348_v30_: str
                        d_348_v30_ = _dafny.CodePoint('v')
                        d_349_v31_: _dafny.Map
                        d_349_v31_ = _dafny.Map({d_344_v27_: d_348_v30_})
                        d_350_v32_: _dafny.Set
                        d_350_v32_ = _dafny.Set({len(d_349_v31_), p2, (0) - (p0)})
                        d_351_v33_: _dafny.MultiSet
                        d_351_v33_ = _dafny.MultiSet([(d_350_v32_ if not(p3) else d_350_v32_), d_350_v32_, d_350_v32_])
                        rhs39_ = p0
                        rhs40_ = d_347_v29_
                        rhs41_ = _dafny.MultiSet([d_350_v32_])
                        lhs23_ = globalState
                        lhs23_.f5 = rhs39_
                        d_347_v29_ = rhs40_
                        d_351_v33_ = rhs41_
                        d_352_v34_: D2
                        d_352_v34_ = D2_DC4(d_346_v28_, _dafny.SeqWithoutIsStrInference([d_344_v27_]), p2)
                        d_352_v34_ = d_352_v34_
                        rhs42_ = p3
                        rhs43_ = p3
                        d_344_v27_ = rhs42_
                        d_344_v27_ = rhs43_
                        d_353_v35_: _dafny.Map
                        d_353_v35_ = _dafny.Map({not(not (d_344_v27_) or (d_344_v27_)): d_346_v28_})
                        d_353_v35_ = (d_353_v35_) | (default__.fm14(p2, d_348_v30_, p2, globalState))
                        d_354_v36_: D1
                        d_354_v36_ = D1_DC1(d_346_v28_)
                        r0 = default__.fm0(p3, (d_354_v36_).cf5, globalState)
                    elif True:
                        d_344_v27_ = not (False) or ((d_344_v27_ if d_344_v27_ else d_344_v27_))
                        d_355_v37_: _dafny.Seq
                        d_355_v37_ = _dafny.SeqWithoutIsStrInference([p2])
                        d_356_v38_: _dafny.Set
                        d_356_v38_ = _dafny.Set({d_355_v37_})
                        d_357_v39_: _dafny.Seq
                        d_357_v39_ = _dafny.SeqWithoutIsStrInference([p2, len(d_356_v38_), p2])
                        d_358_v40_: _dafny.Seq
                        d_358_v40_ = _dafny.SeqWithoutIsStrInference([not(False), p3])
                        d_359_v41_: str
                        d_359_v41_ = _dafny.CodePoint('k')
                        d_360_v42_: D1
                        d_360_v42_ = D1_DC2(len(_dafny.SeqWithoutIsStrInference([p2, (_dafny.MultiSet(d_357_v39_)).cardinality, (_dafny.MultiSet(d_358_v40_)).cardinality])), p2, d_359_v41_)
                        (globalState).f5 = (d_360_v42_).cf7
                        d_344_v27_ = d_344_v27_
                        d_361_v43_: C0
                        nw50_ = C0()
                        nw50_.ctor__(722, d_359_v41_, d_344_v27_, d_344_v27_)
                        d_361_v43_ = nw50_
                        d_361_v43_ = d_361_v43_
                        d_362_v44_: D2
                        d_362_v44_ = D2_DC3(not(d_344_v27_))
                        d_363_v45_: _dafny.Array
                        nw51_ = _dafny.Array(None, 15)
                        nw51_[int(0)] = d_362_v44_
                        nw51_[int(1)] = d_362_v44_
                        nw51_[int(2)] = D2_DC3(d_344_v27_)
                        nw51_[int(3)] = d_362_v44_
                        nw51_[int(4)] = d_362_v44_
                        nw51_[int(5)] = D2_DC3(not(d_344_v27_))
                        nw51_[int(6)] = d_362_v44_
                        nw51_[int(7)] = d_362_v44_
                        nw51_[int(8)] = d_362_v44_
                        nw51_[int(9)] = d_362_v44_
                        nw51_[int(10)] = d_362_v44_
                        nw51_[int(11)] = d_362_v44_
                        nw51_[int(12)] = d_362_v44_
                        nw51_[int(13)] = d_362_v44_
                        nw51_[int(14)] = d_362_v44_
                        d_363_v45_ = nw51_
                        d_364_v46_: _dafny.MultiSet
                        d_364_v46_ = _dafny.MultiSet([p2, d_361_v43_.f13])
                        d_365_v47_: _dafny.Map
                        d_365_v47_ = _dafny.Map({d_363_v45_: (d_361_v43_.f13) == ((d_364_v46_).cardinality)})
                        d_366_v48_: _dafny.Seq
                        d_366_v48_ = _dafny.SeqWithoutIsStrInference([d_363_v45_])
                        d_365_v47_ = (d_365_v47_).set((d_366_v48_)[default__.safeIndex((_dafny.MultiSet(d_355_v37_)).cardinality, len(d_366_v48_))], d_344_v27_)
                    d_367_v49_: _dafny.Array
                    def lambda18_(d_368_p2_):
                        def lambda19_(d_369_i3_):
                            return (d_369_i3_) + (d_368_p2_)

                        return lambda19_

                    init12_ = lambda18_(p2)
                    nw52_ = _dafny.Array(None, 12)
                    for i0_12_ in range(nw52_.length(0)):
                        nw52_[i0_12_] = init12_(i0_12_)
                    d_367_v49_ = nw52_
                    d_370_v50_: str
                    d_370_v50_ = _dafny.CodePoint('g')
                    d_371_v51_: _dafny.Seq
                    d_371_v51_ = _dafny.SeqWithoutIsStrInference([d_346_v28_, d_346_v28_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "molvhg")), (d_346_v28_).set(default__.safeIndex(p0, len(d_346_v28_)), d_370_v50_)])
                    index54_ = default__.safeIndex(988, (d_367_v49_).length(0))
                    (d_367_v49_)[index54_] = len(d_371_v51_)
                    d_372_v52_: _dafny.Array
                    nw53_ = _dafny.Array(_dafny.MultiSet({}), 12)
                    d_372_v52_ = nw53_
                    d_373_v53_: _dafny.MultiSet
                    d_373_v53_ = _dafny.MultiSet([(d_372_v52_ if not(d_344_v27_) else d_372_v52_), d_372_v52_, d_372_v52_, d_372_v52_])
                    d_374_v54_: _dafny.MultiSet
                    d_374_v54_ = _dafny.MultiSet([p3, p3])
                    index55_ = default__.safeIndex(988, (d_367_v49_).length(0))
                    rhs44_ = (p2) + (p0)
                    rhs45_ = p2
                    rhs46_ = (d_373_v53_) - (d_373_v53_)
                    rhs47_ = ((d_374_v54_ if p3 else (_dafny.MultiSet([d_344_v27_, d_344_v27_])).set(d_344_v27_, default__.abs(p2)))).isdisjoint(d_374_v54_)
                    lhs24_ = d_367_v49_
                    lhs25_ = default__.safeIndex(988, (d_367_v49_).length(0))
                    lhs26_ = globalState
                    lhs24_[lhs25_] = rhs44_
                    lhs26_.f5 = rhs45_
                    d_373_v53_ = rhs46_
                    d_344_v27_ = rhs47_
                    d_375_v55_: _dafny.Seq
                    d_375_v55_ = _dafny.SeqWithoutIsStrInference([d_344_v27_, not((self).fm3(globalState)), (p2) == (779)])
                    d_375_v55_ = (_dafny.SeqWithoutIsStrInference([True])) + ((d_375_v55_) + (d_375_v55_))
                    d_376_v56_: _dafny.MultiSet
                    d_376_v56_ = _dafny.MultiSet([p0])
                    d_377_v57_: _dafny.Seq
                    d_377_v57_ = _dafny.SeqWithoutIsStrInference([(0) - (p0), (0) - ((d_376_v56_).cardinality), p2])
                    d_378_v58_: _dafny.Seq
                    d_378_v58_ = _dafny.SeqWithoutIsStrInference([d_377_v57_, d_377_v57_, d_377_v57_])
                    d_379_v59_: C1
                    nw54_ = C1()
                    nw54_.ctor__(_dafny.MultiSet(d_378_v58_), not(p3), d_370_v50_, (p3) or (True), (d_375_v55_)[default__.safeIndex((d_367_v49_)[default__.safeIndex(988, (d_367_v49_).length(0))], len(d_375_v55_))])
                    d_379_v59_ = nw54_
                    pass
            pass
        d_380_v60_: _dafny.Array
        nw55_ = _dafny.Array(_dafny.CodePoint('D'), 18)
        d_380_v60_ = nw55_
        d_381_v61_: _dafny.Array
        def lambda20_(d_382_v27_):
            def lambda21_(d_383_i4_):
                return d_382_v27_

            return lambda21_

        init13_ = lambda20_(d_344_v27_)
        nw56_ = _dafny.Array(None, 21)
        for i0_13_ in range(nw56_.length(0)):
            nw56_[i0_13_] = init13_(i0_13_)
        d_381_v61_ = nw56_
        d_384_v62_: _dafny.MultiSet
        d_384_v62_ = _dafny.MultiSet([d_381_v61_, d_381_v61_])
        d_385_v63_: int
        d_386_v64_: str
        d_387_v65_: bool
        out9_: int
        out10_: str
        out11_: bool
        out9_, out10_, out11_ = (self).m2(not(p3), d_380_v60_, default__.safeModuloInt(p0, p0), d_384_v62_, globalState)
        d_385_v63_ = out9_
        d_386_v64_ = out10_
        d_387_v65_ = out11_
        d_388_v66_: _dafny.Seq
        d_388_v66_ = _dafny.SeqWithoutIsStrInference([d_387_v65_])
        hi1_ = len(d_388_v66_)
        for d_389_i5_ in range(p0, hi1_):
            d_390_v67_: C0
            nw57_ = C0()
            nw57_.ctor__(p0, d_386_v64_, d_387_v65_, d_387_v65_)
            d_390_v67_ = nw57_
            d_391_v68_: _dafny.Map
            d_391_v68_ = _dafny.Map({p0: d_386_v64_})
            d_391_v68_ = (d_391_v68_).set(default__.safeModuloInt(d_389_i5_, 579), d_386_v64_)
            d_392_v69_: _dafny.Array
            nw58_ = _dafny.Array(int(0), 22)
            d_392_v69_ = nw58_
            index56_ = default__.safeIndex(692, (d_392_v69_).length(0))
            (d_392_v69_)[index56_] = (d_385_v63_) - (p2)
            index57_ = default__.safeIndex(692, (d_392_v69_).length(0))
            (d_392_v69_)[index57_] = d_385_v63_
            if d_344_v27_:
                d_393_v70_: _dafny.Seq
                d_393_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xp"))
                d_394_v71_: _dafny.Map
                d_394_v71_ = _dafny.Map({len(d_393_v70_): (_dafny.SeqWithoutIsStrInference([d_386_v64_ for d_395_i6_ in range(default__.abs(-240))])) + (d_393_v70_)})
                d_394_v71_ = (d_394_v71_) | (d_394_v71_)
                d_396_v72_: _dafny.Set
                d_396_v72_ = _dafny.Set({True, False})
                d_397_v73_: _dafny.Map
                d_397_v73_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xas")): len((_dafny.Set({d_344_v27_})).intersection(d_396_v72_))})
                d_397_v73_ = ((_dafny.Map({d_393_v70_: d_385_v63_})) | ((d_397_v73_).set(d_393_v70_, d_390_v67_.f13)) if True else default__.fm15(d_386_v64_, d_390_v67_.f13, p3, default__.fm0(d_344_v27_, _dafny.SeqWithoutIsStrInference([d_386_v64_ for d_398_i7_ in range(default__.abs(171))]), globalState), globalState))
                d_399_v74_: T0
                nw59_ = C2()
                nw59_.ctor__(d_385_v63_, d_387_v65_, d_344_v27_)
                d_399_v74_ = nw59_
                d_400_v75_: _dafny.Set
                d_400_v75_ = _dafny.Set({d_399_v74_})
                d_401_v76_: C1
                nw60_ = C1()
                nw60_.ctor__(_dafny.MultiSet([default__.fm16((_dafny.MultiSet([True])).cardinality, d_387_v65_, d_344_v27_, p3, globalState), _dafny.SeqWithoutIsStrInference([d_390_v67_.f13, d_390_v67_.f13])]), p3, d_386_v64_, (d_388_v66_)[default__.safeIndex((d_392_v69_)[default__.safeIndex(692, (d_392_v69_).length(0))], len(d_388_v66_))], (_dafny.Set({d_399_v74_, d_399_v74_})).issubset(d_400_v75_))
                d_401_v76_ = nw60_
                d_402_v77_: _dafny.Set
                d_402_v77_ = _dafny.Set({d_390_v67_})
                d_344_v27_ = (_dafny.Set({d_390_v67_})).issubset(d_402_v77_)
                d_403_v78_: _dafny.MultiSet
                d_403_v78_ = _dafny.MultiSet([not(d_399_v74_.f9)])
                d_403_v78_ = (_dafny.MultiSet([not((d_399_v74_).f8)])).intersection(d_403_v78_)
            elif True:
                (globalState).f5 = default__.safeDivisionInt(p0, ((d_392_v69_)[default__.safeIndex(692, (d_392_v69_).length(0))] if p3 else (0) - ((d_392_v69_)[default__.safeIndex(692, (d_392_v69_).length(0))])))
                index58_ = default__.safeIndex(37, (d_381_v61_).length(0))
                (d_381_v61_)[index58_] = not (d_387_v65_) or ((self).fm3(globalState))
                index59_ = default__.safeIndex(37, (d_381_v61_).length(0))
                (d_381_v61_)[index59_] = p3
                d_404_v79_: _dafny.Seq
                d_404_v79_ = _dafny.SeqWithoutIsStrInference([p2, d_385_v63_])
                d_405_v80_: _dafny.MultiSet
                d_405_v80_ = _dafny.MultiSet([d_404_v79_])
                d_406_v81_: _dafny.Seq
                d_406_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eopmhpt"))
                d_407_v82_: _dafny.MultiSet
                d_407_v82_ = _dafny.MultiSet([d_406_v81_])
                d_408_v83_: _dafny.Map
                d_408_v83_ = _dafny.Map({(d_392_v69_)[default__.safeIndex(692, (d_392_v69_).length(0))]: -153})
                d_409_v84_: C1
                nw61_ = C1()
                nw61_.ctor__(d_405_v80_, (d_407_v82_).issubset(d_407_v82_), d_386_v64_, (p2) <= (len(d_408_v83_)), p3)
                d_409_v84_ = nw61_
                d_410_v85_: _dafny.MultiSet
                d_410_v85_ = _dafny.MultiSet([d_387_v65_])
                d_411_v86_: _dafny.Map
                d_411_v86_ = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_387_v65_]))) | (d_410_v85_): d_392_v69_})
                d_411_v86_ = (d_411_v86_).set(default__.fm19(d_388_v66_, (d_381_v61_)[default__.safeIndex(37, (d_381_v61_).length(0))], globalState), d_392_v69_)
                d_344_v27_ = not((d_390_v67_).fm5((d_388_v66_) != (d_388_v66_), globalState))
        d_412_v87_: _dafny.Seq
        d_412_v87_ = _dafny.SeqWithoutIsStrInference([p0, d_385_v63_, default__.fm0(d_344_v27_, _dafny.SeqWithoutIsStrInference([d_386_v64_ for d_413_i8_ in range(default__.abs(950))]), globalState), d_385_v63_, len(_dafny.Map({p3: d_385_v63_}))])
        d_414_v88_: D4
        d_414_v88_ = D4_DC7(d_412_v87_)
        r0 = (0) - ((0) - ((len((d_414_v88_).cf15) if p3 else (_dafny.MultiSet([d_386_v64_])).cardinality)))
        r1 = p0
        return r0, r1


class C4:
    def  __init__(self):
        self._f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f7):
        (self)._f7 = f7

    def fm1(self, p0, p1, globalState):
        return not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "biuajrf"))) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnja"))))

    def m1(self, p0, p1, p2, p3, globalState):
        d_415_v0_: str
        d_415_v0_ = _dafny.CodePoint('j')
        d_415_v0_ = d_415_v0_
        hi2_ = (0) - (p2)
        for d_416_i0_ in range(198, hi2_):
            d_417_v1_: _dafny.Array
            def lambda22_(d_418_p3_, d_419_p0_):
                def lambda23_(d_420_i1_):
                    return _dafny.Map({not(d_418_p3_): d_419_p0_})

                return lambda23_

            init14_ = lambda22_(p3, p0)
            nw62_ = _dafny.Array(None, 18)
            for i0_14_ in range(nw62_.length(0)):
                nw62_[i0_14_] = init14_(i0_14_)
            d_417_v1_ = nw62_
            d_421_v2_: _dafny.Map
            d_421_v2_ = _dafny.Map({p0: p3})
            d_422_v3_: _dafny.Seq
            d_422_v3_ = _dafny.SeqWithoutIsStrInference([p3])
            d_423_v4_: _dafny.Array
            nw63_ = _dafny.Array(None, 13)
            nw63_[int(0)] = p3
            nw63_[int(1)] = p3
            nw63_[int(2)] = ((d_421_v2_)[len(d_422_v3_)] if (len(d_422_v3_)) in (d_421_v2_) else not(p3))
            nw63_[int(3)] = p3
            nw63_[int(4)] = p3
            nw63_[int(5)] = p3
            nw63_[int(6)] = p3
            nw63_[int(7)] = p3
            nw63_[int(8)] = p3
            nw63_[int(9)] = p3
            nw63_[int(10)] = p3
            nw63_[int(11)] = False
            nw63_[int(12)] = p3
            d_423_v4_ = nw63_
            d_424_v5_: _dafny.MultiSet
            d_424_v5_ = _dafny.MultiSet([d_423_v4_, d_423_v4_, d_423_v4_, d_423_v4_, d_423_v4_])
            rhs48_ = d_417_v1_
            rhs49_ = d_424_v5_
            rhs50_ = d_416_i0_
            lhs27_ = globalState
            d_417_v1_ = rhs48_
            d_424_v5_ = rhs49_
            lhs27_.f5 = rhs50_
            if p3:
                d_425_v6_: _dafny.Map
                d_425_v6_ = _dafny.Map({not(p3): p0})
                d_426_v7_: _dafny.Set
                d_426_v7_ = _dafny.Set({(0) - (p0), p2})
                d_425_v6_ = (d_425_v6_).set(p3, len(d_426_v7_))
                d_427_v8_: bool
                d_427_v8_ = False
                d_427_v8_ = (d_422_v3_)[default__.safeIndex(default__.safeModuloInt((0) - (p2), p0), len(d_422_v3_))]
                index60_ = default__.safeIndex(372, (d_423_v4_).length(0))
                (d_423_v4_)[index60_] = True
                index61_ = default__.safeIndex(372, (d_423_v4_).length(0))
                (d_423_v4_)[index61_] = (d_427_v8_ if p3 else p3)
                d_428_v9_: _dafny.Array
                nw64_ = _dafny.Array(int(0), 27)
                d_428_v9_ = nw64_
                d_429_v10_: _dafny.Map
                d_429_v10_ = _dafny.Map({p0: d_428_v9_})
                d_430_v11_: D1
                d_430_v11_ = D1_DC1(((self).f7).set(default__.safeIndex(p1, len((self).f7)), ((self).f7)[default__.safeIndex(d_416_i0_, len((self).f7))]))
                d_431_v12_: _dafny.MultiSet
                d_431_v12_ = _dafny.MultiSet([d_427_v8_])
                d_432_v13_: _dafny.Seq
                d_432_v13_ = _dafny.SeqWithoutIsStrInference([p0])
                d_433_v14_: _dafny.Seq
                d_433_v14_ = _dafny.SeqWithoutIsStrInference([d_422_v3_])
                d_434_v15_: _dafny.Seq
                d_434_v15_ = _dafny.SeqWithoutIsStrInference([(d_433_v14_)[default__.safeIndex(p2, len(d_433_v14_))], d_422_v3_])
                d_435_v16_: _dafny.Array
                nw65_ = _dafny.Array(None, 22)
                nw65_[int(0)] = default__.safeDivisionInt(p0, len(d_429_v10_))
                nw65_[int(1)] = len((self).f7)
                nw65_[int(2)] = default__.safeModuloInt(p1, d_416_i0_)
                nw65_[int(3)] = len(((self).f7) + ((d_430_v11_).cf5))
                nw65_[int(4)] = p2
                nw65_[int(5)] = (p1) + (default__.fm0(True, _dafny.SeqWithoutIsStrInference([d_415_v0_ for d_436_i2_ in range(default__.abs(-470))]), globalState))
                nw65_[int(6)] = d_416_i0_
                nw65_[int(7)] = p1
                nw65_[int(8)] = (0) - (default__.fm0(p3, (self).f7, globalState))
                nw65_[int(9)] = (0) - ((d_416_i0_) + (p0))
                nw65_[int(10)] = p2
                nw65_[int(11)] = d_416_i0_
                nw65_[int(12)] = 748
                nw65_[int(13)] = p0
                nw65_[int(14)] = p2
                nw65_[int(15)] = ((d_431_v12_)[p3] if (p3) in (d_431_v12_) else len((d_432_v13_).set(default__.safeIndex(len(d_421_v2_), len(d_432_v13_)), (0) - ((d_431_v12_).cardinality))))
                nw65_[int(16)] = p2
                nw65_[int(17)] = d_416_i0_
                nw65_[int(18)] = d_416_i0_
                nw65_[int(19)] = len((self).f7)
                nw65_[int(20)] = len(d_434_v15_)
                nw65_[int(21)] = default__.fm0(not(not(d_427_v8_)), (self).f7, globalState)
                d_435_v16_ = nw65_
                d_435_v16_ = d_428_v9_
                d_437_v17_: _dafny.Array
                nw66_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                d_437_v17_ = nw66_
                index62_ = default__.safeIndex(21, (d_437_v17_).length(0))
                (d_437_v17_)[index62_] = (self).f7
                d_438_v18_: _dafny.Map
                d_438_v18_ = _dafny.Map({((self).f7) + ((self).f7): (self).f7})
                index63_ = default__.safeIndex(21, (d_437_v17_).length(0))
                rhs51_ = (0) - (p1)
                rhs52_ = ((self).f7).set(default__.safeIndex(d_416_i0_, len((self).f7)), ((self).f7)[default__.safeIndex(p2, len((self).f7))])
                rhs53_ = (d_422_v3_) + (d_422_v3_)
                rhs54_ = d_427_v8_
                rhs55_ = (d_438_v18_) | (default__.fm2(p2, (d_423_v4_)[default__.safeIndex(372, (d_423_v4_).length(0))], p3, globalState))
                lhs28_ = globalState
                lhs29_ = d_437_v17_
                lhs30_ = default__.safeIndex(21, (d_437_v17_).length(0))
                lhs28_.f5 = rhs51_
                lhs29_[lhs30_] = rhs52_
                d_422_v3_ = rhs53_
                d_427_v8_ = rhs54_
                d_438_v18_ = rhs55_
            elif True:
                d_439_v19_: bool
                d_439_v19_ = True
                d_439_v19_ = not(d_439_v19_)
                d_440_v20_: _dafny.Array
                nw67_ = _dafny.Array(int(0), 23)
                d_440_v20_ = nw67_
                index64_ = default__.safeIndex(534, (d_440_v20_).length(0))
                (d_440_v20_)[index64_] = (0) - (p2)
                index65_ = default__.safeIndex(534, (d_440_v20_).length(0))
                (d_440_v20_)[index65_] = default__.safeDivisionInt(p1, (0) - ((d_416_i0_) + (len((self).f7))))
                index66_ = default__.safeIndex(534, (d_440_v20_).length(0))
                (d_440_v20_)[index66_] = (d_440_v20_)[default__.safeIndex(534, (d_440_v20_).length(0))]
                d_441_v21_: _dafny.MultiSet
                d_441_v21_ = _dafny.MultiSet([d_439_v19_])
                index67_ = default__.safeIndex(508, (d_423_v4_).length(0))
                (d_423_v4_)[index67_] = (d_439_v19_) in (d_441_v21_)
                d_442_v22_: D2
                d_442_v22_ = D2_DC3(p3)
                index68_ = default__.safeIndex(508, (d_423_v4_).length(0))
                def iife23_(_pat_let5_0):
                    def iife24_(d_443_dt__update__tmp_h0_):
                        def iife25_(_pat_let6_0):
                            def iife26_(d_444_dt__update_hcf9_h0_):
                                return D2_DC3(d_444_dt__update_hcf9_h0_)
                            return iife26_(_pat_let6_0)
                        return iife25_(True)
                    return iife24_(_pat_let5_0)
                (d_423_v4_)[index68_] = (iife23_(d_442_v22_)).cf9
                index69_ = default__.safeIndex(508, (d_423_v4_).length(0))
                (d_423_v4_)[index69_] = (p2) >= ((len(_dafny.SeqWithoutIsStrInference([not((self).fm1(d_439_v19_, d_439_v19_, globalState))]))) + (len((self).f7)))
            d_445_v23_: _dafny.Array
            nw68_ = _dafny.Array(int(0), 25)
            d_445_v23_ = nw68_
            index70_ = default__.safeIndex(908, (d_445_v23_).length(0))
            (d_445_v23_)[index70_] = 873
            index71_ = default__.safeIndex(908, (d_445_v23_).length(0))
            def iife27_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(397, 852):
                    d_446_v24_: int = compr_13_
                    if ((397) <= (d_446_v24_)) and ((d_446_v24_) < (852)):
                        coll13_[(d_446_v24_) - (p1)] = (d_422_v3_)[default__.safeIndex(p1, len(d_422_v3_))]
                return _dafny.Map(coll13_)
            (d_445_v23_)[index71_] = (0) - ((default__.safeModuloInt(p0, len((self).f7)) if p3 else len(_dafny.Map({iife27_()
            : p2}))))
            index72_ = default__.safeIndex(908, (d_445_v23_).length(0))
            index73_ = default__.safeIndex(908, (d_445_v23_).length(0))
            index74_ = default__.safeIndex(908, (d_445_v23_).length(0))
            rhs56_ = (0) - ((d_445_v23_)[default__.safeIndex(908, (d_445_v23_).length(0))])
            rhs57_ = (d_445_v23_)[default__.safeIndex(908, (d_445_v23_).length(0))]
            rhs58_ = d_416_i0_
            rhs59_ = p2
            lhs31_ = globalState
            lhs32_ = d_445_v23_
            lhs33_ = default__.safeIndex(908, (d_445_v23_).length(0))
            lhs34_ = d_445_v23_
            lhs35_ = default__.safeIndex(908, (d_445_v23_).length(0))
            lhs36_ = d_445_v23_
            lhs37_ = default__.safeIndex(908, (d_445_v23_).length(0))
            lhs31_.f5 = rhs56_
            lhs32_[lhs33_] = rhs57_
            lhs34_[lhs35_] = rhs58_
            lhs36_[lhs37_] = rhs59_
        d_447_v25_: _dafny.Array
        nw69_ = _dafny.Array(_dafny.Map({}), 12)
        d_447_v25_ = nw69_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_447_v25_).length(0)):
            d_448_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_448_i3_)) and ((d_448_i3_) < ((d_447_v25_).length(0)))):
                def iife28_():
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, p3])), p0, p2])).Elements:
                        d_449_v26_: int = compr_14_
                        if (d_449_v26_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, p3])), p0, p2])):
                            coll14_[(d_449_v26_) - (p2)] = p3
                    return _dafny.Map(coll14_)
                (d_447_v25_)[(d_448_i3_)] = (iife28_()
                ) | (_dafny.Map({p0: p3}))
        (globalState).f5 = p1
        if False:
            d_450_v27_: C0
            nw70_ = C0()
            nw70_.ctor__(default__.safeDivisionInt(p1, p1), _dafny.CodePoint('g'), ((self).fm1(p3, p3, globalState)) and (p3), (self).fm1(p3, p3, globalState))
            d_450_v27_ = nw70_
            if p3:
                (globalState).f5 = (d_450_v27_.f13) + (len(_dafny.SeqWithoutIsStrInference([p1 for d_451_i4_ in range(default__.abs(-130))])))
                d_452_v28_: bool
                d_452_v28_ = True
                d_452_v28_ = d_452_v28_
                d_453_v29_: _dafny.Seq
                d_453_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmevn"))
                d_453_v29_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njdcm"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njdcm")))), default__.fm13(p1, globalState))
                d_454_v30_: _dafny.MultiSet
                d_454_v30_ = _dafny.MultiSet([d_415_v0_, d_415_v0_])
                d_455_v31_: _dafny.Map
                d_455_v31_ = _dafny.Map({p1: p3})
                d_456_v32_: _dafny.Set
                d_456_v32_ = _dafny.Set({_dafny.Map({p0: True}), d_455_v31_})
                (d_450_v27_).f13 = (((d_454_v30_).cardinality if p3 else len(d_456_v32_))) + (default__.safeModuloInt(d_450_v27_.f13, (0) - (p0)))
                d_457_v33_: _dafny.Array
                nw71_ = _dafny.Array(_dafny.Array(None, 0), 13)
                d_457_v33_ = nw71_
                d_458_v34_: _dafny.Map
                d_458_v34_ = _dafny.Map({d_457_v33_: d_452_v28_})
                d_459_v35_: _dafny.Seq
                d_459_v35_ = _dafny.SeqWithoutIsStrInference([d_452_v28_])
                d_460_v36_: C2
                nw72_ = C2()
                nw72_.ctor__(p2, True, ((d_458_v34_)[d_457_v33_] if (d_457_v33_) in (d_458_v34_) else (d_459_v35_)[default__.safeIndex(p2, len(d_459_v35_))]))
                d_460_v36_ = nw72_
            elif True:
                d_461_v37_: _dafny.Seq
                d_461_v37_ = _dafny.SeqWithoutIsStrInference([d_450_v27_.f13])
                d_461_v37_ = _dafny.SeqWithoutIsStrInference([d_450_v27_.f13, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([p3])), d_450_v27_.f13), p1])
                d_462_v38_: bool
                d_462_v38_ = True
                d_462_v38_ = d_462_v38_
                d_462_v38_ = p3
                d_463_v39_: _dafny.Array
                def lambda24_(d_464_v27_):
                    def lambda25_(d_465_i5_):
                        return (d_465_i5_) * (d_464_v27_.f13)

                    return lambda25_

                init15_ = lambda24_(d_450_v27_)
                nw73_ = _dafny.Array(None, 20)
                for i0_15_ in range(nw73_.length(0)):
                    nw73_[i0_15_] = init15_(i0_15_)
                d_463_v39_ = nw73_
                index75_ = default__.safeIndex(553, (d_463_v39_).length(0))
                (d_463_v39_)[index75_] = (0) - (p1)
                d_466_v40_: _dafny.Set
                d_466_v40_ = _dafny.Set({d_462_v38_, not(p3), d_462_v38_, d_462_v38_})
                index76_ = default__.safeIndex(553, (d_463_v39_).length(0))
                (d_463_v39_)[index76_] = (len((d_466_v40_ if p3 else d_466_v40_))) * (default__.fm0(p3, (self).f7, globalState))
                d_467_v41_: _dafny.Array
                nw74_ = _dafny.Array(_dafny.Map({}), 19)
                d_467_v41_ = nw74_
                d_467_v41_ = d_467_v41_
            d_468_v42_: C2
            nw75_ = C2()
            nw75_.ctor__(-178, True, p3)
            d_468_v42_ = nw75_
            d_469_v43_: _dafny.Map
            d_469_v43_ = _dafny.Map({p3: d_468_v42_})
            d_470_v44_: _dafny.Map
            d_470_v44_ = _dafny.Map({(d_469_v43_ if p3 else d_469_v43_): p3})
            d_471_v45_: _dafny.Set
            d_471_v45_ = _dafny.Set({False, p3})
            d_470_v44_ = (d_470_v44_).set(d_469_v43_, (d_471_v45_).ispropersubset(d_471_v45_))
            d_415_v0_ = d_415_v0_
            d_472_v46_: _dafny.Array
            def lambda26_(d_473_v27_, d_474_p3_):
                def lambda27_(d_475_i6_):
                    return default__.safeModuloInt(d_475_i6_, len(_dafny.Map({(0) - (d_473_v27_.f13): _dafny.SeqWithoutIsStrInference([d_474_p3_])})))

                return lambda27_

            init16_ = lambda26_(d_450_v27_, p3)
            nw76_ = _dafny.Array(None, 4)
            for i0_16_ in range(nw76_.length(0)):
                nw76_[i0_16_] = init16_(i0_16_)
            d_472_v46_ = nw76_
            index77_ = default__.safeIndex(448, (d_472_v46_).length(0))
            (d_472_v46_)[index77_] = p1
            index78_ = default__.safeIndex(448, (d_472_v46_).length(0))
            (d_472_v46_)[index78_] = p0
        elif True:
            if (p2) != ((0) - (p1)):
                d_476_v47_: bool
                d_476_v47_ = False
                d_477_v48_: _dafny.Seq
                d_477_v48_ = _dafny.SeqWithoutIsStrInference([p3, p3])
                d_476_v47_ = (d_476_v47_) or ((d_477_v48_)[default__.safeIndex(len((self).f7), len(d_477_v48_))])
                d_476_v47_ = d_476_v47_
                d_478_v49_: _dafny.Set
                d_478_v49_ = _dafny.Set({(self).fm1(d_476_v47_, d_476_v47_, globalState)})
                d_479_v50_: D1
                d_479_v50_ = D1_DC2(len(d_478_v49_), p2, _dafny.CodePoint('c'))
                d_480_v51_: _dafny.MultiSet
                d_480_v51_ = _dafny.MultiSet([D1_DC2(p2, 646, d_415_v0_)])
                d_476_v47_ = (_dafny.MultiSet([d_479_v50_, d_479_v50_])).isdisjoint((d_480_v51_) | (_dafny.MultiSet([d_479_v50_])))
                d_476_v47_ = d_476_v47_
                d_481_v52_: C3
                nw77_ = C3()
                nw77_.ctor__()
                d_481_v52_ = nw77_
            elif True:
                (globalState).f5 = p0
                d_482_v53_: _dafny.Map
                d_482_v53_ = _dafny.Map({p2: (_dafny.CodePoint('o') if p3 else d_415_v0_)})
                d_483_v54_: _dafny.Seq
                d_483_v54_ = _dafny.SeqWithoutIsStrInference([p3, p3])
                d_484_v55_: _dafny.Seq
                d_484_v55_ = _dafny.SeqWithoutIsStrInference([p2])
                d_485_v56_: _dafny.Map
                d_485_v56_ = _dafny.Map({(d_484_v55_)[default__.safeIndex(p2, len(d_484_v55_))]: p3})
                d_486_v57_: _dafny.Array
                nw78_ = _dafny.Array(None, 24)
                nw78_[int(0)] = len(d_485_v56_)
                nw78_[int(1)] = p0
                nw78_[int(2)] = p2
                nw78_[int(3)] = p0
                nw78_[int(4)] = p2
                nw78_[int(5)] = p1
                nw78_[int(6)] = p0
                nw78_[int(7)] = len((self).f7)
                nw78_[int(8)] = p0
                nw78_[int(9)] = p0
                nw78_[int(10)] = p0
                nw78_[int(11)] = p2
                nw78_[int(12)] = p0
                nw78_[int(13)] = p0
                nw78_[int(14)] = p1
                nw78_[int(15)] = len(default__.fm20(p3, globalState))
                nw78_[int(16)] = p2
                nw78_[int(17)] = p1
                nw78_[int(18)] = p0
                nw78_[int(19)] = p2
                nw78_[int(20)] = p1
                nw78_[int(21)] = p1
                nw78_[int(22)] = p2
                nw78_[int(23)] = p2
                d_486_v57_ = nw78_
                d_487_v58_: _dafny.Set
                d_487_v58_ = _dafny.Set({d_486_v57_, d_486_v57_, d_486_v57_})
                d_482_v53_ = (d_482_v53_).set((D2_DC4(_dafny.SeqWithoutIsStrInference([d_415_v0_ for d_488_i7_ in range(default__.abs(839))]), d_483_v54_, len(d_487_v58_))).cf12, (d_415_v0_ if p3 else d_415_v0_))
                (globalState).f5 = (len((self).f7)) + ((default__.fm0(p3, (self).f7, globalState) if False else 711))
                d_489_v59_: D1
                d_489_v59_ = D1_DC1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))
                d_490_v60_: _dafny.Seq
                d_490_v60_ = _dafny.SeqWithoutIsStrInference([(self).f7])
                d_489_v59_ = D1_DC1((d_490_v60_)[default__.safeIndex(p2, len(d_490_v60_))])
                (globalState).f5 = default__.safeModuloInt(default__.safeModuloInt(p0, p1), (len((self).f7)) - (p0))
            (globalState).f5 = 697
            d_491_v61_: _dafny.Map
            d_491_v61_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbjxfiomp")): (self).f7})
            d_492_v62_: _dafny.Set
            d_492_v62_ = _dafny.Set({p3, True})
            d_493_v63_: _dafny.Seq
            d_493_v63_ = _dafny.SeqWithoutIsStrInference([p2, p2])
            d_494_v64_: _dafny.MultiSet
            d_494_v64_ = _dafny.MultiSet([d_493_v63_, _dafny.SeqWithoutIsStrInference([p0])])
            d_495_v65_: C1
            nw79_ = C1()
            nw79_.ctor__(d_494_v64_, p3, d_415_v0_, p3, p3)
            d_495_v65_ = nw79_
            d_496_v66_: D4
            d_496_v66_ = D4_DC8(p2, p3, p0, d_495_v65_)
            d_497_v67_: _dafny.Array
            nw80_ = _dafny.Array(None, 14)
            nw80_[int(0)] = len(d_492_v62_)
            nw80_[int(1)] = p0
            nw80_[int(2)] = p1
            nw80_[int(3)] = p2
            nw80_[int(4)] = p2
            nw80_[int(5)] = (0) - (p2)
            nw80_[int(6)] = p0
            nw80_[int(7)] = p2
            nw80_[int(8)] = (0) - (p1)
            nw80_[int(9)] = p0
            nw80_[int(10)] = (0) - ((d_496_v66_).cf18)
            nw80_[int(11)] = 508
            nw80_[int(12)] = p2
            nw80_[int(13)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ll")))
            d_497_v67_ = nw80_
            d_498_v68_: _dafny.Map
            d_498_v68_ = _dafny.Map({p0: d_495_v65_})
            d_499_v69_: _dafny.Array
            def lambda28_(d_500_v0_):
                def lambda29_(d_501_i8_):
                    return d_500_v0_

                return lambda29_

            init17_ = lambda28_(d_415_v0_)
            nw81_ = _dafny.Array(None, 7)
            for i0_17_ in range(nw81_.length(0)):
                nw81_[i0_17_] = init17_(i0_17_)
            d_499_v69_ = nw81_
            d_502_v70_: D0
            d_502_v70_ = D0_DC0(d_491_v61_, d_497_v67_, _dafny.Map({(d_495_v65_).f12: len(d_498_v68_)}), d_499_v69_, d_497_v67_)
            d_503_v71_: _dafny.Map
            d_503_v71_ = _dafny.Map({p3: d_502_v70_})
            d_504_v72_: _dafny.Array
            nw82_ = _dafny.Array(None, 7)
            nw82_[int(0)] = d_503_v71_
            nw82_[int(1)] = _dafny.Map({p3: d_502_v70_})
            nw82_[int(2)] = (d_503_v71_).set(p3, d_502_v70_)
            nw82_[int(3)] = d_503_v71_
            nw82_[int(4)] = d_503_v71_
            nw82_[int(5)] = d_503_v71_
            nw82_[int(6)] = (d_503_v71_) | (d_503_v71_)
            d_504_v72_ = nw82_
            d_504_v72_ = d_504_v72_
            d_505_v73_: _dafny.Map
            d_505_v73_ = _dafny.Map({not(p3): p2})
            d_506_v74_: _dafny.Set
            d_506_v74_ = _dafny.Set({d_505_v73_})
            d_507_v75_: _dafny.MultiSet
            d_507_v75_ = _dafny.MultiSet([False])
            d_508_v76_: int
            out12_: int
            out12_ = default__.m0(not ((d_495_v65_).f12) or (p3), (d_506_v74_) - (d_506_v74_), (p3) not in (d_507_v75_), globalState)
            d_508_v76_ = out12_
            d_509_v77_: bool
            d_509_v77_ = True
            d_509_v77_ = (self).fm1((d_508_v76_) in (_dafny.SeqWithoutIsStrInference([607])), (d_495_v65_).f12, globalState)
        d_510_v78_: _dafny.Map
        d_510_v78_ = _dafny.Map({p2: p2})
        d_510_v78_ = (d_510_v78_).set((0) - (p2), p1)

    @property
    def f7(self):
        return self._f7
