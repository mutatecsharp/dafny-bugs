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
        return len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_0_i0_ in range(default__.abs(62))]) if False else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1_i1_ in range(default__.abs(584))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkumqri")))))

    @staticmethod
    def fm1(p0, p1, p2, p3, globalState):
        return not((((0) - ((0) - (-940))) + (910)) == (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrv"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffpcdlj"))))))

    @staticmethod
    def fm2(p0, p1, globalState):
        return D0_DC0(_dafny.SeqWithoutIsStrInference([561 for d_2_i0_ in range(default__.abs(-974))]))

    @staticmethod
    def fm3(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(224, 957):
                d_3_v0_: int = compr_0_
                if ((224) <= (d_3_v0_)) and ((d_3_v0_) < (957)):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeModuloInt(d_3_v0_, -930)]))
            return _dafny.Set(coll0_)
        return iife0_()
        

    @staticmethod
    def fm4(p0, globalState):
        if (not(False)) or (True):
            return _dafny.CodePoint('a')
        elif True:
            return _dafny.CodePoint('s')

    @staticmethod
    def fm5(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, True])) + (_dafny.SeqWithoutIsStrInference([False]))) + ((_dafny.SeqWithoutIsStrInference([(D2_DC7(True, -582, True, _dafny.Map({_dafny.MultiSet([not(False), True]): 370}))).cf6, False])) + (_dafny.SeqWithoutIsStrInference([True, True])))

    @staticmethod
    def fm6(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_4_i0_ in range(default__.abs(767))])

    @staticmethod
    def fm7(globalState):
        def iife1_():
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(101, 603):
                    d_5_v0_: int = compr_3_
                    if ((101) <= (d_5_v0_)) and ((d_5_v0_) < (603)):
                        coll3_[default__.safeModuloInt(d_5_v0_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, False]))})))] = (D1_DC2(False)).cf2
                return _dafny.Map(coll3_)
            coll1_ = _dafny.Set()
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(101, 603):
                    d_5_v0_: int = compr_2_
                    if ((101) <= (d_5_v0_)) and ((d_5_v0_) < (603)):
                        coll2_[default__.safeModuloInt(d_5_v0_, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False, False]))})))] = (D1_DC2(False)).cf2
                return _dafny.Map(coll2_)
            compr_1_: int
            for compr_1_ in (iife2_()
            ).keys.Elements:
                d_6_v1_: int = compr_1_
                if (d_6_v1_) in (iife3_()
                ):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_6_v1_, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([-790])), 1]))])))]))
            return _dafny.Set(coll1_)
        source0_ = D2_DC7(False, 726, True, (D2_DC7(False, len(_dafny.Map({153: iife1_()
})), False, _dafny.Map({_dafny.MultiSet([not(False)]): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcxukgk")))}))).cf9)
        if source0_.is_DC5:
            d_7___mcc_h0_ = source0_.cf4
            d_8___mcc_h1_ = source0_.cf5
            d_9_cf5_ = d_8___mcc_h1_
            d_10_cf4_ = d_7___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([859])
        elif source0_.is_DC6:
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: _dafny.MultiSet
                for compr_4_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])])).Elements:
                    d_11_v2_: _dafny.MultiSet = compr_4_
                    if (d_11_v2_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True])])):
                        coll4_ = coll4_.union(_dafny.Set([d_11_v2_]))
                return _dafny.Set(coll4_)
            return (_dafny.SeqWithoutIsStrInference([len(iife4_()
) for d_12_i0_ in range(default__.abs(989))])) + (_dafny.SeqWithoutIsStrInference([71]))
        elif source0_.is_DC7:
            d_13___mcc_h2_ = source0_.cf6
            d_14___mcc_h3_ = source0_.cf7
            d_15___mcc_h4_ = source0_.cf8
            d_16___mcc_h5_ = source0_.cf9
            d_17_cf9_ = d_16___mcc_h5_
            d_18_cf8_ = d_15___mcc_h4_
            d_19_cf7_ = d_14___mcc_h3_
            d_20_cf6_ = d_13___mcc_h2_
            return _dafny.SeqWithoutIsStrInference([d_19_cf7_])
        elif True:
            d_21___mcc_h6_ = source0_.cf3
            d_22_cf3_ = d_21___mcc_h6_
            return _dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: not(not(False))})), d_22_cf3_, d_22_cf3_, d_22_cf3_])

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return (_dafny.Set({True, True, not(True)})) - ((_dafny.Set({False, True})) - (_dafny.Set({not(False)})))

    @staticmethod
    def fm9(globalState):
        return (_dafny.MultiSet([False])) | ((_dafny.MultiSet([False, False])).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, True]))))

    @staticmethod
    def m0(globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: int = int(0)
        r1 = default__.fm0(globalState)
        d_23_v0_: int
        d_23_v0_ = 727
        d_24_v1_: _dafny.Set
        d_24_v1_ = _dafny.Set({d_23_v0_})
        d_25_v2_: bool
        d_25_v2_ = False
        d_26_v3_: _dafny.MultiSet
        d_26_v3_ = _dafny.MultiSet([d_25_v2_])
        d_27_v4_: _dafny.Seq
        d_27_v4_ = _dafny.SeqWithoutIsStrInference([d_26_v3_, d_26_v3_])
        d_28_v5_: str
        d_28_v5_ = _dafny.CodePoint('k')
        d_29_v6_: _dafny.Map
        d_29_v6_ = _dafny.Map({(d_27_v4_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_28_v5_])), len(d_27_v4_))]: (d_26_v3_).cardinality})
        d_30_v7_: D2
        d_30_v7_ = D2_DC7((d_24_v1_).ispropersubset(d_24_v1_), d_23_v0_, d_25_v2_, d_29_v6_)
        source1_ = d_30_v7_
        if source1_.is_DC5:
            d_31___mcc_h0_ = source1_.cf4
            d_32___mcc_h1_ = source1_.cf5
            d_33_cf5_ = d_32___mcc_h1_
            d_34_cf4_ = d_31___mcc_h0_
            d_35_v8_: _dafny.Array
            def lambda0_(d_36_i0_):
                return _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sldfemqf")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))})

            init0_ = lambda0_
            nw0_ = _dafny.Array(None, 17)
            for i0_0_ in range(nw0_.length(0)):
                nw0_[i0_0_] = init0_(i0_0_)
            d_35_v8_ = nw0_
            d_37_v9_: _dafny.Set
            d_37_v9_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adcvwjlwm"))})
            index0_ = default__.safeIndex(144, (d_35_v8_).length(0))
            (d_35_v8_)[index0_] = (d_37_v9_) | (d_37_v9_)
            d_38_v10_: _dafny.Seq
            d_38_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbff"))
            d_39_v11_: _dafny.Seq
            d_39_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oktgtfhvg")), d_38_v10_])
            index1_ = default__.safeIndex(144, (d_35_v8_).length(0))
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: _dafny.Seq
                for compr_5_ in (d_39_v11_).Elements:
                    d_40_v12_: _dafny.Seq = compr_5_
                    if (d_40_v12_) in (d_39_v11_):
                        coll5_ = coll5_.union(_dafny.Set([d_40_v12_]))
                return _dafny.Set(coll5_)
            (d_35_v8_)[index1_] = ((d_37_v9_) | (iife5_()
            )).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_28_v5_ for d_41_i1_ in range(default__.abs(589))]), _dafny.SeqWithoutIsStrInference([d_28_v5_ for d_42_i2_ in range(default__.abs(792))])}))
            d_43_v13_: _dafny.Map
            d_43_v13_ = _dafny.Map({(d_28_v5_ if d_25_v2_ else d_28_v5_): not(d_25_v2_)})
            (globalState).f1 = ((d_43_v13_)[default__.fm4(d_25_v2_, globalState)] if (default__.fm4(d_25_v2_, globalState)) in (d_43_v13_) else d_25_v2_)
            d_44_v14_: _dafny.Map
            d_44_v14_ = _dafny.Map({d_33_cf5_: True})
            d_45_v15_: _dafny.Map
            d_45_v15_ = _dafny.Map({(0) - (d_33_cf5_): ((d_44_v14_)[len(default__.fm5(d_33_cf5_, d_34_cf4_, globalState))] if (len(default__.fm5(d_33_cf5_, d_34_cf4_, globalState))) in (d_44_v14_) else d_25_v2_)})
            d_46_v16_: _dafny.Seq
            d_46_v16_ = _dafny.SeqWithoutIsStrInference([True, d_25_v2_, d_25_v2_])
            d_47_v17_: _dafny.Seq
            d_47_v17_ = _dafny.SeqWithoutIsStrInference([d_46_v16_, d_46_v16_, _dafny.SeqWithoutIsStrInference([d_25_v2_, True, d_25_v2_, d_25_v2_, d_25_v2_])])
            d_48_v18_: _dafny.Array
            nw1_ = _dafny.Array(None, 21)
            nw1_[int(0)] = _dafny.SeqWithoutIsStrInference([((d_45_v15_)[d_33_cf5_] if (d_33_cf5_) in (d_45_v15_) else not(True)), d_25_v2_])
            nw1_[int(1)] = d_46_v16_
            nw1_[int(2)] = (d_46_v16_) + (d_46_v16_)
            nw1_[int(3)] = _dafny.SeqWithoutIsStrInference([default__.fm1(d_23_v0_, d_25_v2_, d_25_v2_, True, globalState)])
            nw1_[int(4)] = d_46_v16_
            nw1_[int(5)] = _dafny.SeqWithoutIsStrInference([False, not(d_25_v2_)])
            nw1_[int(6)] = (d_46_v16_) + (d_46_v16_)
            nw1_[int(7)] = d_46_v16_
            nw1_[int(8)] = d_46_v16_
            nw1_[int(9)] = d_46_v16_
            nw1_[int(10)] = (d_47_v17_)[default__.safeIndex(d_23_v0_, len(d_47_v17_))]
            nw1_[int(11)] = (d_46_v16_) + (d_46_v16_)
            nw1_[int(12)] = (d_46_v16_) + (d_46_v16_)
            nw1_[int(13)] = d_46_v16_
            nw1_[int(14)] = (d_46_v16_) + (d_46_v16_)
            nw1_[int(15)] = d_46_v16_
            nw1_[int(16)] = d_46_v16_
            nw1_[int(17)] = d_46_v16_
            nw1_[int(18)] = d_46_v16_
            nw1_[int(19)] = d_46_v16_
            nw1_[int(20)] = d_46_v16_
            d_48_v18_ = nw1_
            index2_ = default__.safeIndex(703, (d_48_v18_).length(0))
            (d_48_v18_)[index2_] = d_46_v16_
            index3_ = default__.safeIndex(703, (d_48_v18_).length(0))
            (d_48_v18_)[index3_] = d_46_v16_
            rhs0_ = (d_25_v2_ if True else (d_25_v2_) or (d_25_v2_))
            rhs1_ = not(d_25_v2_)
            rhs2_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmqcbpj")))
            lhs0_ = globalState
            lhs0_.f1 = rhs0_
            d_25_v2_ = rhs1_
            r1 = rhs2_
        elif source1_.is_DC6:
            d_49_v19_: _dafny.Set
            d_49_v19_ = _dafny.Set({d_25_v2_, d_25_v2_, d_25_v2_})
            d_50_v20_: C0
            nw2_ = C0()
            nw2_.ctor__((_dafny.MultiSet([d_25_v2_, d_25_v2_, d_25_v2_])).cardinality)
            d_50_v20_ = nw2_
            d_51_v21_: _dafny.Set
            d_51_v21_ = _dafny.Set({d_50_v20_, d_50_v20_, d_50_v20_, d_50_v20_, d_50_v20_})
            d_52_v22_: _dafny.Array
            nw3_ = _dafny.Array(None, 11)
            nw3_[int(0)] = d_25_v2_
            nw3_[int(1)] = not (d_25_v2_) or (not(not(False)))
            nw3_[int(2)] = d_25_v2_
            nw3_[int(3)] = d_25_v2_
            nw3_[int(4)] = d_25_v2_
            nw3_[int(5)] = (_dafny.Set({d_25_v2_, True, d_25_v2_, False, d_25_v2_})).ispropersubset(d_49_v19_)
            nw3_[int(6)] = d_25_v2_
            nw3_[int(7)] = (d_23_v0_) != (d_23_v0_)
            nw3_[int(8)] = d_25_v2_
            nw3_[int(9)] = d_25_v2_
            nw3_[int(10)] = (d_51_v21_).ispropersubset(d_51_v21_)
            d_52_v22_ = nw3_
            index4_ = default__.safeIndex(912, (d_52_v22_).length(0))
            (d_52_v22_)[index4_] = d_25_v2_
            d_53_v23_: _dafny.Seq
            d_53_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h"))
            d_54_v24_: D2
            d_54_v24_ = D2_DC4(d_50_v20_.f5)
            pat_let_tv0_ = d_50_v20_
            pat_let_tv1_ = d_50_v20_
            index5_ = default__.safeIndex(912, (d_52_v22_).length(0))
            def iife6_(_pat_let0_0):
                def iife7_(d_55_dt__update__tmp_h0_):
                    def iife8_(_pat_let1_0):
                        def iife9_(d_56_dt__update_hcf3_h0_):
                            return D2_DC4(d_56_dt__update_hcf3_h0_)
                        return iife9_(_pat_let1_0)
                    return iife8_(pat_let_tv0_.f5)
                return iife7_(_pat_let0_0)
            def iife10_(_pat_let2_0):
                def iife11_(d_57_dt__update__tmp_h1_):
                    def iife12_(_pat_let3_0):
                        def iife13_(d_58_dt__update_hcf3_h1_):
                            return D2_DC4(d_58_dt__update_hcf3_h1_)
                        return iife13_(_pat_let3_0)
                    return iife12_(pat_let_tv1_.f5)
                return iife11_(_pat_let2_0)
            (d_52_v22_)[index5_] = ((d_53_v23_) + (d_53_v23_)) < ((default__.fm6(d_23_v0_, iife6_(d_54_v24_), globalState)).set(default__.safeIndex(len(_dafny.Set({d_25_v2_, d_25_v2_, d_25_v2_, d_25_v2_, d_25_v2_})), len(default__.fm6(d_23_v0_, iife10_(d_54_v24_), globalState))), d_28_v5_))
            d_59_v25_: _dafny.Array
            nw4_ = _dafny.Array(_dafny.Seq({}), 1)
            d_59_v25_ = nw4_
            d_60_v26_: _dafny.Array
            nw5_ = _dafny.Array(int(0), 5)
            d_60_v26_ = nw5_
            d_61_v27_: _dafny.Seq
            d_61_v27_ = _dafny.SeqWithoutIsStrInference([d_60_v26_])
            d_62_v28_: _dafny.Seq
            d_62_v28_ = d_61_v27_
            index6_ = default__.safeIndex(226, (d_59_v25_).length(0))
            (d_59_v25_)[index6_] = _dafny.SeqWithoutIsStrInference([d_62_v28_])
            d_63_v29_: D3
            d_63_v29_ = D3_DC10((d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))], (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))], d_23_v0_, d_23_v0_, d_50_v20_.f5)
            d_64_v30_: _dafny.Seq
            d_64_v30_ = _dafny.SeqWithoutIsStrInference([d_61_v27_, (d_62_v28_ if (d_63_v29_).cf13 else d_62_v28_), d_62_v28_, d_62_v28_])
            index7_ = default__.safeIndex(226, (d_59_v25_).length(0))
            (d_59_v25_)[index7_] = d_64_v30_
            d_65_v31_: _dafny.MultiSet
            d_65_v31_ = _dafny.MultiSet([d_50_v20_])
            d_66_v32_: _dafny.Seq
            d_66_v32_ = _dafny.SeqWithoutIsStrInference([(d_65_v31_).cardinality, default__.fm0(globalState)])
            d_67_v33_: _dafny.MultiSet
            d_67_v33_ = _dafny.MultiSet([d_66_v32_])
            d_67_v33_ = (d_67_v33_).intersection(d_67_v33_)
            if (D3_DC10(d_25_v2_, (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))], (d_66_v32_)[default__.safeIndex(d_23_v0_, len(d_66_v32_))], d_50_v20_.f5, d_23_v0_)).cf12:
                d_68_v34_: _dafny.Map
                d_68_v34_ = _dafny.Map({d_28_v5_: d_50_v20_.f5})
                d_69_v35_: _dafny.MultiSet
                d_69_v35_ = _dafny.MultiSet([-939, len(d_68_v34_)])
                (globalState).f1 = ((d_26_v3_).ispropersubset(_dafny.MultiSet([default__.fm1(((d_69_v35_)[d_23_v0_] if (d_23_v0_) in (d_69_v35_) else d_23_v0_), default__.fm1(d_50_v20_.f5, d_25_v2_, (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))], d_25_v2_, globalState), (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))], True, globalState)]))) == (not(d_25_v2_))
                d_70_v36_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_70_v36_ = nw6_
                d_71_v37_: _dafny.MultiSet
                d_71_v37_ = _dafny.MultiSet([d_70_v36_])
                d_72_v38_: _dafny.Map
                d_72_v38_ = _dafny.Map({default__.safeDivisionInt((d_71_v37_).cardinality, -224): (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))]})
                d_72_v38_ = (d_72_v38_).set(d_50_v20_.f5, (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))])
                d_73_v39_: C0
                nw7_ = C0()
                nw7_.ctor__(default__.safeModuloInt(d_50_v20_.f5, 844))
                d_73_v39_ = nw7_
                nw8_ = C0()
                nw8_.ctor__((d_73_v39_.f5) * (d_23_v0_))
                d_73_v39_ = nw8_
                (d_50_v20_).f5 = default__.fm0(globalState)
            elif True:
                d_66_v32_ = d_66_v32_
                (globalState).f1 = d_25_v2_
                d_74_v40_: _dafny.Array
                nw9_ = _dafny.Array(None, 27)
                nw9_[int(0)] = default__.fm4(False, globalState)
                nw9_[int(1)] = d_28_v5_
                nw9_[int(2)] = d_28_v5_
                nw9_[int(3)] = d_28_v5_
                nw9_[int(4)] = d_28_v5_
                nw9_[int(5)] = _dafny.CodePoint('a')
                nw9_[int(6)] = d_28_v5_
                nw9_[int(7)] = d_28_v5_
                nw9_[int(8)] = d_28_v5_
                nw9_[int(9)] = d_28_v5_
                nw9_[int(10)] = d_28_v5_
                nw9_[int(11)] = d_28_v5_
                nw9_[int(12)] = d_28_v5_
                nw9_[int(13)] = d_28_v5_
                nw9_[int(14)] = d_28_v5_
                nw9_[int(15)] = d_28_v5_
                nw9_[int(16)] = d_28_v5_
                nw9_[int(17)] = _dafny.CodePoint('n')
                nw9_[int(18)] = d_28_v5_
                nw9_[int(19)] = d_28_v5_
                nw9_[int(20)] = d_28_v5_
                nw9_[int(21)] = (d_28_v5_ if default__.fm1(d_50_v20_.f5, not((d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))]), d_25_v2_, (d_52_v22_)[default__.safeIndex(912, (d_52_v22_).length(0))], globalState) else _dafny.CodePoint('c'))
                nw9_[int(22)] = d_28_v5_
                nw9_[int(23)] = _dafny.CodePoint('n')
                nw9_[int(24)] = d_28_v5_
                nw9_[int(25)] = d_28_v5_
                nw9_[int(26)] = d_28_v5_
                d_74_v40_ = nw9_
                index8_ = default__.safeIndex(672, (d_74_v40_).length(0))
                (d_74_v40_)[index8_] = d_28_v5_
                d_75_v41_: _dafny.Map
                d_75_v41_ = _dafny.Map({d_53_v23_: (d_66_v32_) + (d_66_v32_)})
                index9_ = default__.safeIndex(672, (d_74_v40_).length(0))
                rhs3_ = d_50_v20_.f5
                rhs4_ = _dafny.CodePoint('v')
                rhs5_ = ((d_50_v20_.f5) * (d_50_v20_.f5)) == (d_23_v0_)
                rhs6_ = ((d_75_v41_)[(d_53_v23_) + (d_53_v23_)] if ((d_53_v23_) + (d_53_v23_)) in (d_75_v41_) else _dafny.SeqWithoutIsStrInference([d_50_v20_.f5 for d_76_i3_ in range(default__.abs(715))]))
                lhs1_ = d_50_v20_
                lhs2_ = d_74_v40_
                lhs3_ = default__.safeIndex(672, (d_74_v40_).length(0))
                lhs4_ = globalState
                lhs1_.f5 = rhs3_
                lhs2_[lhs3_] = rhs4_
                lhs4_.f1 = rhs5_
                d_66_v32_ = rhs6_
                d_77_v42_: C0
                nw10_ = C0()
                nw10_.ctor__(d_50_v20_.f5)
                d_77_v42_ = nw10_
                nw11_ = _dafny.Array(False, 16)
                d_52_v22_ = nw11_
        elif source1_.is_DC7:
            d_78___mcc_h2_ = source1_.cf6
            d_79___mcc_h3_ = source1_.cf7
            d_80___mcc_h4_ = source1_.cf8
            d_81___mcc_h5_ = source1_.cf9
            d_82_cf9_ = d_81___mcc_h5_
            d_83_cf8_ = d_80___mcc_h4_
            d_84_cf7_ = d_79___mcc_h3_
            d_85_cf6_ = d_78___mcc_h2_
            d_86_v43_: D2
            d_86_v43_ = D2_DC4(d_23_v0_)
            d_87_v44_: _dafny.Seq
            d_87_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xautmy"))
            d_88_v45_: D5
            d_88_v45_ = D5_DC12((_dafny.SeqWithoutIsStrInference([d_86_v43_, d_86_v43_, d_86_v43_, d_86_v43_, d_86_v43_])).set(default__.safeIndex(len(d_87_v44_), len(_dafny.SeqWithoutIsStrInference([d_86_v43_, d_86_v43_, d_86_v43_, d_86_v43_, d_86_v43_]))), D2_DC4(d_23_v0_)))
            d_89_v46_: _dafny.Seq
            d_89_v46_ = _dafny.SeqWithoutIsStrInference([d_86_v43_])
            d_90_v47_: _dafny.Seq
            d_90_v47_ = _dafny.SeqWithoutIsStrInference([(0) - (d_84_cf7_)])
            pat_let_tv2_ = d_86_v43_
            d_91_v48_: _dafny.Array
            nw12_ = _dafny.Array(None, 28)
            nw12_[int(0)] = (d_83_cf8_ if d_83_cf8_ else d_25_v2_)
            nw12_[int(1)] = d_85_cf6_
            nw12_[int(2)] = d_25_v2_
            nw12_[int(3)] = True
            nw12_[int(4)] = d_85_cf6_
            nw12_[int(5)] = not(d_83_cf8_)
            nw12_[int(6)] = d_25_v2_
            nw12_[int(7)] = d_25_v2_
            nw12_[int(8)] = (d_84_cf7_) == (len(_dafny.Set({d_84_cf7_})))
            nw12_[int(9)] = d_85_cf6_
            def iife14_(_pat_let4_0):
                def iife15_(d_92_dt__update__tmp_h2_):
                    def iife16_(_pat_let5_0):
                        def iife17_(d_93_dt__update_hcf18_h0_):
                            return D5_DC12(d_93_dt__update_hcf18_h0_)
                        return iife17_(_pat_let5_0)
                    return iife16_(_dafny.SeqWithoutIsStrInference([pat_let_tv2_]))
                return iife15_(_pat_let4_0)
            nw12_[int(10)] = (_dafny.MultiSet(d_89_v46_)).ispropersubset(_dafny.MultiSet((iife14_(d_88_v45_)).cf18))
            nw12_[int(11)] = (d_83_cf8_ if not(d_25_v2_) else d_25_v2_)
            nw12_[int(12)] = d_83_cf8_
            nw12_[int(13)] = d_25_v2_
            nw12_[int(14)] = not(d_85_cf6_)
            nw12_[int(15)] = (len(default__.fm7(globalState))) >= (len(d_87_v44_))
            nw12_[int(16)] = d_83_cf8_
            nw12_[int(17)] = ((0) - (d_23_v0_)) != ((0) - (d_23_v0_))
            nw12_[int(18)] = d_25_v2_
            nw12_[int(19)] = d_25_v2_
            nw12_[int(20)] = default__.fm1(len(d_90_v47_), True, False, d_83_cf8_, globalState)
            nw12_[int(21)] = (d_26_v3_).ispropersubset(d_26_v3_)
            nw12_[int(22)] = d_85_cf6_
            nw12_[int(23)] = d_83_cf8_
            nw12_[int(24)] = not((d_30_v7_).cf8)
            nw12_[int(25)] = d_83_cf8_
            nw12_[int(26)] = (d_87_v44_) != (d_87_v44_)
            nw12_[int(27)] = d_25_v2_
            d_91_v48_ = nw12_
            index10_ = default__.safeIndex(203, (d_91_v48_).length(0))
            (d_91_v48_)[index10_] = d_25_v2_
            d_94_v49_: _dafny.Map
            d_94_v49_ = _dafny.Map({d_23_v0_: default__.fm0(globalState)})
            d_95_v50_: _dafny.Array
            nw13_ = _dafny.Array(int(0), 3)
            d_95_v50_ = nw13_
            index11_ = default__.safeIndex(203, (d_91_v48_).length(0))
            rhs7_ = (d_95_v50_ if True else d_95_v50_)
            rhs8_ = d_95_v50_
            rhs9_ = (d_25_v2_ if d_83_cf8_ else d_85_cf6_)
            rhs10_ = _dafny.Map({d_84_cf7_: (d_90_v47_)[default__.safeIndex((0) - (d_23_v0_), len(d_90_v47_))]})
            rhs11_ = d_91_v48_
            lhs5_ = d_91_v48_
            lhs6_ = default__.safeIndex(203, (d_91_v48_).length(0))
            r0 = rhs7_
            r0 = rhs8_
            lhs5_[lhs6_] = rhs9_
            d_94_v49_ = rhs10_
            d_91_v48_ = rhs11_
            d_96_v51_: C0
            nw14_ = C0()
            nw14_.ctor__(d_23_v0_)
            d_96_v51_ = nw14_
            d_97_v52_: _dafny.Array
            nw15_ = _dafny.Array(None, 6)
            nw15_[int(0)] = d_96_v51_
            nw15_[int(1)] = d_96_v51_
            nw15_[int(2)] = d_96_v51_
            nw15_[int(3)] = (d_96_v51_ if default__.fm1(d_23_v0_, False, False, d_85_cf6_, globalState) else d_96_v51_)
            nw15_[int(4)] = d_96_v51_
            nw15_[int(5)] = d_96_v51_
            d_97_v52_ = nw15_
            index12_ = default__.safeIndex(831, (d_97_v52_).length(0))
            (d_97_v52_)[index12_] = d_96_v51_
            d_98_v53_: _dafny.Map
            d_98_v53_ = _dafny.Map({d_84_cf7_: (_dafny.SeqWithoutIsStrInference([d_28_v5_ for d_99_i4_ in range(default__.abs(216))]) if (d_91_v48_)[default__.safeIndex(203, (d_91_v48_).length(0))] else d_87_v44_)})
            index13_ = default__.safeIndex(831, (d_97_v52_).length(0))
            rhs12_ = d_96_v51_
            rhs13_ = (d_98_v53_) | (d_98_v53_)
            lhs7_ = d_97_v52_
            lhs8_ = default__.safeIndex(831, (d_97_v52_).length(0))
            lhs7_[lhs8_] = rhs12_
            d_98_v53_ = rhs13_
            if d_25_v2_:
                d_94_v49_ = (d_94_v49_).set((0) - (d_96_v51_.f5), default__.fm0(globalState))
                d_100_v54_: _dafny.MultiSet
                d_100_v54_ = _dafny.MultiSet([d_95_v50_])
                d_100_v54_ = d_100_v54_
                d_101_v55_: _dafny.Set
                d_101_v55_ = _dafny.Set({not(not(False))})
                d_102_v56_: _dafny.Array
                nw16_ = _dafny.Array(None, 14)
                nw16_[int(0)] = d_95_v50_
                nw16_[int(1)] = d_95_v50_
                nw16_[int(2)] = d_95_v50_
                nw16_[int(3)] = d_95_v50_
                nw16_[int(4)] = d_95_v50_
                nw16_[int(5)] = d_95_v50_
                nw16_[int(6)] = d_95_v50_
                nw16_[int(7)] = d_95_v50_
                nw16_[int(8)] = d_95_v50_
                nw16_[int(9)] = d_95_v50_
                nw16_[int(10)] = d_95_v50_
                nw16_[int(11)] = d_95_v50_
                nw16_[int(12)] = d_95_v50_
                nw16_[int(13)] = d_95_v50_
                d_102_v56_ = nw16_
                index14_ = default__.safeIndex(765, (d_102_v56_).length(0))
                (d_102_v56_)[index14_] = d_95_v50_
                d_103_v57_: _dafny.Array
                def lambda1_(d_104_v47_):
                    def lambda2_(d_105_i5_):
                        return d_104_v47_

                    return lambda2_

                init1_ = lambda1_(d_90_v47_)
                nw17_ = _dafny.Array(None, 21)
                for i0_1_ in range(nw17_.length(0)):
                    nw17_[i0_1_] = init1_(i0_1_)
                d_103_v57_ = nw17_
                index15_ = default__.safeIndex(719, (d_103_v57_).length(0))
                (d_103_v57_)[index15_] = _dafny.SeqWithoutIsStrInference([len(d_87_v44_), d_84_cf7_])
                d_106_v58_: _dafny.Seq
                d_106_v58_ = _dafny.SeqWithoutIsStrInference([(d_91_v48_)[default__.safeIndex(203, (d_91_v48_).length(0))]])
                index16_ = default__.safeIndex(765, (d_102_v56_).length(0))
                index17_ = default__.safeIndex(719, (d_103_v57_).length(0))
                rhs14_ = default__.fm8(d_85_cf6_, d_83_cf8_, (d_106_v58_) + (d_106_v58_), (d_101_v55_) == (d_101_v55_), globalState)
                rhs15_ = d_95_v50_
                rhs16_ = default__.fm7(globalState)
                lhs9_ = d_102_v56_
                lhs10_ = default__.safeIndex(765, (d_102_v56_).length(0))
                lhs11_ = d_103_v57_
                lhs12_ = default__.safeIndex(719, (d_103_v57_).length(0))
                d_101_v55_ = rhs14_
                lhs9_[lhs10_] = rhs15_
                lhs11_[lhs12_] = rhs16_
                d_107_v59_: _dafny.MultiSet
                d_107_v59_ = _dafny.MultiSet([d_26_v3_, d_26_v3_])
                d_107_v59_ = (_dafny.MultiSet((_dafny.SeqWithoutIsStrInference([d_26_v3_ for d_108_i6_ in range(default__.abs(422))])) + (_dafny.SeqWithoutIsStrInference([d_26_v3_, _dafny.MultiSet(d_106_v58_), default__.fm9(globalState), d_26_v3_])))).intersection(d_107_v59_)
                index18_ = default__.safeIndex(165, (d_95_v50_).length(0))
                (d_95_v50_)[index18_] = d_96_v51_.f5
                index19_ = default__.safeIndex(165, (d_95_v50_).length(0))
                (d_95_v50_)[index19_] = d_84_cf7_
            elif True:
                d_91_v48_ = d_91_v48_
                rhs17_ = default__.safeDivisionInt(d_23_v0_, (0) - (d_84_cf7_))
                rhs18_ = d_85_cf6_
                lhs13_ = d_96_v51_
                lhs14_ = globalState
                lhs13_.f5 = rhs17_
                lhs14_.f1 = rhs18_
                d_83_cf8_ = d_25_v2_
                d_28_v5_ = d_28_v5_
                d_109_v60_: _dafny.Array
                nw18_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
                d_109_v60_ = nw18_
                index20_ = default__.safeIndex(452, (d_109_v60_).length(0))
                (d_109_v60_)[index20_] = d_87_v44_
                index21_ = default__.safeIndex(452, (d_109_v60_).length(0))
                rhs19_ = d_109_v60_
                rhs20_ = ((d_87_v44_) + (_dafny.SeqWithoutIsStrInference([(d_87_v44_)[default__.safeIndex(d_96_v51_.f5, len(d_87_v44_))] for d_110_i7_ in range(default__.abs(621))]))) + (d_87_v44_)
                lhs15_ = d_109_v60_
                lhs16_ = default__.safeIndex(452, (d_109_v60_).length(0))
                d_109_v60_ = rhs19_
                lhs15_[lhs16_] = rhs20_
            d_111_v61_: _dafny.Seq
            d_111_v61_ = _dafny.SeqWithoutIsStrInference([d_95_v50_])
            d_112_v62_: _dafny.MultiSet
            d_112_v62_ = _dafny.MultiSet([-259])
            d_113_v63_: _dafny.Seq
            d_113_v63_ = _dafny.SeqWithoutIsStrInference([d_95_v50_])
            d_111_v61_ = (d_113_v63_ if (d_112_v62_).ispropersubset(d_112_v62_) else d_113_v63_)
        elif True:
            d_114___mcc_h6_ = source1_.cf3
            d_115_cf3_ = d_114___mcc_h6_
            d_116_v64_: _dafny.Array
            nw19_ = _dafny.Array(False, 22)
            d_116_v64_ = nw19_
            index22_ = default__.safeIndex(910, (d_116_v64_).length(0))
            (d_116_v64_)[index22_] = d_25_v2_
            d_117_v65_: _dafny.Map
            d_117_v65_ = _dafny.Map({default__.safeModuloInt(d_23_v0_, d_23_v0_): True})
            d_118_v66_: D2
            d_118_v66_ = D2_DC4(d_23_v0_)
            d_119_v67_: _dafny.Seq
            d_119_v67_ = _dafny.SeqWithoutIsStrInference([d_25_v2_])
            index23_ = default__.safeIndex(910, (d_116_v64_).length(0))
            rhs21_ = d_25_v2_
            rhs22_ = default__.fm1(d_115_cf3_, (d_25_v2_ if not(default__.fm1(d_115_cf3_, d_25_v2_, True, d_25_v2_, globalState)) else (d_119_v67_)[default__.safeIndex(d_115_cf3_, len(d_119_v67_))]), (d_115_cf3_) == (d_23_v0_), not(d_25_v2_), globalState)
            rhs23_ = (d_117_v65_).set(d_23_v0_, d_25_v2_)
            rhs24_ = D2_DC4(d_23_v0_)
            rhs25_ = (0) - (default__.safeModuloInt(d_115_cf3_, d_115_cf3_))
            lhs17_ = d_116_v64_
            lhs18_ = default__.safeIndex(910, (d_116_v64_).length(0))
            lhs19_ = globalState
            lhs17_[lhs18_] = rhs21_
            lhs19_.f1 = rhs22_
            d_117_v65_ = rhs23_
            d_118_v66_ = rhs24_
            d_115_cf3_ = rhs25_
            (globalState).f1 = (d_116_v64_)[default__.safeIndex(910, (d_116_v64_).length(0))]
            d_28_v5_ = d_28_v5_
            d_120_v68_: _dafny.Map
            d_120_v68_ = _dafny.Map({d_28_v5_: 382})
            d_121_v69_: C0
            nw20_ = C0()
            nw20_.ctor__(d_23_v0_)
            d_121_v69_ = nw20_
            d_122_v70_: _dafny.Map
            d_122_v70_ = _dafny.Map({(d_116_v64_)[default__.safeIndex(910, (d_116_v64_).length(0))]: (d_116_v64_)[default__.safeIndex(910, (d_116_v64_).length(0))]})
            d_123_v71_: _dafny.MultiSet
            d_123_v71_ = _dafny.MultiSet([d_28_v5_])
            d_124_v72_: _dafny.Array
            nw21_ = _dafny.Array(None, 13)
            nw21_[int(0)] = d_115_cf3_
            nw21_[int(1)] = len(d_122_v70_)
            nw21_[int(2)] = (d_123_v71_).cardinality
            nw21_[int(3)] = d_121_v69_.f5
            nw21_[int(4)] = 679
            nw21_[int(5)] = (0) - (d_115_cf3_)
            nw21_[int(6)] = d_121_v69_.f5
            nw21_[int(7)] = d_23_v0_
            nw21_[int(8)] = 821
            nw21_[int(9)] = d_121_v69_.f5
            nw21_[int(10)] = d_23_v0_
            nw21_[int(11)] = d_23_v0_
            nw21_[int(12)] = d_121_v69_.f5
            d_124_v72_ = nw21_
            d_125_v73_: _dafny.Seq
            d_125_v73_ = _dafny.SeqWithoutIsStrInference([d_124_v72_, d_124_v72_])
            d_126_v74_: _dafny.Seq
            d_126_v74_ = d_125_v73_
            d_127_v75_: _dafny.Map
            d_127_v75_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_121_v69_])): d_126_v74_})
            d_120_v68_ = (d_120_v68_).set(d_28_v5_, len(d_127_v75_))
        d_128_v76_: C0
        nw22_ = C0()
        nw22_.ctor__((d_23_v0_) * (d_23_v0_))
        d_128_v76_ = nw22_
        d_129_v77_: _dafny.MultiSet
        d_129_v77_ = _dafny.MultiSet([(d_128_v76_.f5) * (d_128_v76_.f5), d_23_v0_])
        d_129_v77_ = (d_129_v77_).intersection(d_129_v77_)
        pat_let_tv3_ = d_29_v6_
        pat_let_tv4_ = d_23_v0_
        def iife18_(_pat_let6_0):
            def iife19_(d_130_dt__update__tmp_h4_):
                def iife20_(_pat_let7_0):
                    def iife21_(d_131_dt__update_hcf9_h0_):
                        def iife22_(_pat_let8_0):
                            def iife23_(d_132_dt__update_hcf7_h0_):
                                return D2_DC7((d_130_dt__update__tmp_h4_).cf6, d_132_dt__update_hcf7_h0_, (d_130_dt__update__tmp_h4_).cf8, d_131_dt__update_hcf9_h0_)
                            return iife23_(_pat_let8_0)
                        return iife22_(pat_let_tv4_)
                    return iife21_(_pat_let7_0)
                return iife20_(pat_let_tv3_)
            return iife19_(_pat_let6_0)
        d_30_v7_ = iife18_(d_30_v7_)
        d_133_v78_: D3
        d_133_v78_ = D3_DC10(d_25_v2_, d_25_v2_, d_23_v0_, d_128_v76_.f5, d_23_v0_)
        d_134_v79_: _dafny.Set
        d_134_v79_ = _dafny.Set({d_25_v2_})
        d_135_v80_: _dafny.Seq
        d_135_v80_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_128_v76_.f5, d_25_v2_, d_25_v2_, d_25_v2_, globalState), d_25_v2_, False])
        d_136_v81_: D3
        d_136_v81_ = D3_DC9(d_129_v77_)
        d_137_v82_: _dafny.MultiSet
        d_137_v82_ = _dafny.MultiSet([d_136_v81_, d_136_v81_])
        d_138_v83_: _dafny.Set
        d_138_v83_ = _dafny.Set({d_24_v1_})
        d_139_v84_: _dafny.Array
        nw23_ = _dafny.Array(None, 20)
        nw23_[int(0)] = (d_133_v78_).cf12
        nw23_[int(1)] = d_25_v2_
        nw23_[int(2)] = d_25_v2_
        nw23_[int(3)] = (d_25_v2_ if not(not(d_25_v2_)) else d_25_v2_)
        nw23_[int(4)] = (d_128_v76_.f5) >= ((0) - (len(d_134_v79_)))
        nw23_[int(5)] = (d_135_v80_)[default__.safeIndex(d_128_v76_.f5, len(d_135_v80_))]
        nw23_[int(6)] = d_25_v2_
        nw23_[int(7)] = d_25_v2_
        nw23_[int(8)] = d_25_v2_
        nw23_[int(9)] = d_25_v2_
        nw23_[int(10)] = d_25_v2_
        nw23_[int(11)] = d_25_v2_
        nw23_[int(12)] = d_25_v2_
        nw23_[int(13)] = d_25_v2_
        nw23_[int(14)] = True
        nw23_[int(15)] = d_25_v2_
        nw23_[int(16)] = d_25_v2_
        nw23_[int(17)] = d_25_v2_
        nw23_[int(18)] = (d_137_v82_) != (_dafny.MultiSet([d_136_v81_]))
        nw23_[int(19)] = (d_138_v83_).issubset(d_138_v83_)
        d_139_v84_ = nw23_
        index24_ = default__.safeIndex(382, (d_139_v84_).length(0))
        (d_139_v84_)[index24_] = d_25_v2_
        index25_ = default__.safeIndex(382, (d_139_v84_).length(0))
        (d_139_v84_)[index25_] = (d_133_v78_).cf12
        nw24_ = _dafny.Array(int(0), 5)
        r0 = nw24_
        d_140_v85_: _dafny.Seq
        d_140_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "matixu"))
        d_141_v86_: _dafny.Map
        d_141_v86_ = _dafny.Map({d_140_v85_: _dafny.SeqWithoutIsStrInference([d_28_v5_])})
        d_142_v87_: _dafny.Seq
        d_142_v87_ = _dafny.SeqWithoutIsStrInference([len((((d_141_v86_)[d_140_v85_] if (d_140_v85_) in (d_141_v86_) else d_140_v85_)) + (_dafny.SeqWithoutIsStrInference([d_28_v5_ for d_143_i8_ in range(default__.abs(814))]))), d_128_v76_.f5, (d_23_v0_) * (d_128_v76_.f5)])
        r1 = (d_142_v87_)[default__.safeIndex(d_23_v0_, len(d_142_v87_))]
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_144_v0_: int
        d_144_v0_ = -560
        d_145_v1_: _dafny.Seq
        d_145_v1_ = _dafny.SeqWithoutIsStrInference([d_144_v0_, d_144_v0_])
        d_146_v2_: _dafny.Array
        def lambda3_(d_147_i0_):
            return _dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')])

        init2_ = lambda3_
        nw25_ = _dafny.Array(None, 19)
        for i0_2_ in range(nw25_.length(0)):
            nw25_[i0_2_] = init2_(i0_2_)
        d_146_v2_ = nw25_
        d_148_globalState_: GlobalState
        nw26_ = GlobalState()
        nw26_.ctor__(791, False, d_145_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpno")), d_146_v2_)
        d_148_globalState_ = nw26_
        d_149_v3_: bool
        d_149_v3_ = True
        d_150_i1_: int
        d_150_i1_ = 0
        with _dafny.label("0"):
            while d_149_v3_:
                with _dafny.c_label("0"):
                    if (d_150_i1_) >= (100):
                        raise _dafny.Break("0")
                    d_150_i1_ = (d_150_i1_) + (1)
                    d_151_v4_: _dafny.Array
                    nw27_ = _dafny.Array(False, 19)
                    d_151_v4_ = nw27_
                    index26_ = default__.safeIndex(258, (d_151_v4_).length(0))
                    (d_151_v4_)[index26_] = False
                    index27_ = default__.safeIndex(258, (d_151_v4_).length(0))
                    (d_151_v4_)[index27_] = not (d_149_v3_) or (d_149_v3_)
                    if not((d_151_v4_)[default__.safeIndex(258, (d_151_v4_).length(0))]):
                        d_152_v5_: _dafny.Array
                        d_153_v6_: int
                        out0_: _dafny.Array
                        out1_: int
                        out0_, out1_ = default__.m0(d_148_globalState_)
                        d_152_v5_ = out0_
                        d_153_v6_ = out1_
                        (d_148_globalState_).f1 = not (d_149_v3_) or (d_149_v3_)
                        d_145_v1_ = ((D0_DC0(d_145_v1_)).cf0).set(default__.safeIndex(default__.fm0(d_148_globalState_), len((D0_DC0(d_145_v1_)).cf0)), d_144_v0_)
                        (d_148_globalState_).f1 = (d_151_v4_)[default__.safeIndex(258, (d_151_v4_).length(0))]
                        d_144_v0_ = default__.safeModuloInt(default__.fm0(d_148_globalState_), default__.safeDivisionInt(d_153_v6_, d_144_v0_))
                    elif True:
                        d_154_v7_: _dafny.Array
                        d_155_v8_: int
                        out2_: _dafny.Array
                        out3_: int
                        out2_, out3_ = default__.m0(d_148_globalState_)
                        d_154_v7_ = out2_
                        d_155_v8_ = out3_
                        d_149_v3_ = (d_151_v4_)[default__.safeIndex(258, (d_151_v4_).length(0))]
                        d_156_v9_: D1
                        d_156_v9_ = D1_DC2((d_151_v4_)[default__.safeIndex(258, (d_151_v4_).length(0))])
                        index28_ = default__.safeIndex(258, (d_151_v4_).length(0))
                        (d_151_v4_)[index28_] = (d_156_v9_).cf2
                        d_157_v10_: C0
                        nw28_ = C0()
                        nw28_.ctor__(d_144_v0_)
                        d_157_v10_ = nw28_
                        rhs26_ = d_155_v8_
                        rhs27_ = default__.safeModuloInt(90, d_144_v0_)
                        lhs20_ = d_157_v10_
                        d_144_v0_ = rhs26_
                        lhs20_.f5 = rhs27_
                    (d_148_globalState_).f1 = True
                    d_158_v11_: _dafny.MultiSet
                    d_158_v11_ = _dafny.MultiSet([d_149_v3_, (d_151_v4_)[default__.safeIndex(258, (d_151_v4_).length(0))]])
                    d_159_v12_: D0
                    d_159_v12_ = D0_DC0((d_145_v1_ if d_149_v3_ else _dafny.SeqWithoutIsStrInference([(d_158_v11_).cardinality])))
                    rhs28_ = (612 if d_149_v3_ else d_144_v0_)
                    rhs29_ = d_159_v12_
                    d_144_v0_ = rhs28_
                    d_159_v12_ = rhs29_
                    pass
            pass
        d_160_v13_: _dafny.Array
        nw29_ = _dafny.Array(int(0), 5)
        d_160_v13_ = nw29_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_160_v13_).length(0)):
            d_161_i2_: int = guard_loop_0_
            if (True) and (((0) <= (d_161_i2_)) and ((d_161_i2_) < ((d_160_v13_).length(0)))):
                (d_160_v13_)[(d_161_i2_)] = default__.safeModuloInt(d_161_i2_, (D2_DC4(d_144_v0_)).cf3)
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_160_v13_).length(0)):
            d_162_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_162_i3_)) and ((d_162_i3_) < ((d_160_v13_).length(0)))):
                (d_160_v13_)[(d_162_i3_)] = (d_162_i3_) * (d_144_v0_)
        (d_148_globalState_).f3 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_163_i4_ in range(default__.abs(-708))])
        d_164_v14_: C0
        nw30_ = C0()
        nw30_.ctor__(190)
        d_164_v14_ = nw30_
        d_165_v15_: _dafny.Seq
        d_165_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqlvtuvhq"))
        d_166_i5_: int
        d_166_i5_ = 0
        with _dafny.label("1"):
            while (d_149_v3_ if (d_144_v0_) >= (d_144_v0_) else (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_170_i6_ in range(default__.abs(671))])) != (d_165_v15_)):
                with _dafny.c_label("1"):
                    if (d_166_i5_) >= (100):
                        raise _dafny.Break("1")
                    d_166_i5_ = (d_166_i5_) + (1)
                    (d_164_v14_).f5 = len(d_165_v15_)
                    d_167_v16_: _dafny.Map
                    d_167_v16_ = _dafny.Map({d_149_v3_: d_164_v14_.f5})
                    d_168_v17_: _dafny.Map
                    d_168_v17_ = _dafny.Map({d_144_v0_: d_167_v16_})
                    d_169_v18_: _dafny.Map
                    d_169_v18_ = _dafny.Map({d_168_v17_: d_164_v14_.f5})
                    d_169_v18_ = (d_169_v18_).set(d_168_v17_, d_164_v14_.f5)
                    d_145_v1_ = d_145_v1_
                    d_144_v0_ = d_164_v14_.f5
                    pass
            pass
        d_171_v19_: str
        d_171_v19_ = _dafny.CodePoint('q')
        d_171_v19_ = (d_165_v15_)[default__.safeIndex(764, len(d_165_v15_))]
        d_172_v20_: _dafny.Array
        def lambda4_(d_173_v0_, d_174_v14_):
            def lambda5_(d_175_i7_):
                return _dafny.Map({d_173_v0_: d_174_v14_.f5})

            return lambda5_

        init3_ = lambda4_(d_144_v0_, d_164_v14_)
        nw31_ = _dafny.Array(None, 17)
        for i0_3_ in range(nw31_.length(0)):
            nw31_[i0_3_] = init3_(i0_3_)
        d_172_v20_ = nw31_
        d_176_v21_: _dafny.Map
        d_176_v21_ = _dafny.Map({d_149_v3_: d_144_v0_})
        d_177_v22_: _dafny.Map
        d_177_v22_ = _dafny.Map({d_164_v14_.f5: len(d_176_v21_)})
        index29_ = default__.safeIndex(590, (d_172_v20_).length(0))
        (d_172_v20_)[index29_] = d_177_v22_
        index30_ = default__.safeIndex(590, (d_172_v20_).length(0))
        (d_172_v20_)[index30_] = ((d_177_v22_) | (_dafny.Map({(0) - (d_164_v14_.f5): d_164_v14_.f5}))).set(d_144_v0_, d_164_v14_.f5)
        hi0_ = d_144_v0_
        for d_178_i8_ in range(232, hi0_):
            d_179_v23_: C0
            nw32_ = C0()
            nw32_.ctor__(d_144_v0_)
            d_179_v23_ = nw32_
            d_180_v24_: _dafny.Array
            d_181_v25_: int
            out4_: _dafny.Array
            out5_: int
            out4_, out5_ = default__.m0(d_148_globalState_)
            d_180_v24_ = out4_
            d_181_v25_ = out5_
            index31_ = default__.safeIndex(212, (d_180_v24_).length(0))
            (d_180_v24_)[index31_] = d_178_i8_
            d_182_v26_: _dafny.Array
            nw33_ = _dafny.Array(None, 13)
            nw33_[int(0)] = not(d_149_v3_)
            nw33_[int(1)] = False
            nw33_[int(2)] = not(d_149_v3_)
            nw33_[int(3)] = d_149_v3_
            nw33_[int(4)] = d_149_v3_
            nw33_[int(5)] = d_149_v3_
            nw33_[int(6)] = d_149_v3_
            nw33_[int(7)] = d_149_v3_
            nw33_[int(8)] = d_149_v3_
            nw33_[int(9)] = d_149_v3_
            nw33_[int(10)] = d_149_v3_
            nw33_[int(11)] = d_149_v3_
            nw33_[int(12)] = d_149_v3_
            d_182_v26_ = nw33_
            d_183_v27_: _dafny.Map
            d_183_v27_ = _dafny.Map({d_182_v26_: default__.fm1(d_181_v25_, False, d_149_v3_, d_149_v3_, d_148_globalState_)})
            index32_ = default__.safeIndex(212, (d_180_v24_).length(0))
            rhs30_ = (0) - (default__.fm0(d_148_globalState_))
            rhs31_ = d_183_v27_
            rhs32_ = (0) - (d_164_v14_.f5)
            lhs21_ = d_180_v24_
            lhs22_ = default__.safeIndex(212, (d_180_v24_).length(0))
            lhs21_[lhs22_] = rhs30_
            d_183_v27_ = rhs31_
            d_181_v25_ = rhs32_
            source2_ = default__.fm2(d_149_v3_, 321, d_148_globalState_)
            if source2_.is_DC1:
                d_184___mcc_h0_ = source2_.cf1
                d_185_cf1_ = d_184___mcc_h0_
                index33_ = default__.safeIndex(212, (d_180_v24_).length(0))
                (d_180_v24_)[index33_] = (d_164_v14_.f5) - ((d_164_v14_.f5) * (default__.fm0(d_148_globalState_)))
                d_186_v28_: C0
                nw34_ = C0()
                nw34_.ctor__(d_144_v0_)
                d_186_v28_ = nw34_
                d_187_v29_: D2
                d_187_v29_ = D2_DC5(d_164_v14_.f5, d_178_i8_)
                pat_let_tv5_ = d_144_v0_
                d_188_v30_: _dafny.Map
                def iife24_(_pat_let9_0):
                    def iife25_(d_189_dt__update__tmp_h0_):
                        def iife26_(_pat_let10_0):
                            def iife27_(d_190_dt__update_hcf5_h0_):
                                return D2_DC5((d_189_dt__update__tmp_h0_).cf4, d_190_dt__update_hcf5_h0_)
                            return iife27_(_pat_let10_0)
                        return iife26_(pat_let_tv5_)
                    return iife25_(_pat_let9_0)
                d_188_v30_ = _dafny.Map({d_165_v15_: iife24_(d_187_v29_)})
                d_188_v30_ = d_188_v30_
                d_191_v31_: _dafny.Seq
                d_191_v31_ = _dafny.SeqWithoutIsStrInference([d_182_v26_])
                d_192_v32_: _dafny.Seq
                d_192_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({not(d_149_v3_): _dafny.SeqWithoutIsStrInference([d_149_v3_])})])
                d_193_v33_: _dafny.Array
                nw35_ = _dafny.Array(None, 25)
                nw35_[int(0)] = d_182_v26_
                nw35_[int(1)] = d_182_v26_
                nw35_[int(2)] = d_182_v26_
                nw35_[int(3)] = d_182_v26_
                nw35_[int(4)] = d_182_v26_
                nw35_[int(5)] = d_182_v26_
                nw35_[int(6)] = d_182_v26_
                nw35_[int(7)] = d_182_v26_
                nw35_[int(8)] = d_182_v26_
                nw35_[int(9)] = d_182_v26_
                nw35_[int(10)] = d_182_v26_
                nw35_[int(11)] = d_182_v26_
                nw35_[int(12)] = d_182_v26_
                nw35_[int(13)] = (d_191_v31_)[default__.safeIndex(len((d_192_v32_)[default__.safeIndex(d_179_v23_.f5, len(d_192_v32_))]), len(d_191_v31_))]
                nw35_[int(14)] = d_182_v26_
                nw35_[int(15)] = d_182_v26_
                nw35_[int(16)] = d_182_v26_
                nw35_[int(17)] = d_182_v26_
                nw35_[int(18)] = d_182_v26_
                nw35_[int(19)] = d_182_v26_
                nw35_[int(20)] = d_182_v26_
                nw35_[int(21)] = d_182_v26_
                nw35_[int(22)] = d_182_v26_
                nw35_[int(23)] = d_182_v26_
                nw35_[int(24)] = d_182_v26_
                d_193_v33_ = nw35_
                d_194_v34_: _dafny.Map
                d_194_v34_ = _dafny.Map({d_164_v14_: d_182_v26_})
                index34_ = default__.safeIndex(372, (d_193_v33_).length(0))
                (d_193_v33_)[index34_] = ((d_194_v34_)[d_179_v23_] if (d_179_v23_) in (d_194_v34_) else d_182_v26_)
                d_195_v35_: _dafny.Map
                d_195_v35_ = _dafny.Map({len((d_172_v20_)[default__.safeIndex(590, (d_172_v20_).length(0))]): d_149_v3_})
                d_196_v36_: _dafny.MultiSet
                d_196_v36_ = _dafny.MultiSet([d_195_v35_])
                d_197_v37_: _dafny.MultiSet
                d_197_v37_ = _dafny.MultiSet([(d_196_v36_).cardinality, d_178_i8_, (0) - ((d_180_v24_)[default__.safeIndex(212, (d_180_v24_).length(0))]), d_186_v28_.f5, d_181_v25_])
                d_198_v38_: _dafny.Seq
                d_198_v38_ = _dafny.SeqWithoutIsStrInference([d_186_v28_])
                index35_ = default__.safeIndex(372, (d_193_v33_).length(0))
                rhs33_ = ((d_197_v37_).set(len(d_198_v38_), default__.abs(len(d_165_v15_)))).isdisjoint(d_197_v37_)
                rhs34_ = d_182_v26_
                lhs23_ = d_148_globalState_
                lhs24_ = d_193_v33_
                lhs25_ = default__.safeIndex(372, (d_193_v33_).length(0))
                lhs23_.f1 = rhs33_
                lhs24_[lhs25_] = rhs34_
            elif True:
                d_199___mcc_h1_ = source2_.cf0
                d_200_cf0_ = d_199___mcc_h1_
                d_201_v39_: C0
                nw36_ = C0()
                nw36_.ctor__((d_180_v24_)[default__.safeIndex(212, (d_180_v24_).length(0))])
                d_201_v39_ = nw36_
                d_202_v40_: _dafny.MultiSet
                d_202_v40_ = _dafny.MultiSet([(d_164_v14_.f5) >= (d_179_v23_.f5)])
                d_203_v41_: _dafny.Seq
                d_203_v41_ = _dafny.SeqWithoutIsStrInference([d_149_v3_, d_149_v3_])
                d_202_v40_ = _dafny.MultiSet([(d_149_v3_ if d_149_v3_ else False), (d_203_v41_)[default__.safeIndex(default__.fm0(d_148_globalState_), len(d_203_v41_))], (d_202_v40_).ispropersubset(d_202_v40_)])
                d_204_v42_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_204_v42_ = nw37_
                d_205_v43_: D0
                d_205_v43_ = D0_DC1(d_204_v42_)
                d_206_v44_: _dafny.Map
                d_206_v44_ = _dafny.Map({d_205_v43_: False})
                d_207_v45_: D3
                d_207_v45_ = D3_DC8(d_206_v44_)
                d_208_v46_: _dafny.Seq
                d_208_v46_ = _dafny.SeqWithoutIsStrInference([(d_206_v44_) | ((d_207_v45_).cf10), d_206_v44_, d_206_v44_, (d_206_v44_).set(D0_DC1(d_204_v42_), d_149_v3_)])
                d_181_v25_ = len(d_208_v46_)
                (d_201_v39_).f5 = (0) - ((0) - (d_181_v25_))
        if d_149_v3_:
            d_209_v47_: _dafny.Array
            def lambda6_(d_210_i9_):
                return False

            init4_ = lambda6_
            nw38_ = _dafny.Array(None, 25)
            for i0_4_ in range(nw38_.length(0)):
                nw38_[i0_4_] = init4_(i0_4_)
            d_209_v47_ = nw38_
            index36_ = default__.safeIndex(360, (d_209_v47_).length(0))
            (d_209_v47_)[index36_] = d_149_v3_
            index37_ = default__.safeIndex(360, (d_209_v47_).length(0))
            (d_209_v47_)[index37_] = not(d_149_v3_)
            (d_164_v14_).f5 = d_144_v0_
            index38_ = default__.safeIndex(587, (d_160_v13_).length(0))
            (d_160_v13_)[index38_] = (((d_172_v20_)[default__.safeIndex(590, (d_172_v20_).length(0))])[(0) - (d_164_v14_.f5)] if ((0) - (d_164_v14_.f5)) in ((d_172_v20_)[default__.safeIndex(590, (d_172_v20_).length(0))]) else d_144_v0_)
            d_211_v48_: _dafny.Seq
            d_211_v48_ = _dafny.SeqWithoutIsStrInference([d_160_v13_, d_160_v13_, d_160_v13_, d_160_v13_])
            d_212_v50_: _dafny.Set
            d_212_v50_ = _dafny.Set({len(d_145_v1_), d_144_v0_})
            d_213_v51_: _dafny.Map
            d_213_v51_ = _dafny.Map({d_164_v14_: d_160_v13_})
            index39_ = default__.safeIndex(360, (d_209_v47_).length(0))
            index40_ = default__.safeIndex(587, (d_160_v13_).length(0))
            def iife28_():
                coll6_ = _dafny.Set()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(803, 238):
                    d_214_v49_: int = compr_6_
                    if ((803) <= (d_214_v49_)) and ((d_214_v49_) < (238)):
                        coll6_ = coll6_.union(_dafny.Set([default__.safeDivisionInt(d_214_v49_, d_164_v14_.f5)]))
                return _dafny.Set(coll6_)
            rhs35_ = ((d_212_v50_) | (default__.fm3((d_209_v47_)[default__.safeIndex(360, (d_209_v47_).length(0))], not(d_149_v3_), d_148_globalState_))).issubset(iife28_()
            )
            rhs36_ = default__.fm0(d_148_globalState_)
            rhs37_ = ((d_211_v48_) + ((_dafny.SeqWithoutIsStrInference([((d_213_v51_)[d_164_v14_] if (d_164_v14_) in (d_213_v51_) else d_160_v13_)])))) + (_dafny.SeqWithoutIsStrInference([d_160_v13_]))
            rhs38_ = d_164_v14_
            lhs26_ = d_209_v47_
            lhs27_ = default__.safeIndex(360, (d_209_v47_).length(0))
            lhs28_ = d_160_v13_
            lhs29_ = default__.safeIndex(587, (d_160_v13_).length(0))
            lhs26_[lhs27_] = rhs35_
            lhs28_[lhs29_] = rhs36_
            d_211_v48_ = rhs37_
            d_164_v14_ = rhs38_
            (d_148_globalState_).f3 = d_165_v15_
            d_215_v52_: _dafny.Array
            d_216_v53_: int
            out6_: _dafny.Array
            out7_: int
            out6_, out7_ = default__.m0(d_148_globalState_)
            d_215_v52_ = out6_
            d_216_v53_ = out7_
        elif True:
            d_144_v0_ = (0) - (d_164_v14_.f5)
            d_217_v54_: _dafny.MultiSet
            d_217_v54_ = _dafny.MultiSet([d_144_v0_])
            if default__.fm1(d_164_v14_.f5, (len(_dafny.SeqWithoutIsStrInference([d_164_v14_.f5 for d_218_i10_ in range(default__.abs(930))]))) != (d_144_v0_), ((d_217_v54_).cardinality) < (d_144_v0_), d_149_v3_, d_148_globalState_):
                d_219_v55_: _dafny.Array
                nw39_ = _dafny.Array(False, 12)
                d_219_v55_ = nw39_
                index41_ = default__.safeIndex(358, (d_219_v55_).length(0))
                (d_219_v55_)[index41_] = not((d_164_v14_.f5) >= ((0) - (len(d_165_v15_))))
                index42_ = default__.safeIndex(358, (d_219_v55_).length(0))
                (d_219_v55_)[index42_] = d_149_v3_
                (d_148_globalState_).f1 = d_149_v3_
                d_220_v56_: _dafny.Seq
                d_220_v56_ = _dafny.SeqWithoutIsStrInference([(d_219_v55_)[default__.safeIndex(358, (d_219_v55_).length(0))], (d_219_v55_)[default__.safeIndex(358, (d_219_v55_).length(0))]])
                d_221_v57_: _dafny.Seq
                d_221_v57_ = _dafny.SeqWithoutIsStrInference([d_220_v56_, d_220_v56_])
                d_222_v58_: _dafny.Set
                d_222_v58_ = _dafny.Set({d_144_v0_, default__.safeModuloInt(d_144_v0_, d_164_v14_.f5), len((d_165_v15_).set(default__.safeIndex((_dafny.MultiSet([default__.fm1(len((d_221_v57_)[default__.safeIndex(d_164_v14_.f5, len(d_221_v57_))]), (d_219_v55_)[default__.safeIndex(358, (d_219_v55_).length(0))], True, d_149_v3_, d_148_globalState_), (d_219_v55_)[default__.safeIndex(358, (d_219_v55_).length(0))]])).cardinality, len(d_165_v15_)), default__.fm4((d_219_v55_)[default__.safeIndex(358, (d_219_v55_).length(0))], d_148_globalState_))), 321})
                def iife29_():
                    coll7_ = _dafny.Set()
                    compr_7_: int
                    for compr_7_ in (d_145_v1_).Elements:
                        d_223_v59_: int = compr_7_
                        if (d_223_v59_) in (d_145_v1_):
                            coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_223_v59_, -437)]))
                    return _dafny.Set(coll7_)
                d_222_v58_ = ((d_222_v58_) | (iife29_()
                )).intersection(_dafny.Set({d_144_v0_}))
                d_171_v19_ = (d_171_v19_ if True else d_171_v19_)
                d_224_v60_: _dafny.Array
                def lambda7_(d_225_v19_):
                    def lambda8_(d_226_i11_):
                        return d_225_v19_

                    return lambda8_

                init5_ = lambda7_(d_171_v19_)
                nw40_ = _dafny.Array(None, 16)
                for i0_5_ in range(nw40_.length(0)):
                    nw40_[i0_5_] = init5_(i0_5_)
                d_224_v60_ = nw40_
                d_227_v61_: _dafny.Map
                d_227_v61_ = _dafny.Map({d_144_v0_: d_224_v60_})
                d_224_v60_ = ((d_227_v61_)[d_144_v0_] if (d_144_v0_) in (d_227_v61_) else (d_224_v60_ if (d_219_v55_)[default__.safeIndex(358, (d_219_v55_).length(0))] else d_224_v60_))
            elif True:
                d_228_v62_: D2
                d_228_v62_ = D2_DC5(d_164_v14_.f5, d_164_v14_.f5)
                d_228_v62_ = d_228_v62_
                d_229_v63_: _dafny.MultiSet
                d_229_v63_ = _dafny.MultiSet([d_149_v3_])
                d_230_v64_: _dafny.Map
                d_230_v64_ = _dafny.Map({d_164_v14_: d_164_v14_})
                d_231_v65_: _dafny.Seq
                d_231_v65_ = _dafny.SeqWithoutIsStrInference([d_149_v3_])
                nw41_ = _dafny.Array(None, 8)
                nw41_[int(0)] = d_164_v14_.f5
                nw41_[int(1)] = d_164_v14_.f5
                nw41_[int(2)] = default__.safeDivisionInt((0) - ((d_229_v63_).cardinality), len(d_165_v15_))
                nw41_[int(3)] = default__.safeDivisionInt(d_164_v14_.f5, d_164_v14_.f5)
                nw41_[int(4)] = len(d_230_v64_)
                nw41_[int(5)] = default__.fm0(d_148_globalState_)
                nw41_[int(6)] = len(d_231_v65_)
                nw41_[int(7)] = d_164_v14_.f5
                d_160_v13_ = nw41_
                d_144_v0_ = (d_164_v14_.f5) - (d_164_v14_.f5)
                (d_164_v14_).f5 = d_144_v0_
                d_232_v66_: _dafny.Array
                d_233_v67_: int
                out8_: _dafny.Array
                out9_: int
                out8_, out9_ = default__.m0(d_148_globalState_)
                d_232_v66_ = out8_
                d_233_v67_ = out9_
            (d_148_globalState_).f1 = d_149_v3_
            d_234_v68_: D0
            d_234_v68_ = D0_DC0(d_145_v1_)
            pat_let_tv6_ = d_145_v1_
            def iife30_(_pat_let11_0):
                def iife31_(d_235_dt__update__tmp_h1_):
                    def iife32_(_pat_let12_0):
                        def iife33_(d_236_dt__update_hcf0_h0_):
                            return D0_DC0(d_236_dt__update_hcf0_h0_)
                        return iife33_(_pat_let12_0)
                    return iife32_(pat_let_tv6_)
                return iife31_(_pat_let11_0)
            d_234_v68_ = iife30_(d_234_v68_)
            d_237_v69_: _dafny.Set
            d_237_v69_ = _dafny.Set({d_149_v3_})
            d_238_v70_: _dafny.Set
            d_238_v70_ = _dafny.Set({d_237_v69_, d_237_v69_})
            d_238_v70_ = (d_238_v70_) | (d_238_v70_)
        d_239_v71_: _dafny.Array
        def lambda9_(d_240_v3_):
            def lambda10_(d_241_i12_):
                return d_240_v3_

            return lambda10_

        init6_ = lambda9_(d_149_v3_)
        nw42_ = _dafny.Array(None, 1)
        for i0_6_ in range(nw42_.length(0)):
            nw42_[i0_6_] = init6_(i0_6_)
        d_239_v71_ = nw42_
        d_239_v71_ = d_239_v71_
        d_242_v72_: _dafny.Map
        d_242_v72_ = _dafny.Map({d_149_v3_: d_239_v71_})
        d_242_v72_ = (d_242_v72_).set(False, d_239_v71_)
        (d_164_v14_).f5 = default__.safeModuloInt(d_164_v14_.f5, default__.safeModuloInt((d_145_v1_)[default__.safeIndex(d_164_v14_.f5, len(d_145_v1_))], default__.fm0(d_148_globalState_)))
        d_243_v73_: _dafny.Array
        d_244_v74_: int
        out10_: _dafny.Array
        out11_: int
        out10_, out11_ = default__.m0(d_148_globalState_)
        d_243_v73_ = out10_
        d_244_v74_ = out11_
        (d_148_globalState_).f1 = d_149_v3_
        hi1_ = d_164_v14_.f5
        for d_245_i13_ in range(d_164_v14_.f5, hi1_):
            d_246_v75_: _dafny.Array
            nw43_ = _dafny.Array(_dafny.Map({}), 6)
            d_246_v75_ = nw43_
            d_247_v76_: _dafny.Array
            nw44_ = _dafny.Array(_dafny.Seq({}), 17)
            d_247_v76_ = nw44_
            d_248_v77_: _dafny.Map
            d_248_v77_ = _dafny.Map({d_247_v76_: d_164_v14_.f5})
            index43_ = default__.safeIndex(843, (d_246_v75_).length(0))
            (d_246_v75_)[index43_] = (d_248_v77_ if d_149_v3_ else d_248_v77_)
            index44_ = default__.safeIndex(843, (d_246_v75_).length(0))
            (d_246_v75_)[index44_] = d_248_v77_
            d_149_v3_ = ((d_145_v1_)[default__.safeIndex(d_144_v0_, len(d_145_v1_))]) < (d_164_v14_.f5)
            d_144_v0_ = (len(d_165_v15_)) + (d_164_v14_.f5)
            (d_164_v14_).f5 = (d_144_v0_) * (d_245_i13_)
        _dafny.print(_dafny.string_of(d_144_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_145_v1_) == (_dafny.SeqWithoutIsStrInference([-560, -560]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[0]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[1]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[2]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[3]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[4]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[5]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[6]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[7]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[8]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[9]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[10]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[11]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[12]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[13]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[14]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[15]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[16]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[17]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_146_v2_)[18]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_148_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_148_globalState_).f2) == (_dafny.SeqWithoutIsStrInference([-560, -560]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_148_globalState_.f3).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[0]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[1]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[2]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[3]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[4]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[5]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[6]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[7]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[8]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[9]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[10]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[11]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[12]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[13]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[14]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[15]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[16]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[17]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_148_globalState_).f4)[18]) == (_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('s')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_149_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_150_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v13_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v13_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v13_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v13_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v13_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_164_v14_.f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_165_v15_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_166_i5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_171_v19_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[0]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[1]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[2]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[3]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[4]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[5]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[6]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[7]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[8]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[9]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[10]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[11]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[12]) == (_dafny.Map({9: 9, -9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[13]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[14]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[15]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_172_v20_)[16]) == (_dafny.Map({9: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v21_) == (_dafny.Map({True: 9}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_v22_) == (_dafny.Map({9: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_239_v71_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_242_v72_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_244_v74_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
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
        return lambda: D1_DC3()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3)
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [('cf4', Any), ('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(_dafny.MultiSet({}))
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

class D3_DC9(D3, NamedTuple('DC9', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)

class D4_DC11(D4, NamedTuple('DC11', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(D0.default()(), int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf19', Any), ('cf20', Any), ('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f1: bool = False
        self.f3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f0: int = int(0)
        self._f2: _dafny.Seq = _dafny.Seq({})
        self._f4: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4

class C0:
    def  __init__(self):
        self.f5: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f5):
        (self).f5 = f5

