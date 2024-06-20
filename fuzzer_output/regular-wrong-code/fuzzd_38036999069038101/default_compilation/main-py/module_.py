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
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(421, 316):
                d_0_v0_: int = compr_0_
                if ((421) <= (d_0_v0_)) and ((d_0_v0_) < (316)):
                    coll0_[default__.safeModuloInt(d_0_v0_, -169)] = False
            return _dafny.Map(coll0_)
        return (iife0_()
        ) | (_dafny.Map({296: not(False)}))

    @staticmethod
    def fm1(globalState):
        return _dafny.CodePoint('l')

    @staticmethod
    def fm2(globalState):
        return not(((_dafny.Map({573: len(_dafny.SeqWithoutIsStrInference([not(not(True)), not(True), False]))})) | (_dafny.Map({(0) - (-447): 235}))) == ((_dafny.Map({len(_dafny.Map({_dafny.CodePoint('q'): True})): 271})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([321 for d_1_i0_ in range(default__.abs(-452))])): 4}))))

    @staticmethod
    def fm3(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([369])).Elements:
                d_2_v0_: int = compr_1_
                if (d_2_v0_) in (_dafny.SeqWithoutIsStrInference([369])):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_2_v0_, 526)]))
            return _dafny.Set(coll1_)
        return _dafny.Map({(iife1_()
        ).isdisjoint(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([261]))})): False})

    @staticmethod
    def fm4(globalState):
        return _dafny.MultiSet([((0) - (len(_dafny.Set({True})))) - (len(_dafny.Map({True: not(False)}))), 247])

    @staticmethod
    def fm5(globalState):
        return D1_DC3((len(_dafny.Map({_dafny.MultiSet([False, False, True, True, not(False)]): 376}))) < (-423), True, default__.safeDivisionInt(117, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sebrnl")))), True)

    @staticmethod
    def fm6(p0, globalState):
        return (0) - (len(((_dafny.Map({_dafny.MultiSet([True]): True})) | (_dafny.Map({_dafny.MultiSet([False]): True}))) | (_dafny.Map({_dafny.MultiSet([False]): True}))))

    @staticmethod
    def fm7(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wfhne")))

    @staticmethod
    def fm8(p0, p1, p2, globalState):
        return _dafny.Set({((D3_DC9(80)).cf17) * (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gy"))): len(_dafny.SeqWithoutIsStrInference([False, True]))})))})

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return (_dafny.Set({not(True)})) | ((_dafny.Set({True})) | (_dafny.Set({False})))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: _dafny.Seq
            for compr_2_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3_i0_ in range(default__.abs(127))]): len(_dafny.Map({True: len(_dafny.Set({(_dafny.MultiSet([211])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xom"))), 419}))}))})).keys.Elements:
                d_4_v0_: _dafny.Seq = compr_2_
                if (d_4_v0_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_3_i0_ in range(default__.abs(127))]): len(_dafny.Map({True: len(_dafny.Set({(_dafny.MultiSet([211])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xom"))), 419}))}))})):
                    coll2_[d_4_v0_] = True
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([352 for d_8_i5_ in range(default__.abs(500))])) for d_9_i4_ in range(default__.abs(948))])).Elements:
                d_10_v1_: int = compr_3_
                if (d_10_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([352 for d_8_i5_ in range(default__.abs(500))])) for d_9_i4_ in range(default__.abs(948))])):
                    coll3_[default__.safeDivisionInt(d_10_v1_, 967)] = True
            return _dafny.Map(coll3_)
        return D1_DC2(len((_dafny.SeqWithoutIsStrInference([972, len(iife2_()
), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_5_i1_ in range(default__.abs(785))])), -884])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([849]))]))), _dafny.CodePoint('r'), (-240) * (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_6_i3_ in range(default__.abs(593))])) for d_7_i2_ in range(default__.abs(534))]))), (_dafny.SeqWithoutIsStrInference([len(iife3_()
), -539])) + ((D1_DC2(70, _dafny.CodePoint('v'), len(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([-71, 213])).cardinality) for d_11_i6_ in range(default__.abs(515))])), _dafny.SeqWithoutIsStrInference([66]), _dafny.CodePoint('d'))).cf7), _dafny.CodePoint('m'))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-935, 330):
                d_12_v0_: int = compr_4_
                if ((-935) <= (d_12_v0_)) and ((d_12_v0_) < (330)):
                    coll4_[(d_12_v0_) * (len(_dafny.Set({not(False)})))] = (_dafny.MultiSet([False])) - (_dafny.MultiSet([not(not(True))]))
            return _dafny.Map(coll4_)
        return iife4_()
        

    @staticmethod
    def m0(p0, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: bool = False
        r3: _dafny.Set = _dafny.Set({})
        d_13_v0_: bool
        d_13_v0_ = False
        d_14_v1_: C0
        nw0_ = C0()
        nw0_.ctor__()
        d_14_v1_ = nw0_
        d_15_v2_: _dafny.Map
        d_15_v2_ = _dafny.Map({d_13_v0_: d_13_v0_})
        d_16_v3_: _dafny.Map
        d_16_v3_ = _dafny.Map({d_14_v1_: d_15_v2_})
        r1 = (p0) >= ((len(_dafny.SeqWithoutIsStrInference([d_13_v0_]))) * (len(d_16_v3_)))
        d_17_v4_: _dafny.Seq
        d_17_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        d_18_v5_: D3
        d_18_v5_ = D3_DC10(not(d_13_v0_), d_17_v4_, d_13_v0_)
        d_19_v6_: D3
        d_19_v6_ = D3_DC11(d_18_v5_)
        d_19_v6_ = d_19_v6_
        d_20_v7_: _dafny.Set
        d_20_v7_ = _dafny.Set({len(d_17_v4_)})
        d_21_i0_: int
        d_21_i0_ = 0
        with _dafny.label("0"):
            while (d_20_v7_).ispropersubset(d_20_v7_):
                with _dafny.c_label("0"):
                    if (d_21_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_21_i0_ = (d_21_i0_) + (1)
                    d_22_v8_: str
                    d_22_v8_ = _dafny.CodePoint('n')
                    d_22_v8_ = d_22_v8_
                    if not(not(default__.fm2(globalState))):
                        d_23_v9_: _dafny.Set
                        d_23_v9_ = _dafny.Set({d_13_v0_})
                        d_24_v10_: _dafny.Seq
                        d_24_v10_ = _dafny.SeqWithoutIsStrInference([d_23_v9_])
                        d_24_v10_ = d_24_v10_
                        d_25_v11_: D3
                        d_25_v11_ = D3_DC10(d_13_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfugb")), d_13_v0_)
                        d_26_v12_: _dafny.Array
                        nw1_ = _dafny.Array(None, 19)
                        nw1_[int(0)] = d_13_v0_
                        nw1_[int(1)] = d_13_v0_
                        nw1_[int(2)] = d_13_v0_
                        nw1_[int(3)] = True
                        nw1_[int(4)] = d_13_v0_
                        nw1_[int(5)] = default__.fm2(globalState)
                        nw1_[int(6)] = d_13_v0_
                        nw1_[int(7)] = d_13_v0_
                        nw1_[int(8)] = d_13_v0_
                        nw1_[int(9)] = d_13_v0_
                        nw1_[int(10)] = d_13_v0_
                        nw1_[int(11)] = d_13_v0_
                        nw1_[int(12)] = d_13_v0_
                        nw1_[int(13)] = d_13_v0_
                        nw1_[int(14)] = d_13_v0_
                        nw1_[int(15)] = d_13_v0_
                        nw1_[int(16)] = d_13_v0_
                        nw1_[int(17)] = d_13_v0_
                        nw1_[int(18)] = d_13_v0_
                        d_26_v12_ = nw1_
                        d_27_v13_: _dafny.Map
                        d_27_v13_ = _dafny.Map({d_26_v12_: d_13_v0_})
                        pat_let_tv0_ = d_27_v13_
                        pat_let_tv1_ = d_26_v12_
                        pat_let_tv2_ = d_26_v12_
                        pat_let_tv3_ = d_27_v13_
                        pat_let_tv4_ = d_13_v0_
                        pat_let_tv5_ = d_20_v7_
                        pat_let_tv6_ = d_20_v7_
                        def iife6_(_pat_let1_0):
                            def iife7_(d_28_dt__update__tmp_h0_):
                                def iife8_(_pat_let2_0):
                                    def iife9_(d_29_dt__update_hcf20_h0_):
                                        return D3_DC10((d_28_dt__update__tmp_h0_).cf18, (d_28_dt__update__tmp_h0_).cf19, d_29_dt__update_hcf20_h0_)
                                    return iife9_(_pat_let2_0)
                                return iife8_(not(((pat_let_tv0_)[pat_let_tv1_] if (pat_let_tv2_) in (pat_let_tv3_) else pat_let_tv4_)))
                            return iife7_(_pat_let1_0)
                        def iife5_(_pat_let0_0):
                            def iife10_(d_30_dt__update__tmp_h1_):
                                def iife11_(_pat_let3_0):
                                    def iife12_(d_31_dt__update_hcf20_h1_):
                                        return D3_DC10((d_30_dt__update__tmp_h1_).cf18, (d_30_dt__update__tmp_h1_).cf19, d_31_dt__update_hcf20_h1_)
                                    return iife12_(_pat_let3_0)
                                return iife11_((pat_let_tv5_).ispropersubset(pat_let_tv6_))
                            return iife10_(_pat_let0_0)
                        d_25_v11_ = iife5_(iife6_(d_25_v11_))
                        d_13_v0_ = d_13_v0_
                        d_14_v1_ = d_14_v1_
                        d_32_v14_: D5
                        d_32_v14_ = D5_DC14(d_13_v0_, d_13_v0_, d_14_v1_, 151)
                        pat_let_tv7_ = d_15_v2_
                        pat_let_tv8_ = d_13_v0_
                        pat_let_tv9_ = d_13_v0_
                        pat_let_tv10_ = d_15_v2_
                        pat_let_tv11_ = d_13_v0_
                        d_33_v15_: _dafny.Seq
                        def iife13_(_pat_let4_0):
                            def iife14_(d_34_dt__update__tmp_h2_):
                                def iife15_(_pat_let5_0):
                                    def iife16_(d_35_dt__update_hcf25_h0_):
                                        return D5_DC14((d_34_dt__update__tmp_h2_).cf24, d_35_dt__update_hcf25_h0_, (d_34_dt__update__tmp_h2_).cf26, (d_34_dt__update__tmp_h2_).cf27)
                                    return iife16_(_pat_let5_0)
                                return iife15_(((pat_let_tv7_)[pat_let_tv8_] if (pat_let_tv9_) in (pat_let_tv10_) else pat_let_tv11_))
                            return iife14_(_pat_let4_0)
                        d_33_v15_ = _dafny.SeqWithoutIsStrInference([(iife13_(d_32_v14_)).cf26])
                        d_14_v1_ = (d_33_v15_)[default__.safeIndex(p0, len(d_33_v15_))]
                    elif True:
                        d_36_v16_: _dafny.Array
                        nw2_ = _dafny.Array(False, 6)
                        d_36_v16_ = nw2_
                        index0_ = default__.safeIndex(351, (d_36_v16_).length(0))
                        (d_36_v16_)[index0_] = (p0) > (p0)
                        d_37_v17_: _dafny.MultiSet
                        d_37_v17_ = _dafny.MultiSet([not(d_13_v0_), d_13_v0_])
                        d_38_v18_: _dafny.Set
                        d_38_v18_ = _dafny.Set({d_13_v0_, d_13_v0_})
                        d_39_v19_: _dafny.Array
                        nw3_ = _dafny.Array(None, 5)
                        nw3_[int(0)] = (p0) - (p0)
                        nw3_[int(1)] = (d_37_v17_).cardinality
                        nw3_[int(2)] = p0
                        nw3_[int(3)] = p0
                        nw3_[int(4)] = ((0) - (len(d_38_v18_)) if False else p0)
                        d_39_v19_ = nw3_
                        index1_ = default__.safeIndex(344, (d_39_v19_).length(0))
                        (d_39_v19_)[index1_] = p0
                        d_40_v20_: _dafny.Array
                        nw4_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                        d_40_v20_ = nw4_
                        d_41_v21_: _dafny.Map
                        d_41_v21_ = _dafny.Map({False: d_40_v20_})
                        d_42_v22_: _dafny.Map
                        d_42_v22_ = d_41_v21_
                        d_43_v23_: _dafny.MultiSet
                        d_43_v23_ = _dafny.MultiSet([p0])
                        index2_ = default__.safeIndex(74, (d_39_v19_).length(0))
                        (d_39_v19_)[index2_] = (len((d_42_v22_))) * ((d_43_v23_).cardinality)
                        index3_ = default__.safeIndex(351, (d_36_v16_).length(0))
                        index4_ = default__.safeIndex(344, (d_39_v19_).length(0))
                        index5_ = default__.safeIndex(74, (d_39_v19_).length(0))
                        rhs0_ = (default__.fm7(p0, globalState)) <= (d_17_v4_)
                        rhs1_ = p0
                        rhs2_ = default__.safeDivisionInt(-370, p0)
                        rhs3_ = (0) - (p0)
                        lhs0_ = d_36_v16_
                        lhs1_ = default__.safeIndex(351, (d_36_v16_).length(0))
                        lhs2_ = d_39_v19_
                        lhs3_ = default__.safeIndex(344, (d_39_v19_).length(0))
                        lhs4_ = globalState
                        lhs5_ = d_39_v19_
                        lhs6_ = default__.safeIndex(74, (d_39_v19_).length(0))
                        lhs0_[lhs1_] = rhs0_
                        lhs2_[lhs3_] = rhs1_
                        lhs4_.f2 = rhs2_
                        lhs5_[lhs6_] = rhs3_
                        d_44_v24_: _dafny.Array
                        def lambda0_(d_45_v8_):
                            def lambda1_(d_46_i1_):
                                return _dafny.SeqWithoutIsStrInference([d_45_v8_ for d_47_i2_ in range(default__.abs(-298))])

                            return lambda1_

                        init0_ = lambda0_(d_22_v8_)
                        nw5_ = _dafny.Array(None, 22)
                        for i0_0_ in range(nw5_.length(0)):
                            nw5_[i0_0_] = init0_(i0_0_)
                        d_44_v24_ = nw5_
                        rhs4_ = d_44_v24_
                        rhs5_ = (d_36_v16_)[default__.safeIndex(351, (d_36_v16_).length(0))]
                        rhs6_ = (d_43_v23_) - (d_43_v23_)
                        rhs7_ = default__.fm7(((d_39_v19_)[default__.safeIndex(344, (d_39_v19_).length(0))] if (d_36_v16_)[default__.safeIndex(351, (d_36_v16_).length(0))] else (d_39_v19_)[default__.safeIndex(344, (d_39_v19_).length(0))]), globalState)
                        lhs7_ = globalState
                        d_44_v24_ = rhs4_
                        r1 = rhs5_
                        d_43_v23_ = rhs6_
                        lhs7_.f11 = rhs7_
                        d_48_v25_: _dafny.Seq
                        d_48_v25_ = _dafny.SeqWithoutIsStrInference([(d_39_v19_)[default__.safeIndex(344, (d_39_v19_).length(0))]])
                        d_49_v26_: _dafny.Seq
                        d_49_v26_ = _dafny.SeqWithoutIsStrInference([p0, (d_39_v19_)[default__.safeIndex(344, (d_39_v19_).length(0))], len(d_48_v25_)])
                        d_50_v27_: _dafny.Seq
                        d_50_v27_ = _dafny.SeqWithoutIsStrInference([d_49_v26_, d_48_v25_])
                        d_51_v28_: _dafny.MultiSet
                        d_51_v28_ = _dafny.MultiSet([d_49_v26_, (d_50_v27_)[default__.safeIndex((d_39_v19_)[default__.safeIndex(344, (d_39_v19_).length(0))], len(d_50_v27_))]])
                        d_52_v29_: _dafny.Map
                        d_52_v29_ = _dafny.Map({d_17_v4_: d_51_v28_})
                        d_52_v29_ = (d_52_v29_).set(d_17_v4_, d_51_v28_)
                        rhs8_ = d_14_v1_
                        rhs9_ = (d_36_v16_)[default__.safeIndex(351, (d_36_v16_).length(0))]
                        rhs10_ = (not(d_13_v0_) if (d_36_v16_)[default__.safeIndex(351, (d_36_v16_).length(0))] else (d_36_v16_)[default__.safeIndex(351, (d_36_v16_).length(0))])
                        d_14_v1_ = rhs8_
                        r1 = rhs9_
                        d_13_v0_ = rhs10_
                        d_53_v30_: C0
                        nw6_ = C0()
                        nw6_.ctor__()
                        d_53_v30_ = nw6_
                    d_54_v32_: _dafny.Seq
                    d_54_v32_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (p0)])
                    d_55_v33_: D1
                    def iife17_():
                        coll5_ = _dafny.Map()
                        compr_5_: int
                        for compr_5_ in (d_20_v7_).Elements:
                            d_56_v31_: int = compr_5_
                            if (d_56_v31_) in (d_20_v7_):
                                coll5_[(d_56_v31_) + (p0)] = ((d_15_v2_)[True] if (True) in (d_15_v2_) else d_13_v0_)
                        return _dafny.Map(coll5_)
                    d_55_v33_ = D1_DC2(214, d_22_v8_, len(iife17_()
), d_54_v32_, d_22_v8_)
                    source0_ = D1_DC4(d_55_v33_)
                    if source0_.is_DC2:
                        d_57___mcc_h0_ = source0_.cf4
                        d_58___mcc_h1_ = source0_.cf5
                        d_59___mcc_h2_ = source0_.cf6
                        d_60___mcc_h3_ = source0_.cf7
                        d_61___mcc_h4_ = source0_.cf8
                        d_62_cf8_ = d_61___mcc_h4_
                        d_63_cf7_ = d_60___mcc_h3_
                        d_64_cf6_ = d_59___mcc_h2_
                        d_65_cf5_ = d_58___mcc_h1_
                        d_66_cf4_ = d_57___mcc_h0_
                        d_67_v34_: _dafny.Array
                        def lambda2_(d_68_cf4_):
                            def lambda3_(d_69_i3_):
                                return (d_69_i3_) + (d_68_cf4_)

                            return lambda3_

                        init1_ = lambda2_(d_66_cf4_)
                        nw7_ = _dafny.Array(None, 5)
                        for i0_1_ in range(nw7_.length(0)):
                            nw7_[i0_1_] = init1_(i0_1_)
                        d_67_v34_ = nw7_
                        index6_ = default__.safeIndex(407, (d_67_v34_).length(0))
                        (d_67_v34_)[index6_] = d_66_cf4_
                        d_70_v35_: D3
                        d_70_v35_ = D3_DC9(d_66_cf4_)
                        index7_ = default__.safeIndex(407, (d_67_v34_).length(0))
                        (d_67_v34_)[index7_] = (d_70_v35_).cf17
                        d_71_v36_: C0
                        nw8_ = C0()
                        nw8_.ctor__()
                        d_71_v36_ = nw8_
                        d_72_v37_: _dafny.Map
                        d_72_v37_ = _dafny.Map({d_66_cf4_: d_64_cf6_})
                        d_72_v37_ = (d_72_v37_).set(d_64_cf6_, (d_64_cf6_) * (d_66_cf4_))
                        d_73_v38_: _dafny.Array
                        nw9_ = _dafny.Array(_dafny.Map({}), 4)
                        d_73_v38_ = nw9_
                        d_74_v39_: _dafny.Array
                        nw10_ = _dafny.Array(None, 6)
                        nw10_[int(0)] = default__.fm2(globalState)
                        nw10_[int(1)] = d_13_v0_
                        nw10_[int(2)] = d_13_v0_
                        nw10_[int(3)] = d_13_v0_
                        nw10_[int(4)] = d_13_v0_
                        nw10_[int(5)] = d_13_v0_
                        d_74_v39_ = nw10_
                        d_75_v40_: _dafny.Map
                        d_75_v40_ = _dafny.Map({d_74_v39_: d_71_v36_})
                        index8_ = default__.safeIndex(630, (d_73_v38_).length(0))
                        (d_73_v38_)[index8_] = d_75_v40_
                        d_76_v41_: _dafny.Seq
                        d_76_v41_ = _dafny.SeqWithoutIsStrInference([d_13_v0_, d_13_v0_])
                        d_77_v42_: D2
                        d_77_v42_ = D2_DC5(d_76_v41_)
                        d_78_v43_: _dafny.Map
                        d_78_v43_ = _dafny.Map({d_77_v42_: d_63_cf7_})
                        d_79_v44_: _dafny.Map
                        d_79_v44_ = _dafny.Map({d_71_v36_: (d_64_cf6_) * (len(((d_78_v43_)[d_77_v42_] if (d_77_v42_) in (d_78_v43_) else d_63_cf7_)))})
                        index9_ = default__.safeIndex(630, (d_73_v38_).length(0))
                        index10_ = default__.safeIndex(407, (d_67_v34_).length(0))
                        rhs11_ = d_14_v1_
                        rhs12_ = d_75_v40_
                        rhs13_ = (len((d_20_v7_) | (d_20_v7_))) - ((0) - (p0))
                        rhs14_ = d_66_cf4_
                        rhs15_ = d_79_v44_
                        lhs8_ = d_73_v38_
                        lhs9_ = default__.safeIndex(630, (d_73_v38_).length(0))
                        lhs10_ = globalState
                        lhs11_ = d_67_v34_
                        lhs12_ = default__.safeIndex(407, (d_67_v34_).length(0))
                        d_71_v36_ = rhs11_
                        lhs8_[lhs9_] = rhs12_
                        lhs10_.f2 = rhs13_
                        lhs11_[lhs12_] = rhs14_
                        d_79_v44_ = rhs15_
                    elif source0_.is_DC3:
                        d_80___mcc_h5_ = source0_.cf9
                        d_81___mcc_h6_ = source0_.cf10
                        d_82___mcc_h7_ = source0_.cf11
                        d_83___mcc_h8_ = source0_.cf12
                        d_84_cf12_ = d_83___mcc_h8_
                        d_85_cf11_ = d_82___mcc_h7_
                        d_86_cf10_ = d_81___mcc_h6_
                        d_87_cf9_ = d_80___mcc_h5_
                        r1 = d_87_cf9_
                        d_88_v45_: C0
                        nw11_ = C0()
                        nw11_.ctor__()
                        d_88_v45_ = nw11_
                        d_89_v47_: _dafny.Array
                        def lambda4_(d_90_cf10_, d_91_p0_, d_92_cf9_):
                            def lambda5_(d_93_i4_):
                                def iife18_():
                                    coll6_ = _dafny.Map()
                                    compr_6_: int
                                    for compr_6_ in _dafny.IntegerRange(-475, 301):
                                        d_94_v46_: int = compr_6_
                                        if ((-475) <= (d_94_v46_)) and ((d_94_v46_) < (301)):
                                            coll6_[(d_94_v46_) * (d_91_p0_)] = _dafny.MultiSet([d_90_cf10_])
                                    return _dafny.Map(coll6_)
                                return (_dafny.Map({d_91_p0_: _dafny.MultiSet([d_92_cf9_, d_90_cf10_])}) if not(d_90_cf10_) else iife18_()
                                )

                            return lambda5_

                        init2_ = lambda4_(d_86_cf10_, p0, d_87_cf9_)
                        nw12_ = _dafny.Array(None, 19)
                        for i0_2_ in range(nw12_.length(0)):
                            nw12_[i0_2_] = init2_(i0_2_)
                        d_89_v47_ = nw12_
                        d_95_v49_: _dafny.Set
                        d_95_v49_ = _dafny.Set({d_84_cf12_, d_86_cf10_, d_13_v0_, d_87_cf9_, d_86_cf10_})
                        d_96_v50_: _dafny.Map
                        d_96_v50_ = _dafny.Map({len(d_17_v4_): d_95_v49_})
                        index11_ = default__.safeIndex(345, (d_89_v47_).length(0))
                        def iife19_():
                            coll7_ = _dafny.Map()
                            compr_7_: int
                            for compr_7_ in (d_96_v50_).keys.Elements:
                                d_97_v48_: int = compr_7_
                                if (d_97_v48_) in (d_96_v50_):
                                    coll7_[(d_97_v48_) * (len(_dafny.SeqWithoutIsStrInference([d_86_cf10_])))] = _dafny.MultiSet([d_86_cf10_, not(d_84_cf12_), d_87_cf9_])
                            return _dafny.Map(coll7_)
                        (d_89_v47_)[index11_] = iife19_()
                        
                        d_98_v51_: _dafny.Seq
                        d_98_v51_ = _dafny.SeqWithoutIsStrInference([d_84_cf12_])
                        d_99_v52_: _dafny.MultiSet
                        d_99_v52_ = _dafny.MultiSet([d_85_cf11_])
                        d_100_v53_: _dafny.Array
                        nw13_ = _dafny.Array(None, 11)
                        nw13_[int(0)] = p0
                        nw13_[int(1)] = p0
                        nw13_[int(2)] = d_85_cf11_
                        nw13_[int(3)] = p0
                        nw13_[int(4)] = 221
                        nw13_[int(5)] = d_85_cf11_
                        nw13_[int(6)] = d_85_cf11_
                        nw13_[int(7)] = len(d_17_v4_)
                        nw13_[int(8)] = d_85_cf11_
                        nw13_[int(9)] = d_85_cf11_
                        nw13_[int(10)] = d_85_cf11_
                        d_100_v53_ = nw13_
                        d_101_v54_: _dafny.Array
                        d_101_v54_ = d_100_v53_
                        d_102_v55_: _dafny.Map
                        d_102_v55_ = _dafny.Map({d_101_v54_: d_13_v0_})
                        d_103_v56_: D2
                        d_103_v56_ = D2_DC7(D2_DC7(D2_DC6()))
                        d_104_v57_: _dafny.Map
                        d_104_v57_ = _dafny.Map({default__.fm2(globalState): d_103_v56_})
                        d_105_v58_: _dafny.Array
                        nw14_ = _dafny.Array(None, 23)
                        nw14_[int(0)] = 858
                        nw14_[int(1)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dnsnwbotg"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xvnoa"))))
                        nw14_[int(2)] = (d_85_cf11_ if d_84_cf12_ else d_85_cf11_)
                        nw14_[int(3)] = p0
                        nw14_[int(4)] = (975) + (655)
                        nw14_[int(5)] = p0
                        nw14_[int(6)] = d_85_cf11_
                        nw14_[int(7)] = default__.safeModuloInt(len(d_17_v4_), p0)
                        nw14_[int(8)] = (d_99_v52_).cardinality
                        nw14_[int(9)] = d_85_cf11_
                        nw14_[int(10)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j")))) * (p0)
                        nw14_[int(11)] = len(d_20_v7_)
                        nw14_[int(12)] = p0
                        nw14_[int(13)] = len(d_102_v55_)
                        nw14_[int(14)] = p0
                        nw14_[int(15)] = (d_85_cf11_) * (p0)
                        nw14_[int(16)] = (len(_dafny.SeqWithoutIsStrInference([d_22_v8_ for d_106_i5_ in range(default__.abs(-769))])) if default__.fm2(globalState) else len(d_104_v57_))
                        nw14_[int(17)] = d_85_cf11_
                        nw14_[int(18)] = d_85_cf11_
                        nw14_[int(19)] = default__.safeModuloInt((0) - ((0) - (p0)), len(d_17_v4_))
                        nw14_[int(20)] = d_85_cf11_
                        nw14_[int(21)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjqydcgsg")))
                        nw14_[int(22)] = d_85_cf11_
                        d_105_v58_ = nw14_
                        index12_ = default__.safeIndex(894, (d_105_v58_).length(0))
                        (d_105_v58_)[index12_] = p0
                        d_107_v59_: D3
                        d_107_v59_ = D3_DC9(p0)
                        index13_ = default__.safeIndex(345, (d_89_v47_).length(0))
                        index14_ = default__.safeIndex(894, (d_105_v58_).length(0))
                        rhs16_ = default__.fm11(not (d_86_cf10_) or (d_13_v0_), 749, d_107_v59_, (d_22_v8_) not in (d_17_v4_), globalState)
                        rhs17_ = d_98_v51_
                        rhs18_ = default__.safeModuloInt(d_85_cf11_, d_85_cf11_)
                        lhs13_ = d_89_v47_
                        lhs14_ = default__.safeIndex(345, (d_89_v47_).length(0))
                        lhs15_ = d_105_v58_
                        lhs16_ = default__.safeIndex(894, (d_105_v58_).length(0))
                        lhs13_[lhs14_] = rhs16_
                        d_98_v51_ = rhs17_
                        lhs15_[lhs16_] = rhs18_
                        d_87_cf9_ = d_86_cf10_
                    elif source0_.is_DC1:
                        d_108___mcc_h9_ = source0_.cf3
                        d_109_cf3_ = d_108___mcc_h9_
                        d_15_v2_ = (d_15_v2_).set(default__.fm2(globalState), d_109_cf3_)
                        d_110_v60_: C0
                        nw15_ = C0()
                        nw15_.ctor__()
                        d_110_v60_ = nw15_
                        d_111_v61_: _dafny.Array
                        nw16_ = _dafny.Array(None, 4)
                        nw16_[int(0)] = (d_15_v2_) | (d_15_v2_)
                        nw16_[int(1)] = d_15_v2_
                        nw16_[int(2)] = _dafny.Map({d_13_v0_: d_109_cf3_})
                        nw16_[int(3)] = d_15_v2_
                        d_111_v61_ = nw16_
                        d_111_v61_ = d_111_v61_
                        d_112_v62_: _dafny.Array
                        def lambda6_(d_113_p0_):
                            def lambda7_(d_114_i6_):
                                return (d_114_i6_) + (d_113_p0_)

                            return lambda7_

                        init3_ = lambda6_(p0)
                        nw17_ = _dafny.Array(None, 6)
                        for i0_3_ in range(nw17_.length(0)):
                            nw17_[i0_3_] = init3_(i0_3_)
                        d_112_v62_ = nw17_
                        d_112_v62_ = d_112_v62_
                    elif True:
                        d_115___mcc_h10_ = source0_.cf13
                        d_116_cf13_ = d_115___mcc_h10_
                        (globalState).f2 = p0
                        d_117_v63_: _dafny.Array
                        nw18_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                        d_117_v63_ = nw18_
                        index15_ = default__.safeIndex(561, (d_117_v63_).length(0))
                        (d_117_v63_)[index15_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbrfdtt"))
                        index16_ = default__.safeIndex(561, (d_117_v63_).length(0))
                        (d_117_v63_)[index16_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rkrne"))
                        d_118_v64_: C0
                        nw19_ = C0()
                        nw19_.ctor__()
                        d_118_v64_ = nw19_
                        rhs19_ = d_54_v32_
                        rhs20_ = d_13_v0_
                        d_54_v32_ = rhs19_
                        d_13_v0_ = rhs20_
                    d_22_v8_ = d_22_v8_
                    pass
            pass
        d_119_v65_: _dafny.Array
        def lambda8_(d_120_v4_):
            def lambda9_(d_121_i8_):
                return ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_122_i9_ in range(default__.abs(100))]), d_120_v4_]))) < (_dafny.SeqWithoutIsStrInference([d_120_v4_]))

            return lambda9_

        init4_ = lambda8_(d_17_v4_)
        nw20_ = _dafny.Array(None, 17)
        for i0_4_ in range(nw20_.length(0)):
            nw20_[i0_4_] = init4_(i0_4_)
        d_119_v65_ = nw20_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_119_v65_).length(0)):
            d_123_i7_: int = guard_loop_0_
            if (True) and (((0) <= (d_123_i7_)) and ((d_123_i7_) < ((d_119_v65_).length(0)))):
                (d_119_v65_)[(d_123_i7_)] = (613) < ((len(d_15_v2_)) - (p0))
        d_124_v66_: _dafny.Array
        nw21_ = _dafny.Array(None, 6)
        d_124_v66_ = nw21_
        index17_ = default__.safeIndex(503, (d_124_v66_).length(0))
        (d_124_v66_)[index17_] = d_14_v1_
        index18_ = default__.safeIndex(503, (d_124_v66_).length(0))
        (d_124_v66_)[index18_] = d_14_v1_
        d_125_v67_: _dafny.Array
        nw22_ = _dafny.Array(None, 4)
        nw22_[int(0)] = d_119_v65_
        nw22_[int(1)] = d_119_v65_
        nw22_[int(2)] = d_119_v65_
        nw22_[int(3)] = d_119_v65_
        d_125_v67_ = nw22_
        index19_ = default__.safeIndex(891, (d_125_v67_).length(0))
        (d_125_v67_)[index19_] = d_119_v65_
        d_126_v68_: D3
        d_126_v68_ = D3_DC8(d_119_v65_)
        index20_ = default__.safeIndex(891, (d_125_v67_).length(0))
        (d_125_v67_)[index20_] = (d_126_v68_).cf16
        d_127_v70_: D1
        d_127_v70_ = D1_DC3(d_13_v0_, d_13_v0_, p0, d_13_v0_)
        d_128_v71_: _dafny.MultiSet
        d_128_v71_ = _dafny.MultiSet([d_127_v70_])
        d_129_v72_: _dafny.Seq
        def iife20_():
            coll8_ = _dafny.Map()
            compr_8_: D1
            for compr_8_ in (d_128_v71_).Elements:
                d_130_v69_: D1 = compr_8_
                if (d_130_v69_) in (d_128_v71_):
                    coll8_[d_130_v69_] = d_13_v0_
            return _dafny.Map(coll8_)
        d_129_v72_ = _dafny.SeqWithoutIsStrInference([iife20_()
        , _dafny.Map({d_127_v70_: True})])
        r0 = d_129_v72_
        r1 = not(d_13_v0_)
        r2 = d_13_v0_
        d_131_v73_: _dafny.Set
        d_131_v73_ = _dafny.Set({(d_13_v0_ if d_13_v0_ else d_13_v0_)})
        r3 = d_131_v73_
        return r0, r1, r2, r3

    @staticmethod
    def Main(noArgsParameter__):
        d_132_v0_: _dafny.Array
        nw23_ = _dafny.Array(False, 19)
        d_132_v0_ = nw23_
        d_133_v1_: _dafny.Array
        def lambda10_(d_134_i0_):
            return _dafny.Set({-34})

        init5_ = lambda10_
        nw24_ = _dafny.Array(None, 21)
        for i0_5_ in range(nw24_.length(0)):
            nw24_[i0_5_] = init5_(i0_5_)
        d_133_v1_ = nw24_
        d_135_globalState_: GlobalState
        nw25_ = GlobalState()
        nw25_.ctor__(False, d_132_v0_, 212, False, 3, True, False, 743, True, d_133_v1_, 700, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_136_i1_ in range(default__.abs(199))]))
        d_135_globalState_ = nw25_
        d_137_v2_: bool
        d_137_v2_ = True
        if d_137_v2_:
            d_138_v3_: int
            d_138_v3_ = 249
            d_139_v4_: _dafny.Map
            d_139_v4_ = _dafny.Map({d_138_v3_: d_137_v2_})
            d_140_v5_: _dafny.Seq
            d_140_v5_ = _dafny.SeqWithoutIsStrInference([(default__.fm0(d_135_globalState_)) | (d_139_v4_)])
            d_140_v5_ = (d_140_v5_ if d_137_v2_ else d_140_v5_)
            d_141_v6_: _dafny.Map
            d_141_v6_ = _dafny.Map({True: (0) - (d_138_v3_)})
            d_142_v7_: _dafny.Seq
            d_142_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xbl"))
            d_143_v8_: _dafny.MultiSet
            d_143_v8_ = _dafny.MultiSet([-795, len(d_142_v7_), d_138_v3_, d_138_v3_, d_138_v3_])
            d_144_v9_: D0
            d_144_v9_ = D0_DC0((0) - (len(d_141_v6_)), d_143_v8_, d_138_v3_)
            source1_ = d_144_v9_
            d_145___mcc_h0_ = source1_.cf0
            d_146___mcc_h1_ = source1_.cf1
            d_147___mcc_h2_ = source1_.cf2
            d_148_cf2_ = d_147___mcc_h2_
            d_149_cf1_ = d_146___mcc_h1_
            d_150_cf0_ = d_145___mcc_h0_
            d_151_v10_: str
            d_151_v10_ = _dafny.CodePoint('g')
            d_152_v11_: _dafny.Array
            nw26_ = _dafny.Array(None, 9)
            nw26_[int(0)] = d_151_v10_
            nw26_[int(1)] = d_151_v10_
            nw26_[int(2)] = default__.fm1(d_135_globalState_)
            nw26_[int(3)] = d_151_v10_
            nw26_[int(4)] = d_151_v10_
            nw26_[int(5)] = d_151_v10_
            nw26_[int(6)] = (d_151_v10_ if d_137_v2_ else d_151_v10_)
            nw26_[int(7)] = default__.fm1(d_135_globalState_)
            nw26_[int(8)] = d_151_v10_
            d_152_v11_ = nw26_
            index21_ = default__.safeIndex(126, (d_152_v11_).length(0))
            (d_152_v11_)[index21_] = d_151_v10_
            index22_ = default__.safeIndex(126, (d_152_v11_).length(0))
            (d_152_v11_)[index22_] = d_151_v10_
            d_153_v12_: _dafny.Map
            d_153_v12_ = _dafny.Map({False: False})
            d_154_v13_: _dafny.Seq
            d_154_v13_ = _dafny.SeqWithoutIsStrInference([(0) - (len(d_153_v12_))])
            d_155_v14_: _dafny.Array
            nw27_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_155_v14_ = nw27_
            d_156_v15_: _dafny.Set
            d_156_v15_ = _dafny.Set({d_137_v2_, d_137_v2_})
            d_157_v16_: _dafny.Array
            nw28_ = _dafny.Array(None, 1)
            nw28_[int(0)] = d_156_v15_
            d_157_v16_ = nw28_
            d_158_v17_: _dafny.Map
            d_158_v17_ = _dafny.Map({d_157_v16_: d_155_v14_})
            d_159_v18_: _dafny.Seq
            d_159_v18_ = _dafny.SeqWithoutIsStrInference([d_157_v16_])
            d_160_v19_: D1
            d_160_v19_ = D1_DC1(d_137_v2_)
            rhs21_ = d_154_v13_
            rhs22_ = ((d_158_v17_)[(d_159_v18_)[default__.safeIndex(d_148_cf2_, len(d_159_v18_))]] if ((d_159_v18_)[default__.safeIndex(d_148_cf2_, len(d_159_v18_))]) in (d_158_v17_) else d_155_v14_)
            rhs23_ = d_137_v2_
            rhs24_ = (d_137_v2_ if (d_148_cf2_) == (d_148_cf2_) else (d_160_v19_).cf3)
            d_154_v13_ = rhs21_
            d_155_v14_ = rhs22_
            d_137_v2_ = rhs23_
            d_137_v2_ = rhs24_
            d_138_v3_ = d_148_cf2_
            d_161_v20_: _dafny.Seq
            d_161_v20_ = _dafny.SeqWithoutIsStrInference([d_137_v2_])
            d_162_v21_: _dafny.Map
            d_162_v21_ = _dafny.Map({d_151_v10_: d_137_v2_})
            d_163_v22_: _dafny.Seq
            d_163_v22_ = _dafny.SeqWithoutIsStrInference([(d_161_v20_)[default__.safeIndex(d_138_v3_, len(d_161_v20_))], False, ((d_162_v21_)[(d_152_v11_)[default__.safeIndex(126, (d_152_v11_).length(0))]] if ((d_152_v11_)[default__.safeIndex(126, (d_152_v11_).length(0))]) in (d_162_v21_) else d_137_v2_), not(d_137_v2_), not(d_137_v2_)])
            d_164_v23_: _dafny.Array
            nw29_ = _dafny.Array(None, 21)
            nw29_[int(0)] = d_149_cf1_
            nw29_[int(1)] = d_149_cf1_
            nw29_[int(2)] = d_149_cf1_
            nw29_[int(3)] = d_149_cf1_
            nw29_[int(4)] = _dafny.MultiSet(d_154_v13_)
            nw29_[int(5)] = d_149_cf1_
            nw29_[int(6)] = d_143_v8_
            nw29_[int(7)] = d_143_v8_
            nw29_[int(8)] = _dafny.MultiSet(d_154_v13_)
            nw29_[int(9)] = _dafny.MultiSet([d_148_cf2_, d_150_cf0_])
            nw29_[int(10)] = d_143_v8_
            nw29_[int(11)] = d_143_v8_
            nw29_[int(12)] = d_149_cf1_
            nw29_[int(13)] = d_143_v8_
            nw29_[int(14)] = d_143_v8_
            nw29_[int(15)] = d_143_v8_
            nw29_[int(16)] = _dafny.MultiSet([d_148_cf2_])
            nw29_[int(17)] = d_143_v8_
            nw29_[int(18)] = d_143_v8_
            nw29_[int(19)] = d_149_cf1_
            nw29_[int(20)] = d_143_v8_
            d_164_v23_ = nw29_
            d_165_v24_: _dafny.Map
            d_165_v24_ = _dafny.Map({d_163_v22_: d_164_v23_})
            d_165_v24_ = (d_165_v24_).set(d_161_v20_, d_164_v23_)
            d_139_v4_ = (d_139_v4_).set(len(d_141_v6_), (not(d_137_v2_)) and (d_137_v2_))
            d_166_v25_: _dafny.Array
            nw30_ = _dafny.Array(_dafny.CodePoint('D'), 1)
            d_166_v25_ = nw30_
            d_167_v26_: _dafny.Map
            d_167_v26_ = _dafny.Map({d_166_v25_: (d_137_v2_ if d_137_v2_ else d_137_v2_)})
            d_168_v29_: _dafny.Map
            d_168_v29_ = _dafny.Map({len(d_142_v7_): d_138_v3_})
            d_169_v30_: _dafny.MultiSet
            def iife21_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in (d_143_v8_).Elements:
                    d_170_v27_: int = compr_9_
                    if (d_170_v27_) in (d_143_v8_):
                        coll9_[(d_170_v27_) - (d_138_v3_)] = len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([len((D2_DC5(_dafny.SeqWithoutIsStrInference([d_137_v2_, d_137_v2_]))).cf14)]))}))
                return _dafny.Map(coll9_)
            def iife22_():
                coll10_ = _dafny.Map()
                compr_10_: int
                for compr_10_ in (d_139_v4_).keys.Elements:
                    d_171_v28_: int = compr_10_
                    if (d_171_v28_) in (d_139_v4_):
                        coll10_[default__.safeDivisionInt(d_171_v28_, d_138_v3_)] = d_138_v3_
                return _dafny.Map(coll10_)
            d_169_v30_ = _dafny.MultiSet([iife21_()
            , iife22_()
            , d_168_v29_])
            d_167_v26_ = (d_167_v26_).set(d_166_v25_, (d_169_v30_).isdisjoint(d_169_v30_))
            d_172_v31_: _dafny.Map
            d_172_v31_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([725])) != (_dafny.SeqWithoutIsStrInference([d_138_v3_, d_138_v3_])): (d_168_v29_) | (_dafny.Map({d_138_v3_: d_138_v3_}))})
            d_168_v29_ = ((d_172_v31_)[d_137_v2_] if (d_137_v2_) in (d_172_v31_) else d_168_v29_)
        elif True:
            d_173_v32_: str
            d_173_v32_ = _dafny.CodePoint('i')
            d_173_v32_ = d_173_v32_
            d_137_v2_ = d_137_v2_
            d_174_v33_: int
            d_174_v33_ = 880
            d_175_v34_: _dafny.Map
            d_175_v34_ = _dafny.Map({(d_174_v33_) != (d_174_v33_): 636})
            d_175_v34_ = (d_175_v34_).set(d_137_v2_, d_174_v33_)
            d_173_v32_ = _dafny.CodePoint('n')
            if (_dafny.SeqWithoutIsStrInference([d_173_v32_ for d_176_i2_ in range(default__.abs(535))])) < (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrvkbowy"))):
                d_177_v35_: _dafny.MultiSet
                d_177_v35_ = _dafny.MultiSet([d_173_v32_, d_173_v32_])
                d_178_v36_: _dafny.MultiSet
                d_178_v36_ = _dafny.MultiSet([d_137_v2_, default__.fm2(d_135_globalState_)])
                d_179_v37_: _dafny.Map
                d_179_v37_ = _dafny.Map({d_174_v33_: d_174_v33_})
                d_180_v38_: _dafny.Array
                nw31_ = _dafny.Array(None, 10)
                nw31_[int(0)] = (d_177_v35_).cardinality
                nw31_[int(1)] = d_174_v33_
                nw31_[int(2)] = len(_dafny.SeqWithoutIsStrInference([(0) - (d_174_v33_) for d_181_i3_ in range(default__.abs(903))]))
                nw31_[int(3)] = d_174_v33_
                nw31_[int(4)] = d_174_v33_
                nw31_[int(5)] = d_174_v33_
                nw31_[int(6)] = d_174_v33_
                nw31_[int(7)] = d_174_v33_
                nw31_[int(8)] = (d_178_v36_).cardinality
                nw31_[int(9)] = len(d_179_v37_)
                d_180_v38_ = nw31_
                d_182_v39_: _dafny.Set
                d_182_v39_ = _dafny.Set({d_180_v38_, d_180_v38_})
                rhs25_ = len(d_182_v39_)
                rhs26_ = default__.fm2(d_135_globalState_)
                rhs27_ = default__.fm2(d_135_globalState_)
                rhs28_ = d_137_v2_
                rhs29_ = d_174_v33_
                lhs17_ = d_135_globalState_
                lhs18_ = d_135_globalState_
                lhs17_.f2 = rhs25_
                d_137_v2_ = rhs26_
                d_137_v2_ = rhs27_
                d_137_v2_ = rhs28_
                lhs18_.f2 = rhs29_
                d_183_v40_: _dafny.Array
                nw32_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 27)
                d_183_v40_ = nw32_
                index23_ = default__.safeIndex(731, (d_183_v40_).length(0))
                (d_183_v40_)[index23_] = _dafny.SeqWithoutIsStrInference([d_173_v32_ for d_184_i4_ in range(default__.abs(266))])
                index24_ = default__.safeIndex(731, (d_183_v40_).length(0))
                (d_183_v40_)[index24_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yjcvetkn"))
                d_185_v41_: _dafny.Array
                nw33_ = _dafny.Array(_dafny.Seq({}), 6)
                d_185_v41_ = nw33_
                d_186_v42_: _dafny.Seq
                d_186_v42_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_174_v33_, d_174_v33_])])
                index25_ = default__.safeIndex(1, (d_185_v41_).length(0))
                (d_185_v41_)[index25_] = d_186_v42_
                index26_ = default__.safeIndex(1, (d_185_v41_).length(0))
                (d_185_v41_)[index26_] = ((d_186_v42_) + (d_186_v42_)) + ((d_186_v42_) + (d_186_v42_))
                index27_ = default__.safeIndex(69, (d_180_v38_).length(0))
                (d_180_v38_)[index27_] = d_174_v33_
                index28_ = default__.safeIndex(69, (d_180_v38_).length(0))
                (d_180_v38_)[index28_] = 59
                d_187_v43_: _dafny.Seq
                d_187_v43_ = _dafny.SeqWithoutIsStrInference([-970])
                (d_135_globalState_).f2 = ((d_178_v36_)[(d_187_v43_) < (d_187_v43_)] if ((d_187_v43_) < (d_187_v43_)) in (d_178_v36_) else (0) - (d_174_v33_))
            elif True:
                d_188_v44_: _dafny.MultiSet
                d_188_v44_ = _dafny.MultiSet([not(default__.fm2(d_135_globalState_))])
                (d_135_globalState_).f2 = (d_188_v44_).cardinality
                d_189_v45_: _dafny.Map
                d_189_v45_ = _dafny.Map({d_137_v2_: d_137_v2_})
                d_137_v2_ = (len(d_189_v45_)) == (d_174_v33_)
                d_189_v45_ = (d_189_v45_).set(d_137_v2_, not (d_137_v2_) or (d_137_v2_))
                d_190_v46_: C0
                nw34_ = C0()
                nw34_.ctor__()
                d_190_v46_ = nw34_
                rhs30_ = d_137_v2_
                rhs31_ = (d_174_v33_) * (d_174_v33_)
                rhs32_ = default__.fm3(d_137_v2_, _dafny.MultiSet([False, d_137_v2_, d_137_v2_]), d_137_v2_, d_174_v33_, d_135_globalState_)
                rhs33_ = len(((d_189_v45_) | (d_189_v45_)) | (d_189_v45_))
                lhs19_ = d_135_globalState_
                d_137_v2_ = rhs30_
                lhs19_.f2 = rhs31_
                d_189_v45_ = rhs32_
                d_174_v33_ = rhs33_
        index29_ = default__.safeIndex(495, (d_132_v0_).length(0))
        (d_132_v0_)[index29_] = d_137_v2_
        index30_ = default__.safeIndex(495, (d_132_v0_).length(0))
        (d_132_v0_)[index30_] = d_137_v2_
        d_191_v47_: _dafny.Array
        nw35_ = _dafny.Array(D1.default()(), 2)
        d_191_v47_ = nw35_
        _ingredientsd_0 = []
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_191_v47_).length(0)):
            d_192_i5_: int = guard_loop_1_
            if (True) and (((0) <= (d_192_i5_)) and ((d_192_i5_) < ((d_191_v47_).length(0)))):
                _ingredientsd_0.append((d_191_v47_, int(d_192_i5_), (D1_DC4(D1_DC1((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))])) if d_137_v2_ else D1_DC4(D1_DC1(d_137_v2_)))))
        for _tupd_0 in _ingredientsd_0:
            _tupd_0[0][_tupd_0[1]] = _tupd_0[2]
        if not((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]):
            d_193_v48_: int
            d_193_v48_ = 890
            d_194_v49_: _dafny.Map
            d_194_v49_ = _dafny.Map({d_193_v48_: default__.fm2(d_135_globalState_)})
            d_195_v50_: _dafny.Seq
            d_196_v51_: bool
            d_197_v52_: bool
            d_198_v53_: _dafny.Set
            out0_: _dafny.Seq
            out1_: bool
            out2_: bool
            out3_: _dafny.Set
            out0_, out1_, out2_, out3_ = default__.m0(len(d_194_v49_), d_135_globalState_)
            d_195_v50_ = out0_
            d_196_v51_ = out1_
            d_197_v52_ = out2_
            d_198_v53_ = out3_
            (d_135_globalState_).f2 = d_193_v48_
            d_199_v54_: _dafny.Array
            nw36_ = _dafny.Array(int(0), 15)
            d_199_v54_ = nw36_
            index31_ = default__.safeIndex(238, (d_199_v54_).length(0))
            (d_199_v54_)[index31_] = d_193_v48_
            d_200_v55_: _dafny.Set
            d_200_v55_ = _dafny.Set({d_193_v48_, d_193_v48_})
            index32_ = default__.safeIndex(238, (d_199_v54_).length(0))
            (d_199_v54_)[index32_] = (len(d_200_v55_)) + (d_193_v48_)
            d_196_v51_ = ((_dafny.Set({d_196_v51_, True})) | (d_198_v53_)).ispropersubset((d_198_v53_) - (d_198_v53_))
            d_201_v56_: C0
            nw37_ = C0()
            nw37_.ctor__()
            d_201_v56_ = nw37_
        elif True:
            d_202_v57_: int
            d_202_v57_ = 970
            d_203_v58_: _dafny.MultiSet
            d_203_v58_ = _dafny.MultiSet([d_202_v57_, d_202_v57_])
            d_203_v58_ = (default__.fm4(d_135_globalState_)).intersection(d_203_v58_)
            index33_ = default__.safeIndex(495, (d_132_v0_).length(0))
            (d_132_v0_)[index33_] = (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]
            d_204_v59_: _dafny.Seq
            d_205_v60_: bool
            d_206_v61_: bool
            d_207_v62_: _dafny.Set
            out4_: _dafny.Seq
            out5_: bool
            out6_: bool
            out7_: _dafny.Set
            out4_, out5_, out6_, out7_ = default__.m0(d_202_v57_, d_135_globalState_)
            d_204_v59_ = out4_
            d_205_v60_ = out5_
            d_206_v61_ = out6_
            d_207_v62_ = out7_
            index34_ = default__.safeIndex(495, (d_132_v0_).length(0))
            (d_132_v0_)[index34_] = not(d_137_v2_)
            d_208_v63_: _dafny.Map
            d_208_v63_ = _dafny.Map({d_202_v57_: False})
            d_209_v64_: _dafny.Seq
            d_209_v64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ljigfrq"))
            d_210_v65_: _dafny.Map
            d_210_v65_ = _dafny.Map({d_202_v57_: d_209_v64_})
            if (d_202_v57_) >= (len(((d_208_v63_).set(len(((d_210_v65_)[d_202_v57_] if (d_202_v57_) in (d_210_v65_) else d_209_v64_)), (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))])) | (d_208_v63_))):
                d_205_v60_ = d_137_v2_
                (d_135_globalState_).f2 = 387
                d_211_v66_: _dafny.Array
                nw38_ = _dafny.Array(int(0), 22)
                d_211_v66_ = nw38_
                d_211_v66_ = d_211_v66_
                d_212_v67_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Seq({}), 2)
                d_212_v67_ = nw39_
                d_213_v68_: _dafny.Array
                def lambda11_(d_214_i6_):
                    return _dafny.CodePoint('p')

                init6_ = lambda11_
                nw40_ = _dafny.Array(None, 25)
                for i0_6_ in range(nw40_.length(0)):
                    nw40_[i0_6_] = init6_(i0_6_)
                d_213_v68_ = nw40_
                d_215_v69_: str
                d_215_v69_ = _dafny.CodePoint('j')
                index35_ = default__.safeIndex(260, (d_213_v68_).length(0))
                (d_213_v68_)[index35_] = d_215_v69_
                d_216_v70_: _dafny.Array
                nw41_ = _dafny.Array(D1.default()(), 8)
                d_216_v70_ = nw41_
                index36_ = default__.safeIndex(482, (d_216_v70_).length(0))
                (d_216_v70_)[index36_] = default__.fm5(d_135_globalState_)
                d_217_v71_: C0
                nw42_ = C0()
                nw42_.ctor__()
                d_217_v71_ = nw42_
                d_218_v72_: _dafny.Map
                d_218_v72_ = _dafny.Map({d_217_v71_: d_202_v57_})
                d_219_v73_: _dafny.Seq
                d_219_v73_ = _dafny.SeqWithoutIsStrInference([(d_202_v57_) + (((d_218_v72_)[d_217_v71_] if (d_217_v71_) in (d_218_v72_) else 777))])
                d_220_v74_: D1
                d_220_v74_ = D1_DC3(d_206_v61_, d_206_v61_, len(_dafny.Set({d_205_v60_})), d_137_v2_)
                pat_let_tv12_ = d_202_v57_
                pat_let_tv13_ = d_132_v0_
                pat_let_tv14_ = d_132_v0_
                index37_ = default__.safeIndex(495, (d_132_v0_).length(0))
                index38_ = default__.safeIndex(260, (d_213_v68_).length(0))
                index39_ = default__.safeIndex(482, (d_216_v70_).length(0))
                def iife23_(_pat_let6_0):
                    def iife24_(d_221_dt__update__tmp_h0_):
                        def iife25_(_pat_let7_0):
                            def iife26_(d_222_dt__update_hcf9_h0_):
                                def iife27_(_pat_let8_0):
                                    def iife28_(d_223_dt__update_hcf12_h0_):
                                        return D1_DC3(d_222_dt__update_hcf9_h0_, (d_221_dt__update__tmp_h0_).cf10, (d_221_dt__update__tmp_h0_).cf11, d_223_dt__update_hcf12_h0_)
                                    return iife28_(_pat_let8_0)
                                return iife27_((pat_let_tv14_)[default__.safeIndex(495, (pat_let_tv13_).length(0))])
                            return iife26_(_pat_let7_0)
                        return iife25_((pat_let_tv12_) > (574))
                    return iife24_(_pat_let6_0)
                rhs34_ = (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]
                rhs35_ = d_212_v67_
                rhs36_ = d_215_v69_
                rhs37_ = iife23_(d_220_v74_)
                rhs38_ = _dafny.SeqWithoutIsStrInference([d_202_v57_ for d_224_i7_ in range(default__.abs(-651))])
                lhs20_ = d_132_v0_
                lhs21_ = default__.safeIndex(495, (d_132_v0_).length(0))
                lhs22_ = d_213_v68_
                lhs23_ = default__.safeIndex(260, (d_213_v68_).length(0))
                lhs24_ = d_216_v70_
                lhs25_ = default__.safeIndex(482, (d_216_v70_).length(0))
                lhs20_[lhs21_] = rhs34_
                d_212_v67_ = rhs35_
                lhs22_[lhs23_] = rhs36_
                lhs24_[lhs25_] = rhs37_
                d_219_v73_ = rhs38_
                d_225_v75_: D0
                d_225_v75_ = D0_DC0(d_202_v57_, _dafny.MultiSet([d_202_v57_]), d_202_v57_)
                d_226_v76_: _dafny.Seq
                d_226_v76_ = _dafny.SeqWithoutIsStrInference([d_132_v0_])
                d_227_v77_: D1
                d_227_v77_ = D1_DC2(default__.safeModuloInt(d_202_v57_, d_202_v57_), default__.fm1(d_135_globalState_), ((d_225_v75_).cf0) * (len(d_226_v76_)), d_219_v73_, (d_215_v69_ if (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))] else (d_213_v68_)[default__.safeIndex(260, (d_213_v68_).length(0))]))
                d_227_v77_ = d_227_v77_
            elif True:
                d_228_v78_: _dafny.Map
                d_228_v78_ = _dafny.Map({351: d_202_v57_})
                d_229_v79_: _dafny.MultiSet
                d_229_v79_ = _dafny.MultiSet([not(d_137_v2_), d_206_v61_, (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]])
                d_230_v80_: _dafny.Seq
                d_230_v80_ = _dafny.SeqWithoutIsStrInference([d_137_v2_])
                d_231_v81_: _dafny.Array
                nw43_ = _dafny.Array(int(0), 7)
                d_231_v81_ = nw43_
                d_232_v82_: _dafny.Set
                d_232_v82_ = _dafny.Set({d_231_v81_, d_231_v81_, d_231_v81_})
                d_233_v83_: _dafny.Array
                nw44_ = _dafny.Array(None, 28)
                nw44_[int(0)] = d_202_v57_
                nw44_[int(1)] = (0) - (d_202_v57_)
                nw44_[int(2)] = (d_202_v57_) * ((d_203_v58_).cardinality)
                nw44_[int(3)] = d_202_v57_
                nw44_[int(4)] = d_202_v57_
                nw44_[int(5)] = (len(d_209_v64_)) + (352)
                nw44_[int(6)] = d_202_v57_
                nw44_[int(7)] = d_202_v57_
                nw44_[int(8)] = d_202_v57_
                nw44_[int(9)] = d_202_v57_
                nw44_[int(10)] = d_202_v57_
                nw44_[int(11)] = d_202_v57_
                nw44_[int(12)] = d_202_v57_
                nw44_[int(13)] = d_202_v57_
                nw44_[int(14)] = d_202_v57_
                nw44_[int(15)] = (d_202_v57_) + (((d_203_v58_)[d_202_v57_] if (d_202_v57_) in (d_203_v58_) else d_202_v57_))
                nw44_[int(16)] = default__.safeDivisionInt(d_202_v57_, len(_dafny.SeqWithoutIsStrInference([53 for d_234_i8_ in range(default__.abs(-382))])))
                nw44_[int(17)] = d_202_v57_
                nw44_[int(18)] = d_202_v57_
                nw44_[int(19)] = (d_202_v57_) - (d_202_v57_)
                nw44_[int(20)] = default__.safeDivisionInt(len(d_228_v78_), d_202_v57_)
                nw44_[int(21)] = 94
                nw44_[int(22)] = d_202_v57_
                nw44_[int(23)] = ((d_229_v79_)[d_206_v61_] if (d_206_v61_) in (d_229_v79_) else d_202_v57_)
                nw44_[int(24)] = len((d_230_v80_ if d_205_v60_ else d_230_v80_))
                nw44_[int(25)] = (599) - (d_202_v57_)
                nw44_[int(26)] = ((d_229_v79_)[(d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]] if ((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]) in (d_229_v79_) else d_202_v57_)
                nw44_[int(27)] = len(d_232_v82_)
                d_233_v83_ = nw44_
                index40_ = default__.safeIndex(703, (d_231_v81_).length(0))
                (d_231_v81_)[index40_] = d_202_v57_
                index41_ = default__.safeIndex(495, (d_132_v0_).length(0))
                index42_ = default__.safeIndex(703, (d_231_v81_).length(0))
                rhs39_ = default__.fm2(d_135_globalState_)
                rhs40_ = True
                rhs41_ = 821
                rhs42_ = d_205_v60_
                lhs26_ = d_132_v0_
                lhs27_ = default__.safeIndex(495, (d_132_v0_).length(0))
                lhs28_ = d_231_v81_
                lhs29_ = default__.safeIndex(703, (d_231_v81_).length(0))
                d_206_v61_ = rhs39_
                lhs26_[lhs27_] = rhs40_
                lhs28_[lhs29_] = rhs41_
                d_206_v61_ = rhs42_
                (d_135_globalState_).f2 = (0) - ((d_231_v81_)[default__.safeIndex(703, (d_231_v81_).length(0))])
                d_235_v84_: D1
                d_235_v84_ = D1_DC3(not(not((d_230_v80_)[default__.safeIndex(d_202_v57_, len(d_230_v80_))])), d_206_v61_, (d_231_v81_)[default__.safeIndex(703, (d_231_v81_).length(0))], d_206_v61_)
                d_206_v61_ = (d_235_v84_).cf12
                index43_ = default__.safeIndex(703, (d_231_v81_).length(0))
                (d_231_v81_)[index43_] = (0) - ((d_202_v57_ if d_205_v60_ else (d_231_v81_)[default__.safeIndex(703, (d_231_v81_).length(0))]))
                d_202_v57_ = -825
        d_236_v85_: int
        d_236_v85_ = -845
        hi0_ = default__.safeDivisionInt(d_236_v85_, default__.fm6(d_236_v85_, d_135_globalState_))
        for d_237_i9_ in range(d_236_v85_, hi0_):
            d_238_v86_: _dafny.Seq
            d_238_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fjukyxth"))
            d_239_v87_: _dafny.Map
            d_239_v87_ = _dafny.Map({d_237_i9_: d_238_v86_})
            d_239_v87_ = (d_239_v87_).set((len(_dafny.Set({d_236_v85_, d_236_v85_, len(default__.fm7(d_237_i9_, d_135_globalState_))})) if (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))] else d_236_v85_), d_238_v86_)
            d_240_v88_: C0
            nw45_ = C0()
            nw45_.ctor__()
            d_240_v88_ = nw45_
            d_241_v89_: D1
            d_241_v89_ = D1_DC1(d_137_v2_)
            d_242_v90_: _dafny.Seq
            d_242_v90_ = _dafny.SeqWithoutIsStrInference([(d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))], (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]])
            pat_let_tv15_ = d_137_v2_
            index44_ = default__.safeIndex(495, (d_132_v0_).length(0))
            index45_ = default__.safeIndex(495, (d_132_v0_).length(0))
            def iife29_(_pat_let9_0):
                def iife30_(d_243_dt__update__tmp_h1_):
                    def iife31_(_pat_let10_0):
                        def iife32_(d_244_dt__update_hcf3_h0_):
                            return D1_DC1(d_244_dt__update_hcf3_h0_)
                        return iife32_(_pat_let10_0)
                    return iife31_(pat_let_tv15_)
                return iife30_(_pat_let9_0)
            rhs43_ = False
            rhs44_ = (d_242_v90_)[default__.safeIndex((d_236_v85_ if d_137_v2_ else d_236_v85_), len(d_242_v90_))]
            rhs45_ = iife29_(D1_DC1((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]))
            rhs46_ = (d_137_v2_) == (((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]) and (d_137_v2_))
            rhs47_ = d_240_v88_
            lhs30_ = d_132_v0_
            lhs31_ = default__.safeIndex(495, (d_132_v0_).length(0))
            lhs32_ = d_132_v0_
            lhs33_ = default__.safeIndex(495, (d_132_v0_).length(0))
            d_137_v2_ = rhs43_
            lhs30_[lhs31_] = rhs44_
            d_241_v89_ = rhs45_
            lhs32_[lhs33_] = rhs46_
            d_240_v88_ = rhs47_
            pat_let_tv16_ = d_132_v0_
            pat_let_tv17_ = d_132_v0_
            index46_ = default__.safeIndex(495, (d_132_v0_).length(0))
            def iife33_(_pat_let11_0):
                def iife34_(d_245_dt__update__tmp_h2_):
                    def iife35_(_pat_let12_0):
                        def iife36_(d_246_dt__update_hcf3_h1_):
                            return D1_DC1(d_246_dt__update_hcf3_h1_)
                        return iife36_(_pat_let12_0)
                    return iife35_((pat_let_tv17_)[default__.safeIndex(495, (pat_let_tv16_).length(0))])
                return iife34_(_pat_let11_0)
            (d_132_v0_)[index46_] = (iife33_(d_241_v89_)).cf3
        d_247_v91_: _dafny.Seq
        d_247_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dbygq"))
        d_248_v92_: _dafny.Seq
        d_248_v92_ = _dafny.SeqWithoutIsStrInference([d_236_v85_, d_236_v85_, d_236_v85_, len(d_247_v91_)])
        d_249_v93_: _dafny.Seq
        d_250_v94_: bool
        d_251_v95_: bool
        d_252_v96_: _dafny.Set
        out8_: _dafny.Seq
        out9_: bool
        out10_: bool
        out11_: _dafny.Set
        out8_, out9_, out10_, out11_ = default__.m0(len((d_248_v92_) + (d_248_v92_)), d_135_globalState_)
        d_249_v93_ = out8_
        d_250_v94_ = out9_
        d_251_v95_ = out10_
        d_252_v96_ = out11_
        d_253_v97_: D3
        d_253_v97_ = D3_DC8(d_132_v0_)
        d_254_v98_: _dafny.Map
        d_254_v98_ = _dafny.Map({d_236_v85_: (d_253_v97_).cf16})
        d_254_v98_ = (d_254_v98_).set(d_236_v85_, d_132_v0_)
        d_255_v99_: C0
        nw46_ = C0()
        nw46_.ctor__()
        d_255_v99_ = nw46_
        d_256_v100_: _dafny.Map
        d_256_v100_ = _dafny.Map({(d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]: d_255_v99_})
        d_255_v99_ = ((d_256_v100_)[d_137_v2_] if (d_137_v2_) in (d_256_v100_) else d_255_v99_)
        d_257_v101_: _dafny.MultiSet
        d_257_v101_ = _dafny.MultiSet([d_236_v85_])
        d_258_v102_: _dafny.Seq
        d_259_v103_: bool
        d_260_v104_: bool
        d_261_v105_: _dafny.Set
        out12_: _dafny.Seq
        out13_: bool
        out14_: bool
        out15_: _dafny.Set
        out12_, out13_, out14_, out15_ = default__.m0(((d_257_v101_) - (d_257_v101_)).cardinality, d_135_globalState_)
        d_258_v102_ = out12_
        d_259_v103_ = out13_
        d_260_v104_ = out14_
        d_261_v105_ = out15_
        d_248_v92_ = ((_dafny.SeqWithoutIsStrInference([d_236_v85_])) + (d_248_v92_)) + (d_248_v92_)
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_132_v0_).length(0)):
            d_262_i10_: int = guard_loop_2_
            if (True) and (((0) <= (d_262_i10_)) and ((d_262_i10_) < ((d_132_v0_).length(0)))):
                (d_132_v0_)[(d_262_i10_)] = d_251_v95_
        _ingredientsd_1 = []
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_132_v0_).length(0)):
            d_263_i11_: int = guard_loop_3_
            if (True) and (((0) <= (d_263_i11_)) and ((d_263_i11_) < ((d_132_v0_).length(0)))):
                _ingredientsd_1.append((d_132_v0_, int(d_263_i11_), (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]))
        for _tupd_1 in _ingredientsd_1:
            _tupd_1[0][_tupd_1[1]] = _tupd_1[2]
        hi1_ = d_236_v85_
        for d_264_i12_ in range(d_236_v85_, hi1_):
            d_265_v106_: str
            d_265_v106_ = _dafny.CodePoint('n')
            d_266_v107_: _dafny.Seq
            d_266_v107_ = _dafny.SeqWithoutIsStrInference([d_248_v92_, d_248_v92_])
            d_267_v108_: D1
            d_267_v108_ = D1_DC2(d_236_v85_, d_265_v106_, d_264_i12_, ((d_266_v107_)[default__.safeIndex(len(d_247_v91_), len(d_266_v107_))]) + (d_248_v92_), d_265_v106_)
            d_268_v109_: D3
            d_268_v109_ = D3_DC10(default__.fm2(d_135_globalState_), d_247_v91_, d_259_v103_)
            d_269_v110_: _dafny.Seq
            d_269_v110_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))) + (d_247_v91_), (d_247_v91_) + ((d_268_v109_).cf19)])
            d_270_v111_: _dafny.Seq
            d_270_v111_ = _dafny.SeqWithoutIsStrInference([d_250_v94_])
            d_271_v112_: _dafny.Seq
            d_271_v112_ = _dafny.SeqWithoutIsStrInference([d_270_v111_])
            rhs48_ = (d_269_v110_)[default__.safeIndex(((d_257_v101_)[len((d_271_v112_)[default__.safeIndex(d_236_v85_, len(d_271_v112_))])] if (len((d_271_v112_)[default__.safeIndex(d_236_v85_, len(d_271_v112_))])) in (d_257_v101_) else d_264_i12_), len(d_269_v110_))]
            rhs49_ = d_267_v108_
            rhs50_ = default__.fm2(d_135_globalState_)
            rhs51_ = d_264_i12_
            lhs34_ = d_135_globalState_
            lhs34_.f11 = rhs48_
            d_267_v108_ = rhs49_
            d_259_v103_ = rhs50_
            d_236_v85_ = rhs51_
            d_272_v113_: _dafny.Map
            d_272_v113_ = _dafny.Map({default__.fm6(54, d_135_globalState_): (d_236_v85_) < (148)})
            d_273_v114_: _dafny.Map
            d_273_v114_ = _dafny.Map({d_264_i12_: (0) - (d_236_v85_)})
            d_274_v115_: _dafny.Map
            d_274_v115_ = _dafny.Map({d_132_v0_: (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]})
            d_275_v116_: _dafny.Map
            d_275_v116_ = _dafny.Map({d_132_v0_: (0) - (len(d_274_v115_))})
            d_276_v117_: _dafny.Array
            nw47_ = _dafny.Array(None, 17)
            nw47_[int(0)] = d_236_v85_
            nw47_[int(1)] = d_236_v85_
            nw47_[int(2)] = d_236_v85_
            nw47_[int(3)] = len((d_273_v114_) | (d_273_v114_))
            nw47_[int(4)] = d_264_i12_
            nw47_[int(5)] = d_264_i12_
            nw47_[int(6)] = d_264_i12_
            nw47_[int(7)] = d_236_v85_
            nw47_[int(8)] = ((d_275_v116_)[d_132_v0_] if (d_132_v0_) in (d_275_v116_) else d_264_i12_)
            nw47_[int(9)] = default__.safeDivisionInt(d_264_i12_, d_236_v85_)
            nw47_[int(10)] = default__.safeModuloInt(878, (0) - (d_236_v85_))
            nw47_[int(11)] = d_264_i12_
            nw47_[int(12)] = len(d_247_v91_)
            nw47_[int(13)] = (d_236_v85_) * (len((d_268_v109_).cf19))
            nw47_[int(14)] = d_236_v85_
            nw47_[int(15)] = d_236_v85_
            nw47_[int(16)] = d_264_i12_
            d_276_v117_ = nw47_
            d_277_v118_: _dafny.Map
            d_277_v118_ = _dafny.Map({default__.fm8(d_236_v85_, d_260_v104_, d_137_v2_, d_135_globalState_): d_272_v113_})
            d_278_v120_: _dafny.Seq
            def iife37_():
                coll11_ = _dafny.Map()
                compr_11_: int
                for compr_11_ in (d_257_v101_).Elements:
                    d_279_v119_: int = compr_11_
                    if (d_279_v119_) in (d_257_v101_):
                        coll11_[(d_279_v119_) * (d_236_v85_)] = True
                return _dafny.Map(coll11_)
            d_278_v120_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_264_i12_: d_260_v104_}), ((d_277_v118_)[_dafny.Set({(0) - (d_236_v85_)})] if (_dafny.Set({(0) - (d_236_v85_)})) in (d_277_v118_) else d_272_v113_), d_272_v113_, iife37_()
            , (d_272_v113_) | (_dafny.Map({-579: (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]}))])
            index47_ = default__.safeIndex(495, (d_132_v0_).length(0))
            rhs52_ = (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]
            rhs53_ = (d_278_v120_)[default__.safeIndex(d_264_i12_, len(d_278_v120_))]
            rhs54_ = d_276_v117_
            lhs35_ = d_132_v0_
            lhs36_ = default__.safeIndex(495, (d_132_v0_).length(0))
            lhs35_[lhs36_] = rhs52_
            d_272_v113_ = rhs53_
            d_276_v117_ = rhs54_
            d_280_v121_: _dafny.Array
            nw48_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_280_v121_ = nw48_
            d_281_v122_: _dafny.Set
            d_281_v122_ = d_252_v96_
            d_282_v123_: _dafny.Map
            d_282_v123_ = _dafny.Map({not((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]): d_261_v105_})
            d_283_v124_: _dafny.Array
            nw49_ = _dafny.Array(None, 26)
            nw49_[int(0)] = d_252_v96_
            nw49_[int(1)] = d_252_v96_
            nw49_[int(2)] = d_261_v105_
            nw49_[int(3)] = d_261_v105_
            nw49_[int(4)] = d_252_v96_
            nw49_[int(5)] = _dafny.Set({False, False})
            nw49_[int(6)] = d_252_v96_
            nw49_[int(7)] = d_252_v96_
            nw49_[int(8)] = d_252_v96_
            nw49_[int(9)] = d_252_v96_
            nw49_[int(10)] = d_252_v96_
            nw49_[int(11)] = d_261_v105_
            nw49_[int(12)] = (d_281_v122_)
            nw49_[int(13)] = d_261_v105_
            nw49_[int(14)] = d_261_v105_
            nw49_[int(15)] = d_252_v96_
            nw49_[int(16)] = d_252_v96_
            nw49_[int(17)] = d_252_v96_
            nw49_[int(18)] = d_261_v105_
            nw49_[int(19)] = d_252_v96_
            nw49_[int(20)] = ((d_282_v123_)[True] if (True) in (d_282_v123_) else d_252_v96_)
            nw49_[int(21)] = _dafny.Set({d_251_v95_, d_250_v94_})
            nw49_[int(22)] = d_261_v105_
            nw49_[int(23)] = d_261_v105_
            nw49_[int(24)] = d_261_v105_
            nw49_[int(25)] = d_261_v105_
            d_283_v124_ = nw49_
            index48_ = default__.safeIndex(601, (d_280_v121_).length(0))
            (d_280_v121_)[index48_] = d_283_v124_
            d_284_v125_: _dafny.Map
            d_284_v125_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_265_v106_ for d_285_i13_ in range(default__.abs(683))]): d_283_v124_})
            index49_ = default__.safeIndex(601, (d_280_v121_).length(0))
            (d_280_v121_)[index49_] = ((d_284_v125_)[d_247_v91_] if (d_247_v91_) in (d_284_v125_) else d_283_v124_)
            (d_135_globalState_).f2 = d_264_i12_
        if (745) >= ((0) - (d_236_v85_)):
            d_286_v126_: _dafny.Array
            nw50_ = _dafny.Array(None, 4)
            nw50_[int(0)] = d_248_v92_
            nw50_[int(1)] = d_248_v92_
            nw50_[int(2)] = d_248_v92_
            nw50_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_236_v85_ for d_287_i14_ in range(default__.abs(525))]) if default__.fm2(d_135_globalState_) else d_248_v92_)
            d_286_v126_ = nw50_
            index50_ = default__.safeIndex(508, (d_286_v126_).length(0))
            (d_286_v126_)[index50_] = _dafny.SeqWithoutIsStrInference([d_236_v85_ for d_288_i15_ in range(default__.abs(688))])
            index51_ = default__.safeIndex(508, (d_286_v126_).length(0))
            (d_286_v126_)[index51_] = _dafny.SeqWithoutIsStrInference([(0) - (d_236_v85_), (d_248_v92_)[default__.safeIndex(d_236_v85_, len(d_248_v92_))]])
            d_289_v127_: str
            d_289_v127_ = _dafny.CodePoint('u')
            d_290_v128_: _dafny.Set
            d_290_v128_ = _dafny.Set({d_289_v127_, d_289_v127_})
            d_291_v129_: _dafny.Map
            d_291_v129_ = _dafny.Map({d_236_v85_: d_290_v128_})
            d_292_v130_: D5
            d_292_v130_ = D5_DC13(d_290_v128_)
            d_259_v103_ = (((d_291_v129_)[d_236_v85_] if (d_236_v85_) in (d_291_v129_) else (d_292_v130_).cf23)).isdisjoint((d_290_v128_).intersection(d_290_v128_))
            d_293_v131_: _dafny.Seq
            d_293_v131_ = _dafny.SeqWithoutIsStrInference([d_137_v2_])
            d_294_v132_: _dafny.Array
            def lambda12_(d_295_v96_):
                def lambda13_(d_296_i16_):
                    return d_295_v96_

                return lambda13_

            init7_ = lambda12_(d_252_v96_)
            nw51_ = _dafny.Array(None, 5)
            for i0_7_ in range(nw51_.length(0)):
                nw51_[i0_7_] = init7_(i0_7_)
            d_294_v132_ = nw51_
            d_297_v133_: _dafny.Set
            d_297_v133_ = d_252_v96_
            index52_ = default__.safeIndex(546, (d_294_v132_).length(0))
            (d_294_v132_)[index52_] = d_297_v133_
            index53_ = default__.safeIndex(495, (d_132_v0_).length(0))
            index54_ = default__.safeIndex(546, (d_294_v132_).length(0))
            rhs55_ = (d_293_v131_) + (d_293_v131_)
            rhs56_ = d_137_v2_
            rhs57_ = d_261_v105_
            rhs58_ = len(d_247_v91_)
            lhs37_ = d_132_v0_
            lhs38_ = default__.safeIndex(495, (d_132_v0_).length(0))
            lhs39_ = d_294_v132_
            lhs40_ = default__.safeIndex(546, (d_294_v132_).length(0))
            lhs41_ = d_135_globalState_
            d_293_v131_ = rhs55_
            lhs37_[lhs38_] = rhs56_
            lhs39_[lhs40_] = rhs57_
            lhs41_.f2 = rhs58_
            d_251_v95_ = d_260_v104_
            (d_135_globalState_).f2 = d_236_v85_
        elif True:
            if not(not(d_137_v2_)):
                d_298_v134_: _dafny.MultiSet
                d_298_v134_ = _dafny.MultiSet([d_247_v91_])
                d_299_v135_: _dafny.Seq
                d_300_v136_: bool
                d_301_v137_: bool
                d_302_v138_: _dafny.Set
                out16_: _dafny.Seq
                out17_: bool
                out18_: bool
                out19_: _dafny.Set
                out16_, out17_, out18_, out19_ = default__.m0((d_298_v134_).cardinality, d_135_globalState_)
                d_299_v135_ = out16_
                d_300_v136_ = out17_
                d_301_v137_ = out18_
                d_302_v138_ = out19_
                d_303_v139_: _dafny.Array
                def lambda14_(d_304_v85_):
                    def lambda15_(d_305_i17_):
                        return (d_305_i17_) + (len(_dafny.Set({948, d_304_v85_, d_304_v85_})))

                    return lambda15_

                init8_ = lambda14_(d_236_v85_)
                nw52_ = _dafny.Array(None, 13)
                for i0_8_ in range(nw52_.length(0)):
                    nw52_[i0_8_] = init8_(i0_8_)
                d_303_v139_ = nw52_
                def lambda16_(d_306_i18_):
                    return (d_306_i18_) * (336)

                init9_ = lambda16_
                nw53_ = _dafny.Array(None, 9)
                for i0_9_ in range(nw53_.length(0)):
                    nw53_[i0_9_] = init9_(i0_9_)
                d_303_v139_ = nw53_
                d_307_v140_: str
                d_307_v140_ = _dafny.CodePoint('h')
                d_307_v140_ = d_307_v140_
                d_137_v2_ = d_260_v104_
                index55_ = default__.safeIndex(495, (d_132_v0_).length(0))
                (d_132_v0_)[index55_] = d_250_v94_
            elif True:
                d_308_v141_: _dafny.Seq
                d_309_v142_: bool
                d_310_v143_: bool
                d_311_v144_: _dafny.Set
                out20_: _dafny.Seq
                out21_: bool
                out22_: bool
                out23_: _dafny.Set
                out20_, out21_, out22_, out23_ = default__.m0(d_236_v85_, d_135_globalState_)
                d_308_v141_ = out20_
                d_309_v142_ = out21_
                d_310_v143_ = out22_
                d_311_v144_ = out23_
                d_255_v99_ = d_255_v99_
                d_255_v99_ = (d_255_v99_ if d_260_v104_ else d_255_v99_)
                d_312_v145_: _dafny.Map
                d_312_v145_ = _dafny.Map({d_236_v85_: (-784) + (d_236_v85_)})
                d_312_v145_ = (d_312_v145_).set(d_236_v85_, d_236_v85_)
                d_313_v146_: _dafny.Set
                d_313_v146_ = _dafny.Set({d_257_v101_})
                d_314_v147_: _dafny.Set
                d_314_v147_ = _dafny.Set({d_255_v99_, d_255_v99_})
                d_315_v148_: _dafny.Map
                d_315_v148_ = _dafny.Map({d_260_v104_: d_137_v2_})
                d_316_v149_: _dafny.Array
                nw54_ = _dafny.Array(None, 20)
                nw54_[int(0)] = d_236_v85_
                nw54_[int(1)] = d_236_v85_
                nw54_[int(2)] = d_236_v85_
                nw54_[int(3)] = d_236_v85_
                nw54_[int(4)] = d_236_v85_
                nw54_[int(5)] = default__.safeDivisionInt(d_236_v85_, default__.fm6(d_236_v85_, d_135_globalState_))
                nw54_[int(6)] = d_236_v85_
                nw54_[int(7)] = d_236_v85_
                nw54_[int(8)] = default__.fm6(d_236_v85_, d_135_globalState_)
                nw54_[int(9)] = d_236_v85_
                nw54_[int(10)] = ((0) - (len(d_313_v146_))) + (d_236_v85_)
                nw54_[int(11)] = default__.fm6(d_236_v85_, d_135_globalState_)
                nw54_[int(12)] = default__.safeModuloInt(d_236_v85_, default__.fm6(len(d_247_v91_), d_135_globalState_))
                nw54_[int(13)] = len(d_314_v147_)
                nw54_[int(14)] = d_236_v85_
                nw54_[int(15)] = d_236_v85_
                nw54_[int(16)] = (0) - (d_236_v85_)
                nw54_[int(17)] = -986
                nw54_[int(18)] = len(d_315_v148_)
                nw54_[int(19)] = default__.safeModuloInt(d_236_v85_, d_236_v85_)
                d_316_v149_ = nw54_
                d_317_v150_: _dafny.Seq
                d_317_v150_ = _dafny.SeqWithoutIsStrInference([d_316_v149_, d_316_v149_, d_316_v149_, d_316_v149_, d_316_v149_])
                d_316_v149_ = (d_317_v150_)[default__.safeIndex((d_236_v85_) * (d_236_v85_), len(d_317_v150_))]
            d_318_v151_: _dafny.Map
            d_318_v151_ = _dafny.Map({d_236_v85_: d_259_v103_})
            d_319_v152_: _dafny.MultiSet
            d_319_v152_ = _dafny.MultiSet([True, d_260_v104_])
            d_320_v153_: _dafny.Map
            d_320_v153_ = _dafny.Map({d_255_v99_: d_319_v152_})
            rhs59_ = not(((d_318_v151_)[default__.safeModuloInt(d_236_v85_, (((d_320_v153_)[d_255_v99_] if (d_255_v99_) in (d_320_v153_) else _dafny.MultiSet([(d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]]))).cardinality)] if (default__.safeModuloInt(d_236_v85_, (((d_320_v153_)[d_255_v99_] if (d_255_v99_) in (d_320_v153_) else _dafny.MultiSet([(d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]]))).cardinality)) in (d_318_v151_) else d_137_v2_))
            rhs60_ = default__.safeModuloInt(d_236_v85_, d_236_v85_)
            lhs42_ = d_135_globalState_
            d_260_v104_ = rhs59_
            lhs42_.f2 = rhs60_
            d_236_v85_ = d_236_v85_
            (d_135_globalState_).f2 = len(d_247_v91_)
            (d_135_globalState_).f2 = ((0) - ((0) - ((0) - (d_236_v85_))) if d_250_v94_ else (default__.fm6(d_236_v85_, d_135_globalState_) if d_260_v104_ else d_236_v85_))
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_132_v0_).length(0)):
            d_321_i19_: int = guard_loop_4_
            if (True) and (((0) <= (d_321_i19_)) and ((d_321_i19_) < ((d_132_v0_).length(0)))):
                (d_132_v0_)[(d_321_i19_)] = d_137_v2_
        d_322_i20_: int
        d_322_i20_ = 0
        with _dafny.label("1"):
            while d_259_v103_:
                with _dafny.c_label("1"):
                    if (d_322_i20_) >= (100):
                        raise _dafny.Break("1")
                    d_322_i20_ = (d_322_i20_) + (1)
                    d_323_v154_: _dafny.Array
                    nw55_ = _dafny.Array(None, 5)
                    nw55_[int(0)] = d_247_v91_
                    nw55_[int(1)] = d_247_v91_
                    nw55_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cugllwi"))
                    nw55_[int(3)] = d_247_v91_
                    nw55_[int(4)] = (d_247_v91_) + (d_247_v91_)
                    d_323_v154_ = nw55_
                    index56_ = default__.safeIndex(240, (d_323_v154_).length(0))
                    (d_323_v154_)[index56_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_324_i21_ in range(default__.abs(38))])
                    index57_ = default__.safeIndex(240, (d_323_v154_).length(0))
                    (d_323_v154_)[index57_] = _dafny.SeqWithoutIsStrInference([(_dafny.CodePoint('m') if d_260_v104_ else _dafny.CodePoint('u')) for d_325_i22_ in range(default__.abs(-999))])
                    if (len(d_248_v92_)) >= (911):
                        d_326_v155_: _dafny.Seq
                        d_327_v156_: bool
                        d_328_v157_: bool
                        d_329_v158_: _dafny.Set
                        out24_: _dafny.Seq
                        out25_: bool
                        out26_: bool
                        out27_: _dafny.Set
                        out24_, out25_, out26_, out27_ = default__.m0(default__.fm6(d_236_v85_, d_135_globalState_), d_135_globalState_)
                        d_326_v155_ = out24_
                        d_327_v156_ = out25_
                        d_328_v157_ = out26_
                        d_329_v158_ = out27_
                        d_330_v159_: _dafny.Map
                        d_330_v159_ = _dafny.Map({(0) - ((d_236_v85_ if d_251_v95_ else d_236_v85_)): d_327_v156_})
                        d_330_v159_ = (d_330_v159_).set((0) - ((d_236_v85_) - (d_236_v85_)), not ((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]) or (d_251_v95_))
                        (d_135_globalState_).f2 = d_236_v85_
                        d_331_v160_: _dafny.MultiSet
                        d_331_v160_ = _dafny.MultiSet([(d_323_v154_)[default__.safeIndex(240, (d_323_v154_).length(0))]])
                        d_332_v161_: str
                        d_332_v161_ = _dafny.CodePoint('q')
                        rhs61_ = _dafny.MultiSet([d_247_v91_, d_247_v91_])
                        rhs62_ = d_332_v161_
                        d_331_v160_ = rhs61_
                        d_332_v161_ = rhs62_
                        d_333_v162_: _dafny.Array
                        nw56_ = _dafny.Array(int(0), 24)
                        d_333_v162_ = nw56_
                        d_334_v163_: _dafny.Set
                        d_334_v163_ = _dafny.Set({d_333_v162_})
                        d_335_v164_: _dafny.Map
                        d_335_v164_ = _dafny.Map({len(d_334_v163_): d_236_v85_})
                        (d_135_globalState_).f2 = default__.fm6((d_236_v85_) - (len(d_335_v164_)), d_135_globalState_)
                    elif True:
                        d_336_v165_: _dafny.Seq
                        d_337_v166_: bool
                        d_338_v167_: bool
                        d_339_v168_: _dafny.Set
                        out28_: _dafny.Seq
                        out29_: bool
                        out30_: bool
                        out31_: _dafny.Set
                        out28_, out29_, out30_, out31_ = default__.m0(d_236_v85_, d_135_globalState_)
                        d_336_v165_ = out28_
                        d_337_v166_ = out29_
                        d_338_v167_ = out30_
                        d_339_v168_ = out31_
                        d_337_v166_ = d_260_v104_
                        (d_135_globalState_).f2 = d_236_v85_
                        d_340_v169_: _dafny.Array
                        nw57_ = _dafny.Array(int(0), 20)
                        d_340_v169_ = nw57_
                        d_341_v170_: _dafny.Map
                        d_341_v170_ = _dafny.Map({not((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]): d_340_v169_})
                        d_342_v171_: D2
                        d_342_v171_ = D2_DC5(_dafny.SeqWithoutIsStrInference([d_137_v2_, d_137_v2_, (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]]))
                        d_343_v172_: D2
                        d_343_v172_ = D2_DC7(d_342_v171_)
                        d_344_v173_: _dafny.Seq
                        d_344_v173_ = _dafny.SeqWithoutIsStrInference([d_343_v172_, D2_DC7(D2_DC6())])
                        d_345_v174_: _dafny.MultiSet
                        d_345_v174_ = _dafny.MultiSet([d_344_v173_])
                        d_346_v175_: _dafny.Map
                        d_346_v175_ = _dafny.Map({_dafny.CodePoint('s'): d_236_v85_})
                        d_347_v176_: _dafny.Map
                        d_347_v176_ = _dafny.Map({d_132_v0_: d_337_v166_})
                        d_348_v177_: _dafny.Array
                        nw58_ = _dafny.Array(None, 21)
                        nw58_[int(0)] = (d_248_v92_)[default__.safeIndex(len(d_341_v170_), len(d_248_v92_))]
                        nw58_[int(1)] = d_236_v85_
                        nw58_[int(2)] = d_236_v85_
                        nw58_[int(3)] = (d_345_v174_).cardinality
                        nw58_[int(4)] = len(d_346_v175_)
                        nw58_[int(5)] = d_236_v85_
                        nw58_[int(6)] = (0) - (d_236_v85_)
                        nw58_[int(7)] = d_236_v85_
                        nw58_[int(8)] = d_236_v85_
                        nw58_[int(9)] = d_236_v85_
                        nw58_[int(10)] = d_236_v85_
                        nw58_[int(11)] = d_236_v85_
                        nw58_[int(12)] = -668
                        nw58_[int(13)] = d_236_v85_
                        nw58_[int(14)] = d_236_v85_
                        nw58_[int(15)] = d_236_v85_
                        nw58_[int(16)] = d_236_v85_
                        nw58_[int(17)] = d_236_v85_
                        nw58_[int(18)] = len(_dafny.Map({d_338_v167_: 288}))
                        nw58_[int(19)] = len(d_347_v176_)
                        nw58_[int(20)] = d_236_v85_
                        d_348_v177_ = nw58_
                        d_349_v178_: _dafny.Map
                        d_349_v178_ = _dafny.Map({-755: (d_340_v169_)})
                        d_350_v179_: _dafny.Array
                        nw59_ = _dafny.Array(None, 15)
                        nw59_[int(0)] = d_348_v177_
                        nw59_[int(1)] = d_348_v177_
                        nw59_[int(2)] = ((d_349_v178_)[(0) - (len(_dafny.Set({d_236_v85_})))] if ((0) - (len(_dafny.Set({d_236_v85_})))) in (d_349_v178_) else d_340_v169_)
                        nw59_[int(3)] = d_348_v177_
                        nw59_[int(4)] = d_340_v169_
                        nw59_[int(5)] = d_348_v177_
                        nw59_[int(6)] = d_348_v177_
                        nw59_[int(7)] = d_348_v177_
                        nw59_[int(8)] = d_348_v177_
                        nw59_[int(9)] = d_348_v177_
                        nw59_[int(10)] = d_348_v177_
                        nw59_[int(11)] = d_340_v169_
                        nw59_[int(12)] = d_340_v169_
                        nw59_[int(13)] = d_340_v169_
                        nw59_[int(14)] = d_340_v169_
                        d_350_v179_ = nw59_
                        index58_ = default__.safeIndex(514, (d_350_v179_).length(0))
                        (d_350_v179_)[index58_] = d_340_v169_
                        index59_ = default__.safeIndex(514, (d_350_v179_).length(0))
                        index60_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        index61_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        rhs63_ = d_250_v94_
                        rhs64_ = d_340_v169_
                        rhs65_ = (d_337_v166_ if (d_236_v85_) > (d_236_v85_) else d_251_v95_)
                        rhs66_ = d_137_v2_
                        lhs43_ = d_350_v179_
                        lhs44_ = default__.safeIndex(514, (d_350_v179_).length(0))
                        lhs45_ = d_132_v0_
                        lhs46_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        lhs47_ = d_132_v0_
                        lhs48_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        d_250_v94_ = rhs63_
                        lhs43_[lhs44_] = rhs64_
                        lhs45_[lhs46_] = rhs65_
                        lhs47_[lhs48_] = rhs66_
                        index62_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        (d_132_v0_)[index62_] = d_259_v103_
                    if d_259_v103_:
                        d_351_v180_: _dafny.MultiSet
                        d_351_v180_ = _dafny.MultiSet([d_250_v94_, not(d_259_v103_), False, d_260_v104_])
                        d_352_v181_: _dafny.Seq
                        d_352_v181_ = _dafny.SeqWithoutIsStrInference([d_351_v180_, d_351_v180_, d_351_v180_, d_351_v180_])
                        d_352_v181_ = (d_352_v181_) + (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_251_v95_, d_260_v104_]) for d_353_i23_ in range(default__.abs(729))]))
                        d_351_v180_ = (d_352_v181_)[default__.safeIndex(d_236_v85_, len(d_352_v181_))]
                        index63_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        (d_132_v0_)[index63_] = not (d_250_v94_) or (not((d_236_v85_) >= ((0) - (d_236_v85_))))
                        index64_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        (d_132_v0_)[index64_] = False
                        d_251_v95_ = d_260_v104_
                    elif True:
                        d_354_v182_: _dafny.Seq
                        d_354_v182_ = _dafny.SeqWithoutIsStrInference([d_261_v105_, d_252_v96_, d_261_v105_])
                        d_355_v183_: _dafny.Set
                        d_355_v183_ = default__.fm9(d_257_v101_, d_260_v104_, d_137_v2_, d_135_globalState_)
                        d_356_v184_: str
                        d_356_v184_ = _dafny.CodePoint('w')
                        d_357_v185_: _dafny.Map
                        d_357_v185_ = _dafny.Map({d_356_v184_: _dafny.Set({d_259_v103_, d_250_v94_, d_259_v103_, d_137_v2_, d_251_v95_})})
                        d_358_v186_: _dafny.Map
                        d_358_v186_ = _dafny.Map({d_356_v184_: d_356_v184_})
                        d_359_v187_: _dafny.Array
                        nw60_ = _dafny.Array(None, 22)
                        nw60_[int(0)] = (d_261_v105_) | (d_252_v96_)
                        nw60_[int(1)] = (_dafny.Set({False, d_259_v103_})) | (d_261_v105_)
                        nw60_[int(2)] = default__.fm9(_dafny.MultiSet([d_236_v85_, d_236_v85_]), True, d_260_v104_, d_135_globalState_)
                        nw60_[int(3)] = _dafny.Set({False})
                        nw60_[int(4)] = (d_354_v182_)[default__.safeIndex(default__.fm6(d_236_v85_, d_135_globalState_), len(d_354_v182_))]
                        nw60_[int(5)] = (_dafny.Set({d_259_v103_, d_260_v104_})).intersection(_dafny.Set({True, d_250_v94_, d_137_v2_, d_260_v104_}))
                        nw60_[int(6)] = d_261_v105_
                        nw60_[int(7)] = (d_355_v183_)
                        nw60_[int(8)] = (d_252_v96_).intersection(_dafny.Set({d_137_v2_, d_260_v104_}))
                        nw60_[int(9)] = d_252_v96_
                        nw60_[int(10)] = d_252_v96_
                        nw60_[int(11)] = d_252_v96_
                        nw60_[int(12)] = ((d_357_v185_)[((d_358_v186_)[d_356_v184_] if (d_356_v184_) in (d_358_v186_) else d_356_v184_)] if (((d_358_v186_)[d_356_v184_] if (d_356_v184_) in (d_358_v186_) else d_356_v184_)) in (d_357_v185_) else d_252_v96_)
                        nw60_[int(13)] = d_252_v96_
                        nw60_[int(14)] = d_252_v96_
                        nw60_[int(15)] = d_261_v105_
                        nw60_[int(16)] = d_252_v96_
                        nw60_[int(17)] = (d_252_v96_ if d_250_v94_ else (d_355_v183_))
                        nw60_[int(18)] = default__.fm9(d_257_v101_, not(d_260_v104_), d_250_v94_, d_135_globalState_)
                        nw60_[int(19)] = d_252_v96_
                        nw60_[int(20)] = d_252_v96_
                        nw60_[int(21)] = d_261_v105_
                        d_359_v187_ = nw60_
                        index65_ = default__.safeIndex(919, (d_359_v187_).length(0))
                        (d_359_v187_)[index65_] = d_261_v105_
                        d_360_v188_: _dafny.Map
                        d_360_v188_ = _dafny.Map({d_356_v184_: d_257_v101_})
                        index66_ = default__.safeIndex(919, (d_359_v187_).length(0))
                        (d_359_v187_)[index66_] = (d_252_v96_ if (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))] else default__.fm9(((d_360_v188_)[d_356_v184_] if (d_356_v184_) in (d_360_v188_) else d_257_v101_), (d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))], d_251_v95_, d_135_globalState_))
                        index67_ = default__.safeIndex(495, (d_132_v0_).length(0))
                        (d_132_v0_)[index67_] = False
                        d_361_v189_: _dafny.Seq
                        d_361_v189_ = _dafny.SeqWithoutIsStrInference([(d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))]])
                        (d_135_globalState_).f2 = default__.safeDivisionInt((d_236_v85_) - (254), (d_236_v85_) + (((d_257_v101_)[d_236_v85_] if (d_236_v85_) in (d_257_v101_) else len(d_361_v189_))))
                        d_137_v2_ = (d_250_v94_) or ((d_132_v0_)[default__.safeIndex(495, (d_132_v0_).length(0))])
                        (d_135_globalState_).f2 = (d_236_v85_) * (default__.fm6(d_236_v85_, d_135_globalState_))
                    d_362_v190_: _dafny.Array
                    nw61_ = _dafny.Array(None, 3)
                    nw61_[int(0)] = d_236_v85_
                    nw61_[int(1)] = (-198) * (d_236_v85_)
                    nw61_[int(2)] = default__.fm6(len(d_247_v91_), d_135_globalState_)
                    d_362_v190_ = nw61_
                    d_363_v191_: _dafny.Map
                    d_363_v191_ = _dafny.Map({d_137_v2_: d_236_v85_})
                    d_364_v192_: _dafny.Map
                    d_364_v192_ = _dafny.Map({d_363_v191_: d_260_v104_})
                    index68_ = default__.safeIndex(157, (d_362_v190_).length(0))
                    (d_362_v190_)[index68_] = (d_236_v85_) * ((0) - (len(d_364_v192_)))
                    d_365_v193_: _dafny.Map
                    d_365_v193_ = _dafny.Map({196: d_236_v85_})
                    d_366_v194_: str
                    d_366_v194_ = _dafny.CodePoint('t')
                    index69_ = default__.safeIndex(157, (d_362_v190_).length(0))
                    (d_362_v190_)[index69_] = ((default__.fm10(d_365_v193_, d_259_v103_, d_260_v104_, False, d_135_globalState_)).cf4) - (default__.safeDivisionInt(len((_dafny.Map({d_366_v194_: d_255_v99_})).set(d_366_v194_, d_255_v99_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwaa")))))
                    pass
            pass
        _dafny.print(_dafny.string_of((d_132_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_132_v0_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[0]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[1]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[2]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[3]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[4]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[5]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[6]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[7]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[8]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[9]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[10]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[11]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[12]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[13]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[14]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[15]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[16]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[17]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[18]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[19]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_133_v1_)[20]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_135_globalState_).f1)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_135_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[0]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[1]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[2]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[3]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[4]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[5]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[6]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[7]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[8]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[9]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[10]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[11]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[12]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[13]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[14]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[15]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[16]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[17]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[18]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[19]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_135_globalState_).f9)[20]) == (_dafny.Set({-34}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_globalState_.f11) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_137_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_191_v47_)[0]).cf13).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_191_v47_)[1]).cf13).cf3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_236_v85_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_247_v91_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_248_v92_) == (_dafny.SeqWithoutIsStrInference([-845, -845, -845, -845, 5, -845, -845, -845, 5]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_249_v93_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({D1_DC3(False, False, 8, False): False}), _dafny.Map({D1_DC3(False, False, 8, False): True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_250_v94_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_251_v95_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v96_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_253_v97_).cf16)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_254_v98_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_256_v100_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_257_v101_) == (_dafny.MultiSet([-845]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_258_v102_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({D1_DC3(False, False, 0, False): False}), _dafny.Map({D1_DC3(False, False, 0, False): True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_259_v103_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_260_v104_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_261_v105_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_322_i20_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC0(int(0), _dafny.MultiSet({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any), ('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)}, {_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0 and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC2(int(0), _dafny.CodePoint('D'), int(0), _dafny.Seq({}), _dafny.CodePoint('D'))
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

class D1_DC2(D1, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any), ('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7 and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC4(D1, NamedTuple('DC4', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC6()
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC5(D2, NamedTuple('DC5', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0))
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
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf18)}, {self.cf19.VerbatimString(True)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC12(D4, NamedTuple('DC12', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(False, False, None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC14(D5, NamedTuple('DC14', [('cf24', Any), ('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC15(D6, NamedTuple('DC15', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC16(D7, NamedTuple('DC16', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D8_DC17)

class D8_DC17(D8, NamedTuple('DC17', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC17({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC17) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self.f11: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f0: bool = False
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f3: bool = False
        self._f4: int = int(0)
        self._f5: bool = False
        self._f6: bool = False
        self._f7: int = int(0)
        self._f8: bool = False
        self._f9: _dafny.Array = _dafny.Array(None, 0)
        self._f10: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self)._f10 = f10
        (self).f11 = f11

    @property
    def f0(self):
        return self._f0
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
    def f8(self):
        return self._f8
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

