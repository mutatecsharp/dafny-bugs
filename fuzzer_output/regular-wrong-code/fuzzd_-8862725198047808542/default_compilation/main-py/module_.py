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
    def fm3(p0, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in (_dafny.SeqWithoutIsStrInference([887 for d_0_i0_ in range(default__.abs(946))])).Elements:
                d_1_v0_: int = compr_0_
                if (d_1_v0_) in (_dafny.SeqWithoutIsStrInference([887 for d_0_i0_ in range(default__.abs(946))])):
                    coll0_[default__.safeDivisionInt(d_1_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lll"))))] = True
            return _dafny.Map(coll0_)
        return ((len(_dafny.Map({True: False}))) + (618)) * ((len(iife0_()
        )) + (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2_i1_ in range(default__.abs(-226))]))))

    @staticmethod
    def fm4(p0, p1, p2, p3, globalState):
        source0_ = (D3_DC11(D3_DC10(_dafny.CodePoint('y'))) if not(not(False)) else D3_DC11(D3_DC10(_dafny.CodePoint('u'))))
        if source0_.is_DC9:
            d_3___mcc_h0_ = source0_.cf8
            d_4___mcc_h1_ = source0_.cf9
            d_5___mcc_h2_ = source0_.cf10
            d_6___mcc_h3_ = source0_.cf11
            d_7_cf11_ = d_6___mcc_h3_
            d_8_cf10_ = d_5___mcc_h2_
            d_9_cf9_ = d_4___mcc_h1_
            d_10_cf8_ = d_3___mcc_h0_
            return _dafny.CodePoint('o')
        elif source0_.is_DC10:
            d_11___mcc_h4_ = source0_.cf12
            d_12_cf12_ = d_11___mcc_h4_
            return d_12_cf12_
        elif source0_.is_DC8:
            d_13___mcc_h5_ = source0_.cf7
            d_14_cf7_ = d_13___mcc_h5_
            return _dafny.CodePoint('b')
        elif True:
            d_15___mcc_h6_ = source0_.cf13
            d_16_cf13_ = d_15___mcc_h6_
            return _dafny.CodePoint('h')

    @staticmethod
    def fm5(p0, p1, p2, p3, globalState):
        return ((_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([not(True)]))) + (_dafny.SeqWithoutIsStrInference([True, True]))

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxm"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uaxq")))

    @staticmethod
    def fm7(p0, p1, p2, p3, globalState):
        if (_dafny.MultiSet([668])) == (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([-644, (_dafny.MultiSet([False, True, not(False), False])).cardinality])), len(_dafny.Map({_dafny.CodePoint('u'): True})), (_dafny.MultiSet([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([-321, 102, 974, 250])): _dafny.CodePoint('h')})])).cardinality])):
            if True:
                return D1_DC4(False)
            elif True:
                return D1_DC4(False)
        elif (D7_DC21(False, _dafny.CodePoint('o'), _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([D0_DC1((0) - (len(_dafny.SeqWithoutIsStrInference([361]))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w")))) for d_17_i0_ in range(default__.abs(-439))]))]))).cf30:
            return D1_DC4(True)
        elif True:
            return D1_DC4(False)

    @staticmethod
    def fm8(p0, p1, p2, p3, globalState):
        return ((_dafny.Map({D1_DC4(False): not(not(True))})) | (_dafny.Map({D1_DC4(False): True}))) | ((_dafny.Map({D1_DC4(True): True})) | (_dafny.Map({D1_DC4(True): True})))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return (_dafny.Set({False})).intersection(_dafny.Set({True, True, False, True, True}))

    @staticmethod
    def fm10(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in _dafny.IntegerRange(896, 351):
                d_18_v0_: int = compr_1_
                if ((896) <= (d_18_v0_)) and ((d_18_v0_) < (351)):
                    coll1_[(d_18_v0_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: False})) for d_19_i0_ in range(default__.abs(-24))]))).cardinality)] = True
            return _dafny.Map(coll1_)
        return (_dafny.Map({871: False})) | ((iife1_()
        ) | (_dafny.Map({821: False})))

    @staticmethod
    def fm11(p0, p1, p2, globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in _dafny.IntegerRange(-758, 774):
                d_20_v0_: int = compr_2_
                if ((-758) <= (d_20_v0_)) and ((d_20_v0_) < (774)):
                    coll2_[default__.safeDivisionInt(d_20_v0_, len(_dafny.SeqWithoutIsStrInference([(0) - ((_dafny.MultiSet([531])).cardinality), len(_dafny.SeqWithoutIsStrInference([-896]))])))] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fixlptlef")))
            return _dafny.Map(coll2_)
        return _dafny.MultiSet([len((_dafny.Map({not(False): len(iife2_()
        )})) | (_dafny.Map({False: 49}))), (len(_dafny.Map({829: len(_dafny.Map({748: False}))}))) - (795)])

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        if True:
            return (_dafny.MultiSet([True, True, not(not(True)), True, True])).issubset(_dafny.MultiSet([False]))
        elif True:
            return ((0) - (len(_dafny.Set({748, 26})))) >= (761)

    @staticmethod
    def fm13(p0, p1, p2, globalState):
        return _dafny.MultiSet([not(not(False)), (_dafny.SeqWithoutIsStrInference([True])) < (_dafny.SeqWithoutIsStrInference([True, not(not(False))])), False])

    @staticmethod
    def fm14(p0, globalState):
        return (_dafny.Map({True: False})) | ((_dafny.Map({True: False})) | (_dafny.Map({False: False})))

    @staticmethod
    def fm15(globalState):
        if not((D7_DC21(True, _dafny.CodePoint('s'), _dafny.MultiSet([756]))).cf30):
            return _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfbfbvbsb"))): _dafny.Map({453: False})})
        elif True:
            def iife3_():
                coll3_ = _dafny.Map()
                compr_3_: int
                for compr_3_ in _dafny.IntegerRange(-111, 887):
                    d_21_v0_: int = compr_3_
                    if ((-111) <= (d_21_v0_)) and ((d_21_v0_) < (887)):
                        coll3_[default__.safeModuloInt(d_21_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g"))))] = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([not(False)])): not(True)})
                return _dafny.Map(coll3_)
            return (_dafny.Map({514: _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdxlfgub")))): False})})) | (iife3_()
            )

    @staticmethod
    def fm16(p0, p1, p2, globalState):
        def iife4_():
            coll4_ = _dafny.Set()
            compr_4_: int
            for compr_4_ in _dafny.IntegerRange(344, 194):
                d_23_v0_: int = compr_4_
                if ((344) <= (d_23_v0_)) and ((d_23_v0_) < (194)):
                    coll4_ = coll4_.union(_dafny.Set([(d_23_v0_) * (-608)]))
            return _dafny.Set(coll4_)
        return ((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxeiish"))): 755}))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({not(False): len(_dafny.SeqWithoutIsStrInference([False, False]))}) for d_22_i0_ in range(default__.abs(-863))]))) + ((_dafny.SeqWithoutIsStrInference([_dafny.Map({True: len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgghpfsb")))}))})])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: len(iife4_()
        )}), _dafny.Map({False: len(_dafny.Set({True, not(False)}))})])))

    @staticmethod
    def fm17(globalState):
        return _dafny.MultiSet([_dafny.CodePoint('a'), _dafny.CodePoint('u'), _dafny.CodePoint('c')])

    @staticmethod
    def fm18(p0, globalState):
        source1_ = D3_DC8(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ycfygtdb")))
        if source1_.is_DC9:
            d_24___mcc_h0_ = source1_.cf8
            d_25___mcc_h1_ = source1_.cf9
            d_26___mcc_h2_ = source1_.cf10
            d_27___mcc_h3_ = source1_.cf11
            d_28_cf11_ = d_27___mcc_h3_
            d_29_cf10_ = d_26___mcc_h2_
            d_30_cf9_ = d_25___mcc_h1_
            d_31_cf8_ = d_24___mcc_h0_
            return D1_DC5(D1_DC5(D1_DC4(d_30_cf9_)))
        elif source1_.is_DC10:
            d_32___mcc_h4_ = source1_.cf12
            d_33_cf12_ = d_32___mcc_h4_
            return D1_DC5(D1_DC5(D1_DC3(not(True))))
        elif source1_.is_DC8:
            d_34___mcc_h5_ = source1_.cf7
            d_35_cf7_ = d_34___mcc_h5_
            return D1_DC5(D1_DC5(D1_DC5(D1_DC4(False))))
        elif True:
            d_36___mcc_h6_ = source1_.cf13
            d_37_cf13_ = d_36___mcc_h6_
            return D1_DC5(D1_DC5(D1_DC5(D1_DC3(not(True)))))

    @staticmethod
    def fm19(p0, p1, globalState):
        if not (False) or ((D7_DC21(False, _dafny.CodePoint('f'), _dafny.MultiSet([len(_dafny.Set({_dafny.Map({False: 914}), _dafny.Map({True: (_dafny.MultiSet([len(_dafny.Map({True: True})), 355, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lloh"))), 384])).cardinality})})), -576]))).cf30):
            return D1_DC3(not(not(False)))
        elif True:
            return D1_DC3(False)

    @staticmethod
    def fm20(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cytq")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))])) + ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcqqyceot"))])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omkpbmjjq"))])))

    @staticmethod
    def fm21(p0, p1, p2, p3, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: str
            for compr_5_ in (_dafny.MultiSet([_dafny.CodePoint('x'), _dafny.CodePoint('m')])).Elements:
                d_38_v0_: str = compr_5_
                if (d_38_v0_) in (_dafny.MultiSet([_dafny.CodePoint('x'), _dafny.CodePoint('m')])):
                    coll5_[d_38_v0_] = _dafny.CodePoint('p')
            return _dafny.Map(coll5_)
        return (_dafny.Map({_dafny.CodePoint('j'): _dafny.CodePoint('g')})) | ((_dafny.Map({_dafny.CodePoint('n'): _dafny.CodePoint('d')})) | (iife5_()
        ))

    @staticmethod
    def m5(p0, p1, globalState):
        r0: bool = False
        r1: _dafny.Map = _dafny.Map({})
        d_39_v0_: _dafny.Seq
        d_39_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ojcommo"))
        d_40_v1_: _dafny.Set
        d_40_v1_ = _dafny.Set({d_39_v0_})
        d_41_i0_: int
        d_41_i0_ = 0
        with _dafny.label("0"):
            while (457) > (len(d_40_v1_)):
                with _dafny.c_label("0"):
                    if (d_41_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_41_i0_ = (d_41_i0_) + (1)
                    d_42_v2_: _dafny.Array
                    nw0_ = _dafny.Array(False, 25)
                    d_42_v2_ = nw0_
                    d_43_v3_: bool
                    d_43_v3_ = True
                    index0_ = default__.safeIndex(675, (d_42_v2_).length(0))
                    (d_42_v2_)[index0_] = default__.fm12(d_43_v3_, True, d_43_v3_, globalState)
                    d_44_v4_: _dafny.Array
                    def lambda0_(d_45_i1_):
                        return (d_45_i1_) + (364)

                    init0_ = lambda0_
                    nw1_ = _dafny.Array(None, 16)
                    for i0_0_ in range(nw1_.length(0)):
                        nw1_[i0_0_] = init0_(i0_0_)
                    d_44_v4_ = nw1_
                    d_46_v5_: int
                    d_46_v5_ = 521
                    index1_ = default__.safeIndex(965, (d_44_v4_).length(0))
                    (d_44_v4_)[index1_] = d_46_v5_
                    d_47_v6_: D3
                    d_47_v6_ = D3_DC8(d_39_v0_)
                    index2_ = default__.safeIndex(675, (d_42_v2_).length(0))
                    index3_ = default__.safeIndex(965, (d_44_v4_).length(0))
                    rhs0_ = (D1_DC3(d_43_v3_)).cf3
                    rhs1_ = d_46_v5_
                    rhs2_ = d_47_v6_
                    lhs0_ = d_42_v2_
                    lhs1_ = default__.safeIndex(675, (d_42_v2_).length(0))
                    lhs2_ = d_44_v4_
                    lhs3_ = default__.safeIndex(965, (d_44_v4_).length(0))
                    lhs0_[lhs1_] = rhs0_
                    lhs2_[lhs3_] = rhs1_
                    d_47_v6_ = rhs2_
                    if True:
                        d_48_v7_: _dafny.Array
                        nw2_ = _dafny.Array(_dafny.Array(None, 0), 19)
                        d_48_v7_ = nw2_
                        index4_ = default__.safeIndex(708, (d_48_v7_).length(0))
                        (d_48_v7_)[index4_] = d_44_v4_
                        d_49_v8_: _dafny.Set
                        d_49_v8_ = _dafny.Set({(d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], d_43_v3_})
                        d_50_v10_: _dafny.Map
                        def iife6_():
                            coll6_ = _dafny.Map()
                            compr_6_: int
                            for compr_6_ in (p1).Elements:
                                d_51_v9_: int = compr_6_
                                if (d_51_v9_) in (p1):
                                    coll6_[(d_51_v9_) - ((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))])] = d_43_v3_
                            return _dafny.Map(coll6_)
                        d_50_v10_ = _dafny.Map({len(iife6_()
                        ): d_43_v3_})
                        index5_ = default__.safeIndex(708, (d_48_v7_).length(0))
                        (d_48_v7_)[index5_] = (d_44_v4_ if (d_49_v8_).issubset(d_49_v8_) else (d_44_v4_ if ((d_50_v10_)[(p0)[default__.safeIndex(d_46_v5_, len(p0))]] if ((p0)[default__.safeIndex(d_46_v5_, len(p0))]) in (d_50_v10_) else d_43_v3_) else d_44_v4_))
                        index6_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        (d_44_v4_)[index6_] = (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]
                        index7_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        (d_44_v4_)[index7_] = (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]
                        d_52_v11_: _dafny.Map
                        d_52_v11_ = _dafny.Map({d_42_v2_: d_46_v5_})
                        d_52_v11_ = ((d_52_v11_) | (d_52_v11_)) | (d_52_v11_)
                        d_46_v5_ = default__.safeDivisionInt((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], default__.fm3(not(False), globalState))
                    elif True:
                        index8_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        (d_44_v4_)[index8_] = (0) - ((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))])
                        d_53_v12_: T0
                        nw3_ = C2()
                        nw3_.ctor__((d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], d_46_v5_, (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))])
                        d_53_v12_ = nw3_
                        d_54_v13_: _dafny.Map
                        d_54_v13_ = _dafny.Map({d_46_v5_: d_53_v12_})
                        d_55_v14_: _dafny.Array
                        nw4_ = _dafny.Array(None, 28)
                        nw4_[int(0)] = d_53_v12_
                        nw4_[int(1)] = d_53_v12_
                        nw4_[int(2)] = d_53_v12_
                        nw4_[int(3)] = d_53_v12_
                        nw4_[int(4)] = d_53_v12_
                        nw4_[int(5)] = d_53_v12_
                        nw4_[int(6)] = d_53_v12_
                        nw4_[int(7)] = d_53_v12_
                        nw4_[int(8)] = d_53_v12_
                        nw4_[int(9)] = d_53_v12_
                        nw4_[int(10)] = d_53_v12_
                        nw4_[int(11)] = d_53_v12_
                        nw4_[int(12)] = d_53_v12_
                        nw4_[int(13)] = d_53_v12_
                        nw4_[int(14)] = ((d_54_v13_)[-282] if (-282) in (d_54_v13_) else d_53_v12_)
                        nw4_[int(15)] = d_53_v12_
                        nw4_[int(16)] = d_53_v12_
                        nw4_[int(17)] = d_53_v12_
                        nw4_[int(18)] = d_53_v12_
                        nw4_[int(19)] = d_53_v12_
                        nw4_[int(20)] = d_53_v12_
                        nw4_[int(21)] = d_53_v12_
                        nw4_[int(22)] = d_53_v12_
                        nw4_[int(23)] = d_53_v12_
                        nw4_[int(24)] = d_53_v12_
                        nw4_[int(25)] = d_53_v12_
                        nw4_[int(26)] = d_53_v12_
                        nw4_[int(27)] = d_53_v12_
                        d_55_v14_ = nw4_
                        d_56_v15_: _dafny.Map
                        d_56_v15_ = _dafny.Map({d_55_v14_: (d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))]})
                        d_57_v16_: _dafny.Map
                        d_57_v16_ = _dafny.Map({d_56_v15_: (p0)[default__.safeIndex((0) - ((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]), len(p0))]})
                        d_57_v16_ = (d_57_v16_).set((d_56_v15_) | (d_56_v15_), (d_53_v12_).f4)
                        d_58_v17_: _dafny.Seq
                        d_58_v17_ = _dafny.SeqWithoutIsStrInference([d_39_v0_])
                        d_59_v18_: _dafny.Array
                        nw5_ = _dafny.Array(None, 12)
                        nw5_[int(0)] = d_39_v0_
                        nw5_[int(1)] = (d_58_v17_)[default__.safeIndex((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], len(d_58_v17_))]
                        nw5_[int(2)] = d_39_v0_
                        nw5_[int(3)] = (d_39_v0_) + (d_39_v0_)
                        nw5_[int(4)] = d_39_v0_
                        nw5_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_60_i2_ in range(default__.abs(-331))])
                        nw5_[int(6)] = d_39_v0_
                        nw5_[int(7)] = d_39_v0_
                        nw5_[int(8)] = d_39_v0_
                        nw5_[int(9)] = d_39_v0_
                        nw5_[int(10)] = d_39_v0_
                        nw5_[int(11)] = d_39_v0_
                        d_59_v18_ = nw5_
                        index9_ = default__.safeIndex(86, (d_59_v18_).length(0))
                        (d_59_v18_)[index9_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owumuar"))
                        index10_ = default__.safeIndex(675, (d_42_v2_).length(0))
                        index11_ = default__.safeIndex(86, (d_59_v18_).length(0))
                        rhs3_ = d_43_v3_
                        rhs4_ = d_43_v3_
                        rhs5_ = (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]
                        rhs6_ = d_39_v0_
                        lhs4_ = globalState
                        lhs5_ = d_42_v2_
                        lhs6_ = default__.safeIndex(675, (d_42_v2_).length(0))
                        lhs7_ = d_53_v12_
                        lhs8_ = d_59_v18_
                        lhs9_ = default__.safeIndex(86, (d_59_v18_).length(0))
                        lhs4_.f2 = rhs3_
                        lhs5_[lhs6_] = rhs4_
                        lhs7_.f5 = rhs5_
                        lhs8_[lhs9_] = rhs6_
                        d_61_v19_: C1
                        nw6_ = C1()
                        nw6_.ctor__(default__.safeModuloInt((0) - (d_53_v12_.f5), d_53_v12_.f5), d_53_v12_.f5)
                        d_61_v19_ = nw6_
                        d_62_v20_: C2
                        nw7_ = C2()
                        nw7_.ctor__((len((_dafny.SeqWithoutIsStrInference([d_43_v3_])).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbblxtt"))), len(_dafny.SeqWithoutIsStrInference([d_43_v3_]))), d_43_v3_))) != ((d_53_v12_).f4), d_46_v5_, d_53_v12_.f5)
                        d_62_v20_ = nw7_
                    d_63_v21_: _dafny.Map
                    d_63_v21_ = _dafny.Map({_dafny.Map({(_dafny.MultiSet([d_43_v3_, False])).cardinality: -512}): d_43_v3_})
                    d_64_v22_: _dafny.Map
                    d_64_v22_ = _dafny.Map({default__.fm3((d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], globalState): d_46_v5_})
                    if ((d_63_v21_)[d_64_v22_] if (d_64_v22_) in (d_63_v21_) else d_43_v3_):
                        d_65_v23_: _dafny.Seq
                        d_65_v23_ = _dafny.SeqWithoutIsStrInference([(d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], d_43_v3_, (d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], (d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], (d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))]])
                        d_66_v24_: _dafny.Map
                        d_66_v24_ = _dafny.Map({default__.fm20((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], d_39_v0_, d_65_v23_, globalState): (d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))]})
                        d_67_v25_: _dafny.Seq
                        d_67_v25_ = _dafny.SeqWithoutIsStrInference([d_39_v0_, d_39_v0_, d_39_v0_])
                        d_66_v24_ = (d_66_v24_).set(d_67_v25_, (d_65_v23_)[default__.safeIndex((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], len(d_65_v23_))])
                        index12_ = default__.safeIndex(675, (d_42_v2_).length(0))
                        (d_42_v2_)[index12_] = False
                        index13_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        index14_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        rhs7_ = (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]
                        rhs8_ = (p0)[default__.safeIndex((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], len(p0))]
                        lhs10_ = d_44_v4_
                        lhs11_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        lhs12_ = d_44_v4_
                        lhs13_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        lhs10_[lhs11_] = rhs7_
                        lhs12_[lhs13_] = rhs8_
                        d_68_v26_: C1
                        nw8_ = C1()
                        nw8_.ctor__((d_46_v5_) * ((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]), (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))])
                        d_68_v26_ = nw8_
                        d_69_v27_: C3
                        nw9_ = C3()
                        nw9_.ctor__((d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))])
                        d_69_v27_ = nw9_
                    elif True:
                        d_39_v0_ = d_39_v0_
                        d_70_v28_: _dafny.MultiSet
                        d_70_v28_ = _dafny.MultiSet([(d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))]])
                        d_71_v29_: _dafny.Seq
                        d_71_v29_ = _dafny.SeqWithoutIsStrInference([(d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))]])
                        index15_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        index16_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        rhs9_ = default__.safeDivisionInt((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], (0) - ((d_70_v28_).cardinality))
                        rhs10_ = (d_71_v29_)[default__.safeIndex(default__.safeModuloInt((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]), len(d_71_v29_))]
                        rhs11_ = 487
                        lhs14_ = d_44_v4_
                        lhs15_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        lhs16_ = globalState
                        lhs17_ = d_44_v4_
                        lhs18_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        lhs14_[lhs15_] = rhs9_
                        lhs16_.f2 = rhs10_
                        lhs17_[lhs18_] = rhs11_
                        d_72_v30_: _dafny.Map
                        d_72_v30_ = _dafny.Map({d_43_v3_: (0) - ((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))])})
                        d_73_v31_: _dafny.Map
                        d_73_v31_ = d_72_v30_
                        d_72_v30_ = ((d_72_v30_).set(d_43_v3_, d_46_v5_)) | ((d_73_v31_))
                        d_74_v32_: str
                        d_74_v32_ = _dafny.CodePoint('f')
                        d_75_v33_: _dafny.Map
                        d_75_v33_ = _dafny.Map({default__.fm4(True, (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))], default__.fm12((d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], d_43_v3_, d_43_v3_, globalState), len(d_39_v0_), globalState): d_74_v32_})
                        d_76_v34_: _dafny.Map
                        d_76_v34_ = _dafny.Map({not(False): d_43_v3_})
                        d_77_v35_: _dafny.MultiSet
                        d_77_v35_ = _dafny.MultiSet([d_75_v33_, default__.fm21(d_46_v5_, ((d_76_v34_)[d_43_v3_] if (d_43_v3_) in (d_76_v34_) else d_43_v3_), len(d_76_v34_), d_46_v5_, globalState), d_75_v33_, _dafny.Map({d_74_v32_: d_74_v32_})])
                        d_78_v36_: C2
                        nw10_ = C2()
                        nw10_.ctor__((d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))], d_46_v5_, d_46_v5_)
                        d_78_v36_ = nw10_
                        d_79_v37_: _dafny.Seq
                        d_79_v37_ = _dafny.SeqWithoutIsStrInference([d_78_v36_])
                        index17_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        rhs12_ = (d_77_v35_) | ((d_77_v35_).set(d_75_v33_, default__.abs(len(d_79_v37_))))
                        rhs13_ = (0) - (d_46_v5_)
                        rhs14_ = (-531) + (((0) - ((0) - (d_46_v5_))) + ((d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]))
                        lhs19_ = d_44_v4_
                        lhs20_ = default__.safeIndex(965, (d_44_v4_).length(0))
                        d_77_v35_ = rhs12_
                        d_46_v5_ = rhs13_
                        lhs19_[lhs20_] = rhs14_
                        d_80_v38_: _dafny.Map
                        d_80_v38_ = _dafny.Map({(D4_DC13((d_42_v2_)[default__.safeIndex(675, (d_42_v2_).length(0))])).cf15: _dafny.MultiSet([_dafny.CodePoint('i')])})
                        d_80_v38_ = (d_80_v38_).set(d_43_v3_, _dafny.MultiSet(d_39_v0_))
                    d_46_v5_ = (d_44_v4_)[default__.safeIndex(965, (d_44_v4_).length(0))]
                    pass
            pass
        (globalState).f2 = False
        d_81_v39_: bool
        d_81_v39_ = False
        r0 = d_81_v39_
        d_82_v40_: _dafny.Array
        nw11_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_82_v40_ = nw11_
        d_82_v40_ = d_82_v40_
        d_83_v41_: int
        d_83_v41_ = -496
        d_84_v42_: _dafny.Set
        d_84_v42_ = _dafny.Set({d_81_v39_, d_81_v39_})
        hi0_ = default__.safeModuloInt(d_83_v41_, len(d_84_v42_))
        for d_85_i3_ in range(d_83_v41_, hi0_):
            d_86_v43_: _dafny.MultiSet
            d_86_v43_ = _dafny.MultiSet([d_81_v39_])
            d_87_v44_: _dafny.Map
            d_87_v44_ = _dafny.Map({(0) - (d_85_i3_): d_86_v43_})
            d_88_v45_: D6
            d_88_v45_ = D6_DC18(d_39_v0_, d_81_v39_, (0) - (d_83_v41_))
            d_89_v46_: _dafny.Set
            d_89_v46_ = _dafny.Set({((d_87_v44_)[(d_88_v45_).cf27] if ((d_88_v45_).cf27) in (d_87_v44_) else d_86_v43_)})
            d_89_v46_ = ((d_89_v46_) | (_dafny.Set({d_86_v43_}))) - (d_89_v46_)
            d_90_v47_: _dafny.Array
            nw12_ = _dafny.Array(None, 26)
            d_90_v47_ = nw12_
            rhs15_ = d_90_v47_
            rhs16_ = (d_81_v39_ if not(d_81_v39_) else d_81_v39_)
            lhs21_ = globalState
            d_90_v47_ = rhs15_
            lhs21_.f2 = rhs16_
            d_91_v48_: _dafny.Map
            d_91_v48_ = _dafny.Map({d_83_v41_: d_39_v0_})
            d_92_v49_: str
            d_92_v49_ = _dafny.CodePoint('l')
            d_91_v48_ = (d_91_v48_).set(d_85_i3_, (d_39_v0_).set(default__.safeIndex(d_83_v41_, len(d_39_v0_)), d_92_v49_))
            d_93_v50_: _dafny.Array
            nw13_ = _dafny.Array(int(0), 15)
            d_93_v50_ = nw13_
            d_94_v51_: _dafny.Seq
            d_94_v51_ = _dafny.SeqWithoutIsStrInference([d_81_v39_, d_81_v39_])
            index18_ = default__.safeIndex(548, (d_93_v50_).length(0))
            (d_93_v50_)[index18_] = (len(d_94_v51_)) - (d_85_i3_)
            d_95_v52_: _dafny.MultiSet
            d_95_v52_ = _dafny.MultiSet([_dafny.CodePoint('p')])
            d_96_v53_: D3
            d_96_v53_ = D3_DC9(len(p0), default__.fm12(d_81_v39_, d_81_v39_, d_81_v39_, globalState), (d_95_v52_).cardinality, d_81_v39_)
            index19_ = default__.safeIndex(548, (d_93_v50_).length(0))
            def iife7_(_pat_let0_0):
                def iife8_(d_97_dt__update__tmp_h0_):
                    def iife9_(_pat_let1_0):
                        def iife10_(d_98_dt__update_hcf11_h0_):
                            return D3_DC9((d_97_dt__update__tmp_h0_).cf8, (d_97_dt__update__tmp_h0_).cf9, (d_97_dt__update__tmp_h0_).cf10, d_98_dt__update_hcf11_h0_)
                        return iife10_(_pat_let1_0)
                    return iife9_(False)
                return iife8_(_pat_let0_0)
            (d_93_v50_)[index19_] = (iife7_(d_96_v53_)).cf8
        (globalState).f2 = d_81_v39_
        d_99_v54_: _dafny.Map
        d_99_v54_ = _dafny.Map({d_83_v41_: not(d_81_v39_)})
        d_100_v55_: D0
        d_100_v55_ = D0_DC0(len(d_99_v54_))
        pat_let_tv0_ = d_81_v39_
        pat_let_tv1_ = d_81_v39_
        pat_let_tv2_ = d_81_v39_
        pat_let_tv3_ = d_83_v41_
        def lambda1_(source2_):
            if source2_.is_DC1:
                d_101___mcc_h0_ = source2_.cf1
                d_102___mcc_h1_ = source2_.cf2
                d_103_cf2_ = d_102___mcc_h1_
                d_104_cf1_ = d_101___mcc_h0_
                return pat_let_tv0_
            elif source2_.is_DC2:
                return pat_let_tv1_
            elif True:
                d_105___mcc_h2_ = source2_.cf0
                d_106_cf0_ = d_105___mcc_h2_
                return pat_let_tv2_

        def iife11_(_pat_let2_0):
            def iife12_(d_107_dt__update__tmp_h1_):
                def iife13_(_pat_let3_0):
                    def iife14_(d_108_dt__update_hcf0_h0_):
                        return D0_DC0(d_108_dt__update_hcf0_h0_)
                    return iife14_(_pat_let3_0)
                return iife13_(pat_let_tv3_)
            return iife12_(_pat_let2_0)
        r0 = lambda1_(iife11_(d_100_v55_))
        d_109_v56_: str
        d_109_v56_ = _dafny.CodePoint('e')
        d_110_v57_: _dafny.Map
        d_110_v57_ = _dafny.Map({(((d_39_v0_) + (d_39_v0_)).set(default__.safeIndex(d_83_v41_, len((d_39_v0_) + (d_39_v0_))), d_109_v56_)).set(default__.safeIndex(d_83_v41_, len(((d_39_v0_) + (d_39_v0_)).set(default__.safeIndex(d_83_v41_, len((d_39_v0_) + (d_39_v0_))), d_109_v56_))), d_109_v56_): default__.fm3(d_81_v39_, globalState)})
        r1 = d_110_v57_
        return r0, r1

    @staticmethod
    def Main(noArgsParameter__):
        d_111_globalState_: GlobalState
        nw14_ = GlobalState()
        nw14_.ctor__(True, False, True)
        d_111_globalState_ = nw14_
        d_112_v0_: str
        d_112_v0_ = _dafny.CodePoint('h')
        d_113_v1_: _dafny.MultiSet
        d_113_v1_ = _dafny.MultiSet([d_112_v0_, d_112_v0_, d_112_v0_, d_112_v0_, d_112_v0_])
        d_114_v2_: C0
        nw15_ = C0()
        nw15_.ctor__(d_113_v1_)
        d_114_v2_ = nw15_
        d_115_v3_: int
        d_115_v3_ = 295
        d_115_v3_ = -786
        d_116_v4_: _dafny.Array
        nw16_ = _dafny.Array(False, 29)
        d_116_v4_ = nw16_
        d_117_v5_: bool
        d_117_v5_ = True
        index20_ = default__.safeIndex(831, (d_116_v4_).length(0))
        (d_116_v4_)[index20_] = d_117_v5_
        index21_ = default__.safeIndex(831, (d_116_v4_).length(0))
        (d_116_v4_)[index21_] = (d_115_v3_) >= (d_115_v3_)
        d_118_v6_: _dafny.Array
        nw17_ = _dafny.Array(D4.default()(), 21)
        d_118_v6_ = nw17_
        d_119_v7_: _dafny.Set
        d_119_v7_ = _dafny.Set({(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]})
        index22_ = default__.safeIndex(885, (d_118_v6_).length(0))
        (d_118_v6_)[index22_] = D4_DC12(d_119_v7_)
        d_120_v8_: D4
        d_120_v8_ = D4_DC12(d_119_v7_)
        index23_ = default__.safeIndex(885, (d_118_v6_).length(0))
        (d_118_v6_)[index23_] = d_120_v8_
        d_121_i0_: int
        d_121_i0_ = 0
        with _dafny.label("1"):
            while False:
                with _dafny.c_label("1"):
                    if (d_121_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_121_i0_ = (d_121_i0_) + (1)
                    d_122_v9_: _dafny.Array
                    def lambda2_(d_123_v0_):
                        def lambda3_(d_124_i1_):
                            return d_123_v0_

                        return lambda3_

                    init1_ = lambda2_(d_112_v0_)
                    nw18_ = _dafny.Array(None, 18)
                    for i0_1_ in range(nw18_.length(0)):
                        nw18_[i0_1_] = init1_(i0_1_)
                    d_122_v9_ = nw18_
                    index24_ = default__.safeIndex(45, (d_122_v9_).length(0))
                    (d_122_v9_)[index24_] = _dafny.CodePoint('h')
                    d_125_v10_: _dafny.Map
                    d_125_v10_ = _dafny.Map({d_115_v3_: d_115_v3_})
                    d_126_v11_: _dafny.Seq
                    d_126_v11_ = _dafny.SeqWithoutIsStrInference([len(d_125_v10_)])
                    d_127_v12_: _dafny.Seq
                    d_127_v12_ = _dafny.SeqWithoutIsStrInference([d_117_v5_, d_117_v5_])
                    index25_ = default__.safeIndex(45, (d_122_v9_).length(0))
                    rhs17_ = (_dafny.CodePoint('t') if (d_127_v12_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_115_v3_, d_115_v3_])), len(d_127_v12_))] else d_112_v0_)
                    rhs18_ = (d_126_v11_) + (d_126_v11_)
                    rhs19_ = (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]
                    lhs22_ = d_122_v9_
                    lhs23_ = default__.safeIndex(45, (d_122_v9_).length(0))
                    lhs22_[lhs23_] = rhs17_
                    d_126_v11_ = rhs18_
                    d_117_v5_ = rhs19_
                    d_128_v13_: _dafny.Seq
                    d_128_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mpqme"))
                    d_128_v13_ = d_128_v13_
                    d_129_v14_: D2
                    d_129_v14_ = D2_DC7()
                    d_129_v14_ = d_129_v14_
                    d_130_v15_: _dafny.Map
                    d_130_v15_ = _dafny.Map({(d_126_v11_)[default__.safeIndex(d_115_v3_, len(d_126_v11_))]: d_117_v5_})
                    (d_111_globalState_).f2 = not(((d_130_v15_ if d_117_v5_ else d_130_v15_)) != (_dafny.Map({d_115_v3_: d_117_v5_})))
                    pass
            pass
        d_119_v7_ = (d_119_v7_) | (d_119_v7_)
        hi1_ = len(d_119_v7_)
        for d_131_i2_ in range(d_115_v3_, hi1_):
            d_132_v16_: _dafny.MultiSet
            d_132_v16_ = _dafny.MultiSet([default__.fm17(d_111_globalState_), d_113_v1_])
            if not(not((((d_132_v16_)[(d_114_v2_).f7] if ((d_114_v2_).f7) in (d_132_v16_) else d_115_v3_)) < (d_131_i2_))):
                d_133_v17_: _dafny.Seq
                d_133_v17_ = _dafny.SeqWithoutIsStrInference([d_131_i2_, d_131_i2_])
                d_134_v18_: _dafny.Seq
                d_134_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqlglyph"))
                (d_111_globalState_).f2 = ((d_133_v17_)[default__.safeIndex(len(d_134_v18_), len(d_133_v17_))]) <= (d_131_i2_)
                d_135_v19_: C3
                nw19_ = C3()
                nw19_.ctor__((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])
                d_135_v19_ = nw19_
                d_136_v20_: D6
                d_136_v20_ = D6_DC17(d_135_v19_)
                d_135_v19_ = (d_136_v20_).cf24
                (d_111_globalState_).f2 = d_117_v5_
                d_137_v21_: _dafny.Array
                nw20_ = _dafny.Array(int(0), 14)
                d_137_v21_ = nw20_
                index26_ = default__.safeIndex(535, (d_137_v21_).length(0))
                (d_137_v21_)[index26_] = len(d_134_v18_)
                index27_ = default__.safeIndex(535, (d_137_v21_).length(0))
                (d_137_v21_)[index27_] = default__.fm3((d_115_v3_) <= (len(_dafny.SeqWithoutIsStrInference([d_134_v18_ for d_138_i3_ in range(default__.abs(946))]))), d_111_globalState_)
                d_139_v22_: _dafny.Map
                d_139_v22_ = _dafny.Map({d_112_v0_: d_135_v19_})
                d_139_v22_ = d_139_v22_
            elif True:
                d_140_v23_: _dafny.Seq
                d_140_v23_ = _dafny.SeqWithoutIsStrInference([867])
                d_141_v24_: _dafny.Map
                d_141_v24_ = _dafny.Map({d_131_i2_: d_115_v3_})
                d_142_v25_: _dafny.Map
                d_142_v25_ = _dafny.Map({(d_140_v23_)[default__.safeIndex(len(d_141_v24_), len(d_140_v23_))]: d_116_v4_})
                d_143_v26_: _dafny.Seq
                d_143_v26_ = _dafny.SeqWithoutIsStrInference([d_116_v4_])
                d_116_v4_ = ((d_142_v25_)[d_131_i2_] if (d_131_i2_) in (d_142_v25_) else (d_143_v26_)[default__.safeIndex(default__.fm3((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_111_globalState_), len(d_143_v26_))])
                d_144_v27_: _dafny.Map
                d_144_v27_ = _dafny.Map({d_117_v5_: d_112_v0_})
                d_144_v27_ = (d_144_v27_).set(not((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]), d_112_v0_)
                d_145_v28_: _dafny.Array
                nw21_ = _dafny.Array(int(0), 10)
                d_145_v28_ = nw21_
                index28_ = default__.safeIndex(436, (d_145_v28_).length(0))
                (d_145_v28_)[index28_] = (d_115_v3_) + (d_131_i2_)
                d_146_v29_: _dafny.Seq
                d_146_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oasbjj"))
                d_147_v30_: _dafny.Map
                d_147_v30_ = _dafny.Map({d_146_v29_: 692})
                index29_ = default__.safeIndex(436, (d_145_v28_).length(0))
                (d_145_v28_)[index29_] = ((d_147_v30_)[d_146_v29_] if (d_146_v29_) in (d_147_v30_) else d_131_i2_)
                rhs20_ = (d_146_v29_)[default__.safeIndex(d_115_v3_, len(d_146_v29_))]
                rhs21_ = d_131_i2_
                rhs22_ = ((0) - (d_131_i2_)) - (d_131_i2_)
                d_112_v0_ = rhs20_
                d_115_v3_ = rhs21_
                d_115_v3_ = rhs22_
                (d_111_globalState_).f2 = (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]
            d_148_v31_: _dafny.Seq
            d_148_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvpces"))
            d_149_v32_: D6
            d_149_v32_ = D6_DC18(d_148_v31_, d_117_v5_, d_131_i2_)
            d_149_v32_ = d_149_v32_
            d_150_v33_: D3
            d_150_v33_ = D3_DC8(d_148_v31_)
            d_148_v31_ = ((d_150_v33_).cf7) + ((default__.fm6((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], len(_dafny.Map({d_117_v5_: d_131_i2_})), (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_111_globalState_)) + (d_148_v31_))
            d_151_v34_: _dafny.Array
            def lambda4_(d_152_v4_):
                def lambda5_(d_153_i4_):
                    return D2_DC6(_dafny.SeqWithoutIsStrInference([(d_152_v4_)[default__.safeIndex(831, (d_152_v4_).length(0))]]))

                return lambda5_

            init2_ = lambda4_(d_116_v4_)
            nw22_ = _dafny.Array(None, 24)
            for i0_2_ in range(nw22_.length(0)):
                nw22_[i0_2_] = init2_(i0_2_)
            d_151_v34_ = nw22_
            d_154_v35_: _dafny.Map
            d_154_v35_ = _dafny.Map({d_151_v34_: not(d_117_v5_)})
            d_155_v36_: _dafny.Array
            nw23_ = _dafny.Array(int(0), 5)
            d_155_v36_ = nw23_
            d_156_v37_: _dafny.Map
            d_156_v37_ = _dafny.Map({d_155_v36_: d_154_v35_})
            d_154_v35_ = ((d_156_v37_)[d_155_v36_] if (d_155_v36_) in (d_156_v37_) else d_154_v35_)
        d_115_v3_ = d_115_v3_
        d_115_v3_ = d_115_v3_
        d_157_v38_: _dafny.Seq
        d_157_v38_ = _dafny.SeqWithoutIsStrInference([d_115_v3_])
        d_158_v39_: _dafny.Array
        nw24_ = _dafny.Array(None, 5)
        nw24_[int(0)] = d_115_v3_
        nw24_[int(1)] = 134
        nw24_[int(2)] = d_115_v3_
        nw24_[int(3)] = (0) - (d_115_v3_)
        nw24_[int(4)] = len(d_157_v38_)
        d_158_v39_ = nw24_
        d_159_v40_: _dafny.Seq
        d_159_v40_ = _dafny.SeqWithoutIsStrInference([d_158_v39_])
        d_159_v40_ = d_159_v40_
        d_160_v41_: _dafny.Seq
        d_160_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oayxq"))
        d_161_v42_: D0
        d_161_v42_ = D0_DC2()
        pat_let_tv4_ = d_160_v41_
        pat_let_tv5_ = d_160_v41_
        pat_let_tv6_ = d_160_v41_
        pat_let_tv7_ = d_160_v41_
        def lambda6_(source3_):
            if source3_.is_DC1:
                d_162___mcc_h0_ = source3_.cf1
                d_163___mcc_h1_ = source3_.cf2
                d_164_cf2_ = d_163___mcc_h1_
                d_165_cf1_ = d_162___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))) + (pat_let_tv4_)
            elif source3_.is_DC2:
                return (pat_let_tv5_) + (pat_let_tv6_)
            elif True:
                d_166___mcc_h2_ = source3_.cf0
                d_167_cf0_ = d_166___mcc_h2_
                return pat_let_tv7_

        d_160_v41_ = lambda6_(d_161_v42_)
        index30_ = default__.safeIndex(831, (d_116_v4_).length(0))
        (d_116_v4_)[index30_] = not((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])
        if d_117_v5_:
            d_168_v43_: _dafny.Map
            d_168_v43_ = _dafny.Map({d_115_v3_: d_158_v39_})
            d_158_v39_ = ((d_168_v43_)[default__.safeDivisionInt(d_115_v3_, d_115_v3_)] if (default__.safeDivisionInt(d_115_v3_, d_115_v3_)) in (d_168_v43_) else d_158_v39_)
            d_169_v44_: C3
            nw25_ = C3()
            nw25_.ctor__((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])
            d_169_v44_ = nw25_
            (d_111_globalState_).f2 = (d_169_v44_).f3
            d_170_v45_: _dafny.Seq
            d_170_v45_ = _dafny.SeqWithoutIsStrInference([(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], True])
            d_171_v46_: _dafny.Array
            def lambda7_(d_172_v38_):
                def lambda8_(d_173_i5_):
                    return d_172_v38_

                return lambda8_

            init3_ = lambda7_(d_157_v38_)
            nw26_ = _dafny.Array(None, 8)
            for i0_3_ in range(nw26_.length(0)):
                nw26_[i0_3_] = init3_(i0_3_)
            d_171_v46_ = nw26_
            rhs23_ = d_170_v45_
            rhs24_ = d_171_v46_
            d_170_v45_ = rhs23_
            d_171_v46_ = rhs24_
            d_174_v47_: D1
            d_174_v47_ = D1_DC4(True)
            d_175_v48_: D1
            d_175_v48_ = D1_DC5(d_174_v47_)
            d_176_v49_: D1
            d_176_v49_ = D1_DC5(d_174_v47_)
            d_177_v50_: D1
            d_177_v50_ = D1_DC5(d_175_v48_)
            d_178_v51_: C1
            nw27_ = C1()
            nw27_.ctor__(d_115_v3_, len(_dafny.Map({d_115_v3_: d_117_v5_})))
            d_178_v51_ = nw27_
            d_179_v52_: _dafny.Seq
            d_179_v52_ = _dafny.SeqWithoutIsStrInference([d_178_v51_, d_178_v51_, d_178_v51_, d_178_v51_])
            d_177_v50_ = default__.fm18(len(d_179_v52_), d_111_globalState_)
        elif True:
            d_180_v53_: _dafny.Array
            nw28_ = _dafny.Array(_dafny.Map({}), 12)
            d_180_v53_ = nw28_
            d_181_v54_: T0
            nw29_ = C2()
            nw29_.ctor__(True, d_115_v3_, d_115_v3_)
            d_181_v54_ = nw29_
            d_182_v55_: _dafny.Map
            d_182_v55_ = _dafny.Map({d_181_v54_: True})
            index31_ = default__.safeIndex(151, (d_180_v53_).length(0))
            (d_180_v53_)[index31_] = d_182_v55_
            index32_ = default__.safeIndex(151, (d_180_v53_).length(0))
            (d_180_v53_)[index32_] = d_182_v55_
            d_160_v41_ = d_160_v41_
            d_183_v56_: _dafny.Map
            d_183_v56_ = _dafny.Map({(d_160_v41_) == (_dafny.SeqWithoutIsStrInference([d_112_v0_ for d_184_i6_ in range(default__.abs(-346))])): d_117_v5_})
            index33_ = default__.safeIndex(831, (d_116_v4_).length(0))
            (d_116_v4_)[index33_] = ((d_183_v56_)[(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]] if ((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]) in (d_183_v56_) else d_117_v5_)
            d_185_v57_: _dafny.Seq
            d_185_v57_ = _dafny.SeqWithoutIsStrInference([(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_117_v5_, (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_117_v5_, (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]])
            d_186_v58_: _dafny.Map
            d_186_v58_ = _dafny.Map({d_119_v7_: (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]})
            if (True if (d_185_v57_)[default__.safeIndex((d_181_v54_).f4, len(d_185_v57_))] else ((d_186_v58_)[default__.fm9((d_181_v54_).f4, d_117_v5_, len(d_160_v41_), d_111_globalState_)] if (default__.fm9((d_181_v54_).f4, d_117_v5_, len(d_160_v41_), d_111_globalState_)) in (d_186_v58_) else d_117_v5_)):
                d_187_v59_: _dafny.MultiSet
                d_187_v59_ = _dafny.MultiSet([d_158_v39_, d_158_v39_, d_158_v39_, d_158_v39_])
                d_188_v60_: _dafny.Seq
                d_188_v60_ = _dafny.SeqWithoutIsStrInference([d_187_v59_, d_187_v59_, d_187_v59_])
                index34_ = default__.safeIndex(831, (d_116_v4_).length(0))
                rhs25_ = d_112_v0_
                rhs26_ = (d_188_v60_)[default__.safeIndex(((d_181_v54_).f4) + ((d_181_v54_).f4), len(d_188_v60_))]
                rhs27_ = not (d_117_v5_) or ((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])
                rhs28_ = d_117_v5_
                rhs29_ = d_181_v54_.f5
                lhs24_ = d_116_v4_
                lhs25_ = default__.safeIndex(831, (d_116_v4_).length(0))
                lhs26_ = d_181_v54_
                d_112_v0_ = rhs25_
                d_187_v59_ = rhs26_
                d_117_v5_ = rhs27_
                lhs24_[lhs25_] = rhs28_
                lhs26_.f5 = rhs29_
                d_189_v61_: C3
                nw30_ = C3()
                nw30_.ctor__((d_185_v57_)[default__.safeIndex(len(d_185_v57_), len(d_185_v57_))])
                d_189_v61_ = nw30_
                d_189_v61_ = d_189_v61_
                d_115_v3_ = (0) - (d_115_v3_)
                d_190_v62_: _dafny.Map
                d_190_v62_ = _dafny.Map({(d_189_v61_).f3: 372})
                (d_181_v54_).f5 = ((d_190_v62_)[d_117_v5_] if (d_117_v5_) in (d_190_v62_) else (d_181_v54_).f4)
                d_191_v63_: _dafny.MultiSet
                d_191_v63_ = _dafny.MultiSet([d_157_v38_, (d_157_v38_) + (d_157_v38_), _dafny.SeqWithoutIsStrInference([(D0_DC1(len(_dafny.Map({(d_181_v54_).f4: d_115_v3_})), d_181_v54_.f5)).cf2 for d_192_i7_ in range(default__.abs(641))])])
                (d_181_v54_).f5 = ((d_191_v63_)[(d_157_v38_ if (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))] else _dafny.SeqWithoutIsStrInference([(d_181_v54_).f4, d_115_v3_]))] if ((d_157_v38_ if (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))] else _dafny.SeqWithoutIsStrInference([(d_181_v54_).f4, d_115_v3_]))) in (d_191_v63_) else d_181_v54_.f5)
            elif True:
                index35_ = default__.safeIndex(831, (d_116_v4_).length(0))
                (d_116_v4_)[index35_] = default__.fm12(False, (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_117_v5_, d_111_globalState_)
                (d_181_v54_).f5 = 605
                index36_ = default__.safeIndex(491, (d_158_v39_).length(0))
                (d_158_v39_)[index36_] = (d_181_v54_).f4
                index37_ = default__.safeIndex(491, (d_158_v39_).length(0))
                (d_158_v39_)[index37_] = d_181_v54_.f5
                index38_ = default__.safeIndex(491, (d_158_v39_).length(0))
                (d_158_v39_)[index38_] = (d_115_v3_) + (d_115_v3_)
                d_193_v64_: _dafny.Array
                nw31_ = _dafny.Array(None, 5)
                nw31_[int(0)] = d_115_v3_
                nw31_[int(1)] = len(d_157_v38_)
                nw31_[int(2)] = 745
                nw31_[int(3)] = (d_181_v54_).f4
                nw31_[int(4)] = default__.safeModuloInt(728, (d_181_v54_).f4)
                d_193_v64_ = nw31_
                d_194_v65_: _dafny.Array
                nw32_ = _dafny.Array(None, 10)
                d_194_v65_ = nw32_
                d_195_v66_: _dafny.Set
                d_195_v66_ = _dafny.Set({d_194_v65_})
                index39_ = default__.safeIndex(491, (d_158_v39_).length(0))
                rhs30_ = d_193_v64_
                rhs31_ = d_115_v3_
                rhs32_ = not((_dafny.Set({d_194_v65_})).issubset(d_195_v66_))
                lhs27_ = d_158_v39_
                lhs28_ = default__.safeIndex(491, (d_158_v39_).length(0))
                lhs29_ = d_111_globalState_
                d_193_v64_ = rhs30_
                lhs27_[lhs28_] = rhs31_
                lhs29_.f2 = rhs32_
            (d_181_v54_).f5 = d_181_v54_.f5
        source4_ = d_161_v42_
        if source4_.is_DC1:
            d_196___mcc_h3_ = source4_.cf1
            d_197___mcc_h4_ = source4_.cf2
            d_198_cf2_ = d_197___mcc_h4_
            d_199_cf1_ = d_196___mcc_h3_
            d_200_v67_: _dafny.Map
            d_200_v67_ = _dafny.Map({(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]: (default__.fm12((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], not(d_117_v5_), d_117_v5_, d_111_globalState_)) in (_dafny.SeqWithoutIsStrInference([d_117_v5_]))})
            d_201_v68_: C2
            nw33_ = C2()
            nw33_.ctor__(d_117_v5_, d_198_cf2_, d_198_cf2_)
            d_201_v68_ = nw33_
            d_202_v69_: _dafny.Map
            d_202_v69_ = _dafny.Map({d_115_v3_: d_201_v68_})
            d_200_v67_ = (d_200_v67_).set(not(not((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])), not((len(d_157_v38_)) >= (len(d_202_v69_))))
            d_198_cf2_ = d_198_cf2_
            d_203_v70_: C1
            nw34_ = C1()
            nw34_.ctor__(default__.fm3((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_111_globalState_), d_115_v3_)
            d_203_v70_ = nw34_
            d_203_v70_ = d_203_v70_
            d_204_v71_: _dafny.MultiSet
            d_204_v71_ = _dafny.MultiSet([d_117_v5_])
            d_205_v72_: D6
            d_205_v72_ = D6_DC18(d_160_v41_, d_117_v5_, ((d_204_v71_).set(d_117_v5_, default__.abs(d_115_v3_))).cardinality)
            d_206_v73_: D6
            d_206_v73_ = D6_DC19(d_205_v72_)
            d_206_v73_ = d_206_v73_
        elif source4_.is_DC2:
            d_207_v74_: C1
            nw35_ = C1()
            nw35_.ctor__((0) - (d_115_v3_), 933)
            d_207_v74_ = nw35_
            d_207_v74_ = d_207_v74_
            d_115_v3_ = (d_115_v3_) + (d_115_v3_)
            d_208_v75_: _dafny.Array
            nw36_ = _dafny.Array(None, 4)
            d_208_v75_ = nw36_
            index40_ = default__.safeIndex(348, (d_208_v75_).length(0))
            (d_208_v75_)[index40_] = d_114_v2_
            index41_ = default__.safeIndex(348, (d_208_v75_).length(0))
            (d_208_v75_)[index41_] = d_114_v2_
            d_209_v76_: _dafny.Array
            nw37_ = _dafny.Array(None, 14)
            d_209_v76_ = nw37_
            d_209_v76_ = d_209_v76_
        elif True:
            d_210___mcc_h5_ = source4_.cf0
            d_211_cf0_ = d_210___mcc_h5_
            d_212_v77_: _dafny.Map
            d_212_v77_ = _dafny.Map({d_115_v3_: d_115_v3_})
            d_213_v78_: _dafny.Map
            d_213_v78_ = _dafny.Map({len(d_212_v77_): d_115_v3_})
            d_214_v79_: D0
            d_214_v79_ = D0_DC0(len(d_213_v78_))
            d_215_v80_: _dafny.Seq
            d_215_v80_ = _dafny.SeqWithoutIsStrInference([d_214_v79_])
            d_215_v80_ = ((d_215_v80_) + (d_215_v80_)).set(default__.safeIndex((d_211_cf0_) + (d_211_cf0_), len((d_215_v80_) + (d_215_v80_))), d_214_v79_)
            d_216_v81_: T0
            nw38_ = C1()
            nw38_.ctor__((0) - (default__.fm3((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_111_globalState_)), len(d_119_v7_))
            d_216_v81_ = nw38_
            d_217_v82_: _dafny.Seq
            d_217_v82_ = _dafny.SeqWithoutIsStrInference([d_216_v81_])
            d_218_v83_: _dafny.Seq
            d_218_v83_ = _dafny.SeqWithoutIsStrInference([(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_117_v5_])
            d_219_v84_: _dafny.Set
            d_219_v84_ = _dafny.Set({d_218_v83_})
            d_220_v85_: _dafny.Map
            d_220_v85_ = _dafny.Map({default__.safeDivisionInt(629, (0) - (d_115_v3_)): (d_217_v82_)[default__.safeIndex(len(d_219_v84_), len(d_217_v82_))]})
            d_220_v85_ = ((d_220_v85_) | (d_220_v85_)) | (d_220_v85_)
            (d_111_globalState_).f2 = False
            d_221_v86_: bool
            d_222_v87_: bool
            d_223_v88_: bool
            out0_: bool
            out1_: bool
            out2_: bool
            out0_, out1_, out2_ = (d_216_v81_).m1((d_158_v39_ if (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))] else d_158_v39_), _dafny.CodePoint('u'), d_111_globalState_)
            d_221_v86_ = out0_
            d_222_v87_ = out1_
            d_223_v88_ = out2_
        source5_ = D3_DC8(d_160_v41_)
        if source5_.is_DC9:
            d_224___mcc_h6_ = source5_.cf8
            d_225___mcc_h7_ = source5_.cf9
            d_226___mcc_h8_ = source5_.cf10
            d_227___mcc_h9_ = source5_.cf11
            d_228_cf11_ = d_227___mcc_h9_
            d_229_cf10_ = d_226___mcc_h8_
            d_230_cf9_ = d_225___mcc_h7_
            d_231_cf8_ = d_224___mcc_h6_
            d_230_cf9_ = (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]
            d_232_v89_: _dafny.MultiSet
            d_232_v89_ = _dafny.MultiSet([d_231_cf8_, -952])
            d_233_v90_: _dafny.Set
            d_233_v90_ = _dafny.Set({((d_232_v89_)[d_231_cf8_] if (d_231_cf8_) in (d_232_v89_) else d_115_v3_), d_229_cf10_, -994, d_115_v3_})
            d_234_v91_: bool
            d_235_v92_: _dafny.Map
            out3_: bool
            out4_: _dafny.Map
            out3_, out4_ = default__.m5((_dafny.SeqWithoutIsStrInference([d_229_cf10_])) + (d_157_v38_), d_233_v90_, d_111_globalState_)
            d_234_v91_ = out3_
            d_235_v92_ = out4_
            d_236_v94_: bool
            d_237_v95_: _dafny.Map
            out5_: bool
            out6_: _dafny.Map
            def iife15_():
                coll7_ = _dafny.Set()
                compr_7_: int
                for compr_7_ in _dafny.IntegerRange(-340, 994):
                    d_238_v93_: int = compr_7_
                    if ((-340) <= (d_238_v93_)) and ((d_238_v93_) < (994)):
                        coll7_ = coll7_.union(_dafny.Set([(d_238_v93_) - (d_231_cf8_)]))
                return _dafny.Set(coll7_)
            out5_, out6_ = default__.m5(d_157_v38_, iife15_()
            , d_111_globalState_)
            d_236_v94_ = out5_
            d_237_v95_ = out6_
            if not(d_234_v91_):
                d_239_v96_: bool
                d_240_v97_: _dafny.Map
                out7_: bool
                out8_: _dafny.Map
                out7_, out8_ = default__.m5(d_157_v38_, d_233_v90_, d_111_globalState_)
                d_239_v96_ = out7_
                d_240_v97_ = out8_
                rhs33_ = _dafny.SeqWithoutIsStrInference([(d_158_v39_ if d_230_cf9_ else d_158_v39_), d_158_v39_, d_158_v39_, d_158_v39_])
                rhs34_ = d_115_v3_
                d_159_v40_ = rhs33_
                d_231_cf8_ = rhs34_
                (d_111_globalState_).f2 = not (d_228_cf11_) or (d_117_v5_)
                d_241_v98_: D1
                d_241_v98_ = D1_DC3(d_230_cf9_)
                pat_let_tv8_ = d_116_v4_
                pat_let_tv9_ = d_116_v4_
                d_242_v99_: _dafny.MultiSet
                def iife16_(_pat_let4_0):
                    def iife17_(d_243_dt__update__tmp_h0_):
                        def iife18_(_pat_let5_0):
                            def iife19_(d_244_dt__update_hcf3_h0_):
                                return D1_DC3(d_244_dt__update_hcf3_h0_)
                            return iife19_(_pat_let5_0)
                        return iife18_((pat_let_tv9_)[default__.safeIndex(831, (pat_let_tv8_).length(0))])
                    return iife17_(_pat_let4_0)
                d_242_v99_ = _dafny.MultiSet([default__.fm19(d_230_cf9_, not((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]), d_111_globalState_), d_241_v98_, d_241_v98_, d_241_v98_, iife16_(d_241_v98_)])
                d_245_v100_: _dafny.Array
                nw39_ = _dafny.Array(None, 6)
                nw39_[int(0)] = d_242_v99_
                nw39_[int(1)] = d_242_v99_
                nw39_[int(2)] = (d_242_v99_).intersection((_dafny.MultiSet([d_241_v98_])).set(D1_DC3(d_236_v94_), default__.abs(d_229_cf10_)))
                nw39_[int(3)] = d_242_v99_
                nw39_[int(4)] = d_242_v99_
                nw39_[int(5)] = d_242_v99_
                d_245_v100_ = nw39_
                d_245_v100_ = d_245_v100_
                d_246_v101_: D7
                d_246_v101_ = D7_DC20(d_157_v38_)
                index42_ = default__.safeIndex(831, (d_116_v4_).length(0))
                rhs35_ = ((d_246_v101_).cf29) <= (d_157_v38_)
                rhs36_ = (default__.fm6(d_117_v5_, d_229_cf10_, (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_111_globalState_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygn")))
                rhs37_ = not((d_229_cf10_) > (d_115_v3_))
                lhs30_ = d_111_globalState_
                lhs31_ = d_116_v4_
                lhs32_ = default__.safeIndex(831, (d_116_v4_).length(0))
                lhs30_.f2 = rhs35_
                d_160_v41_ = rhs36_
                lhs31_[lhs32_] = rhs37_
            elif True:
                d_247_v102_: _dafny.Seq
                d_247_v102_ = _dafny.SeqWithoutIsStrInference([d_160_v41_])
                d_248_v103_: _dafny.MultiSet
                d_248_v103_ = _dafny.MultiSet([(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]])
                d_249_v104_: D6
                d_249_v104_ = D6_DC18(d_160_v41_, True, (d_248_v103_).cardinality)
                d_250_v105_: D6
                d_250_v105_ = D6_DC18(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")), d_234_v91_, (d_249_v104_).cf27)
                d_251_v106_: _dafny.Seq
                d_251_v106_ = _dafny.SeqWithoutIsStrInference([d_160_v41_, (d_247_v102_)[default__.safeIndex(d_115_v3_, len(d_247_v102_))], ((d_249_v104_).cf25) + (d_160_v41_)])
                d_247_v102_ = d_251_v106_
                d_252_v107_: D0
                d_252_v107_ = D0_DC0(d_115_v3_)
                d_231_cf8_ = (d_231_cf8_) + ((d_252_v107_).cf0)
                d_253_v108_: C3
                nw40_ = C3()
                nw40_.ctor__(d_228_cf11_)
                d_253_v108_ = nw40_
                d_254_v109_: T0
                nw41_ = C2()
                nw41_.ctor__(True, (d_232_v89_).cardinality, default__.safeModuloInt(len(d_160_v41_), d_229_cf10_))
                d_254_v109_ = nw41_
                d_255_v110_: C2
                nw42_ = C2()
                nw42_.ctor__((d_253_v108_).f3, ((d_254_v109_).f4) * (d_229_cf10_), 685)
                d_255_v110_ = nw42_
                d_256_v111_: _dafny.Map
                d_256_v111_ = _dafny.Map({d_228_cf11_: (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]})
                rhs38_ = (d_256_v111_) == ((_dafny.Map({d_117_v5_: d_236_v94_})) | (d_256_v111_))
                rhs39_ = (d_253_v108_ if d_234_v91_ else d_253_v108_)
                rhs40_ = d_254_v109_
                rhs41_ = d_255_v110_
                lhs33_ = d_111_globalState_
                lhs33_.f2 = rhs38_
                d_253_v108_ = rhs39_
                d_254_v109_ = rhs40_
                d_255_v110_ = rhs41_
                d_256_v111_ = default__.fm14(543, d_111_globalState_)
                d_230_cf9_ = (d_253_v108_).f3
        elif source5_.is_DC10:
            d_257___mcc_h10_ = source5_.cf12
            d_258_cf12_ = d_257___mcc_h10_
            index43_ = default__.safeIndex(831, (d_116_v4_).length(0))
            (d_116_v4_)[index43_] = not((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])
            d_259_v112_: C2
            nw43_ = C2()
            nw43_.ctor__((d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_115_v3_, d_115_v3_)
            d_259_v112_ = nw43_
            d_260_v113_: C3
            nw44_ = C3()
            nw44_.ctor__((d_259_v112_).f6)
            d_260_v113_ = nw44_
            d_261_v114_: _dafny.Map
            d_261_v114_ = _dafny.Map({(d_259_v112_).f6: d_260_v113_})
            d_262_v115_: _dafny.Map
            d_262_v115_ = _dafny.Map({-744: (d_259_v112_).f6})
            d_263_v116_: _dafny.Map
            d_263_v116_ = _dafny.Map({((d_261_v114_)[d_117_v5_] if (d_117_v5_) in (d_261_v114_) else d_260_v113_): d_262_v115_})
            d_264_v118_: _dafny.Seq
            d_264_v118_ = _dafny.SeqWithoutIsStrInference([d_157_v38_])
            d_265_v119_: _dafny.Seq
            d_265_v119_ = _dafny.SeqWithoutIsStrInference([(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], False, not(d_117_v5_), (d_259_v112_).f6])
            def iife20_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in ((d_264_v118_)[default__.safeIndex(len(d_265_v119_), len(d_264_v118_))]).Elements:
                    d_266_v117_: int = compr_8_
                    if (d_266_v117_) in ((d_264_v118_)[default__.safeIndex(len(d_265_v119_), len(d_264_v118_))]):
                        coll8_[(d_266_v117_) - (len(d_160_v41_))] = (d_259_v112_).f6
                return _dafny.Map(coll8_)
            d_263_v116_ = (d_263_v116_).set(d_260_v113_, iife20_()
            )
            d_115_v3_ = default__.safeModuloInt(d_115_v3_, default__.safeDivisionInt(len(d_157_v38_), d_115_v3_))
        elif source5_.is_DC8:
            d_267___mcc_h11_ = source5_.cf7
            d_268_cf7_ = d_267___mcc_h11_
            index44_ = default__.safeIndex(339, (d_158_v39_).length(0))
            (d_158_v39_)[index44_] = len(_dafny.SeqWithoutIsStrInference([d_115_v3_ for d_269_i8_ in range(default__.abs(-902))]))
            index45_ = default__.safeIndex(339, (d_158_v39_).length(0))
            (d_158_v39_)[index45_] = default__.safeModuloInt(d_115_v3_, d_115_v3_)
            d_270_v120_: _dafny.Seq
            d_270_v120_ = _dafny.SeqWithoutIsStrInference([(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]])
            (d_111_globalState_).f2 = ((d_270_v120_).set(default__.safeIndex(d_115_v3_, len(d_270_v120_)), (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))])) <= (d_270_v120_)
            d_271_v121_: D7
            d_271_v121_ = D7_DC22(len(d_270_v120_))
            d_272_v122_: D0
            d_272_v122_ = D0_DC0((d_271_v121_).cf33)
            d_115_v3_ = (d_272_v122_).cf0
            d_157_v38_ = ((d_157_v38_) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Set({(d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_117_v5_}))]))) + (d_157_v38_)
        elif True:
            d_273___mcc_h12_ = source5_.cf13
            d_274_cf13_ = d_273___mcc_h12_
            d_275_v123_: _dafny.Map
            d_275_v123_ = _dafny.Map({d_117_v5_: (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))]})
            d_276_v124_: _dafny.Map
            d_276_v124_ = _dafny.Map({(d_275_v123_).set(d_117_v5_, d_117_v5_): d_115_v3_})
            d_277_v125_: _dafny.Map
            d_277_v125_ = _dafny.Map({d_276_v124_: d_158_v39_})
            d_277_v125_ = (d_277_v125_).set(d_276_v124_, d_158_v39_)
            d_278_v126_: D8
            d_278_v126_ = D8_DC23(d_114_v2_)
            d_114_v2_ = (d_278_v126_).cf34
            d_279_v127_: C1
            nw45_ = C1()
            nw45_.ctor__((d_115_v3_) - (d_115_v3_), (d_115_v3_) * ((0) - (d_115_v3_)))
            d_279_v127_ = nw45_
            d_117_v5_ = d_117_v5_
        hi2_ = d_115_v3_
        for d_280_i9_ in range((d_157_v38_)[default__.safeIndex(330, len(d_157_v38_))], hi2_):
            d_281_v128_: _dafny.Set
            d_281_v128_ = _dafny.Set({d_115_v3_})
            d_282_v129_: bool
            d_283_v130_: _dafny.Map
            out9_: bool
            out10_: _dafny.Map
            out9_, out10_ = default__.m5(d_157_v38_, d_281_v128_, d_111_globalState_)
            d_282_v129_ = out9_
            d_283_v130_ = out10_
            d_284_v131_: T0
            nw46_ = C2()
            nw46_.ctor__(False, d_115_v3_, d_115_v3_)
            d_284_v131_ = nw46_
            nw47_ = C2()
            nw47_.ctor__(d_282_v129_, d_280_i9_, d_115_v3_)
            d_284_v131_ = nw47_
            d_285_v132_: C3
            nw48_ = C3()
            nw48_.ctor__(default__.fm12(d_117_v5_, d_117_v5_, (d_116_v4_)[default__.safeIndex(831, (d_116_v4_).length(0))], d_111_globalState_))
            d_285_v132_ = nw48_
            index46_ = default__.safeIndex(831, (d_116_v4_).length(0))
            rhs42_ = d_117_v5_
            rhs43_ = (len((_dafny.Set({d_285_v132_, d_285_v132_})).intersection(_dafny.Set({d_285_v132_})))) < (d_284_v131_.f5)
            rhs44_ = d_117_v5_
            rhs45_ = d_285_v132_
            lhs34_ = d_116_v4_
            lhs35_ = default__.safeIndex(831, (d_116_v4_).length(0))
            lhs36_ = d_111_globalState_
            d_117_v5_ = rhs42_
            lhs34_[lhs35_] = rhs43_
            lhs36_.f2 = rhs44_
            d_285_v132_ = rhs45_
            d_115_v3_ = default__.safeModuloInt(default__.safeDivisionInt(-302, (0) - ((0) - (d_284_v131_.f5))), default__.safeModuloInt((0) - (d_280_i9_), (d_284_v131_).f4))
        _dafny.print(_dafny.string_of((d_111_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_111_globalState_).f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_111_globalState_.f2))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_112_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_113_v1_) == (_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_114_v2_).f7) == (_dafny.MultiSet([_dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h'), _dafny.CodePoint('h')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_115_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_116_v4_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_117_v5_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_118_v6_)[3]).cf14) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_119_v7_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_120_v8_).cf14) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_121_i0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_157_v38_) == (_dafny.SeqWithoutIsStrInference([0, 2, 0]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v39_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v39_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v39_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v39_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_158_v39_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_159_v40_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_160_v41_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(int(0), int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2)
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
        return lambda: D1_DC4(False)
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

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf5', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf5)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf5 == __o.cf5
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
    def is_DC6(self) -> bool:
        return isinstance(self, D2_DC6)

class D2_DC7(D2, NamedTuple('DC7', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC9(int(0), False, int(0), False)
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

class D3_DC9(D3, NamedTuple('DC9', [('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC10(D3, NamedTuple('DC10', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC8(D3, NamedTuple('DC8', [('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC8({self.cf7.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC8) and self.cf7 == __o.cf7
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf13', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf13)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf13 == __o.cf13
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D4_DC13)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D4_DC14)
    @property
    def is_DC12(self) -> bool:
        return isinstance(self, D4_DC12)

class D4_DC13(D4, NamedTuple('DC13', [('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf16', Any), ('cf17', Any), ('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)}, {_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17 and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf14)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC16(int(0), _dafny.MultiSet({}), _dafny.Seq({}), D3.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D5_DC15)

class D5_DC16(D5, NamedTuple('DC16', [('cf20', Any), ('cf21', Any), ('cf22', Any), ('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf20 == __o.cf20 and self.cf21 == __o.cf21 and self.cf22 == __o.cf22 and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC15(D5, NamedTuple('DC15', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC15({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC15) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC18(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D6_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D6_DC17)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC18(D6, NamedTuple('DC18', [('cf25', Any), ('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC18({self.cf25.VerbatimString(True)}, {_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC18) and self.cf25 == __o.cf25 and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC17(D6, NamedTuple('DC17', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC17({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC17) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC21(False, _dafny.CodePoint('D'), _dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D7_DC22)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D7_DC20)

class D7_DC21(D7, NamedTuple('DC21', [('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC22(D7, NamedTuple('DC22', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC22({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC22) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC20(D7, NamedTuple('DC20', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC20({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC20) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC24(int(0), _dafny.Array(None, 0), _dafny.Array(None, 0), int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D8_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D8_DC25)

class D8_DC24(D8, NamedTuple('DC24', [('cf35', Any), ('cf36', Any), ('cf37', Any), ('cf38', Any), ('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC24({_dafny.string_of(self.cf35)}, {_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)}, {_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC24) and self.cf35 == __o.cf35 and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38 and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC23(D8, NamedTuple('DC23', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC25(D8, NamedTuple('DC25', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC25({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC25) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC26(D9, NamedTuple('DC26', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    def m1(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f2: bool = False
        self._f0: bool = False
        self._f1: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2):
        (self)._f0 = f0
        (self)._f1 = f1
        (self).f2 = f2

    @property
    def f0(self):
        return self._f0
    @property
    def f1(self):
        return self._f1

class C0:
    def  __init__(self):
        self._f7: _dafny.MultiSet = _dafny.MultiSet({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f7):
        (self)._f7 = f7

    @property
    def f7(self):
        return self._f7

class C1(T0):
    def  __init__(self):
        self._f5: int = int(0)
        self._f4: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f4, f5):
        (self)._f4 = f4
        (self).f5 = f5

    def fm1(self, p0, p1, p2, p3, globalState):
        return True

    def fm2(self, p0, p1, p2, globalState):
        return ((_dafny.MultiSet([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))}))])) | (_dafny.MultiSet([len(_dafny.Map({True: (self).f4}))]))).ispropersubset((_dafny.MultiSet([(self).f4])) - (_dafny.MultiSet([len(_dafny.Map({False: False}))])))

    def m1(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        d_286_v0_: _dafny.Seq
        d_286_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ffyidwc"))
        d_287_v1_: bool
        d_287_v1_ = True
        d_288_v2_: _dafny.Seq
        d_288_v2_ = _dafny.SeqWithoutIsStrInference([d_287_v1_])
        (self).f5 = default__.safeModuloInt(len((d_286_v0_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ykiwthwa")))), (0) - (len(_dafny.Set({(d_288_v2_)[default__.safeIndex(self.f5, len(d_288_v2_))]}))))
        d_289_v3_: _dafny.Set
        d_289_v3_ = _dafny.Set({not(d_287_v1_)})
        r2 = ((d_289_v3_) - (d_289_v3_)).isdisjoint(d_289_v3_)
        (self).f5 = self.f5
        r1 = d_287_v1_
        d_290_v4_: _dafny.Map
        d_290_v4_ = _dafny.Map({(self).f4: not(True)})
        if not((len((d_290_v4_) | (d_290_v4_))) >= ((len(d_289_v3_)) - (self.f5))):
            rhs46_ = (len(d_286_v0_)) != ((self).f4)
            rhs47_ = (self.f5) >= (self.f5)
            r0 = rhs46_
            r1 = rhs47_
            d_291_v5_: _dafny.Array
            nw49_ = _dafny.Array(False, 8)
            d_291_v5_ = nw49_
            d_292_v6_: _dafny.Seq
            d_292_v6_ = _dafny.SeqWithoutIsStrInference([d_291_v5_, d_291_v5_, d_291_v5_, d_291_v5_])
            index47_ = default__.safeIndex(656, (p0).length(0))
            (p0)[index47_] = default__.safeDivisionInt(len(d_292_v6_), self.f5)
            d_293_v7_: _dafny.Seq
            d_293_v7_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_294_v8_: _dafny.Map
            d_294_v8_ = _dafny.Map({((d_293_v7_).set(default__.safeIndex((self).f4, len(d_293_v7_)), self.f5)) + (_dafny.SeqWithoutIsStrInference([len(d_286_v0_)])): self.f5})
            d_295_v9_: D0
            d_295_v9_ = D0_DC0(len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khdwjltpl")))})))
            index48_ = default__.safeIndex(656, (p0).length(0))
            rhs48_ = (self).f4
            rhs49_ = default__.safeDivisionInt((0) - ((d_295_v9_).cf0), (self).f4)
            rhs50_ = d_294_v8_
            lhs37_ = p0
            lhs38_ = default__.safeIndex(656, (p0).length(0))
            lhs39_ = self
            lhs37_[lhs38_] = rhs48_
            lhs39_.f5 = rhs49_
            d_294_v8_ = rhs50_
            (self).f5 = (self).f4
            d_296_v10_: D1
            d_296_v10_ = D1_DC3(d_287_v1_)
            d_297_v11_: _dafny.Map
            d_297_v11_ = _dafny.Map({d_287_v1_: (self).f4})
            d_298_v12_: _dafny.MultiSet
            d_298_v12_ = _dafny.MultiSet([len(d_297_v11_)])
            d_299_v13_: int
            d_300_v14_: _dafny.Seq
            d_301_v15_: _dafny.MultiSet
            d_302_v16_: _dafny.Seq
            out11_: int
            out12_: _dafny.Seq
            out13_: _dafny.MultiSet
            out14_: _dafny.Seq
            out11_, out12_, out13_, out14_ = (self).m4((d_296_v10_).cf3, ((d_298_v12_)[369] if (369) in (d_298_v12_) else default__.fm3(d_287_v1_, globalState)), _dafny.SeqWithoutIsStrInference([p1 for d_303_i0_ in range(default__.abs(682))]), globalState)
            d_299_v13_ = out11_
            d_300_v14_ = out12_
            d_301_v15_ = out13_
            d_302_v16_ = out14_
            d_304_v17_: _dafny.Map
            d_304_v17_ = _dafny.Map({default__.safeModuloInt(d_299_v13_, d_299_v13_): (self.f5) + (933)})
            d_305_v18_: _dafny.MultiSet
            d_305_v18_ = _dafny.MultiSet([False])
            d_304_v17_ = (d_304_v17_).set(default__.safeModuloInt((self).f4, self.f5), ((_dafny.MultiSet([True, d_287_v1_, d_287_v1_, True, d_287_v1_])).intersection((d_305_v18_).set(d_287_v1_, default__.abs(self.f5)))).cardinality)
        elif True:
            d_306_v19_: _dafny.MultiSet
            d_306_v19_ = _dafny.MultiSet([_dafny.CodePoint('i'), p1])
            d_307_v20_: C0
            nw50_ = C0()
            nw50_.ctor__(d_306_v19_)
            d_307_v20_ = nw50_
            d_308_v21_: _dafny.Array
            nw51_ = _dafny.Array(_dafny.CodePoint('D'), 10)
            d_308_v21_ = nw51_
            index49_ = default__.safeIndex(917, (d_308_v21_).length(0))
            (d_308_v21_)[index49_] = p1
            d_309_v22_: _dafny.Seq
            d_309_v22_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            index50_ = default__.safeIndex(917, (d_308_v21_).length(0))
            (d_308_v21_)[index50_] = default__.fm4((self).fm2(not(d_287_v1_), (self).f4, d_287_v1_, globalState), self.f5, d_287_v1_, default__.safeDivisionInt(self.f5, (d_309_v22_)[default__.safeIndex((self).f4, len(d_309_v22_))]), globalState)
            r2 = not(True)
            d_310_v23_: _dafny.Seq
            d_310_v23_ = _dafny.SeqWithoutIsStrInference([d_309_v22_])
            r0 = (d_310_v23_) != (_dafny.SeqWithoutIsStrInference([d_309_v22_ for d_311_i1_ in range(default__.abs(338))]))
            index51_ = default__.safeIndex(339, (p0).length(0))
            (p0)[index51_] = self.f5
            index52_ = default__.safeIndex(339, (p0).length(0))
            (p0)[index52_] = (0) - ((0) - (default__.fm3(d_287_v1_, globalState)))
        d_312_v24_: D1
        d_312_v24_ = D1_DC3(d_287_v1_)
        d_313_v25_: _dafny.Map
        d_313_v25_ = _dafny.Map({True: d_287_v1_})
        d_314_v26_: _dafny.Seq
        d_314_v26_ = _dafny.SeqWithoutIsStrInference([160])
        d_315_v27_: _dafny.MultiSet
        d_315_v27_ = _dafny.MultiSet([self.f5, self.f5, len(d_314_v26_), default__.fm3(d_287_v1_, globalState)])
        d_316_v28_: _dafny.Array
        nw52_ = _dafny.Array(None, 25)
        nw52_[int(0)] = (d_287_v1_ if d_287_v1_ else (self).fm2(d_287_v1_, self.f5, d_287_v1_, globalState))
        nw52_[int(1)] = not (d_287_v1_) or (d_287_v1_)
        nw52_[int(2)] = d_287_v1_
        nw52_[int(3)] = d_287_v1_
        nw52_[int(4)] = d_287_v1_
        nw52_[int(5)] = not(not((d_312_v24_).cf3))
        nw52_[int(6)] = (self).fm1(d_287_v1_, self.f5, _dafny.Set({False, d_287_v1_}), ((d_313_v25_)[d_287_v1_] if (d_287_v1_) in (d_313_v25_) else d_287_v1_), globalState)
        nw52_[int(7)] = d_287_v1_
        nw52_[int(8)] = d_287_v1_
        nw52_[int(9)] = d_287_v1_
        nw52_[int(10)] = d_287_v1_
        nw52_[int(11)] = d_287_v1_
        nw52_[int(12)] = (d_315_v27_).issubset(d_315_v27_)
        nw52_[int(13)] = d_287_v1_
        nw52_[int(14)] = not((default__.fm3(d_287_v1_, globalState)) != (252))
        nw52_[int(15)] = d_287_v1_
        nw52_[int(16)] = d_287_v1_
        nw52_[int(17)] = (True) or (d_287_v1_)
        nw52_[int(18)] = not (d_287_v1_) or (d_287_v1_)
        nw52_[int(19)] = d_287_v1_
        nw52_[int(20)] = d_287_v1_
        nw52_[int(21)] = d_287_v1_
        nw52_[int(22)] = d_287_v1_
        nw52_[int(23)] = (d_312_v24_).cf3
        nw52_[int(24)] = (d_288_v2_)[default__.safeIndex(609, len(d_288_v2_))]
        d_316_v28_ = nw52_
        index53_ = default__.safeIndex(948, (d_316_v28_).length(0))
        (d_316_v28_)[index53_] = not(d_287_v1_)
        index54_ = default__.safeIndex(948, (d_316_v28_).length(0))
        (d_316_v28_)[index54_] = d_287_v1_
        r0 = (d_316_v28_)[default__.safeIndex(948, (d_316_v28_).length(0))]
        d_317_v29_: _dafny.Map
        d_317_v29_ = _dafny.Map({self.f5: self.f5})
        r1 = (len(d_317_v29_)) <= (-967)
        r2 = (d_287_v1_ if (d_316_v28_)[default__.safeIndex(948, (d_316_v28_).length(0))] else d_287_v1_)
        return r0, r1, r2

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: int = int(0)
        d_318_v0_: bool
        d_318_v0_ = True
        (globalState).f2 = d_318_v0_
        if d_318_v0_:
            d_319_v1_: _dafny.MultiSet
            d_319_v1_ = _dafny.MultiSet([_dafny.CodePoint('x')])
            d_320_v2_: C0
            nw53_ = C0()
            nw53_.ctor__(d_319_v1_)
            d_320_v2_ = nw53_
            d_321_v3_: C0
            nw54_ = C0()
            nw54_.ctor__((d_320_v2_).f7)
            d_321_v3_ = nw54_
            d_322_v4_: _dafny.Array
            nw55_ = _dafny.Array(False, 25)
            d_322_v4_ = nw55_
            index55_ = default__.safeIndex(205, (d_322_v4_).length(0))
            (d_322_v4_)[index55_] = not(d_318_v0_)
            index56_ = default__.safeIndex(205, (d_322_v4_).length(0))
            (d_322_v4_)[index56_] = d_318_v0_
            d_321_v3_ = d_321_v3_
            d_323_v5_: D0
            d_323_v5_ = D0_DC0(self.f5)
            source6_ = d_323_v5_
            if source6_.is_DC1:
                d_324___mcc_h0_ = source6_.cf1
                d_325___mcc_h1_ = source6_.cf2
                d_326_cf2_ = d_325___mcc_h1_
                d_327_cf1_ = d_324___mcc_h0_
                d_328_v6_: _dafny.Seq
                d_328_v6_ = _dafny.SeqWithoutIsStrInference([d_326_cf2_])
                d_329_v7_: _dafny.Seq
                d_329_v7_ = _dafny.SeqWithoutIsStrInference([(len(d_328_v6_)) >= (self.f5)])
                d_330_v8_: _dafny.Seq
                d_330_v8_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omrnrrxhq"))
                d_329_v7_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hyq"))) == (d_330_v8_)])
                d_331_v9_: _dafny.Set
                d_331_v9_ = _dafny.Set({d_318_v0_, (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]})
                d_332_v10_: _dafny.Map
                d_332_v10_ = _dafny.Map({True: (d_331_v9_).ispropersubset(d_331_v9_)})
                index57_ = default__.safeIndex(205, (d_322_v4_).length(0))
                (d_322_v4_)[index57_] = ((d_332_v10_)[not((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))])] if (not((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))])) in (d_332_v10_) else (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))])
                d_333_v11_: _dafny.Map
                d_333_v11_ = _dafny.Map({(self).fm2((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))], -72, (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))], globalState): _dafny.SeqWithoutIsStrInference([134, 287])})
                d_328_v6_ = ((d_333_v11_)[d_318_v0_] if (d_318_v0_) in (d_333_v11_) else d_328_v6_)
                (globalState).f2 = d_318_v0_
            elif source6_.is_DC2:
                d_334_v12_: _dafny.Array
                nw56_ = _dafny.Array(D0.default()(), 21)
                d_334_v12_ = nw56_
                index58_ = default__.safeIndex(198, (d_334_v12_).length(0))
                (d_334_v12_)[index58_] = D0_DC2()
                d_335_v13_: _dafny.Map
                d_335_v13_ = _dafny.Map({not ((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]) or (d_318_v0_): d_318_v0_})
                d_336_v14_: _dafny.Set
                d_336_v14_ = _dafny.Set({d_318_v0_})
                d_337_v16_: _dafny.Seq
                d_337_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxo"))
                d_338_v17_: _dafny.Map
                d_338_v17_ = _dafny.Map({(self).f4: d_337_v16_})
                d_339_v18_: _dafny.Array
                nw57_ = _dafny.Array(None, 22)
                nw57_[int(0)] = p0
                nw57_[int(1)] = p0
                nw57_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_318_v0_, (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]])) + ((p0).set(default__.safeIndex((self).f4, len(p0)), d_318_v0_))
                nw57_[int(3)] = p0
                nw57_[int(4)] = _dafny.SeqWithoutIsStrInference([d_318_v0_])
                nw57_[int(5)] = _dafny.SeqWithoutIsStrInference([(d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))], False])
                nw57_[int(6)] = (p0).set(default__.safeIndex(self.f5, len(p0)), (self).fm1(d_318_v0_, (0) - ((self).f4), d_336_v14_, d_318_v0_, globalState))
                nw57_[int(7)] = (p0) + (p0)
                nw57_[int(8)] = p0
                nw57_[int(9)] = _dafny.SeqWithoutIsStrInference([(d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))], (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]])
                nw57_[int(10)] = _dafny.SeqWithoutIsStrInference([(p0)[default__.safeIndex(self.f5, len(p0))]])
                nw57_[int(11)] = _dafny.SeqWithoutIsStrInference([d_318_v0_])
                nw57_[int(12)] = p0
                def iife21_():
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(-360, 250):
                        d_340_v15_: int = compr_9_
                        if ((-360) <= (d_340_v15_)) and ((d_340_v15_) < (250)):
                            coll9_[(d_340_v15_) + (self.f5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dfrfshin"))
                    return _dafny.Map(coll9_)
                nw57_[int(13)] = default__.fm5(iife21_()
                , self.f5, d_318_v0_, (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))], globalState)
                nw57_[int(14)] = p0
                nw57_[int(15)] = (default__.fm5(d_338_v17_, self.f5, (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))], d_318_v0_, globalState)) + (p0)
                nw57_[int(16)] = p0
                nw57_[int(17)] = p0
                nw57_[int(18)] = p0
                nw57_[int(19)] = p0
                nw57_[int(20)] = p0
                nw57_[int(21)] = p0
                d_339_v18_ = nw57_
                index59_ = default__.safeIndex(450, (d_339_v18_).length(0))
                (d_339_v18_)[index59_] = (p0) + (p0)
                d_341_v19_: _dafny.Map
                d_341_v19_ = _dafny.Map({(0) - ((self).f4): self.f5})
                d_342_v20_: D0
                d_342_v20_ = D0_DC2()
                d_343_v21_: _dafny.Map
                d_343_v21_ = _dafny.Map({(self).f4: (d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]})
                index60_ = default__.safeIndex(198, (d_334_v12_).length(0))
                index61_ = default__.safeIndex(450, (d_339_v18_).length(0))
                rhs51_ = (self.f5) <= (len(d_341_v19_))
                rhs52_ = d_342_v20_
                rhs53_ = d_335_v13_
                rhs54_ = p0
                rhs55_ = (0) - ((default__.safeModuloInt((d_323_v5_).cf0, self.f5)) + (default__.safeModuloInt(len(d_343_v21_), (self).f4)))
                lhs40_ = globalState
                lhs41_ = d_334_v12_
                lhs42_ = default__.safeIndex(198, (d_334_v12_).length(0))
                lhs43_ = d_339_v18_
                lhs44_ = default__.safeIndex(450, (d_339_v18_).length(0))
                lhs40_.f2 = rhs51_
                lhs41_[lhs42_] = rhs52_
                d_335_v13_ = rhs53_
                lhs43_[lhs44_] = rhs54_
                r1 = rhs55_
                def iife22_():
                    coll10_ = _dafny.Set()
                    compr_10_: int
                    for compr_10_ in _dafny.IntegerRange(814, 257):
                        d_344_v22_: int = compr_10_
                        if ((814) <= (d_344_v22_)) and ((d_344_v22_) < (257)):
                            coll10_ = coll10_.union(_dafny.Set([(d_344_v22_) + (self.f5)]))
                    return _dafny.Set(coll10_)
                d_341_v19_ = (d_341_v19_).set(default__.safeModuloInt(self.f5, len(iife22_()
                )), default__.safeDivisionInt(self.f5, (self).f4))
                d_345_v23_: D1
                d_345_v23_ = D1_DC3(not (d_318_v0_) or ((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]))
                d_345_v23_ = d_345_v23_
                (self).f5 = (default__.fm3(d_318_v0_, globalState)) - ((self).f4)
            elif True:
                d_346___mcc_h2_ = source6_.cf0
                d_347_cf0_ = d_346___mcc_h2_
                d_323_v5_ = d_323_v5_
                d_348_v24_: _dafny.Set
                d_348_v24_ = _dafny.Set({(d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]})
                (self).f5 = default__.fm3(((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]) not in (d_348_v24_), globalState)
                d_349_v25_: _dafny.Map
                d_349_v25_ = _dafny.Map({(d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]: (self).f4})
                d_349_v25_ = (d_349_v25_).set(not((d_322_v4_)[default__.safeIndex(205, (d_322_v4_).length(0))]), (self).f4)
                d_350_v26_: _dafny.MultiSet
                d_350_v26_ = _dafny.MultiSet([self.f5])
                d_351_v27_: _dafny.Seq
                d_351_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
                r0 = ((0) - ((self.f5 if d_318_v0_ else (d_350_v26_).cardinality))) + (len(d_351_v27_))
        elif True:
            d_352_v28_: str
            d_352_v28_ = _dafny.CodePoint('e')
            d_353_v29_: _dafny.Array
            nw58_ = _dafny.Array(None, 3)
            nw58_[int(0)] = d_352_v28_
            nw58_[int(1)] = d_352_v28_
            nw58_[int(2)] = d_352_v28_
            d_353_v29_ = nw58_
            rhs56_ = d_318_v0_
            rhs57_ = ((self).f4) == ((self).f4)
            rhs58_ = default__.fm3((p0) < (p0), globalState)
            rhs59_ = d_353_v29_
            rhs60_ = (self).f4
            lhs45_ = globalState
            d_318_v0_ = rhs56_
            lhs45_.f2 = rhs57_
            r0 = rhs58_
            d_353_v29_ = rhs59_
            r1 = rhs60_
            d_354_v30_: _dafny.MultiSet
            d_354_v30_ = _dafny.MultiSet([d_352_v28_])
            d_355_v31_: C0
            nw59_ = C0()
            nw59_.ctor__(d_354_v30_)
            d_355_v31_ = nw59_
            d_356_v32_: _dafny.Map
            d_356_v32_ = _dafny.Map({False: self.f5})
            d_356_v32_ = (d_356_v32_).set(d_318_v0_, self.f5)
            d_357_v33_: _dafny.Array
            nw60_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_357_v33_ = nw60_
            d_357_v33_ = d_357_v33_
            (globalState).f2 = d_318_v0_
        d_358_v34_: str
        d_358_v34_ = _dafny.CodePoint('b')
        d_359_v35_: _dafny.Seq
        d_359_v35_ = _dafny.SeqWithoutIsStrInference([d_358_v34_])
        d_360_v36_: C0
        nw61_ = C0()
        nw61_.ctor__(_dafny.MultiSet(d_359_v35_))
        d_360_v36_ = nw61_
        d_361_v37_: _dafny.Array
        def lambda9_(d_362_v34_):
            def lambda10_(d_363_i0_):
                return d_362_v34_

            return lambda10_

        init4_ = lambda9_(d_358_v34_)
        nw62_ = _dafny.Array(None, 21)
        for i0_4_ in range(nw62_.length(0)):
            nw62_[i0_4_] = init4_(i0_4_)
        d_361_v37_ = nw62_
        d_364_v38_: _dafny.Map
        d_364_v38_ = _dafny.Map({d_358_v34_: d_361_v37_})
        d_365_v39_: _dafny.Seq
        d_365_v39_ = _dafny.SeqWithoutIsStrInference([((d_364_v38_)[d_358_v34_] if (d_358_v34_) in (d_364_v38_) else d_361_v37_), d_361_v37_, d_361_v37_])
        if ((d_365_v39_ if d_318_v0_ else d_365_v39_)) == (_dafny.SeqWithoutIsStrInference([d_361_v37_, d_361_v37_])):
            d_366_v40_: _dafny.Array
            nw63_ = _dafny.Array(int(0), 19)
            d_366_v40_ = nw63_
            d_367_v41_: _dafny.Map
            d_367_v41_ = _dafny.Map({d_318_v0_: d_366_v40_})
            d_367_v41_ = (d_367_v41_) | (_dafny.Map({d_318_v0_: d_366_v40_}))
            d_368_v42_: _dafny.Seq
            d_368_v42_ = _dafny.SeqWithoutIsStrInference([len(p0), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgn")))])
            (globalState).f2 = (default__.safeDivisionInt(self.f5, len(d_368_v42_))) < ((self.f5) - (self.f5))
            d_369_v43_: C0
            nw64_ = C0()
            nw64_.ctor__(_dafny.MultiSet(d_359_v35_))
            d_369_v43_ = nw64_
            d_370_v44_: D0
            d_370_v44_ = D0_DC1((self.f5) + (self.f5), self.f5)
            source7_ = d_370_v44_
            if source7_.is_DC1:
                d_371___mcc_h3_ = source7_.cf1
                d_372___mcc_h4_ = source7_.cf2
                d_373_cf2_ = d_372___mcc_h4_
                d_374_cf1_ = d_371___mcc_h3_
                index62_ = default__.safeIndex(379, (d_366_v40_).length(0))
                (d_366_v40_)[index62_] = self.f5
                index63_ = default__.safeIndex(379, (d_366_v40_).length(0))
                (d_366_v40_)[index63_] = d_373_cf2_
                d_375_v45_: int
                d_376_v46_: _dafny.Seq
                d_377_v47_: _dafny.MultiSet
                d_378_v48_: _dafny.Seq
                out15_: int
                out16_: _dafny.Seq
                out17_: _dafny.MultiSet
                out18_: _dafny.Seq
                out15_, out16_, out17_, out18_ = (self).m4(d_318_v0_, (self).f4, default__.fm6(d_318_v0_, self.f5, d_318_v0_, globalState), globalState)
                d_375_v45_ = out15_
                d_376_v46_ = out16_
                d_377_v47_ = out17_
                d_378_v48_ = out18_
                d_367_v41_ = d_367_v41_
                d_379_v49_: _dafny.Array
                def lambda11_(d_380_cf1_):
                    def lambda12_(d_381_i1_):
                        return (d_381_i1_) * (d_380_cf1_)

                    return lambda12_

                init5_ = lambda11_(d_374_cf1_)
                nw65_ = _dafny.Array(None, 27)
                for i0_5_ in range(nw65_.length(0)):
                    nw65_[i0_5_] = init5_(i0_5_)
                d_379_v49_ = nw65_
                d_379_v49_ = d_379_v49_
            elif source7_.is_DC2:
                d_382_v50_: _dafny.Map
                d_382_v50_ = _dafny.Map({(d_318_v0_) == (d_318_v0_): d_318_v0_})
                d_382_v50_ = (d_382_v50_).set(not(d_318_v0_), d_318_v0_)
                d_370_v44_ = d_370_v44_
                d_318_v0_ = d_318_v0_
                d_382_v50_ = (d_382_v50_).set(d_318_v0_, False)
            elif True:
                d_383___mcc_h5_ = source7_.cf0
                d_384_cf0_ = d_383___mcc_h5_
                (globalState).f2 = d_318_v0_
                d_385_v51_: _dafny.Map
                d_385_v51_ = _dafny.Map({False: (self).f4})
                d_386_v52_: D1
                d_386_v52_ = D1_DC4(d_318_v0_)
                pat_let_tv10_ = d_318_v0_
                d_387_v53_: _dafny.Set
                def iife23_(_pat_let6_0):
                    def iife24_(d_388_dt__update__tmp_h0_):
                        def iife25_(_pat_let7_0):
                            def iife26_(d_389_dt__update_hcf4_h0_):
                                return D1_DC4(d_389_dt__update_hcf4_h0_)
                            return iife26_(_pat_let7_0)
                        return iife25_(pat_let_tv10_)
                    return iife24_(_pat_let6_0)
                d_387_v53_ = _dafny.Set({default__.fm7(d_318_v0_, d_385_v51_, d_358_v34_, d_318_v0_, globalState), d_386_v52_, default__.fm7(d_318_v0_, d_385_v51_, d_358_v34_, d_318_v0_, globalState), d_386_v52_, iife23_(d_386_v52_)})
                def iife27_():
                    coll11_ = _dafny.Set()
                    compr_11_: D1
                    for compr_11_ in (default__.fm8(d_358_v34_, d_318_v0_, d_384_cf0_, d_318_v0_, globalState)).keys.Elements:
                        d_390_v54_: D1 = compr_11_
                        if (d_390_v54_) in (default__.fm8(d_358_v34_, d_318_v0_, d_384_cf0_, d_318_v0_, globalState)):
                            coll11_ = coll11_.union(_dafny.Set([d_390_v54_]))
                    return _dafny.Set(coll11_)
                d_387_v53_ = iife27_()
                
                d_391_v55_: _dafny.Array
                def lambda13_(d_392_v0_):
                    def lambda14_(d_393_i2_):
                        return d_392_v0_

                    return lambda14_

                init6_ = lambda13_(d_318_v0_)
                nw66_ = _dafny.Array(None, 10)
                for i0_6_ in range(nw66_.length(0)):
                    nw66_[i0_6_] = init6_(i0_6_)
                d_391_v55_ = nw66_
                d_394_v56_: _dafny.Map
                d_394_v56_ = _dafny.Map({d_384_cf0_: d_358_v34_})
                rhs61_ = ((d_394_v56_)[(self).f4] if ((self).f4) in (d_394_v56_) else d_358_v34_)
                rhs62_ = d_358_v34_
                rhs63_ = not (not (d_318_v0_) or (d_318_v0_)) or (not(d_318_v0_))
                rhs64_ = d_391_v55_
                d_358_v34_ = rhs61_
                d_358_v34_ = rhs62_
                d_318_v0_ = rhs63_
                d_391_v55_ = rhs64_
                (globalState).f2 = (_dafny.CodePoint('u')) in (d_359_v35_)
            r0 = (d_368_v42_)[default__.safeIndex((0) - (((_dafny.MultiSet([d_318_v0_, d_318_v0_])).cardinality) * ((self).f4)), len(d_368_v42_))]
        elif True:
            d_395_v57_: _dafny.Array
            def lambda15_(d_396_v0_):
                def lambda16_(d_397_i3_):
                    return not (d_396_v0_) or (d_396_v0_)

                return lambda16_

            init7_ = lambda15_(d_318_v0_)
            nw67_ = _dafny.Array(None, 25)
            for i0_7_ in range(nw67_.length(0)):
                nw67_[i0_7_] = init7_(i0_7_)
            d_395_v57_ = nw67_
            d_398_v58_: _dafny.Array
            nw68_ = _dafny.Array(_dafny.Map({}), 4)
            d_398_v58_ = nw68_
            d_399_v59_: _dafny.Set
            d_399_v59_ = _dafny.Set({d_318_v0_})
            rhs65_ = d_395_v57_
            rhs66_ = ((d_399_v59_) - (d_399_v59_)) != (d_399_v59_)
            rhs67_ = d_398_v58_
            rhs68_ = (self).fm2(d_318_v0_, self.f5, d_318_v0_, globalState)
            lhs46_ = globalState
            d_395_v57_ = rhs65_
            lhs46_.f2 = rhs66_
            d_398_v58_ = rhs67_
            d_318_v0_ = rhs68_
            if True:
                d_400_v60_: _dafny.Seq
                d_400_v60_ = _dafny.SeqWithoutIsStrInference([(d_360_v36_).f7])
                d_401_v61_: C0
                nw69_ = C0()
                nw69_.ctor__((d_400_v60_)[default__.safeIndex(self.f5, len(d_400_v60_))])
                d_401_v61_ = nw69_
                d_402_v62_: _dafny.Array
                nw70_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_402_v62_ = nw70_
                nw71_ = _dafny.Array(_dafny.Array(None, 0), 22)
                d_402_v62_ = nw71_
                d_403_v63_: C0
                nw72_ = C0()
                nw72_.ctor__((d_360_v36_).f7)
                d_403_v63_ = nw72_
                r0 = ((self).f4) + ((self).f4)
                d_404_v64_: D0
                d_404_v64_ = D0_DC0(self.f5)
                r0 = (d_404_v64_).cf0
            elif True:
                d_405_v65_: _dafny.Array
                nw73_ = _dafny.Array(int(0), 5)
                d_405_v65_ = nw73_
                index64_ = default__.safeIndex(826, (d_405_v65_).length(0))
                (d_405_v65_)[index64_] = self.f5
                index65_ = default__.safeIndex(826, (d_405_v65_).length(0))
                (d_405_v65_)[index65_] = (0) - ((self).f4)
                (globalState).f2 = not(not(d_318_v0_))
                d_395_v57_ = d_395_v57_
                d_405_v65_ = d_405_v65_
                (globalState).f2 = (d_359_v35_) == (d_359_v35_)
            d_406_v66_: _dafny.MultiSet
            d_406_v66_ = _dafny.MultiSet([d_318_v0_])
            d_407_v67_: _dafny.Seq
            d_407_v67_ = _dafny.SeqWithoutIsStrInference([d_359_v35_])
            d_408_v68_: _dafny.Map
            d_408_v68_ = _dafny.Map({len(default__.fm6(d_318_v0_, (self).f4, d_318_v0_, globalState)): False})
            d_409_v69_: _dafny.Set
            d_409_v69_ = _dafny.Set({d_408_v68_, d_408_v68_, d_408_v68_})
            d_410_v70_: _dafny.Map
            d_410_v70_ = _dafny.Map({(d_406_v66_).isdisjoint(d_406_v66_): not((default__.fm6(not(d_318_v0_), self.f5, d_318_v0_, globalState)) < ((d_407_v67_)[default__.safeIndex(len(d_409_v69_), len(d_407_v67_))]))})
            d_410_v70_ = (d_410_v70_).set(d_318_v0_, d_318_v0_)
            d_411_v71_: _dafny.MultiSet
            d_411_v71_ = _dafny.MultiSet([p0])
            rhs69_ = (d_318_v0_) and (d_318_v0_)
            rhs70_ = ((0) - (((d_411_v71_)[p0] if (p0) in (d_411_v71_) else self.f5))) - (215)
            lhs47_ = globalState
            lhs47_.f2 = rhs69_
            r0 = rhs70_
            (self).f5 = default__.safeModuloInt((self).f4, len((d_359_v35_ if d_318_v0_ else (d_359_v35_).set(default__.safeIndex((self).f4, len(d_359_v35_)), d_358_v34_))))
        d_412_v72_: _dafny.MultiSet
        d_412_v72_ = _dafny.MultiSet([d_318_v0_])
        d_412_v72_ = d_412_v72_
        d_413_v73_: _dafny.Map
        d_413_v73_ = _dafny.Map({(d_359_v35_ if d_318_v0_ else d_359_v35_): d_318_v0_})
        d_413_v73_ = (d_413_v73_).set((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_414_i4_ in range(default__.abs(584))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pu"))), d_318_v0_)
        d_415_v74_: _dafny.Map
        d_415_v74_ = _dafny.Map({d_318_v0_: d_358_v34_})
        r0 = default__.safeDivisionInt((len(d_415_v74_)) - ((self).f4), (self).f4)
        r1 = (self).f4
        return r0, r1

    def m4(self, p0, p1, p2, globalState):
        r0: int = int(0)
        r1: _dafny.Seq = _dafny.Seq({})
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: _dafny.Seq = _dafny.Seq({})
        d_416_v0_: _dafny.Seq
        d_416_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_417_v1_: _dafny.MultiSet
        d_417_v1_ = _dafny.MultiSet([len(_dafny.Set({d_416_v0_, _dafny.SeqWithoutIsStrInference([p1 for d_418_i0_ in range(default__.abs(658))])}))])
        d_419_v2_: _dafny.Seq
        d_419_v2_ = _dafny.SeqWithoutIsStrInference([(d_417_v1_).issubset(d_417_v1_)])
        (globalState).f2 = (d_419_v2_)[default__.safeIndex(self.f5, len(d_419_v2_))]
        d_420_v3_: D1
        d_420_v3_ = D1_DC3(p0)
        d_421_v4_: _dafny.Set
        d_421_v4_ = _dafny.Set({D1_DC5(d_420_v3_)})
        d_422_v5_: D1
        d_422_v5_ = D1_DC5(d_420_v3_)
        d_421_v4_ = _dafny.Set({d_422_v5_, d_422_v5_})
        d_423_v6_: _dafny.MultiSet
        d_423_v6_ = _dafny.MultiSet([p0])
        (globalState).f2 = (self).fm1(p0, self.f5, default__.fm9(p1, p0, (d_423_v6_).cardinality, globalState), (p1) >= (self.f5), globalState)
        if p0:
            d_424_v7_: _dafny.Map
            d_424_v7_ = _dafny.Map({p0: d_416_v0_})
            d_425_v8_: _dafny.Map
            d_425_v8_ = _dafny.Map({p0: p1})
            d_426_v9_: _dafny.Set
            d_426_v9_ = _dafny.Set({p0})
            d_427_v10_: D1
            d_427_v10_ = D1_DC4(p0)
            d_428_v11_: _dafny.Seq
            d_428_v11_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_427_v10_).cf4, p0}), d_426_v9_, d_426_v9_])
            d_429_v12_: _dafny.MultiSet
            d_429_v12_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([D1_DC4(p0) for d_430_i1_ in range(default__.abs(431))])])
            d_431_v13_: _dafny.Seq
            d_431_v13_ = _dafny.SeqWithoutIsStrInference([d_427_v10_, d_427_v10_])
            d_432_v14_: _dafny.Array
            nw74_ = _dafny.Array(None, 21)
            nw74_[int(0)] = p1
            nw74_[int(1)] = (0) - ((p1) + ((0) - ((d_416_v0_)[default__.safeIndex(312, len(d_416_v0_))])))
            nw74_[int(2)] = (len(d_424_v7_)) + (self.f5)
            nw74_[int(3)] = len(d_425_v8_)
            nw74_[int(4)] = (0) - (len((d_426_v9_).intersection((d_428_v11_)[default__.safeIndex(p1, len(d_428_v11_))])))
            nw74_[int(5)] = (0) - (p1)
            nw74_[int(6)] = (self).f4
            nw74_[int(7)] = p1
            nw74_[int(8)] = ((d_429_v12_)[d_431_v13_] if (d_431_v13_) in (d_429_v12_) else (self).f4)
            nw74_[int(9)] = self.f5
            nw74_[int(10)] = p1
            nw74_[int(11)] = self.f5
            nw74_[int(12)] = self.f5
            nw74_[int(13)] = -836
            nw74_[int(14)] = default__.fm3((self).fm1(not(p0), default__.fm3(p0, globalState), d_426_v9_, p0, globalState), globalState)
            nw74_[int(15)] = 254
            nw74_[int(16)] = (self).f4
            nw74_[int(17)] = (0) - ((self.f5) * (len(_dafny.SeqWithoutIsStrInference([p0]))))
            nw74_[int(18)] = p1
            nw74_[int(19)] = (d_416_v0_)[default__.safeIndex(self.f5, len(d_416_v0_))]
            nw74_[int(20)] = len((d_416_v0_) + (d_416_v0_))
            d_432_v14_ = nw74_
            index66_ = default__.safeIndex(852, (d_432_v14_).length(0))
            (d_432_v14_)[index66_] = (p1) - (self.f5)
            d_433_v15_: _dafny.Array
            def lambda17_(d_434_p0_):
                def lambda18_(d_435_i2_):
                    return (_dafny.CodePoint('v') if d_434_p0_ else _dafny.CodePoint('f'))

                return lambda18_

            init8_ = lambda17_(p0)
            nw75_ = _dafny.Array(None, 15)
            for i0_8_ in range(nw75_.length(0)):
                nw75_[i0_8_] = init8_(i0_8_)
            d_433_v15_ = nw75_
            d_436_v16_: str
            d_436_v16_ = _dafny.CodePoint('s')
            index67_ = default__.safeIndex(206, (d_433_v15_).length(0))
            (d_433_v15_)[index67_] = d_436_v16_
            d_437_v17_: _dafny.Map
            d_437_v17_ = _dafny.Map({(self).f4: p0})
            index68_ = default__.safeIndex(852, (d_432_v14_).length(0))
            index69_ = default__.safeIndex(206, (d_433_v15_).length(0))
            rhs71_ = (0) - ((self).f4)
            rhs72_ = default__.fm4(p0, -652, not (p0) or (False), default__.fm3(not(False), globalState), globalState)
            rhs73_ = not (p0) or (p0)
            rhs74_ = not(p0)
            rhs75_ = (d_437_v17_ if p0 else (d_437_v17_) | (default__.fm10(self.f5, p0, (self).f4, p1, globalState)))
            lhs48_ = d_432_v14_
            lhs49_ = default__.safeIndex(852, (d_432_v14_).length(0))
            lhs50_ = d_433_v15_
            lhs51_ = default__.safeIndex(206, (d_433_v15_).length(0))
            lhs52_ = globalState
            lhs53_ = globalState
            lhs48_[lhs49_] = rhs71_
            lhs50_[lhs51_] = rhs72_
            lhs52_.f2 = rhs73_
            lhs53_.f2 = rhs74_
            d_437_v17_ = rhs75_
            d_438_v18_: D0
            d_438_v18_ = D0_DC1(self.f5, (d_432_v14_)[default__.safeIndex(852, (d_432_v14_).length(0))])
            d_438_v18_ = d_438_v18_
            r0 = default__.safeModuloInt(default__.safeModuloInt(-676, p1), (self).f4)
            index70_ = default__.safeIndex(852, (d_432_v14_).length(0))
            (d_432_v14_)[index70_] = (self).f4
            (globalState).f2 = p0
        elif True:
            (globalState).f2 = (default__.fm11(-981, p0, (0) - ((0) - (self.f5)), globalState)).issubset((d_417_v1_) - (d_417_v1_))
            (self).f5 = (((self).f4) - ((0) - (p1))) - (self.f5)
            (globalState).f2 = (p1) < (194)
            d_439_v19_: _dafny.Map
            d_439_v19_ = _dafny.Map({p0: p0})
            if (self).fm2(p0, self.f5, ((d_439_v19_)[p0] if (p0) in (d_439_v19_) else p0), globalState):
                d_440_v20_: _dafny.Array
                nw76_ = _dafny.Array(None, 10)
                d_440_v20_ = nw76_
                d_440_v20_ = d_440_v20_
                d_441_v21_: _dafny.Array
                nw77_ = _dafny.Array(_dafny.Seq({}), 11)
                d_441_v21_ = nw77_
                rhs76_ = p1
                rhs77_ = d_441_v21_
                r0 = rhs76_
                d_441_v21_ = rhs77_
                r0 = p1
                d_416_v0_ = d_416_v0_
                d_442_v22_: _dafny.Array
                def lambda19_(d_443_p1_):
                    def lambda20_(d_444_i3_):
                        return (d_444_i3_) * (d_443_p1_)

                    return lambda20_

                init9_ = lambda19_(p1)
                nw78_ = _dafny.Array(None, 18)
                for i0_9_ in range(nw78_.length(0)):
                    nw78_[i0_9_] = init9_(i0_9_)
                d_442_v22_ = nw78_
                rhs78_ = (self).f4
                rhs79_ = d_442_v22_
                rhs80_ = p0
                rhs81_ = (0) - ((-201 if (d_417_v1_).isdisjoint(d_417_v1_) else (0) - (((d_423_v6_).intersection(d_423_v6_)).cardinality)))
                lhs54_ = self
                lhs55_ = globalState
                lhs56_ = self
                lhs54_.f5 = rhs78_
                d_442_v22_ = rhs79_
                lhs55_.f2 = rhs80_
                lhs56_.f5 = rhs81_
            elif True:
                r0 = len((d_419_v2_) + (d_419_v2_))
                d_445_v23_: _dafny.Array
                nw79_ = _dafny.Array(_dafny.Seq({}), 8)
                d_445_v23_ = nw79_
                index71_ = default__.safeIndex(645, (d_445_v23_).length(0))
                (d_445_v23_)[index71_] = d_416_v0_
                index72_ = default__.safeIndex(645, (d_445_v23_).length(0))
                (d_445_v23_)[index72_] = (d_416_v0_).set(default__.safeIndex(default__.fm3(p0, globalState), len(d_416_v0_)), p1)
                d_446_v24_: D0
                d_446_v24_ = D0_DC0(self.f5)
                (self).f5 = (d_446_v24_).cf0
                d_447_v25_: _dafny.Array
                def lambda21_(d_448_p0_):
                    def lambda22_(d_449_i4_):
                        return d_448_p0_

                    return lambda22_

                init10_ = lambda21_(p0)
                nw80_ = _dafny.Array(None, 3)
                for i0_10_ in range(nw80_.length(0)):
                    nw80_[i0_10_] = init10_(i0_10_)
                d_447_v25_ = nw80_
                d_450_v26_: _dafny.Map
                d_450_v26_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_447_v25_, d_447_v25_]): p1})
                d_451_v27_: _dafny.Map
                d_451_v27_ = _dafny.Map({p0: d_447_v25_})
                d_452_v28_: _dafny.Seq
                d_452_v28_ = _dafny.SeqWithoutIsStrInference([((d_451_v27_)[False] if (False) in (d_451_v27_) else d_447_v25_), d_447_v25_, d_447_v25_, d_447_v25_, d_447_v25_])
                d_453_v29_: _dafny.Seq
                d_453_v29_ = _dafny.SeqWithoutIsStrInference([d_452_v28_])
                (globalState).f2 = (p1) < (((d_450_v26_)[(d_453_v29_)[default__.safeIndex((self).f4, len(d_453_v29_))]] if ((d_453_v29_)[default__.safeIndex((self).f4, len(d_453_v29_))]) in (d_450_v26_) else len(p2)))
                d_454_v30_: _dafny.Set
                d_454_v30_ = _dafny.Set({p0})
                index73_ = default__.safeIndex(805, (d_447_v25_).length(0))
                (d_447_v25_)[index73_] = (self).fm1(not((self).fm2(False, p1, p0, globalState)), (0) - (p1), d_454_v30_, p0, globalState)
                index74_ = default__.safeIndex(805, (d_447_v25_).length(0))
                (d_447_v25_)[index74_] = p0
            r0 = default__.safeDivisionInt((self).f4, (self).f4)
        (globalState).f2 = p0
        d_455_v31_: _dafny.MultiSet
        d_455_v31_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k"))])
        d_455_v31_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxhuqgf")), p2])
        r0 = (self).f4
        r1 = _dafny.SeqWithoutIsStrInference([(self).f4 for d_456_i5_ in range(default__.abs(715))])
        r2 = d_417_v1_
        r3 = d_419_v2_
        return r0, r1, r2, r3


class C2(T0):
    def  __init__(self):
        self._f5: int = int(0)
        self._f4: int = int(0)
        self._f6: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f5(self):
        return self._f5
    @f5.setter
    def f5(self, value):
        self._f5 = value
    @property
    def f4(self):
        return self._f4
    def ctor__(self, f6, f4, f5):
        (self)._f6 = f6
        (self)._f4 = f4
        (self).f5 = f5

    def fm0(self, p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.Map({(self).f6: (self).f6})) | (_dafny.Map({(self).f6: (self).f6}))])

    def m1(self, p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        d_457_v0_: _dafny.Array
        nw81_ = _dafny.Array(False, 15)
        d_457_v0_ = nw81_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_457_v0_).length(0)):
            d_458_i0_: int = guard_loop_0_
            if (True) and (((0) <= (d_458_i0_)) and ((d_458_i0_) < ((d_457_v0_).length(0)))):
                (d_457_v0_)[(d_458_i0_)] = (self).f6
        d_459_v1_: str
        d_459_v1_ = _dafny.CodePoint('o')
        d_459_v1_ = d_459_v1_
        (self).f5 = self.f5
        rhs82_ = self.f5
        rhs83_ = (self).f6
        lhs57_ = self
        lhs58_ = globalState
        lhs57_.f5 = rhs82_
        lhs58_.f2 = rhs83_
        d_460_v2_: C0
        nw82_ = C0()
        nw82_.ctor__(_dafny.MultiSet([d_459_v1_]))
        d_460_v2_ = nw82_
        r0 = (self).f6
        d_461_v3_: D1
        d_461_v3_ = D1_DC4((self).f6)
        r0 = not((d_461_v3_).cf4)
        d_462_v4_: _dafny.Seq
        d_462_v4_ = _dafny.SeqWithoutIsStrInference([default__.fm12((self).f6, (self).f6, (self).f6, globalState)])
        r1 = (d_462_v4_)[default__.safeIndex((self).f4, len(d_462_v4_))]
        r2 = (self).f6
        return r0, r1, r2

    def m2(self, p0, globalState):
        r0: bool = False
        r1: int = int(0)
        if False:
            r0 = (self).f6
            if ((self).f6) and ((self).f6):
                r0 = not((self).f6)
                d_463_v0_: _dafny.Seq
                d_463_v0_ = _dafny.SeqWithoutIsStrInference([(self).f6])
                d_464_v1_: _dafny.MultiSet
                d_464_v1_ = _dafny.MultiSet([d_463_v0_])
                d_465_v2_: D2
                d_465_v2_ = D2_DC6(d_463_v0_)
                d_466_v3_: _dafny.Map
                d_466_v3_ = _dafny.Map({p0: (self).f6})
                d_467_v4_: _dafny.Seq
                d_467_v4_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "quepdvot"))
                d_468_v5_: D3
                d_468_v5_ = D3_DC8(d_467_v4_)
                d_469_v6_: _dafny.Set
                d_469_v6_ = _dafny.Set({(self).f6})
                d_470_v7_: str
                d_470_v7_ = _dafny.CodePoint('v')
                d_471_v8_: _dafny.Array
                nw83_ = _dafny.Array(None, 18)
                nw83_[int(0)] = self.f5
                nw83_[int(1)] = ((d_464_v1_)[_dafny.SeqWithoutIsStrInference([not((self).f6)])] if (_dafny.SeqWithoutIsStrInference([not((self).f6)])) in (d_464_v1_) else default__.fm3((self).f6, globalState))
                nw83_[int(2)] = self.f5
                nw83_[int(3)] = len((d_465_v2_).cf6)
                nw83_[int(4)] = (self).f4
                nw83_[int(5)] = len(d_466_v3_)
                nw83_[int(6)] = (self).f4
                nw83_[int(7)] = p0
                nw83_[int(8)] = self.f5
                nw83_[int(9)] = self.f5
                nw83_[int(10)] = -265
                nw83_[int(11)] = len(((d_468_v5_).cf7).set(default__.safeIndex(len(d_469_v6_), len((d_468_v5_).cf7)), d_470_v7_))
                nw83_[int(12)] = (self).f4
                nw83_[int(13)] = p0
                nw83_[int(14)] = (self).f4
                nw83_[int(15)] = p0
                nw83_[int(16)] = self.f5
                nw83_[int(17)] = (self).f4
                d_471_v8_ = nw83_
                d_472_v9_: _dafny.Array
                nw84_ = _dafny.Array(False, 14)
                d_472_v9_ = nw84_
                d_473_v10_: _dafny.Map
                d_473_v10_ = _dafny.Map({d_471_v8_: d_472_v9_})
                d_474_v11_: D1
                d_474_v11_ = D1_DC3((self).f6)
                d_475_v12_: _dafny.Map
                d_475_v12_ = _dafny.Map({d_473_v10_: not(((self).f6) == (default__.fm12((d_474_v11_).cf3, (self).f6, (self).f6, globalState)))})
                d_476_v13_: _dafny.Map
                d_476_v13_ = _dafny.Map({(self).f6: (self).f6})
                d_475_v12_ = (d_475_v12_).set((d_473_v10_).set(d_471_v8_, d_472_v9_), ((d_476_v13_)[(self).f6] if ((self).f6) in (d_476_v13_) else (self).f6))
                d_466_v3_ = (d_466_v3_).set(self.f5, (self).f6)
                (self).f5 = p0
                d_477_v14_: _dafny.MultiSet
                d_477_v14_ = _dafny.MultiSet([d_470_v7_, d_470_v7_])
                d_478_v15_: C0
                nw85_ = C0()
                nw85_.ctor__((d_477_v14_) - (_dafny.MultiSet([d_470_v7_])))
                d_478_v15_ = nw85_
            elif True:
                d_479_v16_: str
                d_479_v16_ = _dafny.CodePoint('v')
                d_480_v17_: _dafny.MultiSet
                d_480_v17_ = _dafny.MultiSet([d_479_v16_])
                d_481_v18_: C0
                nw86_ = C0()
                nw86_.ctor__(d_480_v17_)
                d_481_v18_ = nw86_
                (globalState).f2 = (self).f6
                d_479_v16_ = d_479_v16_
                d_482_v19_: _dafny.Seq
                d_482_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ngmxo"))
                d_483_v20_: D3
                d_483_v20_ = D3_DC8(d_482_v19_)
                d_482_v19_ = (d_483_v20_).cf7
                d_482_v19_ = default__.fm6((self).f6, self.f5, not((self).f6), globalState)
            (self).f5 = (0) - (default__.fm3((self).f6, globalState))
            d_484_v21_: _dafny.MultiSet
            d_484_v21_ = _dafny.MultiSet([(self).f6, (self).f6, (self).f6, (self).f6, (self).f6])
            d_485_v22_: _dafny.Seq
            d_485_v22_ = _dafny.SeqWithoutIsStrInference([(self).f4])
            d_486_v23_: str
            d_486_v23_ = _dafny.CodePoint('t')
            d_487_v24_: _dafny.MultiSet
            d_487_v24_ = _dafny.MultiSet([d_486_v23_, d_486_v23_, d_486_v23_])
            d_488_v25_: C0
            nw87_ = C0()
            nw87_.ctor__(d_487_v24_)
            d_488_v25_ = nw87_
            d_489_v26_: D0
            d_489_v26_ = D0_DC1((self).f4, 469)
            d_490_v27_: D0
            d_490_v27_ = D0_DC0(p0)
            d_491_v28_: D1
            d_491_v28_ = D1_DC4((self).f6)
            d_492_v29_: D3
            d_492_v29_ = D3_DC9(p0, (self).f6, self.f5, (d_491_v28_).cf4)
            d_493_v30_: _dafny.Array
            nw88_ = _dafny.Array(None, 19)
            nw88_[int(0)] = 356
            nw88_[int(1)] = self.f5
            nw88_[int(2)] = (d_484_v21_).cardinality
            nw88_[int(3)] = default__.fm3((self).f6, globalState)
            nw88_[int(4)] = self.f5
            nw88_[int(5)] = (self).f4
            nw88_[int(6)] = (0) - ((d_485_v22_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_488_v25_])), len(d_485_v22_))])
            nw88_[int(7)] = self.f5
            nw88_[int(8)] = ((d_489_v26_).cf1) + (self.f5)
            nw88_[int(9)] = 195
            nw88_[int(10)] = (self).f4
            nw88_[int(11)] = p0
            nw88_[int(12)] = (d_489_v26_).cf2
            nw88_[int(13)] = (d_490_v27_).cf0
            nw88_[int(14)] = (d_485_v22_)[default__.safeIndex((0) - (p0), len(d_485_v22_))]
            nw88_[int(15)] = self.f5
            nw88_[int(16)] = (d_492_v29_).cf10
            nw88_[int(17)] = self.f5
            nw88_[int(18)] = (self).f4
            d_493_v30_ = nw88_
            index75_ = default__.safeIndex(619, (d_493_v30_).length(0))
            (d_493_v30_)[index75_] = (self).f4
            index76_ = default__.safeIndex(619, (d_493_v30_).length(0))
            (d_493_v30_)[index76_] = (self).f4
            d_494_v31_: _dafny.Seq
            d_494_v31_ = _dafny.SeqWithoutIsStrInference([(self).f6, (self).f6])
            d_495_v32_: _dafny.Map
            d_495_v32_ = _dafny.Map({d_493_v30_: d_494_v31_})
            d_484_v21_ = _dafny.MultiSet(((d_495_v32_)[d_493_v30_] if (d_493_v30_) in (d_495_v32_) else d_494_v31_))
        elif True:
            (globalState).f2 = (self).f6
            d_496_v33_: str
            d_496_v33_ = _dafny.CodePoint('e')
            d_497_v34_: _dafny.MultiSet
            d_497_v34_ = _dafny.MultiSet([d_496_v33_])
            d_498_v35_: C0
            nw89_ = C0()
            nw89_.ctor__(d_497_v34_)
            d_498_v35_ = nw89_
            d_499_v36_: D1
            d_499_v36_ = D1_DC4((self).f6)
            d_500_v37_: _dafny.Array
            nw90_ = _dafny.Array(None, 10)
            nw90_[int(0)] = d_499_v36_
            nw90_[int(1)] = d_499_v36_
            nw90_[int(2)] = d_499_v36_
            nw90_[int(3)] = d_499_v36_
            nw90_[int(4)] = d_499_v36_
            nw90_[int(5)] = d_499_v36_
            nw90_[int(6)] = d_499_v36_
            def iife28_(_pat_let8_0):
                def iife29_(d_501_dt__update__tmp_h0_):
                    def iife30_(_pat_let9_0):
                        def iife31_(d_502_dt__update_hcf4_h0_):
                            return D1_DC4(d_502_dt__update_hcf4_h0_)
                        return iife31_(_pat_let9_0)
                    return iife30_((self).f6)
                return iife29_(_pat_let8_0)
            nw90_[int(7)] = iife28_(d_499_v36_)
            nw90_[int(8)] = d_499_v36_
            nw90_[int(9)] = D1_DC4((self).f6)
            d_500_v37_ = nw90_
            d_503_v38_: _dafny.Array
            nw91_ = _dafny.Array(int(0), 24)
            d_503_v38_ = nw91_
            d_504_v39_: _dafny.Seq
            d_504_v39_ = _dafny.SeqWithoutIsStrInference([d_503_v38_])
            rhs84_ = d_498_v35_
            rhs85_ = d_500_v37_
            rhs86_ = (d_504_v39_)[default__.safeIndex(((self).f4) - (default__.fm3((self).f6, globalState)), len(d_504_v39_))]
            rhs87_ = ((self).f6) and ((self).f6)
            rhs88_ = ((0) - ((p0) * ((self).f4))) - ((self).f4)
            lhs59_ = globalState
            lhs60_ = self
            d_498_v35_ = rhs84_
            d_500_v37_ = rhs85_
            d_503_v38_ = rhs86_
            lhs59_.f2 = rhs87_
            lhs60_.f5 = rhs88_
            r0 = not(not((self).f6))
            d_505_v40_: _dafny.Set
            d_505_v40_ = _dafny.Set({(self).f6})
            (globalState).f2 = (d_505_v40_) != (d_505_v40_)
            d_506_v41_: T0
            nw92_ = C1()
            nw92_.ctor__(p0, (self).f4)
            d_506_v41_ = nw92_
            d_507_v42_: _dafny.Array
            nw93_ = _dafny.Array(None, 7)
            nw93_[int(0)] = (d_506_v41_ if (self).f6 else d_506_v41_)
            nw93_[int(1)] = d_506_v41_
            nw93_[int(2)] = d_506_v41_
            nw93_[int(3)] = d_506_v41_
            nw93_[int(4)] = d_506_v41_
            nw93_[int(5)] = d_506_v41_
            nw93_[int(6)] = d_506_v41_
            d_507_v42_ = nw93_
            index77_ = default__.safeIndex(386, (d_507_v42_).length(0))
            (d_507_v42_)[index77_] = d_506_v41_
            index78_ = default__.safeIndex(386, (d_507_v42_).length(0))
            (d_507_v42_)[index78_] = d_506_v41_
        d_508_v43_: _dafny.Array
        nw94_ = _dafny.Array(_dafny.Map({}), 2)
        d_508_v43_ = nw94_
        d_509_v44_: _dafny.Array
        def lambda23_(d_510_p0_):
            def lambda24_(d_511_i0_):
                return default__.safeModuloInt(d_511_i0_, d_510_p0_)

            return lambda24_

        init11_ = lambda23_(p0)
        nw95_ = _dafny.Array(None, 15)
        for i0_11_ in range(nw95_.length(0)):
            nw95_[i0_11_] = init11_(i0_11_)
        d_509_v44_ = nw95_
        d_512_v45_: str
        d_512_v45_ = _dafny.CodePoint('q')
        d_513_v46_: _dafny.Map
        d_513_v46_ = _dafny.Map({d_509_v44_: d_512_v45_})
        index79_ = default__.safeIndex(663, (d_508_v43_).length(0))
        (d_508_v43_)[index79_] = d_513_v46_
        d_514_v47_: _dafny.Map
        d_514_v47_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f4 for d_515_i1_ in range(default__.abs(-702))]): d_513_v46_})
        d_516_v48_: _dafny.Seq
        d_516_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spwqw"))
        d_517_v49_: _dafny.Map
        d_517_v49_ = _dafny.Map({(self).f6: d_516_v48_})
        d_518_v50_: _dafny.Seq
        d_518_v50_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsprbn")))), p0, (0) - (len(((d_517_v49_)[(self).f6] if ((self).f6) in (d_517_v49_) else d_516_v48_))), 703, (self).f4])
        index80_ = default__.safeIndex(663, (d_508_v43_).length(0))
        (d_508_v43_)[index80_] = ((d_514_v47_)[d_518_v50_] if (d_518_v50_) in (d_514_v47_) else d_513_v46_)
        if (self).f6:
            d_519_v51_: _dafny.Seq
            d_519_v51_ = _dafny.SeqWithoutIsStrInference([(self).f6])
            d_520_v52_: _dafny.Array
            nw96_ = _dafny.Array(None, 10)
            nw96_[int(0)] = default__.fm12((self).f6, default__.fm12((self).f6, (self).f6, (self).f6, globalState), (self).f6, globalState)
            nw96_[int(1)] = (self).f6
            def iife32_(_pat_let10_0):
                def iife33_(d_521_dt__update__tmp_h1_):
                    def iife34_(_pat_let11_0):
                        def iife35_(d_522_dt__update_hcf3_h0_):
                            return D1_DC3(d_522_dt__update_hcf3_h0_)
                        return iife35_(_pat_let11_0)
                    return iife34_((self).f6)
                return iife33_(_pat_let10_0)
            nw96_[int(2)] = (iife32_(D1_DC3((self).f6))).cf3
            nw96_[int(3)] = True
            nw96_[int(4)] = (self).f6
            nw96_[int(5)] = (d_519_v51_)[default__.safeIndex(p0, len(d_519_v51_))]
            nw96_[int(6)] = (self).f6
            nw96_[int(7)] = (self).f6
            nw96_[int(8)] = (self).f6
            nw96_[int(9)] = not ((self).f6) or (not((self).f6))
            d_520_v52_ = nw96_
            d_523_v53_: D1
            d_523_v53_ = D1_DC4((self).f6)
            index81_ = default__.safeIndex(29, (d_520_v52_).length(0))
            (d_520_v52_)[index81_] = (d_523_v53_).cf4
            index82_ = default__.safeIndex(29, (d_520_v52_).length(0))
            rhs89_ = self.f5
            rhs90_ = d_520_v52_
            rhs91_ = (self).f6
            lhs61_ = d_520_v52_
            lhs62_ = default__.safeIndex(29, (d_520_v52_).length(0))
            r1 = rhs89_
            d_520_v52_ = rhs90_
            lhs61_[lhs62_] = rhs91_
            if (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]:
                d_524_v54_: _dafny.MultiSet
                d_524_v54_ = _dafny.MultiSet([d_512_v45_])
                d_525_v55_: _dafny.Seq
                d_525_v55_ = _dafny.SeqWithoutIsStrInference([d_524_v54_])
                d_526_v56_: C0
                nw97_ = C0()
                nw97_.ctor__((d_525_v55_)[default__.safeIndex(960, len(d_525_v55_))])
                d_526_v56_ = nw97_
                d_527_v57_: _dafny.Seq
                d_527_v57_ = _dafny.SeqWithoutIsStrInference([d_526_v56_, d_526_v56_])
                d_528_v58_: _dafny.Map
                d_528_v58_ = _dafny.Map({(d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]: (d_527_v57_)[default__.safeIndex((self).f4, len(d_527_v57_))]})
                d_528_v58_ = (d_528_v58_).set(True, d_526_v56_)
                index83_ = default__.safeIndex(29, (d_520_v52_).length(0))
                (d_520_v52_)[index83_] = (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]
                r1 = (default__.fm3(not((self).f6), globalState)) + ((self).f4)
                d_529_v59_: _dafny.MultiSet
                d_529_v59_ = _dafny.MultiSet([d_526_v56_])
                d_530_v60_: _dafny.Map
                d_530_v60_ = _dafny.Map({(p0 if (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))] else len(_dafny.SeqWithoutIsStrInference([d_512_v45_ for d_531_i2_ in range(default__.abs(271))]))): (self).f4})
                index84_ = default__.safeIndex(29, (d_520_v52_).length(0))
                rhs92_ = (0) - ((0) - ((self).f4))
                rhs93_ = _dafny.MultiSet([d_526_v56_, d_526_v56_, d_526_v56_, d_526_v56_, d_526_v56_])
                rhs94_ = not((d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))])
                rhs95_ = (d_530_v60_) | (d_530_v60_)
                rhs96_ = default__.fm12((self.f5) >= (p0), (self).f6, (self).f6, globalState)
                lhs63_ = d_520_v52_
                lhs64_ = default__.safeIndex(29, (d_520_v52_).length(0))
                lhs65_ = globalState
                r1 = rhs92_
                d_529_v59_ = rhs93_
                lhs63_[lhs64_] = rhs94_
                d_530_v60_ = rhs95_
                lhs65_.f2 = rhs96_
                d_532_v61_: _dafny.MultiSet
                d_532_v61_ = _dafny.MultiSet([-358])
                index85_ = default__.safeIndex(29, (d_520_v52_).length(0))
                (d_520_v52_)[index85_] = ((d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]) or ((d_532_v61_).ispropersubset(d_532_v61_))
            elif True:
                (globalState).f2 = ((d_516_v48_).set(default__.safeIndex(p0, len(d_516_v48_)), _dafny.CodePoint('n'))) == (default__.fm6((self).f6, p0, (self).f6, globalState))
                d_533_v62_: _dafny.Map
                d_533_v62_ = _dafny.Map({(self).f6: (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]})
                d_533_v62_ = (d_533_v62_ if (self).f6 else d_533_v62_)
                d_534_v63_: _dafny.Array
                nw98_ = _dafny.Array(_dafny.MultiSet({}), 26)
                d_534_v63_ = nw98_
                d_535_v64_: _dafny.Set
                d_535_v64_ = _dafny.Set({(self).f4, (self).f4, (self).f4})
                index86_ = default__.safeIndex(305, (d_534_v63_).length(0))
                (d_534_v63_)[index86_] = _dafny.MultiSet([d_535_v64_])
                d_536_v65_: _dafny.MultiSet
                d_536_v65_ = _dafny.MultiSet([_dafny.Set({len(d_516_v48_), 548, (D3_DC9((self).f4, (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))], p0, (self).f6)).cf8})])
                index87_ = default__.safeIndex(305, (d_534_v63_).length(0))
                (d_534_v63_)[index87_] = d_536_v65_
                d_535_v64_ = (d_535_v64_).intersection(d_535_v64_)
                r1 = ((self.f5) - (self.f5)) + (default__.fm3((self).f6, globalState))
            if (self).f6:
                (self).f5 = (self).f4
                (self).f5 = (0) - ((self).f4)
                index88_ = default__.safeIndex(29, (d_520_v52_).length(0))
                (d_520_v52_)[index88_] = (self).f6
                d_512_v45_ = d_512_v45_
                (globalState).f2 = (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]
            elif True:
                index89_ = default__.safeIndex(29, (d_520_v52_).length(0))
                (d_520_v52_)[index89_] = (self).f6
                index90_ = default__.safeIndex(29, (d_520_v52_).length(0))
                (d_520_v52_)[index90_] = (self).f6
                (self).f5 = (0) - (self.f5)
                d_537_v66_: _dafny.Map
                d_537_v66_ = _dafny.Map({D2_DC6(_dafny.SeqWithoutIsStrInference([(self).f6, (d_520_v52_)[default__.safeIndex(29, (d_520_v52_).length(0))]])): (self).f6})
                index91_ = default__.safeIndex(26, (d_509_v44_).length(0))
                (d_509_v44_)[index91_] = ((0) - (len(d_537_v66_))) * (self.f5)
                index92_ = default__.safeIndex(26, (d_509_v44_).length(0))
                (d_509_v44_)[index92_] = default__.safeDivisionInt(p0, (self).f4)
                d_538_v67_: T0
                nw99_ = C1()
                nw99_.ctor__((self).f4, (self).f4)
                d_538_v67_ = nw99_
                d_539_v68_: _dafny.Seq
                d_539_v68_ = _dafny.SeqWithoutIsStrInference([d_538_v67_, d_538_v67_])
                d_540_v69_: _dafny.Map
                d_540_v69_ = _dafny.Map({(self).f4: len(d_539_v68_)})
                d_541_v70_: _dafny.Set
                d_541_v70_ = _dafny.Set({len(d_540_v69_)})
                d_542_v72_: _dafny.Seq
                def iife36_():
                    coll12_ = _dafny.Set()
                    compr_12_: int
                    for compr_12_ in _dafny.IntegerRange(390, 552):
                        d_543_v71_: int = compr_12_
                        if ((390) <= (d_543_v71_)) and ((d_543_v71_) < (552)):
                            coll12_ = coll12_.union(_dafny.Set([(d_543_v71_) - (self.f5)]))
                    return _dafny.Set(coll12_)
                d_542_v72_ = _dafny.SeqWithoutIsStrInference([d_541_v70_, iife36_()
                ])
                (globalState).f2 = not(((d_542_v72_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jucpayd"))), len(d_542_v72_))]).isdisjoint(d_541_v70_))
            (self).f5 = (0) - (p0)
            r1 = (self).f4
        elif True:
            r1 = (p0) + (((self).f4) - (209))
            d_544_v73_: _dafny.MultiSet
            d_544_v73_ = _dafny.MultiSet([d_512_v45_])
            r1 = (((d_544_v73_) - (d_544_v73_)).cardinality) - (default__.safeDivisionInt(615, (_dafny.MultiSet([(self).f6, (self).f6])).cardinality))
            d_545_v74_: D3
            d_545_v74_ = D3_DC9((self).f4, (self).f6, (self).f4, (self).f6)
            d_546_v75_: _dafny.Seq
            d_546_v75_ = _dafny.SeqWithoutIsStrInference([True, (d_545_v74_).cf9])
            r0 = ((d_546_v75_) + (d_546_v75_)) <= (d_546_v75_)
            d_547_v76_: _dafny.Array
            nw100_ = _dafny.Array(None, 10)
            nw100_[int(0)] = (d_512_v45_ if (self).f6 else d_512_v45_)
            nw100_[int(1)] = (d_512_v45_ if (self).f6 else d_512_v45_)
            nw100_[int(2)] = ((d_516_v48_)[default__.safeIndex((self).f4, len(d_516_v48_))] if (self).f6 else d_512_v45_)
            nw100_[int(3)] = d_512_v45_
            nw100_[int(4)] = d_512_v45_
            nw100_[int(5)] = d_512_v45_
            nw100_[int(6)] = d_512_v45_
            nw100_[int(7)] = d_512_v45_
            nw100_[int(8)] = d_512_v45_
            nw100_[int(9)] = d_512_v45_
            d_547_v76_ = nw100_
            index93_ = default__.safeIndex(799, (d_547_v76_).length(0))
            (d_547_v76_)[index93_] = d_512_v45_
            index94_ = default__.safeIndex(799, (d_547_v76_).length(0))
            (d_547_v76_)[index94_] = d_512_v45_
            (self).f5 = default__.fm3((self).f6, globalState)
        r0 = (self).f6
        r0 = (self).f6
        index95_ = default__.safeIndex(243, (d_509_v44_).length(0))
        (d_509_v44_)[index95_] = (self.f5) * (p0)
        d_548_v77_: _dafny.Array
        def lambda25_(d_549_v45_):
            def lambda26_(d_550_i3_):
                return _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_549_v45_ for d_551_i4_ in range(default__.abs(20))])), (self).f4, -661, (self).f4})

            return lambda26_

        init12_ = lambda25_(d_512_v45_)
        nw101_ = _dafny.Array(None, 2)
        for i0_12_ in range(nw101_.length(0)):
            nw101_[i0_12_] = init12_(i0_12_)
        d_548_v77_ = nw101_
        index96_ = default__.safeIndex(243, (d_509_v44_).length(0))
        rhs97_ = (d_518_v50_)[default__.safeIndex((self).f4, len(d_518_v50_))]
        rhs98_ = d_548_v77_
        lhs66_ = d_509_v44_
        lhs67_ = default__.safeIndex(243, (d_509_v44_).length(0))
        lhs66_[lhs67_] = rhs97_
        d_548_v77_ = rhs98_
        r0 = (self).f6
        r1 = (self.f5) + ((d_509_v44_)[default__.safeIndex(243, (d_509_v44_).length(0))])
        return r0, r1

    @property
    def f6(self):
        return self._f6

class C3:
    def  __init__(self):
        self._f3: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self, f3):
        (self)._f3 = f3

    def m0(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: _dafny.Map = _dafny.Map({})
        d_552_v0_: C2
        nw102_ = C2()
        nw102_.ctor__(default__.fm12(p2, p2, (self).f3, globalState), p1, p1)
        d_552_v0_ = nw102_
        d_553_v1_: _dafny.Seq
        d_553_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqwgvcpna"))
        if (d_553_v1_) < (d_553_v1_):
            d_554_v2_: _dafny.Seq
            d_554_v2_ = _dafny.SeqWithoutIsStrInference([not((d_552_v0_).f6), default__.fm12((d_552_v0_).f6, p2, (d_552_v0_).f6, globalState), True])
            (globalState).f2 = (d_554_v2_)[default__.safeIndex(p1, len(d_554_v2_))]
            if ((d_552_v0_).f6 if p2 else p2):
                d_555_v3_: _dafny.Map
                d_555_v3_ = _dafny.Map({_dafny.CodePoint('e'): (self).f3})
                d_556_v4_: _dafny.MultiSet
                d_556_v4_ = _dafny.MultiSet([d_555_v3_])
                d_557_v5_: _dafny.Seq
                d_557_v5_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_556_v4_).cardinality)])
                d_558_v6_: _dafny.MultiSet
                d_558_v6_ = _dafny.MultiSet([(d_552_v0_).f6, (d_552_v0_).f6, (d_557_v5_) < (d_557_v5_), (p1) <= (len(d_553_v1_)), not((d_552_v0_).f6)])
                d_559_v7_: _dafny.Seq
                d_559_v7_ = _dafny.SeqWithoutIsStrInference([d_554_v2_])
                d_558_v6_ = (d_558_v6_) | ((d_558_v6_).set(p2, default__.abs(len(d_559_v7_))))
                d_560_v8_: str
                d_560_v8_ = _dafny.CodePoint('r')
                d_560_v8_ = (d_560_v8_ if p2 else _dafny.CodePoint('d'))
                d_561_v9_: _dafny.Map
                d_561_v9_ = _dafny.Map({d_560_v8_: d_553_v1_})
                d_561_v9_ = (d_561_v9_).set(d_560_v8_, d_553_v1_)
                d_562_v10_: _dafny.Set
                d_562_v10_ = _dafny.Set({p2})
                (globalState).f2 = ((d_562_v10_) - (d_562_v10_)).isdisjoint(d_562_v10_)
                d_563_v11_: int
                d_563_v11_ = -107
                d_563_v11_ = p1
            elif True:
                d_564_v12_: _dafny.Array
                nw103_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_564_v12_ = nw103_
                d_565_v13_: str
                d_565_v13_ = _dafny.CodePoint('p')
                index97_ = default__.safeIndex(880, (d_564_v12_).length(0))
                (d_564_v12_)[index97_] = d_565_v13_
                index98_ = default__.safeIndex(880, (d_564_v12_).length(0))
                (d_564_v12_)[index98_] = d_565_v13_
                d_566_v14_: _dafny.Map
                d_566_v14_ = _dafny.Map({p1: len(d_553_v1_)})
                d_567_v15_: _dafny.Map
                d_567_v15_ = _dafny.Map({p1: False})
                d_568_v16_: _dafny.Set
                d_568_v16_ = _dafny.Set({(d_552_v0_).f6})
                d_569_v17_: D4
                d_569_v17_ = D4_DC12(d_568_v16_)
                d_566_v14_ = (_dafny.Map({(default__.fm13((d_552_v0_).f6, len(d_567_v15_), p1, globalState)).cardinality: len((d_569_v17_).cf14)})) | (d_566_v14_)
                d_570_v18_: int
                d_570_v18_ = 822
                d_571_v19_: _dafny.Map
                d_571_v19_ = _dafny.Map({(d_552_v0_).f6: d_570_v18_})
                d_572_v20_: _dafny.Seq
                d_572_v20_ = _dafny.SeqWithoutIsStrInference([d_570_v18_])
                d_573_v21_: T0
                nw104_ = C2()
                nw104_.ctor__((d_552_v0_).f6, len(d_572_v20_), d_570_v18_)
                d_573_v21_ = nw104_
                d_574_v22_: _dafny.Map
                d_574_v22_ = _dafny.Map({d_570_v18_: d_573_v21_})
                d_570_v18_ = ((d_571_v19_)[(d_552_v0_).f6] if ((d_552_v0_).f6) in (d_571_v19_) else (len((default__.fm6(False, len(d_574_v22_), (self).f3, globalState)).set(default__.safeIndex(117, len(default__.fm6(False, len(d_574_v22_), (self).f3, globalState))), (d_564_v12_)[default__.safeIndex(880, (d_564_v12_).length(0))]))) * (p1))
                d_575_v23_: _dafny.Array
                nw105_ = _dafny.Array(None, 26)
                nw105_[int(0)] = p2
                nw105_[int(1)] = (d_552_v0_).f6
                nw105_[int(2)] = (d_552_v0_).f6
                nw105_[int(3)] = (d_552_v0_).f6
                nw105_[int(4)] = (d_552_v0_).f6
                nw105_[int(5)] = (d_552_v0_).f6
                nw105_[int(6)] = (d_552_v0_).f6
                nw105_[int(7)] = (d_552_v0_).f6
                nw105_[int(8)] = (d_552_v0_).f6
                nw105_[int(9)] = p2
                nw105_[int(10)] = p2
                nw105_[int(11)] = (d_552_v0_).f6
                nw105_[int(12)] = p2
                nw105_[int(13)] = (d_552_v0_).f6
                nw105_[int(14)] = p2
                nw105_[int(15)] = (d_552_v0_).f6
                nw105_[int(16)] = not(p2)
                nw105_[int(17)] = (d_552_v0_).f6
                nw105_[int(18)] = p2
                nw105_[int(19)] = (self).f3
                nw105_[int(20)] = (self).f3
                nw105_[int(21)] = (d_552_v0_).f6
                nw105_[int(22)] = (d_552_v0_).f6
                nw105_[int(23)] = (d_552_v0_).f6
                nw105_[int(24)] = (d_552_v0_).f6
                nw105_[int(25)] = (d_552_v0_).f6
                d_575_v23_ = nw105_
                d_576_v24_: _dafny.Seq
                d_576_v24_ = _dafny.SeqWithoutIsStrInference([d_575_v23_, d_575_v23_])
                d_577_v25_: _dafny.Seq
                d_577_v25_ = _dafny.SeqWithoutIsStrInference([d_576_v24_])
                d_578_v26_: _dafny.Seq
                d_578_v26_ = _dafny.SeqWithoutIsStrInference([d_576_v24_, ((d_577_v25_)[default__.safeIndex(p1, len(d_577_v25_))]).set(default__.safeIndex((d_573_v21_).f4, len((d_577_v25_)[default__.safeIndex(p1, len(d_577_v25_))])), d_575_v23_)])
                d_579_v27_: _dafny.Map
                d_579_v27_ = _dafny.Map({d_565_v13_: len((d_578_v26_)[default__.safeIndex(d_573_v21_.f5, len(d_578_v26_))])})
                d_580_v28_: _dafny.Array
                nw106_ = _dafny.Array(None, 3)
                nw106_[int(0)] = p2
                nw106_[int(1)] = (d_579_v27_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({d_565_v13_: len(d_572_v20_)}) for d_581_i0_ in range(default__.abs(614))]))
                nw106_[int(2)] = not(((d_573_v21_).f4) != (p1))
                d_580_v28_ = nw106_
                index99_ = default__.safeIndex(880, (d_564_v12_).length(0))
                rhs99_ = ((d_553_v1_)[default__.safeIndex(p1, len(d_553_v1_))] if ((d_552_v0_).f6 if (d_552_v0_).f6 else p2) else d_565_v13_)
                rhs100_ = (D5_DC15(d_580_v28_)).cf19
                rhs101_ = (0) - (default__.fm3(p2, globalState))
                lhs68_ = d_564_v12_
                lhs69_ = default__.safeIndex(880, (d_564_v12_).length(0))
                lhs68_[lhs69_] = rhs99_
                d_575_v23_ = rhs100_
                d_570_v18_ = rhs101_
                d_574_v22_ = (d_574_v22_).set((p1) + ((0) - (p1)), d_573_v21_)
            d_582_v29_: int
            d_582_v29_ = 446
            d_583_v30_: _dafny.MultiSet
            d_583_v30_ = _dafny.MultiSet([not((self).f3)])
            d_582_v29_ = (d_583_v30_).cardinality
            d_584_v31_: _dafny.Seq
            d_584_v31_ = _dafny.SeqWithoutIsStrInference([(p1 if (d_552_v0_).f6 else p1)])
            d_582_v29_ = (0) - ((d_584_v31_)[default__.safeIndex(p1, len(d_584_v31_))])
            (globalState).f2 = not((p1) < ((p1 if (d_552_v0_).f6 else (0) - (len(d_553_v1_)))))
        elif True:
            d_585_v32_: T0
            nw107_ = C1()
            nw107_.ctor__(p1, p1)
            d_585_v32_ = nw107_
            d_586_v33_: _dafny.Map
            d_586_v33_ = _dafny.Map({(self).f3: d_585_v32_})
            d_587_v34_: _dafny.Array
            nw108_ = _dafny.Array(False, 7)
            d_587_v34_ = nw108_
            d_588_v35_: _dafny.Seq
            d_588_v35_ = _dafny.SeqWithoutIsStrInference([d_587_v34_])
            d_589_v36_: D4
            d_589_v36_ = D4_DC14((d_586_v33_) | (d_586_v33_), p2, d_588_v35_)
            pat_let_tv11_ = d_586_v33_
            def iife37_(_pat_let12_0):
                def iife38_(d_590_dt__update__tmp_h0_):
                    def iife39_(_pat_let13_0):
                        def iife40_(d_591_dt__update_hcf16_h0_):
                            def iife41_(_pat_let14_0):
                                def iife42_(d_592_dt__update_hcf17_h0_):
                                    return D4_DC14(d_591_dt__update_hcf16_h0_, d_592_dt__update_hcf17_h0_, (d_590_dt__update__tmp_h0_).cf18)
                                return iife42_(_pat_let14_0)
                            return iife41_((self).f3)
                        return iife40_(_pat_let13_0)
                    return iife39_(pat_let_tv11_)
                return iife38_(_pat_let12_0)
            d_589_v36_ = iife37_(d_589_v36_)
            (d_585_v32_).f5 = d_585_v32_.f5
            d_593_v37_: _dafny.Array
            def lambda27_(d_594_v0_, d_595_p2_):
                def lambda28_(d_596_i1_):
                    return _dafny.MultiSet([(d_594_v0_).f6, d_595_p2_])

                return lambda28_

            init13_ = lambda27_(d_552_v0_, p2)
            nw109_ = _dafny.Array(None, 6)
            for i0_13_ in range(nw109_.length(0)):
                nw109_[i0_13_] = init13_(i0_13_)
            d_593_v37_ = nw109_
            d_597_v38_: _dafny.MultiSet
            d_597_v38_ = _dafny.MultiSet([(d_552_v0_).f6])
            index100_ = default__.safeIndex(147, (d_593_v37_).length(0))
            (d_593_v37_)[index100_] = (d_597_v38_).intersection(d_597_v38_)
            index101_ = default__.safeIndex(147, (d_593_v37_).length(0))
            (d_593_v37_)[index101_] = _dafny.MultiSet([not((d_552_v0_).f6), (self).f3])
            d_598_v39_: _dafny.Seq
            d_598_v39_ = _dafny.SeqWithoutIsStrInference([p2])
            d_599_v40_: D2
            d_599_v40_ = D2_DC6(d_598_v39_)
            source8_ = d_599_v40_
            if source8_.is_DC7:
                d_600_v41_: _dafny.Array
                nw110_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_600_v41_ = nw110_
                d_601_v42_: str
                d_601_v42_ = _dafny.CodePoint('e')
                d_602_v43_: _dafny.Map
                d_602_v43_ = _dafny.Map({d_601_v42_: default__.fm6(p2, d_585_v32_.f5, not((self).f3), globalState)})
                d_603_v44_: C1
                nw111_ = C1()
                nw111_.ctor__(len(((d_602_v43_)[d_601_v42_] if (d_601_v42_) in (d_602_v43_) else d_553_v1_)), len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([502, (d_585_v32_).f4, len(d_553_v1_), p1])) for d_604_i2_ in range(default__.abs(648))])))
                d_603_v44_ = nw111_
                d_605_v45_: _dafny.Map
                d_605_v45_ = _dafny.Map({d_603_v44_: d_585_v32_.f5})
                rhs102_ = d_587_v34_
                rhs103_ = d_600_v41_
                rhs104_ = (((d_605_v45_)[d_603_v44_] if (d_603_v44_) in (d_605_v45_) else 372)) != (p1)
                lhs70_ = globalState
                d_587_v34_ = rhs102_
                d_600_v41_ = rhs103_
                lhs70_.f2 = rhs104_
                d_606_v46_: _dafny.Array
                nw112_ = _dafny.Array(D3.default()(), 17)
                d_606_v46_ = nw112_
                d_607_v47_: _dafny.Map
                d_607_v47_ = _dafny.Map({d_606_v46_: p1})
                d_608_v48_: _dafny.Map
                d_608_v48_ = _dafny.Map({(d_552_v0_).f6: p2})
                d_609_v49_: _dafny.Map
                d_609_v49_ = _dafny.Map({(d_607_v47_) | (d_607_v47_): (_dafny.Map({(d_552_v0_).f6: p2})) | (d_608_v48_)})
                d_609_v49_ = (d_609_v49_).set(d_607_v47_, d_608_v48_)
                (globalState).f2 = (d_552_v0_).f6
                d_610_v50_: _dafny.Array
                nw113_ = _dafny.Array(_dafny.Map({}), 5)
                d_610_v50_ = nw113_
                index102_ = default__.safeIndex(771, (d_610_v50_).length(0))
                (d_610_v50_)[index102_] = (d_608_v48_).set(True, False)
                index103_ = default__.safeIndex(771, (d_610_v50_).length(0))
                (d_610_v50_)[index103_] = (default__.fm14(d_585_v32_.f5, globalState)) | (d_608_v48_)
            elif True:
                d_611___mcc_h0_ = source8_.cf6
                d_612_cf6_ = d_611___mcc_h0_
                d_613_v51_: C2
                nw114_ = C2()
                nw114_.ctor__(p2, (0) - (d_585_v32_.f5), 645)
                d_613_v51_ = nw114_
                d_614_v52_: D3
                d_614_v52_ = D3_DC8(d_553_v1_)
                d_615_v53_: _dafny.Map
                d_615_v53_ = _dafny.Map({(len(d_553_v1_)) + ((0) - (d_585_v32_.f5)): D5_DC16(d_585_v32_.f5, (d_593_v37_)[default__.safeIndex(147, (d_593_v37_).length(0))], _dafny.SeqWithoutIsStrInference([default__.fm12((self).f3, (d_552_v0_).f6, (d_613_v51_).f6, globalState), default__.fm12((d_552_v0_).f6, p2, not((self).f3), globalState)]), d_614_v52_)})
                d_616_v54_: D5
                d_616_v54_ = D5_DC16(len(d_553_v1_), _dafny.MultiSet(d_598_v39_), d_612_cf6_, d_614_v52_)
                d_615_v53_ = (d_615_v53_).set(len(default__.fm15(globalState)), d_616_v54_)
                d_617_v55_: str
                d_617_v55_ = _dafny.CodePoint('i')
                d_618_v56_: C0
                nw115_ = C0()
                nw115_.ctor__(_dafny.MultiSet([d_617_v55_]))
                d_618_v56_ = nw115_
                d_619_v57_: _dafny.Array
                nw116_ = _dafny.Array(int(0), 21)
                d_619_v57_ = nw116_
                d_620_v58_: _dafny.Seq
                d_620_v58_ = _dafny.SeqWithoutIsStrInference([d_619_v57_])
                d_620_v58_ = _dafny.SeqWithoutIsStrInference([d_619_v57_, d_619_v57_, d_619_v57_, d_619_v57_, d_619_v57_])
            (globalState).f2 = (self).f3
        hi3_ = len(d_553_v1_)
        for d_621_i3_ in range(p1, hi3_):
            d_622_v59_: _dafny.Array
            def lambda29_(d_623_i4_):
                return not((self).f3)

            init14_ = lambda29_
            nw117_ = _dafny.Array(None, 9)
            for i0_14_ in range(nw117_.length(0)):
                nw117_[i0_14_] = init14_(i0_14_)
            d_622_v59_ = nw117_
            index104_ = default__.safeIndex(575, (d_622_v59_).length(0))
            (d_622_v59_)[index104_] = ((self).f3 if p2 else (self).f3)
            index105_ = default__.safeIndex(575, (d_622_v59_).length(0))
            (d_622_v59_)[index105_] = (p2) == ((d_552_v0_).f6)
            d_624_v60_: str
            d_624_v60_ = _dafny.CodePoint('w')
            d_625_v61_: C0
            nw118_ = C0()
            nw118_.ctor__(_dafny.MultiSet([_dafny.CodePoint('y'), d_624_v60_]))
            d_625_v61_ = nw118_
            d_626_v62_: _dafny.MultiSet
            d_626_v62_ = _dafny.MultiSet([d_625_v61_, d_625_v61_])
            d_627_v63_: C1
            nw119_ = C1()
            nw119_.ctor__((d_626_v62_).cardinality, 831)
            d_627_v63_ = nw119_
            nw120_ = C1()
            nw120_.ctor__(default__.fm3((D1_DC4(not(False))).cf4, globalState), d_621_i3_)
            d_627_v63_ = nw120_
            d_628_v64_: int
            d_628_v64_ = 149
            d_629_v65_: _dafny.Set
            d_629_v65_ = _dafny.Set({(d_552_v0_).f6})
            d_630_v66_: _dafny.Set
            d_630_v66_ = _dafny.Set({len(d_629_v65_)})
            d_631_v68_: _dafny.Set
            def iife43_():
                coll13_ = _dafny.Set()
                compr_13_: int
                for compr_13_ in _dafny.IntegerRange(595, 1):
                    d_632_v67_: int = compr_13_
                    if ((595) <= (d_632_v67_)) and ((d_632_v67_) < (1)):
                        coll13_ = coll13_.union(_dafny.Set([default__.safeDivisionInt(d_632_v67_, -992)]))
                return _dafny.Set(coll13_)
            d_631_v68_ = _dafny.Set({d_630_v66_, iife43_()
            })
            d_628_v64_ = ((0) - (len(d_631_v68_))) - (d_621_i3_)
            d_633_v69_: C0
            nw121_ = C0()
            nw121_.ctor__((d_625_v61_).f7)
            d_633_v69_ = nw121_
        d_634_v70_: _dafny.Seq
        d_634_v70_ = _dafny.SeqWithoutIsStrInference([True])
        d_635_v71_: D1
        d_635_v71_ = D1_DC4((d_634_v70_)[default__.safeIndex(p1, len(d_634_v70_))])
        d_636_v72_: D1
        d_636_v72_ = D1_DC5(d_635_v71_)
        d_636_v72_ = d_636_v72_
        d_637_v73_: str
        d_637_v73_ = _dafny.CodePoint('y')
        d_638_v74_: _dafny.Map
        d_638_v74_ = _dafny.Map({p2: d_637_v73_})
        d_638_v74_ = (d_638_v74_).set(p2, d_637_v73_)
        hi4_ = p1
        for d_639_i5_ in range(p1, hi4_):
            d_640_v75_: _dafny.Map
            d_640_v75_ = _dafny.Map({d_639_i5_: (d_552_v0_).f6})
            d_641_v76_: _dafny.MultiSet
            d_641_v76_ = _dafny.MultiSet([len(default__.fm16((d_552_v0_).f6, default__.fm3((d_552_v0_).f6, globalState), (d_552_v0_).f6, globalState)), 504, d_639_i5_])
            d_642_v77_: C1
            nw122_ = C1()
            nw122_.ctor__(-609, (len(d_640_v75_)) - ((d_641_v76_).cardinality))
            d_642_v77_ = nw122_
            d_643_v78_: _dafny.Map
            d_643_v78_ = _dafny.Map({(self).f3: d_639_i5_})
            d_643_v78_ = (d_643_v78_).set(True, (p1 if (d_552_v0_).f6 else p1))
            d_644_v79_: int
            d_644_v79_ = 35
            d_644_v79_ = (0) - (d_644_v79_)
            d_645_v80_: D0
            d_645_v80_ = D0_DC2()
            d_645_v80_ = d_645_v80_
        d_646_v81_: _dafny.Seq
        d_646_v81_ = _dafny.SeqWithoutIsStrInference([len(d_553_v1_), (0) - (p1)])
        r0 = (d_646_v81_) + (d_646_v81_)
        d_647_v82_: _dafny.Map
        d_647_v82_ = _dafny.Map({p1: d_637_v73_})
        d_648_v83_: _dafny.Map
        d_648_v83_ = _dafny.Map({d_647_v82_: (d_552_v0_).f6})
        r1 = d_648_v83_
        return r0, r1

    @property
    def f3(self):
        return self._f3
