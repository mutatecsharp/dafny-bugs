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
    def fm0(p0, p1, p2, p3, globalState):
        return False

    @staticmethod
    def fm2(p0, globalState):
        return (0) - (default__.safeDivisionInt(-490, default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guwwkcoj"))])), -612)))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_0_i0_ in range(default__.abs(309))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dwtkddnh")))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kq"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1_i1_ in range(default__.abs(218))])))

    @staticmethod
    def fm4(globalState):
        return (0) - (((512) + (len(_dafny.Map({True: _dafny.CodePoint('s')}))) if not((True) and (False)) else default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({268: True})), 589, len(_dafny.Map({False: False})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddwiwxcq")))])), -555)))

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return D1_DC3(_dafny.CodePoint('d'))

    @staticmethod
    def fm6(p0, p1, p2, p3, globalState):
        return _dafny.CodePoint('s')

    @staticmethod
    def fm9(p0, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2_i0_ in range(default__.abs(795))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))

    @staticmethod
    def fm10(globalState):
        def iife0_():
            def iife4_():
                def iife6_():
                    coll6_ = _dafny.Set()
                    compr_6_: int
                    for compr_6_ in _dafny.IntegerRange(66, 817):
                        d_9_v0_: int = compr_6_
                        if ((66) <= (d_9_v0_)) and ((d_9_v0_) < (817)):
                            coll6_ = coll6_.union(_dafny.Set([default__.safeModuloInt(d_9_v0_, 401)]))
                    return _dafny.Set(coll6_)
                coll4_ = _dafny.Set()
                def iife5_():
                    coll5_ = _dafny.Set()
                    compr_5_: int
                    for compr_5_ in _dafny.IntegerRange(66, 817):
                        d_7_v0_: int = compr_5_
                        if ((66) <= (d_7_v0_)) and ((d_7_v0_) < (817)):
                            coll5_ = coll5_.union(_dafny.Set([default__.safeModuloInt(d_7_v0_, 401)]))
                    return _dafny.Set(coll5_)
                compr_4_: int
                for compr_4_ in (iife5_()
                ).Elements:
                    d_8_v1_: int = compr_4_
                    if (d_8_v1_) in (iife6_()
                    ):
                        coll4_ = coll4_.union(_dafny.Set([default__.safeModuloInt(d_8_v1_, len(_dafny.Set({True, True})))]))
                return _dafny.Set(coll4_)
            coll0_ = _dafny.Set()
            def iife1_():
                def iife3_():
                    coll3_ = _dafny.Set()
                    compr_3_: int
                    for compr_3_ in _dafny.IntegerRange(66, 817):
                        d_5_v0_: int = compr_3_
                        if ((66) <= (d_5_v0_)) and ((d_5_v0_) < (817)):
                            coll3_ = coll3_.union(_dafny.Set([default__.safeModuloInt(d_5_v0_, 401)]))
                    return _dafny.Set(coll3_)
                coll1_ = _dafny.Set()
                def iife2_():
                    coll2_ = _dafny.Set()
                    compr_2_: int
                    for compr_2_ in _dafny.IntegerRange(66, 817):
                        d_3_v0_: int = compr_2_
                        if ((66) <= (d_3_v0_)) and ((d_3_v0_) < (817)):
                            coll2_ = coll2_.union(_dafny.Set([default__.safeModuloInt(d_3_v0_, 401)]))
                    return _dafny.Set(coll2_)
                compr_1_: int
                for compr_1_ in (iife2_()
                ).Elements:
                    d_4_v1_: int = compr_1_
                    if (d_4_v1_) in (iife3_()
                    ):
                        coll1_ = coll1_.union(_dafny.Set([default__.safeModuloInt(d_4_v1_, len(_dafny.Set({True, True})))]))
                return _dafny.Set(coll1_)
            compr_0_: int
            for compr_0_ in (_dafny.Map({(_dafny.MultiSet([144, len(iife1_()
            ), (0) - (len(_dafny.SeqWithoutIsStrInference([True])))])).cardinality: True})).keys.Elements:
                d_6_v2_: int = compr_0_
                if (d_6_v2_) in (_dafny.Map({(_dafny.MultiSet([144, len(iife4_()
                ), (0) - (len(_dafny.SeqWithoutIsStrInference([True])))])).cardinality: True})):
                    coll0_ = coll0_.union(_dafny.Set([default__.safeDivisionInt(d_6_v2_, (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vup")))))]))
            return _dafny.Set(coll0_)
        def iife7_():
            coll7_ = _dafny.Set()
            compr_7_: int
            for compr_7_ in _dafny.IntegerRange(704, -431):
                d_10_v3_: int = compr_7_
                if ((704) <= (d_10_v3_)) and ((d_10_v3_) < (-431)):
                    coll7_ = coll7_.union(_dafny.Set([(d_10_v3_) - (-676)]))
            return _dafny.Set(coll7_)
        return (iife0_()
        ).intersection((iife7_()
         if not(True) else _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdem"))), len(_dafny.SeqWithoutIsStrInference([True, True, True, not((D0_DC0(True)).cf0), False]))})))

    @staticmethod
    def fm11(p0, p1, globalState):
        return ((_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True, False, False, not(not(not(False))), True]))])) | (_dafny.MultiSet([-769, len(_dafny.SeqWithoutIsStrInference([798]))]))) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-619]))) | (_dafny.MultiSet([len(_dafny.Map({True: True})), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_11_i0_ in range(default__.abs(752))])), 803, len(_dafny.Map({(D1_DC5(False, False)).cf8: True})), 648])))

    @staticmethod
    def fm12(p0, globalState):
        if (True if True else True):
            return (_dafny.SeqWithoutIsStrInference([-402])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mbafox"))), len(_dafny.Set({_dafny.Map({not(False): not(not(False))}), _dafny.Map({False: False}), _dafny.Map({True: True}), _dafny.Map({True: True})}))]))
        elif True:
            return _dafny.SeqWithoutIsStrInference([-643])

    @staticmethod
    def fm13(p0, p1, globalState):
        return D1_DC5(True, (_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([True])})).issubset(_dafny.Set({_dafny.SeqWithoutIsStrInference([False]), _dafny.SeqWithoutIsStrInference([not(True), True, False, True, True])})))

    @staticmethod
    def fm14(p0, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(920, 101):
                d_12_v0_: int = compr_8_
                if ((920) <= (d_12_v0_)) and ((d_12_v0_) < (101)):
                    coll8_[(d_12_v0_) + (len(_dafny.SeqWithoutIsStrInference([False])))] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfuahe"))) for d_13_i0_ in range(default__.abs(621))]))).cardinality
            return _dafny.Map(coll8_)
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D11_DC25(), D11_DC25(), D11_DC25()])), -411])).Elements:
                d_14_v1_: int = compr_9_
                if (d_14_v1_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([D11_DC25(), D11_DC25(), D11_DC25()])), -411])):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeModuloInt(d_14_v1_, 280)]))
            return _dafny.Set(coll9_)
        return (_dafny.Map({False: iife8_()
        })) | ((_dafny.Map({True: _dafny.Map({len(iife9_()
        ): 114})})) | (_dafny.Map({False: _dafny.Map({424: 154})})))

    @staticmethod
    def fm15(p0, p1, p2, globalState):
        return ((_dafny.Set({True})) - (_dafny.Set({False}))) - ((_dafny.Set({False})).intersection(_dafny.Set({False})))

    @staticmethod
    def fm16(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([not((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g')])) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_15_i0_ in range(default__.abs(258))]))), True, (False) or (False)])

    @staticmethod
    def fm17(p0, p1, p2, p3, globalState):
        return D1_DC4((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrkvxhch")) for d_16_i0_ in range(default__.abs(216))]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ugtro"))]))), 149)

    @staticmethod
    def fm18(p0, globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odvbie"))})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "efm")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ctebj"))}))) - ((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmdxukk")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dlv"))})) - (_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "affp")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))})))

    @staticmethod
    def fm19(globalState):
        def iife10_():
            coll10_ = _dafny.Map()
            compr_10_: int
            for compr_10_ in _dafny.IntegerRange(932, 302):
                d_17_v0_: int = compr_10_
                if ((932) <= (d_17_v0_)) and ((d_17_v0_) < (302)):
                    coll10_[default__.safeDivisionInt(d_17_v0_, (0) - (-106))] = False
            return _dafny.Map(coll10_)
        return (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cjrbft")))): 587})) | ((_dafny.Map({-296: 464})) | (_dafny.Map({537: len(_dafny.SeqWithoutIsStrInference([len(iife10_()
        )]))})))

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return (_dafny.Map({len(_dafny.Set({not(False), True})): False})) | ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality]), _dafny.SeqWithoutIsStrInference([795]), _dafny.SeqWithoutIsStrInference([521])])): not(False)})) | (_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({505: True})), len(_dafny.Map({_dafny.CodePoint('q'): len(_dafny.SeqWithoutIsStrInference([not(True)]))})), 265])).cardinality: not(False)})))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        source0_ = D0_DC1((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_18_i0_ in range(default__.abs(908))]))), -506, 562)
        if source0_.is_DC1:
            d_19___mcc_h0_ = source0_.cf1
            d_20___mcc_h1_ = source0_.cf2
            d_21___mcc_h2_ = source0_.cf3
            d_22_cf3_ = d_21___mcc_h2_
            d_23_cf2_ = d_20___mcc_h1_
            d_24_cf1_ = d_19___mcc_h0_
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: D1
                for compr_11_ in (_dafny.SeqWithoutIsStrInference([D1_DC6(D1_DC3(_dafny.CodePoint('f'))), D1_DC6(D1_DC5(False, True))])).Elements:
                    d_25_v0_: D1 = compr_11_
                    if (d_25_v0_) in (_dafny.SeqWithoutIsStrInference([D1_DC6(D1_DC3(_dafny.CodePoint('f'))), D1_DC6(D1_DC5(False, True))])):
                        coll11_[d_25_v0_] = _dafny.CodePoint('y')
                return _dafny.Map(coll11_)
            return iife11_()
            
        elif source0_.is_DC0:
            d_26___mcc_h3_ = source0_.cf0
            d_27_cf0_ = d_26___mcc_h3_
            return _dafny.Map({D1_DC6(D1_DC4(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_28_i1_ in range(default__.abs(888))])])), 720)): _dafny.CodePoint('k')})
        elif True:
            d_29___mcc_h4_ = source0_.cf4
            d_30_cf4_ = d_29___mcc_h4_
            return (D15_DC31(_dafny.Map({D1_DC6(D1_DC6(D1_DC3(_dafny.CodePoint('g')))): _dafny.CodePoint('r')}))).cf43

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_31_v0_: bool
        d_31_v0_ = False
        d_32_v1_: _dafny.Seq
        d_32_v1_ = _dafny.SeqWithoutIsStrInference([d_31_v0_, d_31_v0_, d_31_v0_])
        d_33_v2_: _dafny.Map
        d_33_v2_ = _dafny.Map({d_31_v0_: d_32_v1_})
        d_33_v2_ = (d_33_v2_).set(True, d_32_v1_)
        if (d_32_v1_)[default__.safeIndex(p3, len(d_32_v1_))]:
            d_34_v3_: D3
            d_34_v3_ = D3_DC10(d_31_v0_, p3, d_31_v0_, p0, _dafny.Map({p1: not(d_31_v0_)}))
            d_35_v4_: D3
            d_35_v4_ = D3_DC11(d_34_v3_)
            d_35_v4_ = d_35_v4_
            d_36_v5_: D1
            d_36_v5_ = D1_DC5(d_31_v0_, d_31_v0_)
            d_31_v0_ = (d_36_v5_).cf8
            d_37_v6_: D1
            d_37_v6_ = D1_DC5(d_31_v0_, d_31_v0_)
            d_38_v7_: D1
            d_38_v7_ = D1_DC6(d_37_v6_)
            d_39_v8_: C0
            nw0_ = C0()
            nw0_.ctor__()
            d_39_v8_ = nw0_
            d_40_v9_: _dafny.MultiSet
            d_40_v9_ = _dafny.MultiSet([d_31_v0_, True])
            d_41_v10_: _dafny.Seq
            d_41_v10_ = _dafny.SeqWithoutIsStrInference([p1])
            d_42_v11_: _dafny.Seq
            d_42_v11_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_40_v9_).cardinality), p3, len(d_41_v10_)])
            d_43_v12_: _dafny.Seq
            d_43_v12_ = _dafny.SeqWithoutIsStrInference([p1, len((d_42_v11_).set(default__.safeIndex(p3, len(d_42_v11_)), p1))])
            d_44_v13_: C2
            nw1_ = C2()
            nw1_.ctor__(p1)
            d_44_v13_ = nw1_
            d_45_v14_: _dafny.Map
            d_45_v14_ = _dafny.Map({p1: d_44_v13_})
            d_46_v15_: _dafny.Map
            d_46_v15_ = _dafny.Map({d_31_v0_: len(d_45_v14_)})
            d_47_v16_: C1
            nw2_ = C1()
            nw2_.ctor__(False)
            d_47_v16_ = nw2_
            d_48_v17_: _dafny.Map
            d_48_v17_ = _dafny.Map({d_31_v0_: d_47_v16_})
            d_49_v18_: _dafny.Array
            nw3_ = _dafny.Array(None, 24)
            nw3_[int(0)] = p1
            nw3_[int(1)] = p3
            nw3_[int(2)] = p3
            nw3_[int(3)] = (d_43_v12_)[default__.safeIndex(p1, len(d_43_v12_))]
            nw3_[int(4)] = p1
            nw3_[int(5)] = p3
            nw3_[int(6)] = p1
            nw3_[int(7)] = p3
            nw3_[int(8)] = p1
            nw3_[int(9)] = p3
            nw3_[int(10)] = len(d_41_v10_)
            nw3_[int(11)] = p3
            nw3_[int(12)] = p3
            nw3_[int(13)] = (0) - (p1)
            nw3_[int(14)] = 1
            nw3_[int(15)] = p1
            nw3_[int(16)] = ((d_46_v15_)[d_31_v0_] if (d_31_v0_) in (d_46_v15_) else p3)
            nw3_[int(17)] = 597
            nw3_[int(18)] = p3
            nw3_[int(19)] = p1
            nw3_[int(20)] = -469
            nw3_[int(21)] = d_44_v13_.f13
            nw3_[int(22)] = 896
            nw3_[int(23)] = len(d_48_v17_)
            d_49_v18_ = nw3_
            d_50_v19_: _dafny.Seq
            d_50_v19_ = _dafny.SeqWithoutIsStrInference([d_49_v18_, d_49_v18_])
            d_51_v20_: _dafny.Array
            nw4_ = _dafny.Array(None, 22)
            nw4_[int(0)] = d_49_v18_
            nw4_[int(1)] = d_49_v18_
            nw4_[int(2)] = d_49_v18_
            nw4_[int(3)] = d_49_v18_
            nw4_[int(4)] = d_49_v18_
            nw4_[int(5)] = d_49_v18_
            nw4_[int(6)] = d_49_v18_
            nw4_[int(7)] = d_49_v18_
            nw4_[int(8)] = (d_50_v19_)[default__.safeIndex(p3, len(d_50_v19_))]
            nw4_[int(9)] = d_49_v18_
            nw4_[int(10)] = d_49_v18_
            nw4_[int(11)] = d_49_v18_
            nw4_[int(12)] = d_49_v18_
            nw4_[int(13)] = d_49_v18_
            nw4_[int(14)] = d_49_v18_
            nw4_[int(15)] = d_49_v18_
            nw4_[int(16)] = d_49_v18_
            nw4_[int(17)] = d_49_v18_
            nw4_[int(18)] = d_49_v18_
            nw4_[int(19)] = d_49_v18_
            nw4_[int(20)] = d_49_v18_
            nw4_[int(21)] = d_49_v18_
            d_51_v20_ = nw4_
            rhs0_ = d_38_v7_
            rhs1_ = d_49_v18_
            rhs2_ = d_31_v0_
            rhs3_ = d_39_v8_
            rhs4_ = d_51_v20_
            lhs0_ = globalState
            d_38_v7_ = rhs0_
            lhs0_.f10 = rhs1_
            d_31_v0_ = rhs2_
            d_39_v8_ = rhs3_
            d_51_v20_ = rhs4_
            d_52_v21_: _dafny.Seq
            d_52_v21_ = d_32_v1_
            d_53_v22_: _dafny.Array
            nw5_ = _dafny.Array(None, 13)
            nw5_[int(0)] = (d_32_v1_ if d_47_v16_.f14 else d_32_v1_)
            nw5_[int(1)] = d_32_v1_
            nw5_[int(2)] = d_32_v1_
            nw5_[int(3)] = (_dafny.SeqWithoutIsStrInference([d_47_v16_.f14, d_47_v16_.f14, d_47_v16_.f14, d_47_v16_.f14])) + (d_32_v1_)
            nw5_[int(4)] = _dafny.SeqWithoutIsStrInference([d_47_v16_.f14])
            nw5_[int(5)] = d_32_v1_
            nw5_[int(6)] = d_32_v1_
            nw5_[int(7)] = (d_32_v1_) + (d_32_v1_)
            nw5_[int(8)] = _dafny.SeqWithoutIsStrInference([d_31_v0_, d_47_v16_.f14])
            nw5_[int(9)] = (d_32_v1_) + (d_32_v1_)
            nw5_[int(10)] = (d_32_v1_) + (d_32_v1_)
            nw5_[int(11)] = ((d_52_v21_)) + (d_32_v1_)
            nw5_[int(12)] = d_32_v1_
            d_53_v22_ = nw5_
            d_54_v23_: str
            d_54_v23_ = _dafny.CodePoint('l')
            rhs5_ = d_53_v22_
            rhs6_ = d_54_v23_
            rhs7_ = (p3) >= ((0) - (d_44_v13_.f13))
            lhs1_ = d_47_v16_
            d_53_v22_ = rhs5_
            d_54_v23_ = rhs6_
            lhs1_.f14 = rhs7_
            d_55_v24_: _dafny.Set
            d_55_v24_ = _dafny.Set({d_32_v1_})
            d_55_v24_ = (d_55_v24_) | (d_55_v24_)
        elif True:
            index0_ = default__.safeIndex(491, (p0).length(0))
            (p0)[index0_] = d_31_v0_
            d_56_v25_: _dafny.Array
            nw6_ = _dafny.Array(_dafny.Seq({}), 10)
            d_56_v25_ = nw6_
            d_57_v26_: _dafny.Seq
            d_57_v26_ = _dafny.SeqWithoutIsStrInference([p1, p3])
            index1_ = default__.safeIndex(855, (d_56_v25_).length(0))
            (d_56_v25_)[index1_] = d_57_v26_
            d_58_v27_: _dafny.Seq
            d_58_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwcctfdwq"))
            d_59_v28_: _dafny.Map
            d_59_v28_ = _dafny.Map({p3: not(d_31_v0_)})
            d_60_v29_: D7
            d_60_v29_ = D7_DC17(d_31_v0_)
            d_61_v30_: D11
            d_61_v30_ = D11_DC23(d_57_v26_)
            index2_ = default__.safeIndex(491, (p0).length(0))
            index3_ = default__.safeIndex(855, (d_56_v25_).length(0))
            rhs8_ = ((233) in (d_57_v26_)) and ((d_58_v27_) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))))
            rhs9_ = len((d_59_v28_) | (default__.fm20(d_31_v0_, (0) - (p3), default__.fm0(True, d_31_v0_, (d_60_v29_).cf25, d_31_v0_, globalState), globalState)))
            rhs10_ = (d_31_v0_ if (d_31_v0_ if not(default__.fm0(d_31_v0_, d_31_v0_, d_31_v0_, d_31_v0_, globalState)) else False) else (d_31_v0_ if d_31_v0_ else d_31_v0_))
            rhs11_ = ((d_61_v30_).cf34).set(default__.safeIndex((p1) * (p3), len((d_61_v30_).cf34)), (0) - (p1))
            lhs2_ = p0
            lhs3_ = default__.safeIndex(491, (p0).length(0))
            lhs4_ = globalState
            lhs5_ = d_56_v25_
            lhs6_ = default__.safeIndex(855, (d_56_v25_).length(0))
            lhs2_[lhs3_] = rhs8_
            lhs4_.f2 = rhs9_
            d_31_v0_ = rhs10_
            lhs5_[lhs6_] = rhs11_
            index4_ = default__.safeIndex(491, (p0).length(0))
            (p0)[index4_] = d_31_v0_
            d_62_v31_: C1
            nw7_ = C1()
            nw7_.ctor__(d_31_v0_)
            d_62_v31_ = nw7_
            d_63_v32_: _dafny.Map
            d_63_v32_ = _dafny.Map({(p1) <= (p1): d_62_v31_})
            d_64_v33_: _dafny.Map
            d_64_v33_ = _dafny.Map({(d_57_v26_)[default__.safeIndex(p1, len(d_57_v26_))]: d_62_v31_})
            d_62_v31_ = ((d_63_v32_)[(p0)[default__.safeIndex(491, (p0).length(0))]] if ((p0)[default__.safeIndex(491, (p0).length(0))]) in (d_63_v32_) else ((d_64_v33_)[p3] if (p3) in (d_64_v33_) else d_62_v31_))
            (globalState).f2 = ((d_56_v25_)[default__.safeIndex(855, (d_56_v25_).length(0))])[default__.safeIndex(len(d_58_v27_), len((d_56_v25_)[default__.safeIndex(855, (d_56_v25_).length(0))]))]
            d_65_v34_: _dafny.MultiSet
            d_65_v34_ = _dafny.MultiSet([d_31_v0_, d_31_v0_])
            d_66_v35_: C2
            nw8_ = C2()
            nw8_.ctor__((d_65_v34_).cardinality)
            d_66_v35_ = nw8_
            d_67_v36_: str
            d_67_v36_ = _dafny.CodePoint('x')
            d_68_v37_: _dafny.Map
            d_68_v37_ = _dafny.Map({d_66_v35_: d_67_v36_})
            (globalState).f2 = (p1) - (len(d_68_v37_))
        if d_31_v0_:
            d_31_v0_ = d_31_v0_
            d_69_v38_: C0
            nw9_ = C0()
            nw9_.ctor__()
            d_69_v38_ = nw9_
            d_70_v39_: _dafny.Seq
            d_70_v39_ = _dafny.SeqWithoutIsStrInference([d_69_v38_, d_69_v38_])
            d_71_v40_: _dafny.Seq
            d_71_v40_ = d_70_v39_
            d_72_v41_: _dafny.Map
            d_72_v41_ = _dafny.Map({(d_71_v40_): (0) - (p1)})
            d_73_v42_: _dafny.Map
            d_73_v42_ = _dafny.Map({718: default__.fm4(globalState)})
            d_74_v43_: _dafny.Map
            d_74_v43_ = _dafny.Map({d_31_v0_: d_31_v0_})
            d_75_v44_: _dafny.Seq
            d_75_v44_ = _dafny.SeqWithoutIsStrInference([d_74_v43_])
            d_72_v41_ = (d_72_v41_).set((d_70_v39_).set(default__.safeIndex(p1, len(d_70_v39_)), d_69_v38_), (((d_73_v42_)[len(d_75_v44_)] if (len(d_75_v44_)) in (d_73_v42_) else p3)) - (p1))
            d_76_v45_: _dafny.Set
            d_76_v45_ = _dafny.Set({p3})
            d_77_v46_: _dafny.Map
            d_77_v46_ = _dafny.Map({not((d_76_v45_).ispropersubset(_dafny.Set({p1}))): p3})
            d_77_v46_ = (d_77_v46_).set(default__.fm0(d_31_v0_, d_31_v0_, not(d_31_v0_), d_31_v0_, globalState), p3)
            d_78_v47_: D11
            d_78_v47_ = D11_DC24(p3)
            d_78_v47_ = d_78_v47_
            if not(d_31_v0_):
                d_79_v48_: C0
                nw10_ = C0()
                nw10_.ctor__()
                d_79_v48_ = nw10_
                d_31_v0_ = d_31_v0_
                d_80_v49_: _dafny.Array
                nw11_ = _dafny.Array(int(0), 26)
                d_80_v49_ = nw11_
                d_81_v50_: C2
                nw12_ = C2()
                nw12_.ctor__(p3)
                d_81_v50_ = nw12_
                d_82_v51_: _dafny.Map
                d_82_v51_ = _dafny.Map({d_81_v50_: d_80_v49_})
                d_83_v52_: _dafny.Array
                nw13_ = _dafny.Array(None, 25)
                nw13_[int(0)] = d_80_v49_
                nw13_[int(1)] = d_80_v49_
                nw13_[int(2)] = d_80_v49_
                nw13_[int(3)] = d_80_v49_
                nw13_[int(4)] = d_80_v49_
                nw13_[int(5)] = d_80_v49_
                nw13_[int(6)] = d_80_v49_
                nw13_[int(7)] = d_80_v49_
                nw13_[int(8)] = d_80_v49_
                nw13_[int(9)] = d_80_v49_
                nw13_[int(10)] = d_80_v49_
                nw13_[int(11)] = d_80_v49_
                nw13_[int(12)] = d_80_v49_
                nw13_[int(13)] = d_80_v49_
                nw13_[int(14)] = ((d_82_v51_)[d_81_v50_] if (d_81_v50_) in (d_82_v51_) else d_80_v49_)
                nw13_[int(15)] = d_80_v49_
                nw13_[int(16)] = d_80_v49_
                nw13_[int(17)] = d_80_v49_
                nw13_[int(18)] = d_80_v49_
                nw13_[int(19)] = d_80_v49_
                nw13_[int(20)] = d_80_v49_
                nw13_[int(21)] = d_80_v49_
                nw13_[int(22)] = d_80_v49_
                nw13_[int(23)] = d_80_v49_
                nw13_[int(24)] = d_80_v49_
                d_83_v52_ = nw13_
                d_84_v53_: _dafny.Map
                d_84_v53_ = _dafny.Map({_dafny.CodePoint('k'): d_80_v49_})
                d_85_v54_: str
                d_85_v54_ = _dafny.CodePoint('g')
                index5_ = default__.safeIndex(704, (d_83_v52_).length(0))
                (d_83_v52_)[index5_] = ((d_84_v53_)[d_85_v54_] if (d_85_v54_) in (d_84_v53_) else d_80_v49_)
                index6_ = default__.safeIndex(704, (d_83_v52_).length(0))
                (d_83_v52_)[index6_] = d_80_v49_
                d_86_v55_: _dafny.Map
                d_86_v55_ = _dafny.Map({d_69_v38_: len(d_77_v46_)})
                d_86_v55_ = (d_86_v55_).set(d_69_v38_, p3)
                (globalState).f2 = d_81_v50_.f13
            elif True:
                d_87_v56_: _dafny.Array
                def lambda0_(d_88_p3_):
                    def lambda1_(d_89_i0_):
                        return default__.safeModuloInt(d_89_i0_, d_88_p3_)

                    return lambda1_

                init0_ = lambda0_(p3)
                nw14_ = _dafny.Array(None, 12)
                for i0_0_ in range(nw14_.length(0)):
                    nw14_[i0_0_] = init0_(i0_0_)
                d_87_v56_ = nw14_
                d_90_v57_: _dafny.Map
                d_90_v57_ = _dafny.Map({d_87_v56_: p1})
                d_91_v58_: _dafny.MultiSet
                d_91_v58_ = _dafny.MultiSet([p3, p3])
                d_92_v59_: _dafny.Set
                d_92_v59_ = _dafny.Set({d_31_v0_, d_31_v0_, d_31_v0_})
                d_73_v42_ = (d_73_v42_).set(((d_90_v57_)[d_87_v56_] if (d_87_v56_) in (d_90_v57_) else (d_69_v38_).fm8((d_91_v58_).cardinality, d_92_v59_, p1, globalState)), p1)
                d_93_v60_: _dafny.Array
                def lambda2_(d_94_i1_):
                    return D1_DC6(D1_DC6(D1_DC3(_dafny.CodePoint('h'))))

                init1_ = lambda2_
                nw15_ = _dafny.Array(None, 2)
                for i0_1_ in range(nw15_.length(0)):
                    nw15_[i0_1_] = init1_(i0_1_)
                d_93_v60_ = nw15_
                d_95_v61_: D2
                d_95_v61_ = D2_DC7(d_93_v60_)
                d_96_v62_: str
                d_96_v62_ = _dafny.CodePoint('m')
                rhs12_ = (p1) - (default__.safeDivisionInt(p3, (default__.fm11(d_31_v0_, d_96_v62_, globalState)).cardinality))
                rhs13_ = ((0) - (p3)) >= ((0) - (p1))
                rhs14_ = d_95_v61_
                lhs7_ = globalState
                lhs7_.f2 = rhs12_
                d_31_v0_ = rhs13_
                d_95_v61_ = rhs14_
                (globalState).f2 = (0) - ((0) - (p1))
                d_97_v63_: C2
                nw16_ = C2()
                nw16_.ctor__(p1)
                d_97_v63_ = nw16_
                index7_ = default__.safeIndex(895, (d_87_v56_).length(0))
                (d_87_v56_)[index7_] = default__.fm2(d_31_v0_, globalState)
                index8_ = default__.safeIndex(895, (d_87_v56_).length(0))
                (d_87_v56_)[index8_] = d_97_v63_.f13
        elif True:
            d_98_v64_: D7
            d_98_v64_ = D7_DC18(D7_DC17(d_31_v0_))
            source1_ = d_98_v64_
            if source1_.is_DC17:
                d_99___mcc_h0_ = source1_.cf25
                d_100_cf25_ = d_99___mcc_h0_
                d_100_cf25_ = ((p1) != (p3)) and (not(d_31_v0_))
                d_101_v65_: _dafny.Map
                d_101_v65_ = _dafny.Map({p1: d_100_cf25_})
                d_102_v66_: _dafny.Seq
                d_102_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dpwnr"))
                d_103_v67_: _dafny.Map
                d_103_v67_ = _dafny.Map({not((D3_DC10(d_31_v0_, p3, d_31_v0_, p0, d_101_v65_)).cf16): d_102_v66_})
                d_104_v68_: _dafny.Seq
                d_104_v68_ = _dafny.SeqWithoutIsStrInference([len(d_103_v67_)])
                rhs15_ = (d_104_v68_)[default__.safeIndex(808, len(d_104_v68_))]
                rhs16_ = p1
                lhs8_ = globalState
                lhs9_ = globalState
                lhs8_.f2 = rhs15_
                lhs9_.f2 = rhs16_
                d_105_v69_: C2
                nw17_ = C2()
                nw17_.ctor__((p1 if d_31_v0_ else p1))
                d_105_v69_ = nw17_
                d_105_v69_ = d_105_v69_
                d_106_v70_: C3
                nw18_ = C3()
                nw18_.ctor__(d_100_cf25_, 7)
                d_106_v70_ = nw18_
            elif source1_.is_DC16:
                d_107___mcc_h1_ = source1_.cf24
                d_108_cf24_ = d_107___mcc_h1_
                d_109_v71_: _dafny.Array
                def lambda3_(d_110_i2_):
                    return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q'), _dafny.CodePoint('q'), _dafny.CodePoint('a'), _dafny.CodePoint('y'), _dafny.CodePoint('f')])

                init2_ = lambda3_
                nw19_ = _dafny.Array(None, 23)
                for i0_2_ in range(nw19_.length(0)):
                    nw19_[i0_2_] = init2_(i0_2_)
                d_109_v71_ = nw19_
                d_111_v72_: _dafny.Array
                def lambda4_(d_112_p3_):
                    def lambda5_(d_113_i3_):
                        return (d_113_i3_) - ((0) - (d_112_p3_))

                    return lambda5_

                init3_ = lambda4_(p3)
                nw20_ = _dafny.Array(None, 14)
                for i0_3_ in range(nw20_.length(0)):
                    nw20_[i0_3_] = init3_(i0_3_)
                d_111_v72_ = nw20_
                d_114_v73_: _dafny.Seq
                d_114_v73_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "joaugkus"))
                index9_ = default__.safeIndex(235, (d_111_v72_).length(0))
                (d_111_v72_)[index9_] = len(d_114_v73_)
                d_115_v74_: _dafny.Map
                d_115_v74_ = _dafny.Map({(p1) != (p3): (d_114_v73_) + (d_114_v73_)})
                d_116_v75_: _dafny.Array
                nw21_ = _dafny.Array(_dafny.CodePoint('D'), 16)
                d_116_v75_ = nw21_
                index10_ = default__.safeIndex(235, (d_111_v72_).length(0))
                rhs17_ = (d_109_v71_ if d_31_v0_ else d_109_v71_)
                rhs18_ = p1
                rhs19_ = d_109_v71_
                rhs20_ = (d_115_v74_) | (d_115_v74_)
                rhs21_ = d_116_v75_
                lhs10_ = d_111_v72_
                lhs11_ = default__.safeIndex(235, (d_111_v72_).length(0))
                d_109_v71_ = rhs17_
                lhs10_[lhs11_] = rhs18_
                d_109_v71_ = rhs19_
                d_115_v74_ = rhs20_
                d_116_v75_ = rhs21_
                d_117_v76_: _dafny.Set
                d_117_v76_ = _dafny.Set({d_31_v0_})
                index11_ = default__.safeIndex(979, (p0).length(0))
                (p0)[index11_] = (d_117_v76_).isdisjoint(d_117_v76_)
                d_118_v77_: _dafny.Array
                def lambda6_(d_119_p3_, d_120_p1_, d_121_v72_):
                    def lambda7_(d_122_i4_):
                        return (_dafny.Map({d_119_p3_: d_120_p1_})) | (_dafny.Map({(d_121_v72_)[default__.safeIndex(235, (d_121_v72_).length(0))]: d_119_p3_}))

                    return lambda7_

                init4_ = lambda6_(p3, p1, d_111_v72_)
                nw22_ = _dafny.Array(None, 2)
                for i0_4_ in range(nw22_.length(0)):
                    nw22_[i0_4_] = init4_(i0_4_)
                d_118_v77_ = nw22_
                d_123_v78_: _dafny.Map
                d_123_v78_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nufq"))): p1})
                index12_ = default__.safeIndex(296, (d_118_v77_).length(0))
                (d_118_v77_)[index12_] = d_123_v78_
                d_124_v79_: _dafny.MultiSet
                d_124_v79_ = _dafny.MultiSet([p1])
                index13_ = default__.safeIndex(979, (p0).length(0))
                index14_ = default__.safeIndex(296, (d_118_v77_).length(0))
                index15_ = default__.safeIndex(235, (d_111_v72_).length(0))
                rhs22_ = (p1) - (p3)
                rhs23_ = (d_124_v79_).issubset(d_124_v79_)
                rhs24_ = d_123_v78_
                rhs25_ = p1
                lhs12_ = globalState
                lhs13_ = p0
                lhs14_ = default__.safeIndex(979, (p0).length(0))
                lhs15_ = d_118_v77_
                lhs16_ = default__.safeIndex(296, (d_118_v77_).length(0))
                lhs17_ = d_111_v72_
                lhs18_ = default__.safeIndex(235, (d_111_v72_).length(0))
                lhs12_.f2 = rhs22_
                lhs13_[lhs14_] = rhs23_
                lhs15_[lhs16_] = rhs24_
                lhs17_[lhs18_] = rhs25_
                d_125_v80_: C1
                nw23_ = C1()
                nw23_.ctor__((D7_DC17(False)).cf25)
                d_125_v80_ = nw23_
                d_126_v81_: _dafny.Map
                d_126_v81_ = _dafny.Map({(p0)[default__.safeIndex(979, (p0).length(0))]: d_125_v80_})
                index16_ = default__.safeIndex(979, (p0).length(0))
                (p0)[index16_] = not(not((len((d_114_v73_) + (d_114_v73_))) > (len((d_126_v81_) | (d_126_v81_)))))
                (globalState).f2 = (0) - (((0) - (((d_111_v72_)[default__.safeIndex(235, (d_111_v72_).length(0))]) + (len(d_114_v73_))) if (p0)[default__.safeIndex(979, (p0).length(0))] else (0) - (len((d_117_v76_) | (d_117_v76_)))))
            elif True:
                d_127___mcc_h2_ = source1_.cf26
                d_128_cf26_ = d_127___mcc_h2_
                d_129_v82_: _dafny.MultiSet
                d_129_v82_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwvsosvy"))])
                d_130_v83_: D1
                d_130_v83_ = D1_DC4(d_129_v82_, p1)
                d_131_v84_: _dafny.Seq
                d_131_v84_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qcm"))
                pat_let_tv0_ = d_129_v82_
                pat_let_tv1_ = d_131_v84_
                pat_let_tv2_ = d_131_v84_
                def iife12_(_pat_let0_0):
                    def iife13_(d_132_dt__update__tmp_h0_):
                        def iife14_(_pat_let1_0):
                            def iife15_(d_133_dt__update_hcf6_h0_):
                                return D1_DC4(d_133_dt__update_hcf6_h0_, (d_132_dt__update__tmp_h0_).cf7)
                            return iife15_(_pat_let1_0)
                        return iife14_((pat_let_tv0_).intersection(_dafny.MultiSet([pat_let_tv1_, pat_let_tv2_])))
                    return iife13_(_pat_let0_0)
                d_130_v83_ = iife12_(d_130_v83_)
                (globalState).f2 = default__.safeDivisionInt(p1, p1)
                d_134_v85_: C3
                nw24_ = C3()
                nw24_.ctor__((default__.fm9(d_31_v0_, globalState)) <= (d_131_v84_), p1)
                d_134_v85_ = nw24_
                d_31_v0_ = d_31_v0_
            d_135_v86_: _dafny.MultiSet
            d_135_v86_ = _dafny.MultiSet([False])
            d_136_v87_: _dafny.MultiSet
            d_136_v87_ = _dafny.MultiSet([(d_135_v86_).cardinality, p3])
            d_137_v88_: _dafny.Seq
            d_137_v88_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpbqp"))
            d_138_v89_: _dafny.Map
            d_138_v89_ = _dafny.Map({d_137_v88_: default__.fm0(d_31_v0_, d_31_v0_, d_31_v0_, d_31_v0_, globalState)})
            d_139_v90_: _dafny.Set
            d_139_v90_ = _dafny.Set({(383) > (((d_136_v87_)[p3] if (p3) in (d_136_v87_) else p3)), d_31_v0_, ((d_138_v89_)[d_137_v88_] if (d_137_v88_) in (d_138_v89_) else d_31_v0_)})
            d_139_v90_ = (_dafny.Set({d_31_v0_, d_31_v0_, False})).intersection((d_139_v90_).intersection(d_139_v90_))
            (globalState).f2 = ((d_136_v87_)[530] if (530) in (d_136_v87_) else (_dafny.MultiSet([d_31_v0_, d_31_v0_])).cardinality)
            d_140_v91_: C1
            nw25_ = C1()
            nw25_.ctor__(d_31_v0_)
            d_140_v91_ = nw25_
            d_140_v91_ = d_140_v91_
            d_141_v92_: _dafny.MultiSet
            d_141_v92_ = _dafny.MultiSet([d_137_v88_, d_137_v88_])
            d_142_v93_: _dafny.Set
            d_142_v93_ = _dafny.Set({((d_141_v92_)[d_137_v88_] if (d_137_v88_) in (d_141_v92_) else p1), 311, p1, p1})
            d_142_v93_ = d_142_v93_
        d_143_v94_: D1
        d_143_v94_ = D1_DC5(d_31_v0_, default__.fm0(d_31_v0_, d_31_v0_, False, d_31_v0_, globalState))
        d_31_v0_ = not((d_143_v94_).cf9)
        d_144_v95_: _dafny.Seq
        d_144_v95_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfta"))
        d_145_v96_: _dafny.Map
        d_145_v96_ = _dafny.Map({d_31_v0_: d_144_v95_})
        d_146_v97_: _dafny.MultiSet
        d_146_v97_ = _dafny.MultiSet([p3])
        if (_dafny.MultiSet([len(d_145_v96_)])).ispropersubset(d_146_v97_):
            d_147_v98_: _dafny.Array
            nw26_ = _dafny.Array(False, 3)
            d_147_v98_ = nw26_
            d_147_v98_ = d_147_v98_
            d_148_v99_: C1
            nw27_ = C1()
            nw27_.ctor__(d_31_v0_)
            d_148_v99_ = nw27_
            d_149_v100_: C3
            nw28_ = C3()
            nw28_.ctor__(d_31_v0_, p1)
            d_149_v100_ = nw28_
            d_150_v101_: _dafny.Seq
            d_150_v101_ = _dafny.SeqWithoutIsStrInference([d_149_v100_])
            rhs26_ = (d_32_v1_) + (d_32_v1_)
            rhs27_ = d_148_v99_.f14
            rhs28_ = d_32_v1_
            rhs29_ = p3
            rhs30_ = (p3) + (len(d_150_v101_))
            lhs19_ = globalState
            lhs20_ = globalState
            d_32_v1_ = rhs26_
            d_31_v0_ = rhs27_
            d_32_v1_ = rhs28_
            lhs19_.f2 = rhs29_
            lhs20_.f2 = rhs30_
            d_151_v102_: str
            d_151_v102_ = _dafny.CodePoint('v')
            d_152_v103_: _dafny.Array
            nw29_ = _dafny.Array(None, 21)
            nw29_[int(0)] = d_151_v102_
            nw29_[int(1)] = d_151_v102_
            nw29_[int(2)] = d_151_v102_
            nw29_[int(3)] = d_151_v102_
            nw29_[int(4)] = (d_151_v102_ if default__.fm0(d_148_v99_.f14, (d_149_v100_).f11, d_148_v99_.f14, d_148_v99_.f14, globalState) else d_151_v102_)
            nw29_[int(5)] = d_151_v102_
            nw29_[int(6)] = d_151_v102_
            nw29_[int(7)] = (d_151_v102_ if d_31_v0_ else d_151_v102_)
            nw29_[int(8)] = d_151_v102_
            nw29_[int(9)] = d_151_v102_
            nw29_[int(10)] = d_151_v102_
            nw29_[int(11)] = d_151_v102_
            nw29_[int(12)] = (d_151_v102_ if d_31_v0_ else _dafny.CodePoint('x'))
            nw29_[int(13)] = d_151_v102_
            nw29_[int(14)] = d_151_v102_
            nw29_[int(15)] = d_151_v102_
            nw29_[int(16)] = d_151_v102_
            nw29_[int(17)] = d_151_v102_
            nw29_[int(18)] = d_151_v102_
            nw29_[int(19)] = d_151_v102_
            nw29_[int(20)] = _dafny.CodePoint('y')
            d_152_v103_ = nw29_
            nw30_ = _dafny.Array(_dafny.CodePoint('D'), 8)
            d_152_v103_ = nw30_
            (d_148_v99_).f14 = (p1) == (len(d_32_v1_))
        elif True:
            d_153_v104_: C0
            nw31_ = C0()
            nw31_.ctor__()
            d_153_v104_ = nw31_
            (globalState).f2 = p3
            d_154_v105_: _dafny.Array
            nw32_ = _dafny.Array(int(0), 18)
            d_154_v105_ = nw32_
            d_155_v106_: _dafny.Seq
            d_155_v106_ = _dafny.SeqWithoutIsStrInference([d_154_v105_, d_154_v105_])
            d_156_v107_: _dafny.Set
            d_156_v107_ = _dafny.Set({d_31_v0_})
            d_157_v108_: _dafny.Array
            nw33_ = _dafny.Array(None, 19)
            nw33_[int(0)] = d_154_v105_
            nw33_[int(1)] = d_154_v105_
            nw33_[int(2)] = d_154_v105_
            nw33_[int(3)] = d_154_v105_
            nw33_[int(4)] = d_154_v105_
            nw33_[int(5)] = d_154_v105_
            nw33_[int(6)] = (d_155_v106_)[default__.safeIndex(len(d_156_v107_), len(d_155_v106_))]
            nw33_[int(7)] = d_154_v105_
            nw33_[int(8)] = d_154_v105_
            nw33_[int(9)] = d_154_v105_
            nw33_[int(10)] = d_154_v105_
            nw33_[int(11)] = d_154_v105_
            nw33_[int(12)] = d_154_v105_
            nw33_[int(13)] = d_154_v105_
            nw33_[int(14)] = d_154_v105_
            nw33_[int(15)] = d_154_v105_
            nw33_[int(16)] = d_154_v105_
            nw33_[int(17)] = d_154_v105_
            nw33_[int(18)] = d_154_v105_
            d_157_v108_ = nw33_
            index17_ = default__.safeIndex(744, (d_157_v108_).length(0))
            (d_157_v108_)[index17_] = d_154_v105_
            d_158_v109_: _dafny.MultiSet
            d_158_v109_ = _dafny.MultiSet([d_144_v95_])
            d_159_v110_: D1
            d_159_v110_ = D1_DC6(D1_DC4(d_158_v109_, p1))
            d_160_v111_: D1
            d_160_v111_ = D1_DC3(_dafny.CodePoint('q'))
            d_161_v112_: str
            d_161_v112_ = _dafny.CodePoint('a')
            pat_let_tv3_ = d_160_v111_
            d_162_v113_: _dafny.Map
            def iife16_(_pat_let2_0):
                def iife17_(d_163_dt__update__tmp_h1_):
                    def iife18_(_pat_let3_0):
                        def iife19_(d_164_dt__update_hcf10_h0_):
                            return D1_DC6(d_164_dt__update_hcf10_h0_)
                        return iife19_(_pat_let3_0)
                    return iife18_(pat_let_tv3_)
                return iife17_(_pat_let2_0)
            d_162_v113_ = _dafny.Map({iife16_(d_159_v110_): d_161_v112_})
            index18_ = default__.safeIndex(744, (d_157_v108_).length(0))
            rhs31_ = d_154_v105_
            rhs32_ = d_154_v105_
            rhs33_ = ((default__.fm21(d_31_v0_, d_31_v0_, -627, d_31_v0_, globalState)) | (d_162_v113_)) | (_dafny.Map({d_159_v110_: _dafny.CodePoint('w')}))
            lhs21_ = globalState
            lhs22_ = d_157_v108_
            lhs23_ = default__.safeIndex(744, (d_157_v108_).length(0))
            lhs21_.f10 = rhs31_
            lhs22_[lhs23_] = rhs32_
            d_162_v113_ = rhs33_
            d_31_v0_ = not (d_31_v0_) or (d_31_v0_)
            d_31_v0_ = (d_31_v0_ if d_31_v0_ else not(d_31_v0_))
        if ((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_165_i5_ in range(default__.abs(789))])) + (d_144_v95_)) != (d_144_v95_):
            d_166_v114_: _dafny.Map
            d_166_v114_ = _dafny.Map({p1: p1})
            d_167_v115_: _dafny.Map
            d_167_v115_ = _dafny.Map({p3: d_31_v0_})
            d_31_v0_ = default__.fm0(d_31_v0_, (p3) >= ((0) - (p3)), (p1) != (len(d_166_v114_)), not(((d_167_v115_)[61] if (61) in (d_167_v115_) else d_31_v0_)), globalState)
            d_31_v0_ = d_31_v0_
            d_168_v116_: C2
            nw34_ = C2()
            nw34_.ctor__(p3)
            d_168_v116_ = nw34_
            d_169_v117_: D8
            d_169_v117_ = D8_DC19(d_168_v116_)
            d_169_v117_ = (d_169_v117_ if (p3) >= (d_168_v116_.f13) else d_169_v117_)
            d_170_v118_: _dafny.Array
            def lambda8_(d_171_v95_, d_172_v116_):
                def lambda9_(d_173_i6_):
                    return ((d_171_v95_).set(default__.safeIndex(d_172_v116_.f13, len(d_171_v95_)), _dafny.CodePoint('o'))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_174_i7_ in range(default__.abs(429))]))

                return lambda9_

            init5_ = lambda8_(d_144_v95_, d_168_v116_)
            nw35_ = _dafny.Array(None, 23)
            for i0_5_ in range(nw35_.length(0)):
                nw35_[i0_5_] = init5_(i0_5_)
            d_170_v118_ = nw35_
            index19_ = default__.safeIndex(301, (d_170_v118_).length(0))
            (d_170_v118_)[index19_] = (d_144_v95_) + (default__.fm9(d_31_v0_, globalState))
            index20_ = default__.safeIndex(301, (d_170_v118_).length(0))
            (d_170_v118_)[index20_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_175_i8_ in range(default__.abs(-589))])
            d_176_v119_: _dafny.Set
            d_176_v119_ = _dafny.Set({d_168_v116_.f13})
            d_177_v120_: _dafny.Seq
            d_177_v120_ = _dafny.SeqWithoutIsStrInference([d_176_v119_])
            d_31_v0_ = (d_176_v119_).ispropersubset((d_177_v120_)[default__.safeIndex((0) - (d_168_v116_.f13), len(d_177_v120_))])
        elif True:
            d_178_v121_: _dafny.Array
            nw36_ = _dafny.Array(_dafny.Map({}), 24)
            d_178_v121_ = nw36_
            d_179_v122_: _dafny.Map
            d_179_v122_ = _dafny.Map({d_31_v0_: p1})
            index21_ = default__.safeIndex(978, (d_178_v121_).length(0))
            (d_178_v121_)[index21_] = (_dafny.Map({d_31_v0_: len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_180_i9_ in range(default__.abs(107))]))})) | (d_179_v122_)
            d_181_v123_: _dafny.Seq
            d_181_v123_ = _dafny.SeqWithoutIsStrInference([p1])
            d_182_v124_: _dafny.Set
            d_182_v124_ = _dafny.Set({d_31_v0_})
            d_183_v125_: _dafny.Seq
            d_183_v125_ = _dafny.SeqWithoutIsStrInference([d_182_v124_, d_182_v124_])
            d_184_v127_: _dafny.Array
            nw37_ = _dafny.Array(None, 17)
            nw37_[int(0)] = p3
            nw37_[int(1)] = p3
            nw37_[int(2)] = p3
            nw37_[int(3)] = p3
            nw37_[int(4)] = len(d_181_v123_)
            nw37_[int(5)] = len(d_182_v124_)
            nw37_[int(6)] = p3
            nw37_[int(7)] = (d_146_v97_).cardinality
            def iife20_():
                coll12_ = _dafny.Map()
                compr_12_: int
                for compr_12_ in _dafny.IntegerRange(175, 637):
                    d_185_v126_: int = compr_12_
                    if ((175) <= (d_185_v126_)) and ((d_185_v126_) < (637)):
                        coll12_[(d_185_v126_) - (p1)] = d_31_v0_
                return _dafny.Map(coll12_)
            nw37_[int(8)] = len(iife20_()
            )
            nw37_[int(9)] = p3
            nw37_[int(10)] = (d_146_v97_).cardinality
            nw37_[int(11)] = p1
            nw37_[int(12)] = (0) - (p1)
            nw37_[int(13)] = p3
            nw37_[int(14)] = p3
            nw37_[int(15)] = p1
            nw37_[int(16)] = p3
            d_184_v127_ = nw37_
            d_186_v128_: _dafny.Seq
            d_186_v128_ = _dafny.SeqWithoutIsStrInference([d_184_v127_, d_184_v127_])
            d_187_v129_: D11
            d_187_v129_ = D11_DC26(p3, (d_186_v128_)[default__.safeIndex(p1, len(d_186_v128_))], d_31_v0_)
            d_188_v130_: _dafny.Array
            nw38_ = _dafny.Array(None, 15)
            nw38_[int(0)] = default__.fm15(p3, len(d_181_v123_), default__.fm0(default__.fm0(d_31_v0_, d_31_v0_, d_31_v0_, not(d_31_v0_), globalState), d_31_v0_, d_31_v0_, not(d_31_v0_), globalState), globalState)
            nw38_[int(1)] = d_182_v124_
            nw38_[int(2)] = (d_183_v125_)[default__.safeIndex((d_181_v123_)[default__.safeIndex(p3, len(d_181_v123_))], len(d_183_v125_))]
            nw38_[int(3)] = d_182_v124_
            nw38_[int(4)] = (d_182_v124_).intersection(d_182_v124_)
            nw38_[int(5)] = d_182_v124_
            nw38_[int(6)] = (d_182_v124_ if d_31_v0_ else default__.fm15(p1, p1, d_31_v0_, globalState))
            nw38_[int(7)] = d_182_v124_
            nw38_[int(8)] = (d_182_v124_) | (_dafny.Set({not(d_31_v0_), (d_187_v129_).cf38}))
            nw38_[int(9)] = d_182_v124_
            nw38_[int(10)] = default__.fm15(p1, p1, d_31_v0_, globalState)
            nw38_[int(11)] = d_182_v124_
            nw38_[int(12)] = _dafny.Set({False})
            nw38_[int(13)] = d_182_v124_
            nw38_[int(14)] = d_182_v124_
            d_188_v130_ = nw38_
            index22_ = default__.safeIndex(387, (d_188_v130_).length(0))
            (d_188_v130_)[index22_] = d_182_v124_
            d_189_v131_: _dafny.Set
            d_189_v131_ = _dafny.Set({p1, p3})
            index23_ = default__.safeIndex(978, (d_178_v121_).length(0))
            index24_ = default__.safeIndex(387, (d_188_v130_).length(0))
            rhs34_ = d_179_v122_
            rhs35_ = _dafny.Set({(d_189_v131_).isdisjoint(d_189_v131_)})
            lhs24_ = d_178_v121_
            lhs25_ = default__.safeIndex(978, (d_178_v121_).length(0))
            lhs26_ = d_188_v130_
            lhs27_ = default__.safeIndex(387, (d_188_v130_).length(0))
            lhs24_[lhs25_] = rhs34_
            lhs26_[lhs27_] = rhs35_
            d_190_v132_: D0
            d_190_v132_ = D0_DC0(d_31_v0_)
            d_31_v0_ = default__.fm0(not ((d_190_v132_).cf0) or (d_31_v0_), default__.fm0(d_31_v0_, d_31_v0_, d_31_v0_, d_31_v0_, globalState), ((0) - (p3)) < (p3), default__.fm0(d_31_v0_, default__.fm0(d_31_v0_, d_31_v0_, d_31_v0_, d_31_v0_, globalState), d_31_v0_, d_31_v0_, globalState), globalState)
            d_31_v0_ = not(True)
            d_191_v133_: _dafny.Array
            nw39_ = _dafny.Array(_dafny.Seq({}), 24)
            d_191_v133_ = nw39_
            index25_ = default__.safeIndex(765, (d_191_v133_).length(0))
            (d_191_v133_)[index25_] = d_181_v123_
            index26_ = default__.safeIndex(765, (d_191_v133_).length(0))
            (d_191_v133_)[index26_] = d_181_v123_
            index27_ = default__.safeIndex(45, (d_184_v127_).length(0))
            (d_184_v127_)[index27_] = p3
            index28_ = default__.safeIndex(45, (d_184_v127_).length(0))
            (d_184_v127_)[index28_] = (0) - (p1)

    @staticmethod
    def Main(noArgsParameter__):
        d_192_v0_: bool
        d_192_v0_ = True
        d_193_v1_: _dafny.Set
        d_193_v1_ = _dafny.Set({not(not(d_192_v0_))})
        d_194_v2_: _dafny.Set
        d_194_v2_ = _dafny.Set({_dafny.Set({d_192_v0_, not(not(d_192_v0_))}), d_193_v1_})
        d_195_v3_: int
        d_195_v3_ = 848
        d_196_v4_: _dafny.Seq
        d_196_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk"))
        d_197_v5_: _dafny.Map
        d_197_v5_ = _dafny.Map({d_192_v0_: d_196_v4_})
        d_198_v6_: _dafny.Seq
        d_198_v6_ = _dafny.SeqWithoutIsStrInference([len(d_197_v5_), len(_dafny.Map({d_195_v3_: d_195_v3_})), d_195_v3_])
        d_199_v7_: _dafny.Seq
        d_199_v7_ = _dafny.SeqWithoutIsStrInference([d_193_v1_, d_193_v1_, d_193_v1_])
        d_200_v8_: _dafny.Array
        nw40_ = _dafny.Array(int(0), 14)
        d_200_v8_ = nw40_
        d_201_v9_: _dafny.Map
        d_201_v9_ = _dafny.Map({len((d_199_v7_)[default__.safeIndex(d_195_v3_, len(d_199_v7_))]): d_200_v8_})
        d_202_v10_: _dafny.MultiSet
        d_202_v10_ = _dafny.MultiSet([not(d_192_v0_)])
        d_203_globalState_: GlobalState
        nw41_ = GlobalState()
        nw41_.ctor__(245, False, 862, False, (_dafny.Set({d_193_v1_, d_193_v1_})) - (d_194_v2_), (_dafny.Map({d_195_v3_: d_192_v0_}) if d_192_v0_ else _dafny.Map({d_195_v3_: not(d_192_v0_)})), d_198_v6_, d_198_v6_, False, True, ((d_201_v9_)[((d_202_v10_)[d_192_v0_] if (d_192_v0_) in (d_202_v10_) else d_195_v3_)] if (((d_202_v10_)[d_192_v0_] if (d_192_v0_) in (d_202_v10_) else d_195_v3_)) in (d_201_v9_) else d_200_v8_))
        d_203_globalState_ = nw41_
        (d_203_globalState_).f2 = (default__.safeModuloInt(-661, (0) - (d_195_v3_))) * (d_195_v3_)
        d_204_v11_: _dafny.Map
        d_204_v11_ = _dafny.Map({d_192_v0_: (758) >= (d_195_v3_)})
        d_205_v12_: _dafny.Map
        d_205_v12_ = _dafny.Map({d_192_v0_: d_195_v3_})
        d_204_v11_ = (d_204_v11_).set((d_195_v3_) <= (d_195_v3_), (d_195_v3_) == (len(d_205_v12_)))
        d_206_v13_: _dafny.Array
        def lambda10_(d_207_v0_):
            def lambda11_(d_208_i0_):
                return not(d_207_v0_)

            return lambda11_

        init6_ = lambda10_(d_192_v0_)
        nw42_ = _dafny.Array(None, 15)
        for i0_6_ in range(nw42_.length(0)):
            nw42_[i0_6_] = init6_(i0_6_)
        d_206_v13_ = nw42_
        d_209_v14_: _dafny.Array
        def lambda12_(d_210_v11_):
            def lambda13_(d_211_i1_):
                return d_210_v11_

            return lambda13_

        init7_ = lambda12_(d_204_v11_)
        nw43_ = _dafny.Array(None, 14)
        for i0_7_ in range(nw43_.length(0)):
            nw43_[i0_7_] = init7_(i0_7_)
        d_209_v14_ = nw43_
        default__.m0(d_206_v13_, d_195_v3_, d_209_v14_, d_195_v3_, d_203_globalState_)
        if not(d_192_v0_):
            d_212_v15_: D0
            d_212_v15_ = D0_DC0(True)
            d_213_v16_: _dafny.Seq
            d_213_v16_ = _dafny.SeqWithoutIsStrInference([default__.fm0(d_192_v0_, d_192_v0_, d_192_v0_, (d_212_v15_).cf0, d_203_globalState_)])
            default__.m0(d_206_v13_, (len(d_213_v16_)) * (d_195_v3_), d_209_v14_, len(d_193_v1_), d_203_globalState_)
            d_192_v0_ = d_192_v0_
            d_214_v17_: C2
            nw44_ = C2()
            nw44_.ctor__(d_195_v3_)
            d_214_v17_ = nw44_
            d_192_v0_ = d_192_v0_
            d_192_v0_ = d_192_v0_
        elif True:
            (d_203_globalState_).f2 = d_195_v3_
            d_192_v0_ = d_192_v0_
            d_215_v18_: str
            d_215_v18_ = _dafny.CodePoint('t')
            d_216_v19_: _dafny.Seq
            d_216_v19_ = _dafny.SeqWithoutIsStrInference([d_192_v0_, d_192_v0_])
            source2_ = default__.fm13(d_215_v18_, _dafny.MultiSet(d_216_v19_), d_203_globalState_)
            if source2_.is_DC4:
                d_217___mcc_h0_ = source2_.cf6
                d_218___mcc_h1_ = source2_.cf7
                d_219_cf7_ = d_218___mcc_h1_
                d_220_cf6_ = d_217___mcc_h0_
                rhs36_ = d_192_v0_
                rhs37_ = default__.safeDivisionInt(d_195_v3_, d_219_cf7_)
                rhs38_ = d_206_v13_
                rhs39_ = (0) - (d_219_cf7_)
                lhs28_ = d_203_globalState_
                d_192_v0_ = rhs36_
                lhs28_.f2 = rhs37_
                d_206_v13_ = rhs38_
                d_195_v3_ = rhs39_
                (d_203_globalState_).f2 = d_219_cf7_
                d_221_v20_: C1
                nw45_ = C1()
                nw45_.ctor__(True)
                d_221_v20_ = nw45_
                d_221_v20_ = d_221_v20_
                d_222_v21_: _dafny.Seq
                d_223_v22_: D1
                d_224_v23_: int
                out0_: _dafny.Seq
                out1_: D1
                out2_: int
                out0_, out1_, out2_ = (d_221_v20_).m4(d_205_v12_, d_203_globalState_)
                d_222_v21_ = out0_
                d_223_v22_ = out1_
                d_224_v23_ = out2_
            elif source2_.is_DC5:
                d_225___mcc_h2_ = source2_.cf8
                d_226___mcc_h3_ = source2_.cf9
                d_227_cf9_ = d_226___mcc_h3_
                d_228_cf8_ = d_225___mcc_h2_
                d_229_v24_: D4
                d_229_v24_ = D4_DC13(d_215_v18_)
                d_230_v25_: _dafny.Map
                d_230_v25_ = _dafny.Map({d_227_cf9_: d_229_v24_})
                d_231_v26_: _dafny.Set
                d_231_v26_ = _dafny.Set({d_195_v3_, d_195_v3_})
                d_232_v27_: D3
                d_232_v27_ = D3_DC10(d_192_v0_, d_195_v3_, d_227_cf9_, d_206_v13_, _dafny.Map({len(d_231_v26_): d_227_cf9_}))
                d_230_v25_ = (d_230_v25_).set((d_232_v27_).cf16, d_229_v24_)
                d_204_v11_ = (d_204_v11_).set(d_192_v0_, d_192_v0_)
                d_227_cf9_ = d_227_cf9_
                d_233_v28_: C3
                nw46_ = C3()
                nw46_.ctor__(d_227_cf9_, ((d_198_v6_)[default__.safeIndex(d_195_v3_, len(d_198_v6_))]) * (len(d_216_v19_)))
                d_233_v28_ = nw46_
            elif source2_.is_DC3:
                d_234___mcc_h4_ = source2_.cf5
                d_235_cf5_ = d_234___mcc_h4_
                (d_203_globalState_).f2 = d_195_v3_
                d_192_v0_ = d_192_v0_
                (d_203_globalState_).f2 = default__.safeDivisionInt(d_195_v3_, d_195_v3_)
                d_195_v3_ = default__.fm2(not(d_192_v0_), d_203_globalState_)
            elif True:
                d_236___mcc_h5_ = source2_.cf10
                d_237_cf10_ = d_236___mcc_h5_
                d_238_v29_: D1
                d_238_v29_ = D1_DC3(d_215_v18_)
                d_239_v30_: _dafny.Seq
                d_239_v30_ = _dafny.SeqWithoutIsStrInference([default__.fm5(454, 47, 921, d_216_v19_, d_203_globalState_), d_238_v29_])
                d_240_v31_: _dafny.Map
                d_240_v31_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mawdmqm"))) + (d_196_v4_): len(d_239_v30_)})
                d_240_v31_ = ((d_240_v31_ if True else d_240_v31_)) | (d_240_v31_)
                default__.m0(d_206_v13_, ((0) - ((_dafny.MultiSet(d_216_v19_)).cardinality)) * (d_195_v3_), d_209_v14_, d_195_v3_, d_203_globalState_)
                d_241_v32_: _dafny.Set
                d_241_v32_ = _dafny.Set({len(d_216_v19_), d_195_v3_})
                index29_ = default__.safeIndex(188, (d_200_v8_).length(0))
                (d_200_v8_)[index29_] = len(d_241_v32_)
                index30_ = default__.safeIndex(188, (d_200_v8_).length(0))
                (d_200_v8_)[index30_] = d_195_v3_
                (d_203_globalState_).f2 = d_195_v3_
            (d_203_globalState_).f2 = d_195_v3_
            d_202_v10_ = d_202_v10_
        hi0_ = d_195_v3_
        for d_242_i2_ in range((d_195_v3_) * (d_195_v3_), hi0_):
            d_243_v33_: C2
            nw47_ = C2()
            nw47_.ctor__(d_195_v3_)
            d_243_v33_ = nw47_
            d_244_v34_: C1
            nw48_ = C1()
            nw48_.ctor__(d_192_v0_)
            d_244_v34_ = nw48_
            d_245_v35_: _dafny.Map
            d_245_v35_ = _dafny.Map({d_195_v3_: d_244_v34_})
            d_246_v36_: C0
            nw49_ = C0()
            nw49_.ctor__()
            d_246_v36_ = nw49_
            d_247_v37_: _dafny.Map
            d_247_v37_ = _dafny.Map({d_245_v35_: d_246_v36_})
            d_248_v38_: _dafny.Map
            d_248_v38_ = _dafny.Map({((d_247_v37_)[d_245_v35_] if (d_245_v35_) in (d_247_v37_) else d_246_v36_): d_243_v33_})
            d_249_v39_: _dafny.Seq
            d_249_v39_ = _dafny.SeqWithoutIsStrInference([not(d_244_v34_.f14)])
            d_250_v40_: _dafny.Array
            nw50_ = _dafny.Array(None, 25)
            nw50_[int(0)] = d_243_v33_
            nw50_[int(1)] = d_243_v33_
            nw50_[int(2)] = d_243_v33_
            nw50_[int(3)] = d_243_v33_
            nw50_[int(4)] = d_243_v33_
            nw50_[int(5)] = d_243_v33_
            nw50_[int(6)] = d_243_v33_
            nw50_[int(7)] = d_243_v33_
            nw50_[int(8)] = d_243_v33_
            nw50_[int(9)] = d_243_v33_
            nw50_[int(10)] = d_243_v33_
            nw50_[int(11)] = d_243_v33_
            nw50_[int(12)] = d_243_v33_
            nw50_[int(13)] = (d_243_v33_ if d_192_v0_ else d_243_v33_)
            nw50_[int(14)] = d_243_v33_
            nw50_[int(15)] = d_243_v33_
            nw50_[int(16)] = d_243_v33_
            nw50_[int(17)] = d_243_v33_
            nw50_[int(18)] = d_243_v33_
            nw50_[int(19)] = d_243_v33_
            nw50_[int(20)] = ((d_248_v38_)[d_246_v36_] if (d_246_v36_) in (d_248_v38_) else d_243_v33_)
            nw50_[int(21)] = (d_243_v33_ if (d_249_v39_)[default__.safeIndex(d_243_v33_.f13, len(d_249_v39_))] else d_243_v33_)
            nw50_[int(22)] = d_243_v33_
            nw50_[int(23)] = d_243_v33_
            nw50_[int(24)] = d_243_v33_
            d_250_v40_ = nw50_
            index31_ = default__.safeIndex(930, (d_250_v40_).length(0))
            (d_250_v40_)[index31_] = d_243_v33_
            index32_ = default__.safeIndex(930, (d_250_v40_).length(0))
            rhs40_ = not(((d_195_v3_) * (-807)) <= (d_243_v33_.f13))
            rhs41_ = d_243_v33_
            lhs29_ = d_250_v40_
            lhs30_ = default__.safeIndex(930, (d_250_v40_).length(0))
            d_192_v0_ = rhs40_
            lhs29_[lhs30_] = rhs41_
            d_196_v4_ = (d_196_v4_) + (d_196_v4_)
            d_251_v42_: _dafny.MultiSet
            d_251_v42_ = _dafny.MultiSet([d_242_i2_, d_243_v33_.f13])
            d_252_v43_: _dafny.Set
            d_252_v43_ = _dafny.Set({(d_251_v42_).cardinality})
            def iife21_():
                coll13_ = _dafny.Map()
                compr_13_: int
                for compr_13_ in (d_252_v43_).Elements:
                    d_253_v41_: int = compr_13_
                    if (d_253_v41_) in (d_252_v43_):
                        coll13_[default__.safeModuloInt(d_253_v41_, d_195_v3_)] = d_243_v33_.f13
                return _dafny.Map(coll13_)
            (d_244_v34_).f14 = (d_195_v3_) not in ((iife21_()
            ).set((0) - (d_243_v33_.f13), len(d_252_v43_)))
            d_254_v44_: _dafny.MultiSet
            d_254_v44_ = _dafny.MultiSet([d_200_v8_, d_200_v8_, d_200_v8_, d_200_v8_])
            if (d_200_v8_) in (d_254_v44_):
                (d_203_globalState_).f2 = (0) - (((d_198_v6_)[default__.safeIndex(d_243_v33_.f13, len(d_198_v6_))]) * (d_242_i2_))
                d_196_v4_ = ((d_196_v4_) + (d_196_v4_)) + (d_196_v4_)
                d_255_v45_: _dafny.MultiSet
                d_255_v45_ = _dafny.MultiSet([(d_196_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ux"))), d_196_v4_, (d_196_v4_) + (d_196_v4_), d_196_v4_, d_196_v4_])
                d_255_v45_ = d_255_v45_
                d_256_v46_: _dafny.Array
                def lambda14_(d_257_i3_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))

                init8_ = lambda14_
                nw51_ = _dafny.Array(None, 15)
                for i0_8_ in range(nw51_.length(0)):
                    nw51_[i0_8_] = init8_(i0_8_)
                d_256_v46_ = nw51_
                d_258_v47_: str
                d_258_v47_ = _dafny.CodePoint('d')
                index33_ = default__.safeIndex(63, (d_256_v46_).length(0))
                (d_256_v46_)[index33_] = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfldjy"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")))).set(default__.safeIndex(d_243_v33_.f13, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfldjy"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))))), d_258_v47_)
                d_259_v48_: D2
                d_259_v48_ = D2_DC8(len(d_249_v39_))
                index34_ = default__.safeIndex(63, (d_256_v46_).length(0))
                rhs42_ = (d_243_v33_.f13) * ((d_259_v48_).cf12)
                rhs43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqdmq"))
                rhs44_ = d_243_v33_.f13
                rhs45_ = (d_243_v33_.f13) + (d_243_v33_.f13)
                lhs31_ = d_243_v33_
                lhs32_ = d_256_v46_
                lhs33_ = default__.safeIndex(63, (d_256_v46_).length(0))
                lhs34_ = d_203_globalState_
                lhs35_ = d_203_globalState_
                lhs31_.f13 = rhs42_
                lhs32_[lhs33_] = rhs43_
                lhs34_.f2 = rhs44_
                lhs35_.f2 = rhs45_
                d_260_v49_: _dafny.Map
                d_260_v49_ = _dafny.Map({d_205_v12_: _dafny.Map({d_195_v3_: d_244_v34_})})
                d_260_v49_ = d_260_v49_
            elif True:
                d_204_v11_ = d_204_v11_
                d_261_v50_: str
                d_261_v50_ = _dafny.CodePoint('h')
                d_262_v51_: D1
                d_262_v51_ = D1_DC3(d_261_v50_)
                d_263_v52_: C3
                nw52_ = C3()
                nw52_.ctor__(not((d_246_v36_).fm7(False, d_262_v51_, d_196_v4_, d_244_v34_.f14, d_203_globalState_)), d_195_v3_)
                d_263_v52_ = nw52_
                d_263_v52_ = d_263_v52_
                d_264_v53_: _dafny.Array
                nw53_ = _dafny.Array(D4.default()(), 1)
                d_264_v53_ = nw53_
                d_265_v54_: D4
                d_265_v54_ = D4_DC12(d_193_v1_)
                index35_ = default__.safeIndex(230, (d_264_v53_).length(0))
                (d_264_v53_)[index35_] = d_265_v54_
                index36_ = default__.safeIndex(230, (d_264_v53_).length(0))
                (d_264_v53_)[index36_] = d_265_v54_
                d_266_v55_: D1
                d_266_v55_ = D1_DC5(d_192_v0_, d_244_v34_.f14)
                pat_let_tv4_ = d_244_v34_
                def iife22_(_pat_let4_0):
                    def iife23_(d_267_dt__update__tmp_h0_):
                        def iife24_(_pat_let5_0):
                            def iife25_(d_268_dt__update_hcf9_h0_):
                                return D1_DC5((d_267_dt__update__tmp_h0_).cf8, d_268_dt__update_hcf9_h0_)
                            return iife25_(_pat_let5_0)
                        return iife24_(pat_let_tv4_.f14)
                    return iife23_(_pat_let4_0)
                d_192_v0_ = (iife22_(d_266_v55_)).cf9
                (d_244_v34_).f14 = (d_244_v34_.f14) == (d_244_v34_.f14)
        d_269_v56_: _dafny.Map
        d_269_v56_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sharml")): d_195_v3_})
        d_270_v57_: str
        d_270_v57_ = _dafny.CodePoint('c')
        d_271_v58_: C3
        nw54_ = C3()
        nw54_.ctor__(d_192_v0_, d_195_v3_)
        d_271_v58_ = nw54_
        d_272_v59_: _dafny.Map
        d_272_v59_ = _dafny.Map({d_270_v57_: d_271_v58_})
        d_273_v60_: _dafny.Map
        d_273_v60_ = _dafny.Map({default__.fm4(d_203_globalState_): ((d_272_v59_)[d_270_v57_] if (d_270_v57_) in (d_272_v59_) else d_271_v58_)})
        d_274_v61_: _dafny.Map
        d_274_v61_ = _dafny.Map({d_273_v60_: d_200_v8_})
        d_275_v62_: _dafny.Seq
        d_275_v62_ = _dafny.SeqWithoutIsStrInference([(len(d_269_v56_)) > (len(d_274_v61_))])
        d_275_v62_ = d_275_v62_
        d_276_v63_: bool
        d_277_v64_: _dafny.Map
        d_278_v65_: bool
        out3_: bool
        out4_: _dafny.Map
        out5_: bool
        out3_, out4_, out5_ = (d_271_v58_).m2(d_192_v0_, (d_271_v58_).f11, d_203_globalState_)
        d_276_v63_ = out3_
        d_277_v64_ = out4_
        d_278_v65_ = out5_
        d_279_v66_: _dafny.Map
        d_279_v66_ = _dafny.Map({D0_DC1((d_271_v58_).f12, (d_271_v58_).f12, (d_271_v58_).f12): (d_271_v58_).f11})
        d_280_v67_: _dafny.Map
        d_280_v67_ = _dafny.Map({(d_271_v58_).f12: d_195_v3_})
        d_281_v68_: _dafny.Map
        d_281_v68_ = _dafny.Map({d_192_v0_: d_280_v67_})
        d_282_v69_: _dafny.Map
        d_282_v69_ = d_201_v9_
        d_283_v70_: _dafny.MultiSet
        d_283_v70_ = _dafny.MultiSet([_dafny.Map({d_282_v69_: (d_271_v58_).f12})])
        index37_ = default__.safeIndex(595, (d_200_v8_).length(0))
        (d_200_v8_)[index37_] = (d_283_v70_).cardinality
        index38_ = default__.safeIndex(595, (d_200_v8_).length(0))
        rhs46_ = (0) - (default__.safeModuloInt((d_271_v58_).f12, len((d_204_v11_) | (d_204_v11_))))
        rhs47_ = d_279_v66_
        rhs48_ = default__.fm14((d_196_v4_) + (d_196_v4_), d_203_globalState_)
        rhs49_ = (len(d_196_v4_)) + ((d_271_v58_).f12)
        lhs36_ = d_200_v8_
        lhs37_ = default__.safeIndex(595, (d_200_v8_).length(0))
        d_195_v3_ = rhs46_
        d_279_v66_ = rhs47_
        d_281_v68_ = rhs48_
        lhs36_[lhs37_] = rhs49_
        d_284_i4_: int
        d_284_i4_ = 0
        with _dafny.label("0"):
            while ((d_271_v58_).f12) >= ((d_202_v10_).cardinality):
                with _dafny.c_label("0"):
                    if (d_284_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_284_i4_ = (d_284_i4_) + (1)
                    d_285_v71_: _dafny.Map
                    d_285_v71_ = _dafny.Map({default__.safeModuloInt((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))], (d_271_v58_).f12): d_196_v4_})
                    d_286_v72_: _dafny.Seq
                    d_286_v72_ = _dafny.SeqWithoutIsStrInference([d_271_v58_])
                    d_287_v73_: _dafny.Seq
                    d_287_v73_ = d_286_v72_
                    d_285_v71_ = (d_285_v71_).set(len((d_287_v73_)), _dafny.SeqWithoutIsStrInference([default__.fm6((d_271_v58_).f12, d_192_v0_, d_278_v65_, _dafny.CodePoint('n'), d_203_globalState_)]))
                    d_192_v0_ = ((d_202_v10_).intersection(_dafny.MultiSet([True]))) == (d_202_v10_)
                    d_288_v74_: D4
                    d_288_v74_ = D4_DC13((d_270_v57_ if (d_271_v58_).f11 else d_270_v57_))
                    d_288_v74_ = d_288_v74_
                    d_289_v75_: D1
                    d_289_v75_ = D1_DC5((d_271_v58_).f11, d_192_v0_)
                    d_290_v76_: _dafny.Seq
                    d_290_v76_ = _dafny.SeqWithoutIsStrInference([d_289_v75_])
                    d_290_v76_ = d_290_v76_
                    pass
            pass
        d_291_v77_: C3
        nw55_ = C3()
        nw55_.ctor__(not((d_271_v58_).f11), 651)
        d_291_v77_ = nw55_
        d_292_v78_: _dafny.Map
        d_292_v78_ = _dafny.Map({d_291_v77_: (d_291_v77_).f12})
        d_292_v78_ = (d_292_v78_).set(d_291_v77_, (d_291_v77_).f12)
        d_195_v3_ = (d_271_v58_).f12
        (d_291_v77_).m1(not(d_276_v63_), d_203_globalState_)
        d_293_v79_: _dafny.Array
        nw56_ = _dafny.Array(_dafny.CodePoint('D'), 20)
        d_293_v79_ = nw56_
        d_294_v80_: _dafny.MultiSet
        d_294_v80_ = _dafny.MultiSet([d_293_v79_, d_293_v79_, d_293_v79_])
        d_294_v80_ = d_294_v80_
        source3_ = default__.fm13(d_270_v57_, _dafny.MultiSet([(d_271_v58_).f11, d_278_v65_, not((d_291_v77_).f11)]), d_203_globalState_)
        if source3_.is_DC4:
            d_295___mcc_h6_ = source3_.cf6
            d_296___mcc_h7_ = source3_.cf7
            d_297_cf7_ = d_296___mcc_h7_
            d_298_cf6_ = d_295___mcc_h6_
            d_299_v81_: D1
            d_299_v81_ = D1_DC3(d_270_v57_)
            source4_ = d_299_v81_
            if source4_.is_DC4:
                d_300___mcc_h12_ = source4_.cf6
                d_301___mcc_h13_ = source4_.cf7
                d_302_cf7_ = d_301___mcc_h13_
                d_303_cf6_ = d_300___mcc_h12_
                d_304_v82_: _dafny.Set
                d_304_v82_ = _dafny.Set({832, (d_291_v77_).f12, (d_291_v77_).f12})
                d_305_v83_: C3
                nw57_ = C3()
                nw57_.ctor__(not(d_192_v0_), len(d_304_v82_))
                d_305_v83_ = nw57_
                d_306_v84_: _dafny.Map
                d_306_v84_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nuwybql"))): ((d_305_v83_).f11) or (default__.fm0(False, (d_271_v58_).f11, d_192_v0_, (d_291_v77_).f11, d_203_globalState_))})
                d_307_v85_: _dafny.MultiSet
                d_307_v85_ = _dafny.MultiSet([d_204_v11_])
                d_306_v84_ = (d_306_v84_).set((0) - (((d_307_v85_) - (_dafny.MultiSet([d_204_v11_, d_204_v11_, d_204_v11_]))).cardinality), (d_291_v77_).f11)
                index39_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index39_] = (d_271_v58_).f12
                d_276_v63_ = d_278_v65_
            elif source4_.is_DC5:
                d_308___mcc_h14_ = source4_.cf8
                d_309___mcc_h15_ = source4_.cf9
                d_310_cf9_ = d_309___mcc_h15_
                d_311_cf8_ = d_308___mcc_h14_
                d_312_v86_: _dafny.MultiSet
                d_312_v86_ = _dafny.MultiSet([(d_271_v58_).fm1(d_196_v4_, (d_271_v58_).f12, d_203_globalState_), (d_291_v77_).f12, (d_291_v77_).f12, len(d_196_v4_), len(d_197_v5_)])
                default__.m0(d_206_v13_, (((d_312_v86_)[-408] if (-408) in (d_312_v86_) else (d_291_v77_).f12)) + ((d_291_v77_).f12), d_209_v14_, d_195_v3_, d_203_globalState_)
                d_196_v4_ = (d_196_v4_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_313_i5_ in range(default__.abs(889))]))
                d_314_v87_: bool
                d_315_v88_: _dafny.Map
                d_316_v89_: bool
                out6_: bool
                out7_: _dafny.Map
                out8_: bool
                out6_, out7_, out8_ = (d_271_v58_).m2((d_195_v3_) != (d_297_cf7_), d_276_v63_, d_203_globalState_)
                d_314_v87_ = out6_
                d_315_v88_ = out7_
                d_316_v89_ = out8_
                index40_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index40_] = (0) - ((d_291_v77_).f12)
            elif source4_.is_DC3:
                d_317___mcc_h16_ = source4_.cf5
                d_318_cf5_ = d_317___mcc_h16_
                d_276_v63_ = not(not(not (((d_271_v58_).f11 if False else (d_271_v58_).f11)) or (not((d_271_v58_).f11))))
                def iife26_():
                    coll14_ = _dafny.Set()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(393, 611):
                        d_319_v90_: int = compr_14_
                        if ((393) <= (d_319_v90_)) and ((d_319_v90_) < (611)):
                            coll14_ = coll14_.union(_dafny.Set([default__.safeDivisionInt(d_319_v90_, (d_271_v58_).f12)]))
                    return _dafny.Set(coll14_)
                d_297_cf7_ = (d_297_cf7_) - ((10) - (len(iife26_()
                )))
                d_204_v11_ = (d_204_v11_).set((d_271_v58_).f11, d_276_v63_)
                d_320_v91_: C0
                nw58_ = C0()
                nw58_.ctor__()
                d_320_v91_ = nw58_
            elif True:
                d_321___mcc_h17_ = source4_.cf10
                d_322_cf10_ = d_321___mcc_h17_
                d_323_v92_: C1
                nw59_ = C1()
                nw59_.ctor__((d_291_v77_).f11)
                d_323_v92_ = nw59_
                d_324_v93_: _dafny.Set
                d_324_v93_ = _dafny.Set({d_323_v92_, d_323_v92_, d_323_v92_})
                (d_271_v58_).m1((d_324_v93_).ispropersubset(d_324_v93_), d_203_globalState_)
                d_325_v94_: _dafny.Array
                nw60_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 1)
                d_325_v94_ = nw60_
                index41_ = default__.safeIndex(269, (d_325_v94_).length(0))
                (d_325_v94_)[index41_] = d_196_v4_
                index42_ = default__.safeIndex(269, (d_325_v94_).length(0))
                (d_325_v94_)[index42_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "adpwq"))
                d_326_v95_: _dafny.MultiSet
                d_326_v95_ = _dafny.MultiSet([d_271_v58_])
                index43_ = default__.safeIndex(595, (d_200_v8_).length(0))
                rhs50_ = d_195_v3_
                rhs51_ = d_326_v95_
                rhs52_ = (default__.safeModuloInt((d_291_v77_).f12, (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))])) == (((d_291_v77_).f12) * (d_195_v3_))
                lhs38_ = d_200_v8_
                lhs39_ = default__.safeIndex(595, (d_200_v8_).length(0))
                lhs38_[lhs39_] = rhs50_
                d_326_v95_ = rhs51_
                d_276_v63_ = rhs52_
                d_195_v3_ = (d_271_v58_).f12
            d_206_v13_ = d_206_v13_
            if d_192_v0_:
                (d_203_globalState_).f2 = ((0) - ((d_291_v77_).f12) if (d_271_v58_).f11 else -201)
                d_327_v96_: D7
                d_327_v96_ = D7_DC17(d_192_v0_)
                d_328_v97_: _dafny.Set
                d_328_v97_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([not(d_276_v63_), (d_327_v96_).cf25]))})
                rhs53_ = (_dafny.Set({d_297_cf7_})).issubset(d_328_v97_)
                rhs54_ = len((_dafny.SeqWithoutIsStrInference([d_270_v57_ for d_329_i6_ in range(default__.abs(-994))]) if (d_291_v77_).f11 else _dafny.SeqWithoutIsStrInference([d_270_v57_ for d_330_i7_ in range(default__.abs(-963))])))
                lhs40_ = d_203_globalState_
                d_192_v0_ = rhs53_
                lhs40_.f2 = rhs54_
                d_331_v98_: D0
                d_331_v98_ = D0_DC1((d_271_v58_).f12, len(d_198_v6_), (d_291_v77_).f12)
                d_332_v99_: C1
                nw61_ = C1()
                nw61_.ctor__(((d_331_v98_).cf1) == (d_297_cf7_))
                d_332_v99_ = nw61_
                d_333_v100_: _dafny.Array
                def lambda15_(d_334_i8_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))

                init9_ = lambda15_
                nw62_ = _dafny.Array(None, 8)
                for i0_9_ in range(nw62_.length(0)):
                    nw62_[i0_9_] = init9_(i0_9_)
                d_333_v100_ = nw62_
                index44_ = default__.safeIndex(469, (d_333_v100_).length(0))
                (d_333_v100_)[index44_] = (d_196_v4_) + (d_196_v4_)
                index45_ = default__.safeIndex(469, (d_333_v100_).length(0))
                (d_333_v100_)[index45_] = d_196_v4_
                d_335_v101_: C1
                nw63_ = C1()
                nw63_.ctor__((d_291_v77_).f11)
                d_335_v101_ = nw63_
            elif True:
                d_275_v62_ = (d_275_v62_) + ((d_275_v62_) + (d_275_v62_))
                default__.m0(d_206_v13_, (d_271_v58_).f12, d_209_v14_, (d_291_v77_).f12, d_203_globalState_)
                d_336_v102_: _dafny.Map
                d_336_v102_ = _dafny.Map({d_292_v78_: (d_291_v77_).fm1(d_196_v4_, d_195_v3_, d_203_globalState_)})
                d_297_cf7_ = ((d_271_v58_).f12 if (not((d_271_v58_).f11)) == (default__.fm0(True, (d_271_v58_).f11, (d_271_v58_).f11, (d_271_v58_).f11, d_203_globalState_)) else ((d_336_v102_)[d_292_v78_] if (d_292_v78_) in (d_336_v102_) else (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]))
                (d_203_globalState_).f2 = (d_297_cf7_) - ((d_291_v77_).f12)
                d_206_v13_ = (d_206_v13_ if (d_291_v77_).f11 else d_206_v13_)
            d_337_v103_: _dafny.MultiSet
            d_337_v103_ = _dafny.MultiSet([d_200_v8_, (d_200_v8_ if (d_291_v77_).f11 else d_200_v8_)])
            d_338_v104_: _dafny.Seq
            d_338_v104_ = _dafny.SeqWithoutIsStrInference([d_291_v77_, d_271_v58_])
            d_339_v105_: _dafny.Seq
            d_339_v105_ = d_338_v104_
            d_340_v106_: _dafny.Array
            nw64_ = _dafny.Array(None, 9)
            nw64_[int(0)] = d_338_v104_
            nw64_[int(1)] = d_339_v105_
            nw64_[int(2)] = d_339_v105_
            nw64_[int(3)] = d_338_v104_
            nw64_[int(4)] = d_338_v104_
            nw64_[int(5)] = d_339_v105_
            nw64_[int(6)] = d_339_v105_
            nw64_[int(7)] = (d_338_v104_).set(default__.safeIndex((d_291_v77_).f12, len(d_338_v104_)), d_271_v58_)
            nw64_[int(8)] = d_339_v105_
            d_340_v106_ = nw64_
            d_341_v107_: _dafny.Map
            d_341_v107_ = _dafny.Map({d_340_v106_: _dafny.MultiSet([d_200_v8_])})
            d_337_v103_ = ((d_341_v107_)[d_340_v106_] if (d_340_v106_) in (d_341_v107_) else d_337_v103_)
        elif source3_.is_DC5:
            d_342___mcc_h8_ = source3_.cf8
            d_343___mcc_h9_ = source3_.cf9
            d_344_cf9_ = d_343___mcc_h9_
            d_345_cf8_ = d_342___mcc_h8_
            d_346_v108_: _dafny.MultiSet
            d_346_v108_ = _dafny.MultiSet([default__.fm9((d_291_v77_).f11, d_203_globalState_)])
            d_347_v109_: D1
            d_347_v109_ = D1_DC4(d_346_v108_, (d_291_v77_).f12)
            d_347_v109_ = d_347_v109_
            index46_ = default__.safeIndex(333, (d_206_v13_).length(0))
            (d_206_v13_)[index46_] = (d_271_v58_).f11
            index47_ = default__.safeIndex(333, (d_206_v13_).length(0))
            (d_206_v13_)[index47_] = False
            d_348_v110_: _dafny.Map
            d_348_v110_ = _dafny.Map({d_345_cf8_: d_275_v62_})
            d_349_v111_: _dafny.Map
            d_349_v111_ = _dafny.Map({((d_348_v110_)[d_344_cf9_] if (d_344_cf9_) in (d_348_v110_) else (d_275_v62_).set(default__.safeIndex(d_195_v3_, len(d_275_v62_)), d_276_v63_)): len((d_275_v62_).set(default__.safeIndex(-519, len(d_275_v62_)), not(not(d_345_cf8_))))})
            d_350_v112_: _dafny.MultiSet
            d_350_v112_ = _dafny.MultiSet([len(d_349_v111_)])
            d_195_v3_ = len(_dafny.Set({((d_271_v58_).f12) - ((d_350_v112_).cardinality)}))
            d_351_v113_: D1
            d_351_v113_ = D1_DC5(d_344_cf9_, d_345_cf8_)
            if (d_351_v113_).cf9:
                d_196_v4_ = (d_196_v4_) + (d_196_v4_)
                d_352_v114_: C3
                nw65_ = C3()
                nw65_.ctor__((d_291_v77_).f11, default__.safeDivisionInt(len(default__.fm15((d_291_v77_).f12, (d_291_v77_).f12, d_192_v0_, d_203_globalState_)), d_195_v3_))
                d_352_v114_ = nw65_
                d_353_v115_: _dafny.Seq
                d_353_v115_ = _dafny.SeqWithoutIsStrInference([d_200_v8_, d_200_v8_])
                d_354_v116_: _dafny.Map
                d_354_v116_ = _dafny.Map({d_270_v57_: d_200_v8_})
                d_355_v117_: _dafny.Array
                nw66_ = _dafny.Array(None, 20)
                nw66_[int(0)] = ((d_201_v9_)[83] if (83) in (d_201_v9_) else (d_353_v115_)[default__.safeIndex(-561, len(d_353_v115_))])
                nw66_[int(1)] = d_200_v8_
                nw66_[int(2)] = (d_200_v8_ if (d_291_v77_).f11 else d_200_v8_)
                nw66_[int(3)] = ((d_354_v116_)[d_270_v57_] if (d_270_v57_) in (d_354_v116_) else d_200_v8_)
                nw66_[int(4)] = d_200_v8_
                nw66_[int(5)] = d_200_v8_
                nw66_[int(6)] = d_200_v8_
                nw66_[int(7)] = d_200_v8_
                nw66_[int(8)] = d_200_v8_
                nw66_[int(9)] = d_200_v8_
                nw66_[int(10)] = d_200_v8_
                nw66_[int(11)] = d_200_v8_
                nw66_[int(12)] = d_200_v8_
                nw66_[int(13)] = d_200_v8_
                nw66_[int(14)] = d_200_v8_
                nw66_[int(15)] = d_200_v8_
                nw66_[int(16)] = d_200_v8_
                nw66_[int(17)] = d_200_v8_
                nw66_[int(18)] = d_200_v8_
                nw66_[int(19)] = d_200_v8_
                d_355_v117_ = nw66_
                index48_ = default__.safeIndex(20, (d_355_v117_).length(0))
                (d_355_v117_)[index48_] = d_200_v8_
                index49_ = default__.safeIndex(20, (d_355_v117_).length(0))
                (d_355_v117_)[index49_] = d_200_v8_
                d_352_v114_ = d_271_v58_
                index50_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index50_] = ((0) - ((d_352_v114_).f12)) - (500)
            elif True:
                d_356_v118_: _dafny.Array
                nw67_ = _dafny.Array(False, 5)
                d_356_v118_ = nw67_
                default__.m0(d_356_v118_, default__.safeDivisionInt(d_195_v3_, (d_271_v58_).f12), d_209_v14_, (d_271_v58_).f12, d_203_globalState_)
                d_357_v119_: bool
                d_358_v120_: _dafny.Map
                d_359_v121_: bool
                out9_: bool
                out10_: _dafny.Map
                out11_: bool
                out9_, out10_, out11_ = (d_271_v58_).m2(d_345_cf8_, ((d_271_v58_).f12) <= ((d_291_v77_).f12), d_203_globalState_)
                d_357_v119_ = out9_
                d_358_v120_ = out10_
                d_359_v121_ = out11_
                rhs55_ = (d_198_v6_)[default__.safeIndex(((d_291_v77_).f12 if d_192_v0_ else d_195_v3_), len(d_198_v6_))]
                rhs56_ = (d_275_v62_)[default__.safeIndex(433, len(d_275_v62_))]
                rhs57_ = (d_270_v57_) in (d_196_v4_)
                lhs41_ = d_203_globalState_
                lhs41_.f2 = rhs55_
                d_357_v119_ = rhs56_
                d_276_v63_ = rhs57_
                index51_ = default__.safeIndex(333, (d_206_v13_).length(0))
                rhs58_ = ((d_202_v10_)[d_345_cf8_] if (d_345_cf8_) in (d_202_v10_) else (d_291_v77_).f12)
                rhs59_ = (d_275_v62_) + (default__.fm16(345, (d_351_v113_).cf9, d_203_globalState_))
                rhs60_ = True
                rhs61_ = (d_200_v8_ if not(d_278_v65_) else d_200_v8_)
                rhs62_ = not(d_276_v63_)
                lhs42_ = d_203_globalState_
                lhs43_ = d_206_v13_
                lhs44_ = default__.safeIndex(333, (d_206_v13_).length(0))
                lhs45_ = d_203_globalState_
                lhs42_.f2 = rhs58_
                d_275_v62_ = rhs59_
                lhs43_[lhs44_] = rhs60_
                lhs45_.f10 = rhs61_
                d_344_cf9_ = rhs62_
                d_360_v122_: C2
                nw68_ = C2()
                nw68_.ctor__(d_195_v3_)
                d_360_v122_ = nw68_
        elif source3_.is_DC3:
            d_361___mcc_h10_ = source3_.cf5
            d_362_cf5_ = d_361___mcc_h10_
            d_363_v123_: C0
            nw69_ = C0()
            nw69_.ctor__()
            d_363_v123_ = nw69_
            d_364_v124_: D7
            d_364_v124_ = D7_DC16(d_363_v123_)
            d_364_v124_ = d_364_v124_
            d_365_v125_: C0
            nw70_ = C0()
            nw70_.ctor__()
            d_365_v125_ = nw70_
            default__.m0(d_206_v13_, (d_291_v77_).f12, d_209_v14_, (d_291_v77_).f12, d_203_globalState_)
            index52_ = default__.safeIndex(775, (d_206_v13_).length(0))
            (d_206_v13_)[index52_] = d_276_v63_
            index53_ = default__.safeIndex(775, (d_206_v13_).length(0))
            (d_206_v13_)[index53_] = d_276_v63_
        elif True:
            d_366___mcc_h11_ = source3_.cf10
            d_367_cf10_ = d_366___mcc_h11_
            index54_ = default__.safeIndex(923, (d_206_v13_).length(0))
            (d_206_v13_)[index54_] = True
            index55_ = default__.safeIndex(923, (d_206_v13_).length(0))
            (d_206_v13_)[index55_] = not(d_276_v63_)
            (d_291_v77_).m1((d_291_v77_).f11, d_203_globalState_)
            d_368_v126_: _dafny.Array
            nw71_ = _dafny.Array(False, 9)
            d_368_v126_ = nw71_
            d_369_v127_: _dafny.Map
            d_369_v127_ = _dafny.Map({(d_291_v77_).f12: d_192_v0_})
            d_370_v128_: D3
            d_370_v128_ = D3_DC10((d_206_v13_)[default__.safeIndex(923, (d_206_v13_).length(0))], (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))], True, d_368_v126_, _dafny.Map({(d_271_v58_).f12: (d_291_v77_).f11}))
            pat_let_tv5_ = d_368_v126_
            pat_let_tv6_ = d_271_v58_
            pat_let_tv7_ = d_192_v0_
            def iife27_(_pat_let6_0):
                def iife28_(d_371_dt__update__tmp_h4_):
                    def iife29_(_pat_let7_0):
                        def iife30_(d_372_dt__update_hcf17_h0_):
                            def iife31_(_pat_let8_0):
                                def iife32_(d_373_dt__update_hcf14_h0_):
                                    def iife33_(_pat_let9_0):
                                        def iife34_(d_374_dt__update_hcf16_h0_):
                                            return D3_DC10(d_373_dt__update_hcf14_h0_, (d_371_dt__update__tmp_h4_).cf15, d_374_dt__update_hcf16_h0_, d_372_dt__update_hcf17_h0_, (d_371_dt__update__tmp_h4_).cf18)
                                        return iife34_(_pat_let9_0)
                                    return iife33_(pat_let_tv7_)
                                return iife32_(_pat_let8_0)
                            return iife31_((pat_let_tv6_).f11)
                        return iife30_(_pat_let7_0)
                    return iife29_(pat_let_tv5_)
                return iife28_(_pat_let6_0)
            source5_ = iife27_((D3_DC10(d_278_v65_, (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))], (d_271_v58_).f11, d_368_v126_, d_369_v127_) if (d_291_v77_).f11 else d_370_v128_))
            if source5_.is_DC10:
                d_375___mcc_h18_ = source5_.cf14
                d_376___mcc_h19_ = source5_.cf15
                d_377___mcc_h20_ = source5_.cf16
                d_378___mcc_h21_ = source5_.cf17
                d_379___mcc_h22_ = source5_.cf18
                d_380_cf18_ = d_379___mcc_h22_
                d_381_cf17_ = d_378___mcc_h21_
                d_382_cf16_ = d_377___mcc_h20_
                d_383_cf15_ = d_376___mcc_h19_
                d_384_cf14_ = d_375___mcc_h18_
                index56_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index56_] = (d_271_v58_).fm1(d_196_v4_, len(d_280_v67_), d_203_globalState_)
                d_270_v57_ = _dafny.CodePoint('w')
                d_195_v3_ = ((0) - ((d_271_v58_).f12)) * ((d_271_v58_).f12)
                d_277_v64_ = (d_277_v64_).set(d_275_v62_, default__.fm0(d_192_v0_, (d_271_v58_).f11, d_276_v63_, (d_271_v58_).f11, d_203_globalState_))
            elif source5_.is_DC9:
                d_385___mcc_h23_ = source5_.cf13
                d_386_cf13_ = d_385___mcc_h23_
                index57_ = default__.safeIndex(46, (d_293_v79_).length(0))
                (d_293_v79_)[index57_] = (d_270_v57_ if (d_271_v58_).f11 else d_270_v57_)
                index58_ = default__.safeIndex(46, (d_293_v79_).length(0))
                (d_293_v79_)[index58_] = (d_270_v57_ if (d_276_v63_) not in (d_275_v62_) else d_270_v57_)
                d_387_v129_: _dafny.MultiSet
                d_387_v129_ = _dafny.MultiSet([(d_271_v58_).f12])
                d_387_v129_ = d_387_v129_
                d_388_v130_: D8
                d_388_v130_ = D8_DC20((d_291_v77_).f11, (d_206_v13_)[default__.safeIndex(923, (d_206_v13_).length(0))], ((d_369_v127_)[len(d_204_v11_)] if (len(d_204_v11_)) in (d_369_v127_) else (d_291_v77_).f11), (d_271_v58_).f11)
                d_389_v131_: _dafny.Set
                d_389_v131_ = _dafny.Set({d_196_v4_})
                d_390_v132_: C1
                nw72_ = C1()
                nw72_.ctor__((d_291_v77_).f11)
                d_390_v132_ = nw72_
                d_391_v133_: _dafny.Map
                d_391_v133_ = _dafny.Map({d_390_v132_: d_278_v65_})
                d_278_v65_ = default__.fm0((d_388_v130_).cf28, d_276_v63_, (d_389_v131_).isdisjoint(d_389_v131_), ((d_391_v133_)[d_390_v132_] if (d_390_v132_) in (d_391_v133_) else False), d_203_globalState_)
                d_392_v134_: _dafny.Array
                def lambda16_(d_393_v6_):
                    def lambda17_(d_394_i9_):
                        return d_393_v6_

                    return lambda17_

                init10_ = lambda16_(d_198_v6_)
                nw73_ = _dafny.Array(None, 17)
                for i0_10_ in range(nw73_.length(0)):
                    nw73_[i0_10_] = init10_(i0_10_)
                d_392_v134_ = nw73_
                d_395_v135_: _dafny.Map
                d_395_v135_ = _dafny.Map({d_392_v134_: True})
                index59_ = default__.safeIndex(923, (d_206_v13_).length(0))
                (d_206_v13_)[index59_] = ((d_395_v135_)[d_392_v134_] if (d_392_v134_) in (d_395_v135_) else (d_291_v77_).f11)
            elif True:
                d_396___mcc_h24_ = source5_.cf19
                d_397_cf19_ = d_396___mcc_h24_
                index60_ = default__.safeIndex(923, (d_206_v13_).length(0))
                rhs63_ = default__.safeModuloInt(default__.fm2(d_278_v65_, d_203_globalState_), (d_291_v77_).f12)
                rhs64_ = not (d_192_v0_) or (True)
                lhs46_ = d_203_globalState_
                lhs47_ = d_206_v13_
                lhs48_ = default__.safeIndex(923, (d_206_v13_).length(0))
                lhs46_.f2 = rhs63_
                lhs47_[lhs48_] = rhs64_
                d_398_v136_: _dafny.MultiSet
                d_398_v136_ = _dafny.MultiSet([d_195_v3_, (d_291_v77_).f12, (d_291_v77_).f12])
                (d_203_globalState_).f2 = (((d_202_v10_)[(d_291_v77_).f11] if ((d_291_v77_).f11) in (d_202_v10_) else (d_291_v77_).f12) if ((d_271_v58_).f12) < ((d_398_v136_).cardinality) else (0) - ((d_271_v58_).f12))
                d_278_v65_ = (d_370_v128_).cf16
                d_270_v57_ = (_dafny.CodePoint('c') if (_dafny.Set({not((d_291_v77_).f11), d_278_v65_, (d_291_v77_).f11, (d_271_v58_).f11})).issubset(_dafny.Set({(d_206_v13_)[default__.safeIndex(923, (d_206_v13_).length(0))], not(d_276_v63_), d_276_v63_})) else d_270_v57_)
            d_280_v67_ = (d_280_v67_).set(d_195_v3_, (0) - (len(d_205_v12_)))
        source6_ = D0_DC0(d_192_v0_)
        if source6_.is_DC1:
            d_399___mcc_h25_ = source6_.cf1
            d_400___mcc_h26_ = source6_.cf2
            d_401___mcc_h27_ = source6_.cf3
            d_402_cf3_ = d_401___mcc_h27_
            d_403_cf2_ = d_400___mcc_h26_
            d_404_cf1_ = d_399___mcc_h25_
            source7_ = default__.fm17((d_291_v77_).f11, (d_198_v6_) + (d_198_v6_), (d_275_v62_)[default__.safeIndex(d_403_cf2_, len(d_275_v62_))], d_192_v0_, d_203_globalState_)
            if source7_.is_DC4:
                d_405___mcc_h30_ = source7_.cf6
                d_406___mcc_h31_ = source7_.cf7
                d_407_cf7_ = d_406___mcc_h31_
                d_408_cf6_ = d_405___mcc_h30_
                d_409_v137_: D11
                d_409_v137_ = D11_DC23(d_198_v6_)
                d_198_v6_ = ((d_409_v137_).cf34) + (((d_198_v6_ if (d_271_v58_).f11 else d_198_v6_)).set(default__.safeIndex(len(d_196_v4_), len((d_198_v6_ if (d_271_v58_).f11 else d_198_v6_))), len(d_197_v5_)))
                d_410_v138_: bool
                d_411_v139_: _dafny.Map
                d_412_v140_: bool
                out12_: bool
                out13_: _dafny.Map
                out14_: bool
                out12_, out13_, out14_ = (d_291_v77_).m2((d_271_v58_).f11, ((0) - (d_404_cf1_)) <= ((d_291_v77_).f12), d_203_globalState_)
                d_410_v138_ = out12_
                d_411_v139_ = out13_
                d_412_v140_ = out14_
                d_413_v141_: bool
                d_414_v142_: _dafny.Map
                d_415_v143_: bool
                out15_: bool
                out16_: _dafny.Map
                out17_: bool
                out15_, out16_, out17_ = (d_271_v58_).m2(((d_291_v77_).f11) or (d_276_v63_), ((d_271_v58_).f12) != ((d_291_v77_).f12), d_203_globalState_)
                d_413_v141_ = out15_
                d_414_v142_ = out16_
                d_415_v143_ = out17_
                d_416_v144_: _dafny.Seq
                d_416_v144_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(d_271_v58_).f11, (d_291_v77_).f11]) if d_278_v65_ else d_275_v62_), d_275_v62_, default__.fm16((d_291_v77_).f12, (d_291_v77_).f11, d_203_globalState_)])
                d_416_v144_ = d_416_v144_
            elif source7_.is_DC5:
                d_417___mcc_h32_ = source7_.cf8
                d_418___mcc_h33_ = source7_.cf9
                d_419_cf9_ = d_418___mcc_h33_
                d_420_cf8_ = d_417___mcc_h32_
                d_421_v145_: _dafny.Map
                d_421_v145_ = _dafny.Map({d_402_cf3_: ((d_271_v58_).f12) <= ((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))])})
                d_421_v145_ = d_421_v145_
                (d_203_globalState_).f2 = (d_291_v77_).f12
                d_422_v146_: D0
                d_422_v146_ = D0_DC1(287, (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))], (d_271_v58_).f12)
                d_195_v3_ = (d_198_v6_)[default__.safeIndex(default__.safeDivisionInt((d_422_v146_).cf2, (0) - (((d_202_v10_)[(d_271_v58_).f11] if ((d_271_v58_).f11) in (d_202_v10_) else (d_271_v58_).f12))), len(d_198_v6_))]
                d_196_v4_ = d_196_v4_
            elif source7_.is_DC3:
                d_423___mcc_h34_ = source7_.cf5
                d_424_cf5_ = d_423___mcc_h34_
                d_275_v62_ = d_275_v62_
                d_425_v147_: _dafny.Map
                d_425_v147_ = _dafny.Map({len(d_196_v4_): d_206_v13_})
                d_206_v13_ = ((d_425_v147_)[(d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]] if ((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]) in (d_425_v147_) else d_206_v13_)
                d_196_v4_ = (d_196_v4_) + (d_196_v4_)
                d_276_v63_ = (default__.safeDivisionInt((d_271_v58_).f12, (0) - ((d_271_v58_).f12))) <= ((d_291_v77_).f12)
            elif True:
                d_426___mcc_h35_ = source7_.cf10
                d_427_cf10_ = d_426___mcc_h35_
                d_428_v148_: _dafny.Array
                nw74_ = _dafny.Array(None, 15)
                nw74_[int(0)] = d_206_v13_
                nw74_[int(1)] = d_206_v13_
                nw74_[int(2)] = d_206_v13_
                nw74_[int(3)] = d_206_v13_
                nw74_[int(4)] = d_206_v13_
                nw74_[int(5)] = d_206_v13_
                nw74_[int(6)] = d_206_v13_
                nw74_[int(7)] = d_206_v13_
                nw74_[int(8)] = d_206_v13_
                nw74_[int(9)] = d_206_v13_
                nw74_[int(10)] = d_206_v13_
                nw74_[int(11)] = d_206_v13_
                nw74_[int(12)] = d_206_v13_
                nw74_[int(13)] = d_206_v13_
                nw74_[int(14)] = d_206_v13_
                d_428_v148_ = nw74_
                d_429_v149_: _dafny.Map
                d_429_v149_ = _dafny.Map({673: d_428_v148_})
                d_430_v150_: _dafny.Map
                d_430_v150_ = _dafny.Map({(d_271_v58_).f11: ((d_429_v149_)[(d_271_v58_).f12] if ((d_271_v58_).f12) in (d_429_v149_) else d_428_v148_)})
                d_430_v150_ = (d_430_v150_).set(True, d_428_v148_)
                d_270_v57_ = d_270_v57_
                (d_203_globalState_).f2 = (0) - (d_402_cf3_)
                d_431_v151_: _dafny.Array
                def lambda18_(d_432_v58_):
                    def lambda19_(d_433_i10_):
                        return (D11_DC25() if (d_432_v58_).f11 else D11_DC25())

                    return lambda19_

                init11_ = lambda18_(d_271_v58_)
                nw75_ = _dafny.Array(None, 29)
                for i0_11_ in range(nw75_.length(0)):
                    nw75_[i0_11_] = init11_(i0_11_)
                d_431_v151_ = nw75_
                d_434_v152_: D11
                d_434_v152_ = D11_DC25()
                index61_ = default__.safeIndex(150, (d_431_v151_).length(0))
                (d_431_v151_)[index61_] = d_434_v152_
                d_435_v153_: _dafny.Map
                d_435_v153_ = _dafny.Map({default__.safeDivisionInt(d_402_cf3_, (d_271_v58_).fm1(d_196_v4_, (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))], d_203_globalState_)): _dafny.SeqWithoutIsStrInference([(d_291_v77_).f12])})
                index62_ = default__.safeIndex(595, (d_200_v8_).length(0))
                index63_ = default__.safeIndex(595, (d_200_v8_).length(0))
                index64_ = default__.safeIndex(150, (d_431_v151_).length(0))
                index65_ = default__.safeIndex(595, (d_200_v8_).length(0))
                rhs65_ = len(((d_435_v153_)[d_404_cf1_] if (d_404_cf1_) in (d_435_v153_) else d_198_v6_))
                rhs66_ = False
                rhs67_ = (d_404_cf1_ if (d_270_v57_) in (d_196_v4_) else (d_291_v77_).f12)
                rhs68_ = d_434_v152_
                rhs69_ = default__.safeModuloInt(d_195_v3_, len(d_196_v4_))
                lhs49_ = d_200_v8_
                lhs50_ = default__.safeIndex(595, (d_200_v8_).length(0))
                lhs51_ = d_200_v8_
                lhs52_ = default__.safeIndex(595, (d_200_v8_).length(0))
                lhs53_ = d_431_v151_
                lhs54_ = default__.safeIndex(150, (d_431_v151_).length(0))
                lhs55_ = d_200_v8_
                lhs56_ = default__.safeIndex(595, (d_200_v8_).length(0))
                lhs49_[lhs50_] = rhs65_
                d_278_v65_ = rhs66_
                lhs51_[lhs52_] = rhs67_
                lhs53_[lhs54_] = rhs68_
                lhs55_[lhs56_] = rhs69_
            d_436_v154_: C3
            nw76_ = C3()
            nw76_.ctor__(d_276_v63_, default__.safeModuloInt(len(d_280_v67_), d_402_cf3_))
            d_436_v154_ = nw76_
            source8_ = default__.fm17((d_271_v58_).f11, d_198_v6_, not((len(_dafny.SeqWithoutIsStrInference([(d_271_v58_).f11]))) <= ((d_291_v77_).f12)), (d_196_v4_) <= (d_196_v4_), d_203_globalState_)
            if source8_.is_DC4:
                d_437___mcc_h36_ = source8_.cf6
                d_438___mcc_h37_ = source8_.cf7
                d_439_cf7_ = d_438___mcc_h37_
                d_440_cf6_ = d_437___mcc_h36_
                d_278_v65_ = (d_196_v4_) < (_dafny.SeqWithoutIsStrInference([d_270_v57_ for d_441_i11_ in range(default__.abs(379))]))
                d_278_v65_ = default__.fm0((d_205_v12_) != (d_205_v12_), (d_291_v77_).f11, d_192_v0_, not((d_436_v154_).f11), d_203_globalState_)
                d_276_v63_ = ((d_291_v77_).f11 if (d_271_v58_).f11 else (d_436_v154_).f11)
                d_442_v155_: C1
                nw77_ = C1()
                nw77_.ctor__((not((d_291_v77_).f11)) == ((d_271_v58_).f11))
                d_442_v155_ = nw77_
            elif source8_.is_DC5:
                d_443___mcc_h38_ = source8_.cf8
                d_444___mcc_h39_ = source8_.cf9
                d_445_cf9_ = d_444___mcc_h39_
                d_446_cf8_ = d_443___mcc_h38_
                d_196_v4_ = d_196_v4_
                (d_291_v77_).m1(((d_271_v58_).f12) >= ((d_436_v154_).f12), d_203_globalState_)
                d_447_v156_: C1
                nw78_ = C1()
                nw78_.ctor__((d_275_v62_) < (d_275_v62_))
                d_447_v156_ = nw78_
                d_448_v157_: C3
                nw79_ = C3()
                nw79_.ctor__((d_291_v77_).f11, default__.safeModuloInt(d_402_cf3_, len(default__.fm18(d_403_cf2_, d_203_globalState_))))
                d_448_v157_ = nw79_
            elif source8_.is_DC3:
                d_449___mcc_h40_ = source8_.cf5
                d_450_cf5_ = d_449___mcc_h40_
                d_451_v158_: D1
                d_451_v158_ = D1_DC3(d_450_cf5_)
                d_450_cf5_ = (d_451_v158_).cf5
                d_452_v159_: _dafny.Seq
                d_452_v159_ = _dafny.SeqWithoutIsStrInference([(d_202_v10_).set((d_271_v58_).f11, default__.abs((d_291_v77_).f12))])
                d_270_v57_ = default__.fm6((d_436_v154_).f12, (d_452_v159_) != (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([(d_291_v77_).f11]) for d_453_i12_ in range(default__.abs(712))])), (d_291_v77_).f11, d_450_cf5_, d_203_globalState_)
                (d_203_globalState_).f2 = len(d_198_v6_)
                index66_ = default__.safeIndex(918, (d_206_v13_).length(0))
                (d_206_v13_)[index66_] = True
                index67_ = default__.safeIndex(918, (d_206_v13_).length(0))
                (d_206_v13_)[index67_] = d_276_v63_
            elif True:
                d_454___mcc_h41_ = source8_.cf10
                d_455_cf10_ = d_454___mcc_h41_
                d_456_v160_: _dafny.Map
                d_456_v160_ = _dafny.Map({d_200_v8_: (d_271_v58_).f11})
                d_456_v160_ = (d_456_v160_).set(d_200_v8_, (default__.fm6(d_195_v3_, (d_291_v77_).f11, d_192_v0_, d_270_v57_, d_203_globalState_)) not in (d_196_v4_))
                d_402_cf3_ = (default__.safeDivisionInt((0) - (d_403_cf2_), (d_436_v154_).f12)) + ((d_436_v154_).fm1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sn")), d_404_cf1_, d_203_globalState_))
                d_457_v161_: C0
                nw80_ = C0()
                nw80_.ctor__()
                d_457_v161_ = nw80_
                d_458_v162_: D7
                d_458_v162_ = D7_DC16(d_457_v161_)
                d_459_v163_: _dafny.Map
                d_459_v163_ = _dafny.Map({(d_458_v162_).cf24: len(d_198_v6_)})
                d_460_v164_: _dafny.Map
                d_460_v164_ = _dafny.Map({d_459_v163_: (0) - (len(_dafny.SeqWithoutIsStrInference([D11_DC25() for d_461_i13_ in range(default__.abs(376))])))})
                d_460_v164_ = d_460_v164_
                (d_203_globalState_).f2 = d_195_v3_
            d_462_v165_: D3
            d_462_v165_ = D3_DC9((d_206_v13_ if True else d_206_v13_))
            source9_ = d_462_v165_
            if source9_.is_DC10:
                d_463___mcc_h42_ = source9_.cf14
                d_464___mcc_h43_ = source9_.cf15
                d_465___mcc_h44_ = source9_.cf16
                d_466___mcc_h45_ = source9_.cf17
                d_467___mcc_h46_ = source9_.cf18
                d_468_cf18_ = d_467___mcc_h46_
                d_469_cf17_ = d_466___mcc_h45_
                d_470_cf16_ = d_465___mcc_h44_
                d_471_cf15_ = d_464___mcc_h43_
                d_472_cf14_ = d_463___mcc_h42_
                rhs70_ = (d_271_v58_).f11
                rhs71_ = (d_196_v4_)[default__.safeIndex(d_471_cf15_, len(d_196_v4_))]
                rhs72_ = not(not(False))
                d_470_cf16_ = rhs70_
                d_270_v57_ = rhs71_
                d_276_v63_ = rhs72_
                d_473_v166_: _dafny.Seq
                d_473_v166_ = _dafny.SeqWithoutIsStrInference([(d_275_v62_) + (d_275_v62_)])
                d_473_v166_ = d_473_v166_
                d_471_cf15_ = len((d_280_v67_) | ((d_280_v67_) | (d_280_v67_)))
                d_204_v11_ = d_204_v11_
            elif source9_.is_DC9:
                d_474___mcc_h47_ = source9_.cf13
                d_475_cf13_ = d_474___mcc_h47_
                d_195_v3_ = len(((d_196_v4_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgqwbmtnb")))) + (((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_476_i14_ in range(default__.abs(58))])).set(default__.safeIndex(156, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_477_i14_ in range(default__.abs(58))]))), default__.fm6(d_195_v3_, (d_291_v77_).f11, d_278_v65_, _dafny.CodePoint('b'), d_203_globalState_))) + (d_196_v4_)))
                index68_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index68_] = (d_271_v58_).f12
                d_478_v167_: D0
                d_478_v167_ = D0_DC1((d_271_v58_).f12, (0) - ((d_271_v58_).fm1(_dafny.SeqWithoutIsStrInference([d_270_v57_ for d_479_i15_ in range(default__.abs(637))]), (d_271_v58_).f12, d_203_globalState_)), (d_436_v154_).f12)
                d_195_v3_ = (d_478_v167_).cf3
                index69_ = default__.safeIndex(714, (d_475_cf13_).length(0))
                (d_475_cf13_)[index69_] = not (True) or (d_278_v65_)
                d_480_v168_: D11
                d_480_v168_ = D11_DC26((d_291_v77_).f12, d_200_v8_, (d_436_v154_).f11)
                d_481_v169_: D0
                d_481_v169_ = D0_DC0((d_480_v168_).cf38)
                index70_ = default__.safeIndex(714, (d_475_cf13_).length(0))
                (d_475_cf13_)[index70_] = (d_481_v169_).cf0
            elif True:
                d_482___mcc_h48_ = source9_.cf19
                d_483_cf19_ = d_482___mcc_h48_
                d_403_cf2_ = (len(d_193_v1_)) - (d_403_cf2_)
                d_484_v170_: _dafny.Array
                nw81_ = _dafny.Array(None, 9)
                d_484_v170_ = nw81_
                d_485_v171_: C1
                nw82_ = C1()
                nw82_.ctor__((d_291_v77_).f11)
                d_485_v171_ = nw82_
                index71_ = default__.safeIndex(23, (d_484_v170_).length(0))
                (d_484_v170_)[index71_] = d_485_v171_
                d_486_v172_: C1
                d_486_v172_ = d_485_v171_
                index72_ = default__.safeIndex(23, (d_484_v170_).length(0))
                (d_484_v170_)[index72_] = (d_486_v172_)
                index73_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index73_] = ((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]) * ((d_271_v58_).f12)
                index74_ = default__.safeIndex(32, (d_206_v13_).length(0))
                (d_206_v13_)[index74_] = d_278_v65_
                index75_ = default__.safeIndex(32, (d_206_v13_).length(0))
                (d_206_v13_)[index75_] = (d_436_v154_).f11
        elif source6_.is_DC0:
            d_487___mcc_h28_ = source6_.cf0
            d_488_cf0_ = d_487___mcc_h28_
            d_489_v173_: _dafny.Map
            d_489_v173_ = _dafny.Map({d_278_v65_: d_270_v57_})
            d_490_v174_: bool
            d_491_v175_: _dafny.Map
            d_492_v176_: bool
            out18_: bool
            out19_: _dafny.Map
            out20_: bool
            out18_, out19_, out20_ = (d_291_v77_).m2((len(d_489_v173_)) == ((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]), (d_271_v58_).f11, d_203_globalState_)
            d_490_v174_ = out18_
            d_491_v175_ = out19_
            d_492_v176_ = out20_
            (d_271_v58_).m1(d_278_v65_, d_203_globalState_)
            d_490_v174_ = (d_198_v6_) != ((d_198_v6_) + (d_198_v6_))
            d_493_v177_: _dafny.Map
            d_493_v177_ = _dafny.Map({len(d_204_v11_): d_490_v174_})
            d_494_v178_: D3
            d_494_v178_ = D3_DC10(d_490_v174_, (d_271_v58_).f12, d_276_v63_, d_206_v13_, d_493_v177_)
            d_204_v11_ = (d_204_v11_).set(((d_494_v178_).cf15) >= ((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]), d_276_v63_)
        elif True:
            d_495___mcc_h29_ = source6_.cf4
            d_496_cf4_ = d_495___mcc_h29_
            d_497_v179_: _dafny.Seq
            d_497_v179_ = _dafny.SeqWithoutIsStrInference([d_196_v4_, d_196_v4_, d_196_v4_, _dafny.SeqWithoutIsStrInference([d_270_v57_ for d_498_i16_ in range(default__.abs(570))])])
            if (d_497_v179_) < (_dafny.SeqWithoutIsStrInference([d_196_v4_ for d_499_i17_ in range(default__.abs(24))])):
                d_500_v180_: _dafny.Map
                d_500_v180_ = _dafny.Map({(0) - ((d_271_v58_).f12): (d_271_v58_).f11})
                d_500_v180_ = (d_500_v180_).set(d_195_v3_, True)
                index76_ = default__.safeIndex(578, (d_206_v13_).length(0))
                (d_206_v13_)[index76_] = (d_291_v77_).f11
                index77_ = default__.safeIndex(578, (d_206_v13_).length(0))
                (d_206_v13_)[index77_] = (D3_DC10((d_291_v77_).f11, (d_291_v77_).f12, d_278_v65_, d_206_v13_, d_500_v180_)).cf16
                (d_291_v77_).m1(((d_271_v58_).f12) >= ((d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]), d_203_globalState_)
                d_196_v4_ = d_196_v4_
                (d_291_v77_).m1((d_291_v77_).f11, d_203_globalState_)
            elif True:
                d_501_v181_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.Array(None, 0), 17)
                d_501_v181_ = nw83_
                index78_ = default__.safeIndex(12, (d_501_v181_).length(0))
                (d_501_v181_)[index78_] = d_206_v13_
                index79_ = default__.safeIndex(12, (d_501_v181_).length(0))
                (d_501_v181_)[index79_] = d_206_v13_
                d_502_v182_: C3
                nw84_ = C3()
                nw84_.ctor__((d_291_v77_).f11, (d_291_v77_).f12)
                d_502_v182_ = nw84_
                index80_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index80_] = default__.fm4(d_203_globalState_)
                d_503_v183_: _dafny.Seq
                d_503_v183_ = d_196_v4_
                (d_203_globalState_).f2 = len((d_503_v183_))
                d_504_v184_: D11
                d_504_v184_ = D11_DC24((d_502_v182_).f12)
                index81_ = default__.safeIndex(595, (d_200_v8_).length(0))
                (d_200_v8_)[index81_] = (d_504_v184_).cf35
            d_505_v185_: D0
            d_505_v185_ = D0_DC1(d_195_v3_, (d_291_v77_).f12, (d_271_v58_).f12)
            rhs73_ = ((d_291_v77_).f12) * ((d_198_v6_)[default__.safeIndex(len(d_275_v62_), len(d_198_v6_))])
            rhs74_ = (694) > ((d_505_v185_).cf2)
            rhs75_ = _dafny.Map({(0) - ((d_291_v77_).f12): (d_271_v58_).f12})
            rhs76_ = default__.safeModuloInt(default__.safeModuloInt((d_271_v58_).f12, 697), (d_291_v77_).f12)
            d_195_v3_ = rhs73_
            d_278_v65_ = rhs74_
            d_280_v67_ = rhs75_
            d_195_v3_ = rhs76_
            d_506_v186_: _dafny.Map
            d_506_v186_ = _dafny.Map({d_196_v4_: True})
            d_507_v187_: D0
            d_507_v187_ = D0_DC0(((d_506_v186_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddnbjifo"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddnbjifo"))) in (d_506_v186_) else d_278_v65_))
            source10_ = d_507_v187_
            if source10_.is_DC1:
                d_508___mcc_h49_ = source10_.cf1
                d_509___mcc_h50_ = source10_.cf2
                d_510___mcc_h51_ = source10_.cf3
                d_511_cf3_ = d_510___mcc_h51_
                d_512_cf2_ = d_509___mcc_h50_
                d_513_cf1_ = d_508___mcc_h49_
                d_514_v188_: _dafny.Map
                d_514_v188_ = _dafny.Map({default__.fm19(d_203_globalState_): d_195_v3_})
                d_512_cf2_ = (default__.safeModuloInt((d_291_v77_).f12, d_512_cf2_) if (d_271_v58_).f11 else len(d_514_v188_))
                d_278_v65_ = (d_291_v77_).f11
                d_192_v0_ = ((d_291_v77_).f12) <= (d_511_cf3_)
                d_192_v0_ = (d_291_v77_).f11
            elif source10_.is_DC0:
                d_515___mcc_h52_ = source10_.cf0
                d_516_cf0_ = d_515___mcc_h52_
                d_517_v189_: C1
                nw85_ = C1()
                nw85_.ctor__(False)
                d_517_v189_ = nw85_
                d_516_cf0_ = d_516_cf0_
                d_506_v186_ = (d_506_v186_).set(d_196_v4_, True)
                d_518_v190_: bool
                d_519_v191_: _dafny.Map
                d_520_v192_: bool
                out21_: bool
                out22_: _dafny.Map
                out23_: bool
                out21_, out22_, out23_ = (d_291_v77_).m2(((d_291_v77_).f11 if d_278_v65_ else (d_271_v58_).f11), not(d_276_v63_), d_203_globalState_)
                d_518_v190_ = out21_
                d_519_v191_ = out22_
                d_520_v192_ = out23_
            elif True:
                d_521___mcc_h53_ = source10_.cf4
                d_522_cf4_ = d_521___mcc_h53_
                d_523_v193_: bool
                d_524_v194_: _dafny.Map
                d_525_v195_: bool
                out24_: bool
                out25_: _dafny.Map
                out26_: bool
                out24_, out25_, out26_ = (d_291_v77_).m2(not((d_271_v58_).f11), d_278_v65_, d_203_globalState_)
                d_523_v193_ = out24_
                d_524_v194_ = out25_
                d_525_v195_ = out26_
                d_526_v196_: D7
                d_526_v196_ = D7_DC17((d_291_v77_).f11)
                d_278_v65_ = default__.fm0((d_525_v195_) == ((d_291_v77_).f11), (d_526_v196_).cf25, d_523_v193_, d_276_v63_, d_203_globalState_)
                d_525_v195_ = d_192_v0_
                d_527_v197_: _dafny.Array
                def lambda20_(d_528_v4_):
                    def lambda21_(d_529_i18_):
                        return (d_528_v4_)

                    return lambda21_

                init12_ = lambda20_(d_196_v4_)
                nw86_ = _dafny.Array(None, 19)
                for i0_12_ in range(nw86_.length(0)):
                    nw86_[i0_12_] = init12_(i0_12_)
                d_527_v197_ = nw86_
                index82_ = default__.safeIndex(130, (d_527_v197_).length(0))
                (d_527_v197_)[index82_] = d_196_v4_
                index83_ = default__.safeIndex(130, (d_527_v197_).length(0))
                rhs77_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "c"))) + (d_196_v4_)
                rhs78_ = d_525_v195_
                lhs57_ = d_527_v197_
                lhs58_ = default__.safeIndex(130, (d_527_v197_).length(0))
                lhs57_[lhs58_] = rhs77_
                d_523_v193_ = rhs78_
            d_530_v198_: _dafny.Set
            d_530_v198_ = _dafny.Set({(d_291_v77_).f12, len(_dafny.SeqWithoutIsStrInference([(0) - (len(d_193_v1_)) for d_531_i19_ in range(default__.abs(-9))])), (d_200_v8_)[default__.safeIndex(595, (d_200_v8_).length(0))]})
            d_532_v199_: _dafny.Map
            d_532_v199_ = _dafny.Map({d_530_v198_: d_192_v0_})
            d_533_v200_: _dafny.Seq
            d_533_v200_ = _dafny.SeqWithoutIsStrInference([(d_532_v199_) | (d_532_v199_), d_532_v199_, d_532_v199_, _dafny.Map({d_530_v198_: False})])
            rhs79_ = ((0) - (((d_271_v58_).f12) - (d_195_v3_)) if (d_271_v58_).f11 else default__.safeDivisionInt((0) - ((d_271_v58_).f12), len(d_196_v4_)))
            rhs80_ = (d_533_v200_)[default__.safeIndex(d_195_v3_, len(d_533_v200_))]
            d_195_v3_ = rhs79_
            d_532_v199_ = rhs80_
        _dafny.print(_dafny.string_of(d_192_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_193_v1_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_194_v2_) == (_dafny.Set({_dafny.Set({True})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_195_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_196_v4_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_197_v5_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bk"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_v6_) == (_dafny.SeqWithoutIsStrInference([1, 1, 848]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_199_v7_) == (_dafny.SeqWithoutIsStrInference([_dafny.Set({True}), _dafny.Set({True}), _dafny.Set({True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_200_v8_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_201_v9_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_202_v10_) == (_dafny.MultiSet([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_203_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_203_globalState_).f4) == (_dafny.Set({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_203_globalState_).f5) == (_dafny.Map({848: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_203_globalState_).f6) == (_dafny.SeqWithoutIsStrInference([1, 1, 848]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f7) == (_dafny.SeqWithoutIsStrInference([490]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_203_globalState_.f10)[28]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_204_v11_) == (_dafny.Map({True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_205_v12_) == (_dafny.Map({True: 848}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_206_v13_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[0]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[1]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[2]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[3]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[4]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[5]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[6]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[7]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[8]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[9]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[10]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[11]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[12]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_209_v14_)[13]) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v56_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sharml")): 848}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_270_v57_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_v58_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_271_v58_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_272_v59_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_273_v60_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_274_v61_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_275_v62_) == (_dafny.SeqWithoutIsStrInference([False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_276_v63_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_277_v64_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([False, True, True]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_278_v65_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_279_v66_) == (_dafny.Map({D0_DC1(848, 848, 848): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_280_v67_) == (_dafny.Map({848: 848}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_281_v68_) == (_dafny.Map({False: _dafny.Map({424: 154}), True: _dafny.Map({2: 114})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len((d_282_v69_))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_283_v70_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_284_i4_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v77_).f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_291_v77_).f12))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_292_v78_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_294_v80_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC1(self) -> bool:
        return isinstance(self, D0_DC1)
    @property
    def is_DC0(self) -> bool:
        return isinstance(self, D0_DC0)
    @property
    def is_DC2(self) -> bool:
        return isinstance(self, D0_DC2)

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC0(D0, NamedTuple('DC0', [('cf0', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC0({_dafny.string_of(self.cf0)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC0) and self.cf0 == __o.cf0
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(_dafny.MultiSet({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)

class D1_DC4(D1, NamedTuple('DC4', [('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf5 == __o.cf5
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf10', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf10)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf10 == __o.cf10
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(False, int(0), False, _dafny.Array(None, 0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)

class D3_DC10(D3, NamedTuple('DC10', [('cf14', Any), ('cf15', Any), ('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf14 == __o.cf14 and self.cf15 == __o.cf15 and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC14(D5, NamedTuple('DC14', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC15(D6, NamedTuple('DC15', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC17(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)

class D7_DC17(D7, NamedTuple('DC17', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC18(D7, NamedTuple('DC18', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(False, False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)

class D8_DC20(D8, NamedTuple('DC20', [('cf28', Any), ('cf29', Any), ('cf30', Any), ('cf31', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf28)}, {_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf28 == __o.cf28 and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC21(D9, NamedTuple('DC21', [('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({self.cf32.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D10_DC22)

class D10_DC22(D10, NamedTuple('DC22', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC22({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC22) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC24(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D11_DC24)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D11_DC23)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC24(D11, NamedTuple('DC24', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC24({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC24) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25)
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC26(D11, NamedTuple('DC26', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC23(D11, NamedTuple('DC23', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC23({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC23) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)

class D12_DC28(D12, NamedTuple('DC28', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)

class D13_DC29(D13, NamedTuple('DC29', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D14_DC30)

class D14_DC30(D14, NamedTuple('DC30', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC30({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC30) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC32(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D15_DC32)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D15_DC31)

class D15_DC32(D15, NamedTuple('DC32', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC32({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC32) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC31(D15, NamedTuple('DC31', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC31({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC31) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class GlobalState:
    def  __init__(self):
        self.f2: int = int(0)
        self.f7: _dafny.Seq = _dafny.Seq({})
        self.f10: _dafny.Array = _dafny.Array(None, 0)
        self._f0: int = int(0)
        self._f1: bool = False
        self._f3: bool = False
        self._f4: _dafny.Set = _dafny.Set({})
        self._f5: _dafny.Map = _dafny.Map({})
        self._f6: _dafny.Seq = _dafny.Seq({})
        self._f8: bool = False
        self._f9: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2
        (self)._f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self).f7 = f7
        (self)._f8 = f8
        (self)._f9 = f9
        (self).f10 = f10

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
    def f8(self):
        return self._f8
    @property
    def f9(self):
        return self._f9

class C0:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self):
        pass
        pass

    def fm7(self, p0, p1, p2, p3, globalState):
        return ((107) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wt"))))) < (-225)

    def fm8(self, p0, p1, p2, globalState):
        return ((341) - (-256)) + (len((_dafny.Set({789, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lan"))), len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([150])), len(_dafny.SeqWithoutIsStrInference([False]))})), len(_dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghqr")))})): False}))})).intersection(_dafny.Set({330}))))


class C1:
    def  __init__(self):
        self.f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f14):
        (self).f14 = f14

    def m4(self, p0, globalState):
        r0: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        r1: D1 = D1.default()()
        r2: int = int(0)
        d_534_v0_: int
        d_534_v0_ = -769
        r2 = d_534_v0_
        d_534_v0_ = (d_534_v0_) + (358)
        d_535_v1_: _dafny.Seq
        d_535_v1_ = _dafny.SeqWithoutIsStrInference([341, d_534_v0_])
        d_536_v2_: str
        d_536_v2_ = _dafny.CodePoint('q')
        d_537_v3_: _dafny.MultiSet
        d_537_v3_ = _dafny.MultiSet([d_534_v0_])
        hi1_ = ((d_537_v3_).cardinality) * (d_534_v0_)
        for d_538_i0_ in range((d_535_v1_)[default__.safeIndex(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))).set(default__.safeIndex(d_534_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n")))), d_536_v2_)), len(d_535_v1_))], hi1_):
            d_539_v4_: _dafny.Array
            nw87_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_539_v4_ = nw87_
            d_540_v5_: _dafny.Array
            nw88_ = _dafny.Array(False, 9)
            d_540_v5_ = nw88_
            d_541_v6_: _dafny.MultiSet
            d_541_v6_ = _dafny.MultiSet([d_540_v5_, d_540_v5_])
            d_542_v7_: _dafny.Map
            d_542_v7_ = _dafny.Map({True: self.f14})
            d_543_v8_: _dafny.Seq
            d_543_v8_ = _dafny.SeqWithoutIsStrInference([d_542_v7_, d_542_v7_, d_542_v7_, d_542_v7_])
            d_544_v9_: _dafny.Map
            d_544_v9_ = _dafny.Map({len(d_543_v8_): d_534_v0_})
            d_545_v10_: _dafny.Array
            nw89_ = _dafny.Array(None, 19)
            nw89_[int(0)] = len(d_535_v1_)
            nw89_[int(1)] = 917
            nw89_[int(2)] = d_538_i0_
            nw89_[int(3)] = d_534_v0_
            nw89_[int(4)] = (0) - (len(_dafny.Set({self.f14, self.f14})))
            nw89_[int(5)] = (d_541_v6_).cardinality
            nw89_[int(6)] = d_534_v0_
            nw89_[int(7)] = d_538_i0_
            nw89_[int(8)] = 691
            nw89_[int(9)] = len(d_542_v7_)
            nw89_[int(10)] = d_538_i0_
            nw89_[int(11)] = d_534_v0_
            nw89_[int(12)] = d_538_i0_
            nw89_[int(13)] = d_538_i0_
            nw89_[int(14)] = 731
            nw89_[int(15)] = d_538_i0_
            nw89_[int(16)] = ((d_544_v9_)[d_538_i0_] if (d_538_i0_) in (d_544_v9_) else default__.fm4(globalState))
            nw89_[int(17)] = d_534_v0_
            nw89_[int(18)] = d_538_i0_
            d_545_v10_ = nw89_
            index84_ = default__.safeIndex(770, (d_539_v4_).length(0))
            (d_539_v4_)[index84_] = d_545_v10_
            d_546_v11_: _dafny.Map
            d_546_v11_ = _dafny.Map({d_538_i0_: self.f14})
            index85_ = default__.safeIndex(770, (d_539_v4_).length(0))
            rhs81_ = d_545_v10_
            rhs82_ = self.f14
            rhs83_ = (not(self.f14)) == (((d_546_v11_)[d_538_i0_] if (d_538_i0_) in (d_546_v11_) else self.f14))
            lhs59_ = d_539_v4_
            lhs60_ = default__.safeIndex(770, (d_539_v4_).length(0))
            lhs61_ = self
            lhs62_ = self
            lhs59_[lhs60_] = rhs81_
            lhs61_.f14 = rhs82_
            lhs62_.f14 = rhs83_
            d_547_v12_: _dafny.Set
            d_547_v12_ = _dafny.Set({len(_dafny.Map({len(p0): self.f14})), len(p0)})
            d_548_v13_: _dafny.Map
            d_548_v13_ = _dafny.Map({d_547_v12_: d_534_v0_})
            d_549_v14_: _dafny.MultiSet
            d_549_v14_ = _dafny.MultiSet([(d_538_i0_) >= (len(d_548_v13_))])
            d_549_v14_ = d_549_v14_
            (globalState).f2 = d_538_i0_
            (self).f14 = (self.f14 if self.f14 else self.f14)
        d_550_v15_: _dafny.Map
        d_550_v15_ = _dafny.Map({d_534_v0_: self.f14})
        d_550_v15_ = (d_550_v15_).set((len(d_535_v1_)) + (d_534_v0_), self.f14)
        d_551_v16_: _dafny.Array
        nw90_ = _dafny.Array(_dafny.Map({}), 9)
        d_551_v16_ = nw90_
        d_552_v17_: _dafny.Map
        d_552_v17_ = _dafny.Map({self.f14: not(self.f14)})
        index86_ = default__.safeIndex(345, (d_551_v16_).length(0))
        (d_551_v16_)[index86_] = d_552_v17_
        index87_ = default__.safeIndex(345, (d_551_v16_).length(0))
        (d_551_v16_)[index87_] = _dafny.Map({self.f14: self.f14})
        d_553_v18_: _dafny.Seq
        d_553_v18_ = _dafny.SeqWithoutIsStrInference([not(self.f14), self.f14, self.f14])
        d_554_v19_: _dafny.Array
        def lambda22_(d_555_i1_):
            return self.f14

        init13_ = lambda22_
        nw91_ = _dafny.Array(None, 11)
        for i0_13_ in range(nw91_.length(0)):
            nw91_[i0_13_] = init13_(i0_13_)
        d_554_v19_ = nw91_
        d_556_v20_: D3
        d_556_v20_ = D3_DC10(True, d_534_v0_, self.f14, d_554_v19_, d_550_v15_)
        pat_let_tv8_ = d_550_v15_
        pat_let_tv9_ = d_556_v20_
        pat_let_tv10_ = globalState
        def iife35_(_pat_let10_0):
            def iife36_(d_557_dt__update__tmp_h0_):
                def iife37_(_pat_let11_0):
                    def iife38_(d_558_dt__update_hcf5_h0_):
                        return D1_DC3(d_558_dt__update_hcf5_h0_)
                    return iife38_(_pat_let11_0)
                return iife37_(default__.fm6(len(pat_let_tv8_), (pat_let_tv9_).cf14, not(self.f14), _dafny.CodePoint('a'), pat_let_tv10_))
            return iife36_(_pat_let10_0)
        source11_ = iife35_(default__.fm5(d_534_v0_, d_534_v0_, d_534_v0_, d_553_v18_, globalState))
        if source11_.is_DC4:
            d_559___mcc_h0_ = source11_.cf6
            d_560___mcc_h1_ = source11_.cf7
            d_561_cf7_ = d_560___mcc_h1_
            d_562_cf6_ = d_559___mcc_h0_
            d_563_v21_: C0
            nw92_ = C0()
            nw92_.ctor__()
            d_563_v21_ = nw92_
            d_564_v22_: _dafny.Seq
            d_564_v22_ = _dafny.SeqWithoutIsStrInference([d_554_v19_])
            d_565_v23_: _dafny.Array
            nw93_ = _dafny.Array(None, 23)
            nw93_[int(0)] = d_554_v19_
            nw93_[int(1)] = (d_554_v19_ if self.f14 else d_554_v19_)
            nw93_[int(2)] = d_554_v19_
            nw93_[int(3)] = d_554_v19_
            nw93_[int(4)] = d_554_v19_
            nw93_[int(5)] = d_554_v19_
            nw93_[int(6)] = d_554_v19_
            nw93_[int(7)] = (d_556_v20_).cf17
            nw93_[int(8)] = d_554_v19_
            nw93_[int(9)] = d_554_v19_
            nw93_[int(10)] = d_554_v19_
            nw93_[int(11)] = d_554_v19_
            nw93_[int(12)] = (d_564_v22_)[default__.safeIndex(d_561_cf7_, len(d_564_v22_))]
            nw93_[int(13)] = d_554_v19_
            nw93_[int(14)] = d_554_v19_
            nw93_[int(15)] = d_554_v19_
            nw93_[int(16)] = d_554_v19_
            nw93_[int(17)] = d_554_v19_
            nw93_[int(18)] = d_554_v19_
            nw93_[int(19)] = d_554_v19_
            nw93_[int(20)] = d_554_v19_
            nw93_[int(21)] = d_554_v19_
            nw93_[int(22)] = d_554_v19_
            d_565_v23_ = nw93_
            index88_ = default__.safeIndex(347, (d_565_v23_).length(0))
            (d_565_v23_)[index88_] = d_554_v19_
            index89_ = default__.safeIndex(347, (d_565_v23_).length(0))
            (d_565_v23_)[index89_] = d_554_v19_
            d_566_v24_: _dafny.Map
            d_566_v24_ = _dafny.Map({default__.fm0(self.f14, self.f14, self.f14, self.f14, globalState): d_554_v19_})
            d_554_v19_ = ((d_566_v24_)[self.f14] if (self.f14) in (d_566_v24_) else d_554_v19_)
            d_537_v3_ = d_537_v3_
        elif source11_.is_DC5:
            d_567___mcc_h2_ = source11_.cf8
            d_568___mcc_h3_ = source11_.cf9
            d_569_cf9_ = d_568___mcc_h3_
            d_570_cf8_ = d_567___mcc_h2_
            d_571_v25_: _dafny.Map
            d_571_v25_ = _dafny.Map({d_534_v0_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vsixj"))})
            d_572_v26_: _dafny.Seq
            d_572_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oukgvej"))
            d_573_v27_: _dafny.Map
            d_573_v27_ = _dafny.Map({d_534_v0_: (883 if d_570_cf8_ else len(((d_571_v25_)[d_534_v0_] if (d_534_v0_) in (d_571_v25_) else d_572_v26_)))})
            d_573_v27_ = (d_573_v27_).set(d_534_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lysyusnd"))))
            d_574_v28_: _dafny.Map
            d_574_v28_ = _dafny.Map({(False if d_570_cf8_ else self.f14): d_534_v0_})
            d_574_v28_ = (d_574_v28_).set((self.f14) or (d_570_cf8_), 177)
            d_575_v29_: _dafny.Array
            nw94_ = _dafny.Array(_dafny.CodePoint('D'), 15)
            d_575_v29_ = nw94_
            index90_ = default__.safeIndex(725, (d_575_v29_).length(0))
            (d_575_v29_)[index90_] = d_536_v2_
            index91_ = default__.safeIndex(725, (d_575_v29_).length(0))
            (d_575_v29_)[index91_] = d_536_v2_
            d_576_v30_: C0
            nw95_ = C0()
            nw95_.ctor__()
            d_576_v30_ = nw95_
        elif source11_.is_DC3:
            d_577___mcc_h4_ = source11_.cf5
            d_578_cf5_ = d_577___mcc_h4_
            d_579_v31_: C0
            nw96_ = C0()
            nw96_.ctor__()
            d_579_v31_ = nw96_
            (self).f14 = (d_553_v18_)[default__.safeIndex(((d_537_v3_)[len((d_551_v16_)[default__.safeIndex(345, (d_551_v16_).length(0))])] if (len((d_551_v16_)[default__.safeIndex(345, (d_551_v16_).length(0))])) in (d_537_v3_) else d_534_v0_), len(d_553_v18_))]
            index92_ = default__.safeIndex(870, (d_554_v19_).length(0))
            (d_554_v19_)[index92_] = False
            d_580_v32_: _dafny.Map
            d_580_v32_ = _dafny.Map({d_554_v19_: d_534_v0_})
            d_581_v33_: _dafny.Seq
            d_581_v33_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_554_v19_: d_534_v0_}), _dafny.Map({d_554_v19_: (0) - (d_534_v0_)})])
            index93_ = default__.safeIndex(870, (d_554_v19_).length(0))
            rhs84_ = self.f14
            rhs85_ = ((d_580_v32_ if False else d_580_v32_)) | ((d_581_v33_)[default__.safeIndex((0) - (d_534_v0_), len(d_581_v33_))])
            rhs86_ = (self.f14) == (True)
            rhs87_ = self.f14
            rhs88_ = default__.fm9(self.f14, globalState)
            lhs63_ = d_554_v19_
            lhs64_ = default__.safeIndex(870, (d_554_v19_).length(0))
            lhs65_ = self
            lhs66_ = self
            lhs63_[lhs64_] = rhs84_
            d_580_v32_ = rhs85_
            lhs65_.f14 = rhs86_
            lhs66_.f14 = rhs87_
            r0 = rhs88_
            d_582_v34_: _dafny.Seq
            d_582_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihdetle"))
            d_583_v35_: _dafny.Set
            d_583_v35_ = _dafny.Set({(0) - (d_534_v0_), d_534_v0_})
            d_584_v36_: _dafny.MultiSet
            d_584_v36_ = _dafny.MultiSet([d_582_v34_])
            d_585_v37_: D1
            d_585_v37_ = D1_DC4(d_584_v36_, d_534_v0_)
            d_586_v38_: _dafny.Map
            d_586_v38_ = _dafny.Map({(d_554_v19_)[default__.safeIndex(870, (d_554_v19_).length(0))]: d_553_v18_})
            d_587_v39_: _dafny.Array
            nw97_ = _dafny.Array(None, 17)
            nw97_[int(0)] = d_534_v0_
            nw97_[int(1)] = d_534_v0_
            nw97_[int(2)] = d_534_v0_
            nw97_[int(3)] = d_534_v0_
            nw97_[int(4)] = 137
            nw97_[int(5)] = default__.safeDivisionInt(d_534_v0_, d_534_v0_)
            nw97_[int(6)] = (len(d_582_v34_)) + (len(d_583_v35_))
            nw97_[int(7)] = d_534_v0_
            nw97_[int(8)] = (d_585_v37_).cf7
            nw97_[int(9)] = (d_534_v0_) + (len(d_586_v38_))
            nw97_[int(10)] = ((d_537_v3_)[d_534_v0_] if (d_534_v0_) in (d_537_v3_) else 745)
            nw97_[int(11)] = d_534_v0_
            nw97_[int(12)] = (d_534_v0_) + (d_534_v0_)
            nw97_[int(13)] = default__.fm4(globalState)
            nw97_[int(14)] = default__.safeModuloInt(386, d_534_v0_)
            nw97_[int(15)] = d_534_v0_
            nw97_[int(16)] = (d_535_v1_)[default__.safeIndex(len(d_583_v35_), len(d_535_v1_))]
            d_587_v39_ = nw97_
            index94_ = default__.safeIndex(802, (d_587_v39_).length(0))
            (d_587_v39_)[index94_] = d_534_v0_
            index95_ = default__.safeIndex(802, (d_587_v39_).length(0))
            (d_587_v39_)[index95_] = d_534_v0_
        elif True:
            d_588___mcc_h5_ = source11_.cf10
            d_589_cf10_ = d_588___mcc_h5_
            d_590_v40_: _dafny.Array
            def lambda23_(d_591_v2_):
                def lambda24_(d_592_i2_):
                    return D1_DC6(D1_DC3(d_591_v2_))

                return lambda24_

            init14_ = lambda23_(d_536_v2_)
            nw98_ = _dafny.Array(None, 8)
            for i0_14_ in range(nw98_.length(0)):
                nw98_[i0_14_] = init14_(i0_14_)
            d_590_v40_ = nw98_
            d_593_v41_: D2
            d_593_v41_ = D2_DC7(d_590_v40_)
            d_593_v41_ = d_593_v41_
            (self).f14 = (self.f14 if self.f14 else self.f14)
            d_594_v42_: D2
            d_594_v42_ = D2_DC8(675)
            d_595_v43_: _dafny.MultiSet
            d_595_v43_ = _dafny.MultiSet([d_594_v42_])
            (globalState).f2 = (((_dafny.MultiSet([d_594_v42_])) - (_dafny.MultiSet([d_594_v42_, d_594_v42_])) if self.f14 else ((_dafny.MultiSet([d_594_v42_])).set(d_594_v42_, default__.abs(d_534_v0_))) | (d_595_v43_))).cardinality
            d_596_v44_: C0
            nw99_ = C0()
            nw99_.ctor__()
            d_596_v44_ = nw99_
        d_597_v45_: _dafny.Seq
        d_597_v45_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nwhkcx"))
        d_598_v46_: _dafny.Seq
        d_598_v46_ = _dafny.SeqWithoutIsStrInference([d_597_v45_, (d_597_v45_).set(default__.safeIndex((0) - (d_534_v0_), len(d_597_v45_)), d_536_v2_), d_597_v45_, d_597_v45_])
        r0 = (d_598_v46_)[default__.safeIndex((len(d_535_v1_)) * (d_534_v0_), len(d_598_v46_))]
        d_599_v47_: D1
        d_599_v47_ = D1_DC5(self.f14, (d_597_v45_) < (d_597_v45_))
        r1 = d_599_v47_
        d_600_v48_: _dafny.Seq
        d_600_v48_ = _dafny.SeqWithoutIsStrInference([d_554_v19_, d_554_v19_, d_554_v19_, d_554_v19_])
        r2 = default__.safeModuloInt(len(d_600_v48_), (d_534_v0_) * (len(d_597_v45_)))
        return r0, r1, r2


class C2:
    def  __init__(self):
        self.f13: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f13):
        (self).f13 = f13

    def m3(self, p0, p1, p2, globalState):
        hi2_ = self.f13
        for d_601_i0_ in range(self.f13, hi2_):
            d_602_v0_: bool
            d_602_v0_ = True
            d_603_v1_: str
            d_603_v1_ = _dafny.CodePoint('e')
            d_604_v2_: _dafny.Map
            d_604_v2_ = _dafny.Map({d_603_v1_: d_602_v0_})
            d_605_v3_: _dafny.Map
            d_605_v3_ = _dafny.Map({d_601_i0_: d_602_v0_})
            d_606_v4_: _dafny.Map
            d_606_v4_ = _dafny.Map({len(d_605_v3_): self.f13})
            (self).f13 = ((self.f13 if d_602_v0_ else len(d_604_v2_))) + ((d_601_i0_) - (len(d_606_v4_)))
            d_607_v5_: _dafny.Array
            def lambda25_(d_608_v0_):
                def lambda26_(d_609_i1_):
                    return _dafny.SeqWithoutIsStrInference([d_608_v0_])

                return lambda26_

            init15_ = lambda25_(d_602_v0_)
            nw100_ = _dafny.Array(None, 21)
            for i0_15_ in range(nw100_.length(0)):
                nw100_[i0_15_] = init15_(i0_15_)
            d_607_v5_ = nw100_
            d_607_v5_ = d_607_v5_
            d_610_v6_: _dafny.Seq
            d_610_v6_ = _dafny.SeqWithoutIsStrInference([False])
            d_611_v7_: _dafny.Map
            d_611_v7_ = _dafny.Map({default__.fm0(True, not(d_602_v0_), d_602_v0_, (d_610_v6_)[default__.safeIndex(self.f13, len(d_610_v6_))], globalState): d_602_v0_})
            d_612_v8_: _dafny.Seq
            d_612_v8_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(p1): d_601_i0_})])
            d_613_v9_: _dafny.Array
            nw101_ = _dafny.Array(None, 29)
            nw101_[int(0)] = len(p1)
            nw101_[int(1)] = (p2)[default__.safeIndex(d_601_i0_, len(p2))]
            nw101_[int(2)] = 908
            nw101_[int(3)] = d_601_i0_
            nw101_[int(4)] = len(d_611_v7_)
            nw101_[int(5)] = (self.f13) - (default__.fm2(d_602_v0_, globalState))
            nw101_[int(6)] = d_601_i0_
            nw101_[int(7)] = self.f13
            nw101_[int(8)] = (0) - (d_601_i0_)
            nw101_[int(9)] = 943
            nw101_[int(10)] = self.f13
            nw101_[int(11)] = self.f13
            nw101_[int(12)] = len(p1)
            nw101_[int(13)] = 917
            nw101_[int(14)] = len((d_612_v8_) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_601_i0_: len(_dafny.Set({self.f13}))}) for d_614_i2_ in range(default__.abs(681))])))
            nw101_[int(15)] = (-545) * (d_601_i0_)
            nw101_[int(16)] = d_601_i0_
            nw101_[int(17)] = d_601_i0_
            nw101_[int(18)] = self.f13
            nw101_[int(19)] = self.f13
            nw101_[int(20)] = d_601_i0_
            nw101_[int(21)] = self.f13
            nw101_[int(22)] = (self.f13) - (self.f13)
            nw101_[int(23)] = self.f13
            nw101_[int(24)] = self.f13
            nw101_[int(25)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngpktyvor"))) + (p1))
            nw101_[int(26)] = d_601_i0_
            nw101_[int(27)] = self.f13
            nw101_[int(28)] = len(_dafny.SeqWithoutIsStrInference([d_602_v0_]))
            d_613_v9_ = nw101_
            index96_ = default__.safeIndex(171, (d_613_v9_).length(0))
            (d_613_v9_)[index96_] = len(_dafny.SeqWithoutIsStrInference([d_603_v1_ for d_615_i3_ in range(default__.abs(-155))]))
            d_616_v10_: D0
            d_616_v10_ = D0_DC0(d_602_v0_)
            d_617_v11_: _dafny.Set
            d_617_v11_ = _dafny.Set({d_616_v10_})
            d_618_v12_: _dafny.Map
            d_618_v12_ = _dafny.Map({d_602_v0_: d_617_v11_})
            d_619_v13_: _dafny.Map
            d_619_v13_ = _dafny.Map({d_616_v10_: 680})
            d_620_v15_: _dafny.Map
            def iife39_():
                coll15_ = _dafny.Set()
                compr_15_: D0
                for compr_15_ in (d_619_v13_).keys.Elements:
                    d_621_v14_: D0 = compr_15_
                    if (d_621_v14_) in (d_619_v13_):
                        coll15_ = coll15_.union(_dafny.Set([d_621_v14_]))
                return _dafny.Set(coll15_)
            d_620_v15_ = _dafny.Map({False: ((d_618_v12_)[False] if (False) in (d_618_v12_) else iife39_()
            )})
            d_622_v16_: _dafny.Set
            d_622_v16_ = _dafny.Set({d_602_v0_, d_602_v0_, not(False), d_602_v0_, d_602_v0_})
            index97_ = default__.safeIndex(171, (d_613_v9_).length(0))
            rhs89_ = len(((d_618_v12_)[d_602_v0_] if (d_602_v0_) in (d_618_v12_) else (d_617_v11_ if d_602_v0_ else d_617_v11_)))
            rhs90_ = (_dafny.Set({d_602_v0_, d_602_v0_})).ispropersubset(d_622_v16_)
            lhs67_ = d_613_v9_
            lhs68_ = default__.safeIndex(171, (d_613_v9_).length(0))
            lhs67_[lhs68_] = rhs89_
            d_602_v0_ = rhs90_
            d_602_v0_ = d_602_v0_
        d_623_v17_: _dafny.Seq
        d_623_v17_ = _dafny.SeqWithoutIsStrInference([p0])
        (self).f13 = len(d_623_v17_)
        d_624_v18_: bool
        d_624_v18_ = True
        rhs91_ = (self.f13) != (self.f13)
        rhs92_ = (default__.safeDivisionInt(self.f13, self.f13)) > (self.f13)
        d_624_v18_ = rhs91_
        d_624_v18_ = rhs92_
        d_624_v18_ = False
        hi3_ = default__.safeModuloInt(self.f13, self.f13)
        for d_625_i4_ in range(self.f13, hi3_):
            d_626_v19_: _dafny.Set
            d_626_v19_ = _dafny.Set({d_624_v18_, False, True, default__.fm0(d_624_v18_, d_624_v18_, d_624_v18_, d_624_v18_, globalState)})
            if not(not((d_626_v19_).issubset(d_626_v19_))):
                d_627_v20_: _dafny.Array
                nw102_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_627_v20_ = nw102_
                d_628_v21_: str
                d_628_v21_ = _dafny.CodePoint('l')
                d_629_v22_: _dafny.Map
                d_629_v22_ = _dafny.Map({d_628_v21_: p1})
                d_630_v23_: D1
                d_630_v23_ = D1_DC3(d_628_v21_)
                index98_ = default__.safeIndex(281, (d_627_v20_).length(0))
                (d_627_v20_)[index98_] = ((d_629_v22_)[(d_630_v23_).cf5] if ((d_630_v23_).cf5) in (d_629_v22_) else p1)
                index99_ = default__.safeIndex(281, (d_627_v20_).length(0))
                (d_627_v20_)[index99_] = default__.fm3(d_624_v18_, p1, self.f13, globalState)
                d_631_v24_: _dafny.Array
                nw103_ = _dafny.Array(int(0), 20)
                d_631_v24_ = nw103_
                index100_ = default__.safeIndex(677, (d_631_v24_).length(0))
                (d_631_v24_)[index100_] = (0) - (d_625_i4_)
                index101_ = default__.safeIndex(677, (d_631_v24_).length(0))
                (d_631_v24_)[index101_] = d_625_i4_
                d_632_v25_: _dafny.Seq
                d_632_v25_ = _dafny.SeqWithoutIsStrInference([p2])
                d_632_v25_ = d_632_v25_
                d_633_v26_: _dafny.Map
                d_633_v26_ = _dafny.Map({not(d_624_v18_): d_624_v18_})
                d_634_v27_: _dafny.Array
                nw104_ = _dafny.Array(None, 21)
                nw104_[int(0)] = (default__.fm0(d_624_v18_, d_624_v18_, d_624_v18_, not(d_624_v18_), globalState) if d_624_v18_ else d_624_v18_)
                nw104_[int(1)] = (d_626_v19_).ispropersubset(d_626_v19_)
                nw104_[int(2)] = not(d_624_v18_)
                nw104_[int(3)] = d_624_v18_
                nw104_[int(4)] = d_624_v18_
                nw104_[int(5)] = d_624_v18_
                nw104_[int(6)] = default__.fm0(d_624_v18_, d_624_v18_, d_624_v18_, d_624_v18_, globalState)
                nw104_[int(7)] = (d_624_v18_ if d_624_v18_ else d_624_v18_)
                nw104_[int(8)] = default__.fm0(d_624_v18_, False, d_624_v18_, False, globalState)
                nw104_[int(9)] = d_624_v18_
                nw104_[int(10)] = d_624_v18_
                nw104_[int(11)] = ((d_633_v26_)[d_624_v18_] if (d_624_v18_) in (d_633_v26_) else d_624_v18_)
                nw104_[int(12)] = ((d_633_v26_)[d_624_v18_] if (d_624_v18_) in (d_633_v26_) else d_624_v18_)
                nw104_[int(13)] = d_624_v18_
                nw104_[int(14)] = d_624_v18_
                nw104_[int(15)] = d_624_v18_
                nw104_[int(16)] = d_624_v18_
                nw104_[int(17)] = d_624_v18_
                nw104_[int(18)] = not(d_624_v18_)
                nw104_[int(19)] = d_624_v18_
                nw104_[int(20)] = False
                d_634_v27_ = nw104_
                nw105_ = _dafny.Array(False, 17)
                d_634_v27_ = nw105_
                d_635_v28_: D1
                d_635_v28_ = D1_DC5(d_624_v18_, d_624_v18_)
                d_636_v29_: _dafny.Seq
                d_636_v29_ = _dafny.SeqWithoutIsStrInference([d_635_v28_])
                d_637_v30_: _dafny.Array
                nw106_ = _dafny.Array(D1.default()(), 5)
                d_637_v30_ = nw106_
                d_638_v31_: D2
                d_638_v31_ = D2_DC7(d_637_v30_)
                d_639_v32_: _dafny.MultiSet
                d_639_v32_ = _dafny.MultiSet([(d_627_v20_)[default__.safeIndex(281, (d_627_v20_).length(0))], p1])
                d_640_v33_: D1
                d_640_v33_ = D1_DC4(d_639_v32_, d_625_i4_)
                d_641_v34_: _dafny.Map
                d_641_v34_ = _dafny.Map({d_640_v33_: d_637_v30_})
                d_642_v35_: _dafny.Array
                nw107_ = _dafny.Array(None, 29)
                nw107_[int(0)] = d_637_v30_
                nw107_[int(1)] = d_637_v30_
                nw107_[int(2)] = d_637_v30_
                nw107_[int(3)] = d_637_v30_
                nw107_[int(4)] = d_637_v30_
                nw107_[int(5)] = d_637_v30_
                nw107_[int(6)] = d_637_v30_
                nw107_[int(7)] = d_637_v30_
                nw107_[int(8)] = d_637_v30_
                nw107_[int(9)] = d_637_v30_
                nw107_[int(10)] = d_637_v30_
                nw107_[int(11)] = d_637_v30_
                nw107_[int(12)] = d_637_v30_
                nw107_[int(13)] = d_637_v30_
                nw107_[int(14)] = (d_637_v30_ if not(d_624_v18_) else d_637_v30_)
                nw107_[int(15)] = (d_638_v31_).cf11
                nw107_[int(16)] = d_637_v30_
                nw107_[int(17)] = ((d_641_v34_)[d_640_v33_] if (d_640_v33_) in (d_641_v34_) else d_637_v30_)
                nw107_[int(18)] = d_637_v30_
                nw107_[int(19)] = d_637_v30_
                nw107_[int(20)] = d_637_v30_
                nw107_[int(21)] = d_637_v30_
                nw107_[int(22)] = d_637_v30_
                nw107_[int(23)] = d_637_v30_
                nw107_[int(24)] = d_637_v30_
                nw107_[int(25)] = d_637_v30_
                nw107_[int(26)] = d_637_v30_
                nw107_[int(27)] = (d_637_v30_ if d_624_v18_ else d_637_v30_)
                nw107_[int(28)] = d_637_v30_
                d_642_v35_ = nw107_
                rhs93_ = (_dafny.SeqWithoutIsStrInference([D1_DC5(d_624_v18_, d_624_v18_)])).set(default__.safeIndex(default__.safeDivisionInt(620, self.f13), len(_dafny.SeqWithoutIsStrInference([D1_DC5(d_624_v18_, d_624_v18_)]))), d_635_v28_)
                rhs94_ = not(d_624_v18_)
                rhs95_ = d_642_v35_
                d_636_v29_ = rhs93_
                d_624_v18_ = rhs94_
                d_642_v35_ = rhs95_
            elif True:
                d_643_v36_: _dafny.Map
                d_643_v36_ = _dafny.Map({223: d_624_v18_})
                d_644_v37_: D3
                d_644_v37_ = D3_DC10(d_624_v18_, self.f13, d_624_v18_, p0, d_643_v36_)
                d_645_v38_: str
                d_645_v38_ = _dafny.CodePoint('i')
                d_646_v39_: _dafny.Map
                d_646_v39_ = _dafny.Map({(d_644_v37_).cf17: d_645_v38_})
                d_646_v39_ = (d_646_v39_).set(p0, (d_645_v38_ if d_624_v18_ else d_645_v38_))
                d_647_v40_: _dafny.Array
                nw108_ = _dafny.Array(int(0), 1)
                d_647_v40_ = nw108_
                index102_ = default__.safeIndex(425, (d_647_v40_).length(0))
                (d_647_v40_)[index102_] = d_625_i4_
                index103_ = default__.safeIndex(425, (d_647_v40_).length(0))
                rhs96_ = self.f13
                rhs97_ = d_647_v40_
                rhs98_ = self.f13
                lhs69_ = self
                lhs70_ = globalState
                lhs71_ = d_647_v40_
                lhs72_ = default__.safeIndex(425, (d_647_v40_).length(0))
                lhs69_.f13 = rhs96_
                lhs70_.f10 = rhs97_
                lhs71_[lhs72_] = rhs98_
                d_648_v41_: _dafny.Array
                def lambda27_(d_649_v18_):
                    def lambda28_(d_650_i5_):
                        return d_649_v18_

                    return lambda28_

                init16_ = lambda27_(d_624_v18_)
                nw109_ = _dafny.Array(None, 15)
                for i0_16_ in range(nw109_.length(0)):
                    nw109_[i0_16_] = init16_(i0_16_)
                d_648_v41_ = nw109_
                d_648_v41_ = d_648_v41_
                index104_ = default__.safeIndex(425, (d_647_v40_).length(0))
                rhs99_ = d_624_v18_
                rhs100_ = (d_625_i4_) * (default__.fm2(d_624_v18_, globalState))
                rhs101_ = (self.f13) + (((d_647_v40_)[default__.safeIndex(425, (d_647_v40_).length(0))]) * (self.f13))
                rhs102_ = d_625_i4_
                rhs103_ = self.f13
                lhs73_ = self
                lhs74_ = globalState
                lhs75_ = globalState
                lhs76_ = d_647_v40_
                lhs77_ = default__.safeIndex(425, (d_647_v40_).length(0))
                d_624_v18_ = rhs99_
                lhs73_.f13 = rhs100_
                lhs74_.f2 = rhs101_
                lhs75_.f2 = rhs102_
                lhs76_[lhs77_] = rhs103_
                d_651_v42_: _dafny.Map
                d_651_v42_ = _dafny.Map({default__.safeDivisionInt(458, self.f13): (len(p2)) + (d_625_i4_)})
                d_651_v42_ = d_651_v42_
            d_652_v43_: _dafny.Array
            nw110_ = _dafny.Array(_dafny.Map({}), 13)
            d_652_v43_ = nw110_
            d_652_v43_ = d_652_v43_
            d_653_v44_: str
            d_653_v44_ = _dafny.CodePoint('a')
            d_654_v45_: _dafny.Set
            d_654_v45_ = _dafny.Set({d_653_v44_})
            d_655_v46_: _dafny.Map
            d_655_v46_ = _dafny.Map({len(p1): d_624_v18_})
            d_656_v47_: D3
            d_656_v47_ = D3_DC10(d_624_v18_, len(d_654_v45_), d_624_v18_, (d_623_v17_)[default__.safeIndex(self.f13, len(d_623_v17_))], d_655_v46_)
            if (d_656_v47_).cf16:
                d_657_v48_: _dafny.Map
                d_657_v48_ = _dafny.Map({d_624_v18_: (d_624_v18_) or (d_624_v18_)})
                d_624_v18_ = ((d_657_v48_)[not(not (d_624_v18_) or (d_624_v18_))] if (not(not (d_624_v18_) or (d_624_v18_))) in (d_657_v48_) else False)
                (self).f13 = self.f13
                d_658_v49_: _dafny.Array
                nw111_ = _dafny.Array(False, 17)
                d_658_v49_ = nw111_
                nw112_ = _dafny.Array(False, 3)
                d_658_v49_ = nw112_
                d_659_v50_: _dafny.MultiSet
                d_659_v50_ = _dafny.MultiSet([d_625_i4_])
                rhs104_ = self.f13
                rhs105_ = self.f13
                rhs106_ = (((d_659_v50_)[default__.fm2(True, globalState)] if (default__.fm2(True, globalState)) in (d_659_v50_) else d_625_i4_)) in ((_dafny.Map({(0) - ((0) - (d_625_i4_)): d_624_v18_})) | ((d_655_v46_).set(self.f13, d_624_v18_)))
                rhs107_ = _dafny.CodePoint('c')
                rhs108_ = d_624_v18_
                lhs78_ = self
                lhs79_ = globalState
                lhs78_.f13 = rhs104_
                lhs79_.f2 = rhs105_
                d_624_v18_ = rhs106_
                d_653_v44_ = rhs107_
                d_624_v18_ = rhs108_
                d_660_v51_: _dafny.Array
                nw113_ = _dafny.Array(_dafny.Map({}), 13)
                d_660_v51_ = nw113_
                d_661_v52_: _dafny.MultiSet
                d_661_v52_ = _dafny.MultiSet([d_624_v18_])
                default__.m0(p0, self.f13, d_660_v51_, (d_661_v52_).cardinality, globalState)
            elif True:
                index105_ = default__.safeIndex(573, (p0).length(0))
                (p0)[index105_] = d_624_v18_
                d_662_v53_: _dafny.Array
                nw114_ = _dafny.Array(_dafny.Map({}), 29)
                d_662_v53_ = nw114_
                d_663_v54_: _dafny.Map
                d_663_v54_ = _dafny.Map({p1: True})
                index106_ = default__.safeIndex(729, (d_662_v53_).length(0))
                (d_662_v53_)[index106_] = d_663_v54_
                index107_ = default__.safeIndex(573, (p0).length(0))
                index108_ = default__.safeIndex(729, (d_662_v53_).length(0))
                rhs109_ = d_624_v18_
                rhs110_ = d_663_v54_
                lhs80_ = p0
                lhs81_ = default__.safeIndex(573, (p0).length(0))
                lhs82_ = d_662_v53_
                lhs83_ = default__.safeIndex(729, (d_662_v53_).length(0))
                lhs80_[lhs81_] = rhs109_
                lhs82_[lhs83_] = rhs110_
                d_664_v55_: _dafny.Array
                nw115_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_664_v55_ = nw115_
                d_665_v56_: _dafny.Array
                nw116_ = _dafny.Array(False, 18)
                d_665_v56_ = nw116_
                index109_ = default__.safeIndex(897, (d_664_v55_).length(0))
                (d_664_v55_)[index109_] = d_665_v56_
                index110_ = default__.safeIndex(897, (d_664_v55_).length(0))
                (d_664_v55_)[index110_] = d_665_v56_
                d_666_v57_: _dafny.MultiSet
                d_666_v57_ = _dafny.MultiSet([(p0)[default__.safeIndex(573, (p0).length(0))], True, not((p0)[default__.safeIndex(573, (p0).length(0))])])
                d_666_v57_ = d_666_v57_
                d_667_v58_: C1
                nw117_ = C1()
                nw117_.ctor__(d_624_v18_)
                d_667_v58_ = nw117_
                d_668_v59_: _dafny.Seq
                d_668_v59_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_653_v44_ for d_669_i6_ in range(default__.abs(-930))]), p1])
                d_670_v60_: _dafny.Map
                d_670_v60_ = _dafny.Map({(p0)[default__.safeIndex(573, (p0).length(0))]: True})
                d_671_v61_: _dafny.Map
                d_671_v61_ = _dafny.Map({d_667_v58_.f14: len(d_670_v60_)})
                d_672_v62_: _dafny.Set
                d_672_v62_ = _dafny.Set({len(p2), ((d_671_v61_)[default__.fm0(d_624_v18_, d_667_v58_.f14, d_667_v58_.f14, (p0)[default__.safeIndex(573, (p0).length(0))], globalState)] if (default__.fm0(d_624_v18_, d_667_v58_.f14, d_667_v58_.f14, (p0)[default__.safeIndex(573, (p0).length(0))], globalState)) in (d_671_v61_) else d_625_i4_), self.f13})
                d_673_v63_: _dafny.Array
                def lambda29_(d_674_i4_):
                    def lambda30_(d_675_i7_):
                        return (d_675_i7_) - (d_674_i4_)

                    return lambda30_

                init17_ = lambda29_(d_625_i4_)
                nw118_ = _dafny.Array(None, 5)
                for i0_17_ in range(nw118_.length(0)):
                    nw118_[i0_17_] = init17_(i0_17_)
                d_673_v63_ = nw118_
                d_676_v64_: _dafny.Seq
                d_676_v64_ = _dafny.SeqWithoutIsStrInference([d_673_v63_, d_673_v63_, d_673_v63_, d_673_v63_])
                index111_ = default__.safeIndex(573, (p0).length(0))
                rhs111_ = (d_626_v19_).isdisjoint(_dafny.Set({d_624_v18_}))
                rhs112_ = (p1)[default__.safeIndex(len(d_672_v62_), len(p1))]
                rhs113_ = len((d_676_v64_) + ((_dafny.SeqWithoutIsStrInference([d_673_v63_, d_673_v63_])) + (d_676_v64_)))
                rhs114_ = ((d_668_v59_).set(default__.safeIndex(d_625_i4_, len(d_668_v59_)), p1)).set(default__.safeIndex((0) - (self.f13), len((d_668_v59_).set(default__.safeIndex(d_625_i4_, len(d_668_v59_)), p1))), (p1) + (p1))
                rhs115_ = False
                lhs84_ = self
                lhs85_ = p0
                lhs86_ = default__.safeIndex(573, (p0).length(0))
                d_624_v18_ = rhs111_
                d_653_v44_ = rhs112_
                lhs84_.f13 = rhs113_
                d_668_v59_ = rhs114_
                lhs85_[lhs86_] = rhs115_
            d_677_v65_: _dafny.Map
            d_677_v65_ = _dafny.Map({d_624_v18_: d_626_v19_})
            d_678_v66_: D4
            d_678_v66_ = D4_DC12(d_626_v19_)
            d_679_v67_: _dafny.Map
            d_679_v67_ = _dafny.Map({d_625_i4_: (d_678_v66_).cf20})
            d_680_v68_: _dafny.MultiSet
            d_680_v68_ = _dafny.MultiSet([self.f13])
            d_681_v69_: _dafny.Seq
            d_681_v69_ = _dafny.SeqWithoutIsStrInference([d_624_v18_, False])
            d_682_v70_: _dafny.Map
            d_682_v70_ = _dafny.Map({d_624_v18_: (d_681_v69_)[default__.safeIndex(d_625_i4_, len(d_681_v69_))]})
            d_683_v71_: _dafny.Map
            d_683_v71_ = _dafny.Map({not(d_624_v18_): d_653_v44_})
            d_684_v72_: _dafny.Array
            nw119_ = _dafny.Array(None, 29)
            nw119_[int(0)] = 647
            nw119_[int(1)] = len(((p1 if d_624_v18_ else p1)).set(default__.safeIndex(self.f13, len((p1 if d_624_v18_ else p1))), d_653_v44_))
            nw119_[int(2)] = default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_685_i8_ in range(default__.abs(466))])), (p2)[default__.safeIndex(d_625_i4_, len(p2))])
            nw119_[int(3)] = (self.f13) - (len(d_626_v19_))
            nw119_[int(4)] = len(((d_677_v65_)[d_624_v18_] if (d_624_v18_) in (d_677_v65_) else ((d_679_v67_)[self.f13] if (self.f13) in (d_679_v67_) else d_626_v19_)))
            nw119_[int(5)] = d_625_i4_
            nw119_[int(6)] = default__.safeModuloInt(self.f13, 182)
            nw119_[int(7)] = default__.safeModuloInt(self.f13, self.f13)
            nw119_[int(8)] = d_625_i4_
            nw119_[int(9)] = self.f13
            nw119_[int(10)] = len(p1)
            nw119_[int(11)] = ((d_680_v68_)[len(d_681_v69_)] if (len(d_681_v69_)) in (d_680_v68_) else d_625_i4_)
            nw119_[int(12)] = d_625_i4_
            nw119_[int(13)] = self.f13
            nw119_[int(14)] = d_625_i4_
            nw119_[int(15)] = d_625_i4_
            nw119_[int(16)] = 773
            nw119_[int(17)] = len(d_682_v70_)
            nw119_[int(18)] = default__.fm4(globalState)
            nw119_[int(19)] = len((p1) + (p1))
            nw119_[int(20)] = self.f13
            nw119_[int(21)] = d_625_i4_
            nw119_[int(22)] = d_625_i4_
            nw119_[int(23)] = (0) - ((0) - ((_dafny.MultiSet([d_653_v44_, d_653_v44_, d_653_v44_, ((d_683_v71_)[d_624_v18_] if (d_624_v18_) in (d_683_v71_) else d_653_v44_), d_653_v44_])).cardinality))
            nw119_[int(24)] = self.f13
            nw119_[int(25)] = self.f13
            nw119_[int(26)] = d_625_i4_
            nw119_[int(27)] = d_625_i4_
            nw119_[int(28)] = self.f13
            d_684_v72_ = nw119_
            index112_ = default__.safeIndex(878, (d_684_v72_).length(0))
            (d_684_v72_)[index112_] = self.f13
            index113_ = default__.safeIndex(878, (d_684_v72_).length(0))
            (d_684_v72_)[index113_] = self.f13
        d_624_v18_ = ((_dafny.SeqWithoutIsStrInference([self.f13])) + ((p2).set(default__.safeIndex(self.f13, len(p2)), self.f13))) <= (_dafny.SeqWithoutIsStrInference([self.f13 for d_686_i9_ in range(default__.abs(802))]))


class C3:
    def  __init__(self):
        self._f11: bool = False
        self._f12: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f11, f12):
        (self)._f11 = f11
        (self)._f12 = f12

    def fm1(self, p0, p1, globalState):
        return default__.safeDivisionInt(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngc"))), (_dafny.MultiSet([(self).f11])).cardinality), default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ttwonxung"))), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f12 for d_687_i0_ in range(default__.abs(173))])]))).cardinality))

    def m1(self, p0, globalState):
        d_688_v0_: D0
        d_688_v0_ = D0_DC0(p0)
        d_689_v1_: D0
        d_689_v1_ = D0_DC2(d_688_v0_)
        d_690_v2_: _dafny.Array
        nw120_ = _dafny.Array(_dafny.CodePoint('D'), 14)
        d_690_v2_ = nw120_
        d_691_v3_: _dafny.Map
        d_691_v3_ = _dafny.Map({d_689_v1_: d_690_v2_})
        if (D0_DC2(d_688_v0_)) in (d_691_v3_):
            d_692_v4_: C2
            nw121_ = C2()
            nw121_.ctor__((0) - ((self).f12))
            d_692_v4_ = nw121_
            d_693_v5_: _dafny.MultiSet
            d_693_v5_ = _dafny.MultiSet([(d_692_v4_.f13) - ((self).f12), (self).f12])
            d_693_v5_ = d_693_v5_
            d_694_v6_: _dafny.Array
            nw122_ = _dafny.Array(int(0), 5)
            d_694_v6_ = nw122_
            d_695_v7_: _dafny.Seq
            d_695_v7_ = _dafny.SeqWithoutIsStrInference([p0])
            index114_ = default__.safeIndex(559, (d_694_v6_).length(0))
            (d_694_v6_)[index114_] = len(d_695_v7_)
            index115_ = default__.safeIndex(559, (d_694_v6_).length(0))
            (d_694_v6_)[index115_] = default__.safeModuloInt(-878, len(default__.fm10(globalState)))
            d_696_v8_: bool
            d_696_v8_ = False
            d_697_v9_: _dafny.Set
            d_697_v9_ = _dafny.Set({(d_694_v6_)[default__.safeIndex(559, (d_694_v6_).length(0))]})
            d_696_v8_ = not((d_697_v9_) != (d_697_v9_))
            d_698_v10_: str
            d_698_v10_ = _dafny.CodePoint('e')
            d_699_v11_: _dafny.Seq
            d_699_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdfkv"))
            if default__.fm0((d_698_v10_) in (d_699_v11_), (d_696_v8_ if (self).f11 else (self).f11), (self).f11, (self).f11, globalState):
                (globalState).f2 = ((d_694_v6_)[default__.safeIndex(559, (d_694_v6_).length(0))]) * (d_692_v4_.f13)
                d_696_v8_ = (len(d_695_v7_)) < ((self).f12)
                d_700_v12_: _dafny.Array
                nw123_ = _dafny.Array(False, 23)
                d_700_v12_ = nw123_
                index116_ = default__.safeIndex(560, (d_700_v12_).length(0))
                (d_700_v12_)[index116_] = p0
                d_701_v13_: _dafny.MultiSet
                d_701_v13_ = _dafny.MultiSet([d_695_v7_])
                index117_ = default__.safeIndex(560, (d_700_v12_).length(0))
                (d_700_v12_)[index117_] = (d_701_v13_).issubset((d_701_v13_) | (d_701_v13_))
                (d_692_v4_).f13 = (d_694_v6_)[default__.safeIndex(559, (d_694_v6_).length(0))]
                index118_ = default__.safeIndex(559, (d_694_v6_).length(0))
                (d_694_v6_)[index118_] = ((self).f12) * (d_692_v4_.f13)
            elif True:
                d_689_v1_ = d_689_v1_
                d_702_v14_: C1
                nw124_ = C1()
                nw124_.ctor__((d_693_v5_).ispropersubset(d_693_v5_))
                d_702_v14_ = nw124_
                d_703_v15_: _dafny.Map
                d_703_v15_ = _dafny.Map({(d_694_v6_)[default__.safeIndex(559, (d_694_v6_).length(0))]: d_698_v10_})
                d_704_v16_: _dafny.Map
                d_704_v16_ = _dafny.Map({((d_703_v15_)[(self).f12] if ((self).f12) in (d_703_v15_) else d_698_v10_): (d_698_v10_) in (d_699_v11_)})
                d_696_v8_ = ((d_704_v16_)[_dafny.CodePoint('o')] if (_dafny.CodePoint('o')) in (d_704_v16_) else (self).f11)
                d_705_v17_: _dafny.Array
                nw125_ = _dafny.Array(None, 9)
                nw125_[int(0)] = (d_696_v8_) == ((d_695_v7_)[default__.safeIndex(d_692_v4_.f13, len(d_695_v7_))])
                nw125_[int(1)] = (self).f11
                nw125_[int(2)] = not(d_702_v14_.f14)
                nw125_[int(3)] = (self).f11
                nw125_[int(4)] = (d_695_v7_)[default__.safeIndex(-640, len(d_695_v7_))]
                nw125_[int(5)] = d_702_v14_.f14
                nw125_[int(6)] = p0
                nw125_[int(7)] = d_702_v14_.f14
                nw125_[int(8)] = (self).f11
                d_705_v17_ = nw125_
                d_705_v17_ = d_705_v17_
                d_706_v18_: _dafny.Array
                def lambda31_(d_707_v7_):
                    def lambda32_(d_708_i0_):
                        return d_707_v7_

                    return lambda32_

                init18_ = lambda31_(d_695_v7_)
                nw126_ = _dafny.Array(None, 17)
                for i0_18_ in range(nw126_.length(0)):
                    nw126_[i0_18_] = init18_(i0_18_)
                d_706_v18_ = nw126_
                d_706_v18_ = d_706_v18_
        elif True:
            d_709_v19_: _dafny.Seq
            d_709_v19_ = _dafny.SeqWithoutIsStrInference([(self).f12, (self).f12])
            d_710_v20_: _dafny.Seq
            d_710_v20_ = _dafny.SeqWithoutIsStrInference([(self).f12, len(d_709_v19_), 515])
            d_711_v21_: _dafny.MultiSet
            d_711_v21_ = _dafny.MultiSet([(d_709_v19_) < (d_710_v20_)])
            d_711_v21_ = (d_711_v21_ if p0 else d_711_v21_)
            d_712_v22_: C2
            nw127_ = C2()
            nw127_.ctor__(default__.fm4(globalState))
            d_712_v22_ = nw127_
            d_712_v22_ = d_712_v22_
            d_713_v23_: bool
            d_713_v23_ = True
            d_713_v23_ = d_713_v23_
            d_714_v24_: _dafny.Map
            d_714_v24_ = _dafny.Map({d_712_v22_.f13: ((self).f12) - ((self).f12)})
            rhs116_ = d_713_v23_
            rhs117_ = p0
            rhs118_ = ((self).f12) > (80)
            rhs119_ = p0
            rhs120_ = (d_714_v24_) | (d_714_v24_)
            d_713_v23_ = rhs116_
            d_713_v23_ = rhs117_
            d_713_v23_ = rhs118_
            d_713_v23_ = rhs119_
            d_714_v24_ = rhs120_
            d_715_v25_: C1
            nw128_ = C1()
            nw128_.ctor__((self).f11)
            d_715_v25_ = nw128_
        d_716_v26_: bool
        d_716_v26_ = False
        d_716_v26_ = not(p0)
        d_717_i1_: int
        d_717_i1_ = 0
        with _dafny.label("1"):
            while p0:
                with _dafny.c_label("1"):
                    if (d_717_i1_) >= (100):
                        raise _dafny.Break("1")
                    d_717_i1_ = (d_717_i1_) + (1)
                    d_718_v27_: _dafny.MultiSet
                    d_718_v27_ = _dafny.MultiSet([(self).f11, (self).f11])
                    d_719_v28_: _dafny.Seq
                    d_719_v28_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcu"))
                    d_720_v29_: str
                    d_720_v29_ = _dafny.CodePoint('h')
                    d_721_v30_: _dafny.MultiSet
                    d_721_v30_ = _dafny.MultiSet([(d_719_v28_).set(default__.safeIndex((self).f12, len(d_719_v28_)), d_720_v29_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rbbmc"))])
                    d_722_v31_: bool
                    d_723_v32_: _dafny.Map
                    d_724_v33_: bool
                    out27_: bool
                    out28_: _dafny.Map
                    out29_: bool
                    out27_, out28_, out29_ = (self).m2(((self).f12) >= ((d_718_v27_).cardinality), (d_721_v30_).ispropersubset(d_721_v30_), globalState)
                    d_722_v31_ = out27_
                    d_723_v32_ = out28_
                    d_724_v33_ = out29_
                    d_725_v34_: _dafny.Map
                    d_725_v34_ = _dafny.Map({not(d_716_v26_): _dafny.Set({(self).f11, d_722_v31_, d_722_v31_})})
                    d_726_v35_: _dafny.Set
                    d_726_v35_ = _dafny.Set({(self).f11})
                    d_727_v36_: D4
                    d_727_v36_ = D4_DC12(d_726_v35_)
                    d_725_v34_ = (d_725_v34_).set(p0, (d_727_v36_).cf20)
                    d_719_v28_ = (d_719_v28_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycmdjta")))
                    d_728_v37_: _dafny.Map
                    d_728_v37_ = _dafny.Map({552: d_716_v26_})
                    (globalState).f2 = ((self).f12) + (default__.safeModuloInt((self).f12, len(d_728_v37_)))
                    pass
            pass
        hi4_ = (0) - ((self).f12)
        for d_729_i2_ in range(default__.safeModuloInt((self).f12, (self).f12), hi4_):
            d_730_v38_: _dafny.Seq
            d_730_v38_ = _dafny.SeqWithoutIsStrInference([True])
            d_730_v38_ = ((d_730_v38_).set(default__.safeIndex(d_729_i2_, len(d_730_v38_)), p0)) + (((d_730_v38_).set(default__.safeIndex((self).f12, len(d_730_v38_)), (self).f11)).set(default__.safeIndex((0) - (d_729_i2_), len((d_730_v38_).set(default__.safeIndex((self).f12, len(d_730_v38_)), (self).f11))), not((self).f11)))
            d_731_v39_: _dafny.Array
            def lambda33_(d_732_v26_):
                def lambda34_(d_733_i3_):
                    return d_732_v26_

                return lambda34_

            init19_ = lambda33_(d_716_v26_)
            nw129_ = _dafny.Array(None, 15)
            for i0_19_ in range(nw129_.length(0)):
                nw129_[i0_19_] = init19_(i0_19_)
            d_731_v39_ = nw129_
            d_734_v40_: _dafny.Map
            d_734_v40_ = _dafny.Map({d_731_v39_: d_731_v39_})
            d_735_v41_: _dafny.Seq
            d_735_v41_ = _dafny.SeqWithoutIsStrInference([(self).f12])
            d_736_v42_: _dafny.MultiSet
            d_736_v42_ = _dafny.MultiSet([(self).f12, len(d_735_v41_)])
            d_737_v43_: _dafny.Map
            d_737_v43_ = _dafny.Map({d_716_v26_: d_736_v42_})
            d_738_v44_: _dafny.Map
            d_738_v44_ = _dafny.Map({(self).f11: (self).f11})
            d_739_v45_: _dafny.Map
            d_739_v45_ = _dafny.Map({(self).f12: p0})
            d_740_v46_: _dafny.Map
            d_740_v46_ = _dafny.Map({len(d_738_v44_): len(d_739_v45_)})
            d_741_v47_: _dafny.Array
            nw130_ = _dafny.Array(None, 9)
            nw130_[int(0)] = (self).f12
            nw130_[int(1)] = d_729_i2_
            nw130_[int(2)] = len(d_740_v46_)
            nw130_[int(3)] = d_729_i2_
            nw130_[int(4)] = (self).f12
            nw130_[int(5)] = d_729_i2_
            nw130_[int(6)] = (self).f12
            nw130_[int(7)] = d_729_i2_
            nw130_[int(8)] = d_729_i2_
            d_741_v47_ = nw130_
            d_742_v48_: str
            d_742_v48_ = _dafny.CodePoint('e')
            d_743_v49_: _dafny.Seq
            d_743_v49_ = _dafny.SeqWithoutIsStrInference([d_742_v48_])
            d_744_v50_: _dafny.Seq
            d_744_v50_ = _dafny.SeqWithoutIsStrInference([d_736_v42_])
            d_745_v51_: D4
            d_745_v51_ = D4_DC13(d_742_v48_)
            d_746_v52_: _dafny.MultiSet
            d_746_v52_ = _dafny.MultiSet([d_745_v51_])
            d_747_v53_: _dafny.MultiSet
            d_747_v53_ = d_746_v52_
            d_748_v54_: _dafny.Set
            d_748_v54_ = _dafny.Set({(((d_737_v43_)[d_716_v26_] if (d_716_v26_) in (d_737_v43_) else _dafny.MultiSet([len((_dafny.Map({d_729_i2_: d_741_v47_})))]))).set(len(d_743_v49_), default__.abs(-13)), (d_736_v42_) | (_dafny.MultiSet([(self).f12, (self).f12])), d_736_v42_, (d_744_v50_)[default__.safeIndex((_dafny.MultiSet([len(_dafny.Map({(d_747_v53_): (self).f11})), d_729_i2_, (self).f12])).cardinality, len(d_744_v50_))], (d_736_v42_) - (d_736_v42_)})
            d_749_v55_: _dafny.Set
            d_749_v55_ = _dafny.Set({(self).f11})
            rhs121_ = ((d_749_v55_).intersection(d_749_v55_)).isdisjoint(d_749_v55_)
            rhs122_ = d_734_v40_
            rhs123_ = d_748_v54_
            d_716_v26_ = rhs121_
            d_734_v40_ = rhs122_
            d_748_v54_ = rhs123_
            d_739_v45_ = (d_739_v45_).set(((self).f12) - ((self).f12), d_716_v26_)
            d_716_v26_ = True
        d_750_v56_: str
        d_750_v56_ = _dafny.CodePoint('g')
        if default__.fm0((d_750_v56_) not in (_dafny.SeqWithoutIsStrInference([d_750_v56_ for d_751_i4_ in range(default__.abs(-689))])), d_716_v26_, (d_716_v26_) and ((self).f11), (self).f11, globalState):
            d_752_v57_: C2
            nw131_ = C2()
            nw131_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "my"))))
            d_752_v57_ = nw131_
            if (self).f11:
                d_716_v26_ = not (d_716_v26_) or (not(not(p0)))
                d_753_v58_: _dafny.Seq
                d_753_v58_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoku"))
                d_754_v59_: _dafny.Set
                d_754_v59_ = _dafny.Set({p0, not(not((self).f11)), d_716_v26_, d_716_v26_, d_716_v26_})
                rhs124_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sedfiwcj"))
                rhs125_ = (d_754_v59_) != (d_754_v59_)
                d_753_v58_ = rhs124_
                d_716_v26_ = rhs125_
                d_755_v60_: _dafny.Set
                d_755_v60_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([not(d_716_v26_)]))})
                d_716_v26_ = not(((d_755_v60_) | (_dafny.Set({50, d_752_v57_.f13}))).issubset(d_755_v60_))
                d_716_v26_ = d_716_v26_
                d_756_v61_: D0
                d_756_v61_ = D0_DC0(True)
                d_757_v62_: _dafny.Seq
                d_757_v62_ = _dafny.SeqWithoutIsStrInference([not (p0) or ((d_756_v61_).cf0), not((self).f11)])
                d_757_v62_ = d_757_v62_
            elif True:
                d_758_v63_: C2
                nw132_ = C2()
                nw132_.ctor__(d_752_v57_.f13)
                d_758_v63_ = nw132_
                d_759_v64_: bool
                d_760_v65_: _dafny.Map
                d_761_v66_: bool
                out30_: bool
                out31_: _dafny.Map
                out32_: bool
                out30_, out31_, out32_ = (self).m2(d_716_v26_, d_716_v26_, globalState)
                d_759_v64_ = out30_
                d_760_v65_ = out31_
                d_761_v66_ = out32_
                d_762_v67_: _dafny.Map
                d_762_v67_ = _dafny.Map({d_761_v66_: p0})
                d_763_v68_: _dafny.Seq
                d_763_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yfahqywud"))
                d_764_v70_: _dafny.Seq
                d_764_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_765_i5_ in range(default__.abs(80))])])
                d_766_v71_: _dafny.MultiSet
                def iife40_():
                    coll16_ = _dafny.Map()
                    compr_16_: _dafny.Seq
                    for compr_16_ in (d_764_v70_).Elements:
                        d_767_v69_: _dafny.Seq = compr_16_
                        if (d_767_v69_) in (d_764_v70_):
                            coll16_[d_767_v69_] = (self).f11
                    return _dafny.Map(coll16_)
                d_766_v71_ = _dafny.MultiSet([d_752_v57_.f13, len(d_762_v67_), len(d_763_v68_), len(_dafny.SeqWithoutIsStrInference([len(iife40_()
                )]))])
                d_768_v72_: C1
                nw133_ = C1()
                nw133_.ctor__((d_766_v71_).issubset(default__.fm11(default__.fm0(d_716_v26_, (self).f11, p0, (self).f11, globalState), d_750_v56_, globalState)))
                d_768_v72_ = nw133_
                d_768_v72_ = d_768_v72_
                d_769_v73_: C1
                nw134_ = C1()
                nw134_.ctor__(p0)
                d_769_v73_ = nw134_
                (d_768_v72_).f14 = d_759_v64_
            (globalState).f2 = default__.safeModuloInt(d_752_v57_.f13, (self).f12)
            d_716_v26_ = d_716_v26_
            d_770_v74_: C1
            nw135_ = C1()
            nw135_.ctor__((d_752_v57_.f13) == (184))
            d_770_v74_ = nw135_
        elif True:
            d_771_v75_: _dafny.MultiSet
            d_771_v75_ = _dafny.MultiSet([d_750_v56_, d_750_v56_])
            d_716_v26_ = (d_771_v75_).isdisjoint((d_771_v75_) | (d_771_v75_))
            d_772_v76_: C2
            nw136_ = C2()
            nw136_.ctor__((self).f12)
            d_772_v76_ = nw136_
            d_773_v77_: _dafny.Seq
            d_773_v77_ = _dafny.SeqWithoutIsStrInference([(self).f11, (self).f11])
            d_774_v78_: _dafny.Map
            d_774_v78_ = _dafny.Map({len(d_773_v77_): (self).f11})
            d_775_v79_: _dafny.MultiSet
            d_775_v79_ = _dafny.MultiSet([(self).f11, d_716_v26_, (self).f11])
            d_776_v80_: _dafny.Seq
            d_776_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onxfcxpf"))
            d_777_v81_: _dafny.Array
            nw137_ = _dafny.Array(False, 6)
            d_777_v81_ = nw137_
            d_778_v82_: _dafny.Array
            nw138_ = _dafny.Array(None, 15)
            nw138_[int(0)] = (self).f12
            nw138_[int(1)] = (self).f12
            nw138_[int(2)] = (d_775_v79_).cardinality
            nw138_[int(3)] = (584) - (d_772_v76_.f13)
            nw138_[int(4)] = (0) - ((d_772_v76_.f13) + ((self).f12))
            nw138_[int(5)] = default__.safeModuloInt(len(d_776_v80_), d_772_v76_.f13)
            nw138_[int(6)] = (self).f12
            nw138_[int(7)] = ((self).f12) * (d_772_v76_.f13)
            nw138_[int(8)] = (self).f12
            nw138_[int(9)] = ((self).f12 if (self).f11 else d_772_v76_.f13)
            nw138_[int(10)] = (self).fm1(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lnlmx")), default__.fm2((self).f11, globalState), globalState)
            nw138_[int(11)] = (self).f12
            nw138_[int(12)] = len(_dafny.Map({d_777_v81_: True}))
            nw138_[int(13)] = (self).f12
            nw138_[int(14)] = (self).f12
            d_778_v82_ = nw138_
            index119_ = default__.safeIndex(700, (d_778_v82_).length(0))
            (d_778_v82_)[index119_] = d_772_v76_.f13
            d_779_v83_: C0
            nw139_ = C0()
            nw139_.ctor__()
            d_779_v83_ = nw139_
            d_780_v84_: D7
            d_780_v84_ = D7_DC16(d_779_v83_)
            d_781_v85_: _dafny.Array
            nw140_ = _dafny.Array(None, 27)
            nw140_[int(0)] = d_779_v83_
            nw140_[int(1)] = d_779_v83_
            nw140_[int(2)] = d_779_v83_
            nw140_[int(3)] = d_779_v83_
            nw140_[int(4)] = d_779_v83_
            nw140_[int(5)] = d_779_v83_
            nw140_[int(6)] = d_779_v83_
            nw140_[int(7)] = d_779_v83_
            nw140_[int(8)] = d_779_v83_
            nw140_[int(9)] = d_779_v83_
            nw140_[int(10)] = d_779_v83_
            nw140_[int(11)] = d_779_v83_
            nw140_[int(12)] = d_779_v83_
            nw140_[int(13)] = d_779_v83_
            nw140_[int(14)] = d_779_v83_
            nw140_[int(15)] = d_779_v83_
            nw140_[int(16)] = d_779_v83_
            nw140_[int(17)] = d_779_v83_
            nw140_[int(18)] = d_779_v83_
            nw140_[int(19)] = d_779_v83_
            nw140_[int(20)] = d_779_v83_
            nw140_[int(21)] = d_779_v83_
            nw140_[int(22)] = d_779_v83_
            nw140_[int(23)] = d_779_v83_
            nw140_[int(24)] = (d_780_v84_).cf24
            nw140_[int(25)] = d_779_v83_
            nw140_[int(26)] = d_779_v83_
            d_781_v85_ = nw140_
            d_782_v86_: _dafny.Seq
            d_782_v86_ = _dafny.SeqWithoutIsStrInference([d_781_v85_, d_781_v85_, d_781_v85_, d_781_v85_])
            index120_ = default__.safeIndex(700, (d_778_v82_).length(0))
            rhs126_ = (D8_DC19(d_772_v76_)).cf27
            rhs127_ = (self).f11
            rhs128_ = (d_774_v78_) | (((D3_DC10(d_716_v26_, (self).f12, p0, d_777_v81_, d_774_v78_)).cf18) | (d_774_v78_))
            rhs129_ = (self).f12
            rhs130_ = (d_782_v86_) + (d_782_v86_)
            lhs87_ = d_778_v82_
            lhs88_ = default__.safeIndex(700, (d_778_v82_).length(0))
            d_772_v76_ = rhs126_
            d_716_v26_ = rhs127_
            d_774_v78_ = rhs128_
            lhs87_[lhs88_] = rhs129_
            d_782_v86_ = rhs130_
            d_783_v87_: D4
            d_783_v87_ = D4_DC13(_dafny.CodePoint('f'))
            d_784_v88_: _dafny.Map
            d_784_v88_ = _dafny.Map({p0: d_783_v87_})
            d_785_v89_: _dafny.Seq
            d_785_v89_ = _dafny.SeqWithoutIsStrInference([d_772_v76_.f13, len(d_784_v88_)])
            d_786_v90_: _dafny.Seq
            d_786_v90_ = _dafny.SeqWithoutIsStrInference([len(d_785_v89_), d_772_v76_.f13])
            d_787_v91_: _dafny.Seq
            d_787_v91_ = _dafny.SeqWithoutIsStrInference([len(d_786_v90_)])
            (d_772_v76_).m3(d_777_v81_, d_776_v80_, (d_785_v89_) + (_dafny.SeqWithoutIsStrInference([-298 for d_788_i6_ in range(default__.abs(911))])), globalState)
            if d_716_v26_:
                (d_772_v76_).f13 = d_772_v76_.f13
                rhs131_ = (d_772_v76_.f13) * (((self).f12) + (len(d_787_v91_)))
                rhs132_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebyufrei"))) + (d_776_v80_)
                lhs89_ = d_772_v76_
                lhs89_.f13 = rhs131_
                d_776_v80_ = rhs132_
                (globalState).f2 = (d_778_v82_)[default__.safeIndex(700, (d_778_v82_).length(0))]
                d_776_v80_ = (d_776_v80_).set(default__.safeIndex((self).f12, len(d_776_v80_)), d_750_v56_)
                index121_ = default__.safeIndex(891, (d_777_v81_).length(0))
                (d_777_v81_)[index121_] = p0
                index122_ = default__.safeIndex(891, (d_777_v81_).length(0))
                (d_777_v81_)[index122_] = True
            elif True:
                d_789_v92_: _dafny.MultiSet
                d_789_v92_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wtrbr"))])
                d_790_v93_: D1
                d_790_v93_ = D1_DC4((d_789_v92_) - (d_789_v92_), (0) - ((d_778_v82_)[default__.safeIndex(700, (d_778_v82_).length(0))]))
                d_790_v93_ = d_790_v93_
                (globalState).f2 = ((d_778_v82_)[default__.safeIndex(700, (d_778_v82_).length(0))]) + ((self).f12)
                index123_ = default__.safeIndex(395, (d_777_v81_).length(0))
                (d_777_v81_)[index123_] = p0
                index124_ = default__.safeIndex(395, (d_777_v81_).length(0))
                (d_777_v81_)[index124_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "brsji"))) != (d_776_v80_)
                def iife41_():
                    coll17_ = _dafny.Map()
                    compr_17_: int
                    for compr_17_ in (d_774_v78_).keys.Elements:
                        d_791_v94_: int = compr_17_
                        if (d_791_v94_) in (d_774_v78_):
                            coll17_[default__.safeDivisionInt(d_791_v94_, (self).f12)] = d_750_v56_
                    return _dafny.Map(coll17_)
                (globalState).f2 = (0) - ((0) - (default__.safeModuloInt(d_772_v76_.f13, len(iife41_()
                ))))
                d_792_v95_: _dafny.Map
                d_792_v95_ = _dafny.Map({(d_772_v76_.f13 if d_716_v26_ else (d_778_v82_)[default__.safeIndex(700, (d_778_v82_).length(0))]): d_776_v80_})
                d_792_v95_ = d_792_v95_
            d_793_v96_: _dafny.Set
            d_793_v96_ = _dafny.Set({not(p0), (self).f11})
            d_716_v26_ = (not((_dafny.Set({(d_785_v89_)[default__.safeIndex(d_772_v76_.f13, len(d_785_v89_))], (_dafny.MultiSet([True])).cardinality, (self).f12})).isdisjoint(_dafny.Set({(d_778_v82_)[default__.safeIndex(700, (d_778_v82_).length(0))]})))) in (d_793_v96_)
        d_794_v97_: _dafny.Seq
        d_794_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilyuyef"))
        d_795_v98_: _dafny.MultiSet
        d_795_v98_ = _dafny.MultiSet([d_716_v26_])
        d_796_v99_: _dafny.Map
        d_796_v99_ = _dafny.Map({d_794_v97_: (d_795_v98_).cardinality})
        d_797_v100_: _dafny.Array
        nw141_ = _dafny.Array(int(0), 27)
        d_797_v100_ = nw141_
        d_798_v101_: _dafny.Map
        d_798_v101_ = _dafny.Map({d_796_v99_: d_797_v100_})
        d_799_v102_: _dafny.MultiSet
        d_799_v102_ = _dafny.MultiSet([192])
        d_800_v103_: _dafny.MultiSet
        d_800_v103_ = _dafny.MultiSet([d_799_v102_, d_799_v102_, _dafny.MultiSet([(self).f12]), (_dafny.MultiSet([(self).f12])).set((self).f12, default__.abs((self).f12)), d_799_v102_])
        d_798_v101_ = (d_798_v101_).set(((d_796_v99_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufaahfrp")), (self).f12)).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yaws")), ((d_800_v103_)[d_799_v102_] if (d_799_v102_) in (d_800_v103_) else (d_795_v98_).cardinality)), d_797_v100_)

    def m2(self, p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        r2: bool = False
        d_801_v0_: _dafny.Seq
        d_801_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iteghrvta"))
        d_802_v1_: _dafny.Map
        d_802_v1_ = _dafny.Map({d_801_v0_: (self).f12})
        d_802_v1_ = (d_802_v1_).set(d_801_v0_, (self).f12)
        d_803_v2_: _dafny.Array
        nw142_ = _dafny.Array(False, 23)
        d_803_v2_ = nw142_
        d_804_v3_: _dafny.Array
        nw143_ = _dafny.Array(_dafny.Map({}), 19)
        d_804_v3_ = nw143_
        d_805_v4_: str
        d_805_v4_ = _dafny.CodePoint('w')
        d_806_v5_: _dafny.MultiSet
        d_806_v5_ = _dafny.MultiSet([d_805_v4_])
        d_807_v6_: _dafny.Seq
        d_807_v6_ = _dafny.SeqWithoutIsStrInference([default__.fm2(p0, globalState)])
        default__.m0(d_803_v2_, (self).f12, d_804_v3_, ((d_806_v5_)[d_805_v4_] if (d_805_v4_) in (d_806_v5_) else len(d_807_v6_)), globalState)
        d_808_v7_: C0
        nw144_ = C0()
        nw144_.ctor__()
        d_808_v7_ = nw144_
        d_809_v8_: _dafny.Array
        def lambda35_(d_810_i0_):
            return default__.safeDivisionInt(d_810_i0_, (self).f12)

        init20_ = lambda35_
        nw145_ = _dafny.Array(None, 29)
        for i0_20_ in range(nw145_.length(0)):
            nw145_[i0_20_] = init20_(i0_20_)
        d_809_v8_ = nw145_
        (globalState).f10 = d_809_v8_
        if p0:
            index125_ = default__.safeIndex(312, (d_803_v2_).length(0))
            (d_803_v2_)[index125_] = (p0) == ((self).f11)
            d_811_v9_: _dafny.Map
            d_811_v9_ = _dafny.Map({p1: p1})
            d_812_v10_: _dafny.Set
            d_812_v10_ = _dafny.Set({((d_811_v9_)[p1] if (p1) in (d_811_v9_) else p1), (self).f11})
            d_813_v11_: _dafny.Map
            d_813_v11_ = _dafny.Map({d_812_v10_: p0})
            index126_ = default__.safeIndex(312, (d_803_v2_).length(0))
            (d_803_v2_)[index126_] = ((d_813_v11_)[d_812_v10_] if (d_812_v10_) in (d_813_v11_) else (self).f11)
            if (self).f11:
                d_814_v12_: _dafny.MultiSet
                d_814_v12_ = _dafny.MultiSet([p1])
                r2 = not ((d_814_v12_).ispropersubset(d_814_v12_)) or (p0)
                d_815_v13_: _dafny.MultiSet
                d_815_v13_ = _dafny.MultiSet([d_812_v10_])
                d_816_v14_: _dafny.Seq
                d_816_v14_ = _dafny.SeqWithoutIsStrInference([not(not((d_815_v13_).isdisjoint(d_815_v13_)))])
                d_816_v14_ = d_816_v14_
                index127_ = default__.safeIndex(312, (d_803_v2_).length(0))
                rhs133_ = default__.fm6((self).f12, p0, default__.fm0((self).f11, True, p1, (self).f11, globalState), d_805_v4_, globalState)
                rhs134_ = not((d_803_v2_)[default__.safeIndex(312, (d_803_v2_).length(0))])
                lhs90_ = d_803_v2_
                lhs91_ = default__.safeIndex(312, (d_803_v2_).length(0))
                d_805_v4_ = rhs133_
                lhs90_[lhs91_] = rhs134_
                d_817_v15_: D1
                d_817_v15_ = D1_DC5(p0, p0)
                r2 = (d_808_v7_).fm7(((d_817_v15_).cf9 if p1 else (self).f11), D1_DC3(d_805_v4_), d_801_v0_, (d_803_v2_)[default__.safeIndex(312, (d_803_v2_).length(0))], globalState)
                index128_ = default__.safeIndex(941, (d_809_v8_).length(0))
                (d_809_v8_)[index128_] = (self).f12
                index129_ = default__.safeIndex(941, (d_809_v8_).length(0))
                (d_809_v8_)[index129_] = default__.safeDivisionInt(default__.safeModuloInt((self).f12, (self).f12), (self).f12)
            elif True:
                d_818_v16_: _dafny.Array
                nw146_ = _dafny.Array(D8.default()(), 10)
                d_818_v16_ = nw146_
                d_819_v17_: C2
                nw147_ = C2()
                nw147_.ctor__((self).f12)
                d_819_v17_ = nw147_
                d_820_v18_: D8
                d_820_v18_ = D8_DC19(d_819_v17_)
                index130_ = default__.safeIndex(588, (d_818_v16_).length(0))
                (d_818_v16_)[index130_] = (d_820_v18_ if p1 else D8_DC19(d_819_v17_))
                index131_ = default__.safeIndex(588, (d_818_v16_).length(0))
                (d_818_v16_)[index131_] = d_820_v18_
                d_821_v19_: _dafny.Map
                d_821_v19_ = _dafny.Map({d_819_v17_.f13: len(d_801_v0_)})
                d_822_v20_: _dafny.Map
                d_822_v20_ = _dafny.Map({(d_803_v2_)[default__.safeIndex(312, (d_803_v2_).length(0))]: d_821_v19_})
                d_822_v20_ = (d_822_v20_).set(False, d_821_v19_)
                d_823_v21_: D7
                d_823_v21_ = D7_DC17((d_803_v2_)[default__.safeIndex(312, (d_803_v2_).length(0))])
                d_824_v22_: D7
                d_824_v22_ = D7_DC18(d_823_v21_)
                d_825_v23_: D7
                d_825_v23_ = D7_DC18(d_823_v21_)
                d_826_v24_: _dafny.Map
                d_826_v24_ = _dafny.Map({d_825_v23_: (d_803_v2_)[default__.safeIndex(312, (d_803_v2_).length(0))]})
                index132_ = default__.safeIndex(897, (d_809_v8_).length(0))
                (d_809_v8_)[index132_] = (d_807_v6_)[default__.safeIndex(len(d_826_v24_), len(d_807_v6_))]
                index133_ = default__.safeIndex(897, (d_809_v8_).length(0))
                (d_809_v8_)[index133_] = d_819_v17_.f13
                d_827_v25_: _dafny.Array
                def lambda36_(d_828_v8_):
                    def lambda37_(d_829_i1_):
                        return (d_829_i1_) * ((d_828_v8_)[default__.safeIndex(897, (d_828_v8_).length(0))])

                    return lambda37_

                init21_ = lambda36_(d_809_v8_)
                nw148_ = _dafny.Array(None, 21)
                for i0_21_ in range(nw148_.length(0)):
                    nw148_[i0_21_] = init21_(i0_21_)
                d_827_v25_ = nw148_
                index134_ = default__.safeIndex(312, (d_803_v2_).length(0))
                rhs135_ = d_827_v25_
                rhs136_ = d_819_v17_.f13
                rhs137_ = p0
                rhs138_ = (d_809_v8_)[default__.safeIndex(897, (d_809_v8_).length(0))]
                rhs139_ = (default__.fm12(d_819_v17_.f13, globalState)) + (((d_807_v6_) + (d_807_v6_)).set(default__.safeIndex((self).f12, len((d_807_v6_) + (d_807_v6_))), 406))
                lhs92_ = globalState
                lhs93_ = globalState
                lhs94_ = d_803_v2_
                lhs95_ = default__.safeIndex(312, (d_803_v2_).length(0))
                lhs96_ = globalState
                lhs97_ = globalState
                lhs92_.f10 = rhs135_
                lhs93_.f2 = rhs136_
                lhs94_[lhs95_] = rhs137_
                lhs96_.f2 = rhs138_
                lhs97_.f7 = rhs139_
                d_830_v26_: _dafny.Array
                nw149_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_830_v26_ = nw149_
                d_831_v27_: _dafny.Array
                nw150_ = _dafny.Array(False, 12)
                d_831_v27_ = nw150_
                index135_ = default__.safeIndex(441, (d_830_v26_).length(0))
                (d_830_v26_)[index135_] = d_831_v27_
                index136_ = default__.safeIndex(441, (d_830_v26_).length(0))
                (d_830_v26_)[index136_] = d_831_v27_
            (globalState).f10 = d_809_v8_
            d_811_v9_ = (d_811_v9_).set((self).f11, (d_803_v2_)[default__.safeIndex(312, (d_803_v2_).length(0))])
            d_832_v28_: C2
            nw151_ = C2()
            nw151_.ctor__(((self).f12) * ((self).f12))
            d_832_v28_ = nw151_
        elif True:
            d_833_v29_: _dafny.Map
            d_833_v29_ = _dafny.Map({(self).f12: d_809_v8_})
            d_833_v29_ = d_833_v29_
            (globalState).f7 = d_807_v6_
            d_834_v30_: C2
            nw152_ = C2()
            nw152_.ctor__(len(d_801_v0_))
            d_834_v30_ = nw152_
            d_835_v31_: _dafny.Set
            d_835_v31_ = _dafny.Set({d_834_v30_})
            (globalState).f2 = len((_dafny.Set({d_834_v30_})) - (d_835_v31_))
            d_836_v32_: _dafny.Seq
            d_836_v32_ = d_801_v0_
            d_837_v33_: _dafny.Array
            nw153_ = _dafny.Array(None, 13)
            nw153_[int(0)] = ((d_836_v32_)) + (d_801_v0_)
            nw153_[int(1)] = d_801_v0_
            nw153_[int(2)] = (d_801_v0_) + (d_801_v0_)
            nw153_[int(3)] = d_801_v0_
            nw153_[int(4)] = d_801_v0_
            nw153_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kctfg"))) + (_dafny.SeqWithoutIsStrInference([d_805_v4_ for d_838_i2_ in range(default__.abs(908))]))
            nw153_[int(6)] = d_801_v0_
            nw153_[int(7)] = d_801_v0_
            nw153_[int(8)] = d_801_v0_
            nw153_[int(9)] = (d_801_v0_) + (d_801_v0_)
            nw153_[int(10)] = d_801_v0_
            nw153_[int(11)] = (d_801_v0_) + (default__.fm9(p1, globalState))
            nw153_[int(12)] = d_801_v0_
            d_837_v33_ = nw153_
            index137_ = default__.safeIndex(651, (d_837_v33_).length(0))
            (d_837_v33_)[index137_] = d_801_v0_
            index138_ = default__.safeIndex(651, (d_837_v33_).length(0))
            (d_837_v33_)[index138_] = d_801_v0_
            d_839_v34_: _dafny.Set
            d_839_v34_ = _dafny.Set({(self).f12})
            d_840_v35_: _dafny.MultiSet
            d_840_v35_ = _dafny.MultiSet([(self).f11])
            d_841_v36_: _dafny.Map
            d_841_v36_ = _dafny.Map({(self).f12: (d_840_v35_).cardinality})
            def iife42_():
                coll18_ = _dafny.Set()
                compr_18_: int
                for compr_18_ in (d_841_v36_).keys.Elements:
                    d_842_v37_: int = compr_18_
                    if (d_842_v37_) in (d_841_v36_):
                        coll18_ = coll18_.union(_dafny.Set([default__.safeModuloInt(d_842_v37_, (0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([365]))).cardinality))]))
                return _dafny.Set(coll18_)
            d_839_v34_ = (iife42_()
            ) - (d_839_v34_)
        (globalState).f2 = (self).f12
        r0 = (self).f11
        d_843_v38_: _dafny.Seq
        d_843_v38_ = _dafny.SeqWithoutIsStrInference([not((self).f11), p0, (self).f11])
        d_844_v39_: _dafny.Map
        d_844_v39_ = _dafny.Map({d_843_v38_: p0})
        r1 = d_844_v39_
        r2 = (d_843_v38_) <= (d_843_v38_)
        return r0, r1, r2

    @property
    def f11(self):
        return self._f11
    @property
    def f12(self):
        return self._f12
