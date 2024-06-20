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
    def fm0(p0, p1, p2, globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm1(p0, p1, globalState):
        return len(_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False), False])), (_dafny.MultiSet([not(not(not((False))))])).cardinality, (0) - (-825)]): (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wirhrrwqa"))), (0) - (-305), 846, len(_dafny.SeqWithoutIsStrInference([not(True)]))])).intersection(_dafny.MultiSet([(_dafny.MultiSet([not(True)])).cardinality]))}))

    @staticmethod
    def fm2(p0, p1, globalState):
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([-177]), _dafny.SeqWithoutIsStrInference([len((_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([True]))}))), len(_dafny.SeqWithoutIsStrInference([False, True, True])), len(_dafny.Map({288: not(False)}))])})).issubset(_dafny.Set({_dafny.SeqWithoutIsStrInference([423]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kxsyhfc")))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: False}))])}))

    @staticmethod
    def fm3(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), _dafny.CodePoint('b'), _dafny.CodePoint('e')])

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return False

    @staticmethod
    def fm6(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([206 for d_0_i0_ in range(default__.abs(331))])).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([206 for d_0_i0_ in range(default__.abs(331))])):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_1_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([False]))))]))
            return _dafny.Set(coll0_)
        return (_dafny.Map({True: 800})) | ((_dafny.Map({True: 186})) | (_dafny.Map({not(True): len(iife0_()
        )})))

    @staticmethod
    def fm7(p0, globalState):
        return (D5_DC5(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('d')])}))).cf5

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([len(_dafny.Map({True: 934}))])

    @staticmethod
    def fm9(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(True if True else not(not(False))), True, False, not(((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))), len(_dafny.Map({not(True): True})), 274, 377])).cardinality) != (-936)), not (False) or (False)])

    @staticmethod
    def m0(p0, p1, p2, globalState):
        r0: _dafny.Map = _dafny.Map({})
        d_2_v0_: int
        d_2_v0_ = 896
        (globalState).f12 = (d_2_v0_) >= (d_2_v0_)
        d_3_v1_: _dafny.Set
        d_3_v1_ = _dafny.Set({p2})
        d_4_v2_: _dafny.Map
        d_4_v2_ = _dafny.Map({p1: d_2_v0_})
        d_5_v3_: _dafny.Map
        d_5_v3_ = _dafny.Map({default__.fm1(p1, d_4_v2_, globalState): p2})
        if ((d_3_v1_).intersection(_dafny.Set({((d_5_v3_)[(0) - (d_2_v0_)] if ((0) - (d_2_v0_)) in (d_5_v3_) else p2)}))).issubset(d_3_v1_):
            d_6_v4_: _dafny.Seq
            d_6_v4_ = _dafny.SeqWithoutIsStrInference([p0])
            d_6_v4_ = ((_dafny.SeqWithoutIsStrInference([False])) + (d_6_v4_)) + (d_6_v4_)
            d_7_v5_: _dafny.Array
            nw0_ = _dafny.Array(None, 12)
            nw0_[int(0)] = p2
            nw0_[int(1)] = p2
            nw0_[int(2)] = p1
            nw0_[int(3)] = p2
            nw0_[int(4)] = p2
            nw0_[int(5)] = not(not(p0))
            nw0_[int(6)] = not(p2)
            nw0_[int(7)] = p1
            nw0_[int(8)] = p1
            nw0_[int(9)] = p0
            nw0_[int(10)] = not(p0)
            nw0_[int(11)] = p1
            d_7_v5_ = nw0_
            d_8_v6_: C0
            nw1_ = C0()
            nw1_.ctor__()
            d_8_v6_ = nw1_
            d_9_v7_: _dafny.Map
            d_9_v7_ = _dafny.Map({d_7_v5_: d_8_v6_})
            d_9_v7_ = (d_9_v7_).set(d_7_v5_, d_8_v6_)
            d_6_v4_ = (d_6_v4_) + (d_6_v4_)
            d_10_v8_: _dafny.Map
            d_10_v8_ = _dafny.Map({p2: p0})
            source0_ = d_10_v8_
            d_11___mcc_h0_ = source0_
            d_12_cf1_ = d_11___mcc_h0_
            d_13_v9_: C0
            nw2_ = C0()
            nw2_.ctor__()
            d_13_v9_ = nw2_
            d_14_v10_: _dafny.Map
            d_14_v10_ = _dafny.Map({d_8_v6_: p2})
            d_15_v11_: _dafny.MultiSet
            d_15_v11_ = _dafny.MultiSet([(d_4_v2_) | (default__.fm6(((d_14_v10_)[d_8_v6_] if (d_8_v6_) in (d_14_v10_) else p1), p0, globalState))])
            (globalState).f5 = ((d_15_v11_)[d_4_v2_] if (d_4_v2_) in (d_15_v11_) else d_2_v0_)
            d_16_v12_: _dafny.Array
            def lambda0_(d_17_i0_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mm"))

            init0_ = lambda0_
            nw3_ = _dafny.Array(None, 27)
            for i0_0_ in range(nw3_.length(0)):
                nw3_[i0_0_] = init0_(i0_0_)
            d_16_v12_ = nw3_
            d_18_v13_: _dafny.Seq
            d_18_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "igqyhvkw"))
            d_19_v14_: _dafny.Seq
            d_19_v14_ = d_18_v13_
            index0_ = default__.safeIndex(607, (d_16_v12_).length(0))
            (d_16_v12_)[index0_] = d_19_v14_
            index1_ = default__.safeIndex(607, (d_16_v12_).length(0))
            (d_16_v12_)[index1_] = (d_18_v13_) + (d_18_v13_)
            d_20_v15_: str
            d_20_v15_ = _dafny.CodePoint('j')
            d_21_v16_: int
            d_21_v16_ = d_2_v0_
            d_22_v17_: _dafny.MultiSet
            d_22_v17_ = _dafny.MultiSet([d_13_v9_])
            d_23_v18_: _dafny.Array
            nw4_ = _dafny.Array(None, 10)
            nw4_[int(0)] = (d_22_v17_).cardinality
            nw4_[int(1)] = d_2_v0_
            nw4_[int(2)] = len(d_4_v2_)
            nw4_[int(3)] = len(d_18_v13_)
            nw4_[int(4)] = len(_dafny.SeqWithoutIsStrInference([d_2_v0_]))
            nw4_[int(5)] = len(d_6_v4_)
            nw4_[int(6)] = d_2_v0_
            nw4_[int(7)] = d_2_v0_
            nw4_[int(8)] = len(d_18_v13_)
            nw4_[int(9)] = len(((d_18_v13_).set(default__.safeIndex(d_2_v0_, len(d_18_v13_)), d_20_v15_)).set(default__.safeIndex(d_2_v0_, len((d_18_v13_).set(default__.safeIndex(d_2_v0_, len(d_18_v13_)), d_20_v15_))), d_20_v15_))
            d_23_v18_ = nw4_
            d_24_v19_: _dafny.MultiSet
            d_24_v19_ = _dafny.MultiSet([(d_21_v16_), len(_dafny.Map({not(p0): d_23_v18_})), (d_2_v0_) * (default__.fm1(p0, d_4_v2_, globalState))])
            rhs0_ = (default__.fm1(p0, d_4_v2_, globalState)) > (d_2_v0_)
            rhs1_ = d_20_v15_
            rhs2_ = d_24_v19_
            rhs3_ = p2
            lhs0_ = globalState
            lhs1_ = globalState
            lhs0_.f7 = rhs0_
            d_20_v15_ = rhs1_
            d_24_v19_ = rhs2_
            lhs1_.f7 = rhs3_
            if not((len(d_6_v4_)) > (d_2_v0_)):
                index2_ = default__.safeIndex(353, (d_7_v5_).length(0))
                (d_7_v5_)[index2_] = p1
                index3_ = default__.safeIndex(353, (d_7_v5_).length(0))
                (d_7_v5_)[index3_] = True
                (globalState).f7 = (d_2_v0_) >= (d_2_v0_)
                (globalState).f12 = p0
                d_25_v20_: _dafny.Array
                def lambda1_(d_26_v0_):
                    def lambda2_(d_27_i1_):
                        return default__.safeDivisionInt(d_27_i1_, d_26_v0_)

                    return lambda2_

                init1_ = lambda1_(d_2_v0_)
                nw5_ = _dafny.Array(None, 24)
                for i0_1_ in range(nw5_.length(0)):
                    nw5_[i0_1_] = init1_(i0_1_)
                d_25_v20_ = nw5_
                d_28_v21_: _dafny.Seq
                d_28_v21_ = _dafny.SeqWithoutIsStrInference([d_25_v20_, d_25_v20_])
                d_28_v21_ = d_28_v21_
                (globalState).f3 = d_2_v0_
            elif True:
                (globalState).f12 = (d_2_v0_) < ((d_2_v0_) - ((0) - (d_2_v0_)))
                d_29_v22_: _dafny.Seq
                d_29_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aov"))
                d_30_v23_: _dafny.Map
                d_30_v23_ = _dafny.Map({d_8_v6_: d_29_v22_})
                d_31_v24_: int
                d_31_v24_ = -98
                d_32_v25_: _dafny.Map
                d_32_v25_ = _dafny.Map({(d_31_v24_): d_29_v22_})
                d_33_v26_: _dafny.MultiSet
                d_33_v26_ = _dafny.MultiSet([d_8_v6_])
                d_34_v27_: _dafny.Array
                nw6_ = _dafny.Array(None, 21)
                nw6_[int(0)] = default__.fm1(p2, d_4_v2_, globalState)
                nw6_[int(1)] = len((d_30_v23_) | (d_30_v23_))
                nw6_[int(2)] = d_2_v0_
                nw6_[int(3)] = (0) - (d_2_v0_)
                nw6_[int(4)] = (0) - (default__.fm1(p1, d_4_v2_, globalState))
                nw6_[int(5)] = d_2_v0_
                nw6_[int(6)] = (d_31_v24_)
                nw6_[int(7)] = d_2_v0_
                nw6_[int(8)] = d_2_v0_
                nw6_[int(9)] = len(d_32_v25_)
                nw6_[int(10)] = d_2_v0_
                nw6_[int(11)] = len((d_4_v2_) | (d_4_v2_))
                nw6_[int(12)] = d_2_v0_
                nw6_[int(13)] = d_2_v0_
                nw6_[int(14)] = 597
                nw6_[int(15)] = d_2_v0_
                nw6_[int(16)] = 353
                nw6_[int(17)] = (0) - (default__.safeDivisionInt(d_2_v0_, (0) - (len(d_3_v1_))))
                nw6_[int(18)] = (d_2_v0_) - ((d_33_v26_).cardinality)
                nw6_[int(19)] = d_2_v0_
                nw6_[int(20)] = default__.fm1(p0, d_4_v2_, globalState)
                d_34_v27_ = nw6_
                index4_ = default__.safeIndex(307, (d_34_v27_).length(0))
                (d_34_v27_)[index4_] = d_2_v0_
                index5_ = default__.safeIndex(307, (d_34_v27_).length(0))
                (d_34_v27_)[index5_] = (d_2_v0_) * (d_2_v0_)
                d_35_v28_: str
                d_35_v28_ = _dafny.CodePoint('x')
                d_35_v28_ = d_35_v28_
                d_36_v29_: C0
                nw7_ = C0()
                nw7_.ctor__()
                d_36_v29_ = nw7_
                index6_ = default__.safeIndex(821, (d_7_v5_).length(0))
                (d_7_v5_)[index6_] = p2
                index7_ = default__.safeIndex(821, (d_7_v5_).length(0))
                (d_7_v5_)[index7_] = p0
        elif True:
            d_37_v30_: C0
            nw8_ = C0()
            nw8_.ctor__()
            d_37_v30_ = nw8_
            d_38_v31_: _dafny.Seq
            d_38_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "corwclx"))
            d_39_v32_: _dafny.Map
            d_39_v32_ = _dafny.Map({d_2_v0_: default__.fm8(True, d_2_v0_, d_38_v31_, d_2_v0_, globalState)})
            d_40_v33_: _dafny.MultiSet
            d_40_v33_ = _dafny.MultiSet([d_2_v0_])
            d_39_v32_ = (d_39_v32_).set(d_2_v0_, d_40_v33_)
            d_41_v34_: _dafny.Map
            d_41_v34_ = _dafny.Map({296: d_5_v3_})
            d_41_v34_ = (d_41_v34_).set(default__.safeDivisionInt(d_2_v0_, d_2_v0_), d_5_v3_)
            (globalState).f12 = p2
            d_42_v35_: _dafny.Array
            def lambda3_(d_43_i2_):
                return _dafny.CodePoint('y')

            init2_ = lambda3_
            nw9_ = _dafny.Array(None, 2)
            for i0_2_ in range(nw9_.length(0)):
                nw9_[i0_2_] = init2_(i0_2_)
            d_42_v35_ = nw9_
            d_44_v36_: str
            d_44_v36_ = _dafny.CodePoint('b')
            index8_ = default__.safeIndex(563, (d_42_v35_).length(0))
            (d_42_v35_)[index8_] = d_44_v36_
            index9_ = default__.safeIndex(563, (d_42_v35_).length(0))
            (d_42_v35_)[index9_] = d_44_v36_
        d_45_i3_: int
        d_45_i3_ = 0
        with _dafny.label("0"):
            while p0:
                with _dafny.c_label("0"):
                    if (d_45_i3_) >= (100):
                        raise _dafny.Break("0")
                    d_45_i3_ = (d_45_i3_) + (1)
                    if p1:
                        (globalState).f3 = d_2_v0_
                        d_46_v37_: _dafny.Array
                        def lambda4_(d_47_i4_):
                            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "woh"))

                        init3_ = lambda4_
                        nw10_ = _dafny.Array(None, 5)
                        for i0_3_ in range(nw10_.length(0)):
                            nw10_[i0_3_] = init3_(i0_3_)
                        d_46_v37_ = nw10_
                        d_48_v38_: _dafny.Seq
                        d_48_v38_ = default__.fm3(p0, d_2_v0_, globalState)
                        index10_ = default__.safeIndex(903, (d_46_v37_).length(0))
                        (d_46_v37_)[index10_] = d_48_v38_
                        d_49_v39_: _dafny.Seq
                        d_49_v39_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "siujnxa"))
                        index11_ = default__.safeIndex(903, (d_46_v37_).length(0))
                        (d_46_v37_)[index11_] = d_49_v39_
                        (globalState).f7 = p1
                        d_50_v40_: _dafny.Array
                        nw11_ = _dafny.Array(None, 13)
                        nw11_[int(0)] = p1
                        nw11_[int(1)] = p1
                        nw11_[int(2)] = p2
                        nw11_[int(3)] = p1
                        nw11_[int(4)] = (p2 if p0 else p2)
                        nw11_[int(5)] = p2
                        nw11_[int(6)] = p1
                        nw11_[int(7)] = p0
                        nw11_[int(8)] = p0
                        nw11_[int(9)] = (default__.fm2((0) - (d_2_v0_), p0, globalState)) not in (d_3_v1_)
                        nw11_[int(10)] = p2
                        nw11_[int(11)] = p0
                        nw11_[int(12)] = (p0) and (p1)
                        d_50_v40_ = nw11_
                        d_50_v40_ = d_50_v40_
                        d_51_v41_: C0
                        nw12_ = C0()
                        nw12_.ctor__()
                        d_51_v41_ = nw12_
                        d_52_v42_: _dafny.Map
                        d_52_v42_ = _dafny.Map({p1: p2})
                        d_53_v43_: _dafny.MultiSet
                        d_53_v43_ = _dafny.MultiSet([False, True, ((d_52_v42_)[True] if (True) in (d_52_v42_) else p0)])
                        rhs4_ = p0
                        rhs5_ = d_2_v0_
                        rhs6_ = False
                        rhs7_ = (d_53_v43_).ispropersubset(d_53_v43_)
                        rhs8_ = d_51_v41_
                        lhs2_ = globalState
                        lhs3_ = globalState
                        lhs4_ = globalState
                        lhs5_ = globalState
                        lhs2_.f12 = rhs4_
                        lhs3_.f3 = rhs5_
                        lhs4_.f12 = rhs6_
                        lhs5_.f12 = rhs7_
                        d_51_v41_ = rhs8_
                    elif True:
                        d_54_v44_: C0
                        nw13_ = C0()
                        nw13_.ctor__()
                        d_54_v44_ = nw13_
                        d_55_v45_: _dafny.Seq
                        d_55_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sb"))
                        (globalState).f12 = (d_2_v0_) > (len(d_55_v45_))
                        d_56_v46_: _dafny.MultiSet
                        d_56_v46_ = _dafny.MultiSet([d_2_v0_])
                        (globalState).f7 = ((d_56_v46_) - (d_56_v46_)).ispropersubset(default__.fm8(p0, d_2_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eqibxy")), default__.fm1(p1, _dafny.Map({False: d_2_v0_}), globalState), globalState))
                        d_57_v47_: _dafny.Array
                        def lambda5_(d_58_p1_):
                            def lambda6_(d_59_i5_):
                                return d_58_p1_

                            return lambda6_

                        init4_ = lambda5_(p1)
                        nw14_ = _dafny.Array(None, 14)
                        for i0_4_ in range(nw14_.length(0)):
                            nw14_[i0_4_] = init4_(i0_4_)
                        d_57_v47_ = nw14_
                        d_60_v48_: _dafny.Map
                        d_60_v48_ = _dafny.Map({p0: p2})
                        d_61_v49_: _dafny.Seq
                        d_61_v49_ = _dafny.SeqWithoutIsStrInference([d_60_v48_, d_60_v48_])
                        d_62_v50_: _dafny.Map
                        d_62_v50_ = _dafny.Map({d_57_v47_: (d_61_v49_)[default__.safeIndex(d_2_v0_, len(d_61_v49_))]})
                        d_63_v51_: _dafny.Seq
                        d_63_v51_ = _dafny.SeqWithoutIsStrInference([p2])
                        d_64_v52_: _dafny.Map
                        d_64_v52_ = _dafny.Map({len(d_63_v51_): d_60_v48_})
                        rhs9_ = (0) - (d_2_v0_)
                        rhs10_ = ((d_62_v50_)[d_57_v47_] if (d_57_v47_) in (d_62_v50_) else ((d_64_v52_)[768] if (768) in (d_64_v52_) else d_60_v48_))
                        lhs6_ = globalState
                        lhs6_.f5 = rhs9_
                        r0 = rhs10_
                        d_65_v53_: _dafny.Array
                        def lambda7_(d_66_v45_):
                            def lambda8_(d_67_i6_):
                                return d_66_v45_

                            return lambda8_

                        init5_ = lambda7_(d_55_v45_)
                        nw15_ = _dafny.Array(None, 10)
                        for i0_5_ in range(nw15_.length(0)):
                            nw15_[i0_5_] = init5_(i0_5_)
                        d_65_v53_ = nw15_
                        d_65_v53_ = d_65_v53_
                    (globalState).f12 = not(True)
                    d_68_v54_: _dafny.Array
                    def lambda9_(d_69_p2_):
                        def lambda10_(d_70_i7_):
                            return d_69_p2_

                        return lambda10_

                    init6_ = lambda9_(p2)
                    nw16_ = _dafny.Array(None, 8)
                    for i0_6_ in range(nw16_.length(0)):
                        nw16_[i0_6_] = init6_(i0_6_)
                    d_68_v54_ = nw16_
                    def lambda11_(d_71_p2_, d_72_p0_, d_73_v3_, d_74_p1_):
                        def lambda12_(d_75_i8_):
                            return (((d_73_v3_)[len(_dafny.SeqWithoutIsStrInference([d_72_p0_]))] if (len(_dafny.SeqWithoutIsStrInference([d_72_p0_]))) in (d_73_v3_) else d_74_p1_) if d_71_p2_ else d_74_p1_)

                        return lambda12_

                    init7_ = lambda11_(p2, p0, d_5_v3_, p1)
                    nw17_ = _dafny.Array(None, 23)
                    for i0_7_ in range(nw17_.length(0)):
                        nw17_[i0_7_] = init7_(i0_7_)
                    d_68_v54_ = nw17_
                    d_76_v55_: _dafny.Map
                    d_76_v55_ = _dafny.Map({d_68_v54_: p0})
                    (globalState).f5 = (d_2_v0_) * (len((d_76_v55_) | (d_76_v55_)))
                    pass
            pass
        d_77_v56_: str
        d_77_v56_ = _dafny.CodePoint('w')
        d_78_v57_: _dafny.Seq
        d_78_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rwuaf"))
        if not(not(((d_77_v56_) not in (d_78_v57_) if p0 else p2))):
            d_79_v58_: _dafny.Array
            nw18_ = _dafny.Array(None, 26)
            d_79_v58_ = nw18_
            d_79_v58_ = d_79_v58_
            d_80_v59_: _dafny.Set
            d_80_v59_ = _dafny.Set({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_2_v0_]))).cardinality})
            d_81_v60_: _dafny.Map
            d_81_v60_ = _dafny.Map({p0: d_80_v59_})
            d_81_v60_ = (d_81_v60_).set(p2, d_80_v59_)
            d_82_v61_: C0
            nw19_ = C0()
            nw19_.ctor__()
            d_82_v61_ = nw19_
            (globalState).f12 = p0
            d_83_v62_: _dafny.Seq
            d_83_v62_ = _dafny.SeqWithoutIsStrInference([d_2_v0_, d_2_v0_, 179, (0) - (d_2_v0_)])
            d_84_v63_: _dafny.Map
            d_84_v63_ = _dafny.Map({d_83_v62_: not(not (not(p0)) or (p2))})
            d_84_v63_ = (d_84_v63_).set((d_83_v62_) + (d_83_v62_), p2)
        elif True:
            d_78_v57_ = d_78_v57_
            d_85_v64_: _dafny.MultiSet
            d_85_v64_ = _dafny.MultiSet([p1])
            d_86_v65_: _dafny.Set
            d_86_v65_ = _dafny.Set({d_2_v0_})
            d_87_v66_: _dafny.MultiSet
            d_87_v66_ = _dafny.MultiSet([len(d_78_v57_)])
            d_88_v67_: _dafny.MultiSet
            d_88_v67_ = _dafny.MultiSet([d_87_v66_, d_87_v66_])
            d_89_v68_: _dafny.Map
            d_89_v68_ = _dafny.Map({d_88_v67_: d_4_v2_})
            d_90_v69_: _dafny.Seq
            d_90_v69_ = _dafny.SeqWithoutIsStrInference([d_2_v0_])
            d_91_v70_: _dafny.Array
            nw20_ = _dafny.Array(None, 19)
            nw20_[int(0)] = d_2_v0_
            nw20_[int(1)] = d_2_v0_
            nw20_[int(2)] = d_2_v0_
            nw20_[int(3)] = d_2_v0_
            nw20_[int(4)] = d_2_v0_
            nw20_[int(5)] = d_2_v0_
            nw20_[int(6)] = d_2_v0_
            nw20_[int(7)] = d_2_v0_
            nw20_[int(8)] = (d_85_v64_).cardinality
            nw20_[int(9)] = len(d_78_v57_)
            nw20_[int(10)] = d_2_v0_
            nw20_[int(11)] = d_2_v0_
            nw20_[int(12)] = (0) - (len(d_90_v69_))
            nw20_[int(13)] = d_2_v0_
            nw20_[int(14)] = d_2_v0_
            nw20_[int(15)] = d_2_v0_
            nw20_[int(16)] = d_2_v0_
            nw20_[int(17)] = d_2_v0_
            nw20_[int(18)] = d_2_v0_
            d_91_v70_ = nw20_
            d_92_v71_: _dafny.Seq
            d_92_v71_ = _dafny.SeqWithoutIsStrInference([d_91_v70_, d_91_v70_])
            d_93_v72_: _dafny.Seq
            d_93_v72_ = _dafny.SeqWithoutIsStrInference([(d_92_v71_)[default__.safeIndex(d_2_v0_, len(d_92_v71_))]])
            d_94_v73_: bool
            d_94_v73_ = False
            d_95_v74_: bool
            d_95_v74_ = (d_94_v73_)
            d_96_v75_: C0
            nw21_ = C0()
            nw21_.ctor__()
            d_96_v75_ = nw21_
            d_97_v76_: _dafny.Map
            d_97_v76_ = _dafny.Map({d_96_v75_: d_78_v57_})
            d_98_v77_: _dafny.Array
            nw22_ = _dafny.Array(None, 25)
            nw22_[int(0)] = len(_dafny.SeqWithoutIsStrInference([p2]))
            nw22_[int(1)] = (d_85_v64_).cardinality
            nw22_[int(2)] = d_2_v0_
            nw22_[int(3)] = d_2_v0_
            nw22_[int(4)] = -713
            nw22_[int(5)] = default__.fm1(p0, d_4_v2_, globalState)
            nw22_[int(6)] = (d_2_v0_ if p2 else 144)
            nw22_[int(7)] = len((d_86_v65_).intersection(d_86_v65_))
            nw22_[int(8)] = d_2_v0_
            nw22_[int(9)] = d_2_v0_
            nw22_[int(10)] = default__.fm1(default__.fm2(661, p0, globalState), ((d_89_v68_)[d_88_v67_] if (d_88_v67_) in (d_89_v68_) else d_4_v2_), globalState)
            nw22_[int(11)] = d_2_v0_
            nw22_[int(12)] = len(d_93_v72_)
            nw22_[int(13)] = d_2_v0_
            nw22_[int(14)] = (d_2_v0_ if (d_94_v73_) else d_2_v0_)
            nw22_[int(15)] = d_2_v0_
            nw22_[int(16)] = d_2_v0_
            nw22_[int(17)] = len(d_90_v69_)
            nw22_[int(18)] = len((d_90_v69_) + (d_90_v69_))
            nw22_[int(19)] = d_2_v0_
            nw22_[int(20)] = len(d_97_v76_)
            nw22_[int(21)] = d_2_v0_
            nw22_[int(22)] = d_2_v0_
            nw22_[int(23)] = len(d_3_v1_)
            nw22_[int(24)] = d_2_v0_
            d_98_v77_ = nw22_
            rhs11_ = d_98_v77_
            rhs12_ = d_77_v56_
            d_91_v70_ = rhs11_
            d_77_v56_ = rhs12_
            d_99_v78_: _dafny.Map
            d_99_v78_ = _dafny.Map({p1: p1})
            d_100_v79_: _dafny.Map
            d_100_v79_ = _dafny.Map({_dafny.Set({d_2_v0_, d_2_v0_, d_2_v0_}): default__.fm0(((d_85_v64_)[p2] if (p2) in (d_85_v64_) else d_2_v0_), p1, d_99_v78_, globalState)})
            d_100_v79_ = (d_100_v79_).set((_dafny.Set({d_2_v0_, d_2_v0_})).intersection(_dafny.Set({d_2_v0_, d_2_v0_})), (d_78_v57_)[default__.safeIndex(d_2_v0_, len(d_78_v57_))])
            d_85_v64_ = d_85_v64_
            d_101_v80_: C0
            nw23_ = C0()
            nw23_.ctor__()
            d_101_v80_ = nw23_
        d_102_v81_: _dafny.Map
        d_102_v81_ = _dafny.Map({True: p0})
        d_103_v82_: _dafny.Map
        d_103_v82_ = d_102_v81_
        pat_let_tv0_ = p2
        def lambda13_(source1_):
            d_104___mcc_h1_ = source1_
            d_105_cf1_ = d_104___mcc_h1_
            return pat_let_tv0_

        rhs13_ = lambda13_(d_103_v82_)
        rhs14_ = p0
        lhs7_ = globalState
        lhs8_ = globalState
        lhs7_.f7 = rhs13_
        lhs8_.f12 = rhs14_
        hi0_ = len((d_78_v57_) + (d_78_v57_))
        for d_106_i9_ in range(len(d_3_v1_), hi0_):
            (globalState).f3 = d_2_v0_
            d_107_v83_: _dafny.Seq
            d_107_v83_ = _dafny.SeqWithoutIsStrInference([p2])
            d_108_v84_: bool
            d_108_v84_ = p1
            d_109_v85_: _dafny.Map
            d_109_v85_ = _dafny.Map({p1: d_107_v83_})
            d_110_v86_: _dafny.Array
            nw24_ = _dafny.Array(None, 25)
            nw24_[int(0)] = (d_107_v83_) + (_dafny.SeqWithoutIsStrInference([p0, p2]))
            nw24_[int(1)] = d_107_v83_
            nw24_[int(2)] = (d_107_v83_) + (d_107_v83_)
            nw24_[int(3)] = d_107_v83_
            nw24_[int(4)] = (d_107_v83_) + (d_107_v83_)
            nw24_[int(5)] = d_107_v83_
            nw24_[int(6)] = _dafny.SeqWithoutIsStrInference([p1])
            nw24_[int(7)] = d_107_v83_
            nw24_[int(8)] = (d_107_v83_) + (_dafny.SeqWithoutIsStrInference([p0, default__.fm2(d_106_i9_, p0, globalState)]))
            nw24_[int(9)] = d_107_v83_
            nw24_[int(10)] = d_107_v83_
            nw24_[int(11)] = d_107_v83_
            nw24_[int(12)] = d_107_v83_
            nw24_[int(13)] = (d_107_v83_).set(default__.safeIndex(d_106_i9_, len(d_107_v83_)), p0)
            nw24_[int(14)] = d_107_v83_
            nw24_[int(15)] = (((d_107_v83_) + (d_107_v83_)).set(default__.safeIndex(d_2_v0_, len((d_107_v83_) + (d_107_v83_))), not(p0))).set(default__.safeIndex(d_2_v0_, len(((d_107_v83_) + (d_107_v83_)).set(default__.safeIndex(d_2_v0_, len((d_107_v83_) + (d_107_v83_))), not(p0)))), p1)
            nw24_[int(16)] = default__.fm9((d_108_v84_), globalState)
            nw24_[int(17)] = d_107_v83_
            nw24_[int(18)] = d_107_v83_
            nw24_[int(19)] = _dafny.SeqWithoutIsStrInference([p0, p1, not(not(p0)), p2])
            nw24_[int(20)] = d_107_v83_
            nw24_[int(21)] = _dafny.SeqWithoutIsStrInference([p1])
            nw24_[int(22)] = ((d_107_v83_) + (d_107_v83_)).set(default__.safeIndex(d_2_v0_, len((d_107_v83_) + (d_107_v83_))), p2)
            nw24_[int(23)] = (d_107_v83_) + (d_107_v83_)
            nw24_[int(24)] = ((d_109_v85_)[((d_5_v3_)[d_106_i9_] if (d_106_i9_) in (d_5_v3_) else p1)] if (((d_5_v3_)[d_106_i9_] if (d_106_i9_) in (d_5_v3_) else p1)) in (d_109_v85_) else _dafny.SeqWithoutIsStrInference([p0]))
            d_110_v86_ = nw24_
            index12_ = default__.safeIndex(295, (d_110_v86_).length(0))
            (d_110_v86_)[index12_] = d_107_v83_
            index13_ = default__.safeIndex(295, (d_110_v86_).length(0))
            rhs15_ = (d_107_v83_) + (d_107_v83_)
            rhs16_ = True
            rhs17_ = p1
            lhs9_ = d_110_v86_
            lhs10_ = default__.safeIndex(295, (d_110_v86_).length(0))
            lhs11_ = globalState
            lhs12_ = globalState
            lhs9_[lhs10_] = rhs15_
            lhs11_.f12 = rhs16_
            lhs12_.f7 = rhs17_
            (globalState).f3 = d_106_i9_
            (globalState).f5 = d_106_i9_
        r0 = d_102_v81_
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_111_v0_: _dafny.Seq
        d_111_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "olqswci"))
        d_112_v1_: bool
        d_112_v1_ = True
        d_113_v2_: int
        d_113_v2_ = -891
        d_114_globalState_: GlobalState
        nw25_ = GlobalState()
        nw25_.ctor__(False, 115, True, -895, d_111_v0_, 403, 751, False, d_111_v0_, 958, (_dafny.Map({d_112_v1_: d_113_v2_})).set(d_112_v1_, d_113_v2_), True, False, False, d_111_v0_)
        d_114_globalState_ = nw25_
        if d_112_v1_:
            d_115_v3_: _dafny.Set
            d_115_v3_ = _dafny.Set({(d_112_v1_), d_112_v1_, d_112_v1_, d_112_v1_, d_112_v1_})
            (d_114_globalState_).f5 = ((d_113_v2_) - (len(d_115_v3_)) if True else d_113_v2_)
            d_116_v4_: _dafny.MultiSet
            d_116_v4_ = _dafny.MultiSet([-807])
            d_112_v1_ = (d_116_v4_).ispropersubset((d_116_v4_).set(d_113_v2_, default__.abs(837)))
            d_117_v5_: _dafny.Array
            nw26_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 19)
            d_117_v5_ = nw26_
            index14_ = default__.safeIndex(379, (d_117_v5_).length(0))
            (d_117_v5_)[index14_] = d_111_v0_
            d_118_v6_: bool
            d_118_v6_ = d_112_v1_
            d_119_v7_: _dafny.Seq
            d_119_v7_ = _dafny.SeqWithoutIsStrInference([d_112_v1_, d_112_v1_])
            index15_ = default__.safeIndex(379, (d_117_v5_).length(0))
            rhs18_ = (d_119_v7_) < (d_119_v7_)
            rhs19_ = d_113_v2_
            rhs20_ = (d_111_v0_) + (d_111_v0_)
            rhs21_ = d_118_v6_
            rhs22_ = default__.safeModuloInt((d_113_v2_) + (d_113_v2_), default__.safeDivisionInt(d_113_v2_, d_113_v2_))
            lhs13_ = d_114_globalState_
            lhs14_ = d_114_globalState_
            lhs15_ = d_117_v5_
            lhs16_ = default__.safeIndex(379, (d_117_v5_).length(0))
            lhs17_ = d_114_globalState_
            lhs13_.f12 = rhs18_
            lhs14_.f3 = rhs19_
            lhs15_[lhs16_] = rhs20_
            d_118_v6_ = rhs21_
            lhs17_.f5 = rhs22_
            d_120_v8_: _dafny.Array
            def lambda14_(d_121_v1_):
                def lambda15_(d_122_i0_):
                    return (_dafny.MultiSet([d_121_v1_])).issubset(_dafny.MultiSet([d_121_v1_]))

                return lambda15_

            init8_ = lambda14_(d_112_v1_)
            nw27_ = _dafny.Array(None, 7)
            for i0_8_ in range(nw27_.length(0)):
                nw27_[i0_8_] = init8_(i0_8_)
            d_120_v8_ = nw27_
            index16_ = default__.safeIndex(625, (d_120_v8_).length(0))
            (d_120_v8_)[index16_] = d_112_v1_
            index17_ = default__.safeIndex(625, (d_120_v8_).length(0))
            (d_120_v8_)[index17_] = d_112_v1_
            d_123_v9_: _dafny.Array
            nw28_ = _dafny.Array(int(0), 12)
            d_123_v9_ = nw28_
            index18_ = default__.safeIndex(699, (d_123_v9_).length(0))
            (d_123_v9_)[index18_] = d_113_v2_
            index19_ = default__.safeIndex(699, (d_123_v9_).length(0))
            rhs23_ = d_113_v2_
            rhs24_ = 670
            lhs18_ = d_123_v9_
            lhs19_ = default__.safeIndex(699, (d_123_v9_).length(0))
            lhs20_ = d_114_globalState_
            lhs18_[lhs19_] = rhs23_
            lhs20_.f3 = rhs24_
        elif True:
            d_124_v10_: _dafny.Map
            out0_: _dafny.Map
            out0_ = default__.m0(True, d_112_v1_, d_112_v1_, d_114_globalState_)
            d_124_v10_ = out0_
            d_125_v11_: _dafny.Array
            def lambda16_(d_126_v1_):
                def lambda17_(d_127_i1_):
                    return d_126_v1_

                return lambda17_

            init9_ = lambda16_(d_112_v1_)
            nw29_ = _dafny.Array(None, 6)
            for i0_9_ in range(nw29_.length(0)):
                nw29_[i0_9_] = init9_(i0_9_)
            d_125_v11_ = nw29_
            d_128_v12_: bool
            d_128_v12_ = (d_112_v1_)
            index20_ = default__.safeIndex(38, (d_125_v11_).length(0))
            (d_125_v11_)[index20_] = d_128_v12_
            index21_ = default__.safeIndex(38, (d_125_v11_).length(0))
            (d_125_v11_)[index21_] = (d_128_v12_ if d_112_v1_ else d_128_v12_)
            d_129_v13_: str
            d_129_v13_ = _dafny.CodePoint('e')
            d_130_v14_: _dafny.Map
            d_130_v14_ = d_124_v10_
            d_129_v13_ = default__.fm0(d_113_v2_, d_112_v1_, (d_130_v14_), d_114_globalState_)
            if True:
                d_131_v15_: _dafny.Seq
                d_131_v15_ = _dafny.SeqWithoutIsStrInference([True])
                d_132_v16_: _dafny.Set
                d_132_v16_ = _dafny.Set({d_112_v1_})
                d_133_v17_: _dafny.Set
                d_133_v17_ = _dafny.Set({default__.fm0(len(d_132_v16_), d_112_v1_, d_124_v10_, d_114_globalState_), d_129_v13_})
                d_134_v18_: _dafny.Map
                d_134_v18_ = _dafny.Map({(d_111_v0_)[default__.safeIndex(852, len(d_111_v0_))]: d_113_v2_})
                d_135_v20_: _dafny.Map
                d_135_v20_ = _dafny.Map({not(d_112_v1_): d_113_v2_})
                d_136_v21_: _dafny.Map
                d_136_v21_ = _dafny.Map({_dafny.Map({_dafny.Map({d_112_v1_: d_112_v1_}): d_113_v2_}): d_113_v2_})
                d_137_v22_: _dafny.Map
                d_137_v22_ = _dafny.Map({d_124_v10_: len(d_131_v15_)})
                d_138_v23_: _dafny.Seq
                d_138_v23_ = _dafny.SeqWithoutIsStrInference([d_113_v2_])
                d_139_v24_: _dafny.Array
                nw30_ = _dafny.Array(None, 26)
                nw30_[int(0)] = len((d_131_v15_) + (d_131_v15_))
                nw30_[int(1)] = d_113_v2_
                nw30_[int(2)] = default__.safeModuloInt(d_113_v2_, d_113_v2_)
                nw30_[int(3)] = d_113_v2_
                nw30_[int(4)] = 737
                def iife1_():
                    coll1_ = _dafny.Set()
                    compr_1_: str
                    for compr_1_ in (d_134_v18_).keys.Elements:
                        d_140_v19_: str = compr_1_
                        if (d_140_v19_) in (d_134_v18_):
                            coll1_ = coll1_.union(_dafny.Set([d_140_v19_]))
                    return _dafny.Set(coll1_)
                nw30_[int(5)] = len((d_133_v17_ if d_112_v1_ else iife1_()
                ))
                nw30_[int(6)] = d_113_v2_
                nw30_[int(7)] = d_113_v2_
                nw30_[int(8)] = d_113_v2_
                nw30_[int(9)] = d_113_v2_
                nw30_[int(10)] = default__.fm1(not(default__.fm2(d_113_v2_, d_112_v1_, d_114_globalState_)), d_135_v20_, d_114_globalState_)
                nw30_[int(11)] = d_113_v2_
                nw30_[int(12)] = d_113_v2_
                nw30_[int(13)] = d_113_v2_
                nw30_[int(14)] = (d_113_v2_) - (d_113_v2_)
                nw30_[int(15)] = d_113_v2_
                nw30_[int(16)] = default__.fm1(d_112_v1_, _dafny.Map({d_112_v1_: d_113_v2_}), d_114_globalState_)
                nw30_[int(17)] = ((d_136_v21_)[d_137_v22_] if (d_137_v22_) in (d_136_v21_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_141_i2_ in range(default__.abs(847))])))
                nw30_[int(18)] = (0) - (d_113_v2_)
                nw30_[int(19)] = -851
                nw30_[int(20)] = d_113_v2_
                nw30_[int(21)] = d_113_v2_
                nw30_[int(22)] = (d_113_v2_) * (d_113_v2_)
                nw30_[int(23)] = d_113_v2_
                nw30_[int(24)] = ((d_138_v23_)[default__.safeIndex(default__.fm1(d_112_v1_, d_135_v20_, d_114_globalState_), len(d_138_v23_))] if d_112_v1_ else (0) - (d_113_v2_))
                nw30_[int(25)] = (d_113_v2_) + (d_113_v2_)
                d_139_v24_ = nw30_
                index22_ = default__.safeIndex(86, (d_139_v24_).length(0))
                (d_139_v24_)[index22_] = d_113_v2_
                index23_ = default__.safeIndex(86, (d_139_v24_).length(0))
                (d_139_v24_)[index23_] = len(_dafny.SeqWithoutIsStrInference([d_129_v13_ for d_142_i3_ in range(default__.abs(115))]))
                d_143_v25_: _dafny.Array
                nw31_ = _dafny.Array(False, 23)
                d_143_v25_ = nw31_
                d_144_v26_: _dafny.Map
                d_144_v26_ = _dafny.Map({(d_139_v24_)[default__.safeIndex(86, (d_139_v24_).length(0))]: not(True)})
                index24_ = default__.safeIndex(400, (d_143_v25_).length(0))
                (d_143_v25_)[index24_] = ((d_144_v26_)[d_113_v2_] if (d_113_v2_) in (d_144_v26_) else d_112_v1_)
                index25_ = default__.safeIndex(400, (d_143_v25_).length(0))
                (d_143_v25_)[index25_] = d_112_v1_
                rhs25_ = (default__.fm3(d_112_v1_, (d_139_v24_)[default__.safeIndex(86, (d_139_v24_).length(0))], d_114_globalState_) if (d_143_v25_)[default__.safeIndex(400, (d_143_v25_).length(0))] else ((d_111_v0_).set(default__.safeIndex(d_113_v2_, len(d_111_v0_)), d_129_v13_)) + (d_111_v0_))
                rhs26_ = (d_111_v0_)[default__.safeIndex(d_113_v2_, len(d_111_v0_))]
                rhs27_ = (d_143_v25_)[default__.safeIndex(400, (d_143_v25_).length(0))]
                rhs28_ = d_113_v2_
                rhs29_ = d_112_v1_
                lhs21_ = d_114_globalState_
                lhs22_ = d_114_globalState_
                d_111_v0_ = rhs25_
                d_129_v13_ = rhs26_
                lhs21_.f12 = rhs27_
                lhs22_.f5 = rhs28_
                d_112_v1_ = rhs29_
                d_145_v27_: _dafny.Map
                d_145_v27_ = _dafny.Map({((d_143_v25_)[default__.safeIndex(400, (d_143_v25_).length(0))]) == (default__.fm2(len(d_111_v0_), d_112_v1_, d_114_globalState_)): d_111_v0_})
                d_145_v27_ = d_145_v27_
                d_146_v28_: _dafny.Seq
                d_146_v28_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_139_v24_)[default__.safeIndex(86, (d_139_v24_).length(0))]]), d_138_v23_, _dafny.SeqWithoutIsStrInference([(d_139_v24_)[default__.safeIndex(86, (d_139_v24_).length(0))]]), (d_138_v23_).set(default__.safeIndex((d_139_v24_)[default__.safeIndex(86, (d_139_v24_).length(0))], len(d_138_v23_)), d_113_v2_), d_138_v23_])
                d_147_v29_: _dafny.MultiSet
                d_147_v29_ = _dafny.MultiSet([(d_146_v28_)[default__.safeIndex((d_139_v24_)[default__.safeIndex(86, (d_139_v24_).length(0))], len(d_146_v28_))]])
                d_148_v30_: _dafny.Map
                out1_: _dafny.Map
                out1_ = default__.m0((d_143_v25_)[default__.safeIndex(400, (d_143_v25_).length(0))], (d_143_v25_)[default__.safeIndex(400, (d_143_v25_).length(0))], not((d_147_v29_) == (d_147_v29_)), d_114_globalState_)
                d_148_v30_ = out1_
            elif True:
                d_149_v31_: C0
                nw32_ = C0()
                nw32_.ctor__()
                d_149_v31_ = nw32_
                d_130_v14_ = d_130_v14_
                d_150_v32_: _dafny.Array
                def lambda18_(d_151_i4_):
                    return (d_151_i4_) * (-925)

                init10_ = lambda18_
                nw33_ = _dafny.Array(None, 13)
                for i0_10_ in range(nw33_.length(0)):
                    nw33_[i0_10_] = init10_(i0_10_)
                d_150_v32_ = nw33_
                d_152_v33_: _dafny.Seq
                d_152_v33_ = _dafny.SeqWithoutIsStrInference([d_150_v32_, d_150_v32_, d_150_v32_])
                d_153_v34_: _dafny.Map
                d_153_v34_ = _dafny.Map({d_124_v10_: ((d_152_v33_)[default__.safeIndex(len((d_111_v0_).set(default__.safeIndex(len(_dafny.Map({d_112_v1_: d_113_v2_})), len(d_111_v0_)), d_129_v13_)), len(d_152_v33_))] if d_112_v1_ else d_150_v32_)})
                d_153_v34_ = ((d_153_v34_) | (d_153_v34_)) | (d_153_v34_)
                d_154_v35_: _dafny.MultiSet
                d_154_v35_ = _dafny.MultiSet([d_112_v1_])
                d_155_v36_: _dafny.Map
                d_155_v36_ = _dafny.Map({d_113_v2_: default__.safeModuloInt(d_113_v2_, ((d_154_v35_)[False] if (False) in (d_154_v35_) else d_113_v2_))})
                d_156_v37_: _dafny.Map
                d_156_v37_ = _dafny.Map({d_112_v1_: d_113_v2_})
                d_155_v36_ = (d_155_v36_).set(default__.fm1(d_112_v1_, d_156_v37_, d_114_globalState_), d_113_v2_)
                d_157_v38_: _dafny.Array
                def lambda19_(d_158_v0_):
                    def lambda20_(d_159_i5_):
                        return (d_158_v0_) + (d_158_v0_)

                    return lambda20_

                init11_ = lambda19_(d_111_v0_)
                nw34_ = _dafny.Array(None, 22)
                for i0_11_ in range(nw34_.length(0)):
                    nw34_[i0_11_] = init11_(i0_11_)
                d_157_v38_ = nw34_
                index26_ = default__.safeIndex(205, (d_157_v38_).length(0))
                (d_157_v38_)[index26_] = (d_111_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hydejw")))
                index27_ = default__.safeIndex(205, (d_157_v38_).length(0))
                (d_157_v38_)[index27_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_160_i6_ in range(default__.abs(255))])
            (d_114_globalState_).f7 = d_112_v1_
        d_161_i7_: int
        d_161_i7_ = 0
        with _dafny.label("1"):
            while True:
                with _dafny.c_label("1"):
                    if (d_161_i7_) >= (100):
                        raise _dafny.Break("1")
                    d_161_i7_ = (d_161_i7_) + (1)
                    d_162_v39_: _dafny.Set
                    d_162_v39_ = _dafny.Set({d_112_v1_, False, d_112_v1_, d_112_v1_})
                    d_163_v40_: _dafny.MultiSet
                    d_163_v40_ = _dafny.MultiSet([len(d_111_v0_)])
                    d_164_v41_: _dafny.Seq
                    d_164_v41_ = _dafny.SeqWithoutIsStrInference([d_113_v2_])
                    d_165_v42_: _dafny.Array
                    nw35_ = _dafny.Array(None, 16)
                    nw35_[int(0)] = d_112_v1_
                    nw35_[int(1)] = (d_112_v1_ if d_112_v1_ else d_112_v1_)
                    nw35_[int(2)] = False
                    nw35_[int(3)] = d_112_v1_
                    nw35_[int(4)] = d_112_v1_
                    nw35_[int(5)] = d_112_v1_
                    nw35_[int(6)] = ((0) - (len(d_162_v39_))) > (d_113_v2_)
                    nw35_[int(7)] = (d_163_v40_).isdisjoint((_dafny.MultiSet([d_113_v2_, d_113_v2_])).set(d_113_v2_, default__.abs(d_113_v2_)))
                    nw35_[int(8)] = not((d_164_v41_) < (d_164_v41_))
                    nw35_[int(9)] = d_112_v1_
                    nw35_[int(10)] = d_112_v1_
                    nw35_[int(11)] = d_112_v1_
                    nw35_[int(12)] = d_112_v1_
                    nw35_[int(13)] = d_112_v1_
                    nw35_[int(14)] = (d_112_v1_) in (_dafny.Map({d_112_v1_: d_112_v1_}))
                    nw35_[int(15)] = d_112_v1_
                    d_165_v42_ = nw35_
                    d_166_v43_: _dafny.Seq
                    d_166_v43_ = _dafny.SeqWithoutIsStrInference([d_112_v1_])
                    d_167_v44_: _dafny.Array
                    nw36_ = _dafny.Array(int(0), 8)
                    d_167_v44_ = nw36_
                    rhs30_ = d_165_v42_
                    rhs31_ = d_166_v43_
                    rhs32_ = d_167_v44_
                    rhs33_ = (d_113_v2_) - ((d_113_v2_) - (d_113_v2_))
                    d_165_v42_ = rhs30_
                    d_166_v43_ = rhs31_
                    d_167_v44_ = rhs32_
                    d_113_v2_ = rhs33_
                    index28_ = default__.safeIndex(232, (d_165_v42_).length(0))
                    (d_165_v42_)[index28_] = d_112_v1_
                    index29_ = default__.safeIndex(232, (d_165_v42_).length(0))
                    (d_165_v42_)[index29_] = (d_163_v40_).issubset(_dafny.MultiSet([len(d_111_v0_)]))
                    (d_114_globalState_).f12 = (d_165_v42_)[default__.safeIndex(232, (d_165_v42_).length(0))]
                    d_168_v45_: int
                    d_168_v45_ = len(_dafny.Map({(d_165_v42_)[default__.safeIndex(232, (d_165_v42_).length(0))]: default__.fm2(d_113_v2_, d_112_v1_, d_114_globalState_)}))
                    d_169_v46_: _dafny.Set
                    d_169_v46_ = _dafny.Set({(d_168_v45_), (0) - (d_113_v2_), d_113_v2_})
                    d_170_v47_: _dafny.Map
                    d_170_v47_ = _dafny.Map({(0) - (d_113_v2_): d_113_v2_})
                    d_169_v46_ = (d_169_v46_) - ((_dafny.Set({len(d_170_v47_), d_113_v2_})) - (d_169_v46_))
                    pass
            pass
        d_171_v48_: _dafny.Array
        def lambda21_(d_172_v1_):
            def lambda22_(d_173_i8_):
                return not(d_172_v1_)

            return lambda22_

        init12_ = lambda21_(d_112_v1_)
        nw37_ = _dafny.Array(None, 11)
        for i0_12_ in range(nw37_.length(0)):
            nw37_[i0_12_] = init12_(i0_12_)
        d_171_v48_ = nw37_
        d_174_v49_: _dafny.Map
        d_174_v49_ = _dafny.Map({d_171_v48_: d_112_v1_})
        d_174_v49_ = (d_174_v49_).set(d_171_v48_, d_112_v1_)
        hi1_ = d_113_v2_
        for d_175_i9_ in range(d_113_v2_, hi1_):
            d_176_v50_: C0
            nw38_ = C0()
            nw38_.ctor__()
            d_176_v50_ = nw38_
            d_177_v51_: bool
            d_177_v51_ = d_112_v1_
            d_178_v52_: _dafny.Set
            d_178_v52_ = _dafny.Set({d_112_v1_})
            rhs34_ = d_178_v52_
            rhs35_ = d_112_v1_
            lhs23_ = d_114_globalState_
            d_178_v52_ = rhs34_
            lhs23_.f7 = rhs35_
            d_179_v53_: _dafny.Map
            out2_: _dafny.Map
            out2_ = default__.m0(False, d_112_v1_, d_112_v1_, d_114_globalState_)
            d_179_v53_ = out2_
            d_180_v54_: _dafny.MultiSet
            d_180_v54_ = _dafny.MultiSet([791])
            d_181_v55_: _dafny.Seq
            d_181_v55_ = _dafny.SeqWithoutIsStrInference([d_180_v54_, _dafny.MultiSet([d_113_v2_, d_113_v2_]), d_180_v54_])
            d_182_v56_: _dafny.Seq
            d_182_v56_ = _dafny.SeqWithoutIsStrInference([d_175_i9_, 736])
            rhs36_ = (858) - (d_113_v2_)
            rhs37_ = d_175_i9_
            rhs38_ = d_112_v1_
            rhs39_ = d_112_v1_
            rhs40_ = ((d_181_v55_)[default__.safeIndex(d_113_v2_, len(d_181_v55_))]).set((d_113_v2_ if d_112_v1_ else (d_182_v56_)[default__.safeIndex(d_113_v2_, len(d_182_v56_))]), default__.abs(d_113_v2_))
            lhs24_ = d_114_globalState_
            lhs25_ = d_114_globalState_
            lhs24_.f5 = rhs36_
            lhs25_.f5 = rhs37_
            d_112_v1_ = rhs38_
            d_112_v1_ = rhs39_
            d_180_v54_ = rhs40_
        d_183_v57_: _dafny.Map
        d_183_v57_ = _dafny.Map({d_112_v1_: d_113_v2_})
        hi2_ = (0) - (default__.fm1(d_112_v1_, d_183_v57_, d_114_globalState_))
        for d_184_i10_ in range(d_113_v2_, hi2_):
            if d_112_v1_:
                d_185_v58_: _dafny.Seq
                d_185_v58_ = _dafny.SeqWithoutIsStrInference([d_111_v0_])
                d_186_v59_: C0
                nw39_ = C0()
                nw39_.ctor__()
                d_186_v59_ = nw39_
                d_187_v60_: _dafny.Set
                d_187_v60_ = _dafny.Set({d_186_v59_})
                d_188_v61_: _dafny.Seq
                d_188_v61_ = _dafny.SeqWithoutIsStrInference([d_187_v60_, d_187_v60_])
                d_189_v62_: _dafny.Seq
                d_189_v62_ = d_111_v0_
                d_190_v63_: bool
                d_190_v63_ = d_112_v1_
                d_191_v64_: str
                d_191_v64_ = _dafny.CodePoint('b')
                d_192_v65_: _dafny.Array
                nw40_ = _dafny.Array(None, 20)
                nw40_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apenmlek"))
                nw40_[int(1)] = default__.fm3(d_112_v1_, (0) - (d_184_i10_), d_114_globalState_)
                nw40_[int(2)] = d_111_v0_
                nw40_[int(3)] = d_111_v0_
                nw40_[int(4)] = d_111_v0_
                nw40_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxraijmpf"))
                nw40_[int(6)] = (d_185_v58_)[default__.safeIndex(len((d_188_v61_)[default__.safeIndex(d_113_v2_, len(d_188_v61_))]), len(d_185_v58_))]
                nw40_[int(7)] = ((d_189_v62_)) + (d_111_v0_)
                nw40_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqhbcmux"))
                nw40_[int(9)] = d_111_v0_
                nw40_[int(10)] = ((d_189_v62_)).set(default__.safeIndex(d_184_i10_, len((d_189_v62_))), _dafny.CodePoint('o'))
                nw40_[int(11)] = d_111_v0_
                nw40_[int(12)] = d_111_v0_
                nw40_[int(13)] = d_111_v0_
                nw40_[int(14)] = d_111_v0_
                nw40_[int(15)] = default__.fm3((d_186_v59_).fm4(d_112_v1_, d_114_globalState_), d_184_i10_, d_114_globalState_)
                nw40_[int(16)] = d_111_v0_
                nw40_[int(17)] = (default__.fm3(not((d_190_v63_)), (0) - (d_113_v2_), d_114_globalState_)).set(default__.safeIndex(d_184_i10_, len(default__.fm3(not((d_190_v63_)), (0) - (d_113_v2_), d_114_globalState_))), d_191_v64_)
                nw40_[int(18)] = _dafny.SeqWithoutIsStrInference([d_191_v64_ for d_193_i11_ in range(default__.abs(247))])
                nw40_[int(19)] = (d_111_v0_) + (d_111_v0_)
                d_192_v65_ = nw40_
                index30_ = default__.safeIndex(113, (d_192_v65_).length(0))
                (d_192_v65_)[index30_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "evqq"))
                index31_ = default__.safeIndex(113, (d_192_v65_).length(0))
                (d_192_v65_)[index31_] = d_111_v0_
                (d_114_globalState_).f7 = d_112_v1_
                d_194_v66_: _dafny.Map
                d_194_v66_ = _dafny.Map({d_113_v2_: d_186_v59_})
                d_195_v67_: _dafny.Seq
                d_195_v67_ = _dafny.SeqWithoutIsStrInference([(d_186_v59_ if d_112_v1_ else ((d_194_v66_)[d_113_v2_] if (d_113_v2_) in (d_194_v66_) else d_186_v59_)), d_186_v59_])
                d_186_v59_ = (d_195_v67_)[default__.safeIndex(d_113_v2_, len(d_195_v67_))]
                d_196_v68_: C0
                nw41_ = C0()
                nw41_.ctor__()
                d_196_v68_ = nw41_
                d_197_v69_: _dafny.Array
                nw42_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_197_v69_ = nw42_
                d_198_v70_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 12)
                d_198_v70_ = nw43_
                index32_ = default__.safeIndex(722, (d_197_v69_).length(0))
                (d_197_v69_)[index32_] = d_198_v70_
                index33_ = default__.safeIndex(722, (d_197_v69_).length(0))
                (d_197_v69_)[index33_] = d_198_v70_
            elif True:
                d_199_v71_: _dafny.Map
                out3_: _dafny.Map
                out3_ = default__.m0(d_112_v1_, d_112_v1_, d_112_v1_, d_114_globalState_)
                d_199_v71_ = out3_
                (d_114_globalState_).f7 = d_112_v1_
                (d_114_globalState_).f7 = (d_111_v0_) < (d_111_v0_)
                d_200_v72_: _dafny.Array
                nw44_ = _dafny.Array(int(0), 5)
                d_200_v72_ = nw44_
                index34_ = default__.safeIndex(28, (d_200_v72_).length(0))
                (d_200_v72_)[index34_] = 1
                index35_ = default__.safeIndex(28, (d_200_v72_).length(0))
                (d_200_v72_)[index35_] = (d_184_i10_) - (default__.safeDivisionInt(d_184_i10_, d_113_v2_))
                (d_114_globalState_).f5 = default__.fm1((d_111_v0_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_201_i12_ in range(default__.abs(170))])), d_183_v57_, d_114_globalState_)
            d_202_v73_: _dafny.Seq
            d_202_v73_ = _dafny.SeqWithoutIsStrInference([d_184_i10_, d_113_v2_])
            d_203_v74_: _dafny.Seq
            d_203_v74_ = _dafny.SeqWithoutIsStrInference([d_202_v73_, d_202_v73_])
            d_204_v75_: _dafny.Map
            out4_: _dafny.Map
            out4_ = default__.m0(False, (d_203_v74_) == (_dafny.SeqWithoutIsStrInference([d_202_v73_, _dafny.SeqWithoutIsStrInference([d_184_i10_, d_184_i10_, d_113_v2_])])), d_112_v1_, d_114_globalState_)
            d_204_v75_ = out4_
            (d_114_globalState_).f12 = d_112_v1_
            d_205_v76_: _dafny.Array
            nw45_ = _dafny.Array(_dafny.Map({}), 4)
            d_205_v76_ = nw45_
            d_206_v77_: _dafny.Array
            nw46_ = _dafny.Array(int(0), 17)
            d_206_v77_ = nw46_
            index36_ = default__.safeIndex(26, (d_206_v77_).length(0))
            (d_206_v77_)[index36_] = (d_184_i10_) * (d_184_i10_)
            d_207_v78_: _dafny.MultiSet
            d_207_v78_ = _dafny.MultiSet([d_112_v1_])
            d_208_v79_: _dafny.MultiSet
            d_208_v79_ = _dafny.MultiSet([((d_207_v78_)[d_112_v1_] if (d_112_v1_) in (d_207_v78_) else d_113_v2_)])
            index37_ = default__.safeIndex(79, (d_206_v77_).length(0))
            (d_206_v77_)[index37_] = (d_208_v79_).cardinality
            d_209_v80_: C0
            nw47_ = C0()
            nw47_.ctor__()
            d_209_v80_ = nw47_
            d_210_v81_: _dafny.Map
            d_210_v81_ = _dafny.Map({d_209_v80_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqq"))})
            index38_ = default__.safeIndex(26, (d_206_v77_).length(0))
            index39_ = default__.safeIndex(79, (d_206_v77_).length(0))
            rhs41_ = (0) - (len(((d_210_v81_)[d_209_v80_] if (d_209_v80_) in (d_210_v81_) else d_111_v0_)))
            rhs42_ = d_205_v76_
            rhs43_ = ((0) - (len(d_111_v0_))) + (d_113_v2_)
            rhs44_ = len(d_111_v0_)
            lhs26_ = d_114_globalState_
            lhs27_ = d_206_v77_
            lhs28_ = default__.safeIndex(26, (d_206_v77_).length(0))
            lhs29_ = d_206_v77_
            lhs30_ = default__.safeIndex(79, (d_206_v77_).length(0))
            lhs26_.f5 = rhs41_
            d_205_v76_ = rhs42_
            lhs27_[lhs28_] = rhs43_
            lhs29_[lhs30_] = rhs44_
        if True:
            d_211_v82_: C0
            nw48_ = C0()
            nw48_.ctor__()
            d_211_v82_ = nw48_
            d_183_v57_ = (d_183_v57_).set(d_112_v1_, default__.safeModuloInt(d_113_v2_, d_113_v2_))
            d_212_v83_: _dafny.MultiSet
            d_212_v83_ = _dafny.MultiSet([d_171_v48_])
            d_213_v84_: _dafny.Seq
            d_213_v84_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uknqeb"))), d_113_v2_])
            d_214_v85_: _dafny.Map
            d_214_v85_ = _dafny.Map({d_213_v84_: (d_212_v83_).set(d_171_v48_, default__.abs(d_113_v2_))})
            d_212_v83_ = ((d_214_v85_)[d_213_v84_] if (d_213_v84_) in (d_214_v85_) else (d_212_v83_) | (d_212_v83_))
            (d_114_globalState_).f7 = (False) or (d_112_v1_)
            d_215_v86_: _dafny.Map
            d_215_v86_ = _dafny.Map({d_113_v2_: d_113_v2_})
            d_216_v87_: _dafny.MultiSet
            d_216_v87_ = _dafny.MultiSet([d_112_v1_])
            d_215_v86_ = (d_215_v86_).set(467, ((d_216_v87_ if not(d_112_v1_) else d_216_v87_)).cardinality)
        elif True:
            d_217_v88_: _dafny.MultiSet
            d_217_v88_ = _dafny.MultiSet([d_112_v1_])
            d_218_v89_: _dafny.Seq
            d_218_v89_ = _dafny.SeqWithoutIsStrInference([d_113_v2_])
            d_219_v90_: _dafny.Set
            d_219_v90_ = _dafny.Set({d_171_v48_})
            d_220_v91_: str
            d_220_v91_ = _dafny.CodePoint('d')
            d_221_v92_: _dafny.Seq
            d_221_v92_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_220_v91_})])
            d_222_v93_: _dafny.Map
            d_222_v93_ = _dafny.Map({d_221_v92_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_223_i14_ in range(default__.abs(-475))])})
            d_224_v94_: _dafny.Seq
            d_224_v94_ = _dafny.SeqWithoutIsStrInference([d_112_v1_])
            d_225_v95_: _dafny.Array
            nw49_ = _dafny.Array(None, 4)
            nw49_[int(0)] = d_113_v2_
            nw49_[int(1)] = d_113_v2_
            nw49_[int(2)] = len(d_224_v94_)
            nw49_[int(3)] = d_113_v2_
            d_225_v95_ = nw49_
            d_226_v96_: _dafny.Map
            d_226_v96_ = _dafny.Map({d_225_v95_: d_112_v1_})
            d_227_v97_: _dafny.Array
            nw50_ = _dafny.Array(None, 26)
            nw50_[int(0)] = d_113_v2_
            nw50_[int(1)] = d_113_v2_
            nw50_[int(2)] = d_113_v2_
            nw50_[int(3)] = d_113_v2_
            nw50_[int(4)] = d_113_v2_
            nw50_[int(5)] = -621
            nw50_[int(6)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_228_i13_ in range(default__.abs(-603))])), (d_217_v88_).cardinality)
            nw50_[int(7)] = default__.fm1(d_112_v1_, _dafny.Map({False: 742}), d_114_globalState_)
            nw50_[int(8)] = d_113_v2_
            nw50_[int(9)] = -310
            nw50_[int(10)] = d_113_v2_
            nw50_[int(11)] = len(_dafny.Set({d_112_v1_}))
            nw50_[int(12)] = len(d_183_v57_)
            nw50_[int(13)] = d_113_v2_
            nw50_[int(14)] = (len(d_218_v89_) if not(d_112_v1_) else (d_218_v89_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oe")))), len(d_218_v89_))])
            nw50_[int(15)] = default__.fm1(d_112_v1_, d_183_v57_, d_114_globalState_)
            nw50_[int(16)] = d_113_v2_
            nw50_[int(17)] = default__.fm1(d_112_v1_, d_183_v57_, d_114_globalState_)
            nw50_[int(18)] = default__.safeDivisionInt(d_113_v2_, d_113_v2_)
            nw50_[int(19)] = d_113_v2_
            nw50_[int(20)] = default__.safeModuloInt(d_113_v2_, len(d_219_v90_))
            nw50_[int(21)] = d_113_v2_
            nw50_[int(22)] = len(((d_222_v93_)[d_221_v92_] if (d_221_v92_) in (d_222_v93_) else d_111_v0_))
            nw50_[int(23)] = d_113_v2_
            nw50_[int(24)] = d_113_v2_
            nw50_[int(25)] = len(d_226_v96_)
            d_227_v97_ = nw50_
            index40_ = default__.safeIndex(598, (d_225_v95_).length(0))
            (d_225_v95_)[index40_] = d_113_v2_
            d_229_v98_: int
            d_229_v98_ = d_113_v2_
            index41_ = default__.safeIndex(598, (d_225_v95_).length(0))
            (d_225_v95_)[index41_] = (d_229_v98_)
            d_230_v99_: _dafny.Set
            d_230_v99_ = _dafny.Set({((d_225_v95_)[default__.safeIndex(598, (d_225_v95_).length(0))] if d_112_v1_ else (d_225_v95_)[default__.safeIndex(598, (d_225_v95_).length(0))])})
            d_231_v100_: bool
            d_231_v100_ = d_112_v1_
            d_232_v101_: _dafny.Array
            nw51_ = _dafny.Array(None, 7)
            nw51_[int(0)] = default__.fm2(d_113_v2_, d_112_v1_, d_114_globalState_)
            nw51_[int(1)] = d_231_v100_
            nw51_[int(2)] = d_231_v100_
            nw51_[int(3)] = d_231_v100_
            nw51_[int(4)] = d_112_v1_
            nw51_[int(5)] = default__.fm5(False, 549, d_112_v1_, d_113_v2_, d_114_globalState_)
            nw51_[int(6)] = d_112_v1_
            d_232_v101_ = nw51_
            index42_ = default__.safeIndex(574, (d_232_v101_).length(0))
            (d_232_v101_)[index42_] = d_231_v100_
            d_233_v102_: C0
            nw52_ = C0()
            nw52_.ctor__()
            d_233_v102_ = nw52_
            d_234_v103_: _dafny.Seq
            d_234_v103_ = _dafny.SeqWithoutIsStrInference([d_233_v102_])
            index43_ = default__.safeIndex(574, (d_232_v101_).length(0))
            rhs45_ = d_230_v99_
            rhs46_ = not (d_112_v1_) or ((_dafny.SeqWithoutIsStrInference([d_233_v102_])) == (d_234_v103_))
            rhs47_ = (default__.fm6(d_112_v1_, d_112_v1_, d_114_globalState_)) != (d_183_v57_)
            rhs48_ = ((_dafny.SeqWithoutIsStrInference([not(d_112_v1_)])) + (d_224_v94_)) < (d_224_v94_)
            lhs31_ = d_114_globalState_
            lhs32_ = d_232_v101_
            lhs33_ = default__.safeIndex(574, (d_232_v101_).length(0))
            d_230_v99_ = rhs45_
            lhs31_.f12 = rhs46_
            lhs32_[lhs33_] = rhs47_
            d_112_v1_ = rhs48_
            d_235_v104_: _dafny.Map
            out5_: _dafny.Map
            out5_ = default__.m0(((d_225_v95_)[default__.safeIndex(598, (d_225_v95_).length(0))]) in (d_218_v89_), d_112_v1_, d_112_v1_, d_114_globalState_)
            d_235_v104_ = out5_
            d_236_v105_: _dafny.Array
            nw53_ = _dafny.Array(_dafny.Set({}), 10)
            d_236_v105_ = nw53_
            d_237_v106_: _dafny.Set
            d_237_v106_ = _dafny.Set({d_112_v1_})
            index44_ = default__.safeIndex(53, (d_236_v105_).length(0))
            (d_236_v105_)[index44_] = d_237_v106_
            index45_ = default__.safeIndex(53, (d_236_v105_).length(0))
            rhs49_ = _dafny.MultiSet([True])
            rhs50_ = (d_237_v106_) - ((d_237_v106_ if d_112_v1_ else d_237_v106_))
            rhs51_ = d_113_v2_
            lhs34_ = d_236_v105_
            lhs35_ = default__.safeIndex(53, (d_236_v105_).length(0))
            lhs36_ = d_114_globalState_
            d_217_v88_ = rhs49_
            lhs34_[lhs35_] = rhs50_
            lhs36_.f3 = rhs51_
            d_238_v107_: C0
            nw54_ = C0()
            nw54_.ctor__()
            d_238_v107_ = nw54_
        d_239_v108_: _dafny.Map
        d_239_v108_ = _dafny.Map({d_112_v1_: d_112_v1_})
        d_240_v109_: _dafny.Set
        d_240_v109_ = _dafny.Set({d_239_v108_, d_239_v108_})
        d_241_v110_: int
        d_241_v110_ = len(d_240_v109_)
        d_242_v111_: _dafny.Array
        nw55_ = _dafny.Array(None, 5)
        nw55_[int(0)] = (d_241_v110_)
        nw55_[int(1)] = default__.fm1(d_112_v1_, d_183_v57_, d_114_globalState_)
        nw55_[int(2)] = d_113_v2_
        nw55_[int(3)] = d_113_v2_
        nw55_[int(4)] = d_113_v2_
        d_242_v111_ = nw55_
        nw56_ = _dafny.Array(int(0), 3)
        d_242_v111_ = nw56_
        d_243_v112_: str
        d_243_v112_ = _dafny.CodePoint('j')
        d_244_v113_: _dafny.Seq
        d_244_v113_ = d_111_v0_
        d_245_v114_: _dafny.Map
        d_245_v114_ = _dafny.Map({(d_243_v112_) in (d_111_v0_): ((d_244_v113_)) + (d_111_v0_)})
        d_246_v115_: _dafny.Map
        d_246_v115_ = _dafny.Map({d_112_v1_: d_112_v1_})
        d_245_v114_ = default__.fm7(d_246_v115_, d_114_globalState_)
        d_112_v1_ = d_112_v1_
        d_247_v116_: _dafny.Map
        out6_: _dafny.Map
        out6_ = default__.m0(d_112_v1_, not(d_112_v1_), d_112_v1_, d_114_globalState_)
        d_247_v116_ = out6_
        pat_let_tv1_ = d_113_v2_
        pat_let_tv2_ = d_112_v1_
        pat_let_tv3_ = d_112_v1_
        def lambda23_(source2_):
            d_248___mcc_h0_ = source2_
            d_249_cf3_ = d_248___mcc_h0_
            return (pat_let_tv1_) + (len(_dafny.Map({pat_let_tv2_: _dafny.SeqWithoutIsStrInference([pat_let_tv3_])})))

        d_113_v2_ = lambda23_(d_244_v113_)
        if not(d_112_v1_):
            (d_114_globalState_).f7 = d_112_v1_
            d_250_v117_: _dafny.Seq
            d_250_v117_ = _dafny.SeqWithoutIsStrInference([d_242_v111_])
            d_242_v111_ = (d_242_v111_ if (d_112_v1_) and (d_112_v1_) else (d_250_v117_)[default__.safeIndex(d_113_v2_, len(d_250_v117_))])
            d_251_v118_: _dafny.Map
            d_251_v118_ = _dafny.Map({(d_113_v2_) != (((d_183_v57_)[d_112_v1_] if (d_112_v1_) in (d_183_v57_) else d_113_v2_)): d_242_v111_})
            d_251_v118_ = (d_251_v118_).set(not(False), d_242_v111_)
            (d_114_globalState_).f3 = d_113_v2_
            (d_114_globalState_).f8 = (((d_111_v0_).set(default__.safeIndex(d_113_v2_, len(d_111_v0_)), d_243_v112_)) + ((_dafny.SeqWithoutIsStrInference([d_243_v112_ for d_252_i15_ in range(default__.abs(272))])).set(default__.safeIndex(d_113_v2_, len(_dafny.SeqWithoutIsStrInference([d_243_v112_ for d_253_i15_ in range(default__.abs(272))]))), d_243_v112_)) if ((d_246_v115_)[d_112_v1_] if (d_112_v1_) in (d_246_v115_) else d_112_v1_) else d_111_v0_)
        elif True:
            d_254_v119_: C0
            nw57_ = C0()
            nw57_.ctor__()
            d_254_v119_ = nw57_
            d_255_v120_: _dafny.Map
            d_255_v120_ = _dafny.Map({d_254_v119_: not(d_112_v1_)})
            rhs52_ = default__.fm2(len(_dafny.Map({d_112_v1_: d_254_v119_})), False, d_114_globalState_)
            rhs53_ = d_254_v119_
            rhs54_ = (d_112_v1_ if d_112_v1_ else ((d_255_v120_)[d_254_v119_] if (d_254_v119_) in (d_255_v120_) else d_112_v1_))
            lhs37_ = d_114_globalState_
            d_112_v1_ = rhs52_
            d_254_v119_ = rhs53_
            lhs37_.f7 = rhs54_
            d_256_v121_: bool
            d_256_v121_ = d_112_v1_
            d_113_v2_ = len(((d_247_v116_) | (d_246_v115_)) | (_dafny.Map({d_112_v1_: (d_256_v121_)})))
            d_257_v122_: _dafny.MultiSet
            d_257_v122_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([d_254_v119_]))])
            d_243_v112_ = default__.fm0((d_257_v122_).cardinality, d_112_v1_, (d_247_v116_) | ((d_247_v116_).set(d_112_v1_, d_112_v1_)), d_114_globalState_)
            (d_114_globalState_).f7 = d_112_v1_
            (d_114_globalState_).f7 = True
        d_258_i16_: int
        d_258_i16_ = 0
        with _dafny.label("2"):
            while d_112_v1_:
                with _dafny.c_label("2"):
                    if (d_258_i16_) >= (100):
                        raise _dafny.Break("2")
                    d_258_i16_ = (d_258_i16_) + (1)
                    d_112_v1_ = d_112_v1_
                    d_259_v123_: C0
                    nw58_ = C0()
                    nw58_.ctor__()
                    d_259_v123_ = nw58_
                    d_259_v123_ = d_259_v123_
                    d_260_v124_: C0
                    nw59_ = C0()
                    nw59_.ctor__()
                    d_260_v124_ = nw59_
                    d_261_v125_: _dafny.Map
                    out7_: _dafny.Map
                    out7_ = default__.m0(d_112_v1_, d_112_v1_, (d_112_v1_) or (d_112_v1_), d_114_globalState_)
                    d_261_v125_ = out7_
                    pass
            pass
        hi3_ = (d_113_v2_) - (d_113_v2_)
        for d_262_i17_ in range(d_113_v2_, hi3_):
            source3_ = d_244_v113_
            d_263___mcc_h1_ = source3_
            d_264_cf3_ = d_263___mcc_h1_
            (d_114_globalState_).f5 = d_262_i17_
            d_183_v57_ = (d_183_v57_).set(not(not (d_112_v1_) or (d_112_v1_)), d_262_i17_)
            index46_ = default__.safeIndex(769, (d_171_v48_).length(0))
            (d_171_v48_)[index46_] = default__.fm2(d_113_v2_, d_112_v1_, d_114_globalState_)
            index47_ = default__.safeIndex(769, (d_171_v48_).length(0))
            (d_171_v48_)[index47_] = d_112_v1_
            (d_114_globalState_).f7 = default__.fm2((d_113_v2_) + (d_262_i17_), not(d_112_v1_), d_114_globalState_)
            d_265_v126_: _dafny.Set
            d_265_v126_ = _dafny.Set({(0) - (d_262_i17_)})
            (d_114_globalState_).f12 = not((d_265_v126_).ispropersubset(_dafny.Set({d_113_v2_})))
            d_266_v127_: C0
            nw60_ = C0()
            nw60_.ctor__()
            d_266_v127_ = nw60_
            d_267_v128_: _dafny.Map
            d_267_v128_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_268_i18_ in range(default__.abs(433))]): d_112_v1_})
            d_269_v129_: _dafny.Seq
            d_269_v129_ = _dafny.SeqWithoutIsStrInference([d_112_v1_])
            rhs55_ = ((d_111_v0_ if ((d_267_v128_)[d_111_v0_] if (d_111_v0_) in (d_267_v128_) else False) else d_111_v0_)) + (d_111_v0_)
            rhs56_ = ((d_183_v57_)[not(True)] if (not(True)) in (d_183_v57_) else len(d_269_v129_))
            rhs57_ = d_266_v127_
            rhs58_ = (d_111_v0_) <= ((d_111_v0_) + (d_111_v0_))
            lhs38_ = d_114_globalState_
            d_111_v0_ = rhs55_
            lhs38_.f5 = rhs56_
            d_266_v127_ = rhs57_
            d_112_v1_ = rhs58_
            d_270_v130_: _dafny.Array
            def lambda24_(d_271_v112_):
                def lambda25_(d_272_i20_):
                    return d_271_v112_

                return lambda25_

            init13_ = lambda24_(d_243_v112_)
            nw61_ = _dafny.Array(None, 27)
            for i0_13_ in range(nw61_.length(0)):
                nw61_[i0_13_] = init13_(i0_13_)
            d_270_v130_ = nw61_
            d_273_v131_: _dafny.Map
            d_273_v131_ = _dafny.Map({d_270_v130_: d_112_v1_})
            index48_ = default__.safeIndex(596, (d_171_v48_).length(0))
            (d_171_v48_)[index48_] = (len(_dafny.SeqWithoutIsStrInference([d_262_i17_ for d_274_i19_ in range(default__.abs(-295))]))) > (len(d_273_v131_))
            index49_ = default__.safeIndex(596, (d_171_v48_).length(0))
            (d_171_v48_)[index49_] = not(((d_113_v2_) <= (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvpiwucvt")))) if d_112_v1_ else False))
        d_112_v1_ = not (True) or (d_112_v1_)
        (d_114_globalState_).f12 = not(((d_111_v0_) + (d_111_v0_)) <= (d_111_v0_))
        _dafny.print((d_111_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_112_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_113_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_114_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_114_globalState_).f4).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_114_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_114_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_114_globalState_.f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_114_globalState_).f10) == (_dafny.Map({True: -891}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_114_globalState_.f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_114_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_114_globalState_).f14).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_161_i7_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_171_v48_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_174_v49_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_183_v57_) == (_dafny.Map({True: 0}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_239_v108_)) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_240_v109_) == (_dafny.Set({_dafny.Map({True: True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v110_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_243_v112_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_244_v113_)).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v114_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('d')])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_246_v115_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v116_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_258_i16_))
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
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)

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
        return lambda: int(0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D2_DC2)

class D2_DC2(D2, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D3_DC3)

class D3_DC3(D3, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC3({self.cf3.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D4_DC4)

class D4_DC4(D4, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC6()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D5_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D5_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D5_DC7)

class D5_DC6(D5, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC5(D5, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC5) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC7(D5, NamedTuple('DC7', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC7({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC7) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f3: int = int(0)
        self.f5: int = int(0)
        self.f7: bool = False
        self.f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self.f12: bool = False
        self._f0: bool = False
        self._f1: int = int(0)
        self._f2: bool = False
        self._f4: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f6: int = int(0)
        self._f9: int = int(0)
        self._f10: _dafny.Map = _dafny.Map({})
        self._f11: bool = False
        self._f13: bool = False
        self._f14: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self).f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14

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
    def f4(self):
        return self._f4
    @property
    def f6(self):
        return self._f6
    @property
    def f9(self):
        return self._f9
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm4(self, p0, globalState):
        return True

