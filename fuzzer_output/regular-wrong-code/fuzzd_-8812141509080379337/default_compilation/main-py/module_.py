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
        return 74

    @staticmethod
    def fm1(p0, p1, globalState):
        return ((151) + ((0) - ((_dafny.MultiSet([True])).cardinality))) > (-621)

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "awbiybj"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))

    @staticmethod
    def fm3(p0, p1, globalState):
        return False

    @staticmethod
    def fm6(globalState):
        return _dafny.CodePoint('f')

    @staticmethod
    def fm7(p0, globalState):
        return ((_dafny.Set({not(True)})).intersection(_dafny.Set({True}))) | (_dafny.Set({True, not(True), not(True)}))

    @staticmethod
    def fm8(globalState):
        return (_dafny.Set({_dafny.CodePoint('d')})) - ((_dafny.Set({_dafny.CodePoint('o')})) - (_dafny.Set({_dafny.CodePoint('s')})))

    @staticmethod
    def fm9(p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Set = _dafny.Set({})
        d_0_v0_: bool
        d_0_v0_ = False
        d_1_i0_: int
        d_1_i0_ = 0
        with _dafny.label("0"):
            while d_0_v0_:
                with _dafny.c_label("0"):
                    if (d_1_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_1_i0_ = (d_1_i0_) + (1)
                    d_2_v1_: str
                    d_2_v1_ = _dafny.CodePoint('q')
                    d_3_v2_: _dafny.Seq
                    d_3_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), d_2_v1_, d_2_v1_])
                    d_4_v3_: _dafny.Map
                    d_4_v3_ = _dafny.Map({d_0_v0_: p0})
                    d_5_v4_: _dafny.Set
                    d_5_v4_ = _dafny.Set({p1, p0, p0, len(d_4_v3_), p0})
                    d_6_v5_: C0
                    nw0_ = C0()
                    nw0_.ctor__(d_3_v2_, (p0) + (p1), d_3_v2_, (p1) in (d_5_v4_))
                    d_6_v5_ = nw0_
                    d_6_v5_ = d_6_v5_
                    d_7_v6_: _dafny.Map
                    d_7_v6_ = _dafny.Map({p1: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpidlcdvp"))) + (d_3_v2_)})
                    d_8_v7_: _dafny.Set
                    d_8_v7_ = _dafny.Set({d_0_v0_})
                    d_3_v2_ = ((d_7_v6_)[len((d_8_v7_ if d_0_v0_ else default__.fm7(p0, globalState)))] if (len((d_8_v7_ if d_0_v0_ else default__.fm7(p0, globalState)))) in (d_7_v6_) else ((d_6_v5_).f7) + (d_3_v2_))
                    d_9_v8_: C0
                    nw1_ = C0()
                    nw1_.ctor__(d_3_v2_, p1, (d_6_v5_).f7, d_0_v0_)
                    d_9_v8_ = nw1_
                    (globalState).f1 = p0
                    pass
            pass
        d_10_v9_: str
        d_10_v9_ = _dafny.CodePoint('p')
        d_11_v10_: _dafny.Seq
        d_11_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
        d_12_v11_: C0
        nw2_ = C0()
        nw2_.ctor__(_dafny.SeqWithoutIsStrInference([d_10_v9_]), len(d_11_v10_), d_11_v10_, d_0_v0_)
        d_12_v11_ = nw2_
        d_13_v12_: _dafny.Map
        d_13_v12_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_14_i1_ in range(default__.abs(795))]): (d_12_v11_).f8})
        d_15_v13_: D6
        d_15_v13_ = D6_DC15(p0, d_12_v11_, d_13_v12_, d_0_v0_)
        if (d_15_v13_).cf33:
            d_16_v14_: _dafny.MultiSet
            d_16_v14_ = _dafny.MultiSet([d_0_v0_, d_0_v0_, d_0_v0_])
            (globalState).f1 = ((d_16_v14_) - ((d_16_v14_) | (d_16_v14_))).cardinality
            (globalState).f1 = default__.safeDivisionInt(-28, p1)
            d_17_v15_: _dafny.Map
            d_17_v15_ = _dafny.Map({d_0_v0_: d_10_v9_})
            pat_let_tv0_ = d_13_v12_
            pat_let_tv1_ = d_0_v0_
            pat_let_tv2_ = d_12_v11_
            d_18_v16_: _dafny.Map
            def iife0_(_pat_let0_0):
                def iife1_(d_19_dt__update__tmp_h0_):
                    def iife2_(_pat_let1_0):
                        def iife3_(d_20_dt__update_hcf32_h0_):
                            def iife4_(_pat_let2_0):
                                def iife5_(d_21_dt__update_hcf33_h0_):
                                    def iife6_(_pat_let3_0):
                                        def iife7_(d_22_dt__update_hcf31_h0_):
                                            return D6_DC15((d_19_dt__update__tmp_h0_).cf30, d_22_dt__update_hcf31_h0_, d_20_dt__update_hcf32_h0_, d_21_dt__update_hcf33_h0_)
                                        return iife7_(_pat_let3_0)
                                    return iife6_(pat_let_tv2_)
                                return iife5_(_pat_let2_0)
                            return iife4_(pat_let_tv1_)
                        return iife3_(_pat_let1_0)
                    return iife2_(pat_let_tv0_)
                return iife1_(_pat_let0_0)
            d_18_v16_ = _dafny.Map({(iife0_(d_15_v13_)).cf33: p1})
            d_23_v17_: _dafny.MultiSet
            d_23_v17_ = _dafny.MultiSet([d_10_v9_, d_10_v9_, d_10_v9_])
            d_24_v18_: _dafny.Map
            d_24_v18_ = _dafny.Map({d_12_v11_: d_0_v0_})
            d_25_v19_: _dafny.Seq
            d_25_v19_ = _dafny.SeqWithoutIsStrInference([d_0_v0_])
            d_26_v20_: _dafny.Array
            nw3_ = _dafny.Array(None, 28)
            nw3_[int(0)] = ((d_12_v11_).f8) * ((d_12_v11_).f8)
            nw3_[int(1)] = p0
            nw3_[int(2)] = (p0) * ((0) - (p1))
            nw3_[int(3)] = (p0) + (p1)
            nw3_[int(4)] = -691
            nw3_[int(5)] = (d_12_v11_).f8
            nw3_[int(6)] = p1
            nw3_[int(7)] = (d_12_v11_).f8
            nw3_[int(8)] = len(d_17_v15_)
            nw3_[int(9)] = ((d_18_v16_)[d_0_v0_] if (d_0_v0_) in (d_18_v16_) else len(d_11_v10_))
            nw3_[int(10)] = (d_12_v11_).f8
            nw3_[int(11)] = (d_12_v11_).f8
            nw3_[int(12)] = p0
            nw3_[int(13)] = p1
            nw3_[int(14)] = p1
            nw3_[int(15)] = (d_12_v11_).f8
            nw3_[int(16)] = p1
            nw3_[int(17)] = ((d_12_v11_).f8) * ((d_15_v13_).cf30)
            nw3_[int(18)] = p0
            nw3_[int(19)] = (d_12_v11_).f8
            nw3_[int(20)] = (d_12_v11_).f8
            nw3_[int(21)] = ((d_23_v17_)[d_10_v9_] if (d_10_v9_) in (d_23_v17_) else (d_12_v11_).f8)
            nw3_[int(22)] = (d_12_v11_).f8
            nw3_[int(23)] = (d_12_v11_).f8
            nw3_[int(24)] = p0
            nw3_[int(25)] = (0) - (default__.safeModuloInt(p1, len(d_24_v18_)))
            nw3_[int(26)] = len(d_25_v19_)
            nw3_[int(27)] = (d_12_v11_).f8
            d_26_v20_ = nw3_
            index0_ = default__.safeIndex(572, (d_26_v20_).length(0))
            (d_26_v20_)[index0_] = default__.safeDivisionInt(((d_16_v14_)[d_0_v0_] if (d_0_v0_) in (d_16_v14_) else (d_12_v11_).f8), p1)
            index1_ = default__.safeIndex(572, (d_26_v20_).length(0))
            (d_26_v20_)[index1_] = (d_12_v11_).f8
            d_25_v19_ = d_25_v19_
            (globalState).f0 = default__.fm1(740, d_0_v0_, globalState)
        elif True:
            d_27_v21_: bool
            d_27_v21_ = d_0_v0_
            d_28_v22_: _dafny.Array
            nw4_ = _dafny.Array(int(0), 26)
            d_28_v22_ = nw4_
            d_29_v23_: _dafny.Map
            d_29_v23_ = _dafny.Map({d_27_v21_: d_28_v22_})
            d_29_v23_ = ((d_29_v23_) | (d_29_v23_)) | (d_29_v23_)
            d_30_v24_: C0
            nw5_ = C0()
            nw5_.ctor__(d_11_v10_, (d_12_v11_).f8, (d_12_v11_).f7, (d_0_v0_) and (d_0_v0_))
            d_30_v24_ = nw5_
            d_31_v25_: _dafny.Seq
            d_31_v25_ = _dafny.SeqWithoutIsStrInference([True])
            d_32_v26_: C0
            nw6_ = C0()
            nw6_.ctor__(((d_30_v24_).f7) + (_dafny.SeqWithoutIsStrInference([d_10_v9_, default__.fm6(globalState), d_10_v9_])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qf"))), (_dafny.SeqWithoutIsStrInference([d_10_v9_ for d_33_i2_ in range(default__.abs(793))])).set(default__.safeIndex(p1, len(_dafny.SeqWithoutIsStrInference([d_10_v9_ for d_34_i2_ in range(default__.abs(793))]))), d_10_v9_), not ((d_31_v25_)[default__.safeIndex((d_30_v24_).f8, len(d_31_v25_))]) or (d_0_v0_))
            d_32_v26_ = nw6_
            d_35_v27_: _dafny.Map
            d_35_v27_ = _dafny.Map({d_10_v9_: (d_30_v24_).f8})
            d_36_v28_: _dafny.Set
            d_36_v28_ = _dafny.Set({d_10_v9_})
            d_37_v29_: _dafny.Map
            d_37_v29_ = _dafny.Map({(d_35_v27_) != (d_35_v27_): (d_36_v28_ if d_0_v0_ else default__.fm8(globalState))})
            d_37_v29_ = (d_37_v29_).set(d_0_v0_, d_36_v28_)
            index2_ = default__.safeIndex(142, (d_28_v22_).length(0))
            (d_28_v22_)[index2_] = (d_32_v26_).f8
            d_38_v30_: _dafny.Map
            d_38_v30_ = _dafny.Map({d_28_v22_: d_0_v0_})
            d_39_v31_: _dafny.Map
            d_39_v31_ = _dafny.Map({len(d_38_v30_): 754})
            d_40_v32_: _dafny.Map
            d_40_v32_ = _dafny.Map({d_0_v0_: d_0_v0_})
            d_41_v33_: _dafny.Seq
            d_41_v33_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_39_v31_)), len(d_40_v32_)])
            d_42_v34_: _dafny.Map
            d_42_v34_ = _dafny.Map({not(d_0_v0_): (len(_dafny.SeqWithoutIsStrInference([537 for d_43_i3_ in range(default__.abs(415))]))) + (len(d_41_v33_))})
            d_44_v35_: _dafny.Set
            d_44_v35_ = _dafny.Set({(0) - ((0) - ((d_32_v26_).f8))})
            index3_ = default__.safeIndex(142, (d_28_v22_).length(0))
            (d_28_v22_)[index3_] = ((d_42_v34_)[(d_31_v25_) != (d_31_v25_)] if ((d_31_v25_) != (d_31_v25_)) in (d_42_v34_) else len((_dafny.Set({p0})) | (d_44_v35_)))
        hi0_ = p1
        for d_45_i4_ in range((d_12_v11_).f8, hi0_):
            d_46_v36_: _dafny.Seq
            d_46_v36_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            if default__.fm1((p1 if d_0_v0_ else (d_46_v36_)[default__.safeIndex(p1, len(d_46_v36_))]), (d_11_v10_) < (_dafny.SeqWithoutIsStrInference([d_10_v9_ for d_47_i5_ in range(default__.abs(324))])), globalState):
                d_11_v10_ = (d_12_v11_).f7
                (globalState).f0 = not((d_0_v0_) == (d_0_v0_))
                (globalState).f1 = ((d_12_v11_).f8 if (d_0_v0_ if False else d_0_v0_) else (0) - (p1))
                d_48_v37_: T1
                nw7_ = C0()
                nw7_.ctor__(d_11_v10_, (d_12_v11_).f8, (d_12_v11_).f7, True)
                d_48_v37_ = nw7_
                d_49_v38_: _dafny.Set
                d_49_v38_ = _dafny.Set({d_48_v37_, d_48_v37_})
                d_50_v39_: _dafny.MultiSet
                d_50_v39_ = _dafny.MultiSet([d_0_v0_, d_0_v0_])
                d_51_v40_: _dafny.Map
                d_51_v40_ = _dafny.Map({d_49_v38_: d_50_v39_})
                rhs0_ = len((d_12_v11_).f7)
                rhs1_ = default__.fm6(globalState)
                rhs2_ = default__.safeDivisionInt(len((d_51_v40_) | (d_51_v40_)), (0) - ((d_12_v11_).f8))
                lhs0_ = globalState
                lhs1_ = globalState
                lhs0_.f1 = rhs0_
                d_10_v9_ = rhs1_
                lhs1_.f1 = rhs2_
                d_52_v41_: _dafny.Map
                d_52_v41_ = _dafny.Map({d_0_v0_: (d_12_v11_).f8})
                d_52_v41_ = (d_52_v41_).set(d_0_v0_, default__.safeModuloInt(p0, p0))
            elif True:
                d_10_v9_ = _dafny.CodePoint('i')
                d_53_v42_: _dafny.Seq
                d_53_v42_ = _dafny.SeqWithoutIsStrInference([d_12_v11_, d_12_v11_, d_12_v11_, d_12_v11_, d_12_v11_])
                d_53_v42_ = (d_53_v42_) + (d_53_v42_)
                d_54_v43_: _dafny.Seq
                d_54_v43_ = _dafny.SeqWithoutIsStrInference([d_0_v0_, False, d_0_v0_])
                d_55_v44_: _dafny.Array
                def lambda0_(d_56_v0_):
                    def lambda1_(d_57_i6_):
                        return d_56_v0_

                    return lambda1_

                init0_ = lambda0_(d_0_v0_)
                nw8_ = _dafny.Array(None, 24)
                for i0_0_ in range(nw8_.length(0)):
                    nw8_[i0_0_] = init0_(i0_0_)
                d_55_v44_ = nw8_
                d_58_v45_: _dafny.Seq
                d_58_v45_ = _dafny.SeqWithoutIsStrInference([d_55_v44_, d_55_v44_])
                d_59_v46_: _dafny.Map
                d_59_v46_ = _dafny.Map({len(d_58_v45_): (d_10_v9_) in ((d_12_v11_).f7)})
                rhs3_ = (d_12_v11_).fm4((d_12_v11_).f8, 843, (d_54_v43_) + (default__.fm9(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klg")), p0, globalState)), globalState)
                rhs4_ = len(d_59_v46_)
                lhs2_ = globalState
                lhs3_ = globalState
                lhs2_.f1 = rhs3_
                lhs3_.f1 = rhs4_
                d_12_v11_ = d_12_v11_
                d_12_v11_ = (d_15_v13_).cf31
            d_13_v12_ = (d_13_v12_).set(d_11_v10_, (d_12_v11_).f8)
            d_60_v47_: T1
            nw9_ = C0()
            nw9_.ctor__((d_12_v11_).f7, (d_12_v11_).f8, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cwoxfwq")), d_0_v0_)
            d_60_v47_ = nw9_
            d_61_v48_: bool
            d_61_v48_ = d_0_v0_
            d_62_v49_: _dafny.Map
            d_62_v49_ = _dafny.Map({d_0_v0_: d_0_v0_})
            d_63_v50_: _dafny.Array
            nw10_ = _dafny.Array(None, 29)
            nw10_[int(0)] = d_0_v0_
            nw10_[int(1)] = True
            nw10_[int(2)] = False
            nw10_[int(3)] = d_0_v0_
            nw10_[int(4)] = d_0_v0_
            nw10_[int(5)] = d_0_v0_
            nw10_[int(6)] = d_0_v0_
            nw10_[int(7)] = (d_61_v48_)
            nw10_[int(8)] = d_0_v0_
            nw10_[int(9)] = True
            nw10_[int(10)] = d_60_v47_.f6
            nw10_[int(11)] = d_0_v0_
            nw10_[int(12)] = True
            nw10_[int(13)] = d_60_v47_.f6
            nw10_[int(14)] = d_0_v0_
            nw10_[int(15)] = d_0_v0_
            nw10_[int(16)] = d_60_v47_.f6
            nw10_[int(17)] = d_60_v47_.f6
            nw10_[int(18)] = d_60_v47_.f6
            nw10_[int(19)] = not(d_0_v0_)
            nw10_[int(20)] = d_60_v47_.f6
            nw10_[int(21)] = d_60_v47_.f6
            nw10_[int(22)] = d_0_v0_
            nw10_[int(23)] = d_60_v47_.f6
            nw10_[int(24)] = d_60_v47_.f6
            nw10_[int(25)] = False
            nw10_[int(26)] = d_60_v47_.f6
            nw10_[int(27)] = ((d_62_v49_)[d_60_v47_.f6] if (d_60_v47_.f6) in (d_62_v49_) else d_0_v0_)
            nw10_[int(28)] = d_60_v47_.f6
            d_63_v50_ = nw10_
            d_64_v51_: _dafny.Map
            d_64_v51_ = _dafny.Map({d_60_v47_: d_63_v50_})
            d_64_v51_ = (d_64_v51_).set(d_60_v47_, d_63_v50_)
            if (d_45_i4_) > (p0):
                d_65_v52_: _dafny.Array
                nw11_ = _dafny.Array(int(0), 26)
                d_65_v52_ = nw11_
                index4_ = default__.safeIndex(507, (d_65_v52_).length(0))
                (d_65_v52_)[index4_] = (d_46_v36_)[default__.safeIndex((d_12_v11_).f8, len(d_46_v36_))]
                index5_ = default__.safeIndex(507, (d_65_v52_).length(0))
                (d_65_v52_)[index5_] = (d_12_v11_).f8
                index6_ = default__.safeIndex(507, (d_65_v52_).length(0))
                (d_65_v52_)[index6_] = 749
                d_66_v53_: C0
                nw12_ = C0()
                nw12_.ctor__((d_12_v11_).f7, 617, (d_60_v47_).f5, d_60_v47_.f6)
                d_66_v53_ = nw12_
                index7_ = default__.safeIndex(178, (d_63_v50_).length(0))
                (d_63_v50_)[index7_] = d_60_v47_.f6
                d_67_v54_: D6
                d_67_v54_ = D6_DC13(d_60_v47_.f6, not(d_0_v0_))
                d_68_v55_: _dafny.Array
                def lambda2_(d_69_p0_):
                    def lambda3_(d_70_i7_):
                        return _dafny.Set({d_69_p0_})

                    return lambda3_

                init1_ = lambda2_(p0)
                nw13_ = _dafny.Array(None, 19)
                for i0_1_ in range(nw13_.length(0)):
                    nw13_[i0_1_] = init1_(i0_1_)
                d_68_v55_ = nw13_
                d_71_v56_: _dafny.Set
                d_71_v56_ = _dafny.Set({(d_12_v11_).f8, (d_65_v52_)[default__.safeIndex(507, (d_65_v52_).length(0))], p1, (d_66_v53_).f8})
                index8_ = default__.safeIndex(310, (d_68_v55_).length(0))
                (d_68_v55_)[index8_] = (d_71_v56_) | (d_71_v56_)
                d_72_v57_: _dafny.Array
                nw14_ = _dafny.Array(None, 2)
                nw14_[int(0)] = (d_12_v11_ if d_60_v47_.f6 else d_12_v11_)
                nw14_[int(1)] = d_12_v11_
                d_72_v57_ = nw14_
                d_73_v58_: _dafny.Map
                d_73_v58_ = _dafny.Map({d_12_v11_: d_62_v49_})
                d_74_v59_: _dafny.Seq
                d_74_v59_ = _dafny.SeqWithoutIsStrInference([d_71_v56_])
                pat_let_tv3_ = d_0_v0_
                index9_ = default__.safeIndex(178, (d_63_v50_).length(0))
                index10_ = default__.safeIndex(310, (d_68_v55_).length(0))
                def iife8_(_pat_let4_0):
                    def iife9_(d_75_dt__update__tmp_h1_):
                        def iife10_(_pat_let5_0):
                            def iife11_(d_76_dt__update_hcf26_h0_):
                                return D6_DC13(d_76_dt__update_hcf26_h0_, (d_75_dt__update__tmp_h1_).cf27)
                            return iife11_(_pat_let5_0)
                        return iife10_(pat_let_tv3_)
                    return iife9_(_pat_let4_0)
                rhs5_ = (_dafny.Map({d_0_v0_: d_60_v47_.f6})) | ((((d_73_v58_)[d_12_v11_] if (d_12_v11_) in (d_73_v58_) else d_62_v49_)).set(d_0_v0_, d_0_v0_))
                rhs6_ = d_60_v47_.f6
                rhs7_ = iife8_(d_67_v54_)
                rhs8_ = (d_74_v59_)[default__.safeIndex((d_66_v53_).f8, len(d_74_v59_))]
                rhs9_ = d_72_v57_
                lhs4_ = d_63_v50_
                lhs5_ = default__.safeIndex(178, (d_63_v50_).length(0))
                lhs6_ = d_68_v55_
                lhs7_ = default__.safeIndex(310, (d_68_v55_).length(0))
                d_62_v49_ = rhs5_
                lhs4_[lhs5_] = rhs6_
                d_67_v54_ = rhs7_
                lhs6_[lhs7_] = rhs8_
                d_72_v57_ = rhs9_
                (globalState).f1 = ((0) - ((0) - (p0))) + ((d_12_v11_).f8)
            elif True:
                d_77_v60_: _dafny.Array
                nw15_ = _dafny.Array(_dafny.Seq({}), 2)
                d_77_v60_ = nw15_
                d_78_v61_: T0
                nw16_ = C0()
                nw16_.ctor__(_dafny.SeqWithoutIsStrInference([default__.fm6(globalState)]), (0) - ((d_12_v11_).f8), _dafny.SeqWithoutIsStrInference([d_10_v9_ for d_79_i8_ in range(default__.abs(264))]), d_60_v47_.f6)
                d_78_v61_ = nw16_
                d_80_v62_: _dafny.Seq
                d_80_v62_ = _dafny.SeqWithoutIsStrInference([d_78_v61_])
                d_81_v63_: _dafny.Seq
                d_81_v63_ = _dafny.SeqWithoutIsStrInference([d_80_v62_, d_80_v62_, d_80_v62_, d_80_v62_, d_80_v62_])
                index11_ = default__.safeIndex(46, (d_77_v60_).length(0))
                (d_77_v60_)[index11_] = (d_81_v63_)[default__.safeIndex(d_45_i4_, len(d_81_v63_))]
                index12_ = default__.safeIndex(46, (d_77_v60_).length(0))
                (d_77_v60_)[index12_] = d_80_v62_
                index13_ = default__.safeIndex(334, (d_63_v50_).length(0))
                (d_63_v50_)[index13_] = d_60_v47_.f6
                index14_ = default__.safeIndex(334, (d_63_v50_).length(0))
                (d_63_v50_)[index14_] = d_60_v47_.f6
                d_82_v64_: _dafny.Map
                d_82_v64_ = _dafny.Map({d_60_v47_: (d_12_v11_).f8})
                d_82_v64_ = (d_82_v64_).set(d_60_v47_, len((d_46_v36_) + (d_46_v36_)))
                (d_60_v47_).f6 = (d_63_v50_)[default__.safeIndex(334, (d_63_v50_).length(0))]
                d_83_v65_: C0
                nw17_ = C0()
                nw17_.ctor__((d_12_v11_).f7, (d_12_v11_).f8, (d_78_v61_).f5, not (d_60_v47_.f6) or (True))
                d_83_v65_ = nw17_
        d_11_v10_ = (d_11_v10_) + ((d_12_v11_).f7)
        d_84_v66_: _dafny.Set
        d_84_v66_ = _dafny.Set({d_0_v0_})
        d_85_v67_: _dafny.Seq
        d_85_v67_ = _dafny.SeqWithoutIsStrInference([d_0_v0_, not(d_0_v0_)])
        hi1_ = (_dafny.MultiSet(d_85_v67_)).cardinality
        for d_86_i9_ in range(len(d_84_v66_), hi1_):
            d_87_v68_: _dafny.Array
            nw18_ = _dafny.Array(_dafny.Array(None, 0), 20)
            d_87_v68_ = nw18_
            d_88_v69_: _dafny.Set
            d_88_v69_ = _dafny.Set({p1})
            d_89_v70_: _dafny.Map
            d_89_v70_ = _dafny.Map({_dafny.CodePoint('r'): False})
            d_90_v71_: _dafny.MultiSet
            d_90_v71_ = _dafny.MultiSet([p0])
            d_91_v72_: _dafny.Seq
            d_91_v72_ = _dafny.SeqWithoutIsStrInference([(d_90_v71_).cardinality, p1])
            d_92_v73_: _dafny.Array
            nw19_ = _dafny.Array(None, 11)
            nw19_[int(0)] = len(d_88_v69_)
            nw19_[int(1)] = d_86_i9_
            nw19_[int(2)] = (0) - (len(d_89_v70_))
            nw19_[int(3)] = -43
            nw19_[int(4)] = len(d_91_v72_)
            nw19_[int(5)] = d_86_i9_
            nw19_[int(6)] = (d_90_v71_).cardinality
            nw19_[int(7)] = d_86_i9_
            nw19_[int(8)] = (d_12_v11_).f8
            nw19_[int(9)] = -71
            nw19_[int(10)] = (d_91_v72_)[default__.safeIndex((d_12_v11_).f8, len(d_91_v72_))]
            d_92_v73_ = nw19_
            index15_ = default__.safeIndex(807, (d_87_v68_).length(0))
            (d_87_v68_)[index15_] = d_92_v73_
            index16_ = default__.safeIndex(807, (d_87_v68_).length(0))
            (d_87_v68_)[index16_] = ((d_92_v73_) if d_0_v0_ else d_92_v73_)
            d_93_v74_: C0
            nw20_ = C0()
            nw20_.ctor__((d_12_v11_).f7, p0, (d_12_v11_).f7, d_0_v0_)
            d_93_v74_ = nw20_
            d_94_v75_: _dafny.Array
            nw21_ = _dafny.Array(None, 4)
            d_94_v75_ = nw21_
            index17_ = default__.safeIndex(343, (d_94_v75_).length(0))
            (d_94_v75_)[index17_] = d_12_v11_
            d_95_v76_: _dafny.Array
            nw22_ = _dafny.Array(False, 14)
            d_95_v76_ = nw22_
            index18_ = default__.safeIndex(771, (d_95_v76_).length(0))
            (d_95_v76_)[index18_] = (d_0_v0_) and (d_0_v0_)
            index19_ = default__.safeIndex(858, (d_95_v76_).length(0))
            (d_95_v76_)[index19_] = (not(d_0_v0_)) not in (d_84_v66_)
            index20_ = default__.safeIndex(288, (d_92_v73_).length(0))
            (d_92_v73_)[index20_] = (d_12_v11_).f8
            d_96_v77_: _dafny.Seq
            d_96_v77_ = _dafny.SeqWithoutIsStrInference([d_93_v74_, d_12_v11_, d_12_v11_])
            index21_ = default__.safeIndex(343, (d_94_v75_).length(0))
            index22_ = default__.safeIndex(771, (d_95_v76_).length(0))
            index23_ = default__.safeIndex(858, (d_95_v76_).length(0))
            index24_ = default__.safeIndex(288, (d_92_v73_).length(0))
            rhs10_ = (d_96_v77_)[default__.safeIndex((d_12_v11_).f8, len(d_96_v77_))]
            rhs11_ = not(not(d_0_v0_))
            rhs12_ = (d_93_v74_).f7
            rhs13_ = d_0_v0_
            rhs14_ = (d_12_v11_).f8
            lhs8_ = d_94_v75_
            lhs9_ = default__.safeIndex(343, (d_94_v75_).length(0))
            lhs10_ = d_95_v76_
            lhs11_ = default__.safeIndex(771, (d_95_v76_).length(0))
            lhs12_ = d_95_v76_
            lhs13_ = default__.safeIndex(858, (d_95_v76_).length(0))
            lhs14_ = d_92_v73_
            lhs15_ = default__.safeIndex(288, (d_92_v73_).length(0))
            lhs8_[lhs9_] = rhs10_
            lhs10_[lhs11_] = rhs11_
            d_11_v10_ = rhs12_
            lhs12_[lhs13_] = rhs13_
            lhs14_[lhs15_] = rhs14_
            d_97_v78_: _dafny.Map
            d_97_v78_ = _dafny.Map({(p0 if (d_95_v76_)[default__.safeIndex(771, (d_95_v76_).length(0))] else (d_12_v11_).f8): (d_93_v74_).f8})
            d_98_v79_: D6
            d_98_v79_ = D6_DC13(d_0_v0_, d_0_v0_)
            d_99_v80_: D6
            d_99_v80_ = D6_DC16(d_98_v79_)
            d_100_v81_: _dafny.Map
            d_100_v81_ = _dafny.Map({d_99_v80_: True})
            d_101_v82_: _dafny.Map
            d_101_v82_ = _dafny.Map({d_0_v0_: d_100_v81_})
            index25_ = default__.safeIndex(288, (d_92_v73_).length(0))
            (d_92_v73_)[index25_] = ((d_97_v78_)[(d_93_v74_).f8] if ((d_93_v74_).f8) in (d_97_v78_) else (0) - (len(d_101_v82_)))
        d_13_v12_ = (d_13_v12_).set(d_11_v10_, ((d_12_v11_).f8) * (p1))
        r0 = d_0_v0_
        d_102_v83_: _dafny.Set
        d_102_v83_ = _dafny.Set({(d_12_v11_).f8})
        r1 = d_102_v83_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_103_globalState_: GlobalState
        nw23_ = GlobalState()
        nw23_.ctor__(False, 475, _dafny.CodePoint('r'), 449, False)
        d_103_globalState_ = nw23_
        d_104_v0_: int
        d_104_v0_ = -316
        (d_103_globalState_).f1 = d_104_v0_
        d_105_v1_: _dafny.Array
        def lambda4_(d_106_i0_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eas"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opxmqoxt")))

        init2_ = lambda4_
        nw24_ = _dafny.Array(None, 18)
        for i0_2_ in range(nw24_.length(0)):
            nw24_[i0_2_] = init2_(i0_2_)
        d_105_v1_ = nw24_
        d_107_v2_: _dafny.Seq
        d_107_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eedslel"))
        d_108_v3_: _dafny.Map
        d_108_v3_ = _dafny.Map({True: d_107_v2_})
        index26_ = default__.safeIndex(662, (d_105_v1_).length(0))
        (d_105_v1_)[index26_] = (((d_108_v3_)[False] if (False) in (d_108_v3_) else d_107_v2_)) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_109_i1_ in range(default__.abs(-891))]))
        index27_ = default__.safeIndex(662, (d_105_v1_).length(0))
        (d_105_v1_)[index27_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osndtkw"))
        d_110_v4_: bool
        d_111_v5_: _dafny.Set
        out0_: bool
        out1_: _dafny.Set
        out0_, out1_ = default__.m0(d_104_v0_, (0) - (d_104_v0_), d_103_globalState_)
        d_110_v4_ = out0_
        d_111_v5_ = out1_
        d_112_v6_: _dafny.Array
        nw25_ = _dafny.Array(False, 18)
        d_112_v6_ = nw25_
        d_113_v7_: _dafny.Set
        d_113_v7_ = _dafny.Set({d_112_v6_, d_112_v6_})
        rhs15_ = (_dafny.Set({d_112_v6_})) - (d_113_v7_)
        rhs16_ = d_104_v0_
        rhs17_ = d_110_v4_
        lhs16_ = d_103_globalState_
        d_113_v7_ = rhs15_
        lhs16_.f1 = rhs16_
        d_110_v4_ = rhs17_
        if ((d_113_v7_) - (d_113_v7_)).ispropersubset(_dafny.Set({d_112_v6_})):
            d_114_v8_: _dafny.Seq
            d_114_v8_ = _dafny.SeqWithoutIsStrInference([d_110_v4_])
            d_115_v9_: _dafny.MultiSet
            d_115_v9_ = _dafny.MultiSet([d_104_v0_, d_104_v0_])
            d_116_v10_: _dafny.MultiSet
            d_116_v10_ = _dafny.MultiSet([d_110_v4_])
            d_117_v11_: str
            d_117_v11_ = _dafny.CodePoint('p')
            rhs18_ = (d_110_v4_) in (d_114_v8_)
            rhs19_ = d_110_v4_
            rhs20_ = default__.fm0((_dafny.MultiSet([(d_116_v10_).cardinality])).ispropersubset(d_115_v9_), d_107_v2_, d_110_v4_, d_103_globalState_)
            rhs21_ = not (d_110_v4_) or ((d_117_v11_) in ((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]))
            lhs17_ = d_103_globalState_
            lhs18_ = d_103_globalState_
            lhs19_ = d_103_globalState_
            lhs17_.f0 = rhs18_
            d_110_v4_ = rhs19_
            lhs18_.f1 = rhs20_
            lhs19_.f0 = rhs21_
            if not(d_110_v4_):
                (d_103_globalState_).f1 = len(d_114_v8_)
                d_118_v12_: _dafny.Map
                d_118_v12_ = _dafny.Map({default__.fm1((0) - (d_104_v0_), d_110_v4_, d_103_globalState_): d_110_v4_})
                d_119_v13_: _dafny.Set
                d_119_v13_ = _dafny.Set({d_110_v4_})
                rhs22_ = ((d_118_v12_)[d_110_v4_] if (d_110_v4_) in (d_118_v12_) else (_dafny.Set({d_110_v4_, not(d_110_v4_)})).ispropersubset(d_119_v13_))
                rhs23_ = default__.fm1(default__.safeModuloInt(len(d_107_v2_), d_104_v0_), d_110_v4_, d_103_globalState_)
                d_110_v4_ = rhs22_
                d_110_v4_ = rhs23_
                d_120_v14_: bool
                d_121_v15_: _dafny.Set
                out2_: bool
                out3_: _dafny.Set
                out2_, out3_ = default__.m0(len((d_111_v5_) - (d_111_v5_)), 416, d_103_globalState_)
                d_120_v14_ = out2_
                d_121_v15_ = out3_
                (d_103_globalState_).f0 = d_120_v14_
                d_122_v16_: _dafny.Array
                nw26_ = _dafny.Array(_dafny.Seq({}), 15)
                d_122_v16_ = nw26_
                index28_ = default__.safeIndex(185, (d_122_v16_).length(0))
                (d_122_v16_)[index28_] = d_114_v8_
                index29_ = default__.safeIndex(185, (d_122_v16_).length(0))
                (d_122_v16_)[index29_] = _dafny.SeqWithoutIsStrInference([d_110_v4_])
            elif True:
                d_117_v11_ = _dafny.CodePoint('h')
                d_123_v17_: bool
                d_123_v17_ = d_110_v4_
                d_124_v18_: _dafny.Map
                d_124_v18_ = _dafny.Map({(d_123_v17_): (0) - (d_104_v0_)})
                d_124_v18_ = (d_124_v18_).set(d_110_v4_, (d_104_v0_) + (d_104_v0_))
                (d_103_globalState_).f0 = d_110_v4_
                d_104_v0_ = d_104_v0_
                d_125_v19_: _dafny.MultiSet
                d_125_v19_ = _dafny.MultiSet([d_115_v9_])
                (d_103_globalState_).f0 = (_dafny.MultiSet([765])) not in (d_125_v19_)
            d_126_v21_: D1
            d_126_v21_ = D1_DC3(False, len((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]), d_110_v4_, d_117_v11_)
            def iife12_():
                coll0_ = _dafny.Map()
                compr_0_: int
                for compr_0_ in (d_115_v9_).Elements:
                    d_127_v20_: int = compr_0_
                    if (d_127_v20_) in (d_115_v9_):
                        coll0_[default__.safeDivisionInt(d_127_v20_, len((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]))] = d_110_v4_
                return _dafny.Map(coll0_)
            rhs24_ = (D1_DC3(d_110_v4_, len(iife12_()
), (d_126_v21_).cf9, d_117_v11_)).cf8
            rhs25_ = d_110_v4_
            lhs20_ = d_103_globalState_
            d_104_v0_ = rhs24_
            lhs20_.f0 = rhs25_
            index30_ = default__.safeIndex(662, (d_105_v1_).length(0))
            (d_105_v1_)[index30_] = d_107_v2_
            d_128_v22_: bool
            d_129_v23_: _dafny.Set
            out4_: bool
            out5_: _dafny.Set
            out4_, out5_ = default__.m0((0) - (d_104_v0_), d_104_v0_, d_103_globalState_)
            d_128_v22_ = out4_
            d_129_v23_ = out5_
        elif True:
            d_130_v24_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Map({}), 10)
            d_130_v24_ = nw27_
            d_131_v25_: _dafny.Map
            d_131_v25_ = _dafny.Map({d_104_v0_: True})
            index31_ = default__.safeIndex(351, (d_130_v24_).length(0))
            (d_130_v24_)[index31_] = d_131_v25_
            index32_ = default__.safeIndex(351, (d_130_v24_).length(0))
            (d_130_v24_)[index32_] = (d_131_v25_ if default__.fm1(d_104_v0_, d_110_v4_, d_103_globalState_) else _dafny.Map({d_104_v0_: d_110_v4_}))
            d_132_v26_: _dafny.Seq
            d_132_v26_ = _dafny.SeqWithoutIsStrInference([d_104_v0_, (0) - (len((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]))])
            d_133_v27_: _dafny.MultiSet
            d_133_v27_ = _dafny.MultiSet([d_110_v4_])
            if ((d_104_v0_ if default__.fm1((0) - (d_104_v0_), d_110_v4_, d_103_globalState_) else len(d_132_v26_))) > ((0) - ((((d_133_v27_)[d_110_v4_] if (d_110_v4_) in (d_133_v27_) else d_104_v0_)) - (d_104_v0_))):
                d_134_v28_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.Seq({}), 7)
                d_134_v28_ = nw28_
                d_135_v29_: _dafny.Seq
                d_135_v29_ = _dafny.SeqWithoutIsStrInference([d_110_v4_, d_110_v4_])
                index33_ = default__.safeIndex(977, (d_134_v28_).length(0))
                (d_134_v28_)[index33_] = (d_135_v29_) + (d_135_v29_)
                index34_ = default__.safeIndex(977, (d_134_v28_).length(0))
                (d_134_v28_)[index34_] = d_135_v29_
                (d_103_globalState_).f1 = len((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))])
                d_136_v30_: D1
                d_136_v30_ = D1_DC1(d_104_v0_)
                d_136_v30_ = d_136_v30_
                d_137_v31_: bool
                d_138_v32_: _dafny.Set
                out6_: bool
                out7_: _dafny.Set
                out6_, out7_ = default__.m0(d_104_v0_, default__.safeModuloInt(d_104_v0_, d_104_v0_), d_103_globalState_)
                d_137_v31_ = out6_
                d_138_v32_ = out7_
                nw29_ = _dafny.Array(False, 8)
                d_112_v6_ = nw29_
            elif True:
                d_139_v33_: str
                d_139_v33_ = _dafny.CodePoint('s')
                d_139_v33_ = d_139_v33_
                d_140_v34_: D1
                d_140_v34_ = D1_DC2((_dafny.SeqWithoutIsStrInference([d_104_v0_])) + (d_132_v26_), d_104_v0_, d_104_v0_, (d_107_v2_ if d_110_v4_ else default__.fm2(True, d_110_v4_, d_104_v0_, d_110_v4_, d_103_globalState_)), d_111_v5_)
                d_140_v34_ = d_140_v34_
                index35_ = default__.safeIndex(959, (d_112_v6_).length(0))
                (d_112_v6_)[index35_] = d_110_v4_
                index36_ = default__.safeIndex(959, (d_112_v6_).length(0))
                (d_112_v6_)[index36_] = not ((d_110_v4_ if d_110_v4_ else not(d_110_v4_))) or (d_110_v4_)
                d_141_v35_: _dafny.Seq
                d_141_v35_ = _dafny.SeqWithoutIsStrInference([d_110_v4_, d_110_v4_])
                d_142_v36_: bool
                d_143_v37_: _dafny.Set
                out8_: bool
                out9_: _dafny.Set
                out8_, out9_ = default__.m0(len(d_141_v35_), (len(d_107_v2_)) - (d_104_v0_), d_103_globalState_)
                d_142_v36_ = out8_
                d_143_v37_ = out9_
                index37_ = default__.safeIndex(959, (d_112_v6_).length(0))
                rhs26_ = (d_104_v0_) != (d_104_v0_)
                rhs27_ = d_110_v4_
                lhs21_ = d_112_v6_
                lhs22_ = default__.safeIndex(959, (d_112_v6_).length(0))
                d_110_v4_ = rhs26_
                lhs21_[lhs22_] = rhs27_
            index38_ = default__.safeIndex(662, (d_105_v1_).length(0))
            (d_105_v1_)[index38_] = (d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]
            d_144_v38_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.Map({}), 6)
            d_144_v38_ = nw30_
            d_145_v39_: _dafny.Map
            d_145_v39_ = _dafny.Map({d_110_v4_: d_144_v38_})
            d_145_v39_ = (d_145_v39_).set(False, d_144_v38_)
            d_146_v40_: _dafny.Seq
            d_146_v40_ = _dafny.SeqWithoutIsStrInference([default__.fm1((d_132_v26_)[default__.safeIndex(d_104_v0_, len(d_132_v26_))], False, d_103_globalState_)])
            (d_103_globalState_).f0 = not((_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_110_v4_, True])) + (d_146_v40_))).isdisjoint(d_133_v27_))
        d_147_v41_: _dafny.Array
        def lambda5_(d_148_v0_):
            def lambda6_(d_149_i2_):
                return D1_DC4(D1_DC1(d_148_v0_))

            return lambda6_

        init3_ = lambda5_(d_104_v0_)
        nw31_ = _dafny.Array(None, 9)
        for i0_3_ in range(nw31_.length(0)):
            nw31_[i0_3_] = init3_(i0_3_)
        d_147_v41_ = nw31_
        d_150_v42_: _dafny.MultiSet
        d_150_v42_ = _dafny.MultiSet([d_104_v0_])
        d_151_v43_: _dafny.Map
        d_151_v43_ = _dafny.Map({d_110_v4_: d_150_v42_})
        d_152_v44_: _dafny.Array
        nw32_ = _dafny.Array(None, 14)
        nw32_[int(0)] = len(d_107_v2_)
        nw32_[int(1)] = d_104_v0_
        nw32_[int(2)] = d_104_v0_
        nw32_[int(3)] = len(d_151_v43_)
        nw32_[int(4)] = d_104_v0_
        nw32_[int(5)] = d_104_v0_
        nw32_[int(6)] = 99
        nw32_[int(7)] = d_104_v0_
        nw32_[int(8)] = d_104_v0_
        nw32_[int(9)] = 861
        nw32_[int(10)] = d_104_v0_
        nw32_[int(11)] = d_104_v0_
        nw32_[int(12)] = d_104_v0_
        nw32_[int(13)] = 623
        d_152_v44_ = nw32_
        d_153_v45_: _dafny.Seq
        d_153_v45_ = _dafny.SeqWithoutIsStrInference([d_152_v44_, d_152_v44_])
        d_154_v46_: _dafny.Map
        d_154_v46_ = _dafny.Map({(d_153_v45_)[default__.safeIndex(281, len(d_153_v45_))]: d_110_v4_})
        d_155_v47_: D1
        d_155_v47_ = D1_DC1(len(d_154_v46_))
        d_156_v48_: D1
        d_156_v48_ = D1_DC4(d_155_v47_)
        index39_ = default__.safeIndex(269, (d_147_v41_).length(0))
        (d_147_v41_)[index39_] = d_156_v48_
        d_157_v49_: _dafny.Array
        nw33_ = _dafny.Array(_dafny.Array(None, 0), 23)
        d_157_v49_ = nw33_
        d_158_v50_: _dafny.Map
        d_158_v50_ = _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_159_i3_ in range(default__.abs(524))]))})
        d_160_v51_: _dafny.Set
        d_160_v51_ = _dafny.Set({d_110_v4_})
        pat_let_tv4_ = d_155_v47_
        index40_ = default__.safeIndex(269, (d_147_v41_).length(0))
        def iife13_(_pat_let6_0):
            def iife14_(d_161_dt__update__tmp_h0_):
                def iife15_(_pat_let7_0):
                    def iife16_(d_162_dt__update_hcf11_h0_):
                        return D1_DC4(d_162_dt__update_hcf11_h0_)
                    return iife16_(_pat_let7_0)
                return iife15_(pat_let_tv4_)
            return iife14_(_pat_let6_0)
        rhs28_ = d_104_v0_
        rhs29_ = iife13_(d_156_v48_)
        rhs30_ = default__.fm1((((d_158_v50_)[d_110_v4_] if (d_110_v4_) in (d_158_v50_) else d_104_v0_)) * (d_104_v0_), (d_110_v4_ if d_110_v4_ else True), d_103_globalState_)
        rhs31_ = d_157_v49_
        rhs32_ = (0) - (((d_150_v42_)[len(d_160_v51_)] if (len(d_160_v51_)) in (d_150_v42_) else (0) - (d_104_v0_)))
        lhs23_ = d_147_v41_
        lhs24_ = default__.safeIndex(269, (d_147_v41_).length(0))
        lhs25_ = d_103_globalState_
        lhs26_ = d_103_globalState_
        d_104_v0_ = rhs28_
        lhs23_[lhs24_] = rhs29_
        lhs25_.f0 = rhs30_
        d_157_v49_ = rhs31_
        lhs26_.f1 = rhs32_
        d_104_v0_ = d_104_v0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_112_v6_).length(0)):
            d_163_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_163_i4_)) and ((d_163_i4_) < ((d_112_v6_).length(0)))):
                (d_112_v6_)[(d_163_i4_)] = d_110_v4_
        d_164_v52_: str
        d_164_v52_ = _dafny.CodePoint('f')
        d_165_v53_: D1
        d_165_v53_ = D1_DC3(False, d_104_v0_, True, d_164_v52_)
        d_166_v54_: _dafny.Set
        d_166_v54_ = _dafny.Set({d_165_v53_, d_165_v53_})
        d_166_v54_ = d_166_v54_
        d_167_v55_: _dafny.Seq
        d_167_v55_ = _dafny.SeqWithoutIsStrInference([d_104_v0_])
        d_107_v2_ = (d_107_v2_).set(default__.safeIndex((0) - (len(((d_158_v50_).set(d_110_v4_, len(d_167_v55_))) | (d_158_v50_))), len(d_107_v2_)), d_164_v52_)
        source0_ = d_165_v53_
        if source0_.is_DC2:
            d_168___mcc_h0_ = source0_.cf2
            d_169___mcc_h1_ = source0_.cf3
            d_170___mcc_h2_ = source0_.cf4
            d_171___mcc_h3_ = source0_.cf5
            d_172___mcc_h4_ = source0_.cf6
            d_173_cf6_ = d_172___mcc_h4_
            d_174_cf5_ = d_171___mcc_h3_
            d_175_cf4_ = d_170___mcc_h2_
            d_176_cf3_ = d_169___mcc_h1_
            d_177_cf2_ = d_168___mcc_h0_
            d_110_v4_ = not(False)
            d_178_v56_: _dafny.Map
            d_178_v56_ = _dafny.Map({595: (0) - (d_176_cf3_)})
            d_178_v56_ = (d_178_v56_).set(526, d_104_v0_)
            d_150_v42_ = ((_dafny.MultiSet([d_176_cf3_, d_104_v0_])).intersection(d_150_v42_)) - (d_150_v42_)
            d_179_v57_: bool
            d_179_v57_ = d_110_v4_
            source1_ = d_179_v57_
            d_180___mcc_h11_ = source1_
            d_181_cf0_ = d_180___mcc_h11_
            index41_ = default__.safeIndex(997, (d_112_v6_).length(0))
            (d_112_v6_)[index41_] = d_110_v4_
            d_182_v58_: _dafny.Map
            d_182_v58_ = _dafny.Map({d_110_v4_: d_181_cf0_})
            d_183_v59_: _dafny.Map
            d_183_v59_ = _dafny.Map({(d_165_v53_).cf8: d_164_v52_})
            d_184_v60_: _dafny.MultiSet
            d_184_v60_ = _dafny.MultiSet([((d_183_v59_)[d_176_cf3_] if (d_176_cf3_) in (d_183_v59_) else d_164_v52_)])
            d_185_v61_: _dafny.Map
            d_185_v61_ = _dafny.Map({d_175_cf4_: d_184_v60_})
            index42_ = default__.safeIndex(997, (d_112_v6_).length(0))
            (d_112_v6_)[index42_] = ((d_182_v58_)[(_dafny.MultiSet(d_107_v2_)).isdisjoint(((d_185_v61_)[d_104_v0_] if (d_104_v0_) in (d_185_v61_) else d_184_v60_))] if ((_dafny.MultiSet(d_107_v2_)).isdisjoint(((d_185_v61_)[d_104_v0_] if (d_104_v0_) in (d_185_v61_) else d_184_v60_))) in (d_182_v58_) else d_110_v4_)
            d_174_cf5_ = d_174_cf5_
            (d_103_globalState_).f0 = d_110_v4_
            d_160_v51_ = d_160_v51_
        elif source0_.is_DC3:
            d_186___mcc_h5_ = source0_.cf7
            d_187___mcc_h6_ = source0_.cf8
            d_188___mcc_h7_ = source0_.cf9
            d_189___mcc_h8_ = source0_.cf10
            d_190_cf10_ = d_189___mcc_h8_
            d_191_cf9_ = d_188___mcc_h7_
            d_192_cf8_ = d_187___mcc_h6_
            d_193_cf7_ = d_186___mcc_h5_
            d_111_v5_ = (d_111_v5_) - (d_111_v5_)
            d_194_v62_: bool
            d_194_v62_ = d_193_cf7_
            d_195_v63_: _dafny.Map
            d_195_v63_ = _dafny.Map({d_104_v0_: d_194_v62_})
            d_196_v64_: _dafny.Map
            d_196_v64_ = _dafny.Map({d_112_v6_: d_104_v0_})
            d_197_v65_: bool
            d_198_v66_: _dafny.Set
            out10_: bool
            out11_: _dafny.Set
            out10_, out11_ = default__.m0(default__.safeModuloInt(len(d_195_v63_), d_192_cf8_), len((d_196_v64_) | (d_196_v64_)), d_103_globalState_)
            d_197_v65_ = out10_
            d_198_v66_ = out11_
            if not(d_193_cf7_):
                d_152_v44_ = d_152_v44_
                d_150_v42_ = d_150_v42_
                (d_103_globalState_).f1 = len(d_107_v2_)
                index43_ = default__.safeIndex(653, (d_152_v44_).length(0))
                (d_152_v44_)[index43_] = ((0) - (d_192_cf8_) if d_193_cf7_ else -94)
                index44_ = default__.safeIndex(653, (d_152_v44_).length(0))
                (d_152_v44_)[index44_] = 532
                d_199_v67_: _dafny.Map
                d_199_v67_ = _dafny.Map({(default__.fm3((d_152_v44_)[default__.safeIndex(653, (d_152_v44_).length(0))], d_104_v0_, d_103_globalState_)): d_197_v65_})
                d_199_v67_ = d_199_v67_
            elif True:
                d_200_v68_: _dafny.Seq
                d_200_v68_ = _dafny.SeqWithoutIsStrInference([d_197_v65_])
                d_201_v69_: bool
                d_202_v70_: _dafny.Set
                out12_: bool
                out13_: _dafny.Set
                out12_, out13_ = default__.m0((0) - (len((d_200_v68_) + (d_200_v68_))), (0) - (d_104_v0_), d_103_globalState_)
                d_201_v69_ = out12_
                d_202_v70_ = out13_
                d_203_v71_: bool
                d_204_v72_: _dafny.Set
                out14_: bool
                out15_: _dafny.Set
                out14_, out15_ = default__.m0(320, (d_192_cf8_) * (d_192_cf8_), d_103_globalState_)
                d_203_v71_ = out14_
                d_204_v72_ = out15_
                d_205_v73_: _dafny.Array
                d_205_v73_ = d_152_v44_
                d_206_v74_: _dafny.MultiSet
                d_206_v74_ = _dafny.MultiSet([d_152_v44_, d_152_v44_, (d_205_v73_)])
                rhs33_ = d_193_cf7_
                rhs34_ = len((d_107_v2_) + (d_107_v2_))
                rhs35_ = (d_206_v74_).issubset((d_206_v74_) - (d_206_v74_))
                d_191_cf9_ = rhs33_
                d_192_cf8_ = rhs34_
                d_193_cf7_ = rhs35_
                d_207_v75_: C0
                nw34_ = C0()
                nw34_.ctor__((d_107_v2_).set(default__.safeIndex(d_192_cf8_, len(d_107_v2_)), d_164_v52_), (d_104_v0_) * (d_104_v0_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "homss")), d_193_cf7_)
                d_207_v75_ = nw34_
                d_208_v76_: bool
                d_209_v77_: _dafny.Set
                out16_: bool
                out17_: _dafny.Set
                out16_, out17_ = default__.m0(d_104_v0_, 90, d_103_globalState_)
                d_208_v76_ = out16_
                d_209_v77_ = out17_
            d_191_cf9_ = d_193_cf7_
        elif source0_.is_DC1:
            d_210___mcc_h9_ = source0_.cf1
            d_211_cf1_ = d_210___mcc_h9_
            d_212_v78_: _dafny.Map
            d_212_v78_ = _dafny.Map({d_112_v6_: d_110_v4_})
            d_213_v79_: _dafny.MultiSet
            d_213_v79_ = _dafny.MultiSet([d_110_v4_, not(((d_212_v78_)[d_112_v6_] if (d_112_v6_) in (d_212_v78_) else d_110_v4_)), d_110_v4_, d_110_v4_, d_110_v4_])
            (d_103_globalState_).f0 = ((d_110_v4_) == (d_110_v4_)) == ((d_213_v79_) == (d_213_v79_))
            d_214_v81_: _dafny.Map
            d_214_v81_ = _dafny.Map({d_211_cf1_: d_211_cf1_})
            d_215_v82_: _dafny.Seq
            d_215_v82_ = _dafny.SeqWithoutIsStrInference([d_214_v81_])
            def iife17_():
                coll1_ = _dafny.Map()
                compr_1_: _dafny.Map
                for compr_1_ in (d_215_v82_).Elements:
                    d_216_v80_: _dafny.Map = compr_1_
                    if (d_216_v80_) in (d_215_v82_):
                        coll1_[d_216_v80_] = d_110_v4_
                return _dafny.Map(coll1_)
            (d_103_globalState_).f1 = default__.safeDivisionInt(default__.safeModuloInt(len(iife17_()
            ), d_211_cf1_), (105) + (d_104_v0_))
            index45_ = default__.safeIndex(945, (d_152_v44_).length(0))
            (d_152_v44_)[index45_] = d_104_v0_
            d_217_v83_: _dafny.Map
            d_217_v83_ = _dafny.Map({d_107_v2_: d_110_v4_})
            index46_ = default__.safeIndex(662, (d_105_v1_).length(0))
            index47_ = default__.safeIndex(945, (d_152_v44_).length(0))
            rhs36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
            rhs37_ = d_104_v0_
            rhs38_ = default__.safeDivisionInt(default__.safeDivisionInt(len(d_217_v83_), d_104_v0_), (d_104_v0_) - (d_211_cf1_))
            rhs39_ = (len((d_160_v51_) | (d_160_v51_)) if d_110_v4_ else d_211_cf1_)
            rhs40_ = d_104_v0_
            lhs27_ = d_105_v1_
            lhs28_ = default__.safeIndex(662, (d_105_v1_).length(0))
            lhs29_ = d_152_v44_
            lhs30_ = default__.safeIndex(945, (d_152_v44_).length(0))
            lhs31_ = d_103_globalState_
            lhs27_[lhs28_] = rhs36_
            lhs29_[lhs30_] = rhs37_
            lhs31_.f1 = rhs38_
            d_104_v0_ = rhs39_
            d_211_cf1_ = rhs40_
            if (d_110_v4_) == ((d_110_v4_ if d_110_v4_ else d_110_v4_)):
                index48_ = default__.safeIndex(662, (d_105_v1_).length(0))
                (d_105_v1_)[index48_] = (d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]
                index49_ = default__.safeIndex(662, (d_105_v1_).length(0))
                (d_105_v1_)[index49_] = d_107_v2_
                (d_103_globalState_).f0 = False
                (d_103_globalState_).f0 = d_110_v4_
                d_218_v84_: _dafny.Map
                d_218_v84_ = _dafny.Map({(d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]: d_164_v52_})
                d_218_v84_ = (d_218_v84_).set((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))], d_164_v52_)
            elif True:
                index50_ = default__.safeIndex(230, (d_112_v6_).length(0))
                (d_112_v6_)[index50_] = d_110_v4_
                index51_ = default__.safeIndex(230, (d_112_v6_).length(0))
                (d_112_v6_)[index51_] = d_110_v4_
                (d_103_globalState_).f0 = d_110_v4_
                d_219_v85_: _dafny.Array
                nw35_ = _dafny.Array(_dafny.Map({}), 2)
                d_219_v85_ = nw35_
                d_220_v86_: _dafny.Seq
                d_220_v86_ = _dafny.SeqWithoutIsStrInference([True, (d_112_v6_)[default__.safeIndex(230, (d_112_v6_).length(0))]])
                d_221_v87_: _dafny.Map
                d_221_v87_ = _dafny.Map({False: d_220_v86_})
                d_222_v88_: _dafny.Map
                d_222_v88_ = d_221_v87_
                index52_ = default__.safeIndex(473, (d_219_v85_).length(0))
                (d_219_v85_)[index52_] = d_222_v88_
                d_223_v89_: bool
                d_223_v89_ = (d_112_v6_)[default__.safeIndex(230, (d_112_v6_).length(0))]
                d_224_v90_: T0
                nw36_ = C0()
                nw36_.ctor__(_dafny.SeqWithoutIsStrInference([d_164_v52_]), default__.fm0(d_110_v4_, (d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))], (d_223_v89_), d_103_globalState_), (d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))], (d_112_v6_)[default__.safeIndex(230, (d_112_v6_).length(0))])
                d_224_v90_ = nw36_
                index53_ = default__.safeIndex(473, (d_219_v85_).length(0))
                rhs41_ = d_222_v88_
                rhs42_ = d_224_v90_
                lhs32_ = d_219_v85_
                lhs33_ = default__.safeIndex(473, (d_219_v85_).length(0))
                lhs32_[lhs33_] = rhs41_
                d_224_v90_ = rhs42_
                d_164_v52_ = default__.fm6(d_103_globalState_)
                d_225_v91_: C0
                nw37_ = C0()
                nw37_.ctor__((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))], (d_152_v44_)[default__.safeIndex(945, (d_152_v44_).length(0))], _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lf")), (d_112_v6_)[default__.safeIndex(230, (d_112_v6_).length(0))])
                d_225_v91_ = nw37_
                d_226_v92_: _dafny.Map
                d_226_v92_ = _dafny.Map({d_110_v4_: d_225_v91_})
                d_227_v93_: _dafny.Array
                nw38_ = _dafny.Array(None, 26)
                nw38_[int(0)] = d_225_v91_
                nw38_[int(1)] = d_225_v91_
                nw38_[int(2)] = d_225_v91_
                nw38_[int(3)] = d_225_v91_
                nw38_[int(4)] = d_225_v91_
                nw38_[int(5)] = d_225_v91_
                nw38_[int(6)] = d_225_v91_
                nw38_[int(7)] = d_225_v91_
                nw38_[int(8)] = d_225_v91_
                nw38_[int(9)] = d_225_v91_
                nw38_[int(10)] = d_225_v91_
                nw38_[int(11)] = d_225_v91_
                nw38_[int(12)] = d_225_v91_
                nw38_[int(13)] = d_225_v91_
                nw38_[int(14)] = ((d_226_v92_)[(d_112_v6_)[default__.safeIndex(230, (d_112_v6_).length(0))]] if ((d_112_v6_)[default__.safeIndex(230, (d_112_v6_).length(0))]) in (d_226_v92_) else d_225_v91_)
                nw38_[int(15)] = d_225_v91_
                nw38_[int(16)] = d_225_v91_
                nw38_[int(17)] = d_225_v91_
                nw38_[int(18)] = d_225_v91_
                nw38_[int(19)] = d_225_v91_
                nw38_[int(20)] = d_225_v91_
                nw38_[int(21)] = d_225_v91_
                nw38_[int(22)] = d_225_v91_
                nw38_[int(23)] = d_225_v91_
                nw38_[int(24)] = d_225_v91_
                nw38_[int(25)] = d_225_v91_
                d_227_v93_ = nw38_
                d_228_v94_: _dafny.Array
                nw39_ = _dafny.Array(None, 7)
                nw39_[int(0)] = d_227_v93_
                nw39_[int(1)] = d_227_v93_
                nw39_[int(2)] = d_227_v93_
                nw39_[int(3)] = d_227_v93_
                nw39_[int(4)] = d_227_v93_
                nw39_[int(5)] = (d_227_v93_ if d_110_v4_ else d_227_v93_)
                nw39_[int(6)] = d_227_v93_
                d_228_v94_ = nw39_
                index54_ = default__.safeIndex(477, (d_228_v94_).length(0))
                (d_228_v94_)[index54_] = d_227_v93_
                index55_ = default__.safeIndex(477, (d_228_v94_).length(0))
                (d_228_v94_)[index55_] = d_227_v93_
        elif True:
            d_229___mcc_h10_ = source0_.cf11
            d_230_cf11_ = d_229___mcc_h10_
            d_231_v95_: C0
            nw40_ = C0()
            nw40_.ctor__((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))], d_104_v0_, d_107_v2_, not(not(d_110_v4_)))
            d_231_v95_ = nw40_
            d_232_v96_: _dafny.Map
            d_232_v96_ = _dafny.Map({d_231_v95_: d_112_v6_})
            d_232_v96_ = (d_232_v96_).set(d_231_v95_, d_112_v6_)
            d_233_v97_: bool
            d_233_v97_ = d_110_v4_
            (d_103_globalState_).f0 = (d_110_v4_ if (d_233_v97_) else d_110_v4_)
            d_234_v98_: _dafny.Seq
            d_234_v98_ = _dafny.SeqWithoutIsStrInference([not(d_110_v4_), d_110_v4_])
            d_235_v99_: T0
            nw41_ = C0()
            nw41_.ctor__((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))], (d_231_v95_).fm4((d_231_v95_).f8, d_104_v0_, d_234_v98_, d_103_globalState_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xalwfr")), d_110_v4_)
            d_235_v99_ = nw41_
            d_235_v99_ = d_235_v99_
            d_236_v100_: _dafny.Map
            d_236_v100_ = _dafny.Map({len(((d_235_v99_).f5) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "le")))): d_110_v4_})
            d_236_v100_ = (d_236_v100_).set(default__.safeModuloInt((d_231_v95_).f8, (d_167_v55_)[default__.safeIndex((d_231_v95_).f8, len(d_167_v55_))]), d_110_v4_)
        index56_ = default__.safeIndex(75, (d_112_v6_).length(0))
        (d_112_v6_)[index56_] = not(d_110_v4_)
        index57_ = default__.safeIndex(75, (d_112_v6_).length(0))
        (d_112_v6_)[index57_] = not(False)
        d_107_v2_ = (default__.fm2(((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))] if d_110_v4_ else (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]), d_110_v4_, d_104_v0_, d_110_v4_, d_103_globalState_)).set(default__.safeIndex((d_104_v0_) + (default__.fm0((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))], d_107_v2_, default__.fm1(d_104_v0_, False, d_103_globalState_), d_103_globalState_)), len(default__.fm2(((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))] if d_110_v4_ else (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]), d_110_v4_, d_104_v0_, d_110_v4_, d_103_globalState_))), d_164_v52_)
        index58_ = default__.safeIndex(579, (d_152_v44_).length(0))
        (d_152_v44_)[index58_] = len(d_107_v2_)
        index59_ = default__.safeIndex(579, (d_152_v44_).length(0))
        (d_152_v44_)[index59_] = 82
        d_237_v101_: _dafny.Set
        d_237_v101_ = _dafny.Set({d_152_v44_, d_152_v44_})
        d_238_v102_: _dafny.Map
        d_238_v102_ = _dafny.Map({d_237_v101_: (True) or ((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])})
        d_238_v102_ = (d_238_v102_).set(_dafny.Set({d_152_v44_}), (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])
        hi2_ = (0) - (d_104_v0_)
        for d_239_i5_ in range(d_104_v0_, hi2_):
            d_240_v103_: C0
            nw42_ = C0()
            nw42_.ctor__(d_107_v2_, default__.safeDivisionInt(d_104_v0_, (d_152_v44_)[default__.safeIndex(579, (d_152_v44_).length(0))]), ((d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]) + (((d_108_v3_)[not((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])] if (not((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])) in (d_108_v3_) else _dafny.SeqWithoutIsStrInference([d_164_v52_ for d_241_i6_ in range(default__.abs(902))]))), (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])
            d_240_v103_ = nw42_
            if (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]:
                d_242_v105_: _dafny.MultiSet
                d_242_v105_ = _dafny.MultiSet([(d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))], d_110_v4_])
                d_243_v106_: _dafny.Map
                def iife18_():
                    coll2_ = _dafny.Map()
                    compr_2_: int
                    for compr_2_ in _dafny.IntegerRange(79, -6):
                        d_244_v104_: int = compr_2_
                        if ((79) <= (d_244_v104_)) and ((d_244_v104_) < (-6)):
                            coll2_[(d_244_v104_) + (len((d_240_v103_).f7))] = (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]
                    return _dafny.Map(coll2_)
                d_243_v106_ = _dafny.Map({len(iife18_()
                ): d_242_v105_})
                index60_ = default__.safeIndex(662, (d_105_v1_).length(0))
                rhs43_ = (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]
                rhs44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkbtcsqb"))
                rhs45_ = d_104_v0_
                rhs46_ = not((2) not in ((d_243_v106_ if (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))] else d_243_v106_)))
                rhs47_ = not((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])
                lhs34_ = d_105_v1_
                lhs35_ = default__.safeIndex(662, (d_105_v1_).length(0))
                lhs36_ = d_103_globalState_
                lhs37_ = d_103_globalState_
                d_110_v4_ = rhs43_
                lhs34_[lhs35_] = rhs44_
                lhs36_.f1 = rhs45_
                lhs37_.f0 = rhs46_
                d_110_v4_ = rhs47_
                index61_ = default__.safeIndex(662, (d_105_v1_).length(0))
                rhs48_ = (d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]
                rhs49_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilf"))) + ((d_107_v2_ if True else (d_105_v1_)[default__.safeIndex(662, (d_105_v1_).length(0))]))
                rhs50_ = (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]
                lhs38_ = d_105_v1_
                lhs39_ = default__.safeIndex(662, (d_105_v1_).length(0))
                lhs40_ = d_103_globalState_
                d_107_v2_ = rhs48_
                lhs38_[lhs39_] = rhs49_
                lhs40_.f0 = rhs50_
                d_245_v107_: _dafny.Map
                d_245_v107_ = _dafny.Map({d_110_v4_: d_240_v103_})
                d_246_v108_: _dafny.Seq
                d_246_v108_ = _dafny.SeqWithoutIsStrInference([d_240_v103_])
                d_247_v110_: _dafny.Seq
                def iife19_():
                    coll3_ = _dafny.Set()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(172, -314):
                        d_248_v109_: int = compr_3_
                        if ((172) <= (d_248_v109_)) and ((d_248_v109_) < (-314)):
                            coll3_ = coll3_.union(_dafny.Set([default__.safeDivisionInt(d_248_v109_, d_104_v0_)]))
                    return _dafny.Set(coll3_)
                d_247_v110_ = _dafny.SeqWithoutIsStrInference([d_111_v5_, d_111_v5_, iife19_()
                , d_111_v5_, d_111_v5_])
                d_249_v111_: _dafny.Map
                d_249_v111_ = _dafny.Map({d_239_i5_: (d_111_v5_).ispropersubset((d_247_v110_)[default__.safeIndex(d_104_v0_, len(d_247_v110_))])})
                d_250_v112_: _dafny.Seq
                d_250_v112_ = _dafny.SeqWithoutIsStrInference([d_110_v4_, d_110_v4_, (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))], not(d_110_v4_)])
                rhs51_ = _dafny.Map({(d_110_v4_) or (d_110_v4_): d_240_v103_})
                rhs52_ = ((d_240_v103_).f7)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(d_167_v55_)[default__.safeIndex(len(d_246_v108_), len(d_167_v55_))], len((d_240_v103_).f7), ((d_242_v105_)[(d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]] if ((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]) in (d_242_v105_) else len(default__.fm7((d_240_v103_).f8, d_103_globalState_))), d_239_i5_])), len((d_240_v103_).f7))]
                rhs53_ = ((d_249_v111_)[len((d_250_v112_) + (d_250_v112_))] if (len((d_250_v112_) + (d_250_v112_))) in (d_249_v111_) else d_110_v4_)
                lhs41_ = d_103_globalState_
                d_245_v107_ = rhs51_
                d_164_v52_ = rhs52_
                lhs41_.f0 = rhs53_
                d_251_v113_: _dafny.Array
                nw43_ = _dafny.Array(None, 19)
                nw43_[int(0)] = d_112_v6_
                nw43_[int(1)] = d_112_v6_
                nw43_[int(2)] = d_112_v6_
                nw43_[int(3)] = d_112_v6_
                nw43_[int(4)] = d_112_v6_
                nw43_[int(5)] = d_112_v6_
                nw43_[int(6)] = d_112_v6_
                nw43_[int(7)] = d_112_v6_
                nw43_[int(8)] = (D4_DC7(d_112_v6_)).cf14
                nw43_[int(9)] = d_112_v6_
                nw43_[int(10)] = d_112_v6_
                nw43_[int(11)] = d_112_v6_
                nw43_[int(12)] = d_112_v6_
                nw43_[int(13)] = d_112_v6_
                nw43_[int(14)] = d_112_v6_
                nw43_[int(15)] = d_112_v6_
                nw43_[int(16)] = d_112_v6_
                nw43_[int(17)] = d_112_v6_
                nw43_[int(18)] = d_112_v6_
                d_251_v113_ = nw43_
                d_251_v113_ = d_251_v113_
                index62_ = default__.safeIndex(579, (d_152_v44_).length(0))
                (d_152_v44_)[index62_] = (d_240_v103_).f8
            elif True:
                d_252_v114_: _dafny.Seq
                d_252_v114_ = d_153_v45_
                index63_ = default__.safeIndex(579, (d_152_v44_).length(0))
                rhs54_ = d_152_v44_
                rhs55_ = d_240_v103_
                rhs56_ = not((d_111_v5_) == (d_111_v5_))
                rhs57_ = d_110_v4_
                rhs58_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mxtsyoj")))) - ((len(d_167_v55_)) + (len((d_252_v114_))))
                lhs42_ = d_103_globalState_
                lhs43_ = d_103_globalState_
                lhs44_ = d_152_v44_
                lhs45_ = default__.safeIndex(579, (d_152_v44_).length(0))
                d_152_v44_ = rhs54_
                d_240_v103_ = rhs55_
                lhs42_.f0 = rhs56_
                lhs43_.f0 = rhs57_
                lhs44_[lhs45_] = rhs58_
                d_253_v115_: C0
                nw44_ = C0()
                nw44_.ctor__((d_240_v103_).f7, d_239_i5_, d_107_v2_, (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))])
                d_253_v115_ = nw44_
                d_254_v116_: D6
                d_254_v116_ = D6_DC12(d_150_v42_)
                d_255_v117_: _dafny.MultiSet
                d_255_v117_ = _dafny.MultiSet([d_110_v4_, ((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]) and (d_110_v4_), ((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))] if (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))] else d_110_v4_), ((d_254_v116_).cf25).isdisjoint(d_150_v42_), ((d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]) not in (_dafny.Map({(d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]: (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))]}))])
                index64_ = default__.safeIndex(579, (d_152_v44_).length(0))
                (d_152_v44_)[index64_] = (d_255_v117_).cardinality
                d_256_v118_: _dafny.Array
                nw45_ = _dafny.Array(None, 8)
                nw45_[int(0)] = d_112_v6_
                nw45_[int(1)] = d_112_v6_
                nw45_[int(2)] = d_112_v6_
                nw45_[int(3)] = d_112_v6_
                nw45_[int(4)] = d_112_v6_
                nw45_[int(5)] = d_112_v6_
                nw45_[int(6)] = d_112_v6_
                nw45_[int(7)] = (d_112_v6_ if d_110_v4_ else d_112_v6_)
                d_256_v118_ = nw45_
                index65_ = default__.safeIndex(391, (d_256_v118_).length(0))
                (d_256_v118_)[index65_] = d_112_v6_
                index66_ = default__.safeIndex(391, (d_256_v118_).length(0))
                (d_256_v118_)[index66_] = d_112_v6_
                d_257_v119_: _dafny.Array
                nw46_ = _dafny.Array(_dafny.Map({}), 1)
                d_257_v119_ = nw46_
                d_258_v121_: _dafny.MultiSet
                d_258_v121_ = _dafny.MultiSet([d_107_v2_, (d_240_v103_).f7])
                d_259_v122_: D6
                def iife20_():
                    coll4_ = _dafny.Map()
                    compr_4_: _dafny.Seq
                    for compr_4_ in (d_258_v121_).Elements:
                        d_260_v120_: _dafny.Seq = compr_4_
                        if (d_260_v120_) in (d_258_v121_):
                            coll4_[d_260_v120_] = d_104_v0_
                    return _dafny.Map(coll4_)
                d_259_v122_ = D6_DC15((d_240_v103_).f8, d_253_v115_, iife20_()
, default__.fm1((d_255_v117_).cardinality, (d_112_v6_)[default__.safeIndex(75, (d_112_v6_).length(0))], d_103_globalState_))
                d_261_v123_: D6
                d_261_v123_ = D6_DC16(d_259_v122_)
                d_262_v124_: _dafny.Map
                d_262_v124_ = _dafny.Map({d_261_v123_: len((d_240_v103_).f7)})
                index67_ = default__.safeIndex(218, (d_257_v119_).length(0))
                (d_257_v119_)[index67_] = (d_262_v124_) | (d_262_v124_)
                index68_ = default__.safeIndex(218, (d_257_v119_).length(0))
                (d_257_v119_)[index68_] = (d_262_v124_) | ((d_262_v124_) | (d_262_v124_))
            index69_ = default__.safeIndex(579, (d_152_v44_).length(0))
            (d_152_v44_)[index69_] = d_239_i5_
            (d_103_globalState_).f0 = d_110_v4_
        _dafny.print(_dafny.string_of(d_103_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_103_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_104_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_105_v1_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_107_v2_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_108_v3_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eedslel"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_110_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_v5_) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_112_v6_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_113_v7_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[0]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[1]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[2]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[3]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[4]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[5]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[6]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[7]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_147_v41_)[8]).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v42_) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v43_) == (_dafny.Map({True: _dafny.MultiSet([1])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v44_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_153_v45_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_154_v46_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v47_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_156_v48_).cf11).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v50_) == (_dafny.Map({True: 524}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v51_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_164_v52_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v53_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v53_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v53_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_165_v53_).cf10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v54_) == (_dafny.Set({D1_DC3(False, 1, True, _dafny.CodePoint('f'))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_167_v55_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_237_v101_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_238_v102_)))
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
        return lambda: D1_DC2(_dafny.Seq({}), int(0), int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {self.cf5.VerbatimString(True)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC5(D2, NamedTuple('DC5', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC6(D3, NamedTuple('DC6', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC6({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC6) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC8(int(0), None, False, False, None)
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
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC8(D4, NamedTuple('DC8', [('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC8({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC8) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC7(D4, NamedTuple('DC7', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC7({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC7) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D5_DC11)

class D5_DC11(D5, NamedTuple('DC11', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC11({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC11) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC13(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D6_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D6_DC12)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC13(D6, NamedTuple('DC13', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC13({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC13) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC14(D6, NamedTuple('DC14', [('cf28', Any), ('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC12(D6, NamedTuple('DC12', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC12({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC12) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    @property
    def f6(self):
        return self._f6
    @f6.setter
    def f6(self, value):
        self._f6 = value

class GlobalState:
    def  __init__(self):
        self.f0: bool = False
        self.f1: int = int(0)
        self._f2: str = _dafny.CodePoint('D')
        self._f3: int = int(0)
        self._f4: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4):
        (self).f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4

    @property
    def f2(self):
        return self._f2
    @property
    def f3(self):
        return self._f3
    @property
    def f4(self):
        return self._f4

class C0(T0, T1):
    def  __init__(self):
        self._f6: bool = False
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f7: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f6(self):
        return self._f6
    @f6.setter
    def f6(self, value):
        self._f6 = value
    @property
    def f5(self):
        return self._f5
    def ctor__(self, f7, f8, f5, f6):
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f5 = f5
        (self).f6 = f6

    def fm4(self, p0, p1, p2, globalState):
        return default__.safeDivisionInt((self).f8, ((self).f8 if self.f6 else (self).f8))

    def fm5(self, p0, p1, globalState):
        return (_dafny.Map({self.f6: _dafny.SeqWithoutIsStrInference([self.f6])}))

    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
