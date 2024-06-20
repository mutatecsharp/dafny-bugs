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
        return (default__.safeModuloInt(763, 294)) > (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_0_i0_ in range(default__.abs(934))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "goxbqw")))))

    @staticmethod
    def fm1(p0, p1, p2, globalState):
        if (True) == (True):
            return len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1_i0_ in range(default__.abs(882))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ov"))))
        elif True:
            return len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f'), _dafny.CodePoint('v'), _dafny.CodePoint('d'), _dafny.CodePoint('y')])) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k'), _dafny.CodePoint('p'), _dafny.CodePoint('w'), _dafny.CodePoint('p')])))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.Set({len(_dafny.Map({len(_dafny.Map({710: len(_dafny.Map({816: _dafny.CodePoint('l')}))})): True}))}))) < (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irm")))), (D9_DC25(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ap")), False, 214)).cf60, False])

    @staticmethod
    def fm7(globalState):
        if not((len(_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sfmw"))}))) >= (306)):
            return _dafny.SeqWithoutIsStrInference([True, False])
        elif True:
            return _dafny.SeqWithoutIsStrInference([True, not(not(True))])

    @staticmethod
    def fm9(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        if (True) or (False):
            return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            return _dafny.SeqWithoutIsStrInference([True])

    @staticmethod
    def fm13(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(-834, -776):
                d_2_v0_: int = compr_0_
                if ((-834) <= (d_2_v0_)) and ((d_2_v0_) < (-776)):
                    coll0_[default__.safeModuloInt(d_2_v0_, (_dafny.MultiSet([True])).cardinality)] = 54
            return _dafny.Map(coll0_)
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.Set({-683})).Elements:
                d_3_v1_: int = compr_1_
                if (d_3_v1_) in (_dafny.Set({-683})):
                    coll1_ = coll1_.union(_dafny.Set([default__.safeDivisionInt(d_3_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqnecwgoe"))))]))
            return _dafny.Set(coll1_)
        return ((_dafny.MultiSet([len(iife0_()
        ), len(_dafny.Map({False: (0) - (len(_dafny.Map({False: len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False]))}))})))})), len(_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): False})), len(_dafny.Set({not((D8_DC20(True, False, False)).cf48)}))])).intersection(_dafny.MultiSet([len(_dafny.Set({False, True})), 0, (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(D7_DC18(len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): 552})), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dvkfi")), True, not(False), iife1_()
)).cf43: _dafny.CodePoint('r')}))]))), len(_dafny.SeqWithoutIsStrInference([951, len(_dafny.Map({-863: True})), 159, len(_dafny.Map({not(True): True})), len(_dafny.SeqWithoutIsStrInference([False]))]))]))) | ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-677]))).intersection(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), 861])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "od"))), len(_dafny.SeqWithoutIsStrInference([1]))])))

    @staticmethod
    def fm14(p0, p1, globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return D1_DC2((D9_DC25(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klphfov")), not(True), (0) - (len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arkhat"))): False}))))).cf60, True, False, (_dafny.CodePoint('k') if True else _dafny.CodePoint('g')), _dafny.CodePoint('k'))

    @staticmethod
    def fm16(p0, globalState):
        return (_dafny.Map({True: True})) | (_dafny.Map({not(not(False)): True}))

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({D1_DC3(D1_DC1(_dafny.Set({True})))}) for d_4_i0_ in range(default__.abs(172))]))

    @staticmethod
    def fm18(p0, p1, globalState):
        return ((_dafny.Set({False, False, not(True), True, False})) | (_dafny.Set({True, not(False)}))) | (_dafny.Set({False, True}))

    @staticmethod
    def fm19(p0, globalState):
        return _dafny.MultiSet([(D10_DC27(_dafny.Map({True: 143}))).cf63, (_dafny.Map({False: 176})) | (_dafny.Map({True: 694}))])

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(514, -236):
                d_5_v0_: int = compr_2_
                if ((514) <= (d_5_v0_)) and ((d_5_v0_) < (-236)):
                    def iife3_():
                        coll3_ = _dafny.Map()
                        compr_3_: int
                        for compr_3_ in _dafny.IntegerRange(740, 303):
                            d_6_v1_: int = compr_3_
                            if ((740) <= (d_6_v1_)) and ((d_6_v1_) < (303)):
                                coll3_[(d_6_v1_) + ((_dafny.MultiSet([278])).cardinality)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wag")))
                        return _dafny.Map(coll3_)
                    coll2_ = coll2_.union(_dafny.Set([(d_5_v0_) + (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): len(iife3_()
)})) for d_7_i0_ in range(default__.abs(-183))])))]))
            return _dafny.Set(coll2_)
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(-46, 53):
                d_8_v2_: int = compr_4_
                if ((-46) <= (d_8_v2_)) and ((d_8_v2_) < (53)):
                    coll4_ = coll4_.union(_dafny.Set([(d_8_v2_) * (370)]))
            return _dafny.Set(coll4_)
        return (iife2_()
        ) - ((iife4_()
        ) | (_dafny.Set({-604, (0) - (len(_dafny.Set({(D1_DC2(not(True), True, False, _dafny.CodePoint('g'), _dafny.CodePoint('i'))).cf2}))), 806})))

    @staticmethod
    def fm21(p0, p1, p2, globalState):
        return (D15_DC36(_dafny.Map({_dafny.CodePoint('o'): 123}))).cf78

    @staticmethod
    def fm22(p0, p1, globalState):
        return _dafny.MultiSet([(492) <= (754)])

    @staticmethod
    def fm23(p0, globalState):
        return _dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oot"))): False}))})

    @staticmethod
    def fm24(globalState):
        return _dafny.SeqWithoutIsStrInference([120, (0) - (default__.safeModuloInt(len(_dafny.Set({len(_dafny.Map({False: False}))})), 712)), -813, (941 if not(True) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ldpqylslw"))))])

    @staticmethod
    def fm25(p0, p1, p2, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Seq
            for compr_5_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aua")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_9_i0_ in range(default__.abs(66))])})).Elements:
                d_10_v0_: _dafny.Seq = compr_5_
                if (d_10_v0_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aua")), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_9_i0_ in range(default__.abs(66))])})):
                    coll5_[d_10_v0_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jhdbecap"))
            return _dafny.Map(coll5_)
        return ((iife5_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_11_i1_ in range(default__.abs(-65))]): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_12_i2_ in range(default__.abs(993))])}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqvwwh")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oudm"))}))

    @staticmethod
    def fm26(p0, p1, globalState):
        return D8_DC20((not(True) if False else True), True, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qoxnse"))) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rjxfs"))))

    @staticmethod
    def fm27(p0, p1, globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([396]))])) + (_dafny.SeqWithoutIsStrInference([410, 166]))})

    @staticmethod
    def fm28(globalState):
        if False:
            return True
        elif True:
            return True

    @staticmethod
    def fm29(p0, p1, globalState):
        return (_dafny.Map({-549: _dafny.Map({False: -220})})) | (_dafny.Map({-451: _dafny.Map({True: 540})}))

    @staticmethod
    def m0(globalState):
        d_13_v0_: bool
        d_13_v0_ = False
        (globalState).f3 = d_13_v0_
        d_14_v1_: str
        d_14_v1_ = _dafny.CodePoint('f')
        d_15_v2_: _dafny.Map
        d_15_v2_ = _dafny.Map({d_13_v0_: d_14_v1_})
        d_16_v3_: int
        d_16_v3_ = 692
        d_15_v2_ = (d_15_v2_).set((d_16_v3_) >= (390), d_14_v1_)
        d_17_v4_: _dafny.Seq
        d_17_v4_ = _dafny.SeqWithoutIsStrInference([d_16_v3_, 164])
        d_18_i0_: int
        d_18_i0_ = 0
        with _dafny.label("0"):
            while (d_16_v3_) <= (default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(d_17_v4_)])), d_16_v3_)):
                with _dafny.c_label("0"):
                    if (d_18_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_18_i0_ = (d_18_i0_) + (1)
                    (globalState).f9 = d_16_v3_
                    if not(not(not(not(d_13_v0_)))):
                        d_19_v5_: _dafny.Array
                        def lambda0_(d_20_v3_):
                            def lambda1_(d_21_i1_):
                                return (d_20_v3_) != (d_20_v3_)

                            return lambda1_

                        init0_ = lambda0_(d_16_v3_)
                        nw0_ = _dafny.Array(None, 18)
                        for i0_0_ in range(nw0_.length(0)):
                            nw0_[i0_0_] = init0_(i0_0_)
                        d_19_v5_ = nw0_
                        d_22_v6_: _dafny.Seq
                        d_22_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "luog"))
                        index0_ = default__.safeIndex(830, (d_19_v5_).length(0))
                        (d_19_v5_)[index0_] = (d_22_v6_) != (d_22_v6_)
                        index1_ = default__.safeIndex(830, (d_19_v5_).length(0))
                        (d_19_v5_)[index1_] = default__.fm0(d_14_v1_, globalState)
                        d_23_v7_: _dafny.Array
                        nw1_ = _dafny.Array(int(0), 27)
                        d_23_v7_ = nw1_
                        index2_ = default__.safeIndex(830, (d_19_v5_).length(0))
                        rhs0_ = (_dafny.Set({d_23_v7_})).ispropersubset(_dafny.Set({d_23_v7_}))
                        rhs1_ = (True if (d_19_v5_)[default__.safeIndex(830, (d_19_v5_).length(0))] else (d_19_v5_)[default__.safeIndex(830, (d_19_v5_).length(0))])
                        lhs0_ = globalState
                        lhs1_ = d_19_v5_
                        lhs2_ = default__.safeIndex(830, (d_19_v5_).length(0))
                        lhs0_.f3 = rhs0_
                        lhs1_[lhs2_] = rhs1_
                        (globalState).f7 = False
                        (globalState).f9 = d_16_v3_
                        (globalState).f9 = d_16_v3_
                    elif True:
                        d_24_v8_: _dafny.MultiSet
                        d_24_v8_ = _dafny.MultiSet([d_13_v0_, False])
                        d_25_v9_: _dafny.Seq
                        d_25_v9_ = _dafny.SeqWithoutIsStrInference([d_13_v0_, d_13_v0_])
                        d_26_v10_: _dafny.Set
                        d_26_v10_ = _dafny.Set({d_13_v0_, d_13_v0_})
                        d_27_v11_: _dafny.Map
                        d_27_v11_ = _dafny.Map({_dafny.Set({130, d_16_v3_}): default__.fm1(d_13_v0_, len(d_26_v10_), d_16_v3_, globalState)})
                        d_28_v12_: _dafny.Set
                        d_28_v12_ = _dafny.Set({d_16_v3_})
                        d_29_v13_: _dafny.Seq
                        d_29_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jlgqgkipu"))
                        d_16_v3_ = (((d_24_v8_)[(d_25_v9_)[default__.safeIndex(d_16_v3_, len(d_25_v9_))]] if ((d_25_v9_)[default__.safeIndex(d_16_v3_, len(d_25_v9_))]) in (d_24_v8_) else d_16_v3_)) + (((d_27_v11_)[d_28_v12_] if (d_28_v12_) in (d_27_v11_) else len(d_29_v13_)))
                        d_30_v14_: C2
                        nw2_ = C2()
                        nw2_.ctor__(d_13_v0_, d_24_v8_, default__.fm0((d_29_v13_)[default__.safeIndex(d_16_v3_, len(d_29_v13_))], globalState))
                        d_30_v14_ = nw2_
                        rhs2_ = (d_16_v3_) - ((len(d_29_v13_) if d_30_v14_.f15 else d_16_v3_))
                        rhs3_ = -264
                        rhs4_ = (0) - (d_16_v3_)
                        rhs5_ = d_13_v0_
                        lhs3_ = globalState
                        lhs3_.f9 = rhs2_
                        d_16_v3_ = rhs3_
                        d_16_v3_ = rhs4_
                        d_13_v0_ = rhs5_
                        d_31_v15_: _dafny.MultiSet
                        d_31_v15_ = _dafny.MultiSet([d_16_v3_, d_16_v3_, d_16_v3_, d_16_v3_, len(d_29_v13_)])
                        (globalState).f9 = (default__.safeModuloInt((0) - ((0) - (((d_31_v15_)[d_16_v3_] if (d_16_v3_) in (d_31_v15_) else len(d_17_v4_)))), len(_dafny.Set({d_30_v14_.f15})))) - (d_16_v3_)
                        (globalState).f7 = d_30_v14_.f15
                    d_32_v16_: _dafny.Map
                    d_32_v16_ = _dafny.Map({d_13_v0_: d_16_v3_})
                    d_33_v17_: D10
                    d_33_v17_ = D10_DC27(d_32_v16_)
                    pat_let_tv0_ = d_13_v0_
                    pat_let_tv1_ = d_16_v3_
                    pat_let_tv2_ = d_32_v16_
                    def iife6_(_pat_let0_0):
                        def iife7_(d_34_dt__update__tmp_h0_):
                            def iife8_(_pat_let1_0):
                                def iife9_(d_35_dt__update_hcf63_h0_):
                                    return D10_DC27(d_35_dt__update_hcf63_h0_)
                                return iife9_(_pat_let1_0)
                            return iife8_((_dafny.Map({not(pat_let_tv0_): pat_let_tv1_})) | (pat_let_tv2_))
                        return iife7_(_pat_let0_0)
                    source0_ = iife6_(d_33_v17_)
                    if source0_.is_DC28:
                        d_36___mcc_h0_ = source0_.cf64
                        d_37___mcc_h1_ = source0_.cf65
                        d_38___mcc_h2_ = source0_.cf66
                        d_39___mcc_h3_ = source0_.cf67
                        d_40_cf67_ = d_39___mcc_h3_
                        d_41_cf66_ = d_38___mcc_h2_
                        d_42_cf65_ = d_37___mcc_h1_
                        d_43_cf64_ = d_36___mcc_h0_
                        d_44_v18_: _dafny.Seq
                        d_44_v18_ = _dafny.SeqWithoutIsStrInference([d_41_cf66_])
                        d_45_v19_: _dafny.Map
                        d_45_v19_ = _dafny.Map({d_14_v1_: d_44_v18_})
                        d_45_v19_ = (d_45_v19_).set(d_14_v1_, _dafny.SeqWithoutIsStrInference([d_40_cf67_, d_13_v0_]))
                        d_46_v20_: _dafny.Set
                        d_46_v20_ = _dafny.Set({d_13_v0_})
                        d_40_cf67_ = (d_46_v20_).ispropersubset((d_46_v20_) - (d_46_v20_))
                        d_47_v21_: _dafny.Set
                        d_47_v21_ = _dafny.Set({default__.fm24(globalState)})
                        d_48_v22_: _dafny.MultiSet
                        d_48_v22_ = _dafny.MultiSet([d_14_v1_])
                        d_49_v23_: _dafny.MultiSet
                        d_49_v23_ = d_48_v22_
                        d_50_v24_: _dafny.Seq
                        d_50_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjr"))
                        d_51_v26_: _dafny.Map
                        d_51_v26_ = _dafny.Map({d_16_v3_: d_47_v21_})
                        d_52_v28_: _dafny.Map
                        d_52_v28_ = _dafny.Map({default__.fm24(globalState): d_14_v1_})
                        d_53_v30_: _dafny.Array
                        nw3_ = _dafny.Array(None, 15)
                        nw3_[int(0)] = d_47_v21_
                        nw3_[int(1)] = default__.fm27(d_49_v23_, d_50_v24_, globalState)
                        nw3_[int(2)] = d_47_v21_
                        nw3_[int(3)] = _dafny.Set({d_17_v4_, d_17_v4_})
                        nw3_[int(4)] = default__.fm27(d_48_v22_, d_50_v24_, globalState)
                        nw3_[int(5)] = d_47_v21_
                        nw3_[int(6)] = d_47_v21_
                        def iife10_():
                            coll6_ = _dafny.Set()
                            compr_6_: _dafny.Seq
                            for compr_6_ in (_dafny.SeqWithoutIsStrInference([d_17_v4_ for d_54_i2_ in range(default__.abs(-30))])).Elements:
                                d_55_v25_: _dafny.Seq = compr_6_
                                if (d_55_v25_) in (_dafny.SeqWithoutIsStrInference([d_17_v4_ for d_54_i2_ in range(default__.abs(-30))])):
                                    coll6_ = coll6_.union(_dafny.Set([d_55_v25_]))
                            return _dafny.Set(coll6_)
                        nw3_[int(7)] = iife10_()
                        
                        nw3_[int(8)] = d_47_v21_
                        nw3_[int(9)] = (((d_51_v26_)[d_16_v3_] if (d_16_v3_) in (d_51_v26_) else _dafny.Set({d_17_v4_, d_17_v4_}))) - (_dafny.Set({(d_17_v4_).set(default__.safeIndex(d_16_v3_, len(d_17_v4_)), d_16_v3_)}))
                        nw3_[int(10)] = d_47_v21_
                        def iife11_():
                            coll7_ = _dafny.Set()
                            compr_7_: _dafny.Seq
                            for compr_7_ in (d_47_v21_).Elements:
                                d_56_v27_: _dafny.Seq = compr_7_
                                if (d_56_v27_) in (d_47_v21_):
                                    coll7_ = coll7_.union(_dafny.Set([d_56_v27_]))
                            return _dafny.Set(coll7_)
                        nw3_[int(11)] = iife11_()
                        
                        nw3_[int(12)] = d_47_v21_
                        def iife12_():
                            coll8_ = _dafny.Set()
                            compr_8_: _dafny.Seq
                            for compr_8_ in (d_52_v28_).keys.Elements:
                                d_57_v29_: _dafny.Seq = compr_8_
                                if (d_57_v29_) in (d_52_v28_):
                                    coll8_ = coll8_.union(_dafny.Set([d_57_v29_]))
                            return _dafny.Set(coll8_)
                        nw3_[int(13)] = iife12_()
                        
                        nw3_[int(14)] = d_47_v21_
                        d_53_v30_ = nw3_
                        index3_ = default__.safeIndex(210, (d_53_v30_).length(0))
                        (d_53_v30_)[index3_] = default__.fm27(d_48_v22_, d_50_v24_, globalState)
                        index4_ = default__.safeIndex(210, (d_53_v30_).length(0))
                        (d_53_v30_)[index4_] = d_47_v21_
                        (globalState).f9 = 162
                    elif source0_.is_DC29:
                        d_58___mcc_h4_ = source0_.cf68
                        d_59_cf68_ = d_58___mcc_h4_
                        d_14_v1_ = d_14_v1_
                        (globalState).f9 = (d_17_v4_)[default__.safeIndex((d_16_v3_) + (d_16_v3_), len(d_17_v4_))]
                        d_60_v31_: _dafny.Array
                        def lambda2_(d_61_v3_, d_62_v0_, d_63_cf68_):
                            def lambda3_(d_64_i3_):
                                return (_dafny.Map({D5_DC14(): d_61_v3_})) | (_dafny.Map({D5_DC14(): (_dafny.MultiSet([d_62_v0_, d_63_cf68_])).cardinality}))

                            return lambda3_

                        init1_ = lambda2_(d_16_v3_, d_13_v0_, d_59_cf68_)
                        nw4_ = _dafny.Array(None, 11)
                        for i0_1_ in range(nw4_.length(0)):
                            nw4_[i0_1_] = init1_(i0_1_)
                        d_60_v31_ = nw4_
                        d_65_v32_: D5
                        d_65_v32_ = D5_DC14()
                        d_66_v33_: _dafny.Map
                        d_66_v33_ = _dafny.Map({d_65_v32_: d_16_v3_})
                        index5_ = default__.safeIndex(36, (d_60_v31_).length(0))
                        (d_60_v31_)[index5_] = d_66_v33_
                        index6_ = default__.safeIndex(36, (d_60_v31_).length(0))
                        (d_60_v31_)[index6_] = (d_66_v33_) | (d_66_v33_)
                        d_67_v34_: _dafny.Seq
                        d_67_v34_ = _dafny.SeqWithoutIsStrInference([True])
                        d_68_v35_: D1
                        d_68_v35_ = D1_DC2(d_59_cf68_, d_13_v0_, d_59_cf68_, d_14_v1_, d_14_v1_)
                        d_69_v36_: D1
                        d_69_v36_ = D1_DC2((len(d_67_v34_)) == (d_16_v3_), d_13_v0_, d_59_cf68_, default__.fm14(d_68_v35_, len(d_67_v34_), globalState), d_14_v1_)
                        pat_let_tv3_ = d_13_v0_
                        pat_let_tv4_ = d_14_v1_
                        pat_let_tv5_ = d_59_cf68_
                        def iife14_(_pat_let3_0):
                            def iife15_(d_70_dt__update__tmp_h2_):
                                def iife16_(_pat_let4_0):
                                    def iife17_(d_71_dt__update_hcf3_h0_):
                                        def iife18_(_pat_let5_0):
                                            def iife19_(d_72_dt__update_hcf6_h0_):
                                                return D1_DC2((d_70_dt__update__tmp_h2_).cf2, d_71_dt__update_hcf3_h0_, (d_70_dt__update__tmp_h2_).cf4, (d_70_dt__update__tmp_h2_).cf5, d_72_dt__update_hcf6_h0_)
                                            return iife19_(_pat_let5_0)
                                        return iife18_(pat_let_tv4_)
                                    return iife17_(_pat_let4_0)
                                return iife16_(pat_let_tv3_)
                            return iife15_(_pat_let3_0)
                        def iife13_(_pat_let2_0):
                            def iife20_(d_73_dt__update__tmp_h3_):
                                def iife21_(_pat_let6_0):
                                    def iife22_(d_74_dt__update_hcf4_h0_):
                                        def iife23_(_pat_let7_0):
                                            def iife24_(d_75_dt__update_hcf3_h1_):
                                                return D1_DC2((d_73_dt__update__tmp_h3_).cf2, d_75_dt__update_hcf3_h1_, d_74_dt__update_hcf4_h0_, (d_73_dt__update__tmp_h3_).cf5, (d_73_dt__update__tmp_h3_).cf6)
                                            return iife24_(_pat_let7_0)
                                        return iife23_(pat_let_tv5_)
                                    return iife22_(_pat_let6_0)
                                return iife21_(False)
                            return iife20_(_pat_let2_0)
                        d_68_v35_ = iife13_(iife14_(d_68_v35_))
                    elif True:
                        d_76___mcc_h5_ = source0_.cf63
                        d_77_cf63_ = d_76___mcc_h5_
                        d_78_v37_: _dafny.Array
                        nw5_ = _dafny.Array(int(0), 2)
                        d_78_v37_ = nw5_
                        (globalState).f1 = (d_78_v37_ if d_13_v0_ else d_78_v37_)
                        d_79_v38_: _dafny.MultiSet
                        d_79_v38_ = _dafny.MultiSet([d_13_v0_])
                        d_80_v39_: C2
                        nw6_ = C2()
                        nw6_.ctor__(True, d_79_v38_, d_13_v0_)
                        d_80_v39_ = nw6_
                        d_81_v40_: _dafny.Map
                        d_81_v40_ = _dafny.Map({d_80_v39_: d_16_v3_})
                        d_82_v41_: _dafny.Set
                        d_82_v41_ = _dafny.Set({(len((d_81_v40_).set(d_80_v39_, d_16_v3_)) if d_13_v0_ else d_16_v3_), (d_16_v3_) + (d_16_v3_)})
                        d_82_v41_ = (d_82_v41_) | (d_82_v41_)
                        (globalState).f9 = (d_16_v3_) + (d_16_v3_)
                        d_83_v42_: _dafny.Seq
                        d_83_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bdgl"))
                        d_83_v42_ = d_83_v42_
                    d_84_v43_: _dafny.MultiSet
                    d_84_v43_ = _dafny.MultiSet([d_13_v0_, False, d_13_v0_, False])
                    d_85_v44_: T0
                    nw7_ = C2()
                    nw7_.ctor__(d_13_v0_, d_84_v43_, d_13_v0_)
                    d_85_v44_ = nw7_
                    d_86_v45_: _dafny.MultiSet
                    d_86_v45_ = _dafny.MultiSet([d_85_v44_, d_85_v44_])
                    d_87_v46_: _dafny.Map
                    d_87_v46_ = _dafny.Map({d_16_v3_: d_85_v44_.f12})
                    d_16_v3_ = default__.fm1(default__.fm0(d_14_v1_, globalState), (0) - ((d_86_v45_).cardinality), len(d_87_v46_), globalState)
                    pass
            pass
        d_88_i4_: int
        d_88_i4_ = 0
        with _dafny.label("1"):
            while d_13_v0_:
                with _dafny.c_label("1"):
                    if (d_88_i4_) >= (100):
                        raise _dafny.Break("1")
                    d_88_i4_ = (d_88_i4_) + (1)
                    d_89_v47_: _dafny.Seq
                    d_89_v47_ = _dafny.SeqWithoutIsStrInference([(d_13_v0_) == (d_13_v0_), d_13_v0_, d_13_v0_])
                    if (d_89_v47_)[default__.safeIndex(d_16_v3_, len(d_89_v47_))]:
                        d_90_v48_: _dafny.MultiSet
                        d_90_v48_ = _dafny.MultiSet([d_13_v0_, d_13_v0_])
                        d_91_v49_: C2
                        nw8_ = C2()
                        nw8_.ctor__(d_13_v0_, d_90_v48_, False)
                        d_91_v49_ = nw8_
                        d_90_v48_ = (d_90_v48_ if not((d_16_v3_) < (d_16_v3_)) else d_90_v48_)
                        (globalState).f9 = d_16_v3_
                        d_92_v50_: _dafny.Seq
                        d_92_v50_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                        d_93_v51_: _dafny.Map
                        d_93_v51_ = _dafny.Map({d_16_v3_: d_91_v49_.f15})
                        d_94_v52_: _dafny.Seq
                        d_94_v52_ = _dafny.SeqWithoutIsStrInference([d_93_v51_])
                        d_95_v53_: _dafny.Map
                        d_95_v53_ = _dafny.Map({(0) - (len(((d_92_v50_) + (d_92_v50_)).set(default__.safeIndex(d_16_v3_, len((d_92_v50_) + (d_92_v50_))), d_14_v1_))): (d_16_v3_) - (len((d_94_v52_)[default__.safeIndex(d_16_v3_, len(d_94_v52_))]))})
                        d_95_v53_ = (d_95_v53_).set(d_16_v3_, d_16_v3_)
                        d_96_v54_: T0
                        nw9_ = C1()
                        nw9_.ctor__(d_16_v3_, d_90_v48_, d_91_v49_.f15)
                        d_96_v54_ = nw9_
                        d_97_v55_: _dafny.Map
                        d_97_v55_ = _dafny.Map({d_96_v54_: d_17_v4_})
                        d_98_v56_: _dafny.Set
                        d_98_v56_ = _dafny.Set({d_16_v3_, d_16_v3_, (_dafny.MultiSet(((d_97_v55_)[d_96_v54_] if (d_96_v54_) in (d_97_v55_) else d_17_v4_))).cardinality, len(_dafny.Set({d_14_v1_, d_14_v1_}))})
                        d_99_v57_: _dafny.Map
                        d_99_v57_ = _dafny.Map({_dafny.CodePoint('a'): d_90_v48_})
                        d_100_v58_: C3
                        nw10_ = C3()
                        nw10_.ctor__((_dafny.MultiSet([d_13_v0_, d_13_v0_])).cardinality, (d_98_v56_).ispropersubset(d_98_v56_), ((d_99_v57_)[_dafny.CodePoint('y')] if (_dafny.CodePoint('y')) in (d_99_v57_) else d_90_v48_), (_dafny.Set({d_16_v3_})).ispropersubset(_dafny.Set({d_16_v3_})))
                        d_100_v58_ = nw10_
                    elif True:
                        d_101_v59_: bool
                        d_101_v59_ = not (d_13_v0_) or (True)
                        d_101_v59_ = not (d_13_v0_) or (d_13_v0_)
                        (globalState).f7 = d_13_v0_
                        d_102_v60_: C0
                        nw11_ = C0()
                        nw11_.ctor__((d_13_v0_ if d_13_v0_ else d_13_v0_), d_16_v3_, _dafny.MultiSet([True, d_13_v0_]), (d_16_v3_) not in (d_17_v4_))
                        d_102_v60_ = nw11_
                        d_102_v60_ = d_102_v60_
                        d_103_v61_: _dafny.MultiSet
                        d_103_v61_ = _dafny.MultiSet([d_13_v0_])
                        d_104_v62_: C2
                        nw12_ = C2()
                        nw12_.ctor__(d_13_v0_, (d_103_v61_) | (_dafny.MultiSet([(d_102_v60_).f17])), d_13_v0_)
                        d_104_v62_ = nw12_
                        d_105_v63_: _dafny.Seq
                        d_105_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
                        d_106_v64_: _dafny.MultiSet
                        d_106_v64_ = _dafny.MultiSet([default__.fm1((d_102_v60_).f17, (d_102_v60_).f18, len(d_15_v2_), globalState)])
                        d_107_v65_: _dafny.Map
                        d_107_v65_ = _dafny.Map({d_13_v0_: d_16_v3_})
                        rhs6_ = default__.fm9(d_16_v3_, not ((d_102_v60_).f17) or ((d_102_v60_).fm2(d_106_v64_, (0) - (d_16_v3_), d_107_v65_, d_13_v0_, globalState)), globalState)
                        rhs7_ = d_16_v3_
                        rhs8_ = ((350) + (d_16_v3_)) * (d_16_v3_)
                        lhs4_ = globalState
                        d_105_v63_ = rhs6_
                        lhs4_.f9 = rhs7_
                        d_16_v3_ = rhs8_
                    d_108_v66_: _dafny.Array
                    def lambda4_(d_109_v0_):
                        def lambda5_(d_110_i5_):
                            return (d_109_v0_ if False else d_109_v0_)

                        return lambda5_

                    init2_ = lambda4_(d_13_v0_)
                    nw13_ = _dafny.Array(None, 7)
                    for i0_2_ in range(nw13_.length(0)):
                        nw13_[i0_2_] = init2_(i0_2_)
                    d_108_v66_ = nw13_
                    index7_ = default__.safeIndex(460, (d_108_v66_).length(0))
                    (d_108_v66_)[index7_] = not(d_13_v0_)
                    d_111_v67_: _dafny.Map
                    d_111_v67_ = _dafny.Map({d_13_v0_: d_16_v3_})
                    index8_ = default__.safeIndex(460, (d_108_v66_).length(0))
                    (d_108_v66_)[index8_] = (True) in ((d_111_v67_) | (d_111_v67_))
                    (globalState).f6 = d_89_v47_
                    d_112_v68_: _dafny.MultiSet
                    d_112_v68_ = _dafny.MultiSet([d_14_v1_])
                    d_113_v69_: _dafny.Seq
                    d_113_v69_ = _dafny.SeqWithoutIsStrInference([d_14_v1_])
                    index9_ = default__.safeIndex(460, (d_108_v66_).length(0))
                    (d_108_v66_)[index9_] = (_dafny.MultiSet(d_113_v69_)).ispropersubset(d_112_v68_)
                    pass
            pass
        d_114_v70_: _dafny.Seq
        d_114_v70_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dslsfkrf"))
        d_115_v71_: _dafny.Map
        d_115_v71_ = _dafny.Map({d_114_v70_: len(d_114_v70_)})
        d_116_v72_: _dafny.Set
        d_116_v72_ = _dafny.Set({default__.fm0(d_14_v1_, globalState)})
        d_117_v73_: _dafny.MultiSet
        d_117_v73_ = _dafny.MultiSet([d_13_v0_])
        d_118_v74_: _dafny.Array
        nw14_ = _dafny.Array(None, 26)
        nw14_[int(0)] = d_13_v0_
        nw14_[int(1)] = d_13_v0_
        nw14_[int(2)] = (True) == (d_13_v0_)
        nw14_[int(3)] = d_13_v0_
        nw14_[int(4)] = d_13_v0_
        nw14_[int(5)] = d_13_v0_
        nw14_[int(6)] = (d_115_v71_) != ((d_115_v71_).set(d_114_v70_, d_16_v3_))
        nw14_[int(7)] = d_13_v0_
        nw14_[int(8)] = (d_13_v0_) not in (d_116_v72_)
        nw14_[int(9)] = False
        nw14_[int(10)] = d_13_v0_
        nw14_[int(11)] = not(d_13_v0_)
        nw14_[int(12)] = (d_116_v72_).isdisjoint(d_116_v72_)
        nw14_[int(13)] = True
        nw14_[int(14)] = default__.fm0(d_14_v1_, globalState)
        nw14_[int(15)] = True
        nw14_[int(16)] = (d_13_v0_ if d_13_v0_ else False)
        nw14_[int(17)] = d_13_v0_
        nw14_[int(18)] = d_13_v0_
        nw14_[int(19)] = d_13_v0_
        nw14_[int(20)] = ((d_117_v73_).cardinality) == (len(d_114_v70_))
        nw14_[int(21)] = (d_13_v0_)
        nw14_[int(22)] = d_13_v0_
        nw14_[int(23)] = d_13_v0_
        nw14_[int(24)] = d_13_v0_
        nw14_[int(25)] = d_13_v0_
        d_118_v74_ = nw14_
        d_119_v75_: _dafny.Seq
        d_119_v75_ = _dafny.SeqWithoutIsStrInference([d_118_v74_, d_118_v74_, (d_118_v74_ if d_13_v0_ else d_118_v74_)])
        d_120_v76_: _dafny.MultiSet
        d_120_v76_ = _dafny.MultiSet([d_16_v3_, d_16_v3_, d_16_v3_, default__.fm1(d_13_v0_, d_16_v3_, 496, globalState)])
        rhs9_ = (d_119_v75_)[default__.safeIndex(default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([(d_114_v70_)[default__.safeIndex(877, len(d_114_v70_))] for d_121_i6_ in range(default__.abs(792))])), d_16_v3_), len(d_119_v75_))]
        rhs10_ = d_118_v74_
        rhs11_ = (not(d_13_v0_) if (d_114_v70_) < (d_114_v70_) else (d_16_v3_) in (d_120_v76_))
        d_118_v74_ = rhs9_
        d_118_v74_ = rhs10_
        d_13_v0_ = rhs11_
        d_122_v77_: _dafny.Array
        nw15_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 2)
        d_122_v77_ = nw15_
        d_123_v78_: D5
        d_123_v78_ = D5_DC13(d_16_v3_, d_13_v0_, d_16_v3_, d_122_v77_, d_13_v0_)
        hi0_ = (d_123_v78_).cf30
        for d_124_i7_ in range(d_16_v3_, hi0_):
            d_125_v79_: _dafny.MultiSet
            d_125_v79_ = (_dafny.MultiSet([d_14_v1_])).intersection(_dafny.MultiSet([_dafny.CodePoint('g')]))
            source1_ = d_125_v79_
            d_126___mcc_h6_ = source1_
            d_127_cf77_ = d_126___mcc_h6_
            d_128_v80_: _dafny.Seq
            d_128_v80_ = _dafny.SeqWithoutIsStrInference([True])
            d_129_v81_: _dafny.Seq
            d_129_v81_ = _dafny.SeqWithoutIsStrInference([d_13_v0_, d_13_v0_, False, (d_128_v80_)[default__.safeIndex((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcvbfjwrl")))), len(d_128_v80_))]])
            (globalState).f9 = (d_16_v3_ if (d_124_i7_) in (d_120_v76_) else (d_124_i7_) + (len(d_129_v81_)))
            d_130_v82_: _dafny.Array
            def lambda6_(d_131_v3_):
                def lambda7_(d_132_i8_):
                    return (d_132_i8_) + (d_131_v3_)

                return lambda7_

            init3_ = lambda6_(d_16_v3_)
            nw16_ = _dafny.Array(None, 1)
            for i0_3_ in range(nw16_.length(0)):
                nw16_[i0_3_] = init3_(i0_3_)
            d_130_v82_ = nw16_
            index10_ = default__.safeIndex(618, (d_130_v82_).length(0))
            (d_130_v82_)[index10_] = (d_16_v3_) * (d_16_v3_)
            d_133_v83_: _dafny.MultiSet
            d_133_v83_ = _dafny.MultiSet([d_130_v82_, d_130_v82_])
            d_134_v84_: _dafny.Set
            d_134_v84_ = _dafny.Set({d_124_i7_, d_124_i7_, d_16_v3_})
            d_135_v85_: D2
            d_135_v85_ = D2_DC5((0) - (d_124_i7_), (d_133_v83_).cardinality, len(d_134_v84_), ((_dafny.MultiSet(d_128_v80_)).set(d_13_v0_, default__.abs(default__.fm1(d_13_v0_, d_124_i7_, d_16_v3_, globalState)))).cardinality, d_129_v81_)
            index11_ = default__.safeIndex(618, (d_130_v82_).length(0))
            rhs12_ = d_130_v82_
            rhs13_ = default__.fm1(d_13_v0_, (d_16_v3_) * (d_16_v3_), d_16_v3_, globalState)
            rhs14_ = d_135_v85_
            lhs5_ = globalState
            lhs6_ = d_130_v82_
            lhs7_ = default__.safeIndex(618, (d_130_v82_).length(0))
            lhs5_.f4 = rhs12_
            lhs6_[lhs7_] = rhs13_
            d_135_v85_ = rhs14_
            d_14_v1_ = (d_14_v1_ if True else d_14_v1_)
            d_136_v86_: _dafny.Map
            d_136_v86_ = _dafny.Map({d_14_v1_: d_16_v3_})
            d_136_v86_ = (d_136_v86_).set(d_14_v1_, default__.fm1(d_13_v0_, (d_130_v82_)[default__.safeIndex(618, (d_130_v82_).length(0))], (0) - ((d_130_v82_)[default__.safeIndex(618, (d_130_v82_).length(0))]), globalState))
            d_13_v0_ = False
            d_137_v87_: _dafny.Seq
            d_137_v87_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_13_v0_])])
            d_138_v88_: _dafny.Map
            d_138_v88_ = _dafny.Map({(d_137_v87_)[default__.safeIndex(d_16_v3_, len(d_137_v87_))]: _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_14_v1_ for d_139_i9_ in range(default__.abs(529))]), d_114_v70_, d_114_v70_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bultv")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yvdibk"))])) + (_dafny.SeqWithoutIsStrInference([d_114_v70_, d_114_v70_])))})
            d_140_v89_: _dafny.Seq
            d_140_v89_ = _dafny.SeqWithoutIsStrInference([d_13_v0_])
            d_141_v90_: _dafny.MultiSet
            d_141_v90_ = _dafny.MultiSet([d_114_v70_, d_114_v70_, d_114_v70_])
            d_138_v88_ = (d_138_v88_).set(d_140_v89_, d_141_v90_)
            if not(not(d_13_v0_)):
                d_142_v91_: C4
                nw17_ = C4()
                nw17_.ctor__((_dafny.MultiSet(d_140_v89_)) | (d_117_v73_), not(not(d_13_v0_)))
                d_142_v91_ = nw17_
                (globalState).f9 = default__.safeDivisionInt(d_124_i7_, d_124_i7_)
                d_114_v70_ = d_114_v70_
                d_143_v92_: _dafny.Set
                d_143_v92_ = _dafny.Set({d_124_i7_})
                d_144_v93_: D3
                d_144_v93_ = D3_DC8((0) - ((0) - (d_16_v3_)), d_124_i7_)
                d_145_v94_: _dafny.Map
                d_145_v94_ = _dafny.Map({(len(d_143_v92_)) in (d_143_v92_): d_144_v93_})
                d_145_v94_ = (d_145_v94_).set(d_13_v0_, d_144_v93_)
                d_146_v95_: _dafny.Map
                d_146_v95_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uli"))})
                d_147_v96_: _dafny.Map
                d_147_v96_ = _dafny.Map({default__.fm29((d_140_v89_)[default__.safeIndex(len(d_143_v92_), len(d_140_v89_))], d_16_v3_, globalState): len(_dafny.Set({917, len(d_17_v4_), (d_142_v91_).fm3(d_13_v0_, d_146_v95_, d_13_v0_, d_124_i7_, globalState)}))})
                d_148_v97_: _dafny.Map
                d_148_v97_ = _dafny.Map({d_13_v0_: d_124_i7_})
                d_149_v98_: _dafny.Map
                d_149_v98_ = _dafny.Map({d_124_i7_: d_148_v97_})
                d_150_v99_: _dafny.Map
                d_150_v99_ = _dafny.Map({d_13_v0_: d_13_v0_})
                d_147_v96_ = (d_147_v96_).set((d_149_v98_) | (_dafny.Map({d_16_v3_: d_148_v97_})), (len(d_150_v99_)) * (d_124_i7_))
            elif True:
                d_151_v100_: _dafny.Array
                nw18_ = _dafny.Array(None, 2)
                d_151_v100_ = nw18_
                d_152_v101_: D9
                d_152_v101_ = D9_DC23(_dafny.Map({d_13_v0_: d_151_v100_}))
                d_153_v102_: _dafny.Map
                d_153_v102_ = _dafny.Map({d_16_v3_: _dafny.Map({d_13_v0_: d_151_v100_})})
                d_154_v103_: _dafny.Map
                d_154_v103_ = _dafny.Map({d_13_v0_: d_151_v100_})
                pat_let_tv6_ = d_153_v102_
                pat_let_tv7_ = d_16_v3_
                pat_let_tv8_ = d_16_v3_
                pat_let_tv9_ = d_153_v102_
                pat_let_tv10_ = d_154_v103_
                d_155_v104_: _dafny.Set
                def iife25_(_pat_let8_0):
                    def iife26_(d_156_dt__update__tmp_h6_):
                        def iife27_(_pat_let9_0):
                            def iife28_(d_157_dt__update_hcf55_h0_):
                                return D9_DC23(d_157_dt__update_hcf55_h0_)
                            return iife28_(_pat_let9_0)
                        return iife27_(((pat_let_tv6_)[pat_let_tv7_] if (pat_let_tv8_) in (pat_let_tv9_) else pat_let_tv10_))
                    return iife26_(_pat_let8_0)
                d_155_v104_ = _dafny.Set({iife25_(d_152_v101_), d_152_v101_, (d_152_v101_ if True else D9_DC23(d_154_v103_)), d_152_v101_, d_152_v101_})
                d_155_v104_ = (d_155_v104_) | (d_155_v104_)
                index12_ = default__.safeIndex(705, (d_118_v74_).length(0))
                (d_118_v74_)[index12_] = d_13_v0_
                d_158_v105_: _dafny.Map
                d_158_v105_ = _dafny.Map({not((d_13_v0_ if d_13_v0_ else not(d_13_v0_))): not((d_13_v0_) and (d_13_v0_))})
                index13_ = default__.safeIndex(705, (d_118_v74_).length(0))
                (d_118_v74_)[index13_] = ((d_158_v105_)[(d_13_v0_ if not(d_13_v0_) else d_13_v0_)] if ((d_13_v0_ if not(d_13_v0_) else d_13_v0_)) in (d_158_v105_) else d_13_v0_)
                d_159_v106_: D9
                d_159_v106_ = D9_DC24(d_114_v70_, d_13_v0_, d_16_v3_)
                d_160_v107_: _dafny.Array
                nw19_ = _dafny.Array(None, 20)
                nw19_[int(0)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(1)] = d_13_v0_
                nw19_[int(2)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(3)] = False
                nw19_[int(4)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(5)] = (d_159_v106_).cf57
                nw19_[int(6)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(7)] = not(d_13_v0_)
                nw19_[int(8)] = (d_140_v89_)[default__.safeIndex(d_16_v3_, len(d_140_v89_))]
                nw19_[int(9)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(10)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(11)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(12)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(13)] = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                nw19_[int(14)] = d_13_v0_
                nw19_[int(15)] = d_13_v0_
                nw19_[int(16)] = d_13_v0_
                nw19_[int(17)] = d_13_v0_
                nw19_[int(18)] = d_13_v0_
                nw19_[int(19)] = d_13_v0_
                d_160_v107_ = nw19_
                d_161_v108_: _dafny.Set
                d_161_v108_ = _dafny.Set({d_160_v107_, d_160_v107_})
                d_162_v109_: _dafny.Map
                d_162_v109_ = _dafny.Map({d_124_i7_: len(d_158_v105_)})
                d_163_v110_: _dafny.Map
                d_163_v110_ = _dafny.Map({d_13_v0_: d_162_v109_})
                d_164_v111_: _dafny.Map
                d_164_v111_ = _dafny.Map({(len(d_158_v105_)) != (len(d_161_v108_)): ((d_163_v110_).set((d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))], _dafny.Map({d_16_v3_: len(_dafny.SeqWithoutIsStrInference([305]))}))) | (d_163_v110_)})
                d_164_v111_ = (d_164_v111_).set(((d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))] if (d_140_v89_)[default__.safeIndex(d_124_i7_, len(d_140_v89_))] else d_13_v0_), (d_163_v110_) | (_dafny.Map({d_13_v0_: d_162_v109_})))
                d_165_v112_: C1
                nw20_ = C1()
                nw20_.ctor__(300, _dafny.MultiSet([(d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]]), d_13_v0_)
                d_165_v112_ = nw20_
                d_166_v113_: _dafny.Map
                d_166_v113_ = _dafny.Map({True: d_16_v3_})
                d_167_v114_: _dafny.Array
                nw21_ = _dafny.Array(None, 10)
                nw21_[int(0)] = len(d_166_v113_)
                nw21_[int(1)] = d_16_v3_
                nw21_[int(2)] = d_16_v3_
                nw21_[int(3)] = len(d_140_v89_)
                nw21_[int(4)] = d_124_i7_
                nw21_[int(5)] = 93
                nw21_[int(6)] = (d_17_v4_)[default__.safeIndex(d_165_v112_.f16, len(d_17_v4_))]
                nw21_[int(7)] = d_165_v112_.f16
                nw21_[int(8)] = (0) - (d_16_v3_)
                nw21_[int(9)] = len(d_17_v4_)
                d_167_v114_ = nw21_
                d_168_v115_: _dafny.Seq
                d_168_v115_ = _dafny.SeqWithoutIsStrInference([d_167_v114_])
                d_169_v116_: _dafny.Array
                nw22_ = _dafny.Array(None, 17)
                nw22_[int(0)] = d_160_v107_
                nw22_[int(1)] = d_160_v107_
                nw22_[int(2)] = d_160_v107_
                nw22_[int(3)] = d_160_v107_
                nw22_[int(4)] = d_160_v107_
                nw22_[int(5)] = d_160_v107_
                nw22_[int(6)] = d_160_v107_
                nw22_[int(7)] = d_160_v107_
                nw22_[int(8)] = d_160_v107_
                nw22_[int(9)] = d_160_v107_
                nw22_[int(10)] = d_160_v107_
                nw22_[int(11)] = d_160_v107_
                nw22_[int(12)] = d_160_v107_
                nw22_[int(13)] = d_160_v107_
                nw22_[int(14)] = d_160_v107_
                nw22_[int(15)] = d_160_v107_
                nw22_[int(16)] = d_160_v107_
                d_169_v116_ = nw22_
                d_170_v117_: D8
                d_170_v117_ = D8_DC21(not(d_13_v0_), d_169_v116_, d_16_v3_, d_160_v107_, d_16_v3_)
                rhs15_ = d_165_v112_
                rhs16_ = ((d_170_v117_).cf52 if ((d_168_v115_).set(default__.safeIndex((0) - (d_165_v112_.f16), len(d_168_v115_)), d_167_v114_)) <= (d_168_v115_) else d_160_v107_)
                rhs17_ = (d_118_v74_)[default__.safeIndex(705, (d_118_v74_).length(0))]
                lhs8_ = globalState
                d_165_v112_ = rhs15_
                d_160_v107_ = rhs16_
                lhs8_.f8 = rhs17_
                index14_ = default__.safeIndex(553, (d_167_v114_).length(0))
                (d_167_v114_)[index14_] = d_165_v112_.f16
                index15_ = default__.safeIndex(553, (d_167_v114_).length(0))
                (d_167_v114_)[index15_] = d_124_i7_

    @staticmethod
    def Main(noArgsParameter__):
        d_171_v1_: int
        d_171_v1_ = -351
        d_172_v2_: _dafny.Seq
        d_172_v2_ = _dafny.SeqWithoutIsStrInference([d_171_v1_, 231, d_171_v1_])
        d_173_v3_: str
        d_173_v3_ = _dafny.CodePoint('p')
        d_174_v4_: _dafny.Map
        def iife29_():
            coll9_ = _dafny.Map()
            compr_9_: int
            for compr_9_ in (d_172_v2_).Elements:
                d_175_v0_: int = compr_9_
                if (d_175_v0_) in (d_172_v2_):
                    coll9_[(d_175_v0_) - (d_171_v1_)] = False
            return _dafny.Map(coll9_)
        d_174_v4_ = _dafny.Map({iife29_()
        : d_173_v3_})
        d_176_v5_: _dafny.Array
        def lambda8_(d_177_v1_):
            def lambda9_(d_178_i0_):
                return default__.safeModuloInt(d_178_i0_, d_177_v1_)

            return lambda9_

        init4_ = lambda8_(d_171_v1_)
        nw23_ = _dafny.Array(None, 29)
        for i0_4_ in range(nw23_.length(0)):
            nw23_[i0_4_] = init4_(i0_4_)
        d_176_v5_ = nw23_
        d_179_v6_: bool
        d_179_v6_ = False
        d_180_v7_: _dafny.Seq
        d_180_v7_ = _dafny.SeqWithoutIsStrInference([d_179_v6_, True])
        d_181_globalState_: GlobalState
        nw24_ = GlobalState()
        nw24_.ctor__(d_174_v4_, d_176_v5_, True, False, d_176_v5_, 724, d_180_v7_, True, True, -663, False)
        d_181_globalState_ = nw24_
        default__.m0(d_181_globalState_)
        d_182_i1_: int
        d_182_i1_ = 0
        with _dafny.label("2"):
            while (d_179_v6_ if False else not(not(not(d_179_v6_)))):
                with _dafny.c_label("2"):
                    if (d_182_i1_) >= (100):
                        raise _dafny.Break("2")
                    d_182_i1_ = (d_182_i1_) + (1)
                    (d_181_globalState_).f8 = default__.fm0(d_173_v3_, d_181_globalState_)
                    d_183_v8_: _dafny.Map
                    d_183_v8_ = _dafny.Map({default__.safeModuloInt(d_171_v1_, d_171_v1_): default__.fm0(d_173_v3_, d_181_globalState_)})
                    d_184_v10_: _dafny.Map
                    d_184_v10_ = _dafny.Map({d_179_v6_: d_171_v1_})
                    def iife30_():
                        def iife31_():
                            coll11_ = _dafny.Set()
                            compr_11_: int
                            for compr_11_ in (d_183_v8_).keys.Elements:
                                d_186_v11_: int = compr_11_
                                if (d_186_v11_) in (d_183_v8_):
                                    coll11_ = coll11_.union(_dafny.Set([(d_186_v11_) - (206)]))
                            return _dafny.Set(coll11_)
                        coll10_ = _dafny.Map()
                        compr_10_: int
                        for compr_10_ in (_dafny.SeqWithoutIsStrInference([default__.fm1(d_179_v6_, d_171_v1_, len(d_184_v10_), d_181_globalState_)])).Elements:
                            d_185_v9_: int = compr_10_
                            if (d_185_v9_) in (_dafny.SeqWithoutIsStrInference([default__.fm1(d_179_v6_, d_171_v1_, len(d_184_v10_), d_181_globalState_)])):
                                coll10_[default__.safeDivisionInt(d_185_v9_, (_dafny.MultiSet([d_171_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kengs"))), len(iife31_()
                                )])).cardinality)] = (_dafny.MultiSet([d_179_v6_, d_179_v6_])).ispropersubset(_dafny.MultiSet([d_179_v6_, d_179_v6_]))
                        return _dafny.Map(coll10_)
                    d_183_v8_ = iife30_()
                    
                    d_187_v12_: _dafny.MultiSet
                    d_187_v12_ = _dafny.MultiSet([d_179_v6_])
                    d_188_v13_: C4
                    nw25_ = C4()
                    nw25_.ctor__(d_187_v12_, d_179_v6_)
                    d_188_v13_ = nw25_
                    if d_179_v6_:
                        index16_ = default__.safeIndex(968, (d_176_v5_).length(0))
                        (d_176_v5_)[index16_] = (len(_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_189_i2_ in range(default__.abs(284))]))) - (d_171_v1_)
                        index17_ = default__.safeIndex(968, (d_176_v5_).length(0))
                        (d_176_v5_)[index17_] = d_171_v1_
                        index18_ = default__.safeIndex(968, (d_176_v5_).length(0))
                        (d_176_v5_)[index18_] = d_171_v1_
                        (d_181_globalState_).f9 = d_171_v1_
                        d_190_v14_: C4
                        nw26_ = C4()
                        nw26_.ctor__(d_187_v12_, d_179_v6_)
                        d_190_v14_ = nw26_
                        (d_181_globalState_).f7 = d_179_v6_
                    elif True:
                        d_191_v15_: _dafny.Seq
                        d_191_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fuaink"))
                        d_192_v16_: _dafny.MultiSet
                        d_192_v16_ = _dafny.MultiSet([d_191_v15_, (d_191_v15_ if d_179_v6_ else _dafny.SeqWithoutIsStrInference([d_173_v3_ for d_193_i3_ in range(default__.abs(217))])), (d_191_v15_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdvl"))), d_191_v15_, (d_191_v15_ if ((d_183_v8_)[(0) - (d_171_v1_)] if ((0) - (d_171_v1_)) in (d_183_v8_) else d_179_v6_) else (D9_DC25(d_191_v15_, d_179_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsmf"))))).cf59)])
                        (d_181_globalState_).f9 = (d_192_v16_).cardinality
                        d_184_v10_ = (d_184_v10_).set(d_179_v6_, (d_188_v13_).fm3(d_179_v6_, default__.fm25(d_173_v3_, d_173_v3_, not(d_179_v6_), d_181_globalState_), not(d_179_v6_), d_171_v1_, d_181_globalState_))
                        d_194_v17_: C2
                        nw27_ = C2()
                        nw27_.ctor__(d_179_v6_, d_187_v12_, d_179_v6_)
                        d_194_v17_ = nw27_
                        d_195_v18_: T0
                        nw28_ = C4()
                        nw28_.ctor__(_dafny.MultiSet([d_194_v17_.f15]), d_194_v17_.f15)
                        d_195_v18_ = nw28_
                        d_196_v19_: _dafny.Map
                        d_196_v19_ = _dafny.Map({d_171_v1_: d_195_v18_})
                        d_197_v20_: D10
                        d_197_v20_ = D10_DC28(d_194_v17_, d_194_v17_.f15, d_195_v18_.f12, d_195_v18_.f12)
                        d_198_v21_: _dafny.Array
                        nw29_ = _dafny.Array(None, 8)
                        nw29_[int(0)] = True
                        nw29_[int(1)] = d_194_v17_.f15
                        nw29_[int(2)] = d_194_v17_.f15
                        nw29_[int(3)] = (d_171_v1_) != (len(d_196_v19_))
                        nw29_[int(4)] = (d_197_v20_).cf67
                        nw29_[int(5)] = True
                        nw29_[int(6)] = d_194_v17_.f15
                        nw29_[int(7)] = d_195_v18_.f12
                        d_198_v21_ = nw29_
                        index19_ = default__.safeIndex(591, (d_198_v21_).length(0))
                        (d_198_v21_)[index19_] = (d_171_v1_) >= (489)
                        index20_ = default__.safeIndex(591, (d_198_v21_).length(0))
                        (d_198_v21_)[index20_] = d_195_v18_.f12
                        (d_181_globalState_).f7 = d_194_v17_.f15
                    pass
            pass
        d_199_i4_: int
        d_199_i4_ = 0
        with _dafny.label("3"):
            while d_179_v6_:
                with _dafny.c_label("3"):
                    if (d_199_i4_) >= (100):
                        raise _dafny.Break("3")
                    d_199_i4_ = (d_199_i4_) + (1)
                    d_200_v22_: _dafny.Seq
                    d_200_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrpi"))
                    d_201_v23_: _dafny.MultiSet
                    d_201_v23_ = _dafny.MultiSet([d_179_v6_])
                    d_202_v24_: T0
                    nw30_ = C3()
                    nw30_.ctor__(len(d_200_v22_), not(d_179_v6_), d_201_v23_, d_179_v6_)
                    d_202_v24_ = nw30_
                    d_202_v24_ = (d_202_v24_ if d_202_v24_.f12 else d_202_v24_)
                    d_203_v25_: _dafny.Array
                    def lambda10_(d_204_v24_):
                        def lambda11_(d_205_i5_):
                            return d_204_v24_.f11

                        return lambda11_

                    init5_ = lambda10_(d_202_v24_)
                    nw31_ = _dafny.Array(None, 1)
                    for i0_5_ in range(nw31_.length(0)):
                        nw31_[i0_5_] = init5_(i0_5_)
                    d_203_v25_ = nw31_
                    index21_ = default__.safeIndex(509, (d_203_v25_).length(0))
                    (d_203_v25_)[index21_] = _dafny.MultiSet([not(True), d_202_v24_.f12])
                    d_206_v26_: C2
                    nw32_ = C2()
                    nw32_.ctor__(d_202_v24_.f12, default__.fm22(d_171_v1_, not(d_202_v24_.f12), d_181_globalState_), d_179_v6_)
                    d_206_v26_ = nw32_
                    d_207_v27_: _dafny.Seq
                    d_207_v27_ = _dafny.SeqWithoutIsStrInference([d_206_v26_, d_206_v26_, d_206_v26_])
                    d_208_v28_: C0
                    nw33_ = C0()
                    nw33_.ctor__(d_179_v6_, d_171_v1_, _dafny.MultiSet([d_206_v26_.f15, d_202_v24_.f12]), d_202_v24_.f12)
                    d_208_v28_ = nw33_
                    d_209_v29_: _dafny.Set
                    d_209_v29_ = _dafny.Set({(d_208_v28_).f18})
                    d_210_v30_: _dafny.Map
                    d_210_v30_ = _dafny.Map({d_209_v29_: d_208_v28_})
                    d_211_v31_: _dafny.Array
                    nw34_ = _dafny.Array(None, 22)
                    nw34_[int(0)] = d_208_v28_
                    nw34_[int(1)] = d_208_v28_
                    nw34_[int(2)] = d_208_v28_
                    nw34_[int(3)] = d_208_v28_
                    nw34_[int(4)] = d_208_v28_
                    nw34_[int(5)] = d_208_v28_
                    nw34_[int(6)] = d_208_v28_
                    nw34_[int(7)] = d_208_v28_
                    nw34_[int(8)] = ((d_210_v30_)[_dafny.Set({d_171_v1_, (d_208_v28_).f18})] if (_dafny.Set({d_171_v1_, (d_208_v28_).f18})) in (d_210_v30_) else d_208_v28_)
                    nw34_[int(9)] = d_208_v28_
                    nw34_[int(10)] = d_208_v28_
                    nw34_[int(11)] = d_208_v28_
                    nw34_[int(12)] = d_208_v28_
                    nw34_[int(13)] = d_208_v28_
                    nw34_[int(14)] = d_208_v28_
                    nw34_[int(15)] = d_208_v28_
                    nw34_[int(16)] = d_208_v28_
                    nw34_[int(17)] = d_208_v28_
                    nw34_[int(18)] = d_208_v28_
                    nw34_[int(19)] = d_208_v28_
                    nw34_[int(20)] = d_208_v28_
                    nw34_[int(21)] = d_208_v28_
                    d_211_v31_ = nw34_
                    d_212_v32_: _dafny.Map
                    d_212_v32_ = _dafny.Map({d_179_v6_: d_211_v31_})
                    d_213_v33_: _dafny.Map
                    d_213_v33_ = _dafny.Map({D9_DC23(d_212_v32_): d_202_v24_.f12})
                    d_214_v34_: _dafny.Map
                    d_214_v34_ = _dafny.Map({d_206_v26_.f15: d_212_v32_})
                    d_215_v35_: _dafny.Map
                    d_215_v35_ = _dafny.Map({d_206_v26_.f15: d_171_v1_})
                    d_216_v36_: D9
                    d_216_v36_ = D9_DC24(d_200_v22_, (d_208_v28_).f17, d_171_v1_)
                    d_217_v37_: _dafny.Map
                    d_217_v37_ = _dafny.Map({(d_208_v28_).fm2(_dafny.MultiSet([(d_202_v24_.f11).cardinality]), d_171_v1_, d_215_v35_, False, d_181_globalState_): (d_216_v36_).cf57})
                    d_218_v38_: _dafny.MultiSet
                    d_218_v38_ = _dafny.MultiSet([default__.fm16((d_208_v28_).f18, d_181_globalState_), d_217_v37_])
                    d_219_v39_: _dafny.MultiSet
                    d_219_v39_ = _dafny.MultiSet([d_173_v3_, d_173_v3_, d_173_v3_, d_173_v3_, d_173_v3_])
                    d_220_v40_: _dafny.MultiSet
                    d_220_v40_ = _dafny.MultiSet([(d_208_v28_).f18, 80, (d_208_v28_).f18, d_171_v1_, ((d_219_v39_)[d_173_v3_] if (d_173_v3_) in (d_219_v39_) else (d_208_v28_).f18)])
                    index22_ = default__.safeIndex(509, (d_203_v25_).length(0))
                    rhs18_ = ((d_213_v33_)[D9_DC23(((d_214_v34_)[(d_208_v28_).f17] if ((d_208_v28_).f17) in (d_214_v34_) else d_212_v32_))] if (D9_DC23(((d_214_v34_)[(d_208_v28_).f17] if ((d_208_v28_).f17) in (d_214_v34_) else d_212_v32_))) in (d_213_v33_) else (d_179_v6_) == ((d_208_v28_).f17))
                    rhs19_ = ((d_200_v22_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "svi")))) <= (d_200_v22_)
                    rhs20_ = default__.fm22(((d_218_v38_)[_dafny.Map({(d_208_v28_).f17: (d_208_v28_).f17})] if (_dafny.Map({(d_208_v28_).f17: (d_208_v28_).f17})) in (d_218_v38_) else (d_208_v28_).f18), d_202_v24_.f12, d_181_globalState_)
                    rhs21_ = ((_dafny.SeqWithoutIsStrInference([d_206_v26_])) + (d_207_v27_)) + (d_207_v27_)
                    rhs22_ = not(not((d_202_v24_).fm2(d_220_v40_, (d_208_v28_).f18, default__.fm23((d_208_v28_).f18, d_181_globalState_), d_206_v26_.f15, d_181_globalState_)))
                    lhs9_ = d_181_globalState_
                    lhs10_ = d_181_globalState_
                    lhs11_ = d_203_v25_
                    lhs12_ = default__.safeIndex(509, (d_203_v25_).length(0))
                    lhs13_ = d_181_globalState_
                    lhs9_.f3 = rhs18_
                    lhs10_.f7 = rhs19_
                    lhs11_[lhs12_] = rhs20_
                    d_207_v27_ = rhs21_
                    lhs13_.f7 = rhs22_
                    d_221_v41_: D6
                    d_221_v41_ = D6_DC16(d_171_v1_, (d_219_v39_).ispropersubset((d_219_v39_).set(d_173_v3_, default__.abs(((d_202_v24_.f11)[d_206_v26_.f15] if (d_206_v26_.f15) in (d_202_v24_.f11) else d_171_v1_)))), not(d_179_v6_))
                    source2_ = d_221_v41_
                    if source2_.is_DC16:
                        d_222___mcc_h0_ = source2_.cf36
                        d_223___mcc_h1_ = source2_.cf37
                        d_224___mcc_h2_ = source2_.cf38
                        d_225_cf38_ = d_224___mcc_h2_
                        d_226_cf37_ = d_223___mcc_h1_
                        d_227_cf36_ = d_222___mcc_h0_
                        d_171_v1_ = (d_208_v28_).f18
                        d_228_v42_: _dafny.Map
                        d_228_v42_ = _dafny.Map({d_226_cf37_: d_217_v37_})
                        d_229_v43_: _dafny.Map
                        d_229_v43_ = d_217_v37_
                        rhs23_ = d_202_v24_
                        rhs24_ = ((((((d_228_v42_)[(d_208_v28_).f17] if ((d_208_v28_).f17) in (d_228_v42_) else d_217_v37_)).set(d_206_v26_.f15, d_202_v24_.f12)).set(d_226_cf37_, d_202_v24_.f12) if default__.fm0(d_173_v3_, d_181_globalState_) else d_217_v37_)) | ((d_229_v43_))
                        d_202_v24_ = rhs23_
                        d_217_v37_ = rhs24_
                        d_230_v44_: _dafny.Array
                        nw35_ = _dafny.Array(None, 19)
                        nw35_[int(0)] = (d_208_v28_).f17
                        nw35_[int(1)] = (d_208_v28_).f17
                        nw35_[int(2)] = d_225_cf38_
                        nw35_[int(3)] = d_202_v24_.f12
                        nw35_[int(4)] = d_206_v26_.f15
                        nw35_[int(5)] = d_206_v26_.f15
                        nw35_[int(6)] = (d_208_v28_).f17
                        nw35_[int(7)] = d_226_cf37_
                        nw35_[int(8)] = d_206_v26_.f15
                        nw35_[int(9)] = d_206_v26_.f15
                        nw35_[int(10)] = (d_208_v28_).f17
                        nw35_[int(11)] = (d_208_v28_).f17
                        nw35_[int(12)] = d_202_v24_.f12
                        nw35_[int(13)] = (d_208_v28_).f17
                        nw35_[int(14)] = not(d_206_v26_.f15)
                        nw35_[int(15)] = d_202_v24_.f12
                        nw35_[int(16)] = (d_208_v28_).f17
                        nw35_[int(17)] = d_225_cf38_
                        nw35_[int(18)] = (d_206_v26_).fm2(_dafny.MultiSet([(d_208_v28_).f18, d_171_v1_]), (d_208_v28_).f18, d_215_v35_, d_206_v26_.f15, d_181_globalState_)
                        d_230_v44_ = nw35_
                        d_231_v45_: _dafny.Seq
                        d_231_v45_ = _dafny.SeqWithoutIsStrInference([d_230_v44_, d_230_v44_, d_230_v44_, d_230_v44_, d_230_v44_])
                        d_232_v46_: _dafny.Set
                        d_232_v46_ = _dafny.Set({d_230_v44_, d_230_v44_, (d_231_v45_)[default__.safeIndex(849, len(d_231_v45_))]})
                        d_232_v46_ = (d_232_v46_).intersection(d_232_v46_)
                        d_233_v47_: D1
                        d_233_v47_ = D1_DC2(d_225_cf38_, d_206_v26_.f15, (d_208_v28_).f17, d_173_v3_, d_173_v3_)
                        d_234_v48_: _dafny.Map
                        d_234_v48_ = _dafny.Map({(d_208_v28_).f18: default__.fm14(d_233_v47_, (d_202_v24_.f11).cardinality, d_181_globalState_)})
                        d_235_v49_: D1
                        d_235_v49_ = D1_DC2(d_225_cf38_, ((d_217_v37_)[d_202_v24_.f12] if (d_202_v24_.f12) in (d_217_v37_) else True), False, d_173_v3_, ((d_234_v48_)[(d_208_v28_).f18] if ((d_208_v28_).f18) in (d_234_v48_) else d_173_v3_))
                        d_173_v3_ = default__.fm14(d_235_v49_, d_171_v1_, d_181_globalState_)
                    elif True:
                        d_236___mcc_h3_ = source2_.cf35
                        d_237_cf35_ = d_236___mcc_h3_
                        d_238_v50_: bool
                        d_239_v51_: int
                        out0_: bool
                        out1_: int
                        out0_, out1_ = (d_206_v26_).m5(default__.safeModuloInt((d_208_v28_).f18, (d_208_v28_).f18), d_200_v22_, d_181_globalState_)
                        d_238_v50_ = out0_
                        d_239_v51_ = out1_
                        d_200_v22_ = (d_200_v22_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "alxsdoim"))) + (d_200_v22_))
                        d_239_v51_ = len(((d_180_v7_) + (_dafny.SeqWithoutIsStrInference([d_179_v6_, d_206_v26_.f15]))) + ((_dafny.SeqWithoutIsStrInference([d_202_v24_.f12, d_202_v24_.f12, d_206_v26_.f15])) + (d_180_v7_)))
                        (d_181_globalState_).f3 = ((d_208_v28_).f18) != (len(d_200_v22_))
                    d_240_v52_: bool
                    d_241_v53_: int
                    out2_: bool
                    out3_: int
                    out2_, out3_ = (d_206_v26_).m5(default__.fm1(d_206_v26_.f15, d_171_v1_, (d_208_v28_).f18, d_181_globalState_), d_200_v22_, d_181_globalState_)
                    d_240_v52_ = out2_
                    d_241_v53_ = out3_
                    pass
            pass
        d_172_v2_ = (d_172_v2_) + (d_172_v2_)
        (d_181_globalState_).f3 = d_179_v6_
        d_242_v54_: _dafny.Seq
        d_242_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyxck"))
        d_243_v55_: _dafny.Set
        d_243_v55_ = _dafny.Set({d_171_v1_})
        d_244_v56_: D7
        d_244_v56_ = D7_DC18(len(d_180_v7_), d_242_v54_, d_179_v6_, d_179_v6_, d_243_v55_)
        d_245_v57_: _dafny.Seq
        d_245_v57_ = _dafny.SeqWithoutIsStrInference([(d_244_v56_).cf41])
        if (d_171_v1_) <= (len((d_245_v57_)[default__.safeIndex(d_171_v1_, len(d_245_v57_))])):
            d_246_v58_: _dafny.MultiSet
            d_246_v58_ = _dafny.MultiSet([d_179_v6_])
            d_247_v59_: T0
            nw36_ = C2()
            nw36_.ctor__(d_179_v6_, d_246_v58_, d_179_v6_)
            d_247_v59_ = nw36_
            d_248_v60_: _dafny.MultiSet
            d_248_v60_ = _dafny.MultiSet([d_247_v59_, d_247_v59_, d_247_v59_, d_247_v59_, d_247_v59_])
            d_249_v61_: _dafny.Map
            d_249_v61_ = _dafny.Map({d_171_v1_: d_179_v6_})
            d_250_v62_: _dafny.Map
            d_250_v62_ = _dafny.Map({(0) - (d_171_v1_): d_171_v1_})
            d_251_v63_: _dafny.MultiSet
            d_251_v63_ = _dafny.MultiSet([d_171_v1_, len(d_250_v62_), d_171_v1_])
            d_252_v64_: C1
            nw37_ = C1()
            nw37_.ctor__(len(_dafny.Set({d_171_v1_, d_171_v1_, 466, (d_248_v60_).cardinality, len(d_249_v61_)})), _dafny.MultiSet(default__.fm10(d_247_v59_.f12, d_247_v59_.f12, d_179_v6_, d_181_globalState_)), ((d_251_v63_).cardinality) < (default__.fm1(d_179_v6_, len(d_243_v55_), d_171_v1_, d_181_globalState_)))
            d_252_v64_ = nw37_
            d_253_v66_: _dafny.Array
            nw38_ = _dafny.Array(None, 14)
            nw38_[int(0)] = d_242_v54_
            nw38_[int(1)] = d_242_v54_
            def iife32_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in (_dafny.SeqWithoutIsStrInference([d_252_v64_.f16])).Elements:
                    d_254_v65_: int = compr_12_
                    if (d_254_v65_) in (_dafny.SeqWithoutIsStrInference([d_252_v64_.f16])):
                        coll12_[(d_254_v65_) - (d_171_v1_)] = -244
                return _dafny.Map(coll12_)
            nw38_[int(2)] = _dafny.SeqWithoutIsStrInference([(d_242_v54_)[default__.safeIndex(len(iife32_()
), len(d_242_v54_))] for d_255_i6_ in range(default__.abs(60))])
            nw38_[int(3)] = d_242_v54_
            nw38_[int(4)] = d_242_v54_
            nw38_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
            nw38_[int(6)] = (d_242_v54_).set(default__.safeIndex(d_252_v64_.f16, len(d_242_v54_)), d_173_v3_)
            nw38_[int(7)] = d_242_v54_
            nw38_[int(8)] = d_242_v54_
            nw38_[int(9)] = d_242_v54_
            nw38_[int(10)] = d_242_v54_
            nw38_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "foxyq"))
            nw38_[int(12)] = d_242_v54_
            nw38_[int(13)] = d_242_v54_
            d_253_v66_ = nw38_
            d_256_v67_: D5
            d_256_v67_ = D5_DC13(d_252_v64_.f16, d_179_v6_, (0) - (d_252_v64_.f16), d_253_v66_, d_247_v59_.f12)
            d_171_v1_ = (d_252_v64_.f16) * ((d_256_v67_).cf30)
            rhs25_ = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qtkfwgn")))) - (((d_251_v63_)[len(d_249_v61_)] if (len(d_249_v61_)) in (d_251_v63_) else d_171_v1_))
            rhs26_ = d_171_v1_
            lhs14_ = d_252_v64_
            lhs14_.f16 = rhs25_
            d_171_v1_ = rhs26_
            d_257_v68_: D9
            d_257_v68_ = D9_DC25((d_242_v54_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ak"))), d_247_v59_.f12, d_252_v64_.f16)
            source3_ = d_257_v68_
            if source3_.is_DC24:
                d_258___mcc_h4_ = source3_.cf56
                d_259___mcc_h5_ = source3_.cf57
                d_260___mcc_h6_ = source3_.cf58
                d_261_cf58_ = d_260___mcc_h6_
                d_262_cf57_ = d_259___mcc_h5_
                d_263_cf56_ = d_258___mcc_h4_
                d_264_v69_: _dafny.Map
                d_264_v69_ = _dafny.Map({d_247_v59_.f12: False})
                d_264_v69_ = (d_264_v69_).set((len(_dafny.Set({d_262_cf57_, d_247_v59_.f12}))) > ((0) - (-496)), (d_247_v59_.f12 if d_247_v59_.f12 else d_247_v59_.f12))
                d_265_v70_: _dafny.Map
                d_265_v70_ = _dafny.Map({d_247_v59_.f12: d_252_v64_.f16})
                d_266_v71_: _dafny.Map
                d_266_v71_ = _dafny.Map({d_252_v64_.f16: d_265_v70_})
                d_266_v71_ = (d_266_v71_).set(12, d_265_v70_)
                index23_ = default__.safeIndex(222, (d_176_v5_).length(0))
                (d_176_v5_)[index23_] = len((d_263_cf56_) + (d_263_cf56_))
                index24_ = default__.safeIndex(222, (d_176_v5_).length(0))
                (d_176_v5_)[index24_] = d_252_v64_.f16
                d_173_v3_ = d_173_v3_
            elif source3_.is_DC25:
                d_267___mcc_h7_ = source3_.cf59
                d_268___mcc_h8_ = source3_.cf60
                d_269___mcc_h9_ = source3_.cf61
                d_270_cf61_ = d_269___mcc_h9_
                d_271_cf60_ = d_268___mcc_h8_
                d_272_cf59_ = d_267___mcc_h7_
                d_273_v72_: _dafny.Array
                nw39_ = _dafny.Array(_dafny.Seq({}), 12)
                d_273_v72_ = nw39_
                d_274_v73_: _dafny.Map
                d_274_v73_ = _dafny.Map({((d_249_v61_)[d_252_v64_.f16] if (d_252_v64_.f16) in (d_249_v61_) else d_247_v59_.f12): d_273_v72_})
                d_274_v73_ = (d_274_v73_).set(d_247_v59_.f12, d_273_v72_)
                d_275_v74_: _dafny.Map
                d_275_v74_ = _dafny.Map({d_247_v59_.f12: d_171_v1_})
                d_276_v75_: _dafny.Map
                d_276_v75_ = _dafny.Map({((d_250_v62_)[d_252_v64_.f16] if (d_252_v64_.f16) in (d_250_v62_) else len(d_275_v74_)): _dafny.SeqWithoutIsStrInference([d_271_cf60_])})
                d_277_v76_: C3
                nw40_ = C3()
                nw40_.ctor__(len(((d_276_v75_)[679] if (679) in (d_276_v75_) else d_180_v7_)), d_179_v6_, default__.fm22(d_252_v64_.f16, d_247_v59_.f12, d_181_globalState_), d_247_v59_.f12)
                d_277_v76_ = nw40_
                d_278_v77_: _dafny.Map
                d_278_v77_ = _dafny.Map({d_277_v76_: (len(_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_279_i7_ in range(default__.abs(283))]))) - (d_252_v64_.f16)})
                (d_181_globalState_).f9 = len(d_278_v77_)
                (d_181_globalState_).f8 = d_247_v59_.f12
                d_280_v78_: D6
                d_280_v78_ = D6_DC16(d_252_v64_.f16, d_271_cf60_, d_179_v6_)
                (d_247_v59_).f12 = (d_280_v78_).cf38
            elif source3_.is_DC26:
                d_281___mcc_h10_ = source3_.cf62
                d_282_cf62_ = d_281___mcc_h10_
                (d_181_globalState_).f8 = d_247_v59_.f12
                d_283_v79_: C0
                nw41_ = C0()
                nw41_.ctor__(not(d_247_v59_.f12), default__.safeModuloInt(d_252_v64_.f16, d_252_v64_.f16), d_246_v58_, not(d_179_v6_))
                d_283_v79_ = nw41_
                d_282_cf62_ = -697
                d_284_v80_: _dafny.Array
                nw42_ = _dafny.Array(None, 4)
                d_284_v80_ = nw42_
                d_285_v81_: _dafny.Seq
                d_285_v81_ = _dafny.SeqWithoutIsStrInference([d_247_v59_.f11])
                d_286_v82_: C3
                nw43_ = C3()
                nw43_.ctor__(d_252_v64_.f16, d_179_v6_, (d_285_v81_)[default__.safeIndex(d_171_v1_, len(d_285_v81_))], d_247_v59_.f12)
                d_286_v82_ = nw43_
                index25_ = default__.safeIndex(648, (d_284_v80_).length(0))
                (d_284_v80_)[index25_] = d_286_v82_
                index26_ = default__.safeIndex(648, (d_284_v80_).length(0))
                (d_284_v80_)[index26_] = d_286_v82_
            elif True:
                d_287___mcc_h11_ = source3_.cf55
                d_288_cf55_ = d_287___mcc_h11_
                d_289_v83_: C4
                nw44_ = C4()
                nw44_.ctor__(d_247_v59_.f11, d_247_v59_.f12)
                d_289_v83_ = nw44_
                d_290_v84_: _dafny.Array
                nw45_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_290_v84_ = nw45_
                d_291_v85_: _dafny.Map
                d_291_v85_ = _dafny.Map({d_179_v6_: d_290_v84_})
                d_292_v86_: _dafny.Seq
                d_292_v86_ = _dafny.SeqWithoutIsStrInference([d_246_v58_])
                d_293_v87_: _dafny.Map
                d_293_v87_ = _dafny.Map({((d_291_v85_)[d_179_v6_] if (d_179_v6_) in (d_291_v85_) else d_290_v84_): d_292_v86_})
                d_293_v87_ = (d_293_v87_).set(d_290_v84_, d_292_v86_)
                d_294_v88_: D8
                d_294_v88_ = D8_DC20(d_247_v59_.f12, d_247_v59_.f12, False)
                d_294_v88_ = D8_DC20(d_247_v59_.f12, d_179_v6_, (d_289_v83_).fm4(d_252_v64_.f16, d_252_v64_.f16, d_181_globalState_))
                d_295_v89_: _dafny.Set
                d_295_v89_ = _dafny.Set({d_247_v59_.f12, True, d_247_v59_.f12, d_179_v6_, not(d_247_v59_.f12)})
                d_296_v90_: _dafny.Array
                d_297_v91_: _dafny.Array
                d_298_v92_: int
                out4_: _dafny.Array
                out5_: _dafny.Array
                out6_: int
                out4_, out5_, out6_ = (d_289_v83_).m1(d_295_v89_, d_181_globalState_)
                d_296_v90_ = out4_
                d_297_v91_ = out5_
                d_298_v92_ = out6_
            rhs27_ = d_252_v64_.f16
            rhs28_ = (d_252_v64_.f16 if (d_171_v1_) >= (d_252_v64_.f16) else d_171_v1_)
            rhs29_ = d_171_v1_
            rhs30_ = d_179_v6_
            rhs31_ = d_179_v6_
            lhs15_ = d_252_v64_
            lhs16_ = d_181_globalState_
            lhs17_ = d_181_globalState_
            lhs18_ = d_181_globalState_
            lhs19_ = d_181_globalState_
            lhs15_.f16 = rhs27_
            lhs16_.f9 = rhs28_
            lhs17_.f9 = rhs29_
            lhs18_.f3 = rhs30_
            lhs19_.f8 = rhs31_
        elif True:
            d_299_v93_: _dafny.MultiSet
            d_299_v93_ = _dafny.MultiSet([not(d_179_v6_), False, d_179_v6_])
            d_300_v94_: _dafny.Map
            d_300_v94_ = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))) != (d_299_v93_): d_172_v2_})
            d_300_v94_ = (d_300_v94_).set((d_180_v7_)[default__.safeIndex(d_171_v1_, len(d_180_v7_))], _dafny.SeqWithoutIsStrInference([d_171_v1_ for d_301_i8_ in range(default__.abs(226))]))
            (d_181_globalState_).f8 = (d_171_v1_) > (d_171_v1_)
            index27_ = default__.safeIndex(537, (d_176_v5_).length(0))
            (d_176_v5_)[index27_] = (d_171_v1_) * (d_171_v1_)
            index28_ = default__.safeIndex(537, (d_176_v5_).length(0))
            rhs32_ = (d_171_v1_ if d_179_v6_ else d_171_v1_)
            rhs33_ = (d_172_v2_)[default__.safeIndex(927, len(d_172_v2_))]
            lhs20_ = d_181_globalState_
            lhs21_ = d_176_v5_
            lhs22_ = default__.safeIndex(537, (d_176_v5_).length(0))
            lhs20_.f9 = rhs32_
            lhs21_[lhs22_] = rhs33_
            default__.m0(d_181_globalState_)
            d_302_v95_: _dafny.Map
            d_302_v95_ = _dafny.Map({True: (d_176_v5_)[default__.safeIndex(537, (d_176_v5_).length(0))]})
            d_303_v96_: _dafny.Map
            d_303_v96_ = _dafny.Map({(d_176_v5_)[default__.safeIndex(537, (d_176_v5_).length(0))]: (d_176_v5_)[default__.safeIndex(537, (d_176_v5_).length(0))]})
            (d_181_globalState_).f9 = default__.safeDivisionInt(((d_302_v95_)[not(True)] if (not(True)) in (d_302_v95_) else default__.fm1(d_179_v6_, (d_176_v5_)[default__.safeIndex(537, (d_176_v5_).length(0))], (0) - (len(d_303_v96_)), d_181_globalState_)), (d_176_v5_)[default__.safeIndex(537, (d_176_v5_).length(0))])
        (d_181_globalState_).f3 = (len((d_242_v54_ if d_179_v6_ else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gul"))).set(default__.safeIndex(d_171_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gul")))), _dafny.CodePoint('d'))))) == (d_171_v1_)
        if d_179_v6_:
            d_304_v97_: _dafny.Set
            d_304_v97_ = _dafny.Set({d_179_v6_, d_179_v6_})
            if (d_304_v97_).ispropersubset(d_304_v97_):
                d_305_v98_: _dafny.Map
                d_305_v98_ = _dafny.Map({d_179_v6_: default__.fm0(d_173_v3_, d_181_globalState_)})
                d_306_v99_: _dafny.Map
                d_306_v99_ = _dafny.Map({d_172_v2_: d_305_v98_})
                d_307_v100_: D10
                d_307_v100_ = D10_DC29(d_179_v6_)
                pat_let_tv11_ = d_179_v6_
                d_308_v101_: _dafny.Set
                def iife33_(_pat_let10_0):
                    def iife34_(d_309_dt__update__tmp_h0_):
                        def iife35_(_pat_let11_0):
                            def iife36_(d_310_dt__update_hcf68_h0_):
                                return D10_DC29(d_310_dt__update_hcf68_h0_)
                            return iife36_(_pat_let11_0)
                        return iife35_(pat_let_tv11_)
                    return iife34_(_pat_let10_0)
                d_308_v101_ = _dafny.Set({d_304_v97_, default__.fm18(len((d_306_v99_).set(_dafny.SeqWithoutIsStrInference([len(d_245_v57_)]), d_305_v98_)), (iife33_(d_307_v100_)).cf68, d_181_globalState_), _dafny.Set({d_179_v6_, d_179_v6_})})
                rhs34_ = not(default__.fm0(d_173_v3_, d_181_globalState_))
                rhs35_ = (d_171_v1_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vomrh"))))
                rhs36_ = d_308_v101_
                rhs37_ = d_179_v6_
                rhs38_ = (d_171_v1_) > ((-312) * (d_171_v1_))
                lhs23_ = d_181_globalState_
                lhs24_ = d_181_globalState_
                lhs25_ = d_181_globalState_
                lhs23_.f3 = rhs34_
                d_171_v1_ = rhs35_
                d_308_v101_ = rhs36_
                lhs24_.f7 = rhs37_
                lhs25_.f7 = rhs38_
                d_311_v102_: _dafny.Array
                def lambda12_(d_312_v55_):
                    def lambda13_(d_313_i9_):
                        return d_312_v55_

                    return lambda13_

                init6_ = lambda12_(d_243_v55_)
                nw46_ = _dafny.Array(None, 20)
                for i0_6_ in range(nw46_.length(0)):
                    nw46_[i0_6_] = init6_(i0_6_)
                d_311_v102_ = nw46_
                index29_ = default__.safeIndex(838, (d_311_v102_).length(0))
                (d_311_v102_)[index29_] = d_243_v55_
                index30_ = default__.safeIndex(838, (d_311_v102_).length(0))
                (d_311_v102_)[index30_] = d_243_v55_
                d_314_v103_: _dafny.Map
                d_314_v103_ = _dafny.Map({d_171_v1_: d_179_v6_})
                d_315_v104_: _dafny.Map
                d_315_v104_ = _dafny.Map({d_314_v103_: d_171_v1_})
                (d_181_globalState_).f9 = len((d_315_v104_) | (_dafny.Map({_dafny.Map({d_171_v1_: d_179_v6_}): (0) - (len((d_311_v102_)[default__.safeIndex(838, (d_311_v102_).length(0))]))})))
                d_171_v1_ = d_171_v1_
                (d_181_globalState_).f7 = d_179_v6_
            elif True:
                d_316_v105_: _dafny.Map
                d_316_v105_ = _dafny.Map({d_304_v97_: d_171_v1_})
                index31_ = default__.safeIndex(579, (d_176_v5_).length(0))
                (d_176_v5_)[index31_] = ((d_316_v105_)[d_304_v97_] if (d_304_v97_) in (d_316_v105_) else d_171_v1_)
                index32_ = default__.safeIndex(76, (d_176_v5_).length(0))
                (d_176_v5_)[index32_] = d_171_v1_
                d_317_v106_: C4
                nw47_ = C4()
                nw47_.ctor__(_dafny.MultiSet([d_179_v6_, d_179_v6_]), d_179_v6_)
                d_317_v106_ = nw47_
                d_318_v107_: _dafny.MultiSet
                d_318_v107_ = _dafny.MultiSet([d_317_v106_, d_317_v106_])
                index33_ = default__.safeIndex(579, (d_176_v5_).length(0))
                index34_ = default__.safeIndex(76, (d_176_v5_).length(0))
                rhs39_ = d_179_v6_
                rhs40_ = d_171_v1_
                rhs41_ = default__.safeDivisionInt(574, d_171_v1_)
                rhs42_ = default__.safeDivisionInt(d_171_v1_, -498)
                rhs43_ = len((_dafny.SeqWithoutIsStrInference([d_171_v1_ for d_319_i10_ in range(default__.abs(543))]) if (_dafny.Map({d_318_v107_: d_179_v6_})) == (_dafny.Map({d_318_v107_: d_179_v6_})) else _dafny.SeqWithoutIsStrInference([d_171_v1_ for d_320_i11_ in range(default__.abs(817))])))
                lhs26_ = d_181_globalState_
                lhs27_ = d_176_v5_
                lhs28_ = default__.safeIndex(579, (d_176_v5_).length(0))
                lhs29_ = d_176_v5_
                lhs30_ = default__.safeIndex(76, (d_176_v5_).length(0))
                lhs31_ = d_181_globalState_
                lhs32_ = d_181_globalState_
                lhs26_.f3 = rhs39_
                lhs27_[lhs28_] = rhs40_
                lhs29_[lhs30_] = rhs41_
                lhs31_.f9 = rhs42_
                lhs32_.f9 = rhs43_
                (d_181_globalState_).f9 = 760
                d_321_v108_: _dafny.MultiSet
                d_321_v108_ = _dafny.MultiSet([d_179_v6_, d_179_v6_, False, d_179_v6_])
                d_322_v109_: _dafny.Map
                d_322_v109_ = _dafny.Map({d_171_v1_: d_321_v108_})
                d_323_v110_: D1
                d_323_v110_ = D1_DC2(True, d_179_v6_, d_179_v6_, d_173_v3_, d_173_v3_)
                pat_let_tv12_ = d_173_v3_
                d_324_v111_: C0
                nw48_ = C0()
                def iife37_(_pat_let12_0):
                    def iife38_(d_325_dt__update__tmp_h1_):
                        def iife39_(_pat_let13_0):
                            def iife40_(d_326_dt__update_hcf6_h0_):
                                def iife41_(_pat_let14_0):
                                    def iife42_(d_327_dt__update_hcf2_h0_):
                                        return D1_DC2(d_327_dt__update_hcf2_h0_, (d_325_dt__update__tmp_h1_).cf3, (d_325_dt__update__tmp_h1_).cf4, (d_325_dt__update__tmp_h1_).cf5, d_326_dt__update_hcf6_h0_)
                                    return iife42_(_pat_let14_0)
                                return iife41_(False)
                            return iife40_(_pat_let13_0)
                        return iife39_(pat_let_tv12_)
                    return iife38_(_pat_let12_0)
                nw48_.ctor__(d_179_v6_, d_171_v1_, ((d_322_v109_)[(d_176_v5_)[default__.safeIndex(579, (d_176_v5_).length(0))]] if ((d_176_v5_)[default__.safeIndex(579, (d_176_v5_).length(0))]) in (d_322_v109_) else _dafny.MultiSet([d_179_v6_, d_179_v6_, (iife37_(d_323_v110_)).cf4])), d_179_v6_)
                d_324_v111_ = nw48_
                d_328_v112_: _dafny.Map
                d_328_v112_ = _dafny.Map({d_324_v111_: d_173_v3_})
                d_328_v112_ = (d_328_v112_).set(d_324_v111_, _dafny.CodePoint('d'))
                index35_ = default__.safeIndex(579, (d_176_v5_).length(0))
                index36_ = default__.safeIndex(579, (d_176_v5_).length(0))
                rhs44_ = ((d_324_v111_).f18) * (((d_324_v111_).f18) * (55))
                rhs45_ = default__.safeDivisionInt((d_176_v5_)[default__.safeIndex(579, (d_176_v5_).length(0))], ((d_321_v108_)[d_179_v6_] if (d_179_v6_) in (d_321_v108_) else 542))
                rhs46_ = (0) - ((d_324_v111_).f18)
                lhs33_ = d_176_v5_
                lhs34_ = default__.safeIndex(579, (d_176_v5_).length(0))
                lhs35_ = d_176_v5_
                lhs36_ = default__.safeIndex(579, (d_176_v5_).length(0))
                lhs37_ = d_181_globalState_
                lhs33_[lhs34_] = rhs44_
                lhs35_[lhs36_] = rhs45_
                lhs37_.f9 = rhs46_
                d_329_v113_: _dafny.Array
                d_330_v114_: _dafny.Array
                d_331_v115_: int
                out7_: _dafny.Array
                out8_: _dafny.Array
                out9_: int
                out7_, out8_, out9_ = (d_317_v106_).m1(default__.fm18(269, (d_324_v111_).f17, d_181_globalState_), d_181_globalState_)
                d_329_v113_ = out7_
                d_330_v114_ = out8_
                d_331_v115_ = out9_
            (d_181_globalState_).f3 = (d_171_v1_) != (default__.safeDivisionInt(d_171_v1_, d_171_v1_))
            d_332_v116_: _dafny.Array
            nw49_ = _dafny.Array(None, 9)
            d_332_v116_ = nw49_
            d_333_v117_: _dafny.Seq
            d_333_v117_ = _dafny.SeqWithoutIsStrInference([d_332_v116_])
            d_334_v118_: _dafny.Array
            nw50_ = _dafny.Array(None, 14)
            nw50_[int(0)] = d_332_v116_
            nw50_[int(1)] = d_332_v116_
            nw50_[int(2)] = d_332_v116_
            nw50_[int(3)] = d_332_v116_
            nw50_[int(4)] = d_332_v116_
            nw50_[int(5)] = (d_333_v117_)[default__.safeIndex((0) - (d_171_v1_), len(d_333_v117_))]
            nw50_[int(6)] = d_332_v116_
            nw50_[int(7)] = d_332_v116_
            nw50_[int(8)] = d_332_v116_
            nw50_[int(9)] = d_332_v116_
            nw50_[int(10)] = d_332_v116_
            nw50_[int(11)] = d_332_v116_
            nw50_[int(12)] = d_332_v116_
            nw50_[int(13)] = d_332_v116_
            d_334_v118_ = nw50_
            index37_ = default__.safeIndex(802, (d_334_v118_).length(0))
            (d_334_v118_)[index37_] = (d_332_v116_ if d_179_v6_ else d_332_v116_)
            index38_ = default__.safeIndex(802, (d_334_v118_).length(0))
            (d_334_v118_)[index38_] = d_332_v116_
            d_335_v119_: _dafny.Map
            d_335_v119_ = _dafny.Map({d_173_v3_: d_171_v1_})
            d_336_v120_: _dafny.Set
            d_336_v120_ = _dafny.Set({d_335_v119_})
            d_337_v121_: _dafny.Map
            d_337_v121_ = _dafny.Map({(d_336_v120_).intersection(d_336_v120_): (d_243_v55_) - (d_243_v55_)})
            d_243_v55_ = ((d_337_v121_)[_dafny.Set({d_335_v119_, d_335_v119_, _dafny.Map({d_173_v3_: d_171_v1_})})] if (_dafny.Set({d_335_v119_, d_335_v119_, _dafny.Map({d_173_v3_: d_171_v1_})})) in (d_337_v121_) else ((d_244_v56_).cf44) - (d_243_v55_))
            if default__.fm0(d_173_v3_, d_181_globalState_):
                (d_181_globalState_).f9 = 678
                d_338_v122_: _dafny.Array
                nw51_ = _dafny.Array(False, 28)
                d_338_v122_ = nw51_
                d_338_v122_ = d_338_v122_
                d_339_v123_: _dafny.Array
                nw52_ = _dafny.Array(D8.default()(), 29)
                d_339_v123_ = nw52_
                d_340_v124_: _dafny.Array
                nw53_ = _dafny.Array(None, 8)
                nw53_[int(0)] = d_339_v123_
                nw53_[int(1)] = d_339_v123_
                nw53_[int(2)] = d_339_v123_
                nw53_[int(3)] = d_339_v123_
                nw53_[int(4)] = d_339_v123_
                nw53_[int(5)] = d_339_v123_
                nw53_[int(6)] = d_339_v123_
                nw53_[int(7)] = d_339_v123_
                d_340_v124_ = nw53_
                index39_ = default__.safeIndex(301, (d_340_v124_).length(0))
                (d_340_v124_)[index39_] = d_339_v123_
                index40_ = default__.safeIndex(301, (d_340_v124_).length(0))
                (d_340_v124_)[index40_] = d_339_v123_
                d_341_v125_: _dafny.Map
                d_341_v125_ = _dafny.Map({d_173_v3_: d_179_v6_})
                d_341_v125_ = (d_341_v125_).set(d_173_v3_, d_179_v6_)
                d_342_v126_: _dafny.Array
                nw54_ = _dafny.Array(_dafny.Array(None, 0), 8)
                d_342_v126_ = nw54_
                d_343_v127_: D8
                d_343_v127_ = D8_DC21(d_179_v6_, d_342_v126_, d_171_v1_, d_338_v122_, 325)
                d_344_v128_: _dafny.MultiSet
                d_344_v128_ = _dafny.MultiSet([not (default__.fm0(d_173_v3_, d_181_globalState_)) or (not((d_343_v127_).cf49)), d_179_v6_, d_179_v6_, d_179_v6_])
                d_344_v128_ = ((default__.fm22(d_171_v1_, d_179_v6_, d_181_globalState_)) | (d_344_v128_)).intersection((d_344_v128_).intersection(d_344_v128_))
            elif True:
                (d_181_globalState_).f3 = d_179_v6_
                (d_181_globalState_).f9 = -58
                default__.m0(d_181_globalState_)
                d_345_v129_: _dafny.Array
                nw55_ = _dafny.Array(None, 16)
                d_345_v129_ = nw55_
                d_346_v130_: _dafny.Map
                d_346_v130_ = _dafny.Map({d_345_v129_: d_179_v6_})
                d_346_v130_ = (d_346_v130_).set(d_345_v129_, d_179_v6_)
                d_172_v2_ = _dafny.SeqWithoutIsStrInference([(d_171_v1_) * (d_171_v1_), 975, d_171_v1_])
        elif True:
            (d_181_globalState_).f7 = d_179_v6_
            d_347_v131_: _dafny.MultiSet
            d_347_v131_ = _dafny.MultiSet([d_179_v6_])
            d_348_v132_: C4
            nw56_ = C4()
            nw56_.ctor__(d_347_v131_, d_179_v6_)
            d_348_v132_ = nw56_
            d_349_v133_: T0
            nw57_ = C0()
            nw57_.ctor__(d_179_v6_, d_171_v1_, (d_347_v131_).set(d_179_v6_, default__.abs(d_171_v1_)), d_179_v6_)
            d_349_v133_ = nw57_
            d_349_v133_ = d_349_v133_
            d_350_v134_: D8
            d_350_v134_ = D8_DC20(d_179_v6_, d_179_v6_, d_349_v133_.f12)
            d_351_v135_: C0
            nw58_ = C0()
            nw58_.ctor__(not(d_349_v133_.f12), d_171_v1_, d_349_v133_.f11, (d_350_v134_).cf48)
            d_351_v135_ = nw58_
            d_352_v136_: _dafny.Map
            d_352_v136_ = _dafny.Map({d_179_v6_: d_171_v1_})
            d_353_v137_: _dafny.Set
            d_353_v137_ = _dafny.Set({d_352_v136_})
            d_354_v138_: _dafny.Set
            d_354_v138_ = _dafny.Set({True, (d_351_v135_).f17})
            d_355_v139_: C0
            nw59_ = C0()
            nw59_.ctor__(d_179_v6_, len(d_353_v137_), d_349_v133_.f11, ((d_351_v135_).f18) >= (len(d_354_v138_)))
            d_355_v139_ = nw59_
        d_356_v140_: _dafny.Map
        d_356_v140_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muwivve"))).set(default__.safeIndex(d_171_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muwivve")))), d_173_v3_): (default__.fm26(d_171_v1_, not(False), d_181_globalState_)).cf46})
        def iife43_():
            coll13_ = _dafny.Set()
            compr_13_: int
            for compr_13_ in _dafny.IntegerRange(567, 942):
                d_359_v141_: int = compr_13_
                if ((567) <= (d_359_v141_)) and ((d_359_v141_) < (942)):
                    coll13_ = coll13_.union(_dafny.Set([(d_359_v141_) + (d_171_v1_)]))
            return _dafny.Set(coll13_)
        if ((d_356_v140_)[_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_357_i12_ in range(default__.abs(820))])] if (_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_358_i12_ in range(default__.abs(820))])) in (d_356_v140_) else (d_243_v55_).ispropersubset(iife43_()
        )):
            d_360_v142_: _dafny.Array
            def lambda14_(d_361_v6_):
                def lambda15_(d_362_i13_):
                    return d_361_v6_

                return lambda15_

            init7_ = lambda14_(d_179_v6_)
            nw60_ = _dafny.Array(None, 7)
            for i0_7_ in range(nw60_.length(0)):
                nw60_[i0_7_] = init7_(i0_7_)
            d_360_v142_ = nw60_
            d_363_v143_: _dafny.Seq
            d_363_v143_ = _dafny.SeqWithoutIsStrInference([d_360_v142_])
            d_364_v144_: _dafny.Array
            nw61_ = _dafny.Array(None, 23)
            nw61_[int(0)] = d_360_v142_
            nw61_[int(1)] = d_360_v142_
            nw61_[int(2)] = d_360_v142_
            nw61_[int(3)] = (d_363_v143_)[default__.safeIndex(default__.fm1(d_179_v6_, default__.fm1(d_179_v6_, d_171_v1_, d_171_v1_, d_181_globalState_), len(_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_365_i14_ in range(default__.abs(810))])), d_181_globalState_), len(d_363_v143_))]
            nw61_[int(4)] = d_360_v142_
            nw61_[int(5)] = d_360_v142_
            nw61_[int(6)] = d_360_v142_
            nw61_[int(7)] = d_360_v142_
            nw61_[int(8)] = d_360_v142_
            nw61_[int(9)] = d_360_v142_
            nw61_[int(10)] = d_360_v142_
            nw61_[int(11)] = d_360_v142_
            nw61_[int(12)] = d_360_v142_
            nw61_[int(13)] = d_360_v142_
            nw61_[int(14)] = d_360_v142_
            nw61_[int(15)] = d_360_v142_
            nw61_[int(16)] = d_360_v142_
            nw61_[int(17)] = d_360_v142_
            nw61_[int(18)] = d_360_v142_
            nw61_[int(19)] = d_360_v142_
            nw61_[int(20)] = d_360_v142_
            nw61_[int(21)] = d_360_v142_
            nw61_[int(22)] = d_360_v142_
            d_364_v144_ = nw61_
            d_366_v145_: D8
            d_366_v145_ = D8_DC21(d_179_v6_, d_364_v144_, d_171_v1_, d_360_v142_, default__.fm1(d_179_v6_, 759, d_171_v1_, d_181_globalState_))
            d_367_v146_: _dafny.MultiSet
            d_367_v146_ = _dafny.MultiSet([d_179_v6_, (d_366_v145_).cf49])
            d_368_v147_: C2
            nw62_ = C2()
            nw62_.ctor__(True, d_367_v146_, not(d_179_v6_))
            d_368_v147_ = nw62_
            d_369_v148_: _dafny.Map
            d_369_v148_ = _dafny.Map({d_368_v147_: d_360_v142_})
            pat_let_tv13_ = d_364_v144_
            d_370_v149_: _dafny.Array
            nw63_ = _dafny.Array(None, 23)
            nw63_[int(0)] = d_360_v142_
            nw63_[int(1)] = d_360_v142_
            nw63_[int(2)] = d_360_v142_
            nw63_[int(3)] = d_360_v142_
            nw63_[int(4)] = ((d_369_v148_)[d_368_v147_] if (d_368_v147_) in (d_369_v148_) else d_360_v142_)
            nw63_[int(5)] = d_360_v142_
            nw63_[int(6)] = d_360_v142_
            nw63_[int(7)] = d_360_v142_
            def iife44_(_pat_let15_0):
                def iife45_(d_371_dt__update__tmp_h2_):
                    def iife46_(_pat_let16_0):
                        def iife47_(d_372_dt__update_hcf50_h0_):
                            return D8_DC21((d_371_dt__update__tmp_h2_).cf49, d_372_dt__update_hcf50_h0_, (d_371_dt__update__tmp_h2_).cf51, (d_371_dt__update__tmp_h2_).cf52, (d_371_dt__update__tmp_h2_).cf53)
                        return iife47_(_pat_let16_0)
                    return iife46_(pat_let_tv13_)
                return iife45_(_pat_let15_0)
            nw63_[int(8)] = (iife44_(d_366_v145_)).cf52
            nw63_[int(9)] = d_360_v142_
            nw63_[int(10)] = d_360_v142_
            nw63_[int(11)] = d_360_v142_
            nw63_[int(12)] = d_360_v142_
            nw63_[int(13)] = d_360_v142_
            nw63_[int(14)] = d_360_v142_
            nw63_[int(15)] = d_360_v142_
            nw63_[int(16)] = d_360_v142_
            nw63_[int(17)] = d_360_v142_
            nw63_[int(18)] = d_360_v142_
            nw63_[int(19)] = d_360_v142_
            nw63_[int(20)] = d_360_v142_
            nw63_[int(21)] = d_360_v142_
            nw63_[int(22)] = d_360_v142_
            d_370_v149_ = nw63_
            index41_ = default__.safeIndex(586, (d_364_v144_).length(0))
            (d_364_v144_)[index41_] = d_360_v142_
            d_373_v150_: _dafny.Map
            d_373_v150_ = _dafny.Map({d_171_v1_: d_368_v147_.f15})
            d_374_v151_: _dafny.Map
            d_374_v151_ = _dafny.Map({((d_373_v150_)[d_171_v1_] if (d_171_v1_) in (d_373_v150_) else d_368_v147_.f15): d_360_v142_})
            d_375_v152_: _dafny.MultiSet
            d_375_v152_ = _dafny.MultiSet([d_171_v1_, d_171_v1_, len(d_180_v7_)])
            d_376_v153_: _dafny.Map
            d_376_v153_ = _dafny.Map({d_179_v6_: d_171_v1_})
            index42_ = default__.safeIndex(586, (d_364_v144_).length(0))
            (d_364_v144_)[index42_] = ((d_374_v151_)[(d_368_v147_).fm2(d_375_v152_, d_171_v1_, d_376_v153_, not(d_179_v6_), d_181_globalState_)] if ((d_368_v147_).fm2(d_375_v152_, d_171_v1_, d_376_v153_, not(d_179_v6_), d_181_globalState_)) in (d_374_v151_) else d_360_v142_)
            arr0_ = (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]
            index43_ = default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))
            arr0_[index43_] = d_368_v147_.f15
            arr1_ = (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]
            index44_ = default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))
            arr1_[index44_] = d_179_v6_
            d_377_v154_: C0
            nw64_ = C0()
            nw64_.ctor__(((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))])[default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))], d_171_v1_, d_367_v146_, d_179_v6_)
            d_377_v154_ = nw64_
            if d_368_v147_.f15:
                d_179_v6_ = ((len(d_376_v153_)) * ((d_377_v154_).f18)) >= ((d_377_v154_).f18)
                (d_181_globalState_).f8 = ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))])[default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))]
                d_378_v155_: _dafny.Set
                d_378_v155_ = _dafny.Set({d_173_v3_, d_173_v3_})
                d_379_v156_: _dafny.Array
                nw65_ = _dafny.Array(None, 27)
                nw65_[int(0)] = d_242_v54_
                nw65_[int(1)] = _dafny.SeqWithoutIsStrInference([d_173_v3_ for d_380_i15_ in range(default__.abs(-298))])
                nw65_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ecbkysbx"))
                nw65_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mekb"))
                nw65_[int(4)] = d_242_v54_
                nw65_[int(5)] = d_242_v54_
                nw65_[int(6)] = d_242_v54_
                nw65_[int(7)] = d_242_v54_
                nw65_[int(8)] = d_242_v54_
                nw65_[int(9)] = d_242_v54_
                nw65_[int(10)] = d_242_v54_
                nw65_[int(11)] = d_242_v54_
                nw65_[int(12)] = d_242_v54_
                nw65_[int(13)] = d_242_v54_
                nw65_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnflkujmg"))
                nw65_[int(15)] = d_242_v54_
                nw65_[int(16)] = d_242_v54_
                nw65_[int(17)] = d_242_v54_
                nw65_[int(18)] = default__.fm9(len(d_378_v155_), True, d_181_globalState_)
                nw65_[int(19)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_381_i16_ in range(default__.abs(432))])
                nw65_[int(20)] = d_242_v54_
                nw65_[int(21)] = d_242_v54_
                nw65_[int(22)] = d_242_v54_
                nw65_[int(23)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))
                nw65_[int(24)] = _dafny.SeqWithoutIsStrInference([d_173_v3_ for d_382_i17_ in range(default__.abs(325))])
                nw65_[int(25)] = d_242_v54_
                nw65_[int(26)] = d_242_v54_
                d_379_v156_ = nw65_
                d_383_v157_: D5
                d_383_v157_ = D5_DC13(len(d_376_v153_), True, d_171_v1_, d_379_v156_, (d_377_v154_).f17)
                d_384_v158_: _dafny.Set
                d_384_v158_ = _dafny.Set({d_383_v157_, d_383_v157_})
                d_385_v159_: _dafny.Seq
                d_385_v159_ = _dafny.SeqWithoutIsStrInference([d_384_v158_])
                d_386_v160_: D12
                d_386_v160_ = D12_DC31(d_385_v159_)
                d_387_v161_: _dafny.Seq
                d_387_v161_ = _dafny.SeqWithoutIsStrInference([d_385_v159_, _dafny.SeqWithoutIsStrInference([d_384_v158_])])
                (d_368_v147_).f15 = ((d_386_v160_).cf70) == ((d_387_v161_)[default__.safeIndex((d_377_v154_).f18, len(d_387_v161_))])
                d_388_v162_: C1
                nw66_ = C1()
                nw66_.ctor__((0) - ((d_377_v154_).f18), _dafny.MultiSet([d_179_v6_]), ((d_377_v154_).f18) == ((d_377_v154_).f18))
                d_388_v162_ = nw66_
                d_389_v163_: _dafny.MultiSet
                d_389_v163_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygwxn"))])
                d_390_v164_: _dafny.Map
                d_390_v164_ = _dafny.Map({(d_377_v154_).f17: d_389_v163_})
                d_389_v163_ = (((d_390_v164_)[d_179_v6_] if (d_179_v6_) in (d_390_v164_) else d_389_v163_)) - (d_389_v163_)
            elif True:
                default__.m0(d_181_globalState_)
                index45_ = default__.safeIndex(527, (d_176_v5_).length(0))
                (d_176_v5_)[index45_] = (d_377_v154_).f18
                index46_ = default__.safeIndex(527, (d_176_v5_).length(0))
                (d_176_v5_)[index46_] = ((d_367_v146_)[(d_368_v147_).fm2(d_375_v152_, (d_377_v154_).f18, _dafny.Map({True: (d_377_v154_).f18}), d_179_v6_, d_181_globalState_)] if ((d_368_v147_).fm2(d_375_v152_, (d_377_v154_).f18, _dafny.Map({True: (d_377_v154_).f18}), d_179_v6_, d_181_globalState_)) in (d_367_v146_) else (d_377_v154_).f18)
                d_391_v165_: _dafny.MultiSet
                d_391_v165_ = _dafny.MultiSet([d_173_v3_])
                d_392_v166_: _dafny.Map
                d_392_v166_ = _dafny.Map({((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))])[default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))]: d_391_v165_})
                d_393_v167_: _dafny.Set
                d_393_v167_ = _dafny.Set({not(d_179_v6_)})
                d_394_v168_: _dafny.Array
                nw67_ = _dafny.Array(None, 14)
                nw67_[int(0)] = (0) - ((d_377_v154_).f18)
                nw67_[int(1)] = (-420 if (d_377_v154_).f17 else (((d_392_v166_)[False] if (False) in (d_392_v166_) else d_391_v165_)).cardinality)
                nw67_[int(2)] = (0) - (((_dafny.MultiSet([(d_377_v154_).f17, d_179_v6_])).cardinality) - ((d_377_v154_).f18))
                nw67_[int(3)] = d_171_v1_
                nw67_[int(4)] = len(d_393_v167_)
                nw67_[int(5)] = d_171_v1_
                nw67_[int(6)] = (d_171_v1_) - (844)
                nw67_[int(7)] = ((d_377_v154_).fm12(d_181_globalState_)) * ((d_377_v154_).f18)
                nw67_[int(8)] = 620
                nw67_[int(9)] = len(_dafny.Map({610: (d_377_v154_).f18}))
                nw67_[int(10)] = (d_377_v154_).f18
                nw67_[int(11)] = default__.safeModuloInt(d_171_v1_, (d_377_v154_).f18)
                nw67_[int(12)] = (d_176_v5_)[default__.safeIndex(527, (d_176_v5_).length(0))]
                nw67_[int(13)] = default__.safeDivisionInt(d_171_v1_, len(d_242_v54_))
                d_394_v168_ = nw67_
                (d_181_globalState_).f1 = d_394_v168_
                d_395_v169_: _dafny.Set
                d_395_v169_ = _dafny.Set({d_173_v3_, d_173_v3_})
                d_396_v170_: _dafny.Map
                d_396_v170_ = _dafny.Map({len(d_395_v169_): d_375_v152_})
                d_397_v171_: _dafny.Map
                d_397_v171_ = _dafny.Map({d_171_v1_: d_367_v146_})
                d_398_v172_: _dafny.Array
                nw68_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 20)
                d_398_v172_ = nw68_
                d_399_v173_: D5
                d_399_v173_ = D5_DC13(d_171_v1_, False, len(d_397_v171_), d_398_v172_, (d_377_v154_).f17)
                index47_ = default__.safeIndex(527, (d_176_v5_).length(0))
                rhs47_ = (d_242_v54_) + (d_242_v54_)
                rhs48_ = (d_377_v154_).f18
                rhs49_ = (d_368_v147_).fm2(_dafny.MultiSet([(d_377_v154_).f18, (((d_396_v170_)[len(d_243_v55_)] if (len(d_243_v55_)) in (d_396_v170_) else d_375_v152_)).cardinality]), (d_399_v173_).cf32, (default__.fm23((0) - ((d_176_v5_)[default__.safeIndex(527, (d_176_v5_).length(0))]), d_181_globalState_)).set(False, d_171_v1_), ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))])[default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))], d_181_globalState_)
                rhs50_ = ((_dafny.MultiSet(d_172_v2_)).intersection(default__.fm13((d_242_v54_)[default__.safeIndex((d_377_v154_).f18, len(d_242_v54_))], False, d_181_globalState_))) | (d_375_v152_)
                lhs38_ = d_176_v5_
                lhs39_ = default__.safeIndex(527, (d_176_v5_).length(0))
                lhs40_ = d_181_globalState_
                d_242_v54_ = rhs47_
                lhs38_[lhs39_] = rhs48_
                lhs40_.f7 = rhs49_
                d_375_v152_ = rhs50_
                (d_181_globalState_).f9 = (d_377_v154_).fm12(d_181_globalState_)
            d_400_v174_: D5
            d_400_v174_ = D5_DC14()
            source4_ = d_400_v174_
            if source4_.is_DC13:
                d_401___mcc_h12_ = source4_.cf30
                d_402___mcc_h13_ = source4_.cf31
                d_403___mcc_h14_ = source4_.cf32
                d_404___mcc_h15_ = source4_.cf33
                d_405___mcc_h16_ = source4_.cf34
                d_406_cf34_ = d_405___mcc_h16_
                d_407_cf33_ = d_404___mcc_h15_
                d_408_cf32_ = d_403___mcc_h14_
                d_409_cf31_ = d_402___mcc_h13_
                d_410_cf30_ = d_401___mcc_h12_
                d_411_v175_: C4
                nw69_ = C4()
                nw69_.ctor__(d_367_v146_, ((d_377_v154_).f18) in (d_243_v55_))
                d_411_v175_ = nw69_
                index48_ = default__.safeIndex(539, (d_176_v5_).length(0))
                (d_176_v5_)[index48_] = 135
                d_412_v176_: C3
                nw70_ = C3()
                nw70_.ctor__(d_408_cf32_, (d_377_v154_).f17, d_367_v146_, False)
                d_412_v176_ = nw70_
                d_413_v177_: _dafny.Map
                d_413_v177_ = _dafny.Map({(d_412_v176_).f13: d_410_cf30_})
                d_414_v178_: _dafny.Map
                d_414_v178_ = _dafny.Map({d_412_v176_: len(d_413_v177_)})
                d_415_v179_: _dafny.MultiSet
                d_415_v179_ = _dafny.MultiSet([_dafny.MultiSet([((d_414_v178_)[d_412_v176_] if (d_412_v176_) in (d_414_v178_) else (d_377_v154_).f18), default__.fm1(not(d_368_v147_.f15), d_171_v1_, d_171_v1_, d_181_globalState_), d_171_v1_]), (d_375_v152_) - (d_375_v152_), d_375_v152_, d_375_v152_])
                d_416_v180_: D9
                d_416_v180_ = D9_DC24(d_242_v54_, d_368_v147_.f15, d_171_v1_)
                index49_ = default__.safeIndex(539, (d_176_v5_).length(0))
                rhs51_ = d_173_v3_
                rhs52_ = d_411_v175_
                rhs53_ = (d_243_v55_).intersection(d_243_v55_)
                rhs54_ = len((d_416_v180_).cf56)
                rhs55_ = d_415_v179_
                lhs41_ = d_176_v5_
                lhs42_ = default__.safeIndex(539, (d_176_v5_).length(0))
                d_173_v3_ = rhs51_
                d_411_v175_ = rhs52_
                d_243_v55_ = rhs53_
                lhs41_[lhs42_] = rhs54_
                d_415_v179_ = rhs55_
                (d_181_globalState_).f9 = (d_377_v154_).f18
                d_417_v181_: _dafny.Seq
                d_417_v181_ = _dafny.SeqWithoutIsStrInference([d_367_v146_])
                d_418_v182_: T0
                nw71_ = C4()
                nw71_.ctor__((d_417_v181_)[default__.safeIndex((d_176_v5_)[default__.safeIndex(539, (d_176_v5_).length(0))], len(d_417_v181_))], not(d_406_cf34_))
                d_418_v182_ = nw71_
                d_419_v183_: _dafny.Map
                d_419_v183_ = _dafny.Map({d_418_v182_: (d_242_v54_).set(default__.safeIndex((0) - ((d_412_v176_).f13), len(d_242_v54_)), d_173_v3_)})
                d_420_v184_: _dafny.Seq
                d_420_v184_ = _dafny.SeqWithoutIsStrInference([default__.fm23(len(d_172_v2_), d_181_globalState_), default__.fm23((d_176_v5_)[default__.safeIndex(539, (d_176_v5_).length(0))], d_181_globalState_)])
                d_421_v185_: _dafny.MultiSet
                d_421_v185_ = _dafny.MultiSet([d_418_v182_, d_418_v182_])
                d_422_v186_: _dafny.Map
                d_422_v186_ = _dafny.Map({467: d_370_v149_})
                d_423_v187_: _dafny.Set
                d_423_v187_ = _dafny.Set({d_418_v182_.f12})
                d_424_v188_: _dafny.Set
                d_424_v188_ = _dafny.Set({d_423_v187_})
                arr2_ = (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]
                index50_ = default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))
                rhs56_ = d_419_v183_
                rhs57_ = ((d_421_v185_)[d_418_v182_] if (d_418_v182_) in (d_421_v185_) else default__.safeDivisionInt((d_377_v154_).f18, 222))
                rhs58_ = ((d_242_v54_) + (d_242_v54_)) < (_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_425_i18_ in range(default__.abs(805))]))
                rhs59_ = D8_DC21(d_368_v147_.f15, ((d_422_v186_)[(d_176_v5_)[default__.safeIndex(539, (d_176_v5_).length(0))]] if ((d_176_v5_)[default__.safeIndex(539, (d_176_v5_).length(0))]) in (d_422_v186_) else d_364_v144_), default__.safeModuloInt((d_176_v5_)[default__.safeIndex(539, (d_176_v5_).length(0))], (d_377_v154_).f18), (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))], (len(d_424_v188_)) * (default__.fm1(d_409_cf31_, d_410_cf30_, d_410_cf30_, d_181_globalState_)))
                rhs60_ = (d_420_v184_) + (d_420_v184_)
                lhs43_ = (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]
                lhs44_ = default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))
                d_419_v183_ = rhs56_
                d_410_cf30_ = rhs57_
                lhs43_[lhs44_] = rhs58_
                d_366_v145_ = rhs59_
                d_420_v184_ = rhs60_
                d_426_v189_: int
                d_427_v190_: bool
                out10_: int
                out11_: bool
                out10_, out11_ = (d_412_v176_).m2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jb")), (d_412_v176_).f14, d_407_cf33_, d_181_globalState_)
                d_426_v189_ = out10_
                d_427_v190_ = out11_
            elif source4_.is_DC14:
                (d_368_v147_).f15 = True
                (d_181_globalState_).f3 = d_368_v147_.f15
                d_376_v153_ = (d_376_v153_).set((d_377_v154_).fm2(d_375_v152_, len(d_363_v143_), d_376_v153_, d_179_v6_, d_181_globalState_), (d_377_v154_).f18)
                arr3_ = (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]
                index51_ = default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))
                rhs61_ = not(not(False))
                rhs62_ = (d_375_v152_) - ((d_375_v152_) | (d_375_v152_))
                rhs63_ = (d_377_v154_).f17
                rhs64_ = not(((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))])[default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))])
                lhs45_ = d_181_globalState_
                lhs46_ = (d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]
                lhs47_ = default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))
                d_179_v6_ = rhs61_
                d_375_v152_ = rhs62_
                lhs45_.f3 = rhs63_
                lhs46_[lhs47_] = rhs64_
            elif True:
                d_428___mcc_h17_ = source4_.cf29
                d_429_cf29_ = d_428___mcc_h17_
                d_430_v191_: C2
                nw72_ = C2()
                nw72_.ctor__(d_368_v147_.f15, d_367_v146_, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))])[default__.safeIndex(94, ((d_364_v144_)[default__.safeIndex(586, (d_364_v144_).length(0))]).length(0))])
                d_430_v191_ = nw72_
                (d_181_globalState_).f9 = len(d_242_v54_)
                d_431_v192_: bool
                d_432_v193_: int
                out12_: bool
                out13_: int
                out12_, out13_ = (d_430_v191_).m5(765, (d_242_v54_) + (d_242_v54_), d_181_globalState_)
                d_431_v192_ = out12_
                d_432_v193_ = out13_
                d_432_v193_ = (0) - (d_171_v1_)
        elif True:
            d_433_v194_: _dafny.Array
            nw73_ = _dafny.Array(D6.default()(), 6)
            d_433_v194_ = nw73_
            d_434_v195_: _dafny.Map
            d_434_v195_ = _dafny.Map({d_179_v6_: d_179_v6_})
            d_435_v196_: C1
            nw74_ = C1()
            nw74_.ctor__(d_171_v1_, (_dafny.MultiSet(d_180_v7_)).set(d_179_v6_, default__.abs(len(d_434_v195_))), d_179_v6_)
            d_435_v196_ = nw74_
            d_436_v197_: _dafny.MultiSet
            d_436_v197_ = _dafny.MultiSet([d_171_v1_, d_435_v196_.f16])
            d_437_v198_: _dafny.Map
            d_437_v198_ = _dafny.Map({d_435_v196_: d_436_v197_})
            d_438_v199_: D6
            d_438_v199_ = D6_DC15(((d_437_v198_)[d_435_v196_] if (d_435_v196_) in (d_437_v198_) else d_436_v197_))
            index52_ = default__.safeIndex(601, (d_433_v194_).length(0))
            (d_433_v194_)[index52_] = d_438_v199_
            index53_ = default__.safeIndex(601, (d_433_v194_).length(0))
            (d_433_v194_)[index53_] = d_438_v199_
            d_439_v200_: D1
            d_439_v200_ = D1_DC2(False, d_179_v6_, d_179_v6_, d_173_v3_, d_173_v3_)
            d_440_v201_: D1
            d_440_v201_ = D1_DC3(d_439_v200_)
            d_441_v202_: D1
            d_441_v202_ = D1_DC3(d_440_v201_)
            d_442_v203_: D1
            d_442_v203_ = D1_DC3(d_441_v202_)
            d_443_v204_: D1
            d_443_v204_ = D1_DC3(d_442_v203_)
            d_443_v204_ = d_443_v204_
            d_444_v206_: _dafny.Map
            d_444_v206_ = _dafny.Map({len(d_434_v195_): d_434_v195_})
            def iife48_():
                coll14_ = _dafny.Map()
                compr_14_: int
                for compr_14_ in _dafny.IntegerRange(359, -879):
                    d_445_v205_: int = compr_14_
                    if ((359) <= (d_445_v205_)) and ((d_445_v205_) < (-879)):
                        coll14_[default__.safeDivisionInt(d_445_v205_, (_dafny.MultiSet(d_180_v7_)).cardinality)] = (d_434_v195_).set(not(d_179_v6_), d_179_v6_)
                return _dafny.Map(coll14_)
            d_171_v1_ = default__.safeDivisionInt(len((iife48_()
            ) | ((d_444_v206_).set(len(d_242_v54_), d_434_v195_))), len(d_242_v54_))
            d_446_v207_: _dafny.Array
            def lambda16_(d_447_v7_, d_448_v1_):
                def lambda17_(d_449_i19_):
                    return (d_447_v7_)[default__.safeIndex(d_448_v1_, len(d_447_v7_))]

                return lambda17_

            init8_ = lambda16_(d_180_v7_, d_171_v1_)
            nw75_ = _dafny.Array(None, 15)
            for i0_8_ in range(nw75_.length(0)):
                nw75_[i0_8_] = init8_(i0_8_)
            d_446_v207_ = nw75_
            index54_ = default__.safeIndex(717, (d_446_v207_).length(0))
            (d_446_v207_)[index54_] = d_179_v6_
            d_450_v208_: _dafny.Set
            d_450_v208_ = _dafny.Set({True})
            index55_ = default__.safeIndex(717, (d_446_v207_).length(0))
            (d_446_v207_)[index55_] = not ((_dafny.Set({d_179_v6_, d_179_v6_, d_179_v6_})) != (d_450_v208_)) or ((d_179_v6_ if d_179_v6_ else d_179_v6_))
            d_451_v209_: _dafny.Set
            d_451_v209_ = _dafny.Set({(d_245_v57_)[default__.safeIndex(d_435_v196_.f16, len(d_245_v57_))]})
            def iife49_():
                coll15_ = _dafny.Set()
                compr_15_: _dafny.Seq
                for compr_15_ in (_dafny.Map({d_242_v54_: d_179_v6_})).keys.Elements:
                    d_452_v210_: _dafny.Seq = compr_15_
                    if (d_452_v210_) in (_dafny.Map({d_242_v54_: d_179_v6_})):
                        coll15_ = coll15_.union(_dafny.Set([d_452_v210_]))
                return _dafny.Set(coll15_)
            if ((d_451_v209_).intersection(iife49_()
            )).ispropersubset(d_451_v209_):
                d_180_v7_ = ((d_180_v7_) + (d_180_v7_)).set(default__.safeIndex(d_171_v1_, len((d_180_v7_) + (d_180_v7_))), d_179_v6_)
                d_242_v54_ = _dafny.SeqWithoutIsStrInference([d_173_v3_ for d_453_i20_ in range(default__.abs(32))])
                d_454_v211_: _dafny.MultiSet
                d_454_v211_ = _dafny.MultiSet([(d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))], d_179_v6_, (d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))]])
                d_455_v212_: C4
                nw76_ = C4()
                nw76_.ctor__(d_454_v211_, False)
                d_455_v212_ = nw76_
                d_456_v213_: C0
                nw77_ = C0()
                nw77_.ctor__((d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))], 449, d_454_v211_, (d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))])
                d_456_v213_ = nw77_
                d_179_v6_ = (d_456_v213_).f17
            elif True:
                d_457_v214_: _dafny.MultiSet
                d_457_v214_ = _dafny.MultiSet([(d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))]])
                d_458_v215_: _dafny.Array
                nw78_ = _dafny.Array(None, 20)
                nw78_[int(0)] = d_446_v207_
                nw78_[int(1)] = d_446_v207_
                nw78_[int(2)] = d_446_v207_
                nw78_[int(3)] = d_446_v207_
                nw78_[int(4)] = d_446_v207_
                nw78_[int(5)] = d_446_v207_
                nw78_[int(6)] = d_446_v207_
                nw78_[int(7)] = d_446_v207_
                nw78_[int(8)] = d_446_v207_
                nw78_[int(9)] = d_446_v207_
                nw78_[int(10)] = d_446_v207_
                nw78_[int(11)] = d_446_v207_
                nw78_[int(12)] = d_446_v207_
                nw78_[int(13)] = d_446_v207_
                nw78_[int(14)] = d_446_v207_
                nw78_[int(15)] = d_446_v207_
                nw78_[int(16)] = d_446_v207_
                nw78_[int(17)] = d_446_v207_
                nw78_[int(18)] = d_446_v207_
                nw78_[int(19)] = d_446_v207_
                d_458_v215_ = nw78_
                d_459_v216_: D8
                d_459_v216_ = D8_DC21(False, d_458_v215_, d_171_v1_, d_446_v207_, len(_dafny.Set({(d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))], (d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))]})))
                d_460_v217_: C4
                nw79_ = C4()
                nw79_.ctor__(d_457_v214_, (d_459_v216_).cf49)
                d_460_v217_ = nw79_
                d_461_v218_: _dafny.Map
                d_461_v218_ = _dafny.Map({True: d_460_v217_})
                d_462_v219_: _dafny.Map
                d_462_v219_ = _dafny.Map({d_435_v196_.f16: (d_446_v207_)[default__.safeIndex(717, (d_446_v207_).length(0))]})
                d_461_v218_ = (d_461_v218_).set((not(d_179_v6_)) and (((d_462_v219_)[d_435_v196_.f16] if (d_435_v196_.f16) in (d_462_v219_) else d_179_v6_)), d_460_v217_)
                d_457_v214_ = d_457_v214_
                (d_181_globalState_).f1 = (d_176_v5_ if (d_460_v217_).fm4(d_435_v196_.f16, d_435_v196_.f16, d_181_globalState_) else d_176_v5_)
                (d_181_globalState_).f9 = d_435_v196_.f16
                (d_435_v196_).f16 = d_171_v1_
        d_463_v220_: _dafny.Set
        d_463_v220_ = _dafny.Set({d_179_v6_})
        d_464_v221_: C0
        nw80_ = C0()
        nw80_.ctor__(d_179_v6_, len(d_463_v220_), _dafny.MultiSet([d_179_v6_, d_179_v6_]), d_179_v6_)
        d_464_v221_ = nw80_
        d_465_v222_: _dafny.MultiSet
        d_465_v222_ = _dafny.MultiSet([(d_464_v221_).f17, d_179_v6_, d_179_v6_])
        d_466_v223_: C3
        nw81_ = C3()
        nw81_.ctor__(d_171_v1_, (d_464_v221_).f17, d_465_v222_, d_179_v6_)
        d_466_v223_ = nw81_
        d_467_v224_: _dafny.Map
        d_467_v224_ = _dafny.Map({d_171_v1_: d_466_v223_})
        d_468_v225_: _dafny.Seq
        d_468_v225_ = _dafny.SeqWithoutIsStrInference([d_465_v222_, d_465_v222_, d_465_v222_, default__.fm22(len(d_467_v224_), d_179_v6_, d_181_globalState_)])
        d_469_v226_: _dafny.Map
        d_469_v226_ = _dafny.Map({d_464_v221_: ((d_468_v225_)[default__.safeIndex((d_464_v221_).f18, len(d_468_v225_))]).isdisjoint(default__.fm22((d_466_v223_).f13, (d_466_v223_).f14, d_181_globalState_))})
        d_469_v226_ = (d_469_v226_).set(d_464_v221_, not((d_466_v223_).f14))
        d_470_v227_: _dafny.Array
        def lambda18_(d_471_v54_):
            def lambda19_(d_472_i22_):
                return d_471_v54_

            return lambda19_

        init9_ = lambda18_(d_242_v54_)
        nw82_ = _dafny.Array(None, 27)
        for i0_9_ in range(nw82_.length(0)):
            nw82_[i0_9_] = init9_(i0_9_)
        d_470_v227_ = nw82_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_470_v227_).length(0)):
            d_473_i21_: int = guard_loop_0_
            if (True) and (((0) <= (d_473_i21_)) and ((d_473_i21_) < ((d_470_v227_).length(0)))):
                (d_470_v227_)[(d_473_i21_)] = d_242_v54_
        d_474_v228_: C1
        nw83_ = C1()
        nw83_.ctor__((d_466_v223_).f13, d_465_v222_, (d_464_v221_).f17)
        d_474_v228_ = nw83_
        d_475_v229_: _dafny.Seq
        d_475_v229_ = _dafny.SeqWithoutIsStrInference([d_474_v228_, d_474_v228_])
        d_476_i23_: int
        d_476_i23_ = 0
        with _dafny.label("4"):
            while not(not((((d_475_v229_) + (_dafny.SeqWithoutIsStrInference([d_474_v228_]))).set(default__.safeIndex(d_171_v1_, len((d_475_v229_) + (_dafny.SeqWithoutIsStrInference([d_474_v228_])))), d_474_v228_)) == (d_475_v229_))):
                with _dafny.c_label("4"):
                    if (d_476_i23_) >= (100):
                        raise _dafny.Break("4")
                    d_476_i23_ = (d_476_i23_) + (1)
                    d_477_v230_: T0
                    nw84_ = C1()
                    nw84_.ctor__((d_466_v223_).f13, _dafny.MultiSet([d_179_v6_]), (d_464_v221_).f17)
                    d_477_v230_ = nw84_
                    d_478_v231_: _dafny.Set
                    d_478_v231_ = _dafny.Set({d_477_v230_})
                    d_245_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_173_v3_ for d_479_i24_ in range(default__.abs(485))]), d_242_v54_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ot"))).set(default__.safeIndex((0) - (len(d_478_v231_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ot")))), _dafny.CodePoint('r')), d_242_v54_])
                    d_480_v233_: _dafny.Array
                    def lambda20_(d_481_v221_, d_482_v1_, d_483_v223_, d_484_v2_, d_485_v228_):
                        def lambda21_(d_486_i25_):
                            def iife50_():
                                coll16_ = _dafny.Map()
                                compr_16_: _dafny.Seq
                                for compr_16_ in (_dafny.MultiSet([(d_484_v2_).set(default__.safeIndex(d_485_v228_.f16, len(d_484_v2_)), d_482_v1_), d_484_v2_])).Elements:
                                    d_487_v232_: _dafny.Seq = compr_16_
                                    if (d_487_v232_) in (_dafny.MultiSet([(d_484_v2_).set(default__.safeIndex(d_485_v228_.f16, len(d_484_v2_)), d_482_v1_), d_484_v2_])):
                                        coll16_[d_487_v232_] = (d_483_v223_).f13
                                return _dafny.Map(coll16_)
                            return (_dafny.Map({_dafny.SeqWithoutIsStrInference([(d_481_v221_).f18, d_482_v1_]): (d_483_v223_).f13})) | (iife50_()
                            )

                        return lambda21_

                    init10_ = lambda20_(d_464_v221_, d_171_v1_, d_466_v223_, d_172_v2_, d_474_v228_)
                    nw85_ = _dafny.Array(None, 8)
                    for i0_10_ in range(nw85_.length(0)):
                        nw85_[i0_10_] = init10_(i0_10_)
                    d_480_v233_ = nw85_
                    d_488_v234_: _dafny.Map
                    d_488_v234_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([len(d_242_v54_)]): d_171_v1_})
                    d_489_v235_: _dafny.Map
                    d_489_v235_ = d_488_v234_
                    index56_ = default__.safeIndex(824, (d_480_v233_).length(0))
                    (d_480_v233_)[index56_] = (d_489_v235_)
                    index57_ = default__.safeIndex(824, (d_480_v233_).length(0))
                    def iife51_():
                        coll17_ = _dafny.Map()
                        compr_17_: _dafny.Seq
                        for compr_17_ in (_dafny.SeqWithoutIsStrInference([d_172_v2_ for d_490_i26_ in range(default__.abs(82))])).Elements:
                            d_491_v236_: _dafny.Seq = compr_17_
                            if (d_491_v236_) in (_dafny.SeqWithoutIsStrInference([d_172_v2_ for d_490_i26_ in range(default__.abs(82))])):
                                coll17_[d_491_v236_] = (d_464_v221_).f18
                        return _dafny.Map(coll17_)
                    (d_480_v233_)[index57_] = ((d_488_v234_) | (_dafny.Map({d_172_v2_: d_171_v1_}))) | ((d_488_v234_) | (iife51_()
                    ))
                    d_492_v237_: _dafny.MultiSet
                    d_492_v237_ = _dafny.MultiSet([d_474_v228_.f16, 813, (d_464_v221_).f18, len(d_180_v7_)])
                    d_493_v238_: _dafny.Map
                    d_493_v238_ = _dafny.Map({d_477_v230_.f12: (d_466_v223_).f13})
                    d_494_v239_: _dafny.Seq
                    d_494_v239_ = _dafny.SeqWithoutIsStrInference([d_172_v2_, (d_172_v2_ if (d_474_v228_).fm2(d_492_v237_, 673, d_493_v238_, (d_466_v223_).fm6((d_466_v223_).f14, d_171_v1_, d_181_globalState_), d_181_globalState_) else d_172_v2_), (d_172_v2_) + (d_172_v2_)])
                    d_495_v240_: _dafny.Seq
                    d_495_v240_ = _dafny.SeqWithoutIsStrInference([d_494_v239_, (d_494_v239_) + (d_494_v239_)])
                    d_496_v241_: D9
                    d_496_v241_ = D9_DC24(d_242_v54_, True, 707)
                    d_494_v239_ = ((d_495_v240_)[default__.safeIndex((d_466_v223_).f13, len(d_495_v240_))]).set(default__.safeIndex((d_464_v221_).f18, len((d_495_v240_)[default__.safeIndex((d_466_v223_).f13, len(d_495_v240_))])), (_dafny.SeqWithoutIsStrInference([(d_496_v241_).cf58])) + (d_172_v2_))
                    (d_477_v230_).f12 = not(d_179_v6_)
                    pass
            pass
        default__.m0(d_181_globalState_)
        if d_179_v6_:
            rhs65_ = d_242_v54_
            rhs66_ = d_172_v2_
            d_242_v54_ = rhs65_
            d_172_v2_ = rhs66_
            d_497_v242_: _dafny.Map
            d_497_v242_ = _dafny.Map({len(d_242_v54_): (d_464_v221_).f17})
            d_498_v243_: T0
            nw86_ = C0()
            nw86_.ctor__((d_466_v223_).f14, (d_464_v221_).f18, d_465_v222_, d_179_v6_)
            d_498_v243_ = nw86_
            d_499_v244_: _dafny.Map
            d_499_v244_ = _dafny.Map({d_173_v3_: d_498_v243_})
            d_500_v245_: _dafny.Seq
            d_500_v245_ = _dafny.SeqWithoutIsStrInference([(d_499_v244_).set(d_173_v3_, d_498_v243_)])
            d_501_v246_: C1
            nw87_ = C1()
            nw87_.ctor__(default__.safeModuloInt(847, len(d_242_v54_)), _dafny.MultiSet([((d_497_v242_)[len(d_500_v245_)] if (len(d_500_v245_)) in (d_497_v242_) else d_498_v243_.f12), (d_466_v223_).f14]), not((d_464_v221_).f17))
            d_501_v246_ = nw87_
            if not ((d_466_v223_).fm6(not(not(d_179_v6_)), (d_464_v221_).fm12(d_181_globalState_), d_181_globalState_)) or (d_498_v243_.f12):
                d_502_v247_: _dafny.Array
                nw88_ = _dafny.Array(False, 10)
                d_502_v247_ = nw88_
                d_503_v248_: _dafny.Seq
                d_503_v248_ = _dafny.SeqWithoutIsStrInference([d_502_v247_])
                d_504_v249_: _dafny.MultiSet
                d_504_v249_ = _dafny.MultiSet([d_474_v228_.f16, d_171_v1_, len(d_503_v248_), (0) - (d_474_v228_.f16), d_474_v228_.f16])
                d_505_v250_: _dafny.Map
                d_505_v250_ = _dafny.Map({d_171_v1_: d_504_v249_})
                d_506_v251_: _dafny.Array
                nw89_ = _dafny.Array(None, 12)
                nw89_[int(0)] = default__.fm13(d_173_v3_, not((d_464_v221_).f17), d_181_globalState_)
                nw89_[int(1)] = d_504_v249_
                nw89_[int(2)] = d_504_v249_
                nw89_[int(3)] = d_504_v249_
                nw89_[int(4)] = ((d_505_v250_)[(d_464_v221_).f18] if ((d_464_v221_).f18) in (d_505_v250_) else _dafny.MultiSet([d_474_v228_.f16]))
                nw89_[int(5)] = d_504_v249_
                nw89_[int(6)] = (d_504_v249_).set((d_466_v223_).f13, default__.abs((d_466_v223_).f13))
                nw89_[int(7)] = d_504_v249_
                nw89_[int(8)] = (_dafny.MultiSet([len(d_172_v2_)])) | (d_504_v249_)
                nw89_[int(9)] = _dafny.MultiSet([d_474_v228_.f16, d_501_v246_.f16, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "co")))])
                nw89_[int(10)] = (d_504_v249_ if d_498_v243_.f12 else d_504_v249_)
                nw89_[int(11)] = (_dafny.MultiSet([162, d_474_v228_.f16])).set((0) - (d_171_v1_), default__.abs(d_474_v228_.f16))
                d_506_v251_ = nw89_
                index58_ = default__.safeIndex(845, (d_506_v251_).length(0))
                (d_506_v251_)[index58_] = default__.fm13(_dafny.CodePoint('h'), (d_466_v223_).f14, d_181_globalState_)
                index59_ = default__.safeIndex(845, (d_506_v251_).length(0))
                (d_506_v251_)[index59_] = d_504_v249_
                d_507_v252_: _dafny.Map
                d_507_v252_ = _dafny.Map({(d_172_v2_).set(default__.safeIndex(d_501_v246_.f16, len(d_172_v2_)), (d_464_v221_).f18): 354})
                d_508_v253_: _dafny.Map
                d_508_v253_ = d_507_v252_
                d_509_v254_: _dafny.Map
                d_509_v254_ = _dafny.Map({d_508_v253_: (d_466_v223_).f14})
                d_510_v255_: _dafny.Array
                nw90_ = _dafny.Array(D8.default()(), 10)
                d_510_v255_ = nw90_
                d_511_v256_: _dafny.Map
                d_511_v256_ = _dafny.Map({d_509_v254_: d_510_v255_})
                index60_ = default__.safeIndex(120, (d_176_v5_).length(0))
                (d_176_v5_)[index60_] = default__.safeModuloInt((d_464_v221_).f18, d_474_v228_.f16)
                index61_ = default__.safeIndex(120, (d_176_v5_).length(0))
                rhs67_ = d_511_v256_
                rhs68_ = ((d_466_v223_).f13) != ((d_464_v221_).f18)
                rhs69_ = ((d_466_v223_ if (d_464_v221_).f17 else d_466_v223_) if (d_180_v7_)[default__.safeIndex((d_464_v221_).f18, len(d_180_v7_))] else d_466_v223_)
                rhs70_ = d_474_v228_.f16
                lhs48_ = d_498_v243_
                lhs49_ = d_176_v5_
                lhs50_ = default__.safeIndex(120, (d_176_v5_).length(0))
                d_511_v256_ = rhs67_
                lhs48_.f12 = rhs68_
                d_466_v223_ = rhs69_
                lhs49_[lhs50_] = rhs70_
                d_512_v257_: C2
                nw91_ = C2()
                nw91_.ctor__(d_498_v243_.f12, d_498_v243_.f11, d_498_v243_.f12)
                d_512_v257_ = nw91_
                d_513_v258_: D12
                d_513_v258_ = D12_DC32(default__.safeDivisionInt((d_464_v221_).f18, len(_dafny.Map({-861: d_501_v246_}))), d_512_v257_, d_464_v221_, (d_466_v223_).f13)
                d_514_v259_: _dafny.Map
                d_514_v259_ = _dafny.Map({d_501_v246_.f16: d_464_v221_})
                d_515_v260_: _dafny.Seq
                d_515_v260_ = _dafny.SeqWithoutIsStrInference([d_464_v221_, ((d_514_v259_)[(d_464_v221_).f18] if ((d_464_v221_).f18) in (d_514_v259_) else d_464_v221_), d_464_v221_])
                d_513_v258_ = D12_DC32(d_474_v228_.f16, d_512_v257_, (d_515_v260_)[default__.safeIndex(len(d_463_v220_), len(d_515_v260_))], (0) - (len(d_242_v54_)))
                d_516_v261_: C3
                nw92_ = C3()
                nw92_.ctor__((0) - ((d_466_v223_).f13), d_498_v243_.f12, (_dafny.MultiSet([False, (d_464_v221_).f17])) | (d_465_v222_), (d_464_v221_).f17)
                d_516_v261_ = nw92_
                d_517_v262_: _dafny.Array
                nw93_ = _dafny.Array(None, 13)
                nw93_[int(0)] = ((d_466_v223_).f13) + (d_474_v228_.f16)
                nw93_[int(1)] = (d_466_v223_).f13
                nw93_[int(2)] = (d_474_v228_.f16) + ((0) - (d_501_v246_.f16))
                nw93_[int(3)] = default__.fm1(d_512_v257_.f15, (d_466_v223_).f13, len(d_243_v55_), d_181_globalState_)
                nw93_[int(4)] = (0) - (((d_466_v223_).f13) * ((d_464_v221_).f18))
                nw93_[int(5)] = d_474_v228_.f16
                nw93_[int(6)] = (d_466_v223_).f13
                nw93_[int(7)] = (0) - (d_474_v228_.f16)
                nw93_[int(8)] = (d_466_v223_).f13
                nw93_[int(9)] = d_474_v228_.f16
                nw93_[int(10)] = (len(d_242_v54_) if not(not(d_512_v257_.f15)) else d_171_v1_)
                nw93_[int(11)] = len(d_180_v7_)
                nw93_[int(12)] = (len(d_242_v54_)) - ((d_464_v221_).f18)
                d_517_v262_ = nw93_
                rhs71_ = not(not (True) or (d_498_v243_.f12))
                rhs72_ = d_517_v262_
                rhs73_ = ((d_171_v1_) + ((d_464_v221_).f18) if ((d_516_v261_).f13) <= (-876) else 979)
                lhs51_ = d_181_globalState_
                lhs52_ = d_181_globalState_
                lhs51_.f3 = rhs71_
                lhs52_.f1 = rhs72_
                d_171_v1_ = rhs73_
            elif True:
                d_179_v6_ = d_498_v243_.f12
                d_518_v263_: _dafny.Map
                d_518_v263_ = _dafny.Map({d_172_v2_: d_474_v228_.f16})
                d_519_v264_: _dafny.Map
                d_519_v264_ = _dafny.Map({d_172_v2_: d_501_v246_.f16})
                d_520_v265_: _dafny.Array
                nw94_ = _dafny.Array(None, 4)
                nw94_[int(0)] = d_518_v263_
                nw94_[int(1)] = d_519_v264_
                nw94_[int(2)] = d_518_v263_
                nw94_[int(3)] = d_518_v263_
                d_520_v265_ = nw94_
                d_521_v266_: _dafny.Map
                d_521_v266_ = _dafny.Map({(d_466_v223_).f14: d_520_v265_})
                d_522_v267_: _dafny.Array
                nw95_ = _dafny.Array(None, 10)
                nw95_[int(0)] = (_dafny.Set({len(d_463_v220_), d_171_v1_, d_474_v228_.f16})).ispropersubset(d_243_v55_)
                nw95_[int(1)] = ((d_466_v223_).f13) > (533)
                nw95_[int(2)] = ((d_464_v221_).f17) and ((d_466_v223_).f14)
                nw95_[int(3)] = d_179_v6_
                nw95_[int(4)] = (d_466_v223_).f14
                nw95_[int(5)] = (d_466_v223_).f14
                nw95_[int(6)] = d_498_v243_.f12
                nw95_[int(7)] = (not(d_498_v243_.f12)) not in (d_521_v266_)
                nw95_[int(8)] = (d_180_v7_)[default__.safeIndex(default__.fm1(d_179_v6_, d_171_v1_, d_474_v228_.f16, d_181_globalState_), len(d_180_v7_))]
                nw95_[int(9)] = (d_464_v221_).f17
                d_522_v267_ = nw95_
                index62_ = default__.safeIndex(622, (d_522_v267_).length(0))
                (d_522_v267_)[index62_] = (d_466_v223_).f14
                index63_ = default__.safeIndex(622, (d_522_v267_).length(0))
                rhs74_ = ((d_172_v2_) <= (d_172_v2_)) or ((d_464_v221_).f17)
                rhs75_ = (d_464_v221_).f18
                rhs76_ = (d_464_v221_).f17
                rhs77_ = d_173_v3_
                lhs53_ = d_181_globalState_
                lhs54_ = d_181_globalState_
                lhs55_ = d_522_v267_
                lhs56_ = default__.safeIndex(622, (d_522_v267_).length(0))
                lhs53_.f7 = rhs74_
                lhs54_.f9 = rhs75_
                lhs55_[lhs56_] = rhs76_
                d_173_v3_ = rhs77_
                d_523_v268_: int
                out14_: int
                out14_ = (d_474_v228_).m6(d_181_globalState_)
                d_523_v268_ = out14_
                d_524_v269_: _dafny.Map
                d_524_v269_ = _dafny.Map({False: (d_466_v223_).f13})
                d_525_v270_: _dafny.Map
                d_525_v270_ = _dafny.Map({(d_466_v223_).f14: (d_466_v223_).f14})
                d_526_v271_: _dafny.MultiSet
                d_526_v271_ = _dafny.MultiSet([len(d_172_v2_), (d_464_v221_).fm12(d_181_globalState_), len(d_525_v270_), d_474_v228_.f16, (d_464_v221_).f18])
                d_527_v272_: _dafny.Map
                d_527_v272_ = _dafny.Map({d_176_v5_: True})
                d_528_v273_: _dafny.Map
                d_528_v273_ = _dafny.Map({((d_524_v269_)[(d_498_v243_).fm2(d_526_v271_, (0) - (len(d_525_v270_)), d_524_v269_, (d_522_v267_)[default__.safeIndex(622, (d_522_v267_).length(0))], d_181_globalState_)] if ((d_498_v243_).fm2(d_526_v271_, (0) - (len(d_525_v270_)), d_524_v269_, (d_522_v267_)[default__.safeIndex(622, (d_522_v267_).length(0))], d_181_globalState_)) in (d_524_v269_) else len(d_527_v272_)): d_522_v267_})
                d_529_v274_: _dafny.Set
                d_529_v274_ = _dafny.Set({d_173_v3_})
                d_530_v275_: _dafny.Map
                d_530_v275_ = _dafny.Map({(d_529_v274_) == (d_529_v274_): (d_522_v267_ if (d_466_v223_).f14 else d_522_v267_)})
                rhs78_ = d_528_v273_
                rhs79_ = ((d_522_v267_)[default__.safeIndex(622, (d_522_v267_).length(0))] if (d_466_v223_).f14 else d_498_v243_.f12)
                rhs80_ = len(d_242_v54_)
                rhs81_ = ((d_464_v221_).f18) + ((0) - (d_474_v228_.f16))
                rhs82_ = (d_530_v275_) | (d_530_v275_)
                lhs57_ = d_181_globalState_
                lhs58_ = d_474_v228_
                lhs59_ = d_474_v228_
                d_528_v273_ = rhs78_
                lhs57_.f3 = rhs79_
                lhs58_.f16 = rhs80_
                lhs59_.f16 = rhs81_
                d_530_v275_ = rhs82_
                d_525_v270_ = (d_525_v270_).set((d_464_v221_).f17, (d_466_v223_).f14)
            d_466_v223_ = d_466_v223_
            d_531_v276_: D1
            d_531_v276_ = D1_DC2((d_464_v221_).f17, (d_466_v223_).f14, (d_464_v221_).f17, d_173_v3_, d_173_v3_)
            pat_let_tv14_ = d_466_v223_
            d_532_v277_: _dafny.Array
            nw96_ = _dafny.Array(None, 12)
            nw96_[int(0)] = default__.fm14(d_531_v276_, d_501_v246_.f16, d_181_globalState_)
            nw96_[int(1)] = d_173_v3_
            nw96_[int(2)] = d_173_v3_
            nw96_[int(3)] = d_173_v3_
            def iife52_(_pat_let17_0):
                def iife53_(d_533_dt__update__tmp_h3_):
                    def iife54_(_pat_let18_0):
                        def iife55_(d_534_dt__update_hcf3_h0_):
                            return D1_DC2((d_533_dt__update__tmp_h3_).cf2, d_534_dt__update_hcf3_h0_, (d_533_dt__update__tmp_h3_).cf4, (d_533_dt__update__tmp_h3_).cf5, (d_533_dt__update__tmp_h3_).cf6)
                        return iife55_(_pat_let18_0)
                    return iife54_((pat_let_tv14_).f14)
                return iife53_(_pat_let17_0)
            nw96_[int(4)] = default__.fm14(iife52_(d_531_v276_), (d_464_v221_).f18, d_181_globalState_)
            nw96_[int(5)] = d_173_v3_
            nw96_[int(6)] = d_173_v3_
            nw96_[int(7)] = d_173_v3_
            nw96_[int(8)] = d_173_v3_
            nw96_[int(9)] = d_173_v3_
            nw96_[int(10)] = ((d_242_v54_)[default__.safeIndex(len(_dafny.Map({(d_466_v223_).fm6(d_179_v6_, (d_466_v223_).f13, d_181_globalState_): d_171_v1_})), len(d_242_v54_))] if not(not((d_466_v223_).f14)) else d_173_v3_)
            nw96_[int(11)] = d_173_v3_
            d_532_v277_ = nw96_
            index64_ = default__.safeIndex(314, (d_532_v277_).length(0))
            (d_532_v277_)[index64_] = default__.fm14(d_531_v276_, (d_464_v221_).f18, d_181_globalState_)
            d_535_v278_: C2
            nw97_ = C2()
            nw97_.ctor__((d_497_v242_) == (d_497_v242_), d_498_v243_.f11, d_498_v243_.f12)
            d_535_v278_ = nw97_
            d_536_v279_: _dafny.MultiSet
            d_536_v279_ = _dafny.MultiSet([len(d_180_v7_), len(d_172_v2_), d_474_v228_.f16, d_501_v246_.f16])
            d_537_v280_: _dafny.Map
            d_537_v280_ = _dafny.Map({-467: d_536_v279_})
            d_538_v281_: _dafny.Array
            nw98_ = _dafny.Array(None, 12)
            nw98_[int(0)] = (d_536_v279_).set(d_474_v228_.f16, default__.abs(d_171_v1_))
            nw98_[int(1)] = (d_536_v279_).intersection(d_536_v279_)
            nw98_[int(2)] = d_536_v279_
            nw98_[int(3)] = d_536_v279_
            nw98_[int(4)] = d_536_v279_
            nw98_[int(5)] = _dafny.MultiSet(d_172_v2_)
            nw98_[int(6)] = ((d_537_v280_)[(d_464_v221_).f18] if ((d_464_v221_).f18) in (d_537_v280_) else d_536_v279_)
            nw98_[int(7)] = d_536_v279_
            nw98_[int(8)] = d_536_v279_
            nw98_[int(9)] = (default__.fm13(d_173_v3_, (d_464_v221_).f17, d_181_globalState_)).set((d_464_v221_).f18, default__.abs((0) - (len(d_242_v54_))))
            nw98_[int(10)] = (_dafny.MultiSet([d_501_v246_.f16])) | (d_536_v279_)
            nw98_[int(11)] = d_536_v279_
            d_538_v281_ = nw98_
            index65_ = default__.safeIndex(81, (d_538_v281_).length(0))
            (d_538_v281_)[index65_] = d_536_v279_
            index66_ = default__.safeIndex(314, (d_532_v277_).length(0))
            index67_ = default__.safeIndex(81, (d_538_v281_).length(0))
            rhs83_ = d_474_v228_.f16
            rhs84_ = d_173_v3_
            rhs85_ = d_535_v278_
            rhs86_ = d_536_v279_
            lhs60_ = d_181_globalState_
            lhs61_ = d_532_v277_
            lhs62_ = default__.safeIndex(314, (d_532_v277_).length(0))
            lhs63_ = d_538_v281_
            lhs64_ = default__.safeIndex(81, (d_538_v281_).length(0))
            lhs60_.f9 = rhs83_
            lhs61_[lhs62_] = rhs84_
            d_535_v278_ = rhs85_
            lhs63_[lhs64_] = rhs86_
        elif True:
            d_539_v282_: T0
            nw99_ = C1()
            nw99_.ctor__((0) - ((d_465_v222_).cardinality), d_465_v222_, True)
            d_539_v282_ = nw99_
            d_540_v283_: _dafny.Map
            d_540_v283_ = _dafny.Map({d_539_v282_: (len(((d_245_v57_)[default__.safeIndex(len(d_243_v55_), len(d_245_v57_))]).set(default__.safeIndex((d_464_v221_).f18, len((d_245_v57_)[default__.safeIndex(len(d_243_v55_), len(d_245_v57_))])), d_173_v3_)) if (d_464_v221_).f17 else len(d_180_v7_))})
            d_540_v283_ = (d_540_v283_).set(d_539_v282_, default__.fm1((d_180_v7_)[default__.safeIndex((d_464_v221_).f18, len(d_180_v7_))], (0) - ((d_464_v221_).f18), (d_466_v223_).f13, d_181_globalState_))
            d_541_v284_: _dafny.Map
            d_541_v284_ = _dafny.Map({not((d_466_v223_).f14): False})
            d_542_v285_: _dafny.Seq
            d_542_v285_ = _dafny.SeqWithoutIsStrInference([d_541_v284_, (d_541_v284_).set(d_539_v282_.f12, not((d_466_v223_).fm6((d_466_v223_).f14, ((d_465_v222_)[(d_466_v223_).f14] if ((d_466_v223_).f14) in (d_465_v222_) else d_171_v1_), d_181_globalState_))), d_541_v284_])
            d_542_v285_ = _dafny.SeqWithoutIsStrInference([(d_541_v284_) | ((d_541_v284_).set((d_464_v221_).f17, d_179_v6_)), d_541_v284_, d_541_v284_, (d_541_v284_) | (d_541_v284_)])
            d_543_v286_: _dafny.Array
            def lambda22_(d_544_i27_):
                return True

            init11_ = lambda22_
            nw100_ = _dafny.Array(None, 25)
            for i0_11_ in range(nw100_.length(0)):
                nw100_[i0_11_] = init11_(i0_11_)
            d_543_v286_ = nw100_
            d_545_v287_: _dafny.Map
            d_545_v287_ = _dafny.Map({False: d_543_v286_})
            d_546_v288_: _dafny.Array
            nw101_ = _dafny.Array(None, 22)
            nw101_[int(0)] = d_543_v286_
            nw101_[int(1)] = d_543_v286_
            nw101_[int(2)] = d_543_v286_
            nw101_[int(3)] = d_543_v286_
            nw101_[int(4)] = d_543_v286_
            nw101_[int(5)] = d_543_v286_
            nw101_[int(6)] = d_543_v286_
            nw101_[int(7)] = d_543_v286_
            nw101_[int(8)] = d_543_v286_
            nw101_[int(9)] = d_543_v286_
            nw101_[int(10)] = d_543_v286_
            nw101_[int(11)] = d_543_v286_
            nw101_[int(12)] = d_543_v286_
            nw101_[int(13)] = d_543_v286_
            nw101_[int(14)] = d_543_v286_
            nw101_[int(15)] = ((d_545_v287_)[d_179_v6_] if (d_179_v6_) in (d_545_v287_) else d_543_v286_)
            nw101_[int(16)] = d_543_v286_
            nw101_[int(17)] = d_543_v286_
            nw101_[int(18)] = d_543_v286_
            nw101_[int(19)] = d_543_v286_
            nw101_[int(20)] = d_543_v286_
            nw101_[int(21)] = d_543_v286_
            d_546_v288_ = nw101_
            d_547_v289_: D8
            d_547_v289_ = D8_DC21((d_464_v221_).f17, d_546_v288_, (d_464_v221_).f18, d_543_v286_, default__.safeModuloInt((d_466_v223_).f13, (d_464_v221_).f18))
            source5_ = d_547_v289_
            if source5_.is_DC20:
                d_548___mcc_h18_ = source5_.cf46
                d_549___mcc_h19_ = source5_.cf47
                d_550___mcc_h20_ = source5_.cf48
                d_551_cf48_ = d_550___mcc_h20_
                d_552_cf47_ = d_549___mcc_h19_
                d_553_cf46_ = d_548___mcc_h18_
                d_554_v290_: _dafny.Seq
                d_554_v290_ = _dafny.SeqWithoutIsStrInference([d_463_v220_])
                d_171_v1_ = (0) - (len(((d_554_v290_)[default__.safeIndex(d_474_v228_.f16, len(d_554_v290_))]) - ((d_463_v220_).intersection(d_463_v220_))))
                d_555_v291_: _dafny.Map
                d_555_v291_ = _dafny.Map({803: d_173_v3_})
                d_173_v3_ = ((d_555_v291_)[d_474_v228_.f16] if (d_474_v228_.f16) in (d_555_v291_) else d_173_v3_)
                d_556_v292_: _dafny.MultiSet
                d_556_v292_ = _dafny.MultiSet([d_173_v3_, d_173_v3_, d_173_v3_, (d_242_v54_)[default__.safeIndex((d_464_v221_).f18, len(d_242_v54_))]])
                d_557_v293_: _dafny.MultiSet
                d_557_v293_ = _dafny.MultiSet([d_173_v3_])
                d_558_v294_: _dafny.Seq
                d_558_v294_ = _dafny.SeqWithoutIsStrInference([d_556_v292_])
                d_552_cf47_ = ((d_558_v294_)[default__.safeIndex(len(d_242_v54_), len(d_558_v294_))]).issubset((d_556_v292_ if True else (d_557_v293_)))
                (d_181_globalState_).f3 = (d_466_v223_).f14
            elif source5_.is_DC21:
                d_559___mcc_h21_ = source5_.cf49
                d_560___mcc_h22_ = source5_.cf50
                d_561___mcc_h23_ = source5_.cf51
                d_562___mcc_h24_ = source5_.cf52
                d_563___mcc_h25_ = source5_.cf53
                d_564_cf53_ = d_563___mcc_h25_
                d_565_cf52_ = d_562___mcc_h24_
                d_566_cf51_ = d_561___mcc_h23_
                d_567_cf50_ = d_560___mcc_h22_
                d_568_cf49_ = d_559___mcc_h21_
                d_566_cf51_ = len((d_245_v57_)[default__.safeIndex((d_171_v1_ if d_539_v282_.f12 else (d_464_v221_).f18), len(d_245_v57_))])
                (d_539_v282_).f12 = d_568_cf49_
                d_565_cf52_ = d_543_v286_
                d_569_v295_: _dafny.MultiSet
                d_569_v295_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_474_v228_.f16 for d_570_i28_ in range(default__.abs(235))])])
                d_566_cf51_ = ((d_569_v295_)[_dafny.SeqWithoutIsStrInference([(0) - (-198) for d_571_i29_ in range(default__.abs(899))])] if (_dafny.SeqWithoutIsStrInference([(0) - (-198) for d_572_i29_ in range(default__.abs(899))])) in (d_569_v295_) else default__.safeDivisionInt((d_464_v221_).f18, (d_464_v221_).f18))
            elif source5_.is_DC19:
                d_573___mcc_h26_ = source5_.cf45
                d_574_cf45_ = d_573___mcc_h26_
                default__.m0(d_181_globalState_)
                (d_181_globalState_).f9 = ((d_466_v223_).f13 if d_539_v282_.f12 else (d_466_v223_).f13)
                (d_181_globalState_).f9 = ((d_464_v221_).f18) * (d_474_v228_.f16)
                d_539_v282_ = d_539_v282_
            elif True:
                d_575___mcc_h27_ = source5_.cf54
                d_576_cf54_ = d_575___mcc_h27_
                (d_539_v282_).f12 = d_539_v282_.f12
                d_577_v296_: D1
                d_577_v296_ = D1_DC2(d_539_v282_.f12, (d_464_v221_).f17, (d_466_v223_).f14, d_173_v3_, d_173_v3_)
                d_173_v3_ = default__.fm14(d_577_v296_, d_171_v1_, d_181_globalState_)
                d_578_v297_: _dafny.Map
                d_578_v297_ = _dafny.Map({(d_466_v223_).f14: (d_464_v221_).f18})
                d_579_v298_: _dafny.Array
                nw102_ = _dafny.Array(_dafny.Seq({}), 1)
                d_579_v298_ = nw102_
                d_580_v299_: _dafny.Map
                d_580_v299_ = _dafny.Map({d_579_v298_: d_578_v297_})
                d_581_v300_: _dafny.Seq
                d_581_v300_ = _dafny.SeqWithoutIsStrInference([d_578_v297_, d_578_v297_])
                d_582_v301_: _dafny.Map
                d_582_v301_ = _dafny.Map({len(d_242_v54_): (d_464_v221_).f18})
                d_583_v302_: _dafny.MultiSet
                d_583_v302_ = _dafny.MultiSet([d_474_v228_.f16, (d_466_v223_).f13])
                d_584_v303_: _dafny.Map
                d_584_v303_ = _dafny.Map({_dafny.Set({((d_582_v301_)[(d_583_v302_).cardinality] if ((d_583_v302_).cardinality) in (d_582_v301_) else d_474_v228_.f16), d_171_v1_}): (d_464_v221_).f18})
                d_585_v304_: _dafny.Array
                nw103_ = _dafny.Array(None, 19)
                nw103_[int(0)] = d_578_v297_
                nw103_[int(1)] = d_578_v297_
                nw103_[int(2)] = ((d_580_v299_)[d_579_v298_] if (d_579_v298_) in (d_580_v299_) else d_578_v297_)
                nw103_[int(3)] = d_578_v297_
                nw103_[int(4)] = default__.fm23(d_474_v228_.f16, d_181_globalState_)
                nw103_[int(5)] = d_578_v297_
                nw103_[int(6)] = _dafny.Map({(d_464_v221_).f17: (d_466_v223_).f13})
                nw103_[int(7)] = d_578_v297_
                nw103_[int(8)] = (d_578_v297_) | (d_578_v297_)
                nw103_[int(9)] = (d_578_v297_) | (d_578_v297_)
                nw103_[int(10)] = _dafny.Map({(d_464_v221_).f17: (d_464_v221_).f18})
                nw103_[int(11)] = d_578_v297_
                nw103_[int(12)] = d_578_v297_
                nw103_[int(13)] = d_578_v297_
                nw103_[int(14)] = d_578_v297_
                nw103_[int(15)] = (d_578_v297_) | (d_578_v297_)
                nw103_[int(16)] = d_578_v297_
                nw103_[int(17)] = d_578_v297_
                nw103_[int(18)] = (d_578_v297_) | ((d_581_v300_)[default__.safeIndex(len(d_584_v303_), len(d_581_v300_))])
                d_585_v304_ = nw103_
                d_586_v305_: _dafny.Array
                def lambda23_(d_587_v1_):
                    def lambda24_(d_588_i30_):
                        return D9_DC26(d_587_v1_)

                    return lambda24_

                init12_ = lambda23_(d_171_v1_)
                nw104_ = _dafny.Array(None, 19)
                for i0_12_ in range(nw104_.length(0)):
                    nw104_[i0_12_] = init12_(i0_12_)
                d_586_v305_ = nw104_
                rhs87_ = (_dafny.MultiSet([d_586_v305_])).cardinality
                rhs88_ = d_474_v228_
                rhs89_ = (d_464_v221_).f18
                rhs90_ = (d_466_v223_).f14
                rhs91_ = d_585_v304_
                lhs65_ = d_474_v228_
                lhs66_ = d_474_v228_
                lhs67_ = d_181_globalState_
                lhs65_.f16 = rhs87_
                d_474_v228_ = rhs88_
                lhs66_.f16 = rhs89_
                lhs67_.f8 = rhs90_
                d_585_v304_ = rhs91_
                d_589_v306_: _dafny.Map
                d_589_v306_ = _dafny.Map({True: d_474_v228_})
                d_590_v307_: _dafny.Array
                nw105_ = _dafny.Array(None, 4)
                nw105_[int(0)] = ((d_589_v306_)[d_179_v6_] if (d_179_v6_) in (d_589_v306_) else d_474_v228_)
                nw105_[int(1)] = d_474_v228_
                nw105_[int(2)] = d_474_v228_
                nw105_[int(3)] = d_474_v228_
                d_590_v307_ = nw105_
                d_590_v307_ = d_590_v307_
            default__.m0(d_181_globalState_)
            d_543_v286_ = d_543_v286_
        (d_181_globalState_).f7 = d_179_v6_
        d_591_i31_: int
        d_591_i31_ = 0
        with _dafny.label("5"):
            while (d_464_v221_).f17:
                with _dafny.c_label("5"):
                    if (d_591_i31_) >= (100):
                        raise _dafny.Break("5")
                    d_591_i31_ = (d_591_i31_) + (1)
                    (d_474_v228_).f16 = d_171_v1_
                    d_592_v308_: C2
                    nw106_ = C2()
                    nw106_.ctor__((d_464_v221_).f17, d_465_v222_, (d_466_v223_).f14)
                    d_592_v308_ = nw106_
                    d_593_v309_: D12
                    d_593_v309_ = D12_DC32(len(d_242_v54_), d_592_v308_, d_464_v221_, (d_466_v223_).f13)
                    d_594_v310_: _dafny.Map
                    d_594_v310_ = _dafny.Map({d_172_v2_: d_593_v309_})
                    d_595_v311_: _dafny.Map
                    d_595_v311_ = _dafny.Map({d_474_v228_.f16: len(_dafny.Set({d_592_v308_.f15}))})
                    d_596_v312_: _dafny.Map
                    d_596_v312_ = _dafny.Map({(d_466_v223_).f13: len(d_595_v311_)})
                    d_594_v310_ = (d_594_v310_).set(d_172_v2_, D12_DC32(d_474_v228_.f16, d_592_v308_, d_464_v221_, ((d_595_v311_)[-252] if (-252) in (d_595_v311_) else (d_464_v221_).f18)))
                    d_597_v313_: _dafny.Seq
                    d_597_v313_ = _dafny.SeqWithoutIsStrInference([d_464_v221_])
                    d_598_v314_: _dafny.MultiSet
                    d_598_v314_ = _dafny.MultiSet([len(d_597_v313_)])
                    d_599_v315_: D9
                    d_599_v315_ = D9_DC25(d_242_v54_, d_592_v308_.f15, (d_464_v221_).f18)
                    d_600_v316_: _dafny.Map
                    d_600_v316_ = _dafny.Map({(d_474_v228_).fm2(d_598_v314_, d_474_v228_.f16, _dafny.Map({(d_599_v315_).cf60: d_474_v228_.f16}), (d_466_v223_).f14, d_181_globalState_): 373})
                    d_601_v317_: C4
                    nw107_ = C4()
                    nw107_.ctor__(d_465_v222_, (d_464_v221_).f17)
                    d_601_v317_ = nw107_
                    d_602_v318_: _dafny.Map
                    d_602_v318_ = _dafny.Map({((d_600_v316_)[d_592_v308_.f15] if (d_592_v308_.f15) in (d_600_v316_) else 312): d_601_v317_})
                    d_602_v318_ = (d_602_v318_).set(615, (d_601_v317_ if (d_466_v223_).f14 else d_601_v317_))
                    d_603_v319_: int
                    d_604_v320_: bool
                    out15_: int
                    out16_: bool
                    out15_, out16_ = (d_466_v223_).m2(d_242_v54_, (d_466_v223_).f14, d_470_v227_, d_181_globalState_)
                    d_603_v319_ = out15_
                    d_604_v320_ = out16_
                    pass
            pass
        _dafny.print(_dafny.string_of(d_171_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_172_v2_) == (_dafny.SeqWithoutIsStrInference([-351, 231, -351, -351, 231, -351]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_173_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v4_) == (_dafny.Map({_dafny.Map({0: False, 582: False}): _dafny.CodePoint('p')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v5_)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_179_v6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_180_v7_) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_181_globalState_).f0) == (_dafny.Map({_dafny.Map({0: False, 582: False}): _dafny.CodePoint('p')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f1)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_).f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_globalState_.f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f4)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_.f6) == (_dafny.SeqWithoutIsStrInference([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_globalState_.f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_181_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_181_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_182_i1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_199_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_242_v54_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v55_) == (_dafny.Set({-351}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v56_).cf40))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_244_v56_).cf41).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v56_).cf42))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_244_v56_).cf43))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_244_v56_).cf44) == (_dafny.Set({-351}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_245_v57_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eyxck"))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_356_v140_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mupivve")): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_463_v220_) == (_dafny.Set({False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_464_v221_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_464_v221_).f18))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_465_v222_) == (_dafny.MultiSet([False, False, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_466_v223_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_466_v223_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_467_v224_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_468_v225_) == (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([False, False, False]), _dafny.MultiSet([False, False, False]), _dafny.MultiSet([False, False, False]), _dafny.MultiSet([True])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_469_v226_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[3]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[10]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[11]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[12]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[13]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[14]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[15]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[16]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[17]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[18]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[19]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[20]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[21]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[22]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[23]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[24]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[25]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_470_v227_)[26]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_474_v228_.f16))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_475_v229_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_476_i23_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_591_i31_))
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
        return lambda: D1_DC2(False, False, False, _dafny.CodePoint('D'), _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D1_DC2)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D1_DC1)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC2(D1, NamedTuple('DC2', [('cf2', Any), ('cf3', Any), ('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)}, {_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3 and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC1(D1, NamedTuple('DC1', [('cf1', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC1({_dafny.string_of(self.cf1)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC1) and self.cf1 == __o.cf1
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC5(int(0), int(0), int(0), int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D2_DC4)

class D2_DC5(D2, NamedTuple('DC5', [('cf9', Any), ('cf10', Any), ('cf11', Any), ('cf12', Any), ('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC5({_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC5) and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC4(D2, NamedTuple('DC4', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC4({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC4) and self.cf8 == __o.cf8
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC7(int(0), int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D3_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D3_DC8)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D3_DC6)

class D3_DC7(D3, NamedTuple('DC7', [('cf15', Any), ('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC7({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC7) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf18', Any), ('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

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
        return lambda: D4_DC11(False, int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D4_DC11)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D4_DC10)

class D4_DC11(D4, NamedTuple('DC11', [('cf25', Any), ('cf26', Any), ('cf27', Any), ('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)}, {_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27 and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC13(int(0), False, int(0), _dafny.Array(None, 0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D5_DC12)

class D5_DC13(D5, NamedTuple('DC13', [('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any), ('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33 and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC14(D5, NamedTuple('DC14', [])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14)
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC12(D5, NamedTuple('DC12', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC12({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC12) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(int(0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, False, _dafny.Set({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [('cf40', Any), ('cf41', Any), ('cf42', Any), ('cf43', Any), ('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf40)}, {self.cf41.VerbatimString(True)}, {_dafny.string_of(self.cf42)}, {_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42 and self.cf43 == __o.cf43 and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D8_DC21)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC20(D8, NamedTuple('DC20', [('cf46', Any), ('cf47', Any), ('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf46)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf49', Any), ('cf50', Any), ('cf51', Any), ('cf52', Any), ('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)}, {_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50 and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf45)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC24(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC24(D9, NamedTuple('DC24', [('cf56', Any), ('cf57', Any), ('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({self.cf56.VerbatimString(True)}, {_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf56 == __o.cf56 and self.cf57 == __o.cf57 and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC25(D9, NamedTuple('DC25', [('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({self.cf59.VerbatimString(True)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC23(D9, NamedTuple('DC23', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC28(None, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D10_DC29)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC28(D10, NamedTuple('DC28', [('cf64', Any), ('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC29(D10, NamedTuple('DC29', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC29({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC29) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)

class D11_DC30(D11, NamedTuple('DC30', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(int(0), None, None, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)

class D12_DC32(D12, NamedTuple('DC32', [('cf71', Any), ('cf72', Any), ('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D13_DC34)

class D13_DC34(D13, NamedTuple('DC34', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC34({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC34) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D14_DC35)

class D14_DC35(D14, NamedTuple('DC35', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC35({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC35) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC37(_dafny.Seq({}), False, False, _dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D15_DC37)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D15_DC36)

class D15_DC37(D15, NamedTuple('DC37', [('cf79', Any), ('cf80', Any), ('cf81', Any), ('cf82', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC37({_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)}, {_dafny.string_of(self.cf81)}, {_dafny.string_of(self.cf82)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC37) and self.cf79 == __o.cf79 and self.cf80 == __o.cf80 and self.cf81 == __o.cf81 and self.cf82 == __o.cf82
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC36(D15, NamedTuple('DC36', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC36({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC36) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value

class GlobalState:
    def  __init__(self):
        self.f1: _dafny.Array = _dafny.Array(None, 0)
        self.f3: bool = False
        self.f4: _dafny.Array = _dafny.Array(None, 0)
        self.f6: _dafny.Seq = _dafny.Seq({})
        self.f7: bool = False
        self.f8: bool = False
        self.f9: int = int(0)
        self._f0: _dafny.Map = _dafny.Map({})
        self._f2: bool = False
        self._f5: int = int(0)
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self).f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self).f4 = f4
        (self)._f5 = f5
        (self).f6 = f6
        (self).f7 = f7
        (self).f8 = f8
        (self).f9 = f9
        (self)._f10 = f10

    @property
    def f0(self):
        return self._f0
    @property
    def f2(self):
        return self._f2
    @property
    def f5(self):
        return self._f5
    @property
    def f10(self):
        return self._f10

class C0(T0):
    def  __init__(self):
        self._f11: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        self._f17: bool = False
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f17, f18, f11, f12):
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f11 = f11
        (self).f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        return self.f12

    def fm11(self, p0, p1, p2, globalState):
        return _dafny.Map({default__.safeDivisionInt((self).f18, (self).f18): D1_DC2(self.f12, False, (self).f17, _dafny.CodePoint('f'), _dafny.CodePoint('x'))})

    def fm12(self, globalState):
        return default__.safeDivisionInt((self).f18, 291)

    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18

class C1(T0):
    def  __init__(self):
        self._f11: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        self.f16: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f16, f11, f12):
        (self).f16 = f16
        (self).f11 = f11
        (self).f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        source6_ = D2_DC5(self.f16, self.f16, self.f16, -458, _dafny.SeqWithoutIsStrInference([self.f12]))
        if source6_.is_DC5:
            d_605___mcc_h0_ = source6_.cf9
            d_606___mcc_h1_ = source6_.cf10
            d_607___mcc_h2_ = source6_.cf11
            d_608___mcc_h3_ = source6_.cf12
            d_609___mcc_h4_ = source6_.cf13
            d_610_cf13_ = d_609___mcc_h4_
            d_611_cf12_ = d_608___mcc_h3_
            d_612_cf11_ = d_607___mcc_h2_
            d_613_cf10_ = d_606___mcc_h1_
            d_614_cf9_ = d_605___mcc_h0_
            return self.f12
        elif True:
            d_615___mcc_h5_ = source6_.cf8
            d_616_cf8_ = d_615___mcc_h5_
            return not((self.f12 if self.f12 else False))

    def m6(self, globalState):
        r0: int = int(0)
        d_617_i0_: int
        d_617_i0_ = 0
        with _dafny.label("6"):
            while self.f12:
                with _dafny.c_label("6"):
                    if (d_617_i0_) >= (100):
                        raise _dafny.Break("6")
                    d_617_i0_ = (d_617_i0_) + (1)
                    if ((self.f16 if self.f12 else len(_dafny.Map({self.f12: self.f12})))) == (default__.safeDivisionInt(self.f16, self.f16)):
                        d_618_v0_: D2
                        d_618_v0_ = D2_DC4(43)
                        (globalState).f9 = default__.safeModuloInt((d_618_v0_).cf8, self.f16)
                        d_619_v1_: _dafny.Array
                        def lambda25_(d_620_i1_):
                            return (d_620_i1_) + (self.f16)

                        init13_ = lambda25_
                        nw108_ = _dafny.Array(None, 27)
                        for i0_13_ in range(nw108_.length(0)):
                            nw108_[i0_13_] = init13_(i0_13_)
                        d_619_v1_ = nw108_
                        d_621_v2_: _dafny.Seq
                        d_621_v2_ = _dafny.SeqWithoutIsStrInference([d_619_v1_])
                        d_622_v3_: _dafny.Map
                        d_622_v3_ = _dafny.Map({self.f12: len(((d_621_v2_) + (d_621_v2_)).set(default__.safeIndex(self.f16, len((d_621_v2_) + (d_621_v2_))), (d_621_v2_)[default__.safeIndex(self.f16, len(d_621_v2_))]))})
                        d_622_v3_ = (d_622_v3_) | (d_622_v3_)
                        index68_ = default__.safeIndex(524, (d_619_v1_).length(0))
                        (d_619_v1_)[index68_] = (0) - (self.f16)
                        index69_ = default__.safeIndex(524, (d_619_v1_).length(0))
                        (d_619_v1_)[index69_] = 276
                        d_623_v4_: _dafny.Array
                        nw109_ = _dafny.Array(False, 3)
                        d_623_v4_ = nw109_
                        d_623_v4_ = d_623_v4_
                        d_624_v5_: _dafny.Array
                        nw110_ = _dafny.Array(_dafny.Array(None, 0), 27)
                        d_624_v5_ = nw110_
                        index70_ = default__.safeIndex(596, (d_624_v5_).length(0))
                        (d_624_v5_)[index70_] = d_623_v4_
                        index71_ = default__.safeIndex(240, (d_623_v4_).length(0))
                        (d_623_v4_)[index71_] = False
                        d_625_v6_: _dafny.Array
                        nw111_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 22)
                        d_625_v6_ = nw111_
                        d_626_v7_: str
                        d_626_v7_ = _dafny.CodePoint('t')
                        d_627_v8_: D1
                        d_627_v8_ = D1_DC2(self.f12, self.f12, True, d_626_v7_, d_626_v7_)
                        index72_ = default__.safeIndex(789, (d_625_v6_).length(0))
                        (d_625_v6_)[index72_] = default__.fm9((d_619_v1_)[default__.safeIndex(524, (d_619_v1_).length(0))], (d_627_v8_).cf3, globalState)
                        d_628_v9_: _dafny.Seq
                        d_628_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcdmomb"))
                        d_629_v10_: _dafny.MultiSet
                        d_629_v10_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_626_v7_ for d_630_i2_ in range(default__.abs(248))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqudam")), d_628_v9_])
                        d_631_v11_: _dafny.Map
                        d_631_v11_ = _dafny.Map({self.f12: self.f12})
                        d_632_v12_: _dafny.Seq
                        d_632_v12_ = _dafny.SeqWithoutIsStrInference([len(d_631_v11_), (d_619_v1_)[default__.safeIndex(524, (d_619_v1_).length(0))], len(d_628_v9_), (d_619_v1_)[default__.safeIndex(524, (d_619_v1_).length(0))]])
                        d_633_v13_: _dafny.Seq
                        d_633_v13_ = _dafny.SeqWithoutIsStrInference([d_628_v9_])
                        d_634_v14_: _dafny.MultiSet
                        d_634_v14_ = _dafny.MultiSet([self.f16])
                        d_635_v15_: _dafny.Map
                        d_635_v15_ = _dafny.Map({(d_619_v1_)[default__.safeIndex(524, (d_619_v1_).length(0))]: self.f16})
                        index73_ = default__.safeIndex(596, (d_624_v5_).length(0))
                        index74_ = default__.safeIndex(240, (d_623_v4_).length(0))
                        index75_ = default__.safeIndex(789, (d_625_v6_).length(0))
                        rhs92_ = d_623_v4_
                        rhs93_ = (d_629_v10_).ispropersubset((d_629_v10_) - (d_629_v10_))
                        rhs94_ = (_dafny.SeqWithoutIsStrInference([len(d_631_v11_)])) == (d_632_v12_)
                        rhs95_ = (d_628_v9_) + ((d_633_v13_)[default__.safeIndex(default__.fm1((self).fm2(d_634_v14_, self.f16, _dafny.Map({self.f12: -762}), self.f12, globalState), (d_619_v1_)[default__.safeIndex(524, (d_619_v1_).length(0))], len(d_635_v15_), globalState), len(d_633_v13_))])
                        lhs68_ = d_624_v5_
                        lhs69_ = default__.safeIndex(596, (d_624_v5_).length(0))
                        lhs70_ = self
                        lhs71_ = d_623_v4_
                        lhs72_ = default__.safeIndex(240, (d_623_v4_).length(0))
                        lhs73_ = d_625_v6_
                        lhs74_ = default__.safeIndex(789, (d_625_v6_).length(0))
                        lhs68_[lhs69_] = rhs92_
                        lhs70_.f12 = rhs93_
                        lhs71_[lhs72_] = rhs94_
                        lhs73_[lhs74_] = rhs95_
                    elif True:
                        d_636_v16_: str
                        d_636_v16_ = _dafny.CodePoint('t')
                        d_637_v17_: _dafny.Seq
                        d_637_v17_ = _dafny.SeqWithoutIsStrInference([d_636_v16_])
                        d_638_v18_: _dafny.Seq
                        d_638_v18_ = _dafny.SeqWithoutIsStrInference([self.f12, (_dafny.CodePoint('p')) not in (d_637_v17_), self.f12, self.f12])
                        (globalState).f6 = (d_638_v18_).set(default__.safeIndex((self.f16) * (self.f16), len(d_638_v18_)), True)
                        d_639_v19_: _dafny.MultiSet
                        d_639_v19_ = _dafny.MultiSet([len(d_637_v17_)])
                        d_640_v20_: _dafny.Array
                        nw112_ = _dafny.Array(int(0), 27)
                        d_640_v20_ = nw112_
                        d_641_v21_: _dafny.Map
                        d_641_v21_ = _dafny.Map({d_638_v18_: d_640_v20_})
                        (globalState).f9 = (((d_639_v19_)[len(d_641_v21_)] if (len(d_641_v21_)) in (d_639_v19_) else 354) if not((d_638_v18_)[default__.safeIndex(self.f16, len(d_638_v18_))]) else -844)
                        (globalState).f3 = not(self.f12)
                        (globalState).f9 = self.f16
                        d_642_v22_: D1
                        d_642_v22_ = D1_DC2(self.f12, self.f12, self.f12, d_636_v16_, d_636_v16_)
                        (globalState).f9 = (0) - (len(_dafny.Map({(self.f12 if self.f12 else self.f12): (d_642_v22_).cf5})))
                    d_643_v23_: str
                    d_643_v23_ = _dafny.CodePoint('q')
                    d_644_v24_: _dafny.Seq
                    d_644_v24_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_643_v23_, globalState)])
                    d_645_v25_: _dafny.Seq
                    d_645_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfreelwss"))
                    d_646_v26_: _dafny.Array
                    nw113_ = _dafny.Array(None, 23)
                    nw113_[int(0)] = self.f12
                    nw113_[int(1)] = (self.f12 if self.f12 else False)
                    nw113_[int(2)] = self.f12
                    nw113_[int(3)] = self.f12
                    nw113_[int(4)] = self.f12
                    nw113_[int(5)] = True
                    nw113_[int(6)] = (self.f12) == (not(self.f12))
                    nw113_[int(7)] = (d_644_v24_) == (d_644_v24_)
                    nw113_[int(8)] = self.f12
                    nw113_[int(9)] = not (self.f12) or (self.f12)
                    nw113_[int(10)] = (self.f12) or (not(self.f12))
                    nw113_[int(11)] = self.f12
                    nw113_[int(12)] = self.f12
                    nw113_[int(13)] = (len(_dafny.Map({self.f12: self.f12}))) > (self.f16)
                    nw113_[int(14)] = self.f12
                    nw113_[int(15)] = self.f12
                    nw113_[int(16)] = (len(d_645_v25_)) > (self.f16)
                    nw113_[int(17)] = not(self.f12)
                    nw113_[int(18)] = self.f12
                    nw113_[int(19)] = (self.f12 if self.f12 else self.f12)
                    nw113_[int(20)] = self.f12
                    nw113_[int(21)] = self.f12
                    nw113_[int(22)] = self.f12
                    d_646_v26_ = nw113_
                    index76_ = default__.safeIndex(19, (d_646_v26_).length(0))
                    (d_646_v26_)[index76_] = self.f12
                    d_647_v27_: _dafny.Map
                    d_647_v27_ = _dafny.Map({not(self.f12): (d_645_v25_) == (d_645_v25_)})
                    d_648_v28_: _dafny.Seq
                    d_648_v28_ = _dafny.SeqWithoutIsStrInference([d_644_v24_, d_644_v24_, _dafny.SeqWithoutIsStrInference([self.f12, self.f12]), d_644_v24_, default__.fm10(self.f12, self.f12, True, globalState)])
                    d_649_v29_: _dafny.MultiSet
                    d_649_v29_ = _dafny.MultiSet([len((d_648_v28_)[default__.safeIndex(self.f16, len(d_648_v28_))])])
                    d_650_v30_: _dafny.Map
                    d_650_v30_ = _dafny.Map({self.f12: self.f16})
                    index77_ = default__.safeIndex(19, (d_646_v26_).length(0))
                    rhs96_ = (self.f16) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddcprsn"))))
                    rhs97_ = ((d_647_v27_)[self.f12] if (self.f12) in (d_647_v27_) else not (self.f12) or ((self).fm2(d_649_v29_, self.f16, d_650_v30_, self.f12, globalState)))
                    lhs75_ = globalState
                    lhs76_ = d_646_v26_
                    lhs77_ = default__.safeIndex(19, (d_646_v26_).length(0))
                    lhs75_.f9 = rhs96_
                    lhs76_[lhs77_] = rhs97_
                    d_651_v31_: _dafny.Array
                    nw114_ = _dafny.Array(_dafny.Array(None, 0), 2)
                    d_651_v31_ = nw114_
                    index78_ = default__.safeIndex(811, (d_651_v31_).length(0))
                    (d_651_v31_)[index78_] = d_646_v26_
                    index79_ = default__.safeIndex(811, (d_651_v31_).length(0))
                    rhs98_ = d_649_v29_
                    rhs99_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvme"))) + (d_645_v25_)).set(default__.safeIndex((self.f16) - (self.f16), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvme"))) + (d_645_v25_))), d_643_v23_)
                    rhs100_ = d_646_v26_
                    rhs101_ = d_646_v26_
                    lhs78_ = d_651_v31_
                    lhs79_ = default__.safeIndex(811, (d_651_v31_).length(0))
                    d_649_v29_ = rhs98_
                    d_645_v25_ = rhs99_
                    lhs78_[lhs79_] = rhs100_
                    d_646_v26_ = rhs101_
                    d_652_v32_: _dafny.MultiSet
                    d_652_v32_ = _dafny.MultiSet([d_643_v23_, d_643_v23_, d_643_v23_])
                    d_653_v33_: _dafny.Seq
                    d_653_v33_ = _dafny.SeqWithoutIsStrInference([self.f16, self.f16])
                    d_654_v34_: _dafny.Map
                    d_654_v34_ = _dafny.Map({d_652_v32_: d_653_v33_})
                    d_654_v34_ = (d_654_v34_).set(d_652_v32_, (d_653_v33_).set(default__.safeIndex(self.f16, len(d_653_v33_)), self.f16))
                    pass
            pass
        d_655_v35_: _dafny.Array
        nw115_ = _dafny.Array(int(0), 26)
        d_655_v35_ = nw115_
        index80_ = default__.safeIndex(829, (d_655_v35_).length(0))
        (d_655_v35_)[index80_] = self.f16
        d_656_v36_: _dafny.Seq
        d_656_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bso"))
        index81_ = default__.safeIndex(829, (d_655_v35_).length(0))
        (d_655_v35_)[index81_] = default__.safeDivisionInt(len(d_656_v36_), len(_dafny.SeqWithoutIsStrInference([D1_DC3(D1_DC2(self.f12, self.f12, self.f12, _dafny.CodePoint('k'), _dafny.CodePoint('d'))) for d_657_i3_ in range(default__.abs(52))])))
        d_658_v37_: _dafny.Map
        d_658_v37_ = _dafny.Map({True: self.f12})
        d_658_v37_ = d_658_v37_
        d_659_v38_: D2
        d_659_v38_ = D2_DC4((0) - (self.f16))
        d_660_v39_: _dafny.Array
        def lambda26_(d_661_i4_):
            return False

        init14_ = lambda26_
        nw116_ = _dafny.Array(None, 28)
        for i0_14_ in range(nw116_.length(0)):
            nw116_[i0_14_] = init14_(i0_14_)
        d_660_v39_ = nw116_
        d_662_v40_: _dafny.Map
        d_662_v40_ = _dafny.Map({(d_659_v38_ if self.f12 else d_659_v38_): d_660_v39_})
        index82_ = default__.safeIndex(829, (d_655_v35_).length(0))
        rhs102_ = (self.f16) + ((self.f16) * ((d_655_v35_)[default__.safeIndex(829, (d_655_v35_).length(0))]))
        rhs103_ = d_662_v40_
        lhs80_ = d_655_v35_
        lhs81_ = default__.safeIndex(829, (d_655_v35_).length(0))
        lhs80_[lhs81_] = rhs102_
        d_662_v40_ = rhs103_
        d_663_v41_: _dafny.Seq
        d_663_v41_ = _dafny.SeqWithoutIsStrInference([self.f12])
        d_664_v42_: _dafny.Map
        d_664_v42_ = _dafny.Map({self.f16: d_663_v41_})
        d_665_v43_: _dafny.Array
        nw117_ = _dafny.Array(None, 10)
        nw117_[int(0)] = d_663_v41_
        nw117_[int(1)] = d_663_v41_
        nw117_[int(2)] = d_663_v41_
        nw117_[int(3)] = (d_663_v41_).set(default__.safeIndex((d_655_v35_)[default__.safeIndex(829, (d_655_v35_).length(0))], len(d_663_v41_)), self.f12)
        nw117_[int(4)] = d_663_v41_
        nw117_[int(5)] = d_663_v41_
        nw117_[int(6)] = default__.fm10(not(self.f12), self.f12, (d_663_v41_)[default__.safeIndex(self.f16, len(d_663_v41_))], globalState)
        nw117_[int(7)] = d_663_v41_
        nw117_[int(8)] = (((d_664_v42_)[default__.fm1(self.f12, self.f16, self.f16, globalState)] if (default__.fm1(self.f12, self.f16, self.f16, globalState)) in (d_664_v42_) else d_663_v41_)) + (d_663_v41_)
        nw117_[int(9)] = _dafny.SeqWithoutIsStrInference([self.f12])
        d_665_v43_ = nw117_
        d_666_v44_: bool
        d_666_v44_ = self.f12
        d_667_v45_: _dafny.Set
        d_667_v45_ = _dafny.Set({self.f12})
        d_668_v46_: _dafny.Seq
        d_668_v46_ = _dafny.SeqWithoutIsStrInference([d_663_v41_, d_663_v41_, (_dafny.SeqWithoutIsStrInference([self.f12, (d_666_v44_), self.f12])).set(default__.safeIndex(len(d_667_v45_), len(_dafny.SeqWithoutIsStrInference([self.f12, (d_666_v44_), self.f12]))), self.f12), d_663_v41_, d_663_v41_])
        d_669_v47_: _dafny.Map
        d_669_v47_ = _dafny.Map({(d_655_v35_)[default__.safeIndex(829, (d_655_v35_).length(0))]: self.f12})
        index83_ = default__.safeIndex(21, (d_665_v43_).length(0))
        (d_665_v43_)[index83_] = (d_668_v46_)[default__.safeIndex(len(d_669_v47_), len(d_668_v46_))]
        index84_ = default__.safeIndex(21, (d_665_v43_).length(0))
        (d_665_v43_)[index84_] = d_663_v41_
        d_666_v44_ = d_666_v44_
        r0 = len(_dafny.Map({not(self.f12): (d_655_v35_)[default__.safeIndex(829, (d_655_v35_).length(0))]}))
        return r0

    def m7(self, p0, p1, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        d_670_v0_: bool
        d_670_v0_ = self.f12
        d_671_v1_: str
        d_671_v1_ = _dafny.CodePoint('s')
        d_672_v2_: D1
        d_672_v2_ = D1_DC2((d_670_v0_), self.f12, self.f12, d_671_v1_, _dafny.CodePoint('w'))
        source7_ = d_672_v2_
        if source7_.is_DC2:
            d_673___mcc_h0_ = source7_.cf2
            d_674___mcc_h1_ = source7_.cf3
            d_675___mcc_h2_ = source7_.cf4
            d_676___mcc_h3_ = source7_.cf5
            d_677___mcc_h4_ = source7_.cf6
            d_678_cf6_ = d_677___mcc_h4_
            d_679_cf5_ = d_676___mcc_h3_
            d_680_cf4_ = d_675___mcc_h2_
            d_681_cf3_ = d_674___mcc_h1_
            d_682_cf2_ = d_673___mcc_h0_
            (globalState).f9 = self.f16
            d_683_v3_: _dafny.Seq
            d_683_v3_ = _dafny.SeqWithoutIsStrInference([self.f16, self.f16, self.f16])
            (self).f16 = (d_683_v3_)[default__.safeIndex(self.f16, len(d_683_v3_))]
            rhs104_ = d_681_cf3_
            rhs105_ = (d_683_v3_) + ((d_683_v3_) + (d_683_v3_))
            rhs106_ = False
            lhs82_ = globalState
            lhs83_ = globalState
            lhs82_.f7 = rhs104_
            d_683_v3_ = rhs105_
            lhs83_.f3 = rhs106_
            if ((p0) - (_dafny.MultiSet([self.f16]))).ispropersubset(p0):
                (globalState).f7 = (d_680_cf4_ if d_681_cf3_ else not (self.f12) or (self.f12))
                d_684_v4_: _dafny.Map
                d_684_v4_ = _dafny.Map({len(default__.fm9(self.f16, d_681_cf3_, globalState)): d_681_cf3_})
                d_685_v5_: _dafny.Array
                def lambda27_(d_686_cf4_):
                    def lambda28_(d_687_i0_):
                        return d_686_cf4_

                    return lambda28_

                init15_ = lambda27_(d_680_cf4_)
                nw118_ = _dafny.Array(None, 6)
                for i0_15_ in range(nw118_.length(0)):
                    nw118_[i0_15_] = init15_(i0_15_)
                d_685_v5_ = nw118_
                d_688_v6_: _dafny.Map
                d_688_v6_ = _dafny.Map({d_685_v5_: d_680_cf4_})
                d_684_v4_ = (d_684_v4_).set((0) - (self.f16), ((d_688_v6_)[d_685_v5_] if (d_685_v5_) in (d_688_v6_) else d_680_cf4_))
                d_689_v7_: _dafny.Seq
                d_689_v7_ = _dafny.SeqWithoutIsStrInference([d_680_cf4_])
                rhs107_ = len(default__.fm9(-404, d_681_cf3_, globalState))
                rhs108_ = default__.safeDivisionInt(self.f16, 302)
                rhs109_ = d_682_cf2_
                rhs110_ = (_dafny.MultiSet(d_689_v7_)) | ((self.f11 if d_680_cf4_ else _dafny.MultiSet([d_682_cf2_, d_681_cf3_, d_681_cf3_])))
                rhs111_ = d_681_cf3_
                lhs84_ = globalState
                lhs85_ = globalState
                lhs86_ = self
                lhs87_ = globalState
                lhs84_.f9 = rhs107_
                lhs85_.f9 = rhs108_
                d_681_cf3_ = rhs109_
                lhs86_.f11 = rhs110_
                lhs87_.f3 = rhs111_
                (globalState).f9 = self.f16
                d_690_v8_: _dafny.Seq
                d_690_v8_ = _dafny.SeqWithoutIsStrInference([p1])
                d_691_v9_: _dafny.Set
                d_691_v9_ = _dafny.Set({self.f16})
                d_692_v10_: _dafny.Array
                nw119_ = _dafny.Array(None, 23)
                nw119_[int(0)] = p1
                nw119_[int(1)] = p1
                nw119_[int(2)] = (d_690_v8_)[default__.safeIndex(self.f16, len(d_690_v8_))]
                nw119_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exlcjc"))
                nw119_[int(4)] = p1
                nw119_[int(5)] = p1
                nw119_[int(6)] = p1
                nw119_[int(7)] = default__.fm9(len(d_691_v9_), d_681_cf3_, globalState)
                nw119_[int(8)] = p1
                nw119_[int(9)] = _dafny.SeqWithoutIsStrInference([d_671_v1_ for d_693_i1_ in range(default__.abs(527))])
                nw119_[int(10)] = p1
                nw119_[int(11)] = _dafny.SeqWithoutIsStrInference([d_678_cf6_ for d_694_i2_ in range(default__.abs(110))])
                nw119_[int(12)] = p1
                nw119_[int(13)] = p1
                nw119_[int(14)] = p1
                nw119_[int(15)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mw"))).set(default__.safeIndex(len(d_689_v7_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mw")))), d_678_cf6_)
                nw119_[int(16)] = p1
                nw119_[int(17)] = p1
                nw119_[int(18)] = p1
                nw119_[int(19)] = p1
                nw119_[int(20)] = _dafny.SeqWithoutIsStrInference([d_671_v1_ for d_695_i3_ in range(default__.abs(793))])
                nw119_[int(21)] = _dafny.SeqWithoutIsStrInference([d_678_cf6_ for d_696_i4_ in range(default__.abs(233))])
                nw119_[int(22)] = p1
                d_692_v10_ = nw119_
                d_697_v11_: _dafny.Map
                d_697_v11_ = _dafny.Map({d_679_cf5_: d_692_v10_})
                d_698_v12_: _dafny.Map
                d_698_v12_ = _dafny.Map({d_682_cf2_: d_692_v10_})
                d_699_v13_: _dafny.Seq
                d_699_v13_ = _dafny.SeqWithoutIsStrInference([d_692_v10_])
                d_700_v14_: _dafny.Array
                nw120_ = _dafny.Array(None, 19)
                nw120_[int(0)] = d_692_v10_
                nw120_[int(1)] = d_692_v10_
                nw120_[int(2)] = d_692_v10_
                nw120_[int(3)] = ((d_697_v11_)[d_671_v1_] if (d_671_v1_) in (d_697_v11_) else d_692_v10_)
                nw120_[int(4)] = (d_692_v10_ if self.f12 else d_692_v10_)
                nw120_[int(5)] = d_692_v10_
                nw120_[int(6)] = d_692_v10_
                nw120_[int(7)] = d_692_v10_
                nw120_[int(8)] = d_692_v10_
                nw120_[int(9)] = ((d_698_v12_)[d_680_cf4_] if (d_680_cf4_) in (d_698_v12_) else d_692_v10_)
                nw120_[int(10)] = d_692_v10_
                nw120_[int(11)] = d_692_v10_
                nw120_[int(12)] = d_692_v10_
                nw120_[int(13)] = d_692_v10_
                nw120_[int(14)] = ((d_699_v13_)[default__.safeIndex(self.f16, len(d_699_v13_))] if d_680_cf4_ else d_692_v10_)
                nw120_[int(15)] = d_692_v10_
                nw120_[int(16)] = d_692_v10_
                nw120_[int(17)] = ((d_698_v12_)[d_682_cf2_] if (d_682_cf2_) in (d_698_v12_) else d_692_v10_)
                nw120_[int(18)] = d_692_v10_
                d_700_v14_ = nw120_
                index85_ = default__.safeIndex(160, (d_700_v14_).length(0))
                (d_700_v14_)[index85_] = d_692_v10_
                index86_ = default__.safeIndex(160, (d_700_v14_).length(0))
                (d_700_v14_)[index86_] = d_692_v10_
            elif True:
                d_701_v15_: _dafny.Map
                d_701_v15_ = _dafny.Map({self.f12: self.f16})
                (globalState).f7 = (self).fm2(((p0).set(self.f16, default__.abs(self.f16))) | (p0), default__.safeDivisionInt(768, (0) - (self.f16)), d_701_v15_, (self.f16) == (self.f16), globalState)
                d_671_v1_ = d_671_v1_
                (globalState).f9 = (self.f16) - (self.f16)
                d_702_v16_: _dafny.Seq
                d_702_v16_ = _dafny.SeqWithoutIsStrInference([d_680_cf4_, d_682_cf2_])
                d_703_v17_: _dafny.Map
                d_703_v17_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_671_v1_ for d_704_i5_ in range(default__.abs(374))]): (d_702_v16_) + (d_702_v16_)})
                d_705_v18_: T0
                nw121_ = C0()
                nw121_.ctor__(d_681_cf3_, self.f16, self.f11, d_682_cf2_)
                d_705_v18_ = nw121_
                d_706_v19_: _dafny.Map
                d_706_v19_ = _dafny.Map({d_705_v18_: d_702_v16_})
                d_703_v17_ = (d_703_v17_).set((p1 if d_681_cf3_ else p1), ((d_706_v19_)[d_705_v18_] if (d_705_v18_) in (d_706_v19_) else d_702_v16_))
                d_707_v20_: _dafny.Array
                nw122_ = _dafny.Array(None, 26)
                d_707_v20_ = nw122_
                d_708_v21_: _dafny.Map
                d_708_v21_ = _dafny.Map({self.f16: self.f16})
                d_709_v22_: _dafny.Map
                d_709_v22_ = _dafny.Map({d_678_cf6_: p1})
                rhs112_ = default__.safeDivisionInt(len((d_702_v16_) + (d_702_v16_)), self.f16)
                rhs113_ = (d_702_v16_)[default__.safeIndex((self.f16) * (self.f16), len(d_702_v16_))]
                rhs114_ = default__.safeModuloInt(default__.safeDivisionInt((0) - (len(_dafny.Map({False: d_680_cf4_}))), len(_dafny.SeqWithoutIsStrInference([d_682_cf2_]))), (self.f16) + (self.f16))
                rhs115_ = d_707_v20_
                rhs116_ = (self.f16 if d_682_cf2_ else (0) - ((self.f16) - (((d_708_v21_)[self.f16] if (self.f16) in (d_708_v21_) else len(d_709_v22_)))))
                lhs88_ = globalState
                lhs89_ = globalState
                lhs90_ = self
                lhs91_ = globalState
                lhs88_.f9 = rhs112_
                lhs89_.f7 = rhs113_
                lhs90_.f16 = rhs114_
                d_707_v20_ = rhs115_
                lhs91_.f9 = rhs116_
        elif source7_.is_DC1:
            d_710___mcc_h5_ = source7_.cf1
            d_711_cf1_ = d_710___mcc_h5_
            d_712_v23_: D1
            d_712_v23_ = D1_DC1(d_711_cf1_)
            d_713_v24_: _dafny.Map
            d_713_v24_ = _dafny.Map({d_712_v23_: d_711_cf1_})
            d_714_v25_: _dafny.Array
            nw123_ = _dafny.Array(None, 8)
            nw123_[int(0)] = (p0).cardinality
            nw123_[int(1)] = len((p1) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))))
            nw123_[int(2)] = 135
            nw123_[int(3)] = self.f16
            nw123_[int(4)] = (self.f16) * (default__.fm1(self.f12, 467, self.f16, globalState))
            nw123_[int(5)] = self.f16
            nw123_[int(6)] = self.f16
            nw123_[int(7)] = len(((d_713_v24_)[d_712_v23_] if (d_712_v23_) in (d_713_v24_) else d_711_cf1_))
            d_714_v25_ = nw123_
            (globalState).f1 = d_714_v25_
            d_715_v27_: _dafny.Map
            def iife56_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(893, 501):
                    d_716_v26_: int = compr_18_
                    if ((893) <= (d_716_v26_)) and ((d_716_v26_) < (501)):
                        coll18_[(d_716_v26_) + (len(_dafny.Map({self.f16: d_671_v1_})))] = self.f12
                return _dafny.Map(coll18_)
            d_715_v27_ = _dafny.Map({self.f16: _dafny.Map({(iife56_()
            ).set(73, self.f12): self.f16})})
            d_717_v29_: _dafny.Set
            d_717_v29_ = _dafny.Set({self.f16})
            def iife57_():
                coll19_ = _dafny.Map()
                compr_19_: int
                for compr_19_ in (d_717_v29_).Elements:
                    d_718_v28_: int = compr_19_
                    if (d_718_v28_) in (d_717_v29_):
                        coll19_[default__.safeModuloInt(d_718_v28_, 553)] = self.f12
                return _dafny.Map(coll19_)
            if (self.f12 if self.f12 else (self.f16) < (len(((d_715_v27_)[-670] if (-670) in (d_715_v27_) else _dafny.Map({iife57_()
            : self.f16}))))):
                (globalState).f7 = self.f12
                (self).f12 = self.f12
                rhs117_ = default__.fm1(self.f12, self.f16, self.f16, globalState)
                rhs118_ = False
                lhs92_ = self
                lhs93_ = globalState
                lhs92_.f16 = rhs117_
                lhs93_.f3 = rhs118_
                (globalState).f9 = default__.safeModuloInt(default__.fm1(self.f12, 834, ((self.f11)[self.f12] if (self.f12) in (self.f11) else 849), globalState), (self.f16) + (945))
                d_719_v30_: _dafny.Array
                nw124_ = _dafny.Array(_dafny.Array(None, 0), 21)
                d_719_v30_ = nw124_
                d_720_v31_: _dafny.Map
                d_720_v31_ = _dafny.Map({_dafny.MultiSet([self.f16]): d_719_v30_})
                d_721_v32_: D2
                d_721_v32_ = D2_DC4(self.f16)
                d_722_v33_: _dafny.Map
                d_722_v33_ = _dafny.Map({self.f12: d_721_v32_})
                d_720_v31_ = (d_720_v31_).set((p0).intersection((p0).set(len(d_722_v33_), default__.abs(self.f16))), d_719_v30_)
            elif True:
                d_723_v34_: _dafny.Seq
                d_723_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pekwurnd"))
                d_724_v35_: _dafny.Map
                d_724_v35_ = _dafny.Map({d_671_v1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tfxhr"))})
                d_723_v34_ = ((d_724_v35_)[_dafny.CodePoint('w')] if (_dafny.CodePoint('w')) in (d_724_v35_) else p1)
                d_725_v36_: _dafny.Array
                def lambda29_(d_726_i6_):
                    return not((self.f12 if True else True))

                init16_ = lambda29_
                nw125_ = _dafny.Array(None, 14)
                for i0_16_ in range(nw125_.length(0)):
                    nw125_[i0_16_] = init16_(i0_16_)
                d_725_v36_ = nw125_
                index87_ = default__.safeIndex(726, (d_725_v36_).length(0))
                (d_725_v36_)[index87_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gdkv")))) > (self.f16)
                index88_ = default__.safeIndex(726, (d_725_v36_).length(0))
                (d_725_v36_)[index88_] = self.f12
                d_727_v37_: D3
                d_727_v37_ = D3_DC6(d_725_v36_)
                d_725_v36_ = (d_727_v37_).cf14
                d_728_v38_: _dafny.Array
                nw126_ = _dafny.Array(_dafny.MultiSet({}), 9)
                d_728_v38_ = nw126_
                index89_ = default__.safeIndex(284, (d_728_v38_).length(0))
                (d_728_v38_)[index89_] = default__.fm13(d_671_v1_, self.f12, globalState)
                index90_ = default__.safeIndex(284, (d_728_v38_).length(0))
                (d_728_v38_)[index90_] = (p0).intersection(p0)
                index91_ = default__.safeIndex(726, (d_725_v36_).length(0))
                (d_725_v36_)[index91_] = self.f12
            d_729_v39_: _dafny.Map
            d_729_v39_ = _dafny.Map({not(False): self.f16})
            d_730_v40_: _dafny.Seq
            d_730_v40_ = _dafny.SeqWithoutIsStrInference([d_729_v39_, d_729_v39_])
            d_730_v40_ = (((_dafny.SeqWithoutIsStrInference([d_729_v39_])).set(default__.safeIndex(len(p1), len(_dafny.SeqWithoutIsStrInference([d_729_v39_]))), d_729_v39_)) + (_dafny.SeqWithoutIsStrInference([d_729_v39_])) if True else d_730_v40_)
            d_731_v41_: _dafny.Array
            nw127_ = _dafny.Array(False, 26)
            d_731_v41_ = nw127_
            d_731_v41_ = d_731_v41_
        elif True:
            d_732___mcc_h6_ = source7_.cf7
            d_733_cf7_ = d_732___mcc_h6_
            if not(default__.fm0(d_671_v1_, globalState)):
                d_734_v42_: _dafny.Array
                nw128_ = _dafny.Array(None, 24)
                nw128_[int(0)] = d_671_v1_
                nw128_[int(1)] = d_671_v1_
                nw128_[int(2)] = d_671_v1_
                nw128_[int(3)] = d_671_v1_
                nw128_[int(4)] = _dafny.CodePoint('f')
                nw128_[int(5)] = _dafny.CodePoint('w')
                nw128_[int(6)] = d_671_v1_
                nw128_[int(7)] = d_671_v1_
                nw128_[int(8)] = d_671_v1_
                nw128_[int(9)] = d_671_v1_
                nw128_[int(10)] = d_671_v1_
                nw128_[int(11)] = _dafny.CodePoint('c')
                nw128_[int(12)] = d_671_v1_
                nw128_[int(13)] = d_671_v1_
                nw128_[int(14)] = d_671_v1_
                nw128_[int(15)] = _dafny.CodePoint('r')
                nw128_[int(16)] = d_671_v1_
                nw128_[int(17)] = d_671_v1_
                nw128_[int(18)] = d_671_v1_
                nw128_[int(19)] = d_671_v1_
                nw128_[int(20)] = d_671_v1_
                nw128_[int(21)] = d_671_v1_
                nw128_[int(22)] = d_671_v1_
                nw128_[int(23)] = _dafny.CodePoint('w')
                d_734_v42_ = nw128_
                d_735_v43_: _dafny.Seq
                d_735_v43_ = _dafny.SeqWithoutIsStrInference([d_734_v42_])
                d_735_v43_ = d_735_v43_
                d_736_v44_: _dafny.Map
                d_736_v44_ = _dafny.Map({588: self.f12})
                d_736_v44_ = (d_736_v44_).set(self.f16, self.f12)
                d_737_v45_: _dafny.Set
                d_737_v45_ = _dafny.Set({self.f12, self.f12, self.f12, self.f12})
                (globalState).f8 = (self.f16) == (len(d_737_v45_))
                d_738_v46_: _dafny.Array
                nw129_ = _dafny.Array(False, 11)
                d_738_v46_ = nw129_
                d_738_v46_ = d_738_v46_
                d_739_v47_: _dafny.Seq
                d_739_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "njgjri"))
                d_740_v48_: _dafny.Array
                def lambda30_(d_741_i7_):
                    return (d_741_i7_) * ((0) - (self.f16))

                init17_ = lambda30_
                nw130_ = _dafny.Array(None, 27)
                for i0_17_ in range(nw130_.length(0)):
                    nw130_[i0_17_] = init17_(i0_17_)
                d_740_v48_ = nw130_
                d_742_v49_: D4
                d_742_v49_ = D4_DC10(d_740_v48_)
                d_743_v50_: _dafny.Seq
                d_743_v50_ = _dafny.SeqWithoutIsStrInference([d_740_v48_])
                pat_let_tv15_ = d_740_v48_
                d_744_v51_: _dafny.Array
                nw131_ = _dafny.Array(None, 11)
                nw131_[int(0)] = d_740_v48_
                nw131_[int(1)] = d_740_v48_
                nw131_[int(2)] = (d_740_v48_ if self.f12 else d_740_v48_)
                def iife58_(_pat_let19_0):
                    def iife59_(d_745_dt__update__tmp_h0_):
                        def iife60_(_pat_let20_0):
                            def iife61_(d_746_dt__update_hcf24_h0_):
                                return D4_DC10(d_746_dt__update_hcf24_h0_)
                            return iife61_(_pat_let20_0)
                        return iife60_(pat_let_tv15_)
                    return iife59_(_pat_let19_0)
                nw131_[int(3)] = (iife58_(d_742_v49_)).cf24
                nw131_[int(4)] = (d_743_v50_)[default__.safeIndex(self.f16, len(d_743_v50_))]
                nw131_[int(5)] = d_740_v48_
                nw131_[int(6)] = d_740_v48_
                nw131_[int(7)] = d_740_v48_
                nw131_[int(8)] = d_740_v48_
                nw131_[int(9)] = d_740_v48_
                nw131_[int(10)] = d_740_v48_
                d_744_v51_ = nw131_
                d_747_v52_: _dafny.Map
                d_747_v52_ = _dafny.Map({False: d_671_v1_})
                d_748_v53_: _dafny.Map
                d_748_v53_ = _dafny.Map({self.f12: self.f16})
                index92_ = default__.safeIndex(387, (d_738_v46_).length(0))
                (d_738_v46_)[index92_] = ((self).fm2(_dafny.MultiSet([((self.f11)[not(self.f12)] if (not(self.f12)) in (self.f11) else len(d_747_v52_))]), self.f16, d_748_v53_, self.f12, globalState)) and (self.f12)
                index93_ = default__.safeIndex(387, (d_738_v46_).length(0))
                rhs119_ = (d_739_v47_).set(default__.safeIndex(self.f16, len(d_739_v47_)), default__.fm14(d_672_v2_, self.f16, globalState))
                rhs120_ = d_744_v51_
                rhs121_ = ((d_739_v47_).set(default__.safeIndex(self.f16, len(d_739_v47_)), _dafny.CodePoint('c'))) < (((d_739_v47_).set(default__.safeIndex(self.f16, len(d_739_v47_)), d_671_v1_)) + (p1))
                lhs94_ = d_738_v46_
                lhs95_ = default__.safeIndex(387, (d_738_v46_).length(0))
                d_739_v47_ = rhs119_
                d_744_v51_ = rhs120_
                lhs94_[lhs95_] = rhs121_
            elif True:
                d_749_v54_: _dafny.Map
                d_749_v54_ = _dafny.Map({self.f16: self.f12})
                d_671_v1_ = (d_671_v1_ if (d_749_v54_) == (d_749_v54_) else d_671_v1_)
                d_750_v55_: _dafny.Array
                def lambda31_(d_751_i8_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nol"))

                init18_ = lambda31_
                nw132_ = _dafny.Array(None, 1)
                for i0_18_ in range(nw132_.length(0)):
                    nw132_[i0_18_] = init18_(i0_18_)
                d_750_v55_ = nw132_
                index94_ = default__.safeIndex(481, (d_750_v55_).length(0))
                (d_750_v55_)[index94_] = p1
                index95_ = default__.safeIndex(481, (d_750_v55_).length(0))
                (d_750_v55_)[index95_] = p1
                (globalState).f9 = (self.f16) * (self.f16)
                d_752_v56_: _dafny.Seq
                d_752_v56_ = _dafny.SeqWithoutIsStrInference([self.f16, 278])
                d_753_v57_: D5
                d_753_v57_ = D5_DC12(d_752_v56_)
                d_752_v56_ = ((d_753_v57_).cf29).set(default__.safeIndex(((_dafny.MultiSet([self.f12, self.f12])).cardinality if self.f12 else (0) - (len((d_750_v55_)[default__.safeIndex(481, (d_750_v55_).length(0))]))), len((d_753_v57_).cf29)), self.f16)
                d_754_v58_: _dafny.Seq
                d_754_v58_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f12])
                (globalState).f9 = default__.safeDivisionInt(self.f16, (self.f16) - (len(d_754_v58_)))
            d_755_v59_: _dafny.Array
            nw133_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 29)
            d_755_v59_ = nw133_
            r0 = d_755_v59_
            source8_ = d_670_v0_
            d_756___mcc_h7_ = source8_
            d_757_cf0_ = d_756___mcc_h7_
            d_758_v60_: _dafny.Array
            nw134_ = _dafny.Array(int(0), 24)
            d_758_v60_ = nw134_
            index96_ = default__.safeIndex(693, (d_758_v60_).length(0))
            (d_758_v60_)[index96_] = self.f16
            d_759_v61_: _dafny.Seq
            d_759_v61_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jljlfkl"))
            d_760_v62_: _dafny.Array
            nw135_ = _dafny.Array(D1.default()(), 21)
            d_760_v62_ = nw135_
            d_761_v63_: D1
            d_761_v63_ = D1_DC3(default__.fm15(len(d_759_v61_), self.f16, d_757_cf0_, globalState))
            d_762_v64_: D1
            d_762_v64_ = D1_DC3(d_761_v63_)
            index97_ = default__.safeIndex(188, (d_760_v62_).length(0))
            (d_760_v62_)[index97_] = d_762_v64_
            d_763_v65_: D3
            d_763_v65_ = D3_DC8(self.f16, default__.fm1(True, self.f16, self.f16, globalState))
            d_764_v66_: _dafny.Seq
            d_764_v66_ = _dafny.SeqWithoutIsStrInference([True, self.f12])
            d_765_v67_: _dafny.Map
            d_765_v67_ = _dafny.Map({D3_DC8(self.f16, len(_dafny.SeqWithoutIsStrInference([self.f16, len(d_764_v66_)]))): self.f16})
            d_766_v68_: _dafny.Seq
            d_766_v68_ = _dafny.SeqWithoutIsStrInference([self.f16])
            d_767_v69_: _dafny.Set
            d_767_v69_ = _dafny.Set({default__.fm9(self.f16, self.f12, globalState)})
            d_768_v70_: _dafny.MultiSet
            d_768_v70_ = _dafny.MultiSet([p1, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sarchm")), p1])
            d_769_v71_: _dafny.Array
            nw136_ = _dafny.Array(None, 25)
            nw136_[int(0)] = d_671_v1_
            nw136_[int(1)] = d_671_v1_
            nw136_[int(2)] = d_671_v1_
            nw136_[int(3)] = d_671_v1_
            nw136_[int(4)] = d_671_v1_
            nw136_[int(5)] = d_671_v1_
            nw136_[int(6)] = default__.fm14(d_672_v2_, self.f16, globalState)
            nw136_[int(7)] = d_671_v1_
            nw136_[int(8)] = d_671_v1_
            nw136_[int(9)] = d_671_v1_
            nw136_[int(10)] = d_671_v1_
            nw136_[int(11)] = d_671_v1_
            nw136_[int(12)] = d_671_v1_
            nw136_[int(13)] = d_671_v1_
            nw136_[int(14)] = d_671_v1_
            nw136_[int(15)] = default__.fm14(default__.fm15(self.f16, self.f16, self.f12, globalState), 651, globalState)
            nw136_[int(16)] = _dafny.CodePoint('o')
            nw136_[int(17)] = d_671_v1_
            nw136_[int(18)] = d_671_v1_
            nw136_[int(19)] = _dafny.CodePoint('p')
            nw136_[int(20)] = d_671_v1_
            nw136_[int(21)] = d_671_v1_
            nw136_[int(22)] = d_671_v1_
            nw136_[int(23)] = (d_759_v61_)[default__.safeIndex(self.f16, len(d_759_v61_))]
            nw136_[int(24)] = d_671_v1_
            d_769_v71_ = nw136_
            d_770_v72_: _dafny.Map
            d_770_v72_ = _dafny.Map({d_769_v71_: False})
            d_771_v73_: _dafny.Set
            d_771_v73_ = _dafny.Set({False, d_757_cf0_, not(d_757_cf0_)})
            d_772_v74_: _dafny.Set
            d_772_v74_ = _dafny.Set({self.f16})
            d_773_v75_: D4
            d_773_v75_ = D4_DC11(d_757_cf0_, self.f16, default__.fm1(self.f12, (0) - (len(_dafny.SeqWithoutIsStrInference([self.f16]))), self.f16, globalState), len(d_772_v74_))
            d_774_v76_: _dafny.Map
            d_774_v76_ = _dafny.Map({d_759_v61_: len(d_766_v68_)})
            d_775_v77_: C0
            nw137_ = C0()
            nw137_.ctor__(d_757_cf0_, len(d_774_v76_), self.f11, (d_764_v66_)[default__.safeIndex(((p0)[(self.f11).cardinality] if ((self.f11).cardinality) in (p0) else self.f16), len(d_764_v66_))])
            d_775_v77_ = nw137_
            d_776_v78_: _dafny.Set
            d_776_v78_ = _dafny.Set({d_775_v77_, d_775_v77_, d_775_v77_, d_775_v77_})
            d_777_v79_: _dafny.Array
            nw138_ = _dafny.Array(None, 11)
            nw138_[int(0)] = (d_763_v65_) in (d_765_v67_)
            nw138_[int(1)] = (self.f16) <= (len(d_766_v68_))
            nw138_[int(2)] = (-92) == (len(d_767_v69_))
            nw138_[int(3)] = self.f12
            nw138_[int(4)] = (d_759_v61_) in (d_768_v70_)
            nw138_[int(5)] = ((d_770_v72_)[d_769_v71_] if (d_769_v71_) in (d_770_v72_) else self.f12)
            nw138_[int(6)] = d_757_cf0_
            nw138_[int(7)] = (d_771_v73_).issubset(_dafny.Set({(d_773_v75_).cf25}))
            nw138_[int(8)] = d_757_cf0_
            nw138_[int(9)] = d_757_cf0_
            nw138_[int(10)] = (len(d_776_v78_)) > ((d_775_v77_).f18)
            d_777_v79_ = nw138_
            d_778_v80_: _dafny.Map
            d_778_v80_ = _dafny.Map({(d_775_v77_).f17: 53})
            index98_ = default__.safeIndex(765, (d_777_v79_).length(0))
            (d_777_v79_)[index98_] = (d_775_v77_).fm2(p0, self.f16, d_778_v80_, (d_775_v77_).f17, globalState)
            d_779_v81_: _dafny.Map
            d_779_v81_ = _dafny.Map({p1: (d_775_v77_).f17})
            index99_ = default__.safeIndex(693, (d_758_v60_).length(0))
            index100_ = default__.safeIndex(188, (d_760_v62_).length(0))
            index101_ = default__.safeIndex(765, (d_777_v79_).length(0))
            rhs122_ = self.f16
            rhs123_ = d_759_v61_
            rhs124_ = d_762_v64_
            rhs125_ = d_777_v79_
            rhs126_ = not(((d_779_v81_)[p1] if (p1) in (d_779_v81_) else d_757_cf0_))
            lhs96_ = d_758_v60_
            lhs97_ = default__.safeIndex(693, (d_758_v60_).length(0))
            lhs98_ = d_760_v62_
            lhs99_ = default__.safeIndex(188, (d_760_v62_).length(0))
            lhs100_ = d_777_v79_
            lhs101_ = default__.safeIndex(765, (d_777_v79_).length(0))
            lhs96_[lhs97_] = rhs122_
            d_759_v61_ = rhs123_
            lhs98_[lhs99_] = rhs124_
            d_777_v79_ = rhs125_
            lhs100_[lhs101_] = rhs126_
            index102_ = default__.safeIndex(693, (d_758_v60_).length(0))
            (d_758_v60_)[index102_] = -18
            index103_ = default__.safeIndex(736, (d_755_v59_).length(0))
            (d_755_v59_)[index103_] = p1
            index104_ = default__.safeIndex(736, (d_755_v59_).length(0))
            (d_755_v59_)[index104_] = default__.fm9((d_758_v60_)[default__.safeIndex(693, (d_758_v60_).length(0))], ((0) - ((d_758_v60_)[default__.safeIndex(693, (d_758_v60_).length(0))])) >= (((d_778_v80_)[(d_775_v77_).f17] if ((d_775_v77_).f17) in (d_778_v80_) else self.f16)), globalState)
            d_672_v2_ = d_672_v2_
            d_780_v82_: _dafny.Array
            nw139_ = _dafny.Array(_dafny.CodePoint('D'), 11)
            d_780_v82_ = nw139_
            index105_ = default__.safeIndex(113, (d_780_v82_).length(0))
            (d_780_v82_)[index105_] = d_671_v1_
            index106_ = default__.safeIndex(113, (d_780_v82_).length(0))
            (d_780_v82_)[index106_] = (d_671_v1_ if self.f12 else d_671_v1_)
        hi1_ = self.f16
        for d_781_i9_ in range(self.f16, hi1_):
            d_782_v83_: _dafny.Map
            d_782_v83_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([self.f16 for d_783_i10_ in range(default__.abs(434))])): self.f12})
            d_784_v84_: D4
            d_784_v84_ = D4_DC11(not(False), d_781_i9_, d_781_i9_, len(d_782_v83_))
            d_785_v85_: C0
            nw140_ = C0()
            nw140_.ctor__(((d_784_v84_).cf25) == (self.f12), (0) - (d_781_i9_), self.f11, False)
            d_785_v85_ = nw140_
            d_786_v86_: _dafny.Array
            nw141_ = _dafny.Array(_dafny.Map({}), 22)
            d_786_v86_ = nw141_
            d_787_v87_: _dafny.Seq
            d_787_v87_ = _dafny.SeqWithoutIsStrInference([self.f16, (d_785_v85_).f18])
            d_788_v88_: _dafny.Seq
            d_788_v88_ = _dafny.SeqWithoutIsStrInference([len(d_787_v87_)])
            d_789_v89_: _dafny.Map
            d_789_v89_ = _dafny.Map({self.f12: d_788_v88_})
            index107_ = default__.safeIndex(225, (d_786_v86_).length(0))
            (d_786_v86_)[index107_] = (d_789_v89_) | (d_789_v89_)
            d_790_v90_: _dafny.Array
            nw142_ = _dafny.Array(False, 2)
            d_790_v90_ = nw142_
            d_791_v91_: _dafny.Seq
            d_791_v91_ = _dafny.SeqWithoutIsStrInference([d_790_v90_, d_790_v90_])
            index108_ = default__.safeIndex(225, (d_786_v86_).length(0))
            (d_786_v86_)[index108_] = ((d_789_v89_ if self.f12 else d_789_v89_) if (d_791_v91_) == (d_791_v91_) else (d_789_v89_) | (d_789_v89_))
            (self).f12 = ((d_785_v85_).f17) and (((d_785_v85_).f17) and (self.f12))
            d_792_v92_: _dafny.Map
            d_792_v92_ = _dafny.Map({(d_785_v85_).f18: _dafny.SeqWithoutIsStrInference([(d_785_v85_).f18 for d_793_i11_ in range(default__.abs(66))])})
            d_794_v93_: _dafny.Set
            d_794_v93_ = _dafny.Set({d_787_v87_, d_787_v87_, ((d_792_v92_)[d_781_i9_] if (d_781_i9_) in (d_792_v92_) else d_788_v88_)})
            d_795_v94_: _dafny.Map
            d_795_v94_ = _dafny.Map({d_794_v93_: self.f12})
            d_795_v94_ = (d_795_v94_).set(d_794_v93_, (d_785_v85_).f17)
        d_796_v95_: _dafny.Array
        nw143_ = _dafny.Array(False, 20)
        d_796_v95_ = nw143_
        index109_ = default__.safeIndex(617, (d_796_v95_).length(0))
        (d_796_v95_)[index109_] = self.f12
        index110_ = default__.safeIndex(617, (d_796_v95_).length(0))
        (d_796_v95_)[index110_] = (self.f11).ispropersubset((self.f11) | (_dafny.MultiSet([self.f12, True])))
        d_797_v96_: _dafny.Array
        nw144_ = _dafny.Array(None, 16)
        nw144_[int(0)] = _dafny.CodePoint('b')
        nw144_[int(1)] = d_671_v1_
        nw144_[int(2)] = d_671_v1_
        nw144_[int(3)] = d_671_v1_
        nw144_[int(4)] = d_671_v1_
        nw144_[int(5)] = d_671_v1_
        nw144_[int(6)] = d_671_v1_
        nw144_[int(7)] = default__.fm14(d_672_v2_, (0) - (self.f16), globalState)
        nw144_[int(8)] = d_671_v1_
        nw144_[int(9)] = d_671_v1_
        nw144_[int(10)] = d_671_v1_
        nw144_[int(11)] = d_671_v1_
        nw144_[int(12)] = d_671_v1_
        nw144_[int(13)] = d_671_v1_
        nw144_[int(14)] = d_671_v1_
        nw144_[int(15)] = d_671_v1_
        d_797_v96_ = nw144_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_797_v96_).length(0)):
            d_798_i12_: int = guard_loop_1_
            if (True) and (((0) <= (d_798_i12_)) and ((d_798_i12_) < ((d_797_v96_).length(0)))):
                (d_797_v96_)[(d_798_i12_)] = d_671_v1_
        d_799_v97_: _dafny.Array
        def lambda32_(d_800_p1_):
            def lambda33_(d_801_i13_):
                return d_800_p1_

            return lambda33_

        init19_ = lambda32_(p1)
        nw145_ = _dafny.Array(None, 3)
        for i0_19_ in range(nw145_.length(0)):
            nw145_[i0_19_] = init19_(i0_19_)
        d_799_v97_ = nw145_
        d_802_v98_: D5
        d_802_v98_ = D5_DC13(self.f16, (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], self.f16, d_799_v97_, (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))])
        d_803_v99_: _dafny.Seq
        d_803_v99_ = _dafny.SeqWithoutIsStrInference([(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], self.f12, False])
        d_804_v100_: _dafny.Array
        nw146_ = _dafny.Array(None, 3)
        nw146_[int(0)] = p0
        nw146_[int(1)] = p0
        nw146_[int(2)] = ((p0).set((d_802_v98_).cf32, default__.abs((0) - (self.f16)))).set(len((d_803_v99_).set(default__.safeIndex(142, len(d_803_v99_)), (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))])), default__.abs(self.f16))
        d_804_v100_ = nw146_
        index111_ = default__.safeIndex(505, (d_804_v100_).length(0))
        (d_804_v100_)[index111_] = (p0) - (p0)
        index112_ = default__.safeIndex(505, (d_804_v100_).length(0))
        (d_804_v100_)[index112_] = (p0).intersection(p0)
        if (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]:
            (self).f12 = (default__.safeDivisionInt((_dafny.MultiSet([(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]])).cardinality, self.f16)) <= (self.f16)
            (globalState).f9 = self.f16
            (globalState).f7 = default__.fm0(d_671_v1_, globalState)
            d_805_v101_: D3
            d_805_v101_ = D3_DC6(d_796_v95_)
            source9_ = d_805_v101_
            if source9_.is_DC7:
                d_806___mcc_h8_ = source9_.cf15
                d_807___mcc_h9_ = source9_.cf16
                d_808___mcc_h10_ = source9_.cf17
                d_809_cf17_ = d_808___mcc_h10_
                d_810_cf16_ = d_807___mcc_h9_
                d_811_cf15_ = d_806___mcc_h8_
                (globalState).f7 = True
                rhs127_ = (0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eokqyem"))), d_810_cf16_))
                rhs128_ = default__.safeModuloInt((d_810_cf16_) + (d_811_cf15_), self.f16)
                rhs129_ = d_809_cf17_
                rhs130_ = (0) - ((d_809_cf17_).f18)
                lhs102_ = globalState
                lhs103_ = globalState
                d_811_cf15_ = rhs127_
                lhs102_.f9 = rhs128_
                d_809_cf17_ = rhs129_
                lhs103_.f9 = rhs130_
                d_812_v102_: _dafny.Map
                d_812_v102_ = _dafny.Map({d_811_cf15_: self.f16})
                d_813_v103_: _dafny.Set
                d_813_v103_ = _dafny.Set({d_811_cf15_, (d_809_cf17_).f18, d_810_cf16_, (228) * (len(d_812_v102_)), (d_809_cf17_).fm12(globalState)})
                d_813_v103_ = d_813_v103_
                d_814_v104_: _dafny.Set
                d_814_v104_ = _dafny.Set({(d_809_cf17_).f17})
                d_815_v105_: _dafny.Map
                d_815_v105_ = _dafny.Map({True: self.f12})
                d_816_v106_: _dafny.Set
                d_816_v106_ = _dafny.Set({self.f12, (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], not(self.f12), (default__.fm16(len(d_814_v104_), globalState)) != (d_815_v105_)})
                d_817_v107_: _dafny.Seq
                d_817_v107_ = _dafny.SeqWithoutIsStrInference([d_811_cf15_, ((d_804_v100_)[default__.safeIndex(505, (d_804_v100_).length(0))]).cardinality])
                index113_ = default__.safeIndex(617, (d_796_v95_).length(0))
                def iife62_():
                    coll20_ = _dafny.Set()
                    compr_20_: int
                    for compr_20_ in ((d_817_v107_).set(default__.safeIndex(self.f16, len(d_817_v107_)), d_811_cf15_)).Elements:
                        d_818_v108_: int = compr_20_
                        if (d_818_v108_) in ((d_817_v107_).set(default__.safeIndex(self.f16, len(d_817_v107_)), d_811_cf15_)):
                            coll20_ = coll20_.union(_dafny.Set([(d_818_v108_) - (992)]))
                    return _dafny.Set(coll20_)
                rhs131_ = ((d_813_v103_) - (d_813_v103_)).ispropersubset(iife62_()
                )
                rhs132_ = len(d_812_v102_)
                rhs133_ = (d_816_v106_).intersection(_dafny.Set({(d_809_cf17_).f17}))
                rhs134_ = (d_814_v104_).issubset(_dafny.Set({self.f12}))
                lhs104_ = d_796_v95_
                lhs105_ = default__.safeIndex(617, (d_796_v95_).length(0))
                lhs106_ = globalState
                lhs104_[lhs105_] = rhs131_
                d_810_cf16_ = rhs132_
                d_816_v106_ = rhs133_
                lhs106_.f8 = rhs134_
            elif source9_.is_DC8:
                d_819___mcc_h11_ = source9_.cf18
                d_820___mcc_h12_ = source9_.cf19
                d_821_cf19_ = d_820___mcc_h12_
                d_822_cf18_ = d_819___mcc_h11_
                index114_ = default__.safeIndex(474, (d_799_v97_).length(0))
                (d_799_v97_)[index114_] = p1
                pat_let_tv16_ = d_796_v95_
                index115_ = default__.safeIndex(474, (d_799_v97_).length(0))
                def iife63_(_pat_let21_0):
                    def iife64_(d_823_dt__update__tmp_h1_):
                        def iife65_(_pat_let22_0):
                            def iife66_(d_824_dt__update_hcf14_h0_):
                                return D3_DC6(d_824_dt__update_hcf14_h0_)
                            return iife66_(_pat_let22_0)
                        return iife65_(pat_let_tv16_)
                    return iife64_(_pat_let21_0)
                rhs135_ = p1
                rhs136_ = ((_dafny.SeqWithoutIsStrInference([(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], self.f12]) if not(not(False)) else d_803_v99_)) + (d_803_v99_)
                rhs137_ = not(default__.fm0(d_671_v1_, globalState))
                rhs138_ = self.f12
                rhs139_ = (iife63_(d_805_v101_)).cf14
                lhs107_ = d_799_v97_
                lhs108_ = default__.safeIndex(474, (d_799_v97_).length(0))
                lhs109_ = globalState
                lhs110_ = globalState
                lhs111_ = globalState
                lhs107_[lhs108_] = rhs135_
                lhs109_.f6 = rhs136_
                lhs110_.f7 = rhs137_
                lhs111_.f3 = rhs138_
                d_796_v95_ = rhs139_
                d_825_v109_: T0
                nw147_ = C0()
                nw147_.ctor__(False, self.f16, self.f11, (d_803_v99_)[default__.safeIndex(d_821_cf19_, len(d_803_v99_))])
                d_825_v109_ = nw147_
                d_826_v110_: _dafny.Map
                d_826_v110_ = _dafny.Map({d_821_cf19_: d_825_v109_})
                d_827_v111_: _dafny.Map
                d_827_v111_ = _dafny.Map({(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]: self.f16})
                (globalState).f7 = (self).fm2(default__.fm13(d_671_v1_, (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], globalState), default__.safeModuloInt(len(d_826_v110_), d_822_cf18_), (d_827_v111_) | (d_827_v111_), (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], globalState)
                d_828_v112_: _dafny.Seq
                d_828_v112_ = _dafny.SeqWithoutIsStrInference([d_822_cf18_, d_822_cf18_])
                d_829_v113_: _dafny.Array
                nw148_ = _dafny.Array(int(0), 25)
                d_829_v113_ = nw148_
                d_830_v114_: _dafny.Map
                d_830_v114_ = _dafny.Map({_dafny.MultiSet(((d_828_v112_ if self.f12 else d_828_v112_)).set(default__.safeIndex(self.f16, len((d_828_v112_ if self.f12 else d_828_v112_))), self.f16)): d_829_v113_})
                d_830_v114_ = (d_830_v114_).set((p0) - ((d_804_v100_)[default__.safeIndex(505, (d_804_v100_).length(0))]), d_829_v113_)
                d_831_v115_: C0
                nw149_ = C0()
                nw149_.ctor__((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], self.f16, (d_825_v109_.f11) | (self.f11), ((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]) or ((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]))
                d_831_v115_ = nw149_
                d_831_v115_ = d_831_v115_
            elif source9_.is_DC9:
                d_832___mcc_h13_ = source9_.cf20
                d_833___mcc_h14_ = source9_.cf21
                d_834___mcc_h15_ = source9_.cf22
                d_835___mcc_h16_ = source9_.cf23
                d_836_cf23_ = d_835___mcc_h16_
                d_837_cf22_ = d_834___mcc_h15_
                d_838_cf21_ = d_833___mcc_h14_
                d_839_cf20_ = d_832___mcc_h13_
                d_840_v116_: _dafny.Array
                nw150_ = _dafny.Array(int(0), 16)
                d_840_v116_ = nw150_
                index116_ = default__.safeIndex(12, (d_840_v116_).length(0))
                (d_840_v116_)[index116_] = self.f16
                d_841_v117_: _dafny.Map
                d_841_v117_ = _dafny.Map({p1: self.f12})
                index117_ = default__.safeIndex(12, (d_840_v116_).length(0))
                (d_840_v116_)[index117_] = default__.safeModuloInt(len(d_841_v117_), (d_837_cf22_).f18)
                d_842_v118_: D1
                d_842_v118_ = D1_DC2(self.f12, (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], self.f12, d_671_v1_, d_671_v1_)
                d_843_v119_: D1
                d_843_v119_ = D1_DC3(d_842_v118_)
                d_844_v120_: D1
                d_844_v120_ = D1_DC3(d_843_v119_)
                d_845_v121_: D1
                d_845_v121_ = D1_DC3(d_842_v118_)
                pat_let_tv17_ = d_842_v118_
                d_846_v122_: _dafny.Set
                def iife67_(_pat_let23_0):
                    def iife68_(d_847_dt__update__tmp_h2_):
                        def iife69_(_pat_let24_0):
                            def iife70_(d_848_dt__update_hcf7_h0_):
                                return D1_DC3(d_848_dt__update_hcf7_h0_)
                            return iife70_(_pat_let24_0)
                        return iife69_(pat_let_tv17_)
                    return iife68_(_pat_let23_0)
                d_846_v122_ = _dafny.Set({iife67_(D1_DC3(d_844_v120_)), d_845_v121_})
                d_849_v123_: _dafny.MultiSet
                d_849_v123_ = _dafny.MultiSet([(d_846_v122_) | (d_846_v122_), d_846_v122_])
                d_849_v123_ = (d_849_v123_) - ((d_849_v123_) - (default__.fm17((d_837_cf22_).f17, not(False), (d_837_cf22_).f18, 190, globalState)))
                d_839_cf20_ = (0) - (d_839_cf20_)
                d_850_v124_: _dafny.Map
                d_850_v124_ = _dafny.Map({self.f12: d_796_v95_})
                d_850_v124_ = (d_850_v124_).set((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], d_796_v95_)
            elif True:
                d_851___mcc_h17_ = source9_.cf14
                d_852_cf14_ = d_851___mcc_h17_
                d_853_v125_: _dafny.Set
                d_853_v125_ = _dafny.Set({self.f16})
                d_854_v126_: _dafny.Map
                d_854_v126_ = _dafny.Map({self.f12: self.f16})
                d_855_v127_: _dafny.Set
                d_855_v127_ = _dafny.Set({self.f12, (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]})
                d_856_v128_: D2
                d_856_v128_ = D2_DC4(self.f16)
                d_857_v129_: _dafny.Array
                nw151_ = _dafny.Array(None, 6)
                nw151_[int(0)] = self.f16
                nw151_[int(1)] = self.f16
                nw151_[int(2)] = (0) - (len(_dafny.Map({(D2_DC5(len(d_803_v99_), self.f16, self.f16, self.f16, d_803_v99_)).cf9: self.f12})))
                nw151_[int(3)] = len((d_853_v125_) - (_dafny.Set({(0) - (self.f16)})))
                nw151_[int(4)] = (-623) - (len((d_854_v126_).set(not(self.f12), default__.fm1(not((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]), len(d_855_v127_), self.f16, globalState))))
                nw151_[int(5)] = default__.fm1(self.f12, (d_856_v128_).cf8, 797, globalState)
                d_857_v129_ = nw151_
                index118_ = default__.safeIndex(273, (d_857_v129_).length(0))
                (d_857_v129_)[index118_] = self.f16
                d_858_v130_: _dafny.Map
                d_858_v130_ = _dafny.Map({_dafny.CodePoint('o'): (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]})
                d_859_v131_: _dafny.MultiSet
                d_859_v131_ = _dafny.MultiSet([d_671_v1_, (d_672_v2_).cf5, d_671_v1_, (p1)[default__.safeIndex(self.f16, len(p1))]])
                d_860_v132_: _dafny.Seq
                d_860_v132_ = _dafny.SeqWithoutIsStrInference([self.f16, len((d_858_v130_) | (d_858_v130_)), default__.safeModuloInt(len((d_803_v99_).set(default__.safeIndex(671, len(d_803_v99_)), (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))])), self.f16), self.f16, ((d_859_v131_) - (d_859_v131_)).cardinality])
                index119_ = default__.safeIndex(273, (d_857_v129_).length(0))
                (d_857_v129_)[index119_] = (d_860_v132_)[default__.safeIndex(self.f16, len(d_860_v132_))]
                d_861_v133_: C0
                nw152_ = C0()
                nw152_.ctor__(self.f12, self.f16, (_dafny.MultiSet([(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], self.f12])) | (self.f11), (self.f16) == (self.f16))
                d_861_v133_ = nw152_
                index120_ = default__.safeIndex(273, (d_857_v129_).length(0))
                (d_857_v129_)[index120_] = self.f16
                index121_ = default__.safeIndex(273, (d_857_v129_).length(0))
                (d_857_v129_)[index121_] = (d_857_v129_)[default__.safeIndex(273, (d_857_v129_).length(0))]
            d_862_v135_: _dafny.Set
            d_862_v135_ = _dafny.Set({self.f16})
            d_863_v136_: C0
            nw153_ = C0()
            def iife71_():
                coll21_ = _dafny.Set()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(447, 623):
                    d_864_v134_: int = compr_21_
                    if ((447) <= (d_864_v134_)) and ((d_864_v134_) < (623)):
                        coll21_ = coll21_.union(_dafny.Set([default__.safeModuloInt(d_864_v134_, self.f16)]))
                return _dafny.Set(coll21_)
            nw153_.ctor__(False, self.f16, _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]])).set(default__.safeIndex(self.f16, len(_dafny.SeqWithoutIsStrInference([(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]]))), self.f12)), (d_862_v135_).ispropersubset(iife71_()
            ))
            d_863_v136_ = nw153_
        elif True:
            (globalState).f9 = self.f16
            d_865_v137_: _dafny.Map
            d_865_v137_ = _dafny.Map({not((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]): self.f16})
            (globalState).f3 = (default__.safeModuloInt(-201, ((d_865_v137_)[(d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]] if ((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]) in (d_865_v137_) else self.f16))) != (len(_dafny.SeqWithoutIsStrInference([(0) - (self.f16) for d_866_i14_ in range(default__.abs(606))])))
            if ((self.f16) == (default__.fm1(self.f12, self.f16, self.f16, globalState))) not in (_dafny.Map({self.f12: self.f16})):
                d_867_v138_: _dafny.Set
                d_867_v138_ = _dafny.Set({self.f12, self.f12})
                d_868_v139_: C0
                nw154_ = C0()
                nw154_.ctor__(self.f12, len(_dafny.Map({self.f16: len(d_867_v138_)})), (self.f11).set(self.f12, default__.abs(len(p1))), self.f12)
                d_868_v139_ = nw154_
                d_869_v140_: D3
                d_869_v140_ = D3_DC9(self.f16, len(default__.fm18(self.f16, self.f12, globalState)), d_868_v139_, (d_868_v139_).f18)
                d_870_v141_: _dafny.Map
                d_870_v141_ = _dafny.Map({self.f16: default__.safeDivisionInt(self.f16, len(_dafny.Map({d_869_v140_: (d_868_v139_).f17})))})
                d_870_v141_ = ((d_870_v141_) | (d_870_v141_)) | (d_870_v141_)
                index122_ = default__.safeIndex(617, (d_796_v95_).length(0))
                (d_796_v95_)[index122_] = True
                d_871_v142_: _dafny.Array
                def lambda34_(d_872_i15_):
                    return (d_872_i15_) * (263)

                init20_ = lambda34_
                nw155_ = _dafny.Array(None, 10)
                for i0_20_ in range(nw155_.length(0)):
                    nw155_[i0_20_] = init20_(i0_20_)
                d_871_v142_ = nw155_
                index123_ = default__.safeIndex(122, (d_871_v142_).length(0))
                (d_871_v142_)[index123_] = self.f16
                d_873_v143_: D3
                d_873_v143_ = D3_DC6(d_796_v95_)
                index124_ = default__.safeIndex(712, (d_871_v142_).length(0))
                (d_871_v142_)[index124_] = self.f16
                index125_ = default__.safeIndex(122, (d_871_v142_).length(0))
                index126_ = default__.safeIndex(712, (d_871_v142_).length(0))
                rhs140_ = d_671_v1_
                rhs141_ = (0) - (default__.safeDivisionInt(self.f16, (self.f16) + ((self.f11).cardinality)))
                rhs142_ = d_873_v143_
                rhs143_ = self.f16
                lhs112_ = d_871_v142_
                lhs113_ = default__.safeIndex(122, (d_871_v142_).length(0))
                lhs114_ = d_871_v142_
                lhs115_ = default__.safeIndex(712, (d_871_v142_).length(0))
                d_671_v1_ = rhs140_
                lhs112_[lhs113_] = rhs141_
                d_873_v143_ = rhs142_
                lhs114_[lhs115_] = rhs143_
                (globalState).f3 = (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]
                index127_ = default__.safeIndex(617, (d_796_v95_).length(0))
                index128_ = default__.safeIndex(505, (d_804_v100_).length(0))
                rhs144_ = self.f12
                rhs145_ = p0
                rhs146_ = (d_868_v139_).fm2(default__.fm13(d_671_v1_, False, globalState), (0) - (self.f16), d_865_v137_, not (self.f12) or ((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]), globalState)
                rhs147_ = (d_868_v139_).f17
                lhs116_ = d_796_v95_
                lhs117_ = default__.safeIndex(617, (d_796_v95_).length(0))
                lhs118_ = d_804_v100_
                lhs119_ = default__.safeIndex(505, (d_804_v100_).length(0))
                lhs120_ = globalState
                lhs121_ = self
                lhs116_[lhs117_] = rhs144_
                lhs118_[lhs119_] = rhs145_
                lhs120_.f7 = rhs146_
                lhs121_.f12 = rhs147_
            elif True:
                d_874_v144_: _dafny.Array
                nw156_ = _dafny.Array(None, 24)
                d_874_v144_ = nw156_
                d_875_v145_: D3
                d_875_v145_ = D3_DC6(d_796_v95_)
                pat_let_tv18_ = d_796_v95_
                pat_let_tv19_ = d_796_v95_
                index129_ = default__.safeIndex(190, (d_874_v144_).length(0))
                def iife73_(_pat_let26_0):
                    def iife74_(d_876_dt__update__tmp_h3_):
                        def iife75_(_pat_let27_0):
                            def iife76_(d_877_dt__update_hcf14_h1_):
                                return D3_DC6(d_877_dt__update_hcf14_h1_)
                            return iife76_(_pat_let27_0)
                        return iife75_(pat_let_tv18_)
                    return iife74_(_pat_let26_0)
                def iife72_(_pat_let25_0):
                    def iife77_(d_878_dt__update__tmp_h4_):
                        def iife78_(_pat_let28_0):
                            def iife79_(d_879_dt__update_hcf14_h2_):
                                return D3_DC6(d_879_dt__update_hcf14_h2_)
                            return iife79_(_pat_let28_0)
                        return iife78_(pat_let_tv19_)
                    return iife77_(_pat_let25_0)
                (d_874_v144_)[index129_] = iife72_(iife73_(d_875_v145_))
                index130_ = default__.safeIndex(190, (d_874_v144_).length(0))
                (d_874_v144_)[index130_] = d_875_v145_
                d_880_v146_: C0
                nw157_ = C0()
                nw157_.ctor__(self.f12, ((d_804_v100_)[default__.safeIndex(505, (d_804_v100_).length(0))]).cardinality, (self.f11).intersection(self.f11), not(not ((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]) or ((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))])))
                d_880_v146_ = nw157_
                d_881_v147_: T0
                nw158_ = C0()
                nw158_.ctor__(((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))] if (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))] else self.f12), ((d_880_v146_).f18) * (self.f16), (self.f11).set(self.f12, default__.abs((d_880_v146_).f18)), self.f12)
                d_881_v147_ = nw158_
                d_881_v147_ = d_881_v147_
                d_882_v148_: C0
                nw159_ = C0()
                nw159_.ctor__((d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))], (d_880_v146_).f18, (d_881_v147_.f11) | (d_881_v147_.f11), self.f12)
                d_882_v148_ = nw159_
                d_883_v149_: _dafny.Seq
                d_883_v149_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggsc"))
                d_883_v149_ = p1
            d_884_v150_: _dafny.Seq
            d_884_v150_ = _dafny.SeqWithoutIsStrInference([d_796_v95_])
            d_885_v151_: _dafny.MultiSet
            d_885_v151_ = _dafny.MultiSet([d_671_v1_])
            d_886_v152_: _dafny.Map
            d_886_v152_ = _dafny.Map({False: d_885_v151_})
            d_887_v153_: _dafny.Array
            nw160_ = _dafny.Array(None, 10)
            nw160_[int(0)] = d_796_v95_
            nw160_[int(1)] = d_796_v95_
            nw160_[int(2)] = d_796_v95_
            nw160_[int(3)] = d_796_v95_
            nw160_[int(4)] = (d_884_v150_)[default__.safeIndex(len(d_886_v152_), len(d_884_v150_))]
            nw160_[int(5)] = d_796_v95_
            nw160_[int(6)] = d_796_v95_
            nw160_[int(7)] = (d_884_v150_)[default__.safeIndex(self.f16, len(d_884_v150_))]
            nw160_[int(8)] = d_796_v95_
            nw160_[int(9)] = d_796_v95_
            d_887_v153_ = nw160_
            index131_ = default__.safeIndex(807, (d_887_v153_).length(0))
            (d_887_v153_)[index131_] = d_796_v95_
            index132_ = default__.safeIndex(807, (d_887_v153_).length(0))
            (d_887_v153_)[index132_] = d_796_v95_
            (globalState).f3 = (d_796_v95_)[default__.safeIndex(617, (d_796_v95_).length(0))]
        r0 = d_799_v97_
        return r0


class C2(T0):
    def  __init__(self):
        self._f11: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        self.f15: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f15, f11, f12):
        (self).f15 = f15
        (self).f11 = f11
        (self).f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        return self.f12

    def fm8(self, p0, p1, globalState):
        return (_dafny.SeqWithoutIsStrInference([622])) + (_dafny.SeqWithoutIsStrInference([959 for d_888_i0_ in range(default__.abs(392))]))

    def m4(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r2: int = int(0)
        r3: bool = False
        if (len(_dafny.SeqWithoutIsStrInference([p0]))) == (p0):
            if p1:
                d_889_v0_: _dafny.Array
                nw161_ = _dafny.Array(False, 16)
                d_889_v0_ = nw161_
                d_889_v0_ = (d_889_v0_ if (self.f12) and (p1) else d_889_v0_)
                r2 = default__.safeDivisionInt(p0, p0)
                d_890_v1_: _dafny.Seq
                d_890_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbffrd"))
                (globalState).f7 = ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_891_i0_ in range(default__.abs(5))])) + (d_890_v1_)) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xycvusgt")))
                d_892_v2_: _dafny.Array
                nw162_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_892_v2_ = nw162_
                d_893_v3_: _dafny.Map
                d_893_v3_ = _dafny.Map({d_892_v2_: self.f15})
                d_893_v3_ = (d_893_v3_).set(d_892_v2_, self.f15)
                d_894_v4_: _dafny.Map
                d_894_v4_ = _dafny.Map({_dafny.CodePoint('j'): self.f12})
                d_895_v5_: _dafny.Set
                d_895_v5_ = _dafny.Set({d_889_v0_})
                d_896_v6_: _dafny.Map
                d_896_v6_ = _dafny.Map({len(d_894_v4_): (d_895_v5_).ispropersubset(d_895_v5_)})
                rhs148_ = (d_890_v1_) + (d_890_v1_)
                rhs149_ = d_896_v6_
                rhs150_ = d_889_v0_
                rhs151_ = self.f12
                rhs152_ = not (self.f15) or (self.f12)
                lhs122_ = globalState
                lhs123_ = globalState
                d_890_v1_ = rhs148_
                d_896_v6_ = rhs149_
                d_889_v0_ = rhs150_
                lhs122_.f7 = rhs151_
                lhs123_.f3 = rhs152_
            elif True:
                d_897_v7_: _dafny.Seq
                d_897_v7_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_897_v7_ = d_897_v7_
                d_898_v8_: _dafny.Array
                nw163_ = _dafny.Array(False, 28)
                d_898_v8_ = nw163_
                index133_ = default__.safeIndex(166, (d_898_v8_).length(0))
                (d_898_v8_)[index133_] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_899_i1_ in range(default__.abs(23))])) != (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_900_i2_ in range(default__.abs(924))]))
                index134_ = default__.safeIndex(166, (d_898_v8_).length(0))
                (d_898_v8_)[index134_] = self.f12
                d_901_v9_: _dafny.Array
                nw164_ = _dafny.Array(int(0), 4)
                d_901_v9_ = nw164_
                index135_ = default__.safeIndex(459, (d_901_v9_).length(0))
                (d_901_v9_)[index135_] = default__.fm1(self.f15, p0, p0, globalState)
                d_902_v10_: _dafny.Map
                d_902_v10_ = _dafny.Map({p0: self.f12})
                index136_ = default__.safeIndex(459, (d_901_v9_).length(0))
                (d_901_v9_)[index136_] = (default__.safeModuloInt(p0, p0)) - (len(d_902_v10_))
                default__.m0(globalState)
                d_903_v11_: _dafny.Map
                d_903_v11_ = _dafny.Map({d_898_v8_: d_901_v9_})
                d_903_v11_ = (d_903_v11_).set(d_898_v8_, d_901_v9_)
            (globalState).f9 = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xswkq"))))
            (globalState).f9 = p0
            r2 = ((self.f11)[not (self.f12) or (self.f15)] if (not (self.f12) or (self.f15)) in (self.f11) else p0)
            d_904_v12_: _dafny.Set
            d_904_v12_ = _dafny.Set({-479, p0, p0})
            d_905_v13_: _dafny.Map
            d_905_v13_ = _dafny.Map({not (p1) or (self.f12): _dafny.Set({len(d_904_v12_)})})
            d_904_v12_ = ((d_905_v13_)[False] if (False) in (d_905_v13_) else d_904_v12_)
        elif True:
            d_906_v14_: _dafny.Seq
            d_906_v14_ = _dafny.SeqWithoutIsStrInference([p0])
            d_907_v16_: C0
            nw165_ = C0()
            def iife80_():
                coll22_ = _dafny.Set()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(279, 299):
                    d_908_v15_: int = compr_22_
                    if ((279) <= (d_908_v15_)) and ((d_908_v15_) < (299)):
                        coll22_ = coll22_.union(_dafny.Set([default__.safeModuloInt(d_908_v15_, p0)]))
                return _dafny.Set(coll22_)
            nw165_.ctor__(self.f12, p0, self.f11, (_dafny.Set({len(d_906_v14_)})) != (iife80_()
            ))
            d_907_v16_ = nw165_
            d_909_v17_: _dafny.Map
            d_909_v17_ = _dafny.Map({(d_907_v16_).f17: p0})
            d_909_v17_ = (d_909_v17_).set(p1, p0)
            d_910_v18_: _dafny.Seq
            d_910_v18_ = _dafny.SeqWithoutIsStrInference([self.f15])
            d_911_v19_: _dafny.Array
            nw166_ = _dafny.Array(None, 2)
            nw166_[int(0)] = (d_907_v16_).f17
            nw166_[int(1)] = not((len(d_910_v18_)) > ((d_907_v16_).f18))
            d_911_v19_ = nw166_
            d_912_v20_: _dafny.Map
            d_912_v20_ = _dafny.Map({(d_907_v16_).f18: self.f12})
            d_913_v21_: _dafny.Map
            d_913_v21_ = _dafny.Map({self.f15: ((d_912_v20_)[694] if (694) in (d_912_v20_) else p1)})
            index137_ = default__.safeIndex(21, (d_911_v19_).length(0))
            (d_911_v19_)[index137_] = ((d_913_v21_)[not(not(True))] if (not(not(True))) in (d_913_v21_) else (d_907_v16_).f17)
            index138_ = default__.safeIndex(21, (d_911_v19_).length(0))
            (d_911_v19_)[index138_] = self.f15
            (globalState).f3 = (p0) == (default__.safeDivisionInt(p0, p0))
            d_909_v17_ = (d_909_v17_).set((d_911_v19_)[default__.safeIndex(21, (d_911_v19_).length(0))], (d_907_v16_).f18)
        d_914_i3_: int
        d_914_i3_ = 0
        with _dafny.label("7"):
            while p1:
                with _dafny.c_label("7"):
                    if (d_914_i3_) >= (100):
                        raise _dafny.Break("7")
                    d_914_i3_ = (d_914_i3_) + (1)
                    d_915_v22_: _dafny.Seq
                    d_915_v22_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_916_v23_: _dafny.Map
                    d_916_v23_ = _dafny.Map({(0) - (p0): True})
                    d_915_v22_ = (self).fm8(((d_916_v23_)[p0] if (p0) in (d_916_v23_) else self.f12), p1, globalState)
                    d_917_v24_: _dafny.Set
                    d_917_v24_ = _dafny.Set({self.f12})
                    d_918_v25_: str
                    d_918_v25_ = _dafny.CodePoint('q')
                    d_919_v26_: _dafny.Map
                    d_919_v26_ = _dafny.Map({(len(d_917_v24_)) - (p0): d_918_v25_})
                    d_919_v26_ = d_919_v26_
                    d_920_v27_: _dafny.MultiSet
                    d_920_v27_ = _dafny.MultiSet([100, p0, p0, 494, p0])
                    (globalState).f9 = (((self.f11).cardinality if not(self.f15) else ((d_920_v27_)[p0] if (p0) in (d_920_v27_) else 316))) * (p0)
                    d_921_v28_: _dafny.Map
                    d_921_v28_ = _dafny.Map({(d_915_v22_).set(default__.safeIndex(-740, len(d_915_v22_)), p0): 954})
                    d_922_v30_: _dafny.MultiSet
                    d_922_v30_ = _dafny.MultiSet([d_915_v22_])
                    d_923_v31_: _dafny.Seq
                    d_923_v31_ = _dafny.SeqWithoutIsStrInference([d_921_v28_])
                    def iife81_():
                        coll23_ = _dafny.Map()
                        compr_23_: _dafny.Seq
                        for compr_23_ in (d_922_v30_).Elements:
                            d_924_v29_: _dafny.Seq = compr_23_
                            if (d_924_v29_) in (d_922_v30_):
                                coll23_[d_924_v29_] = p0
                        return _dafny.Map(coll23_)
                    d_921_v28_ = ((d_921_v28_) | (iife81_()
                    )) | ((d_923_v31_)[default__.safeIndex(p0, len(d_923_v31_))])
                    pass
            pass
        (self).f15 = self.f15
        d_925_v32_: _dafny.Seq
        d_925_v32_ = _dafny.SeqWithoutIsStrInference([self.f12])
        d_926_v33_: _dafny.Seq
        d_926_v33_ = _dafny.SeqWithoutIsStrInference([default__.fm1(p1, p0, p0, globalState), p0, p0])
        d_927_v34_: _dafny.Map
        d_927_v34_ = _dafny.Map({((self.f11).cardinality) * (((self.f11)[(d_925_v32_)[default__.safeIndex(len(d_926_v33_), len(d_925_v32_))]] if ((d_925_v32_)[default__.safeIndex(len(d_926_v33_), len(d_925_v32_))]) in (self.f11) else p0)): p0})
        d_928_v35_: _dafny.Seq
        d_928_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hptlhvxc"))
        d_929_v36_: str
        d_929_v36_ = _dafny.CodePoint('c')
        d_927_v34_ = (d_927_v34_).set((len((d_928_v35_).set(default__.safeIndex((0) - (p0), len(d_928_v35_)), d_929_v36_)) if self.f12 else p0), len(d_928_v35_))
        d_930_v37_: _dafny.Array
        def lambda35_(d_931_i5_):
            return (d_931_i5_) + (758)

        init21_ = lambda35_
        nw167_ = _dafny.Array(None, 29)
        for i0_21_ in range(nw167_.length(0)):
            nw167_[i0_21_] = init21_(i0_21_)
        d_930_v37_ = nw167_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_930_v37_).length(0)):
            d_932_i4_: int = guard_loop_2_
            if (True) and (((0) <= (d_932_i4_)) and ((d_932_i4_) < ((d_930_v37_).length(0)))):
                (d_930_v37_)[(d_932_i4_)] = (d_932_i4_) * ((D2_DC4(-115)).cf8)
        r2 = ((0) - (p0)) - ((p0) + ((0) - (p0)))
        d_933_v38_: _dafny.MultiSet
        d_933_v38_ = _dafny.MultiSet([p0])
        d_934_v39_: D6
        d_934_v39_ = D6_DC15(d_933_v38_)
        d_935_v40_: _dafny.Set
        d_935_v40_ = _dafny.Set({p1, True})
        d_936_v41_: _dafny.Map
        d_936_v41_ = _dafny.Map({self.f12: p0})
        r0 = (self).fm2(((d_934_v39_).cf35).intersection(_dafny.MultiSet([p0, len(d_928_v35_), len(d_935_v40_)])), p0, d_936_v41_, self.f15, globalState)
        r1 = _dafny.SeqWithoutIsStrInference([d_929_v36_ for d_937_i6_ in range(default__.abs(880))])
        r2 = default__.safeModuloInt(p0, default__.fm1(p1, ((self.f11)[self.f15] if (self.f15) in (self.f11) else -55), len(d_935_v40_), globalState))
        r3 = not(self.f12)
        return r0, r1, r2, r3

    def m5(self, p0, p1, globalState):
        r0: bool = False
        r1: int = int(0)
        r1 = p0
        d_938_v0_: str
        d_938_v0_ = _dafny.CodePoint('a')
        d_939_v1_: D1
        d_939_v1_ = D1_DC2(self.f15, self.f12, True, d_938_v0_, d_938_v0_)
        d_940_v2_: _dafny.Map
        d_940_v2_ = _dafny.Map({self.f15: default__.fm14(d_939_v1_, 344, globalState)})
        d_940_v2_ = (d_940_v2_).set(self.f15, d_938_v0_)
        d_941_v3_: _dafny.Array
        nw168_ = _dafny.Array(None, 3)
        nw168_[int(0)] = self.f15
        nw168_[int(1)] = self.f12
        nw168_[int(2)] = self.f15
        d_941_v3_ = nw168_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_941_v3_).length(0)):
            d_942_i0_: int = guard_loop_3_
            if (True) and (((0) <= (d_942_i0_)) and ((d_942_i0_) < ((d_941_v3_).length(0)))):
                (d_941_v3_)[(d_942_i0_)] = not(self.f15)
        d_943_v4_: _dafny.Set
        d_943_v4_ = _dafny.Set({self.f12, self.f15})
        d_944_v5_: D1
        d_944_v5_ = D1_DC1(d_943_v4_)
        d_945_v6_: _dafny.Set
        d_945_v6_ = _dafny.Set({d_944_v5_})
        d_946_v7_: _dafny.Map
        d_946_v7_ = _dafny.Map({self.f15: (d_945_v6_).ispropersubset(_dafny.Set({D1_DC1(d_943_v4_)}))})
        if ((d_946_v7_)[self.f15] if (self.f15) in (d_946_v7_) else self.f15):
            (globalState).f9 = p0
            (self).f12 = self.f12
            d_947_v8_: _dafny.Map
            d_947_v8_ = _dafny.Map({False: p0})
            d_948_v9_: _dafny.Seq
            d_948_v9_ = _dafny.SeqWithoutIsStrInference([p1, ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_949_i1_ in range(default__.abs(591))])).set(default__.safeIndex(((d_947_v8_)[self.f15] if (self.f15) in (d_947_v8_) else p0), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_950_i1_ in range(default__.abs(591))]))), d_938_v0_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_951_i1_ in range(default__.abs(591))])).set(default__.safeIndex(((d_947_v8_)[self.f15] if (self.f15) in (d_947_v8_) else p0), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_952_i1_ in range(default__.abs(591))]))), d_938_v0_))), d_938_v0_)])
            d_953_v10_: _dafny.Seq
            d_953_v10_ = _dafny.SeqWithoutIsStrInference([(0) - (len(p1)), (len(d_948_v9_)) - (p0), (p0 if self.f15 else p0), p0, p0])
            d_954_v11_: C0
            nw169_ = C0()
            nw169_.ctor__(self.f12, p0, self.f11, self.f15)
            d_954_v11_ = nw169_
            (globalState).f9 = (d_953_v10_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_954_v11_, d_954_v11_])), len(d_953_v10_))]
            d_955_v12_: _dafny.Seq
            d_955_v12_ = _dafny.SeqWithoutIsStrInference([(d_953_v10_) + (d_953_v10_), d_953_v10_])
            d_953_v10_ = (d_955_v12_)[default__.safeIndex((d_954_v11_).f18, len(d_955_v12_))]
            d_956_v13_: _dafny.Set
            d_956_v13_ = _dafny.Set({d_938_v0_})
            d_957_v14_: _dafny.Map
            d_957_v14_ = _dafny.Map({D5_DC14(): not(False)})
            d_958_v15_: D7
            d_958_v15_ = D7_DC17(d_957_v14_)
            d_959_v16_: _dafny.Array
            nw170_ = _dafny.Array(None, 26)
            nw170_[int(0)] = p0
            nw170_[int(1)] = p0
            nw170_[int(2)] = (d_954_v11_).f18
            nw170_[int(3)] = p0
            nw170_[int(4)] = (300) * ((d_954_v11_).f18)
            nw170_[int(5)] = (631) - ((d_954_v11_).f18)
            nw170_[int(6)] = (p0) + (p0)
            nw170_[int(7)] = p0
            nw170_[int(8)] = p0
            nw170_[int(9)] = ((d_954_v11_).f18) + (p0)
            nw170_[int(10)] = len(d_943_v4_)
            nw170_[int(11)] = 0
            nw170_[int(12)] = default__.safeModuloInt(len(d_956_v13_), p0)
            nw170_[int(13)] = (0) - (len((d_958_v15_).cf39))
            nw170_[int(14)] = p0
            nw170_[int(15)] = p0
            nw170_[int(16)] = default__.fm1(self.f12, (d_954_v11_).f18, (d_954_v11_).f18, globalState)
            nw170_[int(17)] = (d_954_v11_).f18
            nw170_[int(18)] = ((self.f11)[self.f12] if (self.f12) in (self.f11) else (d_954_v11_).f18)
            nw170_[int(19)] = (d_954_v11_).f18
            nw170_[int(20)] = (d_954_v11_).f18
            nw170_[int(21)] = len(p1)
            nw170_[int(22)] = ((0) - ((d_954_v11_).f18)) * ((d_954_v11_).f18)
            nw170_[int(23)] = ((d_954_v11_).f18) + ((self.f11).cardinality)
            nw170_[int(24)] = (d_954_v11_).f18
            nw170_[int(25)] = (d_954_v11_).f18
            d_959_v16_ = nw170_
            index139_ = default__.safeIndex(371, (d_959_v16_).length(0))
            (d_959_v16_)[index139_] = (d_953_v10_)[default__.safeIndex(p0, len(d_953_v10_))]
            index140_ = default__.safeIndex(371, (d_959_v16_).length(0))
            (d_959_v16_)[index140_] = (0) - ((p0) + (635))
        elif True:
            d_960_v17_: _dafny.Array
            nw171_ = _dafny.Array(_dafny.Seq({}), 17)
            d_960_v17_ = nw171_
            index141_ = default__.safeIndex(387, (d_960_v17_).length(0))
            (d_960_v17_)[index141_] = (self).fm8(self.f15, self.f12, globalState)
            d_961_v18_: _dafny.Seq
            d_961_v18_ = _dafny.SeqWithoutIsStrInference([p0])
            d_962_v19_: D5
            d_962_v19_ = D5_DC12(d_961_v18_)
            index142_ = default__.safeIndex(387, (d_960_v17_).length(0))
            (d_960_v17_)[index142_] = ((D5_DC12(d_961_v18_)).cf29) + ((d_962_v19_).cf29)
            d_963_v20_: _dafny.Array
            def lambda36_(d_964_p0_):
                def lambda37_(d_965_i2_):
                    return default__.safeModuloInt(d_965_i2_, d_964_p0_)

                return lambda37_

            init22_ = lambda36_(p0)
            nw172_ = _dafny.Array(None, 25)
            for i0_22_ in range(nw172_.length(0)):
                nw172_[i0_22_] = init22_(i0_22_)
            d_963_v20_ = nw172_
            d_966_v21_: _dafny.Map
            d_966_v21_ = _dafny.Map({self.f15: d_963_v20_})
            d_967_v22_: _dafny.Seq
            d_967_v22_ = _dafny.SeqWithoutIsStrInference([((d_966_v21_)[self.f15] if (self.f15) in (d_966_v21_) else d_963_v20_)])
            index143_ = default__.safeIndex(260, (d_941_v3_).length(0))
            (d_941_v3_)[index143_] = self.f12
            d_968_v23_: _dafny.Seq
            d_968_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "muarlkhi"))
            index144_ = default__.safeIndex(260, (d_941_v3_).length(0))
            rhs153_ = ((d_967_v22_) + (d_967_v22_)) + (d_967_v22_)
            rhs154_ = (d_946_v7_) == (d_946_v7_)
            rhs155_ = p1
            lhs124_ = d_941_v3_
            lhs125_ = default__.safeIndex(260, (d_941_v3_).length(0))
            d_967_v22_ = rhs153_
            lhs124_[lhs125_] = rhs154_
            d_968_v23_ = rhs155_
            d_969_v24_: _dafny.Map
            d_969_v24_ = _dafny.Map({p0: d_968_v23_})
            d_970_v25_: _dafny.MultiSet
            d_970_v25_ = _dafny.MultiSet([p0])
            rhs156_ = (p1) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_971_i3_ in range(default__.abs(370))]))
            rhs157_ = default__.safeDivisionInt(default__.safeModuloInt(((d_970_v25_)[p0] if (p0) in (d_970_v25_) else p0), p0), default__.safeDivisionInt((0) - ((0) - ((0) - (p0))), len(p1)))
            rhs158_ = ((d_969_v24_).set(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_972_i4_ in range(default__.abs(108))])), p1)) | (d_969_v24_)
            lhs126_ = globalState
            lhs127_ = globalState
            lhs126_.f3 = rhs156_
            lhs127_.f9 = rhs157_
            d_969_v24_ = rhs158_
            d_973_v26_: _dafny.Map
            d_973_v26_ = _dafny.Map({p0: (d_961_v18_)[default__.safeIndex(p0, len(d_961_v18_))]})
            d_974_v27_: _dafny.Set
            d_974_v27_ = _dafny.Set({d_973_v26_, d_973_v26_})
            d_975_v28_: _dafny.Map
            d_975_v28_ = _dafny.Map({self.f12: default__.fm1(not(self.f12), p0, len(d_974_v27_), globalState)})
            d_975_v28_ = d_975_v28_
            d_976_v29_: _dafny.Map
            d_976_v29_ = _dafny.Map({D6_DC15(d_970_v25_): d_970_v25_})
            d_977_v30_: D6
            d_977_v30_ = D6_DC15(d_970_v25_)
            d_978_v31_: _dafny.Set
            d_978_v31_ = _dafny.Set({len(d_975_v28_)})
            d_979_v32_: D4
            d_979_v32_ = D4_DC11((self).fm2(((d_976_v29_)[d_977_v30_] if (d_977_v30_) in (d_976_v29_) else d_970_v25_), p0, d_975_v28_, (d_941_v3_)[default__.safeIndex(260, (d_941_v3_).length(0))], globalState), p0, p0, default__.fm1((d_941_v3_)[default__.safeIndex(260, (d_941_v3_).length(0))], (0) - (len(d_978_v31_)), len(d_946_v7_), globalState))
            d_980_v33_: _dafny.Map
            d_980_v33_ = _dafny.Map({d_963_v20_: (d_979_v32_).cf25})
            d_980_v33_ = (d_980_v33_).set(d_963_v20_, self.f12)
        d_981_v34_: D8
        d_981_v34_ = D8_DC19(self.f11)
        d_982_v35_: _dafny.Seq
        d_982_v35_ = _dafny.SeqWithoutIsStrInference([self.f15])
        d_983_v36_: C1
        nw173_ = C1()
        nw173_.ctor__(len(p1), ((d_981_v34_).cf45).set((d_982_v35_)[default__.safeIndex(-857, len(d_982_v35_))], default__.abs(-462)), self.f12)
        d_983_v36_ = nw173_
        (globalState).f3 = default__.fm0(d_938_v0_, globalState)
        d_984_v37_: _dafny.Set
        d_984_v37_ = _dafny.Set({p0})
        d_985_v38_: bool
        d_985_v38_ = False
        r0 = (d_985_v38_ if (_dafny.Set({p0, d_983_v36_.f16})).issubset(d_984_v37_) else d_985_v38_)
        r1 = (d_983_v36_.f16) + (p0)
        return r0, r1


class C3(T0):
    def  __init__(self):
        self._f11: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        self._f13: int = int(0)
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f13, f14, f11, f12):
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f11 = f11
        (self).f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        return not ((self).f14) or (self.f12)

    def fm6(self, p0, p1, globalState):
        if (self.f12 if True else self.f12):
            if self.f12:
                return (self).f14
            elif True:
                return not(self.f12)
        elif True:
            return ((self).f13) < ((self).f13)

    def m2(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: bool = False
        if (self).f14:
            d_986_v0_: _dafny.Array
            nw174_ = _dafny.Array(_dafny.MultiSet({}), 29)
            d_986_v0_ = nw174_
            d_987_v1_: str
            d_987_v1_ = _dafny.CodePoint('d')
            d_988_v2_: _dafny.MultiSet
            d_988_v2_ = _dafny.MultiSet([_dafny.CodePoint('a'), d_987_v1_])
            d_989_v3_: _dafny.Seq
            d_989_v3_ = _dafny.SeqWithoutIsStrInference([d_988_v2_])
            index145_ = default__.safeIndex(206, (d_986_v0_).length(0))
            (d_986_v0_)[index145_] = (d_989_v3_)[default__.safeIndex((self).f13, len(d_989_v3_))]
            index146_ = default__.safeIndex(206, (d_986_v0_).length(0))
            (d_986_v0_)[index146_] = _dafny.MultiSet((p0) + ((p0) + (((_dafny.SeqWithoutIsStrInference([d_987_v1_, d_987_v1_])).set(default__.safeIndex((self).f13, len(_dafny.SeqWithoutIsStrInference([d_987_v1_, d_987_v1_]))), d_987_v1_)).set(default__.safeIndex((self).f13, len((_dafny.SeqWithoutIsStrInference([d_987_v1_, d_987_v1_])).set(default__.safeIndex((self).f13, len(_dafny.SeqWithoutIsStrInference([d_987_v1_, d_987_v1_]))), d_987_v1_))), _dafny.CodePoint('i')))))
            if p1:
                d_990_v4_: _dafny.Set
                d_990_v4_ = _dafny.Set({(self).f13, (self).f13})
                d_991_v5_: _dafny.Array
                nw175_ = _dafny.Array(None, 5)
                nw175_[int(0)] = self.f12
                nw175_[int(1)] = ((self).f13) != ((self).f13)
                nw175_[int(2)] = (d_990_v4_).isdisjoint(d_990_v4_)
                nw175_[int(3)] = (self).f14
                nw175_[int(4)] = p1
                d_991_v5_ = nw175_
                index147_ = default__.safeIndex(1, (d_991_v5_).length(0))
                (d_991_v5_)[index147_] = (p1) == ((self).f14)
                d_992_v6_: _dafny.Seq
                d_992_v6_ = _dafny.SeqWithoutIsStrInference([d_991_v5_, d_991_v5_, d_991_v5_])
                d_993_v7_: _dafny.Map
                d_993_v7_ = _dafny.Map({d_992_v6_: self.f12})
                index148_ = default__.safeIndex(1, (d_991_v5_).length(0))
                (d_991_v5_)[index148_] = ((d_993_v7_)[d_992_v6_] if (d_992_v6_) in (d_993_v7_) else (False if p1 else p1))
                d_994_v8_: _dafny.Array
                nw176_ = _dafny.Array(int(0), 15)
                d_994_v8_ = nw176_
                index149_ = default__.safeIndex(557, (d_994_v8_).length(0))
                (d_994_v8_)[index149_] = (self).f13
                index150_ = default__.safeIndex(557, (d_994_v8_).length(0))
                (d_994_v8_)[index150_] = (self).f13
                d_995_v9_: _dafny.Map
                d_995_v9_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_994_v8_, d_994_v8_, d_994_v8_])): p0})
                d_995_v9_ = (d_995_v9_).set(((d_994_v8_)[default__.safeIndex(557, (d_994_v8_).length(0))]) * ((self).f13), p0)
                d_996_v10_: D1
                d_996_v10_ = D1_DC2(True, (self).f14, (d_991_v5_)[default__.safeIndex(1, (d_991_v5_).length(0))], d_987_v1_, _dafny.CodePoint('t'))
                d_997_v11_: D1
                d_997_v11_ = D1_DC3(d_996_v10_)
                d_998_v12_: _dafny.Map
                d_998_v12_ = _dafny.Map({-76: d_997_v11_})
                d_998_v12_ = (d_998_v12_).set((self).f13, D1_DC3(d_996_v10_))
                d_999_v13_: bool
                d_999_v13_ = p1
                rhs159_ = d_999_v13_
                rhs160_ = p1
                lhs128_ = globalState
                d_999_v13_ = rhs159_
                lhs128_.f7 = rhs160_
            elif True:
                d_1000_v14_: _dafny.Seq
                d_1000_v14_ = _dafny.SeqWithoutIsStrInference([self.f12])
                d_1001_v15_: D2
                d_1001_v15_ = D2_DC5((self).f13, (self).f13, (self).f13, (self).f13, d_1000_v14_)
                pat_let_tv20_ = globalState
                def iife82_(_pat_let29_0):
                    def iife83_(d_1002_dt__update__tmp_h0_):
                        def iife84_(_pat_let30_0):
                            def iife85_(d_1003_dt__update_hcf13_h0_):
                                def iife86_(_pat_let31_0):
                                    def iife87_(d_1004_dt__update_hcf10_h0_):
                                        return D2_DC5((d_1002_dt__update__tmp_h0_).cf9, d_1004_dt__update_hcf10_h0_, (d_1002_dt__update__tmp_h0_).cf11, (d_1002_dt__update__tmp_h0_).cf12, d_1003_dt__update_hcf13_h0_)
                                    return iife87_(_pat_let31_0)
                                return iife86_((self).f13)
                            return iife85_(_pat_let30_0)
                        return iife84_(default__.fm7(pat_let_tv20_))
                    return iife83_(_pat_let29_0)
                r0 = (len(p0)) * ((iife82_(d_1001_v15_)).cf10)
                (globalState).f9 = (self).f13
                d_1005_v16_: C1
                nw177_ = C1()
                nw177_.ctor__((self).f13, self.f11, p1)
                d_1005_v16_ = nw177_
                d_1006_v17_: int
                out17_: int
                out17_ = (d_1005_v16_).m6(globalState)
                d_1006_v17_ = out17_
                d_1007_v18_: C1
                nw178_ = C1()
                nw178_.ctor__((self).f13, self.f11, (self).f14)
                d_1007_v18_ = nw178_
            if not((self).f14):
                d_1008_v19_: _dafny.Seq
                d_1008_v19_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k')])
                d_1009_v20_: _dafny.Array
                nw179_ = _dafny.Array(int(0), 11)
                d_1009_v20_ = nw179_
                index151_ = default__.safeIndex(265, (d_1009_v20_).length(0))
                (d_1009_v20_)[index151_] = (self).f13
                d_1010_v21_: _dafny.Set
                d_1010_v21_ = _dafny.Set({(self).f13, (self).f13, 0})
                d_1011_v23_: D7
                def iife88_():
                    coll24_ = _dafny.Set()
                    compr_24_: int
                    for compr_24_ in (d_1010_v21_).Elements:
                        d_1013_v22_: int = compr_24_
                        if (d_1013_v22_) in (d_1010_v21_):
                            coll24_ = coll24_.union(_dafny.Set([(d_1013_v22_) + ((_dafny.MultiSet([False])).cardinality)]))
                    return _dafny.Set(coll24_)
                d_1011_v23_ = D7_DC18((self).f13, _dafny.SeqWithoutIsStrInference([d_987_v1_ for d_1012_i0_ in range(default__.abs(787))]), self.f12, (self).f14, iife88_()
)
                index152_ = default__.safeIndex(265, (d_1009_v20_).length(0))
                rhs161_ = (43 if p1 else (self).f13)
                rhs162_ = (d_1011_v23_).cf41
                rhs163_ = ((self).f13) * ((self).f13)
                rhs164_ = not(self.f12)
                lhs129_ = d_1009_v20_
                lhs130_ = default__.safeIndex(265, (d_1009_v20_).length(0))
                lhs131_ = globalState
                r0 = rhs161_
                d_1008_v19_ = rhs162_
                lhs129_[lhs130_] = rhs163_
                lhs131_.f7 = rhs164_
                d_1014_v24_: _dafny.Map
                d_1014_v24_ = _dafny.Map({p1: (self).f14})
                d_1015_v25_: _dafny.Array
                nw180_ = _dafny.Array(None, 10)
                d_1015_v25_ = nw180_
                d_1016_v26_: _dafny.Map
                d_1016_v26_ = _dafny.Map({(_dafny.MultiSet([p1])).ispropersubset(self.f11): d_1015_v25_})
                d_1017_v27_: D4
                d_1017_v27_ = D4_DC11(False, (d_1009_v20_)[default__.safeIndex(265, (d_1009_v20_).length(0))], (self).f13, (self).f13)
                d_1018_v28_: _dafny.Map
                d_1018_v28_ = _dafny.Map({(d_1009_v20_)[default__.safeIndex(265, (d_1009_v20_).length(0))]: d_1009_v20_})
                rhs165_ = d_1014_v24_
                rhs166_ = d_1016_v26_
                rhs167_ = D4_DC11((D8_DC20(p1, self.f12, p1)).cf48, (d_1009_v20_)[default__.safeIndex(265, (d_1009_v20_).length(0))], (d_1009_v20_)[default__.safeIndex(265, (d_1009_v20_).length(0))], len((d_1018_v28_ if (self).f14 else d_1018_v28_)))
                d_1014_v24_ = rhs165_
                d_1016_v26_ = rhs166_
                d_1017_v27_ = rhs167_
                d_1019_v29_: _dafny.Array
                nw181_ = _dafny.Array(_dafny.Seq({}), 23)
                d_1019_v29_ = nw181_
                d_1020_v30_: _dafny.Array
                nw182_ = _dafny.Array(D6.default()(), 2)
                d_1020_v30_ = nw182_
                index153_ = default__.safeIndex(684, (d_1019_v29_).length(0))
                (d_1019_v29_)[index153_] = _dafny.SeqWithoutIsStrInference([d_1020_v30_, d_1020_v30_, d_1020_v30_, d_1020_v30_, d_1020_v30_])
                index154_ = default__.safeIndex(684, (d_1019_v29_).length(0))
                (d_1019_v29_)[index154_] = _dafny.SeqWithoutIsStrInference([d_1020_v30_, d_1020_v30_])
                d_1021_v31_: bool
                d_1021_v31_ = self.f12
                d_1022_v32_: _dafny.Map
                d_1022_v32_ = _dafny.Map({(d_1021_v31_): d_987_v1_})
                (self).m3(((self).f13) - (len(d_1022_v32_)), globalState)
                index155_ = default__.safeIndex(265, (d_1009_v20_).length(0))
                rhs168_ = (self).f13
                rhs169_ = 386
                rhs170_ = self.f11
                rhs171_ = ((d_1009_v20_)[default__.safeIndex(265, (d_1009_v20_).length(0))]) + (default__.safeModuloInt(918, (d_1017_v27_).cf28))
                lhs132_ = globalState
                lhs133_ = self
                lhs134_ = d_1009_v20_
                lhs135_ = default__.safeIndex(265, (d_1009_v20_).length(0))
                r0 = rhs168_
                lhs132_.f9 = rhs169_
                lhs133_.f11 = rhs170_
                lhs134_[lhs135_] = rhs171_
            elif True:
                r1 = False
                (globalState).f9 = (self).f13
                d_1023_v33_: _dafny.Array
                def lambda38_(d_1024_p0_):
                    def lambda39_(d_1025_i1_):
                        return (_dafny.Set({_dafny.Set({len(d_1024_p0_)})})) | (_dafny.Set({_dafny.Set({len(d_1024_p0_), (self).f13})}))

                    return lambda39_

                init23_ = lambda38_(p0)
                nw183_ = _dafny.Array(None, 14)
                for i0_23_ in range(nw183_.length(0)):
                    nw183_[i0_23_] = init23_(i0_23_)
                d_1023_v33_ = nw183_
                d_1026_v34_: _dafny.Map
                d_1026_v34_ = _dafny.Map({(self).f13: self.f12})
                d_1027_v35_: _dafny.Set
                d_1027_v35_ = _dafny.Set({(self).f13, len(d_1026_v34_)})
                d_1028_v36_: _dafny.Seq
                d_1028_v36_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_1029_v37_: T0
                nw184_ = C1()
                nw184_.ctor__((self).f13, _dafny.MultiSet([self.f12, True]), True)
                d_1029_v37_ = nw184_
                d_1030_v38_: _dafny.Seq
                d_1030_v38_ = _dafny.SeqWithoutIsStrInference([d_1029_v37_])
                d_1031_v39_: _dafny.Map
                d_1031_v39_ = _dafny.Map({p1: (d_1030_v38_)[default__.safeIndex((self).f13, len(d_1030_v38_))]})
                d_1032_v40_: D2
                d_1032_v40_ = D2_DC5(len(d_1028_v36_), len(d_1031_v39_), (self).f13, (self).f13, _dafny.SeqWithoutIsStrInference([d_1029_v37_.f12, d_1029_v37_.f12]))
                d_1033_v41_: _dafny.Set
                d_1033_v41_ = _dafny.Set({d_1027_v35_, d_1027_v35_, _dafny.Set({(d_1032_v40_).cf12, len(_dafny.SeqWithoutIsStrInference([d_987_v1_ for d_1034_i2_ in range(default__.abs(706))]))})})
                index156_ = default__.safeIndex(974, (d_1023_v33_).length(0))
                (d_1023_v33_)[index156_] = d_1033_v41_
                d_1035_v42_: _dafny.Array
                nw185_ = _dafny.Array(False, 10)
                d_1035_v42_ = nw185_
                index157_ = default__.safeIndex(974, (d_1023_v33_).length(0))
                rhs172_ = d_1033_v41_
                rhs173_ = d_1035_v42_
                rhs174_ = (p1 if p1 else self.f12)
                lhs136_ = d_1023_v33_
                lhs137_ = default__.safeIndex(974, (d_1023_v33_).length(0))
                lhs138_ = d_1029_v37_
                lhs136_[lhs137_] = rhs172_
                d_1035_v42_ = rhs173_
                lhs138_.f12 = rhs174_
                index158_ = default__.safeIndex(888, (d_1035_v42_).length(0))
                (d_1035_v42_)[index158_] = d_1029_v37_.f12
                index159_ = default__.safeIndex(888, (d_1035_v42_).length(0))
                (d_1035_v42_)[index159_] = p1
                d_1036_v43_: _dafny.MultiSet
                d_1036_v43_ = _dafny.MultiSet([d_1032_v40_, d_1032_v40_])
                d_1037_v44_: _dafny.Seq
                d_1037_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1032_v40_, d_1032_v40_]), (d_1036_v43_).set(d_1032_v40_, default__.abs((self).f13)), d_1036_v43_])
                d_1037_v44_ = (d_1037_v44_) + (_dafny.SeqWithoutIsStrInference([d_1036_v43_ for d_1038_i3_ in range(default__.abs(-912))]))
            d_1039_v45_: _dafny.Array
            nw186_ = _dafny.Array(None, 27)
            d_1039_v45_ = nw186_
            d_1040_v46_: _dafny.Map
            d_1040_v46_ = _dafny.Map({p1: d_1039_v45_})
            d_1041_v47_: _dafny.Seq
            d_1041_v47_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_1042_v48_: _dafny.Map
            d_1042_v48_ = _dafny.Map({default__.fm1((self).f14, (self).f13, len(d_1041_v47_), globalState): _dafny.Map({(self).f14: d_1039_v45_})})
            d_1043_v49_: D9
            d_1043_v49_ = D9_DC23(d_1040_v46_)
            d_1044_v50_: _dafny.Array
            nw187_ = _dafny.Array(None, 25)
            nw187_[int(0)] = d_1040_v46_
            nw187_[int(1)] = (d_1040_v46_ if self.f12 else d_1040_v46_)
            nw187_[int(2)] = (d_1040_v46_).set(True, d_1039_v45_)
            nw187_[int(3)] = d_1040_v46_
            nw187_[int(4)] = (d_1040_v46_) | (d_1040_v46_)
            nw187_[int(5)] = d_1040_v46_
            nw187_[int(6)] = d_1040_v46_
            nw187_[int(7)] = d_1040_v46_
            nw187_[int(8)] = d_1040_v46_
            nw187_[int(9)] = ((d_1042_v48_)[((self.f11)[True] if (True) in (self.f11) else (self).f13)] if (((self.f11)[True] if (True) in (self.f11) else (self).f13)) in (d_1042_v48_) else d_1040_v46_)
            nw187_[int(10)] = (d_1043_v49_).cf55
            nw187_[int(11)] = (d_1040_v46_ if self.f12 else d_1040_v46_)
            nw187_[int(12)] = _dafny.Map({self.f12: d_1039_v45_})
            nw187_[int(13)] = d_1040_v46_
            nw187_[int(14)] = d_1040_v46_
            nw187_[int(15)] = (_dafny.Map({(self).f14: d_1039_v45_})).set(self.f12, d_1039_v45_)
            nw187_[int(16)] = d_1040_v46_
            nw187_[int(17)] = d_1040_v46_
            nw187_[int(18)] = d_1040_v46_
            nw187_[int(19)] = (d_1040_v46_) | (d_1040_v46_)
            nw187_[int(20)] = d_1040_v46_
            nw187_[int(21)] = d_1040_v46_
            nw187_[int(22)] = d_1040_v46_
            nw187_[int(23)] = _dafny.Map({True: d_1039_v45_})
            nw187_[int(24)] = (d_1040_v46_).set(self.f12, d_1039_v45_)
            d_1044_v50_ = nw187_
            index160_ = default__.safeIndex(514, (d_1044_v50_).length(0))
            (d_1044_v50_)[index160_] = (d_1040_v46_ if p1 else d_1040_v46_)
            d_1045_v51_: _dafny.MultiSet
            d_1045_v51_ = _dafny.MultiSet([(0) - ((self).f13)])
            d_1046_v52_: _dafny.Seq
            d_1046_v52_ = _dafny.SeqWithoutIsStrInference([d_1045_v51_, d_1045_v51_])
            d_1047_v53_: _dafny.Map
            d_1047_v53_ = _dafny.Map({(self).f14: (self).f13})
            index161_ = default__.safeIndex(514, (d_1044_v50_).length(0))
            rhs175_ = (self).f13
            rhs176_ = (d_1040_v46_) | ((d_1043_v49_).cf55)
            rhs177_ = not((self).fm2((d_1046_v52_)[default__.safeIndex((0) - (len(d_1041_v47_)), len(d_1046_v52_))], (self).f13, (d_1047_v53_).set(False, (self).f13), (self).f14, globalState))
            rhs178_ = default__.safeModuloInt(default__.safeDivisionInt((self).f13, (self).f13), ((self.f11).cardinality) + ((self).f13))
            lhs139_ = d_1044_v50_
            lhs140_ = default__.safeIndex(514, (d_1044_v50_).length(0))
            lhs141_ = globalState
            r0 = rhs175_
            lhs139_[lhs140_] = rhs176_
            lhs141_.f3 = rhs177_
            r0 = rhs178_
            d_1048_v54_: _dafny.Array
            nw188_ = _dafny.Array(int(0), 7)
            d_1048_v54_ = nw188_
            index162_ = default__.safeIndex(68, (d_1048_v54_).length(0))
            (d_1048_v54_)[index162_] = (self).f13
            index163_ = default__.safeIndex(68, (d_1048_v54_).length(0))
            (d_1048_v54_)[index163_] = (self).f13
        elif True:
            d_1049_v55_: str
            d_1049_v55_ = _dafny.CodePoint('y')
            d_1050_v56_: _dafny.MultiSet
            d_1050_v56_ = _dafny.MultiSet([d_1049_v55_, d_1049_v55_])
            d_1051_v57_: _dafny.Array
            nw189_ = _dafny.Array(_dafny.Seq({}), 8)
            d_1051_v57_ = nw189_
            d_1052_v58_: _dafny.Seq
            d_1052_v58_ = _dafny.SeqWithoutIsStrInference([(self).f13])
            index164_ = default__.safeIndex(204, (d_1051_v57_).length(0))
            (d_1051_v57_)[index164_] = d_1052_v58_
            index165_ = default__.safeIndex(204, (d_1051_v57_).length(0))
            rhs179_ = _dafny.MultiSet([d_1049_v55_])
            rhs180_ = (self).f14
            rhs181_ = default__.safeDivisionInt((0) - ((self).f13), (self).f13)
            rhs182_ = ((d_1052_v58_) + (d_1052_v58_)) + (d_1052_v58_)
            lhs142_ = globalState
            lhs143_ = d_1051_v57_
            lhs144_ = default__.safeIndex(204, (d_1051_v57_).length(0))
            d_1050_v56_ = rhs179_
            r1 = rhs180_
            lhs142_.f9 = rhs181_
            lhs143_[lhs144_] = rhs182_
            d_1053_v59_: _dafny.MultiSet
            d_1053_v59_ = _dafny.MultiSet([(self).f13])
            d_1054_v60_: D1
            d_1054_v60_ = D1_DC2(self.f12, self.f12, self.f12, d_1049_v55_, d_1049_v55_)
            d_1055_v61_: D1
            d_1055_v61_ = D1_DC2(p1, False, not(p1), d_1049_v55_, default__.fm14(d_1054_v60_, (self).f13, globalState))
            d_1056_v62_: _dafny.Array
            nw190_ = _dafny.Array(None, 23)
            nw190_[int(0)] = p0
            nw190_[int(1)] = _dafny.SeqWithoutIsStrInference([d_1049_v55_ for d_1057_i4_ in range(default__.abs(833))])
            nw190_[int(2)] = p0
            nw190_[int(3)] = p0
            nw190_[int(4)] = p0
            nw190_[int(5)] = p0
            nw190_[int(6)] = p0
            nw190_[int(7)] = p0
            nw190_[int(8)] = p0
            nw190_[int(9)] = p0
            nw190_[int(10)] = p0
            nw190_[int(11)] = p0
            nw190_[int(12)] = p0
            nw190_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ppym"))
            nw190_[int(14)] = (p0) + (p0)
            nw190_[int(15)] = p0
            nw190_[int(16)] = p0
            nw190_[int(17)] = (p0) + (p0)
            nw190_[int(18)] = (p0 if (self).f14 else (p0).set(default__.safeIndex((0) - (((d_1053_v59_).set((self).f13, default__.abs((0) - (((d_1053_v59_)[(self).f13] if ((self).f13) in (d_1053_v59_) else (self).f13))))).cardinality), len(p0)), (d_1055_v61_).cf6))
            nw190_[int(19)] = p0
            nw190_[int(20)] = _dafny.SeqWithoutIsStrInference([d_1049_v55_ for d_1058_i5_ in range(default__.abs(-542))])
            nw190_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "drfnfyqmg"))
            nw190_[int(22)] = p0
            d_1056_v62_ = nw190_
            d_1059_v63_: _dafny.Map
            d_1059_v63_ = _dafny.Map({(self).f14: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "whjni")))})
            d_1060_v64_: _dafny.Seq
            d_1060_v64_ = _dafny.SeqWithoutIsStrInference([True])
            d_1061_v65_: D2
            d_1061_v65_ = D2_DC5((self).f13, (d_1050_v56_).cardinality, (self).f13, ((d_1059_v63_)[(self).f14] if ((self).f14) in (d_1059_v63_) else (self).f13), d_1060_v64_)
            pat_let_tv21_ = d_1060_v64_
            d_1062_v66_: _dafny.Map
            def iife89_(_pat_let32_0):
                def iife90_(d_1063_dt__update__tmp_h1_):
                    def iife91_(_pat_let33_0):
                        def iife92_(d_1064_dt__update_hcf13_h1_):
                            def iife93_(_pat_let34_0):
                                def iife94_(d_1065_dt__update_hcf9_h0_):
                                    return D2_DC5(d_1065_dt__update_hcf9_h0_, (d_1063_dt__update__tmp_h1_).cf10, (d_1063_dt__update__tmp_h1_).cf11, (d_1063_dt__update__tmp_h1_).cf12, d_1064_dt__update_hcf13_h1_)
                                return iife94_(_pat_let34_0)
                            return iife93_((self).f13)
                        return iife92_(_pat_let33_0)
                    return iife91_(pat_let_tv21_)
                return iife90_(_pat_let32_0)
            d_1062_v66_ = _dafny.Map({(d_1061_v65_ if self.f12 else iife89_(d_1061_v65_)): d_1056_v62_})
            d_1066_v67_: _dafny.Array
            nw191_ = _dafny.Array(int(0), 26)
            d_1066_v67_ = nw191_
            d_1067_v68_: _dafny.Seq
            d_1067_v68_ = _dafny.SeqWithoutIsStrInference([d_1066_v67_, d_1066_v67_])
            d_1068_v70_: _dafny.Set
            d_1068_v70_ = _dafny.Set({(self).f13, (self).f13})
            d_1069_v71_: _dafny.Seq
            def iife95_():
                coll25_ = _dafny.Set()
                compr_25_: int
                for compr_25_ in (d_1053_v59_).Elements:
                    d_1070_v69_: int = compr_25_
                    if (d_1070_v69_) in (d_1053_v59_):
                        coll25_ = coll25_.union(_dafny.Set([(d_1070_v69_) * (-388)]))
                return _dafny.Set(coll25_)
            d_1069_v71_ = _dafny.SeqWithoutIsStrInference([iife95_()
            , d_1068_v70_, _dafny.Set({len(d_1060_v64_), (self).f13})])
            d_1071_v72_: _dafny.Seq
            d_1071_v72_ = _dafny.SeqWithoutIsStrInference([d_1066_v67_, d_1066_v67_, (d_1067_v68_)[default__.safeIndex(((d_1051_v57_)[default__.safeIndex(204, (d_1051_v57_).length(0))])[default__.safeIndex(default__.fm1(p1, (self).f13, len(d_1069_v71_), globalState), len((d_1051_v57_)[default__.safeIndex(204, (d_1051_v57_).length(0))]))], len(d_1067_v68_))]])
            d_1072_v73_: D9
            d_1072_v73_ = D9_DC24(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c")), self.f12, (self).f13)
            index166_ = default__.safeIndex(204, (d_1051_v57_).length(0))
            def iife96_(_pat_let35_0):
                def iife97_(d_1073_dt__update__tmp_h3_):
                    def iife98_(_pat_let36_0):
                        def iife99_(d_1074_dt__update_hcf9_h2_):
                            return D2_DC5(d_1074_dt__update_hcf9_h2_, (d_1073_dt__update__tmp_h3_).cf10, (d_1073_dt__update__tmp_h3_).cf11, (d_1073_dt__update__tmp_h3_).cf12, (d_1073_dt__update__tmp_h3_).cf13)
                        return iife99_(_pat_let36_0)
                    return iife98_(697)
                return iife97_(_pat_let35_0)
            def iife100_(_pat_let37_0):
                def iife101_(d_1075_dt__update__tmp_h2_):
                    def iife102_(_pat_let38_0):
                        def iife103_(d_1076_dt__update_hcf9_h1_):
                            return D2_DC5(d_1076_dt__update_hcf9_h1_, (d_1075_dt__update__tmp_h2_).cf10, (d_1075_dt__update__tmp_h2_).cf11, (d_1075_dt__update__tmp_h2_).cf12, (d_1075_dt__update__tmp_h2_).cf13)
                        return iife103_(_pat_let38_0)
                    return iife102_(697)
                return iife101_(_pat_let37_0)
            def iife104_(_pat_let39_0):
                def iife105_(d_1078_dt__update__tmp_h4_):
                    def iife106_(_pat_let40_0):
                        def iife107_(d_1079_dt__update_hcf58_h0_):
                            return D9_DC24((d_1078_dt__update__tmp_h4_).cf56, (d_1078_dt__update__tmp_h4_).cf57, d_1079_dt__update_hcf58_h0_)
                        return iife107_(_pat_let40_0)
                    return iife106_((self).f13)
                return iife105_(_pat_let39_0)
            rhs183_ = -676
            rhs184_ = ((d_1062_v66_)[iife96_(d_1061_v65_)] if (iife100_(d_1061_v65_)) in (d_1062_v66_) else p2)
            rhs185_ = (d_1052_v58_) + (_dafny.SeqWithoutIsStrInference([(self).f13 for d_1077_i6_ in range(default__.abs(-730))]))
            rhs186_ = (d_1067_v68_)[default__.safeIndex(default__.safeModuloInt((self).f13, (self).f13), len(d_1067_v68_))]
            rhs187_ = (iife104_(d_1072_v73_)).cf57
            lhs145_ = globalState
            lhs146_ = d_1051_v57_
            lhs147_ = default__.safeIndex(204, (d_1051_v57_).length(0))
            lhs148_ = globalState
            lhs149_ = globalState
            lhs145_.f9 = rhs183_
            d_1056_v62_ = rhs184_
            lhs146_[lhs147_] = rhs185_
            lhs148_.f4 = rhs186_
            lhs149_.f7 = rhs187_
            d_1080_v74_: _dafny.Map
            d_1080_v74_ = _dafny.Map({(self).f13: self.f12})
            d_1080_v74_ = (d_1080_v74_).set(default__.safeDivisionInt((self).f13, (self).f13), not (p1) or (self.f12))
            r0 = (0) - ((self).f13)
            d_1081_v75_: C0
            nw192_ = C0()
            nw192_.ctor__((self).f14, (self).f13, (self.f11).set(self.f12, default__.abs((self).f13)), (self).f14)
            d_1081_v75_ = nw192_
            d_1082_v76_: D3
            d_1082_v76_ = D3_DC9((self).f13, default__.safeModuloInt(((d_1053_v59_)[default__.fm1(self.f12, (self).f13, (self).f13, globalState)] if (default__.fm1(self.f12, (self).f13, (self).f13, globalState)) in (d_1053_v59_) else (self).f13), (0) - (len(((d_1051_v57_)[default__.safeIndex(204, (d_1051_v57_).length(0))]).set(default__.safeIndex(len((d_1051_v57_)[default__.safeIndex(204, (d_1051_v57_).length(0))]), len((d_1051_v57_)[default__.safeIndex(204, (d_1051_v57_).length(0))])), -968)))), d_1081_v75_, (d_1081_v75_).f18)
            d_1082_v76_ = d_1082_v76_
        d_1083_v77_: _dafny.Seq
        d_1083_v77_ = _dafny.SeqWithoutIsStrInference([not((self).f14), self.f12, p1])
        rhs188_ = (self).f14
        rhs189_ = (d_1083_v77_)[default__.safeIndex(len(_dafny.Map({(self).f13: (self).f13})), len(d_1083_v77_))]
        lhs150_ = globalState
        r1 = rhs188_
        lhs150_.f3 = rhs189_
        d_1084_v78_: _dafny.Map
        d_1084_v78_ = _dafny.Map({not(self.f12): (self).f13})
        d_1085_v79_: _dafny.MultiSet
        d_1085_v79_ = _dafny.MultiSet([d_1084_v78_])
        d_1086_v80_: D10
        d_1086_v80_ = D10_DC27(d_1084_v78_)
        d_1087_v81_: _dafny.Seq
        d_1087_v81_ = _dafny.SeqWithoutIsStrInference([len((d_1084_v78_).set(p1, (0) - ((self).f13))), (self).f13, (self).f13])
        d_1088_v82_: _dafny.Array
        nw193_ = _dafny.Array(None, 13)
        nw193_[int(0)] = _dafny.MultiSet([_dafny.Map({(self).f14: len(p0)}), d_1084_v78_, _dafny.Map({self.f12: (self).f13})])
        nw193_[int(1)] = _dafny.MultiSet([d_1084_v78_, d_1084_v78_])
        nw193_[int(2)] = (d_1085_v79_).intersection(d_1085_v79_)
        nw193_[int(3)] = d_1085_v79_
        nw193_[int(4)] = d_1085_v79_
        nw193_[int(5)] = _dafny.MultiSet([(d_1086_v80_).cf63])
        nw193_[int(6)] = d_1085_v79_
        nw193_[int(7)] = d_1085_v79_
        nw193_[int(8)] = d_1085_v79_
        nw193_[int(9)] = d_1085_v79_
        nw193_[int(10)] = d_1085_v79_
        nw193_[int(11)] = (d_1085_v79_).set(d_1084_v78_, default__.abs(len(d_1087_v81_)))
        nw193_[int(12)] = (d_1085_v79_) - (default__.fm19((self).f13, globalState))
        d_1088_v82_ = nw193_
        index167_ = default__.safeIndex(599, (d_1088_v82_).length(0))
        (d_1088_v82_)[index167_] = d_1085_v79_
        index168_ = default__.safeIndex(599, (d_1088_v82_).length(0))
        (d_1088_v82_)[index168_] = (d_1085_v79_) - ((d_1085_v79_) | (d_1085_v79_))
        d_1089_v83_: _dafny.Map
        d_1089_v83_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f13]): (self).f13})
        if (d_1087_v81_) in (d_1089_v83_):
            d_1090_v84_: _dafny.Map
            d_1090_v84_ = _dafny.Map({d_1083_v77_: p1})
            d_1090_v84_ = (d_1090_v84_).set((d_1083_v77_) + ((d_1083_v77_).set(default__.safeIndex((self.f11).cardinality, len(d_1083_v77_)), not(p1))), self.f12)
            d_1091_v85_: C2
            nw194_ = C2()
            nw194_.ctor__((self).f14, self.f11, (self).f14)
            d_1091_v85_ = nw194_
            d_1091_v85_ = d_1091_v85_
            d_1092_v86_: _dafny.Seq
            d_1092_v86_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kfcc"))
            d_1092_v86_ = (d_1092_v86_) + (p0)
            d_1093_v87_: D5
            d_1093_v87_ = D5_DC14()
            d_1094_v88_: D7
            d_1094_v88_ = D7_DC17(_dafny.Map({d_1093_v87_: p1}))
            d_1095_v89_: _dafny.Map
            d_1095_v89_ = _dafny.Map({d_1093_v87_: self.f12})
            pat_let_tv22_ = d_1095_v89_
            pat_let_tv23_ = d_1095_v89_
            def iife108_(_pat_let41_0):
                def iife109_(d_1096_dt__update__tmp_h5_):
                    def iife110_(_pat_let42_0):
                        def iife111_(d_1097_dt__update_hcf39_h0_):
                            return D7_DC17(d_1097_dt__update_hcf39_h0_)
                        return iife111_(_pat_let42_0)
                    return iife110_((pat_let_tv22_) | (pat_let_tv23_))
                return iife109_(_pat_let41_0)
            d_1094_v88_ = iife108_(d_1094_v88_)
            (globalState).f7 = (d_1091_v85_.f15 if d_1091_v85_.f15 else self.f12)
        elif True:
            (globalState).f8 = (self).f14
            d_1098_v90_: str
            d_1098_v90_ = _dafny.CodePoint('s')
            d_1099_v91_: D1
            d_1099_v91_ = D1_DC2(self.f12, False, (self).f14, d_1098_v90_, d_1098_v90_)
            d_1100_v92_: _dafny.Map
            d_1100_v92_ = _dafny.Map({len(d_1087_v81_): default__.fm14(d_1099_v91_, (self).f13, globalState)})
            d_1101_v93_: _dafny.Set
            d_1101_v93_ = _dafny.Set({(self).f14})
            d_1102_v94_: _dafny.MultiSet
            d_1102_v94_ = _dafny.MultiSet([(self).f13])
            d_1103_v95_: _dafny.Map
            d_1103_v95_ = _dafny.Map({(0) - ((self).f13): (self).f13})
            d_1104_v96_: _dafny.Set
            d_1104_v96_ = _dafny.Set({(self).f13, len(d_1103_v95_), (self).f13})
            d_1105_v97_: _dafny.Map
            d_1105_v97_ = _dafny.Map({(self).fm2(d_1102_v94_, (self).f13, d_1084_v78_, p1, globalState): d_1104_v96_})
            d_1106_v98_: _dafny.Array
            nw195_ = _dafny.Array(None, 24)
            nw195_[int(0)] = (0) - (len(d_1087_v81_))
            nw195_[int(1)] = default__.fm1((self).f14, (self).f13, 375, globalState)
            nw195_[int(2)] = len(d_1100_v92_)
            nw195_[int(3)] = (self).f13
            nw195_[int(4)] = (self).f13
            nw195_[int(5)] = (self).f13
            nw195_[int(6)] = (self).f13
            nw195_[int(7)] = len(d_1101_v93_)
            nw195_[int(8)] = (self).f13
            nw195_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mwn")))
            nw195_[int(10)] = len(p0)
            nw195_[int(11)] = (self).f13
            nw195_[int(12)] = (self).f13
            nw195_[int(13)] = (self).f13
            nw195_[int(14)] = (self).f13
            nw195_[int(15)] = len(d_1105_v97_)
            nw195_[int(16)] = (self).f13
            nw195_[int(17)] = (self).f13
            nw195_[int(18)] = default__.fm1(p1, len(d_1084_v78_), (self).f13, globalState)
            nw195_[int(19)] = (self).f13
            nw195_[int(20)] = (0) - ((self).f13)
            nw195_[int(21)] = (self).f13
            nw195_[int(22)] = (self).f13
            nw195_[int(23)] = (self).f13
            d_1106_v98_ = nw195_
            d_1107_v99_: _dafny.Set
            d_1107_v99_ = _dafny.Set({d_1106_v98_})
            d_1107_v99_ = ((d_1107_v99_) - (d_1107_v99_) if (self).f14 else (d_1107_v99_) | (d_1107_v99_))
            d_1108_v100_: _dafny.Map
            d_1108_v100_ = _dafny.Map({(self).f14: d_1106_v98_})
            d_1108_v100_ = (d_1108_v100_).set(self.f12, d_1106_v98_)
            (self).f12 = ((_dafny.SeqWithoutIsStrInference([(self).f13 for d_1109_i7_ in range(default__.abs(558))])) + (d_1087_v81_)) != (d_1087_v81_)
            d_1110_v101_: _dafny.Seq
            d_1110_v101_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            d_1110_v101_ = p0
        d_1111_v102_: C0
        nw196_ = C0()
        nw196_.ctor__(not((self).f14), (self).f13, self.f11, p1)
        d_1111_v102_ = nw196_
        (self).f12 = self.f12
        r0 = (0) - (default__.fm1(self.f12, (self).f13, default__.safeModuloInt((self).f13, 120), globalState))
        d_1112_v103_: D5
        d_1112_v103_ = D5_DC12(_dafny.SeqWithoutIsStrInference([382 for d_1113_i8_ in range(default__.abs(179))]))
        d_1114_v104_: _dafny.MultiSet
        d_1114_v104_ = _dafny.MultiSet([d_1112_v103_, d_1112_v103_, d_1112_v103_])
        r1 = (d_1114_v104_).issubset(d_1114_v104_)
        return r0, r1

    def m3(self, p0, globalState):
        d_1115_v0_: _dafny.Array
        nw197_ = _dafny.Array(int(0), 20)
        d_1115_v0_ = nw197_
        index169_ = default__.safeIndex(3, (d_1115_v0_).length(0))
        (d_1115_v0_)[index169_] = p0
        index170_ = default__.safeIndex(3, (d_1115_v0_).length(0))
        (d_1115_v0_)[index170_] = p0
        d_1116_v1_: _dafny.Array
        nw198_ = _dafny.Array(D4.default()(), 13)
        d_1116_v1_ = nw198_
        d_1117_v2_: D4
        d_1117_v2_ = D4_DC10(d_1115_v0_)
        pat_let_tv24_ = d_1115_v0_
        index171_ = default__.safeIndex(525, (d_1116_v1_).length(0))
        def iife112_(_pat_let43_0):
            def iife113_(d_1118_dt__update__tmp_h0_):
                def iife114_(_pat_let44_0):
                    def iife115_(d_1119_dt__update_hcf24_h0_):
                        return D4_DC10(d_1119_dt__update_hcf24_h0_)
                    return iife115_(_pat_let44_0)
                return iife114_(pat_let_tv24_)
            return iife113_(_pat_let43_0)
        (d_1116_v1_)[index171_] = iife112_(d_1117_v2_)
        index172_ = default__.safeIndex(525, (d_1116_v1_).length(0))
        (d_1116_v1_)[index172_] = D4_DC10(d_1115_v0_)
        if (self).f14:
            d_1120_v3_: str
            d_1120_v3_ = _dafny.CodePoint('m')
            d_1121_v4_: D1
            d_1121_v4_ = D1_DC2((self).f14, (self).f14, self.f12, d_1120_v3_, d_1120_v3_)
            d_1122_v5_: _dafny.Map
            d_1122_v5_ = _dafny.Map({p0: default__.fm14(d_1121_v4_, p0, globalState)})
            d_1123_v6_: _dafny.Set
            d_1123_v6_ = _dafny.Set({p0, (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})
            rhs190_ = not((d_1123_v6_).issubset(default__.fm20(d_1120_v3_, 678, 388, d_1122_v5_, globalState)))
            rhs191_ = -307
            lhs151_ = self
            lhs152_ = globalState
            lhs151_.f12 = rhs190_
            lhs152_.f9 = rhs191_
            index173_ = default__.safeIndex(3, (d_1115_v0_).length(0))
            (d_1115_v0_)[index173_] = (self).f13
            d_1124_v7_: _dafny.Map
            d_1124_v7_ = _dafny.Map({d_1115_v0_: len(d_1123_v6_)})
            d_1125_v8_: _dafny.Seq
            d_1125_v8_ = _dafny.SeqWithoutIsStrInference([d_1115_v0_, d_1115_v0_])
            d_1126_v9_: _dafny.Set
            d_1126_v9_ = _dafny.Set({self.f12})
            d_1127_v10_: _dafny.Seq
            d_1127_v10_ = _dafny.SeqWithoutIsStrInference([len(d_1126_v9_), p0, (self).f13, (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]])
            d_1124_v7_ = (d_1124_v7_).set((d_1125_v8_)[default__.safeIndex(len(d_1127_v10_), len(d_1125_v8_))], p0)
            d_1128_v11_: _dafny.Seq
            d_1128_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hoftygcb"))
            d_1128_v11_ = d_1128_v11_
            (globalState).f9 = (p0) * ((self).f13)
        elif True:
            (globalState).f9 = 701
            d_1129_v12_: _dafny.Array
            def lambda40_(d_1130_i0_):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ipdww"))

            init24_ = lambda40_
            nw199_ = _dafny.Array(None, 28)
            for i0_24_ in range(nw199_.length(0)):
                nw199_[i0_24_] = init24_(i0_24_)
            d_1129_v12_ = nw199_
            d_1131_v13_: D5
            d_1131_v13_ = D5_DC13((self).f13, self.f12, (p0) - (p0), d_1129_v12_, (self).f14)
            source10_ = d_1131_v13_
            if source10_.is_DC13:
                d_1132___mcc_h0_ = source10_.cf30
                d_1133___mcc_h1_ = source10_.cf31
                d_1134___mcc_h2_ = source10_.cf32
                d_1135___mcc_h3_ = source10_.cf33
                d_1136___mcc_h4_ = source10_.cf34
                d_1137_cf34_ = d_1136___mcc_h4_
                d_1138_cf33_ = d_1135___mcc_h3_
                d_1139_cf32_ = d_1134___mcc_h2_
                d_1140_cf31_ = d_1133___mcc_h1_
                d_1141_cf30_ = d_1132___mcc_h0_
                d_1142_v14_: C2
                nw200_ = C2()
                nw200_.ctor__(self.f12, (self.f11) - (_dafny.MultiSet([(self).f14])), d_1137_cf34_)
                d_1142_v14_ = nw200_
                d_1142_v14_ = d_1142_v14_
                d_1139_cf32_ = default__.safeModuloInt((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "klkgrol")))) * (d_1139_cf32_))
                d_1143_v15_: _dafny.Array
                def lambda41_(d_1144_i1_):
                    return self.f12

                init25_ = lambda41_
                nw201_ = _dafny.Array(None, 8)
                for i0_25_ in range(nw201_.length(0)):
                    nw201_[i0_25_] = init25_(i0_25_)
                d_1143_v15_ = nw201_
                index174_ = default__.safeIndex(803, (d_1143_v15_).length(0))
                (d_1143_v15_)[index174_] = self.f12
                index175_ = default__.safeIndex(803, (d_1143_v15_).length(0))
                (d_1143_v15_)[index175_] = (self).f14
                d_1145_v16_: _dafny.Map
                d_1145_v16_ = _dafny.Map({d_1142_v14_.f15: d_1143_v15_})
                (globalState).f3 = ((d_1145_v16_) | (d_1145_v16_)) == (d_1145_v16_)
            elif source10_.is_DC14:
                (globalState).f9 = default__.fm1((self).f14, 12, (len(_dafny.SeqWithoutIsStrInference([86, 617, (self.f11).cardinality]))) - ((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]), globalState)
                d_1146_v17_: _dafny.Seq
                d_1146_v17_ = _dafny.SeqWithoutIsStrInference([self.f12])
                d_1147_v18_: _dafny.Map
                d_1147_v18_ = _dafny.Map({_dafny.Map({False: False}): p0})
                d_1148_v19_: _dafny.Seq
                d_1148_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                d_1149_v20_: _dafny.Set
                d_1149_v20_ = _dafny.Set({len(d_1148_v19_)})
                d_1150_v21_: C0
                nw202_ = C0()
                nw202_.ctor__(self.f12, p0, (self.f11) - ((_dafny.MultiSet(d_1146_v17_)).set((self).f14, default__.abs(len(d_1146_v17_)))), (_dafny.Set({(0) - (((d_1147_v18_)[_dafny.Map({(False): (self).f14})] if (_dafny.Map({(False): (self).f14})) in (d_1147_v18_) else -628)), (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})).issubset(d_1149_v20_))
                d_1150_v21_ = nw202_
                nw203_ = C0()
                nw203_.ctor__((self).f14, (self).f13, _dafny.MultiSet([(d_1150_v21_).f17]), (self).f14)
                d_1150_v21_ = nw203_
                d_1151_v22_: D9
                d_1151_v22_ = D9_DC25(d_1148_v19_, (d_1150_v21_).f17, p0)
                d_1152_v23_: str
                d_1152_v23_ = _dafny.CodePoint('y')
                d_1148_v19_ = ((d_1151_v22_).cf59).set(default__.safeIndex((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], len((d_1151_v22_).cf59)), d_1152_v23_)
                d_1153_v24_: D2
                d_1153_v24_ = D2_DC5(p0, (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], 223, p0, d_1146_v17_)
                rhs192_ = self.f12
                rhs193_ = d_1152_v23_
                rhs194_ = (default__.safeModuloInt((d_1150_v21_).f18, (d_1150_v21_).fm12(globalState))) < ((len((d_1153_v24_).cf13)) + ((self).f13))
                rhs195_ = default__.safeModuloInt(len(d_1149_v20_), len(_dafny.Map({d_1152_v23_: self.f12})))
                lhs153_ = globalState
                lhs154_ = globalState
                lhs155_ = globalState
                lhs153_.f8 = rhs192_
                d_1152_v23_ = rhs193_
                lhs154_.f7 = rhs194_
                lhs155_.f9 = rhs195_
            elif True:
                d_1154___mcc_h5_ = source10_.cf29
                d_1155_cf29_ = d_1154___mcc_h5_
                d_1156_v25_: T0
                nw204_ = C2()
                nw204_.ctor__((self).f14, self.f11, self.f12)
                d_1156_v25_ = nw204_
                d_1156_v25_ = d_1156_v25_
                d_1157_v26_: str
                d_1157_v26_ = _dafny.CodePoint('c')
                d_1158_v27_: _dafny.Seq
                d_1158_v27_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_1159_v28_: _dafny.Map
                d_1159_v28_ = _dafny.Map({len(d_1158_v27_): (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})
                d_1160_v29_: C0
                nw205_ = C0()
                nw205_.ctor__(default__.fm0(d_1157_v26_, globalState), default__.fm1(d_1156_v25_.f12, (self).f13, (self).f13, globalState), _dafny.MultiSet([not((self).f14), (self).f14, self.f12, self.f12, not((self).f14)]), not((((d_1159_v28_)[(self).f13] if ((self).f13) in (d_1159_v28_) else (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))])) == (-424)))
                d_1160_v29_ = nw205_
                d_1161_v30_: _dafny.Set
                d_1161_v30_ = _dafny.Set({self.f12, (self).f14})
                d_1162_v31_: C0
                nw206_ = C0()
                nw206_.ctor__((_dafny.Set({self.f12, (self).f14})).ispropersubset(d_1161_v30_), 502, _dafny.MultiSet([(self).f14, (d_1160_v29_).f17]), not((d_1160_v29_).f17))
                d_1162_v31_ = nw206_
                d_1163_v32_: _dafny.Array
                def lambda42_(d_1164_v31_, d_1165_v25_):
                    def lambda43_(d_1166_i2_):
                        return (_dafny.Map({(d_1164_v31_).f17: (d_1164_v31_).f17})) not in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_1165_v25_.f12: (d_1164_v31_).f17})]))

                    return lambda43_

                init26_ = lambda42_(d_1162_v31_, d_1156_v25_)
                nw207_ = _dafny.Array(None, 6)
                for i0_26_ in range(nw207_.length(0)):
                    nw207_[i0_26_] = init26_(i0_26_)
                d_1163_v32_ = nw207_
                index176_ = default__.safeIndex(186, (d_1163_v32_).length(0))
                (d_1163_v32_)[index176_] = (d_1162_v31_).f17
                index177_ = default__.safeIndex(186, (d_1163_v32_).length(0))
                (d_1163_v32_)[index177_] = not (default__.fm0(d_1157_v26_, globalState)) or (not((d_1160_v29_).f17))
            index178_ = default__.safeIndex(3, (d_1115_v0_).length(0))
            (d_1115_v0_)[index178_] = (default__.safeModuloInt((0) - ((self).f13), (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))])) + ((self).f13)
            d_1167_v33_: _dafny.Seq
            d_1167_v33_ = _dafny.SeqWithoutIsStrInference([d_1115_v0_])
            d_1168_v34_: str
            d_1168_v34_ = _dafny.CodePoint('l')
            d_1169_v35_: D1
            d_1169_v35_ = D1_DC2((self).f14, not((self).f14), self.f12, _dafny.CodePoint('c'), d_1168_v34_)
            source11_ = (D5_DC13((self.f11).cardinality, self.f12, len(d_1167_v33_), d_1129_v12_, (d_1169_v35_).cf2) if self.f12 else d_1131_v13_)
            if source11_.is_DC13:
                d_1170___mcc_h6_ = source11_.cf30
                d_1171___mcc_h7_ = source11_.cf31
                d_1172___mcc_h8_ = source11_.cf32
                d_1173___mcc_h9_ = source11_.cf33
                d_1174___mcc_h10_ = source11_.cf34
                d_1175_cf34_ = d_1174___mcc_h10_
                d_1176_cf33_ = d_1173___mcc_h9_
                d_1177_cf32_ = d_1172___mcc_h8_
                d_1178_cf31_ = d_1171___mcc_h7_
                d_1179_cf30_ = d_1170___mcc_h6_
                (globalState).f3 = d_1175_cf34_
                d_1180_v36_: _dafny.Array
                def lambda44_(d_1181_cf34_):
                    def lambda45_(d_1182_i3_):
                        return d_1181_cf34_

                    return lambda45_

                init27_ = lambda44_(d_1175_cf34_)
                nw208_ = _dafny.Array(None, 19)
                for i0_27_ in range(nw208_.length(0)):
                    nw208_[i0_27_] = init27_(i0_27_)
                d_1180_v36_ = nw208_
                index179_ = default__.safeIndex(763, (d_1180_v36_).length(0))
                (d_1180_v36_)[index179_] = not(d_1175_cf34_)
                d_1183_v38_: _dafny.Set
                d_1183_v38_ = _dafny.Set({_dafny.CodePoint('c'), d_1168_v34_, d_1168_v34_})
                index180_ = default__.safeIndex(763, (d_1180_v36_).length(0))
                def iife116_():
                    coll26_ = _dafny.Set()
                    compr_26_: str
                    for compr_26_ in (default__.fm21(d_1175_cf34_, True, d_1178_cf31_, globalState)).keys.Elements:
                        d_1184_v37_: str = compr_26_
                        if (d_1184_v37_) in (default__.fm21(d_1175_cf34_, True, d_1178_cf31_, globalState)):
                            coll26_ = coll26_.union(_dafny.Set([d_1184_v37_]))
                    return _dafny.Set(coll26_)
                (d_1180_v36_)[index180_] = (iife116_()
                ).issubset(d_1183_v38_)
                d_1185_v39_: _dafny.Seq
                d_1185_v39_ = _dafny.SeqWithoutIsStrInference([d_1180_v36_, d_1180_v36_, d_1180_v36_, d_1180_v36_, d_1180_v36_])
                d_1185_v39_ = d_1185_v39_
                d_1186_v40_: _dafny.Map
                d_1186_v40_ = _dafny.Map({d_1115_v0_: (self).f13})
                d_1168_v34_ = (d_1168_v34_ if (d_1186_v40_) == (d_1186_v40_) else d_1168_v34_)
            elif source11_.is_DC14:
                d_1187_v41_: _dafny.Seq
                d_1187_v41_ = _dafny.SeqWithoutIsStrInference([self.f12])
                index181_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                index182_ = default__.safeIndex(525, (d_1116_v1_).length(0))
                rhs196_ = (self).f13
                rhs197_ = default__.safeDivisionInt(len(d_1187_v41_), default__.fm1(self.f12, p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scarno"))), globalState))
                rhs198_ = D4_DC10(d_1115_v0_)
                lhs156_ = d_1115_v0_
                lhs157_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                lhs158_ = globalState
                lhs159_ = d_1116_v1_
                lhs160_ = default__.safeIndex(525, (d_1116_v1_).length(0))
                lhs156_[lhs157_] = rhs196_
                lhs158_.f9 = rhs197_
                lhs159_[lhs160_] = rhs198_
                d_1188_v42_: _dafny.Seq
                d_1188_v42_ = _dafny.SeqWithoutIsStrInference([d_1187_v41_, d_1187_v41_, (d_1187_v41_).set(default__.safeIndex(815, len(d_1187_v41_)), self.f12)])
                d_1189_v43_: _dafny.Map
                d_1189_v43_ = _dafny.Map({False: d_1188_v42_})
                d_1189_v43_ = (d_1189_v43_).set((self).f14, (d_1188_v42_) + (d_1188_v42_))
                d_1190_v44_: _dafny.Array
                nw209_ = _dafny.Array(False, 28)
                d_1190_v44_ = nw209_
                index183_ = default__.safeIndex(468, (d_1190_v44_).length(0))
                (d_1190_v44_)[index183_] = (self).f14
                index184_ = default__.safeIndex(468, (d_1190_v44_).length(0))
                rhs199_ = p0
                rhs200_ = self.f12
                rhs201_ = self.f12
                lhs161_ = globalState
                lhs162_ = d_1190_v44_
                lhs163_ = default__.safeIndex(468, (d_1190_v44_).length(0))
                lhs164_ = globalState
                lhs161_.f9 = rhs199_
                lhs162_[lhs163_] = rhs200_
                lhs164_.f8 = rhs201_
                d_1191_v45_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1191_v45_ = nw210_
                d_1191_v45_ = d_1191_v45_
            elif True:
                d_1192___mcc_h11_ = source11_.cf29
                d_1193_cf29_ = d_1192___mcc_h11_
                d_1194_v46_: _dafny.Seq
                d_1194_v46_ = _dafny.SeqWithoutIsStrInference([self.f12])
                d_1195_v47_: T0
                nw211_ = C0()
                nw211_.ctor__(True, (self).f13, _dafny.MultiSet(d_1194_v46_), False)
                d_1195_v47_ = nw211_
                d_1196_v48_: _dafny.Seq
                d_1196_v48_ = _dafny.SeqWithoutIsStrInference([d_1195_v47_])
                d_1197_v49_: _dafny.Map
                d_1197_v49_ = _dafny.Map({p0: (d_1196_v48_) + (d_1196_v48_)})
                d_1197_v49_ = (d_1197_v49_).set(p0, (d_1196_v48_) + (d_1196_v48_))
                d_1198_v50_: C2
                nw212_ = C2()
                nw212_.ctor__(d_1195_v47_.f12, self.f11, (self).f14)
                d_1198_v50_ = nw212_
                d_1199_v51_: _dafny.Array
                nw213_ = _dafny.Array(_dafny.Set({}), 5)
                d_1199_v51_ = nw213_
                d_1200_v52_: _dafny.Seq
                d_1200_v52_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tdx"))
                d_1201_v53_: _dafny.Set
                d_1201_v53_ = _dafny.Set({len(d_1200_v52_), len(_dafny.SeqWithoutIsStrInference([d_1168_v34_ for d_1202_i4_ in range(default__.abs(583))]))})
                index185_ = default__.safeIndex(668, (d_1199_v51_).length(0))
                (d_1199_v51_)[index185_] = d_1201_v53_
                index186_ = default__.safeIndex(668, (d_1199_v51_).length(0))
                def iife117_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(188, 25):
                        d_1203_v54_: int = compr_27_
                        if ((188) <= (d_1203_v54_)) and ((d_1203_v54_) < (25)):
                            coll27_ = coll27_.union(_dafny.Set([(d_1203_v54_) - ((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))])]))
                    return _dafny.Set(coll27_)
                (d_1199_v51_)[index186_] = iife117_()
                
                (globalState).f8 = (d_1198_v50_.f15 if default__.fm0(d_1168_v34_, globalState) else False)
            (globalState).f7 = (self).fm6(not(self.f12), (self).f13, globalState)
        d_1204_i5_: int
        d_1204_i5_ = 0
        with _dafny.label("8"):
            while self.f12:
                with _dafny.c_label("8"):
                    if (d_1204_i5_) >= (100):
                        raise _dafny.Break("8")
                    d_1204_i5_ = (d_1204_i5_) + (1)
                    default__.m0(globalState)
                    index187_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                    (d_1115_v0_)[index187_] = (-310) * (default__.safeModuloInt((0) - ((self).f13), p0))
                    d_1205_v55_: str
                    d_1205_v55_ = _dafny.CodePoint('p')
                    d_1206_v56_: _dafny.Map
                    d_1206_v56_ = _dafny.Map({(d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]: d_1205_v55_})
                    if default__.fm0(((d_1206_v56_)[p0] if (p0) in (d_1206_v56_) else d_1205_v55_), globalState):
                        d_1207_v57_: _dafny.Array
                        def lambda46_(d_1208_i6_):
                            return ((self).f13) < ((0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ad"))))))

                        init28_ = lambda46_
                        nw214_ = _dafny.Array(None, 25)
                        for i0_28_ in range(nw214_.length(0)):
                            nw214_[i0_28_] = init28_(i0_28_)
                        d_1207_v57_ = nw214_
                        d_1209_v58_: _dafny.Seq
                        d_1209_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmboh"))
                        d_1210_v59_: _dafny.Set
                        d_1210_v59_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1205_v55_ for d_1211_i7_ in range(default__.abs(134))]))})
                        d_1212_v60_: D7
                        d_1212_v60_ = D7_DC18(p0, d_1209_v58_, (self).f14, (self).f14, d_1210_v59_)
                        d_1213_v61_: _dafny.Seq
                        d_1213_v61_ = _dafny.SeqWithoutIsStrInference([self.f12])
                        pat_let_tv25_ = d_1210_v59_
                        nw215_ = _dafny.Array(None, 11)
                        nw215_[int(0)] = not((self).f14)
                        nw215_[int(1)] = (self).f14
                        nw215_[int(2)] = (self).f14
                        nw215_[int(3)] = False
                        nw215_[int(4)] = self.f12
                        nw215_[int(5)] = (d_1205_v55_) not in (d_1209_v58_)
                        nw215_[int(6)] = (self).f14
                        def iife118_(_pat_let45_0):
                            def iife119_(d_1214_dt__update__tmp_h1_):
                                def iife120_(_pat_let46_0):
                                    def iife121_(d_1215_dt__update_hcf44_h0_):
                                        return D7_DC18((d_1214_dt__update__tmp_h1_).cf40, (d_1214_dt__update__tmp_h1_).cf41, (d_1214_dt__update__tmp_h1_).cf42, (d_1214_dt__update__tmp_h1_).cf43, d_1215_dt__update_hcf44_h0_)
                                    return iife121_(_pat_let46_0)
                                return iife120_(pat_let_tv25_)
                            return iife119_(_pat_let45_0)
                        nw215_[int(7)] = not((iife118_(d_1212_v60_)).cf43)
                        nw215_[int(8)] = (self).f14
                        nw215_[int(9)] = self.f12
                        nw215_[int(10)] = (d_1213_v61_) != (d_1213_v61_)
                        d_1207_v57_ = nw215_
                        d_1216_v62_: C2
                        nw216_ = C2()
                        nw216_.ctor__(self.f12, (_dafny.MultiSet([(self).f14])).set(default__.fm0(_dafny.CodePoint('c'), globalState), default__.abs((self).f13)), (self).f14)
                        d_1216_v62_ = nw216_
                        d_1217_v63_: _dafny.Map
                        d_1217_v63_ = _dafny.Map({p0: d_1207_v57_})
                        (globalState).f8 = not ((d_1217_v63_) == (d_1217_v63_)) or (d_1216_v62_.f15)
                        d_1218_v64_: C2
                        nw217_ = C2()
                        nw217_.ctor__(self.f12, self.f11, (self).f14)
                        d_1218_v64_ = nw217_
                        (globalState).f9 = (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]
                    elif True:
                        d_1219_v65_: _dafny.Map
                        d_1219_v65_ = _dafny.Map({self.f12: (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})
                        d_1220_v66_: _dafny.Seq
                        d_1220_v66_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tpnnpacy")))])
                        d_1221_v67_: _dafny.Seq
                        d_1221_v67_ = _dafny.SeqWithoutIsStrInference([d_1220_v66_, d_1220_v66_])
                        d_1222_v68_: _dafny.Map
                        d_1222_v68_ = _dafny.Map({d_1219_v65_: ((d_1221_v67_)[default__.safeIndex((d_1220_v66_)[default__.safeIndex(p0, len(d_1220_v66_))], len(d_1221_v67_))]) + (d_1220_v66_)})
                        d_1222_v68_ = d_1222_v68_
                        d_1223_v69_: _dafny.Array
                        nw218_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_1223_v69_ = nw218_
                        index188_ = default__.safeIndex(251, (d_1223_v69_).length(0))
                        (d_1223_v69_)[index188_] = d_1115_v0_
                        index189_ = default__.safeIndex(251, (d_1223_v69_).length(0))
                        (d_1223_v69_)[index189_] = d_1115_v0_
                        d_1224_v70_: C2
                        nw219_ = C2()
                        nw219_.ctor__(default__.fm0(d_1205_v55_, globalState), (default__.fm22(p0, self.f12, globalState)).set((self).f14, default__.abs((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))])), (self).f14)
                        d_1224_v70_ = nw219_
                        d_1225_v71_: _dafny.Map
                        d_1225_v71_ = _dafny.Map({(self.f11).set(d_1224_v70_.f15, default__.abs(p0)): (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})
                        d_1226_v72_: _dafny.Map
                        d_1226_v72_ = _dafny.Map({d_1205_v55_: (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})
                        d_1227_v74_: _dafny.Array
                        nw220_ = _dafny.Array(None, 8)
                        nw220_[int(0)] = (d_1226_v72_).set(d_1205_v55_, (self).f13)
                        nw220_[int(1)] = d_1226_v72_
                        nw220_[int(2)] = d_1226_v72_
                        nw220_[int(3)] = d_1226_v72_
                        def iife122_():
                            coll28_ = _dafny.Map()
                            compr_28_: str
                            for compr_28_ in (_dafny.Set({d_1205_v55_})).Elements:
                                d_1228_v73_: str = compr_28_
                                if (d_1228_v73_) in (_dafny.Set({d_1205_v55_})):
                                    coll28_[d_1228_v73_] = (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]
                            return _dafny.Map(coll28_)
                        nw220_[int(4)] = iife122_()
                        
                        nw220_[int(5)] = d_1226_v72_
                        nw220_[int(6)] = d_1226_v72_
                        nw220_[int(7)] = d_1226_v72_
                        d_1227_v74_ = nw220_
                        index190_ = default__.safeIndex(354, (d_1227_v74_).length(0))
                        (d_1227_v74_)[index190_] = _dafny.Map({d_1205_v55_: (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})
                        d_1229_v76_: _dafny.MultiSet
                        d_1229_v76_ = _dafny.MultiSet([self.f11])
                        d_1230_v78_: _dafny.Seq
                        d_1230_v78_ = _dafny.SeqWithoutIsStrInference([self.f11, self.f11])
                        d_1231_v79_: _dafny.Set
                        d_1231_v79_ = _dafny.Set({self.f11, (d_1230_v78_)[default__.safeIndex(p0, len(d_1230_v78_))]})
                        index191_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                        index192_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                        index193_ = default__.safeIndex(354, (d_1227_v74_).length(0))
                        def iife123_():
                            coll29_ = _dafny.Map()
                            compr_29_: _dafny.MultiSet
                            for compr_29_ in (d_1229_v76_).Elements:
                                d_1232_v75_: _dafny.MultiSet = compr_29_
                                if (d_1232_v75_) in (d_1229_v76_):
                                    coll29_[d_1232_v75_] = len(d_1219_v65_)
                            return _dafny.Map(coll29_)
                        def iife124_():
                            coll30_ = _dafny.Map()
                            compr_30_: _dafny.MultiSet
                            for compr_30_ in (d_1231_v79_).Elements:
                                d_1233_v77_: _dafny.MultiSet = compr_30_
                                if (d_1233_v77_) in (d_1231_v79_):
                                    coll30_[d_1233_v77_] = (self).f13
                            return _dafny.Map(coll30_)
                        rhs202_ = d_1224_v70_
                        rhs203_ = (default__.safeDivisionInt(default__.fm1((self).f14, (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], (self).f13, globalState), (0) - ((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]))) * (p0)
                        rhs204_ = ((iife123_()
                        ) | (iife124_()
                        ) if ((self).f13) != (411) else (_dafny.Map({self.f11: (self).f13})) | (d_1225_v71_))
                        rhs205_ = default__.safeDivisionInt(815, len(d_1219_v65_))
                        rhs206_ = d_1226_v72_
                        lhs165_ = d_1115_v0_
                        lhs166_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                        lhs167_ = d_1115_v0_
                        lhs168_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                        lhs169_ = d_1227_v74_
                        lhs170_ = default__.safeIndex(354, (d_1227_v74_).length(0))
                        d_1224_v70_ = rhs202_
                        lhs165_[lhs166_] = rhs203_
                        d_1225_v71_ = rhs204_
                        lhs167_[lhs168_] = rhs205_
                        lhs169_[lhs170_] = rhs206_
                        d_1234_v80_: T0
                        nw221_ = C0()
                        nw221_.ctor__(self.f12, (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], self.f11, self.f12)
                        d_1234_v80_ = nw221_
                        d_1235_v81_: _dafny.Seq
                        d_1235_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xcisubu"))
                        d_1236_v82_: _dafny.Map
                        d_1236_v82_ = _dafny.Map({d_1234_v80_: d_1235_v81_})
                        d_1236_v82_ = (d_1236_v82_).set(d_1234_v80_, d_1235_v81_)
                        (globalState).f3 = default__.fm0(d_1205_v55_, globalState)
                    d_1237_v83_: _dafny.Set
                    d_1237_v83_ = _dafny.Set({p0})
                    d_1238_v84_: _dafny.Array
                    nw222_ = _dafny.Array(None, 2)
                    nw222_[int(0)] = not (self.f12) or ((self).f14)
                    nw222_[int(1)] = (d_1237_v83_).issubset(d_1237_v83_)
                    d_1238_v84_ = nw222_
                    index194_ = default__.safeIndex(705, (d_1238_v84_).length(0))
                    (d_1238_v84_)[index194_] = (default__.fm0(d_1205_v55_, globalState) if self.f12 else (self).f14)
                    d_1239_v85_: _dafny.Seq
                    d_1239_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbi"))
                    index195_ = default__.safeIndex(705, (d_1238_v84_).length(0))
                    (d_1238_v84_)[index195_] = ((default__.fm9(p0, False, globalState) if (self).fm6((self).f14, p0, globalState) else d_1239_v85_)) == (d_1239_v85_)
                    pass
            pass
        d_1240_v86_: _dafny.Set
        d_1240_v86_ = _dafny.Set({self.f12, False, (self).f14})
        d_1241_v87_: _dafny.Seq
        d_1241_v87_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjje"))
        d_1242_v88_: _dafny.Set
        d_1242_v88_ = _dafny.Set({(self).f13})
        hi2_ = (D7_DC18(len(d_1240_v86_), d_1241_v87_, False, (self).f14, d_1242_v88_)).cf40
        for d_1243_i8_ in range((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], hi2_):
            if (self).f14:
                d_1244_v89_: _dafny.Array
                def lambda47_(d_1245_i9_):
                    return (self).f14

                init29_ = lambda47_
                nw223_ = _dafny.Array(None, 21)
                for i0_29_ in range(nw223_.length(0)):
                    nw223_[i0_29_] = init29_(i0_29_)
                d_1244_v89_ = nw223_
                d_1246_v90_: _dafny.Map
                d_1246_v90_ = _dafny.Map({d_1243_i8_: d_1244_v89_})
                d_1246_v90_ = ((d_1246_v90_) | (d_1246_v90_)) | (d_1246_v90_)
                d_1247_v91_: _dafny.Seq
                d_1247_v91_ = _dafny.SeqWithoutIsStrInference([d_1243_i8_])
                (globalState).f9 = ((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))] if self.f12 else ((d_1247_v91_)[default__.safeIndex(p0, len(d_1247_v91_))]) - ((self).f13))
                d_1246_v90_ = d_1246_v90_
                index196_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                (d_1115_v0_)[index196_] = default__.safeModuloInt((self).f13, (self).f13)
                d_1248_v92_: _dafny.Map
                d_1248_v92_ = _dafny.Map({(self).f14: p0})
                d_1249_v93_: _dafny.Seq
                d_1249_v93_ = _dafny.SeqWithoutIsStrInference([d_1248_v92_])
                d_1250_v94_: _dafny.Map
                d_1250_v94_ = _dafny.Map({(self).f14: d_1241_v87_})
                d_1251_v95_: _dafny.Map
                d_1251_v95_ = _dafny.Map({p0: d_1248_v92_})
                d_1252_v96_: _dafny.Map
                d_1252_v96_ = _dafny.Map({self.f12: self.f12})
                d_1253_v97_: _dafny.Array
                nw224_ = _dafny.Array(None, 19)
                nw224_[int(0)] = d_1248_v92_
                nw224_[int(1)] = d_1248_v92_
                nw224_[int(2)] = ((d_1249_v93_)[default__.safeIndex(len(d_1250_v94_), len(d_1249_v93_))]) | (d_1248_v92_)
                nw224_[int(3)] = d_1248_v92_
                nw224_[int(4)] = (_dafny.Map({self.f12: (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]})) | (d_1248_v92_)
                nw224_[int(5)] = d_1248_v92_
                nw224_[int(6)] = (d_1248_v92_) | (d_1248_v92_)
                nw224_[int(7)] = (d_1248_v92_).set(False, (self).f13)
                nw224_[int(8)] = (default__.fm23(p0, globalState)) | (d_1248_v92_)
                nw224_[int(9)] = d_1248_v92_
                nw224_[int(10)] = d_1248_v92_
                nw224_[int(11)] = (default__.fm23((self).f13, globalState)).set((self).f14, (self).f13)
                nw224_[int(12)] = (d_1248_v92_) | (d_1248_v92_)
                nw224_[int(13)] = d_1248_v92_
                nw224_[int(14)] = (d_1248_v92_) | (d_1248_v92_)
                nw224_[int(15)] = (d_1248_v92_) | (_dafny.Map({(self).f14: p0}))
                nw224_[int(16)] = ((d_1251_v95_)[len(d_1252_v96_)] if (len(d_1252_v96_)) in (d_1251_v95_) else _dafny.Map({False: d_1243_i8_}))
                nw224_[int(17)] = (d_1248_v92_) | (d_1248_v92_)
                nw224_[int(18)] = d_1248_v92_
                d_1253_v97_ = nw224_
                rhs207_ = d_1253_v97_
                rhs208_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l"))
                d_1253_v97_ = rhs207_
                d_1241_v87_ = rhs208_
            elif True:
                d_1254_v98_: _dafny.Array
                def lambda48_(d_1255_i10_):
                    return True

                init30_ = lambda48_
                nw225_ = _dafny.Array(None, 2)
                for i0_30_ in range(nw225_.length(0)):
                    nw225_[i0_30_] = init30_(i0_30_)
                d_1254_v98_ = nw225_
                index197_ = default__.safeIndex(627, (d_1254_v98_).length(0))
                (d_1254_v98_)[index197_] = self.f12
                index198_ = default__.safeIndex(627, (d_1254_v98_).length(0))
                (d_1254_v98_)[index198_] = (self).f14
                (globalState).f1 = d_1115_v0_
                d_1256_v99_: _dafny.Seq
                d_1256_v99_ = _dafny.SeqWithoutIsStrInference([d_1240_v86_, (d_1240_v86_) - (d_1240_v86_)])
                d_1257_v100_: str
                d_1257_v100_ = _dafny.CodePoint('q')
                d_1258_v101_: _dafny.Seq
                d_1258_v101_ = _dafny.SeqWithoutIsStrInference([d_1241_v87_])
                index199_ = default__.safeIndex(627, (d_1254_v98_).length(0))
                rhs209_ = len((d_1256_v99_)[default__.safeIndex((0) - (default__.fm1(True, d_1243_i8_, (0) - (d_1243_i8_), globalState)), len(d_1256_v99_))])
                rhs210_ = (self.f12 if (self).f14 else (d_1254_v98_)[default__.safeIndex(627, (d_1254_v98_).length(0))])
                rhs211_ = (d_1241_v87_).set(default__.safeIndex(((self.f11) | (self.f11)).cardinality, len(d_1241_v87_)), d_1257_v100_)
                rhs212_ = (_dafny.CodePoint('d')) in ((d_1258_v101_)[default__.safeIndex((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], len(d_1258_v101_))])
                rhs213_ = not(self.f12)
                lhs171_ = globalState
                lhs172_ = self
                lhs173_ = globalState
                lhs174_ = d_1254_v98_
                lhs175_ = default__.safeIndex(627, (d_1254_v98_).length(0))
                lhs171_.f9 = rhs209_
                lhs172_.f12 = rhs210_
                d_1241_v87_ = rhs211_
                lhs173_.f8 = rhs212_
                lhs174_[lhs175_] = rhs213_
                default__.m0(globalState)
                d_1259_v102_: _dafny.Array
                nw226_ = _dafny.Array(None, 10)
                d_1259_v102_ = nw226_
                d_1260_v103_: _dafny.Map
                d_1260_v103_ = _dafny.Map({(d_1254_v98_)[default__.safeIndex(627, (d_1254_v98_).length(0))]: d_1259_v102_})
                d_1261_v104_: D9
                d_1261_v104_ = D9_DC23(d_1260_v103_)
                pat_let_tv26_ = d_1260_v103_
                index200_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                index201_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                def iife125_(_pat_let47_0):
                    def iife126_(d_1262_dt__update__tmp_h2_):
                        def iife127_(_pat_let48_0):
                            def iife128_(d_1263_dt__update_hcf55_h0_):
                                return D9_DC23(d_1263_dt__update_hcf55_h0_)
                            return iife128_(_pat_let48_0)
                        return iife127_(pat_let_tv26_)
                    return iife126_(_pat_let47_0)
                rhs214_ = iife125_(d_1261_v104_)
                rhs215_ = (self).f13
                rhs216_ = p0
                rhs217_ = -637
                lhs176_ = d_1115_v0_
                lhs177_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                lhs178_ = globalState
                lhs179_ = d_1115_v0_
                lhs180_ = default__.safeIndex(3, (d_1115_v0_).length(0))
                d_1261_v104_ = rhs214_
                lhs176_[lhs177_] = rhs215_
                lhs178_.f9 = rhs216_
                lhs179_[lhs180_] = rhs217_
            index202_ = default__.safeIndex(3, (d_1115_v0_).length(0))
            (d_1115_v0_)[index202_] = (self).f13
            d_1264_v105_: _dafny.Seq
            d_1264_v105_ = _dafny.SeqWithoutIsStrInference([self.f12, self.f12])
            d_1265_v106_: _dafny.Map
            d_1265_v106_ = _dafny.Map({False: (self).f14})
            d_1266_v107_: _dafny.Array
            nw227_ = _dafny.Array(None, 8)
            nw227_[int(0)] = (self).f14
            nw227_[int(1)] = not((True if (self).f14 else self.f12))
            nw227_[int(2)] = not(((0) - (len((d_1264_v105_).set(default__.safeIndex(p0, len(d_1264_v105_)), self.f12)))) == ((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]))
            nw227_[int(3)] = ((d_1265_v106_)[(self).f14] if ((self).f14) in (d_1265_v106_) else self.f12)
            nw227_[int(4)] = self.f12
            nw227_[int(5)] = (self).f14
            nw227_[int(6)] = self.f12
            nw227_[int(7)] = self.f12
            d_1266_v107_ = nw227_
            index203_ = default__.safeIndex(174, (d_1266_v107_).length(0))
            (d_1266_v107_)[index203_] = (d_1241_v87_) == (d_1241_v87_)
            d_1267_v108_: _dafny.Array
            nw228_ = _dafny.Array(D8.default()(), 9)
            d_1267_v108_ = nw228_
            d_1268_v109_: D8
            d_1268_v109_ = D8_DC19(self.f11)
            d_1269_v110_: D8
            d_1269_v110_ = D8_DC22(d_1268_v109_)
            index204_ = default__.safeIndex(659, (d_1267_v108_).length(0))
            (d_1267_v108_)[index204_] = d_1269_v110_
            index205_ = default__.safeIndex(174, (d_1266_v107_).length(0))
            index206_ = default__.safeIndex(3, (d_1115_v0_).length(0))
            index207_ = default__.safeIndex(659, (d_1267_v108_).length(0))
            rhs218_ = not((d_1242_v88_).issubset(d_1242_v88_))
            rhs219_ = (((self).f13) + (466) if self.f12 else d_1243_i8_)
            rhs220_ = d_1269_v110_
            rhs221_ = (self).f13
            rhs222_ = p0
            lhs181_ = d_1266_v107_
            lhs182_ = default__.safeIndex(174, (d_1266_v107_).length(0))
            lhs183_ = d_1115_v0_
            lhs184_ = default__.safeIndex(3, (d_1115_v0_).length(0))
            lhs185_ = d_1267_v108_
            lhs186_ = default__.safeIndex(659, (d_1267_v108_).length(0))
            lhs187_ = globalState
            lhs188_ = globalState
            lhs181_[lhs182_] = rhs218_
            lhs183_[lhs184_] = rhs219_
            lhs185_[lhs186_] = rhs220_
            lhs187_.f9 = rhs221_
            lhs188_.f9 = rhs222_
            d_1270_v111_: _dafny.Map
            d_1270_v111_ = _dafny.Map({d_1115_v0_: self.f12})
            d_1270_v111_ = (d_1270_v111_).set(d_1115_v0_, (d_1241_v87_) == (d_1241_v87_))
        d_1271_v112_: str
        d_1271_v112_ = _dafny.CodePoint('x')
        d_1272_v113_: _dafny.Seq
        d_1272_v113_ = _dafny.SeqWithoutIsStrInference([not(self.f12), default__.fm0(d_1271_v112_, globalState)])
        d_1273_v114_: _dafny.Map
        d_1273_v114_ = _dafny.Map({not(self.f12): len(d_1272_v113_)})
        d_1274_v115_: _dafny.Seq
        d_1274_v115_ = _dafny.SeqWithoutIsStrInference([p0, (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tyfnjxj"))), (self).f13, ((d_1273_v114_)[self.f12] if (self.f12) in (d_1273_v114_) else (d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))])])
        index208_ = default__.safeIndex(3, (d_1115_v0_).length(0))
        (d_1115_v0_)[index208_] = (((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))]) + (len(d_1274_v115_))) + ((d_1115_v0_)[default__.safeIndex(3, (d_1115_v0_).length(0))])

    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C4(T0):
    def  __init__(self):
        self._f11: _dafny.MultiSet = _dafny.MultiSet({})
        self._f12: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f11(self):
        return self._f11
    @f11.setter
    def f11(self, value):
        self._f11 = value
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    def ctor__(self, f11, f12):
        (self).f11 = f11
        (self).f12 = f12

    def fm2(self, p0, p1, p2, p3, globalState):
        return (896) != (((0) - (-394)) * (46))

    def fm3(self, p0, p1, p2, p3, globalState):
        return 9

    def fm4(self, p0, p1, globalState):
        return ((-538 if self.f12 else -626)) == (((_dafny.MultiSet([908])) | (_dafny.MultiSet([-317, 163]))).cardinality)

    def m1(self, p0, globalState):
        r0: _dafny.Array = _dafny.Array(None, 0)
        r1: _dafny.Array = _dafny.Array(None, 0)
        r2: int = int(0)
        d_1275_v0_: int
        d_1275_v0_ = -835
        d_1276_v1_: _dafny.Seq
        d_1276_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fahubdda"))
        rhs223_ = (0) - ((default__.fm1(self.f12, d_1275_v0_, d_1275_v0_, globalState)) * (d_1275_v0_))
        rhs224_ = self.f12
        rhs225_ = (self.f11).set(self.f12, default__.abs((0) - ((36) + (d_1275_v0_))))
        rhs226_ = (default__.fm5(self.f12, self.f12, self.f12, d_1275_v0_, globalState)) + (default__.fm5(self.f12, self.f12, self.f12, len(d_1276_v1_), globalState))
        lhs189_ = globalState
        lhs190_ = globalState
        lhs191_ = self
        lhs192_ = globalState
        lhs189_.f9 = rhs223_
        lhs190_.f7 = rhs224_
        lhs191_.f11 = rhs225_
        lhs192_.f6 = rhs226_
        d_1277_v2_: _dafny.Array
        nw229_ = _dafny.Array(False, 3)
        d_1277_v2_ = nw229_
        d_1278_v3_: _dafny.Set
        d_1278_v3_ = _dafny.Set({d_1277_v2_, d_1277_v2_, d_1277_v2_, d_1277_v2_})
        d_1279_v4_: _dafny.Seq
        d_1279_v4_ = _dafny.SeqWithoutIsStrInference([d_1275_v0_, default__.fm1(self.f12, d_1275_v0_, d_1275_v0_, globalState), d_1275_v0_, (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kd")))) + (d_1275_v0_)])
        d_1280_v5_: _dafny.Seq
        d_1280_v5_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        d_1281_v6_: D1
        d_1281_v6_ = D1_DC1(p0)
        d_1282_v7_: _dafny.Array
        nw230_ = _dafny.Array(None, 29)
        nw230_[int(0)] = p0
        nw230_[int(1)] = p0
        nw230_[int(2)] = (p0).intersection(p0)
        nw230_[int(3)] = (p0) | (p0)
        nw230_[int(4)] = _dafny.Set({self.f12, (self.f12), self.f12, not(self.f12)})
        nw230_[int(5)] = (p0) | (p0)
        nw230_[int(6)] = p0
        nw230_[int(7)] = (p0) | (p0)
        nw230_[int(8)] = p0
        nw230_[int(9)] = (p0) - ((d_1280_v5_)[default__.safeIndex(d_1275_v0_, len(d_1280_v5_))])
        nw230_[int(10)] = p0
        nw230_[int(11)] = (d_1281_v6_).cf1
        nw230_[int(12)] = p0
        nw230_[int(13)] = p0
        nw230_[int(14)] = p0
        nw230_[int(15)] = p0
        nw230_[int(16)] = p0
        nw230_[int(17)] = (d_1280_v5_)[default__.safeIndex(-249, len(d_1280_v5_))]
        nw230_[int(18)] = p0
        nw230_[int(19)] = (p0).intersection(p0)
        nw230_[int(20)] = p0
        nw230_[int(21)] = (p0) - (_dafny.Set({self.f12, self.f12}))
        nw230_[int(22)] = (p0 if self.f12 else p0)
        nw230_[int(23)] = (_dafny.Set({not(False)})) | (p0)
        nw230_[int(24)] = (p0).intersection(_dafny.Set({self.f12}))
        nw230_[int(25)] = p0
        nw230_[int(26)] = _dafny.Set({self.f12, self.f12})
        nw230_[int(27)] = (p0) - (p0)
        nw230_[int(28)] = p0
        d_1282_v7_ = nw230_
        index209_ = default__.safeIndex(185, (d_1282_v7_).length(0))
        (d_1282_v7_)[index209_] = (p0) - (p0)
        d_1283_v8_: _dafny.Seq
        d_1283_v8_ = _dafny.SeqWithoutIsStrInference([d_1278_v3_])
        index210_ = default__.safeIndex(185, (d_1282_v7_).length(0))
        rhs227_ = (d_1283_v8_)[default__.safeIndex(default__.safeModuloInt(662, (0) - (d_1275_v0_)), len(d_1283_v8_))]
        rhs228_ = d_1279_v4_
        rhs229_ = p0
        rhs230_ = self.f12
        rhs231_ = 922
        lhs193_ = d_1282_v7_
        lhs194_ = default__.safeIndex(185, (d_1282_v7_).length(0))
        lhs195_ = globalState
        d_1278_v3_ = rhs227_
        d_1279_v4_ = rhs228_
        lhs193_[lhs194_] = rhs229_
        lhs195_.f3 = rhs230_
        d_1275_v0_ = rhs231_
        d_1277_v2_ = d_1277_v2_
        d_1284_i0_: int
        d_1284_i0_ = 0
        with _dafny.label("9"):
            while self.f12:
                with _dafny.c_label("9"):
                    if (d_1284_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1284_i0_ = (d_1284_i0_) + (1)
                    (globalState).f9 = (d_1275_v0_) - (d_1275_v0_)
                    (globalState).f3 = not(self.f12)
                    if self.f12:
                        d_1285_v9_: D1
                        d_1285_v9_ = D1_DC1((d_1282_v7_)[default__.safeIndex(185, (d_1282_v7_).length(0))])
                        d_1286_v10_: _dafny.MultiSet
                        d_1286_v10_ = _dafny.MultiSet([D1_DC3(d_1285_v9_)])
                        d_1287_v11_: D1
                        d_1287_v11_ = D1_DC3(d_1285_v9_)
                        (globalState).f3 = ((d_1286_v10_) - (d_1286_v10_)).ispropersubset((d_1286_v10_).set(d_1287_v11_, default__.abs(d_1275_v0_)))
                        d_1288_v12_: C3
                        nw231_ = C3()
                        nw231_.ctor__((0) - ((d_1275_v0_ if self.f12 else d_1275_v0_)), self.f12, (_dafny.MultiSet([self.f12, self.f12])) - (_dafny.MultiSet([not(self.f12), self.f12])), (self.f11).isdisjoint(self.f11))
                        d_1288_v12_ = nw231_
                        d_1289_v13_: _dafny.Map
                        d_1289_v13_ = _dafny.Map({(d_1288_v12_).f13: (0) - ((0) - ((d_1288_v12_).f13))})
                        d_1290_v14_: _dafny.Set
                        d_1290_v14_ = _dafny.Set({(d_1288_v12_).f13, (d_1288_v12_).f13, ((d_1289_v13_)[-708] if (-708) in (d_1289_v13_) else (d_1288_v12_).f13)})
                        rhs232_ = d_1290_v14_
                        rhs233_ = d_1275_v0_
                        rhs234_ = d_1275_v0_
                        rhs235_ = (d_1279_v4_ if (d_1288_v12_).f14 else (d_1279_v4_) + (d_1279_v4_))
                        lhs196_ = globalState
                        lhs197_ = globalState
                        d_1290_v14_ = rhs232_
                        lhs196_.f9 = rhs233_
                        lhs197_.f9 = rhs234_
                        d_1279_v4_ = rhs235_
                        d_1291_v15_: _dafny.Array
                        nw232_ = _dafny.Array(_dafny.Array(None, 0), 19)
                        d_1291_v15_ = nw232_
                        index211_ = default__.safeIndex(748, (d_1291_v15_).length(0))
                        (d_1291_v15_)[index211_] = (d_1277_v2_ if self.f12 else d_1277_v2_)
                        index212_ = default__.safeIndex(748, (d_1291_v15_).length(0))
                        rhs236_ = (d_1275_v0_) - (((d_1288_v12_).f13 if (d_1288_v12_).f14 else d_1275_v0_))
                        rhs237_ = d_1277_v2_
                        lhs198_ = d_1291_v15_
                        lhs199_ = default__.safeIndex(748, (d_1291_v15_).length(0))
                        d_1275_v0_ = rhs236_
                        lhs198_[lhs199_] = rhs237_
                        d_1292_v16_: _dafny.Map
                        d_1292_v16_ = _dafny.Map({(d_1279_v4_)[default__.safeIndex((d_1288_v12_).f13, len(d_1279_v4_))]: d_1276_v1_})
                        rhs238_ = (d_1288_v12_).f13
                        rhs239_ = (d_1288_v12_).f14
                        rhs240_ = (d_1288_v12_).f14
                        rhs241_ = ((d_1292_v16_)[d_1275_v0_] if (d_1275_v0_) in (d_1292_v16_) else (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tghqo"))) + (d_1276_v1_))
                        lhs200_ = globalState
                        lhs201_ = globalState
                        lhs202_ = globalState
                        lhs200_.f9 = rhs238_
                        lhs201_.f8 = rhs239_
                        lhs202_.f8 = rhs240_
                        d_1276_v1_ = rhs241_
                    elif True:
                        d_1293_v17_: D9
                        d_1293_v17_ = D9_DC25(d_1276_v1_, self.f12, d_1275_v0_)
                        d_1294_v18_: _dafny.Seq
                        d_1294_v18_ = _dafny.SeqWithoutIsStrInference([d_1293_v17_, d_1293_v17_])
                        d_1295_v19_: C3
                        nw233_ = C3()
                        nw233_.ctor__(d_1275_v0_, self.f12, self.f11, self.f12)
                        d_1295_v19_ = nw233_
                        d_1296_v20_: _dafny.Seq
                        d_1296_v20_ = _dafny.SeqWithoutIsStrInference([d_1295_v19_, d_1295_v19_])
                        d_1297_v21_: _dafny.Map
                        d_1297_v21_ = _dafny.Map({not((d_1295_v19_).f14): self.f12})
                        d_1298_v22_: _dafny.Seq
                        d_1298_v22_ = _dafny.SeqWithoutIsStrInference([True, self.f12, (d_1295_v19_).f14])
                        d_1299_v23_: str
                        d_1299_v23_ = _dafny.CodePoint('v')
                        d_1300_v24_: D1
                        d_1300_v24_ = D1_DC2((d_1295_v19_).f14, self.f12, False, d_1299_v23_, d_1299_v23_)
                        d_1301_v25_: T0
                        nw234_ = C1()
                        nw234_.ctor__(len(d_1298_v22_), _dafny.MultiSet([(d_1300_v24_).cf4]), not(True))
                        d_1301_v25_ = nw234_
                        d_1302_v26_: _dafny.Set
                        d_1302_v26_ = _dafny.Set({d_1275_v0_})
                        d_1303_v27_: _dafny.Array
                        def lambda49_(d_1304_v1_):
                            def lambda50_(d_1305_i1_):
                                return d_1304_v1_

                            return lambda50_

                        init31_ = lambda49_(d_1276_v1_)
                        nw235_ = _dafny.Array(None, 7)
                        for i0_31_ in range(nw235_.length(0)):
                            nw235_[i0_31_] = init31_(i0_31_)
                        d_1303_v27_ = nw235_
                        d_1306_v28_: D5
                        d_1306_v28_ = D5_DC13((d_1295_v19_).f13, True, (d_1295_v19_).f13, d_1303_v27_, d_1301_v25_.f12)
                        d_1307_v29_: _dafny.Map
                        d_1307_v29_ = _dafny.Map({(d_1306_v28_).cf34: (d_1295_v19_).f13})
                        d_1308_v30_: C0
                        nw236_ = C0()
                        nw236_.ctor__(True, d_1275_v0_, self.f11, d_1301_v25_.f12)
                        d_1308_v30_ = nw236_
                        d_1309_v31_: _dafny.Seq
                        d_1309_v31_ = _dafny.SeqWithoutIsStrInference([d_1308_v30_])
                        d_1310_v32_: _dafny.Array
                        nw237_ = _dafny.Array(None, 19)
                        nw237_[int(0)] = d_1275_v0_
                        nw237_[int(1)] = d_1275_v0_
                        nw237_[int(2)] = (((self.f11)[self.f12] if (self.f12) in (self.f11) else d_1275_v0_)) + (d_1275_v0_)
                        nw237_[int(3)] = len(d_1294_v18_)
                        nw237_[int(4)] = d_1275_v0_
                        nw237_[int(5)] = d_1275_v0_
                        nw237_[int(6)] = d_1275_v0_
                        nw237_[int(7)] = d_1275_v0_
                        nw237_[int(8)] = len(d_1296_v20_)
                        nw237_[int(9)] = len(d_1297_v21_)
                        nw237_[int(10)] = len(_dafny.Map({(d_1295_v19_).f13: d_1301_v25_}))
                        nw237_[int(11)] = default__.safeDivisionInt((d_1295_v19_).f13, (d_1295_v19_).f13)
                        nw237_[int(12)] = len((d_1302_v26_) - (d_1302_v26_))
                        nw237_[int(13)] = (0) - ((d_1295_v19_).f13)
                        nw237_[int(14)] = ((d_1307_v29_)[d_1301_v25_.f12] if (d_1301_v25_.f12) in (d_1307_v29_) else len(d_1309_v31_))
                        nw237_[int(15)] = (d_1308_v30_).f18
                        nw237_[int(16)] = 278
                        nw237_[int(17)] = (d_1295_v19_).f13
                        nw237_[int(18)] = (d_1295_v19_).f13
                        d_1310_v32_ = nw237_
                        index213_ = default__.safeIndex(355, (d_1310_v32_).length(0))
                        (d_1310_v32_)[index213_] = default__.safeDivisionInt(d_1275_v0_, d_1275_v0_)
                        index214_ = default__.safeIndex(355, (d_1310_v32_).length(0))
                        (d_1310_v32_)[index214_] = d_1275_v0_
                        d_1311_v33_: C0
                        nw238_ = C0()
                        nw238_.ctor__((d_1295_v19_).f14, default__.safeModuloInt(len(p0), (d_1295_v19_).f13), self.f11, ((_dafny.MultiSet([d_1279_v4_])).cardinality) != (d_1275_v0_))
                        d_1311_v33_ = nw238_
                        index215_ = default__.safeIndex(162, (d_1277_v2_).length(0))
                        (d_1277_v2_)[index215_] = (d_1308_v30_).f17
                        index216_ = default__.safeIndex(162, (d_1277_v2_).length(0))
                        (d_1277_v2_)[index216_] = (d_1308_v30_).f17
                        (globalState).f8 = (d_1301_v25_.f12 if ((d_1295_v19_).f13) == (591) else self.f12)
                        d_1312_v34_: _dafny.Map
                        d_1312_v34_ = _dafny.Map({(0) - (((d_1310_v32_)[default__.safeIndex(355, (d_1310_v32_).length(0))]) + (len(_dafny.SeqWithoutIsStrInference([d_1299_v23_ for d_1313_i2_ in range(default__.abs(792))])))): _dafny.SeqWithoutIsStrInference([(d_1295_v19_).f13 for d_1314_i3_ in range(default__.abs(336))])})
                        d_1312_v34_ = (d_1312_v34_ if ((d_1311_v33_).f17 if not((d_1277_v2_)[default__.safeIndex(162, (d_1277_v2_).length(0))]) else (d_1311_v33_).f17) else (_dafny.Map({(d_1311_v33_).f18: d_1279_v4_})) | ((_dafny.Map({(d_1310_v32_)[default__.safeIndex(355, (d_1310_v32_).length(0))]: d_1279_v4_})).set(907, d_1279_v4_)))
                    d_1276_v1_ = default__.fm9((_dafny.MultiSet([d_1275_v0_])).cardinality, self.f12, globalState)
                    pass
            pass
        d_1315_v35_: str
        d_1315_v35_ = _dafny.CodePoint('m')
        d_1316_v36_: _dafny.Array
        nw239_ = _dafny.Array(None, 5)
        nw239_[int(0)] = d_1315_v35_
        nw239_[int(1)] = d_1315_v35_
        nw239_[int(2)] = d_1315_v35_
        nw239_[int(3)] = d_1315_v35_
        nw239_[int(4)] = d_1315_v35_
        d_1316_v36_ = nw239_
        d_1317_v37_: D1
        d_1317_v37_ = D1_DC2(self.f12, self.f12, self.f12, d_1315_v35_, d_1315_v35_)
        d_1318_v38_: _dafny.Array
        nw240_ = _dafny.Array(_dafny.Array(None, 0), 19)
        d_1318_v38_ = nw240_
        d_1319_v39_: T0
        nw241_ = C3()
        nw241_.ctor__(d_1275_v0_, self.f12, self.f11, self.f12)
        d_1319_v39_ = nw241_
        d_1320_v40_: _dafny.Set
        d_1320_v40_ = _dafny.Set({d_1319_v39_, d_1319_v39_})
        d_1321_v41_: D8
        d_1321_v41_ = D8_DC21(True, d_1318_v38_, len(d_1276_v1_), d_1277_v2_, len(d_1320_v40_))
        index217_ = default__.safeIndex(164, (d_1316_v36_).length(0))
        (d_1316_v36_)[index217_] = default__.fm14(d_1317_v37_, (0) - ((d_1321_v41_).cf53), globalState)
        index218_ = default__.safeIndex(185, (d_1282_v7_).length(0))
        index219_ = default__.safeIndex(164, (d_1316_v36_).length(0))
        rhs242_ = len(d_1279_v4_)
        rhs243_ = (p0).intersection(_dafny.Set({True}))
        rhs244_ = d_1315_v35_
        rhs245_ = default__.fm24(globalState)
        lhs203_ = globalState
        lhs204_ = d_1282_v7_
        lhs205_ = default__.safeIndex(185, (d_1282_v7_).length(0))
        lhs206_ = d_1316_v36_
        lhs207_ = default__.safeIndex(164, (d_1316_v36_).length(0))
        lhs203_.f9 = rhs242_
        lhs204_[lhs205_] = rhs243_
        lhs206_[lhs207_] = rhs244_
        d_1279_v4_ = rhs245_
        hi3_ = d_1275_v0_
        for d_1322_i4_ in range(default__.safeDivisionInt(d_1275_v0_, d_1275_v0_), hi3_):
            index220_ = default__.safeIndex(185, (d_1282_v7_).length(0))
            (d_1282_v7_)[index220_] = p0
            d_1323_v42_: _dafny.Array
            nw242_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
            d_1323_v42_ = nw242_
            index221_ = default__.safeIndex(383, (d_1323_v42_).length(0))
            (d_1323_v42_)[index221_] = d_1276_v1_
            index222_ = default__.safeIndex(383, (d_1323_v42_).length(0))
            (d_1323_v42_)[index222_] = d_1276_v1_
            d_1324_v43_: _dafny.Seq
            d_1324_v43_ = _dafny.SeqWithoutIsStrInference([self.f12])
            d_1325_v44_: _dafny.Map
            d_1325_v44_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1319_v39_.f12, self.f12])): d_1324_v43_})
            d_1326_v45_: _dafny.Map
            d_1326_v45_ = _dafny.Map({(d_1316_v36_)[default__.safeIndex(164, (d_1316_v36_).length(0))]: d_1319_v39_.f12})
            rhs246_ = ((d_1325_v44_)[d_1319_v39_.f11] if (d_1319_v39_.f11) in (d_1325_v44_) else d_1324_v43_)
            rhs247_ = d_1275_v0_
            rhs248_ = ((d_1316_v36_)[default__.safeIndex(164, (d_1316_v36_).length(0))]) in ((d_1326_v45_) | (d_1326_v45_))
            lhs208_ = globalState
            lhs209_ = globalState
            lhs208_.f6 = rhs246_
            d_1275_v0_ = rhs247_
            lhs209_.f7 = rhs248_
            d_1327_v46_: _dafny.Map
            d_1327_v46_ = _dafny.Map({d_1319_v39_: d_1319_v39_})
            d_1327_v46_ = _dafny.Map({d_1319_v39_: d_1319_v39_})
        d_1328_v47_: _dafny.Array
        nw243_ = _dafny.Array(_dafny.Map({}), 5)
        d_1328_v47_ = nw243_
        r0 = d_1328_v47_
        d_1329_v48_: _dafny.Array
        nw244_ = _dafny.Array(_dafny.Seq({}), 10)
        d_1329_v48_ = nw244_
        r1 = d_1329_v48_
        r2 = d_1275_v0_
        return r0, r1, r2

