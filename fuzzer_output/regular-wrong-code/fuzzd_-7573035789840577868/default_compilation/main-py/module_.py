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
        return not((_dafny.SeqWithoutIsStrInference([D0_DC1(False), D0_DC1(not(True))])) == ((_dafny.SeqWithoutIsStrInference([D0_DC1(False)])) + (_dafny.SeqWithoutIsStrInference([D0_DC1(not(not(True)))]))))

    @staticmethod
    def fm1(p0, globalState):
        return D1_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hu")), (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bgppqsue"))) if False else len(_dafny.SeqWithoutIsStrInference([_dafny.Map({850: 630})])))), -415)

    @staticmethod
    def fm2(globalState):
        return D1_DC6((0) - ((D11_DC30(False, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([not(False), False, True, False, True]))])), True)).cf53), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qipvo"))))

    @staticmethod
    def fm3(p0, p1, p2, globalState):
        return (((0) - (len(_dafny.Set({True})))) + (384)) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kifiqbht"))))

    @staticmethod
    def fm5(p0, globalState):
        source0_ = D0_DC2(D0_DC2(D0_DC2(D0_DC2(D0_DC1(True)))))
        if source0_.is_DC1:
            d_0___mcc_h0_ = source0_.cf1
            d_1_cf1_ = d_0___mcc_h0_
            return ((0) - (-50)) - (435)
        elif source0_.is_DC0:
            d_2___mcc_h1_ = source0_.cf0
            d_3_cf0_ = d_2___mcc_h1_
            return (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_4_i0_ in range(default__.abs(95))]))))
        elif True:
            d_5___mcc_h2_ = source0_.cf2
            d_6_cf2_ = d_5___mcc_h2_
            return (0) - (len((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False]))))

    @staticmethod
    def fm9(p0, p1, p2, globalState):
        return (((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rch")))) if True else (0) - ((_dafny.MultiSet([True])).cardinality))) - (len((_dafny.Set({len(_dafny.Map({161: False})), 240, len(_dafny.Map({False: len(_dafny.Set({not(True), False}))}))})) | (_dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpetoj"))), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_7_i0_ in range(default__.abs(216))]))})), 90}))))

    @staticmethod
    def fm10(globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(387, 476):
                d_8_v0_: int = compr_0_
                if ((387) <= (d_8_v0_)) and ((d_8_v0_) < (476)):
                    coll0_[(d_8_v0_) * (853)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_9_i0_ in range(default__.abs(758))]))
            return _dafny.Map(coll0_)
        return ((iife0_()
        ) | (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_10_i1_ in range(default__.abs(187))]))): 784}))) | ((_dafny.Map({572: len(_dafny.SeqWithoutIsStrInference([249]))})) | (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bkuqdvcwu")))})))

    @staticmethod
    def fm11(p0, p1, p2, p3, globalState):
        return D0_DC2(D0_DC0(False))

    @staticmethod
    def fm12(p0, p1, p2, globalState):
        source1_ = D0_DC0(True)
        if source1_.is_DC1:
            d_11___mcc_h0_ = source1_.cf1
            d_12_cf1_ = d_11___mcc_h0_
            return _dafny.CodePoint('a')
        elif source1_.is_DC0:
            d_13___mcc_h1_ = source1_.cf0
            d_14_cf0_ = d_13___mcc_h1_
            return _dafny.CodePoint('r')
        elif True:
            d_15___mcc_h2_ = source1_.cf2
            d_16_cf2_ = d_15___mcc_h2_
            return _dafny.CodePoint('b')

    @staticmethod
    def fm13(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([372 for d_17_i0_ in range(default__.abs(-955))])) + ((D9_DC24(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqeuo")))])), (_dafny.MultiSet([True])).cardinality, len(_dafny.Map({(_dafny.MultiSet([True])).cardinality: len(_dafny.Set({456, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dst"))), (0) - (-802), 587}))}))}))]))).cf42), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([True])).cardinality, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-557, -934])), 811, (0) - (len(_dafny.Set({-612, len(_dafny.SeqWithoutIsStrInference([False]))})))})), (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wggx")))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d")))])])

    @staticmethod
    def fm16(p0, p1, globalState):
        if True:
            if not(False):
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lr"))
            elif True:
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrdlm"))
        elif True:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mtnbe"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sunbm")))

    @staticmethod
    def fm17(p0, p1, globalState):
        return _dafny.Map({990: not((False) == (True))})

    @staticmethod
    def fm18(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.MultiSet([510])).Elements:
                d_18_v0_: int = compr_1_
                if (d_18_v0_) in (_dafny.MultiSet([510])):
                    coll1_ = coll1_.union(_dafny.Set([(d_18_v0_) - (len(_dafny.SeqWithoutIsStrInference([True])))]))
            return _dafny.Set(coll1_)
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: int
            for compr_2_ in (_dafny.Map({958: 174})).keys.Elements:
                d_19_v1_: int = compr_2_
                if (d_19_v1_) in (_dafny.Map({958: 174})):
                    coll2_ = coll2_.union(_dafny.Set([(d_19_v1_) - (688)]))
            return _dafny.Set(coll2_)
        return (iife1_()
        ) - (iife2_()
        )

    @staticmethod
    def fm19(p0, p1, globalState):
        return _dafny.SeqWithoutIsStrInference([(779 if True else 725), len(_dafny.Map({_dafny.CodePoint('k'): -617}))])

    @staticmethod
    def fm20(globalState):
        return _dafny.Set({(-355) >= (-391), True})

    @staticmethod
    def fm21(globalState):
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not(True)]))])).Elements:
                d_22_v0_: int = compr_3_
                if (d_22_v0_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not(True)]))])):
                    coll3_[(d_22_v0_) * (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcolpleqj"))})))] = not(True)
            return _dafny.Map(coll3_)
        return D1_DC5((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_20_i0_ in range(default__.abs(524))])), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqjjx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvlppdqpi")))), ((0) - ((_dafny.MultiSet([True])).cardinality)) * (len(_dafny.SeqWithoutIsStrInference([712, (0) - (len(_dafny.Map({867: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_21_i1_ in range(default__.abs(513))])): _dafny.CodePoint('u')}))}))), len(iife3_()
), len(_dafny.SeqWithoutIsStrInference([False, not(False)]))]))))

    @staticmethod
    def fm22(globalState):
        return _dafny.Map({(_dafny.Set({False})).issubset(_dafny.Set({False})): not((D12_DC33(-467, False, _dafny.MultiSet([-662]), False, True)).cf61)})

    @staticmethod
    def fm23(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([(0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_23_i0_ in range(default__.abs(-322))])), 119, 382, 970, len(_dafny.Map({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])]))]): False}))])))])) - (_dafny.MultiSet([len(_dafny.Map({-465: False})), -733, 919]))

    @staticmethod
    def fm24(p0, p1, globalState):
        def iife4_():
            coll4_ = _dafny.Map()
            compr_4_: int
            for compr_4_ in (_dafny.Set({len(_dafny.Map({True: len(_dafny.Set({_dafny.CodePoint('n')}))})), len(_dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgnsevrhf"))), 120])).cardinality: True}))})).Elements:
                d_24_v0_: int = compr_4_
                if (d_24_v0_) in (_dafny.Set({len(_dafny.Map({True: len(_dafny.Set({_dafny.CodePoint('n')}))})), len(_dafny.Map({(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgnsevrhf"))), 120])).cardinality: True}))})):
                    coll4_[default__.safeDivisionInt(d_24_v0_, -51)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uw")))
            return _dafny.Map(coll4_)
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: _dafny.Map
            for compr_5_ in (_dafny.MultiSet([_dafny.Map({False: 783})])).Elements:
                d_25_v1_: _dafny.Map = compr_5_
                if (d_25_v1_) in (_dafny.MultiSet([_dafny.Map({False: 783})])):
                    coll5_[d_25_v1_] = len(_dafny.SeqWithoutIsStrInference([365]))
            return _dafny.Map(coll5_)
        if (iife4_()
        ) != (_dafny.Map({-855: len(_dafny.Set({_dafny.Map({221: (0) - (len(iife5_()
        ))}), _dafny.Map({535: len(_dafny.SeqWithoutIsStrInference([False]))})}))})):
            return (_dafny.Map({False: 437})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference([785]))}))
        elif True:
            def iife6_():
                coll6_ = _dafny.Map()
                compr_6_: int
                for compr_6_ in _dafny.IntegerRange(773, 283):
                    d_26_v2_: int = compr_6_
                    if ((773) <= (d_26_v2_)) and ((d_26_v2_) < (283)):
                        coll6_[default__.safeDivisionInt(d_26_v2_, len(_dafny.SeqWithoutIsStrInference([not(False), True])))] = True
                return _dafny.Map(coll6_)
            return (_dafny.Map({False: (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eceq"))))})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(iife6_()
            ), 537]))}))

    @staticmethod
    def fm25(p0, p1, globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tusbtf"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hbix"))): 972})

    @staticmethod
    def fm26(p0, p1, p2, p3, globalState):
        return D6_DC20(not((654) > (len(_dafny.Set({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([-259]))}))})))), len(_dafny.Map({not(False): 422})), False)

    @staticmethod
    def fm27(p0, p1, p2, globalState):
        return (_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([True, False, False]))

    @staticmethod
    def fm28(p0, p1, globalState):
        return ((_dafny.MultiSet([False, not(True)])) - (_dafny.MultiSet([True]))).intersection(_dafny.MultiSet([True, True]))

    @staticmethod
    def fm29(p0, p1, globalState):
        def iife7_():
            def iife13_():
                def iife16_():
                    def iife17_():
                        coll17_ = _dafny.Set()
                        compr_17_: int
                        for compr_17_ in _dafny.IntegerRange(10, 160):
                            d_33_v3_: int = compr_17_
                            if ((10) <= (d_33_v3_)) and ((d_33_v3_) < (160)):
                                coll17_ = coll17_.union(_dafny.Set([(d_33_v3_) - (428)]))
                        return _dafny.Set(coll17_)
                    coll16_ = _dafny.Map()
                    compr_16_: int
                    for compr_16_ in _dafny.IntegerRange(292, 523):
                        d_27_v2_: int = compr_16_
                        if ((292) <= (d_27_v2_)) and ((d_27_v2_) < (523)):
                            coll16_[(d_27_v2_) - (len(_dafny.SeqWithoutIsStrInference([len(iife17_()
                            ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))])))] = True
                    return _dafny.Map(coll16_)
                coll13_ = _dafny.Map()
                def iife14_():
                    def iife15_():
                        coll15_ = _dafny.Set()
                        compr_15_: int
                        for compr_15_ in _dafny.IntegerRange(10, 160):
                            d_32_v3_: int = compr_15_
                            if ((10) <= (d_32_v3_)) and ((d_32_v3_) < (160)):
                                coll15_ = coll15_.union(_dafny.Set([(d_32_v3_) - (428)]))
                        return _dafny.Set(coll15_)
                    coll14_ = _dafny.Map()
                    compr_14_: int
                    for compr_14_ in _dafny.IntegerRange(292, 523):
                        d_27_v2_: int = compr_14_
                        if ((292) <= (d_27_v2_)) and ((d_27_v2_) < (523)):
                            coll14_[(d_27_v2_) - (len(_dafny.SeqWithoutIsStrInference([len(iife15_()
                            ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))])))] = True
                    return _dafny.Map(coll14_)
                compr_13_: _dafny.Map
                for compr_13_ in (_dafny.Map({(iife14_()
                ): (0) - (len(_dafny.Map({not(True): False})))})).keys.Elements:
                    d_29_v1_: _dafny.Map = compr_13_
                    if (d_29_v1_) in (_dafny.Map({(iife16_()
                    ): (0) - (len(_dafny.Map({not(True): False})))})):
                        coll13_[d_29_v1_] = 669
                return _dafny.Map(coll13_)
            coll7_ = _dafny.Map()
            def iife8_():
                def iife11_():
                    def iife12_():
                        coll12_ = _dafny.Set()
                        compr_12_: int
                        for compr_12_ in _dafny.IntegerRange(10, 160):
                            d_30_v3_: int = compr_12_
                            if ((10) <= (d_30_v3_)) and ((d_30_v3_) < (160)):
                                coll12_ = coll12_.union(_dafny.Set([(d_30_v3_) - (428)]))
                        return _dafny.Set(coll12_)
                    coll11_ = _dafny.Map()
                    compr_11_: int
                    for compr_11_ in _dafny.IntegerRange(292, 523):
                        d_27_v2_: int = compr_11_
                        if ((292) <= (d_27_v2_)) and ((d_27_v2_) < (523)):
                            coll11_[(d_27_v2_) - (len(_dafny.SeqWithoutIsStrInference([len(iife12_()
                            ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))])))] = True
                    return _dafny.Map(coll11_)
                coll8_ = _dafny.Map()
                def iife9_():
                    def iife10_():
                        coll10_ = _dafny.Set()
                        compr_10_: int
                        for compr_10_ in _dafny.IntegerRange(10, 160):
                            d_28_v3_: int = compr_10_
                            if ((10) <= (d_28_v3_)) and ((d_28_v3_) < (160)):
                                coll10_ = coll10_.union(_dafny.Set([(d_28_v3_) - (428)]))
                        return _dafny.Set(coll10_)
                    coll9_ = _dafny.Map()
                    compr_9_: int
                    for compr_9_ in _dafny.IntegerRange(292, 523):
                        d_27_v2_: int = compr_9_
                        if ((292) <= (d_27_v2_)) and ((d_27_v2_) < (523)):
                            coll9_[(d_27_v2_) - (len(_dafny.SeqWithoutIsStrInference([len(iife10_()
                            ), len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: True}))]))])))] = True
                    return _dafny.Map(coll9_)
                compr_8_: _dafny.Map
                for compr_8_ in (_dafny.Map({(iife9_()
                ): (0) - (len(_dafny.Map({not(True): False})))})).keys.Elements:
                    d_29_v1_: _dafny.Map = compr_8_
                    if (d_29_v1_) in (_dafny.Map({(iife11_()
                    ): (0) - (len(_dafny.Map({not(True): False})))})):
                        coll8_[d_29_v1_] = 669
                return _dafny.Map(coll8_)
            compr_7_: _dafny.Map
            for compr_7_ in (iife8_()
            ).keys.Elements:
                d_31_v0_: _dafny.Map = compr_7_
                if (d_31_v0_) in (iife13_()
                ):
                    coll7_[d_31_v0_] = (0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lukayiqqc")))])))
            return _dafny.Map(coll7_)
        return (iife7_()
        ) | (_dafny.Map({_dafny.Map({len(_dafny.Set({True, True, not(False), False})): True}): 502}))

    @staticmethod
    def fm30(p0, globalState):
        if True:
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: _dafny.Seq
                for compr_18_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(True), True]), _dafny.SeqWithoutIsStrInference([False, not(True), False, False])])).Elements:
                    d_34_v0_: _dafny.Seq = compr_18_
                    if (d_34_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not(True), True]), _dafny.SeqWithoutIsStrInference([False, not(True), False, False])])):
                        coll18_[d_34_v0_] = False
                return _dafny.Map(coll18_)
            return (iife18_()
            ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, not(not(False))]): False}))
        elif True:
            return _dafny.Map({_dafny.SeqWithoutIsStrInference([True]): False})

    @staticmethod
    def fm31(p0, p1, globalState):
        return D10_DC27((_dafny.MultiSet([False])) | (_dafny.MultiSet([False])))

    @staticmethod
    def fm32(p0, p1, p2, p3, globalState):
        return D2_DC7(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Map({not(False): _dafny.SeqWithoutIsStrInference([False])})))])))

    @staticmethod
    def fm33(p0, p1, globalState):
        return D10_DC28(False, (59) - (len(_dafny.Map({933: True}))))

    @staticmethod
    def fm34(p0, p1, p2, globalState):
        def iife19_():
            coll19_ = _dafny.Set()
            compr_19_: _dafny.Seq
            for compr_19_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False, not(False)]) for d_35_i0_ in range(default__.abs(609))])).Elements:
                d_36_v0_: _dafny.Seq = compr_19_
                if (d_36_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, False, not(False)]) for d_35_i0_ in range(default__.abs(609))])):
                    coll19_ = coll19_.union(_dafny.Set([d_36_v0_]))
            return _dafny.Set(coll19_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([False, True, True, False]), _dafny.SeqWithoutIsStrInference([True, False]), _dafny.SeqWithoutIsStrInference([not(True)])})).intersection(iife19_()
        )

    @staticmethod
    def m9(p0, p1, globalState):
        r0: bool = False
        r1: bool = False
        r2: int = int(0)
        d_37_v0_: _dafny.Seq
        d_37_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dcrcctt"))
        d_38_v1_: int
        d_38_v1_ = -916
        d_39_v2_: _dafny.Map
        d_39_v2_ = _dafny.Map({d_37_v0_: default__.fm12(d_38_v1_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_40_i1_ in range(default__.abs(255))]), d_37_v0_, globalState)})
        d_41_v3_: _dafny.Seq
        d_41_v3_ = _dafny.SeqWithoutIsStrInference([p1])
        d_42_v4_: _dafny.Set
        d_42_v4_ = _dafny.Set({(d_41_v3_)[default__.safeIndex(875, len(d_41_v3_))], p1, p1})
        hi0_ = len(d_42_v4_)
        for d_43_i0_ in range(len(d_39_v2_), hi0_):
            (globalState).f17 = p1
            d_44_v5_: _dafny.Seq
            d_44_v5_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([p1])).set(default__.safeIndex(d_38_v1_, len(_dafny.SeqWithoutIsStrInference([p1]))), p0), d_41_v3_])
            if not(not((not (False) or (p1)) not in ((d_44_v5_)[default__.safeIndex(d_38_v1_, len(d_44_v5_))]))):
                d_45_v6_: D0
                d_45_v6_ = D0_DC1(p0)
                d_46_v7_: C2
                nw0_ = C2()
                nw0_.ctor__(d_45_v6_, p1)
                d_46_v7_ = nw0_
                d_47_v8_: C5
                nw1_ = C5()
                nw1_.ctor__(d_43_i0_)
                d_47_v8_ = nw1_
                d_37_v0_ = d_37_v0_
                d_48_v9_: _dafny.Array
                def lambda0_(d_49_v8_):
                    def lambda1_(d_50_i2_):
                        return _dafny.Set({(d_49_v8_).f18})

                    return lambda1_

                init0_ = lambda0_(d_47_v8_)
                nw2_ = _dafny.Array(None, 9)
                for i0_0_ in range(nw2_.length(0)):
                    nw2_[i0_0_] = init0_(i0_0_)
                d_48_v9_ = nw2_
                d_51_v10_: str
                d_51_v10_ = _dafny.CodePoint('e')
                d_52_v11_: _dafny.Array
                nw3_ = _dafny.Array(None, 25)
                nw3_[int(0)] = False
                nw3_[int(1)] = False
                nw3_[int(2)] = (d_46_v7_).f21
                nw3_[int(3)] = (d_46_v7_).f21
                nw3_[int(4)] = p1
                nw3_[int(5)] = not((d_41_v3_)[default__.safeIndex(d_43_i0_, len(d_41_v3_))])
                nw3_[int(6)] = p0
                nw3_[int(7)] = (d_46_v7_).f21
                nw3_[int(8)] = (d_46_v7_).f21
                nw3_[int(9)] = False
                nw3_[int(10)] = (d_46_v7_).f21
                nw3_[int(11)] = False
                nw3_[int(12)] = p1
                nw3_[int(13)] = p1
                nw3_[int(14)] = p0
                nw3_[int(15)] = p1
                nw3_[int(16)] = p0
                nw3_[int(17)] = (d_46_v7_).f21
                nw3_[int(18)] = p0
                nw3_[int(19)] = (d_46_v7_).f21
                nw3_[int(20)] = True
                nw3_[int(21)] = default__.fm0(p1, d_51_v10_, (d_46_v7_).f21, globalState)
                nw3_[int(22)] = p0
                nw3_[int(23)] = (d_46_v7_).f21
                nw3_[int(24)] = p0
                d_52_v11_ = nw3_
                d_53_v12_: D1
                d_53_v12_ = D1_DC4((d_47_v8_).f18, p1, d_52_v11_)
                rhs0_ = d_37_v0_
                rhs1_ = d_48_v9_
                rhs2_ = d_53_v12_
                rhs3_ = d_43_i0_
                d_37_v0_ = rhs0_
                d_48_v9_ = rhs1_
                d_53_v12_ = rhs2_
                d_38_v1_ = rhs3_
                d_54_v13_: T0
                nw4_ = C0()
                nw4_.ctor__(p1)
                d_54_v13_ = nw4_
                d_55_v14_: C4
                nw5_ = C4()
                nw5_.ctor__((d_46_v7_).f21)
                d_55_v14_ = nw5_
                d_56_v15_: _dafny.Map
                d_56_v15_ = _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference([d_54_v13_, d_54_v13_, d_54_v13_]))): d_55_v14_})
                r2 = default__.fm3(len((d_56_v15_).set(d_43_i0_, d_55_v14_)), (0) - (d_43_i0_), p1, globalState)
            elif True:
                d_57_v16_: _dafny.Array
                def lambda2_(d_58_v0_):
                    def lambda3_(d_59_i3_):
                        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hi"))) + (d_58_v0_)

                    return lambda3_

                init1_ = lambda2_(d_37_v0_)
                nw6_ = _dafny.Array(None, 26)
                for i0_1_ in range(nw6_.length(0)):
                    nw6_[i0_1_] = init1_(i0_1_)
                d_57_v16_ = nw6_
                index0_ = default__.safeIndex(645, (d_57_v16_).length(0))
                (d_57_v16_)[index0_] = (d_37_v0_) + (d_37_v0_)
                d_60_v17_: D0
                d_60_v17_ = D0_DC1(p1)
                d_61_v18_: C2
                nw7_ = C2()
                nw7_.ctor__(d_60_v17_, p1)
                d_61_v18_ = nw7_
                d_62_v19_: _dafny.Array
                def lambda4_(d_63_v3_):
                    def lambda5_(d_64_i4_):
                        return d_63_v3_

                    return lambda5_

                init2_ = lambda4_(d_41_v3_)
                nw8_ = _dafny.Array(None, 14)
                for i0_2_ in range(nw8_.length(0)):
                    nw8_[i0_2_] = init2_(i0_2_)
                d_62_v19_ = nw8_
                d_65_v20_: D14
                d_65_v20_ = D14_DC41(d_61_v18_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sos")), d_62_v19_)
                pat_let_tv0_ = d_62_v19_
                pat_let_tv1_ = d_61_v18_
                d_66_v21_: _dafny.Map
                def iife20_(_pat_let0_0):
                    def iife21_(d_67_dt__update__tmp_h0_):
                        def iife22_(_pat_let1_0):
                            def iife23_(d_68_dt__update_hcf74_h0_):
                                def iife24_(_pat_let2_0):
                                    def iife25_(d_69_dt__update_hcf72_h0_):
                                        return D14_DC41(d_69_dt__update_hcf72_h0_, (d_67_dt__update__tmp_h0_).cf73, d_68_dt__update_hcf74_h0_)
                                    return iife25_(_pat_let2_0)
                                return iife24_(pat_let_tv1_)
                            return iife23_(_pat_let1_0)
                        return iife22_(pat_let_tv0_)
                    return iife21_(_pat_let0_0)
                d_66_v21_ = _dafny.Map({(iife20_(d_65_v20_)).cf73: d_37_v0_})
                d_70_v22_: _dafny.Set
                d_70_v22_ = _dafny.Set({d_38_v1_})
                index1_ = default__.safeIndex(645, (d_57_v16_).length(0))
                rhs4_ = ((d_66_v21_)[d_37_v0_] if (d_37_v0_) in (d_66_v21_) else d_37_v0_)
                rhs5_ = (d_61_v18_).f21
                rhs6_ = (d_61_v18_).f21
                rhs7_ = ((d_70_v22_) | (d_70_v22_)).issubset(d_70_v22_)
                lhs0_ = d_57_v16_
                lhs1_ = default__.safeIndex(645, (d_57_v16_).length(0))
                lhs2_ = globalState
                lhs3_ = globalState
                lhs4_ = globalState
                lhs0_[lhs1_] = rhs4_
                lhs2_.f17 = rhs5_
                lhs3_.f11 = rhs6_
                lhs4_.f15 = rhs7_
                d_71_v23_: _dafny.Map
                d_71_v23_ = _dafny.Map({p1: False})
                d_72_v24_: _dafny.Map
                d_72_v24_ = _dafny.Map({p0: d_71_v23_})
                d_72_v24_ = d_72_v24_
                rhs8_ = default__.fm3(default__.safeDivisionInt(d_43_i0_, d_38_v1_), (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_73_i5_ in range(default__.abs(148))]))) - (d_38_v1_), p0, globalState)
                rhs9_ = d_43_i0_
                r2 = rhs8_
                r2 = rhs9_
                d_74_v25_: C4
                nw9_ = C4()
                nw9_.ctor__((d_61_v18_).f21)
                d_74_v25_ = nw9_
                d_75_v26_: bool
                out0_: bool
                out0_ = (d_61_v18_).m5((p0) and (False), d_41_v3_, globalState)
                d_75_v26_ = out0_
            d_38_v1_ = default__.safeModuloInt(-352, len(_dafny.SeqWithoutIsStrInference([d_43_i0_])))
            d_76_v27_: _dafny.Array
            def lambda6_(d_77_p1_):
                def lambda7_(d_78_i6_):
                    return d_77_p1_

                return lambda7_

            init3_ = lambda6_(p1)
            nw10_ = _dafny.Array(None, 3)
            for i0_3_ in range(nw10_.length(0)):
                nw10_[i0_3_] = init3_(i0_3_)
            d_76_v27_ = nw10_
            index2_ = default__.safeIndex(689, (d_76_v27_).length(0))
            (d_76_v27_)[index2_] = p0
            index3_ = default__.safeIndex(689, (d_76_v27_).length(0))
            (d_76_v27_)[index3_] = p0
        d_79_v28_: _dafny.MultiSet
        d_79_v28_ = _dafny.MultiSet([p0])
        if (d_79_v28_).issubset((_dafny.MultiSet([p0])).intersection(d_79_v28_)):
            d_80_v29_: _dafny.Array
            nw11_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_80_v29_ = nw11_
            d_81_v30_: _dafny.MultiSet
            d_81_v30_ = _dafny.MultiSet([d_38_v1_])
            d_82_v31_: _dafny.Map
            d_82_v31_ = _dafny.Map({d_80_v29_: ((0) - ((d_81_v30_).cardinality)) * (d_38_v1_)})
            d_82_v31_ = (d_82_v31_).set(d_80_v29_, d_38_v1_)
            d_83_v32_: C0
            nw12_ = C0()
            nw12_.ctor__(p0)
            d_83_v32_ = nw12_
            d_84_v33_: _dafny.Set
            d_84_v33_ = _dafny.Set({d_83_v32_})
            d_38_v1_ = len(d_84_v33_)
            d_85_v34_: _dafny.Array
            nw13_ = _dafny.Array(_dafny.Seq({}), 8)
            d_85_v34_ = nw13_
            d_86_v35_: T0
            nw14_ = C4()
            nw14_.ctor__((d_83_v32_).f23)
            d_86_v35_ = nw14_
            d_87_v36_: _dafny.Seq
            d_87_v36_ = _dafny.SeqWithoutIsStrInference([d_86_v35_, d_86_v35_, d_86_v35_])
            index4_ = default__.safeIndex(361, (d_85_v34_).length(0))
            (d_85_v34_)[index4_] = d_87_v36_
            index5_ = default__.safeIndex(361, (d_85_v34_).length(0))
            rhs10_ = (d_38_v1_) > ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gvcj"))) if True else (0) - (d_38_v1_)))
            rhs11_ = (429) != (d_38_v1_)
            rhs12_ = d_87_v36_
            lhs5_ = globalState
            lhs6_ = globalState
            lhs7_ = d_85_v34_
            lhs8_ = default__.safeIndex(361, (d_85_v34_).length(0))
            lhs5_.f17 = rhs10_
            lhs6_.f11 = rhs11_
            lhs7_[lhs8_] = rhs12_
            d_88_v37_: D12
            d_88_v37_ = D12_DC34()
            source2_ = (d_88_v37_ if (d_83_v32_).f23 else d_88_v37_)
            if source2_.is_DC32:
                d_89___mcc_h0_ = source2_.cf56
                d_90_cf56_ = d_89___mcc_h0_
                d_91_v38_: D0
                d_91_v38_ = D0_DC1(p1)
                d_92_v39_: C2
                nw15_ = C2()
                nw15_.ctor__(d_91_v38_, (d_83_v32_).f23)
                d_92_v39_ = nw15_
                r2 = d_90_cf56_
                (globalState).f11 = True
                d_93_v40_: _dafny.Map
                d_93_v40_ = _dafny.Map({(d_92_v39_).f21: d_90_cf56_})
                d_94_v41_: _dafny.Array
                nw16_ = _dafny.Array(int(0), 4)
                d_94_v41_ = nw16_
                d_95_v42_: D3
                d_95_v42_ = D3_DC11(d_93_v40_, d_90_cf56_, d_94_v41_)
                rhs13_ = d_38_v1_
                rhs14_ = (0) - ((((_dafny.MultiSet([(d_83_v32_).f23, (d_83_v32_).f23])).intersection(_dafny.MultiSet([(d_83_v32_).f23, p0]))).cardinality) + (len(d_42_v4_)))
                rhs15_ = d_95_v42_
                d_90_cf56_ = rhs13_
                d_38_v1_ = rhs14_
                d_95_v42_ = rhs15_
            elif source2_.is_DC33:
                d_96___mcc_h1_ = source2_.cf57
                d_97___mcc_h2_ = source2_.cf58
                d_98___mcc_h3_ = source2_.cf59
                d_99___mcc_h4_ = source2_.cf60
                d_100___mcc_h5_ = source2_.cf61
                d_101_cf61_ = d_100___mcc_h5_
                d_102_cf60_ = d_99___mcc_h4_
                d_103_cf59_ = d_98___mcc_h3_
                d_104_cf58_ = d_97___mcc_h2_
                d_105_cf57_ = d_96___mcc_h1_
                r2 = d_38_v1_
                d_106_v43_: _dafny.Map
                d_106_v43_ = _dafny.Map({d_83_v32_: default__.fm20(globalState)})
                d_42_v4_ = ((d_42_v4_) - (d_42_v4_)) | (((d_106_v43_)[d_83_v32_] if (d_83_v32_) in (d_106_v43_) else d_42_v4_))
                d_107_v44_: str
                d_107_v44_ = _dafny.CodePoint('o')
                d_108_v45_: _dafny.Map
                d_108_v45_ = _dafny.Map({d_107_v44_: d_107_v44_})
                d_108_v45_ = d_108_v45_
                d_109_v46_: D11
                d_109_v46_ = D11_DC29(d_42_v4_)
                d_110_v47_: _dafny.Seq
                d_110_v47_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_105_cf57_: len((d_109_v46_).cf51)})])
                d_111_v48_: D0
                d_111_v48_ = D0_DC1(d_102_cf60_)
                d_112_v49_: C2
                nw17_ = C2()
                nw17_.ctor__(d_111_v48_, d_104_cf58_)
                d_112_v49_ = nw17_
                d_113_v50_: _dafny.Array
                def lambda8_(d_114_cf60_):
                    def lambda9_(d_115_i7_):
                        return _dafny.SeqWithoutIsStrInference([d_114_cf60_])

                    return lambda9_

                init4_ = lambda8_(d_102_cf60_)
                nw18_ = _dafny.Array(None, 9)
                for i0_4_ in range(nw18_.length(0)):
                    nw18_[i0_4_] = init4_(i0_4_)
                d_113_v50_ = nw18_
                d_116_v51_: _dafny.MultiSet
                d_116_v51_ = _dafny.MultiSet([D14_DC41(d_112_v49_, d_37_v0_, d_113_v50_)])
                d_117_v52_: _dafny.Map
                d_117_v52_ = _dafny.Map({len(default__.fm27(d_105_cf57_, d_105_cf57_, len((d_110_v47_)[default__.safeIndex((d_116_v51_).cardinality, len(d_110_v47_))]), globalState)): 533})
                d_118_v53_: _dafny.Set
                d_118_v53_ = _dafny.Set({d_117_v52_})
                d_119_v54_: D6
                d_119_v54_ = D6_DC20((d_83_v32_).f23, d_105_cf57_, p0)
                d_120_v55_: _dafny.Map
                d_120_v55_ = _dafny.Map({p0: 72})
                d_121_v56_: _dafny.Set
                d_121_v56_ = _dafny.Set({d_38_v1_, (default__.fm23((d_83_v32_).f23, d_120_v55_, d_38_v1_, d_107_v44_, globalState)).cardinality})
                d_122_v57_: _dafny.MultiSet
                d_122_v57_ = _dafny.MultiSet([d_120_v55_])
                d_123_v58_: _dafny.Seq
                d_123_v58_ = _dafny.SeqWithoutIsStrInference([949])
                d_124_v59_: _dafny.Array
                nw19_ = _dafny.Array(None, 26)
                nw19_[int(0)] = (d_83_v32_).fm14(len(d_118_v53_), globalState)
                nw19_[int(1)] = d_105_cf57_
                nw19_[int(2)] = d_38_v1_
                nw19_[int(3)] = d_105_cf57_
                nw19_[int(4)] = d_38_v1_
                nw19_[int(5)] = (d_119_v54_).cf37
                nw19_[int(6)] = d_105_cf57_
                nw19_[int(7)] = default__.safeDivisionInt(d_38_v1_, d_38_v1_)
                nw19_[int(8)] = (d_38_v1_ if (d_83_v32_).f23 else d_38_v1_)
                nw19_[int(9)] = d_105_cf57_
                nw19_[int(10)] = d_105_cf57_
                nw19_[int(11)] = (D4_DC14(p0, d_38_v1_, d_121_v56_, d_107_v44_)).cf24
                nw19_[int(12)] = d_105_cf57_
                nw19_[int(13)] = d_105_cf57_
                nw19_[int(14)] = d_38_v1_
                nw19_[int(15)] = d_105_cf57_
                nw19_[int(16)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kuj")) if (d_112_v49_).f21 else d_37_v0_))
                nw19_[int(17)] = 12
                nw19_[int(18)] = (((d_122_v57_).set(d_120_v55_, default__.abs(len(d_123_v58_)))) | (d_122_v57_)).cardinality
                nw19_[int(19)] = (0) - ((d_79_v28_).cardinality)
                nw19_[int(20)] = d_38_v1_
                nw19_[int(21)] = len(d_37_v0_)
                nw19_[int(22)] = d_105_cf57_
                nw19_[int(23)] = d_105_cf57_
                nw19_[int(24)] = (0) - (d_105_cf57_)
                nw19_[int(25)] = d_38_v1_
                d_124_v59_ = nw19_
                d_125_v60_: _dafny.Map
                d_125_v60_ = _dafny.Map({(d_105_cf57_) + (d_38_v1_): d_124_v59_})
                rhs16_ = d_38_v1_
                rhs17_ = ((d_125_v60_)[(0) - (default__.safeDivisionInt(d_105_cf57_, d_105_cf57_))] if ((0) - (default__.safeDivisionInt(d_105_cf57_, d_105_cf57_))) in (d_125_v60_) else d_124_v59_)
                r2 = rhs16_
                d_124_v59_ = rhs17_
            elif source2_.is_DC34:
                d_126_v61_: _dafny.Map
                d_126_v61_ = _dafny.Map({p1: d_38_v1_})
                d_127_v62_: _dafny.Array
                nw20_ = _dafny.Array(int(0), 8)
                d_127_v62_ = nw20_
                d_128_v63_: D3
                d_128_v63_ = D3_DC11(d_126_v61_, len(d_37_v0_), d_127_v62_)
                d_129_v64_: C1
                nw21_ = C1()
                nw21_.ctor__((d_128_v63_).cf20)
                d_129_v64_ = nw21_
                d_130_v65_: _dafny.Seq
                d_130_v65_ = _dafny.SeqWithoutIsStrInference([d_129_v64_])
                (globalState).f15 = (len((d_130_v65_) + (d_130_v65_))) > (d_38_v1_)
                d_131_v66_: _dafny.Seq
                d_131_v66_ = _dafny.SeqWithoutIsStrInference([d_38_v1_])
                d_132_v67_: _dafny.Seq
                d_132_v67_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p1: (d_83_v32_).f23})]))) for d_133_i8_ in range(default__.abs(513))])).set(default__.safeIndex(d_38_v1_, len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p1: (d_83_v32_).f23})]))) for d_134_i8_ in range(default__.abs(513))]))), d_38_v1_), d_131_v66_, d_131_v66_])
                d_135_v68_: _dafny.MultiSet
                d_135_v68_ = _dafny.MultiSet([(d_132_v67_)[default__.safeIndex(d_38_v1_, len(d_132_v67_))], d_131_v66_])
                d_135_v68_ = ((d_135_v68_) | (d_135_v68_)) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_38_v1_ for d_136_i9_ in range(default__.abs(648))])])) | (d_135_v68_))
                d_137_v69_: _dafny.Array
                def lambda10_(d_138_v32_):
                    def lambda11_(d_139_i10_):
                        return ((d_138_v32_).f23) or ((d_138_v32_).f23)

                    return lambda11_

                init5_ = lambda10_(d_83_v32_)
                nw22_ = _dafny.Array(None, 25)
                for i0_5_ in range(nw22_.length(0)):
                    nw22_[i0_5_] = init5_(i0_5_)
                d_137_v69_ = nw22_
                index6_ = default__.safeIndex(922, (d_137_v69_).length(0))
                (d_137_v69_)[index6_] = (d_83_v32_).f23
                d_140_v70_: _dafny.Set
                d_140_v70_ = _dafny.Set({d_41_v3_})
                index7_ = default__.safeIndex(922, (d_137_v69_).length(0))
                rhs18_ = d_86_v35_
                rhs19_ = (default__.fm34((d_131_v66_)[default__.safeIndex(d_38_v1_, len(d_131_v66_))], d_38_v1_, True, globalState)).issubset(d_140_v70_)
                rhs20_ = p0
                lhs9_ = d_137_v69_
                lhs10_ = default__.safeIndex(922, (d_137_v69_).length(0))
                lhs11_ = globalState
                d_86_v35_ = rhs18_
                lhs9_[lhs10_] = rhs19_
                lhs11_.f11 = rhs20_
                d_141_v71_: D12
                d_141_v71_ = D12_DC33(d_38_v1_, False, d_81_v30_, p1, p0)
                d_38_v1_ = (d_141_v71_).cf57
            elif source2_.is_DC31:
                d_142___mcc_h6_ = source2_.cf55
                d_143_cf55_ = d_142___mcc_h6_
                d_144_v72_: _dafny.Array
                def lambda12_(d_145_v1_):
                    def lambda13_(d_146_i11_):
                        return default__.safeDivisionInt(d_146_i11_, d_145_v1_)

                    return lambda13_

                init6_ = lambda12_(d_38_v1_)
                nw23_ = _dafny.Array(None, 22)
                for i0_6_ in range(nw23_.length(0)):
                    nw23_[i0_6_] = init6_(i0_6_)
                d_144_v72_ = nw23_
                d_147_v73_: C1
                nw24_ = C1()
                nw24_.ctor__((d_144_v72_ if p1 else d_144_v72_))
                d_147_v73_ = nw24_
                d_148_v74_: D0
                d_148_v74_ = D0_DC1((d_83_v32_).f23)
                d_149_v75_: C2
                nw25_ = C2()
                nw25_.ctor__(d_148_v74_, ((d_83_v32_).f23) == (p1))
                d_149_v75_ = nw25_
                d_149_v75_ = d_149_v75_
                (globalState).f11 = (d_83_v32_).f23
                d_150_v76_: D12
                d_150_v76_ = D12_DC32(d_38_v1_)
                d_151_v77_: _dafny.Map
                d_151_v77_ = _dafny.Map({d_150_v76_: d_144_v72_})
                d_144_v72_ = ((d_151_v77_)[d_150_v76_] if (d_150_v76_) in (d_151_v77_) else (d_147_v73_).f22)
            elif True:
                d_152___mcc_h7_ = source2_.cf62
                d_153_cf62_ = d_152___mcc_h7_
                (globalState).f11 = not (p0) or ((d_83_v32_).f23)
                d_154_v78_: _dafny.Array
                nw26_ = _dafny.Array(int(0), 26)
                d_154_v78_ = nw26_
                d_155_v79_: C1
                nw27_ = C1()
                nw27_.ctor__(d_154_v78_)
                d_155_v79_ = nw27_
                d_156_v80_: _dafny.Map
                d_156_v80_ = _dafny.Map({d_155_v79_: (0) - (d_38_v1_)})
                d_38_v1_ = (0) - (len((d_156_v80_) | ((d_156_v80_).set(d_155_v79_, 161))))
                d_157_v81_: str
                d_157_v81_ = _dafny.CodePoint('x')
                d_157_v81_ = d_157_v81_
                d_158_v82_: _dafny.Map
                d_158_v82_ = _dafny.Map({p1: d_38_v1_})
                d_158_v82_ = (d_158_v82_).set((d_83_v32_).f23, (0) - (default__.safeModuloInt(d_38_v1_, d_38_v1_)))
            (globalState).f11 = p1
        elif True:
            d_159_v83_: _dafny.Array
            nw28_ = _dafny.Array(False, 8)
            d_159_v83_ = nw28_
            index8_ = default__.safeIndex(434, (d_159_v83_).length(0))
            (d_159_v83_)[index8_] = p1
            index9_ = default__.safeIndex(434, (d_159_v83_).length(0))
            (d_159_v83_)[index9_] = not((d_38_v1_) < (d_38_v1_))
            index10_ = default__.safeIndex(434, (d_159_v83_).length(0))
            (d_159_v83_)[index10_] = (d_159_v83_)[default__.safeIndex(434, (d_159_v83_).length(0))]
            d_160_v84_: _dafny.Seq
            d_160_v84_ = _dafny.SeqWithoutIsStrInference([d_37_v0_, d_37_v0_, d_37_v0_, d_37_v0_])
            d_161_v85_: _dafny.Seq
            d_161_v85_ = _dafny.SeqWithoutIsStrInference([d_160_v84_])
            d_160_v84_ = (d_161_v85_)[default__.safeIndex(len(d_37_v0_), len(d_161_v85_))]
            d_162_v86_: _dafny.Seq
            d_162_v86_ = _dafny.SeqWithoutIsStrInference([(0) - (d_38_v1_), 256, len(d_37_v0_), d_38_v1_])
            (globalState).f11 = ((default__.fm19(_dafny.SeqWithoutIsStrInference([d_38_v1_]), p1, globalState)).set(default__.safeIndex(d_38_v1_, len(default__.fm19(_dafny.SeqWithoutIsStrInference([d_38_v1_]), p1, globalState))), d_38_v1_)) == (d_162_v86_)
            d_163_v87_: _dafny.Map
            d_163_v87_ = _dafny.Map({p1: d_38_v1_})
            d_163_v87_ = (d_163_v87_) | ((d_163_v87_) | (d_163_v87_))
        d_41_v3_ = d_41_v3_
        (globalState).f15 = p0
        d_164_v88_: _dafny.Array
        nw29_ = _dafny.Array(None, 29)
        d_164_v88_ = nw29_
        d_165_v89_: _dafny.Array
        nw30_ = _dafny.Array(None, 8)
        nw30_[int(0)] = d_38_v1_
        nw30_[int(1)] = d_38_v1_
        nw30_[int(2)] = d_38_v1_
        nw30_[int(3)] = d_38_v1_
        nw30_[int(4)] = d_38_v1_
        nw30_[int(5)] = d_38_v1_
        nw30_[int(6)] = d_38_v1_
        nw30_[int(7)] = d_38_v1_
        d_165_v89_ = nw30_
        d_166_v90_: D13
        d_166_v90_ = D13_DC38(default__.fm12(-248, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_167_i12_ in range(default__.abs(-943))]), d_37_v0_, globalState), p0, (0) - (d_38_v1_), d_165_v89_)
        d_168_v91_: C1
        nw31_ = C1()
        nw31_.ctor__((d_166_v90_).cf69)
        d_168_v91_ = nw31_
        index11_ = default__.safeIndex(771, (d_164_v88_).length(0))
        (d_164_v88_)[index11_] = d_168_v91_
        index12_ = default__.safeIndex(771, (d_164_v88_).length(0))
        (d_164_v88_)[index12_] = d_168_v91_
        (globalState).f11 = (d_38_v1_) == ((len(_dafny.SeqWithoutIsStrInference([False, p1]))) * (d_38_v1_))
        r0 = not((d_38_v1_) < (d_38_v1_))
        r1 = p1
        r2 = d_38_v1_
        return r0, r1, r2

    @staticmethod
    def Main(noArgsParameter__):
        d_169_v0_: bool
        d_169_v0_ = True
        d_170_v1_: _dafny.Seq
        d_170_v1_ = _dafny.SeqWithoutIsStrInference([d_169_v0_, d_169_v0_])
        d_171_v2_: int
        d_171_v2_ = 644
        d_172_v3_: _dafny.Array
        nw32_ = _dafny.Array(int(0), 3)
        d_172_v3_ = nw32_
        d_173_v4_: _dafny.Array
        nw33_ = _dafny.Array(None, 28)
        nw33_[int(0)] = d_172_v3_
        nw33_[int(1)] = d_172_v3_
        nw33_[int(2)] = d_172_v3_
        nw33_[int(3)] = d_172_v3_
        nw33_[int(4)] = d_172_v3_
        nw33_[int(5)] = d_172_v3_
        nw33_[int(6)] = d_172_v3_
        nw33_[int(7)] = d_172_v3_
        nw33_[int(8)] = d_172_v3_
        nw33_[int(9)] = d_172_v3_
        nw33_[int(10)] = d_172_v3_
        nw33_[int(11)] = d_172_v3_
        nw33_[int(12)] = d_172_v3_
        nw33_[int(13)] = d_172_v3_
        nw33_[int(14)] = d_172_v3_
        nw33_[int(15)] = d_172_v3_
        nw33_[int(16)] = d_172_v3_
        nw33_[int(17)] = d_172_v3_
        nw33_[int(18)] = d_172_v3_
        nw33_[int(19)] = d_172_v3_
        nw33_[int(20)] = d_172_v3_
        nw33_[int(21)] = d_172_v3_
        nw33_[int(22)] = d_172_v3_
        nw33_[int(23)] = d_172_v3_
        nw33_[int(24)] = d_172_v3_
        nw33_[int(25)] = d_172_v3_
        nw33_[int(26)] = d_172_v3_
        nw33_[int(27)] = d_172_v3_
        d_173_v4_ = nw33_
        d_174_v5_: D0
        d_174_v5_ = D0_DC1(d_169_v0_)
        d_175_v6_: _dafny.Seq
        d_175_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "df"))
        d_176_v7_: _dafny.Map
        d_176_v7_ = _dafny.Map({d_171_v2_: d_171_v2_})
        d_177_globalState_: GlobalState
        nw34_ = GlobalState()
        nw34_.ctor__(_dafny.CodePoint('x'), (d_170_v1_).set(default__.safeIndex(d_171_v2_, len(d_170_v1_)), d_169_v0_), d_173_v4_, False, (d_170_v1_) + (_dafny.SeqWithoutIsStrInference([d_169_v0_])), _dafny.CodePoint('e'), -835, False, (d_175_v6_ if (d_174_v5_).cf1 else d_175_v6_), False, True, False, d_176_v7_, True, -305, True, (d_170_v1_) + (d_170_v1_), False)
        d_177_globalState_ = nw34_
        d_178_v8_: str
        d_178_v8_ = _dafny.CodePoint('c')
        if default__.fm0(d_169_v0_, d_178_v8_, d_169_v0_, d_177_globalState_):
            d_179_v9_: _dafny.Map
            d_179_v9_ = _dafny.Map({False: d_169_v0_})
            d_180_v10_: _dafny.Seq
            d_180_v10_ = _dafny.SeqWithoutIsStrInference([d_171_v2_])
            d_181_v11_: D0
            d_181_v11_ = D0_DC0(d_169_v0_)
            d_182_v12_: _dafny.Array
            nw35_ = _dafny.Array(None, 15)
            nw35_[int(0)] = (d_170_v1_)[default__.safeIndex(470, len(d_170_v1_))]
            nw35_[int(1)] = d_169_v0_
            nw35_[int(2)] = d_169_v0_
            nw35_[int(3)] = False
            nw35_[int(4)] = False
            nw35_[int(5)] = d_169_v0_
            nw35_[int(6)] = not (d_169_v0_) or (d_169_v0_)
            nw35_[int(7)] = d_169_v0_
            nw35_[int(8)] = (True) in (d_179_v9_)
            nw35_[int(9)] = d_169_v0_
            nw35_[int(10)] = d_169_v0_
            nw35_[int(11)] = (d_180_v10_) < (d_180_v10_)
            nw35_[int(12)] = d_169_v0_
            nw35_[int(13)] = (D0_DC1((d_181_v11_).cf0)).cf1
            nw35_[int(14)] = d_169_v0_
            d_182_v12_ = nw35_
            index13_ = default__.safeIndex(513, (d_182_v12_).length(0))
            (d_182_v12_)[index13_] = d_169_v0_
            index14_ = default__.safeIndex(513, (d_182_v12_).length(0))
            (d_182_v12_)[index14_] = d_169_v0_
            rhs21_ = _dafny.SeqWithoutIsStrInference([d_178_v8_ for d_183_i0_ in range(default__.abs(829))])
            rhs22_ = -916
            d_175_v6_ = rhs21_
            d_171_v2_ = rhs22_
            index15_ = default__.safeIndex(513, (d_182_v12_).length(0))
            (d_182_v12_)[index15_] = (d_169_v0_) or ((d_182_v12_)[default__.safeIndex(513, (d_182_v12_).length(0))])
            d_184_v13_: C3
            nw36_ = C3()
            nw36_.ctor__()
            d_184_v13_ = nw36_
            d_185_v14_: D6
            d_185_v14_ = D6_DC20(d_169_v0_, d_171_v2_, d_169_v0_)
            d_186_v15_: T0
            nw37_ = C2()
            nw37_.ctor__(d_174_v5_, (d_185_v14_).cf36)
            d_186_v15_ = nw37_
            d_187_v16_: _dafny.Array
            nw38_ = _dafny.Array(_dafny.Array(None, 0), 9)
            d_187_v16_ = nw38_
            index16_ = default__.safeIndex(31, (d_187_v16_).length(0))
            (d_187_v16_)[index16_] = (d_182_v12_ if d_169_v0_ else d_182_v12_)
            d_188_v17_: _dafny.Array
            nw39_ = _dafny.Array(None, 8)
            d_188_v17_ = nw39_
            d_189_v18_: T1
            nw40_ = C4()
            nw40_.ctor__((d_182_v12_)[default__.safeIndex(513, (d_182_v12_).length(0))])
            d_189_v18_ = nw40_
            index17_ = default__.safeIndex(74, (d_188_v17_).length(0))
            (d_188_v17_)[index17_] = d_189_v18_
            d_190_v19_: _dafny.Seq
            d_190_v19_ = _dafny.SeqWithoutIsStrInference([d_186_v15_, d_186_v15_, d_186_v15_])
            index18_ = default__.safeIndex(31, (d_187_v16_).length(0))
            index19_ = default__.safeIndex(74, (d_188_v17_).length(0))
            rhs23_ = d_182_v12_
            rhs24_ = (d_190_v19_)[default__.safeIndex(d_171_v2_, len(d_190_v19_))]
            rhs25_ = d_182_v12_
            rhs26_ = d_169_v0_
            rhs27_ = d_189_v18_
            lhs12_ = d_187_v16_
            lhs13_ = default__.safeIndex(31, (d_187_v16_).length(0))
            lhs14_ = d_177_globalState_
            lhs15_ = d_188_v17_
            lhs16_ = default__.safeIndex(74, (d_188_v17_).length(0))
            d_182_v12_ = rhs23_
            d_186_v15_ = rhs24_
            lhs12_[lhs13_] = rhs25_
            lhs14_.f17 = rhs26_
            lhs15_[lhs16_] = rhs27_
        elif True:
            d_191_v20_: _dafny.Array
            def lambda14_(d_192_v8_):
                def lambda15_(d_193_i1_):
                    return d_192_v8_

                return lambda15_

            init7_ = lambda14_(d_178_v8_)
            nw41_ = _dafny.Array(None, 13)
            for i0_7_ in range(nw41_.length(0)):
                nw41_[i0_7_] = init7_(i0_7_)
            d_191_v20_ = nw41_
            index20_ = default__.safeIndex(872, (d_191_v20_).length(0))
            (d_191_v20_)[index20_] = d_178_v8_
            index21_ = default__.safeIndex(872, (d_191_v20_).length(0))
            (d_191_v20_)[index21_] = d_178_v8_
            d_194_v21_: _dafny.Map
            d_194_v21_ = _dafny.Map({d_169_v0_: (0) - (d_171_v2_)})
            d_195_v22_: D13
            d_195_v22_ = D13_DC38((d_191_v20_)[default__.safeIndex(872, (d_191_v20_).length(0))], (len(d_194_v21_)) > (len((_dafny.SeqWithoutIsStrInference([d_178_v8_ for d_196_i2_ in range(default__.abs(191))])).set(default__.safeIndex(701, len(_dafny.SeqWithoutIsStrInference([d_178_v8_ for d_197_i2_ in range(default__.abs(191))]))), (d_191_v20_)[default__.safeIndex(872, (d_191_v20_).length(0))]))), (d_171_v2_) * ((0) - (d_171_v2_)), d_172_v3_)
            d_195_v22_ = d_195_v22_
            d_198_v23_: _dafny.Seq
            d_198_v23_ = _dafny.SeqWithoutIsStrInference([(d_171_v2_) - (len(d_194_v21_)), (d_171_v2_) - (409), default__.safeModuloInt(d_171_v2_, d_171_v2_), (d_171_v2_) - (default__.fm9(d_171_v2_, d_171_v2_, False, d_177_globalState_))])
            d_198_v23_ = (d_198_v23_) + ((_dafny.SeqWithoutIsStrInference([d_171_v2_ for d_199_i3_ in range(default__.abs(963))])) + (_dafny.SeqWithoutIsStrInference([d_171_v2_])))
            d_200_v24_: _dafny.Set
            d_200_v24_ = _dafny.Set({len(((d_175_v6_) + (d_175_v6_)).set(default__.safeIndex(d_171_v2_, len((d_175_v6_) + (d_175_v6_))), (d_191_v20_)[default__.safeIndex(872, (d_191_v20_).length(0))])), d_171_v2_, (len(d_198_v23_)) * (d_171_v2_), d_171_v2_, (D3_DC11(d_194_v21_, d_171_v2_, d_172_v3_)).cf19})
            d_201_v25_: _dafny.Seq
            d_201_v25_ = _dafny.SeqWithoutIsStrInference([d_200_v24_])
            d_200_v24_ = (d_201_v25_)[default__.safeIndex(d_171_v2_, len(d_201_v25_))]
            d_202_v26_: C2
            nw42_ = C2()
            nw42_.ctor__(d_174_v5_, d_169_v0_)
            d_202_v26_ = nw42_
            d_203_v27_: _dafny.Array
            def lambda16_(d_204_v1_):
                def lambda17_(d_205_i4_):
                    return d_204_v1_

                return lambda17_

            init8_ = lambda16_(d_170_v1_)
            nw43_ = _dafny.Array(None, 11)
            for i0_8_ in range(nw43_.length(0)):
                nw43_[i0_8_] = init8_(i0_8_)
            d_203_v27_ = nw43_
            d_206_v28_: D14
            d_206_v28_ = D14_DC41(d_202_v26_, d_175_v6_, d_203_v27_)
            source3_ = d_206_v28_
            if source3_.is_DC41:
                d_207___mcc_h0_ = source3_.cf72
                d_208___mcc_h1_ = source3_.cf73
                d_209___mcc_h2_ = source3_.cf74
                d_210_cf74_ = d_209___mcc_h2_
                d_211_cf73_ = d_208___mcc_h1_
                d_212_cf72_ = d_207___mcc_h0_
                d_213_v29_: _dafny.Array
                nw44_ = _dafny.Array(_dafny.MultiSet({}), 21)
                d_213_v29_ = nw44_
                d_214_v31_: D4
                def iife26_():
                    coll20_ = _dafny.Set()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(602, 158):
                        d_215_v30_: int = compr_20_
                        if ((602) <= (d_215_v30_)) and ((d_215_v30_) < (158)):
                            coll20_ = coll20_.union(_dafny.Set([(d_215_v30_) + (d_171_v2_)]))
                    return _dafny.Set(coll20_)
                d_214_v31_ = D4_DC14((d_212_cf72_).f21, d_171_v2_, iife26_()
, (d_191_v20_)[default__.safeIndex(872, (d_191_v20_).length(0))])
                d_216_v32_: D10
                d_216_v32_ = D10_DC28((d_202_v26_).f21, len(d_194_v21_))
                d_217_v33_: _dafny.Map
                d_217_v33_ = _dafny.Map({(d_216_v32_).cf49: (d_202_v26_).f21})
                d_218_v34_: _dafny.Map
                d_218_v34_ = _dafny.Map({(d_212_cf72_).f21: ((d_217_v33_)[default__.fm0((d_202_v26_).f21, (d_191_v20_)[default__.safeIndex(872, (d_191_v20_).length(0))], (d_202_v26_).f21, d_177_globalState_)] if (default__.fm0((d_202_v26_).f21, (d_191_v20_)[default__.safeIndex(872, (d_191_v20_).length(0))], (d_202_v26_).f21, d_177_globalState_)) in (d_217_v33_) else not((d_202_v26_).f21))})
                d_219_v35_: _dafny.Array
                nw45_ = _dafny.Array(None, 19)
                nw45_[int(0)] = not((d_214_v31_).cf23)
                nw45_[int(1)] = (d_212_cf72_).f21
                nw45_[int(2)] = (d_212_cf72_).f21
                nw45_[int(3)] = True
                nw45_[int(4)] = True
                nw45_[int(5)] = (d_202_v26_).f21
                nw45_[int(6)] = (d_202_v26_).f21
                nw45_[int(7)] = not(((d_218_v34_)[(d_202_v26_).f21] if ((d_202_v26_).f21) in (d_218_v34_) else (d_202_v26_).f21))
                nw45_[int(8)] = not((d_212_cf72_).f21)
                nw45_[int(9)] = (d_202_v26_).f21
                nw45_[int(10)] = d_169_v0_
                nw45_[int(11)] = False
                nw45_[int(12)] = (d_212_cf72_).f21
                nw45_[int(13)] = d_169_v0_
                nw45_[int(14)] = (d_202_v26_).f21
                nw45_[int(15)] = True
                nw45_[int(16)] = d_169_v0_
                nw45_[int(17)] = not((d_202_v26_).f21)
                nw45_[int(18)] = not((d_202_v26_).f21)
                d_219_v35_ = nw45_
                d_220_v36_: _dafny.MultiSet
                d_220_v36_ = _dafny.MultiSet([d_219_v35_, d_219_v35_, d_219_v35_])
                index22_ = default__.safeIndex(199, (d_213_v29_).length(0))
                (d_213_v29_)[index22_] = d_220_v36_
                index23_ = default__.safeIndex(199, (d_213_v29_).length(0))
                (d_213_v29_)[index23_] = d_220_v36_
                d_221_v37_: D3
                d_221_v37_ = D3_DC10(d_171_v2_)
                d_222_v38_: _dafny.MultiSet
                d_222_v38_ = _dafny.MultiSet([d_221_v37_])
                d_223_v39_: _dafny.MultiSet
                d_223_v39_ = _dafny.MultiSet([not(True), (d_202_v26_).f21, False, (d_222_v38_) != (d_222_v38_)])
                d_223_v39_ = d_223_v39_
                d_224_v40_: bool
                d_225_v41_: bool
                out1_: bool
                out2_: bool
                out1_, out2_ = (d_212_cf72_).m1((d_214_v31_).cf23, ((d_194_v21_)[(d_202_v26_).f21] if ((d_202_v26_).f21) in (d_194_v21_) else -498), False, (D1_DC4(812, (d_202_v26_).f21, d_219_v35_)).cf4, d_177_globalState_)
                d_224_v40_ = out1_
                d_225_v41_ = out2_
                d_226_v42_: _dafny.Array
                def lambda18_(d_227_cf72_):
                    def lambda19_(d_228_i5_):
                        return _dafny.MultiSet([(d_227_cf72_).f21])

                    return lambda19_

                init9_ = lambda18_(d_212_cf72_)
                nw46_ = _dafny.Array(None, 27)
                for i0_9_ in range(nw46_.length(0)):
                    nw46_[i0_9_] = init9_(i0_9_)
                d_226_v42_ = nw46_
                rhs28_ = d_225_v41_
                rhs29_ = _dafny.CodePoint('j')
                rhs30_ = d_226_v42_
                lhs17_ = d_177_globalState_
                lhs17_.f15 = rhs28_
                d_178_v8_ = rhs29_
                d_226_v42_ = rhs30_
            elif True:
                d_229___mcc_h3_ = source3_.cf71
                d_230_cf71_ = d_229___mcc_h3_
                d_231_v43_: T1
                nw47_ = C2()
                nw47_.ctor__(d_174_v5_, (d_202_v26_).f21)
                d_231_v43_ = nw47_
                d_231_v43_ = (d_231_v43_ if False else d_231_v43_)
                d_232_v44_: _dafny.Array
                nw48_ = _dafny.Array(False, 26)
                d_232_v44_ = nw48_
                d_232_v44_ = d_232_v44_
                d_233_v45_: bool
                d_234_v46_: bool
                d_235_v47_: bool
                d_236_v48_: _dafny.Seq
                out3_: bool
                out4_: bool
                out5_: bool
                out6_: _dafny.Seq
                out3_, out4_, out5_, out6_ = (d_202_v26_).m6(d_171_v2_, (_dafny.SeqWithoutIsStrInference([d_178_v8_ for d_237_i6_ in range(default__.abs(-465))])) + (d_175_v6_), ((d_202_v26_).f21 if (d_202_v26_).f21 else (d_202_v26_).f21), d_177_globalState_)
                d_233_v45_ = out3_
                d_234_v46_ = out4_
                d_235_v47_ = out5_
                d_236_v48_ = out6_
                index24_ = default__.safeIndex(272, (d_172_v3_).length(0))
                (d_172_v3_)[index24_] = d_171_v2_
                index25_ = default__.safeIndex(272, (d_172_v3_).length(0))
                (d_172_v3_)[index25_] = d_171_v2_
        d_238_v49_: D12
        d_238_v49_ = D12_DC35(D12_DC34())
        d_239_v50_: D12
        d_239_v50_ = D12_DC35(d_238_v49_)
        d_240_v51_: _dafny.Map
        d_240_v51_ = _dafny.Map({d_239_v50_: d_171_v2_})
        d_240_v51_ = (d_240_v51_).set(d_239_v50_, d_171_v2_)
        d_241_v52_: _dafny.MultiSet
        d_241_v52_ = _dafny.MultiSet([not(d_169_v0_), d_169_v0_])
        d_242_v53_: _dafny.Seq
        d_242_v53_ = _dafny.SeqWithoutIsStrInference([D10_DC27(d_241_v52_)])
        d_243_v54_: _dafny.Seq
        d_243_v54_ = _dafny.SeqWithoutIsStrInference([d_171_v2_])
        hi1_ = (d_243_v54_)[default__.safeIndex(425, len(d_243_v54_))]
        for d_244_i7_ in range((0) - (len(d_242_v53_)), hi1_):
            d_245_v55_: _dafny.Map
            d_245_v55_ = _dafny.Map({d_169_v0_: d_178_v8_})
            (d_177_globalState_).f11 = (default__.fm0(d_169_v0_, default__.fm12(len(d_245_v55_), d_175_v6_, d_175_v6_, d_177_globalState_), d_169_v0_, d_177_globalState_) if d_169_v0_ else d_169_v0_)
            d_246_v56_: bool
            d_247_v57_: bool
            d_248_v58_: int
            out7_: bool
            out8_: bool
            out9_: int
            out7_, out8_, out9_ = default__.m9(d_169_v0_, (False if d_169_v0_ else False), d_177_globalState_)
            d_246_v56_ = out7_
            d_247_v57_ = out8_
            d_248_v58_ = out9_
            d_178_v8_ = (d_175_v6_)[default__.safeIndex(d_248_v58_, len(d_175_v6_))]
            d_249_v59_: _dafny.Array
            nw49_ = _dafny.Array(False, 13)
            d_249_v59_ = nw49_
            d_250_v60_: _dafny.Seq
            d_250_v60_ = _dafny.SeqWithoutIsStrInference([d_249_v59_])
            if (d_250_v60_) < ((d_250_v60_) + (d_250_v60_)):
                d_251_v61_: _dafny.Map
                d_251_v61_ = _dafny.Map({d_178_v8_: default__.fm24(d_246_v56_, d_248_v58_, d_177_globalState_)})
                d_252_v62_: _dafny.Map
                d_252_v62_ = _dafny.Map({d_247_v57_: d_244_i7_})
                d_253_v63_: D3
                d_253_v63_ = D3_DC11(d_252_v62_, -703, d_172_v3_)
                d_254_v64_: _dafny.Array
                nw50_ = _dafny.Array(None, 4)
                nw50_[int(0)] = ((d_251_v61_)[d_178_v8_] if (d_178_v8_) in (d_251_v61_) else (d_253_v63_).cf18)
                nw50_[int(1)] = _dafny.Map({d_169_v0_: d_171_v2_})
                nw50_[int(2)] = d_252_v62_
                nw50_[int(3)] = d_252_v62_
                d_254_v64_ = nw50_
                index26_ = default__.safeIndex(843, (d_254_v64_).length(0))
                (d_254_v64_)[index26_] = d_252_v62_
                index27_ = default__.safeIndex(843, (d_254_v64_).length(0))
                (d_254_v64_)[index27_] = d_252_v62_
                d_255_v65_: C0
                nw51_ = C0()
                nw51_.ctor__(default__.fm0(d_169_v0_, d_178_v8_, d_247_v57_, d_177_globalState_))
                d_255_v65_ = nw51_
                d_256_v66_: _dafny.Map
                d_256_v66_ = _dafny.Map({d_241_v52_: d_175_v6_})
                d_257_v67_: bool
                d_258_v68_: bool
                d_259_v69_: int
                out10_: bool
                out11_: bool
                out12_: int
                out10_, out11_, out12_ = default__.m9((d_255_v65_).fm15(d_246_v56_, d_247_v57_, (0) - (d_248_v58_), d_252_v62_, d_177_globalState_), ((((d_256_v66_)[_dafny.MultiSet([not(d_247_v57_)])] if (_dafny.MultiSet([not(d_247_v57_)])) in (d_256_v66_) else d_175_v6_)).set(default__.safeIndex(d_244_i7_, len(((d_256_v66_)[_dafny.MultiSet([not(d_247_v57_)])] if (_dafny.MultiSet([not(d_247_v57_)])) in (d_256_v66_) else d_175_v6_))), d_178_v8_)) <= (d_175_v6_), d_177_globalState_)
                d_257_v67_ = out10_
                d_258_v68_ = out11_
                d_259_v69_ = out12_
                index28_ = default__.safeIndex(885, (d_172_v3_).length(0))
                (d_172_v3_)[index28_] = (d_248_v58_) * (d_248_v58_)
                index29_ = default__.safeIndex(885, (d_172_v3_).length(0))
                (d_172_v3_)[index29_] = d_171_v2_
                d_260_v70_: bool
                d_261_v71_: bool
                d_262_v72_: int
                out13_: bool
                out14_: bool
                out15_: int
                out13_, out14_, out15_ = default__.m9((d_255_v65_).f23, (d_175_v6_) <= (d_175_v6_), d_177_globalState_)
                d_260_v70_ = out13_
                d_261_v71_ = out14_
                d_262_v72_ = out15_
            elif True:
                d_171_v2_ = (d_244_i7_) + (default__.safeDivisionInt(873, d_244_i7_))
                d_169_v0_ = True
                d_263_v73_: _dafny.Map
                d_263_v73_ = _dafny.Map({(d_244_i7_) - ((0) - (d_171_v2_)): ((0) - (d_248_v58_)) < (d_248_v58_)})
                d_263_v73_ = d_263_v73_
                (d_177_globalState_).f15 = d_246_v56_
                d_264_v74_: bool
                d_265_v75_: bool
                d_266_v76_: int
                out16_: bool
                out17_: bool
                out18_: int
                out16_, out17_, out18_ = default__.m9(d_169_v0_, d_246_v56_, d_177_globalState_)
                d_264_v74_ = out16_
                d_265_v75_ = out17_
                d_266_v76_ = out18_
        d_267_v77_: C0
        nw52_ = C0()
        nw52_.ctor__(d_169_v0_)
        d_267_v77_ = nw52_
        d_268_v78_: _dafny.MultiSet
        d_268_v78_ = _dafny.MultiSet([d_267_v77_])
        rhs31_ = (d_169_v0_) == ((d_267_v77_).f23)
        rhs32_ = d_268_v78_
        lhs18_ = d_177_globalState_
        lhs18_.f11 = rhs31_
        d_268_v78_ = rhs32_
        d_269_v79_: _dafny.Map
        d_269_v79_ = _dafny.Map({(d_267_v77_).f23: d_171_v2_})
        if (d_267_v77_).fm15(d_169_v0_, d_169_v0_, d_171_v2_, d_269_v79_, d_177_globalState_):
            d_270_v80_: _dafny.Array
            def lambda20_(d_271_v8_):
                def lambda21_(d_272_i8_):
                    return d_271_v8_

                return lambda21_

            init10_ = lambda20_(d_178_v8_)
            nw53_ = _dafny.Array(None, 19)
            for i0_10_ in range(nw53_.length(0)):
                nw53_[i0_10_] = init10_(i0_10_)
            d_270_v80_ = nw53_
            index30_ = default__.safeIndex(593, (d_270_v80_).length(0))
            (d_270_v80_)[index30_] = d_178_v8_
            index31_ = default__.safeIndex(593, (d_270_v80_).length(0))
            (d_270_v80_)[index31_] = d_178_v8_
            d_273_v81_: C2
            nw54_ = C2()
            nw54_.ctor__(d_174_v5_, (d_267_v77_).f23)
            d_273_v81_ = nw54_
            d_274_v82_: bool
            d_275_v83_: bool
            d_276_v84_: bool
            d_277_v85_: _dafny.Seq
            out19_: bool
            out20_: bool
            out21_: bool
            out22_: _dafny.Seq
            out19_, out20_, out21_, out22_ = (d_273_v81_).m6(default__.safeDivisionInt(d_171_v2_, d_171_v2_), d_175_v6_, (d_267_v77_).f23, d_177_globalState_)
            d_274_v82_ = out19_
            d_275_v83_ = out20_
            d_276_v84_ = out21_
            d_277_v85_ = out22_
            d_278_v86_: _dafny.Map
            d_278_v86_ = _dafny.Map({_dafny.CodePoint('o'): (0) - (d_171_v2_)})
            if ((d_178_v8_) in (d_278_v86_)) or ((d_275_v83_) or (d_274_v82_)):
                d_279_v87_: D12
                d_279_v87_ = D12_DC34()
                d_280_v88_: _dafny.Map
                d_280_v88_ = _dafny.Map({d_279_v87_: d_267_v77_})
                d_280_v88_ = (d_280_v88_).set(d_279_v87_, d_267_v77_)
                d_281_v89_: _dafny.Set
                d_281_v89_ = _dafny.Set({d_243_v54_, _dafny.SeqWithoutIsStrInference([d_171_v2_ for d_282_i9_ in range(default__.abs(-160))])})
                d_283_v90_: _dafny.Seq
                d_283_v90_ = _dafny.SeqWithoutIsStrInference([d_281_v89_, d_281_v89_, d_281_v89_])
                (d_177_globalState_).f17 = (d_281_v89_).ispropersubset((d_283_v90_)[default__.safeIndex(d_171_v2_, len(d_283_v90_))])
                d_171_v2_ = len(_dafny.Map({(d_270_v80_)[default__.safeIndex(593, (d_270_v80_).length(0))]: d_241_v52_}))
                d_284_v91_: _dafny.Array
                nw55_ = _dafny.Array(_dafny.Array(None, 0), 4)
                d_284_v91_ = nw55_
                d_285_v92_: _dafny.MultiSet
                d_285_v92_ = _dafny.MultiSet([d_171_v2_, d_171_v2_])
                d_286_v93_: _dafny.Seq
                d_286_v93_ = _dafny.SeqWithoutIsStrInference([d_243_v54_])
                d_287_v94_: D8
                d_287_v94_ = D8_DC23(d_285_v92_)
                pat_let_tv2_ = d_285_v92_
                d_288_v95_: _dafny.Array
                nw56_ = _dafny.Array(None, 5)
                nw56_[int(0)] = d_285_v92_
                nw56_[int(1)] = _dafny.MultiSet((d_286_v93_)[default__.safeIndex(d_171_v2_, len(d_286_v93_))])
                def iife27_(_pat_let3_0):
                    def iife28_(d_289_dt__update__tmp_h0_):
                        def iife29_(_pat_let4_0):
                            def iife30_(d_290_dt__update_hcf41_h0_):
                                return D8_DC23(d_290_dt__update_hcf41_h0_)
                            return iife30_(_pat_let4_0)
                        return iife29_(pat_let_tv2_)
                    return iife28_(_pat_let3_0)
                nw56_[int(2)] = (iife27_(d_287_v94_)).cf41
                nw56_[int(3)] = d_285_v92_
                nw56_[int(4)] = d_285_v92_
                d_288_v95_ = nw56_
                index32_ = default__.safeIndex(556, (d_284_v91_).length(0))
                (d_284_v91_)[index32_] = d_288_v95_
                index33_ = default__.safeIndex(556, (d_284_v91_).length(0))
                (d_284_v91_)[index33_] = d_288_v95_
                (d_177_globalState_).f15 = d_274_v82_
            elif True:
                d_291_v96_: bool
                out23_: bool
                out23_ = (d_273_v81_).m5(not(not (True) or ((d_267_v77_).f23)), d_170_v1_, d_177_globalState_)
                d_291_v96_ = out23_
                d_292_v97_: D9
                d_292_v97_ = D9_DC25(d_171_v2_, (d_273_v81_).f21, d_171_v2_, d_267_v77_)
                d_293_v98_: _dafny.Map
                d_293_v98_ = _dafny.Map({d_171_v2_: d_292_v97_})
                d_294_v99_: C4
                nw57_ = C4()
                nw57_.ctor__((d_273_v81_).f21)
                d_294_v99_ = nw57_
                d_295_v100_: _dafny.Map
                d_295_v100_ = _dafny.Map({d_294_v99_: d_275_v83_})
                d_296_v101_: _dafny.Array
                nw58_ = _dafny.Array(False, 14)
                d_296_v101_ = nw58_
                d_297_v102_: _dafny.Map
                d_297_v102_ = _dafny.Map({d_296_v101_: (d_267_v77_).f23})
                d_298_v103_: bool
                d_299_v104_: bool
                d_300_v105_: bool
                d_301_v106_: _dafny.Seq
                out24_: bool
                out25_: bool
                out26_: bool
                out27_: _dafny.Seq
                out24_, out25_, out26_, out27_ = (d_273_v81_).m6((len(d_293_v98_)) + (len(d_295_v100_)), (d_277_v85_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yoiqoa"))), (D1_DC4(len(d_297_v102_), (d_267_v77_).fm15((d_273_v81_).f21, (d_273_v81_).f21, -511, d_269_v79_, d_177_globalState_), d_296_v101_)).cf5, d_177_globalState_)
                d_298_v103_ = out24_
                d_299_v104_ = out25_
                d_300_v105_ = out26_
                d_301_v106_ = out27_
                d_302_v107_: T1
                nw59_ = C2()
                nw59_.ctor__(d_273_v81_.f20, False)
                d_302_v107_ = nw59_
                d_303_v108_: C2
                nw60_ = C2()
                nw60_.ctor__(d_174_v5_, (d_302_v107_) in (_dafny.MultiSet([d_302_v107_, d_302_v107_])))
                d_303_v108_ = nw60_
                d_274_v82_ = default__.fm0(d_275_v83_, d_178_v8_, d_291_v96_, d_177_globalState_)
                d_304_v109_: C5
                nw61_ = C5()
                nw61_.ctor__(default__.fm5(len(_dafny.SeqWithoutIsStrInference([(d_273_v81_).f21])), d_177_globalState_))
                d_304_v109_ = nw61_
            d_305_v110_: _dafny.Set
            d_305_v110_ = _dafny.Set({d_275_v83_, d_275_v83_})
            d_306_v111_: _dafny.Seq
            d_306_v111_ = _dafny.SeqWithoutIsStrInference([d_305_v110_])
            d_307_v112_: bool
            d_308_v113_: bool
            d_309_v114_: bool
            d_310_v115_: _dafny.Seq
            out28_: bool
            out29_: bool
            out30_: bool
            out31_: _dafny.Seq
            out28_, out29_, out30_, out31_ = (d_273_v81_).m6(len(d_306_v111_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g') for d_311_i10_ in range(default__.abs(-938))]), (d_273_v81_).f21, d_177_globalState_)
            d_307_v112_ = out28_
            d_308_v113_ = out29_
            d_309_v114_ = out30_
            d_310_v115_ = out31_
        elif True:
            d_312_v116_: _dafny.Array
            nw62_ = _dafny.Array(_dafny.Seq({}), 23)
            d_312_v116_ = nw62_
            d_312_v116_ = d_312_v116_
            (d_177_globalState_).f17 = not((d_267_v77_).f23)
            d_313_v117_: _dafny.Map
            d_313_v117_ = _dafny.Map({d_171_v2_: d_269_v79_})
            d_313_v117_ = (d_313_v117_).set((d_171_v2_) + (d_171_v2_), (d_269_v79_) | (d_269_v79_))
            d_169_v0_ = (d_267_v77_).f23
            index34_ = default__.safeIndex(480, (d_172_v3_).length(0))
            (d_172_v3_)[index34_] = d_171_v2_
            index35_ = default__.safeIndex(480, (d_172_v3_).length(0))
            (d_172_v3_)[index35_] = d_171_v2_
        d_314_v118_: bool
        d_315_v119_: bool
        d_316_v120_: int
        out32_: bool
        out33_: bool
        out34_: int
        out32_, out33_, out34_ = default__.m9(d_169_v0_, d_169_v0_, d_177_globalState_)
        d_314_v118_ = out32_
        d_315_v119_ = out33_
        d_316_v120_ = out34_
        d_317_v121_: _dafny.Map
        d_317_v121_ = _dafny.Map({True: d_314_v118_})
        d_318_v122_: D13
        d_318_v122_ = D13_DC36((_dafny.SeqWithoutIsStrInference([d_317_v121_, d_317_v121_, d_317_v121_, d_317_v121_, d_317_v121_])).set(default__.safeIndex(d_171_v2_, len(_dafny.SeqWithoutIsStrInference([d_317_v121_, d_317_v121_, d_317_v121_, d_317_v121_, d_317_v121_]))), d_317_v121_))
        pat_let_tv3_ = d_178_v8_
        pat_let_tv4_ = d_178_v8_
        pat_let_tv5_ = d_267_v77_
        pat_let_tv6_ = d_316_v120_
        pat_let_tv7_ = d_316_v120_
        pat_let_tv8_ = d_178_v8_
        pat_let_tv9_ = d_178_v8_
        def lambda22_(source4_):
            if source4_.is_DC37:
                d_319___mcc_h4_ = source4_.cf64
                d_320___mcc_h5_ = source4_.cf65
                d_321_cf65_ = d_320___mcc_h5_
                d_322_cf64_ = d_319___mcc_h4_
                return pat_let_tv3_
            elif source4_.is_DC38:
                d_323___mcc_h6_ = source4_.cf66
                d_324___mcc_h7_ = source4_.cf67
                d_325___mcc_h8_ = source4_.cf68
                d_326___mcc_h9_ = source4_.cf69
                d_327_cf69_ = d_326___mcc_h9_
                d_328_cf68_ = d_325___mcc_h8_
                d_329_cf67_ = d_324___mcc_h7_
                d_330_cf66_ = d_323___mcc_h6_
                return pat_let_tv4_
            elif source4_.is_DC36:
                d_331___mcc_h10_ = source4_.cf63
                d_332_cf63_ = d_331___mcc_h10_
                return (D4_DC14((pat_let_tv5_).f23, pat_let_tv6_, _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdb"))), pat_let_tv7_}), pat_let_tv8_)).cf26
            elif True:
                d_333___mcc_h11_ = source4_.cf70
                d_334_cf70_ = d_333___mcc_h11_
                return pat_let_tv9_

        d_178_v8_ = lambda22_(d_318_v122_)
        hi2_ = d_171_v2_
        for d_335_i11_ in range(d_316_v120_, hi2_):
            d_336_v123_: _dafny.Array
            nw63_ = _dafny.Array(None, 16)
            d_336_v123_ = nw63_
            d_337_v124_: C2
            nw64_ = C2()
            nw64_.ctor__(d_174_v5_, True)
            d_337_v124_ = nw64_
            d_338_v125_: _dafny.Map
            d_338_v125_ = _dafny.Map({d_335_i11_: d_337_v124_})
            index36_ = default__.safeIndex(581, (d_336_v123_).length(0))
            (d_336_v123_)[index36_] = ((d_338_v125_)[d_171_v2_] if (d_171_v2_) in (d_338_v125_) else d_337_v124_)
            d_339_v126_: _dafny.Seq
            d_339_v126_ = _dafny.SeqWithoutIsStrInference([d_337_v124_, d_337_v124_, d_337_v124_, d_337_v124_])
            index37_ = default__.safeIndex(581, (d_336_v123_).length(0))
            (d_336_v123_)[index37_] = (d_339_v126_)[default__.safeIndex(d_335_i11_, len(d_339_v126_))]
            d_340_v127_: _dafny.Seq
            d_340_v127_ = _dafny.SeqWithoutIsStrInference([d_318_v122_])
            d_341_v128_: _dafny.Seq
            d_341_v128_ = _dafny.SeqWithoutIsStrInference([d_340_v127_])
            if ((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_318_v122_ for d_342_i13_ in range(default__.abs(856))]) for d_343_i12_ in range(default__.abs(112))])) + (d_341_v128_)) <= (d_341_v128_):
                d_169_v0_ = (True) or (d_314_v118_)
                d_344_v129_: _dafny.Set
                d_344_v129_ = _dafny.Set({d_171_v2_})
                d_344_v129_ = (_dafny.Set({len(d_175_v6_)})).intersection(d_344_v129_)
                d_316_v120_ = 371
                d_345_v130_: bool
                d_346_v131_: bool
                d_347_v132_: bool
                d_348_v133_: _dafny.Seq
                out35_: bool
                out36_: bool
                out37_: bool
                out38_: _dafny.Seq
                out35_, out36_, out37_, out38_ = (d_337_v124_).m6(default__.safeDivisionInt(len(d_175_v6_), len(d_269_v79_)), (d_175_v6_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cqmm"))), (d_267_v77_).f23, d_177_globalState_)
                d_345_v130_ = out35_
                d_346_v131_ = out36_
                d_347_v132_ = out37_
                d_348_v133_ = out38_
                d_349_v134_: D9
                d_349_v134_ = D9_DC24(_dafny.SeqWithoutIsStrInference([d_316_v120_ for d_350_i14_ in range(default__.abs(-737))]))
                d_243_v54_ = (d_349_v134_).cf42
            elif True:
                d_351_v135_: bool
                d_352_v136_: bool
                d_353_v137_: int
                out39_: bool
                out40_: bool
                out41_: int
                out39_, out40_, out41_ = default__.m9((d_337_v124_).f21, (d_267_v77_).f23, d_177_globalState_)
                d_351_v135_ = out39_
                d_352_v136_ = out40_
                d_353_v137_ = out41_
                d_175_v6_ = (d_175_v6_).set(default__.safeIndex((0) - (d_353_v137_), len(d_175_v6_)), d_178_v8_)
                d_354_v138_: _dafny.Set
                d_354_v138_ = _dafny.Set({d_175_v6_, _dafny.SeqWithoutIsStrInference([d_178_v8_ for d_355_i15_ in range(default__.abs(-434))])})
                d_316_v120_ = (d_335_i11_ if (d_267_v77_).f23 else default__.safeModuloInt(len(d_354_v138_), d_171_v2_))
                (d_177_globalState_).f15 = ((d_175_v6_ if d_351_v135_ else d_175_v6_)) <= ((d_175_v6_) + (d_175_v6_))
                d_356_v139_: C0
                nw65_ = C0()
                nw65_.ctor__(True)
                d_356_v139_ = nw65_
            d_357_v140_: T1
            nw66_ = C4()
            nw66_.ctor__(d_314_v118_)
            d_357_v140_ = nw66_
            d_358_v141_: _dafny.Set
            d_358_v141_ = _dafny.Set({d_357_v140_, d_357_v140_, d_357_v140_})
            d_359_v142_: D12
            d_359_v142_ = D12_DC31(d_358_v141_)
            pat_let_tv10_ = d_358_v141_
            def iife31_(_pat_let5_0):
                def iife32_(d_360_dt__update__tmp_h1_):
                    def iife33_(_pat_let6_0):
                        def iife34_(d_361_dt__update_hcf55_h0_):
                            return D12_DC31(d_361_dt__update_hcf55_h0_)
                        return iife34_(_pat_let6_0)
                    return iife33_(pat_let_tv10_)
                return iife32_(_pat_let5_0)
            source5_ = iife31_(d_359_v142_)
            if source5_.is_DC32:
                d_362___mcc_h12_ = source5_.cf56
                d_363_cf56_ = d_362___mcc_h12_
                d_364_v143_: _dafny.Seq
                d_364_v143_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_337_v124_).f21: d_314_v118_}), d_317_v121_, d_317_v121_])
                d_365_v144_: D13
                d_365_v144_ = D13_DC36(d_364_v143_)
                d_366_v145_: D13
                d_366_v145_ = D13_DC39(d_365_v144_)
                d_366_v145_ = d_366_v145_
                d_367_v146_: D13
                d_367_v146_ = D13_DC38(d_178_v8_, (d_267_v77_).f23, d_335_i11_, d_172_v3_)
                d_368_v147_: bool
                d_369_v148_: bool
                d_370_v149_: int
                out42_: bool
                out43_: bool
                out44_: int
                out42_, out43_, out44_ = default__.m9((d_169_v0_ if (d_267_v77_).f23 else (d_337_v124_).f21), (d_367_v146_).cf67, d_177_globalState_)
                d_368_v147_ = out42_
                d_369_v148_ = out43_
                d_370_v149_ = out44_
                d_371_v150_: _dafny.Set
                d_371_v150_ = _dafny.Set({d_175_v6_})
                d_314_v118_ = (d_371_v150_).issubset(d_371_v150_)
                d_372_v151_: C1
                nw67_ = C1()
                nw67_.ctor__(d_172_v3_)
                d_372_v151_ = nw67_
            elif source5_.is_DC33:
                d_373___mcc_h13_ = source5_.cf57
                d_374___mcc_h14_ = source5_.cf58
                d_375___mcc_h15_ = source5_.cf59
                d_376___mcc_h16_ = source5_.cf60
                d_377___mcc_h17_ = source5_.cf61
                d_378_cf61_ = d_377___mcc_h17_
                d_379_cf60_ = d_376___mcc_h16_
                d_380_cf59_ = d_375___mcc_h15_
                d_381_cf58_ = d_374___mcc_h14_
                d_382_cf57_ = d_373___mcc_h13_
                d_383_v152_: _dafny.Map
                d_383_v152_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_178_v8_ for d_384_i16_ in range(default__.abs(727))]): False})
                d_382_cf57_ = default__.safeDivisionInt((0) - (default__.fm3(-604, d_316_v120_, not(((d_383_v152_)[d_175_v6_] if (d_175_v6_) in (d_383_v152_) else False)), d_177_globalState_)), d_171_v2_)
                d_385_v153_: C4
                nw68_ = C4()
                nw68_.ctor__((d_337_v124_).f21)
                d_385_v153_ = nw68_
                d_357_v140_ = d_357_v140_
                (d_177_globalState_).f11 = (d_267_v77_).f23
            elif source5_.is_DC34:
                d_386_v154_: _dafny.Map
                d_386_v154_ = _dafny.Map({d_171_v2_: d_169_v0_})
                d_386_v154_ = (d_386_v154_).set(918, ((_dafny.MultiSet(d_170_v1_)).cardinality) < (d_335_i11_))
                d_387_v155_: D13
                d_387_v155_ = D13_DC37(-142, _dafny.SeqWithoutIsStrInference([d_243_v54_]))
                d_388_v156_: _dafny.Set
                d_388_v156_ = _dafny.Set({d_171_v2_, d_171_v2_, 368})
                d_389_v157_: D13
                d_389_v157_ = D13_DC38(d_178_v8_, d_315_v119_, d_335_i11_, d_172_v3_)
                d_390_v158_: _dafny.Array
                nw69_ = _dafny.Array(None, 26)
                nw69_[int(0)] = d_178_v8_
                nw69_[int(1)] = ((d_175_v6_)[default__.safeIndex(len(d_388_v156_), len(d_175_v6_))] if not(True) else d_178_v8_)
                nw69_[int(2)] = d_178_v8_
                nw69_[int(3)] = d_178_v8_
                nw69_[int(4)] = d_178_v8_
                nw69_[int(5)] = (d_389_v157_).cf66
                nw69_[int(6)] = d_178_v8_
                nw69_[int(7)] = d_178_v8_
                nw69_[int(8)] = d_178_v8_
                nw69_[int(9)] = d_178_v8_
                nw69_[int(10)] = d_178_v8_
                nw69_[int(11)] = d_178_v8_
                nw69_[int(12)] = d_178_v8_
                nw69_[int(13)] = d_178_v8_
                nw69_[int(14)] = (d_175_v6_)[default__.safeIndex(d_335_i11_, len(d_175_v6_))]
                nw69_[int(15)] = d_178_v8_
                nw69_[int(16)] = d_178_v8_
                nw69_[int(17)] = d_178_v8_
                nw69_[int(18)] = d_178_v8_
                nw69_[int(19)] = d_178_v8_
                nw69_[int(20)] = default__.fm12(d_316_v120_, d_175_v6_, d_175_v6_, d_177_globalState_)
                nw69_[int(21)] = d_178_v8_
                nw69_[int(22)] = d_178_v8_
                nw69_[int(23)] = d_178_v8_
                nw69_[int(24)] = d_178_v8_
                nw69_[int(25)] = default__.fm12(len(d_170_v1_), d_175_v6_, _dafny.SeqWithoutIsStrInference([d_178_v8_ for d_391_i17_ in range(default__.abs(454))]), d_177_globalState_)
                d_390_v158_ = nw69_
                index38_ = default__.safeIndex(147, (d_390_v158_).length(0))
                (d_390_v158_)[index38_] = d_178_v8_
                index39_ = default__.safeIndex(147, (d_390_v158_).length(0))
                rhs33_ = d_387_v155_
                rhs34_ = d_388_v156_
                rhs35_ = d_316_v120_
                rhs36_ = d_171_v2_
                rhs37_ = d_178_v8_
                lhs19_ = d_390_v158_
                lhs20_ = default__.safeIndex(147, (d_390_v158_).length(0))
                d_387_v155_ = rhs33_
                d_388_v156_ = rhs34_
                d_171_v2_ = rhs35_
                d_316_v120_ = rhs36_
                lhs19_[lhs20_] = rhs37_
                index40_ = default__.safeIndex(463, (d_172_v3_).length(0))
                (d_172_v3_)[index40_] = d_171_v2_
                d_392_v159_: _dafny.MultiSet
                d_392_v159_ = _dafny.MultiSet([default__.safeDivisionInt(d_335_i11_, len(d_386_v154_))])
                index41_ = default__.safeIndex(463, (d_172_v3_).length(0))
                (d_172_v3_)[index41_] = (d_392_v159_).cardinality
                d_393_v160_: _dafny.Map
                d_393_v160_ = _dafny.Map({True: d_170_v1_})
                d_394_v161_: _dafny.Seq
                d_394_v161_ = _dafny.SeqWithoutIsStrInference([((d_393_v160_)[True] if (True) in (d_393_v160_) else d_170_v1_)])
                index42_ = default__.safeIndex(463, (d_172_v3_).length(0))
                (d_172_v3_)[index42_] = (d_171_v2_) + (len((d_394_v161_)[default__.safeIndex(default__.fm5(d_316_v120_, d_177_globalState_), len(d_394_v161_))]))
            elif source5_.is_DC31:
                d_395___mcc_h18_ = source5_.cf55
                d_396_cf55_ = d_395___mcc_h18_
                d_397_v162_: bool
                d_398_v163_: bool
                out45_: bool
                out46_: bool
                out45_, out46_ = (d_357_v140_).m1((d_337_v124_).f21, len(d_175_v6_), not(d_314_v118_), default__.safeModuloInt(((d_269_v79_)[d_314_v118_] if (d_314_v118_) in (d_269_v79_) else -122), -560), d_177_globalState_)
                d_397_v162_ = out45_
                d_398_v163_ = out46_
                rhs38_ = (d_335_i11_) <= ((d_335_i11_) - (len(default__.fm19(d_243_v54_, d_397_v162_, d_177_globalState_))))
                rhs39_ = (-772 if (d_337_v124_).f21 else (0) - (d_171_v2_))
                rhs40_ = d_315_v119_
                rhs41_ = (d_337_v124_).f21
                lhs21_ = d_177_globalState_
                lhs22_ = d_177_globalState_
                d_398_v163_ = rhs38_
                d_171_v2_ = rhs39_
                lhs21_.f15 = rhs40_
                lhs22_.f15 = rhs41_
                d_399_v164_: D3
                d_399_v164_ = D3_DC11(d_269_v79_, (d_241_v52_).cardinality, d_172_v3_)
                rhs42_ = ((d_399_v164_).cf18) | (d_269_v79_)
                rhs43_ = False
                d_269_v79_ = rhs42_
                d_398_v163_ = rhs43_
                d_175_v6_ = ((d_175_v6_ if d_398_v163_ else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "esrnngllh")))) + (d_175_v6_)
            elif True:
                d_400___mcc_h19_ = source5_.cf62
                d_401_cf62_ = d_400___mcc_h19_
                d_171_v2_ = d_316_v120_
                d_239_v50_ = d_239_v50_
                (d_177_globalState_).f1 = _dafny.SeqWithoutIsStrInference([(d_337_v124_).f21])
                d_402_v165_: _dafny.Set
                d_402_v165_ = _dafny.Set({d_172_v3_})
                d_403_v166_: D15
                d_403_v166_ = D15_DC42(d_357_v140_)
                d_404_v167_: _dafny.Seq
                d_404_v167_ = _dafny.SeqWithoutIsStrInference([d_357_v140_, d_357_v140_])
                d_405_v168_: _dafny.Array
                nw70_ = _dafny.Array(None, 27)
                nw70_[int(0)] = d_357_v140_
                nw70_[int(1)] = (D15_DC42(d_357_v140_)).cf75
                nw70_[int(2)] = d_357_v140_
                nw70_[int(3)] = (d_403_v166_).cf75
                nw70_[int(4)] = d_357_v140_
                nw70_[int(5)] = d_357_v140_
                nw70_[int(6)] = d_357_v140_
                nw70_[int(7)] = d_357_v140_
                nw70_[int(8)] = d_357_v140_
                nw70_[int(9)] = d_357_v140_
                nw70_[int(10)] = d_357_v140_
                nw70_[int(11)] = d_357_v140_
                nw70_[int(12)] = d_357_v140_
                nw70_[int(13)] = d_357_v140_
                nw70_[int(14)] = d_357_v140_
                nw70_[int(15)] = d_357_v140_
                nw70_[int(16)] = d_357_v140_
                nw70_[int(17)] = d_357_v140_
                nw70_[int(18)] = d_357_v140_
                nw70_[int(19)] = d_357_v140_
                nw70_[int(20)] = d_357_v140_
                nw70_[int(21)] = d_357_v140_
                nw70_[int(22)] = d_357_v140_
                nw70_[int(23)] = d_357_v140_
                nw70_[int(24)] = (d_404_v167_)[default__.safeIndex(len(default__.fm19(d_243_v54_, (d_267_v77_).f23, d_177_globalState_)), len(d_404_v167_))]
                nw70_[int(25)] = d_357_v140_
                nw70_[int(26)] = d_357_v140_
                d_405_v168_ = nw70_
                index43_ = default__.safeIndex(995, (d_405_v168_).length(0))
                (d_405_v168_)[index43_] = d_357_v140_
                d_406_v169_: _dafny.Seq
                d_406_v169_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tetmmpkc")), d_175_v6_, d_175_v6_])
                index44_ = default__.safeIndex(995, (d_405_v168_).length(0))
                rhs44_ = d_402_v165_
                rhs45_ = d_357_v140_
                rhs46_ = (d_406_v169_)[default__.safeIndex(d_316_v120_, len(d_406_v169_))]
                lhs23_ = d_405_v168_
                lhs24_ = default__.safeIndex(995, (d_405_v168_).length(0))
                d_402_v165_ = rhs44_
                lhs23_[lhs24_] = rhs45_
                d_175_v6_ = rhs46_
            d_171_v2_ = d_316_v120_
        d_407_v170_: _dafny.Array
        def lambda23_(d_408_v1_):
            def lambda24_(d_409_i19_):
                return (d_408_v1_) <= (d_408_v1_)

            return lambda24_

        init11_ = lambda23_(d_170_v1_)
        nw71_ = _dafny.Array(None, 28)
        for i0_11_ in range(nw71_.length(0)):
            nw71_[i0_11_] = init11_(i0_11_)
        d_407_v170_ = nw71_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_407_v170_).length(0)):
            d_410_i18_: int = guard_loop_0_
            if (True) and (((0) <= (d_410_i18_)) and ((d_410_i18_) < ((d_407_v170_).length(0)))):
                (d_407_v170_)[(d_410_i18_)] = False
        d_411_v171_: _dafny.Map
        d_411_v171_ = _dafny.Map({d_171_v2_: d_169_v0_})
        hi3_ = (((d_269_v79_)[(d_267_v77_).f23] if ((d_267_v77_).f23) in (d_269_v79_) else d_171_v2_)) - (d_316_v120_)
        for d_412_i20_ in range(((d_267_v77_).fm14(len(_dafny.Map({d_411_v171_: d_316_v120_})), d_177_globalState_)) * ((0) - (d_316_v120_)), hi3_):
            d_316_v120_ = (0) - (d_316_v120_)
            d_413_v172_: D10
            d_413_v172_ = D10_DC27(d_241_v52_)
            source6_ = d_413_v172_
            if source6_.is_DC28:
                d_414___mcc_h20_ = source6_.cf49
                d_415___mcc_h21_ = source6_.cf50
                d_416_cf50_ = d_415___mcc_h21_
                d_417_cf49_ = d_414___mcc_h20_
                d_418_v173_: _dafny.Seq
                d_418_v173_ = _dafny.SeqWithoutIsStrInference([d_243_v54_])
                d_419_v174_: D13
                d_419_v174_ = D13_DC39(D13_DC37(d_171_v2_, d_418_v173_))
                d_420_v175_: _dafny.Map
                d_420_v175_ = _dafny.Map({_dafny.CodePoint('q'): d_419_v174_})
                d_420_v175_ = d_420_v175_
                d_421_v176_: _dafny.Set
                d_421_v176_ = _dafny.Set({d_172_v3_, d_172_v3_})
                d_422_v177_: _dafny.Map
                d_422_v177_ = _dafny.Map({d_316_v120_: d_421_v176_})
                d_423_v178_: T0
                nw72_ = C2()
                nw72_.ctor__(d_174_v5_, (d_267_v77_).f23)
                d_423_v178_ = nw72_
                d_424_v179_: _dafny.Map
                d_424_v179_ = _dafny.Map({d_407_v170_: d_423_v178_})
                d_169_v0_ = (d_267_v77_).fm15((d_421_v176_).issubset(((d_422_v177_)[d_171_v2_] if (d_171_v2_) in (d_422_v177_) else d_421_v176_)), d_315_v119_, default__.safeDivisionInt(len(d_424_v179_), d_316_v120_), d_269_v79_, d_177_globalState_)
                d_315_v119_ = d_417_cf49_
                d_425_v180_: C5
                nw73_ = C5()
                nw73_.ctor__(d_412_i20_)
                d_425_v180_ = nw73_
            elif True:
                d_426___mcc_h22_ = source6_.cf48
                d_427_cf48_ = d_426___mcc_h22_
                index45_ = default__.safeIndex(421, (d_407_v170_).length(0))
                (d_407_v170_)[index45_] = True
                index46_ = default__.safeIndex(421, (d_407_v170_).length(0))
                (d_407_v170_)[index46_] = d_315_v119_
                d_176_v7_ = (d_176_v7_).set(d_316_v120_, len(d_175_v6_))
                d_316_v120_ = d_171_v2_
                d_428_v181_: _dafny.Array
                def lambda25_(d_429_v0_):
                    def lambda26_(d_430_i21_):
                        return d_429_v0_

                    return lambda26_

                init12_ = lambda25_(d_169_v0_)
                nw74_ = _dafny.Array(None, 2)
                for i0_12_ in range(nw74_.length(0)):
                    nw74_[i0_12_] = init12_(i0_12_)
                d_428_v181_ = nw74_
                d_431_v182_: D1
                d_431_v182_ = D1_DC4((0) - (d_316_v120_), (d_267_v77_).f23, d_428_v181_)
                d_432_v183_: _dafny.Seq
                d_432_v183_ = _dafny.SeqWithoutIsStrInference([d_170_v1_, d_170_v1_, _dafny.SeqWithoutIsStrInference([False, d_315_v119_])])
                rhs47_ = (d_267_v77_).f23
                rhs48_ = (d_316_v120_) == ((d_431_v182_).cf4)
                rhs49_ = d_315_v119_
                rhs50_ = (default__.safeDivisionInt(d_171_v2_, d_316_v120_)) - (default__.safeModuloInt(d_316_v120_, len((d_432_v183_)[default__.safeIndex(d_412_i20_, len(d_432_v183_))])))
                lhs25_ = d_177_globalState_
                lhs26_ = d_177_globalState_
                lhs25_.f11 = rhs47_
                d_315_v119_ = rhs48_
                lhs26_.f11 = rhs49_
                d_171_v2_ = rhs50_
            d_433_v184_: C1
            nw75_ = C1()
            nw75_.ctor__(d_172_v3_)
            d_433_v184_ = nw75_
            index47_ = default__.safeIndex(974, ((d_433_v184_).f22).length(0))
            ((d_433_v184_).f22)[index47_] = d_316_v120_
            index48_ = default__.safeIndex(974, ((d_433_v184_).f22).length(0))
            ((d_433_v184_).f22)[index48_] = d_316_v120_
        hi4_ = d_171_v2_
        for d_434_i22_ in range(d_316_v120_, hi4_):
            d_435_v185_: bool
            d_436_v186_: bool
            d_437_v187_: int
            out47_: bool
            out48_: bool
            out49_: int
            out47_, out48_, out49_ = default__.m9((d_267_v77_).f23, False, d_177_globalState_)
            d_435_v185_ = out47_
            d_436_v186_ = out48_
            d_437_v187_ = out49_
            d_438_v188_: C4
            nw76_ = C4()
            nw76_.ctor__(((d_267_v77_).fm15(False, (d_267_v77_).f23, d_437_v187_, d_269_v79_, d_177_globalState_)) or (d_435_v185_))
            d_438_v188_ = nw76_
            d_439_v189_: C2
            nw77_ = C2()
            nw77_.ctor__(d_174_v5_, (d_267_v77_).f23)
            d_439_v189_ = nw77_
            d_176_v7_ = (d_176_v7_).set(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmdrnblv"))), d_437_v187_), d_171_v2_)
        d_440_v190_: bool
        d_441_v191_: bool
        d_442_v192_: int
        out50_: bool
        out51_: bool
        out52_: int
        out50_, out51_, out52_ = default__.m9(d_169_v0_, True, d_177_globalState_)
        d_440_v190_ = out50_
        d_441_v191_ = out51_
        d_442_v192_ = out52_
        d_317_v121_ = (d_317_v121_).set(d_169_v0_, not(d_315_v119_))
        d_443_v193_: C4
        nw78_ = C4()
        nw78_.ctor__(d_315_v119_)
        d_443_v193_ = nw78_
        d_444_v194_: _dafny.MultiSet
        d_444_v194_ = _dafny.MultiSet([d_407_v170_])
        d_445_v195_: _dafny.MultiSet
        d_445_v195_ = _dafny.MultiSet([d_171_v2_])
        rhs51_ = (d_175_v6_).set(default__.safeIndex((d_442_v192_) - (len((d_243_v54_).set(default__.safeIndex(d_442_v192_, len(d_243_v54_)), ((d_444_v194_)[d_407_v170_] if (d_407_v170_) in (d_444_v194_) else d_171_v2_)))), len(d_175_v6_)), _dafny.CodePoint('w'))
        rhs52_ = ((d_445_v195_) | (d_445_v195_)).cardinality
        rhs53_ = d_443_v193_
        d_175_v6_ = rhs51_
        d_442_v192_ = rhs52_
        d_443_v193_ = rhs53_
        (d_177_globalState_).f1 = _dafny.SeqWithoutIsStrInference([True])
        d_446_v196_: C1
        nw79_ = C1()
        nw79_.ctor__(d_172_v3_)
        d_446_v196_ = nw79_
        _dafny.print(_dafny.string_of(d_169_v0_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_170_v1_) == (_dafny.SeqWithoutIsStrInference([True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_171_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_174_v5_).cf1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_175_v6_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c'), _dafny.CodePoint('c')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_176_v7_) == (_dafny.Map({644: 644}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_.f1) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_177_globalState_).f4) == (_dafny.SeqWithoutIsStrInference([True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_177_globalState_).f8).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_177_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_177_globalState_).f12) == (_dafny.Map({644: 644}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_).f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_177_globalState_.f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_177_globalState_.f16) == (_dafny.SeqWithoutIsStrInference([True, True, True, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_177_globalState_.f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_178_v8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_240_v51_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_241_v52_) == (_dafny.MultiSet([False, True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_242_v53_) == (_dafny.SeqWithoutIsStrInference([D10_DC27(_dafny.MultiSet([False, True]))]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_243_v54_) == (_dafny.SeqWithoutIsStrInference([-916]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_267_v77_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_268_v78_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_269_v79_) == (_dafny.Map({True: -916}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_314_v118_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_315_v119_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_316_v120_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_317_v121_) == (_dafny.Map({True: False}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_318_v122_).cf63) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({True: True}), _dafny.Map({True: True}), _dafny.Map({True: True}), _dafny.Map({True: True}), _dafny.Map({True: True})]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[24]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[25]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[26]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_407_v170_)[27]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_411_v171_) == (_dafny.Map({1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_440_v190_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_441_v191_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_442_v192_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_443_v193_).f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_444_v194_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_445_v195_) == (_dafny.MultiSet([1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(False)
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

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2
    def __hash__(self) -> int:
        return super().__hash__()


class D1:
    @classmethod
    def default(cls, ):
        return lambda: D1_DC4(int(0), False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D1_DC5)
    @property
    def is_DC6(self) -> bool:
        return isinstance(self, D1_DC6)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf4', Any), ('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({self.cf7.VerbatimString(True)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC6(D1, NamedTuple('DC6', [('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC6({_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC6) and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC8(_dafny.Array(None, 0), int(0), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)

class D2_DC8(D2, NamedTuple('DC8', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {_dafny.string_of(self.cf15)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: D3_DC10(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC10(self) -> bool:
        return isinstance(self, D3_DC10)
    @property
    def is_DC11(self) -> bool:
        return isinstance(self, D3_DC11)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC10(D3, NamedTuple('DC10', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC10({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC10) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC11(D3, NamedTuple('DC11', [('cf18', Any), ('cf19', Any), ('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC11({_dafny.string_of(self.cf18)}, {_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC11) and self.cf18 == __o.cf18 and self.cf19 == __o.cf19 and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D3_DC9(D3, NamedTuple('DC9', [('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf16 == __o.cf16
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC13(int(0))
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
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D4_DC15)

class D4_DC13(D4, NamedTuple('DC13', [('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC14(D4, NamedTuple('DC14', [('cf23', Any), ('cf24', Any), ('cf25', Any), ('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC14({_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)}, {_dafny.string_of(self.cf25)}, {_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC14) and self.cf23 == __o.cf23 and self.cf24 == __o.cf24 and self.cf25 == __o.cf25 and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC15(D4, NamedTuple('DC15', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC15({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC15) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC17(False, D2.default()(), False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D5_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D5_DC16)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D5_DC18)

class D5_DC17(D5, NamedTuple('DC17', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any), ('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC17({_dafny.string_of(self.cf29)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)}, {_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC17) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32 and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC16(D5, NamedTuple('DC16', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC16({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC16) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC18(D5, NamedTuple('DC18', [('cf34', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC18({_dafny.string_of(self.cf34)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC18) and self.cf34 == __o.cf34
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC20(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D6_DC20)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D6_DC19)

class D6_DC20(D6, NamedTuple('DC20', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC20({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC20) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC19(D6, NamedTuple('DC19', [('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC19({_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC19) and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D7_DC21)

class D7_DC21(D7, NamedTuple('DC21', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC21({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC21) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC23(_dafny.MultiSet({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D8_DC23)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D8_DC22)

class D8_DC23(D8, NamedTuple('DC23', [('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC23({_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC23) and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf40', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf40)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf40 == __o.cf40
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC25(int(0), False, int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D9_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D9_DC24)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D9_DC26)

class D9_DC25(D9, NamedTuple('DC25', [('cf43', Any), ('cf44', Any), ('cf45', Any), ('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC25({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {_dafny.string_of(self.cf45)}, {_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC25) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45 and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC24(D9, NamedTuple('DC24', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC24({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC24) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC26(D9, NamedTuple('DC26', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC26({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC26) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC28(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D10_DC28)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D10_DC27)

class D10_DC28(D10, NamedTuple('DC28', [('cf49', Any), ('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC28({_dafny.string_of(self.cf49)}, {_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC28) and self.cf49 == __o.cf49 and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC27(D10, NamedTuple('DC27', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC27({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC27) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC30(False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D11_DC30)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D11_DC29)

class D11_DC30(D11, NamedTuple('DC30', [('cf52', Any), ('cf53', Any), ('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC30({_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC30) and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC29(D11, NamedTuple('DC29', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC29({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC29) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC32(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D12_DC32)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D12_DC33)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D12_DC34)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D12_DC31)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D12_DC35)

class D12_DC32(D12, NamedTuple('DC32', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC32({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC32) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC33(D12, NamedTuple('DC33', [('cf57', Any), ('cf58', Any), ('cf59', Any), ('cf60', Any), ('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC33({_dafny.string_of(self.cf57)}, {_dafny.string_of(self.cf58)}, {_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC33) and self.cf57 == __o.cf57 and self.cf58 == __o.cf58 and self.cf59 == __o.cf59 and self.cf60 == __o.cf60 and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC34(D12, NamedTuple('DC34', [])):
    def __dafnystr__(self) -> str:
        return f'D12.DC34'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC34)
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC31(D12, NamedTuple('DC31', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC31({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC31) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC35(D12, NamedTuple('DC35', [('cf62', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC35({_dafny.string_of(self.cf62)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC35) and self.cf62 == __o.cf62
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC37(int(0), _dafny.Seq({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D13_DC37)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D13_DC38)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D13_DC36)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D13_DC39)

class D13_DC37(D13, NamedTuple('DC37', [('cf64', Any), ('cf65', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC37({_dafny.string_of(self.cf64)}, {_dafny.string_of(self.cf65)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC37) and self.cf64 == __o.cf64 and self.cf65 == __o.cf65
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC38(D13, NamedTuple('DC38', [('cf66', Any), ('cf67', Any), ('cf68', Any), ('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC38({_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)}, {_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC38) and self.cf66 == __o.cf66 and self.cf67 == __o.cf67 and self.cf68 == __o.cf68 and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC36(D13, NamedTuple('DC36', [('cf63', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC36({_dafny.string_of(self.cf63)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC36) and self.cf63 == __o.cf63
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC39(D13, NamedTuple('DC39', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC39({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC39) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: D14_DC41(None, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D14_DC41)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D14_DC40)

class D14_DC41(D14, NamedTuple('DC41', [('cf72', Any), ('cf73', Any), ('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC41({_dafny.string_of(self.cf72)}, {self.cf73.VerbatimString(True)}, {_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC41) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()

class D14_DC40(D14, NamedTuple('DC40', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC40({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC40) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC43(_dafny.Set({}), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D15_DC43)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D15_DC44)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D15_DC42)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D15_DC45)

class D15_DC43(D15, NamedTuple('DC43', [('cf76', Any), ('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC43({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC43) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC44(D15, NamedTuple('DC44', [('cf78', Any), ('cf79', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC44({_dafny.string_of(self.cf78)}, {_dafny.string_of(self.cf79)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC44) and self.cf78 == __o.cf78 and self.cf79 == __o.cf79
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC42(D15, NamedTuple('DC42', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC42({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC42) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC45(D15, NamedTuple('DC45', [('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC45({_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC45) and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D16_DC46)

class D16_DC46(D16, NamedTuple('DC46', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC46({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC46) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass

class T1:
    pass
    def m1(self, p0, p1, p2, p3, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f1: _dafny.Seq = _dafny.Seq({})
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self.f11: bool = False
        self.f15: bool = False
        self.f16: _dafny.Seq = _dafny.Seq({})
        self.f17: bool = False
        self._f0: str = _dafny.CodePoint('D')
        self._f3: bool = False
        self._f4: _dafny.Seq = _dafny.Seq({})
        self._f5: str = _dafny.CodePoint('D')
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f9: bool = False
        self._f10: bool = False
        self._f12: _dafny.Map = _dafny.Map({})
        self._f13: bool = False
        self._f14: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17):
        (self)._f0 = f0
        (self).f1 = f1
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
        (self)._f12 = f12
        (self)._f13 = f13
        (self)._f14 = f14
        (self).f15 = f15
        (self).f16 = f16
        (self).f17 = f17

    @property
    def f0(self):
        return self._f0
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
    @property
    def f12(self):
        return self._f12
    @property
    def f13(self):
        return self._f13
    @property
    def f14(self):
        return self._f14

class C0(T0):
    def  __init__(self):
        self._f23: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f23):
        (self)._f23 = f23

    def fm14(self, p0, globalState):
        if ((self).f23 if (self).f23 else False):
            return (len(_dafny.SeqWithoutIsStrInference([(self).f23]))) - (493)
        elif True:
            return (969) - (len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opfaxvbk"))})))

    def fm15(self, p0, p1, p2, p3, globalState):
        return not ((self).f23) or ((D0_DC0((self).f23)).cf0)

    @property
    def f23(self):
        return self._f23

class C1(T0):
    def  __init__(self):
        self._f22: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    def ctor__(self, f22):
        (self)._f22 = f22

    def m7(self, p0, p1, globalState):
        d_447_v0_: _dafny.Seq
        d_447_v0_ = _dafny.SeqWithoutIsStrInference([True])
        d_448_v1_: bool
        d_448_v1_ = True
        d_449_v2_: _dafny.Seq
        d_449_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlrnbcijs"))
        hi5_ = p1
        for d_450_i0_ in range((default__.fm9(p0, len(d_447_v0_), d_448_v1_, globalState)) + ((D1_DC5(d_449_v2_, p0, len(d_449_v2_))).cf9), hi5_):
            d_451_v3_: int
            d_451_v3_ = 596
            d_451_v3_ = default__.safeModuloInt(p1, d_451_v3_)
            d_452_v5_: str
            d_452_v5_ = _dafny.CodePoint('j')
            d_453_v6_: _dafny.Set
            d_453_v6_ = _dafny.Set({d_452_v5_})
            d_454_v7_: _dafny.Map
            d_454_v7_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_448_v1_, d_448_v1_, d_448_v1_]): (self).f22})
            d_455_v8_: _dafny.Map
            def iife35_():
                coll21_ = _dafny.Map()
                compr_21_: str
                for compr_21_ in (d_453_v6_).Elements:
                    d_456_v4_: str = compr_21_
                    if (d_456_v4_) in (d_453_v6_):
                        coll21_[d_456_v4_] = d_450_i0_
                return _dafny.Map(coll21_)
            d_455_v8_ = _dafny.Map({iife35_()
            : len(d_454_v7_)})
            d_457_v9_: _dafny.Map
            d_457_v9_ = _dafny.Map({_dafny.CodePoint('m'): 706})
            d_455_v8_ = (d_455_v8_).set(d_457_v9_, p1)
            d_458_v10_: _dafny.Array
            nw80_ = _dafny.Array(False, 2)
            d_458_v10_ = nw80_
            d_459_v11_: D1
            d_459_v11_ = D1_DC4(p1, d_448_v1_, d_458_v10_)
            source7_ = d_459_v11_
            if source7_.is_DC4:
                d_460___mcc_h0_ = source7_.cf4
                d_461___mcc_h1_ = source7_.cf5
                d_462___mcc_h2_ = source7_.cf6
                d_463_cf6_ = d_462___mcc_h2_
                d_464_cf5_ = d_461___mcc_h1_
                d_465_cf4_ = d_460___mcc_h0_
                d_466_v12_: _dafny.Set
                d_466_v12_ = _dafny.Set({p1, p1})
                index49_ = default__.safeIndex(913, (d_458_v10_).length(0))
                (d_458_v10_)[index49_] = (_dafny.Set({p0, d_451_v3_, d_451_v3_})).issubset(d_466_v12_)
                index50_ = default__.safeIndex(913, (d_458_v10_).length(0))
                (d_458_v10_)[index50_] = not(d_448_v1_)
                d_467_v13_: _dafny.MultiSet
                d_467_v13_ = _dafny.MultiSet([p1])
                d_468_v14_: D2
                d_468_v14_ = D2_DC7(d_467_v13_)
                d_448_v1_ = (((d_468_v14_).cf12).cardinality) == ((p1) + (p1))
                d_469_v15_: _dafny.Map
                d_469_v15_ = _dafny.Map({False: len(d_449_v2_)})
                index51_ = default__.safeIndex(913, (d_458_v10_).length(0))
                (d_458_v10_)[index51_] = (default__.safeModuloInt(d_465_cf4_, len(d_469_v15_))) == ((d_451_v3_) + (p1))
                d_470_v16_: _dafny.Map
                d_470_v16_ = _dafny.Map({p0: 624})
                d_471_v17_: D0
                d_471_v17_ = D0_DC1((d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))])
                d_472_v18_: D0
                d_472_v18_ = D0_DC2(d_471_v17_)
                pat_let_tv11_ = d_448_v1_
                pat_let_tv12_ = d_458_v10_
                pat_let_tv13_ = d_458_v10_
                pat_let_tv14_ = globalState
                d_473_v19_: _dafny.Map
                def iife36_(_pat_let7_0):
                    def iife37_(d_474_dt__update__tmp_h0_):
                        def iife38_(_pat_let8_0):
                            def iife39_(d_475_dt__update_hcf2_h0_):
                                return D0_DC2(d_475_dt__update_hcf2_h0_)
                            return iife39_(_pat_let8_0)
                        return iife38_(default__.fm11(d_450_i0_, pat_let_tv11_, (pat_let_tv13_)[default__.safeIndex(913, (pat_let_tv12_).length(0))], d_450_i0_, pat_let_tv14_))
                    return iife37_(_pat_let7_0)
                d_473_v19_ = _dafny.Map({iife36_(d_472_v18_): d_472_v18_})
                nw81_ = _dafny.Array(None, 15)
                nw81_[int(0)] = False
                nw81_[int(1)] = (d_451_v3_) > (d_450_i0_)
                nw81_[int(2)] = (d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))]
                nw81_[int(3)] = (not(d_448_v1_)) == ((d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))])
                nw81_[int(4)] = (-192) > (p0)
                nw81_[int(5)] = ((d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))]) == (d_464_cf5_)
                nw81_[int(6)] = (d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))]
                nw81_[int(7)] = (d_470_v16_) != (default__.fm10(globalState))
                nw81_[int(8)] = True
                nw81_[int(9)] = d_448_v1_
                nw81_[int(10)] = (default__.fm9(d_465_cf4_, d_450_i0_, d_464_cf5_, globalState)) < (len(d_473_v19_))
                nw81_[int(11)] = d_464_cf5_
                nw81_[int(12)] = (d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))]
                nw81_[int(13)] = ((d_458_v10_)[default__.safeIndex(913, (d_458_v10_).length(0))]) or (d_448_v1_)
                nw81_[int(14)] = (d_464_cf5_ if d_448_v1_ else d_464_cf5_)
                d_463_cf6_ = nw81_
            elif source7_.is_DC5:
                d_476___mcc_h3_ = source7_.cf7
                d_477___mcc_h4_ = source7_.cf8
                d_478___mcc_h5_ = source7_.cf9
                d_479_cf9_ = d_478___mcc_h5_
                d_480_cf8_ = d_477___mcc_h4_
                d_481_cf7_ = d_476___mcc_h3_
                d_452_v5_ = d_452_v5_
                d_482_v20_: _dafny.Map
                d_482_v20_ = _dafny.Map({d_448_v1_: d_447_v0_})
                d_482_v20_ = (d_482_v20_).set(default__.fm0(d_448_v1_, d_452_v5_, d_448_v1_, globalState), d_447_v0_)
                d_481_cf7_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss"))).set(default__.safeIndex(d_450_i0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ss")))), d_452_v5_)
                index52_ = default__.safeIndex(384, (d_458_v10_).length(0))
                (d_458_v10_)[index52_] = d_448_v1_
                d_483_v21_: _dafny.Set
                d_483_v21_ = _dafny.Set({p0})
                d_484_v22_: D0
                d_484_v22_ = D0_DC0(d_448_v1_)
                index53_ = default__.safeIndex(384, (d_458_v10_).length(0))
                (d_458_v10_)[index53_] = ((d_484_v22_).cf0 if (d_483_v21_).ispropersubset(d_483_v21_) else d_448_v1_)
            elif source7_.is_DC6:
                d_485___mcc_h6_ = source7_.cf10
                d_486___mcc_h7_ = source7_.cf11
                d_487_cf11_ = d_486___mcc_h7_
                d_488_cf10_ = d_485___mcc_h6_
                d_489_v23_: _dafny.Set
                d_489_v23_ = _dafny.Set({d_448_v1_})
                d_490_v24_: _dafny.Map
                d_490_v24_ = _dafny.Map({d_452_v5_: (d_489_v23_).issubset(d_489_v23_)})
                d_491_v25_: _dafny.MultiSet
                d_491_v25_ = _dafny.MultiSet([69, d_487_cf11_])
                d_490_v24_ = (d_490_v24_).set(d_452_v5_, (_dafny.MultiSet([d_487_cf11_])).issubset(d_491_v25_))
                (globalState).f11 = d_448_v1_
                d_492_v26_: _dafny.Seq
                d_492_v26_ = _dafny.SeqWithoutIsStrInference([d_449_v2_])
                d_493_v27_: _dafny.Map
                d_493_v27_ = _dafny.Map({(self).f22: d_450_i0_})
                d_452_v5_ = default__.fm12((0) - ((0) - (default__.safeModuloInt(p1, p1))), (((d_449_v2_) + ((d_492_v26_)[default__.safeIndex((0) - (d_451_v3_), len(d_492_v26_))])).set(default__.safeIndex(((d_493_v27_)[(self).f22] if ((self).f22) in (d_493_v27_) else d_487_cf11_), len((d_449_v2_) + ((d_492_v26_)[default__.safeIndex((0) - (d_451_v3_), len(d_492_v26_))]))), d_452_v5_)).set(default__.safeIndex((0) - (d_450_i0_), len(((d_449_v2_) + ((d_492_v26_)[default__.safeIndex((0) - (d_451_v3_), len(d_492_v26_))])).set(default__.safeIndex(((d_493_v27_)[(self).f22] if ((self).f22) in (d_493_v27_) else d_487_cf11_), len((d_449_v2_) + ((d_492_v26_)[default__.safeIndex((0) - (d_451_v3_), len(d_492_v26_))]))), d_452_v5_))), d_452_v5_), (d_492_v26_)[default__.safeIndex(p0, len(d_492_v26_))], globalState)
                (globalState).f11 = d_448_v1_
            elif True:
                d_494___mcc_h8_ = source7_.cf3
                d_495_cf3_ = d_494___mcc_h8_
                d_496_v28_: _dafny.Array
                nw82_ = _dafny.Array(_dafny.Seq({}), 18)
                d_496_v28_ = nw82_
                d_497_v29_: _dafny.Seq
                d_497_v29_ = _dafny.SeqWithoutIsStrInference([d_450_i0_])
                index54_ = default__.safeIndex(882, (d_496_v28_).length(0))
                (d_496_v28_)[index54_] = (_dafny.SeqWithoutIsStrInference([d_451_v3_, p1])) + (d_497_v29_)
                d_498_v30_: _dafny.Array
                def lambda27_(d_499_v1_):
                    def lambda28_(d_500_i1_):
                        return _dafny.Set({d_499_v1_, d_499_v1_})

                    return lambda28_

                init13_ = lambda27_(d_448_v1_)
                nw83_ = _dafny.Array(None, 1)
                for i0_13_ in range(nw83_.length(0)):
                    nw83_[i0_13_] = init13_(i0_13_)
                d_498_v30_ = nw83_
                index55_ = default__.safeIndex(882, (d_496_v28_).length(0))
                rhs54_ = _dafny.SeqWithoutIsStrInference([d_452_v5_ for d_501_i2_ in range(default__.abs(266))])
                rhs55_ = d_497_v29_
                rhs56_ = d_498_v30_
                rhs57_ = (len(d_449_v2_)) + ((d_451_v3_ if d_448_v1_ else 578))
                lhs27_ = d_496_v28_
                lhs28_ = default__.safeIndex(882, (d_496_v28_).length(0))
                d_449_v2_ = rhs54_
                lhs27_[lhs28_] = rhs55_
                d_498_v30_ = rhs56_
                d_451_v3_ = rhs57_
                d_502_v31_: _dafny.Seq
                d_502_v31_ = _dafny.SeqWithoutIsStrInference([(d_496_v28_)[default__.safeIndex(882, (d_496_v28_).length(0))], (d_496_v28_)[default__.safeIndex(882, (d_496_v28_).length(0))]])
                d_503_v32_: _dafny.Map
                d_503_v32_ = _dafny.Map({d_502_v31_: default__.safeDivisionInt(p1, default__.fm9(p0, len(d_449_v2_), d_448_v1_, globalState))})
                d_504_v33_: _dafny.Set
                d_504_v33_ = _dafny.Set({464})
                d_503_v32_ = (d_503_v32_).set((_dafny.SeqWithoutIsStrInference([(d_496_v28_)[default__.safeIndex(882, (d_496_v28_).length(0))], d_497_v29_])).set(default__.safeIndex(len(d_504_v33_), len(_dafny.SeqWithoutIsStrInference([(d_496_v28_)[default__.safeIndex(882, (d_496_v28_).length(0))], d_497_v29_]))), (d_496_v28_)[default__.safeIndex(882, (d_496_v28_).length(0))]), p0)
                d_451_v3_ = len((default__.fm13(d_450_i0_, d_450_i0_, globalState)) + (d_502_v31_))
                d_505_v34_: _dafny.Map
                d_505_v34_ = _dafny.Map({d_451_v3_: d_448_v1_})
                d_506_v35_: _dafny.Set
                d_506_v35_ = _dafny.Set({d_505_v34_, d_505_v34_, _dafny.Map({-625: d_448_v1_})})
                d_507_v37_: D3
                def iife40_():
                    coll22_ = _dafny.Map()
                    compr_22_: int
                    for compr_22_ in (d_497_v29_).Elements:
                        d_508_v36_: int = compr_22_
                        if (d_508_v36_) in (d_497_v29_):
                            coll22_[(d_508_v36_) * (p0)] = d_448_v1_
                    return _dafny.Map(coll22_)
                d_507_v37_ = D3_DC9(_dafny.Set({iife40_()
, d_505_v34_, _dafny.Map({(_dafny.MultiSet([d_448_v1_])).cardinality: d_448_v1_})}))
                (globalState).f15 = ((d_506_v35_).intersection(d_506_v35_)).issubset((d_507_v37_).cf16)
            d_451_v3_ = p1
        d_448_v1_ = True
        d_509_v38_: _dafny.Map
        d_509_v38_ = _dafny.Map({p1: d_448_v1_})
        d_510_v39_: _dafny.MultiSet
        d_510_v39_ = _dafny.MultiSet([d_448_v1_])
        d_511_v40_: _dafny.Seq
        d_511_v40_ = _dafny.SeqWithoutIsStrInference([d_510_v39_])
        d_509_v38_ = (d_509_v38_).set(p1, (p1) > (((d_511_v40_)[default__.safeIndex(p1, len(d_511_v40_))]).cardinality))
        d_512_v41_: _dafny.Array
        def lambda29_(d_513_p0_):
            def lambda30_(d_514_i3_):
                return (d_513_p0_) >= (-797)

            return lambda30_

        init14_ = lambda29_(p0)
        nw84_ = _dafny.Array(None, 4)
        for i0_14_ in range(nw84_.length(0)):
            nw84_[i0_14_] = init14_(i0_14_)
        d_512_v41_ = nw84_
        index56_ = default__.safeIndex(199, (d_512_v41_).length(0))
        (d_512_v41_)[index56_] = not (d_448_v1_) or (d_448_v1_)
        d_515_v42_: _dafny.MultiSet
        d_515_v42_ = _dafny.MultiSet([p1])
        index57_ = default__.safeIndex(199, (d_512_v41_).length(0))
        (d_512_v41_)[index57_] = (default__.safeDivisionInt(-892, p1)) not in (d_515_v42_)
        (globalState).f15 = False
        d_516_i4_: int
        d_516_i4_ = 0
        with _dafny.label("0"):
            while (d_512_v41_)[default__.safeIndex(199, (d_512_v41_).length(0))]:
                with _dafny.c_label("0"):
                    if (d_516_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_516_i4_ = (d_516_i4_) + (1)
                    d_517_v43_: C0
                    nw85_ = C0()
                    nw85_.ctor__(d_448_v1_)
                    d_517_v43_ = nw85_
                    d_518_v44_: int
                    d_518_v44_ = -941
                    d_519_v45_: _dafny.Seq
                    d_519_v45_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_518_v44_ = len(d_519_v45_)
                    d_518_v44_ = p0
                    d_448_v1_ = d_448_v1_
                    pass
            pass

    def m8(self, p0, p1, globalState):
        r0: _dafny.Set = _dafny.Set({})
        r1: bool = False
        r2: int = int(0)
        r3: _dafny.Map = _dafny.Map({})
        if p1:
            d_520_v0_: _dafny.Seq
            d_520_v0_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
            d_521_v1_: _dafny.Map
            d_521_v1_ = _dafny.Map({p1: 798})
            d_522_v2_: _dafny.Set
            d_522_v2_ = _dafny.Set({p0, ((d_521_v1_)[False] if (False) in (d_521_v1_) else p0)})
            d_523_v3_: _dafny.Seq
            d_523_v3_ = _dafny.SeqWithoutIsStrInference([p1])
            d_524_v4_: C0
            nw86_ = C0()
            nw86_.ctor__(p1)
            d_524_v4_ = nw86_
            d_525_v5_: _dafny.Map
            d_525_v5_ = _dafny.Map({p1: d_524_v4_})
            d_526_v6_: str
            d_526_v6_ = _dafny.CodePoint('m')
            d_527_v7_: _dafny.Array
            def lambda31_(d_528_v3_):
                def lambda32_(d_529_i0_):
                    return d_528_v3_

                return lambda32_

            init15_ = lambda31_(d_523_v3_)
            nw87_ = _dafny.Array(None, 4)
            for i0_15_ in range(nw87_.length(0)):
                nw87_[i0_15_] = init15_(i0_15_)
            d_527_v7_ = nw87_
            d_530_v8_: D1
            d_530_v8_ = D1_DC3(d_527_v7_)
            d_531_v9_: _dafny.Map
            d_531_v9_ = _dafny.Map({d_530_v8_: p1})
            d_532_v10_: _dafny.Map
            d_532_v10_ = _dafny.Map({(d_525_v5_).set((d_524_v4_).f23, d_524_v4_): default__.fm0((d_524_v4_).f23, d_526_v6_, ((d_531_v9_)[D1_DC3(d_527_v7_)] if (D1_DC3(d_527_v7_)) in (d_531_v9_) else (d_524_v4_).f23), globalState)})
            d_533_v11_: _dafny.Seq
            d_533_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ufps"))
            d_534_v12_: _dafny.MultiSet
            d_534_v12_ = _dafny.MultiSet([(d_524_v4_).fm14(p0, globalState)])
            d_535_v13_: _dafny.Seq
            d_535_v13_ = _dafny.SeqWithoutIsStrInference([d_534_v12_])
            d_536_v14_: _dafny.Array
            nw88_ = _dafny.Array(None, 24)
            nw88_[int(0)] = p1
            nw88_[int(1)] = p1
            nw88_[int(2)] = (p1) and (p1)
            nw88_[int(3)] = p1
            nw88_[int(4)] = True
            nw88_[int(5)] = not((p0) in (d_520_v0_))
            nw88_[int(6)] = p1
            nw88_[int(7)] = p1
            nw88_[int(8)] = p1
            nw88_[int(9)] = (d_522_v2_) != (_dafny.Set({p0, p0, p0}))
            nw88_[int(10)] = p1
            nw88_[int(11)] = p1
            nw88_[int(12)] = (d_523_v3_)[default__.safeIndex(p0, len(d_523_v3_))]
            nw88_[int(13)] = p1
            nw88_[int(14)] = p1
            nw88_[int(15)] = ((d_532_v10_)[d_525_v5_] if (d_525_v5_) in (d_532_v10_) else p1)
            nw88_[int(16)] = (d_520_v0_) <= (d_520_v0_)
            nw88_[int(17)] = (d_526_v6_) not in (_dafny.SeqWithoutIsStrInference([d_526_v6_ for d_537_i1_ in range(default__.abs(514))]))
            nw88_[int(18)] = not((len(d_533_v11_)) <= (len(d_535_v13_)))
            nw88_[int(19)] = p1
            nw88_[int(20)] = True
            nw88_[int(21)] = (default__.fm0((d_524_v4_).f23, (d_533_v11_)[default__.safeIndex(p0, len(d_533_v11_))], p1, globalState)) == ((d_524_v4_).f23)
            nw88_[int(22)] = p1
            nw88_[int(23)] = p1
            d_536_v14_ = nw88_
            index58_ = default__.safeIndex(252, (d_536_v14_).length(0))
            (d_536_v14_)[index58_] = (d_524_v4_).f23
            d_538_v15_: _dafny.Map
            d_538_v15_ = _dafny.Map({p0: (d_524_v4_).f23})
            d_539_v16_: _dafny.Set
            d_539_v16_ = _dafny.Set({d_522_v2_})
            d_540_v17_: _dafny.Set
            d_540_v17_ = _dafny.Set({p1})
            index59_ = default__.safeIndex(252, (d_536_v14_).length(0))
            (d_536_v14_)[index59_] = not ((d_524_v4_).f23) or ((_dafny.Set({p1, ((d_538_v15_)[len(d_539_v16_)] if (len(d_539_v16_)) in (d_538_v15_) else False)})).isdisjoint(d_540_v17_))
            d_541_v18_: D0
            d_541_v18_ = D0_DC0(False)
            d_541_v18_ = d_541_v18_
            d_542_v19_: C0
            nw89_ = C0()
            nw89_.ctor__((not(p1) if p1 else (d_536_v14_)[default__.safeIndex(252, (d_536_v14_).length(0))]))
            d_542_v19_ = nw89_
            d_526_v6_ = d_526_v6_
            index60_ = default__.safeIndex(323, ((self).f22).length(0))
            ((self).f22)[index60_] = len(d_523_v3_)
            index61_ = default__.safeIndex(323, ((self).f22).length(0))
            ((self).f22)[index61_] = (p0) + (len(_dafny.SeqWithoutIsStrInference([p0, len(_dafny.SeqWithoutIsStrInference([(d_536_v14_)[default__.safeIndex(252, (d_536_v14_).length(0))], (d_524_v4_).f23]))])))
        elif True:
            d_543_v20_: _dafny.Map
            d_543_v20_ = _dafny.Map({p0: p1})
            d_544_v21_: _dafny.Set
            d_544_v21_ = _dafny.Set({d_543_v20_})
            source8_ = D3_DC9((d_544_v21_) | (d_544_v21_))
            if source8_.is_DC10:
                d_545___mcc_h0_ = source8_.cf17
                d_546_cf17_ = d_545___mcc_h0_
                d_547_v22_: _dafny.Seq
                d_547_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cnnnwq"))
                d_548_v23_: _dafny.MultiSet
                d_548_v23_ = _dafny.MultiSet([d_547_v22_])
                d_549_v24_: _dafny.Seq
                d_549_v24_ = _dafny.SeqWithoutIsStrInference([d_547_v22_])
                r2 = ((0) - ((d_546_cf17_) + ((d_548_v23_).cardinality))) + (len((d_549_v24_) + (d_549_v24_)))
                d_550_v25_: str
                d_550_v25_ = _dafny.CodePoint('r')
                d_543_v20_ = _dafny.Map({871: default__.fm0(p1, d_550_v25_, p1, globalState)})
                d_551_v26_: D1
                d_551_v26_ = D1_DC5(d_547_v22_, d_546_cf17_, len(_dafny.SeqWithoutIsStrInference([d_550_v25_ for d_552_i2_ in range(default__.abs(-709))])))
                d_546_cf17_ = (d_551_v26_).cf8
                d_553_v27_: _dafny.Seq
                d_553_v27_ = _dafny.SeqWithoutIsStrInference([not(p1)])
                d_554_v28_: _dafny.Array
                nw90_ = _dafny.Array(None, 18)
                nw90_[int(0)] = d_553_v27_
                nw90_[int(1)] = d_553_v27_
                nw90_[int(2)] = d_553_v27_
                nw90_[int(3)] = d_553_v27_
                nw90_[int(4)] = d_553_v27_
                nw90_[int(5)] = (d_553_v27_).set(default__.safeIndex(p0, len(d_553_v27_)), p1)
                nw90_[int(6)] = _dafny.SeqWithoutIsStrInference([not(p1)])
                nw90_[int(7)] = d_553_v27_
                nw90_[int(8)] = d_553_v27_
                nw90_[int(9)] = d_553_v27_
                nw90_[int(10)] = d_553_v27_
                nw90_[int(11)] = d_553_v27_
                nw90_[int(12)] = d_553_v27_
                nw90_[int(13)] = d_553_v27_
                nw90_[int(14)] = d_553_v27_
                nw90_[int(15)] = d_553_v27_
                nw90_[int(16)] = d_553_v27_
                nw90_[int(17)] = d_553_v27_
                d_554_v28_ = nw90_
                d_555_v29_: D1
                d_555_v29_ = D1_DC3(d_554_v28_)
                pat_let_tv15_ = d_554_v28_
                pat_let_tv16_ = d_554_v28_
                def iife41_(_pat_let9_0):
                    def iife42_(d_556_dt__update__tmp_h0_):
                        def iife43_(_pat_let10_0):
                            def iife44_(d_557_dt__update_hcf3_h0_):
                                return D1_DC3(d_557_dt__update_hcf3_h0_)
                            return iife44_(_pat_let10_0)
                        return iife43_((pat_let_tv15_ if False else pat_let_tv16_))
                    return iife42_(_pat_let9_0)
                rhs58_ = iife41_(d_555_v29_)
                rhs59_ = p1
                rhs60_ = (p0) - ((d_546_cf17_) * (d_546_cf17_))
                rhs61_ = p1
                rhs62_ = d_547_v22_
                lhs29_ = globalState
                lhs30_ = globalState
                d_555_v29_ = rhs58_
                lhs29_.f15 = rhs59_
                d_546_cf17_ = rhs60_
                lhs30_.f17 = rhs61_
                d_547_v22_ = rhs62_
            elif source8_.is_DC11:
                d_558___mcc_h1_ = source8_.cf18
                d_559___mcc_h2_ = source8_.cf19
                d_560___mcc_h3_ = source8_.cf20
                d_561_cf20_ = d_560___mcc_h3_
                d_562_cf19_ = d_559___mcc_h2_
                d_563_cf18_ = d_558___mcc_h1_
                r1 = p1
                d_564_v30_: str
                d_564_v30_ = _dafny.CodePoint('x')
                d_565_v31_: _dafny.Array
                nw91_ = _dafny.Array(None, 9)
                nw91_[int(0)] = d_564_v30_
                nw91_[int(1)] = d_564_v30_
                nw91_[int(2)] = d_564_v30_
                nw91_[int(3)] = d_564_v30_
                nw91_[int(4)] = d_564_v30_
                nw91_[int(5)] = d_564_v30_
                nw91_[int(6)] = d_564_v30_
                nw91_[int(7)] = d_564_v30_
                nw91_[int(8)] = d_564_v30_
                d_565_v31_ = nw91_
                d_566_v32_: _dafny.Seq
                d_566_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "urvmnj"))
                d_567_v33_: _dafny.Map
                d_567_v33_ = _dafny.Map({d_561_cf20_: d_566_v32_})
                d_568_v34_: _dafny.MultiSet
                d_568_v34_ = _dafny.MultiSet([d_562_cf19_])
                nw92_ = _dafny.Array(None, 24)
                nw92_[int(0)] = _dafny.CodePoint('i')
                nw92_[int(1)] = d_564_v30_
                nw92_[int(2)] = d_564_v30_
                nw92_[int(3)] = d_564_v30_
                nw92_[int(4)] = d_564_v30_
                nw92_[int(5)] = default__.fm12(d_562_cf19_, ((d_567_v33_)[(self).f22] if ((self).f22) in (d_567_v33_) else d_566_v32_), default__.fm16(False, (d_568_v34_).cardinality, globalState), globalState)
                nw92_[int(6)] = d_564_v30_
                nw92_[int(7)] = d_564_v30_
                nw92_[int(8)] = d_564_v30_
                nw92_[int(9)] = _dafny.CodePoint('m')
                nw92_[int(10)] = d_564_v30_
                nw92_[int(11)] = d_564_v30_
                nw92_[int(12)] = d_564_v30_
                nw92_[int(13)] = d_564_v30_
                nw92_[int(14)] = d_564_v30_
                nw92_[int(15)] = d_564_v30_
                nw92_[int(16)] = d_564_v30_
                nw92_[int(17)] = d_564_v30_
                nw92_[int(18)] = d_564_v30_
                nw92_[int(19)] = d_564_v30_
                nw92_[int(20)] = d_564_v30_
                nw92_[int(21)] = d_564_v30_
                nw92_[int(22)] = d_564_v30_
                nw92_[int(23)] = d_564_v30_
                d_565_v31_ = nw92_
                d_569_v35_: _dafny.Array
                nw93_ = _dafny.Array(None, 12)
                nw93_[int(0)] = p1
                nw93_[int(1)] = p1
                nw93_[int(2)] = p1
                nw93_[int(3)] = p1
                nw93_[int(4)] = p1
                nw93_[int(5)] = True
                nw93_[int(6)] = p1
                nw93_[int(7)] = p1
                nw93_[int(8)] = p1
                nw93_[int(9)] = p1
                nw93_[int(10)] = p1
                nw93_[int(11)] = p1
                d_569_v35_ = nw93_
                d_570_v36_: D1
                d_570_v36_ = D1_DC5(_dafny.SeqWithoutIsStrInference([d_564_v30_ for d_571_i3_ in range(default__.abs(-656))]), d_562_cf19_, d_562_cf19_)
                d_572_v37_: _dafny.Seq
                d_572_v37_ = _dafny.SeqWithoutIsStrInference([d_566_v32_, d_566_v32_, d_566_v32_, d_566_v32_, d_566_v32_])
                d_573_v38_: _dafny.MultiSet
                d_573_v38_ = _dafny.MultiSet([p1])
                d_574_v39_: _dafny.Seq
                d_574_v39_ = _dafny.SeqWithoutIsStrInference([True])
                d_575_v40_: _dafny.Map
                d_575_v40_ = _dafny.Map({not(p1): d_574_v39_})
                d_576_v41_: _dafny.Array
                nw94_ = _dafny.Array(None, 26)
                nw94_[int(0)] = p1
                nw94_[int(1)] = p1
                nw94_[int(2)] = p1
                nw94_[int(3)] = not(default__.fm0((D1_DC4(d_562_cf19_, p1, d_569_v35_)).cf5, d_564_v30_, p1, globalState))
                nw94_[int(4)] = p1
                nw94_[int(5)] = not (p1) or (p1)
                nw94_[int(6)] = False
                nw94_[int(7)] = p1
                nw94_[int(8)] = False
                nw94_[int(9)] = False
                nw94_[int(10)] = (len((d_570_v36_).cf7)) > (d_562_cf19_)
                nw94_[int(11)] = p1
                nw94_[int(12)] = not(p1)
                nw94_[int(13)] = default__.fm0(p1, d_564_v30_, default__.fm0(p1, d_564_v30_, p1, globalState), globalState)
                nw94_[int(14)] = p1
                nw94_[int(15)] = (d_572_v37_) == (d_572_v37_)
                nw94_[int(16)] = p1
                nw94_[int(17)] = p1
                nw94_[int(18)] = (_dafny.SeqWithoutIsStrInference([d_564_v30_ for d_577_i4_ in range(default__.abs(-100))])) < (d_566_v32_)
                nw94_[int(19)] = p1
                nw94_[int(20)] = (d_573_v38_) == (d_573_v38_)
                nw94_[int(21)] = p1
                nw94_[int(22)] = p1
                nw94_[int(23)] = not(p1)
                nw94_[int(24)] = p1
                nw94_[int(25)] = (((d_575_v40_)[p1] if (p1) in (d_575_v40_) else d_574_v39_)) == ((d_574_v39_).set(default__.safeIndex(p0, len(d_574_v39_)), False))
                d_576_v41_ = nw94_
                index62_ = default__.safeIndex(852, (d_576_v41_).length(0))
                (d_576_v41_)[index62_] = p1
                index63_ = default__.safeIndex(852, (d_576_v41_).length(0))
                (d_576_v41_)[index63_] = default__.fm0(p1, d_564_v30_, p1, globalState)
                r2 = default__.fm9(-77, d_562_cf19_, (d_576_v41_)[default__.safeIndex(852, (d_576_v41_).length(0))], globalState)
            elif True:
                d_578___mcc_h4_ = source8_.cf16
                d_579_cf16_ = d_578___mcc_h4_
                d_580_v42_: _dafny.Array
                nw95_ = _dafny.Array(False, 22)
                d_580_v42_ = nw95_
                d_580_v42_ = (d_580_v42_ if p1 else d_580_v42_)
                index64_ = default__.safeIndex(525, ((self).f22).length(0))
                ((self).f22)[index64_] = p0
                index65_ = default__.safeIndex(525, ((self).f22).length(0))
                ((self).f22)[index65_] = 276
                d_581_v43_: str
                d_581_v43_ = _dafny.CodePoint('a')
                (globalState).f15 = default__.fm0((p1 if False else p1), d_581_v43_, p1, globalState)
                d_582_v44_: _dafny.Seq
                d_582_v44_ = _dafny.SeqWithoutIsStrInference([((self).f22)[default__.safeIndex(525, ((self).f22).length(0))]])
                index66_ = default__.safeIndex(525, ((self).f22).length(0))
                ((self).f22)[index66_] = default__.safeDivisionInt(((d_582_v44_)[default__.safeIndex(p0, len(d_582_v44_))] if p1 else (0) - (p0)), ((self).f22)[default__.safeIndex(525, ((self).f22).length(0))])
            d_583_v45_: _dafny.Array
            nw96_ = _dafny.Array(None, 3)
            nw96_[int(0)] = (p0) <= (667)
            nw96_[int(1)] = not((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_584_i5_ in range(default__.abs(-190))]))) >= (p0))
            nw96_[int(2)] = p1
            d_583_v45_ = nw96_
            index67_ = default__.safeIndex(123, (d_583_v45_).length(0))
            (d_583_v45_)[index67_] = True
            d_585_v46_: _dafny.Seq
            d_585_v46_ = _dafny.SeqWithoutIsStrInference([len(default__.fm16(not(p1), p0, globalState))])
            index68_ = default__.safeIndex(123, (d_583_v45_).length(0))
            (d_583_v45_)[index68_] = ((543) in (d_585_v46_)) == (p1)
            d_586_v47_: _dafny.Array
            nw97_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_586_v47_ = nw97_
            d_586_v47_ = d_586_v47_
            d_587_v48_: _dafny.Map
            d_587_v48_ = _dafny.Map({(self).f22: (d_583_v45_)[default__.safeIndex(123, (d_583_v45_).length(0))]})
            d_588_v49_: D0
            d_588_v49_ = D0_DC0((p0) <= (p0))
            pat_let_tv17_ = d_583_v45_
            pat_let_tv18_ = d_583_v45_
            def iife45_(_pat_let11_0):
                def iife46_(d_589_dt__update__tmp_h1_):
                    def iife47_(_pat_let12_0):
                        def iife48_(d_590_dt__update_hcf0_h0_):
                            return D0_DC0(d_590_dt__update_hcf0_h0_)
                        return iife48_(_pat_let12_0)
                    return iife47_((pat_let_tv18_)[default__.safeIndex(123, (pat_let_tv17_).length(0))])
                return iife46_(_pat_let11_0)
            rhs63_ = p0
            rhs64_ = (d_587_v48_) | (d_587_v48_)
            rhs65_ = iife45_(D0_DC0(p1))
            rhs66_ = p0
            rhs67_ = p0
            r2 = rhs63_
            d_587_v48_ = rhs64_
            d_588_v49_ = rhs65_
            r2 = rhs66_
            r2 = rhs67_
            d_591_v50_: _dafny.Seq
            d_591_v50_ = _dafny.SeqWithoutIsStrInference([(d_583_v45_)[default__.safeIndex(123, (d_583_v45_).length(0))], p1])
            d_592_v51_: _dafny.MultiSet
            d_592_v51_ = _dafny.MultiSet([d_591_v50_, _dafny.SeqWithoutIsStrInference([(d_583_v45_)[default__.safeIndex(123, (d_583_v45_).length(0))]])])
            r2 = (p0 if ((d_592_v51_).set(d_591_v50_, default__.abs(p0))).issubset((d_592_v51_).set(d_591_v50_, default__.abs(p0))) else p0)
        d_593_v52_: C0
        nw98_ = C0()
        nw98_.ctor__(not (p1) or (p1))
        d_593_v52_ = nw98_
        d_594_v53_: _dafny.Seq
        d_594_v53_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fnymofdaa"))
        d_595_v54_: _dafny.MultiSet
        d_595_v54_ = _dafny.MultiSet([p1, p1])
        d_596_v55_: D1
        d_596_v55_ = D1_DC6(p0, (d_595_v54_).cardinality)
        pat_let_tv19_ = d_594_v53_
        pat_let_tv20_ = d_594_v53_
        pat_let_tv21_ = d_594_v53_
        pat_let_tv22_ = d_594_v53_
        pat_let_tv23_ = d_594_v53_
        def lambda33_(source9_):
            if source9_.is_DC4:
                d_597___mcc_h5_ = source9_.cf4
                d_598___mcc_h6_ = source9_.cf5
                d_599___mcc_h7_ = source9_.cf6
                d_600_cf6_ = d_599___mcc_h7_
                d_601_cf5_ = d_598___mcc_h6_
                d_602_cf4_ = d_597___mcc_h5_
                return (pat_let_tv19_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgsdhwirx")))
            elif source9_.is_DC5:
                d_603___mcc_h8_ = source9_.cf7
                d_604___mcc_h9_ = source9_.cf8
                d_605___mcc_h10_ = source9_.cf9
                d_606_cf9_ = d_605___mcc_h10_
                d_607_cf8_ = d_604___mcc_h9_
                d_608_cf7_ = d_603___mcc_h8_
                return pat_let_tv20_
            elif source9_.is_DC6:
                d_609___mcc_h11_ = source9_.cf10
                d_610___mcc_h12_ = source9_.cf11
                d_611_cf11_ = d_610___mcc_h12_
                d_612_cf10_ = d_609___mcc_h11_
                return (pat_let_tv21_) + (pat_let_tv22_)
            elif True:
                d_613___mcc_h13_ = source9_.cf3
                d_614_cf3_ = d_613___mcc_h13_
                return pat_let_tv23_

        d_594_v53_ = lambda33_(d_596_v55_)
        d_615_v56_: _dafny.Map
        d_615_v56_ = _dafny.Map({p1: d_594_v53_})
        d_616_v57_: D0
        d_616_v57_ = D0_DC1(p1)
        d_617_v58_: _dafny.Set
        d_617_v58_ = _dafny.Set({d_616_v57_})
        d_618_v59_: D1
        d_618_v59_ = D1_DC5(((d_615_v56_)[True] if (True) in (d_615_v56_) else d_594_v53_), p0, (0) - (len(d_617_v58_)))
        d_619_v60_: D3
        d_619_v60_ = D3_DC10(p0)
        d_620_v61_: _dafny.Map
        d_620_v61_ = _dafny.Map({p0: (self).f22})
        d_618_v59_ = D1_DC5(d_594_v53_, (d_619_v60_).cf17, (p0) - (len(d_620_v61_)))
        d_621_v62_: _dafny.Array
        def lambda34_(d_622_p0_):
            def lambda35_(d_623_i7_):
                return default__.safeDivisionInt(d_623_i7_, (_dafny.MultiSet([d_622_p0_, d_622_p0_, d_622_p0_, d_622_p0_, d_622_p0_])).cardinality)

            return lambda35_

        init16_ = lambda34_(p0)
        nw99_ = _dafny.Array(None, 11)
        for i0_16_ in range(nw99_.length(0)):
            nw99_[i0_16_] = init16_(i0_16_)
        d_621_v62_ = nw99_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_621_v62_).length(0)):
            d_624_i6_: int = guard_loop_1_
            if (True) and (((0) <= (d_624_i6_)) and ((d_624_i6_) < ((d_621_v62_).length(0)))):
                (d_621_v62_)[(d_624_i6_)] = default__.safeDivisionInt(d_624_i6_, (d_595_v54_).cardinality)
        d_625_v63_: str
        d_625_v63_ = _dafny.CodePoint('i')
        d_626_v64_: _dafny.Map
        d_626_v64_ = _dafny.Map({p0: d_625_v63_})
        d_626_v64_ = (d_626_v64_).set(p0, d_625_v63_)
        d_627_v65_: _dafny.Set
        d_627_v65_ = _dafny.Set({(0) - (p0)})
        r0 = d_627_v65_
        d_628_v66_: _dafny.Map
        d_628_v66_ = _dafny.Map({p1: (d_593_v52_).f23})
        d_629_v67_: _dafny.Map
        d_629_v67_ = _dafny.Map({p1: p0})
        d_630_v68_: _dafny.Map
        d_630_v68_ = _dafny.Map({p0: ((d_628_v66_)[p1] if (p1) in (d_628_v66_) else (d_593_v52_).fm15(False, (d_593_v52_).f23, p0, d_629_v67_, globalState))})
        r1 = ((d_630_v68_)[(((d_629_v67_)[p1] if (p1) in (d_629_v67_) else p0)) + (462)] if ((((d_629_v67_)[p1] if (p1) in (d_629_v67_) else p0)) + (462)) in (d_630_v68_) else p1)
        r2 = p0
        d_631_v70_: T0
        nw100_ = C0()
        nw100_.ctor__((d_593_v52_).f23)
        d_631_v70_ = nw100_
        d_632_v71_: _dafny.Seq
        d_632_v71_ = _dafny.SeqWithoutIsStrInference([d_631_v70_])
        d_633_v73_: _dafny.MultiSet
        d_633_v73_ = _dafny.MultiSet([p0])
        d_634_v74_: _dafny.Seq
        d_634_v74_ = _dafny.SeqWithoutIsStrInference([p0, (d_633_v73_).cardinality, p0])
        d_635_v75_: _dafny.Seq
        def iife49_():
            coll23_ = _dafny.Map()
            compr_23_: int
            for compr_23_ in (d_634_v74_).Elements:
                d_636_v72_: int = compr_23_
                if (d_636_v72_) in (d_634_v74_):
                    coll23_[(d_636_v72_) * (len(_dafny.SeqWithoutIsStrInference([(d_593_v52_).f23])))] = 96
            return _dafny.Map(coll23_)
        d_635_v75_ = _dafny.SeqWithoutIsStrInference([len(d_632_v71_), p0, len(iife49_()
        ), p0])
        def iife50_():
            coll24_ = _dafny.Map()
            compr_24_: int
            for compr_24_ in (_dafny.MultiSet(d_635_v75_)).Elements:
                d_637_v69_: int = compr_24_
                if (d_637_v69_) in (_dafny.MultiSet(d_635_v75_)):
                    coll24_[default__.safeDivisionInt(d_637_v69_, p0)] = (d_593_v52_).f23
            return _dafny.Map(coll24_)
        r3 = iife50_()
        
        return r0, r1, r2, r3

    @property
    def f22(self):
        return self._f22

class C2(T0, T1):
    def  __init__(self):
        self.f20: D0 = D0.default()()
        self._f21: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    def ctor__(self, f20, f21):
        (self).f20 = f20
        (self)._f21 = f21

    def fm8(self, p0, globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qr"))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tglyx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        d_638_i0_: int
        d_638_i0_ = 0
        with _dafny.label("1"):
            while p2:
                with _dafny.c_label("1"):
                    if (d_638_i0_) >= (100):
                        raise _dafny.Break("1")
                    d_638_i0_ = (d_638_i0_) + (1)
                    (globalState).f11 = not(p0)
                    d_639_v0_: _dafny.Seq
                    d_639_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qletf"))
                    rhs68_ = p2
                    rhs69_ = ((d_639_v0_) + (d_639_v0_)) <= (d_639_v0_)
                    lhs31_ = globalState
                    lhs32_ = globalState
                    lhs31_.f15 = rhs68_
                    lhs32_.f15 = rhs69_
                    d_640_v1_: _dafny.MultiSet
                    d_640_v1_ = _dafny.MultiSet([328])
                    d_641_v2_: _dafny.Seq
                    d_641_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_642_i1_ in range(default__.abs(456))])])
                    (globalState).f17 = ((d_640_v1_).set(p3, default__.abs(len(d_641_v2_)))).isdisjoint(d_640_v1_)
                    d_643_v3_: _dafny.Array
                    nw101_ = _dafny.Array(int(0), 15)
                    d_643_v3_ = nw101_
                    d_644_v4_: T0
                    nw102_ = C1()
                    nw102_.ctor__(d_643_v3_)
                    d_644_v4_ = nw102_
                    d_645_v5_: D4
                    d_645_v5_ = D4_DC12(d_644_v4_)
                    d_646_v6_: _dafny.Seq
                    d_646_v6_ = _dafny.SeqWithoutIsStrInference([d_644_v4_, d_644_v4_])
                    d_647_v7_: _dafny.Map
                    d_647_v7_ = _dafny.Map({p3: d_643_v3_})
                    d_648_v8_: _dafny.Map
                    d_648_v8_ = _dafny.Map({((d_647_v7_)[p1] if (p1) in (d_647_v7_) else d_643_v3_): p1})
                    d_649_v9_: _dafny.Array
                    nw103_ = _dafny.Array(None, 12)
                    nw103_[int(0)] = (d_644_v4_ if (self).f21 else (d_645_v5_).cf21)
                    nw103_[int(1)] = d_644_v4_
                    nw103_[int(2)] = d_644_v4_
                    nw103_[int(3)] = d_644_v4_
                    nw103_[int(4)] = d_644_v4_
                    nw103_[int(5)] = d_644_v4_
                    nw103_[int(6)] = (d_646_v6_)[default__.safeIndex(((d_648_v8_)[d_643_v3_] if (d_643_v3_) in (d_648_v8_) else p1), len(d_646_v6_))]
                    nw103_[int(7)] = d_644_v4_
                    nw103_[int(8)] = (d_644_v4_ if (self).f21 else d_644_v4_)
                    nw103_[int(9)] = d_644_v4_
                    nw103_[int(10)] = d_644_v4_
                    nw103_[int(11)] = d_644_v4_
                    d_649_v9_ = nw103_
                    index69_ = default__.safeIndex(743, (d_649_v9_).length(0))
                    (d_649_v9_)[index69_] = (d_646_v6_)[default__.safeIndex(199, len(d_646_v6_))]
                    index70_ = default__.safeIndex(743, (d_649_v9_).length(0))
                    (d_649_v9_)[index70_] = d_644_v4_
                    pass
            pass
        d_650_v10_: int
        d_650_v10_ = -235
        d_651_v11_: _dafny.Seq
        d_651_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
        d_652_v12_: _dafny.Map
        d_652_v12_ = _dafny.Map({p3: (self).f21})
        d_650_v10_ = default__.fm9((0) - (len((d_651_v11_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_653_i2_ in range(default__.abs(278))])))), default__.safeModuloInt(p3, d_650_v10_), ((d_652_v12_)[p3] if (p3) in (d_652_v12_) else p0), globalState)
        if not (not(not(not(p2)))) or (((self).f21 if (self).f21 else (self).f21)):
            d_654_v13_: _dafny.Seq
            d_654_v13_ = _dafny.SeqWithoutIsStrInference([d_650_v10_])
            d_650_v10_ = (p3) * ((len(d_654_v13_)) - (p3))
            d_650_v10_ = default__.fm9(d_650_v10_, p3, p0, globalState)
            (globalState).f11 = not((self).f21)
            d_655_v14_: _dafny.Array
            def lambda36_(d_656_i3_):
                return (self).f21

            init17_ = lambda36_
            nw104_ = _dafny.Array(None, 20)
            for i0_17_ in range(nw104_.length(0)):
                nw104_[i0_17_] = init17_(i0_17_)
            d_655_v14_ = nw104_
            d_657_v15_: _dafny.Seq
            d_657_v15_ = _dafny.SeqWithoutIsStrInference([d_655_v14_])
            d_657_v15_ = d_657_v15_
            d_650_v10_ = (0) - (p3)
        elif True:
            d_658_v16_: _dafny.Array
            nw105_ = _dafny.Array(False, 21)
            d_658_v16_ = nw105_
            index71_ = default__.safeIndex(133, (d_658_v16_).length(0))
            (d_658_v16_)[index71_] = (len(d_651_v11_)) > (d_650_v10_)
            index72_ = default__.safeIndex(133, (d_658_v16_).length(0))
            (d_658_v16_)[index72_] = p2
            d_658_v16_ = d_658_v16_
            if p2:
                d_659_v17_: _dafny.Array
                def lambda37_(d_660_v12_):
                    def lambda38_(d_661_i4_):
                        return d_660_v12_

                    return lambda38_

                init18_ = lambda37_(d_652_v12_)
                nw106_ = _dafny.Array(None, 15)
                for i0_18_ in range(nw106_.length(0)):
                    nw106_[i0_18_] = init18_(i0_18_)
                d_659_v17_ = nw106_
                d_662_v18_: _dafny.Seq
                d_662_v18_ = _dafny.SeqWithoutIsStrInference([not(p2), p0])
                index73_ = default__.safeIndex(32, (d_659_v17_).length(0))
                (d_659_v17_)[index73_] = default__.fm17((d_662_v18_)[default__.safeIndex(p3, len(d_662_v18_))], (self).f21, globalState)
                index74_ = default__.safeIndex(32, (d_659_v17_).length(0))
                (d_659_v17_)[index74_] = d_652_v12_
                d_663_v19_: _dafny.Array
                nw107_ = _dafny.Array(int(0), 7)
                d_663_v19_ = nw107_
                d_664_v20_: _dafny.Set
                d_664_v20_ = _dafny.Set({(self).f21})
                index75_ = default__.safeIndex(216, (d_663_v19_).length(0))
                (d_663_v19_)[index75_] = len((d_664_v20_) | (d_664_v20_))
                d_665_v21_: str
                d_665_v21_ = _dafny.CodePoint('g')
                index76_ = default__.safeIndex(216, (d_663_v19_).length(0))
                (d_663_v19_)[index76_] = len(((d_651_v11_).set(default__.safeIndex(p1, len(d_651_v11_)), d_665_v21_)) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pddyfcmi"))) + (d_651_v11_)))
                d_666_v22_: _dafny.MultiSet
                d_666_v22_ = _dafny.MultiSet([p1, p3])
                (globalState).f17 = (((_dafny.MultiSet([(0) - ((0) - (len((D1_DC5(d_651_v11_, d_650_v10_, (d_663_v19_)[default__.safeIndex(216, (d_663_v19_).length(0))])).cf7))), p1])) | (d_666_v22_)).cardinality) < ((d_663_v19_)[default__.safeIndex(216, (d_663_v19_).length(0))])
                (globalState).f15 = (p1) != (len((d_651_v11_ if not((d_658_v16_)[default__.safeIndex(133, (d_658_v16_).length(0))]) else _dafny.SeqWithoutIsStrInference([d_665_v21_ for d_667_i5_ in range(default__.abs(84))]))))
                (self).f20 = self.f20
            elif True:
                d_668_v23_: _dafny.Seq
                d_668_v23_ = _dafny.SeqWithoutIsStrInference([((d_652_v12_)[p3] if (p3) in (d_652_v12_) else (self).f21), p2])
                d_669_v24_: D4
                d_669_v24_ = D4_DC14((d_668_v23_)[default__.safeIndex(p3, len(d_668_v23_))], (0) - (d_650_v10_), _dafny.Set({(_dafny.MultiSet([p1])).cardinality, -869, len(d_652_v12_)}), _dafny.CodePoint('l'))
                d_670_v25_: _dafny.Set
                d_670_v25_ = _dafny.Set({d_658_v16_})
                d_671_v26_: _dafny.Map
                d_671_v26_ = _dafny.Map({(d_658_v16_)[default__.safeIndex(133, (d_658_v16_).length(0))]: d_670_v25_})
                rhs70_ = d_669_v24_
                rhs71_ = d_650_v10_
                rhs72_ = ((d_670_v25_ if p0 else ((d_671_v26_)[(d_658_v16_)[default__.safeIndex(133, (d_658_v16_).length(0))]] if ((d_658_v16_)[default__.safeIndex(133, (d_658_v16_).length(0))]) in (d_671_v26_) else d_670_v25_))) | ((d_670_v25_ if p2 else d_670_v25_))
                d_669_v24_ = rhs70_
                d_650_v10_ = rhs71_
                d_670_v25_ = rhs72_
                d_672_v27_: str
                d_672_v27_ = _dafny.CodePoint('w')
                d_673_v28_: _dafny.MultiSet
                d_673_v28_ = _dafny.MultiSet([d_672_v27_])
                d_673_v28_ = d_673_v28_
                (globalState).f17 = p2
                def iife51_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in _dafny.IntegerRange(447, -428):
                        d_674_v29_: int = compr_25_
                        if ((447) <= (d_674_v29_)) and ((d_674_v29_) < (-428)):
                            coll25_[default__.safeModuloInt(d_674_v29_, len(_dafny.SeqWithoutIsStrInference([(d_658_v16_)[default__.safeIndex(133, (d_658_v16_).length(0))]])))] = d_668_v23_
                    return _dafny.Map(coll25_)
                d_650_v10_ = default__.safeDivisionInt(d_650_v10_, (len(iife51_()
                )) * (p1))
                d_675_v30_: _dafny.Set
                d_675_v30_ = _dafny.Set({-433})
                d_676_v31_: _dafny.Set
                d_676_v31_ = _dafny.Set({p3, len(d_675_v30_), p3, default__.fm9(d_650_v10_, p1, p0, globalState)})
                (globalState).f11 = (d_676_v31_).issubset(_dafny.Set({default__.fm9(d_650_v10_, p1, p0, globalState)}))
            index77_ = default__.safeIndex(133, (d_658_v16_).length(0))
            (d_658_v16_)[index77_] = (p3) != (p1)
            d_677_v32_: C0
            nw108_ = C0()
            nw108_.ctor__((d_658_v16_)[default__.safeIndex(133, (d_658_v16_).length(0))])
            d_677_v32_ = nw108_
        if p2:
            d_678_v33_: _dafny.Array
            def lambda39_(d_679_p3_):
                def lambda40_(d_680_i6_):
                    return default__.safeDivisionInt(d_680_i6_, d_679_p3_)

                return lambda40_

            init19_ = lambda39_(p3)
            nw109_ = _dafny.Array(None, 17)
            for i0_19_ in range(nw109_.length(0)):
                nw109_[i0_19_] = init19_(i0_19_)
            d_678_v33_ = nw109_
            d_681_v34_: _dafny.Seq
            d_681_v34_ = _dafny.SeqWithoutIsStrInference([p1, d_650_v10_])
            d_682_v35_: _dafny.Map
            d_682_v35_ = _dafny.Map({(self).f21: (self).f21})
            d_683_v36_: _dafny.Map
            d_683_v36_ = _dafny.Map({d_681_v34_: d_682_v35_})
            index78_ = default__.safeIndex(322, (d_678_v33_).length(0))
            (d_678_v33_)[index78_] = len(d_683_v36_)
            index79_ = default__.safeIndex(322, (d_678_v33_).length(0))
            (d_678_v33_)[index79_] = default__.safeDivisionInt((0) - (default__.safeDivisionInt(p3, (0) - (d_650_v10_))), (d_650_v10_ if not((self).f21) else d_650_v10_))
            d_684_v37_: _dafny.Array
            nw110_ = _dafny.Array(_dafny.Set({}), 10)
            d_684_v37_ = nw110_
            d_685_v38_: _dafny.MultiSet
            d_685_v38_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yyioc"))])
            d_686_v39_: _dafny.Set
            d_686_v39_ = _dafny.Set({p0})
            d_687_v40_: D4
            d_687_v40_ = D4_DC13(d_650_v10_)
            index80_ = default__.safeIndex(322, (d_678_v33_).length(0))
            index81_ = default__.safeIndex(322, (d_678_v33_).length(0))
            rhs73_ = p1
            rhs74_ = d_684_v37_
            rhs75_ = (((d_685_v38_)[d_651_v11_] if (d_651_v11_) in (d_685_v38_) else len(d_686_v39_))) - ((d_687_v40_).cf22)
            rhs76_ = d_651_v11_
            lhs33_ = d_678_v33_
            lhs34_ = default__.safeIndex(322, (d_678_v33_).length(0))
            lhs35_ = d_678_v33_
            lhs36_ = default__.safeIndex(322, (d_678_v33_).length(0))
            lhs33_[lhs34_] = rhs73_
            d_684_v37_ = rhs74_
            lhs35_[lhs36_] = rhs75_
            d_651_v11_ = rhs76_
            d_688_v41_: _dafny.Array
            def lambda41_(d_689_v11_):
                def lambda42_(d_690_i7_):
                    return d_689_v11_

                return lambda42_

            init20_ = lambda41_(d_651_v11_)
            nw111_ = _dafny.Array(None, 20)
            for i0_20_ in range(nw111_.length(0)):
                nw111_[i0_20_] = init20_(i0_20_)
            d_688_v41_ = nw111_
            index82_ = default__.safeIndex(845, (d_688_v41_).length(0))
            (d_688_v41_)[index82_] = d_651_v11_
            d_691_v42_: _dafny.Map
            d_691_v42_ = _dafny.Map({d_650_v10_: p1})
            d_692_v43_: _dafny.Set
            d_692_v43_ = _dafny.Set({((d_691_v42_)[p1] if (p1) in (d_691_v42_) else 615)})
            d_693_v44_: str
            d_693_v44_ = _dafny.CodePoint('v')
            index83_ = default__.safeIndex(845, (d_688_v41_).length(0))
            rhs77_ = ((d_651_v11_) + (d_651_v11_)).set(default__.safeIndex(default__.fm9(p3, p3, False, globalState), len((d_651_v11_) + (d_651_v11_))), d_693_v44_)
            rhs78_ = _dafny.Set({(p1) + ((d_678_v33_)[default__.safeIndex(322, (d_678_v33_).length(0))]), p1, len(d_681_v34_)})
            lhs37_ = d_688_v41_
            lhs38_ = default__.safeIndex(845, (d_688_v41_).length(0))
            lhs37_[lhs38_] = rhs77_
            d_692_v43_ = rhs78_
            d_694_v45_: D1
            d_694_v45_ = D1_DC6(p3, (d_678_v33_)[default__.safeIndex(322, (d_678_v33_).length(0))])
            d_694_v45_ = D1_DC6(p3, p3)
            d_695_v46_: C0
            nw112_ = C0()
            nw112_.ctor__((self).f21)
            d_695_v46_ = nw112_
        elif True:
            d_650_v10_ = 458
            d_650_v10_ = (d_650_v10_) - (p3)
            d_696_v47_: _dafny.Array
            nw113_ = _dafny.Array(int(0), 12)
            d_696_v47_ = nw113_
            d_697_v48_: C1
            nw114_ = C1()
            nw114_.ctor__(d_696_v47_)
            d_697_v48_ = nw114_
            d_698_v49_: _dafny.Map
            d_698_v49_ = _dafny.Map({d_697_v48_: p0})
            d_699_v50_: _dafny.Map
            d_699_v50_ = _dafny.Map({((d_698_v49_)[d_697_v48_] if (d_697_v48_) in (d_698_v49_) else (self).f21): (self).f21})
            d_699_v50_ = (d_699_v50_).set(((d_652_v12_)[d_650_v10_] if (d_650_v10_) in (d_652_v12_) else p0), p2)
            d_650_v10_ = p1
            d_700_v51_: _dafny.Seq
            d_700_v51_ = _dafny.SeqWithoutIsStrInference([(self).f21, (d_651_v11_) != (d_651_v11_), p0])
            rhs79_ = not((d_700_v51_)[default__.safeIndex(p1, len(d_700_v51_))])
            rhs80_ = p0
            rhs81_ = p2
            lhs39_ = globalState
            lhs40_ = globalState
            lhs41_ = globalState
            lhs39_.f15 = rhs79_
            lhs40_.f11 = rhs80_
            lhs41_.f15 = rhs81_
        if (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_701_i8_ in range(default__.abs(-184))])) != (d_651_v11_):
            d_702_v52_: _dafny.Set
            d_702_v52_ = _dafny.Set({p3})
            d_650_v10_ = (p1 if (d_702_v52_) == (_dafny.Set({len(d_652_v12_)})) else (d_650_v10_) + (p3))
            d_652_v12_ = (d_652_v12_).set(707, (default__.fm18((self).f21, d_650_v10_, p3, globalState)).issubset(d_702_v52_))
            d_650_v10_ = d_650_v10_
            d_650_v10_ = d_650_v10_
            d_703_v53_: _dafny.Map
            d_703_v53_ = _dafny.Map({(self).f21: p1})
            d_704_v54_: _dafny.Seq
            d_704_v54_ = _dafny.SeqWithoutIsStrInference([p2, (self).f21])
            d_705_v55_: _dafny.Seq
            d_705_v55_ = _dafny.SeqWithoutIsStrInference([p1])
            d_706_v56_: _dafny.Array
            nw115_ = _dafny.Array(None, 20)
            nw115_[int(0)] = len(d_704_v54_)
            nw115_[int(1)] = p3
            nw115_[int(2)] = d_650_v10_
            nw115_[int(3)] = p1
            nw115_[int(4)] = p3
            nw115_[int(5)] = (d_705_v55_)[default__.safeIndex(len(_dafny.Set({(self).f21, True, False, p2})), len(d_705_v55_))]
            nw115_[int(6)] = (0) - (p1)
            nw115_[int(7)] = d_650_v10_
            nw115_[int(8)] = d_650_v10_
            nw115_[int(9)] = p1
            nw115_[int(10)] = (d_705_v55_)[default__.safeIndex(p1, len(d_705_v55_))]
            nw115_[int(11)] = len(d_652_v12_)
            nw115_[int(12)] = p1
            nw115_[int(13)] = p3
            nw115_[int(14)] = len(d_705_v55_)
            nw115_[int(15)] = p1
            nw115_[int(16)] = (0) - (p3)
            nw115_[int(17)] = p3
            nw115_[int(18)] = d_650_v10_
            nw115_[int(19)] = d_650_v10_
            d_706_v56_ = nw115_
            d_707_v57_: _dafny.MultiSet
            d_707_v57_ = _dafny.MultiSet([False, p2, p2, p2])
            d_708_v58_: D4
            d_708_v58_ = D4_DC14((self).f21, d_650_v10_, d_702_v52_, (d_651_v11_)[default__.safeIndex((d_707_v57_).cardinality, len(d_651_v11_))])
            d_709_v59_: _dafny.Map
            d_709_v59_ = _dafny.Map({D3_DC11(d_703_v53_, 919, d_706_v56_): d_708_v58_})
            d_650_v10_ = len(d_709_v59_)
        elif True:
            d_710_v60_: _dafny.Array
            def lambda43_(d_711_v10_, d_712_p3_, d_713_p2_):
                def lambda44_(d_714_i9_):
                    return (_dafny.Set({len(_dafny.Map({d_711_v10_: _dafny.CodePoint('s')})), d_712_p3_, len(_dafny.SeqWithoutIsStrInference([d_713_p2_])), d_712_p3_})).isdisjoint(_dafny.Set({d_712_p3_}))

                return lambda44_

            init21_ = lambda43_(d_650_v10_, p3, p2)
            nw116_ = _dafny.Array(None, 26)
            for i0_21_ in range(nw116_.length(0)):
                nw116_[i0_21_] = init21_(i0_21_)
            d_710_v60_ = nw116_
            d_710_v60_ = d_710_v60_
            (globalState).f15 = (p0) and (p2)
            if (d_650_v10_) < (d_650_v10_):
                d_715_v61_: str
                d_715_v61_ = _dafny.CodePoint('x')
                (globalState).f17 = (d_715_v61_) not in (d_651_v11_)
                d_716_v62_: _dafny.Array
                def lambda45_(d_717_p3_):
                    def lambda46_(d_718_i10_):
                        return D4_DC13(d_717_p3_)

                    return lambda46_

                init22_ = lambda45_(p3)
                nw117_ = _dafny.Array(None, 4)
                for i0_22_ in range(nw117_.length(0)):
                    nw117_[i0_22_] = init22_(i0_22_)
                d_716_v62_ = nw117_
                d_719_v63_: _dafny.Map
                d_719_v63_ = _dafny.Map({p3: d_716_v62_})
                d_720_v64_: _dafny.Array
                nw118_ = _dafny.Array(None, 21)
                nw118_[int(0)] = d_716_v62_
                nw118_[int(1)] = d_716_v62_
                nw118_[int(2)] = d_716_v62_
                nw118_[int(3)] = d_716_v62_
                nw118_[int(4)] = d_716_v62_
                nw118_[int(5)] = d_716_v62_
                nw118_[int(6)] = (d_716_v62_ if p0 else d_716_v62_)
                nw118_[int(7)] = (d_716_v62_ if (self).f21 else d_716_v62_)
                nw118_[int(8)] = d_716_v62_
                nw118_[int(9)] = d_716_v62_
                nw118_[int(10)] = d_716_v62_
                nw118_[int(11)] = d_716_v62_
                nw118_[int(12)] = d_716_v62_
                nw118_[int(13)] = d_716_v62_
                nw118_[int(14)] = d_716_v62_
                nw118_[int(15)] = d_716_v62_
                nw118_[int(16)] = d_716_v62_
                nw118_[int(17)] = d_716_v62_
                nw118_[int(18)] = d_716_v62_
                nw118_[int(19)] = d_716_v62_
                nw118_[int(20)] = ((d_719_v63_)[p1] if (p1) in (d_719_v63_) else d_716_v62_)
                d_720_v64_ = nw118_
                d_720_v64_ = (d_720_v64_ if p0 else d_720_v64_)
                d_721_v65_: T0
                nw119_ = C0()
                nw119_.ctor__(p2)
                d_721_v65_ = nw119_
                d_722_v66_: _dafny.Map
                d_722_v66_ = _dafny.Map({p2: d_721_v65_})
                d_723_v67_: D4
                d_723_v67_ = D4_DC12(((d_722_v66_)[(self).f21] if ((self).f21) in (d_722_v66_) else d_721_v65_))
                pat_let_tv24_ = d_721_v65_
                def iife52_(_pat_let13_0):
                    def iife53_(d_724_dt__update__tmp_h0_):
                        def iife54_(_pat_let14_0):
                            def iife55_(d_725_dt__update_hcf21_h0_):
                                return D4_DC12(d_725_dt__update_hcf21_h0_)
                            return iife55_(_pat_let14_0)
                        return iife54_(pat_let_tv24_)
                    return iife53_(_pat_let13_0)
                d_723_v67_ = (d_723_v67_ if p2 else iife52_(d_723_v67_))
                index84_ = default__.safeIndex(565, (d_710_v60_).length(0))
                (d_710_v60_)[index84_] = p0
                d_726_v68_: _dafny.Array
                nw120_ = _dafny.Array(None, 4)
                nw120_[int(0)] = d_650_v10_
                nw120_[int(1)] = p1
                nw120_[int(2)] = 9
                nw120_[int(3)] = p3
                d_726_v68_ = nw120_
                d_727_v69_: _dafny.Map
                d_727_v69_ = _dafny.Map({d_710_v60_: d_726_v68_})
                d_728_v70_: _dafny.Seq
                d_728_v70_ = _dafny.SeqWithoutIsStrInference([p1])
                d_729_v71_: _dafny.Set
                d_729_v71_ = _dafny.Set({True})
                d_730_v72_: _dafny.Array
                nw121_ = _dafny.Array(None, 24)
                nw121_[int(0)] = 505
                nw121_[int(1)] = p1
                nw121_[int(2)] = p3
                nw121_[int(3)] = len(d_727_v69_)
                nw121_[int(4)] = 499
                nw121_[int(5)] = len(d_728_v70_)
                nw121_[int(6)] = len(_dafny.SeqWithoutIsStrInference([d_715_v61_ for d_731_i11_ in range(default__.abs(634))]))
                nw121_[int(7)] = d_650_v10_
                nw121_[int(8)] = d_650_v10_
                nw121_[int(9)] = d_650_v10_
                nw121_[int(10)] = len(d_729_v71_)
                nw121_[int(11)] = p1
                nw121_[int(12)] = p1
                nw121_[int(13)] = p1
                nw121_[int(14)] = 235
                nw121_[int(15)] = default__.fm9(len(d_651_v11_), p3, True, globalState)
                nw121_[int(16)] = d_650_v10_
                nw121_[int(17)] = p3
                nw121_[int(18)] = 181
                nw121_[int(19)] = 772
                nw121_[int(20)] = len(d_651_v11_)
                nw121_[int(21)] = d_650_v10_
                nw121_[int(22)] = p3
                nw121_[int(23)] = -85
                d_730_v72_ = nw121_
                d_732_v73_: _dafny.Array
                nw122_ = _dafny.Array(None, 3)
                nw122_[int(0)] = d_730_v72_
                nw122_[int(1)] = d_726_v68_
                nw122_[int(2)] = d_730_v72_
                d_732_v73_ = nw122_
                index85_ = default__.safeIndex(565, (d_710_v60_).length(0))
                rhs82_ = d_732_v73_
                rhs83_ = p2
                lhs42_ = globalState
                lhs43_ = d_710_v60_
                lhs44_ = default__.safeIndex(565, (d_710_v60_).length(0))
                lhs42_.f2 = rhs82_
                lhs43_[lhs44_] = rhs83_
                d_733_v74_: _dafny.Map
                d_733_v74_ = _dafny.Map({(self).f21: p3})
                d_733_v74_ = (d_733_v74_).set(p2, p1)
            elif True:
                d_650_v10_ = (0) - (p3)
                (globalState).f11 = (d_651_v11_) != (d_651_v11_)
                (globalState).f17 = not((self).f21)
                d_734_v75_: str
                d_734_v75_ = _dafny.CodePoint('b')
                d_735_v76_: _dafny.Array
                def lambda47_(d_736_p1_):
                    def lambda48_(d_737_i12_):
                        return (d_737_i12_) * (d_736_p1_)

                    return lambda48_

                init23_ = lambda47_(p1)
                nw123_ = _dafny.Array(None, 2)
                for i0_23_ in range(nw123_.length(0)):
                    nw123_[i0_23_] = init23_(i0_23_)
                d_735_v76_ = nw123_
                index86_ = default__.safeIndex(944, (d_735_v76_).length(0))
                (d_735_v76_)[index86_] = (p1 if p0 else 185)
                d_738_v77_: _dafny.Set
                d_738_v77_ = _dafny.Set({p2})
                d_739_v78_: _dafny.Seq
                d_739_v78_ = _dafny.SeqWithoutIsStrInference([p0, p0, p2, p2, p0])
                index87_ = default__.safeIndex(944, (d_735_v76_).length(0))
                rhs84_ = False
                rhs85_ = _dafny.CodePoint('q')
                rhs86_ = d_650_v10_
                rhs87_ = default__.fm9((default__.fm9(len(d_651_v11_), p1, p2, globalState)) + (len(d_738_v77_)), default__.safeDivisionInt(d_650_v10_, 636), (p3) != (len(d_739_v78_)), globalState)
                rhs88_ = 958
                lhs45_ = globalState
                lhs46_ = d_735_v76_
                lhs47_ = default__.safeIndex(944, (d_735_v76_).length(0))
                lhs45_.f17 = rhs84_
                d_734_v75_ = rhs85_
                d_650_v10_ = rhs86_
                lhs46_[lhs47_] = rhs87_
                d_650_v10_ = rhs88_
                (globalState).f15 = p0
            d_740_v79_: _dafny.Array
            nw124_ = _dafny.Array(int(0), 18)
            d_740_v79_ = nw124_
            d_741_v80_: C1
            nw125_ = C1()
            nw125_.ctor__(d_740_v79_)
            d_741_v80_ = nw125_
            index88_ = default__.safeIndex(954, (d_710_v60_).length(0))
            (d_710_v60_)[index88_] = not((self).f21)
            index89_ = default__.safeIndex(954, (d_710_v60_).length(0))
            (d_710_v60_)[index89_] = not((self).f21)
        d_742_v81_: _dafny.MultiSet
        d_742_v81_ = _dafny.MultiSet([True])
        hi6_ = (d_742_v81_).cardinality
        for d_743_i13_ in range(-878, hi6_):
            d_744_v82_: _dafny.Map
            d_744_v82_ = _dafny.Map({p0: len(_dafny.Map({True: (self).f21}))})
            d_745_v83_: _dafny.Set
            d_745_v83_ = _dafny.Set({p3})
            d_746_v84_: str
            d_746_v84_ = _dafny.CodePoint('p')
            d_747_v85_: D4
            d_747_v85_ = D4_DC14((self).f21, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_748_i14_ in range(default__.abs(477))])), d_743_i13_, d_650_v10_])), d_745_v83_, d_746_v84_)
            d_749_v86_: _dafny.Map
            d_749_v86_ = _dafny.Map({((d_744_v82_)[p2] if (p2) in (d_744_v82_) else (d_747_v85_).cf24): 264})
            d_650_v10_ = (d_650_v10_) - (len(((d_749_v86_).set(603, p3)) | (d_749_v86_)))
            d_652_v12_ = (d_652_v12_).set(p1, (d_650_v10_) not in (d_749_v86_))
            d_750_v87_: _dafny.Set
            d_750_v87_ = _dafny.Set({p0})
            d_751_v88_: _dafny.Seq
            d_751_v88_ = _dafny.SeqWithoutIsStrInference([d_750_v87_])
            d_650_v10_ = (p1) * (len((d_751_v88_) + (d_751_v88_)))
            d_752_v89_: _dafny.Array
            nw126_ = _dafny.Array(_dafny.Array(None, 0), 16)
            d_752_v89_ = nw126_
            d_753_v90_: _dafny.Map
            d_753_v90_ = _dafny.Map({p2: (self).f21})
            d_754_v91_: _dafny.Map
            d_754_v91_ = _dafny.Map({_dafny.Map({d_743_i13_: p2}): False})
            d_755_v93_: _dafny.Seq
            d_755_v93_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({_dafny.CodePoint('k')}))])
            d_756_v94_: _dafny.Array
            nw127_ = _dafny.Array(None, 29)
            nw127_[int(0)] = p0
            nw127_[int(1)] = p2
            nw127_[int(2)] = p0
            nw127_[int(3)] = (self).f21
            nw127_[int(4)] = p0
            nw127_[int(5)] = p0
            nw127_[int(6)] = not((self).f21)
            nw127_[int(7)] = False
            nw127_[int(8)] = (self).f21
            nw127_[int(9)] = p0
            nw127_[int(10)] = False
            nw127_[int(11)] = not(p0)
            nw127_[int(12)] = (self).f21
            nw127_[int(13)] = p2
            nw127_[int(14)] = p0
            nw127_[int(15)] = p0
            nw127_[int(16)] = default__.fm0(p0, d_746_v84_, (self).f21, globalState)
            nw127_[int(17)] = default__.fm0((self).f21, d_746_v84_, (self).f21, globalState)
            nw127_[int(18)] = p2
            nw127_[int(19)] = p2
            nw127_[int(20)] = (self).f21
            def iife56_():
                coll26_ = _dafny.Map()
                compr_26_: int
                for compr_26_ in (d_755_v93_).Elements:
                    d_757_v92_: int = compr_26_
                    if (d_757_v92_) in (d_755_v93_):
                        coll26_[default__.safeModuloInt(d_757_v92_, d_743_i13_)] = (self).f21
                return _dafny.Map(coll26_)
            def iife57_():
                coll27_ = _dafny.Map()
                compr_27_: int
                for compr_27_ in (d_755_v93_).Elements:
                    d_758_v92_: int = compr_27_
                    if (d_758_v92_) in (d_755_v93_):
                        coll27_[default__.safeModuloInt(d_758_v92_, d_743_i13_)] = (self).f21
                return _dafny.Map(coll27_)
            nw127_[int(21)] = not(((d_753_v90_)[(self).f21] if ((self).f21) in (d_753_v90_) else ((d_754_v91_)[iife56_()
            ] if (iife57_()
            ) in (d_754_v91_) else p0)))
            nw127_[int(22)] = p2
            nw127_[int(23)] = True
            nw127_[int(24)] = p2
            nw127_[int(25)] = p0
            nw127_[int(26)] = p0
            nw127_[int(27)] = (self).f21
            nw127_[int(28)] = (self).f21
            d_756_v94_ = nw127_
            index90_ = default__.safeIndex(339, (d_752_v89_).length(0))
            (d_752_v89_)[index90_] = d_756_v94_
            index91_ = default__.safeIndex(339, (d_752_v89_).length(0))
            (d_752_v89_)[index91_] = (D1_DC4(d_650_v10_, (self).f21, d_756_v94_)).cf6
        r0 = p2
        d_759_v95_: _dafny.Map
        d_759_v95_ = _dafny.Map({p0: p2})
        d_760_v96_: str
        d_760_v96_ = _dafny.CodePoint('h')
        d_761_v97_: D3
        d_761_v97_ = D3_DC10(d_650_v10_)
        d_762_v98_: _dafny.Map
        d_762_v98_ = _dafny.Map({(self).f21: default__.fm11(d_650_v10_, ((d_759_v95_)[default__.fm0(p0, d_760_v96_, p2, globalState)] if (default__.fm0(p0, d_760_v96_, p2, globalState)) in (d_759_v95_) else (self).f21), p2, (d_761_v97_).cf17, globalState)})
        r1 = not(((0) - (p3)) > ((len(d_762_v98_)) - (p3)))
        return r0, r1

    def m5(self, p0, p1, globalState):
        r0: bool = False
        d_763_v0_: D4
        d_763_v0_ = D4_DC13(len(_dafny.SeqWithoutIsStrInference([(self).f21])))
        d_764_v1_: int
        d_764_v1_ = 992
        d_765_v2_: str
        d_765_v2_ = _dafny.CodePoint('e')
        d_766_v3_: D4
        d_766_v3_ = D4_DC14(False, (0) - (d_764_v1_), _dafny.Set({d_764_v1_}), d_765_v2_)
        d_767_v4_: _dafny.Seq
        d_767_v4_ = _dafny.SeqWithoutIsStrInference([d_764_v1_])
        d_768_v5_: D4
        d_768_v5_ = D4_DC14(not((self).f21), d_764_v1_, _dafny.Set({d_764_v1_, d_764_v1_, (d_766_v3_).cf24, len(default__.fm19(d_767_v4_, True, globalState))}), d_765_v2_)
        pat_let_tv25_ = d_763_v0_
        pat_let_tv26_ = d_764_v1_
        pat_let_tv27_ = d_763_v0_
        pat_let_tv28_ = d_763_v0_
        def lambda49_(source10_):
            if source10_.is_DC13:
                d_769___mcc_h0_ = source10_.cf22
                d_770_cf22_ = d_769___mcc_h0_
                d_771_dt__update__tmp_h0_ = pat_let_tv25_
                d_772_dt__update_hcf22_h0_ = pat_let_tv26_
                return D4_DC13(d_772_dt__update_hcf22_h0_)
            elif source10_.is_DC14:
                d_773___mcc_h1_ = source10_.cf23
                d_774___mcc_h2_ = source10_.cf24
                d_775___mcc_h3_ = source10_.cf25
                d_776___mcc_h4_ = source10_.cf26
                d_777_cf26_ = d_776___mcc_h4_
                d_778_cf25_ = d_775___mcc_h3_
                d_779_cf24_ = d_774___mcc_h2_
                d_780_cf23_ = d_773___mcc_h1_
                return D4_DC13(d_779_cf24_)
            elif source10_.is_DC12:
                d_781___mcc_h5_ = source10_.cf21
                d_782_cf21_ = d_781___mcc_h5_
                return pat_let_tv27_
            elif True:
                d_783___mcc_h6_ = source10_.cf27
                d_784_cf27_ = d_783___mcc_h6_
                return pat_let_tv28_

        d_763_v0_ = lambda49_(d_766_v3_)
        if (self).f21:
            d_785_v6_: _dafny.Set
            d_785_v6_ = _dafny.Set({len(p1), d_764_v1_, d_764_v1_})
            d_786_v7_: _dafny.Map
            d_786_v7_ = _dafny.Map({p0: not((d_785_v6_).ispropersubset(d_785_v6_))})
            d_764_v1_ = len(d_786_v7_)
            d_787_v9_: _dafny.Array
            nw128_ = _dafny.Array(False, 28)
            d_787_v9_ = nw128_
            d_788_v10_: D1
            d_788_v10_ = D1_DC4((0) - (d_764_v1_), False, d_787_v9_)
            d_789_v11_: _dafny.Map
            d_789_v11_ = _dafny.Map({d_764_v1_: (d_788_v10_).cf5})
            d_790_v12_: D3
            def iife58_():
                coll28_ = _dafny.Map()
                compr_28_: int
                for compr_28_ in _dafny.IntegerRange(527, 468):
                    d_791_v8_: int = compr_28_
                    if ((527) <= (d_791_v8_)) and ((d_791_v8_) < (468)):
                        coll28_[(d_791_v8_) + (d_764_v1_)] = (self).f21
                return _dafny.Map(coll28_)
            d_790_v12_ = D3_DC9(_dafny.Set({iife58_()
, d_789_v11_}))
            d_792_v13_: _dafny.Set
            d_792_v13_ = _dafny.Set({d_789_v11_, d_789_v11_})
            d_790_v12_ = D3_DC9(d_792_v13_)
            d_793_v14_: _dafny.Array
            nw129_ = _dafny.Array(None, 5)
            nw129_[int(0)] = d_764_v1_
            nw129_[int(1)] = 418
            nw129_[int(2)] = d_764_v1_
            nw129_[int(3)] = d_764_v1_
            nw129_[int(4)] = d_764_v1_
            d_793_v14_ = nw129_
            d_794_v15_: C1
            nw130_ = C1()
            nw130_.ctor__(d_793_v14_)
            d_794_v15_ = nw130_
            d_795_v16_: _dafny.Map
            d_795_v16_ = _dafny.Map({d_764_v1_: d_764_v1_})
            d_796_v17_: _dafny.Seq
            d_796_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "csvwaexvd"))
            d_789_v11_ = (d_789_v11_).set((0) - (((d_795_v16_)[d_764_v1_] if (d_764_v1_) in (d_795_v16_) else len(d_796_v17_))), not(p0))
            d_797_v18_: _dafny.Set
            d_797_v18_ = _dafny.Set({p0})
            d_797_v18_ = d_797_v18_
        elif True:
            d_798_v19_: _dafny.Map
            d_798_v19_ = _dafny.Map({d_764_v1_: len((_dafny.Map({p0: d_764_v1_})).set(p0, d_764_v1_))})
            d_799_v20_: _dafny.Array
            nw131_ = _dafny.Array(_dafny.Seq({}), 4)
            d_799_v20_ = nw131_
            d_800_v21_: _dafny.Seq
            d_800_v21_ = _dafny.SeqWithoutIsStrInference([d_799_v20_])
            d_801_v22_: D1
            d_801_v22_ = D1_DC3((d_800_v21_)[default__.safeIndex(len(default__.fm20(globalState)), len(d_800_v21_))])
            d_802_v23_: _dafny.MultiSet
            d_802_v23_ = _dafny.MultiSet([d_801_v22_])
            d_798_v19_ = (d_798_v19_).set(default__.safeModuloInt(d_764_v1_, d_764_v1_), ((d_802_v23_).cardinality) - (d_764_v1_))
            d_764_v1_ = ((0) - (d_764_v1_)) + ((d_764_v1_) - (len(p1)))
            d_803_v24_: _dafny.Seq
            d_803_v24_ = _dafny.SeqWithoutIsStrInference([d_765_v2_, d_765_v2_, d_765_v2_])
            d_804_v25_: T0
            nw132_ = C0()
            nw132_.ctor__(((d_803_v24_)[default__.safeIndex(d_764_v1_, len(d_803_v24_))]) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgaehayk"))))
            d_804_v25_ = nw132_
            d_804_v25_ = d_804_v25_
            d_803_v24_ = (d_803_v24_) + (d_803_v24_)
            rhs89_ = default__.safeDivisionInt(default__.fm9(d_764_v1_, len(p1), (self).f21, globalState), default__.fm9(d_764_v1_, (d_767_v4_)[default__.safeIndex(d_764_v1_, len(d_767_v4_))], p0, globalState))
            rhs90_ = (d_803_v24_ if p0 else d_803_v24_)
            d_764_v1_ = rhs89_
            d_803_v24_ = rhs90_
        d_805_i0_: int
        d_805_i0_ = 0
        with _dafny.label("2"):
            while (self).f21:
                with _dafny.c_label("2"):
                    if (d_805_i0_) >= (100):
                        raise _dafny.Break("2")
                    d_805_i0_ = (d_805_i0_) + (1)
                    if ((self).f21 if p0 else not((self).f21)):
                        (globalState).f17 = not((self).f21)
                        d_806_v26_: _dafny.Array
                        def lambda50_(d_807_p0_):
                            def lambda51_(d_808_i1_):
                                return d_807_p0_

                            return lambda51_

                        init24_ = lambda50_(p0)
                        nw133_ = _dafny.Array(None, 23)
                        for i0_24_ in range(nw133_.length(0)):
                            nw133_[i0_24_] = init24_(i0_24_)
                        d_806_v26_ = nw133_
                        d_809_v27_: _dafny.Map
                        d_809_v27_ = _dafny.Map({d_806_v26_: default__.fm9((0) - (d_764_v1_), len(_dafny.SeqWithoutIsStrInference([p0])), p0, globalState)})
                        d_809_v27_ = (d_809_v27_).set(d_806_v26_, (d_764_v1_) - (d_764_v1_))
                        d_810_v28_: _dafny.Array
                        nw134_ = _dafny.Array(int(0), 16)
                        d_810_v28_ = nw134_
                        d_810_v28_ = d_810_v28_
                        d_811_v29_: _dafny.Set
                        d_811_v29_ = _dafny.Set({len(d_767_v4_), d_764_v1_})
                        d_764_v1_ = default__.fm9(len((_dafny.Set({d_764_v1_})) | (d_811_v29_)), default__.safeDivisionInt(d_764_v1_, d_764_v1_), not((self).f21), globalState)
                        (globalState).f15 = (self).f21
                    elif True:
                        d_812_v30_: _dafny.Array
                        nw135_ = _dafny.Array(False, 11)
                        d_812_v30_ = nw135_
                        index92_ = default__.safeIndex(241, (d_812_v30_).length(0))
                        (d_812_v30_)[index92_] = p0
                        d_813_v31_: _dafny.MultiSet
                        d_813_v31_ = _dafny.MultiSet([d_812_v30_])
                        index93_ = default__.safeIndex(241, (d_812_v30_).length(0))
                        (d_812_v30_)[index93_] = default__.fm0(not((self).f21), d_765_v2_, (d_813_v31_).issubset(d_813_v31_), globalState)
                        d_814_v32_: _dafny.Seq
                        d_814_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cb"))
                        d_814_v32_ = (d_814_v32_) + ((d_814_v32_) + (d_814_v32_))
                        d_764_v1_ = d_764_v1_
                        d_815_v33_: C0
                        nw136_ = C0()
                        nw136_.ctor__((self).f21)
                        d_815_v33_ = nw136_
                        d_816_v34_: _dafny.Array
                        nw137_ = _dafny.Array(int(0), 28)
                        d_816_v34_ = nw137_
                        index94_ = default__.safeIndex(653, (d_816_v34_).length(0))
                        (d_816_v34_)[index94_] = d_764_v1_
                        index95_ = default__.safeIndex(653, (d_816_v34_).length(0))
                        (d_816_v34_)[index95_] = d_764_v1_
                    if (p0) and ((self).f21):
                        d_817_v35_: C0
                        nw138_ = C0()
                        nw138_.ctor__((self).f21)
                        d_817_v35_ = nw138_
                        (globalState).f17 = (d_817_v35_).f23
                        d_765_v2_ = d_765_v2_
                        d_818_v36_: _dafny.Array
                        nw139_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 11)
                        d_818_v36_ = nw139_
                        d_819_v37_: _dafny.Seq
                        d_819_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbewbhblg"))
                        index96_ = default__.safeIndex(779, (d_818_v36_).length(0))
                        (d_818_v36_)[index96_] = d_819_v37_
                        index97_ = default__.safeIndex(779, (d_818_v36_).length(0))
                        (d_818_v36_)[index97_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jksnhhduo"))
                        d_820_v38_: _dafny.Set
                        d_820_v38_ = _dafny.Set({478})
                        def iife59_():
                            coll29_ = _dafny.Set()
                            compr_29_: int
                            for compr_29_ in _dafny.IntegerRange(-957, 919):
                                d_821_v39_: int = compr_29_
                                if ((-957) <= (d_821_v39_)) and ((d_821_v39_) < (919)):
                                    coll29_ = coll29_.union(_dafny.Set([(d_821_v39_) * (d_764_v1_)]))
                            return _dafny.Set(coll29_)
                        d_764_v1_ = len(((_dafny.Set({len((d_818_v36_)[default__.safeIndex(779, (d_818_v36_).length(0))]), (0) - ((0) - (d_764_v1_))})) - (d_820_v38_)).intersection(iife59_()
                        ))
                    elif True:
                        d_822_v40_: _dafny.Seq
                        d_822_v40_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))
                        d_823_v41_: D1
                        d_823_v41_ = D1_DC5(d_822_v40_, (0) - (d_764_v1_), d_764_v1_)
                        d_764_v1_ = len((_dafny.SeqWithoutIsStrInference([D1_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dyabgd")), d_764_v1_, d_764_v1_) for d_824_i2_ in range(default__.abs(165))]) if p0 else _dafny.SeqWithoutIsStrInference([d_823_v41_, d_823_v41_])))
                        d_764_v1_ = (len(d_767_v4_)) * (d_764_v1_)
                        d_825_v42_: C0
                        nw140_ = C0()
                        nw140_.ctor__((len((d_822_v40_).set(default__.safeIndex(d_764_v1_, len(d_822_v40_)), d_765_v2_))) < (d_764_v1_))
                        d_825_v42_ = nw140_
                        d_826_v43_: _dafny.Seq
                        d_826_v43_ = _dafny.SeqWithoutIsStrInference([d_822_v40_])
                        d_764_v1_ = len(((d_822_v40_) + (((d_826_v43_)[default__.safeIndex(d_764_v1_, len(d_826_v43_))]).set(default__.safeIndex(d_764_v1_, len((d_826_v43_)[default__.safeIndex(d_764_v1_, len(d_826_v43_))])), d_765_v2_)) if (d_825_v42_).f23 else d_822_v40_))
                        rhs91_ = default__.safeDivisionInt((0) - (d_764_v1_), len((self).fm8(p0, globalState)))
                        rhs92_ = (p1)[default__.safeIndex(518, len(p1))]
                        rhs93_ = (-637 if (d_825_v42_).f23 else d_764_v1_)
                        lhs48_ = globalState
                        d_764_v1_ = rhs91_
                        lhs48_.f15 = rhs92_
                        d_764_v1_ = rhs93_
                    d_827_v44_: _dafny.Array
                    def lambda52_(d_828_i3_):
                        return (self).f21

                    init25_ = lambda52_
                    nw141_ = _dafny.Array(None, 22)
                    for i0_25_ in range(nw141_.length(0)):
                        nw141_[i0_25_] = init25_(i0_25_)
                    d_827_v44_ = nw141_
                    d_827_v44_ = d_827_v44_
                    (globalState).f11 = (self).f21
                    pass
            pass
        d_829_i4_: int
        d_829_i4_ = 0
        with _dafny.label("3"):
            while (d_764_v1_) != ((_dafny.MultiSet([d_764_v1_, d_764_v1_, d_764_v1_])).cardinality):
                with _dafny.c_label("3"):
                    if (d_829_i4_) >= (100):
                        raise _dafny.Break("3")
                    d_829_i4_ = (d_829_i4_) + (1)
                    (globalState).f17 = (self).f21
                    d_830_v45_: _dafny.Array
                    def lambda53_(d_831_v4_):
                        def lambda54_(d_832_i5_):
                            return d_831_v4_

                        return lambda54_

                    init26_ = lambda53_(d_767_v4_)
                    nw142_ = _dafny.Array(None, 13)
                    for i0_26_ in range(nw142_.length(0)):
                        nw142_[i0_26_] = init26_(i0_26_)
                    d_830_v45_ = nw142_
                    index98_ = default__.safeIndex(158, (d_830_v45_).length(0))
                    (d_830_v45_)[index98_] = default__.fm19(d_767_v4_, p0, globalState)
                    index99_ = default__.safeIndex(158, (d_830_v45_).length(0))
                    (d_830_v45_)[index99_] = (d_767_v4_) + (d_767_v4_)
                    rhs94_ = (d_764_v1_) not in ((d_830_v45_)[default__.safeIndex(158, (d_830_v45_).length(0))])
                    rhs95_ = 459
                    lhs49_ = globalState
                    lhs49_.f11 = rhs94_
                    d_764_v1_ = rhs95_
                    d_764_v1_ = default__.fm9(d_764_v1_, (d_764_v1_ if (self).f21 else d_764_v1_), p0, globalState)
                    pass
            pass
        d_833_v46_: _dafny.Array
        nw143_ = _dafny.Array(_dafny.Seq({}), 19)
        d_833_v46_ = nw143_
        d_834_v47_: D1
        d_834_v47_ = D1_DC3(d_833_v46_)
        d_835_v48_: _dafny.Array
        nw144_ = _dafny.Array(None, 7)
        nw144_[int(0)] = d_833_v46_
        nw144_[int(1)] = d_833_v46_
        nw144_[int(2)] = d_833_v46_
        nw144_[int(3)] = d_833_v46_
        nw144_[int(4)] = d_833_v46_
        nw144_[int(5)] = d_833_v46_
        nw144_[int(6)] = (d_834_v47_).cf3
        d_835_v48_ = nw144_
        index100_ = default__.safeIndex(649, (d_835_v48_).length(0))
        (d_835_v48_)[index100_] = d_833_v46_
        index101_ = default__.safeIndex(649, (d_835_v48_).length(0))
        (d_835_v48_)[index101_] = d_833_v46_
        if (self).f21:
            d_836_v49_: _dafny.Seq
            d_836_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trlsw"))
            d_764_v1_ = (len(d_836_v49_)) + (len((d_767_v4_) + (d_767_v4_)))
            if not((p1)[default__.safeIndex(default__.fm9(d_764_v1_, len(p1), not((self).f21), globalState), len(p1))]):
                d_837_v50_: _dafny.Array
                nw145_ = _dafny.Array(False, 5)
                d_837_v50_ = nw145_
                index102_ = default__.safeIndex(474, (d_837_v50_).length(0))
                (d_837_v50_)[index102_] = p0
                index103_ = default__.safeIndex(474, (d_837_v50_).length(0))
                rhs96_ = (self).f21
                rhs97_ = d_764_v1_
                lhs50_ = d_837_v50_
                lhs51_ = default__.safeIndex(474, (d_837_v50_).length(0))
                lhs50_[lhs51_] = rhs96_
                d_764_v1_ = rhs97_
                d_764_v1_ = (0) - (d_764_v1_)
                d_837_v50_ = d_837_v50_
                d_838_v51_: _dafny.MultiSet
                d_838_v51_ = _dafny.MultiSet([p0, (d_837_v50_)[default__.safeIndex(474, (d_837_v50_).length(0))], (self).f21])
                index104_ = default__.safeIndex(474, (d_837_v50_).length(0))
                (d_837_v50_)[index104_] = (d_838_v51_).issubset(d_838_v51_)
                d_764_v1_ = default__.safeModuloInt(468, d_764_v1_)
            elif True:
                d_836_v49_ = (default__.fm21(globalState)).cf7
                d_839_v52_: _dafny.Array
                nw146_ = _dafny.Array(int(0), 2)
                d_839_v52_ = nw146_
                d_840_v53_: _dafny.MultiSet
                d_840_v53_ = _dafny.MultiSet([d_839_v52_])
                (globalState).f15 = ((self).f21 if (_dafny.MultiSet([d_839_v52_, d_839_v52_])).issubset(d_840_v53_) else (self).f21)
                d_841_v54_: _dafny.Set
                d_841_v54_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p0, p0])})
                d_842_v55_: _dafny.MultiSet
                d_842_v55_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qn"))])
                d_843_v56_: _dafny.Seq
                d_843_v56_ = _dafny.SeqWithoutIsStrInference([d_842_v55_])
                d_844_v57_: _dafny.Array
                nw147_ = _dafny.Array(None, 20)
                nw147_[int(0)] = p0
                nw147_[int(1)] = (self).f21
                nw147_[int(2)] = False
                nw147_[int(3)] = (d_841_v54_).ispropersubset(d_841_v54_)
                nw147_[int(4)] = default__.fm0((self).f21, d_765_v2_, (self).f21, globalState)
                nw147_[int(5)] = p0
                nw147_[int(6)] = p0
                nw147_[int(7)] = (self).f21
                nw147_[int(8)] = (d_842_v55_).isdisjoint((d_843_v56_)[default__.safeIndex(d_764_v1_, len(d_843_v56_))])
                nw147_[int(9)] = not(not(False))
                nw147_[int(10)] = (self).f21
                nw147_[int(11)] = p0
                nw147_[int(12)] = p0
                nw147_[int(13)] = (self).f21
                nw147_[int(14)] = True
                nw147_[int(15)] = p0
                nw147_[int(16)] = p0
                nw147_[int(17)] = not((self).f21)
                nw147_[int(18)] = default__.fm0((self).f21, d_765_v2_, p0, globalState)
                nw147_[int(19)] = True
                d_844_v57_ = nw147_
                index105_ = default__.safeIndex(792, (d_844_v57_).length(0))
                (d_844_v57_)[index105_] = p0
                index106_ = default__.safeIndex(792, (d_844_v57_).length(0))
                (d_844_v57_)[index106_] = ((0) - (d_764_v1_)) >= ((d_764_v1_) + (-440))
                d_845_v58_: _dafny.Map
                d_845_v58_ = _dafny.Map({(self).f21: p0})
                d_846_v59_: _dafny.Map
                d_846_v59_ = _dafny.Map({True: (d_845_v58_) | ((d_845_v58_).set((d_844_v57_)[default__.safeIndex(792, (d_844_v57_).length(0))], (self).f21))})
                d_846_v59_ = (d_846_v59_).set((self).f21, (default__.fm22(globalState)).set(p0, (self).f21))
                d_847_v60_: _dafny.Map
                d_847_v60_ = _dafny.Map({d_764_v1_: len(d_836_v49_)})
                index107_ = default__.safeIndex(792, (d_844_v57_).length(0))
                (d_844_v57_)[index107_] = not((d_764_v1_) not in (d_847_v60_))
            (globalState).f15 = default__.fm0((d_764_v1_) != (d_764_v1_), d_765_v2_, (self).f21, globalState)
            d_848_v61_: _dafny.Map
            d_848_v61_ = _dafny.Map({p1: d_764_v1_})
            d_848_v61_ = (d_848_v61_).set((p1) + (p1), (d_764_v1_) + (d_764_v1_))
            d_849_v62_: _dafny.MultiSet
            d_849_v62_ = _dafny.MultiSet([d_764_v1_])
            d_850_v63_: _dafny.Set
            d_850_v63_ = _dafny.Set({d_849_v62_, _dafny.MultiSet(d_767_v4_), d_849_v62_, d_849_v62_})
            d_851_v64_: _dafny.Seq
            d_851_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_849_v62_, d_849_v62_}), d_850_v63_, d_850_v63_, d_850_v63_, d_850_v63_])
            d_852_v65_: _dafny.MultiSet
            d_852_v65_ = _dafny.MultiSet([((d_851_v64_)[default__.safeIndex(d_764_v1_, len(d_851_v64_))]).ispropersubset(d_850_v63_), (self).f21, p0, (self).f21])
            d_852_v65_ = d_852_v65_
        elif True:
            if ((default__.fm9(d_764_v1_, d_764_v1_, (self).f21, globalState)) - (d_764_v1_)) != ((d_764_v1_ if (self).f21 else d_764_v1_)):
                d_853_v66_: _dafny.Seq
                d_853_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "upxiueqmn"))
                d_854_v67_: _dafny.Map
                d_854_v67_ = _dafny.Map({d_768_v5_: d_853_v66_})
                d_854_v67_ = (d_854_v67_).set(d_768_v5_, _dafny.SeqWithoutIsStrInference([d_765_v2_ for d_855_i6_ in range(default__.abs(581))]))
                d_856_v68_: _dafny.Map
                d_856_v68_ = _dafny.Map({(self).f21: d_764_v1_})
                d_857_v69_: _dafny.MultiSet
                d_857_v69_ = _dafny.MultiSet([d_764_v1_])
                d_858_v70_: _dafny.Seq
                d_858_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_764_v1_]), _dafny.MultiSet([len(d_853_v66_)]), d_857_v69_, d_857_v69_, _dafny.MultiSet(d_767_v4_)])
                d_856_v68_ = (d_856_v68_).set((default__.fm23((self).f21, d_856_v68_, default__.fm9(d_764_v1_, d_764_v1_, p0, globalState), d_765_v2_, globalState)).ispropersubset((d_858_v70_)[default__.safeIndex((0) - ((d_857_v69_).cardinality), len(d_858_v70_))]), d_764_v1_)
                (globalState).f11 = default__.fm0((self).f21, _dafny.CodePoint('x'), False, globalState)
                (globalState).f11 = p0
                d_859_v71_: _dafny.Array
                def lambda55_(d_860_v1_, d_861_v4_):
                    def lambda56_(d_862_i7_):
                        return _dafny.Map({len(_dafny.Set({_dafny.SeqWithoutIsStrInference([d_860_v1_]), d_861_v4_})): d_860_v1_})

                    return lambda56_

                init27_ = lambda55_(d_764_v1_, d_767_v4_)
                nw148_ = _dafny.Array(None, 16)
                for i0_27_ in range(nw148_.length(0)):
                    nw148_[i0_27_] = init27_(i0_27_)
                d_859_v71_ = nw148_
                d_863_v72_: _dafny.Map
                d_863_v72_ = _dafny.Map({d_764_v1_: (d_857_v69_).cardinality})
                index108_ = default__.safeIndex(828, (d_859_v71_).length(0))
                (d_859_v71_)[index108_] = d_863_v72_
                index109_ = default__.safeIndex(828, (d_859_v71_).length(0))
                (d_859_v71_)[index109_] = d_863_v72_
            elif True:
                d_864_v73_: _dafny.Array
                nw149_ = _dafny.Array(False, 22)
                d_864_v73_ = nw149_
                d_865_v74_: _dafny.Set
                d_865_v74_ = _dafny.Set({d_864_v73_})
                d_866_v75_: _dafny.Seq
                d_866_v75_ = _dafny.SeqWithoutIsStrInference([d_864_v73_, d_864_v73_])
                d_867_v76_: _dafny.Array
                nw150_ = _dafny.Array(None, 16)
                nw150_[int(0)] = d_865_v74_
                nw150_[int(1)] = (d_865_v74_).intersection(d_865_v74_)
                nw150_[int(2)] = d_865_v74_
                nw150_[int(3)] = d_865_v74_
                nw150_[int(4)] = d_865_v74_
                nw150_[int(5)] = d_865_v74_
                nw150_[int(6)] = _dafny.Set({d_864_v73_, d_864_v73_, d_864_v73_, d_864_v73_, (d_866_v75_)[default__.safeIndex(d_764_v1_, len(d_866_v75_))]})
                nw150_[int(7)] = d_865_v74_
                nw150_[int(8)] = (d_865_v74_) - (d_865_v74_)
                nw150_[int(9)] = d_865_v74_
                nw150_[int(10)] = d_865_v74_
                nw150_[int(11)] = _dafny.Set({d_864_v73_})
                nw150_[int(12)] = d_865_v74_
                nw150_[int(13)] = d_865_v74_
                nw150_[int(14)] = d_865_v74_
                nw150_[int(15)] = _dafny.Set({d_864_v73_})
                d_867_v76_ = nw150_
                index110_ = default__.safeIndex(518, (d_867_v76_).length(0))
                (d_867_v76_)[index110_] = d_865_v74_
                index111_ = default__.safeIndex(518, (d_867_v76_).length(0))
                (d_867_v76_)[index111_] = d_865_v74_
                d_764_v1_ = d_764_v1_
                d_868_v77_: C0
                nw151_ = C0()
                nw151_.ctor__(False)
                d_868_v77_ = nw151_
                index112_ = default__.safeIndex(427, (d_864_v73_).length(0))
                (d_864_v73_)[index112_] = (self).f21
                index113_ = default__.safeIndex(229, (d_864_v73_).length(0))
                (d_864_v73_)[index113_] = True
                d_869_v78_: _dafny.Array
                nw152_ = _dafny.Array(int(0), 24)
                d_869_v78_ = nw152_
                d_870_v79_: _dafny.MultiSet
                d_870_v79_ = _dafny.MultiSet([d_764_v1_, d_764_v1_])
                d_871_v80_: _dafny.Seq
                d_871_v80_ = _dafny.SeqWithoutIsStrInference([d_870_v79_, d_870_v79_])
                d_872_v81_: _dafny.Map
                d_872_v81_ = _dafny.Map({d_764_v1_: p0})
                d_873_v82_: _dafny.Set
                d_873_v82_ = _dafny.Set({d_764_v1_, ((d_870_v79_).set((d_767_v4_)[default__.safeIndex(len(d_767_v4_), len(d_767_v4_))], default__.abs(d_764_v1_))).cardinality, ((d_870_v79_).intersection((d_871_v80_)[default__.safeIndex(len(d_872_v81_), len(d_871_v80_))])).cardinality})
                d_874_v83_: _dafny.MultiSet
                d_874_v83_ = _dafny.MultiSet([p1, p1, p1])
                d_875_v84_: _dafny.Map
                d_875_v84_ = _dafny.Map({d_874_v83_: d_764_v1_})
                d_876_v85_: _dafny.Map
                d_876_v85_ = _dafny.Map({(self).f21: d_764_v1_})
                index114_ = default__.safeIndex(427, (d_864_v73_).length(0))
                index115_ = default__.safeIndex(229, (d_864_v73_).length(0))
                def iife60_():
                    coll30_ = _dafny.Set()
                    compr_30_: int
                    for compr_30_ in _dafny.IntegerRange(961, 442):
                        d_877_v86_: int = compr_30_
                        if ((961) <= (d_877_v86_)) and ((d_877_v86_) < (442)):
                            coll30_ = coll30_.union(_dafny.Set([(d_877_v86_) * (d_764_v1_)]))
                    return _dafny.Set(coll30_)
                rhs98_ = p0
                rhs99_ = default__.fm0(not((d_868_v77_).fm15((d_868_v77_).f23, p0, ((d_875_v84_)[(d_874_v83_).set(_dafny.SeqWithoutIsStrInference([p0]), default__.abs(d_764_v1_))] if ((d_874_v83_).set(_dafny.SeqWithoutIsStrInference([p0]), default__.abs(d_764_v1_))) in (d_875_v84_) else d_764_v1_), default__.fm24((d_868_v77_).fm15(p0, p0, len(d_767_v4_), d_876_v85_, globalState), len(d_873_v82_), globalState), globalState)), (d_768_v5_).cf26, True, globalState)
                rhs100_ = d_869_v78_
                rhs101_ = (d_873_v82_) | ((iife60_()
                ) - (_dafny.Set({d_764_v1_, 282, d_764_v1_})))
                lhs52_ = d_864_v73_
                lhs53_ = default__.safeIndex(427, (d_864_v73_).length(0))
                lhs54_ = d_864_v73_
                lhs55_ = default__.safeIndex(229, (d_864_v73_).length(0))
                lhs52_[lhs53_] = rhs98_
                lhs54_[lhs55_] = rhs99_
                d_869_v78_ = rhs100_
                d_873_v82_ = rhs101_
                d_878_v87_: _dafny.Map
                d_878_v87_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_765_v2_]): d_764_v1_})
                d_879_v88_: _dafny.Seq
                d_879_v88_ = _dafny.SeqWithoutIsStrInference([d_765_v2_, d_765_v2_])
                d_878_v87_ = (d_878_v87_).set(d_879_v88_, d_764_v1_)
            d_880_v89_: _dafny.Map
            d_880_v89_ = _dafny.Map({d_764_v1_: (self).f21})
            d_881_v90_: _dafny.Seq
            d_881_v90_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_764_v1_: (self).f21})])
            d_882_v91_: _dafny.Set
            d_882_v91_ = _dafny.Set({d_880_v89_, (d_881_v90_)[default__.safeIndex(d_764_v1_, len(d_881_v90_))]})
            d_883_v92_: D3
            d_883_v92_ = D3_DC9((d_882_v91_) - (d_882_v91_))
            d_883_v92_ = d_883_v92_
            d_884_v93_: C0
            nw153_ = C0()
            nw153_.ctor__((self).f21)
            d_884_v93_ = nw153_
            d_885_v94_: _dafny.Map
            d_885_v94_ = _dafny.Map({p0: (0) - (d_764_v1_)})
            d_886_v95_: _dafny.Array
            nw154_ = _dafny.Array(None, 12)
            nw154_[int(0)] = (self).f21
            nw154_[int(1)] = p0
            nw154_[int(2)] = p0
            nw154_[int(3)] = ((d_884_v93_).fm14((0) - (d_764_v1_), globalState)) == (d_764_v1_)
            nw154_[int(4)] = p0
            nw154_[int(5)] = (self).f21
            nw154_[int(6)] = p0
            nw154_[int(7)] = False
            nw154_[int(8)] = (self).f21
            nw154_[int(9)] = (self).f21
            nw154_[int(10)] = (d_884_v93_).f23
            nw154_[int(11)] = (d_884_v93_).fm15(not((d_884_v93_).f23), not((d_884_v93_).f23), (d_884_v93_).fm14(d_764_v1_, globalState), d_885_v94_, globalState)
            d_886_v95_ = nw154_
            d_886_v95_ = (d_886_v95_ if (self).f21 else (d_886_v95_ if p0 else d_886_v95_))
            source11_ = self.f20
            if source11_.is_DC1:
                d_887___mcc_h7_ = source11_.cf1
                d_888_cf1_ = d_887___mcc_h7_
                index116_ = default__.safeIndex(956, (d_886_v95_).length(0))
                (d_886_v95_)[index116_] = p0
                index117_ = default__.safeIndex(956, (d_886_v95_).length(0))
                (d_886_v95_)[index117_] = (d_884_v93_).fm15((d_884_v93_).f23, ((self).f21) or ((d_884_v93_).f23), 395, d_885_v94_, globalState)
                d_889_v96_: _dafny.Seq
                d_889_v96_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aln"))
                d_890_v97_: _dafny.Map
                d_890_v97_ = _dafny.Map({d_764_v1_: d_889_v96_})
                (globalState).f15 = (d_889_v96_) <= (((d_890_v97_)[d_764_v1_] if (d_764_v1_) in (d_890_v97_) else d_889_v96_))
                d_891_v98_: _dafny.Array
                nw155_ = _dafny.Array(_dafny.Seq({}), 8)
                d_891_v98_ = nw155_
                d_891_v98_ = d_891_v98_
                d_892_v99_: _dafny.Array
                nw156_ = _dafny.Array(int(0), 10)
                d_892_v99_ = nw156_
                index118_ = default__.safeIndex(334, (d_892_v99_).length(0))
                (d_892_v99_)[index118_] = d_764_v1_
                d_893_v100_: _dafny.MultiSet
                d_893_v100_ = _dafny.MultiSet([d_764_v1_, d_764_v1_, d_764_v1_])
                index119_ = default__.safeIndex(334, (d_892_v99_).length(0))
                rhs102_ = default__.safeModuloInt(d_764_v1_, d_764_v1_)
                rhs103_ = d_764_v1_
                rhs104_ = (d_893_v100_) - (d_893_v100_)
                lhs56_ = d_892_v99_
                lhs57_ = default__.safeIndex(334, (d_892_v99_).length(0))
                lhs56_[lhs57_] = rhs102_
                d_764_v1_ = rhs103_
                d_893_v100_ = rhs104_
            elif source11_.is_DC0:
                d_894___mcc_h8_ = source11_.cf0
                d_895_cf0_ = d_894___mcc_h8_
                d_896_v101_: _dafny.Array
                nw157_ = _dafny.Array(_dafny.Array(None, 0), 25)
                d_896_v101_ = nw157_
                d_897_v102_: _dafny.Seq
                d_897_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uajufq"))
                d_898_v103_: _dafny.Array
                nw158_ = _dafny.Array(int(0), 18)
                d_898_v103_ = nw158_
                d_899_v104_: _dafny.Map
                d_899_v104_ = _dafny.Map({d_897_v102_: d_898_v103_})
                index120_ = default__.safeIndex(686, (d_896_v101_).length(0))
                (d_896_v101_)[index120_] = ((d_899_v104_)[d_897_v102_] if (d_897_v102_) in (d_899_v104_) else d_898_v103_)
                index121_ = default__.safeIndex(686, (d_896_v101_).length(0))
                (d_896_v101_)[index121_] = d_898_v103_
                (globalState).f17 = (d_768_v5_).cf23
                (globalState).f15 = (d_884_v93_).f23
                d_764_v1_ = d_764_v1_
            elif True:
                d_900___mcc_h9_ = source11_.cf2
                d_901_cf2_ = d_900___mcc_h9_
                r0 = not(((self).f21) and (not((d_884_v93_).f23)))
                d_902_v105_: _dafny.Array
                nw159_ = _dafny.Array(int(0), 22)
                d_902_v105_ = nw159_
                d_903_v106_: D3
                d_903_v106_ = D3_DC11(d_885_v94_, d_764_v1_, d_902_v105_)
                d_904_v107_: _dafny.MultiSet
                d_904_v107_ = _dafny.MultiSet([D3_DC11(d_885_v94_, d_764_v1_, d_902_v105_), d_903_v106_, d_903_v106_, d_903_v106_])
                index122_ = default__.safeIndex(2, (d_886_v95_).length(0))
                (d_886_v95_)[index122_] = (d_904_v107_).ispropersubset(d_904_v107_)
                d_905_v108_: _dafny.Seq
                d_905_v108_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vpii"))
                d_906_v109_: _dafny.Set
                d_906_v109_ = _dafny.Set({(d_884_v93_).f23, default__.fm0(p0, d_765_v2_, (d_884_v93_).f23, globalState)})
                d_907_v110_: _dafny.Map
                d_907_v110_ = _dafny.Map({d_905_v108_: d_906_v109_})
                index123_ = default__.safeIndex(2, (d_886_v95_).length(0))
                (d_886_v95_)[index123_] = ((d_907_v110_) | (d_907_v110_)) == (d_907_v110_)
                d_908_v111_: D1
                d_908_v111_ = D1_DC5(d_905_v108_, d_764_v1_, (0) - (default__.fm9(d_764_v1_, d_764_v1_, p0, globalState)))
                d_909_v113_: _dafny.MultiSet
                d_909_v113_ = _dafny.MultiSet([740])
                d_910_v114_: _dafny.MultiSet
                def iife61_():
                    coll31_ = _dafny.Map()
                    compr_31_: int
                    for compr_31_ in (d_909_v113_).Elements:
                        d_911_v112_: int = compr_31_
                        if (d_911_v112_) in (d_909_v113_):
                            coll31_[(d_911_v112_) * (d_764_v1_)] = p0
                    return _dafny.Map(coll31_)
                d_910_v114_ = _dafny.MultiSet([default__.safeDivisionInt(d_764_v1_, len(iife61_()
                )), d_764_v1_])
                rhs105_ = d_908_v111_
                rhs106_ = d_764_v1_
                rhs107_ = d_909_v113_
                rhs108_ = (446 if True else (d_909_v113_).cardinality)
                d_908_v111_ = rhs105_
                d_764_v1_ = rhs106_
                d_909_v113_ = rhs107_
                d_764_v1_ = rhs108_
                rhs109_ = d_765_v2_
                rhs110_ = (0) - (d_764_v1_)
                rhs111_ = d_764_v1_
                rhs112_ = d_765_v2_
                rhs113_ = d_902_v105_
                d_765_v2_ = rhs109_
                d_764_v1_ = rhs110_
                d_764_v1_ = rhs111_
                d_765_v2_ = rhs112_
                d_902_v105_ = rhs113_
        r0 = not(not((self).f21))
        return r0

    def m6(self, p0, p1, p2, globalState):
        r0: bool = False
        r1: bool = False
        r2: bool = False
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_912_i0_: int
        d_912_i0_ = 0
        with _dafny.label("4"):
            while p2:
                with _dafny.c_label("4"):
                    if (d_912_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_912_i0_ = (d_912_i0_) + (1)
                    if (self).f21:
                        d_913_v0_: D3
                        d_913_v0_ = D3_DC10(p0)
                        d_914_v1_: _dafny.Seq
                        d_914_v1_ = _dafny.SeqWithoutIsStrInference([False, (self).f21])
                        d_915_v2_: _dafny.Map
                        d_915_v2_ = _dafny.Map({p2: d_914_v1_})
                        d_916_v3_: _dafny.Map
                        d_916_v3_ = _dafny.Map({len(p1): d_915_v2_})
                        d_917_v4_: _dafny.Map
                        d_917_v4_ = _dafny.Map({default__.fm9((d_913_v0_).cf17, p0, p2, globalState): ((d_916_v3_)[p0] if (p0) in (d_916_v3_) else d_915_v2_)})
                        d_918_v5_: _dafny.Seq
                        d_918_v5_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_919_v6_: _dafny.Set
                        d_919_v6_ = _dafny.Set({len(_dafny.Set({154}))})
                        d_920_v7_: _dafny.Map
                        d_920_v7_ = _dafny.Map({p0: len(d_919_v6_)})
                        d_921_v8_: _dafny.Map
                        d_921_v8_ = _dafny.Map({p2: p2})
                        d_922_v9_: D1
                        d_922_v9_ = D1_DC5(p1, p0, len(p1))
                        d_923_v10_: _dafny.Array
                        nw160_ = _dafny.Array(None, 22)
                        nw160_[int(0)] = p0
                        nw160_[int(1)] = p0
                        nw160_[int(2)] = p0
                        nw160_[int(3)] = len(((d_916_v3_)[len(p1)] if (len(p1)) in (d_916_v3_) else d_915_v2_))
                        nw160_[int(4)] = (_dafny.MultiSet([(self).f21])).cardinality
                        nw160_[int(5)] = p0
                        nw160_[int(6)] = p0
                        nw160_[int(7)] = p0
                        nw160_[int(8)] = len(p1)
                        nw160_[int(9)] = p0
                        nw160_[int(10)] = p0
                        nw160_[int(11)] = p0
                        nw160_[int(12)] = (0) - ((d_918_v5_)[default__.safeIndex(default__.fm9(len(d_920_v7_), p0, (self).f21, globalState), len(d_918_v5_))])
                        nw160_[int(13)] = p0
                        nw160_[int(14)] = len(d_914_v1_)
                        nw160_[int(15)] = p0
                        nw160_[int(16)] = len(d_921_v8_)
                        nw160_[int(17)] = p0
                        nw160_[int(18)] = (d_922_v9_).cf9
                        nw160_[int(19)] = p0
                        nw160_[int(20)] = p0
                        nw160_[int(21)] = 659
                        d_923_v10_ = nw160_
                        d_924_v11_: C1
                        nw161_ = C1()
                        nw161_.ctor__(d_923_v10_)
                        d_924_v11_ = nw161_
                        d_925_v12_: _dafny.Array
                        nw162_ = _dafny.Array(D3.default()(), 18)
                        d_925_v12_ = nw162_
                        d_926_v13_: _dafny.Array
                        def lambda57_(d_927_p2_):
                            def lambda58_(d_928_i1_):
                                return d_927_p2_

                            return lambda58_

                        init28_ = lambda57_(p2)
                        nw163_ = _dafny.Array(None, 14)
                        for i0_28_ in range(nw163_.length(0)):
                            nw163_[i0_28_] = init28_(i0_28_)
                        d_926_v13_ = nw163_
                        index124_ = default__.safeIndex(941, (d_926_v13_).length(0))
                        (d_926_v13_)[index124_] = (self).f21
                        index125_ = default__.safeIndex(941, (d_926_v13_).length(0))
                        rhs114_ = d_925_v12_
                        rhs115_ = (self).f21
                        lhs58_ = d_926_v13_
                        lhs59_ = default__.safeIndex(941, (d_926_v13_).length(0))
                        d_925_v12_ = rhs114_
                        lhs58_[lhs59_] = rhs115_
                        d_929_v14_: C1
                        nw164_ = C1()
                        nw164_.ctor__((d_924_v11_).f22)
                        d_929_v14_ = nw164_
                        d_926_v13_ = d_926_v13_
                        r2 = p2
                    elif True:
                        d_930_v15_: int
                        d_930_v15_ = 714
                        d_931_v16_: _dafny.Seq
                        d_931_v16_ = _dafny.SeqWithoutIsStrInference([not(p2)])
                        d_932_v17_: _dafny.Seq
                        d_932_v17_ = _dafny.SeqWithoutIsStrInference([d_931_v16_])
                        d_933_v18_: _dafny.Seq
                        d_933_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_931_v16_ for d_934_i2_ in range(default__.abs(857))]), d_932_v17_])
                        d_935_v19_: _dafny.Map
                        d_935_v19_ = _dafny.Map({p2: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_936_i3_ in range(default__.abs(222))])})
                        d_937_v20_: str
                        d_937_v20_ = _dafny.CodePoint('c')
                        d_938_v21_: _dafny.MultiSet
                        d_938_v21_ = _dafny.MultiSet([d_930_v15_])
                        d_939_v22_: _dafny.Seq
                        d_939_v22_ = _dafny.SeqWithoutIsStrInference([p0, d_930_v15_])
                        d_940_v23_: _dafny.Map
                        d_940_v23_ = _dafny.Map({d_930_v15_: 300})
                        d_941_v24_: _dafny.Map
                        d_941_v24_ = _dafny.Map({len(d_940_v23_): d_930_v15_})
                        d_942_v25_: _dafny.Array
                        nw165_ = _dafny.Array(None, 22)
                        nw165_[int(0)] = p0
                        nw165_[int(1)] = p0
                        nw165_[int(2)] = d_930_v15_
                        nw165_[int(3)] = d_930_v15_
                        nw165_[int(4)] = len(default__.fm16(default__.fm0(default__.fm0(p2, d_937_v20_, False, globalState), d_937_v20_, p2, globalState), p0, globalState))
                        nw165_[int(5)] = default__.fm9(p0, len(d_931_v16_), p2, globalState)
                        nw165_[int(6)] = p0
                        nw165_[int(7)] = (d_938_v21_).cardinality
                        nw165_[int(8)] = p0
                        nw165_[int(9)] = d_930_v15_
                        nw165_[int(10)] = 418
                        nw165_[int(11)] = p0
                        nw165_[int(12)] = d_930_v15_
                        nw165_[int(13)] = p0
                        nw165_[int(14)] = len(d_939_v22_)
                        nw165_[int(15)] = d_930_v15_
                        nw165_[int(16)] = 594
                        nw165_[int(17)] = len(d_941_v24_)
                        nw165_[int(18)] = p0
                        nw165_[int(19)] = 284
                        nw165_[int(20)] = d_930_v15_
                        nw165_[int(21)] = len(_dafny.Map({(self).f21: p0}))
                        d_942_v25_ = nw165_
                        d_943_v26_: _dafny.Seq
                        d_943_v26_ = _dafny.SeqWithoutIsStrInference([p1])
                        d_944_v27_: _dafny.Map
                        d_944_v27_ = _dafny.Map({p2: len(_dafny.SeqWithoutIsStrInference([d_937_v20_ for d_945_i4_ in range(default__.abs(214))]))})
                        d_946_v28_: _dafny.Array
                        nw166_ = _dafny.Array(None, 22)
                        nw166_[int(0)] = len(d_931_v16_)
                        nw166_[int(1)] = d_930_v15_
                        nw166_[int(2)] = p0
                        nw166_[int(3)] = (0) - (p0)
                        nw166_[int(4)] = p0
                        nw166_[int(5)] = p0
                        nw166_[int(6)] = p0
                        nw166_[int(7)] = d_930_v15_
                        nw166_[int(8)] = p0
                        nw166_[int(9)] = len((d_933_v18_)[default__.safeIndex(d_930_v15_, len(d_933_v18_))])
                        nw166_[int(10)] = len(d_935_v19_)
                        nw166_[int(11)] = p0
                        nw166_[int(12)] = d_930_v15_
                        nw166_[int(13)] = len(p1)
                        nw166_[int(14)] = d_930_v15_
                        nw166_[int(15)] = default__.fm9(83, d_930_v15_, (self).f21, globalState)
                        nw166_[int(16)] = len(_dafny.SeqWithoutIsStrInference([d_942_v25_, d_942_v25_]))
                        nw166_[int(17)] = len((d_943_v26_)[default__.safeIndex(p0, len(d_943_v26_))])
                        nw166_[int(18)] = d_930_v15_
                        nw166_[int(19)] = len(d_931_v16_)
                        nw166_[int(20)] = len(d_944_v27_)
                        nw166_[int(21)] = len(p1)
                        d_946_v28_ = nw166_
                        d_930_v15_ = (D3_DC11(_dafny.Map({(self).f21: d_930_v15_}), d_930_v15_, d_942_v25_)).cf19
                        d_947_v29_: _dafny.Map
                        d_947_v29_ = _dafny.Map({_dafny.CodePoint('a'): d_930_v15_})
                        d_930_v15_ = len(_dafny.SeqWithoutIsStrInference([d_939_v22_, _dafny.SeqWithoutIsStrInference([d_930_v15_, d_930_v15_]), _dafny.SeqWithoutIsStrInference([544, len(d_947_v29_)])]))
                        (globalState).f15 = (self).f21
                        d_930_v15_ = default__.safeDivisionInt((d_939_v22_)[default__.safeIndex(d_930_v15_, len(d_939_v22_))], p0)
                        d_948_v30_: _dafny.Set
                        d_948_v30_ = _dafny.Set({(self).f21})
                        d_930_v15_ = (len(d_948_v30_)) * (d_930_v15_)
                    d_949_v31_: _dafny.Array
                    nw167_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                    d_949_v31_ = nw167_
                    d_950_v32_: _dafny.Map
                    d_950_v32_ = _dafny.Map({p1: d_949_v31_})
                    d_951_v33_: _dafny.Map
                    d_951_v33_ = _dafny.Map({(p1) + (p1): d_950_v32_})
                    d_951_v33_ = (d_951_v33_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ubtj")), d_950_v32_)
                    d_952_v34_: _dafny.Array
                    def lambda59_(d_953_p0_):
                        def lambda60_(d_954_i5_):
                            return default__.safeModuloInt(d_954_i5_, d_953_p0_)

                        return lambda60_

                    init29_ = lambda59_(p0)
                    nw168_ = _dafny.Array(None, 28)
                    for i0_29_ in range(nw168_.length(0)):
                        nw168_[i0_29_] = init29_(i0_29_)
                    d_952_v34_ = nw168_
                    d_955_v35_: _dafny.Seq
                    d_955_v35_ = _dafny.SeqWithoutIsStrInference([d_952_v34_])
                    d_956_v36_: C1
                    nw169_ = C1()
                    nw169_.ctor__(d_952_v34_)
                    d_956_v36_ = nw169_
                    d_957_v37_: D5
                    d_957_v37_ = D5_DC16(d_956_v36_)
                    d_958_v38_: _dafny.Seq
                    d_958_v38_ = _dafny.SeqWithoutIsStrInference([(d_957_v37_).cf28])
                    d_959_v39_: C1
                    nw170_ = C1()
                    nw170_.ctor__((d_955_v35_)[default__.safeIndex(len(d_958_v38_), len(d_955_v35_))])
                    d_959_v39_ = nw170_
                    d_960_v40_: _dafny.Map
                    d_960_v40_ = _dafny.Map({(self).f21: d_952_v34_})
                    d_960_v40_ = (d_960_v40_).set(((self).f21) and (p2), (d_959_v39_).f22)
                    pass
            pass
        d_961_i6_: int
        d_961_i6_ = 0
        with _dafny.label("5"):
            while p2:
                with _dafny.c_label("5"):
                    if (d_961_i6_) >= (100):
                        raise _dafny.Break("5")
                    d_961_i6_ = (d_961_i6_) + (1)
                    d_962_v41_: _dafny.Array
                    nw171_ = _dafny.Array(False, 23)
                    d_962_v41_ = nw171_
                    d_963_v42_: D1
                    d_963_v42_ = D1_DC4(451, not((self).f21), d_962_v41_)
                    source12_ = d_963_v42_
                    if source12_.is_DC4:
                        d_964___mcc_h0_ = source12_.cf4
                        d_965___mcc_h1_ = source12_.cf5
                        d_966___mcc_h2_ = source12_.cf6
                        d_967_cf6_ = d_966___mcc_h2_
                        d_968_cf5_ = d_965___mcc_h1_
                        d_969_cf4_ = d_964___mcc_h0_
                        r2 = p2
                        d_970_v43_: C0
                        nw172_ = C0()
                        nw172_.ctor__(p2)
                        d_970_v43_ = nw172_
                        d_971_v44_: C0
                        nw173_ = C0()
                        nw173_.ctor__(not((self).f21))
                        d_971_v44_ = nw173_
                        d_972_v45_: str
                        d_972_v45_ = _dafny.CodePoint('b')
                        d_973_v46_: _dafny.Map
                        d_973_v46_ = _dafny.Map({(p1 if (d_971_v44_).f23 else (p1).set(default__.safeIndex((0) - (d_969_cf4_), len(p1)), d_972_v45_)): len(_dafny.Set({(self).f21, d_968_cf5_, (d_971_v44_).f23, (self).f21}))})
                        d_974_v47_: _dafny.Map
                        d_974_v47_ = _dafny.Map({(self).f21: (d_970_v43_).f23})
                        d_973_v46_ = default__.fm25((d_970_v43_).f23, (d_974_v47_) == (d_974_v47_), globalState)
                    elif source12_.is_DC5:
                        d_975___mcc_h3_ = source12_.cf7
                        d_976___mcc_h4_ = source12_.cf8
                        d_977___mcc_h5_ = source12_.cf9
                        d_978_cf9_ = d_977___mcc_h5_
                        d_979_cf8_ = d_976___mcc_h4_
                        d_980_cf7_ = d_975___mcc_h3_
                        d_981_v48_: _dafny.Seq
                        d_981_v48_ = _dafny.SeqWithoutIsStrInference([len(p1)])
                        d_982_v49_: _dafny.Array
                        nw174_ = _dafny.Array(None, 17)
                        nw174_[int(0)] = d_979_cf8_
                        nw174_[int(1)] = (p0) * (d_978_cf9_)
                        nw174_[int(2)] = d_979_cf8_
                        nw174_[int(3)] = default__.safeDivisionInt(d_978_cf9_, p0)
                        nw174_[int(4)] = d_978_cf9_
                        nw174_[int(5)] = p0
                        nw174_[int(6)] = p0
                        nw174_[int(7)] = d_979_cf8_
                        nw174_[int(8)] = 723
                        nw174_[int(9)] = 785
                        nw174_[int(10)] = d_978_cf9_
                        nw174_[int(11)] = p0
                        nw174_[int(12)] = len(_dafny.Set({(d_981_v48_)[default__.safeIndex(d_979_cf8_, len(d_981_v48_))], d_979_cf8_, d_978_cf9_, p0}))
                        nw174_[int(13)] = d_979_cf8_
                        nw174_[int(14)] = (d_981_v48_)[default__.safeIndex(p0, len(d_981_v48_))]
                        nw174_[int(15)] = p0
                        nw174_[int(16)] = default__.safeModuloInt(p0, p0)
                        d_982_v49_ = nw174_
                        index126_ = default__.safeIndex(620, (d_982_v49_).length(0))
                        (d_982_v49_)[index126_] = d_978_cf9_
                        d_983_v50_: _dafny.Map
                        d_983_v50_ = _dafny.Map({d_982_v49_: d_982_v49_})
                        index127_ = default__.safeIndex(620, (d_982_v49_).length(0))
                        (d_982_v49_)[index127_] = len(d_983_v50_)
                        (globalState).f11 = (self).f21
                        d_984_v51_: _dafny.Seq
                        d_984_v51_ = _dafny.SeqWithoutIsStrInference([(self).f21])
                        index128_ = default__.safeIndex(696, (d_962_v41_).length(0))
                        (d_962_v41_)[index128_] = (d_984_v51_)[default__.safeIndex(len(_dafny.Set({(self).f21, (self).f21, (d_984_v51_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xlt"))), len(d_984_v51_))]})), len(d_984_v51_))]
                        d_985_v52_: _dafny.Array
                        nw175_ = _dafny.Array(_dafny.CodePoint('D'), 3)
                        d_985_v52_ = nw175_
                        d_986_v53_: str
                        d_986_v53_ = _dafny.CodePoint('h')
                        index129_ = default__.safeIndex(339, (d_985_v52_).length(0))
                        (d_985_v52_)[index129_] = d_986_v53_
                        d_987_v54_: _dafny.Set
                        d_987_v54_ = _dafny.Set({not((self).f21)})
                        d_988_v55_: T0
                        nw176_ = C0()
                        nw176_.ctor__((d_987_v54_).issubset(d_987_v54_))
                        d_988_v55_ = nw176_
                        index130_ = default__.safeIndex(696, (d_962_v41_).length(0))
                        index131_ = default__.safeIndex(339, (d_985_v52_).length(0))
                        rhs116_ = p2
                        rhs117_ = d_978_cf9_
                        rhs118_ = (d_987_v54_).issubset((_dafny.Set({(self).f21})) - (d_987_v54_))
                        rhs119_ = d_986_v53_
                        rhs120_ = d_988_v55_
                        lhs60_ = d_962_v41_
                        lhs61_ = default__.safeIndex(696, (d_962_v41_).length(0))
                        lhs62_ = d_985_v52_
                        lhs63_ = default__.safeIndex(339, (d_985_v52_).length(0))
                        lhs60_[lhs61_] = rhs116_
                        d_978_cf9_ = rhs117_
                        r0 = rhs118_
                        lhs62_[lhs63_] = rhs119_
                        d_988_v55_ = rhs120_
                        index132_ = default__.safeIndex(620, (d_982_v49_).length(0))
                        (d_982_v49_)[index132_] = default__.safeDivisionInt((0) - (d_978_cf9_), d_979_cf8_)
                    elif source12_.is_DC6:
                        d_989___mcc_h6_ = source12_.cf10
                        d_990___mcc_h7_ = source12_.cf11
                        d_991_cf11_ = d_990___mcc_h7_
                        d_992_cf10_ = d_989___mcc_h6_
                        d_992_cf10_ = d_991_cf11_
                        d_993_v56_: _dafny.MultiSet
                        d_993_v56_ = _dafny.MultiSet([p0, p0])
                        d_994_v57_: _dafny.Map
                        d_994_v57_ = _dafny.Map({d_993_v56_: d_992_cf10_})
                        r1 = (d_994_v57_) == (d_994_v57_)
                        d_995_v58_: _dafny.Seq
                        d_995_v58_ = _dafny.SeqWithoutIsStrInference([d_962_v41_, d_962_v41_, d_962_v41_, d_962_v41_, d_962_v41_])
                        r2 = (d_962_v41_) in (d_995_v58_)
                        r3 = p1
                    elif True:
                        d_996___mcc_h8_ = source12_.cf3
                        d_997_cf3_ = d_996___mcc_h8_
                        d_998_v59_: _dafny.Map
                        d_998_v59_ = _dafny.Map({p0: p0})
                        d_999_v60_: _dafny.MultiSet
                        d_999_v60_ = _dafny.MultiSet([d_998_v59_])
                        d_1000_v61_: D6
                        d_1000_v61_ = D6_DC19(_dafny.MultiSet([d_998_v59_, d_998_v59_]))
                        d_999_v60_ = (d_999_v60_) - ((d_1000_v61_).cf35)
                        d_1001_v62_: _dafny.Array
                        nw177_ = _dafny.Array(_dafny.Seq({}), 28)
                        d_1001_v62_ = nw177_
                        d_1002_v63_: _dafny.Set
                        d_1002_v63_ = _dafny.Set({len(p1), p0, p0})
                        d_1003_v64_: _dafny.Seq
                        d_1003_v64_ = _dafny.SeqWithoutIsStrInference([len(d_1002_v63_), p0])
                        index133_ = default__.safeIndex(502, (d_1001_v62_).length(0))
                        (d_1001_v62_)[index133_] = d_1003_v64_
                        index134_ = default__.safeIndex(502, (d_1001_v62_).length(0))
                        (d_1001_v62_)[index134_] = d_1003_v64_
                        d_1004_v65_: _dafny.Array
                        nw178_ = _dafny.Array(int(0), 19)
                        d_1004_v65_ = nw178_
                        index135_ = default__.safeIndex(796, (d_1004_v65_).length(0))
                        (d_1004_v65_)[index135_] = len((p1) + (default__.fm16((self).f21, p0, globalState)))
                        d_1005_v66_: D6
                        d_1005_v66_ = D6_DC20((self).f21, p0, True)
                        index136_ = default__.safeIndex(796, (d_1004_v65_).length(0))
                        rhs121_ = p2
                        rhs122_ = (d_1005_v66_).cf36
                        rhs123_ = (0) - (p0)
                        lhs64_ = globalState
                        lhs65_ = globalState
                        lhs66_ = d_1004_v65_
                        lhs67_ = default__.safeIndex(796, (d_1004_v65_).length(0))
                        lhs64_.f17 = rhs121_
                        lhs65_.f17 = rhs122_
                        lhs66_[lhs67_] = rhs123_
                        d_1006_v67_: _dafny.Seq
                        d_1006_v67_ = _dafny.SeqWithoutIsStrInference([(p2 if (self).f21 else (self).f21), (self).f21, (self).f21, p2])
                        (globalState).f16 = (d_1006_v67_).set(default__.safeIndex(p0, len(d_1006_v67_)), (self).f21)
                    d_1007_v68_: _dafny.Set
                    d_1007_v68_ = _dafny.Set({not((self).f21), p2, p2})
                    d_1008_v69_: _dafny.Seq
                    d_1008_v69_ = _dafny.SeqWithoutIsStrInference([default__.fm20(globalState)])
                    d_1009_v70_: int
                    d_1009_v70_ = 616
                    rhs124_ = (d_1007_v68_).intersection(_dafny.Set({True}))
                    rhs125_ = (p1) + (p1)
                    rhs126_ = (d_1008_v69_) + (_dafny.SeqWithoutIsStrInference([d_1007_v68_, d_1007_v68_]))
                    rhs127_ = 396
                    d_1007_v68_ = rhs124_
                    r3 = rhs125_
                    d_1008_v69_ = rhs126_
                    d_1009_v70_ = rhs127_
                    d_1010_v71_: str
                    d_1010_v71_ = _dafny.CodePoint('w')
                    d_1011_v72_: _dafny.Map
                    d_1011_v72_ = _dafny.Map({d_1009_v70_: default__.fm0((self).f21, d_1010_v71_, p2, globalState)})
                    d_1012_v73_: _dafny.MultiSet
                    d_1012_v73_ = _dafny.MultiSet([(self).f21])
                    d_1013_v74_: _dafny.Array
                    nw179_ = _dafny.Array(None, 27)
                    nw179_[int(0)] = p0
                    nw179_[int(1)] = p0
                    nw179_[int(2)] = d_1009_v70_
                    nw179_[int(3)] = p0
                    nw179_[int(4)] = len(p1)
                    nw179_[int(5)] = (0) - (p0)
                    nw179_[int(6)] = p0
                    nw179_[int(7)] = d_1009_v70_
                    nw179_[int(8)] = 443
                    nw179_[int(9)] = -109
                    nw179_[int(10)] = d_1009_v70_
                    nw179_[int(11)] = 882
                    nw179_[int(12)] = 55
                    nw179_[int(13)] = p0
                    nw179_[int(14)] = d_1009_v70_
                    nw179_[int(15)] = d_1009_v70_
                    nw179_[int(16)] = default__.fm9(p0, len(d_1007_v68_), (self).f21, globalState)
                    nw179_[int(17)] = d_1009_v70_
                    nw179_[int(18)] = len(d_1011_v72_)
                    nw179_[int(19)] = p0
                    nw179_[int(20)] = d_1009_v70_
                    nw179_[int(21)] = d_1009_v70_
                    nw179_[int(22)] = (d_1012_v73_).cardinality
                    nw179_[int(23)] = p0
                    nw179_[int(24)] = p0
                    nw179_[int(25)] = p0
                    nw179_[int(26)] = d_1009_v70_
                    d_1013_v74_ = nw179_
                    d_1014_v75_: C1
                    nw180_ = C1()
                    nw180_.ctor__(d_1013_v74_)
                    d_1014_v75_ = nw180_
                    d_1015_v76_: T0
                    nw181_ = C0()
                    nw181_.ctor__((self).f21)
                    d_1015_v76_ = nw181_
                    d_1016_v77_: D4
                    d_1016_v77_ = D4_DC12(d_1015_v76_)
                    d_1017_v78_: _dafny.Map
                    d_1017_v78_ = _dafny.Map({(d_1012_v73_).cardinality: d_1015_v76_})
                    d_1018_v79_: D2
                    d_1018_v79_ = D2_DC8(d_962_v41_, p0, p0)
                    d_1019_v80_: _dafny.Map
                    d_1019_v80_ = _dafny.Map({d_1018_v79_: d_1015_v76_})
                    d_1020_v81_: _dafny.Seq
                    d_1020_v81_ = _dafny.SeqWithoutIsStrInference([d_1009_v70_])
                    d_1021_v82_: _dafny.Map
                    d_1021_v82_ = _dafny.Map({(self).f21: d_1020_v81_})
                    d_1022_v83_: _dafny.Array
                    nw182_ = _dafny.Array(None, 25)
                    nw182_[int(0)] = d_1015_v76_
                    nw182_[int(1)] = d_1015_v76_
                    nw182_[int(2)] = (d_1016_v77_).cf21
                    nw182_[int(3)] = d_1015_v76_
                    nw182_[int(4)] = d_1015_v76_
                    nw182_[int(5)] = (d_1016_v77_).cf21
                    nw182_[int(6)] = d_1015_v76_
                    nw182_[int(7)] = d_1015_v76_
                    nw182_[int(8)] = d_1015_v76_
                    nw182_[int(9)] = ((d_1017_v78_)[p0] if (p0) in (d_1017_v78_) else d_1015_v76_)
                    nw182_[int(10)] = d_1015_v76_
                    nw182_[int(11)] = d_1015_v76_
                    nw182_[int(12)] = d_1015_v76_
                    nw182_[int(13)] = d_1015_v76_
                    nw182_[int(14)] = d_1015_v76_
                    nw182_[int(15)] = (d_1016_v77_).cf21
                    nw182_[int(16)] = d_1015_v76_
                    nw182_[int(17)] = d_1015_v76_
                    nw182_[int(18)] = d_1015_v76_
                    nw182_[int(19)] = d_1015_v76_
                    nw182_[int(20)] = d_1015_v76_
                    nw182_[int(21)] = d_1015_v76_
                    nw182_[int(22)] = d_1015_v76_
                    nw182_[int(23)] = d_1015_v76_
                    nw182_[int(24)] = ((d_1019_v80_)[D2_DC8(d_962_v41_, len(d_1021_v82_), (default__.fm26(False, _dafny.CodePoint('m'), (self).f21, p2, globalState)).cf37)] if (D2_DC8(d_962_v41_, len(d_1021_v82_), (default__.fm26(False, _dafny.CodePoint('m'), (self).f21, p2, globalState)).cf37)) in (d_1019_v80_) else d_1015_v76_)
                    d_1022_v83_ = nw182_
                    index137_ = default__.safeIndex(948, (d_1022_v83_).length(0))
                    (d_1022_v83_)[index137_] = d_1015_v76_
                    d_1023_v84_: _dafny.Seq
                    d_1023_v84_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_1024_v85_: _dafny.MultiSet
                    d_1024_v85_ = _dafny.MultiSet([p0, 60, p0])
                    index138_ = default__.safeIndex(948, (d_1022_v83_).length(0))
                    rhs128_ = d_1014_v75_
                    rhs129_ = (d_1009_v70_) > (d_1009_v70_)
                    rhs130_ = d_1015_v76_
                    rhs131_ = not(((_dafny.MultiSet([p0, 115, len(d_1023_v84_)])).issubset(d_1024_v85_) if True else (self).f21))
                    lhs68_ = globalState
                    lhs69_ = d_1022_v83_
                    lhs70_ = default__.safeIndex(948, (d_1022_v83_).length(0))
                    lhs71_ = globalState
                    d_1014_v75_ = rhs128_
                    lhs68_.f11 = rhs129_
                    lhs69_[lhs70_] = rhs130_
                    lhs71_.f11 = rhs131_
                    d_1025_v86_: _dafny.Array
                    nw183_ = _dafny.Array(None, 2)
                    nw183_[int(0)] = d_1010_v71_
                    nw183_[int(1)] = d_1010_v71_
                    d_1025_v86_ = nw183_
                    d_1025_v86_ = d_1025_v86_
                    pass
            pass
        if not(((p1) + (p1)) == (p1)):
            d_1026_v87_: _dafny.Set
            d_1026_v87_ = _dafny.Set({p0, p0})
            d_1027_v88_: _dafny.Map
            d_1027_v88_ = _dafny.Map({(self).f21: d_1026_v87_})
            d_1028_v89_: _dafny.Seq
            d_1028_v89_ = _dafny.SeqWithoutIsStrInference([(self).f21])
            d_1029_v90_: bool
            out53_: bool
            out53_ = (self).m5((d_1026_v87_) == (((d_1027_v88_)[False] if (False) in (d_1027_v88_) else d_1026_v87_)), d_1028_v89_, globalState)
            d_1029_v90_ = out53_
            (globalState).f11 = True
            d_1030_v91_: _dafny.Array
            def lambda61_(d_1031_i7_):
                return (d_1031_i7_) * (-579)

            init30_ = lambda61_
            nw184_ = _dafny.Array(None, 7)
            for i0_30_ in range(nw184_.length(0)):
                nw184_[i0_30_] = init30_(i0_30_)
            d_1030_v91_ = nw184_
            d_1032_v92_: C1
            nw185_ = C1()
            nw185_.ctor__(d_1030_v91_)
            d_1032_v92_ = nw185_
            d_1033_v93_: _dafny.MultiSet
            d_1033_v93_ = _dafny.MultiSet([p0, p0, p0])
            d_1034_v94_: _dafny.Map
            d_1034_v94_ = _dafny.Map({p2: p0})
            d_1035_v95_: str
            d_1035_v95_ = _dafny.CodePoint('h')
            d_1033_v93_ = (d_1033_v93_).intersection(default__.fm23(d_1029_v90_, (d_1034_v94_).set(p2, p0), p0, d_1035_v95_, globalState))
            d_1036_v96_: _dafny.Map
            d_1036_v96_ = _dafny.Map({((d_1033_v93_)[p0] if (p0) in (d_1033_v93_) else p0): p0})
            d_1037_v97_: _dafny.Seq
            d_1037_v97_ = _dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([default__.fm16(d_1029_v90_, p0, globalState)]))), p0])
            d_1036_v96_ = (d_1036_v96_).set((d_1037_v97_)[default__.safeIndex(p0, len(d_1037_v97_))], default__.safeModuloInt(p0, p0))
        elif True:
            d_1038_v98_: _dafny.Array
            nw186_ = _dafny.Array(int(0), 10)
            d_1038_v98_ = nw186_
            index139_ = default__.safeIndex(151, (d_1038_v98_).length(0))
            (d_1038_v98_)[index139_] = p0
            index140_ = default__.safeIndex(151, (d_1038_v98_).length(0))
            (d_1038_v98_)[index140_] = p0
            (globalState).f15 = (self).f21
            index141_ = default__.safeIndex(151, (d_1038_v98_).length(0))
            (d_1038_v98_)[index141_] = ((0) - (p0)) + ((d_1038_v98_)[default__.safeIndex(151, (d_1038_v98_).length(0))])
            d_1039_v99_: D6
            d_1039_v99_ = D6_DC20((self).f21, (d_1038_v98_)[default__.safeIndex(151, (d_1038_v98_).length(0))], (self).f21)
            (globalState).f17 = (d_1039_v99_).cf36
            d_1040_v100_: _dafny.Array
            nw187_ = _dafny.Array(False, 15)
            d_1040_v100_ = nw187_
            d_1041_v101_: _dafny.Seq
            d_1041_v101_ = _dafny.SeqWithoutIsStrInference([d_1040_v100_, d_1040_v100_])
            d_1042_v102_: _dafny.Set
            d_1042_v102_ = _dafny.Set({d_1040_v100_, d_1040_v100_})
            d_1043_v103_: _dafny.Map
            d_1043_v103_ = _dafny.Map({p2: p0})
            d_1044_v104_: _dafny.Map
            d_1044_v104_ = _dafny.Map({len(d_1043_v103_): (d_1038_v98_)[default__.safeIndex(151, (d_1038_v98_).length(0))]})
            d_1045_v106_: str
            d_1045_v106_ = _dafny.CodePoint('u')
            d_1046_v107_: D4
            def iife62_():
                coll32_ = _dafny.Set()
                compr_32_: int
                for compr_32_ in (d_1044_v104_).keys.Elements:
                    d_1047_v105_: int = compr_32_
                    if (d_1047_v105_) in (d_1044_v104_):
                        coll32_ = coll32_.union(_dafny.Set([(d_1047_v105_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lp"))))]))
                return _dafny.Set(coll32_)
            d_1046_v107_ = D4_DC14(False, p0, iife62_()
, d_1045_v106_)
            d_1048_v108_: _dafny.Map
            d_1048_v108_ = _dafny.Map({(d_1042_v102_).issubset(_dafny.Set({(d_1041_v101_)[default__.safeIndex((d_1038_v98_)[default__.safeIndex(151, (d_1038_v98_).length(0))], len(d_1041_v101_))]})): default__.fm12((d_1046_v107_).cf24, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1049_i8_ in range(default__.abs(-558))]), _dafny.SeqWithoutIsStrInference([d_1045_v106_ for d_1050_i9_ in range(default__.abs(464))]), globalState)})
            d_1048_v108_ = (d_1048_v108_).set((self).f21, d_1045_v106_)
        d_1051_i10_: int
        d_1051_i10_ = 0
        with _dafny.label("6"):
            while (p0) != (p0):
                with _dafny.c_label("6"):
                    if (d_1051_i10_) >= (100):
                        raise _dafny.Break("6")
                    d_1051_i10_ = (d_1051_i10_) + (1)
                    d_1052_v109_: _dafny.Map
                    d_1052_v109_ = _dafny.Map({p0: (self).f21})
                    d_1053_v110_: _dafny.Set
                    d_1053_v110_ = _dafny.Set({d_1052_v109_})
                    d_1054_v111_: D3
                    d_1054_v111_ = D3_DC9(d_1053_v110_)
                    d_1054_v111_ = d_1054_v111_
                    d_1055_v112_: _dafny.Array
                    nw188_ = _dafny.Array(_dafny.Array(None, 0), 8)
                    d_1055_v112_ = nw188_
                    d_1056_v113_: str
                    d_1056_v113_ = _dafny.CodePoint('i')
                    d_1057_v114_: _dafny.Map
                    d_1057_v114_ = _dafny.Map({len(p1): p1})
                    d_1058_v115_: _dafny.Seq
                    d_1058_v115_ = _dafny.SeqWithoutIsStrInference([p0, len(_dafny.SeqWithoutIsStrInference([p0]))])
                    d_1059_v116_: D3
                    d_1059_v116_ = D3_DC10(p0)
                    d_1060_v117_: _dafny.Map
                    d_1060_v117_ = _dafny.Map({len(d_1058_v115_): d_1059_v116_})
                    d_1061_v118_: _dafny.Array
                    nw189_ = _dafny.Array(None, 21)
                    nw189_[int(0)] = p1
                    nw189_[int(1)] = p1
                    nw189_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "snqcwt"))
                    nw189_[int(3)] = p1
                    nw189_[int(4)] = p1
                    nw189_[int(5)] = p1
                    nw189_[int(6)] = p1
                    nw189_[int(7)] = p1
                    nw189_[int(8)] = p1
                    nw189_[int(9)] = p1
                    nw189_[int(10)] = p1
                    nw189_[int(11)] = p1
                    nw189_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cx"))
                    nw189_[int(13)] = (p1).set(default__.safeIndex(p0, len(p1)), d_1056_v113_)
                    nw189_[int(14)] = p1
                    nw189_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygblryv"))
                    nw189_[int(16)] = ((d_1057_v114_)[len(d_1060_v117_)] if (len(d_1060_v117_)) in (d_1057_v114_) else p1)
                    nw189_[int(17)] = p1
                    nw189_[int(18)] = p1
                    nw189_[int(19)] = p1
                    nw189_[int(20)] = _dafny.SeqWithoutIsStrInference([d_1056_v113_ for d_1062_i11_ in range(default__.abs(-367))])
                    d_1061_v118_ = nw189_
                    d_1063_v119_: _dafny.Array
                    d_1063_v119_ = d_1061_v118_
                    d_1064_v120_: _dafny.Seq
                    d_1064_v120_ = _dafny.SeqWithoutIsStrInference([(d_1061_v118_)])
                    index142_ = default__.safeIndex(246, (d_1055_v112_).length(0))
                    (d_1055_v112_)[index142_] = (d_1064_v120_)[default__.safeIndex(215, len(d_1064_v120_))]
                    d_1065_v121_: _dafny.MultiSet
                    d_1065_v121_ = _dafny.MultiSet([p2])
                    index143_ = default__.safeIndex(246, (d_1055_v112_).length(0))
                    rhs132_ = d_1061_v118_
                    rhs133_ = (419) >= (default__.safeModuloInt(p0, p0))
                    rhs134_ = (d_1065_v121_) - ((_dafny.MultiSet([p2])).set(False, default__.abs(p0)))
                    lhs72_ = d_1055_v112_
                    lhs73_ = default__.safeIndex(246, (d_1055_v112_).length(0))
                    lhs74_ = globalState
                    lhs72_[lhs73_] = rhs132_
                    lhs74_.f15 = rhs133_
                    d_1065_v121_ = rhs134_
                    d_1066_v122_: _dafny.Seq
                    d_1066_v122_ = _dafny.SeqWithoutIsStrInference([p2])
                    d_1067_v123_: _dafny.Array
                    nw190_ = _dafny.Array(None, 13)
                    nw190_[int(0)] = d_1066_v122_
                    nw190_[int(1)] = _dafny.SeqWithoutIsStrInference([p2, p2])
                    nw190_[int(2)] = d_1066_v122_
                    nw190_[int(3)] = (d_1066_v122_ if p2 else d_1066_v122_)
                    nw190_[int(4)] = d_1066_v122_
                    nw190_[int(5)] = (d_1066_v122_) + (_dafny.SeqWithoutIsStrInference([True, (self).f21]))
                    nw190_[int(6)] = d_1066_v122_
                    nw190_[int(7)] = d_1066_v122_
                    nw190_[int(8)] = d_1066_v122_
                    nw190_[int(9)] = d_1066_v122_
                    nw190_[int(10)] = d_1066_v122_
                    nw190_[int(11)] = (d_1066_v122_) + (d_1066_v122_)
                    nw190_[int(12)] = (d_1066_v122_) + (d_1066_v122_)
                    d_1067_v123_ = nw190_
                    index144_ = default__.safeIndex(539, (d_1067_v123_).length(0))
                    (d_1067_v123_)[index144_] = (d_1066_v122_) + (default__.fm27(p0, p0, (0) - (p0), globalState))
                    index145_ = default__.safeIndex(539, (d_1067_v123_).length(0))
                    (d_1067_v123_)[index145_] = (_dafny.SeqWithoutIsStrInference([p2])).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference([p2]))), p2)
                    d_1068_v124_: _dafny.Array
                    def lambda62_(d_1069_p0_):
                        def lambda63_(d_1070_i12_):
                            return default__.safeModuloInt(d_1070_i12_, d_1069_p0_)

                        return lambda63_

                    init31_ = lambda62_(p0)
                    nw191_ = _dafny.Array(None, 7)
                    for i0_31_ in range(nw191_.length(0)):
                        nw191_[i0_31_] = init31_(i0_31_)
                    d_1068_v124_ = nw191_
                    d_1071_v125_: _dafny.Seq
                    d_1071_v125_ = _dafny.SeqWithoutIsStrInference([d_1068_v124_])
                    d_1071_v125_ = _dafny.SeqWithoutIsStrInference([d_1068_v124_])
                    pass
            pass
        hi7_ = p0
        for d_1072_i13_ in range(p0, hi7_):
            d_1073_v126_: _dafny.Array
            def lambda64_(d_1074_i14_):
                return _dafny.SeqWithoutIsStrInference([(self).f21, (self).f21])

            init32_ = lambda64_
            nw192_ = _dafny.Array(None, 24)
            for i0_32_ in range(nw192_.length(0)):
                nw192_[i0_32_] = init32_(i0_32_)
            d_1073_v126_ = nw192_
            d_1075_v127_: _dafny.Seq
            d_1075_v127_ = _dafny.SeqWithoutIsStrInference([(self).f21])
            index146_ = default__.safeIndex(351, (d_1073_v126_).length(0))
            (d_1073_v126_)[index146_] = (d_1075_v127_) + (d_1075_v127_)
            d_1076_v128_: _dafny.Array
            nw193_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_1076_v128_ = nw193_
            d_1077_v129_: int
            d_1077_v129_ = 808
            d_1078_v130_: D8
            d_1078_v130_ = D8_DC22(d_1076_v128_)
            index147_ = default__.safeIndex(351, (d_1073_v126_).length(0))
            rhs135_ = not((default__.safeDivisionInt(-12, (0) - (d_1077_v129_))) >= (default__.fm9(p0, p0, False, globalState)))
            rhs136_ = d_1075_v127_
            rhs137_ = (d_1078_v130_).cf40
            rhs138_ = (p0) - (d_1077_v129_)
            rhs139_ = d_1072_i13_
            lhs75_ = d_1073_v126_
            lhs76_ = default__.safeIndex(351, (d_1073_v126_).length(0))
            r1 = rhs135_
            lhs75_[lhs76_] = rhs136_
            d_1076_v128_ = rhs137_
            d_1077_v129_ = rhs138_
            d_1077_v129_ = rhs139_
            d_1079_v131_: _dafny.Map
            d_1079_v131_ = _dafny.Map({(self).f21: d_1072_i13_})
            d_1080_v133_: _dafny.Array
            def lambda65_(d_1081_v131_, d_1082_p0_):
                def lambda66_(d_1083_i15_):
                    def iife63_():
                        coll33_ = _dafny.Map()
                        compr_33_: int
                        for compr_33_ in _dafny.IntegerRange(-447, 915):
                            d_1084_v132_: int = compr_33_
                            if ((-447) <= (d_1084_v132_)) and ((d_1084_v132_) < (915)):
                                coll33_[(d_1084_v132_) - (d_1082_p0_)] = d_1082_p0_
                        return _dafny.Map(coll33_)
                    return (iife63_()
                    ) | (_dafny.Map({515: len(d_1081_v131_)}))

                return lambda66_

            init33_ = lambda65_(d_1079_v131_, p0)
            nw194_ = _dafny.Array(None, 15)
            for i0_33_ in range(nw194_.length(0)):
                nw194_[i0_33_] = init33_(i0_33_)
            d_1080_v133_ = nw194_
            d_1085_v134_: _dafny.MultiSet
            d_1085_v134_ = _dafny.MultiSet([p2])
            d_1086_v135_: _dafny.Map
            d_1086_v135_ = _dafny.Map({(d_1085_v134_).cardinality: len(p1)})
            index148_ = default__.safeIndex(876, (d_1080_v133_).length(0))
            (d_1080_v133_)[index148_] = d_1086_v135_
            index149_ = default__.safeIndex(876, (d_1080_v133_).length(0))
            rhs140_ = (0) - (default__.safeModuloInt(p0, default__.fm9(682, len(p1), (self).f21, globalState)))
            rhs141_ = d_1079_v131_
            rhs142_ = default__.fm9(799, p0, (_dafny.MultiSet([not(False), (self).f21])).ispropersubset(default__.fm28(d_1072_i13_, p2, globalState)), globalState)
            rhs143_ = default__.fm10(globalState)
            lhs77_ = d_1080_v133_
            lhs78_ = default__.safeIndex(876, (d_1080_v133_).length(0))
            d_1077_v129_ = rhs140_
            d_1079_v131_ = rhs141_
            d_1077_v129_ = rhs142_
            lhs77_[lhs78_] = rhs143_
            d_1087_v136_: _dafny.Map
            d_1087_v136_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_1088_i16_ in range(default__.abs(322))])): (self).f21})
            d_1089_v137_: _dafny.Seq
            d_1089_v137_ = _dafny.SeqWithoutIsStrInference([d_1072_i13_])
            d_1090_v138_: _dafny.Set
            d_1090_v138_ = _dafny.Set({(self).f21, p2})
            d_1091_v139_: str
            d_1091_v139_ = _dafny.CodePoint('q')
            d_1092_v140_: _dafny.Map
            d_1092_v140_ = _dafny.Map({d_1091_v139_: d_1091_v139_})
            d_1093_v141_: _dafny.MultiSet
            d_1093_v141_ = _dafny.MultiSet([len(d_1092_v140_), d_1077_v129_, d_1072_i13_, d_1077_v129_, len(_dafny.Map({p2: d_1077_v129_}))])
            d_1094_v142_: _dafny.Array
            nw195_ = _dafny.Array(None, 27)
            nw195_[int(0)] = p0
            nw195_[int(1)] = p0
            nw195_[int(2)] = d_1072_i13_
            nw195_[int(3)] = len((p1).set(default__.safeIndex(d_1077_v129_, len(p1)), _dafny.CodePoint('t')))
            nw195_[int(4)] = -593
            nw195_[int(5)] = (0) - (p0)
            nw195_[int(6)] = d_1077_v129_
            nw195_[int(7)] = d_1077_v129_
            nw195_[int(8)] = len(d_1087_v136_)
            nw195_[int(9)] = len(d_1089_v137_)
            nw195_[int(10)] = d_1077_v129_
            nw195_[int(11)] = d_1077_v129_
            nw195_[int(12)] = d_1077_v129_
            nw195_[int(13)] = p0
            nw195_[int(14)] = d_1077_v129_
            nw195_[int(15)] = len(d_1089_v137_)
            nw195_[int(16)] = default__.fm9(d_1077_v129_, len(p1), p2, globalState)
            nw195_[int(17)] = (_dafny.MultiSet(p1)).cardinality
            nw195_[int(18)] = p0
            nw195_[int(19)] = p0
            nw195_[int(20)] = d_1077_v129_
            nw195_[int(21)] = p0
            nw195_[int(22)] = d_1072_i13_
            nw195_[int(23)] = len(d_1090_v138_)
            nw195_[int(24)] = ((d_1093_v141_)[len(p1)] if (len(p1)) in (d_1093_v141_) else d_1077_v129_)
            nw195_[int(25)] = d_1077_v129_
            nw195_[int(26)] = len(_dafny.Map({p0: (self).f21}))
            d_1094_v142_ = nw195_
            d_1095_v143_: _dafny.MultiSet
            d_1095_v143_ = _dafny.MultiSet([d_1094_v142_])
            d_1096_v144_: _dafny.Seq
            d_1096_v144_ = _dafny.SeqWithoutIsStrInference([d_1095_v143_])
            d_1077_v129_ = ((d_1096_v144_)[default__.safeIndex(p0, len(d_1096_v144_))]).cardinality
            d_1097_v145_: _dafny.Array
            nw196_ = _dafny.Array(_dafny.CodePoint('D'), 5)
            d_1097_v145_ = nw196_
            d_1098_v146_: _dafny.Seq
            d_1098_v146_ = _dafny.SeqWithoutIsStrInference([d_1097_v145_, d_1097_v145_])
            d_1099_v147_: _dafny.Array
            nw197_ = _dafny.Array(None, 26)
            nw197_[int(0)] = d_1097_v145_
            nw197_[int(1)] = d_1097_v145_
            nw197_[int(2)] = d_1097_v145_
            nw197_[int(3)] = d_1097_v145_
            nw197_[int(4)] = d_1097_v145_
            nw197_[int(5)] = d_1097_v145_
            nw197_[int(6)] = d_1097_v145_
            nw197_[int(7)] = d_1097_v145_
            nw197_[int(8)] = d_1097_v145_
            nw197_[int(9)] = d_1097_v145_
            nw197_[int(10)] = d_1097_v145_
            nw197_[int(11)] = d_1097_v145_
            nw197_[int(12)] = d_1097_v145_
            nw197_[int(13)] = d_1097_v145_
            nw197_[int(14)] = d_1097_v145_
            nw197_[int(15)] = (d_1098_v146_)[default__.safeIndex(p0, len(d_1098_v146_))]
            nw197_[int(16)] = d_1097_v145_
            nw197_[int(17)] = d_1097_v145_
            nw197_[int(18)] = d_1097_v145_
            nw197_[int(19)] = (d_1097_v145_ if (self).f21 else d_1097_v145_)
            nw197_[int(20)] = d_1097_v145_
            nw197_[int(21)] = d_1097_v145_
            nw197_[int(22)] = d_1097_v145_
            nw197_[int(23)] = d_1097_v145_
            nw197_[int(24)] = d_1097_v145_
            nw197_[int(25)] = d_1097_v145_
            d_1099_v147_ = nw197_
            d_1100_v148_: _dafny.Seq
            d_1100_v148_ = _dafny.SeqWithoutIsStrInference([(d_1073_v126_)[default__.safeIndex(351, (d_1073_v126_).length(0))], (d_1075_v127_) + ((d_1073_v126_)[default__.safeIndex(351, (d_1073_v126_).length(0))])])
            d_1101_v149_: _dafny.Seq
            d_1101_v149_ = _dafny.SeqWithoutIsStrInference([d_1094_v142_])
            d_1102_v150_: _dafny.Seq
            d_1102_v150_ = _dafny.SeqWithoutIsStrInference([d_1094_v142_, (d_1101_v149_)[default__.safeIndex(d_1077_v129_, len(d_1101_v149_))], d_1094_v142_, d_1094_v142_, d_1094_v142_])
            d_1103_v151_: C1
            nw198_ = C1()
            nw198_.ctor__((d_1102_v150_)[default__.safeIndex((0) - (p0), len(d_1102_v150_))])
            d_1103_v151_ = nw198_
            d_1104_v152_: D5
            d_1104_v152_ = D5_DC16(d_1103_v151_)
            d_1105_v153_: _dafny.Array
            nw199_ = _dafny.Array(False, 14)
            d_1105_v153_ = nw199_
            rhs144_ = d_1099_v147_
            rhs145_ = (d_1100_v148_ if p2 else d_1100_v148_)
            rhs146_ = D5_DC16(d_1103_v151_)
            rhs147_ = d_1105_v153_
            d_1099_v147_ = rhs144_
            d_1100_v148_ = rhs145_
            d_1104_v152_ = rhs146_
            d_1105_v153_ = rhs147_
        d_1106_v154_: _dafny.Map
        d_1106_v154_ = _dafny.Map({not(p2): (self).f21})
        d_1106_v154_ = (d_1106_v154_).set(not(p2), (self).f21)
        d_1107_v155_: str
        d_1107_v155_ = _dafny.CodePoint('e')
        d_1108_v156_: _dafny.MultiSet
        d_1108_v156_ = _dafny.MultiSet([d_1107_v155_])
        d_1109_v157_: _dafny.Seq
        d_1109_v157_ = _dafny.SeqWithoutIsStrInference([(d_1108_v156_).cardinality])
        d_1110_v158_: D0
        d_1110_v158_ = D0_DC0((self).f21)
        d_1111_v159_: D6
        d_1111_v159_ = D6_DC20(p2, 642, p2)
        d_1112_v160_: _dafny.Array
        nw200_ = _dafny.Array(None, 22)
        nw200_[int(0)] = default__.fm0(p2, d_1107_v155_, True, globalState)
        nw200_[int(1)] = p2
        nw200_[int(2)] = (self).f21
        nw200_[int(3)] = p2
        nw200_[int(4)] = p2
        nw200_[int(5)] = (self).f21
        nw200_[int(6)] = False
        nw200_[int(7)] = (self).f21
        nw200_[int(8)] = (d_1110_v158_).cf0
        nw200_[int(9)] = p2
        nw200_[int(10)] = default__.fm0(False, d_1107_v155_, (d_1111_v159_).cf36, globalState)
        nw200_[int(11)] = (self).f21
        nw200_[int(12)] = p2
        nw200_[int(13)] = (self).f21
        nw200_[int(14)] = (self).f21
        nw200_[int(15)] = not((self).f21)
        nw200_[int(16)] = (self).f21
        nw200_[int(17)] = False
        nw200_[int(18)] = False
        nw200_[int(19)] = False
        nw200_[int(20)] = (self).f21
        nw200_[int(21)] = (self).f21
        d_1112_v160_ = nw200_
        pat_let_tv29_ = p0
        pat_let_tv30_ = p0
        d_1113_v161_: _dafny.Array
        nw201_ = _dafny.Array(None, 19)
        nw201_[int(0)] = p0
        nw201_[int(1)] = p0
        nw201_[int(2)] = p0
        nw201_[int(3)] = p0
        nw201_[int(4)] = p0
        nw201_[int(5)] = p0
        nw201_[int(6)] = 130
        nw201_[int(7)] = p0
        nw201_[int(8)] = p0
        nw201_[int(9)] = p0
        nw201_[int(10)] = 714
        nw201_[int(11)] = len(d_1109_v157_)
        def iife64_(_pat_let15_0):
            def iife65_(d_1114_dt__update__tmp_h1_):
                def iife66_(_pat_let16_0):
                    def iife67_(d_1115_dt__update_hcf15_h0_):
                        def iife68_(_pat_let17_0):
                            def iife69_(d_1116_dt__update_hcf14_h0_):
                                return D2_DC8((d_1114_dt__update__tmp_h1_).cf13, d_1116_dt__update_hcf14_h0_, d_1115_dt__update_hcf15_h0_)
                            return iife69_(_pat_let17_0)
                        return iife68_(pat_let_tv30_)
                    return iife67_(_pat_let16_0)
                return iife66_((0) - (pat_let_tv29_))
            return iife65_(_pat_let15_0)
        nw201_[int(12)] = (iife64_(D2_DC8(d_1112_v160_, p0, len(d_1106_v154_)))).cf15
        nw201_[int(13)] = p0
        nw201_[int(14)] = p0
        nw201_[int(15)] = p0
        nw201_[int(16)] = 20
        nw201_[int(17)] = p0
        nw201_[int(18)] = p0
        d_1113_v161_ = nw201_
        d_1117_v162_: C1
        nw202_ = C1()
        nw202_.ctor__(d_1113_v161_)
        d_1117_v162_ = nw202_
        d_1118_v163_: _dafny.Set
        d_1118_v163_ = _dafny.Set({d_1117_v162_})
        d_1119_v164_: _dafny.Map
        d_1119_v164_ = _dafny.Map({p1: d_1118_v163_})
        r0 = not(((d_1118_v163_) == (((d_1119_v164_)[p1] if (p1) in (d_1119_v164_) else d_1118_v163_))) or ((248) >= (p0)))
        r1 = default__.fm0(p2, d_1107_v155_, (p0) > (p0), globalState)
        r2 = ((0) - ((0) - (default__.safeDivisionInt(p0, p0)))) != (p0)
        r3 = p1
        return r0, r1, r2, r3

    @property
    def f21(self):
        return self._f21

class C3:
    def  __init__(self):
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    def ctor__(self):
        pass
        pass

    def fm6(self, p0, globalState):
        return ((0) - ((len(_dafny.Map({True: True}))) + ((_dafny.MultiSet([(0) - (len(_dafny.Set({len(_dafny.Set({not(True)})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b"))), 697, 758, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dssq")))})))])).cardinality))) == ((414) - (985))

    def fm7(self, p0, p1, p2, globalState):
        source13_ = D0_DC0(True)
        if source13_.is_DC1:
            d_1120___mcc_h0_ = source13_.cf1
            d_1121_cf1_ = d_1120___mcc_h0_
            return d_1121_cf1_
        elif source13_.is_DC0:
            d_1122___mcc_h1_ = source13_.cf0
            d_1123_cf0_ = d_1122___mcc_h1_
            return d_1123_cf0_
        elif True:
            d_1124___mcc_h2_ = source13_.cf2
            d_1125_cf2_ = d_1124___mcc_h2_
            return False

    def m4(self, p0, p1, p2, p3, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        r3: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        (globalState).f11 = True
        d_1126_i0_: int
        d_1126_i0_ = 0
        with _dafny.label("7"):
            while p1:
                with _dafny.c_label("7"):
                    if (d_1126_i0_) >= (100):
                        raise _dafny.Break("7")
                    d_1126_i0_ = (d_1126_i0_) + (1)
                    d_1127_v0_: _dafny.Array
                    def lambda67_(d_1128_p0_):
                        def lambda68_(d_1129_i1_):
                            return (d_1129_i1_) * (d_1128_p0_)

                        return lambda68_

                    init34_ = lambda67_(p0)
                    nw203_ = _dafny.Array(None, 23)
                    for i0_34_ in range(nw203_.length(0)):
                        nw203_[i0_34_] = init34_(i0_34_)
                    d_1127_v0_ = nw203_
                    d_1130_v1_: C1
                    nw204_ = C1()
                    nw204_.ctor__(d_1127_v0_)
                    d_1130_v1_ = nw204_
                    d_1131_v2_: _dafny.Seq
                    d_1131_v2_ = _dafny.SeqWithoutIsStrInference([p2, p2])
                    d_1132_v3_: _dafny.Set
                    d_1132_v3_ = _dafny.Set({p0, p0})
                    d_1133_v4_: _dafny.Seq
                    d_1133_v4_ = _dafny.SeqWithoutIsStrInference([d_1132_v3_])
                    d_1134_v5_: _dafny.Seq
                    d_1134_v5_ = _dafny.SeqWithoutIsStrInference([(d_1131_v2_)[default__.safeIndex(len((d_1133_v4_)[default__.safeIndex(p0, len(d_1133_v4_))]), len(d_1131_v2_))], p2])
                    d_1135_v6_: _dafny.Map
                    d_1135_v6_ = _dafny.Map({(0) - (-984): d_1134_v5_})
                    d_1136_v7_: _dafny.Seq
                    d_1136_v7_ = _dafny.SeqWithoutIsStrInference([p1])
                    index150_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                    (d_1127_v0_)[index150_] = (len(((d_1135_v6_)[p0] if (p0) in (d_1135_v6_) else (d_1134_v5_).set(default__.safeIndex(p0, len(d_1134_v5_)), p2)))) + ((_dafny.MultiSet(d_1136_v7_)).cardinality)
                    index151_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                    (d_1127_v0_)[index151_] = (p0) * (p0)
                    d_1137_v8_: _dafny.Array
                    nw205_ = _dafny.Array(_dafny.CodePoint('D'), 17)
                    d_1137_v8_ = nw205_
                    d_1138_v9_: _dafny.Map
                    d_1138_v9_ = _dafny.Map({d_1137_v8_: not(True)})
                    if ((d_1138_v9_)[d_1137_v8_] if (d_1137_v8_) in (d_1138_v9_) else (p1 if True else p1)):
                        (globalState).f15 = (self).fm7((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], p1, (p0) * (-584), globalState)
                        index152_ = default__.safeIndex(346, (p2).length(0))
                        (p2)[index152_] = (p1 if False else False)
                        d_1139_v10_: _dafny.Seq
                        d_1139_v10_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wngf"))
                        index153_ = default__.safeIndex(346, (p2).length(0))
                        rhs148_ = d_1139_v10_
                        rhs149_ = p1
                        lhs79_ = p2
                        lhs80_ = default__.safeIndex(346, (p2).length(0))
                        r3 = rhs148_
                        lhs79_[lhs80_] = rhs149_
                        d_1140_v11_: _dafny.Map
                        d_1140_v11_ = _dafny.Map({len(d_1139_v10_): p1})
                        d_1141_v12_: _dafny.Map
                        d_1141_v12_ = _dafny.Map({p0: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})
                        d_1142_v13_: _dafny.Map
                        d_1142_v13_ = _dafny.Map({(p2)[default__.safeIndex(346, (p2).length(0))]: ((d_1140_v11_)[p0] if (p0) in (d_1140_v11_) else (self).fm6(d_1141_v12_, globalState))})
                        d_1143_v14_: D6
                        d_1143_v14_ = D6_DC20((d_1139_v10_) < ((d_1139_v10_).set(default__.safeIndex((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], len(d_1139_v10_)), _dafny.CodePoint('v'))), (0) - ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]), ((d_1142_v13_)[(p2)[default__.safeIndex(346, (p2).length(0))]] if ((p2)[default__.safeIndex(346, (p2).length(0))]) in (d_1142_v13_) else p1))
                        d_1143_v14_ = D6_DC20(True, default__.fm9(p0, len(d_1139_v10_), (p2)[default__.safeIndex(346, (p2).length(0))], globalState), p1)
                        (globalState).f11 = p1
                        d_1144_v15_: _dafny.Array
                        def lambda69_(d_1145_p1_):
                            def lambda70_(d_1146_i2_):
                                return d_1145_p1_

                            return lambda70_

                        init35_ = lambda69_(p1)
                        nw206_ = _dafny.Array(None, 25)
                        for i0_35_ in range(nw206_.length(0)):
                            nw206_[i0_35_] = init35_(i0_35_)
                        d_1144_v15_ = nw206_
                        d_1147_v16_: _dafny.Array
                        def lambda71_(d_1148_v10_):
                            def lambda72_(d_1149_i3_):
                                return _dafny.SeqWithoutIsStrInference([len(d_1148_v10_)])

                            return lambda72_

                        init36_ = lambda71_(d_1139_v10_)
                        nw207_ = _dafny.Array(None, 28)
                        for i0_36_ in range(nw207_.length(0)):
                            nw207_[i0_36_] = init36_(i0_36_)
                        d_1147_v16_ = nw207_
                        d_1150_v17_: _dafny.MultiSet
                        d_1150_v17_ = _dafny.MultiSet([p0])
                        index154_ = default__.safeIndex(988, (d_1147_v16_).length(0))
                        (d_1147_v16_)[index154_] = _dafny.SeqWithoutIsStrInference([(d_1150_v17_).cardinality])
                        d_1151_v18_: _dafny.Map
                        d_1151_v18_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([default__.fm16(p1, p0, globalState), default__.fm16((p2)[default__.safeIndex(346, (p2).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cfbueqff"))), globalState)])): d_1144_v15_})
                        d_1152_v19_: _dafny.MultiSet
                        d_1152_v19_ = _dafny.MultiSet([d_1139_v10_, d_1139_v10_, d_1139_v10_, d_1139_v10_, d_1139_v10_])
                        d_1153_v20_: _dafny.Seq
                        d_1153_v20_ = _dafny.SeqWithoutIsStrInference([(d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], p0, (0) - (default__.safeDivisionInt(p0, (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]))])
                        index155_ = default__.safeIndex(988, (d_1147_v16_).length(0))
                        index156_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        rhs150_ = ((d_1151_v18_)[d_1152_v19_] if (d_1152_v19_) in (d_1151_v18_) else d_1144_v15_)
                        rhs151_ = d_1153_v20_
                        rhs152_ = p0
                        lhs81_ = d_1147_v16_
                        lhs82_ = default__.safeIndex(988, (d_1147_v16_).length(0))
                        lhs83_ = d_1127_v0_
                        lhs84_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        d_1144_v15_ = rhs150_
                        lhs81_[lhs82_] = rhs151_
                        lhs83_[lhs84_] = rhs152_
                    elif True:
                        index157_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        (d_1127_v0_)[index157_] = p0
                        d_1154_v21_: C0
                        nw208_ = C0()
                        nw208_.ctor__((p1 if default__.fm0(p1, p3, p1, globalState) else p1))
                        d_1154_v21_ = nw208_
                        d_1155_v22_: _dafny.Seq
                        d_1155_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xergafjq"))
                        r0 = ((_dafny.SeqWithoutIsStrInference([(d_1154_v21_).f23, p1, default__.fm0(p1, (d_1155_v22_)[default__.safeIndex((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], len(d_1155_v22_))], p1, globalState)])) + (d_1136_v7_)).set(default__.safeIndex((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], len((_dafny.SeqWithoutIsStrInference([(d_1154_v21_).f23, p1, default__.fm0(p1, (d_1155_v22_)[default__.safeIndex((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], len(d_1155_v22_))], p1, globalState)])) + (d_1136_v7_))), p1)
                        d_1156_v23_: _dafny.MultiSet
                        d_1156_v23_ = _dafny.MultiSet([len(_dafny.Map({p3: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})), (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], p0])
                        r1 = (((d_1156_v23_)[(d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]] if ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]) in (d_1156_v23_) else p0)) - ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))])
                        d_1157_v24_: _dafny.Map
                        d_1157_v24_ = _dafny.Map({p3: d_1130_v1_})
                        d_1157_v24_ = (d_1157_v24_).set((D4_DC14(p1, (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], d_1132_v3_, _dafny.CodePoint('s'))).cf26, d_1130_v1_)
                    if p1:
                        d_1158_v25_: _dafny.Map
                        d_1158_v25_ = _dafny.Map({p0: d_1132_v3_})
                        r1 = (len((((d_1158_v25_)[(0) - ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))])] if ((0) - ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))])) in (d_1158_v25_) else d_1132_v3_)).intersection(d_1132_v3_))) * ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))])
                        d_1159_v26_: _dafny.MultiSet
                        d_1159_v26_ = _dafny.MultiSet([p1])
                        index158_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        rhs153_ = d_1159_v26_
                        rhs154_ = ((p1) and (not(not(p1))) if p1 else p1)
                        rhs155_ = (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]
                        lhs85_ = globalState
                        lhs86_ = d_1127_v0_
                        lhs87_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        d_1159_v26_ = rhs153_
                        lhs85_.f15 = rhs154_
                        lhs86_[lhs87_] = rhs155_
                        d_1160_v27_: _dafny.MultiSet
                        d_1160_v27_ = _dafny.MultiSet([p0])
                        d_1161_v28_: D6
                        d_1161_v28_ = D6_DC20(not (not(p1)) or (True), ((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]) + ((d_1160_v27_).cardinality), p1)
                        d_1162_v29_: _dafny.Set
                        d_1162_v29_ = _dafny.Set({p1})
                        index159_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        rhs156_ = (d_1161_v28_ if (d_1162_v29_).isdisjoint(d_1162_v29_) else d_1161_v28_)
                        rhs157_ = p1
                        rhs158_ = p1
                        rhs159_ = default__.safeDivisionInt((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))])
                        lhs88_ = globalState
                        lhs89_ = globalState
                        lhs90_ = d_1127_v0_
                        lhs91_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        d_1161_v28_ = rhs156_
                        lhs88_.f15 = rhs157_
                        lhs89_.f11 = rhs158_
                        lhs90_[lhs91_] = rhs159_
                        d_1163_v30_: _dafny.Seq
                        d_1163_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gghwthy"))
                        d_1164_v31_: _dafny.Seq
                        d_1164_v31_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqjkhiid")), (d_1163_v30_).set(default__.safeIndex((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], len(d_1163_v30_)), p3)])
                        d_1165_v32_: _dafny.Map
                        d_1165_v32_ = _dafny.Map({p0: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})
                        d_1164_v31_ = (d_1164_v31_ if (self).fm6(d_1165_v32_, globalState) else _dafny.SeqWithoutIsStrInference([d_1163_v30_]))
                        r1 = 277
                    elif True:
                        index160_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        (d_1127_v0_)[index160_] = (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]
                        d_1166_v33_: _dafny.Map
                        d_1166_v33_ = _dafny.Map({p0: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})
                        d_1167_v35_: _dafny.Seq
                        d_1167_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))
                        d_1168_v36_: _dafny.MultiSet
                        def iife70_():
                            coll34_ = _dafny.Map()
                            compr_34_: int
                            for compr_34_ in _dafny.IntegerRange(617, 957):
                                d_1169_v34_: int = compr_34_
                                if ((617) <= (d_1169_v34_)) and ((d_1169_v34_) < (957)):
                                    coll34_[(d_1169_v34_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uvscymnab"))))] = p0
                            return _dafny.Map(coll34_)
                        d_1168_v36_ = _dafny.MultiSet([d_1166_v33_, iife70_()
                        , (_dafny.Map({len(d_1167_v35_): (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})) | (d_1166_v33_)])
                        d_1170_v37_: _dafny.Map
                        d_1170_v37_ = _dafny.Map({p1: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})
                        d_1171_v38_: _dafny.Seq
                        d_1171_v38_ = _dafny.SeqWithoutIsStrInference([(d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], p0, default__.fm9(len(d_1170_v37_), p0, p1, globalState)])
                        index161_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        rhs160_ = False
                        rhs161_ = (d_1132_v3_).issubset(d_1132_v3_)
                        rhs162_ = (d_1130_v1_).f22
                        rhs163_ = ((d_1168_v36_)[(_dafny.Map({54: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})) | (d_1166_v33_)] if ((_dafny.Map({54: (d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))]})) | (d_1166_v33_)) in (d_1168_v36_) else (d_1171_v38_)[default__.safeIndex((d_1127_v0_)[default__.safeIndex(598, (d_1127_v0_).length(0))], len(d_1171_v38_))])
                        rhs164_ = p0
                        lhs92_ = globalState
                        lhs93_ = globalState
                        lhs94_ = d_1127_v0_
                        lhs95_ = default__.safeIndex(598, (d_1127_v0_).length(0))
                        lhs92_.f11 = rhs160_
                        lhs93_.f11 = rhs161_
                        d_1127_v0_ = rhs162_
                        r1 = rhs163_
                        lhs94_[lhs95_] = rhs164_
                        d_1172_v39_: D9
                        d_1172_v39_ = D9_DC24(d_1171_v38_)
                        rhs165_ = (216) != ((len((d_1172_v39_).cf42)) - (970))
                        rhs166_ = not((False) == (False))
                        lhs96_ = globalState
                        lhs97_ = globalState
                        lhs96_.f15 = rhs165_
                        lhs97_.f11 = rhs166_
                        d_1173_v40_: _dafny.Map
                        d_1173_v40_ = _dafny.Map({d_1171_v38_: p1})
                        d_1173_v40_ = (d_1173_v40_).set(d_1171_v38_, p1)
                        index162_ = default__.safeIndex(936, (p2).length(0))
                        (p2)[index162_] = p1
                        d_1174_v41_: _dafny.Array
                        def lambda73_(d_1175_v35_):
                            def lambda74_(d_1176_i4_):
                                return d_1175_v35_

                            return lambda74_

                        init37_ = lambda73_(d_1167_v35_)
                        nw209_ = _dafny.Array(None, 24)
                        for i0_37_ in range(nw209_.length(0)):
                            nw209_[i0_37_] = init37_(i0_37_)
                        d_1174_v41_ = nw209_
                        index163_ = default__.safeIndex(792, (d_1174_v41_).length(0))
                        (d_1174_v41_)[index163_] = default__.fm16(True, p0, globalState)
                        d_1177_v42_: _dafny.Array
                        nw210_ = _dafny.Array(None, 5)
                        nw210_[int(0)] = (d_1130_v1_).f22
                        nw210_[int(1)] = (d_1130_v1_).f22
                        nw210_[int(2)] = d_1127_v0_
                        nw210_[int(3)] = (d_1130_v1_).f22
                        nw210_[int(4)] = (d_1130_v1_).f22
                        d_1177_v42_ = nw210_
                        index164_ = default__.safeIndex(936, (p2).length(0))
                        index165_ = default__.safeIndex(792, (d_1174_v41_).length(0))
                        rhs167_ = p1
                        rhs168_ = (d_1177_v42_ if p1 else d_1177_v42_)
                        rhs169_ = d_1167_v35_
                        lhs98_ = p2
                        lhs99_ = default__.safeIndex(936, (p2).length(0))
                        lhs100_ = globalState
                        lhs101_ = d_1174_v41_
                        lhs102_ = default__.safeIndex(792, (d_1174_v41_).length(0))
                        lhs98_[lhs99_] = rhs167_
                        lhs100_.f2 = rhs168_
                        lhs101_[lhs102_] = rhs169_
                    pass
            pass
        d_1178_v43_: _dafny.MultiSet
        d_1178_v43_ = _dafny.MultiSet([640])
        d_1178_v43_ = d_1178_v43_
        d_1179_v46_: _dafny.Array
        def lambda75_(d_1180_p0_, d_1181_p1_):
            def lambda76_(d_1182_i5_):
                def iife71_():
                    coll35_ = _dafny.Map()
                    compr_35_: int
                    for compr_35_ in (_dafny.SeqWithoutIsStrInference([d_1180_p0_])).Elements:
                        d_1183_v44_: int = compr_35_
                        if (d_1183_v44_) in (_dafny.SeqWithoutIsStrInference([d_1180_p0_])):
                            coll35_[(d_1183_v44_) * (d_1180_p0_)] = d_1181_p1_
                    return _dafny.Map(coll35_)
                def iife72_():
                    coll36_ = _dafny.Map()
                    compr_36_: int
                    for compr_36_ in (_dafny.MultiSet([d_1180_p0_])).Elements:
                        d_1184_v45_: int = compr_36_
                        if (d_1184_v45_) in (_dafny.MultiSet([d_1180_p0_])):
                            coll36_[(d_1184_v45_) - (d_1180_p0_)] = d_1181_p1_
                    return _dafny.Map(coll36_)
                return (iife71_()
                ) | (iife72_()
                )

            return lambda76_

        init38_ = lambda75_(p0, p1)
        nw211_ = _dafny.Array(None, 23)
        for i0_38_ in range(nw211_.length(0)):
            nw211_[i0_38_] = init38_(i0_38_)
        d_1179_v46_ = nw211_
        nw212_ = _dafny.Array(_dafny.Map({}), 7)
        d_1179_v46_ = nw212_
        d_1185_v47_: _dafny.Map
        d_1185_v47_ = _dafny.Map({p0: default__.fm9(p0, 383, not(p1), globalState)})
        d_1186_v48_: _dafny.Map
        d_1186_v48_ = _dafny.Map({False: p1})
        d_1185_v47_ = (d_1185_v47_).set(len((d_1186_v48_) | (d_1186_v48_)), p0)
        hi8_ = p0
        for d_1187_i6_ in range(p0, hi8_):
            (globalState).f11 = p1
            (globalState).f11 = (p1) and (p1)
            d_1188_v49_: D6
            d_1188_v49_ = D6_DC20(False, p0, p1)
            d_1189_v50_: _dafny.Map
            d_1189_v50_ = _dafny.Map({p0: p1})
            pat_let_tv31_ = p1
            pat_let_tv32_ = d_1189_v50_
            d_1190_v51_: _dafny.Seq
            def iife73_(_pat_let18_0):
                def iife74_(d_1191_dt__update__tmp_h0_):
                    def iife75_(_pat_let19_0):
                        def iife76_(d_1192_dt__update_hcf36_h0_):
                            def iife77_(_pat_let20_0):
                                def iife78_(d_1193_dt__update_hcf37_h0_):
                                    return D6_DC20(d_1192_dt__update_hcf36_h0_, d_1193_dt__update_hcf37_h0_, (d_1191_dt__update__tmp_h0_).cf38)
                                return iife78_(_pat_let20_0)
                            return iife77_(len(pat_let_tv32_))
                        return iife76_(_pat_let19_0)
                    return iife75_(not(pat_let_tv31_))
                return iife74_(_pat_let18_0)
            d_1190_v51_ = _dafny.SeqWithoutIsStrInference([d_1188_v49_, iife73_(d_1188_v49_)])
            d_1194_v52_: str
            d_1194_v52_ = _dafny.CodePoint('j')
            rhs170_ = d_1190_v51_
            rhs171_ = p0
            rhs172_ = (d_1194_v52_ if True else d_1194_v52_)
            d_1190_v51_ = rhs170_
            r1 = rhs171_
            d_1194_v52_ = rhs172_
            d_1195_v53_: D0
            d_1195_v53_ = D0_DC1(True)
            pat_let_tv33_ = p1
            d_1196_v54_: C2
            nw213_ = C2()
            def iife79_(_pat_let21_0):
                def iife80_(d_1197_dt__update__tmp_h1_):
                    def iife81_(_pat_let22_0):
                        def iife82_(d_1198_dt__update_hcf1_h0_):
                            return D0_DC1(d_1198_dt__update_hcf1_h0_)
                        return iife82_(_pat_let22_0)
                    return iife81_(not(pat_let_tv33_))
                return iife80_(_pat_let21_0)
            nw213_.ctor__(iife79_(d_1195_v53_), not((p1) and (p1)))
            d_1196_v54_ = nw213_
        d_1199_v55_: _dafny.Seq
        d_1199_v55_ = _dafny.SeqWithoutIsStrInference([p1, p1, False, True, p1])
        r0 = (d_1199_v55_) + (d_1199_v55_)
        r1 = p0
        d_1200_v56_: _dafny.Seq
        d_1200_v56_ = _dafny.SeqWithoutIsStrInference([d_1186_v48_])
        r2 = _dafny.Map({default__.safeModuloInt(len(d_1200_v56_), p0): (p1) == (not(p1))})
        d_1201_v57_: _dafny.Seq
        d_1201_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "suhxbcocc"))
        r3 = ((d_1201_v57_) + (d_1201_v57_)) + (d_1201_v57_)
        return r0, r1, r2, r3


class C4(T0, T1):
    def  __init__(self):
        self._f19: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f19):
        (self)._f19 = f19

    def fm4(self, p0, p1, p2, p3, globalState):
        return (self).f19

    def m1(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: bool = False
        d_1202_v0_: _dafny.Array
        nw214_ = _dafny.Array(False, 7)
        d_1202_v0_ = nw214_
        d_1203_v1_: str
        d_1203_v1_ = _dafny.CodePoint('u')
        d_1204_v2_: _dafny.Seq
        d_1204_v2_ = _dafny.SeqWithoutIsStrInference([d_1203_v1_])
        d_1205_v3_: _dafny.Map
        d_1205_v3_ = _dafny.Map({d_1204_v2_: d_1202_v0_})
        d_1206_v4_: _dafny.Seq
        d_1206_v4_ = _dafny.SeqWithoutIsStrInference([d_1202_v0_])
        d_1207_v5_: _dafny.Array
        nw215_ = _dafny.Array(None, 19)
        nw215_[int(0)] = d_1202_v0_
        nw215_[int(1)] = d_1202_v0_
        nw215_[int(2)] = d_1202_v0_
        nw215_[int(3)] = (d_1202_v0_ if p0 else d_1202_v0_)
        nw215_[int(4)] = d_1202_v0_
        nw215_[int(5)] = d_1202_v0_
        nw215_[int(6)] = d_1202_v0_
        nw215_[int(7)] = d_1202_v0_
        nw215_[int(8)] = ((d_1205_v3_)[_dafny.SeqWithoutIsStrInference([d_1203_v1_])] if (_dafny.SeqWithoutIsStrInference([d_1203_v1_])) in (d_1205_v3_) else d_1202_v0_)
        nw215_[int(9)] = d_1202_v0_
        nw215_[int(10)] = d_1202_v0_
        nw215_[int(11)] = d_1202_v0_
        nw215_[int(12)] = d_1202_v0_
        nw215_[int(13)] = (d_1206_v4_)[default__.safeIndex(-8, len(d_1206_v4_))]
        nw215_[int(14)] = d_1202_v0_
        nw215_[int(15)] = d_1202_v0_
        nw215_[int(16)] = (D1_DC4(51, p2, d_1202_v0_)).cf6
        nw215_[int(17)] = d_1202_v0_
        nw215_[int(18)] = d_1202_v0_
        d_1207_v5_ = nw215_
        index166_ = default__.safeIndex(100, (d_1207_v5_).length(0))
        (d_1207_v5_)[index166_] = d_1202_v0_
        d_1208_v6_: _dafny.Seq
        d_1208_v6_ = _dafny.SeqWithoutIsStrInference([d_1204_v2_])
        d_1209_v7_: _dafny.MultiSet
        d_1209_v7_ = _dafny.MultiSet([d_1203_v1_])
        d_1210_v8_: _dafny.Map
        d_1210_v8_ = _dafny.Map({((d_1209_v7_)[d_1203_v1_] if (d_1203_v1_) in (d_1209_v7_) else p1): p2})
        index167_ = default__.safeIndex(100, (d_1207_v5_).length(0))
        rhs173_ = d_1202_v0_
        rhs174_ = default__.fm0(not((p1) != (len(_dafny.Map({p0: d_1203_v1_})))), d_1203_v1_, (d_1210_v8_) != (d_1210_v8_), globalState)
        rhs175_ = d_1208_v6_
        lhs103_ = d_1207_v5_
        lhs104_ = default__.safeIndex(100, (d_1207_v5_).length(0))
        lhs105_ = globalState
        lhs103_[lhs104_] = rhs173_
        lhs105_.f15 = rhs174_
        d_1208_v6_ = rhs175_
        if not((self).f19):
            index168_ = default__.safeIndex(100, (d_1207_v5_).length(0))
            (d_1207_v5_)[index168_] = (d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]
            d_1211_v9_: _dafny.MultiSet
            d_1211_v9_ = _dafny.MultiSet([p1, 1])
            (globalState).f11 = (d_1211_v9_).isdisjoint(d_1211_v9_)
            d_1212_v10_: int
            d_1212_v10_ = 938
            d_1213_v11_: _dafny.Seq
            d_1213_v11_ = _dafny.SeqWithoutIsStrInference([p2])
            d_1212_v10_ = default__.fm5(default__.safeDivisionInt(len(d_1213_v11_), p1), globalState)
            d_1204_v2_ = d_1204_v2_
            d_1214_v12_: _dafny.Map
            d_1214_v12_ = _dafny.Map({p3: d_1202_v0_})
            d_1215_v13_: _dafny.Map
            d_1215_v13_ = _dafny.Map({p0: p2})
            d_1214_v12_ = (d_1214_v12_).set(default__.safeModuloInt(len((d_1213_v11_).set(default__.safeIndex(len(d_1215_v13_), len(d_1213_v11_)), p0)), p3), (d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))])
        elif True:
            if (self).f19:
                d_1216_v14_: _dafny.Seq
                d_1216_v14_ = _dafny.SeqWithoutIsStrInference([p1, p1])
                d_1217_v15_: _dafny.Array
                nw216_ = _dafny.Array(None, 11)
                nw216_[int(0)] = p3
                nw216_[int(1)] = (d_1216_v14_)[default__.safeIndex(p3, len(d_1216_v14_))]
                nw216_[int(2)] = p1
                nw216_[int(3)] = p3
                nw216_[int(4)] = p3
                nw216_[int(5)] = p3
                nw216_[int(6)] = p1
                nw216_[int(7)] = p3
                nw216_[int(8)] = p3
                nw216_[int(9)] = p1
                nw216_[int(10)] = len(d_1208_v6_)
                d_1217_v15_ = nw216_
                d_1218_v16_: _dafny.Map
                d_1218_v16_ = _dafny.Map({(p2) or ((self).f19): d_1217_v15_})
                d_1218_v16_ = (d_1218_v16_) | ((d_1218_v16_) | (d_1218_v16_))
                d_1219_v17_: int
                d_1219_v17_ = 919
                d_1220_v18_: _dafny.MultiSet
                d_1220_v18_ = _dafny.MultiSet([d_1207_v5_])
                d_1219_v17_ = ((d_1220_v18_)[d_1207_v5_] if (d_1207_v5_) in (d_1220_v18_) else p3)
                d_1221_v19_: _dafny.Map
                d_1221_v19_ = _dafny.Map({p3: (0) - (p1)})
                d_1221_v19_ = _dafny.Map({d_1219_v17_: default__.safeModuloInt(p1, d_1219_v17_)})
                d_1222_v20_: C3
                nw217_ = C3()
                nw217_.ctor__()
                d_1222_v20_ = nw217_
                d_1223_v21_: _dafny.Seq
                d_1223_v21_ = _dafny.SeqWithoutIsStrInference([p0, p2])
                rhs176_ = not(p2)
                rhs177_ = not ((self).f19) or (p0)
                rhs178_ = (((d_1221_v19_)[(_dafny.MultiSet(d_1223_v21_)).cardinality] if ((_dafny.MultiSet(d_1223_v21_)).cardinality) in (d_1221_v19_) else d_1219_v17_)) > (p3)
                rhs179_ = p2
                rhs180_ = p1
                lhs106_ = globalState
                lhs107_ = globalState
                lhs108_ = globalState
                r1 = rhs176_
                lhs106_.f15 = rhs177_
                lhs107_.f17 = rhs178_
                lhs108_.f17 = rhs179_
                d_1219_v17_ = rhs180_
            elif True:
                d_1224_v22_: _dafny.Array
                nw218_ = _dafny.Array(_dafny.MultiSet({}), 1)
                d_1224_v22_ = nw218_
                index169_ = default__.safeIndex(114, (d_1224_v22_).length(0))
                (d_1224_v22_)[index169_] = _dafny.MultiSet([p0])
                d_1225_v23_: _dafny.MultiSet
                d_1225_v23_ = _dafny.MultiSet([p2])
                d_1226_v24_: D10
                d_1226_v24_ = D10_DC27(d_1225_v23_)
                index170_ = default__.safeIndex(114, (d_1224_v22_).length(0))
                (d_1224_v22_)[index170_] = ((d_1226_v24_).cf48 if (self).f19 else d_1225_v23_)
                d_1227_v25_: _dafny.Map
                d_1227_v25_ = _dafny.Map({p2: default__.safeDivisionInt(len(d_1204_v2_), (0) - (len(d_1210_v8_)))})
                d_1228_v26_: _dafny.MultiSet
                d_1228_v26_ = _dafny.MultiSet([p1])
                d_1227_v25_ = (d_1227_v25_).set(not(((d_1228_v26_).cardinality) <= (p3)), (0) - (len(d_1227_v25_)))
                d_1210_v8_ = (d_1210_v8_).set(p1, p2)
                r0 = (self).f19
                d_1229_v27_: _dafny.Array
                nw219_ = _dafny.Array(None, 21)
                d_1229_v27_ = nw219_
                d_1229_v27_ = d_1229_v27_
            d_1230_v28_: _dafny.Map
            d_1230_v28_ = _dafny.Map({d_1203_v1_: p3})
            d_1231_v29_: _dafny.Set
            d_1231_v29_ = _dafny.Set({default__.fm9(p1, p3, p2, globalState)})
            d_1232_v30_: D4
            d_1232_v30_ = D4_DC14(p0, ((d_1230_v28_)[d_1203_v1_] if (d_1203_v1_) in (d_1230_v28_) else p3), d_1231_v29_, d_1203_v1_)
            d_1233_v31_: D6
            d_1233_v31_ = D6_DC20(default__.fm0(p0, (d_1232_v30_).cf26, p0, globalState), p1, (self).f19)
            def iife83_(_pat_let23_0):
                def iife84_(d_1234_dt__update__tmp_h0_):
                    def iife85_(_pat_let24_0):
                        def iife86_(d_1235_dt__update_hcf37_h0_):
                            return D6_DC20((d_1234_dt__update__tmp_h0_).cf36, d_1235_dt__update_hcf37_h0_, (d_1234_dt__update__tmp_h0_).cf38)
                        return iife86_(_pat_let24_0)
                    return iife85_(115)
                return iife84_(_pat_let23_0)
            if not((iife83_(d_1233_v31_)).cf38):
                d_1236_v32_: int
                d_1236_v32_ = 203
                d_1237_v33_: _dafny.Seq
                d_1237_v33_ = _dafny.SeqWithoutIsStrInference([p3])
                d_1236_v32_ = (d_1237_v33_)[default__.safeIndex(p3, len(d_1237_v33_))]
                r0 = p2
                (globalState).f15 = not(p2)
                d_1238_v34_: _dafny.Array
                nw220_ = _dafny.Array(int(0), 16)
                d_1238_v34_ = nw220_
                index171_ = default__.safeIndex(446, (d_1238_v34_).length(0))
                (d_1238_v34_)[index171_] = p3
                d_1239_v35_: D1
                d_1239_v35_ = D1_DC4((0) - (p1), (self).f19, d_1202_v0_)
                d_1240_v36_: _dafny.Set
                d_1240_v36_ = _dafny.Set({False, (self).f19, (self).f19, p2, (self).f19})
                d_1241_v37_: D2
                d_1241_v37_ = D2_DC8((d_1239_v35_).cf6, 770, len(d_1240_v36_))
                index172_ = default__.safeIndex(446, (d_1238_v34_).length(0))
                rhs181_ = d_1236_v32_
                rhs182_ = (d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]
                rhs183_ = (d_1241_v37_).cf13
                lhs109_ = d_1238_v34_
                lhs110_ = default__.safeIndex(446, (d_1238_v34_).length(0))
                lhs109_[lhs110_] = rhs181_
                d_1202_v0_ = rhs182_
                d_1202_v0_ = rhs183_
                d_1242_v38_: D1
                d_1242_v38_ = D1_DC5(d_1204_v2_, (0) - (p3), p3)
                d_1243_v39_: _dafny.Map
                d_1243_v39_ = _dafny.Map({(self).f19: d_1204_v2_})
                d_1244_v40_: _dafny.Array
                nw221_ = _dafny.Array(None, 28)
                nw221_[int(0)] = d_1204_v2_
                nw221_[int(1)] = d_1204_v2_
                nw221_[int(2)] = d_1204_v2_
                nw221_[int(3)] = (d_1204_v2_ if True else d_1204_v2_)
                nw221_[int(4)] = (d_1204_v2_) + (d_1204_v2_)
                nw221_[int(5)] = d_1204_v2_
                nw221_[int(6)] = (d_1204_v2_) + ((d_1242_v38_).cf7)
                nw221_[int(7)] = (d_1204_v2_) + (d_1204_v2_)
                nw221_[int(8)] = (d_1208_v6_)[default__.safeIndex(p3, len(d_1208_v6_))]
                nw221_[int(9)] = (d_1242_v38_).cf7
                nw221_[int(10)] = d_1204_v2_
                nw221_[int(11)] = (d_1204_v2_) + ((d_1208_v6_)[default__.safeIndex(p1, len(d_1208_v6_))])
                nw221_[int(12)] = d_1204_v2_
                nw221_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpbix"))
                nw221_[int(14)] = d_1204_v2_
                nw221_[int(15)] = d_1204_v2_
                nw221_[int(16)] = (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1245_i0_ in range(default__.abs(49))])) + (d_1204_v2_)
                nw221_[int(17)] = ((_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1246_i1_ in range(default__.abs(210))])).set(default__.safeIndex(398, len(_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1247_i1_ in range(default__.abs(210))]))), d_1203_v1_)) + (((d_1243_v39_)[p2] if (p2) in (d_1243_v39_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "noveiuir"))))
                nw221_[int(18)] = ((default__.fm16(not(p2), (d_1238_v34_)[default__.safeIndex(446, (d_1238_v34_).length(0))], globalState)) + (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1248_i2_ in range(default__.abs(129))]))).set(default__.safeIndex(p3, len((default__.fm16(not(p2), (d_1238_v34_)[default__.safeIndex(446, (d_1238_v34_).length(0))], globalState)) + (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1249_i2_ in range(default__.abs(129))])))), d_1203_v1_)
                nw221_[int(19)] = (d_1204_v2_ if p2 else _dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1250_i3_ in range(default__.abs(898))]))
                nw221_[int(20)] = d_1204_v2_
                nw221_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vafiqo"))
                nw221_[int(22)] = d_1204_v2_
                nw221_[int(23)] = _dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1251_i4_ in range(default__.abs(-200))])
                nw221_[int(24)] = (d_1204_v2_) + (_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1252_i5_ in range(default__.abs(717))]))
                nw221_[int(25)] = (d_1204_v2_) + (d_1204_v2_)
                nw221_[int(26)] = d_1204_v2_
                nw221_[int(27)] = d_1204_v2_
                d_1244_v40_ = nw221_
                index173_ = default__.safeIndex(469, (d_1244_v40_).length(0))
                (d_1244_v40_)[index173_] = d_1204_v2_
                d_1253_v41_: _dafny.MultiSet
                d_1253_v41_ = _dafny.MultiSet([p2, ((d_1238_v34_)[default__.safeIndex(446, (d_1238_v34_).length(0))]) >= (d_1236_v32_), p2])
                d_1254_v42_: C0
                nw222_ = C0()
                nw222_.ctor__(p0)
                d_1254_v42_ = nw222_
                d_1255_v43_: D9
                d_1255_v43_ = D9_DC25(p3, p0, p1, d_1254_v42_)
                d_1256_v44_: _dafny.MultiSet
                d_1256_v44_ = _dafny.MultiSet([(d_1255_v43_).cf45, len(default__.fm16((d_1254_v42_).f23, (0) - (-311), globalState))])
                index174_ = default__.safeIndex(469, (d_1244_v40_).length(0))
                index175_ = default__.safeIndex(446, (d_1238_v34_).length(0))
                rhs184_ = d_1240_v36_
                rhs185_ = d_1204_v2_
                rhs186_ = default__.safeModuloInt(default__.safeDivisionInt((d_1256_v44_).cardinality, len(d_1204_v2_)), len(d_1231_v29_))
                rhs187_ = (d_1253_v41_).set((d_1240_v36_).ispropersubset(d_1240_v36_), default__.abs(d_1236_v32_))
                rhs188_ = ((d_1210_v8_)[((d_1238_v34_)[default__.safeIndex(446, (d_1238_v34_).length(0))]) + (p1)] if (((d_1238_v34_)[default__.safeIndex(446, (d_1238_v34_).length(0))]) + (p1)) in (d_1210_v8_) else p0)
                lhs111_ = d_1244_v40_
                lhs112_ = default__.safeIndex(469, (d_1244_v40_).length(0))
                lhs113_ = d_1238_v34_
                lhs114_ = default__.safeIndex(446, (d_1238_v34_).length(0))
                lhs115_ = globalState
                d_1240_v36_ = rhs184_
                lhs111_[lhs112_] = rhs185_
                lhs113_[lhs114_] = rhs186_
                d_1253_v41_ = rhs187_
                lhs115_.f15 = rhs188_
            elif True:
                d_1257_v45_: D4
                d_1257_v45_ = D4_DC13(p1)
                d_1258_v46_: _dafny.Map
                d_1258_v46_ = _dafny.Map({(self).f19: d_1257_v45_})
                d_1258_v46_ = (d_1258_v46_).set(p0, d_1257_v45_)
                d_1259_v47_: C3
                nw223_ = C3()
                nw223_.ctor__()
                d_1259_v47_ = nw223_
                d_1260_v48_: int
                d_1260_v48_ = 1
                d_1261_v49_: T1
                nw224_ = C2()
                nw224_.ctor__(D0_DC1(p0), p2)
                d_1261_v49_ = nw224_
                rhs189_ = d_1260_v48_
                rhs190_ = d_1261_v49_
                rhs191_ = (self).f19
                lhs116_ = globalState
                d_1260_v48_ = rhs189_
                d_1261_v49_ = rhs190_
                lhs116_.f17 = rhs191_
                d_1262_v50_: _dafny.Set
                d_1262_v50_ = _dafny.Set({(self).f19, p0, p0})
                d_1263_v51_: _dafny.Array
                nw225_ = _dafny.Array(None, 27)
                nw225_[int(0)] = d_1260_v48_
                nw225_[int(1)] = d_1260_v48_
                nw225_[int(2)] = p1
                nw225_[int(3)] = p1
                nw225_[int(4)] = d_1260_v48_
                nw225_[int(5)] = 242
                nw225_[int(6)] = p3
                nw225_[int(7)] = p1
                nw225_[int(8)] = d_1260_v48_
                nw225_[int(9)] = p3
                nw225_[int(10)] = d_1260_v48_
                nw225_[int(11)] = (0) - (d_1260_v48_)
                nw225_[int(12)] = p1
                nw225_[int(13)] = d_1260_v48_
                nw225_[int(14)] = d_1260_v48_
                nw225_[int(15)] = p3
                nw225_[int(16)] = len(d_1204_v2_)
                nw225_[int(17)] = d_1260_v48_
                nw225_[int(18)] = len(d_1262_v50_)
                nw225_[int(19)] = len(_dafny.SeqWithoutIsStrInference([p0, False]))
                nw225_[int(20)] = p1
                nw225_[int(21)] = p3
                nw225_[int(22)] = d_1260_v48_
                nw225_[int(23)] = p1
                nw225_[int(24)] = p3
                nw225_[int(25)] = p1
                nw225_[int(26)] = p3
                d_1263_v51_ = nw225_
                d_1264_v52_: C1
                nw226_ = C1()
                nw226_.ctor__(d_1263_v51_)
                d_1264_v52_ = nw226_
                d_1265_v53_: _dafny.Seq
                d_1265_v53_ = _dafny.SeqWithoutIsStrInference([d_1264_v52_])
                d_1260_v48_ = len(d_1265_v53_)
                d_1210_v8_ = (d_1210_v8_).set(d_1260_v48_, default__.fm0((self).f19, _dafny.CodePoint('a'), (self).f19, globalState))
            d_1204_v2_ = d_1204_v2_
            d_1203_v1_ = _dafny.CodePoint('x')
            if ((0) - ((0) - (p3))) >= (p3):
                d_1266_v54_: _dafny.Array
                def lambda77_(d_1267_v1_):
                    def lambda78_(d_1268_i6_):
                        return d_1267_v1_

                    return lambda78_

                init39_ = lambda77_(d_1203_v1_)
                nw227_ = _dafny.Array(None, 20)
                for i0_39_ in range(nw227_.length(0)):
                    nw227_[i0_39_] = init39_(i0_39_)
                d_1266_v54_ = nw227_
                d_1269_v55_: _dafny.Map
                d_1269_v55_ = _dafny.Map({d_1266_v54_: default__.fm5(p3, globalState)})
                d_1270_v56_: _dafny.Map
                d_1270_v56_ = _dafny.Map({False: p3})
                d_1271_v57_: _dafny.Map
                d_1271_v57_ = _dafny.Map({((d_1270_v56_)[(self).f19] if ((self).f19) in (d_1270_v56_) else p1): d_1266_v54_})
                d_1272_v58_: _dafny.Seq
                d_1272_v58_ = _dafny.SeqWithoutIsStrInference([False])
                d_1269_v55_ = (d_1269_v55_).set(((d_1271_v57_)[p1] if (p1) in (d_1271_v57_) else d_1266_v54_), len(d_1272_v58_))
                d_1273_v59_: _dafny.Seq
                d_1273_v59_ = _dafny.SeqWithoutIsStrInference([p3, p1, p3])
                d_1274_v60_: _dafny.Map
                d_1274_v60_ = _dafny.Map({(self).f19: (d_1273_v59_).set(default__.safeIndex(p3, len(d_1273_v59_)), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbydlk"))))})
                d_1275_v61_: _dafny.Map
                d_1275_v61_ = _dafny.Map({704: len(((d_1274_v60_)[p2] if (p2) in (d_1274_v60_) else _dafny.SeqWithoutIsStrInference([len(d_1204_v2_) for d_1276_i7_ in range(default__.abs(975))])))})
                d_1277_v62_: _dafny.Set
                d_1277_v62_ = _dafny.Set({p2})
                d_1278_v63_: _dafny.Map
                d_1278_v63_ = _dafny.Map({True: p2})
                d_1279_v64_: D9
                d_1279_v64_ = D9_DC24(default__.fm19(d_1273_v59_, ((d_1278_v63_)[p0] if (p0) in (d_1278_v63_) else (d_1272_v58_)[default__.safeIndex(p3, len(d_1272_v58_))]), globalState))
                d_1280_v65_: D9
                d_1280_v65_ = D9_DC26(d_1279_v64_)
                d_1281_v66_: _dafny.MultiSet
                d_1281_v66_ = _dafny.MultiSet([d_1280_v65_, d_1280_v65_])
                d_1282_v67_: _dafny.Array
                nw228_ = _dafny.Array(None, 19)
                nw228_[int(0)] = len(d_1272_v58_)
                nw228_[int(1)] = p3
                nw228_[int(2)] = p1
                nw228_[int(3)] = 870
                nw228_[int(4)] = p1
                nw228_[int(5)] = p1
                nw228_[int(6)] = (0) - (p3)
                nw228_[int(7)] = p3
                nw228_[int(8)] = 738
                nw228_[int(9)] = p3
                nw228_[int(10)] = p3
                nw228_[int(11)] = p1
                nw228_[int(12)] = p3
                nw228_[int(13)] = p3
                nw228_[int(14)] = -922
                nw228_[int(15)] = p1
                nw228_[int(16)] = p1
                nw228_[int(17)] = p1
                nw228_[int(18)] = p1
                d_1282_v67_ = nw228_
                d_1283_v68_: _dafny.Map
                d_1283_v68_ = _dafny.Map({len(d_1273_v59_): d_1282_v67_})
                d_1284_v69_: _dafny.Seq
                d_1284_v69_ = _dafny.SeqWithoutIsStrInference([((d_1283_v68_)[default__.fm9(len(d_1231_v29_), p1, (self).f19, globalState)] if (default__.fm9(len(d_1231_v29_), p1, (self).f19, globalState)) in (d_1283_v68_) else d_1282_v67_)])
                d_1285_v70_: _dafny.Array
                nw229_ = _dafny.Array(None, 23)
                nw229_[int(0)] = (0) - ((((d_1275_v61_)[p1] if (p1) in (d_1275_v61_) else p3)) * (-157))
                nw229_[int(1)] = p1
                nw229_[int(2)] = (len(d_1231_v29_)) * (p3)
                nw229_[int(3)] = p1
                nw229_[int(4)] = default__.fm9(p1, 81, p2, globalState)
                nw229_[int(5)] = p1
                nw229_[int(6)] = default__.safeDivisionInt(p3, p3)
                nw229_[int(7)] = -328
                nw229_[int(8)] = p1
                nw229_[int(9)] = (default__.fm9(p1, p3, p2, globalState) if p0 else len(d_1277_v62_))
                nw229_[int(10)] = len((d_1270_v56_ if p2 else _dafny.Map({p0: -174})))
                nw229_[int(11)] = (p1) - (p1)
                nw229_[int(12)] = p3
                nw229_[int(13)] = p1
                nw229_[int(14)] = 258
                nw229_[int(15)] = (len(_dafny.Map({d_1281_v66_: p1}))) * (p1)
                nw229_[int(16)] = default__.fm9(p1, (0) - (p3), (self).f19, globalState)
                nw229_[int(17)] = default__.safeDivisionInt(p1, 589)
                nw229_[int(18)] = p1
                nw229_[int(19)] = (730) + (p1)
                nw229_[int(20)] = default__.safeDivisionInt((0) - (p1), p3)
                nw229_[int(21)] = len((d_1284_v69_) + (d_1284_v69_))
                nw229_[int(22)] = p1
                d_1285_v70_ = nw229_
                index176_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                (d_1282_v67_)[index176_] = p1
                d_1286_v71_: int
                d_1286_v71_ = 358
                index177_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                rhs192_ = not(p2)
                rhs193_ = p3
                rhs194_ = d_1286_v71_
                lhs117_ = globalState
                lhs118_ = d_1282_v67_
                lhs119_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                lhs117_.f11 = rhs192_
                lhs118_[lhs119_] = rhs193_
                d_1286_v71_ = rhs194_
                d_1287_v72_: D4
                d_1287_v72_ = D4_DC13(len(default__.fm20(globalState)))
                d_1288_v73_: _dafny.Map
                d_1288_v73_ = _dafny.Map({(d_1272_v58_)[default__.safeIndex(863, len(d_1272_v58_))]: d_1287_v72_})
                d_1288_v73_ = (d_1288_v73_).set(True, d_1287_v72_)
                d_1289_v74_: C1
                nw230_ = C1()
                nw230_.ctor__(d_1282_v67_)
                d_1289_v74_ = nw230_
                d_1290_v75_: _dafny.Seq
                d_1290_v75_ = _dafny.SeqWithoutIsStrInference([d_1289_v74_])
                d_1291_v76_: _dafny.MultiSet
                d_1291_v76_ = _dafny.MultiSet([not(p2), p2, not(False), False, (self).f19])
                index178_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                index179_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                rhs195_ = (0) - ((p3) * (len((d_1290_v75_) + (d_1290_v75_))))
                rhs196_ = ((0) - (((d_1291_v76_)[p0] if (p0) in (d_1291_v76_) else d_1286_v71_))) + ((d_1282_v67_)[default__.safeIndex(357, (d_1282_v67_).length(0))])
                lhs120_ = d_1282_v67_
                lhs121_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                lhs122_ = d_1282_v67_
                lhs123_ = default__.safeIndex(357, (d_1282_v67_).length(0))
                lhs120_[lhs121_] = rhs195_
                lhs122_[lhs123_] = rhs196_
                r0 = (p3) == (d_1286_v71_)
            elif True:
                d_1292_v78_: _dafny.Map
                d_1292_v78_ = _dafny.Map({(_dafny.MultiSet(d_1204_v2_)).cardinality: p1})
                d_1293_v79_: _dafny.Map
                def iife87_():
                    coll37_ = _dafny.Map()
                    compr_37_: int
                    for compr_37_ in (d_1292_v78_).keys.Elements:
                        d_1294_v77_: int = compr_37_
                        if (d_1294_v77_) in (d_1292_v78_):
                            coll37_[default__.safeDivisionInt(d_1294_v77_, p1)] = p3
                    return _dafny.Map(coll37_)
                d_1293_v79_ = _dafny.Map({(d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]: iife87_()
                })
                d_1204_v2_ = (default__.fm16(p2, len(d_1293_v79_), globalState)).set(default__.safeIndex(p3, len(default__.fm16(p2, len(d_1293_v79_), globalState))), (d_1204_v2_)[default__.safeIndex(p1, len(d_1204_v2_))])
                d_1295_v80_: D0
                d_1295_v80_ = D0_DC1(p2)
                d_1296_v81_: C2
                nw231_ = C2()
                nw231_.ctor__(d_1295_v80_, p2)
                d_1296_v81_ = nw231_
                d_1297_v82_: _dafny.Seq
                d_1297_v82_ = _dafny.SeqWithoutIsStrInference([(d_1296_v81_).f21])
                d_1298_v83_: _dafny.Set
                d_1298_v83_ = _dafny.Set({d_1297_v82_, d_1297_v82_, d_1297_v82_})
                nw232_ = C2()
                nw232_.ctor__(d_1295_v80_, (d_1298_v83_).issubset(d_1298_v83_))
                d_1296_v81_ = nw232_
                d_1299_v84_: C0
                nw233_ = C0()
                nw233_.ctor__(not((self).f19))
                d_1299_v84_ = nw233_
                index180_ = default__.safeIndex(758, (d_1202_v0_).length(0))
                (d_1202_v0_)[index180_] = (d_1296_v81_).f21
                d_1300_v85_: int
                d_1300_v85_ = -626
                d_1301_v86_: _dafny.Set
                d_1301_v86_ = _dafny.Set({p0})
                index181_ = default__.safeIndex(758, (d_1202_v0_).length(0))
                rhs197_ = (len(d_1301_v86_)) >= (default__.safeDivisionInt(p3, d_1300_v85_))
                rhs198_ = p1
                lhs124_ = d_1202_v0_
                lhs125_ = default__.safeIndex(758, (d_1202_v0_).length(0))
                lhs124_[lhs125_] = rhs197_
                d_1300_v85_ = rhs198_
                r0 = ((_dafny.Set({(d_1299_v84_).f23})) | (d_1301_v86_)).issubset((d_1301_v86_) - (d_1301_v86_))
        d_1302_v87_: _dafny.Array
        nw234_ = _dafny.Array(_dafny.Set({}), 26)
        d_1302_v87_ = nw234_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1302_v87_).length(0)):
            d_1303_i8_: int = guard_loop_2_
            if (True) and (((0) <= (d_1303_i8_)) and ((d_1303_i8_) < ((d_1302_v87_).length(0)))):
                (d_1302_v87_)[(d_1303_i8_)] = (D11_DC29(_dafny.Set({(self).f19}))).cf51
        arr0_ = (d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]
        index182_ = default__.safeIndex(820, ((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]).length(0))
        arr0_[index182_] = (len(d_1204_v2_)) == (p1)
        arr1_ = (d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]
        index183_ = default__.safeIndex(820, ((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]).length(0))
        arr1_[index183_] = p2
        hi9_ = (818) + (162)
        for d_1304_i9_ in range(p3, hi9_):
            d_1305_v88_: _dafny.Array
            nw235_ = _dafny.Array(int(0), 29)
            d_1305_v88_ = nw235_
            d_1306_v89_: _dafny.Seq
            d_1306_v89_ = _dafny.SeqWithoutIsStrInference([p0, default__.fm0(((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))])[default__.safeIndex(820, ((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]).length(0))], _dafny.CodePoint('a'), p0, globalState)])
            index184_ = default__.safeIndex(780, (d_1305_v88_).length(0))
            (d_1305_v88_)[index184_] = default__.safeModuloInt(len(d_1306_v89_), p3)
            index185_ = default__.safeIndex(780, (d_1305_v88_).length(0))
            (d_1305_v88_)[index185_] = (p1) - ((p1 if p2 else d_1304_i9_))
            d_1307_v90_: _dafny.Map
            d_1307_v90_ = _dafny.Map({p0: d_1304_i9_})
            d_1308_v91_: _dafny.Seq
            d_1308_v91_ = _dafny.SeqWithoutIsStrInference([p3])
            d_1309_v92_: D1
            d_1309_v92_ = D1_DC5(_dafny.SeqWithoutIsStrInference([d_1203_v1_ for d_1310_i10_ in range(default__.abs(-248))]), (d_1305_v88_)[default__.safeIndex(780, (d_1305_v88_).length(0))], p1)
            d_1311_v93_: D3
            d_1311_v93_ = D3_DC11(d_1307_v90_, default__.safeModuloInt((d_1308_v91_)[default__.safeIndex(len((d_1309_v92_).cf7), len(d_1308_v91_))], p1), d_1305_v88_)
            pat_let_tv34_ = d_1307_v90_
            def iife88_(_pat_let25_0):
                def iife89_(d_1312_dt__update__tmp_h1_):
                    def iife90_(_pat_let26_0):
                        def iife91_(d_1313_dt__update_hcf18_h0_):
                            return D3_DC11(d_1313_dt__update_hcf18_h0_, (d_1312_dt__update__tmp_h1_).cf19, (d_1312_dt__update__tmp_h1_).cf20)
                        return iife91_(_pat_let26_0)
                    return iife90_(pat_let_tv34_)
                return iife89_(_pat_let25_0)
            d_1311_v93_ = iife88_(D3_DC11(d_1307_v90_, len(_dafny.Map({p2: (d_1305_v88_)[default__.safeIndex(780, (d_1305_v88_).length(0))]})), d_1305_v88_))
            d_1308_v91_ = d_1308_v91_
            d_1314_v94_: _dafny.Map
            d_1314_v94_ = _dafny.Map({(self).f19: d_1203_v1_})
            d_1314_v94_ = d_1314_v94_
        d_1315_v95_: _dafny.Seq
        d_1315_v95_ = _dafny.SeqWithoutIsStrInference([(self).f19])
        d_1316_v96_: _dafny.Map
        d_1316_v96_ = _dafny.Map({d_1315_v95_: p1})
        d_1317_i11_: int
        d_1317_i11_ = 0
        with _dafny.label("8"):
            while (_dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): p3})) != (d_1316_v96_):
                with _dafny.c_label("8"):
                    if (d_1317_i11_) >= (100):
                        raise _dafny.Break("8")
                    d_1317_i11_ = (d_1317_i11_) + (1)
                    d_1315_v95_ = d_1315_v95_
                    d_1318_v97_: _dafny.Set
                    d_1318_v97_ = _dafny.Set({p0})
                    d_1319_v98_: _dafny.MultiSet
                    d_1319_v98_ = _dafny.MultiSet([(self).f19])
                    d_1320_v99_: _dafny.Array
                    nw236_ = _dafny.Array(None, 9)
                    nw236_[int(0)] = p1
                    nw236_[int(1)] = p1
                    nw236_[int(2)] = (d_1209_v7_).cardinality
                    nw236_[int(3)] = default__.fm5(p1, globalState)
                    nw236_[int(4)] = p1
                    nw236_[int(5)] = (p3) + (default__.fm9(len(d_1318_v97_), p1, (self).f19, globalState))
                    nw236_[int(6)] = ((_dafny.MultiSet([p0])) | (d_1319_v98_)).cardinality
                    nw236_[int(7)] = p3
                    nw236_[int(8)] = p1
                    d_1320_v99_ = nw236_
                    index186_ = default__.safeIndex(552, (d_1320_v99_).length(0))
                    (d_1320_v99_)[index186_] = (655) - (p3)
                    index187_ = default__.safeIndex(881, (d_1320_v99_).length(0))
                    (d_1320_v99_)[index187_] = p3
                    d_1321_v100_: D5
                    d_1321_v100_ = D5_DC17(p2, D2_DC8((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))], p1, -273), True, (self).f19, p1)
                    index188_ = default__.safeIndex(552, (d_1320_v99_).length(0))
                    index189_ = default__.safeIndex(881, (d_1320_v99_).length(0))
                    rhs199_ = not(((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))])[default__.safeIndex(820, ((d_1207_v5_)[default__.safeIndex(100, (d_1207_v5_).length(0))]).length(0))])
                    rhs200_ = (d_1321_v100_).cf33
                    rhs201_ = (((0) - (len(d_1315_v95_))) * (p1)) - (p1)
                    lhs126_ = d_1320_v99_
                    lhs127_ = default__.safeIndex(552, (d_1320_v99_).length(0))
                    lhs128_ = d_1320_v99_
                    lhs129_ = default__.safeIndex(881, (d_1320_v99_).length(0))
                    r1 = rhs199_
                    lhs126_[lhs127_] = rhs200_
                    lhs128_[lhs129_] = rhs201_
                    d_1322_v101_: C3
                    nw237_ = C3()
                    nw237_.ctor__()
                    d_1322_v101_ = nw237_
                    d_1323_v102_: _dafny.Seq
                    d_1323_v102_ = _dafny.SeqWithoutIsStrInference([d_1320_v99_, d_1320_v99_])
                    d_1323_v102_ = ((d_1323_v102_) + (d_1323_v102_)) + (d_1323_v102_)
                    pass
            pass
        r0 = (self).f19
        d_1324_v103_: C3
        nw238_ = C3()
        nw238_.ctor__()
        d_1324_v103_ = nw238_
        d_1325_v104_: _dafny.Seq
        d_1325_v104_ = _dafny.SeqWithoutIsStrInference([d_1324_v103_])
        r1 = default__.fm0((_dafny.MultiSet([d_1324_v103_])).issubset(_dafny.MultiSet(d_1325_v104_)), (d_1203_v1_ if True else default__.fm12(p3, d_1204_v2_, d_1204_v2_, globalState)), p2, globalState)
        return r0, r1

    def m2(self, globalState):
        d_1326_v0_: _dafny.MultiSet
        d_1326_v0_ = _dafny.MultiSet([(self).f19])
        d_1327_v1_: int
        d_1327_v1_ = 931
        d_1328_v2_: D10
        d_1328_v2_ = D10_DC27((d_1326_v0_) | ((d_1326_v0_).set(not((self).f19), default__.abs(d_1327_v1_))))
        d_1329_v3_: _dafny.Seq
        d_1329_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yrd"))
        d_1330_v4_: str
        d_1330_v4_ = _dafny.CodePoint('s')
        d_1331_v5_: _dafny.Map
        d_1331_v5_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbgn")): d_1330_v4_})
        d_1332_v6_: _dafny.Seq
        d_1332_v6_ = _dafny.SeqWithoutIsStrInference([not(default__.fm0((self).f19, d_1330_v4_, (self).f19, globalState))])
        d_1333_v7_: _dafny.Array
        def lambda79_(d_1334_v1_):
            def lambda80_(d_1335_i1_):
                return (d_1335_i1_) + (d_1334_v1_)

            return lambda80_

        init40_ = lambda79_(d_1327_v1_)
        nw239_ = _dafny.Array(None, 28)
        for i0_40_ in range(nw239_.length(0)):
            nw239_[i0_40_] = init40_(i0_40_)
        d_1333_v7_ = nw239_
        d_1336_v8_: T0
        nw240_ = C1()
        nw240_.ctor__(d_1333_v7_)
        d_1336_v8_ = nw240_
        d_1337_v9_: _dafny.Seq
        d_1337_v9_ = _dafny.SeqWithoutIsStrInference([d_1336_v8_, d_1336_v8_])
        d_1338_v10_: _dafny.Set
        d_1338_v10_ = _dafny.Set({(self).f19})
        d_1339_v11_: _dafny.Map
        d_1339_v11_ = _dafny.Map({d_1327_v1_: d_1338_v10_})
        d_1340_v12_: _dafny.Array
        nw241_ = _dafny.Array(None, 20)
        nw241_[int(0)] = d_1330_v4_
        nw241_[int(1)] = d_1330_v4_
        nw241_[int(2)] = (((d_1331_v5_)[_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apnh"))] if (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "apnh"))) in (d_1331_v5_) else d_1330_v4_) if (self).f19 else d_1330_v4_)
        nw241_[int(3)] = d_1330_v4_
        nw241_[int(4)] = d_1330_v4_
        nw241_[int(5)] = d_1330_v4_
        nw241_[int(6)] = d_1330_v4_
        nw241_[int(7)] = default__.fm12(len(d_1332_v6_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('e') for d_1341_i0_ in range(default__.abs(223))]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxjyg")), globalState)
        nw241_[int(8)] = d_1330_v4_
        nw241_[int(9)] = (d_1329_v3_)[default__.safeIndex(default__.fm5(len(d_1337_v9_), globalState), len(d_1329_v3_))]
        nw241_[int(10)] = d_1330_v4_
        nw241_[int(11)] = d_1330_v4_
        nw241_[int(12)] = d_1330_v4_
        nw241_[int(13)] = default__.fm12(d_1327_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnd")), d_1329_v3_, globalState)
        nw241_[int(14)] = d_1330_v4_
        nw241_[int(15)] = d_1330_v4_
        nw241_[int(16)] = d_1330_v4_
        nw241_[int(17)] = default__.fm12(d_1327_v1_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "op")), d_1329_v3_, globalState)
        nw241_[int(18)] = d_1330_v4_
        nw241_[int(19)] = default__.fm12(len(d_1339_v11_), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hdoqfle")), d_1329_v3_, globalState)
        d_1340_v12_ = nw241_
        rhs202_ = D10_DC27(d_1326_v0_)
        rhs203_ = ((d_1326_v0_)[(self).f19] if ((self).f19) in (d_1326_v0_) else default__.fm5(len(d_1332_v6_), globalState))
        rhs204_ = (d_1329_v3_) + (d_1329_v3_)
        rhs205_ = d_1340_v12_
        d_1328_v2_ = rhs202_
        d_1327_v1_ = rhs203_
        d_1329_v3_ = rhs204_
        d_1340_v12_ = rhs205_
        d_1342_v13_: D10
        d_1342_v13_ = D10_DC28((self).f19, d_1327_v1_)
        d_1343_v14_: _dafny.Seq
        d_1343_v14_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1327_v1_), (d_1342_v13_).cf50, ((0) - (d_1327_v1_)) * (len(d_1329_v3_)), d_1327_v1_])
        d_1343_v14_ = ((d_1343_v14_ if (self).f19 else (((d_1343_v14_).set(default__.safeIndex(d_1327_v1_, len(d_1343_v14_)), len(d_1329_v3_))).set(default__.safeIndex(d_1327_v1_, len((d_1343_v14_).set(default__.safeIndex(d_1327_v1_, len(d_1343_v14_)), len(d_1329_v3_)))), len(default__.fm29(d_1330_v4_, 420, globalState)))) + (d_1343_v14_))).set(default__.safeIndex(d_1327_v1_, len((d_1343_v14_ if (self).f19 else (((d_1343_v14_).set(default__.safeIndex(d_1327_v1_, len(d_1343_v14_)), len(d_1329_v3_))).set(default__.safeIndex(d_1327_v1_, len((d_1343_v14_).set(default__.safeIndex(d_1327_v1_, len(d_1343_v14_)), len(d_1329_v3_)))), len(default__.fm29(d_1330_v4_, 420, globalState)))) + (d_1343_v14_)))), d_1327_v1_)
        d_1344_v15_: _dafny.MultiSet
        d_1344_v15_ = _dafny.MultiSet([d_1327_v1_])
        d_1345_v16_: D2
        d_1345_v16_ = D2_DC7(d_1344_v15_)
        pat_let_tv35_ = d_1327_v1_
        def lambda81_(source14_):
            if source14_.is_DC8:
                d_1346___mcc_h0_ = source14_.cf13
                d_1347___mcc_h1_ = source14_.cf14
                d_1348___mcc_h2_ = source14_.cf15
                d_1349_cf15_ = d_1348___mcc_h2_
                d_1350_cf14_ = d_1347___mcc_h1_
                d_1351_cf13_ = d_1346___mcc_h0_
                if (self).f19:
                    return len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxsw")))
                elif True:
                    return d_1350_cf14_
            elif True:
                d_1352___mcc_h3_ = source14_.cf12
                d_1353_cf12_ = d_1352___mcc_h3_
                return pat_let_tv35_

        d_1327_v1_ = (0) - (lambda81_(d_1345_v16_))
        d_1354_i2_: int
        d_1354_i2_ = 0
        with _dafny.label("9"):
            while False:
                with _dafny.c_label("9"):
                    if (d_1354_i2_) >= (100):
                        raise _dafny.Break("9")
                    d_1354_i2_ = (d_1354_i2_) + (1)
                    rhs206_ = ((self).f19) and ((self).f19)
                    rhs207_ = (0) - ((0) - ((default__.safeDivisionInt(d_1327_v1_, d_1327_v1_)) + (d_1327_v1_)))
                    rhs208_ = not ((self).f19) or (not(default__.fm0((self).f19, _dafny.CodePoint('l'), (self).f19, globalState)))
                    rhs209_ = (self).f19
                    rhs210_ = (self).f19
                    lhs130_ = globalState
                    lhs131_ = globalState
                    lhs132_ = globalState
                    lhs133_ = globalState
                    lhs130_.f11 = rhs206_
                    d_1327_v1_ = rhs207_
                    lhs131_.f11 = rhs208_
                    lhs132_.f17 = rhs209_
                    lhs133_.f15 = rhs210_
                    d_1355_v17_: C3
                    nw242_ = C3()
                    nw242_.ctor__()
                    d_1355_v17_ = nw242_
                    d_1326_v0_ = _dafny.MultiSet([((self).f19) == ((self).f19)])
                    d_1356_v18_: _dafny.Array
                    def lambda82_(d_1357_i3_):
                        return (self).f19

                    init41_ = lambda82_
                    nw243_ = _dafny.Array(None, 11)
                    for i0_41_ in range(nw243_.length(0)):
                        nw243_[i0_41_] = init41_(i0_41_)
                    d_1356_v18_ = nw243_
                    d_1358_v19_: _dafny.Map
                    d_1358_v19_ = _dafny.Map({(self).f19: d_1327_v1_})
                    d_1359_v20_: _dafny.Map
                    d_1359_v20_ = _dafny.Map({((d_1358_v19_)[(self).f19] if ((self).f19) in (d_1358_v19_) else len((d_1329_v3_).set(default__.safeIndex(d_1327_v1_, len(d_1329_v3_)), d_1330_v4_))): False})
                    d_1360_v21_: _dafny.Set
                    d_1360_v21_ = _dafny.Set({d_1327_v1_})
                    d_1361_v22_: D4
                    d_1361_v22_ = D4_DC14((self).f19, d_1327_v1_, d_1360_v21_, d_1330_v4_)
                    d_1362_v23_: _dafny.MultiSet
                    d_1362_v23_ = _dafny.MultiSet([(d_1361_v22_).cf26])
                    index190_ = default__.safeIndex(304, (d_1356_v18_).length(0))
                    (d_1356_v18_)[index190_] = ((d_1359_v20_)[(d_1362_v23_).cardinality] if ((d_1362_v23_).cardinality) in (d_1359_v20_) else (self).f19)
                    index191_ = default__.safeIndex(304, (d_1356_v18_).length(0))
                    (d_1356_v18_)[index191_] = (self).fm4((self).f19, _dafny.Map({d_1329_v3_: d_1358_v19_}), d_1330_v4_, d_1327_v1_, globalState)
                    pass
            pass
        (globalState).f17 = not((self).f19)
        d_1363_v24_: D9
        d_1363_v24_ = D9_DC24(d_1343_v14_)
        d_1364_v25_: _dafny.Map
        d_1364_v25_ = _dafny.Map({(self).f19: D9_DC26(d_1363_v24_)})
        d_1365_v26_: _dafny.Array
        nw244_ = _dafny.Array(None, 8)
        nw244_[int(0)] = d_1364_v25_
        nw244_[int(1)] = d_1364_v25_
        nw244_[int(2)] = d_1364_v25_
        nw244_[int(3)] = d_1364_v25_
        nw244_[int(4)] = d_1364_v25_
        nw244_[int(5)] = d_1364_v25_
        nw244_[int(6)] = d_1364_v25_
        nw244_[int(7)] = (d_1364_v25_) | (d_1364_v25_)
        d_1365_v26_ = nw244_
        index192_ = default__.safeIndex(181, (d_1365_v26_).length(0))
        (d_1365_v26_)[index192_] = d_1364_v25_
        index193_ = default__.safeIndex(181, (d_1365_v26_).length(0))
        (d_1365_v26_)[index193_] = d_1364_v25_

    def m3(self, p0, p1, globalState):
        r0: bool = False
        d_1366_v0_: _dafny.Map
        d_1366_v0_ = _dafny.Map({(self).f19: (self).f19})
        if (((d_1366_v0_)[(self).f19] if ((self).f19) in (d_1366_v0_) else ((d_1366_v0_)[p1] if (p1) in (d_1366_v0_) else (self).f19)) if (self).f19 else p1):
            d_1367_v1_: _dafny.Seq
            d_1367_v1_ = _dafny.SeqWithoutIsStrInference([p1])
            d_1368_v2_: _dafny.Seq
            d_1368_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ghvjg"))
            d_1369_v3_: _dafny.Map
            d_1369_v3_ = _dafny.Map({p1: len(d_1368_v2_)})
            d_1370_v4_: _dafny.Map
            d_1370_v4_ = _dafny.Map({not((self).f19): d_1369_v3_})
            d_1371_v5_: _dafny.Map
            d_1371_v5_ = _dafny.Map({d_1368_v2_: ((d_1370_v4_)[((d_1366_v0_)[(self).f19] if ((self).f19) in (d_1366_v0_) else (self).f19)] if (((d_1366_v0_)[(self).f19] if ((self).f19) in (d_1366_v0_) else (self).f19)) in (d_1370_v4_) else d_1369_v3_)})
            d_1372_v6_: str
            d_1372_v6_ = _dafny.CodePoint('o')
            d_1373_v7_: D1
            d_1373_v7_ = D1_DC6(p0, 943)
            d_1374_v8_: D11
            d_1374_v8_ = D11_DC30(p1, (d_1373_v7_).cf11, (self).f19)
            d_1375_v9_: _dafny.Map
            d_1375_v9_ = _dafny.Map({d_1367_v1_: p1})
            d_1376_v11_: _dafny.MultiSet
            d_1376_v11_ = _dafny.MultiSet([d_1367_v1_])
            d_1377_v14_: _dafny.MultiSet
            d_1377_v14_ = _dafny.MultiSet([True, (self).f19])
            d_1378_v15_: _dafny.Map
            d_1378_v15_ = _dafny.Map({(d_1377_v14_).cardinality: (self).f19})
            d_1379_v16_: _dafny.Map
            d_1379_v16_ = _dafny.Map({p0: d_1378_v15_})
            d_1380_v17_: _dafny.Seq
            d_1380_v17_ = _dafny.SeqWithoutIsStrInference([len(((d_1379_v16_)[p0] if (p0) in (d_1379_v16_) else d_1378_v15_)), p0, default__.fm5((d_1377_v14_).cardinality, globalState), (0) - (p0)])
            d_1381_v18_: _dafny.Seq
            d_1381_v18_ = _dafny.SeqWithoutIsStrInference([(d_1380_v17_)[default__.safeIndex(p0, len(d_1380_v17_))], p0, p0])
            d_1382_v20_: _dafny.Seq
            d_1382_v20_ = _dafny.SeqWithoutIsStrInference([d_1367_v1_])
            d_1383_v21_: _dafny.Array
            nw245_ = _dafny.Array(None, 15)
            nw245_[int(0)] = (((_dafny.Map({d_1367_v1_: p1})).set(_dafny.SeqWithoutIsStrInference([p1]), (self).fm4((self).f19, d_1371_v5_, d_1372_v6_, p0, globalState))).set(d_1367_v1_, (d_1374_v8_).cf52)).set(_dafny.SeqWithoutIsStrInference([(self).f19]), (self).f19)
            nw245_[int(1)] = d_1375_v9_
            nw245_[int(2)] = _dafny.Map({_dafny.SeqWithoutIsStrInference([True, p1]): True})
            def iife92_():
                coll38_ = _dafny.Map()
                compr_38_: _dafny.Seq
                for compr_38_ in (d_1376_v11_).Elements:
                    d_1384_v10_: _dafny.Seq = compr_38_
                    if (d_1384_v10_) in (d_1376_v11_):
                        coll38_[d_1384_v10_] = not(p1)
                return _dafny.Map(coll38_)
            nw245_[int(3)] = (iife92_()
            ) | (d_1375_v9_)
            nw245_[int(4)] = (d_1375_v9_) | (d_1375_v9_)
            nw245_[int(5)] = d_1375_v9_
            nw245_[int(6)] = _dafny.Map({default__.fm27(p0, p0, p0, globalState): p1})
            nw245_[int(7)] = d_1375_v9_
            def iife93_():
                coll39_ = _dafny.Map()
                compr_39_: _dafny.Seq
                for compr_39_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f19, not(p1)])])).Elements:
                    d_1385_v12_: _dafny.Seq = compr_39_
                    if (d_1385_v12_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f19, not(p1)])])):
                        coll39_[d_1385_v12_] = (self).f19
                return _dafny.Map(coll39_)
            nw245_[int(8)] = iife93_()
            
            nw245_[int(9)] = d_1375_v9_
            def iife94_():
                coll40_ = _dafny.Map()
                compr_40_: _dafny.Seq
                for compr_40_ in (_dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): (self).f19})).keys.Elements:
                    d_1386_v13_: _dafny.Seq = compr_40_
                    if (d_1386_v13_) in (_dafny.Map({_dafny.SeqWithoutIsStrInference([p1]): (self).f19})):
                        coll40_[d_1386_v13_] = p1
                return _dafny.Map(coll40_)
            nw245_[int(10)] = iife94_()
            
            nw245_[int(11)] = default__.fm30(d_1381_v18_, globalState)
            nw245_[int(12)] = (_dafny.Map({d_1367_v1_: default__.fm0((self).f19, d_1372_v6_, p1, globalState)})) | (d_1375_v9_)
            def iife95_():
                coll41_ = _dafny.Map()
                compr_41_: _dafny.Seq
                for compr_41_ in (d_1382_v20_).Elements:
                    d_1387_v19_: _dafny.Seq = compr_41_
                    if (d_1387_v19_) in (d_1382_v20_):
                        coll41_[d_1387_v19_] = (self).f19
                return _dafny.Map(coll41_)
            nw245_[int(13)] = iife95_()
            
            nw245_[int(14)] = (d_1375_v9_) | (d_1375_v9_)
            d_1383_v21_ = nw245_
            index194_ = default__.safeIndex(336, (d_1383_v21_).length(0))
            (d_1383_v21_)[index194_] = d_1375_v9_
            index195_ = default__.safeIndex(336, (d_1383_v21_).length(0))
            (d_1383_v21_)[index195_] = d_1375_v9_
            if (d_1367_v1_) < (d_1367_v1_):
                d_1388_v22_: _dafny.Array
                nw246_ = _dafny.Array(int(0), 13)
                d_1388_v22_ = nw246_
                d_1389_v23_: _dafny.Map
                d_1389_v23_ = _dafny.Map({d_1388_v22_: (d_1380_v17_)[default__.safeIndex(-865, len(d_1380_v17_))]})
                d_1389_v23_ = (d_1389_v23_).set(d_1388_v22_, p0)
                d_1390_v24_: int
                d_1390_v24_ = 47
                d_1391_v25_: _dafny.MultiSet
                d_1391_v25_ = _dafny.MultiSet([D10_DC27(d_1377_v14_)])
                d_1390_v24_ = ((d_1391_v25_)[default__.fm31(True, (d_1381_v18_)[default__.safeIndex((0) - (d_1390_v24_), len(d_1381_v18_))], globalState)] if (default__.fm31(True, (d_1381_v18_)[default__.safeIndex((0) - (d_1390_v24_), len(d_1381_v18_))], globalState)) in (d_1391_v25_) else d_1390_v24_)
                d_1392_v26_: _dafny.Array
                nw247_ = _dafny.Array(False, 5)
                d_1392_v26_ = nw247_
                d_1393_v27_: _dafny.MultiSet
                d_1393_v27_ = _dafny.MultiSet([d_1392_v26_])
                d_1394_v28_: _dafny.Map
                d_1394_v28_ = _dafny.Map({d_1393_v27_: (0) - ((default__.fm5(p0, globalState)) * (d_1390_v24_))})
                d_1394_v28_ = (d_1394_v28_).set(d_1393_v27_, 592)
                d_1395_v29_: D9
                d_1395_v29_ = D9_DC26(D9_DC24(d_1381_v18_))
                d_1396_v30_: C0
                nw248_ = C0()
                nw248_.ctor__(p1)
                d_1396_v30_ = nw248_
                d_1397_v31_: D9
                d_1397_v31_ = D9_DC26(D9_DC26(D9_DC25(d_1390_v24_, (self).f19, d_1390_v24_, d_1396_v30_)))
                pat_let_tv36_ = d_1397_v31_
                d_1398_v32_: _dafny.Array
                nw249_ = _dafny.Array(None, 9)
                nw249_[int(0)] = d_1395_v29_
                nw249_[int(1)] = D9_DC26(D9_DC26(d_1397_v31_))
                nw249_[int(2)] = d_1395_v29_
                nw249_[int(3)] = d_1395_v29_
                nw249_[int(4)] = D9_DC26(D9_DC24(d_1380_v17_))
                nw249_[int(5)] = (d_1395_v29_ if not(p1) else d_1395_v29_)
                nw249_[int(6)] = d_1395_v29_
                nw249_[int(7)] = d_1395_v29_
                def iife96_(_pat_let27_0):
                    def iife97_(d_1399_dt__update__tmp_h0_):
                        def iife98_(_pat_let28_0):
                            def iife99_(d_1400_dt__update_hcf47_h0_):
                                return D9_DC26(d_1400_dt__update_hcf47_h0_)
                            return iife99_(_pat_let28_0)
                        return iife98_(pat_let_tv36_)
                    return iife97_(_pat_let27_0)
                nw249_[int(8)] = iife96_(d_1395_v29_)
                d_1398_v32_ = nw249_
                index196_ = default__.safeIndex(259, (d_1398_v32_).length(0))
                (d_1398_v32_)[index196_] = d_1395_v29_
                d_1401_v33_: D4
                d_1401_v33_ = D4_DC14((self).f19, p0, default__.fm18((d_1396_v30_).f23, len(d_1380_v17_), -821, globalState), d_1372_v6_)
                d_1402_v34_: _dafny.Set
                d_1402_v34_ = _dafny.Set({(d_1380_v17_)[default__.safeIndex(d_1390_v24_, len(d_1380_v17_))]})
                d_1403_v35_: _dafny.Set
                d_1403_v35_ = _dafny.Set({(d_1401_v33_).cf25, _dafny.Set({len(d_1368_v2_)}), d_1402_v34_, d_1402_v34_, _dafny.Set({(d_1374_v8_).cf53, p0})})
                index197_ = default__.safeIndex(259, (d_1398_v32_).length(0))
                rhs211_ = p0
                rhs212_ = D9_DC26(d_1397_v31_)
                rhs213_ = not((d_1403_v35_) == (d_1403_v35_))
                lhs134_ = d_1398_v32_
                lhs135_ = default__.safeIndex(259, (d_1398_v32_).length(0))
                lhs136_ = globalState
                d_1390_v24_ = rhs211_
                lhs134_[lhs135_] = rhs212_
                lhs136_.f15 = rhs213_
                index198_ = default__.safeIndex(933, (d_1388_v22_).length(0))
                (d_1388_v22_)[index198_] = default__.safeDivisionInt(d_1390_v24_, len(d_1367_v1_))
                d_1404_v36_: _dafny.MultiSet
                d_1404_v36_ = _dafny.MultiSet([_dafny.CodePoint('u')])
                d_1405_v37_: D4
                d_1405_v37_ = D4_DC13((0) - (len(d_1368_v2_)))
                d_1406_v38_: _dafny.Map
                d_1406_v38_ = _dafny.Map({d_1405_v37_: (0) - (d_1390_v24_)})
                index199_ = default__.safeIndex(933, (d_1388_v22_).length(0))
                rhs214_ = ((d_1396_v30_).f23) or ((d_1404_v36_).ispropersubset(d_1404_v36_))
                rhs215_ = (d_1396_v30_).f23
                rhs216_ = ((d_1406_v38_)[d_1405_v37_] if (d_1405_v37_) in (d_1406_v38_) else d_1390_v24_)
                lhs137_ = globalState
                lhs138_ = globalState
                lhs139_ = d_1388_v22_
                lhs140_ = default__.safeIndex(933, (d_1388_v22_).length(0))
                lhs137_.f17 = rhs214_
                lhs138_.f17 = rhs215_
                lhs139_[lhs140_] = rhs216_
            elif True:
                d_1407_v39_: _dafny.Map
                d_1407_v39_ = _dafny.Map({d_1374_v8_: not(p1)})
                d_1408_v40_: _dafny.Map
                d_1408_v40_ = _dafny.Map({d_1367_v1_: (d_1407_v39_).set(d_1374_v8_, p1)})
                d_1409_v41_: _dafny.Array
                nw250_ = _dafny.Array(None, 20)
                nw250_[int(0)] = not((self).f19)
                nw250_[int(1)] = (self).f19
                nw250_[int(2)] = (self).f19
                nw250_[int(3)] = not ((self).f19) or (not(not((self).f19)))
                nw250_[int(4)] = (self).f19
                nw250_[int(5)] = not((d_1367_v1_) not in (d_1408_v40_))
                nw250_[int(6)] = p1
                nw250_[int(7)] = (self).f19
                nw250_[int(8)] = p1
                nw250_[int(9)] = p1
                nw250_[int(10)] = ((d_1366_v0_)[p1] if (p1) in (d_1366_v0_) else True)
                nw250_[int(11)] = True
                nw250_[int(12)] = p1
                nw250_[int(13)] = ((self).f19 if (self).fm4((self).fm4((self).f19, d_1371_v5_, d_1372_v6_, (0) - (p0), globalState), _dafny.Map({d_1368_v2_: d_1369_v3_}), d_1372_v6_, len(_dafny.SeqWithoutIsStrInference([d_1372_v6_])), globalState) else True)
                nw250_[int(14)] = (self).f19
                nw250_[int(15)] = (self).f19
                nw250_[int(16)] = (self).f19
                nw250_[int(17)] = ((d_1378_v15_)[p0] if (p0) in (d_1378_v15_) else False)
                nw250_[int(18)] = p1
                nw250_[int(19)] = p1
                d_1409_v41_ = nw250_
                d_1410_v42_: D0
                d_1410_v42_ = D0_DC1((self).f19)
                d_1411_v43_: T1
                nw251_ = C2()
                nw251_.ctor__(d_1410_v42_, p1)
                d_1411_v43_ = nw251_
                d_1412_v44_: D12
                d_1412_v44_ = D12_DC31(_dafny.Set({d_1411_v43_, d_1411_v43_}))
                index200_ = default__.safeIndex(987, (d_1409_v41_).length(0))
                (d_1409_v41_)[index200_] = (d_1411_v43_) not in ((d_1412_v44_).cf55)
                d_1413_v45_: int
                d_1413_v45_ = 579
                d_1414_v46_: _dafny.Array
                def lambda83_(d_1415_v45_):
                    def lambda84_(d_1416_i0_):
                        return default__.safeModuloInt(d_1416_i0_, d_1415_v45_)

                    return lambda84_

                init42_ = lambda83_(d_1413_v45_)
                nw252_ = _dafny.Array(None, 8)
                for i0_42_ in range(nw252_.length(0)):
                    nw252_[i0_42_] = init42_(i0_42_)
                d_1414_v46_ = nw252_
                d_1417_v47_: D3
                d_1417_v47_ = D3_DC11(_dafny.Map({(self).f19: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))}), (0) - (d_1413_v45_), d_1414_v46_)
                index201_ = default__.safeIndex(987, (d_1409_v41_).length(0))
                rhs217_ = d_1380_v17_
                rhs218_ = (d_1367_v1_)[default__.safeIndex((d_1417_v47_).cf19, len(d_1367_v1_))]
                rhs219_ = p0
                lhs141_ = d_1409_v41_
                lhs142_ = default__.safeIndex(987, (d_1409_v41_).length(0))
                d_1380_v17_ = rhs217_
                lhs141_[lhs142_] = rhs218_
                d_1413_v45_ = rhs219_
                d_1418_v48_: _dafny.Set
                d_1418_v48_ = _dafny.Set({d_1413_v45_})
                d_1413_v45_ = len((d_1418_v48_) - (d_1418_v48_))
                d_1413_v45_ = (d_1413_v45_) * (len((d_1368_v2_ if not(p1) else _dafny.SeqWithoutIsStrInference([d_1372_v6_ for d_1419_i1_ in range(default__.abs(271))]))))
                (globalState).f17 = not((d_1409_v41_)[default__.safeIndex(987, (d_1409_v41_).length(0))])
                d_1420_v49_: C2
                nw253_ = C2()
                nw253_.ctor__(d_1410_v42_, (d_1409_v41_)[default__.safeIndex(987, (d_1409_v41_).length(0))])
                d_1420_v49_ = nw253_
            if ((d_1368_v2_) + (default__.fm16((self).f19, p0, globalState))) <= (d_1368_v2_):
                d_1421_v50_: _dafny.Array
                def lambda85_(d_1422_i2_):
                    return (d_1422_i2_) - (-313)

                init43_ = lambda85_
                nw254_ = _dafny.Array(None, 7)
                for i0_43_ in range(nw254_.length(0)):
                    nw254_[i0_43_] = init43_(i0_43_)
                d_1421_v50_ = nw254_
                index202_ = default__.safeIndex(241, (d_1421_v50_).length(0))
                (d_1421_v50_)[index202_] = len(_dafny.SeqWithoutIsStrInference([d_1372_v6_ for d_1423_i3_ in range(default__.abs(626))]))
                index203_ = default__.safeIndex(1, (d_1421_v50_).length(0))
                (d_1421_v50_)[index203_] = len(d_1368_v2_)
                d_1424_v51_: int
                d_1424_v51_ = -819
                d_1425_v52_: _dafny.Set
                d_1425_v52_ = _dafny.Set({p0, p0})
                d_1426_v53_: _dafny.Map
                d_1426_v53_ = _dafny.Map({d_1372_v6_: (self).f19})
                index204_ = default__.safeIndex(241, (d_1421_v50_).length(0))
                index205_ = default__.safeIndex(1, (d_1421_v50_).length(0))
                rhs220_ = len(d_1366_v0_)
                rhs221_ = (d_1368_v2_) + (default__.fm16(default__.fm0((self).f19, d_1372_v6_, (self).f19, globalState), p0, globalState))
                rhs222_ = p0
                rhs223_ = default__.safeDivisionInt((len(d_1425_v52_)) + (d_1424_v51_), len((d_1426_v53_) | (_dafny.Map({d_1372_v6_: not(not(not(p1)))}))))
                lhs143_ = d_1421_v50_
                lhs144_ = default__.safeIndex(241, (d_1421_v50_).length(0))
                lhs145_ = d_1421_v50_
                lhs146_ = default__.safeIndex(1, (d_1421_v50_).length(0))
                lhs143_[lhs144_] = rhs220_
                d_1368_v2_ = rhs221_
                lhs145_[lhs146_] = rhs222_
                d_1424_v51_ = rhs223_
                index206_ = default__.safeIndex(241, (d_1421_v50_).length(0))
                (d_1421_v50_)[index206_] = (d_1424_v51_ if (d_1424_v51_) <= ((d_1421_v50_)[default__.safeIndex(241, (d_1421_v50_).length(0))]) else default__.fm5(p0, globalState))
                nw255_ = _dafny.Array(int(0), 9)
                d_1421_v50_ = nw255_
                d_1427_v54_: D1
                d_1427_v54_ = D1_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iteqm")), p0, (d_1421_v50_)[default__.safeIndex(241, (d_1421_v50_).length(0))])
                d_1428_v55_: _dafny.Set
                d_1428_v55_ = _dafny.Set({d_1368_v2_, (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fofwaytk"))) + (d_1368_v2_), (d_1427_v54_).cf7})
                d_1428_v55_ = _dafny.Set({d_1368_v2_, d_1368_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frcpgrc"))})
                d_1429_v56_: D2
                d_1429_v56_ = D2_DC7(_dafny.MultiSet(d_1380_v17_))
                d_1430_v57_: _dafny.Map
                d_1430_v57_ = _dafny.Map({not(True): d_1372_v6_})
                d_1431_v58_: D4
                d_1431_v58_ = D4_DC14(p1, d_1424_v51_, d_1425_v52_, d_1372_v6_)
                d_1429_v56_ = default__.fm32(p1, (0) - (default__.safeDivisionInt(len(d_1430_v57_), 905)), (d_1421_v50_)[default__.safeIndex(241, (d_1421_v50_).length(0))], default__.fm0(not((self).f19), (d_1431_v58_).cf26, p1, globalState), globalState)
            elif True:
                d_1432_v59_: int
                d_1432_v59_ = -135
                d_1432_v59_ = p0
                d_1433_v60_: _dafny.Set
                d_1433_v60_ = _dafny.Set({d_1432_v59_})
                d_1434_v61_: D4
                d_1434_v61_ = D4_DC14(True, d_1432_v59_, d_1433_v60_, d_1372_v6_)
                d_1435_v62_: _dafny.Set
                d_1435_v62_ = _dafny.Set({default__.fm0(not(p1), (d_1434_v61_).cf26, p1, globalState), (self).f19, p1})
                d_1436_v63_: _dafny.Array
                nw256_ = _dafny.Array(None, 2)
                nw256_[int(0)] = p0
                nw256_[int(1)] = len(d_1435_v62_)
                d_1436_v63_ = nw256_
                d_1437_v64_: _dafny.Seq
                d_1437_v64_ = _dafny.SeqWithoutIsStrInference([(d_1436_v63_ if (self).f19 else d_1436_v63_), d_1436_v63_, d_1436_v63_])
                d_1438_v65_: D3
                d_1438_v65_ = D3_DC10(len(d_1378_v15_))
                d_1439_v66_: _dafny.MultiSet
                d_1439_v66_ = _dafny.MultiSet([d_1438_v65_, d_1438_v65_])
                rhs224_ = _dafny.SeqWithoutIsStrInference([d_1436_v63_, d_1436_v63_, d_1436_v63_, d_1436_v63_, d_1436_v63_])
                rhs225_ = len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "je"))) + (d_1368_v2_) if (self).f19 else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_1440_i4_ in range(default__.abs(816))])))
                rhs226_ = ((d_1367_v1_)[default__.safeIndex((d_1380_v17_)[default__.safeIndex(((d_1439_v66_)[d_1438_v65_] if (d_1438_v65_) in (d_1439_v66_) else d_1432_v59_), len(d_1380_v17_))], len(d_1367_v1_))] if p1 else p1)
                lhs147_ = globalState
                d_1437_v64_ = rhs224_
                d_1432_v59_ = rhs225_
                lhs147_.f17 = rhs226_
                d_1432_v59_ = (len((d_1366_v0_) | (_dafny.Map({(self).f19: True}))) if (d_1368_v2_) != (_dafny.SeqWithoutIsStrInference([d_1372_v6_ for d_1441_i5_ in range(default__.abs(26))])) else -108)
                d_1442_v67_: _dafny.Map
                d_1442_v67_ = _dafny.Map({(_dafny.MultiSet([not(p1), (self).f19])).issubset(default__.fm28(len(d_1368_v2_), True, globalState)): default__.fm33((self).f19, True, globalState)})
                d_1443_v68_: D10
                d_1443_v68_ = D10_DC28(p1, d_1432_v59_)
                d_1442_v67_ = _dafny.Map({p1: d_1443_v68_})
                d_1372_v6_ = d_1372_v6_
            (globalState).f15 = (self).f19
            d_1444_v69_: D0
            d_1444_v69_ = D0_DC1((self).f19)
            d_1445_v70_: T1
            nw257_ = C2()
            nw257_.ctor__(d_1444_v69_, p1)
            d_1445_v70_ = nw257_
            d_1446_v71_: D12
            d_1446_v71_ = D12_DC31((_dafny.Set({d_1445_v70_, d_1445_v70_})) - (_dafny.Set({d_1445_v70_})))
            d_1446_v71_ = d_1446_v71_
        elif True:
            d_1447_v72_: str
            d_1447_v72_ = _dafny.CodePoint('r')
            d_1447_v72_ = (d_1447_v72_ if (self).f19 else d_1447_v72_)
            d_1448_v73_: _dafny.Map
            d_1448_v73_ = _dafny.Map({p1: p0})
            d_1449_v74_: _dafny.Array
            nw258_ = _dafny.Array(False, 14)
            d_1449_v74_ = nw258_
            d_1450_v75_: D1
            d_1450_v75_ = D1_DC4(p0, (self).f19, d_1449_v74_)
            d_1451_v76_: _dafny.Seq
            d_1451_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ird"))
            d_1452_v77_: _dafny.Map
            d_1452_v77_ = _dafny.Map({(d_1450_v75_).cf5: d_1451_v76_})
            d_1448_v73_ = (d_1448_v73_).set(False, len(d_1452_v77_))
            d_1366_v0_ = (d_1366_v0_).set(p1, (not(default__.fm0((self).f19, d_1447_v72_, p1, globalState))) or (False))
            r0 = ((self).f19) or ((self).f19)
            d_1453_v78_: D0
            d_1453_v78_ = D0_DC1(True)
            d_1454_v79_: C2
            nw259_ = C2()
            nw259_.ctor__(d_1453_v78_, (self).f19)
            d_1454_v79_ = nw259_
            d_1455_v80_: _dafny.Map
            d_1455_v80_ = _dafny.Map({-116: d_1454_v79_})
            d_1456_v81_: _dafny.Set
            d_1456_v81_ = _dafny.Set({d_1455_v80_})
            d_1457_v82_: _dafny.MultiSet
            d_1457_v82_ = _dafny.MultiSet([p0])
            d_1458_v83_: D8
            d_1458_v83_ = D8_DC23(d_1457_v82_)
            d_1459_v84_: _dafny.Seq
            d_1459_v84_ = _dafny.SeqWithoutIsStrInference([(d_1454_v79_).f21, p1])
            d_1460_v85_: _dafny.Map
            d_1460_v85_ = _dafny.Map({p0: (d_1454_v79_).f21})
            d_1461_v86_: C3
            nw260_ = C3()
            nw260_.ctor__()
            d_1461_v86_ = nw260_
            d_1462_v87_: _dafny.Seq
            d_1462_v87_ = _dafny.SeqWithoutIsStrInference([d_1458_v83_, D8_DC23((default__.fm23((d_1459_v84_)[default__.safeIndex((0) - (p0), len(d_1459_v84_))], ((_dafny.Map({(d_1454_v79_).f21: len(d_1451_v76_)})).set((d_1454_v79_).f21, len(d_1460_v85_))).set(not((d_1454_v79_).f21), p0), (0) - (len(_dafny.Map({p0: d_1461_v86_}))), d_1447_v72_, globalState)).set(-748, default__.abs(p0))), D8_DC23(d_1457_v82_), d_1458_v83_])
            d_1463_v88_: _dafny.Map
            d_1463_v88_ = _dafny.Map({(d_1456_v81_).issubset(d_1456_v81_): d_1462_v87_})
            d_1463_v88_ = (d_1463_v88_) | (d_1463_v88_)
        d_1464_v89_: str
        d_1464_v89_ = _dafny.CodePoint('y')
        d_1465_v90_: int
        d_1465_v90_ = 975
        d_1466_v91_: _dafny.Array
        nw261_ = _dafny.Array(False, 28)
        d_1466_v91_ = nw261_
        d_1467_v92_: D12
        d_1467_v92_ = D12_DC32(d_1465_v90_)
        d_1468_v93_: _dafny.Seq
        d_1468_v93_ = _dafny.SeqWithoutIsStrInference([(self).f19, default__.fm0((self).f19, d_1464_v89_, (self).f19, globalState), (self).f19, False, True])
        d_1469_v94_: _dafny.Map
        d_1469_v94_ = _dafny.Map({p1: d_1465_v90_})
        d_1470_v95_: _dafny.Seq
        d_1470_v95_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([p1]))])
        d_1471_v96_: _dafny.Seq
        d_1471_v96_ = _dafny.SeqWithoutIsStrInference([(p0) * (d_1465_v90_), p0, (d_1467_v92_).cf56, (p0 if (d_1468_v93_)[default__.safeIndex(((d_1469_v94_)[(self).f19] if ((self).f19) in (d_1469_v94_) else 589), len(d_1468_v93_))] else d_1465_v90_), len((_dafny.SeqWithoutIsStrInference([d_1465_v90_ for d_1472_i6_ in range(default__.abs(389))])) + (d_1470_v95_))])
        d_1473_v97_: _dafny.Array
        nw262_ = _dafny.Array(None, 2)
        nw262_[int(0)] = d_1468_v93_
        nw262_[int(1)] = d_1468_v93_
        d_1473_v97_ = nw262_
        d_1474_v98_: D1
        d_1474_v98_ = D1_DC3(d_1473_v97_)
        d_1475_v99_: _dafny.MultiSet
        d_1475_v99_ = _dafny.MultiSet([d_1474_v98_])
        d_1476_v100_: _dafny.Map
        d_1476_v100_ = _dafny.Map({p0: d_1475_v99_})
        d_1477_v101_: _dafny.Map
        d_1477_v101_ = _dafny.Map({d_1476_v100_: default__.safeDivisionInt(p0, d_1465_v90_)})
        d_1478_v102_: _dafny.Map
        d_1478_v102_ = _dafny.Map({((0) - (p0)) != (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({p1, p1})), d_1465_v90_]))): d_1466_v91_})
        d_1479_v103_: _dafny.Map
        d_1479_v103_ = _dafny.Map({len(d_1469_v94_): d_1470_v95_})
        rhs227_ = d_1464_v89_
        rhs228_ = ((d_1477_v101_)[d_1476_v100_] if (d_1476_v100_) in (d_1477_v101_) else p0)
        rhs229_ = ((d_1478_v102_)[(self).f19] if ((self).f19) in (d_1478_v102_) else d_1466_v91_)
        rhs230_ = p1
        rhs231_ = (((d_1479_v103_)[d_1465_v90_] if (d_1465_v90_) in (d_1479_v103_) else d_1471_v96_)) + (_dafny.SeqWithoutIsStrInference([p0 for d_1480_i7_ in range(default__.abs(906))]))
        d_1464_v89_ = rhs227_
        d_1465_v90_ = rhs228_
        d_1466_v91_ = rhs229_
        r0 = rhs230_
        d_1471_v96_ = rhs231_
        d_1481_v104_: _dafny.Seq
        d_1481_v104_ = _dafny.SeqWithoutIsStrInference([d_1366_v0_, (d_1366_v0_) | (d_1366_v0_)])
        d_1482_v105_: D13
        d_1482_v105_ = D13_DC36(_dafny.SeqWithoutIsStrInference([d_1366_v0_]))
        d_1481_v104_ = (((d_1482_v105_).cf63 if (p0) > (p0) else _dafny.SeqWithoutIsStrInference([d_1366_v0_ for d_1483_i8_ in range(default__.abs(31))]))).set(default__.safeIndex((d_1465_v90_) + ((0) - (d_1465_v90_)), len(((d_1482_v105_).cf63 if (p0) > (p0) else _dafny.SeqWithoutIsStrInference([d_1366_v0_ for d_1484_i8_ in range(default__.abs(31))])))), d_1366_v0_)
        d_1485_v106_: D0
        d_1485_v106_ = D0_DC1((self).f19)
        d_1486_v107_: C2
        nw263_ = C2()
        nw263_.ctor__((d_1485_v106_ if not((self).f19) else d_1485_v106_), (True if not(p1) else (self).f19))
        d_1486_v107_ = nw263_
        d_1487_v108_: _dafny.Array
        nw264_ = _dafny.Array(int(0), 16)
        d_1487_v108_ = nw264_
        d_1488_v109_: D13
        d_1488_v109_ = D13_DC38(d_1464_v89_, (d_1486_v107_).f21, d_1465_v90_, d_1487_v108_)
        d_1489_v110_: _dafny.Array
        nw265_ = _dafny.Array(None, 5)
        nw265_[int(0)] = d_1488_v109_
        nw265_[int(1)] = d_1488_v109_
        nw265_[int(2)] = d_1488_v109_
        nw265_[int(3)] = d_1488_v109_
        nw265_[int(4)] = (D13_DC38(d_1464_v89_, p1, d_1465_v90_, d_1487_v108_) if (d_1486_v107_).f21 else d_1488_v109_)
        d_1489_v110_ = nw265_
        d_1490_v111_: D14
        d_1490_v111_ = D14_DC40(d_1489_v110_)
        d_1489_v110_ = (d_1490_v111_).cf71
        d_1491_i9_: int
        d_1491_i9_ = 0
        with _dafny.label("10"):
            while (default__.safeModuloInt(p0, d_1465_v90_)) not in (_dafny.Map({d_1465_v90_: d_1465_v90_})):
                with _dafny.c_label("10"):
                    if (d_1491_i9_) >= (100):
                        raise _dafny.Break("10")
                    d_1491_i9_ = (d_1491_i9_) + (1)
                    d_1492_v112_: _dafny.Map
                    d_1492_v112_ = _dafny.Map({d_1464_v89_: d_1464_v89_})
                    d_1493_v113_: _dafny.Seq
                    d_1493_v113_ = _dafny.SeqWithoutIsStrInference([d_1464_v89_, _dafny.CodePoint('k'), d_1464_v89_, d_1464_v89_, _dafny.CodePoint('e')])
                    d_1492_v112_ = (d_1492_v112_).set((d_1493_v113_)[default__.safeIndex(14, len(d_1493_v113_))], d_1464_v89_)
                    rhs232_ = not (False) or ((d_1486_v107_).f21)
                    rhs233_ = d_1486_v107_
                    lhs148_ = globalState
                    lhs148_.f17 = rhs232_
                    d_1486_v107_ = rhs233_
                    index207_ = default__.safeIndex(762, (d_1487_v108_).length(0))
                    (d_1487_v108_)[index207_] = p0
                    index208_ = default__.safeIndex(26, (d_1466_v91_).length(0))
                    (d_1466_v91_)[index208_] = p1
                    d_1494_v114_: _dafny.Set
                    d_1494_v114_ = _dafny.Set({d_1464_v89_, d_1464_v89_, d_1464_v89_})
                    index209_ = default__.safeIndex(762, (d_1487_v108_).length(0))
                    index210_ = default__.safeIndex(26, (d_1466_v91_).length(0))
                    rhs234_ = (p0 if not((d_1494_v114_).issubset(d_1494_v114_)) else len(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxxgp"))) + (d_1493_v113_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxxgp"))) + (d_1493_v113_))), d_1464_v89_)))
                    rhs235_ = (self).f19
                    lhs149_ = d_1487_v108_
                    lhs150_ = default__.safeIndex(762, (d_1487_v108_).length(0))
                    lhs151_ = d_1466_v91_
                    lhs152_ = default__.safeIndex(26, (d_1466_v91_).length(0))
                    lhs149_[lhs150_] = rhs234_
                    lhs151_[lhs152_] = rhs235_
                    d_1465_v90_ = len((d_1493_v113_) + (d_1493_v113_))
                    pass
            pass
        r0 = (d_1486_v107_.f20).cf1
        return r0

    @property
    def f19(self):
        return self._f19

class C5:
    def  __init__(self):
        self._f18: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    def ctor__(self, f18):
        (self)._f18 = f18

    def m0(self, p0, p1, globalState):
        r0: int = int(0)
        d_1495_v0_: bool
        d_1495_v0_ = True
        d_1496_v1_: _dafny.Seq
        d_1496_v1_ = _dafny.SeqWithoutIsStrInference([d_1495_v0_])
        if (d_1496_v1_) < (d_1496_v1_):
            d_1497_v2_: _dafny.Array
            def lambda86_(d_1498_v1_):
                def lambda87_(d_1499_i0_):
                    return d_1498_v1_

                return lambda87_

            init44_ = lambda86_(d_1496_v1_)
            nw266_ = _dafny.Array(None, 26)
            for i0_44_ in range(nw266_.length(0)):
                nw266_[i0_44_] = init44_(i0_44_)
            d_1497_v2_ = nw266_
            d_1500_v3_: D1
            d_1500_v3_ = D1_DC3(d_1497_v2_)
            d_1497_v2_ = (d_1500_v3_).cf3
            d_1501_v4_: _dafny.Array
            def lambda88_(d_1502_i1_):
                return D1_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "odrkr")), (self).f18, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tbdrju"))))

            init45_ = lambda88_
            nw267_ = _dafny.Array(None, 25)
            for i0_45_ in range(nw267_.length(0)):
                nw267_[i0_45_] = init45_(i0_45_)
            d_1501_v4_ = nw267_
            d_1503_v5_: _dafny.Seq
            d_1503_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bnr"))
            index211_ = default__.safeIndex(378, (d_1501_v4_).length(0))
            (d_1501_v4_)[index211_] = D1_DC5(d_1503_v5_, (self).f18, (self).f18)
            d_1504_v6_: _dafny.Seq
            d_1504_v6_ = _dafny.SeqWithoutIsStrInference([797])
            index212_ = default__.safeIndex(378, (d_1501_v4_).length(0))
            rhs236_ = default__.fm1((self).f18, globalState)
            rhs237_ = (_dafny.MultiSet([(self).f18])).ispropersubset(_dafny.MultiSet(d_1504_v6_))
            rhs238_ = (p0).isdisjoint(p0)
            rhs239_ = (d_1496_v1_)[default__.safeIndex(((self).f18) * (len(d_1504_v6_)), len(d_1496_v1_))]
            lhs153_ = d_1501_v4_
            lhs154_ = default__.safeIndex(378, (d_1501_v4_).length(0))
            lhs155_ = globalState
            lhs156_ = globalState
            lhs153_[lhs154_] = rhs236_
            d_1495_v0_ = rhs237_
            lhs155_.f15 = rhs238_
            lhs156_.f11 = rhs239_
            d_1505_v7_: _dafny.Map
            d_1505_v7_ = _dafny.Map({(self).f18: (self).f18})
            d_1505_v7_ = (d_1505_v7_).set((self).f18, (166) - ((self).f18))
            d_1506_v8_: _dafny.Array
            nw268_ = _dafny.Array(False, 11)
            d_1506_v8_ = nw268_
            d_1507_v9_: D1
            d_1507_v9_ = D1_DC6((self).f18, default__.safeDivisionInt((self).f18, (self).f18))
            d_1508_v10_: D0
            d_1508_v10_ = D0_DC1(d_1495_v0_)
            d_1509_v11_: D0
            d_1509_v11_ = D0_DC2(d_1508_v10_)
            rhs240_ = d_1506_v8_
            rhs241_ = (((self).f18) * (len(d_1496_v1_)) if False else (self).f18)
            rhs242_ = (self).f18
            rhs243_ = default__.fm2(globalState)
            rhs244_ = d_1509_v11_
            d_1506_v8_ = rhs240_
            r0 = rhs241_
            r0 = rhs242_
            d_1507_v9_ = rhs243_
            d_1509_v11_ = rhs244_
            d_1510_v12_: _dafny.Seq
            d_1510_v12_ = _dafny.SeqWithoutIsStrInference([d_1503_v5_, d_1503_v5_, d_1503_v5_, d_1503_v5_, d_1503_v5_])
            if not((d_1510_v12_) != (d_1510_v12_)):
                d_1511_v13_: _dafny.Seq
                d_1511_v13_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1495_v0_])])
                d_1512_v14_: _dafny.MultiSet
                d_1512_v14_ = _dafny.MultiSet([d_1495_v0_])
                d_1513_v15_: _dafny.Map
                d_1513_v15_ = _dafny.Map({d_1495_v0_: (d_1512_v14_).cardinality})
                d_1514_v16_: _dafny.MultiSet
                d_1514_v16_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([(d_1503_v5_)[default__.safeIndex(855, len(d_1503_v5_))] for d_1515_i2_ in range(default__.abs(228))]))])
                d_1516_v17_: _dafny.Array
                nw269_ = _dafny.Array(None, 23)
                nw269_[int(0)] = (self).f18
                nw269_[int(1)] = (self).f18
                nw269_[int(2)] = default__.safeDivisionInt(len((d_1511_v13_)[default__.safeIndex((self).f18, len(d_1511_v13_))]), (d_1512_v14_).cardinality)
                nw269_[int(3)] = len(p0)
                nw269_[int(4)] = ((self).f18) + ((self).f18)
                nw269_[int(5)] = (self).f18
                nw269_[int(6)] = (self).f18
                nw269_[int(7)] = default__.safeDivisionInt(len(d_1504_v6_), -560)
                nw269_[int(8)] = 871
                nw269_[int(9)] = (0) - ((self).f18)
                nw269_[int(10)] = 231
                nw269_[int(11)] = ((0) - ((self).f18)) * ((self).f18)
                nw269_[int(12)] = ((d_1513_v15_)[d_1495_v0_] if (d_1495_v0_) in (d_1513_v15_) else (self).f18)
                nw269_[int(13)] = (0) - ((self).f18)
                nw269_[int(14)] = (self).f18
                nw269_[int(15)] = (((d_1514_v16_).set(len(d_1505_v7_), default__.abs(655))).intersection(d_1514_v16_)).cardinality
                nw269_[int(16)] = (self).f18
                nw269_[int(17)] = (self).f18
                nw269_[int(18)] = default__.fm3((d_1504_v6_)[default__.safeIndex((self).f18, len(d_1504_v6_))], (self).f18, d_1495_v0_, globalState)
                nw269_[int(19)] = (default__.fm3((self).f18, (self).f18, d_1495_v0_, globalState)) * ((self).f18)
                nw269_[int(20)] = (self).f18
                nw269_[int(21)] = (self).f18
                nw269_[int(22)] = len(d_1503_v5_)
                d_1516_v17_ = nw269_
                d_1517_v18_: str
                d_1517_v18_ = _dafny.CodePoint('j')
                d_1518_v19_: _dafny.MultiSet
                d_1518_v19_ = _dafny.MultiSet([d_1503_v5_, _dafny.SeqWithoutIsStrInference([d_1517_v18_])])
                index213_ = default__.safeIndex(560, (d_1516_v17_).length(0))
                (d_1516_v17_)[index213_] = (((d_1518_v19_)[d_1503_v5_] if (d_1503_v5_) in (d_1518_v19_) else (self).f18)) + ((self).f18)
                index214_ = default__.safeIndex(115, (d_1506_v8_).length(0))
                (d_1506_v8_)[index214_] = d_1495_v0_
                d_1519_v20_: D0
                d_1519_v20_ = D0_DC1(d_1495_v0_)
                index215_ = default__.safeIndex(560, (d_1516_v17_).length(0))
                index216_ = default__.safeIndex(115, (d_1506_v8_).length(0))
                rhs245_ = (0) - ((self).f18)
                rhs246_ = d_1495_v0_
                rhs247_ = d_1519_v20_
                lhs157_ = d_1516_v17_
                lhs158_ = default__.safeIndex(560, (d_1516_v17_).length(0))
                lhs159_ = d_1506_v8_
                lhs160_ = default__.safeIndex(115, (d_1506_v8_).length(0))
                lhs157_[lhs158_] = rhs245_
                lhs159_[lhs160_] = rhs246_
                d_1519_v20_ = rhs247_
                (globalState).f11 = ((d_1496_v1_ if d_1495_v0_ else _dafny.SeqWithoutIsStrInference([True, d_1495_v0_]))) == (d_1496_v1_)
                d_1520_v21_: C0
                nw270_ = C0()
                nw270_.ctor__(d_1495_v0_)
                d_1520_v21_ = nw270_
                d_1503_v5_ = (d_1503_v5_) + (d_1503_v5_)
                d_1521_v22_: T1
                nw271_ = C2()
                nw271_.ctor__(D0_DC1(d_1495_v0_), ((d_1506_v8_)[default__.safeIndex(115, (d_1506_v8_).length(0))] if (d_1520_v21_).f23 else (d_1506_v8_)[default__.safeIndex(115, (d_1506_v8_).length(0))]))
                d_1521_v22_ = nw271_
                nw272_ = C4()
                nw272_.ctor__((d_1506_v8_)[default__.safeIndex(115, (d_1506_v8_).length(0))])
                d_1521_v22_ = nw272_
            elif True:
                (globalState).f17 = (p0).issubset(p0)
                d_1522_v23_: _dafny.Map
                d_1522_v23_ = _dafny.Map({_dafny.Set({d_1495_v0_}): d_1495_v0_})
                r0 = default__.fm9((self).f18, default__.fm5(len(d_1504_v6_), globalState), ((d_1522_v23_)[default__.fm20(globalState)] if (default__.fm20(globalState)) in (d_1522_v23_) else True), globalState)
                d_1523_v24_: _dafny.Map
                d_1523_v24_ = _dafny.Map({(self).f18: d_1495_v0_})
                d_1524_v25_: str
                d_1524_v25_ = _dafny.CodePoint('n')
                (globalState).f15 = (d_1495_v0_ if default__.fm0(((d_1523_v24_)[(self).f18] if ((self).f18) in (d_1523_v24_) else d_1495_v0_), d_1524_v25_, d_1495_v0_, globalState) else default__.fm0(d_1495_v0_, _dafny.CodePoint('q'), d_1495_v0_, globalState))
                d_1495_v0_ = ((self).f18) < ((0) - ((self).f18))
                (globalState).f16 = d_1496_v1_
        elif True:
            if d_1495_v0_:
                d_1525_v26_: str
                d_1525_v26_ = _dafny.CodePoint('c')
                d_1526_v27_: _dafny.Array
                def lambda89_(d_1527_v0_):
                    def lambda90_(d_1528_i3_):
                        return d_1527_v0_

                    return lambda90_

                init46_ = lambda89_(d_1495_v0_)
                nw273_ = _dafny.Array(None, 13)
                for i0_46_ in range(nw273_.length(0)):
                    nw273_[i0_46_] = init46_(i0_46_)
                d_1526_v27_ = nw273_
                index217_ = default__.safeIndex(67, (d_1526_v27_).length(0))
                (d_1526_v27_)[index217_] = d_1495_v0_
                index218_ = default__.safeIndex(67, (d_1526_v27_).length(0))
                rhs248_ = _dafny.CodePoint('o')
                rhs249_ = not(d_1495_v0_)
                lhs161_ = d_1526_v27_
                lhs162_ = default__.safeIndex(67, (d_1526_v27_).length(0))
                d_1525_v26_ = rhs248_
                lhs161_[lhs162_] = rhs249_
                r0 = (self).f18
                r0 = (self).f18
                d_1525_v26_ = d_1525_v26_
                d_1529_v28_: _dafny.Map
                d_1529_v28_ = _dafny.Map({default__.safeModuloInt((self).f18, (self).f18): (0) - ((self).f18)})
                d_1530_v29_: _dafny.Array
                def lambda91_(d_1531_v0_, d_1532_p0_, d_1533_v26_):
                    def lambda92_(d_1534_i4_):
                        return (D4_DC14(d_1531_v0_, (self).f18, d_1532_p0_, d_1533_v26_)).cf26

                    return lambda92_

                init47_ = lambda91_(d_1495_v0_, p0, d_1525_v26_)
                nw274_ = _dafny.Array(None, 18)
                for i0_47_ in range(nw274_.length(0)):
                    nw274_[i0_47_] = init47_(i0_47_)
                d_1530_v29_ = nw274_
                d_1535_v30_: _dafny.Seq
                d_1535_v30_ = _dafny.SeqWithoutIsStrInference([d_1530_v29_, d_1530_v29_])
                d_1529_v28_ = (d_1529_v28_).set(default__.safeModuloInt((self).f18, len(d_1535_v30_)), (self).f18)
            elif True:
                d_1536_v31_: _dafny.Seq
                d_1536_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xm"))
                d_1537_v32_: _dafny.Array
                nw275_ = _dafny.Array(None, 5)
                nw275_[int(0)] = d_1495_v0_
                nw275_[int(1)] = d_1495_v0_
                nw275_[int(2)] = d_1495_v0_
                nw275_[int(3)] = (d_1536_v31_) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_1538_i5_ in range(default__.abs(542))]))
                nw275_[int(4)] = not ((d_1496_v1_)[default__.safeIndex((0) - ((self).f18), len(d_1496_v1_))]) or (d_1495_v0_)
                d_1537_v32_ = nw275_
                index219_ = default__.safeIndex(131, (d_1537_v32_).length(0))
                (d_1537_v32_)[index219_] = d_1495_v0_
                index220_ = default__.safeIndex(131, (d_1537_v32_).length(0))
                (d_1537_v32_)[index220_] = True
                d_1539_v33_: _dafny.MultiSet
                d_1539_v33_ = _dafny.MultiSet([(self).f18])
                d_1540_v34_: D2
                d_1540_v34_ = D2_DC7(d_1539_v33_)
                d_1541_v35_: _dafny.Seq
                d_1541_v35_ = _dafny.SeqWithoutIsStrInference([d_1540_v34_, d_1540_v34_, d_1540_v34_, d_1540_v34_])
                index221_ = default__.safeIndex(131, (d_1537_v32_).length(0))
                (d_1537_v32_)[index221_] = (_dafny.SeqWithoutIsStrInference([D2_DC7(_dafny.MultiSet([(self).f18])) for d_1542_i6_ in range(default__.abs(149))])) <= (d_1541_v35_)
                r0 = (0) - (len((d_1536_v31_) + (d_1536_v31_)))
                d_1543_v36_: _dafny.Map
                d_1543_v36_ = _dafny.Map({(self).f18: (self).f18})
                d_1544_v37_: _dafny.Array
                nw276_ = _dafny.Array(None, 14)
                nw276_[int(0)] = len((d_1536_v31_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kkptvhk"))))
                nw276_[int(1)] = (self).f18
                nw276_[int(2)] = (self).f18
                nw276_[int(3)] = len(d_1543_v36_)
                nw276_[int(4)] = (self).f18
                nw276_[int(5)] = (self).f18
                nw276_[int(6)] = (len(d_1536_v31_)) * (151)
                nw276_[int(7)] = ((self).f18) * ((self).f18)
                nw276_[int(8)] = len(d_1543_v36_)
                nw276_[int(9)] = 453
                nw276_[int(10)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ddlyqlq")))
                nw276_[int(11)] = (self).f18
                nw276_[int(12)] = default__.safeModuloInt((self).f18, (self).f18)
                nw276_[int(13)] = (self).f18
                d_1544_v37_ = nw276_
                index222_ = default__.safeIndex(152, (d_1544_v37_).length(0))
                (d_1544_v37_)[index222_] = (self).f18
                index223_ = default__.safeIndex(152, (d_1544_v37_).length(0))
                (d_1544_v37_)[index223_] = 476
                d_1545_v38_: _dafny.Map
                d_1545_v38_ = _dafny.Map({(d_1537_v32_)[default__.safeIndex(131, (d_1537_v32_).length(0))]: -546})
                d_1546_v39_: D3
                d_1546_v39_ = D3_DC11(d_1545_v38_, (self).f18, d_1544_v37_)
                d_1546_v39_ = D3_DC11(d_1545_v38_, (self).f18, d_1544_v37_)
            r0 = (self).f18
            d_1547_v40_: _dafny.MultiSet
            d_1547_v40_ = _dafny.MultiSet([d_1495_v0_])
            r0 = default__.safeDivisionInt(default__.fm5(((D10_DC27(d_1547_v40_)).cf48).cardinality, globalState), (self).f18)
            d_1548_v41_: _dafny.Seq
            d_1548_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "datkd"))
            d_1548_v41_ = (d_1548_v41_) + (d_1548_v41_)
            d_1549_v42_: C4
            nw277_ = C4()
            nw277_.ctor__(d_1495_v0_)
            d_1549_v42_ = nw277_
            d_1550_v43_: _dafny.Map
            d_1550_v43_ = _dafny.Map({d_1549_v42_: (self).f18})
            d_1551_v44_: D0
            d_1551_v44_ = D0_DC1(d_1495_v0_)
            d_1552_v45_: C2
            nw278_ = C2()
            nw278_.ctor__(d_1551_v44_, d_1495_v0_)
            d_1552_v45_ = nw278_
            d_1553_v46_: _dafny.Array
            nw279_ = _dafny.Array(_dafny.Seq({}), 10)
            d_1553_v46_ = nw279_
            d_1554_v47_: D14
            d_1554_v47_ = D14_DC41(d_1552_v45_, d_1548_v41_, d_1553_v46_)
            d_1555_v48_: D14
            d_1555_v48_ = D14_DC41((d_1554_v47_).cf72, d_1548_v41_, d_1553_v46_)
            d_1556_v49_: _dafny.Map
            d_1556_v49_ = _dafny.Map({d_1555_v48_: (self).f18})
            d_1557_v50_: _dafny.Seq
            d_1557_v50_ = _dafny.SeqWithoutIsStrInference([(self).f18])
            d_1558_v51_: _dafny.Map
            d_1558_v51_ = _dafny.Map({d_1495_v0_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jdetgprt")))})
            d_1559_v52_: _dafny.Map
            d_1559_v52_ = _dafny.Map({d_1548_v41_: (d_1558_v51_).set((d_1549_v42_).f19, (self).f18)})
            d_1560_v53_: str
            d_1560_v53_ = _dafny.CodePoint('c')
            d_1561_v54_: _dafny.Array
            nw280_ = _dafny.Array(None, 20)
            nw280_[int(0)] = default__.fm5((self).f18, globalState)
            nw280_[int(1)] = ((d_1550_v43_)[d_1549_v42_] if (d_1549_v42_) in (d_1550_v43_) else (self).f18)
            nw280_[int(2)] = (self).f18
            nw280_[int(3)] = (self).f18
            nw280_[int(4)] = ((self).f18) + ((self).f18)
            nw280_[int(5)] = len((p0).intersection(p0))
            nw280_[int(6)] = ((d_1556_v49_)[d_1554_v47_] if (d_1554_v47_) in (d_1556_v49_) else (self).f18)
            nw280_[int(7)] = (self).f18
            nw280_[int(8)] = (self).f18
            nw280_[int(9)] = (self).f18
            nw280_[int(10)] = default__.fm5(28, globalState)
            nw280_[int(11)] = (self).f18
            nw280_[int(12)] = (667) + ((self).f18)
            nw280_[int(13)] = (self).f18
            nw280_[int(14)] = (self).f18
            nw280_[int(15)] = (self).f18
            nw280_[int(16)] = len(default__.fm19(d_1557_v50_, d_1495_v0_, globalState))
            nw280_[int(17)] = (self).f18
            nw280_[int(18)] = (self).f18
            nw280_[int(19)] = ((0) - ((self).f18) if (d_1549_v42_).fm4((d_1549_v42_).f19, d_1559_v52_, d_1560_v53_, (self).f18, globalState) else (self).f18)
            d_1561_v54_ = nw280_
            index224_ = default__.safeIndex(546, (d_1561_v54_).length(0))
            (d_1561_v54_)[index224_] = 611
            index225_ = default__.safeIndex(546, (d_1561_v54_).length(0))
            (d_1561_v54_)[index225_] = (self).f18
        d_1562_v55_: _dafny.Seq
        d_1562_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftlcnhg"))
        d_1563_v56_: _dafny.Map
        d_1563_v56_ = _dafny.Map({(self).f18: (self).f18})
        d_1564_v57_: D0
        d_1564_v57_ = D0_DC1(d_1495_v0_)
        d_1565_v58_: str
        d_1565_v58_ = _dafny.CodePoint('b')
        d_1566_v59_: _dafny.Map
        d_1566_v59_ = _dafny.Map({d_1564_v57_: d_1565_v58_})
        d_1567_v60_: _dafny.Map
        d_1567_v60_ = _dafny.Map({(d_1496_v1_)[default__.safeIndex((self).f18, len(d_1496_v1_))]: (self).f18})
        d_1568_v61_: _dafny.MultiSet
        d_1568_v61_ = _dafny.MultiSet([d_1495_v0_])
        d_1569_v62_: _dafny.Set
        d_1569_v62_ = _dafny.Set({d_1568_v61_, d_1568_v61_})
        d_1570_v63_: _dafny.Array
        nw281_ = _dafny.Array(None, 23)
        nw281_[int(0)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1571_i7_ in range(default__.abs(768))]))
        nw281_[int(1)] = (self).f18
        nw281_[int(2)] = len(d_1562_v55_)
        nw281_[int(3)] = (self).f18
        nw281_[int(4)] = 306
        nw281_[int(5)] = (self).f18
        nw281_[int(6)] = (self).f18
        nw281_[int(7)] = (self).f18
        nw281_[int(8)] = (0) - ((self).f18)
        nw281_[int(9)] = (self).f18
        nw281_[int(10)] = len((d_1563_v56_).set((self).f18, (self).f18))
        nw281_[int(11)] = (self).f18
        nw281_[int(12)] = (self).f18
        nw281_[int(13)] = (0) - (len(d_1566_v59_))
        nw281_[int(14)] = ((d_1567_v60_)[False] if (False) in (d_1567_v60_) else (self).f18)
        nw281_[int(15)] = (0) - (len(d_1569_v62_))
        nw281_[int(16)] = (self).f18
        nw281_[int(17)] = (self).f18
        nw281_[int(18)] = (self).f18
        nw281_[int(19)] = len(d_1562_v55_)
        nw281_[int(20)] = (0) - ((self).f18)
        nw281_[int(21)] = (self).f18
        nw281_[int(22)] = (self).f18
        d_1570_v63_ = nw281_
        d_1572_v64_: _dafny.Map
        d_1572_v64_ = _dafny.Map({d_1570_v63_: d_1562_v55_})
        d_1572_v64_ = d_1572_v64_
        d_1573_i8_: int
        d_1573_i8_ = 0
        with _dafny.label("11"):
            while d_1495_v0_:
                with _dafny.c_label("11"):
                    if (d_1573_i8_) >= (100):
                        raise _dafny.Break("11")
                    d_1573_i8_ = (d_1573_i8_) + (1)
                    d_1574_v65_: _dafny.Map
                    d_1574_v65_ = _dafny.Map({(self).f18: d_1495_v0_})
                    d_1575_v66_: _dafny.Set
                    d_1575_v66_ = _dafny.Set({(default__.fm17(d_1495_v0_, d_1495_v0_, globalState)).set((self).f18, default__.fm0(d_1495_v0_, _dafny.CodePoint('i'), d_1495_v0_, globalState)), d_1574_v65_})
                    source15_ = D3_DC9(d_1575_v66_)
                    if source15_.is_DC10:
                        d_1576___mcc_h0_ = source15_.cf17
                        d_1577_cf17_ = d_1576___mcc_h0_
                        (globalState).f11 = ((self).f18) >= ((0) - ((0) - ((0) - (d_1577_cf17_))))
                        d_1578_v67_: _dafny.Array
                        nw282_ = _dafny.Array(False, 17)
                        d_1578_v67_ = nw282_
                        d_1578_v67_ = d_1578_v67_
                        index226_ = default__.safeIndex(40, (d_1578_v67_).length(0))
                        (d_1578_v67_)[index226_] = d_1495_v0_
                        d_1579_v68_: _dafny.Array
                        nw283_ = _dafny.Array(None, 11)
                        nw283_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1565_v58_ for d_1580_i9_ in range(default__.abs(508))])
                        nw283_[int(1)] = d_1562_v55_
                        nw283_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mk"))
                        nw283_[int(3)] = d_1562_v55_
                        nw283_[int(4)] = d_1562_v55_
                        nw283_[int(5)] = d_1562_v55_
                        nw283_[int(6)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qpit"))
                        nw283_[int(7)] = d_1562_v55_
                        nw283_[int(8)] = d_1562_v55_
                        nw283_[int(9)] = (d_1562_v55_).set(default__.safeIndex(d_1577_cf17_, len(d_1562_v55_)), d_1565_v58_)
                        nw283_[int(10)] = d_1562_v55_
                        d_1579_v68_ = nw283_
                        d_1581_v69_: _dafny.Array
                        nw284_ = _dafny.Array(None, 17)
                        nw284_[int(0)] = d_1579_v68_
                        nw284_[int(1)] = d_1579_v68_
                        nw284_[int(2)] = d_1579_v68_
                        nw284_[int(3)] = d_1579_v68_
                        nw284_[int(4)] = (d_1579_v68_ if d_1495_v0_ else d_1579_v68_)
                        nw284_[int(5)] = d_1579_v68_
                        nw284_[int(6)] = d_1579_v68_
                        nw284_[int(7)] = d_1579_v68_
                        nw284_[int(8)] = d_1579_v68_
                        nw284_[int(9)] = d_1579_v68_
                        nw284_[int(10)] = d_1579_v68_
                        nw284_[int(11)] = d_1579_v68_
                        nw284_[int(12)] = d_1579_v68_
                        nw284_[int(13)] = d_1579_v68_
                        nw284_[int(14)] = d_1579_v68_
                        nw284_[int(15)] = d_1579_v68_
                        nw284_[int(16)] = (d_1579_v68_ if d_1495_v0_ else d_1579_v68_)
                        d_1581_v69_ = nw284_
                        index227_ = default__.safeIndex(274, (d_1581_v69_).length(0))
                        (d_1581_v69_)[index227_] = d_1579_v68_
                        index228_ = default__.safeIndex(40, (d_1578_v67_).length(0))
                        index229_ = default__.safeIndex(274, (d_1581_v69_).length(0))
                        rhs250_ = (d_1496_v1_)[default__.safeIndex(default__.safeDivisionInt((self).f18, (self).f18), len(d_1496_v1_))]
                        rhs251_ = d_1579_v68_
                        lhs163_ = d_1578_v67_
                        lhs164_ = default__.safeIndex(40, (d_1578_v67_).length(0))
                        lhs165_ = d_1581_v69_
                        lhs166_ = default__.safeIndex(274, (d_1581_v69_).length(0))
                        lhs163_[lhs164_] = rhs250_
                        lhs165_[lhs166_] = rhs251_
                        d_1563_v56_ = (d_1563_v56_).set(d_1577_cf17_, len(d_1562_v55_))
                    elif source15_.is_DC11:
                        d_1582___mcc_h1_ = source15_.cf18
                        d_1583___mcc_h2_ = source15_.cf19
                        d_1584___mcc_h3_ = source15_.cf20
                        d_1585_cf20_ = d_1584___mcc_h3_
                        d_1586_cf19_ = d_1583___mcc_h2_
                        d_1587_cf18_ = d_1582___mcc_h1_
                        d_1586_cf19_ = default__.fm9((self).f18, (self).f18, d_1495_v0_, globalState)
                        d_1588_v70_: _dafny.MultiSet
                        d_1588_v70_ = _dafny.MultiSet([d_1586_cf19_])
                        d_1589_v71_: _dafny.Set
                        d_1589_v71_ = _dafny.Set({d_1588_v70_})
                        (globalState).f15 = ((d_1574_v65_)[((self).f18) + (d_1586_cf19_)] if (((self).f18) + (d_1586_cf19_)) in (d_1574_v65_) else (d_1589_v71_).issubset(d_1589_v71_))
                        d_1565_v58_ = d_1565_v58_
                        d_1495_v0_ = d_1495_v0_
                    elif True:
                        d_1590___mcc_h4_ = source15_.cf16
                        d_1591_cf16_ = d_1590___mcc_h4_
                        (globalState).f17 = d_1495_v0_
                        d_1592_v72_: T1
                        nw285_ = C2()
                        nw285_.ctor__(d_1564_v57_, d_1495_v0_)
                        d_1592_v72_ = nw285_
                        d_1593_v73_: _dafny.Array
                        nw286_ = _dafny.Array(None, 14)
                        nw286_[int(0)] = d_1592_v72_
                        nw286_[int(1)] = d_1592_v72_
                        nw286_[int(2)] = d_1592_v72_
                        nw286_[int(3)] = d_1592_v72_
                        nw286_[int(4)] = d_1592_v72_
                        nw286_[int(5)] = d_1592_v72_
                        nw286_[int(6)] = d_1592_v72_
                        nw286_[int(7)] = d_1592_v72_
                        nw286_[int(8)] = d_1592_v72_
                        nw286_[int(9)] = d_1592_v72_
                        nw286_[int(10)] = d_1592_v72_
                        nw286_[int(11)] = d_1592_v72_
                        nw286_[int(12)] = d_1592_v72_
                        nw286_[int(13)] = d_1592_v72_
                        d_1593_v73_ = nw286_
                        index230_ = default__.safeIndex(333, (d_1593_v73_).length(0))
                        (d_1593_v73_)[index230_] = d_1592_v72_
                        index231_ = default__.safeIndex(333, (d_1593_v73_).length(0))
                        (d_1593_v73_)[index231_] = d_1592_v72_
                        (globalState).f15 = (d_1495_v0_ if d_1495_v0_ else d_1495_v0_)
                        index232_ = default__.safeIndex(498, (d_1570_v63_).length(0))
                        (d_1570_v63_)[index232_] = (0) - (len((d_1562_v55_) + (d_1562_v55_)))
                        index233_ = default__.safeIndex(498, (d_1570_v63_).length(0))
                        (d_1570_v63_)[index233_] = default__.safeModuloInt((self).f18, (self).f18)
                    d_1570_v63_ = d_1570_v63_
                    d_1594_v74_: _dafny.Array
                    nw287_ = _dafny.Array(None, 13)
                    nw287_[int(0)] = d_1495_v0_
                    nw287_[int(1)] = d_1495_v0_
                    nw287_[int(2)] = d_1495_v0_
                    nw287_[int(3)] = d_1495_v0_
                    nw287_[int(4)] = d_1495_v0_
                    nw287_[int(5)] = d_1495_v0_
                    nw287_[int(6)] = d_1495_v0_
                    nw287_[int(7)] = d_1495_v0_
                    nw287_[int(8)] = d_1495_v0_
                    nw287_[int(9)] = d_1495_v0_
                    nw287_[int(10)] = d_1495_v0_
                    nw287_[int(11)] = not(d_1495_v0_)
                    nw287_[int(12)] = d_1495_v0_
                    d_1594_v74_ = nw287_
                    d_1595_v75_: _dafny.Map
                    d_1595_v75_ = _dafny.Map({((d_1568_v61_)[d_1495_v0_] if (d_1495_v0_) in (d_1568_v61_) else (self).f18): d_1594_v74_})
                    d_1596_v76_: _dafny.Array
                    nw288_ = _dafny.Array(None, 29)
                    d_1596_v76_ = nw288_
                    d_1597_v77_: C4
                    nw289_ = C4()
                    nw289_.ctor__(d_1495_v0_)
                    d_1597_v77_ = nw289_
                    index234_ = default__.safeIndex(698, (d_1596_v76_).length(0))
                    (d_1596_v76_)[index234_] = d_1597_v77_
                    d_1598_v78_: _dafny.Map
                    d_1598_v78_ = _dafny.Map({d_1495_v0_: d_1565_v58_})
                    index235_ = default__.safeIndex(698, (d_1596_v76_).length(0))
                    rhs252_ = (d_1595_v75_) | (_dafny.Map({(self).f18: d_1594_v74_}))
                    rhs253_ = (default__.fm16(default__.fm0(d_1495_v0_, ((d_1598_v78_)[(d_1597_v77_).f19] if ((d_1597_v77_).f19) in (d_1598_v78_) else d_1565_v58_), (d_1597_v77_).f19, globalState), (self).f18, globalState)) < (_dafny.SeqWithoutIsStrInference([d_1565_v58_ for d_1599_i10_ in range(default__.abs(432))]))
                    rhs254_ = (0) - ((self).f18)
                    rhs255_ = d_1597_v77_
                    lhs167_ = d_1596_v76_
                    lhs168_ = default__.safeIndex(698, (d_1596_v76_).length(0))
                    d_1595_v75_ = rhs252_
                    d_1495_v0_ = rhs253_
                    r0 = rhs254_
                    lhs167_[lhs168_] = rhs255_
                    d_1600_v79_: _dafny.Map
                    d_1600_v79_ = _dafny.Map({D4_DC13((self).f18): d_1495_v0_})
                    (globalState).f15 = not((len(d_1600_v79_)) not in ((p0) | (p0)))
                    pass
            pass
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1570_v63_).length(0)):
            d_1601_i11_: int = guard_loop_3_
            if (True) and (((0) <= (d_1601_i11_)) and ((d_1601_i11_) < ((d_1570_v63_).length(0)))):
                (d_1570_v63_)[(d_1601_i11_)] = default__.safeModuloInt(d_1601_i11_, (D11_DC30(False, (self).f18, d_1495_v0_)).cf53)
        d_1562_v55_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "up"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "calngaa")))
        (globalState).f17 = not (not(d_1495_v0_)) or (d_1495_v0_)
        r0 = default__.safeModuloInt((0) - ((self).f18), (0) - (default__.safeDivisionInt((self).f18, (self).f18)))
        return r0

    @property
    def f18(self):
        return self._f18
