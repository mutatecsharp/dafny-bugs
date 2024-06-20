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
        return len(_dafny.Map({_dafny.MultiSet([_dafny.CodePoint('j')]): _dafny.MultiSet([len(_dafny.Set({D9_DC24(not(False), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_0_i0_ in range(default__.abs(466))])))})), 596])}))

    @staticmethod
    def fm1(p0, globalState):
        return _dafny.MultiSet(((_dafny.SeqWithoutIsStrInference([_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hpvvkqoa"))))}), _dafny.Map({False: len(_dafny.Set({False}))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(False): -524})]))) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 766})])))

    @staticmethod
    def fm2(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: str
            for compr_0_ in (_dafny.Map({_dafny.CodePoint('c'): _dafny.Map({not(True): False})})).keys.Elements:
                d_1_v0_: str = compr_0_
                if (d_1_v0_) in (_dafny.Map({_dafny.CodePoint('c'): _dafny.Map({not(True): False})})):
                    coll0_[d_1_v0_] = False
            return _dafny.Map(coll0_)
        return (len((_dafny.SeqWithoutIsStrInference([not(False), not(False), True, True, False])) + (_dafny.SeqWithoutIsStrInference([False])))) <= ((len(iife0_()
        )) - (len(_dafny.Map({-803: len(_dafny.SeqWithoutIsStrInference([-538, -54]))}))))

    @staticmethod
    def fm3(p0, globalState):
        return _dafny.CodePoint('y')

    @staticmethod
    def fm4(p0, p1, globalState):
        return _dafny.Map({not ((D1_DC2(False)).cf2) or (not(not(True))): _dafny.SeqWithoutIsStrInference([True])})

    @staticmethod
    def fm5(p0, p1, p2, globalState):
        return D0_DC1(_dafny.CodePoint('c'))

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, not(True), not(True)])) + (_dafny.SeqWithoutIsStrInference([False]))) + (_dafny.SeqWithoutIsStrInference([False, True, False]))

    @staticmethod
    def fm11(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return D3_DC7((_dafny.Set({_dafny.MultiSet([409]), _dafny.MultiSet([(_dafny.MultiSet([176])).cardinality, 766]), _dafny.MultiSet([len(_dafny.Set({True, False}))])})).intersection(_dafny.Set({_dafny.MultiSet([-517, (0) - (len(_dafny.SeqWithoutIsStrInference([False, True])))]), _dafny.MultiSet([264])})))

    @staticmethod
    def fm13(p0, globalState):
        return ((_dafny.Map({_dafny.CodePoint('m'): not(True)})) | (_dafny.Map({_dafny.CodePoint('n'): not(False)}))) | ((_dafny.Map({_dafny.CodePoint('r'): True})) | (_dafny.Map({_dafny.CodePoint('g'): False})))

    @staticmethod
    def fm14(p0, p1, globalState):
        return _dafny.Set({D1_DC2(True)})

    @staticmethod
    def fm15(globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({False: not(False)}))) | (_dafny.Map({True: False}))

    @staticmethod
    def fm16(globalState):
        return ((_dafny.SeqWithoutIsStrInference([-115])) + (_dafny.SeqWithoutIsStrInference([879, len(_dafny.SeqWithoutIsStrInference([(D7_DC19(_dafny.CodePoint('n'), False)).cf29 for d_2_i0_ in range(default__.abs(-823))]))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.Map({not(True): False}) for d_3_i1_ in range(default__.abs(590))]))])) + (_dafny.SeqWithoutIsStrInference([373])))

    @staticmethod
    def fm17(p0, p1, globalState):
        return (_dafny.MultiSet([False, False])).intersection((_dafny.MultiSet([False])) | (_dafny.MultiSet([False])))

    @staticmethod
    def fm18(globalState):
        return _dafny.MultiSet([_dafny.Map({len(_dafny.Map({True: (D4_DC11(len(_dafny.SeqWithoutIsStrInference([D1_DC3(_dafny.CodePoint('t'))])), _dafny.CodePoint('f'))).cf14})): (_dafny.MultiSet((D6_DC16(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({916: 337})))]))).cf22)).cardinality}), _dafny.Map({44: 907})])

    @staticmethod
    def fm19(globalState):
        def iife1_():
            def iife4_():
                def iife5_():
                    coll5_ = _dafny.Map()
                    compr_5_: int
                    for compr_5_ in _dafny.IntegerRange(916, 690):
                        d_6_v2_: int = compr_5_
                        if ((916) <= (d_6_v2_)) and ((d_6_v2_) < (690)):
                            coll5_[(d_6_v2_) - (893)] = len(_dafny.Map({len(_dafny.Map({True: 90})): _dafny.CodePoint('c')}))
                    return _dafny.Map(coll5_)
                coll4_ = _dafny.Map()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(375, 105):
                    d_4_v1_: int = compr_4_
                    if ((375) <= (d_4_v1_)) and ((d_4_v1_) < (105)):
                        coll4_[(d_4_v1_) * (len(iife5_()
                        ))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_5_i0_ in range(default__.abs(434))]))
                return _dafny.Map(coll4_)
            coll1_ = _dafny.Map()
            def iife2_():
                def iife3_():
                    coll3_ = _dafny.Map()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(916, 690):
                        d_6_v2_: int = compr_3_
                        if ((916) <= (d_6_v2_)) and ((d_6_v2_) < (690)):
                            coll3_[(d_6_v2_) - (893)] = len(_dafny.Map({len(_dafny.Map({True: 90})): _dafny.CodePoint('c')}))
                    return _dafny.Map(coll3_)
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(375, 105):
                    d_4_v1_: int = compr_2_
                    if ((375) <= (d_4_v1_)) and ((d_4_v1_) < (105)):
                        coll2_[(d_4_v1_) * (len(iife3_()
                        ))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_5_i0_ in range(default__.abs(434))]))
                return _dafny.Map(coll2_)
            compr_1_: int
            for compr_1_ in (iife2_()
            ).keys.Elements:
                d_7_v0_: int = compr_1_
                if (d_7_v0_) in (iife4_()
                ):
                    coll1_[default__.safeModuloInt(d_7_v0_, -34)] = 360
            return _dafny.Map(coll1_)
        source0_ = (D4_DC10(_dafny.Map({516: 525})) if False else D4_DC10(iife1_()
))
        if source0_.is_DC11:
            d_8___mcc_h0_ = source0_.cf13
            d_9___mcc_h1_ = source0_.cf14
            d_10_cf14_ = d_9___mcc_h1_
            d_11_cf13_ = d_8___mcc_h0_
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: _dafny.MultiSet
                for compr_6_ in (_dafny.Set({_dafny.MultiSet([d_11_cf13_, d_11_cf13_, d_11_cf13_])})).Elements:
                    d_12_v3_: _dafny.MultiSet = compr_6_
                    if (d_12_v3_) in (_dafny.Set({_dafny.MultiSet([d_11_cf13_, d_11_cf13_, d_11_cf13_])})):
                        coll6_ = coll6_.union(_dafny.Set([d_12_v3_]))
                return _dafny.Set(coll6_)
            return D3_DC9(D3_DC9(D3_DC7(iife6_()
)))
        elif source0_.is_DC12:
            d_13___mcc_h2_ = source0_.cf15
            d_14___mcc_h3_ = source0_.cf16
            d_15___mcc_h4_ = source0_.cf17
            d_16_cf17_ = d_15___mcc_h4_
            d_17_cf16_ = d_14___mcc_h3_
            d_18_cf15_ = d_13___mcc_h2_
            return D3_DC9(D3_DC8(d_18_cf15_))
        elif source0_.is_DC10:
            d_19___mcc_h5_ = source0_.cf12
            d_20_cf12_ = d_19___mcc_h5_
            return D3_DC9(D3_DC8(923))
        elif True:
            d_21___mcc_h6_ = source0_.cf18
            d_22_cf18_ = d_21___mcc_h6_
            return D3_DC9(D3_DC7(_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uh"))), len(_dafny.Map({True: len(_dafny.Map({573: True}))}))]), _dafny.MultiSet([len(_dafny.Set({not(not(True))}))])})))

    @staticmethod
    def fm20(p0, p1, globalState):
        def iife7_():
            coll7_ = _dafny.Map()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(910, 290):
                d_23_v0_: int = compr_7_
                if ((910) <= (d_23_v0_)) and ((d_23_v0_) < (290)):
                    coll7_[default__.safeModuloInt(d_23_v0_, 50)] = _dafny.CodePoint('q')
            return _dafny.Map(coll7_)
        source1_ = (D7_DC18(_dafny.Map({_dafny.Map({_dafny.CodePoint('m'): False}): iife7_()
})) if not(True) else D7_DC18(_dafny.Map({_dafny.Map({_dafny.CodePoint('m'): True}): _dafny.Map({37: _dafny.CodePoint('y')})})))
        if source1_.is_DC19:
            d_24___mcc_h0_ = source1_.cf29
            d_25___mcc_h1_ = source1_.cf30
            d_26_cf30_ = d_25___mcc_h1_
            d_27_cf29_ = d_24___mcc_h0_
            return (_dafny.Set({d_26_cf30_, False, d_26_cf30_})).intersection(_dafny.Set({d_26_cf30_}))
        elif True:
            d_28___mcc_h2_ = source1_.cf28
            d_29_cf28_ = d_28___mcc_h2_
            return (_dafny.Set({True})) | (_dafny.Set({True}))

    @staticmethod
    def fm21(p0, p1, globalState):
        return _dafny.Map({False: (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqgqi")))) * (-128)})

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(719, 317):
                d_30_v0_: int = compr_8_
                if ((719) <= (d_30_v0_)) and ((d_30_v0_) < (317)):
                    coll8_[(d_30_v0_) * ((0) - (len(_dafny.Map({True: _dafny.CodePoint('r')}))))] = ((D6_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyc")), False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eda"))), True, False)).cf25) - (47)
            return _dafny.Map(coll8_)
        return iife8_()
        

    @staticmethod
    def fm23(globalState):
        return D4_DC10((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.SeqWithoutIsStrInference([True]))})) | (_dafny.Map({len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([118]))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbwtruw")))})): 304})))

    @staticmethod
    def m0(globalState):
        d_31_v0_: bool
        d_31_v0_ = True
        (globalState).f3 = d_31_v0_
        d_32_v1_: int
        d_32_v1_ = 524
        d_33_v2_: _dafny.Map
        d_33_v2_ = _dafny.Map({default__.fm0(_dafny.CodePoint('x'), globalState): d_32_v1_})
        d_33_v2_ = (d_33_v2_).set(d_32_v1_, ((0) - (d_32_v1_)) - (d_32_v1_))
        d_34_v3_: _dafny.Array
        def lambda0_(d_35_v0_):
            def lambda1_(d_36_i0_):
                return (D7_DC19(_dafny.CodePoint('o'), d_35_v0_)).cf30

            return lambda1_

        init0_ = lambda0_(d_31_v0_)
        nw0_ = _dafny.Array(None, 26)
        for i0_0_ in range(nw0_.length(0)):
            nw0_[i0_0_] = init0_(i0_0_)
        d_34_v3_ = nw0_
        d_37_v4_: D8
        d_37_v4_ = D8_DC21(d_32_v1_)
        pat_let_tv0_ = d_31_v0_
        pat_let_tv1_ = d_31_v0_
        pat_let_tv2_ = d_31_v0_
        pat_let_tv3_ = d_31_v0_
        def lambda2_(source2_):
            if source2_.is_DC21:
                d_38___mcc_h0_ = source2_.cf32
                d_39_cf32_ = d_38___mcc_h0_
                return not((_dafny.MultiSet([665, d_39_cf32_])).isdisjoint(_dafny.MultiSet([d_39_cf32_, (0) - ((0) - ((D6_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ekoegwtab")), pat_let_tv0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mkjxih"))), pat_let_tv1_, pat_let_tv2_)).cf25))])))
            elif True:
                d_40___mcc_h1_ = source2_.cf31
                d_41_cf31_ = d_40___mcc_h1_
                return pat_let_tv3_

        rhs0_ = (d_32_v1_) == (d_32_v1_)
        rhs1_ = lambda2_(d_37_v4_)
        rhs2_ = (d_34_v3_ if d_31_v0_ else d_34_v3_)
        lhs0_ = globalState
        lhs0_.f3 = rhs0_
        d_31_v0_ = rhs1_
        d_34_v3_ = rhs2_
        d_42_v5_: _dafny.Map
        d_42_v5_ = _dafny.Map({d_31_v0_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jouwqcwf")))})
        d_43_v6_: str
        d_43_v6_ = _dafny.CodePoint('x')
        d_44_v7_: D4
        d_44_v7_ = D4_DC11(d_32_v1_, d_43_v6_)
        d_45_v8_: _dafny.MultiSet
        d_45_v8_ = _dafny.MultiSet([d_31_v0_])
        d_46_v9_: C2
        nw1_ = C2()
        nw1_.ctor__((d_44_v7_).cf13, d_32_v1_, ((d_45_v8_)[not(not(d_31_v0_))] if (not(not(d_31_v0_))) in (d_45_v8_) else d_32_v1_), len(d_33_v2_))
        d_46_v9_ = nw1_
        d_47_v10_: _dafny.Map
        d_47_v10_ = _dafny.Map({(d_42_v5_) | (d_42_v5_): d_46_v9_})
        d_48_v11_: _dafny.Seq
        d_48_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyfp"))
        d_47_v10_ = (d_47_v10_).set(((d_42_v5_).set(d_31_v0_, (0) - (len(d_48_v11_)))) | (d_42_v5_), d_46_v9_)
        d_49_v12_: _dafny.Set
        d_49_v12_ = _dafny.Set({d_34_v3_, d_34_v3_, d_34_v3_})
        hi0_ = (d_46_v9_).f9
        for d_50_i1_ in range(len(d_49_v12_), hi0_):
            if d_31_v0_:
                d_51_v13_: T0
                nw2_ = C2()
                nw2_.ctor__(d_46_v9_.f10, d_32_v1_, (d_46_v9_).f9, (d_46_v9_).f9)
                d_51_v13_ = nw2_
                d_52_v14_: _dafny.Seq
                d_52_v14_ = _dafny.SeqWithoutIsStrInference([d_51_v13_])
                d_53_v15_: _dafny.Map
                d_53_v15_ = _dafny.Map({d_31_v0_: d_52_v14_})
                (d_46_v9_).f10 = len(d_53_v15_)
                (globalState).f3 = not ((d_31_v0_) == (d_31_v0_)) or (d_31_v0_)
                (d_46_v9_).f10 = (d_51_v13_).f7
                d_34_v3_ = (d_34_v3_ if d_31_v0_ else (d_34_v3_ if d_31_v0_ else d_34_v3_))
                d_54_v16_: _dafny.Seq
                d_54_v16_ = _dafny.SeqWithoutIsStrInference([d_32_v1_, (d_46_v9_).f9])
                d_48_v11_ = (default__.fm11(d_31_v0_, (0) - (d_50_i1_), globalState)).set(default__.safeIndex((d_54_v16_)[default__.safeIndex(d_50_i1_, len(d_54_v16_))], len(default__.fm11(d_31_v0_, (0) - (d_50_i1_), globalState))), d_43_v6_)
            elif True:
                d_55_v17_: _dafny.Map
                d_55_v17_ = _dafny.Map({True: d_45_v8_})
                (globalState).f1 = not((((d_55_v17_)[not(d_31_v0_)] if (not(d_31_v0_)) in (d_55_v17_) else d_45_v8_)).issubset(default__.fm17(d_31_v0_, (d_46_v9_).f9, globalState)))
                d_56_v18_: _dafny.Map
                d_56_v18_ = _dafny.Map({d_43_v6_: d_31_v0_})
                d_57_v20_: _dafny.Map
                def iife9_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(791, 145):
                        d_58_v19_: int = compr_9_
                        if ((791) <= (d_58_v19_)) and ((d_58_v19_) < (145)):
                            coll9_[(d_58_v19_) + (293)] = d_43_v6_
                    return _dafny.Map(coll9_)
                d_57_v20_ = _dafny.Map({d_56_v18_: iife9_()
                })
                d_59_v21_: C4
                nw3_ = C4()
                nw3_.ctor__((d_57_v20_) | (d_57_v20_))
                d_59_v21_ = nw3_
                d_60_v22_: _dafny.Seq
                d_60_v22_ = _dafny.SeqWithoutIsStrInference([d_31_v0_])
                d_61_v23_: D4
                d_61_v23_ = D4_DC12((d_46_v9_).f9, d_60_v22_, d_31_v0_)
                d_62_v24_: C1
                nw4_ = C1()
                nw4_.ctor__(d_31_v0_, (0) - ((d_45_v8_).cardinality), (d_46_v9_).f9, len((d_61_v23_).cf16))
                d_62_v24_ = nw4_
                d_63_v25_: _dafny.Array
                nw5_ = _dafny.Array(None, 3)
                nw5_[int(0)] = d_62_v24_
                nw5_[int(1)] = d_62_v24_
                nw5_[int(2)] = d_62_v24_
                d_63_v25_ = nw5_
                index0_ = default__.safeIndex(875, (d_63_v25_).length(0))
                (d_63_v25_)[index0_] = d_62_v24_
                index1_ = default__.safeIndex(875, (d_63_v25_).length(0))
                (d_63_v25_)[index1_] = d_62_v24_
                d_32_v1_ = (0) - (d_50_i1_)
                d_64_v26_: _dafny.Array
                nw6_ = _dafny.Array(_dafny.Map({}), 2)
                d_64_v26_ = nw6_
                d_64_v26_ = (d_64_v26_ if d_31_v0_ else d_64_v26_)
            d_65_v27_: _dafny.Array
            nw7_ = _dafny.Array(_dafny.Array(None, 0), 25)
            d_65_v27_ = nw7_
            d_66_v28_: D4
            d_66_v28_ = D4_DC13(default__.fm23(globalState))
            d_67_v29_: D4
            d_67_v29_ = D4_DC13(d_66_v28_)
            d_68_v30_: _dafny.Seq
            d_68_v30_ = _dafny.SeqWithoutIsStrInference([d_67_v29_])
            d_69_v31_: _dafny.Map
            d_69_v31_ = _dafny.Map({d_65_v27_: d_68_v30_})
            d_69_v31_ = (d_69_v31_).set(d_65_v27_, d_68_v30_)
            d_70_v32_: _dafny.Seq
            d_70_v32_ = _dafny.SeqWithoutIsStrInference([d_31_v0_, True, d_31_v0_, False])
            (d_46_v9_).f10 = (0) - (len(((d_70_v32_ if d_31_v0_ else d_70_v32_)) + (d_70_v32_)))
            d_71_v34_: _dafny.Map
            d_71_v34_ = _dafny.Map({not(False): d_33_v2_})
            d_72_v35_: _dafny.Seq
            d_72_v35_ = _dafny.SeqWithoutIsStrInference([(d_46_v9_).f9, d_50_i1_, len(((d_71_v34_)[True] if (True) in (d_71_v34_) else d_33_v2_)), (d_46_v9_).fm7((default__.fm11(d_31_v0_, (d_46_v9_).f9, globalState)).set(default__.safeIndex(d_50_i1_, len(default__.fm11(d_31_v0_, (d_46_v9_).f9, globalState))), d_43_v6_), (d_46_v9_).f9, globalState)])
            d_73_v36_: C1
            nw8_ = C1()
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in (d_72_v35_).Elements:
                    d_74_v33_: int = compr_10_
                    if (d_74_v33_) in (d_72_v35_):
                        coll10_[(d_74_v33_) + (d_46_v9_.f10)] = d_46_v9_.f10
                return _dafny.Map(coll10_)
            nw8_.ctor__((((d_42_v5_)[d_31_v0_] if (d_31_v0_) in (d_42_v5_) else (d_46_v9_).f9)) <= ((0) - (len(_dafny.SeqWithoutIsStrInference([d_31_v0_])))), len(iife10_()
            ), (d_46_v9_.f10) + ((d_46_v9_).f9), d_46_v9_.f10)
            d_73_v36_ = nw8_
            d_73_v36_ = d_73_v36_
        d_75_v37_: _dafny.Set
        d_75_v37_ = _dafny.Set({d_31_v0_, d_31_v0_, d_31_v0_})
        index2_ = default__.safeIndex(468, (d_34_v3_).length(0))
        (d_34_v3_)[index2_] = (d_75_v37_).ispropersubset(d_75_v37_)
        index3_ = default__.safeIndex(468, (d_34_v3_).length(0))
        (d_34_v3_)[index3_] = default__.fm2(d_31_v0_, globalState)

    @staticmethod
    def Main(noArgsParameter__):
        d_76_globalState_: GlobalState
        nw9_ = GlobalState()
        nw9_.ctor__(True, True, 135, True, False, _dafny.CodePoint('e'))
        d_76_globalState_ = nw9_
        d_77_v0_: _dafny.Seq
        d_77_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jc"))
        d_77_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqitbtm"))
        d_78_v1_: int
        d_78_v1_ = 414
        d_79_v2_: bool
        d_79_v2_ = True
        d_80_v3_: _dafny.Seq
        d_80_v3_ = _dafny.SeqWithoutIsStrInference([d_79_v2_, True])
        d_81_v4_: str
        d_81_v4_ = _dafny.CodePoint('j')
        d_82_v5_: _dafny.Map
        d_82_v5_ = _dafny.Map({d_79_v2_: len(d_80_v3_)})
        d_83_v6_: _dafny.Map
        d_83_v6_ = _dafny.Map({d_80_v3_: ((d_82_v5_)[not(True)] if (not(True)) in (d_82_v5_) else d_78_v1_)})
        d_78_v1_ = default__.safeDivisionInt(d_78_v1_, len((_dafny.Map({d_80_v3_: default__.fm0(d_81_v4_, d_76_globalState_)})) | (d_83_v6_)))
        if not(d_79_v2_):
            d_84_v7_: _dafny.Array
            nw10_ = _dafny.Array(None, 2)
            nw10_[int(0)] = d_78_v1_
            nw10_[int(1)] = d_78_v1_
            d_84_v7_ = nw10_
            d_85_v8_: _dafny.Map
            d_85_v8_ = _dafny.Map({d_79_v2_: d_81_v4_})
            d_86_v9_: _dafny.Map
            d_86_v9_ = _dafny.Map({(d_85_v8_).set(d_79_v2_, d_81_v4_): default__.fm2(d_79_v2_, d_76_globalState_)})
            d_87_v10_: _dafny.Map
            d_87_v10_ = _dafny.Map({default__.fm1(((d_86_v9_)[d_85_v8_] if (d_85_v8_) in (d_86_v9_) else d_79_v2_), d_76_globalState_): d_84_v7_})
            d_88_v11_: _dafny.MultiSet
            d_88_v11_ = _dafny.MultiSet([d_82_v5_])
            d_89_v12_: D0
            d_89_v12_ = D0_DC0(d_84_v7_)
            d_90_v13_: _dafny.Array
            nw11_ = _dafny.Array(None, 28)
            nw11_[int(0)] = d_84_v7_
            nw11_[int(1)] = d_84_v7_
            nw11_[int(2)] = d_84_v7_
            nw11_[int(3)] = d_84_v7_
            nw11_[int(4)] = d_84_v7_
            nw11_[int(5)] = d_84_v7_
            nw11_[int(6)] = d_84_v7_
            nw11_[int(7)] = d_84_v7_
            nw11_[int(8)] = d_84_v7_
            nw11_[int(9)] = d_84_v7_
            nw11_[int(10)] = d_84_v7_
            nw11_[int(11)] = d_84_v7_
            nw11_[int(12)] = d_84_v7_
            nw11_[int(13)] = d_84_v7_
            nw11_[int(14)] = ((d_87_v10_)[d_88_v11_] if (d_88_v11_) in (d_87_v10_) else d_84_v7_)
            nw11_[int(15)] = d_84_v7_
            nw11_[int(16)] = d_84_v7_
            nw11_[int(17)] = d_84_v7_
            nw11_[int(18)] = d_84_v7_
            nw11_[int(19)] = d_84_v7_
            nw11_[int(20)] = d_84_v7_
            nw11_[int(21)] = d_84_v7_
            nw11_[int(22)] = d_84_v7_
            nw11_[int(23)] = (d_84_v7_ if d_79_v2_ else (d_89_v12_).cf0)
            nw11_[int(24)] = d_84_v7_
            nw11_[int(25)] = d_84_v7_
            nw11_[int(26)] = d_84_v7_
            nw11_[int(27)] = d_84_v7_
            d_90_v13_ = nw11_
            rhs3_ = d_90_v13_
            rhs4_ = (_dafny.SeqWithoutIsStrInference([d_81_v4_ for d_91_i0_ in range(default__.abs(-42))])) + ((d_77_v0_) + (d_77_v0_))
            d_90_v13_ = rhs3_
            d_77_v0_ = rhs4_
            (d_76_globalState_).f3 = d_79_v2_
            d_79_v2_ = d_79_v2_
            d_92_v14_: D0
            d_92_v14_ = D0_DC1(d_81_v4_)
            d_92_v14_ = d_92_v14_
            d_78_v1_ = (466) + (len(d_77_v0_))
        elif True:
            default__.m0(d_76_globalState_)
            if (d_80_v3_)[default__.safeIndex((d_78_v1_ if d_79_v2_ else d_78_v1_), len(d_80_v3_))]:
                d_93_v15_: _dafny.Set
                d_93_v15_ = _dafny.Set({d_81_v4_, d_81_v4_, d_81_v4_, default__.fm3(d_79_v2_, d_76_globalState_)})
                d_94_v16_: _dafny.Seq
                d_94_v16_ = _dafny.SeqWithoutIsStrInference([d_93_v15_, d_93_v15_])
                d_94_v16_ = (d_94_v16_) + (_dafny.SeqWithoutIsStrInference([d_93_v15_, d_93_v15_, d_93_v15_]))
                d_77_v0_ = d_77_v0_
                d_95_v17_: _dafny.Array
                def lambda3_(d_96_v1_):
                    def lambda4_(d_97_i1_):
                        return (d_97_i1_) - (d_96_v1_)

                    return lambda4_

                init1_ = lambda3_(d_78_v1_)
                nw12_ = _dafny.Array(None, 10)
                for i0_1_ in range(nw12_.length(0)):
                    nw12_[i0_1_] = init1_(i0_1_)
                d_95_v17_ = nw12_
                index4_ = default__.safeIndex(415, (d_95_v17_).length(0))
                (d_95_v17_)[index4_] = (d_78_v1_) * (d_78_v1_)
                d_98_v18_: _dafny.Array
                nw13_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 8)
                d_98_v18_ = nw13_
                index5_ = default__.safeIndex(395, (d_98_v18_).length(0))
                (d_98_v18_)[index5_] = d_77_v0_
                d_99_v19_: _dafny.Set
                d_99_v19_ = _dafny.Set({d_79_v2_})
                d_100_v20_: _dafny.Array
                nw14_ = _dafny.Array(None, 8)
                nw14_[int(0)] = d_81_v4_
                nw14_[int(1)] = (d_77_v0_)[default__.safeIndex(len(d_99_v19_), len(d_77_v0_))]
                nw14_[int(2)] = default__.fm3(d_79_v2_, d_76_globalState_)
                nw14_[int(3)] = d_81_v4_
                nw14_[int(4)] = d_81_v4_
                nw14_[int(5)] = d_81_v4_
                nw14_[int(6)] = d_81_v4_
                nw14_[int(7)] = d_81_v4_
                d_100_v20_ = nw14_
                index6_ = default__.safeIndex(438, (d_100_v20_).length(0))
                (d_100_v20_)[index6_] = _dafny.CodePoint('t')
                index7_ = default__.safeIndex(415, (d_95_v17_).length(0))
                index8_ = default__.safeIndex(395, (d_98_v18_).length(0))
                index9_ = default__.safeIndex(438, (d_100_v20_).length(0))
                rhs5_ = (0) - (len(default__.fm4((0) - (d_78_v1_), d_79_v2_, d_76_globalState_)))
                rhs6_ = d_77_v0_
                rhs7_ = d_81_v4_
                rhs8_ = d_81_v4_
                rhs9_ = default__.safeModuloInt(default__.fm0(d_81_v4_, d_76_globalState_), default__.fm0(d_81_v4_, d_76_globalState_))
                lhs1_ = d_95_v17_
                lhs2_ = default__.safeIndex(415, (d_95_v17_).length(0))
                lhs3_ = d_98_v18_
                lhs4_ = default__.safeIndex(395, (d_98_v18_).length(0))
                lhs5_ = d_100_v20_
                lhs6_ = default__.safeIndex(438, (d_100_v20_).length(0))
                lhs1_[lhs2_] = rhs5_
                lhs3_[lhs4_] = rhs6_
                d_81_v4_ = rhs7_
                lhs5_[lhs6_] = rhs8_
                d_78_v1_ = rhs9_
                d_101_v21_: D0
                d_101_v21_ = D0_DC1(d_81_v4_)
                d_101_v21_ = D0_DC1((d_100_v20_)[default__.safeIndex(438, (d_100_v20_).length(0))])
                index10_ = default__.safeIndex(395, (d_98_v18_).length(0))
                (d_98_v18_)[index10_] = ((d_98_v18_)[default__.safeIndex(395, (d_98_v18_).length(0))]) + ((d_98_v18_)[default__.safeIndex(395, (d_98_v18_).length(0))])
            elif True:
                default__.m0(d_76_globalState_)
                d_102_v22_: _dafny.Seq
                d_102_v22_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_81_v4_, d_76_globalState_), 800])
                d_103_v23_: C1
                nw15_ = C1()
                nw15_.ctor__(d_79_v2_, d_78_v1_, len(d_102_v22_), len((d_77_v0_) + (d_77_v0_)))
                d_103_v23_ = nw15_
                d_104_v24_: _dafny.MultiSet
                d_104_v24_ = _dafny.MultiSet([d_79_v2_, (d_103_v23_).f12])
                d_105_v25_: _dafny.Map
                d_105_v25_ = _dafny.Map({((d_104_v24_).set((d_103_v23_).f12, default__.abs(d_78_v1_))).cardinality: (d_77_v0_).set(default__.safeIndex(d_103_v23_.f13, len(d_77_v0_)), _dafny.CodePoint('q'))})
                d_106_v26_: T0
                nw16_ = C2()
                nw16_.ctor__(len(d_105_v25_), default__.safeDivisionInt(852, d_78_v1_), (d_103_v23_.f13 if (d_103_v23_).f12 else d_103_v23_.f13), (d_102_v22_)[default__.safeIndex(d_78_v1_, len(d_102_v22_))])
                d_106_v26_ = nw16_
                d_106_v26_ = d_106_v26_
                d_81_v4_ = d_81_v4_
                (d_76_globalState_).f3 = (105) != ((d_106_v26_).f7)
            d_107_v28_: D1
            d_107_v28_ = D1_DC3(d_81_v4_)
            d_108_v29_: _dafny.Map
            d_108_v29_ = _dafny.Map({default__.fm13(d_78_v1_, d_76_globalState_): d_107_v28_})
            d_109_v30_: C4
            nw17_ = C4()
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: _dafny.Map
                for compr_11_ in (d_108_v29_).keys.Elements:
                    d_110_v27_: _dafny.Map = compr_11_
                    if (d_110_v27_) in (d_108_v29_):
                        coll11_[d_110_v27_] = _dafny.Map({d_78_v1_: _dafny.CodePoint('y')})
                return _dafny.Map(coll11_)
            nw17_.ctor__(iife11_()
            )
            d_109_v30_ = nw17_
            d_109_v30_ = d_109_v30_
            d_111_v31_: C0
            nw18_ = C0()
            nw18_.ctor__(d_79_v2_)
            d_111_v31_ = nw18_
            d_112_v32_: _dafny.Array
            nw19_ = _dafny.Array(None, 12)
            nw19_[int(0)] = d_111_v31_
            nw19_[int(1)] = d_111_v31_
            nw19_[int(2)] = d_111_v31_
            nw19_[int(3)] = d_111_v31_
            nw19_[int(4)] = d_111_v31_
            nw19_[int(5)] = d_111_v31_
            nw19_[int(6)] = d_111_v31_
            nw19_[int(7)] = d_111_v31_
            nw19_[int(8)] = d_111_v31_
            nw19_[int(9)] = d_111_v31_
            nw19_[int(10)] = d_111_v31_
            nw19_[int(11)] = d_111_v31_
            d_112_v32_ = nw19_
            d_112_v32_ = d_112_v32_
            d_78_v1_ = default__.safeModuloInt(d_78_v1_, (d_78_v1_) * ((0) - (len(default__.fm21(len(d_80_v3_), d_78_v1_, d_76_globalState_)))))
        d_78_v1_ = -958
        d_113_v34_: _dafny.Map
        def iife12_():
            coll12_ = _dafny.Map()
            compr_12_: str
            for compr_12_ in (d_77_v0_).Elements:
                d_114_v33_: str = compr_12_
                if (d_114_v33_) in (d_77_v0_):
                    coll12_[d_114_v33_] = d_79_v2_
            return _dafny.Map(coll12_)
        d_113_v34_ = _dafny.Map({iife12_()
        : _dafny.Map({d_78_v1_: _dafny.CodePoint('q')})})
        d_115_v35_: D7
        d_115_v35_ = D7_DC18(d_113_v34_)
        d_116_v36_: C4
        nw20_ = C4()
        nw20_.ctor__((D7_DC18((d_115_v35_).cf28)).cf28)
        d_116_v36_ = nw20_
        if d_79_v2_:
            if (d_79_v2_) or (d_79_v2_):
                d_117_v37_: C0
                nw21_ = C0()
                nw21_.ctor__(d_79_v2_)
                d_117_v37_ = nw21_
                d_118_v38_: _dafny.Seq
                d_118_v38_ = _dafny.SeqWithoutIsStrInference([d_117_v37_])
                d_119_v39_: _dafny.Map
                d_119_v39_ = _dafny.Map({(d_117_v37_).f11: d_118_v38_})
                d_118_v38_ = (d_118_v38_) + (((d_119_v39_)[(d_117_v37_).f11] if ((d_117_v37_).f11) in (d_119_v39_) else d_118_v38_))
                d_120_v40_: _dafny.Array
                nw22_ = _dafny.Array(int(0), 5)
                d_120_v40_ = nw22_
                d_121_v41_: _dafny.Map
                d_121_v41_ = _dafny.Map({default__.fm2(d_79_v2_, d_76_globalState_): not(False)})
                d_122_v42_: _dafny.Seq
                d_122_v42_ = _dafny.SeqWithoutIsStrInference([897])
                d_123_v43_: D0
                d_123_v43_ = D0_DC1(d_81_v4_)
                d_124_v44_: _dafny.Array
                nw23_ = _dafny.Array(False, 22)
                d_124_v44_ = nw23_
                d_125_v45_: _dafny.Seq
                d_125_v45_ = _dafny.SeqWithoutIsStrInference([d_124_v44_, d_124_v44_, d_124_v44_])
                d_126_v46_: D8
                d_126_v46_ = D8_DC20(d_125_v45_)
                nw24_ = _dafny.Array(None, 17)
                nw24_[int(0)] = 935
                nw24_[int(1)] = d_78_v1_
                nw24_[int(2)] = len((d_121_v41_).set(d_79_v2_, (d_117_v37_).f11))
                nw24_[int(3)] = len((d_122_v42_) + (d_122_v42_))
                nw24_[int(4)] = d_78_v1_
                nw24_[int(5)] = (d_78_v1_) - (d_78_v1_)
                nw24_[int(6)] = (195) - (d_78_v1_)
                nw24_[int(7)] = d_78_v1_
                nw24_[int(8)] = (0) - (default__.fm0((d_123_v43_).cf1, d_76_globalState_))
                nw24_[int(9)] = default__.safeDivisionInt((0) - (d_78_v1_), d_78_v1_)
                nw24_[int(10)] = (len((d_77_v0_).set(default__.safeIndex(d_78_v1_, len(d_77_v0_)), d_81_v4_))) * (d_78_v1_)
                nw24_[int(11)] = len(_dafny.Map({d_79_v2_: True}))
                nw24_[int(12)] = len((d_122_v42_ if d_79_v2_ else _dafny.SeqWithoutIsStrInference([d_78_v1_, d_78_v1_, d_78_v1_])))
                nw24_[int(13)] = (0) - (d_78_v1_)
                nw24_[int(14)] = d_78_v1_
                nw24_[int(15)] = len((d_126_v46_).cf31)
                nw24_[int(16)] = default__.safeModuloInt(d_78_v1_, d_78_v1_)
                d_120_v40_ = nw24_
                d_127_v47_: C1
                nw25_ = C1()
                nw25_.ctor__((d_117_v37_).f11, d_78_v1_, d_78_v1_, d_78_v1_)
                d_127_v47_ = nw25_
                d_128_v48_: _dafny.Map
                d_128_v48_ = _dafny.Map({d_127_v47_: (d_127_v47_).f12})
                d_129_v49_: _dafny.Map
                d_129_v49_ = _dafny.Map({d_78_v1_: d_127_v47_})
                d_128_v48_ = (d_128_v48_).set(((d_129_v49_)[default__.fm0(d_81_v4_, d_76_globalState_)] if (default__.fm0(d_81_v4_, d_76_globalState_)) in (d_129_v49_) else d_127_v47_), d_79_v2_)
                index11_ = default__.safeIndex(785, (d_124_v44_).length(0))
                (d_124_v44_)[index11_] = d_79_v2_
                index12_ = default__.safeIndex(785, (d_124_v44_).length(0))
                (d_124_v44_)[index12_] = False
                d_130_v50_: D6
                d_130_v50_ = D6_DC17(d_77_v0_, d_79_v2_, len(d_80_v3_), default__.fm2((d_124_v44_)[default__.safeIndex(785, (d_124_v44_).length(0))], d_76_globalState_), (d_124_v44_)[default__.safeIndex(785, (d_124_v44_).length(0))])
                (d_127_v47_).f13 = (d_130_v50_).cf25
            elif True:
                d_131_v51_: _dafny.Array
                nw26_ = _dafny.Array(int(0), 10)
                d_131_v51_ = nw26_
                d_132_v52_: _dafny.Set
                d_132_v52_ = _dafny.Set({d_78_v1_})
                d_133_v53_: _dafny.Seq
                d_133_v53_ = _dafny.SeqWithoutIsStrInference([len(d_82_v5_), d_78_v1_])
                index13_ = default__.safeIndex(150, (d_131_v51_).length(0))
                (d_131_v51_)[index13_] = default__.safeModuloInt(len(d_132_v52_), (d_133_v53_)[default__.safeIndex(d_78_v1_, len(d_133_v53_))])
                index14_ = default__.safeIndex(150, (d_131_v51_).length(0))
                rhs10_ = (0) - ((937) - ((d_78_v1_) * (d_78_v1_)))
                rhs11_ = d_79_v2_
                rhs12_ = d_78_v1_
                lhs7_ = d_76_globalState_
                lhs8_ = d_131_v51_
                lhs9_ = default__.safeIndex(150, (d_131_v51_).length(0))
                d_78_v1_ = rhs10_
                lhs7_.f1 = rhs11_
                lhs8_[lhs9_] = rhs12_
                d_131_v51_ = d_131_v51_
                (d_76_globalState_).f3 = ((d_131_v51_)[default__.safeIndex(150, (d_131_v51_).length(0))]) <= (172)
                d_134_v54_: _dafny.Set
                d_134_v54_ = _dafny.Set({d_79_v2_})
                d_135_v55_: _dafny.Seq
                d_135_v55_ = _dafny.SeqWithoutIsStrInference([d_134_v54_])
                d_136_v56_: _dafny.Array
                nw27_ = _dafny.Array(None, 13)
                nw27_[int(0)] = d_134_v54_
                nw27_[int(1)] = d_134_v54_
                nw27_[int(2)] = d_134_v54_
                nw27_[int(3)] = (_dafny.Set({d_79_v2_, d_79_v2_, (D6_DC17(d_77_v0_, d_79_v2_, d_78_v1_, (d_80_v3_)[default__.safeIndex((d_131_v51_)[default__.safeIndex(150, (d_131_v51_).length(0))], len(d_80_v3_))], not(True))).cf24})) - (d_134_v54_)
                nw27_[int(4)] = (d_135_v55_)[default__.safeIndex((d_131_v51_)[default__.safeIndex(150, (d_131_v51_).length(0))], len(d_135_v55_))]
                nw27_[int(5)] = d_134_v54_
                nw27_[int(6)] = d_134_v54_
                nw27_[int(7)] = _dafny.Set({d_79_v2_, d_79_v2_})
                nw27_[int(8)] = _dafny.Set({d_79_v2_, d_79_v2_, d_79_v2_})
                nw27_[int(9)] = d_134_v54_
                nw27_[int(10)] = d_134_v54_
                nw27_[int(11)] = d_134_v54_
                nw27_[int(12)] = (_dafny.Set({d_79_v2_, d_79_v2_})).intersection(d_134_v54_)
                d_136_v56_ = nw27_
                index15_ = default__.safeIndex(456, (d_136_v56_).length(0))
                (d_136_v56_)[index15_] = d_134_v54_
                d_137_v57_: _dafny.Array
                nw28_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 21)
                d_137_v57_ = nw28_
                index16_ = default__.safeIndex(497, (d_137_v57_).length(0))
                (d_137_v57_)[index16_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vyygsh"))) + (d_77_v0_)
                d_138_v58_: D7
                d_138_v58_ = D7_DC19((d_77_v0_)[default__.safeIndex(d_78_v1_, len(d_77_v0_))], d_79_v2_)
                index17_ = default__.safeIndex(456, (d_136_v56_).length(0))
                index18_ = default__.safeIndex(497, (d_137_v57_).length(0))
                rhs13_ = (d_138_v58_).cf30
                rhs14_ = d_134_v54_
                rhs15_ = (d_77_v0_) + (d_77_v0_)
                lhs10_ = d_76_globalState_
                lhs11_ = d_136_v56_
                lhs12_ = default__.safeIndex(456, (d_136_v56_).length(0))
                lhs13_ = d_137_v57_
                lhs14_ = default__.safeIndex(497, (d_137_v57_).length(0))
                lhs10_.f1 = rhs13_
                lhs11_[lhs12_] = rhs14_
                lhs13_[lhs14_] = rhs15_
                d_139_v59_: C3
                nw29_ = C3()
                nw29_.ctor__(d_78_v1_, d_78_v1_)
                d_139_v59_ = nw29_
                d_139_v59_ = d_139_v59_
            d_140_v60_: C0
            nw30_ = C0()
            nw30_.ctor__(d_79_v2_)
            d_140_v60_ = nw30_
            d_140_v60_ = d_140_v60_
            d_77_v0_ = ((d_77_v0_) + (d_77_v0_)) + ((d_77_v0_) + (d_77_v0_))
            d_77_v0_ = d_77_v0_
            d_141_v61_: _dafny.Array
            nw31_ = _dafny.Array(_dafny.Map({}), 7)
            d_141_v61_ = nw31_
            d_142_v62_: _dafny.Seq
            d_142_v62_ = _dafny.SeqWithoutIsStrInference([d_78_v1_, d_78_v1_, d_78_v1_])
            d_143_v63_: C2
            nw32_ = C2()
            nw32_.ctor__(d_78_v1_, (d_142_v62_)[default__.safeIndex(d_78_v1_, len(d_142_v62_))], d_78_v1_, d_78_v1_)
            d_143_v63_ = nw32_
            index19_ = default__.safeIndex(49, (d_141_v61_).length(0))
            (d_141_v61_)[index19_] = _dafny.Map({d_143_v63_: (d_143_v63_).fm7(d_77_v0_, d_143_v63_.f10, d_76_globalState_)})
            d_144_v64_: _dafny.Map
            d_144_v64_ = _dafny.Map({d_143_v63_: d_143_v63_.f10})
            index20_ = default__.safeIndex(49, (d_141_v61_).length(0))
            (d_141_v61_)[index20_] = d_144_v64_
        elif True:
            d_78_v1_ = d_78_v1_
            d_145_v65_: _dafny.Map
            d_145_v65_ = _dafny.Map({d_79_v2_: _dafny.Map({(0) - (default__.fm0(d_81_v4_, d_76_globalState_)): d_78_v1_})})
            d_146_v66_: D1
            d_146_v66_ = D1_DC2(d_79_v2_)
            d_147_v67_: _dafny.Map
            d_147_v67_ = _dafny.Map({((d_82_v5_)[True] if (True) in (d_82_v5_) else d_78_v1_): d_78_v1_})
            d_145_v65_ = (d_145_v65_).set((d_146_v66_).cf2, d_147_v67_)
            d_148_v68_: _dafny.Array
            nw33_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_148_v68_ = nw33_
            d_149_v69_: C3
            nw34_ = C3()
            nw34_.ctor__(777, d_78_v1_)
            d_149_v69_ = nw34_
            d_150_v70_: _dafny.Array
            nw35_ = _dafny.Array(None, 4)
            nw35_[int(0)] = d_149_v69_
            nw35_[int(1)] = d_149_v69_
            nw35_[int(2)] = d_149_v69_
            nw35_[int(3)] = d_149_v69_
            d_150_v70_ = nw35_
            d_151_v71_: D6
            d_151_v71_ = D6_DC17(d_77_v0_, d_79_v2_, d_78_v1_, d_79_v2_, d_79_v2_)
            d_152_v72_: _dafny.Map
            d_152_v72_ = _dafny.Map({d_79_v2_: True})
            d_153_v73_: _dafny.Array
            nw36_ = _dafny.Array(None, 14)
            nw36_[int(0)] = (0) - (d_78_v1_)
            nw36_[int(1)] = d_78_v1_
            nw36_[int(2)] = d_78_v1_
            nw36_[int(3)] = d_78_v1_
            nw36_[int(4)] = d_78_v1_
            nw36_[int(5)] = (0) - ((_dafny.MultiSet([d_150_v70_, d_150_v70_, d_150_v70_])).cardinality)
            nw36_[int(6)] = (d_149_v69_).fm7((d_151_v71_).cf23, d_78_v1_, d_76_globalState_)
            nw36_[int(7)] = d_78_v1_
            nw36_[int(8)] = (d_149_v69_).fm7(d_77_v0_, d_78_v1_, d_76_globalState_)
            nw36_[int(9)] = 608
            nw36_[int(10)] = len(d_152_v72_)
            nw36_[int(11)] = (d_149_v69_).fm7(d_77_v0_, d_78_v1_, d_76_globalState_)
            nw36_[int(12)] = (0) - (len(d_82_v5_))
            nw36_[int(13)] = d_78_v1_
            d_153_v73_ = nw36_
            index21_ = default__.safeIndex(546, (d_148_v68_).length(0))
            (d_148_v68_)[index21_] = d_153_v73_
            d_154_v74_: _dafny.MultiSet
            d_154_v74_ = _dafny.MultiSet([d_79_v2_])
            index22_ = default__.safeIndex(546, (d_148_v68_).length(0))
            nw37_ = _dafny.Array(None, 10)
            nw37_[int(0)] = d_78_v1_
            nw37_[int(1)] = (d_78_v1_) - (591)
            nw37_[int(2)] = d_78_v1_
            nw37_[int(3)] = (d_78_v1_) * ((d_154_v74_).cardinality)
            nw37_[int(4)] = (d_78_v1_) - (len(d_152_v72_))
            nw37_[int(5)] = d_78_v1_
            nw37_[int(6)] = (d_78_v1_) * (-336)
            nw37_[int(7)] = d_78_v1_
            nw37_[int(8)] = default__.fm0((d_77_v0_)[default__.safeIndex(161, len(d_77_v0_))], d_76_globalState_)
            nw37_[int(9)] = d_78_v1_
            (d_148_v68_)[index22_] = nw37_
            d_155_v75_: D4
            d_155_v75_ = D4_DC12(d_78_v1_, _dafny.SeqWithoutIsStrInference([d_79_v2_]), not(not(d_79_v2_)))
            d_156_v76_: D7
            d_156_v76_ = D7_DC19(d_81_v4_, d_79_v2_)
            d_157_v77_: _dafny.Set
            d_157_v77_ = _dafny.Set({d_78_v1_, (0) - (((d_154_v74_)[d_79_v2_] if (d_79_v2_) in (d_154_v74_) else 179)), d_78_v1_})
            d_158_v78_: _dafny.Array
            nw38_ = _dafny.Array(None, 25)
            nw38_[int(0)] = d_79_v2_
            nw38_[int(1)] = d_79_v2_
            nw38_[int(2)] = d_79_v2_
            nw38_[int(3)] = (d_79_v2_) or (False)
            nw38_[int(4)] = (d_79_v2_) or (d_79_v2_)
            nw38_[int(5)] = (d_80_v3_) <= (d_80_v3_)
            nw38_[int(6)] = not (d_79_v2_) or (d_79_v2_)
            nw38_[int(7)] = d_79_v2_
            nw38_[int(8)] = d_79_v2_
            nw38_[int(9)] = not (d_79_v2_) or ((d_155_v75_).cf17)
            nw38_[int(10)] = d_79_v2_
            nw38_[int(11)] = d_79_v2_
            nw38_[int(12)] = d_79_v2_
            nw38_[int(13)] = d_79_v2_
            nw38_[int(14)] = d_79_v2_
            nw38_[int(15)] = d_79_v2_
            nw38_[int(16)] = ((0) - ((d_155_v75_).cf15)) > (d_78_v1_)
            nw38_[int(17)] = d_79_v2_
            nw38_[int(18)] = (d_156_v76_).cf30
            nw38_[int(19)] = not(not ((d_80_v3_)[default__.safeIndex((0) - (len(d_157_v77_)), len(d_80_v3_))]) or (d_79_v2_))
            nw38_[int(20)] = d_79_v2_
            nw38_[int(21)] = d_79_v2_
            nw38_[int(22)] = d_79_v2_
            nw38_[int(23)] = not (d_79_v2_) or (d_79_v2_)
            nw38_[int(24)] = not(d_79_v2_)
            d_158_v78_ = nw38_
            index23_ = default__.safeIndex(995, (d_158_v78_).length(0))
            (d_158_v78_)[index23_] = d_79_v2_
            index24_ = default__.safeIndex(995, (d_158_v78_).length(0))
            (d_158_v78_)[index24_] = d_79_v2_
            index25_ = default__.safeIndex(652, (d_153_v73_).length(0))
            (d_153_v73_)[index25_] = d_78_v1_
            index26_ = default__.safeIndex(652, (d_153_v73_).length(0))
            (d_153_v73_)[index26_] = ((d_78_v1_) + (d_78_v1_)) - (d_78_v1_)
        (d_76_globalState_).f3 = d_79_v2_
        (d_76_globalState_).f1 = d_79_v2_
        d_159_v79_: _dafny.Seq
        d_159_v79_ = _dafny.SeqWithoutIsStrInference([d_78_v1_])
        d_159_v79_ = d_159_v79_
        d_160_v80_: _dafny.Array
        nw39_ = _dafny.Array(None, 16)
        nw39_[int(0)] = d_81_v4_
        nw39_[int(1)] = (d_77_v0_)[default__.safeIndex(len(d_80_v3_), len(d_77_v0_))]
        nw39_[int(2)] = d_81_v4_
        nw39_[int(3)] = _dafny.CodePoint('a')
        nw39_[int(4)] = d_81_v4_
        nw39_[int(5)] = d_81_v4_
        nw39_[int(6)] = _dafny.CodePoint('f')
        nw39_[int(7)] = d_81_v4_
        nw39_[int(8)] = d_81_v4_
        nw39_[int(9)] = d_81_v4_
        nw39_[int(10)] = d_81_v4_
        nw39_[int(11)] = d_81_v4_
        nw39_[int(12)] = d_81_v4_
        nw39_[int(13)] = d_81_v4_
        nw39_[int(14)] = d_81_v4_
        nw39_[int(15)] = _dafny.CodePoint('j')
        d_160_v80_ = nw39_
        d_161_v81_: D9
        d_161_v81_ = D9_DC22(d_160_v80_)
        d_160_v80_ = (d_161_v81_).cf33
        hi1_ = (d_78_v1_ if d_79_v2_ else len(_dafny.Set({len(default__.fm16(d_76_globalState_)), d_78_v1_})))
        for d_162_i2_ in range(d_78_v1_, hi1_):
            if d_79_v2_:
                d_78_v1_ = default__.safeModuloInt(d_162_i2_, d_162_i2_)
                d_163_v82_: _dafny.Map
                d_163_v82_ = _dafny.Map({default__.safeDivisionInt((0) - (d_78_v1_), d_162_i2_): d_162_i2_})
                d_163_v82_ = d_163_v82_
                d_164_v83_: C3
                nw40_ = C3()
                nw40_.ctor__(d_162_i2_, d_162_i2_)
                d_164_v83_ = nw40_
                d_164_v83_ = d_164_v83_
                d_165_v84_: _dafny.Array
                nw41_ = _dafny.Array(int(0), 11)
                d_165_v84_ = nw41_
                d_166_v85_: _dafny.Array
                d_167_v86_: bool
                d_168_v87_: _dafny.Seq
                out0_: _dafny.Array
                out1_: bool
                out2_: _dafny.Seq
                out0_, out1_, out2_ = (d_164_v83_).m4(d_165_v84_, d_76_globalState_)
                d_166_v85_ = out0_
                d_167_v86_ = out1_
                d_168_v87_ = out2_
                d_169_v88_: _dafny.Array
                nw42_ = _dafny.Array(None, 4)
                nw42_[int(0)] = d_167_v86_
                nw42_[int(1)] = d_79_v2_
                nw42_[int(2)] = d_167_v86_
                nw42_[int(3)] = d_79_v2_
                d_169_v88_ = nw42_
                d_170_v89_: _dafny.Map
                d_170_v89_ = _dafny.Map({d_81_v4_: d_79_v2_})
                index27_ = default__.safeIndex(649, (d_169_v88_).length(0))
                (d_169_v88_)[index27_] = ((d_170_v89_)[d_81_v4_] if (d_81_v4_) in (d_170_v89_) else d_79_v2_)
                index28_ = default__.safeIndex(649, (d_169_v88_).length(0))
                rhs16_ = d_79_v2_
                rhs17_ = d_81_v4_
                rhs18_ = not (d_167_v86_) or (d_167_v86_)
                lhs15_ = d_169_v88_
                lhs16_ = default__.safeIndex(649, (d_169_v88_).length(0))
                lhs17_ = d_76_globalState_
                lhs15_[lhs16_] = rhs16_
                d_81_v4_ = rhs17_
                lhs17_.f3 = rhs18_
            elif True:
                d_171_v90_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 16)
                d_171_v90_ = nw43_
                index29_ = default__.safeIndex(532, (d_171_v90_).length(0))
                (d_171_v90_)[index29_] = default__.fm0(d_81_v4_, d_76_globalState_)
                d_172_v91_: _dafny.Seq
                d_172_v91_ = _dafny.SeqWithoutIsStrInference([d_77_v0_])
                index30_ = default__.safeIndex(532, (d_171_v90_).length(0))
                (d_171_v90_)[index30_] = (0) - ((len(d_172_v91_)) * (d_162_i2_))
                d_78_v1_ = (default__.safeDivisionInt(d_78_v1_, d_162_i2_) if (d_79_v2_ if True else default__.fm2(d_79_v2_, d_76_globalState_)) else d_78_v1_)
                index31_ = default__.safeIndex(532, (d_171_v90_).length(0))
                (d_171_v90_)[index31_] = ((0) - (d_162_i2_)) + (default__.fm0(d_81_v4_, d_76_globalState_))
                d_173_v92_: _dafny.Map
                d_173_v92_ = _dafny.Map({(d_171_v90_)[default__.safeIndex(532, (d_171_v90_).length(0))]: d_79_v2_})
                d_174_v93_: _dafny.Map
                d_174_v93_ = _dafny.Map({d_79_v2_: d_79_v2_})
                d_175_v94_: _dafny.Array
                nw44_ = _dafny.Array(None, 7)
                nw44_[int(0)] = True
                nw44_[int(1)] = not(d_79_v2_)
                nw44_[int(2)] = d_79_v2_
                nw44_[int(3)] = d_79_v2_
                nw44_[int(4)] = ((d_173_v92_)[d_78_v1_] if (d_78_v1_) in (d_173_v92_) else d_79_v2_)
                nw44_[int(5)] = ((d_174_v93_)[d_79_v2_] if (d_79_v2_) in (d_174_v93_) else d_79_v2_)
                nw44_[int(6)] = d_79_v2_
                d_175_v94_ = nw44_
                index32_ = default__.safeIndex(732, (d_175_v94_).length(0))
                (d_175_v94_)[index32_] = d_79_v2_
                index33_ = default__.safeIndex(732, (d_175_v94_).length(0))
                (d_175_v94_)[index33_] = (d_80_v3_)[default__.safeIndex(default__.safeModuloInt(d_78_v1_, d_162_i2_), len(d_80_v3_))]
                (d_76_globalState_).f1 = (d_79_v2_ if d_79_v2_ else (_dafny.Set({(d_175_v94_)[default__.safeIndex(732, (d_175_v94_).length(0))]})).ispropersubset(_dafny.Set({(d_175_v94_)[default__.safeIndex(732, (d_175_v94_).length(0))], True})))
            source3_ = D8_DC21(default__.safeModuloInt(-625, d_162_i2_))
            if source3_.is_DC21:
                d_176___mcc_h0_ = source3_.cf32
                d_177_cf32_ = d_176___mcc_h0_
                d_178_v95_: _dafny.Array
                def lambda5_(d_179_v79_):
                    def lambda6_(d_180_i3_):
                        return d_179_v79_

                    return lambda6_

                init2_ = lambda5_(d_159_v79_)
                nw45_ = _dafny.Array(None, 11)
                for i0_2_ in range(nw45_.length(0)):
                    nw45_[i0_2_] = init2_(i0_2_)
                d_178_v95_ = nw45_
                index34_ = default__.safeIndex(443, (d_178_v95_).length(0))
                (d_178_v95_)[index34_] = d_159_v79_
                index35_ = default__.safeIndex(443, (d_178_v95_).length(0))
                (d_178_v95_)[index35_] = (d_159_v79_) + (d_159_v79_)
                (d_116_v36_).m1(d_76_globalState_)
                d_181_v97_: _dafny.MultiSet
                d_181_v97_ = _dafny.MultiSet([d_78_v1_, d_162_i2_])
                d_182_v98_: C2
                nw46_ = C2()
                def iife13_():
                    coll13_ = _dafny.Map()
                    compr_13_: int
                    for compr_13_ in (d_181_v97_).Elements:
                        d_183_v96_: int = compr_13_
                        if (d_183_v96_) in (d_181_v97_):
                            coll13_[default__.safeDivisionInt(d_183_v96_, (0) - (d_162_i2_))] = d_79_v2_
                    return _dafny.Map(coll13_)
                nw46_.ctor__(len(iife13_()
                ), (0) - (d_177_cf32_), d_177_cf32_, d_177_cf32_)
                d_182_v98_ = nw46_
                d_184_v99_: C0
                nw47_ = C0()
                nw47_.ctor__(d_79_v2_)
                d_184_v99_ = nw47_
                d_185_v100_: _dafny.Seq
                d_185_v100_ = _dafny.SeqWithoutIsStrInference([d_184_v99_, d_184_v99_])
                d_186_v101_: D6
                d_186_v101_ = D6_DC17(d_77_v0_, d_79_v2_, -505, True, d_79_v2_)
                d_82_v5_ = (d_82_v5_).set((len(d_185_v100_)) <= ((d_182_v98_).fm7((d_186_v101_).cf23, d_78_v1_, d_76_globalState_)), d_162_i2_)
            elif True:
                d_187___mcc_h1_ = source3_.cf31
                d_188_cf31_ = d_187___mcc_h1_
                d_189_v102_: _dafny.MultiSet
                d_189_v102_ = _dafny.MultiSet([d_81_v4_, d_81_v4_])
                d_190_v103_: D4
                d_190_v103_ = D4_DC12(d_78_v1_, d_80_v3_, d_79_v2_)
                d_191_v104_: _dafny.Map
                d_191_v104_ = _dafny.Map({d_79_v2_: (d_82_v5_).set(False, d_162_i2_)})
                d_192_v105_: _dafny.Array
                nw48_ = _dafny.Array(None, 27)
                nw48_[int(0)] = d_82_v5_
                nw48_[int(1)] = d_82_v5_
                nw48_[int(2)] = _dafny.Map({d_79_v2_: 797})
                nw48_[int(3)] = _dafny.Map({d_79_v2_: d_78_v1_})
                nw48_[int(4)] = d_82_v5_
                nw48_[int(5)] = _dafny.Map({d_79_v2_: d_162_i2_})
                nw48_[int(6)] = d_82_v5_
                nw48_[int(7)] = d_82_v5_
                nw48_[int(8)] = _dafny.Map({True: d_162_i2_})
                nw48_[int(9)] = _dafny.Map({d_79_v2_: d_78_v1_})
                nw48_[int(10)] = _dafny.Map({False: d_78_v1_})
                nw48_[int(11)] = default__.fm21(len(d_82_v5_), ((d_189_v102_)[d_81_v4_] if (d_81_v4_) in (d_189_v102_) else len(_dafny.SeqWithoutIsStrInference([(d_190_v103_).cf15]))), d_76_globalState_)
                nw48_[int(12)] = d_82_v5_
                nw48_[int(13)] = default__.fm21((0) - (d_162_i2_), d_78_v1_, d_76_globalState_)
                nw48_[int(14)] = ((d_191_v104_)[d_79_v2_] if (d_79_v2_) in (d_191_v104_) else d_82_v5_)
                nw48_[int(15)] = d_82_v5_
                nw48_[int(16)] = d_82_v5_
                nw48_[int(17)] = d_82_v5_
                nw48_[int(18)] = d_82_v5_
                nw48_[int(19)] = d_82_v5_
                nw48_[int(20)] = d_82_v5_
                nw48_[int(21)] = d_82_v5_
                nw48_[int(22)] = d_82_v5_
                nw48_[int(23)] = d_82_v5_
                nw48_[int(24)] = d_82_v5_
                nw48_[int(25)] = d_82_v5_
                nw48_[int(26)] = _dafny.Map({d_79_v2_: d_78_v1_})
                d_192_v105_ = nw48_
                d_193_v106_: _dafny.Array
                nw49_ = _dafny.Array(None, 11)
                nw49_[int(0)] = d_192_v105_
                nw49_[int(1)] = d_192_v105_
                nw49_[int(2)] = d_192_v105_
                nw49_[int(3)] = d_192_v105_
                nw49_[int(4)] = d_192_v105_
                nw49_[int(5)] = d_192_v105_
                nw49_[int(6)] = d_192_v105_
                nw49_[int(7)] = d_192_v105_
                nw49_[int(8)] = d_192_v105_
                nw49_[int(9)] = d_192_v105_
                nw49_[int(10)] = d_192_v105_
                d_193_v106_ = nw49_
                index36_ = default__.safeIndex(491, (d_193_v106_).length(0))
                (d_193_v106_)[index36_] = d_192_v105_
                index37_ = default__.safeIndex(491, (d_193_v106_).length(0))
                (d_193_v106_)[index37_] = d_192_v105_
                d_78_v1_ = len(d_77_v0_)
                d_194_v107_: _dafny.Array
                def lambda7_(d_195_i4_):
                    return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_196_i5_ in range(default__.abs(117))])

                init3_ = lambda7_
                nw50_ = _dafny.Array(None, 24)
                for i0_3_ in range(nw50_.length(0)):
                    nw50_[i0_3_] = init3_(i0_3_)
                d_194_v107_ = nw50_
                index38_ = default__.safeIndex(158, (d_194_v107_).length(0))
                (d_194_v107_)[index38_] = d_77_v0_
                d_197_v108_: _dafny.Array
                def lambda8_(d_198_i6_):
                    return (d_198_i6_) * (281)

                init4_ = lambda8_
                nw51_ = _dafny.Array(None, 29)
                for i0_4_ in range(nw51_.length(0)):
                    nw51_[i0_4_] = init4_(i0_4_)
                d_197_v108_ = nw51_
                index39_ = default__.safeIndex(352, (d_197_v108_).length(0))
                (d_197_v108_)[index39_] = (d_78_v1_ if d_79_v2_ else d_78_v1_)
                index40_ = default__.safeIndex(158, (d_194_v107_).length(0))
                index41_ = default__.safeIndex(352, (d_197_v108_).length(0))
                rhs19_ = d_79_v2_
                rhs20_ = d_77_v0_
                rhs21_ = d_78_v1_
                rhs22_ = d_79_v2_
                lhs18_ = d_76_globalState_
                lhs19_ = d_194_v107_
                lhs20_ = default__.safeIndex(158, (d_194_v107_).length(0))
                lhs21_ = d_197_v108_
                lhs22_ = default__.safeIndex(352, (d_197_v108_).length(0))
                lhs23_ = d_76_globalState_
                lhs18_.f1 = rhs19_
                lhs19_[lhs20_] = rhs20_
                lhs21_[lhs22_] = rhs21_
                lhs23_.f3 = rhs22_
                d_199_v109_: T0
                nw52_ = C2()
                nw52_.ctor__(d_162_i2_, d_78_v1_, d_162_i2_, (d_197_v108_)[default__.safeIndex(352, (d_197_v108_).length(0))])
                d_199_v109_ = nw52_
                d_200_v110_: _dafny.Seq
                d_200_v110_ = _dafny.SeqWithoutIsStrInference([d_199_v109_])
                index42_ = default__.safeIndex(352, (d_197_v108_).length(0))
                (d_197_v108_)[index42_] = (0) - (((d_197_v108_)[default__.safeIndex(352, (d_197_v108_).length(0))] if d_79_v2_ else (len(d_200_v110_) if d_79_v2_ else (d_199_v109_).f8)))
            (d_76_globalState_).f1 = (d_80_v3_)[default__.safeIndex(d_78_v1_, len(d_80_v3_))]
            if not((d_78_v1_) > (d_162_i2_)):
                d_201_v111_: _dafny.Map
                d_201_v111_ = _dafny.Map({d_162_i2_: d_162_i2_})
                d_202_v112_: _dafny.Set
                d_202_v112_ = _dafny.Set({True})
                d_203_v113_: _dafny.Set
                d_203_v113_ = _dafny.Set({len(d_202_v112_), d_162_i2_, d_162_i2_})
                d_204_v114_: _dafny.Map
                d_204_v114_ = _dafny.Map({(_dafny.MultiSet([d_79_v2_])).set(d_79_v2_, default__.abs(len(d_203_v113_))): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rxrjsdrs")))})
                d_201_v111_ = default__.fm22(d_78_v1_, (d_162_i2_) - (len(d_204_v114_)), True, d_159_v79_, d_76_globalState_)
                d_205_v115_: _dafny.Array
                def lambda9_(d_206_v2_):
                    def lambda10_(d_207_i7_):
                        return d_206_v2_

                    return lambda10_

                init5_ = lambda9_(d_79_v2_)
                nw53_ = _dafny.Array(None, 14)
                for i0_5_ in range(nw53_.length(0)):
                    nw53_[i0_5_] = init5_(i0_5_)
                d_205_v115_ = nw53_
                d_208_v116_: _dafny.MultiSet
                d_208_v116_ = _dafny.MultiSet([d_205_v115_, d_205_v115_, d_205_v115_])
                d_209_v117_: _dafny.Array
                nw54_ = _dafny.Array(int(0), 27)
                d_209_v117_ = nw54_
                d_210_v118_: _dafny.Map
                d_210_v118_ = _dafny.Map({d_208_v116_: _dafny.Map({default__.fm11(d_79_v2_, d_162_i2_, d_76_globalState_): d_209_v117_})})
                d_211_v119_: _dafny.Map
                d_211_v119_ = _dafny.Map({d_77_v0_: d_209_v117_})
                d_210_v118_ = (d_210_v118_).set((_dafny.MultiSet([d_205_v115_, d_205_v115_])) - (d_208_v116_), d_211_v119_)
                d_212_v120_: D0
                d_212_v120_ = D0_DC0(d_209_v117_)
                d_213_v121_: _dafny.Map
                d_213_v121_ = _dafny.Map({d_78_v1_: _dafny.Map({338: d_79_v2_})})
                d_214_v122_: _dafny.Map
                d_214_v122_ = _dafny.Map({d_78_v1_: d_79_v2_})
                d_215_v123_: C0
                nw55_ = C0()
                nw55_.ctor__(d_79_v2_)
                d_215_v123_ = nw55_
                d_216_v124_: _dafny.Map
                d_216_v124_ = _dafny.Map({((d_213_v121_)[d_162_i2_] if (d_162_i2_) in (d_213_v121_) else d_214_v122_): d_215_v123_})
                d_217_v125_: T0
                nw56_ = C3()
                nw56_.ctor__(len(d_216_v124_), d_78_v1_)
                d_217_v125_ = nw56_
                d_218_v126_: _dafny.Map
                d_218_v126_ = _dafny.Map({d_162_i2_: (d_217_v125_ if d_79_v2_ else d_217_v125_)})
                rhs23_ = d_212_v120_
                rhs24_ = (_dafny.Map({((d_82_v5_)[(d_215_v123_).f11] if ((d_215_v123_).f11) in (d_82_v5_) else (0) - ((d_217_v125_).f8)): d_217_v125_})) | ((d_218_v126_) | (_dafny.Map({d_78_v1_: d_217_v125_})))
                rhs25_ = d_79_v2_
                rhs26_ = (d_217_v125_).fm7(d_77_v0_, (len(d_77_v0_)) * ((d_217_v125_).f7), d_76_globalState_)
                lhs24_ = d_76_globalState_
                d_212_v120_ = rhs23_
                d_218_v126_ = rhs24_
                lhs24_.f3 = rhs25_
                d_78_v1_ = rhs26_
                d_214_v122_ = (d_214_v122_).set(len(d_77_v0_), not(d_79_v2_))
                d_219_v127_: _dafny.MultiSet
                d_219_v127_ = _dafny.MultiSet([False])
                d_220_v128_: C1
                nw57_ = C1()
                nw57_.ctor__(((d_219_v127_).cardinality) != (d_78_v1_), d_78_v1_, 220, (d_217_v125_).f7)
                d_220_v128_ = nw57_
            elif True:
                d_79_v2_ = d_79_v2_
                d_221_v129_: _dafny.Array
                nw58_ = _dafny.Array(int(0), 25)
                d_221_v129_ = nw58_
                index43_ = default__.safeIndex(915, (d_221_v129_).length(0))
                (d_221_v129_)[index43_] = d_78_v1_
                index44_ = default__.safeIndex(915, (d_221_v129_).length(0))
                (d_221_v129_)[index44_] = (0) - (d_162_i2_)
                d_222_v130_: _dafny.Array
                nw59_ = _dafny.Array(False, 5)
                d_222_v130_ = nw59_
                index45_ = default__.safeIndex(457, (d_222_v130_).length(0))
                (d_222_v130_)[index45_] = False
                index46_ = default__.safeIndex(457, (d_222_v130_).length(0))
                rhs27_ = (d_80_v3_)[default__.safeIndex(default__.safeDivisionInt(len(d_159_v79_), d_162_i2_), len(d_80_v3_))]
                rhs28_ = ((d_221_v129_)[default__.safeIndex(915, (d_221_v129_).length(0))]) >= (((d_221_v129_)[default__.safeIndex(915, (d_221_v129_).length(0))]) - (d_78_v1_))
                lhs25_ = d_222_v130_
                lhs26_ = default__.safeIndex(457, (d_222_v130_).length(0))
                lhs27_ = d_76_globalState_
                lhs25_[lhs26_] = rhs27_
                lhs27_.f3 = rhs28_
                index47_ = default__.safeIndex(457, (d_222_v130_).length(0))
                (d_222_v130_)[index47_] = (d_222_v130_)[default__.safeIndex(457, (d_222_v130_).length(0))]
                d_223_v131_: _dafny.Map
                d_223_v131_ = _dafny.Map({d_78_v1_: (d_222_v130_)[default__.safeIndex(457, (d_222_v130_).length(0))]})
                (d_116_v36_).m2(d_223_v131_, (d_80_v3_) + (d_80_v3_), not((d_222_v130_)[default__.safeIndex(457, (d_222_v130_).length(0))]), d_76_globalState_)
        d_78_v1_ = (0) - ((d_78_v1_) * ((d_78_v1_) + (d_78_v1_)))
        d_79_v2_ = d_79_v2_
        d_224_v132_: _dafny.Set
        d_224_v132_ = _dafny.Set({(d_77_v0_)[default__.safeIndex(d_78_v1_, len(d_77_v0_))]})
        d_225_v134_: _dafny.MultiSet
        def iife14_():
            coll14_ = _dafny.Set()
            compr_14_: str
            for compr_14_ in (d_224_v132_).Elements:
                d_226_v133_: str = compr_14_
                if (d_226_v133_) in (d_224_v132_):
                    coll14_ = coll14_.union(_dafny.Set([d_226_v133_]))
            return _dafny.Set(coll14_)
        d_225_v134_ = _dafny.MultiSet([d_224_v132_, iife14_()
        , d_224_v132_])
        d_227_i8_: int
        d_227_i8_ = 0
        with _dafny.label("0"):
            while (d_225_v134_).issubset(d_225_v134_):
                with _dafny.c_label("0"):
                    if (d_227_i8_) >= (100):
                        raise _dafny.Break("0")
                    d_227_i8_ = (d_227_i8_) + (1)
                    d_228_v135_: _dafny.Array
                    nw60_ = _dafny.Array(_dafny.Set({}), 17)
                    d_228_v135_ = nw60_
                    d_229_v136_: T0
                    nw61_ = C1()
                    nw61_.ctor__(d_79_v2_, d_78_v1_, d_78_v1_, d_78_v1_)
                    d_229_v136_ = nw61_
                    d_230_v137_: D2
                    d_230_v137_ = D2_DC5(len(_dafny.Set({d_79_v2_})), d_77_v0_, d_229_v136_)
                    d_231_v138_: D2
                    d_231_v138_ = D2_DC6(d_230_v137_)
                    d_232_v139_: _dafny.Map
                    d_232_v139_ = _dafny.Map({d_81_v4_: d_231_v138_})
                    d_233_v140_: _dafny.Set
                    d_233_v140_ = _dafny.Set({d_232_v139_})
                    index48_ = default__.safeIndex(378, (d_228_v135_).length(0))
                    (d_228_v135_)[index48_] = d_233_v140_
                    index49_ = default__.safeIndex(378, (d_228_v135_).length(0))
                    (d_228_v135_)[index49_] = d_233_v140_
                    d_234_v141_: D6
                    d_234_v141_ = D6_DC17(d_77_v0_, d_79_v2_, d_78_v1_, d_79_v2_, d_79_v2_)
                    d_235_v142_: C1
                    nw62_ = C1()
                    nw62_.ctor__(not((d_234_v141_).cf26), (d_229_v136_).f7, d_78_v1_, ((d_229_v136_).f7) + ((d_229_v136_).fm7(d_77_v0_, (d_229_v136_).f7, d_76_globalState_)))
                    d_235_v142_ = nw62_
                    d_235_v142_ = d_235_v142_
                    d_236_v143_: _dafny.Array
                    def lambda11_(d_237_v136_):
                        def lambda12_(d_238_i9_):
                            return (d_238_i9_) + ((d_237_v136_).f7)

                        return lambda12_

                    init6_ = lambda11_(d_229_v136_)
                    nw63_ = _dafny.Array(None, 2)
                    for i0_6_ in range(nw63_.length(0)):
                        nw63_[i0_6_] = init6_(i0_6_)
                    d_236_v143_ = nw63_
                    index50_ = default__.safeIndex(692, (d_236_v143_).length(0))
                    (d_236_v143_)[index50_] = (d_229_v136_).f7
                    index51_ = default__.safeIndex(692, (d_236_v143_).length(0))
                    (d_236_v143_)[index51_] = d_78_v1_
                    d_239_v144_: _dafny.Map
                    d_239_v144_ = _dafny.Map({d_79_v2_: d_79_v2_})
                    index52_ = default__.safeIndex(692, (d_236_v143_).length(0))
                    (d_236_v143_)[index52_] = (len(d_239_v144_)) + ((d_78_v1_) * ((d_235_v142_).fm7(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qyb")), len(d_159_v79_), d_76_globalState_)))
                    pass
            pass
        d_240_v145_: _dafny.Array
        def lambda13_(d_241_i10_):
            return D1_DC2(False)

        init7_ = lambda13_
        nw64_ = _dafny.Array(None, 3)
        for i0_7_ in range(nw64_.length(0)):
            nw64_[i0_7_] = init7_(i0_7_)
        d_240_v145_ = nw64_
        d_242_v146_: D1
        d_242_v146_ = D1_DC2(False)
        index53_ = default__.safeIndex(11, (d_240_v145_).length(0))
        (d_240_v145_)[index53_] = d_242_v146_
        index54_ = default__.safeIndex(11, (d_240_v145_).length(0))
        (d_240_v145_)[index54_] = D1_DC2(d_79_v2_)
        if d_79_v2_:
            (d_76_globalState_).f3 = not ((d_79_v2_) or (d_79_v2_)) or (d_79_v2_)
            d_243_v147_: _dafny.Map
            d_243_v147_ = _dafny.Map({d_78_v1_: d_79_v2_})
            d_244_v148_: _dafny.Map
            d_244_v148_ = _dafny.Map({d_81_v4_: d_80_v3_})
            d_245_v149_: _dafny.Array
            nw65_ = _dafny.Array(False, 29)
            d_245_v149_ = nw65_
            d_246_v150_: _dafny.Set
            d_246_v150_ = _dafny.Set({d_245_v149_, d_245_v149_, d_245_v149_})
            (d_116_v36_).m2(d_243_v147_, ((d_244_v148_)[d_81_v4_] if (d_81_v4_) in (d_244_v148_) else d_80_v3_), (d_246_v150_).ispropersubset(_dafny.Set({d_245_v149_, d_245_v149_})), d_76_globalState_)
            d_247_v151_: D6
            d_247_v151_ = D6_DC17(_dafny.SeqWithoutIsStrInference([d_81_v4_ for d_248_i11_ in range(default__.abs(833))]), d_79_v2_, len(d_77_v0_), d_79_v2_, d_79_v2_)
            d_247_v151_ = d_247_v151_
            (d_116_v36_).m1(d_76_globalState_)
            d_249_v152_: C4
            nw66_ = C4()
            nw66_.ctor__((d_116_v36_).f6)
            d_249_v152_ = nw66_
        elif True:
            default__.m0(d_76_globalState_)
            (d_76_globalState_).f1 = d_79_v2_
            d_250_v153_: _dafny.Map
            d_250_v153_ = _dafny.Map({d_82_v5_: not(d_79_v2_)})
            d_78_v1_ = default__.safeDivisionInt((0) - (d_78_v1_), len(d_250_v153_))
            d_251_v154_: _dafny.Array
            nw67_ = _dafny.Array(None, 2)
            d_251_v154_ = nw67_
            d_251_v154_ = d_251_v154_
            (d_116_v36_).m1(d_76_globalState_)
        _dafny.print(_dafny.string_of((d_76_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_76_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_76_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_76_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_77_v0_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_78_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_79_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_80_v3_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_81_v4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_82_v5_) == (_dafny.Map({True: 2}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_83_v6_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, True]): 414}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v34_) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('p'): True, _dafny.CodePoint('q'): True, _dafny.CodePoint('i'): True, _dafny.CodePoint('t'): True, _dafny.CodePoint('b'): True, _dafny.CodePoint('m'): True}): _dafny.Map({-958: _dafny.CodePoint('q')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_115_v35_).cf28) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('p'): True, _dafny.CodePoint('q'): True, _dafny.CodePoint('i'): True, _dafny.CodePoint('t'): True, _dafny.CodePoint('b'): True, _dafny.CodePoint('m'): True}): _dafny.Map({-958: _dafny.CodePoint('q')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_116_v36_).f6) == (_dafny.Map({_dafny.Map({_dafny.CodePoint('p'): True, _dafny.CodePoint('q'): True, _dafny.CodePoint('i'): True, _dafny.CodePoint('t'): True, _dafny.CodePoint('b'): True, _dafny.CodePoint('m'): True}): _dafny.Map({-958: _dafny.CodePoint('q')})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v79_) == (_dafny.SeqWithoutIsStrInference([-958]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_160_v80_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_161_v81_).cf33)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_224_v132_) == (_dafny.Set({_dafny.CodePoint('p')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_225_v134_) == (_dafny.MultiSet([_dafny.Set({_dafny.CodePoint('p')}), _dafny.Set({_dafny.CodePoint('p')}), _dafny.Set({_dafny.CodePoint('p')})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_227_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_v145_)[0]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_v145_)[1]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_240_v145_)[2]).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v146_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.CodePoint('D'))
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
        return lambda: D1_DC3(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
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
        return lambda: D2_DC5(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC5(D2, NamedTuple('DC5', [('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf5)}, {self.cf6.VerbatimString(True)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf4 == __o.cf4
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
        return lambda: D3_DC8(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC8(D3, NamedTuple('DC8', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D4_DC11(int(0), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)

class D4_DC11(D4, NamedTuple('DC11', [('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC15(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC15(D5, NamedTuple('DC15', [('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({self.cf20.VerbatimString(True)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)

class D6_DC17(D6, NamedTuple('DC17', [('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({self.cf23.VerbatimString(True)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC16(D6, NamedTuple('DC16', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC19(_dafny.CodePoint('D'), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D7_DC19)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC19(D7, NamedTuple('DC19', [('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC19({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC19) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC21(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)

class D8_DC21(D8, NamedTuple('DC21', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC23(int(0), int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)

class D9_DC23(D9, NamedTuple('DC23', [('cf34', Any), ('cf35', Any), ('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35 and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC22(D9, NamedTuple('DC22', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class GlobalState:
    def  __init__(self):
        self.f1: bool = False
        self.f3: bool = False
        self._f0: bool = False
        self._f2: int = int(0)
        self._f4: bool = False
        self._f5: str = _dafny.CodePoint('D')
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f4(self):
        return self._f4
    @property
    def f5(self):
        return self._f5

class C0:
    def  __init__(self):
        self._f11: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f11):
        (self)._f11 = f11

    @property
    def f11(self):
        return self._f11

class C1(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        self.f13: int = int(0)
        self._f12: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f12, f13, f7, f8):
        (self)._f12 = f12
        (self).f13 = f13
        (self)._f7 = f7
        (self)._f8 = f8

    def fm6(self, p0, p1, globalState):
        return ((_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(self).f12]) for d_252_i0_ in range(default__.abs(175))]))])})) - ((D3_DC7(_dafny.Set({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f12: (self).f12})), (self).f8]))}))).cf9)) | ((_dafny.Set({_dafny.MultiSet([-398]), _dafny.MultiSet([(self).f8, (self).f8]), _dafny.MultiSet([(self).f7]), _dafny.MultiSet([639]), _dafny.MultiSet([(self).f7])})).intersection(_dafny.Set({_dafny.MultiSet([len(_dafny.Map({(self).f12: _dafny.CodePoint('v')})), (self).f7, len(_dafny.SeqWithoutIsStrInference([self.f13 for d_253_i1_ in range(default__.abs(239))]))])})))

    def fm7(self, p0, p1, globalState):
        return (self).f7

    def m6(self, p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: D3 = D3.default()()
        hi2_ = (self).f7
        for d_254_i0_ in range(p0, hi2_):
            (globalState).f1 = ((self).f7) < (p0)
            d_255_v0_: C0
            nw68_ = C0()
            nw68_.ctor__((self).f12)
            d_255_v0_ = nw68_
            d_256_v1_: _dafny.Array
            nw69_ = _dafny.Array(int(0), 12)
            d_256_v1_ = nw69_
            index55_ = default__.safeIndex(23, (d_256_v1_).length(0))
            (d_256_v1_)[index55_] = default__.safeModuloInt((self).f7, (self).f8)
            d_257_v2_: _dafny.Array
            nw70_ = _dafny.Array(D3.default()(), 26)
            d_257_v2_ = nw70_
            d_258_v3_: _dafny.MultiSet
            d_258_v3_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpakoftbq"))), self.f13, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_259_i1_ in range(default__.abs(307))]))])
            d_260_v4_: _dafny.Seq
            d_260_v4_ = _dafny.SeqWithoutIsStrInference([d_258_v3_])
            d_261_v6_: D3
            def iife15_():
                coll15_ = _dafny.Set()
                compr_15_: _dafny.MultiSet
                for compr_15_ in (d_260_v4_).Elements:
                    d_262_v5_: _dafny.MultiSet = compr_15_
                    if (d_262_v5_) in (d_260_v4_):
                        coll15_ = coll15_.union(_dafny.Set([d_262_v5_]))
                return _dafny.Set(coll15_)
            d_261_v6_ = D3_DC7(iife15_()
)
            index56_ = default__.safeIndex(597, (d_257_v2_).length(0))
            (d_257_v2_)[index56_] = d_261_v6_
            d_263_v7_: _dafny.MultiSet
            d_263_v7_ = _dafny.MultiSet([(self).f12, (d_255_v0_).f11, (self).f12])
            d_264_v8_: _dafny.Map
            d_264_v8_ = _dafny.Map({440: d_263_v7_})
            d_265_v10_: D4
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(674, 811):
                    d_266_v9_: int = compr_16_
                    if ((674) <= (d_266_v9_)) and ((d_266_v9_) < (811)):
                        coll16_[(d_266_v9_) + (68)] = d_254_i0_
                return _dafny.Map(coll16_)
            d_265_v10_ = D4_DC10(iife16_()
)
            d_267_v11_: _dafny.Set
            d_267_v11_ = _dafny.Set({(self).f12})
            index57_ = default__.safeIndex(23, (d_256_v1_).length(0))
            index58_ = default__.safeIndex(597, (d_257_v2_).length(0))
            rhs29_ = (((d_264_v8_)[len((d_265_v10_).cf12)] if (len((d_265_v10_).cf12)) in (d_264_v8_) else d_263_v7_)).cardinality
            rhs30_ = default__.fm12((not((d_255_v0_).f11)) or (True), (d_255_v0_).f11, default__.fm3((self).f12, globalState), globalState)
            rhs31_ = (((self).f7) != (len(d_267_v11_)) if (self).f12 else (p0) == (len(_dafny.Map({(d_255_v0_).f11: (self).f8}))))
            lhs28_ = d_256_v1_
            lhs29_ = default__.safeIndex(23, (d_256_v1_).length(0))
            lhs30_ = d_257_v2_
            lhs31_ = default__.safeIndex(597, (d_257_v2_).length(0))
            lhs32_ = globalState
            lhs28_[lhs29_] = rhs29_
            lhs30_[lhs31_] = rhs30_
            lhs32_.f3 = rhs31_
            d_256_v1_ = d_256_v1_
        d_268_v12_: _dafny.Seq
        d_268_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhflilcl"))
        d_268_v12_ = d_268_v12_
        d_269_v13_: C0
        nw71_ = C0()
        nw71_.ctor__((self).f12)
        d_269_v13_ = nw71_
        d_270_v14_: D1
        d_270_v14_ = D1_DC2((d_269_v13_).f11)
        d_271_v15_: C0
        nw72_ = C0()
        nw72_.ctor__(((self).f12 if (d_270_v14_).cf2 else (self).f12))
        d_271_v15_ = nw72_
        d_272_v16_: _dafny.Set
        d_272_v16_ = _dafny.Set({default__.safeModuloInt(self.f13, self.f13), self.f13, (self).f7, (self).f7})
        d_272_v16_ = d_272_v16_
        (self).f13 = default__.safeDivisionInt(p0, (self).fm7(d_268_v12_, (self).f7, globalState))
        d_273_v17_: _dafny.Seq
        d_273_v17_ = _dafny.SeqWithoutIsStrInference([(d_269_v13_).f11])
        r0 = (d_273_v17_) + (_dafny.SeqWithoutIsStrInference([(self).f12]))
        d_274_v18_: D3
        d_274_v18_ = D3_DC8(((self).f8) - (p0))
        r1 = d_274_v18_
        return r0, r1

    @property
    def f12(self):
        return self._f12

class C2(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        self.f10: int = int(0)
        self._f9: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f9, f10, f7, f8):
        (self)._f9 = f9
        (self).f10 = f10
        (self)._f7 = f7
        (self)._f8 = f8

    def fm6(self, p0, p1, globalState):
        return (_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, False, not(False), True]))]), _dafny.MultiSet([(self).f8])})) - (_dafny.Set({_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ql")))})): (self).f8}))]), _dafny.MultiSet([(self).f8, (self).f7, (self).f9]), _dafny.MultiSet([846, (self).f8]), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f9}) for d_275_i0_ in range(default__.abs(23))]))])}))

    def fm7(self, p0, p1, globalState):
        return (self).f8

    def fm9(self, p0, p1, globalState):
        return ((D1_DC2(not(False))).cf2) and ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "see"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mu"))))

    def fm10(self, p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([False])).issubset((_dafny.MultiSet([True, False])).intersection(_dafny.MultiSet([True])))

    def m5(self, p0, p1, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (globalState).f1 = ((self.f10) * ((self).f9)) >= (self.f10)
        d_276_i0_: int
        d_276_i0_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_276_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_276_i0_ = (d_276_i0_) + (1)
                    d_277_v0_: bool
                    d_277_v0_ = True
                    if not (d_277_v0_) or (True):
                        (self).f10 = (self).f7
                        d_278_v1_: _dafny.Array
                        def lambda14_(d_279_v0_):
                            def lambda15_(d_280_i1_):
                                return d_279_v0_

                            return lambda15_

                        init8_ = lambda14_(d_277_v0_)
                        nw73_ = _dafny.Array(None, 8)
                        for i0_8_ in range(nw73_.length(0)):
                            nw73_[i0_8_] = init8_(i0_8_)
                        d_278_v1_ = nw73_
                        d_278_v1_ = d_278_v1_
                        d_281_v2_: str
                        d_281_v2_ = _dafny.CodePoint('e')
                        d_282_v3_: C0
                        nw74_ = C0()
                        nw74_.ctor__((self).fm10(not(d_277_v0_), d_277_v0_, d_277_v0_, d_281_v2_, globalState))
                        d_282_v3_ = nw74_
                        d_283_v4_: _dafny.Seq
                        d_283_v4_ = _dafny.SeqWithoutIsStrInference([d_277_v0_])
                        d_284_v5_: D1
                        d_284_v5_ = D1_DC3(d_281_v2_)
                        d_285_v6_: _dafny.Map
                        d_285_v6_ = _dafny.Map({default__.safeModuloInt((self).f7, (_dafny.MultiSet(d_283_v4_)).cardinality): d_284_v5_})
                        d_285_v6_ = (_dafny.Map({561: d_284_v5_})) | (d_285_v6_)
                        d_286_v7_: _dafny.Array
                        nw75_ = _dafny.Array(int(0), 14)
                        d_286_v7_ = nw75_
                        d_286_v7_ = d_286_v7_
                    elif True:
                        d_287_v8_: _dafny.Set
                        d_287_v8_ = _dafny.Set({(self).f8, self.f10})
                        d_287_v8_ = d_287_v8_
                        d_288_v9_: D2
                        d_288_v9_ = D2_DC4((self).f7)
                        (self).f10 = (d_288_v9_).cf4
                        d_289_v10_: _dafny.Set
                        d_289_v10_ = _dafny.Set({d_277_v0_, default__.fm2(d_277_v0_, globalState), d_277_v0_, d_277_v0_})
                        d_290_v11_: _dafny.Map
                        d_290_v11_ = _dafny.Map({(d_277_v0_ if not(d_277_v0_) else d_277_v0_): len(d_289_v10_)})
                        d_291_v12_: str
                        d_291_v12_ = _dafny.CodePoint('w')
                        d_292_v13_: _dafny.Array
                        def lambda16_(d_293_v12_):
                            def lambda17_(d_294_i2_):
                                return d_293_v12_

                            return lambda17_

                        init9_ = lambda16_(d_291_v12_)
                        nw76_ = _dafny.Array(None, 1)
                        for i0_9_ in range(nw76_.length(0)):
                            nw76_[i0_9_] = init9_(i0_9_)
                        d_292_v13_ = nw76_
                        d_295_v14_: D1
                        d_295_v14_ = D1_DC2(d_277_v0_)
                        rhs32_ = p1
                        rhs33_ = d_277_v0_
                        rhs34_ = (_dafny.CodePoint('x') if (d_295_v14_).cf2 else d_291_v12_)
                        rhs35_ = d_292_v13_
                        rhs36_ = d_277_v0_
                        lhs33_ = globalState
                        lhs34_ = globalState
                        d_290_v11_ = rhs32_
                        lhs33_.f1 = rhs33_
                        d_291_v12_ = rhs34_
                        d_292_v13_ = rhs35_
                        lhs34_.f3 = rhs36_
                        (self).f10 = (self).f8
                        d_296_v15_: C0
                        nw77_ = C0()
                        nw77_.ctor__(d_277_v0_)
                        d_296_v15_ = nw77_
                    (globalState).f3 = (self.f10) < (((self).f8) + (-750))
                    d_297_v16_: D2
                    d_297_v16_ = D2_DC4((self).f9)
                    d_298_v17_: _dafny.Map
                    def iife17_(_pat_let0_0):
                        def iife18_(d_299_dt__update__tmp_h0_):
                            def iife19_(_pat_let1_0):
                                def iife20_(d_300_dt__update_hcf4_h0_):
                                    return D2_DC4(d_300_dt__update_hcf4_h0_)
                                return iife20_(_pat_let1_0)
                            return iife19_((self).f9)
                        return iife18_(_pat_let0_0)
                    d_298_v17_ = _dafny.Map({(self).f8: iife17_(d_297_v16_)})
                    d_301_v18_: str
                    d_301_v18_ = _dafny.CodePoint('h')
                    d_302_v19_: _dafny.Seq
                    d_302_v19_ = _dafny.SeqWithoutIsStrInference([d_277_v0_, d_277_v0_, d_277_v0_])
                    d_303_v20_: _dafny.MultiSet
                    d_303_v20_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_277_v0_]), d_302_v19_, d_302_v19_, d_302_v19_])
                    d_304_v21_: _dafny.Array
                    nw78_ = _dafny.Array(None, 19)
                    nw78_[int(0)] = (self).f8
                    nw78_[int(1)] = (self).f9
                    nw78_[int(2)] = (self).f7
                    nw78_[int(3)] = (self).f7
                    nw78_[int(4)] = (self).f9
                    nw78_[int(5)] = self.f10
                    nw78_[int(6)] = (self).f7
                    nw78_[int(7)] = (self).f9
                    nw78_[int(8)] = self.f10
                    nw78_[int(9)] = (self).f7
                    nw78_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdqprjgij")))
                    nw78_[int(11)] = (self).f8
                    nw78_[int(12)] = len(d_298_v17_)
                    nw78_[int(13)] = (0) - (default__.fm0(d_301_v18_, globalState))
                    nw78_[int(14)] = -826
                    nw78_[int(15)] = (self).f9
                    nw78_[int(16)] = (self).f7
                    nw78_[int(17)] = (self).f7
                    nw78_[int(18)] = (d_303_v20_).cardinality
                    d_304_v21_ = nw78_
                    d_305_v22_: D0
                    d_305_v22_ = D0_DC0(d_304_v21_)
                    pat_let_tv4_ = d_304_v21_
                    d_306_v23_: _dafny.Array
                    nw79_ = _dafny.Array(None, 29)
                    def iife21_(_pat_let2_0):
                        def iife22_(d_307_dt__update__tmp_h1_):
                            def iife23_(_pat_let3_0):
                                def iife24_(d_308_dt__update_hcf0_h0_):
                                    return D0_DC0(d_308_dt__update_hcf0_h0_)
                                return iife24_(_pat_let3_0)
                            return iife23_(pat_let_tv4_)
                        return iife22_(_pat_let2_0)
                    nw79_[int(0)] = iife21_(d_305_v22_)
                    nw79_[int(1)] = d_305_v22_
                    nw79_[int(2)] = d_305_v22_
                    nw79_[int(3)] = d_305_v22_
                    nw79_[int(4)] = d_305_v22_
                    nw79_[int(5)] = D0_DC0((d_305_v22_).cf0)
                    nw79_[int(6)] = d_305_v22_
                    nw79_[int(7)] = d_305_v22_
                    nw79_[int(8)] = d_305_v22_
                    nw79_[int(9)] = d_305_v22_
                    nw79_[int(10)] = d_305_v22_
                    nw79_[int(11)] = d_305_v22_
                    nw79_[int(12)] = D0_DC0(d_304_v21_)
                    nw79_[int(13)] = d_305_v22_
                    nw79_[int(14)] = d_305_v22_
                    nw79_[int(15)] = d_305_v22_
                    nw79_[int(16)] = d_305_v22_
                    nw79_[int(17)] = d_305_v22_
                    nw79_[int(18)] = d_305_v22_
                    nw79_[int(19)] = D0_DC0(d_304_v21_)
                    nw79_[int(20)] = D0_DC0(d_304_v21_)
                    nw79_[int(21)] = d_305_v22_
                    nw79_[int(22)] = d_305_v22_
                    nw79_[int(23)] = d_305_v22_
                    nw79_[int(24)] = d_305_v22_
                    nw79_[int(25)] = d_305_v22_
                    nw79_[int(26)] = d_305_v22_
                    nw79_[int(27)] = d_305_v22_
                    nw79_[int(28)] = d_305_v22_
                    d_306_v23_ = nw79_
                    d_306_v23_ = d_306_v23_
                    d_309_v24_: D0
                    d_309_v24_ = D0_DC1(d_301_v18_)
                    d_310_v25_: _dafny.Set
                    d_310_v25_ = _dafny.Set({d_309_v24_, d_309_v24_})
                    pat_let_tv5_ = d_301_v18_
                    d_311_v26_: _dafny.Map
                    def iife25_(_pat_let4_0):
                        def iife26_(d_312_dt__update__tmp_h2_):
                            def iife27_(_pat_let5_0):
                                def iife28_(d_313_dt__update_hcf1_h0_):
                                    return D0_DC1(d_313_dt__update_hcf1_h0_)
                                return iife28_(_pat_let5_0)
                            return iife27_(pat_let_tv5_)
                        return iife26_(_pat_let4_0)
                    d_311_v26_ = _dafny.Map({d_277_v0_: (d_310_v25_).ispropersubset(_dafny.Set({d_309_v24_, iife25_(d_309_v24_)}))})
                    d_314_v27_: T0
                    nw80_ = C1()
                    nw80_.ctor__(d_277_v0_, self.f10, (self).f7, (0) - (self.f10))
                    d_314_v27_ = nw80_
                    d_315_v28_: D2
                    d_315_v28_ = D2_DC5((0) - ((self).f8), default__.fm11(d_277_v0_, self.f10, globalState), d_314_v27_)
                    d_316_v29_: _dafny.Map
                    def iife29_(_pat_let6_0):
                        def iife30_(d_317_dt__update__tmp_h3_):
                            def iife31_(_pat_let7_0):
                                def iife32_(d_318_dt__update_hcf5_h0_):
                                    return D2_DC5(d_318_dt__update_hcf5_h0_, (d_317_dt__update__tmp_h3_).cf6, (d_317_dt__update__tmp_h3_).cf7)
                                return iife32_(_pat_let7_0)
                            return iife31_(self.f10)
                        return iife30_(_pat_let6_0)
                    d_316_v29_ = _dafny.Map({iife29_(D2_DC5((self).f9, p0, d_314_v27_)): d_277_v0_})
                    if ((d_311_v26_)[(d_315_v28_) not in (d_316_v29_)] if ((d_315_v28_) not in (d_316_v29_)) in (d_311_v26_) else (d_303_v20_) != (d_303_v20_)):
                        d_319_v30_: _dafny.MultiSet
                        d_319_v30_ = _dafny.MultiSet([(self).f8])
                        d_320_v31_: _dafny.Set
                        d_320_v31_ = _dafny.Set({d_319_v30_})
                        d_321_v32_: D3
                        d_321_v32_ = D3_DC9(D3_DC7(d_320_v31_))
                        d_322_v33_: _dafny.Seq
                        d_322_v33_ = _dafny.SeqWithoutIsStrInference([d_321_v32_, d_321_v32_, d_321_v32_])
                        (globalState).f1 = (d_321_v32_) not in (d_322_v33_)
                        (globalState).f1 = default__.fm2((not(d_277_v0_) if d_277_v0_ else d_277_v0_), globalState)
                        d_323_v34_: _dafny.Array
                        def lambda18_(d_324_v27_):
                            def lambda19_(d_325_i3_):
                                return ((0) - ((d_324_v27_).f7)) <= ((d_324_v27_).f7)

                            return lambda19_

                        init10_ = lambda18_(d_314_v27_)
                        nw81_ = _dafny.Array(None, 3)
                        for i0_10_ in range(nw81_.length(0)):
                            nw81_[i0_10_] = init10_(i0_10_)
                        d_323_v34_ = nw81_
                        d_326_v35_: _dafny.Seq
                        d_326_v35_ = _dafny.SeqWithoutIsStrInference([d_314_v27_])
                        index59_ = default__.safeIndex(951, (d_323_v34_).length(0))
                        (d_323_v34_)[index59_] = (len(d_326_v35_)) != ((d_314_v27_).f8)
                        index60_ = default__.safeIndex(951, (d_323_v34_).length(0))
                        (d_323_v34_)[index60_] = d_277_v0_
                        index61_ = default__.safeIndex(680, (d_304_v21_).length(0))
                        (d_304_v21_)[index61_] = (d_314_v27_).f8
                        index62_ = default__.safeIndex(680, (d_304_v21_).length(0))
                        (d_304_v21_)[index62_] = default__.safeModuloInt((self).f9, (d_314_v27_).f8)
                        (self).f10 = (((self).f8 if (d_323_v34_)[default__.safeIndex(951, (d_323_v34_).length(0))] else default__.fm0(d_301_v18_, globalState))) * ((0) - ((d_314_v27_).f8))
                    elif True:
                        d_327_v36_: _dafny.Array
                        def lambda20_(d_328_v0_):
                            def lambda21_(d_329_i4_):
                                return d_328_v0_

                            return lambda21_

                        init11_ = lambda20_(d_277_v0_)
                        nw82_ = _dafny.Array(None, 22)
                        for i0_11_ in range(nw82_.length(0)):
                            nw82_[i0_11_] = init11_(i0_11_)
                        d_327_v36_ = nw82_
                        d_330_v37_: _dafny.Map
                        d_330_v37_ = _dafny.Map({d_277_v0_: d_327_v36_})
                        d_330_v37_ = d_330_v37_
                        (globalState).f1 = ((d_314_v27_).f7) == ((self).f7)
                        d_331_v38_: _dafny.Set
                        d_331_v38_ = _dafny.Set({(self).f9})
                        d_332_v39_: _dafny.Seq
                        d_332_v39_ = _dafny.SeqWithoutIsStrInference([-828, len(d_331_v38_), (d_314_v27_).f8, (d_314_v27_).f7, len(p0)])
                        d_332_v39_ = ((d_332_v39_) + (d_332_v39_)) + (d_332_v39_)
                        index63_ = default__.safeIndex(321, (d_304_v21_).length(0))
                        (d_304_v21_)[index63_] = (self).f8
                        index64_ = default__.safeIndex(321, (d_304_v21_).length(0))
                        rhs37_ = d_301_v18_
                        rhs38_ = (d_314_v27_).f8
                        lhs35_ = d_304_v21_
                        lhs36_ = default__.safeIndex(321, (d_304_v21_).length(0))
                        d_301_v18_ = rhs37_
                        lhs35_[lhs36_] = rhs38_
                        (globalState).f1 = d_277_v0_
                    pass
            pass
        d_333_v40_: _dafny.Array
        nw83_ = _dafny.Array(_dafny.Array(None, 0), 1)
        d_333_v40_ = nw83_
        nw84_ = _dafny.Array(_dafny.Array(None, 0), 3)
        d_333_v40_ = nw84_
        d_334_v41_: bool
        d_334_v41_ = True
        (globalState).f1 = (d_334_v41_) or (d_334_v41_)
        if (False if d_334_v41_ else d_334_v41_):
            (globalState).f1 = d_334_v41_
            if d_334_v41_:
                d_335_v42_: _dafny.Array
                def lambda22_(d_336_v41_):
                    def lambda23_(d_337_i5_):
                        return (d_336_v41_) and (d_336_v41_)

                    return lambda23_

                init12_ = lambda22_(d_334_v41_)
                nw85_ = _dafny.Array(None, 20)
                for i0_12_ in range(nw85_.length(0)):
                    nw85_[i0_12_] = init12_(i0_12_)
                d_335_v42_ = nw85_
                d_335_v42_ = d_335_v42_
                d_338_v43_: _dafny.Array
                nw86_ = _dafny.Array(D0.default()(), 2)
                d_338_v43_ = nw86_
                d_339_v44_: _dafny.MultiSet
                d_339_v44_ = _dafny.MultiSet([d_338_v43_])
                d_339_v44_ = ((_dafny.MultiSet([d_338_v43_, d_338_v43_])).intersection(d_339_v44_)).set(d_338_v43_, default__.abs(319))
                r0 = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cget"))) + (p0)
                index65_ = default__.safeIndex(710, (d_335_v42_).length(0))
                (d_335_v42_)[index65_] = d_334_v41_
                index66_ = default__.safeIndex(710, (d_335_v42_).length(0))
                rhs39_ = not(d_334_v41_)
                rhs40_ = self.f10
                rhs41_ = (self).f9
                rhs42_ = True
                rhs43_ = default__.fm0(_dafny.CodePoint('l'), globalState)
                lhs37_ = d_335_v42_
                lhs38_ = default__.safeIndex(710, (d_335_v42_).length(0))
                lhs39_ = self
                lhs40_ = self
                lhs41_ = globalState
                lhs42_ = self
                lhs37_[lhs38_] = rhs39_
                lhs39_.f10 = rhs40_
                lhs40_.f10 = rhs41_
                lhs41_.f1 = rhs42_
                lhs42_.f10 = rhs43_
                d_340_v45_: _dafny.Seq
                d_340_v45_ = _dafny.SeqWithoutIsStrInference([(d_335_v42_)[default__.safeIndex(710, (d_335_v42_).length(0))]])
                d_341_v46_: _dafny.Set
                d_341_v46_ = _dafny.Set({d_340_v45_, d_340_v45_, d_340_v45_})
                d_334_v41_ = not((len((d_341_v46_).intersection(_dafny.Set({d_340_v45_})))) != (default__.safeModuloInt(-139, -421)))
            elif True:
                (self).f10 = len(p0)
                rhs44_ = self.f10
                rhs45_ = d_334_v41_
                lhs43_ = self
                lhs44_ = globalState
                lhs43_.f10 = rhs44_
                lhs44_.f1 = rhs45_
                (globalState).f1 = d_334_v41_
                d_342_v47_: _dafny.Array
                nw87_ = _dafny.Array(False, 3)
                d_342_v47_ = nw87_
                index67_ = default__.safeIndex(729, (d_342_v47_).length(0))
                (d_342_v47_)[index67_] = d_334_v41_
                index68_ = default__.safeIndex(729, (d_342_v47_).length(0))
                (d_342_v47_)[index68_] = d_334_v41_
                d_343_v48_: C0
                nw88_ = C0()
                nw88_.ctor__(d_334_v41_)
                d_343_v48_ = nw88_
            d_344_v49_: _dafny.Seq
            d_344_v49_ = _dafny.SeqWithoutIsStrInference([d_334_v41_])
            d_345_v50_: _dafny.Map
            d_345_v50_ = _dafny.Map({default__.fm2(d_334_v41_, globalState): d_344_v49_})
            d_346_v51_: _dafny.Seq
            d_346_v51_ = _dafny.SeqWithoutIsStrInference([len(p0), (self).f7, self.f10, self.f10, (self).f8])
            d_347_v52_: T0
            nw89_ = C1()
            nw89_.ctor__(d_334_v41_, (self).f8, (self).f9, len(d_344_v49_))
            d_347_v52_ = nw89_
            d_348_v53_: _dafny.Map
            d_348_v53_ = _dafny.Map({d_334_v41_: d_347_v52_})
            d_349_v54_: _dafny.MultiSet
            d_349_v54_ = _dafny.MultiSet([(_dafny.MultiSet(d_344_v49_)).cardinality, len(d_344_v49_)])
            d_350_v55_: str
            d_350_v55_ = _dafny.CodePoint('y')
            d_351_v56_: _dafny.Array
            nw90_ = _dafny.Array(None, 29)
            nw90_[int(0)] = self.f10
            nw90_[int(1)] = (0) - (self.f10)
            nw90_[int(2)] = (len(d_344_v49_)) + ((self).f8)
            nw90_[int(3)] = (self).f7
            nw90_[int(4)] = len(d_345_v50_)
            nw90_[int(5)] = self.f10
            nw90_[int(6)] = (0) - ((self).f8)
            nw90_[int(7)] = (self).f7
            nw90_[int(8)] = (self).f7
            nw90_[int(9)] = (len(d_346_v51_) if d_334_v41_ else (self).f8)
            nw90_[int(10)] = (self).fm7(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_352_i6_ in range(default__.abs(389))]), 454, globalState)
            nw90_[int(11)] = default__.safeModuloInt((0) - ((self).f7), (self).f9)
            nw90_[int(12)] = self.f10
            nw90_[int(13)] = ((0) - (len(d_348_v53_))) - ((self).f7)
            nw90_[int(14)] = (d_347_v52_).f8
            nw90_[int(15)] = (self).f9
            nw90_[int(16)] = (d_347_v52_).f8
            nw90_[int(17)] = (self).f7
            nw90_[int(18)] = (self).f8
            nw90_[int(19)] = self.f10
            nw90_[int(20)] = (0) - ((self).f7)
            nw90_[int(21)] = (self).f9
            nw90_[int(22)] = (self).f7
            nw90_[int(23)] = (0) - (((self).f8) - (-46))
            nw90_[int(24)] = (d_347_v52_).f8
            nw90_[int(25)] = (((d_349_v54_)[(self).f9] if ((self).f9) in (d_349_v54_) else (self).f8)) + ((self).f9)
            nw90_[int(26)] = self.f10
            nw90_[int(27)] = default__.fm0(d_350_v55_, globalState)
            nw90_[int(28)] = self.f10
            d_351_v56_ = nw90_
            index69_ = default__.safeIndex(627, (d_351_v56_).length(0))
            (d_351_v56_)[index69_] = len(d_346_v51_)
            d_353_v57_: _dafny.Map
            d_353_v57_ = _dafny.Map({d_350_v55_: d_334_v41_})
            d_354_v58_: _dafny.Seq
            d_354_v58_ = _dafny.SeqWithoutIsStrInference([d_347_v52_, d_347_v52_])
            d_355_v59_: _dafny.Map
            d_355_v59_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_350_v55_ for d_356_i7_ in range(default__.abs(516))])): d_354_v58_})
            index70_ = default__.safeIndex(627, (d_351_v56_).length(0))
            (d_351_v56_)[index70_] = len((d_353_v57_) | (default__.fm13(len(((d_355_v59_)[self.f10] if (self.f10) in (d_355_v59_) else d_354_v58_)), globalState)))
            d_357_v60_: _dafny.Map
            d_357_v60_ = _dafny.Map({_dafny.Set({d_334_v41_}): d_344_v49_})
            d_358_v61_: C1
            nw91_ = C1()
            nw91_.ctor__(d_334_v41_, (0) - (len(p0)), self.f10, len(d_357_v60_))
            d_358_v61_ = nw91_
            d_359_v62_: D1
            d_359_v62_ = D1_DC2(d_334_v41_)
            pat_let_tv6_ = d_334_v41_
            d_360_v63_: _dafny.Set
            def iife33_(_pat_let8_0):
                def iife34_(d_361_dt__update__tmp_h4_):
                    def iife35_(_pat_let9_0):
                        def iife36_(d_362_dt__update_hcf2_h0_):
                            return D1_DC2(d_362_dt__update_hcf2_h0_)
                        return iife36_(_pat_let9_0)
                    return iife35_(pat_let_tv6_)
                return iife34_(_pat_let8_0)
            d_360_v63_ = _dafny.Set({iife33_(d_359_v62_)})
            (globalState).f3 = (d_360_v63_).ispropersubset(default__.fm14(d_358_v61_.f13, not((d_358_v61_).f12), globalState))
        elif True:
            if (p0) <= (p0):
                d_363_v64_: str
                d_363_v64_ = _dafny.CodePoint('s')
                d_364_v65_: _dafny.Map
                d_364_v65_ = _dafny.Map({d_334_v41_: d_363_v64_})
                d_365_v66_: _dafny.Map
                d_365_v66_ = _dafny.Map({_dafny.CodePoint('r'): ((d_364_v65_)[d_334_v41_] if (d_334_v41_) in (d_364_v65_) else d_363_v64_)})
                d_366_v67_: _dafny.Map
                d_366_v67_ = _dafny.Map({d_365_v66_: (self).f9})
                d_367_v68_: _dafny.Seq
                d_367_v68_ = _dafny.SeqWithoutIsStrInference([d_334_v41_])
                d_368_v69_: C1
                nw92_ = C1()
                nw92_.ctor__(not(d_334_v41_), ((d_366_v67_)[d_365_v66_] if (d_365_v66_) in (d_366_v67_) else len(default__.fm15(globalState))), (self).f9, len(d_367_v68_))
                d_368_v69_ = nw92_
                d_369_v70_: _dafny.Seq
                d_370_v71_: D3
                out3_: _dafny.Seq
                out4_: D3
                out3_, out4_ = (d_368_v69_).m6(d_368_v69_.f13, globalState)
                d_369_v70_ = out3_
                d_370_v71_ = out4_
                d_371_v73_: C0
                nw93_ = C0()
                nw93_.ctor__(not(d_334_v41_))
                d_371_v73_ = nw93_
                d_372_v74_: _dafny.MultiSet
                d_372_v74_ = _dafny.MultiSet([d_371_v73_, d_371_v73_])
                d_373_v75_: _dafny.Map
                d_373_v75_ = _dafny.Map({(self).fm7((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glko"))).set(default__.safeIndex(283, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "glko")))), _dafny.CodePoint('o')), -128, globalState): d_372_v74_})
                d_374_v76_: _dafny.Set
                d_374_v76_ = _dafny.Set({(d_368_v69_).f12})
                d_375_v77_: _dafny.Array
                nw94_ = _dafny.Array(None, 19)
                nw94_[int(0)] = (d_368_v69_).f12
                nw94_[int(1)] = d_334_v41_
                nw94_[int(2)] = (d_368_v69_).f12
                nw94_[int(3)] = d_334_v41_
                nw94_[int(4)] = (d_368_v69_).f12
                nw94_[int(5)] = (d_371_v73_).f11
                nw94_[int(6)] = d_334_v41_
                nw94_[int(7)] = d_334_v41_
                nw94_[int(8)] = d_334_v41_
                nw94_[int(9)] = (d_368_v69_).f12
                nw94_[int(10)] = not((d_368_v69_).f12)
                nw94_[int(11)] = (d_368_v69_).f12
                nw94_[int(12)] = d_334_v41_
                nw94_[int(13)] = not((d_371_v73_).f11)
                nw94_[int(14)] = (d_368_v69_).f12
                nw94_[int(15)] = (d_371_v73_).f11
                nw94_[int(16)] = False
                nw94_[int(17)] = (d_368_v69_).f12
                nw94_[int(18)] = d_334_v41_
                d_375_v77_ = nw94_
                d_376_v78_: _dafny.Seq
                d_376_v78_ = _dafny.SeqWithoutIsStrInference([d_375_v77_])
                d_377_v79_: _dafny.Map
                d_377_v79_ = _dafny.Map({(self).f9: (d_371_v73_).f11})
                d_378_v80_: _dafny.Array
                nw95_ = _dafny.Array(None, 19)
                nw95_[int(0)] = (len(p0)) * (len(_dafny.SeqWithoutIsStrInference([d_363_v64_ for d_379_i8_ in range(default__.abs(547))])))
                def iife37_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in (_dafny.SeqWithoutIsStrInference([(self).f7])).Elements:
                        d_380_v72_: int = compr_17_
                        if (d_380_v72_) in (_dafny.SeqWithoutIsStrInference([(self).f7])):
                            coll17_[default__.safeDivisionInt(d_380_v72_, 950)] = self.f10
                    return _dafny.Map(coll17_)
                nw95_[int(1)] = (len((iife37_()
                ).set(d_368_v69_.f13, d_368_v69_.f13))) * ((self).f7)
                nw95_[int(2)] = (self).f9
                nw95_[int(3)] = len(_dafny.Map({((d_373_v75_)[(self).f8] if ((self).f8) in (d_373_v75_) else d_372_v74_): (self).f9}))
                nw95_[int(4)] = ((self).f7 if d_334_v41_ else self.f10)
                nw95_[int(5)] = (self).f9
                nw95_[int(6)] = (76) - ((self).f9)
                nw95_[int(7)] = ((self).f7) - (self.f10)
                nw95_[int(8)] = len((p0) + (p0))
                nw95_[int(9)] = (0) - (d_368_v69_.f13)
                nw95_[int(10)] = -325
                nw95_[int(11)] = (self).f9
                nw95_[int(12)] = ((self).f8) * (len(d_374_v76_))
                nw95_[int(13)] = default__.safeDivisionInt((0) - ((0) - (self.f10)), d_368_v69_.f13)
                nw95_[int(14)] = d_368_v69_.f13
                nw95_[int(15)] = (self).f7
                nw95_[int(16)] = (self).f8
                nw95_[int(17)] = (len(_dafny.SeqWithoutIsStrInference([self.f10, (self).f8]))) * (len(d_376_v78_))
                nw95_[int(18)] = len((d_377_v79_) | (d_377_v79_))
                d_378_v80_ = nw95_
                index71_ = default__.safeIndex(530, (d_378_v80_).length(0))
                (d_378_v80_)[index71_] = 793
                index72_ = default__.safeIndex(530, (d_378_v80_).length(0))
                (d_378_v80_)[index72_] = len(d_374_v76_)
                (globalState).f1 = d_334_v41_
                (self).f10 = (0) - (((d_378_v80_)[default__.safeIndex(530, (d_378_v80_).length(0))]) - (-91))
            elif True:
                (globalState).f3 = (not (d_334_v41_) or (d_334_v41_)) or (d_334_v41_)
                d_381_v81_: str
                d_381_v81_ = _dafny.CodePoint('y')
                r0 = (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_382_i9_ in range(default__.abs(495))])) + (p0)).set(default__.safeIndex((self).f9, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_383_i9_ in range(default__.abs(495))])) + (p0))), d_381_v81_)) + (p0)
                (globalState).f1 = False
                d_384_v82_: _dafny.Array
                nw96_ = _dafny.Array(None, 8)
                nw96_[int(0)] = (self).f8
                nw96_[int(1)] = (self).f9
                nw96_[int(2)] = len(_dafny.SeqWithoutIsStrInference([(self).f7 for d_385_i10_ in range(default__.abs(468))]))
                nw96_[int(3)] = (self).f8
                nw96_[int(4)] = len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f7])).cardinality]))
                nw96_[int(5)] = 404
                nw96_[int(6)] = self.f10
                nw96_[int(7)] = (self).f8
                d_384_v82_ = nw96_
                d_386_v83_: _dafny.Seq
                d_386_v83_ = _dafny.SeqWithoutIsStrInference([d_384_v82_, d_384_v82_])
                d_387_v84_: _dafny.Array
                nw97_ = _dafny.Array(None, 8)
                nw97_[int(0)] = d_334_v41_
                nw97_[int(1)] = d_334_v41_
                nw97_[int(2)] = False
                nw97_[int(3)] = d_334_v41_
                nw97_[int(4)] = d_334_v41_
                nw97_[int(5)] = d_334_v41_
                nw97_[int(6)] = d_334_v41_
                nw97_[int(7)] = d_334_v41_
                d_387_v84_ = nw97_
                d_388_v85_: _dafny.Map
                d_388_v85_ = _dafny.Map({len(d_386_v83_): d_387_v84_})
                d_389_v86_: _dafny.Seq
                d_389_v86_ = _dafny.SeqWithoutIsStrInference([len(d_388_v85_), (self).f9, (self).f8])
                d_390_v87_: _dafny.Seq
                d_390_v87_ = _dafny.SeqWithoutIsStrInference([(d_389_v86_)[default__.safeIndex(len(p0), len(d_389_v86_))], ((self).f7) - ((self).f7)])
                d_391_v88_: _dafny.MultiSet
                d_391_v88_ = _dafny.MultiSet([default__.fm16(globalState), default__.fm16(globalState), d_389_v86_, d_390_v87_])
                d_392_v89_: _dafny.Map
                d_392_v89_ = _dafny.Map({d_391_v88_: (self).f7})
                rhs46_ = not(d_334_v41_)
                rhs47_ = (_dafny.SeqWithoutIsStrInference([self.f10, self.f10, ((d_392_v89_)[d_391_v88_] if (d_391_v88_) in (d_392_v89_) else self.f10)])).set(default__.safeIndex((498) + ((self).f8), len(_dafny.SeqWithoutIsStrInference([self.f10, self.f10, ((d_392_v89_)[d_391_v88_] if (d_391_v88_) in (d_392_v89_) else self.f10)]))), (self).f9)
                d_334_v41_ = rhs46_
                d_390_v87_ = rhs47_
                r0 = p0
            if d_334_v41_:
                d_393_v90_: D2
                d_393_v90_ = D2_DC4(self.f10)
                d_394_v91_: _dafny.Set
                d_394_v91_ = _dafny.Set({d_334_v41_})
                d_395_v92_: _dafny.Array
                nw98_ = _dafny.Array(None, 4)
                nw98_[int(0)] = (self).f8
                nw98_[int(1)] = default__.safeModuloInt((default__.fm17(d_334_v41_, self.f10, globalState)).cardinality, (self).f8)
                nw98_[int(2)] = (d_393_v90_).cf4
                nw98_[int(3)] = (0) - (default__.safeModuloInt(self.f10, len(d_394_v91_)))
                d_395_v92_ = nw98_
                rhs48_ = d_395_v92_
                rhs49_ = self.f10
                lhs45_ = self
                d_395_v92_ = rhs48_
                lhs45_.f10 = rhs49_
                (globalState).f1 = ((_dafny.Set({d_334_v41_})).intersection(d_394_v91_)).issubset(d_394_v91_)
                (globalState).f3 = d_334_v41_
                d_393_v90_ = d_393_v90_
                (globalState).f1 = d_334_v41_
            elif True:
                d_396_v93_: _dafny.Seq
                d_396_v93_ = _dafny.SeqWithoutIsStrInference([len(p0)])
                d_397_v94_: _dafny.Set
                d_397_v94_ = _dafny.Set({p0})
                d_398_v95_: _dafny.Seq
                d_398_v95_ = _dafny.SeqWithoutIsStrInference([d_334_v41_])
                d_399_v96_: _dafny.Array
                nw99_ = _dafny.Array(False, 2)
                d_399_v96_ = nw99_
                d_400_v97_: _dafny.Map
                d_400_v97_ = _dafny.Map({d_334_v41_: d_399_v96_})
                d_401_v98_: _dafny.Seq
                d_401_v98_ = _dafny.SeqWithoutIsStrInference([d_399_v96_])
                d_402_v99_: _dafny.MultiSet
                d_402_v99_ = _dafny.MultiSet([d_400_v97_, (d_400_v97_).set(d_334_v41_, (d_401_v98_)[default__.safeIndex(-514, len(d_401_v98_))]), d_400_v97_, _dafny.Map({False: d_399_v96_}), d_400_v97_])
                d_403_v100_: _dafny.Array
                nw100_ = _dafny.Array(None, 16)
                nw100_[int(0)] = (default__.fm16(globalState)) < (d_396_v93_)
                nw100_[int(1)] = d_334_v41_
                nw100_[int(2)] = d_334_v41_
                nw100_[int(3)] = d_334_v41_
                nw100_[int(4)] = (_dafny.Set({p0, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_404_i11_ in range(default__.abs(93))])})).issubset(d_397_v94_)
                nw100_[int(5)] = ((self).f7) != ((self).f8)
                nw100_[int(6)] = (d_334_v41_ if d_334_v41_ else True)
                nw100_[int(7)] = (d_334_v41_) and (d_334_v41_)
                nw100_[int(8)] = d_334_v41_
                nw100_[int(9)] = (d_334_v41_ if default__.fm2(d_334_v41_, globalState) else d_334_v41_)
                nw100_[int(10)] = d_334_v41_
                nw100_[int(11)] = (d_398_v95_) < (d_398_v95_)
                nw100_[int(12)] = d_334_v41_
                nw100_[int(13)] = (((d_402_v99_)[d_400_v97_] if (d_400_v97_) in (d_402_v99_) else (self).f8)) > ((self).f9)
                nw100_[int(14)] = d_334_v41_
                nw100_[int(15)] = d_334_v41_
                d_403_v100_ = nw100_
                index73_ = default__.safeIndex(867, (d_399_v96_).length(0))
                (d_399_v96_)[index73_] = not(d_334_v41_)
                d_405_v101_: _dafny.Map
                d_405_v101_ = _dafny.Map({(p0) + (p0): (self).f7})
                d_406_v102_: str
                d_406_v102_ = _dafny.CodePoint('l')
                d_407_v103_: _dafny.Array
                nw101_ = _dafny.Array(None, 10)
                nw101_[int(0)] = d_406_v102_
                nw101_[int(1)] = d_406_v102_
                nw101_[int(2)] = d_406_v102_
                nw101_[int(3)] = d_406_v102_
                nw101_[int(4)] = d_406_v102_
                nw101_[int(5)] = d_406_v102_
                nw101_[int(6)] = (d_406_v102_ if d_334_v41_ else d_406_v102_)
                nw101_[int(7)] = d_406_v102_
                nw101_[int(8)] = default__.fm3(d_334_v41_, globalState)
                nw101_[int(9)] = d_406_v102_
                d_407_v103_ = nw101_
                index74_ = default__.safeIndex(867, (d_399_v96_).length(0))
                rhs50_ = d_334_v41_
                rhs51_ = d_405_v101_
                rhs52_ = len(_dafny.Map({(self).f7: (self).f8}))
                rhs53_ = d_407_v103_
                rhs54_ = d_399_v96_
                lhs46_ = d_399_v96_
                lhs47_ = default__.safeIndex(867, (d_399_v96_).length(0))
                lhs48_ = self
                lhs46_[lhs47_] = rhs50_
                d_405_v101_ = rhs51_
                lhs48_.f10 = rhs52_
                d_407_v103_ = rhs53_
                d_399_v96_ = rhs54_
                d_408_v104_: _dafny.Set
                d_408_v104_ = _dafny.Set({False, False})
                d_409_v105_: C1
                nw102_ = C1()
                nw102_.ctor__(True, len(d_408_v104_), ((self).f7) - (self.f10), ((self).f9 if not(d_334_v41_) else (0) - ((self).f8)))
                d_409_v105_ = nw102_
                d_410_v106_: _dafny.Array
                nw103_ = _dafny.Array(None, 15)
                nw103_[int(0)] = d_408_v104_
                nw103_[int(1)] = (d_408_v104_) | (d_408_v104_)
                nw103_[int(2)] = d_408_v104_
                nw103_[int(3)] = d_408_v104_
                nw103_[int(4)] = d_408_v104_
                nw103_[int(5)] = (d_408_v104_ if (d_399_v96_)[default__.safeIndex(867, (d_399_v96_).length(0))] else d_408_v104_)
                nw103_[int(6)] = (d_408_v104_) - (d_408_v104_)
                nw103_[int(7)] = d_408_v104_
                nw103_[int(8)] = (d_408_v104_ if (d_399_v96_)[default__.safeIndex(867, (d_399_v96_).length(0))] else d_408_v104_)
                nw103_[int(9)] = d_408_v104_
                nw103_[int(10)] = _dafny.Set({(d_409_v105_).f12})
                nw103_[int(11)] = _dafny.Set({(d_409_v105_).f12, d_334_v41_, (d_409_v105_).f12, False, d_334_v41_})
                nw103_[int(12)] = d_408_v104_
                nw103_[int(13)] = (d_408_v104_) - (d_408_v104_)
                nw103_[int(14)] = d_408_v104_
                d_410_v106_ = nw103_
                index75_ = default__.safeIndex(904, (d_410_v106_).length(0))
                (d_410_v106_)[index75_] = d_408_v104_
                index76_ = default__.safeIndex(904, (d_410_v106_).length(0))
                rhs55_ = d_408_v104_
                rhs56_ = d_403_v100_
                rhs57_ = (p0)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_409_v105_.f13 for d_411_i12_ in range(default__.abs(387))])), len(p0))]
                lhs49_ = d_410_v106_
                lhs50_ = default__.safeIndex(904, (d_410_v106_).length(0))
                lhs49_[lhs50_] = rhs55_
                d_403_v100_ = rhs56_
                d_406_v102_ = rhs57_
                r0 = p0
                (globalState).f1 = False
            (self).f10 = self.f10
            d_412_v107_: _dafny.Array
            nw104_ = _dafny.Array(False, 6)
            d_412_v107_ = nw104_
            d_413_v108_: str
            d_413_v108_ = _dafny.CodePoint('i')
            d_414_v109_: D4
            d_414_v109_ = D4_DC11((self).f7, d_413_v108_)
            index77_ = default__.safeIndex(674, (d_412_v107_).length(0))
            (d_412_v107_)[index77_] = ((d_414_v109_).cf13) != (284)
            d_415_v110_: _dafny.MultiSet
            d_415_v110_ = _dafny.MultiSet([(self).fm9(d_334_v41_, self.f10, globalState), d_334_v41_])
            d_416_v111_: _dafny.Seq
            d_416_v111_ = _dafny.SeqWithoutIsStrInference([d_334_v41_, d_334_v41_])
            d_417_v112_: _dafny.Seq
            d_417_v112_ = _dafny.SeqWithoutIsStrInference([d_415_v110_, _dafny.MultiSet([d_334_v41_])])
            d_418_v113_: _dafny.Array
            nw105_ = _dafny.Array(None, 12)
            nw105_[int(0)] = d_415_v110_
            nw105_[int(1)] = (d_415_v110_).intersection(d_415_v110_)
            nw105_[int(2)] = d_415_v110_
            nw105_[int(3)] = d_415_v110_
            nw105_[int(4)] = (_dafny.MultiSet(d_416_v111_)).set(d_334_v41_, default__.abs((self).f9))
            nw105_[int(5)] = (d_415_v110_ if d_334_v41_ else d_415_v110_)
            nw105_[int(6)] = d_415_v110_
            nw105_[int(7)] = d_415_v110_
            nw105_[int(8)] = d_415_v110_
            nw105_[int(9)] = d_415_v110_
            nw105_[int(10)] = _dafny.MultiSet([d_334_v41_, d_334_v41_, d_334_v41_])
            nw105_[int(11)] = (d_417_v112_)[default__.safeIndex((self).f7, len(d_417_v112_))]
            d_418_v113_ = nw105_
            index78_ = default__.safeIndex(918, (d_418_v113_).length(0))
            (d_418_v113_)[index78_] = _dafny.MultiSet([d_334_v41_])
            d_419_v114_: _dafny.Set
            d_419_v114_ = _dafny.Set({d_413_v108_, d_413_v108_})
            d_420_v115_: _dafny.Map
            d_420_v115_ = _dafny.Map({d_419_v114_: self.f10})
            index79_ = default__.safeIndex(674, (d_412_v107_).length(0))
            index80_ = default__.safeIndex(918, (d_418_v113_).length(0))
            rhs58_ = (d_334_v41_ if True else not((d_420_v115_) == (d_420_v115_)))
            rhs59_ = d_415_v110_
            rhs60_ = default__.safeDivisionInt(((self).f8) * ((0) - (self.f10)), (self).f7)
            lhs51_ = d_412_v107_
            lhs52_ = default__.safeIndex(674, (d_412_v107_).length(0))
            lhs53_ = d_418_v113_
            lhs54_ = default__.safeIndex(918, (d_418_v113_).length(0))
            lhs55_ = self
            lhs51_[lhs52_] = rhs58_
            lhs53_[lhs54_] = rhs59_
            lhs55_.f10 = rhs60_
            d_421_v116_: _dafny.Set
            d_421_v116_ = _dafny.Set({(self).f9, 658, 163})
            d_422_v117_: _dafny.Map
            d_422_v117_ = _dafny.Map({d_421_v116_: True})
            d_423_v118_: _dafny.Map
            d_423_v118_ = _dafny.Map({(d_412_v107_)[default__.safeIndex(674, (d_412_v107_).length(0))]: d_422_v117_})
            d_424_v119_: _dafny.Map
            d_424_v119_ = _dafny.Map({((d_423_v118_)[d_334_v41_] if (d_334_v41_) in (d_423_v118_) else d_422_v117_): d_334_v41_})
            index81_ = default__.safeIndex(674, (d_412_v107_).length(0))
            rhs61_ = (d_412_v107_)[default__.safeIndex(674, (d_412_v107_).length(0))]
            rhs62_ = (d_412_v107_ if d_334_v41_ else d_412_v107_)
            rhs63_ = False
            rhs64_ = ((d_424_v119_)[d_422_v117_] if (d_422_v117_) in (d_424_v119_) else (d_412_v107_)[default__.safeIndex(674, (d_412_v107_).length(0))])
            lhs56_ = d_412_v107_
            lhs57_ = default__.safeIndex(674, (d_412_v107_).length(0))
            d_334_v41_ = rhs61_
            d_412_v107_ = rhs62_
            d_334_v41_ = rhs63_
            lhs56_[lhs57_] = rhs64_
        d_425_v120_: _dafny.Seq
        d_425_v120_ = _dafny.SeqWithoutIsStrInference([(self).f8])
        d_426_v121_: str
        d_426_v121_ = _dafny.CodePoint('q')
        r0 = ((p0).set(default__.safeIndex(len(d_425_v120_), len(p0)), d_426_v121_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhgenyj")))
        r0 = p0
        return r0

    @property
    def f9(self):
        return self._f9

class C3(T0):
    def  __init__(self):
        self._f7: int = int(0)
        self._f8: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f7(self):
        return self._f7
    @property
    def f8(self):
        return self._f8
    def ctor__(self, f7, f8):
        (self)._f7 = f7
        (self)._f8 = f8

    def fm6(self, p0, p1, globalState):
        return (_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False])), (self).f7, (0) - ((self).f8)]), _dafny.MultiSet([(self).f8, (self).f8]), _dafny.MultiSet([(self).f7, (self).f7, (self).f8, (self).f7])})) - (_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qt"))]))]), _dafny.MultiSet([670])}))

    def fm7(self, p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('r'), _dafny.CodePoint('e')])) | ((_dafny.MultiSet([_dafny.CodePoint('g'), _dafny.CodePoint('b')])) | (_dafny.MultiSet([_dafny.CodePoint('q'), _dafny.CodePoint('w')])))).cardinality

    def m3(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: int = int(0)
        r3: bool = False
        r2 = default__.safeDivisionInt(p3, (self).f8)
        d_427_v0_: _dafny.Array
        nw106_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 15)
        d_427_v0_ = nw106_
        d_427_v0_ = d_427_v0_
        d_428_v1_: _dafny.Seq
        d_428_v1_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(self).f8, len(p0)])), ((self).f8) * (p3)])
        rhs65_ = d_428_v1_
        rhs66_ = p1
        d_428_v1_ = rhs65_
        r3 = rhs66_
        d_429_v2_: _dafny.Map
        d_429_v2_ = _dafny.Map({len(_dafny.Map({p1: p1})): (0) - (len(p0))})
        d_429_v2_ = (d_429_v2_).set((0) - ((self).f8), (self).f7)
        d_430_v3_: _dafny.Seq
        d_430_v3_ = _dafny.SeqWithoutIsStrInference([not(p1), True])
        d_431_v4_: _dafny.MultiSet
        d_431_v4_ = _dafny.MultiSet([D1_DC2(p1), p2])
        d_430_v3_ = default__.fm8(d_428_v1_, -163, p1, d_431_v4_, globalState)
        d_432_v5_: str
        d_432_v5_ = _dafny.CodePoint('c')
        d_433_v6_: D1
        d_433_v6_ = D1_DC3(d_432_v5_)
        def iife38_(_pat_let10_0):
            def iife39_(d_434_dt__update__tmp_h0_):
                def iife40_(_pat_let11_0):
                    def iife41_(d_435_dt__update_hcf3_h0_):
                        return D1_DC3(d_435_dt__update_hcf3_h0_)
                    return iife41_(_pat_let11_0)
                return iife40_(_dafny.CodePoint('f'))
            return iife39_(_pat_let10_0)
        source4_ = (iife38_(d_433_v6_) if default__.fm2(p1, globalState) else d_433_v6_)
        if source4_.is_DC3:
            d_436___mcc_h0_ = source4_.cf3
            d_437_cf3_ = d_436___mcc_h0_
            d_438_v7_: _dafny.MultiSet
            d_438_v7_ = _dafny.MultiSet([p1])
            d_439_v8_: _dafny.Seq
            d_439_v8_ = _dafny.SeqWithoutIsStrInference([d_438_v7_, d_438_v7_])
            d_439_v8_ = d_439_v8_
            if p1:
                d_440_v9_: C2
                nw107_ = C2()
                nw107_.ctor__(p3, p3, default__.fm0(default__.fm3(False, globalState), globalState), default__.safeDivisionInt((0) - (p3), (self).f7))
                d_440_v9_ = nw107_
                (d_440_v9_).f10 = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")))) - (len((d_430_v3_).set(default__.safeIndex((self).f8, len(d_430_v3_)), p1)))) * ((761 if p1 else (d_440_v9_).f9))
                d_441_v10_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 10)
                d_441_v10_ = nw108_
                index82_ = default__.safeIndex(922, (d_441_v10_).length(0))
                (d_441_v10_)[index82_] = d_440_v9_.f10
                index83_ = default__.safeIndex(922, (d_441_v10_).length(0))
                (d_441_v10_)[index83_] = default__.fm0(d_437_cf3_, globalState)
                d_442_v11_: D2
                d_442_v11_ = D2_DC4((self).f8)
                d_443_v12_: D2
                d_443_v12_ = D2_DC6(d_442_v11_)
                d_444_v13_: D2
                d_444_v13_ = D2_DC6(d_443_v12_)
                d_445_v14_: D2
                d_445_v14_ = D2_DC6(d_442_v11_)
                d_446_v15_: D2
                d_446_v15_ = D2_DC6(d_445_v14_)
                d_447_v16_: _dafny.Map
                d_447_v16_ = _dafny.Map({d_446_v15_: (self).f7})
                d_448_v17_: _dafny.Map
                d_448_v17_ = _dafny.Map({p1: _dafny.Map({D2_DC6(d_445_v14_): -176})})
                d_449_v18_: _dafny.Map
                d_449_v18_ = _dafny.Map({d_437_cf3_: p1})
                d_450_v19_: _dafny.Map
                d_450_v19_ = _dafny.Map({(d_447_v16_) | (((d_448_v17_)[p1] if (p1) in (d_448_v17_) else d_447_v16_)): ((d_449_v18_)[d_437_cf3_] if (d_437_cf3_) in (d_449_v18_) else p1)})
                d_450_v19_ = _dafny.Map({d_447_v16_: p1})
                d_451_v20_: C0
                nw109_ = C0()
                nw109_.ctor__(p1)
                d_451_v20_ = nw109_
            elif True:
                d_452_v21_: _dafny.Map
                d_452_v21_ = _dafny.Map({(False) == (False): p1})
                d_452_v21_ = (d_452_v21_).set(p1, p1)
                d_453_v22_: T0
                nw110_ = C1()
                nw110_.ctor__(p1, default__.fm0(d_437_cf3_, globalState), (self).f7, (self).f7)
                d_453_v22_ = nw110_
                d_454_v23_: D2
                d_454_v23_ = D2_DC5((self).f7, p0, d_453_v22_)
                d_455_v24_: D2
                d_455_v24_ = D2_DC6(d_454_v23_)
                d_456_v25_: _dafny.Map
                d_456_v25_ = _dafny.Map({(self).f8: _dafny.Map({p3: d_455_v24_})})
                d_456_v25_ = d_456_v25_
                d_457_v26_: _dafny.Array
                nw111_ = _dafny.Array(int(0), 14)
                d_457_v26_ = nw111_
                d_457_v26_ = (D0_DC0(d_457_v26_)).cf0
                d_458_v27_: D0
                d_458_v27_ = D0_DC0(d_457_v26_)
                d_459_v28_: D0
                d_459_v28_ = D0_DC0((d_458_v27_).cf0)
                d_458_v27_ = d_458_v27_
                r2 = default__.safeDivisionInt(486, (d_453_v22_).f7)
            index84_ = default__.safeIndex(759, (d_427_v0_).length(0))
            (d_427_v0_)[index84_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hql"))) + (p0)
            d_460_v29_: _dafny.Array
            nw112_ = _dafny.Array(int(0), 20)
            d_460_v29_ = nw112_
            d_461_v30_: _dafny.Set
            d_461_v30_ = _dafny.Set({d_460_v29_, d_460_v29_})
            index85_ = default__.safeIndex(654, (d_460_v29_).length(0))
            (d_460_v29_)[index85_] = default__.safeDivisionInt(len(d_461_v30_), p3)
            d_462_v32_: _dafny.Set
            d_462_v32_ = _dafny.Set({d_437_cf3_})
            d_463_v33_: _dafny.Set
            d_463_v33_ = _dafny.Set({len(d_462_v32_)})
            index86_ = default__.safeIndex(759, (d_427_v0_).length(0))
            index87_ = default__.safeIndex(654, (d_460_v29_).length(0))
            def iife42_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in (d_463_v33_).Elements:
                    d_464_v31_: int = compr_18_
                    if (d_464_v31_) in (d_463_v33_):
                        coll18_[(d_464_v31_) - ((self).f7)] = (D2_DC4(p3)).cf4
                return _dafny.Map(coll18_)
            rhs67_ = default__.fm2(p1, globalState)
            rhs68_ = (p0) + (default__.fm11(p1, len(p0), globalState))
            rhs69_ = p3
            rhs70_ = len(((d_429_v2_) | (d_429_v2_)) | (iife42_()
            ))
            lhs58_ = globalState
            lhs59_ = d_427_v0_
            lhs60_ = default__.safeIndex(759, (d_427_v0_).length(0))
            lhs61_ = d_460_v29_
            lhs62_ = default__.safeIndex(654, (d_460_v29_).length(0))
            lhs58_.f1 = rhs67_
            lhs59_[lhs60_] = rhs68_
            lhs61_[lhs62_] = rhs69_
            r0 = rhs70_
            default__.m0(globalState)
        elif True:
            d_465___mcc_h1_ = source4_.cf2
            d_466_cf2_ = d_465___mcc_h1_
            d_467_v34_: _dafny.Array
            nw113_ = _dafny.Array(False, 28)
            d_467_v34_ = nw113_
            index88_ = default__.safeIndex(558, (d_467_v34_).length(0))
            (d_467_v34_)[index88_] = p1
            index89_ = default__.safeIndex(558, (d_467_v34_).length(0))
            (d_467_v34_)[index89_] = d_466_cf2_
            d_468_v35_: _dafny.Array
            nw114_ = _dafny.Array(int(0), 14)
            d_468_v35_ = nw114_
            index90_ = default__.safeIndex(175, (d_468_v35_).length(0))
            (d_468_v35_)[index90_] = default__.safeModuloInt(default__.fm0(_dafny.CodePoint('d'), globalState), p3)
            index91_ = default__.safeIndex(175, (d_468_v35_).length(0))
            (d_468_v35_)[index91_] = (p3) - ((0) - ((self).f7))
            (globalState).f3 = d_466_cf2_
            d_466_cf2_ = (d_428_v1_) <= ((d_428_v1_) + (d_428_v1_))
        r0 = default__.safeDivisionInt(len(d_430_v3_), (self).f7)
        d_469_v36_: D4
        d_469_v36_ = D4_DC11((self).f7, d_432_v5_)
        d_470_v37_: D4
        d_470_v37_ = D4_DC13(d_469_v36_)
        pat_let_tv7_ = d_432_v5_
        pat_let_tv8_ = p0
        pat_let_tv9_ = d_432_v5_
        pat_let_tv10_ = d_432_v5_
        def lambda24_(source5_):
            if source5_.is_DC11:
                d_471___mcc_h2_ = source5_.cf13
                d_472___mcc_h3_ = source5_.cf14
                d_473_cf14_ = d_472___mcc_h3_
                d_474_cf13_ = d_471___mcc_h2_
                return _dafny.SeqWithoutIsStrInference([d_473_cf14_, pat_let_tv7_])
            elif source5_.is_DC12:
                d_475___mcc_h4_ = source5_.cf15
                d_476___mcc_h5_ = source5_.cf16
                d_477___mcc_h6_ = source5_.cf17
                d_478_cf17_ = d_477___mcc_h6_
                d_479_cf16_ = d_476___mcc_h5_
                d_480_cf15_ = d_475___mcc_h4_
                return pat_let_tv8_
            elif source5_.is_DC10:
                d_481___mcc_h7_ = source5_.cf12
                d_482_cf12_ = d_481___mcc_h7_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv9_])
            elif True:
                d_483___mcc_h8_ = source5_.cf18
                d_484_cf18_ = d_483___mcc_h8_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv10_ for d_485_i0_ in range(default__.abs(873))])

        r1 = lambda24_(d_470_v37_)
        r2 = (self).f8
        r3 = p1
        return r0, r1, r2, r3

    def m4(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        d_486_v0_: bool
        d_486_v0_ = False
        d_487_v1_: _dafny.Seq
        d_487_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xeoexah"))
        d_488_v2_: T0
        nw115_ = C1()
        nw115_.ctor__(d_486_v0_, (self).f7, len(d_487_v1_), (self).f7)
        d_488_v2_ = nw115_
        d_489_v3_: _dafny.Seq
        d_489_v3_ = _dafny.SeqWithoutIsStrInference([d_488_v2_])
        d_489_v3_ = d_489_v3_
        d_490_v4_: int
        d_490_v4_ = 96
        d_490_v4_ = default__.safeDivisionInt(d_490_v4_, -305)
        d_491_v5_: _dafny.Array
        nw116_ = _dafny.Array(_dafny.Array(None, 0), 6)
        d_491_v5_ = nw116_
        d_492_v6_: _dafny.Array
        nw117_ = _dafny.Array(_dafny.Map({}), 29)
        d_492_v6_ = nw117_
        index92_ = default__.safeIndex(189, (d_491_v5_).length(0))
        (d_491_v5_)[index92_] = d_492_v6_
        index93_ = default__.safeIndex(189, (d_491_v5_).length(0))
        (d_491_v5_)[index93_] = d_492_v6_
        hi3_ = d_490_v4_
        for d_493_i0_ in range(d_490_v4_, hi3_):
            d_494_v7_: _dafny.Map
            d_494_v7_ = _dafny.Map({d_486_v0_: d_486_v0_})
            d_495_v8_: _dafny.Array
            nw118_ = _dafny.Array(None, 25)
            nw118_[int(0)] = False
            nw118_[int(1)] = d_486_v0_
            nw118_[int(2)] = False
            nw118_[int(3)] = d_486_v0_
            nw118_[int(4)] = ((d_494_v7_)[((d_494_v7_)[d_486_v0_] if (d_486_v0_) in (d_494_v7_) else d_486_v0_)] if (((d_494_v7_)[d_486_v0_] if (d_486_v0_) in (d_494_v7_) else d_486_v0_)) in (d_494_v7_) else d_486_v0_)
            nw118_[int(5)] = d_486_v0_
            nw118_[int(6)] = d_486_v0_
            nw118_[int(7)] = d_486_v0_
            nw118_[int(8)] = d_486_v0_
            nw118_[int(9)] = d_486_v0_
            nw118_[int(10)] = d_486_v0_
            nw118_[int(11)] = d_486_v0_
            nw118_[int(12)] = d_486_v0_
            nw118_[int(13)] = d_486_v0_
            nw118_[int(14)] = d_486_v0_
            nw118_[int(15)] = False
            nw118_[int(16)] = d_486_v0_
            nw118_[int(17)] = True
            nw118_[int(18)] = d_486_v0_
            nw118_[int(19)] = not(d_486_v0_)
            nw118_[int(20)] = d_486_v0_
            nw118_[int(21)] = False
            nw118_[int(22)] = d_486_v0_
            nw118_[int(23)] = d_486_v0_
            nw118_[int(24)] = d_486_v0_
            d_495_v8_ = nw118_
            d_496_v9_: _dafny.Map
            d_496_v9_ = _dafny.Map({d_486_v0_: d_495_v8_})
            d_497_v10_: D5
            d_497_v10_ = D5_DC14(d_495_v8_)
            d_498_v11_: _dafny.Array
            nw119_ = _dafny.Array(None, 24)
            nw119_[int(0)] = d_495_v8_
            nw119_[int(1)] = d_495_v8_
            nw119_[int(2)] = d_495_v8_
            nw119_[int(3)] = d_495_v8_
            nw119_[int(4)] = d_495_v8_
            nw119_[int(5)] = d_495_v8_
            nw119_[int(6)] = d_495_v8_
            nw119_[int(7)] = ((d_496_v9_)[d_486_v0_] if (d_486_v0_) in (d_496_v9_) else d_495_v8_)
            nw119_[int(8)] = (d_497_v10_).cf19
            nw119_[int(9)] = d_495_v8_
            nw119_[int(10)] = d_495_v8_
            nw119_[int(11)] = d_495_v8_
            nw119_[int(12)] = (d_497_v10_).cf19
            nw119_[int(13)] = d_495_v8_
            nw119_[int(14)] = d_495_v8_
            nw119_[int(15)] = d_495_v8_
            nw119_[int(16)] = d_495_v8_
            nw119_[int(17)] = d_495_v8_
            nw119_[int(18)] = d_495_v8_
            nw119_[int(19)] = d_495_v8_
            nw119_[int(20)] = d_495_v8_
            nw119_[int(21)] = d_495_v8_
            nw119_[int(22)] = d_495_v8_
            nw119_[int(23)] = d_495_v8_
            d_498_v11_ = nw119_
            d_499_v12_: str
            d_499_v12_ = _dafny.CodePoint('y')
            d_500_v13_: _dafny.Map
            d_500_v13_ = _dafny.Map({p0: len((d_487_v1_).set(default__.safeIndex(-796, len(d_487_v1_)), d_499_v12_))})
            d_501_v14_: _dafny.Set
            d_501_v14_ = _dafny.Set({d_487_v1_, d_487_v1_})
            rhs71_ = d_498_v11_
            rhs72_ = ((0) - (len(d_500_v13_))) - (default__.safeModuloInt((d_488_v2_).f8, len(d_501_v14_)))
            rhs73_ = (d_488_v2_).f8
            d_498_v11_ = rhs71_
            d_490_v4_ = rhs72_
            d_490_v4_ = rhs73_
            d_502_v15_: _dafny.Map
            d_502_v15_ = _dafny.Map({(d_488_v2_).f8: _dafny.SeqWithoutIsStrInference([p0])})
            d_502_v15_ = (d_502_v15_).set((d_488_v2_).f8, _dafny.SeqWithoutIsStrInference([p0]))
            (globalState).f3 = default__.fm2(d_486_v0_, globalState)
            d_495_v8_ = (d_495_v8_ if d_486_v0_ else d_495_v8_)
        d_503_v16_: str
        d_503_v16_ = _dafny.CodePoint('s')
        d_504_v17_: C2
        nw120_ = C2()
        nw120_.ctor__(default__.fm0(d_503_v16_, globalState), d_490_v4_, (d_488_v2_).f8, 250)
        d_504_v17_ = nw120_
        d_505_v18_: C2
        nw121_ = C2()
        nw121_.ctor__((d_488_v2_).f8, d_490_v4_, (d_488_v2_).f8, (self).f7)
        d_505_v18_ = nw121_
        r0 = ((d_492_v6_ if d_486_v0_ else (d_491_v5_)[default__.safeIndex(189, (d_491_v5_).length(0))]) if d_486_v0_ else (d_491_v5_)[default__.safeIndex(189, (d_491_v5_).length(0))])
        d_506_v19_: _dafny.MultiSet
        d_506_v19_ = _dafny.MultiSet([not(d_486_v0_)])
        r1 = ((d_486_v0_ if d_486_v0_ else d_486_v0_)) in (d_506_v19_)
        d_507_v20_: _dafny.Seq
        d_507_v20_ = _dafny.SeqWithoutIsStrInference([996, d_505_v18_.f10])
        d_508_v21_: D6
        d_508_v21_ = D6_DC16(d_507_v20_)
        def iife43_():
            coll19_ = _dafny.Map()
            compr_19_: D1
            for compr_19_ in (_dafny.Map({D1_DC3(_dafny.CodePoint('n')): (self).f8})).keys.Elements:
                d_509_v22_: D1 = compr_19_
                if (d_509_v22_) in (_dafny.Map({D1_DC3(_dafny.CodePoint('n')): (self).f8})):
                    coll19_[d_509_v22_] = d_486_v0_
            return _dafny.Map(coll19_)
        r2 = ((d_507_v20_) + ((d_508_v21_).cf22)) + (_dafny.SeqWithoutIsStrInference([len(iife43_()
) for d_510_i1_ in range(default__.abs(786))]))
        return r0, r1, r2


class C4:
    def  __init__(self):
        self._f6: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f6):
        (self)._f6 = f6

    def m1(self, globalState):
        d_511_v0_: D0
        d_511_v0_ = D0_DC1(_dafny.CodePoint('e'))
        d_512_v1_: bool
        d_512_v1_ = True
        d_513_v2_: str
        d_513_v2_ = _dafny.CodePoint('t')
        d_514_v3_: D1
        d_514_v3_ = D1_DC2(d_512_v1_)
        d_515_v4_: _dafny.Array
        nw122_ = _dafny.Array(None, 23)
        nw122_[int(0)] = d_511_v0_
        nw122_[int(1)] = d_511_v0_
        nw122_[int(2)] = d_511_v0_
        nw122_[int(3)] = d_511_v0_
        nw122_[int(4)] = d_511_v0_
        nw122_[int(5)] = d_511_v0_
        nw122_[int(6)] = default__.fm5(d_512_v1_, d_512_v1_, d_512_v1_, globalState)
        nw122_[int(7)] = D0_DC1(d_513_v2_)
        nw122_[int(8)] = D0_DC1(d_513_v2_)
        nw122_[int(9)] = D0_DC1(d_513_v2_)
        nw122_[int(10)] = d_511_v0_
        nw122_[int(11)] = d_511_v0_
        nw122_[int(12)] = default__.fm5(d_512_v1_, default__.fm2((d_514_v3_).cf2, globalState), d_512_v1_, globalState)
        nw122_[int(13)] = D0_DC1(d_513_v2_)
        nw122_[int(14)] = d_511_v0_
        nw122_[int(15)] = default__.fm5(d_512_v1_, d_512_v1_, False, globalState)
        nw122_[int(16)] = d_511_v0_
        nw122_[int(17)] = d_511_v0_
        nw122_[int(18)] = default__.fm5(d_512_v1_, d_512_v1_, d_512_v1_, globalState)
        nw122_[int(19)] = d_511_v0_
        nw122_[int(20)] = d_511_v0_
        nw122_[int(21)] = d_511_v0_
        nw122_[int(22)] = D0_DC1(d_513_v2_)
        d_515_v4_ = nw122_
        d_516_v5_: int
        d_516_v5_ = -361
        d_517_v6_: _dafny.MultiSet
        d_517_v6_ = _dafny.MultiSet([d_516_v5_])
        d_518_v7_: _dafny.Seq
        d_518_v7_ = _dafny.SeqWithoutIsStrInference([d_512_v1_, d_512_v1_])
        d_519_v8_: _dafny.MultiSet
        d_519_v8_ = _dafny.MultiSet([(d_517_v6_).cardinality, len(d_518_v7_)])
        d_520_v9_: _dafny.Map
        d_520_v9_ = _dafny.Map({d_515_v4_: (d_517_v6_).cardinality})
        d_520_v9_ = (d_520_v9_).set(d_515_v4_, 59)
        d_512_v1_ = d_512_v1_
        d_521_i0_: int
        d_521_i0_ = 0
        with _dafny.label("2"):
            while d_512_v1_:
                with _dafny.c_label("2"):
                    if (d_521_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_521_i0_ = (d_521_i0_) + (1)
                    d_522_v10_: _dafny.Array
                    def lambda25_(d_523_v5_):
                        def lambda26_(d_524_i1_):
                            return default__.safeModuloInt(d_524_i1_, (0) - (d_523_v5_))

                        return lambda26_

                    init13_ = lambda25_(d_516_v5_)
                    nw123_ = _dafny.Array(None, 18)
                    for i0_13_ in range(nw123_.length(0)):
                        nw123_[i0_13_] = init13_(i0_13_)
                    d_522_v10_ = nw123_
                    d_525_v11_: D0
                    d_525_v11_ = D0_DC0(d_522_v10_)
                    d_526_v12_: _dafny.Seq
                    d_526_v12_ = _dafny.SeqWithoutIsStrInference([d_522_v10_])
                    pat_let_tv11_ = d_526_v12_
                    pat_let_tv12_ = d_516_v5_
                    pat_let_tv13_ = d_526_v12_
                    def iife44_(_pat_let12_0):
                        def iife45_(d_527_dt__update__tmp_h0_):
                            def iife46_(_pat_let13_0):
                                def iife47_(d_528_dt__update_hcf0_h0_):
                                    return D0_DC0(d_528_dt__update_hcf0_h0_)
                                return iife47_(_pat_let13_0)
                            return iife46_((pat_let_tv11_)[default__.safeIndex(pat_let_tv12_, len(pat_let_tv13_))])
                        return iife45_(_pat_let12_0)
                    source6_ = iife44_(d_525_v11_)
                    if source6_.is_DC1:
                        d_529___mcc_h0_ = source6_.cf1
                        d_530_cf1_ = d_529___mcc_h0_
                        d_531_v13_: C3
                        nw124_ = C3()
                        nw124_.ctor__(d_516_v5_, (d_516_v5_ if d_512_v1_ else d_516_v5_))
                        d_531_v13_ = nw124_
                        d_532_v14_: _dafny.Map
                        d_532_v14_ = _dafny.Map({524: d_516_v5_})
                        d_533_v15_: _dafny.Map
                        d_533_v15_ = _dafny.Map({d_512_v1_: _dafny.Map({d_516_v5_: d_516_v5_})})
                        d_534_v17_: _dafny.Seq
                        def iife48_():
                            coll20_ = _dafny.Map()
                            compr_20_: int
                            for compr_20_ in _dafny.IntegerRange(856, 532):
                                d_535_v16_: int = compr_20_
                                if ((856) <= (d_535_v16_)) and ((d_535_v16_) < (532)):
                                    coll20_[default__.safeModuloInt(d_535_v16_, d_516_v5_)] = d_516_v5_
                            return _dafny.Map(coll20_)
                        d_534_v17_ = _dafny.SeqWithoutIsStrInference([((d_533_v15_)[d_512_v1_] if (d_512_v1_) in (d_533_v15_) else iife48_()
                        )])
                        d_536_v18_: _dafny.Map
                        d_536_v18_ = _dafny.Map({d_513_v2_: not((d_532_v14_) in (d_534_v17_))})
                        d_536_v18_ = (d_536_v18_).set(d_513_v2_, True)
                        d_537_v19_: _dafny.Array
                        def lambda27_(d_538_v1_):
                            def lambda28_(d_539_i2_):
                                return d_538_v1_

                            return lambda28_

                        init14_ = lambda27_(d_512_v1_)
                        nw125_ = _dafny.Array(None, 26)
                        for i0_14_ in range(nw125_.length(0)):
                            nw125_[i0_14_] = init14_(i0_14_)
                        d_537_v19_ = nw125_
                        index94_ = default__.safeIndex(679, (d_537_v19_).length(0))
                        (d_537_v19_)[index94_] = (d_512_v1_ if d_512_v1_ else d_512_v1_)
                        index95_ = default__.safeIndex(679, (d_537_v19_).length(0))
                        (d_537_v19_)[index95_] = d_512_v1_
                        d_540_v20_: _dafny.Seq
                        d_540_v20_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjhkseht"))
                        d_541_v21_: D6
                        d_541_v21_ = D6_DC17(d_540_v20_, d_512_v1_, d_516_v5_, False, not((d_537_v19_)[default__.safeIndex(679, (d_537_v19_).length(0))]))
                        (globalState).f3 = (d_540_v20_) != ((d_541_v21_).cf23)
                    elif True:
                        d_542___mcc_h1_ = source6_.cf0
                        d_543_cf0_ = d_542___mcc_h1_
                        d_544_v22_: _dafny.Seq
                        d_544_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jsstabmc"))
                        d_544_v22_ = (d_544_v22_) + ((d_544_v22_ if d_512_v1_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))))
                        d_545_v23_: _dafny.Array
                        def lambda29_(d_546_v1_):
                            def lambda30_(d_547_i3_):
                                return d_546_v1_

                            return lambda30_

                        init15_ = lambda29_(d_512_v1_)
                        nw126_ = _dafny.Array(None, 2)
                        for i0_15_ in range(nw126_.length(0)):
                            nw126_[i0_15_] = init15_(i0_15_)
                        d_545_v23_ = nw126_
                        d_545_v23_ = d_545_v23_
                        d_544_v22_ = default__.fm11(d_512_v1_, d_516_v5_, globalState)
                        d_516_v5_ = (0) - (d_516_v5_)
                    d_513_v2_ = d_513_v2_
                    d_548_v24_: _dafny.Array
                    nw127_ = _dafny.Array(None, 26)
                    nw127_[int(0)] = d_512_v1_
                    nw127_[int(1)] = d_512_v1_
                    nw127_[int(2)] = not(d_512_v1_)
                    nw127_[int(3)] = d_512_v1_
                    nw127_[int(4)] = d_512_v1_
                    nw127_[int(5)] = d_512_v1_
                    nw127_[int(6)] = default__.fm2(d_512_v1_, globalState)
                    nw127_[int(7)] = False
                    nw127_[int(8)] = d_512_v1_
                    nw127_[int(9)] = d_512_v1_
                    nw127_[int(10)] = d_512_v1_
                    nw127_[int(11)] = d_512_v1_
                    nw127_[int(12)] = d_512_v1_
                    nw127_[int(13)] = True
                    nw127_[int(14)] = not(False)
                    nw127_[int(15)] = d_512_v1_
                    nw127_[int(16)] = d_512_v1_
                    nw127_[int(17)] = d_512_v1_
                    nw127_[int(18)] = d_512_v1_
                    nw127_[int(19)] = d_512_v1_
                    nw127_[int(20)] = True
                    nw127_[int(21)] = d_512_v1_
                    nw127_[int(22)] = default__.fm2(d_512_v1_, globalState)
                    nw127_[int(23)] = d_512_v1_
                    nw127_[int(24)] = d_512_v1_
                    nw127_[int(25)] = not(not(d_512_v1_))
                    d_548_v24_ = nw127_
                    d_549_v25_: _dafny.MultiSet
                    d_549_v25_ = _dafny.MultiSet([d_548_v24_, d_548_v24_])
                    d_550_v26_: _dafny.Map
                    d_550_v26_ = _dafny.Map({d_549_v25_: d_512_v1_})
                    d_550_v26_ = (d_550_v26_).set(d_549_v25_, (d_516_v5_) == (d_516_v5_))
                    d_551_v27_: _dafny.Seq
                    d_551_v27_ = _dafny.SeqWithoutIsStrInference([d_516_v5_])
                    index96_ = default__.safeIndex(829, (d_522_v10_).length(0))
                    (d_522_v10_)[index96_] = len((_dafny.SeqWithoutIsStrInference([d_513_v2_ for d_552_i4_ in range(default__.abs(-759))])) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))).set(default__.safeIndex(len(d_551_v27_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t")))), d_513_v2_)))
                    d_553_v28_: _dafny.MultiSet
                    d_553_v28_ = _dafny.MultiSet([_dafny.Map({len(d_518_v7_): d_516_v5_})])
                    index97_ = default__.safeIndex(829, (d_522_v10_).length(0))
                    (d_522_v10_)[index97_] = ((d_553_v28_).intersection((d_553_v28_) | (default__.fm18(globalState)))).cardinality
                    pass
            pass
        d_554_v29_: _dafny.Seq
        d_554_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "geidfsen"))
        d_555_v30_: _dafny.Map
        d_555_v30_ = _dafny.Map({d_516_v5_: 188})
        pat_let_tv14_ = d_512_v1_
        d_556_v32_: D6
        def iife50_():
            coll21_ = _dafny.Set()
            compr_21_: int
            for compr_21_ in (d_555_v30_).keys.Elements:
                d_557_v31_: int = compr_21_
                if (d_557_v31_) in (d_555_v30_):
                    coll21_ = coll21_.union(_dafny.Set([(d_557_v31_) - ((_dafny.MultiSet([143])).cardinality)]))
            return _dafny.Set(coll21_)
        def iife49_(_pat_let14_0):
            def iife51_(d_558_dt__update__tmp_h1_):
                def iife52_(_pat_let15_0):
                    def iife53_(d_559_dt__update_hcf17_h0_):
                        return D4_DC12((d_558_dt__update__tmp_h1_).cf15, (d_558_dt__update__tmp_h1_).cf16, d_559_dt__update_hcf17_h0_)
                    return iife53_(_pat_let15_0)
                return iife52_(not(pat_let_tv14_))
            return iife51_(_pat_let14_0)
        d_556_v32_ = D6_DC17(d_554_v29_, (iife49_(D4_DC12(len(iife50_()
), _dafny.SeqWithoutIsStrInference([d_512_v1_]), False))).cf17, ((d_519_v8_)[d_516_v5_] if (d_516_v5_) in (d_519_v8_) else d_516_v5_), (d_518_v7_)[default__.safeIndex((0) - (d_516_v5_), len(d_518_v7_))], False)
        d_512_v1_ = (d_556_v32_).cf27
        d_560_v33_: _dafny.Array
        def lambda31_(d_561_v1_):
            def lambda32_(d_562_i5_):
                return d_561_v1_

            return lambda32_

        init16_ = lambda31_(d_512_v1_)
        nw128_ = _dafny.Array(None, 10)
        for i0_16_ in range(nw128_.length(0)):
            nw128_[i0_16_] = init16_(i0_16_)
        d_560_v33_ = nw128_
        index98_ = default__.safeIndex(61, (d_560_v33_).length(0))
        (d_560_v33_)[index98_] = d_512_v1_
        index99_ = default__.safeIndex(61, (d_560_v33_).length(0))
        (d_560_v33_)[index99_] = not((d_518_v7_)[default__.safeIndex(d_516_v5_, len(d_518_v7_))])
        (globalState).f1 = (d_560_v33_)[default__.safeIndex(61, (d_560_v33_).length(0))]

    def m2(self, p0, p1, p2, globalState):
        d_563_v0_: _dafny.Array
        nw129_ = _dafny.Array(int(0), 27)
        d_563_v0_ = nw129_
        d_564_v1_: D0
        d_564_v1_ = D0_DC0(d_563_v0_)
        d_565_v2_: _dafny.Array
        nw130_ = _dafny.Array(False, 6)
        d_565_v2_ = nw130_
        d_566_v3_: _dafny.Map
        d_566_v3_ = _dafny.Map({(d_564_v1_).cf0: d_565_v2_})
        d_566_v3_ = (d_566_v3_).set(d_563_v0_, d_565_v2_)
        d_567_v4_: _dafny.Array
        nw131_ = _dafny.Array(_dafny.Map({}), 9)
        d_567_v4_ = nw131_
        d_568_v5_: int
        d_568_v5_ = -872
        d_569_v6_: _dafny.Seq
        d_569_v6_ = _dafny.SeqWithoutIsStrInference([d_568_v5_])
        d_570_v7_: _dafny.Map
        d_570_v7_ = _dafny.Map({d_568_v5_: len((D6_DC16(d_569_v6_)).cf22)})
        index100_ = default__.safeIndex(285, (d_567_v4_).length(0))
        (d_567_v4_)[index100_] = d_570_v7_
        d_571_v8_: str
        d_571_v8_ = _dafny.CodePoint('a')
        d_572_v9_: _dafny.MultiSet
        d_572_v9_ = _dafny.MultiSet([d_563_v0_, d_563_v0_, d_563_v0_, d_563_v0_, d_563_v0_])
        d_573_v10_: _dafny.Seq
        d_573_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jmjl"))
        d_574_v11_: _dafny.Map
        d_574_v11_ = _dafny.Map({True: p2})
        index101_ = default__.safeIndex(285, (d_567_v4_).length(0))
        rhs74_ = d_570_v7_
        rhs75_ = d_571_v8_
        rhs76_ = (D4_DC12(len(d_573_v10_), p1, p2)).cf17
        rhs77_ = len(d_574_v11_)
        rhs78_ = _dafny.MultiSet([d_563_v0_])
        lhs63_ = d_567_v4_
        lhs64_ = default__.safeIndex(285, (d_567_v4_).length(0))
        lhs65_ = globalState
        lhs63_[lhs64_] = rhs74_
        d_571_v8_ = rhs75_
        lhs65_.f3 = rhs76_
        d_568_v5_ = rhs77_
        d_572_v9_ = rhs78_
        d_575_v12_: _dafny.MultiSet
        d_575_v12_ = _dafny.MultiSet([d_568_v5_, d_568_v5_])
        d_576_v13_: _dafny.MultiSet
        d_576_v13_ = _dafny.MultiSet([d_571_v8_, d_571_v8_])
        d_577_v14_: _dafny.MultiSet
        d_577_v14_ = _dafny.MultiSet([((d_575_v12_)[d_568_v5_] if (d_568_v5_) in (d_575_v12_) else (d_576_v13_).cardinality)])
        d_578_v15_: _dafny.Set
        d_578_v15_ = _dafny.Set({d_575_v12_})
        d_579_v16_: D3
        d_579_v16_ = D3_DC9(D3_DC7(d_578_v15_))
        d_579_v16_ = default__.fm19(globalState)
        (globalState).f1 = ((_dafny.SeqWithoutIsStrInference([d_571_v8_ for d_580_i0_ in range(default__.abs(169))])) != (d_573_v10_)) or (not (p2) or (p2))
        if not((p1)[default__.safeIndex(d_568_v5_, len(p1))]):
            index102_ = default__.safeIndex(747, (d_565_v2_).length(0))
            (d_565_v2_)[index102_] = p2
            index103_ = default__.safeIndex(776, (d_565_v2_).length(0))
            (d_565_v2_)[index103_] = not((p1)[default__.safeIndex(d_568_v5_, len(p1))])
            index104_ = default__.safeIndex(747, (d_565_v2_).length(0))
            index105_ = default__.safeIndex(776, (d_565_v2_).length(0))
            rhs79_ = p2
            rhs80_ = len((_dafny.SeqWithoutIsStrInference([d_571_v8_ for d_581_i1_ in range(default__.abs(311))])) + (d_573_v10_))
            rhs81_ = default__.fm0(d_571_v8_, globalState)
            rhs82_ = d_568_v5_
            rhs83_ = p2
            lhs66_ = d_565_v2_
            lhs67_ = default__.safeIndex(747, (d_565_v2_).length(0))
            lhs68_ = d_565_v2_
            lhs69_ = default__.safeIndex(776, (d_565_v2_).length(0))
            lhs66_[lhs67_] = rhs79_
            d_568_v5_ = rhs80_
            d_568_v5_ = rhs81_
            d_568_v5_ = rhs82_
            lhs68_[lhs69_] = rhs83_
            (globalState).f1 = (d_568_v5_) != ((d_568_v5_) - (d_568_v5_))
            index106_ = default__.safeIndex(747, (d_565_v2_).length(0))
            (d_565_v2_)[index106_] = not(p2)
            (globalState).f1 = default__.fm2((d_565_v2_)[default__.safeIndex(747, (d_565_v2_).length(0))], globalState)
            index107_ = default__.safeIndex(285, (d_567_v4_).length(0))
            (d_567_v4_)[index107_] = ((d_567_v4_)[default__.safeIndex(285, (d_567_v4_).length(0))]).set(d_568_v5_, default__.safeModuloInt(d_568_v5_, d_568_v5_))
        elif True:
            d_582_v17_: C2
            nw132_ = C2()
            nw132_.ctor__(d_568_v5_, (_dafny.MultiSet(d_569_v6_)).cardinality, d_568_v5_, d_568_v5_)
            d_582_v17_ = nw132_
            (globalState).f3 = default__.fm2(False, globalState)
            (globalState).f3 = p2
            d_583_v18_: D6
            d_583_v18_ = D6_DC17(d_573_v10_, False, d_582_v17_.f10, p2, p2)
            d_568_v5_ = default__.safeModuloInt(default__.safeDivisionInt((d_582_v17_).f9, (d_582_v17_).f9), len((_dafny.SeqWithoutIsStrInference([d_571_v8_ for d_584_i2_ in range(default__.abs(432))])) + ((d_583_v18_).cf23)))
            d_585_v19_: D3
            d_585_v19_ = D3_DC8(d_582_v17_.f10)
            pat_let_tv15_ = d_582_v17_
            d_586_v20_: _dafny.MultiSet
            def iife54_(_pat_let16_0):
                def iife55_(d_587_dt__update__tmp_h0_):
                    def iife56_(_pat_let17_0):
                        def iife57_(d_588_dt__update_hcf10_h0_):
                            return D3_DC8(d_588_dt__update_hcf10_h0_)
                        return iife57_(_pat_let17_0)
                    return iife56_((pat_let_tv15_).f9)
                return iife55_(_pat_let16_0)
            d_586_v20_ = _dafny.MultiSet([d_585_v19_, d_585_v19_, iife54_(d_585_v19_)])
            d_589_v21_: D4
            d_589_v21_ = D4_DC11((_dafny.MultiSet([d_568_v5_, (d_582_v17_).f9])).cardinality, d_571_v8_)
            d_590_v22_: _dafny.Set
            d_590_v22_ = _dafny.Set({True})
            d_591_v23_: D2
            d_591_v23_ = D2_DC4(776)
            d_592_v24_: _dafny.Seq
            d_592_v24_ = _dafny.SeqWithoutIsStrInference([d_590_v22_, default__.fm20(d_591_v23_, p2, globalState)])
            d_593_v25_: _dafny.Array
            nw133_ = _dafny.Array(None, 24)
            nw133_[int(0)] = d_573_v10_
            nw133_[int(1)] = d_573_v10_
            nw133_[int(2)] = (d_573_v10_).set(default__.safeIndex(172, len(d_573_v10_)), d_571_v8_)
            nw133_[int(3)] = ((d_573_v10_).set(default__.safeIndex(((d_586_v20_)[d_585_v19_] if (d_585_v19_) in (d_586_v20_) else (0) - ((d_589_v21_).cf13)), len(d_573_v10_)), d_571_v8_)).set(default__.safeIndex(len((d_592_v24_)[default__.safeIndex(d_582_v17_.f10, len(d_592_v24_))]), len((d_573_v10_).set(default__.safeIndex(((d_586_v20_)[d_585_v19_] if (d_585_v19_) in (d_586_v20_) else (0) - ((d_589_v21_).cf13)), len(d_573_v10_)), d_571_v8_))), d_571_v8_)
            nw133_[int(4)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))) + (d_573_v10_)
            nw133_[int(5)] = _dafny.SeqWithoutIsStrInference([d_571_v8_ for d_594_i3_ in range(default__.abs(-519))])
            nw133_[int(6)] = d_573_v10_
            nw133_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sval"))
            nw133_[int(8)] = d_573_v10_
            nw133_[int(9)] = d_573_v10_
            nw133_[int(10)] = (d_573_v10_) + (_dafny.SeqWithoutIsStrInference([d_571_v8_ for d_595_i4_ in range(default__.abs(743))]))
            nw133_[int(11)] = d_573_v10_
            nw133_[int(12)] = (d_573_v10_) + (d_573_v10_)
            nw133_[int(13)] = d_573_v10_
            nw133_[int(14)] = d_573_v10_
            nw133_[int(15)] = d_573_v10_
            nw133_[int(16)] = d_573_v10_
            nw133_[int(17)] = d_573_v10_
            nw133_[int(18)] = (d_573_v10_) + ((d_573_v10_).set(default__.safeIndex(d_568_v5_, len(d_573_v10_)), d_571_v8_))
            nw133_[int(19)] = d_573_v10_
            nw133_[int(20)] = d_573_v10_
            nw133_[int(21)] = (d_573_v10_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tswjrvwl")))
            nw133_[int(22)] = (d_573_v10_) + (d_573_v10_)
            nw133_[int(23)] = d_573_v10_
            d_593_v25_ = nw133_
            index108_ = default__.safeIndex(376, (d_593_v25_).length(0))
            (d_593_v25_)[index108_] = d_573_v10_
            index109_ = default__.safeIndex(376, (d_593_v25_).length(0))
            (d_593_v25_)[index109_] = (d_573_v10_ if not(not (p2) or (not(not(p2)))) else (d_573_v10_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hupyfkie"))))
        d_596_v26_: _dafny.Array
        nw134_ = _dafny.Array(None, 4)
        d_596_v26_ = nw134_
        d_597_v27_: C3
        nw135_ = C3()
        nw135_.ctor__((0) - (len(_dafny.Set({p2, True}))), 383)
        d_597_v27_ = nw135_
        index110_ = default__.safeIndex(884, (d_596_v26_).length(0))
        (d_596_v26_)[index110_] = d_597_v27_
        index111_ = default__.safeIndex(884, (d_596_v26_).length(0))
        (d_596_v26_)[index111_] = d_597_v27_

    @property
    def f6(self):
        return self._f6
