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
        return (D0_DC2(False, False, -941)).cf4

    @staticmethod
    def fm1(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csbalwk"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_0_i0_ in range(default__.abs(-85))]))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1_i1_ in range(default__.abs(394))]))

    @staticmethod
    def fm2(globalState):
        if not (not(True)) or (False):
            return (_dafny.Set({-71})).isdisjoint(_dafny.Set({-563}))
        elif True:
            return False

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False, not(not(not(False)))])) + (_dafny.SeqWithoutIsStrInference([True]))) + (_dafny.SeqWithoutIsStrInference([not(not(True)), True]))

    @staticmethod
    def fm7(p0, globalState):
        return ((_dafny.Map({False: False})) | (_dafny.Map({True: False}))) | ((_dafny.Map({True: True})) | (_dafny.Map({False: not(True)})))

    @staticmethod
    def fm8(p0, globalState):
        return (_dafny.Map({450: False})) | (_dafny.Map({-488: False}))

    @staticmethod
    def fm9(globalState):
        return ((_dafny.Set({False, not(True)})) - (_dafny.Set({False, True}))).intersection(_dafny.Set({False}))

    @staticmethod
    def fm10(p0, p1, globalState):
        if not(not(False)):
            if True:
                return D0_DC2(False, True, (0) - (len(_dafny.Map({981: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2_i0_ in range(default__.abs(758))]))}))))
            elif True:
                return D0_DC2(not(False), (D3_DC8(431, -307, True)).cf19, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))).cardinality)
        elif True:
            return D0_DC2((D0_DC2(not(True), True, len(_dafny.SeqWithoutIsStrInference([943])))).cf3, True, 567)

    @staticmethod
    def fm11(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: _dafny.Map
            for compr_0_ in (_dafny.Set({_dafny.Map({False: False}), _dafny.Map({True: True})})).Elements:
                d_3_v0_: _dafny.Map = compr_0_
                if (d_3_v0_) in (_dafny.Set({_dafny.Map({False: False}), _dafny.Map({True: True})})):
                    coll0_ = coll0_.union(_dafny.Set([d_3_v0_]))
            return _dafny.Set(coll0_)
        return D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ft")), (len(_dafny.Set({True}))) + ((D3_DC8(len(_dafny.Set({len(_dafny.Set({True, True}))})), 798, not(not(True)))).cf17), (_dafny.Set({_dafny.Map({True: True})})).ispropersubset(iife0_()
), ((_dafny.MultiSet([True, True]) if True else _dafny.MultiSet([(D3_DC8(171, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "otasatp"))), False)).cf19, False, False]))).cardinality, (_dafny.MultiSet([len(_dafny.Map({(0) - ((0) - ((_dafny.MultiSet([False])).cardinality)): _dafny.SeqWithoutIsStrInference([False])})), 118])).isdisjoint(_dafny.MultiSet([587])))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        return _dafny.MultiSet([239])

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        if not((False) == (not(not(True)))):
            def iife1_():
                coll1_ = _dafny.Map()
                compr_1_: D3
                for compr_1_ in (_dafny.MultiSet([D3_DC8(764, 168, not(False))])).Elements:
                    d_4_v0_: D3 = compr_1_
                    if (d_4_v0_) in (_dafny.MultiSet([D3_DC8(764, 168, not(False))])):
                        coll1_[d_4_v0_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u')])
                return _dafny.Map(coll1_)
            return D4_DC9(iife1_()
)
        elif True:
            return D4_DC9(_dafny.Map({D3_DC8(-664, len(_dafny.SeqWithoutIsStrInference([90])), True): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])}))
        elif True:
            return D4_DC9(_dafny.Map({D3_DC8(len(_dafny.Map({-837: True})), len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i')]))})), True): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_5_i0_ in range(default__.abs(82))])}))

    @staticmethod
    def fm14(p0, p1, globalState):
        return _dafny.Set({305})

    @staticmethod
    def fm15(p0, p1, globalState):
        return D0_DC0(_dafny.SeqWithoutIsStrInference([not(False), True]))

    @staticmethod
    def fm16(p0, p1, globalState):
        return _dafny.MultiSet([not(True), not((not(not((D0_DC2(True, False, 735)).cf2))) and (False))])

    @staticmethod
    def fm17(p0, p1, p2, globalState):
        source0_ = (D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mjffy")), (_dafny.MultiSet([False])).cardinality, False, len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference([False, False, not(False)])})), False) if not(False) else D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")), 139, not(True), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dop"))), False))
        if source0_.is_DC6:
            d_6___mcc_h0_ = source0_.cf11
            d_7___mcc_h1_ = source0_.cf12
            d_8___mcc_h2_ = source0_.cf13
            d_9___mcc_h3_ = source0_.cf14
            d_10___mcc_h4_ = source0_.cf15
            d_11_cf15_ = d_10___mcc_h4_
            d_12_cf14_ = d_9___mcc_h3_
            d_13_cf13_ = d_8___mcc_h2_
            d_14_cf12_ = d_7___mcc_h1_
            d_15_cf11_ = d_6___mcc_h0_
            def iife2_():
                coll2_ = _dafny.Map()
                compr_2_: int
                for compr_2_ in _dafny.IntegerRange(790, 970):
                    d_16_v0_: int = compr_2_
                    if ((790) <= (d_16_v0_)) and ((d_16_v0_) < (970)):
                        coll2_[default__.safeDivisionInt(d_16_v0_, d_12_cf14_)] = 857
                return _dafny.Map(coll2_)
            return (_dafny.SeqWithoutIsStrInference([iife2_()
 for d_17_i0_ in range(default__.abs(780))])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({-290: d_14_cf12_}) for d_18_i1_ in range(default__.abs(701))]))
        elif True:
            d_19___mcc_h5_ = source0_.cf10
            d_20_cf10_ = d_19___mcc_h5_
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(-768, 394):
                    d_21_v1_: int = compr_3_
                    if ((-768) <= (d_21_v1_)) and ((d_21_v1_) < (394)):
                        coll3_[(d_21_v1_) * (-880)] = 586
                return _dafny.Map(coll3_)
            return _dafny.SeqWithoutIsStrInference([(iife3_()
) for d_22_i2_ in range(default__.abs(980))])

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        if p0:
            d_23_v0_: bool
            d_23_v0_ = True
            d_24_v1_: _dafny.Seq
            d_24_v1_ = _dafny.SeqWithoutIsStrInference([p3])
            d_23_v0_ = (p1) not in (d_24_v1_)
            d_25_v2_: D3
            d_25_v2_ = D3_DC8(-73, p3, d_23_v0_)
            d_26_v3_: C0
            nw0_ = C0()
            nw0_.ctor__(d_25_v2_)
            d_26_v3_ = nw0_
            d_27_v4_: _dafny.Seq
            d_27_v4_ = _dafny.SeqWithoutIsStrInference([d_23_v0_])
            d_28_v5_: _dafny.Map
            d_28_v5_ = _dafny.Map({len(d_27_v4_): len(d_27_v4_)})
            d_29_v6_: _dafny.Map
            d_29_v6_ = _dafny.Map({p2: ((_dafny.MultiSet([p0, p2])).set(p0, default__.abs(p3))).cardinality})
            d_30_v7_: _dafny.Array
            nw1_ = _dafny.Array(False, 13)
            d_30_v7_ = nw1_
            d_31_v8_: D4
            d_31_v8_ = D4_DC12(d_23_v0_, _dafny.SeqWithoutIsStrInference([d_28_v5_, d_28_v5_, d_28_v5_, _dafny.Map({p3: len(d_29_v6_)}), d_28_v5_]), p1, d_30_v7_, p0)
            d_32_v9_: _dafny.MultiSet
            d_32_v9_ = _dafny.MultiSet([_dafny.Map({d_23_v0_: (d_31_v8_).cf28}), d_29_v6_, d_29_v6_])
            d_33_v10_: D4
            d_33_v10_ = D4_DC11(d_26_v3_, ((d_32_v9_)[_dafny.Map({d_23_v0_: (d_26_v3_).fm6(p0, p1, not(True), globalState)})] if (_dafny.Map({d_23_v0_: (d_26_v3_).fm6(p0, p1, not(True), globalState)})) in (d_32_v9_) else p3), p3)
            d_33_v10_ = d_33_v10_
            index0_ = default__.safeIndex(74, (d_30_v7_).length(0))
            (d_30_v7_)[index0_] = False
            index1_ = default__.safeIndex(74, (d_30_v7_).length(0))
            (d_30_v7_)[index1_] = default__.fm2(globalState)
            index2_ = default__.safeIndex(74, (d_30_v7_).length(0))
            (d_30_v7_)[index2_] = p0
            if (p2) and (default__.fm2(globalState)):
                d_34_v11_: C1
                nw2_ = C1()
                nw2_.ctor__()
                d_34_v11_ = nw2_
                d_35_v12_: _dafny.Array
                nw3_ = _dafny.Array(_dafny.Array(None, 0), 2)
                d_35_v12_ = nw3_
                d_36_v13_: str
                d_36_v13_ = _dafny.CodePoint('j')
                d_37_v14_: _dafny.Map
                d_37_v14_ = _dafny.Map({d_36_v13_: d_27_v4_})
                d_38_v15_: D0
                d_38_v15_ = D0_DC0(_dafny.SeqWithoutIsStrInference([not(p2)]))
                d_39_v16_: _dafny.Array
                nw4_ = _dafny.Array(None, 16)
                nw4_[int(0)] = _dafny.SeqWithoutIsStrInference([p0])
                nw4_[int(1)] = d_27_v4_
                nw4_[int(2)] = d_27_v4_
                nw4_[int(3)] = _dafny.SeqWithoutIsStrInference([p2, (d_30_v7_)[default__.safeIndex(74, (d_30_v7_).length(0))]])
                nw4_[int(4)] = d_27_v4_
                nw4_[int(5)] = d_27_v4_
                nw4_[int(6)] = _dafny.SeqWithoutIsStrInference([False, True, d_23_v0_])
                nw4_[int(7)] = d_27_v4_
                nw4_[int(8)] = _dafny.SeqWithoutIsStrInference([True])
                nw4_[int(9)] = d_27_v4_
                nw4_[int(10)] = d_27_v4_
                nw4_[int(11)] = _dafny.SeqWithoutIsStrInference([d_23_v0_, (d_30_v7_)[default__.safeIndex(74, (d_30_v7_).length(0))]])
                nw4_[int(12)] = d_27_v4_
                nw4_[int(13)] = d_27_v4_
                nw4_[int(14)] = d_27_v4_
                nw4_[int(15)] = ((d_37_v14_)[d_36_v13_] if (d_36_v13_) in (d_37_v14_) else (d_38_v15_).cf0)
                d_39_v16_ = nw4_
                index3_ = default__.safeIndex(557, (d_35_v12_).length(0))
                (d_35_v12_)[index3_] = (d_39_v16_ if not((d_30_v7_)[default__.safeIndex(74, (d_30_v7_).length(0))]) else d_39_v16_)
                index4_ = default__.safeIndex(557, (d_35_v12_).length(0))
                (d_35_v12_)[index4_] = d_39_v16_
                d_40_v17_: _dafny.Seq
                d_40_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
                d_41_v18_: _dafny.Map
                d_41_v18_ = _dafny.Map({(p1) > (len(d_40_v17_)): _dafny.CodePoint('i')})
                d_42_v19_: str
                d_42_v19_ = d_36_v13_
                d_41_v18_ = (d_41_v18_).set(p2, (d_42_v19_))
                d_25_v2_ = d_25_v2_
                d_43_v20_: C0
                nw5_ = C0()
                nw5_.ctor__(d_25_v2_)
                d_43_v20_ = nw5_
            elif True:
                d_44_v21_: C1
                nw6_ = C1()
                nw6_.ctor__()
                d_44_v21_ = nw6_
                d_45_v22_: C1
                nw7_ = C1()
                nw7_.ctor__()
                d_45_v22_ = nw7_
                (globalState).f0 = p3
                d_46_v23_: _dafny.Seq
                d_46_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mmk"))
                d_46_v23_ = (d_46_v23_).set(default__.safeIndex(p1, len(d_46_v23_)), (d_46_v23_)[default__.safeIndex(p1, len(d_46_v23_))])
                d_47_v24_: _dafny.Seq
                d_47_v24_ = _dafny.SeqWithoutIsStrInference([d_26_v3_])
                d_48_v25_: D1
                d_48_v25_ = D1_DC4(p3, not(p0), True, False)
                d_49_v26_: _dafny.Array
                nw8_ = _dafny.Array(None, 19)
                nw8_[int(0)] = d_26_v3_
                nw8_[int(1)] = d_26_v3_
                nw8_[int(2)] = d_26_v3_
                nw8_[int(3)] = d_26_v3_
                nw8_[int(4)] = (d_47_v24_)[default__.safeIndex((d_48_v25_).cf6, len(d_47_v24_))]
                nw8_[int(5)] = d_26_v3_
                nw8_[int(6)] = d_26_v3_
                nw8_[int(7)] = d_26_v3_
                nw8_[int(8)] = d_26_v3_
                nw8_[int(9)] = d_26_v3_
                nw8_[int(10)] = d_26_v3_
                nw8_[int(11)] = d_26_v3_
                nw8_[int(12)] = d_26_v3_
                nw8_[int(13)] = d_26_v3_
                nw8_[int(14)] = d_26_v3_
                nw8_[int(15)] = d_26_v3_
                nw8_[int(16)] = d_26_v3_
                nw8_[int(17)] = d_26_v3_
                nw8_[int(18)] = d_26_v3_
                d_49_v26_ = nw8_
                d_50_v27_: _dafny.Map
                d_50_v27_ = _dafny.Map({d_49_v26_: d_23_v0_})
                d_50_v27_ = (d_50_v27_).set(d_49_v26_, p2)
        elif True:
            d_51_v28_: _dafny.Map
            d_51_v28_ = _dafny.Map({(D3_DC8(p3, p3, p0)).cf18: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "slgt")))})
            d_52_v29_: _dafny.MultiSet
            d_52_v29_ = _dafny.MultiSet([p0, (_dafny.Map({334: p3})) in (_dafny.Map({d_51_v28_: D1_DC4(p1, p2, not(p2), False)})), p2, p0, p0])
            d_52_v29_ = (d_52_v29_) - (d_52_v29_)
            if p2:
                (globalState).f0 = 70
                d_53_v30_: C0
                nw9_ = C0()
                nw9_.ctor__(D3_DC8(p3, default__.fm0(globalState), p2))
                d_53_v30_ = nw9_
                d_54_v31_: str
                d_54_v31_ = _dafny.CodePoint('c')
                d_55_v32_: _dafny.Array
                nw10_ = _dafny.Array(False, 15)
                d_55_v32_ = nw10_
                d_56_v33_: _dafny.Map
                d_56_v33_ = _dafny.Map({d_54_v31_: d_55_v32_})
                d_56_v33_ = (d_56_v33_).set(d_54_v31_, d_55_v32_)
                d_57_v34_: bool
                d_57_v34_ = True
                d_57_v34_ = not((d_57_v34_) or (p2))
                d_58_v35_: _dafny.Seq
                d_58_v35_ = _dafny.SeqWithoutIsStrInference([not(p0), p0])
                d_59_v36_: D4
                d_59_v36_ = D4_DC9(_dafny.Map({(d_53_v30_).f14: _dafny.SeqWithoutIsStrInference([d_54_v31_, d_54_v31_])}))
                d_60_v37_: _dafny.Seq
                d_60_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "osydgy"))
                d_61_v38_: _dafny.Map
                d_61_v38_ = _dafny.Map({(d_53_v30_).f14: d_60_v37_})
                d_62_v39_: _dafny.MultiSet
                d_62_v39_ = _dafny.MultiSet([p1, len(default__.fm1(p0, (0) - (p1), globalState)), p1])
                pat_let_tv0_ = d_61_v38_
                pat_let_tv1_ = d_61_v38_
                pat_let_tv2_ = d_61_v38_
                pat_let_tv3_ = d_61_v38_
                pat_let_tv4_ = d_61_v38_
                d_63_v40_: _dafny.Array
                nw11_ = _dafny.Array(None, 20)
                nw11_[int(0)] = d_59_v36_
                nw11_[int(1)] = d_59_v36_
                nw11_[int(2)] = default__.fm13(p3, d_60_v37_, d_57_v34_, globalState)
                nw11_[int(3)] = d_59_v36_
                nw11_[int(4)] = d_59_v36_
                nw11_[int(5)] = d_59_v36_
                nw11_[int(6)] = D4_DC9(d_61_v38_)
                nw11_[int(7)] = d_59_v36_
                nw11_[int(8)] = d_59_v36_
                nw11_[int(9)] = D4_DC9(d_61_v38_)
                nw11_[int(10)] = (d_59_v36_ if True else d_59_v36_)
                def iife5_(_pat_let1_0):
                    def iife6_(d_64_dt__update__tmp_h0_):
                        def iife7_(_pat_let2_0):
                            def iife8_(d_65_dt__update_hcf20_h0_):
                                return D4_DC9(d_65_dt__update_hcf20_h0_)
                            return iife8_(_pat_let2_0)
                        return iife7_(pat_let_tv0_)
                    return iife6_(_pat_let1_0)
                def iife4_(_pat_let0_0):
                    def iife9_(d_66_dt__update__tmp_h1_):
                        def iife10_(_pat_let3_0):
                            def iife11_(d_67_dt__update_hcf20_h1_):
                                return D4_DC9(d_67_dt__update_hcf20_h1_)
                            return iife11_(_pat_let3_0)
                        return iife10_(pat_let_tv1_)
                    return iife9_(_pat_let0_0)
                nw11_[int(11)] = iife4_(iife5_(d_59_v36_))
                nw11_[int(12)] = d_59_v36_
                def iife13_(_pat_let5_0):
                    def iife14_(d_68_dt__update__tmp_h2_):
                        def iife15_(_pat_let6_0):
                            def iife16_(d_69_dt__update_hcf20_h2_):
                                return D4_DC9(d_69_dt__update_hcf20_h2_)
                            return iife16_(_pat_let6_0)
                        return iife15_(pat_let_tv2_)
                    return iife14_(_pat_let5_0)
                def iife12_(_pat_let4_0):
                    def iife17_(d_70_dt__update__tmp_h3_):
                        def iife18_(_pat_let7_0):
                            def iife19_(d_71_dt__update_hcf20_h3_):
                                return D4_DC9(d_71_dt__update_hcf20_h3_)
                            return iife19_(_pat_let7_0)
                        return iife18_(pat_let_tv3_)
                    return iife17_(_pat_let4_0)
                nw11_[int(13)] = iife12_(iife13_(d_59_v36_))
                def iife20_(_pat_let8_0):
                    def iife21_(d_72_dt__update__tmp_h4_):
                        def iife22_(_pat_let9_0):
                            def iife23_(d_73_dt__update_hcf20_h4_):
                                return D4_DC9(d_73_dt__update_hcf20_h4_)
                            return iife23_(_pat_let9_0)
                        return iife22_(pat_let_tv4_)
                    return iife21_(_pat_let8_0)
                nw11_[int(14)] = (iife20_(d_59_v36_) if p0 else d_59_v36_)
                nw11_[int(15)] = d_59_v36_
                nw11_[int(16)] = d_59_v36_
                nw11_[int(17)] = D4_DC9(d_61_v38_)
                nw11_[int(18)] = D4_DC9(d_61_v38_)
                nw11_[int(19)] = default__.fm13((d_62_v39_).cardinality, d_60_v37_, d_57_v34_, globalState)
                d_63_v40_ = nw11_
                index5_ = default__.safeIndex(103, (d_63_v40_).length(0))
                (d_63_v40_)[index5_] = d_59_v36_
                d_74_v41_: _dafny.Map
                d_74_v41_ = _dafny.Map({p3: default__.fm1(p0, len(d_60_v37_), globalState)})
                d_75_v42_: C1
                nw12_ = C1()
                nw12_.ctor__()
                d_75_v42_ = nw12_
                d_76_v43_: _dafny.Map
                d_76_v43_ = _dafny.Map({p2: d_75_v42_})
                d_77_v44_: _dafny.Map
                d_77_v44_ = _dafny.Map({len(d_76_v43_): d_57_v34_})
                pat_let_tv5_ = d_61_v38_
                index6_ = default__.safeIndex(103, (d_63_v40_).length(0))
                def iife24_(_pat_let10_0):
                    def iife25_(d_78_dt__update__tmp_h5_):
                        def iife26_(_pat_let11_0):
                            def iife27_(d_79_dt__update_hcf20_h5_):
                                return D4_DC9(d_79_dt__update_hcf20_h5_)
                            return iife27_(_pat_let11_0)
                        return iife26_(pat_let_tv5_)
                    return iife25_(_pat_let10_0)
                rhs0_ = default__.fm3(p1, default__.safeModuloInt(p1, p3), (len(d_74_v41_)) - (p1), (d_77_v44_ if p0 else d_77_v44_), globalState)
                rhs1_ = (iife24_(d_59_v36_) if p0 else d_59_v36_)
                lhs0_ = d_63_v40_
                lhs1_ = default__.safeIndex(103, (d_63_v40_).length(0))
                d_58_v35_ = rhs0_
                lhs0_[lhs1_] = rhs1_
            elif True:
                d_80_v45_: bool
                d_80_v45_ = True
                d_81_v46_: _dafny.Array
                nw13_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_81_v46_ = nw13_
                d_82_v47_: _dafny.Array
                def lambda0_(d_83_p3_):
                    def lambda1_(d_84_i0_):
                        return (d_84_i0_) - ((0) - (d_83_p3_))

                    return lambda1_

                init0_ = lambda0_(p3)
                nw14_ = _dafny.Array(None, 8)
                for i0_0_ in range(nw14_.length(0)):
                    nw14_[i0_0_] = init0_(i0_0_)
                d_82_v47_ = nw14_
                index7_ = default__.safeIndex(803, (d_81_v46_).length(0))
                (d_81_v46_)[index7_] = d_82_v47_
                d_85_v48_: D3
                d_85_v48_ = D3_DC8(p1, p3, p0)
                d_86_v49_: C0
                nw15_ = C0()
                nw15_.ctor__(d_85_v48_)
                d_86_v49_ = nw15_
                d_87_v50_: _dafny.Seq
                d_87_v50_ = _dafny.SeqWithoutIsStrInference([d_86_v49_])
                d_88_v51_: _dafny.Seq
                d_88_v51_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2])
                d_89_v52_: _dafny.Set
                d_89_v52_ = _dafny.Set({p1, len(d_88_v51_)})
                d_90_v53_: _dafny.Seq
                d_90_v53_ = _dafny.SeqWithoutIsStrInference([d_89_v52_])
                d_91_v54_: D3
                d_91_v54_ = D3_DC8(p3, len((d_87_v50_).set(default__.safeIndex(len(d_90_v53_), len(d_87_v50_)), d_86_v49_)), p0)
                d_92_v55_: _dafny.Array
                nw16_ = _dafny.Array(None, 12)
                nw16_[int(0)] = d_80_v45_
                nw16_[int(1)] = p2
                nw16_[int(2)] = d_80_v45_
                nw16_[int(3)] = p2
                nw16_[int(4)] = d_80_v45_
                nw16_[int(5)] = p0
                nw16_[int(6)] = d_80_v45_
                nw16_[int(7)] = d_80_v45_
                nw16_[int(8)] = (d_85_v48_).cf19
                nw16_[int(9)] = not(p2)
                nw16_[int(10)] = p0
                nw16_[int(11)] = p0
                d_92_v55_ = nw16_
                d_93_v56_: _dafny.Seq
                d_93_v56_ = _dafny.SeqWithoutIsStrInference([d_92_v55_])
                d_94_v57_: _dafny.Map
                d_94_v57_ = _dafny.Map({d_92_v55_: d_92_v55_})
                d_95_v58_: C1
                nw17_ = C1()
                nw17_.ctor__()
                d_95_v58_ = nw17_
                d_96_v59_: _dafny.Map
                d_96_v59_ = _dafny.Map({d_95_v58_: d_92_v55_})
                d_97_v60_: _dafny.Array
                nw18_ = _dafny.Array(None, 12)
                nw18_[int(0)] = (d_93_v56_)[default__.safeIndex(p3, len(d_93_v56_))]
                nw18_[int(1)] = d_92_v55_
                nw18_[int(2)] = d_92_v55_
                nw18_[int(3)] = d_92_v55_
                nw18_[int(4)] = d_92_v55_
                nw18_[int(5)] = ((d_94_v57_)[d_92_v55_] if (d_92_v55_) in (d_94_v57_) else d_92_v55_)
                nw18_[int(6)] = d_92_v55_
                nw18_[int(7)] = d_92_v55_
                nw18_[int(8)] = d_92_v55_
                nw18_[int(9)] = d_92_v55_
                nw18_[int(10)] = (d_92_v55_ if True else d_92_v55_)
                nw18_[int(11)] = ((d_96_v59_)[d_95_v58_] if (d_95_v58_) in (d_96_v59_) else d_92_v55_)
                d_97_v60_ = nw18_
                index8_ = default__.safeIndex(247, (d_97_v60_).length(0))
                (d_97_v60_)[index8_] = d_92_v55_
                index9_ = default__.safeIndex(803, (d_81_v46_).length(0))
                index10_ = default__.safeIndex(247, (d_97_v60_).length(0))
                rhs2_ = d_80_v45_
                rhs3_ = d_82_v47_
                rhs4_ = d_92_v55_
                rhs5_ = p2
                lhs2_ = d_81_v46_
                lhs3_ = default__.safeIndex(803, (d_81_v46_).length(0))
                lhs4_ = d_97_v60_
                lhs5_ = default__.safeIndex(247, (d_97_v60_).length(0))
                d_80_v45_ = rhs2_
                lhs2_[lhs3_] = rhs3_
                lhs4_[lhs5_] = rhs4_
                d_80_v45_ = rhs5_
                d_98_v61_: _dafny.Array
                def lambda2_(d_99_v52_):
                    def lambda3_(d_100_i1_):
                        return d_99_v52_

                    return lambda3_

                init1_ = lambda2_(d_89_v52_)
                nw19_ = _dafny.Array(None, 26)
                for i0_1_ in range(nw19_.length(0)):
                    nw19_[i0_1_] = init1_(i0_1_)
                d_98_v61_ = nw19_
                index11_ = default__.safeIndex(590, (d_98_v61_).length(0))
                (d_98_v61_)[index11_] = d_89_v52_
                d_101_v62_: str
                d_101_v62_ = _dafny.CodePoint('e')
                index12_ = default__.safeIndex(590, (d_98_v61_).length(0))
                (d_98_v61_)[index12_] = ((default__.fm14(d_101_v62_, (0) - (p1), globalState)).intersection(d_89_v52_)) | (d_89_v52_)
                d_102_v63_: _dafny.MultiSet
                d_102_v63_ = _dafny.MultiSet([p3])
                d_103_v64_: _dafny.MultiSet
                d_103_v64_ = _dafny.MultiSet([default__.fm15((0) - ((d_102_v63_).cardinality), p1, globalState), D0_DC0(d_88_v51_)])
                d_104_v65_: _dafny.Map
                d_104_v65_ = _dafny.Map({not (d_80_v45_) or (p0): d_103_v64_})
                d_104_v65_ = (d_104_v65_).set(not(p0), d_103_v64_)
                d_80_v45_ = (d_88_v51_)[default__.safeIndex(p3, len(d_88_v51_))]
                d_105_v66_: _dafny.Seq
                d_105_v66_ = _dafny.SeqWithoutIsStrInference([d_52_v29_])
                d_105_v66_ = (d_105_v66_ if p0 else _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0])), d_52_v29_, d_52_v29_, default__.fm16(p3, p1, globalState), d_52_v29_]))
            d_106_v67_: _dafny.Seq
            d_106_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iwvnrgqkx"))
            d_107_v68_: _dafny.Map
            d_107_v68_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgktoekf"))): True})
            (globalState).f0 = (0) - (default__.safeModuloInt(len(d_106_v67_), default__.safeModuloInt(489, len(d_107_v68_))))
            d_108_v69_: _dafny.Map
            d_108_v69_ = _dafny.Map({p1: d_52_v29_})
            d_109_v70_: _dafny.Seq
            d_109_v70_ = _dafny.SeqWithoutIsStrInference([p0])
            d_108_v69_ = (d_108_v69_).set(p3, _dafny.MultiSet(d_109_v70_))
            d_110_v71_: bool
            d_110_v71_ = True
            d_110_v71_ = (default__.safeModuloInt(p1, p3)) >= (p3)
        (globalState).f0 = p1
        d_111_v72_: _dafny.Seq
        d_111_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oncu"))
        d_112_v73_: _dafny.Seq
        d_112_v73_ = _dafny.SeqWithoutIsStrInference([3, 236])
        hi0_ = (0) - ((d_112_v73_)[default__.safeIndex(p1, len(d_112_v73_))])
        for d_113_i2_ in range((0) - (len(d_111_v72_)), hi0_):
            d_114_v74_: _dafny.Map
            d_114_v74_ = _dafny.Map({p3: d_111_v72_})
            d_115_v75_: str
            d_115_v75_ = _dafny.CodePoint('i')
            d_116_v76_: _dafny.Map
            d_116_v76_ = _dafny.Map({d_114_v74_: d_115_v75_})
            d_116_v76_ = (d_116_v76_).set(d_114_v74_, d_115_v75_)
            d_111_v72_ = d_111_v72_
            (globalState).f0 = p3
            d_117_v77_: _dafny.Array
            def lambda4_(d_118_p1_):
                def lambda5_(d_119_i3_):
                    return (d_119_i3_) - (d_118_p1_)

                return lambda5_

            init2_ = lambda4_(p1)
            nw20_ = _dafny.Array(None, 7)
            for i0_2_ in range(nw20_.length(0)):
                nw20_[i0_2_] = init2_(i0_2_)
            d_117_v77_ = nw20_
            index13_ = default__.safeIndex(883, (d_117_v77_).length(0))
            (d_117_v77_)[index13_] = d_113_i2_
            d_120_v78_: _dafny.MultiSet
            d_120_v78_ = _dafny.MultiSet([p3, p3])
            d_121_v79_: _dafny.Seq
            d_121_v79_ = _dafny.SeqWithoutIsStrInference([d_120_v78_])
            d_122_v80_: C1
            nw21_ = C1()
            nw21_.ctor__()
            d_122_v80_ = nw21_
            d_123_v81_: _dafny.Map
            d_123_v81_ = _dafny.Map({d_122_v80_: d_122_v80_})
            index14_ = default__.safeIndex(883, (d_117_v77_).length(0))
            (d_117_v77_)[index14_] = (0) - ((p3) - ((len(d_121_v79_) if p0 else (0) - (len(d_123_v81_)))))
        (globalState).f0 = p1
        d_124_v82_: _dafny.Array
        nw22_ = _dafny.Array(_dafny.Seq({}), 28)
        d_124_v82_ = nw22_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_124_v82_).length(0)):
            d_125_i4_: int = guard_loop_0_
            if (True) and (((0) <= (d_125_i4_)) and ((d_125_i4_) < ((d_124_v82_).length(0)))):
                (d_124_v82_)[(d_125_i4_)] = ((d_112_v73_) + (d_112_v73_)) + (d_112_v73_)
        d_126_v83_: _dafny.Array
        nw23_ = _dafny.Array(None, 27)
        nw23_[int(0)] = False
        nw23_[int(1)] = p0
        nw23_[int(2)] = p2
        nw23_[int(3)] = not(p0)
        nw23_[int(4)] = p2
        nw23_[int(5)] = p0
        nw23_[int(6)] = p0
        nw23_[int(7)] = p2
        nw23_[int(8)] = p2
        nw23_[int(9)] = p0
        nw23_[int(10)] = p0
        nw23_[int(11)] = p0
        nw23_[int(12)] = p0
        nw23_[int(13)] = p2
        nw23_[int(14)] = p2
        nw23_[int(15)] = p0
        nw23_[int(16)] = p2
        nw23_[int(17)] = default__.fm2(globalState)
        nw23_[int(18)] = p0
        nw23_[int(19)] = p2
        nw23_[int(20)] = p0
        nw23_[int(21)] = p0
        nw23_[int(22)] = p2
        nw23_[int(23)] = p0
        nw23_[int(24)] = not(p0)
        nw23_[int(25)] = p0
        nw23_[int(26)] = p0
        d_126_v83_ = nw23_
        d_127_v84_: D4
        d_127_v84_ = D4_DC12(p2, default__.fm17(p3, p2, False, globalState), p1, d_126_v83_, p2)
        d_128_v85_: _dafny.Map
        d_128_v85_ = _dafny.Map({p3: (0) - (default__.fm0(globalState))})
        d_129_v86_: _dafny.Seq
        d_129_v86_ = _dafny.SeqWithoutIsStrInference([d_128_v85_])
        pat_let_tv6_ = d_126_v83_
        pat_let_tv7_ = d_112_v73_
        pat_let_tv8_ = d_129_v86_
        d_130_v87_: _dafny.Array
        nw24_ = _dafny.Array(None, 1)
        def iife28_(_pat_let12_0):
            def iife29_(d_131_dt__update__tmp_h6_):
                def iife30_(_pat_let13_0):
                    def iife31_(d_132_dt__update_hcf29_h0_):
                        def iife32_(_pat_let14_0):
                            def iife33_(d_133_dt__update_hcf28_h0_):
                                def iife34_(_pat_let15_0):
                                    def iife35_(d_134_dt__update_hcf27_h0_):
                                        return D4_DC12((d_131_dt__update__tmp_h6_).cf26, d_134_dt__update_hcf27_h0_, d_133_dt__update_hcf28_h0_, d_132_dt__update_hcf29_h0_, (d_131_dt__update__tmp_h6_).cf30)
                                    return iife35_(_pat_let15_0)
                                return iife34_(pat_let_tv8_)
                            return iife33_(_pat_let14_0)
                        return iife32_(len(pat_let_tv7_))
                    return iife31_(_pat_let13_0)
                return iife30_(pat_let_tv6_)
            return iife29_(_pat_let12_0)
        nw24_[int(0)] = iife28_(d_127_v84_)
        d_130_v87_ = nw24_
        index15_ = default__.safeIndex(156, (d_130_v87_).length(0))
        (d_130_v87_)[index15_] = d_127_v84_
        index16_ = default__.safeIndex(156, (d_130_v87_).length(0))
        (d_130_v87_)[index16_] = d_127_v84_

    @staticmethod
    def Main(noArgsParameter__):
        d_135_v0_: bool
        d_135_v0_ = True
        d_136_v1_: _dafny.Set
        d_136_v1_ = _dafny.Set({d_135_v0_, d_135_v0_, d_135_v0_})
        d_137_v2_: _dafny.Array
        def lambda6_(d_138_v0_):
            def lambda7_(d_139_i0_):
                return (D0_DC0(_dafny.SeqWithoutIsStrInference([d_138_v0_, d_138_v0_]))).cf0

            return lambda7_

        init3_ = lambda6_(d_135_v0_)
        nw25_ = _dafny.Array(None, 1)
        for i0_3_ in range(nw25_.length(0)):
            nw25_[i0_3_] = init3_(i0_3_)
        d_137_v2_ = nw25_
        d_140_v3_: int
        d_140_v3_ = -730
        d_141_v4_: _dafny.Map
        d_141_v4_ = _dafny.Map({d_135_v0_: d_140_v3_})
        d_142_v5_: _dafny.Array
        nw26_ = _dafny.Array(None, 21)
        nw26_[int(0)] = d_141_v4_
        nw26_[int(1)] = _dafny.Map({d_135_v0_: d_140_v3_})
        nw26_[int(2)] = d_141_v4_
        nw26_[int(3)] = d_141_v4_
        nw26_[int(4)] = d_141_v4_
        nw26_[int(5)] = _dafny.Map({d_135_v0_: d_140_v3_})
        nw26_[int(6)] = d_141_v4_
        nw26_[int(7)] = d_141_v4_
        nw26_[int(8)] = d_141_v4_
        nw26_[int(9)] = d_141_v4_
        nw26_[int(10)] = _dafny.Map({d_135_v0_: d_140_v3_})
        nw26_[int(11)] = d_141_v4_
        nw26_[int(12)] = d_141_v4_
        nw26_[int(13)] = d_141_v4_
        nw26_[int(14)] = d_141_v4_
        nw26_[int(15)] = d_141_v4_
        nw26_[int(16)] = d_141_v4_
        nw26_[int(17)] = d_141_v4_
        nw26_[int(18)] = d_141_v4_
        nw26_[int(19)] = d_141_v4_
        nw26_[int(20)] = _dafny.Map({d_135_v0_: d_140_v3_})
        d_142_v5_ = nw26_
        d_143_globalState_: GlobalState
        nw27_ = GlobalState()
        nw27_.ctor__(317, (_dafny.Set({d_135_v0_})) - (d_136_v1_), d_137_v2_, d_142_v5_, _dafny.CodePoint('u'), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tinqkjdyt")), 892, 437, True, _dafny.CodePoint('o'), False, -34, False, False)
        d_143_globalState_ = nw27_
        d_144_v6_: _dafny.Seq
        d_144_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whhg"))
        hi1_ = len(d_144_v6_)
        for d_145_i1_ in range(len(d_136_v1_), hi1_):
            d_146_v7_: _dafny.Array
            nw28_ = _dafny.Array(int(0), 15)
            d_146_v7_ = nw28_
            d_147_v8_: _dafny.Seq
            d_147_v8_ = _dafny.SeqWithoutIsStrInference([d_146_v7_])
            d_148_v9_: _dafny.Map
            d_148_v9_ = _dafny.Map({(d_140_v3_) + (len(d_144_v6_)): d_147_v8_})
            d_148_v9_ = (d_148_v9_).set((0) - (default__.fm0(d_143_globalState_)), d_147_v8_)
            d_140_v3_ = (0) - (d_140_v3_)
            d_149_v10_: _dafny.Seq
            d_149_v10_ = _dafny.SeqWithoutIsStrInference([d_140_v3_, d_145_i1_])
            if (default__.safeDivisionInt(d_140_v3_, ((_dafny.MultiSet(d_149_v10_)).set(d_140_v3_, default__.abs(d_145_i1_))).cardinality)) > (d_140_v3_):
                d_150_v11_: _dafny.Seq
                d_150_v11_ = _dafny.SeqWithoutIsStrInference([d_135_v0_])
                d_151_v12_: D0
                d_151_v12_ = D0_DC2(d_135_v0_, d_135_v0_, (0) - (len(d_150_v11_)))
                d_140_v3_ = (d_151_v12_).cf4
                d_152_v13_: _dafny.Array
                nw29_ = _dafny.Array(False, 24)
                d_152_v13_ = nw29_
                d_153_v14_: _dafny.Seq
                d_153_v14_ = _dafny.SeqWithoutIsStrInference([d_152_v13_])
                index17_ = default__.safeIndex(49, (d_152_v13_).length(0))
                (d_152_v13_)[index17_] = (d_153_v14_) != (d_153_v14_)
                index18_ = default__.safeIndex(49, (d_152_v13_).length(0))
                (d_152_v13_)[index18_] = (d_135_v0_) and (d_135_v0_)
                d_154_v15_: _dafny.Map
                d_154_v15_ = _dafny.Map({len(d_144_v6_): d_135_v0_})
                default__.m0((d_140_v3_) < (d_145_i1_), len(d_154_v15_), (d_152_v13_)[default__.safeIndex(49, (d_152_v13_).length(0))], d_140_v3_, d_143_globalState_)
                d_155_v16_: _dafny.MultiSet
                d_155_v16_ = _dafny.MultiSet([d_145_i1_])
                d_156_v17_: _dafny.Map
                d_156_v17_ = _dafny.Map({d_145_i1_: default__.safeModuloInt(((d_155_v16_)[d_140_v3_] if (d_140_v3_) in (d_155_v16_) else d_145_i1_), d_145_i1_)})
                d_157_v18_: str
                d_157_v18_ = _dafny.CodePoint('i')
                d_156_v17_ = (d_156_v17_).set(len(d_144_v6_), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_158_i2_ in range(default__.abs(357))])).set(default__.safeIndex(d_140_v3_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_159_i2_ in range(default__.abs(357))]))), d_157_v18_)))
                (d_143_globalState_).f0 = default__.safeModuloInt(d_145_i1_, d_145_i1_)
            elif True:
                d_140_v3_ = d_140_v3_
                default__.m0(not(d_135_v0_), d_140_v3_, (d_140_v3_) > (d_140_v3_), d_145_i1_, d_143_globalState_)
                default__.m0(d_135_v0_, d_145_i1_, False, (d_140_v3_) * (d_140_v3_), d_143_globalState_)
                d_160_v19_: _dafny.Map
                d_160_v19_ = _dafny.Map({d_140_v3_: d_144_v6_})
                d_161_v20_: _dafny.Map
                d_161_v20_ = _dafny.Map({d_140_v3_: (d_144_v6_) < (((d_160_v19_)[d_145_i1_] if (d_145_i1_) in (d_160_v19_) else d_144_v6_))})
                d_162_v21_: _dafny.Map
                d_162_v21_ = _dafny.Map({d_135_v0_: d_135_v0_})
                d_135_v0_ = ((d_161_v20_)[d_145_i1_] if (d_145_i1_) in (d_161_v20_) else ((d_162_v21_)[d_135_v0_] if (d_135_v0_) in (d_162_v21_) else d_135_v0_))
                d_161_v20_ = d_161_v20_
            d_163_v22_: D0
            d_163_v22_ = D0_DC2(d_135_v0_, d_135_v0_, d_145_i1_)
            d_164_v23_: _dafny.Map
            d_164_v23_ = _dafny.Map({d_135_v0_: (d_163_v22_).cf2})
            default__.m0(((d_164_v23_)[d_135_v0_] if (d_135_v0_) in (d_164_v23_) else d_135_v0_), d_140_v3_, d_135_v0_, (d_140_v3_) + ((0) - (d_140_v3_)), d_143_globalState_)
        default__.m0((d_140_v3_) <= (d_140_v3_), (d_140_v3_) - (d_140_v3_), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bffs"))) <= (d_144_v6_), d_140_v3_, d_143_globalState_)
        hi2_ = (0) - (d_140_v3_)
        for d_165_i3_ in range(d_140_v3_, hi2_):
            d_135_v0_ = d_135_v0_
            default__.m0(d_135_v0_, (d_165_i3_) - (default__.fm0(d_143_globalState_)), False, d_140_v3_, d_143_globalState_)
            d_166_v24_: _dafny.Seq
            d_166_v24_ = _dafny.SeqWithoutIsStrInference([d_135_v0_, d_135_v0_])
            d_167_v25_: _dafny.Map
            d_167_v25_ = _dafny.Map({d_165_i3_: d_140_v3_})
            d_168_v26_: _dafny.Seq
            d_168_v26_ = _dafny.SeqWithoutIsStrInference([d_167_v25_])
            d_169_v27_: str
            d_169_v27_ = _dafny.CodePoint('g')
            d_170_v28_: _dafny.Seq
            d_170_v28_ = _dafny.SeqWithoutIsStrInference([len((default__.fm1(d_135_v0_, d_140_v3_, d_143_globalState_)).set(default__.safeIndex(len(d_168_v26_), len(default__.fm1(d_135_v0_, d_140_v3_, d_143_globalState_))), d_169_v27_))])
            d_171_v29_: _dafny.Map
            d_171_v29_ = _dafny.Map({d_165_i3_: not(d_135_v0_)})
            default__.m0((d_135_v0_) in (d_166_v24_), (d_170_v28_)[default__.safeIndex(d_140_v3_, len(d_170_v28_))], ((d_171_v29_)[d_165_i3_] if (d_165_i3_) in (d_171_v29_) else d_135_v0_), len((d_166_v24_).set(default__.safeIndex(d_165_i3_, len(d_166_v24_)), d_135_v0_)), d_143_globalState_)
            d_172_v30_: _dafny.Map
            d_172_v30_ = _dafny.Map({d_135_v0_: default__.fm2(d_143_globalState_)})
            d_140_v3_ = (len(d_172_v30_)) * (d_165_i3_)
        d_173_v31_: D0
        d_173_v31_ = D0_DC1((d_140_v3_) + (d_140_v3_))
        source1_ = d_173_v31_
        if source1_.is_DC1:
            d_174___mcc_h0_ = source1_.cf1
            d_175_cf1_ = d_174___mcc_h0_
            d_135_v0_ = default__.fm2(d_143_globalState_)
            d_176_v32_: _dafny.Seq
            d_176_v32_ = _dafny.SeqWithoutIsStrInference([False])
            d_176_v32_ = d_176_v32_
            d_177_v33_: _dafny.Map
            d_177_v33_ = _dafny.Map({d_175_cf1_: d_135_v0_})
            d_178_v34_: _dafny.Seq
            d_178_v34_ = _dafny.SeqWithoutIsStrInference([d_140_v3_])
            d_176_v32_ = (d_176_v32_) + ((d_176_v32_) + (default__.fm3((0) - (d_175_cf1_), -816, d_175_cf1_, (d_177_v33_).set(len(d_178_v34_), d_135_v0_), d_143_globalState_)))
            d_179_v35_: D0
            d_179_v35_ = D0_DC2(d_135_v0_, d_135_v0_, d_140_v3_)
            d_180_v36_: _dafny.Array
            nw30_ = _dafny.Array(None, 12)
            nw30_[int(0)] = d_135_v0_
            nw30_[int(1)] = not(d_135_v0_)
            nw30_[int(2)] = (d_179_v35_).cf2
            nw30_[int(3)] = d_135_v0_
            nw30_[int(4)] = d_135_v0_
            nw30_[int(5)] = d_135_v0_
            nw30_[int(6)] = d_135_v0_
            nw30_[int(7)] = d_135_v0_
            nw30_[int(8)] = False
            nw30_[int(9)] = d_135_v0_
            nw30_[int(10)] = not(not((d_179_v35_).cf2))
            nw30_[int(11)] = d_135_v0_
            d_180_v36_ = nw30_
            d_181_v37_: _dafny.MultiSet
            d_181_v37_ = _dafny.MultiSet([d_180_v36_, d_180_v36_, d_180_v36_])
            d_135_v0_ = not (not((d_180_v36_) in (d_181_v37_))) or (d_135_v0_)
        elif source1_.is_DC2:
            d_182___mcc_h1_ = source1_.cf2
            d_183___mcc_h2_ = source1_.cf3
            d_184___mcc_h3_ = source1_.cf4
            d_185_cf4_ = d_184___mcc_h3_
            d_186_cf3_ = d_183___mcc_h2_
            d_187_cf2_ = d_182___mcc_h1_
            d_188_v38_: _dafny.Seq
            d_188_v38_ = _dafny.SeqWithoutIsStrInference([d_135_v0_, True, d_135_v0_])
            default__.m0((d_188_v38_) <= (d_188_v38_), d_140_v3_, (d_186_cf3_) in (_dafny.Map({d_186_cf3_: (0) - (len(d_144_v6_))})), d_140_v3_, d_143_globalState_)
            d_189_v39_: C1
            nw31_ = C1()
            nw31_.ctor__()
            d_189_v39_ = nw31_
            d_190_v40_: _dafny.Array
            def lambda8_(d_191_v0_, d_192_cf3_, d_193_cf2_):
                def lambda9_(d_194_i4_):
                    return (_dafny.Map({d_191_v0_: d_192_cf3_})) | (_dafny.Map({(D0_DC2(d_191_v0_, d_193_cf2_, (_dafny.MultiSet([(_dafny.MultiSet([_dafny.CodePoint('y')])).cardinality])).cardinality)).cf3: d_192_cf3_}))

                return lambda9_

            init4_ = lambda8_(d_135_v0_, d_186_cf3_, d_187_cf2_)
            nw32_ = _dafny.Array(None, 5)
            for i0_4_ in range(nw32_.length(0)):
                nw32_[i0_4_] = init4_(i0_4_)
            d_190_v40_ = nw32_
            index19_ = default__.safeIndex(78, (d_190_v40_).length(0))
            (d_190_v40_)[index19_] = default__.fm7(True, d_143_globalState_)
            d_195_v41_: _dafny.Map
            d_195_v41_ = _dafny.Map({d_186_cf3_: d_135_v0_})
            index20_ = default__.safeIndex(78, (d_190_v40_).length(0))
            (d_190_v40_)[index20_] = (d_195_v41_) | (default__.fm7(d_135_v0_, d_143_globalState_))
            d_196_v42_: _dafny.Array
            def lambda10_(d_197_cf3_):
                def lambda11_(d_198_i5_):
                    return d_197_cf3_

                return lambda11_

            init5_ = lambda10_(d_186_cf3_)
            nw33_ = _dafny.Array(None, 6)
            for i0_5_ in range(nw33_.length(0)):
                nw33_[i0_5_] = init5_(i0_5_)
            d_196_v42_ = nw33_
            index21_ = default__.safeIndex(193, (d_196_v42_).length(0))
            (d_196_v42_)[index21_] = d_186_cf3_
            index22_ = default__.safeIndex(903, (d_196_v42_).length(0))
            (d_196_v42_)[index22_] = (d_140_v3_) > (d_185_cf4_)
            index23_ = default__.safeIndex(417, (d_196_v42_).length(0))
            (d_196_v42_)[index23_] = (d_185_cf4_) < (d_140_v3_)
            index24_ = default__.safeIndex(193, (d_196_v42_).length(0))
            index25_ = default__.safeIndex(903, (d_196_v42_).length(0))
            index26_ = default__.safeIndex(417, (d_196_v42_).length(0))
            rhs6_ = (d_185_cf4_) <= (d_185_cf4_)
            rhs7_ = d_186_cf3_
            rhs8_ = False
            rhs9_ = d_186_cf3_
            lhs6_ = d_196_v42_
            lhs7_ = default__.safeIndex(193, (d_196_v42_).length(0))
            lhs8_ = d_196_v42_
            lhs9_ = default__.safeIndex(903, (d_196_v42_).length(0))
            lhs10_ = d_196_v42_
            lhs11_ = default__.safeIndex(417, (d_196_v42_).length(0))
            lhs6_[lhs7_] = rhs6_
            d_187_cf2_ = rhs7_
            lhs8_[lhs9_] = rhs8_
            lhs10_[lhs11_] = rhs9_
        elif True:
            d_199___mcc_h4_ = source1_.cf0
            d_200_cf0_ = d_199___mcc_h4_
            default__.m0(d_135_v0_, d_140_v3_, d_135_v0_, (d_140_v3_) + (d_140_v3_), d_143_globalState_)
            d_201_v43_: _dafny.MultiSet
            d_201_v43_ = _dafny.MultiSet([d_141_v4_])
            default__.m0(d_135_v0_, ((d_201_v43_).cardinality) - (d_140_v3_), d_135_v0_, len(d_144_v6_), d_143_globalState_)
            d_202_v44_: _dafny.Array
            nw34_ = _dafny.Array(None, 15)
            nw34_[int(0)] = d_135_v0_
            nw34_[int(1)] = d_135_v0_
            nw34_[int(2)] = d_135_v0_
            nw34_[int(3)] = d_135_v0_
            nw34_[int(4)] = d_135_v0_
            nw34_[int(5)] = d_135_v0_
            nw34_[int(6)] = d_135_v0_
            nw34_[int(7)] = default__.fm2(d_143_globalState_)
            nw34_[int(8)] = d_135_v0_
            nw34_[int(9)] = d_135_v0_
            nw34_[int(10)] = d_135_v0_
            nw34_[int(11)] = d_135_v0_
            nw34_[int(12)] = d_135_v0_
            nw34_[int(13)] = d_135_v0_
            nw34_[int(14)] = d_135_v0_
            d_202_v44_ = nw34_
            d_203_v45_: _dafny.Map
            d_203_v45_ = _dafny.Map({d_202_v44_: d_140_v3_})
            d_204_v46_: _dafny.Map
            d_204_v46_ = _dafny.Map({(d_203_v45_) | (_dafny.Map({d_202_v44_: d_140_v3_})): (d_144_v6_) + (d_144_v6_)})
            d_204_v46_ = (d_204_v46_).set((d_203_v45_) | (d_203_v45_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jbs")))
            d_205_v47_: str
            d_205_v47_ = _dafny.CodePoint('j')
            (d_143_globalState_).f9 = d_205_v47_
        d_135_v0_ = d_135_v0_
        d_206_v48_: D1
        d_206_v48_ = D1_DC4(423, d_135_v0_, d_135_v0_, d_135_v0_)
        pat_let_tv9_ = d_140_v3_
        pat_let_tv10_ = d_141_v4_
        def lambda12_(source2_):
            if source2_.is_DC4:
                d_207___mcc_h5_ = source2_.cf6
                d_208___mcc_h6_ = source2_.cf7
                d_209___mcc_h7_ = source2_.cf8
                d_210___mcc_h8_ = source2_.cf9
                d_211_cf9_ = d_210___mcc_h8_
                d_212_cf8_ = d_209___mcc_h7_
                d_213_cf7_ = d_208___mcc_h6_
                d_214_cf6_ = d_207___mcc_h5_
                return not((_dafny.Set({pat_let_tv9_, len(pat_let_tv10_)})).issubset(_dafny.Set({d_214_cf6_})))
            elif True:
                d_215___mcc_h9_ = source2_.cf5
                d_216_cf5_ = d_215___mcc_h9_
                return False

        if lambda12_(d_206_v48_):
            d_217_v49_: _dafny.Seq
            d_217_v49_ = _dafny.SeqWithoutIsStrInference([d_140_v3_, ((d_141_v4_)[d_135_v0_] if (d_135_v0_) in (d_141_v4_) else d_140_v3_), d_140_v3_, 390, d_140_v3_])
            rhs10_ = default__.safeDivisionInt(d_140_v3_, len(d_217_v49_))
            rhs11_ = d_140_v3_
            rhs12_ = d_140_v3_
            lhs12_ = d_143_globalState_
            lhs13_ = d_143_globalState_
            lhs12_.f0 = rhs10_
            lhs13_.f0 = rhs11_
            d_140_v3_ = rhs12_
            d_218_v50_: _dafny.Array
            nw35_ = _dafny.Array(int(0), 12)
            d_218_v50_ = nw35_
            d_219_v51_: _dafny.Seq
            d_219_v51_ = _dafny.SeqWithoutIsStrInference([d_218_v50_, d_218_v50_, d_218_v50_])
            d_219_v51_ = d_219_v51_
            index27_ = default__.safeIndex(338, (d_218_v50_).length(0))
            (d_218_v50_)[index27_] = d_140_v3_
            d_220_v52_: _dafny.MultiSet
            d_220_v52_ = _dafny.MultiSet([d_135_v0_])
            d_221_v53_: _dafny.Map
            d_221_v53_ = _dafny.Map({default__.fm0(d_143_globalState_): d_140_v3_})
            index28_ = default__.safeIndex(338, (d_218_v50_).length(0))
            (d_218_v50_)[index28_] = ((d_220_v52_)[(d_135_v0_ if d_135_v0_ else d_135_v0_)] if ((d_135_v0_ if d_135_v0_ else d_135_v0_)) in (d_220_v52_) else len(d_221_v53_))
            if default__.fm2(d_143_globalState_):
                d_222_v54_: _dafny.Map
                d_222_v54_ = _dafny.Map({d_140_v3_: d_135_v0_})
                d_222_v54_ = (default__.fm8(d_140_v3_, d_143_globalState_)) | (default__.fm8(473, d_143_globalState_))
                d_223_v55_: _dafny.Array
                def lambda13_(d_224_i6_):
                    return _dafny.CodePoint('w')

                init6_ = lambda13_
                nw36_ = _dafny.Array(None, 9)
                for i0_6_ in range(nw36_.length(0)):
                    nw36_[i0_6_] = init6_(i0_6_)
                d_223_v55_ = nw36_
                d_225_v56_: str
                d_225_v56_ = _dafny.CodePoint('k')
                index29_ = default__.safeIndex(726, (d_223_v55_).length(0))
                (d_223_v55_)[index29_] = d_225_v56_
                index30_ = default__.safeIndex(726, (d_223_v55_).length(0))
                (d_223_v55_)[index30_] = _dafny.CodePoint('g')
                (d_143_globalState_).f0 = default__.safeDivisionInt(default__.safeModuloInt((d_218_v50_)[default__.safeIndex(338, (d_218_v50_).length(0))], d_140_v3_), (d_218_v50_)[default__.safeIndex(338, (d_218_v50_).length(0))])
                (d_143_globalState_).f0 = d_140_v3_
                d_226_v57_: _dafny.Array
                nw37_ = _dafny.Array(_dafny.Seq({}), 2)
                d_226_v57_ = nw37_
                index31_ = default__.safeIndex(403, (d_226_v57_).length(0))
                (d_226_v57_)[index31_] = _dafny.SeqWithoutIsStrInference([d_221_v53_, d_221_v53_, d_221_v53_])
                d_227_v58_: _dafny.Seq
                d_227_v58_ = _dafny.SeqWithoutIsStrInference([d_221_v53_, d_221_v53_])
                index32_ = default__.safeIndex(403, (d_226_v57_).length(0))
                (d_226_v57_)[index32_] = (d_227_v58_) + (d_227_v58_)
            elif True:
                d_228_v59_: _dafny.Map
                d_228_v59_ = _dafny.Map({not(not(d_135_v0_)): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlxmbhbd"))})
                d_228_v59_ = (d_228_v59_) | ((d_228_v59_) | (_dafny.Map({d_135_v0_: d_144_v6_})))
                d_229_v60_: _dafny.Set
                d_229_v60_ = _dafny.Set({d_140_v3_})
                d_229_v60_ = d_229_v60_
                d_230_v61_: _dafny.Map
                d_230_v61_ = _dafny.Map({D3_DC8(d_140_v3_, d_140_v3_, d_135_v0_): d_135_v0_})
                d_231_v63_: _dafny.Map
                d_231_v63_ = _dafny.Map({default__.fm9(d_143_globalState_): d_144_v6_})
                d_232_v64_: _dafny.Set
                d_232_v64_ = _dafny.Set({D1_DC3(d_231_v63_)})
                d_233_v65_: D3
                d_233_v65_ = D3_DC8((d_218_v50_)[default__.safeIndex(338, (d_218_v50_).length(0))], len(d_232_v64_), d_135_v0_)
                d_234_v66_: _dafny.Array
                nw38_ = _dafny.Array(None, 2)
                nw38_[int(0)] = d_230_v61_
                def iife36_():
                    coll4_ = _dafny.Map()
                    compr_4_: D3
                    for compr_4_ in (_dafny.SeqWithoutIsStrInference([d_233_v65_])).Elements:
                        d_235_v62_: D3 = compr_4_
                        if (d_235_v62_) in (_dafny.SeqWithoutIsStrInference([d_233_v65_])):
                            coll4_[d_235_v62_] = d_135_v0_
                    return _dafny.Map(coll4_)
                nw38_[int(1)] = iife36_()
                
                d_234_v66_ = nw38_
                index33_ = default__.safeIndex(714, (d_234_v66_).length(0))
                (d_234_v66_)[index33_] = _dafny.Map({d_233_v65_: True})
                index34_ = default__.safeIndex(714, (d_234_v66_).length(0))
                (d_234_v66_)[index34_] = (d_230_v61_) | (_dafny.Map({d_233_v65_: default__.fm2(d_143_globalState_)}))
                d_135_v0_ = True
                d_236_v67_: D0
                d_236_v67_ = D0_DC2(d_135_v0_, d_135_v0_, (d_218_v50_)[default__.safeIndex(338, (d_218_v50_).length(0))])
                pat_let_tv11_ = d_135_v0_
                d_237_v68_: _dafny.Array
                nw39_ = _dafny.Array(None, 11)
                nw39_[int(0)] = d_236_v67_
                nw39_[int(1)] = d_236_v67_
                nw39_[int(2)] = d_236_v67_
                nw39_[int(3)] = d_236_v67_
                nw39_[int(4)] = D0_DC2(d_135_v0_, d_135_v0_, default__.fm0(d_143_globalState_))
                nw39_[int(5)] = d_236_v67_
                nw39_[int(6)] = d_236_v67_
                def iife37_(_pat_let16_0):
                    def iife38_(d_238_dt__update__tmp_h0_):
                        def iife39_(_pat_let17_0):
                            def iife40_(d_239_dt__update_hcf2_h0_):
                                return D0_DC2(d_239_dt__update_hcf2_h0_, (d_238_dt__update__tmp_h0_).cf3, (d_238_dt__update__tmp_h0_).cf4)
                            return iife40_(_pat_let17_0)
                        return iife39_(pat_let_tv11_)
                    return iife38_(_pat_let16_0)
                nw39_[int(7)] = iife37_(d_236_v67_)
                nw39_[int(8)] = d_236_v67_
                nw39_[int(9)] = d_236_v67_
                def iife41_(_pat_let18_0):
                    def iife42_(d_240_dt__update__tmp_h1_):
                        def iife43_(_pat_let19_0):
                            def iife44_(d_241_dt__update_hcf3_h0_):
                                return D0_DC2((d_240_dt__update__tmp_h1_).cf2, d_241_dt__update_hcf3_h0_, (d_240_dt__update__tmp_h1_).cf4)
                            return iife44_(_pat_let19_0)
                        return iife43_(False)
                    return iife42_(_pat_let18_0)
                nw39_[int(10)] = iife41_(default__.fm10(d_135_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oysattfu")), d_143_globalState_))
                d_237_v68_ = nw39_
                index35_ = default__.safeIndex(408, (d_237_v68_).length(0))
                (d_237_v68_)[index35_] = d_236_v67_
                index36_ = default__.safeIndex(408, (d_237_v68_).length(0))
                rhs13_ = (len(d_136_v1_)) + (d_140_v3_)
                rhs14_ = _dafny.CodePoint('e')
                rhs15_ = default__.fm10(d_135_v0_, d_144_v6_, d_143_globalState_)
                rhs16_ = (0) - ((d_218_v50_)[default__.safeIndex(338, (d_218_v50_).length(0))])
                lhs14_ = d_143_globalState_
                lhs15_ = d_143_globalState_
                lhs16_ = d_237_v68_
                lhs17_ = default__.safeIndex(408, (d_237_v68_).length(0))
                lhs14_.f0 = rhs13_
                lhs15_.f4 = rhs14_
                lhs16_[lhs17_] = rhs15_
                d_140_v3_ = rhs16_
            (d_143_globalState_).f0 = d_140_v3_
        elif True:
            (d_143_globalState_).f0 = 660
            d_242_v69_: _dafny.Array
            nw40_ = _dafny.Array(False, 22)
            d_242_v69_ = nw40_
            index37_ = default__.safeIndex(779, (d_242_v69_).length(0))
            (d_242_v69_)[index37_] = not(d_135_v0_)
            index38_ = default__.safeIndex(779, (d_242_v69_).length(0))
            (d_242_v69_)[index38_] = (d_140_v3_) == (d_140_v3_)
            d_243_v70_: _dafny.Seq
            d_243_v70_ = _dafny.SeqWithoutIsStrInference([d_144_v6_])
            index39_ = default__.safeIndex(779, (d_242_v69_).length(0))
            rhs17_ = (d_140_v3_) != (d_140_v3_)
            rhs18_ = not((d_140_v3_) != (default__.safeDivisionInt(default__.fm0(d_143_globalState_), len(d_243_v70_))))
            lhs18_ = d_242_v69_
            lhs19_ = default__.safeIndex(779, (d_242_v69_).length(0))
            d_135_v0_ = rhs17_
            lhs18_[lhs19_] = rhs18_
            d_135_v0_ = (d_140_v3_) <= (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({d_140_v3_: (_dafny.MultiSet([d_135_v0_, (d_242_v69_)[default__.safeIndex(779, (d_242_v69_).length(0))]])).cardinality}) for d_244_i7_ in range(default__.abs(486))])))
            d_245_v71_: _dafny.Seq
            d_245_v71_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_144_v6_: d_242_v69_})])
            (d_143_globalState_).f0 = default__.safeDivisionInt((0) - (d_140_v3_), len((d_245_v71_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_246_i8_ in range(default__.abs(576))])), len(d_245_v71_))]))
        d_247_v72_: _dafny.Map
        d_247_v72_ = _dafny.Map({d_140_v3_: d_135_v0_})
        d_247_v72_ = (d_247_v72_) | (d_247_v72_)
        d_248_v73_: D0
        d_248_v73_ = D0_DC2(d_135_v0_, d_135_v0_, d_140_v3_)
        pat_let_tv12_ = d_135_v0_
        pat_let_tv13_ = d_140_v3_
        def lambda14_(source3_):
            if source3_.is_DC1:
                d_249___mcc_h10_ = source3_.cf1
                d_250_cf1_ = d_249___mcc_h10_
                return d_250_cf1_
            elif source3_.is_DC2:
                d_251___mcc_h11_ = source3_.cf2
                d_252___mcc_h12_ = source3_.cf3
                d_253___mcc_h13_ = source3_.cf4
                d_254_cf4_ = d_253___mcc_h13_
                d_255_cf3_ = d_252___mcc_h12_
                d_256_cf2_ = d_251___mcc_h11_
                return (_dafny.MultiSet([pat_let_tv12_])).cardinality
            elif True:
                d_257___mcc_h14_ = source3_.cf0
                d_258_cf0_ = d_257___mcc_h14_
                return (667) * (pat_let_tv13_)

        d_140_v3_ = lambda14_(d_248_v73_)
        d_135_v0_ = d_135_v0_
        d_259_v74_: _dafny.Map
        d_259_v74_ = _dafny.Map({d_206_v48_: d_144_v6_})
        d_260_v75_: D2
        d_260_v75_ = D2_DC6((d_144_v6_) + (((d_259_v74_)[D1_DC4((0) - (len(d_144_v6_)), d_135_v0_, d_135_v0_, d_135_v0_)] if (D1_DC4((0) - (len(d_144_v6_)), d_135_v0_, d_135_v0_, d_135_v0_)) in (d_259_v74_) else d_144_v6_)), d_140_v3_, d_135_v0_, d_140_v3_, d_135_v0_)
        d_260_v75_ = d_260_v75_
        d_261_v76_: _dafny.Map
        d_261_v76_ = _dafny.Map({d_140_v3_: (D0_DC1(d_140_v3_)).cf1})
        d_262_v77_: _dafny.Seq
        d_262_v77_ = _dafny.SeqWithoutIsStrInference([d_140_v3_])
        d_263_v78_: _dafny.Seq
        d_263_v78_ = _dafny.SeqWithoutIsStrInference([d_261_v76_, _dafny.Map({len(d_144_v6_): len(d_262_v77_)}), d_261_v76_, d_261_v76_, d_261_v76_])
        d_264_v79_: _dafny.Seq
        d_264_v79_ = _dafny.SeqWithoutIsStrInference([d_135_v0_])
        d_265_v80_: D3
        d_265_v80_ = D3_DC8(d_140_v3_, d_140_v3_, False)
        d_266_v81_: C0
        nw41_ = C0()
        nw41_.ctor__(d_265_v80_)
        d_266_v81_ = nw41_
        d_267_v82_: _dafny.Seq
        d_267_v82_ = _dafny.SeqWithoutIsStrInference([d_266_v81_, d_266_v81_])
        d_268_v83_: _dafny.MultiSet
        d_268_v83_ = _dafny.MultiSet([(d_267_v82_)[default__.safeIndex(d_140_v3_, len(d_267_v82_))]])
        d_269_v84_: _dafny.MultiSet
        d_269_v84_ = _dafny.MultiSet([d_140_v3_, (d_268_v83_).cardinality, (0) - (d_140_v3_)])
        d_270_v85_: _dafny.Array
        nw42_ = _dafny.Array(None, 21)
        nw42_[int(0)] = d_135_v0_
        nw42_[int(1)] = d_135_v0_
        nw42_[int(2)] = True
        nw42_[int(3)] = (d_263_v78_) < (d_263_v78_)
        nw42_[int(4)] = d_135_v0_
        nw42_[int(5)] = not((d_135_v0_) in (d_136_v1_))
        nw42_[int(6)] = d_135_v0_
        nw42_[int(7)] = not(d_135_v0_)
        nw42_[int(8)] = d_135_v0_
        nw42_[int(9)] = d_135_v0_
        nw42_[int(10)] = d_135_v0_
        nw42_[int(11)] = ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ryigcb")), d_144_v6_, d_144_v6_])).cardinality) == (d_140_v3_)
        nw42_[int(12)] = (_dafny.MultiSet([len(d_264_v79_)])).ispropersubset(d_269_v84_)
        nw42_[int(13)] = (d_135_v0_) or (d_135_v0_)
        nw42_[int(14)] = d_135_v0_
        nw42_[int(15)] = d_135_v0_
        nw42_[int(16)] = d_135_v0_
        nw42_[int(17)] = (d_135_v0_) or (d_135_v0_)
        nw42_[int(18)] = d_135_v0_
        nw42_[int(19)] = not(d_135_v0_)
        nw42_[int(20)] = not(d_135_v0_)
        d_270_v85_ = nw42_
        index40_ = default__.safeIndex(938, (d_270_v85_).length(0))
        (d_270_v85_)[index40_] = d_135_v0_
        d_271_v86_: str
        d_271_v86_ = _dafny.CodePoint('d')
        d_272_v87_: _dafny.Map
        d_272_v87_ = _dafny.Map({d_271_v86_: (0) - ((d_266_v81_).fm6(d_135_v0_, (D2_DC6(d_144_v6_, d_140_v3_, ((d_247_v72_)[d_140_v3_] if (d_140_v3_) in (d_247_v72_) else d_135_v0_), d_140_v3_, ((d_247_v72_)[540] if (540) in (d_247_v72_) else d_135_v0_))).cf12, not(d_135_v0_), d_143_globalState_))})
        d_273_v88_: _dafny.MultiSet
        d_273_v88_ = _dafny.MultiSet([d_135_v0_])
        index41_ = default__.safeIndex(938, (d_270_v85_).length(0))
        (d_270_v85_)[index41_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lthcc")))) > (((d_272_v87_)[d_271_v86_] if (d_271_v86_) in (d_272_v87_) else ((d_273_v88_)[d_135_v0_] if (d_135_v0_) in (d_273_v88_) else d_140_v3_)))
        default__.m0(d_135_v0_, (d_140_v3_) * (d_140_v3_), d_135_v0_, d_140_v3_, d_143_globalState_)
        d_274_v89_: _dafny.Map
        d_274_v89_ = _dafny.Map({d_265_v80_: d_144_v6_})
        d_275_v90_: D4
        d_275_v90_ = D4_DC9(_dafny.Map({(d_266_v81_).f14: d_144_v6_}))
        d_274_v89_ = (d_275_v90_).cf20
        d_276_v91_: _dafny.MultiSet
        d_276_v91_ = _dafny.MultiSet([d_264_v79_])
        rhs19_ = ((d_276_v91_) | (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_135_v0_]), d_264_v79_]))).cardinality
        rhs20_ = d_262_v77_
        rhs21_ = (default__.fm11(d_271_v86_, d_143_globalState_)).cf14
        lhs20_ = d_143_globalState_
        lhs21_ = d_143_globalState_
        lhs20_.f0 = rhs19_
        d_262_v77_ = rhs20_
        lhs21_.f0 = rhs21_
        d_277_i9_: int
        d_277_i9_ = 0
        with _dafny.label("0"):
            while d_135_v0_:
                with _dafny.c_label("0"):
                    if (d_277_i9_) >= (100):
                        raise _dafny.Break("0")
                    d_277_i9_ = (d_277_i9_) + (1)
                    default__.m0((d_270_v85_)[default__.safeIndex(938, (d_270_v85_).length(0))], d_140_v3_, (d_270_v85_)[default__.safeIndex(938, (d_270_v85_).length(0))], d_140_v3_, d_143_globalState_)
                    d_278_v92_: D0
                    d_278_v92_ = D0_DC0(_dafny.SeqWithoutIsStrInference([(d_270_v85_)[default__.safeIndex(938, (d_270_v85_).length(0))], d_135_v0_]))
                    source4_ = d_278_v92_
                    if source4_.is_DC1:
                        d_279___mcc_h15_ = source4_.cf1
                        d_280_cf1_ = d_279___mcc_h15_
                        d_281_v93_: _dafny.Array
                        nw43_ = _dafny.Array(_dafny.MultiSet({}), 21)
                        d_281_v93_ = nw43_
                        index42_ = default__.safeIndex(335, (d_281_v93_).length(0))
                        (d_281_v93_)[index42_] = d_269_v84_
                        index43_ = default__.safeIndex(938, (d_270_v85_).length(0))
                        index44_ = default__.safeIndex(335, (d_281_v93_).length(0))
                        rhs22_ = d_280_cf1_
                        rhs23_ = d_135_v0_
                        rhs24_ = _dafny.CodePoint('o')
                        rhs25_ = (_dafny.MultiSet((d_262_v77_) + (d_262_v77_))).intersection(default__.fm12(d_271_v86_, (d_270_v85_)[default__.safeIndex(938, (d_270_v85_).length(0))], d_280_cf1_, d_143_globalState_))
                        rhs26_ = d_267_v82_
                        lhs22_ = d_270_v85_
                        lhs23_ = default__.safeIndex(938, (d_270_v85_).length(0))
                        lhs24_ = d_143_globalState_
                        lhs25_ = d_281_v93_
                        lhs26_ = default__.safeIndex(335, (d_281_v93_).length(0))
                        d_140_v3_ = rhs22_
                        lhs22_[lhs23_] = rhs23_
                        lhs24_.f9 = rhs24_
                        lhs25_[lhs26_] = rhs25_
                        d_267_v82_ = rhs26_
                        d_282_v94_: C1
                        nw44_ = C1()
                        nw44_.ctor__()
                        d_282_v94_ = nw44_
                        d_283_v95_: C0
                        nw45_ = C0()
                        nw45_.ctor__((d_266_v81_).f14)
                        d_283_v95_ = nw45_
                        d_284_v96_: C0
                        nw46_ = C0()
                        nw46_.ctor__((d_266_v81_).f14)
                        d_284_v96_ = nw46_
                    elif source4_.is_DC2:
                        d_285___mcc_h16_ = source4_.cf2
                        d_286___mcc_h17_ = source4_.cf3
                        d_287___mcc_h18_ = source4_.cf4
                        d_288_cf4_ = d_287___mcc_h18_
                        d_289_cf3_ = d_286___mcc_h17_
                        d_290_cf2_ = d_285___mcc_h16_
                        d_291_v97_: _dafny.Array
                        nw47_ = _dafny.Array(_dafny.Array(None, 0), 1)
                        d_291_v97_ = nw47_
                        index45_ = default__.safeIndex(336, (d_291_v97_).length(0))
                        (d_291_v97_)[index45_] = d_270_v85_
                        d_292_v98_: _dafny.Array
                        nw48_ = _dafny.Array(None, 2)
                        nw48_[int(0)] = d_144_v6_
                        nw48_[int(1)] = d_144_v6_
                        d_292_v98_ = nw48_
                        index46_ = default__.safeIndex(580, (d_292_v98_).length(0))
                        (d_292_v98_)[index46_] = (d_144_v6_) + (d_144_v6_)
                        d_293_v99_: _dafny.Array
                        def lambda15_(d_294_cf4_):
                            def lambda16_(d_295_i10_):
                                return (d_295_i10_) - (d_294_cf4_)

                            return lambda16_

                        init7_ = lambda15_(d_288_cf4_)
                        nw49_ = _dafny.Array(None, 24)
                        for i0_7_ in range(nw49_.length(0)):
                            nw49_[i0_7_] = init7_(i0_7_)
                        d_293_v99_ = nw49_
                        index47_ = default__.safeIndex(40, (d_293_v99_).length(0))
                        (d_293_v99_)[index47_] = -345
                        index48_ = default__.safeIndex(336, (d_291_v97_).length(0))
                        index49_ = default__.safeIndex(580, (d_292_v98_).length(0))
                        index50_ = default__.safeIndex(40, (d_293_v99_).length(0))
                        rhs27_ = d_270_v85_
                        rhs28_ = d_144_v6_
                        rhs29_ = d_288_cf4_
                        rhs30_ = (default__.safeModuloInt(d_140_v3_, d_140_v3_)) > (d_288_cf4_)
                        lhs27_ = d_291_v97_
                        lhs28_ = default__.safeIndex(336, (d_291_v97_).length(0))
                        lhs29_ = d_292_v98_
                        lhs30_ = default__.safeIndex(580, (d_292_v98_).length(0))
                        lhs31_ = d_293_v99_
                        lhs32_ = default__.safeIndex(40, (d_293_v99_).length(0))
                        lhs27_[lhs28_] = rhs27_
                        lhs29_[lhs30_] = rhs28_
                        lhs31_[lhs32_] = rhs29_
                        d_290_cf2_ = rhs30_
                        index51_ = default__.safeIndex(40, (d_293_v99_).length(0))
                        rhs31_ = default__.fm0(d_143_globalState_)
                        rhs32_ = not((D4_DC12(d_135_v0_, d_263_v78_, (d_173_v31_).cf1, d_270_v85_, d_135_v0_)).cf26)
                        lhs33_ = d_293_v99_
                        lhs34_ = default__.safeIndex(40, (d_293_v99_).length(0))
                        lhs33_[lhs34_] = rhs31_
                        d_289_cf3_ = rhs32_
                        d_296_v100_: _dafny.Array
                        nw50_ = _dafny.Array(None, 16)
                        d_296_v100_ = nw50_
                        d_297_v101_: _dafny.Map
                        d_297_v101_ = _dafny.Map({d_289_cf3_: d_296_v100_})
                        d_297_v101_ = _dafny.Map({(d_270_v85_)[default__.safeIndex(938, (d_270_v85_).length(0))]: d_296_v100_})
                        d_140_v3_ = d_288_cf4_
                    elif True:
                        d_298___mcc_h19_ = source4_.cf0
                        d_299_cf0_ = d_298___mcc_h19_
                        d_300_v102_: _dafny.Array
                        nw51_ = _dafny.Array(int(0), 27)
                        d_300_v102_ = nw51_
                        d_301_v103_: _dafny.Map
                        d_301_v103_ = _dafny.Map({d_135_v0_: d_300_v102_})
                        d_302_v104_: D3
                        d_302_v104_ = D3_DC7(((d_301_v103_)[(d_264_v79_)[default__.safeIndex(d_140_v3_, len(d_264_v79_))]] if ((d_264_v79_)[default__.safeIndex(d_140_v3_, len(d_264_v79_))]) in (d_301_v103_) else d_300_v102_))
                        d_302_v104_ = d_302_v104_
                        d_303_v105_: C1
                        nw52_ = C1()
                        nw52_.ctor__()
                        d_303_v105_ = nw52_
                        d_304_v106_: _dafny.Seq
                        d_304_v106_ = _dafny.SeqWithoutIsStrInference([d_303_v105_])
                        index52_ = default__.safeIndex(938, (d_270_v85_).length(0))
                        (d_270_v85_)[index52_] = (default__.safeDivisionInt(len(d_299_cf0_), d_140_v3_)) < (len((d_304_v106_) + (_dafny.SeqWithoutIsStrInference([d_303_v105_, d_303_v105_]))))
                        (d_143_globalState_).f0 = default__.safeModuloInt(d_140_v3_, d_140_v3_)
                        (d_143_globalState_).f0 = d_140_v3_
                    d_305_v107_: C1
                    nw53_ = C1()
                    nw53_.ctor__()
                    d_305_v107_ = nw53_
                    d_306_v108_: _dafny.Seq
                    d_306_v108_ = _dafny.SeqWithoutIsStrInference([d_305_v107_, d_305_v107_, d_305_v107_, d_305_v107_])
                    (d_143_globalState_).f0 = (d_140_v3_) + (len((d_306_v108_) + (d_306_v108_)))
                    (d_143_globalState_).f0 = (d_140_v3_) * (d_140_v3_)
                    pass
            pass
        d_307_v109_: _dafny.Map
        d_307_v109_ = _dafny.Map({d_144_v6_: (d_270_v85_)[default__.safeIndex(938, (d_270_v85_).length(0))]})
        d_308_v110_: _dafny.Map
        d_308_v110_ = _dafny.Map({d_136_v1_: ((d_261_v76_)[d_140_v3_] if (d_140_v3_) in (d_261_v76_) else d_140_v3_)})
        d_307_v109_ = (d_307_v109_).set(d_144_v6_, (_dafny.Set({d_135_v0_})) not in (d_308_v110_))
        _dafny.print(_dafny.string_of(d_135_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v1_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_137_v2_)[0]) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_140_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_141_v4_) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[0]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[1]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[2]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[3]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[4]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[5]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[6]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[7]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[8]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[9]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[10]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[11]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[12]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[13]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[14]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[15]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[16]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[17]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[18]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[19]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_142_v5_)[20]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_).f1) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_143_globalState_).f2)[0]) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[0]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[1]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[2]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[3]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[4]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[5]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[6]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[7]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[8]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[9]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[10]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[11]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[12]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[13]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[14]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[15]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[16]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[17]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[18]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[19]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_143_globalState_.f3)[20]) == (_dafny.Map({True: -730}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_143_globalState_).f5).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_143_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_143_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_144_v6_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_173_v31_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v48_).cf6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v48_).cf7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v48_).cf8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v48_).cf9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_247_v72_) == (_dafny.Map({1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v73_).cf2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v73_).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v73_).cf4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_259_v74_) == (_dafny.Map({D1_DC4(423, True, True, True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whhg"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_260_v75_).cf11).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v75_).cf12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v75_).cf13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v75_).cf14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_260_v75_).cf15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v76_) == (_dafny.Map({1: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_262_v77_) == (_dafny.SeqWithoutIsStrInference([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_263_v78_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({1: 1}), _dafny.Map({4: 1}), _dafny.Map({1: 1}), _dafny.Map({1: 1}), _dafny.Map({1: 1})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_264_v79_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v80_).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v80_).cf18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_265_v80_).cf19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_266_v81_).f14).cf17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_266_v81_).f14).cf18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_266_v81_).f14).cf19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_267_v82_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v83_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v84_) == (_dafny.MultiSet([1, 1, -1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_270_v85_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_271_v86_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_272_v87_) == (_dafny.Map({_dafny.CodePoint('d'): 615}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_273_v88_) == (_dafny.MultiSet([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_274_v89_) == (_dafny.Map({D3_DC8(1, 1, False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whhg"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_275_v90_).cf20) == (_dafny.Map({D3_DC8(1, 1, False): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whhg"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_276_v91_) == (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_277_i9_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_307_v109_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whhg")): False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_308_v110_) == (_dafny.Map({_dafny.Set({True}): 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4
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
        return lambda: D1_DC4(int(0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({self.cf11.VerbatimString(True)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({self.cf10.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC8(int(0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)

class D3_DC8(D3, NamedTuple('DC8', [('cf17', Any), ('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf17 == __o.cf17 and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC7(D3, NamedTuple('DC7', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC10(_dafny.Seq({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D4_DC9)

class D4_DC10(D4, NamedTuple('DC10', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC11(D4, NamedTuple('DC11', [('cf23', Any), ('cf24', Any), ('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf26', Any), ('cf27', Any), ('cf28', Any), ('cf29', Any), ('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC9(D4, NamedTuple('DC9', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC9({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC9) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.CodePoint('D')
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC13(D5, NamedTuple('DC13', [('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D6_DC14)

class D6_DC14(D6, NamedTuple('DC14', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC14({_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC14) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f3: _dafny.Array = _dafny.Array(None, 0)
        self.f4: str = _dafny.CodePoint('D')
        self.f9: str = _dafny.CodePoint('D')
        self._f1: _dafny.Set = _dafny.Set({})
        self._f2: _dafny.Array = _dafny.Array(None, 0)
        self._f5: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f6: int = int(0)
        self._f7: int = int(0)
        self._f8: bool = False
        self._f10: bool = False
        self._f11: int = int(0)
        self._f12: bool = False
        self._f13: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13):
        (self).f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self).f9 = f9
        (self)._f10 = f10
        (self)._f11 = f11
        (self)._f12 = f12
        (self)._f13 = f13

    @property
    def f1(self):
        return self._f1
    @property
    def f2(self):
        return self._f2
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
    def f10(self):
        return self._f10
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
        self._f14: D3 = D3.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f14):
        (self)._f14 = f14

    def fm6(self, p0, p1, p2, globalState):
        return (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([-847, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqfinko")))])), 15)) - (615)

    @property
    def f14(self):
        return self._f14

class C1:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self):
        pass
        pass

    def fm4(self, globalState):
        def iife45_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Set
            for compr_5_ in (_dafny.Map({_dafny.Set({True}): False})).keys.Elements:
                d_309_v0_: _dafny.Set = compr_5_
                if (d_309_v0_) in (_dafny.Map({_dafny.Set({True}): False})):
                    coll5_[d_309_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "byiwjxmen"))
            return _dafny.Map(coll5_)
        return ((D1_DC3(_dafny.Map({_dafny.Set({True, False, False}): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "saf"))}))).cf5) | (iife45_()
        )

    def fm5(self, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_310_i0_ in range(default__.abs(209))]): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kipi"))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbvdrrtth")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))}))) | (_dafny.Map({(D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rt")))).cf10: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goeperti"))}))

    def m1(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        r2: _dafny.Set = _dafny.Set({})
        r3: int = int(0)
        d_311_v0_: int
        d_311_v0_ = 52
        d_312_v1_: D0
        d_312_v1_ = D0_DC1(d_311_v0_)
        pat_let_tv14_ = d_311_v0_
        def iife46_(_pat_let20_0):
            def iife47_(d_313_dt__update__tmp_h0_):
                def iife48_(_pat_let21_0):
                    def iife49_(d_314_dt__update_hcf1_h0_):
                        return D0_DC1(d_314_dt__update_hcf1_h0_)
                    return iife49_(_pat_let21_0)
                return iife48_(pat_let_tv14_)
            return iife47_(_pat_let20_0)
        r3 = (iife46_(d_312_v1_)).cf1
        hi3_ = (d_311_v0_) + (d_311_v0_)
        for d_315_i0_ in range(len(p1), hi3_):
            r1 = not(p2)
            rhs33_ = p2
            rhs34_ = d_315_i0_
            rhs35_ = p2
            r1 = rhs33_
            d_311_v0_ = rhs34_
            r1 = rhs35_
            r1 = p2
            d_316_v2_: _dafny.Seq
            d_316_v2_ = _dafny.SeqWithoutIsStrInference([len((p1) + (p1))])
            rhs36_ = (d_316_v2_)[default__.safeIndex(d_311_v0_, len(d_316_v2_))]
            rhs37_ = default__.fm2(globalState)
            lhs35_ = globalState
            lhs35_.f0 = rhs36_
            r1 = rhs37_
        d_317_v3_: _dafny.Array
        nw54_ = _dafny.Array(False, 7)
        d_317_v3_ = nw54_
        index53_ = default__.safeIndex(729, (d_317_v3_).length(0))
        (d_317_v3_)[index53_] = p2
        index54_ = default__.safeIndex(729, (d_317_v3_).length(0))
        (d_317_v3_)[index54_] = not (not(p2)) or (p2)
        index55_ = default__.safeIndex(729, (d_317_v3_).length(0))
        rhs38_ = not((d_311_v0_) >= (len(p1)))
        rhs39_ = p2
        lhs36_ = d_317_v3_
        lhs37_ = default__.safeIndex(729, (d_317_v3_).length(0))
        r1 = rhs38_
        lhs36_[lhs37_] = rhs39_
        hi4_ = d_311_v0_
        for d_318_i1_ in range((d_311_v0_) - (d_311_v0_), hi4_):
            d_319_v4_: D0
            d_319_v4_ = D0_DC2((d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))], (d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))], d_318_i1_)
            d_320_v5_: _dafny.Seq
            d_320_v5_ = _dafny.SeqWithoutIsStrInference([d_319_v4_, d_319_v4_])
            d_321_v6_: _dafny.Map
            d_321_v6_ = _dafny.Map({d_311_v0_: 178})
            d_322_v7_: _dafny.MultiSet
            d_322_v7_ = _dafny.MultiSet([d_321_v6_, d_321_v6_])
            d_323_v8_: _dafny.Seq
            d_323_v8_ = _dafny.SeqWithoutIsStrInference([p2, (d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))]])
            d_324_v9_: _dafny.Set
            d_324_v9_ = _dafny.Set({d_323_v8_})
            d_325_v10_: _dafny.Seq
            d_325_v10_ = _dafny.SeqWithoutIsStrInference([len(d_320_v5_), (0) - ((d_322_v7_).cardinality), d_318_i1_, len(d_324_v9_)])
            index56_ = default__.safeIndex(729, (d_317_v3_).length(0))
            rhs40_ = d_311_v0_
            rhs41_ = (d_325_v10_) == ((d_325_v10_) + (d_325_v10_))
            lhs38_ = d_317_v3_
            lhs39_ = default__.safeIndex(729, (d_317_v3_).length(0))
            r3 = rhs40_
            lhs38_[lhs39_] = rhs41_
            if default__.fm2(globalState):
                d_326_v11_: _dafny.Map
                d_326_v11_ = _dafny.Map({p2: default__.fm2(globalState)})
                d_327_v12_: _dafny.MultiSet
                d_327_v12_ = _dafny.MultiSet([default__.fm2(globalState), p2])
                d_326_v11_ = (d_326_v11_).set((_dafny.MultiSet([(d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))]])).ispropersubset(d_327_v12_), (d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))])
                d_326_v11_ = (d_326_v11_).set(p2, p2)
                d_328_v13_: _dafny.Array
                nw55_ = _dafny.Array(int(0), 23)
                d_328_v13_ = nw55_
                d_329_v14_: D3
                d_329_v14_ = D3_DC7(d_328_v13_)
                d_330_v15_: _dafny.Map
                d_330_v15_ = _dafny.Map({(d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))]: (d_329_v14_).cf16})
                d_331_v16_: _dafny.Seq
                d_331_v16_ = _dafny.SeqWithoutIsStrInference([d_328_v13_, d_328_v13_, d_328_v13_, d_328_v13_])
                d_330_v15_ = (d_330_v15_).set((d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))], (d_331_v16_)[default__.safeIndex(d_318_i1_, len(d_331_v16_))])
                d_332_v17_: _dafny.Array
                nw56_ = _dafny.Array(_dafny.Seq({}), 29)
                d_332_v17_ = nw56_
                index57_ = default__.safeIndex(222, (d_332_v17_).length(0))
                (d_332_v17_)[index57_] = (d_331_v16_) + (d_331_v16_)
                d_333_v18_: _dafny.Map
                d_333_v18_ = _dafny.Map({d_311_v0_: p2})
                d_334_v19_: _dafny.Seq
                d_334_v19_ = _dafny.SeqWithoutIsStrInference([d_333_v18_, d_333_v18_])
                index58_ = default__.safeIndex(222, (d_332_v17_).length(0))
                (d_332_v17_)[index58_] = (((d_331_v16_) + (d_331_v16_)) + (d_331_v16_)).set(default__.safeIndex(default__.safeDivisionInt(len(d_334_v19_), len(p1)), len(((d_331_v16_) + (d_331_v16_)) + (d_331_v16_))), ((d_330_v15_)[p2] if (p2) in (d_330_v15_) else d_328_v13_))
                d_326_v11_ = (d_326_v11_).set(p2, p2)
            elif True:
                d_311_v0_ = len(d_323_v8_)
                d_335_v20_: str
                d_335_v20_ = _dafny.CodePoint('q')
                (globalState).f4 = d_335_v20_
                d_321_v6_ = (d_321_v6_).set(d_318_i1_, d_311_v0_)
                index59_ = default__.safeIndex(729, (d_317_v3_).length(0))
                (d_317_v3_)[index59_] = default__.fm2(globalState)
                d_336_v21_: _dafny.Array
                nw57_ = _dafny.Array(int(0), 1)
                d_336_v21_ = nw57_
                d_337_v22_: _dafny.Seq
                d_337_v22_ = _dafny.SeqWithoutIsStrInference([d_336_v21_, d_336_v21_])
                d_338_v23_: _dafny.Map
                d_338_v23_ = _dafny.Map({D3_DC7((d_337_v22_)[default__.safeIndex(716, len(d_337_v22_))]): d_311_v0_})
                d_339_v24_: _dafny.Array
                nw58_ = _dafny.Array(None, 1)
                nw58_[int(0)] = d_338_v23_
                d_339_v24_ = nw58_
                index60_ = default__.safeIndex(445, (d_339_v24_).length(0))
                (d_339_v24_)[index60_] = (d_338_v23_) | (d_338_v23_)
                index61_ = default__.safeIndex(445, (d_339_v24_).length(0))
                (d_339_v24_)[index61_] = d_338_v23_
            (globalState).f0 = d_311_v0_
            d_340_v25_: D3
            d_340_v25_ = D3_DC8(d_318_i1_, d_311_v0_, (d_317_v3_)[default__.safeIndex(729, (d_317_v3_).length(0))])
            d_341_v26_: C0
            nw59_ = C0()
            nw59_.ctor__(d_340_v25_)
            d_341_v26_ = nw59_
        (globalState).f0 = default__.safeDivisionInt(d_311_v0_, (d_311_v0_) * (len(p1)))
        r0 = (0) - (len(p1))
        r1 = not(p2)
        d_342_v27_: _dafny.Set
        d_342_v27_ = _dafny.Set({d_311_v0_})
        r2 = (d_342_v27_) | (d_342_v27_)
        r3 = default__.safeDivisionInt(((0) - (d_311_v0_)) - (d_311_v0_), d_311_v0_)
        return r0, r1, r2, r3

