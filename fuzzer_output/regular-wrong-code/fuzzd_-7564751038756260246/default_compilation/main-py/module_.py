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
        return (D7_DC16(163, 125, True, False, 276)).cf26

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        return ((_dafny.Set({387})).intersection(_dafny.Set({-805}))) | ((_dafny.Set({624, (D6_DC13(510)).cf19})))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        return _dafny.CodePoint('f')

    @staticmethod
    def fm3(globalState):
        return ((_dafny.SeqWithoutIsStrInference([True])) + (_dafny.SeqWithoutIsStrInference([True]))) + (_dafny.SeqWithoutIsStrInference([True, False]))

    @staticmethod
    def fm4(p0, p1, globalState):
        return 549

    @staticmethod
    def fm5(p0, p1, globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([False, True, not(True), not(False), False])): (0) - ((0) - (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([True])), 192)))})

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)])))

    @staticmethod
    def fm7(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(525, 56):
                d_0_v0_: int = compr_0_
                if ((525) <= (d_0_v0_)) and ((d_0_v0_) < (56)):
                    coll0_[(d_0_v0_) - ((_dafny.MultiSet([False])).cardinality)] = not(not(False))
            return _dafny.Map(coll0_)
        return (_dafny.Map({315: False})) | (iife0_()
        )

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drfsgt"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtfucath"))))

    @staticmethod
    def fm9(globalState):
        if False:
            return D3_DC4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_1_i0_ in range(default__.abs(864))]))
        elif True:
            return D3_DC4(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_2_i1_ in range(default__.abs(-403))]))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        source0_ = (D7_DC15((0) - ((0) - (len(_dafny.Map({983: len(_dafny.Map({False: True}))})))), len(_dafny.Map({False: _dafny.MultiSet([True])}))) if False else D7_DC15(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(False)])])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ujipcla")))))
        if source0_.is_DC15:
            d_3___mcc_h0_ = source0_.cf21
            d_4___mcc_h1_ = source0_.cf22
            d_5_cf22_ = d_4___mcc_h1_
            d_6_cf21_ = d_3___mcc_h0_
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: _dafny.Map
                for compr_1_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htgdju"))): d_6_cf21_}), _dafny.Map({d_6_cf21_: d_5_cf22_}), _dafny.Map({d_6_cf21_: d_6_cf21_})]))).Elements:
                    d_7_v0_: _dafny.Map = compr_1_
                    if (d_7_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htgdju"))): d_6_cf21_}), _dafny.Map({d_6_cf21_: d_5_cf22_}), _dafny.Map({d_6_cf21_: d_6_cf21_})]))):
                        coll1_[d_7_v0_] = True
                return _dafny.Map(coll1_)
            return iife1_()
            
        elif source0_.is_DC16:
            d_8___mcc_h2_ = source0_.cf23
            d_9___mcc_h3_ = source0_.cf24
            d_10___mcc_h4_ = source0_.cf25
            d_11___mcc_h5_ = source0_.cf26
            d_12___mcc_h6_ = source0_.cf27
            d_13_cf27_ = d_12___mcc_h6_
            d_14_cf26_ = d_11___mcc_h5_
            d_15_cf25_ = d_10___mcc_h4_
            d_16_cf24_ = d_9___mcc_h3_
            d_17_cf23_ = d_8___mcc_h2_
            return _dafny.Map({_dafny.Map({d_17_cf23_: d_16_cf24_}): d_15_cf25_})
        elif source0_.is_DC14:
            d_18___mcc_h7_ = source0_.cf20
            d_19_cf20_ = d_18___mcc_h7_
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in (_dafny.Map({51: True})).keys.Elements:
                    d_20_v1_: int = compr_2_
                    if (d_20_v1_) in (_dafny.Map({51: True})):
                        coll2_[(d_20_v1_) * (len(_dafny.Set({False, False})))] = 420
                return _dafny.Map(coll2_)
            def iife3_():
                def iife6_():
                    coll6_ = _dafny.Map()
                    compr_6_: _dafny.Seq
                    for compr_6_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])).Elements:
                        d_21_v3_: _dafny.Seq = compr_6_
                        if (d_21_v3_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])):
                            coll6_[d_21_v3_] = -542
                    return _dafny.Map(coll6_)
                def iife7_():
                    coll7_ = _dafny.Map()
                    compr_7_: int
                    for compr_7_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_22_i0_ in range(default__.abs(672))]))])).Elements:
                        d_23_v4_: int = compr_7_
                        if (d_23_v4_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_22_i0_ in range(default__.abs(672))]))])):
                            coll7_[(d_23_v4_) + (405)] = 190
                    return _dafny.Map(coll7_)
                coll3_ = _dafny.Map()
                def iife4_():
                    coll4_ = _dafny.Map()
                    compr_4_: _dafny.Seq
                    for compr_4_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])).Elements:
                        d_21_v3_: _dafny.Seq = compr_4_
                        if (d_21_v3_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True, True])])):
                            coll4_[d_21_v3_] = -542
                    return _dafny.Map(coll4_)
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: int
                    for compr_5_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_22_i0_ in range(default__.abs(672))]))])).Elements:
                        d_23_v4_: int = compr_5_
                        if (d_23_v4_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_22_i0_ in range(default__.abs(672))]))])):
                            coll5_[(d_23_v4_) + (405)] = 190
                    return _dafny.Map(coll5_)
                compr_3_: int
                for compr_3_ in (_dafny.MultiSet([len(iife4_()
                ), len(iife5_()
                ), len(_dafny.SeqWithoutIsStrInference([83 for d_24_i1_ in range(default__.abs(344))])), 350])).Elements:
                    d_25_v2_: int = compr_3_
                    if (d_25_v2_) in (_dafny.MultiSet([len(iife6_()
                    ), len(iife7_()
                    ), len(_dafny.SeqWithoutIsStrInference([83 for d_24_i1_ in range(default__.abs(344))])), 350])):
                        coll3_[default__.safeDivisionInt(d_25_v2_, len(_dafny.Set({not(True)})))] = (0) - (len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([True, False, True]))})))
                return _dafny.Map(coll3_)
            return (_dafny.Map({iife2_()
            : False})) | (_dafny.Map({iife3_()
            : not(True)}))
        elif True:
            d_26___mcc_h8_ = source0_.cf28
            d_27_cf28_ = d_26___mcc_h8_
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: _dafny.Map
                for compr_8_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.Set({658, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_28_i2_ in range(default__.abs(528))])), -688, -524, -555}))}), _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvgbw"))), len(_dafny.Map({True: 394}))]))).cardinality: -341})])).Elements:
                    d_29_v5_: _dafny.Map = compr_8_
                    if (d_29_v5_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.Set({658, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_28_i2_ in range(default__.abs(528))])), -688, -524, -555}))}), _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvgbw"))), len(_dafny.Map({True: 394}))]))).cardinality: -341})])):
                        coll8_[d_29_v5_] = False
                return _dafny.Map(coll8_)
            return iife8_()
            

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in _dafny.IntegerRange(492, 841):
                d_30_v0_: int = compr_9_
                if ((492) <= (d_30_v0_)) and ((d_30_v0_) < (841)):
                    coll9_ = coll9_.union(_dafny.Set([(d_30_v0_) - ((0) - (len(_dafny.SeqWithoutIsStrInference([256]))))]))
            return _dafny.Set(coll9_)
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, False])).cardinality])).Elements:
                d_32_v1_: int = compr_10_
                if (d_32_v1_) in (_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False, False])).cardinality])):
                    coll10_[(d_32_v1_) + (445)] = len(_dafny.Map({True: 908}))
            return _dafny.Map(coll10_)
        return (_dafny.Map({False: (0) - (len(iife9_()
        ))})) | ((_dafny.Map({not(not(True)): len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpdwcbhu"))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_31_i0_ in range(default__.abs(893))]))}))})) | (_dafny.Map({False: len(iife10_()
        )})))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_33_i0_: int
        d_33_i0_ = 0
        with _dafny.label("0"):
            while p3:
                with _dafny.c_label("0"):
                    if (d_33_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_33_i0_ = (d_33_i0_) + (1)
                    d_34_v0_: _dafny.Map
                    d_34_v0_ = _dafny.Map({False: p0})
                    (globalState).f8 = len((d_34_v0_).set(p3, p0))
                    d_35_v1_: C0
                    nw0_ = C0()
                    nw0_.ctor__()
                    d_35_v1_ = nw0_
                    d_36_v2_: C0
                    nw1_ = C0()
                    nw1_.ctor__()
                    d_36_v2_ = nw1_
                    d_37_v3_: D3
                    d_37_v3_ = D3_DC5(p3, p1)
                    d_38_v4_: _dafny.Seq
                    d_38_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkerwni"))
                    d_39_v5_: _dafny.Set
                    d_39_v5_ = _dafny.Set({False, p0})
                    d_40_v6_: _dafny.Map
                    d_40_v6_ = _dafny.Map({d_36_v2_: p3})
                    d_41_v7_: _dafny.Seq
                    d_41_v7_ = _dafny.SeqWithoutIsStrInference([d_40_v6_])
                    d_42_v8_: _dafny.Array
                    nw2_ = _dafny.Array(False, 18)
                    d_42_v8_ = nw2_
                    d_43_v9_: _dafny.Map
                    d_43_v9_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([p0, p3])): d_42_v8_})
                    d_44_v10_: _dafny.Array
                    nw3_ = _dafny.Array(None, 21)
                    nw3_[int(0)] = p1
                    nw3_[int(1)] = default__.safeDivisionInt(p1, p1)
                    nw3_[int(2)] = (d_37_v3_).cf7
                    nw3_[int(3)] = p1
                    nw3_[int(4)] = (0) - (p1)
                    nw3_[int(5)] = p1
                    nw3_[int(6)] = p1
                    nw3_[int(7)] = len(d_38_v4_)
                    nw3_[int(8)] = len(d_39_v5_)
                    nw3_[int(9)] = p1
                    nw3_[int(10)] = len(d_38_v4_)
                    nw3_[int(11)] = p1
                    nw3_[int(12)] = 526
                    nw3_[int(13)] = p1
                    nw3_[int(14)] = p1
                    nw3_[int(15)] = default__.safeDivisionInt(p1, len((d_41_v7_).set(default__.safeIndex(p1, len(d_41_v7_)), d_40_v6_)))
                    nw3_[int(16)] = p1
                    nw3_[int(17)] = len(d_38_v4_)
                    nw3_[int(18)] = p1
                    nw3_[int(19)] = len((d_43_v9_) | (d_43_v9_))
                    nw3_[int(20)] = p1
                    d_44_v10_ = nw3_
                    index0_ = default__.safeIndex(786, (d_44_v10_).length(0))
                    (d_44_v10_)[index0_] = len((d_34_v0_).set(p3, True))
                    index1_ = default__.safeIndex(786, (d_44_v10_).length(0))
                    (d_44_v10_)[index1_] = p1
                    pass
            pass
        d_45_v11_: _dafny.Seq
        d_45_v11_ = _dafny.SeqWithoutIsStrInference([(0) - (p1)])
        d_46_v12_: _dafny.Seq
        d_46_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owdvbmnf"))
        d_47_v13_: C0
        nw4_ = C0()
        nw4_.ctor__()
        d_47_v13_ = nw4_
        d_48_v14_: _dafny.Seq
        d_48_v14_ = _dafny.SeqWithoutIsStrInference([d_47_v13_, d_47_v13_])
        d_49_v15_: _dafny.Seq
        d_49_v15_ = _dafny.SeqWithoutIsStrInference([d_48_v14_, d_48_v14_, d_48_v14_, _dafny.SeqWithoutIsStrInference([d_47_v13_, d_47_v13_, d_47_v13_])])
        d_50_v16_: _dafny.Map
        d_50_v16_ = _dafny.Map({(d_45_v11_)[default__.safeIndex(len(d_46_v12_), len(d_45_v11_))]: (d_48_v14_) in (d_49_v15_)})
        d_51_i1_: int
        d_51_i1_ = 0
        with _dafny.label("1"):
            while ((d_50_v16_)[p1] if (p1) in (d_50_v16_) else p3):
                with _dafny.c_label("1"):
                    if (d_51_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_51_i1_ = (d_51_i1_) + (1)
                    if not(not(not(p0))):
                        d_52_v17_: _dafny.Seq
                        d_52_v17_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                        d_53_v18_: _dafny.Map
                        d_53_v18_ = _dafny.Map({d_52_v17_: d_47_v13_})
                        d_54_v19_: _dafny.Map
                        d_54_v19_ = _dafny.Map({((d_53_v18_)[d_52_v17_] if (d_52_v17_) in (d_53_v18_) else d_47_v13_): (d_45_v11_) < (d_45_v11_)})
                        d_54_v19_ = _dafny.Map({d_47_v13_: p3})
                        (globalState).f8 = (default__.fm4(p3, p1, globalState)) * (p1)
                        d_55_v20_: _dafny.Array
                        def lambda0_(d_56_p3_, d_57_v11_, d_58_p1_):
                            def lambda1_(d_59_i2_):
                                return ((_dafny.MultiSet([d_56_p3_])).cardinality) < ((d_57_v11_)[default__.safeIndex(d_58_p1_, len(d_57_v11_))])

                            return lambda1_

                        init0_ = lambda0_(p3, d_45_v11_, p1)
                        nw5_ = _dafny.Array(None, 22)
                        for i0_0_ in range(nw5_.length(0)):
                            nw5_[i0_0_] = init0_(i0_0_)
                        d_55_v20_ = nw5_
                        d_60_v21_: _dafny.MultiSet
                        d_60_v21_ = _dafny.MultiSet([p3])
                        index2_ = default__.safeIndex(304, (d_55_v20_).length(0))
                        (d_55_v20_)[index2_] = (d_60_v21_) != (d_60_v21_)
                        d_61_v22_: _dafny.MultiSet
                        d_61_v22_ = _dafny.MultiSet([p1])
                        d_62_v23_: _dafny.Map
                        d_62_v23_ = _dafny.Map({-270: p1})
                        d_63_v24_: _dafny.Set
                        d_63_v24_ = _dafny.Set({357})
                        d_64_v25_: _dafny.Map
                        d_64_v25_ = _dafny.Map({d_47_v13_: d_63_v24_})
                        index3_ = default__.safeIndex(304, (d_55_v20_).length(0))
                        (d_55_v20_)[index3_] = ((0) - (((d_45_v11_)[default__.safeIndex((d_61_v22_).cardinality, len(d_45_v11_))]) - ((0) - (((d_62_v23_)[p1] if (p1) in (d_62_v23_) else len(d_64_v25_)))))) > (len(d_52_v17_))
                        (globalState).f8 = 182
                        d_65_v26_: _dafny.Set
                        d_65_v26_ = _dafny.Set({d_47_v13_, d_47_v13_})
                        d_66_v27_: _dafny.Map
                        d_66_v27_ = _dafny.Map({p1: d_65_v26_})
                        index4_ = default__.safeIndex(304, (d_55_v20_).length(0))
                        (d_55_v20_)[index4_] = (d_65_v26_).isdisjoint(((d_66_v27_)[p1] if (p1) in (d_66_v27_) else d_65_v26_))
                    elif True:
                        d_67_v28_: bool
                        d_67_v28_ = False
                        rhs0_ = (d_45_v11_)[default__.safeIndex((136) + (p1), len(d_45_v11_))]
                        rhs1_ = not(d_67_v28_)
                        rhs2_ = d_46_v12_
                        rhs3_ = (p2) in (d_46_v12_)
                        lhs0_ = globalState
                        lhs0_.f8 = rhs0_
                        d_67_v28_ = rhs1_
                        d_46_v12_ = rhs2_
                        d_67_v28_ = rhs3_
                        d_68_v29_: _dafny.Map
                        d_68_v29_ = _dafny.Map({not(p3): p0})
                        d_69_v30_: _dafny.Seq
                        d_69_v30_ = _dafny.SeqWithoutIsStrInference([d_68_v29_, _dafny.Map({p3: d_67_v28_})])
                        d_69_v30_ = ((d_69_v30_) + (d_69_v30_)) + (d_69_v30_)
                        d_70_v31_: _dafny.Map
                        d_70_v31_ = _dafny.Map({d_47_v13_: p2})
                        d_70_v31_ = (d_70_v31_).set(d_47_v13_, p2)
                        d_71_v32_: C0
                        nw6_ = C0()
                        nw6_.ctor__()
                        d_71_v32_ = nw6_
                        d_72_v33_: _dafny.MultiSet
                        d_72_v33_ = _dafny.MultiSet([p2, _dafny.CodePoint('q'), p2, _dafny.CodePoint('f'), p2])
                        d_73_v34_: _dafny.Set
                        d_73_v34_ = _dafny.Set({p1})
                        d_74_v35_: _dafny.Map
                        d_74_v35_ = _dafny.Map({p1: p1})
                        d_75_v36_: _dafny.Array
                        nw7_ = _dafny.Array(None, 19)
                        nw7_[int(0)] = p1
                        nw7_[int(1)] = 293
                        nw7_[int(2)] = p1
                        nw7_[int(3)] = ((d_72_v33_)[p2] if (p2) in (d_72_v33_) else p1)
                        nw7_[int(4)] = len((d_45_v11_) + (d_45_v11_))
                        nw7_[int(5)] = p1
                        nw7_[int(6)] = (0) - (default__.fm4(p0, p1, globalState))
                        nw7_[int(7)] = p1
                        nw7_[int(8)] = p1
                        nw7_[int(9)] = len(_dafny.Map({d_71_v32_: len(d_45_v11_)}))
                        nw7_[int(10)] = (0) - ((D3_DC5(False, 662)).cf7)
                        nw7_[int(11)] = p1
                        nw7_[int(12)] = 185
                        nw7_[int(13)] = (len(d_73_v34_)) + (p1)
                        nw7_[int(14)] = default__.safeModuloInt(p1, p1)
                        nw7_[int(15)] = len(_dafny.SeqWithoutIsStrInference([326 for d_76_i3_ in range(default__.abs(862))]))
                        nw7_[int(16)] = p1
                        nw7_[int(17)] = default__.safeModuloInt(-257, p1)
                        nw7_[int(18)] = len((d_74_v35_) | (_dafny.Map({p1: p1})))
                        d_75_v36_ = nw7_
                        index5_ = default__.safeIndex(249, (d_75_v36_).length(0))
                        (d_75_v36_)[index5_] = 506
                        d_77_v37_: _dafny.Map
                        d_77_v37_ = _dafny.Map({p0: d_46_v12_})
                        index6_ = default__.safeIndex(249, (d_75_v36_).length(0))
                        (d_75_v36_)[index6_] = (p1) + (default__.safeDivisionInt(15, len(d_77_v37_)))
                    d_78_v38_: bool
                    d_78_v38_ = False
                    d_78_v38_ = p0
                    d_79_v39_: C0
                    nw8_ = C0()
                    nw8_.ctor__()
                    d_79_v39_ = nw8_
                    d_80_v40_: _dafny.Array
                    nw9_ = _dafny.Array(int(0), 26)
                    d_80_v40_ = nw9_
                    d_81_v41_: _dafny.Map
                    d_81_v41_ = _dafny.Map({d_80_v40_: (d_46_v12_) + ((d_46_v12_).set(default__.safeIndex(p1, len(d_46_v12_)), _dafny.CodePoint('e')))})
                    d_82_v42_: D3
                    d_82_v42_ = D3_DC5(True, p1)
                    rhs4_ = (d_81_v41_) | (d_81_v41_)
                    rhs5_ = p3
                    rhs6_ = (d_82_v42_).cf7
                    rhs7_ = default__.fm4(d_78_v38_, (0) - (p1), globalState)
                    lhs1_ = globalState
                    lhs2_ = globalState
                    d_81_v41_ = rhs4_
                    d_78_v38_ = rhs5_
                    lhs1_.f0 = rhs6_
                    lhs2_.f0 = rhs7_
                    pass
            pass
        source1_ = default__.fm9(globalState)
        if source1_.is_DC5:
            d_83___mcc_h0_ = source1_.cf6
            d_84___mcc_h1_ = source1_.cf7
            d_85_cf7_ = d_84___mcc_h1_
            d_86_cf6_ = d_83___mcc_h0_
            d_87_v43_: C0
            nw10_ = C0()
            nw10_.ctor__()
            d_87_v43_ = nw10_
            d_88_v44_: str
            d_88_v44_ = _dafny.CodePoint('g')
            d_88_v44_ = (default__.fm2(_dafny.Map({p1: (0) - (p1)}), not(not(d_86_cf6_)), d_85_cf7_, globalState) if p0 else d_88_v44_)
            d_86_cf6_ = p0
            d_86_cf6_ = (p3 if not(d_86_cf6_) else p0)
        elif True:
            d_89___mcc_h2_ = source1_.cf5
            d_90_cf5_ = d_89___mcc_h2_
            d_91_v45_: _dafny.Map
            d_91_v45_ = _dafny.Map({d_45_v11_: d_90_cf5_})
            d_92_v46_: D6
            d_92_v46_ = D6_DC10(_dafny.SeqWithoutIsStrInference([p1]))
            d_91_v45_ = (d_91_v45_).set(((d_45_v11_).set(default__.safeIndex(len((d_92_v46_).cf14), len(d_45_v11_)), (d_45_v11_)[default__.safeIndex(p1, len(d_45_v11_))])).set(default__.safeIndex(p1, len((d_45_v11_).set(default__.safeIndex(len((d_92_v46_).cf14), len(d_45_v11_)), (d_45_v11_)[default__.safeIndex(p1, len(d_45_v11_))]))), p1), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imvdwrgjx")))
            d_93_v47_: C0
            nw11_ = C0()
            nw11_.ctor__()
            d_93_v47_ = nw11_
            (globalState).f0 = p1
            d_94_v48_: _dafny.Array
            nw12_ = _dafny.Array(int(0), 10)
            d_94_v48_ = nw12_
            index7_ = default__.safeIndex(289, (d_94_v48_).length(0))
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in _dafny.IntegerRange(400, 602):
                    d_95_v49_: int = compr_11_
                    if ((400) <= (d_95_v49_)) and ((d_95_v49_) < (602)):
                        coll11_[default__.safeDivisionInt(d_95_v49_, p1)] = p2
                return _dafny.Map(coll11_)
            (d_94_v48_)[index7_] = (len(d_90_cf5_)) + (len(iife11_()
            ))
            index8_ = default__.safeIndex(289, (d_94_v48_).length(0))
            (d_94_v48_)[index8_] = p1
        d_96_v50_: _dafny.MultiSet
        d_96_v50_ = _dafny.MultiSet([p1])
        d_97_v51_: _dafny.Seq
        d_97_v51_ = _dafny.SeqWithoutIsStrInference([(not(p3) if not(p3) else p0), (d_96_v50_).ispropersubset(d_96_v50_)])
        d_97_v51_ = d_97_v51_
        d_98_v52_: C0
        nw13_ = C0()
        nw13_.ctor__()
        d_98_v52_ = nw13_
        hi0_ = (0) - ((p1) + (p1))
        for d_99_i4_ in range(p1, hi0_):
            d_100_v53_: _dafny.Array
            def lambda2_(d_101_p1_):
                def lambda3_(d_102_i5_):
                    return (d_102_i5_) * (d_101_p1_)

                return lambda3_

            init1_ = lambda2_(p1)
            nw14_ = _dafny.Array(None, 3)
            for i0_1_ in range(nw14_.length(0)):
                nw14_[i0_1_] = init1_(i0_1_)
            d_100_v53_ = nw14_
            index9_ = default__.safeIndex(646, (d_100_v53_).length(0))
            (d_100_v53_)[index9_] = (751) - (p1)
            d_103_v54_: _dafny.Array
            nw15_ = _dafny.Array(False, 6)
            d_103_v54_ = nw15_
            d_104_v55_: _dafny.Map
            d_104_v55_ = _dafny.Map({d_103_v54_: (d_97_v51_)[default__.safeIndex(p1, len(d_97_v51_))]})
            index10_ = default__.safeIndex(646, (d_100_v53_).length(0))
            rhs8_ = (0) - (((p1 if p0 else d_99_i4_)) + (default__.safeModuloInt(d_99_i4_, d_99_i4_)))
            rhs9_ = (d_46_v12_ if p3 else ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "br"))).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "br")))), p2)) + (d_46_v12_))
            rhs10_ = default__.safeModuloInt(d_99_i4_, p1)
            rhs11_ = d_104_v55_
            rhs12_ = default__.safeDivisionInt(d_99_i4_, (p1) - (p1))
            lhs3_ = globalState
            lhs4_ = globalState
            lhs5_ = d_100_v53_
            lhs6_ = default__.safeIndex(646, (d_100_v53_).length(0))
            lhs7_ = globalState
            lhs3_.f0 = rhs8_
            lhs4_.f2 = rhs9_
            lhs5_[lhs6_] = rhs10_
            d_104_v55_ = rhs11_
            lhs7_.f8 = rhs12_
            if False:
                index11_ = default__.safeIndex(643, (d_103_v54_).length(0))
                (d_103_v54_)[index11_] = p3
                index12_ = default__.safeIndex(643, (d_103_v54_).length(0))
                (d_103_v54_)[index12_] = (not((p1) != (d_99_i4_)) if not(not(p3)) else (False) and (p0))
                d_105_v56_: _dafny.Map
                d_105_v56_ = _dafny.Map({p0: (d_103_v54_)[default__.safeIndex(643, (d_103_v54_).length(0))]})
                d_106_v57_: _dafny.Map
                d_106_v57_ = _dafny.Map({(0) - ((d_100_v53_)[default__.safeIndex(646, (d_100_v53_).length(0))]): d_105_v56_})
                d_106_v57_ = (d_106_v57_).set((0) - (d_99_i4_), d_105_v56_)
                index13_ = default__.safeIndex(643, (d_103_v54_).length(0))
                (d_103_v54_)[index13_] = p3
                d_107_v58_: _dafny.Map
                d_107_v58_ = _dafny.Map({default__.fm10(p0, d_99_i4_, -914, globalState): p1})
                d_108_v60_: _dafny.Map
                d_108_v60_ = _dafny.Map({736: p1})
                d_109_v61_: _dafny.Set
                d_109_v61_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dxrmfw")), d_46_v12_})
                d_110_v62_: D1
                d_110_v62_ = D1_DC2(d_109_v61_, p2)
                d_111_v63_: _dafny.Map
                d_111_v63_ = _dafny.Map({d_108_v60_: d_110_v62_})
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: _dafny.Map
                    for compr_12_ in (d_111_v63_).keys.Elements:
                        d_112_v59_: _dafny.Map = compr_12_
                        if (d_112_v59_) in (d_111_v63_):
                            coll12_[d_112_v59_] = p0
                    return _dafny.Map(coll12_)
                d_107_v58_ = (d_107_v58_).set(iife12_()
                , p1)
                index14_ = default__.safeIndex(643, (d_103_v54_).length(0))
                (d_103_v54_)[index14_] = p3
            elif True:
                d_113_v64_: bool
                d_113_v64_ = True
                d_113_v64_ = ((d_100_v53_)[default__.safeIndex(646, (d_100_v53_).length(0))]) != (d_99_i4_)
                d_113_v64_ = (d_45_v11_) < ((d_45_v11_) + ((d_45_v11_).set(default__.safeIndex((d_100_v53_)[default__.safeIndex(646, (d_100_v53_).length(0))], len(d_45_v11_)), p1)))
                d_114_v65_: _dafny.Seq
                d_114_v65_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_48_v14_)).intersection(_dafny.MultiSet([d_98_v52_])), _dafny.MultiSet((d_48_v14_ if p0 else d_48_v14_)), _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([d_98_v52_]) if p0 else d_48_v14_)).set(default__.safeIndex(p1, len((_dafny.SeqWithoutIsStrInference([d_98_v52_]) if p0 else d_48_v14_))), d_47_v13_))])
                rhs13_ = d_47_v13_
                rhs14_ = d_47_v13_
                rhs15_ = d_114_v65_
                d_47_v13_ = rhs13_
                d_98_v52_ = rhs14_
                d_114_v65_ = rhs15_
                d_115_v66_: D4
                d_115_v66_ = D4_DC8()
                d_115_v66_ = d_115_v66_
                d_116_v67_: _dafny.Array
                nw16_ = _dafny.Array(_dafny.Seq({}), 11)
                d_116_v67_ = nw16_
                d_117_v68_: _dafny.Map
                d_117_v68_ = _dafny.Map({p0: (_dafny.MultiSet([(d_97_v51_)[default__.safeIndex(p1, len(d_97_v51_))]])).cardinality})
                d_118_v69_: _dafny.Seq
                d_118_v69_ = _dafny.SeqWithoutIsStrInference([default__.fm11(len(d_45_v11_), (d_100_v53_)[default__.safeIndex(646, (d_100_v53_).length(0))], d_117_v68_, globalState)])
                index15_ = default__.safeIndex(913, (d_116_v67_).length(0))
                (d_116_v67_)[index15_] = d_118_v69_
                d_119_v70_: _dafny.Set
                d_119_v70_ = _dafny.Set({p1})
                index16_ = default__.safeIndex(913, (d_116_v67_).length(0))
                index17_ = default__.safeIndex(646, (d_100_v53_).length(0))
                rhs16_ = (d_113_v64_ if (d_46_v12_) <= (d_46_v12_) else p0)
                rhs17_ = d_118_v69_
                rhs18_ = (d_100_v53_)[default__.safeIndex(646, (d_100_v53_).length(0))]
                rhs19_ = default__.safeDivisionInt((d_100_v53_)[default__.safeIndex(646, (d_100_v53_).length(0))], len(d_119_v70_))
                lhs8_ = d_116_v67_
                lhs9_ = default__.safeIndex(913, (d_116_v67_).length(0))
                lhs10_ = d_100_v53_
                lhs11_ = default__.safeIndex(646, (d_100_v53_).length(0))
                lhs12_ = globalState
                d_113_v64_ = rhs16_
                lhs8_[lhs9_] = rhs17_
                lhs10_[lhs11_] = rhs18_
                lhs12_.f0 = rhs19_
            index18_ = default__.safeIndex(400, (d_103_v54_).length(0))
            (d_103_v54_)[index18_] = (d_97_v51_)[default__.safeIndex(d_99_i4_, len(d_97_v51_))]
            d_120_v71_: _dafny.MultiSet
            d_120_v71_ = _dafny.MultiSet([d_100_v53_, d_100_v53_, d_100_v53_])
            index19_ = default__.safeIndex(400, (d_103_v54_).length(0))
            (d_103_v54_)[index19_] = (d_120_v71_).isdisjoint(d_120_v71_)
            d_121_v72_: _dafny.Map
            d_121_v72_ = _dafny.Map({(p0) == (p3): (d_96_v50_).ispropersubset(d_96_v50_)})
            d_122_v73_: _dafny.Map
            d_122_v73_ = _dafny.Map({d_47_v13_: (d_103_v54_)[default__.safeIndex(400, (d_103_v54_).length(0))]})
            d_123_v74_: D7
            d_123_v74_ = D7_DC14(d_122_v73_)
            d_124_v75_: _dafny.Map
            d_124_v75_ = _dafny.Map({d_96_v50_: len(((d_123_v74_).cf20).set(d_47_v13_, (d_103_v54_)[default__.safeIndex(400, (d_103_v54_).length(0))]))})
            d_121_v72_ = (d_121_v72_).set((len(d_124_v75_)) < ((0) - (p1)), (d_103_v54_)[default__.safeIndex(400, (d_103_v54_).length(0))])

    @staticmethod
    def Main(noArgsParameter__):
        d_125_v0_: _dafny.Seq
        d_125_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "si"))
        d_126_v1_: _dafny.Array
        nw17_ = _dafny.Array(_dafny.CodePoint('D'), 11)
        d_126_v1_ = nw17_
        d_127_v2_: bool
        d_127_v2_ = False
        d_128_v3_: _dafny.Array
        nw18_ = _dafny.Array(None, 3)
        nw18_[int(0)] = not(d_127_v2_)
        nw18_[int(1)] = not(d_127_v2_)
        nw18_[int(2)] = d_127_v2_
        d_128_v3_ = nw18_
        d_129_v4_: str
        d_129_v4_ = _dafny.CodePoint('j')
        d_130_v5_: _dafny.Set
        d_130_v5_ = _dafny.Set({d_129_v4_})
        d_131_globalState_: GlobalState
        nw19_ = GlobalState()
        nw19_.ctor__(-164, 378, d_125_v0_, d_126_v1_, True, True, d_128_v3_, 223, 583, d_130_v5_, 782)
        d_131_globalState_ = nw19_
        d_127_v2_ = not((not(d_127_v2_)) and (d_127_v2_))
        d_132_v6_: int
        d_132_v6_ = 520
        d_133_v7_: _dafny.Map
        d_133_v7_ = _dafny.Map({(d_126_v1_ if d_127_v2_ else d_126_v1_): (d_132_v6_ if default__.fm0(d_131_globalState_) else d_132_v6_)})
        d_134_v8_: _dafny.Seq
        d_134_v8_ = _dafny.SeqWithoutIsStrInference([d_127_v2_])
        d_135_v9_: _dafny.Set
        d_135_v9_ = _dafny.Set({d_132_v6_, d_132_v6_, len(d_134_v8_)})
        d_133_v7_ = (d_133_v7_).set(d_126_v1_, (len(d_135_v9_)) - (d_132_v6_))
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_128_v3_).length(0)):
            d_136_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_136_i0_)) and ((d_136_i0_) < ((d_128_v3_).length(0)))):
                (d_128_v3_)[(d_136_i0_)] = (_dafny.MultiSet([False, d_127_v2_, False, d_127_v2_])).issubset((_dafny.MultiSet([False])))
        d_137_v10_: _dafny.MultiSet
        d_137_v10_ = _dafny.MultiSet([not(d_127_v2_), d_127_v2_])
        d_138_v11_: _dafny.MultiSet
        d_138_v11_ = d_137_v10_
        source2_ = d_138_v11_
        d_139___mcc_h0_ = source2_
        d_140_cf0_ = d_139___mcc_h0_
        if not(not (d_127_v2_) or (not((d_135_v9_).ispropersubset(d_135_v9_)))):
            default__.m0((default__.fm0(d_131_globalState_) if d_127_v2_ else (d_134_v8_)[default__.safeIndex(d_132_v6_, len(d_134_v8_))]), (_dafny.MultiSet(d_134_v8_)).cardinality, d_129_v4_, d_127_v2_, d_131_globalState_)
            d_141_v12_: _dafny.Set
            d_141_v12_ = _dafny.Set({False})
            d_142_v13_: _dafny.Seq
            d_142_v13_ = _dafny.SeqWithoutIsStrInference([len(d_134_v8_)])
            d_143_v14_: _dafny.Map
            d_143_v14_ = _dafny.Map({(d_135_v9_) - (_dafny.Set({len(d_141_v12_), -308})): ((d_142_v13_)[default__.safeIndex(d_132_v6_, len(d_142_v13_))]) + (d_132_v6_)})
            rhs20_ = default__.fm0(d_131_globalState_)
            rhs21_ = d_132_v6_
            rhs22_ = d_132_v6_
            rhs23_ = (d_143_v14_).set(default__.fm1(_dafny.CodePoint('d'), d_127_v2_, d_127_v2_, d_131_globalState_), 80)
            lhs13_ = d_131_globalState_
            lhs14_ = d_131_globalState_
            d_127_v2_ = rhs20_
            lhs13_.f8 = rhs21_
            lhs14_.f8 = rhs22_
            d_143_v14_ = rhs23_
            d_127_v2_ = d_127_v2_
            (d_131_globalState_).f0 = d_132_v6_
            d_144_v15_: _dafny.MultiSet
            d_144_v15_ = _dafny.MultiSet([d_134_v8_, _dafny.SeqWithoutIsStrInference([d_127_v2_, d_127_v2_])])
            d_145_v16_: _dafny.Map
            d_145_v16_ = _dafny.Map({d_144_v15_: d_125_v0_})
            d_146_v17_: _dafny.Seq
            d_146_v17_ = _dafny.SeqWithoutIsStrInference([d_144_v15_, d_144_v15_, d_144_v15_])
            d_145_v16_ = (d_145_v16_).set((d_146_v17_)[default__.safeIndex(d_132_v6_, len(d_146_v17_))], d_125_v0_)
        elif True:
            index20_ = default__.safeIndex(178, (d_126_v1_).length(0))
            (d_126_v1_)[index20_] = d_129_v4_
            d_147_v18_: _dafny.Map
            d_147_v18_ = _dafny.Map({d_132_v6_: d_132_v6_})
            index21_ = default__.safeIndex(178, (d_126_v1_).length(0))
            (d_126_v1_)[index21_] = default__.fm2(d_147_v18_, d_127_v2_, d_132_v6_, d_131_globalState_)
            d_148_v19_: _dafny.Set
            d_148_v19_ = _dafny.Set({d_127_v2_})
            d_149_v20_: _dafny.Map
            d_149_v20_ = _dafny.Map({d_129_v4_: _dafny.SeqWithoutIsStrInference([d_129_v4_ for d_150_i1_ in range(default__.abs(430))])})
            d_151_v21_: _dafny.Map
            d_151_v21_ = _dafny.Map({d_148_v19_: (d_149_v20_) != (d_149_v20_)})
            d_151_v21_ = (d_151_v21_).set(d_148_v19_, d_127_v2_)
            (d_131_globalState_).f0 = d_132_v6_
            (d_131_globalState_).f2 = d_125_v0_
            default__.m0((d_134_v8_)[default__.safeIndex(d_132_v6_, len(d_134_v8_))], d_132_v6_, d_129_v4_, d_127_v2_, d_131_globalState_)
        (d_131_globalState_).f0 = d_132_v6_
        d_152_v22_: D1
        d_152_v22_ = D1_DC1(True)
        if (d_152_v22_).cf1:
            d_153_v23_: _dafny.Array
            nw20_ = _dafny.Array(_dafny.Array(None, 0), 18)
            d_153_v23_ = nw20_
            d_154_v24_: _dafny.Map
            d_154_v24_ = _dafny.Map({d_128_v3_: d_127_v2_})
            d_155_v25_: _dafny.Map
            d_155_v25_ = _dafny.Map({d_153_v23_: (d_132_v6_) * (len(d_154_v24_))})
            d_155_v25_ = (d_155_v25_).set(d_153_v23_, d_132_v6_)
            d_134_v8_ = (default__.fm3(d_131_globalState_)) + (d_134_v8_)
            (d_131_globalState_).f0 = (d_132_v6_) - (d_132_v6_)
            d_156_v26_: C0
            nw21_ = C0()
            nw21_.ctor__()
            d_156_v26_ = nw21_
            d_157_v27_: _dafny.Array
            nw22_ = _dafny.Array(_dafny.Array(None, 0), 11)
            d_157_v27_ = nw22_
            d_158_v28_: _dafny.Map
            d_158_v28_ = _dafny.Map({d_127_v2_: d_125_v0_})
            d_159_v29_: _dafny.Map
            d_159_v29_ = _dafny.Map({d_156_v26_: ((d_158_v28_)[not(d_127_v2_)] if (not(d_127_v2_)) in (d_158_v28_) else _dafny.SeqWithoutIsStrInference([d_129_v4_ for d_160_i2_ in range(default__.abs(23))]))})
            d_161_v30_: _dafny.Array
            nw23_ = _dafny.Array(None, 8)
            nw23_[int(0)] = _dafny.SeqWithoutIsStrInference([d_129_v4_, _dafny.CodePoint('g'), d_129_v4_, d_129_v4_])
            nw23_[int(1)] = ((d_159_v29_)[d_156_v26_] if (d_156_v26_) in (d_159_v29_) else _dafny.SeqWithoutIsStrInference([d_129_v4_ for d_162_i3_ in range(default__.abs(-93))]))
            nw23_[int(2)] = d_125_v0_
            nw23_[int(3)] = d_125_v0_
            nw23_[int(4)] = d_125_v0_
            nw23_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_163_i4_ in range(default__.abs(123))])
            nw23_[int(6)] = d_125_v0_
            nw23_[int(7)] = d_125_v0_
            d_161_v30_ = nw23_
            index22_ = default__.safeIndex(264, (d_157_v27_).length(0))
            (d_157_v27_)[index22_] = d_161_v30_
            index23_ = default__.safeIndex(325, (d_126_v1_).length(0))
            (d_126_v1_)[index23_] = d_129_v4_
            d_164_v31_: _dafny.Array
            nw24_ = _dafny.Array(None, 27)
            d_164_v31_ = nw24_
            d_165_v32_: _dafny.Map
            d_165_v32_ = _dafny.Map({d_164_v31_: d_161_v30_})
            index24_ = default__.safeIndex(264, (d_157_v27_).length(0))
            index25_ = default__.safeIndex(325, (d_126_v1_).length(0))
            rhs24_ = ((d_165_v32_)[d_164_v31_] if (d_164_v31_) in (d_165_v32_) else d_161_v30_)
            rhs25_ = d_153_v23_
            rhs26_ = d_129_v4_
            rhs27_ = d_132_v6_
            lhs15_ = d_157_v27_
            lhs16_ = default__.safeIndex(264, (d_157_v27_).length(0))
            lhs17_ = d_126_v1_
            lhs18_ = default__.safeIndex(325, (d_126_v1_).length(0))
            lhs19_ = d_131_globalState_
            lhs15_[lhs16_] = rhs24_
            d_153_v23_ = rhs25_
            lhs17_[lhs18_] = rhs26_
            lhs19_.f0 = rhs27_
        elif True:
            d_166_v33_: _dafny.Array
            nw25_ = _dafny.Array(_dafny.Array(None, 0), 27)
            d_166_v33_ = nw25_
            index26_ = default__.safeIndex(214, (d_166_v33_).length(0))
            (d_166_v33_)[index26_] = d_126_v1_
            index27_ = default__.safeIndex(214, (d_166_v33_).length(0))
            rhs28_ = d_126_v1_
            rhs29_ = (0) - ((d_132_v6_) + (d_132_v6_))
            lhs20_ = d_166_v33_
            lhs21_ = default__.safeIndex(214, (d_166_v33_).length(0))
            lhs22_ = d_131_globalState_
            lhs20_[lhs21_] = rhs28_
            lhs22_.f8 = rhs29_
            d_167_v34_: _dafny.Map
            d_167_v34_ = _dafny.Map({d_127_v2_: d_132_v6_})
            d_167_v34_ = (d_167_v34_).set(d_127_v2_, d_132_v6_)
            d_168_v35_: _dafny.Map
            d_168_v35_ = _dafny.Map({d_127_v2_: d_127_v2_})
            d_169_v36_: _dafny.Seq
            d_169_v36_ = _dafny.SeqWithoutIsStrInference([d_135_v9_])
            d_170_v37_: _dafny.Map
            d_170_v37_ = _dafny.Map({d_132_v6_: len(d_125_v0_)})
            d_171_v38_: _dafny.MultiSet
            d_171_v38_ = _dafny.MultiSet([d_129_v4_])
            rhs30_ = (d_169_v36_)[default__.safeIndex((d_132_v6_) * (d_132_v6_), len(d_169_v36_))]
            rhs31_ = default__.fm4(d_127_v2_, d_132_v6_, d_131_globalState_)
            rhs32_ = default__.safeDivisionInt(d_132_v6_, d_132_v6_)
            rhs33_ = not ((len(d_170_v37_)) == ((d_171_v38_).cardinality)) or ((d_127_v2_) and (d_127_v2_))
            rhs34_ = ((d_168_v35_) | ((d_168_v35_).set(d_127_v2_, True))).set(default__.fm0(d_131_globalState_), d_127_v2_)
            lhs23_ = d_131_globalState_
            d_135_v9_ = rhs30_
            d_132_v6_ = rhs31_
            lhs23_.f0 = rhs32_
            d_127_v2_ = rhs33_
            d_168_v35_ = rhs34_
            d_172_v39_: _dafny.Seq
            d_172_v39_ = _dafny.SeqWithoutIsStrInference([d_170_v37_])
            d_173_v40_: _dafny.Map
            d_173_v40_ = _dafny.Map({d_132_v6_: d_127_v2_})
            d_174_v41_: _dafny.Map
            d_174_v41_ = _dafny.Map({len(_dafny.Set({((d_173_v40_)[d_132_v6_] if (d_132_v6_) in (d_173_v40_) else d_127_v2_)})): d_127_v2_})
            d_175_v43_: _dafny.Array
            nw26_ = _dafny.Array(None, 10)
            nw26_[int(0)] = d_170_v37_
            nw26_[int(1)] = (d_172_v39_)[default__.safeIndex(43, len(d_172_v39_))]
            nw26_[int(2)] = default__.fm5(183, d_127_v2_, d_131_globalState_)
            nw26_[int(3)] = (d_170_v37_) | (d_170_v37_)
            nw26_[int(4)] = d_170_v37_
            nw26_[int(5)] = d_170_v37_
            nw26_[int(6)] = (default__.fm5(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aywqiw"))), d_127_v2_, d_131_globalState_)) | (d_170_v37_)
            def iife13_():
                coll13_ = _dafny.Set()
                compr_13_: int
                for compr_13_ in (d_174_v41_).keys.Elements:
                    d_176_v42_: int = compr_13_
                    if (d_176_v42_) in (d_174_v41_):
                        coll13_ = coll13_.union(_dafny.Set([(d_176_v42_) * (51)]))
                return _dafny.Set(coll13_)
            nw26_[int(7)] = _dafny.Map({377: len(iife13_()
            )})
            nw26_[int(8)] = (d_170_v37_) | (default__.fm5(d_132_v6_, d_127_v2_, d_131_globalState_))
            nw26_[int(9)] = d_170_v37_
            d_175_v43_ = nw26_
            index28_ = default__.safeIndex(809, (d_175_v43_).length(0))
            (d_175_v43_)[index28_] = _dafny.Map({d_132_v6_: len(d_134_v8_)})
            d_177_v44_: _dafny.Set
            d_177_v44_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nagnmac"))})
            d_178_v45_: D1
            d_178_v45_ = D1_DC2(d_177_v44_, d_129_v4_)
            d_179_v46_: _dafny.Array
            nw27_ = _dafny.Array(None, 7)
            d_179_v46_ = nw27_
            d_180_v47_: _dafny.Map
            d_180_v47_ = _dafny.Map({d_140_cf0_: (d_179_v46_ if d_127_v2_ else d_179_v46_)})
            d_181_v48_: _dafny.Map
            d_181_v48_ = _dafny.Map({d_132_v6_: d_170_v37_})
            index29_ = default__.safeIndex(809, (d_175_v43_).length(0))
            rhs35_ = ((d_170_v37_).set(d_132_v6_, d_132_v6_)) | (((d_181_v48_)[d_132_v6_] if (d_132_v6_) in (d_181_v48_) else d_170_v37_))
            rhs36_ = d_178_v45_
            rhs37_ = (d_180_v47_) | ((d_180_v47_) | (_dafny.Map({d_138_v11_: d_179_v46_})))
            rhs38_ = (d_132_v6_) >= (d_132_v6_)
            lhs24_ = d_175_v43_
            lhs25_ = default__.safeIndex(809, (d_175_v43_).length(0))
            lhs24_[lhs25_] = rhs35_
            d_178_v45_ = rhs36_
            d_180_v47_ = rhs37_
            d_127_v2_ = rhs38_
            d_127_v2_ = False
        default__.m0(d_127_v2_, (0) - (d_132_v6_), d_129_v4_, (-278) >= (d_132_v6_), d_131_globalState_)
        d_127_v2_ = d_127_v2_
        index30_ = default__.safeIndex(849, (d_128_v3_).length(0))
        (d_128_v3_)[index30_] = d_127_v2_
        index31_ = default__.safeIndex(849, (d_128_v3_).length(0))
        (d_128_v3_)[index31_] = not(d_127_v2_)
        d_182_v49_: _dafny.Seq
        d_182_v49_ = _dafny.SeqWithoutIsStrInference([default__.fm4((d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))], d_132_v6_, d_131_globalState_)])
        d_183_v50_: _dafny.Map
        d_183_v50_ = _dafny.Map({d_182_v49_: d_132_v6_})
        d_184_v51_: _dafny.Map
        d_184_v51_ = _dafny.Map({(d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]: len((d_125_v0_).set(default__.safeIndex(122, len(d_125_v0_)), d_129_v4_))})
        d_185_v52_: _dafny.Map
        d_185_v52_ = _dafny.Map({d_183_v50_: d_184_v51_})
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: _dafny.Seq
            for compr_14_ in (d_183_v50_).keys.Elements:
                d_186_v53_: _dafny.Seq = compr_14_
                if (d_186_v53_) in (d_183_v50_):
                    coll14_[d_186_v53_] = len(d_135_v9_)
            return _dafny.Map(coll14_)
        d_185_v52_ = (d_185_v52_).set((_dafny.Map({d_182_v49_: d_132_v6_})) | (iife14_()
        ), _dafny.Map({(d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]: d_132_v6_}))
        if (d_127_v2_) or ((d_132_v6_) >= (d_132_v6_)):
            (d_131_globalState_).f8 = d_132_v6_
            d_187_v54_: _dafny.Map
            d_187_v54_ = _dafny.Map({d_128_v3_: d_127_v2_})
            d_187_v54_ = ((d_187_v54_) | (d_187_v54_)).set(d_128_v3_, True)
            d_188_v55_: _dafny.Seq
            d_188_v55_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnjsvda"))])
            d_189_v56_: _dafny.Seq
            d_189_v56_ = ((d_134_v8_).set(default__.safeIndex(len(d_188_v55_), len(d_134_v8_)), d_127_v2_)).set(default__.safeIndex(len(d_125_v0_), len((d_134_v8_).set(default__.safeIndex(len(d_188_v55_), len(d_134_v8_)), d_127_v2_))), False)
            d_134_v8_ = (d_189_v56_)
            d_190_v57_: C0
            nw28_ = C0()
            nw28_.ctor__()
            d_190_v57_ = nw28_
            d_191_v58_: _dafny.Seq
            d_191_v58_ = _dafny.SeqWithoutIsStrInference([d_190_v57_])
            d_192_v59_: _dafny.Set
            d_192_v59_ = _dafny.Set({d_190_v57_})
            d_193_v60_: _dafny.Map
            d_193_v60_ = _dafny.Map({d_132_v6_: d_190_v57_})
            d_194_v61_: D3
            d_194_v61_ = D3_DC4(_dafny.SeqWithoutIsStrInference([d_129_v4_ for d_195_i5_ in range(default__.abs(244))]))
            d_196_v62_: _dafny.Seq
            d_196_v62_ = _dafny.SeqWithoutIsStrInference([d_138_v11_, default__.fm6(_dafny.CodePoint('i'), (d_194_v61_).cf5, d_127_v2_, d_131_globalState_)])
            default__.m0((_dafny.Set({(d_191_v58_)[default__.safeIndex(d_132_v6_, len(d_191_v58_))], d_190_v57_, d_190_v57_, d_190_v57_, d_190_v57_})).isdisjoint(d_192_v59_), len((d_193_v60_) | (_dafny.Map({d_132_v6_: d_190_v57_}))), d_129_v4_, (d_196_v62_) != (_dafny.SeqWithoutIsStrInference([d_138_v11_, d_138_v11_])), d_131_globalState_)
            d_197_v63_: _dafny.Array
            def lambda4_(d_198_v6_):
                def lambda5_(d_199_i6_):
                    return (d_199_i6_) * (d_198_v6_)

                return lambda5_

            init2_ = lambda4_(d_132_v6_)
            nw29_ = _dafny.Array(None, 3)
            for i0_2_ in range(nw29_.length(0)):
                nw29_[i0_2_] = init2_(i0_2_)
            d_197_v63_ = nw29_
            d_197_v63_ = d_197_v63_
        elif True:
            d_200_v64_: _dafny.Map
            d_200_v64_ = _dafny.Map({d_132_v6_: d_132_v6_})
            d_200_v64_ = (_dafny.Map({len(d_125_v0_): d_132_v6_})) | (_dafny.Map({d_132_v6_: d_132_v6_}))
            d_201_v65_: D3
            d_201_v65_ = D3_DC4(d_125_v0_)
            d_202_v66_: _dafny.Seq
            d_202_v66_ = _dafny.SeqWithoutIsStrInference([d_201_v65_, d_201_v65_])
            d_203_v67_: _dafny.Map
            d_203_v67_ = _dafny.Map({default__.fm0(d_131_globalState_): d_202_v66_})
            d_204_v68_: _dafny.Map
            d_204_v68_ = _dafny.Map({not((d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]): ((d_203_v67_)[(d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]] if ((d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]) in (d_203_v67_) else _dafny.SeqWithoutIsStrInference([d_201_v65_]))})
            d_205_v69_: D4
            d_205_v69_ = D4_DC6(d_202_v66_)
            d_203_v67_ = (d_203_v67_).set((d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))], (d_205_v69_).cf8)
            d_206_v70_: _dafny.Map
            d_206_v70_ = _dafny.Map({d_132_v6_: d_127_v2_})
            index32_ = default__.safeIndex(849, (d_128_v3_).length(0))
            (d_128_v3_)[index32_] = (d_206_v70_) != (default__.fm7(d_131_globalState_))
            index33_ = default__.safeIndex(849, (d_128_v3_).length(0))
            (d_128_v3_)[index33_] = not(d_127_v2_)
            d_135_v9_ = d_135_v9_
        d_127_v2_ = d_127_v2_
        d_207_v71_: _dafny.Array
        nw30_ = _dafny.Array(_dafny.Map({}), 4)
        d_207_v71_ = nw30_
        def iife15_():
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: _dafny.Seq
                for compr_17_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_209_i8_ in range(default__.abs(862))])).Elements:
                    d_210_v73_: _dafny.Seq = compr_17_
                    if (d_210_v73_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_209_i8_ in range(default__.abs(862))])):
                        coll17_[d_210_v73_] = _dafny.Map({not(True): d_132_v6_})
                return _dafny.Map(coll17_)
            coll15_ = _dafny.Map()
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: _dafny.Seq
                for compr_16_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_209_i8_ in range(default__.abs(862))])).Elements:
                    d_210_v73_: _dafny.Seq = compr_16_
                    if (d_210_v73_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False]) for d_209_i8_ in range(default__.abs(862))])):
                        coll16_[d_210_v73_] = _dafny.Map({not(True): d_132_v6_})
                return _dafny.Map(coll16_)
            compr_15_: _dafny.Seq
            for compr_15_ in (iife16_()
            ).keys.Elements:
                d_211_v72_: _dafny.Seq = compr_15_
                if (d_211_v72_) in (iife17_()
                ):
                    coll15_[d_211_v72_] = (d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]
            return _dafny.Map(coll15_)
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_207_v71_).length(0)):
            d_208_i7_: int = guard_loop_1_
            if (True) and (((0) <= (d_208_i7_)) and ((d_208_i7_) < ((d_207_v71_).length(0)))):
                _ingredientsd_0.append((d_207_v71_, int(d_208_i7_), ((_dafny.Map({d_134_v8_: (d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]})) | (iife15_()
                )) | ((_dafny.Map({d_134_v8_: not((d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))])}) if (d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))] else _dafny.Map({d_134_v8_: (d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]})))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        d_212_v74_: C0
        nw31_ = C0()
        nw31_.ctor__()
        d_212_v74_ = nw31_
        d_213_v75_: _dafny.Map
        d_213_v75_ = _dafny.Map({d_132_v6_: d_132_v6_})
        d_125_v0_ = (default__.fm8(d_132_v6_, d_213_v75_, len(d_182_v49_), d_131_globalState_)) + ((d_125_v0_) + (d_125_v0_))
        d_214_v76_: _dafny.Map
        d_214_v76_ = _dafny.Map({d_212_v74_: d_132_v6_})
        d_214_v76_ = (d_214_v76_).set(d_212_v74_, d_132_v6_)
        d_127_v2_ = (d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]
        d_215_v77_: _dafny.Array
        nw32_ = _dafny.Array(int(0), 23)
        d_215_v77_ = nw32_
        d_215_v77_ = d_215_v77_
        d_216_v78_: _dafny.MultiSet
        d_216_v78_ = _dafny.MultiSet([d_132_v6_])
        hi1_ = ((d_216_v78_).cardinality) * (d_132_v6_)
        for d_217_i9_ in range(default__.fm4(False, len((d_134_v8_).set(default__.safeIndex(79, len(d_134_v8_)), (D1_DC1((d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))])).cf1)), d_131_globalState_), hi1_):
            d_218_v79_: _dafny.Seq
            d_218_v79_ = _dafny.SeqWithoutIsStrInference([d_128_v3_, d_128_v3_, d_128_v3_])
            d_218_v79_ = _dafny.SeqWithoutIsStrInference([d_128_v3_])
            d_219_v80_: _dafny.Array
            d_219_v80_ = d_215_v77_
            d_220_v81_: _dafny.Seq
            d_220_v81_ = _dafny.SeqWithoutIsStrInference([d_215_v77_, (d_219_v80_), d_215_v77_, d_215_v77_])
            d_220_v81_ = d_220_v81_
            d_221_v82_: D3
            d_221_v82_ = D3_DC5(d_127_v2_, (d_182_v49_)[default__.safeIndex(d_217_i9_, len(d_182_v49_))])
            index34_ = default__.safeIndex(849, (d_128_v3_).length(0))
            (d_128_v3_)[index34_] = (d_132_v6_) > (((d_221_v82_).cf7) * (-792))
            index35_ = default__.safeIndex(849, (d_128_v3_).length(0))
            index36_ = default__.safeIndex(849, (d_128_v3_).length(0))
            rhs39_ = not(d_127_v2_)
            rhs40_ = ((d_137_v10_).intersection(d_137_v10_)).ispropersubset(d_137_v10_)
            rhs41_ = len(((_dafny.SeqWithoutIsStrInference([(d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]])) + (d_134_v8_)) + (_dafny.SeqWithoutIsStrInference([d_127_v2_, (d_128_v3_)[default__.safeIndex(849, (d_128_v3_).length(0))]])))
            rhs42_ = (_dafny.Set({d_132_v6_})) - (d_135_v9_)
            rhs43_ = d_212_v74_
            lhs26_ = d_128_v3_
            lhs27_ = default__.safeIndex(849, (d_128_v3_).length(0))
            lhs28_ = d_128_v3_
            lhs29_ = default__.safeIndex(849, (d_128_v3_).length(0))
            lhs30_ = d_131_globalState_
            lhs26_[lhs27_] = rhs39_
            lhs28_[lhs29_] = rhs40_
            lhs30_.f8 = rhs41_
            d_135_v9_ = rhs42_
            d_212_v74_ = rhs43_
        _dafny.print((d_125_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v1_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_126_v1_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_127_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v3_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v3_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_128_v3_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_129_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v5_) == (_dafny.Set({_dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_131_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_131_globalState_.f2).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_globalState_).f3)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_globalState_).f3)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_globalState_).f6)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_globalState_).f6)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_globalState_).f6)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_131_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_131_globalState_).f9) == (_dafny.Set({_dafny.CodePoint('j')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_132_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_133_v7_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_134_v8_) == (_dafny.SeqWithoutIsStrInference([True, True, False, False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v9_) == (_dafny.Set({520, 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v10_) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_v11_)) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_182_v49_) == (_dafny.SeqWithoutIsStrInference([549]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v50_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([549]): 520}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_184_v51_) == (_dafny.Map({False: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_185_v52_) == (_dafny.Map({_dafny.Map({_dafny.SeqWithoutIsStrInference([549]): 520}): _dafny.Map({False: 2}), _dafny.Map({_dafny.SeqWithoutIsStrInference([549]): 2}): _dafny.Map({False: 520})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v71_)[0]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True, False, False, True]): False, _dafny.SeqWithoutIsStrInference([False]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v71_)[1]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True, False, False, True]): False, _dafny.SeqWithoutIsStrInference([False]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v71_)[2]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True, False, False, True]): False, _dafny.SeqWithoutIsStrInference([False]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_207_v71_)[3]) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True, False, False, True]): False, _dafny.SeqWithoutIsStrInference([False]): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_213_v75_) == (_dafny.Map({520: 520}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_214_v76_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_216_v78_) == (_dafny.MultiSet([520]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
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
        return lambda: D1_DC2(_dafny.Set({}), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D2_DC3)

class D2_DC3(D2, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC5(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D3_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D3_DC4)

class D3_DC5(D3, NamedTuple('DC5', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC5({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC5) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC4(D3, NamedTuple('DC4', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC4({self.cf5.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC4) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC7(False, None, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D4_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D4_DC8)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D4_DC6)

class D4_DC7(D4, NamedTuple('DC7', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC8(D4, NamedTuple('DC8', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC6(D4, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D5_DC9)

class D5_DC9(D5, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC9({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC11(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D6_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D6_DC10)

class D6_DC11(D6, NamedTuple('DC11', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC11({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC11) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC13(D6, NamedTuple('DC13', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC10(D6, NamedTuple('DC10', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC10({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC10) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC15(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D7_DC15)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D7_DC14)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC15(D7, NamedTuple('DC15', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC15({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC15) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC14(D7, NamedTuple('DC14', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC14({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC14) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC18(D8, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f8: int = int(0)
        self._f1: int = int(0)
        self._f3: _dafny.Array = _dafny.Array(None, 0)
        self._f4: bool = False
        self._f5: bool = False
        self._f6: _dafny.Array = _dafny.Array(None, 0)
        self._f7: int = int(0)
        self._f9: _dafny.Set = _dafny.Set({})
        self._f10: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self).f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10

    @property
    def f1(self):
        return self._f1
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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

