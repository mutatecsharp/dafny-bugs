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
        return len((_dafny.Set({not(False)})) | (_dafny.Set({False})))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([not(False)])).intersection((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([False])) if False else _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(True)]))))

    @staticmethod
    def fm2(p0, p1, globalState):
        return not(True)

    @staticmethod
    def fm3(globalState):
        return (_dafny.Set({True, False, False})) - ((_dafny.Set({True})).intersection(_dafny.Set({False, not(False)})))

    @staticmethod
    def fm4(p0, p1, p2, globalState):
        return ((D5_DC14(_dafny.SeqWithoutIsStrInference([(D4_DC13(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({642})), len(_dafny.Map({True: False})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_0_i0_ in range(default__.abs(336))]))])))).cf17])) if True else D5_DC14(_dafny.SeqWithoutIsStrInference([675, 827, len(_dafny.Map({718: -405}))])))).cf18

    @staticmethod
    def fm5(p0, globalState):
        source0_ = D2_DC8(True, False)
        if source0_.is_DC7:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ammm"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1_i0_ in range(default__.abs(832))]))
        elif source0_.is_DC8:
            d_2___mcc_h0_ = source0_.cf9
            d_3___mcc_h1_ = source0_.cf10
            d_4_cf10_ = d_3___mcc_h1_
            d_5_cf9_ = d_2___mcc_h0_
            return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "widm"))
        elif True:
            d_6___mcc_h2_ = source0_.cf8
            d_7_cf8_ = d_6___mcc_h2_
            if True:
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ult"))
            elif True:
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krthaq"))

    @staticmethod
    def fm8(globalState):
        source1_ = (D5_DC14(_dafny.SeqWithoutIsStrInference([268])) if not(False) else D5_DC14(_dafny.SeqWithoutIsStrInference([-267])))
        if source1_.is_DC15:
            d_8___mcc_h0_ = source1_.cf19
            d_9___mcc_h1_ = source1_.cf20
            d_10_cf20_ = d_9___mcc_h1_
            d_11_cf19_ = d_8___mcc_h0_
            if d_10_cf20_:
                return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({d_10_cf20_: d_10_cf20_})))]))
            elif True:
                return _dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('n'): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "isbarghd"))}))])
        elif True:
            d_12___mcc_h2_ = source1_.cf18
            d_13_cf18_ = d_12___mcc_h2_
            return _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not(False)]))])

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: str
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_14_i0_ in range(default__.abs(951))])).Elements:
                d_15_v0_: str = compr_0_
                if (d_15_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_14_i0_ in range(default__.abs(951))])):
                    coll0_ = coll0_.union(_dafny.Set([d_15_v0_]))
            return _dafny.Set(coll0_)
        return (iife0_()
        ).intersection(_dafny.Set({_dafny.CodePoint('f'), _dafny.CodePoint('m'), _dafny.CodePoint('i'), _dafny.CodePoint('f'), _dafny.CodePoint('x')}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: str
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])).Elements:
                d_16_v0_: str = compr_1_
                if (d_16_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v')])):
                    coll1_ = coll1_.union(_dafny.Set([d_16_v0_]))
            return _dafny.Set(coll1_)
        return _dafny.SeqWithoutIsStrInference([False, (-386) > (486), not((iife1_()
        ).issubset(_dafny.Set({_dafny.CodePoint('s'), _dafny.CodePoint('q')}))), not((False if False else not(not(True))))])

    @staticmethod
    def fm11(globalState):
        if (_dafny.Set({True, not(False)})).issubset(_dafny.Set({True, not(False)})):
            return _dafny.Map({_dafny.SeqWithoutIsStrInference([False]): not(False)})
        elif True:
            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, True]): True}))

    @staticmethod
    def fm12(p0, p1, p2, p3, globalState):
        return (_dafny.Map({-464: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "creunkjsj"))})) | ((_dafny.Map({(_dafny.MultiSet([True])).cardinality: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft"))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvvvveiw"))): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_17_i0_ in range(default__.abs(272))])})))

    @staticmethod
    def m0(p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_18_v0_: _dafny.Array
        nw0_ = _dafny.Array(False, 26)
        d_18_v0_ = nw0_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_18_v0_).length(0)):
            d_19_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_19_i0_)) and ((d_19_i0_) < ((d_18_v0_).length(0)))):
                (d_18_v0_)[(d_19_i0_)] = False
        d_20_v1_: int
        d_20_v1_ = 555
        d_21_v2_: D0
        d_21_v2_ = D0_DC0(p0)
        d_22_v3_: C0
        nw1_ = C0()
        nw1_.ctor__(d_20_v1_, p0, d_21_v2_)
        d_22_v3_ = nw1_
        d_23_v4_: str
        d_23_v4_ = _dafny.CodePoint('p')
        d_24_v5_: _dafny.Set
        d_24_v5_ = _dafny.Set({d_23_v4_})
        d_25_v6_: _dafny.Set
        d_25_v6_ = _dafny.Set({d_24_v5_, d_24_v5_, d_24_v5_, d_24_v5_, default__.fm9((d_22_v3_).f12, d_20_v1_, len(p1), globalState)})
        hi0_ = len(_dafny.Map({d_22_v3_: len(d_25_v6_)}))
        for d_26_i1_ in range((d_20_v1_ if p0 else d_20_v1_), hi0_):
            d_27_v7_: _dafny.Seq
            d_27_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lotbhuwwl"))
            d_28_v8_: _dafny.Set
            d_28_v8_ = _dafny.Set({len(d_27_v7_)})
            d_29_v9_: _dafny.Seq
            d_29_v9_ = _dafny.SeqWithoutIsStrInference([p1])
            d_30_v10_: _dafny.Map
            d_30_v10_ = _dafny.Map({(d_22_v3_).f12: d_26_i1_})
            d_31_v11_: _dafny.MultiSet
            d_31_v11_ = _dafny.MultiSet([d_18_v0_])
            d_32_v12_: _dafny.MultiSet
            d_32_v12_ = d_31_v11_
            d_33_v13_: _dafny.Array
            nw2_ = _dafny.Array(None, 22)
            nw2_[int(0)] = p1
            nw2_[int(1)] = (p1).set(default__.safeIndex(d_26_i1_, len(p1)), not(not(p0)))
            nw2_[int(2)] = p1
            nw2_[int(3)] = p1
            nw2_[int(4)] = p1
            nw2_[int(5)] = p1
            nw2_[int(6)] = (d_29_v9_)[default__.safeIndex(d_20_v1_, len(d_29_v9_))]
            nw2_[int(7)] = (p1) + ((_dafny.SeqWithoutIsStrInference([not(True)])).set(default__.safeIndex((d_22_v3_).f12, len(_dafny.SeqWithoutIsStrInference([not(True)]))), p0))
            nw2_[int(8)] = (p1) + (p1)
            nw2_[int(9)] = default__.fm10(880, p0, (d_22_v3_).f12, d_26_i1_, globalState)
            nw2_[int(10)] = p1
            nw2_[int(11)] = p1
            nw2_[int(12)] = p1
            nw2_[int(13)] = (p1 if p0 else _dafny.SeqWithoutIsStrInference([p0]))
            nw2_[int(14)] = default__.fm10((d_22_v3_).f12, p0, ((d_30_v10_)[((d_32_v12_)).cardinality] if (((d_32_v12_)).cardinality) in (d_30_v10_) else d_26_i1_), d_26_i1_, globalState)
            nw2_[int(15)] = p1
            nw2_[int(16)] = (p1 if p0 else (p1).set(default__.safeIndex(d_20_v1_, len(p1)), p0))
            nw2_[int(17)] = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0])
            nw2_[int(18)] = (p1 if p0 else p1)
            nw2_[int(19)] = _dafny.SeqWithoutIsStrInference([p0, p0, p0, p0, not(p0)])
            nw2_[int(20)] = _dafny.SeqWithoutIsStrInference([default__.fm2(p0, p1, globalState), p0])
            nw2_[int(21)] = p1
            d_33_v13_ = nw2_
            index0_ = default__.safeIndex(907, (d_33_v13_).length(0))
            (d_33_v13_)[index0_] = p1
            index1_ = default__.safeIndex(907, (d_33_v13_).length(0))
            rhs0_ = d_28_v8_
            rhs1_ = p0
            rhs2_ = p1
            rhs3_ = d_22_v3_
            rhs4_ = p0
            lhs0_ = globalState
            lhs1_ = d_33_v13_
            lhs2_ = default__.safeIndex(907, (d_33_v13_).length(0))
            lhs3_ = globalState
            d_28_v8_ = rhs0_
            lhs0_.f5 = rhs1_
            lhs1_[lhs2_] = rhs2_
            d_22_v3_ = rhs3_
            lhs3_.f5 = rhs4_
            d_34_v14_: D2
            d_34_v14_ = D2_DC7()
            d_35_v15_: _dafny.Set
            d_35_v15_ = _dafny.Set({d_34_v14_})
            if (((d_22_v3_).f12) * (len(d_35_v15_))) > (d_26_i1_):
                d_36_v16_: _dafny.Array
                nw3_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_36_v16_ = nw3_
                d_36_v16_ = d_36_v16_
                d_37_v17_: C0
                nw4_ = C0()
                nw4_.ctor__((len(d_27_v7_)) * (d_26_i1_), (p0 if p0 else p0), d_21_v2_)
                d_37_v17_ = nw4_
                pat_let_tv0_ = p0
                pat_let_tv1_ = p0
                pat_let_tv2_ = p0
                def iife2_(_pat_let0_0):
                    def iife3_(d_38_dt__update__tmp_h0_):
                        def iife4_(_pat_let1_0):
                            def iife5_(d_39_dt__update_hcf0_h0_):
                                return D0_DC0(d_39_dt__update_hcf0_h0_)
                            return iife5_(_pat_let1_0)
                        return iife4_((pat_let_tv0_ if not(pat_let_tv1_) else pat_let_tv2_))
                    return iife3_(_pat_let0_0)
                rhs5_ = not(not (p0) or (True))
                rhs6_ = (d_26_i1_) - (default__.safeDivisionInt((d_37_v17_).f12, (0) - (d_26_i1_)))
                rhs7_ = iife2_(d_21_v2_)
                rhs8_ = not(not (p0) or (p0))
                lhs4_ = globalState
                lhs5_ = globalState
                lhs4_.f8 = rhs5_
                d_20_v1_ = rhs6_
                d_21_v2_ = rhs7_
                lhs5_.f8 = rhs8_
                d_40_v18_: _dafny.Seq
                d_40_v18_ = _dafny.SeqWithoutIsStrInference([(d_22_v3_).f12])
                d_41_v19_: _dafny.Map
                d_41_v19_ = _dafny.Map({d_40_v18_: d_37_v17_})
                d_42_v20_: _dafny.Map
                d_42_v20_ = _dafny.Map({(_dafny.Map({d_40_v18_: d_37_v17_})) | (d_41_v19_): (d_30_v10_) | (d_30_v10_)})
                d_42_v20_ = (d_42_v20_).set((d_41_v19_).set(d_40_v18_, d_37_v17_), _dafny.Map({485: (d_22_v3_).f12}))
                d_43_v21_: _dafny.Seq
                d_43_v21_ = _dafny.SeqWithoutIsStrInference([d_37_v17_, d_22_v3_, d_37_v17_, d_22_v3_])
                d_43_v21_ = ((_dafny.SeqWithoutIsStrInference([d_37_v17_, d_22_v3_])) + (d_43_v21_)) + (d_43_v21_)
            elif True:
                d_44_v22_: _dafny.Array
                nw5_ = _dafny.Array(None, 4)
                nw5_[int(0)] = -460
                nw5_[int(1)] = (d_22_v3_).f12
                nw5_[int(2)] = default__.safeModuloInt((d_22_v3_).f12, (d_22_v3_).f12)
                nw5_[int(3)] = (d_22_v3_).f12
                d_44_v22_ = nw5_
                d_44_v22_ = d_44_v22_
                d_45_v23_: _dafny.MultiSet
                d_45_v23_ = _dafny.MultiSet([default__.fm0((0) - (d_20_v1_), _dafny.Map({p1: p0}), globalState)])
                d_46_v24_: D0
                d_46_v24_ = D0_DC2(p0, (d_45_v23_).cardinality)
                d_47_v25_: _dafny.Map
                d_47_v25_ = _dafny.Map({((d_33_v13_)[default__.safeIndex(907, (d_33_v13_).length(0))]).set(default__.safeIndex(9, len((d_33_v13_)[default__.safeIndex(907, (d_33_v13_).length(0))])), p0): p0})
                d_48_v26_: _dafny.Seq
                d_48_v26_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({p1: p0}), d_47_v25_, d_47_v25_])
                rhs9_ = (0) - ((d_22_v3_).f12)
                rhs10_ = ((d_45_v23_)[(d_22_v3_).f12] if ((d_22_v3_).f12) in (d_45_v23_) else default__.fm0((d_22_v3_).f12, (d_48_v26_)[default__.safeIndex(d_26_i1_, len(d_48_v26_))], globalState))
                rhs11_ = d_23_v4_
                rhs12_ = d_46_v24_
                d_20_v1_ = rhs9_
                d_20_v1_ = rhs10_
                d_23_v4_ = rhs11_
                d_46_v24_ = rhs12_
                (globalState).f5 = True
                d_49_v27_: _dafny.Seq
                d_49_v27_ = _dafny.SeqWithoutIsStrInference([default__.fm0(len(d_27_v7_), default__.fm11(globalState), globalState), len((d_33_v13_)[default__.safeIndex(907, (d_33_v13_).length(0))])])
                d_50_v28_: _dafny.Map
                d_50_v28_ = _dafny.Map({d_23_v4_: (d_22_v3_).f12})
                (globalState).f5 = ((d_49_v27_)[default__.safeIndex(571, len(d_49_v27_))]) in ((default__.fm12(d_30_v10_, d_28_v8_, len(d_50_v28_), p0, globalState)).set((d_22_v3_).f12, d_27_v7_))
                d_51_v29_: T0
                nw6_ = C0()
                nw6_.ctor__(d_20_v1_, p0, d_21_v2_)
                d_51_v29_ = nw6_
                d_52_v30_: _dafny.Map
                d_52_v30_ = _dafny.Map({len(p1): d_51_v29_})
                d_53_v31_: _dafny.Set
                d_53_v31_ = _dafny.Set({d_52_v30_})
                d_53_v31_ = d_53_v31_
            d_54_v32_: _dafny.Set
            d_54_v32_ = _dafny.Set({p0})
            d_55_v33_: _dafny.Map
            d_55_v33_ = _dafny.Map({459: d_54_v32_})
            d_56_v34_: _dafny.MultiSet
            d_56_v34_ = _dafny.MultiSet([len((d_55_v33_) | (d_55_v33_)), ((d_22_v3_).f12) + (d_20_v1_), d_20_v1_, 865, (d_26_i1_) + ((d_22_v3_).f12)])
            d_56_v34_ = d_56_v34_
            d_22_v3_ = d_22_v3_
        d_57_v35_: _dafny.Map
        d_57_v35_ = _dafny.Map({p0: _dafny.MultiSet([p0])})
        d_58_v36_: _dafny.Map
        d_58_v36_ = _dafny.Map({p0: p0})
        (globalState).f8 = (d_22_v3_).fm7(d_21_v2_, d_57_v35_, (d_20_v1_ if ((d_58_v36_)[p0] if (p0) in (d_58_v36_) else p0) else 609), globalState)
        d_59_v37_: T0
        nw7_ = C0()
        nw7_.ctor__((0) - (-566), p0, d_21_v2_)
        d_59_v37_ = nw7_
        d_60_v38_: _dafny.Map
        d_60_v38_ = _dafny.Map({(d_22_v3_).f12: d_59_v37_})
        d_61_v39_: _dafny.MultiSet
        d_61_v39_ = _dafny.MultiSet([d_59_v37_, ((d_60_v38_)[(d_22_v3_).f12] if ((d_22_v3_).f12) in (d_60_v38_) else d_59_v37_)])
        d_62_v40_: _dafny.Map
        d_62_v40_ = _dafny.Map({d_61_v39_: not((d_20_v1_) != (922))})
        (globalState).f8 = ((d_62_v40_)[d_61_v39_] if (d_61_v39_) in (d_62_v40_) else (d_59_v37_).f10)
        d_63_i2_: int
        d_63_i2_ = 0
        with _dafny.label("0"):
            while (D2_DC8((d_59_v37_).f10, p0)).cf10:
                with _dafny.c_label("0"):
                    if (d_63_i2_) >= (100):
                        raise _dafny.Break("0")
                    d_63_i2_ = (d_63_i2_) + (1)
                    rhs13_ = 107
                    rhs14_ = d_20_v1_
                    rhs15_ = False
                    lhs6_ = globalState
                    d_20_v1_ = rhs13_
                    d_20_v1_ = rhs14_
                    lhs6_.f8 = rhs15_
                    if ((d_22_v3_).f12) == ((d_22_v3_).f12):
                        d_20_v1_ = (0) - (d_20_v1_)
                        d_64_v41_: _dafny.Seq
                        d_64_v41_ = _dafny.SeqWithoutIsStrInference([(d_22_v3_).f12, d_20_v1_])
                        d_65_v42_: D2
                        d_65_v42_ = D2_DC8((d_59_v37_).f10, (d_59_v37_).f10)
                        d_66_v43_: _dafny.MultiSet
                        d_66_v43_ = _dafny.MultiSet([(d_59_v37_).f10])
                        d_67_v44_: _dafny.Seq
                        d_67_v44_ = _dafny.SeqWithoutIsStrInference([d_66_v43_])
                        d_68_v45_: _dafny.Array
                        def lambda0_(d_69_v1_):
                            def lambda1_(d_70_i3_):
                                return (d_70_i3_) * (d_69_v1_)

                            return lambda1_

                        init0_ = lambda0_(d_20_v1_)
                        nw8_ = _dafny.Array(None, 18)
                        for i0_0_ in range(nw8_.length(0)):
                            nw8_[i0_0_] = init0_(i0_0_)
                        d_68_v45_ = nw8_
                        d_71_v46_: _dafny.Map
                        d_71_v46_ = _dafny.Map({(p1)[default__.safeIndex((d_22_v3_).f12, len(p1))]: d_68_v45_})
                        rhs16_ = (d_64_v41_) + ((d_64_v41_ if (d_59_v37_).f10 else _dafny.SeqWithoutIsStrInference([(d_22_v3_).f12])))
                        rhs17_ = not ((d_66_v43_) not in (d_67_v44_)) or (p0)
                        rhs18_ = (854) * (((d_22_v3_).f12) * (len(d_71_v46_)))
                        rhs19_ = d_65_v42_
                        rhs20_ = ((d_60_v38_)[(d_22_v3_).f12] if ((d_22_v3_).f12) in (d_60_v38_) else d_59_v37_)
                        lhs7_ = globalState
                        d_64_v41_ = rhs16_
                        lhs7_.f8 = rhs17_
                        d_20_v1_ = rhs18_
                        d_65_v42_ = rhs19_
                        d_59_v37_ = rhs20_
                        index2_ = default__.safeIndex(287, (d_68_v45_).length(0))
                        (d_68_v45_)[index2_] = (d_22_v3_).f12
                        d_72_v47_: _dafny.MultiSet
                        d_72_v47_ = _dafny.MultiSet([d_18_v0_])
                        d_73_v48_: _dafny.Set
                        d_73_v48_ = _dafny.Set({d_64_v41_})
                        d_74_v49_: _dafny.Seq
                        d_74_v49_ = _dafny.SeqWithoutIsStrInference([d_22_v3_])
                        d_75_v50_: _dafny.Map
                        d_75_v50_ = _dafny.Map({(d_73_v48_).issubset(d_73_v48_): (d_74_v49_) + (d_74_v49_)})
                        index3_ = default__.safeIndex(287, (d_68_v45_).length(0))
                        rhs21_ = (((d_72_v47_) | (d_72_v47_)).cardinality) * ((d_22_v3_).f12)
                        rhs22_ = (d_22_v3_).f12
                        rhs23_ = (d_22_v3_).f12
                        rhs24_ = len(d_75_v50_)
                        rhs25_ = ((d_59_v37_).f10) or ((d_59_v37_).f10)
                        lhs8_ = d_68_v45_
                        lhs9_ = default__.safeIndex(287, (d_68_v45_).length(0))
                        lhs10_ = globalState
                        lhs8_[lhs9_] = rhs21_
                        d_20_v1_ = rhs22_
                        d_20_v1_ = rhs23_
                        d_20_v1_ = rhs24_
                        lhs10_.f8 = rhs25_
                        d_59_v37_ = d_59_v37_
                        d_76_v51_: _dafny.Array
                        def lambda2_(d_77_v37_, d_78_p0_, d_79_v3_):
                            def lambda3_(d_80_i4_):
                                return D1_DC4((d_77_v37_).f10, _dafny.Map({d_78_p0_: (d_79_v3_).f12}), not((d_77_v37_).f10))

                            return lambda3_

                        init1_ = lambda2_(d_59_v37_, p0, d_22_v3_)
                        nw9_ = _dafny.Array(None, 1)
                        for i0_1_ in range(nw9_.length(0)):
                            nw9_[i0_1_] = init1_(i0_1_)
                        d_76_v51_ = nw9_
                        d_81_v52_: _dafny.Map
                        d_81_v52_ = _dafny.Map({(d_59_v37_).f10: (d_22_v3_).f12})
                        d_82_v53_: D1
                        d_82_v53_ = D1_DC4((d_59_v37_).f10, d_81_v52_, True)
                        index4_ = default__.safeIndex(111, (d_76_v51_).length(0))
                        (d_76_v51_)[index4_] = d_82_v53_
                        index5_ = default__.safeIndex(111, (d_76_v51_).length(0))
                        (d_76_v51_)[index5_] = d_82_v53_
                    elif True:
                        d_83_v54_: D0
                        d_83_v54_ = D0_DC1()
                        d_83_v54_ = d_83_v54_
                        d_84_v55_: D2
                        d_84_v55_ = D2_DC8((d_59_v37_).f10, (d_59_v37_).f10)
                        d_85_v56_: _dafny.MultiSet
                        d_85_v56_ = _dafny.MultiSet([(d_84_v55_).cf9, (d_59_v37_).f10, (d_59_v37_).f10])
                        (globalState).f8 = ((d_59_v37_).f10 if (d_59_v37_).f10 else (d_22_v3_).fm7(D0_DC0(False), _dafny.Map({(d_59_v37_).f10: d_85_v56_}), (d_22_v3_).f12, globalState))
                        d_59_v37_ = d_59_v37_
                        d_59_v37_ = d_59_v37_
                        d_86_v57_: _dafny.Array
                        nw10_ = _dafny.Array(None, 9)
                        d_86_v57_ = nw10_
                        d_87_v58_: _dafny.Map
                        d_87_v58_ = _dafny.Map({(d_22_v3_).f12: d_22_v3_})
                        index6_ = default__.safeIndex(387, (d_86_v57_).length(0))
                        (d_86_v57_)[index6_] = (d_22_v3_ if (d_59_v37_).f10 else ((d_87_v58_)[(d_22_v3_).f12] if ((d_22_v3_).f12) in (d_87_v58_) else d_22_v3_))
                        index7_ = default__.safeIndex(387, (d_86_v57_).length(0))
                        (d_86_v57_)[index7_] = d_22_v3_
                    d_88_v59_: D0
                    d_88_v59_ = D0_DC2((d_59_v37_).f10, (d_22_v3_).f12)
                    d_89_v60_: _dafny.Map
                    d_89_v60_ = _dafny.Map({p1: (d_59_v37_).f10})
                    d_90_v61_: D4
                    d_90_v61_ = D4_DC12((d_59_v37_).f10, d_59_v37_, True, (d_59_v37_).f10)
                    d_91_v62_: _dafny.Map
                    d_91_v62_ = _dafny.Map({d_22_v3_: d_90_v61_})
                    d_92_v63_: _dafny.Seq
                    d_92_v63_ = _dafny.SeqWithoutIsStrInference([-463])
                    d_93_v64_: _dafny.Array
                    nw11_ = _dafny.Array(None, 19)
                    nw11_[int(0)] = -181
                    nw11_[int(1)] = (d_22_v3_).f12
                    nw11_[int(2)] = 658
                    nw11_[int(3)] = (d_88_v59_).cf2
                    nw11_[int(4)] = d_20_v1_
                    nw11_[int(5)] = (d_22_v3_).f12
                    nw11_[int(6)] = default__.fm0((d_22_v3_).f12, d_89_v60_, globalState)
                    nw11_[int(7)] = (d_22_v3_).f12
                    nw11_[int(8)] = len(d_91_v62_)
                    nw11_[int(9)] = -413
                    nw11_[int(10)] = d_20_v1_
                    nw11_[int(11)] = d_20_v1_
                    nw11_[int(12)] = d_20_v1_
                    nw11_[int(13)] = (d_22_v3_).f12
                    nw11_[int(14)] = 528
                    nw11_[int(15)] = d_20_v1_
                    nw11_[int(16)] = (d_22_v3_).f12
                    nw11_[int(17)] = d_20_v1_
                    nw11_[int(18)] = len(d_92_v63_)
                    d_93_v64_ = nw11_
                    d_94_v65_: D4
                    d_94_v65_ = D4_DC10(d_93_v64_)
                    d_95_v66_: _dafny.Map
                    d_95_v66_ = _dafny.Map({17: (d_94_v65_).cf12})
                    d_95_v66_ = (d_95_v66_).set((d_20_v1_ if p0 else -648), d_93_v64_)
                    (globalState).f8 = ((d_59_v37_).f10) == ((default__.fm2((d_59_v37_).f10, p1, globalState)) and ((d_59_v37_).f10))
                    pass
            pass
        nw12_ = C0()
        nw12_.ctor__((d_22_v3_).f12, p0, (d_59_v37_).f11)
        d_59_v37_ = nw12_
        d_96_v67_: _dafny.Seq
        d_96_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qf"))
        r0 = (d_96_v67_) + (d_96_v67_)
        return r0

    @staticmethod
    def Main(noArgsParameter__):
        d_97_v0_: bool
        d_97_v0_ = False
        d_98_v1_: _dafny.Seq
        d_98_v1_ = _dafny.SeqWithoutIsStrInference([d_97_v0_, d_97_v0_])
        d_99_v2_: int
        d_99_v2_ = 147
        d_100_v3_: _dafny.MultiSet
        d_100_v3_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), d_98_v1_, (d_98_v1_).set(default__.safeIndex(-758, len(d_98_v1_)), False), (d_98_v1_).set(default__.safeIndex(d_99_v2_, len(d_98_v1_)), not(d_97_v0_))])
        d_101_v4_: D0
        d_101_v4_ = D0_DC0(d_97_v0_)
        d_102_v5_: _dafny.Array
        nw13_ = _dafny.Array(None, 11)
        nw13_[int(0)] = d_97_v0_
        nw13_[int(1)] = (d_101_v4_).cf0
        nw13_[int(2)] = d_97_v0_
        nw13_[int(3)] = d_97_v0_
        nw13_[int(4)] = d_97_v0_
        nw13_[int(5)] = d_97_v0_
        nw13_[int(6)] = d_97_v0_
        nw13_[int(7)] = d_97_v0_
        nw13_[int(8)] = d_97_v0_
        nw13_[int(9)] = d_97_v0_
        nw13_[int(10)] = d_97_v0_
        d_102_v5_ = nw13_
        d_103_v6_: _dafny.Map
        d_103_v6_ = _dafny.Map({d_97_v0_: d_99_v2_})
        d_104_v7_: D1
        d_104_v7_ = D1_DC4(d_97_v0_, d_103_v6_, d_97_v0_)
        d_105_v8_: _dafny.Map
        d_105_v8_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_97_v0_])])): d_97_v0_})
        d_106_globalState_: GlobalState
        nw14_ = GlobalState()
        nw14_.ctor__(d_100_v3_, d_102_v5_, -299, True, -378, True, (d_104_v7_).cf5, d_105_v8_, False, True)
        d_106_globalState_ = nw14_
        d_107_v9_: _dafny.MultiSet
        d_107_v9_ = _dafny.MultiSet([not(d_97_v0_), d_97_v0_, d_97_v0_])
        if not(((d_107_v9_).intersection(d_107_v9_)).issubset(d_107_v9_)):
            d_108_v10_: _dafny.Map
            d_108_v10_ = _dafny.Map({d_98_v1_: d_97_v0_})
            d_109_v11_: _dafny.Seq
            d_109_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dsjek"))
            d_99_v2_ = (default__.fm0(d_99_v2_, d_108_v10_, d_106_globalState_)) + ((len(d_109_v11_)) + (d_99_v2_))
            d_110_v12_: _dafny.Array
            def lambda4_(d_111_v2_):
                def lambda5_(d_112_i0_):
                    return default__.safeModuloInt(d_112_i0_, d_111_v2_)

                return lambda5_

            init2_ = lambda4_(d_99_v2_)
            nw15_ = _dafny.Array(None, 28)
            for i0_2_ in range(nw15_.length(0)):
                nw15_[i0_2_] = init2_(i0_2_)
            d_110_v12_ = nw15_
            d_110_v12_ = d_110_v12_
            d_113_v13_: _dafny.Seq
            out0_: _dafny.Seq
            out0_ = default__.m0(d_97_v0_, d_98_v1_, d_106_globalState_)
            d_113_v13_ = out0_
            index8_ = default__.safeIndex(136, (d_110_v12_).length(0))
            (d_110_v12_)[index8_] = d_99_v2_
            d_114_v14_: _dafny.Set
            d_114_v14_ = _dafny.Set({d_97_v0_, d_97_v0_, False, d_97_v0_, d_97_v0_})
            d_115_v15_: _dafny.Set
            d_115_v15_ = _dafny.Set({len(d_114_v14_)})
            index9_ = default__.safeIndex(136, (d_110_v12_).length(0))
            (d_110_v12_)[index9_] = (len(d_115_v15_)) + (d_99_v2_)
            d_116_v16_: _dafny.Map
            d_116_v16_ = _dafny.Map({d_97_v0_: d_97_v0_})
            d_117_v17_: _dafny.Seq
            d_117_v17_ = _dafny.SeqWithoutIsStrInference([(446 if not(((d_116_v16_)[True] if (True) in (d_116_v16_) else d_97_v0_)) else (d_110_v12_)[default__.safeIndex(136, (d_110_v12_).length(0))])])
            d_118_v18_: _dafny.Map
            d_118_v18_ = _dafny.Map({d_97_v0_: d_117_v17_})
            d_117_v17_ = ((d_117_v17_) + ((((d_118_v18_)[d_97_v0_] if (d_97_v0_) in (d_118_v18_) else d_117_v17_)).set(default__.safeIndex(d_99_v2_, len(((d_118_v18_)[d_97_v0_] if (d_97_v0_) in (d_118_v18_) else d_117_v17_))), (d_110_v12_)[default__.safeIndex(136, (d_110_v12_).length(0))]))) + (_dafny.SeqWithoutIsStrInference([(d_110_v12_)[default__.safeIndex(136, (d_110_v12_).length(0))], 584]))
        elif True:
            d_107_v9_ = (d_107_v9_) | (default__.fm1(d_99_v2_, d_97_v0_, d_99_v2_, d_97_v0_, d_106_globalState_))
            (d_106_globalState_).f8 = (D1_DC4(d_97_v0_, d_103_v6_, d_97_v0_)).cf6
            d_119_v19_: D1
            d_119_v19_ = D1_DC3(d_103_v6_)
            d_120_v20_: D1
            d_120_v20_ = D1_DC5(d_119_v19_)
            index10_ = default__.safeIndex(158, (d_102_v5_).length(0))
            (d_102_v5_)[index10_] = False
            index11_ = default__.safeIndex(158, (d_102_v5_).length(0))
            rhs26_ = d_99_v2_
            rhs27_ = d_105_v8_
            rhs28_ = d_120_v20_
            rhs29_ = d_99_v2_
            rhs30_ = d_97_v0_
            lhs11_ = d_102_v5_
            lhs12_ = default__.safeIndex(158, (d_102_v5_).length(0))
            d_99_v2_ = rhs26_
            d_105_v8_ = rhs27_
            d_120_v20_ = rhs28_
            d_99_v2_ = rhs29_
            lhs11_[lhs12_] = rhs30_
            index12_ = default__.safeIndex(158, (d_102_v5_).length(0))
            (d_102_v5_)[index12_] = default__.fm2((d_102_v5_)[default__.safeIndex(158, (d_102_v5_).length(0))], (d_98_v1_) + (d_98_v1_), d_106_globalState_)
            index13_ = default__.safeIndex(158, (d_102_v5_).length(0))
            rhs31_ = not (d_97_v0_) or (d_97_v0_)
            rhs32_ = (d_102_v5_)[default__.safeIndex(158, (d_102_v5_).length(0))]
            lhs13_ = d_106_globalState_
            lhs14_ = d_102_v5_
            lhs15_ = default__.safeIndex(158, (d_102_v5_).length(0))
            lhs13_.f8 = rhs31_
            lhs14_[lhs15_] = rhs32_
        (d_106_globalState_).f8 = d_97_v0_
        d_99_v2_ = d_99_v2_
        hi1_ = d_99_v2_
        for d_121_i1_ in range(d_99_v2_, hi1_):
            d_122_v21_: _dafny.Set
            d_122_v21_ = _dafny.Set({(0) - (d_99_v2_), d_121_i1_, d_121_i1_})
            d_123_v22_: _dafny.Seq
            d_123_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgy"))
            d_124_v23_: D0
            d_124_v23_ = D0_DC1()
            d_125_v24_: _dafny.Map
            d_125_v24_ = _dafny.Map({d_123_v22_: d_124_v23_})
            d_99_v2_ = (len(d_122_v21_) if (d_107_v9_) != (d_107_v9_) else (d_121_i1_) * (len(d_125_v24_)))
            d_126_v25_: _dafny.Map
            d_126_v25_ = _dafny.Map({d_121_i1_: d_99_v2_})
            d_107_v9_ = (_dafny.MultiSet([d_97_v0_, d_97_v0_, d_97_v0_, d_97_v0_])).set((((d_126_v25_)[d_99_v2_] if (d_99_v2_) in (d_126_v25_) else len(d_123_v22_))) <= (d_121_i1_), default__.abs(d_121_i1_))
            (d_106_globalState_).f8 = d_97_v0_
            if not (not((d_97_v0_) == (not(d_97_v0_)))) or (d_97_v0_):
                d_127_v26_: str
                d_127_v26_ = _dafny.CodePoint('f')
                d_128_v27_: _dafny.Array
                nw16_ = _dafny.Array(None, 3)
                nw16_[int(0)] = d_127_v26_
                nw16_[int(1)] = (_dafny.CodePoint('y') if not(d_97_v0_) else d_127_v26_)
                nw16_[int(2)] = d_127_v26_
                d_128_v27_ = nw16_
                d_128_v27_ = d_128_v27_
                (d_106_globalState_).f8 = d_97_v0_
                d_98_v1_ = ((d_98_v1_) + (d_98_v1_)) + (d_98_v1_)
                d_129_v28_: _dafny.Seq
                out1_: _dafny.Seq
                out1_ = default__.m0(d_97_v0_, (d_98_v1_) + (_dafny.SeqWithoutIsStrInference([not(d_97_v0_)])), d_106_globalState_)
                d_129_v28_ = out1_
                index14_ = default__.safeIndex(470, (d_102_v5_).length(0))
                (d_102_v5_)[index14_] = True
                d_130_v29_: _dafny.Array
                def lambda6_(d_131_v2_):
                    def lambda7_(d_132_i2_):
                        return default__.safeModuloInt(d_132_i2_, d_131_v2_)

                    return lambda7_

                init3_ = lambda6_(d_99_v2_)
                nw17_ = _dafny.Array(None, 2)
                for i0_3_ in range(nw17_.length(0)):
                    nw17_[i0_3_] = init3_(i0_3_)
                d_130_v29_ = nw17_
                index15_ = default__.safeIndex(391, (d_130_v29_).length(0))
                (d_130_v29_)[index15_] = (0) - (d_99_v2_)
                d_133_v30_: D2
                d_133_v30_ = D2_DC6(default__.fm3(d_106_globalState_))
                d_134_v31_: _dafny.Set
                d_134_v31_ = _dafny.Set({d_97_v0_})
                d_135_v32_: _dafny.Seq
                d_135_v32_ = _dafny.SeqWithoutIsStrInference([d_99_v2_, -223])
                index16_ = default__.safeIndex(470, (d_102_v5_).length(0))
                index17_ = default__.safeIndex(391, (d_130_v29_).length(0))
                rhs33_ = (((d_133_v30_).cf8) - (d_134_v31_)).isdisjoint((_dafny.Set({d_97_v0_, d_97_v0_}) if d_97_v0_ else d_134_v31_))
                rhs34_ = default__.safeDivisionInt(d_99_v2_, d_99_v2_)
                rhs35_ = (d_135_v32_)[default__.safeIndex((d_121_i1_) * (d_121_i1_), len(d_135_v32_))]
                rhs36_ = (0) - ((d_121_i1_) + (325))
                lhs16_ = d_102_v5_
                lhs17_ = default__.safeIndex(470, (d_102_v5_).length(0))
                lhs18_ = d_130_v29_
                lhs19_ = default__.safeIndex(391, (d_130_v29_).length(0))
                lhs16_[lhs17_] = rhs33_
                d_99_v2_ = rhs34_
                d_99_v2_ = rhs35_
                lhs18_[lhs19_] = rhs36_
            elif True:
                d_123_v22_ = d_123_v22_
                d_99_v2_ = d_121_i1_
                d_136_v33_: _dafny.Map
                d_136_v33_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([False]): d_97_v0_})
                (d_106_globalState_).f5 = default__.fm2((default__.fm0(d_121_i1_, d_136_v33_, d_106_globalState_)) != (d_99_v2_), d_98_v1_, d_106_globalState_)
                d_101_v4_ = d_101_v4_
                d_137_v34_: str
                d_137_v34_ = _dafny.CodePoint('i')
                d_137_v34_ = d_137_v34_
        d_138_v36_: _dafny.Map
        def iife6_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Seq
            for compr_2_ in (_dafny.Map({d_98_v1_: d_97_v0_})).keys.Elements:
                d_139_v35_: _dafny.Seq = compr_2_
                if (d_139_v35_) in (_dafny.Map({d_98_v1_: d_97_v0_})):
                    coll2_[d_139_v35_] = d_97_v0_
            return _dafny.Map(coll2_)
        d_138_v36_ = _dafny.Map({default__.fm0(d_99_v2_, iife6_()
        , d_106_globalState_): d_99_v2_})
        d_140_v37_: _dafny.MultiSet
        d_140_v37_ = _dafny.MultiSet([717])
        d_141_v38_: _dafny.Seq
        d_141_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "toc"))
        d_142_v39_: str
        d_142_v39_ = _dafny.CodePoint('y')
        d_143_v40_: _dafny.Array
        def lambda8_(d_144_v2_):
            def lambda9_(d_145_i3_):
                return default__.safeDivisionInt(d_145_i3_, d_144_v2_)

            return lambda9_

        init4_ = lambda8_(d_99_v2_)
        nw18_ = _dafny.Array(None, 11)
        for i0_4_ in range(nw18_.length(0)):
            nw18_[i0_4_] = init4_(i0_4_)
        d_143_v40_ = nw18_
        d_146_v41_: _dafny.Seq
        d_146_v41_ = _dafny.SeqWithoutIsStrInference([d_143_v40_])
        rhs37_ = (d_138_v36_) == ((_dafny.Map({249: ((d_140_v37_)[d_99_v2_] if (d_99_v2_) in (d_140_v37_) else len(d_141_v38_))})) | (d_138_v36_))
        rhs38_ = len(default__.fm4(d_142_v39_, (default__.fm5(d_142_v39_, d_106_globalState_)) + (d_141_v38_), d_97_v0_, d_106_globalState_))
        rhs39_ = d_97_v0_
        rhs40_ = (0) - (len(((_dafny.SeqWithoutIsStrInference([d_143_v40_])) + (d_146_v41_)) + ((d_146_v41_) + (_dafny.SeqWithoutIsStrInference([d_143_v40_, d_143_v40_, d_143_v40_, d_143_v40_])))))
        lhs20_ = d_106_globalState_
        d_97_v0_ = rhs37_
        d_99_v2_ = rhs38_
        lhs20_.f8 = rhs39_
        d_99_v2_ = rhs40_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_102_v5_).length(0)):
            d_147_i4_: int = guard_loop_1_
            if (True) and (((0) <= (d_147_i4_)) and ((d_147_i4_) < ((d_102_v5_).length(0)))):
                (d_102_v5_)[(d_147_i4_)] = (d_99_v2_) > ((d_99_v2_) * (d_99_v2_))
        d_141_v38_ = d_141_v38_
        d_99_v2_ = d_99_v2_
        d_148_v43_: _dafny.Map
        def iife7_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(728, 918):
                d_149_v42_: int = compr_3_
                if ((728) <= (d_149_v42_)) and ((d_149_v42_) < (918)):
                    coll3_[default__.safeDivisionInt(d_149_v42_, d_99_v2_)] = d_97_v0_
            return _dafny.Map(coll3_)
        d_148_v43_ = _dafny.Map({len(_dafny.Map({d_99_v2_: d_99_v2_})): iife7_()
        })
        index18_ = default__.safeIndex(797, (d_102_v5_).length(0))
        (d_102_v5_)[index18_] = (927) not in (d_148_v43_)
        d_150_v45_: _dafny.Set
        d_150_v45_ = _dafny.Set({d_98_v1_})
        index19_ = default__.safeIndex(797, (d_102_v5_).length(0))
        def iife8_():
            coll4_ = _dafny.Map()
            compr_4_: _dafny.Seq
            for compr_4_ in (d_150_v45_).Elements:
                d_151_v44_: _dafny.Seq = compr_4_
                if (d_151_v44_) in (d_150_v45_):
                    coll4_[d_151_v44_] = d_97_v0_
            return _dafny.Map(coll4_)
        rhs41_ = d_97_v0_
        rhs42_ = ((d_142_v39_ if d_97_v0_ else d_142_v39_) if True else d_142_v39_)
        rhs43_ = default__.safeModuloInt(d_99_v2_, default__.safeModuloInt(d_99_v2_, default__.fm0(d_99_v2_, iife8_()
        , d_106_globalState_)))
        rhs44_ = (601) + (d_99_v2_)
        lhs21_ = d_102_v5_
        lhs22_ = default__.safeIndex(797, (d_102_v5_).length(0))
        lhs21_[lhs22_] = rhs41_
        d_142_v39_ = rhs42_
        d_99_v2_ = rhs43_
        d_99_v2_ = rhs44_
        d_152_v46_: _dafny.Map
        d_152_v46_ = _dafny.Map({(d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))]: (d_97_v0_) == ((d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))])})
        d_153_v48_: _dafny.Seq
        d_153_v48_ = _dafny.SeqWithoutIsStrInference([d_99_v2_])
        def iife9_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in (d_153_v48_).Elements:
                d_154_v47_: int = compr_5_
                if (d_154_v47_) in (d_153_v48_):
                    coll5_[(d_154_v47_) - (d_99_v2_)] = d_140_v37_
            return _dafny.Map(coll5_)
        d_152_v46_ = (d_152_v46_).set(((d_140_v37_).cardinality) in (iife9_()
        ), (d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))])
        d_155_v49_: C0
        nw19_ = C0()
        nw19_.ctor__(d_99_v2_, d_97_v0_, d_101_v4_)
        d_155_v49_ = nw19_
        d_99_v2_ = ((d_155_v49_).f12) * (((d_155_v49_).f12) - ((d_153_v48_)[default__.safeIndex(130, len(d_153_v48_))]))
        hi2_ = d_99_v2_
        for d_156_i5_ in range((len(d_105_v8_)) + ((d_155_v49_).f12), hi2_):
            (d_106_globalState_).f8 = d_97_v0_
            d_157_v50_: C0
            nw20_ = C0()
            nw20_.ctor__((0) - (((d_155_v49_).f12) * (len(d_103_v6_))), ((d_105_v8_)[(d_155_v49_).f12] if ((d_155_v49_).f12) in (d_105_v8_) else True), d_101_v4_)
            d_157_v50_ = nw20_
            index20_ = default__.safeIndex(797, (d_102_v5_).length(0))
            (d_102_v5_)[index20_] = not(((default__.fm8(d_106_globalState_)).intersection(_dafny.MultiSet(d_153_v48_))).ispropersubset(d_140_v37_))
            d_158_v51_: _dafny.Seq
            d_158_v51_ = _dafny.SeqWithoutIsStrInference([d_157_v50_, d_157_v50_, d_157_v50_])
            if (len(d_158_v51_)) == (d_99_v2_):
                d_159_v52_: _dafny.Seq
                out2_: _dafny.Seq
                out2_ = default__.m0((d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))], d_98_v1_, d_106_globalState_)
                d_159_v52_ = out2_
                d_160_v53_: _dafny.Seq
                out3_: _dafny.Seq
                out3_ = default__.m0((d_155_v49_).fm6(d_97_v0_, d_99_v2_, True, d_106_globalState_), d_98_v1_, d_106_globalState_)
                d_160_v53_ = out3_
                (d_106_globalState_).f8 = d_97_v0_
                rhs45_ = d_97_v0_
                rhs46_ = default__.safeDivisionInt((0) - ((d_155_v49_).f12), ((d_155_v49_).f12) - (-61))
                lhs23_ = d_106_globalState_
                lhs23_.f5 = rhs45_
                d_99_v2_ = rhs46_
                index21_ = default__.safeIndex(797, (d_102_v5_).length(0))
                (d_102_v5_)[index21_] = d_97_v0_
            elif True:
                d_99_v2_ = (d_157_v50_).f12
                d_99_v2_ = (0) - (default__.safeDivisionInt((d_99_v2_) - (458), (0) - ((d_155_v49_).f12)))
                d_142_v39_ = d_142_v39_
                d_161_v54_: D2
                d_161_v54_ = D2_DC6(_dafny.Set({(d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))]}))
                d_162_v55_: _dafny.Map
                d_162_v55_ = _dafny.Map({d_161_v54_: (d_157_v50_).f12})
                d_99_v2_ = ((d_162_v55_)[d_161_v54_] if (d_161_v54_) in (d_162_v55_) else ((d_155_v49_).f12) + ((d_155_v49_).f12))
                d_163_v56_: _dafny.Seq
                out4_: _dafny.Seq
                out4_ = default__.m0((d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))], d_98_v1_, d_106_globalState_)
                d_163_v56_ = out4_
        pat_let_tv3_ = d_102_v5_
        pat_let_tv4_ = d_102_v5_
        index22_ = default__.safeIndex(797, (d_102_v5_).length(0))
        def iife10_(_pat_let2_0):
            def iife11_(d_164_dt__update__tmp_h0_):
                def iife12_(_pat_let3_0):
                    def iife13_(d_165_dt__update_hcf10_h0_):
                        return D2_DC8((d_164_dt__update__tmp_h0_).cf9, d_165_dt__update_hcf10_h0_)
                    return iife13_(_pat_let3_0)
                return iife12_((pat_let_tv4_)[default__.safeIndex(797, (pat_let_tv3_).length(0))])
            return iife11_(_pat_let2_0)
        (d_102_v5_)[index22_] = (iife10_(D2_DC8(d_97_v0_, (d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))]))).cf9
        d_166_v57_: C0
        nw21_ = C0()
        nw21_.ctor__((d_155_v49_).f12, d_97_v0_, D0_DC0((d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))]))
        d_166_v57_ = nw21_
        d_167_v58_: _dafny.Seq
        out5_: _dafny.Seq
        out5_ = default__.m0((d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))], _dafny.SeqWithoutIsStrInference([(d_102_v5_)[default__.safeIndex(797, (d_102_v5_).length(0))]]), d_106_globalState_)
        d_167_v58_ = out5_
        _dafny.print(_dafny.string_of(d_97_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_98_v1_) == (_dafny.SeqWithoutIsStrInference([False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_99_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_100_v3_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_101_v4_).cf0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_102_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_103_v6_) == (_dafny.Map({False: 147}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_104_v7_).cf5) == (_dafny.Map({False: 147}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_104_v7_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_105_v8_) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f0) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, False]), _dafny.SeqWithoutIsStrInference([False, True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_106_globalState_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f6) == (_dafny.Map({False: 147}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_106_globalState_).f7) == (_dafny.Map({1: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_106_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_106_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_107_v9_) == (_dafny.MultiSet([True, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_v36_) == (_dafny.Map({2: 147}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_140_v37_) == (_dafny.MultiSet([717]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_141_v38_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_142_v39_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_v40_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_146_v41_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v43_) == (_dafny.Map({1: _dafny.Map({-104: False, -105: False, -106: False, -107: False, -108: False, -109: False, -110: False, -111: False, -112: False, -113: False, -114: False, -115: False, -116: False, -117: False, -118: False, -119: False, -120: False, -121: False, -122: False, -123: False, -124: False, -125: False, -126: False, -127: False, -128: False, -129: False, -130: False, -131: False})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v45_) == (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, False])}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_152_v46_) == (_dafny.Map({False: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_153_v48_) == (_dafny.SeqWithoutIsStrInference([594]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_155_v49_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_166_v57_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_167_v58_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1)
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D1_DC4(False, _dafny.Map({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7()
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

class D2_DC7(D2, NamedTuple('DC7', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf9', Any), ('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC9(D3, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11)
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f5: bool = False
        self.f8: bool = False
        self._f0: _dafny.MultiSet = _dafny.MultiSet({})
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: int = int(0)
        self._f3: bool = False
        self._f4: int = int(0)
        self._f6: _dafny.Map = _dafny.Map({})
        self._f7: _dafny.Map = _dafny.Map({})
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self).f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self)._f9 = f9

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
    def f6(self):
        return self._f6
    @property
    def f7(self):
        return self._f7
    @property
    def f9(self):
        return self._f9

class C0(T0):
    def  __init__(self):
        self._f10: bool = False
        self._f11: D0 = D0.default()()
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f10(self):
        return self._f10
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f12, f10, f11):
        (self)._f12 = f12
        (self)._f10 = f10
        (self)._f11 = f11

    def fm6(self, p0, p1, p2, globalState):
        return ((self).f12) > ((0) - ((self).f12))

    def fm7(self, p0, p1, p2, globalState):
        return not (False) or (False)

    @property
    def f12(self):
        return self._f12
