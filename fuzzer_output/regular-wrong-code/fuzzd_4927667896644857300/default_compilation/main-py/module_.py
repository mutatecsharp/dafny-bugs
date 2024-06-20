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
        source0_ = D0_DC1(-4)
        if source0_.is_DC1:
            d_0___mcc_h0_ = source0_.cf1
            d_1_cf1_ = d_0___mcc_h0_
            return _dafny.SeqWithoutIsStrInference([not(True)])
        elif source0_.is_DC2:
            d_2___mcc_h1_ = source0_.cf2
            d_3___mcc_h2_ = source0_.cf3
            d_4_cf3_ = d_3___mcc_h2_
            d_5_cf2_ = d_2___mcc_h1_
            return _dafny.SeqWithoutIsStrInference([False])
        elif True:
            d_6___mcc_h3_ = source0_.cf0
            d_7_cf0_ = d_6___mcc_h3_
            return (_dafny.SeqWithoutIsStrInference([False])) + (_dafny.SeqWithoutIsStrInference([False, not(not(False))]))

    @staticmethod
    def fm1(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Map()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(440, -849):
                d_8_v0_: int = compr_0_
                if ((440) <= (d_8_v0_)) and ((d_8_v0_) < (-849)):
                    coll0_[(d_8_v0_) + (245)] = 299
            return _dafny.Map(coll0_)
        return ((_dafny.Set({_dafny.Map({True: 525})}) if True else _dafny.Set({_dafny.Map({False: -401}), _dafny.Map({True: len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jb")))}))})}))) - ((_dafny.Set({_dafny.Map({False: len(iife0_()
        )})})).intersection(_dafny.Set({_dafny.Map({True: len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kp"))): False}))})})))

    @staticmethod
    def fm2(p0, p1, p2, p3, globalState):
        return (0) - ((len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_9_i0_ in range(default__.abs(873))]))) + ((0) - ((len(_dafny.Map({True: len(_dafny.Map({(D1_DC5(False, False, (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))))), True, True)).cf9: False}))}))) - ((_dafny.MultiSet([len(_dafny.Map({D20_DC48(_dafny.MultiSet([-469])): 362})), 810])).cardinality))))

    @staticmethod
    def fm3(globalState):
        return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mromwiucb"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")))

    @staticmethod
    def fm4(globalState):
        return not ((_dafny.Map({True: 634})) == (_dafny.Map({True: 112}))) or (True)

    @staticmethod
    def fm5(p0, p1, globalState):
        return _dafny.CodePoint('n')

    @staticmethod
    def fm6(p0, p1, p2, globalState):
        def iife1_():
            coll1_ = _dafny.Set()
            compr_1_: int
            for compr_1_ in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([181]))).Elements:
                d_14_v0_: int = compr_1_
                if (d_14_v0_) in (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([181]))):
                    coll1_ = coll1_.union(_dafny.Set([(d_14_v0_) + (-324)]))
            return _dafny.Set(coll1_)
        return ((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([696])) for d_10_i0_ in range(default__.abs(835))])) + (_dafny.SeqWithoutIsStrInference([-980, len(_dafny.Map({not(not(False)): not(not(False))}))]))) + ((_dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('m') for d_11_i2_ in range(default__.abs(695))]))))})) for d_12_i1_ in range(default__.abs(655))])) + (_dafny.SeqWithoutIsStrInference([-640, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_13_i3_ in range(default__.abs(141))])): _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mno"))): True})})), len(iife1_()
        ), len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_15_i4_ in range(default__.abs(794))])), -636])))

    @staticmethod
    def fm10(p0, p1, p2, globalState):
        return (_dafny.MultiSet([_dafny.Map({not(False): False})])).intersection((_dafny.MultiSet([_dafny.Map({True: False})])).intersection(_dafny.MultiSet([_dafny.Map({False: not(True)})])))

    @staticmethod
    def fm18(p0, p1, globalState):
        return ((_dafny.Map({True: _dafny.CodePoint('u')})) | (_dafny.Map({True: _dafny.CodePoint('e')}))) | ((_dafny.Map({False: _dafny.CodePoint('l')})) | (_dafny.Map({False: _dafny.CodePoint('u')})))

    @staticmethod
    def fm20(p0, p1, p2, p3, globalState):
        if False:
            return D1_DC3(False)
        elif True:
            return D1_DC3(False)

    @staticmethod
    def fm21(p0, globalState):
        return (_dafny.Map({924: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qeevq")))})) | ((_dafny.Map({(_dafny.MultiSet([len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xe"))): not(False)}))])).cardinality: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lslmesr")))})) | (_dafny.Map({399: len(_dafny.SeqWithoutIsStrInference([D7_DC18(True)]))})))

    @staticmethod
    def fm22(p0, p1, p2, globalState):
        return ((_dafny.Map({_dafny.SeqWithoutIsStrInference([False]): len(_dafny.Map({False: (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), 957]))).cardinality}))})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([not(True), True]): -479}))) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([True, (D12_DC30(False, not(False), not(False), 436, True)).cf55]): (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t')])), -118, len(_dafny.Set({False, True, True, True})), len(_dafny.SeqWithoutIsStrInference([False])), -115])).cardinality}))

    @staticmethod
    def fm23(globalState):
        def iife2_():
            coll2_ = _dafny.Map()
            compr_2_: int
            for compr_2_ in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True, False, False, False]))])).Elements:
                d_16_v0_: int = compr_2_
                if (d_16_v0_) in (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False, True, False, False, False]))])):
                    coll2_[(d_16_v0_) * (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_17_i0_ in range(default__.abs(890))])))] = True
            return _dafny.Map(coll2_)
        def iife3_():
            coll3_ = _dafny.Map()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(608, 838):
                d_18_v1_: int = compr_3_
                if ((608) <= (d_18_v1_)) and ((d_18_v1_) < (838)):
                    coll3_[(d_18_v1_) + (190)] = True
            return _dafny.Map(coll3_)
        return ((_dafny.Map({738: True})) | (iife2_()
        )) | (iife3_()
        )

    @staticmethod
    def fm24(p0, p1, globalState):
        return (_dafny.Set({False, not(True), True, False})) | (_dafny.Set({not(False)}))

    @staticmethod
    def fm25(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([(False if not(not(False)) else True)])

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return ((_dafny.MultiSet([len(_dafny.Map({False: True}))])).intersection(_dafny.MultiSet([(D0_DC2(574, -350)).cf3]))).intersection((D20_DC48((D20_DC48(_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.Map({False: False}))]))).cf89)).cf89)

    @staticmethod
    def fm27(p0, globalState):
        def iife4_():
            def iife6_():
                coll6_ = _dafny.Set()
                compr_6_: D4
                for compr_6_ in (_dafny.Set({D4_DC10(_dafny.CodePoint('g'))})).Elements:
                    d_21_v1_: D4 = compr_6_
                    if (d_21_v1_) in (_dafny.Set({D4_DC10(_dafny.CodePoint('g'))})):
                        coll6_ = coll6_.union(_dafny.Set([d_21_v1_]))
                return _dafny.Set(coll6_)
            coll4_ = _dafny.Map()
            def iife5_():
                coll5_ = _dafny.Set()
                compr_5_: D4
                for compr_5_ in (_dafny.Set({D4_DC10(_dafny.CodePoint('g'))})).Elements:
                    d_19_v1_: D4 = compr_5_
                    if (d_19_v1_) in (_dafny.Set({D4_DC10(_dafny.CodePoint('g'))})):
                        coll5_ = coll5_.union(_dafny.Set([d_19_v1_]))
                return _dafny.Set(coll5_)
            compr_4_: D4
            for compr_4_ in (iife5_()
            ).Elements:
                d_20_v0_: D4 = compr_4_
                if (d_20_v0_) in (iife6_()
                ):
                    def iife7_():
                        coll7_ = _dafny.Map()
                        compr_7_: int
                        for compr_7_ in _dafny.IntegerRange(-672, 151):
                            d_22_v2_: int = compr_7_
                            if ((-672) <= (d_22_v2_)) and ((d_22_v2_) < (151)):
                                coll7_[(d_22_v2_) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jul"))))] = 532
                        return _dafny.Map(coll7_)
                    coll4_[d_20_v0_] = len(iife7_()
                    )
            return _dafny.Map(coll4_)
        return (iife4_()
        ) | ((_dafny.Map({D4_DC10(_dafny.CodePoint('w')): len(_dafny.Map({False: 998}))})) | (_dafny.Map({D4_DC10(_dafny.CodePoint('d')): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_23_i0_ in range(default__.abs(5))]))})))

    @staticmethod
    def fm28(p0, globalState):
        def iife8_():
            coll8_ = _dafny.Map()
            compr_8_: int
            for compr_8_ in _dafny.IntegerRange(68, 222):
                d_24_v0_: int = compr_8_
                if ((68) <= (d_24_v0_)) and ((d_24_v0_) < (222)):
                    coll8_[default__.safeDivisionInt(d_24_v0_, 547)] = True
            return _dafny.Map(coll8_)
        return _dafny.Set({D8_DC21(True, False, 100), D8_DC21(not(True), True, len(_dafny.Set({168}))), D8_DC21(True, True, len(iife8_()
))})

    @staticmethod
    def fm29(p0, globalState):
        return D1_DC5(True, (D12_DC29(False)).cf50, -663, False, not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgnxwcy"))) <= ((D2_DC7(True, False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a')]))).cf15)))

    @staticmethod
    def fm30(p0, p1, globalState):
        return _dafny.Map({not((_dafny.CodePoint('q')) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "st")))): (-677) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len((D16_DC37(_dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tytqu")), _dafny.CodePoint('f'))).cf65), -68])) for d_25_i0_ in range(default__.abs(-343))])))})

    @staticmethod
    def fm31(p0, p1, p2, p3, globalState):
        def iife9_():
            coll9_ = _dafny.Set()
            compr_9_: int
            for compr_9_ in (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkogaj")))): _dafny.CodePoint('q')})).keys.Elements:
                d_26_v0_: int = compr_9_
                if (d_26_v0_) in (_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fkogaj")))): _dafny.CodePoint('q')})):
                    coll9_ = coll9_.union(_dafny.Set([default__.safeDivisionInt(d_26_v0_, 551)]))
            return _dafny.Set(coll9_)
        return (_dafny.MultiSet([iife9_()
        , _dafny.Set({-965})])).intersection((_dafny.MultiSet([_dafny.Set({(_dafny.MultiSet([(0) - (len(_dafny.Map({_dafny.CodePoint('x'): True}))), 922])).cardinality})])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.Set({len(_dafny.Set({False}))})]))))

    @staticmethod
    def fm32(p0, globalState):
        return D0_DC0(_dafny.Map({593: len(_dafny.Map({False: False}))}))

    @staticmethod
    def fm33(p0, globalState):
        source1_ = (D8_DC20((D12_DC30(True, False, True, len(_dafny.Map({True: False})), False)).cf54, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([True, True, True, False]), _dafny.SeqWithoutIsStrInference([not(False)]), _dafny.SeqWithoutIsStrInference([True, not(False)]), _dafny.SeqWithoutIsStrInference([True]), _dafny.SeqWithoutIsStrInference([False])]))) if True else D8_DC20(-413, len(_dafny.Set({160}))))
        if source1_.is_DC20:
            d_27___mcc_h0_ = source1_.cf31
            d_28___mcc_h1_ = source1_.cf32
            d_29_cf32_ = d_28___mcc_h1_
            d_30_cf31_ = d_27___mcc_h0_
            def iife10_():
                coll10_ = _dafny.Map()
                compr_10_: _dafny.Seq
                for compr_10_ in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_29_cf32_])])).Elements:
                    d_31_v0_: _dafny.Seq = compr_10_
                    if (d_31_v0_) in (_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_29_cf32_])])):
                        coll10_[d_31_v0_] = _dafny.CodePoint('f')
                return _dafny.Map(coll10_)
            return iife10_()
            
        elif source1_.is_DC21:
            d_32___mcc_h2_ = source1_.cf33
            d_33___mcc_h3_ = source1_.cf34
            d_34___mcc_h4_ = source1_.cf35
            d_35_cf35_ = d_34___mcc_h4_
            d_36_cf34_ = d_33___mcc_h3_
            d_37_cf33_ = d_32___mcc_h2_
            def iife11_():
                coll11_ = _dafny.Map()
                compr_11_: _dafny.Seq
                for compr_11_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_35_cf35_ for d_38_i0_ in range(default__.abs(405))])})).Elements:
                    d_39_v1_: _dafny.Seq = compr_11_
                    if (d_39_v1_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([d_35_cf35_ for d_38_i0_ in range(default__.abs(405))])})):
                        coll11_[d_39_v1_] = _dafny.CodePoint('x')
                return _dafny.Map(coll11_)
            return iife11_()
            
        elif source1_.is_DC19:
            d_40___mcc_h5_ = source1_.cf30
            d_41_cf30_ = d_40___mcc_h5_
            if not(False):
                def iife12_():
                    coll12_ = _dafny.Map()
                    compr_12_: _dafny.Seq
                    for compr_12_ in (_dafny.Set({_dafny.SeqWithoutIsStrInference([781])})).Elements:
                        d_42_v2_: _dafny.Seq = compr_12_
                        if (d_42_v2_) in (_dafny.Set({_dafny.SeqWithoutIsStrInference([781])})):
                            coll12_[d_42_v2_] = _dafny.CodePoint('g')
                    return _dafny.Map(coll12_)
                return iife12_()
                
            elif True:
                def iife13_():
                    coll13_ = _dafny.Map()
                    compr_13_: _dafny.Seq
                    for compr_13_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([627]) for d_43_i1_ in range(default__.abs(490))])).Elements:
                        d_44_v3_: _dafny.Seq = compr_13_
                        if (d_44_v3_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([627]) for d_43_i1_ in range(default__.abs(490))])):
                            coll13_[d_44_v3_] = _dafny.CodePoint('k')
                    return _dafny.Map(coll13_)
                return iife13_()
                
        elif True:
            d_45___mcc_h6_ = source1_.cf36
            d_46_cf36_ = d_45___mcc_h6_
            return _dafny.Map({_dafny.SeqWithoutIsStrInference([-559]): _dafny.CodePoint('y')})

    @staticmethod
    def fm34(globalState):
        return _dafny.Set({default__.safeModuloInt(len(_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rrwa")))})), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "opxxf")))), 897, ((_dafny.MultiSet([188, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xmiisjrm")))])) | (_dafny.MultiSet([-276]))).cardinality})

    @staticmethod
    def fm35(p0, p1, p2, p3, globalState):
        def iife14_():
            coll14_ = _dafny.Map()
            compr_14_: int
            for compr_14_ in (_dafny.SeqWithoutIsStrInference([-723])).Elements:
                d_47_v0_: int = compr_14_
                if (d_47_v0_) in (_dafny.SeqWithoutIsStrInference([-723])):
                    coll14_[(d_47_v0_) + (171)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
            return _dafny.Map(coll14_)
        return ((_dafny.Map({len(_dafny.SeqWithoutIsStrInference([165])): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gxs"))})) | (iife14_()
        )) | (_dafny.Map({79: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "asold"))}))

    @staticmethod
    def fm36(globalState):
        return (_dafny.Set({D8_DC21(True, False, (_dafny.MultiSet([False, False, True])).cardinality)}) if True else _dafny.Set({D8_DC21(True, True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqb"))))}))

    @staticmethod
    def fm37(p0, p1, p2, globalState):
        return D8_DC22(D8_DC20(703, -144))

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return D15_DC34(True, (0) - (default__.safeDivisionInt(962, 458)))

    @staticmethod
    def fm39(p0, globalState):
        def iife15_():
            def iife23_():
                coll23_ = _dafny.Map()
                compr_23_: int
                for compr_23_ in _dafny.IntegerRange(123, 410):
                    d_49_v1_: int = compr_23_
                    if ((123) <= (d_49_v1_)) and ((d_49_v1_) < (410)):
                        coll23_[default__.safeDivisionInt(d_49_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmhpep"))))] = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('j'): False})) for d_50_i1_ in range(default__.abs(56))]))).cardinality: -42})
                return _dafny.Map(coll23_)
            def iife24_():
                coll24_ = _dafny.Map()
                compr_24_: int
                for compr_24_ in _dafny.IntegerRange(204, 267):
                    d_51_v2_: int = compr_24_
                    if ((204) <= (d_51_v2_)) and ((d_51_v2_) < (267)):
                        coll24_[default__.safeDivisionInt(d_51_v2_, 60)] = _dafny.CodePoint('g')
                return _dafny.Map(coll24_)
            def iife25_():
                coll25_ = _dafny.Map()
                compr_25_: int
                for compr_25_ in _dafny.IntegerRange(167, 217):
                    d_52_v3_: int = compr_25_
                    if ((167) <= (d_52_v3_)) and ((d_52_v3_) < (217)):
                        coll25_[(d_52_v3_) + (344)] = _dafny.CodePoint('q')
                return _dafny.Map(coll25_)
            def iife26_():
                def iife28_():
                    coll28_ = _dafny.Set()
                    compr_28_: int
                    for compr_28_ in _dafny.IntegerRange(-217, -953):
                        d_60_v5_: int = compr_28_
                        if ((-217) <= (d_60_v5_)) and ((d_60_v5_) < (-953)):
                            coll28_ = coll28_.union(_dafny.Set([(d_60_v5_) * (len(_dafny.Map({False: 11})))]))
                    return _dafny.Set(coll28_)
                coll26_ = _dafny.Map()
                def iife27_():
                    coll27_ = _dafny.Set()
                    compr_27_: int
                    for compr_27_ in _dafny.IntegerRange(-217, -953):
                        d_59_v5_: int = compr_27_
                        if ((-217) <= (d_59_v5_)) and ((d_59_v5_) < (-953)):
                            coll27_ = coll27_.union(_dafny.Set([(d_59_v5_) * (len(_dafny.Map({False: 11})))]))
                    return _dafny.Set(coll27_)
                compr_26_: int
                for compr_26_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(iife27_()
                )})).keys.Elements:
                    d_55_v4_: int = compr_26_
                    if (d_55_v4_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(iife28_()
                    )})):
                        coll26_[default__.safeDivisionInt(d_55_v4_, 650)] = _dafny.CodePoint('x')
                return _dafny.Map(coll26_)
            def iife29_():
                coll29_ = _dafny.Map()
                compr_29_: int
                for compr_29_ in _dafny.IntegerRange(614, 189):
                    d_57_v6_: int = compr_29_
                    if ((614) <= (d_57_v6_)) and ((d_57_v6_) < (189)):
                        coll29_[default__.safeDivisionInt(d_57_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvfmnoywj"))))] = _dafny.CodePoint('f')
                return _dafny.Map(coll29_)
            coll15_ = _dafny.Map()
            def iife16_():
                coll16_ = _dafny.Map()
                compr_16_: int
                for compr_16_ in _dafny.IntegerRange(123, 410):
                    d_49_v1_: int = compr_16_
                    if ((123) <= (d_49_v1_)) and ((d_49_v1_) < (410)):
                        coll16_[default__.safeDivisionInt(d_49_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nmhpep"))))] = _dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.CodePoint('j'): False})) for d_50_i1_ in range(default__.abs(56))]))).cardinality: -42})
                return _dafny.Map(coll16_)
            def iife17_():
                coll17_ = _dafny.Map()
                compr_17_: int
                for compr_17_ in _dafny.IntegerRange(204, 267):
                    d_51_v2_: int = compr_17_
                    if ((204) <= (d_51_v2_)) and ((d_51_v2_) < (267)):
                        coll17_[default__.safeDivisionInt(d_51_v2_, 60)] = _dafny.CodePoint('g')
                return _dafny.Map(coll17_)
            def iife18_():
                coll18_ = _dafny.Map()
                compr_18_: int
                for compr_18_ in _dafny.IntegerRange(167, 217):
                    d_52_v3_: int = compr_18_
                    if ((167) <= (d_52_v3_)) and ((d_52_v3_) < (217)):
                        coll18_[(d_52_v3_) + (344)] = _dafny.CodePoint('q')
                return _dafny.Map(coll18_)
            def iife19_():
                def iife21_():
                    coll21_ = _dafny.Set()
                    compr_21_: int
                    for compr_21_ in _dafny.IntegerRange(-217, -953):
                        d_56_v5_: int = compr_21_
                        if ((-217) <= (d_56_v5_)) and ((d_56_v5_) < (-953)):
                            coll21_ = coll21_.union(_dafny.Set([(d_56_v5_) * (len(_dafny.Map({False: 11})))]))
                    return _dafny.Set(coll21_)
                coll19_ = _dafny.Map()
                def iife20_():
                    coll20_ = _dafny.Set()
                    compr_20_: int
                    for compr_20_ in _dafny.IntegerRange(-217, -953):
                        d_54_v5_: int = compr_20_
                        if ((-217) <= (d_54_v5_)) and ((d_54_v5_) < (-953)):
                            coll20_ = coll20_.union(_dafny.Set([(d_54_v5_) * (len(_dafny.Map({False: 11})))]))
                    return _dafny.Set(coll20_)
                compr_19_: int
                for compr_19_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(iife20_()
                )})).keys.Elements:
                    d_55_v4_: int = compr_19_
                    if (d_55_v4_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(iife21_()
                    )})):
                        coll19_[default__.safeDivisionInt(d_55_v4_, 650)] = _dafny.CodePoint('x')
                return _dafny.Map(coll19_)
            def iife22_():
                coll22_ = _dafny.Map()
                compr_22_: int
                for compr_22_ in _dafny.IntegerRange(614, 189):
                    d_57_v6_: int = compr_22_
                    if ((614) <= (d_57_v6_)) and ((d_57_v6_) < (189)):
                        coll22_[default__.safeDivisionInt(d_57_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvfmnoywj"))))] = _dafny.CodePoint('f')
                return _dafny.Map(coll22_)
            compr_15_: _dafny.Map
            for compr_15_ in ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_48_i0_ in range(default__.abs(940))])): _dafny.CodePoint('o')}), _dafny.Map({len(iife16_()
            ): _dafny.CodePoint('h')}), iife17_()
            , iife18_()
            ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_53_i2_ in range(default__.abs(69))])): _dafny.CodePoint('o')}), iife19_()
            , _dafny.Map({379: _dafny.CodePoint('f')}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egagps")))): _dafny.CodePoint('c')}), iife22_()
            ]))).Elements:
                d_58_v0_: _dafny.Map = compr_15_
                if (d_58_v0_) in ((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_48_i0_ in range(default__.abs(940))])): _dafny.CodePoint('o')}), _dafny.Map({len(iife23_()
                ): _dafny.CodePoint('h')}), iife24_()
                , iife25_()
                ])) + (_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_53_i2_ in range(default__.abs(69))])): _dafny.CodePoint('o')}), iife26_()
                , _dafny.Map({379: _dafny.CodePoint('f')}), _dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "egagps")))): _dafny.CodePoint('c')}), iife29_()
                ]))):
                    coll15_[d_58_v0_] = not(True)
            return _dafny.Map(coll15_)
        return iife15_()
        

    @staticmethod
    def fm40(p0, globalState):
        def iife30_():
            coll30_ = _dafny.Map()
            compr_30_: _dafny.Seq
            for compr_30_ in ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgeav")): _dafny.Map({D1_DC5(not(False), True, 504, True, True): not(False)})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")): _dafny.Map({D1_DC5(not(False), True, -294, not(False), False): True})}))).keys.Elements:
                d_61_v0_: _dafny.Seq = compr_30_
                if (d_61_v0_) in ((_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qgeav")): _dafny.Map({D1_DC5(not(False), True, 504, True, True): not(False)})})) | (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e")): _dafny.Map({D1_DC5(not(False), True, -294, not(False), False): True})}))):
                    coll30_[d_61_v0_] = (0) - (((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality) - (449))
            return _dafny.Map(coll30_)
        return iife30_()
        

    @staticmethod
    def fm41(p0, globalState):
        def iife31_():
            coll31_ = _dafny.Set()
            compr_31_: str
            for compr_31_ in (_dafny.Map({_dafny.CodePoint('e'): 133})).keys.Elements:
                d_62_v0_: str = compr_31_
                if (d_62_v0_) in (_dafny.Map({_dafny.CodePoint('e'): 133})):
                    coll31_ = coll31_.union(_dafny.Set([d_62_v0_]))
            return _dafny.Set(coll31_)
        return (_dafny.Set({_dafny.CodePoint('i')})) - ((iife31_()
        ) - (_dafny.Set({_dafny.CodePoint('j')})))

    @staticmethod
    def fm42(globalState):
        return _dafny.Map({(len(_dafny.SeqWithoutIsStrInference([122]))) + ((0) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "l")))])))): D1_DC5(True, not(False), (_dafny.MultiSet([len(_dafny.Map({True: len(_dafny.Set({len(_dafny.Set({False})), -262, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gifktq"))), -138, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxqco")))}))}))])).cardinality, True, False)})

    @staticmethod
    def fm43(globalState):
        return D0_DC1((len(_dafny.Map({not(True): (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, False]))).cardinality}))) - ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fqgrojepe")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlsvtg"))])).cardinality))

    @staticmethod
    def fm44(p0, p1, p2, p3, globalState):
        source2_ = D0_DC2(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_63_i0_ in range(default__.abs(314))])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p"))))
        if source2_.is_DC1:
            d_64___mcc_h0_ = source2_.cf1
            d_65_cf1_ = d_64___mcc_h0_
            return _dafny.MultiSet([_dafny.Set({True, True}), _dafny.Set({False, not(True)}), _dafny.Set({not(False), not(True)}), _dafny.Set({True})])
        elif source2_.is_DC2:
            d_66___mcc_h1_ = source2_.cf2
            d_67___mcc_h2_ = source2_.cf3
            d_68_cf3_ = d_67___mcc_h2_
            d_69_cf2_ = d_66___mcc_h1_
            return _dafny.MultiSet([_dafny.Set({not(True)})])
        elif True:
            d_70___mcc_h3_ = source2_.cf0
            d_71_cf0_ = d_70___mcc_h3_
            return _dafny.MultiSet([_dafny.Set({False, False}), _dafny.Set({False})])

    @staticmethod
    def fm45(p0, globalState):
        return D22_DC56(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohkasfpnp")), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmtj"))), (D7_DC18(True)).cf29, (len(_dafny.SeqWithoutIsStrInference([False, True]))) >= ((0) - (len(_dafny.SeqWithoutIsStrInference([True, not(True), False, False])))))

    @staticmethod
    def fm46(p0, globalState):
        return ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ynim"))})) | (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fj"))}))) | ((_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aw"))})) | (_dafny.Map({not(True): _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agjpjwmqp"))})))

    @staticmethod
    def fm47(p0, globalState):
        return _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: -9})), len(_dafny.Set({not(True), True})), 959, 207, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kwuvakx")))]) if True else _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xgqjjacir")))])), _dafny.SeqWithoutIsStrInference([147])])

    @staticmethod
    def fm48(p0, p1, globalState):
        return ((_dafny.Map({False: True})) | (_dafny.Map({False: not(True)}))) | ((_dafny.Map({True: True})) | (_dafny.Map({not(False): True})))

    @staticmethod
    def fm49(p0, p1, p2, p3, globalState):
        source3_ = D22_DC57(_dafny.CodePoint('q'), len(_dafny.Map({469: 233})), len(_dafny.Map({True: -799})), not(True))
        if source3_.is_DC56:
            d_72___mcc_h0_ = source3_.cf105
            d_73___mcc_h1_ = source3_.cf106
            d_74___mcc_h2_ = source3_.cf107
            d_75___mcc_h3_ = source3_.cf108
            d_76_cf108_ = d_75___mcc_h3_
            d_77_cf107_ = d_74___mcc_h2_
            d_78_cf106_ = d_73___mcc_h1_
            d_79_cf105_ = d_72___mcc_h0_
            return (_dafny.SeqWithoutIsStrInference([D12_DC30(d_76_cf108_, (D22_DC57(_dafny.CodePoint('a'), len(d_79_cf105_), d_78_cf106_, d_76_cf108_)).cf112, d_77_cf107_, d_78_cf106_, True), D12_DC30(d_77_cf107_, d_76_cf108_, d_76_cf108_, d_78_cf106_, d_77_cf107_)])) + (_dafny.SeqWithoutIsStrInference([D12_DC30(d_77_cf107_, d_76_cf108_, True, d_78_cf106_, d_77_cf107_)]))
        elif source3_.is_DC57:
            d_80___mcc_h4_ = source3_.cf109
            d_81___mcc_h5_ = source3_.cf110
            d_82___mcc_h6_ = source3_.cf111
            d_83___mcc_h7_ = source3_.cf112
            d_84_cf112_ = d_83___mcc_h7_
            d_85_cf111_ = d_82___mcc_h6_
            d_86_cf110_ = d_81___mcc_h5_
            d_87_cf109_ = d_80___mcc_h4_
            return _dafny.SeqWithoutIsStrInference([D12_DC30(d_84_cf112_, d_84_cf112_, True, d_85_cf111_, not(d_84_cf112_)) for d_88_i0_ in range(default__.abs(-892))])
        elif source3_.is_DC55:
            d_89___mcc_h8_ = source3_.cf104
            d_90_cf104_ = d_89___mcc_h8_
            return (_dafny.SeqWithoutIsStrInference([D12_DC30(False, False, False, len(_dafny.Set({(_dafny.MultiSet([len(_dafny.Map({len(_dafny.Map({False: 298})): len(_dafny.SeqWithoutIsStrInference([724]))})), 750, 962, len(_dafny.Map({_dafny.CodePoint('j'): not(True)})), len(_dafny.SeqWithoutIsStrInference([True, True, False, False, False]))])).cardinality})), True)])) + (_dafny.SeqWithoutIsStrInference([D12_DC30(False, not(False), True, len(_dafny.SeqWithoutIsStrInference([True, False])), False)]))
        elif True:
            d_91___mcc_h9_ = source3_.cf113
            d_92_cf113_ = d_91___mcc_h9_
            return (_dafny.SeqWithoutIsStrInference([D12_DC30(False, True, False, 836, True), D12_DC30(not(False), False, False, len(_dafny.SeqWithoutIsStrInference([True, True])), not(True)), D12_DC30(True, True, False, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xegqw"))), False)])) + (_dafny.SeqWithoutIsStrInference([D12_DC30(True, False, True, 321, False), D12_DC30(False, False, not(False), 216, True)]))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.Set({True, True, False})) | (_dafny.Set({False, True, True})), _dafny.Set({True, False}), (_dafny.Set({False})) - (_dafny.Set({False, True}))])

    @staticmethod
    def fm51(globalState):
        def iife32_():
            coll32_ = _dafny.Map()
            compr_32_: _dafny.Map
            for compr_32_ in (_dafny.SeqWithoutIsStrInference([_dafny.Map({32: not(True)}), _dafny.Map({-925: not(True)})])).Elements:
                d_93_v0_: _dafny.Map = compr_32_
                if (d_93_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.Map({32: not(True)}), _dafny.Map({-925: not(True)})])):
                    coll32_[d_93_v0_] = False
            return _dafny.Map(coll32_)
        def iife33_():
            def iife35_():
                coll35_ = _dafny.Set()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(-806, 531):
                    d_96_v2_: int = compr_35_
                    if ((-806) <= (d_96_v2_)) and ((d_96_v2_) < (531)):
                        coll35_ = coll35_.union(_dafny.Set([default__.safeModuloInt(d_96_v2_, len(_dafny.SeqWithoutIsStrInference([True, False, True])))]))
                return _dafny.Set(coll35_)
            coll33_ = _dafny.Map()
            def iife34_():
                coll34_ = _dafny.Set()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(-806, 531):
                    d_94_v2_: int = compr_34_
                    if ((-806) <= (d_94_v2_)) and ((d_94_v2_) < (531)):
                        coll34_ = coll34_.union(_dafny.Set([default__.safeModuloInt(d_94_v2_, len(_dafny.SeqWithoutIsStrInference([True, False, True])))]))
                return _dafny.Set(coll34_)
            compr_33_: int
            for compr_33_ in (iife34_()
            ).Elements:
                d_95_v1_: int = compr_33_
                if (d_95_v1_) in (iife35_()
                ):
                    coll33_[(d_95_v1_) + (len(_dafny.SeqWithoutIsStrInference([True])))] = not(False)
            return _dafny.Map(coll33_)
        def iife36_():
            coll36_ = _dafny.Map()
            compr_36_: _dafny.Map
            for compr_36_ in (_dafny.MultiSet([_dafny.Map({-700: False}), _dafny.Map({(0) - (-190): True})])).Elements:
                d_98_v3_: _dafny.Map = compr_36_
                if (d_98_v3_) in (_dafny.MultiSet([_dafny.Map({-700: False}), _dafny.Map({(0) - (-190): True})])):
                    coll36_[d_98_v3_] = False
            return _dafny.Map(coll36_)
        return ((iife32_()
        ) | (_dafny.Map({iife33_()
        : True}))) | ((_dafny.Map({_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_97_i0_ in range(default__.abs(-753))])): False}): not(True)})) | (iife36_()
        ))

    @staticmethod
    def fm52(p0, p1, p2, p3, globalState):
        return _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wdunh")), ((D16_DC37(_dafny.SeqWithoutIsStrInference([not(True)]), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o")), _dafny.CodePoint('d'))).cf66) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cvteawr"))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jtjh"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "t"))), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "krx"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fpg")))])

    @staticmethod
    def fm53(p0, p1, p2, p3, globalState):
        def iife37_():
            coll37_ = _dafny.Set()
            compr_37_: _dafny.Seq
            for compr_37_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([623]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([-901])).cardinality]), _dafny.SeqWithoutIsStrInference([947 for d_102_i3_ in range(default__.abs(251))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])])).Elements:
                d_103_v0_: _dafny.Seq = compr_37_
                if (d_103_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([623]), _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([-901])).cardinality]), _dafny.SeqWithoutIsStrInference([947 for d_102_i3_ in range(default__.abs(251))]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({False}))])])):
                    coll37_ = coll37_.union(_dafny.Set([d_103_v0_]))
            return _dafny.Set(coll37_)
        return (_dafny.Set({_dafny.SeqWithoutIsStrInference([450, len(_dafny.Map({not(False): not(False)})), len(_dafny.SeqWithoutIsStrInference([True, True])), 924]), _dafny.SeqWithoutIsStrInference([831 for d_99_i0_ in range(default__.abs(842))]), _dafny.SeqWithoutIsStrInference([15, len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oxlytiaam"))), len(_dafny.Map({True: False}))])) for d_100_i1_ in range(default__.abs(148))])), -438])})).intersection((_dafny.Set({_dafny.SeqWithoutIsStrInference([-832]), _dafny.SeqWithoutIsStrInference([-197 for d_101_i2_ in range(default__.abs(-109))])})) - (iife37_()
        ))

    @staticmethod
    def fm54(globalState):
        return D6_DC16((_dafny.Set({(D7_DC18(not(True))).cf29})).isdisjoint(_dafny.Set({not(True)})), (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygv"))) < (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_104_i0_ in range(default__.abs(213))])))

    @staticmethod
    def m0(p0, p1, p2, p3, globalState):
        d_105_v0_: bool
        d_105_v0_ = False
        d_106_v1_: D8
        d_106_v1_ = D8_DC21(d_105_v0_, d_105_v0_, (0) - (p3))
        d_107_v2_: _dafny.Map
        d_107_v2_ = _dafny.Map({(d_106_v1_).cf34: d_105_v0_})
        index0_ = default__.safeIndex(666, (p0).length(0))
        (p0)[index0_] = (d_105_v0_ if default__.fm4(globalState) else d_105_v0_)
        d_108_v3_: _dafny.Array
        nw0_ = _dafny.Array(_dafny.Array(None, 0), 11)
        d_108_v3_ = nw0_
        d_109_v4_: D1
        d_109_v4_ = D1_DC5(d_105_v0_, d_105_v0_, p3, d_105_v0_, d_105_v0_)
        d_110_v5_: _dafny.Seq
        d_110_v5_ = _dafny.SeqWithoutIsStrInference([d_109_v4_, d_109_v4_])
        d_111_v6_: _dafny.Seq
        d_111_v6_ = _dafny.SeqWithoutIsStrInference([len(d_110_v5_)])
        d_112_v7_: _dafny.Seq
        d_112_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bjqgtso"))
        d_113_v8_: D22
        d_113_v8_ = D22_DC56(d_112_v7_, p3, d_105_v0_, d_105_v0_)
        d_114_v9_: D18
        d_114_v9_ = D18_DC41(d_111_v6_)
        d_115_v10_: _dafny.Seq
        d_115_v10_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p3 for d_116_i1_ in range(default__.abs(-891))]), d_111_v6_, d_111_v6_, ((d_114_v9_).cf71).set(default__.safeIndex(p2, len((d_114_v9_).cf71)), p3), d_111_v6_])
        d_117_v11_: _dafny.Array
        nw1_ = _dafny.Array(None, 19)
        nw1_[int(0)] = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ummi"))) for d_118_i0_ in range(default__.abs(544))])
        nw1_[int(1)] = d_111_v6_
        nw1_[int(2)] = default__.fm6(d_105_v0_, p2, d_112_v7_, globalState)
        nw1_[int(3)] = _dafny.SeqWithoutIsStrInference([p3, 870, (d_113_v8_).cf106])
        nw1_[int(4)] = (d_115_v10_)[default__.safeIndex(p3, len(d_115_v10_))]
        nw1_[int(5)] = d_111_v6_
        nw1_[int(6)] = d_111_v6_
        nw1_[int(7)] = d_111_v6_
        nw1_[int(8)] = d_111_v6_
        nw1_[int(9)] = d_111_v6_
        nw1_[int(10)] = d_111_v6_
        nw1_[int(11)] = d_111_v6_
        nw1_[int(12)] = _dafny.SeqWithoutIsStrInference([p2 for d_119_i2_ in range(default__.abs(985))])
        nw1_[int(13)] = d_111_v6_
        nw1_[int(14)] = d_111_v6_
        nw1_[int(15)] = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_105_v0_, (D20_DC49(_dafny.CodePoint('n'), p3, d_105_v0_, ((d_107_v2_)[d_105_v0_] if (d_105_v0_) in (d_107_v2_) else False), d_105_v0_)).cf92})) for d_120_i3_ in range(default__.abs(744))])
        nw1_[int(16)] = d_111_v6_
        nw1_[int(17)] = d_111_v6_
        nw1_[int(18)] = d_111_v6_
        d_117_v11_ = nw1_
        index1_ = default__.safeIndex(651, (d_108_v3_).length(0))
        (d_108_v3_)[index1_] = d_117_v11_
        index2_ = default__.safeIndex(666, (p0).length(0))
        index3_ = default__.safeIndex(651, (d_108_v3_).length(0))
        rhs0_ = d_107_v2_
        rhs1_ = d_105_v0_
        rhs2_ = (p2) <= (p3)
        rhs3_ = d_117_v11_
        rhs4_ = (p2) != ((365 if default__.fm4(globalState) else p2))
        lhs0_ = globalState
        lhs1_ = p0
        lhs2_ = default__.safeIndex(666, (p0).length(0))
        lhs3_ = d_108_v3_
        lhs4_ = default__.safeIndex(651, (d_108_v3_).length(0))
        lhs5_ = globalState
        d_107_v2_ = rhs0_
        lhs0_.f11 = rhs1_
        lhs1_[lhs2_] = rhs2_
        lhs3_[lhs4_] = rhs3_
        lhs5_.f11 = rhs4_
        d_105_v0_ = (D8_DC21((p0)[default__.safeIndex(666, (p0).length(0))], d_105_v0_, p3)).cf33
        (globalState).f14 = p2
        (globalState).f0 = (0) - (p3)
        d_121_v12_: _dafny.Array
        nw2_ = _dafny.Array(int(0), 15)
        d_121_v12_ = nw2_
        d_121_v12_ = d_121_v12_
        d_122_i4_: int
        d_122_i4_ = 0
        with _dafny.label("0"):
            while not(True):
                with _dafny.c_label("0"):
                    if (d_122_i4_) >= (100):
                        raise _dafny.Break("0")
                    d_122_i4_ = (d_122_i4_) + (1)
                    (globalState).f0 = len((d_112_v7_) + (d_112_v7_))
                    d_123_v13_: str
                    d_123_v13_ = _dafny.CodePoint('n')
                    (globalState).f11 = (d_123_v13_) in (d_112_v7_)
                    index4_ = default__.safeIndex(666, (p0).length(0))
                    rhs5_ = (d_114_v9_).cf71
                    rhs6_ = d_105_v0_
                    lhs6_ = p0
                    lhs7_ = default__.safeIndex(666, (p0).length(0))
                    d_111_v6_ = rhs5_
                    lhs6_[lhs7_] = rhs6_
                    (globalState).f1 = d_105_v0_
                    pass
            pass

    @staticmethod
    def Main(noArgsParameter__):
        d_124_v0_: _dafny.Array
        def lambda0_(d_125_i0_):
            return _dafny.CodePoint('o')

        init0_ = lambda0_
        nw3_ = _dafny.Array(None, 24)
        for i0_0_ in range(nw3_.length(0)):
            nw3_[i0_0_] = init0_(i0_0_)
        d_124_v0_ = nw3_
        d_126_v1_: int
        d_126_v1_ = 554
        d_127_v2_: _dafny.Seq
        d_127_v2_ = _dafny.SeqWithoutIsStrInference([d_126_v1_, d_126_v1_])
        d_128_v3_: bool
        d_128_v3_ = True
        d_129_v4_: _dafny.Seq
        d_129_v4_ = _dafny.SeqWithoutIsStrInference([d_128_v3_])
        d_130_v5_: _dafny.Map
        d_130_v5_ = _dafny.Map({d_126_v1_: len(_dafny.Map({len(d_129_v4_): d_128_v3_}))})
        d_131_v6_: _dafny.MultiSet
        d_131_v6_ = _dafny.MultiSet([d_126_v1_])
        d_132_v7_: D0
        d_132_v7_ = D0_DC0((d_130_v5_).set(d_126_v1_, (d_131_v6_).cardinality))
        d_133_v8_: _dafny.Map
        d_133_v8_ = _dafny.Map({d_128_v3_: d_128_v3_})
        d_134_v9_: _dafny.Array
        nw4_ = _dafny.Array(_dafny.MultiSet({}), 3)
        d_134_v9_ = nw4_
        d_135_v10_: _dafny.Array
        nw5_ = _dafny.Array(int(0), 11)
        d_135_v10_ = nw5_
        d_136_v11_: _dafny.Map
        d_136_v11_ = _dafny.Map({d_126_v1_: d_128_v3_})
        d_137_v12_: _dafny.Seq
        d_137_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgkmsxlai"))
        d_138_globalState_: GlobalState
        nw6_ = GlobalState()
        nw6_.ctor__(734, False, d_124_v0_, False, 454, True, 354, False, True, d_127_v2_, (d_132_v7_).cf0, False, _dafny.SeqWithoutIsStrInference([((d_133_v8_)[d_128_v3_] if (d_128_v3_) in (d_133_v8_) else False), False]), False, 935, d_134_v9_, d_129_v4_, True, d_135_v10_, 511, d_136_v11_, d_131_v6_, 257, -900, 539, (d_137_v12_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lv"))), True, True)
        d_138_globalState_ = nw6_
        d_139_v13_: _dafny.Array
        nw7_ = _dafny.Array(None, 25)
        nw7_[int(0)] = (_dafny.SeqWithoutIsStrInference([d_128_v3_, d_128_v3_])) + (d_129_v4_)
        nw7_[int(1)] = d_129_v4_
        nw7_[int(2)] = (default__.fm0(not(d_128_v3_), d_138_globalState_)) + (d_129_v4_)
        nw7_[int(3)] = d_129_v4_
        nw7_[int(4)] = d_129_v4_
        nw7_[int(5)] = d_129_v4_
        nw7_[int(6)] = _dafny.SeqWithoutIsStrInference([d_128_v3_])
        nw7_[int(7)] = d_129_v4_
        nw7_[int(8)] = d_129_v4_
        nw7_[int(9)] = d_129_v4_
        nw7_[int(10)] = d_129_v4_
        nw7_[int(11)] = (d_129_v4_) + (d_129_v4_)
        nw7_[int(12)] = (d_129_v4_) + (d_129_v4_)
        nw7_[int(13)] = _dafny.SeqWithoutIsStrInference([d_128_v3_, d_128_v3_, d_128_v3_, d_128_v3_])
        nw7_[int(14)] = d_129_v4_
        nw7_[int(15)] = d_129_v4_
        nw7_[int(16)] = (d_129_v4_).set(default__.safeIndex(d_126_v1_, len(d_129_v4_)), d_128_v3_)
        nw7_[int(17)] = d_129_v4_
        nw7_[int(18)] = _dafny.SeqWithoutIsStrInference([d_128_v3_])
        nw7_[int(19)] = d_129_v4_
        nw7_[int(20)] = d_129_v4_
        nw7_[int(21)] = d_129_v4_
        nw7_[int(22)] = (d_129_v4_) + (default__.fm0(not(d_128_v3_), d_138_globalState_))
        nw7_[int(23)] = d_129_v4_
        nw7_[int(24)] = (d_129_v4_) + (d_129_v4_)
        d_139_v13_ = nw7_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_139_v13_).length(0)):
            d_140_i1_: int = guard_loop_0_
            if (True) and (((0) <= (d_140_i1_)) and ((d_140_i1_) < ((d_139_v13_).length(0)))):
                (d_139_v13_)[(d_140_i1_)] = d_129_v4_
        index5_ = default__.safeIndex(223, (d_135_v10_).length(0))
        (d_135_v10_)[index5_] = d_126_v1_
        index6_ = default__.safeIndex(223, (d_135_v10_).length(0))
        (d_135_v10_)[index6_] = d_126_v1_
        (d_138_globalState_).f27 = d_128_v3_
        d_141_v14_: _dafny.Array
        def lambda1_(d_142_v1_, d_143_v3_, d_144_v10_):
            def lambda2_(d_145_i2_):
                return (_dafny.Set({_dafny.Map({(D1_DC4(d_142_v1_, d_143_v3_)).cf6: (d_144_v10_)[default__.safeIndex(223, (d_144_v10_).length(0))]}), _dafny.Map({d_143_v3_: d_142_v1_})})) | (_dafny.Set({_dafny.Map({not(d_143_v3_): d_142_v1_})}))

            return lambda2_

        init1_ = lambda1_(d_126_v1_, d_128_v3_, d_135_v10_)
        nw8_ = _dafny.Array(None, 21)
        for i0_1_ in range(nw8_.length(0)):
            nw8_[i0_1_] = init1_(i0_1_)
        d_141_v14_ = nw8_
        d_146_v15_: _dafny.Map
        d_146_v15_ = _dafny.Map({d_128_v3_: (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]})
        d_147_v16_: _dafny.Set
        d_147_v16_ = _dafny.Set({d_146_v15_, d_146_v15_})
        index7_ = default__.safeIndex(925, (d_141_v14_).length(0))
        (d_141_v14_)[index7_] = d_147_v16_
        index8_ = default__.safeIndex(925, (d_141_v14_).length(0))
        (d_141_v14_)[index8_] = default__.fm1(len(d_129_v4_), _dafny.Set({d_128_v3_, d_128_v3_, False, d_128_v3_, d_128_v3_}), d_138_globalState_)
        d_148_v17_: _dafny.Set
        d_148_v17_ = _dafny.Set({d_128_v3_, d_128_v3_, d_128_v3_})
        d_149_v18_: _dafny.Map
        d_149_v18_ = _dafny.Map({d_128_v3_: (d_137_v12_)[default__.safeIndex((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], len(d_137_v12_))]})
        (d_138_globalState_).f24 = default__.fm2(default__.fm3(d_138_globalState_), d_148_v17_, d_149_v18_, (d_126_v1_) <= ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]), d_138_globalState_)
        d_150_v19_: _dafny.Map
        d_150_v19_ = _dafny.Map({d_127_v2_: d_128_v3_})
        d_151_v20_: _dafny.Array
        nw9_ = _dafny.Array(None, 17)
        nw9_[int(0)] = True
        nw9_[int(1)] = d_128_v3_
        nw9_[int(2)] = default__.fm4(d_138_globalState_)
        nw9_[int(3)] = d_128_v3_
        nw9_[int(4)] = d_128_v3_
        nw9_[int(5)] = default__.fm4(d_138_globalState_)
        nw9_[int(6)] = d_128_v3_
        nw9_[int(7)] = d_128_v3_
        nw9_[int(8)] = False
        nw9_[int(9)] = (228) < (len(d_150_v19_))
        nw9_[int(10)] = True
        nw9_[int(11)] = not(d_128_v3_)
        nw9_[int(12)] = d_128_v3_
        nw9_[int(13)] = d_128_v3_
        nw9_[int(14)] = d_128_v3_
        nw9_[int(15)] = d_128_v3_
        nw9_[int(16)] = d_128_v3_
        d_151_v20_ = nw9_
        d_151_v20_ = d_151_v20_
        d_152_v21_: _dafny.Seq
        d_152_v21_ = _dafny.SeqWithoutIsStrInference([d_151_v20_, d_151_v20_, (D2_DC6(d_151_v20_)).cf12, d_151_v20_, d_151_v20_])
        d_151_v20_ = (d_152_v21_)[default__.safeIndex(default__.safeDivisionInt((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_126_v1_), len(d_152_v21_))]
        d_153_v22_: _dafny.Seq
        d_153_v22_ = _dafny.SeqWithoutIsStrInference([d_135_v10_, d_135_v10_])
        d_154_i3_: int
        d_154_i3_ = 0
        with _dafny.label("1"):
            while ((d_135_v10_) in (d_153_v22_) if d_128_v3_ else d_128_v3_):
                with _dafny.c_label("1"):
                    if (d_154_i3_) >= (100):
                        raise _dafny.Break("1")
                    d_154_i3_ = (d_154_i3_) + (1)
                    (d_138_globalState_).f24 = (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]
                    d_151_v20_ = d_151_v20_
                    default__.m0(d_151_v20_, d_135_v10_, 898, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_138_globalState_)
                    d_155_v23_: _dafny.Array
                    nw10_ = _dafny.Array(_dafny.Map({}), 16)
                    d_155_v23_ = nw10_
                    d_156_v24_: _dafny.Array
                    d_156_v24_ = d_155_v23_
                    d_157_v25_: _dafny.Array
                    nw11_ = _dafny.Array(None, 23)
                    nw11_[int(0)] = d_155_v23_
                    nw11_[int(1)] = d_155_v23_
                    nw11_[int(2)] = d_155_v23_
                    nw11_[int(3)] = d_155_v23_
                    nw11_[int(4)] = (d_155_v23_ if d_128_v3_ else d_155_v23_)
                    nw11_[int(5)] = (d_155_v23_ if not(d_128_v3_) else d_155_v23_)
                    nw11_[int(6)] = d_155_v23_
                    nw11_[int(7)] = d_155_v23_
                    nw11_[int(8)] = d_155_v23_
                    nw11_[int(9)] = d_155_v23_
                    nw11_[int(10)] = (d_155_v23_ if True else d_155_v23_)
                    nw11_[int(11)] = d_155_v23_
                    nw11_[int(12)] = d_155_v23_
                    nw11_[int(13)] = d_155_v23_
                    nw11_[int(14)] = d_155_v23_
                    nw11_[int(15)] = (d_156_v24_)
                    nw11_[int(16)] = d_155_v23_
                    nw11_[int(17)] = d_155_v23_
                    nw11_[int(18)] = d_155_v23_
                    nw11_[int(19)] = d_155_v23_
                    nw11_[int(20)] = d_155_v23_
                    nw11_[int(21)] = d_155_v23_
                    nw11_[int(22)] = d_155_v23_
                    d_157_v25_ = nw11_
                    index9_ = default__.safeIndex(528, (d_157_v25_).length(0))
                    (d_157_v25_)[index9_] = d_155_v23_
                    index10_ = default__.safeIndex(528, (d_157_v25_).length(0))
                    (d_157_v25_)[index10_] = d_155_v23_
                    pass
            pass
        d_158_v26_: str
        d_158_v26_ = _dafny.CodePoint('j')
        d_159_v27_: _dafny.Map
        d_159_v27_ = _dafny.Map({d_128_v3_: default__.fm3(d_138_globalState_)})
        d_160_v28_: _dafny.Array
        nw12_ = _dafny.Array(None, 10)
        nw12_[int(0)] = (_dafny.SeqWithoutIsStrInference([(d_137_v12_)[default__.safeIndex((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], len(d_137_v12_))] for d_161_i4_ in range(default__.abs(368))])) + (d_137_v12_)
        nw12_[int(1)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wsoc"))
        nw12_[int(2)] = (d_137_v12_).set(default__.safeIndex((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], len(d_137_v12_)), d_158_v26_)
        nw12_[int(3)] = _dafny.SeqWithoutIsStrInference([d_158_v26_ for d_162_i5_ in range(default__.abs(944))])
        nw12_[int(4)] = (d_137_v12_) + (d_137_v12_)
        nw12_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbo"))
        nw12_[int(6)] = d_137_v12_
        nw12_[int(7)] = d_137_v12_
        nw12_[int(8)] = d_137_v12_
        nw12_[int(9)] = ((d_159_v27_)[d_128_v3_] if (d_128_v3_) in (d_159_v27_) else d_137_v12_)
        d_160_v28_ = nw12_
        index11_ = default__.safeIndex(54, (d_160_v28_).length(0))
        (d_160_v28_)[index11_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
        index12_ = default__.safeIndex(54, (d_160_v28_).length(0))
        (d_160_v28_)[index12_] = ((d_159_v27_)[True] if (True) in (d_159_v27_) else d_137_v12_)
        index13_ = default__.safeIndex(54, (d_160_v28_).length(0))
        (d_160_v28_)[index13_] = ((d_137_v12_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obandijdf"))) + (d_137_v12_))).set(default__.safeIndex((0) - (d_126_v1_), len((d_137_v12_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "obandijdf"))) + (d_137_v12_)))), default__.fm5((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_128_v3_, d_138_globalState_))
        d_163_i6_: int
        d_163_i6_ = 0
        with _dafny.label("2"):
            while d_128_v3_:
                with _dafny.c_label("2"):
                    if (d_163_i6_) >= (100):
                        raise _dafny.Break("2")
                    d_163_i6_ = (d_163_i6_) + (1)
                    d_164_v29_: _dafny.Seq
                    d_164_v29_ = _dafny.SeqWithoutIsStrInference([d_130_v5_])
                    pat_let_tv0_ = d_130_v5_
                    d_165_v30_: _dafny.Array
                    nw13_ = _dafny.Array(None, 6)
                    nw13_[int(0)] = D0_DC0(d_130_v5_)
                    nw13_[int(1)] = D0_DC0((d_164_v29_)[default__.safeIndex(d_126_v1_, len(d_164_v29_))])
                    nw13_[int(2)] = d_132_v7_
                    def iife38_(_pat_let0_0):
                        def iife39_(d_166_dt__update__tmp_h0_):
                            def iife40_(_pat_let1_0):
                                def iife41_(d_167_dt__update_hcf0_h0_):
                                    return D0_DC0(d_167_dt__update_hcf0_h0_)
                                return iife41_(_pat_let1_0)
                            return iife40_(pat_let_tv0_)
                        return iife39_(_pat_let0_0)
                    nw13_[int(3)] = iife38_(d_132_v7_)
                    nw13_[int(4)] = d_132_v7_
                    nw13_[int(5)] = d_132_v7_
                    d_165_v30_ = nw13_
                    d_168_v31_: _dafny.Map
                    d_168_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_165_v30_]): d_131_v6_})
                    d_169_v32_: _dafny.Seq
                    d_169_v32_ = _dafny.SeqWithoutIsStrInference([d_165_v30_, d_165_v30_, d_165_v30_])
                    d_168_v31_ = (d_168_v31_).set(d_169_v32_, d_131_v6_)
                    d_170_v33_: D0
                    d_170_v33_ = D0_DC2((0) - ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]), d_126_v1_)
                    def iife42_(_pat_let2_0):
                        def iife43_(d_171_dt__update__tmp_h1_):
                            def iife44_(_pat_let3_0):
                                def iife45_(d_172_dt__update_hcf2_h0_):
                                    return D0_DC2(d_172_dt__update_hcf2_h0_, (d_171_dt__update__tmp_h1_).cf3)
                                return iife45_(_pat_let3_0)
                            return iife44_(220)
                        return iife43_(_pat_let2_0)
                    source4_ = iife42_(d_170_v33_)
                    if source4_.is_DC1:
                        d_173___mcc_h0_ = source4_.cf1
                        d_174_cf1_ = d_173___mcc_h0_
                        d_174_cf1_ = default__.safeDivisionInt(default__.fm2((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))], d_148_v17_, d_149_v18_, d_128_v3_, d_138_globalState_), default__.safeDivisionInt(d_174_cf1_, d_174_cf1_))
                        d_175_v34_: _dafny.Map
                        d_175_v34_ = _dafny.Map({d_151_v20_: (_dafny.SeqWithoutIsStrInference([d_128_v3_])) + ((d_129_v4_).set(default__.safeIndex((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], len(d_129_v4_)), d_128_v3_))})
                        d_175_v34_ = (d_175_v34_).set(d_151_v20_, d_129_v4_)
                        (d_138_globalState_).f27 = False
                        d_176_v35_: D0
                        d_176_v35_ = D0_DC1(44)
                        d_177_v36_: _dafny.Map
                        d_177_v36_ = _dafny.Map({d_151_v20_: d_176_v35_})
                        d_177_v36_ = (d_177_v36_).set(d_151_v20_, d_176_v35_)
                    elif source4_.is_DC2:
                        d_178___mcc_h1_ = source4_.cf2
                        d_179___mcc_h2_ = source4_.cf3
                        d_180_cf3_ = d_179___mcc_h2_
                        d_181_cf2_ = d_178___mcc_h1_
                        d_182_v37_: _dafny.Array
                        nw14_ = _dafny.Array(D1.default()(), 14)
                        d_182_v37_ = nw14_
                        d_183_v38_: _dafny.MultiSet
                        d_183_v38_ = _dafny.MultiSet([d_158_v26_, d_158_v26_, d_158_v26_])
                        d_184_v39_: D1
                        d_184_v39_ = D1_DC5(d_128_v3_, d_128_v3_, (0) - (d_126_v1_), ((d_136_v11_)[((d_183_v38_)[d_158_v26_] if (d_158_v26_) in (d_183_v38_) else d_181_cf2_)] if (((d_183_v38_)[d_158_v26_] if (d_158_v26_) in (d_183_v38_) else d_181_cf2_)) in (d_136_v11_) else False), d_128_v3_)
                        index14_ = default__.safeIndex(113, (d_182_v37_).length(0))
                        (d_182_v37_)[index14_] = d_184_v39_
                        index15_ = default__.safeIndex(113, (d_182_v37_).length(0))
                        (d_182_v37_)[index15_] = D1_DC5((d_180_cf3_) > (d_126_v1_), d_128_v3_, d_180_cf3_, d_128_v3_, d_128_v3_)
                        d_146_v15_ = (d_146_v15_).set((d_128_v3_ if d_128_v3_ else d_128_v3_), d_180_cf3_)
                        default__.m0(d_151_v20_, d_135_v10_, (d_184_v39_).cf9, len(_dafny.SeqWithoutIsStrInference([d_128_v3_, d_128_v3_])), d_138_globalState_)
                        d_185_v40_: _dafny.Seq
                        d_185_v40_ = _dafny.SeqWithoutIsStrInference([d_148_v17_, _dafny.Set({False})])
                        d_186_v41_: _dafny.Map
                        d_186_v41_ = _dafny.Map({681: d_149_v18_})
                        d_187_v42_: _dafny.Set
                        d_187_v42_ = _dafny.Set({d_180_cf3_})
                        default__.m0(d_151_v20_, d_135_v10_, default__.fm2((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))], (d_185_v40_)[default__.safeIndex(d_126_v1_, len(d_185_v40_))], ((d_186_v41_)[len(d_187_v42_)] if (len(d_187_v42_)) in (d_186_v41_) else d_149_v18_), d_128_v3_, d_138_globalState_), d_181_cf2_, d_138_globalState_)
                    elif True:
                        d_188___mcc_h3_ = source4_.cf0
                        d_189_cf0_ = d_188___mcc_h3_
                        index16_ = default__.safeIndex(961, (d_151_v20_).length(0))
                        (d_151_v20_)[index16_] = (not(d_128_v3_) if d_128_v3_ else d_128_v3_)
                        index17_ = default__.safeIndex(961, (d_151_v20_).length(0))
                        (d_151_v20_)[index17_] = d_128_v3_
                        default__.m0((d_152_v21_)[default__.safeIndex(606, len(d_152_v21_))], d_135_v10_, d_126_v1_, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_138_globalState_)
                        d_190_v43_: _dafny.Map
                        d_190_v43_ = _dafny.Map({(d_126_v1_) + ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]): d_135_v10_})
                        index18_ = default__.safeIndex(223, (d_135_v10_).length(0))
                        (d_135_v10_)[index18_] = len(d_190_v43_)
                        (d_138_globalState_).f1 = not(default__.fm4(d_138_globalState_))
                    index19_ = default__.safeIndex(955, (d_151_v20_).length(0))
                    (d_151_v20_)[index19_] = default__.fm4(d_138_globalState_)
                    index20_ = default__.safeIndex(955, (d_151_v20_).length(0))
                    (d_151_v20_)[index20_] = False
                    d_191_v44_: _dafny.Seq
                    d_191_v44_ = _dafny.SeqWithoutIsStrInference([d_133_v8_])
                    d_192_v45_: _dafny.Map
                    d_192_v45_ = _dafny.Map({(d_191_v44_) + (_dafny.SeqWithoutIsStrInference([d_133_v8_])): d_128_v3_})
                    d_193_v46_: _dafny.MultiSet
                    d_193_v46_ = _dafny.MultiSet([d_128_v3_])
                    index21_ = default__.safeIndex(955, (d_151_v20_).length(0))
                    (d_151_v20_)[index21_] = ((d_192_v45_)[(d_191_v44_).set(default__.safeIndex((d_193_v46_).cardinality, len(d_191_v44_)), d_133_v8_)] if ((d_191_v44_).set(default__.safeIndex((d_193_v46_).cardinality, len(d_191_v44_)), d_133_v8_)) in (d_192_v45_) else d_128_v3_)
                    pass
            pass
        if (d_126_v1_) == ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]):
            (d_138_globalState_).f27 = (d_128_v3_ if d_128_v3_ else (_dafny.SeqWithoutIsStrInference([240])) != (default__.fm6(d_128_v3_, (d_131_v6_).cardinality, d_137_v12_, d_138_globalState_)))
            d_194_v47_: D0
            d_194_v47_ = D0_DC1(default__.safeModuloInt(d_126_v1_, len((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))])))
            source5_ = d_194_v47_
            if source5_.is_DC1:
                d_195___mcc_h4_ = source5_.cf1
                d_196_cf1_ = d_195___mcc_h4_
                d_130_v5_ = (d_130_v5_) | (d_130_v5_)
                index22_ = default__.safeIndex(417, (d_151_v20_).length(0))
                (d_151_v20_)[index22_] = True
                index23_ = default__.safeIndex(417, (d_151_v20_).length(0))
                index24_ = default__.safeIndex(223, (d_135_v10_).length(0))
                rhs7_ = d_158_v26_
                rhs8_ = d_128_v3_
                rhs9_ = (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]
                lhs8_ = d_151_v20_
                lhs9_ = default__.safeIndex(417, (d_151_v20_).length(0))
                lhs10_ = d_135_v10_
                lhs11_ = default__.safeIndex(223, (d_135_v10_).length(0))
                d_158_v26_ = rhs7_
                lhs8_[lhs9_] = rhs8_
                lhs10_[lhs11_] = rhs9_
                (d_138_globalState_).f27 = default__.fm4(d_138_globalState_)
                d_197_v48_: _dafny.Map
                d_197_v48_ = _dafny.Map({d_196_cf1_: d_135_v10_})
                (d_138_globalState_).f19 = (d_196_cf1_) + (len(d_197_v48_))
            elif source5_.is_DC2:
                d_198___mcc_h5_ = source5_.cf2
                d_199___mcc_h6_ = source5_.cf3
                d_200_cf3_ = d_199___mcc_h6_
                d_201_cf2_ = d_198___mcc_h5_
                (d_138_globalState_).f27 = d_128_v3_
                (d_138_globalState_).f27 = (len(d_133_v8_)) <= (-162)
                d_202_v49_: _dafny.Seq
                d_202_v49_ = _dafny.SeqWithoutIsStrInference([d_137_v12_, d_137_v12_])
                d_128_v3_ = (len((d_202_v49_) + (d_202_v49_))) >= (((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]) + (d_201_cf2_))
                (d_138_globalState_).f27 = d_128_v3_
            elif True:
                d_203___mcc_h7_ = source5_.cf0
                d_204_cf0_ = d_203___mcc_h7_
                d_205_v50_: D10
                d_205_v50_ = D10_DC24(default__.fm24(d_129_v4_, d_128_v3_, d_138_globalState_))
                d_206_v51_: C0
                nw15_ = C0()
                nw15_.ctor__((d_205_v50_).cf38)
                d_206_v51_ = nw15_
                (d_138_globalState_).f27 = not(not((d_129_v4_)[default__.safeIndex(default__.fm2(default__.fm3(d_138_globalState_), d_148_v17_, d_149_v18_, d_128_v3_, d_138_globalState_), len(d_129_v4_))]))
                default__.m0(d_151_v20_, d_135_v10_, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_126_v1_, d_138_globalState_)
                index25_ = default__.safeIndex(223, (d_135_v10_).length(0))
                (d_135_v10_)[index25_] = (0) - (((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]) + (851))
            d_207_v52_: D1
            d_207_v52_ = D1_DC5(d_128_v3_, d_128_v3_, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_128_v3_, not(d_128_v3_))
            d_159_v27_ = (d_159_v27_).set((d_207_v52_).cf8, (d_137_v12_) + (d_137_v12_))
            d_208_v53_: D1
            d_208_v53_ = D1_DC3(d_128_v3_)
            d_209_v54_: C10
            nw16_ = C10()
            nw16_.ctor__(d_148_v17_, d_208_v53_)
            d_209_v54_ = nw16_
            d_210_v55_: T0
            nw17_ = C5()
            nw17_.ctor__((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_128_v3_, d_208_v53_)
            d_210_v55_ = nw17_
            d_211_v56_: _dafny.Map
            d_211_v56_ = _dafny.Map({d_209_v54_: (d_210_v55_ if d_128_v3_ else d_210_v55_)})
            d_211_v56_ = (d_211_v56_).set(d_209_v54_, d_210_v55_)
            d_135_v10_ = d_135_v10_
        elif True:
            d_212_v57_: _dafny.Array
            def lambda3_(d_213_v2_):
                def lambda4_(d_214_i7_):
                    return d_213_v2_

                return lambda4_

            init2_ = lambda3_(d_127_v2_)
            nw18_ = _dafny.Array(None, 7)
            for i0_2_ in range(nw18_.length(0)):
                nw18_[i0_2_] = init2_(i0_2_)
            d_212_v57_ = nw18_
            d_215_v58_: _dafny.Array
            d_215_v58_ = d_212_v57_
            d_216_v59_: _dafny.Set
            d_216_v59_ = _dafny.Set({d_215_v58_})
            rhs10_ = d_135_v10_
            rhs11_ = ((d_216_v59_).intersection(d_216_v59_)).isdisjoint(d_216_v59_)
            lhs12_ = d_138_globalState_
            d_135_v10_ = rhs10_
            lhs12_.f27 = rhs11_
            d_217_v60_: _dafny.Array
            nw19_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_217_v60_ = nw19_
            d_218_v61_: _dafny.Map
            d_218_v61_ = _dafny.Map({default__.fm43(d_138_globalState_): d_217_v60_})
            d_219_v62_: D0
            d_219_v62_ = D0_DC1((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])
            d_217_v60_ = ((d_218_v61_)[(d_219_v62_ if d_128_v3_ else D0_DC1(d_126_v1_))] if ((d_219_v62_ if d_128_v3_ else D0_DC1(d_126_v1_))) in (d_218_v61_) else d_217_v60_)
            if True:
                index26_ = default__.safeIndex(54, (d_160_v28_).length(0))
                (d_160_v28_)[index26_] = (d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))]
                d_220_v63_: _dafny.Map
                d_220_v63_ = _dafny.Map({d_151_v20_: d_128_v3_})
                d_220_v63_ = (d_220_v63_).set(d_151_v20_, default__.fm4(d_138_globalState_))
                d_221_v64_: D2
                d_221_v64_ = D2_DC6(d_151_v20_)
                pat_let_tv1_ = d_151_v20_
                def iife46_(_pat_let4_0):
                    def iife47_(d_222_dt__update__tmp_h2_):
                        def iife48_(_pat_let5_0):
                            def iife49_(d_223_dt__update_hcf12_h0_):
                                return D2_DC6(d_223_dt__update_hcf12_h0_)
                            return iife49_(_pat_let5_0)
                        return iife48_(pat_let_tv1_)
                    return iife47_(_pat_let4_0)
                d_151_v20_ = (iife46_(d_221_v64_)).cf12
                d_224_v65_: C4
                nw20_ = C4()
                nw20_.ctor__(d_128_v3_, d_160_v28_)
                d_224_v65_ = nw20_
                d_225_v66_: D25
                d_225_v66_ = D25_DC63(len(d_133_v8_), (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_224_v65_)
                d_226_v67_: D1
                d_226_v67_ = D1_DC3((d_224_v65_).f37)
                d_227_v68_: T0
                nw21_ = C1()
                nw21_.ctor__(default__.safeDivisionInt(len(d_127_v2_), (d_225_v66_).cf120), (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], (d_226_v67_ if True else d_226_v67_))
                d_227_v68_ = nw21_
                d_228_v69_: _dafny.Array
                nw22_ = _dafny.Array(None, 25)
                d_228_v69_ = nw22_
                d_229_v70_: C7
                nw23_ = C7()
                nw23_.ctor__(d_226_v67_)
                d_229_v70_ = nw23_
                index27_ = default__.safeIndex(255, (d_228_v69_).length(0))
                (d_228_v69_)[index27_] = d_229_v70_
                index28_ = default__.safeIndex(867, (d_151_v20_).length(0))
                (d_151_v20_)[index28_] = (d_129_v4_) != (_dafny.SeqWithoutIsStrInference([(d_224_v65_).f37, d_128_v3_]))
                index29_ = default__.safeIndex(255, (d_228_v69_).length(0))
                index30_ = default__.safeIndex(867, (d_151_v20_).length(0))
                rhs12_ = d_227_v68_
                rhs13_ = d_229_v70_
                rhs14_ = d_128_v3_
                rhs15_ = d_128_v3_
                lhs13_ = d_228_v69_
                lhs14_ = default__.safeIndex(255, (d_228_v69_).length(0))
                lhs15_ = d_151_v20_
                lhs16_ = default__.safeIndex(867, (d_151_v20_).length(0))
                d_227_v68_ = rhs12_
                lhs13_[lhs14_] = rhs13_
                d_128_v3_ = rhs14_
                lhs15_[lhs16_] = rhs15_
                d_230_v71_: _dafny.Set
                d_230_v71_ = _dafny.Set({d_158_v26_})
                d_231_v72_: D4
                d_231_v72_ = D4_DC12(True, d_230_v71_)
                (d_138_globalState_).f27 = (d_229_v70_).fm7((d_231_v72_).cf21, not((d_224_v65_).f37), d_138_globalState_)
            elif True:
                d_232_v73_: D16
                d_232_v73_ = D16_DC37(d_129_v4_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qx")), d_158_v26_)
                d_158_v26_ = (d_232_v73_).cf67
                d_233_v74_: C6
                nw24_ = C6()
                nw24_.ctor__((d_126_v1_) + ((0) - (d_126_v1_)))
                d_233_v74_ = nw24_
                default__.m0(d_151_v20_, d_135_v10_, ((d_233_v74_).f34) - ((d_233_v74_).f34), len((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))]), d_138_globalState_)
                d_135_v10_ = d_135_v10_
                d_234_v75_: D8
                d_234_v75_ = D8_DC20(d_126_v1_, (d_233_v74_).f34)
                d_235_v76_: _dafny.Map
                d_235_v76_ = _dafny.Map({d_234_v75_: d_126_v1_})
                d_235_v76_ = (d_235_v76_).set(d_234_v75_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gm"))))
            d_236_v77_: D1
            d_236_v77_ = D1_DC3(d_128_v3_)
            d_237_v78_: T0
            nw25_ = C5()
            nw25_.ctor__(default__.fm2(d_137_v12_, d_148_v17_, _dafny.Map({not(d_128_v3_): d_158_v26_}), d_128_v3_, d_138_globalState_), d_128_v3_, d_236_v77_)
            d_237_v78_ = nw25_
            d_237_v78_ = d_237_v78_
            index31_ = default__.safeIndex(923, (d_124_v0_).length(0))
            (d_124_v0_)[index31_] = d_158_v26_
            index32_ = default__.safeIndex(923, (d_124_v0_).length(0))
            (d_124_v0_)[index32_] = d_158_v26_
        d_238_i8_: int
        d_238_i8_ = 0
        with _dafny.label("3"):
            while False:
                with _dafny.c_label("3"):
                    if (d_238_i8_) >= (100):
                        raise _dafny.Break("3")
                    d_238_i8_ = (d_238_i8_) + (1)
                    if ((d_126_v1_) + (d_126_v1_)) == (d_126_v1_):
                        (d_138_globalState_).f27 = False
                        index33_ = default__.safeIndex(291, (d_151_v20_).length(0))
                        (d_151_v20_)[index33_] = d_128_v3_
                        index34_ = default__.safeIndex(291, (d_151_v20_).length(0))
                        (d_151_v20_)[index34_] = True
                        d_239_v79_: _dafny.Set
                        d_239_v79_ = _dafny.Set({default__.fm2(((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))]).set(default__.safeIndex(d_126_v1_, len((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))])), d_158_v26_), d_148_v17_, d_149_v18_, (d_151_v20_)[default__.safeIndex(291, (d_151_v20_).length(0))], d_138_globalState_)})
                        (d_138_globalState_).f14 = ((d_146_v15_)[(d_239_v79_).ispropersubset(d_239_v79_)] if ((d_239_v79_).ispropersubset(d_239_v79_)) in (d_146_v15_) else (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])
                        d_240_v80_: _dafny.Array
                        nw26_ = _dafny.Array(D1.default()(), 25)
                        d_240_v80_ = nw26_
                        rhs16_ = ((d_133_v8_) | (d_133_v8_)).set((default__.fm54(d_138_globalState_)).cf27, ((0) - ((0) - ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]))) >= (d_126_v1_))
                        rhs17_ = not(not((d_151_v20_)[default__.safeIndex(291, (d_151_v20_).length(0))]))
                        rhs18_ = len((d_146_v15_ if d_128_v3_ else d_146_v15_))
                        rhs19_ = d_240_v80_
                        rhs20_ = d_137_v12_
                        lhs17_ = d_138_globalState_
                        lhs18_ = d_138_globalState_
                        d_133_v8_ = rhs16_
                        lhs17_.f11 = rhs17_
                        lhs18_.f24 = rhs18_
                        d_240_v80_ = rhs19_
                        d_137_v12_ = rhs20_
                        d_241_v81_: _dafny.Map
                        d_241_v81_ = _dafny.Map({d_128_v3_: d_131_v6_})
                        d_241_v81_ = d_241_v81_
                    elif True:
                        (d_138_globalState_).f11 = not(((d_136_v11_)[(d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]] if ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]) in (d_136_v11_) else d_128_v3_))
                        d_133_v8_ = d_133_v8_
                        (d_138_globalState_).f27 = d_128_v3_
                        d_242_v82_: C8
                        nw27_ = C8()
                        nw27_.ctor__(d_126_v1_)
                        d_242_v82_ = nw27_
                        (d_138_globalState_).f14 = (d_126_v1_ if d_128_v3_ else (d_242_v82_.f33) + (len(d_129_v4_)))
                    d_133_v8_ = (d_133_v8_).set((d_129_v4_)[default__.safeIndex((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], len(d_129_v4_))], d_128_v3_)
                    d_243_v83_: D21
                    d_243_v83_ = D21_DC52(len(_dafny.SeqWithoutIsStrInference([not(d_128_v3_)])), not(d_128_v3_), d_128_v3_, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_128_v3_)
                    pat_let_tv2_ = d_128_v3_
                    pat_let_tv3_ = d_135_v10_
                    pat_let_tv4_ = d_135_v10_
                    pat_let_tv5_ = d_126_v1_
                    def iife50_(_pat_let6_0):
                        def iife51_(d_244_dt__update__tmp_h3_):
                            def iife52_(_pat_let7_0):
                                def iife53_(d_245_dt__update_hcf99_h0_):
                                    def iife54_(_pat_let8_0):
                                        def iife55_(d_246_dt__update_hcf98_h0_):
                                            def iife56_(_pat_let9_0):
                                                def iife57_(d_247_dt__update_hcf97_h0_):
                                                    def iife58_(_pat_let10_0):
                                                        def iife59_(d_248_dt__update_hcf100_h0_):
                                                            return D21_DC52(d_247_dt__update_hcf97_h0_, d_246_dt__update_hcf98_h0_, d_245_dt__update_hcf99_h0_, d_248_dt__update_hcf100_h0_, (d_244_dt__update__tmp_h3_).cf101)
                                                        return iife59_(_pat_let10_0)
                                                    return iife58_(pat_let_tv5_)
                                                return iife57_(_pat_let9_0)
                                            return iife56_((pat_let_tv4_)[default__.safeIndex(223, (pat_let_tv3_).length(0))])
                                        return iife55_(_pat_let8_0)
                                    return iife54_(pat_let_tv2_)
                                return iife53_(_pat_let7_0)
                            return iife52_(True)
                        return iife51_(_pat_let6_0)
                    rhs21_ = (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]
                    rhs22_ = (default__.safeDivisionInt((iife50_(d_243_v83_)).cf97, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])) * (d_126_v1_)
                    rhs23_ = 751
                    rhs24_ = len(d_148_v17_)
                    lhs19_ = d_138_globalState_
                    lhs20_ = d_138_globalState_
                    lhs21_ = d_138_globalState_
                    d_126_v1_ = rhs21_
                    lhs19_.f14 = rhs22_
                    lhs20_.f14 = rhs23_
                    lhs21_.f24 = rhs24_
                    d_249_v84_: C9
                    nw28_ = C9()
                    nw28_.ctor__(d_128_v3_, d_129_v4_)
                    d_249_v84_ = nw28_
                    d_250_v85_: C3
                    nw29_ = C3()
                    nw29_.ctor__(D1_DC3(d_249_v84_.f31))
                    d_250_v85_ = nw29_
                    d_251_v86_: _dafny.Map
                    d_251_v86_ = _dafny.Map({d_249_v84_: d_250_v85_})
                    default__.m0(d_151_v20_, d_135_v10_, len(((d_251_v86_).set(d_249_v84_, d_250_v85_)) | (d_251_v86_)), d_126_v1_, d_138_globalState_)
                    pass
            pass
        index35_ = default__.safeIndex(54, (d_160_v28_).length(0))
        (d_160_v28_)[index35_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjcam"))
        index36_ = default__.safeIndex(262, (d_124_v0_).length(0))
        (d_124_v0_)[index36_] = d_158_v26_
        d_252_v87_: D18
        d_252_v87_ = D18_DC42(d_127_v2_, len(d_129_v4_), d_128_v3_, d_128_v3_)
        d_253_v88_: _dafny.Map
        d_253_v88_ = _dafny.Map({(d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))]: d_128_v3_})
        d_254_v89_: _dafny.Seq
        d_254_v89_ = _dafny.SeqWithoutIsStrInference([default__.fm30(d_126_v1_, d_253_v88_, d_138_globalState_), d_146_v15_, d_146_v15_])
        pat_let_tv6_ = d_128_v3_
        pat_let_tv7_ = d_126_v1_
        pat_let_tv8_ = d_128_v3_
        pat_let_tv9_ = d_158_v26_
        pat_let_tv10_ = d_158_v26_
        pat_let_tv11_ = d_158_v26_
        pat_let_tv12_ = d_158_v26_
        pat_let_tv13_ = d_133_v8_
        index37_ = default__.safeIndex(262, (d_124_v0_).length(0))
        def iife60_(_pat_let11_0):
            def iife61_(d_257_dt__update__tmp_h4_):
                def iife62_(_pat_let12_0):
                    def iife63_(d_258_dt__update_hcf74_h0_):
                        def iife64_(_pat_let13_0):
                            def iife65_(d_259_dt__update_hcf73_h0_):
                                def iife66_(_pat_let14_0):
                                    def iife67_(d_260_dt__update_hcf75_h0_):
                                        return D18_DC42((d_257_dt__update__tmp_h4_).cf72, d_259_dt__update_hcf73_h0_, d_258_dt__update_hcf74_h0_, d_260_dt__update_hcf75_h0_)
                                    return iife67_(_pat_let14_0)
                                return iife66_(pat_let_tv8_)
                            return iife65_(_pat_let13_0)
                        return iife64_(pat_let_tv7_)
                    return iife63_(_pat_let12_0)
                return iife62_(pat_let_tv6_)
            return iife61_(_pat_let11_0)
        def lambda5_(source6_):
            if source6_.is_DC20:
                d_263___mcc_h8_ = source6_.cf31
                d_264___mcc_h9_ = source6_.cf32
                d_265_cf32_ = d_264___mcc_h9_
                d_266_cf31_ = d_263___mcc_h8_
                return pat_let_tv9_
            elif source6_.is_DC21:
                d_267___mcc_h10_ = source6_.cf33
                d_268___mcc_h11_ = source6_.cf34
                d_269___mcc_h12_ = source6_.cf35
                d_270_cf35_ = d_269___mcc_h12_
                d_271_cf34_ = d_268___mcc_h11_
                d_272_cf33_ = d_267___mcc_h10_
                return pat_let_tv10_
            elif source6_.is_DC19:
                d_273___mcc_h13_ = source6_.cf30
                d_274_cf30_ = d_273___mcc_h13_
                return pat_let_tv11_
            elif True:
                d_275___mcc_h14_ = source6_.cf36
                d_276_cf36_ = d_275___mcc_h14_
                return pat_let_tv12_

        def iife69_():
            coll38_ = _dafny.Set()
            compr_38_: int
            for compr_38_ in (d_136_v11_).keys.Elements:
                d_277_v90_: int = compr_38_
                if (d_277_v90_) in (d_136_v11_):
                    coll38_ = coll38_.union(_dafny.Set([default__.safeDivisionInt(d_277_v90_, 897)]))
            return _dafny.Set(coll38_)
        def iife68_(_pat_let15_0):
            def iife70_(d_278_dt__update__tmp_h5_):
                def iife71_(_pat_let16_0):
                    def iife72_(d_279_dt__update_hcf36_h0_):
                        return D8_DC22(d_279_dt__update_hcf36_h0_)
                    return iife72_(_pat_let16_0)
                return iife71_(D8_DC19(pat_let_tv13_))
            return iife70_(_pat_let15_0)
        rhs25_ = ((0) - (len((_dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_137_v12_): len(d_137_v12_)}), d_130_v5_, (d_130_v5_).set(d_126_v1_, d_126_v1_)])) + (_dafny.SeqWithoutIsStrInference([d_130_v5_ for d_255_i9_ in range(default__.abs(175))]))))) - ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])
        rhs26_ = _dafny.CodePoint('a')
        rhs27_ = (_dafny.SeqWithoutIsStrInference([d_158_v26_ for d_256_i10_ in range(default__.abs(613))])).set(default__.safeIndex((iife60_(d_252_v87_)).cf73, len(_dafny.SeqWithoutIsStrInference([d_158_v26_ for d_261_i10_ in range(default__.abs(613))]))), d_158_v26_)
        rhs28_ = len((d_254_v89_) + ((_dafny.SeqWithoutIsStrInference([d_146_v15_ for d_262_i11_ in range(default__.abs(511))])) + (d_254_v89_)))
        rhs29_ = lambda5_(iife68_(default__.fm37(d_128_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxhk")), len(iife69_()
        ), d_138_globalState_)))
        lhs22_ = d_138_globalState_
        lhs23_ = d_124_v0_
        lhs24_ = default__.safeIndex(262, (d_124_v0_).length(0))
        lhs25_ = d_138_globalState_
        lhs22_.f24 = rhs25_
        lhs23_[lhs24_] = rhs26_
        d_137_v12_ = rhs27_
        lhs25_.f0 = rhs28_
        d_158_v26_ = rhs29_
        if default__.fm4(d_138_globalState_):
            if d_128_v3_:
                (d_138_globalState_).f27 = False
                (d_138_globalState_).f11 = not((d_128_v3_ if d_128_v3_ else d_128_v3_))
                (d_138_globalState_).f24 = (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]
                (d_138_globalState_).f1 = ((d_136_v11_)[len((d_136_v11_) | (d_136_v11_))] if (len((d_136_v11_) | (d_136_v11_))) in (d_136_v11_) else d_128_v3_)
                (d_138_globalState_).f27 = not (d_128_v3_) or (d_128_v3_)
            elif True:
                d_158_v26_ = ((d_124_v0_)[default__.safeIndex(262, (d_124_v0_).length(0))] if d_128_v3_ else d_158_v26_)
                d_280_v91_: D0
                d_280_v91_ = D0_DC2(d_126_v1_, d_126_v1_)
                (d_138_globalState_).f14 = (d_280_v91_).cf2
                d_281_v92_: _dafny.Array
                nw30_ = _dafny.Array(_dafny.Map({}), 23)
                d_281_v92_ = nw30_
                d_282_v93_: D7
                d_282_v93_ = D7_DC18(d_128_v3_)
                d_283_v94_: _dafny.Map
                d_283_v94_ = _dafny.Map({d_158_v26_: d_282_v93_})
                index38_ = default__.safeIndex(950, (d_281_v92_).length(0))
                (d_281_v92_)[index38_] = d_283_v94_
                index39_ = default__.safeIndex(950, (d_281_v92_).length(0))
                index40_ = default__.safeIndex(54, (d_160_v28_).length(0))
                rhs30_ = ((d_126_v1_) - (len(d_127_v2_)) if not(d_128_v3_) else (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])
                rhs31_ = (d_283_v94_) | (d_283_v94_)
                rhs32_ = default__.fm3(d_138_globalState_)
                lhs26_ = d_138_globalState_
                lhs27_ = d_281_v92_
                lhs28_ = default__.safeIndex(950, (d_281_v92_).length(0))
                lhs29_ = d_160_v28_
                lhs30_ = default__.safeIndex(54, (d_160_v28_).length(0))
                lhs26_.f19 = rhs30_
                lhs27_[lhs28_] = rhs31_
                lhs29_[lhs30_] = rhs32_
                d_147_v16_ = d_147_v16_
                d_130_v5_ = (d_130_v5_).set(default__.safeModuloInt(d_126_v1_, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]), d_126_v1_)
            index41_ = default__.safeIndex(517, (d_151_v20_).length(0))
            (d_151_v20_)[index41_] = ((d_136_v11_)[(0) - ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])] if ((0) - ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])) in (d_136_v11_) else d_128_v3_)
            d_284_v95_: _dafny.MultiSet
            d_284_v95_ = _dafny.MultiSet([d_158_v26_, d_158_v26_, (d_124_v0_)[default__.safeIndex(262, (d_124_v0_).length(0))]])
            d_285_v96_: C3
            nw31_ = C3()
            nw31_.ctor__(D1_DC3((D19_DC46((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))], d_128_v3_, d_128_v3_, d_128_v3_)).cf86))
            d_285_v96_ = nw31_
            d_286_v97_: _dafny.Map
            d_286_v97_ = _dafny.Map({d_128_v3_: d_285_v96_})
            d_287_v98_: _dafny.Set
            d_287_v98_ = _dafny.Set({(0) - (len(d_286_v97_)), (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]})
            d_288_v99_: T0
            nw32_ = C7()
            nw32_.ctor__(default__.fm20(d_128_v3_, d_128_v3_, (d_284_v95_).cardinality, d_287_v98_, d_138_globalState_))
            d_288_v99_ = nw32_
            d_289_v100_: _dafny.Map
            d_289_v100_ = _dafny.Map({d_151_v20_: d_126_v1_})
            d_290_v101_: _dafny.Array
            nw33_ = _dafny.Array(None, 14)
            nw33_[int(0)] = d_289_v100_
            nw33_[int(1)] = d_289_v100_
            nw33_[int(2)] = _dafny.Map({d_151_v20_: (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]})
            nw33_[int(3)] = d_289_v100_
            nw33_[int(4)] = (_dafny.Map({d_151_v20_: 102})) | (d_289_v100_)
            nw33_[int(5)] = d_289_v100_
            nw33_[int(6)] = _dafny.Map({d_151_v20_: 777})
            nw33_[int(7)] = (d_289_v100_) | (d_289_v100_)
            nw33_[int(8)] = d_289_v100_
            nw33_[int(9)] = (_dafny.Map({d_151_v20_: d_126_v1_})) | (d_289_v100_)
            nw33_[int(10)] = (d_289_v100_) | (d_289_v100_)
            nw33_[int(11)] = _dafny.Map({d_151_v20_: d_126_v1_})
            nw33_[int(12)] = d_289_v100_
            nw33_[int(13)] = (d_289_v100_).set(d_151_v20_, (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])
            d_290_v101_ = nw33_
            index42_ = default__.safeIndex(651, (d_290_v101_).length(0))
            (d_290_v101_)[index42_] = (d_289_v100_) | (d_289_v100_)
            index43_ = default__.safeIndex(517, (d_151_v20_).length(0))
            index44_ = default__.safeIndex(651, (d_290_v101_).length(0))
            rhs33_ = (d_285_v96_).fm7(d_128_v3_, d_128_v3_, d_138_globalState_)
            rhs34_ = (default__.fm48(d_158_v26_, d_126_v1_, d_138_globalState_)) | (d_133_v8_)
            rhs35_ = d_288_v99_
            rhs36_ = (d_289_v100_ if (d_128_v3_ if d_128_v3_ else True) else d_289_v100_)
            rhs37_ = not(d_128_v3_)
            lhs31_ = d_151_v20_
            lhs32_ = default__.safeIndex(517, (d_151_v20_).length(0))
            lhs33_ = d_290_v101_
            lhs34_ = default__.safeIndex(651, (d_290_v101_).length(0))
            lhs35_ = d_138_globalState_
            lhs31_[lhs32_] = rhs33_
            d_133_v8_ = rhs34_
            d_288_v99_ = rhs35_
            lhs33_[lhs34_] = rhs36_
            lhs35_.f27 = rhs37_
            d_291_v102_: C7
            nw34_ = C7()
            nw34_.ctor__((d_288_v99_).f28)
            d_291_v102_ = nw34_
            (d_138_globalState_).f27 = (d_151_v20_)[default__.safeIndex(517, (d_151_v20_).length(0))]
            (d_138_globalState_).f27 = not((d_151_v20_)[default__.safeIndex(517, (d_151_v20_).length(0))])
        elif True:
            index45_ = default__.safeIndex(223, (d_135_v10_).length(0))
            (d_135_v10_)[index45_] = (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]
            d_292_v103_: _dafny.MultiSet
            d_292_v103_ = _dafny.MultiSet([((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]) not in (d_127_v2_), d_128_v3_, (d_128_v3_) and (d_128_v3_), d_128_v3_, d_128_v3_])
            d_292_v103_ = (d_292_v103_) | ((d_292_v103_).intersection(d_292_v103_))
            d_128_v3_ = ((-769) + ((0) - ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]))) < (default__.fm2(((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))]).set(default__.safeIndex(len(_dafny.Set({len((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))])})), len((d_160_v28_)[default__.safeIndex(54, (d_160_v28_).length(0))])), ((d_149_v18_)[d_128_v3_] if (d_128_v3_) in (d_149_v18_) else (d_124_v0_)[default__.safeIndex(262, (d_124_v0_).length(0))])), d_148_v17_, d_149_v18_, d_128_v3_, d_138_globalState_))
            default__.m0(d_151_v20_, d_135_v10_, ((d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))]) - (((d_292_v103_)[d_128_v3_] if (d_128_v3_) in (d_292_v103_) else (d_135_v10_)[default__.safeIndex(223, (d_135_v10_).length(0))])), 811, d_138_globalState_)
            (d_138_globalState_).f11 = not((d_126_v1_) != (len(d_148_v17_)))
        _dafny.print(_dafny.string_of((d_124_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_124_v0_)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_126_v1_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_127_v2_) == (_dafny.SeqWithoutIsStrInference([554, 554]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_128_v3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_129_v4_) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_130_v5_) == (_dafny.Map({554: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_131_v6_) == (_dafny.MultiSet([554]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_132_v7_).cf0) == (_dafny.Map({554: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_133_v8_) == (_dafny.Map({False: False, True: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_135_v10_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_136_v11_) == (_dafny.Map({554: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_137_v12_) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f1))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[17]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[18]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[19]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[20]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[21]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[22]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f2)[23]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f3))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f9) == (_dafny.SeqWithoutIsStrInference([554, 554]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f10) == (_dafny.Map({554: 1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f11))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f12) == (_dafny.SeqWithoutIsStrInference([True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f13))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f14))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f16) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f17))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f18)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f19))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_.f20) == (_dafny.Map({554: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_138_globalState_).f21) == (_dafny.MultiSet([554]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f22))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f23))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f24))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_138_globalState_).f25).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_138_globalState_).f26))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_138_globalState_.f27))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[0]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[1]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[2]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[3]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[4]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[5]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[6]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[7]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[8]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[9]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[10]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[11]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[12]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[13]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[14]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[15]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[16]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[17]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[18]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[19]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[20]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[21]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[22]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[23]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_139_v13_)[24]) == (_dafny.SeqWithoutIsStrInference([True]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[0]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[1]) == (_dafny.Set({_dafny.Map({True: 525})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[2]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[3]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[4]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[5]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[6]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[7]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[8]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[9]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[10]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[11]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[12]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[13]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[14]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[15]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[16]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[17]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[18]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[19]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_141_v14_)[20]) == (_dafny.Set({_dafny.Map({True: 554}), _dafny.Map({False: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_146_v15_) == (_dafny.Map({True: 554}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_147_v16_) == (_dafny.Set({_dafny.Map({True: 554})}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_148_v17_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_149_v18_) == (_dafny.Map({True: _dafny.CodePoint('x')}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_150_v19_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference([554, 554]): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_151_v20_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_152_v21_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_153_v22_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_154_i3_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_158_v26_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_159_v27_) == (_dafny.Map({True: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mgkmsxlaimgkmsxlai"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[0]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[1]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[2]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_160_v28_)[3]) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j'), _dafny.CodePoint('j')]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[4]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[5]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[6]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[7]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[8]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_160_v28_)[9]).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_163_i6_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_238_i8_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_252_v87_).cf72) == (_dafny.SeqWithoutIsStrInference([554, 554]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v87_).cf73))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v87_).cf74))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_252_v87_).cf75))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_253_v88_) == (_dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pjcam")): True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_254_v89_) == (_dafny.SeqWithoutIsStrInference([_dafny.Map({False: -1020}), _dafny.Map({True: 554}), _dafny.Map({True: 554})]))))
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

class D0_DC2(D0, NamedTuple('DC2', [('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
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
        return lambda: D1_DC4(int(0), False)
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

class D1_DC4(D1, NamedTuple('DC4', [('cf5', Any), ('cf6', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf5 == __o.cf5 and self.cf6 == __o.cf6
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC5(D1, NamedTuple('DC5', [('cf7', Any), ('cf8', Any), ('cf9', Any), ('cf10', Any), ('cf11', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC5({_dafny.string_of(self.cf7)}, {_dafny.string_of(self.cf8)}, {_dafny.string_of(self.cf9)}, {_dafny.string_of(self.cf10)}, {_dafny.string_of(self.cf11)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC5) and self.cf7 == __o.cf7 and self.cf8 == __o.cf8 and self.cf9 == __o.cf9 and self.cf10 == __o.cf10 and self.cf11 == __o.cf11
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf4', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf4)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf4 == __o.cf4
    def __hash__(self) -> int:
        return super().__hash__()


class D2:
    @classmethod
    def default(cls, ):
        return lambda: D2_DC7(False, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
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

class D2_DC7(D2, NamedTuple('DC7', [('cf13', Any), ('cf14', Any), ('cf15', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf13)}, {_dafny.string_of(self.cf14)}, {self.cf15.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf13 == __o.cf13 and self.cf14 == __o.cf14 and self.cf15 == __o.cf15
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf16', Any), ('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf16)}, {_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf16 == __o.cf16 and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC6(D2, NamedTuple('DC6', [('cf12', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6({_dafny.string_of(self.cf12)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6) and self.cf12 == __o.cf12
    def __hash__(self) -> int:
        return super().__hash__()


class D3:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC9(D3, NamedTuple('DC9', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(_dafny.Seq({}))
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

class D4_DC11(D4, NamedTuple('DC11', [('cf20', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf20)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf20 == __o.cf20
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf21', Any), ('cf22', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf21)}, {_dafny.string_of(self.cf22)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf21 == __o.cf21 and self.cf22 == __o.cf22
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf19', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf19)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf19 == __o.cf19
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC13(D4, NamedTuple('DC13', [('cf23', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC13({_dafny.string_of(self.cf23)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC13) and self.cf23 == __o.cf23
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)

class D5_DC14(D5, NamedTuple('DC14', [('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: D6_DC16(False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D6_DC16)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC16(D6, NamedTuple('DC16', [('cf26', Any), ('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC16({_dafny.string_of(self.cf26)}, {_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC16) and self.cf26 == __o.cf26 and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()

class D6_DC15(D6, NamedTuple('DC15', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC18(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D7_DC18)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)

class D7_DC18(D7, NamedTuple('DC18', [('cf29', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC18({_dafny.string_of(self.cf29)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC18) and self.cf29 == __o.cf29
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC17(D7, NamedTuple('DC17', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC20(int(0), int(0))
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

class D8_DC20(D8, NamedTuple('DC20', [('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC21(D8, NamedTuple('DC21', [('cf33', Any), ('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC21({_dafny.string_of(self.cf33)}, {_dafny.string_of(self.cf34)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC21) and self.cf33 == __o.cf33 and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC19(D8, NamedTuple('DC19', [('cf30', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({_dafny.string_of(self.cf30)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf30 == __o.cf30
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC22(D8, NamedTuple('DC22', [('cf36', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC22({_dafny.string_of(self.cf36)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC22) and self.cf36 == __o.cf36
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D9_DC23)

class D9_DC23(D9, NamedTuple('DC23', [('cf37', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC23({_dafny.string_of(self.cf37)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC23) and self.cf37 == __o.cf37
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC25(_dafny.Array(None, 0), False, _dafny.Seq({}), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D10_DC25)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)

class D10_DC25(D10, NamedTuple('DC25', [('cf39', Any), ('cf40', Any), ('cf41', Any), ('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC25({_dafny.string_of(self.cf39)}, {_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)}, {_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC25) and self.cf39 == __o.cf39 and self.cf40 == __o.cf40 and self.cf41 == __o.cf41 and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC24(D10, NamedTuple('DC24', [('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)

class D11_DC26(D11, NamedTuple('DC26', [('cf43', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf43)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf43 == __o.cf43
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: D12_DC28(False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), False, int(0), D2.default()())
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D12_DC29)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D12_DC30)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D12_DC27)

class D12_DC28(D12, NamedTuple('DC28', [('cf45', Any), ('cf46', Any), ('cf47', Any), ('cf48', Any), ('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf45)}, {self.cf46.VerbatimString(True)}, {_dafny.string_of(self.cf47)}, {_dafny.string_of(self.cf48)}, {_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf45 == __o.cf45 and self.cf46 == __o.cf46 and self.cf47 == __o.cf47 and self.cf48 == __o.cf48 and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC29(D12, NamedTuple('DC29', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC29({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC29) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC30(D12, NamedTuple('DC30', [('cf51', Any), ('cf52', Any), ('cf53', Any), ('cf54', Any), ('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC30({_dafny.string_of(self.cf51)}, {_dafny.string_of(self.cf52)}, {_dafny.string_of(self.cf53)}, {_dafny.string_of(self.cf54)}, {_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC30) and self.cf51 == __o.cf51 and self.cf52 == __o.cf52 and self.cf53 == __o.cf53 and self.cf54 == __o.cf54 and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()

class D12_DC27(D12, NamedTuple('DC27', [('cf44', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC27({_dafny.string_of(self.cf44)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC27) and self.cf44 == __o.cf44
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)

class D13_DC31(D13, NamedTuple('DC31', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC32(D14, NamedTuple('DC32', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: D15_DC34(False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D15_DC34)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)

class D15_DC34(D15, NamedTuple('DC34', [('cf59', Any), ('cf60', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC34({_dafny.string_of(self.cf59)}, {_dafny.string_of(self.cf60)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC34) and self.cf59 == __o.cf59 and self.cf60 == __o.cf60
    def __hash__(self) -> int:
        return super().__hash__()

class D15_DC33(D15, NamedTuple('DC33', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: D16_DC36(False, False, _dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D16_DC36)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D16_DC37)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D16_DC35)

class D16_DC36(D16, NamedTuple('DC36', [('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC36({_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC36) and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC37(D16, NamedTuple('DC37', [('cf65', Any), ('cf66', Any), ('cf67', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC37({_dafny.string_of(self.cf65)}, {self.cf66.VerbatimString(True)}, {_dafny.string_of(self.cf67)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC37) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67
    def __hash__(self) -> int:
        return super().__hash__()

class D16_DC35(D16, NamedTuple('DC35', [('cf61', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC35({_dafny.string_of(self.cf61)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC35) and self.cf61 == __o.cf61
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: D17_DC39(_dafny.Array(None, 0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D17_DC39)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D17_DC38)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D17_DC40)

class D17_DC39(D17, NamedTuple('DC39', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC39({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC39) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC38(D17, NamedTuple('DC38', [('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC38({_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC38) and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D17_DC40(D17, NamedTuple('DC40', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC40({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC40) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: D18_DC42(_dafny.Seq({}), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D18_DC42)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D18_DC43)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D18_DC41)

class D18_DC42(D18, NamedTuple('DC42', [('cf72', Any), ('cf73', Any), ('cf74', Any), ('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC42({_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)}, {_dafny.string_of(self.cf74)}, {_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC42) and self.cf72 == __o.cf72 and self.cf73 == __o.cf73 and self.cf74 == __o.cf74 and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC43(D18, NamedTuple('DC43', [('cf76', Any), ('cf77', Any), ('cf78', Any), ('cf79', Any), ('cf80', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC43({_dafny.string_of(self.cf76)}, {_dafny.string_of(self.cf77)}, {self.cf78.VerbatimString(True)}, {_dafny.string_of(self.cf79)}, {_dafny.string_of(self.cf80)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC43) and self.cf76 == __o.cf76 and self.cf77 == __o.cf77 and self.cf78 == __o.cf78 and self.cf79 == __o.cf79 and self.cf80 == __o.cf80
    def __hash__(self) -> int:
        return super().__hash__()

class D18_DC41(D18, NamedTuple('DC41', [('cf71', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC41({_dafny.string_of(self.cf71)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC41) and self.cf71 == __o.cf71
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: D19_DC45(int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D19_DC45)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D19_DC46)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D19_DC44)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D19_DC47)

class D19_DC45(D19, NamedTuple('DC45', [('cf82', Any), ('cf83', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC45({_dafny.string_of(self.cf82)}, {_dafny.string_of(self.cf83)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC45) and self.cf82 == __o.cf82 and self.cf83 == __o.cf83
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC46(D19, NamedTuple('DC46', [('cf84', Any), ('cf85', Any), ('cf86', Any), ('cf87', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC46({_dafny.string_of(self.cf84)}, {_dafny.string_of(self.cf85)}, {_dafny.string_of(self.cf86)}, {_dafny.string_of(self.cf87)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC46) and self.cf84 == __o.cf84 and self.cf85 == __o.cf85 and self.cf86 == __o.cf86 and self.cf87 == __o.cf87
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC44(D19, NamedTuple('DC44', [('cf81', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC44({_dafny.string_of(self.cf81)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC44) and self.cf81 == __o.cf81
    def __hash__(self) -> int:
        return super().__hash__()

class D19_DC47(D19, NamedTuple('DC47', [('cf88', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC47({_dafny.string_of(self.cf88)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC47) and self.cf88 == __o.cf88
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC49(_dafny.CodePoint('D'), int(0), False, False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC49(self) -> bool:
        return isinstance(self, D20_DC49)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D20_DC48)
    @property
    def is_DC50(self) -> bool:
        return isinstance(self, D20_DC50)

class D20_DC49(D20, NamedTuple('DC49', [('cf90', Any), ('cf91', Any), ('cf92', Any), ('cf93', Any), ('cf94', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC49({_dafny.string_of(self.cf90)}, {_dafny.string_of(self.cf91)}, {_dafny.string_of(self.cf92)}, {_dafny.string_of(self.cf93)}, {_dafny.string_of(self.cf94)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC49) and self.cf90 == __o.cf90 and self.cf91 == __o.cf91 and self.cf92 == __o.cf92 and self.cf93 == __o.cf93 and self.cf94 == __o.cf94
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC48(D20, NamedTuple('DC48', [('cf89', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC48({_dafny.string_of(self.cf89)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC48) and self.cf89 == __o.cf89
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC50(D20, NamedTuple('DC50', [('cf95', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC50({_dafny.string_of(self.cf95)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC50) and self.cf95 == __o.cf95
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC52(int(0), False, False, int(0), False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC52(self) -> bool:
        return isinstance(self, D21_DC52)
    @property
    def is_DC53(self) -> bool:
        return isinstance(self, D21_DC53)
    @property
    def is_DC51(self) -> bool:
        return isinstance(self, D21_DC51)
    @property
    def is_DC54(self) -> bool:
        return isinstance(self, D21_DC54)

class D21_DC52(D21, NamedTuple('DC52', [('cf97', Any), ('cf98', Any), ('cf99', Any), ('cf100', Any), ('cf101', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC52({_dafny.string_of(self.cf97)}, {_dafny.string_of(self.cf98)}, {_dafny.string_of(self.cf99)}, {_dafny.string_of(self.cf100)}, {_dafny.string_of(self.cf101)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC52) and self.cf97 == __o.cf97 and self.cf98 == __o.cf98 and self.cf99 == __o.cf99 and self.cf100 == __o.cf100 and self.cf101 == __o.cf101
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC53(D21, NamedTuple('DC53', [('cf102', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC53({_dafny.string_of(self.cf102)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC53) and self.cf102 == __o.cf102
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC51(D21, NamedTuple('DC51', [('cf96', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC51({_dafny.string_of(self.cf96)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC51) and self.cf96 == __o.cf96
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC54(D21, NamedTuple('DC54', [('cf103', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC54({_dafny.string_of(self.cf103)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC54) and self.cf103 == __o.cf103
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC56(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0), False, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC56(self) -> bool:
        return isinstance(self, D22_DC56)
    @property
    def is_DC57(self) -> bool:
        return isinstance(self, D22_DC57)
    @property
    def is_DC55(self) -> bool:
        return isinstance(self, D22_DC55)
    @property
    def is_DC58(self) -> bool:
        return isinstance(self, D22_DC58)

class D22_DC56(D22, NamedTuple('DC56', [('cf105', Any), ('cf106', Any), ('cf107', Any), ('cf108', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC56({self.cf105.VerbatimString(True)}, {_dafny.string_of(self.cf106)}, {_dafny.string_of(self.cf107)}, {_dafny.string_of(self.cf108)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC56) and self.cf105 == __o.cf105 and self.cf106 == __o.cf106 and self.cf107 == __o.cf107 and self.cf108 == __o.cf108
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC57(D22, NamedTuple('DC57', [('cf109', Any), ('cf110', Any), ('cf111', Any), ('cf112', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC57({_dafny.string_of(self.cf109)}, {_dafny.string_of(self.cf110)}, {_dafny.string_of(self.cf111)}, {_dafny.string_of(self.cf112)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC57) and self.cf109 == __o.cf109 and self.cf110 == __o.cf110 and self.cf111 == __o.cf111 and self.cf112 == __o.cf112
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC55(D22, NamedTuple('DC55', [('cf104', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC55({_dafny.string_of(self.cf104)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC55) and self.cf104 == __o.cf104
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC58(D22, NamedTuple('DC58', [('cf113', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC58({_dafny.string_of(self.cf113)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC58) and self.cf113 == __o.cf113
    def __hash__(self) -> int:
        return super().__hash__()


class D23:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC59(self) -> bool:
        return isinstance(self, D23_DC59)

class D23_DC59(D23, NamedTuple('DC59', [('cf114', Any)])):
    def __dafnystr__(self) -> str:
        return f'D23.DC59({_dafny.string_of(self.cf114)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D23_DC59) and self.cf114 == __o.cf114
    def __hash__(self) -> int:
        return super().__hash__()


class D24:
    @classmethod
    def default(cls, ):
        return lambda: D24_DC61(False, False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC61(self) -> bool:
        return isinstance(self, D24_DC61)
    @property
    def is_DC60(self) -> bool:
        return isinstance(self, D24_DC60)

class D24_DC61(D24, NamedTuple('DC61', [('cf116', Any), ('cf117', Any), ('cf118', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC61({_dafny.string_of(self.cf116)}, {_dafny.string_of(self.cf117)}, {_dafny.string_of(self.cf118)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC61) and self.cf116 == __o.cf116 and self.cf117 == __o.cf117 and self.cf118 == __o.cf118
    def __hash__(self) -> int:
        return super().__hash__()

class D24_DC60(D24, NamedTuple('DC60', [('cf115', Any)])):
    def __dafnystr__(self) -> str:
        return f'D24.DC60({_dafny.string_of(self.cf115)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D24_DC60) and self.cf115 == __o.cf115
    def __hash__(self) -> int:
        return super().__hash__()


class D25:
    @classmethod
    def default(cls, ):
        return lambda: D25_DC63(int(0), int(0), None)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC63(self) -> bool:
        return isinstance(self, D25_DC63)
    @property
    def is_DC62(self) -> bool:
        return isinstance(self, D25_DC62)

class D25_DC63(D25, NamedTuple('DC63', [('cf120', Any), ('cf121', Any), ('cf122', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC63({_dafny.string_of(self.cf120)}, {_dafny.string_of(self.cf121)}, {_dafny.string_of(self.cf122)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC63) and self.cf120 == __o.cf120 and self.cf121 == __o.cf121 and self.cf122 == __o.cf122
    def __hash__(self) -> int:
        return super().__hash__()

class D25_DC62(D25, NamedTuple('DC62', [('cf119', Any)])):
    def __dafnystr__(self) -> str:
        return f'D25.DC62({_dafny.string_of(self.cf119)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D25_DC62) and self.cf119 == __o.cf119
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m1(self, globalState):
        pass

    def m2(self, p0, p1, globalState):
        pass


class GlobalState:
    def  __init__(self):
        self.f0: int = int(0)
        self.f1: bool = False
        self.f2: _dafny.Array = _dafny.Array(None, 0)
        self.f11: bool = False
        self.f12: _dafny.Seq = _dafny.Seq({})
        self.f14: int = int(0)
        self.f16: _dafny.Seq = _dafny.Seq({})
        self.f19: int = int(0)
        self.f20: _dafny.Map = _dafny.Map({})
        self.f24: int = int(0)
        self.f27: bool = False
        self._f3: bool = False
        self._f4: int = int(0)
        self._f5: bool = False
        self._f6: int = int(0)
        self._f7: bool = False
        self._f8: bool = False
        self._f9: _dafny.Seq = _dafny.Seq({})
        self._f10: _dafny.Map = _dafny.Map({})
        self._f13: bool = False
        self._f15: _dafny.Array = _dafny.Array(None, 0)
        self._f17: bool = False
        self._f18: _dafny.Array = _dafny.Array(None, 0)
        self._f21: _dafny.MultiSet = _dafny.MultiSet({})
        self._f22: int = int(0)
        self._f23: int = int(0)
        self._f25: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f26: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9, f10, f11, f12, f13, f14, f15, f16, f17, f18, f19, f20, f21, f22, f23, f24, f25, f26, f27):
        (self).f0 = f0
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
        (self).f12 = f12
        (self)._f13 = f13
        (self).f14 = f14
        (self)._f15 = f15
        (self).f16 = f16
        (self)._f17 = f17
        (self)._f18 = f18
        (self).f19 = f19
        (self).f20 = f20
        (self)._f21 = f21
        (self)._f22 = f22
        (self)._f23 = f23
        (self).f24 = f24
        (self)._f25 = f25
        (self)._f26 = f26
        (self).f27 = f27

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
    def f13(self):
        return self._f13
    @property
    def f15(self):
        return self._f15
    @property
    def f17(self):
        return self._f17
    @property
    def f18(self):
        return self._f18
    @property
    def f21(self):
        return self._f21
    @property
    def f22(self):
        return self._f22
    @property
    def f23(self):
        return self._f23
    @property
    def f25(self):
        return self._f25
    @property
    def f26(self):
        return self._f26

class C0:
    def  __init__(self):
        self._f43: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f43):
        (self)._f43 = f43

    def fm19(self, globalState):
        return _dafny.Map({(-359 if True else len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): False}))]))): (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohnxf")))}) if not(True) else _dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, not(True)])), -58]))}))})

    @property
    def f43(self):
        return self._f43

class C1(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        self.f42: int = int(0)
        self._f41: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f41, f42, f28):
        (self)._f41 = f41
        (self).f42 = f42
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        if (_dafny.MultiSet([(self).f41])).ispropersubset(_dafny.MultiSet([(self).f41])):
            return not (False) or (False)
        elif True:
            return False

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('w')

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_293_v0_: str
        d_293_v0_ = _dafny.CodePoint('x')
        d_294_v1_: _dafny.Map
        d_294_v1_ = _dafny.Map({self.f42: d_293_v0_})
        d_295_v2_: _dafny.Seq
        d_295_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frvhu"))
        d_296_v3_: bool
        d_296_v3_ = True
        d_297_v4_: _dafny.Set
        d_297_v4_ = _dafny.Set({d_296_v3_, not(d_296_v3_), d_296_v3_, d_296_v3_})
        d_298_v5_: _dafny.Seq
        d_298_v5_ = _dafny.SeqWithoutIsStrInference([d_296_v3_, False])
        d_294_v1_ = (d_294_v1_).set(default__.fm2((d_295_v2_).set(default__.safeIndex(self.f42, len(d_295_v2_)), d_293_v0_), d_297_v4_, default__.fm18((self).fm7(d_296_v3_, d_296_v3_, globalState), (d_298_v5_)[default__.safeIndex(self.f42, len(d_298_v5_))], globalState), default__.fm4(globalState), globalState), d_293_v0_)
        d_299_v6_: _dafny.Array
        def lambda6_(d_300_i0_):
            return default__.safeModuloInt(d_300_i0_, self.f42)

        init3_ = lambda6_
        nw35_ = _dafny.Array(None, 2)
        for i0_3_ in range(nw35_.length(0)):
            nw35_[i0_3_] = init3_(i0_3_)
        d_299_v6_ = nw35_
        d_301_v7_: _dafny.Map
        d_301_v7_ = _dafny.Map({d_299_v6_: self.f42})
        (globalState).f14 = ((d_301_v7_)[d_299_v6_] if (d_299_v6_) in (d_301_v7_) else (self).f41)
        rhs38_ = (self).f41
        rhs39_ = (_dafny.SeqWithoutIsStrInference([d_293_v0_ for d_302_i1_ in range(default__.abs(-813))])).set(default__.safeIndex((self).f41, len(_dafny.SeqWithoutIsStrInference([d_293_v0_ for d_303_i1_ in range(default__.abs(-813))]))), d_293_v0_)
        lhs36_ = self
        lhs36_.f42 = rhs38_
        r1 = rhs39_
        d_304_v8_: C0
        nw36_ = C0()
        nw36_.ctor__((d_297_v4_) | (d_297_v4_))
        d_304_v8_ = nw36_
        if False:
            d_305_v9_: _dafny.Map
            d_305_v9_ = _dafny.Map({d_296_v3_: d_296_v3_})
            d_305_v9_ = (d_305_v9_).set(d_296_v3_, d_296_v3_)
            d_306_v10_: C0
            nw37_ = C0()
            nw37_.ctor__((d_304_v8_).f43)
            d_306_v10_ = nw37_
            (globalState).f24 = (len((d_295_v2_) + (d_295_v2_))) + (self.f42)
            d_307_v11_: D1
            d_307_v11_ = D1_DC4(self.f42, d_296_v3_)
            index46_ = default__.safeIndex(993, (d_299_v6_).length(0))
            (d_299_v6_)[index46_] = (0) - ((d_307_v11_).cf5)
            d_308_v12_: _dafny.Set
            d_308_v12_ = _dafny.Set({d_293_v0_, d_293_v0_})
            index47_ = default__.safeIndex(993, (d_299_v6_).length(0))
            (d_299_v6_)[index47_] = ((self).f41 if not(d_296_v3_) else len(d_308_v12_))
            d_309_v13_: _dafny.Map
            d_309_v13_ = _dafny.Map({self.f42: (d_296_v3_) or (d_296_v3_)})
            d_309_v13_ = (d_309_v13_).set((self).f41, d_296_v3_)
        elif True:
            (globalState).f0 = (self).f41
            if d_296_v3_:
                index48_ = default__.safeIndex(244, (d_299_v6_).length(0))
                (d_299_v6_)[index48_] = (self).f41
                index49_ = default__.safeIndex(244, (d_299_v6_).length(0))
                (d_299_v6_)[index49_] = default__.safeDivisionInt(self.f42, (self).f41)
                (globalState).f14 = self.f42
                (globalState).f27 = (self).fm7(d_296_v3_, d_296_v3_, globalState)
                (globalState).f11 = (d_295_v2_) < (d_295_v2_)
                d_310_v14_: _dafny.Set
                d_310_v14_ = _dafny.Set({(self).f41})
                d_311_v15_: _dafny.Set
                d_311_v15_ = _dafny.Set({(self).f28, default__.fm20(d_296_v3_, d_296_v3_, (self).f41, d_310_v14_, globalState), (self).f28, D1_DC3(d_296_v3_)})
                d_311_v15_ = d_311_v15_
            elif True:
                d_312_v16_: _dafny.Seq
                d_312_v16_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_296_v3_, True])), self.f42, (self.f42) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "knsnppx")))), 865])
                (globalState).f14 = (d_312_v16_)[default__.safeIndex(self.f42, len(d_312_v16_))]
                d_313_v17_: _dafny.Array
                nw38_ = _dafny.Array(False, 7)
                d_313_v17_ = nw38_
                d_314_v18_: _dafny.Map
                d_314_v18_ = _dafny.Map({d_296_v3_: d_293_v0_})
                default__.m0(d_313_v17_, (d_299_v6_ if d_296_v3_ else d_299_v6_), default__.safeDivisionInt(default__.fm2(d_295_v2_, (d_304_v8_).f43, d_314_v18_, d_296_v3_, globalState), (self).f41), (self.f42) + ((0) - (self.f42)), globalState)
                (globalState).f11 = d_296_v3_
                d_315_v19_: _dafny.Seq
                d_315_v19_ = _dafny.SeqWithoutIsStrInference([d_313_v17_, d_313_v17_, d_313_v17_])
                d_315_v19_ = _dafny.SeqWithoutIsStrInference([d_313_v17_])
                d_316_v20_: C0
                nw39_ = C0()
                nw39_.ctor__(d_297_v4_)
                d_316_v20_ = nw39_
            d_317_v21_: _dafny.MultiSet
            d_317_v21_ = _dafny.MultiSet([d_295_v2_])
            (globalState).f1 = (d_317_v21_) == ((d_317_v21_ if False else d_317_v21_))
            (globalState).f16 = (d_298_v5_).set(default__.safeIndex((self).f41, len(d_298_v5_)), d_296_v3_)
            (globalState).f19 = default__.safeDivisionInt(self.f42, (self).f41)
        d_318_v22_: _dafny.Set
        d_318_v22_ = _dafny.Set({d_293_v0_})
        d_319_v23_: D4
        d_319_v23_ = D4_DC12(d_296_v3_, d_318_v22_)
        d_320_v24_: _dafny.Seq
        d_320_v24_ = _dafny.SeqWithoutIsStrInference([d_319_v23_, d_319_v23_])
        (self).f42 = (0) - (len(((_dafny.SeqWithoutIsStrInference([d_319_v23_])) + (_dafny.SeqWithoutIsStrInference([d_319_v23_]))) + (d_320_v24_)))
        d_321_v25_: D2
        d_321_v25_ = D2_DC8(d_296_v3_, d_299_v6_)
        pat_let_tv14_ = d_296_v3_
        def iife73_(_pat_let17_0):
            def iife74_(d_322_dt__update__tmp_h0_):
                def iife75_(_pat_let18_0):
                    def iife76_(d_323_dt__update_hcf16_h0_):
                        return D2_DC8(d_323_dt__update_hcf16_h0_, (d_322_dt__update__tmp_h0_).cf17)
                    return iife76_(_pat_let18_0)
                return iife75_((D1_DC4(self.f42, pat_let_tv14_)).cf6)
            return iife74_(_pat_let17_0)
        r0 = iife73_(d_321_v25_)
        r1 = (_dafny.SeqWithoutIsStrInference([d_293_v0_ for d_324_i2_ in range(default__.abs(990))])) + (d_295_v2_)
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_325_v0_: _dafny.Seq
        d_325_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        if not((p0) in ((d_325_v0_) + (d_325_v0_))):
            (globalState).f1 = not (p0) or (p0)
            d_326_v1_: _dafny.Array
            def lambda7_(d_327_p0_):
                def lambda8_(d_328_i0_):
                    return _dafny.Map({d_327_p0_: d_327_p0_})

                return lambda8_

            init4_ = lambda7_(p0)
            nw40_ = _dafny.Array(None, 27)
            for i0_4_ in range(nw40_.length(0)):
                nw40_[i0_4_] = init4_(i0_4_)
            d_326_v1_ = nw40_
            source7_ = d_326_v1_
            d_329___mcc_h0_ = source7_
            d_330_cf18_ = d_329___mcc_h0_
            d_331_v2_: _dafny.Array
            def lambda9_(d_332_p0_):
                def lambda10_(d_333_i1_):
                    return D1_DC5(d_332_p0_, d_332_p0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ymy"))), not(d_332_p0_), not(d_332_p0_))

                return lambda10_

            init5_ = lambda9_(p0)
            nw41_ = _dafny.Array(None, 7)
            for i0_5_ in range(nw41_.length(0)):
                nw41_[i0_5_] = init5_(i0_5_)
            d_331_v2_ = nw41_
            d_334_v3_: _dafny.Seq
            d_334_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jvduqat"))
            d_335_v4_: _dafny.Set
            d_335_v4_ = _dafny.Set({p0})
            rhs40_ = (self).f41
            rhs41_ = d_331_v2_
            rhs42_ = default__.fm4(globalState)
            rhs43_ = d_334_v3_
            rhs44_ = (d_335_v4_).intersection(d_335_v4_)
            lhs37_ = globalState
            lhs38_ = globalState
            lhs37_.f24 = rhs40_
            d_331_v2_ = rhs41_
            lhs38_.f27 = rhs42_
            d_334_v3_ = rhs43_
            d_335_v4_ = rhs44_
            (globalState).f11 = p0
            d_336_v5_: _dafny.Array
            def lambda11_(d_337_v4_):
                def lambda12_(d_338_i2_):
                    return (d_337_v4_).ispropersubset(d_337_v4_)

                return lambda12_

            init6_ = lambda11_(d_335_v4_)
            nw42_ = _dafny.Array(None, 5)
            for i0_6_ in range(nw42_.length(0)):
                nw42_[i0_6_] = init6_(i0_6_)
            d_336_v5_ = nw42_
            d_336_v5_ = d_336_v5_
            (globalState).f27 = p0
            (globalState).f0 = self.f42
            d_339_v6_: _dafny.Array
            def lambda13_(d_340_p0_):
                def lambda14_(d_341_i3_):
                    return d_340_p0_

                return lambda14_

            init7_ = lambda13_(p0)
            nw43_ = _dafny.Array(None, 20)
            for i0_7_ in range(nw43_.length(0)):
                nw43_[i0_7_] = init7_(i0_7_)
            d_339_v6_ = nw43_
            index50_ = default__.safeIndex(543, (d_339_v6_).length(0))
            (d_339_v6_)[index50_] = (not(p0) if p0 else p0)
            d_342_v7_: _dafny.Map
            d_342_v7_ = _dafny.Map({self.f42: p0})
            index51_ = default__.safeIndex(543, (d_339_v6_).length(0))
            (d_339_v6_)[index51_] = (d_325_v0_)[default__.safeIndex((0) - (len(d_342_v7_)), len(d_325_v0_))]
            d_343_v8_: _dafny.Array
            def lambda15_(d_344_i4_):
                return default__.safeDivisionInt(d_344_i4_, self.f42)

            init8_ = lambda15_
            nw44_ = _dafny.Array(None, 2)
            for i0_8_ in range(nw44_.length(0)):
                nw44_[i0_8_] = init8_(i0_8_)
            d_343_v8_ = nw44_
            index52_ = default__.safeIndex(584, (d_343_v8_).length(0))
            (d_343_v8_)[index52_] = self.f42
            d_345_v9_: _dafny.MultiSet
            d_345_v9_ = _dafny.MultiSet([(self).f41])
            d_346_v10_: _dafny.Set
            d_346_v10_ = _dafny.Set({p0})
            index53_ = default__.safeIndex(584, (d_343_v8_).length(0))
            (d_343_v8_)[index53_] = ((self.f42) * (self.f42)) - ((self.f42) * (((d_345_v9_).set(self.f42, default__.abs(len(d_346_v10_)))).cardinality))
        elif True:
            d_347_v11_: _dafny.Seq
            d_347_v11_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f41)])
            (globalState).f24 = (d_347_v11_)[default__.safeIndex(default__.safeDivisionInt((self).f41, self.f42), len(d_347_v11_))]
            (globalState).f14 = (self).f41
            d_348_v12_: _dafny.MultiSet
            d_348_v12_ = _dafny.MultiSet([self.f42, self.f42])
            d_349_v13_: _dafny.Array
            nw45_ = _dafny.Array(int(0), 5)
            d_349_v13_ = nw45_
            d_350_v14_: D2
            d_350_v14_ = D2_DC8(((self).f41) <= ((d_348_v12_).cardinality), d_349_v13_)
            source8_ = d_350_v14_
            if source8_.is_DC7:
                d_351___mcc_h1_ = source8_.cf13
                d_352___mcc_h2_ = source8_.cf14
                d_353___mcc_h3_ = source8_.cf15
                d_354_cf15_ = d_353___mcc_h3_
                d_355_cf14_ = d_352___mcc_h2_
                d_356_cf13_ = d_351___mcc_h1_
                d_357_v15_: _dafny.Set
                d_357_v15_ = _dafny.Set({p0, d_356_cf13_})
                d_358_v16_: C0
                nw46_ = C0()
                nw46_.ctor__(d_357_v15_)
                d_358_v16_ = nw46_
                d_359_v17_: _dafny.Set
                d_359_v17_ = _dafny.Set({_dafny.Map({(self).f41: p0})})
                d_360_v18_: _dafny.Map
                d_360_v18_ = _dafny.Map({d_359_v17_: default__.safeModuloInt((self).f41, len(d_357_v15_))})
                rhs45_ = default__.fm0((_dafny.SeqWithoutIsStrInference([d_355_cf14_, (d_325_v0_)[default__.safeIndex(self.f42, len(d_325_v0_))]])) <= ((d_325_v0_).set(default__.safeIndex(self.f42, len(d_325_v0_)), d_356_cf13_)), globalState)
                rhs46_ = self.f42
                rhs47_ = ((d_354_cf15_) + (_dafny.SeqWithoutIsStrInference([p1 for d_361_i5_ in range(default__.abs(598))]))) <= ((d_354_cf15_) + (_dafny.SeqWithoutIsStrInference([p1 for d_362_i6_ in range(default__.abs(580))])))
                rhs48_ = ((d_360_v18_)[d_359_v17_] if (d_359_v17_) in (d_360_v18_) else (self).f41)
                lhs39_ = globalState
                lhs40_ = globalState
                lhs41_ = globalState
                d_325_v0_ = rhs45_
                lhs39_.f24 = rhs46_
                lhs40_.f11 = rhs47_
                lhs41_.f14 = rhs48_
                d_363_v19_: str
                d_363_v19_ = _dafny.CodePoint('u')
                rhs49_ = _dafny.CodePoint('h')
                rhs50_ = (self).f41
                lhs42_ = globalState
                d_363_v19_ = rhs49_
                lhs42_.f19 = rhs50_
                (globalState).f27 = ((self).f41) != ((self).f41)
            elif source8_.is_DC8:
                d_364___mcc_h4_ = source8_.cf16
                d_365___mcc_h5_ = source8_.cf17
                d_366_cf17_ = d_365___mcc_h5_
                d_367_cf16_ = d_364___mcc_h4_
                d_368_v20_: _dafny.Set
                d_368_v20_ = _dafny.Set({p1})
                d_369_v21_: D4
                d_369_v21_ = D4_DC12(p0, d_368_v20_)
                d_370_v22_: _dafny.Set
                d_370_v22_ = _dafny.Set({d_369_v21_})
                rhs51_ = (d_370_v22_).ispropersubset(d_370_v22_)
                rhs52_ = (self).f41
                rhs53_ = self.f42
                rhs54_ = p0
                rhs55_ = p0
                lhs43_ = globalState
                lhs44_ = globalState
                lhs45_ = globalState
                lhs46_ = globalState
                d_367_cf16_ = rhs51_
                lhs43_.f24 = rhs52_
                lhs44_.f0 = rhs53_
                lhs45_.f1 = rhs54_
                lhs46_.f1 = rhs55_
                d_371_v23_: _dafny.Array
                def lambda16_(d_372_v12_):
                    def lambda17_(d_373_i7_):
                        return (d_372_v12_).issubset(d_372_v12_)

                    return lambda17_

                init9_ = lambda16_(d_348_v12_)
                nw47_ = _dafny.Array(None, 3)
                for i0_9_ in range(nw47_.length(0)):
                    nw47_[i0_9_] = init9_(i0_9_)
                d_371_v23_ = nw47_
                index54_ = default__.safeIndex(63, (d_371_v23_).length(0))
                (d_371_v23_)[index54_] = False
                d_374_v24_: _dafny.Seq
                d_374_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hh"))
                d_375_v25_: _dafny.Set
                d_375_v25_ = _dafny.Set({d_367_cf16_})
                d_376_v26_: _dafny.Map
                d_376_v26_ = _dafny.Map({p0: _dafny.CodePoint('i')})
                d_377_v27_: _dafny.Map
                d_377_v27_ = _dafny.Map({(self).f41: p0})
                index55_ = default__.safeIndex(63, (d_371_v23_).length(0))
                (d_371_v23_)[index55_] = (self.f42) == ((default__.fm2(d_374_v24_, d_375_v25_, d_376_v26_, p0, globalState) if d_367_cf16_ else len(d_377_v27_)))
                d_378_v28_: D2
                d_378_v28_ = D2_DC7(default__.fm4(globalState), p0, (d_374_v24_ if d_367_cf16_ else d_374_v24_))
                rhs56_ = default__.fm4(globalState)
                rhs57_ = d_378_v28_
                rhs58_ = d_371_v23_
                lhs47_ = globalState
                lhs47_.f1 = rhs56_
                d_378_v28_ = rhs57_
                d_371_v23_ = rhs58_
                d_379_v29_: _dafny.Array
                def lambda18_(d_380_v24_):
                    def lambda19_(d_381_i8_):
                        return (d_380_v24_) + (d_380_v24_)

                    return lambda19_

                init10_ = lambda18_(d_374_v24_)
                nw48_ = _dafny.Array(None, 6)
                for i0_10_ in range(nw48_.length(0)):
                    nw48_[i0_10_] = init10_(i0_10_)
                d_379_v29_ = nw48_
                index56_ = default__.safeIndex(585, (d_379_v29_).length(0))
                (d_379_v29_)[index56_] = _dafny.SeqWithoutIsStrInference([p1 for d_382_i9_ in range(default__.abs(523))])
                index57_ = default__.safeIndex(585, (d_379_v29_).length(0))
                (d_379_v29_)[index57_] = ((d_374_v24_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mlkpbgw")))) + (d_374_v24_)
            elif True:
                d_383___mcc_h6_ = source8_.cf12
                d_384_cf12_ = d_383___mcc_h6_
                (globalState).f11 = not(p0)
                rhs59_ = p0
                rhs60_ = d_347_v11_
                lhs48_ = globalState
                lhs48_.f27 = rhs59_
                d_347_v11_ = rhs60_
                (globalState).f1 = (d_325_v0_)[default__.safeIndex((self).f41, len(d_325_v0_))]
                d_385_v30_: _dafny.Set
                d_385_v30_ = _dafny.Set({-664})
                d_386_v31_: _dafny.Map
                d_386_v31_ = _dafny.Map({(len(default__.fm3(globalState)) if p0 else (self).f41): default__.safeModuloInt(self.f42, len(d_385_v30_))})
                d_386_v31_ = (d_386_v31_).set((self).f41, 787)
            d_387_v32_: _dafny.Map
            d_387_v32_ = _dafny.Map({self.f42: d_348_v12_})
            d_388_v33_: _dafny.Seq
            d_388_v33_ = _dafny.SeqWithoutIsStrInference([d_348_v12_, _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f41])), ((d_387_v32_)[self.f42] if (self.f42) in (d_387_v32_) else d_348_v12_), d_348_v12_, d_348_v12_])
            d_389_v34_: _dafny.Set
            d_389_v34_ = _dafny.Set({p0})
            d_390_v35_: _dafny.Map
            d_390_v35_ = _dafny.Map({(self.f42) < ((self).f41): (len(d_388_v33_)) * (len(d_389_v34_))})
            d_390_v35_ = (d_390_v35_) | ((d_390_v35_) | (d_390_v35_))
            (globalState).f24 = (d_347_v11_)[default__.safeIndex(-164, len(d_347_v11_))]
        if p0:
            d_391_v36_: _dafny.Array
            nw49_ = _dafny.Array(int(0), 27)
            d_391_v36_ = nw49_
            index58_ = default__.safeIndex(263, (d_391_v36_).length(0))
            (d_391_v36_)[index58_] = (self).f41
            d_392_v37_: _dafny.Seq
            d_392_v37_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ug"))
            index59_ = default__.safeIndex(263, (d_391_v36_).length(0))
            (d_391_v36_)[index59_] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "butbgpmp"))) + (d_392_v37_))
            d_393_v38_: _dafny.Array
            nw50_ = _dafny.Array(None, 7)
            nw50_[int(0)] = p0
            nw50_[int(1)] = p0
            nw50_[int(2)] = p0
            nw50_[int(3)] = p0
            nw50_[int(4)] = p0
            nw50_[int(5)] = p0
            nw50_[int(6)] = p0
            d_393_v38_ = nw50_
            default__.m0(d_393_v38_, d_391_v36_, len(default__.fm3(globalState)), (self).f41, globalState)
            d_394_v39_: str
            d_394_v39_ = _dafny.CodePoint('o')
            d_394_v39_ = p1
            d_395_v40_: _dafny.Map
            d_395_v40_ = _dafny.Map({(792) + (len(d_392_v37_)): default__.safeDivisionInt((d_391_v36_)[default__.safeIndex(263, (d_391_v36_).length(0))], self.f42)})
            index60_ = default__.safeIndex(263, (d_391_v36_).length(0))
            rhs61_ = (d_391_v36_)[default__.safeIndex(263, (d_391_v36_).length(0))]
            rhs62_ = ((d_395_v40_)[self.f42] if (self.f42) in (d_395_v40_) else (0) - ((self.f42) + (452)))
            rhs63_ = (d_391_v36_)[default__.safeIndex(263, (d_391_v36_).length(0))]
            rhs64_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ebsaabwkq"))
            rhs65_ = p0
            lhs49_ = d_391_v36_
            lhs50_ = default__.safeIndex(263, (d_391_v36_).length(0))
            lhs51_ = globalState
            lhs52_ = globalState
            lhs53_ = globalState
            lhs49_[lhs50_] = rhs61_
            lhs51_.f14 = rhs62_
            lhs52_.f0 = rhs63_
            d_392_v37_ = rhs64_
            lhs53_.f1 = rhs65_
            d_396_v41_: _dafny.Array
            nw51_ = _dafny.Array(_dafny.Array(None, 0), 17)
            d_396_v41_ = nw51_
            d_396_v41_ = d_396_v41_
        elif True:
            (self).f42 = (self).f41
            d_397_v42_: _dafny.Array
            nw52_ = _dafny.Array(False, 4)
            d_397_v42_ = nw52_
            d_398_v43_: _dafny.Map
            d_398_v43_ = _dafny.Map({(self).f41: 994})
            d_399_v44_: _dafny.MultiSet
            d_399_v44_ = _dafny.MultiSet([len(d_398_v43_)])
            index61_ = default__.safeIndex(315, (d_397_v42_).length(0))
            (d_397_v42_)[index61_] = not((d_399_v44_) == (d_399_v44_))
            index62_ = default__.safeIndex(315, (d_397_v42_).length(0))
            (d_397_v42_)[index62_] = (self.f42) > (self.f42)
            d_400_v45_: _dafny.Map
            d_400_v45_ = _dafny.Map({(d_397_v42_)[default__.safeIndex(315, (d_397_v42_).length(0))]: self.f42})
            d_401_v46_: _dafny.Set
            d_401_v46_ = _dafny.Set({_dafny.Map({p0: (self).f41}), d_400_v45_, d_400_v45_, d_400_v45_})
            d_402_v47_: _dafny.Seq
            d_402_v47_ = _dafny.SeqWithoutIsStrInference([len(d_401_v46_)])
            index63_ = default__.safeIndex(315, (d_397_v42_).length(0))
            rhs66_ = (d_402_v47_)[default__.safeIndex(len(d_402_v47_), len(d_402_v47_))]
            rhs67_ = p0
            lhs54_ = self
            lhs55_ = d_397_v42_
            lhs56_ = default__.safeIndex(315, (d_397_v42_).length(0))
            lhs54_.f42 = rhs66_
            lhs55_[lhs56_] = rhs67_
            (globalState).f14 = (self).f41
            d_403_v48_: _dafny.Seq
            d_403_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
            d_404_v49_: D1
            d_404_v49_ = D1_DC4(len(d_403_v48_), (d_397_v42_)[default__.safeIndex(315, (d_397_v42_).length(0))])
            d_405_v50_: _dafny.Array
            nw53_ = _dafny.Array(int(0), 11)
            d_405_v50_ = nw53_
            d_406_v51_: D2
            d_406_v51_ = D2_DC8((d_404_v49_).cf6, d_405_v50_)
            d_406_v51_ = d_406_v51_
        d_407_i10_: int
        d_407_i10_ = 0
        with _dafny.label("4"):
            while p0:
                with _dafny.c_label("4"):
                    if (d_407_i10_) >= (100):
                        raise _dafny.Break("4")
                    d_407_i10_ = (d_407_i10_) + (1)
                    d_408_v52_: _dafny.Array
                    nw54_ = _dafny.Array(int(0), 5)
                    d_408_v52_ = nw54_
                    index64_ = default__.safeIndex(655, (d_408_v52_).length(0))
                    (d_408_v52_)[index64_] = default__.safeDivisionInt(len(d_325_v0_), self.f42)
                    index65_ = default__.safeIndex(655, (d_408_v52_).length(0))
                    (d_408_v52_)[index65_] = (D0_DC1((0) - (self.f42))).cf1
                    d_409_v53_: _dafny.Array
                    nw55_ = _dafny.Array(False, 11)
                    d_409_v53_ = nw55_
                    index66_ = default__.safeIndex(786, (d_409_v53_).length(0))
                    (d_409_v53_)[index66_] = p0
                    index67_ = default__.safeIndex(786, (d_409_v53_).length(0))
                    (d_409_v53_)[index67_] = p0
                    d_410_v54_: _dafny.Seq
                    d_410_v54_ = _dafny.SeqWithoutIsStrInference([self.f42])
                    d_411_v55_: _dafny.Seq
                    d_411_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ulfbjea"))
                    d_412_v56_: _dafny.Map
                    d_412_v56_ = _dafny.Map({len((d_410_v54_).set(default__.safeIndex((self).f41, len(d_410_v54_)), len(default__.fm6((d_409_v53_)[default__.safeIndex(786, (d_409_v53_).length(0))], (d_408_v52_)[default__.safeIndex(655, (d_408_v52_).length(0))], d_411_v55_, globalState)))): self.f42})
                    d_413_v58_: _dafny.Seq
                    def iife77_():
                        coll39_ = _dafny.Map()
                        compr_39_: int
                        for compr_39_ in _dafny.IntegerRange(434, 190):
                            d_414_v57_: int = compr_39_
                            if ((434) <= (d_414_v57_)) and ((d_414_v57_) < (190)):
                                coll39_[default__.safeDivisionInt(d_414_v57_, self.f42)] = self.f42
                        return _dafny.Map(coll39_)
                    d_413_v58_ = _dafny.SeqWithoutIsStrInference([default__.fm21(_dafny.CodePoint('g'), globalState), (d_412_v56_) | (d_412_v56_), iife77_()
                    ])
                    d_415_v59_: _dafny.Map
                    d_415_v59_ = _dafny.Map({default__.fm4(globalState): self.f42})
                    index68_ = default__.safeIndex(655, (d_408_v52_).length(0))
                    rhs68_ = default__.safeModuloInt(len(d_415_v59_), (d_408_v52_)[default__.safeIndex(655, (d_408_v52_).length(0))])
                    rhs69_ = ((d_413_v58_) + (d_413_v58_)).set(default__.safeIndex(((self).f41) - ((d_408_v52_)[default__.safeIndex(655, (d_408_v52_).length(0))]), len((d_413_v58_) + (d_413_v58_))), (d_412_v56_).set((self).f41, 193))
                    rhs70_ = (self).f41
                    rhs71_ = (_dafny.SeqWithoutIsStrInference([p1 for d_416_i11_ in range(default__.abs(-297))])) + (((d_411_v55_) + (d_411_v55_)).set(default__.safeIndex(self.f42, len((d_411_v55_) + (d_411_v55_))), p1))
                    lhs57_ = d_408_v52_
                    lhs58_ = default__.safeIndex(655, (d_408_v52_).length(0))
                    r0 = rhs68_
                    d_413_v58_ = rhs69_
                    lhs57_[lhs58_] = rhs70_
                    d_411_v55_ = rhs71_
                    rhs72_ = p0
                    rhs73_ = p0
                    lhs59_ = globalState
                    lhs60_ = globalState
                    lhs59_.f11 = rhs72_
                    lhs60_.f11 = rhs73_
                    pass
            pass
        d_417_i12_: int
        d_417_i12_ = 0
        with _dafny.label("5"):
            while p0:
                with _dafny.c_label("5"):
                    if (d_417_i12_) >= (100):
                        raise _dafny.Break("5")
                    d_417_i12_ = (d_417_i12_) + (1)
                    (globalState).f14 = (self).f41
                    d_418_v60_: _dafny.Seq
                    d_418_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "skvlfa"))
                    d_418_v60_ = (d_418_v60_) + ((d_418_v60_) + (d_418_v60_))
                    d_419_v61_: _dafny.Set
                    d_419_v61_ = _dafny.Set({p0, p0})
                    d_420_v62_: _dafny.Map
                    d_420_v62_ = _dafny.Map({p0: p1})
                    d_421_v63_: _dafny.Map
                    d_421_v63_ = _dafny.Map({default__.fm2(d_418_v60_, d_419_v61_, d_420_v62_, p0, globalState): self.f42})
                    d_422_v64_: D0
                    d_422_v64_ = D0_DC0(d_421_v63_)
                    d_423_v65_: _dafny.MultiSet
                    d_423_v65_ = _dafny.MultiSet([p0, False])
                    d_424_v66_: _dafny.Map
                    d_424_v66_ = _dafny.Map({d_422_v64_: (p0) in (d_423_v65_)})
                    d_424_v66_ = d_424_v66_
                    (globalState).f27 = ((p1) not in (d_418_v60_) if p0 else p0)
                    pass
            pass
        d_425_i13_: int
        d_425_i13_ = 0
        with _dafny.label("6"):
            while p0:
                with _dafny.c_label("6"):
                    if (d_425_i13_) >= (100):
                        raise _dafny.Break("6")
                    d_425_i13_ = (d_425_i13_) + (1)
                    d_426_v67_: _dafny.Seq
                    d_426_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdvsmayoh"))
                    d_427_v68_: _dafny.Set
                    d_427_v68_ = _dafny.Set({p0, p0, p0, p0})
                    d_428_v69_: D1
                    d_428_v69_ = D1_DC5(p0, not(p0), default__.fm2(d_426_v67_, d_427_v68_, default__.fm18(p0, p0, globalState), not(p0), globalState), not(p0), p0)
                    (globalState).f14 = (d_428_v69_).cf9
                    (globalState).f24 = (self).f41
                    d_429_v70_: _dafny.Seq
                    d_429_v70_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0}), d_427_v68_])
                    d_430_v71_: C0
                    nw56_ = C0()
                    nw56_.ctor__((d_429_v70_)[default__.safeIndex(384, len(d_429_v70_))])
                    d_430_v71_ = nw56_
                    (globalState).f11 = not(p0)
                    pass
            pass
        (globalState).f27 = (self).fm7(not(p0), (self.f42) == (self.f42), globalState)
        r0 = (self.f42) - ((self).f41)
        return r0

    @property
    def f41(self):
        return self._f41

class C2(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        self._f40: _dafny.Set = _dafny.Set({})
        self._f39: _dafny.Map = _dafny.Map({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f39, f40, f28):
        (self)._f39 = f39
        (self)._f40 = f40
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        return not(((_dafny.Set({(self).f28})) - (_dafny.Set({(self).f28}))).isdisjoint(_dafny.Set({(self).f28, (self).f28})))

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('h')

    def fm17(self, p0, p1, globalState):
        return len((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False])])) + (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, True])])))

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_431_v0_: bool
        d_431_v0_ = False
        d_432_v1_: int
        d_432_v1_ = 170
        d_433_v2_: _dafny.MultiSet
        d_433_v2_ = _dafny.MultiSet([len((self).f40), (0) - (d_432_v1_)])
        d_434_v3_: _dafny.Map
        d_434_v3_ = _dafny.Map({d_431_v0_: False})
        d_435_v4_: _dafny.MultiSet
        d_435_v4_ = _dafny.MultiSet([d_434_v3_])
        d_436_v5_: _dafny.Array
        nw57_ = _dafny.Array(None, 29)
        nw57_[int(0)] = d_431_v0_
        nw57_[int(1)] = (d_431_v0_) and (d_431_v0_)
        nw57_[int(2)] = (d_433_v2_).issubset(d_433_v2_)
        nw57_[int(3)] = (d_433_v2_).ispropersubset(d_433_v2_)
        nw57_[int(4)] = d_431_v0_
        nw57_[int(5)] = d_431_v0_
        nw57_[int(6)] = (len(d_434_v3_)) == (d_432_v1_)
        nw57_[int(7)] = d_431_v0_
        nw57_[int(8)] = d_431_v0_
        nw57_[int(9)] = not (d_431_v0_) or (d_431_v0_)
        nw57_[int(10)] = (d_435_v4_).ispropersubset(d_435_v4_)
        nw57_[int(11)] = d_431_v0_
        nw57_[int(12)] = d_431_v0_
        nw57_[int(13)] = d_431_v0_
        nw57_[int(14)] = d_431_v0_
        nw57_[int(15)] = d_431_v0_
        nw57_[int(16)] = (d_431_v0_) == (d_431_v0_)
        nw57_[int(17)] = d_431_v0_
        nw57_[int(18)] = d_431_v0_
        nw57_[int(19)] = d_431_v0_
        nw57_[int(20)] = d_431_v0_
        nw57_[int(21)] = d_431_v0_
        nw57_[int(22)] = d_431_v0_
        nw57_[int(23)] = True
        nw57_[int(24)] = d_431_v0_
        nw57_[int(25)] = d_431_v0_
        nw57_[int(26)] = d_431_v0_
        nw57_[int(27)] = d_431_v0_
        nw57_[int(28)] = d_431_v0_
        d_436_v5_ = nw57_
        d_437_v6_: _dafny.Seq
        d_437_v6_ = _dafny.SeqWithoutIsStrInference([d_432_v1_, d_432_v1_])
        d_438_v7_: _dafny.Map
        d_438_v7_ = _dafny.Map({d_432_v1_: len(d_437_v6_)})
        index69_ = default__.safeIndex(480, (d_436_v5_).length(0))
        (d_436_v5_)[index69_] = ((_dafny.MultiSet([d_438_v7_, d_438_v7_, d_438_v7_])).cardinality) > ((self).fm17(_dafny.SeqWithoutIsStrInference([d_432_v1_]), d_432_v1_, globalState))
        index70_ = default__.safeIndex(480, (d_436_v5_).length(0))
        (d_436_v5_)[index70_] = not((d_437_v6_) != (d_437_v6_))
        if d_431_v0_:
            d_439_v8_: str
            d_439_v8_ = _dafny.CodePoint('v')
            d_440_v9_: _dafny.Seq
            d_440_v9_ = _dafny.SeqWithoutIsStrInference([d_439_v8_, d_439_v8_, d_439_v8_, d_439_v8_])
            d_441_v10_: D2
            d_441_v10_ = D2_DC7(d_431_v0_, d_431_v0_, d_440_v9_)
            d_442_v11_: _dafny.Set
            d_442_v11_ = _dafny.Set({d_441_v10_, d_441_v10_, d_441_v10_, d_441_v10_, D2_DC7((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], d_431_v0_, d_440_v9_)})
            d_443_v12_: _dafny.Map
            d_443_v12_ = _dafny.Map({d_432_v1_: d_431_v0_})
            rhs74_ = (d_442_v11_).isdisjoint((d_442_v11_).intersection(d_442_v11_))
            rhs75_ = d_443_v12_
            lhs61_ = globalState
            lhs62_ = globalState
            lhs61_.f11 = rhs74_
            lhs62_.f20 = rhs75_
            d_444_v13_: C1
            nw58_ = C1()
            nw58_.ctor__(d_432_v1_, d_432_v1_, (self).f28)
            d_444_v13_ = nw58_
            (globalState).f24 = ((0) - ((d_433_v2_).cardinality)) * (d_432_v1_)
            (globalState).f24 = (d_444_v13_).f41
            (globalState).f1 = d_431_v0_
        elif True:
            d_445_v14_: _dafny.Array
            nw59_ = _dafny.Array(_dafny.Map({}), 7)
            d_445_v14_ = nw59_
            d_446_v15_: _dafny.Array
            d_446_v15_ = d_445_v14_
            d_447_v16_: _dafny.Seq
            d_447_v16_ = _dafny.SeqWithoutIsStrInference([(d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]])
            d_448_v17_: _dafny.Map
            d_448_v17_ = _dafny.Map({d_447_v16_: d_432_v1_})
            rhs76_ = d_446_v15_
            rhs77_ = d_432_v1_
            rhs78_ = (d_448_v17_) | ((default__.fm22((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], d_432_v1_, d_432_v1_, globalState)) | ((self).f39))
            rhs79_ = default__.fm3(globalState)
            lhs63_ = globalState
            d_446_v15_ = rhs76_
            lhs63_.f0 = rhs77_
            d_448_v17_ = rhs78_
            r1 = rhs79_
            d_449_v18_: _dafny.Array
            nw60_ = _dafny.Array(_dafny.Map({}), 29)
            d_449_v18_ = nw60_
            d_450_v19_: _dafny.Seq
            d_450_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcqon"))
            d_451_v20_: _dafny.Map
            d_451_v20_ = _dafny.Map({d_450_v19_: default__.fm23(globalState)})
            index71_ = default__.safeIndex(119, (d_449_v18_).length(0))
            (d_449_v18_)[index71_] = (d_451_v20_) | (d_451_v20_)
            d_452_v22_: _dafny.Map
            d_452_v22_ = _dafny.Map({d_450_v19_: d_437_v6_})
            index72_ = default__.safeIndex(119, (d_449_v18_).length(0))
            def iife78_():
                coll40_ = _dafny.Map()
                compr_40_: _dafny.Seq
                for compr_40_ in (d_452_v22_).keys.Elements:
                    d_453_v21_: _dafny.Seq = compr_40_
                    if (d_453_v21_) in (d_452_v22_):
                        coll40_[d_453_v21_] = _dafny.Map({d_432_v1_: d_431_v0_})
                return _dafny.Map(coll40_)
            (d_449_v18_)[index72_] = (d_451_v20_) | (iife78_()
            )
            (globalState).f27 = not(((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] if (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] else d_431_v0_))
            d_454_v23_: _dafny.Map
            d_454_v23_ = _dafny.Map({d_437_v6_: default__.fm4(globalState)})
            d_455_v24_: _dafny.Set
            d_455_v24_ = _dafny.Set({not((d_431_v0_) == ((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))])), (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], (((d_454_v23_)[d_437_v6_] if (d_437_v6_) in (d_454_v23_) else d_431_v0_)) or (d_431_v0_), (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], d_431_v0_})
            d_455_v24_ = ((self).f40).intersection(d_455_v24_)
            d_456_v25_: _dafny.Set
            d_456_v25_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywxubw"))})
            d_456_v25_ = ((d_456_v25_).intersection(d_456_v25_)).intersection((d_456_v25_) - (d_456_v25_))
        d_457_v26_: D0
        d_457_v26_ = D0_DC0(d_438_v7_)
        pat_let_tv15_ = d_436_v5_
        pat_let_tv16_ = d_436_v5_
        pat_let_tv17_ = d_432_v1_
        def lambda20_(source9_):
            if source9_.is_DC1:
                d_458___mcc_h0_ = source9_.cf1
                d_459_cf1_ = d_458___mcc_h0_
                return d_459_cf1_
            elif source9_.is_DC2:
                d_460___mcc_h1_ = source9_.cf2
                d_461___mcc_h2_ = source9_.cf3
                d_462_cf3_ = d_461___mcc_h2_
                d_463_cf2_ = d_460___mcc_h1_
                return (0) - ((0) - (len(_dafny.SeqWithoutIsStrInference([(pat_let_tv16_)[default__.safeIndex(480, (pat_let_tv15_).length(0))]]))))
            elif True:
                d_464___mcc_h3_ = source9_.cf0
                d_465_cf0_ = d_464___mcc_h3_
                return pat_let_tv17_

        d_432_v1_ = lambda20_((d_457_v26_ if (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] else D0_DC0(d_438_v7_)))
        d_466_v27_: C1
        nw61_ = C1()
        nw61_.ctor__(d_432_v1_, d_432_v1_, (self).f28)
        d_466_v27_ = nw61_
        d_467_v28_: _dafny.Seq
        d_467_v28_ = _dafny.SeqWithoutIsStrInference([(d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]])
        d_468_v29_: D1
        d_468_v29_ = D1_DC4(len(d_467_v28_), (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))])
        hi0_ = (d_468_v29_).cf5
        for d_469_i0_ in range((d_466_v27_.f42 if (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] else d_432_v1_), hi0_):
            if (self).fm7(((d_434_v3_)[not((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))])] if (not((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))])) in (d_434_v3_) else d_431_v0_), (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], globalState):
                d_470_v30_: _dafny.Seq
                d_470_v30_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jurwunavv"))
                d_471_v31_: _dafny.Seq
                d_471_v31_ = _dafny.SeqWithoutIsStrInference([d_470_v30_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rdenx")), d_470_v30_, d_470_v30_, d_470_v30_])
                d_472_v32_: str
                d_472_v32_ = _dafny.CodePoint('q')
                d_473_v33_: _dafny.Map
                d_473_v33_ = _dafny.Map({d_431_v0_: d_472_v32_})
                d_474_v34_: _dafny.Set
                d_474_v34_ = _dafny.Set({d_470_v30_, d_470_v30_})
                (globalState).f1 = (default__.fm2((d_471_v31_)[default__.safeIndex(d_466_v27_.f42, len(d_471_v31_))], (self).f40, d_473_v33_, d_431_v0_, globalState)) <= (len((_dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xledywls")), d_470_v30_})) - (d_474_v34_)))
                (globalState).f27 = d_431_v0_
                (globalState).f19 = (d_437_v6_)[default__.safeIndex((d_433_v2_).cardinality, len(d_437_v6_))]
                (globalState).f14 = default__.safeDivisionInt((0) - ((d_466_v27_).f41), (d_466_v27_).f41)
                (globalState).f11 = d_431_v0_
            elif True:
                (globalState).f11 = (d_466_v27_.f42) <= ((d_466_v27_).f41)
                d_475_v35_: _dafny.Seq
                d_475_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "och"))
                d_476_v36_: _dafny.Array
                nw62_ = _dafny.Array(_dafny.CodePoint('D'), 28)
                d_476_v36_ = nw62_
                d_477_v37_: _dafny.Map
                d_477_v37_ = _dafny.Map({d_476_v36_: (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]})
                d_478_v38_: _dafny.Map
                d_478_v38_ = _dafny.Map({len(d_477_v37_): not(True)})
                d_479_v39_: str
                d_479_v39_ = _dafny.CodePoint('v')
                d_480_v40_: _dafny.Map
                d_480_v40_ = _dafny.Map({((d_478_v38_)[(d_466_v27_).f41] if ((d_466_v27_).f41) in (d_478_v38_) else d_431_v0_): d_479_v39_})
                d_481_v41_: _dafny.Set
                d_481_v41_ = _dafny.Set({(d_466_v27_).f41})
                d_482_v42_: _dafny.Array
                nw63_ = _dafny.Array(None, 21)
                nw63_[int(0)] = default__.fm2(d_475_v35_, (self).f40, d_480_v40_, d_431_v0_, globalState)
                nw63_[int(1)] = (d_466_v27_.f42 if (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] else len(d_475_v35_))
                nw63_[int(2)] = (0) - (((0) - ((d_466_v27_).f41)) + (463))
                nw63_[int(3)] = (d_466_v27_).f41
                nw63_[int(4)] = (d_466_v27_.f42) - (d_469_i0_)
                nw63_[int(5)] = len(d_475_v35_)
                nw63_[int(6)] = len(_dafny.SeqWithoutIsStrInference([d_431_v0_]))
                nw63_[int(7)] = d_466_v27_.f42
                nw63_[int(8)] = 171
                nw63_[int(9)] = default__.fm2(d_475_v35_, (self).f40, d_480_v40_, (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], globalState)
                nw63_[int(10)] = (d_468_v29_).cf5
                nw63_[int(11)] = (d_466_v27_).f41
                nw63_[int(12)] = d_466_v27_.f42
                nw63_[int(13)] = ((0) - (d_466_v27_.f42)) * (d_466_v27_.f42)
                nw63_[int(14)] = d_469_i0_
                nw63_[int(15)] = (0) - ((0) - (len((_dafny.Map({(0) - (len(d_475_v35_)): d_481_v41_})).set(((d_438_v7_)[(0) - ((d_466_v27_).f41)] if ((0) - ((d_466_v27_).f41)) in (d_438_v7_) else (d_466_v27_).f41), d_481_v41_))))
                nw63_[int(16)] = (d_466_v27_).f41
                nw63_[int(17)] = 313
                nw63_[int(18)] = -142
                nw63_[int(19)] = (self).fm17(d_437_v6_, (d_466_v27_).f41, globalState)
                nw63_[int(20)] = (d_466_v27_).f41
                d_482_v42_ = nw63_
                index73_ = default__.safeIndex(346, (d_482_v42_).length(0))
                (d_482_v42_)[index73_] = (d_432_v1_) + (d_432_v1_)
                index74_ = default__.safeIndex(346, (d_482_v42_).length(0))
                (d_482_v42_)[index74_] = default__.safeModuloInt((d_466_v27_).f41, d_466_v27_.f42)
                (globalState).f14 = d_432_v1_
                d_482_v42_ = d_482_v42_
                (globalState).f27 = (default__.fm4(globalState) if d_431_v0_ else True)
            d_483_v43_: _dafny.MultiSet
            d_483_v43_ = _dafny.MultiSet([d_436_v5_])
            d_483_v43_ = ((_dafny.MultiSet([d_436_v5_])) - (d_483_v43_)) - (d_483_v43_)
            d_484_v44_: _dafny.Array
            nw64_ = _dafny.Array(_dafny.Map({}), 5)
            d_484_v44_ = nw64_
            d_484_v44_ = d_484_v44_
            d_485_v45_: _dafny.Array
            def lambda21_(d_486_v6_):
                def lambda22_(d_487_i1_):
                    return (d_486_v6_) + (d_486_v6_)

                return lambda22_

            init11_ = lambda21_(d_437_v6_)
            nw65_ = _dafny.Array(None, 29)
            for i0_11_ in range(nw65_.length(0)):
                nw65_[i0_11_] = init11_(i0_11_)
            d_485_v45_ = nw65_
            index75_ = default__.safeIndex(599, (d_485_v45_).length(0))
            (d_485_v45_)[index75_] = (d_437_v6_) + (d_437_v6_)
            d_488_v46_: _dafny.Seq
            d_488_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vhb"))
            index76_ = default__.safeIndex(599, (d_485_v45_).length(0))
            index77_ = default__.safeIndex(480, (d_436_v5_).length(0))
            rhs80_ = d_431_v0_
            rhs81_ = ((_dafny.SeqWithoutIsStrInference([-457 for d_489_i2_ in range(default__.abs(642))])) + (_dafny.SeqWithoutIsStrInference([-703 for d_490_i3_ in range(default__.abs(404))]))) + (d_437_v6_)
            rhs82_ = (default__.fm3(globalState)) != (d_488_v46_)
            lhs64_ = globalState
            lhs65_ = d_485_v45_
            lhs66_ = default__.safeIndex(599, (d_485_v45_).length(0))
            lhs67_ = d_436_v5_
            lhs68_ = default__.safeIndex(480, (d_436_v5_).length(0))
            lhs64_.f27 = rhs80_
            lhs65_[lhs66_] = rhs81_
            lhs67_[lhs68_] = rhs82_
        d_491_v47_: str
        d_491_v47_ = _dafny.CodePoint('m')
        d_492_v48_: _dafny.Set
        d_492_v48_ = _dafny.Set({d_491_v47_})
        d_493_v49_: D4
        d_493_v49_ = D4_DC12(d_431_v0_, d_492_v48_)
        if (d_493_v49_).cf21:
            d_494_v51_: D4
            def iife79_():
                coll41_ = _dafny.Set()
                compr_41_: str
                for compr_41_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])).Elements:
                    d_495_v50_: str = compr_41_
                    if (d_495_v50_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x')])):
                        coll41_ = coll41_.union(_dafny.Set([d_495_v50_]))
                return _dafny.Set(coll41_)
            d_494_v51_ = D4_DC12(d_431_v0_, iife79_()
)
            d_496_v52_: D4
            d_496_v52_ = D4_DC13(d_494_v51_)
            rhs83_ = (default__.safeModuloInt(d_432_v1_, d_466_v27_.f42)) - ((d_466_v27_.f42) + (d_466_v27_.f42))
            rhs84_ = d_496_v52_
            rhs85_ = (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]
            lhs69_ = globalState
            lhs70_ = globalState
            lhs69_.f14 = rhs83_
            d_496_v52_ = rhs84_
            lhs70_.f27 = rhs85_
            d_497_v53_: _dafny.Array
            nw66_ = _dafny.Array(int(0), 21)
            d_497_v53_ = nw66_
            rhs86_ = d_466_v27_.f42
            rhs87_ = d_497_v53_
            lhs71_ = d_466_v27_
            lhs71_.f42 = rhs86_
            d_497_v53_ = rhs87_
            d_498_v54_: _dafny.Seq
            d_498_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtdv"))
            d_499_v55_: C1
            nw67_ = C1()
            nw67_.ctor__(((d_466_v27_).f41) - (d_466_v27_.f42), len(d_498_v54_), (self).f28)
            d_499_v55_ = nw67_
            pat_let_tv18_ = d_436_v5_
            pat_let_tv19_ = d_436_v5_
            def iife80_(_pat_let19_0):
                def iife81_(d_500_dt__update__tmp_h0_):
                    def iife82_(_pat_let20_0):
                        def iife83_(d_501_dt__update_hcf6_h0_):
                            return D1_DC4((d_500_dt__update__tmp_h0_).cf5, d_501_dt__update_hcf6_h0_)
                        return iife83_(_pat_let20_0)
                    return iife82_((pat_let_tv19_)[default__.safeIndex(480, (pat_let_tv18_).length(0))])
                return iife81_(_pat_let19_0)
            (globalState).f19 = (0) - (len(default__.fm0((iife80_(d_468_v29_)).cf6, globalState)))
            (globalState).f27 = d_431_v0_
        elif True:
            d_502_v56_: _dafny.Map
            d_502_v56_ = _dafny.Map({(d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]: d_466_v27_.f42})
            d_503_v57_: _dafny.Map
            d_503_v57_ = _dafny.Map({d_431_v0_: d_491_v47_})
            d_504_v58_: _dafny.Map
            d_504_v58_ = _dafny.Map({d_466_v27_.f42: (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]})
            d_505_v59_: _dafny.Set
            d_505_v59_ = _dafny.Set({d_504_v58_, d_504_v58_, d_504_v58_})
            d_506_v60_: _dafny.Array
            nw68_ = _dafny.Array(None, 22)
            nw68_[int(0)] = default__.safeDivisionInt((0) - (d_432_v1_), (d_466_v27_).f41)
            nw68_[int(1)] = ((d_466_v27_).f41) - ((d_466_v27_).f41)
            nw68_[int(2)] = len((d_467_v28_) + (d_467_v28_))
            nw68_[int(3)] = (d_466_v27_).f41
            nw68_[int(4)] = -910
            nw68_[int(5)] = d_466_v27_.f42
            nw68_[int(6)] = (((d_502_v56_)[(d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]] if ((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]) in (d_502_v56_) else default__.fm2(default__.fm3(globalState), (self).f40, d_503_v57_, d_431_v0_, globalState))) - (417)
            nw68_[int(7)] = d_466_v27_.f42
            nw68_[int(8)] = len(_dafny.SeqWithoutIsStrInference([d_431_v0_, (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], not((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))])]))
            nw68_[int(9)] = (d_466_v27_).f41
            nw68_[int(10)] = (d_466_v27_).f41
            nw68_[int(11)] = d_466_v27_.f42
            nw68_[int(12)] = (d_466_v27_.f42) * ((d_466_v27_).f41)
            nw68_[int(13)] = d_432_v1_
            nw68_[int(14)] = d_466_v27_.f42
            nw68_[int(15)] = len(d_505_v59_)
            nw68_[int(16)] = (d_466_v27_).f41
            nw68_[int(17)] = d_432_v1_
            nw68_[int(18)] = (d_466_v27_).f41
            nw68_[int(19)] = ((d_466_v27_).f41 if d_431_v0_ else (d_466_v27_).f41)
            nw68_[int(20)] = d_466_v27_.f42
            nw68_[int(21)] = ((d_466_v27_).f41) - ((d_466_v27_).f41)
            d_506_v60_ = nw68_
            index78_ = default__.safeIndex(895, (d_506_v60_).length(0))
            (d_506_v60_)[index78_] = len(d_502_v56_)
            index79_ = default__.safeIndex(895, (d_506_v60_).length(0))
            (d_506_v60_)[index79_] = d_466_v27_.f42
            d_507_v61_: D2
            d_507_v61_ = D2_DC6(d_436_v5_)
            d_436_v5_ = (d_507_v61_).cf12
            (d_466_v27_).f42 = (d_466_v27_).f41
            (globalState).f27 = (not(d_431_v0_)) == (not((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]))
            d_508_v62_: _dafny.MultiSet
            d_508_v62_ = _dafny.MultiSet([d_491_v47_])
            if (d_508_v62_).isdisjoint(d_508_v62_):
                (globalState).f1 = (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]
                d_509_v63_: C1
                nw69_ = C1()
                nw69_.ctor__(len((d_504_v58_) | (d_504_v58_)), default__.safeDivisionInt((d_466_v27_).f41, (d_506_v60_)[default__.safeIndex(895, (d_506_v60_).length(0))]), (self).f28)
                d_509_v63_ = nw69_
                (globalState).f16 = ((d_467_v28_) + (d_467_v28_)) + (_dafny.SeqWithoutIsStrInference([(d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))]]))
                d_510_v64_: _dafny.Array
                nw70_ = _dafny.Array(None, 26)
                d_510_v64_ = nw70_
                d_511_v65_: C0
                nw71_ = C0()
                nw71_.ctor__(_dafny.Set({d_431_v0_}))
                d_511_v65_ = nw71_
                index80_ = default__.safeIndex(673, (d_510_v64_).length(0))
                (d_510_v64_)[index80_] = d_511_v65_
                d_512_v66_: D2
                d_512_v66_ = D2_DC8((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], d_506_v60_)
                d_513_v67_: _dafny.Array
                nw72_ = _dafny.Array(None, 15)
                nw72_[int(0)] = d_506_v60_
                nw72_[int(1)] = d_506_v60_
                nw72_[int(2)] = d_506_v60_
                nw72_[int(3)] = d_506_v60_
                nw72_[int(4)] = d_506_v60_
                nw72_[int(5)] = d_506_v60_
                nw72_[int(6)] = d_506_v60_
                nw72_[int(7)] = d_506_v60_
                nw72_[int(8)] = d_506_v60_
                nw72_[int(9)] = d_506_v60_
                nw72_[int(10)] = d_506_v60_
                nw72_[int(11)] = d_506_v60_
                nw72_[int(12)] = (d_512_v66_).cf17
                nw72_[int(13)] = d_506_v60_
                nw72_[int(14)] = d_506_v60_
                d_513_v67_ = nw72_
                index81_ = default__.safeIndex(112, (d_513_v67_).length(0))
                (d_513_v67_)[index81_] = d_506_v60_
                d_514_v68_: _dafny.Set
                d_514_v68_ = _dafny.Set({(d_466_v27_).f41})
                d_515_v69_: _dafny.Seq
                d_515_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rh"))
                index82_ = default__.safeIndex(673, (d_510_v64_).length(0))
                index83_ = default__.safeIndex(112, (d_513_v67_).length(0))
                index84_ = default__.safeIndex(895, (d_506_v60_).length(0))
                rhs88_ = d_511_v65_
                rhs89_ = d_506_v60_
                rhs90_ = False
                rhs91_ = (d_431_v0_ if (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] else (d_514_v68_).isdisjoint(d_514_v68_))
                rhs92_ = len(d_515_v69_)
                lhs72_ = d_510_v64_
                lhs73_ = default__.safeIndex(673, (d_510_v64_).length(0))
                lhs74_ = d_513_v67_
                lhs75_ = default__.safeIndex(112, (d_513_v67_).length(0))
                lhs76_ = globalState
                lhs77_ = globalState
                lhs78_ = d_506_v60_
                lhs79_ = default__.safeIndex(895, (d_506_v60_).length(0))
                lhs72_[lhs73_] = rhs88_
                lhs74_[lhs75_] = rhs89_
                lhs76_.f11 = rhs90_
                lhs77_.f1 = rhs91_
                lhs78_[lhs79_] = rhs92_
                (d_466_v27_).f42 = d_509_v63_.f42
            elif True:
                d_516_v70_: _dafny.Map
                d_516_v70_ = _dafny.Map({(d_466_v27_).f41: d_491_v47_})
                (globalState).f14 = len((d_516_v70_).set(((d_466_v27_).f41) + (len(d_437_v6_)), d_491_v47_))
                d_517_v71_: C1
                nw73_ = C1()
                nw73_.ctor__((d_466_v27_).f41, len((d_467_v28_) + (d_467_v28_)), (self).f28)
                d_517_v71_ = nw73_
                d_436_v5_ = d_436_v5_
                index85_ = default__.safeIndex(480, (d_436_v5_).length(0))
                (d_436_v5_)[index85_] = not((d_467_v28_)[default__.safeIndex(d_517_v71_.f42, len(d_467_v28_))])
                d_518_v72_: _dafny.Array
                nw74_ = _dafny.Array(_dafny.Seq({}), 28)
                d_518_v72_ = nw74_
                d_519_v73_: _dafny.Array
                nw75_ = _dafny.Array(None, 13)
                nw75_[int(0)] = d_518_v72_
                nw75_[int(1)] = d_518_v72_
                nw75_[int(2)] = d_518_v72_
                nw75_[int(3)] = d_518_v72_
                nw75_[int(4)] = d_518_v72_
                nw75_[int(5)] = d_518_v72_
                nw75_[int(6)] = (d_518_v72_ if (d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))] else d_518_v72_)
                nw75_[int(7)] = d_518_v72_
                nw75_[int(8)] = d_518_v72_
                nw75_[int(9)] = d_518_v72_
                nw75_[int(10)] = d_518_v72_
                nw75_[int(11)] = d_518_v72_
                nw75_[int(12)] = d_518_v72_
                d_519_v73_ = nw75_
                index86_ = default__.safeIndex(24, (d_519_v73_).length(0))
                (d_519_v73_)[index86_] = d_518_v72_
                index87_ = default__.safeIndex(24, (d_519_v73_).length(0))
                (d_519_v73_)[index87_] = d_518_v72_
        d_520_v74_: _dafny.Array
        nw76_ = _dafny.Array(int(0), 20)
        d_520_v74_ = nw76_
        d_521_v75_: _dafny.Array
        nw77_ = _dafny.Array(None, 9)
        nw77_[int(0)] = d_466_v27_.f42
        nw77_[int(1)] = d_466_v27_.f42
        nw77_[int(2)] = d_466_v27_.f42
        nw77_[int(3)] = len(_dafny.Map({d_520_v74_: d_466_v27_.f42}))
        nw77_[int(4)] = (d_466_v27_).f41
        nw77_[int(5)] = (d_466_v27_).f41
        nw77_[int(6)] = d_432_v1_
        nw77_[int(7)] = (d_466_v27_).f41
        nw77_[int(8)] = (d_466_v27_).f41
        d_521_v75_ = nw77_
        r0 = D2_DC8((d_436_v5_)[default__.safeIndex(480, (d_436_v5_).length(0))], d_520_v74_)
        d_522_v76_: _dafny.Seq
        d_522_v76_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nh"))
        r1 = d_522_v76_
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_523_v0_: int
        d_523_v0_ = 578
        hi1_ = (0) - (d_523_v0_)
        for d_524_i0_ in range(default__.safeModuloInt(d_523_v0_, d_523_v0_), hi1_):
            if (p0 if p0 else (-414) <= (d_524_i0_)):
                d_525_v1_: _dafny.Seq
                d_525_v1_ = _dafny.SeqWithoutIsStrInference([p0])
                d_526_v2_: _dafny.Map
                d_526_v2_ = _dafny.Map({d_524_i0_: default__.fm4(globalState)})
                d_527_v3_: _dafny.Map
                d_527_v3_ = _dafny.Map({d_525_v1_: (d_526_v2_) == (d_526_v2_)})
                d_527_v3_ = (d_527_v3_).set(d_525_v1_, p0)
                d_528_v4_: _dafny.Array
                nw78_ = _dafny.Array(False, 17)
                d_528_v4_ = nw78_
                index88_ = default__.safeIndex(175, (d_528_v4_).length(0))
                (d_528_v4_)[index88_] = ((self).f40).isdisjoint((self).f40)
                d_529_v5_: _dafny.Array
                def lambda23_(d_530_i0_):
                    def lambda24_(d_531_i1_):
                        return (d_531_i1_) * (d_530_i0_)

                    return lambda24_

                init12_ = lambda23_(d_524_i0_)
                nw79_ = _dafny.Array(None, 17)
                for i0_12_ in range(nw79_.length(0)):
                    nw79_[i0_12_] = init12_(i0_12_)
                d_529_v5_ = nw79_
                d_532_v6_: _dafny.MultiSet
                d_532_v6_ = _dafny.MultiSet([d_529_v5_])
                index89_ = default__.safeIndex(175, (d_528_v4_).length(0))
                (d_528_v4_)[index89_] = not((d_532_v6_).issubset(d_532_v6_))
                d_533_v7_: _dafny.Array
                nw80_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_533_v7_ = nw80_
                d_534_v8_: _dafny.Map
                d_534_v8_ = _dafny.Map({p1: d_533_v7_})
                d_535_v9_: _dafny.Seq
                d_535_v9_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cujvw"))
                d_536_v10_: _dafny.Map
                d_536_v10_ = _dafny.Map({((d_534_v8_)[p1] if (p1) in (d_534_v8_) else d_533_v7_): (d_535_v9_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfoh")))})
                d_536_v10_ = (d_536_v10_).set(d_533_v7_, d_535_v9_)
                d_537_v11_: str
                d_537_v11_ = _dafny.CodePoint('f')
                d_538_v12_: D4
                d_538_v12_ = D4_DC10(d_537_v11_)
                d_537_v11_ = (d_538_v12_).cf19
                (globalState).f24 = d_523_v0_
            elif True:
                d_539_v13_: _dafny.Array
                nw81_ = _dafny.Array(int(0), 2)
                d_539_v13_ = nw81_
                d_540_v14_: _dafny.Map
                d_540_v14_ = _dafny.Map({p0: d_539_v13_})
                d_541_v15_: _dafny.Map
                d_541_v15_ = _dafny.Map({d_524_i0_: d_523_v0_})
                d_542_v16_: _dafny.Seq
                d_542_v16_ = _dafny.SeqWithoutIsStrInference([d_523_v0_, d_523_v0_])
                d_543_v17_: C1
                nw82_ = C1()
                nw82_.ctor__((d_523_v0_) + (len(d_540_v14_)), ((d_541_v15_)[(self).fm17(d_542_v16_, d_524_i0_, globalState)] if ((self).fm17(d_542_v16_, d_524_i0_, globalState)) in (d_541_v15_) else d_523_v0_), (self).f28)
                d_543_v17_ = nw82_
                d_544_v18_: _dafny.Array
                nw83_ = _dafny.Array(_dafny.MultiSet({}), 23)
                d_544_v18_ = nw83_
                d_545_v19_: _dafny.MultiSet
                d_545_v19_ = _dafny.MultiSet([(d_543_v17_).f41, d_523_v0_])
                index90_ = default__.safeIndex(748, (d_544_v18_).length(0))
                (d_544_v18_)[index90_] = d_545_v19_
                d_546_v20_: _dafny.Array
                nw84_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_546_v20_ = nw84_
                index91_ = default__.safeIndex(982, (d_546_v20_).length(0))
                (d_546_v20_)[index91_] = p1
                d_547_v21_: _dafny.Array
                nw85_ = _dafny.Array(False, 7)
                d_547_v21_ = nw85_
                index92_ = default__.safeIndex(682, (d_547_v21_).length(0))
                (d_547_v21_)[index92_] = p0
                d_548_v22_: _dafny.Seq
                d_548_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aulkhrkor"))
                d_549_v23_: _dafny.Map
                d_549_v23_ = _dafny.Map({p0: p1})
                index93_ = default__.safeIndex(748, (d_544_v18_).length(0))
                index94_ = default__.safeIndex(982, (d_546_v20_).length(0))
                index95_ = default__.safeIndex(682, (d_547_v21_).length(0))
                rhs93_ = d_542_v16_
                rhs94_ = ((_dafny.MultiSet([(d_543_v17_).f41, default__.fm2((d_548_v22_).set(default__.safeIndex(d_523_v0_, len(d_548_v22_)), p1), (self).f40, d_549_v23_, p0, globalState)])) - (d_545_v19_)) - (_dafny.MultiSet([d_523_v0_, (d_543_v17_).f41, (d_543_v17_).f41, (0) - (d_523_v0_)]))
                rhs95_ = d_542_v16_
                rhs96_ = p1
                rhs97_ = False
                lhs80_ = d_544_v18_
                lhs81_ = default__.safeIndex(748, (d_544_v18_).length(0))
                lhs82_ = d_546_v20_
                lhs83_ = default__.safeIndex(982, (d_546_v20_).length(0))
                lhs84_ = d_547_v21_
                lhs85_ = default__.safeIndex(682, (d_547_v21_).length(0))
                d_542_v16_ = rhs93_
                lhs80_[lhs81_] = rhs94_
                d_542_v16_ = rhs95_
                lhs82_[lhs83_] = rhs96_
                lhs84_[lhs85_] = rhs97_
                (globalState).f24 = default__.fm2(_dafny.SeqWithoutIsStrInference([(d_548_v22_)[default__.safeIndex((0) - ((d_543_v17_).f41), len(d_548_v22_))] for d_550_i2_ in range(default__.abs(618))]), (self).f40, d_549_v23_, not(p0), globalState)
                index96_ = default__.safeIndex(682, (d_547_v21_).length(0))
                (d_547_v21_)[index96_] = True
                d_551_v24_: D1
                d_551_v24_ = D1_DC3((-797) > (d_543_v17_.f42))
                d_551_v24_ = d_551_v24_
            d_552_v25_: _dafny.Array
            nw86_ = _dafny.Array(False, 11)
            d_552_v25_ = nw86_
            d_553_v26_: _dafny.Seq
            d_553_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lrtsg"))
            d_554_v27_: _dafny.Map
            d_554_v27_ = _dafny.Map({(0) - (d_523_v0_): p0})
            d_555_v28_: _dafny.Map
            d_555_v28_ = _dafny.Map({p0: p1})
            index97_ = default__.safeIndex(539, (d_552_v25_).length(0))
            (d_552_v25_)[index97_] = (d_523_v0_) >= (default__.fm2(d_553_v26_, _dafny.Set({((d_554_v27_)[d_523_v0_] if (d_523_v0_) in (d_554_v27_) else p0)}), d_555_v28_, p0, globalState))
            index98_ = default__.safeIndex(539, (d_552_v25_).length(0))
            (d_552_v25_)[index98_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "trt"))) == (d_553_v26_)
            d_556_v29_: _dafny.Map
            d_556_v29_ = _dafny.Map({p0: d_524_i0_})
            d_556_v29_ = (d_556_v29_).set(False, (0) - (d_523_v0_))
            d_557_v30_: _dafny.Array
            nw87_ = _dafny.Array(int(0), 14)
            d_557_v30_ = nw87_
            d_557_v30_ = d_557_v30_
        if p0:
            d_558_v31_: str
            d_558_v31_ = _dafny.CodePoint('d')
            d_558_v31_ = d_558_v31_
            d_559_v32_: _dafny.Seq
            d_559_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xdlri"))
            d_560_v33_: _dafny.Seq
            d_560_v33_ = _dafny.SeqWithoutIsStrInference([p0])
            d_561_v34_: _dafny.Map
            d_561_v34_ = _dafny.Map({default__.fm2(d_559_v32_, default__.fm24((d_560_v33_).set(default__.safeIndex(d_523_v0_, len(d_560_v33_)), p0), p0, globalState), _dafny.Map({True: d_558_v31_}), p0, globalState): (d_523_v0_) == (d_523_v0_)})
            (globalState).f20 = d_561_v34_
            d_562_v35_: _dafny.Array
            nw88_ = _dafny.Array(D1.default()(), 19)
            d_562_v35_ = nw88_
            index99_ = default__.safeIndex(700, (d_562_v35_).length(0))
            (d_562_v35_)[index99_] = D1_DC4(d_523_v0_, p0)
            d_563_v36_: _dafny.Seq
            d_563_v36_ = _dafny.SeqWithoutIsStrInference([d_523_v0_, d_523_v0_, d_523_v0_, d_523_v0_])
            index100_ = default__.safeIndex(700, (d_562_v35_).length(0))
            (d_562_v35_)[index100_] = D1_DC4((d_563_v36_)[default__.safeIndex(d_523_v0_, len(d_563_v36_))], not (p0) or (p0))
            (globalState).f27 = default__.fm4(globalState)
            d_564_v37_: _dafny.Seq
            d_564_v37_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0, p0]), d_560_v33_, d_560_v33_, d_560_v33_, d_560_v33_])
            (globalState).f27 = not((self).fm7(p0, ((d_564_v37_)[default__.safeIndex(d_523_v0_, len(d_564_v37_))]) == (d_560_v33_), globalState))
        elif True:
            (globalState).f0 = d_523_v0_
            d_565_v38_: _dafny.Array
            nw89_ = _dafny.Array(None, 3)
            nw89_[int(0)] = p0
            nw89_[int(1)] = (p0) == (p0)
            nw89_[int(2)] = p0
            d_565_v38_ = nw89_
            index101_ = default__.safeIndex(92, (d_565_v38_).length(0))
            (d_565_v38_)[index101_] = p0
            d_566_v39_: _dafny.Map
            d_566_v39_ = _dafny.Map({-247: d_523_v0_})
            index102_ = default__.safeIndex(92, (d_565_v38_).length(0))
            (d_565_v38_)[index102_] = ((d_523_v0_) + ((0) - (len(_dafny.SeqWithoutIsStrInference([d_566_v39_]))))) <= ((0) - (d_523_v0_))
            d_567_v40_: _dafny.Seq
            d_567_v40_ = _dafny.SeqWithoutIsStrInference([d_523_v0_])
            d_568_v41_: _dafny.Array
            nw90_ = _dafny.Array(None, 17)
            nw90_[int(0)] = 395
            nw90_[int(1)] = d_523_v0_
            nw90_[int(2)] = d_523_v0_
            nw90_[int(3)] = d_523_v0_
            nw90_[int(4)] = d_523_v0_
            nw90_[int(5)] = d_523_v0_
            nw90_[int(6)] = d_523_v0_
            nw90_[int(7)] = d_523_v0_
            nw90_[int(8)] = d_523_v0_
            nw90_[int(9)] = d_523_v0_
            nw90_[int(10)] = d_523_v0_
            nw90_[int(11)] = d_523_v0_
            nw90_[int(12)] = d_523_v0_
            nw90_[int(13)] = 797
            nw90_[int(14)] = d_523_v0_
            nw90_[int(15)] = len(d_567_v40_)
            nw90_[int(16)] = -886
            d_568_v41_ = nw90_
            d_569_v42_: _dafny.Seq
            d_569_v42_ = _dafny.SeqWithoutIsStrInference([d_568_v41_])
            r0 = (len(d_569_v42_) if (d_565_v38_)[default__.safeIndex(92, (d_565_v38_).length(0))] else default__.safeModuloInt(d_523_v0_, d_523_v0_))
            d_570_v43_: C0
            nw91_ = C0()
            nw91_.ctor__((self).f40)
            d_570_v43_ = nw91_
            d_571_v44_: _dafny.Seq
            d_571_v44_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "neqjhkexi"))
            d_572_v45_: D1
            d_572_v45_ = D1_DC4((d_567_v40_)[default__.safeIndex(d_523_v0_, len(d_567_v40_))], not((d_565_v38_)[default__.safeIndex(92, (d_565_v38_).length(0))]))
            d_573_v46_: int
            d_574_v47_: bool
            out0_: int
            out1_: bool
            out0_, out1_ = (self).m15(p0, (636) - ((self).fm17(_dafny.SeqWithoutIsStrInference([len(d_571_v44_), 613, d_523_v0_, d_523_v0_, d_523_v0_]), d_523_v0_, globalState)), d_572_v45_, (d_565_v38_)[default__.safeIndex(92, (d_565_v38_).length(0))], globalState)
            d_573_v46_ = out0_
            d_574_v47_ = out1_
        d_575_v48_: _dafny.Array
        nw92_ = _dafny.Array(_dafny.Seq({}), 13)
        d_575_v48_ = nw92_
        d_576_v49_: _dafny.Seq
        d_576_v49_ = _dafny.SeqWithoutIsStrInference([d_523_v0_])
        index103_ = default__.safeIndex(89, (d_575_v48_).length(0))
        (d_575_v48_)[index103_] = (_dafny.SeqWithoutIsStrInference([(0) - (d_523_v0_)])) + (d_576_v49_)
        d_577_v50_: _dafny.Set
        d_577_v50_ = _dafny.Set({d_523_v0_})
        d_578_v51_: _dafny.MultiSet
        d_578_v51_ = _dafny.MultiSet([p0, (_dafny.Set({d_523_v0_, 524, d_523_v0_})).issubset(d_577_v50_)])
        d_579_v52_: _dafny.Array
        nw93_ = _dafny.Array(False, 19)
        d_579_v52_ = nw93_
        index104_ = default__.safeIndex(429, (d_579_v52_).length(0))
        (d_579_v52_)[index104_] = p0
        d_580_v53_: _dafny.Seq
        d_580_v53_ = _dafny.SeqWithoutIsStrInference([p0, p0, p0])
        d_581_v54_: _dafny.Map
        d_581_v54_ = _dafny.Map({p0: d_580_v53_})
        index105_ = default__.safeIndex(89, (d_575_v48_).length(0))
        index106_ = default__.safeIndex(429, (d_579_v52_).length(0))
        rhs98_ = (p0) and (p0)
        rhs99_ = d_576_v49_
        rhs100_ = (default__.fm25((default__.fm25(d_523_v0_, p0, d_580_v53_, d_523_v0_, globalState)).cardinality, p0, ((d_581_v54_)[(d_580_v53_)[default__.safeIndex(d_523_v0_, len(d_580_v53_))]] if ((d_580_v53_)[default__.safeIndex(d_523_v0_, len(d_580_v53_))]) in (d_581_v54_) else _dafny.SeqWithoutIsStrInference([p0])), d_523_v0_, globalState) if not((not(p0)) == (p0)) else d_578_v51_)
        rhs101_ = p0
        rhs102_ = p0
        lhs86_ = globalState
        lhs87_ = d_575_v48_
        lhs88_ = default__.safeIndex(89, (d_575_v48_).length(0))
        lhs89_ = globalState
        lhs90_ = d_579_v52_
        lhs91_ = default__.safeIndex(429, (d_579_v52_).length(0))
        lhs86_.f27 = rhs98_
        lhs87_[lhs88_] = rhs99_
        d_578_v51_ = rhs100_
        lhs89_.f1 = rhs101_
        lhs90_[lhs91_] = rhs102_
        d_582_v55_: _dafny.Seq
        d_582_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nldq"))
        hi2_ = d_523_v0_
        for d_583_i3_ in range(len(d_582_v55_), hi2_):
            (globalState).f19 = 292
            d_584_v56_: _dafny.Array
            nw94_ = _dafny.Array(None, 7)
            nw94_[int(0)] = p1
            nw94_[int(1)] = p1
            nw94_[int(2)] = p1
            nw94_[int(3)] = p1
            nw94_[int(4)] = p1
            nw94_[int(5)] = p1
            nw94_[int(6)] = p1
            d_584_v56_ = nw94_
            index107_ = default__.safeIndex(574, (d_584_v56_).length(0))
            (d_584_v56_)[index107_] = _dafny.CodePoint('l')
            index108_ = default__.safeIndex(574, (d_584_v56_).length(0))
            (d_584_v56_)[index108_] = p1
            d_585_v57_: _dafny.Map
            d_585_v57_ = _dafny.Map({p0: d_583_i3_})
            d_577_v50_ = _dafny.Set({len((d_582_v55_).set(default__.safeIndex((0) - (d_523_v0_), len(d_582_v55_)), p1)), (0) - ((0) - (len(((d_585_v57_).set((d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))], d_583_i3_)).set((d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))], 304))))})
            d_586_v58_: C1
            nw95_ = C1()
            nw95_.ctor__((0) - (d_523_v0_), len(d_580_v53_), (self).f28)
            d_586_v58_ = nw95_
        d_587_i4_: int
        d_587_i4_ = 0
        with _dafny.label("7"):
            while not(p0):
                with _dafny.c_label("7"):
                    if (d_587_i4_) >= (100):
                        raise _dafny.Break("7")
                    d_587_i4_ = (d_587_i4_) + (1)
                    def iife84_():
                        coll42_ = _dafny.Set()
                        compr_42_: int
                        for compr_42_ in _dafny.IntegerRange(984, 105):
                            d_588_v59_: int = compr_42_
                            if ((984) <= (d_588_v59_)) and ((d_588_v59_) < (105)):
                                coll42_ = coll42_.union(_dafny.Set([(d_588_v59_) - (len((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))]))]))
                        return _dafny.Set(coll42_)
                    if (d_580_v53_)[default__.safeIndex(len(iife84_()
                    ), len(d_580_v53_))]:
                        d_589_v60_: _dafny.Map
                        d_589_v60_ = _dafny.Map({not(True): p1})
                        d_590_v61_: _dafny.Map
                        d_590_v61_ = _dafny.Map({(d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))]: len(d_582_v55_)})
                        d_591_v62_: D0
                        d_591_v62_ = D0_DC2(d_523_v0_, len(d_580_v53_))
                        d_592_v63_: _dafny.Map
                        d_592_v63_ = _dafny.Map({d_591_v62_: p0})
                        d_593_v64_: _dafny.Array
                        nw96_ = _dafny.Array(None, 27)
                        nw96_[int(0)] = 48
                        nw96_[int(1)] = -273
                        nw96_[int(2)] = d_523_v0_
                        nw96_[int(3)] = d_523_v0_
                        nw96_[int(4)] = len(d_589_v60_)
                        nw96_[int(5)] = -604
                        nw96_[int(6)] = (0) - (d_523_v0_)
                        nw96_[int(7)] = len((d_580_v53_) + (d_580_v53_))
                        nw96_[int(8)] = d_523_v0_
                        nw96_[int(9)] = d_523_v0_
                        nw96_[int(10)] = (_dafny.MultiSet([d_523_v0_])).cardinality
                        nw96_[int(11)] = d_523_v0_
                        nw96_[int(12)] = default__.safeModuloInt(d_523_v0_, (0) - (((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))])[default__.safeIndex(len((d_590_v61_).set(p0, d_523_v0_)), len((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))]))]))
                        nw96_[int(13)] = (d_523_v0_) + (len(d_580_v53_))
                        nw96_[int(14)] = len(((d_576_v49_ if (d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))] else (d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))])).set(default__.safeIndex(429, len((d_576_v49_ if (d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))] else (d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))]))), d_523_v0_))
                        nw96_[int(15)] = 30
                        nw96_[int(16)] = d_523_v0_
                        nw96_[int(17)] = (0) - ((274) * (d_523_v0_))
                        nw96_[int(18)] = d_523_v0_
                        nw96_[int(19)] = len(d_592_v63_)
                        nw96_[int(20)] = d_523_v0_
                        nw96_[int(21)] = d_523_v0_
                        nw96_[int(22)] = default__.safeDivisionInt(len((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))]), d_523_v0_)
                        nw96_[int(23)] = 583
                        nw96_[int(24)] = (d_523_v0_ if (d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))] else len((d_582_v55_).set(default__.safeIndex(d_523_v0_, len(d_582_v55_)), p1)))
                        nw96_[int(25)] = d_523_v0_
                        nw96_[int(26)] = (d_523_v0_) - (-629)
                        d_593_v64_ = nw96_
                        d_594_v65_: _dafny.Seq
                        d_594_v65_ = _dafny.SeqWithoutIsStrInference([(self).f40])
                        index109_ = default__.safeIndex(35, (d_593_v64_).length(0))
                        (d_593_v64_)[index109_] = (self).fm17((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))], len(d_594_v65_), globalState)
                        index110_ = default__.safeIndex(35, (d_593_v64_).length(0))
                        (d_593_v64_)[index110_] = (d_523_v0_) * (default__.safeModuloInt(d_523_v0_, d_523_v0_))
                        d_595_v66_: D2
                        d_595_v66_ = D2_DC7((d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))], (d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))], _dafny.SeqWithoutIsStrInference([p1, p1]))
                        d_596_v67_: _dafny.Array
                        nw97_ = _dafny.Array(_dafny.Map({}), 11)
                        d_596_v67_ = nw97_
                        d_597_v68_: _dafny.Array
                        d_597_v68_ = d_596_v67_
                        d_598_v69_: _dafny.MultiSet
                        d_598_v69_ = _dafny.MultiSet([d_597_v68_])
                        d_599_v70_: _dafny.Map
                        d_599_v70_ = _dafny.Map({(d_595_v66_).cf15: (((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))])[default__.safeIndex((d_593_v64_)[default__.safeIndex(35, (d_593_v64_).length(0))], len((d_575_v48_)[default__.safeIndex(89, (d_575_v48_).length(0))]))] if (d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))] else (d_598_v69_).cardinality)})
                        (globalState).f24 = ((d_599_v70_)[d_582_v55_] if (d_582_v55_) in (d_599_v70_) else (0) - (d_523_v0_))
                        index111_ = default__.safeIndex(35, (d_593_v64_).length(0))
                        (d_593_v64_)[index111_] = ((0) - (d_523_v0_)) * (d_523_v0_)
                        (globalState).f1 = (d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))]
                        (globalState).f14 = (d_593_v64_)[default__.safeIndex(35, (d_593_v64_).length(0))]
                    elif True:
                        index112_ = default__.safeIndex(429, (d_579_v52_).length(0))
                        rhs103_ = d_575_v48_
                        rhs104_ = p0
                        lhs92_ = d_579_v52_
                        lhs93_ = default__.safeIndex(429, (d_579_v52_).length(0))
                        d_575_v48_ = rhs103_
                        lhs92_[lhs93_] = rhs104_
                        d_600_v71_: C0
                        nw98_ = C0()
                        nw98_.ctor__((self).f40)
                        d_600_v71_ = nw98_
                        d_601_v72_: _dafny.MultiSet
                        d_601_v72_ = _dafny.MultiSet([d_523_v0_])
                        d_602_v73_: _dafny.Map
                        d_602_v73_ = _dafny.Map({(d_601_v72_).issubset(d_601_v72_): (d_523_v0_) * (167)})
                        (globalState).f0 = ((d_602_v73_)[p0] if (p0) in (d_602_v73_) else (0) - (-238))
                        (globalState).f11 = default__.fm4(globalState)
                        index113_ = default__.safeIndex(429, (d_579_v52_).length(0))
                        (d_579_v52_)[index113_] = p0
                    d_603_v74_: _dafny.Array
                    nw99_ = _dafny.Array(_dafny.Array(None, 0), 6)
                    d_603_v74_ = nw99_
                    d_604_v75_: _dafny.Array
                    nw100_ = _dafny.Array(None, 17)
                    nw100_[int(0)] = d_603_v74_
                    nw100_[int(1)] = d_603_v74_
                    nw100_[int(2)] = d_603_v74_
                    nw100_[int(3)] = d_603_v74_
                    nw100_[int(4)] = d_603_v74_
                    nw100_[int(5)] = d_603_v74_
                    nw100_[int(6)] = d_603_v74_
                    nw100_[int(7)] = d_603_v74_
                    nw100_[int(8)] = d_603_v74_
                    nw100_[int(9)] = d_603_v74_
                    nw100_[int(10)] = (d_603_v74_ if p0 else d_603_v74_)
                    nw100_[int(11)] = d_603_v74_
                    nw100_[int(12)] = d_603_v74_
                    nw100_[int(13)] = d_603_v74_
                    nw100_[int(14)] = d_603_v74_
                    nw100_[int(15)] = d_603_v74_
                    nw100_[int(16)] = d_603_v74_
                    d_604_v75_ = nw100_
                    index114_ = default__.safeIndex(107, (d_604_v75_).length(0))
                    (d_604_v75_)[index114_] = d_603_v74_
                    index115_ = default__.safeIndex(107, (d_604_v75_).length(0))
                    (d_604_v75_)[index115_] = d_603_v74_
                    (globalState).f24 = (d_523_v0_) + (d_523_v0_)
                    d_605_v76_: _dafny.Map
                    d_605_v76_ = _dafny.Map({d_523_v0_: ((d_580_v53_).set(default__.safeIndex(d_523_v0_, len(d_580_v53_)), p0)) == (d_580_v53_)})
                    d_606_v77_: _dafny.Array
                    def lambda25_(d_607_v0_):
                        def lambda26_(d_608_i5_):
                            return default__.safeDivisionInt(d_608_i5_, d_607_v0_)

                        return lambda26_

                    init13_ = lambda25_(d_523_v0_)
                    nw101_ = _dafny.Array(None, 20)
                    for i0_13_ in range(nw101_.length(0)):
                        nw101_[i0_13_] = init13_(i0_13_)
                    d_606_v77_ = nw101_
                    d_609_v78_: _dafny.Set
                    d_609_v78_ = _dafny.Set({d_606_v77_})
                    d_605_v76_ = (d_605_v76_).set(len(d_609_v78_), p0)
                    pass
            pass
        d_610_v79_: _dafny.Map
        d_610_v79_ = _dafny.Map({d_523_v0_: 165})
        d_611_v80_: _dafny.Map
        d_611_v80_ = _dafny.Map({_dafny.MultiSet(d_580_v53_): len(d_610_v79_)})
        d_611_v80_ = (d_611_v80_).set((d_578_v51_).set((d_579_v52_)[default__.safeIndex(429, (d_579_v52_).length(0))], default__.abs(d_523_v0_)), d_523_v0_)
        r0 = (d_523_v0_) - (d_523_v0_)
        return r0

    def m15(self, p0, p1, p2, p3, globalState):
        r0: int = int(0)
        r1: bool = False
        d_612_i0_: int
        d_612_i0_ = 0
        with _dafny.label("8"):
            while default__.fm4(globalState):
                with _dafny.c_label("8"):
                    if (d_612_i0_) >= (100):
                        raise _dafny.Break("8")
                    d_612_i0_ = (d_612_i0_) + (1)
                    r1 = p3
                    d_613_v0_: _dafny.Array
                    nw102_ = _dafny.Array(_dafny.CodePoint('D'), 6)
                    d_613_v0_ = nw102_
                    d_614_v1_: str
                    d_614_v1_ = _dafny.CodePoint('t')
                    index116_ = default__.safeIndex(518, (d_613_v0_).length(0))
                    (d_613_v0_)[index116_] = d_614_v1_
                    index117_ = default__.safeIndex(518, (d_613_v0_).length(0))
                    (d_613_v0_)[index117_] = d_614_v1_
                    d_615_v2_: _dafny.Array
                    nw103_ = _dafny.Array(int(0), 5)
                    d_615_v2_ = nw103_
                    index118_ = default__.safeIndex(899, (d_615_v2_).length(0))
                    (d_615_v2_)[index118_] = p1
                    d_616_v3_: _dafny.Map
                    d_616_v3_ = _dafny.Map({p1: p1})
                    d_617_v4_: _dafny.Set
                    d_617_v4_ = _dafny.Set({p1})
                    d_618_v5_: _dafny.Seq
                    d_618_v5_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({133, len(d_616_v3_), p1}), d_617_v4_])
                    d_619_v6_: _dafny.Seq
                    d_619_v6_ = d_618_v5_
                    index119_ = default__.safeIndex(899, (d_615_v2_).length(0))
                    (d_615_v2_)[index119_] = len((d_618_v5_) + ((d_619_v6_)))
                    d_620_v7_: _dafny.Array
                    def lambda27_(d_621_p0_):
                        def lambda28_(d_622_i1_):
                            return D1_DC3(d_621_p0_)

                        return lambda28_

                    init14_ = lambda27_(p0)
                    nw104_ = _dafny.Array(None, 16)
                    for i0_14_ in range(nw104_.length(0)):
                        nw104_[i0_14_] = init14_(i0_14_)
                    d_620_v7_ = nw104_
                    d_623_v8_: _dafny.MultiSet
                    d_623_v8_ = _dafny.MultiSet([d_620_v7_, d_620_v7_, d_620_v7_, d_620_v7_])
                    d_623_v8_ = d_623_v8_
                    pass
            pass
        d_624_v9_: str
        d_624_v9_ = _dafny.CodePoint('p')
        d_625_v10_: _dafny.Map
        d_625_v10_ = _dafny.Map({d_624_v9_: p1})
        d_626_v11_: _dafny.Seq
        d_626_v11_ = _dafny.SeqWithoutIsStrInference([d_625_v10_])
        if (self).fm7(p0, (d_626_v11_) <= (d_626_v11_), globalState):
            d_627_v12_: _dafny.Map
            d_627_v12_ = _dafny.Map({p1: not(p3)})
            d_628_v13_: _dafny.Seq
            d_628_v13_ = _dafny.SeqWithoutIsStrInference([len(d_627_v12_)])
            d_628_v13_ = d_628_v13_
            (globalState).f11 = p0
            d_629_v14_: _dafny.Array
            nw105_ = _dafny.Array(False, 3)
            d_629_v14_ = nw105_
            d_630_v15_: _dafny.Array
            nw106_ = _dafny.Array(int(0), 1)
            d_630_v15_ = nw106_
            d_631_v16_: _dafny.Map
            d_631_v16_ = _dafny.Map({d_629_v14_: d_630_v15_})
            d_631_v16_ = (d_631_v16_).set(d_629_v14_, d_630_v15_)
            d_632_v17_: C0
            nw107_ = C0()
            nw107_.ctor__((self).f40)
            d_632_v17_ = nw107_
            d_633_v18_: _dafny.Seq
            d_633_v18_ = _dafny.SeqWithoutIsStrInference([p0])
            d_634_v19_: _dafny.MultiSet
            d_634_v19_ = _dafny.MultiSet([p3, p3, p0, p3, p3])
            (globalState).f19 = (default__.safeDivisionInt(-331, len(_dafny.Map({p3: d_624_v9_}))) if (672) < (len(d_633_v18_)) else ((d_634_v19_)[default__.fm4(globalState)] if (default__.fm4(globalState)) in (d_634_v19_) else len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrm")))))
        elif True:
            d_635_v20_: _dafny.Array
            def lambda29_(d_636_i2_):
                return default__.safeDivisionInt(d_636_i2_, 494)

            init15_ = lambda29_
            nw108_ = _dafny.Array(None, 22)
            for i0_15_ in range(nw108_.length(0)):
                nw108_[i0_15_] = init15_(i0_15_)
            d_635_v20_ = nw108_
            d_637_v21_: _dafny.Array
            nw109_ = _dafny.Array(None, 24)
            nw109_[int(0)] = d_635_v20_
            nw109_[int(1)] = d_635_v20_
            nw109_[int(2)] = d_635_v20_
            nw109_[int(3)] = d_635_v20_
            nw109_[int(4)] = d_635_v20_
            nw109_[int(5)] = d_635_v20_
            nw109_[int(6)] = d_635_v20_
            nw109_[int(7)] = d_635_v20_
            nw109_[int(8)] = d_635_v20_
            nw109_[int(9)] = d_635_v20_
            nw109_[int(10)] = d_635_v20_
            nw109_[int(11)] = d_635_v20_
            nw109_[int(12)] = d_635_v20_
            nw109_[int(13)] = d_635_v20_
            nw109_[int(14)] = d_635_v20_
            nw109_[int(15)] = d_635_v20_
            nw109_[int(16)] = d_635_v20_
            nw109_[int(17)] = d_635_v20_
            nw109_[int(18)] = d_635_v20_
            nw109_[int(19)] = d_635_v20_
            nw109_[int(20)] = d_635_v20_
            nw109_[int(21)] = d_635_v20_
            nw109_[int(22)] = d_635_v20_
            nw109_[int(23)] = d_635_v20_
            d_637_v21_ = nw109_
            rhs105_ = True
            rhs106_ = d_637_v21_
            rhs107_ = (p0) and (p3)
            lhs94_ = globalState
            lhs95_ = globalState
            lhs94_.f27 = rhs105_
            d_637_v21_ = rhs106_
            lhs95_.f1 = rhs107_
            d_638_v22_: D2
            d_638_v22_ = D2_DC8(p0, d_635_v20_)
            source10_ = d_638_v22_
            if source10_.is_DC7:
                d_639___mcc_h0_ = source10_.cf13
                d_640___mcc_h1_ = source10_.cf14
                d_641___mcc_h2_ = source10_.cf15
                d_642_cf15_ = d_641___mcc_h2_
                d_643_cf14_ = d_640___mcc_h1_
                d_644_cf13_ = d_639___mcc_h0_
                d_645_v23_: _dafny.Map
                d_645_v23_ = _dafny.Map({p3: _dafny.SeqWithoutIsStrInference([d_624_v9_ for d_646_i3_ in range(default__.abs(611))])})
                d_647_v24_: _dafny.Map
                d_647_v24_ = _dafny.Map({p1: d_642_cf15_})
                d_648_v25_: _dafny.Map
                d_648_v25_ = _dafny.Map({p1: p0})
                d_649_v26_: _dafny.Seq
                d_649_v26_ = _dafny.SeqWithoutIsStrInference([not(p0), ((d_648_v25_)[p1] if (p1) in (d_648_v25_) else p0)])
                d_650_v27_: _dafny.Map
                d_650_v27_ = _dafny.Map({p1: d_624_v9_})
                d_651_v28_: _dafny.Set
                d_651_v28_ = _dafny.Set({(0) - (-756)})
                rhs108_ = (0) - (len(((d_642_cf15_) + (((d_645_v23_)[p3] if (p3) in (d_645_v23_) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_652_i4_ in range(default__.abs(836))])))) + ((_dafny.SeqWithoutIsStrInference([d_624_v9_ for d_653_i5_ in range(default__.abs(73))])) + (((d_647_v24_)[p1] if (p1) in (d_647_v24_) else d_642_cf15_)))))
                rhs109_ = (default__.safeModuloInt(len(d_649_v26_), len(d_650_v27_))) > (p1)
                rhs110_ = (d_651_v28_).issubset(d_651_v28_)
                rhs111_ = p0
                lhs96_ = globalState
                lhs97_ = globalState
                lhs96_.f19 = rhs108_
                lhs97_.f1 = rhs109_
                d_643_cf14_ = rhs110_
                d_643_cf14_ = rhs111_
                d_654_v29_: _dafny.Array
                nw110_ = _dafny.Array(False, 16)
                d_654_v29_ = nw110_
                d_654_v29_ = d_654_v29_
                d_655_v30_: _dafny.Map
                d_655_v30_ = _dafny.Map({p3: d_635_v20_})
                index120_ = default__.safeIndex(516, (d_654_v29_).length(0))
                (d_654_v29_)[index120_] = (True) or (p3)
                d_656_v31_: _dafny.Array
                nw111_ = _dafny.Array(_dafny.Array(None, 0), 6)
                d_656_v31_ = nw111_
                d_657_v32_: _dafny.Seq
                d_657_v32_ = _dafny.SeqWithoutIsStrInference([p1])
                d_658_v33_: _dafny.Array
                nw112_ = _dafny.Array(None, 3)
                nw112_[int(0)] = d_657_v32_
                nw112_[int(1)] = _dafny.SeqWithoutIsStrInference([p1])
                nw112_[int(2)] = d_657_v32_
                d_658_v33_ = nw112_
                index121_ = default__.safeIndex(56, (d_656_v31_).length(0))
                (d_656_v31_)[index121_] = d_658_v33_
                d_659_v34_: _dafny.MultiSet
                d_659_v34_ = _dafny.MultiSet([p1, p1, p1, 720])
                index122_ = default__.safeIndex(516, (d_654_v29_).length(0))
                index123_ = default__.safeIndex(56, (d_656_v31_).length(0))
                rhs112_ = ((d_655_v30_) | ((d_655_v30_).set(d_644_cf13_, d_635_v20_))) | (d_655_v30_)
                rhs113_ = not((d_659_v34_).ispropersubset(d_659_v34_))
                rhs114_ = (d_658_v33_ if p0 else d_658_v33_)
                lhs98_ = d_654_v29_
                lhs99_ = default__.safeIndex(516, (d_654_v29_).length(0))
                lhs100_ = d_656_v31_
                lhs101_ = default__.safeIndex(56, (d_656_v31_).length(0))
                d_655_v30_ = rhs112_
                lhs98_[lhs99_] = rhs113_
                lhs100_[lhs101_] = rhs114_
                d_660_v35_: D2
                d_660_v35_ = D2_DC7(False, d_643_cf14_, default__.fm3(globalState))
                d_661_v36_: _dafny.Array
                nw113_ = _dafny.Array(None, 24)
                nw113_[int(0)] = d_642_cf15_
                nw113_[int(1)] = d_642_cf15_
                nw113_[int(2)] = (d_642_cf15_) + (default__.fm3(globalState))
                nw113_[int(3)] = d_642_cf15_
                nw113_[int(4)] = (d_660_v35_).cf15
                nw113_[int(5)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mopoyayn"))
                nw113_[int(6)] = (d_642_cf15_) + (d_642_cf15_)
                nw113_[int(7)] = d_642_cf15_
                nw113_[int(8)] = d_642_cf15_
                nw113_[int(9)] = d_642_cf15_
                nw113_[int(10)] = d_642_cf15_
                nw113_[int(11)] = d_642_cf15_
                nw113_[int(12)] = d_642_cf15_
                nw113_[int(13)] = d_642_cf15_
                nw113_[int(14)] = d_642_cf15_
                nw113_[int(15)] = d_642_cf15_
                nw113_[int(16)] = d_642_cf15_
                nw113_[int(17)] = d_642_cf15_
                nw113_[int(18)] = (default__.fm3(globalState)) + (d_642_cf15_)
                nw113_[int(19)] = d_642_cf15_
                nw113_[int(20)] = d_642_cf15_
                nw113_[int(21)] = (d_642_cf15_) + (_dafny.SeqWithoutIsStrInference([d_624_v9_ for d_662_i6_ in range(default__.abs(-628))]))
                nw113_[int(22)] = (d_642_cf15_) + (d_642_cf15_)
                nw113_[int(23)] = default__.fm3(globalState)
                d_661_v36_ = nw113_
                index124_ = default__.safeIndex(498, (d_661_v36_).length(0))
                (d_661_v36_)[index124_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "itmx"))
                d_663_v37_: _dafny.Map
                d_663_v37_ = _dafny.Map({p3: False})
                index125_ = default__.safeIndex(498, (d_661_v36_).length(0))
                index126_ = default__.safeIndex(516, (d_654_v29_).length(0))
                rhs115_ = d_642_cf15_
                rhs116_ = (d_651_v28_).ispropersubset((d_651_v28_).intersection(d_651_v28_))
                rhs117_ = d_663_v37_
                rhs118_ = not(p0)
                rhs119_ = ((p1) + (p1)) < (default__.safeModuloInt(p1, p1))
                lhs102_ = d_661_v36_
                lhs103_ = default__.safeIndex(498, (d_661_v36_).length(0))
                lhs104_ = d_654_v29_
                lhs105_ = default__.safeIndex(516, (d_654_v29_).length(0))
                lhs106_ = globalState
                lhs102_[lhs103_] = rhs115_
                lhs104_[lhs105_] = rhs116_
                d_663_v37_ = rhs117_
                d_643_cf14_ = rhs118_
                lhs106_.f1 = rhs119_
            elif source10_.is_DC8:
                d_664___mcc_h3_ = source10_.cf16
                d_665___mcc_h4_ = source10_.cf17
                d_666_cf17_ = d_665___mcc_h4_
                d_667_cf16_ = d_664___mcc_h3_
                d_668_v38_: _dafny.Seq
                d_668_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fft"))
                d_668_v38_ = (d_668_v38_) + ((d_668_v38_) + (d_668_v38_))
                (globalState).f1 = p0
                d_669_v39_: _dafny.Map
                d_669_v39_ = _dafny.Map({p0: (p1) - (len(d_668_v38_))})
                d_670_v40_: _dafny.MultiSet
                d_670_v40_ = _dafny.MultiSet([p0])
                d_669_v39_ = ((d_669_v39_) | (d_669_v39_)).set(p0, ((d_670_v40_).cardinality) + (len(d_668_v38_)))
                (globalState).f0 = p1
            elif True:
                d_671___mcc_h5_ = source10_.cf12
                d_672_cf12_ = d_671___mcc_h5_
                d_673_v41_: _dafny.Seq
                d_673_v41_ = _dafny.SeqWithoutIsStrInference([d_624_v9_, d_624_v9_])
                d_674_v42_: _dafny.Set
                d_674_v42_ = _dafny.Set({d_673_v41_})
                index127_ = default__.safeIndex(517, (d_672_cf12_).length(0))
                (d_672_cf12_)[index127_] = (_dafny.Set({default__.fm3(globalState)})).ispropersubset(d_674_v42_)
                d_675_v43_: _dafny.Map
                d_675_v43_ = _dafny.Map({p0: p1})
                index128_ = default__.safeIndex(517, (d_672_cf12_).length(0))
                (d_672_cf12_)[index128_] = not((not((not(False) if not(True) else p3))) in (d_675_v43_))
                d_676_v44_: _dafny.MultiSet
                d_676_v44_ = _dafny.MultiSet([p3, p0, p0, p3, p0])
                d_677_v45_: _dafny.Map
                d_677_v45_ = _dafny.Map({(d_673_v41_)[default__.safeIndex(p1, len(d_673_v41_))]: (d_676_v44_).ispropersubset(d_676_v44_)})
                d_678_v46_: _dafny.Map
                d_678_v46_ = _dafny.Map({not(p0): d_624_v9_})
                d_677_v45_ = (d_677_v45_).set(default__.fm5(default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ohnr")), (self).f40, d_678_v46_, p3, globalState), (d_672_cf12_)[default__.safeIndex(517, (d_672_cf12_).length(0))], globalState), (d_672_cf12_)[default__.safeIndex(517, (d_672_cf12_).length(0))])
                d_679_v47_: D4
                d_679_v47_ = D4_DC10(d_624_v9_)
                pat_let_tv20_ = d_624_v9_
                def iife85_(_pat_let21_0):
                    def iife86_(d_680_dt__update__tmp_h0_):
                        def iife87_(_pat_let22_0):
                            def iife88_(d_681_dt__update_hcf19_h0_):
                                return D4_DC10(d_681_dt__update_hcf19_h0_)
                            return iife88_(_pat_let22_0)
                        return iife87_(pat_let_tv20_)
                    return iife86_(_pat_let21_0)
                d_679_v47_ = iife85_(d_679_v47_)
                d_682_v48_: D2
                d_682_v48_ = D2_DC7(p3, p3, (d_673_v41_) + (d_673_v41_))
                d_683_v49_: _dafny.Seq
                d_683_v49_ = _dafny.SeqWithoutIsStrInference([p1, p1, p1])
                d_684_v50_: D0
                d_684_v50_ = D0_DC1(p1)
                d_685_v51_: _dafny.Map
                d_685_v51_ = _dafny.Map({d_684_v50_: 144})
                rhs120_ = (self).fm17(d_683_v49_, -678, globalState)
                rhs121_ = D2_DC7((d_672_cf12_)[default__.safeIndex(517, (d_672_cf12_).length(0))], False, _dafny.SeqWithoutIsStrInference([d_624_v9_]))
                rhs122_ = len(d_685_v51_)
                rhs123_ = _dafny.CodePoint('t')
                lhs107_ = globalState
                lhs108_ = globalState
                lhs107_.f14 = rhs120_
                d_682_v48_ = rhs121_
                lhs108_.f24 = rhs122_
                d_624_v9_ = rhs123_
            d_686_v52_: _dafny.Array
            nw114_ = _dafny.Array(_dafny.Array(None, 0), 14)
            d_686_v52_ = nw114_
            d_687_v53_: _dafny.Array
            nw115_ = _dafny.Array(_dafny.CodePoint('D'), 29)
            d_687_v53_ = nw115_
            index129_ = default__.safeIndex(700, (d_686_v52_).length(0))
            (d_686_v52_)[index129_] = d_687_v53_
            index130_ = default__.safeIndex(700, (d_686_v52_).length(0))
            nw116_ = _dafny.Array(_dafny.CodePoint('D'), 27)
            (d_686_v52_)[index130_] = nw116_
            (globalState).f1 = ((self).f28).cf4
            d_688_v54_: _dafny.Seq
            d_688_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "colx"))
            d_688_v54_ = d_688_v54_
        if p0:
            d_689_v55_: _dafny.Seq
            d_689_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kjcqanx"))
            d_690_v56_: _dafny.Map
            d_690_v56_ = _dafny.Map({p0: d_624_v9_})
            d_691_v57_: _dafny.Map
            d_691_v57_ = _dafny.Map({default__.fm2(d_689_v55_, _dafny.Set({False}), d_690_v56_, p3, globalState): len((self).f40)})
            d_692_v58_: D0
            d_692_v58_ = D0_DC0(d_691_v57_)
            rhs124_ = default__.fm3(globalState)
            rhs125_ = (0) - (p1)
            rhs126_ = (0) - (default__.safeModuloInt(default__.safeDivisionInt(586, 859), (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqfbcws")))) + (p1))))
            rhs127_ = d_692_v58_
            lhs109_ = globalState
            lhs110_ = globalState
            d_689_v55_ = rhs124_
            lhs109_.f19 = rhs125_
            lhs110_.f24 = rhs126_
            d_692_v58_ = rhs127_
            d_693_v59_: _dafny.Array
            nw117_ = _dafny.Array(False, 10)
            d_693_v59_ = nw117_
            index131_ = default__.safeIndex(252, (d_693_v59_).length(0))
            (d_693_v59_)[index131_] = (p3 if p3 else p3)
            d_694_v60_: _dafny.Seq
            d_694_v60_ = _dafny.SeqWithoutIsStrInference([d_689_v55_])
            index132_ = default__.safeIndex(252, (d_693_v59_).length(0))
            (d_693_v59_)[index132_] = ((_dafny.SeqWithoutIsStrInference([d_689_v55_ for d_695_i7_ in range(default__.abs(669))])) + (d_694_v60_)) == ((d_694_v60_) + (d_694_v60_))
            if p3:
                (globalState).f11 = False
                (globalState).f1 = (p1) == ((p1) + ((D0_DC1(p1)).cf1))
                d_696_v61_: D4
                d_696_v61_ = D4_DC10(d_624_v9_)
                d_696_v61_ = D4_DC10(d_624_v9_)
                (globalState).f27 = not(False)
                (globalState).f19 = (p1) - ((0) - (p1))
            elif True:
                d_697_v62_: _dafny.Array
                nw118_ = _dafny.Array(_dafny.Array(None, 0), 16)
                d_697_v62_ = nw118_
                d_698_v63_: _dafny.Seq
                d_698_v63_ = _dafny.SeqWithoutIsStrInference([len(d_689_v55_)])
                d_699_v64_: _dafny.Seq
                d_699_v64_ = _dafny.SeqWithoutIsStrInference([(d_693_v59_)[default__.safeIndex(252, (d_693_v59_).length(0))]])
                d_700_v65_: _dafny.MultiSet
                d_700_v65_ = _dafny.MultiSet([p0, p0, p0])
                d_701_v66_: _dafny.Array
                nw119_ = _dafny.Array(None, 29)
                nw119_[int(0)] = p1
                nw119_[int(1)] = p1
                nw119_[int(2)] = p1
                nw119_[int(3)] = p1
                nw119_[int(4)] = len(d_691_v57_)
                nw119_[int(5)] = p1
                nw119_[int(6)] = 507
                nw119_[int(7)] = p1
                nw119_[int(8)] = 965
                nw119_[int(9)] = p1
                nw119_[int(10)] = p1
                nw119_[int(11)] = (d_698_v63_)[default__.safeIndex(p1, len(d_698_v63_))]
                nw119_[int(12)] = p1
                nw119_[int(13)] = p1
                nw119_[int(14)] = (d_698_v63_)[default__.safeIndex((0) - (p1), len(d_698_v63_))]
                nw119_[int(15)] = p1
                nw119_[int(16)] = (0) - (p1)
                nw119_[int(17)] = p1
                nw119_[int(18)] = p1
                nw119_[int(19)] = p1
                nw119_[int(20)] = p1
                nw119_[int(21)] = p1
                nw119_[int(22)] = (0) - (p1)
                nw119_[int(23)] = len(d_699_v64_)
                nw119_[int(24)] = (0) - (p1)
                nw119_[int(25)] = p1
                nw119_[int(26)] = p1
                nw119_[int(27)] = ((d_700_v65_)[p3] if (p3) in (d_700_v65_) else p1)
                nw119_[int(28)] = p1
                d_701_v66_ = nw119_
                index133_ = default__.safeIndex(384, (d_697_v62_).length(0))
                (d_697_v62_)[index133_] = d_701_v66_
                index134_ = default__.safeIndex(384, (d_697_v62_).length(0))
                (d_697_v62_)[index134_] = d_701_v66_
                index135_ = default__.safeIndex(666, (d_701_v66_).length(0))
                (d_701_v66_)[index135_] = p1
                index136_ = default__.safeIndex(666, (d_701_v66_).length(0))
                (d_701_v66_)[index136_] = default__.safeDivisionInt(p1, (p1) - (p1))
                d_702_v68_: D2
                d_702_v68_ = D2_DC7((d_693_v59_)[default__.safeIndex(252, (d_693_v59_).length(0))], p3, _dafny.SeqWithoutIsStrInference([d_624_v9_]))
                d_703_v69_: _dafny.Map
                def iife89_():
                    coll43_ = _dafny.Set()
                    compr_43_: int
                    for compr_43_ in _dafny.IntegerRange(312, 531):
                        d_704_v67_: int = compr_43_
                        if ((312) <= (d_704_v67_)) and ((d_704_v67_) < (531)):
                            coll43_ = coll43_.union(_dafny.Set([default__.safeDivisionInt(d_704_v67_, len(_dafny.Set({(d_701_v66_)[default__.safeIndex(666, (d_701_v66_).length(0))]})))]))
                    return _dafny.Set(coll43_)
                d_703_v69_ = _dafny.Map({len(iife89_()
                ): d_702_v68_})
                d_705_v71_: _dafny.Map
                def iife90_():
                    coll44_ = _dafny.Map()
                    compr_44_: int
                    for compr_44_ in (d_691_v57_).keys.Elements:
                        d_706_v70_: int = compr_44_
                        if (d_706_v70_) in (d_691_v57_):
                            coll44_[(d_706_v70_) * (p1)] = d_702_v68_
                    return _dafny.Map(coll44_)
                d_705_v71_ = _dafny.Map({(d_701_v66_)[default__.safeIndex(666, (d_701_v66_).length(0))]: iife90_()
                })
                d_707_v72_: _dafny.Set
                d_707_v72_ = _dafny.Set({p1})
                d_708_v73_: _dafny.Map
                d_708_v73_ = _dafny.Map({default__.fm4(globalState): len(d_707_v72_)})
                d_709_v74_: _dafny.Map
                d_709_v74_ = _dafny.Map({len(d_708_v73_): True})
                d_710_v75_: _dafny.Array
                nw120_ = _dafny.Array(None, 13)
                nw120_[int(0)] = d_703_v69_
                nw120_[int(1)] = _dafny.Map({p1: d_702_v68_})
                nw120_[int(2)] = ((d_705_v71_)[224] if (224) in (d_705_v71_) else _dafny.Map({len(d_709_v74_): d_702_v68_}))
                nw120_[int(3)] = d_703_v69_
                nw120_[int(4)] = d_703_v69_
                nw120_[int(5)] = (d_703_v69_) | (_dafny.Map({(_dafny.MultiSet([p1])).cardinality: D2_DC7((d_693_v59_)[default__.safeIndex(252, (d_693_v59_).length(0))], (d_693_v59_)[default__.safeIndex(252, (d_693_v59_).length(0))], d_689_v55_)}))
                nw120_[int(6)] = d_703_v69_
                nw120_[int(7)] = d_703_v69_
                nw120_[int(8)] = d_703_v69_
                nw120_[int(9)] = (d_703_v69_).set((d_701_v66_)[default__.safeIndex(666, (d_701_v66_).length(0))], d_702_v68_)
                nw120_[int(10)] = (d_703_v69_) | (d_703_v69_)
                nw120_[int(11)] = d_703_v69_
                nw120_[int(12)] = d_703_v69_
                d_710_v75_ = nw120_
                index137_ = default__.safeIndex(324, (d_710_v75_).length(0))
                (d_710_v75_)[index137_] = d_703_v69_
                d_711_v76_: T0
                nw121_ = C1()
                nw121_.ctor__(p1, (d_701_v66_)[default__.safeIndex(666, (d_701_v66_).length(0))], D1_DC3((d_693_v59_)[default__.safeIndex(252, (d_693_v59_).length(0))]))
                d_711_v76_ = nw121_
                d_712_v77_: _dafny.Seq
                d_712_v77_ = _dafny.SeqWithoutIsStrInference([d_711_v76_, d_711_v76_])
                index138_ = default__.safeIndex(324, (d_710_v75_).length(0))
                (d_710_v75_)[index138_] = (d_703_v69_).set(len(((D6_DC15(d_712_v77_)).cf25).set(default__.safeIndex((d_701_v66_)[default__.safeIndex(666, (d_701_v66_).length(0))], len((D6_DC15(d_712_v77_)).cf25)), d_711_v76_)), d_702_v68_)
                pat_let_tv21_ = d_693_v59_
                pat_let_tv22_ = d_693_v59_
                d_713_v78_: C1
                nw122_ = C1()
                def iife91_(_pat_let23_0):
                    def iife92_(d_714_dt__update__tmp_h1_):
                        def iife93_(_pat_let24_0):
                            def iife94_(d_715_dt__update_hcf4_h0_):
                                return D1_DC3(d_715_dt__update_hcf4_h0_)
                            return iife94_(_pat_let24_0)
                        return iife93_((pat_let_tv22_)[default__.safeIndex(252, (pat_let_tv21_).length(0))])
                    return iife92_(_pat_let23_0)
                nw122_.ctor__(p1, default__.safeDivisionInt(p1, p1), (iife91_((self).f28) if not(True) else (d_711_v76_).f28))
                d_713_v78_ = nw122_
                (globalState).f27 = True
            (globalState).f14 = p1
            d_716_v79_: D0
            d_716_v79_ = D0_DC2(len(d_689_v55_), p1)
            (globalState).f24 = (d_716_v79_).cf3
        elif True:
            d_717_v80_: _dafny.Seq
            d_717_v80_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fvnkssion"))
            (globalState).f19 = default__.safeModuloInt(p1, (default__.fm26(p3, d_717_v80_, not(p0), globalState)).cardinality)
            d_718_v81_: _dafny.Array
            nw123_ = _dafny.Array(False, 24)
            d_718_v81_ = nw123_
            d_719_v82_: _dafny.Map
            d_719_v82_ = _dafny.Map({p1: True})
            index139_ = default__.safeIndex(692, (d_718_v81_).length(0))
            (d_718_v81_)[index139_] = (len(d_719_v82_)) <= (p1)
            d_720_v83_: _dafny.Seq
            d_720_v83_ = _dafny.SeqWithoutIsStrInference([p1])
            index140_ = default__.safeIndex(692, (d_718_v81_).length(0))
            (d_718_v81_)[index140_] = not ((d_624_v9_) in (d_717_v80_)) or (((d_720_v83_)[default__.safeIndex(default__.fm2(d_717_v80_, _dafny.Set({p0}), _dafny.Map({p0: d_624_v9_}), True, globalState), len(d_720_v83_))]) < (p1))
            d_721_v84_: _dafny.Set
            d_721_v84_ = _dafny.Set({d_624_v9_})
            d_722_v85_: _dafny.Map
            d_722_v85_ = _dafny.Map({d_624_v9_: d_624_v9_})
            d_723_v86_: _dafny.Seq
            d_723_v86_ = _dafny.SeqWithoutIsStrInference([d_718_v81_])
            d_724_v87_: _dafny.Array
            nw124_ = _dafny.Array(None, 18)
            nw124_[int(0)] = p1
            nw124_[int(1)] = default__.safeDivisionInt(p1, p1)
            nw124_[int(2)] = p1
            nw124_[int(3)] = ((0) - (p1)) * ((0) - ((self).fm17(_dafny.SeqWithoutIsStrInference([p1]), p1, globalState)))
            nw124_[int(4)] = len((d_720_v83_) + ((d_720_v83_).set(default__.safeIndex(p1, len(d_720_v83_)), p1)))
            nw124_[int(5)] = len(d_721_v84_)
            nw124_[int(6)] = p1
            nw124_[int(7)] = len(d_722_v85_)
            nw124_[int(8)] = p1
            nw124_[int(9)] = 848
            nw124_[int(10)] = default__.safeModuloInt((self).fm17(d_720_v83_, p1, globalState), len(_dafny.Map({p1: p1})))
            nw124_[int(11)] = len(d_723_v86_)
            nw124_[int(12)] = p1
            nw124_[int(13)] = 101
            nw124_[int(14)] = (p1) - (p1)
            nw124_[int(15)] = p1
            nw124_[int(16)] = default__.safeModuloInt(len(d_717_v80_), p1)
            nw124_[int(17)] = len(d_720_v83_)
            d_724_v87_ = nw124_
            index141_ = default__.safeIndex(579, (d_724_v87_).length(0))
            (d_724_v87_)[index141_] = p1
            index142_ = default__.safeIndex(579, (d_724_v87_).length(0))
            rhs128_ = d_624_v9_
            rhs129_ = (len(_dafny.SeqWithoutIsStrInference([d_624_v9_]))) * (p1)
            rhs130_ = p1
            lhs111_ = d_724_v87_
            lhs112_ = default__.safeIndex(579, (d_724_v87_).length(0))
            lhs113_ = globalState
            d_624_v9_ = rhs128_
            lhs111_[lhs112_] = rhs129_
            lhs113_.f14 = rhs130_
            nw125_ = _dafny.Array(None, 1)
            nw125_[int(0)] = (0) - (p1)
            d_724_v87_ = nw125_
            d_725_v88_: D4
            d_725_v88_ = D4_DC10(d_624_v9_)
            d_726_v89_: _dafny.Map
            d_726_v89_ = _dafny.Map({d_725_v88_: (d_724_v87_)[default__.safeIndex(579, (d_724_v87_).length(0))]})
            d_727_v91_: _dafny.Map
            d_727_v91_ = _dafny.Map({D4_DC10(d_624_v9_): p3})
            d_728_v92_: D7
            d_728_v92_ = D7_DC17(d_726_v89_)
            d_729_v93_: _dafny.Seq
            d_729_v93_ = _dafny.SeqWithoutIsStrInference([d_726_v89_, d_726_v89_])
            d_730_v94_: _dafny.Map
            d_730_v94_ = _dafny.Map({(d_718_v81_)[default__.safeIndex(692, (d_718_v81_).length(0))]: True})
            d_731_v95_: D8
            d_731_v95_ = D8_DC19(d_730_v94_)
            d_732_v96_: _dafny.Set
            d_732_v96_ = _dafny.Set({len((d_731_v95_).cf30), (d_724_v87_)[default__.safeIndex(579, (d_724_v87_).length(0))]})
            d_733_v98_: _dafny.Array
            nw126_ = _dafny.Array(None, 27)
            nw126_[int(0)] = d_726_v89_
            def iife95_():
                coll45_ = _dafny.Map()
                compr_45_: D4
                for compr_45_ in (d_727_v91_).keys.Elements:
                    d_734_v90_: D4 = compr_45_
                    if (d_734_v90_) in (d_727_v91_):
                        coll45_[d_734_v90_] = (d_724_v87_)[default__.safeIndex(579, (d_724_v87_).length(0))]
                return _dafny.Map(coll45_)
            nw126_[int(1)] = iife95_()
            
            nw126_[int(2)] = (d_728_v92_).cf28
            nw126_[int(3)] = default__.fm27(len(d_717_v80_), globalState)
            nw126_[int(4)] = d_726_v89_
            nw126_[int(5)] = d_726_v89_
            nw126_[int(6)] = _dafny.Map({d_725_v88_: p1})
            nw126_[int(7)] = d_726_v89_
            nw126_[int(8)] = d_726_v89_
            nw126_[int(9)] = d_726_v89_
            nw126_[int(10)] = d_726_v89_
            nw126_[int(11)] = d_726_v89_
            nw126_[int(12)] = (d_729_v93_)[default__.safeIndex(len(d_732_v96_), len(d_729_v93_))]
            nw126_[int(13)] = (d_726_v89_).set(d_725_v88_, (0) - ((d_724_v87_)[default__.safeIndex(579, (d_724_v87_).length(0))]))
            nw126_[int(14)] = d_726_v89_
            nw126_[int(15)] = d_726_v89_
            nw126_[int(16)] = d_726_v89_
            nw126_[int(17)] = (d_726_v89_) | (d_726_v89_)
            nw126_[int(18)] = d_726_v89_
            nw126_[int(19)] = d_726_v89_
            def iife96_():
                coll46_ = _dafny.Map()
                compr_46_: D4
                for compr_46_ in (_dafny.SeqWithoutIsStrInference([d_725_v88_ for d_735_i8_ in range(default__.abs(326))])).Elements:
                    d_736_v97_: D4 = compr_46_
                    if (d_736_v97_) in (_dafny.SeqWithoutIsStrInference([d_725_v88_ for d_735_i8_ in range(default__.abs(326))])):
                        coll46_[d_736_v97_] = (d_724_v87_)[default__.safeIndex(579, (d_724_v87_).length(0))]
                return _dafny.Map(coll46_)
            nw126_[int(20)] = (d_726_v89_ if default__.fm4(globalState) else iife96_()
            )
            nw126_[int(21)] = (d_726_v89_) | (d_726_v89_)
            nw126_[int(22)] = d_726_v89_
            nw126_[int(23)] = (d_726_v89_) | (d_726_v89_)
            nw126_[int(24)] = (d_726_v89_ if True else d_726_v89_)
            nw126_[int(25)] = d_726_v89_
            nw126_[int(26)] = d_726_v89_
            d_733_v98_ = nw126_
            index143_ = default__.safeIndex(266, (d_733_v98_).length(0))
            (d_733_v98_)[index143_] = d_726_v89_
            index144_ = default__.safeIndex(266, (d_733_v98_).length(0))
            (d_733_v98_)[index144_] = (d_729_v93_)[default__.safeIndex((d_724_v87_)[default__.safeIndex(579, (d_724_v87_).length(0))], len(d_729_v93_))]
        d_737_v99_: T0
        nw127_ = C1()
        nw127_.ctor__(p1, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "okusjfq"))), (self).f28)
        d_737_v99_ = nw127_
        d_738_v100_: _dafny.Map
        d_738_v100_ = _dafny.Map({p1: (d_737_v99_ if p0 else d_737_v99_)})
        d_738_v100_ = (d_738_v100_).set(p1, (d_737_v99_ if p0 else d_737_v99_))
        d_739_v101_: _dafny.MultiSet
        d_739_v101_ = _dafny.MultiSet([p1])
        hi3_ = p1
        for d_740_i9_ in range(((d_739_v101_)[p1] if (p1) in (d_739_v101_) else p1), hi3_):
            d_741_v102_: _dafny.Seq
            d_741_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bvleeo"))
            d_742_v103_: _dafny.Seq
            d_742_v103_ = _dafny.SeqWithoutIsStrInference([d_741_v102_])
            d_743_v104_: _dafny.Array
            nw128_ = _dafny.Array(None, 17)
            nw128_[int(0)] = (d_741_v102_) + (d_741_v102_)
            nw128_[int(1)] = (d_741_v102_) + (_dafny.SeqWithoutIsStrInference([d_624_v9_ for d_744_i10_ in range(default__.abs(469))]))
            nw128_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pqagahw"))
            nw128_[int(3)] = default__.fm3(globalState)
            nw128_[int(4)] = (d_742_v103_)[default__.safeIndex(p1, len(d_742_v103_))]
            nw128_[int(5)] = (D2_DC7(p0, p0, d_741_v102_)).cf15
            nw128_[int(6)] = default__.fm3(globalState)
            nw128_[int(7)] = (_dafny.SeqWithoutIsStrInference([d_624_v9_ for d_745_i11_ in range(default__.abs(445))])) + (d_741_v102_)
            nw128_[int(8)] = (d_741_v102_ if p0 else d_741_v102_)
            nw128_[int(9)] = d_741_v102_
            nw128_[int(10)] = d_741_v102_
            nw128_[int(11)] = d_741_v102_
            nw128_[int(12)] = d_741_v102_
            nw128_[int(13)] = d_741_v102_
            nw128_[int(14)] = (d_741_v102_ if p3 else _dafny.SeqWithoutIsStrInference([d_624_v9_ for d_746_i12_ in range(default__.abs(594))]))
            nw128_[int(15)] = (d_741_v102_) + (d_741_v102_)
            nw128_[int(16)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pyfkhq"))) + (d_741_v102_)
            d_743_v104_ = nw128_
            index145_ = default__.safeIndex(490, (d_743_v104_).length(0))
            (d_743_v104_)[index145_] = (d_741_v102_) + (d_741_v102_)
            index146_ = default__.safeIndex(490, (d_743_v104_).length(0))
            (d_743_v104_)[index146_] = (default__.fm3(globalState)) + (d_741_v102_)
            d_624_v9_ = d_624_v9_
            d_747_v105_: D4
            d_747_v105_ = D4_DC10(d_624_v9_)
            pat_let_tv23_ = d_624_v9_
            def iife97_(_pat_let25_0):
                def iife98_(d_748_dt__update__tmp_h2_):
                    def iife99_(_pat_let26_0):
                        def iife100_(d_749_dt__update_hcf19_h1_):
                            return D4_DC10(d_749_dt__update_hcf19_h1_)
                        return iife100_(_pat_let26_0)
                    return iife99_(pat_let_tv23_)
                return iife98_(_pat_let25_0)
            source11_ = iife97_(d_747_v105_)
            if source11_.is_DC11:
                d_750___mcc_h6_ = source11_.cf20
                d_751_cf20_ = d_750___mcc_h6_
                r0 = (d_740_i9_) - (p1)
                d_752_v106_: _dafny.Array
                nw129_ = _dafny.Array(None, 18)
                nw129_[int(0)] = p0
                nw129_[int(1)] = True
                nw129_[int(2)] = p0
                nw129_[int(3)] = p3
                nw129_[int(4)] = p3
                nw129_[int(5)] = p0
                nw129_[int(6)] = ((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))]) <= ((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))])
                nw129_[int(7)] = default__.fm4(globalState)
                nw129_[int(8)] = p0
                nw129_[int(9)] = (p3 if False else p3)
                nw129_[int(10)] = p3
                nw129_[int(11)] = (d_624_v9_) not in ((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))])
                nw129_[int(12)] = p0
                nw129_[int(13)] = ((d_737_v99_).f28).cf4
                nw129_[int(14)] = p3
                nw129_[int(15)] = not(p3)
                nw129_[int(16)] = (d_740_i9_) > (p1)
                nw129_[int(17)] = p3
                d_752_v106_ = nw129_
                index147_ = default__.safeIndex(318, (d_752_v106_).length(0))
                (d_752_v106_)[index147_] = (d_751_cf20_)[default__.safeIndex(d_740_i9_, len(d_751_cf20_))]
                index148_ = default__.safeIndex(318, (d_752_v106_).length(0))
                (d_752_v106_)[index148_] = p3
                d_753_v107_: _dafny.Array
                nw130_ = _dafny.Array(_dafny.CodePoint('D'), 29)
                d_753_v107_ = nw130_
                d_754_v108_: _dafny.Map
                d_754_v108_ = _dafny.Map({d_753_v107_: (d_752_v106_ if (d_752_v106_)[default__.safeIndex(318, (d_752_v106_).length(0))] else d_752_v106_)})
                rhs131_ = d_624_v9_
                rhs132_ = ((_dafny.Map({d_753_v107_: d_752_v106_})) | (d_754_v108_)) | (d_754_v108_)
                d_624_v9_ = rhs131_
                d_754_v108_ = rhs132_
                d_755_v109_: _dafny.Array
                nw131_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_755_v109_ = nw131_
                d_755_v109_ = d_755_v109_
            elif source11_.is_DC12:
                d_756___mcc_h7_ = source11_.cf21
                d_757___mcc_h8_ = source11_.cf22
                d_758_cf22_ = d_757___mcc_h8_
                d_759_cf21_ = d_756___mcc_h7_
                d_760_v110_: _dafny.Seq
                d_760_v110_ = _dafny.SeqWithoutIsStrInference([p0])
                d_761_v111_: _dafny.Seq
                d_761_v111_ = _dafny.SeqWithoutIsStrInference([len((self).f40)])
                d_762_v112_: _dafny.Array
                nw132_ = _dafny.Array(None, 17)
                nw132_[int(0)] = p0
                nw132_[int(1)] = d_759_cf21_
                nw132_[int(2)] = p3
                nw132_[int(3)] = p3
                nw132_[int(4)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "exijh")))) > (p1)
                nw132_[int(5)] = p0
                nw132_[int(6)] = not(d_759_cf21_)
                nw132_[int(7)] = p0
                nw132_[int(8)] = p0
                nw132_[int(9)] = (d_739_v101_) != (d_739_v101_)
                nw132_[int(10)] = (d_760_v110_)[default__.safeIndex((d_761_v111_)[default__.safeIndex(p1, len(d_761_v111_))], len(d_760_v110_))]
                nw132_[int(11)] = p3
                nw132_[int(12)] = (True) and (p3)
                nw132_[int(13)] = not(p0)
                nw132_[int(14)] = (d_737_v99_).fm7(True, p0, globalState)
                nw132_[int(15)] = p0
                nw132_[int(16)] = not(d_759_cf21_)
                d_762_v112_ = nw132_
                index149_ = default__.safeIndex(851, (d_762_v112_).length(0))
                (d_762_v112_)[index149_] = p0
                index150_ = default__.safeIndex(851, (d_762_v112_).length(0))
                rhs133_ = ((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))]) < ((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))])
                rhs134_ = default__.fm4(globalState)
                rhs135_ = d_624_v9_
                rhs136_ = d_740_i9_
                rhs137_ = ((not(p3)) or (d_759_cf21_) if d_759_cf21_ else p3)
                lhs114_ = globalState
                lhs115_ = globalState
                lhs116_ = d_762_v112_
                lhs117_ = default__.safeIndex(851, (d_762_v112_).length(0))
                lhs114_.f1 = rhs133_
                lhs115_.f1 = rhs134_
                d_624_v9_ = rhs135_
                r0 = rhs136_
                lhs116_[lhs117_] = rhs137_
                d_763_v114_: _dafny.MultiSet
                d_763_v114_ = _dafny.MultiSet([p0, p3, p3, p0, (d_762_v112_)[default__.safeIndex(851, (d_762_v112_).length(0))]])
                def iife101_():
                    coll47_ = _dafny.Map()
                    compr_47_: int
                    for compr_47_ in (_dafny.MultiSet([((d_763_v114_)[p3] if (p3) in (d_763_v114_) else d_740_i9_)])).Elements:
                        d_764_v113_: int = compr_47_
                        if (d_764_v113_) in (_dafny.MultiSet([((d_763_v114_)[p3] if (p3) in (d_763_v114_) else d_740_i9_)])):
                            coll47_[default__.safeModuloInt(d_764_v113_, d_740_i9_)] = 205
                    return _dafny.Map(coll47_)
                (globalState).f27 = ((self).fm8(iife101_()
                , not(p0), (d_762_v112_)[default__.safeIndex(851, (d_762_v112_).length(0))], globalState)) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gr")))
                (globalState).f14 = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fmaq")))) * (p1)
                d_765_v115_: _dafny.Array
                nw133_ = _dafny.Array(D2.default()(), 22)
                d_765_v115_ = nw133_
                d_766_v116_: D2
                d_766_v116_ = D2_DC6(d_762_v112_)
                index151_ = default__.safeIndex(957, (d_765_v115_).length(0))
                (d_765_v115_)[index151_] = d_766_v116_
                index152_ = default__.safeIndex(957, (d_765_v115_).length(0))
                (d_765_v115_)[index152_] = D2_DC6(d_762_v112_)
            elif source11_.is_DC10:
                d_767___mcc_h9_ = source11_.cf19
                d_768_cf19_ = d_767___mcc_h9_
                d_769_v117_: _dafny.Array
                nw134_ = _dafny.Array(False, 13)
                d_769_v117_ = nw134_
                d_770_v118_: _dafny.Array
                nw135_ = _dafny.Array(int(0), 28)
                d_770_v118_ = nw135_
                d_771_v119_: _dafny.Map
                d_771_v119_ = _dafny.Map({len((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))]): p3})
                d_772_v120_: _dafny.Map
                d_772_v120_ = _dafny.Map({len(d_771_v119_): not(True)})
                d_773_v121_: _dafny.Seq
                d_773_v121_ = _dafny.SeqWithoutIsStrInference([len(d_771_v119_)])
                d_774_v122_: _dafny.Map
                d_774_v122_ = _dafny.Map({p0: len(d_773_v121_)})
                default__.m0(d_769_v117_, d_770_v118_, (len((d_743_v104_)[default__.safeIndex(490, (d_743_v104_).length(0))])) * (len(d_773_v121_)), len(d_774_v122_), globalState)
                d_775_v123_: _dafny.Map
                d_775_v123_ = _dafny.Map({(self).f28: p3})
                d_776_v124_: _dafny.Map
                d_776_v124_ = _dafny.Map({d_775_v123_: (0) - (len(d_741_v102_))})
                d_776_v124_ = _dafny.Map({(d_775_v123_).set(D1_DC3(p3), p3): d_740_i9_})
                (globalState).f0 = (d_740_i9_ if p3 else p1)
                (globalState).f19 = d_740_i9_
            elif True:
                d_777___mcc_h10_ = source11_.cf23
                d_778_cf23_ = d_777___mcc_h10_
                d_779_v125_: _dafny.Map
                d_779_v125_ = _dafny.Map({p0: not (p0) or (p0)})
                d_779_v125_ = (d_779_v125_).set((p0) and (p0), p3)
                d_741_v102_ = d_741_v102_
                d_780_v126_: D0
                d_780_v126_ = D0_DC2(p1, (d_740_i9_ if False else d_740_i9_))
                d_781_v127_: _dafny.Seq
                d_781_v127_ = _dafny.SeqWithoutIsStrInference([p1])
                pat_let_tv24_ = d_781_v127_
                def iife102_(_pat_let27_0):
                    def iife103_(d_782_dt__update__tmp_h3_):
                        def iife104_(_pat_let28_0):
                            def iife105_(d_783_dt__update_hcf2_h0_):
                                return D0_DC2(d_783_dt__update_hcf2_h0_, (d_782_dt__update__tmp_h3_).cf3)
                            return iife105_(_pat_let28_0)
                        return iife104_((0) - ((0) - ((0) - (len(pat_let_tv24_)))))
                    return iife103_(_pat_let27_0)
                d_780_v126_ = iife102_(D0_DC2(d_740_i9_, len(d_781_v127_)))
                (globalState).f27 = True
            d_784_v128_: _dafny.Array
            def lambda30_(d_785_i13_):
                return (d_785_i13_) - (780)

            init16_ = lambda30_
            nw136_ = _dafny.Array(None, 25)
            for i0_16_ in range(nw136_.length(0)):
                nw136_[i0_16_] = init16_(i0_16_)
            d_784_v128_ = nw136_
            d_784_v128_ = d_784_v128_
        d_786_v129_: C1
        nw137_ = C1()
        nw137_.ctor__(p1, 108, (d_737_v99_).f28)
        d_786_v129_ = nw137_
        r0 = p1
        d_787_v130_: D8
        d_787_v130_ = D8_DC21(p3, default__.fm4(globalState), (d_739_v101_).cardinality)
        r1 = (d_787_v130_).cf34
        return r0, r1

    @property
    def f40(self):
        return self._f40
    @property
    def f39(self):
        return self._f39

class C3(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f28):
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        return not (False) or (False)

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('d')

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_788_v0_: bool
        d_788_v0_ = False
        d_789_v1_: str
        d_789_v1_ = _dafny.CodePoint('i')
        d_790_v2_: _dafny.Set
        d_790_v2_ = _dafny.Set({d_789_v1_})
        d_791_v3_: D4
        d_791_v3_ = D4_DC12(d_788_v0_, d_790_v2_)
        d_792_v4_: D4
        d_792_v4_ = D4_DC13(d_791_v3_)
        pat_let_tv25_ = d_789_v1_
        pat_let_tv26_ = d_789_v1_
        def lambda31_(source12_):
            if source12_.is_DC11:
                d_793___mcc_h0_ = source12_.cf20
                d_794_cf20_ = d_793___mcc_h0_
                return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaegn"))) + (_dafny.SeqWithoutIsStrInference([pat_let_tv25_ for d_795_i0_ in range(default__.abs(11))]))
            elif source12_.is_DC12:
                d_796___mcc_h1_ = source12_.cf21
                d_797___mcc_h2_ = source12_.cf22
                d_798_cf22_ = d_797___mcc_h2_
                d_799_cf21_ = d_796___mcc_h1_
                return _dafny.SeqWithoutIsStrInference([pat_let_tv26_ for d_800_i1_ in range(default__.abs(24))])
            elif source12_.is_DC10:
                d_801___mcc_h3_ = source12_.cf19
                d_802_cf19_ = d_801___mcc_h3_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bimyub"))
            elif True:
                d_803___mcc_h4_ = source12_.cf23
                d_804_cf23_ = d_803___mcc_h4_
                return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhsqnulxr"))

        r1 = lambda31_(d_792_v4_)
        d_805_v5_: _dafny.MultiSet
        d_805_v5_ = _dafny.MultiSet([d_788_v0_])
        d_806_v6_: int
        d_806_v6_ = 707
        d_807_v7_: _dafny.Map
        d_807_v7_ = _dafny.Map({((d_805_v5_)[d_788_v0_] if (d_788_v0_) in (d_805_v5_) else d_806_v6_): d_788_v0_})
        d_808_v8_: _dafny.Set
        d_808_v8_ = _dafny.Set({d_788_v0_, d_788_v0_})
        d_809_v9_: _dafny.Map
        d_809_v9_ = _dafny.Map({len(d_807_v7_): len(d_808_v8_)})
        d_810_v10_: _dafny.Seq
        d_810_v10_ = _dafny.SeqWithoutIsStrInference([d_806_v6_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vnr")))])
        d_811_v11_: _dafny.Seq
        d_811_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ggrj"))
        d_812_v12_: D1
        d_812_v12_ = D1_DC4(d_806_v6_, d_788_v0_)
        d_813_v13_: _dafny.Set
        d_813_v13_ = _dafny.Set({d_806_v6_, d_806_v6_})
        d_814_v14_: _dafny.Array
        nw138_ = _dafny.Array(None, 17)
        nw138_[int(0)] = (len(d_809_v9_)) in (d_810_v10_)
        nw138_[int(1)] = d_788_v0_
        nw138_[int(2)] = (len((d_811_v11_).set(default__.safeIndex(d_806_v6_, len(d_811_v11_)), d_789_v1_))) > (d_806_v6_)
        nw138_[int(3)] = d_788_v0_
        nw138_[int(4)] = d_788_v0_
        nw138_[int(5)] = (d_812_v12_).cf6
        nw138_[int(6)] = True
        nw138_[int(7)] = False
        nw138_[int(8)] = d_788_v0_
        nw138_[int(9)] = d_788_v0_
        nw138_[int(10)] = False
        nw138_[int(11)] = d_788_v0_
        nw138_[int(12)] = (d_813_v13_).issubset(d_813_v13_)
        nw138_[int(13)] = d_788_v0_
        nw138_[int(14)] = d_788_v0_
        nw138_[int(15)] = (d_806_v6_) == (d_806_v6_)
        nw138_[int(16)] = (d_810_v10_) <= (d_810_v10_)
        d_814_v14_ = nw138_
        d_814_v14_ = d_814_v14_
        d_815_v15_: _dafny.Array
        def lambda32_(d_816_v6_, d_817_v0_):
            def lambda33_(d_818_i3_):
                return (d_818_i3_) * (len(_dafny.Map({d_816_v6_: _dafny.SeqWithoutIsStrInference([d_817_v0_])})))

            return lambda33_

        init17_ = lambda32_(d_806_v6_, d_788_v0_)
        nw139_ = _dafny.Array(None, 25)
        for i0_17_ in range(nw139_.length(0)):
            nw139_[i0_17_] = init17_(i0_17_)
        d_815_v15_ = nw139_
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_815_v15_).length(0)):
            d_819_i2_: int = guard_loop_1_
            if (True) and (((0) <= (d_819_i2_)) and ((d_819_i2_) < ((d_815_v15_).length(0)))):
                (d_815_v15_)[(d_819_i2_)] = (d_819_i2_) * ((d_806_v6_) - (len(d_811_v11_)))
        d_788_v0_ = d_788_v0_
        d_820_i4_: int
        d_820_i4_ = 0
        with _dafny.label("9"):
            while not(not(True)):
                with _dafny.c_label("9"):
                    if (d_820_i4_) >= (100):
                        raise _dafny.Break("9")
                    d_820_i4_ = (d_820_i4_) + (1)
                    d_821_v16_: _dafny.Map
                    d_821_v16_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ot")): (D4_DC12(d_788_v0_, d_790_v2_)).cf21})
                    d_822_v17_: _dafny.Seq
                    d_822_v17_ = _dafny.SeqWithoutIsStrInference([d_788_v0_])
                    d_823_v18_: _dafny.Seq
                    d_823_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({((d_821_v16_)[d_811_v11_] if (d_811_v11_) in (d_821_v16_) else (d_822_v17_)[default__.safeIndex(818, len(d_822_v17_))]): d_788_v0_})])
                    rhs138_ = d_823_v18_
                    rhs139_ = d_814_v14_
                    rhs140_ = (d_788_v0_ if (649) > (d_806_v6_) else d_788_v0_)
                    lhs118_ = globalState
                    d_823_v18_ = rhs138_
                    d_814_v14_ = rhs139_
                    lhs118_.f11 = rhs140_
                    d_824_v19_: bool
                    out2_: bool
                    out2_ = (self).m13(229, globalState)
                    d_824_v19_ = out2_
                    d_825_v20_: C2
                    nw140_ = C2()
                    nw140_.ctor__(default__.fm22(d_788_v0_, d_806_v6_, len(d_808_v8_), globalState), d_808_v8_, (self).f28)
                    d_825_v20_ = nw140_
                    d_826_v21_: D8
                    d_826_v21_ = D8_DC21(d_824_v19_, d_788_v0_, d_806_v6_)
                    d_827_v22_: _dafny.Set
                    d_827_v22_ = _dafny.Set({d_826_v21_})
                    d_828_v23_: _dafny.Set
                    d_828_v23_ = d_827_v22_
                    d_829_v24_: _dafny.Map
                    d_829_v24_ = _dafny.Map({len(d_811_v11_): d_827_v22_})
                    d_830_v25_: _dafny.Map
                    d_830_v25_ = _dafny.Map({(self).fm7(d_824_v19_, not(d_824_v19_), globalState): d_827_v22_})
                    d_831_v27_: _dafny.Array
                    nw141_ = _dafny.Array(None, 28)
                    nw141_[int(0)] = (d_827_v22_ if not(not(d_824_v19_)) else d_827_v22_)
                    nw141_[int(1)] = (d_827_v22_).intersection(d_827_v22_)
                    nw141_[int(2)] = d_827_v22_
                    nw141_[int(3)] = ((_dafny.Set({d_826_v21_}))) | (d_827_v22_)
                    nw141_[int(4)] = d_827_v22_
                    nw141_[int(5)] = d_827_v22_
                    nw141_[int(6)] = _dafny.Set({D8_DC21(d_788_v0_, True, len(_dafny.Map({d_824_v19_: len(d_811_v11_)})))})
                    nw141_[int(7)] = d_827_v22_
                    nw141_[int(8)] = _dafny.Set({d_826_v21_})
                    nw141_[int(9)] = d_827_v22_
                    nw141_[int(10)] = ((d_829_v24_)[d_806_v6_] if (d_806_v6_) in (d_829_v24_) else default__.fm28(not(((d_807_v7_)[d_806_v6_] if (d_806_v6_) in (d_807_v7_) else (d_822_v17_)[default__.safeIndex(len(d_821_v16_), len(d_822_v17_))])), globalState))
                    nw141_[int(11)] = ((d_830_v25_)[d_824_v19_] if (d_824_v19_) in (d_830_v25_) else default__.fm28(False, globalState))
                    nw141_[int(12)] = _dafny.Set({d_826_v21_, d_826_v21_})
                    def iife106_():
                        coll48_ = _dafny.Set()
                        compr_48_: D8
                        for compr_48_ in (_dafny.MultiSet([D8_DC21(d_788_v0_, d_824_v19_, d_806_v6_), d_826_v21_, d_826_v21_])).Elements:
                            d_832_v26_: D8 = compr_48_
                            if (d_832_v26_) in (_dafny.MultiSet([D8_DC21(d_788_v0_, d_824_v19_, d_806_v6_), d_826_v21_, d_826_v21_])):
                                coll48_ = coll48_.union(_dafny.Set([d_832_v26_]))
                        return _dafny.Set(coll48_)
                    nw141_[int(13)] = iife106_()
                    
                    nw141_[int(14)] = d_827_v22_
                    nw141_[int(15)] = d_827_v22_
                    nw141_[int(16)] = d_827_v22_
                    nw141_[int(17)] = d_827_v22_
                    nw141_[int(18)] = (d_827_v22_) - (d_827_v22_)
                    nw141_[int(19)] = d_827_v22_
                    nw141_[int(20)] = d_827_v22_
                    nw141_[int(21)] = (d_827_v22_).intersection(d_827_v22_)
                    nw141_[int(22)] = default__.fm28(True, globalState)
                    nw141_[int(23)] = d_827_v22_
                    nw141_[int(24)] = d_827_v22_
                    nw141_[int(25)] = (d_827_v22_ if d_824_v19_ else d_827_v22_)
                    nw141_[int(26)] = (d_828_v23_)
                    nw141_[int(27)] = (d_827_v22_) | (d_827_v22_)
                    d_831_v27_ = nw141_
                    d_831_v27_ = d_831_v27_
                    pass
            pass
        d_788_v0_ = d_788_v0_
        d_833_v28_: _dafny.Seq
        d_833_v28_ = _dafny.SeqWithoutIsStrInference([d_815_v15_, d_815_v15_])
        d_834_v29_: D2
        d_834_v29_ = D2_DC8(d_788_v0_, (d_833_v28_)[default__.safeIndex(d_806_v6_, len(d_833_v28_))])
        r0 = d_834_v29_
        r1 = d_811_v11_
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        if (p0) and (p0):
            d_835_v0_: _dafny.Array
            nw142_ = _dafny.Array(int(0), 22)
            d_835_v0_ = nw142_
            d_835_v0_ = d_835_v0_
            (globalState).f27 = (p0 if p0 else p0)
            if (not (p0) or (p0)) or (p0):
                (globalState).f27 = p0
                d_836_v1_: _dafny.Array
                nw143_ = _dafny.Array(_dafny.Map({}), 25)
                d_836_v1_ = nw143_
                d_837_v2_: int
                d_837_v2_ = 130
                index153_ = default__.safeIndex(462, (d_836_v1_).length(0))
                (d_836_v1_)[index153_] = _dafny.Map({(0) - (d_837_v2_): p0})
                d_838_v3_: _dafny.Seq
                d_838_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "omvn"))
                d_839_v4_: _dafny.MultiSet
                d_839_v4_ = _dafny.MultiSet([d_838_v3_])
                d_840_v5_: _dafny.Map
                d_840_v5_ = _dafny.Map({((d_839_v4_)[(d_838_v3_).set(default__.safeIndex(-328, len(d_838_v3_)), p1)] if ((d_838_v3_).set(default__.safeIndex(-328, len(d_838_v3_)), p1)) in (d_839_v4_) else d_837_v2_): p0})
                index154_ = default__.safeIndex(462, (d_836_v1_).length(0))
                (d_836_v1_)[index154_] = d_840_v5_
                d_841_v6_: _dafny.Array
                nw144_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 23)
                d_841_v6_ = nw144_
                d_842_v7_: _dafny.Seq
                d_842_v7_ = _dafny.SeqWithoutIsStrInference([p0])
                d_843_v8_: _dafny.Seq
                d_843_v8_ = _dafny.SeqWithoutIsStrInference([d_842_v7_, d_842_v7_, d_842_v7_])
                d_844_v9_: _dafny.Seq
                d_844_v9_ = _dafny.SeqWithoutIsStrInference([d_837_v2_, d_837_v2_])
                d_845_v10_: _dafny.MultiSet
                d_845_v10_ = _dafny.MultiSet([p0])
                d_846_v11_: _dafny.Seq
                d_846_v11_ = _dafny.SeqWithoutIsStrInference([d_845_v10_])
                d_847_v12_: _dafny.Array
                nw145_ = _dafny.Array(None, 9)
                nw145_[int(0)] = p0
                nw145_[int(1)] = (p0) not in ((d_843_v8_)[default__.safeIndex(d_837_v2_, len(d_843_v8_))])
                nw145_[int(2)] = (d_837_v2_) in (d_844_v9_)
                nw145_[int(3)] = p0
                nw145_[int(4)] = p0
                nw145_[int(5)] = (not(p0)) and (p0)
                nw145_[int(6)] = p0
                nw145_[int(7)] = (d_837_v2_) not in ((d_836_v1_)[default__.safeIndex(462, (d_836_v1_).length(0))])
                nw145_[int(8)] = (_dafny.MultiSet([(self).fm7(p0, p0, globalState)])).issubset(((d_846_v11_)[default__.safeIndex(d_837_v2_, len(d_846_v11_))]).set(p0, default__.abs(d_837_v2_)))
                d_847_v12_ = nw145_
                index155_ = default__.safeIndex(114, (d_847_v12_).length(0))
                (d_847_v12_)[index155_] = p0
                d_848_v13_: D1
                d_848_v13_ = D1_DC4(d_837_v2_, p0)
                index156_ = default__.safeIndex(114, (d_847_v12_).length(0))
                rhs141_ = d_841_v6_
                rhs142_ = not((d_848_v13_).cf6)
                rhs143_ = (self).fm7(p0, p0, globalState)
                rhs144_ = p0
                lhs119_ = globalState
                lhs120_ = globalState
                lhs121_ = d_847_v12_
                lhs122_ = default__.safeIndex(114, (d_847_v12_).length(0))
                d_841_v6_ = rhs141_
                lhs119_.f27 = rhs142_
                lhs120_.f27 = rhs143_
                lhs121_[lhs122_] = rhs144_
                d_849_v14_: _dafny.Map
                d_849_v14_ = _dafny.Map({(d_847_v12_)[default__.safeIndex(114, (d_847_v12_).length(0))]: p0})
                d_850_v15_: bool
                out3_: bool
                out3_ = (self).m13(default__.safeModuloInt(d_837_v2_, len(d_849_v14_)), globalState)
                d_850_v15_ = out3_
                d_851_v16_: _dafny.Seq
                d_851_v16_ = _dafny.SeqWithoutIsStrInference([d_835_v0_, d_835_v0_])
                d_852_v17_: _dafny.Map
                d_852_v17_ = _dafny.Map({p0: d_835_v0_})
                d_853_v18_: D2
                d_853_v18_ = D2_DC8(p0, d_835_v0_)
                d_854_v19_: _dafny.Array
                nw146_ = _dafny.Array(None, 26)
                nw146_[int(0)] = d_835_v0_
                nw146_[int(1)] = d_835_v0_
                nw146_[int(2)] = d_835_v0_
                nw146_[int(3)] = (d_851_v16_)[default__.safeIndex(d_837_v2_, len(d_851_v16_))]
                nw146_[int(4)] = d_835_v0_
                nw146_[int(5)] = d_835_v0_
                nw146_[int(6)] = d_835_v0_
                nw146_[int(7)] = d_835_v0_
                nw146_[int(8)] = d_835_v0_
                nw146_[int(9)] = d_835_v0_
                nw146_[int(10)] = ((d_852_v17_)[(d_847_v12_)[default__.safeIndex(114, (d_847_v12_).length(0))]] if ((d_847_v12_)[default__.safeIndex(114, (d_847_v12_).length(0))]) in (d_852_v17_) else d_835_v0_)
                nw146_[int(11)] = ((d_852_v17_)[p0] if (p0) in (d_852_v17_) else d_835_v0_)
                nw146_[int(12)] = d_835_v0_
                nw146_[int(13)] = d_835_v0_
                nw146_[int(14)] = d_835_v0_
                nw146_[int(15)] = d_835_v0_
                nw146_[int(16)] = (d_853_v18_).cf17
                nw146_[int(17)] = d_835_v0_
                nw146_[int(18)] = d_835_v0_
                nw146_[int(19)] = d_835_v0_
                nw146_[int(20)] = d_835_v0_
                nw146_[int(21)] = d_835_v0_
                nw146_[int(22)] = d_835_v0_
                nw146_[int(23)] = d_835_v0_
                nw146_[int(24)] = d_835_v0_
                nw146_[int(25)] = d_835_v0_
                d_854_v19_ = nw146_
                d_854_v19_ = d_854_v19_
            elif True:
                d_855_v20_: _dafny.Array
                nw147_ = _dafny.Array(None, 9)
                nw147_[int(0)] = p0
                nw147_[int(1)] = True
                nw147_[int(2)] = p0
                nw147_[int(3)] = p0
                nw147_[int(4)] = p0
                nw147_[int(5)] = p0
                nw147_[int(6)] = p0
                nw147_[int(7)] = p0
                nw147_[int(8)] = p0
                d_855_v20_ = nw147_
                d_856_v21_: D2
                d_856_v21_ = D2_DC6(d_855_v20_)
                d_857_v22_: _dafny.Set
                d_857_v22_ = _dafny.Set({_dafny.CodePoint('p'), p1, p1, _dafny.CodePoint('f'), p1})
                d_858_v23_: D4
                d_858_v23_ = D4_DC12(p0, d_857_v22_)
                d_859_v24_: _dafny.Map
                d_859_v24_ = _dafny.Map({d_856_v21_: d_858_v23_})
                d_860_v25_: _dafny.Map
                d_860_v25_ = _dafny.Map({d_859_v24_: (p1) not in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pi")))})
                d_860_v25_ = (d_860_v25_).set(_dafny.Map({d_856_v21_: d_858_v23_}), p0)
                d_861_v26_: int
                d_861_v26_ = 806
                d_862_v27_: D1
                d_862_v27_ = D1_DC5(p0, p0, d_861_v26_, default__.fm4(globalState), p0)
                d_863_v28_: _dafny.Map
                d_863_v28_ = _dafny.Map({d_862_v27_: p0})
                d_864_v29_: _dafny.Set
                d_864_v29_ = _dafny.Set({p0, False, not(((d_863_v28_)[default__.fm29(d_861_v26_, globalState)] if (default__.fm29(d_861_v26_, globalState)) in (d_863_v28_) else p0))})
                d_865_v30_: D2
                d_865_v30_ = D2_DC8(p0, d_835_v0_)
                d_866_v31_: _dafny.Map
                d_866_v31_ = _dafny.Map({d_865_v30_: d_864_v29_})
                d_867_v32_: D10
                d_867_v32_ = D10_DC24(d_864_v29_)
                d_868_v33_: _dafny.Map
                d_868_v33_ = _dafny.Map({d_861_v26_: d_861_v26_})
                d_869_v34_: _dafny.Map
                d_869_v34_ = _dafny.Map({p0: (len(d_857_v22_)) * (len(d_868_v33_))})
                d_870_v35_: _dafny.Seq
                d_870_v35_ = _dafny.SeqWithoutIsStrInference([d_861_v26_, len(_dafny.SeqWithoutIsStrInference([745])), d_861_v26_, len(_dafny.Set({-677}))])
                pat_let_tv27_ = p0
                pat_let_tv28_ = p0
                def iife107_(_pat_let29_0):
                    def iife108_(d_871_dt__update__tmp_h1_):
                        def iife109_(_pat_let30_0):
                            def iife110_(d_872_dt__update_hcf16_h1_):
                                return D2_DC8(d_872_dt__update_hcf16_h1_, (d_871_dt__update__tmp_h1_).cf17)
                            return iife110_(_pat_let30_0)
                        return iife109_(pat_let_tv27_)
                    return iife108_(_pat_let29_0)
                def iife111_(_pat_let31_0):
                    def iife112_(d_873_dt__update__tmp_h0_):
                        def iife113_(_pat_let32_0):
                            def iife114_(d_874_dt__update_hcf16_h0_):
                                return D2_DC8(d_874_dt__update_hcf16_h0_, (d_873_dt__update__tmp_h0_).cf17)
                            return iife114_(_pat_let32_0)
                        return iife113_(pat_let_tv28_)
                    return iife112_(_pat_let31_0)
                rhs145_ = ((d_866_v31_)[iife107_(d_865_v30_)] if (iife111_(d_865_v30_)) in (d_866_v31_) else (d_867_v32_).cf38)
                rhs146_ = ((d_869_v34_)[p0] if (p0) in (d_869_v34_) else len((d_870_v35_) + (d_870_v35_)))
                rhs147_ = d_861_v26_
                rhs148_ = d_861_v26_
                lhs123_ = globalState
                lhs124_ = globalState
                lhs125_ = globalState
                d_864_v29_ = rhs145_
                lhs123_.f14 = rhs146_
                lhs124_.f24 = rhs147_
                lhs125_.f14 = rhs148_
                (globalState).f27 = p0
                d_875_v36_: _dafny.Seq
                d_875_v36_ = _dafny.SeqWithoutIsStrInference([d_855_v20_])
                d_876_v37_: _dafny.Array
                nw148_ = _dafny.Array(None, 19)
                nw148_[int(0)] = d_855_v20_
                nw148_[int(1)] = d_855_v20_
                nw148_[int(2)] = d_855_v20_
                nw148_[int(3)] = d_855_v20_
                nw148_[int(4)] = d_855_v20_
                nw148_[int(5)] = d_855_v20_
                nw148_[int(6)] = d_855_v20_
                nw148_[int(7)] = d_855_v20_
                nw148_[int(8)] = d_855_v20_
                nw148_[int(9)] = (D2_DC6(d_855_v20_)).cf12
                nw148_[int(10)] = d_855_v20_
                nw148_[int(11)] = d_855_v20_
                nw148_[int(12)] = d_855_v20_
                nw148_[int(13)] = d_855_v20_
                nw148_[int(14)] = d_855_v20_
                nw148_[int(15)] = d_855_v20_
                nw148_[int(16)] = (d_875_v36_)[default__.safeIndex(d_861_v26_, len(d_875_v36_))]
                nw148_[int(17)] = d_855_v20_
                nw148_[int(18)] = d_855_v20_
                d_876_v37_ = nw148_
                index157_ = default__.safeIndex(677, (d_876_v37_).length(0))
                (d_876_v37_)[index157_] = d_855_v20_
                index158_ = default__.safeIndex(677, (d_876_v37_).length(0))
                (d_876_v37_)[index158_] = (d_855_v20_ if p0 else (d_856_v21_).cf12)
                d_877_v38_: _dafny.Seq
                d_877_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mudnlnri"))
                (globalState).f19 = (10) - (len(d_877_v38_))
            (globalState).f27 = p0
            d_878_v39_: int
            d_878_v39_ = 122
            d_879_v40_: _dafny.Seq
            d_879_v40_ = _dafny.SeqWithoutIsStrInference([d_878_v39_, d_878_v39_])
            d_880_v41_: _dafny.Set
            d_880_v41_ = _dafny.Set({(d_879_v40_)[default__.safeIndex(d_878_v39_, len(d_879_v40_))]})
            d_881_v42_: C1
            nw149_ = C1()
            nw149_.ctor__(default__.safeDivisionInt(d_878_v39_, d_878_v39_), (d_878_v39_) * (d_878_v39_), default__.fm20(not(False), p0, d_878_v39_, d_880_v41_, globalState))
            d_881_v42_ = nw149_
        elif True:
            d_882_v43_: _dafny.Array
            def lambda34_(d_883_i0_):
                return False

            init18_ = lambda34_
            nw150_ = _dafny.Array(None, 18)
            for i0_18_ in range(nw150_.length(0)):
                nw150_[i0_18_] = init18_(i0_18_)
            d_882_v43_ = nw150_
            index159_ = default__.safeIndex(304, (d_882_v43_).length(0))
            (d_882_v43_)[index159_] = p0
            d_884_v44_: int
            d_884_v44_ = 184
            d_885_v45_: _dafny.Map
            d_885_v45_ = _dafny.Map({p0: p0})
            d_886_v46_: _dafny.MultiSet
            d_886_v46_ = _dafny.MultiSet([p0, p0, p0, p0, p0])
            d_887_v47_: _dafny.Seq
            d_887_v47_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            index160_ = default__.safeIndex(304, (d_882_v43_).length(0))
            rhs149_ = not(p0)
            rhs150_ = d_884_v44_
            rhs151_ = (not(p0) if ((d_885_v45_)[p0] if (p0) in (d_885_v45_) else p0) else p0)
            rhs152_ = d_884_v44_
            rhs153_ = (-888) <= (((d_886_v46_)[p0] if (p0) in (d_886_v46_) else (_dafny.MultiSet([len(d_887_v47_), d_884_v44_])).cardinality))
            lhs126_ = globalState
            lhs127_ = globalState
            lhs128_ = d_882_v43_
            lhs129_ = default__.safeIndex(304, (d_882_v43_).length(0))
            lhs130_ = globalState
            lhs131_ = globalState
            lhs126_.f27 = rhs149_
            lhs127_.f14 = rhs150_
            lhs128_[lhs129_] = rhs151_
            lhs130_.f24 = rhs152_
            lhs131_.f27 = rhs153_
            d_888_v48_: _dafny.Seq
            d_888_v48_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bksjwejlq"))
            d_888_v48_ = (d_888_v48_) + ((d_888_v48_) + (d_888_v48_))
            d_889_v49_: _dafny.Seq
            d_889_v49_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0})])
            d_890_v50_: _dafny.Map
            d_890_v50_ = _dafny.Map({d_889_v49_: len((_dafny.SeqWithoutIsStrInference([-605 for d_891_i1_ in range(default__.abs(307))])).set(default__.safeIndex(d_884_v44_, len(_dafny.SeqWithoutIsStrInference([-605 for d_892_i1_ in range(default__.abs(307))]))), d_884_v44_))})
            d_890_v50_ = (d_890_v50_).set(d_889_v49_, d_884_v44_)
            d_893_v51_: _dafny.Seq
            d_893_v51_ = _dafny.SeqWithoutIsStrInference([d_884_v44_])
            (globalState).f0 = (d_893_v51_)[default__.safeIndex((0) - (len(d_885_v45_)), len(d_893_v51_))]
            d_894_v52_: _dafny.Set
            d_894_v52_ = _dafny.Set({(d_882_v43_)[default__.safeIndex(304, (d_882_v43_).length(0))], (d_882_v43_)[default__.safeIndex(304, (d_882_v43_).length(0))]})
            d_895_v53_: _dafny.Map
            d_895_v53_ = _dafny.Map({(d_882_v43_)[default__.safeIndex(304, (d_882_v43_).length(0))]: p1})
            r0 = default__.safeModuloInt(default__.fm2(d_888_v48_, d_894_v52_, d_895_v53_, p0, globalState), d_884_v44_)
        d_896_i2_: int
        d_896_i2_ = 0
        with _dafny.label("10"):
            while p0:
                with _dafny.c_label("10"):
                    if (d_896_i2_) >= (100):
                        raise _dafny.Break("10")
                    d_896_i2_ = (d_896_i2_) + (1)
                    if not(p0):
                        d_897_v54_: int
                        d_897_v54_ = 318
                        d_898_v55_: _dafny.Seq
                        d_898_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "etkpecews"))
                        d_899_v56_: _dafny.Map
                        d_899_v56_ = _dafny.Map({d_897_v54_: d_898_v55_})
                        d_899_v56_ = (d_899_v56_).set(d_897_v54_, d_898_v55_)
                        d_900_v57_: _dafny.Array
                        nw151_ = _dafny.Array(False, 5)
                        d_900_v57_ = nw151_
                        d_900_v57_ = d_900_v57_
                        d_901_v58_: str
                        d_901_v58_ = _dafny.CodePoint('b')
                        d_901_v58_ = d_901_v58_
                        d_902_v59_: _dafny.MultiSet
                        d_902_v59_ = _dafny.MultiSet([d_901_v58_])
                        d_903_v60_: bool
                        out4_: bool
                        out4_ = (self).m13(default__.safeModuloInt(d_897_v54_, ((d_902_v59_)[p1] if (p1) in (d_902_v59_) else d_897_v54_)), globalState)
                        d_903_v60_ = out4_
                        d_901_v58_ = p1
                    elif True:
                        d_904_v61_: _dafny.Array
                        def lambda35_(d_905_p0_):
                            def lambda36_(d_906_i3_):
                                return (d_906_i3_) - ((_dafny.MultiSet([(_dafny.MultiSet([d_905_p0_, d_905_p0_, d_905_p0_])).cardinality])).cardinality)

                            return lambda36_

                        init19_ = lambda35_(p0)
                        nw152_ = _dafny.Array(None, 5)
                        for i0_19_ in range(nw152_.length(0)):
                            nw152_[i0_19_] = init19_(i0_19_)
                        d_904_v61_ = nw152_
                        d_907_v62_: int
                        d_907_v62_ = -654
                        index161_ = default__.safeIndex(524, (d_904_v61_).length(0))
                        (d_904_v61_)[index161_] = (d_907_v62_ if not(p0) else (0) - (d_907_v62_))
                        index162_ = default__.safeIndex(524, (d_904_v61_).length(0))
                        (d_904_v61_)[index162_] = d_907_v62_
                        d_908_v63_: _dafny.Seq
                        d_908_v63_ = _dafny.SeqWithoutIsStrInference([p0])
                        (globalState).f11 = not(((d_908_v63_) + (d_908_v63_)) <= (d_908_v63_))
                        r0 = d_907_v62_
                        (globalState).f0 = default__.safeDivisionInt(d_907_v62_, 717)
                        (globalState).f11 = p0
                    d_909_v64_: _dafny.Array
                    def lambda37_(d_910_i4_):
                        return (d_910_i4_) + (260)

                    init20_ = lambda37_
                    nw153_ = _dafny.Array(None, 28)
                    for i0_20_ in range(nw153_.length(0)):
                        nw153_[i0_20_] = init20_(i0_20_)
                    d_909_v64_ = nw153_
                    index163_ = default__.safeIndex(31, (d_909_v64_).length(0))
                    (d_909_v64_)[index163_] = 387
                    d_911_v65_: int
                    d_911_v65_ = 855
                    index164_ = default__.safeIndex(31, (d_909_v64_).length(0))
                    (d_909_v64_)[index164_] = default__.safeModuloInt(d_911_v65_, d_911_v65_)
                    d_912_v66_: _dafny.Array
                    nw154_ = _dafny.Array(False, 13)
                    d_912_v66_ = nw154_
                    index165_ = default__.safeIndex(329, (d_912_v66_).length(0))
                    (d_912_v66_)[index165_] = p0
                    d_913_v67_: _dafny.Seq
                    d_913_v67_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))
                    d_914_v68_: _dafny.Set
                    d_914_v68_ = _dafny.Set({d_913_v67_, d_913_v67_})
                    d_915_v69_: _dafny.MultiSet
                    d_915_v69_ = _dafny.MultiSet([(d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))]])
                    d_916_v70_: _dafny.Map
                    d_916_v70_ = _dafny.Map({(0) - ((d_915_v69_).cardinality): p0})
                    index166_ = default__.safeIndex(329, (d_912_v66_).length(0))
                    rhs154_ = p0
                    rhs155_ = ((d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))]) < (len(d_916_v70_))
                    rhs156_ = p0
                    rhs157_ = d_914_v68_
                    lhs132_ = globalState
                    lhs133_ = globalState
                    lhs134_ = d_912_v66_
                    lhs135_ = default__.safeIndex(329, (d_912_v66_).length(0))
                    lhs132_.f1 = rhs154_
                    lhs133_.f1 = rhs155_
                    lhs134_[lhs135_] = rhs156_
                    d_914_v68_ = rhs157_
                    if p0:
                        d_917_v71_: _dafny.Seq
                        d_917_v71_ = _dafny.SeqWithoutIsStrInference([d_909_v64_])
                        d_918_v72_: _dafny.Set
                        d_918_v72_ = _dafny.Set({True})
                        d_919_v73_: _dafny.Seq
                        d_919_v73_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_909_v64_, d_909_v64_])])
                        d_920_v74_: _dafny.Array
                        nw155_ = _dafny.Array(None, 24)
                        nw155_[int(0)] = _dafny.SeqWithoutIsStrInference([d_909_v64_, (d_917_v71_)[default__.safeIndex(len(d_918_v72_), len(d_917_v71_))]])
                        nw155_[int(1)] = d_917_v71_
                        nw155_[int(2)] = d_917_v71_
                        nw155_[int(3)] = d_917_v71_
                        nw155_[int(4)] = d_917_v71_
                        nw155_[int(5)] = (d_917_v71_) + (d_917_v71_)
                        nw155_[int(6)] = d_917_v71_
                        nw155_[int(7)] = d_917_v71_
                        nw155_[int(8)] = d_917_v71_
                        nw155_[int(9)] = (d_917_v71_).set(default__.safeIndex(d_911_v65_, len(d_917_v71_)), d_909_v64_)
                        nw155_[int(10)] = d_917_v71_
                        nw155_[int(11)] = (d_917_v71_) + (d_917_v71_)
                        nw155_[int(12)] = d_917_v71_
                        nw155_[int(13)] = ((d_917_v71_) + (d_917_v71_)).set(default__.safeIndex((0) - ((d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))]), len((d_917_v71_) + (d_917_v71_))), d_909_v64_)
                        nw155_[int(14)] = (d_919_v73_)[default__.safeIndex((d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))], len(d_919_v73_))]
                        nw155_[int(15)] = d_917_v71_
                        nw155_[int(16)] = _dafny.SeqWithoutIsStrInference([d_909_v64_])
                        nw155_[int(17)] = _dafny.SeqWithoutIsStrInference([d_909_v64_])
                        nw155_[int(18)] = (d_917_v71_) + (d_917_v71_)
                        nw155_[int(19)] = _dafny.SeqWithoutIsStrInference([d_909_v64_])
                        nw155_[int(20)] = d_917_v71_
                        nw155_[int(21)] = d_917_v71_
                        nw155_[int(22)] = d_917_v71_
                        nw155_[int(23)] = d_917_v71_
                        d_920_v74_ = nw155_
                        d_921_v75_: _dafny.Map
                        d_921_v75_ = _dafny.Map({p0: d_909_v64_})
                        index167_ = default__.safeIndex(474, (d_920_v74_).length(0))
                        (d_920_v74_)[index167_] = (_dafny.SeqWithoutIsStrInference([((d_921_v75_)[p0] if (p0) in (d_921_v75_) else d_909_v64_)])).set(default__.safeIndex(d_911_v65_, len(_dafny.SeqWithoutIsStrInference([((d_921_v75_)[p0] if (p0) in (d_921_v75_) else d_909_v64_)]))), d_909_v64_)
                        d_922_v76_: _dafny.Map
                        d_922_v76_ = _dafny.Map({p1: (d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))]})
                        index168_ = default__.safeIndex(474, (d_920_v74_).length(0))
                        rhs158_ = (len(d_922_v76_)) + ((d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))])
                        rhs159_ = d_917_v71_
                        lhs136_ = globalState
                        lhs137_ = d_920_v74_
                        lhs138_ = default__.safeIndex(474, (d_920_v74_).length(0))
                        lhs136_.f0 = rhs158_
                        lhs137_[lhs138_] = rhs159_
                        d_923_v77_: _dafny.Set
                        d_923_v77_ = _dafny.Set({(d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))]})
                        d_924_v78_: _dafny.Map
                        d_924_v78_ = _dafny.Map({(d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))]: len(d_923_v77_)})
                        d_925_v79_: _dafny.Map
                        d_925_v79_ = _dafny.Map({len(d_924_v78_): len(d_913_v67_)})
                        d_926_v80_: _dafny.Map
                        d_926_v80_ = _dafny.Map({(d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))]: d_925_v79_})
                        d_926_v80_ = (d_926_v80_).set(p0, d_925_v79_)
                        d_927_v81_: C1
                        nw156_ = C1()
                        nw156_.ctor__(default__.safeDivisionInt(default__.fm2((d_913_v67_).set(default__.safeIndex(d_911_v65_, len(d_913_v67_)), p1), d_918_v72_, _dafny.Map({p0: p1}), p0, globalState), d_911_v65_), (d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))], (self).f28)
                        d_927_v81_ = nw156_
                        d_909_v64_ = (d_909_v64_ if (self).fm7(False, p0, globalState) else d_909_v64_)
                        d_912_v66_ = (D2_DC6(d_912_v66_)).cf12
                    elif True:
                        d_928_v82_: _dafny.Map
                        d_928_v82_ = _dafny.Map({p1: (d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))]})
                        d_929_v83_: _dafny.MultiSet
                        d_930_v84_: int
                        d_931_v85_: _dafny.MultiSet
                        d_932_v86_: int
                        out5_: _dafny.MultiSet
                        out6_: int
                        out7_: _dafny.MultiSet
                        out8_: int
                        out5_, out6_, out7_, out8_ = (self).m14((d_928_v82_) | (d_928_v82_), globalState)
                        d_929_v83_ = out5_
                        d_930_v84_ = out6_
                        d_931_v85_ = out7_
                        d_932_v86_ = out8_
                        (globalState).f11 = (d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))]
                        d_933_v87_: _dafny.Seq
                        d_933_v87_ = _dafny.SeqWithoutIsStrInference([p0])
                        d_934_v88_: _dafny.Map
                        d_934_v88_ = _dafny.Map({d_933_v87_: 512})
                        d_935_v89_: _dafny.Map
                        d_935_v89_ = _dafny.Map({d_909_v64_: (d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))]})
                        d_936_v91_: T0
                        nw157_ = C2()
                        def iife115_():
                            coll49_ = _dafny.Set()
                            compr_49_: int
                            for compr_49_ in _dafny.IntegerRange(779, 149):
                                d_937_v90_: int = compr_49_
                                if ((779) <= (d_937_v90_)) and ((d_937_v90_) < (149)):
                                    coll49_ = coll49_.union(_dafny.Set([default__.safeDivisionInt(d_937_v90_, 6)]))
                            return _dafny.Set(coll49_)
                        nw157_.ctor__((d_934_v88_).set(d_933_v87_, (d_909_v64_)[default__.safeIndex(31, (d_909_v64_).length(0))]), default__.fm24(d_933_v87_, p0, globalState), default__.fm20((d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))], (d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))], ((d_935_v89_)[d_909_v64_] if (d_909_v64_) in (d_935_v89_) else d_930_v84_), iife115_()
                        , globalState))
                        d_936_v91_ = nw157_
                        d_936_v91_ = d_936_v91_
                        (globalState).f27 = (d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))]
                        d_909_v64_ = (d_909_v64_ if (d_912_v66_)[default__.safeIndex(329, (d_912_v66_).length(0))] else d_909_v64_)
                    pass
            pass
        d_938_v92_: _dafny.Seq
        d_938_v92_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wadqsyrq"))
        d_939_v93_: _dafny.Set
        d_939_v93_ = _dafny.Set({not(p0)})
        d_940_v94_: _dafny.Map
        d_940_v94_ = _dafny.Map({p0: p1})
        d_941_v95_: int
        d_941_v95_ = 783
        d_942_v96_: D0
        d_942_v96_ = D0_DC2(-756, d_941_v95_)
        d_943_v97_: _dafny.MultiSet
        d_943_v97_ = _dafny.MultiSet([p0, p0, p0])
        d_944_v98_: _dafny.Map
        d_944_v98_ = _dafny.Map({not(False): True})
        d_945_v99_: _dafny.Array
        nw158_ = _dafny.Array(None, 22)
        nw158_[int(0)] = p0
        nw158_[int(1)] = p0
        nw158_[int(2)] = not(p0)
        nw158_[int(3)] = False
        nw158_[int(4)] = p0
        nw158_[int(5)] = p0
        nw158_[int(6)] = p0
        nw158_[int(7)] = p0
        nw158_[int(8)] = p0
        nw158_[int(9)] = p0
        nw158_[int(10)] = p0
        nw158_[int(11)] = p0
        nw158_[int(12)] = p0
        nw158_[int(13)] = ((self).f28).cf4
        nw158_[int(14)] = True
        nw158_[int(15)] = p0
        nw158_[int(16)] = p0
        nw158_[int(17)] = p0
        nw158_[int(18)] = False
        nw158_[int(19)] = p0
        nw158_[int(20)] = ((d_944_v98_)[p0] if (p0) in (d_944_v98_) else p0)
        nw158_[int(21)] = True
        d_945_v99_ = nw158_
        d_946_v100_: _dafny.Map
        d_946_v100_ = _dafny.Map({d_945_v99_: d_941_v95_})
        d_947_v101_: _dafny.MultiSet
        d_947_v101_ = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([not(p0)])), default__.safeDivisionInt(default__.fm2(d_938_v92_, d_939_v93_, d_940_v94_, p0, globalState), d_941_v95_), (0) - ((0) - ((d_942_v96_).cf3)), (d_943_v97_).cardinality, len(d_946_v100_)])
        d_948_v102_: _dafny.Seq
        d_948_v102_ = _dafny.SeqWithoutIsStrInference([d_941_v95_])
        d_949_v103_: _dafny.Map
        d_949_v103_ = _dafny.Map({188: p0})
        rhs160_ = (default__.fm26(p0, (d_938_v92_).set(default__.safeIndex(d_941_v95_, len(d_938_v92_)), p1), p0, globalState)) | (_dafny.MultiSet((d_948_v102_) + (_dafny.SeqWithoutIsStrInference([d_941_v95_, d_941_v95_, d_941_v95_]))))
        rhs161_ = (len(d_949_v103_)) * (d_941_v95_)
        lhs139_ = globalState
        d_947_v101_ = rhs160_
        lhs139_.f0 = rhs161_
        if p0:
            if True:
                (globalState).f27 = (default__.fm2(d_938_v92_, d_939_v93_, d_940_v94_, p0, globalState)) < (d_941_v95_)
                d_950_v105_: _dafny.Map
                d_950_v105_ = (d_949_v103_).set(d_941_v95_, p0)
                d_951_v106_: _dafny.Seq
                d_951_v106_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({len(d_938_v92_): p0}), (d_950_v105_)])
                d_952_v107_: C1
                nw159_ = C1()
                def iife116_():
                    coll50_ = _dafny.Map()
                    compr_50_: _dafny.Map
                    for compr_50_ in (d_951_v106_).Elements:
                        d_953_v104_: _dafny.Map = compr_50_
                        if (d_953_v104_) in (d_951_v106_):
                            coll50_[d_953_v104_] = not(p0)
                    return _dafny.Map(coll50_)
                nw159_.ctor__((d_941_v95_) * (d_941_v95_), len(iife116_()
                ), (self).f28)
                d_952_v107_ = nw159_
                d_954_v108_: _dafny.Map
                d_954_v108_ = _dafny.Map({p0: d_941_v95_})
                d_955_v109_: C1
                nw160_ = C1()
                nw160_.ctor__(default__.safeModuloInt((d_952_v107_).f41, len(d_954_v108_)), (d_952_v107_).f41, (self).f28)
                d_955_v109_ = nw160_
                d_956_v110_: _dafny.Seq
                d_956_v110_ = _dafny.SeqWithoutIsStrInference([False, True, ((d_944_v98_)[True] if (True) in (d_944_v98_) else p0)])
                d_957_v111_: _dafny.Map
                d_957_v111_ = _dafny.Map({d_956_v110_: (d_952_v107_).f41})
                d_958_v112_: _dafny.Map
                d_958_v112_ = _dafny.Map({p0: d_957_v111_})
                d_959_v113_: C2
                nw161_ = C2()
                nw161_.ctor__(((d_958_v112_)[p0] if (p0) in (d_958_v112_) else d_957_v111_), d_939_v93_, (self).f28)
                d_959_v113_ = nw161_
                d_960_v114_: D2
                d_960_v114_ = D2_DC6(d_945_v99_)
                pat_let_tv29_ = d_945_v99_
                def iife117_(_pat_let33_0):
                    def iife118_(d_961_dt__update__tmp_h2_):
                        def iife119_(_pat_let34_0):
                            def iife120_(d_962_dt__update_hcf12_h0_):
                                return D2_DC6(d_962_dt__update_hcf12_h0_)
                            return iife120_(_pat_let34_0)
                        return iife119_(pat_let_tv29_)
                    return iife118_(_pat_let33_0)
                d_945_v99_ = (iife117_(d_960_v114_)).cf12
            elif True:
                (globalState).f27 = not (p0) or ((d_947_v101_).issubset(default__.fm26(p0, d_938_v92_, p0, globalState)))
                (globalState).f1 = p0
                (globalState).f27 = (d_941_v95_) <= (d_941_v95_)
                d_963_v115_: _dafny.Seq
                d_963_v115_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                d_964_v116_: _dafny.Seq
                d_964_v116_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_963_v115_: len(d_948_v102_)})])
                d_965_v117_: _dafny.Map
                d_965_v117_ = _dafny.Map({d_963_v115_: d_941_v95_})
                d_966_v118_: _dafny.Map
                d_966_v118_ = _dafny.Map({p0: d_939_v93_})
                pat_let_tv30_ = p0
                d_967_v119_: C2
                nw162_ = C2()
                def iife121_(_pat_let35_0):
                    def iife122_(d_968_dt__update__tmp_h3_):
                        def iife123_(_pat_let36_0):
                            def iife124_(d_969_dt__update_hcf4_h0_):
                                return D1_DC3(d_969_dt__update_hcf4_h0_)
                            return iife124_(_pat_let36_0)
                        return iife123_(pat_let_tv30_)
                    return iife122_(_pat_let35_0)
                nw162_.ctor__(((d_964_v116_)[default__.safeIndex(d_941_v95_, len(d_964_v116_))]) | (d_965_v117_), ((d_966_v118_)[True] if (True) in (d_966_v118_) else d_939_v93_), iife121_((self).f28))
                d_967_v119_ = nw162_
                d_970_v120_: str
                d_970_v120_ = _dafny.CodePoint('w')
                d_970_v120_ = p1
            if True:
                pat_let_tv31_ = p0
                d_971_v121_: C1
                nw163_ = C1()
                def iife125_(_pat_let37_0):
                    def iife126_(d_972_dt__update__tmp_h4_):
                        def iife127_(_pat_let38_0):
                            def iife128_(d_973_dt__update_hcf4_h1_):
                                return D1_DC3(d_973_dt__update_hcf4_h1_)
                            return iife128_(_pat_let38_0)
                        return iife127_(pat_let_tv31_)
                    return iife126_(_pat_let37_0)
                nw163_.ctor__(default__.safeModuloInt(d_941_v95_, d_941_v95_), d_941_v95_, iife125_((self).f28))
                d_971_v121_ = nw163_
                d_974_v122_: _dafny.Array
                nw164_ = _dafny.Array(int(0), 19)
                d_974_v122_ = nw164_
                d_975_v123_: _dafny.Map
                d_975_v123_ = _dafny.Map({d_974_v122_: not(p0)})
                d_975_v123_ = (d_975_v123_).set(d_974_v122_, (d_943_v97_).issubset(d_943_v97_))
                index169_ = default__.safeIndex(889, (d_974_v122_).length(0))
                (d_974_v122_)[index169_] = d_941_v95_
                index170_ = default__.safeIndex(889, (d_974_v122_).length(0))
                (d_974_v122_)[index170_] = default__.fm2(d_938_v92_, (d_939_v93_).intersection(d_939_v93_), d_940_v94_, True, globalState)
                d_976_v124_: D1
                d_976_v124_ = D1_DC4((d_974_v122_)[default__.safeIndex(889, (d_974_v122_).length(0))], p0)
                d_977_v125_: C1
                nw165_ = C1()
                nw165_.ctor__((d_976_v124_).cf5, d_971_v121_.f42, (self).f28)
                d_977_v125_ = nw165_
                d_978_v126_: C0
                nw166_ = C0()
                nw166_.ctor__(d_939_v93_)
                d_978_v126_ = nw166_
            elif True:
                d_979_v127_: _dafny.Seq
                d_979_v127_ = _dafny.SeqWithoutIsStrInference([default__.fm4(globalState)])
                d_980_v128_: _dafny.Seq
                d_980_v128_ = _dafny.SeqWithoutIsStrInference([d_979_v127_])
                (globalState).f11 = (len((d_938_v92_) + (d_938_v92_))) <= ((_dafny.MultiSet((d_979_v127_) + ((d_980_v128_)[default__.safeIndex(26, len(d_980_v128_))]))).cardinality)
                d_981_v129_: T0
                nw167_ = C1()
                nw167_.ctor__(len(d_949_v103_), d_941_v95_, (self).f28)
                d_981_v129_ = nw167_
                d_982_v130_: _dafny.Map
                d_982_v130_ = _dafny.Map({d_981_v129_: d_938_v92_})
                d_982_v130_ = (d_982_v130_).set(d_981_v129_, d_938_v92_)
                d_983_v131_: _dafny.Array
                nw168_ = _dafny.Array(int(0), 29)
                d_983_v131_ = nw168_
                index171_ = default__.safeIndex(719, (d_983_v131_).length(0))
                (d_983_v131_)[index171_] = (0) - ((d_941_v95_) + (d_941_v95_))
                index172_ = default__.safeIndex(719, (d_983_v131_).length(0))
                (d_983_v131_)[index172_] = d_941_v95_
                d_984_v132_: str
                d_984_v132_ = _dafny.CodePoint('c')
                d_984_v132_ = p1
                (globalState).f0 = (0) - ((d_983_v131_)[default__.safeIndex(719, (d_983_v131_).length(0))])
            d_943_v97_ = (d_943_v97_) | (d_943_v97_)
            (globalState).f19 = default__.fm2(d_938_v92_, (d_939_v93_) - (d_939_v93_), d_940_v94_, p0, globalState)
            if not(default__.fm4(globalState)):
                d_985_v133_: _dafny.Array
                nw169_ = _dafny.Array(int(0), 27)
                d_985_v133_ = nw169_
                index173_ = default__.safeIndex(100, (d_985_v133_).length(0))
                (d_985_v133_)[index173_] = 258
                d_986_v134_: _dafny.Map
                d_986_v134_ = _dafny.Map({p1: p0})
                index174_ = default__.safeIndex(151, (d_945_v99_).length(0))
                (d_945_v99_)[index174_] = (D7_DC18(p0)).cf29
                index175_ = default__.safeIndex(100, (d_985_v133_).length(0))
                index176_ = default__.safeIndex(151, (d_945_v99_).length(0))
                rhs162_ = d_941_v95_
                rhs163_ = d_941_v95_
                rhs164_ = ((_dafny.Map({p1: p0})) | (d_986_v134_)) | ((d_986_v134_ if p0 else d_986_v134_))
                rhs165_ = (d_943_v97_).ispropersubset((d_943_v97_) - (d_943_v97_))
                lhs140_ = globalState
                lhs141_ = d_985_v133_
                lhs142_ = default__.safeIndex(100, (d_985_v133_).length(0))
                lhs143_ = d_945_v99_
                lhs144_ = default__.safeIndex(151, (d_945_v99_).length(0))
                lhs140_.f19 = rhs162_
                lhs141_[lhs142_] = rhs163_
                d_986_v134_ = rhs164_
                lhs143_[lhs144_] = rhs165_
                (globalState).f24 = (d_985_v133_)[default__.safeIndex(100, (d_985_v133_).length(0))]
                (globalState).f24 = len(d_944_v98_)
                d_987_v135_: D1
                d_987_v135_ = D1_DC5(p0, p0, d_941_v95_, (d_945_v99_)[default__.safeIndex(151, (d_945_v99_).length(0))], (d_945_v99_)[default__.safeIndex(151, (d_945_v99_).length(0))])
                rhs166_ = ((d_945_v99_)[default__.safeIndex(151, (d_945_v99_).length(0))] if p0 else (d_945_v99_)[default__.safeIndex(151, (d_945_v99_).length(0))])
                rhs167_ = not(((d_987_v135_).cf10) == (p0))
                lhs145_ = globalState
                lhs146_ = globalState
                lhs145_.f1 = rhs166_
                lhs146_.f11 = rhs167_
                (globalState).f27 = p0
            elif True:
                d_988_v136_: _dafny.MultiSet
                d_988_v136_ = _dafny.MultiSet([d_943_v97_])
                d_989_v137_: _dafny.Array
                nw170_ = _dafny.Array(None, 13)
                nw170_[int(0)] = d_944_v98_
                nw170_[int(1)] = d_944_v98_
                nw170_[int(2)] = d_944_v98_
                nw170_[int(3)] = (d_944_v98_).set(False, p0)
                nw170_[int(4)] = d_944_v98_
                nw170_[int(5)] = d_944_v98_
                nw170_[int(6)] = d_944_v98_
                nw170_[int(7)] = d_944_v98_
                nw170_[int(8)] = d_944_v98_
                nw170_[int(9)] = _dafny.Map({p0: not(p0)})
                nw170_[int(10)] = d_944_v98_
                nw170_[int(11)] = d_944_v98_
                nw170_[int(12)] = d_944_v98_
                d_989_v137_ = nw170_
                d_990_v138_: _dafny.Array
                d_990_v138_ = d_989_v137_
                d_991_v139_: _dafny.Seq
                d_991_v139_ = _dafny.SeqWithoutIsStrInference([d_990_v138_])
                d_992_v140_: _dafny.Seq
                d_992_v140_ = _dafny.SeqWithoutIsStrInference([d_991_v139_, d_991_v139_, d_991_v139_, d_991_v139_, d_991_v139_])
                d_993_v141_: _dafny.Map
                d_993_v141_ = _dafny.Map({p1: d_945_v99_})
                d_994_v142_: D12
                d_994_v142_ = D12_DC27(d_993_v141_)
                d_995_v143_: _dafny.Map
                d_995_v143_ = _dafny.Map({d_941_v95_: d_941_v95_})
                d_996_v144_: _dafny.Seq
                d_996_v144_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vk")), d_938_v92_])
                d_997_v145_: _dafny.Seq
                d_997_v145_ = _dafny.SeqWithoutIsStrInference([p0])
                d_998_v146_: _dafny.Array
                nw171_ = _dafny.Array(None, 18)
                nw171_[int(0)] = len(d_995_v143_)
                nw171_[int(1)] = d_941_v95_
                nw171_[int(2)] = (0) - (len(d_996_v144_))
                nw171_[int(3)] = d_941_v95_
                nw171_[int(4)] = d_941_v95_
                nw171_[int(5)] = d_941_v95_
                nw171_[int(6)] = 228
                nw171_[int(7)] = d_941_v95_
                nw171_[int(8)] = d_941_v95_
                nw171_[int(9)] = len(d_939_v93_)
                nw171_[int(10)] = d_941_v95_
                nw171_[int(11)] = (_dafny.MultiSet([p0, p0, p0])).cardinality
                nw171_[int(12)] = d_941_v95_
                nw171_[int(13)] = -523
                nw171_[int(14)] = d_941_v95_
                nw171_[int(15)] = d_941_v95_
                nw171_[int(16)] = len(d_997_v145_)
                nw171_[int(17)] = (0) - ((d_947_v101_).cardinality)
                d_998_v146_ = nw171_
                d_999_v147_: D10
                d_999_v147_ = D10_DC25(d_998_v146_, p0, d_997_v145_, d_941_v95_)
                d_1000_v148_: _dafny.Array
                nw172_ = _dafny.Array(None, 22)
                nw172_[int(0)] = (d_941_v95_) * (d_941_v95_)
                nw172_[int(1)] = d_941_v95_
                nw172_[int(2)] = d_941_v95_
                nw172_[int(3)] = (0) - (default__.safeDivisionInt(len(d_938_v92_), len(d_939_v93_)))
                nw172_[int(4)] = d_941_v95_
                nw172_[int(5)] = d_941_v95_
                nw172_[int(6)] = (0) - (len((d_992_v140_)[default__.safeIndex(d_941_v95_, len(d_992_v140_))]))
                nw172_[int(7)] = default__.safeDivisionInt(default__.fm2(d_938_v92_, _dafny.Set({p0}), _dafny.Map({p0: p1}), p0, globalState), (0) - (d_941_v95_))
                nw172_[int(8)] = d_941_v95_
                nw172_[int(9)] = d_941_v95_
                nw172_[int(10)] = (d_941_v95_) * (d_941_v95_)
                nw172_[int(11)] = (d_941_v95_) * (-184)
                nw172_[int(12)] = len(((d_994_v142_).cf44) | (d_993_v141_))
                nw172_[int(13)] = d_941_v95_
                nw172_[int(14)] = d_941_v95_
                nw172_[int(15)] = d_941_v95_
                nw172_[int(16)] = len((_dafny.SeqWithoutIsStrInference([p1 for d_1001_i5_ in range(default__.abs(-763))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wbtg"))))
                nw172_[int(17)] = d_941_v95_
                nw172_[int(18)] = d_941_v95_
                nw172_[int(19)] = (d_999_v147_).cf42
                nw172_[int(20)] = len(d_997_v145_)
                nw172_[int(21)] = d_941_v95_
                d_1000_v148_ = nw172_
                rhs168_ = ((_dafny.MultiSet([d_943_v97_])).intersection((_dafny.MultiSet([_dafny.MultiSet(d_997_v145_), d_943_v97_, d_943_v97_, _dafny.MultiSet([p0])])).set(d_943_v97_, default__.abs(11)))) - (d_988_v136_)
                rhs169_ = d_998_v146_
                d_988_v136_ = rhs168_
                d_1000_v148_ = rhs169_
                (globalState).f0 = default__.fm2((_dafny.SeqWithoutIsStrInference([p1 for d_1002_i6_ in range(default__.abs(-901))])) + (d_938_v92_), (d_939_v93_) | (default__.fm24(d_997_v145_, p0, globalState)), (d_940_v94_) | (d_940_v94_), p0, globalState)
                (globalState).f14 = 731
                (globalState).f1 = p0
                d_1003_v149_: _dafny.Array
                nw173_ = _dafny.Array(D12.default()(), 8)
                d_1003_v149_ = nw173_
                d_1004_v150_: _dafny.Map
                d_1004_v150_ = _dafny.Map({p0: d_1003_v149_})
                d_1005_v151_: _dafny.Map
                d_1005_v151_ = d_949_v103_
                d_1006_v152_: _dafny.Map
                d_1006_v152_ = _dafny.Map({d_1005_v151_: (_dafny.SeqWithoutIsStrInference([269, (d_947_v101_).cardinality, d_941_v95_, d_941_v95_])) < (default__.fm6(p0, len(d_948_v102_), d_938_v92_, globalState))})
                d_1007_v153_: _dafny.Map
                d_1007_v153_ = _dafny.Map({d_997_v145_: d_1006_v152_})
                rhs170_ = (d_1004_v150_) | ((_dafny.Map({p0: d_1003_v149_})) | ((d_1004_v150_).set(p0, d_1003_v149_)))
                rhs171_ = ((d_1007_v153_)[d_997_v145_] if (d_997_v145_) in (d_1007_v153_) else d_1006_v152_)
                d_1004_v150_ = rhs170_
                d_1006_v152_ = rhs171_
        elif True:
            (globalState).f0 = (d_948_v102_)[default__.safeIndex(-164, len(d_948_v102_))]
            (globalState).f24 = d_941_v95_
            d_1008_v154_: _dafny.Map
            d_1008_v154_ = _dafny.Map({d_941_v95_: default__.safeDivisionInt(len(default__.fm3(globalState)), len(d_938_v92_))})
            d_1008_v154_ = (d_1008_v154_).set(d_941_v95_, d_941_v95_)
            d_1009_v155_: _dafny.Map
            d_1009_v155_ = _dafny.Map({d_938_v92_: p1})
            d_1009_v155_ = (d_1009_v155_).set(d_938_v92_, (d_938_v92_)[default__.safeIndex(d_941_v95_, len(d_938_v92_))])
            d_1010_v156_: _dafny.Array
            nw174_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_1010_v156_ = nw174_
            d_1011_v157_: _dafny.Array
            nw175_ = _dafny.Array(None, 10)
            nw175_[int(0)] = p1
            nw175_[int(1)] = _dafny.CodePoint('p')
            nw175_[int(2)] = p1
            nw175_[int(3)] = p1
            nw175_[int(4)] = p1
            nw175_[int(5)] = p1
            nw175_[int(6)] = p1
            nw175_[int(7)] = p1
            nw175_[int(8)] = p1
            nw175_[int(9)] = p1
            d_1011_v157_ = nw175_
            index177_ = default__.safeIndex(614, (d_1010_v156_).length(0))
            (d_1010_v156_)[index177_] = d_1011_v157_
            index178_ = default__.safeIndex(614, (d_1010_v156_).length(0))
            (d_1010_v156_)[index178_] = d_1011_v157_
        d_1012_v158_: _dafny.Seq
        d_1012_v158_ = _dafny.SeqWithoutIsStrInference([p0, p0])
        hi4_ = d_941_v95_
        for d_1013_i7_ in range(len(d_1012_v158_), hi4_):
            d_1014_v159_: str
            d_1014_v159_ = _dafny.CodePoint('i')
            d_1015_v160_: _dafny.Map
            d_1015_v160_ = _dafny.Map({d_948_v102_: d_1013_i7_})
            d_1016_v161_: _dafny.Map
            d_1016_v161_ = _dafny.Map({748: (d_1015_v160_).set(_dafny.SeqWithoutIsStrInference([d_941_v95_ for d_1017_i8_ in range(default__.abs(-993))]), (_dafny.MultiSet([d_1013_i7_])).cardinality)})
            d_1018_v163_: _dafny.Map
            d_1018_v163_ = _dafny.Map({(0) - (d_1013_i7_): d_948_v102_})
            d_1019_v164_: _dafny.Seq
            d_1019_v164_ = _dafny.SeqWithoutIsStrInference([((d_1018_v163_)[510] if (510) in (d_1018_v163_) else d_948_v102_), d_948_v102_])
            d_1020_v165_: _dafny.Map
            d_1020_v165_ = _dafny.Map({_dafny.CodePoint('i'): d_941_v95_})
            d_1021_v166_: _dafny.Array
            nw176_ = _dafny.Array(None, 13)
            nw176_[int(0)] = d_941_v95_
            nw176_[int(1)] = d_1013_i7_
            nw176_[int(2)] = (0) - (((d_943_v97_).intersection(d_943_v97_)).cardinality)
            nw176_[int(3)] = (0) - (d_941_v95_)
            def iife129_():
                coll51_ = _dafny.Map()
                compr_51_: _dafny.Seq
                for compr_51_ in (d_1019_v164_).Elements:
                    d_1022_v162_: _dafny.Seq = compr_51_
                    if (d_1022_v162_) in (d_1019_v164_):
                        coll51_[d_1022_v162_] = (D1_DC5(False, ((d_949_v103_)[d_941_v95_] if (d_941_v95_) in (d_949_v103_) else p0), d_1013_i7_, True, not(False))).cf9
                return _dafny.Map(coll51_)
            nw176_[int(4)] = len(((d_1016_v161_)[d_1013_i7_] if (d_1013_i7_) in (d_1016_v161_) else iife129_()
            ))
            nw176_[int(5)] = d_941_v95_
            nw176_[int(6)] = len(d_938_v92_)
            nw176_[int(7)] = (d_1013_i7_) - (len(d_938_v92_))
            nw176_[int(8)] = (len(d_1020_v165_)) * (d_1013_i7_)
            nw176_[int(9)] = (default__.fm2(d_938_v92_, d_939_v93_, d_940_v94_, p0, globalState)) - (len(d_938_v92_))
            nw176_[int(10)] = (d_943_v97_).cardinality
            nw176_[int(11)] = d_1013_i7_
            nw176_[int(12)] = (D8_DC20((0) - (d_941_v95_), d_941_v95_)).cf31
            d_1021_v166_ = nw176_
            index179_ = default__.safeIndex(261, (d_1021_v166_).length(0))
            (d_1021_v166_)[index179_] = -507
            index180_ = default__.safeIndex(261, (d_1021_v166_).length(0))
            rhs172_ = p1
            rhs173_ = default__.safeDivisionInt(d_941_v95_, 448)
            rhs174_ = p0
            lhs147_ = d_1021_v166_
            lhs148_ = default__.safeIndex(261, (d_1021_v166_).length(0))
            lhs149_ = globalState
            d_1014_v159_ = rhs172_
            lhs147_[lhs148_] = rhs173_
            lhs149_.f27 = rhs174_
            d_948_v102_ = d_948_v102_
            (globalState).f11 = p0
            d_1023_v167_: _dafny.Map
            d_1023_v167_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoscp")): p1})
            d_1024_v168_: _dafny.Map
            d_1024_v168_ = _dafny.Map({p0: (d_947_v101_).set((d_1021_v166_)[default__.safeIndex(261, (d_1021_v166_).length(0))], default__.abs((d_1021_v166_)[default__.safeIndex(261, (d_1021_v166_).length(0))]))})
            index181_ = default__.safeIndex(543, (d_945_v99_).length(0))
            (d_945_v99_)[index181_] = (_dafny.MultiSet([d_941_v95_, len(d_1023_v167_), len(d_1012_v158_), d_1013_i7_])).isdisjoint(((d_1024_v168_)[default__.fm4(globalState)] if (default__.fm4(globalState)) in (d_1024_v168_) else _dafny.MultiSet([d_1013_i7_])))
            d_1025_v169_: _dafny.Set
            d_1025_v169_ = _dafny.Set({d_1013_i7_, d_1013_i7_})
            d_1026_v170_: _dafny.Map
            d_1026_v170_ = _dafny.Map({d_941_v95_: len(d_1025_v169_)})
            d_1027_v171_: D2
            d_1027_v171_ = D2_DC7(p0, p0, _dafny.SeqWithoutIsStrInference([p1, p1, (self).fm8(d_1026_v170_, p0, p0, globalState), d_1014_v159_, p1]))
            index182_ = default__.safeIndex(543, (d_945_v99_).length(0))
            (d_945_v99_)[index182_] = (self).fm7((d_1027_v171_).cf14, True, globalState)
        d_1028_i9_: int
        d_1028_i9_ = 0
        with _dafny.label("11"):
            while p0:
                with _dafny.c_label("11"):
                    if (d_1028_i9_) >= (100):
                        raise _dafny.Break("11")
                    d_1028_i9_ = (d_1028_i9_) + (1)
                    d_1029_v172_: C1
                    nw177_ = C1()
                    nw177_.ctor__(d_941_v95_, (d_941_v95_) + (d_941_v95_), (self).f28)
                    d_1029_v172_ = nw177_
                    d_1030_v173_: _dafny.Map
                    d_1030_v173_ = _dafny.Map({(d_941_v95_) > (len(d_938_v92_)): d_1029_v172_.f42})
                    d_1031_v174_: _dafny.Set
                    d_1031_v174_ = _dafny.Set({(d_1029_v172_).f41, (0) - ((d_1029_v172_).f41)})
                    d_1032_v175_: _dafny.Set
                    d_1032_v175_ = _dafny.Set({d_1029_v172_.f42, len(d_1031_v174_), (d_1029_v172_).f41, len(_dafny.SeqWithoutIsStrInference([p0]))})
                    rhs175_ = default__.safeDivisionInt((d_1029_v172_).f41, (0) - (len((d_1032_v175_).intersection(d_1032_v175_))))
                    rhs176_ = d_1030_v173_
                    rhs177_ = (d_943_v97_).issubset((d_943_v97_).intersection(d_943_v97_))
                    rhs178_ = (0) - ((d_1029_v172_).f41)
                    lhs150_ = globalState
                    lhs151_ = globalState
                    lhs150_.f14 = rhs175_
                    d_1030_v173_ = rhs176_
                    lhs151_.f1 = rhs177_
                    d_941_v95_ = rhs178_
                    d_1033_v176_: _dafny.Array
                    nw178_ = _dafny.Array(None, 6)
                    nw178_[int(0)] = d_1029_v172_.f42
                    nw178_[int(1)] = (d_941_v95_) + (622)
                    nw178_[int(2)] = default__.fm2(d_938_v92_, d_939_v93_, d_940_v94_, True, globalState)
                    nw178_[int(3)] = d_1029_v172_.f42
                    nw178_[int(4)] = d_941_v95_
                    nw178_[int(5)] = (0) - (default__.safeModuloInt(((d_943_v97_)[p0] if (p0) in (d_943_v97_) else (d_1029_v172_).f41), 260))
                    d_1033_v176_ = nw178_
                    index183_ = default__.safeIndex(693, (d_1033_v176_).length(0))
                    (d_1033_v176_)[index183_] = d_941_v95_
                    index184_ = default__.safeIndex(693, (d_1033_v176_).length(0))
                    (d_1033_v176_)[index184_] = d_1029_v172_.f42
                    d_1034_v177_: D8
                    d_1034_v177_ = D8_DC21((p0 if p0 else p0), p0, (0) - (d_941_v95_))
                    d_1034_v177_ = D8_DC21(p0, (d_943_v97_) != (d_943_v97_), (d_1033_v176_)[default__.safeIndex(693, (d_1033_v176_).length(0))])
                    pass
            pass
        r0 = default__.safeDivisionInt(-604, d_941_v95_)
        return r0

    def m13(self, p0, globalState):
        r0: bool = False
        d_1035_v0_: bool
        d_1035_v0_ = False
        d_1036_v1_: _dafny.Set
        d_1036_v1_ = _dafny.Set({True, d_1035_v0_})
        hi5_ = len(d_1036_v1_)
        for d_1037_i0_ in range((0) - (p0), hi5_):
            d_1038_v2_: _dafny.Seq
            d_1038_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))
            d_1039_v3_: _dafny.Seq
            d_1039_v3_ = _dafny.SeqWithoutIsStrInference([len(d_1038_v2_), (0) - (len(d_1038_v2_)), d_1037_i0_])
            d_1040_v4_: _dafny.Set
            d_1040_v4_ = _dafny.Set({771})
            d_1041_v5_: _dafny.MultiSet
            d_1041_v5_ = _dafny.MultiSet([(d_1039_v3_) < (_dafny.SeqWithoutIsStrInference([len(d_1040_v4_), p0]))])
            rhs179_ = _dafny.MultiSet([d_1035_v0_, (d_1035_v0_) or (d_1035_v0_), not(d_1035_v0_), d_1035_v0_])
            rhs180_ = d_1035_v0_
            lhs152_ = globalState
            d_1041_v5_ = rhs179_
            lhs152_.f11 = rhs180_
            (globalState).f1 = ((0) - (d_1037_i0_)) >= (p0)
            d_1042_v6_: _dafny.Map
            d_1042_v6_ = _dafny.Map({d_1035_v0_: not(d_1035_v0_)})
            d_1042_v6_ = (d_1042_v6_).set(d_1035_v0_, d_1035_v0_)
            if d_1035_v0_:
                d_1043_v7_: _dafny.Map
                d_1043_v7_ = _dafny.Map({(0) - (len(d_1039_v3_)): d_1035_v0_})
                d_1044_v8_: _dafny.MultiSet
                d_1044_v8_ = _dafny.MultiSet([len(d_1043_v7_), d_1037_i0_])
                d_1045_v9_: _dafny.Map
                d_1045_v9_ = _dafny.Map({(d_1044_v8_).isdisjoint(_dafny.MultiSet([p0, p0])): default__.safeModuloInt(619, d_1037_i0_)})
                d_1046_v10_: _dafny.Array
                nw179_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
                d_1046_v10_ = nw179_
                d_1047_v11_: _dafny.Array
                nw180_ = _dafny.Array(False, 4)
                d_1047_v11_ = nw180_
                index185_ = default__.safeIndex(368, (d_1047_v11_).length(0))
                (d_1047_v11_)[index185_] = default__.fm4(globalState)
                index186_ = default__.safeIndex(346, (d_1047_v11_).length(0))
                (d_1047_v11_)[index186_] = not(d_1035_v0_)
                d_1048_v13_: _dafny.Map
                d_1048_v13_ = _dafny.Map({d_1038_v2_: d_1035_v0_})
                d_1049_v14_: _dafny.Seq
                d_1049_v14_ = _dafny.SeqWithoutIsStrInference([d_1036_v1_, d_1036_v1_])
                d_1050_v15_: str
                d_1050_v15_ = _dafny.CodePoint('e')
                d_1051_v16_: _dafny.Map
                d_1051_v16_ = _dafny.Map({d_1035_v0_: d_1050_v15_})
                d_1052_v17_: _dafny.Set
                d_1052_v17_ = _dafny.Set({d_1036_v1_, d_1036_v1_, d_1036_v1_, _dafny.Set({d_1035_v0_}), (d_1049_v14_)[default__.safeIndex(default__.fm2((d_1038_v2_).set(default__.safeIndex(p0, len(d_1038_v2_)), d_1050_v15_), d_1036_v1_, d_1051_v16_, True, globalState), len(d_1049_v14_))]})
                d_1053_v18_: _dafny.Seq
                d_1053_v18_ = _dafny.SeqWithoutIsStrInference([d_1052_v17_, _dafny.Set({d_1036_v1_})])
                d_1054_v19_: D8
                d_1054_v19_ = D8_DC21(d_1035_v0_, d_1035_v0_, d_1037_i0_)
                d_1055_v20_: _dafny.Seq
                d_1055_v20_ = _dafny.SeqWithoutIsStrInference([d_1035_v0_, d_1035_v0_, not((d_1054_v19_).cf33)])
                d_1056_v21_: _dafny.Seq
                d_1056_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_1035_v0_: d_1050_v15_})])
                index187_ = default__.safeIndex(368, (d_1047_v11_).length(0))
                index188_ = default__.safeIndex(346, (d_1047_v11_).length(0))
                def iife130_():
                    coll52_ = _dafny.Map()
                    compr_52_: _dafny.Seq
                    for compr_52_ in (d_1048_v13_).keys.Elements:
                        d_1057_v12_: _dafny.Seq = compr_52_
                        if (d_1057_v12_) in (d_1048_v13_):
                            coll52_[d_1057_v12_] = d_1035_v0_
                    return _dafny.Map(coll52_)
                rhs181_ = default__.fm30(d_1037_i0_, iife130_()
                , globalState)
                rhs182_ = (d_1046_v10_ if d_1035_v0_ else (d_1046_v10_ if d_1035_v0_ else d_1046_v10_))
                rhs183_ = ((d_1053_v18_)[default__.safeIndex(p0, len(d_1053_v18_))]).ispropersubset(d_1052_v17_)
                rhs184_ = ((d_1038_v2_ if d_1035_v0_ else d_1038_v2_)) + (((d_1038_v2_).set(default__.safeIndex(d_1037_i0_, len(d_1038_v2_)), d_1050_v15_)) + (d_1038_v2_))
                rhs185_ = (d_1055_v20_)[default__.safeIndex(default__.fm2(d_1038_v2_, d_1036_v1_, (d_1056_v21_)[default__.safeIndex(d_1037_i0_, len(d_1056_v21_))], (d_1055_v20_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1037_i0_ for d_1058_i1_ in range(default__.abs(-77))])), len(d_1055_v20_))], globalState), len(d_1055_v20_))]
                lhs153_ = d_1047_v11_
                lhs154_ = default__.safeIndex(368, (d_1047_v11_).length(0))
                lhs155_ = d_1047_v11_
                lhs156_ = default__.safeIndex(346, (d_1047_v11_).length(0))
                d_1045_v9_ = rhs181_
                d_1046_v10_ = rhs182_
                lhs153_[lhs154_] = rhs183_
                d_1038_v2_ = rhs184_
                lhs155_[lhs156_] = rhs185_
                d_1059_v22_: _dafny.Map
                d_1059_v22_ = _dafny.Map({d_1050_v15_: d_1037_i0_})
                d_1060_v23_: C1
                nw181_ = C1()
                nw181_.ctor__(len(d_1038_v2_), (p0) * (len(((d_1059_v22_).set(d_1050_v15_, (0) - (p0))).set(d_1050_v15_, d_1037_i0_))), default__.fm20(d_1035_v0_, (d_1047_v11_)[default__.safeIndex(368, (d_1047_v11_).length(0))], (0) - (p0), d_1040_v4_, globalState))
                d_1060_v23_ = nw181_
                d_1060_v23_ = d_1060_v23_
                d_1061_v24_: C1
                nw182_ = C1()
                nw182_.ctor__(-976, d_1037_i0_, (self).f28)
                d_1061_v24_ = nw182_
                (globalState).f14 = default__.safeModuloInt(default__.safeDivisionInt(d_1037_i0_, d_1061_v24_.f42), p0)
                rhs186_ = (d_1047_v11_)[default__.safeIndex(368, (d_1047_v11_).length(0))]
                rhs187_ = default__.fm4(globalState)
                rhs188_ = d_1060_v23_.f42
                rhs189_ = d_1035_v0_
                lhs157_ = globalState
                lhs158_ = globalState
                lhs159_ = d_1060_v23_
                lhs160_ = globalState
                lhs157_.f27 = rhs186_
                lhs158_.f11 = rhs187_
                lhs159_.f42 = rhs188_
                lhs160_.f27 = rhs189_
            elif True:
                (globalState).f0 = d_1037_i0_
                (globalState).f1 = d_1035_v0_
                (globalState).f14 = (0) - ((0) - ((len(d_1038_v2_)) + (d_1037_i0_)))
                (globalState).f11 = d_1035_v0_
                d_1062_v25_: _dafny.Seq
                d_1062_v25_ = _dafny.SeqWithoutIsStrInference([d_1035_v0_])
                d_1063_v26_: _dafny.Map
                d_1063_v26_ = _dafny.Map({d_1062_v25_: d_1037_i0_})
                d_1064_v27_: C2
                nw183_ = C2()
                nw183_.ctor__(d_1063_v26_, _dafny.Set({(d_1062_v25_)[default__.safeIndex((d_1039_v3_)[default__.safeIndex(p0, len(d_1039_v3_))], len(d_1062_v25_))], not(d_1035_v0_)}), (self).f28)
                d_1064_v27_ = nw183_
        d_1065_v28_: _dafny.Array
        def lambda38_(d_1066_i2_):
            return default__.safeDivisionInt(d_1066_i2_, -217)

        init21_ = lambda38_
        nw184_ = _dafny.Array(None, 23)
        for i0_21_ in range(nw184_.length(0)):
            nw184_[i0_21_] = init21_(i0_21_)
        d_1065_v28_ = nw184_
        index189_ = default__.safeIndex(776, (d_1065_v28_).length(0))
        (d_1065_v28_)[index189_] = 91
        d_1067_v29_: _dafny.Array
        def lambda39_(d_1068_v0_):
            def lambda40_(d_1069_i3_):
                return d_1068_v0_

            return lambda40_

        init22_ = lambda39_(d_1035_v0_)
        nw185_ = _dafny.Array(None, 3)
        for i0_22_ in range(nw185_.length(0)):
            nw185_[i0_22_] = init22_(i0_22_)
        d_1067_v29_ = nw185_
        d_1070_v30_: _dafny.Set
        d_1070_v30_ = _dafny.Set({d_1067_v29_})
        d_1071_v31_: _dafny.Array
        nw186_ = _dafny.Array(None, 2)
        nw186_[int(0)] = d_1070_v30_
        nw186_[int(1)] = d_1070_v30_
        d_1071_v31_ = nw186_
        index190_ = default__.safeIndex(287, (d_1071_v31_).length(0))
        (d_1071_v31_)[index190_] = d_1070_v30_
        d_1072_v32_: _dafny.Map
        d_1072_v32_ = _dafny.Map({(p0) + (p0): _dafny.CodePoint('q')})
        index191_ = default__.safeIndex(776, (d_1065_v28_).length(0))
        index192_ = default__.safeIndex(287, (d_1071_v31_).length(0))
        rhs190_ = p0
        rhs191_ = ((d_1070_v30_).intersection(d_1070_v30_) if d_1035_v0_ else (d_1070_v30_) | (d_1070_v30_))
        rhs192_ = p0
        rhs193_ = d_1072_v32_
        lhs161_ = d_1065_v28_
        lhs162_ = default__.safeIndex(776, (d_1065_v28_).length(0))
        lhs163_ = d_1071_v31_
        lhs164_ = default__.safeIndex(287, (d_1071_v31_).length(0))
        lhs165_ = globalState
        lhs161_[lhs162_] = rhs190_
        lhs163_[lhs164_] = rhs191_
        lhs165_.f19 = rhs192_
        d_1072_v32_ = rhs193_
        d_1073_v33_: str
        d_1073_v33_ = _dafny.CodePoint('g')
        d_1074_v34_: _dafny.Seq
        d_1074_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxqeqvw"))
        d_1075_v35_: _dafny.Map
        d_1075_v35_ = _dafny.Map({d_1035_v0_: d_1073_v33_})
        hi6_ = default__.fm2(d_1074_v34_, d_1036_v1_, d_1075_v35_, default__.fm4(globalState), globalState)
        for d_1076_i4_ in range(default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gcni")), d_1036_v1_, _dafny.Map({d_1035_v0_: d_1073_v33_}), d_1035_v0_, globalState), hi6_):
            d_1077_v36_: _dafny.Seq
            d_1077_v36_ = _dafny.SeqWithoutIsStrInference([d_1035_v0_])
            d_1078_v37_: C2
            nw187_ = C2()
            nw187_.ctor__(_dafny.Map({d_1077_v36_: d_1076_i4_}), d_1036_v1_, D1_DC3(d_1035_v0_))
            d_1078_v37_ = nw187_
            index193_ = default__.safeIndex(776, (d_1065_v28_).length(0))
            (d_1065_v28_)[index193_] = (0) - (p0)
            d_1079_v38_: _dafny.Map
            d_1079_v38_ = _dafny.Map({d_1035_v0_: p0})
            d_1080_v39_: _dafny.MultiSet
            d_1080_v39_ = _dafny.MultiSet([len(d_1079_v38_)])
            source13_ = D1_DC4(default__.safeModuloInt(p0, (d_1080_v39_).cardinality), d_1035_v0_)
            if source13_.is_DC4:
                d_1081___mcc_h0_ = source13_.cf5
                d_1082___mcc_h1_ = source13_.cf6
                d_1083_cf6_ = d_1082___mcc_h1_
                d_1084_cf5_ = d_1081___mcc_h0_
                rhs194_ = (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]
                rhs195_ = d_1074_v34_
                lhs166_ = globalState
                lhs166_.f14 = rhs194_
                d_1074_v34_ = rhs195_
                d_1085_v40_: _dafny.Map
                d_1085_v40_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqrd")): d_1083_cf6_})
                d_1086_v41_: C1
                nw188_ = C1()
                nw188_.ctor__(default__.safeModuloInt(d_1084_cf5_, len(d_1085_v40_)), p0, (self).f28)
                d_1086_v41_ = nw188_
                r0 = d_1083_cf6_
                (globalState).f19 = (0) - (default__.safeDivisionInt(d_1084_cf5_, (688) + (d_1084_cf5_)))
            elif source13_.is_DC5:
                d_1087___mcc_h2_ = source13_.cf7
                d_1088___mcc_h3_ = source13_.cf8
                d_1089___mcc_h4_ = source13_.cf9
                d_1090___mcc_h5_ = source13_.cf10
                d_1091___mcc_h6_ = source13_.cf11
                d_1092_cf11_ = d_1091___mcc_h6_
                d_1093_cf10_ = d_1090___mcc_h5_
                d_1094_cf9_ = d_1089___mcc_h4_
                d_1095_cf8_ = d_1088___mcc_h3_
                d_1096_cf7_ = d_1087___mcc_h2_
                index194_ = default__.safeIndex(825, (d_1067_v29_).length(0))
                (d_1067_v29_)[index194_] = (d_1077_v36_)[default__.safeIndex(-851, len(d_1077_v36_))]
                d_1097_v42_: C0
                nw189_ = C0()
                nw189_.ctor__(_dafny.Set({False}))
                d_1097_v42_ = nw189_
                d_1098_v43_: _dafny.Seq
                d_1098_v43_ = _dafny.SeqWithoutIsStrInference([d_1097_v42_, d_1097_v42_])
                d_1099_v44_: _dafny.Array
                nw190_ = _dafny.Array(None, 22)
                nw190_[int(0)] = d_1079_v38_
                nw190_[int(1)] = (d_1079_v38_) | (d_1079_v38_)
                nw190_[int(2)] = (d_1079_v38_) | (d_1079_v38_)
                nw190_[int(3)] = _dafny.Map({d_1093_cf10_: d_1094_cf9_})
                nw190_[int(4)] = d_1079_v38_
                nw190_[int(5)] = d_1079_v38_
                nw190_[int(6)] = d_1079_v38_
                nw190_[int(7)] = d_1079_v38_
                nw190_[int(8)] = d_1079_v38_
                nw190_[int(9)] = (d_1079_v38_) | (d_1079_v38_)
                nw190_[int(10)] = d_1079_v38_
                nw190_[int(11)] = d_1079_v38_
                nw190_[int(12)] = d_1079_v38_
                nw190_[int(13)] = d_1079_v38_
                nw190_[int(14)] = d_1079_v38_
                nw190_[int(15)] = (_dafny.Map({d_1092_cf11_: len(d_1098_v43_)})) | (d_1079_v38_)
                nw190_[int(16)] = d_1079_v38_
                nw190_[int(17)] = d_1079_v38_
                nw190_[int(18)] = d_1079_v38_
                nw190_[int(19)] = (d_1079_v38_) | (d_1079_v38_)
                nw190_[int(20)] = ((d_1079_v38_).set(not(d_1035_v0_), (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))])) | (d_1079_v38_)
                nw190_[int(21)] = (_dafny.Map({d_1095_cf8_: d_1076_i4_})) | (d_1079_v38_)
                d_1099_v44_ = nw190_
                d_1100_v45_: _dafny.Map
                d_1100_v45_ = _dafny.Map({d_1096_cf7_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_1101_i5_ in range(default__.abs(397))])})
                d_1102_v46_: _dafny.MultiSet
                d_1102_v46_ = _dafny.MultiSet([d_1035_v0_, d_1035_v0_, d_1092_cf11_])
                index195_ = default__.safeIndex(825, (d_1067_v29_).length(0))
                rhs196_ = ((_dafny.MultiSet([d_1093_cf10_])) | (d_1102_v46_)).isdisjoint(_dafny.MultiSet(d_1077_v36_))
                rhs197_ = d_1099_v44_
                rhs198_ = d_1100_v45_
                rhs199_ = d_1096_cf7_
                rhs200_ = False
                lhs167_ = d_1067_v29_
                lhs168_ = default__.safeIndex(825, (d_1067_v29_).length(0))
                lhs169_ = globalState
                lhs170_ = globalState
                lhs167_[lhs168_] = rhs196_
                d_1099_v44_ = rhs197_
                d_1100_v45_ = rhs198_
                lhs169_.f11 = rhs199_
                lhs170_.f1 = rhs200_
                d_1092_cf11_ = (self).fm7((d_1076_i4_) in (d_1080_v39_), d_1092_cf11_, globalState)
                d_1103_v47_: _dafny.Map
                d_1103_v47_ = _dafny.Map({d_1074_v34_: d_1077_v36_})
                d_1103_v47_ = (d_1103_v47_).set(d_1074_v34_, d_1077_v36_)
                index196_ = default__.safeIndex(825, (d_1067_v29_).length(0))
                (d_1067_v29_)[index196_] = d_1092_cf11_
            elif True:
                d_1104___mcc_h7_ = source13_.cf4
                d_1105_cf4_ = d_1104___mcc_h7_
                (globalState).f27 = not((d_1077_v36_)[default__.safeIndex(p0, len(d_1077_v36_))])
                d_1106_v48_: _dafny.Array
                nw191_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1106_v48_ = nw191_
                index197_ = default__.safeIndex(76, (d_1106_v48_).length(0))
                (d_1106_v48_)[index197_] = d_1065_v28_
                index198_ = default__.safeIndex(76, (d_1106_v48_).length(0))
                nw192_ = _dafny.Array(int(0), 16)
                (d_1106_v48_)[index198_] = nw192_
                (globalState).f19 = default__.fm2((d_1074_v34_) + (d_1074_v34_), d_1036_v1_, d_1075_v35_, ((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]) < (p0), globalState)
                d_1107_v49_: _dafny.Set
                d_1107_v49_ = _dafny.Set({p0})
                d_1108_v50_: _dafny.Seq
                d_1108_v50_ = _dafny.SeqWithoutIsStrInference([d_1107_v49_])
                d_1109_v51_: _dafny.Seq
                d_1109_v51_ = d_1108_v50_
                d_1110_v52_: _dafny.Array
                nw193_ = _dafny.Array(_dafny.Set({}), 14)
                d_1110_v52_ = nw193_
                d_1111_v53_: _dafny.Set
                d_1111_v53_ = _dafny.Set({(d_1106_v48_)[default__.safeIndex(76, (d_1106_v48_).length(0))], (d_1106_v48_)[default__.safeIndex(76, (d_1106_v48_).length(0))]})
                index199_ = default__.safeIndex(960, (d_1110_v52_).length(0))
                (d_1110_v52_)[index199_] = d_1111_v53_
                d_1112_v54_: _dafny.MultiSet
                d_1112_v54_ = _dafny.MultiSet([d_1107_v49_])
                d_1113_v55_: _dafny.Map
                d_1113_v55_ = _dafny.Map({(d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]: d_1107_v49_})
                d_1114_v56_: _dafny.Seq
                d_1114_v56_ = _dafny.SeqWithoutIsStrInference([p0, d_1076_i4_])
                index200_ = default__.safeIndex(960, (d_1110_v52_).length(0))
                rhs201_ = d_1109_v51_
                rhs202_ = (d_1105_cf4_ if d_1035_v0_ else (d_1076_i4_) > ((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]))
                rhs203_ = (d_1111_v53_) - (d_1111_v53_)
                rhs204_ = ((((d_1112_v54_).set(((d_1113_v55_)[(d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]] if ((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]) in (d_1113_v55_) else d_1107_v49_), default__.abs(d_1076_i4_)) if d_1035_v0_ else d_1112_v54_)).intersection((d_1112_v54_).intersection(default__.fm31((d_1078_v37_).fm17(d_1114_v56_, (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))], globalState), (0) - (p0), d_1105_cf4_, default__.fm4(globalState), globalState)))).cardinality
                lhs171_ = globalState
                lhs172_ = d_1110_v52_
                lhs173_ = default__.safeIndex(960, (d_1110_v52_).length(0))
                lhs174_ = globalState
                d_1109_v51_ = rhs201_
                lhs171_.f11 = rhs202_
                lhs172_[lhs173_] = rhs203_
                lhs174_.f24 = rhs204_
            (globalState).f27 = default__.fm4(globalState)
        if d_1035_v0_:
            d_1115_v57_: _dafny.Seq
            d_1115_v57_ = _dafny.SeqWithoutIsStrInference([False])
            d_1116_v58_: _dafny.Map
            d_1116_v58_ = _dafny.Map({d_1115_v57_: (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]})
            d_1117_v59_: C2
            nw194_ = C2()
            nw194_.ctor__(d_1116_v58_, (d_1036_v1_).intersection(d_1036_v1_), (self).f28)
            d_1117_v59_ = nw194_
            def iife131_():
                coll53_ = _dafny.Map()
                compr_53_: int
                for compr_53_ in _dafny.IntegerRange(855, -148):
                    d_1118_v60_: int = compr_53_
                    if ((855) <= (d_1118_v60_)) and ((d_1118_v60_) < (-148)):
                        coll53_[(d_1118_v60_) - (len(_dafny.SeqWithoutIsStrInference([d_1073_v33_ for d_1119_i6_ in range(default__.abs(698))])))] = _dafny.Map({d_1035_v0_: (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]})
                return _dafny.Map(coll53_)
            source14_ = default__.fm32(iife131_()
            , globalState)
            if source14_.is_DC1:
                d_1120___mcc_h8_ = source14_.cf1
                d_1121_cf1_ = d_1120___mcc_h8_
                (globalState).f19 = 752
                d_1067_v29_ = d_1067_v29_
                d_1122_v61_: D8
                d_1122_v61_ = D8_DC20(len(d_1074_v34_), d_1121_cf1_)
                d_1123_v62_: _dafny.MultiSet
                d_1123_v62_ = _dafny.MultiSet([d_1122_v61_])
                (globalState).f1 = (_dafny.MultiSet([d_1122_v61_])) == (d_1123_v62_)
                d_1124_v63_: C2
                nw195_ = C2()
                nw195_.ctor__(d_1116_v58_, d_1036_v1_, (self).f28)
                d_1124_v63_ = nw195_
            elif source14_.is_DC2:
                d_1125___mcc_h9_ = source14_.cf2
                d_1126___mcc_h10_ = source14_.cf3
                d_1127_cf3_ = d_1126___mcc_h10_
                d_1128_cf2_ = d_1125___mcc_h9_
                index201_ = default__.safeIndex(334, (d_1067_v29_).length(0))
                (d_1067_v29_)[index201_] = d_1035_v0_
                index202_ = default__.safeIndex(334, (d_1067_v29_).length(0))
                (d_1067_v29_)[index202_] = d_1035_v0_
                (globalState).f11 = ((40) - (328)) > (p0)
                rhs205_ = (d_1074_v34_) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjrs"))).set(default__.safeIndex((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))], len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjrs")))), d_1073_v33_))
                rhs206_ = not(not((d_1074_v34_) < (d_1074_v34_)))
                rhs207_ = d_1127_cf3_
                rhs208_ = p0
                lhs175_ = globalState
                lhs176_ = globalState
                lhs177_ = globalState
                d_1074_v34_ = rhs205_
                lhs175_.f27 = rhs206_
                lhs176_.f19 = rhs207_
                lhs177_.f14 = rhs208_
                d_1129_v64_: C0
                nw196_ = C0()
                nw196_.ctor__(((d_1117_v59_).f40) | ((d_1117_v59_).f40))
                d_1129_v64_ = nw196_
            elif True:
                d_1130___mcc_h11_ = source14_.cf0
                d_1131_cf0_ = d_1130___mcc_h11_
                d_1073_v33_ = d_1073_v33_
                (globalState).f27 = d_1035_v0_
                d_1132_v65_: _dafny.MultiSet
                d_1132_v65_ = _dafny.MultiSet([p0])
                d_1133_v66_: _dafny.MultiSet
                d_1133_v66_ = _dafny.MultiSet([(d_1132_v65_).cardinality])
                d_1134_v67_: _dafny.Seq
                d_1134_v67_ = _dafny.SeqWithoutIsStrInference([-988, -770])
                (globalState).f0 = default__.safeModuloInt((p0 if d_1035_v0_ else default__.fm2(d_1074_v34_, (d_1117_v59_).f40, (default__.fm18(d_1035_v0_, d_1035_v0_, globalState)).set(d_1035_v0_, d_1073_v33_), default__.fm4(globalState), globalState)), ((d_1133_v66_).set((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))], default__.abs((d_1117_v59_).fm17(d_1134_v67_, p0, globalState)))).cardinality)
                (globalState).f24 = default__.fm2(_dafny.SeqWithoutIsStrInference([d_1073_v33_ for d_1135_i7_ in range(default__.abs(870))]), ((d_1117_v59_).f40 if d_1035_v0_ else (d_1117_v59_).f40), d_1075_v35_, d_1035_v0_, globalState)
            d_1136_v68_: _dafny.Map
            d_1136_v68_ = _dafny.Map({True: d_1035_v0_})
            d_1137_v69_: _dafny.MultiSet
            d_1137_v69_ = _dafny.MultiSet([d_1136_v68_, (d_1136_v68_) | (d_1136_v68_)])
            d_1137_v69_ = d_1137_v69_
            index203_ = default__.safeIndex(776, (d_1065_v28_).length(0))
            (d_1065_v28_)[index203_] = (0) - (p0)
            d_1138_v70_: _dafny.Array
            nw197_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 16)
            d_1138_v70_ = nw197_
            index204_ = default__.safeIndex(806, (d_1138_v70_).length(0))
            (d_1138_v70_)[index204_] = d_1074_v34_
            index205_ = default__.safeIndex(806, (d_1138_v70_).length(0))
            (d_1138_v70_)[index205_] = d_1074_v34_
        elif True:
            index206_ = default__.safeIndex(308, (d_1067_v29_).length(0))
            (d_1067_v29_)[index206_] = d_1035_v0_
            d_1139_v71_: _dafny.Seq
            d_1139_v71_ = _dafny.SeqWithoutIsStrInference([p0])
            index207_ = default__.safeIndex(776, (d_1065_v28_).length(0))
            index208_ = default__.safeIndex(308, (d_1067_v29_).length(0))
            rhs209_ = (0) - (len(d_1139_v71_))
            rhs210_ = not(d_1035_v0_)
            lhs178_ = d_1065_v28_
            lhs179_ = default__.safeIndex(776, (d_1065_v28_).length(0))
            lhs180_ = d_1067_v29_
            lhs181_ = default__.safeIndex(308, (d_1067_v29_).length(0))
            lhs178_[lhs179_] = rhs209_
            lhs180_[lhs181_] = rhs210_
            d_1140_v72_: _dafny.Seq
            d_1140_v72_ = _dafny.SeqWithoutIsStrInference([not((d_1067_v29_)[default__.safeIndex(308, (d_1067_v29_).length(0))])])
            d_1141_v73_: _dafny.MultiSet
            d_1141_v73_ = _dafny.MultiSet([d_1140_v72_, _dafny.SeqWithoutIsStrInference([d_1035_v0_, d_1035_v0_])])
            rhs211_ = d_1035_v0_
            rhs212_ = (d_1141_v73_) | (d_1141_v73_)
            rhs213_ = (d_1140_v72_) != (d_1140_v72_)
            lhs182_ = globalState
            d_1035_v0_ = rhs211_
            d_1141_v73_ = rhs212_
            lhs182_.f27 = rhs213_
            index209_ = default__.safeIndex(776, (d_1065_v28_).length(0))
            (d_1065_v28_)[index209_] = (default__.safeDivisionInt((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))], (_dafny.MultiSet([(d_1067_v29_)[default__.safeIndex(308, (d_1067_v29_).length(0))]])).cardinality) if ((0) - (p0)) < ((d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]) else (p0 if d_1035_v0_ else 777))
            d_1142_v74_: _dafny.Array
            nw198_ = _dafny.Array(False, 23)
            d_1142_v74_ = nw198_
            default__.m0(d_1142_v74_, d_1065_v28_, (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))], (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))], globalState)
            (globalState).f0 = len(d_1074_v34_)
        (globalState).f24 = p0
        index210_ = default__.safeIndex(776, (d_1065_v28_).length(0))
        (d_1065_v28_)[index210_] = default__.safeDivisionInt(default__.safeModuloInt(p0, (d_1065_v28_)[default__.safeIndex(776, (d_1065_v28_).length(0))]), len(_dafny.SeqWithoutIsStrInference([d_1035_v0_])))
        r0 = d_1035_v0_
        return r0

    def m14(self, p0, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        r1: int = int(0)
        r2: _dafny.MultiSet = _dafny.MultiSet({})
        r3: int = int(0)
        d_1143_v0_: _dafny.Array
        def lambda41_(d_1144_i0_):
            return default__.safeDivisionInt(d_1144_i0_, -828)

        init23_ = lambda41_
        nw199_ = _dafny.Array(None, 5)
        for i0_23_ in range(nw199_.length(0)):
            nw199_[i0_23_] = init23_(i0_23_)
        d_1143_v0_ = nw199_
        d_1145_v1_: int
        d_1145_v1_ = 127
        d_1146_v2_: _dafny.Seq
        d_1146_v2_ = _dafny.SeqWithoutIsStrInference([d_1145_v1_, (0) - (d_1145_v1_)])
        nw200_ = _dafny.Array(None, 4)
        nw200_[int(0)] = d_1145_v1_
        nw200_[int(1)] = (d_1146_v2_)[default__.safeIndex(d_1145_v1_, len(d_1146_v2_))]
        nw200_[int(2)] = d_1145_v1_
        nw200_[int(3)] = default__.safeModuloInt(d_1145_v1_, d_1145_v1_)
        d_1143_v0_ = nw200_
        d_1147_v3_: bool
        d_1147_v3_ = False
        if d_1147_v3_:
            d_1148_v4_: C1
            nw201_ = C1()
            nw201_.ctor__(default__.safeModuloInt(d_1145_v1_, 662), d_1145_v1_, (self).f28)
            d_1148_v4_ = nw201_
            d_1149_v5_: _dafny.Seq
            d_1149_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hmjt"))
            d_1150_v6_: _dafny.Set
            d_1150_v6_ = _dafny.Set({d_1147_v3_, d_1147_v3_, d_1147_v3_})
            d_1151_v7_: _dafny.Seq
            d_1151_v7_ = _dafny.SeqWithoutIsStrInference([d_1150_v6_, d_1150_v6_])
            d_1152_v8_: _dafny.Map
            d_1152_v8_ = _dafny.Map({d_1145_v1_: d_1148_v4_.f42})
            d_1153_v9_: _dafny.Map
            d_1153_v9_ = _dafny.Map({not(d_1147_v3_): (self).fm8(d_1152_v8_, d_1147_v3_, d_1147_v3_, globalState)})
            d_1154_v10_: _dafny.Array
            nw202_ = _dafny.Array(_dafny.Array(None, 0), 5)
            d_1154_v10_ = nw202_
            d_1155_v11_: _dafny.Map
            d_1155_v11_ = _dafny.Map({d_1154_v10_: _dafny.MultiSet(d_1146_v2_)})
            d_1156_v12_: _dafny.MultiSet
            d_1156_v12_ = _dafny.MultiSet([d_1148_v4_.f42, 399, d_1148_v4_.f42])
            if not((default__.fm2(d_1149_v5_, (d_1151_v7_)[default__.safeIndex((d_1148_v4_).f41, len(d_1151_v7_))], d_1153_v9_, d_1147_v3_, globalState)) in (((d_1155_v11_)[d_1154_v10_] if (d_1154_v10_) in (d_1155_v11_) else d_1156_v12_))):
                d_1157_v13_: C0
                nw203_ = C0()
                nw203_.ctor__((_dafny.Set({d_1147_v3_, (d_1148_v4_).fm7(d_1147_v3_, d_1147_v3_, globalState), d_1147_v3_})).intersection(d_1150_v6_))
                d_1157_v13_ = nw203_
                d_1158_v14_: _dafny.Map
                d_1158_v14_ = _dafny.Map({d_1147_v3_: 449})
                d_1159_v15_: str
                d_1159_v15_ = _dafny.CodePoint('s')
                d_1160_v16_: _dafny.MultiSet
                d_1160_v16_ = _dafny.MultiSet([d_1159_v15_])
                d_1161_v17_: _dafny.Array
                def lambda42_(d_1162_i1_):
                    return True

                init24_ = lambda42_
                nw204_ = _dafny.Array(None, 28)
                for i0_24_ in range(nw204_.length(0)):
                    nw204_[i0_24_] = init24_(i0_24_)
                d_1161_v17_ = nw204_
                d_1163_v18_: _dafny.Map
                d_1163_v18_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(0) - (((d_1158_v14_)[d_1147_v3_] if (d_1147_v3_) in (d_1158_v14_) else d_1148_v4_.f42)), d_1145_v1_, ((d_1160_v16_).set(d_1159_v15_, default__.abs(53))).cardinality, len(_dafny.SeqWithoutIsStrInference([d_1145_v1_]))]): d_1161_v17_})
                d_1163_v18_ = (d_1163_v18_).set(_dafny.SeqWithoutIsStrInference([len(d_1146_v2_) for d_1164_i2_ in range(default__.abs(980))]), d_1161_v17_)
                d_1159_v15_ = d_1159_v15_
                d_1165_v19_: _dafny.Map
                d_1165_v19_ = _dafny.Map({d_1161_v17_: ((d_1148_v4_).f41) + (d_1148_v4_.f42)})
                d_1166_v20_: _dafny.Map
                d_1166_v20_ = _dafny.Map({d_1145_v1_: not(d_1147_v3_)})
                rhs214_ = ((d_1165_v19_)[d_1161_v17_] if (d_1161_v17_) in (d_1165_v19_) else d_1148_v4_.f42)
                rhs215_ = d_1166_v20_
                rhs216_ = (d_1146_v2_).set(default__.safeIndex((d_1148_v4_).f41, len(d_1146_v2_)), d_1148_v4_.f42)
                rhs217_ = default__.fm4(globalState)
                lhs183_ = globalState
                lhs184_ = globalState
                lhs185_ = globalState
                lhs183_.f0 = rhs214_
                lhs184_.f20 = rhs215_
                d_1146_v2_ = rhs216_
                lhs185_.f11 = rhs217_
                (globalState).f27 = d_1147_v3_
            elif True:
                d_1167_v21_: _dafny.Map
                d_1167_v21_ = _dafny.Map({803: d_1147_v3_})
                d_1168_v22_: _dafny.Map
                d_1168_v22_ = _dafny.Map({d_1147_v3_: (d_1148_v4_).f41})
                d_1169_v23_: C1
                nw205_ = C1()
                nw205_.ctor__(default__.safeDivisionInt(len(default__.fm3(globalState)), 618), default__.safeModuloInt(len(d_1167_v21_), (0) - (((d_1168_v22_)[d_1147_v3_] if (d_1147_v3_) in (d_1168_v22_) else d_1145_v1_))), D1_DC3(d_1147_v3_))
                d_1169_v23_ = nw205_
                d_1170_v24_: str
                d_1170_v24_ = _dafny.CodePoint('h')
                d_1170_v24_ = _dafny.CodePoint('y')
                d_1171_v25_: _dafny.Array
                nw206_ = _dafny.Array(None, 8)
                nw206_[int(0)] = d_1149_v5_
                nw206_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_1170_v24_ for d_1172_i3_ in range(default__.abs(498))])) + (d_1149_v5_)
                nw206_[int(2)] = (d_1149_v5_) + (d_1149_v5_)
                nw206_[int(3)] = ((d_1149_v5_).set(default__.safeIndex((0) - (d_1145_v1_), len(d_1149_v5_)), d_1170_v24_)) + (d_1149_v5_)
                nw206_[int(4)] = d_1149_v5_
                nw206_[int(5)] = (d_1149_v5_) + (d_1149_v5_)
                nw206_[int(6)] = d_1149_v5_
                nw206_[int(7)] = d_1149_v5_
                d_1171_v25_ = nw206_
                index211_ = default__.safeIndex(953, (d_1171_v25_).length(0))
                (d_1171_v25_)[index211_] = d_1149_v5_
                d_1173_v26_: _dafny.Map
                d_1173_v26_ = _dafny.Map({d_1148_v4_.f42: d_1149_v5_})
                index212_ = default__.safeIndex(953, (d_1171_v25_).length(0))
                (d_1171_v25_)[index212_] = (((d_1173_v26_)[len(d_1149_v5_)] if (len(d_1149_v5_)) in (d_1173_v26_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbvq")))) + (default__.fm3(globalState))
                d_1147_v3_ = not((d_1170_v24_) not in (d_1149_v5_))
                (globalState).f24 = d_1169_v23_.f42
            d_1174_v27_: str
            d_1174_v27_ = _dafny.CodePoint('d')
            d_1175_v28_: _dafny.Map
            d_1175_v28_ = _dafny.Map({d_1146_v2_: d_1174_v27_})
            d_1176_v29_: _dafny.Map
            d_1176_v29_ = _dafny.Map({d_1147_v3_: d_1175_v28_})
            d_1176_v29_ = (d_1176_v29_).set(not(d_1147_v3_), (d_1175_v28_) | (default__.fm33(False, globalState)))
            d_1177_v30_: _dafny.Map
            d_1177_v30_ = _dafny.Map({d_1174_v27_: (d_1156_v12_).isdisjoint(d_1156_v12_)})
            d_1177_v30_ = (d_1177_v30_).set(d_1174_v27_, not (d_1147_v3_) or (d_1147_v3_))
            d_1178_v31_: _dafny.Seq
            d_1178_v31_ = _dafny.SeqWithoutIsStrInference([d_1147_v3_, not(d_1147_v3_), d_1147_v3_, d_1147_v3_])
            (globalState).f16 = ((d_1178_v31_) + (d_1178_v31_)) + (d_1178_v31_)
        elif True:
            d_1179_v32_: _dafny.Set
            d_1179_v32_ = _dafny.Set({d_1143_v0_, d_1143_v0_, d_1143_v0_, d_1143_v0_, d_1143_v0_})
            d_1179_v32_ = d_1179_v32_
            d_1180_v33_: _dafny.Map
            d_1180_v33_ = _dafny.Map({d_1147_v3_: d_1145_v1_})
            d_1181_v34_: _dafny.Map
            d_1181_v34_ = _dafny.Map({len(d_1180_v33_): d_1147_v3_})
            index213_ = default__.safeIndex(320, (d_1143_v0_).length(0))
            (d_1143_v0_)[index213_] = len(d_1181_v34_)
            d_1182_v35_: _dafny.Seq
            d_1182_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xjka"))
            index214_ = default__.safeIndex(320, (d_1143_v0_).length(0))
            (d_1143_v0_)[index214_] = default__.safeDivisionInt((len(d_1182_v35_)) * (d_1145_v1_), d_1145_v1_)
            d_1183_v36_: _dafny.Set
            d_1183_v36_ = _dafny.Set({d_1147_v3_, d_1147_v3_})
            d_1184_v37_: _dafny.Array
            def lambda43_(d_1185_v3_):
                def lambda44_(d_1186_i4_):
                    return d_1185_v3_

                return lambda44_

            init25_ = lambda43_(d_1147_v3_)
            nw207_ = _dafny.Array(None, 22)
            for i0_25_ in range(nw207_.length(0)):
                nw207_[i0_25_] = init25_(i0_25_)
            d_1184_v37_ = nw207_
            d_1187_v38_: _dafny.Map
            d_1187_v38_ = _dafny.Map({(d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))]: d_1184_v37_})
            d_1188_v39_: _dafny.Map
            d_1188_v39_ = _dafny.Map({len(d_1187_v38_): (d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))]})
            d_1189_v40_: _dafny.Map
            d_1189_v40_ = _dafny.Map({d_1183_v36_: d_1188_v39_})
            d_1190_v41_: _dafny.Map
            d_1190_v41_ = d_1189_v40_
            d_1189_v40_ = (d_1189_v40_) | ((d_1190_v41_))
            d_1147_v3_ = (d_1182_v35_) < ((d_1182_v35_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1191_i5_ in range(default__.abs(-811))])))
            if not(not (d_1147_v3_) or (d_1147_v3_)):
                index215_ = default__.safeIndex(462, (d_1184_v37_).length(0))
                (d_1184_v37_)[index215_] = True
                d_1192_v42_: _dafny.Set
                d_1192_v42_ = _dafny.Set({d_1184_v37_, d_1184_v37_, d_1184_v37_})
                index216_ = default__.safeIndex(462, (d_1184_v37_).length(0))
                rhs218_ = d_1147_v3_
                rhs219_ = (0) - (len(d_1192_v42_))
                rhs220_ = d_1182_v35_
                lhs186_ = d_1184_v37_
                lhs187_ = default__.safeIndex(462, (d_1184_v37_).length(0))
                lhs188_ = globalState
                lhs186_[lhs187_] = rhs218_
                lhs188_.f24 = rhs219_
                d_1182_v35_ = rhs220_
                d_1193_v43_: _dafny.Array
                nw208_ = _dafny.Array(False, 24)
                d_1193_v43_ = nw208_
                d_1194_v44_: _dafny.MultiSet
                d_1194_v44_ = _dafny.MultiSet([d_1193_v43_, d_1193_v43_, (d_1193_v43_ if d_1147_v3_ else d_1193_v43_), d_1193_v43_])
                index217_ = default__.safeIndex(320, (d_1143_v0_).length(0))
                index218_ = default__.safeIndex(320, (d_1143_v0_).length(0))
                rhs221_ = ((d_1194_v44_)[d_1193_v43_] if (d_1193_v43_) in (d_1194_v44_) else d_1145_v1_)
                rhs222_ = (0) - ((0) - (((d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))] if (d_1184_v37_)[default__.safeIndex(462, (d_1184_v37_).length(0))] else (d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))])))
                rhs223_ = d_1147_v3_
                lhs189_ = d_1143_v0_
                lhs190_ = default__.safeIndex(320, (d_1143_v0_).length(0))
                lhs191_ = d_1143_v0_
                lhs192_ = default__.safeIndex(320, (d_1143_v0_).length(0))
                lhs193_ = globalState
                lhs189_[lhs190_] = rhs221_
                lhs191_[lhs192_] = rhs222_
                lhs193_.f1 = rhs223_
                (globalState).f27 = d_1147_v3_
                d_1195_v45_: _dafny.Array
                nw209_ = _dafny.Array(_dafny.Array(None, 0), 15)
                d_1195_v45_ = nw209_
                index219_ = default__.safeIndex(392, (d_1195_v45_).length(0))
                (d_1195_v45_)[index219_] = d_1193_v43_
                d_1196_v46_: _dafny.Array
                nw210_ = _dafny.Array(_dafny.Seq({}), 19)
                d_1196_v46_ = nw210_
                d_1197_v47_: _dafny.Seq
                d_1197_v47_ = _dafny.SeqWithoutIsStrInference([d_1196_v46_])
                d_1198_v48_: _dafny.Array
                nw211_ = _dafny.Array(None, 27)
                nw211_[int(0)] = d_1196_v46_
                nw211_[int(1)] = d_1196_v46_
                nw211_[int(2)] = d_1196_v46_
                nw211_[int(3)] = d_1196_v46_
                nw211_[int(4)] = d_1196_v46_
                nw211_[int(5)] = d_1196_v46_
                nw211_[int(6)] = d_1196_v46_
                nw211_[int(7)] = d_1196_v46_
                nw211_[int(8)] = d_1196_v46_
                nw211_[int(9)] = (d_1197_v47_)[default__.safeIndex((d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))], len(d_1197_v47_))]
                nw211_[int(10)] = d_1196_v46_
                nw211_[int(11)] = d_1196_v46_
                nw211_[int(12)] = d_1196_v46_
                nw211_[int(13)] = d_1196_v46_
                nw211_[int(14)] = d_1196_v46_
                nw211_[int(15)] = d_1196_v46_
                nw211_[int(16)] = d_1196_v46_
                nw211_[int(17)] = d_1196_v46_
                nw211_[int(18)] = d_1196_v46_
                nw211_[int(19)] = d_1196_v46_
                nw211_[int(20)] = d_1196_v46_
                nw211_[int(21)] = d_1196_v46_
                nw211_[int(22)] = d_1196_v46_
                nw211_[int(23)] = d_1196_v46_
                nw211_[int(24)] = d_1196_v46_
                nw211_[int(25)] = d_1196_v46_
                nw211_[int(26)] = d_1196_v46_
                d_1198_v48_ = nw211_
                index220_ = default__.safeIndex(431, (d_1198_v48_).length(0))
                (d_1198_v48_)[index220_] = d_1196_v46_
                index221_ = default__.safeIndex(392, (d_1195_v45_).length(0))
                index222_ = default__.safeIndex(431, (d_1198_v48_).length(0))
                rhs224_ = not(((d_1184_v37_)[default__.safeIndex(462, (d_1184_v37_).length(0))]) and (d_1147_v3_))
                rhs225_ = (d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))]
                rhs226_ = d_1193_v43_
                rhs227_ = d_1196_v46_
                rhs228_ = (0) - (default__.safeDivisionInt(d_1145_v1_, -683))
                lhs194_ = globalState
                lhs195_ = globalState
                lhs196_ = d_1195_v45_
                lhs197_ = default__.safeIndex(392, (d_1195_v45_).length(0))
                lhs198_ = d_1198_v48_
                lhs199_ = default__.safeIndex(431, (d_1198_v48_).length(0))
                lhs200_ = globalState
                lhs194_.f11 = rhs224_
                lhs195_.f19 = rhs225_
                lhs196_[lhs197_] = rhs226_
                lhs198_[lhs199_] = rhs227_
                lhs200_.f0 = rhs228_
                (globalState).f24 = d_1145_v1_
            elif True:
                d_1199_v50_: _dafny.Seq
                d_1199_v50_ = _dafny.SeqWithoutIsStrInference([d_1147_v3_, d_1147_v3_])
                d_1200_v52_: C2
                nw212_ = C2()
                def iife132_():
                    def iife134_():
                        coll56_ = _dafny.Set()
                        compr_56_: _dafny.Seq
                        for compr_56_ in (_dafny.Set({d_1199_v50_})).Elements:
                            d_1203_v51_: _dafny.Seq = compr_56_
                            if (d_1203_v51_) in (_dafny.Set({d_1199_v50_})):
                                coll56_ = coll56_.union(_dafny.Set([d_1203_v51_]))
                        return _dafny.Set(coll56_)
                    coll54_ = _dafny.Map()
                    def iife133_():
                        coll55_ = _dafny.Set()
                        compr_55_: _dafny.Seq
                        for compr_55_ in (_dafny.Set({d_1199_v50_})).Elements:
                            d_1201_v51_: _dafny.Seq = compr_55_
                            if (d_1201_v51_) in (_dafny.Set({d_1199_v50_})):
                                coll55_ = coll55_.union(_dafny.Set([d_1201_v51_]))
                        return _dafny.Set(coll55_)
                    compr_54_: _dafny.Seq
                    for compr_54_ in (iife133_()
                    ).Elements:
                        d_1202_v49_: _dafny.Seq = compr_54_
                        if (d_1202_v49_) in (iife134_()
                        ):
                            coll54_[d_1202_v49_] = len(_dafny.Map({(d_1143_v0_)[default__.safeIndex(320, (d_1143_v0_).length(0))]: _dafny.CodePoint('q')}))
                    return _dafny.Map(coll54_)
                nw212_.ctor__(iife132_()
                , (d_1183_v36_) | (_dafny.Set({d_1147_v3_})), D1_DC3(d_1147_v3_))
                d_1200_v52_ = nw212_
                d_1184_v37_ = d_1184_v37_
                d_1204_v53_: str
                d_1204_v53_ = _dafny.CodePoint('q')
                d_1205_v54_: D2
                d_1205_v54_ = D2_DC7(d_1147_v3_, (default__.fm4(globalState) if d_1147_v3_ else d_1147_v3_), (d_1182_v35_) + (_dafny.SeqWithoutIsStrInference([d_1204_v53_, d_1204_v53_])))
                d_1205_v54_ = d_1205_v54_
                rhs229_ = not(d_1147_v3_)
                rhs230_ = d_1184_v37_
                lhs201_ = globalState
                lhs201_.f11 = rhs229_
                d_1184_v37_ = rhs230_
                d_1206_v55_: _dafny.MultiSet
                d_1206_v55_ = _dafny.MultiSet([d_1199_v50_])
                d_1147_v3_ = (d_1206_v55_).ispropersubset(d_1206_v55_)
        (globalState).f0 = d_1145_v1_
        d_1207_v56_: _dafny.Array
        nw213_ = _dafny.Array(False, 8)
        d_1207_v56_ = nw213_
        d_1208_v57_: _dafny.Seq
        d_1208_v57_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
        d_1209_v58_: str
        d_1209_v58_ = _dafny.CodePoint('h')
        default__.m0(d_1207_v56_, d_1143_v0_, d_1145_v1_, (d_1145_v1_ if d_1147_v3_ else (0) - (len((d_1208_v57_).set(default__.safeIndex((0) - (d_1145_v1_), len(d_1208_v57_)), d_1209_v58_)))), globalState)
        d_1210_v59_: _dafny.Seq
        d_1210_v59_ = _dafny.SeqWithoutIsStrInference([d_1147_v3_])
        d_1211_v60_: _dafny.Map
        d_1211_v60_ = _dafny.Map({d_1210_v59_: -793})
        d_1212_v61_: C2
        nw214_ = C2()
        nw214_.ctor__(d_1211_v60_, (_dafny.Set({d_1147_v3_, d_1147_v3_, d_1147_v3_, d_1147_v3_}) if d_1147_v3_ else _dafny.Set({d_1147_v3_})), (self).f28)
        d_1212_v61_ = nw214_
        d_1212_v61_ = d_1212_v61_
        r3 = d_1145_v1_
        d_1213_v62_: _dafny.Seq
        d_1213_v62_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smtljvbj"))).set(default__.safeIndex(d_1145_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "smtljvbj")))), d_1209_v58_), d_1208_v57_])
        d_1214_v63_: _dafny.MultiSet
        d_1214_v63_ = _dafny.MultiSet([d_1208_v57_, d_1208_v57_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ptmphaiv"))])
        d_1215_v64_: _dafny.Map
        d_1215_v64_ = _dafny.Map({957: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1208_v57_, d_1208_v57_]))})
        r0 = ((_dafny.MultiSet(d_1213_v62_)) - (d_1214_v63_)).intersection(((d_1215_v64_)[d_1145_v1_] if (d_1145_v1_) in (d_1215_v64_) else d_1214_v63_))
        d_1216_v65_: _dafny.Map
        d_1216_v65_ = _dafny.Map({d_1147_v3_: d_1209_v58_})
        d_1217_v66_: _dafny.Map
        d_1217_v66_ = _dafny.Map({not(d_1147_v3_): d_1147_v3_})
        r1 = (d_1145_v1_) - (default__.safeDivisionInt(default__.fm2(d_1208_v57_, (d_1212_v61_).f40, d_1216_v65_, ((d_1217_v66_)[d_1147_v3_] if (d_1147_v3_) in (d_1217_v66_) else ((d_1217_v66_)[d_1147_v3_] if (d_1147_v3_) in (d_1217_v66_) else True)), globalState), d_1145_v1_))
        d_1218_v67_: _dafny.MultiSet
        d_1218_v67_ = _dafny.MultiSet([639, (d_1145_v1_) * (d_1145_v1_), d_1145_v1_])
        r2 = d_1218_v67_
        r3 = (d_1218_v67_).cardinality
        return r0, r1, r2, r3


class C4:
    def  __init__(self):
        self.f38: _dafny.Array = _dafny.Array(None, 0)
        self._f37: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    def ctor__(self, f37, f38):
        (self)._f37 = f37
        (self).f38 = f38

    def m12(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1219_v0_: _dafny.Seq
        d_1219_v0_ = _dafny.SeqWithoutIsStrInference([p3, p3])
        d_1220_v1_: int
        d_1220_v1_ = -108
        d_1221_v2_: _dafny.Seq
        d_1221_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gwryo"))
        d_1222_v3_: _dafny.Seq
        d_1222_v3_ = _dafny.SeqWithoutIsStrInference([940, len(d_1221_v2_)])
        d_1223_v4_: _dafny.Map
        d_1223_v4_ = _dafny.Map({d_1220_v1_: p1})
        d_1224_v5_: D2
        d_1224_v5_ = D2_DC8(((d_1223_v4_)[113] if (113) in (d_1223_v4_) else (self).f37), p3)
        d_1225_v6_: _dafny.Map
        d_1225_v6_ = _dafny.Map({d_1222_v3_: (d_1224_v5_).cf17})
        d_1226_v7_: _dafny.Array
        nw215_ = _dafny.Array(None, 23)
        nw215_[int(0)] = p3
        nw215_[int(1)] = p3
        nw215_[int(2)] = p3
        nw215_[int(3)] = p3
        nw215_[int(4)] = p3
        nw215_[int(5)] = p3
        nw215_[int(6)] = p3
        nw215_[int(7)] = p3
        nw215_[int(8)] = p3
        nw215_[int(9)] = (d_1219_v0_)[default__.safeIndex(d_1220_v1_, len(d_1219_v0_))]
        nw215_[int(10)] = p3
        nw215_[int(11)] = p3
        nw215_[int(12)] = p3
        nw215_[int(13)] = p3
        nw215_[int(14)] = ((d_1225_v6_)[d_1222_v3_] if (d_1222_v3_) in (d_1225_v6_) else p3)
        nw215_[int(15)] = p3
        nw215_[int(16)] = p3
        nw215_[int(17)] = p3
        nw215_[int(18)] = p3
        nw215_[int(19)] = p3
        nw215_[int(20)] = p3
        nw215_[int(21)] = p3
        nw215_[int(22)] = p3
        d_1226_v7_ = nw215_
        index223_ = default__.safeIndex(582, (d_1226_v7_).length(0))
        (d_1226_v7_)[index223_] = p3
        index224_ = default__.safeIndex(582, (d_1226_v7_).length(0))
        (d_1226_v7_)[index224_] = p3
        d_1227_v8_: _dafny.Set
        d_1227_v8_ = _dafny.Set({(self).f37, (self).f37, p1})
        d_1228_v9_: C0
        nw216_ = C0()
        nw216_.ctor__(d_1227_v8_)
        d_1228_v9_ = nw216_
        d_1229_v10_: _dafny.Map
        d_1229_v10_ = _dafny.Map({d_1221_v2_: True})
        d_1230_v12_: _dafny.Seq
        d_1230_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p0 for d_1231_i0_ in range(default__.abs(165))]), _dafny.SeqWithoutIsStrInference([p0 for d_1232_i1_ in range(default__.abs(144))]), d_1221_v2_])
        d_1233_v13_: _dafny.Map
        def iife135_():
            coll57_ = _dafny.Map()
            compr_57_: _dafny.Seq
            for compr_57_ in (d_1230_v12_).Elements:
                d_1234_v11_: _dafny.Seq = compr_57_
                if (d_1234_v11_) in (d_1230_v12_):
                    coll57_[d_1234_v11_] = ((d_1223_v4_)[d_1220_v1_] if (d_1220_v1_) in (d_1223_v4_) else True)
            return _dafny.Map(coll57_)
        d_1233_v13_ = _dafny.Map({(d_1229_v10_) | (iife135_()
        ): (_dafny.MultiSet([d_1220_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nrjgh")))])).cardinality})
        d_1233_v13_ = (d_1233_v13_).set((d_1229_v10_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xhomqv")), p1), d_1220_v1_)
        d_1235_v14_: _dafny.Set
        d_1235_v14_ = _dafny.Set({696})
        d_1236_v15_: _dafny.Seq
        d_1236_v15_ = _dafny.SeqWithoutIsStrInference([d_1235_v14_])
        d_1237_v16_: _dafny.Seq
        d_1237_v16_ = d_1236_v15_
        pat_let_tv32_ = d_1220_v1_
        def lambda45_(source15_):
            d_1238___mcc_h0_ = source15_
            d_1239_cf24_ = d_1238___mcc_h0_
            return pat_let_tv32_

        d_1220_v1_ = lambda45_(d_1237_v16_)
        d_1240_v17_: _dafny.Set
        d_1240_v17_ = _dafny.Set({p0, _dafny.CodePoint('d')})
        d_1241_v18_: D4
        d_1241_v18_ = D4_DC12(False, d_1240_v17_)
        source16_ = D4_DC13(d_1241_v18_)
        if source16_.is_DC11:
            d_1242___mcc_h1_ = source16_.cf20
            d_1243_cf20_ = d_1242___mcc_h1_
            d_1244_v19_: _dafny.Array
            nw217_ = _dafny.Array(None, 16)
            nw217_[int(0)] = not (p1) or ((self).f37)
            nw217_[int(1)] = (self).f37
            nw217_[int(2)] = (d_1221_v2_) == (d_1221_v2_)
            nw217_[int(3)] = (self).f37
            nw217_[int(4)] = not(True)
            nw217_[int(5)] = not(p1)
            nw217_[int(6)] = (self).f37
            nw217_[int(7)] = p1
            nw217_[int(8)] = not((len(d_1222_v3_)) <= (d_1220_v1_))
            nw217_[int(9)] = (self).f37
            nw217_[int(10)] = default__.fm4(globalState)
            nw217_[int(11)] = p1
            nw217_[int(12)] = (self).f37
            nw217_[int(13)] = p1
            nw217_[int(14)] = (self).f37
            nw217_[int(15)] = (self).f37
            d_1244_v19_ = nw217_
            index225_ = default__.safeIndex(731, (d_1244_v19_).length(0))
            (d_1244_v19_)[index225_] = not (not((self).f37)) or ((self).f37)
            d_1245_v20_: _dafny.MultiSet
            d_1245_v20_ = _dafny.MultiSet([p1, (self).f37])
            d_1246_v21_: _dafny.Seq
            d_1246_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1}), (d_1228_v9_).f43, d_1227_v8_])
            index226_ = default__.safeIndex(731, (d_1244_v19_).length(0))
            (d_1244_v19_)[index226_] = not((_dafny.Set({(0) - (d_1220_v1_), d_1220_v1_, len(_dafny.Set({(d_1245_v20_).cardinality, d_1220_v1_, d_1220_v1_})), default__.fm2(d_1221_v2_, (d_1246_v21_)[default__.safeIndex(d_1220_v1_, len(d_1246_v21_))], _dafny.Map({(self).f37: p0}), p1, globalState)})).isdisjoint(d_1235_v14_))
            d_1247_v22_: _dafny.Array
            def lambda46_(d_1248_v3_):
                def lambda47_(d_1249_i2_):
                    return d_1248_v3_

                return lambda47_

            init26_ = lambda46_(d_1222_v3_)
            nw218_ = _dafny.Array(None, 18)
            for i0_26_ in range(nw218_.length(0)):
                nw218_[i0_26_] = init26_(i0_26_)
            d_1247_v22_ = nw218_
            d_1250_v23_: _dafny.Map
            d_1250_v23_ = _dafny.Map({not((d_1244_v19_)[default__.safeIndex(731, (d_1244_v19_).length(0))]): p0})
            arr0_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
            index227_ = default__.safeIndex(311, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
            arr0_[index227_] = default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fcvvufh")), d_1227_v8_, d_1250_v23_, (d_1244_v19_)[default__.safeIndex(731, (d_1244_v19_).length(0))], globalState)
            d_1251_v24_: _dafny.Array
            d_1251_v24_ = d_1247_v22_
            arr1_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
            index228_ = default__.safeIndex(311, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
            rhs231_ = (d_1221_v2_) + (d_1221_v2_)
            rhs232_ = (d_1251_v24_)
            rhs233_ = d_1244_v19_
            rhs234_ = (0) - ((0) - (d_1220_v1_))
            lhs202_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
            lhs203_ = default__.safeIndex(311, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
            d_1221_v2_ = rhs231_
            d_1247_v22_ = rhs232_
            d_1244_v19_ = rhs233_
            lhs202_[lhs203_] = rhs234_
            d_1252_v25_: _dafny.Array
            nw219_ = _dafny.Array(_dafny.Array(None, 0), 15)
            d_1252_v25_ = nw219_
            index229_ = default__.safeIndex(248, (d_1252_v25_).length(0))
            (d_1252_v25_)[index229_] = d_1244_v19_
            d_1253_v26_: D6
            d_1253_v26_ = D6_DC16(p1, (d_1244_v19_)[default__.safeIndex(731, (d_1244_v19_).length(0))])
            index230_ = default__.safeIndex(248, (d_1252_v25_).length(0))
            nw220_ = _dafny.Array(None, 13)
            nw220_[int(0)] = p1
            nw220_[int(1)] = (self).f37
            nw220_[int(2)] = (d_1244_v19_)[default__.safeIndex(731, (d_1244_v19_).length(0))]
            nw220_[int(3)] = p1
            nw220_[int(4)] = (d_1220_v1_) == (d_1220_v1_)
            nw220_[int(5)] = (d_1244_v19_)[default__.safeIndex(731, (d_1244_v19_).length(0))]
            nw220_[int(6)] = not((d_1253_v26_).cf26)
            nw220_[int(7)] = (self).f37
            nw220_[int(8)] = p1
            nw220_[int(9)] = (((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(311, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))]) == (d_1220_v1_)
            nw220_[int(10)] = (d_1224_v5_).cf16
            nw220_[int(11)] = p1
            nw220_[int(12)] = True
            (d_1252_v25_)[index230_] = nw220_
            d_1254_v27_: C0
            nw221_ = C0()
            nw221_.ctor__(_dafny.Set({not((self).f37), (d_1244_v19_)[default__.safeIndex(731, (d_1244_v19_).length(0))]}))
            d_1254_v27_ = nw221_
        elif source16_.is_DC12:
            d_1255___mcc_h2_ = source16_.cf21
            d_1256___mcc_h3_ = source16_.cf22
            d_1257_cf22_ = d_1256___mcc_h3_
            d_1258_cf21_ = d_1255___mcc_h2_
            d_1221_v2_ = (d_1221_v2_) + (d_1221_v2_)
            d_1259_v28_: D1
            d_1259_v28_ = D1_DC3(p1)
            d_1260_v29_: C3
            nw222_ = C3()
            nw222_.ctor__(d_1259_v28_)
            d_1260_v29_ = nw222_
            d_1261_v30_: _dafny.Array
            def lambda48_(d_1262_i3_):
                return (self).f37

            init27_ = lambda48_
            nw223_ = _dafny.Array(None, 2)
            for i0_27_ in range(nw223_.length(0)):
                nw223_[i0_27_] = init27_(i0_27_)
            d_1261_v30_ = nw223_
            index231_ = default__.safeIndex(877, (d_1261_v30_).length(0))
            (d_1261_v30_)[index231_] = not (False) or (p1)
            index232_ = default__.safeIndex(877, (d_1261_v30_).length(0))
            def iife136_():
                coll58_ = _dafny.Set()
                compr_58_: int
                for compr_58_ in _dafny.IntegerRange(-687, -868):
                    d_1263_v31_: int = compr_58_
                    if ((-687) <= (d_1263_v31_)) and ((d_1263_v31_) < (-868)):
                        coll58_ = coll58_.union(_dafny.Set([(d_1263_v31_) + (d_1220_v1_)]))
                return _dafny.Set(coll58_)
            (d_1261_v30_)[index232_] = not((default__.safeModuloInt(len(iife136_()
            ), 845)) == ((len(_dafny.Map({(self).f37: 160}))) - (d_1220_v1_)))
            if (d_1261_v30_)[default__.safeIndex(877, (d_1261_v30_).length(0))]:
                d_1264_v32_: C1
                nw224_ = C1()
                nw224_.ctor__(d_1220_v1_, d_1220_v1_, d_1259_v28_)
                d_1264_v32_ = nw224_
                d_1265_v33_: _dafny.Seq
                d_1265_v33_ = _dafny.SeqWithoutIsStrInference([d_1258_cf21_, not((d_1261_v30_)[default__.safeIndex(877, (d_1261_v30_).length(0))]), (self).f37])
                d_1266_v34_: _dafny.Seq
                d_1266_v34_ = _dafny.SeqWithoutIsStrInference([d_1265_v33_])
                d_1267_v35_: D2
                d_1267_v35_ = D2_DC6(d_1261_v30_)
                d_1268_v36_: _dafny.Set
                d_1268_v36_ = _dafny.Set({d_1267_v35_})
                d_1269_v37_: _dafny.Map
                d_1269_v37_ = _dafny.Map({(d_1266_v34_)[default__.safeIndex(d_1220_v1_, len(d_1266_v34_))]: (_dafny.Set({d_1267_v35_, d_1267_v35_, d_1267_v35_})).ispropersubset(d_1268_v36_)})
                d_1269_v37_ = d_1269_v37_
                d_1270_v38_: _dafny.Map
                d_1270_v38_ = _dafny.Map({(d_1264_v32_).f41: d_1261_v30_})
                d_1271_v39_: _dafny.Seq
                d_1271_v39_ = _dafny.SeqWithoutIsStrInference([d_1261_v30_])
                index233_ = default__.safeIndex(877, (d_1261_v30_).length(0))
                rhs235_ = d_1235_v14_
                rhs236_ = (d_1264_v32_).fm7(p1, (len((D15_DC33((d_1270_v38_).set(896, (d_1271_v39_)[default__.safeIndex(d_1220_v1_, len(d_1271_v39_))]))).cf58)) > ((d_1264_v32_).f41), globalState)
                lhs204_ = d_1261_v30_
                lhs205_ = default__.safeIndex(877, (d_1261_v30_).length(0))
                d_1235_v14_ = rhs235_
                lhs204_[lhs205_] = rhs236_
                d_1272_v40_: _dafny.Map
                d_1272_v40_ = _dafny.Map({(D2_DC6(d_1261_v30_)).cf12: d_1264_v32_.f42})
                d_1220_v1_ = ((d_1272_v40_)[d_1261_v30_] if (d_1261_v30_) in (d_1272_v40_) else (d_1264_v32_).f41)
                arr2_ = self.f38
                index234_ = default__.safeIndex(657, (self.f38).length(0))
                arr2_[index234_] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "loet"))).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([(self).f37, d_1258_cf21_])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "loet")))), p0)
                arr3_ = self.f38
                index235_ = default__.safeIndex(657, (self.f38).length(0))
                arr3_[index235_] = default__.fm3(globalState)
            elif True:
                d_1273_v41_: _dafny.Map
                d_1273_v41_ = _dafny.Map({d_1258_cf21_: True})
                d_1273_v41_ = (d_1273_v41_).set((self).f37, not((p1) and (d_1258_cf21_)))
                (globalState).f14 = ((0) - (d_1220_v1_)) - (d_1220_v1_)
                (globalState).f24 = d_1220_v1_
                d_1274_v42_: _dafny.Seq
                d_1274_v42_ = _dafny.SeqWithoutIsStrInference([d_1261_v30_, d_1261_v30_])
                d_1275_v43_: _dafny.MultiSet
                d_1275_v43_ = _dafny.MultiSet([(d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))], p3])
                d_1276_v44_: _dafny.Array
                nw225_ = _dafny.Array(None, 17)
                nw225_[int(0)] = d_1261_v30_
                nw225_[int(1)] = d_1261_v30_
                nw225_[int(2)] = d_1261_v30_
                nw225_[int(3)] = d_1261_v30_
                nw225_[int(4)] = d_1261_v30_
                nw225_[int(5)] = d_1261_v30_
                nw225_[int(6)] = d_1261_v30_
                nw225_[int(7)] = d_1261_v30_
                nw225_[int(8)] = d_1261_v30_
                nw225_[int(9)] = d_1261_v30_
                nw225_[int(10)] = d_1261_v30_
                nw225_[int(11)] = d_1261_v30_
                nw225_[int(12)] = d_1261_v30_
                nw225_[int(13)] = (d_1274_v42_)[default__.safeIndex((0) - ((d_1275_v43_).cardinality), len(d_1274_v42_))]
                nw225_[int(14)] = d_1261_v30_
                nw225_[int(15)] = d_1261_v30_
                nw225_[int(16)] = d_1261_v30_
                d_1276_v44_ = nw225_
                d_1276_v44_ = d_1276_v44_
                d_1277_v45_: _dafny.Array
                def lambda49_(d_1278_p0_):
                    def lambda50_(d_1279_i4_):
                        return d_1278_p0_

                    return lambda50_

                init28_ = lambda49_(p0)
                nw226_ = _dafny.Array(None, 5)
                for i0_28_ in range(nw226_.length(0)):
                    nw226_[i0_28_] = init28_(i0_28_)
                d_1277_v45_ = nw226_
                index236_ = default__.safeIndex(107, (d_1277_v45_).length(0))
                (d_1277_v45_)[index236_] = p0
                index237_ = default__.safeIndex(107, (d_1277_v45_).length(0))
                (d_1277_v45_)[index237_] = p0
        elif source16_.is_DC10:
            d_1280___mcc_h4_ = source16_.cf19
            d_1281_cf19_ = d_1280___mcc_h4_
            d_1282_v46_: _dafny.MultiSet
            d_1282_v46_ = _dafny.MultiSet([d_1220_v1_])
            d_1283_v47_: _dafny.Map
            d_1283_v47_ = _dafny.Map({d_1282_v46_: d_1281_cf19_})
            d_1284_v48_: _dafny.Map
            d_1284_v48_ = _dafny.Map({len(d_1283_v47_): d_1223_v4_})
            d_1285_v49_: _dafny.Map
            d_1285_v49_ = _dafny.Map({d_1220_v1_: d_1220_v1_})
            d_1286_v50_: _dafny.Map
            d_1286_v50_ = _dafny.Map({d_1284_v48_: d_1285_v49_})
            d_1287_v51_: _dafny.Set
            d_1287_v51_ = _dafny.Set({((d_1286_v50_)[_dafny.Map({d_1220_v1_: d_1223_v4_})] if (_dafny.Map({d_1220_v1_: d_1223_v4_})) in (d_1286_v50_) else _dafny.Map({len(default__.fm34(globalState)): 793})), d_1285_v49_})
            rhs237_ = not(default__.fm4(globalState))
            rhs238_ = len(d_1287_v51_)
            rhs239_ = (self).f37
            lhs206_ = globalState
            lhs207_ = globalState
            r0 = rhs237_
            lhs206_.f0 = rhs238_
            lhs207_.f1 = rhs239_
            d_1288_v52_: C1
            nw227_ = C1()
            nw227_.ctor__(844, default__.safeModuloInt(d_1220_v1_, ((d_1282_v46_)[d_1220_v1_] if (d_1220_v1_) in (d_1282_v46_) else d_1220_v1_)), D1_DC3((self).f37))
            d_1288_v52_ = nw227_
            d_1221_v2_ = (d_1221_v2_).set(default__.safeIndex(d_1288_v52_.f42, len(d_1221_v2_)), p0)
            d_1281_cf19_ = default__.fm5((d_1220_v1_) * (d_1220_v1_), (self).f37, globalState)
        elif True:
            d_1289___mcc_h5_ = source16_.cf23
            d_1290_cf23_ = d_1289___mcc_h5_
            d_1291_v53_: D4
            d_1291_v53_ = D4_DC10(p0)
            d_1292_v54_: _dafny.Map
            d_1292_v54_ = _dafny.Map({d_1291_v53_: d_1220_v1_})
            d_1293_v55_: D7
            d_1293_v55_ = D7_DC17(d_1292_v54_)
            source17_ = d_1293_v55_
            if source17_.is_DC18:
                d_1294___mcc_h6_ = source17_.cf29
                d_1295_cf29_ = d_1294___mcc_h6_
                d_1296_v56_: _dafny.Array
                nw228_ = _dafny.Array(False, 16)
                d_1296_v56_ = nw228_
                d_1296_v56_ = d_1296_v56_
                d_1297_v57_: _dafny.Map
                d_1297_v57_ = _dafny.Map({(p1) == ((self).f37): d_1220_v1_})
                d_1297_v57_ = d_1297_v57_
                d_1223_v4_ = (d_1223_v4_).set((0) - (default__.safeModuloInt(len(d_1235_v14_), d_1220_v1_)), (default__.fm4(globalState)) and (not(True)))
                index238_ = default__.safeIndex(467, (d_1296_v56_).length(0))
                (d_1296_v56_)[index238_] = p1
                d_1298_v58_: _dafny.MultiSet
                d_1298_v58_ = _dafny.MultiSet([_dafny.Map({65: d_1220_v1_})])
                index239_ = default__.safeIndex(467, (d_1296_v56_).length(0))
                (d_1296_v56_)[index239_] = (d_1298_v58_).issubset(d_1298_v58_)
            elif True:
                d_1299___mcc_h7_ = source17_.cf28
                d_1300_cf28_ = d_1299___mcc_h7_
                d_1301_v59_: _dafny.Array
                nw229_ = _dafny.Array(None, 23)
                nw229_[int(0)] = p0
                nw229_[int(1)] = p0
                nw229_[int(2)] = p0
                nw229_[int(3)] = _dafny.CodePoint('p')
                nw229_[int(4)] = p0
                nw229_[int(5)] = p0
                nw229_[int(6)] = p0
                nw229_[int(7)] = p0
                nw229_[int(8)] = p0
                nw229_[int(9)] = p0
                nw229_[int(10)] = p0
                nw229_[int(11)] = p0
                nw229_[int(12)] = p0
                nw229_[int(13)] = p0
                nw229_[int(14)] = p0
                nw229_[int(15)] = p0
                nw229_[int(16)] = p0
                nw229_[int(17)] = p0
                nw229_[int(18)] = p0
                nw229_[int(19)] = _dafny.CodePoint('x')
                nw229_[int(20)] = p0
                nw229_[int(21)] = p0
                nw229_[int(22)] = default__.fm5((0) - (d_1220_v1_), p1, globalState)
                d_1301_v59_ = nw229_
                index240_ = default__.safeIndex(144, (d_1301_v59_).length(0))
                (d_1301_v59_)[index240_] = p0
                index241_ = default__.safeIndex(144, (d_1301_v59_).length(0))
                (d_1301_v59_)[index241_] = p0
                d_1302_v60_: D1
                d_1302_v60_ = D1_DC3(p1)
                d_1303_v61_: C3
                nw230_ = C3()
                nw230_.ctor__(d_1302_v60_)
                d_1303_v61_ = nw230_
                d_1304_v62_: _dafny.Map
                d_1304_v62_ = _dafny.Map({(self).f37: p3})
                d_1304_v62_ = (d_1304_v62_).set(p1, p3)
                d_1305_v63_: _dafny.Map
                d_1305_v63_ = _dafny.Map({len(d_1223_v4_): d_1220_v1_})
                (globalState).f11 = (d_1220_v1_) not in (d_1305_v63_)
            arr4_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
            index242_ = default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
            arr4_[index242_] = d_1220_v1_
            arr5_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
            index243_ = default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
            arr5_[index243_] = d_1220_v1_
            d_1306_v64_: D1
            d_1306_v64_ = D1_DC5(p1, p1, d_1220_v1_, (self).f37, not((p0) in (d_1221_v2_)))
            source18_ = d_1306_v64_
            if source18_.is_DC4:
                d_1307___mcc_h8_ = source18_.cf5
                d_1308___mcc_h9_ = source18_.cf6
                d_1309_cf6_ = d_1308___mcc_h9_
                d_1310_cf5_ = d_1307___mcc_h8_
                (globalState).f14 = (0) - ((d_1222_v3_)[default__.safeIndex(((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))], len(d_1222_v3_))])
                (globalState).f27 = not(p1)
                d_1311_v65_: _dafny.Array
                def lambda51_(d_1312_v1_):
                    def lambda52_(d_1313_i5_):
                        return _dafny.Map({_dafny.Map({True: d_1312_v1_}): (self).f37})

                    return lambda52_

                init29_ = lambda51_(d_1220_v1_)
                nw231_ = _dafny.Array(None, 26)
                for i0_29_ in range(nw231_.length(0)):
                    nw231_[i0_29_] = init29_(i0_29_)
                d_1311_v65_ = nw231_
                d_1314_v66_: _dafny.Seq
                d_1314_v66_ = _dafny.SeqWithoutIsStrInference([(self).f37])
                d_1315_v67_: _dafny.Map
                d_1315_v67_ = _dafny.Map({p1: len(d_1314_v66_)})
                d_1316_v68_: _dafny.Map
                d_1316_v68_ = _dafny.Map({d_1315_v67_: d_1309_cf6_})
                index244_ = default__.safeIndex(506, (d_1311_v65_).length(0))
                (d_1311_v65_)[index244_] = d_1316_v68_
                d_1317_v69_: _dafny.MultiSet
                d_1317_v69_ = _dafny.MultiSet([d_1310_cf5_])
                d_1318_v70_: _dafny.Map
                d_1318_v70_ = _dafny.Map({True: p0})
                d_1319_v71_: _dafny.Map
                d_1319_v71_ = _dafny.Map({(self).f37: d_1317_v69_})
                d_1320_v72_: _dafny.Array
                nw232_ = _dafny.Array(None, 12)
                nw232_[int(0)] = d_1317_v69_
                nw232_[int(1)] = d_1317_v69_
                nw232_[int(2)] = d_1317_v69_
                nw232_[int(3)] = d_1317_v69_
                nw232_[int(4)] = (d_1317_v69_).set(((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))], default__.abs(d_1310_cf5_))
                nw232_[int(5)] = _dafny.MultiSet([d_1310_cf5_, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))]])
                nw232_[int(6)] = d_1317_v69_
                nw232_[int(7)] = d_1317_v69_
                nw232_[int(8)] = _dafny.MultiSet([((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))], default__.fm2(d_1221_v2_, default__.fm24(_dafny.SeqWithoutIsStrInference([(self).f37, d_1309_cf6_, (self).f37, True]), d_1309_cf6_, globalState), d_1318_v70_, p1, globalState)])
                nw232_[int(9)] = d_1317_v69_
                nw232_[int(10)] = (d_1317_v69_) | (d_1317_v69_)
                nw232_[int(11)] = ((d_1319_v71_)[(self).f37] if ((self).f37) in (d_1319_v71_) else d_1317_v69_)
                d_1320_v72_ = nw232_
                index245_ = default__.safeIndex(691, (d_1320_v72_).length(0))
                (d_1320_v72_)[index245_] = d_1317_v69_
                d_1321_v74_: _dafny.MultiSet
                d_1321_v74_ = _dafny.MultiSet([d_1315_v67_])
                index246_ = default__.safeIndex(506, (d_1311_v65_).length(0))
                index247_ = default__.safeIndex(691, (d_1320_v72_).length(0))
                def iife137_():
                    coll59_ = _dafny.Map()
                    compr_59_: _dafny.Map
                    for compr_59_ in ((d_1321_v74_) | (d_1321_v74_)).Elements:
                        d_1322_v73_: _dafny.Map = compr_59_
                        if (d_1322_v73_) in ((d_1321_v74_) | (d_1321_v74_)):
                            coll59_[d_1322_v73_] = (d_1310_cf5_) == (d_1310_cf5_)
                    return _dafny.Map(coll59_)
                rhs240_ = iife137_()

                rhs241_ = _dafny.SeqWithoutIsStrInference([len((_dafny.SeqWithoutIsStrInference([p0 for d_1323_i7_ in range(default__.abs(-626))])).set(default__.safeIndex(780, len(_dafny.SeqWithoutIsStrInference([p0 for d_1324_i7_ in range(default__.abs(-626))]))), p0)) for d_1325_i6_ in range(default__.abs(135))])
                rhs242_ = (d_1317_v69_).intersection((d_1317_v69_) | (d_1317_v69_))
                lhs208_ = d_1311_v65_
                lhs209_ = default__.safeIndex(506, (d_1311_v65_).length(0))
                lhs210_ = d_1320_v72_
                lhs211_ = default__.safeIndex(691, (d_1320_v72_).length(0))
                lhs208_[lhs209_] = rhs240_
                d_1222_v3_ = rhs241_
                lhs210_[lhs211_] = rhs242_
                d_1326_v75_: _dafny.Array
                nw233_ = _dafny.Array(False, 21)
                d_1326_v75_ = nw233_
                index248_ = default__.safeIndex(906, (d_1326_v75_).length(0))
                (d_1326_v75_)[index248_] = (self).f37
                d_1327_v76_: D8
                d_1327_v76_ = D8_DC20(d_1310_cf5_, len((_dafny.Map({d_1309_cf6_: len(d_1222_v3_)})).set(True, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))])))
                index249_ = default__.safeIndex(906, (d_1326_v75_).length(0))
                arr6_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
                index250_ = default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
                rhs243_ = not ((p1 if (self).f37 else d_1309_cf6_)) or (d_1309_cf6_)
                rhs244_ = default__.safeDivisionInt((d_1327_v76_).cf31, 659)
                rhs245_ = default__.fm4(globalState)
                rhs246_ = d_1310_cf5_
                rhs247_ = self.f38
                lhs212_ = globalState
                lhs213_ = globalState
                lhs214_ = d_1326_v75_
                lhs215_ = default__.safeIndex(906, (d_1326_v75_).length(0))
                lhs216_ = (d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]
                lhs217_ = default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))
                lhs218_ = self
                lhs212_.f27 = rhs243_
                lhs213_.f24 = rhs244_
                lhs214_[lhs215_] = rhs245_
                lhs216_[lhs217_] = rhs246_
                lhs218_.f38 = rhs247_
            elif source18_.is_DC5:
                d_1328___mcc_h10_ = source18_.cf7
                d_1329___mcc_h11_ = source18_.cf8
                d_1330___mcc_h12_ = source18_.cf9
                d_1331___mcc_h13_ = source18_.cf10
                d_1332___mcc_h14_ = source18_.cf11
                d_1333_cf11_ = d_1332___mcc_h14_
                d_1334_cf10_ = d_1331___mcc_h13_
                d_1335_cf9_ = d_1330___mcc_h12_
                d_1336_cf8_ = d_1329___mcc_h11_
                d_1337_cf7_ = d_1328___mcc_h10_
                d_1338_v77_: _dafny.Array
                def lambda53_(d_1339_cf10_):
                    def lambda54_(d_1340_i8_):
                        return d_1339_cf10_

                    return lambda54_

                init30_ = lambda53_(d_1334_cf10_)
                nw234_ = _dafny.Array(None, 28)
                for i0_30_ in range(nw234_.length(0)):
                    nw234_[i0_30_] = init30_(i0_30_)
                d_1338_v77_ = nw234_
                index251_ = default__.safeIndex(801, (d_1338_v77_).length(0))
                (d_1338_v77_)[index251_] = ((d_1223_v4_)[(0) - (((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))])] if ((0) - (((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))])) in (d_1223_v4_) else False)
                d_1341_v78_: _dafny.Seq
                d_1341_v78_ = _dafny.SeqWithoutIsStrInference([d_1337_cf7_])
                d_1342_v79_: D1
                d_1342_v79_ = D1_DC3(p1)
                d_1343_v80_: C2
                nw235_ = C2()
                nw235_.ctor__(_dafny.Map({d_1341_v78_: 834}), (d_1228_v9_).f43, d_1342_v79_)
                d_1343_v80_ = nw235_
                d_1344_v81_: _dafny.Map
                d_1344_v81_ = _dafny.Map({d_1343_v80_: not((self).f37)})
                index252_ = default__.safeIndex(801, (d_1338_v77_).length(0))
                rhs248_ = not(d_1336_cf8_)
                rhs249_ = not((self).f37)
                rhs250_ = (167) == (default__.safeModuloInt(d_1335_cf9_, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))]))
                rhs251_ = (d_1344_v81_) | (d_1344_v81_)
                lhs219_ = globalState
                lhs220_ = d_1338_v77_
                lhs221_ = default__.safeIndex(801, (d_1338_v77_).length(0))
                lhs219_.f27 = rhs248_
                lhs220_[lhs221_] = rhs249_
                d_1337_cf7_ = rhs250_
                d_1344_v81_ = rhs251_
                (globalState).f24 = ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))]
                d_1345_v82_: _dafny.MultiSet
                d_1345_v82_ = _dafny.MultiSet([p1])
                d_1346_v83_: _dafny.Map
                d_1346_v83_ = _dafny.Map({len(d_1221_v2_): ((d_1345_v82_)[p1] if (p1) in (d_1345_v82_) else 781)})
                d_1347_v84_: _dafny.Map
                d_1347_v84_ = _dafny.Map({True: p0})
                d_1348_v85_: _dafny.Map
                d_1348_v85_ = _dafny.Map({d_1220_v1_: d_1347_v84_})
                index253_ = default__.safeIndex(801, (d_1338_v77_).length(0))
                def iife138_():
                    coll60_ = _dafny.Map()
                    compr_60_: int
                    for compr_60_ in _dafny.IntegerRange(561, 261):
                        d_1349_v86_: int = compr_60_
                        if ((561) <= (d_1349_v86_)) and ((d_1349_v86_) < (261)):
                            coll60_[(d_1349_v86_) * (((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))])] = d_1335_cf9_
                    return _dafny.Map(coll60_)
                def iife139_():
                    coll61_ = _dafny.Map()
                    compr_61_: int
                    for compr_61_ in _dafny.IntegerRange(561, 261):
                        d_1350_v86_: int = compr_61_
                        if ((561) <= (d_1350_v86_)) and ((d_1350_v86_) < (261)):
                            coll61_[(d_1350_v86_) * (((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))])] = d_1335_cf9_
                    return _dafny.Map(coll61_)
                rhs252_ = not((default__.safeDivisionInt(d_1220_v1_, len(d_1341_v78_))) <= (d_1220_v1_))
                rhs253_ = (len(d_1346_v83_)) not in ((d_1222_v3_) + (d_1222_v3_))
                rhs254_ = default__.fm2((d_1221_v2_) + ((D2_DC7(d_1333_cf11_, (self).f37, (d_1221_v2_).set(default__.safeIndex(((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))], len(d_1221_v2_)), p0))).cf15), (d_1228_v9_).f43, (((d_1348_v85_)[len(iife138_()
                )] if (len(iife139_()
                )) in (d_1348_v85_) else _dafny.Map({p1: p0})) if d_1333_cf11_ else d_1347_v84_), (d_1338_v77_)[default__.safeIndex(801, (d_1338_v77_).length(0))], globalState)
                rhs255_ = d_1337_cf7_
                lhs222_ = d_1338_v77_
                lhs223_ = default__.safeIndex(801, (d_1338_v77_).length(0))
                lhs224_ = globalState
                lhs225_ = globalState
                lhs222_[lhs223_] = rhs252_
                lhs224_.f11 = rhs253_
                lhs225_.f14 = rhs254_
                d_1334_cf10_ = rhs255_
                d_1351_v87_: _dafny.Map
                d_1351_v87_ = _dafny.Map({True: not(d_1336_cf8_)})
                d_1336_cf8_ = not((d_1220_v1_) <= (len(d_1351_v87_)))
            elif True:
                d_1352___mcc_h15_ = source18_.cf4
                d_1353_cf4_ = d_1352___mcc_h15_
                d_1222_v3_ = (d_1222_v3_) + (d_1222_v3_)
                d_1354_v88_: _dafny.Array
                nw236_ = _dafny.Array(None, 15)
                nw236_[int(0)] = d_1353_cf4_
                nw236_[int(1)] = (self).f37
                nw236_[int(2)] = d_1353_cf4_
                nw236_[int(3)] = (self).f37
                nw236_[int(4)] = (self).f37
                nw236_[int(5)] = p1
                nw236_[int(6)] = (self).f37
                nw236_[int(7)] = d_1353_cf4_
                nw236_[int(8)] = p1
                nw236_[int(9)] = (self).f37
                nw236_[int(10)] = p1
                nw236_[int(11)] = d_1353_cf4_
                nw236_[int(12)] = (self).f37
                nw236_[int(13)] = d_1353_cf4_
                nw236_[int(14)] = False
                d_1354_v88_ = nw236_
                d_1355_v89_: _dafny.Seq
                d_1355_v89_ = _dafny.SeqWithoutIsStrInference([d_1354_v88_, d_1354_v88_])
                d_1356_v90_: _dafny.Array
                nw237_ = _dafny.Array(None, 20)
                nw237_[int(0)] = d_1354_v88_
                nw237_[int(1)] = d_1354_v88_
                nw237_[int(2)] = d_1354_v88_
                nw237_[int(3)] = d_1354_v88_
                nw237_[int(4)] = d_1354_v88_
                nw237_[int(5)] = d_1354_v88_
                nw237_[int(6)] = d_1354_v88_
                nw237_[int(7)] = d_1354_v88_
                nw237_[int(8)] = d_1354_v88_
                nw237_[int(9)] = d_1354_v88_
                nw237_[int(10)] = d_1354_v88_
                nw237_[int(11)] = d_1354_v88_
                nw237_[int(12)] = (d_1355_v89_)[default__.safeIndex(((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))], len(d_1355_v89_))]
                nw237_[int(13)] = d_1354_v88_
                nw237_[int(14)] = d_1354_v88_
                nw237_[int(15)] = (d_1354_v88_ if p1 else d_1354_v88_)
                nw237_[int(16)] = d_1354_v88_
                nw237_[int(17)] = d_1354_v88_
                nw237_[int(18)] = d_1354_v88_
                nw237_[int(19)] = d_1354_v88_
                d_1356_v90_ = nw237_
                index254_ = default__.safeIndex(21, (d_1356_v90_).length(0))
                (d_1356_v90_)[index254_] = d_1354_v88_
                index255_ = default__.safeIndex(21, (d_1356_v90_).length(0))
                (d_1356_v90_)[index255_] = d_1354_v88_
                (globalState).f27 = ((self).f37) and ((self).f37)
                d_1357_v91_: _dafny.MultiSet
                d_1357_v91_ = _dafny.MultiSet([p0, p0])
                d_1358_v92_: _dafny.Seq
                d_1358_v92_ = _dafny.SeqWithoutIsStrInference([d_1353_cf4_])
                d_1359_v93_: D1
                d_1359_v93_ = D1_DC3((self).f37)
                d_1360_v94_: C1
                nw238_ = C1()
                nw238_.ctor__(((d_1357_v91_)[p0] if (p0) in (d_1357_v91_) else len(d_1358_v92_)), ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))])[default__.safeIndex(676, ((d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]).length(0))], d_1359_v93_)
                d_1360_v94_ = nw238_
            d_1361_v95_: D1
            d_1361_v95_ = D1_DC3(p1)
            d_1362_v96_: C3
            nw239_ = C3()
            nw239_.ctor__(d_1361_v95_)
            d_1362_v96_ = nw239_
            d_1363_v97_: _dafny.Set
            d_1363_v97_ = _dafny.Set({d_1362_v96_, d_1362_v96_, d_1362_v96_})
            d_1364_v98_: _dafny.Seq
            d_1364_v98_ = _dafny.SeqWithoutIsStrInference([d_1223_v4_])
            d_1365_v99_: _dafny.Map
            d_1365_v99_ = _dafny.Map({d_1363_v97_: (d_1364_v98_)[default__.safeIndex(d_1220_v1_, len(d_1364_v98_))]})
            d_1223_v4_ = (((d_1365_v99_)[d_1363_v97_] if (d_1363_v97_) in (d_1365_v99_) else _dafny.Map({d_1220_v1_: p1}))) | ((d_1223_v4_) | (d_1223_v4_))
        if p1:
            (globalState).f14 = d_1220_v1_
            (globalState).f0 = d_1220_v1_
            d_1366_v100_: _dafny.Map
            d_1366_v100_ = _dafny.Map({(d_1226_v7_)[default__.safeIndex(582, (d_1226_v7_).length(0))]: (self).f37})
            d_1366_v100_ = (d_1366_v100_).set(p3, p1)
            index256_ = default__.safeIndex(582, (d_1226_v7_).length(0))
            def lambda55_(d_1367_i9_):
                return default__.safeModuloInt(d_1367_i9_, 848)

            init31_ = lambda55_
            nw240_ = _dafny.Array(None, 3)
            for i0_31_ in range(nw240_.length(0)):
                nw240_[i0_31_] = init31_(i0_31_)
            (d_1226_v7_)[index256_] = nw240_
            arr7_ = self.f38
            index257_ = default__.safeIndex(753, (self.f38).length(0))
            arr7_[index257_] = d_1221_v2_
            arr8_ = self.f38
            index258_ = default__.safeIndex(753, (self.f38).length(0))
            arr8_[index258_] = (d_1230_v12_)[default__.safeIndex(len((d_1221_v2_) + (default__.fm3(globalState))), len(d_1230_v12_))]
        elif True:
            def iife140_(_pat_let39_0):
                def iife141_(d_1368_dt__update__tmp_h0_):
                    def iife142_(_pat_let40_0):
                        def iife143_(d_1369_dt__update_hcf4_h0_):
                            return D1_DC3(d_1369_dt__update_hcf4_h0_)
                        return iife143_(_pat_let40_0)
                    return iife142_((self).f37)
                return iife141_(_pat_let39_0)
            source19_ = iife140_(D1_DC3(p1))
            if source19_.is_DC4:
                d_1370___mcc_h16_ = source19_.cf5
                d_1371___mcc_h17_ = source19_.cf6
                d_1372_cf6_ = d_1371___mcc_h17_
                d_1373_cf5_ = d_1370___mcc_h16_
                (globalState).f14 = d_1373_cf5_
                d_1374_v101_: _dafny.MultiSet
                d_1374_v101_ = _dafny.MultiSet([d_1373_cf5_, len(_dafny.Set({p0})), d_1373_cf5_])
                d_1375_v102_: _dafny.Seq
                d_1375_v102_ = _dafny.SeqWithoutIsStrInference([p1])
                d_1376_v103_: _dafny.Map
                d_1376_v103_ = _dafny.Map({(d_1374_v101_).cardinality: d_1375_v102_})
                d_1377_v104_: D0
                d_1377_v104_ = D0_DC1(((_dafny.MultiSet(((d_1376_v103_)[len(d_1221_v2_)] if (len(d_1221_v2_)) in (d_1376_v103_) else d_1375_v102_))) | (_dafny.MultiSet([p1]))).cardinality)
                d_1377_v104_ = d_1377_v104_
                d_1378_v105_: _dafny.Map
                d_1378_v105_ = _dafny.Map({len(d_1221_v2_): d_1221_v2_})
                d_1379_v107_: _dafny.Map
                def iife144_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(273, 73):
                        d_1380_v106_: int = compr_62_
                        if ((273) <= (d_1380_v106_)) and ((d_1380_v106_) < (73)):
                            coll62_[(d_1380_v106_) + (470)] = d_1221_v2_
                    return _dafny.Map(coll62_)
                d_1379_v107_ = _dafny.Map({d_1373_cf5_: iife144_()
                })
                d_1381_v109_: _dafny.Map
                d_1381_v109_ = _dafny.Map({d_1372_cf6_: d_1373_cf5_})
                d_1382_v112_: _dafny.Array
                nw241_ = _dafny.Array(None, 20)
                nw241_[int(0)] = (d_1378_v105_) | (((d_1379_v107_)[d_1373_cf5_] if (d_1373_cf5_) in (d_1379_v107_) else d_1378_v105_))
                nw241_[int(1)] = d_1378_v105_
                nw241_[int(2)] = d_1378_v105_
                nw241_[int(3)] = d_1378_v105_
                def iife145_():
                    coll63_ = _dafny.Map()
                    compr_63_: int
                    for compr_63_ in (d_1222_v3_).Elements:
                        d_1383_v108_: int = compr_63_
                        if (d_1383_v108_) in (d_1222_v3_):
                            coll63_[(d_1383_v108_) - (d_1220_v1_)] = d_1221_v2_
                    return _dafny.Map(coll63_)
                nw241_[int(4)] = iife145_()
                
                nw241_[int(5)] = _dafny.Map({len(d_1381_v109_): (d_1221_v2_).set(default__.safeIndex(d_1373_cf5_, len(d_1221_v2_)), p0)})
                nw241_[int(6)] = (_dafny.Map({d_1220_v1_: d_1221_v2_})) | (d_1378_v105_)
                nw241_[int(7)] = d_1378_v105_
                def iife146_():
                    coll64_ = _dafny.Map()
                    compr_64_: int
                    for compr_64_ in (d_1374_v101_).Elements:
                        d_1384_v110_: int = compr_64_
                        if (d_1384_v110_) in (d_1374_v101_):
                            coll64_[(d_1384_v110_) - (d_1373_cf5_)] = d_1221_v2_
                    return _dafny.Map(coll64_)
                nw241_[int(8)] = iife146_()
                
                def iife147_():
                    coll65_ = _dafny.Map()
                    compr_65_: int
                    for compr_65_ in _dafny.IntegerRange(57, 762):
                        d_1385_v111_: int = compr_65_
                        if ((57) <= (d_1385_v111_)) and ((d_1385_v111_) < (762)):
                            coll65_[(d_1385_v111_) + (len(d_1221_v2_))] = d_1221_v2_
                    return _dafny.Map(coll65_)
                nw241_[int(9)] = (iife147_()
                ).set((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([p0, p0]))).cardinality, d_1221_v2_)
                nw241_[int(10)] = d_1378_v105_
                nw241_[int(11)] = d_1378_v105_
                nw241_[int(12)] = (_dafny.Map({d_1220_v1_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sf"))})).set(51, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "st")))
                nw241_[int(13)] = d_1378_v105_
                nw241_[int(14)] = d_1378_v105_
                nw241_[int(15)] = d_1378_v105_
                nw241_[int(16)] = d_1378_v105_
                nw241_[int(17)] = d_1378_v105_
                nw241_[int(18)] = (d_1378_v105_ if p1 else d_1378_v105_)
                nw241_[int(19)] = (d_1378_v105_) | (d_1378_v105_)
                d_1382_v112_ = nw241_
                index259_ = default__.safeIndex(691, (d_1382_v112_).length(0))
                def iife148_():
                    coll66_ = _dafny.Map()
                    compr_66_: int
                    for compr_66_ in _dafny.IntegerRange(977, 817):
                        d_1386_v113_: int = compr_66_
                        if ((977) <= (d_1386_v113_)) and ((d_1386_v113_) < (817)):
                            coll66_[(d_1386_v113_) + (d_1373_cf5_)] = d_1221_v2_
                    return _dafny.Map(coll66_)
                (d_1382_v112_)[index259_] = (iife148_()
                ) | (_dafny.Map({d_1220_v1_: d_1221_v2_}))
                d_1387_v114_: _dafny.Map
                d_1387_v114_ = _dafny.Map({d_1375_v102_: (d_1378_v105_).set(146, d_1221_v2_)})
                index260_ = default__.safeIndex(691, (d_1382_v112_).length(0))
                (d_1382_v112_)[index260_] = (default__.fm35(_dafny.MultiSet(d_1222_v3_), len(d_1221_v2_), len(_dafny.SeqWithoutIsStrInference([(d_1222_v3_)[default__.safeIndex(d_1373_cf5_, len(d_1222_v3_))]])), d_1373_cf5_, globalState)) | (((d_1387_v114_)[d_1375_v102_] if (d_1375_v102_) in (d_1387_v114_) else d_1378_v105_))
                d_1388_v115_: D1
                d_1388_v115_ = D1_DC3((d_1375_v102_)[default__.safeIndex(len(d_1222_v3_), len(d_1375_v102_))])
                d_1389_v116_: C1
                nw242_ = C1()
                nw242_.ctor__(default__.safeModuloInt(d_1220_v1_, (D15_DC34(False, (0) - (len(d_1222_v3_)))).cf60), (0) - ((d_1220_v1_) - (d_1373_cf5_)), d_1388_v115_)
                d_1389_v116_ = nw242_
            elif source19_.is_DC5:
                d_1390___mcc_h18_ = source19_.cf7
                d_1391___mcc_h19_ = source19_.cf8
                d_1392___mcc_h20_ = source19_.cf9
                d_1393___mcc_h21_ = source19_.cf10
                d_1394___mcc_h22_ = source19_.cf11
                d_1395_cf11_ = d_1394___mcc_h22_
                d_1396_cf10_ = d_1393___mcc_h21_
                d_1397_cf9_ = d_1392___mcc_h20_
                d_1398_cf8_ = d_1391___mcc_h19_
                d_1399_cf7_ = d_1390___mcc_h18_
                d_1400_v117_: _dafny.Seq
                d_1400_v117_ = _dafny.SeqWithoutIsStrInference([d_1395_cf11_, True])
                d_1401_v118_: D10
                d_1401_v118_ = D10_DC25(p3, d_1398_cf8_, d_1400_v117_, d_1397_cf9_)
                d_1402_v119_: _dafny.Array
                nw243_ = _dafny.Array(None, 24)
                nw243_[int(0)] = d_1400_v117_
                nw243_[int(1)] = d_1400_v117_
                nw243_[int(2)] = d_1400_v117_
                nw243_[int(3)] = d_1400_v117_
                nw243_[int(4)] = d_1400_v117_
                nw243_[int(5)] = d_1400_v117_
                nw243_[int(6)] = d_1400_v117_
                nw243_[int(7)] = d_1400_v117_
                nw243_[int(8)] = (d_1400_v117_).set(default__.safeIndex(128, len(d_1400_v117_)), (self).f37)
                nw243_[int(9)] = (d_1400_v117_) + (d_1400_v117_)
                nw243_[int(10)] = d_1400_v117_
                nw243_[int(11)] = ((d_1400_v117_) + (d_1400_v117_)).set(default__.safeIndex(len(d_1221_v2_), len((d_1400_v117_) + (d_1400_v117_))), not(d_1395_cf11_))
                nw243_[int(12)] = d_1400_v117_
                nw243_[int(13)] = (_dafny.SeqWithoutIsStrInference([d_1395_cf11_])).set(default__.safeIndex(11, len(_dafny.SeqWithoutIsStrInference([d_1395_cf11_]))), not(p1))
                nw243_[int(14)] = d_1400_v117_
                nw243_[int(15)] = (d_1401_v118_).cf41
                nw243_[int(16)] = (d_1400_v117_) + (d_1400_v117_)
                nw243_[int(17)] = d_1400_v117_
                nw243_[int(18)] = d_1400_v117_
                nw243_[int(19)] = (_dafny.SeqWithoutIsStrInference([p1])) + (_dafny.SeqWithoutIsStrInference([d_1396_cf10_]))
                nw243_[int(20)] = d_1400_v117_
                nw243_[int(21)] = d_1400_v117_
                nw243_[int(22)] = d_1400_v117_
                nw243_[int(23)] = default__.fm0(p1, globalState)
                d_1402_v119_ = nw243_
                index261_ = default__.safeIndex(208, (d_1402_v119_).length(0))
                (d_1402_v119_)[index261_] = d_1400_v117_
                index262_ = default__.safeIndex(208, (d_1402_v119_).length(0))
                (d_1402_v119_)[index262_] = d_1400_v117_
                d_1403_v120_: D1
                d_1403_v120_ = D1_DC3(d_1396_cf10_)
                d_1404_v121_: C3
                nw244_ = C3()
                nw244_.ctor__(d_1403_v120_)
                d_1404_v121_ = nw244_
                d_1405_v122_: _dafny.Array
                def lambda56_(d_1406_p0_):
                    def lambda57_(d_1407_i10_):
                        return _dafny.MultiSet([d_1406_p0_])

                    return lambda57_

                init32_ = lambda56_(p0)
                nw245_ = _dafny.Array(None, 27)
                for i0_32_ in range(nw245_.length(0)):
                    nw245_[i0_32_] = init32_(i0_32_)
                d_1405_v122_ = nw245_
                d_1408_v123_: _dafny.Map
                d_1408_v123_ = _dafny.Map({d_1405_v122_: (_dafny.SeqWithoutIsStrInference([d_1220_v1_ for d_1409_i11_ in range(default__.abs(540))])) + (_dafny.SeqWithoutIsStrInference([d_1397_cf9_]))})
                d_1410_v124_: _dafny.Map
                d_1410_v124_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jytuxwrx"))})
                d_1408_v123_ = (d_1408_v123_).set(d_1405_v122_, _dafny.SeqWithoutIsStrInference([len(((d_1410_v124_)[p0] if (p0) in (d_1410_v124_) else _dafny.SeqWithoutIsStrInference([p0 for d_1411_i12_ in range(default__.abs(702))])))]))
                d_1412_v125_: D15
                d_1412_v125_ = D15_DC34(d_1398_cf8_, len(d_1221_v2_))
                pat_let_tv33_ = d_1220_v1_
                def iife149_(_pat_let41_0):
                    def iife150_(d_1413_dt__update__tmp_h1_):
                        def iife151_(_pat_let42_0):
                            def iife152_(d_1414_dt__update_hcf60_h0_):
                                return D15_DC34((d_1413_dt__update__tmp_h1_).cf59, d_1414_dt__update_hcf60_h0_)
                            return iife152_(_pat_let42_0)
                        return iife151_(pat_let_tv33_)
                    return iife150_(_pat_let41_0)
                d_1229_v10_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qocqs")): (d_1396_cf10_) and ((iife149_(d_1412_v125_)).cf59)})
            elif True:
                d_1415___mcc_h23_ = source19_.cf4
                d_1416_cf4_ = d_1415___mcc_h23_
                r0 = (p1) and (((d_1223_v4_)[(0) - (d_1220_v1_)] if ((0) - (d_1220_v1_)) in (d_1223_v4_) else d_1416_cf4_))
                d_1417_v126_: _dafny.MultiSet
                d_1417_v126_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([p0 for d_1418_i13_ in range(default__.abs(925))])) + (d_1221_v2_), (d_1221_v2_ if p1 else d_1221_v2_), default__.fm3(globalState), d_1221_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wrpsv"))])
                d_1417_v126_ = d_1417_v126_
                (globalState).f24 = (default__.safeModuloInt(d_1220_v1_, d_1220_v1_)) + (len((d_1228_v9_).f43))
                d_1419_v127_: str
                d_1419_v127_ = _dafny.CodePoint('r')
                d_1420_v128_: _dafny.Map
                d_1420_v128_ = _dafny.Map({p1: p0})
                d_1421_v130_: _dafny.Map
                d_1421_v130_ = _dafny.Map({(0) - (d_1220_v1_): 154})
                d_1422_v132_: _dafny.Array
                nw246_ = _dafny.Array(None, 29)
                nw246_[int(0)] = p1
                nw246_[int(1)] = (d_1220_v1_) == (default__.fm2(d_1221_v2_, (d_1228_v9_).f43, d_1420_v128_, (self).f37, globalState))
                nw246_[int(2)] = p1
                nw246_[int(3)] = (self).f37
                nw246_[int(4)] = ((0) - (d_1220_v1_)) > (-540)
                nw246_[int(5)] = p1
                nw246_[int(6)] = not((self).f37)
                nw246_[int(7)] = (p1) and (d_1416_cf4_)
                nw246_[int(8)] = False
                nw246_[int(9)] = p1
                def iife153_():
                    coll67_ = _dafny.Set()
                    compr_67_: str
                    for compr_67_ in (_dafny.Set({p0, _dafny.CodePoint('d')})).Elements:
                        d_1423_v129_: str = compr_67_
                        if (d_1423_v129_) in (_dafny.Set({p0, _dafny.CodePoint('d')})):
                            coll67_ = coll67_.union(_dafny.Set([d_1423_v129_]))
                    return _dafny.Set(coll67_)
                nw246_[int(10)] = (iife153_()
                ).ispropersubset(d_1240_v17_)
                nw246_[int(11)] = d_1416_cf4_
                nw246_[int(12)] = d_1416_cf4_
                nw246_[int(13)] = not (p1) or (p1)
                nw246_[int(14)] = p1
                nw246_[int(15)] = d_1416_cf4_
                nw246_[int(16)] = not(default__.fm4(globalState))
                def iife154_():
                    coll68_ = _dafny.Map()
                    compr_68_: int
                    for compr_68_ in _dafny.IntegerRange(598, 996):
                        d_1424_v131_: int = compr_68_
                        if ((598) <= (d_1424_v131_)) and ((d_1424_v131_) < (996)):
                            coll68_[(d_1424_v131_) * (d_1220_v1_)] = 256
                    return _dafny.Map(coll68_)
                nw246_[int(17)] = (d_1421_v130_) == (iife154_()
                )
                nw246_[int(18)] = (self).f37
                nw246_[int(19)] = p1
                nw246_[int(20)] = not(p1)
                nw246_[int(21)] = d_1416_cf4_
                nw246_[int(22)] = p1
                nw246_[int(23)] = p1
                nw246_[int(24)] = p1
                nw246_[int(25)] = p1
                nw246_[int(26)] = default__.fm4(globalState)
                nw246_[int(27)] = (self).f37
                nw246_[int(28)] = (self).f37
                d_1422_v132_ = nw246_
                d_1425_v133_: _dafny.MultiSet
                d_1425_v133_ = _dafny.MultiSet([d_1220_v1_, d_1220_v1_])
                index263_ = default__.safeIndex(615, (d_1422_v132_).length(0))
                (d_1422_v132_)[index263_] = (d_1425_v133_).issubset(_dafny.MultiSet([d_1220_v1_]))
                d_1426_v134_: _dafny.Seq
                d_1426_v134_ = _dafny.SeqWithoutIsStrInference([d_1416_cf4_, ((self).f37) or (not(False)), p1, p1, d_1416_cf4_])
                index264_ = default__.safeIndex(615, (d_1422_v132_).length(0))
                rhs256_ = d_1419_v127_
                rhs257_ = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1427_i14_ in range(default__.abs(330))])) < (_dafny.SeqWithoutIsStrInference([p0 for d_1428_i15_ in range(default__.abs(188))]))
                rhs258_ = (d_1426_v134_)[default__.safeIndex(d_1220_v1_, len(d_1426_v134_))]
                lhs226_ = globalState
                lhs227_ = d_1422_v132_
                lhs228_ = default__.safeIndex(615, (d_1422_v132_).length(0))
                d_1419_v127_ = rhs256_
                lhs226_.f11 = rhs257_
                lhs227_[lhs228_] = rhs258_
            (globalState).f1 = (self).f37
            (globalState).f11 = (self).f37
            d_1429_v135_: _dafny.Map
            d_1429_v135_ = (_dafny.Map({d_1220_v1_: p1}) if (self).f37 else d_1223_v4_)
            d_1429_v135_ = d_1429_v135_
            (globalState).f11 = True
        r0 = (self).f37
        pat_let_tv34_ = p0
        def lambda58_(source20_):
            d_1430___mcc_h24_ = source20_
            d_1431_cf37_ = d_1430___mcc_h24_
            return _dafny.SeqWithoutIsStrInference([pat_let_tv34_ for d_1432_i16_ in range(default__.abs(-476))])

        r1 = lambda58_(default__.fm36(globalState))
        return r0, r1

    @property
    def f37(self):
        return self._f37

class C5(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        self._f35: int = int(0)
        self._f36: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f35, f36, f28):
        (self)._f35 = f35
        (self)._f36 = f36
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        return (self).f36

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('k')

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1433_v0_: C1
        nw247_ = C1()
        nw247_.ctor__((self).f35, (self).f35, D1_DC3(not((self).f36)))
        d_1433_v0_ = nw247_
        if (self).f36:
            d_1434_v1_: _dafny.Seq
            d_1434_v1_ = _dafny.SeqWithoutIsStrInference([(self).f36, (self).f36])
            d_1435_v2_: _dafny.Seq
            d_1435_v2_ = _dafny.SeqWithoutIsStrInference([d_1433_v0_.f42])
            d_1436_v3_: _dafny.Seq
            d_1436_v3_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_1435_v2_)[default__.safeIndex((self).f35, len(d_1435_v2_))]), (0) - ((d_1433_v0_).f41)])
            d_1437_v4_: _dafny.Array
            nw248_ = _dafny.Array(_dafny.Seq({}), 9)
            d_1437_v4_ = nw248_
            d_1438_v5_: _dafny.Array
            d_1438_v5_ = d_1437_v4_
            d_1439_v6_: _dafny.Map
            d_1439_v6_ = _dafny.Map({d_1438_v5_: d_1433_v0_.f42})
            d_1440_v7_: _dafny.MultiSet
            d_1440_v7_ = _dafny.MultiSet([(self).f35, (self).f35])
            d_1441_v8_: _dafny.Array
            nw249_ = _dafny.Array(None, 23)
            nw249_[int(0)] = (d_1433_v0_).f41
            nw249_[int(1)] = len(d_1434_v1_)
            nw249_[int(2)] = len((d_1434_v1_).set(default__.safeIndex(d_1433_v0_.f42, len(d_1434_v1_)), (self).f36))
            nw249_[int(3)] = (self).f35
            nw249_[int(4)] = 334
            nw249_[int(5)] = d_1433_v0_.f42
            nw249_[int(6)] = (d_1435_v2_)[default__.safeIndex((0) - (d_1433_v0_.f42), len(d_1435_v2_))]
            nw249_[int(7)] = (d_1433_v0_).f41
            nw249_[int(8)] = d_1433_v0_.f42
            nw249_[int(9)] = (0) - ((d_1433_v0_).f41)
            nw249_[int(10)] = d_1433_v0_.f42
            nw249_[int(11)] = len(d_1439_v6_)
            nw249_[int(12)] = (d_1433_v0_).f41
            nw249_[int(13)] = 399
            nw249_[int(14)] = (0) - ((self).f35)
            nw249_[int(15)] = (d_1433_v0_).f41
            nw249_[int(16)] = (d_1433_v0_).f41
            nw249_[int(17)] = d_1433_v0_.f42
            nw249_[int(18)] = (0) - ((d_1433_v0_).f41)
            nw249_[int(19)] = d_1433_v0_.f42
            nw249_[int(20)] = (0) - ((d_1440_v7_).cardinality)
            nw249_[int(21)] = (0) - ((0) - ((d_1433_v0_).f41))
            nw249_[int(22)] = (d_1433_v0_).f41
            d_1441_v8_ = nw249_
            d_1442_v9_: _dafny.MultiSet
            d_1442_v9_ = _dafny.MultiSet([d_1441_v8_, d_1441_v8_])
            (globalState).f14 = (d_1442_v9_).cardinality
            d_1443_v10_: str
            d_1443_v10_ = _dafny.CodePoint('k')
            d_1444_v11_: _dafny.Seq
            d_1444_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jkw"))
            d_1445_v12_: _dafny.Map
            d_1445_v12_ = _dafny.Map({(d_1433_v0_).f41: d_1444_v11_})
            d_1446_v13_: _dafny.Array
            nw250_ = _dafny.Array(None, 17)
            nw250_[int(0)] = (default__.fm3(globalState)).set(default__.safeIndex((self).f35, len(default__.fm3(globalState))), d_1443_v10_)
            nw250_[int(1)] = d_1444_v11_
            nw250_[int(2)] = ((d_1445_v12_)[default__.fm2(d_1444_v11_, _dafny.Set({not((d_1433_v0_).fm7((self).f36, (self).f36, globalState)), (self).f36}), _dafny.Map({(self).f36: d_1443_v10_}), not((self).f36), globalState)] if (default__.fm2(d_1444_v11_, _dafny.Set({not((d_1433_v0_).fm7((self).f36, (self).f36, globalState)), (self).f36}), _dafny.Map({(self).f36: d_1443_v10_}), not((self).f36), globalState)) in (d_1445_v12_) else d_1444_v11_)
            nw250_[int(3)] = (d_1444_v11_).set(default__.safeIndex((self).f35, len(d_1444_v11_)), d_1443_v10_)
            nw250_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1443_v10_ for d_1447_i0_ in range(default__.abs(742))])
            nw250_[int(5)] = default__.fm3(globalState)
            nw250_[int(6)] = d_1444_v11_
            nw250_[int(7)] = d_1444_v11_
            nw250_[int(8)] = d_1444_v11_
            nw250_[int(9)] = _dafny.SeqWithoutIsStrInference([d_1443_v10_ for d_1448_i1_ in range(default__.abs(898))])
            nw250_[int(10)] = d_1444_v11_
            nw250_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wgehews"))
            nw250_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1443_v10_ for d_1449_i2_ in range(default__.abs(-46))])
            nw250_[int(13)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eagkks"))
            nw250_[int(14)] = d_1444_v11_
            nw250_[int(15)] = d_1444_v11_
            nw250_[int(16)] = default__.fm3(globalState)
            d_1446_v13_ = nw250_
            d_1450_v14_: C4
            nw251_ = C4()
            nw251_.ctor__(((self).f36) and ((self).f36), (d_1446_v13_ if (self).f36 else d_1446_v13_))
            d_1450_v14_ = nw251_
            (globalState).f19 = ((d_1433_v0_.f42) - ((d_1433_v0_).f41)) - ((d_1433_v0_).f41)
            d_1451_v15_: _dafny.Map
            d_1451_v15_ = _dafny.Map({(d_1433_v0_).f41: (d_1433_v0_).f41})
            d_1452_v16_: _dafny.Map
            d_1452_v16_ = _dafny.Map({d_1433_v0_.f42: d_1443_v10_})
            d_1453_v17_: _dafny.Map
            d_1453_v17_ = _dafny.Map({(self).f36: d_1443_v10_})
            d_1451_v15_ = (d_1451_v15_).set(((d_1433_v0_).f41) + (len(d_1452_v16_)), default__.fm2(d_1444_v11_, _dafny.Set({(d_1450_v14_).f37, (d_1450_v14_).f37}), d_1453_v17_, (d_1434_v1_)[default__.safeIndex((d_1433_v0_).f41, len(d_1434_v1_))], globalState))
            (globalState).f14 = 676
        elif True:
            (globalState).f0 = d_1433_v0_.f42
            d_1454_v18_: _dafny.Array
            nw252_ = _dafny.Array(False, 11)
            d_1454_v18_ = nw252_
            index265_ = default__.safeIndex(86, (d_1454_v18_).length(0))
            (d_1454_v18_)[index265_] = (self).f36
            index266_ = default__.safeIndex(86, (d_1454_v18_).length(0))
            (d_1454_v18_)[index266_] = (self).f36
            d_1455_v19_: _dafny.Seq
            d_1455_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wxdpojc"))
            d_1456_v20_: _dafny.Set
            d_1456_v20_ = _dafny.Set({not((d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]), (d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))], (d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]})
            d_1457_v21_: str
            d_1457_v21_ = _dafny.CodePoint('o')
            d_1458_v22_: _dafny.Map
            d_1458_v22_ = _dafny.Map({(self).f36: d_1457_v21_})
            d_1459_v23_: _dafny.Map
            d_1459_v23_ = _dafny.Map({d_1455_v19_: (d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]})
            d_1460_v24_: _dafny.Array
            nw253_ = _dafny.Array(None, 28)
            nw253_[int(0)] = (self).f35
            nw253_[int(1)] = (self).f35
            nw253_[int(2)] = (0) - ((d_1433_v0_).f41)
            nw253_[int(3)] = (d_1433_v0_).f41
            nw253_[int(4)] = (0) - ((d_1433_v0_.f42 if (self).f36 else 995))
            nw253_[int(5)] = default__.safeModuloInt(default__.fm2(d_1455_v19_, d_1456_v20_, d_1458_v22_, (self).f36, globalState), d_1433_v0_.f42)
            nw253_[int(6)] = (d_1433_v0_).f41
            nw253_[int(7)] = d_1433_v0_.f42
            nw253_[int(8)] = d_1433_v0_.f42
            nw253_[int(9)] = default__.safeModuloInt(d_1433_v0_.f42, d_1433_v0_.f42)
            nw253_[int(10)] = (d_1433_v0_).f41
            nw253_[int(11)] = default__.safeModuloInt((d_1433_v0_).f41, d_1433_v0_.f42)
            nw253_[int(12)] = (self).f35
            nw253_[int(13)] = (d_1433_v0_).f41
            nw253_[int(14)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrdrde"))))
            nw253_[int(15)] = default__.safeDivisionInt(d_1433_v0_.f42, (d_1433_v0_).f41)
            nw253_[int(16)] = (d_1433_v0_.f42) - (len(d_1459_v23_))
            nw253_[int(17)] = d_1433_v0_.f42
            nw253_[int(18)] = default__.fm2(d_1455_v19_, d_1456_v20_, d_1458_v22_, (self).f36, globalState)
            nw253_[int(19)] = (0) - (d_1433_v0_.f42)
            nw253_[int(20)] = d_1433_v0_.f42
            nw253_[int(21)] = d_1433_v0_.f42
            nw253_[int(22)] = default__.safeModuloInt(d_1433_v0_.f42, (d_1433_v0_).f41)
            nw253_[int(23)] = 264
            nw253_[int(24)] = (-590 if (self).f36 else d_1433_v0_.f42)
            nw253_[int(25)] = 977
            nw253_[int(26)] = (default__.fm2(d_1455_v19_, d_1456_v20_, d_1458_v22_, (self).f36, globalState) if (self).f36 else 825)
            nw253_[int(27)] = (0) - (default__.safeModuloInt(default__.fm2((d_1455_v19_).set(default__.safeIndex(d_1433_v0_.f42, len(d_1455_v19_)), d_1457_v21_), d_1456_v20_, d_1458_v22_, (d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))], globalState), (d_1433_v0_).f41))
            d_1460_v24_ = nw253_
            index267_ = default__.safeIndex(866, (d_1460_v24_).length(0))
            (d_1460_v24_)[index267_] = d_1433_v0_.f42
            d_1461_v25_: _dafny.Seq
            d_1461_v25_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]])).cardinality, (d_1433_v0_).f41])
            d_1462_v26_: _dafny.Map
            d_1462_v26_ = _dafny.Map({(self).f36: len(d_1461_v25_)})
            d_1463_v27_: D0
            d_1463_v27_ = D0_DC2(d_1433_v0_.f42, ((d_1462_v26_)[(d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]] if ((d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]) in (d_1462_v26_) else (d_1433_v0_).f41))
            index268_ = default__.safeIndex(866, (d_1460_v24_).length(0))
            (d_1460_v24_)[index268_] = (default__.safeModuloInt((d_1463_v27_).cf3, (self).f35)) - (len(default__.fm24(_dafny.SeqWithoutIsStrInference([(d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]]), (d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))], globalState)))
            d_1464_v28_: _dafny.Array
            nw254_ = _dafny.Array(_dafny.Array(None, 0), 12)
            d_1464_v28_ = nw254_
            d_1465_v29_: _dafny.Set
            d_1465_v29_ = _dafny.Set({default__.safeDivisionInt((d_1433_v0_).f41, len(d_1461_v25_)), (d_1460_v24_)[default__.safeIndex(866, (d_1460_v24_).length(0))], (d_1460_v24_)[default__.safeIndex(866, (d_1460_v24_).length(0))]})
            d_1466_v30_: _dafny.Seq
            d_1466_v30_ = _dafny.SeqWithoutIsStrInference([not((d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]), (d_1454_v18_)[default__.safeIndex(86, (d_1454_v18_).length(0))]])
            d_1467_v31_: _dafny.Map
            d_1467_v31_ = _dafny.Map({d_1465_v29_: d_1433_v0_.f42})
            d_1468_v32_: _dafny.Map
            d_1468_v32_ = _dafny.Map({len((d_1466_v30_).set(default__.safeIndex(len(d_1466_v30_), len(d_1466_v30_)), (self).f36)): len(d_1467_v31_)})
            d_1469_v33_: _dafny.Map
            d_1469_v33_ = _dafny.Map({d_1456_v20_: d_1468_v32_})
            d_1470_v34_: _dafny.Map
            d_1470_v34_ = d_1469_v33_
            index269_ = default__.safeIndex(86, (d_1454_v18_).length(0))
            def iife155_():
                coll69_ = _dafny.Set()
                compr_69_: int
                for compr_69_ in (d_1465_v29_).Elements:
                    d_1472_v35_: int = compr_69_
                    if (d_1472_v35_) in (d_1465_v29_):
                        coll69_ = coll69_.union(_dafny.Set([default__.safeModuloInt(d_1472_v35_, len(_dafny.Set({len(_dafny.Set({753}))})))]))
                return _dafny.Set(coll69_)
            rhs259_ = (d_1433_v0_.f42) >= ((len(d_1455_v19_)) + (len(_dafny.SeqWithoutIsStrInference([d_1457_v21_ for d_1471_i3_ in range(default__.abs(674))]))))
            rhs260_ = d_1464_v28_
            rhs261_ = (d_1465_v29_) - (iife155_()
            )
            rhs262_ = (d_1469_v33_) | (_dafny.Map({d_1456_v20_: d_1468_v32_}))
            rhs263_ = (self).f36
            lhs229_ = globalState
            lhs230_ = d_1454_v18_
            lhs231_ = default__.safeIndex(86, (d_1454_v18_).length(0))
            lhs229_.f27 = rhs259_
            d_1464_v28_ = rhs260_
            d_1465_v29_ = rhs261_
            d_1470_v34_ = rhs262_
            lhs230_[lhs231_] = rhs263_
            d_1473_v36_: _dafny.Array
            nw255_ = _dafny.Array(D12.default()(), 14)
            d_1473_v36_ = nw255_
            d_1473_v36_ = d_1473_v36_
        d_1474_i4_: int
        d_1474_i4_ = 0
        with _dafny.label("12"):
            while ((d_1433_v0_).f41) >= ((self).f35):
                with _dafny.c_label("12"):
                    if (d_1474_i4_) >= (100):
                        raise _dafny.Break("12")
                    d_1474_i4_ = (d_1474_i4_) + (1)
                    d_1475_v37_: _dafny.Map
                    d_1475_v37_ = _dafny.Map({(self).f36: (d_1433_v0_.f42) == ((d_1433_v0_).f41)})
                    d_1475_v37_ = (_dafny.Map({(self).f36: (self).f36})).set((self).f36, (186) >= ((d_1433_v0_).f41))
                    d_1476_v38_: _dafny.Array
                    def lambda59_(d_1477_v0_):
                        def lambda60_(d_1478_i5_):
                            return default__.safeDivisionInt(d_1478_i5_, (d_1477_v0_).f41)

                        return lambda60_

                    init33_ = lambda59_(d_1433_v0_)
                    nw256_ = _dafny.Array(None, 9)
                    for i0_33_ in range(nw256_.length(0)):
                        nw256_[i0_33_] = init33_(i0_33_)
                    d_1476_v38_ = nw256_
                    index270_ = default__.safeIndex(708, (d_1476_v38_).length(0))
                    (d_1476_v38_)[index270_] = (d_1433_v0_).f41
                    index271_ = default__.safeIndex(708, (d_1476_v38_).length(0))
                    (d_1476_v38_)[index271_] = d_1433_v0_.f42
                    d_1476_v38_ = d_1476_v38_
                    if (self).f36:
                        d_1479_v39_: D8
                        d_1479_v39_ = D8_DC19((d_1475_v37_) | (_dafny.Map({(self).f36: (self).f36})))
                        d_1479_v39_ = d_1479_v39_
                        d_1480_v40_: C3
                        nw257_ = C3()
                        nw257_.ctor__((self).f28)
                        d_1480_v40_ = nw257_
                        (globalState).f19 = d_1433_v0_.f42
                        d_1481_v41_: _dafny.Seq
                        d_1481_v41_ = _dafny.SeqWithoutIsStrInference([(self).f36])
                        d_1482_v42_: _dafny.Set
                        d_1482_v42_ = _dafny.Set({(self).f36, (self).f36})
                        d_1483_v43_: C2
                        nw258_ = C2()
                        nw258_.ctor__(_dafny.Map({d_1481_v41_: (d_1476_v38_)[default__.safeIndex(708, (d_1476_v38_).length(0))]}), (d_1482_v42_) | (d_1482_v42_), (self).f28)
                        d_1483_v43_ = nw258_
                        d_1484_v44_: _dafny.Map
                        d_1484_v44_ = _dafny.Map({default__.fm37((self).f36, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtyuy")), d_1433_v0_.f42, globalState): (0) - (len(((d_1483_v43_).f40) | (default__.fm24(d_1481_v41_, (self).f36, globalState))))})
                        d_1485_v45_: D8
                        d_1485_v45_ = D8_DC19(d_1475_v37_)
                        d_1486_v46_: D8
                        d_1486_v46_ = D8_DC22(d_1485_v45_)
                        d_1484_v44_ = (d_1484_v44_).set((d_1486_v46_ if (self).f36 else d_1486_v46_), -540)
                    elif True:
                        d_1487_v47_: _dafny.Map
                        d_1487_v47_ = _dafny.Map({986: (d_1433_v0_).f41})
                        (globalState).f0 = default__.safeDivisionInt((d_1433_v0_).f41, ((d_1487_v47_)[(self).f35] if ((self).f35) in (d_1487_v47_) else (d_1476_v38_)[default__.safeIndex(708, (d_1476_v38_).length(0))]))
                        d_1488_v48_: _dafny.Seq
                        d_1488_v48_ = _dafny.SeqWithoutIsStrInference([(self).f36])
                        d_1489_v49_: _dafny.Map
                        d_1489_v49_ = _dafny.Map({d_1488_v48_: (self).f36})
                        d_1489_v49_ = d_1489_v49_
                        (globalState).f24 = (d_1476_v38_)[default__.safeIndex(708, (d_1476_v38_).length(0))]
                        d_1490_v50_: _dafny.Array
                        nw259_ = _dafny.Array(D4.default()(), 5)
                        d_1490_v50_ = nw259_
                        rhs264_ = not((self).f36)
                        rhs265_ = default__.fm3(globalState)
                        rhs266_ = (d_1490_v50_ if (self).f36 else d_1490_v50_)
                        lhs232_ = globalState
                        lhs232_.f1 = rhs264_
                        r1 = rhs265_
                        d_1490_v50_ = rhs266_
                        d_1491_v51_: _dafny.Seq
                        d_1491_v51_ = _dafny.SeqWithoutIsStrInference([626, (0) - ((0) - ((d_1476_v38_)[default__.safeIndex(708, (d_1476_v38_).length(0))]))])
                        rhs267_ = (self).fm7(True, ((self).f36 if (self).f36 else (self).f36), globalState)
                        rhs268_ = False
                        rhs269_ = (d_1491_v51_)[default__.safeIndex((d_1433_v0_).f41, len(d_1491_v51_))]
                        rhs270_ = 614
                        rhs271_ = -181
                        lhs233_ = globalState
                        lhs234_ = globalState
                        lhs235_ = globalState
                        lhs236_ = d_1433_v0_
                        lhs237_ = d_1433_v0_
                        lhs233_.f11 = rhs267_
                        lhs234_.f11 = rhs268_
                        lhs235_.f0 = rhs269_
                        lhs236_.f42 = rhs270_
                        lhs237_.f42 = rhs271_
                    pass
            pass
        d_1492_v52_: str
        d_1492_v52_ = _dafny.CodePoint('m')
        d_1493_v53_: _dafny.Map
        d_1493_v53_ = _dafny.Map({(self).f35: d_1492_v52_})
        (globalState).f27 = (len((d_1493_v53_ if (self).f36 else d_1493_v53_))) != ((self).f35)
        d_1494_v54_: D12
        d_1494_v54_ = D12_DC29((self).f36)
        def lambda61_(source21_):
            if source21_.is_DC28:
                d_1495___mcc_h0_ = source21_.cf45
                d_1496___mcc_h1_ = source21_.cf46
                d_1497___mcc_h2_ = source21_.cf47
                d_1498___mcc_h3_ = source21_.cf48
                d_1499___mcc_h4_ = source21_.cf49
                d_1500_cf49_ = d_1499___mcc_h4_
                d_1501_cf48_ = d_1498___mcc_h3_
                d_1502_cf47_ = d_1497___mcc_h2_
                d_1503_cf46_ = d_1496___mcc_h1_
                d_1504_cf45_ = d_1495___mcc_h0_
                return ((self).f36) or (d_1504_cf45_)
            elif source21_.is_DC29:
                d_1505___mcc_h5_ = source21_.cf50
                d_1506_cf50_ = d_1505___mcc_h5_
                return (self).f36
            elif source21_.is_DC30:
                d_1507___mcc_h6_ = source21_.cf51
                d_1508___mcc_h7_ = source21_.cf52
                d_1509___mcc_h8_ = source21_.cf53
                d_1510___mcc_h9_ = source21_.cf54
                d_1511___mcc_h10_ = source21_.cf55
                d_1512_cf55_ = d_1511___mcc_h10_
                d_1513_cf54_ = d_1510___mcc_h9_
                d_1514_cf53_ = d_1509___mcc_h8_
                d_1515_cf52_ = d_1508___mcc_h7_
                d_1516_cf51_ = d_1507___mcc_h6_
                return d_1514_cf53_
            elif True:
                d_1517___mcc_h11_ = source21_.cf44
                d_1518_cf44_ = d_1517___mcc_h11_
                return not(True)

        (globalState).f1 = lambda61_(d_1494_v54_)
        d_1519_i6_: int
        d_1519_i6_ = 0
        with _dafny.label("13"):
            while not(not(not (((self).f36 if True else (self).f36)) or ((self).f36))):
                with _dafny.c_label("13"):
                    if (d_1519_i6_) >= (100):
                        raise _dafny.Break("13")
                    d_1519_i6_ = (d_1519_i6_) + (1)
                    d_1520_v55_: _dafny.Set
                    d_1520_v55_ = _dafny.Set({(self).f36})
                    (globalState).f19 = ((len(_dafny.Map({(d_1433_v0_).f41: (self).f36}))) - ((d_1433_v0_).f41)) - (default__.fm2(default__.fm3(globalState), d_1520_v55_, default__.fm18(not((self).f36), True, globalState), True, globalState))
                    d_1521_v56_: _dafny.Array
                    nw260_ = _dafny.Array(int(0), 22)
                    d_1521_v56_ = nw260_
                    index272_ = default__.safeIndex(253, (d_1521_v56_).length(0))
                    (d_1521_v56_)[index272_] = (_dafny.MultiSet([(0) - ((0) - (d_1433_v0_.f42)), d_1433_v0_.f42, (self).f35, (d_1433_v0_).f41, (self).f35])).cardinality
                    d_1522_v57_: _dafny.Map
                    d_1522_v57_ = _dafny.Map({d_1521_v56_: (self).f36})
                    d_1523_v58_: _dafny.Map
                    d_1523_v58_ = _dafny.Map({(d_1433_v0_).f41: (self).f36})
                    index273_ = default__.safeIndex(253, (d_1521_v56_).length(0))
                    (d_1521_v56_)[index273_] = (len(d_1522_v57_)) + (len(_dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1492_v52_ for d_1524_i7_ in range(default__.abs(520))])): False}), d_1523_v58_])))
                    d_1525_v59_: _dafny.Array
                    nw261_ = _dafny.Array(_dafny.Array(None, 0), 17)
                    d_1525_v59_ = nw261_
                    d_1526_v60_: _dafny.Seq
                    d_1526_v60_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ajymacajy"))
                    d_1527_v61_: _dafny.Map
                    d_1527_v61_ = _dafny.Map({(d_1433_v0_).f41: d_1526_v60_})
                    d_1528_v62_: _dafny.Array
                    nw262_ = _dafny.Array(None, 11)
                    nw262_[int(0)] = d_1526_v60_
                    nw262_[int(1)] = d_1526_v60_
                    nw262_[int(2)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sk"))
                    nw262_[int(3)] = d_1526_v60_
                    nw262_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "arrnfthm"))
                    nw262_[int(5)] = d_1526_v60_
                    nw262_[int(6)] = d_1526_v60_
                    nw262_[int(7)] = d_1526_v60_
                    nw262_[int(8)] = d_1526_v60_
                    nw262_[int(9)] = d_1526_v60_
                    nw262_[int(10)] = ((d_1527_v61_)[(d_1433_v0_).f41] if ((d_1433_v0_).f41) in (d_1527_v61_) else d_1526_v60_)
                    d_1528_v62_ = nw262_
                    index274_ = default__.safeIndex(909, (d_1525_v59_).length(0))
                    (d_1525_v59_)[index274_] = d_1528_v62_
                    index275_ = default__.safeIndex(909, (d_1525_v59_).length(0))
                    (d_1525_v59_)[index275_] = d_1528_v62_
                    d_1529_v63_: _dafny.MultiSet
                    d_1529_v63_ = _dafny.MultiSet([(0) - ((d_1521_v56_)[default__.safeIndex(253, (d_1521_v56_).length(0))])])
                    d_1530_v64_: _dafny.Seq
                    d_1530_v64_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet([d_1433_v0_.f42]), (d_1529_v63_) - (_dafny.MultiSet([828])), (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1433_v0_.f42]))) - (d_1529_v63_), d_1529_v63_])
                    index276_ = default__.safeIndex(253, (d_1521_v56_).length(0))
                    rhs272_ = (len(_dafny.SeqWithoutIsStrInference([d_1492_v52_ for d_1531_i8_ in range(default__.abs(233))]))) + (len(d_1526_v60_))
                    rhs273_ = (d_1530_v64_)[default__.safeIndex((d_1521_v56_)[default__.safeIndex(253, (d_1521_v56_).length(0))], len(d_1530_v64_))]
                    lhs238_ = d_1521_v56_
                    lhs239_ = default__.safeIndex(253, (d_1521_v56_).length(0))
                    lhs238_[lhs239_] = rhs272_
                    d_1529_v63_ = rhs273_
                    pass
            pass
        d_1532_v65_: _dafny.Array
        nw263_ = _dafny.Array(int(0), 22)
        d_1532_v65_ = nw263_
        d_1533_v66_: D2
        d_1533_v66_ = D2_DC8((self).f36, d_1532_v65_)
        r0 = d_1533_v66_
        r1 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('h') for d_1534_i9_ in range(default__.abs(589))])
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        hi7_ = (self).f35
        for d_1535_i0_ in range((self).f35, hi7_):
            d_1536_v0_: _dafny.Seq
            d_1536_v0_ = _dafny.SeqWithoutIsStrInference([(self).f36, False])
            d_1537_v1_: _dafny.Array
            nw264_ = _dafny.Array(None, 7)
            nw264_[int(0)] = (self).f36
            nw264_[int(1)] = not((self).f36)
            nw264_[int(2)] = p0
            nw264_[int(3)] = (self).f36
            nw264_[int(4)] = (self).f36
            nw264_[int(5)] = p0
            nw264_[int(6)] = (d_1536_v0_)[default__.safeIndex((self).f35, len(d_1536_v0_))]
            d_1537_v1_ = nw264_
            d_1538_v2_: _dafny.Map
            d_1538_v2_ = _dafny.Map({-292: d_1537_v1_})
            d_1539_v3_: D15
            d_1539_v3_ = D15_DC33((d_1538_v2_).set(d_1535_i0_, d_1537_v1_))
            source22_ = d_1539_v3_
            if source22_.is_DC34:
                d_1540___mcc_h0_ = source22_.cf59
                d_1541___mcc_h1_ = source22_.cf60
                d_1542_cf60_ = d_1541___mcc_h1_
                d_1543_cf59_ = d_1540___mcc_h0_
                (globalState).f11 = d_1543_cf59_
                d_1544_v4_: _dafny.Seq
                d_1544_v4_ = _dafny.SeqWithoutIsStrInference([177])
                d_1545_v5_: _dafny.MultiSet
                d_1545_v5_ = _dafny.MultiSet([d_1535_i0_, d_1542_cf60_])
                (globalState).f11 = (_dafny.MultiSet(d_1544_v4_)).isdisjoint((d_1545_v5_) | (d_1545_v5_))
                d_1546_v6_: _dafny.Seq
                d_1546_v6_ = _dafny.SeqWithoutIsStrInference([d_1536_v0_, d_1536_v0_, d_1536_v0_])
                d_1547_v7_: _dafny.Seq
                d_1547_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bmhkrq"))
                d_1548_v8_: _dafny.Seq
                d_1548_v8_ = _dafny.SeqWithoutIsStrInference([d_1547_v7_, d_1547_v7_, d_1547_v7_])
                d_1549_v9_: _dafny.MultiSet
                d_1549_v9_ = _dafny.MultiSet([p1, _dafny.CodePoint('h')])
                d_1550_v10_: _dafny.Set
                d_1550_v10_ = _dafny.Set({d_1543_cf59_})
                d_1551_v11_: _dafny.Map
                d_1551_v11_ = _dafny.Map({(self).f36: p1})
                d_1552_v12_: _dafny.Array
                nw265_ = _dafny.Array(None, 7)
                nw265_[int(0)] = (self).f35
                nw265_[int(1)] = (self).f35
                nw265_[int(2)] = 5
                nw265_[int(3)] = (d_1549_v9_).cardinality
                nw265_[int(4)] = d_1535_i0_
                nw265_[int(5)] = default__.fm2(d_1547_v7_, d_1550_v10_, d_1551_v11_, p0, globalState)
                nw265_[int(6)] = (self).f35
                d_1552_v12_ = nw265_
                d_1553_v13_: D2
                d_1553_v13_ = D2_DC8(d_1543_cf59_, d_1552_v12_)
                d_1554_v14_: _dafny.Map
                d_1554_v14_ = _dafny.Map({d_1535_i0_: (d_1553_v13_).cf17})
                d_1555_v15_: _dafny.Map
                d_1555_v15_ = _dafny.Map({len(d_1554_v14_): d_1547_v7_})
                d_1556_v16_: _dafny.Array
                nw266_ = _dafny.Array(None, 23)
                nw266_[int(0)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kogvc"))) + (d_1547_v7_)
                nw266_[int(1)] = d_1547_v7_
                nw266_[int(2)] = _dafny.SeqWithoutIsStrInference([p1 for d_1557_i1_ in range(default__.abs(-851))])
                nw266_[int(3)] = (d_1547_v7_) + (d_1547_v7_)
                nw266_[int(4)] = (d_1547_v7_).set(default__.safeIndex(d_1535_i0_, len(d_1547_v7_)), p1)
                nw266_[int(5)] = (d_1548_v8_)[default__.safeIndex(d_1542_cf60_, len(d_1548_v8_))]
                nw266_[int(6)] = d_1547_v7_
                nw266_[int(7)] = d_1547_v7_
                nw266_[int(8)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1558_i2_ in range(default__.abs(566))])
                nw266_[int(9)] = d_1547_v7_
                nw266_[int(10)] = (((d_1555_v15_)[d_1535_i0_] if (d_1535_i0_) in (d_1555_v15_) else d_1547_v7_)) + (default__.fm3(globalState))
                nw266_[int(11)] = d_1547_v7_
                nw266_[int(12)] = (d_1547_v7_).set(default__.safeIndex((0) - (d_1535_i0_), len(d_1547_v7_)), p1)
                nw266_[int(13)] = (d_1547_v7_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qdvhe")))
                nw266_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sksexkses"))
                nw266_[int(15)] = (d_1547_v7_ if p0 else default__.fm3(globalState))
                nw266_[int(16)] = d_1547_v7_
                nw266_[int(17)] = d_1547_v7_
                nw266_[int(18)] = d_1547_v7_
                nw266_[int(19)] = (d_1547_v7_) + (d_1547_v7_)
                nw266_[int(20)] = _dafny.SeqWithoutIsStrInference([p1 for d_1559_i3_ in range(default__.abs(463))])
                nw266_[int(21)] = ((d_1547_v7_).set(default__.safeIndex(d_1542_cf60_, len(d_1547_v7_)), p1)) + (d_1547_v7_)
                nw266_[int(22)] = d_1547_v7_
                d_1556_v16_ = nw266_
                d_1560_v17_: str
                d_1560_v17_ = _dafny.CodePoint('e')
                rhs274_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, default__.fm4(globalState), d_1543_cf59_, default__.fm4(globalState)])])
                rhs275_ = d_1556_v16_
                rhs276_ = d_1538_v2_
                rhs277_ = (d_1560_v17_ if (self).f36 else d_1560_v17_)
                rhs278_ = (d_1544_v4_) + ((d_1544_v4_ if p0 else (_dafny.SeqWithoutIsStrInference([d_1542_cf60_])).set(default__.safeIndex(d_1535_i0_, len(_dafny.SeqWithoutIsStrInference([d_1542_cf60_]))), 295)))
                d_1546_v6_ = rhs274_
                d_1556_v16_ = rhs275_
                d_1538_v2_ = rhs276_
                d_1560_v17_ = rhs277_
                d_1544_v4_ = rhs278_
                (globalState).f27 = d_1543_cf59_
            elif True:
                d_1561___mcc_h2_ = source22_.cf58
                d_1562_cf58_ = d_1561___mcc_h2_
                (globalState).f14 = d_1535_i0_
                d_1563_v18_: _dafny.Array
                nw267_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 10)
                d_1563_v18_ = nw267_
                d_1564_v19_: C4
                nw268_ = C4()
                nw268_.ctor__((self).f36, d_1563_v18_)
                d_1564_v19_ = nw268_
                d_1565_v20_: _dafny.Set
                d_1565_v20_ = _dafny.Set({p0})
                rhs279_ = (853) >= ((default__.fm25((self).f35, (self).f36, d_1536_v0_, d_1535_i0_, globalState)).cardinality)
                rhs280_ = d_1535_i0_
                rhs281_ = d_1564_v19_
                rhs282_ = default__.safeModuloInt((0) - (len(d_1565_v20_)), 382)
                rhs283_ = (default__.safeModuloInt((self).f35, 987)) >= (d_1535_i0_)
                lhs240_ = globalState
                lhs241_ = globalState
                lhs240_.f11 = rhs279_
                r0 = rhs280_
                d_1564_v19_ = rhs281_
                r0 = rhs282_
                lhs241_.f27 = rhs283_
                (globalState).f11 = p0
                (globalState).f11 = not((d_1564_v19_).f37)
            d_1536_v0_ = (d_1536_v0_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "occ"))) for d_1566_i5_ in range(default__.abs(733))]) for d_1567_i4_ in range(default__.abs(965))])), len(d_1536_v0_)), (self).f36)
            d_1568_v21_: _dafny.Set
            d_1568_v21_ = _dafny.Set({(self).f36, (self).f36})
            d_1569_v22_: C0
            nw269_ = C0()
            nw269_.ctor__(d_1568_v21_)
            d_1569_v22_ = nw269_
            d_1537_v1_ = (d_1537_v1_ if (p0 if p0 else False) else d_1537_v1_)
        hi8_ = (self).f35
        for d_1570_i6_ in range(37, hi8_):
            d_1571_v23_: _dafny.Set
            d_1571_v23_ = _dafny.Set({(self).f35})
            d_1572_v24_: _dafny.Map
            d_1572_v24_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1571_v23_]): d_1570_i6_})
            (globalState).f0 = ((d_1572_v24_)[_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f35}) for d_1573_i7_ in range(default__.abs(934))])] if (_dafny.SeqWithoutIsStrInference([_dafny.Set({(self).f35}) for d_1574_i7_ in range(default__.abs(934))])) in (d_1572_v24_) else (self).f35)
            d_1575_v25_: _dafny.Seq
            d_1575_v25_ = _dafny.SeqWithoutIsStrInference([d_1570_i6_, (0) - ((self).f35)])
            d_1576_v26_: _dafny.Array
            nw270_ = _dafny.Array(int(0), 8)
            d_1576_v26_ = nw270_
            d_1577_v27_: _dafny.Map
            d_1577_v27_ = _dafny.Map({d_1576_v26_: d_1570_i6_})
            d_1578_v28_: D8
            d_1578_v28_ = D8_DC21(p0, (self).f36, (self).f35)
            rhs284_ = (d_1576_v26_) not in ((d_1577_v27_) | (d_1577_v27_))
            rhs285_ = ((d_1578_v28_).cf35) <= (71)
            rhs286_ = d_1575_v25_
            rhs287_ = not(((p0 if not(False) else (self).f36)) == ((self).f36))
            lhs242_ = globalState
            lhs243_ = globalState
            lhs244_ = globalState
            lhs242_.f11 = rhs284_
            lhs243_.f27 = rhs285_
            d_1575_v25_ = rhs286_
            lhs244_.f11 = rhs287_
            r0 = (865 if True else (self).f35)
            if (d_1575_v25_) == (d_1575_v25_):
                d_1579_v29_: _dafny.Map
                d_1579_v29_ = _dafny.Map({p0: d_1570_i6_})
                d_1580_v30_: _dafny.MultiSet
                d_1580_v30_ = _dafny.MultiSet([(d_1575_v25_) < (d_1575_v25_)])
                d_1581_v31_: _dafny.Set
                d_1581_v31_ = _dafny.Set({True, True})
                d_1582_v32_: C0
                nw271_ = C0()
                nw271_.ctor__(d_1581_v31_)
                d_1582_v32_ = nw271_
                d_1583_v33_: _dafny.Map
                d_1583_v33_ = _dafny.Map({d_1582_v32_: d_1579_v29_})
                d_1584_v35_: _dafny.Seq
                d_1584_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nxso"))
                d_1585_v36_: _dafny.Set
                d_1585_v36_ = _dafny.Set({d_1584_v35_})
                def iife156_():
                    coll70_ = _dafny.Map()
                    compr_70_: _dafny.Seq
                    for compr_70_ in (d_1585_v36_).Elements:
                        d_1586_v34_: _dafny.Seq = compr_70_
                        if (d_1586_v34_) in (d_1585_v36_):
                            coll70_[d_1586_v34_] = not((self).f36)
                    return _dafny.Map(coll70_)
                rhs288_ = p0
                rhs289_ = (d_1579_v29_) | (((d_1583_v33_)[d_1582_v32_] if (d_1582_v32_) in (d_1583_v33_) else (default__.fm30(d_1570_i6_, iife156_()
                , globalState)).set((self).f36, len(d_1584_v35_))))
                rhs290_ = d_1570_i6_
                rhs291_ = _dafny.MultiSet([(self).f36])
                lhs245_ = globalState
                lhs246_ = globalState
                lhs245_.f27 = rhs288_
                d_1579_v29_ = rhs289_
                lhs246_.f0 = rhs290_
                d_1580_v30_ = rhs291_
                (globalState).f0 = (0) - ((0) - (d_1570_i6_))
                d_1587_v37_: _dafny.Map
                d_1587_v37_ = _dafny.Map({not(p0): p1})
                d_1588_v38_: D1
                d_1588_v38_ = D1_DC5(p0, (self).f36, (self).f35, (self).f36, (self).f36)
                d_1589_v39_: _dafny.Seq
                d_1590_v40_: bool
                d_1591_v41_: _dafny.Seq
                d_1592_v42_: int
                out9_: _dafny.Seq
                out10_: bool
                out11_: _dafny.Seq
                out12_: int
                out9_, out10_, out11_, out12_ = (self).m11(p0, default__.fm2(default__.fm3(globalState), _dafny.Set({not(p0)}), d_1587_v37_, p0, globalState), (d_1588_v38_).cf7, globalState)
                d_1589_v39_ = out9_
                d_1590_v40_ = out10_
                d_1591_v41_ = out11_
                d_1592_v42_ = out12_
                (globalState).f27 = d_1590_v40_
                d_1593_v43_: _dafny.Array
                nw272_ = _dafny.Array(_dafny.Seq({}), 13)
                d_1593_v43_ = nw272_
                d_1594_v44_: _dafny.Array
                d_1594_v44_ = d_1593_v43_
                d_1594_v44_ = d_1594_v44_
            elif True:
                d_1595_v45_: _dafny.Array
                nw273_ = _dafny.Array(_dafny.CodePoint('D'), 10)
                d_1595_v45_ = nw273_
                d_1596_v46_: _dafny.Map
                d_1596_v46_ = _dafny.Map({p0: d_1595_v45_})
                d_1597_v47_: D16
                d_1597_v47_ = D16_DC35(d_1595_v45_)
                d_1598_v48_: _dafny.Map
                d_1598_v48_ = _dafny.Map({d_1570_i6_: d_1595_v45_})
                d_1599_v49_: _dafny.Array
                nw274_ = _dafny.Array(None, 27)
                nw274_[int(0)] = d_1595_v45_
                nw274_[int(1)] = d_1595_v45_
                nw274_[int(2)] = d_1595_v45_
                nw274_[int(3)] = d_1595_v45_
                nw274_[int(4)] = d_1595_v45_
                nw274_[int(5)] = d_1595_v45_
                nw274_[int(6)] = d_1595_v45_
                nw274_[int(7)] = d_1595_v45_
                nw274_[int(8)] = d_1595_v45_
                nw274_[int(9)] = ((d_1596_v46_)[p0] if (p0) in (d_1596_v46_) else d_1595_v45_)
                nw274_[int(10)] = ((d_1596_v46_)[(self).f36] if ((self).f36) in (d_1596_v46_) else d_1595_v45_)
                nw274_[int(11)] = d_1595_v45_
                nw274_[int(12)] = d_1595_v45_
                nw274_[int(13)] = (d_1597_v47_).cf61
                nw274_[int(14)] = d_1595_v45_
                nw274_[int(15)] = d_1595_v45_
                nw274_[int(16)] = ((d_1598_v48_)[(self).f35] if ((self).f35) in (d_1598_v48_) else d_1595_v45_)
                nw274_[int(17)] = d_1595_v45_
                nw274_[int(18)] = d_1595_v45_
                nw274_[int(19)] = d_1595_v45_
                nw274_[int(20)] = d_1595_v45_
                nw274_[int(21)] = d_1595_v45_
                nw274_[int(22)] = d_1595_v45_
                nw274_[int(23)] = d_1595_v45_
                nw274_[int(24)] = d_1595_v45_
                nw274_[int(25)] = d_1595_v45_
                nw274_[int(26)] = d_1595_v45_
                d_1599_v49_ = nw274_
                index277_ = default__.safeIndex(948, (d_1599_v49_).length(0))
                (d_1599_v49_)[index277_] = d_1595_v45_
                index278_ = default__.safeIndex(948, (d_1599_v49_).length(0))
                (d_1599_v49_)[index278_] = d_1595_v45_
                index279_ = default__.safeIndex(378, (d_1576_v26_).length(0))
                (d_1576_v26_)[index279_] = 610
                d_1600_v50_: _dafny.Array
                nw275_ = _dafny.Array(None, 29)
                nw275_[int(0)] = (self).f36
                nw275_[int(1)] = (self).f36
                nw275_[int(2)] = (self).f36
                nw275_[int(3)] = p0
                nw275_[int(4)] = default__.fm4(globalState)
                nw275_[int(5)] = p0
                nw275_[int(6)] = (self).f36
                nw275_[int(7)] = (self).f36
                nw275_[int(8)] = (self).f36
                nw275_[int(9)] = (self).f36
                nw275_[int(10)] = p0
                nw275_[int(11)] = p0
                nw275_[int(12)] = (self).f36
                nw275_[int(13)] = (self).f36
                nw275_[int(14)] = p0
                nw275_[int(15)] = (self).f36
                nw275_[int(16)] = (self).f36
                nw275_[int(17)] = False
                nw275_[int(18)] = p0
                nw275_[int(19)] = p0
                nw275_[int(20)] = p0
                nw275_[int(21)] = (self).f36
                nw275_[int(22)] = p0
                nw275_[int(23)] = p0
                nw275_[int(24)] = (self).f36
                nw275_[int(25)] = p0
                nw275_[int(26)] = True
                nw275_[int(27)] = p0
                nw275_[int(28)] = p0
                d_1600_v50_ = nw275_
                d_1601_v51_: _dafny.Map
                d_1601_v51_ = _dafny.Map({p1: d_1600_v50_})
                d_1602_v52_: D12
                d_1602_v52_ = D12_DC27(d_1601_v51_)
                index280_ = default__.safeIndex(803, (d_1600_v50_).length(0))
                (d_1600_v50_)[index280_] = not((self).f36)
                index281_ = default__.safeIndex(378, (d_1576_v26_).length(0))
                index282_ = default__.safeIndex(803, (d_1600_v50_).length(0))
                rhs292_ = d_1570_i6_
                rhs293_ = d_1570_i6_
                rhs294_ = d_1602_v52_
                rhs295_ = not((self).f36)
                lhs247_ = d_1576_v26_
                lhs248_ = default__.safeIndex(378, (d_1576_v26_).length(0))
                lhs249_ = globalState
                lhs250_ = d_1600_v50_
                lhs251_ = default__.safeIndex(803, (d_1600_v50_).length(0))
                lhs247_[lhs248_] = rhs292_
                lhs249_.f24 = rhs293_
                d_1602_v52_ = rhs294_
                lhs250_[lhs251_] = rhs295_
                index283_ = default__.safeIndex(803, (d_1600_v50_).length(0))
                (d_1600_v50_)[index283_] = False
                d_1603_v53_: _dafny.MultiSet
                d_1603_v53_ = _dafny.MultiSet([d_1570_i6_])
                d_1604_v54_: _dafny.Map
                d_1604_v54_ = _dafny.Map({d_1603_v53_: default__.fm3(globalState)})
                d_1605_v55_: _dafny.Seq
                d_1605_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sm"))
                d_1606_v56_: _dafny.Seq
                d_1606_v56_ = _dafny.SeqWithoutIsStrInference([(d_1604_v54_).set(_dafny.MultiSet([(d_1576_v26_)[default__.safeIndex(378, (d_1576_v26_).length(0))]]), d_1605_v55_)])
                d_1607_v57_: _dafny.Seq
                d_1607_v57_ = _dafny.SeqWithoutIsStrInference([(self).f36, not((d_1600_v50_)[default__.safeIndex(803, (d_1600_v50_).length(0))]), (d_1600_v50_)[default__.safeIndex(803, (d_1600_v50_).length(0))]])
                d_1608_v58_: _dafny.MultiSet
                d_1608_v58_ = _dafny.MultiSet([d_1607_v57_, d_1607_v57_])
                d_1609_v59_: _dafny.Map
                d_1609_v59_ = _dafny.Map({(d_1606_v56_)[default__.safeIndex((self).f35, len(d_1606_v56_))]: d_1608_v58_})
                d_1610_v61_: _dafny.Seq
                d_1610_v61_ = _dafny.SeqWithoutIsStrInference([d_1603_v53_])
                def iife157_():
                    coll71_ = _dafny.Map()
                    compr_71_: _dafny.MultiSet
                    for compr_71_ in (d_1610_v61_).Elements:
                        d_1611_v60_: _dafny.MultiSet = compr_71_
                        if (d_1611_v60_) in (d_1610_v61_):
                            coll71_[d_1611_v60_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kaxulpph"))
                    return _dafny.Map(coll71_)
                d_1609_v59_ = (d_1609_v59_).set(iife157_()
                , (d_1608_v58_) - (d_1608_v58_))
                d_1612_v62_: _dafny.Map
                d_1612_v62_ = _dafny.Map({p0: (self).f35})
                d_1612_v62_ = (d_1612_v62_).set(p0, (0) - ((d_1570_i6_) + ((d_1576_v26_)[default__.safeIndex(378, (d_1576_v26_).length(0))])))
        hi9_ = ((self).f35) + ((self).f35)
        for d_1613_i8_ in range((self).f35, hi9_):
            d_1614_v63_: _dafny.Seq
            d_1614_v63_ = _dafny.SeqWithoutIsStrInference([(self).f36, p0])
            (globalState).f24 = len(((d_1614_v63_) + ((d_1614_v63_) + (d_1614_v63_))).set(default__.safeIndex((self).f35, len((d_1614_v63_) + ((d_1614_v63_) + (d_1614_v63_)))), False))
            if not(p0):
                d_1615_v64_: _dafny.Map
                d_1615_v64_ = _dafny.Map({len(d_1614_v63_): (d_1614_v63_)[default__.safeIndex((self).f35, len(d_1614_v63_))]})
                (globalState).f20 = d_1615_v64_
                r0 = (d_1613_i8_) - (d_1613_i8_)
                (globalState).f1 = p0
                d_1616_v65_: _dafny.Array
                nw276_ = _dafny.Array(None, 6)
                nw276_[int(0)] = d_1613_i8_
                nw276_[int(1)] = d_1613_i8_
                nw276_[int(2)] = (self).f35
                nw276_[int(3)] = d_1613_i8_
                nw276_[int(4)] = d_1613_i8_
                nw276_[int(5)] = len(_dafny.SeqWithoutIsStrInference([d_1613_i8_ for d_1617_i9_ in range(default__.abs(51))]))
                d_1616_v65_ = nw276_
                index284_ = default__.safeIndex(376, (d_1616_v65_).length(0))
                (d_1616_v65_)[index284_] = (self).f35
                index285_ = default__.safeIndex(376, (d_1616_v65_).length(0))
                (d_1616_v65_)[index285_] = len(((d_1614_v63_) + (default__.fm0((self).f36, globalState))) + (_dafny.SeqWithoutIsStrInference([p0])))
                d_1618_v66_: _dafny.Seq
                d_1618_v66_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bccnilcqs"))
                d_1619_v67_: _dafny.MultiSet
                d_1619_v67_ = _dafny.MultiSet([(_dafny.SeqWithoutIsStrInference([p1 for d_1620_i10_ in range(default__.abs(704))])) + ((d_1618_v66_).set(default__.safeIndex((self).f35, len(d_1618_v66_)), default__.fm5(626, p0, globalState))), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sdc")), d_1618_v66_, d_1618_v66_, default__.fm3(globalState)])
                d_1619_v67_ = d_1619_v67_
            elif True:
                (globalState).f1 = p0
                d_1621_v68_: _dafny.Seq
                d_1621_v68_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sysqjqqk"))
                d_1622_v69_: _dafny.Array
                nw277_ = _dafny.Array(None, 26)
                nw277_[int(0)] = (self).f35
                nw277_[int(1)] = (self).f35
                nw277_[int(2)] = d_1613_i8_
                nw277_[int(3)] = (self).f35
                nw277_[int(4)] = d_1613_i8_
                nw277_[int(5)] = d_1613_i8_
                nw277_[int(6)] = d_1613_i8_
                nw277_[int(7)] = (self).f35
                nw277_[int(8)] = d_1613_i8_
                nw277_[int(9)] = d_1613_i8_
                nw277_[int(10)] = d_1613_i8_
                nw277_[int(11)] = (self).f35
                nw277_[int(12)] = d_1613_i8_
                nw277_[int(13)] = (self).f35
                nw277_[int(14)] = len(d_1621_v68_)
                nw277_[int(15)] = (self).f35
                nw277_[int(16)] = d_1613_i8_
                nw277_[int(17)] = d_1613_i8_
                nw277_[int(18)] = (self).f35
                nw277_[int(19)] = d_1613_i8_
                nw277_[int(20)] = (self).f35
                nw277_[int(21)] = (self).f35
                nw277_[int(22)] = d_1613_i8_
                nw277_[int(23)] = (self).f35
                nw277_[int(24)] = (self).f35
                nw277_[int(25)] = (self).f35
                d_1622_v69_ = nw277_
                d_1623_v70_: D2
                d_1623_v70_ = D2_DC8(p0, d_1622_v69_)
                (globalState).f11 = (d_1623_v70_).cf16
                d_1624_v71_: _dafny.MultiSet
                d_1624_v71_ = _dafny.MultiSet([(self).f36, (self).f36])
                d_1625_v72_: _dafny.Set
                d_1625_v72_ = _dafny.Set({p0, p0, (self).f36, True, (self).f36})
                d_1626_v73_: _dafny.Map
                d_1626_v73_ = _dafny.Map({len(d_1625_v72_): d_1613_i8_})
                d_1627_v74_: _dafny.Map
                d_1627_v74_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qjhrkfk")): ((d_1626_v73_)[(self).f35] if ((self).f35) in (d_1626_v73_) else (0) - (len(d_1614_v63_)))})
                d_1628_v75_: _dafny.Map
                d_1628_v75_ = _dafny.Map({(d_1624_v71_).set((self).f36, default__.abs(len(d_1627_v74_))): ((self).f36) and (p0)})
                d_1629_v76_: D17
                d_1629_v76_ = D17_DC38(d_1624_v71_)
                d_1628_v75_ = (d_1628_v75_).set(((d_1629_v76_).cf68).set((d_1614_v63_)[default__.safeIndex(d_1613_i8_, len(d_1614_v63_))], default__.abs(d_1613_i8_)), (self).f36)
                (globalState).f19 = (self).f35
                (globalState).f11 = (len((d_1625_v72_).intersection(_dafny.Set({(self).f36})))) >= (d_1613_i8_)
            d_1630_v77_: _dafny.Set
            d_1630_v77_ = _dafny.Set({(self).f36})
            d_1631_v78_: C0
            nw278_ = C0()
            nw278_.ctor__((d_1630_v77_).intersection(d_1630_v77_))
            d_1631_v78_ = nw278_
            r0 = (d_1613_i8_) + ((self).f35)
        d_1632_v79_: _dafny.Seq
        d_1632_v79_ = _dafny.SeqWithoutIsStrInference([True, (self).f36])
        d_1633_v80_: _dafny.Map
        d_1633_v80_ = _dafny.Map({d_1632_v79_: (self).f35})
        d_1634_v81_: _dafny.Map
        d_1634_v81_ = _dafny.Map({(self).f35: p0})
        d_1635_v82_: C2
        nw279_ = C2()
        nw279_.ctor__(d_1633_v80_, default__.fm24(_dafny.SeqWithoutIsStrInference([((d_1634_v81_)[(self).f35] if ((self).f35) in (d_1634_v81_) else p0), (self).f36]), p0, globalState), (self).f28)
        d_1635_v82_ = nw279_
        d_1636_v83_: _dafny.Array
        def lambda62_(d_1637_i12_):
            return default__.safeDivisionInt(d_1637_i12_, (self).f35)

        init34_ = lambda62_
        nw280_ = _dafny.Array(None, 19)
        for i0_34_ in range(nw280_.length(0)):
            nw280_[i0_34_] = init34_(i0_34_)
        d_1636_v83_ = nw280_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1636_v83_).length(0)):
            d_1638_i11_: int = guard_loop_2_
            if (True) and (((0) <= (d_1638_i11_)) and ((d_1638_i11_) < ((d_1636_v83_).length(0)))):
                (d_1636_v83_)[(d_1638_i11_)] = (d_1638_i11_) - ((-408) + (-240))
        d_1639_v84_: _dafny.Array
        nw281_ = _dafny.Array(_dafny.Seq({}), 24)
        d_1639_v84_ = nw281_
        d_1640_v85_: D16
        d_1640_v85_ = D16_DC36(p0, (self).f36, d_1639_v84_)
        pat_let_tv35_ = p0
        d_1641_v86_: _dafny.MultiSet
        def iife158_(_pat_let43_0):
            def iife159_(d_1642_dt__update__tmp_h0_):
                def iife160_(_pat_let44_0):
                    def iife161_(d_1643_dt__update_hcf63_h0_):
                        return D16_DC36((d_1642_dt__update__tmp_h0_).cf62, d_1643_dt__update_hcf63_h0_, (d_1642_dt__update__tmp_h0_).cf64)
                    return iife161_(_pat_let44_0)
                return iife160_(pat_let_tv35_)
            return iife159_(_pat_let43_0)
        d_1641_v86_ = _dafny.MultiSet([(p0 if (self).f36 else (iife158_(d_1640_v85_)).cf62)])
        d_1641_v86_ = d_1641_v86_
        r0 = (self).f35
        return r0

    def m11(self, p0, p1, p2, globalState):
        r0: _dafny.Seq = _dafny.Seq({})
        r1: bool = False
        r2: _dafny.Seq = _dafny.Seq({})
        r3: int = int(0)
        d_1644_v0_: _dafny.Array
        def lambda63_(d_1645_p1_):
            def lambda64_(d_1646_i0_):
                return (d_1645_p1_) >= (d_1645_p1_)

            return lambda64_

        init35_ = lambda63_(p1)
        nw282_ = _dafny.Array(None, 22)
        for i0_35_ in range(nw282_.length(0)):
            nw282_[i0_35_] = init35_(i0_35_)
        d_1644_v0_ = nw282_
        index286_ = default__.safeIndex(997, (d_1644_v0_).length(0))
        (d_1644_v0_)[index286_] = not (p0) or ((self).f36)
        d_1647_v1_: _dafny.Seq
        d_1647_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lv"))
        d_1648_v2_: _dafny.Seq
        d_1648_v2_ = _dafny.SeqWithoutIsStrInference([(self).f35, (p1) * (p1), ((self).f35 if p0 else p1), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uxsjsd")))])
        index287_ = default__.safeIndex(997, (d_1644_v0_).length(0))
        rhs296_ = False
        rhs297_ = (self).fm7(True, False, globalState)
        rhs298_ = (d_1647_v1_) + (d_1647_v1_)
        rhs299_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbhxkruxg"))
        rhs300_ = (d_1648_v2_)[default__.safeIndex(p1, len(d_1648_v2_))]
        lhs252_ = d_1644_v0_
        lhs253_ = default__.safeIndex(997, (d_1644_v0_).length(0))
        lhs254_ = globalState
        lhs252_[lhs253_] = rhs296_
        lhs254_.f1 = rhs297_
        d_1647_v1_ = rhs298_
        d_1647_v1_ = rhs299_
        r3 = rhs300_
        d_1649_v3_: _dafny.Map
        d_1649_v3_ = _dafny.Map({d_1647_v1_: (self).f35})
        d_1650_v4_: C1
        nw283_ = C1()
        nw283_.ctor__(((self).f35) * (len(d_1649_v3_)), (self).f35, (self).f28)
        d_1650_v4_ = nw283_
        d_1651_v5_: C3
        nw284_ = C3()
        nw284_.ctor__((self).f28)
        d_1651_v5_ = nw284_
        d_1652_v6_: _dafny.Seq
        d_1652_v6_ = _dafny.SeqWithoutIsStrInference([d_1651_v5_])
        d_1651_v5_ = (d_1652_v6_)[default__.safeIndex(d_1650_v4_.f42, len(d_1652_v6_))]
        d_1653_v7_: _dafny.Set
        d_1653_v7_ = _dafny.Set({(d_1644_v0_)[default__.safeIndex(997, (d_1644_v0_).length(0))], True, (d_1644_v0_)[default__.safeIndex(997, (d_1644_v0_).length(0))]})
        d_1654_v8_: _dafny.MultiSet
        d_1654_v8_ = _dafny.MultiSet([d_1653_v7_, d_1653_v7_])
        d_1654_v8_ = ((d_1654_v8_).set(d_1653_v7_, default__.abs((d_1650_v4_).f41))) - (_dafny.MultiSet([d_1653_v7_, d_1653_v7_]))
        (globalState).f19 = ((self).f35 if False else (p1) * (d_1650_v4_.f42))
        (d_1650_v4_).f42 = (self).f35
        pat_let_tv36_ = d_1648_v2_
        def iife162_(_pat_let45_0):
            def iife163_(d_1655_dt__update__tmp_h0_):
                def iife164_(_pat_let46_0):
                    def iife165_(d_1656_dt__update_hcf71_h0_):
                        return D18_DC41(d_1656_dt__update_hcf71_h0_)
                    return iife165_(_pat_let46_0)
                return iife164_(pat_let_tv36_)
            return iife163_(_pat_let45_0)
        r0 = (iife162_(D18_DC41(d_1648_v2_))).cf71
        r1 = not((_dafny.MultiSet((D18_DC41(d_1648_v2_)).cf71)).ispropersubset(default__.fm26(not((self).f36), d_1647_v1_, p2, globalState)))
        d_1657_v9_: _dafny.Seq
        d_1657_v9_ = _dafny.SeqWithoutIsStrInference([(self).f36, p0, p0, (d_1644_v0_)[default__.safeIndex(997, (d_1644_v0_).length(0))], p2])
        r2 = (d_1657_v9_) + (_dafny.SeqWithoutIsStrInference([p2, (self).f36]))
        d_1658_v10_: str
        d_1658_v10_ = _dafny.CodePoint('v')
        d_1659_v11_: _dafny.Map
        d_1659_v11_ = _dafny.Map({(d_1647_v1_).set(default__.safeIndex((self).f35, len(d_1647_v1_)), d_1658_v10_): False})
        d_1660_v12_: _dafny.MultiSet
        d_1660_v12_ = _dafny.MultiSet([default__.fm30(len(_dafny.SeqWithoutIsStrInference([(d_1644_v0_)[default__.safeIndex(997, (d_1644_v0_).length(0))], (d_1644_v0_)[default__.safeIndex(997, (d_1644_v0_).length(0))], p0, p0, p0])), d_1659_v11_, globalState)])
        d_1661_v13_: _dafny.Map
        d_1661_v13_ = _dafny.Map({(self).f36: d_1650_v4_.f42})
        r3 = default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_1662_i1_ in range(default__.abs(265))])) + ((d_1647_v1_).set(default__.safeIndex(d_1650_v4_.f42, len(d_1647_v1_)), d_1658_v10_))), default__.safeModuloInt((self).f35, ((d_1660_v12_)[d_1661_v13_] if (d_1661_v13_) in (d_1660_v12_) else p1)))
        return r0, r1, r2, r3

    @property
    def f35(self):
        return self._f35
    @property
    def f36(self):
        return self._f36

class C6:
    def  __init__(self):
        self._f34: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f34):
        (self)._f34 = f34

    def fm16(self, p0, p1, globalState):
        return ((-879 if True else 824)) <= (default__.safeModuloInt((self).f34, 492))

    def m9(self, globalState):
        r0: bool = False
        d_1663_v0_: _dafny.Seq
        d_1663_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcntgxs"))
        d_1664_v1_: bool
        d_1664_v1_ = True
        d_1665_v2_: _dafny.Set
        d_1665_v2_ = _dafny.Set({d_1664_v1_})
        d_1666_v3_: str
        d_1666_v3_ = _dafny.CodePoint('e')
        hi10_ = (default__.fm2(d_1663_v0_, d_1665_v2_, _dafny.Map({True: d_1666_v3_}), not(d_1664_v1_), globalState)) + ((self).f34)
        for d_1667_i0_ in range((self).f34, hi10_):
            d_1668_v4_: D1
            d_1668_v4_ = D1_DC3(d_1664_v1_)
            d_1669_v5_: C3
            nw285_ = C3()
            nw285_.ctor__(d_1668_v4_)
            d_1669_v5_ = nw285_
            (globalState).f1 = d_1664_v1_
            d_1670_v6_: _dafny.Seq
            d_1670_v6_ = _dafny.SeqWithoutIsStrInference([d_1664_v1_])
            d_1671_v7_: _dafny.Seq
            d_1671_v7_ = _dafny.SeqWithoutIsStrInference([(d_1670_v6_) + (d_1670_v6_)])
            d_1671_v7_ = d_1671_v7_
            if d_1664_v1_:
                d_1672_v8_: _dafny.Array
                def lambda65_(d_1673_v1_):
                    def lambda66_(d_1674_i1_):
                        return default__.safeDivisionInt(d_1674_i1_, len(_dafny.Map({d_1673_v1_: d_1673_v1_})))

                    return lambda66_

                init36_ = lambda65_(d_1664_v1_)
                nw286_ = _dafny.Array(None, 20)
                for i0_36_ in range(nw286_.length(0)):
                    nw286_[i0_36_] = init36_(i0_36_)
                d_1672_v8_ = nw286_
                rhs301_ = (d_1672_v8_ if d_1664_v1_ else d_1672_v8_)
                rhs302_ = d_1664_v1_
                d_1672_v8_ = rhs301_
                d_1664_v1_ = rhs302_
                d_1675_v9_: _dafny.Array
                def lambda67_(d_1676_v1_, d_1677_v3_):
                    def lambda68_(d_1678_i2_):
                        return D4_DC12(d_1676_v1_, _dafny.Set({d_1677_v3_, d_1677_v3_}))

                    return lambda68_

                init37_ = lambda67_(d_1664_v1_, d_1666_v3_)
                nw287_ = _dafny.Array(None, 2)
                for i0_37_ in range(nw287_.length(0)):
                    nw287_[i0_37_] = init37_(i0_37_)
                d_1675_v9_ = nw287_
                d_1679_v10_: _dafny.Set
                d_1679_v10_ = _dafny.Set({d_1666_v3_})
                d_1680_v11_: D4
                d_1680_v11_ = D4_DC12(d_1664_v1_, d_1679_v10_)
                index288_ = default__.safeIndex(326, (d_1675_v9_).length(0))
                (d_1675_v9_)[index288_] = d_1680_v11_
                index289_ = default__.safeIndex(326, (d_1675_v9_).length(0))
                (d_1675_v9_)[index289_] = d_1680_v11_
                (globalState).f16 = ((_dafny.SeqWithoutIsStrInference([d_1664_v1_, d_1664_v1_, True])) + (d_1670_v6_)).set(default__.safeIndex((d_1667_i0_ if d_1664_v1_ else default__.fm2(d_1663_v0_, d_1665_v2_, _dafny.Map({d_1664_v1_: d_1666_v3_}), True, globalState)), len((_dafny.SeqWithoutIsStrInference([d_1664_v1_, d_1664_v1_, True])) + (d_1670_v6_))), True)
                d_1681_v12_: D4
                d_1681_v12_ = D4_DC10(d_1666_v3_)
                d_1682_v13_: _dafny.Map
                d_1682_v13_ = _dafny.Map({(d_1681_v12_).cf19: d_1665_v2_})
                d_1682_v13_ = (d_1682_v13_).set(d_1666_v3_, d_1665_v2_)
                d_1683_v14_: _dafny.Seq
                d_1683_v14_ = _dafny.SeqWithoutIsStrInference([d_1667_i0_, d_1667_i0_])
                d_1684_v15_: _dafny.Set
                d_1684_v15_ = _dafny.Set({d_1667_i0_})
                d_1685_v16_: _dafny.Map
                d_1685_v16_ = _dafny.Map({d_1664_v1_: d_1684_v15_})
                d_1686_v17_: _dafny.Map
                d_1686_v17_ = _dafny.Map({d_1672_v8_: len(((d_1685_v16_)[d_1664_v1_] if (d_1664_v1_) in (d_1685_v16_) else d_1684_v15_))})
                d_1687_v18_: _dafny.MultiSet
                d_1687_v18_ = _dafny.MultiSet([(d_1686_v17_) | (d_1686_v17_), d_1686_v17_, (d_1686_v17_) | (d_1686_v17_), d_1686_v17_, _dafny.Map({d_1672_v8_: len(d_1683_v14_)})])
                d_1688_v19_: _dafny.Map
                d_1688_v19_ = _dafny.Map({(self).f34: _dafny.Map({d_1672_v8_: (self).f34})})
                d_1689_v20_: D19
                d_1689_v20_ = D19_DC44(((d_1688_v19_)[d_1667_i0_] if (d_1667_i0_) in (d_1688_v19_) else _dafny.Map({d_1672_v8_: (self).f34})))
                rhs303_ = ((_dafny.SeqWithoutIsStrInference([-984, d_1667_i0_])) == (d_1683_v14_) if True else not(d_1664_v1_))
                rhs304_ = ((d_1687_v18_)[(d_1686_v17_) | ((d_1689_v20_).cf81)] if ((d_1686_v17_) | ((d_1689_v20_).cf81)) in (d_1687_v18_) else 288)
                lhs255_ = globalState
                lhs256_ = globalState
                lhs255_.f1 = rhs303_
                lhs256_.f24 = rhs304_
            elif True:
                (globalState).f0 = len(d_1663_v0_)
                (globalState).f27 = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_1690_i3_ in range(default__.abs(947))])) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q")))
                d_1691_v21_: _dafny.Seq
                d_1691_v21_ = _dafny.SeqWithoutIsStrInference([(self).f34, d_1667_i0_])
                (globalState).f0 = default__.safeDivisionInt((d_1691_v21_)[default__.safeIndex(-853, len(d_1691_v21_))], (self).f34)
                (globalState).f1 = ((self).f34) > (d_1667_i0_)
                d_1692_v22_: _dafny.Array
                nw288_ = _dafny.Array(False, 28)
                d_1692_v22_ = nw288_
                index290_ = default__.safeIndex(761, (d_1692_v22_).length(0))
                (d_1692_v22_)[index290_] = (self).fm16(d_1665_v2_, d_1667_i0_, globalState)
                d_1693_v23_: _dafny.MultiSet
                d_1693_v23_ = _dafny.MultiSet([(d_1692_v22_ if d_1664_v1_ else d_1692_v22_), d_1692_v22_, d_1692_v22_])
                d_1694_v24_: _dafny.Map
                d_1694_v24_ = _dafny.Map({d_1664_v1_: d_1666_v3_})
                index291_ = default__.safeIndex(761, (d_1692_v22_).length(0))
                rhs305_ = default__.fm2(d_1663_v0_, d_1665_v2_, d_1694_v24_, d_1664_v1_, globalState)
                rhs306_ = d_1664_v1_
                rhs307_ = not(not((d_1667_i0_) > ((self).f34)))
                rhs308_ = (((self).f34) - (-58)) != (default__.safeDivisionInt((self).f34, len(d_1691_v21_)))
                rhs309_ = d_1693_v23_
                lhs257_ = globalState
                lhs258_ = globalState
                lhs259_ = globalState
                lhs260_ = d_1692_v22_
                lhs261_ = default__.safeIndex(761, (d_1692_v22_).length(0))
                lhs257_.f24 = rhs305_
                lhs258_.f11 = rhs306_
                lhs259_.f1 = rhs307_
                lhs260_[lhs261_] = rhs308_
                d_1693_v23_ = rhs309_
        d_1695_v25_: _dafny.Array
        nw289_ = _dafny.Array(None, 14)
        nw289_[int(0)] = d_1664_v1_
        nw289_[int(1)] = d_1664_v1_
        nw289_[int(2)] = False
        nw289_[int(3)] = not(True)
        nw289_[int(4)] = d_1664_v1_
        nw289_[int(5)] = (len(d_1663_v0_)) != ((self).f34)
        nw289_[int(6)] = d_1664_v1_
        nw289_[int(7)] = False
        nw289_[int(8)] = d_1664_v1_
        nw289_[int(9)] = False
        nw289_[int(10)] = d_1664_v1_
        nw289_[int(11)] = d_1664_v1_
        nw289_[int(12)] = d_1664_v1_
        nw289_[int(13)] = d_1664_v1_
        d_1695_v25_ = nw289_
        guard_loop_3_: int
        for guard_loop_3_ in _dafny.IntegerRange(0, (d_1695_v25_).length(0)):
            d_1696_i4_: int = guard_loop_3_
            if (True) and (((0) <= (d_1696_i4_)) and ((d_1696_i4_) < ((d_1695_v25_).length(0)))):
                (d_1695_v25_)[(d_1696_i4_)] = (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f34]))).issubset(_dafny.MultiSet([len(_dafny.Set({_dafny.Set({(self).f34})})), len(_dafny.SeqWithoutIsStrInference([(self).f34, (self).f34]))]))
        d_1697_v26_: _dafny.Array
        nw290_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_1697_v26_ = nw290_
        d_1698_v27_: C4
        nw291_ = C4()
        nw291_.ctor__(d_1664_v1_, d_1697_v26_)
        d_1698_v27_ = nw291_
        index292_ = default__.safeIndex(783, (d_1695_v25_).length(0))
        (d_1695_v25_)[index292_] = (d_1698_v27_).f37
        d_1699_v28_: _dafny.Seq
        d_1699_v28_ = _dafny.SeqWithoutIsStrInference([not(d_1664_v1_)])
        index293_ = default__.safeIndex(783, (d_1695_v25_).length(0))
        (d_1695_v25_)[index293_] = ((d_1698_v27_).f37) not in (d_1699_v28_)
        d_1663_v0_ = ((d_1663_v0_) + ((d_1663_v0_).set(default__.safeIndex((self).f34, len(d_1663_v0_)), d_1666_v3_)) if (False if d_1664_v1_ else True) else d_1663_v0_)
        pat_let_tv37_ = d_1695_v25_
        pat_let_tv38_ = d_1695_v25_
        def iife166_(_pat_let47_0):
            def iife167_(d_1700_dt__update__tmp_h0_):
                def iife168_(_pat_let48_0):
                    def iife169_(d_1701_dt__update_hcf50_h0_):
                        return D12_DC29(d_1701_dt__update_hcf50_h0_)
                    return iife169_(_pat_let48_0)
                return iife168_((pat_let_tv38_)[default__.safeIndex(783, (pat_let_tv37_).length(0))])
            return iife167_(_pat_let47_0)
        source23_ = iife166_(D12_DC29((d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))]))
        if source23_.is_DC28:
            d_1702___mcc_h0_ = source23_.cf45
            d_1703___mcc_h1_ = source23_.cf46
            d_1704___mcc_h2_ = source23_.cf47
            d_1705___mcc_h3_ = source23_.cf48
            d_1706___mcc_h4_ = source23_.cf49
            d_1707_cf49_ = d_1706___mcc_h4_
            d_1708_cf48_ = d_1705___mcc_h3_
            d_1709_cf47_ = d_1704___mcc_h2_
            d_1710_cf46_ = d_1703___mcc_h1_
            d_1711_cf45_ = d_1702___mcc_h0_
            (globalState).f1 = False
            (globalState).f1 = (self).fm16(d_1665_v2_, ((0) - (d_1708_cf48_)) * (280), globalState)
            (globalState).f19 = d_1708_cf48_
            d_1664_v1_ = d_1709_cf47_
        elif source23_.is_DC29:
            d_1712___mcc_h5_ = source23_.cf50
            d_1713_cf50_ = d_1712___mcc_h5_
            d_1695_v25_ = d_1695_v25_
            d_1714_v29_: C3
            nw292_ = C3()
            nw292_.ctor__(D1_DC3((d_1698_v27_).f37))
            d_1714_v29_ = nw292_
            d_1715_v30_: _dafny.Seq
            d_1715_v30_ = _dafny.SeqWithoutIsStrInference([d_1665_v2_, d_1665_v2_])
            d_1716_v31_: C0
            nw293_ = C0()
            nw293_.ctor__((d_1715_v30_)[default__.safeIndex((self).f34, len(d_1715_v30_))])
            d_1716_v31_ = nw293_
            d_1717_v32_: _dafny.Map
            d_1717_v32_ = _dafny.Map({d_1664_v1_: d_1666_v3_})
            (globalState).f0 = default__.fm2(_dafny.SeqWithoutIsStrInference([d_1666_v3_ for d_1718_i5_ in range(default__.abs(94))]), d_1665_v2_, d_1717_v32_, (d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))], globalState)
        elif source23_.is_DC30:
            d_1719___mcc_h6_ = source23_.cf51
            d_1720___mcc_h7_ = source23_.cf52
            d_1721___mcc_h8_ = source23_.cf53
            d_1722___mcc_h9_ = source23_.cf54
            d_1723___mcc_h10_ = source23_.cf55
            d_1724_cf55_ = d_1723___mcc_h10_
            d_1725_cf54_ = d_1722___mcc_h9_
            d_1726_cf53_ = d_1721___mcc_h8_
            d_1727_cf52_ = d_1720___mcc_h7_
            d_1728_cf51_ = d_1719___mcc_h6_
            d_1729_v33_: C1
            nw294_ = C1()
            nw294_.ctor__(default__.safeDivisionInt((self).f34, len(d_1699_v28_)), d_1725_cf54_, D1_DC3(not((d_1699_v28_)[default__.safeIndex((self).f34, len(d_1699_v28_))])))
            d_1729_v33_ = nw294_
            d_1730_v34_: D1
            d_1730_v34_ = D1_DC3((d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))])
            source24_ = d_1730_v34_
            if source24_.is_DC4:
                d_1731___mcc_h12_ = source24_.cf5
                d_1732___mcc_h13_ = source24_.cf6
                d_1733_cf6_ = d_1732___mcc_h13_
                d_1734_cf5_ = d_1731___mcc_h12_
                d_1735_v35_: _dafny.Array
                nw295_ = _dafny.Array(int(0), 15)
                d_1735_v35_ = nw295_
                d_1736_v36_: _dafny.Seq
                d_1736_v36_ = _dafny.SeqWithoutIsStrInference([d_1663_v0_])
                d_1737_v37_: _dafny.Map
                d_1737_v37_ = _dafny.Map({d_1664_v1_: d_1666_v3_})
                d_1738_v38_: _dafny.Map
                d_1738_v38_ = _dafny.Map({d_1728_cf51_: d_1727_cf52_})
                index294_ = default__.safeIndex(783, (d_1695_v25_).length(0))
                rhs310_ = (_dafny.SeqWithoutIsStrInference([d_1666_v3_ for d_1739_i6_ in range(default__.abs(482))])) != (d_1663_v0_)
                rhs311_ = d_1695_v25_
                rhs312_ = default__.fm2((d_1736_v36_)[default__.safeIndex((d_1729_v33_).f41, len(d_1736_v36_))], d_1665_v2_, d_1737_v37_, not (((d_1738_v38_)[d_1726_cf53_] if (d_1726_cf53_) in (d_1738_v38_) else d_1727_cf52_)) or ((d_1698_v27_).f37), globalState)
                rhs313_ = d_1735_v35_
                lhs262_ = d_1695_v25_
                lhs263_ = default__.safeIndex(783, (d_1695_v25_).length(0))
                lhs264_ = globalState
                lhs262_[lhs263_] = rhs310_
                d_1695_v25_ = rhs311_
                lhs264_.f24 = rhs312_
                d_1735_v35_ = rhs313_
                d_1740_v39_: _dafny.Seq
                d_1740_v39_ = _dafny.SeqWithoutIsStrInference([(d_1734_cf5_) * (d_1725_cf54_), default__.safeDivisionInt((d_1729_v33_).f41, d_1729_v33_.f42), d_1734_cf5_, (default__.fm2(_dafny.SeqWithoutIsStrInference([d_1666_v3_ for d_1741_i7_ in range(default__.abs(892))]), _dafny.Set({False}), d_1737_v37_, (d_1698_v27_).f37, globalState)) - (len(d_1663_v0_))])
                index295_ = default__.safeIndex(959, (d_1735_v35_).length(0))
                (d_1735_v35_)[index295_] = d_1729_v33_.f42
                d_1742_v40_: _dafny.Map
                d_1742_v40_ = _dafny.Map({d_1728_cf51_: d_1740_v39_})
                index296_ = default__.safeIndex(959, (d_1735_v35_).length(0))
                index297_ = default__.safeIndex(783, (d_1695_v25_).length(0))
                rhs314_ = default__.fm6(default__.fm4(globalState), 37, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_1743_i8_ in range(default__.abs(699))]), globalState)
                rhs315_ = (default__.safeModuloInt((self).f34, d_1725_cf54_)) * ((d_1729_v33_).f41)
                rhs316_ = len(d_1742_v40_)
                rhs317_ = d_1663_v0_
                rhs318_ = (d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))]
                lhs265_ = globalState
                lhs266_ = d_1735_v35_
                lhs267_ = default__.safeIndex(959, (d_1735_v35_).length(0))
                lhs268_ = d_1695_v25_
                lhs269_ = default__.safeIndex(783, (d_1695_v25_).length(0))
                d_1740_v39_ = rhs314_
                lhs265_.f19 = rhs315_
                lhs266_[lhs267_] = rhs316_
                d_1663_v0_ = rhs317_
                lhs268_[lhs269_] = rhs318_
                d_1744_v41_: _dafny.MultiSet
                d_1744_v41_ = _dafny.MultiSet([(self).f34])
                d_1745_v42_: _dafny.Map
                d_1745_v42_ = _dafny.Map({(d_1729_v33_).f41: ((D20_DC48(d_1744_v41_)).cf89).cardinality})
                (globalState).f0 = default__.safeModuloInt(default__.safeModuloInt(len(d_1745_v42_), d_1729_v33_.f42), len(default__.fm3(globalState)))
                index298_ = default__.safeIndex(959, (d_1735_v35_).length(0))
                (d_1735_v35_)[index298_] = (d_1735_v35_)[default__.safeIndex(959, (d_1735_v35_).length(0))]
            elif source24_.is_DC5:
                d_1746___mcc_h14_ = source24_.cf7
                d_1747___mcc_h15_ = source24_.cf8
                d_1748___mcc_h16_ = source24_.cf9
                d_1749___mcc_h17_ = source24_.cf10
                d_1750___mcc_h18_ = source24_.cf11
                d_1751_cf11_ = d_1750___mcc_h18_
                d_1752_cf10_ = d_1749___mcc_h17_
                d_1753_cf9_ = d_1748___mcc_h16_
                d_1754_cf8_ = d_1747___mcc_h15_
                d_1755_cf7_ = d_1746___mcc_h14_
                d_1756_v43_: _dafny.MultiSet
                d_1756_v43_ = _dafny.MultiSet([905])
                d_1757_v44_: _dafny.Map
                d_1757_v44_ = _dafny.Map({d_1724_cf55_: (d_1756_v43_).cardinality})
                d_1758_v45_: _dafny.Seq
                d_1758_v45_ = _dafny.SeqWithoutIsStrInference([d_1757_v44_, d_1757_v44_, d_1757_v44_])
                d_1759_v46_: _dafny.Map
                d_1759_v46_ = _dafny.Map({d_1727_cf52_: d_1758_v45_})
                d_1759_v46_ = (d_1759_v46_).set(d_1664_v1_, d_1758_v45_)
                d_1760_v47_: _dafny.Array
                nw296_ = _dafny.Array(int(0), 16)
                d_1760_v47_ = nw296_
                index299_ = default__.safeIndex(509, (d_1760_v47_).length(0))
                (d_1760_v47_)[index299_] = default__.safeDivisionInt((d_1729_v33_).f41, 879)
                d_1761_v49_: _dafny.Set
                d_1761_v49_ = _dafny.Set({d_1725_cf54_})
                index300_ = default__.safeIndex(509, (d_1760_v47_).length(0))
                def iife170_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in (d_1761_v49_).Elements:
                        d_1762_v48_: int = compr_72_
                        if (d_1762_v48_) in (d_1761_v49_):
                            coll72_[default__.safeDivisionInt(d_1762_v48_, len(d_1663_v0_))] = d_1755_cf7_
                    return _dafny.Map(coll72_)
                rhs319_ = (d_1757_v44_) | (d_1757_v44_)
                rhs320_ = (len(iife170_()
                )) + (len(d_1699_v28_))
                rhs321_ = 986
                lhs270_ = globalState
                lhs271_ = d_1760_v47_
                lhs272_ = default__.safeIndex(509, (d_1760_v47_).length(0))
                d_1757_v44_ = rhs319_
                lhs270_.f0 = rhs320_
                lhs271_[lhs272_] = rhs321_
                d_1763_v50_: C3
                nw297_ = C3()
                nw297_.ctor__(d_1730_v34_)
                d_1763_v50_ = nw297_
                d_1764_v51_: _dafny.Set
                d_1764_v51_ = _dafny.Set({d_1763_v50_, d_1763_v50_, d_1763_v50_, d_1763_v50_, d_1763_v50_})
                rhs322_ = (len(d_1663_v0_)) <= (default__.safeDivisionInt(d_1725_cf54_, d_1725_cf54_))
                rhs323_ = not ((d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))]) or ((d_1727_cf52_) or (d_1728_cf51_))
                rhs324_ = (d_1666_v3_ if d_1728_cf51_ else (d_1666_v3_ if True else d_1666_v3_))
                rhs325_ = (default__.fm2(default__.fm3(globalState), _dafny.Set({d_1664_v1_}), _dafny.Map({d_1728_cf51_: d_1666_v3_}), (d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))], globalState) if d_1726_cf53_ else d_1729_v33_.f42)
                rhs326_ = d_1764_v51_
                lhs273_ = globalState
                d_1752_cf10_ = rhs322_
                d_1727_cf52_ = rhs323_
                d_1666_v3_ = rhs324_
                lhs273_.f19 = rhs325_
                d_1764_v51_ = rhs326_
                (d_1729_v33_).f42 = d_1729_v33_.f42
            elif True:
                d_1765___mcc_h19_ = source24_.cf4
                d_1766_cf4_ = d_1765___mcc_h19_
                (globalState).f11 = (d_1727_cf52_) or (not((d_1698_v27_).f37))
                d_1767_v52_: _dafny.Array
                nw298_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1767_v52_ = nw298_
                d_1767_v52_ = d_1767_v52_
                d_1768_v53_: _dafny.Array
                nw299_ = _dafny.Array(int(0), 17)
                d_1768_v53_ = nw299_
                d_1769_v54_: _dafny.Set
                d_1769_v54_ = _dafny.Set({d_1666_v3_, d_1666_v3_, d_1666_v3_, d_1666_v3_, d_1666_v3_})
                index301_ = default__.safeIndex(803, (d_1768_v53_).length(0))
                (d_1768_v53_)[index301_] = (len(d_1769_v54_) if (d_1698_v27_).f37 else d_1729_v33_.f42)
                index302_ = default__.safeIndex(803, (d_1768_v53_).length(0))
                rhs327_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dgm"))) + (d_1663_v0_)) == ((d_1663_v0_) + (d_1663_v0_))
                rhs328_ = (self).f34
                rhs329_ = d_1729_v33_.f42
                rhs330_ = d_1766_cf4_
                rhs331_ = (d_1666_v3_ if d_1664_v1_ else d_1666_v3_)
                lhs274_ = d_1768_v53_
                lhs275_ = default__.safeIndex(803, (d_1768_v53_).length(0))
                lhs276_ = globalState
                lhs277_ = globalState
                d_1728_cf51_ = rhs327_
                lhs274_[lhs275_] = rhs328_
                lhs276_.f24 = rhs329_
                lhs277_.f27 = rhs330_
                d_1666_v3_ = rhs331_
                index303_ = default__.safeIndex(803, (d_1768_v53_).length(0))
                (d_1768_v53_)[index303_] = -56
            (globalState).f0 = (d_1729_v33_).f41
            (globalState).f11 = not ((d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))]) or ((d_1698_v27_).f37)
        elif True:
            d_1770___mcc_h11_ = source23_.cf44
            d_1771_cf44_ = d_1770___mcc_h11_
            d_1772_v55_: _dafny.Array
            nw300_ = _dafny.Array(_dafny.Map({}), 27)
            d_1772_v55_ = nw300_
            d_1773_v57_: _dafny.Map
            def iife171_():
                coll73_ = _dafny.Map()
                compr_73_: str
                for compr_73_ in (_dafny.SeqWithoutIsStrInference([d_1666_v3_])).Elements:
                    d_1774_v56_: str = compr_73_
                    if (d_1774_v56_) in (_dafny.SeqWithoutIsStrInference([d_1666_v3_])):
                        coll73_[d_1774_v56_] = D16_DC37(d_1699_v28_, d_1663_v0_, d_1666_v3_)
                return _dafny.Map(coll73_)
            d_1773_v57_ = _dafny.Map({(d_1698_v27_).f37: iife171_()
            })
            d_1775_v58_: D16
            d_1775_v58_ = D16_DC37(d_1699_v28_, _dafny.SeqWithoutIsStrInference([d_1666_v3_ for d_1776_i9_ in range(default__.abs(792))]), d_1666_v3_)
            d_1777_v59_: _dafny.Map
            d_1777_v59_ = _dafny.Map({d_1666_v3_: d_1775_v58_})
            index304_ = default__.safeIndex(529, (d_1772_v55_).length(0))
            (d_1772_v55_)[index304_] = (d_1773_v57_).set(d_1664_v1_, d_1777_v59_)
            d_1778_v60_: _dafny.Seq
            d_1778_v60_ = _dafny.SeqWithoutIsStrInference([d_1663_v0_, _dafny.SeqWithoutIsStrInference([d_1666_v3_, _dafny.CodePoint('p')])])
            d_1779_v62_: _dafny.Map
            d_1779_v62_ = _dafny.Map({d_1666_v3_: (d_1698_v27_).f37})
            index305_ = default__.safeIndex(529, (d_1772_v55_).length(0))
            def iife172_():
                coll74_ = _dafny.Map()
                compr_74_: str
                for compr_74_ in (d_1779_v62_).keys.Elements:
                    d_1780_v61_: str = compr_74_
                    if (d_1780_v61_) in (d_1779_v62_):
                        coll74_[d_1780_v61_] = d_1775_v58_
                return _dafny.Map(coll74_)
            rhs332_ = (d_1773_v57_) | ((d_1773_v57_).set(False, iife172_()
            ))
            rhs333_ = (d_1663_v0_)[default__.safeIndex(default__.safeDivisionInt((self).f34, (self).f34), len(d_1663_v0_))]
            rhs334_ = (self).f34
            rhs335_ = (d_1778_v60_) + (d_1778_v60_)
            lhs278_ = d_1772_v55_
            lhs279_ = default__.safeIndex(529, (d_1772_v55_).length(0))
            lhs280_ = globalState
            lhs278_[lhs279_] = rhs332_
            d_1666_v3_ = rhs333_
            lhs280_.f24 = rhs334_
            d_1778_v60_ = rhs335_
            d_1781_v63_: _dafny.Seq
            d_1781_v63_ = _dafny.SeqWithoutIsStrInference([(self).f34])
            (globalState).f24 = (d_1781_v63_)[default__.safeIndex((self).f34, len(d_1781_v63_))]
            (globalState).f19 = (0) - (((0) - ((self).f34)) * ((self).f34))
            d_1782_v64_: _dafny.Map
            d_1782_v64_ = _dafny.Map({(self).f34: (d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))]})
            d_1782_v64_ = (d_1782_v64_).set((d_1781_v63_)[default__.safeIndex(378, len(d_1781_v63_))], (d_1699_v28_)[default__.safeIndex((0) - (-696), len(d_1699_v28_))])
        d_1783_v65_: _dafny.Map
        d_1783_v65_ = _dafny.Map({(d_1698_v27_).f37: _dafny.CodePoint('x')})
        r0 = (default__.fm2(_dafny.SeqWithoutIsStrInference([d_1666_v3_ for d_1784_i10_ in range(default__.abs(494))]), d_1665_v2_, d_1783_v65_, (d_1695_v25_)[default__.safeIndex(783, (d_1695_v25_).length(0))], globalState)) != ((self).f34)
        return r0

    def m10(self, p0, p1, p2, globalState):
        d_1785_v0_: _dafny.Seq
        d_1785_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xqqvwtki"))
        d_1786_v1_: _dafny.Set
        d_1786_v1_ = _dafny.Set({p2})
        d_1787_v2_: str
        d_1787_v2_ = _dafny.CodePoint('p')
        d_1788_v3_: _dafny.Map
        d_1788_v3_ = _dafny.Map({p2: d_1787_v2_})
        d_1789_v4_: _dafny.Array
        nw301_ = _dafny.Array(int(0), 27)
        d_1789_v4_ = nw301_
        d_1790_v5_: _dafny.MultiSet
        d_1790_v5_ = _dafny.MultiSet([d_1789_v4_])
        d_1791_v6_: _dafny.Map
        d_1791_v6_ = _dafny.Map({p1: p2})
        d_1792_v7_: _dafny.MultiSet
        d_1792_v7_ = _dafny.MultiSet([d_1787_v2_])
        d_1793_v8_: _dafny.Map
        d_1793_v8_ = _dafny.Map({p2: d_1786_v1_})
        d_1794_v9_: _dafny.Array
        def lambda69_(d_1795_p2_):
            def lambda70_(d_1796_i0_):
                return d_1795_p2_

            return lambda70_

        init38_ = lambda69_(p2)
        nw302_ = _dafny.Array(None, 11)
        for i0_38_ in range(nw302_.length(0)):
            nw302_[i0_38_] = init38_(i0_38_)
        d_1794_v9_ = nw302_
        d_1797_v10_: _dafny.MultiSet
        d_1797_v10_ = _dafny.MultiSet([d_1794_v9_, d_1794_v9_, d_1794_v9_])
        d_1798_v11_: _dafny.MultiSet
        d_1798_v11_ = _dafny.MultiSet([(d_1785_v0_).set(default__.safeIndex((self).f34, len(d_1785_v0_)), d_1787_v2_)])
        d_1799_v12_: _dafny.Seq
        d_1799_v12_ = _dafny.SeqWithoutIsStrInference([(self).f34, p1, p1])
        d_1800_v13_: _dafny.Array
        nw303_ = _dafny.Array(None, 29)
        nw303_[int(0)] = default__.fm2(d_1785_v0_, d_1786_v1_, d_1788_v3_, p2, globalState)
        nw303_[int(1)] = p1
        nw303_[int(2)] = (self).f34
        nw303_[int(3)] = ((d_1790_v5_)[d_1789_v4_] if (d_1789_v4_) in (d_1790_v5_) else (self).f34)
        nw303_[int(4)] = (self).f34
        nw303_[int(5)] = p1
        nw303_[int(6)] = default__.safeModuloInt((self).f34, (self).f34)
        nw303_[int(7)] = (210 if p2 else -671)
        nw303_[int(8)] = len((d_1791_v6_) | (d_1791_v6_))
        nw303_[int(9)] = ((self).f34) + (p1)
        nw303_[int(10)] = (self).f34
        nw303_[int(11)] = len(_dafny.SeqWithoutIsStrInference([d_1789_v4_, d_1789_v4_]))
        nw303_[int(12)] = (self).f34
        nw303_[int(13)] = ((d_1792_v7_).set(d_1787_v2_, default__.abs((self).f34))).cardinality
        nw303_[int(14)] = (0) - (default__.fm2(d_1785_v0_, ((d_1793_v8_)[p2] if (p2) in (d_1793_v8_) else d_1786_v1_), d_1788_v3_, p2, globalState))
        nw303_[int(15)] = (self).f34
        nw303_[int(16)] = len(default__.fm3(globalState))
        nw303_[int(17)] = 60
        nw303_[int(18)] = p1
        nw303_[int(19)] = p1
        nw303_[int(20)] = (0) - (((d_1797_v10_)[d_1794_v9_] if (d_1794_v9_) in (d_1797_v10_) else len((d_1785_v0_).set(default__.safeIndex(p1, len(d_1785_v0_)), d_1787_v2_))))
        nw303_[int(21)] = p1
        nw303_[int(22)] = (self).f34
        nw303_[int(23)] = (d_1798_v11_).cardinality
        nw303_[int(24)] = 217
        nw303_[int(25)] = len(default__.fm3(globalState))
        nw303_[int(26)] = ((_dafny.MultiSet([p1, p1]) if p2 else _dafny.MultiSet(d_1799_v12_))).cardinality
        nw303_[int(27)] = -518
        nw303_[int(28)] = ((0) - ((self).f34)) - ((self).f34)
        d_1800_v13_ = nw303_
        index306_ = default__.safeIndex(274, (d_1800_v13_).length(0))
        def iife173_():
            coll75_ = _dafny.Set()
            compr_75_: int
            for compr_75_ in _dafny.IntegerRange(762, -763):
                d_1801_v14_: int = compr_75_
                if ((762) <= (d_1801_v14_)) and ((d_1801_v14_) < (-763)):
                    coll75_ = coll75_.union(_dafny.Set([(d_1801_v14_) - (p1)]))
            return _dafny.Set(coll75_)
        (d_1800_v13_)[index306_] = len((_dafny.Set({(self).f34, (self).f34})) | (iife173_()
        ))
        d_1802_v15_: _dafny.MultiSet
        d_1802_v15_ = _dafny.MultiSet([True, p2, p2])
        index307_ = default__.safeIndex(274, (d_1800_v13_).length(0))
        (d_1800_v13_)[index307_] = (((d_1802_v15_)[p2] if (p2) in (d_1802_v15_) else (self).f34)) + ((self).f34)
        if (default__.safeModuloInt((d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))], (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))])) > (default__.safeDivisionInt(103, p1)):
            d_1803_v16_: D6
            d_1803_v16_ = D6_DC16(p2, p2)
            d_1804_v17_: _dafny.Map
            d_1804_v17_ = _dafny.Map({_dafny.CodePoint('e'): p2})
            d_1805_v18_: _dafny.Map
            d_1805_v18_ = _dafny.Map({d_1804_v17_: not(p2)})
            d_1806_v19_: _dafny.Map
            d_1806_v19_ = _dafny.Map({(self).f34: (_dafny.MultiSet(d_1799_v12_)).cardinality})
            d_1807_v20_: _dafny.MultiSet
            d_1807_v20_ = _dafny.MultiSet([(0) - (((d_1806_v19_)[len(d_1785_v0_)] if (len(d_1785_v0_)) in (d_1806_v19_) else -937))])
            pat_let_tv39_ = globalState
            index308_ = default__.safeIndex(274, (d_1800_v13_).length(0))
            def iife174_(_pat_let49_0):
                def iife175_(d_1808_dt__update__tmp_h0_):
                    def iife176_(_pat_let50_0):
                        def iife177_(d_1809_dt__update_hcf27_h0_):
                            return D6_DC16((d_1808_dt__update__tmp_h0_).cf26, d_1809_dt__update_hcf27_h0_)
                        return iife177_(_pat_let50_0)
                    return iife176_(not(default__.fm4(pat_let_tv39_)))
                return iife175_(_pat_let49_0)
            rhs336_ = iife174_(d_1803_v16_)
            rhs337_ = ((_dafny.MultiSet([p1, (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]]) if False else d_1807_v20_)).cardinality
            rhs338_ = d_1791_v6_
            rhs339_ = _dafny.Map({(d_1804_v17_).set(d_1787_v2_, p2): (d_1785_v0_) <= (d_1785_v0_)})
            lhs281_ = d_1800_v13_
            lhs282_ = default__.safeIndex(274, (d_1800_v13_).length(0))
            lhs283_ = globalState
            d_1803_v16_ = rhs336_
            lhs281_[lhs282_] = rhs337_
            lhs283_.f20 = rhs338_
            d_1805_v18_ = rhs339_
            d_1810_v21_: D15
            d_1810_v21_ = D15_DC34(p2, len(_dafny.Map({d_1806_v19_: p2})))
            index309_ = default__.safeIndex(783, (d_1794_v9_).length(0))
            (d_1794_v9_)[index309_] = (d_1810_v21_).cf59
            index310_ = default__.safeIndex(783, (d_1794_v9_).length(0))
            (d_1794_v9_)[index310_] = p2
            d_1811_v22_: D1
            d_1811_v22_ = D1_DC3(p2)
            d_1812_v23_: T0
            nw304_ = C3()
            nw304_.ctor__(d_1811_v22_)
            d_1812_v23_ = nw304_
            d_1813_v24_: _dafny.Seq
            d_1813_v24_ = _dafny.SeqWithoutIsStrInference([d_1812_v23_])
            d_1814_v25_: _dafny.Seq
            d_1814_v25_ = _dafny.SeqWithoutIsStrInference([p2, p2, p2, p2, (d_1794_v9_)[default__.safeIndex(783, (d_1794_v9_).length(0))]])
            d_1815_v26_: _dafny.Map
            d_1815_v26_ = _dafny.Map({(0) - ((0) - (len(d_1786_v1_))): d_1787_v2_})
            d_1816_v27_: _dafny.Map
            d_1816_v27_ = _dafny.Map({len(d_1814_v25_): d_1815_v26_})
            d_1817_v28_: _dafny.Seq
            d_1817_v28_ = _dafny.SeqWithoutIsStrInference([(d_1813_v24_)[default__.safeIndex(len(((d_1816_v27_)[len(_dafny.SeqWithoutIsStrInference([p2]))] if (len(_dafny.SeqWithoutIsStrInference([p2]))) in (d_1816_v27_) else _dafny.Map({(d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]: _dafny.CodePoint('g')}))), len(d_1813_v24_))], d_1812_v23_, (d_1813_v24_)[default__.safeIndex(len(d_1785_v0_), len(d_1813_v24_))]])
            if (d_1817_v28_) not in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1812_v23_, d_1812_v23_]), d_1817_v28_])):
                d_1818_v29_: _dafny.Set
                d_1818_v29_ = _dafny.Set({d_1787_v2_, d_1787_v2_, d_1787_v2_, _dafny.CodePoint('r')})
                d_1819_v30_: D0
                d_1819_v30_ = D0_DC0(d_1806_v19_)
                d_1820_v31_: _dafny.Map
                d_1820_v31_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_1821_i1_ in range(default__.abs(-423))]): (d_1794_v9_)[default__.safeIndex(783, (d_1794_v9_).length(0))]})
                d_1818_v29_ = ((d_1818_v29_) | (_dafny.Set({(d_1812_v23_).fm8(d_1806_v19_, p2, p2, globalState)}))).intersection(_dafny.Set({(d_1812_v23_).fm8((d_1819_v30_).cf0, (d_1814_v25_)[default__.safeIndex((self).f34, len(d_1814_v25_))], ((d_1820_v31_)[d_1785_v0_] if (d_1785_v0_) in (d_1820_v31_) else p2), globalState), d_1787_v2_}))
                d_1822_v32_: _dafny.Array
                nw305_ = _dafny.Array(_dafny.Set({}), 2)
                d_1822_v32_ = nw305_
                index311_ = default__.safeIndex(793, (d_1822_v32_).length(0))
                def iife178_():
                    coll76_ = _dafny.Set()
                    compr_76_: int
                    for compr_76_ in _dafny.IntegerRange(256, 964):
                        d_1823_v33_: int = compr_76_
                        if ((256) <= (d_1823_v33_)) and ((d_1823_v33_) < (964)):
                            coll76_ = coll76_.union(_dafny.Set([default__.safeDivisionInt(d_1823_v33_, 282)]))
                    return _dafny.Set(coll76_)
                (d_1822_v32_)[index311_] = iife178_()
                
                d_1824_v34_: _dafny.Set
                d_1824_v34_ = _dafny.Set({(855 if p2 else p1), len(d_1785_v0_)})
                index312_ = default__.safeIndex(793, (d_1822_v32_).length(0))
                (d_1822_v32_)[index312_] = d_1824_v34_
                d_1787_v2_ = (d_1785_v0_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1787_v2_ for d_1825_i2_ in range(default__.abs(225))])), len(d_1785_v0_))]
                d_1826_v35_: _dafny.Array
                nw306_ = _dafny.Array(None, 22)
                nw306_[int(0)] = d_1787_v2_
                nw306_[int(1)] = d_1787_v2_
                nw306_[int(2)] = d_1787_v2_
                nw306_[int(3)] = d_1787_v2_
                nw306_[int(4)] = d_1787_v2_
                nw306_[int(5)] = d_1787_v2_
                nw306_[int(6)] = d_1787_v2_
                nw306_[int(7)] = d_1787_v2_
                nw306_[int(8)] = d_1787_v2_
                nw306_[int(9)] = d_1787_v2_
                nw306_[int(10)] = d_1787_v2_
                nw306_[int(11)] = d_1787_v2_
                nw306_[int(12)] = d_1787_v2_
                nw306_[int(13)] = (d_1812_v23_).fm8(d_1806_v19_, not(p2), p2, globalState)
                nw306_[int(14)] = d_1787_v2_
                nw306_[int(15)] = d_1787_v2_
                nw306_[int(16)] = d_1787_v2_
                nw306_[int(17)] = d_1787_v2_
                nw306_[int(18)] = ((d_1815_v26_)[739] if (739) in (d_1815_v26_) else d_1787_v2_)
                nw306_[int(19)] = d_1787_v2_
                nw306_[int(20)] = d_1787_v2_
                nw306_[int(21)] = d_1787_v2_
                d_1826_v35_ = nw306_
                index313_ = default__.safeIndex(784, (d_1826_v35_).length(0))
                (d_1826_v35_)[index313_] = d_1787_v2_
                index314_ = default__.safeIndex(784, (d_1826_v35_).length(0))
                (d_1826_v35_)[index314_] = d_1787_v2_
                d_1827_v36_: _dafny.Array
                def lambda71_(d_1828_p1_, d_1829_v6_, d_1830_v25_, d_1831_v13_):
                    def lambda72_(d_1832_i3_):
                        return ((d_1829_v6_)[d_1828_p1_] if (d_1828_p1_) in (d_1829_v6_) else (d_1830_v25_)[default__.safeIndex((d_1831_v13_)[default__.safeIndex(274, (d_1831_v13_).length(0))], len(d_1830_v25_))])

                    return lambda72_

                init39_ = lambda71_(p1, d_1791_v6_, d_1814_v25_, d_1800_v13_)
                nw307_ = _dafny.Array(None, 26)
                for i0_39_ in range(nw307_.length(0)):
                    nw307_[i0_39_] = init39_(i0_39_)
                d_1827_v36_ = nw307_
                d_1827_v36_ = d_1827_v36_
            elif True:
                d_1833_v37_: D12
                d_1833_v37_ = D12_DC30(p2, (d_1794_v9_)[default__.safeIndex(783, (d_1794_v9_).length(0))], p2, (self).f34, (d_1794_v9_)[default__.safeIndex(783, (d_1794_v9_).length(0))])
                index315_ = default__.safeIndex(783, (d_1794_v9_).length(0))
                (d_1794_v9_)[index315_] = (d_1833_v37_).cf51
                rhs340_ = True
                rhs341_ = (len((d_1785_v0_ if (d_1794_v9_)[default__.safeIndex(783, (d_1794_v9_).length(0))] else d_1785_v0_))) + (p1)
                lhs284_ = globalState
                lhs285_ = globalState
                lhs284_.f11 = rhs340_
                lhs285_.f0 = rhs341_
                (globalState).f27 = p2
                d_1834_v38_: _dafny.Map
                d_1834_v38_ = _dafny.Map({(d_1794_v9_)[default__.safeIndex(783, (d_1794_v9_).length(0))]: (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]})
                d_1834_v38_ = (d_1834_v38_).set(p2, 226)
                (globalState).f1 = not(p2)
            d_1835_v39_: _dafny.Map
            d_1835_v39_ = _dafny.Map({d_1814_v25_: (self).f34})
            d_1836_v40_: C2
            nw308_ = C2()
            nw308_.ctor__((d_1835_v39_ if not(p2) else d_1835_v39_), d_1786_v1_, (d_1812_v23_).f28)
            d_1836_v40_ = nw308_
            (globalState).f11 = (d_1785_v0_) < (d_1785_v0_)
        elif True:
            (globalState).f19 = ((d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]) + (p1)
            d_1837_v41_: _dafny.Map
            d_1837_v41_ = _dafny.Map({len(p0): p1})
            d_1838_v42_: _dafny.Map
            d_1838_v42_ = _dafny.Map({p2: _dafny.Map({p2: p2})})
            d_1839_v43_: _dafny.Map
            d_1839_v43_ = _dafny.Map({(d_1837_v41_).set(p1, len(d_1838_v42_)): (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]})
            rhs342_ = default__.safeDivisionInt(((d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]) - (len(d_1799_v12_)), p1)
            rhs343_ = (d_1839_v43_) == ((d_1839_v43_) | (d_1839_v43_))
            rhs344_ = d_1787_v2_
            rhs345_ = p2
            rhs346_ = p2
            lhs286_ = globalState
            lhs287_ = globalState
            lhs288_ = globalState
            lhs289_ = globalState
            lhs286_.f14 = rhs342_
            lhs287_.f11 = rhs343_
            d_1787_v2_ = rhs344_
            lhs288_.f1 = rhs345_
            lhs289_.f1 = rhs346_
            (globalState).f19 = (p1) - ((d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))])
            (globalState).f0 = (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))]
            d_1840_v44_: D1
            d_1840_v44_ = D1_DC3(not(p2))
            d_1841_v45_: C1
            nw309_ = C1()
            nw309_.ctor__((self).f34, ((0) - (len(d_1799_v12_))) * (len(default__.fm0(p2, globalState))), d_1840_v44_)
            d_1841_v45_ = nw309_
        d_1842_v46_: _dafny.Map
        d_1842_v46_ = _dafny.Map({d_1787_v2_: d_1794_v9_})
        d_1843_v47_: D12
        d_1843_v47_ = D12_DC27(d_1842_v46_)
        d_1843_v47_ = d_1843_v47_
        (globalState).f1 = p2
        d_1844_v48_: _dafny.Array
        nw310_ = _dafny.Array(_dafny.Array(None, 0), 4)
        d_1844_v48_ = nw310_
        index316_ = default__.safeIndex(596, (d_1844_v48_).length(0))
        (d_1844_v48_)[index316_] = d_1794_v9_
        index317_ = default__.safeIndex(596, (d_1844_v48_).length(0))
        (d_1844_v48_)[index317_] = d_1794_v9_
        if p2:
            d_1845_v49_: _dafny.Map
            d_1845_v49_ = _dafny.Map({p2: (self).f34})
            d_1846_v50_: D12
            d_1846_v50_ = D12_DC30(p2, not(p2), p2, (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))], p2)
            d_1847_v51_: D15
            d_1847_v51_ = D15_DC34(p2, (d_1846_v50_).cf54)
            rhs347_ = (self).f34
            rhs348_ = d_1845_v49_
            rhs349_ = default__.fm2(d_1785_v0_, d_1786_v1_, d_1788_v3_, p2, globalState)
            rhs350_ = default__.fm38(((self).f34) + (p1), p2, p1, (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))], globalState)
            lhs290_ = globalState
            lhs291_ = globalState
            lhs290_.f14 = rhs347_
            d_1845_v49_ = rhs348_
            lhs291_.f0 = rhs349_
            d_1847_v51_ = rhs350_
            default__.m0((d_1844_v48_)[default__.safeIndex(596, (d_1844_v48_).length(0))], d_1800_v13_, (self).f34, default__.fm2((d_1785_v0_).set(default__.safeIndex((self).f34, len(d_1785_v0_)), d_1787_v2_), d_1786_v1_, _dafny.Map({p2: d_1787_v2_}), p2, globalState), globalState)
            index318_ = default__.safeIndex(106, (d_1794_v9_).length(0))
            (d_1794_v9_)[index318_] = p2
            index319_ = default__.safeIndex(106, (d_1794_v9_).length(0))
            (d_1794_v9_)[index319_] = not(p2)
            d_1848_v52_: _dafny.Array
            nw311_ = _dafny.Array(_dafny.Seq({}), 5)
            d_1848_v52_ = nw311_
            d_1849_v53_: _dafny.Array
            d_1849_v53_ = d_1848_v52_
            d_1850_v54_: _dafny.Seq
            d_1850_v54_ = _dafny.SeqWithoutIsStrInference([True])
            d_1851_v55_: _dafny.Map
            d_1851_v55_ = _dafny.Map({(d_1850_v54_) + (d_1850_v54_): default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ntvemf")), d_1786_v1_, d_1788_v3_, p2, globalState)})
            index320_ = default__.safeIndex(596, (d_1844_v48_).length(0))
            rhs351_ = p2
            rhs352_ = d_1849_v53_
            rhs353_ = (d_1844_v48_)[default__.safeIndex(596, (d_1844_v48_).length(0))]
            rhs354_ = (d_1851_v55_) | ((default__.fm22(False, (d_1800_v13_)[default__.safeIndex(274, (d_1800_v13_).length(0))], 174, globalState)) | (_dafny.Map({d_1850_v54_: p1})))
            lhs292_ = globalState
            lhs293_ = d_1844_v48_
            lhs294_ = default__.safeIndex(596, (d_1844_v48_).length(0))
            lhs292_.f1 = rhs351_
            d_1849_v53_ = rhs352_
            lhs293_[lhs294_] = rhs353_
            d_1851_v55_ = rhs354_
            (globalState).f24 = default__.safeModuloInt(p1, default__.safeDivisionInt(p1, 854))
        elif True:
            index321_ = default__.safeIndex(274, (d_1800_v13_).length(0))
            (d_1800_v13_)[index321_] = default__.safeDivisionInt(default__.safeModuloInt((self).f34, len(d_1785_v0_)), p1)
            (globalState).f27 = p2
            index322_ = default__.safeIndex(596, (d_1844_v48_).length(0))
            (d_1844_v48_)[index322_] = (d_1844_v48_)[default__.safeIndex(596, (d_1844_v48_).length(0))]
            d_1789_v4_ = d_1789_v4_
            d_1852_v56_: bool
            out13_: bool
            out13_ = (self).m9(globalState)
            d_1852_v56_ = out13_

    @property
    def f34(self):
        return self._f34

class C7(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f28):
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        return False

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('g')

    def fm15(self, p0, p1, p2, globalState):
        return default__.safeModuloInt(811, (258 if True else len(_dafny.SeqWithoutIsStrInference([961, -652, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uescm")))]))))

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_1853_v0_: int
        d_1853_v0_ = 279
        d_1854_i0_: int
        d_1854_i0_ = 0
        with _dafny.label("14"):
            while ((0) - (d_1853_v0_)) <= (d_1853_v0_):
                with _dafny.c_label("14"):
                    if (d_1854_i0_) >= (100):
                        raise _dafny.Break("14")
                    d_1854_i0_ = (d_1854_i0_) + (1)
                    d_1855_v1_: bool
                    d_1855_v1_ = False
                    if not((self).fm7(d_1855_v1_, False, globalState)):
                        d_1856_v2_: C5
                        nw312_ = C5()
                        nw312_.ctor__(d_1853_v0_, d_1855_v1_, (self).f28)
                        d_1856_v2_ = nw312_
                        d_1857_v3_: _dafny.Seq
                        d_1857_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))
                        d_1858_v4_: str
                        d_1858_v4_ = _dafny.CodePoint('p')
                        d_1859_v5_: D16
                        d_1859_v5_ = D16_DC37(_dafny.SeqWithoutIsStrInference([(d_1856_v2_).f36]), d_1857_v3_, d_1858_v4_)
                        d_1860_v6_: _dafny.Map
                        d_1860_v6_ = _dafny.Map({d_1859_v5_: (d_1856_v2_).f36})
                        d_1861_v7_: _dafny.Set
                        d_1861_v7_ = _dafny.Set({(d_1856_v2_).f36, d_1855_v1_, not(((d_1860_v6_)[d_1859_v5_] if (d_1859_v5_) in (d_1860_v6_) else d_1855_v1_)), (self).fm7((d_1856_v2_).f36, (d_1856_v2_).f36, globalState), (d_1856_v2_).f36})
                        d_1862_v8_: _dafny.Seq
                        d_1862_v8_ = _dafny.SeqWithoutIsStrInference([d_1861_v7_, d_1861_v7_, d_1861_v7_])
                        d_1862_v8_ = d_1862_v8_
                        d_1863_v9_: _dafny.Set
                        d_1863_v9_ = _dafny.Set({(d_1856_v2_).f35})
                        d_1864_v10_: _dafny.Set
                        d_1864_v10_ = _dafny.Set({d_1853_v0_, (0) - (d_1853_v0_), len(d_1863_v9_)})
                        (globalState).f27 = (d_1855_v1_) and ((d_1864_v10_).isdisjoint(d_1863_v9_))
                        d_1865_v11_: _dafny.MultiSet
                        d_1865_v11_ = _dafny.MultiSet([len(_dafny.Set({(d_1856_v2_).f36, d_1855_v1_}))])
                        d_1866_v12_: _dafny.Array
                        nw313_ = _dafny.Array(None, 9)
                        nw313_[int(0)] = ((d_1865_v11_) | (d_1865_v11_)).cardinality
                        nw313_[int(1)] = ((0) - ((d_1856_v2_).f35)) * ((0) - ((d_1856_v2_).f35))
                        nw313_[int(2)] = d_1853_v0_
                        nw313_[int(3)] = (d_1853_v0_) + (-518)
                        nw313_[int(4)] = (d_1856_v2_).f35
                        nw313_[int(5)] = (len(d_1857_v3_)) * ((d_1865_v11_).cardinality)
                        nw313_[int(6)] = default__.safeModuloInt((d_1856_v2_).f35, (d_1856_v2_).f35)
                        nw313_[int(7)] = (d_1856_v2_).f35
                        nw313_[int(8)] = d_1853_v0_
                        d_1866_v12_ = nw313_
                        index323_ = default__.safeIndex(185, (d_1866_v12_).length(0))
                        (d_1866_v12_)[index323_] = (d_1856_v2_).f35
                        d_1867_v13_: _dafny.Array
                        def lambda73_(d_1868_v2_):
                            def lambda74_(d_1869_i1_):
                                return (d_1868_v2_).f36

                            return lambda74_

                        init40_ = lambda73_(d_1856_v2_)
                        nw314_ = _dafny.Array(None, 1)
                        for i0_40_ in range(nw314_.length(0)):
                            nw314_[i0_40_] = init40_(i0_40_)
                        d_1867_v13_ = nw314_
                        d_1870_v14_: _dafny.Array
                        def lambda75_(d_1871_v4_):
                            def lambda76_(d_1872_i2_):
                                return d_1871_v4_

                            return lambda76_

                        init41_ = lambda75_(d_1858_v4_)
                        nw315_ = _dafny.Array(None, 12)
                        for i0_41_ in range(nw315_.length(0)):
                            nw315_[i0_41_] = init41_(i0_41_)
                        d_1870_v14_ = nw315_
                        d_1873_v15_: D17
                        d_1873_v15_ = D17_DC39(d_1870_v14_)
                        d_1874_v16_: _dafny.Seq
                        d_1874_v16_ = _dafny.SeqWithoutIsStrInference([(d_1873_v15_).cf69, d_1870_v14_, d_1870_v14_, d_1870_v14_, d_1870_v14_])
                        d_1875_v17_: _dafny.Seq
                        d_1875_v17_ = _dafny.SeqWithoutIsStrInference([d_1855_v1_])
                        index324_ = default__.safeIndex(185, (d_1866_v12_).length(0))
                        rhs355_ = (d_1874_v16_)[default__.safeIndex((len(d_1861_v7_)) + ((0) - ((d_1856_v2_).f35)), len(d_1874_v16_))]
                        rhs356_ = (d_1855_v1_) and ((d_1856_v2_).f36)
                        rhs357_ = len(d_1875_v17_)
                        rhs358_ = d_1857_v3_
                        rhs359_ = d_1867_v13_
                        lhs295_ = globalState
                        lhs296_ = d_1866_v12_
                        lhs297_ = default__.safeIndex(185, (d_1866_v12_).length(0))
                        lhs295_.f2 = rhs355_
                        d_1855_v1_ = rhs356_
                        lhs296_[lhs297_] = rhs357_
                        r1 = rhs358_
                        d_1867_v13_ = rhs359_
                        (globalState).f27 = d_1855_v1_
                    elif True:
                        d_1876_v18_: _dafny.Array
                        nw316_ = _dafny.Array(False, 23)
                        d_1876_v18_ = nw316_
                        d_1876_v18_ = d_1876_v18_
                        d_1877_v19_: _dafny.Seq
                        d_1877_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ivbvvuiux"))
                        r1 = d_1877_v19_
                        d_1878_v20_: _dafny.Array
                        nw317_ = _dafny.Array(_dafny.Array(None, 0), 16)
                        d_1878_v20_ = nw317_
                        d_1878_v20_ = d_1878_v20_
                        d_1879_v21_: _dafny.Array
                        nw318_ = _dafny.Array(D18.default()(), 11)
                        d_1879_v21_ = nw318_
                        d_1879_v21_ = d_1879_v21_
                        d_1880_v22_: _dafny.Map
                        d_1880_v22_ = _dafny.Map({len(d_1877_v19_): d_1855_v1_})
                        d_1881_v24_: _dafny.MultiSet
                        d_1881_v24_ = _dafny.MultiSet([d_1853_v0_])
                        d_1882_v25_: _dafny.Set
                        d_1882_v25_ = _dafny.Set({False})
                        d_1883_v26_: _dafny.Seq
                        d_1883_v26_ = _dafny.SeqWithoutIsStrInference([d_1853_v0_, 325, d_1853_v0_])
                        d_1884_v27_: _dafny.Array
                        nw319_ = _dafny.Array(None, 26)
                        nw319_[int(0)] = len(d_1877_v19_)
                        def iife179_():
                            coll77_ = _dafny.Set()
                            compr_77_: int
                            for compr_77_ in (d_1880_v22_).keys.Elements:
                                d_1885_v23_: int = compr_77_
                                if (d_1885_v23_) in (d_1880_v22_):
                                    coll77_ = coll77_.union(_dafny.Set([default__.safeModuloInt(d_1885_v23_, -588)]))
                            return _dafny.Set(coll77_)
                        nw319_[int(1)] = len(iife179_()
                        )
                        nw319_[int(2)] = d_1853_v0_
                        nw319_[int(3)] = d_1853_v0_
                        nw319_[int(4)] = d_1853_v0_
                        nw319_[int(5)] = d_1853_v0_
                        nw319_[int(6)] = d_1853_v0_
                        nw319_[int(7)] = d_1853_v0_
                        nw319_[int(8)] = 461
                        nw319_[int(9)] = d_1853_v0_
                        nw319_[int(10)] = (0) - (d_1853_v0_)
                        nw319_[int(11)] = (d_1881_v24_).cardinality
                        nw319_[int(12)] = d_1853_v0_
                        nw319_[int(13)] = d_1853_v0_
                        nw319_[int(14)] = d_1853_v0_
                        nw319_[int(15)] = len(d_1877_v19_)
                        nw319_[int(16)] = 848
                        nw319_[int(17)] = d_1853_v0_
                        nw319_[int(18)] = d_1853_v0_
                        nw319_[int(19)] = d_1853_v0_
                        nw319_[int(20)] = len(d_1882_v25_)
                        nw319_[int(21)] = (d_1883_v26_)[default__.safeIndex(d_1853_v0_, len(d_1883_v26_))]
                        nw319_[int(22)] = d_1853_v0_
                        nw319_[int(23)] = 442
                        nw319_[int(24)] = d_1853_v0_
                        nw319_[int(25)] = d_1853_v0_
                        d_1884_v27_ = nw319_
                        d_1886_v28_: _dafny.Map
                        d_1886_v28_ = _dafny.Map({not(d_1855_v1_): d_1855_v1_})
                        default__.m0(d_1876_v18_, d_1884_v27_, len((d_1886_v28_ if d_1855_v1_ else d_1886_v28_)), len(d_1880_v22_), globalState)
                    (globalState).f11 = (d_1853_v0_) != (d_1853_v0_)
                    if ((0) - (d_1853_v0_)) == (d_1853_v0_):
                        r1 = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_1887_i3_ in range(default__.abs(515))])
                        d_1888_v29_: _dafny.Map
                        d_1888_v29_ = _dafny.Map({d_1853_v0_: d_1855_v1_})
                        d_1889_v30_: C5
                        nw320_ = C5()
                        nw320_.ctor__(d_1853_v0_, ((d_1888_v29_)[d_1853_v0_] if (d_1853_v0_) in (d_1888_v29_) else d_1855_v1_), (self).f28)
                        d_1889_v30_ = nw320_
                        d_1890_v31_: _dafny.Seq
                        d_1890_v31_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "epvq"))
                        r1 = d_1890_v31_
                        d_1891_v32_: _dafny.Set
                        d_1891_v32_ = _dafny.Set({d_1853_v0_, (0) - (len(_dafny.SeqWithoutIsStrInference([d_1855_v1_, True, False, d_1855_v1_])))})
                        d_1892_v33_: D21
                        d_1892_v33_ = D21_DC51(d_1891_v32_)
                        (globalState).f27 = (d_1891_v32_) == ((d_1891_v32_) | ((d_1892_v33_).cf96))
                        (globalState).f0 = d_1853_v0_
                    elif True:
                        d_1893_v34_: _dafny.Map
                        d_1893_v34_ = _dafny.Map({d_1855_v1_: d_1855_v1_})
                        d_1894_v35_: _dafny.Seq
                        d_1894_v35_ = _dafny.SeqWithoutIsStrInference([not(((d_1893_v34_)[d_1855_v1_] if (d_1855_v1_) in (d_1893_v34_) else d_1855_v1_)), d_1855_v1_, d_1855_v1_])
                        d_1895_v36_: _dafny.Seq
                        d_1895_v36_ = _dafny.SeqWithoutIsStrInference([d_1855_v1_, (d_1894_v35_)[default__.safeIndex(207, len(d_1894_v35_))]])
                        d_1896_v37_: _dafny.Seq
                        d_1896_v37_ = _dafny.SeqWithoutIsStrInference([d_1895_v36_])
                        d_1897_v38_: C6
                        nw321_ = C6()
                        nw321_.ctor__(d_1853_v0_)
                        d_1897_v38_ = nw321_
                        d_1898_v39_: _dafny.Map
                        d_1898_v39_ = _dafny.Map({d_1855_v1_: d_1897_v38_})
                        d_1899_v40_: _dafny.Set
                        d_1899_v40_ = _dafny.Set({d_1898_v39_, d_1898_v39_})
                        d_1900_v41_: _dafny.MultiSet
                        d_1900_v41_ = _dafny.MultiSet([len(d_1899_v40_), d_1853_v0_])
                        d_1901_v42_: _dafny.Seq
                        d_1901_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hncioyd"))
                        rhs360_ = (_dafny.SeqWithoutIsStrInference([d_1855_v1_])) + ((d_1896_v37_)[default__.safeIndex(((d_1900_v41_)[(d_1897_v38_).f34] if ((d_1897_v38_).f34) in (d_1900_v41_) else (d_1897_v38_).f34), len(d_1896_v37_))])
                        rhs361_ = d_1855_v1_
                        rhs362_ = d_1895_v36_
                        rhs363_ = len((d_1901_v42_) + (d_1901_v42_))
                        lhs298_ = globalState
                        lhs299_ = globalState
                        lhs300_ = globalState
                        lhs301_ = globalState
                        lhs298_.f16 = rhs360_
                        lhs299_.f27 = rhs361_
                        lhs300_.f12 = rhs362_
                        lhs301_.f19 = rhs363_
                        d_1902_v43_: _dafny.Set
                        d_1902_v43_ = _dafny.Set({not(d_1855_v1_), d_1855_v1_})
                        d_1903_v44_: str
                        d_1903_v44_ = _dafny.CodePoint('g')
                        d_1904_v45_: D2
                        d_1904_v45_ = D2_DC7((d_1855_v1_ if d_1855_v1_ else d_1855_v1_), ((d_1901_v42_)[default__.safeIndex(default__.fm2(default__.fm3(globalState), d_1902_v43_, _dafny.Map({False: d_1903_v44_}), d_1855_v1_, globalState), len(d_1901_v42_))]) not in (d_1901_v42_), d_1901_v42_)
                        d_1905_v46_: _dafny.Array
                        def lambda77_(d_1906_v0_):
                            def lambda78_(d_1907_i4_):
                                return (d_1907_i4_) * (d_1906_v0_)

                            return lambda78_

                        init42_ = lambda77_(d_1853_v0_)
                        nw322_ = _dafny.Array(None, 27)
                        for i0_42_ in range(nw322_.length(0)):
                            nw322_[i0_42_] = init42_(i0_42_)
                        d_1905_v46_ = nw322_
                        index325_ = default__.safeIndex(859, (d_1905_v46_).length(0))
                        (d_1905_v46_)[index325_] = ((d_1897_v38_).f34) - ((d_1897_v38_).f34)
                        index326_ = default__.safeIndex(859, (d_1905_v46_).length(0))
                        rhs364_ = d_1904_v45_
                        rhs365_ = default__.safeDivisionInt(d_1853_v0_, (d_1897_v38_).f34)
                        rhs366_ = default__.safeDivisionInt(len(d_1894_v35_), d_1853_v0_)
                        lhs302_ = globalState
                        lhs303_ = d_1905_v46_
                        lhs304_ = default__.safeIndex(859, (d_1905_v46_).length(0))
                        d_1904_v45_ = rhs364_
                        lhs302_.f24 = rhs365_
                        lhs303_[lhs304_] = rhs366_
                        d_1908_v47_: bool
                        out14_: bool
                        out14_ = (d_1897_v38_).m9(globalState)
                        d_1908_v47_ = out14_
                        d_1909_v48_: _dafny.Map
                        d_1909_v48_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_1908_v47_, d_1908_v47_]): len(d_1895_v36_)})
                        d_1910_v49_: C2
                        nw323_ = C2()
                        nw323_.ctor__(d_1909_v48_, d_1902_v43_, (self).f28)
                        d_1910_v49_ = nw323_
                        d_1910_v49_ = d_1910_v49_
                        d_1911_v50_: D19
                        d_1911_v50_ = D19_DC46(d_1853_v0_, not(d_1908_v47_), d_1855_v1_, not(d_1855_v1_))
                        (globalState).f1 = (d_1911_v50_).cf86
                    d_1912_v51_: _dafny.MultiSet
                    d_1912_v51_ = _dafny.MultiSet([d_1853_v0_])
                    d_1913_v52_: str
                    d_1913_v52_ = _dafny.CodePoint('b')
                    d_1914_v53_: D20
                    d_1914_v53_ = D20_DC49(d_1913_v52_, d_1853_v0_, d_1855_v1_, d_1855_v1_, d_1855_v1_)
                    d_1915_v54_: _dafny.Set
                    d_1915_v54_ = _dafny.Set({d_1855_v1_})
                    d_1916_v55_: _dafny.Seq
                    d_1916_v55_ = _dafny.SeqWithoutIsStrInference([d_1855_v1_])
                    d_1917_v56_: _dafny.Seq
                    d_1917_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kowhcca"))
                    d_1918_v57_: _dafny.Seq
                    d_1918_v57_ = _dafny.SeqWithoutIsStrInference([len(d_1917_v56_)])
                    d_1919_v58_: _dafny.MultiSet
                    d_1919_v58_ = _dafny.MultiSet([_dafny.CodePoint('l')])
                    d_1920_v60_: _dafny.Map
                    def iife180_():
                        coll78_ = _dafny.Set()
                        compr_78_: int
                        for compr_78_ in _dafny.IntegerRange(549, 983):
                            d_1921_v59_: int = compr_78_
                            if ((549) <= (d_1921_v59_)) and ((d_1921_v59_) < (983)):
                                coll78_ = coll78_.union(_dafny.Set([(d_1921_v59_) + (len(_dafny.Set({d_1853_v0_})))]))
                        return _dafny.Set(coll78_)
                    d_1920_v60_ = _dafny.Map({((d_1919_v58_)[(d_1917_v56_)[default__.safeIndex(d_1853_v0_, len(d_1917_v56_))]] if ((d_1917_v56_)[default__.safeIndex(d_1853_v0_, len(d_1917_v56_))]) in (d_1919_v58_) else len(iife180_()
                    )): d_1855_v1_})
                    d_1922_v61_: _dafny.Map
                    d_1922_v61_ = _dafny.Map({d_1913_v52_: d_1920_v60_})
                    d_1923_v62_: _dafny.Array
                    nw324_ = _dafny.Array(None, 26)
                    nw324_[int(0)] = d_1853_v0_
                    nw324_[int(1)] = ((d_1912_v51_)[d_1853_v0_] if (d_1853_v0_) in (d_1912_v51_) else d_1853_v0_)
                    nw324_[int(2)] = ((d_1914_v53_).cf91) + (d_1853_v0_)
                    nw324_[int(3)] = d_1853_v0_
                    nw324_[int(4)] = len(d_1915_v54_)
                    nw324_[int(5)] = (len(d_1916_v55_)) + ((D19_DC46(996, d_1855_v1_, d_1855_v1_, d_1855_v1_)).cf84)
                    nw324_[int(6)] = d_1853_v0_
                    nw324_[int(7)] = d_1853_v0_
                    nw324_[int(8)] = d_1853_v0_
                    nw324_[int(9)] = (d_1853_v0_) + (d_1853_v0_)
                    nw324_[int(10)] = default__.safeModuloInt(d_1853_v0_, len(d_1916_v55_))
                    nw324_[int(11)] = (0) - ((d_1918_v57_)[default__.safeIndex(d_1853_v0_, len(d_1918_v57_))])
                    nw324_[int(12)] = default__.safeModuloInt(d_1853_v0_, d_1853_v0_)
                    nw324_[int(13)] = 854
                    nw324_[int(14)] = (d_1853_v0_) - (d_1853_v0_)
                    nw324_[int(15)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jg"))) + (d_1917_v56_))
                    nw324_[int(16)] = (0) - ((d_1918_v57_)[default__.safeIndex(d_1853_v0_, len(d_1918_v57_))])
                    nw324_[int(17)] = (0) - (d_1853_v0_)
                    nw324_[int(18)] = d_1853_v0_
                    nw324_[int(19)] = d_1853_v0_
                    nw324_[int(20)] = len((d_1922_v61_) | (d_1922_v61_))
                    nw324_[int(21)] = d_1853_v0_
                    nw324_[int(22)] = d_1853_v0_
                    nw324_[int(23)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rpwmkcsw")))
                    nw324_[int(24)] = d_1853_v0_
                    nw324_[int(25)] = d_1853_v0_
                    d_1923_v62_ = nw324_
                    index327_ = default__.safeIndex(563, (d_1923_v62_).length(0))
                    (d_1923_v62_)[index327_] = default__.safeDivisionInt((0) - (-331), d_1853_v0_)
                    index328_ = default__.safeIndex(563, (d_1923_v62_).length(0))
                    (d_1923_v62_)[index328_] = d_1853_v0_
                    pass
            pass
        d_1924_v63_: _dafny.Array
        nw325_ = _dafny.Array(int(0), 6)
        d_1924_v63_ = nw325_
        index329_ = default__.safeIndex(833, (d_1924_v63_).length(0))
        (d_1924_v63_)[index329_] = d_1853_v0_
        d_1925_v64_: bool
        d_1925_v64_ = False
        d_1926_v65_: _dafny.MultiSet
        d_1926_v65_ = _dafny.MultiSet([d_1925_v64_])
        d_1927_v66_: _dafny.Map
        d_1927_v66_ = _dafny.Map({d_1925_v64_: d_1926_v65_})
        index330_ = default__.safeIndex(833, (d_1924_v63_).length(0))
        (d_1924_v63_)[index330_] = (d_1853_v0_) * (len(d_1927_v66_))
        d_1928_v67_: _dafny.Array
        nw326_ = _dafny.Array(_dafny.Seq({}), 1)
        d_1928_v67_ = nw326_
        index331_ = default__.safeIndex(64, (d_1928_v67_).length(0))
        (d_1928_v67_)[index331_] = _dafny.SeqWithoutIsStrInference([d_1925_v64_, d_1925_v64_, d_1925_v64_, d_1925_v64_])
        d_1929_v68_: _dafny.Seq
        d_1929_v68_ = _dafny.SeqWithoutIsStrInference([d_1925_v64_])
        d_1930_v69_: _dafny.Map
        d_1930_v69_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1931_i5_ in range(default__.abs(885))])): d_1929_v68_})
        d_1932_v70_: _dafny.Set
        d_1932_v70_ = _dafny.Set({(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], d_1853_v0_, len(((d_1930_v69_)[(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]] if ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) in (d_1930_v69_) else d_1929_v68_)), (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]})
        d_1933_v71_: _dafny.Map
        d_1933_v71_ = _dafny.Map({d_1925_v64_: d_1932_v70_})
        index332_ = default__.safeIndex(64, (d_1928_v67_).length(0))
        rhs367_ = (self).fm7((d_1932_v70_).ispropersubset(((d_1933_v71_)[True] if (True) in (d_1933_v71_) else d_1932_v70_)), d_1925_v64_, globalState)
        rhs368_ = (_dafny.SeqWithoutIsStrInference([d_1925_v64_]) if d_1925_v64_ else d_1929_v68_)
        lhs305_ = d_1928_v67_
        lhs306_ = default__.safeIndex(64, (d_1928_v67_).length(0))
        d_1925_v64_ = rhs367_
        lhs305_[lhs306_] = rhs368_
        d_1934_v72_: _dafny.Map
        d_1934_v72_ = _dafny.Map({d_1925_v64_: (0) - ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])})
        d_1935_v73_: D0
        d_1935_v73_ = D0_DC1(len(d_1934_v72_))
        d_1936_v74_: _dafny.MultiSet
        d_1936_v74_ = _dafny.MultiSet([(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], (d_1935_v73_).cf1])
        pat_let_tv40_ = d_1925_v64_
        pat_let_tv41_ = d_1925_v64_
        def lambda79_(source25_):
            if source25_.is_DC52:
                d_1937___mcc_h0_ = source25_.cf97
                d_1938___mcc_h1_ = source25_.cf98
                d_1939___mcc_h2_ = source25_.cf99
                d_1940___mcc_h3_ = source25_.cf100
                d_1941___mcc_h4_ = source25_.cf101
                d_1942_cf101_ = d_1941___mcc_h4_
                d_1943_cf100_ = d_1940___mcc_h3_
                d_1944_cf99_ = d_1939___mcc_h2_
                d_1945_cf98_ = d_1938___mcc_h1_
                d_1946_cf97_ = d_1937___mcc_h0_
                return d_1942_cf101_
            elif source25_.is_DC53:
                d_1947___mcc_h5_ = source25_.cf102
                d_1948_cf102_ = d_1947___mcc_h5_
                return pat_let_tv40_
            elif source25_.is_DC51:
                d_1949___mcc_h6_ = source25_.cf96
                d_1950_cf96_ = d_1949___mcc_h6_
                return not((992) == (976))
            elif True:
                d_1951___mcc_h7_ = source25_.cf103
                d_1952_cf103_ = d_1951___mcc_h7_
                return pat_let_tv41_

        if lambda79_(D21_DC52((d_1936_v74_).cardinality, d_1925_v64_, False, (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], d_1925_v64_)):
            d_1953_v75_: _dafny.Set
            d_1953_v75_ = _dafny.Set({d_1932_v70_, d_1932_v70_})
            d_1954_v76_: _dafny.Seq
            d_1954_v76_ = _dafny.SeqWithoutIsStrInference([d_1953_v75_, d_1953_v75_])
            d_1955_v77_: _dafny.Seq
            d_1955_v77_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))
            (globalState).f11 = ((d_1954_v76_)[default__.safeIndex(len(d_1955_v77_), len(d_1954_v76_))]).ispropersubset(d_1953_v75_)
            d_1956_v78_: _dafny.Seq
            d_1956_v78_ = _dafny.SeqWithoutIsStrInference([d_1853_v0_, (0) - (len(d_1955_v77_))])
            d_1957_v79_: _dafny.Seq
            d_1957_v79_ = _dafny.SeqWithoutIsStrInference([(d_1956_v78_)[default__.safeIndex((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], len(d_1956_v78_))], (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]])
            d_1958_v80_: _dafny.Seq
            d_1958_v80_ = _dafny.SeqWithoutIsStrInference([d_1956_v78_])
            (globalState).f0 = len(d_1958_v80_)
            d_1959_v81_: _dafny.Array
            nw327_ = _dafny.Array(_dafny.Seq({}), 29)
            d_1959_v81_ = nw327_
            d_1960_v82_: _dafny.Map
            d_1960_v82_ = _dafny.Map({d_1959_v81_: not((d_1934_v72_) != (d_1934_v72_))})
            d_1961_v83_: _dafny.Array
            d_1961_v83_ = d_1959_v81_
            d_1960_v82_ = (d_1960_v82_).set(d_1961_v83_, d_1925_v64_)
            (globalState).f1 = d_1925_v64_
            d_1962_v84_: _dafny.Map
            d_1962_v84_ = _dafny.Map({d_1853_v0_: not(d_1925_v64_)})
            (globalState).f14 = len((d_1962_v84_) | (d_1962_v84_))
        elif True:
            if (d_1925_v64_ if (d_1926_v65_).isdisjoint(d_1926_v65_) else d_1925_v64_):
                d_1963_v85_: _dafny.Seq
                d_1963_v85_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilwl"))
                d_1964_v86_: _dafny.Seq
                d_1964_v86_ = _dafny.SeqWithoutIsStrInference([(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], (0) - (d_1853_v0_), len(default__.fm6(not(d_1925_v64_), ((d_1934_v72_)[d_1925_v64_] if (d_1925_v64_) in (d_1934_v72_) else (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]), d_1963_v85_, globalState))])
                d_1964_v86_ = d_1964_v86_
                d_1965_v87_: _dafny.Set
                d_1965_v87_ = _dafny.Set({d_1925_v64_, d_1925_v64_})
                d_1966_v88_: T0
                nw328_ = C2()
                nw328_.ctor__(_dafny.Map({_dafny.SeqWithoutIsStrInference([d_1925_v64_, True, d_1925_v64_, d_1925_v64_, not(True)]): (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]}), d_1965_v87_, (self).f28)
                d_1966_v88_ = nw328_
                d_1967_v89_: str
                d_1968_v90_: int
                d_1969_v91_: _dafny.Map
                d_1970_v92_: _dafny.Set
                out15_: str
                out16_: int
                out17_: _dafny.Map
                out18_: _dafny.Set
                out15_, out16_, out17_, out18_ = (self).m8((0) - (len((d_1928_v67_)[default__.safeIndex(64, (d_1928_v67_).length(0))])), d_1925_v64_, (0) - (((d_1964_v86_)[default__.safeIndex((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], len(d_1964_v86_))]) + (len(_dafny.SeqWithoutIsStrInference([d_1925_v64_, d_1925_v64_])))), d_1966_v88_, globalState)
                d_1967_v89_ = out15_
                d_1968_v90_ = out16_
                d_1969_v91_ = out17_
                d_1970_v92_ = out18_
                (globalState).f11 = ((d_1936_v74_).cardinality) >= (len(default__.fm6(((d_1928_v67_)[default__.safeIndex(64, (d_1928_v67_).length(0))])[default__.safeIndex(d_1968_v90_, len((d_1928_v67_)[default__.safeIndex(64, (d_1928_v67_).length(0))]))], d_1968_v90_, d_1963_v85_, globalState)))
                d_1971_v93_: _dafny.Map
                d_1971_v93_ = _dafny.Map({_dafny.Map({d_1968_v90_: d_1967_v89_}): True})
                d_1971_v93_ = (default__.fm39(d_1968_v90_, globalState)) | (d_1971_v93_)
                d_1972_v94_: _dafny.Seq
                d_1972_v94_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ihkanrsjf"))])
                d_1973_v95_: D8
                d_1973_v95_ = D8_DC20((self).fm15(d_1925_v64_, len(d_1929_v68_), len((d_1972_v94_)[default__.safeIndex((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], len(d_1972_v94_))]), globalState), (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
                index333_ = default__.safeIndex(833, (d_1924_v63_).length(0))
                (d_1924_v63_)[index333_] = (d_1973_v95_).cf32
            elif True:
                d_1974_v96_: _dafny.Set
                d_1974_v96_ = _dafny.Set({d_1925_v64_, d_1925_v64_, d_1925_v64_, not(not(d_1925_v64_)), default__.fm4(globalState)})
                d_1975_v97_: _dafny.Seq
                d_1975_v97_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jj"))
                d_1976_v98_: _dafny.Seq
                d_1976_v98_ = _dafny.SeqWithoutIsStrInference([len(d_1975_v97_)])
                nw329_ = _dafny.Array(None, 25)
                nw329_[int(0)] = d_1853_v0_
                nw329_[int(1)] = 995
                nw329_[int(2)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(3)] = default__.safeDivisionInt(len(d_1974_v96_), (0) - (d_1853_v0_))
                nw329_[int(4)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(5)] = ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) * ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
                nw329_[int(6)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(7)] = ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbo"))))
                nw329_[int(8)] = 230
                nw329_[int(9)] = default__.safeDivisionInt((d_1976_v98_)[default__.safeIndex(len(d_1976_v98_), len(d_1976_v98_))], (d_1976_v98_)[default__.safeIndex(d_1853_v0_, len(d_1976_v98_))])
                nw329_[int(10)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(11)] = -545
                nw329_[int(12)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(13)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(14)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(15)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(16)] = d_1853_v0_
                nw329_[int(17)] = ((d_1926_v65_)[d_1925_v64_] if (d_1925_v64_) in (d_1926_v65_) else -208)
                nw329_[int(18)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(19)] = (0) - ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
                nw329_[int(20)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                nw329_[int(21)] = d_1853_v0_
                nw329_[int(22)] = (0) - (d_1853_v0_)
                nw329_[int(23)] = default__.safeDivisionInt(d_1853_v0_, (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
                nw329_[int(24)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                d_1924_v63_ = nw329_
                d_1977_v99_: D15
                d_1977_v99_ = D15_DC34(d_1925_v64_, (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
                d_1978_v100_: _dafny.Array
                nw330_ = _dafny.Array(None, 16)
                nw330_[int(0)] = False
                nw330_[int(1)] = d_1925_v64_
                nw330_[int(2)] = d_1925_v64_
                nw330_[int(3)] = (d_1977_v99_).cf59
                nw330_[int(4)] = not (True) or (d_1925_v64_)
                nw330_[int(5)] = d_1925_v64_
                nw330_[int(6)] = d_1925_v64_
                nw330_[int(7)] = d_1925_v64_
                nw330_[int(8)] = d_1925_v64_
                nw330_[int(9)] = d_1925_v64_
                nw330_[int(10)] = d_1925_v64_
                nw330_[int(11)] = True
                nw330_[int(12)] = (False if True else (self).fm7(d_1925_v64_, d_1925_v64_, globalState))
                nw330_[int(13)] = ((self).fm7(d_1925_v64_, d_1925_v64_, globalState)) == (d_1925_v64_)
                nw330_[int(14)] = d_1925_v64_
                nw330_[int(15)] = (d_1936_v74_).issubset(d_1936_v74_)
                d_1978_v100_ = nw330_
                d_1979_v101_: _dafny.Map
                d_1979_v101_ = _dafny.Map({d_1925_v64_: d_1929_v68_})
                d_1980_v102_: _dafny.Seq
                d_1980_v102_ = _dafny.SeqWithoutIsStrInference([(d_1926_v65_ if d_1925_v64_ else d_1926_v65_), d_1926_v65_, d_1926_v65_, d_1926_v65_])
                rhs369_ = not((d_1853_v0_) != (((0) - (len(((d_1979_v101_)[d_1925_v64_] if (d_1925_v64_) in (d_1979_v101_) else (d_1928_v67_)[default__.safeIndex(64, (d_1928_v67_).length(0))])))) + ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])))
                rhs370_ = not(not (((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) not in (_dafny.MultiSet(d_1976_v98_))) or (False))
                rhs371_ = not((_dafny.MultiSet(d_1976_v98_)).isdisjoint((d_1936_v74_) - (d_1936_v74_)))
                rhs372_ = d_1978_v100_
                rhs373_ = (0) - (((d_1980_v102_)[default__.safeIndex(default__.safeModuloInt((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], d_1853_v0_), len(d_1980_v102_))]).cardinality)
                lhs307_ = globalState
                lhs308_ = globalState
                lhs309_ = globalState
                lhs307_.f11 = rhs369_
                lhs308_.f1 = rhs370_
                lhs309_.f11 = rhs371_
                d_1978_v100_ = rhs372_
                d_1853_v0_ = rhs373_
                d_1981_v103_: C6
                nw331_ = C6()
                nw331_.ctor__((0) - ((self).fm15(d_1925_v64_, 285, (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], globalState)))
                d_1981_v103_ = nw331_
                d_1981_v103_ = d_1981_v103_
                d_1982_v104_: _dafny.Array
                nw332_ = _dafny.Array(_dafny.CodePoint('D'), 27)
                d_1982_v104_ = nw332_
                d_1983_v105_: _dafny.Map
                d_1983_v105_ = _dafny.Map({d_1982_v104_: (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]})
                (globalState).f0 = default__.safeModuloInt(((0) - (len(d_1983_v105_)) if d_1925_v64_ else len(d_1975_v97_)), d_1853_v0_)
                (globalState).f1 = not (d_1925_v64_) or ((d_1975_v97_) != (d_1975_v97_))
            if ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) != ((980) - ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])):
                d_1984_v106_: str
                d_1984_v106_ = _dafny.CodePoint('c')
                d_1985_v107_: _dafny.Seq
                d_1985_v107_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vllishi"))
                d_1986_v108_: _dafny.Map
                d_1986_v108_ = _dafny.Map({(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]: d_1925_v64_})
                d_1987_v109_: _dafny.Array
                nw333_ = _dafny.Array(None, 12)
                nw333_[int(0)] = (d_1925_v64_) or (d_1925_v64_)
                nw333_[int(1)] = d_1925_v64_
                nw333_[int(2)] = d_1925_v64_
                nw333_[int(3)] = (d_1984_v106_) not in (d_1985_v107_)
                nw333_[int(4)] = (d_1925_v64_) == (not(d_1925_v64_))
                nw333_[int(5)] = d_1925_v64_
                nw333_[int(6)] = d_1925_v64_
                nw333_[int(7)] = d_1925_v64_
                nw333_[int(8)] = ((d_1986_v108_)[(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]] if ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) in (d_1986_v108_) else not(True))
                nw333_[int(9)] = d_1925_v64_
                nw333_[int(10)] = d_1925_v64_
                nw333_[int(11)] = True
                d_1987_v109_ = nw333_
                index334_ = default__.safeIndex(184, (d_1987_v109_).length(0))
                (d_1987_v109_)[index334_] = d_1925_v64_
                index335_ = default__.safeIndex(184, (d_1987_v109_).length(0))
                (d_1987_v109_)[index335_] = (d_1985_v107_) != (d_1985_v107_)
                d_1988_v110_: _dafny.Map
                d_1988_v110_ = _dafny.Map({False: d_1984_v106_})
                index336_ = default__.safeIndex(184, (d_1987_v109_).length(0))
                index337_ = default__.safeIndex(833, (d_1924_v63_).length(0))
                rhs374_ = ((d_1928_v67_)[default__.safeIndex(64, (d_1928_v67_).length(0))]) + (d_1929_v68_)
                rhs375_ = d_1925_v64_
                rhs376_ = d_1925_v64_
                rhs377_ = default__.safeDivisionInt((0) - ((d_1853_v0_) * (default__.fm2(d_1985_v107_, _dafny.Set({not(d_1925_v64_)}), d_1988_v110_, (d_1987_v109_)[default__.safeIndex(184, (d_1987_v109_).length(0))], globalState))), 184)
                rhs378_ = d_1984_v106_
                lhs310_ = globalState
                lhs311_ = globalState
                lhs312_ = d_1987_v109_
                lhs313_ = default__.safeIndex(184, (d_1987_v109_).length(0))
                lhs314_ = d_1924_v63_
                lhs315_ = default__.safeIndex(833, (d_1924_v63_).length(0))
                lhs310_.f16 = rhs374_
                lhs311_.f27 = rhs375_
                lhs312_[lhs313_] = rhs376_
                lhs314_[lhs315_] = rhs377_
                d_1984_v106_ = rhs378_
                d_1984_v106_ = d_1984_v106_
                d_1989_v111_: C6
                nw334_ = C6()
                nw334_.ctor__((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
                d_1989_v111_ = nw334_
                d_1990_v112_: _dafny.Array
                def lambda80_(d_1991_v107_):
                    def lambda81_(d_1992_i6_):
                        return d_1991_v107_

                    return lambda81_

                init43_ = lambda80_(d_1985_v107_)
                nw335_ = _dafny.Array(None, 5)
                for i0_43_ in range(nw335_.length(0)):
                    nw335_[i0_43_] = init43_(i0_43_)
                d_1990_v112_ = nw335_
                d_1993_v113_: D22
                d_1993_v113_ = D22_DC55(d_1990_v112_)
                d_1994_v114_: C4
                nw336_ = C4()
                nw336_.ctor__(True, (d_1993_v113_).cf104)
                d_1994_v114_ = nw336_
            elif True:
                d_1995_v115_: str
                d_1995_v115_ = _dafny.CodePoint('f')
                d_1996_v116_: _dafny.Map
                d_1996_v116_ = _dafny.Map({((d_1926_v65_)[not(d_1925_v64_)] if (not(d_1925_v64_)) in (d_1926_v65_) else d_1853_v0_): (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]})
                d_1997_v117_: _dafny.Map
                d_1997_v117_ = _dafny.Map({d_1995_v115_: (len(d_1996_v116_)) not in (d_1996_v116_)})
                (globalState).f1 = ((d_1997_v117_)[d_1995_v115_] if (d_1995_v115_) in (d_1997_v117_) else d_1925_v64_)
                index338_ = default__.safeIndex(833, (d_1924_v63_).length(0))
                (d_1924_v63_)[index338_] = d_1853_v0_
                d_1995_v115_ = d_1995_v115_
                d_1998_v118_: _dafny.Seq
                d_1998_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qoovryke"))
                d_1999_v119_: _dafny.Map
                d_1999_v119_ = _dafny.Map({d_1925_v64_: (len(d_1998_v118_)) > (-415)})
                d_1999_v119_ = (d_1999_v119_).set(d_1925_v64_, (d_1929_v68_)[default__.safeIndex(len(d_1929_v68_), len(d_1929_v68_))])
                nw337_ = _dafny.Array(int(0), 14)
                d_1924_v63_ = nw337_
            d_2000_v120_: _dafny.Map
            d_2000_v120_ = _dafny.Map({d_1853_v0_: d_1853_v0_})
            d_2001_v122_: _dafny.Map
            d_2001_v122_ = _dafny.Map({d_1853_v0_: (self).fm8(_dafny.Map({(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]: (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]}), d_1925_v64_, d_1925_v64_, globalState)})
            d_2002_v123_: _dafny.Set
            d_2002_v123_ = _dafny.Set({d_1925_v64_, False})
            d_2003_v124_: _dafny.Map
            d_2003_v124_ = _dafny.Map({902: d_1925_v64_})
            d_2004_v125_: D17
            d_2004_v125_ = D17_DC38(d_1926_v65_)
            d_2005_v126_: _dafny.Map
            d_2005_v126_ = _dafny.Map({d_1925_v64_: d_1925_v64_})
            d_2006_v127_: _dafny.Array
            nw338_ = _dafny.Array(None, 25)
            nw338_[int(0)] = d_1925_v64_
            nw338_[int(1)] = (d_1925_v64_) == (d_1925_v64_)
            nw338_[int(2)] = d_1925_v64_
            nw338_[int(3)] = not(d_1925_v64_)
            nw338_[int(4)] = d_1925_v64_
            nw338_[int(5)] = d_1925_v64_
            nw338_[int(6)] = False
            def iife181_():
                coll79_ = _dafny.Map()
                compr_79_: int
                for compr_79_ in (d_2001_v122_).keys.Elements:
                    d_2007_v121_: int = compr_79_
                    if (d_2007_v121_) in (d_2001_v122_):
                        coll79_[default__.safeDivisionInt(d_2007_v121_, d_1853_v0_)] = (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]
                return _dafny.Map(coll79_)
            nw338_[int(7)] = ((d_2000_v120_).set(d_1853_v0_, (d_1936_v74_).cardinality)) != (iife181_()
            )
            nw338_[int(8)] = d_1925_v64_
            nw338_[int(9)] = d_1925_v64_
            nw338_[int(10)] = not(d_1925_v64_)
            nw338_[int(11)] = default__.fm4(globalState)
            nw338_[int(12)] = not((d_1925_v64_) not in (d_2002_v123_))
            nw338_[int(13)] = d_1925_v64_
            nw338_[int(14)] = d_1925_v64_
            nw338_[int(15)] = ((d_2003_v124_)[(0) - (((d_2004_v125_).cf68).cardinality)] if ((0) - (((d_2004_v125_).cf68).cardinality)) in (d_2003_v124_) else d_1925_v64_)
            nw338_[int(16)] = (False) or (d_1925_v64_)
            nw338_[int(17)] = d_1925_v64_
            nw338_[int(18)] = not(d_1925_v64_)
            nw338_[int(19)] = d_1925_v64_
            nw338_[int(20)] = d_1925_v64_
            nw338_[int(21)] = (((d_2005_v126_)[False] if (False) in (d_2005_v126_) else not(d_1925_v64_))) == (d_1925_v64_)
            nw338_[int(22)] = not (d_1925_v64_) or (not(not(True)))
            nw338_[int(23)] = not (((d_2005_v126_)[d_1925_v64_] if (d_1925_v64_) in (d_2005_v126_) else d_1925_v64_)) or (d_1925_v64_)
            nw338_[int(24)] = (d_1853_v0_) < ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])
            d_2006_v127_ = nw338_
            index339_ = default__.safeIndex(609, (d_2006_v127_).length(0))
            (d_2006_v127_)[index339_] = d_1925_v64_
            index340_ = default__.safeIndex(609, (d_2006_v127_).length(0))
            (d_2006_v127_)[index340_] = not (((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]) > ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])) or (d_1925_v64_)
            d_2008_v128_: _dafny.Seq
            d_2008_v128_ = _dafny.SeqWithoutIsStrInference([(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], d_1853_v0_])
            d_2009_v129_: _dafny.Map
            d_2009_v129_ = _dafny.Map({default__.fm5((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))], d_1925_v64_, globalState): d_2008_v128_})
            d_2010_v130_: _dafny.Seq
            d_2010_v130_ = _dafny.SeqWithoutIsStrInference([len(d_2009_v129_)])
            d_2011_v131_: _dafny.Map
            d_2011_v131_ = _dafny.Map({(d_2010_v130_) == (_dafny.SeqWithoutIsStrInference([len((d_1928_v67_)[default__.safeIndex(64, (d_1928_v67_).length(0))])])): d_2006_v127_})
            d_2011_v131_ = (d_2011_v131_).set(d_1925_v64_, d_2006_v127_)
            (globalState).f14 = ((0) - (d_1853_v0_)) * (d_1853_v0_)
        d_2012_v132_: _dafny.Seq
        d_2012_v132_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxjwsitb"))
        d_2013_v133_: _dafny.Map
        d_2013_v133_ = _dafny.Map({(d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]: d_2012_v132_})
        d_2013_v133_ = (d_2013_v133_).set(default__.safeModuloInt(d_1853_v0_, (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]), d_2012_v132_)
        d_2014_v134_: _dafny.Map
        d_2014_v134_ = _dafny.Map({d_1853_v0_: d_1925_v64_})
        d_2015_v135_: _dafny.Map
        d_2015_v135_ = _dafny.Map({(0) - ((d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))]): d_1853_v0_})
        d_2016_v136_: C5
        nw339_ = C5()
        nw339_.ctor__((0) - (default__.safeDivisionInt(d_1853_v0_, (d_1924_v63_)[default__.safeIndex(833, (d_1924_v63_).length(0))])), (d_1925_v64_ if d_1925_v64_ else ((d_2014_v134_)[((d_2015_v135_)[d_1853_v0_] if (d_1853_v0_) in (d_2015_v135_) else d_1853_v0_)] if (((d_2015_v135_)[d_1853_v0_] if (d_1853_v0_) in (d_2015_v135_) else d_1853_v0_)) in (d_2014_v134_) else d_1925_v64_)), (self).f28)
        d_2016_v136_ = nw339_
        r0 = D2_DC8((d_2016_v136_).f36, d_1924_v63_)
        d_2017_v137_: _dafny.Array
        nw340_ = _dafny.Array(None, 6)
        nw340_[int(0)] = (d_2016_v136_).f36
        nw340_[int(1)] = d_1925_v64_
        nw340_[int(2)] = (d_2016_v136_).f36
        nw340_[int(3)] = not((d_2016_v136_).f36)
        nw340_[int(4)] = (d_2016_v136_).f36
        nw340_[int(5)] = d_1925_v64_
        d_2017_v137_ = nw340_
        d_2018_v138_: D2
        d_2018_v138_ = D2_DC6(d_2017_v137_)
        d_2019_v139_: D12
        d_2019_v139_ = D12_DC28((d_2016_v136_).f36, d_2012_v132_, (d_2016_v136_).f36, ((d_1936_v74_).set((0) - ((d_2016_v136_).f35), default__.abs((d_2016_v136_).f35))).cardinality, d_2018_v138_)
        r1 = ((d_2012_v132_ if d_1925_v64_ else (default__.fm3(globalState)).set(default__.safeIndex((d_2016_v136_).f35, len(default__.fm3(globalState))), _dafny.CodePoint('s')))) + ((d_2019_v139_).cf46)
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_2020_v0_: int
        d_2020_v0_ = -581
        d_2021_v1_: _dafny.MultiSet
        d_2021_v1_ = _dafny.MultiSet([p0, p0, p0, False, p0])
        d_2022_v2_: _dafny.Map
        d_2022_v2_ = _dafny.Map({d_2020_v0_: d_2021_v1_})
        d_2023_v3_: _dafny.MultiSet
        d_2023_v3_ = _dafny.MultiSet([d_2020_v0_])
        d_2024_v4_: _dafny.Map
        d_2024_v4_ = _dafny.Map({d_2020_v0_: d_2023_v3_})
        d_2022_v2_ = (d_2022_v2_).set((((d_2024_v4_)[(0) - ((d_2023_v3_).cardinality)] if ((0) - ((d_2023_v3_).cardinality)) in (d_2024_v4_) else d_2023_v3_)).cardinality, d_2021_v1_)
        d_2025_v5_: _dafny.Seq
        d_2025_v5_ = _dafny.SeqWithoutIsStrInference([not(p0)])
        d_2026_v6_: _dafny.Seq
        d_2026_v6_ = _dafny.SeqWithoutIsStrInference([len((d_2025_v5_) + (d_2025_v5_))])
        d_2027_v7_: _dafny.Seq
        d_2027_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "frol"))
        d_2028_v8_: D16
        d_2028_v8_ = D16_DC37(d_2025_v5_, d_2027_v7_, p1)
        d_2029_v9_: _dafny.Array
        nw341_ = _dafny.Array(None, 4)
        nw341_[int(0)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gpylep"))
        nw341_[int(1)] = (d_2027_v7_) + ((d_2028_v8_).cf66)
        nw341_[int(2)] = d_2027_v7_
        nw341_[int(3)] = _dafny.SeqWithoutIsStrInference([p1 for d_2030_i0_ in range(default__.abs(444))])
        d_2029_v9_ = nw341_
        rhs379_ = default__.fm6((d_2027_v7_) == (d_2027_v7_), d_2020_v0_, d_2027_v7_, globalState)
        rhs380_ = d_2029_v9_
        rhs381_ = default__.safeModuloInt(d_2020_v0_, d_2020_v0_)
        d_2026_v6_ = rhs379_
        d_2029_v9_ = rhs380_
        r0 = rhs381_
        d_2031_v10_: D1
        d_2031_v10_ = D1_DC5(p0, p0, len(d_2027_v7_), p0, p0)
        source26_ = d_2031_v10_
        if source26_.is_DC4:
            d_2032___mcc_h0_ = source26_.cf5
            d_2033___mcc_h1_ = source26_.cf6
            d_2034_cf6_ = d_2033___mcc_h1_
            d_2035_cf5_ = d_2032___mcc_h0_
            if False:
                d_2036_v11_: C5
                nw342_ = C5()
                nw342_.ctor__(876, p0, (self).f28)
                d_2036_v11_ = nw342_
                (globalState).f27 = p0
                d_2037_v12_: _dafny.Map
                d_2037_v12_ = _dafny.Map({p0: (d_2036_v11_).fm7((d_2036_v11_).f36, d_2034_cf6_, globalState)})
                d_2038_v13_: C5
                nw343_ = C5()
                nw343_.ctor__(-715, ((d_2037_v12_)[(d_2036_v11_).f36] if ((d_2036_v11_).f36) in (d_2037_v12_) else True), (self).f28)
                d_2038_v13_ = nw343_
                rhs382_ = d_2027_v7_
                rhs383_ = d_2027_v7_
                d_2027_v7_ = rhs382_
                d_2027_v7_ = rhs383_
                d_2039_v14_: str
                d_2039_v14_ = _dafny.CodePoint('w')
                rhs384_ = d_2039_v14_
                rhs385_ = default__.safeDivisionInt(default__.safeModuloInt((d_2038_v13_).f35, len(d_2027_v7_)), (d_2036_v11_).f35)
                lhs316_ = globalState
                d_2039_v14_ = rhs384_
                lhs316_.f14 = rhs385_
            elif True:
                d_2040_v15_: _dafny.Map
                d_2040_v15_ = _dafny.Map({d_2034_cf6_: 485})
                d_2041_v16_: _dafny.Map
                d_2041_v16_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_2020_v0_, d_2020_v0_])): d_2035_cf5_})
                (globalState).f19 = default__.safeModuloInt(((d_2040_v15_)[p0] if (p0) in (d_2040_v15_) else d_2035_cf5_), (len(d_2041_v16_)) + (d_2035_cf5_))
                d_2042_v17_: C6
                nw344_ = C6()
                nw344_.ctor__(d_2035_cf5_)
                d_2042_v17_ = nw344_
                (globalState).f24 = default__.safeDivisionInt(955, d_2020_v0_)
                (globalState).f14 = (d_2020_v0_) - ((0) - (d_2035_cf5_))
                (globalState).f27 = default__.fm4(globalState)
            d_2043_v18_: D15
            d_2043_v18_ = D15_DC34(d_2034_cf6_, (0) - (d_2020_v0_))
            d_2044_v19_: _dafny.Array
            nw345_ = _dafny.Array(None, 24)
            nw345_[int(0)] = not(p0)
            nw345_[int(1)] = d_2034_cf6_
            nw345_[int(2)] = default__.fm4(globalState)
            nw345_[int(3)] = (d_2035_cf5_) <= (d_2020_v0_)
            nw345_[int(4)] = d_2034_cf6_
            nw345_[int(5)] = d_2034_cf6_
            nw345_[int(6)] = d_2034_cf6_
            nw345_[int(7)] = p0
            nw345_[int(8)] = False
            nw345_[int(9)] = d_2034_cf6_
            nw345_[int(10)] = p0
            nw345_[int(11)] = d_2034_cf6_
            nw345_[int(12)] = (len(_dafny.SeqWithoutIsStrInference([p0]))) == ((0) - (d_2020_v0_))
            nw345_[int(13)] = d_2034_cf6_
            nw345_[int(14)] = p0
            nw345_[int(15)] = (p0 if d_2034_cf6_ else d_2034_cf6_)
            nw345_[int(16)] = d_2034_cf6_
            nw345_[int(17)] = (d_2043_v18_).cf59
            nw345_[int(18)] = True
            nw345_[int(19)] = p0
            nw345_[int(20)] = d_2034_cf6_
            nw345_[int(21)] = d_2034_cf6_
            nw345_[int(22)] = d_2034_cf6_
            nw345_[int(23)] = d_2034_cf6_
            d_2044_v19_ = nw345_
            d_2044_v19_ = d_2044_v19_
            d_2020_v0_ = (d_2020_v0_) - (default__.safeModuloInt(d_2035_cf5_, d_2035_cf5_))
            d_2045_v20_: _dafny.Set
            d_2045_v20_ = _dafny.Set({d_2034_cf6_, d_2034_cf6_, True, False, p0})
            d_2046_v21_: C0
            nw346_ = C0()
            nw346_.ctor__(d_2045_v20_)
            d_2046_v21_ = nw346_
        elif source26_.is_DC5:
            d_2047___mcc_h2_ = source26_.cf7
            d_2048___mcc_h3_ = source26_.cf8
            d_2049___mcc_h4_ = source26_.cf9
            d_2050___mcc_h5_ = source26_.cf10
            d_2051___mcc_h6_ = source26_.cf11
            d_2052_cf11_ = d_2051___mcc_h6_
            d_2053_cf10_ = d_2050___mcc_h5_
            d_2054_cf9_ = d_2049___mcc_h4_
            d_2055_cf8_ = d_2048___mcc_h3_
            d_2056_cf7_ = d_2047___mcc_h2_
            d_2057_v24_: _dafny.Set
            def iife182_():
                coll80_ = _dafny.Set()
                compr_80_: int
                for compr_80_ in _dafny.IntegerRange(-136, 176):
                    d_2058_v22_: int = compr_80_
                    if ((-136) <= (d_2058_v22_)) and ((d_2058_v22_) < (176)):
                        def iife183_():
                            coll81_ = _dafny.Map()
                            compr_81_: int
                            for compr_81_ in _dafny.IntegerRange(-340, 432):
                                d_2059_v23_: int = compr_81_
                                if ((-340) <= (d_2059_v23_)) and ((d_2059_v23_) < (432)):
                                    coll81_[default__.safeDivisionInt(d_2059_v23_, len(d_2025_v5_))] = d_2020_v0_
                            return _dafny.Map(coll81_)
                        coll80_ = coll80_.union(_dafny.Set([(d_2058_v22_) - (len(iife183_()
))]))
                return _dafny.Set(coll80_)
            d_2057_v24_ = _dafny.Set({len(iife182_()
            )})
            d_2060_v25_: _dafny.Map
            d_2060_v25_ = _dafny.Map({d_2057_v24_: d_2053_cf10_})
            d_2061_v26_: _dafny.Seq
            d_2061_v26_ = _dafny.SeqWithoutIsStrInference([(d_2060_v25_).set(d_2057_v24_, d_2055_cf8_), d_2060_v25_, d_2060_v25_, _dafny.Map({d_2057_v24_: d_2053_cf10_}), d_2060_v25_])
            d_2062_v28_: _dafny.Set
            d_2062_v28_ = _dafny.Set({_dafny.Set({d_2054_cf9_, len(d_2026_v6_), d_2020_v0_})})
            def iife184_():
                coll82_ = _dafny.Set()
                compr_82_: _dafny.Set
                for compr_82_ in ((d_2061_v26_)[default__.safeIndex(d_2054_cf9_, len(d_2061_v26_))]).keys.Elements:
                    d_2063_v27_: _dafny.Set = compr_82_
                    if (d_2063_v27_) in ((d_2061_v26_)[default__.safeIndex(d_2054_cf9_, len(d_2061_v26_))]):
                        coll82_ = coll82_.union(_dafny.Set([d_2063_v27_]))
                return _dafny.Set(coll82_)
            d_2056_cf7_ = ((d_2062_v28_).intersection(d_2062_v28_)).issubset(iife184_()
            )
            d_2064_v29_: _dafny.Array
            nw347_ = _dafny.Array(None, 1)
            nw347_[int(0)] = d_2020_v0_
            d_2064_v29_ = nw347_
            d_2064_v29_ = d_2064_v29_
            (globalState).f1 = (default__.safeModuloInt(d_2020_v0_, d_2054_cf9_)) >= (default__.safeDivisionInt(d_2020_v0_, d_2054_cf9_))
            d_2065_v30_: _dafny.Array
            def lambda82_(d_2066_v5_):
                def lambda83_(d_2067_i1_):
                    return (D4_DC11(d_2066_v5_) if True else D4_DC11(d_2066_v5_))

                return lambda83_

            init44_ = lambda82_(d_2025_v5_)
            nw348_ = _dafny.Array(None, 8)
            for i0_44_ in range(nw348_.length(0)):
                nw348_[i0_44_] = init44_(i0_44_)
            d_2065_v30_ = nw348_
            d_2068_v31_: _dafny.Seq
            d_2068_v31_ = _dafny.SeqWithoutIsStrInference([d_2025_v5_, d_2025_v5_, d_2025_v5_, d_2025_v5_, default__.fm0(default__.fm4(globalState), globalState)])
            d_2069_v32_: D4
            d_2069_v32_ = D4_DC11(((d_2068_v31_)[default__.safeIndex(d_2020_v0_, len(d_2068_v31_))]).set(default__.safeIndex(d_2054_cf9_, len((d_2068_v31_)[default__.safeIndex(d_2020_v0_, len(d_2068_v31_))])), True))
            index341_ = default__.safeIndex(622, (d_2065_v30_).length(0))
            (d_2065_v30_)[index341_] = d_2069_v32_
            index342_ = default__.safeIndex(622, (d_2065_v30_).length(0))
            (d_2065_v30_)[index342_] = d_2069_v32_
        elif True:
            d_2070___mcc_h7_ = source26_.cf4
            d_2071_cf4_ = d_2070___mcc_h7_
            d_2072_v33_: str
            d_2072_v33_ = _dafny.CodePoint('b')
            d_2072_v33_ = p1
            d_2073_v34_: _dafny.Set
            d_2073_v34_ = _dafny.Set({True})
            d_2074_v35_: _dafny.Map
            d_2074_v35_ = _dafny.Map({p0: p0})
            d_2075_v36_: _dafny.Map
            d_2075_v36_ = _dafny.Map({((d_2074_v35_)[(self).fm7(d_2071_cf4_, d_2071_cf4_, globalState)] if ((self).fm7(d_2071_cf4_, d_2071_cf4_, globalState)) in (d_2074_v35_) else d_2071_cf4_): d_2072_v33_})
            d_2076_v38_: _dafny.MultiSet
            d_2076_v38_ = _dafny.MultiSet([d_2027_v7_, d_2027_v7_])
            d_2077_v39_: _dafny.Map
            d_2077_v39_ = _dafny.Map({d_2020_v0_: d_2020_v0_})
            d_2078_v40_: _dafny.Map
            d_2078_v40_ = _dafny.Map({d_2020_v0_: d_2077_v39_})
            d_2079_v41_: _dafny.Map
            d_2079_v41_ = _dafny.Map({False: len(((d_2078_v40_)[len(d_2025_v5_)] if (len(d_2025_v5_)) in (d_2078_v40_) else d_2077_v39_))})
            d_2080_v42_: _dafny.Map
            d_2080_v42_ = _dafny.Map({d_2020_v0_: d_2071_cf4_})
            d_2081_v43_: _dafny.Map
            d_2081_v43_ = _dafny.Map({d_2080_v42_: d_2071_cf4_})
            d_2082_v44_: _dafny.Array
            nw349_ = _dafny.Array(None, 24)
            nw349_[int(0)] = (default__.fm2(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_2083_i2_ in range(default__.abs(-871))]), d_2073_v34_, d_2075_v36_, d_2071_cf4_, globalState)) != (len(_dafny.Map({p0: d_2071_cf4_})))
            nw349_[int(1)] = (d_2071_cf4_) == (d_2071_cf4_)
            nw349_[int(2)] = (_dafny.SeqWithoutIsStrInference([d_2072_v33_ for d_2084_i3_ in range(default__.abs(-1))])) not in (default__.fm40(d_2071_cf4_, globalState))
            nw349_[int(3)] = d_2071_cf4_
            nw349_[int(4)] = d_2071_cf4_
            nw349_[int(5)] = p0
            def iife185_():
                coll83_ = _dafny.Map()
                compr_83_: _dafny.Seq
                for compr_83_ in (d_2076_v38_).Elements:
                    d_2085_v37_: _dafny.Seq = compr_83_
                    if (d_2085_v37_) in (d_2076_v38_):
                        coll83_[d_2085_v37_] = p0
                return _dafny.Map(coll83_)
            nw349_[int(6)] = (default__.fm30(len(d_2027_v7_), iife185_()
            , globalState)) == (d_2079_v41_)
            nw349_[int(7)] = p0
            nw349_[int(8)] = (d_2020_v0_) <= (len(d_2027_v7_))
            nw349_[int(9)] = not (d_2071_cf4_) or (d_2071_cf4_)
            nw349_[int(10)] = not(p0)
            nw349_[int(11)] = p0
            nw349_[int(12)] = p0
            nw349_[int(13)] = default__.fm4(globalState)
            nw349_[int(14)] = p0
            nw349_[int(15)] = p0
            nw349_[int(16)] = False
            nw349_[int(17)] = (len(d_2027_v7_)) >= (d_2020_v0_)
            nw349_[int(18)] = d_2071_cf4_
            nw349_[int(19)] = ((d_2081_v43_)[(_dafny.Map({d_2020_v0_: d_2071_cf4_})).set(d_2020_v0_, p0)] if ((_dafny.Map({d_2020_v0_: d_2071_cf4_})).set(d_2020_v0_, p0)) in (d_2081_v43_) else p0)
            nw349_[int(20)] = d_2071_cf4_
            nw349_[int(21)] = False
            nw349_[int(22)] = p0
            nw349_[int(23)] = (d_2020_v0_) < (d_2020_v0_)
            d_2082_v44_ = nw349_
            index343_ = default__.safeIndex(691, (d_2082_v44_).length(0))
            (d_2082_v44_)[index343_] = d_2071_cf4_
            index344_ = default__.safeIndex(691, (d_2082_v44_).length(0))
            (d_2082_v44_)[index344_] = d_2071_cf4_
            d_2080_v42_ = (d_2080_v42_).set(d_2020_v0_, (self).fm7(p0, p0, globalState))
            d_2086_v45_: D4
            d_2086_v45_ = D4_DC12(d_2071_cf4_, default__.fm41(d_2020_v0_, globalState))
            d_2087_v46_: _dafny.Array
            nw350_ = _dafny.Array(None, 13)
            d_2087_v46_ = nw350_
            d_2088_v47_: _dafny.Map
            d_2088_v47_ = _dafny.Map({d_2086_v45_: d_2087_v46_})
            d_2089_v48_: C4
            nw351_ = C4()
            nw351_.ctor__(not((d_2082_v44_)[default__.safeIndex(691, (d_2082_v44_).length(0))]), d_2029_v9_)
            d_2089_v48_ = nw351_
            d_2090_v49_: _dafny.Map
            d_2090_v49_ = _dafny.Map({d_2089_v48_: d_2088_v47_})
            d_2088_v47_ = ((d_2090_v49_)[d_2089_v48_] if (d_2089_v48_) in (d_2090_v49_) else _dafny.Map({d_2086_v45_: d_2087_v46_}))
        d_2091_v50_: _dafny.Map
        d_2091_v50_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0]): d_2020_v0_})
        d_2092_v51_: _dafny.Set
        d_2092_v51_ = _dafny.Set({p0, p0, default__.fm4(globalState)})
        d_2093_v52_: C2
        nw352_ = C2()
        nw352_.ctor__(d_2091_v50_, d_2092_v51_, (self).f28)
        d_2093_v52_ = nw352_
        r0 = (d_2020_v0_) - (d_2020_v0_)
        (globalState).f16 = d_2025_v5_
        r0 = len((d_2027_v7_) + (_dafny.SeqWithoutIsStrInference([p1 for d_2094_i4_ in range(default__.abs(320))])))
        return r0

    def m8(self, p0, p1, p2, p3, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Map = _dafny.Map({})
        r3: _dafny.Set = _dafny.Set({})
        (globalState).f19 = p0
        d_2095_v0_: _dafny.Seq
        d_2095_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))
        d_2096_v1_: _dafny.Seq
        d_2096_v1_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2097_v2_: _dafny.Seq
        d_2097_v2_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2098_v3_: _dafny.MultiSet
        d_2098_v3_ = _dafny.MultiSet([default__.fm24(d_2097_v2_, p1, globalState)])
        d_2099_v4_: _dafny.MultiSet
        d_2099_v4_ = d_2098_v3_
        d_2100_v6_: _dafny.Map
        d_2100_v6_ = _dafny.Map({p1: p1})
        d_2101_v7_: D8
        d_2101_v7_ = D8_DC19(d_2100_v6_)
        d_2102_v8_: _dafny.Seq
        d_2102_v8_ = _dafny.SeqWithoutIsStrInference([d_2101_v7_])
        d_2103_v9_: _dafny.Set
        d_2103_v9_ = _dafny.Set({p0, len(d_2102_v8_)})
        d_2104_v10_: C1
        nw353_ = C1()
        def iife186_():
            coll84_ = _dafny.Set()
            compr_84_: _dafny.Set
            for compr_84_ in ((d_2099_v4_)).Elements:
                d_2105_v5_: _dafny.Set = compr_84_
                if (d_2105_v5_) in ((d_2099_v4_)):
                    coll84_ = coll84_.union(_dafny.Set([d_2105_v5_]))
            return _dafny.Set(coll84_)
        nw353_.ctor__((default__.fm2(d_2095_v0_, _dafny.Set({p1}), _dafny.Map({p1: _dafny.CodePoint('p')}), not(False), globalState)) + ((_dafny.MultiSet(d_2096_v1_)).cardinality), default__.safeDivisionInt(len(iife186_()
        ), p0), default__.fm20(p1, p1, len(d_2096_v1_), d_2103_v9_, globalState))
        d_2104_v10_ = nw353_
        hi11_ = d_2104_v10_.f42
        for d_2106_i0_ in range(d_2104_v10_.f42, hi11_):
            d_2100_v6_ = d_2100_v6_
            d_2107_v11_: D4
            d_2107_v11_ = D4_DC11(d_2097_v2_)
            d_2108_v12_: _dafny.MultiSet
            d_2108_v12_ = _dafny.MultiSet([len((d_2107_v11_).cf20), (d_2104_v10_).f41])
            (globalState).f27 = not((d_2108_v12_).ispropersubset(d_2108_v12_))
            d_2095_v0_ = ((d_2095_v0_) + (d_2095_v0_)) + (d_2095_v0_)
            (globalState).f19 = d_2104_v10_.f42
        d_2109_v13_: _dafny.Seq
        d_2109_v13_ = _dafny.SeqWithoutIsStrInference([p3, p3, p3, p3])
        d_2110_v14_: D6
        d_2110_v14_ = D6_DC15(d_2109_v13_)
        d_2111_v15_: _dafny.Map
        d_2111_v15_ = _dafny.Map({d_2104_v10_.f42: p1})
        d_2112_v16_: _dafny.Map
        d_2112_v16_ = _dafny.Map({(0) - ((d_2104_v10_).f41): p0})
        d_2113_v17_: _dafny.Map
        d_2113_v17_ = _dafny.Map({d_2112_v16_: p1})
        d_2114_v18_: _dafny.Seq
        d_2114_v18_ = _dafny.SeqWithoutIsStrInference([d_2095_v0_, d_2095_v0_])
        d_2115_v19_: _dafny.Map
        d_2115_v19_ = _dafny.Map({d_2097_v2_: len(d_2114_v18_)})
        d_2116_v20_: D8
        d_2116_v20_ = D8_DC21(p1, p1, (d_2104_v10_).f41)
        d_2117_v21_: _dafny.Array
        nw354_ = _dafny.Array(None, 24)
        nw354_[int(0)] = not (p1) or (p1)
        nw354_[int(1)] = p1
        nw354_[int(2)] = p1
        nw354_[int(3)] = not (p1) or ((p3).fm7(p1, p1, globalState))
        nw354_[int(4)] = p1
        nw354_[int(5)] = (p0) <= (d_2104_v10_.f42)
        nw354_[int(6)] = p1
        nw354_[int(7)] = not((p1) == (((d_2111_v15_)[535] if (535) in (d_2111_v15_) else p1)))
        nw354_[int(8)] = p1
        nw354_[int(9)] = (d_2113_v17_) == (d_2113_v17_)
        nw354_[int(10)] = p1
        nw354_[int(11)] = (d_2097_v2_) not in (d_2115_v19_)
        nw354_[int(12)] = p1
        nw354_[int(13)] = ((d_2100_v6_)[p1] if (p1) in (d_2100_v6_) else p1)
        nw354_[int(14)] = p1
        nw354_[int(15)] = (p1) and ((self).fm7(p1, (d_2116_v20_).cf33, globalState))
        nw354_[int(16)] = p1
        nw354_[int(17)] = not(p1)
        nw354_[int(18)] = p1
        nw354_[int(19)] = not((d_2095_v0_) == (d_2095_v0_))
        nw354_[int(20)] = p1
        nw354_[int(21)] = p1
        nw354_[int(22)] = p1
        nw354_[int(23)] = p1
        d_2117_v21_ = nw354_
        d_2118_v22_: D1
        d_2118_v22_ = D1_DC5(p1, not(p1), p2, p1, p1)
        d_2119_v23_: _dafny.Map
        d_2119_v23_ = _dafny.Map({p0: d_2118_v22_})
        pat_let_tv42_ = p2
        pat_let_tv43_ = p1
        d_2120_v24_: _dafny.Set
        def iife187_(_pat_let51_0):
            def iife188_(d_2121_dt__update__tmp_h0_):
                def iife189_(_pat_let52_0):
                    def iife190_(d_2122_dt__update_hcf9_h0_):
                        def iife191_(_pat_let53_0):
                            def iife192_(d_2123_dt__update_hcf8_h0_):
                                return D1_DC5((d_2121_dt__update__tmp_h0_).cf7, d_2123_dt__update_hcf8_h0_, d_2122_dt__update_hcf9_h0_, (d_2121_dt__update__tmp_h0_).cf10, (d_2121_dt__update__tmp_h0_).cf11)
                            return iife192_(_pat_let53_0)
                        return iife191_(pat_let_tv43_)
                    return iife190_(_pat_let52_0)
                return iife189_(pat_let_tv42_)
            return iife188_(_pat_let51_0)
        d_2120_v24_ = _dafny.Set({d_2119_v23_, d_2119_v23_, _dafny.Map({d_2104_v10_.f42: iife187_(d_2118_v22_)}), d_2119_v23_, default__.fm42(globalState)})
        d_2124_v25_: D21
        d_2124_v25_ = D21_DC53(d_2120_v24_)
        pat_let_tv44_ = p2
        pat_let_tv45_ = d_2104_v10_
        pat_let_tv46_ = d_2095_v0_
        pat_let_tv47_ = d_2104_v10_
        def lambda84_(source27_):
            if source27_.is_DC52:
                d_2125___mcc_h0_ = source27_.cf97
                d_2126___mcc_h1_ = source27_.cf98
                d_2127___mcc_h2_ = source27_.cf99
                d_2128___mcc_h3_ = source27_.cf100
                d_2129___mcc_h4_ = source27_.cf101
                d_2130_cf101_ = d_2129___mcc_h4_
                d_2131_cf100_ = d_2128___mcc_h3_
                d_2132_cf99_ = d_2127___mcc_h2_
                d_2133_cf98_ = d_2126___mcc_h1_
                d_2134_cf97_ = d_2125___mcc_h0_
                return pat_let_tv44_
            elif source27_.is_DC53:
                d_2135___mcc_h5_ = source27_.cf102
                d_2136_cf102_ = d_2135___mcc_h5_
                return pat_let_tv45_.f42
            elif source27_.is_DC51:
                d_2137___mcc_h6_ = source27_.cf96
                d_2138_cf96_ = d_2137___mcc_h6_
                return len(pat_let_tv46_)
            elif True:
                d_2139___mcc_h7_ = source27_.cf103
                d_2140_cf103_ = d_2139___mcc_h7_
                return (pat_let_tv47_).f41

        rhs386_ = d_2110_v14_
        rhs387_ = d_2117_v21_
        rhs388_ = p1
        rhs389_ = (d_2104_v10_).f41
        rhs390_ = lambda84_(d_2124_v25_)
        lhs317_ = globalState
        lhs318_ = globalState
        lhs319_ = globalState
        d_2110_v14_ = rhs386_
        d_2117_v21_ = rhs387_
        lhs317_.f11 = rhs388_
        lhs318_.f19 = rhs389_
        lhs319_.f14 = rhs390_
        if (d_2097_v2_)[default__.safeIndex(p0, len(d_2097_v2_))]:
            (d_2104_v10_).f42 = 480
            d_2112_v16_ = (d_2112_v16_).set(default__.safeDivisionInt(d_2104_v10_.f42, 209), p0)
            (globalState).f27 = not(p1)
            (globalState).f14 = p0
            d_2141_v26_: _dafny.Array
            def lambda85_(d_2142_v10_):
                def lambda86_(d_2143_i1_):
                    return (_dafny.MultiSet([(d_2142_v10_).f41])).intersection(_dafny.MultiSet([(d_2142_v10_).f41]))

                return lambda86_

            init45_ = lambda85_(d_2104_v10_)
            nw355_ = _dafny.Array(None, 11)
            for i0_45_ in range(nw355_.length(0)):
                nw355_[i0_45_] = init45_(i0_45_)
            d_2141_v26_ = nw355_
            index345_ = default__.safeIndex(213, (d_2141_v26_).length(0))
            (d_2141_v26_)[index345_] = _dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2144_i2_ in range(default__.abs(385))]))])
            d_2145_v27_: D12
            d_2145_v27_ = D12_DC29(p1)
            d_2146_v28_: _dafny.MultiSet
            d_2146_v28_ = _dafny.MultiSet([(d_2145_v27_).cf50])
            d_2147_v29_: _dafny.MultiSet
            d_2147_v29_ = _dafny.MultiSet([p2])
            d_2148_v30_: D18
            d_2148_v30_ = D18_DC41(d_2096_v1_)
            d_2149_v31_: str
            d_2149_v31_ = _dafny.CodePoint('n')
            d_2150_v32_: _dafny.Map
            d_2150_v32_ = _dafny.Map({d_2146_v28_: p0})
            d_2151_v33_: _dafny.Map
            d_2151_v33_ = _dafny.Map({d_2096_v1_: p1})
            d_2152_v34_: _dafny.Map
            d_2152_v34_ = _dafny.Map({d_2096_v1_: d_2117_v21_})
            d_2153_v35_: _dafny.Set
            d_2153_v35_ = _dafny.Set({p1})
            d_2154_v36_: _dafny.Map
            d_2154_v36_ = _dafny.Map({p1: d_2149_v31_})
            d_2155_v37_: _dafny.Map
            d_2155_v37_ = _dafny.Map({d_2149_v31_: d_2117_v21_})
            d_2156_v38_: D12
            d_2156_v38_ = D12_DC27(d_2155_v37_)
            d_2157_v39_: D18
            d_2157_v39_ = D18_DC43(default__.fm2(d_2095_v0_, d_2153_v35_, d_2154_v36_, p1, globalState), 896, d_2095_v0_, d_2156_v38_, p1)
            index346_ = default__.safeIndex(213, (d_2141_v26_).length(0))
            rhs391_ = d_2147_v29_
            rhs392_ = ((d_2148_v30_).cf71).set(default__.safeIndex(default__.safeDivisionInt(len((d_2095_v0_).set(default__.safeIndex(280, len(d_2095_v0_)), d_2149_v31_)), 380), len((d_2148_v30_).cf71)), ((d_2150_v32_)[d_2146_v28_] if (d_2146_v28_) in (d_2150_v32_) else len(d_2151_v33_)))
            rhs393_ = d_2146_v28_
            rhs394_ = (default__.fm2(d_2095_v0_, d_2153_v35_, d_2154_v36_, p1, globalState) if (len((d_2152_v34_).set(_dafny.SeqWithoutIsStrInference([d_2104_v10_.f42]), d_2117_v21_))) >= ((d_2157_v39_).cf77) else (0) - (p2))
            lhs320_ = d_2141_v26_
            lhs321_ = default__.safeIndex(213, (d_2141_v26_).length(0))
            lhs322_ = globalState
            lhs320_[lhs321_] = rhs391_
            d_2096_v1_ = rhs392_
            d_2146_v28_ = rhs393_
            lhs322_.f24 = rhs394_
        elif True:
            (d_2104_v10_).f42 = (d_2104_v10_).f41
            d_2158_v40_: _dafny.MultiSet
            d_2158_v40_ = _dafny.MultiSet([p1, p1])
            d_2159_v41_: _dafny.MultiSet
            d_2159_v41_ = _dafny.MultiSet([((d_2158_v40_)[p1] if (p1) in (d_2158_v40_) else 184)])
            d_2160_v43_: _dafny.MultiSet
            d_2160_v43_ = _dafny.MultiSet([d_2096_v1_, d_2096_v1_])
            d_2161_v44_: _dafny.Array
            nw356_ = _dafny.Array(None, 20)
            nw356_[int(0)] = (d_2104_v10_).f41
            nw356_[int(1)] = (0) - ((d_2104_v10_).f41)
            nw356_[int(2)] = (0) - ((d_2104_v10_).f41)
            nw356_[int(3)] = -390
            nw356_[int(4)] = len(d_2097_v2_)
            nw356_[int(5)] = (d_2096_v1_)[default__.safeIndex(p2, len(d_2096_v1_))]
            nw356_[int(6)] = (d_2104_v10_).f41
            nw356_[int(7)] = len(d_2095_v0_)
            nw356_[int(8)] = d_2104_v10_.f42
            nw356_[int(9)] = (d_2104_v10_).f41
            nw356_[int(10)] = d_2104_v10_.f42
            nw356_[int(11)] = p2
            nw356_[int(12)] = ((d_2159_v41_)[p0] if (p0) in (d_2159_v41_) else (d_2159_v41_).cardinality)
            nw356_[int(13)] = -628
            nw356_[int(14)] = (0) - ((d_2104_v10_).f41)
            nw356_[int(15)] = d_2104_v10_.f42
            nw356_[int(16)] = len(d_2095_v0_)
            def iife193_():
                coll85_ = _dafny.Map()
                compr_85_: _dafny.Seq
                for compr_85_ in (d_2160_v43_).Elements:
                    d_2162_v42_: _dafny.Seq = compr_85_
                    if (d_2162_v42_) in (d_2160_v43_):
                        coll85_[d_2162_v42_] = p1
                return _dafny.Map(coll85_)
            nw356_[int(17)] = len(iife193_()
            )
            nw356_[int(18)] = (d_2096_v1_)[default__.safeIndex((self).fm15(p1, (d_2104_v10_).f41, 281, globalState), len(d_2096_v1_))]
            nw356_[int(19)] = (_dafny.MultiSet(d_2096_v1_)).cardinality
            d_2161_v44_ = nw356_
            d_2163_v45_: _dafny.MultiSet
            d_2163_v45_ = _dafny.MultiSet([d_2161_v44_, d_2161_v44_, d_2161_v44_])
            r1 = ((d_2163_v45_)[d_2161_v44_] if (d_2161_v44_) in (d_2163_v45_) else 836)
            (globalState).f11 = p1
            d_2164_v46_: _dafny.Map
            d_2164_v46_ = _dafny.Map({d_2095_v0_: d_2096_v1_})
            (d_2104_v10_).f42 = len((d_2164_v46_) | (d_2164_v46_))
            (globalState).f11 = ((d_2104_v10_).f41) != ((d_2104_v10_).f41)
        d_2165_v47_: str
        d_2165_v47_ = _dafny.CodePoint('t')
        d_2166_v48_: D2
        d_2166_v48_ = D2_DC7(p1, p1, _dafny.SeqWithoutIsStrInference([d_2165_v47_, d_2165_v47_]))
        d_2167_v49_: _dafny.Map
        d_2167_v49_ = _dafny.Map({(d_2104_v10_).f41: d_2165_v47_})
        d_2168_v50_: _dafny.Map
        d_2168_v50_ = _dafny.Map({p1: p0})
        d_2169_v51_: _dafny.Array
        nw357_ = _dafny.Array(None, 25)
        nw357_[int(0)] = d_2095_v0_
        nw357_[int(1)] = (d_2095_v0_).set(default__.safeIndex(p2, len(d_2095_v0_)), d_2165_v47_)
        nw357_[int(2)] = d_2095_v0_
        nw357_[int(3)] = d_2095_v0_
        nw357_[int(4)] = d_2095_v0_
        nw357_[int(5)] = d_2095_v0_
        nw357_[int(6)] = (d_2166_v48_).cf15
        nw357_[int(7)] = d_2095_v0_
        nw357_[int(8)] = d_2095_v0_
        nw357_[int(9)] = d_2095_v0_
        nw357_[int(10)] = d_2095_v0_
        nw357_[int(11)] = (d_2095_v0_).set(default__.safeIndex(len(d_2095_v0_), len(d_2095_v0_)), _dafny.CodePoint('l'))
        nw357_[int(12)] = d_2095_v0_
        nw357_[int(13)] = d_2095_v0_
        nw357_[int(14)] = d_2095_v0_
        nw357_[int(15)] = _dafny.SeqWithoutIsStrInference([d_2165_v47_ for d_2170_i3_ in range(default__.abs(-501))])
        nw357_[int(16)] = d_2095_v0_
        nw357_[int(17)] = d_2095_v0_
        nw357_[int(18)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qnmyost"))
        nw357_[int(19)] = d_2095_v0_
        nw357_[int(20)] = (d_2095_v0_).set(default__.safeIndex(p2, len(d_2095_v0_)), ((d_2167_v49_)[p0] if (p0) in (d_2167_v49_) else d_2165_v47_))
        nw357_[int(21)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vahiro"))
        nw357_[int(22)] = d_2095_v0_
        nw357_[int(23)] = d_2095_v0_
        nw357_[int(24)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvmyitwy"))).set(default__.safeIndex(len(d_2168_v50_), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvmyitwy")))), d_2165_v47_)
        d_2169_v51_ = nw357_
        d_2171_v52_: D22
        d_2171_v52_ = D22_DC55(d_2169_v51_)
        source28_ = d_2171_v52_
        if source28_.is_DC56:
            d_2172___mcc_h8_ = source28_.cf105
            d_2173___mcc_h9_ = source28_.cf106
            d_2174___mcc_h10_ = source28_.cf107
            d_2175___mcc_h11_ = source28_.cf108
            d_2176_cf108_ = d_2175___mcc_h11_
            d_2177_cf107_ = d_2174___mcc_h10_
            d_2178_cf106_ = d_2173___mcc_h9_
            d_2179_cf105_ = d_2172___mcc_h8_
            d_2180_v53_: _dafny.Array
            def lambda87_(d_2181_cf106_):
                def lambda88_(d_2182_i4_):
                    return (d_2182_i4_) * (d_2181_cf106_)

                return lambda88_

            init46_ = lambda87_(d_2178_cf106_)
            nw358_ = _dafny.Array(None, 13)
            for i0_46_ in range(nw358_.length(0)):
                nw358_[i0_46_] = init46_(i0_46_)
            d_2180_v53_ = nw358_
            d_2183_v54_: _dafny.Map
            d_2183_v54_ = _dafny.Map({_dafny.Map({-361: d_2180_v53_}): (d_2104_v10_).f41})
            d_2184_v55_: _dafny.Map
            d_2184_v55_ = _dafny.Map({(d_2104_v10_).f41: d_2180_v53_})
            d_2185_v56_: _dafny.Seq
            d_2185_v56_ = _dafny.SeqWithoutIsStrInference([d_2184_v55_, _dafny.Map({d_2104_v10_.f42: d_2180_v53_}), d_2184_v55_])
            (globalState).f14 = (self).fm15((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hliohr"))) < (d_2179_cf105_), d_2104_v10_.f42, ((d_2183_v54_)[(d_2185_v56_)[default__.safeIndex(d_2104_v10_.f42, len(d_2185_v56_))]] if ((d_2185_v56_)[default__.safeIndex(d_2104_v10_.f42, len(d_2185_v56_))]) in (d_2183_v54_) else len(d_2095_v0_)), globalState)
            d_2186_v57_: _dafny.Set
            d_2186_v57_ = _dafny.Set({d_2180_v53_})
            d_2186_v57_ = _dafny.Set({d_2180_v53_, d_2180_v53_, d_2180_v53_, d_2180_v53_, d_2180_v53_})
            d_2187_v58_: _dafny.Array
            nw359_ = _dafny.Array(_dafny.Map({}), 22)
            d_2187_v58_ = nw359_
            d_2188_v59_: _dafny.Map
            d_2188_v59_ = _dafny.Map({d_2180_v53_: d_2117_v21_})
            index347_ = default__.safeIndex(507, (d_2187_v58_).length(0))
            (d_2187_v58_)[index347_] = d_2188_v59_
            index348_ = default__.safeIndex(507, (d_2187_v58_).length(0))
            (d_2187_v58_)[index348_] = d_2188_v59_
            d_2189_v60_: _dafny.Map
            d_2189_v60_ = _dafny.Map({(d_2097_v2_)[default__.safeIndex(len(d_2179_cf105_), len(d_2097_v2_))]: d_2180_v53_})
            d_2189_v60_ = ((_dafny.Map({d_2177_cf107_: d_2180_v53_})).set(d_2177_cf107_, d_2180_v53_)) | (_dafny.Map({d_2176_cf108_: d_2180_v53_}))
        elif source28_.is_DC57:
            d_2190___mcc_h12_ = source28_.cf109
            d_2191___mcc_h13_ = source28_.cf110
            d_2192___mcc_h14_ = source28_.cf111
            d_2193___mcc_h15_ = source28_.cf112
            d_2194_cf112_ = d_2193___mcc_h15_
            d_2195_cf111_ = d_2192___mcc_h14_
            d_2196_cf110_ = d_2191___mcc_h13_
            d_2197_cf109_ = d_2190___mcc_h12_
            (globalState).f19 = 113
            (globalState).f14 = (d_2096_v1_)[default__.safeIndex((d_2104_v10_).f41, len(d_2096_v1_))]
            d_2198_v61_: _dafny.Seq
            d_2198_v61_ = _dafny.SeqWithoutIsStrInference([d_2098_v3_])
            d_2199_v62_: _dafny.Array
            nw360_ = _dafny.Array(None, 2)
            nw360_[int(0)] = d_2099_v4_
            nw360_[int(1)] = (d_2198_v61_)[default__.safeIndex((d_2104_v10_).f41, len(d_2198_v61_))]
            d_2199_v62_ = nw360_
            index349_ = default__.safeIndex(700, (d_2199_v62_).length(0))
            (d_2199_v62_)[index349_] = d_2099_v4_
            d_2200_v63_: D12
            d_2200_v63_ = D12_DC29(d_2194_cf112_)
            d_2201_v64_: _dafny.Set
            d_2201_v64_ = _dafny.Set({p1, d_2194_cf112_})
            index350_ = default__.safeIndex(700, (d_2199_v62_).length(0))
            rhs395_ = d_2099_v4_
            rhs396_ = (d_2104_v10_).fm8(default__.fm21(d_2165_v47_, globalState), p1, not (p1) or (d_2194_cf112_), globalState)
            rhs397_ = default__.fm2((d_2095_v0_) + (d_2095_v0_), d_2201_v64_, _dafny.Map({d_2194_cf112_: d_2197_cf109_}), not(((d_2111_v15_)[d_2196_cf110_] if (d_2196_cf110_) in (d_2111_v15_) else p1)), globalState)
            rhs398_ = d_2200_v63_
            lhs323_ = d_2199_v62_
            lhs324_ = default__.safeIndex(700, (d_2199_v62_).length(0))
            lhs325_ = globalState
            lhs323_[lhs324_] = rhs395_
            d_2165_v47_ = rhs396_
            lhs325_.f19 = rhs397_
            d_2200_v63_ = rhs398_
            d_2202_v65_: _dafny.Map
            d_2202_v65_ = _dafny.Map({not(p1): (d_2095_v0_) + (d_2095_v0_)})
            d_2203_v66_: D19
            d_2203_v66_ = D19_DC45(p2, not(d_2194_cf112_))
            pat_let_tv48_ = d_2194_cf112_
            d_2204_v67_: _dafny.Map
            def iife194_(_pat_let54_0):
                def iife195_(d_2205_dt__update__tmp_h1_):
                    def iife196_(_pat_let55_0):
                        def iife197_(d_2206_dt__update_hcf83_h0_):
                            return D19_DC45((d_2205_dt__update__tmp_h1_).cf82, d_2206_dt__update_hcf83_h0_)
                        return iife197_(_pat_let55_0)
                    return iife196_(not(pat_let_tv48_))
                return iife195_(_pat_let54_0)
            d_2204_v67_ = _dafny.Map({iife194_(d_2203_v66_): d_2117_v21_})
            d_2202_v65_ = (d_2202_v65_).set((D19_DC45(d_2104_v10_.f42, False)) in (d_2204_v67_), _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2207_i5_ in range(default__.abs(187))]))
        elif source28_.is_DC55:
            d_2208___mcc_h16_ = source28_.cf104
            d_2209_cf104_ = d_2208___mcc_h16_
            d_2210_v68_: _dafny.Array
            nw361_ = _dafny.Array(D0.default()(), 17)
            d_2210_v68_ = nw361_
            d_2211_v69_: D0
            d_2211_v69_ = D0_DC1(len(d_2168_v50_))
            index351_ = default__.safeIndex(999, (d_2210_v68_).length(0))
            (d_2210_v68_)[index351_] = d_2211_v69_
            index352_ = default__.safeIndex(999, (d_2210_v68_).length(0))
            (d_2210_v68_)[index352_] = default__.fm43(globalState)
            d_2212_v70_: _dafny.Set
            d_2212_v70_ = _dafny.Set({p1, p1, p1, p1})
            d_2213_v71_: _dafny.Map
            d_2213_v71_ = _dafny.Map({False: d_2165_v47_})
            d_2214_v72_: _dafny.MultiSet
            d_2214_v72_ = _dafny.MultiSet([p1, p1])
            d_2215_v73_: _dafny.MultiSet
            d_2215_v73_ = _dafny.MultiSet([_dafny.MultiSet(d_2097_v2_), d_2214_v72_])
            rhs399_ = (p0) > (default__.safeModuloInt(default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmqxejnad")), d_2212_v70_, d_2213_v71_, p1, globalState), d_2104_v10_.f42))
            rhs400_ = (d_2215_v73_) != (d_2215_v73_)
            lhs326_ = globalState
            lhs327_ = globalState
            lhs326_.f11 = rhs399_
            lhs327_.f1 = rhs400_
            if (d_2117_v21_) not in (_dafny.SeqWithoutIsStrInference([d_2117_v21_, d_2117_v21_])):
                (globalState).f27 = (d_2095_v0_) != (d_2095_v0_)
                (globalState).f27 = not(p1)
                (globalState).f0 = p0
                d_2216_v74_: _dafny.Map
                d_2216_v74_ = _dafny.Map({(d_2214_v72_ if p1 else _dafny.MultiSet([p1])): (len(d_2096_v1_)) - (((d_2214_v72_)[p1] if (p1) in (d_2214_v72_) else d_2104_v10_.f42))})
                d_2216_v74_ = (d_2216_v74_).set(d_2214_v72_, ((d_2168_v50_)[p1] if (p1) in (d_2168_v50_) else p2))
                d_2217_v75_: _dafny.Array
                def lambda89_(d_2218_v0_):
                    def lambda90_(d_2219_i6_):
                        return (d_2219_i6_) * (len(d_2218_v0_))

                    return lambda90_

                init47_ = lambda89_(d_2095_v0_)
                nw362_ = _dafny.Array(None, 9)
                for i0_47_ in range(nw362_.length(0)):
                    nw362_[i0_47_] = init47_(i0_47_)
                d_2217_v75_ = nw362_
                d_2217_v75_ = d_2217_v75_
            elif True:
                (globalState).f11 = not(False)
                d_2220_v76_: _dafny.Array
                nw363_ = _dafny.Array(int(0), 16)
                d_2220_v76_ = nw363_
                index353_ = default__.safeIndex(322, (d_2220_v76_).length(0))
                (d_2220_v76_)[index353_] = (d_2104_v10_).f41
                index354_ = default__.safeIndex(322, (d_2220_v76_).length(0))
                (d_2220_v76_)[index354_] = d_2104_v10_.f42
                d_2221_v77_: C4
                nw364_ = C4()
                nw364_.ctor__(p1, d_2209_cf104_)
                d_2221_v77_ = nw364_
                d_2222_v78_: _dafny.Array
                nw365_ = _dafny.Array(_dafny.Seq({}), 4)
                d_2222_v78_ = nw365_
                d_2223_v79_: _dafny.Seq
                d_2223_v79_ = _dafny.SeqWithoutIsStrInference([d_2222_v78_, d_2222_v78_, d_2222_v78_, d_2222_v78_])
                d_2224_v80_: _dafny.Array
                nw366_ = _dafny.Array(None, 4)
                nw366_[int(0)] = d_2222_v78_
                nw366_[int(1)] = d_2222_v78_
                nw366_[int(2)] = (d_2223_v79_)[default__.safeIndex(len(d_2212_v70_), len(d_2223_v79_))]
                nw366_[int(3)] = d_2222_v78_
                d_2224_v80_ = nw366_
                index355_ = default__.safeIndex(573, (d_2224_v80_).length(0))
                (d_2224_v80_)[index355_] = d_2222_v78_
                index356_ = default__.safeIndex(573, (d_2224_v80_).length(0))
                (d_2224_v80_)[index356_] = d_2222_v78_
                (globalState).f27 = ((d_2095_v0_) + (d_2095_v0_)) != (default__.fm3(globalState))
            d_2225_v81_: D6
            d_2225_v81_ = D6_DC16(p1, p1)
            d_2225_v81_ = d_2225_v81_
        elif True:
            d_2226___mcc_h17_ = source28_.cf113
            d_2227_cf113_ = d_2226___mcc_h17_
            (globalState).f19 = (d_2104_v10_).f41
            d_2228_v82_: _dafny.Set
            d_2228_v82_ = _dafny.Set({not (p1) or (p1), p1, p1, p1})
            d_2229_v84_: _dafny.Map
            d_2229_v84_ = _dafny.Map({True: d_2165_v47_})
            d_2230_v85_: _dafny.Seq
            d_2230_v85_ = _dafny.SeqWithoutIsStrInference([d_2103_v9_])
            d_2231_v86_: _dafny.MultiSet
            d_2231_v86_ = _dafny.MultiSet([43, p0])
            d_2232_v87_: _dafny.Array
            nw367_ = _dafny.Array(None, 28)
            nw367_[int(0)] = (d_2104_v10_).f41
            nw367_[int(1)] = (d_2104_v10_).f41
            def iife198_():
                coll86_ = _dafny.Map()
                compr_86_: int
                for compr_86_ in (d_2111_v15_).keys.Elements:
                    d_2233_v83_: int = compr_86_
                    if (d_2233_v83_) in (d_2111_v15_):
                        coll86_[(d_2233_v83_) * (d_2104_v10_.f42)] = p1
                return _dafny.Map(coll86_)
            nw367_[int(2)] = len(iife198_()
            )
            nw367_[int(3)] = 50
            nw367_[int(4)] = p2
            nw367_[int(5)] = (d_2104_v10_).f41
            nw367_[int(6)] = len(d_2168_v50_)
            nw367_[int(7)] = (d_2104_v10_).f41
            nw367_[int(8)] = d_2104_v10_.f42
            nw367_[int(9)] = d_2104_v10_.f42
            nw367_[int(10)] = d_2104_v10_.f42
            nw367_[int(11)] = p2
            nw367_[int(12)] = len(d_2229_v84_)
            nw367_[int(13)] = (d_2104_v10_).f41
            nw367_[int(14)] = p0
            nw367_[int(15)] = d_2104_v10_.f42
            nw367_[int(16)] = (d_2104_v10_).f41
            nw367_[int(17)] = len((d_2230_v85_)[default__.safeIndex((d_2104_v10_).f41, len(d_2230_v85_))])
            nw367_[int(18)] = (0) - (p2)
            nw367_[int(19)] = (_dafny.MultiSet([-150])).cardinality
            nw367_[int(20)] = p0
            nw367_[int(21)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qujafgmvo")))
            nw367_[int(22)] = len(d_2096_v1_)
            nw367_[int(23)] = (d_2231_v86_).cardinality
            nw367_[int(24)] = p0
            nw367_[int(25)] = d_2104_v10_.f42
            nw367_[int(26)] = d_2104_v10_.f42
            nw367_[int(27)] = p0
            d_2232_v87_ = nw367_
            d_2234_v88_: _dafny.Map
            d_2234_v88_ = _dafny.Map({d_2232_v87_: (d_2104_v10_).f41})
            d_2235_v89_: _dafny.Map
            d_2235_v89_ = _dafny.Map({D19_DC44(d_2234_v88_): p1})
            d_2236_v90_: D19
            d_2236_v90_ = D19_DC44(d_2234_v88_)
            d_2228_v82_ = _dafny.Set({p1, (((d_2235_v89_)[d_2236_v90_] if (d_2236_v90_) in (d_2235_v89_) else False)) == (p1)})
            source29_ = default__.fm44((0) - (d_2104_v10_.f42), (p1) or (p1), (d_2095_v0_) + (d_2095_v0_), -853, globalState)
            d_2237___mcc_h18_ = source29_
            d_2238_cf114_ = d_2237___mcc_h18_
            rhs401_ = d_2117_v21_
            rhs402_ = (_dafny.SeqWithoutIsStrInference([402])) + (d_2096_v1_)
            d_2117_v21_ = rhs401_
            d_2096_v1_ = rhs402_
            (globalState).f14 = (0) - ((d_2104_v10_).f41)
            d_2096_v1_ = d_2096_v1_
            d_2239_v91_: C2
            nw368_ = C2()
            nw368_.ctor__(d_2115_v19_, d_2228_v82_, (self).f28)
            d_2239_v91_ = nw368_
            (globalState).f27 = (len(d_2095_v0_)) > ((d_2104_v10_.f42) - (p2))
        r0 = d_2165_v47_
        r1 = (p0) - (p0)
        d_2240_v92_: D24
        d_2240_v92_ = D24_DC60(d_2168_v50_)
        r2 = (d_2168_v50_) | ((d_2240_v92_).cf115)
        r3 = _dafny.Set({d_2095_v0_})
        return r0, r1, r2, r3


class C8:
    def  __init__(self):
        self.f33: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f33):
        (self).f33 = f33

    def fm14(self, p0, p1, p2, globalState):
        return self.f33

    def m7(self, globalState):
        d_2241_v0_: bool
        d_2241_v0_ = True
        d_2242_v1_: _dafny.Set
        d_2242_v1_ = _dafny.Set({d_2241_v0_, d_2241_v0_})
        d_2243_v2_: _dafny.Seq
        d_2243_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nsnrubs"))
        d_2244_v3_: _dafny.Array
        nw369_ = _dafny.Array(None, 13)
        nw369_[int(0)] = not (d_2241_v0_) or (True)
        nw369_[int(1)] = not(d_2241_v0_)
        nw369_[int(2)] = d_2241_v0_
        nw369_[int(3)] = d_2241_v0_
        nw369_[int(4)] = d_2241_v0_
        nw369_[int(5)] = d_2241_v0_
        nw369_[int(6)] = False
        nw369_[int(7)] = default__.fm4(globalState)
        nw369_[int(8)] = d_2241_v0_
        nw369_[int(9)] = d_2241_v0_
        nw369_[int(10)] = (len(d_2242_v1_)) >= (len(d_2243_v2_))
        nw369_[int(11)] = (self.f33) <= ((0) - (self.f33))
        nw369_[int(12)] = default__.fm4(globalState)
        d_2244_v3_ = nw369_
        guard_loop_4_: int
        for guard_loop_4_ in _dafny.IntegerRange(0, (d_2244_v3_).length(0)):
            d_2245_i0_: int = guard_loop_4_
            if (True) and (((0) <= (d_2245_i0_)) and ((d_2245_i0_) < ((d_2244_v3_).length(0)))):
                (d_2244_v3_)[(d_2245_i0_)] = (975) != (self.f33)
        d_2246_v4_: _dafny.Seq
        d_2246_v4_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([d_2241_v0_]))])
        d_2247_v5_: _dafny.Set
        d_2247_v5_ = _dafny.Set({d_2243_v2_, default__.fm3(globalState), d_2243_v2_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ke")), d_2243_v2_})
        d_2248_v6_: _dafny.Array
        def lambda91_(d_2249_i1_):
            return default__.safeDivisionInt(d_2249_i1_, self.f33)

        init48_ = lambda91_
        nw370_ = _dafny.Array(None, 27)
        for i0_48_ in range(nw370_.length(0)):
            nw370_[i0_48_] = init48_(i0_48_)
        d_2248_v6_ = nw370_
        d_2250_v7_: _dafny.Map
        d_2250_v7_ = _dafny.Map({((d_2246_v4_)[default__.safeIndex(len(d_2247_v5_), len(d_2246_v4_))]) * (self.f33): d_2248_v6_})
        d_2251_v8_: _dafny.Map
        d_2251_v8_ = _dafny.Map({d_2241_v0_: d_2241_v0_})
        d_2252_v9_: _dafny.Seq
        d_2252_v9_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2241_v0_: d_2241_v0_}), d_2251_v8_])
        d_2253_v10_: D1
        d_2253_v10_ = D1_DC5(d_2241_v0_, d_2241_v0_, len(d_2252_v9_), d_2241_v0_, d_2241_v0_)
        d_2254_v11_: _dafny.Seq
        d_2254_v11_ = _dafny.SeqWithoutIsStrInference([d_2248_v6_])
        rhs403_ = (self.f33) * (len(_dafny.Map({self.f33: d_2253_v10_})))
        rhs404_ = (d_2250_v7_) | ((_dafny.Map({self.f33: (d_2254_v11_)[default__.safeIndex(self.f33, len(d_2254_v11_))]})) | (d_2250_v7_))
        rhs405_ = d_2241_v0_
        lhs328_ = globalState
        lhs329_ = globalState
        lhs328_.f19 = rhs403_
        d_2250_v7_ = rhs404_
        lhs329_.f11 = rhs405_
        index357_ = default__.safeIndex(0, (d_2248_v6_).length(0))
        (d_2248_v6_)[index357_] = self.f33
        d_2255_v12_: D1
        d_2255_v12_ = D1_DC4(self.f33, True)
        pat_let_tv49_ = d_2253_v10_
        index358_ = default__.safeIndex(0, (d_2248_v6_).length(0))
        def lambda92_(source30_):
            if source30_.is_DC4:
                d_2256___mcc_h0_ = source30_.cf5
                d_2257___mcc_h1_ = source30_.cf6
                d_2258_cf6_ = d_2257___mcc_h1_
                d_2259_cf5_ = d_2256___mcc_h0_
                return default__.safeModuloInt(self.f33, -349)
            elif source30_.is_DC5:
                d_2260___mcc_h2_ = source30_.cf7
                d_2261___mcc_h3_ = source30_.cf8
                d_2262___mcc_h4_ = source30_.cf9
                d_2263___mcc_h5_ = source30_.cf10
                d_2264___mcc_h6_ = source30_.cf11
                d_2265_cf11_ = d_2264___mcc_h6_
                d_2266_cf10_ = d_2263___mcc_h5_
                d_2267_cf9_ = d_2262___mcc_h4_
                d_2268_cf8_ = d_2261___mcc_h3_
                d_2269_cf7_ = d_2260___mcc_h2_
                return d_2267_cf9_
            elif True:
                d_2270___mcc_h7_ = source30_.cf4
                d_2271_cf4_ = d_2270___mcc_h7_
                return (pat_let_tv49_).cf9

        (d_2248_v6_)[index358_] = lambda92_(d_2255_v12_)
        source31_ = d_2255_v12_
        if source31_.is_DC4:
            d_2272___mcc_h8_ = source31_.cf5
            d_2273___mcc_h9_ = source31_.cf6
            d_2274_cf6_ = d_2273___mcc_h9_
            d_2275_cf5_ = d_2272___mcc_h8_
            d_2243_v2_ = (d_2243_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kyns")))
            if (d_2275_cf5_) < (len(_dafny.Map({d_2244_v3_: (d_2246_v4_)[default__.safeIndex((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], len(d_2246_v4_))]}))):
                d_2276_v13_: _dafny.Array
                def lambda93_(d_2277_v0_):
                    def lambda94_(d_2278_i2_):
                        return _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2277_v0_: self.f33}), _dafny.Map({d_2277_v0_: self.f33})])

                    return lambda94_

                init49_ = lambda93_(d_2241_v0_)
                nw371_ = _dafny.Array(None, 13)
                for i0_49_ in range(nw371_.length(0)):
                    nw371_[i0_49_] = init49_(i0_49_)
                d_2276_v13_ = nw371_
                d_2279_v14_: _dafny.Seq
                d_2279_v14_ = _dafny.SeqWithoutIsStrInference([d_2274_cf6_])
                d_2280_v15_: _dafny.Map
                d_2280_v15_ = _dafny.Map({d_2241_v0_: len(d_2279_v14_)})
                index359_ = default__.safeIndex(751, (d_2276_v13_).length(0))
                (d_2276_v13_)[index359_] = _dafny.SeqWithoutIsStrInference([d_2280_v15_, d_2280_v15_])
                d_2281_v16_: _dafny.Seq
                d_2281_v16_ = _dafny.SeqWithoutIsStrInference([d_2280_v15_, d_2280_v15_])
                d_2282_v17_: _dafny.Seq
                d_2282_v17_ = _dafny.SeqWithoutIsStrInference([((d_2281_v16_).set(default__.safeIndex(528, len(d_2281_v16_)), d_2280_v15_)).set(default__.safeIndex(408, len((d_2281_v16_).set(default__.safeIndex(528, len(d_2281_v16_)), d_2280_v15_))), d_2280_v15_)])
                d_2283_v18_: str
                d_2283_v18_ = _dafny.CodePoint('p')
                d_2284_v19_: D4
                d_2284_v19_ = D4_DC10(d_2283_v18_)
                pat_let_tv50_ = d_2283_v18_
                d_2285_v20_: _dafny.Map
                def iife199_(_pat_let56_0):
                    def iife200_(d_2286_dt__update__tmp_h0_):
                        def iife201_(_pat_let57_0):
                            def iife202_(d_2287_dt__update_hcf19_h0_):
                                return D4_DC10(d_2287_dt__update_hcf19_h0_)
                            return iife202_(_pat_let57_0)
                        return iife201_(pat_let_tv50_)
                    return iife200_(_pat_let56_0)
                d_2285_v20_ = _dafny.Map({d_2241_v0_: (iife199_(d_2284_v19_)).cf19})
                index360_ = default__.safeIndex(751, (d_2276_v13_).length(0))
                (d_2276_v13_)[index360_] = (d_2281_v16_) + ((((d_2282_v17_)[default__.safeIndex((0) - (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_2242_v1_, d_2285_v20_, d_2274_cf6_, globalState)), len(d_2282_v17_))]).set(default__.safeIndex((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], len((d_2282_v17_)[default__.safeIndex((0) - (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_2242_v1_, d_2285_v20_, d_2274_cf6_, globalState)), len(d_2282_v17_))])), d_2280_v15_)).set(default__.safeIndex(d_2275_cf5_, len(((d_2282_v17_)[default__.safeIndex((0) - (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_2242_v1_, d_2285_v20_, d_2274_cf6_, globalState)), len(d_2282_v17_))]).set(default__.safeIndex((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], len((d_2282_v17_)[default__.safeIndex((0) - (default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_2242_v1_, d_2285_v20_, d_2274_cf6_, globalState)), len(d_2282_v17_))])), d_2280_v15_))), d_2280_v15_))
                d_2288_v21_: _dafny.Array
                nw372_ = _dafny.Array(None, 8)
                nw372_[int(0)] = d_2243_v2_
                nw372_[int(1)] = d_2243_v2_
                nw372_[int(2)] = d_2243_v2_
                nw372_[int(3)] = d_2243_v2_
                nw372_[int(4)] = d_2243_v2_
                nw372_[int(5)] = d_2243_v2_
                nw372_[int(6)] = d_2243_v2_
                nw372_[int(7)] = (d_2243_v2_) + (d_2243_v2_)
                d_2288_v21_ = nw372_
                index361_ = default__.safeIndex(588, (d_2288_v21_).length(0))
                (d_2288_v21_)[index361_] = d_2243_v2_
                index362_ = default__.safeIndex(588, (d_2288_v21_).length(0))
                rhs406_ = d_2274_cf6_
                rhs407_ = d_2243_v2_
                rhs408_ = (d_2275_cf5_) - (self.f33)
                lhs330_ = globalState
                lhs331_ = d_2288_v21_
                lhs332_ = default__.safeIndex(588, (d_2288_v21_).length(0))
                lhs333_ = globalState
                lhs330_.f1 = rhs406_
                lhs331_[lhs332_] = rhs407_
                lhs333_.f0 = rhs408_
                d_2248_v6_ = d_2248_v6_
                d_2289_v22_: C4
                nw373_ = C4()
                nw373_.ctor__(not((d_2283_v18_) in ((d_2288_v21_)[default__.safeIndex(588, (d_2288_v21_).length(0))])), d_2288_v21_)
                d_2289_v22_ = nw373_
                d_2290_v23_: _dafny.Map
                d_2290_v23_ = _dafny.Map({(d_2275_cf5_) * ((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))]): (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))]})
                d_2290_v23_ = (d_2290_v23_).set((0) - (((0) - (d_2275_cf5_)) - (self.f33)), self.f33)
            elif True:
                d_2291_v24_: str
                d_2291_v24_ = _dafny.CodePoint('t')
                d_2292_v25_: _dafny.Map
                d_2292_v25_ = _dafny.Map({d_2291_v24_: d_2291_v24_})
                d_2251_v8_ = (d_2251_v8_).set(d_2274_cf6_, (d_2292_v25_) == (d_2292_v25_))
                d_2293_v26_: _dafny.Array
                nw374_ = _dafny.Array(None, 17)
                nw374_[int(0)] = d_2291_v24_
                nw374_[int(1)] = d_2291_v24_
                nw374_[int(2)] = d_2291_v24_
                nw374_[int(3)] = d_2291_v24_
                nw374_[int(4)] = d_2291_v24_
                nw374_[int(5)] = default__.fm5(len(d_2242_v1_), True, globalState)
                nw374_[int(6)] = d_2291_v24_
                nw374_[int(7)] = d_2291_v24_
                nw374_[int(8)] = _dafny.CodePoint('k')
                nw374_[int(9)] = d_2291_v24_
                nw374_[int(10)] = d_2291_v24_
                nw374_[int(11)] = d_2291_v24_
                nw374_[int(12)] = d_2291_v24_
                nw374_[int(13)] = d_2291_v24_
                nw374_[int(14)] = _dafny.CodePoint('e')
                nw374_[int(15)] = d_2291_v24_
                nw374_[int(16)] = d_2291_v24_
                d_2293_v26_ = nw374_
                d_2294_v27_: D16
                d_2294_v27_ = D16_DC35(d_2293_v26_)
                d_2295_v28_: C0
                nw375_ = C0()
                nw375_.ctor__(d_2242_v1_)
                d_2295_v28_ = nw375_
                d_2296_v29_: _dafny.Map
                d_2296_v29_ = _dafny.Map({d_2295_v28_: d_2274_cf6_})
                d_2297_v30_: _dafny.Map
                d_2297_v30_ = _dafny.Map({d_2274_cf6_: d_2291_v24_})
                rhs409_ = default__.safeModuloInt((len(_dafny.Map({d_2275_cf5_: len(_dafny.SeqWithoutIsStrInference([len(d_2296_v29_), (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], (d_2246_v4_)[default__.safeIndex(d_2275_cf5_, len(d_2246_v4_))]]))}))) + ((0) - ((d_2246_v4_)[default__.safeIndex(default__.fm2(d_2243_v2_, d_2242_v1_, d_2297_v30_, d_2241_v0_, globalState), len(d_2246_v4_))])), ((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))]) * (self.f33))
                rhs410_ = default__.fm4(globalState)
                rhs411_ = d_2294_v27_
                lhs334_ = globalState
                lhs335_ = globalState
                lhs334_.f19 = rhs409_
                lhs335_.f27 = rhs410_
                d_2294_v27_ = rhs411_
                d_2243_v2_ = (d_2243_v2_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ircgj")))
                index363_ = default__.safeIndex(393, (d_2244_v3_).length(0))
                (d_2244_v3_)[index363_] = d_2241_v0_
                d_2298_v31_: _dafny.Array
                nw376_ = _dafny.Array(_dafny.Map({}), 19)
                d_2298_v31_ = nw376_
                d_2299_v32_: _dafny.Map
                d_2299_v32_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vatxeihkp")): _dafny.Map({d_2274_cf6_: d_2274_cf6_})})
                index364_ = default__.safeIndex(834, (d_2298_v31_).length(0))
                (d_2298_v31_)[index364_] = d_2299_v32_
                index365_ = default__.safeIndex(393, (d_2244_v3_).length(0))
                index366_ = default__.safeIndex(834, (d_2298_v31_).length(0))
                rhs412_ = 36
                rhs413_ = not((d_2246_v4_) != ((_dafny.SeqWithoutIsStrInference([707])) + (_dafny.SeqWithoutIsStrInference([-49 for d_2300_i3_ in range(default__.abs(684))]))))
                rhs414_ = d_2299_v32_
                rhs415_ = d_2274_cf6_
                lhs336_ = globalState
                lhs337_ = d_2244_v3_
                lhs338_ = default__.safeIndex(393, (d_2244_v3_).length(0))
                lhs339_ = d_2298_v31_
                lhs340_ = default__.safeIndex(834, (d_2298_v31_).length(0))
                lhs341_ = globalState
                lhs336_.f14 = rhs412_
                lhs337_[lhs338_] = rhs413_
                lhs339_[lhs340_] = rhs414_
                lhs341_.f11 = rhs415_
                (globalState).f11 = (self.f33) == (self.f33)
            (globalState).f14 = len((d_2243_v2_) + (_dafny.SeqWithoutIsStrInference([(d_2243_v2_)[default__.safeIndex(d_2275_cf5_, len(d_2243_v2_))] for d_2301_i4_ in range(default__.abs(316))])))
            d_2302_v33_: D1
            d_2302_v33_ = D1_DC3(d_2241_v0_)
            d_2303_v34_: C5
            nw377_ = C5()
            nw377_.ctor__(d_2275_cf5_, d_2274_cf6_, d_2302_v33_)
            d_2303_v34_ = nw377_
            d_2303_v34_ = d_2303_v34_
        elif source31_.is_DC5:
            d_2304___mcc_h10_ = source31_.cf7
            d_2305___mcc_h11_ = source31_.cf8
            d_2306___mcc_h12_ = source31_.cf9
            d_2307___mcc_h13_ = source31_.cf10
            d_2308___mcc_h14_ = source31_.cf11
            d_2309_cf11_ = d_2308___mcc_h14_
            d_2310_cf10_ = d_2307___mcc_h13_
            d_2311_cf9_ = d_2306___mcc_h12_
            d_2312_cf8_ = d_2305___mcc_h11_
            d_2313_cf7_ = d_2304___mcc_h10_
            d_2314_v35_: _dafny.Seq
            d_2314_v35_ = _dafny.SeqWithoutIsStrInference([d_2312_cf8_])
            d_2315_v36_: _dafny.MultiSet
            d_2315_v36_ = _dafny.MultiSet([len(d_2314_v35_), d_2311_cf9_])
            d_2316_v37_: C6
            nw378_ = C6()
            nw378_.ctor__((d_2246_v4_)[default__.safeIndex((d_2315_v36_).cardinality, len(d_2246_v4_))])
            d_2316_v37_ = nw378_
            d_2316_v37_ = d_2316_v37_
            (globalState).f14 = self.f33
            (globalState).f0 = (0) - ((((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))] if not(not(d_2241_v0_)) else self.f33)) - ((d_2316_v37_).f34))
            d_2317_v38_: _dafny.Map
            d_2317_v38_ = _dafny.Map({(d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))]: (d_2316_v37_).f34})
            d_2318_v39_: str
            d_2318_v39_ = _dafny.CodePoint('a')
            index367_ = default__.safeIndex(0, (d_2248_v6_).length(0))
            (d_2248_v6_)[index367_] = (default__.safeModuloInt((d_2316_v37_).f34, (d_2316_v37_).f34)) - (default__.safeModuloInt(len((d_2243_v2_).set(default__.safeIndex(len(d_2317_v38_), len(d_2243_v2_)), d_2318_v39_)), (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))]))
        elif True:
            d_2319___mcc_h15_ = source31_.cf4
            d_2320_cf4_ = d_2319___mcc_h15_
            default__.m0(d_2244_v3_, d_2248_v6_, (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], globalState)
            d_2321_v40_: _dafny.Array
            def lambda95_(d_2322_v2_):
                def lambda96_(d_2323_i5_):
                    return d_2322_v2_

                return lambda96_

            init50_ = lambda95_(d_2243_v2_)
            nw379_ = _dafny.Array(None, 15)
            for i0_50_ in range(nw379_.length(0)):
                nw379_[i0_50_] = init50_(i0_50_)
            d_2321_v40_ = nw379_
            index368_ = default__.safeIndex(123, (d_2321_v40_).length(0))
            (d_2321_v40_)[index368_] = d_2243_v2_
            index369_ = default__.safeIndex(123, (d_2321_v40_).length(0))
            (d_2321_v40_)[index369_] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_2324_i6_ in range(default__.abs(884))])
            d_2325_v41_: _dafny.MultiSet
            d_2325_v41_ = _dafny.MultiSet([50])
            rhs416_ = (((d_2325_v41_).set(41, default__.abs(len(d_2246_v4_)))).intersection(_dafny.MultiSet([(d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))], (0) - ((0) - ((0) - ((d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))]))), self.f33]))).ispropersubset(_dafny.MultiSet([self.f33]))
            rhs417_ = d_2241_v0_
            lhs342_ = globalState
            lhs343_ = globalState
            lhs342_.f11 = rhs416_
            lhs343_.f1 = rhs417_
            (globalState).f0 = default__.safeModuloInt((0) - ((len((d_2321_v40_)[default__.safeIndex(123, (d_2321_v40_).length(0))]) if d_2320_cf4_ else (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))])), (d_2248_v6_)[default__.safeIndex(0, (d_2248_v6_).length(0))])
        d_2254_v11_ = d_2254_v11_
        d_2241_v0_ = d_2241_v0_


class C9:
    def  __init__(self):
        self.f31: bool = False
        self._f32: _dafny.Seq = _dafny.Seq({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C9"
    def ctor__(self, f31, f32):
        (self).f31 = f31
        (self)._f32 = f32

    def fm13(self, p0, p1, p2, globalState):
        def iife203_():
            coll87_ = _dafny.Map()
            compr_87_: int
            for compr_87_ in _dafny.IntegerRange(631, 773):
                d_2326_v0_: int = compr_87_
                if ((631) <= (d_2326_v0_)) and ((d_2326_v0_) < (773)):
                    coll87_[default__.safeModuloInt(d_2326_v0_, -201)] = _dafny.Map({len(_dafny.Set({(_dafny.MultiSet([len(_dafny.Map({190: self.f31}))])).cardinality})): len(_dafny.Map({self.f31: self.f31}))})
            return _dafny.Map(coll87_)
        return (0) - ((D0_DC2(455, len(_dafny.Set({_dafny.Map({self.f31: -660}), _dafny.Map({self.f31: len(iife203_()
)})})))).cf2)

    def m5(self, globalState):
        d_2327_v0_: int
        d_2327_v0_ = 727
        hi12_ = d_2327_v0_
        for d_2328_i0_ in range(d_2327_v0_, hi12_):
            d_2329_v1_: _dafny.Array
            nw380_ = _dafny.Array(None, 20)
            nw380_[int(0)] = self.f31
            nw380_[int(1)] = self.f31
            nw380_[int(2)] = not(self.f31)
            nw380_[int(3)] = self.f31
            nw380_[int(4)] = self.f31
            nw380_[int(5)] = self.f31
            nw380_[int(6)] = self.f31
            nw380_[int(7)] = self.f31
            nw380_[int(8)] = not(self.f31)
            nw380_[int(9)] = self.f31
            nw380_[int(10)] = self.f31
            nw380_[int(11)] = True
            nw380_[int(12)] = self.f31
            nw380_[int(13)] = self.f31
            nw380_[int(14)] = self.f31
            nw380_[int(15)] = self.f31
            nw380_[int(16)] = self.f31
            nw380_[int(17)] = False
            nw380_[int(18)] = self.f31
            nw380_[int(19)] = self.f31
            d_2329_v1_ = nw380_
            d_2330_v2_: _dafny.Array
            def lambda97_(d_2331_i0_):
                def lambda98_(d_2332_i1_):
                    return (d_2332_i1_) - (d_2331_i0_)

                return lambda98_

            init51_ = lambda97_(d_2328_i0_)
            nw381_ = _dafny.Array(None, 7)
            for i0_51_ in range(nw381_.length(0)):
                nw381_[i0_51_] = init51_(i0_51_)
            d_2330_v2_ = nw381_
            d_2333_v3_: _dafny.Seq
            d_2333_v3_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y')])
            d_2334_v4_: _dafny.Set
            d_2334_v4_ = _dafny.Set({d_2327_v0_, d_2327_v0_})
            d_2335_v5_: _dafny.Seq
            d_2335_v5_ = _dafny.SeqWithoutIsStrInference([len(d_2334_v4_)])
            d_2336_v6_: _dafny.MultiSet
            d_2336_v6_ = _dafny.MultiSet([len(d_2333_v3_), (_dafny.MultiSet(d_2335_v5_)).cardinality])
            d_2337_v7_: _dafny.Seq
            d_2337_v7_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({True: d_2328_i0_})), ((d_2336_v6_)[d_2328_i0_] if (d_2328_i0_) in (d_2336_v6_) else d_2327_v0_)])
            default__.m0(d_2329_v1_, d_2330_v2_, len((d_2337_v7_) + (d_2335_v5_)), d_2328_i0_, globalState)
            d_2338_v8_: _dafny.MultiSet
            d_2338_v8_ = _dafny.MultiSet([(not(self.f31)) and (self.f31), (D1_DC5(self.f31, True, d_2327_v0_, self.f31, self.f31)).cf10, self.f31])
            d_2338_v8_ = _dafny.MultiSet([default__.fm4(globalState)])
            d_2339_v9_: D2
            d_2339_v9_ = D2_DC8(self.f31, d_2330_v2_)
            d_2340_v10_: _dafny.MultiSet
            d_2340_v10_ = _dafny.MultiSet([d_2339_v9_, d_2339_v9_, d_2339_v9_])
            d_2340_v10_ = (d_2340_v10_) - (d_2340_v10_)
            d_2341_v11_: D2
            d_2341_v11_ = D2_DC6(d_2329_v1_)
            rhs418_ = (len((self).f32)) * ((0) - ((d_2328_i0_) + (d_2328_i0_)))
            rhs419_ = d_2341_v11_
            rhs420_ = d_2333_v3_
            lhs344_ = globalState
            lhs344_.f19 = rhs418_
            d_2341_v11_ = rhs419_
            d_2333_v3_ = rhs420_
        d_2342_v12_: _dafny.Set
        d_2342_v12_ = _dafny.Set({self.f31, self.f31})
        d_2343_v13_: C0
        nw382_ = C0()
        nw382_.ctor__(d_2342_v12_)
        d_2343_v13_ = nw382_
        hi13_ = (d_2327_v0_) - (d_2327_v0_)
        for d_2344_i2_ in range((d_2327_v0_) + (d_2327_v0_), hi13_):
            d_2345_v14_: _dafny.Array
            nw383_ = _dafny.Array(None, 8)
            nw383_[int(0)] = self.f31
            nw383_[int(1)] = not(self.f31)
            nw383_[int(2)] = self.f31
            nw383_[int(3)] = self.f31
            nw383_[int(4)] = not(self.f31)
            nw383_[int(5)] = self.f31
            nw383_[int(6)] = self.f31
            nw383_[int(7)] = default__.fm4(globalState)
            d_2345_v14_ = nw383_
            d_2346_v15_: _dafny.Array
            nw384_ = _dafny.Array(int(0), 29)
            d_2346_v15_ = nw384_
            d_2347_v16_: _dafny.Seq
            d_2347_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pgcywtq"))
            default__.m0(d_2345_v14_, d_2346_v15_, (d_2327_v0_) + (d_2327_v0_), len(d_2347_v16_), globalState)
            d_2348_v17_: D22
            d_2348_v17_ = D22_DC56(d_2347_v16_, len(d_2347_v16_), self.f31, self.f31)
            d_2349_v18_: _dafny.Map
            d_2349_v18_ = _dafny.Map({self.f31: d_2347_v16_})
            d_2350_v19_: str
            d_2350_v19_ = _dafny.CodePoint('p')
            d_2351_v20_: _dafny.Array
            nw385_ = _dafny.Array(None, 18)
            nw385_[int(0)] = d_2348_v17_
            nw385_[int(1)] = d_2348_v17_
            nw385_[int(2)] = d_2348_v17_
            nw385_[int(3)] = d_2348_v17_
            def iife204_(_pat_let58_0):
                def iife205_(d_2352_dt__update__tmp_h0_):
                    def iife206_(_pat_let59_0):
                        def iife207_(d_2353_dt__update_hcf106_h0_):
                            return D22_DC56((d_2352_dt__update__tmp_h0_).cf105, d_2353_dt__update_hcf106_h0_, (d_2352_dt__update__tmp_h0_).cf107, (d_2352_dt__update__tmp_h0_).cf108)
                        return iife207_(_pat_let59_0)
                    return iife206_(len((self).f32))
                return iife205_(_pat_let58_0)
            nw385_[int(4)] = iife204_(d_2348_v17_)
            nw385_[int(5)] = default__.fm45(True, globalState)
            nw385_[int(6)] = d_2348_v17_
            nw385_[int(7)] = D22_DC56((((d_2349_v18_)[self.f31] if (self.f31) in (d_2349_v18_) else d_2347_v16_)).set(default__.safeIndex(d_2327_v0_, len(((d_2349_v18_)[self.f31] if (self.f31) in (d_2349_v18_) else d_2347_v16_))), d_2350_v19_), d_2327_v0_, self.f31, self.f31)
            nw385_[int(8)] = d_2348_v17_
            nw385_[int(9)] = d_2348_v17_
            nw385_[int(10)] = d_2348_v17_
            nw385_[int(11)] = d_2348_v17_
            nw385_[int(12)] = d_2348_v17_
            nw385_[int(13)] = D22_DC56(d_2347_v16_, d_2327_v0_, self.f31, self.f31)
            nw385_[int(14)] = d_2348_v17_
            nw385_[int(15)] = d_2348_v17_
            nw385_[int(16)] = d_2348_v17_
            nw385_[int(17)] = d_2348_v17_
            d_2351_v20_ = nw385_
            pat_let_tv51_ = d_2347_v16_
            index370_ = default__.safeIndex(686, (d_2351_v20_).length(0))
            def iife208_(_pat_let60_0):
                def iife209_(d_2354_dt__update__tmp_h1_):
                    def iife210_(_pat_let61_0):
                        def iife211_(d_2355_dt__update_hcf107_h0_):
                            def iife212_(_pat_let62_0):
                                def iife213_(d_2356_dt__update_hcf105_h0_):
                                    return D22_DC56(d_2356_dt__update_hcf105_h0_, (d_2354_dt__update__tmp_h1_).cf106, d_2355_dt__update_hcf107_h0_, (d_2354_dt__update__tmp_h1_).cf108)
                                return iife213_(_pat_let62_0)
                            return iife212_(pat_let_tv51_)
                        return iife211_(_pat_let61_0)
                    return iife210_(self.f31)
                return iife209_(_pat_let60_0)
            (d_2351_v20_)[index370_] = iife208_(d_2348_v17_)
            pat_let_tv52_ = d_2347_v16_
            index371_ = default__.safeIndex(686, (d_2351_v20_).length(0))
            def iife214_(_pat_let63_0):
                def iife215_(d_2357_dt__update__tmp_h2_):
                    def iife216_(_pat_let64_0):
                        def iife217_(d_2358_dt__update_hcf107_h1_):
                            def iife218_(_pat_let65_0):
                                def iife219_(d_2359_dt__update_hcf105_h1_):
                                    return D22_DC56(d_2359_dt__update_hcf105_h1_, (d_2357_dt__update__tmp_h2_).cf106, d_2358_dt__update_hcf107_h1_, (d_2357_dt__update__tmp_h2_).cf108)
                                return iife219_(_pat_let65_0)
                            return iife218_(pat_let_tv52_)
                        return iife217_(_pat_let64_0)
                    return iife216_(False)
                return iife215_(_pat_let63_0)
            rhs421_ = iife214_(d_2348_v17_)
            rhs422_ = self.f31
            rhs423_ = not((self.f31 if self.f31 else self.f31))
            lhs345_ = d_2351_v20_
            lhs346_ = default__.safeIndex(686, (d_2351_v20_).length(0))
            lhs347_ = globalState
            lhs348_ = globalState
            lhs345_[lhs346_] = rhs421_
            lhs347_.f1 = rhs422_
            lhs348_.f27 = rhs423_
            d_2360_v21_: _dafny.Seq
            d_2360_v21_ = _dafny.SeqWithoutIsStrInference([d_2347_v16_, d_2347_v16_, d_2347_v16_])
            d_2361_v22_: _dafny.MultiSet
            d_2361_v22_ = _dafny.MultiSet([self.f31])
            d_2362_v23_: _dafny.Array
            nw386_ = _dafny.Array(None, 1)
            nw386_[int(0)] = (d_2360_v21_)[default__.safeIndex(((d_2361_v22_)[self.f31] if (self.f31) in (d_2361_v22_) else len(d_2347_v16_)), len(d_2360_v21_))]
            d_2362_v23_ = nw386_
            index372_ = default__.safeIndex(891, (d_2362_v23_).length(0))
            (d_2362_v23_)[index372_] = d_2347_v16_
            index373_ = default__.safeIndex(891, (d_2362_v23_).length(0))
            (d_2362_v23_)[index373_] = (default__.fm3(globalState)) + (d_2347_v16_)
            (globalState).f11 = not(default__.fm4(globalState))
        d_2363_v24_: _dafny.Array
        nw387_ = _dafny.Array(_dafny.Array(None, 0), 29)
        d_2363_v24_ = nw387_
        d_2364_v25_: _dafny.Array
        nw388_ = _dafny.Array(int(0), 15)
        d_2364_v25_ = nw388_
        index374_ = default__.safeIndex(504, (d_2363_v24_).length(0))
        (d_2363_v24_)[index374_] = d_2364_v25_
        index375_ = default__.safeIndex(504, (d_2363_v24_).length(0))
        (d_2363_v24_)[index375_] = d_2364_v25_
        d_2365_v26_: _dafny.Seq
        d_2365_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "n"))
        d_2366_v27_: _dafny.Map
        d_2366_v27_ = _dafny.Map({d_2365_v26_: d_2327_v0_})
        d_2366_v27_ = (d_2366_v27_).set(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tvkotmruk")), d_2327_v0_)
        d_2367_v28_: _dafny.Map
        d_2367_v28_ = _dafny.Map({d_2327_v0_: True})
        d_2368_v29_: str
        d_2368_v29_ = _dafny.CodePoint('o')
        d_2369_v30_: _dafny.Set
        d_2369_v30_ = _dafny.Set({d_2368_v29_})
        d_2370_v31_: D4
        d_2370_v31_ = D4_DC12(((d_2367_v28_)[d_2327_v0_] if (d_2327_v0_) in (d_2367_v28_) else self.f31), d_2369_v30_)
        d_2371_v32_: D4
        d_2371_v32_ = D4_DC13(d_2370_v31_)
        source32_ = (d_2371_v32_ if (not(self.f31) if self.f31 else self.f31) else d_2371_v32_)
        if source32_.is_DC11:
            d_2372___mcc_h0_ = source32_.cf20
            d_2373_cf20_ = d_2372___mcc_h0_
            d_2374_v33_: _dafny.Array
            nw389_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 26)
            d_2374_v33_ = nw389_
            d_2375_v34_: _dafny.Seq
            d_2375_v34_ = _dafny.SeqWithoutIsStrInference([d_2365_v26_])
            index376_ = default__.safeIndex(878, (d_2374_v33_).length(0))
            (d_2374_v33_)[index376_] = (_dafny.SeqWithoutIsStrInference([d_2368_v29_ for d_2376_i3_ in range(default__.abs(224))])) + ((d_2375_v34_)[default__.safeIndex(d_2327_v0_, len(d_2375_v34_))])
            index377_ = default__.safeIndex(878, (d_2374_v33_).length(0))
            (d_2374_v33_)[index377_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdi"))
            d_2377_v35_: _dafny.MultiSet
            d_2377_v35_ = _dafny.MultiSet([len(d_2365_v26_), default__.safeDivisionInt(d_2327_v0_, (_dafny.MultiSet([self.f31])).cardinality), d_2327_v0_])
            d_2378_v36_: _dafny.Seq
            d_2378_v36_ = _dafny.SeqWithoutIsStrInference([(d_2377_v35_) - (d_2377_v35_), (_dafny.MultiSet(default__.fm6(self.f31, d_2327_v0_, (d_2374_v33_)[default__.safeIndex(878, (d_2374_v33_).length(0))], globalState))).intersection(d_2377_v35_), d_2377_v35_, d_2377_v35_])
            d_2377_v35_ = (d_2378_v36_)[default__.safeIndex((d_2327_v0_ if self.f31 else d_2327_v0_), len(d_2378_v36_))]
            d_2379_v37_: _dafny.Map
            d_2379_v37_ = _dafny.Map({self.f31: d_2368_v29_})
            d_2367_v28_ = (d_2367_v28_).set(default__.fm2(d_2365_v26_, d_2342_v12_, d_2379_v37_, False, globalState), self.f31)
            d_2380_v38_: _dafny.Array
            nw390_ = _dafny.Array(False, 1)
            d_2380_v38_ = nw390_
            index378_ = default__.safeIndex(62, (d_2380_v38_).length(0))
            (d_2380_v38_)[index378_] = default__.fm4(globalState)
            index379_ = default__.safeIndex(62, (d_2380_v38_).length(0))
            (d_2380_v38_)[index379_] = self.f31
        elif source32_.is_DC12:
            d_2381___mcc_h1_ = source32_.cf21
            d_2382___mcc_h2_ = source32_.cf22
            d_2383_cf22_ = d_2382___mcc_h2_
            d_2384_cf21_ = d_2381___mcc_h1_
            (globalState).f0 = d_2327_v0_
            d_2385_v39_: _dafny.Seq
            d_2385_v39_ = _dafny.SeqWithoutIsStrInference([d_2327_v0_])
            d_2386_v40_: D18
            d_2386_v40_ = D18_DC41(d_2385_v39_)
            d_2385_v39_ = (((_dafny.SeqWithoutIsStrInference([d_2327_v0_])) + (d_2385_v39_)) + ((d_2386_v40_).cf71)).set(default__.safeIndex((d_2327_v0_) - (d_2327_v0_), len(((_dafny.SeqWithoutIsStrInference([d_2327_v0_])) + (d_2385_v39_)) + ((d_2386_v40_).cf71))), d_2327_v0_)
            d_2387_v41_: _dafny.MultiSet
            d_2387_v41_ = _dafny.MultiSet([self.f31, self.f31, default__.fm4(globalState)])
            d_2388_v42_: _dafny.Map
            d_2388_v42_ = _dafny.Map({True: d_2387_v41_})
            d_2389_v43_: D8
            d_2389_v43_ = D8_DC20(d_2327_v0_, len(d_2388_v42_))
            (globalState).f0 = (d_2389_v43_).cf31
            d_2390_v44_: _dafny.Array
            nw391_ = _dafny.Array(_dafny.Set({}), 25)
            d_2390_v44_ = nw391_
            index380_ = default__.safeIndex(642, (d_2390_v44_).length(0))
            (d_2390_v44_)[index380_] = (d_2343_v13_).f43
            index381_ = default__.safeIndex(642, (d_2390_v44_).length(0))
            (d_2390_v44_)[index381_] = (_dafny.Set({d_2384_cf21_})) | (default__.fm24((self).f32, d_2384_cf21_, globalState))
        elif source32_.is_DC10:
            d_2391___mcc_h3_ = source32_.cf19
            d_2392_cf19_ = d_2391___mcc_h3_
            if not(not(not(self.f31))):
                d_2393_v45_: _dafny.Set
                d_2393_v45_ = _dafny.Set({_dafny.Map({d_2327_v0_: d_2392_cf19_})})
                rhs424_ = d_2365_v26_
                rhs425_ = len((d_2393_v45_) | (d_2393_v45_))
                lhs349_ = globalState
                d_2365_v26_ = rhs424_
                lhs349_.f0 = rhs425_
                (globalState).f19 = default__.safeModuloInt(362, d_2327_v0_)
                d_2394_v46_: _dafny.Array
                nw392_ = _dafny.Array(_dafny.Map({}), 22)
                d_2394_v46_ = nw392_
                index382_ = default__.safeIndex(785, (d_2394_v46_).length(0))
                (d_2394_v46_)[index382_] = default__.fm46(d_2327_v0_, globalState)
                d_2395_v47_: _dafny.Map
                d_2395_v47_ = _dafny.Map({len((self).f32): d_2365_v26_})
                d_2396_v48_: _dafny.Map
                d_2396_v48_ = _dafny.Map({self.f31: ((d_2395_v47_)[544] if (544) in (d_2395_v47_) else d_2365_v26_)})
                index383_ = default__.safeIndex(785, (d_2394_v46_).length(0))
                (d_2394_v46_)[index383_] = ((_dafny.Map({not(self.f31): d_2365_v26_}) if self.f31 else d_2396_v48_)).set(self.f31, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cf")))
                (globalState).f14 = (d_2327_v0_) * ((0) - (-562))
                index384_ = default__.safeIndex(634, (d_2364_v25_).length(0))
                (d_2364_v25_)[index384_] = d_2327_v0_
                index385_ = default__.safeIndex(634, (d_2364_v25_).length(0))
                (d_2364_v25_)[index385_] = d_2327_v0_
            elif True:
                (globalState).f20 = d_2367_v28_
                d_2397_v49_: _dafny.Map
                d_2397_v49_ = _dafny.Map({(self.f31 if self.f31 else self.f31): self.f31})
                d_2397_v49_ = (d_2397_v49_).set(self.f31, ((self).f32)[default__.safeIndex(d_2327_v0_, len((self).f32))])
                d_2398_v50_: D1
                d_2398_v50_ = D1_DC3(self.f31)
                d_2399_v51_: C5
                nw393_ = C5()
                nw393_.ctor__(d_2327_v0_, self.f31, d_2398_v50_)
                d_2399_v51_ = nw393_
                d_2400_v52_: _dafny.Set
                d_2400_v52_ = _dafny.Set({d_2399_v51_})
                d_2401_v53_: _dafny.Map
                d_2401_v53_ = _dafny.Map({d_2397_v49_: (d_2400_v52_) - (d_2400_v52_)})
                d_2401_v53_ = (d_2401_v53_).set(d_2397_v49_, (d_2400_v52_) | (d_2400_v52_))
                d_2402_v54_: _dafny.MultiSet
                d_2402_v54_ = _dafny.MultiSet([(d_2399_v51_).f36])
                d_2403_v55_: _dafny.Map
                d_2403_v55_ = _dafny.Map({(d_2399_v51_).f35: d_2402_v54_})
                d_2403_v55_ = (d_2403_v55_).set(default__.safeDivisionInt(-708, (d_2399_v51_).f35), (default__.fm25(len(d_2397_v49_), True, (self).f32, d_2327_v0_, globalState)) | (d_2402_v54_))
                d_2404_v56_: _dafny.Seq
                d_2404_v56_ = _dafny.SeqWithoutIsStrInference([d_2327_v0_])
                d_2405_v57_: _dafny.Array
                nw394_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_2405_v57_ = nw394_
                d_2406_v60_: _dafny.Array
                def lambda99_(d_2407_v51_, d_2408_v0_, d_2409_v26_, d_2410_v28_):
                    def lambda100_(d_2411_i4_):
                        def iife220_():
                            coll88_ = _dafny.Set()
                            compr_88_: D8
                            for compr_88_ in (_dafny.SeqWithoutIsStrInference([D8_DC21((d_2407_v51_).f36, not((d_2407_v51_).f36), d_2408_v0_) for d_2412_i5_ in range(default__.abs(513))])).Elements:
                                d_2413_v58_: D8 = compr_88_
                                if (d_2413_v58_) in (_dafny.SeqWithoutIsStrInference([D8_DC21((d_2407_v51_).f36, not((d_2407_v51_).f36), d_2408_v0_) for d_2412_i5_ in range(default__.abs(513))])):
                                    coll88_ = coll88_.union(_dafny.Set([d_2413_v58_]))
                            return _dafny.Set(coll88_)
                        def iife221_():
                            coll89_ = _dafny.Set()
                            compr_89_: D8
                            for compr_89_ in (_dafny.Set({D8_DC21((d_2407_v51_).f36, self.f31, len(d_2409_v26_)), D8_DC21(((d_2410_v28_)[(d_2407_v51_).f35] if ((d_2407_v51_).f35) in (d_2410_v28_) else (d_2407_v51_).f36), not(self.f31), d_2408_v0_), D8_DC21((d_2407_v51_).f36, (d_2407_v51_).f36, (d_2407_v51_).f35)})).Elements:
                                d_2414_v59_: D8 = compr_89_
                                if (d_2414_v59_) in (_dafny.Set({D8_DC21((d_2407_v51_).f36, self.f31, len(d_2409_v26_)), D8_DC21(((d_2410_v28_)[(d_2407_v51_).f35] if ((d_2407_v51_).f35) in (d_2410_v28_) else (d_2407_v51_).f36), not(self.f31), d_2408_v0_), D8_DC21((d_2407_v51_).f36, (d_2407_v51_).f36, (d_2407_v51_).f35)})):
                                    coll89_ = coll89_.union(_dafny.Set([d_2414_v59_]))
                            return _dafny.Set(coll89_)
                        return _dafny.SeqWithoutIsStrInference([_dafny.Set({D8_DC21((d_2407_v51_).f36, (d_2407_v51_).f36, (d_2407_v51_).f35)}), iife220_()
                        , iife221_()
                        ])

                    return lambda100_

                init52_ = lambda99_(d_2399_v51_, d_2327_v0_, d_2365_v26_, d_2367_v28_)
                nw395_ = _dafny.Array(None, 8)
                for i0_52_ in range(nw395_.length(0)):
                    nw395_[i0_52_] = init52_(i0_52_)
                d_2406_v60_ = nw395_
                index386_ = default__.safeIndex(686, (d_2405_v57_).length(0))
                (d_2405_v57_)[index386_] = d_2406_v60_
                index387_ = default__.safeIndex(686, (d_2405_v57_).length(0))
                rhs426_ = (_dafny.SeqWithoutIsStrInference([d_2327_v0_])) + ((_dafny.SeqWithoutIsStrInference([432])) + (d_2404_v56_))
                rhs427_ = len(d_2365_v26_)
                rhs428_ = d_2406_v60_
                rhs429_ = not ((len(d_2365_v26_)) not in (_dafny.MultiSet(default__.fm6((d_2399_v51_).f36, d_2327_v0_, d_2365_v26_, globalState)))) or (self.f31)
                rhs430_ = (d_2366_v27_).set(d_2365_v26_, len((d_2367_v28_) | (d_2367_v28_)))
                lhs350_ = d_2405_v57_
                lhs351_ = default__.safeIndex(686, (d_2405_v57_).length(0))
                lhs352_ = globalState
                d_2404_v56_ = rhs426_
                d_2327_v0_ = rhs427_
                lhs350_[lhs351_] = rhs428_
                lhs352_.f1 = rhs429_
                d_2366_v27_ = rhs430_
            d_2365_v26_ = (d_2365_v26_) + ((((D22_DC56(_dafny.SeqWithoutIsStrInference([d_2368_v29_ for d_2415_i6_ in range(default__.abs(591))]), len(_dafny.SeqWithoutIsStrInference([self.f31])), self.f31, default__.fm4(globalState))).cf105).set(default__.safeIndex(len(d_2365_v26_), len((D22_DC56(_dafny.SeqWithoutIsStrInference([d_2368_v29_ for d_2416_i6_ in range(default__.abs(591))]), len(_dafny.SeqWithoutIsStrInference([self.f31])), self.f31, default__.fm4(globalState))).cf105)), d_2368_v29_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nfugcqnqm"))))
            d_2365_v26_ = d_2365_v26_
            d_2417_v61_: _dafny.Seq
            d_2417_v61_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2327_v0_), default__.safeModuloInt(d_2327_v0_, d_2327_v0_)])
            d_2418_v62_: D19
            d_2418_v62_ = D19_DC45(d_2327_v0_, not(self.f31))
            d_2419_v63_: _dafny.Set
            d_2419_v63_ = _dafny.Set({(d_2418_v62_).cf82, d_2327_v0_})
            (globalState).f24 = (d_2417_v61_)[default__.safeIndex(len(d_2419_v63_), len(d_2417_v61_))]
        elif True:
            d_2420___mcc_h4_ = source32_.cf23
            d_2421_cf23_ = d_2420___mcc_h4_
            d_2422_v64_: _dafny.Array
            def lambda101_(d_2423_v0_):
                def lambda102_(d_2424_i7_):
                    return ((self).f32)[default__.safeIndex(d_2423_v0_, len((self).f32))]

                return lambda102_

            init53_ = lambda101_(d_2327_v0_)
            nw396_ = _dafny.Array(None, 13)
            for i0_53_ in range(nw396_.length(0)):
                nw396_[i0_53_] = init53_(i0_53_)
            d_2422_v64_ = nw396_
            index388_ = default__.safeIndex(800, (d_2422_v64_).length(0))
            (d_2422_v64_)[index388_] = self.f31
            index389_ = default__.safeIndex(800, (d_2422_v64_).length(0))
            (d_2422_v64_)[index389_] = not(True)
            d_2425_v65_: _dafny.Seq
            d_2425_v65_ = _dafny.SeqWithoutIsStrInference([d_2327_v0_, d_2327_v0_])
            d_2426_v66_: _dafny.Map
            d_2426_v66_ = _dafny.Map({d_2425_v65_: (d_2422_v64_)[default__.safeIndex(800, (d_2422_v64_).length(0))]})
            d_2426_v66_ = (d_2426_v66_).set((d_2425_v65_) + (d_2425_v65_), self.f31)
            (globalState).f11 = False
            (globalState).f19 = d_2327_v0_

    def m6(self, globalState):
        r0: bool = False
        r1: _dafny.Seq = _dafny.Seq({})
        r2: int = int(0)
        r3: int = int(0)
        d_2427_v0_: _dafny.Array
        nw397_ = _dafny.Array(_dafny.Seq({}), 20)
        d_2427_v0_ = nw397_
        guard_loop_5_: int
        for guard_loop_5_ in _dafny.IntegerRange(0, (d_2427_v0_).length(0)):
            d_2428_i0_: int = guard_loop_5_
            if (True) and (((0) <= (d_2428_i0_)) and ((d_2428_i0_) < ((d_2427_v0_).length(0)))):
                (d_2427_v0_)[(d_2428_i0_)] = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.Set({153, 266, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmofyvv")))}), _dafny.Set({726})])])
        (self).f31 = (self.f31) and (self.f31)
        d_2429_v1_: _dafny.Seq
        d_2429_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hgtw"))
        (globalState).f14 = (0) - (len((d_2429_v1_ if (len(d_2429_v1_)) < (len((self).f32)) else _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2430_i1_ in range(default__.abs(841))]))))
        d_2431_v2_: _dafny.Array
        def lambda103_(d_2432_i2_):
            return self.f31

        init54_ = lambda103_
        nw398_ = _dafny.Array(None, 11)
        for i0_54_ in range(nw398_.length(0)):
            nw398_[i0_54_] = init54_(i0_54_)
        d_2431_v2_ = nw398_
        index390_ = default__.safeIndex(455, (d_2431_v2_).length(0))
        (d_2431_v2_)[index390_] = self.f31
        d_2433_v3_: int
        d_2433_v3_ = 309
        index391_ = default__.safeIndex(455, (d_2431_v2_).length(0))
        (d_2431_v2_)[index391_] = (d_2433_v3_) > ((940) - (d_2433_v3_))
        if (d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]:
            d_2434_v4_: str
            d_2434_v4_ = _dafny.CodePoint('w')
            d_2434_v4_ = d_2434_v4_
            d_2435_v5_: _dafny.MultiSet
            d_2435_v5_ = _dafny.MultiSet([d_2431_v2_])
            d_2436_v6_: _dafny.Map
            d_2436_v6_ = _dafny.Map({((self).fm13(len(default__.fm23(globalState)), d_2433_v3_, self.f31, globalState)) <= (d_2433_v3_): (d_2435_v5_) | (d_2435_v5_)})
            d_2435_v5_ = ((d_2436_v6_)[(d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]] if ((d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]) in (d_2436_v6_) else d_2435_v5_)
            d_2437_v7_: _dafny.Seq
            d_2437_v7_ = _dafny.SeqWithoutIsStrInference([d_2433_v3_, d_2433_v3_, d_2433_v3_])
            index392_ = default__.safeIndex(455, (d_2431_v2_).length(0))
            (d_2431_v2_)[index392_] = (_dafny.MultiSet([d_2437_v7_])).ispropersubset(default__.fm47(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rqqdcgfg")), globalState))
            d_2438_v8_: D1
            d_2438_v8_ = D1_DC3(default__.fm4(globalState))
            d_2439_v9_: C7
            nw399_ = C7()
            nw399_.ctor__(d_2438_v8_)
            d_2439_v9_ = nw399_
            r3 = (0) - (default__.safeModuloInt(d_2433_v3_, d_2433_v3_))
        elif True:
            d_2433_v3_ = d_2433_v3_
            index393_ = default__.safeIndex(455, (d_2431_v2_).length(0))
            def iife222_():
                coll90_ = _dafny.Set()
                compr_90_: int
                for compr_90_ in _dafny.IntegerRange(659, -299):
                    d_2440_v10_: int = compr_90_
                    if ((659) <= (d_2440_v10_)) and ((d_2440_v10_) < (-299)):
                        coll90_ = coll90_.union(_dafny.Set([default__.safeDivisionInt(d_2440_v10_, d_2433_v3_)]))
                return _dafny.Set(coll90_)
            rhs431_ = False
            rhs432_ = (d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]
            rhs433_ = (self).fm13((0) - (d_2433_v3_), d_2433_v3_, (d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))], globalState)
            rhs434_ = (d_2433_v3_) + (default__.safeDivisionInt(len(iife222_()
            ), d_2433_v3_))
            lhs353_ = d_2431_v2_
            lhs354_ = default__.safeIndex(455, (d_2431_v2_).length(0))
            lhs355_ = globalState
            lhs356_ = globalState
            lhs357_ = globalState
            lhs353_[lhs354_] = rhs431_
            lhs355_.f27 = rhs432_
            lhs356_.f0 = rhs433_
            lhs357_.f24 = rhs434_
            d_2441_v11_: str
            d_2441_v11_ = _dafny.CodePoint('i')
            (globalState).f14 = (d_2433_v3_) - (len((D16_DC37(_dafny.SeqWithoutIsStrInference([False]), d_2429_v1_, d_2441_v11_)).cf65))
            d_2442_v12_: _dafny.Set
            d_2442_v12_ = _dafny.Set({-232})
            d_2442_v12_ = d_2442_v12_
            d_2443_v13_: D1
            d_2443_v13_ = D1_DC3((d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))])
            d_2444_v14_: C1
            nw400_ = C1()
            nw400_.ctor__(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwqac"))), d_2433_v3_, d_2443_v13_)
            d_2444_v14_ = nw400_
        d_2445_v15_: D22
        d_2445_v15_ = D22_DC56(d_2429_v1_, (0) - (len((d_2429_v1_ if (d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))] else d_2429_v1_))), (d_2433_v3_) <= (len((self).f32)), (((self).f32).set(default__.safeIndex(d_2433_v3_, len((self).f32)), self.f31)) != ((self).f32))
        d_2445_v15_ = d_2445_v15_
        r0 = (len(d_2429_v1_)) <= (d_2433_v3_)
        d_2446_v16_: _dafny.Seq
        d_2446_v16_ = _dafny.SeqWithoutIsStrInference([48])
        r1 = (d_2446_v16_) + (_dafny.SeqWithoutIsStrInference([(0) - (d_2433_v3_), d_2433_v3_]))
        d_2447_v17_: _dafny.MultiSet
        d_2447_v17_ = _dafny.MultiSet([(d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))], (d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]])
        r2 = ((self).fm13((d_2447_v17_).cardinality, d_2433_v3_, self.f31, globalState) if not((d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]) else d_2433_v3_)
        d_2448_v18_: _dafny.Map
        d_2448_v18_ = _dafny.Map({self.f31: self.f31})
        d_2449_v19_: D24
        d_2449_v19_ = D24_DC61(not((d_2431_v2_)[default__.safeIndex(455, (d_2431_v2_).length(0))]), self.f31, (0) - (len(d_2448_v18_)))
        r3 = (d_2449_v19_).cf118
        return r0, r1, r2, r3

    @property
    def f32(self):
        return self._f32

class C10(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        self._f30: _dafny.Set = _dafny.Set({})
        pass

    def __dafnystr__(self) -> str:
        return "_module.C10"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f30, f28):
        (self)._f30 = f30
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        source33_ = D0_DC0(_dafny.Map({len(_dafny.Set({-614})): 871}))
        if source33_.is_DC1:
            d_2450___mcc_h0_ = source33_.cf1
            d_2451_cf1_ = d_2450___mcc_h0_
            return (True) or (False)
        elif source33_.is_DC2:
            d_2452___mcc_h1_ = source33_.cf2
            d_2453___mcc_h2_ = source33_.cf3
            d_2454_cf3_ = d_2453___mcc_h2_
            d_2455_cf2_ = d_2452___mcc_h1_
            return (_dafny.Map({585: d_2455_cf2_})) == (_dafny.Map({d_2454_cf3_: d_2455_cf2_}))
        elif True:
            d_2456___mcc_h3_ = source33_.cf0
            d_2457_cf0_ = d_2456___mcc_h3_
            return not((True) == (not(False)))

    def fm8(self, p0, p1, p2, globalState):
        return _dafny.CodePoint('y')

    def fm11(self, p0, p1, p2, globalState):
        return ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qiecn"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xxfru")))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkycwpcs")))

    def fm12(self, globalState):
        return default__.safeDivisionInt(-519, (0) - ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "toyav")))) + (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ywdle"))))))

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2458_v0_: bool
        d_2458_v0_ = True
        d_2459_i0_: int
        d_2459_i0_ = 0
        with _dafny.label("15"):
            while d_2458_v0_:
                with _dafny.c_label("15"):
                    if (d_2459_i0_) >= (100):
                        raise _dafny.Break("15")
                    d_2459_i0_ = (d_2459_i0_) + (1)
                    d_2460_v1_: int
                    d_2460_v1_ = 868
                    d_2461_v2_: _dafny.MultiSet
                    d_2461_v2_ = _dafny.MultiSet([d_2460_v1_, d_2460_v1_, d_2460_v1_])
                    d_2462_v3_: _dafny.Seq
                    d_2462_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wcnqasj"))
                    d_2463_v4_: str
                    d_2463_v4_ = _dafny.CodePoint('i')
                    d_2464_v5_: _dafny.Map
                    d_2464_v5_ = _dafny.Map({d_2458_v0_: d_2463_v4_})
                    d_2465_v6_: _dafny.Seq
                    d_2465_v6_ = _dafny.SeqWithoutIsStrInference([d_2458_v0_, d_2458_v0_])
                    d_2466_v7_: _dafny.Map
                    d_2466_v7_ = _dafny.Map({d_2458_v0_: (d_2465_v6_)[default__.safeIndex(d_2460_v1_, len(d_2465_v6_))]})
                    d_2467_v8_: _dafny.MultiSet
                    d_2467_v8_ = _dafny.MultiSet([d_2458_v0_, d_2458_v0_])
                    d_2468_v9_: _dafny.Array
                    nw401_ = _dafny.Array(None, 25)
                    nw401_[int(0)] = d_2460_v1_
                    nw401_[int(1)] = ((d_2461_v2_) - (d_2461_v2_)).cardinality
                    nw401_[int(2)] = d_2460_v1_
                    nw401_[int(3)] = d_2460_v1_
                    nw401_[int(4)] = default__.safeDivisionInt(d_2460_v1_, d_2460_v1_)
                    nw401_[int(5)] = d_2460_v1_
                    nw401_[int(6)] = d_2460_v1_
                    nw401_[int(7)] = d_2460_v1_
                    nw401_[int(8)] = default__.safeDivisionInt(d_2460_v1_, len(d_2462_v3_))
                    nw401_[int(9)] = d_2460_v1_
                    nw401_[int(10)] = d_2460_v1_
                    nw401_[int(11)] = 392
                    nw401_[int(12)] = d_2460_v1_
                    nw401_[int(13)] = default__.safeModuloInt(d_2460_v1_, 800)
                    nw401_[int(14)] = default__.fm2((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2469_i1_ in range(default__.abs(-466))])).set(default__.safeIndex(d_2460_v1_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('w') for d_2470_i1_ in range(default__.abs(-466))]))), d_2463_v4_), (self).f30, d_2464_v5_, ((d_2466_v7_)[(self).fm7(d_2458_v0_, d_2458_v0_, globalState)] if ((self).fm7(d_2458_v0_, d_2458_v0_, globalState)) in (d_2466_v7_) else True), globalState)
                    nw401_[int(15)] = 84
                    nw401_[int(16)] = d_2460_v1_
                    nw401_[int(17)] = (0) - (len(d_2462_v3_))
                    nw401_[int(18)] = (0) - ((d_2460_v1_) - (d_2460_v1_))
                    nw401_[int(19)] = d_2460_v1_
                    nw401_[int(20)] = d_2460_v1_
                    nw401_[int(21)] = d_2460_v1_
                    nw401_[int(22)] = (d_2460_v1_) - (d_2460_v1_)
                    nw401_[int(23)] = (d_2467_v8_).cardinality
                    nw401_[int(24)] = 23
                    d_2468_v9_ = nw401_
                    index394_ = default__.safeIndex(507, (d_2468_v9_).length(0))
                    (d_2468_v9_)[index394_] = 693
                    index395_ = default__.safeIndex(507, (d_2468_v9_).length(0))
                    (d_2468_v9_)[index395_] = (0) - (default__.safeDivisionInt((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")))), len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_2460_v1_]) for d_2471_i2_ in range(default__.abs(411))]))))
                    (globalState).f11 = True
                    (globalState).f0 = (len(d_2466_v7_)) - (d_2460_v1_)
                    d_2472_v10_: C8
                    nw402_ = C8()
                    nw402_.ctor__(d_2460_v1_)
                    d_2472_v10_ = nw402_
                    pass
            pass
        if d_2458_v0_:
            d_2473_v11_: _dafny.Array
            nw403_ = _dafny.Array(_dafny.Map({}), 7)
            d_2473_v11_ = nw403_
            index396_ = default__.safeIndex(462, (d_2473_v11_).length(0))
            def iife223_():
                coll91_ = _dafny.Map()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(667, 40):
                    d_2474_v12_: int = compr_91_
                    if ((667) <= (d_2474_v12_)) and ((d_2474_v12_) < (40)):
                        def iife224_():
                            coll92_ = _dafny.Map()
                            compr_92_: D4
                            for compr_92_ in (_dafny.Map({D4_DC10(_dafny.CodePoint('p')): 532})).keys.Elements:
                                d_2475_v13_: D4 = compr_92_
                                if (d_2475_v13_) in (_dafny.Map({D4_DC10(_dafny.CodePoint('p')): 532})):
                                    coll92_[d_2475_v13_] = -380
                            return _dafny.Map(coll92_)
                        coll91_[default__.safeDivisionInt(d_2474_v12_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True, d_2458_v0_, not(d_2458_v0_), d_2458_v0_]))).cardinality)] = D7_DC17(iife224_()
)
                return _dafny.Map(coll91_)
            (d_2473_v11_)[index396_] = iife223_()
            
            d_2476_v14_: _dafny.Seq
            d_2476_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ftnbci"))
            d_2477_v15_: str
            d_2477_v15_ = _dafny.CodePoint('m')
            d_2478_v16_: D4
            d_2478_v16_ = D4_DC10(d_2477_v15_)
            d_2479_v17_: int
            d_2479_v17_ = 584
            d_2480_v18_: _dafny.Map
            d_2480_v18_ = _dafny.Map({d_2478_v16_: d_2479_v17_})
            d_2481_v19_: D7
            d_2481_v19_ = D7_DC17(d_2480_v18_)
            d_2482_v20_: _dafny.Map
            d_2482_v20_ = _dafny.Map({len(d_2476_v14_): d_2481_v19_})
            index397_ = default__.safeIndex(462, (d_2473_v11_).length(0))
            (d_2473_v11_)[index397_] = d_2482_v20_
            (globalState).f24 = d_2479_v17_
            d_2477_v15_ = d_2477_v15_
            (globalState).f14 = (466) * (len((self).f30))
            d_2483_v21_: _dafny.Array
            def lambda104_(d_2484_v0_):
                def lambda105_(d_2485_i3_):
                    return d_2484_v0_

                return lambda105_

            init55_ = lambda104_(d_2458_v0_)
            nw404_ = _dafny.Array(None, 22)
            for i0_55_ in range(nw404_.length(0)):
                nw404_[i0_55_] = init55_(i0_55_)
            d_2483_v21_ = nw404_
            d_2486_v22_: _dafny.Seq
            d_2486_v22_ = _dafny.SeqWithoutIsStrInference([d_2483_v21_])
            d_2487_v23_: _dafny.Array
            nw405_ = _dafny.Array(_dafny.Map({}), 12)
            d_2487_v23_ = nw405_
            d_2488_v24_: _dafny.Array
            d_2488_v24_ = d_2487_v23_
            d_2489_v25_: _dafny.Map
            d_2489_v25_ = _dafny.Map({d_2479_v17_: (d_2488_v24_ if d_2458_v0_ else d_2488_v24_)})
            rhs435_ = (d_2479_v17_) <= ((d_2479_v17_) * (d_2479_v17_))
            rhs436_ = d_2486_v22_
            rhs437_ = ((len((d_2476_v14_).set(default__.safeIndex(d_2479_v17_, len(d_2476_v14_)), _dafny.CodePoint('p')))) - (d_2479_v17_) if d_2458_v0_ else (0) - (d_2479_v17_))
            rhs438_ = d_2489_v25_
            lhs358_ = globalState
            d_2458_v0_ = rhs435_
            d_2486_v22_ = rhs436_
            lhs358_.f24 = rhs437_
            d_2489_v25_ = rhs438_
        elif True:
            d_2490_v26_: int
            d_2490_v26_ = 679
            d_2491_v27_: _dafny.Seq
            d_2491_v27_ = _dafny.SeqWithoutIsStrInference([d_2490_v26_])
            d_2492_v28_: C1
            nw406_ = C1()
            nw406_.ctor__(((d_2491_v27_)[default__.safeIndex(d_2490_v26_, len(d_2491_v27_))]) * (d_2490_v26_), d_2490_v26_, (D1_DC3(d_2458_v0_) if d_2458_v0_ else (self).f28))
            d_2492_v28_ = nw406_
            d_2493_v29_: _dafny.Array
            nw407_ = _dafny.Array(_dafny.Map({}), 11)
            d_2493_v29_ = nw407_
            d_2494_v30_: _dafny.MultiSet
            d_2494_v30_ = _dafny.MultiSet([(d_2492_v28_).f41])
            d_2495_v31_: C0
            nw408_ = C0()
            nw408_.ctor__((self).f30)
            d_2495_v31_ = nw408_
            d_2496_v32_: _dafny.Map
            d_2496_v32_ = _dafny.Map({((d_2494_v30_)[d_2490_v26_] if (d_2490_v26_) in (d_2494_v30_) else d_2490_v26_): d_2495_v31_})
            d_2497_v33_: _dafny.Map
            d_2497_v33_ = _dafny.Map({d_2490_v26_: (0) - (len(d_2496_v32_))})
            index398_ = default__.safeIndex(758, (d_2493_v29_).length(0))
            (d_2493_v29_)[index398_] = d_2497_v33_
            index399_ = default__.safeIndex(758, (d_2493_v29_).length(0))
            (d_2493_v29_)[index399_] = d_2497_v33_
            (globalState).f27 = (d_2492_v28_.f42) >= ((d_2492_v28_.f42) + (d_2490_v26_))
            (globalState).f11 = d_2458_v0_
            source34_ = (self).f28
            if source34_.is_DC4:
                d_2498___mcc_h0_ = source34_.cf5
                d_2499___mcc_h1_ = source34_.cf6
                d_2500_cf6_ = d_2499___mcc_h1_
                d_2501_cf5_ = d_2498___mcc_h0_
                d_2502_v34_: _dafny.Array
                nw409_ = _dafny.Array(int(0), 19)
                d_2502_v34_ = nw409_
                d_2503_v35_: _dafny.Array
                nw410_ = _dafny.Array(None, 26)
                nw410_[int(0)] = d_2502_v34_
                nw410_[int(1)] = d_2502_v34_
                nw410_[int(2)] = d_2502_v34_
                nw410_[int(3)] = d_2502_v34_
                nw410_[int(4)] = d_2502_v34_
                nw410_[int(5)] = d_2502_v34_
                nw410_[int(6)] = d_2502_v34_
                nw410_[int(7)] = d_2502_v34_
                nw410_[int(8)] = d_2502_v34_
                nw410_[int(9)] = d_2502_v34_
                nw410_[int(10)] = d_2502_v34_
                nw410_[int(11)] = d_2502_v34_
                nw410_[int(12)] = d_2502_v34_
                nw410_[int(13)] = d_2502_v34_
                nw410_[int(14)] = d_2502_v34_
                nw410_[int(15)] = d_2502_v34_
                nw410_[int(16)] = d_2502_v34_
                nw410_[int(17)] = d_2502_v34_
                nw410_[int(18)] = d_2502_v34_
                nw410_[int(19)] = d_2502_v34_
                nw410_[int(20)] = d_2502_v34_
                nw410_[int(21)] = d_2502_v34_
                nw410_[int(22)] = d_2502_v34_
                nw410_[int(23)] = d_2502_v34_
                nw410_[int(24)] = d_2502_v34_
                nw410_[int(25)] = d_2502_v34_
                d_2503_v35_ = nw410_
                d_2504_v36_: _dafny.Map
                d_2504_v36_ = _dafny.Map({d_2503_v35_: d_2503_v35_})
                d_2505_v37_: _dafny.Seq
                d_2505_v37_ = _dafny.SeqWithoutIsStrInference([d_2503_v35_, d_2503_v35_])
                d_2506_v38_: _dafny.Seq
                d_2506_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "htrrxc"))
                d_2504_v36_ = (d_2504_v36_).set(d_2503_v35_, (d_2505_v37_)[default__.safeIndex(len(d_2506_v38_), len(d_2505_v37_))])
                d_2507_v39_: _dafny.Seq
                d_2507_v39_ = _dafny.SeqWithoutIsStrInference([d_2500_cf6_, d_2500_cf6_])
                d_2508_v40_: D4
                d_2508_v40_ = D4_DC11(d_2507_v39_)
                d_2508_v40_ = d_2508_v40_
                (globalState).f11 = (d_2494_v30_).ispropersubset(d_2494_v30_)
                index400_ = default__.safeIndex(40, (d_2502_v34_).length(0))
                (d_2502_v34_)[index400_] = d_2490_v26_
                index401_ = default__.safeIndex(40, (d_2502_v34_).length(0))
                (d_2502_v34_)[index401_] = (d_2491_v27_)[default__.safeIndex(len((d_2506_v38_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yhytnwlse")))), len(d_2491_v27_))]
            elif source34_.is_DC5:
                d_2509___mcc_h2_ = source34_.cf7
                d_2510___mcc_h3_ = source34_.cf8
                d_2511___mcc_h4_ = source34_.cf9
                d_2512___mcc_h5_ = source34_.cf10
                d_2513___mcc_h6_ = source34_.cf11
                d_2514_cf11_ = d_2513___mcc_h6_
                d_2515_cf10_ = d_2512___mcc_h5_
                d_2516_cf9_ = d_2511___mcc_h4_
                d_2517_cf8_ = d_2510___mcc_h3_
                d_2518_cf7_ = d_2509___mcc_h2_
                index402_ = default__.safeIndex(758, (d_2493_v29_).length(0))
                (d_2493_v29_)[index402_] = ((d_2493_v29_)[default__.safeIndex(758, (d_2493_v29_).length(0))]).set((d_2492_v28_).f41, d_2516_cf9_)
                d_2519_v41_: _dafny.Seq
                d_2519_v41_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xwnb"))
                d_2520_v42_: str
                d_2520_v42_ = _dafny.CodePoint('m')
                r1 = (d_2519_v41_).set(default__.safeIndex((d_2492_v28_).f41, len(d_2519_v41_)), d_2520_v42_)
                d_2519_v41_ = d_2519_v41_
                d_2521_v43_: _dafny.Array
                nw411_ = _dafny.Array(_dafny.Map({}), 22)
                d_2521_v43_ = nw411_
                d_2522_v44_: _dafny.Set
                d_2522_v44_ = _dafny.Set({len(d_2519_v41_), (d_2492_v28_).f41, d_2490_v26_})
                d_2523_v45_: _dafny.Map
                d_2523_v45_ = _dafny.Map({d_2491_v27_: d_2522_v44_})
                index403_ = default__.safeIndex(620, (d_2521_v43_).length(0))
                def iife225_():
                    coll93_ = _dafny.Set()
                    compr_93_: int
                    for compr_93_ in (d_2494_v30_).Elements:
                        d_2524_v46_: int = compr_93_
                        if (d_2524_v46_) in (d_2494_v30_):
                            coll93_ = coll93_.union(_dafny.Set([default__.safeModuloInt(d_2524_v46_, (D19_DC46(-871, False, False, False)).cf84)]))
                    return _dafny.Set(coll93_)
                (d_2521_v43_)[index403_] = (d_2523_v45_).set(d_2491_v27_, _dafny.Set({(d_2492_v28_).f41, len(iife225_()
                )}))
                index404_ = default__.safeIndex(620, (d_2521_v43_).length(0))
                (d_2521_v43_)[index404_] = ((d_2523_v45_) | (d_2523_v45_)) | (d_2523_v45_)
            elif True:
                d_2525___mcc_h7_ = source34_.cf4
                d_2526_cf4_ = d_2525___mcc_h7_
                d_2527_v47_: _dafny.Array
                def lambda106_(d_2528_i4_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oyofknkne"))

                init56_ = lambda106_
                nw412_ = _dafny.Array(None, 27)
                for i0_56_ in range(nw412_.length(0)):
                    nw412_[i0_56_] = init56_(i0_56_)
                d_2527_v47_ = nw412_
                d_2529_v48_: C4
                nw413_ = C4()
                nw413_.ctor__(d_2526_cf4_, d_2527_v47_)
                d_2529_v48_ = nw413_
                d_2530_v51_: _dafny.Seq
                d_2530_v51_ = _dafny.SeqWithoutIsStrInference([d_2458_v0_])
                d_2531_v52_: _dafny.Set
                d_2531_v52_ = _dafny.Set({d_2530_v51_, d_2530_v51_})
                d_2532_v53_: C2
                nw414_ = C2()
                def iife226_():
                    def iife228_():
                        coll96_ = _dafny.Map()
                        compr_96_: _dafny.Seq
                        for compr_96_ in (d_2531_v52_).Elements:
                            d_2533_v50_: _dafny.Seq = compr_96_
                            if (d_2533_v50_) in (d_2531_v52_):
                                coll96_[d_2533_v50_] = len(_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_2492_v28_).f41: (d_2529_v48_).f37}) for d_2534_i5_ in range(default__.abs(983))])): d_2497_v33_}))}))
                        return _dafny.Map(coll96_)
                    coll94_ = _dafny.Map()
                    def iife227_():
                        coll95_ = _dafny.Map()
                        compr_95_: _dafny.Seq
                        for compr_95_ in (d_2531_v52_).Elements:
                            d_2533_v50_: _dafny.Seq = compr_95_
                            if (d_2533_v50_) in (d_2531_v52_):
                                coll95_[d_2533_v50_] = len(_dafny.Set({len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_2492_v28_).f41: (d_2529_v48_).f37}) for d_2534_i5_ in range(default__.abs(983))])): d_2497_v33_}))}))
                        return _dafny.Map(coll95_)
                    compr_94_: _dafny.Seq
                    for compr_94_ in (iife227_()
                    ).keys.Elements:
                        d_2535_v49_: _dafny.Seq = compr_94_
                        if (d_2535_v49_) in (iife228_()
                        ):
                            coll94_[d_2535_v49_] = len(_dafny.Map({d_2492_v28_.f42: (d_2492_v28_).f41}))
                    return _dafny.Map(coll94_)
                nw414_.ctor__(iife226_()
                , ((self).f30) | ((d_2495_v31_).f43), D1_DC3((d_2529_v48_).f37))
                d_2532_v53_ = nw414_
                d_2536_v54_: _dafny.Seq
                d_2536_v54_ = _dafny.SeqWithoutIsStrInference([d_2494_v30_])
                (globalState).f0 = default__.safeDivisionInt(len(d_2536_v54_), d_2492_v28_.f42)
                d_2537_v55_: _dafny.Array
                nw415_ = _dafny.Array(False, 28)
                d_2537_v55_ = nw415_
                index405_ = default__.safeIndex(237, (d_2537_v55_).length(0))
                (d_2537_v55_)[index405_] = d_2458_v0_
                index406_ = default__.safeIndex(237, (d_2537_v55_).length(0))
                (d_2537_v55_)[index406_] = (d_2530_v51_)[default__.safeIndex((d_2492_v28_).f41, len(d_2530_v51_))]
        d_2538_v56_: _dafny.Array
        def lambda107_(d_2539_i6_):
            return _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "w"))), (_dafny.MultiSet([False])).cardinality, -875])

        init57_ = lambda107_
        nw416_ = _dafny.Array(None, 29)
        for i0_57_ in range(nw416_.length(0)):
            nw416_[i0_57_] = init57_(i0_57_)
        d_2538_v56_ = nw416_
        d_2540_v57_: int
        d_2540_v57_ = 935
        d_2541_v58_: _dafny.MultiSet
        d_2541_v58_ = _dafny.MultiSet([d_2540_v57_])
        d_2542_v59_: _dafny.Seq
        d_2542_v59_ = _dafny.SeqWithoutIsStrInference([(d_2541_v58_).cardinality, d_2540_v57_])
        index407_ = default__.safeIndex(744, (d_2538_v56_).length(0))
        (d_2538_v56_)[index407_] = (d_2542_v59_) + (d_2542_v59_)
        index408_ = default__.safeIndex(744, (d_2538_v56_).length(0))
        (d_2538_v56_)[index408_] = (d_2542_v59_) + (d_2542_v59_)
        d_2543_v60_: C1
        nw417_ = C1()
        nw417_.ctor__(d_2540_v57_, (d_2540_v57_) + (d_2540_v57_), (self).f28)
        d_2543_v60_ = nw417_
        d_2544_v61_: _dafny.Map
        d_2544_v61_ = _dafny.Map({d_2458_v0_: 549})
        (globalState).f14 = len(d_2544_v61_)
        d_2545_v62_: D0
        d_2545_v62_ = D0_DC1(d_2540_v57_)
        source35_ = d_2545_v62_
        if source35_.is_DC1:
            d_2546___mcc_h8_ = source35_.cf1
            d_2547_cf1_ = d_2546___mcc_h8_
            d_2548_v63_: _dafny.Seq
            d_2548_v63_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hivp"))
            d_2549_v64_: D19
            d_2549_v64_ = D19_DC46(len(d_2548_v63_), d_2458_v0_, d_2458_v0_, d_2458_v0_)
            source36_ = d_2549_v64_
            if source36_.is_DC45:
                d_2550___mcc_h12_ = source36_.cf82
                d_2551___mcc_h13_ = source36_.cf83
                d_2552_cf83_ = d_2551___mcc_h13_
                d_2553_cf82_ = d_2550___mcc_h12_
                (globalState).f1 = d_2552_cf83_
                d_2458_v0_ = (self).fm7(d_2552_cf83_, False, globalState)
                d_2554_v65_: _dafny.Seq
                d_2554_v65_ = _dafny.SeqWithoutIsStrInference([d_2552_cf83_, d_2552_cf83_, d_2552_cf83_])
                d_2555_v66_: str
                d_2555_v66_ = _dafny.CodePoint('k')
                d_2556_v67_: _dafny.Map
                d_2556_v67_ = _dafny.Map({not(d_2552_cf83_): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y'), d_2555_v66_])})
                d_2557_v68_: _dafny.Map
                d_2557_v68_ = _dafny.Map({d_2555_v66_: len(d_2548_v63_)})
                d_2558_v69_: D8
                d_2558_v69_ = D8_DC20(406, -749)
                d_2559_v70_: _dafny.MultiSet
                d_2559_v70_ = _dafny.MultiSet([False, d_2458_v0_])
                d_2560_v71_: _dafny.Array
                nw418_ = _dafny.Array(None, 18)
                nw418_[int(0)] = (d_2543_v60_).f41
                nw418_[int(1)] = len(d_2554_v65_)
                nw418_[int(2)] = d_2543_v60_.f42
                nw418_[int(3)] = d_2540_v57_
                nw418_[int(4)] = d_2543_v60_.f42
                nw418_[int(5)] = default__.fm2((d_2548_v63_).set(default__.safeIndex(len(((d_2556_v67_)[d_2552_cf83_] if (d_2552_cf83_) in (d_2556_v67_) else d_2548_v63_)), len(d_2548_v63_)), d_2555_v66_), (self).f30, _dafny.Map({d_2552_cf83_: d_2555_v66_}), d_2458_v0_, globalState)
                nw418_[int(6)] = (-440) * ((0) - ((d_2543_v60_).f41))
                nw418_[int(7)] = (d_2543_v60_).f41
                nw418_[int(8)] = len(d_2557_v68_)
                nw418_[int(9)] = (d_2543_v60_).f41
                nw418_[int(10)] = 609
                nw418_[int(11)] = (d_2543_v60_).f41
                nw418_[int(12)] = d_2540_v57_
                nw418_[int(13)] = (0) - ((188 if d_2458_v0_ else (0) - ((d_2558_v69_).cf31)))
                nw418_[int(14)] = len((d_2554_v65_) + (d_2554_v65_))
                nw418_[int(15)] = d_2540_v57_
                nw418_[int(16)] = ((d_2559_v70_)[d_2552_cf83_] if (d_2552_cf83_) in (d_2559_v70_) else d_2553_cf82_)
                nw418_[int(17)] = (0) - (len(d_2548_v63_))
                d_2560_v71_ = nw418_
                index409_ = default__.safeIndex(969, (d_2560_v71_).length(0))
                (d_2560_v71_)[index409_] = d_2540_v57_
                index410_ = default__.safeIndex(969, (d_2560_v71_).length(0))
                (d_2560_v71_)[index410_] = 979
                (globalState).f11 = (d_2554_v65_)[default__.safeIndex((d_2543_v60_).f41, len(d_2554_v65_))]
            elif source36_.is_DC46:
                d_2561___mcc_h14_ = source36_.cf84
                d_2562___mcc_h15_ = source36_.cf85
                d_2563___mcc_h16_ = source36_.cf86
                d_2564___mcc_h17_ = source36_.cf87
                d_2565_cf87_ = d_2564___mcc_h17_
                d_2566_cf86_ = d_2563___mcc_h16_
                d_2567_cf85_ = d_2562___mcc_h15_
                d_2568_cf84_ = d_2561___mcc_h14_
                d_2569_v72_: _dafny.Map
                d_2569_v72_ = _dafny.Map({d_2458_v0_: d_2458_v0_})
                d_2570_v73_: _dafny.Map
                d_2570_v73_ = _dafny.Map({d_2569_v72_: d_2547_cf1_})
                d_2571_v74_: str
                d_2571_v74_ = _dafny.CodePoint('t')
                (d_2543_v60_).f42 = ((d_2570_v73_)[(default__.fm48(d_2571_v74_, d_2547_cf1_, globalState)) | (d_2569_v72_)] if ((default__.fm48(d_2571_v74_, d_2547_cf1_, globalState)) | (d_2569_v72_)) in (d_2570_v73_) else default__.safeModuloInt((d_2543_v60_).f41, (d_2543_v60_).f41))
                d_2572_v75_: _dafny.Set
                d_2572_v75_ = _dafny.Set({d_2541_v58_})
                d_2573_v76_: _dafny.Map
                d_2573_v76_ = _dafny.Map({d_2572_v75_: (default__.fm3(globalState)) + (d_2548_v63_)})
                d_2547_cf1_ = len(((d_2573_v76_)[d_2572_v75_] if (d_2572_v75_) in (d_2573_v76_) else d_2548_v63_))
                d_2574_v77_: _dafny.Array
                nw419_ = _dafny.Array(None, 20)
                nw419_[int(0)] = True
                nw419_[int(1)] = (len(_dafny.SeqWithoutIsStrInference([d_2571_v74_ for d_2575_i7_ in range(default__.abs(525))]))) == ((d_2543_v60_).f41)
                nw419_[int(2)] = not(d_2566_cf86_)
                nw419_[int(3)] = ((d_2543_v60_).f41) == (-999)
                nw419_[int(4)] = not(d_2566_cf86_)
                nw419_[int(5)] = d_2567_cf85_
                nw419_[int(6)] = d_2565_cf87_
                nw419_[int(7)] = (d_2543_v60_).fm7(d_2458_v0_, False, globalState)
                nw419_[int(8)] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tkyq")))) >= ((d_2543_v60_).f41)
                nw419_[int(9)] = d_2458_v0_
                nw419_[int(10)] = d_2565_cf87_
                nw419_[int(11)] = d_2458_v0_
                nw419_[int(12)] = d_2458_v0_
                nw419_[int(13)] = d_2565_cf87_
                nw419_[int(14)] = d_2566_cf86_
                nw419_[int(15)] = not (d_2567_cf85_) or (d_2565_cf87_)
                nw419_[int(16)] = d_2458_v0_
                nw419_[int(17)] = (d_2568_cf84_) >= ((0) - ((d_2543_v60_).f41))
                nw419_[int(18)] = d_2567_cf85_
                nw419_[int(19)] = d_2565_cf87_
                d_2574_v77_ = nw419_
                index411_ = default__.safeIndex(317, (d_2574_v77_).length(0))
                (d_2574_v77_)[index411_] = True
                index412_ = default__.safeIndex(317, (d_2574_v77_).length(0))
                (d_2574_v77_)[index412_] = not((default__.safeModuloInt((d_2543_v60_).f41, len(d_2569_v72_))) == (587))
                (globalState).f27 = (d_2567_cf85_ if (d_2574_v77_)[default__.safeIndex(317, (d_2574_v77_).length(0))] else not((not(d_2567_cf85_)) and (d_2565_cf87_)))
            elif source36_.is_DC44:
                d_2576___mcc_h18_ = source36_.cf81
                d_2577_cf81_ = d_2576___mcc_h18_
                d_2578_v78_: _dafny.Array
                nw420_ = _dafny.Array(D1.default()(), 6)
                d_2578_v78_ = nw420_
                index413_ = default__.safeIndex(209, (d_2578_v78_).length(0))
                (d_2578_v78_)[index413_] = default__.fm29(d_2540_v57_, globalState)
                d_2579_v79_: _dafny.Seq
                d_2579_v79_ = _dafny.SeqWithoutIsStrInference([not(d_2458_v0_), d_2458_v0_])
                d_2580_v80_: D1
                d_2580_v80_ = D1_DC5(d_2458_v0_, d_2458_v0_, (d_2543_v60_.f42 if d_2458_v0_ else (d_2541_v58_).cardinality), d_2458_v0_, (_dafny.MultiSet([False])).ispropersubset(default__.fm25(len(d_2548_v63_), default__.fm4(globalState), d_2579_v79_, 799, globalState)))
                index414_ = default__.safeIndex(209, (d_2578_v78_).length(0))
                (d_2578_v78_)[index414_] = d_2580_v80_
                d_2581_v81_: C0
                nw421_ = C0()
                nw421_.ctor__((self).f30)
                d_2581_v81_ = nw421_
                d_2582_v82_: _dafny.Array
                nw422_ = _dafny.Array(False, 26)
                d_2582_v82_ = nw422_
                d_2583_v83_: _dafny.Map
                d_2583_v83_ = _dafny.Map({d_2543_v60_: d_2543_v60_.f42})
                d_2584_v84_: _dafny.Map
                d_2584_v84_ = _dafny.Map({not(d_2458_v0_): default__.fm5(len(d_2583_v83_), d_2458_v0_, globalState)})
                d_2585_v85_: _dafny.MultiSet
                d_2585_v85_ = _dafny.MultiSet([d_2458_v0_, d_2458_v0_, d_2458_v0_])
                d_2586_v86_: _dafny.Map
                d_2586_v86_ = _dafny.Map({d_2547_cf1_: d_2547_cf1_})
                d_2587_v87_: _dafny.Array
                nw423_ = _dafny.Array(None, 29)
                nw423_[int(0)] = d_2543_v60_.f42
                nw423_[int(1)] = d_2547_cf1_
                nw423_[int(2)] = d_2543_v60_.f42
                nw423_[int(3)] = d_2543_v60_.f42
                nw423_[int(4)] = d_2543_v60_.f42
                nw423_[int(5)] = default__.fm2(d_2548_v63_, _dafny.Set({d_2458_v0_, d_2458_v0_}), d_2584_v84_, d_2458_v0_, globalState)
                nw423_[int(6)] = len(default__.fm49(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfrojel"))), d_2540_v57_, d_2458_v0_, d_2540_v57_, globalState))
                nw423_[int(7)] = (d_2585_v85_).cardinality
                nw423_[int(8)] = d_2543_v60_.f42
                nw423_[int(9)] = d_2543_v60_.f42
                nw423_[int(10)] = d_2543_v60_.f42
                nw423_[int(11)] = (d_2543_v60_).f41
                nw423_[int(12)] = (d_2543_v60_).f41
                nw423_[int(13)] = d_2543_v60_.f42
                nw423_[int(14)] = 421
                nw423_[int(15)] = (0) - (d_2543_v60_.f42)
                nw423_[int(16)] = (d_2543_v60_).f41
                nw423_[int(17)] = d_2540_v57_
                nw423_[int(18)] = d_2540_v57_
                nw423_[int(19)] = (self).fm12(globalState)
                nw423_[int(20)] = d_2540_v57_
                nw423_[int(21)] = (d_2543_v60_).f41
                nw423_[int(22)] = (d_2543_v60_).f41
                nw423_[int(23)] = -824
                nw423_[int(24)] = d_2543_v60_.f42
                nw423_[int(25)] = len(d_2586_v86_)
                nw423_[int(26)] = d_2540_v57_
                nw423_[int(27)] = (d_2543_v60_).f41
                nw423_[int(28)] = (d_2585_v85_).cardinality
                d_2587_v87_ = nw423_
                d_2588_v88_: _dafny.Seq
                d_2588_v88_ = _dafny.SeqWithoutIsStrInference([d_2584_v84_])
                default__.m0(d_2582_v82_, d_2587_v87_, (0) - (default__.fm2(d_2548_v63_, (self).f30, (d_2588_v88_)[default__.safeIndex(len(d_2542_v59_), len(d_2588_v88_))], d_2458_v0_, globalState)), (d_2543_v60_).f41, globalState)
                d_2589_v89_: C8
                nw424_ = C8()
                nw424_.ctor__(d_2543_v60_.f42)
                d_2589_v89_ = nw424_
            elif True:
                d_2590___mcc_h19_ = source36_.cf88
                d_2591_cf88_ = d_2590___mcc_h19_
                (globalState).f24 = (d_2543_v60_).f41
                d_2547_cf1_ = ((d_2543_v60_).f41 if d_2458_v0_ else (d_2543_v60_).f41)
                d_2592_v90_: _dafny.Set
                d_2592_v90_ = _dafny.Set({d_2540_v57_, -508})
                d_2592_v90_ = d_2592_v90_
                d_2593_v91_: _dafny.Array
                nw425_ = _dafny.Array(None, 5)
                nw425_[int(0)] = d_2458_v0_
                nw425_[int(1)] = d_2458_v0_
                nw425_[int(2)] = (d_2592_v90_).ispropersubset(d_2592_v90_)
                nw425_[int(3)] = d_2458_v0_
                nw425_[int(4)] = d_2458_v0_
                d_2593_v91_ = nw425_
                index415_ = default__.safeIndex(686, (d_2593_v91_).length(0))
                (d_2593_v91_)[index415_] = d_2458_v0_
                index416_ = default__.safeIndex(686, (d_2593_v91_).length(0))
                (d_2593_v91_)[index416_] = (d_2548_v63_) < (d_2548_v63_)
            (globalState).f27 = True
            d_2541_v58_ = _dafny.MultiSet((d_2538_v56_)[default__.safeIndex(744, (d_2538_v56_).length(0))])
            d_2594_v92_: _dafny.Array
            nw426_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_2594_v92_ = nw426_
            d_2595_v93_: _dafny.Array
            nw427_ = _dafny.Array(int(0), 28)
            d_2595_v93_ = nw427_
            index417_ = default__.safeIndex(66, (d_2594_v92_).length(0))
            (d_2594_v92_)[index417_] = d_2595_v93_
            d_2596_v94_: _dafny.Seq
            d_2596_v94_ = _dafny.SeqWithoutIsStrInference([d_2458_v0_, d_2458_v0_, d_2458_v0_, d_2458_v0_, d_2458_v0_])
            d_2597_v95_: str
            d_2597_v95_ = _dafny.CodePoint('f')
            d_2598_v96_: _dafny.Seq
            d_2598_v96_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({d_2458_v0_, d_2458_v0_})])
            d_2599_v97_: _dafny.Array
            nw428_ = _dafny.Array(None, 29)
            nw428_[int(0)] = d_2458_v0_
            nw428_[int(1)] = d_2458_v0_
            nw428_[int(2)] = False
            nw428_[int(3)] = (d_2596_v94_)[default__.safeIndex(d_2540_v57_, len(d_2596_v94_))]
            nw428_[int(4)] = not((not(d_2458_v0_)) or (d_2458_v0_))
            nw428_[int(5)] = (d_2597_v95_) not in (d_2548_v63_)
            nw428_[int(6)] = d_2458_v0_
            nw428_[int(7)] = (d_2458_v0_) == (default__.fm4(globalState))
            nw428_[int(8)] = (d_2458_v0_ if False else d_2458_v0_)
            nw428_[int(9)] = d_2458_v0_
            nw428_[int(10)] = (d_2543_v60_.f42) < ((d_2543_v60_).f41)
            nw428_[int(11)] = d_2458_v0_
            nw428_[int(12)] = d_2458_v0_
            nw428_[int(13)] = d_2458_v0_
            nw428_[int(14)] = d_2458_v0_
            nw428_[int(15)] = d_2458_v0_
            nw428_[int(16)] = d_2458_v0_
            nw428_[int(17)] = (d_2547_cf1_) <= ((d_2543_v60_).f41)
            nw428_[int(18)] = d_2458_v0_
            nw428_[int(19)] = (d_2543_v60_.f42) == (d_2547_cf1_)
            nw428_[int(20)] = d_2458_v0_
            nw428_[int(21)] = (d_2548_v63_) <= ((d_2548_v63_).set(default__.safeIndex(len((d_2538_v56_)[default__.safeIndex(744, (d_2538_v56_).length(0))]), len(d_2548_v63_)), d_2597_v95_))
            nw428_[int(22)] = d_2458_v0_
            nw428_[int(23)] = d_2458_v0_
            nw428_[int(24)] = (d_2598_v96_) < (default__.fm50(False, d_2458_v0_, d_2458_v0_, globalState))
            nw428_[int(25)] = ((self).f30).ispropersubset((self).f30)
            nw428_[int(26)] = not (d_2458_v0_) or (not(d_2458_v0_))
            nw428_[int(27)] = d_2458_v0_
            nw428_[int(28)] = d_2458_v0_
            d_2599_v97_ = nw428_
            index418_ = default__.safeIndex(197, (d_2599_v97_).length(0))
            (d_2599_v97_)[index418_] = (d_2596_v94_)[default__.safeIndex(-544, len(d_2596_v94_))]
            d_2600_v98_: _dafny.Seq
            d_2600_v98_ = _dafny.SeqWithoutIsStrInference([d_2595_v93_, d_2595_v93_, d_2595_v93_])
            index419_ = default__.safeIndex(66, (d_2594_v92_).length(0))
            index420_ = default__.safeIndex(197, (d_2599_v97_).length(0))
            rhs439_ = (d_2600_v98_)[default__.safeIndex(d_2543_v60_.f42, len(d_2600_v98_))]
            rhs440_ = d_2458_v0_
            rhs441_ = not (d_2458_v0_) or (d_2458_v0_)
            lhs359_ = d_2594_v92_
            lhs360_ = default__.safeIndex(66, (d_2594_v92_).length(0))
            lhs361_ = globalState
            lhs362_ = d_2599_v97_
            lhs363_ = default__.safeIndex(197, (d_2599_v97_).length(0))
            lhs359_[lhs360_] = rhs439_
            lhs361_.f1 = rhs440_
            lhs362_[lhs363_] = rhs441_
        elif source35_.is_DC2:
            d_2601___mcc_h9_ = source35_.cf2
            d_2602___mcc_h10_ = source35_.cf3
            d_2603_cf3_ = d_2602___mcc_h10_
            d_2604_cf2_ = d_2601___mcc_h9_
            d_2605_v99_: _dafny.Set
            d_2605_v99_ = _dafny.Set({(d_2543_v60_).f41, (d_2541_v58_).cardinality, (d_2543_v60_).f41})
            d_2603_cf3_ = (len(d_2605_v99_)) + (default__.safeModuloInt(-627, (d_2543_v60_).f41))
            d_2606_v100_: _dafny.Seq
            d_2606_v100_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmck"))
            d_2607_v101_: _dafny.Map
            d_2607_v101_ = _dafny.Map({len(d_2606_v100_): d_2458_v0_})
            if (d_2543_v60_).fm7((len(d_2607_v101_)) != ((d_2543_v60_).f41), d_2458_v0_, globalState):
                (globalState).f14 = ((d_2543_v60_).f41) - (d_2604_cf2_)
                d_2608_v102_: D17
                d_2608_v102_ = D17_DC38(_dafny.MultiSet([d_2458_v0_]))
                d_2609_v103_: _dafny.MultiSet
                d_2609_v103_ = _dafny.MultiSet([d_2458_v0_])
                d_2610_v104_: _dafny.MultiSet
                d_2610_v104_ = _dafny.MultiSet([d_2608_v102_, D17_DC38(d_2609_v103_), d_2608_v102_, d_2608_v102_, D17_DC38(d_2609_v103_)])
                d_2611_v105_: _dafny.Array
                nw429_ = _dafny.Array(None, 29)
                nw429_[int(0)] = d_2543_v60_.f42
                nw429_[int(1)] = d_2543_v60_.f42
                nw429_[int(2)] = ((d_2541_v58_)[d_2604_cf2_] if (d_2604_cf2_) in (d_2541_v58_) else d_2543_v60_.f42)
                nw429_[int(3)] = d_2543_v60_.f42
                nw429_[int(4)] = d_2604_cf2_
                nw429_[int(5)] = (d_2543_v60_).f41
                nw429_[int(6)] = (d_2541_v58_).cardinality
                nw429_[int(7)] = d_2543_v60_.f42
                nw429_[int(8)] = (d_2543_v60_).f41
                nw429_[int(9)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "unfuer")))
                nw429_[int(10)] = (d_2610_v104_).cardinality
                nw429_[int(11)] = d_2543_v60_.f42
                nw429_[int(12)] = d_2604_cf2_
                nw429_[int(13)] = (d_2541_v58_).cardinality
                nw429_[int(14)] = len(d_2606_v100_)
                nw429_[int(15)] = 394
                nw429_[int(16)] = d_2604_cf2_
                nw429_[int(17)] = d_2543_v60_.f42
                nw429_[int(18)] = (0) - ((d_2543_v60_).f41)
                nw429_[int(19)] = (self).fm12(globalState)
                nw429_[int(20)] = d_2603_cf3_
                nw429_[int(21)] = (d_2543_v60_).f41
                nw429_[int(22)] = d_2604_cf2_
                nw429_[int(23)] = d_2543_v60_.f42
                nw429_[int(24)] = (d_2543_v60_).f41
                nw429_[int(25)] = d_2543_v60_.f42
                nw429_[int(26)] = d_2604_cf2_
                nw429_[int(27)] = (d_2543_v60_).f41
                nw429_[int(28)] = (d_2541_v58_).cardinality
                d_2611_v105_ = nw429_
                d_2612_v106_: _dafny.Map
                d_2612_v106_ = _dafny.Map({d_2458_v0_: d_2611_v105_})
                d_2613_v107_: _dafny.Array
                nw430_ = _dafny.Array(None, 14)
                nw430_[int(0)] = d_2611_v105_
                nw430_[int(1)] = d_2611_v105_
                nw430_[int(2)] = d_2611_v105_
                nw430_[int(3)] = d_2611_v105_
                nw430_[int(4)] = d_2611_v105_
                nw430_[int(5)] = ((d_2612_v106_)[d_2458_v0_] if (d_2458_v0_) in (d_2612_v106_) else d_2611_v105_)
                nw430_[int(6)] = d_2611_v105_
                nw430_[int(7)] = d_2611_v105_
                nw430_[int(8)] = d_2611_v105_
                nw430_[int(9)] = d_2611_v105_
                nw430_[int(10)] = d_2611_v105_
                nw430_[int(11)] = d_2611_v105_
                nw430_[int(12)] = d_2611_v105_
                nw430_[int(13)] = d_2611_v105_
                d_2613_v107_ = nw430_
                d_2614_v108_: _dafny.Map
                d_2614_v108_ = _dafny.Map({d_2613_v107_: d_2605_v99_})
                (globalState).f27 = (((d_2614_v108_)[d_2613_v107_] if (d_2613_v107_) in (d_2614_v108_) else _dafny.Set({d_2603_cf3_}))) == (_dafny.Set({(d_2543_v60_).f41, d_2604_cf2_, (d_2543_v60_).f41}))
                (globalState).f14 = d_2543_v60_.f42
                (globalState).f11 = not (d_2458_v0_) or (d_2458_v0_)
                d_2615_v109_: str
                d_2615_v109_ = _dafny.CodePoint('p')
                d_2616_v110_: D4
                d_2616_v110_ = D4_DC10(d_2615_v109_)
                d_2615_v109_ = (d_2616_v110_).cf19
            elif True:
                d_2617_v111_: _dafny.Array
                nw431_ = _dafny.Array(int(0), 25)
                d_2617_v111_ = nw431_
                d_2618_v112_: _dafny.Map
                d_2618_v112_ = _dafny.Map({d_2458_v0_: d_2458_v0_})
                d_2619_v113_: _dafny.MultiSet
                d_2619_v113_ = _dafny.MultiSet([d_2618_v112_, d_2618_v112_, d_2618_v112_, d_2618_v112_])
                d_2620_v114_: _dafny.Map
                d_2620_v114_ = _dafny.Map({(d_2619_v113_).issubset(d_2619_v113_): d_2617_v111_})
                rhs442_ = ((d_2620_v114_)[not (d_2458_v0_) or (d_2458_v0_)] if (not (d_2458_v0_) or (d_2458_v0_)) in (d_2620_v114_) else d_2617_v111_)
                rhs443_ = (d_2543_v60_).f41
                rhs444_ = (d_2543_v60_).f41
                rhs445_ = (0) - ((d_2543_v60_).f41)
                lhs364_ = globalState
                lhs365_ = globalState
                lhs366_ = globalState
                d_2617_v111_ = rhs442_
                lhs364_.f14 = rhs443_
                lhs365_.f14 = rhs444_
                lhs366_.f14 = rhs445_
                d_2621_v115_: str
                d_2621_v115_ = _dafny.CodePoint('w')
                d_2622_v116_: int
                out19_: int
                out19_ = (d_2543_v60_).m2(d_2458_v0_, d_2621_v115_, globalState)
                d_2622_v116_ = out19_
                (globalState).f14 = d_2543_v60_.f42
                d_2623_v117_: D19
                d_2623_v117_ = D19_DC45(d_2543_v60_.f42, d_2458_v0_)
                d_2624_v120_: _dafny.Array
                nw432_ = _dafny.Array(None, 20)
                nw432_[int(0)] = (d_2623_v117_).cf83
                nw432_[int(1)] = (d_2458_v0_) and (False)
                def iife229_():
                    coll97_ = _dafny.Map()
                    compr_97_: int
                    for compr_97_ in _dafny.IntegerRange(124, -212):
                        d_2625_v118_: int = compr_97_
                        if ((124) <= (d_2625_v118_)) and ((d_2625_v118_) < (-212)):
                            coll97_[default__.safeModuloInt(d_2625_v118_, (d_2543_v60_).f41)] = d_2458_v0_
                    return _dafny.Map(coll97_)
                nw432_[int(2)] = (len(iife229_()
                )) == ((d_2543_v60_).f41)
                nw432_[int(3)] = d_2458_v0_
                nw432_[int(4)] = d_2458_v0_
                nw432_[int(5)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mqgq"))) == (_dafny.SeqWithoutIsStrInference([d_2621_v115_ for d_2626_i8_ in range(default__.abs(543))]))
                nw432_[int(6)] = d_2458_v0_
                nw432_[int(7)] = default__.fm4(globalState)
                nw432_[int(8)] = d_2458_v0_
                nw432_[int(9)] = d_2458_v0_
                nw432_[int(10)] = d_2458_v0_
                nw432_[int(11)] = d_2458_v0_
                nw432_[int(12)] = d_2458_v0_
                nw432_[int(13)] = d_2458_v0_
                nw432_[int(14)] = (False if d_2458_v0_ else d_2458_v0_)
                nw432_[int(15)] = d_2458_v0_
                def iife230_():
                    coll98_ = _dafny.Set()
                    compr_98_: int
                    for compr_98_ in _dafny.IntegerRange(294, 458):
                        d_2627_v119_: int = compr_98_
                        if ((294) <= (d_2627_v119_)) and ((d_2627_v119_) < (458)):
                            coll98_ = coll98_.union(_dafny.Set([default__.safeModuloInt(d_2627_v119_, (d_2543_v60_).f41)]))
                    return _dafny.Set(coll98_)
                nw432_[int(16)] = (iife230_()
                ) == (d_2605_v99_)
                nw432_[int(17)] = False
                nw432_[int(18)] = d_2458_v0_
                nw432_[int(19)] = d_2458_v0_
                d_2624_v120_ = nw432_
                index421_ = default__.safeIndex(965, (d_2624_v120_).length(0))
                (d_2624_v120_)[index421_] = d_2458_v0_
                index422_ = default__.safeIndex(965, (d_2624_v120_).length(0))
                (d_2624_v120_)[index422_] = d_2458_v0_
                d_2628_v121_: _dafny.Seq
                d_2628_v121_ = _dafny.SeqWithoutIsStrInference([d_2458_v0_])
                d_2629_v122_: D10
                d_2629_v122_ = D10_DC25(d_2617_v111_, ((d_2606_v100_)[default__.safeIndex(d_2603_cf3_, len(d_2606_v100_))]) in (d_2606_v100_), (d_2628_v121_) + (d_2628_v121_), default__.safeModuloInt(d_2603_cf3_, (d_2543_v60_).f41))
                d_2629_v122_ = d_2629_v122_
            if not(d_2458_v0_):
                d_2630_v123_: _dafny.Map
                d_2630_v123_ = _dafny.Map({d_2458_v0_: ((self).f30).intersection(_dafny.Set({d_2458_v0_}))})
                d_2631_v124_: str
                d_2631_v124_ = _dafny.CodePoint('r')
                d_2632_v125_: _dafny.Map
                d_2632_v125_ = _dafny.Map({d_2631_v124_: 329})
                d_2633_v126_: D10
                d_2633_v126_ = D10_DC24((self).f30)
                d_2630_v123_ = _dafny.Map({((D24_DC61(False, d_2458_v0_, d_2540_v57_)).cf118) == (len(d_2632_v125_)): (d_2633_v126_).cf38})
                d_2634_v127_: _dafny.MultiSet
                d_2634_v127_ = _dafny.MultiSet([d_2458_v0_, d_2458_v0_, d_2458_v0_])
                d_2635_v128_: _dafny.Array
                def lambda108_(d_2636_i9_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fxlybiaqv"))

                init58_ = lambda108_
                nw433_ = _dafny.Array(None, 8)
                for i0_58_ in range(nw433_.length(0)):
                    nw433_[i0_58_] = init58_(i0_58_)
                d_2635_v128_ = nw433_
                d_2637_v129_: C4
                nw434_ = C4()
                nw434_.ctor__((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jipxaed")))) <= ((0) - ((d_2634_v127_).cardinality)), d_2635_v128_)
                d_2637_v129_ = nw434_
                (globalState).f1 = ((d_2541_v58_).ispropersubset(_dafny.MultiSet([d_2604_cf2_, 626, (d_2543_v60_).f41])) if d_2458_v0_ else d_2458_v0_)
                d_2638_v130_: _dafny.Array
                def lambda109_(d_2639_v129_):
                    def lambda110_(d_2640_i10_):
                        return not((d_2639_v129_).f37)

                    return lambda110_

                init59_ = lambda109_(d_2637_v129_)
                nw435_ = _dafny.Array(None, 27)
                for i0_59_ in range(nw435_.length(0)):
                    nw435_[i0_59_] = init59_(i0_59_)
                d_2638_v130_ = nw435_
                index423_ = default__.safeIndex(747, (d_2638_v130_).length(0))
                (d_2638_v130_)[index423_] = ((d_2543_v60_).f41) > (d_2543_v60_.f42)
                index424_ = default__.safeIndex(747, (d_2638_v130_).length(0))
                (d_2638_v130_)[index424_] = not((d_2637_v129_).f37)
                d_2631_v124_ = d_2631_v124_
            elif True:
                d_2641_v131_: _dafny.Seq
                d_2641_v131_ = _dafny.SeqWithoutIsStrInference([d_2458_v0_])
                d_2642_v132_: _dafny.Map
                d_2642_v132_ = _dafny.Map({_dafny.Set({(d_2641_v131_)[default__.safeIndex(d_2543_v60_.f42, len(d_2641_v131_))]}): len(d_2607_v101_)})
                rhs446_ = ((d_2543_v60_).f41) - ((0) - ((d_2543_v60_).f41))
                rhs447_ = d_2642_v132_
                rhs448_ = d_2543_v60_.f42
                lhs367_ = d_2543_v60_
                lhs368_ = globalState
                lhs367_.f42 = rhs446_
                d_2642_v132_ = rhs447_
                lhs368_.f14 = rhs448_
                (globalState).f11 = d_2458_v0_
                d_2643_v133_: C9
                nw436_ = C9()
                nw436_.ctor__((D19_DC45(d_2543_v60_.f42, d_2458_v0_)).cf83, (default__.fm0(d_2458_v0_, globalState)).set(default__.safeIndex(d_2540_v57_, len(default__.fm0(d_2458_v0_, globalState))), d_2458_v0_))
                d_2643_v133_ = nw436_
                d_2644_v134_: _dafny.MultiSet
                d_2644_v134_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2544_v61_])])
                d_2645_v135_: D24
                d_2645_v135_ = D24_DC60(d_2544_v61_)
                d_2646_v136_: _dafny.Seq
                d_2646_v136_ = _dafny.SeqWithoutIsStrInference([(d_2645_v135_).cf115])
                d_2647_v137_: _dafny.Array
                nw437_ = _dafny.Array(None, 20)
                nw437_[int(0)] = d_2540_v57_
                nw437_[int(1)] = (len((d_2643_v133_).f32)) * (len(d_2542_v59_))
                nw437_[int(2)] = (0) - ((d_2543_v60_).f41)
                nw437_[int(3)] = (d_2603_cf3_) * (len(d_2542_v59_))
                nw437_[int(4)] = d_2540_v57_
                nw437_[int(5)] = d_2543_v60_.f42
                nw437_[int(6)] = d_2604_cf2_
                nw437_[int(7)] = ((d_2543_v60_).f41) * ((d_2543_v60_).f41)
                nw437_[int(8)] = (d_2543_v60_).f41
                nw437_[int(9)] = (d_2543_v60_).f41
                nw437_[int(10)] = (d_2603_cf3_ if d_2458_v0_ else d_2604_cf2_)
                nw437_[int(11)] = (d_2543_v60_).f41
                nw437_[int(12)] = (d_2543_v60_).f41
                nw437_[int(13)] = (0) - ((d_2543_v60_).f41)
                nw437_[int(14)] = (d_2543_v60_).f41
                nw437_[int(15)] = ((d_2644_v134_)[d_2646_v136_] if (d_2646_v136_) in (d_2644_v134_) else d_2604_cf2_)
                nw437_[int(16)] = d_2603_cf3_
                nw437_[int(17)] = default__.safeModuloInt(d_2603_cf3_, 786)
                nw437_[int(18)] = d_2543_v60_.f42
                nw437_[int(19)] = (d_2543_v60_).f41
                d_2647_v137_ = nw437_
                d_2648_v138_: str
                d_2648_v138_ = _dafny.CodePoint('w')
                d_2649_v139_: _dafny.Map
                d_2649_v139_ = _dafny.Map({d_2643_v133_.f31: d_2648_v138_})
                index425_ = default__.safeIndex(488, (d_2647_v137_).length(0))
                (d_2647_v137_)[index425_] = (default__.fm2(d_2606_v100_, (self).f30, d_2649_v139_, d_2458_v0_, globalState)) + (d_2543_v60_.f42)
                index426_ = default__.safeIndex(488, (d_2647_v137_).length(0))
                rhs449_ = (self).fm12(globalState)
                rhs450_ = d_2603_cf3_
                rhs451_ = d_2543_v60_.f42
                rhs452_ = d_2541_v58_
                lhs369_ = globalState
                lhs370_ = d_2647_v137_
                lhs371_ = default__.safeIndex(488, (d_2647_v137_).length(0))
                lhs372_ = globalState
                lhs369_.f14 = rhs449_
                lhs370_[lhs371_] = rhs450_
                lhs372_.f14 = rhs451_
                d_2541_v58_ = rhs452_
                d_2650_v140_: _dafny.Map
                d_2650_v140_ = _dafny.Map({d_2643_v133_.f31: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hvnykr"))})
                d_2651_v141_: C8
                nw438_ = C8()
                nw438_.ctor__(len(((d_2650_v140_)[d_2458_v0_] if (d_2458_v0_) in (d_2650_v140_) else _dafny.SeqWithoutIsStrInference([d_2648_v138_ for d_2652_i11_ in range(default__.abs(872))]))))
                d_2651_v141_ = nw438_
            (globalState).f14 = (d_2604_cf2_ if (d_2458_v0_) and (d_2458_v0_) else (0) - ((d_2603_cf3_) + (d_2540_v57_)))
        elif True:
            d_2653___mcc_h11_ = source35_.cf0
            d_2654_cf0_ = d_2653___mcc_h11_
            d_2655_v142_: _dafny.Set
            d_2655_v142_ = _dafny.Set({848, (d_2543_v60_).f41})
            (d_2543_v60_).f42 = len((d_2655_v142_) - (d_2655_v142_))
            d_2656_v143_: C7
            nw439_ = C7()
            nw439_.ctor__((self).f28)
            d_2656_v143_ = nw439_
            d_2657_v144_: _dafny.Seq
            d_2657_v144_ = _dafny.SeqWithoutIsStrInference([d_2458_v0_])
            d_2658_v145_: _dafny.Map
            d_2658_v145_ = _dafny.Map({d_2657_v144_: (d_2543_v60_).f41})
            d_2659_v146_: C2
            nw440_ = C2()
            nw440_.ctor__(d_2658_v145_, ((self).f30).intersection((self).f30), D1_DC3(d_2458_v0_))
            d_2659_v146_ = nw440_
            (globalState).f11 = not(d_2458_v0_)
        d_2660_v147_: _dafny.Array
        nw441_ = _dafny.Array(int(0), 17)
        d_2660_v147_ = nw441_
        r0 = D2_DC8(default__.fm4(globalState), d_2660_v147_)
        r1 = default__.fm3(globalState)
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_2661_v0_: int
        d_2661_v0_ = 230
        d_2662_v1_: D12
        d_2662_v1_ = D12_DC29(False)
        d_2663_v2_: D21
        d_2663_v2_ = D21_DC52(d_2661_v0_, p0, p0, (0) - (d_2661_v0_), (d_2662_v1_).cf50)
        d_2664_i0_: int
        d_2664_i0_ = 0
        with _dafny.label("16"):
            while (d_2663_v2_).cf98:
                with _dafny.c_label("16"):
                    if (d_2664_i0_) >= (100):
                        raise _dafny.Break("16")
                    d_2664_i0_ = (d_2664_i0_) + (1)
                    d_2665_v3_: _dafny.Seq
                    d_2665_v3_ = _dafny.SeqWithoutIsStrInference([d_2661_v0_])
                    d_2666_v4_: _dafny.Seq
                    d_2666_v4_ = _dafny.SeqWithoutIsStrInference([p0])
                    d_2667_v5_: D20
                    d_2667_v5_ = D20_DC49(default__.fm5(d_2661_v0_, p0, globalState), (d_2661_v0_) * (len(d_2665_v3_)), p0, p0, ((self).f30) != (default__.fm24(d_2666_v4_, p0, globalState)))
                    d_2667_v5_ = d_2667_v5_
                    if p0:
                        (globalState).f24 = d_2661_v0_
                        d_2668_v6_: _dafny.Array
                        nw442_ = _dafny.Array(False, 12)
                        d_2668_v6_ = nw442_
                        d_2669_v7_: _dafny.Seq
                        d_2669_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "spkw"))
                        index427_ = default__.safeIndex(430, (d_2668_v6_).length(0))
                        (d_2668_v6_)[index427_] = ((d_2669_v7_).set(default__.safeIndex(683, len(d_2669_v7_)), _dafny.CodePoint('h'))) != ((d_2669_v7_).set(default__.safeIndex(d_2661_v0_, len(d_2669_v7_)), p1))
                        index428_ = default__.safeIndex(430, (d_2668_v6_).length(0))
                        (d_2668_v6_)[index428_] = (D12_DC29(True)).cf50
                        d_2670_v8_: _dafny.Map
                        d_2670_v8_ = _dafny.Map({p1: p0})
                        d_2671_v9_: C7
                        nw443_ = C7()
                        nw443_.ctor__((self).f28)
                        d_2671_v9_ = nw443_
                        d_2672_v10_: _dafny.Map
                        d_2672_v10_ = _dafny.Map({d_2671_v9_: p0})
                        d_2673_v11_: _dafny.Array
                        nw444_ = _dafny.Array(None, 15)
                        nw444_[int(0)] = len(d_2666_v4_)
                        nw444_[int(1)] = 546
                        nw444_[int(2)] = d_2661_v0_
                        nw444_[int(3)] = (d_2665_v3_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([p1 for d_2674_i1_ in range(default__.abs(968))])), len(d_2665_v3_))]
                        nw444_[int(4)] = d_2661_v0_
                        nw444_[int(5)] = d_2661_v0_
                        nw444_[int(6)] = len(d_2669_v7_)
                        nw444_[int(7)] = d_2661_v0_
                        nw444_[int(8)] = (0) - ((self).fm12(globalState))
                        nw444_[int(9)] = d_2661_v0_
                        nw444_[int(10)] = len(d_2670_v8_)
                        nw444_[int(11)] = d_2661_v0_
                        nw444_[int(12)] = len(d_2672_v10_)
                        nw444_[int(13)] = d_2661_v0_
                        nw444_[int(14)] = d_2661_v0_
                        d_2673_v11_ = nw444_
                        d_2675_v12_: _dafny.Map
                        d_2675_v12_ = _dafny.Map({380: d_2673_v11_})
                        d_2675_v12_ = d_2675_v12_
                        d_2676_v13_: _dafny.Map
                        d_2676_v13_ = _dafny.Map({(d_2668_v6_)[default__.safeIndex(430, (d_2668_v6_).length(0))]: p1})
                        (globalState).f14 = default__.fm2(d_2669_v7_, ((self).f30) | ((self).f30), d_2676_v13_, p0, globalState)
                        d_2673_v11_ = d_2673_v11_
                    elif True:
                        d_2677_v14_: _dafny.Seq
                        d_2677_v14_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lvixuq"))
                        d_2677_v14_ = ((d_2677_v14_).set(default__.safeIndex(d_2661_v0_, len(d_2677_v14_)), p1)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "auclg")))
                        d_2678_v15_: _dafny.Map
                        d_2678_v15_ = _dafny.Map({p0: (p0) and (p0)})
                        d_2678_v15_ = (d_2678_v15_).set(p0, p0)
                        d_2679_v16_: _dafny.Map
                        d_2679_v16_ = _dafny.Map({p0: p1})
                        d_2680_v17_: _dafny.Seq
                        d_2680_v17_ = _dafny.SeqWithoutIsStrInference([d_2665_v3_, _dafny.SeqWithoutIsStrInference([729, default__.fm2(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "heh")), (self).f30, d_2679_v16_, False, globalState), d_2661_v0_])])
                        rhs453_ = d_2677_v14_
                        rhs454_ = ((d_2680_v17_) + (_dafny.SeqWithoutIsStrInference([d_2665_v3_ for d_2681_i2_ in range(default__.abs(830))]))) == ((d_2680_v17_) + (d_2680_v17_))
                        lhs373_ = globalState
                        d_2677_v14_ = rhs453_
                        lhs373_.f27 = rhs454_
                        (globalState).f24 = default__.safeModuloInt(len(d_2678_v15_), (0) - (d_2661_v0_))
                        (globalState).f0 = len(d_2666_v4_)
                    d_2682_v18_: _dafny.Seq
                    d_2682_v18_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ijmnfs"))
                    d_2683_v19_: _dafny.Map
                    d_2683_v19_ = _dafny.Map({d_2661_v0_: (d_2682_v18_) <= (_dafny.SeqWithoutIsStrInference([p1 for d_2684_i3_ in range(default__.abs(824))]))})
                    d_2683_v19_ = (d_2683_v19_).set(d_2661_v0_, p0)
                    d_2682_v18_ = d_2682_v18_
                    pass
            pass
        (globalState).f24 = d_2661_v0_
        (globalState).f11 = True
        d_2685_v20_: _dafny.Seq
        d_2685_v20_ = _dafny.SeqWithoutIsStrInference([(0) - (d_2661_v0_), d_2661_v0_])
        d_2686_v21_: D19
        d_2686_v21_ = D19_DC47(D19_DC45((d_2685_v20_)[default__.safeIndex(d_2661_v0_, len(d_2685_v20_))], not(p0)))
        source37_ = d_2686_v21_
        if source37_.is_DC45:
            d_2687___mcc_h0_ = source37_.cf82
            d_2688___mcc_h1_ = source37_.cf83
            d_2689_cf83_ = d_2688___mcc_h1_
            d_2690_cf82_ = d_2687___mcc_h0_
            if ((d_2661_v0_ if d_2689_cf83_ else d_2690_cf82_)) > ((d_2690_cf82_) + ((0) - (d_2661_v0_))):
                d_2691_v22_: _dafny.Seq
                d_2691_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qscqitpvi"))
                d_2691_v22_ = d_2691_v22_
                d_2692_v23_: _dafny.Array
                def lambda111_(d_2693_p1_, d_2694_p0_):
                    def lambda112_(d_2695_i4_):
                        return _dafny.Map({_dafny.MultiSet([d_2693_p1_, d_2693_p1_]): d_2694_p0_})

                    return lambda112_

                init60_ = lambda111_(p1, p0)
                nw445_ = _dafny.Array(None, 19)
                for i0_60_ in range(nw445_.length(0)):
                    nw445_[i0_60_] = init60_(i0_60_)
                d_2692_v23_ = nw445_
                d_2696_v24_: _dafny.MultiSet
                d_2696_v24_ = _dafny.MultiSet([_dafny.CodePoint('g'), p1, p1, p1, p1])
                d_2697_v25_: _dafny.Map
                d_2697_v25_ = _dafny.Map({d_2696_v24_: p0})
                index429_ = default__.safeIndex(155, (d_2692_v23_).length(0))
                (d_2692_v23_)[index429_] = d_2697_v25_
                index430_ = default__.safeIndex(155, (d_2692_v23_).length(0))
                (d_2692_v23_)[index430_] = d_2697_v25_
                d_2698_v26_: _dafny.Array
                def lambda113_(d_2699_i5_):
                    return True

                init61_ = lambda113_
                nw446_ = _dafny.Array(None, 24)
                for i0_61_ in range(nw446_.length(0)):
                    nw446_[i0_61_] = init61_(i0_61_)
                d_2698_v26_ = nw446_
                d_2700_v27_: _dafny.Array
                def lambda114_(d_2701_cf82_):
                    def lambda115_(d_2702_i6_):
                        return (d_2702_i6_) * (d_2701_cf82_)

                    return lambda115_

                init62_ = lambda114_(d_2690_cf82_)
                nw447_ = _dafny.Array(None, 9)
                for i0_62_ in range(nw447_.length(0)):
                    nw447_[i0_62_] = init62_(i0_62_)
                d_2700_v27_ = nw447_
                default__.m0(d_2698_v26_, d_2700_v27_, default__.safeDivisionInt(-218, (0) - (d_2661_v0_)), default__.safeModuloInt(d_2661_v0_, (d_2696_v24_).cardinality), globalState)
                default__.m0(d_2698_v26_, d_2700_v27_, d_2690_cf82_, 993, globalState)
                d_2703_v28_: T0
                nw448_ = C1()
                nw448_.ctor__(d_2661_v0_, d_2690_cf82_, (self).f28)
                d_2703_v28_ = nw448_
                d_2704_v29_: _dafny.Seq
                d_2704_v29_ = _dafny.SeqWithoutIsStrInference([d_2703_v28_])
                (globalState).f1 = not((len((_dafny.SeqWithoutIsStrInference([d_2703_v28_, d_2703_v28_, d_2703_v28_, d_2703_v28_])) + ((D6_DC15(d_2704_v29_)).cf25))) > (d_2661_v0_))
            elif True:
                d_2690_cf82_ = default__.safeDivisionInt(d_2690_cf82_, d_2690_cf82_)
                d_2705_v30_: _dafny.Array
                def lambda116_(d_2706_cf83_):
                    def lambda117_(d_2707_i7_):
                        return d_2706_cf83_

                    return lambda117_

                init63_ = lambda116_(d_2689_cf83_)
                nw449_ = _dafny.Array(None, 7)
                for i0_63_ in range(nw449_.length(0)):
                    nw449_[i0_63_] = init63_(i0_63_)
                d_2705_v30_ = nw449_
                index431_ = default__.safeIndex(167, (d_2705_v30_).length(0))
                (d_2705_v30_)[index431_] = p0
                index432_ = default__.safeIndex(167, (d_2705_v30_).length(0))
                (d_2705_v30_)[index432_] = p0
                d_2708_v31_: _dafny.Map
                d_2708_v31_ = _dafny.Map({-130: p1})
                d_2709_v32_: _dafny.Map
                d_2709_v32_ = _dafny.Map({(d_2661_v0_) <= (len(d_2708_v31_)): 531})
                d_2709_v32_ = (d_2709_v32_).set(p0, default__.safeDivisionInt(d_2690_cf82_, d_2690_cf82_))
                d_2710_v33_: _dafny.Seq
                d_2710_v33_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xisi"))
                d_2711_v34_: C3
                nw450_ = C3()
                nw450_.ctor__(D1_DC3(p0))
                d_2711_v34_ = nw450_
                d_2712_v35_: _dafny.Seq
                d_2712_v35_ = _dafny.SeqWithoutIsStrInference([d_2711_v34_])
                d_2713_v36_: _dafny.MultiSet
                d_2713_v36_ = _dafny.MultiSet([d_2711_v34_])
                (globalState).f14 = default__.safeDivisionInt(len(d_2710_v33_), ((_dafny.MultiSet(d_2712_v35_)) | (d_2713_v36_)).cardinality)
                (globalState).f19 = d_2661_v0_
            (globalState).f11 = not(not(((d_2685_v20_)[default__.safeIndex(d_2661_v0_, len(d_2685_v20_))]) != ((-571) - (d_2690_cf82_))))
            (globalState).f0 = d_2661_v0_
            d_2714_v37_: _dafny.Map
            d_2714_v37_ = _dafny.Map({d_2661_v0_: d_2685_v20_})
            d_2714_v37_ = (d_2714_v37_).set(len((_dafny.Map({p0: p1})).set(p0, p1)), d_2685_v20_)
        elif source37_.is_DC46:
            d_2715___mcc_h2_ = source37_.cf84
            d_2716___mcc_h3_ = source37_.cf85
            d_2717___mcc_h4_ = source37_.cf86
            d_2718___mcc_h5_ = source37_.cf87
            d_2719_cf87_ = d_2718___mcc_h5_
            d_2720_cf86_ = d_2717___mcc_h4_
            d_2721_cf85_ = d_2716___mcc_h3_
            d_2722_cf84_ = d_2715___mcc_h2_
            d_2723_v38_: D8
            d_2723_v38_ = D8_DC21(p0, not(d_2719_cf87_), d_2661_v0_)
            d_2724_v39_: _dafny.Set
            d_2724_v39_ = _dafny.Set({D8_DC21(d_2721_cf85_, d_2720_cf86_, default__.fm2(_dafny.SeqWithoutIsStrInference([p1 for d_2725_i8_ in range(default__.abs(71))]), (self).f30, _dafny.Map({p0: p1}), d_2720_cf86_, globalState))})
            d_2726_v40_: _dafny.Set
            d_2726_v40_ = (_dafny.Set({d_2723_v38_}) if d_2721_cf85_ else d_2724_v39_)
            source38_ = d_2726_v40_
            d_2727___mcc_h8_ = source38_
            d_2728_cf37_ = d_2727___mcc_h8_
            d_2729_v41_: _dafny.Map
            d_2729_v41_ = _dafny.Map({(d_2720_cf86_) and (d_2721_cf85_): d_2661_v0_})
            d_2730_v42_: _dafny.Map
            d_2730_v42_ = _dafny.Map({d_2722_cf84_: d_2722_cf84_})
            d_2729_v41_ = (d_2729_v41_).set((d_2730_v42_) != (d_2730_v42_), -378)
            d_2731_v43_: _dafny.Array
            nw451_ = _dafny.Array(_dafny.Set({}), 27)
            d_2731_v43_ = nw451_
            index433_ = default__.safeIndex(946, (d_2731_v43_).length(0))
            (d_2731_v43_)[index433_] = _dafny.Set({p0, d_2721_cf85_})
            index434_ = default__.safeIndex(946, (d_2731_v43_).length(0))
            rhs455_ = d_2721_cf85_
            rhs456_ = ((self).f30).intersection((self).f30)
            rhs457_ = d_2722_cf84_
            rhs458_ = d_2719_cf87_
            lhs374_ = globalState
            lhs375_ = d_2731_v43_
            lhs376_ = default__.safeIndex(946, (d_2731_v43_).length(0))
            lhs377_ = globalState
            lhs374_.f1 = rhs455_
            lhs375_[lhs376_] = rhs456_
            r0 = rhs457_
            lhs377_.f11 = rhs458_
            (globalState).f19 = d_2661_v0_
            (globalState).f11 = d_2720_cf86_
            d_2732_v44_: str
            d_2732_v44_ = _dafny.CodePoint('k')
            d_2733_v45_: _dafny.Seq
            d_2733_v45_ = _dafny.SeqWithoutIsStrInference([(_dafny.SeqWithoutIsStrInference([d_2722_cf84_ for d_2734_i9_ in range(default__.abs(790))])) + (d_2685_v20_)])
            rhs459_ = p1
            rhs460_ = ((d_2733_v45_) + (d_2733_v45_)) + (_dafny.SeqWithoutIsStrInference([d_2685_v20_ for d_2735_i10_ in range(default__.abs(411))]))
            d_2732_v44_ = rhs459_
            d_2733_v45_ = rhs460_
            d_2736_v46_: _dafny.Map
            d_2736_v46_ = _dafny.Map({d_2720_cf86_: d_2722_cf84_})
            d_2737_v47_: _dafny.Map
            d_2737_v47_ = _dafny.Map({_dafny.Set({not(p0)}): d_2736_v46_})
            d_2738_v48_: _dafny.Set
            d_2738_v48_ = _dafny.Set({(((d_2737_v47_)[(self).f30] if ((self).f30) in (d_2737_v47_) else d_2736_v46_) if p0 else d_2736_v46_)})
            def iife231_():
                coll99_ = _dafny.Set()
                compr_99_: _dafny.Map
                for compr_99_ in (_dafny.Map({d_2736_v46_: d_2719_cf87_})).keys.Elements:
                    d_2739_v49_: _dafny.Map = compr_99_
                    if (d_2739_v49_) in (_dafny.Map({d_2736_v46_: d_2719_cf87_})):
                        coll99_ = coll99_.union(_dafny.Set([d_2739_v49_]))
                return _dafny.Set(coll99_)
            d_2738_v48_ = (default__.fm1(d_2722_cf84_, _dafny.Set({False, d_2721_cf85_}), globalState) if d_2720_cf86_ else (iife231_()
            ).intersection(d_2738_v48_))
            if False:
                d_2732_v44_ = p1
                d_2661_v0_ = d_2722_cf84_
                d_2740_v50_: _dafny.Array
                def lambda118_(d_2741_cf84_):
                    def lambda119_(d_2742_i11_):
                        return (d_2742_i11_) + (d_2741_cf84_)

                    return lambda119_

                init64_ = lambda118_(d_2722_cf84_)
                nw452_ = _dafny.Array(None, 28)
                for i0_64_ in range(nw452_.length(0)):
                    nw452_[i0_64_] = init64_(i0_64_)
                d_2740_v50_ = nw452_
                index435_ = default__.safeIndex(773, (d_2740_v50_).length(0))
                (d_2740_v50_)[index435_] = d_2661_v0_
                d_2743_v51_: _dafny.Map
                d_2743_v51_ = _dafny.Map({d_2661_v0_: d_2719_cf87_})
                d_2744_v52_: _dafny.Set
                d_2744_v52_ = _dafny.Set({d_2661_v0_})
                index436_ = default__.safeIndex(773, (d_2740_v50_).length(0))
                rhs461_ = d_2722_cf84_
                rhs462_ = d_2743_v51_
                rhs463_ = not(default__.fm4(globalState))
                rhs464_ = d_2722_cf84_
                rhs465_ = (d_2744_v52_).ispropersubset(d_2744_v52_)
                lhs378_ = d_2740_v50_
                lhs379_ = default__.safeIndex(773, (d_2740_v50_).length(0))
                lhs380_ = globalState
                lhs381_ = globalState
                lhs382_ = globalState
                lhs383_ = globalState
                lhs378_[lhs379_] = rhs461_
                lhs380_.f20 = rhs462_
                lhs381_.f1 = rhs463_
                lhs382_.f24 = rhs464_
                lhs383_.f1 = rhs465_
                d_2745_v53_: _dafny.Seq
                d_2745_v53_ = _dafny.SeqWithoutIsStrInference([p0, d_2719_cf87_])
                d_2746_v54_: C9
                nw453_ = C9()
                nw453_.ctor__(d_2720_cf86_, d_2745_v53_)
                d_2746_v54_ = nw453_
                (globalState).f11 = d_2721_cf85_
            elif True:
                d_2747_v55_: _dafny.Seq
                d_2747_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ve"))
                d_2747_v55_ = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shpr"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vqyi")))
                d_2748_v56_: _dafny.Array
                nw454_ = _dafny.Array(int(0), 17)
                d_2748_v56_ = nw454_
                d_2748_v56_ = d_2748_v56_
                d_2749_v57_: D10
                d_2749_v57_ = D10_DC25(d_2748_v56_, p0, _dafny.SeqWithoutIsStrInference([p0]), d_2661_v0_)
                d_2750_v58_: _dafny.MultiSet
                d_2750_v58_ = _dafny.MultiSet([(d_2749_v57_).cf42, (0) - (d_2661_v0_)])
                d_2751_v59_: _dafny.Array
                def lambda120_(d_2752_v46_):
                    def lambda121_(d_2753_i12_):
                        return d_2752_v46_

                    return lambda121_

                init65_ = lambda120_(d_2736_v46_)
                nw455_ = _dafny.Array(None, 15)
                for i0_65_ in range(nw455_.length(0)):
                    nw455_[i0_65_] = init65_(i0_65_)
                d_2751_v59_ = nw455_
                index437_ = default__.safeIndex(478, (d_2751_v59_).length(0))
                (d_2751_v59_)[index437_] = (d_2736_v46_) | (d_2736_v46_)
                d_2754_v60_: _dafny.Seq
                d_2754_v60_ = _dafny.SeqWithoutIsStrInference([d_2750_v58_, (d_2750_v58_) - (d_2750_v58_), (d_2750_v58_) | (d_2750_v58_)])
                d_2755_v61_: _dafny.MultiSet
                d_2755_v61_ = _dafny.MultiSet([d_2720_cf86_, d_2719_cf87_])
                index438_ = default__.safeIndex(478, (d_2751_v59_).length(0))
                rhs466_ = (d_2754_v60_)[default__.safeIndex(((d_2755_v61_)[d_2720_cf86_] if (d_2720_cf86_) in (d_2755_v61_) else d_2661_v0_), len(d_2754_v60_))]
                rhs467_ = default__.safeDivisionInt(d_2722_cf84_, (d_2755_v61_).cardinality)
                rhs468_ = (D22_DC56(d_2747_v55_, d_2722_cf84_, p0, (self).fm7(not(d_2721_cf85_), p0, globalState))).cf107
                rhs469_ = d_2736_v46_
                lhs384_ = globalState
                lhs385_ = globalState
                lhs386_ = d_2751_v59_
                lhs387_ = default__.safeIndex(478, (d_2751_v59_).length(0))
                d_2750_v58_ = rhs466_
                lhs384_.f19 = rhs467_
                lhs385_.f1 = rhs468_
                lhs386_[lhs387_] = rhs469_
                d_2756_v62_: _dafny.Seq
                d_2756_v62_ = _dafny.SeqWithoutIsStrInference([d_2719_cf87_, d_2721_cf85_, d_2719_cf87_])
                d_2757_v63_: _dafny.Seq
                d_2757_v63_ = _dafny.SeqWithoutIsStrInference([d_2756_v62_, (d_2749_v57_).cf41])
                d_2758_v64_: _dafny.Map
                d_2758_v64_ = _dafny.Map({d_2722_cf84_: d_2721_cf85_})
                d_2759_v65_: _dafny.Map
                d_2759_v65_ = _dafny.Map({(d_2758_v64_) | (d_2758_v64_): d_2721_cf85_})
                d_2760_v66_: _dafny.Map
                d_2760_v66_ = _dafny.Map({d_2722_cf84_: d_2661_v0_})
                rhs470_ = len(d_2747_v55_)
                rhs471_ = ((901) == ((0) - (((d_2760_v66_)[len(d_2747_v55_)] if (len(d_2747_v55_)) in (d_2760_v66_) else d_2722_cf84_)))) and (default__.fm4(globalState))
                rhs472_ = _dafny.SeqWithoutIsStrInference([(d_2756_v62_) + ((d_2756_v62_).set(default__.safeIndex(((d_2755_v61_)[d_2721_cf85_] if (d_2721_cf85_) in (d_2755_v61_) else d_2661_v0_), len(d_2756_v62_)), True)), (d_2756_v62_) + (d_2756_v62_), d_2756_v62_])
                rhs473_ = (d_2758_v64_) | ((d_2758_v64_) | ((d_2758_v64_).set(d_2722_cf84_, d_2719_cf87_)))
                rhs474_ = default__.fm51(globalState)
                lhs388_ = globalState
                lhs389_ = globalState
                lhs390_ = globalState
                lhs388_.f0 = rhs470_
                lhs389_.f11 = rhs471_
                d_2757_v63_ = rhs472_
                lhs390_.f20 = rhs473_
                d_2759_v65_ = rhs474_
                d_2761_v67_: _dafny.Map
                d_2761_v67_ = _dafny.Map({d_2720_cf86_: d_2719_cf87_})
                d_2762_v68_: C1
                nw456_ = C1()
                nw456_.ctor__(d_2722_cf84_, (d_2722_cf84_) + (len(d_2761_v67_)), D1_DC3(d_2720_cf86_))
                d_2762_v68_ = nw456_
                d_2762_v68_ = d_2762_v68_
        elif source37_.is_DC44:
            d_2763___mcc_h6_ = source37_.cf81
            d_2764_cf81_ = d_2763___mcc_h6_
            d_2765_v69_: D8
            d_2765_v69_ = D8_DC20(default__.safeDivisionInt(d_2661_v0_, d_2661_v0_), d_2661_v0_)
            source39_ = d_2765_v69_
            if source39_.is_DC20:
                d_2766___mcc_h9_ = source39_.cf31
                d_2767___mcc_h10_ = source39_.cf32
                d_2768_cf32_ = d_2767___mcc_h10_
                d_2769_cf31_ = d_2766___mcc_h9_
                d_2770_v70_: _dafny.Array
                nw457_ = _dafny.Array(_dafny.CodePoint('D'), 24)
                d_2770_v70_ = nw457_
                index439_ = default__.safeIndex(253, (d_2770_v70_).length(0))
                (d_2770_v70_)[index439_] = _dafny.CodePoint('f')
                d_2771_v71_: _dafny.Seq
                d_2771_v71_ = _dafny.SeqWithoutIsStrInference([p0])
                d_2772_v72_: _dafny.Seq
                d_2772_v72_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "m"))
                d_2773_v73_: D16
                d_2773_v73_ = D16_DC37(d_2771_v71_, d_2772_v72_, p1)
                index440_ = default__.safeIndex(253, (d_2770_v70_).length(0))
                (d_2770_v70_)[index440_] = (d_2773_v73_).cf67
                d_2774_v74_: D2
                d_2774_v74_ = D2_DC7(p0, p0, d_2772_v72_)
                d_2775_v75_: _dafny.Map
                d_2775_v75_ = _dafny.Map({d_2769_cf31_: _dafny.SeqWithoutIsStrInference([(d_2774_v74_).cf13])})
                (globalState).f14 = default__.safeModuloInt(len(d_2775_v75_), (d_2769_cf31_) - (211))
                (globalState).f11 = p0
                (globalState).f11 = p0
            elif source39_.is_DC21:
                d_2776___mcc_h11_ = source39_.cf33
                d_2777___mcc_h12_ = source39_.cf34
                d_2778___mcc_h13_ = source39_.cf35
                d_2779_cf35_ = d_2778___mcc_h13_
                d_2780_cf34_ = d_2777___mcc_h12_
                d_2781_cf33_ = d_2776___mcc_h11_
                d_2782_v76_: _dafny.Array
                nw458_ = _dafny.Array(False, 18)
                d_2782_v76_ = nw458_
                d_2782_v76_ = (d_2782_v76_ if ((self).fm12(globalState)) > (d_2661_v0_) else d_2782_v76_)
                (globalState).f0 = d_2661_v0_
                d_2783_v77_: _dafny.Seq
                d_2783_v77_ = _dafny.SeqWithoutIsStrInference([d_2780_cf34_])
                d_2784_v78_: _dafny.MultiSet
                d_2784_v78_ = _dafny.MultiSet([d_2783_v77_])
                d_2785_v79_: _dafny.Map
                d_2785_v79_ = _dafny.Map({d_2784_v78_: d_2780_cf34_})
                d_2785_v79_ = (d_2785_v79_).set(d_2784_v78_, p0)
                d_2780_cf34_ = d_2780_cf34_
            elif source39_.is_DC19:
                d_2786___mcc_h14_ = source39_.cf30
                d_2787_cf30_ = d_2786___mcc_h14_
                (globalState).f19 = d_2661_v0_
                d_2788_v80_: _dafny.Map
                d_2788_v80_ = _dafny.Map({p1: p0})
                d_2789_v81_: _dafny.Seq
                d_2789_v81_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sksnehuv"))
                d_2790_v82_: _dafny.Map
                d_2790_v82_ = _dafny.Map({((d_2788_v80_)[default__.fm5(d_2661_v0_, p0, globalState)] if (default__.fm5(d_2661_v0_, p0, globalState)) in (d_2788_v80_) else p0): d_2789_v81_})
                d_2790_v82_ = (d_2790_v82_).set(p0, d_2789_v81_)
                d_2791_v83_: _dafny.MultiSet
                d_2791_v83_ = _dafny.MultiSet([(self).fm12(globalState), d_2661_v0_])
                def iife232_():
                    coll100_ = _dafny.Map()
                    compr_100_: str
                    for compr_100_ in (d_2789_v81_).Elements:
                        d_2792_v84_: str = compr_100_
                        if (d_2792_v84_) in (d_2789_v81_):
                            coll100_[d_2792_v84_] = p0
                    return _dafny.Map(coll100_)
                d_2791_v83_ = ((d_2791_v83_).intersection(d_2791_v83_)) - (_dafny.MultiSet([len(iife232_()
                ), d_2661_v0_]))
                d_2793_v85_: _dafny.Seq
                d_2793_v85_ = _dafny.SeqWithoutIsStrInference([((d_2788_v80_)[p1] if (p1) in (d_2788_v80_) else p0), p0, p0, True])
                d_2794_v86_: _dafny.Seq
                d_2794_v86_ = _dafny.SeqWithoutIsStrInference([p0, p0, (d_2793_v85_)[default__.safeIndex(d_2661_v0_, len(d_2793_v85_))]])
                d_2795_v87_: _dafny.Map
                d_2795_v87_ = _dafny.Map({d_2793_v85_: d_2661_v0_})
                d_2796_v88_: C2
                nw459_ = C2()
                nw459_.ctor__((d_2795_v87_) | (d_2795_v87_), (self).f30, (self).f28)
                d_2796_v88_ = nw459_
                d_2796_v88_ = d_2796_v88_
            elif True:
                d_2797___mcc_h15_ = source39_.cf36
                d_2798_cf36_ = d_2797___mcc_h15_
                d_2799_v89_: _dafny.Array
                nw460_ = _dafny.Array(None, 23)
                d_2799_v89_ = nw460_
                d_2800_v90_: T0
                nw461_ = C7()
                nw461_.ctor__((self).f28)
                d_2800_v90_ = nw461_
                index441_ = default__.safeIndex(954, (d_2799_v89_).length(0))
                (d_2799_v89_)[index441_] = d_2800_v90_
                index442_ = default__.safeIndex(954, (d_2799_v89_).length(0))
                (d_2799_v89_)[index442_] = d_2800_v90_
                (globalState).f24 = default__.safeDivisionInt((self).fm12(globalState), d_2661_v0_)
                d_2801_v91_: _dafny.Seq
                d_2801_v91_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gppgxu"))
                d_2802_v92_: _dafny.Array
                def lambda122_(d_2803_p0_, d_2804_v91_, d_2805_p1_):
                    def lambda123_(d_2806_i13_):
                        return _dafny.Map({len(_dafny.Map({d_2803_p0_: d_2804_v91_})): d_2805_p1_})

                    return lambda123_

                init66_ = lambda122_(p0, d_2801_v91_, p1)
                nw462_ = _dafny.Array(None, 6)
                for i0_66_ in range(nw462_.length(0)):
                    nw462_[i0_66_] = init66_(i0_66_)
                d_2802_v92_ = nw462_
                d_2807_v93_: _dafny.Map
                d_2807_v93_ = _dafny.Map({d_2801_v91_: d_2802_v92_})
                d_2807_v93_ = (d_2807_v93_).set(d_2801_v91_, d_2802_v92_)
                (globalState).f0 = (d_2765_v69_).cf31
            if (d_2661_v0_) <= (default__.safeDivisionInt(d_2661_v0_, d_2661_v0_)):
                d_2808_v94_: _dafny.Array
                nw463_ = _dafny.Array(False, 24)
                d_2808_v94_ = nw463_
                index443_ = default__.safeIndex(853, (d_2808_v94_).length(0))
                (d_2808_v94_)[index443_] = (p0 if p0 else p0)
                index444_ = default__.safeIndex(853, (d_2808_v94_).length(0))
                (d_2808_v94_)[index444_] = p0
                (globalState).f1 = not(((d_2661_v0_) * (d_2661_v0_)) != ((0) - (d_2661_v0_)))
                d_2809_v95_: _dafny.Array
                def lambda124_(d_2810_v0_):
                    def lambda125_(d_2811_i14_):
                        return (d_2811_i14_) - (d_2810_v0_)

                    return lambda125_

                init67_ = lambda124_(d_2661_v0_)
                nw464_ = _dafny.Array(None, 28)
                for i0_67_ in range(nw464_.length(0)):
                    nw464_[i0_67_] = init67_(i0_67_)
                d_2809_v95_ = nw464_
                index445_ = default__.safeIndex(43, (d_2809_v95_).length(0))
                (d_2809_v95_)[index445_] = d_2661_v0_
                index446_ = default__.safeIndex(43, (d_2809_v95_).length(0))
                (d_2809_v95_)[index446_] = d_2661_v0_
                d_2812_v96_: _dafny.Array
                def lambda126_(d_2813_v20_):
                    def lambda127_(d_2814_i15_):
                        return d_2813_v20_

                    return lambda127_

                init68_ = lambda126_(d_2685_v20_)
                nw465_ = _dafny.Array(None, 5)
                for i0_68_ in range(nw465_.length(0)):
                    nw465_[i0_68_] = init68_(i0_68_)
                d_2812_v96_ = nw465_
                index447_ = default__.safeIndex(306, (d_2812_v96_).length(0))
                (d_2812_v96_)[index447_] = d_2685_v20_
                d_2815_v97_: _dafny.Array
                def lambda128_(d_2816_v94_):
                    def lambda129_(d_2817_i16_):
                        return _dafny.SeqWithoutIsStrInference([(d_2816_v94_)[default__.safeIndex(853, (d_2816_v94_).length(0))]])

                    return lambda129_

                init69_ = lambda128_(d_2808_v94_)
                nw466_ = _dafny.Array(None, 19)
                for i0_69_ in range(nw466_.length(0)):
                    nw466_[i0_69_] = init69_(i0_69_)
                d_2815_v97_ = nw466_
                d_2818_v98_: _dafny.Seq
                d_2818_v98_ = _dafny.SeqWithoutIsStrInference([p0])
                index448_ = default__.safeIndex(238, (d_2815_v97_).length(0))
                (d_2815_v97_)[index448_] = d_2818_v98_
                d_2819_v99_: _dafny.MultiSet
                d_2819_v99_ = _dafny.MultiSet([d_2808_v94_])
                d_2820_v100_: _dafny.Map
                d_2820_v100_ = _dafny.Map({p0: p0})
                d_2821_v101_: _dafny.Map
                d_2821_v101_ = _dafny.Map({len(d_2820_v100_): (d_2808_v94_)[default__.safeIndex(853, (d_2808_v94_).length(0))]})
                d_2822_v102_: _dafny.Seq
                d_2822_v102_ = _dafny.SeqWithoutIsStrInference([d_2821_v101_, d_2821_v101_])
                d_2823_v103_: _dafny.Seq
                d_2823_v103_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qhxilhr"))
                index449_ = default__.safeIndex(306, (d_2812_v96_).length(0))
                index450_ = default__.safeIndex(238, (d_2815_v97_).length(0))
                rhs475_ = d_2685_v20_
                rhs476_ = (d_2808_v94_)[default__.safeIndex(853, (d_2808_v94_).length(0))]
                rhs477_ = (d_2809_v95_)[default__.safeIndex(43, (d_2809_v95_).length(0))]
                rhs478_ = (((d_2818_v98_) + (default__.fm0(p0, globalState))).set(default__.safeIndex(default__.safeModuloInt(((d_2819_v99_)[d_2808_v94_] if (d_2808_v94_) in (d_2819_v99_) else d_2661_v0_), (d_2809_v95_)[default__.safeIndex(43, (d_2809_v95_).length(0))]), len((d_2818_v98_) + (default__.fm0(p0, globalState)))), True)).set(default__.safeIndex(len(d_2822_v102_), len(((d_2818_v98_) + (default__.fm0(p0, globalState))).set(default__.safeIndex(default__.safeModuloInt(((d_2819_v99_)[d_2808_v94_] if (d_2808_v94_) in (d_2819_v99_) else d_2661_v0_), (d_2809_v95_)[default__.safeIndex(43, (d_2809_v95_).length(0))]), len((d_2818_v98_) + (default__.fm0(p0, globalState)))), True))), (p1) in (d_2823_v103_))
                rhs479_ = p0
                lhs391_ = d_2812_v96_
                lhs392_ = default__.safeIndex(306, (d_2812_v96_).length(0))
                lhs393_ = globalState
                lhs394_ = globalState
                lhs395_ = d_2815_v97_
                lhs396_ = default__.safeIndex(238, (d_2815_v97_).length(0))
                lhs397_ = globalState
                lhs391_[lhs392_] = rhs475_
                lhs393_.f11 = rhs476_
                lhs394_.f19 = rhs477_
                lhs395_[lhs396_] = rhs478_
                lhs397_.f11 = rhs479_
                d_2824_v104_: _dafny.Set
                d_2824_v104_ = _dafny.Set({(0) - (d_2661_v0_), (d_2809_v95_)[default__.safeIndex(43, (d_2809_v95_).length(0))], d_2661_v0_, (d_2809_v95_)[default__.safeIndex(43, (d_2809_v95_).length(0))]})
                d_2824_v104_ = d_2824_v104_
            elif True:
                d_2825_v105_: _dafny.Array
                nw467_ = _dafny.Array(D2.default()(), 7)
                d_2825_v105_ = nw467_
                d_2826_v106_: _dafny.Array
                def lambda130_(d_2827_v0_):
                    def lambda131_(d_2828_i17_):
                        return (d_2828_i17_) * (d_2827_v0_)

                    return lambda131_

                init70_ = lambda130_(d_2661_v0_)
                nw468_ = _dafny.Array(None, 24)
                for i0_70_ in range(nw468_.length(0)):
                    nw468_[i0_70_] = init70_(i0_70_)
                d_2826_v106_ = nw468_
                d_2829_v107_: D2
                d_2829_v107_ = D2_DC8(p0, d_2826_v106_)
                index451_ = default__.safeIndex(241, (d_2825_v105_).length(0))
                (d_2825_v105_)[index451_] = d_2829_v107_
                pat_let_tv53_ = p0
                index452_ = default__.safeIndex(241, (d_2825_v105_).length(0))
                def iife233_(_pat_let66_0):
                    def iife234_(d_2830_dt__update__tmp_h0_):
                        def iife235_(_pat_let67_0):
                            def iife236_(d_2831_dt__update_hcf16_h0_):
                                return D2_DC8(d_2831_dt__update_hcf16_h0_, (d_2830_dt__update__tmp_h0_).cf17)
                            return iife236_(_pat_let67_0)
                        return iife235_(pat_let_tv53_)
                    return iife234_(_pat_let66_0)
                (d_2825_v105_)[index452_] = iife233_(d_2829_v107_)
                d_2832_v108_: _dafny.Map
                d_2832_v108_ = _dafny.Map({p0: (d_2661_v0_) - (d_2661_v0_)})
                d_2833_v109_: _dafny.Seq
                d_2833_v109_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "agdyntgc"))
                d_2832_v108_ = (d_2832_v108_).set(p0, default__.safeDivisionInt(len(d_2833_v109_), len(d_2833_v109_)))
                d_2826_v106_ = d_2826_v106_
                d_2834_v110_: _dafny.Array
                nw469_ = _dafny.Array(False, 21)
                d_2834_v110_ = nw469_
                index453_ = default__.safeIndex(752, (d_2834_v110_).length(0))
                (d_2834_v110_)[index453_] = p0
                index454_ = default__.safeIndex(752, (d_2834_v110_).length(0))
                (d_2834_v110_)[index454_] = p0
                d_2835_v111_: _dafny.Map
                d_2835_v111_ = _dafny.Map({d_2661_v0_: d_2661_v0_})
                d_2836_v112_: _dafny.Array
                nw470_ = _dafny.Array(None, 4)
                nw470_[int(0)] = p1
                nw470_[int(1)] = p1
                nw470_[int(2)] = default__.fm5(d_2661_v0_, (d_2834_v110_)[default__.safeIndex(752, (d_2834_v110_).length(0))], globalState)
                nw470_[int(3)] = (p1 if p0 else (self).fm8(d_2835_v111_, p0, p0, globalState))
                d_2836_v112_ = nw470_
                index455_ = default__.safeIndex(709, (d_2836_v112_).length(0))
                (d_2836_v112_)[index455_] = p1
                index456_ = default__.safeIndex(709, (d_2836_v112_).length(0))
                (d_2836_v112_)[index456_] = p1
            d_2837_v113_: _dafny.Array
            nw471_ = _dafny.Array(False, 17)
            d_2837_v113_ = nw471_
            d_2837_v113_ = d_2837_v113_
            index457_ = default__.safeIndex(250, (d_2837_v113_).length(0))
            (d_2837_v113_)[index457_] = (p0) and (p0)
            d_2838_v114_: D18
            d_2838_v114_ = D18_DC42(d_2685_v20_, d_2661_v0_, p0, p0)
            index458_ = default__.safeIndex(250, (d_2837_v113_).length(0))
            rhs480_ = (d_2838_v114_).cf75
            rhs481_ = p0
            lhs398_ = d_2837_v113_
            lhs399_ = default__.safeIndex(250, (d_2837_v113_).length(0))
            lhs400_ = globalState
            lhs398_[lhs399_] = rhs480_
            lhs400_.f11 = rhs481_
        elif True:
            d_2839___mcc_h7_ = source37_.cf88
            d_2840_cf88_ = d_2839___mcc_h7_
            (globalState).f14 = default__.safeDivisionInt(561, d_2661_v0_)
            d_2841_v115_: T0
            nw472_ = C3()
            nw472_.ctor__((self).f28)
            d_2841_v115_ = nw472_
            d_2842_v116_: _dafny.Seq
            d_2842_v116_ = _dafny.SeqWithoutIsStrInference([d_2841_v115_])
            d_2843_v117_: _dafny.Array
            nw473_ = _dafny.Array(None, 16)
            nw473_[int(0)] = d_2841_v115_
            nw473_[int(1)] = d_2841_v115_
            nw473_[int(2)] = d_2841_v115_
            nw473_[int(3)] = d_2841_v115_
            nw473_[int(4)] = d_2841_v115_
            nw473_[int(5)] = d_2841_v115_
            nw473_[int(6)] = d_2841_v115_
            nw473_[int(7)] = (d_2842_v116_)[default__.safeIndex(d_2661_v0_, len(d_2842_v116_))]
            nw473_[int(8)] = d_2841_v115_
            nw473_[int(9)] = d_2841_v115_
            nw473_[int(10)] = (d_2841_v115_ if p0 else d_2841_v115_)
            nw473_[int(11)] = d_2841_v115_
            nw473_[int(12)] = d_2841_v115_
            nw473_[int(13)] = (d_2841_v115_ if p0 else d_2841_v115_)
            nw473_[int(14)] = d_2841_v115_
            nw473_[int(15)] = d_2841_v115_
            d_2843_v117_ = nw473_
            index459_ = default__.safeIndex(903, (d_2843_v117_).length(0))
            (d_2843_v117_)[index459_] = d_2841_v115_
            index460_ = default__.safeIndex(903, (d_2843_v117_).length(0))
            rhs482_ = (d_2685_v20_)[default__.safeIndex(d_2661_v0_, len(d_2685_v20_))]
            rhs483_ = (d_2841_v115_ if p0 else d_2841_v115_)
            lhs401_ = globalState
            lhs402_ = d_2843_v117_
            lhs403_ = default__.safeIndex(903, (d_2843_v117_).length(0))
            lhs401_.f24 = rhs482_
            lhs402_[lhs403_] = rhs483_
            d_2844_v118_: _dafny.Seq
            d_2844_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wjtrcmxa"))
            (globalState).f0 = (d_2661_v0_) - (len(d_2844_v118_))
            d_2845_v119_: _dafny.Seq
            d_2845_v119_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nqhww")), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kpc"))])
            d_2846_v120_: _dafny.MultiSet
            d_2846_v120_ = _dafny.MultiSet([p0, p0])
            (globalState).f1 = ((d_2845_v119_)[default__.safeIndex(((d_2846_v120_)[p0] if (p0) in (d_2846_v120_) else d_2661_v0_), len(d_2845_v119_))]) not in (default__.fm52(p0, d_2661_v0_, True, len((self).f30), globalState))
        hi14_ = d_2661_v0_
        for d_2847_i18_ in range((0) - ((939) + (d_2661_v0_)), hi14_):
            d_2848_v121_: _dafny.Array
            nw474_ = _dafny.Array(_dafny.CodePoint('D'), 24)
            d_2848_v121_ = nw474_
            d_2849_v122_: _dafny.Map
            d_2849_v122_ = _dafny.Map({p0: d_2848_v121_})
            d_2849_v122_ = (d_2849_v122_).set(p0, d_2848_v121_)
            d_2850_v123_: _dafny.Seq
            d_2850_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vumynhidj"))
            d_2851_v124_: _dafny.Map
            d_2851_v124_ = _dafny.Map({(d_2850_v123_).set(default__.safeIndex(d_2847_i18_, len(d_2850_v123_)), p1): d_2847_i18_})
            d_2852_v125_: _dafny.Seq
            d_2852_v125_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p0}), (self).f30])
            d_2853_v126_: _dafny.Map
            d_2853_v126_ = _dafny.Map({p0: p1})
            d_2851_v124_ = (d_2851_v124_).set(d_2850_v123_, default__.fm2(default__.fm3(globalState), (d_2852_v125_)[default__.safeIndex(d_2661_v0_, len(d_2852_v125_))], d_2853_v126_, p0, globalState))
            d_2854_v127_: _dafny.Seq
            d_2854_v127_ = _dafny.SeqWithoutIsStrInference([default__.fm4(globalState), p0, default__.fm4(globalState), p0])
            d_2855_v128_: _dafny.Map
            d_2855_v128_ = _dafny.Map({p0: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jocjsyb")))})
            d_2856_v129_: _dafny.Map
            d_2856_v129_ = _dafny.Map({(d_2854_v127_) + (d_2854_v127_): d_2855_v128_})
            d_2857_v130_: _dafny.Map
            d_2857_v130_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p1 for d_2858_i19_ in range(default__.abs(264))]): d_2850_v123_})
            d_2856_v129_ = (d_2856_v129_).set(d_2854_v127_, _dafny.Map({p0: len(d_2857_v130_)}))
            d_2859_v131_: _dafny.Map
            d_2859_v131_ = _dafny.Map({(self).f30: d_2661_v0_})
            rhs484_ = default__.fm3(globalState)
            rhs485_ = d_2850_v123_
            rhs486_ = p0
            rhs487_ = d_2859_v131_
            lhs404_ = globalState
            d_2850_v123_ = rhs484_
            d_2850_v123_ = rhs485_
            lhs404_.f27 = rhs486_
            d_2859_v131_ = rhs487_
        d_2860_v132_: _dafny.Seq
        d_2860_v132_ = _dafny.SeqWithoutIsStrInference([p0])
        d_2861_v133_: D4
        d_2861_v133_ = D4_DC11(d_2860_v132_)
        d_2862_v134_: D4
        d_2862_v134_ = D4_DC13((D4_DC10(p1) if p0 else d_2861_v133_))
        d_2863_v135_: _dafny.Array
        def lambda132_(d_2864_v0_):
            def lambda133_(d_2865_i20_):
                return (d_2865_i20_) * (d_2864_v0_)

            return lambda133_

        init71_ = lambda132_(d_2661_v0_)
        nw475_ = _dafny.Array(None, 2)
        for i0_71_ in range(nw475_.length(0)):
            nw475_[i0_71_] = init71_(i0_71_)
        d_2863_v135_ = nw475_
        d_2866_v136_: _dafny.Map
        d_2866_v136_ = _dafny.Map({d_2863_v135_: p0})
        rhs488_ = d_2862_v134_
        rhs489_ = (self).fm12(globalState)
        rhs490_ = ((d_2866_v136_)[d_2863_v135_] if (d_2863_v135_) in (d_2866_v136_) else p0)
        rhs491_ = d_2860_v132_
        rhs492_ = d_2661_v0_
        lhs405_ = globalState
        lhs406_ = globalState
        lhs407_ = globalState
        d_2862_v134_ = rhs488_
        r0 = rhs489_
        lhs405_.f11 = rhs490_
        lhs406_.f12 = rhs491_
        lhs407_.f19 = rhs492_
        d_2867_v137_: _dafny.Seq
        d_2867_v137_ = _dafny.SeqWithoutIsStrInference([d_2860_v132_, d_2860_v132_, d_2860_v132_, d_2860_v132_])
        r0 = len((_dafny.SeqWithoutIsStrInference([d_2860_v132_ for d_2868_i21_ in range(default__.abs(550))])) + (d_2867_v137_))
        return r0

    @property
    def f30(self):
        return self._f30

class C11(T0):
    def  __init__(self):
        self._f28: D1 = D1.default()()
        pass

    def __dafnystr__(self) -> str:
        return "_module.C11"
    @property
    def f28(self):
        return self._f28
    def ctor__(self, f28):
        (self)._f28 = f28

    def fm7(self, p0, p1, globalState):
        return (D1_DC5(not(True), True, 450, False, False)).cf10

    def fm8(self, p0, p1, p2, globalState):
        if not(True):
            return _dafny.CodePoint('v')
        elif True:
            return _dafny.CodePoint('o')
        elif True:
            return _dafny.CodePoint('q')

    def fm9(self, p0, p1, p2, globalState):
        return (348) not in ((_dafny.Set({-436, len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2869_i0_ in range(default__.abs(634))]))}))}) if False else _dafny.Set({len(_dafny.Map({True: 348})), len(_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "twhdbjrx")))}))})))

    def m1(self, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_2870_v0_: _dafny.Array
        def lambda134_(d_2871_i0_):
            return D1_DC5(False, True, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhgk"))), True, True)

        init72_ = lambda134_
        nw476_ = _dafny.Array(None, 19)
        for i0_72_ in range(nw476_.length(0)):
            nw476_[i0_72_] = init72_(i0_72_)
        d_2870_v0_ = nw476_
        d_2872_v1_: bool
        d_2872_v1_ = True
        d_2873_v2_: int
        d_2873_v2_ = 211
        d_2874_v3_: D1
        d_2874_v3_ = D1_DC5(d_2872_v1_, d_2872_v1_, d_2873_v2_, d_2872_v1_, d_2872_v1_)
        index461_ = default__.safeIndex(433, (d_2870_v0_).length(0))
        (d_2870_v0_)[index461_] = d_2874_v3_
        index462_ = default__.safeIndex(433, (d_2870_v0_).length(0))
        (d_2870_v0_)[index462_] = d_2874_v3_
        d_2875_v4_: _dafny.Array
        nw477_ = _dafny.Array(None, 26)
        nw477_[int(0)] = d_2872_v1_
        nw477_[int(1)] = True
        nw477_[int(2)] = True
        nw477_[int(3)] = d_2872_v1_
        nw477_[int(4)] = d_2872_v1_
        nw477_[int(5)] = d_2872_v1_
        nw477_[int(6)] = d_2872_v1_
        nw477_[int(7)] = not(d_2872_v1_)
        nw477_[int(8)] = d_2872_v1_
        nw477_[int(9)] = d_2872_v1_
        nw477_[int(10)] = d_2872_v1_
        nw477_[int(11)] = d_2872_v1_
        nw477_[int(12)] = d_2872_v1_
        nw477_[int(13)] = d_2872_v1_
        nw477_[int(14)] = not(d_2872_v1_)
        nw477_[int(15)] = d_2872_v1_
        nw477_[int(16)] = True
        nw477_[int(17)] = d_2872_v1_
        nw477_[int(18)] = d_2872_v1_
        nw477_[int(19)] = d_2872_v1_
        nw477_[int(20)] = d_2872_v1_
        nw477_[int(21)] = d_2872_v1_
        nw477_[int(22)] = d_2872_v1_
        nw477_[int(23)] = d_2872_v1_
        nw477_[int(24)] = d_2872_v1_
        nw477_[int(25)] = d_2872_v1_
        d_2875_v4_ = nw477_
        d_2876_v5_: _dafny.Seq
        d_2876_v5_ = _dafny.SeqWithoutIsStrInference([d_2875_v4_, d_2875_v4_])
        d_2877_v6_: D2
        d_2877_v6_ = D2_DC6(d_2875_v4_)
        d_2878_v7_: _dafny.Array
        nw478_ = _dafny.Array(None, 26)
        nw478_[int(0)] = (d_2876_v5_)[default__.safeIndex(d_2873_v2_, len(d_2876_v5_))]
        nw478_[int(1)] = d_2875_v4_
        nw478_[int(2)] = d_2875_v4_
        nw478_[int(3)] = d_2875_v4_
        nw478_[int(4)] = d_2875_v4_
        nw478_[int(5)] = d_2875_v4_
        nw478_[int(6)] = d_2875_v4_
        nw478_[int(7)] = d_2875_v4_
        nw478_[int(8)] = (d_2877_v6_).cf12
        nw478_[int(9)] = d_2875_v4_
        nw478_[int(10)] = d_2875_v4_
        nw478_[int(11)] = d_2875_v4_
        nw478_[int(12)] = d_2875_v4_
        nw478_[int(13)] = d_2875_v4_
        nw478_[int(14)] = (d_2876_v5_)[default__.safeIndex(d_2873_v2_, len(d_2876_v5_))]
        nw478_[int(15)] = d_2875_v4_
        nw478_[int(16)] = d_2875_v4_
        nw478_[int(17)] = d_2875_v4_
        nw478_[int(18)] = d_2875_v4_
        nw478_[int(19)] = d_2875_v4_
        nw478_[int(20)] = d_2875_v4_
        nw478_[int(21)] = (d_2875_v4_ if False else d_2875_v4_)
        nw478_[int(22)] = d_2875_v4_
        nw478_[int(23)] = d_2875_v4_
        nw478_[int(24)] = d_2875_v4_
        nw478_[int(25)] = d_2875_v4_
        d_2878_v7_ = nw478_
        d_2878_v7_ = d_2878_v7_
        d_2879_v8_: _dafny.Map
        d_2879_v8_ = _dafny.Map({d_2872_v1_: d_2872_v1_})
        d_2880_v9_: _dafny.Array
        nw479_ = _dafny.Array(None, 25)
        nw479_[int(0)] = _dafny.Map({d_2872_v1_: default__.fm4(globalState)})
        nw479_[int(1)] = d_2879_v8_
        nw479_[int(2)] = d_2879_v8_
        nw479_[int(3)] = d_2879_v8_
        nw479_[int(4)] = _dafny.Map({d_2872_v1_: d_2872_v1_})
        nw479_[int(5)] = d_2879_v8_
        nw479_[int(6)] = d_2879_v8_
        nw479_[int(7)] = d_2879_v8_
        nw479_[int(8)] = (d_2879_v8_).set(d_2872_v1_, True)
        nw479_[int(9)] = d_2879_v8_
        nw479_[int(10)] = d_2879_v8_
        nw479_[int(11)] = _dafny.Map({d_2872_v1_: True})
        nw479_[int(12)] = d_2879_v8_
        nw479_[int(13)] = d_2879_v8_
        nw479_[int(14)] = _dafny.Map({False: d_2872_v1_})
        nw479_[int(15)] = d_2879_v8_
        nw479_[int(16)] = d_2879_v8_
        nw479_[int(17)] = d_2879_v8_
        nw479_[int(18)] = d_2879_v8_
        nw479_[int(19)] = d_2879_v8_
        nw479_[int(20)] = d_2879_v8_
        nw479_[int(21)] = d_2879_v8_
        nw479_[int(22)] = d_2879_v8_
        nw479_[int(23)] = d_2879_v8_
        nw479_[int(24)] = _dafny.Map({d_2872_v1_: d_2872_v1_})
        d_2880_v9_ = nw479_
        d_2881_v10_: _dafny.Array
        d_2881_v10_ = d_2880_v9_
        d_2882_v11_: _dafny.MultiSet
        d_2882_v11_ = _dafny.MultiSet([d_2881_v10_, d_2881_v10_])
        d_2883_v12_: _dafny.Seq
        d_2883_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "begefqf"))
        d_2884_v13_: _dafny.Set
        d_2884_v13_ = _dafny.Set({d_2872_v1_, d_2872_v1_, d_2872_v1_, d_2872_v1_})
        d_2885_v14_: _dafny.Seq
        d_2885_v14_ = _dafny.SeqWithoutIsStrInference([False, d_2872_v1_, d_2872_v1_, True])
        d_2886_v15_: str
        d_2886_v15_ = _dafny.CodePoint('b')
        d_2887_v16_: _dafny.Map
        d_2887_v16_ = _dafny.Map({(d_2885_v14_)[default__.safeIndex(d_2873_v2_, len(d_2885_v14_))]: d_2886_v15_})
        d_2888_v17_: _dafny.Seq
        d_2888_v17_ = _dafny.SeqWithoutIsStrInference([d_2873_v2_])
        d_2889_v18_: _dafny.MultiSet
        d_2889_v18_ = _dafny.MultiSet([len(d_2883_v12_), d_2873_v2_])
        d_2890_v19_: _dafny.Seq
        d_2890_v19_ = _dafny.SeqWithoutIsStrInference([d_2883_v12_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ytyvpvjxm")), d_2883_v12_, d_2883_v12_, d_2883_v12_])
        d_2891_v20_: _dafny.Array
        nw480_ = _dafny.Array(None, 7)
        nw480_[int(0)] = d_2873_v2_
        nw480_[int(1)] = (default__.fm2(d_2883_v12_, d_2884_v13_, d_2887_v16_, d_2872_v1_, globalState)) + ((d_2888_v17_)[default__.safeIndex((d_2889_v18_).cardinality, len(d_2888_v17_))])
        nw480_[int(2)] = default__.fm2(d_2883_v12_, d_2884_v13_, d_2887_v16_, not(not(d_2872_v1_)), globalState)
        nw480_[int(3)] = default__.fm2((d_2890_v19_)[default__.safeIndex(d_2873_v2_, len(d_2890_v19_))], d_2884_v13_, d_2887_v16_, True, globalState)
        nw480_[int(4)] = d_2873_v2_
        nw480_[int(5)] = 733
        nw480_[int(6)] = (0) - ((0) - ((0) - (d_2873_v2_)))
        d_2891_v20_ = nw480_
        index463_ = default__.safeIndex(664, (d_2891_v20_).length(0))
        (d_2891_v20_)[index463_] = d_2873_v2_
        d_2892_v21_: _dafny.MultiSet
        d_2892_v21_ = _dafny.MultiSet([_dafny.Map({d_2872_v1_: d_2872_v1_})])
        index464_ = default__.safeIndex(664, (d_2891_v20_).length(0))
        rhs493_ = not((d_2872_v1_) and (d_2872_v1_))
        rhs494_ = (len((d_2883_v12_) + (d_2883_v12_))) - (d_2873_v2_)
        rhs495_ = (_dafny.MultiSet([d_2881_v10_, d_2881_v10_])).set((d_2881_v10_), default__.abs(d_2873_v2_))
        rhs496_ = (((d_2892_v21_) | (default__.fm10(d_2873_v2_, d_2872_v1_, d_2872_v1_, globalState))).cardinality) * (default__.fm2(d_2883_v12_, d_2884_v13_, d_2887_v16_, d_2872_v1_, globalState))
        rhs497_ = d_2886_v15_
        lhs408_ = globalState
        lhs409_ = d_2891_v20_
        lhs410_ = default__.safeIndex(664, (d_2891_v20_).length(0))
        lhs408_.f1 = rhs493_
        d_2873_v2_ = rhs494_
        d_2882_v11_ = rhs495_
        lhs409_[lhs410_] = rhs496_
        d_2886_v15_ = rhs497_
        d_2875_v4_ = d_2875_v4_
        index465_ = default__.safeIndex(154, (d_2878_v7_).length(0))
        (d_2878_v7_)[index465_] = d_2875_v4_
        index466_ = default__.safeIndex(154, (d_2878_v7_).length(0))
        (d_2878_v7_)[index466_] = d_2875_v4_
        hi15_ = (len(d_2888_v17_)) - (895)
        for d_2893_i1_ in range((d_2891_v20_)[default__.safeIndex(664, (d_2891_v20_).length(0))], hi15_):
            d_2894_v22_: _dafny.Array
            nw481_ = _dafny.Array(_dafny.CodePoint('D'), 9)
            d_2894_v22_ = nw481_
            d_2895_v23_: _dafny.Set
            d_2895_v23_ = _dafny.Set({d_2894_v22_})
            d_2896_v24_: _dafny.Seq
            d_2896_v24_ = _dafny.SeqWithoutIsStrInference([d_2881_v10_, d_2881_v10_])
            rhs498_ = d_2895_v23_
            rhs499_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lp"))) + (d_2883_v12_)) + (d_2883_v12_)
            rhs500_ = (len(d_2896_v24_)) - ((d_2891_v20_)[default__.safeIndex(664, (d_2891_v20_).length(0))])
            rhs501_ = d_2893_i1_
            lhs411_ = globalState
            lhs412_ = globalState
            d_2895_v23_ = rhs498_
            r1 = rhs499_
            lhs411_.f24 = rhs500_
            lhs412_.f0 = rhs501_
            d_2897_v25_: _dafny.Map
            d_2897_v25_ = _dafny.Map({d_2885_v14_: (d_2891_v20_)[default__.safeIndex(664, (d_2891_v20_).length(0))]})
            d_2898_v26_: T0
            nw482_ = C2()
            nw482_.ctor__(d_2897_v25_, d_2884_v13_, (self).f28)
            d_2898_v26_ = nw482_
            d_2899_v27_: _dafny.Map
            d_2899_v27_ = _dafny.Map({not(not(d_2872_v1_)): (d_2893_i1_ if d_2872_v1_ else d_2873_v2_)})
            rhs502_ = d_2872_v1_
            rhs503_ = d_2898_v26_
            rhs504_ = d_2872_v1_
            rhs505_ = d_2899_v27_
            lhs413_ = globalState
            d_2872_v1_ = rhs502_
            d_2898_v26_ = rhs503_
            lhs413_.f1 = rhs504_
            d_2899_v27_ = rhs505_
            d_2900_v28_: C1
            nw483_ = C1()
            nw483_.ctor__((d_2891_v20_)[default__.safeIndex(664, (d_2891_v20_).length(0))], (0) - (d_2873_v2_), (d_2898_v26_).f28)
            d_2900_v28_ = nw483_
            arr9_ = (d_2878_v7_)[default__.safeIndex(154, (d_2878_v7_).length(0))]
            index467_ = default__.safeIndex(647, ((d_2878_v7_)[default__.safeIndex(154, (d_2878_v7_).length(0))]).length(0))
            arr9_[index467_] = d_2872_v1_
            arr10_ = (d_2878_v7_)[default__.safeIndex(154, (d_2878_v7_).length(0))]
            index468_ = default__.safeIndex(647, ((d_2878_v7_)[default__.safeIndex(154, (d_2878_v7_).length(0))]).length(0))
            arr10_[index468_] = d_2872_v1_
        d_2901_v29_: D2
        d_2901_v29_ = D2_DC8(d_2872_v1_, d_2891_v20_)
        pat_let_tv54_ = d_2872_v1_
        def iife237_(_pat_let68_0):
            def iife238_(d_2902_dt__update__tmp_h0_):
                def iife239_(_pat_let69_0):
                    def iife240_(d_2903_dt__update_hcf16_h0_):
                        return D2_DC8(d_2903_dt__update_hcf16_h0_, (d_2902_dt__update__tmp_h0_).cf17)
                    return iife240_(_pat_let69_0)
                return iife239_(pat_let_tv54_)
            return iife238_(_pat_let68_0)
        r0 = iife237_(d_2901_v29_)
        d_2904_v30_: _dafny.Map
        d_2904_v30_ = _dafny.Map({d_2872_v1_: d_2883_v12_})
        r1 = (d_2883_v12_) + (((d_2904_v30_)[d_2872_v1_] if (d_2872_v1_) in (d_2904_v30_) else _dafny.SeqWithoutIsStrInference([d_2886_v15_ for d_2905_i2_ in range(default__.abs(793))])))
        return r0, r1

    def m2(self, p0, p1, globalState):
        r0: int = int(0)
        d_2906_v0_: _dafny.Array
        nw484_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
        d_2906_v0_ = nw484_
        d_2907_v1_: _dafny.Seq
        d_2907_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nlija"))
        index469_ = default__.safeIndex(733, (d_2906_v0_).length(0))
        (d_2906_v0_)[index469_] = d_2907_v1_
        d_2908_v2_: _dafny.Seq
        d_2908_v2_ = _dafny.SeqWithoutIsStrInference([True])
        d_2909_v3_: int
        d_2909_v3_ = 992
        d_2910_v4_: D1
        d_2910_v4_ = D1_DC4((0) - (d_2909_v3_), p0)
        d_2911_v5_: _dafny.Map
        d_2911_v5_ = _dafny.Map({p0: (d_2910_v4_ if (d_2908_v2_)[default__.safeIndex(19, len(d_2908_v2_))] else d_2910_v4_)})
        index470_ = default__.safeIndex(733, (d_2906_v0_).length(0))
        rhs506_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ev"))
        rhs507_ = p0
        rhs508_ = d_2911_v5_
        rhs509_ = p0
        lhs414_ = d_2906_v0_
        lhs415_ = default__.safeIndex(733, (d_2906_v0_).length(0))
        lhs416_ = globalState
        lhs417_ = globalState
        lhs414_[lhs415_] = rhs506_
        lhs416_.f1 = rhs507_
        d_2911_v5_ = rhs508_
        lhs417_.f11 = rhs509_
        d_2912_v6_: _dafny.Array
        def lambda135_(d_2913_p0_):
            def lambda136_(d_2914_i0_):
                return d_2913_p0_

            return lambda136_

        init73_ = lambda135_(p0)
        nw485_ = _dafny.Array(None, 6)
        for i0_73_ in range(nw485_.length(0)):
            nw485_[i0_73_] = init73_(i0_73_)
        d_2912_v6_ = nw485_
        d_2915_v7_: _dafny.MultiSet
        d_2915_v7_ = _dafny.MultiSet([d_2912_v6_, d_2912_v6_])
        d_2915_v7_ = d_2915_v7_
        d_2916_v8_: _dafny.Set
        d_2916_v8_ = _dafny.Set({_dafny.CodePoint('t')})
        d_2917_v9_: D4
        d_2917_v9_ = D4_DC12(False, d_2916_v8_)
        d_2917_v9_ = d_2917_v9_
        d_2918_v10_: C0
        nw486_ = C0()
        nw486_.ctor__(_dafny.Set({p0}))
        d_2918_v10_ = nw486_
        d_2919_v11_: _dafny.MultiSet
        d_2919_v11_ = _dafny.MultiSet([d_2918_v10_, d_2918_v10_])
        (globalState).f1 = (d_2909_v3_) >= (default__.safeDivisionInt(d_2909_v3_, (d_2919_v11_).cardinality))
        if not(((d_2907_v1_).set(default__.safeIndex(len((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]), len(d_2907_v1_)), p1)) == (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gepy")))):
            d_2920_v12_: _dafny.Array
            nw487_ = _dafny.Array(int(0), 16)
            d_2920_v12_ = nw487_
            index471_ = default__.safeIndex(458, (d_2920_v12_).length(0))
            (d_2920_v12_)[index471_] = d_2909_v3_
            index472_ = default__.safeIndex(458, (d_2920_v12_).length(0))
            rhs510_ = p0
            rhs511_ = d_2909_v3_
            rhs512_ = p0
            lhs418_ = globalState
            lhs419_ = d_2920_v12_
            lhs420_ = default__.safeIndex(458, (d_2920_v12_).length(0))
            lhs421_ = globalState
            lhs418_.f11 = rhs510_
            lhs419_[lhs420_] = rhs511_
            lhs421_.f1 = rhs512_
            (globalState).f27 = p0
            (globalState).f11 = (self).fm7(True, p0, globalState)
            d_2921_v13_: _dafny.Array
            nw488_ = _dafny.Array(_dafny.Array(None, 0), 23)
            d_2921_v13_ = nw488_
            d_2922_v14_: _dafny.MultiSet
            d_2922_v14_ = _dafny.MultiSet([len((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]), len(d_2908_v2_)])
            d_2923_v15_: _dafny.Map
            d_2923_v15_ = _dafny.Map({d_2922_v14_: d_2908_v2_})
            d_2924_v16_: _dafny.Array
            nw489_ = _dafny.Array(None, 8)
            nw489_[int(0)] = d_2908_v2_
            nw489_[int(1)] = ((d_2923_v15_)[_dafny.MultiSet([(d_2920_v12_)[default__.safeIndex(458, (d_2920_v12_).length(0))]])] if (_dafny.MultiSet([(d_2920_v12_)[default__.safeIndex(458, (d_2920_v12_).length(0))]])) in (d_2923_v15_) else d_2908_v2_)
            nw489_[int(2)] = d_2908_v2_
            nw489_[int(3)] = d_2908_v2_
            nw489_[int(4)] = d_2908_v2_
            nw489_[int(5)] = (d_2908_v2_).set(default__.safeIndex((d_2920_v12_)[default__.safeIndex(458, (d_2920_v12_).length(0))], len(d_2908_v2_)), p0)
            nw489_[int(6)] = _dafny.SeqWithoutIsStrInference([True])
            nw489_[int(7)] = _dafny.SeqWithoutIsStrInference([default__.fm4(globalState), p0, default__.fm4(globalState)])
            d_2924_v16_ = nw489_
            index473_ = default__.safeIndex(907, (d_2921_v13_).length(0))
            (d_2921_v13_)[index473_] = d_2924_v16_
            index474_ = default__.safeIndex(538, (d_2912_v6_).length(0))
            (d_2912_v6_)[index474_] = (d_2909_v3_) < (d_2909_v3_)
            d_2925_v17_: D25
            d_2925_v17_ = D25_DC62(d_2924_v16_)
            d_2926_v18_: _dafny.Seq
            d_2926_v18_ = _dafny.SeqWithoutIsStrInference([d_2924_v16_, (d_2925_v17_).cf119, d_2924_v16_])
            index475_ = default__.safeIndex(907, (d_2921_v13_).length(0))
            index476_ = default__.safeIndex(538, (d_2912_v6_).length(0))
            rhs513_ = (d_2926_v18_)[default__.safeIndex((d_2920_v12_)[default__.safeIndex(458, (d_2920_v12_).length(0))], len(d_2926_v18_))]
            rhs514_ = d_2912_v6_
            rhs515_ = default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))) + (d_2907_v1_)), d_2909_v3_)
            rhs516_ = p0
            lhs422_ = d_2921_v13_
            lhs423_ = default__.safeIndex(907, (d_2921_v13_).length(0))
            lhs424_ = globalState
            lhs425_ = d_2912_v6_
            lhs426_ = default__.safeIndex(538, (d_2912_v6_).length(0))
            lhs422_[lhs423_] = rhs513_
            d_2912_v6_ = rhs514_
            lhs424_.f14 = rhs515_
            lhs425_[lhs426_] = rhs516_
            d_2927_v19_: D12
            d_2927_v19_ = D12_DC29(p0)
            d_2928_v20_: _dafny.Map
            d_2928_v20_ = _dafny.Map({(d_2908_v2_).set(default__.safeIndex((d_2920_v12_)[default__.safeIndex(458, (d_2920_v12_).length(0))], len(d_2908_v2_)), (d_2927_v19_).cf50): (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('l') for d_2929_i1_ in range(default__.abs(520))])) + (d_2907_v1_)})
            index477_ = default__.safeIndex(733, (d_2906_v0_).length(0))
            (d_2906_v0_)[index477_] = ((d_2928_v20_)[d_2908_v2_] if (d_2908_v2_) in (d_2928_v20_) else default__.fm3(globalState))
        elif True:
            (globalState).f1 = p0
            (globalState).f1 = True
            d_2930_v21_: str
            d_2930_v21_ = _dafny.CodePoint('a')
            d_2930_v21_ = p1
            index478_ = default__.safeIndex(561, (d_2912_v6_).length(0))
            (d_2912_v6_)[index478_] = p0
            index479_ = default__.safeIndex(561, (d_2912_v6_).length(0))
            (d_2912_v6_)[index479_] = p0
            (globalState).f1 = (d_2912_v6_)[default__.safeIndex(561, (d_2912_v6_).length(0))]
        d_2931_v22_: _dafny.MultiSet
        d_2931_v22_ = _dafny.MultiSet([True])
        d_2932_v23_: _dafny.Seq
        d_2932_v23_ = _dafny.SeqWithoutIsStrInference([len(d_2907_v1_)])
        d_2933_v24_: _dafny.MultiSet
        d_2933_v24_ = _dafny.MultiSet([len(d_2932_v23_), len((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))])])
        d_2934_v25_: _dafny.MultiSet
        d_2934_v25_ = _dafny.MultiSet([d_2933_v24_])
        d_2935_v26_: D2
        d_2935_v26_ = D2_DC6(d_2912_v6_)
        d_2936_v27_: D12
        d_2936_v27_ = D12_DC28((d_2931_v22_).isdisjoint(d_2931_v22_), d_2907_v1_, p0, ((d_2934_v25_)[d_2933_v24_] if (d_2933_v24_) in (d_2934_v25_) else d_2909_v3_), d_2935_v26_)
        source40_ = d_2936_v27_
        if source40_.is_DC28:
            d_2937___mcc_h0_ = source40_.cf45
            d_2938___mcc_h1_ = source40_.cf46
            d_2939___mcc_h2_ = source40_.cf47
            d_2940___mcc_h3_ = source40_.cf48
            d_2941___mcc_h4_ = source40_.cf49
            d_2942_cf49_ = d_2941___mcc_h4_
            d_2943_cf48_ = d_2940___mcc_h3_
            d_2944_cf47_ = d_2939___mcc_h2_
            d_2945_cf46_ = d_2938___mcc_h1_
            d_2946_cf45_ = d_2937___mcc_h0_
            d_2947_v28_: D1
            d_2947_v28_ = D1_DC3((_dafny.MultiSet(d_2908_v2_)) == (_dafny.MultiSet(d_2908_v2_)))
            d_2948_v29_: _dafny.MultiSet
            d_2948_v29_ = _dafny.MultiSet([p1, p1])
            d_2949_v30_: _dafny.Map
            d_2949_v30_ = _dafny.Map({d_2946_cf45_: d_2946_cf45_})
            rhs517_ = d_2947_v28_
            rhs518_ = (_dafny.MultiSet([p1, p1])).isdisjoint(d_2948_v29_)
            rhs519_ = not((d_2943_cf48_) == (len(d_2949_v30_)))
            lhs427_ = globalState
            d_2947_v28_ = rhs517_
            lhs427_.f27 = rhs518_
            d_2944_cf47_ = rhs519_
            d_2950_v31_: C3
            nw490_ = C3()
            nw490_.ctor__((self).f28)
            d_2950_v31_ = nw490_
            d_2951_v32_: _dafny.Array
            def lambda137_(d_2952_p1_):
                def lambda138_(d_2953_i2_):
                    return d_2952_p1_

                return lambda138_

            init74_ = lambda137_(p1)
            nw491_ = _dafny.Array(None, 8)
            for i0_74_ in range(nw491_.length(0)):
                nw491_[i0_74_] = init74_(i0_74_)
            d_2951_v32_ = nw491_
            index480_ = default__.safeIndex(799, (d_2951_v32_).length(0))
            (d_2951_v32_)[index480_] = p1
            index481_ = default__.safeIndex(799, (d_2951_v32_).length(0))
            (d_2951_v32_)[index481_] = ((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))])[default__.safeIndex((d_2943_cf48_) - (len(d_2932_v23_)), len((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]))]
            d_2954_v33_: _dafny.Array
            def lambda139_(d_2955_v30_):
                def lambda140_(d_2956_i3_):
                    return D8_DC19(d_2955_v30_)

                return lambda140_

            init75_ = lambda139_(d_2949_v30_)
            nw492_ = _dafny.Array(None, 11)
            for i0_75_ in range(nw492_.length(0)):
                nw492_[i0_75_] = init75_(i0_75_)
            d_2954_v33_ = nw492_
            d_2957_v34_: _dafny.Seq
            d_2957_v34_ = _dafny.SeqWithoutIsStrInference([d_2949_v30_, d_2949_v30_, d_2949_v30_, d_2949_v30_])
            d_2958_v35_: D8
            d_2958_v35_ = D8_DC19((d_2957_v34_)[default__.safeIndex(d_2909_v3_, len(d_2957_v34_))])
            index482_ = default__.safeIndex(474, (d_2954_v33_).length(0))
            (d_2954_v33_)[index482_] = d_2958_v35_
            d_2959_v36_: _dafny.Set
            d_2959_v36_ = _dafny.Set({(0) - (d_2909_v3_), d_2943_cf48_})
            index483_ = default__.safeIndex(255, (d_2912_v6_).length(0))
            (d_2912_v6_)[index483_] = (default__.fm34(globalState)).ispropersubset(d_2959_v36_)
            d_2960_v37_: _dafny.Array
            def lambda141_(d_2961_cf48_):
                def lambda142_(d_2962_i4_):
                    return (d_2962_i4_) + (d_2961_cf48_)

                return lambda142_

            init76_ = lambda141_(d_2943_cf48_)
            nw493_ = _dafny.Array(None, 23)
            for i0_76_ in range(nw493_.length(0)):
                nw493_[i0_76_] = init76_(i0_76_)
            d_2960_v37_ = nw493_
            index484_ = default__.safeIndex(785, (d_2960_v37_).length(0))
            (d_2960_v37_)[index484_] = d_2909_v3_
            d_2963_v38_: _dafny.Array
            nw494_ = _dafny.Array(D12.default()(), 19)
            d_2963_v38_ = nw494_
            d_2964_v39_: _dafny.Map
            d_2964_v39_ = _dafny.Map({d_2963_v38_: d_2945_cf46_})
            index485_ = default__.safeIndex(474, (d_2954_v33_).length(0))
            index486_ = default__.safeIndex(255, (d_2912_v6_).length(0))
            index487_ = default__.safeIndex(785, (d_2960_v37_).length(0))
            rhs520_ = d_2958_v35_
            rhs521_ = d_2946_cf45_
            rhs522_ = len((d_2964_v39_) | (((_dafny.Map({d_2963_v38_: d_2945_cf46_})).set(d_2963_v38_, d_2945_cf46_) if p0 else d_2964_v39_)))
            rhs523_ = d_2943_cf48_
            rhs524_ = d_2909_v3_
            lhs428_ = d_2954_v33_
            lhs429_ = default__.safeIndex(474, (d_2954_v33_).length(0))
            lhs430_ = d_2912_v6_
            lhs431_ = default__.safeIndex(255, (d_2912_v6_).length(0))
            lhs432_ = d_2960_v37_
            lhs433_ = default__.safeIndex(785, (d_2960_v37_).length(0))
            lhs434_ = globalState
            lhs428_[lhs429_] = rhs520_
            lhs430_[lhs431_] = rhs521_
            lhs432_[lhs433_] = rhs522_
            lhs434_.f19 = rhs523_
            r0 = rhs524_
        elif source40_.is_DC29:
            d_2965___mcc_h5_ = source40_.cf50
            d_2966_cf50_ = d_2965___mcc_h5_
            d_2967_v40_: C4
            nw495_ = C4()
            nw495_.ctor__(p0, d_2906_v0_)
            d_2967_v40_ = nw495_
            d_2968_v41_: _dafny.Map
            d_2968_v41_ = _dafny.Map({p1: d_2909_v3_})
            d_2969_v42_: _dafny.Array
            nw496_ = _dafny.Array(None, 13)
            nw496_[int(0)] = d_2909_v3_
            nw496_[int(1)] = d_2909_v3_
            nw496_[int(2)] = 324
            nw496_[int(3)] = d_2909_v3_
            nw496_[int(4)] = d_2909_v3_
            nw496_[int(5)] = d_2909_v3_
            nw496_[int(6)] = d_2909_v3_
            nw496_[int(7)] = len(d_2907_v1_)
            nw496_[int(8)] = d_2909_v3_
            nw496_[int(9)] = (D25_DC63(d_2909_v3_, len(d_2907_v1_), d_2967_v40_)).cf121
            nw496_[int(10)] = d_2909_v3_
            nw496_[int(11)] = d_2909_v3_
            nw496_[int(12)] = ((d_2968_v41_)[p1] if (p1) in (d_2968_v41_) else d_2909_v3_)
            d_2969_v42_ = nw496_
            d_2970_v43_: _dafny.Set
            d_2970_v43_ = _dafny.Set({d_2969_v42_})
            d_2971_v44_: _dafny.Map
            d_2971_v44_ = _dafny.Map({_dafny.Set({d_2909_v3_}): d_2970_v43_})
            d_2972_v45_: _dafny.Set
            d_2972_v45_ = _dafny.Set({d_2909_v3_})
            d_2971_v44_ = (d_2971_v44_).set(d_2972_v45_, d_2970_v43_)
            index488_ = default__.safeIndex(733, (d_2906_v0_).length(0))
            (d_2906_v0_)[index488_] = (d_2907_v1_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_2973_i5_ in range(default__.abs(-797))]))
            index489_ = default__.safeIndex(733, (d_2906_v0_).length(0))
            (d_2906_v0_)[index489_] = (d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]
            rhs525_ = not (not ((d_2967_v40_).f37) or ((d_2967_v40_).f37)) or (not((d_2933_v24_).ispropersubset(d_2933_v24_)))
            rhs526_ = (d_2967_v40_).f37
            lhs435_ = globalState
            lhs436_ = globalState
            lhs435_.f11 = rhs525_
            lhs436_.f1 = rhs526_
        elif source40_.is_DC30:
            d_2974___mcc_h6_ = source40_.cf51
            d_2975___mcc_h7_ = source40_.cf52
            d_2976___mcc_h8_ = source40_.cf53
            d_2977___mcc_h9_ = source40_.cf54
            d_2978___mcc_h10_ = source40_.cf55
            d_2979_cf55_ = d_2978___mcc_h10_
            d_2980_cf54_ = d_2977___mcc_h9_
            d_2981_cf53_ = d_2976___mcc_h8_
            d_2982_cf52_ = d_2975___mcc_h7_
            d_2983_cf51_ = d_2974___mcc_h6_
            d_2984_v46_: C5
            nw497_ = C5()
            nw497_.ctor__((d_2980_cf54_) + (d_2909_v3_), d_2983_cf51_, (self).f28)
            d_2984_v46_ = nw497_
            d_2985_v47_: C4
            nw498_ = C4()
            nw498_.ctor__(d_2983_cf51_, d_2906_v0_)
            d_2985_v47_ = nw498_
            d_2986_v48_: D25
            d_2986_v48_ = D25_DC63(d_2980_cf54_, len((d_2908_v2_).set(default__.safeIndex(d_2909_v3_, len(d_2908_v2_)), p0)), d_2985_v47_)
            rhs527_ = _dafny.SeqWithoutIsStrInference([(p1 if (d_2985_v47_).f37 else p1) for d_2987_i6_ in range(default__.abs(-845))])
            rhs528_ = d_2980_cf54_
            rhs529_ = d_2984_v46_
            rhs530_ = D25_DC63((d_2984_v46_).f35, d_2909_v3_, d_2985_v47_)
            d_2907_v1_ = rhs527_
            d_2909_v3_ = rhs528_
            d_2984_v46_ = rhs529_
            d_2986_v48_ = rhs530_
            d_2988_v49_: _dafny.Map
            d_2988_v49_ = _dafny.Map({True: (d_2907_v1_)[default__.safeIndex(d_2980_cf54_, len(d_2907_v1_))]})
            (globalState).f24 = default__.fm2((d_2907_v1_).set(default__.safeIndex(635, len(d_2907_v1_)), p1), (d_2918_v10_).f43, d_2988_v49_, not(d_2982_cf52_), globalState)
            if (len(d_2907_v1_)) <= (((d_2931_v22_)[p0] if (p0) in (d_2931_v22_) else len(d_2907_v1_))):
                d_2989_v50_: C10
                nw499_ = C10()
                nw499_.ctor__((d_2918_v10_).f43, (self).f28)
                d_2989_v50_ = nw499_
                d_2990_v51_: _dafny.MultiSet
                d_2990_v51_ = _dafny.MultiSet([d_2910_v4_, d_2910_v4_])
                d_2991_v52_: _dafny.Set
                d_2991_v52_ = _dafny.Set({d_2990_v51_})
                index490_ = default__.safeIndex(355, (d_2912_v6_).length(0))
                (d_2912_v6_)[index490_] = (d_2991_v52_).issubset(d_2991_v52_)
                d_2992_v53_: _dafny.Seq
                d_2992_v53_ = _dafny.SeqWithoutIsStrInference([d_2907_v1_, d_2907_v1_, _dafny.SeqWithoutIsStrInference([p1 for d_2993_i9_ in range(default__.abs(576))]), d_2907_v1_, (d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]])
                index491_ = default__.safeIndex(355, (d_2912_v6_).length(0))
                (d_2912_v6_)[index491_] = (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([p1 for d_2994_i8_ in range(default__.abs(760))]) for d_2995_i7_ in range(default__.abs(980))])) <= (d_2992_v53_)
                d_2996_v54_: C8
                nw500_ = C8()
                nw500_.ctor__((0) - ((d_2980_cf54_) - ((d_2984_v46_).f35)))
                d_2996_v54_ = nw500_
                d_2996_v54_ = d_2996_v54_
                (globalState).f1 = (d_2984_v46_).f36
                index492_ = default__.safeIndex(355, (d_2912_v6_).length(0))
                rhs531_ = (466) >= (len(d_2907_v1_))
                rhs532_ = (d_2984_v46_).f35
                rhs533_ = ((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]) != (((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]) + (d_2907_v1_))
                lhs437_ = d_2912_v6_
                lhs438_ = default__.safeIndex(355, (d_2912_v6_).length(0))
                lhs439_ = globalState
                lhs437_[lhs438_] = rhs531_
                d_2980_cf54_ = rhs532_
                lhs439_.f1 = rhs533_
            elif True:
                d_2997_v55_: C10
                nw501_ = C10()
                nw501_.ctor__((d_2918_v10_).f43, D1_DC3(d_2981_cf53_))
                d_2997_v55_ = nw501_
                (globalState).f0 = (d_2909_v3_ if d_2979_cf55_ else (d_2980_cf54_ if (d_2985_v47_).f37 else 686))
                d_2998_v56_: _dafny.Array
                def lambda143_(d_2999_v3_):
                    def lambda144_(d_3000_i10_):
                        return (d_3000_i10_) - ((0) - (d_2999_v3_))

                    return lambda144_

                init77_ = lambda143_(d_2909_v3_)
                nw502_ = _dafny.Array(None, 7)
                for i0_77_ in range(nw502_.length(0)):
                    nw502_[i0_77_] = init77_(i0_77_)
                d_2998_v56_ = nw502_
                d_3001_v57_: _dafny.Seq
                d_3001_v57_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(d_2984_v46_).f35, len(d_2908_v2_), 26, d_2909_v3_])])
                rhs534_ = ((d_2918_v10_).f43).ispropersubset((d_2997_v55_).f30)
                rhs535_ = (0) - (len((_dafny.SeqWithoutIsStrInference([d_2932_v23_ for d_3002_i11_ in range(default__.abs(-652))])) + ((d_3001_v57_) + (d_3001_v57_))))
                rhs536_ = d_2998_v56_
                rhs537_ = (d_2984_v46_).f36
                rhs538_ = d_2932_v23_
                lhs440_ = globalState
                lhs441_ = globalState
                lhs442_ = globalState
                lhs440_.f27 = rhs534_
                lhs441_.f19 = rhs535_
                d_2998_v56_ = rhs536_
                lhs442_.f27 = rhs537_
                d_2932_v23_ = rhs538_
                (globalState).f0 = d_2980_cf54_
                (globalState).f1 = (d_2933_v24_).isdisjoint(d_2933_v24_)
            d_3003_v58_: D19
            d_3003_v58_ = D19_DC45(d_2909_v3_, not((d_2984_v46_).f36))
            d_3004_v59_: D19
            d_3004_v59_ = D19_DC47(d_3003_v58_)
            d_3005_v60_: D19
            d_3005_v60_ = D19_DC47(d_3004_v59_)
            source41_ = d_3005_v60_
            if source41_.is_DC45:
                d_3006___mcc_h12_ = source41_.cf82
                d_3007___mcc_h13_ = source41_.cf83
                d_3008_cf83_ = d_3007___mcc_h13_
                d_3009_cf82_ = d_3006___mcc_h12_
                d_3010_v61_: _dafny.Array
                nw503_ = _dafny.Array(None, 1)
                nw503_[int(0)] = d_2980_cf54_
                d_3010_v61_ = nw503_
                index493_ = default__.safeIndex(913, (d_3010_v61_).length(0))
                (d_3010_v61_)[index493_] = default__.safeDivisionInt((d_2984_v46_).f35, (d_2984_v46_).f35)
                index494_ = default__.safeIndex(913, (d_3010_v61_).length(0))
                (d_3010_v61_)[index494_] = ((d_2984_v46_).f35) - (d_2980_cf54_)
                d_3011_v62_: _dafny.Map
                d_3011_v62_ = _dafny.Map({p0: (d_2984_v46_).fm7((d_2984_v46_).fm7(p0, (d_2984_v46_).f36, globalState), False, globalState)})
                d_3012_v63_: _dafny.Map
                d_3012_v63_ = _dafny.Map({d_3011_v62_: (d_2984_v46_).f35})
                (globalState).f27 = ((_dafny.Map({d_2979_cf55_: (d_2984_v46_).f36})) | (default__.fm48(p1, (d_3010_v61_)[default__.safeIndex(913, (d_3010_v61_).length(0))], globalState))) not in ((d_3012_v63_) | (d_3012_v63_))
                d_3013_v64_: _dafny.Array
                nw504_ = _dafny.Array(_dafny.Array(None, 0), 18)
                d_3013_v64_ = nw504_
                d_3014_v65_: D2
                d_3015_v66_: _dafny.Set
                out20_: D2
                out21_: _dafny.Set
                out20_, out21_ = (self).m3(d_3013_v64_, d_2908_v2_, d_2907_v1_, globalState)
                d_3014_v65_ = out20_
                d_3015_v66_ = out21_
                d_3016_v67_: _dafny.Array
                def lambda145_(d_3017_i12_):
                    return _dafny.CodePoint('d')

                init78_ = lambda145_
                nw505_ = _dafny.Array(None, 7)
                for i0_78_ in range(nw505_.length(0)):
                    nw505_[i0_78_] = init78_(i0_78_)
                d_3016_v67_ = nw505_
                index495_ = default__.safeIndex(598, (d_3016_v67_).length(0))
                (d_3016_v67_)[index495_] = p1
                index496_ = default__.safeIndex(598, (d_3016_v67_).length(0))
                (d_3016_v67_)[index496_] = p1
            elif source41_.is_DC46:
                d_3018___mcc_h14_ = source41_.cf84
                d_3019___mcc_h15_ = source41_.cf85
                d_3020___mcc_h16_ = source41_.cf86
                d_3021___mcc_h17_ = source41_.cf87
                d_3022_cf87_ = d_3021___mcc_h17_
                d_3023_cf86_ = d_3020___mcc_h16_
                d_3024_cf85_ = d_3019___mcc_h15_
                d_3025_cf84_ = d_3018___mcc_h14_
                index497_ = default__.safeIndex(733, (d_2906_v0_).length(0))
                (d_2906_v0_)[index497_] = ((d_2907_v1_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uetbm")))) + (d_2907_v1_)
                d_3026_v68_: C1
                nw506_ = C1()
                nw506_.ctor__((-965) + ((d_2984_v46_).f35), default__.fm2((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))], (d_2918_v10_).f43, d_2988_v49_, False, globalState), (self).f28)
                d_3026_v68_ = nw506_
                d_3027_v69_: _dafny.Array
                nw507_ = _dafny.Array(_dafny.Set({}), 26)
                d_3027_v69_ = nw507_
                index498_ = default__.safeIndex(725, (d_3027_v69_).length(0))
                (d_3027_v69_)[index498_] = d_2916_v8_
                index499_ = default__.safeIndex(725, (d_3027_v69_).length(0))
                (d_3027_v69_)[index499_] = (_dafny.Set({((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))])[default__.safeIndex(197, len((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]))], p1}) if d_2983_cf51_ else d_2916_v8_)
                d_3028_v70_: _dafny.Map
                d_3028_v70_ = _dafny.Map({d_2981_cf53_: (d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]})
                d_2979_cf55_ = (d_3028_v70_) != (d_3028_v70_)
            elif source41_.is_DC44:
                d_3029___mcc_h18_ = source41_.cf81
                d_3030_cf81_ = d_3029___mcc_h18_
                d_3031_v71_: _dafny.Map
                d_3031_v71_ = _dafny.Map({(d_2985_v47_).f37: d_2980_cf54_})
                d_3032_v72_: D24
                d_3032_v72_ = D24_DC60(d_3031_v71_)
                index500_ = default__.safeIndex(733, (d_2906_v0_).length(0))
                rhs539_ = (d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]
                rhs540_ = d_2909_v3_
                rhs541_ = d_3032_v72_
                lhs443_ = d_2906_v0_
                lhs444_ = default__.safeIndex(733, (d_2906_v0_).length(0))
                lhs445_ = globalState
                lhs443_[lhs444_] = rhs539_
                lhs445_.f19 = rhs540_
                d_3032_v72_ = rhs541_
                d_3033_v73_: _dafny.Map
                d_3033_v73_ = _dafny.Map({not(d_2983_cf51_): d_2988_v49_})
                d_3034_v74_: D15
                d_3034_v74_ = D15_DC34((d_2984_v46_).f36, d_2909_v3_)
                d_3035_v75_: _dafny.Map
                d_3035_v75_ = _dafny.Map({default__.fm2(d_2907_v1_, (d_2918_v10_).f43, d_2988_v49_, d_2982_cf52_, globalState): default__.fm2((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))], (d_2918_v10_).f43, ((d_3033_v73_)[d_2982_cf52_] if (d_2982_cf52_) in (d_3033_v73_) else _dafny.Map({(d_3034_v74_).cf59: p1})), (d_2984_v46_).f36, globalState)})
                d_3035_v75_ = (d_3035_v75_).set(d_2980_cf54_, len(d_2907_v1_))
                d_3036_v76_: _dafny.Array
                def lambda146_(d_3037_cf54_):
                    def lambda147_(d_3038_i13_):
                        return (d_3038_i13_) * (d_3037_cf54_)

                    return lambda147_

                init79_ = lambda146_(d_2980_cf54_)
                nw508_ = _dafny.Array(None, 20)
                for i0_79_ in range(nw508_.length(0)):
                    nw508_[i0_79_] = init79_(i0_79_)
                d_3036_v76_ = nw508_
                index501_ = default__.safeIndex(90, (d_3036_v76_).length(0))
                (d_3036_v76_)[index501_] = d_2909_v3_
                d_3039_v77_: _dafny.Set
                d_3039_v77_ = _dafny.Set({161, (d_2984_v46_).f35, (d_2984_v46_).f35})
                d_3040_v78_: _dafny.Array
                nw509_ = _dafny.Array(_dafny.Seq({}), 5)
                d_3040_v78_ = nw509_
                index502_ = default__.safeIndex(56, (d_3040_v78_).length(0))
                (d_3040_v78_)[index502_] = (d_2932_v23_) + (d_2932_v23_)
                d_3041_v79_: _dafny.Map
                d_3041_v79_ = _dafny.Map({default__.fm4(globalState): d_2932_v23_})
                index503_ = default__.safeIndex(90, (d_3036_v76_).length(0))
                index504_ = default__.safeIndex(56, (d_3040_v78_).length(0))
                rhs542_ = (0) - ((0) - (default__.safeDivisionInt(((d_2984_v46_).f35) * (d_2909_v3_), (d_2984_v46_).f35)))
                rhs543_ = d_3039_v77_
                rhs544_ = (((d_3041_v79_)[d_2982_cf52_] if (d_2982_cf52_) in (d_3041_v79_) else _dafny.SeqWithoutIsStrInference([default__.fm2(d_2907_v1_, (d_2918_v10_).f43, d_2988_v49_, d_2983_cf51_, globalState), d_2980_cf54_]))) + (d_2932_v23_)
                lhs446_ = d_3036_v76_
                lhs447_ = default__.safeIndex(90, (d_3036_v76_).length(0))
                lhs448_ = d_3040_v78_
                lhs449_ = default__.safeIndex(56, (d_3040_v78_).length(0))
                lhs446_[lhs447_] = rhs542_
                d_3039_v77_ = rhs543_
                lhs448_[lhs449_] = rhs544_
                d_3036_v76_ = (d_3036_v76_ if (d_2984_v46_).f36 else d_3036_v76_)
            elif True:
                d_3042___mcc_h19_ = source41_.cf88
                d_3043_cf88_ = d_3042___mcc_h19_
                d_3044_v80_: _dafny.Array
                nw510_ = _dafny.Array(int(0), 23)
                d_3044_v80_ = nw510_
                index505_ = default__.safeIndex(40, (d_3044_v80_).length(0))
                (d_3044_v80_)[index505_] = d_2909_v3_
                index506_ = default__.safeIndex(40, (d_3044_v80_).length(0))
                (d_3044_v80_)[index506_] = (0) - ((0) - ((d_2932_v23_)[default__.safeIndex((d_2984_v46_).f35, len(d_2932_v23_))]))
                d_3045_v81_: C5
                nw511_ = C5()
                nw511_.ctor__(d_2909_v3_, ((d_3044_v80_)[default__.safeIndex(40, (d_3044_v80_).length(0))]) > ((d_2984_v46_).f35), (self).f28)
                d_3045_v81_ = nw511_
                d_3046_v82_: D2
                d_3046_v82_ = D2_DC7((d_2985_v47_).f37, d_2982_cf52_, d_2907_v1_)
                d_3047_v83_: _dafny.Map
                d_3047_v83_ = _dafny.Map({((d_2918_v10_).f43).ispropersubset(_dafny.Set({(d_2984_v46_).f36})): (d_3046_v82_).cf13})
                d_3047_v83_ = (d_3047_v83_).set((d_3045_v81_).f36, d_2979_cf55_)
                index507_ = default__.safeIndex(40, (d_3044_v80_).length(0))
                rhs545_ = (d_2981_cf53_ if (True) and (p0) else (self).fm9((d_3045_v81_).f35, (d_2985_v47_).f37, not(False), globalState))
                rhs546_ = (d_2984_v46_).f35
                lhs450_ = globalState
                lhs451_ = d_3044_v80_
                lhs452_ = default__.safeIndex(40, (d_3044_v80_).length(0))
                lhs450_.f1 = rhs545_
                lhs451_[lhs452_] = rhs546_
        elif True:
            d_3048___mcc_h11_ = source40_.cf44
            d_3049_cf44_ = d_3048___mcc_h11_
            d_3050_v84_: _dafny.Array
            def lambda148_(d_3051_v3_):
                def lambda149_(d_3052_i14_):
                    return (d_3052_i14_) - (d_3051_v3_)

                return lambda149_

            init80_ = lambda148_(d_2909_v3_)
            nw512_ = _dafny.Array(None, 15)
            for i0_80_ in range(nw512_.length(0)):
                nw512_[i0_80_] = init80_(i0_80_)
            d_3050_v84_ = nw512_
            index508_ = default__.safeIndex(419, (d_3050_v84_).length(0))
            (d_3050_v84_)[index508_] = (696) - (d_2909_v3_)
            d_3053_v85_: _dafny.Map
            d_3053_v85_ = _dafny.Map({p0: _dafny.CodePoint('b')})
            d_3054_v86_: C8
            nw513_ = C8()
            nw513_.ctor__(len(_dafny.Set({d_2912_v6_})))
            d_3054_v86_ = nw513_
            d_3055_v87_: _dafny.Map
            d_3055_v87_ = _dafny.Map({d_3054_v86_: (_dafny.MultiSet(default__.fm0(not(p0), globalState))).isdisjoint(d_2931_v22_)})
            index509_ = default__.safeIndex(419, (d_3050_v84_).length(0))
            rhs547_ = default__.fm2((default__.fm3(globalState)) + (d_2907_v1_), ((d_2918_v10_).f43).intersection((d_2918_v10_).f43), d_3053_v85_, False, globalState)
            rhs548_ = d_2909_v3_
            rhs549_ = ((d_3055_v87_)[d_3054_v86_] if (d_3054_v86_) in (d_3055_v87_) else p0)
            rhs550_ = (0) - (d_2909_v3_)
            lhs453_ = d_3050_v84_
            lhs454_ = default__.safeIndex(419, (d_3050_v84_).length(0))
            lhs455_ = globalState
            lhs456_ = globalState
            lhs453_[lhs454_] = rhs547_
            d_2909_v3_ = rhs548_
            lhs455_.f11 = rhs549_
            lhs456_.f19 = rhs550_
            (d_3054_v86_).f33 = (0) - (d_2909_v3_)
            d_3056_v89_: _dafny.Map
            d_3056_v89_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([p0, p0]): p0})
            d_3057_v90_: C2
            nw514_ = C2()
            def iife241_():
                coll101_ = _dafny.Map()
                compr_101_: _dafny.Seq
                for compr_101_ in (d_3056_v89_).keys.Elements:
                    d_3058_v88_: _dafny.Seq = compr_101_
                    if (d_3058_v88_) in (d_3056_v89_):
                        coll101_[d_3058_v88_] = d_3054_v86_.f33
                return _dafny.Map(coll101_)
            nw514_.ctor__(iife241_()
            , ((d_2918_v10_).f43).intersection((d_2918_v10_).f43), (self).f28)
            d_3057_v90_ = nw514_
            d_3059_v91_: _dafny.Map
            d_3059_v91_ = _dafny.Map({p0: (d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))]})
            (globalState).f27 = (((d_3059_v91_)[p0] if (p0) in (d_3059_v91_) else (d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))])) != ((((d_2936_v27_).cf46) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))).set(default__.safeIndex((d_3050_v84_)[default__.safeIndex(419, (d_3050_v84_).length(0))], len(((d_2936_v27_).cf46) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))), p1))
        d_3060_v92_: _dafny.Map
        d_3060_v92_ = _dafny.Map({p0: p1})
        r0 = (default__.fm2((d_2906_v0_)[default__.safeIndex(733, (d_2906_v0_).length(0))], (d_2918_v10_).f43, d_3060_v92_, p0, globalState)) - (d_2909_v3_)
        return r0

    def m3(self, p0, p1, p2, globalState):
        r0: D2 = D2.default()()
        r1: _dafny.Set = _dafny.Set({})
        d_3061_v0_: _dafny.Array
        def lambda150_(d_3062_p2_):
            def lambda151_(d_3063_i0_):
                return (d_3063_i0_) - (len(d_3062_p2_))

            return lambda151_

        init81_ = lambda150_(p2)
        nw515_ = _dafny.Array(None, 22)
        for i0_81_ in range(nw515_.length(0)):
            nw515_[i0_81_] = init81_(i0_81_)
        d_3061_v0_ = nw515_
        d_3061_v0_ = d_3061_v0_
        (globalState).f0 = 748
        d_3064_i1_: int
        d_3064_i1_ = 0
        with _dafny.label("17"):
            while default__.fm4(globalState):
                with _dafny.c_label("17"):
                    if (d_3064_i1_) >= (100):
                        raise _dafny.Break("17")
                    d_3064_i1_ = (d_3064_i1_) + (1)
                    d_3065_v1_: _dafny.Array
                    def lambda152_(d_3066_i2_):
                        return (396) <= (999)

                    init82_ = lambda152_
                    nw516_ = _dafny.Array(None, 4)
                    for i0_82_ in range(nw516_.length(0)):
                        nw516_[i0_82_] = init82_(i0_82_)
                    d_3065_v1_ = nw516_
                    d_3067_v2_: bool
                    d_3067_v2_ = True
                    index510_ = default__.safeIndex(322, (d_3065_v1_).length(0))
                    (d_3065_v1_)[index510_] = not(d_3067_v2_)
                    d_3068_v3_: int
                    d_3068_v3_ = 280
                    index511_ = default__.safeIndex(322, (d_3065_v1_).length(0))
                    (d_3065_v1_)[index511_] = (self).fm7(d_3067_v2_, (p1)[default__.safeIndex(d_3068_v3_, len(p1))], globalState)
                    d_3069_v4_: D0
                    d_3069_v4_ = D0_DC1(d_3068_v3_)
                    rhs551_ = False
                    rhs552_ = (0) - ((d_3069_v4_).cf1)
                    lhs457_ = globalState
                    lhs458_ = globalState
                    lhs457_.f27 = rhs551_
                    lhs458_.f0 = rhs552_
                    (globalState).f27 = (d_3065_v1_)[default__.safeIndex(322, (d_3065_v1_).length(0))]
                    d_3070_v5_: C7
                    nw517_ = C7()
                    nw517_.ctor__((self).f28)
                    d_3070_v5_ = nw517_
                    pass
            pass
        d_3071_v6_: _dafny.Seq
        d_3071_v6_ = _dafny.SeqWithoutIsStrInference([d_3061_v0_])
        d_3072_v7_: int
        d_3072_v7_ = 68
        d_3073_v8_: _dafny.Map
        d_3073_v8_ = _dafny.Map({(d_3071_v6_)[default__.safeIndex(d_3072_v7_, len(d_3071_v6_))]: d_3072_v7_})
        d_3074_v9_: str
        d_3074_v9_ = _dafny.CodePoint('t')
        d_3075_v10_: _dafny.Array
        nw518_ = _dafny.Array(_dafny.Map({}), 5)
        d_3075_v10_ = nw518_
        d_3076_v11_: bool
        d_3076_v11_ = True
        d_3077_v12_: _dafny.Map
        d_3077_v12_ = _dafny.Map({d_3076_v11_: d_3072_v7_})
        index512_ = default__.safeIndex(669, (d_3075_v10_).length(0))
        (d_3075_v10_)[index512_] = d_3077_v12_
        d_3078_v13_: _dafny.Seq
        d_3078_v13_ = _dafny.SeqWithoutIsStrInference([p2])
        d_3079_v14_: _dafny.Map
        d_3079_v14_ = _dafny.Map({d_3072_v7_: (d_3078_v13_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_3074_v9_ for d_3080_i3_ in range(default__.abs(894))])), len(d_3078_v13_))]})
        d_3081_v15_: _dafny.Map
        d_3081_v15_ = _dafny.Map({d_3076_v11_: d_3074_v9_})
        index513_ = default__.safeIndex(669, (d_3075_v10_).length(0))
        rhs553_ = d_3073_v8_
        rhs554_ = 915
        rhs555_ = default__.fm5(default__.fm2(((d_3079_v14_)[d_3072_v7_] if (d_3072_v7_) in (d_3079_v14_) else p2), _dafny.Set({default__.fm4(globalState)}), d_3081_v15_, d_3076_v11_, globalState), not (True) or (False), globalState)
        rhs556_ = d_3077_v12_
        lhs459_ = d_3075_v10_
        lhs460_ = default__.safeIndex(669, (d_3075_v10_).length(0))
        d_3073_v8_ = rhs553_
        d_3072_v7_ = rhs554_
        d_3074_v9_ = rhs555_
        lhs459_[lhs460_] = rhs556_
        (globalState).f27 = d_3076_v11_
        d_3082_v16_: _dafny.Array
        nw519_ = _dafny.Array(D22.default()(), 9)
        d_3082_v16_ = nw519_
        guard_loop_6_: int
        for guard_loop_6_ in _dafny.IntegerRange(0, (d_3082_v16_).length(0)):
            d_3083_i4_: int = guard_loop_6_
            if (True) and (((0) <= (d_3083_i4_)) and ((d_3083_i4_) < ((d_3082_v16_).length(0)))):
                (d_3082_v16_)[(d_3083_i4_)] = D22_DC56((p2) + ((d_3078_v13_)[default__.safeIndex(d_3072_v7_, len(d_3078_v13_))]), d_3072_v7_, d_3076_v11_, d_3076_v11_)
        d_3084_v17_: D2
        d_3084_v17_ = D2_DC7(d_3076_v11_, not(True), ((p2) + (p2)).set(default__.safeIndex(d_3072_v7_, len((p2) + (p2))), d_3074_v9_))
        r0 = d_3084_v17_
        d_3085_v18_: _dafny.Set
        d_3085_v18_ = _dafny.Set({d_3076_v11_, d_3076_v11_, d_3076_v11_, d_3076_v11_})
        r1 = d_3085_v18_
        return r0, r1

    def m4(self, p0, p1, globalState):
        r0: int = int(0)
        d_3086_v0_: bool
        d_3086_v0_ = True
        d_3087_v1_: _dafny.Seq
        d_3087_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jpupckeo"))
        d_3088_v2_: int
        d_3088_v2_ = 807
        d_3089_v3_: _dafny.Map
        d_3089_v3_ = _dafny.Map({len(d_3087_v1_): d_3088_v2_})
        d_3090_v4_: D0
        d_3090_v4_ = D0_DC0(d_3089_v3_)
        pat_let_tv55_ = d_3086_v0_
        pat_let_tv56_ = p1
        pat_let_tv57_ = p1
        pat_let_tv58_ = d_3086_v0_
        pat_let_tv59_ = d_3088_v2_
        def lambda153_(source42_):
            if source42_.is_DC1:
                d_3091___mcc_h0_ = source42_.cf1
                d_3092_cf1_ = d_3091___mcc_h0_
                return d_3092_cf1_
            elif source42_.is_DC2:
                d_3093___mcc_h1_ = source42_.cf2
                d_3094___mcc_h2_ = source42_.cf3
                d_3095_cf3_ = d_3094___mcc_h2_
                d_3096_cf2_ = d_3093___mcc_h1_
                return d_3095_cf3_
            elif True:
                d_3097___mcc_h3_ = source42_.cf0
                d_3098_cf0_ = d_3097___mcc_h3_
                if (pat_let_tv55_) in (pat_let_tv56_):
                    return (pat_let_tv57_)[pat_let_tv58_]
                elif True:
                    return pat_let_tv59_

        r0 = lambda153_((d_3090_v4_ if d_3086_v0_ else d_3090_v4_))
        d_3099_v5_: _dafny.Map
        d_3099_v5_ = _dafny.Map({d_3086_v0_: p0})
        hi16_ = (0) - (d_3088_v2_)
        for d_3100_i0_ in range(len((d_3099_v5_) | (d_3099_v5_)), hi16_):
            d_3101_v6_: str
            d_3101_v6_ = _dafny.CodePoint('l')
            d_3102_v7_: _dafny.Set
            d_3102_v7_ = _dafny.Set({d_3101_v6_, d_3101_v6_, default__.fm5(d_3100_i0_, d_3086_v0_, globalState)})
            d_3103_v8_: C1
            nw520_ = C1()
            nw520_.ctor__(d_3088_v2_, len(d_3102_v7_), (self).f28)
            d_3103_v8_ = nw520_
            d_3104_v9_: T0
            nw521_ = C5()
            nw521_.ctor__(d_3100_i0_, d_3086_v0_, (self).f28)
            d_3104_v9_ = nw521_
            d_3105_v10_: _dafny.Array
            nw522_ = _dafny.Array(None, 20)
            d_3105_v10_ = nw522_
            index514_ = default__.safeIndex(55, (d_3105_v10_).length(0))
            (d_3105_v10_)[index514_] = d_3104_v9_
            d_3106_v11_: C6
            nw523_ = C6()
            nw523_.ctor__(d_3103_v8_.f42)
            d_3106_v11_ = nw523_
            d_3107_v12_: _dafny.Seq
            d_3107_v12_ = _dafny.SeqWithoutIsStrInference([not(d_3086_v0_), d_3086_v0_])
            d_3108_v13_: _dafny.Map
            d_3108_v13_ = _dafny.Map({d_3107_v12_: 267})
            d_3109_v14_: _dafny.Set
            d_3109_v14_ = _dafny.Set({not(d_3086_v0_)})
            d_3110_v15_: C2
            nw524_ = C2()
            nw524_.ctor__(d_3108_v13_, d_3109_v14_, (self).f28)
            d_3110_v15_ = nw524_
            d_3111_v16_: _dafny.Map
            d_3111_v16_ = _dafny.Map({d_3110_v15_: d_3086_v0_})
            index515_ = default__.safeIndex(55, (d_3105_v10_).length(0))
            rhs557_ = (d_3110_v15_) in (d_3111_v16_)
            rhs558_ = d_3104_v9_
            rhs559_ = d_3104_v9_
            rhs560_ = d_3106_v11_
            lhs461_ = globalState
            lhs462_ = d_3105_v10_
            lhs463_ = default__.safeIndex(55, (d_3105_v10_).length(0))
            lhs461_.f11 = rhs557_
            d_3104_v9_ = rhs558_
            lhs462_[lhs463_] = rhs559_
            d_3106_v11_ = rhs560_
            index516_ = default__.safeIndex(552, (p0).length(0))
            (p0)[index516_] = (p1).ispropersubset(p1)
            d_3112_v17_: _dafny.Map
            d_3112_v17_ = _dafny.Map({((p1)[d_3086_v0_] if (d_3086_v0_) in (p1) else (d_3103_v8_).f41): d_3086_v0_})
            d_3113_v18_: _dafny.Map
            d_3113_v18_ = _dafny.Map({_dafny.Map({d_3088_v2_: d_3086_v0_}): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_3114_i1_ in range(default__.abs(963))])})
            index517_ = default__.safeIndex(552, (p0).length(0))
            (p0)[index517_] = not(not((d_3112_v17_) not in (d_3113_v18_)))
            d_3087_v1_ = d_3087_v1_
        d_3115_v19_: _dafny.Seq
        d_3115_v19_ = _dafny.SeqWithoutIsStrInference([d_3086_v0_, False])
        d_3116_v20_: _dafny.Array
        nw525_ = _dafny.Array(int(0), 7)
        d_3116_v20_ = nw525_
        d_3117_v21_: D2
        d_3117_v21_ = D2_DC8(d_3086_v0_, d_3116_v20_)
        d_3118_v22_: _dafny.Map
        d_3118_v22_ = _dafny.Map({-866: True})
        d_3119_v23_: _dafny.Seq
        d_3119_v23_ = _dafny.SeqWithoutIsStrInference([d_3088_v2_])
        d_3120_v24_: _dafny.Set
        d_3120_v24_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([d_3088_v2_]), d_3119_v23_})
        d_3121_v25_: _dafny.MultiSet
        d_3121_v25_ = _dafny.MultiSet([len(d_3087_v1_)])
        d_3122_v26_: _dafny.Set
        d_3122_v26_ = _dafny.Set({d_3086_v0_, True})
        d_3123_v27_: _dafny.Map
        d_3123_v27_ = _dafny.Map({not(default__.fm4(globalState)): d_3086_v0_})
        d_3124_v28_: _dafny.Set
        d_3124_v28_ = _dafny.Set({(0) - (((d_3121_v25_)[len(d_3123_v27_)] if (len(d_3123_v27_)) in (d_3121_v25_) else d_3088_v2_))})
        d_3125_v29_: _dafny.Array
        nw526_ = _dafny.Array(None, 26)
        nw526_[int(0)] = not(d_3086_v0_)
        nw526_[int(1)] = d_3086_v0_
        nw526_[int(2)] = d_3086_v0_
        nw526_[int(3)] = (d_3115_v19_)[default__.safeIndex((0) - (d_3088_v2_), len(d_3115_v19_))]
        nw526_[int(4)] = (d_3117_v21_).cf16
        nw526_[int(5)] = d_3086_v0_
        nw526_[int(6)] = d_3086_v0_
        nw526_[int(7)] = d_3086_v0_
        nw526_[int(8)] = True
        nw526_[int(9)] = (d_3086_v0_) == (d_3086_v0_)
        nw526_[int(10)] = d_3086_v0_
        nw526_[int(11)] = ((self).fm7(d_3086_v0_, ((d_3118_v22_)[d_3088_v2_] if (d_3088_v2_) in (d_3118_v22_) else d_3086_v0_), globalState) if False else ((self).f28).cf4)
        nw526_[int(12)] = d_3086_v0_
        nw526_[int(13)] = not(d_3086_v0_)
        nw526_[int(14)] = (default__.fm53(d_3088_v2_, d_3088_v2_, d_3086_v0_, d_3086_v0_, globalState)).issubset(d_3120_v24_)
        nw526_[int(15)] = d_3086_v0_
        nw526_[int(16)] = d_3086_v0_
        nw526_[int(17)] = (d_3088_v2_) not in (d_3121_v25_)
        nw526_[int(18)] = d_3086_v0_
        nw526_[int(19)] = not((default__.fm5(d_3088_v2_, d_3086_v0_, globalState)) not in (d_3087_v1_))
        nw526_[int(20)] = not(d_3086_v0_)
        nw526_[int(21)] = (d_3122_v26_).ispropersubset(d_3122_v26_)
        nw526_[int(22)] = (True if d_3086_v0_ else d_3086_v0_)
        nw526_[int(23)] = (_dafny.Set({d_3088_v2_, (0) - (d_3088_v2_)})).ispropersubset(d_3124_v28_)
        nw526_[int(24)] = d_3086_v0_
        nw526_[int(25)] = d_3086_v0_
        d_3125_v29_ = nw526_
        guard_loop_7_: int
        for guard_loop_7_ in _dafny.IntegerRange(0, (d_3125_v29_).length(0)):
            d_3126_i2_: int = guard_loop_7_
            if (True) and (((0) <= (d_3126_i2_)) and ((d_3126_i2_) < ((d_3125_v29_).length(0)))):
                (d_3125_v29_)[(d_3126_i2_)] = not(not ((d_3086_v0_) in (_dafny.Map({d_3086_v0_: d_3088_v2_}))) or ((-290) < (len(d_3087_v1_))))
        d_3127_v30_: _dafny.Array
        def lambda154_(d_3128_v25_):
            def lambda155_(d_3129_i4_):
                return (d_3128_v25_) | (d_3128_v25_)

            return lambda155_

        init83_ = lambda154_(d_3121_v25_)
        nw527_ = _dafny.Array(None, 14)
        for i0_83_ in range(nw527_.length(0)):
            nw527_[i0_83_] = init83_(i0_83_)
        d_3127_v30_ = nw527_
        guard_loop_8_: int
        for guard_loop_8_ in _dafny.IntegerRange(0, (d_3127_v30_).length(0)):
            d_3130_i3_: int = guard_loop_8_
            if (True) and (((0) <= (d_3130_i3_)) and ((d_3130_i3_) < ((d_3127_v30_).length(0)))):
                (d_3127_v30_)[(d_3130_i3_)] = (d_3121_v25_).set(d_3088_v2_, default__.abs(default__.safeDivisionInt(d_3088_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))))))
        d_3089_v3_ = (d_3089_v3_).set(d_3088_v2_, d_3088_v2_)
        d_3131_v31_: _dafny.Map
        d_3131_v31_ = _dafny.Map({d_3122_v26_: len(d_3119_v23_)})
        (globalState).f27 = ((d_3088_v2_) + (d_3088_v2_)) <= ((len(d_3124_v28_)) + ((0) - (len(d_3131_v31_))))
        d_3132_v32_: _dafny.Seq
        d_3132_v32_ = _dafny.SeqWithoutIsStrInference([d_3124_v28_])
        d_3133_v33_: _dafny.Seq
        d_3133_v33_ = d_3132_v32_
        d_3134_v34_: _dafny.MultiSet
        d_3134_v34_ = _dafny.MultiSet([d_3133_v33_])
        r0 = (d_3134_v34_).cardinality
        return r0

