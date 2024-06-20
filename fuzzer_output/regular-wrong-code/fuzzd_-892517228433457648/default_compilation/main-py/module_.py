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
        return (0) - (-194)

    @staticmethod
    def fm1(p0, p1, globalState):
        return not(((False if True else True) if False else (_dafny.SeqWithoutIsStrInference([506, 281])) <= (_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([False, True])), len(_dafny.Set({_dafny.CodePoint('s'), _dafny.CodePoint('b')}))]))))

    @staticmethod
    def fm2(p0, p1, p2, globalState):
        if (False if True else False):
            return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_0_i0_ in range(default__.abs(43))])
        elif True:
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sgxcrrbck"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('k') for d_1_i1_ in range(default__.abs(992))]))

    @staticmethod
    def fm3(p0, p1, globalState):
        return _dafny.CodePoint('c')

    @staticmethod
    def fm4(p0, p1, globalState):
        return D0_DC2((D0_DC2(_dafny.MultiSet([not(False), True]), False, (0) - (len(_dafny.SeqWithoutIsStrInference([True]))), False)).cf4, True, -858, (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_2_i0_ in range(default__.abs(183))])) == (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_3_i1_ in range(default__.abs(-240))])))

    @staticmethod
    def fm7(globalState):
        return _dafny.Map({len(_dafny.SeqWithoutIsStrInference([720 for d_4_i0_ in range(default__.abs(897))])): ((D0_DC2(_dafny.MultiSet([not(True)]), True, 611, False)).cf6 if False else -491)})

    @staticmethod
    def fm14(p0, p1, globalState):
        def iife0_():
            coll0_ = _dafny.Set()
            compr_0_: int
            for compr_0_ in _dafny.IntegerRange(869, -384):
                d_5_v0_: int = compr_0_
                if ((869) <= (d_5_v0_)) and ((d_5_v0_) < (-384)):
                    coll0_ = coll0_.union(_dafny.Set([(d_5_v0_) * ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([421, -384, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfdgtb"))), 389])).cardinality])) for d_6_i1_ in range(default__.abs(-635))]))).cardinality)]))
            return _dafny.Set(coll0_)
        return (_dafny.MultiSet([_dafny.Set({len(_dafny.Set({True}))})])) - (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([iife0_()
 for d_7_i0_ in range(default__.abs(-621))])))

    @staticmethod
    def fm15(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(_dafny.Map({True: True})) | (_dafny.Map({False: False})) for d_8_i0_ in range(default__.abs(286))])

    @staticmethod
    def fm16(globalState):
        return _dafny.Set({(False if not(not(False)) else not(not(False))), False, True})

    @staticmethod
    def fm17(globalState):
        return D2_DC7(not((380) != ((0) - ((0) - (len(_dafny.Set({-846, len(_dafny.SeqWithoutIsStrInference([True]))})))))), (0) - (default__.safeModuloInt(-811, 752)), (not(True) if True else False), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uyiex")))

    @staticmethod
    def fm18(p0, p1, p2, p3, globalState):
        def iife1_():
            coll1_ = _dafny.Map()
            compr_1_: int
            for compr_1_ in (_dafny.SeqWithoutIsStrInference([78])).Elements:
                d_9_v0_: int = compr_1_
                if (d_9_v0_) in (_dafny.SeqWithoutIsStrInference([78])):
                    coll1_[(d_9_v0_) + ((D20_DC40(False, _dafny.CodePoint('s'), _dafny.CodePoint('e'), len(_dafny.Map({False: True})))).cf68)] = 823
            return _dafny.Map(coll1_)
        def iife2_():
            coll2_ = _dafny.Set()
            compr_2_: _dafny.MultiSet
            for compr_2_ in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('w'): 176}))]) for d_11_i1_ in range(default__.abs(757))])).Elements:
                d_12_v1_: _dafny.MultiSet = compr_2_
                if (d_12_v1_) in (_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([len(_dafny.Map({_dafny.CodePoint('w'): 176}))]) for d_11_i1_ in range(default__.abs(757))])):
                    coll2_ = coll2_.union(_dafny.Set([d_12_v1_]))
            return _dafny.Set(coll2_)
        return (_dafny.Set({_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference([False]))]), (_dafny.MultiSet([329])), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.MultiSet([True]), _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))]))]))})) | ((_dafny.Set({_dafny.MultiSet([(_dafny.MultiSet([671, len(_dafny.SeqWithoutIsStrInference([False]))])).cardinality, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxlacl")))]), (_dafny.MultiSet([(_dafny.MultiSet([len(_dafny.Map({not(True): 511})), 604, len(_dafny.Set({-660, 685})), len(iife1_()
        )])).cardinality, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_10_i0_ in range(default__.abs(702))]))]))})) | (iife2_()
        ))

    @staticmethod
    def fm21(p0, globalState):
        def iife3_():
            coll3_ = _dafny.Set()
            compr_3_: int
            for compr_3_ in _dafny.IntegerRange(793, 442):
                d_13_v0_: int = compr_3_
                if ((793) <= (d_13_v0_)) and ((d_13_v0_) < (442)):
                    coll3_ = coll3_.union(_dafny.Set([(d_13_v0_) - (len(_dafny.SeqWithoutIsStrInference([not(False)])))]))
            return _dafny.Set(coll3_)
        return (_dafny.Map({iife3_()
        : len(_dafny.Set({True}))})) | (_dafny.Map({_dafny.Set({(0) - (len(_dafny.Map({len(_dafny.Map({True: False})): False}))), (0) - ((0) - (-545))}): 51}))

    @staticmethod
    def fm22(p0, p1, p2, p3, globalState):
        return (_dafny.Map({_dafny.CodePoint('y'): True})) | ((_dafny.Map({_dafny.CodePoint('c'): False}) if True else _dafny.Map({_dafny.CodePoint('p'): True})))

    @staticmethod
    def fm23(p0, globalState):
        if False:
            def iife4_():
                coll4_ = _dafny.Set()
                compr_4_: int
                for compr_4_ in _dafny.IntegerRange(244, 762):
                    d_14_v0_: int = compr_4_
                    if ((244) <= (d_14_v0_)) and ((d_14_v0_) < (762)):
                        coll4_ = coll4_.union(_dafny.Set([default__.safeDivisionInt(d_14_v0_, len(_dafny.Set({False})))]))
                return _dafny.Set(coll4_)
            return (iife4_()
            ).intersection(_dafny.Set({-694, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference([True])): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_15_i0_ in range(default__.abs(-387))]))}))}))
        elif True:
            return (_dafny.Set({(_dafny.MultiSet([not(True), False, False, True])).cardinality})).intersection(_dafny.Set({348}))

    @staticmethod
    def fm24(p0, p1, p2, globalState):
        return _dafny.Map({default__.safeModuloInt(91, 768): _dafny.CodePoint('o')})

    @staticmethod
    def fm25(p0, globalState):
        return (_dafny.Map({False: -954})) | (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([False]))}))

    @staticmethod
    def fm26(p0, p1, p2, globalState):
        return (D5_DC13(_dafny.SeqWithoutIsStrInference([-784]))).cf25

    @staticmethod
    def fm27(p0, globalState):
        def iife5_():
            coll5_ = _dafny.Map()
            compr_5_: int
            for compr_5_ in _dafny.IntegerRange(284, 491):
                d_16_v0_: int = compr_5_
                if ((284) <= (d_16_v0_)) and ((d_16_v0_) < (491)):
                    coll5_[default__.safeDivisionInt(d_16_v0_, len(_dafny.Map({True: -314})))] = not(False)
            return _dafny.Map(coll5_)
        def iife6_():
            coll6_ = _dafny.Map()
            compr_6_: int
            for compr_6_ in (_dafny.Map({-785: len(_dafny.SeqWithoutIsStrInference([277 for d_17_i0_ in range(default__.abs(744))]))})).keys.Elements:
                d_18_v1_: int = compr_6_
                if (d_18_v1_) in (_dafny.Map({-785: len(_dafny.SeqWithoutIsStrInference([277 for d_17_i0_ in range(default__.abs(744))]))})):
                    coll6_[(d_18_v1_) - (len(_dafny.Map({not(True): 245})))] = True
            return _dafny.Map(coll6_)
        return ((iife5_()
        ) | (_dafny.Map({(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([846, 273]))).cardinality: (D2_DC7(False, 544, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xoq")))).cf11}))) | ((iife6_()
        ) | (_dafny.Map({39: False})))

    @staticmethod
    def fm28(p0, p1, p2, p3, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ieinao")) for d_19_i0_ in range(default__.abs(-791))])

    @staticmethod
    def fm31(globalState):
        return ((_dafny.Set({_dafny.SeqWithoutIsStrInference([-281]), _dafny.SeqWithoutIsStrInference([len(_dafny.Set({-172}))])})).intersection(_dafny.Set({_dafny.SeqWithoutIsStrInference([386, 340, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ropewiho"))), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "feeuekkg"))), len(_dafny.Map({373: 155}))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([False])).cardinality, len(_dafny.Set({True}))]))])}))) | ((_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({-885, 23})) for d_20_i0_ in range(default__.abs(859))]))]), _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: _dafny.SeqWithoutIsStrInference([998 for d_21_i1_ in range(default__.abs(494))])}))]))])})) | (_dafny.Set({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "irwxgif"))) for d_22_i2_ in range(default__.abs(-50))])})))

    @staticmethod
    def fm32(globalState):
        def iife7_():
            def iife9_():
                coll9_ = _dafny.Map()
                compr_9_: int
                for compr_9_ in (_dafny.Map({-128: 460})).keys.Elements:
                    d_23_v0_: int = compr_9_
                    if (d_23_v0_) in (_dafny.Map({-128: 460})):
                        coll9_[(d_23_v0_) * ((0) - (len(_dafny.Map({False: 629}))))] = _dafny.SeqWithoutIsStrInference([363])
                return _dafny.Map(coll9_)
            coll7_ = _dafny.Set()
            def iife8_():
                coll8_ = _dafny.Map()
                compr_8_: int
                for compr_8_ in (_dafny.Map({-128: 460})).keys.Elements:
                    d_23_v0_: int = compr_8_
                    if (d_23_v0_) in (_dafny.Map({-128: 460})):
                        coll8_[(d_23_v0_) * ((0) - (len(_dafny.Map({False: 629}))))] = _dafny.SeqWithoutIsStrInference([363])
                return _dafny.Map(coll8_)
            compr_7_: int
            for compr_7_ in (iife8_()
            ).keys.Elements:
                d_24_v1_: int = compr_7_
                if (d_24_v1_) in (iife9_()
                ):
                    coll7_ = coll7_.union(_dafny.Set([default__.safeModuloInt(d_24_v1_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r"))))]))
            return _dafny.Set(coll7_)
        def iife10_():
            coll10_ = _dafny.Set()
            compr_10_: int
            for compr_10_ in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([not(True), False, False, False, not(True)]))})).Elements:
                d_25_v2_: int = compr_10_
                if (d_25_v2_) in (_dafny.Set({len(_dafny.SeqWithoutIsStrInference([not(True), False, False, False, not(True)]))})):
                    coll10_ = coll10_.union(_dafny.Set([default__.safeModuloInt(d_25_v2_, 168)]))
            return _dafny.Set(coll10_)
        def iife11_():
            coll11_ = _dafny.Set()
            compr_11_: int
            for compr_11_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eneorssfn")))})).keys.Elements:
                d_26_v3_: int = compr_11_
                if (d_26_v3_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference([False])): len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eneorssfn")))})):
                    coll11_ = coll11_.union(_dafny.Set([(d_26_v3_) * ((0) - (len(_dafny.Map({True: True}))))]))
            return _dafny.Set(coll11_)
        def iife12_():
            coll12_ = _dafny.Set()
            compr_12_: int
            for compr_12_ in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({490: not(False)}))])).Elements:
                d_27_v4_: int = compr_12_
                if (d_27_v4_) in (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({490: not(False)}))])):
                    coll12_ = coll12_.union(_dafny.Set([default__.safeModuloInt(d_27_v4_, -749)]))
            return _dafny.Set(coll12_)
        return _dafny.SeqWithoutIsStrInference([(_dafny.Set({34})).intersection(iife7_()
        ), iife10_()
        , _dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ikr")))}), iife11_()
        , iife12_()
        ])

    @staticmethod
    def fm33(p0, globalState):
        return (_dafny.MultiSet([not(False), False])) - (_dafny.MultiSet([True]))

    @staticmethod
    def fm34(p0, p1, globalState):
        return ((_dafny.Map({True: not(not(True))})) | (_dafny.Map({False: True}))) | (_dafny.Map({True: not(True)}))

    @staticmethod
    def fm35(p0, p1, globalState):
        def iife13_():
            coll13_ = _dafny.Map()
            compr_13_: int
            for compr_13_ in (_dafny.Set({304})).Elements:
                d_30_v0_: int = compr_13_
                if (d_30_v0_) in (_dafny.Set({304})):
                    def iife14_():
                        coll14_ = _dafny.Set()
                        compr_14_: int
                        for compr_14_ in _dafny.IntegerRange(361, 168):
                            d_31_v1_: int = compr_14_
                            if ((361) <= (d_31_v1_)) and ((d_31_v1_) < (168)):
                                coll14_ = coll14_.union(_dafny.Set([(d_31_v1_) - (833)]))
                        return _dafny.Set(coll14_)
                    coll13_[default__.safeModuloInt(d_30_v0_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lbekjpmmt"))))] = len(iife14_()
                    )
            return _dafny.Map(coll13_)
        return ((_dafny.SeqWithoutIsStrInference([D5_DC13(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({False: len(_dafny.Map({(0) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([not(False)]))).cardinality): 280}))})) for d_28_i0_ in range(default__.abs(351))])), D5_DC13(_dafny.SeqWithoutIsStrInference([-705 for d_29_i1_ in range(default__.abs(714))])), D5_DC13(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xs")))): False}))]))])) + (_dafny.SeqWithoutIsStrInference([D5_DC13(_dafny.SeqWithoutIsStrInference([872]))]))) + (_dafny.SeqWithoutIsStrInference([D5_DC13(_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([len(iife13_()
)])).cardinality for d_32_i2_ in range(default__.abs(675))]))]))

    @staticmethod
    def fm36(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nurarvw")))) + (-32), (219) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pybfvhx"))))])

    @staticmethod
    def fm37(p0, p1, p2, p3, globalState):
        if not(False):
            return _dafny.Set({(_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('p'), _dafny.CodePoint('i')])).cardinality})
        elif True:
            def iife15_():
                coll15_ = _dafny.Set()
                compr_15_: int
                for compr_15_ in _dafny.IntegerRange(-57, 768):
                    d_33_v0_: int = compr_15_
                    if ((-57) <= (d_33_v0_)) and ((d_33_v0_) < (768)):
                        coll15_ = coll15_.union(_dafny.Set([(d_33_v0_) * (307)]))
                return _dafny.Set(coll15_)
            return iife15_()
            

    @staticmethod
    def fm38(p0, p1, p2, p3, globalState):
        return (_dafny.MultiSet([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "teicle")))])) | (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([True])), len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "levf")))])))

    @staticmethod
    def fm39(p0, globalState):
        return ((_dafny.SeqWithoutIsStrInference([not(False), False])) + (_dafny.SeqWithoutIsStrInference([(D22_DC46(False)).cf76]))) + ((_dafny.SeqWithoutIsStrInference([not(False)])) + (_dafny.SeqWithoutIsStrInference([False])))

    @staticmethod
    def fm40(p0, globalState):
        return D0_DC0((len(_dafny.Map({913: -795})) if not(True) else 764))

    @staticmethod
    def fm43(p0, p1, p2, p3, globalState):
        return D4_DC10(_dafny.Set({True, True}))

    @staticmethod
    def fm44(p0, p1, globalState):
        return ((_dafny.SeqWithoutIsStrInference([True, not(True)])) + (_dafny.SeqWithoutIsStrInference([False]))) + (_dafny.SeqWithoutIsStrInference([False]))

    @staticmethod
    def fm45(p0, p1, globalState):
        return ((_dafny.MultiSet([_dafny.CodePoint('f'), _dafny.CodePoint('k'), _dafny.CodePoint('o'), _dafny.CodePoint('n')]) if False else _dafny.MultiSet([_dafny.CodePoint('j')]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('g'), _dafny.CodePoint('o'), _dafny.CodePoint('k'), _dafny.CodePoint('k')])))

    @staticmethod
    def fm46(globalState):
        return _dafny.MultiSet([((_dafny.MultiSet([622])).cardinality) < ((D8_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vapfcanir")), 939)).cf35)])

    @staticmethod
    def fm49(p0, globalState):
        return D2_DC5(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('y') for d_34_i0_ in range(default__.abs(-885))]))

    @staticmethod
    def fm50(p0, p1, p2, globalState):
        return (_dafny.Map({True: len(_dafny.SeqWithoutIsStrInference([not(False)]))})) | (_dafny.Map({False: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "utwya")))}))

    @staticmethod
    def fm51(globalState):
        def iife16_():
            coll16_ = _dafny.Map()
            compr_16_: D0
            for compr_16_ in (_dafny.SeqWithoutIsStrInference([D0_DC0((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))))), D0_DC0(444), D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqobdcs"))))])).Elements:
                d_35_v0_: D0 = compr_16_
                if (d_35_v0_) in (_dafny.SeqWithoutIsStrInference([D0_DC0((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "o"))))), D0_DC0(444), D0_DC0(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tqobdcs"))))])):
                    coll16_[d_35_v0_] = _dafny.CodePoint('p')
            return _dafny.Map(coll16_)
        def iife17_():
            coll17_ = _dafny.Map()
            compr_17_: D0
            for compr_17_ in (_dafny.Map({D0_DC0(len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jminm")))), 48]))): len(_dafny.SeqWithoutIsStrInference([not(not(False))]))})).keys.Elements:
                d_36_v1_: D0 = compr_17_
                if (d_36_v1_) in (_dafny.Map({D0_DC0(len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jminm")))), 48]))): len(_dafny.SeqWithoutIsStrInference([not(not(False))]))})):
                    coll17_[d_36_v1_] = _dafny.CodePoint('v')
            return _dafny.Map(coll17_)
        return ((iife16_()
        ) | (_dafny.Map({D0_DC0(len(_dafny.SeqWithoutIsStrInference([(0) - (-783)]))): _dafny.CodePoint('l')}))) | (iife17_()
        )

    @staticmethod
    def fm52(globalState):
        def iife18_():
            coll18_ = _dafny.Map()
            compr_18_: int
            for compr_18_ in _dafny.IntegerRange(-508, 187):
                d_37_v0_: int = compr_18_
                if ((-508) <= (d_37_v0_)) and ((d_37_v0_) < (187)):
                    coll18_[(d_37_v0_) - (len(_dafny.SeqWithoutIsStrInference([-834, 291, 447, (_dafny.MultiSet([True])).cardinality, len(_dafny.Map({True: _dafny.CodePoint('f')}))])))] = 58
            return _dafny.Map(coll18_)
        def iife19_():
            def iife21_():
                coll21_ = _dafny.Map()
                compr_21_: int
                for compr_21_ in _dafny.IntegerRange(405, 737):
                    d_38_v2_: int = compr_21_
                    if ((405) <= (d_38_v2_)) and ((d_38_v2_) < (737)):
                        coll21_[default__.safeDivisionInt(d_38_v2_, 754)] = len(_dafny.Map({True: 666}))
                return _dafny.Map(coll21_)
            coll19_ = _dafny.Map()
            def iife20_():
                coll20_ = _dafny.Map()
                compr_20_: int
                for compr_20_ in _dafny.IntegerRange(405, 737):
                    d_38_v2_: int = compr_20_
                    if ((405) <= (d_38_v2_)) and ((d_38_v2_) < (737)):
                        coll20_[default__.safeDivisionInt(d_38_v2_, 754)] = len(_dafny.Map({True: 666}))
                return _dafny.Map(coll20_)
            compr_19_: int
            for compr_19_ in (_dafny.Map({len(_dafny.Set({_dafny.Map({-31: len(_dafny.Map({True: (0) - ((_dafny.MultiSet([(_dafny.MultiSet([True, False])).cardinality])).cardinality)}))}), iife20_()
            })): 988})).keys.Elements:
                d_39_v1_: int = compr_19_
                if (d_39_v1_) in (_dafny.Map({len(_dafny.Set({_dafny.Map({-31: len(_dafny.Map({True: (0) - ((_dafny.MultiSet([(_dafny.MultiSet([True, False])).cardinality])).cardinality)}))}), iife21_()
                })): 988})):
                    coll19_[default__.safeModuloInt(d_39_v1_, len(_dafny.SeqWithoutIsStrInference([37, 64])))] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_40_i0_ in range(default__.abs(394))]))
            return _dafny.Map(coll19_)
        return ((_dafny.Map({(0) - (len(_dafny.Map({False: True}))): (0) - ((_dafny.MultiSet([False, True])).cardinality)})) | (iife18_()
        )) | (iife19_()
        )

    @staticmethod
    def fm53(globalState):
        return (_dafny.Set({False, not(not(not(False))), True, True})) | (((D4_DC10(_dafny.Set({True, not(True)}))).cf18) | (_dafny.Set({False, True})))

    @staticmethod
    def fm54(p0, p1, p2, globalState):
        if (not(True)) == (False):
            return _dafny.SeqWithoutIsStrInference([(0) - (-476), 634, 444])
        elif True:
            return (_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.Set({990, 676}))), 891, (0) - ((D9_DC22(693, _dafny.Map({317: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmcxvlhm"))}))).cf40)])))])) + (_dafny.SeqWithoutIsStrInference([-25, len(_dafny.Map({False: 855}))]))

    @staticmethod
    def fm55(p0, p1, p2, globalState):
        return D5_DC14((False) and (False))

    @staticmethod
    def fm56(p0, p1, globalState):
        def iife22_():
            coll22_ = _dafny.Set()
            compr_22_: int
            for compr_22_ in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qac"))): True})).keys.Elements:
                d_41_v0_: int = compr_22_
                if (d_41_v0_) in (_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qac"))): True})):
                    coll22_ = coll22_.union(_dafny.Set([default__.safeModuloInt(d_41_v0_, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ylqio")))]))).cardinality)]))
            return _dafny.Set(coll22_)
        return ((_dafny.MultiSet([850])) | (_dafny.MultiSet([len(_dafny.Set({True}))]))) - ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({690: _dafny.CodePoint('v')})), len(_dafny.SeqWithoutIsStrInference([True, False, True, True])), 591]))).intersection(_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([-448, len(_dafny.Set({_dafny.Set({len(_dafny.Map({True: iife22_()
        }))}), _dafny.Set({len(_dafny.SeqWithoutIsStrInference([_dafny.Map({False: 436})]))})})), 507, len(_dafny.Map({490: len(_dafny.Set({True}))})), 96]))))

    @staticmethod
    def fm57(globalState):
        return D1_DC4(_dafny.CodePoint('j'))

    @staticmethod
    def fm58(p0, p1, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kbab"))) for d_42_i0_ in range(default__.abs(842))])) + (_dafny.SeqWithoutIsStrInference([len(_dafny.Map({not(True): _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([True]))}))])))

    @staticmethod
    def fm59(globalState):
        return _dafny.Map({(_dafny.SeqWithoutIsStrInference([not(True)])) + (_dafny.SeqWithoutIsStrInference([True, True])): (-846) + (141)})

    @staticmethod
    def fm60(p0, p1, globalState):
        return D8_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "scicnjy")), len(_dafny.SeqWithoutIsStrInference([-492, -464, 441])))

    @staticmethod
    def fm61(globalState):
        def iife23_():
            coll23_ = _dafny.Set()
            compr_23_: int
            for compr_23_ in _dafny.IntegerRange(-130, 889):
                d_43_v0_: int = compr_23_
                if ((-130) <= (d_43_v0_)) and ((d_43_v0_) < (889)):
                    coll23_ = coll23_.union(_dafny.Set([(d_43_v0_) * (19)]))
            return _dafny.Set(coll23_)
        return (_dafny.Map({len(_dafny.Map({False: True})): _dafny.SeqWithoutIsStrInference([True, True])})) | ((_dafny.Map({939: _dafny.SeqWithoutIsStrInference([False, False])})) | (_dafny.Map({(0) - (len(iife23_()
        )): _dafny.SeqWithoutIsStrInference([True, True, True])})))

    @staticmethod
    def fm62(p0, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.Map({len(_dafny.Map({(_dafny.MultiSet([False, False])).cardinality: (_dafny.MultiSet([True])).cardinality})): False})])

    @staticmethod
    def fm63(globalState):
        return _dafny.Set({(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kici"))) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_44_i0_ in range(default__.abs(998))]))})

    @staticmethod
    def m0(globalState):
        d_45_v0_: bool
        d_45_v0_ = False
        d_46_i0_: int
        d_46_i0_ = 0
        with _dafny.label("0"):
            while d_45_v0_:
                with _dafny.c_label("0"):
                    if (d_46_i0_) >= (100):
                        raise _dafny.Break("0")
                    d_46_i0_ = (d_46_i0_) + (1)
                    if d_45_v0_:
                        d_47_v1_: _dafny.Seq
                        d_47_v1_ = _dafny.SeqWithoutIsStrInference([d_45_v0_, d_45_v0_])
                        d_48_v2_: int
                        d_48_v2_ = -142
                        (globalState).f9 = default__.safeModuloInt(772, (len(d_47_v1_)) + (d_48_v2_))
                        d_49_v3_: str
                        d_49_v3_ = _dafny.CodePoint('m')
                        d_50_v4_: _dafny.Set
                        d_50_v4_ = _dafny.Set({default__.fm1(d_45_v0_, d_49_v3_, globalState), d_45_v0_, d_45_v0_})
                        (globalState).f8 = (d_50_v4_).isdisjoint(d_50_v4_)
                        d_51_v5_: _dafny.Array
                        nw0_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 6)
                        d_51_v5_ = nw0_
                        d_52_v6_: _dafny.Array
                        nw1_ = _dafny.Array(_dafny.Map({}), 22)
                        d_52_v6_ = nw1_
                        d_53_v7_: C5
                        nw2_ = C5()
                        nw2_.ctor__(d_52_v6_)
                        d_53_v7_ = nw2_
                        d_54_v8_: _dafny.MultiSet
                        d_54_v8_ = _dafny.MultiSet([d_53_v7_, d_53_v7_])
                        d_55_v9_: _dafny.Seq
                        d_55_v9_ = _dafny.SeqWithoutIsStrInference([d_49_v3_])
                        d_56_v10_: _dafny.Map
                        d_56_v10_ = _dafny.Map({((d_54_v8_).cardinality) <= (d_48_v2_): d_55_v9_})
                        d_57_v11_: _dafny.Set
                        d_57_v11_ = _dafny.Set({len(d_55_v9_), len(d_47_v1_)})
                        d_58_v12_: D21
                        d_58_v12_ = D21_DC42(d_51_v5_)
                        d_59_v13_: _dafny.Seq
                        d_59_v13_ = _dafny.SeqWithoutIsStrInference([d_56_v10_, d_56_v10_])
                        rhs0_ = d_45_v0_
                        rhs1_ = (default__.safeModuloInt((0) - (len(d_57_v11_)), d_48_v2_)) * (d_48_v2_)
                        rhs2_ = (d_58_v12_).cf70
                        rhs3_ = (d_59_v13_)[default__.safeIndex(d_48_v2_, len(d_59_v13_))]
                        lhs0_ = globalState
                        lhs1_ = globalState
                        lhs0_.f8 = rhs0_
                        lhs1_.f9 = rhs1_
                        d_51_v5_ = rhs2_
                        d_56_v10_ = rhs3_
                        (globalState).f9 = (len(((d_47_v1_).set(default__.safeIndex(d_48_v2_, len(d_47_v1_)), d_45_v0_)) + (_dafny.SeqWithoutIsStrInference([d_45_v0_])))) - (d_48_v2_)
                        d_60_v14_: _dafny.Array
                        nw3_ = _dafny.Array(int(0), 11)
                        d_60_v14_ = nw3_
                        index0_ = default__.safeIndex(217, (d_60_v14_).length(0))
                        (d_60_v14_)[index0_] = d_48_v2_
                        d_61_v15_: _dafny.Seq
                        d_61_v15_ = _dafny.SeqWithoutIsStrInference([d_48_v2_])
                        index1_ = default__.safeIndex(217, (d_60_v14_).length(0))
                        (d_60_v14_)[index1_] = (d_61_v15_)[default__.safeIndex(94, len(d_61_v15_))]
                    elif True:
                        d_62_v16_: _dafny.Seq
                        d_62_v16_ = _dafny.SeqWithoutIsStrInference([d_45_v0_, d_45_v0_])
                        d_63_v17_: int
                        d_63_v17_ = 55
                        d_64_v18_: D8
                        d_64_v18_ = D8_DC20(D0_DC2(_dafny.MultiSet(d_62_v16_), d_45_v0_, d_63_v17_, d_45_v0_), d_45_v0_, d_45_v0_)
                        d_45_v0_ = (d_64_v18_).cf37
                        d_65_v19_: _dafny.Array
                        def lambda0_(d_66_v0_):
                            def lambda1_(d_67_i1_):
                                return (_dafny.Set({d_66_v0_, d_66_v0_, d_66_v0_, True})).intersection(_dafny.Set({not(d_66_v0_), d_66_v0_, not(d_66_v0_), not(d_66_v0_), d_66_v0_}))

                            return lambda1_

                        init0_ = lambda0_(d_45_v0_)
                        nw4_ = _dafny.Array(None, 10)
                        for i0_0_ in range(nw4_.length(0)):
                            nw4_[i0_0_] = init0_(i0_0_)
                        d_65_v19_ = nw4_
                        d_68_v20_: _dafny.Set
                        d_68_v20_ = _dafny.Set({d_45_v0_, d_45_v0_, d_45_v0_, d_45_v0_, d_45_v0_})
                        index2_ = default__.safeIndex(229, (d_65_v19_).length(0))
                        (d_65_v19_)[index2_] = d_68_v20_
                        index3_ = default__.safeIndex(229, (d_65_v19_).length(0))
                        (d_65_v19_)[index3_] = d_68_v20_
                        d_69_v21_: _dafny.Seq
                        d_69_v21_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vtof"))
                        d_70_v22_: C6
                        nw5_ = C6()
                        nw5_.ctor__(d_63_v17_, d_69_v21_)
                        d_70_v22_ = nw5_
                        d_63_v17_ = (d_70_v22_).f15
                        (globalState).f9 = (d_70_v22_).f15
                    d_71_v23_: int
                    d_71_v23_ = 885
                    d_72_v24_: _dafny.Set
                    d_72_v24_ = _dafny.Set({(0) - ((d_71_v23_) + (d_71_v23_)), (0) - (-673)})
                    d_72_v24_ = (d_72_v24_) - ((_dafny.Set({d_71_v23_})).intersection(d_72_v24_))
                    if not(not (True) or (d_45_v0_)):
                        d_73_v25_: _dafny.Array
                        def lambda2_(d_74_v0_):
                            def lambda3_(d_75_i2_):
                                return d_74_v0_

                            return lambda3_

                        init1_ = lambda2_(d_45_v0_)
                        nw6_ = _dafny.Array(None, 8)
                        for i0_1_ in range(nw6_.length(0)):
                            nw6_[i0_1_] = init1_(i0_1_)
                        d_73_v25_ = nw6_
                        index4_ = default__.safeIndex(872, (d_73_v25_).length(0))
                        (d_73_v25_)[index4_] = False
                        d_76_v26_: _dafny.Map
                        d_76_v26_ = _dafny.Map({not(False): d_73_v25_})
                        d_77_v27_: _dafny.Map
                        d_77_v27_ = _dafny.Map({-216: d_45_v0_})
                        d_78_v28_: _dafny.Set
                        d_78_v28_ = _dafny.Set({((d_76_v26_)[((d_77_v27_)[(0) - ((0) - (d_71_v23_))] if ((0) - ((0) - (d_71_v23_))) in (d_77_v27_) else d_45_v0_)] if (((d_77_v27_)[(0) - ((0) - (d_71_v23_))] if ((0) - ((0) - (d_71_v23_))) in (d_77_v27_) else d_45_v0_)) in (d_76_v26_) else d_73_v25_), d_73_v25_, d_73_v25_, d_73_v25_, d_73_v25_})
                        d_79_v29_: _dafny.MultiSet
                        d_79_v29_ = _dafny.MultiSet([672])
                        d_80_v30_: str
                        d_80_v30_ = _dafny.CodePoint('h')
                        d_81_v31_: _dafny.MultiSet
                        d_81_v31_ = _dafny.MultiSet([d_45_v0_])
                        index5_ = default__.safeIndex(872, (d_73_v25_).length(0))
                        rhs4_ = len((d_78_v28_ if d_45_v0_ else (d_78_v28_) | (d_78_v28_)))
                        rhs5_ = (d_79_v29_).issubset(d_79_v29_)
                        rhs6_ = default__.fm1(d_45_v0_, d_80_v30_, globalState)
                        rhs7_ = (True) not in (d_81_v31_)
                        lhs2_ = globalState
                        lhs3_ = d_73_v25_
                        lhs4_ = default__.safeIndex(872, (d_73_v25_).length(0))
                        lhs5_ = globalState
                        lhs2_.f9 = rhs4_
                        d_45_v0_ = rhs5_
                        lhs3_[lhs4_] = rhs6_
                        lhs5_.f8 = rhs7_
                        d_45_v0_ = d_45_v0_
                        (globalState).f9 = d_71_v23_
                        d_82_v32_: D0
                        d_82_v32_ = D0_DC0(d_71_v23_)
                        d_83_v33_: C1
                        nw7_ = C1()
                        nw7_.ctor__(d_45_v0_, d_45_v0_, d_82_v32_, d_45_v0_)
                        d_83_v33_ = nw7_
                        d_84_v34_: D1
                        d_84_v34_ = D1_DC4(d_80_v30_)
                        d_85_v35_: _dafny.Seq
                        d_85_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "enequbdi"))
                        d_86_v36_: _dafny.Seq
                        d_86_v36_ = _dafny.SeqWithoutIsStrInference([d_85_v35_])
                        d_87_v37_: C6
                        nw8_ = C6()
                        nw8_.ctor__((d_83_v33_).fm9(d_84_v34_, d_72_v24_, globalState), (d_86_v36_)[default__.safeIndex(d_71_v23_, len(d_86_v36_))])
                        d_87_v37_ = nw8_
                    elif True:
                        d_88_v38_: _dafny.Map
                        d_88_v38_ = _dafny.Map({d_71_v23_: not(d_45_v0_)})
                        d_89_v40_: _dafny.Array
                        nw9_ = _dafny.Array(None, 13)
                        nw9_[int(0)] = d_88_v38_
                        nw9_[int(1)] = d_88_v38_
                        nw9_[int(2)] = d_88_v38_
                        def iife24_():
                            coll24_ = _dafny.Map()
                            compr_24_: int
                            for compr_24_ in _dafny.IntegerRange(647, 676):
                                d_90_v39_: int = compr_24_
                                if ((647) <= (d_90_v39_)) and ((d_90_v39_) < (676)):
                                    coll24_[(d_90_v39_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('o') for d_91_i3_ in range(default__.abs(174))])))] = True
                            return _dafny.Map(coll24_)
                        nw9_[int(3)] = iife24_()
                        
                        nw9_[int(4)] = d_88_v38_
                        nw9_[int(5)] = d_88_v38_
                        nw9_[int(6)] = d_88_v38_
                        nw9_[int(7)] = d_88_v38_
                        nw9_[int(8)] = d_88_v38_
                        nw9_[int(9)] = d_88_v38_
                        nw9_[int(10)] = d_88_v38_
                        nw9_[int(11)] = d_88_v38_
                        nw9_[int(12)] = d_88_v38_
                        d_89_v40_ = nw9_
                        d_92_v41_: C2
                        nw10_ = C2()
                        nw10_.ctor__(d_89_v40_)
                        d_92_v41_ = nw10_
                        d_92_v41_ = d_92_v41_
                        d_71_v23_ = (d_71_v23_ if d_45_v0_ else (d_71_v23_) * (default__.fm0(d_45_v0_, globalState)))
                        d_93_v42_: T0
                        nw11_ = C3()
                        nw11_.ctor__(d_45_v0_, not((d_45_v0_) and (d_45_v0_)), d_89_v40_)
                        d_93_v42_ = nw11_
                        d_93_v42_ = d_93_v42_
                        d_94_v43_: _dafny.Seq
                        d_94_v43_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxobqbdo"))
                        d_95_v44_: str
                        d_95_v44_ = _dafny.CodePoint('m')
                        d_96_v45_: _dafny.MultiSet
                        d_96_v45_ = _dafny.MultiSet([d_71_v23_])
                        d_97_v46_: _dafny.Map
                        d_97_v46_ = _dafny.Map({d_71_v23_: len(d_88_v38_)})
                        d_98_v47_: _dafny.Array
                        nw12_ = _dafny.Array(None, 23)
                        nw12_[int(0)] = (d_92_v41_).fm19(globalState)
                        nw12_[int(1)] = ((d_94_v43_).set(default__.safeIndex(d_71_v23_, len(d_94_v43_)), d_95_v44_)) + (d_94_v43_)
                        nw12_[int(2)] = d_94_v43_
                        nw12_[int(3)] = d_94_v43_
                        nw12_[int(4)] = (d_94_v43_) + (d_94_v43_)
                        nw12_[int(5)] = d_94_v43_
                        nw12_[int(6)] = d_94_v43_
                        nw12_[int(7)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aamc"))
                        nw12_[int(8)] = default__.fm2(d_71_v23_, d_96_v45_, d_97_v46_, globalState)
                        nw12_[int(9)] = d_94_v43_
                        nw12_[int(10)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s"))
                        nw12_[int(11)] = d_94_v43_
                        nw12_[int(12)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mrwlfrsf"))
                        nw12_[int(13)] = ((d_94_v43_) + (d_94_v43_)).set(default__.safeIndex(d_71_v23_, len((d_94_v43_) + (d_94_v43_))), d_95_v44_)
                        nw12_[int(14)] = (d_94_v43_) + (d_94_v43_)
                        nw12_[int(15)] = d_94_v43_
                        nw12_[int(16)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_99_i4_ in range(default__.abs(905))])
                        nw12_[int(17)] = d_94_v43_
                        nw12_[int(18)] = d_94_v43_
                        nw12_[int(19)] = d_94_v43_
                        nw12_[int(20)] = d_94_v43_
                        nw12_[int(21)] = d_94_v43_
                        nw12_[int(22)] = d_94_v43_
                        d_98_v47_ = nw12_
                        index6_ = default__.safeIndex(240, (d_98_v47_).length(0))
                        (d_98_v47_)[index6_] = (d_94_v43_) + (d_94_v43_)
                        d_100_v48_: D2
                        d_100_v48_ = D2_DC5(d_94_v43_)
                        index7_ = default__.safeIndex(240, (d_98_v47_).length(0))
                        rhs8_ = (d_94_v43_) + (d_94_v43_)
                        rhs9_ = default__.fm1(default__.fm1(d_45_v0_, d_95_v44_, globalState), d_95_v44_, globalState)
                        rhs10_ = (d_94_v43_) <= ((d_100_v48_).cf10)
                        lhs6_ = d_98_v47_
                        lhs7_ = default__.safeIndex(240, (d_98_v47_).length(0))
                        lhs8_ = globalState
                        lhs9_ = globalState
                        lhs6_[lhs7_] = rhs8_
                        lhs8_.f8 = rhs9_
                        lhs9_.f8 = rhs10_
                        d_101_v49_: _dafny.Array
                        nw13_ = _dafny.Array(None, 5)
                        d_101_v49_ = nw13_
                        d_102_v50_: _dafny.Map
                        d_102_v50_ = _dafny.Map({d_101_v49_: (0) - (d_71_v23_)})
                        d_103_v51_: D22
                        d_103_v51_ = D22_DC45(d_101_v49_)
                        d_102_v50_ = (d_102_v50_).set((d_103_v51_).cf75, -627)
                    (globalState).f8 = d_45_v0_
                    pass
            pass
        d_104_v52_: _dafny.Array
        nw14_ = _dafny.Array(_dafny.Map({}), 5)
        d_104_v52_ = nw14_
        d_105_v53_: C3
        nw15_ = C3()
        nw15_.ctor__(d_45_v0_, False, d_104_v52_)
        d_105_v53_ = nw15_
        d_106_v54_: int
        d_106_v54_ = 546
        hi0_ = (0) - (default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nklwdo"))), 124))
        for d_107_i5_ in range(d_106_v54_, hi0_):
            d_108_v55_: C2
            nw16_ = C2()
            nw16_.ctor__(d_104_v52_)
            d_108_v55_ = nw16_
            d_109_v56_: _dafny.Seq
            d_109_v56_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rilg"))
            d_110_v57_: _dafny.Seq
            d_110_v57_ = _dafny.SeqWithoutIsStrInference([d_105_v53_.f21, d_105_v53_.f21])
            d_111_v58_: _dafny.MultiSet
            d_111_v58_ = _dafny.MultiSet([d_105_v53_.f21])
            d_112_v59_: D0
            d_112_v59_ = D0_DC0((0) - ((d_111_v58_).cardinality))
            d_113_v60_: str
            d_113_v60_ = _dafny.CodePoint('x')
            d_114_v61_: T2
            nw17_ = C1()
            nw17_.ctor__((d_106_v54_) < ((0) - (len(d_109_v56_))), (d_110_v57_)[default__.safeIndex(d_106_v54_, len(d_110_v57_))], d_112_v59_, (d_113_v60_) in (d_109_v56_))
            d_114_v61_ = nw17_
            d_114_v61_ = d_114_v61_
            if ((939 if (d_114_v61_).f13 else d_106_v54_)) <= (d_107_i5_):
                d_115_v62_: _dafny.Map
                d_115_v62_ = _dafny.Map({d_107_i5_: (d_105_v53_).f22})
                (globalState).f9 = len((d_115_v62_) | (d_115_v62_))
                d_116_v63_: _dafny.Array
                def lambda4_(d_117_v54_):
                    def lambda5_(d_118_i6_):
                        return (d_118_i6_) * (d_117_v54_)

                    return lambda5_

                init2_ = lambda4_(d_106_v54_)
                nw18_ = _dafny.Array(None, 17)
                for i0_2_ in range(nw18_.length(0)):
                    nw18_[i0_2_] = init2_(i0_2_)
                d_116_v63_ = nw18_
                d_119_v64_: _dafny.Seq
                d_119_v64_ = _dafny.SeqWithoutIsStrInference([d_115_v62_, d_115_v62_])
                d_120_v65_: _dafny.Set
                d_120_v65_ = _dafny.Set({len(d_119_v64_), d_106_v54_})
                d_121_v66_: C5
                nw19_ = C5()
                nw19_.ctor__(d_104_v52_)
                d_121_v66_ = nw19_
                d_122_v67_: _dafny.Map
                d_122_v67_ = _dafny.Map({d_121_v66_: d_107_i5_})
                d_123_v68_: _dafny.MultiSet
                d_123_v68_ = _dafny.MultiSet([d_122_v67_, d_122_v67_, d_122_v67_])
                rhs11_ = d_116_v63_
                rhs12_ = d_109_v56_
                rhs13_ = (d_120_v65_).issubset(d_120_v65_)
                rhs14_ = default__.safeDivisionInt((len(d_109_v56_)) - ((d_123_v68_).cardinality), d_106_v54_)
                rhs15_ = (d_110_v57_) + ((default__.fm44(_dafny.SeqWithoutIsStrInference([d_110_v57_ for d_124_i7_ in range(default__.abs(41))]), d_113_v60_, globalState)) + (d_110_v57_))
                lhs10_ = globalState
                lhs11_ = globalState
                d_116_v63_ = rhs11_
                d_109_v56_ = rhs12_
                lhs10_.f8 = rhs13_
                lhs11_.f9 = rhs14_
                d_110_v57_ = rhs15_
                d_125_v69_: _dafny.Array
                nw20_ = _dafny.Array(False, 19)
                d_125_v69_ = nw20_
                d_126_v70_: T0
                nw21_ = C4()
                nw21_.ctor__(d_104_v52_)
                d_126_v70_ = nw21_
                d_127_v71_: _dafny.Map
                d_127_v71_ = _dafny.Map({(d_114_v61_).f13: d_126_v70_})
                d_128_v72_: _dafny.MultiSet
                d_128_v72_ = _dafny.MultiSet([d_116_v63_, d_116_v63_])
                d_129_v73_: D0
                d_129_v73_ = D0_DC1(d_128_v72_, len(d_110_v57_), d_106_v54_)
                index8_ = default__.safeIndex(158, (d_125_v69_).length(0))
                (d_125_v69_)[index8_] = (len(d_127_v71_)) != ((d_129_v73_).cf3)
                index9_ = default__.safeIndex(158, (d_125_v69_).length(0))
                (d_125_v69_)[index9_] = default__.fm1((d_114_v61_).f13, d_113_v60_, globalState)
                rhs16_ = d_107_i5_
                rhs17_ = (d_105_v53_).f22
                lhs12_ = globalState
                lhs13_ = d_105_v53_
                lhs12_.f9 = rhs16_
                lhs13_.f21 = rhs17_
                d_130_v74_: C6
                nw22_ = C6()
                nw22_.ctor__(d_106_v54_, d_109_v56_)
                d_130_v74_ = nw22_
            elif True:
                d_131_v75_: _dafny.Map
                d_131_v75_ = _dafny.Map({(d_114_v61_).f13: (d_105_v53_).f22})
                d_132_v76_: _dafny.MultiSet
                d_132_v76_ = _dafny.MultiSet([d_131_v75_])
                d_132_v76_ = (d_132_v76_).intersection(((d_132_v76_).set(_dafny.Map({d_105_v53_.f21: (d_105_v53_).f22}), default__.abs((0) - (len(d_109_v56_)))) if (d_105_v53_).f22 else d_132_v76_))
                (globalState).f9 = default__.safeDivisionInt(d_106_v54_, d_106_v54_)
                d_133_v77_: _dafny.MultiSet
                d_133_v77_ = _dafny.MultiSet([d_106_v54_])
                d_134_v78_: _dafny.Map
                d_134_v78_ = _dafny.Map({d_107_i5_: d_106_v54_})
                d_135_v79_: _dafny.Map
                d_135_v79_ = _dafny.Map({d_106_v54_: d_134_v78_})
                d_136_v80_: _dafny.Map
                d_136_v80_ = d_134_v78_
                d_109_v56_ = default__.fm2(d_106_v54_, d_133_v77_, ((d_135_v79_)[len(_dafny.Map({d_105_v53_.f21: d_113_v60_}))] if (len(_dafny.Map({d_105_v53_.f21: d_113_v60_}))) in (d_135_v79_) else (d_136_v80_)), globalState)
                d_137_v81_: _dafny.Array
                def lambda6_(d_138_v53_):
                    def lambda7_(d_139_i8_):
                        return d_138_v53_.f21

                    return lambda7_

                init3_ = lambda6_(d_105_v53_)
                nw23_ = _dafny.Array(None, 10)
                for i0_3_ in range(nw23_.length(0)):
                    nw23_[i0_3_] = init3_(i0_3_)
                d_137_v81_ = nw23_
                index10_ = default__.safeIndex(204, (d_137_v81_).length(0))
                (d_137_v81_)[index10_] = d_105_v53_.f21
                index11_ = default__.safeIndex(204, (d_137_v81_).length(0))
                (d_137_v81_)[index11_] = (d_105_v53_).f22
                d_140_v82_: _dafny.Set
                d_140_v82_ = _dafny.Set({(d_105_v53_).f22})
                d_141_v83_: _dafny.Map
                d_141_v83_ = _dafny.Map({((d_114_v61_).f13) not in (d_140_v82_): d_109_v56_})
                d_141_v83_ = (d_141_v83_).set((d_114_v61_).f13, d_109_v56_)
            d_142_v84_: _dafny.Map
            d_142_v84_ = _dafny.Map({not (d_45_v0_) or ((d_105_v53_).f22): d_105_v53_.f21})
            d_143_v85_: _dafny.Map
            d_143_v85_ = d_142_v84_
            d_142_v84_ = (default__.fm34(not(d_105_v53_.f21), _dafny.MultiSet([(d_105_v53_).f22]), globalState)) | ((d_142_v84_) | ((d_143_v85_)))
        d_144_v86_: _dafny.Set
        d_144_v86_ = _dafny.Set({d_106_v54_})
        d_145_v87_: _dafny.Map
        d_145_v87_ = _dafny.Map({default__.safeModuloInt(d_106_v54_, len(d_144_v86_)): d_45_v0_})
        if ((d_145_v87_)[d_106_v54_] if (d_106_v54_) in (d_145_v87_) else (d_45_v0_) == (False)):
            d_106_v54_ = (len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_146_i9_ in range(default__.abs(366))])) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp"))))) - (default__.safeModuloInt(d_106_v54_, d_106_v54_))
            (d_105_v53_).f21 = d_45_v0_
            (d_105_v53_).f21 = (not (d_105_v53_.f21) or (d_105_v53_.f21) if d_45_v0_ else d_105_v53_.f21)
            d_145_v87_ = (d_145_v87_).set(d_106_v54_, d_105_v53_.f21)
            d_147_v88_: _dafny.Array
            def lambda8_(d_148_i10_):
                return default__.safeModuloInt(d_148_i10_, 650)

            init4_ = lambda8_
            nw24_ = _dafny.Array(None, 24)
            for i0_4_ in range(nw24_.length(0)):
                nw24_[i0_4_] = init4_(i0_4_)
            d_147_v88_ = nw24_
            d_149_v89_: _dafny.Map
            d_149_v89_ = _dafny.Map({d_106_v54_: d_147_v88_})
            d_150_v90_: D9
            d_150_v90_ = D9_DC21(((d_149_v89_)[d_106_v54_] if (d_106_v54_) in (d_149_v89_) else d_147_v88_))
            pat_let_tv0_ = d_147_v88_
            def iife25_(_pat_let0_0):
                def iife26_(d_151_dt__update__tmp_h0_):
                    def iife27_(_pat_let1_0):
                        def iife28_(d_152_dt__update_hcf39_h0_):
                            return D9_DC21(d_152_dt__update_hcf39_h0_)
                        return iife28_(_pat_let1_0)
                    return iife27_(pat_let_tv0_)
                return iife26_(_pat_let0_0)
            source0_ = iife25_(d_150_v90_)
            if source0_.is_DC22:
                d_153___mcc_h0_ = source0_.cf40
                d_154___mcc_h1_ = source0_.cf41
                d_155_cf41_ = d_154___mcc_h1_
                d_156_cf40_ = d_153___mcc_h0_
                d_157_v91_: _dafny.Seq
                d_157_v91_ = _dafny.SeqWithoutIsStrInference([(d_105_v53_).f22, d_105_v53_.f21])
                index12_ = default__.safeIndex(253, (d_147_v88_).length(0))
                (d_147_v88_)[index12_] = default__.fm0((d_157_v91_)[default__.safeIndex(d_106_v54_, len(d_157_v91_))], globalState)
                index13_ = default__.safeIndex(253, (d_147_v88_).length(0))
                (d_147_v88_)[index13_] = (0) - (d_106_v54_)
                d_158_v92_: _dafny.Map
                d_158_v92_ = _dafny.Map({d_145_v87_: False})
                d_159_v93_: _dafny.Seq
                d_159_v93_ = _dafny.SeqWithoutIsStrInference([d_106_v54_])
                d_160_v94_: _dafny.Map
                d_160_v94_ = _dafny.Map({(d_105_v53_).f22: d_159_v93_})
                d_161_v95_: D7
                d_161_v95_ = D7_DC16(d_157_v91_)
                d_162_v96_: str
                d_162_v96_ = _dafny.CodePoint('l')
                d_163_v97_: _dafny.Map
                d_163_v97_ = _dafny.Map({d_162_v96_: d_105_v53_.f21})
                d_164_v98_: _dafny.MultiSet
                d_164_v98_ = _dafny.MultiSet([len(d_163_v97_), d_106_v54_])
                d_165_v99_: _dafny.MultiSet
                d_165_v99_ = d_164_v98_
                d_166_v100_: _dafny.Set
                d_166_v100_ = _dafny.Set({d_165_v99_})
                d_167_v101_: _dafny.Seq
                d_167_v101_ = _dafny.SeqWithoutIsStrInference([default__.safeDivisionInt(len(d_160_v94_), d_106_v54_), -234, len((d_161_v95_).cf28), d_156_cf40_, len(d_166_v100_)])
                d_168_v102_: _dafny.Seq
                d_168_v102_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                d_169_v103_: _dafny.Seq
                d_169_v103_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_105_v53_.f21])])
                index14_ = default__.safeIndex(253, (d_147_v88_).length(0))
                rhs18_ = d_158_v92_
                rhs19_ = d_159_v93_
                rhs20_ = (d_168_v102_) < (d_168_v102_)
                rhs21_ = len(((d_157_v91_).set(default__.safeIndex(d_106_v54_, len(d_157_v91_)), False)) + (((_dafny.SeqWithoutIsStrInference([(d_105_v53_).f22, (d_105_v53_).f22, False, default__.fm1((d_105_v53_).f22, d_162_v96_, globalState)])) + (default__.fm44(d_169_v103_, d_162_v96_, globalState))).set(default__.safeIndex(d_106_v54_, len((_dafny.SeqWithoutIsStrInference([(d_105_v53_).f22, (d_105_v53_).f22, False, default__.fm1((d_105_v53_).f22, d_162_v96_, globalState)])) + (default__.fm44(d_169_v103_, d_162_v96_, globalState)))), d_105_v53_.f21)))
                lhs14_ = d_105_v53_
                lhs15_ = d_147_v88_
                lhs16_ = default__.safeIndex(253, (d_147_v88_).length(0))
                d_158_v92_ = rhs18_
                d_159_v93_ = rhs19_
                lhs14_.f21 = rhs20_
                lhs15_[lhs16_] = rhs21_
                d_170_v104_: D0
                d_170_v104_ = D0_DC0(d_156_cf40_)
                d_171_v105_: C7
                nw25_ = C7()
                nw25_.ctor__((d_105_v53_).f22, d_104_v52_, d_170_v104_, (d_105_v53_).f22)
                d_171_v105_ = nw25_
                d_172_v106_: _dafny.Seq
                d_172_v106_ = _dafny.SeqWithoutIsStrInference([d_171_v105_])
                d_173_v107_: _dafny.Array
                nw26_ = _dafny.Array(None, 27)
                nw26_[int(0)] = (d_171_v105_ if d_105_v53_.f21 else d_171_v105_)
                nw26_[int(1)] = d_171_v105_
                nw26_[int(2)] = d_171_v105_
                nw26_[int(3)] = d_171_v105_
                nw26_[int(4)] = d_171_v105_
                nw26_[int(5)] = d_171_v105_
                nw26_[int(6)] = d_171_v105_
                nw26_[int(7)] = d_171_v105_
                nw26_[int(8)] = d_171_v105_
                nw26_[int(9)] = d_171_v105_
                nw26_[int(10)] = d_171_v105_
                nw26_[int(11)] = d_171_v105_
                nw26_[int(12)] = d_171_v105_
                nw26_[int(13)] = d_171_v105_
                nw26_[int(14)] = d_171_v105_
                nw26_[int(15)] = d_171_v105_
                nw26_[int(16)] = d_171_v105_
                nw26_[int(17)] = d_171_v105_
                nw26_[int(18)] = d_171_v105_
                nw26_[int(19)] = d_171_v105_
                nw26_[int(20)] = d_171_v105_
                nw26_[int(21)] = d_171_v105_
                nw26_[int(22)] = (d_172_v106_)[default__.safeIndex(d_106_v54_, len(d_172_v106_))]
                nw26_[int(23)] = d_171_v105_
                nw26_[int(24)] = d_171_v105_
                nw26_[int(25)] = d_171_v105_
                nw26_[int(26)] = d_171_v105_
                d_173_v107_ = nw26_
                index15_ = default__.safeIndex(922, (d_173_v107_).length(0))
                (d_173_v107_)[index15_] = d_171_v105_
                index16_ = default__.safeIndex(922, (d_173_v107_).length(0))
                (d_173_v107_)[index16_] = d_171_v105_
                (globalState).f9 = -780
            elif True:
                d_174___mcc_h2_ = source0_.cf39
                d_175_cf39_ = d_174___mcc_h2_
                (d_105_v53_).f21 = default__.fm1((d_105_v53_).f22, _dafny.CodePoint('q'), globalState)
                d_176_v108_: str
                d_176_v108_ = _dafny.CodePoint('e')
                d_177_v109_: _dafny.Seq
                d_177_v109_ = _dafny.SeqWithoutIsStrInference([d_176_v108_, d_176_v108_, default__.fm3(d_106_v54_, d_105_v53_.f21, globalState)])
                d_178_v110_: _dafny.Map
                d_178_v110_ = _dafny.Map({False: (0) - (len((d_177_v109_) + (d_177_v109_)))})
                d_179_v111_: _dafny.Seq
                d_179_v111_ = _dafny.SeqWithoutIsStrInference([d_177_v109_, d_177_v109_])
                rhs22_ = (d_179_v111_)[default__.safeIndex(default__.safeModuloInt(591, d_106_v54_), len(d_179_v111_))]
                rhs23_ = (d_178_v110_) | (_dafny.Map({(d_105_v53_).f22: d_106_v54_}))
                rhs24_ = d_106_v54_
                d_177_v109_ = rhs22_
                d_178_v110_ = rhs23_
                d_106_v54_ = rhs24_
                (globalState).f9 = d_106_v54_
                d_180_v112_: T0
                nw27_ = C3()
                nw27_.ctor__(not(d_45_v0_), (d_105_v53_).f22, d_104_v52_)
                d_180_v112_ = nw27_
                d_181_v113_: _dafny.Map
                d_181_v113_ = _dafny.Map({d_105_v53_.f21: d_180_v112_})
                d_181_v113_ = (d_181_v113_).set(d_45_v0_, d_180_v112_)
        elif True:
            d_182_v114_: _dafny.Array
            nw28_ = _dafny.Array(None, 21)
            d_182_v114_ = nw28_
            d_183_v115_: C4
            nw29_ = C4()
            nw29_.ctor__(d_104_v52_)
            d_183_v115_ = nw29_
            index17_ = default__.safeIndex(840, (d_182_v114_).length(0))
            (d_182_v114_)[index17_] = d_183_v115_
            index18_ = default__.safeIndex(840, (d_182_v114_).length(0))
            (d_182_v114_)[index18_] = d_183_v115_
            d_106_v54_ = d_106_v54_
            d_184_v116_: C2
            nw30_ = C2()
            nw30_.ctor__(d_104_v52_)
            d_184_v116_ = nw30_
            d_185_v117_: D8
            d_185_v117_ = D8_DC19(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_186_i11_ in range(default__.abs(647))]), d_106_v54_)
            d_187_v118_: _dafny.Seq
            d_187_v118_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jp"))
            d_188_v119_: str
            d_188_v119_ = _dafny.CodePoint('j')
            d_185_v117_ = D8_DC19(((d_187_v118_).set(default__.safeIndex(762, len(d_187_v118_)), d_188_v119_) if d_105_v53_.f21 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bpmojegsf"))), default__.safeDivisionInt(d_106_v54_, d_106_v54_))
            d_189_v120_: C8
            nw31_ = C8()
            nw31_.ctor__(not(d_105_v53_.f21))
            d_189_v120_ = nw31_
        d_190_v121_: D8
        d_190_v121_ = D8_DC18(d_144_v86_)
        d_191_v122_: _dafny.Seq
        d_191_v122_ = _dafny.SeqWithoutIsStrInference([d_144_v86_, d_144_v86_, d_144_v86_, d_144_v86_])
        pat_let_tv1_ = d_191_v122_
        pat_let_tv2_ = d_106_v54_
        pat_let_tv3_ = d_191_v122_
        def iife29_(_pat_let2_0):
            def iife30_(d_192_dt__update__tmp_h1_):
                def iife31_(_pat_let3_0):
                    def iife32_(d_193_dt__update_hcf33_h0_):
                        return D8_DC18(d_193_dt__update_hcf33_h0_)
                    return iife32_(_pat_let3_0)
                return iife31_((pat_let_tv1_)[default__.safeIndex(pat_let_tv2_, len(pat_let_tv3_))])
            return iife30_(_pat_let2_0)
        d_190_v121_ = iife29_(d_190_v121_)
        d_194_v123_: _dafny.Array
        nw32_ = _dafny.Array(None, 1)
        nw32_[int(0)] = d_45_v0_
        d_194_v123_ = nw32_
        index19_ = default__.safeIndex(991, (d_194_v123_).length(0))
        (d_194_v123_)[index19_] = d_45_v0_
        index20_ = default__.safeIndex(991, (d_194_v123_).length(0))
        (d_194_v123_)[index20_] = not(d_105_v53_.f21)

    @staticmethod
    def Main(noArgsParameter__):
        d_195_v0_: _dafny.Array
        def lambda9_(d_196_i0_):
            return (d_196_i0_) * (-343)

        init5_ = lambda9_
        nw33_ = _dafny.Array(None, 17)
        for i0_5_ in range(nw33_.length(0)):
            nw33_[i0_5_] = init5_(i0_5_)
        d_195_v0_ = nw33_
        d_197_v1_: _dafny.Array
        nw34_ = _dafny.Array(None, 1)
        nw34_[int(0)] = d_195_v0_
        d_197_v1_ = nw34_
        d_198_globalState_: GlobalState
        nw35_ = GlobalState()
        nw35_.ctor__(False, d_195_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "k")), d_197_v1_, -63, 660, 763, 917, False, 720)
        d_198_globalState_ = nw35_
        d_199_v2_: int
        d_199_v2_ = -797
        hi1_ = 901
        for d_200_i1_ in range(d_199_v2_, hi1_):
            d_201_v3_: bool
            d_201_v3_ = False
            d_202_v4_: _dafny.Set
            d_202_v4_ = _dafny.Set({not(d_201_v3_)})
            d_203_v5_: _dafny.Seq
            d_203_v5_ = _dafny.SeqWithoutIsStrInference([d_202_v4_, d_202_v4_, d_202_v4_])
            d_204_v6_: _dafny.Map
            d_204_v6_ = _dafny.Map({d_201_v3_: len(_dafny.Map({d_201_v3_: len((d_203_v5_)[default__.safeIndex(default__.fm0(d_201_v3_, d_198_globalState_), len(d_203_v5_))])}))})
            d_205_v7_: str
            d_205_v7_ = _dafny.CodePoint('s')
            d_206_v8_: _dafny.MultiSet
            d_206_v8_ = _dafny.MultiSet([d_205_v7_, d_205_v7_, d_205_v7_])
            (d_198_globalState_).f9 = ((d_204_v6_)[(d_206_v8_).ispropersubset(d_206_v8_)] if ((d_206_v8_).ispropersubset(d_206_v8_)) in (d_204_v6_) else d_199_v2_)
            d_207_v9_: _dafny.Set
            d_207_v9_ = _dafny.Set({d_200_i1_})
            (d_198_globalState_).f8 = ((d_207_v9_) - (d_207_v9_)).ispropersubset(d_207_v9_)
            d_208_v10_: _dafny.MultiSet
            d_208_v10_ = _dafny.MultiSet([d_201_v3_])
            d_208_v10_ = ((d_208_v10_) | (d_208_v10_)).set(False, default__.abs((d_200_i1_) + (d_200_i1_)))
            (d_198_globalState_).f9 = (d_199_v2_) * ((0) - (d_199_v2_))
        default__.m0(d_198_globalState_)
        d_209_v11_: bool
        d_209_v11_ = True
        d_210_i2_: int
        d_210_i2_ = 0
        with _dafny.label("1"):
            while d_209_v11_:
                with _dafny.c_label("1"):
                    if (d_210_i2_) >= (100):
                        raise _dafny.Break("1")
                    d_210_i2_ = (d_210_i2_) + (1)
                    if d_209_v11_:
                        d_199_v2_ = (d_199_v2_) + (d_199_v2_)
                        d_211_v12_: _dafny.Seq
                        d_211_v12_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wwabxmyaf"))
                        d_212_v13_: _dafny.Map
                        d_212_v13_ = _dafny.Map({d_211_v12_: d_199_v2_})
                        index21_ = default__.safeIndex(606, (d_195_v0_).length(0))
                        (d_195_v0_)[index21_] = len(d_212_v13_)
                        index22_ = default__.safeIndex(606, (d_195_v0_).length(0))
                        (d_195_v0_)[index22_] = (d_199_v2_) * (d_199_v2_)
                        d_213_v14_: _dafny.MultiSet
                        d_213_v14_ = _dafny.MultiSet([(d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))]])
                        d_214_v15_: str
                        d_214_v15_ = _dafny.CodePoint('r')
                        d_215_v16_: _dafny.Seq
                        d_215_v16_ = _dafny.SeqWithoutIsStrInference([d_209_v11_, d_209_v11_, default__.fm1(d_209_v11_, d_214_v15_, d_198_globalState_), d_209_v11_])
                        (d_198_globalState_).f9 = (0) - ((((d_213_v14_)[d_199_v2_] if (d_199_v2_) in (d_213_v14_) else (d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))])) + ((len(d_215_v16_)) - (d_199_v2_)))
                        d_216_v17_: _dafny.Array
                        nw36_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 14)
                        d_216_v17_ = nw36_
                        index23_ = default__.safeIndex(856, (d_216_v17_).length(0))
                        (d_216_v17_)[index23_] = d_211_v12_
                        d_217_v18_: _dafny.Array
                        nw37_ = _dafny.Array(_dafny.Array(None, 0), 4)
                        d_217_v18_ = nw37_
                        d_218_v19_: _dafny.Array
                        nw38_ = _dafny.Array(_dafny.Set({}), 27)
                        d_218_v19_ = nw38_
                        index24_ = default__.safeIndex(127, (d_217_v18_).length(0))
                        (d_217_v18_)[index24_] = d_218_v19_
                        index25_ = default__.safeIndex(856, (d_216_v17_).length(0))
                        index26_ = default__.safeIndex(127, (d_217_v18_).length(0))
                        rhs25_ = len((default__.fm2((d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))], _dafny.MultiSet([(d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))]]), _dafny.Map({d_199_v2_: d_199_v2_}), d_198_globalState_)) + ((d_211_v12_).set(default__.safeIndex((d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))], len(d_211_v12_)), d_214_v15_)))
                        rhs26_ = d_211_v12_
                        rhs27_ = d_209_v11_
                        rhs28_ = d_218_v19_
                        rhs29_ = d_199_v2_
                        lhs17_ = d_198_globalState_
                        lhs18_ = d_216_v17_
                        lhs19_ = default__.safeIndex(856, (d_216_v17_).length(0))
                        lhs20_ = d_198_globalState_
                        lhs21_ = d_217_v18_
                        lhs22_ = default__.safeIndex(127, (d_217_v18_).length(0))
                        lhs23_ = d_198_globalState_
                        lhs17_.f9 = rhs25_
                        lhs18_[lhs19_] = rhs26_
                        lhs20_.f8 = rhs27_
                        lhs21_[lhs22_] = rhs28_
                        lhs23_.f9 = rhs29_
                        d_209_v11_ = ((d_199_v2_ if d_209_v11_ else d_199_v2_)) == (default__.safeDivisionInt((d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))], (d_195_v0_)[default__.safeIndex(606, (d_195_v0_).length(0))]))
                    elif True:
                        (d_198_globalState_).f8 = d_209_v11_
                        (d_198_globalState_).f8 = d_209_v11_
                        (d_198_globalState_).f9 = default__.fm0(True, d_198_globalState_)
                        d_219_v20_: _dafny.Array
                        nw39_ = _dafny.Array(False, 26)
                        d_219_v20_ = nw39_
                        d_220_v21_: str
                        d_220_v21_ = _dafny.CodePoint('f')
                        d_221_v22_: _dafny.Seq
                        d_221_v22_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jgl"))])
                        d_222_v23_: _dafny.MultiSet
                        d_222_v23_ = _dafny.MultiSet([d_199_v2_])
                        d_223_v24_: _dafny.Seq
                        d_223_v24_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bwtvm"))
                        d_224_v25_: _dafny.Map
                        d_224_v25_ = _dafny.Map({d_199_v2_: d_219_v20_})
                        rhs30_ = len((d_221_v22_) + ((_dafny.SeqWithoutIsStrInference([default__.fm2(d_199_v2_, d_222_v23_, _dafny.Map({d_199_v2_: d_199_v2_}), d_198_globalState_)])) + (_dafny.SeqWithoutIsStrInference([d_223_v24_]))))
                        rhs31_ = ((d_224_v25_)[(0) - (d_199_v2_)] if ((0) - (d_199_v2_)) in (d_224_v25_) else d_219_v20_)
                        rhs32_ = d_220_v21_
                        lhs24_ = d_198_globalState_
                        lhs24_.f9 = rhs30_
                        d_219_v20_ = rhs31_
                        d_220_v21_ = rhs32_
                        d_225_v26_: _dafny.MultiSet
                        d_225_v26_ = _dafny.MultiSet([d_209_v11_, d_209_v11_])
                        d_226_v27_: D0
                        d_226_v27_ = D0_DC2(d_225_v26_, True, d_199_v2_, d_209_v11_)
                        rhs33_ = (d_199_v2_) - (d_199_v2_)
                        rhs34_ = ((d_226_v27_).cf6) != (d_199_v2_)
                        rhs35_ = ((0) - (default__.fm0(d_209_v11_, d_198_globalState_)) if (d_223_v24_) == (d_223_v24_) else d_199_v2_)
                        lhs25_ = d_198_globalState_
                        lhs26_ = d_198_globalState_
                        lhs27_ = d_198_globalState_
                        lhs25_.f9 = rhs33_
                        lhs26_.f8 = rhs34_
                        lhs27_.f9 = rhs35_
                    (d_198_globalState_).f8 = True
                    d_227_v28_: _dafny.Map
                    d_227_v28_ = _dafny.Map({d_209_v11_: not(d_209_v11_)})
                    d_228_v29_: _dafny.Map
                    d_228_v29_ = _dafny.Map({((d_227_v28_)[d_209_v11_] if (d_209_v11_) in (d_227_v28_) else d_209_v11_): d_209_v11_})
                    rhs36_ = d_209_v11_
                    rhs37_ = ((d_228_v29_ if d_209_v11_ else d_227_v28_)) | ((d_227_v28_) | (d_228_v29_))
                    rhs38_ = (d_199_v2_) == (d_199_v2_)
                    lhs28_ = d_198_globalState_
                    lhs28_.f8 = rhs36_
                    d_228_v29_ = rhs37_
                    d_209_v11_ = rhs38_
                    default__.m0(d_198_globalState_)
                    pass
            pass
        default__.m0(d_198_globalState_)
        if d_209_v11_:
            if d_209_v11_:
                d_229_v30_: _dafny.Array
                nw40_ = _dafny.Array(False, 14)
                d_229_v30_ = nw40_
                index27_ = default__.safeIndex(35, (d_229_v30_).length(0))
                (d_229_v30_)[index27_] = d_209_v11_
                d_230_v31_: _dafny.Map
                d_230_v31_ = _dafny.Map({d_209_v11_: d_199_v2_})
                d_231_v32_: _dafny.Map
                d_231_v32_ = _dafny.Map({((d_230_v31_)[d_209_v11_] if (d_209_v11_) in (d_230_v31_) else d_199_v2_): not(d_209_v11_)})
                d_232_v33_: _dafny.Map
                d_232_v33_ = _dafny.Map({d_209_v11_: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_233_i3_ in range(default__.abs(655))])})
                index28_ = default__.safeIndex(35, (d_229_v30_).length(0))
                (d_229_v30_)[index28_] = ((d_231_v32_)[(len(d_232_v33_)) * ((_dafny.MultiSet([d_209_v11_])).cardinality)] if ((len(d_232_v33_)) * ((_dafny.MultiSet([d_209_v11_])).cardinality)) in (d_231_v32_) else not(d_209_v11_))
                d_234_v34_: str
                d_234_v34_ = _dafny.CodePoint('q')
                (d_198_globalState_).f8 = default__.fm1((d_229_v30_)[default__.safeIndex(35, (d_229_v30_).length(0))], d_234_v34_, d_198_globalState_)
                d_230_v31_ = d_230_v31_
                d_234_v34_ = default__.fm3(len(_dafny.SeqWithoutIsStrInference([d_209_v11_])), (d_229_v30_)[default__.safeIndex(35, (d_229_v30_).length(0))], d_198_globalState_)
                d_229_v30_ = d_229_v30_
            elif True:
                d_235_v35_: _dafny.Seq
                d_235_v35_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nliacvc"))
                d_235_v35_ = d_235_v35_
                d_236_v36_: _dafny.Array
                nw41_ = _dafny.Array(False, 17)
                d_236_v36_ = nw41_
                index29_ = default__.safeIndex(728, (d_236_v36_).length(0))
                (d_236_v36_)[index29_] = d_209_v11_
                index30_ = default__.safeIndex(728, (d_236_v36_).length(0))
                (d_236_v36_)[index30_] = (default__.fm0(d_209_v11_, d_198_globalState_)) >= ((d_199_v2_) + (d_199_v2_))
                d_237_v37_: _dafny.MultiSet
                d_237_v37_ = _dafny.MultiSet([d_199_v2_, d_199_v2_, (0) - (d_199_v2_)])
                d_237_v37_ = d_237_v37_
                d_238_v38_: _dafny.MultiSet
                d_238_v38_ = _dafny.MultiSet([d_209_v11_, d_209_v11_])
                d_239_v39_: str
                d_239_v39_ = _dafny.CodePoint('e')
                d_240_v40_: D0
                d_240_v40_ = D0_DC2((d_238_v38_) | (d_238_v38_), (d_199_v2_) > (d_199_v2_), d_199_v2_, default__.fm1((d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))], d_239_v39_, d_198_globalState_))
                d_241_v41_: _dafny.Seq
                d_241_v41_ = _dafny.SeqWithoutIsStrInference([(d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))], (d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]])
                d_242_v42_: _dafny.Array
                nw42_ = _dafny.Array(None, 21)
                nw42_[int(0)] = d_241_v41_
                nw42_[int(1)] = d_241_v41_
                nw42_[int(2)] = d_241_v41_
                nw42_[int(3)] = _dafny.SeqWithoutIsStrInference([(d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]])
                nw42_[int(4)] = (d_241_v41_) + (d_241_v41_)
                nw42_[int(5)] = d_241_v41_
                nw42_[int(6)] = (d_241_v41_) + ((d_241_v41_).set(default__.safeIndex(d_199_v2_, len(d_241_v41_)), (d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]))
                nw42_[int(7)] = d_241_v41_
                nw42_[int(8)] = d_241_v41_
                nw42_[int(9)] = d_241_v41_
                nw42_[int(10)] = _dafny.SeqWithoutIsStrInference([(d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))], True])
                nw42_[int(11)] = d_241_v41_
                nw42_[int(12)] = d_241_v41_
                nw42_[int(13)] = (d_241_v41_) + (d_241_v41_)
                nw42_[int(14)] = d_241_v41_
                nw42_[int(15)] = (((d_241_v41_).set(default__.safeIndex(len(d_241_v41_), len(d_241_v41_)), d_209_v11_) if True else d_241_v41_)).set(default__.safeIndex(d_199_v2_, len(((d_241_v41_).set(default__.safeIndex(len(d_241_v41_), len(d_241_v41_)), d_209_v11_) if True else d_241_v41_))), False)
                nw42_[int(16)] = d_241_v41_
                nw42_[int(17)] = d_241_v41_
                nw42_[int(18)] = d_241_v41_
                nw42_[int(19)] = (d_241_v41_) + (d_241_v41_)
                nw42_[int(20)] = d_241_v41_
                d_242_v42_ = nw42_
                index31_ = default__.safeIndex(54, (d_242_v42_).length(0))
                (d_242_v42_)[index31_] = ((_dafny.SeqWithoutIsStrInference([(d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))], (d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]])).set(default__.safeIndex(d_199_v2_, len(_dafny.SeqWithoutIsStrInference([(d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))], (d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]]))), d_209_v11_)) + (d_241_v41_)
                d_243_v43_: _dafny.Set
                d_243_v43_ = _dafny.Set({(d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))], not((d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]), d_209_v11_, True, (d_236_v36_)[default__.safeIndex(728, (d_236_v36_).length(0))]})
                d_244_v44_: _dafny.Seq
                d_244_v44_ = _dafny.SeqWithoutIsStrInference([(0) - (d_199_v2_)])
                d_245_v45_: _dafny.Map
                d_245_v45_ = _dafny.Map({d_235_v35_: d_199_v2_})
                index32_ = default__.safeIndex(728, (d_236_v36_).length(0))
                index33_ = default__.safeIndex(54, (d_242_v42_).length(0))
                rhs39_ = default__.fm4(len(d_243_v43_), d_199_v2_, d_198_globalState_)
                rhs40_ = d_236_v36_
                rhs41_ = not((len(d_244_v44_)) <= ((d_199_v2_) + (len(d_245_v45_))))
                rhs42_ = d_241_v41_
                rhs43_ = default__.fm3(((0) - (d_199_v2_)) * (d_199_v2_), False, d_198_globalState_)
                lhs29_ = d_236_v36_
                lhs30_ = default__.safeIndex(728, (d_236_v36_).length(0))
                lhs31_ = d_242_v42_
                lhs32_ = default__.safeIndex(54, (d_242_v42_).length(0))
                d_240_v40_ = rhs39_
                d_236_v36_ = rhs40_
                lhs29_[lhs30_] = rhs41_
                lhs31_[lhs32_] = rhs42_
                d_239_v39_ = rhs43_
                default__.m0(d_198_globalState_)
            d_246_v46_: _dafny.Seq
            d_246_v46_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyv"))
            d_247_v47_: str
            d_247_v47_ = _dafny.CodePoint('k')
            d_246_v46_ = ((d_246_v46_) + (d_246_v46_)).set(default__.safeIndex((0) - (default__.safeDivisionInt(d_199_v2_, d_199_v2_)), len((d_246_v46_) + (d_246_v46_))), d_247_v47_)
            d_248_v48_: D5
            d_248_v48_ = D5_DC14(d_209_v11_)
            d_249_v49_: _dafny.Array
            def lambda10_(d_250_v2_, d_251_v11_):
                def lambda11_(d_252_i4_):
                    return _dafny.Map({d_250_v2_: d_251_v11_})

                return lambda11_

            init6_ = lambda10_(d_199_v2_, d_209_v11_)
            nw43_ = _dafny.Array(None, 2)
            for i0_6_ in range(nw43_.length(0)):
                nw43_[i0_6_] = init6_(i0_6_)
            d_249_v49_ = nw43_
            d_253_v50_: C3
            nw44_ = C3()
            nw44_.ctor__((d_248_v48_).cf26, d_209_v11_, (d_249_v49_ if d_209_v11_ else d_249_v49_))
            d_253_v50_ = nw44_
            d_254_v51_: D0
            d_254_v51_ = D0_DC0(d_199_v2_)
            d_255_v52_: C1
            nw45_ = C1()
            nw45_.ctor__(d_253_v50_.f21, d_253_v50_.f21, d_254_v51_, d_209_v11_)
            d_255_v52_ = nw45_
            d_256_v53_: _dafny.Seq
            d_256_v53_ = _dafny.SeqWithoutIsStrInference([d_255_v52_, d_255_v52_])
            d_257_v54_: _dafny.MultiSet
            d_257_v54_ = _dafny.MultiSet([d_256_v53_])
            d_258_v55_: _dafny.Map
            d_258_v55_ = _dafny.Map({d_195_v0_: d_256_v53_})
            d_259_v56_: bool
            d_260_v57_: _dafny.Array
            out0_: bool
            out1_: _dafny.Array
            out0_, out1_ = (d_253_v50_).m13(d_199_v2_, (d_253_v50_).f22, 641, (_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([((d_258_v55_)[d_195_v0_] if (d_195_v0_) in (d_258_v55_) else d_256_v53_)]))).issubset(d_257_v54_), d_198_globalState_)
            d_259_v56_ = out0_
            d_260_v57_ = out1_
            if (d_253_v50_).f22:
                index34_ = default__.safeIndex(190, (d_195_v0_).length(0))
                (d_195_v0_)[index34_] = d_199_v2_
                index35_ = default__.safeIndex(190, (d_195_v0_).length(0))
                (d_195_v0_)[index35_] = d_199_v2_
                d_261_v58_: T1
                nw46_ = C2()
                nw46_.ctor__(d_249_v49_)
                d_261_v58_ = nw46_
                d_261_v58_ = d_261_v58_
                d_262_v59_: _dafny.Array
                nw47_ = _dafny.Array(False, 5)
                d_262_v59_ = nw47_
                d_263_v60_: _dafny.Seq
                d_263_v60_ = _dafny.SeqWithoutIsStrInference([len(d_246_v46_)])
                d_264_v61_: _dafny.Map
                d_264_v61_ = _dafny.Map({d_199_v2_: d_253_v50_.f21})
                d_265_v62_: _dafny.Seq
                d_265_v62_ = _dafny.SeqWithoutIsStrInference([((d_195_v0_)[default__.safeIndex(190, (d_195_v0_).length(0))]) + (d_199_v2_), d_199_v2_, len(d_263_v60_), (d_255_v52_).fm30(d_264_v61_, d_198_globalState_), d_199_v2_])
                d_266_v63_: _dafny.Seq
                d_266_v63_ = _dafny.SeqWithoutIsStrInference([(d_253_v50_).f22])
                rhs44_ = d_199_v2_
                rhs45_ = (d_263_v60_)[default__.safeIndex(len(d_266_v63_), len(d_263_v60_))]
                rhs46_ = default__.safeDivisionInt(d_199_v2_, (0) - ((d_199_v2_) + (d_199_v2_)))
                rhs47_ = d_262_v59_
                rhs48_ = (((d_195_v0_)[default__.safeIndex(190, (d_195_v0_).length(0))] if (d_253_v50_).f22 else (d_195_v0_)[default__.safeIndex(190, (d_195_v0_).length(0))])) + (d_199_v2_)
                lhs33_ = d_198_globalState_
                lhs34_ = d_198_globalState_
                lhs35_ = d_198_globalState_
                lhs33_.f9 = rhs44_
                lhs34_.f9 = rhs45_
                d_199_v2_ = rhs46_
                d_262_v59_ = rhs47_
                lhs35_.f9 = rhs48_
                d_267_v64_: D2
                d_267_v64_ = D2_DC6()
                d_268_v65_: _dafny.Map
                d_268_v65_ = _dafny.Map({d_267_v64_: True})
                d_269_v66_: _dafny.Map
                d_269_v66_ = _dafny.Map({True: (d_255_v52_).f20})
                d_270_v67_: _dafny.Map
                d_270_v67_ = _dafny.Map({d_269_v66_: d_199_v2_})
                d_268_v65_ = (d_268_v65_).set(d_267_v64_, (977) >= (((d_270_v67_)[d_269_v66_] if (d_269_v66_) in (d_270_v67_) else len(d_266_v63_))))
                d_271_v68_: _dafny.Array
                def lambda12_(d_272_v2_, d_273_v0_):
                    def lambda13_(d_274_i5_):
                        return (_dafny.MultiSet([d_272_v2_])).intersection(_dafny.MultiSet([(d_273_v0_)[default__.safeIndex(190, (d_273_v0_).length(0))], (d_273_v0_)[default__.safeIndex(190, (d_273_v0_).length(0))], (d_273_v0_)[default__.safeIndex(190, (d_273_v0_).length(0))]]))

                    return lambda13_

                init7_ = lambda12_(d_199_v2_, d_195_v0_)
                nw48_ = _dafny.Array(None, 27)
                for i0_7_ in range(nw48_.length(0)):
                    nw48_[i0_7_] = init7_(i0_7_)
                d_271_v68_ = nw48_
                d_275_v69_: _dafny.MultiSet
                d_275_v69_ = _dafny.MultiSet([(d_195_v0_)[default__.safeIndex(190, (d_195_v0_).length(0))]])
                index36_ = default__.safeIndex(748, (d_271_v68_).length(0))
                (d_271_v68_)[index36_] = (d_275_v69_) - (d_275_v69_)
                d_276_v70_: C6
                nw49_ = C6()
                nw49_.ctor__(d_199_v2_, d_246_v46_)
                d_276_v70_ = nw49_
                d_277_v71_: _dafny.Map
                d_277_v71_ = _dafny.Map({d_276_v70_: (d_255_v52_).f19})
                d_278_v72_: _dafny.Map
                d_278_v72_ = _dafny.Map({(d_276_v70_).f15: len(d_276_v70_.f16)})
                index37_ = default__.safeIndex(748, (d_271_v68_).length(0))
                (d_271_v68_)[index37_] = (d_275_v69_) | ((d_275_v69_).intersection(((d_275_v69_).set(len(d_277_v71_), default__.abs(571))).set(len(d_278_v72_), default__.abs((d_195_v0_)[default__.safeIndex(190, (d_195_v0_).length(0))]))))
            elif True:
                (d_198_globalState_).f9 = d_199_v2_
                d_279_v73_: C5
                nw50_ = C5()
                nw50_.ctor__(d_249_v49_)
                d_279_v73_ = nw50_
                d_280_v74_: _dafny.Set
                d_280_v74_ = _dafny.Set({d_279_v73_, d_279_v73_})
                rhs49_ = ((d_280_v74_).intersection(d_280_v74_)).ispropersubset((_dafny.Set({d_279_v73_, d_279_v73_, d_279_v73_, d_279_v73_}) if d_209_v11_ else d_280_v74_))
                rhs50_ = (d_253_v50_).f22
                d_259_v56_ = rhs49_
                d_209_v11_ = rhs50_
                (d_198_globalState_).f9 = (d_255_v52_).fm30(_dafny.Map({d_199_v2_: (d_255_v52_).f20}), d_198_globalState_)
                d_281_v75_: _dafny.Map
                d_281_v75_ = _dafny.Map({d_199_v2_: _dafny.Map({(d_255_v52_).f19: d_199_v2_})})
                d_282_v76_: _dafny.Map
                d_282_v76_ = _dafny.Map({d_199_v2_: d_253_v50_.f21})
                d_283_v77_: _dafny.Map
                d_283_v77_ = _dafny.Map({(d_255_v52_).f19: (d_255_v52_).fm30(d_282_v76_, d_198_globalState_)})
                d_284_v78_: bool
                out2_: bool
                out2_ = (d_253_v50_).m12(d_259_v56_, d_259_v56_, ((d_281_v75_)[d_199_v2_] if (d_199_v2_) in (d_281_v75_) else d_283_v77_), d_198_globalState_)
                d_284_v78_ = out2_
                d_285_v79_: C0
                nw51_ = C0()
                nw51_.ctor__(d_284_v78_)
                d_285_v79_ = nw51_
                d_286_v80_: _dafny.Array
                nw52_ = _dafny.Array(_dafny.Map({}), 27)
                d_286_v80_ = nw52_
                d_287_v81_: T1
                nw53_ = C5()
                nw53_.ctor__(d_249_v49_)
                d_287_v81_ = nw53_
                d_288_v82_: _dafny.Map
                d_288_v82_ = _dafny.Map({d_287_v81_: (d_253_v50_).f22})
                index38_ = default__.safeIndex(910, (d_286_v80_).length(0))
                (d_286_v80_)[index38_] = d_288_v82_
                d_289_v83_: D4
                d_289_v83_ = D4_DC11(d_199_v2_, d_199_v2_, (d_253_v50_).f22)
                pat_let_tv4_ = d_199_v2_
                index39_ = default__.safeIndex(910, (d_286_v80_).length(0))
                def iife33_(_pat_let4_0):
                    def iife34_(d_290_dt__update__tmp_h0_):
                        def iife35_(_pat_let5_0):
                            def iife36_(d_291_dt__update_hcf19_h0_):
                                return D4_DC11(d_291_dt__update_hcf19_h0_, (d_290_dt__update__tmp_h0_).cf20, (d_290_dt__update__tmp_h0_).cf21)
                            return iife36_(_pat_let5_0)
                        return iife35_(pat_let_tv4_)
                    return iife34_(_pat_let4_0)
                rhs51_ = d_285_v79_
                rhs52_ = default__.safeDivisionInt(default__.safeModuloInt(d_199_v2_, d_199_v2_), d_199_v2_)
                rhs53_ = _dafny.Map({d_287_v81_: not(d_259_v56_)})
                rhs54_ = (d_199_v2_) * ((d_199_v2_ if (d_253_v50_).f22 else d_199_v2_))
                rhs55_ = (iife33_(d_289_v83_)).cf21
                lhs36_ = d_286_v80_
                lhs37_ = default__.safeIndex(910, (d_286_v80_).length(0))
                lhs38_ = d_198_globalState_
                lhs39_ = d_253_v50_
                d_285_v79_ = rhs51_
                d_199_v2_ = rhs52_
                lhs36_[lhs37_] = rhs53_
                lhs38_.f9 = rhs54_
                lhs39_.f21 = rhs55_
        elif True:
            d_292_v84_: _dafny.Seq
            d_292_v84_ = _dafny.SeqWithoutIsStrInference([d_209_v11_, False])
            d_293_v85_: str
            d_293_v85_ = _dafny.CodePoint('y')
            if (d_292_v84_) == (default__.fm44(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_209_v11_])]), d_293_v85_, d_198_globalState_)):
                (d_198_globalState_).f8 = default__.fm1(d_209_v11_, d_293_v85_, d_198_globalState_)
                d_294_v86_: _dafny.Map
                d_294_v86_ = _dafny.Map({d_199_v2_: d_209_v11_})
                d_295_v87_: _dafny.MultiSet
                d_295_v87_ = _dafny.MultiSet([d_209_v11_])
                d_296_v88_: _dafny.Seq
                d_296_v88_ = _dafny.SeqWithoutIsStrInference([d_199_v2_, len(d_294_v86_), (d_295_v87_).cardinality])
                d_297_v89_: _dafny.Seq
                d_297_v89_ = _dafny.SeqWithoutIsStrInference([d_296_v88_])
                d_199_v2_ = (len(((_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_199_v2_, d_199_v2_, d_199_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bojpia")))]) for d_298_i6_ in range(default__.abs(632))])).set(default__.safeIndex(d_199_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_199_v2_, d_199_v2_, d_199_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bojpia")))]) for d_299_i6_ in range(default__.abs(632))]))), d_296_v88_)) + (d_297_v89_)) if d_209_v11_ else d_199_v2_)
                d_300_v90_: _dafny.Seq
                d_300_v90_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nq"))
                (d_198_globalState_).f8 = default__.fm1(d_209_v11_, (d_300_v90_)[default__.safeIndex((0) - (d_199_v2_), len(d_300_v90_))], d_198_globalState_)
                d_301_v91_: _dafny.Map
                d_301_v91_ = _dafny.Map({d_199_v2_: d_294_v86_})
                def iife37_():
                    coll25_ = _dafny.Map()
                    compr_25_: int
                    for compr_25_ in _dafny.IntegerRange(943, 750):
                        d_302_v92_: int = compr_25_
                        if ((943) <= (d_302_v92_)) and ((d_302_v92_) < (750)):
                            coll25_[(d_302_v92_) * (d_199_v2_)] = (d_292_v84_)[default__.safeIndex(-98, len(d_292_v84_))]
                    return _dafny.Map(coll25_)
                d_301_v91_ = (d_301_v91_).set((len(d_292_v84_)) + (d_199_v2_), iife37_()
                )
                d_303_v93_: _dafny.Set
                d_303_v93_ = _dafny.Set({d_300_v90_, d_300_v90_, d_300_v90_, d_300_v90_, d_300_v90_})
                d_304_v94_: _dafny.Array
                nw54_ = _dafny.Array(None, 12)
                nw54_[int(0)] = d_303_v93_
                nw54_[int(1)] = d_303_v93_
                nw54_[int(2)] = d_303_v93_
                nw54_[int(3)] = default__.fm63(d_198_globalState_)
                nw54_[int(4)] = d_303_v93_
                nw54_[int(5)] = d_303_v93_
                nw54_[int(6)] = d_303_v93_
                nw54_[int(7)] = d_303_v93_
                nw54_[int(8)] = d_303_v93_
                nw54_[int(9)] = d_303_v93_
                nw54_[int(10)] = d_303_v93_
                nw54_[int(11)] = d_303_v93_
                d_304_v94_ = nw54_
                index40_ = default__.safeIndex(481, (d_304_v94_).length(0))
                (d_304_v94_)[index40_] = d_303_v93_
                index41_ = default__.safeIndex(481, (d_304_v94_).length(0))
                (d_304_v94_)[index41_] = default__.fm63(d_198_globalState_)
            elif True:
                d_305_v95_: _dafny.Array
                nw55_ = _dafny.Array(_dafny.Map({}), 22)
                d_305_v95_ = nw55_
                d_306_v96_: _dafny.Array
                def lambda14_(d_307_v2_, d_308_v11_):
                    def lambda15_(d_309_i7_):
                        return _dafny.Map({len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvf"))), d_307_v2_})): d_308_v11_})

                    return lambda15_

                init8_ = lambda14_(d_199_v2_, d_209_v11_)
                nw56_ = _dafny.Array(None, 6)
                for i0_8_ in range(nw56_.length(0)):
                    nw56_[i0_8_] = init8_(i0_8_)
                d_306_v96_ = nw56_
                d_310_v97_: T1
                nw57_ = C2()
                nw57_.ctor__(d_306_v96_)
                d_310_v97_ = nw57_
                d_311_v98_: _dafny.Map
                d_311_v98_ = _dafny.Map({d_209_v11_: d_310_v97_})
                index42_ = default__.safeIndex(567, (d_305_v95_).length(0))
                (d_305_v95_)[index42_] = (d_311_v98_) | (d_311_v98_)
                index43_ = default__.safeIndex(567, (d_305_v95_).length(0))
                (d_305_v95_)[index43_] = ((_dafny.Map({d_209_v11_: d_310_v97_})) | (d_311_v98_)) | (d_311_v98_)
                d_312_v99_: _dafny.Set
                d_312_v99_ = _dafny.Set({d_209_v11_})
                index44_ = default__.safeIndex(469, (d_195_v0_).length(0))
                (d_195_v0_)[index44_] = len(d_312_v99_)
                index45_ = default__.safeIndex(469, (d_195_v0_).length(0))
                (d_195_v0_)[index45_] = default__.safeDivisionInt(d_199_v2_, d_199_v2_)
                d_313_v100_: D0
                d_313_v100_ = D0_DC0((0) - (-905))
                d_314_v101_: C7
                nw58_ = C7()
                nw58_.ctor__(d_209_v11_, (d_310_v97_).f11, d_313_v100_, d_209_v11_)
                d_314_v101_ = nw58_
                d_314_v101_ = d_314_v101_
                (d_198_globalState_).f8 = True
                d_315_v102_: _dafny.Array
                nw59_ = _dafny.Array(int(0), 29)
                d_315_v102_ = nw59_
                d_316_v103_: str
                d_317_v104_: int
                d_318_v105_: D2
                d_319_v106_: bool
                out3_: str
                out4_: int
                out5_: D2
                out6_: bool
                out3_, out4_, out5_, out6_ = (d_314_v101_).m4(d_199_v2_, d_315_v102_, d_198_globalState_)
                d_316_v103_ = out3_
                d_317_v104_ = out4_
                d_318_v105_ = out5_
                d_319_v106_ = out6_
            d_209_v11_ = d_209_v11_
            d_320_v107_: _dafny.Set
            d_320_v107_ = _dafny.Set({d_209_v11_, not(d_209_v11_)})
            if ((d_320_v107_) | (_dafny.Set({d_209_v11_}))).ispropersubset(d_320_v107_):
                d_321_v108_: _dafny.Array
                nw60_ = _dafny.Array(None, 17)
                nw60_[int(0)] = (d_292_v84_)[default__.safeIndex(d_199_v2_, len(d_292_v84_))]
                nw60_[int(1)] = d_209_v11_
                nw60_[int(2)] = d_209_v11_
                nw60_[int(3)] = d_209_v11_
                nw60_[int(4)] = d_209_v11_
                nw60_[int(5)] = d_209_v11_
                nw60_[int(6)] = d_209_v11_
                nw60_[int(7)] = d_209_v11_
                nw60_[int(8)] = True
                nw60_[int(9)] = d_209_v11_
                nw60_[int(10)] = not(True)
                nw60_[int(11)] = d_209_v11_
                nw60_[int(12)] = d_209_v11_
                nw60_[int(13)] = not(False)
                nw60_[int(14)] = d_209_v11_
                nw60_[int(15)] = True
                nw60_[int(16)] = d_209_v11_
                d_321_v108_ = nw60_
                d_322_v109_: D10
                d_322_v109_ = D10_DC24(d_321_v108_, d_195_v0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rsxa")))
                d_195_v0_ = (d_322_v109_).cf44
                (d_198_globalState_).f8 = d_209_v11_
                d_323_v110_: _dafny.Array
                nw61_ = _dafny.Array(None, 19)
                d_323_v110_ = nw61_
                d_324_v112_: _dafny.Array
                def lambda16_(d_325_v2_):
                    def lambda17_(d_326_i8_):
                        def iife38_():
                            coll26_ = _dafny.Map()
                            compr_26_: int
                            for compr_26_ in (_dafny.Map({(0) - (d_325_v2_): 229})).keys.Elements:
                                d_327_v111_: int = compr_26_
                                if (d_327_v111_) in (_dafny.Map({(0) - (d_325_v2_): 229})):
                                    coll26_[default__.safeModuloInt(d_327_v111_, d_325_v2_)] = True
                            return _dafny.Map(coll26_)
                        return iife38_()
                        

                    return lambda17_

                init9_ = lambda16_(d_199_v2_)
                nw62_ = _dafny.Array(None, 22)
                for i0_9_ in range(nw62_.length(0)):
                    nw62_[i0_9_] = init9_(i0_9_)
                d_324_v112_ = nw62_
                d_328_v113_: C3
                nw63_ = C3()
                nw63_.ctor__(not(d_209_v11_), d_209_v11_, d_324_v112_)
                d_328_v113_ = nw63_
                index46_ = default__.safeIndex(943, (d_323_v110_).length(0))
                (d_323_v110_)[index46_] = d_328_v113_
                d_329_v114_: C3
                d_329_v114_ = d_328_v113_
                index47_ = default__.safeIndex(943, (d_323_v110_).length(0))
                (d_323_v110_)[index47_] = (d_329_v114_)
                d_330_v115_: _dafny.Seq
                d_330_v115_ = _dafny.SeqWithoutIsStrInference([d_199_v2_])
                d_331_v116_: D7
                d_331_v116_ = D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvj")), d_321_v108_, (d_328_v113_).f22, len(d_330_v115_))
                d_332_v117_: _dafny.Array
                nw64_ = _dafny.Array(None, 17)
                nw64_[int(0)] = d_321_v108_
                nw64_[int(1)] = d_321_v108_
                nw64_[int(2)] = d_321_v108_
                nw64_[int(3)] = d_321_v108_
                nw64_[int(4)] = d_321_v108_
                nw64_[int(5)] = d_321_v108_
                nw64_[int(6)] = d_321_v108_
                nw64_[int(7)] = d_321_v108_
                nw64_[int(8)] = (d_331_v116_).cf30
                nw64_[int(9)] = (d_321_v108_ if (d_328_v113_).f22 else d_321_v108_)
                nw64_[int(10)] = d_321_v108_
                nw64_[int(11)] = d_321_v108_
                nw64_[int(12)] = d_321_v108_
                nw64_[int(13)] = d_321_v108_
                nw64_[int(14)] = d_321_v108_
                nw64_[int(15)] = d_321_v108_
                nw64_[int(16)] = d_321_v108_
                d_332_v117_ = nw64_
                index48_ = default__.safeIndex(612, (d_332_v117_).length(0))
                (d_332_v117_)[index48_] = d_321_v108_
                index49_ = default__.safeIndex(612, (d_332_v117_).length(0))
                (d_332_v117_)[index49_] = d_321_v108_
                (d_198_globalState_).f8 = d_328_v113_.f21
            elif True:
                d_333_v118_: _dafny.Seq
                d_333_v118_ = _dafny.SeqWithoutIsStrInference([d_199_v2_])
                d_333_v118_ = d_333_v118_
                d_334_v119_: _dafny.MultiSet
                d_334_v119_ = _dafny.MultiSet([d_209_v11_, d_209_v11_, not(not(d_209_v11_))])
                d_335_v120_: _dafny.Map
                d_335_v120_ = _dafny.Map({d_334_v119_: (d_209_v11_) and (d_209_v11_)})
                d_335_v120_ = (d_335_v120_).set(d_334_v119_, d_209_v11_)
                d_336_v121_: _dafny.MultiSet
                d_336_v121_ = _dafny.MultiSet([d_199_v2_, d_199_v2_, default__.fm0(d_209_v11_, d_198_globalState_), 79, d_199_v2_])
                d_337_v122_: _dafny.Array
                def lambda18_(d_338_v2_, d_339_v11_):
                    def lambda19_(d_340_i9_):
                        return _dafny.Map({d_338_v2_: d_339_v11_})

                    return lambda19_

                init10_ = lambda18_(d_199_v2_, d_209_v11_)
                nw65_ = _dafny.Array(None, 23)
                for i0_10_ in range(nw65_.length(0)):
                    nw65_[i0_10_] = init10_(i0_10_)
                d_337_v122_ = nw65_
                d_341_v123_: T1
                nw66_ = C5()
                nw66_.ctor__(d_337_v122_)
                d_341_v123_ = nw66_
                d_342_v124_: _dafny.Map
                d_342_v124_ = _dafny.Map({(d_336_v121_) | (d_336_v121_): d_341_v123_})
                d_342_v124_ = (d_342_v124_).set(d_336_v121_, d_341_v123_)
                d_209_v11_ = d_209_v11_
                d_343_v125_: _dafny.Seq
                d_343_v125_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "khkrqxbg"))
                d_344_v126_: _dafny.Map
                d_344_v126_ = _dafny.Map({579: len(_dafny.Map({d_209_v11_: d_199_v2_}))})
                d_345_v127_: _dafny.Array
                nw67_ = _dafny.Array(None, 13)
                nw67_[int(0)] = _dafny.Map({d_199_v2_: (d_334_v119_).cardinality})
                nw67_[int(1)] = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([702, d_199_v2_])): len(d_343_v125_)})
                nw67_[int(2)] = default__.fm7(d_198_globalState_)
                nw67_[int(3)] = (d_344_v126_).set(d_199_v2_, (0) - (d_199_v2_))
                nw67_[int(4)] = default__.fm7(d_198_globalState_)
                nw67_[int(5)] = d_344_v126_
                nw67_[int(6)] = d_344_v126_
                nw67_[int(7)] = d_344_v126_
                nw67_[int(8)] = d_344_v126_
                nw67_[int(9)] = _dafny.Map({d_199_v2_: 204})
                nw67_[int(10)] = d_344_v126_
                nw67_[int(11)] = d_344_v126_
                nw67_[int(12)] = d_344_v126_
                d_345_v127_ = nw67_
                d_346_v128_: _dafny.Seq
                d_346_v128_ = _dafny.SeqWithoutIsStrInference([d_344_v126_])
                index50_ = default__.safeIndex(948, (d_345_v127_).length(0))
                (d_345_v127_)[index50_] = (d_344_v126_) | ((d_346_v128_)[default__.safeIndex((0) - (d_199_v2_), len(d_346_v128_))])
                index51_ = default__.safeIndex(948, (d_345_v127_).length(0))
                (d_345_v127_)[index51_] = (d_344_v126_) | (d_344_v126_)
            d_347_v129_: _dafny.Map
            d_347_v129_ = _dafny.Map({d_209_v11_: (0) - (d_199_v2_)})
            d_348_v130_: _dafny.Seq
            d_348_v130_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cbvy"))
            d_349_v131_: _dafny.Map
            d_349_v131_ = _dafny.Map({len(d_348_v130_): not(d_209_v11_)})
            d_350_v132_: _dafny.Seq
            d_350_v132_ = _dafny.SeqWithoutIsStrInference([((d_347_v129_)[d_209_v11_] if (d_209_v11_) in (d_347_v129_) else len(d_349_v131_))])
            d_350_v132_ = d_350_v132_
            d_199_v2_ = d_199_v2_
        d_351_v133_: _dafny.Map
        d_351_v133_ = _dafny.Map({d_199_v2_: d_209_v11_})
        d_352_v135_: _dafny.MultiSet
        d_352_v135_ = _dafny.MultiSet([d_199_v2_, d_199_v2_])
        d_353_v136_: _dafny.Array
        nw68_ = _dafny.Array(_dafny.Map({}), 10)
        d_353_v136_ = nw68_
        d_354_v137_: T1
        nw69_ = C2()
        nw69_.ctor__(d_353_v136_)
        d_354_v137_ = nw69_
        d_355_v138_: _dafny.Map
        d_355_v138_ = _dafny.Map({d_199_v2_: d_354_v137_})
        d_356_v141_: _dafny.Map
        d_356_v141_ = _dafny.Map({d_199_v2_: d_199_v2_})
        d_357_v142_: _dafny.Array
        nw70_ = _dafny.Array(None, 23)
        nw70_[int(0)] = _dafny.Map({d_199_v2_: d_209_v11_})
        nw70_[int(1)] = _dafny.Map({d_199_v2_: d_209_v11_})
        nw70_[int(2)] = d_351_v133_
        def iife39_():
            coll27_ = _dafny.Map()
            compr_27_: int
            for compr_27_ in (d_352_v135_).Elements:
                d_358_v134_: int = compr_27_
                if (d_358_v134_) in (d_352_v135_):
                    coll27_[(d_358_v134_) - (d_199_v2_)] = True
            return _dafny.Map(coll27_)
        nw70_[int(3)] = iife39_()
        
        nw70_[int(4)] = d_351_v133_
        nw70_[int(5)] = d_351_v133_
        nw70_[int(6)] = d_351_v133_
        nw70_[int(7)] = d_351_v133_
        nw70_[int(8)] = _dafny.Map({d_199_v2_: d_209_v11_})
        nw70_[int(9)] = d_351_v133_
        nw70_[int(10)] = d_351_v133_
        nw70_[int(11)] = d_351_v133_
        nw70_[int(12)] = d_351_v133_
        nw70_[int(13)] = _dafny.Map({len(d_355_v138_): d_209_v11_})
        nw70_[int(14)] = d_351_v133_
        nw70_[int(15)] = d_351_v133_
        def iife40_():
            coll28_ = _dafny.Map()
            compr_28_: int
            for compr_28_ in _dafny.IntegerRange(793, 340):
                d_359_v139_: int = compr_28_
                if ((793) <= (d_359_v139_)) and ((d_359_v139_) < (340)):
                    coll28_[default__.safeModuloInt(d_359_v139_, len(d_351_v133_))] = d_209_v11_
            return _dafny.Map(coll28_)
        nw70_[int(16)] = iife40_()
        
        nw70_[int(17)] = _dafny.Map({d_199_v2_: d_209_v11_})
        nw70_[int(18)] = d_351_v133_
        nw70_[int(19)] = d_351_v133_
        nw70_[int(20)] = d_351_v133_
        nw70_[int(21)] = _dafny.Map({d_199_v2_: d_209_v11_})
        def iife41_():
            coll29_ = _dafny.Map()
            compr_29_: int
            for compr_29_ in (d_356_v141_).keys.Elements:
                d_360_v140_: int = compr_29_
                if (d_360_v140_) in (d_356_v141_):
                    coll29_[default__.safeModuloInt(d_360_v140_, d_199_v2_)] = d_209_v11_
            return _dafny.Map(coll29_)
        nw70_[int(22)] = iife41_()
        
        d_357_v142_ = nw70_
        d_361_v143_: T1
        nw71_ = C2()
        nw71_.ctor__(d_357_v142_)
        d_361_v143_ = nw71_
        d_362_v144_: _dafny.Set
        d_362_v144_ = _dafny.Set({d_199_v2_})
        d_363_v145_: _dafny.Map
        d_363_v145_ = _dafny.Map({d_354_v137_: len(d_362_v144_)})
        (d_198_globalState_).f9 = (len(d_363_v145_)) * ((d_199_v2_) + (894))
        d_364_v146_: _dafny.Array
        nw72_ = _dafny.Array(_dafny.MultiSet({}), 5)
        d_364_v146_ = nw72_
        d_365_v147_: C8
        nw73_ = C8()
        nw73_.ctor__(d_209_v11_)
        d_365_v147_ = nw73_
        d_366_v148_: _dafny.MultiSet
        d_366_v148_ = _dafny.MultiSet([d_365_v147_, d_365_v147_])
        index52_ = default__.safeIndex(239, (d_364_v146_).length(0))
        (d_364_v146_)[index52_] = (d_366_v148_) - ((d_366_v148_).set(d_365_v147_, default__.abs(d_199_v2_)))
        d_367_v149_: _dafny.Seq
        d_367_v149_ = _dafny.SeqWithoutIsStrInference([d_365_v147_, d_365_v147_, d_365_v147_, d_365_v147_, d_365_v147_])
        index53_ = default__.safeIndex(239, (d_364_v146_).length(0))
        (d_364_v146_)[index53_] = _dafny.MultiSet(d_367_v149_)
        d_368_v150_: _dafny.Seq
        d_368_v150_ = _dafny.SeqWithoutIsStrInference([(d_365_v147_).f10, False, d_209_v11_, (d_365_v147_).f10, False])
        d_369_i10_: int
        d_369_i10_ = 0
        with _dafny.label("2"):
            while (d_209_v11_ if ((d_365_v147_).f10 if not(d_209_v11_) else (d_365_v147_).f10) else (d_368_v150_)[default__.safeIndex(d_199_v2_, len(d_368_v150_))]):
                with _dafny.c_label("2"):
                    if (d_369_i10_) >= (100):
                        raise _dafny.Break("2")
                    d_369_i10_ = (d_369_i10_) + (1)
                    d_370_v151_: int
                    d_371_v152_: _dafny.Set
                    out7_: int
                    out8_: _dafny.Set
                    out7_, out8_ = (d_354_v137_).m3((d_365_v147_).f10, d_198_globalState_)
                    d_370_v151_ = out7_
                    d_371_v152_ = out8_
                    d_372_v153_: _dafny.Array
                    nw74_ = _dafny.Array(_dafny.Set({}), 17)
                    d_372_v153_ = nw74_
                    index54_ = default__.safeIndex(407, (d_372_v153_).length(0))
                    (d_372_v153_)[index54_] = _dafny.Set({d_195_v0_, d_195_v0_})
                    d_373_v154_: C4
                    nw75_ = C4()
                    nw75_.ctor__(d_353_v136_)
                    d_373_v154_ = nw75_
                    d_374_v155_: _dafny.Map
                    d_374_v155_ = _dafny.Map({d_373_v154_: d_195_v0_})
                    d_375_v156_: _dafny.Seq
                    d_375_v156_ = _dafny.SeqWithoutIsStrInference([d_373_v154_])
                    d_376_v157_: _dafny.Set
                    d_376_v157_ = _dafny.Set({d_195_v0_, d_195_v0_, ((d_374_v155_)[(d_375_v156_)[default__.safeIndex(d_370_v151_, len(d_375_v156_))]] if ((d_375_v156_)[default__.safeIndex(d_370_v151_, len(d_375_v156_))]) in (d_374_v155_) else d_195_v0_)})
                    index55_ = default__.safeIndex(407, (d_372_v153_).length(0))
                    (d_372_v153_)[index55_] = d_376_v157_
                    d_377_v158_: _dafny.Seq
                    d_377_v158_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "go"))
                    d_378_v159_: str
                    d_378_v159_ = _dafny.CodePoint('c')
                    d_379_v160_: _dafny.Array
                    nw76_ = _dafny.Array(None, 13)
                    nw76_[int(0)] = d_377_v158_
                    nw76_[int(1)] = d_377_v158_
                    nw76_[int(2)] = (d_377_v158_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gbsy")))
                    nw76_[int(3)] = d_377_v158_
                    nw76_[int(4)] = (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_380_i11_ in range(default__.abs(-187))])) + (d_377_v158_)
                    nw76_[int(5)] = d_377_v158_
                    nw76_[int(6)] = d_377_v158_
                    nw76_[int(7)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c') for d_381_i12_ in range(default__.abs(909))])
                    nw76_[int(8)] = (d_377_v158_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "idw")))
                    nw76_[int(9)] = d_377_v158_
                    nw76_[int(10)] = (d_377_v158_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xnjnwfbu")))
                    nw76_[int(11)] = (d_377_v158_).set(default__.safeIndex(len(d_363_v145_), len(d_377_v158_)), d_378_v159_)
                    nw76_[int(12)] = d_377_v158_
                    d_379_v160_ = nw76_
                    rhs56_ = (d_370_v151_) - (d_199_v2_)
                    rhs57_ = d_379_v160_
                    rhs58_ = d_365_v147_
                    d_370_v151_ = rhs56_
                    d_379_v160_ = rhs57_
                    d_365_v147_ = rhs58_
                    d_382_v161_: C8
                    nw77_ = C8()
                    nw77_.ctor__(True)
                    d_382_v161_ = nw77_
                    pass
            pass
        if d_209_v11_:
            d_383_v162_: _dafny.Seq
            d_383_v162_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ultd"))
            d_384_v163_: _dafny.Seq
            d_384_v163_ = _dafny.SeqWithoutIsStrInference([d_383_v162_])
            d_385_v164_: str
            d_385_v164_ = _dafny.CodePoint('w')
            d_386_v165_: _dafny.Seq
            d_386_v165_ = _dafny.SeqWithoutIsStrInference([d_384_v163_])
            d_387_v166_: _dafny.Map
            d_387_v166_ = _dafny.Map({(d_365_v147_).f10: (d_365_v147_).f10})
            d_388_v167_: _dafny.Array
            nw78_ = _dafny.Array(False, 19)
            d_388_v167_ = nw78_
            d_389_v168_: D7
            d_389_v168_ = D7_DC17(d_383_v162_, d_388_v167_, (d_365_v147_).f10, d_199_v2_)
            d_390_v169_: _dafny.Map
            d_390_v169_ = _dafny.Map({(d_365_v147_).f10: d_199_v2_})
            d_391_v170_: _dafny.Map
            d_391_v170_ = _dafny.Map({d_385_v164_: len(d_390_v169_)})
            d_392_v172_: _dafny.Seq
            d_392_v172_ = _dafny.SeqWithoutIsStrInference([d_383_v162_ for d_393_i14_ in range(default__.abs(779))])
            d_394_v173_: _dafny.Array
            nw79_ = _dafny.Array(None, 16)
            nw79_[int(0)] = d_384_v163_
            nw79_[int(1)] = (d_384_v163_) + (_dafny.SeqWithoutIsStrInference([d_383_v162_ for d_395_i13_ in range(default__.abs(326))]))
            nw79_[int(2)] = (d_384_v163_ if (d_365_v147_).f10 else _dafny.SeqWithoutIsStrInference([(d_383_v162_).set(default__.safeIndex(d_199_v2_, len(d_383_v162_)), d_385_v164_), d_383_v162_, (d_383_v162_).set(default__.safeIndex((0) - (d_199_v2_), len(d_383_v162_)), d_385_v164_)]))
            nw79_[int(3)] = d_384_v163_
            nw79_[int(4)] = (d_386_v165_)[default__.safeIndex(default__.fm0((d_365_v147_).f10, d_198_globalState_), len(d_386_v165_))]
            nw79_[int(5)] = d_384_v163_
            nw79_[int(6)] = d_384_v163_
            nw79_[int(7)] = d_384_v163_
            nw79_[int(8)] = d_384_v163_
            nw79_[int(9)] = d_384_v163_
            nw79_[int(10)] = d_384_v163_
            nw79_[int(11)] = d_384_v163_
            nw79_[int(12)] = default__.fm28(d_385_v164_, len(d_387_v166_), d_199_v2_, (d_365_v147_).f10, d_198_globalState_)
            nw79_[int(13)] = d_384_v163_
            nw79_[int(14)] = d_384_v163_
            def iife42_():
                coll30_ = _dafny.Map()
                compr_30_: int
                for compr_30_ in (d_362_v144_).Elements:
                    d_396_v171_: int = compr_30_
                    if (d_396_v171_) in (d_362_v144_):
                        coll30_[(d_396_v171_) * (len(d_383_v162_))] = d_199_v2_
                return _dafny.Map(coll30_)
            nw79_[int(15)] = (_dafny.SeqWithoutIsStrInference([d_383_v162_, d_383_v162_, (d_389_v168_).cf29, default__.fm2(d_199_v2_, _dafny.MultiSet([((d_391_v170_)[d_385_v164_] if (d_385_v164_) in (d_391_v170_) else len(d_383_v162_)), d_199_v2_]), iife42_()
            , d_198_globalState_)])) + ((d_392_v172_))
            d_394_v173_ = nw79_
            index56_ = default__.safeIndex(43, (d_394_v173_).length(0))
            (d_394_v173_)[index56_] = (d_384_v163_) + (d_384_v163_)
            index57_ = default__.safeIndex(43, (d_394_v173_).length(0))
            (d_394_v173_)[index57_] = d_384_v163_
            index58_ = default__.safeIndex(70, (d_388_v167_).length(0))
            (d_388_v167_)[index58_] = d_209_v11_
            index59_ = default__.safeIndex(70, (d_388_v167_).length(0))
            (d_388_v167_)[index59_] = (d_365_v147_).f10
            d_397_v174_: _dafny.Array
            nw80_ = _dafny.Array(_dafny.Seq({}), 1)
            d_397_v174_ = nw80_
            d_398_v175_: _dafny.Seq
            d_398_v175_ = _dafny.SeqWithoutIsStrInference([d_199_v2_])
            index60_ = default__.safeIndex(286, (d_397_v174_).length(0))
            (d_397_v174_)[index60_] = d_398_v175_
            index61_ = default__.safeIndex(814, (d_195_v0_).length(0))
            (d_195_v0_)[index61_] = d_199_v2_
            d_399_v176_: _dafny.MultiSet
            d_399_v176_ = _dafny.MultiSet([not(False), not((len(d_387_v166_)) != (len(d_383_v162_))), (d_365_v147_).f10, (d_388_v167_)[default__.safeIndex(70, (d_388_v167_).length(0))], (775) not in (d_398_v175_)])
            index62_ = default__.safeIndex(286, (d_397_v174_).length(0))
            index63_ = default__.safeIndex(814, (d_195_v0_).length(0))
            rhs59_ = default__.fm54(False, (d_368_v150_)[default__.safeIndex(d_199_v2_, len(d_368_v150_))], (d_365_v147_).f10, d_198_globalState_)
            rhs60_ = ((d_399_v176_)[default__.fm1((d_365_v147_).f10, _dafny.CodePoint('i'), d_198_globalState_)] if (default__.fm1((d_365_v147_).f10, _dafny.CodePoint('i'), d_198_globalState_)) in (d_399_v176_) else d_199_v2_)
            rhs61_ = default__.safeDivisionInt(len(d_398_v175_), (0) - (d_199_v2_))
            lhs40_ = d_397_v174_
            lhs41_ = default__.safeIndex(286, (d_397_v174_).length(0))
            lhs42_ = d_198_globalState_
            lhs43_ = d_195_v0_
            lhs44_ = default__.safeIndex(814, (d_195_v0_).length(0))
            lhs40_[lhs41_] = rhs59_
            lhs42_.f9 = rhs60_
            lhs43_[lhs44_] = rhs61_
            d_400_v177_: D0
            d_400_v177_ = D0_DC0((d_195_v0_)[default__.safeIndex(814, (d_195_v0_).length(0))])
            d_401_v178_: _dafny.Map
            d_401_v178_ = _dafny.Map({d_383_v162_: d_400_v177_})
            d_401_v178_ = (d_401_v178_).set(d_383_v162_, d_400_v177_)
            d_390_v169_ = _dafny.Map({((d_195_v0_)[default__.safeIndex(814, (d_195_v0_).length(0))]) != (560): d_199_v2_})
        elif True:
            d_209_v11_ = d_209_v11_
            d_402_v179_: _dafny.Array
            nw81_ = _dafny.Array(_dafny.Set({}), 27)
            d_402_v179_ = nw81_
            d_403_v180_: _dafny.Seq
            d_403_v180_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nvdnk"))
            d_404_v181_: _dafny.Set
            d_404_v181_ = _dafny.Set({d_403_v180_})
            index64_ = default__.safeIndex(625, (d_402_v179_).length(0))
            (d_402_v179_)[index64_] = d_404_v181_
            index65_ = default__.safeIndex(625, (d_402_v179_).length(0))
            rhs62_ = d_404_v181_
            rhs63_ = not((d_365_v147_).f10)
            rhs64_ = d_209_v11_
            rhs65_ = (d_368_v150_)[default__.safeIndex(d_199_v2_, len(d_368_v150_))]
            rhs66_ = d_403_v180_
            lhs45_ = d_402_v179_
            lhs46_ = default__.safeIndex(625, (d_402_v179_).length(0))
            lhs47_ = d_198_globalState_
            lhs48_ = d_198_globalState_
            lhs45_[lhs46_] = rhs62_
            d_209_v11_ = rhs63_
            lhs47_.f8 = rhs64_
            lhs48_.f8 = rhs65_
            d_403_v180_ = rhs66_
            (d_198_globalState_).f8 = (d_365_v147_).f10
            d_405_v182_: str
            d_405_v182_ = _dafny.CodePoint('t')
            d_406_v183_: _dafny.Set
            d_406_v183_ = _dafny.Set({d_405_v182_})
            d_407_v184_: _dafny.Set
            d_407_v184_ = d_406_v183_
            source1_ = d_407_v184_
            d_408___mcc_h0_ = source1_
            d_409_cf27_ = d_408___mcc_h0_
            index66_ = default__.safeIndex(682, (d_195_v0_).length(0))
            (d_195_v0_)[index66_] = d_199_v2_
            d_410_v185_: _dafny.MultiSet
            d_410_v185_ = _dafny.MultiSet([d_195_v0_, d_195_v0_])
            d_411_v186_: D0
            d_411_v186_ = D0_DC1(d_410_v185_, d_199_v2_, d_199_v2_)
            d_412_v187_: _dafny.MultiSet
            d_412_v187_ = _dafny.MultiSet([(d_365_v147_).f10])
            index67_ = default__.safeIndex(682, (d_195_v0_).length(0))
            (d_195_v0_)[index67_] = ((d_411_v186_).cf2) - (((d_412_v187_) - ((_dafny.MultiSet(d_368_v150_)).set((d_365_v147_).f10, default__.abs(d_199_v2_)))).cardinality)
            d_403_v180_ = d_403_v180_
            d_413_v188_: C8
            nw82_ = C8()
            nw82_.ctor__(not((d_365_v147_).f10))
            d_413_v188_ = nw82_
            d_414_v189_: D2
            d_414_v189_ = D2_DC8(d_199_v2_, d_405_v182_)
            d_415_v190_: C7
            nw83_ = C7()
            nw83_.ctor__(True, (d_354_v137_).f11, default__.fm40((d_414_v189_).cf16, d_198_globalState_), (d_365_v147_).f10)
            d_415_v190_ = nw83_
            d_416_v191_: _dafny.Seq
            d_416_v191_ = _dafny.SeqWithoutIsStrInference([d_415_v190_])
            d_417_v192_: T0
            nw84_ = C3()
            nw84_.ctor__(d_209_v11_, (d_365_v147_).f10, (d_354_v137_).f11)
            d_417_v192_ = nw84_
            d_418_v193_: _dafny.Map
            d_418_v193_ = _dafny.Map({d_416_v191_: d_417_v192_})
            (d_198_globalState_).f8 = (d_416_v191_) not in ((d_418_v193_) | (d_418_v193_))
            d_419_v194_: _dafny.Array
            nw85_ = _dafny.Array(None, 21)
            nw85_[int(0)] = (d_365_v147_).f10
            nw85_[int(1)] = (d_365_v147_).f10
            nw85_[int(2)] = (d_365_v147_).f10
            nw85_[int(3)] = (d_365_v147_).f10
            nw85_[int(4)] = (d_365_v147_).f10
            nw85_[int(5)] = (d_365_v147_).f10
            nw85_[int(6)] = (d_365_v147_).f10
            nw85_[int(7)] = (d_365_v147_).f10
            nw85_[int(8)] = (d_365_v147_).f10
            nw85_[int(9)] = (d_365_v147_).f10
            nw85_[int(10)] = (d_365_v147_).f10
            nw85_[int(11)] = (d_365_v147_).f10
            nw85_[int(12)] = d_209_v11_
            nw85_[int(13)] = (d_365_v147_).f10
            nw85_[int(14)] = (d_365_v147_).f10
            nw85_[int(15)] = (d_365_v147_).f10
            nw85_[int(16)] = (d_365_v147_).f10
            nw85_[int(17)] = (d_365_v147_).f10
            nw85_[int(18)] = (d_365_v147_).f10
            nw85_[int(19)] = d_209_v11_
            nw85_[int(20)] = d_209_v11_
            d_419_v194_ = nw85_
            d_420_v195_: D7
            d_420_v195_ = D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")), d_419_v194_, d_209_v11_, len(_dafny.SeqWithoutIsStrInference([d_199_v2_ for d_421_i15_ in range(default__.abs(929))])))
            d_422_v196_: _dafny.Array
            nw86_ = _dafny.Array(None, 6)
            nw86_[int(0)] = (d_365_v147_).f10
            nw86_[int(1)] = (d_365_v147_).f10
            nw86_[int(2)] = not((d_365_v147_).f10)
            nw86_[int(3)] = not (((d_351_v133_)[d_199_v2_] if (d_199_v2_) in (d_351_v133_) else d_209_v11_)) or (d_209_v11_)
            nw86_[int(4)] = (d_365_v147_).f10
            nw86_[int(5)] = (d_368_v150_)[default__.safeIndex((d_420_v195_).cf32, len(d_368_v150_))]
            d_422_v196_ = nw86_
            index68_ = default__.safeIndex(941, (d_422_v196_).length(0))
            (d_422_v196_)[index68_] = d_209_v11_
            index69_ = default__.safeIndex(941, (d_422_v196_).length(0))
            (d_422_v196_)[index69_] = (d_365_v147_).f10
        d_423_v197_: _dafny.Seq
        d_423_v197_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "negiseois"))
        d_424_v198_: C6
        nw87_ = C6()
        nw87_.ctor__(d_199_v2_, (d_423_v197_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('q') for d_425_i16_ in range(default__.abs(120))])))
        d_424_v198_ = nw87_
        hi2_ = default__.safeDivisionInt(d_199_v2_, (d_424_v198_).f15)
        for d_426_i17_ in range(len(d_424_v198_.f16), hi2_):
            d_427_v199_: str
            d_427_v199_ = _dafny.CodePoint('j')
            (d_198_globalState_).f9 = (((d_424_v198_).f15 if (d_365_v147_).f10 else d_199_v2_) if default__.fm1((d_365_v147_).f10, d_427_v199_, d_198_globalState_) else (d_199_v2_) + (len(d_424_v198_.f16)))
            d_428_v200_: D2
            d_428_v200_ = D2_DC5(d_424_v198_.f16)
            d_428_v200_ = D2_DC5(d_424_v198_.f16)
            (d_198_globalState_).f9 = (0) - ((d_352_v135_).cardinality)
            d_429_v201_: _dafny.Seq
            d_429_v201_ = _dafny.SeqWithoutIsStrInference([(d_424_v198_).f15])
            d_430_v202_: _dafny.Array
            nw88_ = _dafny.Array(None, 5)
            nw88_[int(0)] = d_429_v201_
            nw88_[int(1)] = (_dafny.SeqWithoutIsStrInference([(d_424_v198_).f15])) + (d_429_v201_)
            nw88_[int(2)] = (D5_DC13(d_429_v201_)).cf25
            nw88_[int(3)] = d_429_v201_
            nw88_[int(4)] = (d_429_v201_ if d_209_v11_ else d_429_v201_)
            d_430_v202_ = nw88_
            index70_ = default__.safeIndex(603, (d_430_v202_).length(0))
            (d_430_v202_)[index70_] = _dafny.SeqWithoutIsStrInference([(d_424_v198_).f15 for d_431_i18_ in range(default__.abs(159))])
            index71_ = default__.safeIndex(603, (d_430_v202_).length(0))
            (d_430_v202_)[index71_] = d_429_v201_
        d_432_v203_: D7
        d_432_v203_ = D7_DC16(d_368_v150_)
        d_209_v11_ = (d_368_v150_)[default__.safeIndex((len(((d_432_v203_).cf28).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([len(d_423_v197_), d_199_v2_, len(d_424_v198_.f16)])), len((d_432_v203_).cf28)), d_209_v11_))) * (d_199_v2_), len(d_368_v150_))]
        d_433_v204_: _dafny.Array
        nw89_ = _dafny.Array(False, 10)
        d_433_v204_ = nw89_
        d_434_v205_: _dafny.Seq
        d_434_v205_ = _dafny.SeqWithoutIsStrInference([d_195_v0_])
        index72_ = default__.safeIndex(690, (d_433_v204_).length(0))
        (d_433_v204_)[index72_] = ((d_424_v198_).f15) >= (len(d_434_v205_))
        d_435_v206_: _dafny.Set
        d_435_v206_ = _dafny.Set({(d_365_v147_).f10})
        d_436_v207_: _dafny.Array
        nw90_ = _dafny.Array(_dafny.Seq({}), 11)
        d_436_v207_ = nw90_
        d_437_v208_: _dafny.Seq
        d_437_v208_ = _dafny.SeqWithoutIsStrInference([510, d_199_v2_])
        index73_ = default__.safeIndex(521, (d_436_v207_).length(0))
        (d_436_v207_)[index73_] = d_437_v208_
        d_438_v209_: D7
        d_438_v209_ = D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ervymfer")), d_433_v204_, d_209_v11_, d_199_v2_)
        d_439_v210_: _dafny.Map
        d_439_v210_ = _dafny.Map({(0) - ((d_424_v198_).f15): d_424_v198_.f16})
        d_440_v211_: D9
        d_440_v211_ = D9_DC22((0) - ((d_424_v198_).f15), d_439_v210_)
        d_441_v212_: _dafny.Seq
        d_441_v212_ = _dafny.SeqWithoutIsStrInference([d_437_v208_])
        index74_ = default__.safeIndex(690, (d_433_v204_).length(0))
        index75_ = default__.safeIndex(521, (d_436_v207_).length(0))
        rhs67_ = (d_438_v209_).cf30
        rhs68_ = (((d_365_v147_).f10 if d_209_v11_ else d_209_v11_)) == (True)
        rhs69_ = default__.fm53(d_198_globalState_)
        rhs70_ = (((d_437_v208_) + (d_437_v208_)) + (d_437_v208_)).set(default__.safeIndex(d_199_v2_, len(((d_437_v208_) + (d_437_v208_)) + (d_437_v208_))), ((d_352_v135_)[(d_440_v211_).cf40] if ((d_440_v211_).cf40) in (d_352_v135_) else len((_dafny.SeqWithoutIsStrInference([len((d_441_v212_)[default__.safeIndex(88, len(d_441_v212_))])])).set(default__.safeIndex((d_424_v198_).f15, len(_dafny.SeqWithoutIsStrInference([len((d_441_v212_)[default__.safeIndex(88, len(d_441_v212_))])]))), d_199_v2_))))
        lhs49_ = d_433_v204_
        lhs50_ = default__.safeIndex(690, (d_433_v204_).length(0))
        lhs51_ = d_436_v207_
        lhs52_ = default__.safeIndex(521, (d_436_v207_).length(0))
        d_433_v204_ = rhs67_
        lhs49_[lhs50_] = rhs68_
        d_435_v206_ = rhs69_
        lhs51_[lhs52_] = rhs70_
        hi3_ = ((0) - (-60)) * (len(d_423_v197_))
        for d_442_i19_ in range((d_424_v198_).f15, hi3_):
            d_199_v2_ = (_dafny.MultiSet([d_438_v209_, d_438_v209_, d_438_v209_])).cardinality
            index76_ = default__.safeIndex(690, (d_433_v204_).length(0))
            (d_433_v204_)[index76_] = not(d_209_v11_)
            d_443_v213_: _dafny.Array
            nw91_ = _dafny.Array(None, 7)
            nw91_[int(0)] = _dafny.Map({672: ((d_352_v135_).set((d_424_v198_).f15, default__.abs(619))).cardinality})
            nw91_[int(1)] = d_356_v141_
            nw91_[int(2)] = d_356_v141_
            nw91_[int(3)] = d_356_v141_
            nw91_[int(4)] = d_356_v141_
            nw91_[int(5)] = d_356_v141_
            nw91_[int(6)] = d_356_v141_
            d_443_v213_ = nw91_
            d_444_v214_: str
            d_444_v214_ = _dafny.CodePoint('a')
            d_445_v215_: _dafny.Map
            d_445_v215_ = _dafny.Map({d_443_v213_: not (False) or (default__.fm1((d_365_v147_).f10, d_444_v214_, d_198_globalState_))})
            d_445_v215_ = (d_445_v215_).set(d_443_v213_, (d_365_v147_).fm6(d_198_globalState_))
            (d_198_globalState_).f9 = (0) - ((d_424_v198_).f15)
        d_446_v216_: D13
        d_446_v216_ = D13_DC29(_dafny.Map({d_199_v2_: (d_365_v147_).f10}))
        pat_let_tv5_ = d_368_v150_
        pat_let_tv6_ = d_368_v150_
        pat_let_tv7_ = d_424_v198_
        pat_let_tv8_ = d_424_v198_
        pat_let_tv9_ = d_424_v198_
        def lambda20_(source2_):
            if source2_.is_DC30:
                d_447___mcc_h1_ = source2_.cf51
                d_448_cf51_ = d_447___mcc_h1_
                return (_dafny.MultiSet([pat_let_tv5_])).ispropersubset(_dafny.MultiSet([pat_let_tv6_]))
            elif source2_.is_DC31:
                d_449___mcc_h2_ = source2_.cf52
                d_450_cf52_ = d_449___mcc_h2_
                return ((pat_let_tv7_.f16)[default__.safeIndex((pat_let_tv8_).f15, len(pat_let_tv9_.f16))]) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfaxiys")))
            elif True:
                d_451___mcc_h3_ = source2_.cf50
                d_452_cf50_ = d_451___mcc_h3_
                return False

        (d_198_globalState_).f8 = lambda20_(d_446_v216_)
        if (d_365_v147_).f10:
            d_453_v217_: str
            d_453_v217_ = _dafny.CodePoint('g')
            (d_198_globalState_).f8 = default__.fm1((d_435_v206_).ispropersubset(d_435_v206_), d_453_v217_, d_198_globalState_)
            (d_198_globalState_).f8 = (d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))]
            d_454_v218_: _dafny.Array
            nw92_ = _dafny.Array(_dafny.Seq({}), 4)
            d_454_v218_ = nw92_
            index77_ = default__.safeIndex(947, (d_454_v218_).length(0))
            (d_454_v218_)[index77_] = d_368_v150_
            d_455_v219_: _dafny.Array
            def lambda21_(d_456_v217_):
                def lambda22_(d_457_i20_):
                    return d_456_v217_

                return lambda22_

            init11_ = lambda21_(d_453_v217_)
            nw93_ = _dafny.Array(None, 24)
            for i0_11_ in range(nw93_.length(0)):
                nw93_[i0_11_] = init11_(i0_11_)
            d_455_v219_ = nw93_
            d_458_v220_: _dafny.Map
            d_458_v220_ = _dafny.Map({(d_365_v147_).f10: d_455_v219_})
            d_459_v221_: _dafny.Map
            d_459_v221_ = _dafny.Map({((d_458_v220_)[(d_365_v147_).f10] if ((d_365_v147_).f10) in (d_458_v220_) else d_455_v219_): (_dafny.MultiSet([(d_365_v147_).f10])).set((d_365_v147_).f10, default__.abs((d_424_v198_).f15))})
            d_460_v222_: _dafny.Map
            d_460_v222_ = _dafny.Map({d_361_v143_: False})
            d_461_v223_: C3
            nw94_ = C3()
            nw94_.ctor__(d_209_v11_, (d_365_v147_).f10, (d_354_v137_).f11)
            d_461_v223_ = nw94_
            d_462_v224_: _dafny.Map
            d_462_v224_ = _dafny.Map({d_461_v223_: d_361_v143_})
            d_463_v225_: _dafny.MultiSet
            d_463_v225_ = _dafny.MultiSet([((d_460_v222_)[((d_462_v224_)[d_461_v223_] if (d_461_v223_) in (d_462_v224_) else d_361_v143_)] if (((d_462_v224_)[d_461_v223_] if (d_461_v223_) in (d_462_v224_) else d_361_v143_)) in (d_460_v222_) else (d_365_v147_).f10)])
            d_464_v226_: _dafny.Map
            d_464_v226_ = _dafny.Map({d_433_v204_: len((d_459_v221_).set(d_455_v219_, d_463_v225_))})
            d_465_v227_: _dafny.Array
            nw95_ = _dafny.Array(None, 17)
            d_465_v227_ = nw95_
            d_466_v228_: _dafny.Seq
            d_466_v228_ = _dafny.SeqWithoutIsStrInference([d_465_v227_, d_465_v227_])
            index78_ = default__.safeIndex(690, (d_433_v204_).length(0))
            index79_ = default__.safeIndex(947, (d_454_v218_).length(0))
            index80_ = default__.safeIndex(690, (d_433_v204_).length(0))
            rhs71_ = (((d_356_v141_)[(d_424_v198_).f15] if ((d_424_v198_).f15) in (d_356_v141_) else ((d_464_v226_)[d_433_v204_] if (d_433_v204_) in (d_464_v226_) else len(_dafny.SeqWithoutIsStrInference([(d_365_v147_).f10]))))) != (((d_424_v198_).f15) + ((d_424_v198_).f15))
            rhs72_ = default__.fm0((d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))], d_198_globalState_)
            rhs73_ = ((_dafny.Map({d_354_v137_: (d_424_v198_).f15})).set(d_354_v137_, d_199_v2_)) | (d_363_v145_)
            rhs74_ = ((d_368_v150_) + (d_368_v150_)) + (d_368_v150_)
            rhs75_ = not ((d_365_v147_).f10) or (((0) - (len(d_466_v228_))) != ((d_424_v198_).f15))
            lhs53_ = d_433_v204_
            lhs54_ = default__.safeIndex(690, (d_433_v204_).length(0))
            lhs55_ = d_454_v218_
            lhs56_ = default__.safeIndex(947, (d_454_v218_).length(0))
            lhs57_ = d_433_v204_
            lhs58_ = default__.safeIndex(690, (d_433_v204_).length(0))
            lhs53_[lhs54_] = rhs71_
            d_199_v2_ = rhs72_
            d_363_v145_ = rhs73_
            lhs55_[lhs56_] = rhs74_
            lhs57_[lhs58_] = rhs75_
            d_467_v229_: C5
            nw96_ = C5()
            nw96_.ctor__((d_361_v143_).f11)
            d_467_v229_ = nw96_
            d_468_v230_: _dafny.Array
            nw97_ = _dafny.Array(None, 8)
            nw97_[int(0)] = d_467_v229_
            nw97_[int(1)] = d_467_v229_
            nw97_[int(2)] = d_467_v229_
            nw97_[int(3)] = d_467_v229_
            nw97_[int(4)] = d_467_v229_
            nw97_[int(5)] = d_467_v229_
            nw97_[int(6)] = d_467_v229_
            nw97_[int(7)] = d_467_v229_
            d_468_v230_ = nw97_
            d_469_v231_: _dafny.Array
            d_469_v231_ = d_468_v230_
            rhs76_ = (d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))]
            rhs77_ = d_469_v231_
            lhs59_ = d_198_globalState_
            lhs59_.f8 = rhs76_
            d_469_v231_ = rhs77_
            index81_ = default__.safeIndex(833, (d_195_v0_).length(0))
            (d_195_v0_)[index81_] = (d_424_v198_).f15
            d_470_v232_: _dafny.Array
            nw98_ = _dafny.Array(None, 2)
            d_470_v232_ = nw98_
            d_471_v233_: D0
            d_471_v233_ = D0_DC0(d_199_v2_)
            d_472_v234_: T2
            nw99_ = C1()
            nw99_.ctor__((d_365_v147_).f10, (d_365_v147_).f10, d_471_v233_, (d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))])
            d_472_v234_ = nw99_
            d_473_v235_: D20
            d_473_v235_ = D20_DC38(d_472_v234_)
            d_474_v236_: _dafny.Map
            d_474_v236_ = _dafny.Map({(d_424_v198_).f15: (d_473_v235_).cf59})
            index82_ = default__.safeIndex(487, (d_470_v232_).length(0))
            (d_470_v232_)[index82_] = ((d_474_v236_)[1] if (1) in (d_474_v236_) else d_472_v234_)
            d_475_v237_: _dafny.MultiSet
            d_475_v237_ = _dafny.MultiSet([d_461_v223_])
            index83_ = default__.safeIndex(833, (d_195_v0_).length(0))
            index84_ = default__.safeIndex(487, (d_470_v232_).length(0))
            rhs78_ = not (not((d_435_v206_).issubset(d_435_v206_))) or ((d_475_v237_).ispropersubset(d_475_v237_))
            rhs79_ = default__.safeDivisionInt(len(default__.fm39((d_424_v198_).f15, d_198_globalState_)), (d_424_v198_).f15)
            rhs80_ = (d_435_v206_).intersection(_dafny.Set({(d_472_v234_).f13}))
            rhs81_ = (d_424_v198_.f16 if (d_365_v147_).fm6(d_198_globalState_) else d_424_v198_.f16)
            rhs82_ = d_472_v234_
            lhs60_ = d_198_globalState_
            lhs61_ = d_195_v0_
            lhs62_ = default__.safeIndex(833, (d_195_v0_).length(0))
            lhs63_ = d_470_v232_
            lhs64_ = default__.safeIndex(487, (d_470_v232_).length(0))
            lhs60_.f8 = rhs78_
            lhs61_[lhs62_] = rhs79_
            d_435_v206_ = rhs80_
            d_423_v197_ = rhs81_
            lhs63_[lhs64_] = rhs82_
        elif True:
            (d_198_globalState_).f9 = len(d_424_v198_.f16)
            d_476_v238_: _dafny.MultiSet
            d_476_v238_ = _dafny.MultiSet([(d_356_v141_) != (d_356_v141_), False])
            d_476_v238_ = d_476_v238_
            d_199_v2_ = d_199_v2_
            d_477_v239_: _dafny.Map
            d_477_v239_ = _dafny.Map({(d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))]: (d_424_v198_).f15})
            d_478_v240_: str
            d_478_v240_ = _dafny.CodePoint('i')
            d_479_v241_: _dafny.Seq
            d_479_v241_ = _dafny.SeqWithoutIsStrInference([d_477_v239_, default__.fm50(d_478_v240_, (d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))], (d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))], d_198_globalState_)])
            d_479_v241_ = (d_479_v241_) + ((d_479_v241_) + (d_479_v241_))
            d_480_v242_: _dafny.Map
            d_480_v242_ = _dafny.Map({d_423_v197_: (d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))]})
            d_481_v243_: D0
            d_481_v243_ = D0_DC2(d_476_v238_, ((d_480_v242_)[d_424_v198_.f16] if (d_424_v198_.f16) in (d_480_v242_) else (d_365_v147_).f10), (d_424_v198_).f15, (d_365_v147_).f10)
            d_482_v244_: _dafny.Map
            d_482_v244_ = _dafny.Map({True: d_481_v243_})
            d_482_v244_ = (d_482_v244_).set((d_433_v204_)[default__.safeIndex(690, (d_433_v204_).length(0))], d_481_v243_)
        _dafny.print(_dafny.string_of((d_195_v0_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_195_v0_)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_197_v1_)[0])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_globalState_).f0))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_).f1)[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_198_globalState_).f2).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[1]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[2]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[3]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[4]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[5]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[6]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[7]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[8]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[9]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[10]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[11]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[12]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[13]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[14]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[15]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_198_globalState_.f3)[0])[16]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_globalState_).f4))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_globalState_).f5))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_globalState_).f6))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_198_globalState_).f7))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_198_globalState_.f8))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_198_globalState_.f9))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_199_v2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_209_v11_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_210_i2_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_351_v133_) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_352_v135_) == (_dafny.MultiSet([-1, -1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_355_v138_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_356_v141_) == (_dafny.Map({-1: -1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[0]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[1]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[2]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[3]) == (_dafny.Map({0: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[4]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[5]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[6]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[7]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[8]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[9]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[10]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[11]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[12]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[13]) == (_dafny.Map({1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[14]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[15]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[16]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[17]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[18]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[19]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[20]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[21]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_357_v142_)[22]) == (_dafny.Map({0: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[0]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[1]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[2]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[3]) == (_dafny.Map({0: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[4]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[5]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[6]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[7]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[8]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[9]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[10]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[11]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[12]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[13]) == (_dafny.Map({1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[14]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[15]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[16]) == (_dafny.Map({}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[17]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[18]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[19]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[20]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[21]) == (_dafny.Map({-1: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((((d_361_v143_).f11)[22]) == (_dafny.Map({0: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_362_v144_) == (_dafny.Set({-1}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_363_v145_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_364_v146_)[4]).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_365_v147_).f10))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_366_v148_).cardinality))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_367_v149_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_368_v150_) == (_dafny.SeqWithoutIsStrInference([True, False, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(d_369_i10_))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_423_v197_).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_424_v198_).f15))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print((d_424_v198_.f16).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_432_v203_).cf28) == (_dafny.SeqWithoutIsStrInference([True, False, True, True, False]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_433_v204_)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(len(d_434_v205_)))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_435_v206_) == (_dafny.Set({True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_436_v207_)[4]) == (_dafny.SeqWithoutIsStrInference([1, -1, 510, -1, 510, -1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_437_v208_) == (_dafny.SeqWithoutIsStrInference([510, -1]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(((d_438_v209_).cf29).VerbatimString(False))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_438_v209_).cf30)[0]))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_438_v209_).cf31))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_438_v209_).cf32))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_439_v210_) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "negiseoisqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_440_v211_).cf40))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_440_v211_).cf41) == (_dafny.Map({1: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "negiseoisqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq"))}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of((d_441_v212_) == (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([510, -1])]))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))
        _dafny.print(_dafny.string_of(((d_446_v216_).cf50) == (_dafny.Map({3: True}))))
        _dafny.print((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "\n"))).VerbatimString(False))


class D0:
    @classmethod
    def default(cls, ):
        return lambda: D0_DC1(_dafny.MultiSet({}), int(0), int(0))
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

class D0_DC1(D0, NamedTuple('DC1', [('cf1', Any), ('cf2', Any), ('cf3', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC1({_dafny.string_of(self.cf1)}, {_dafny.string_of(self.cf2)}, {_dafny.string_of(self.cf3)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC1) and self.cf1 == __o.cf1 and self.cf2 == __o.cf2 and self.cf3 == __o.cf3
    def __hash__(self) -> int:
        return super().__hash__()

class D0_DC2(D0, NamedTuple('DC2', [('cf4', Any), ('cf5', Any), ('cf6', Any), ('cf7', Any)])):
    def __dafnystr__(self) -> str:
        return f'D0.DC2({_dafny.string_of(self.cf4)}, {_dafny.string_of(self.cf5)}, {_dafny.string_of(self.cf6)}, {_dafny.string_of(self.cf7)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D0_DC2) and self.cf4 == __o.cf4 and self.cf5 == __o.cf5 and self.cf6 == __o.cf6 and self.cf7 == __o.cf7
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
        return lambda: D1_DC4(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC4(self) -> bool:
        return isinstance(self, D1_DC4)
    @property
    def is_DC3(self) -> bool:
        return isinstance(self, D1_DC3)

class D1_DC4(D1, NamedTuple('DC4', [('cf9', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC4({_dafny.string_of(self.cf9)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC4) and self.cf9 == __o.cf9
    def __hash__(self) -> int:
        return super().__hash__()

class D1_DC3(D1, NamedTuple('DC3', [('cf8', Any)])):
    def __dafnystr__(self) -> str:
        return f'D1.DC3({_dafny.string_of(self.cf8)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D1_DC3) and self.cf8 == __o.cf8
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
    def is_DC7(self) -> bool:
        return isinstance(self, D2_DC7)
    @property
    def is_DC8(self) -> bool:
        return isinstance(self, D2_DC8)
    @property
    def is_DC5(self) -> bool:
        return isinstance(self, D2_DC5)

class D2_DC6(D2, NamedTuple('DC6', [])):
    def __dafnystr__(self) -> str:
        return f'D2.DC6'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC6)
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC7(D2, NamedTuple('DC7', [('cf11', Any), ('cf12', Any), ('cf13', Any), ('cf14', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC7({_dafny.string_of(self.cf11)}, {_dafny.string_of(self.cf12)}, {_dafny.string_of(self.cf13)}, {self.cf14.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC7) and self.cf11 == __o.cf11 and self.cf12 == __o.cf12 and self.cf13 == __o.cf13 and self.cf14 == __o.cf14
    def __hash__(self) -> int:
        return super().__hash__()

class D2_DC8(D2, NamedTuple('DC8', [('cf15', Any), ('cf16', Any)])):
    def __dafnystr__(self) -> str:
        return f'D2.DC8({_dafny.string_of(self.cf15)}, {_dafny.string_of(self.cf16)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D2_DC8) and self.cf15 == __o.cf15 and self.cf16 == __o.cf16
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
        return lambda: _dafny.MultiSet({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC9(self) -> bool:
        return isinstance(self, D3_DC9)

class D3_DC9(D3, NamedTuple('DC9', [('cf17', Any)])):
    def __dafnystr__(self) -> str:
        return f'D3.DC9({_dafny.string_of(self.cf17)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D3_DC9) and self.cf17 == __o.cf17
    def __hash__(self) -> int:
        return super().__hash__()


class D4:
    @classmethod
    def default(cls, ):
        return lambda: D4_DC11(int(0), int(0), False)
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

class D4_DC11(D4, NamedTuple('DC11', [('cf19', Any), ('cf20', Any), ('cf21', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC11({_dafny.string_of(self.cf19)}, {_dafny.string_of(self.cf20)}, {_dafny.string_of(self.cf21)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC11) and self.cf19 == __o.cf19 and self.cf20 == __o.cf20 and self.cf21 == __o.cf21
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC12(D4, NamedTuple('DC12', [('cf22', Any), ('cf23', Any), ('cf24', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC12({_dafny.string_of(self.cf22)}, {_dafny.string_of(self.cf23)}, {_dafny.string_of(self.cf24)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC12) and self.cf22 == __o.cf22 and self.cf23 == __o.cf23 and self.cf24 == __o.cf24
    def __hash__(self) -> int:
        return super().__hash__()

class D4_DC10(D4, NamedTuple('DC10', [('cf18', Any)])):
    def __dafnystr__(self) -> str:
        return f'D4.DC10({_dafny.string_of(self.cf18)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D4_DC10) and self.cf18 == __o.cf18
    def __hash__(self) -> int:
        return super().__hash__()


class D5:
    @classmethod
    def default(cls, ):
        return lambda: D5_DC14(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC14(self) -> bool:
        return isinstance(self, D5_DC14)
    @property
    def is_DC13(self) -> bool:
        return isinstance(self, D5_DC13)

class D5_DC14(D5, NamedTuple('DC14', [('cf26', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC14({_dafny.string_of(self.cf26)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC14) and self.cf26 == __o.cf26
    def __hash__(self) -> int:
        return super().__hash__()

class D5_DC13(D5, NamedTuple('DC13', [('cf25', Any)])):
    def __dafnystr__(self) -> str:
        return f'D5.DC13({_dafny.string_of(self.cf25)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D5_DC13) and self.cf25 == __o.cf25
    def __hash__(self) -> int:
        return super().__hash__()


class D6:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC15(self) -> bool:
        return isinstance(self, D6_DC15)

class D6_DC15(D6, NamedTuple('DC15', [('cf27', Any)])):
    def __dafnystr__(self) -> str:
        return f'D6.DC15({_dafny.string_of(self.cf27)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D6_DC15) and self.cf27 == __o.cf27
    def __hash__(self) -> int:
        return super().__hash__()


class D7:
    @classmethod
    def default(cls, ):
        return lambda: D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), _dafny.Array(None, 0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC17(self) -> bool:
        return isinstance(self, D7_DC17)
    @property
    def is_DC16(self) -> bool:
        return isinstance(self, D7_DC16)

class D7_DC17(D7, NamedTuple('DC17', [('cf29', Any), ('cf30', Any), ('cf31', Any), ('cf32', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC17({self.cf29.VerbatimString(True)}, {_dafny.string_of(self.cf30)}, {_dafny.string_of(self.cf31)}, {_dafny.string_of(self.cf32)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC17) and self.cf29 == __o.cf29 and self.cf30 == __o.cf30 and self.cf31 == __o.cf31 and self.cf32 == __o.cf32
    def __hash__(self) -> int:
        return super().__hash__()

class D7_DC16(D7, NamedTuple('DC16', [('cf28', Any)])):
    def __dafnystr__(self) -> str:
        return f'D7.DC16({_dafny.string_of(self.cf28)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D7_DC16) and self.cf28 == __o.cf28
    def __hash__(self) -> int:
        return super().__hash__()


class D8:
    @classmethod
    def default(cls, ):
        return lambda: D8_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC19(self) -> bool:
        return isinstance(self, D8_DC19)
    @property
    def is_DC20(self) -> bool:
        return isinstance(self, D8_DC20)
    @property
    def is_DC18(self) -> bool:
        return isinstance(self, D8_DC18)

class D8_DC19(D8, NamedTuple('DC19', [('cf34', Any), ('cf35', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC19({self.cf34.VerbatimString(True)}, {_dafny.string_of(self.cf35)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC19) and self.cf34 == __o.cf34 and self.cf35 == __o.cf35
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC20(D8, NamedTuple('DC20', [('cf36', Any), ('cf37', Any), ('cf38', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC20({_dafny.string_of(self.cf36)}, {_dafny.string_of(self.cf37)}, {_dafny.string_of(self.cf38)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC20) and self.cf36 == __o.cf36 and self.cf37 == __o.cf37 and self.cf38 == __o.cf38
    def __hash__(self) -> int:
        return super().__hash__()

class D8_DC18(D8, NamedTuple('DC18', [('cf33', Any)])):
    def __dafnystr__(self) -> str:
        return f'D8.DC18({_dafny.string_of(self.cf33)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D8_DC18) and self.cf33 == __o.cf33
    def __hash__(self) -> int:
        return super().__hash__()


class D9:
    @classmethod
    def default(cls, ):
        return lambda: D9_DC22(int(0), _dafny.Map({}))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC22(self) -> bool:
        return isinstance(self, D9_DC22)
    @property
    def is_DC21(self) -> bool:
        return isinstance(self, D9_DC21)

class D9_DC22(D9, NamedTuple('DC22', [('cf40', Any), ('cf41', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC22({_dafny.string_of(self.cf40)}, {_dafny.string_of(self.cf41)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC22) and self.cf40 == __o.cf40 and self.cf41 == __o.cf41
    def __hash__(self) -> int:
        return super().__hash__()

class D9_DC21(D9, NamedTuple('DC21', [('cf39', Any)])):
    def __dafnystr__(self) -> str:
        return f'D9.DC21({_dafny.string_of(self.cf39)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D9_DC21) and self.cf39 == __o.cf39
    def __hash__(self) -> int:
        return super().__hash__()


class D10:
    @classmethod
    def default(cls, ):
        return lambda: D10_DC24(_dafny.Array(None, 0), _dafny.Array(None, 0), _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC24(self) -> bool:
        return isinstance(self, D10_DC24)
    @property
    def is_DC23(self) -> bool:
        return isinstance(self, D10_DC23)

class D10_DC24(D10, NamedTuple('DC24', [('cf43', Any), ('cf44', Any), ('cf45', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC24({_dafny.string_of(self.cf43)}, {_dafny.string_of(self.cf44)}, {self.cf45.VerbatimString(True)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC24) and self.cf43 == __o.cf43 and self.cf44 == __o.cf44 and self.cf45 == __o.cf45
    def __hash__(self) -> int:
        return super().__hash__()

class D10_DC23(D10, NamedTuple('DC23', [('cf42', Any)])):
    def __dafnystr__(self) -> str:
        return f'D10.DC23({_dafny.string_of(self.cf42)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D10_DC23) and self.cf42 == __o.cf42
    def __hash__(self) -> int:
        return super().__hash__()


class D11:
    @classmethod
    def default(cls, ):
        return lambda: D11_DC26(int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC26(self) -> bool:
        return isinstance(self, D11_DC26)
    @property
    def is_DC25(self) -> bool:
        return isinstance(self, D11_DC25)
    @property
    def is_DC27(self) -> bool:
        return isinstance(self, D11_DC27)

class D11_DC26(D11, NamedTuple('DC26', [('cf47', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC26({_dafny.string_of(self.cf47)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC26) and self.cf47 == __o.cf47
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC25(D11, NamedTuple('DC25', [('cf46', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC25({_dafny.string_of(self.cf46)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC25) and self.cf46 == __o.cf46
    def __hash__(self) -> int:
        return super().__hash__()

class D11_DC27(D11, NamedTuple('DC27', [('cf48', Any)])):
    def __dafnystr__(self) -> str:
        return f'D11.DC27({_dafny.string_of(self.cf48)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D11_DC27) and self.cf48 == __o.cf48
    def __hash__(self) -> int:
        return super().__hash__()


class D12:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC28(self) -> bool:
        return isinstance(self, D12_DC28)

class D12_DC28(D12, NamedTuple('DC28', [('cf49', Any)])):
    def __dafnystr__(self) -> str:
        return f'D12.DC28({_dafny.string_of(self.cf49)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D12_DC28) and self.cf49 == __o.cf49
    def __hash__(self) -> int:
        return super().__hash__()


class D13:
    @classmethod
    def default(cls, ):
        return lambda: D13_DC30(_dafny.CodePoint('D'))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC30(self) -> bool:
        return isinstance(self, D13_DC30)
    @property
    def is_DC31(self) -> bool:
        return isinstance(self, D13_DC31)
    @property
    def is_DC29(self) -> bool:
        return isinstance(self, D13_DC29)

class D13_DC30(D13, NamedTuple('DC30', [('cf51', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC30({_dafny.string_of(self.cf51)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC30) and self.cf51 == __o.cf51
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC31(D13, NamedTuple('DC31', [('cf52', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC31({_dafny.string_of(self.cf52)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC31) and self.cf52 == __o.cf52
    def __hash__(self) -> int:
        return super().__hash__()

class D13_DC29(D13, NamedTuple('DC29', [('cf50', Any)])):
    def __dafnystr__(self) -> str:
        return f'D13.DC29({_dafny.string_of(self.cf50)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D13_DC29) and self.cf50 == __o.cf50
    def __hash__(self) -> int:
        return super().__hash__()


class D14:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Set({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC32(self) -> bool:
        return isinstance(self, D14_DC32)

class D14_DC32(D14, NamedTuple('DC32', [('cf53', Any)])):
    def __dafnystr__(self) -> str:
        return f'D14.DC32({_dafny.string_of(self.cf53)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D14_DC32) and self.cf53 == __o.cf53
    def __hash__(self) -> int:
        return super().__hash__()


class D15:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Array(None, 0)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC33(self) -> bool:
        return isinstance(self, D15_DC33)

class D15_DC33(D15, NamedTuple('DC33', [('cf54', Any)])):
    def __dafnystr__(self) -> str:
        return f'D15.DC33({_dafny.string_of(self.cf54)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D15_DC33) and self.cf54 == __o.cf54
    def __hash__(self) -> int:
        return super().__hash__()


class D16:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Map({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC34(self) -> bool:
        return isinstance(self, D16_DC34)

class D16_DC34(D16, NamedTuple('DC34', [('cf55', Any)])):
    def __dafnystr__(self) -> str:
        return f'D16.DC34({_dafny.string_of(self.cf55)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D16_DC34) and self.cf55 == __o.cf55
    def __hash__(self) -> int:
        return super().__hash__()


class D17:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC35(self) -> bool:
        return isinstance(self, D17_DC35)

class D17_DC35(D17, NamedTuple('DC35', [('cf56', Any)])):
    def __dafnystr__(self) -> str:
        return f'D17.DC35({_dafny.string_of(self.cf56)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D17_DC35) and self.cf56 == __o.cf56
    def __hash__(self) -> int:
        return super().__hash__()


class D18:
    @classmethod
    def default(cls, ):
        return lambda: None
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC36(self) -> bool:
        return isinstance(self, D18_DC36)

class D18_DC36(D18, NamedTuple('DC36', [('cf57', Any)])):
    def __dafnystr__(self) -> str:
        return f'D18.DC36({_dafny.string_of(self.cf57)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D18_DC36) and self.cf57 == __o.cf57
    def __hash__(self) -> int:
        return super().__hash__()


class D19:
    @classmethod
    def default(cls, ):
        return lambda: _dafny.Seq({})
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC37(self) -> bool:
        return isinstance(self, D19_DC37)

class D19_DC37(D19, NamedTuple('DC37', [('cf58', Any)])):
    def __dafnystr__(self) -> str:
        return f'D19.DC37({_dafny.string_of(self.cf58)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D19_DC37) and self.cf58 == __o.cf58
    def __hash__(self) -> int:
        return super().__hash__()


class D20:
    @classmethod
    def default(cls, ):
        return lambda: D20_DC39(int(0), False, None, None, False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC39(self) -> bool:
        return isinstance(self, D20_DC39)
    @property
    def is_DC40(self) -> bool:
        return isinstance(self, D20_DC40)
    @property
    def is_DC38(self) -> bool:
        return isinstance(self, D20_DC38)
    @property
    def is_DC41(self) -> bool:
        return isinstance(self, D20_DC41)

class D20_DC39(D20, NamedTuple('DC39', [('cf60', Any), ('cf61', Any), ('cf62', Any), ('cf63', Any), ('cf64', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC39({_dafny.string_of(self.cf60)}, {_dafny.string_of(self.cf61)}, {_dafny.string_of(self.cf62)}, {_dafny.string_of(self.cf63)}, {_dafny.string_of(self.cf64)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC39) and self.cf60 == __o.cf60 and self.cf61 == __o.cf61 and self.cf62 == __o.cf62 and self.cf63 == __o.cf63 and self.cf64 == __o.cf64
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC40(D20, NamedTuple('DC40', [('cf65', Any), ('cf66', Any), ('cf67', Any), ('cf68', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC40({_dafny.string_of(self.cf65)}, {_dafny.string_of(self.cf66)}, {_dafny.string_of(self.cf67)}, {_dafny.string_of(self.cf68)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC40) and self.cf65 == __o.cf65 and self.cf66 == __o.cf66 and self.cf67 == __o.cf67 and self.cf68 == __o.cf68
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC38(D20, NamedTuple('DC38', [('cf59', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC38({_dafny.string_of(self.cf59)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC38) and self.cf59 == __o.cf59
    def __hash__(self) -> int:
        return super().__hash__()

class D20_DC41(D20, NamedTuple('DC41', [('cf69', Any)])):
    def __dafnystr__(self) -> str:
        return f'D20.DC41({_dafny.string_of(self.cf69)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D20_DC41) and self.cf69 == __o.cf69
    def __hash__(self) -> int:
        return super().__hash__()


class D21:
    @classmethod
    def default(cls, ):
        return lambda: D21_DC43(int(0), False, int(0))
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC43(self) -> bool:
        return isinstance(self, D21_DC43)
    @property
    def is_DC42(self) -> bool:
        return isinstance(self, D21_DC42)
    @property
    def is_DC44(self) -> bool:
        return isinstance(self, D21_DC44)

class D21_DC43(D21, NamedTuple('DC43', [('cf71', Any), ('cf72', Any), ('cf73', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC43({_dafny.string_of(self.cf71)}, {_dafny.string_of(self.cf72)}, {_dafny.string_of(self.cf73)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC43) and self.cf71 == __o.cf71 and self.cf72 == __o.cf72 and self.cf73 == __o.cf73
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC42(D21, NamedTuple('DC42', [('cf70', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC42({_dafny.string_of(self.cf70)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC42) and self.cf70 == __o.cf70
    def __hash__(self) -> int:
        return super().__hash__()

class D21_DC44(D21, NamedTuple('DC44', [('cf74', Any)])):
    def __dafnystr__(self) -> str:
        return f'D21.DC44({_dafny.string_of(self.cf74)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D21_DC44) and self.cf74 == __o.cf74
    def __hash__(self) -> int:
        return super().__hash__()


class D22:
    @classmethod
    def default(cls, ):
        return lambda: D22_DC46(False)
    def __ne__(self, __o: object) -> bool:
        return not self.__eq__(__o)
    @property
    def is_DC46(self) -> bool:
        return isinstance(self, D22_DC46)
    @property
    def is_DC47(self) -> bool:
        return isinstance(self, D22_DC47)
    @property
    def is_DC45(self) -> bool:
        return isinstance(self, D22_DC45)
    @property
    def is_DC48(self) -> bool:
        return isinstance(self, D22_DC48)

class D22_DC46(D22, NamedTuple('DC46', [('cf76', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC46({_dafny.string_of(self.cf76)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC46) and self.cf76 == __o.cf76
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC47(D22, NamedTuple('DC47', [('cf77', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC47({_dafny.string_of(self.cf77)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC47) and self.cf77 == __o.cf77
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC45(D22, NamedTuple('DC45', [('cf75', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC45({_dafny.string_of(self.cf75)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC45) and self.cf75 == __o.cf75
    def __hash__(self) -> int:
        return super().__hash__()

class D22_DC48(D22, NamedTuple('DC48', [('cf78', Any)])):
    def __dafnystr__(self) -> str:
        return f'D22.DC48({_dafny.string_of(self.cf78)})'
    def __eq__(self, __o: object) -> bool:
        return isinstance(__o, D22_DC48) and self.cf78 == __o.cf78
    def __hash__(self) -> int:
        return super().__hash__()


class T0:
    pass
    def m3(self, p0, globalState):
        pass


class T1:
    pass
    def m4(self, p0, p1, globalState):
        pass

    def m5(self, globalState):
        pass


class T2:
    pass
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value

class GlobalState:
    def  __init__(self):
        self.f3: _dafny.Array = _dafny.Array(None, 0)
        self.f8: bool = False
        self.f9: int = int(0)
        self._f0: bool = False
        self._f1: _dafny.Array = _dafny.Array(None, 0)
        self._f2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f4: int = int(0)
        self._f5: int = int(0)
        self._f6: int = int(0)
        self._f7: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.GlobalState"
    def ctor__(self, f0, f1, f2, f3, f4, f5, f6, f7, f8, f9):
        (self)._f0 = f0
        (self)._f1 = f1
        (self)._f2 = f2
        (self).f3 = f3
        (self)._f4 = f4
        (self)._f5 = f5
        (self)._f6 = f6
        (self)._f7 = f7
        (self).f8 = f8
        (self).f9 = f9

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

class C0:
    def  __init__(self):
        self._f17: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C0"
    def ctor__(self, f17):
        (self)._f17 = f17

    def fm12(self, p0, p1, globalState):
        return _dafny.MultiSet((_dafny.SeqWithoutIsStrInference([6])) + (_dafny.SeqWithoutIsStrInference([70, -717, len(_dafny.Set({77}))])))

    def fm13(self, p0, globalState):
        source3_ = _dafny.MultiSet([570, 969])
        d_483___mcc_h0_ = source3_
        d_484_cf17_ = d_483___mcc_h0_
        return (_dafny.Set({(self).f17, (self).f17})).intersection((D4_DC10(_dafny.Set({True, (self).f17}))).cf18)

    @property
    def f17(self):
        return self._f17

class C1(T2):
    def  __init__(self):
        self._f12: D0 = D0.default()()
        self._f13: bool = False
        self._f19: bool = False
        self._f20: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C1"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f19, f20, f12, f13):
        (self)._f19 = f19
        (self)._f20 = f20
        (self).f12 = f12
        (self)._f13 = f13

    def fm8(self, globalState):
        source4_ = D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fwsgch")))
        if source4_.is_DC6:
            return D2_DC6()
        elif source4_.is_DC7:
            d_485___mcc_h0_ = source4_.cf11
            d_486___mcc_h1_ = source4_.cf12
            d_487___mcc_h2_ = source4_.cf13
            d_488___mcc_h3_ = source4_.cf14
            d_489_cf14_ = d_488___mcc_h3_
            d_490_cf13_ = d_487___mcc_h2_
            d_491_cf12_ = d_486___mcc_h1_
            d_492_cf11_ = d_485___mcc_h0_
            return D2_DC6()
        elif source4_.is_DC8:
            d_493___mcc_h4_ = source4_.cf15
            d_494___mcc_h5_ = source4_.cf16
            d_495_cf16_ = d_494___mcc_h5_
            d_496_cf15_ = d_493___mcc_h4_
            return D2_DC6()
        elif True:
            d_497___mcc_h6_ = source4_.cf10
            d_498_cf10_ = d_497___mcc_h6_
            if (self).f19:
                return D2_DC6()
            elif True:
                return D2_DC6()

    def fm9(self, p0, p1, globalState):
        return ((0) - (len(_dafny.Map({not((self).f19): 420})))) - (111)

    def fm29(self, p0, globalState):
        def iife43_():
            coll31_ = _dafny.Map()
            compr_31_: _dafny.Seq
            for compr_31_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([451, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anmfxxj"))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_499_i0_ in range(default__.abs(-920))]))}))])])).Elements:
                d_500_v0_: _dafny.Seq = compr_31_
                if (d_500_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([451, len(_dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "anmfxxj"))): len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('u') for d_499_i0_ in range(default__.abs(-920))]))}))])])):
                    coll31_[d_500_v0_] = (self).f13
            return _dafny.Map(coll31_)
        return default__.safeDivisionInt(default__.safeModuloInt(-454, 983), (0) - (len((iife43_()
        ) | (_dafny.Map({_dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet([(self).f19])).cardinality]): len(_dafny.Map({_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_501_i2_ in range(default__.abs(778))]): (self).f19}))})) for d_502_i1_ in range(default__.abs(-466))])), (_dafny.MultiSet([(self).f13])).cardinality, 542, 131, -34]): not((self).f13)})))))

    def fm30(self, p0, globalState):
        return len(_dafny.Map({len(_dafny.Set({(self).f13, (self).f20, (self).f13})): (len(_dafny.SeqWithoutIsStrInference([164])) if (self).f13 else (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qsxjgq")))))}))

    def m9(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.MultiSet = _dafny.MultiSet({})
        d_503_v0_: int
        d_503_v0_ = -644
        if (d_503_v0_) <= (d_503_v0_):
            if (self).f20:
                d_504_v1_: _dafny.Array
                nw100_ = _dafny.Array(False, 4)
                d_504_v1_ = nw100_
                d_505_v2_: _dafny.Set
                d_505_v2_ = _dafny.Set({d_504_v1_, d_504_v1_})
                (globalState).f8 = (len((_dafny.Set({d_504_v1_})) - (d_505_v2_))) < (791)
                d_506_v3_: str
                d_506_v3_ = _dafny.CodePoint('w')
                d_506_v3_ = d_506_v3_
                d_507_v4_: D1
                d_507_v4_ = D1_DC4(d_506_v3_)
                d_508_v5_: _dafny.Map
                d_508_v5_ = _dafny.Map({d_506_v3_: d_507_v4_})
                d_508_v5_ = (d_508_v5_).set(_dafny.CodePoint('y'), d_507_v4_)
                d_509_v6_: _dafny.Seq
                d_509_v6_ = _dafny.SeqWithoutIsStrInference([d_503_v0_])
                d_509_v6_ = d_509_v6_
                d_510_v7_: C0
                nw101_ = C0()
                nw101_.ctor__((self).f19)
                d_510_v7_ = nw101_
                d_511_v8_: _dafny.Seq
                d_511_v8_ = _dafny.SeqWithoutIsStrInference([d_510_v7_])
                (globalState).f8 = ((d_511_v8_) + (_dafny.SeqWithoutIsStrInference([d_510_v7_, d_510_v7_]))) < (d_511_v8_)
            elif True:
                d_512_v9_: _dafny.Set
                d_512_v9_ = _dafny.Set({True, (self).f19})
                rhs83_ = ((self).f20 if (_dafny.Set({(self).f20})).isdisjoint(d_512_v9_) else (self).f20)
                rhs84_ = (self).f19
                rhs85_ = d_503_v0_
                lhs65_ = globalState
                lhs66_ = globalState
                lhs67_ = globalState
                lhs65_.f8 = rhs83_
                lhs66_.f8 = rhs84_
                lhs67_.f9 = rhs85_
                d_513_v10_: C0
                nw102_ = C0()
                nw102_.ctor__(not((len(d_512_v9_)) <= (len(d_512_v9_))))
                d_513_v10_ = nw102_
                d_514_v11_: str
                d_514_v11_ = _dafny.CodePoint('y')
                d_515_v12_: _dafny.Seq
                d_515_v12_ = _dafny.SeqWithoutIsStrInference([default__.fm3(len(d_512_v9_), (d_513_v10_).f17, globalState), d_514_v11_])
                d_516_v13_: _dafny.Seq
                d_516_v13_ = _dafny.SeqWithoutIsStrInference([not((self).f20), not((self).f19)])
                d_517_v14_: _dafny.Map
                d_517_v14_ = _dafny.Map({d_503_v0_: d_503_v0_})
                d_518_v15_: _dafny.Seq
                d_518_v15_ = _dafny.SeqWithoutIsStrInference([d_517_v14_])
                d_519_v16_: _dafny.Seq
                d_519_v16_ = _dafny.SeqWithoutIsStrInference([d_518_v15_, d_518_v15_, _dafny.SeqWithoutIsStrInference([d_517_v14_])])
                d_520_v18_: _dafny.Set
                d_520_v18_ = _dafny.Set({d_503_v0_})
                d_521_v19_: _dafny.Map
                d_521_v19_ = _dafny.Map({d_503_v0_: (d_513_v10_).f17})
                def iife44_():
                    coll32_ = _dafny.Map()
                    compr_32_: int
                    for compr_32_ in (_dafny.Set({len(d_520_v18_)})).Elements:
                        d_522_v17_: int = compr_32_
                        if (d_522_v17_) in (_dafny.Set({len(d_520_v18_)})):
                            coll32_[default__.safeDivisionInt(d_522_v17_, d_503_v0_)] = (self).f20
                    return _dafny.Map(coll32_)
                rhs86_ = ((len((d_515_v12_).set(default__.safeIndex(d_503_v0_, len(d_515_v12_)), d_514_v11_))) - (d_503_v0_)) - (default__.safeDivisionInt(d_503_v0_, len(d_516_v13_)))
                rhs87_ = (self).f13
                rhs88_ = (d_503_v0_) + (d_503_v0_)
                rhs89_ = default__.safeDivisionInt((d_503_v0_) - (len((d_519_v16_)[default__.safeIndex(d_503_v0_, len(d_519_v16_))])), len(_dafny.Map({(self).fm30(iife44_()
                , globalState): d_521_v19_})))
                lhs68_ = globalState
                lhs69_ = globalState
                lhs68_.f9 = rhs86_
                lhs69_.f8 = rhs87_
                r0 = rhs88_
                r0 = rhs89_
                d_523_v20_: _dafny.Map
                d_523_v20_ = _dafny.Map({d_515_v12_: D4_DC10(d_512_v9_)})
                d_524_v21_: D4
                d_524_v21_ = D4_DC10(d_512_v9_)
                d_523_v20_ = (d_523_v20_).set(d_515_v12_, d_524_v21_)
                (globalState).f9 = 857
            d_525_v22_: _dafny.Seq
            d_525_v22_ = _dafny.SeqWithoutIsStrInference([d_503_v0_])
            d_525_v22_ = _dafny.SeqWithoutIsStrInference([d_503_v0_ for d_526_i0_ in range(default__.abs(720))])
            d_527_v23_: _dafny.Array
            nw103_ = _dafny.Array(int(0), 11)
            d_527_v23_ = nw103_
            d_528_v24_: _dafny.MultiSet
            d_528_v24_ = _dafny.MultiSet([d_527_v23_, d_527_v23_, d_527_v23_])
            d_529_v25_: _dafny.MultiSet
            d_529_v25_ = _dafny.MultiSet([(self).f13])
            d_530_v26_: _dafny.Seq
            d_530_v26_ = _dafny.SeqWithoutIsStrInference([d_529_v25_])
            d_531_v27_: _dafny.Seq
            d_531_v27_ = _dafny.SeqWithoutIsStrInference([(d_530_v26_)[default__.safeIndex(d_503_v0_, len(d_530_v26_))]])
            d_532_v28_: D0
            d_532_v28_ = D0_DC1((d_528_v24_).set(d_527_v23_, default__.abs(d_503_v0_)), len(d_531_v27_), (len(d_525_v22_)) * (d_503_v0_))
            d_532_v28_ = d_532_v28_
            d_533_v29_: _dafny.Seq
            d_533_v29_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vfyu"))
            d_533_v29_ = ((d_533_v29_) + (d_533_v29_)) + ((d_533_v29_) + (d_533_v29_))
            d_534_v30_: _dafny.Seq
            d_534_v30_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f20])
            d_535_v31_: _dafny.MultiSet
            d_535_v31_ = _dafny.MultiSet([len(d_534_v30_)])
            d_533_v29_ = default__.fm2((d_532_v28_).cf3, d_535_v31_, _dafny.Map({d_503_v0_: d_503_v0_}), globalState)
        elif True:
            d_536_v32_: C0
            nw104_ = C0()
            nw104_.ctor__((self).f19)
            d_536_v32_ = nw104_
            d_537_v33_: _dafny.Set
            d_537_v33_ = _dafny.Set({(0) - (d_503_v0_), len(default__.fm31(globalState)), d_503_v0_})
            d_538_v34_: _dafny.Seq
            d_538_v34_ = _dafny.SeqWithoutIsStrInference([d_537_v33_, _dafny.Set({d_503_v0_}), d_537_v33_, (D8_DC18(d_537_v33_)).cf33, d_537_v33_])
            d_538_v34_ = default__.fm32(globalState)
            r0 = d_503_v0_
            d_539_v35_: _dafny.MultiSet
            d_539_v35_ = _dafny.MultiSet([(d_536_v32_).f17])
            d_539_v35_ = _dafny.MultiSet([(not(True)) and ((self).f19)])
            d_540_v36_: str
            d_540_v36_ = _dafny.CodePoint('t')
            d_540_v36_ = d_540_v36_
        d_541_v37_: _dafny.Array
        def lambda23_(d_542_i1_):
            return (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ejs"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))

        init12_ = lambda23_
        nw105_ = _dafny.Array(None, 26)
        for i0_12_ in range(nw105_.length(0)):
            nw105_[i0_12_] = init12_(i0_12_)
        d_541_v37_ = nw105_
        d_541_v37_ = d_541_v37_
        d_543_i2_: int
        d_543_i2_ = 0
        with _dafny.label("3"):
            while not (not((self).f13)) or ((self).f20):
                with _dafny.c_label("3"):
                    if (d_543_i2_) >= (100):
                        raise _dafny.Break("3")
                    d_543_i2_ = (d_543_i2_) + (1)
                    d_544_v38_: str
                    d_544_v38_ = _dafny.CodePoint('h')
                    (globalState).f8 = (default__.fm1((self).f19, d_544_v38_, globalState)) and (False)
                    d_545_v39_: _dafny.Seq
                    d_545_v39_ = _dafny.SeqWithoutIsStrInference([(self).f19, (self).f19])
                    d_546_v40_: _dafny.Map
                    d_546_v40_ = _dafny.Map({d_503_v0_: d_545_v39_})
                    d_547_v41_: _dafny.Seq
                    d_547_v41_ = _dafny.SeqWithoutIsStrInference([d_503_v0_, d_503_v0_])
                    d_548_v42_: _dafny.Map
                    d_548_v42_ = _dafny.Map({(self).f19: d_545_v39_})
                    d_549_v43_: _dafny.Array
                    nw106_ = _dafny.Array(None, 17)
                    nw106_[int(0)] = ((d_546_v40_)[d_503_v0_] if (d_503_v0_) in (d_546_v40_) else _dafny.SeqWithoutIsStrInference([(self).f13]))
                    nw106_[int(1)] = d_545_v39_
                    nw106_[int(2)] = d_545_v39_
                    nw106_[int(3)] = d_545_v39_
                    nw106_[int(4)] = d_545_v39_
                    nw106_[int(5)] = d_545_v39_
                    nw106_[int(6)] = d_545_v39_
                    nw106_[int(7)] = ((d_545_v39_).set(default__.safeIndex((0) - (len(d_547_v41_)), len(d_545_v39_)), True) if (self).f13 else d_545_v39_)
                    nw106_[int(8)] = (d_545_v39_) + (d_545_v39_)
                    nw106_[int(9)] = d_545_v39_
                    nw106_[int(10)] = d_545_v39_
                    nw106_[int(11)] = d_545_v39_
                    nw106_[int(12)] = (d_545_v39_) + (d_545_v39_)
                    nw106_[int(13)] = (((d_548_v42_)[(self).f20] if ((self).f20) in (d_548_v42_) else d_545_v39_)).set(default__.safeIndex(d_503_v0_, len(((d_548_v42_)[(self).f20] if ((self).f20) in (d_548_v42_) else d_545_v39_))), (self).f20)
                    nw106_[int(14)] = d_545_v39_
                    nw106_[int(15)] = d_545_v39_
                    nw106_[int(16)] = d_545_v39_
                    d_549_v43_ = nw106_
                    index85_ = default__.safeIndex(989, (d_549_v43_).length(0))
                    (d_549_v43_)[index85_] = d_545_v39_
                    index86_ = default__.safeIndex(989, (d_549_v43_).length(0))
                    (d_549_v43_)[index86_] = d_545_v39_
                    d_550_v44_: _dafny.MultiSet
                    d_550_v44_ = _dafny.MultiSet([d_503_v0_])
                    d_551_v45_: D1
                    d_551_v45_ = D1_DC4(_dafny.CodePoint('v'))
                    d_552_v46_: _dafny.Map
                    d_552_v46_ = _dafny.Map({d_544_v38_: (self).f20})
                    d_553_v47_: _dafny.Set
                    d_553_v47_ = _dafny.Set({d_503_v0_})
                    d_554_v48_: _dafny.Map
                    d_554_v48_ = _dafny.Map({(d_545_v39_).set(default__.safeIndex(len(d_552_v46_), len(d_545_v39_)), (self).f19): d_553_v47_})
                    rhs90_ = (self).f13
                    rhs91_ = (_dafny.MultiSet([(self).fm9(d_551_v45_, ((d_554_v48_)[d_545_v39_] if (d_545_v39_) in (d_554_v48_) else d_553_v47_), globalState)])) | (d_550_v44_)
                    rhs92_ = d_503_v0_
                    lhs70_ = globalState
                    lhs70_.f8 = rhs90_
                    d_550_v44_ = rhs91_
                    d_503_v0_ = rhs92_
                    d_555_v49_: _dafny.Array
                    def lambda24_(d_556_v0_):
                        def lambda25_(d_557_i3_):
                            return (d_556_v0_) == (d_556_v0_)

                        return lambda25_

                    init13_ = lambda24_(d_503_v0_)
                    nw107_ = _dafny.Array(None, 3)
                    for i0_13_ in range(nw107_.length(0)):
                        nw107_[i0_13_] = init13_(i0_13_)
                    d_555_v49_ = nw107_
                    index87_ = default__.safeIndex(602, (d_555_v49_).length(0))
                    (d_555_v49_)[index87_] = (self).f19
                    index88_ = default__.safeIndex(602, (d_555_v49_).length(0))
                    (d_555_v49_)[index88_] = default__.fm1((self).f13, _dafny.CodePoint('c'), globalState)
                    pass
            pass
        d_558_v50_: _dafny.Array
        nw108_ = _dafny.Array(int(0), 18)
        d_558_v50_ = nw108_
        index89_ = default__.safeIndex(648, (d_558_v50_).length(0))
        (d_558_v50_)[index89_] = 115
        index90_ = default__.safeIndex(648, (d_558_v50_).length(0))
        (d_558_v50_)[index90_] = -973
        d_559_v51_: _dafny.Set
        d_559_v51_ = _dafny.Set({_dafny.CodePoint('t')})
        source5_ = ((d_559_v51_ if (self).f20 else d_559_v51_) if (self).f20 else d_559_v51_)
        d_560___mcc_h0_ = source5_
        d_561_cf27_ = d_560___mcc_h0_
        (globalState).f8 = not((self).f19)
        d_562_v52_: str
        d_562_v52_ = _dafny.CodePoint('i')
        d_563_v53_: _dafny.Seq
        d_563_v53_ = _dafny.SeqWithoutIsStrInference([(self).f20, default__.fm1((self).f20, d_562_v52_, globalState), (self).f20])
        d_563_v53_ = (d_563_v53_) + (d_563_v53_)
        (globalState).f8 = (self).f13
        d_564_v54_: _dafny.Map
        d_564_v54_ = _dafny.Map({(d_503_v0_ if (self).f19 else (0) - (d_503_v0_)): (self).f13})
        d_564_v54_ = (d_564_v54_).set((d_558_v50_)[default__.safeIndex(648, (d_558_v50_).length(0))], not((self).f19))
        d_565_v55_: _dafny.Seq
        d_565_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "q"))
        rhs93_ = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_566_i4_ in range(default__.abs(-233))])
        rhs94_ = (d_503_v0_) - (len(d_565_v55_))
        lhs71_ = globalState
        d_565_v55_ = rhs93_
        lhs71_.f9 = rhs94_
        r0 = d_503_v0_
        r1 = p0
        return r0, r1

    def m10(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        d_567_i0_: int
        d_567_i0_ = 0
        with _dafny.label("4"):
            while (self).f19:
                with _dafny.c_label("4"):
                    if (d_567_i0_) >= (100):
                        raise _dafny.Break("4")
                    d_567_i0_ = (d_567_i0_) + (1)
                    d_568_v0_: int
                    d_568_v0_ = 362
                    d_569_v1_: _dafny.Seq
                    d_569_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gt"))
                    d_570_v2_: str
                    d_570_v2_ = _dafny.CodePoint('l')
                    (globalState).f9 = (d_568_v0_) * (len(((d_569_v1_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([D4_DC12((self).f20, (self).f13, (self).f19) for d_571_i1_ in range(default__.abs(-789))])), len(d_569_v1_)), d_570_v2_)) + (d_569_v1_)))
                    d_569_v1_ = ((d_569_v1_) + (d_569_v1_)) + (d_569_v1_)
                    if (self).f20:
                        d_572_v3_: _dafny.Map
                        d_572_v3_ = _dafny.Map({(self).f20: (self).f19})
                        d_572_v3_ = (d_572_v3_).set((self).f13, (self).f19)
                        d_573_v4_: _dafny.Set
                        d_573_v4_ = _dafny.Set({(self).f20})
                        d_574_v5_: _dafny.Map
                        d_574_v5_ = _dafny.Map({d_573_v4_: (self).f19})
                        d_575_v6_: _dafny.Map
                        d_575_v6_ = _dafny.Map({(self).f19: d_574_v5_})
                        d_576_v7_: _dafny.Map
                        d_576_v7_ = _dafny.Map({((self).f19) or ((self).f13): ((d_575_v6_)[(self).f13] if ((self).f13) in (d_575_v6_) else d_574_v5_)})
                        d_576_v7_ = (d_576_v7_).set((self).f20, d_574_v5_)
                        d_577_v8_: _dafny.Map
                        d_577_v8_ = _dafny.Map({False: self.f12})
                        d_578_v9_: _dafny.Map
                        d_578_v9_ = _dafny.Map({-747: True})
                        d_579_v10_: _dafny.Set
                        d_579_v10_ = _dafny.Set({d_568_v0_})
                        d_580_v11_: _dafny.Seq
                        d_580_v11_ = _dafny.SeqWithoutIsStrInference([len(d_579_v10_), -703])
                        d_581_v12_: _dafny.Array
                        nw109_ = _dafny.Array(None, 15)
                        nw109_[int(0)] = d_568_v0_
                        nw109_[int(1)] = -270
                        nw109_[int(2)] = d_568_v0_
                        nw109_[int(3)] = d_568_v0_
                        nw109_[int(4)] = d_568_v0_
                        nw109_[int(5)] = (d_568_v0_ if (self).f13 else len(d_577_v8_))
                        nw109_[int(6)] = (d_568_v0_) * (d_568_v0_)
                        nw109_[int(7)] = len(d_578_v9_)
                        nw109_[int(8)] = d_568_v0_
                        nw109_[int(9)] = d_568_v0_
                        nw109_[int(10)] = (0) - ((d_580_v11_)[default__.safeIndex(d_568_v0_, len(d_580_v11_))])
                        nw109_[int(11)] = d_568_v0_
                        nw109_[int(12)] = d_568_v0_
                        nw109_[int(13)] = d_568_v0_
                        nw109_[int(14)] = d_568_v0_
                        d_581_v12_ = nw109_
                        d_582_v13_: D9
                        d_582_v13_ = D9_DC21(d_581_v12_)
                        d_581_v12_ = (d_582_v13_).cf39
                        d_580_v11_ = (((d_580_v11_) + ((d_580_v11_).set(default__.safeIndex(len(d_569_v1_), len(d_580_v11_)), d_568_v0_))) + (d_580_v11_)).set(default__.safeIndex((0) - (d_568_v0_), len(((d_580_v11_) + ((d_580_v11_).set(default__.safeIndex(len(d_569_v1_), len(d_580_v11_)), d_568_v0_))) + (d_580_v11_))), (len(d_572_v3_) if (self).f19 else d_568_v0_))
                        d_583_v14_: _dafny.Map
                        d_583_v14_ = _dafny.Map({d_579_v10_: False})
                        d_583_v14_ = (d_583_v14_).set(d_579_v10_, (self).f20)
                    elif True:
                        d_584_v15_: D4
                        d_584_v15_ = D4_DC12((self).f20, (self).f13, (self).f19)
                        rhs95_ = d_584_v15_
                        rhs96_ = (self).f19
                        rhs97_ = default__.safeModuloInt((d_568_v0_) + (d_568_v0_), d_568_v0_)
                        lhs72_ = globalState
                        lhs73_ = globalState
                        d_584_v15_ = rhs95_
                        lhs72_.f8 = rhs96_
                        lhs73_.f9 = rhs97_
                        d_585_v16_: _dafny.Map
                        d_585_v16_ = _dafny.Map({len((d_569_v1_) + (d_569_v1_)): default__.safeDivisionInt(d_568_v0_, default__.fm0((self).f19, globalState))})
                        d_586_v17_: D1
                        d_586_v17_ = D1_DC4(d_570_v2_)
                        d_587_v18_: _dafny.Set
                        d_587_v18_ = _dafny.Set({d_568_v0_, d_568_v0_})
                        d_588_v19_: D8
                        d_588_v19_ = D8_DC19(d_569_v1_, (self).fm9(d_586_v17_, d_587_v18_, globalState))
                        d_589_v20_: _dafny.Array
                        nw110_ = _dafny.Array(None, 11)
                        nw110_[int(0)] = (d_569_v1_).set(default__.safeIndex(d_568_v0_, len(d_569_v1_)), d_570_v2_)
                        nw110_[int(1)] = d_569_v1_
                        nw110_[int(2)] = d_569_v1_
                        nw110_[int(3)] = (d_569_v1_) + (((d_588_v19_).cf34).set(default__.safeIndex(d_568_v0_, len((d_588_v19_).cf34)), d_570_v2_))
                        nw110_[int(4)] = d_569_v1_
                        nw110_[int(5)] = d_569_v1_
                        nw110_[int(6)] = d_569_v1_
                        nw110_[int(7)] = d_569_v1_
                        nw110_[int(8)] = d_569_v1_
                        nw110_[int(9)] = (d_569_v1_) + (_dafny.SeqWithoutIsStrInference([d_570_v2_ for d_590_i2_ in range(default__.abs(421))]))
                        nw110_[int(10)] = d_569_v1_
                        d_589_v20_ = nw110_
                        d_591_v21_: _dafny.Array
                        nw111_ = _dafny.Array(int(0), 25)
                        d_591_v21_ = nw111_
                        d_592_v22_: _dafny.Map
                        d_592_v22_ = _dafny.Map({d_591_v21_: d_568_v0_})
                        d_593_v23_: _dafny.Set
                        d_593_v23_ = _dafny.Set({False, (self).f19, False})
                        d_594_v24_: _dafny.Seq
                        d_594_v24_ = _dafny.SeqWithoutIsStrInference([d_569_v1_])
                        d_595_v25_: _dafny.Map
                        d_595_v25_ = _dafny.Map({d_594_v24_: d_569_v1_})
                        rhs98_ = _dafny.Map({d_568_v0_: d_568_v0_})
                        rhs99_ = (d_593_v23_) == (d_593_v23_)
                        rhs100_ = d_589_v20_
                        rhs101_ = d_592_v22_
                        rhs102_ = ((d_595_v25_)[(d_594_v24_) + (d_594_v24_)] if ((d_594_v24_) + (d_594_v24_)) in (d_595_v25_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nor")))
                        lhs74_ = globalState
                        d_585_v16_ = rhs98_
                        lhs74_.f8 = rhs99_
                        d_589_v20_ = rhs100_
                        d_592_v22_ = rhs101_
                        r2 = rhs102_
                        d_596_v26_: _dafny.Array
                        def lambda26_(d_597_i3_):
                            return (self).f13

                        init14_ = lambda26_
                        nw112_ = _dafny.Array(None, 22)
                        for i0_14_ in range(nw112_.length(0)):
                            nw112_[i0_14_] = init14_(i0_14_)
                        d_596_v26_ = nw112_
                        index91_ = default__.safeIndex(233, (d_596_v26_).length(0))
                        (d_596_v26_)[index91_] = not(True)
                        index92_ = default__.safeIndex(233, (d_596_v26_).length(0))
                        (d_596_v26_)[index92_] = (self).f13
                        d_598_v27_: _dafny.MultiSet
                        d_598_v27_ = _dafny.MultiSet([d_568_v0_])
                        d_599_v28_: _dafny.Map
                        d_599_v28_ = _dafny.Map({d_568_v0_: (len(default__.fm2(d_568_v0_, d_598_v27_, d_585_v16_, globalState))) <= (72)})
                        d_599_v28_ = (d_599_v28_).set((0) - ((0) - (default__.safeDivisionInt(d_568_v0_, len(d_593_v23_)))), not((self).f13))
                        d_600_v29_: _dafny.Map
                        d_600_v29_ = _dafny.Map({(d_598_v27_).cardinality: d_570_v2_})
                        d_601_v30_: _dafny.Set
                        d_601_v30_ = _dafny.Set({d_600_v29_})
                        d_602_v31_: _dafny.Seq
                        d_602_v31_ = _dafny.SeqWithoutIsStrInference([d_600_v29_])
                        def iife45_():
                            coll33_ = _dafny.Set()
                            compr_33_: _dafny.Map
                            for compr_33_ in (d_602_v31_).Elements:
                                d_603_v32_: _dafny.Map = compr_33_
                                if (d_603_v32_) in (d_602_v31_):
                                    coll33_ = coll33_.union(_dafny.Set([d_603_v32_]))
                            return _dafny.Set(coll33_)
                        (globalState).f8 = (d_601_v30_) != (iife45_()
                        )
                    (globalState).f8 = (self).f20
                    pass
            pass
        d_604_v33_: int
        d_604_v33_ = -346
        d_605_v34_: _dafny.Seq
        d_605_v34_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ot"))
        d_606_v35_: D2
        d_606_v35_ = D2_DC7(((self).f13 if (self).f19 else (self).f13), d_604_v33_, True, d_605_v34_)
        rhs103_ = (d_604_v33_) - (d_604_v33_)
        rhs104_ = d_606_v35_
        lhs75_ = globalState
        lhs75_.f9 = rhs103_
        d_606_v35_ = rhs104_
        hi4_ = d_604_v33_
        for d_607_i4_ in range(d_604_v33_, hi4_):
            d_608_v36_: _dafny.Array
            nw113_ = _dafny.Array(None, 3)
            nw113_[int(0)] = True
            nw113_[int(1)] = (self).f20
            nw113_[int(2)] = (self).f20
            d_608_v36_ = nw113_
            d_609_v37_: _dafny.MultiSet
            d_609_v37_ = _dafny.MultiSet([d_608_v36_])
            d_610_v38_: _dafny.Map
            d_610_v38_ = _dafny.Map({(d_609_v37_).intersection(d_609_v37_): _dafny.CodePoint('y')})
            d_611_v39_: str
            d_611_v39_ = _dafny.CodePoint('o')
            d_610_v38_ = (d_610_v38_).set(d_609_v37_, d_611_v39_)
            d_612_v40_: _dafny.MultiSet
            d_612_v40_ = _dafny.MultiSet([(self).f13, (self).f13, (self).f20, (self).f20, (self).f20])
            d_613_v41_: _dafny.Seq
            d_613_v41_ = _dafny.SeqWithoutIsStrInference([True, (self).f19, (self).f19, (self).f13, (self).f13])
            (globalState).f9 = ((d_612_v40_).intersection((_dafny.MultiSet(d_613_v41_)) | (d_612_v40_))).cardinality
            (globalState).f8 = not(not((self).f20))
            if (default__.fm33(default__.fm34((self).f13, d_612_v40_, globalState), globalState)).ispropersubset(d_612_v40_):
                r0 = d_611_v39_
                d_614_v42_: _dafny.Array
                nw114_ = _dafny.Array(None, 5)
                nw114_[int(0)] = d_604_v33_
                nw114_[int(1)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgwlycu"))).set(default__.safeIndex(d_607_i4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fgwlycu")))), _dafny.CodePoint('x')))
                nw114_[int(2)] = d_604_v33_
                nw114_[int(3)] = d_607_i4_
                nw114_[int(4)] = d_604_v33_
                d_614_v42_ = nw114_
                index93_ = default__.safeIndex(50, (d_614_v42_).length(0))
                (d_614_v42_)[index93_] = (914) - (d_607_i4_)
                index94_ = default__.safeIndex(50, (d_614_v42_).length(0))
                (d_614_v42_)[index94_] = d_604_v33_
                d_615_v43_: _dafny.Map
                d_615_v43_ = _dafny.Map({(self).f20: d_605_v34_})
                d_616_v44_: _dafny.Set
                d_616_v44_ = _dafny.Set({d_607_i4_})
                d_617_v45_: _dafny.Set
                d_617_v45_ = _dafny.Set({d_616_v44_})
                d_618_v46_: _dafny.Seq
                d_618_v46_ = _dafny.SeqWithoutIsStrInference([(d_614_v42_)[default__.safeIndex(50, (d_614_v42_).length(0))]])
                d_619_v47_: _dafny.MultiSet
                d_619_v47_ = _dafny.MultiSet([len(d_618_v46_)])
                d_620_v48_: _dafny.Map
                d_620_v48_ = _dafny.Map({len(d_616_v44_): d_607_i4_})
                d_615_v43_ = (d_615_v43_).set((d_617_v45_).issubset(d_617_v45_), (default__.fm2(d_604_v33_, d_619_v47_, d_620_v48_, globalState)) + (d_605_v34_))
                d_621_v49_: C0
                nw115_ = C0()
                nw115_.ctor__((self).f19)
                d_621_v49_ = nw115_
                d_622_v50_: D5
                d_622_v50_ = D5_DC13(d_618_v46_)
                d_623_v51_: _dafny.Seq
                d_623_v51_ = _dafny.SeqWithoutIsStrInference([d_622_v50_])
                d_624_v52_: _dafny.MultiSet
                d_624_v52_ = _dafny.MultiSet([default__.fm35((self).f13, (self).f19, globalState), d_623_v51_])
                d_604_v33_ = (0) - (default__.safeDivisionInt(default__.safeDivisionInt(((d_624_v52_)[d_623_v51_] if (d_623_v51_) in (d_624_v52_) else d_604_v33_), ((d_620_v48_)[(d_614_v42_)[default__.safeIndex(50, (d_614_v42_).length(0))]] if ((d_614_v42_)[default__.safeIndex(50, (d_614_v42_).length(0))]) in (d_620_v48_) else d_607_i4_)), d_604_v33_))
            elif True:
                d_625_v53_: C0
                nw116_ = C0()
                nw116_.ctor__((388) != (827))
                d_625_v53_ = nw116_
                d_626_v54_: _dafny.Set
                d_626_v54_ = _dafny.Set({d_604_v33_})
                d_627_v55_: _dafny.Map
                d_627_v55_ = _dafny.Map({(0) - (default__.safeModuloInt(-928, d_607_i4_)): d_626_v54_})
                d_627_v55_ = (d_627_v55_).set(d_607_i4_, d_626_v54_)
                d_628_v56_: _dafny.Seq
                d_628_v56_ = _dafny.SeqWithoutIsStrInference([d_607_i4_])
                d_629_v57_: _dafny.Seq
                d_629_v57_ = _dafny.SeqWithoutIsStrInference([-527, d_604_v33_, len(d_628_v56_)])
                d_630_v58_: _dafny.MultiSet
                d_630_v58_ = _dafny.MultiSet([(d_629_v57_) + (d_628_v56_)])
                d_630_v58_ = ((_dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_607_i4_]), default__.fm36((self).f20, globalState)])).set((_dafny.SeqWithoutIsStrInference([d_604_v33_ for d_631_i5_ in range(default__.abs(143))])).set(default__.safeIndex(d_604_v33_, len(_dafny.SeqWithoutIsStrInference([d_604_v33_ for d_632_i5_ in range(default__.abs(143))]))), d_604_v33_), default__.abs(d_607_i4_))).intersection((d_630_v58_).set(d_628_v56_, default__.abs(d_607_i4_)))
                d_633_v59_: _dafny.Map
                d_633_v59_ = _dafny.Map({default__.safeModuloInt(len(d_613_v41_), d_604_v33_): (d_629_v57_)[default__.safeIndex((self).fm9(D1_DC4(d_611_v39_), d_626_v54_, globalState), len(d_629_v57_))]})
                d_633_v59_ = (d_633_v59_).set(d_604_v33_, (0) - ((d_607_i4_) * (d_604_v33_)))
                index95_ = default__.safeIndex(125, (d_608_v36_).length(0))
                (d_608_v36_)[index95_] = (self).f13
                index96_ = default__.safeIndex(125, (d_608_v36_).length(0))
                (d_608_v36_)[index96_] = default__.fm1((self).f20, d_611_v39_, globalState)
        hi5_ = (d_604_v33_) + (-932)
        for d_634_i6_ in range(d_604_v33_, hi5_):
            (globalState).f8 = (self).f19
            (globalState).f8 = (246) >= (d_604_v33_)
            d_635_v60_: _dafny.Array
            nw117_ = _dafny.Array(False, 24)
            d_635_v60_ = nw117_
            d_636_v61_: str
            d_636_v61_ = _dafny.CodePoint('o')
            index97_ = default__.safeIndex(452, (d_635_v60_).length(0))
            (d_635_v60_)[index97_] = ((self).fm9(D1_DC4(d_636_v61_), default__.fm37(False, d_604_v33_, d_604_v33_, d_634_i6_, globalState), globalState)) <= (d_634_i6_)
            d_637_v62_: _dafny.Map
            d_637_v62_ = _dafny.Map({d_634_i6_: (self).f19})
            index98_ = default__.safeIndex(452, (d_635_v60_).length(0))
            (d_635_v60_)[index98_] = (default__.safeDivisionInt((self).fm30(d_637_v62_, globalState), d_634_i6_)) >= (default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference([d_636_v61_ for d_638_i7_ in range(default__.abs(-842))])).set(default__.safeIndex(d_604_v33_, len(_dafny.SeqWithoutIsStrInference([d_636_v61_ for d_639_i7_ in range(default__.abs(-842))]))), _dafny.CodePoint('w'))), d_634_i6_))
            d_640_v63_: _dafny.MultiSet
            d_640_v63_ = _dafny.MultiSet([d_604_v33_])
            d_641_v64_: _dafny.Map
            d_641_v64_ = _dafny.Map({len(d_605_v34_): d_604_v33_})
            d_642_v65_: _dafny.Map
            d_642_v65_ = _dafny.Map({d_641_v64_: len(d_605_v34_)})
            d_643_v66_: _dafny.Map
            d_643_v66_ = _dafny.Map({default__.fm3(len(_dafny.SeqWithoutIsStrInference([(d_635_v60_)[default__.safeIndex(452, (d_635_v60_).length(0))], (self).f20, (self).f20, (self).f19])), (self).f13, globalState): d_642_v65_})
            (globalState).f9 = ((d_640_v63_).set(d_634_i6_, default__.abs(len((d_642_v65_) | (((d_643_v66_)[d_636_v61_] if (d_636_v61_) in (d_643_v66_) else d_642_v65_)))))).cardinality
        d_644_v67_: _dafny.Array
        def lambda27_(d_645_v33_):
            def lambda28_(d_646_i8_):
                return (d_646_i8_) - (d_645_v33_)

            return lambda28_

        init15_ = lambda27_(d_604_v33_)
        nw118_ = _dafny.Array(None, 11)
        for i0_15_ in range(nw118_.length(0)):
            nw118_[i0_15_] = init15_(i0_15_)
        d_644_v67_ = nw118_
        d_647_v68_: _dafny.Map
        d_647_v68_ = _dafny.Map({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rnrfgxihe")): (self).f13})
        index99_ = default__.safeIndex(26, (d_644_v67_).length(0))
        (d_644_v67_)[index99_] = len(d_647_v68_)
        d_648_v69_: _dafny.Seq
        d_648_v69_ = _dafny.SeqWithoutIsStrInference([d_604_v33_])
        d_649_v70_: _dafny.MultiSet
        d_649_v70_ = _dafny.MultiSet([d_648_v69_, d_648_v69_, d_648_v69_, (d_648_v69_ if False else d_648_v69_)])
        index100_ = default__.safeIndex(26, (d_644_v67_).length(0))
        (d_644_v67_)[index100_] = ((d_649_v70_)[d_648_v69_] if (d_648_v69_) in (d_649_v70_) else d_604_v33_)
        if (self).f13:
            d_650_v71_: _dafny.Set
            d_650_v71_ = _dafny.Set({(self).f20})
            (globalState).f8 = ((0) - (len(d_650_v71_))) < ((d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))])
            d_651_v72_: str
            d_651_v72_ = _dafny.CodePoint('l')
            d_652_v73_: _dafny.Set
            d_652_v73_ = _dafny.Set({d_651_v72_})
            d_653_v74_: _dafny.MultiSet
            d_653_v74_ = _dafny.MultiSet([(d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))], d_604_v33_])
            d_654_v75_: _dafny.MultiSet
            d_654_v75_ = d_653_v74_
            d_655_v76_: D4
            d_655_v76_ = D4_DC12(True, default__.fm1((self).f19, d_651_v72_, globalState), (self).f19)
            d_656_v77_: _dafny.Map
            d_656_v77_ = _dafny.Map({d_604_v33_: d_604_v33_})
            d_657_v78_: C0
            nw119_ = C0()
            nw119_.ctor__((self).f13)
            d_657_v78_ = nw119_
            d_658_v79_: _dafny.Set
            d_658_v79_ = _dafny.Set({d_657_v78_})
            d_659_v80_: _dafny.Seq
            d_659_v80_ = _dafny.SeqWithoutIsStrInference([d_650_v71_])
            d_660_v81_: _dafny.Map
            d_660_v81_ = _dafny.Map({not((d_657_v78_).f17): (self).f20})
            d_661_v82_: _dafny.MultiSet
            d_661_v82_ = _dafny.MultiSet([(self).f19])
            d_662_v83_: D0
            d_662_v83_ = D0_DC2(d_661_v82_, (d_657_v78_).f17, len(d_650_v71_), True)
            d_663_v84_: D8
            d_663_v84_ = D8_DC20(d_662_v83_, (self).f20, default__.fm1((self).f19, d_651_v72_, globalState))
            d_664_v85_: _dafny.Array
            nw120_ = _dafny.Array(None, 29)
            nw120_[int(0)] = ((d_654_v75_)).isdisjoint(d_653_v74_)
            nw120_[int(1)] = True
            nw120_[int(2)] = (self).f19
            nw120_[int(3)] = True
            nw120_[int(4)] = not((d_650_v71_).issubset(d_650_v71_))
            nw120_[int(5)] = (self).f13
            nw120_[int(6)] = (self).f20
            nw120_[int(7)] = (self).f19
            nw120_[int(8)] = (self).f19
            nw120_[int(9)] = (self).f20
            nw120_[int(10)] = (self).f20
            nw120_[int(11)] = (self).f13
            nw120_[int(12)] = not((self).f13)
            nw120_[int(13)] = (not((self).f13)) == ((self).f19)
            def iife46_(_pat_let6_0):
                def iife47_(d_665_dt__update__tmp_h0_):
                    def iife48_(_pat_let7_0):
                        def iife49_(d_666_dt__update_hcf22_h0_):
                            def iife50_(_pat_let8_0):
                                def iife51_(d_667_dt__update_hcf23_h0_):
                                    return D4_DC12(d_666_dt__update_hcf22_h0_, d_667_dt__update_hcf23_h0_, (d_665_dt__update__tmp_h0_).cf24)
                                return iife51_(_pat_let8_0)
                            return iife50_((self).f13)
                        return iife49_(_pat_let7_0)
                    return iife48_((self).f13)
                return iife47_(_pat_let6_0)
            nw120_[int(14)] = (iife46_(d_655_v76_)).cf23
            nw120_[int(15)] = default__.fm1(True, d_651_v72_, globalState)
            nw120_[int(16)] = (self).f19
            nw120_[int(17)] = ((d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))]) == (len(d_656_v77_))
            nw120_[int(18)] = (self).f20
            nw120_[int(19)] = (self).f19
            nw120_[int(20)] = (d_658_v79_) != (_dafny.Set({d_657_v78_}))
            nw120_[int(21)] = not ((self).f20) or (True)
            nw120_[int(22)] = ((d_659_v80_)[default__.safeIndex(-231, len(d_659_v80_))]).isdisjoint(_dafny.Set({(d_657_v78_).f17}))
            nw120_[int(23)] = not ((self).f20) or (default__.fm1((self).f20, _dafny.CodePoint('u'), globalState))
            nw120_[int(24)] = ((d_662_v83_).cf7 if ((d_660_v81_)[(self).f13] if ((self).f13) in (d_660_v81_) else (d_657_v78_).f17) else (d_657_v78_).f17)
            nw120_[int(25)] = not((d_657_v78_).f17)
            nw120_[int(26)] = (d_657_v78_).f17
            nw120_[int(27)] = (self).f20
            nw120_[int(28)] = (d_663_v84_).cf37
            d_664_v85_ = nw120_
            index101_ = default__.safeIndex(141, (d_664_v85_).length(0))
            (d_664_v85_)[index101_] = not((self).f19)
            index102_ = default__.safeIndex(141, (d_664_v85_).length(0))
            index103_ = default__.safeIndex(26, (d_644_v67_).length(0))
            rhs105_ = d_652_v73_
            rhs106_ = (d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))]
            rhs107_ = ((d_657_v78_).f17 if (self).f20 else True)
            rhs108_ = len(d_660_v81_)
            lhs76_ = globalState
            lhs77_ = d_664_v85_
            lhs78_ = default__.safeIndex(141, (d_664_v85_).length(0))
            lhs79_ = d_644_v67_
            lhs80_ = default__.safeIndex(26, (d_644_v67_).length(0))
            d_652_v73_ = rhs105_
            lhs76_.f9 = rhs106_
            lhs77_[lhs78_] = rhs107_
            lhs79_[lhs80_] = rhs108_
            d_604_v33_ = (0) - ((d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))])
            index104_ = default__.safeIndex(141, (d_664_v85_).length(0))
            (d_664_v85_)[index104_] = (d_664_v85_)[default__.safeIndex(141, (d_664_v85_).length(0))]
            d_668_v86_: D9
            d_668_v86_ = D9_DC21(d_644_v67_)
            rhs109_ = d_668_v86_
            rhs110_ = (d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))]
            d_668_v86_ = rhs109_
            d_604_v33_ = rhs110_
        elif True:
            d_604_v33_ = 193
            (globalState).f8 = ((self).f20) or ((self).f13)
            index105_ = default__.safeIndex(26, (d_644_v67_).length(0))
            (d_644_v67_)[index105_] = ((d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))]) + ((d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))])
            d_669_v87_: _dafny.Array
            nw121_ = _dafny.Array(_dafny.Array(None, 0), 10)
            d_669_v87_ = nw121_
            d_670_v88_: str
            d_670_v88_ = _dafny.CodePoint('i')
            d_671_v89_: D1
            d_671_v89_ = D1_DC4(d_670_v88_)
            d_672_v90_: _dafny.Set
            d_672_v90_ = _dafny.Set({(d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))], 415, (d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))]})
            rhs111_ = (self).fm9(d_671_v89_, d_672_v90_, globalState)
            rhs112_ = d_669_v87_
            d_604_v33_ = rhs111_
            d_669_v87_ = rhs112_
            d_673_v91_: _dafny.MultiSet
            d_673_v91_ = _dafny.MultiSet([d_605_v34_, d_605_v34_])
            d_674_v92_: _dafny.Map
            d_674_v92_ = _dafny.Map({(d_673_v91_).cardinality: d_670_v88_})
            d_675_v93_: _dafny.Map
            d_675_v93_ = _dafny.Map({(self).f13: (d_674_v92_) | (d_674_v92_)})
            d_676_v94_: _dafny.Array
            def lambda29_(d_677_i9_):
                return (_dafny.Set({not((self).f20)})).ispropersubset(_dafny.Set({(self).f20, True}))

            init16_ = lambda29_
            nw122_ = _dafny.Array(None, 4)
            for i0_16_ in range(nw122_.length(0)):
                nw122_[i0_16_] = init16_(i0_16_)
            d_676_v94_ = nw122_
            d_678_v95_: _dafny.Seq
            d_678_v95_ = _dafny.SeqWithoutIsStrInference([(self).f19])
            def iife52_():
                coll34_ = _dafny.Map()
                compr_34_: int
                for compr_34_ in _dafny.IntegerRange(160, 441):
                    d_679_v96_: int = compr_34_
                    if ((160) <= (d_679_v96_)) and ((d_679_v96_) < (441)):
                        coll34_[(d_679_v96_) - (d_604_v33_)] = (d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))]
                return _dafny.Map(coll34_)
            rhs113_ = (d_675_v93_).set(not((self).f19), d_674_v92_)
            rhs114_ = (D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "p")), d_676_v94_, (self).f19, len(d_678_v95_))).cf32
            rhs115_ = (d_644_v67_ if ((self).f13 if (self).f13 else default__.fm1(True, default__.fm3((d_644_v67_)[default__.safeIndex(26, (d_644_v67_).length(0))], False, globalState), globalState)) else d_644_v67_)
            rhs116_ = d_676_v94_
            rhs117_ = (not((self).f19)) and ((len(d_605_v34_)) >= (len(iife52_()
            )))
            lhs81_ = globalState
            lhs82_ = globalState
            d_675_v93_ = rhs113_
            lhs81_.f9 = rhs114_
            d_644_v67_ = rhs115_
            d_676_v94_ = rhs116_
            lhs82_.f8 = rhs117_
        r0 = _dafny.CodePoint('y')
        d_680_v97_: _dafny.Set
        d_680_v97_ = _dafny.Set({_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "r")), d_605_v34_})
        r1 = (d_648_v69_)[default__.safeIndex(len(d_680_v97_), len(d_648_v69_))]
        r2 = d_605_v34_
        return r0, r1, r2

    @property
    def f19(self):
        return self._f19
    @property
    def f20(self):
        return self._f20

class C2(T1, T0):
    def  __init__(self):
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C2"
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f11):
        (self)._f11 = f11

    def fm19(self, globalState):
        return _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_681_i0_ in range(default__.abs(129))])

    def fm20(self, p0, p1, globalState):
        if (False) or (not(False)):
            return (-248) * (-552)
        elif True:
            return default__.safeModuloInt(len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.Map({True: 525}))}))])), len(_dafny.Map({-220: (_dafny.MultiSet([True, not(False)])).cardinality})))

    def m4(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: D2 = D2.default()()
        r3: bool = False
        (globalState).f9 = p0
        d_682_v0_: bool
        d_682_v0_ = True
        d_683_v1_: _dafny.MultiSet
        d_683_v1_ = _dafny.MultiSet([True, d_682_v0_, d_682_v0_])
        d_684_v2_: _dafny.Seq
        d_684_v2_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({962: (d_683_v1_).cardinality})])
        d_685_v3_: _dafny.Map
        d_685_v3_ = _dafny.Map({(d_684_v2_)[default__.safeIndex(default__.fm0(d_682_v0_, globalState), len(d_684_v2_))]: p0})
        d_686_v4_: _dafny.Map
        d_686_v4_ = _dafny.Map({p0: p0})
        d_685_v3_ = (d_685_v3_).set(d_686_v4_, 581)
        (globalState).f8 = d_682_v0_
        d_687_v5_: _dafny.Seq
        d_687_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jfnnndul"))
        d_687_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bogafiid"))
        d_688_v6_: D0
        d_688_v6_ = D0_DC0(715)
        pat_let_tv10_ = p0
        def iife53_(_pat_let9_0):
            def iife54_(d_689_dt__update__tmp_h0_):
                def iife55_(_pat_let10_0):
                    def iife56_(d_690_dt__update_hcf0_h0_):
                        return D0_DC0(d_690_dt__update_hcf0_h0_)
                    return iife56_(_pat_let10_0)
                return iife55_(pat_let_tv10_)
            return iife54_(_pat_let9_0)
        d_688_v6_ = iife53_(d_688_v6_)
        d_682_v0_ = (_dafny.CodePoint('r')) in (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnktrnh")))
        d_691_v7_: str
        d_691_v7_ = _dafny.CodePoint('k')
        r0 = d_691_v7_
        r1 = p0
        d_692_v8_: D2
        d_692_v8_ = D2_DC8(p0, d_691_v7_)
        r2 = d_692_v8_
        r3 = d_682_v0_
        return r0, r1, r2, r3

    def m5(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: bool = False
        r2: bool = False
        d_693_v0_: bool
        d_693_v0_ = True
        d_694_v1_: C0
        nw123_ = C0()
        nw123_.ctor__(not(not(d_693_v0_)))
        d_694_v1_ = nw123_
        d_695_v2_: int
        d_695_v2_ = 354
        d_696_v3_: _dafny.Seq
        d_696_v3_ = _dafny.SeqWithoutIsStrInference([(d_694_v1_).f17, d_693_v0_, d_693_v0_])
        hi6_ = default__.safeModuloInt(len(d_696_v3_), d_695_v2_)
        for d_697_i0_ in range(d_695_v2_, hi6_):
            d_698_v4_: _dafny.Array
            nw124_ = _dafny.Array(D1.default()(), 14)
            d_698_v4_ = nw124_
            d_699_v5_: _dafny.Seq
            d_699_v5_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfrdl"))
            d_700_v6_: str
            d_700_v6_ = _dafny.CodePoint('c')
            d_701_v7_: _dafny.Seq
            d_701_v7_ = _dafny.SeqWithoutIsStrInference([len((d_699_v5_).set(default__.safeIndex(592, len(d_699_v5_)), d_700_v6_)), d_697_i0_, d_697_i0_])
            d_702_v8_: _dafny.Map
            d_702_v8_ = _dafny.Map({d_698_v4_: (0) - ((0) - (default__.safeDivisionInt(len(d_701_v7_), d_697_i0_)))})
            d_702_v8_ = _dafny.Map({d_698_v4_: default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnfg"))).set(default__.safeIndex(d_695_v2_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wnfg")))), _dafny.CodePoint('f'))), d_695_v2_)})
            d_703_v9_: _dafny.Array
            nw125_ = _dafny.Array(False, 8)
            d_703_v9_ = nw125_
            index106_ = default__.safeIndex(213, (d_703_v9_).length(0))
            (d_703_v9_)[index106_] = d_693_v0_
            d_704_v10_: _dafny.MultiSet
            d_704_v10_ = _dafny.MultiSet([(d_694_v1_).f17])
            d_705_v11_: _dafny.Map
            d_705_v11_ = _dafny.Map({(d_694_v1_).f17: d_697_i0_})
            d_706_v12_: _dafny.Seq
            d_706_v12_ = _dafny.SeqWithoutIsStrInference([_dafny.MultiSet(d_696_v3_), (d_704_v10_).set((d_694_v1_).f17, default__.abs(((d_705_v11_)[(d_694_v1_).f17] if ((d_694_v1_).f17) in (d_705_v11_) else d_695_v2_)))])
            index107_ = default__.safeIndex(213, (d_703_v9_).length(0))
            (d_703_v9_)[index107_] = (d_704_v10_).ispropersubset((d_706_v12_)[default__.safeIndex(d_697_i0_, len(d_706_v12_))])
            d_707_v13_: C0
            nw126_ = C0()
            nw126_.ctor__((d_695_v2_) != (d_695_v2_))
            d_707_v13_ = nw126_
            d_708_v14_: D2
            d_708_v14_ = D2_DC7(d_693_v0_, (d_704_v10_).cardinality, (d_694_v1_).f17, d_699_v5_)
            rhs118_ = (d_694_v1_).f17
            rhs119_ = (d_708_v14_).cf13
            lhs83_ = globalState
            r2 = rhs118_
            lhs83_.f8 = rhs119_
        d_695_v2_ = 978
        hi7_ = (0) - (d_695_v2_)
        for d_709_i1_ in range((d_695_v2_) * (600), hi7_):
            d_710_v15_: str
            d_710_v15_ = _dafny.CodePoint('o')
            (globalState).f8 = (D4_DC12(d_693_v0_, not((d_694_v1_).f17), not(default__.fm1((d_694_v1_).f17, d_710_v15_, globalState)))).cf24
            r2 = d_693_v0_
            d_711_v16_: _dafny.Set
            d_711_v16_ = _dafny.Set({d_695_v2_, d_695_v2_})
            d_712_v17_: _dafny.MultiSet
            d_712_v17_ = _dafny.MultiSet([d_709_i1_, len(d_711_v16_)])
            d_713_v18_: D5
            d_713_v18_ = D5_DC13(_dafny.SeqWithoutIsStrInference([(0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_714_i3_ in range(default__.abs(135))]))) for d_715_i2_ in range(default__.abs(-209))]))
            d_716_v20_: _dafny.Map
            d_716_v20_ = _dafny.Map({default__.fm0(d_693_v0_, globalState): d_693_v0_})
            d_717_v21_: _dafny.Array
            nw127_ = _dafny.Array(None, 9)
            nw127_[int(0)] = (d_694_v1_).f17
            nw127_[int(1)] = d_693_v0_
            nw127_[int(2)] = (d_694_v1_).f17
            nw127_[int(3)] = (d_694_v1_).f17
            nw127_[int(4)] = d_693_v0_
            nw127_[int(5)] = default__.fm1((d_694_v1_).f17, d_710_v15_, globalState)
            nw127_[int(6)] = d_693_v0_
            nw127_[int(7)] = (d_712_v17_).isdisjoint(_dafny.MultiSet((d_713_v18_).cf25))
            def iife57_():
                coll35_ = _dafny.Map()
                compr_35_: int
                for compr_35_ in _dafny.IntegerRange(714, 223):
                    d_718_v19_: int = compr_35_
                    if ((714) <= (d_718_v19_)) and ((d_718_v19_) < (223)):
                        coll35_[default__.safeDivisionInt(d_718_v19_, d_709_i1_)] = (d_694_v1_).f17
                return _dafny.Map(coll35_)
            nw127_[int(8)] = (iife57_()
            ) != (d_716_v20_)
            d_717_v21_ = nw127_
            d_717_v21_ = d_717_v21_
            source6_ = D2_DC8(d_695_v2_, d_710_v15_)
            if source6_.is_DC6:
                d_719_v22_: _dafny.Array
                nw128_ = _dafny.Array(D0.default()(), 23)
                d_719_v22_ = nw128_
                d_720_v23_: _dafny.Array
                def lambda30_(d_721_i1_):
                    def lambda31_(d_722_i4_):
                        return (d_722_i4_) - (d_721_i1_)

                    return lambda31_

                init17_ = lambda30_(d_709_i1_)
                nw129_ = _dafny.Array(None, 17)
                for i0_17_ in range(nw129_.length(0)):
                    nw129_[i0_17_] = init17_(i0_17_)
                d_720_v23_ = nw129_
                d_723_v24_: _dafny.MultiSet
                d_723_v24_ = _dafny.MultiSet([d_720_v23_])
                d_724_v25_: D0
                d_724_v25_ = D0_DC1(d_723_v24_, (self).fm20((d_694_v1_).f17, d_709_i1_, globalState), -331)
                index108_ = default__.safeIndex(230, (d_719_v22_).length(0))
                (d_719_v22_)[index108_] = d_724_v25_
                index109_ = default__.safeIndex(230, (d_719_v22_).length(0))
                (d_719_v22_)[index109_] = d_724_v25_
                d_725_v26_: _dafny.Seq
                d_725_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aolhi"))
                d_725_v26_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fr"))
                d_726_v27_: _dafny.MultiSet
                d_726_v27_ = _dafny.MultiSet([(d_694_v1_).f17])
                d_727_v28_: _dafny.Set
                d_727_v28_ = _dafny.Set({_dafny.MultiSet([True, (d_694_v1_).f17]), d_726_v27_})
                d_728_v29_: D0
                d_728_v29_ = D0_DC2(d_726_v27_, (d_694_v1_).f17, d_709_i1_, (d_694_v1_).f17)
                d_729_v30_: _dafny.Map
                d_729_v30_ = _dafny.Map({d_695_v2_: _dafny.SeqWithoutIsStrInference([d_710_v15_ for d_730_i5_ in range(default__.abs(-978))])})
                def iife58_():
                    coll36_ = _dafny.Set()
                    compr_36_: int
                    for compr_36_ in (d_729_v30_).keys.Elements:
                        d_731_v31_: int = compr_36_
                        if (d_731_v31_) in (d_729_v30_):
                            coll36_ = coll36_.union(_dafny.Set([default__.safeDivisionInt(d_731_v31_, len(_dafny.Set({938})))]))
                    return _dafny.Set(coll36_)
                rhs120_ = (_dafny.Set({(d_728_v29_).cf4})).issubset(d_727_v28_)
                rhs121_ = (iife58_()
                ) in (default__.fm21((0) - ((self).fm20((d_694_v1_).f17, 203, globalState)), globalState))
                lhs84_ = globalState
                d_693_v0_ = rhs120_
                lhs84_.f8 = rhs121_
                (globalState).f9 = d_709_i1_
            elif source6_.is_DC7:
                d_732___mcc_h0_ = source6_.cf11
                d_733___mcc_h1_ = source6_.cf12
                d_734___mcc_h2_ = source6_.cf13
                d_735___mcc_h3_ = source6_.cf14
                d_736_cf14_ = d_735___mcc_h3_
                d_737_cf13_ = d_734___mcc_h2_
                d_738_cf12_ = d_733___mcc_h1_
                d_739_cf11_ = d_732___mcc_h0_
                d_736_cf14_ = d_736_cf14_
                r1 = d_739_cf11_
                d_740_v32_: C0
                nw130_ = C0()
                nw130_.ctor__((_dafny.Set({d_738_cf12_, -941})).ispropersubset(d_711_v16_))
                d_740_v32_ = nw130_
                (globalState).f9 = default__.safeModuloInt(d_695_v2_, 271)
            elif source6_.is_DC8:
                d_741___mcc_h4_ = source6_.cf15
                d_742___mcc_h5_ = source6_.cf16
                d_743_cf16_ = d_742___mcc_h5_
                d_744_cf15_ = d_741___mcc_h4_
                index110_ = default__.safeIndex(967, (d_717_v21_).length(0))
                (d_717_v21_)[index110_] = d_693_v0_
                index111_ = default__.safeIndex(967, (d_717_v21_).length(0))
                (d_717_v21_)[index111_] = (d_695_v2_) > ((d_744_cf15_ if False else (d_712_v17_).cardinality))
                d_745_v33_: _dafny.Array
                nw131_ = _dafny.Array(None, 3)
                nw131_[int(0)] = d_693_v0_
                nw131_[int(1)] = (d_694_v1_).f17
                nw131_[int(2)] = (d_694_v1_).f17
                d_745_v33_ = nw131_
                d_746_v34_: _dafny.Map
                d_746_v34_ = _dafny.Map({d_744_cf15_: d_745_v33_})
                d_747_v35_: _dafny.Seq
                d_747_v35_ = _dafny.SeqWithoutIsStrInference([d_695_v2_])
                d_748_v36_: _dafny.Seq
                d_748_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ormmod"))
                index112_ = default__.safeIndex(967, (d_717_v21_).length(0))
                (d_717_v21_)[index112_] = ((((d_747_v35_).set(default__.safeIndex(len(d_748_v36_), len(d_747_v35_)), d_709_i1_)).set(default__.safeIndex(d_744_cf15_, len((d_747_v35_).set(default__.safeIndex(len(d_748_v36_), len(d_747_v35_)), d_709_i1_))), d_709_i1_)) < (d_747_v35_) if (d_746_v34_) != (d_746_v34_) else (d_694_v1_).f17)
                d_749_v37_: C0
                nw132_ = C0()
                nw132_.ctor__((d_694_v1_).f17)
                d_749_v37_ = nw132_
                d_750_v38_: D5
                d_750_v38_ = D5_DC14(True)
                d_751_v39_: _dafny.MultiSet
                d_751_v39_ = _dafny.MultiSet([(d_750_v38_).cf26])
                (globalState).f8 = (D0_DC2(d_751_v39_, (d_694_v1_).f17, d_709_i1_, d_693_v0_)).cf5
            elif True:
                d_752___mcc_h6_ = source6_.cf10
                d_753_cf10_ = d_752___mcc_h6_
                d_754_v40_: C0
                nw133_ = C0()
                nw133_.ctor__(False)
                d_754_v40_ = nw133_
                d_710_v15_ = d_710_v15_
                d_755_v41_: _dafny.Array
                nw134_ = _dafny.Array(int(0), 2)
                d_755_v41_ = nw134_
                d_755_v41_ = d_755_v41_
                rhs122_ = (d_694_v1_).f17
                rhs123_ = d_709_i1_
                rhs124_ = (0) - ((0) - ((d_709_i1_) * (len((d_753_cf10_) + (d_753_cf10_)))))
                rhs125_ = d_695_v2_
                lhs85_ = globalState
                lhs86_ = globalState
                lhs87_ = globalState
                lhs85_.f8 = rhs122_
                d_695_v2_ = rhs123_
                lhs86_.f9 = rhs124_
                lhs87_.f9 = rhs125_
        if not(True):
            (globalState).f8 = not((d_695_v2_) <= ((0) - ((d_695_v2_ if d_693_v0_ else d_695_v2_))))
            d_756_v42_: _dafny.MultiSet
            d_756_v42_ = _dafny.MultiSet([(d_694_v1_).f17, (d_694_v1_).f17, d_693_v0_, True])
            d_757_v43_: str
            d_757_v43_ = _dafny.CodePoint('y')
            d_758_v44_: _dafny.MultiSet
            d_758_v44_ = _dafny.MultiSet([d_757_v43_])
            d_759_v45_: _dafny.Seq
            d_759_v45_ = _dafny.SeqWithoutIsStrInference([(D0_DC2(d_756_v42_, (D5_DC14(False)).cf26, d_695_v2_, (d_694_v1_).f17)).cf6, ((d_758_v44_)[default__.fm3(d_695_v2_, (d_694_v1_).f17, globalState)] if (default__.fm3(d_695_v2_, (d_694_v1_).f17, globalState)) in (d_758_v44_) else d_695_v2_)])
            d_760_v46_: _dafny.Set
            d_760_v46_ = _dafny.Set({(0) - (d_695_v2_)})
            d_761_v47_: D0
            d_761_v47_ = D0_DC2(d_756_v42_, d_693_v0_, d_695_v2_, d_693_v0_)
            d_762_v48_: _dafny.Map
            d_762_v48_ = _dafny.Map({(d_761_v47_).cf7: (d_694_v1_).f17})
            d_763_v49_: _dafny.MultiSet
            d_763_v49_ = _dafny.MultiSet([d_695_v2_])
            d_693_v0_ = not(not((_dafny.MultiSet(d_759_v45_)).ispropersubset((_dafny.MultiSet([len(d_760_v46_), len(d_762_v48_), d_695_v2_, 596, 650])) | (d_763_v49_))))
            d_764_v50_: _dafny.Map
            d_764_v50_ = _dafny.Map({d_760_v46_: (d_694_v1_).f17})
            d_764_v50_ = (d_764_v50_).set(d_760_v46_, (d_694_v1_).f17)
            d_765_v51_: _dafny.Set
            d_765_v51_ = _dafny.Set({d_693_v0_})
            d_766_v52_: D4
            d_766_v52_ = D4_DC10(d_765_v51_)
            d_766_v52_ = d_766_v52_
            d_693_v0_ = (d_695_v2_) != (len(_dafny.SeqWithoutIsStrInference([d_757_v43_ for d_767_i6_ in range(default__.abs(-411))])))
        elif True:
            d_768_v53_: str
            d_768_v53_ = _dafny.CodePoint('m')
            d_769_v54_: _dafny.Seq
            d_769_v54_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ehbrla"))
            if (d_768_v53_) not in (d_769_v54_):
                (globalState).f9 = d_695_v2_
                (globalState).f9 = d_695_v2_
                d_693_v0_ = (d_694_v1_).f17
                d_770_v55_: _dafny.Array
                nw135_ = _dafny.Array(False, 23)
                d_770_v55_ = nw135_
                d_771_v56_: _dafny.Seq
                d_771_v56_ = _dafny.SeqWithoutIsStrInference([d_770_v55_])
                d_772_v57_: _dafny.MultiSet
                d_772_v57_ = _dafny.MultiSet([d_768_v53_, d_768_v53_, _dafny.CodePoint('i'), (d_769_v54_)[default__.safeIndex(d_695_v2_, len(d_769_v54_))]])
                d_771_v56_ = (d_771_v56_).set(default__.safeIndex((d_695_v2_) * (((d_772_v57_)[d_768_v53_] if (d_768_v53_) in (d_772_v57_) else d_695_v2_)), len(d_771_v56_)), d_770_v55_)
                d_773_v58_: _dafny.Array
                def lambda32_(d_774_v54_):
                    def lambda33_(d_775_i7_):
                        return D2_DC5(d_774_v54_)

                    return lambda33_

                init18_ = lambda32_(d_769_v54_)
                nw136_ = _dafny.Array(None, 21)
                for i0_18_ in range(nw136_.length(0)):
                    nw136_[i0_18_] = init18_(i0_18_)
                d_773_v58_ = nw136_
                d_773_v58_ = d_773_v58_
            elif True:
                d_776_v59_: _dafny.MultiSet
                d_776_v59_ = _dafny.MultiSet([d_695_v2_])
                d_776_v59_ = d_776_v59_
                d_777_v60_: _dafny.Array
                nw137_ = _dafny.Array(_dafny.Array(None, 0), 26)
                d_777_v60_ = nw137_
                d_778_v61_: _dafny.Array
                def lambda34_(d_779_v2_):
                    def lambda35_(d_780_i8_):
                        return (d_780_i8_) * (d_779_v2_)

                    return lambda35_

                init19_ = lambda34_(d_695_v2_)
                nw138_ = _dafny.Array(None, 14)
                for i0_19_ in range(nw138_.length(0)):
                    nw138_[i0_19_] = init19_(i0_19_)
                d_778_v61_ = nw138_
                index113_ = default__.safeIndex(534, (d_777_v60_).length(0))
                (d_777_v60_)[index113_] = d_778_v61_
                d_781_v62_: _dafny.Array
                nw139_ = _dafny.Array(_dafny.Seq({}), 9)
                d_781_v62_ = nw139_
                d_782_v63_: _dafny.Map
                d_782_v63_ = _dafny.Map({d_781_v62_: d_693_v0_})
                d_783_v64_: _dafny.Seq
                d_783_v64_ = _dafny.SeqWithoutIsStrInference([d_778_v61_])
                d_784_v65_: _dafny.Map
                d_784_v65_ = _dafny.Map({d_778_v61_: d_696_v3_})
                d_785_v66_: _dafny.Set
                d_785_v66_ = _dafny.Set({(d_694_v1_).f17})
                d_786_v67_: _dafny.Set
                d_786_v67_ = _dafny.Set({d_785_v66_, d_785_v66_, d_785_v66_})
                index114_ = default__.safeIndex(534, (d_777_v60_).length(0))
                nw140_ = _dafny.Array(None, 11)
                nw140_[int(0)] = d_695_v2_
                nw140_[int(1)] = (0) - (d_695_v2_)
                nw140_[int(2)] = -427
                nw140_[int(3)] = default__.fm0(((d_782_v63_)[d_781_v62_] if (d_781_v62_) in (d_782_v63_) else (d_694_v1_).f17), globalState)
                nw140_[int(4)] = default__.safeModuloInt(len(d_696_v3_), d_695_v2_)
                nw140_[int(5)] = -64
                nw140_[int(6)] = (0) - ((d_695_v2_) * (d_695_v2_))
                nw140_[int(7)] = d_695_v2_
                nw140_[int(8)] = len((_dafny.Map({(d_783_v64_)[default__.safeIndex(d_695_v2_, len(d_783_v64_))]: d_696_v3_})) | (((d_784_v65_).set(d_778_v61_, _dafny.SeqWithoutIsStrInference([(d_694_v1_).f17]))).set(d_778_v61_, (d_696_v3_).set(default__.safeIndex(d_695_v2_, len(d_696_v3_)), (d_694_v1_).f17))))
                nw140_[int(9)] = len((d_786_v67_).intersection(d_786_v67_))
                nw140_[int(10)] = d_695_v2_
                (d_777_v60_)[index114_] = nw140_
                (globalState).f8 = not((d_694_v1_).f17)
                d_787_v68_: D5
                d_787_v68_ = D5_DC14((d_694_v1_).f17)
                d_788_v69_: _dafny.Map
                d_788_v69_ = _dafny.Map({d_768_v53_: (d_787_v68_).cf26})
                d_789_v70_: _dafny.Map
                d_789_v70_ = _dafny.Map({d_693_v0_: _dafny.Map({d_768_v53_: (d_694_v1_).f17})})
                d_790_v71_: _dafny.Seq
                d_790_v71_ = _dafny.SeqWithoutIsStrInference([d_788_v69_, d_788_v69_])
                d_791_v75_: _dafny.MultiSet
                d_791_v75_ = _dafny.MultiSet([d_768_v53_, d_768_v53_])
                d_792_v77_: D4
                d_792_v77_ = D4_DC12((d_694_v1_).f17, (d_694_v1_).f17, (d_694_v1_).f17)
                d_793_v78_: _dafny.Array
                nw141_ = _dafny.Array(None, 27)
                nw141_[int(0)] = d_788_v69_
                nw141_[int(1)] = ((d_789_v70_)[d_693_v0_] if (d_693_v0_) in (d_789_v70_) else default__.fm22((d_694_v1_).f17, (d_694_v1_).f17, d_768_v53_, default__.fm3(d_695_v2_, (d_694_v1_).f17, globalState), globalState))
                nw141_[int(2)] = (d_788_v69_ if (d_694_v1_).f17 else d_788_v69_)
                nw141_[int(3)] = _dafny.Map({d_768_v53_: (d_694_v1_).f17})
                nw141_[int(4)] = (d_788_v69_) | ((d_790_v71_)[default__.safeIndex(d_695_v2_, len(d_790_v71_))])
                def iife59_():
                    coll37_ = _dafny.Map()
                    compr_37_: str
                    for compr_37_ in (_dafny.SeqWithoutIsStrInference([d_768_v53_])).Elements:
                        d_794_v72_: str = compr_37_
                        if (d_794_v72_) in (_dafny.SeqWithoutIsStrInference([d_768_v53_])):
                            coll37_[d_794_v72_] = True
                    return _dafny.Map(coll37_)
                nw141_[int(5)] = iife59_()
                
                nw141_[int(6)] = (_dafny.Map({d_768_v53_: (d_694_v1_).f17}) if (d_694_v1_).f17 else d_788_v69_)
                nw141_[int(7)] = (_dafny.Map({d_768_v53_: (d_694_v1_).f17}) if d_693_v0_ else _dafny.Map({(d_769_v54_)[default__.safeIndex(d_695_v2_, len(d_769_v54_))]: (d_694_v1_).f17}))
                nw141_[int(8)] = _dafny.Map({d_768_v53_: (d_694_v1_).f17})
                def iife60_():
                    coll38_ = _dafny.Map()
                    compr_38_: str
                    for compr_38_ in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_795_i9_ in range(default__.abs(-618))])).Elements:
                        d_796_v73_: str = compr_38_
                        if (d_796_v73_) in (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_795_i9_ in range(default__.abs(-618))])):
                            coll38_[d_796_v73_] = (d_694_v1_).f17
                    return _dafny.Map(coll38_)
                nw141_[int(9)] = iife60_()
                
                nw141_[int(10)] = (d_788_v69_) | (d_788_v69_)
                nw141_[int(11)] = (_dafny.Map({d_768_v53_: not((d_694_v1_).f17)})) | (d_788_v69_)
                nw141_[int(12)] = (d_788_v69_).set(d_768_v53_, (d_694_v1_).f17)
                nw141_[int(13)] = (d_788_v69_).set(d_768_v53_, not((d_694_v1_).f17))
                nw141_[int(14)] = _dafny.Map({d_768_v53_: (d_694_v1_).f17})
                def iife61_():
                    coll39_ = _dafny.Map()
                    compr_39_: str
                    for compr_39_ in (d_791_v75_).Elements:
                        d_797_v74_: str = compr_39_
                        if (d_797_v74_) in (d_791_v75_):
                            coll39_[d_797_v74_] = (d_694_v1_).f17
                    return _dafny.Map(coll39_)
                nw141_[int(15)] = iife61_()
                
                nw141_[int(16)] = d_788_v69_
                def iife62_():
                    coll40_ = _dafny.Map()
                    compr_40_: str
                    for compr_40_ in (d_788_v69_).keys.Elements:
                        d_798_v76_: str = compr_40_
                        if (d_798_v76_) in (d_788_v69_):
                            coll40_[d_798_v76_] = d_693_v0_
                    return _dafny.Map(coll40_)
                nw141_[int(17)] = iife62_()
                
                nw141_[int(18)] = d_788_v69_
                nw141_[int(19)] = (d_788_v69_) | (d_788_v69_)
                nw141_[int(20)] = d_788_v69_
                nw141_[int(21)] = _dafny.Map({d_768_v53_: (d_792_v77_).cf23})
                nw141_[int(22)] = d_788_v69_
                nw141_[int(23)] = d_788_v69_
                nw141_[int(24)] = d_788_v69_
                nw141_[int(25)] = default__.fm22(d_693_v0_, (d_694_v1_).f17, d_768_v53_, d_768_v53_, globalState)
                nw141_[int(26)] = d_788_v69_
                d_793_v78_ = nw141_
                d_793_v78_ = d_793_v78_
                d_799_v79_: _dafny.Seq
                d_799_v79_ = _dafny.SeqWithoutIsStrInference([d_696_v3_])
                d_799_v79_ = _dafny.SeqWithoutIsStrInference([d_696_v3_])
            d_800_v80_: C0
            nw142_ = C0()
            nw142_.ctor__(d_693_v0_)
            d_800_v80_ = nw142_
            d_801_v81_: D5
            d_801_v81_ = D5_DC14(not(d_693_v0_))
            source7_ = (d_801_v81_ if (d_800_v80_).f17 else d_801_v81_)
            if source7_.is_DC14:
                d_802___mcc_h7_ = source7_.cf26
                d_803_cf26_ = d_802___mcc_h7_
                d_804_v82_: _dafny.Seq
                d_804_v82_ = _dafny.SeqWithoutIsStrInference([d_769_v54_])
                d_805_v83_: _dafny.Map
                d_805_v83_ = _dafny.Map({(d_694_v1_).f17: d_695_v2_})
                d_806_v84_: _dafny.Array
                nw143_ = _dafny.Array(None, 26)
                nw143_[int(0)] = (d_694_v1_).f17
                nw143_[int(1)] = not(d_803_cf26_)
                nw143_[int(2)] = default__.fm1(d_803_cf26_, d_768_v53_, globalState)
                nw143_[int(3)] = (len(d_804_v82_)) < (d_695_v2_)
                nw143_[int(4)] = (d_694_v1_).f17
                nw143_[int(5)] = False
                nw143_[int(6)] = (d_694_v1_).f17
                nw143_[int(7)] = not(True)
                nw143_[int(8)] = (d_694_v1_).f17
                nw143_[int(9)] = (d_695_v2_) < (d_695_v2_)
                nw143_[int(10)] = (d_694_v1_).f17
                nw143_[int(11)] = d_803_cf26_
                nw143_[int(12)] = (d_694_v1_).f17
                nw143_[int(13)] = not ((d_694_v1_).f17) or (False)
                nw143_[int(14)] = ((d_800_v80_).f17) or ((d_694_v1_).f17)
                nw143_[int(15)] = (d_694_v1_).f17
                nw143_[int(16)] = (d_696_v3_) <= (d_696_v3_)
                nw143_[int(17)] = not((d_695_v2_) >= (d_695_v2_))
                nw143_[int(18)] = (d_805_v83_) != (d_805_v83_)
                nw143_[int(19)] = default__.fm1((d_800_v80_).f17, d_768_v53_, globalState)
                nw143_[int(20)] = not((not(d_693_v0_) if (d_800_v80_).f17 else (d_694_v1_).f17))
                nw143_[int(21)] = (d_694_v1_).f17
                nw143_[int(22)] = default__.fm1((d_694_v1_).f17, d_768_v53_, globalState)
                nw143_[int(23)] = not((d_800_v80_).f17)
                nw143_[int(24)] = ((d_694_v1_).f17 if d_693_v0_ else (d_694_v1_).f17)
                nw143_[int(25)] = (d_803_cf26_) and ((d_800_v80_).f17)
                d_806_v84_ = nw143_
                index115_ = default__.safeIndex(589, (d_806_v84_).length(0))
                (d_806_v84_)[index115_] = d_803_cf26_
                index116_ = default__.safeIndex(589, (d_806_v84_).length(0))
                (d_806_v84_)[index116_] = d_693_v0_
                index117_ = default__.safeIndex(589, (d_806_v84_).length(0))
                (d_806_v84_)[index117_] = d_803_cf26_
                d_807_v85_: _dafny.Map
                d_807_v85_ = _dafny.Map({d_695_v2_: (d_800_v80_).f17})
                d_808_v86_: _dafny.MultiSet
                d_808_v86_ = _dafny.MultiSet([(d_694_v1_).f17])
                d_809_v87_: D0
                d_809_v87_ = D0_DC2(d_808_v86_, (d_800_v80_).f17, d_695_v2_, (d_806_v84_)[default__.safeIndex(589, (d_806_v84_).length(0))])
                d_807_v85_ = (d_807_v85_).set(d_695_v2_, (d_809_v87_).cf7)
                d_810_v88_: C0
                nw144_ = C0()
                nw144_.ctor__(((d_806_v84_)[default__.safeIndex(589, (d_806_v84_).length(0))] if (d_800_v80_).f17 else (d_806_v84_)[default__.safeIndex(589, (d_806_v84_).length(0))]))
                d_810_v88_ = nw144_
            elif True:
                d_811___mcc_h8_ = source7_.cf25
                d_812_cf25_ = d_811___mcc_h8_
                d_813_v89_: _dafny.Array
                def lambda36_(d_814_v53_, d_815_v2_):
                    def lambda37_(d_816_i10_):
                        return _dafny.Map({d_814_v53_: d_815_v2_})

                    return lambda37_

                init20_ = lambda36_(d_768_v53_, d_695_v2_)
                nw145_ = _dafny.Array(None, 19)
                for i0_20_ in range(nw145_.length(0)):
                    nw145_[i0_20_] = init20_(i0_20_)
                d_813_v89_ = nw145_
                d_817_v90_: _dafny.Map
                d_817_v90_ = _dafny.Map({_dafny.Map({d_694_v1_: True}): d_768_v53_})
                d_818_v91_: _dafny.Map
                d_818_v91_ = _dafny.Map({d_694_v1_: d_693_v0_})
                d_819_v92_: _dafny.Set
                d_819_v92_ = _dafny.Set({d_768_v53_, ((d_817_v90_)[d_818_v91_] if (d_818_v91_) in (d_817_v90_) else d_768_v53_)})
                d_820_v93_: _dafny.Set
                d_820_v93_ = d_819_v92_
                rhs126_ = d_695_v2_
                rhs127_ = d_813_v89_
                rhs128_ = (d_820_v93_)
                rhs129_ = (d_694_v1_).f17
                lhs88_ = globalState
                lhs88_.f9 = rhs126_
                d_813_v89_ = rhs127_
                d_819_v92_ = rhs128_
                d_693_v0_ = rhs129_
                d_821_v94_: _dafny.Set
                d_821_v94_ = _dafny.Set({(d_694_v1_).f17})
                d_822_v95_: D4
                d_822_v95_ = D4_DC10(d_821_v94_)
                d_823_v96_: int
                out9_: int
                out9_ = (self).m8(d_822_v95_, d_695_v2_, globalState)
                d_823_v96_ = out9_
                d_824_v97_: _dafny.Map
                d_824_v97_ = _dafny.Map({d_768_v53_: 437})
                rhs130_ = False
                rhs131_ = d_824_v97_
                lhs89_ = globalState
                lhs89_.f8 = rhs130_
                d_824_v97_ = rhs131_
                d_825_v98_: _dafny.Set
                d_825_v98_ = _dafny.Set({default__.fm23(d_693_v0_, globalState)})
                d_826_v99_: _dafny.Set
                d_826_v99_ = _dafny.Set({d_695_v2_})
                d_825_v98_ = _dafny.Set({d_826_v99_, _dafny.Set({385, len(default__.fm24(d_769_v54_, (d_694_v1_).f17, (d_694_v1_).f17, globalState))})})
            source8_ = D2_DC6()
            if source8_.is_DC6:
                (globalState).f8 = not((d_695_v2_) >= (d_695_v2_))
                (globalState).f8 = d_693_v0_
                d_827_v100_: _dafny.Array
                def lambda38_(d_828_v53_):
                    def lambda39_(d_829_i11_):
                        return _dafny.SeqWithoutIsStrInference([d_828_v53_ for d_830_i12_ in range(default__.abs(-151))])

                    return lambda39_

                init21_ = lambda38_(d_768_v53_)
                nw146_ = _dafny.Array(None, 1)
                for i0_21_ in range(nw146_.length(0)):
                    nw146_[i0_21_] = init21_(i0_21_)
                d_827_v100_ = nw146_
                d_827_v100_ = d_827_v100_
                d_769_v54_ = (d_769_v54_).set(default__.safeIndex((0) - (d_695_v2_), len(d_769_v54_)), _dafny.CodePoint('a'))
            elif source8_.is_DC7:
                d_831___mcc_h9_ = source8_.cf11
                d_832___mcc_h10_ = source8_.cf12
                d_833___mcc_h11_ = source8_.cf13
                d_834___mcc_h12_ = source8_.cf14
                d_835_cf14_ = d_834___mcc_h12_
                d_836_cf13_ = d_833___mcc_h11_
                d_837_cf12_ = d_832___mcc_h10_
                d_838_cf11_ = d_831___mcc_h9_
                d_839_v101_: D2
                d_839_v101_ = D2_DC8(d_837_cf12_, _dafny.CodePoint('t'))
                d_840_v102_: _dafny.Seq
                d_840_v102_ = _dafny.SeqWithoutIsStrInference([d_837_cf12_])
                d_841_v103_: _dafny.Set
                d_841_v103_ = _dafny.Set({d_695_v2_, d_837_cf12_, (d_839_v101_).cf15, (0) - (len((_dafny.SeqWithoutIsStrInference([d_837_cf12_])) + (d_840_v102_)))})
                d_842_v104_: _dafny.Map
                d_842_v104_ = _dafny.Map({d_693_v0_: d_835_cf14_})
                d_843_v105_: _dafny.MultiSet
                d_843_v105_ = _dafny.MultiSet([((d_842_v104_)[not(True)] if (not(True)) in (d_842_v104_) else d_835_cf14_), _dafny.SeqWithoutIsStrInference([d_768_v53_ for d_844_i13_ in range(default__.abs(314))]), d_769_v54_, (d_835_cf14_).set(default__.safeIndex(d_695_v2_, len(d_835_cf14_)), d_768_v53_)])
                def iife63_():
                    coll41_ = _dafny.Set()
                    compr_41_: int
                    for compr_41_ in (_dafny.MultiSet(d_840_v102_)).Elements:
                        d_845_v106_: int = compr_41_
                        if (d_845_v106_) in (_dafny.MultiSet(d_840_v102_)):
                            coll41_ = coll41_.union(_dafny.Set([(d_845_v106_) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_846_i14_ in range(default__.abs(-93))])))]))
                    return _dafny.Set(coll41_)
                rhs132_ = (d_836_cf13_ if (d_843_v105_).issubset(_dafny.MultiSet([d_835_cf14_, d_835_cf14_, d_769_v54_, d_835_cf14_, d_769_v54_])) else (d_694_v1_).f17)
                rhs133_ = (default__.fm23(d_693_v0_, globalState)) | (iife63_()
                )
                rhs134_ = d_837_cf12_
                d_838_cf11_ = rhs132_
                d_841_v103_ = rhs133_
                d_837_cf12_ = rhs134_
                d_847_v107_: _dafny.Array
                nw147_ = _dafny.Array(None, 21)
                d_847_v107_ = nw147_
                index118_ = default__.safeIndex(75, (d_847_v107_).length(0))
                (d_847_v107_)[index118_] = d_694_v1_
                index119_ = default__.safeIndex(75, (d_847_v107_).length(0))
                (d_847_v107_)[index119_] = d_800_v80_
                d_848_v108_: _dafny.Array
                nw148_ = _dafny.Array(False, 9)
                d_848_v108_ = nw148_
                d_848_v108_ = (d_848_v108_ if (d_800_v80_).f17 else d_848_v108_)
                d_849_v109_: _dafny.Array
                nw149_ = _dafny.Array(None, 7)
                nw149_[int(0)] = (0) - (d_837_cf12_)
                nw149_[int(1)] = d_695_v2_
                nw149_[int(2)] = d_695_v2_
                nw149_[int(3)] = d_837_cf12_
                nw149_[int(4)] = default__.safeDivisionInt((0) - (d_837_cf12_), d_837_cf12_)
                nw149_[int(5)] = d_695_v2_
                nw149_[int(6)] = (self).fm20((d_694_v1_).f17, d_837_cf12_, globalState)
                d_849_v109_ = nw149_
                d_850_v110_: _dafny.MultiSet
                d_850_v110_ = _dafny.MultiSet([(d_694_v1_).f17])
                index120_ = default__.safeIndex(720, (d_849_v109_).length(0))
                (d_849_v109_)[index120_] = ((d_850_v110_).set(d_836_cf13_, default__.abs(d_837_cf12_))).cardinality
                index121_ = default__.safeIndex(720, (d_849_v109_).length(0))
                (d_849_v109_)[index121_] = (0) - (d_695_v2_)
            elif source8_.is_DC8:
                d_851___mcc_h13_ = source8_.cf15
                d_852___mcc_h14_ = source8_.cf16
                d_853_cf16_ = d_852___mcc_h14_
                d_854_cf15_ = d_851___mcc_h13_
                d_855_v111_: _dafny.Array
                nw150_ = _dafny.Array(False, 29)
                d_855_v111_ = nw150_
                d_855_v111_ = d_855_v111_
                d_856_v112_: _dafny.Seq
                d_856_v112_ = _dafny.SeqWithoutIsStrInference([363])
                d_857_v113_: _dafny.Map
                d_857_v113_ = _dafny.Map({d_769_v54_: len(d_769_v54_)})
                (globalState).f8 = (default__.safeDivisionInt(len((d_856_v112_).set(default__.safeIndex(d_695_v2_, len(d_856_v112_)), -597)), len(((d_856_v112_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_694_v1_).f17: (0) - (d_854_cf15_)}), default__.fm25(d_695_v2_, globalState)])), len(d_856_v112_)), ((d_857_v113_)[d_769_v54_] if (d_769_v54_) in (d_857_v113_) else d_695_v2_))).set(default__.safeIndex(d_854_cf15_, len((d_856_v112_).set(default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([_dafny.Map({(d_694_v1_).f17: (0) - (d_854_cf15_)}), default__.fm25(d_695_v2_, globalState)])), len(d_856_v112_)), ((d_857_v113_)[d_769_v54_] if (d_769_v54_) in (d_857_v113_) else d_695_v2_)))), d_695_v2_)))) >= ((d_854_cf15_) + ((0) - (d_854_cf15_)))
                d_858_v114_: C0
                nw151_ = C0()
                nw151_.ctor__((d_694_v1_).f17)
                d_858_v114_ = nw151_
                d_859_v115_: C0
                nw152_ = C0()
                nw152_.ctor__((d_694_v1_).f17)
                d_859_v115_ = nw152_
            elif True:
                d_860___mcc_h15_ = source8_.cf10
                d_861_cf10_ = d_860___mcc_h15_
                d_862_v116_: _dafny.Array
                nw153_ = _dafny.Array(None, 15)
                nw153_[int(0)] = d_769_v54_
                nw153_[int(1)] = d_769_v54_
                nw153_[int(2)] = d_769_v54_
                nw153_[int(3)] = d_861_cf10_
                nw153_[int(4)] = d_861_cf10_
                nw153_[int(5)] = (d_769_v54_).set(default__.safeIndex(d_695_v2_, len(d_769_v54_)), d_768_v53_)
                nw153_[int(6)] = d_861_cf10_
                nw153_[int(7)] = d_861_cf10_
                nw153_[int(8)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "triwllfx"))
                nw153_[int(9)] = d_769_v54_
                nw153_[int(10)] = d_861_cf10_
                nw153_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pnlwkn"))
                nw153_[int(12)] = d_769_v54_
                nw153_[int(13)] = (d_769_v54_) + (d_861_cf10_)
                nw153_[int(14)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v"))
                d_862_v116_ = nw153_
                index122_ = default__.safeIndex(655, (d_862_v116_).length(0))
                (d_862_v116_)[index122_] = d_861_cf10_
                index123_ = default__.safeIndex(655, (d_862_v116_).length(0))
                (d_862_v116_)[index123_] = d_861_cf10_
                d_863_v117_: _dafny.Array
                nw154_ = _dafny.Array(_dafny.Array(None, 0), 11)
                d_863_v117_ = nw154_
                d_864_v118_: _dafny.Array
                nw155_ = _dafny.Array(None, 6)
                nw155_[int(0)] = (d_694_v1_).f17
                nw155_[int(1)] = (d_800_v80_).f17
                nw155_[int(2)] = (d_694_v1_).f17
                nw155_[int(3)] = (d_694_v1_).f17
                nw155_[int(4)] = (d_694_v1_).f17
                nw155_[int(5)] = (d_800_v80_).f17
                d_864_v118_ = nw155_
                index124_ = default__.safeIndex(976, (d_863_v117_).length(0))
                (d_863_v117_)[index124_] = d_864_v118_
                index125_ = default__.safeIndex(976, (d_863_v117_).length(0))
                (d_863_v117_)[index125_] = d_864_v118_
                (globalState).f8 = default__.fm1((d_694_v1_).f17, d_768_v53_, globalState)
                d_865_v119_: _dafny.Seq
                d_865_v119_ = _dafny.SeqWithoutIsStrInference([(d_862_v116_)[default__.safeIndex(655, (d_862_v116_).length(0))], d_769_v54_])
                d_693_v0_ = ((d_865_v119_)[default__.safeIndex(d_695_v2_, len(d_865_v119_))]) < ((d_862_v116_)[default__.safeIndex(655, (d_862_v116_).length(0))])
            d_866_v120_: D2
            d_866_v120_ = D2_DC6()
            d_867_v121_: _dafny.Map
            d_867_v121_ = _dafny.Map({(d_694_v1_).f17: d_866_v120_})
            d_867_v121_ = (d_867_v121_).set((d_694_v1_).f17, d_866_v120_)
        if d_693_v0_:
            d_868_v122_: _dafny.Array
            def lambda40_(d_869_v1_):
                def lambda41_(d_870_i15_):
                    return not((d_869_v1_).f17)

                return lambda41_

            init22_ = lambda40_(d_694_v1_)
            nw156_ = _dafny.Array(None, 16)
            for i0_22_ in range(nw156_.length(0)):
                nw156_[i0_22_] = init22_(i0_22_)
            d_868_v122_ = nw156_
            d_871_v123_: _dafny.Array
            nw157_ = _dafny.Array(None, 3)
            nw157_[int(0)] = d_868_v122_
            nw157_[int(1)] = d_868_v122_
            nw157_[int(2)] = d_868_v122_
            d_871_v123_ = nw157_
            index126_ = default__.safeIndex(926, (d_871_v123_).length(0))
            (d_871_v123_)[index126_] = d_868_v122_
            index127_ = default__.safeIndex(926, (d_871_v123_).length(0))
            (d_871_v123_)[index127_] = d_868_v122_
            d_872_v124_: _dafny.MultiSet
            d_872_v124_ = _dafny.MultiSet([False])
            d_873_v125_: _dafny.Seq
            d_873_v125_ = _dafny.SeqWithoutIsStrInference([(d_872_v124_).cardinality])
            d_873_v125_ = d_873_v125_
            d_693_v0_ = (d_694_v1_).f17
            d_874_v126_: _dafny.Seq
            d_874_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bldcpsr"))
            d_875_v127_: _dafny.Seq
            d_875_v127_ = _dafny.SeqWithoutIsStrInference([d_874_v126_, d_874_v126_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_876_i16_ in range(default__.abs(592))])])
            d_877_v128_: _dafny.Map
            d_877_v128_ = _dafny.Map({((d_873_v125_)[default__.safeIndex(d_695_v2_, len(d_873_v125_))]) + (len(d_875_v127_)): (not(d_693_v0_) if (d_694_v1_).f17 else d_693_v0_)})
            d_878_v129_: _dafny.Set
            d_878_v129_ = _dafny.Set({D4_DC11(d_695_v2_, d_695_v2_, True)})
            rhs135_ = (d_877_v128_) | (d_877_v128_)
            rhs136_ = ((d_694_v1_).f17) == ((d_694_v1_).f17)
            rhs137_ = (d_878_v129_).ispropersubset(d_878_v129_)
            lhs90_ = globalState
            d_877_v128_ = rhs135_
            lhs90_.f8 = rhs136_
            d_693_v0_ = rhs137_
            d_879_v130_: str
            d_879_v130_ = _dafny.CodePoint('l')
            (globalState).f9 = default__.safeDivisionInt((0) - (d_695_v2_), len(_dafny.Map({default__.fm1((d_694_v1_).f17, d_879_v130_, globalState): d_695_v2_})))
        elif True:
            if ((d_693_v0_ if (d_694_v1_).f17 else d_693_v0_)) == (((d_694_v1_).f17 if not(False) else (d_694_v1_).f17)):
                d_695_v2_ = d_695_v2_
                d_880_v131_: _dafny.Seq
                d_880_v131_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ccokdbo"))
                d_881_v132_: _dafny.Map
                d_881_v132_ = _dafny.Map({(d_694_v1_).f17: d_695_v2_})
                d_882_v133_: _dafny.Seq
                d_882_v133_ = _dafny.SeqWithoutIsStrInference([d_695_v2_, d_695_v2_, len(d_881_v132_)])
                d_883_v134_: str
                d_883_v134_ = _dafny.CodePoint('w')
                rhs138_ = (not (d_693_v0_) or ((d_694_v1_).f17)) == ((d_882_v133_) != ((d_882_v133_).set(default__.safeIndex(len(default__.fm26(not((d_694_v1_).f17), len((_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_884_i17_ in range(default__.abs(958))])).set(default__.safeIndex(d_695_v2_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_885_i17_ in range(default__.abs(958))]))), d_883_v134_)), d_693_v0_, globalState)), len(d_882_v133_)), d_695_v2_)))
                rhs139_ = d_695_v2_
                rhs140_ = d_880_v131_
                lhs91_ = globalState
                r2 = rhs138_
                lhs91_.f9 = rhs139_
                d_880_v131_ = rhs140_
                (globalState).f9 = (d_695_v2_) + ((0) - (default__.safeDivisionInt(d_695_v2_, d_695_v2_)))
                d_886_v135_: _dafny.Array
                nw158_ = _dafny.Array(int(0), 9)
                d_886_v135_ = nw158_
                index128_ = default__.safeIndex(128, (d_886_v135_).length(0))
                (d_886_v135_)[index128_] = d_695_v2_
                index129_ = default__.safeIndex(128, (d_886_v135_).length(0))
                (d_886_v135_)[index129_] = len(_dafny.Set({(d_694_v1_).f17, d_693_v0_, not(not((len(d_880_v131_)) < (d_695_v2_)))}))
                d_887_v136_: _dafny.Map
                d_887_v136_ = _dafny.Map({(d_886_v135_)[default__.safeIndex(128, (d_886_v135_).length(0))]: d_693_v0_})
                d_888_v137_: C0
                nw159_ = C0()
                nw159_.ctor__(((d_887_v136_)[(d_886_v135_)[default__.safeIndex(128, (d_886_v135_).length(0))]] if ((d_886_v135_)[default__.safeIndex(128, (d_886_v135_).length(0))]) in (d_887_v136_) else True))
                d_888_v137_ = nw159_
            elif True:
                d_889_v138_: _dafny.Seq
                d_889_v138_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wqtu"))
                d_890_v139_: _dafny.MultiSet
                d_890_v139_ = _dafny.MultiSet([d_695_v2_])
                d_891_v140_: _dafny.Map
                d_891_v140_ = _dafny.Map({d_695_v2_: len(default__.fm27(_dafny.SeqWithoutIsStrInference([d_696_v3_]), globalState))})
                d_889_v138_ = (default__.fm2(d_695_v2_, d_890_v139_, d_891_v140_, globalState)) + (d_889_v138_)
                d_892_v141_: str
                d_892_v141_ = _dafny.CodePoint('u')
                d_889_v138_ = (((d_889_v138_) + (d_889_v138_)) + ((d_889_v138_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_893_i18_ in range(default__.abs(27))])))).set(default__.safeIndex(default__.safeModuloInt(d_695_v2_, d_695_v2_), len(((d_889_v138_) + (d_889_v138_)) + ((d_889_v138_) + (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_894_i18_ in range(default__.abs(27))]))))), d_892_v141_)
                d_895_v142_: _dafny.Array
                def lambda42_(d_896_v1_):
                    def lambda43_(d_897_i19_):
                        return (d_896_v1_).f17

                    return lambda43_

                init23_ = lambda42_(d_694_v1_)
                nw160_ = _dafny.Array(None, 27)
                for i0_23_ in range(nw160_.length(0)):
                    nw160_[i0_23_] = init23_(i0_23_)
                d_895_v142_ = nw160_
                index130_ = default__.safeIndex(50, (d_895_v142_).length(0))
                (d_895_v142_)[index130_] = (d_694_v1_).f17
                d_898_v143_: _dafny.Array
                nw161_ = _dafny.Array(int(0), 12)
                d_898_v143_ = nw161_
                index131_ = default__.safeIndex(275, (d_898_v143_).length(0))
                (d_898_v143_)[index131_] = (d_695_v2_ if (d_694_v1_).f17 else d_695_v2_)
                d_899_v144_: _dafny.Map
                d_899_v144_ = _dafny.Map({d_695_v2_: d_696_v3_})
                index132_ = default__.safeIndex(50, (d_895_v142_).length(0))
                index133_ = default__.safeIndex(275, (d_898_v143_).length(0))
                rhs141_ = (d_696_v3_) <= (((d_899_v144_)[d_695_v2_] if (d_695_v2_) in (d_899_v144_) else d_696_v3_))
                rhs142_ = (d_694_v1_).f17
                rhs143_ = (d_889_v138_) + (d_889_v138_)
                rhs144_ = d_695_v2_
                rhs145_ = d_895_v142_
                lhs92_ = globalState
                lhs93_ = d_895_v142_
                lhs94_ = default__.safeIndex(50, (d_895_v142_).length(0))
                lhs95_ = d_898_v143_
                lhs96_ = default__.safeIndex(275, (d_898_v143_).length(0))
                lhs92_.f8 = rhs141_
                lhs93_[lhs94_] = rhs142_
                d_889_v138_ = rhs143_
                lhs95_[lhs96_] = rhs144_
                d_895_v142_ = rhs145_
                d_900_v145_: _dafny.Seq
                d_900_v145_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_892_v141_ for d_901_i20_ in range(default__.abs(144))]), d_889_v138_])
                d_902_v146_: _dafny.MultiSet
                d_902_v146_ = _dafny.MultiSet([(d_895_v142_)[default__.safeIndex(50, (d_895_v142_).length(0))]])
                d_903_v147_: D4
                d_903_v147_ = D4_DC12((d_900_v145_) < (default__.fm28(_dafny.CodePoint('o'), d_695_v2_, -300, d_693_v0_, globalState)), (d_902_v146_).issubset(d_902_v146_), (d_895_v142_)[default__.safeIndex(50, (d_895_v142_).length(0))])
                d_904_v148_: _dafny.Seq
                d_904_v148_ = _dafny.SeqWithoutIsStrInference([d_695_v2_])
                index134_ = default__.safeIndex(50, (d_895_v142_).length(0))
                index135_ = default__.safeIndex(275, (d_898_v143_).length(0))
                rhs146_ = D4_DC12(d_693_v0_, not(default__.fm1(not((d_895_v142_)[default__.safeIndex(50, (d_895_v142_).length(0))]), d_892_v141_, globalState)), (d_694_v1_).f17)
                rhs147_ = (d_694_v1_).f17
                rhs148_ = (d_904_v148_)[default__.safeIndex(d_695_v2_, len(d_904_v148_))]
                rhs149_ = (((d_898_v143_)[default__.safeIndex(275, (d_898_v143_).length(0))]) - (44)) + (d_695_v2_)
                lhs97_ = d_895_v142_
                lhs98_ = default__.safeIndex(50, (d_895_v142_).length(0))
                lhs99_ = globalState
                lhs100_ = d_898_v143_
                lhs101_ = default__.safeIndex(275, (d_898_v143_).length(0))
                d_903_v147_ = rhs146_
                lhs97_[lhs98_] = rhs147_
                lhs99_.f9 = rhs148_
                lhs100_[lhs101_] = rhs149_
                d_905_v149_: D1
                d_905_v149_ = D1_DC4(d_892_v141_)
                d_906_v150_: _dafny.Map
                d_906_v150_ = _dafny.Map({d_905_v149_: d_695_v2_})
                d_907_v151_: _dafny.Map
                d_907_v151_ = _dafny.Map({(d_694_v1_).f17: (d_906_v150_) | (d_906_v150_)})
                d_907_v151_ = (d_907_v151_).set((d_895_v142_)[default__.safeIndex(50, (d_895_v142_).length(0))], (d_906_v150_) | (d_906_v150_))
            d_908_v152_: D7
            d_908_v152_ = D7_DC16(d_696_v3_)
            d_909_v153_: _dafny.Set
            d_909_v153_ = _dafny.Set({d_693_v0_})
            d_910_v154_: _dafny.Map
            d_910_v154_ = _dafny.Map({(d_696_v3_)[default__.safeIndex(len(d_909_v153_), len(d_696_v3_))]: -789})
            (globalState).f8 = not ((not((d_694_v1_).f17)) in (((d_908_v152_).cf28).set(default__.safeIndex(((d_910_v154_)[not((d_694_v1_).f17)] if (not((d_694_v1_).f17)) in (d_910_v154_) else d_695_v2_), len((d_908_v152_).cf28)), d_693_v0_))) or ((d_695_v2_) == (d_695_v2_))
            if d_693_v0_:
                d_909_v153_ = d_909_v153_
                d_911_v155_: _dafny.Seq
                d_911_v155_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaxdpco"))
                d_912_v156_: D2
                d_912_v156_ = D2_DC5(d_911_v155_)
                d_913_v157_: _dafny.Map
                d_913_v157_ = _dafny.Map({d_911_v155_: d_912_v156_})
                pat_let_tv11_ = d_911_v155_
                def iife64_(_pat_let11_0):
                    def iife65_(d_914_dt__update__tmp_h0_):
                        def iife66_(_pat_let12_0):
                            def iife67_(d_915_dt__update_hcf10_h0_):
                                return D2_DC5(d_915_dt__update_hcf10_h0_)
                            return iife67_(_pat_let12_0)
                        return iife66_(pat_let_tv11_)
                    return iife65_(_pat_let11_0)
                d_913_v157_ = (d_913_v157_).set(d_911_v155_, iife64_(d_912_v156_))
                (globalState).f9 = ((0) - (d_695_v2_)) - (d_695_v2_)
                d_916_v158_: str
                d_916_v158_ = _dafny.CodePoint('y')
                r0 = d_916_v158_
                d_917_v159_: _dafny.Array
                nw162_ = _dafny.Array(int(0), 10)
                d_917_v159_ = nw162_
                index136_ = default__.safeIndex(851, (d_917_v159_).length(0))
                (d_917_v159_)[index136_] = len((d_696_v3_) + (d_696_v3_))
                d_918_v161_: _dafny.Map
                def iife68_():
                    coll42_ = _dafny.Set()
                    compr_42_: str
                    for compr_42_ in (d_911_v155_).Elements:
                        d_919_v160_: str = compr_42_
                        if (d_919_v160_) in (d_911_v155_):
                            coll42_ = coll42_.union(_dafny.Set([d_919_v160_]))
                    return _dafny.Set(coll42_)
                d_918_v161_ = _dafny.Map({iife68_()
                : d_694_v1_})
                d_920_v162_: _dafny.Map
                d_920_v162_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_916_v158_ for d_921_i21_ in range(default__.abs(932))])): d_693_v0_})
                index137_ = default__.safeIndex(851, (d_917_v159_).length(0))
                rhs150_ = ((d_910_v154_)[((d_920_v162_)[(_dafny.MultiSet([d_695_v2_])).cardinality] if ((_dafny.MultiSet([d_695_v2_])).cardinality) in (d_920_v162_) else (d_694_v1_).f17)] if (((d_920_v162_)[(_dafny.MultiSet([d_695_v2_])).cardinality] if ((_dafny.MultiSet([d_695_v2_])).cardinality) in (d_920_v162_) else (d_694_v1_).f17)) in (d_910_v154_) else (self).fm20((d_694_v1_).f17, d_695_v2_, globalState))
                rhs151_ = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_922_i22_ in range(default__.abs(791))]))
                rhs152_ = d_918_v161_
                rhs153_ = d_911_v155_
                rhs154_ = not(d_693_v0_)
                lhs102_ = globalState
                lhs103_ = d_917_v159_
                lhs104_ = default__.safeIndex(851, (d_917_v159_).length(0))
                lhs105_ = globalState
                lhs102_.f9 = rhs150_
                lhs103_[lhs104_] = rhs151_
                d_918_v161_ = rhs152_
                d_911_v155_ = rhs153_
                lhs105_.f8 = rhs154_
            elif True:
                (globalState).f8 = (d_694_v1_).f17
                d_923_v163_: C0
                nw163_ = C0()
                nw163_.ctor__((d_694_v1_).f17)
                d_923_v163_ = nw163_
                d_924_v164_: _dafny.Seq
                d_924_v164_ = _dafny.SeqWithoutIsStrInference([d_695_v2_])
                (globalState).f9 = len(d_924_v164_)
                rhs155_ = d_695_v2_
                rhs156_ = (((d_694_v1_).f17) and ((d_694_v1_).f17)) and (((d_694_v1_).f17 if True else (d_694_v1_).f17))
                d_695_v2_ = rhs155_
                r1 = rhs156_
                d_925_v165_: str
                d_925_v165_ = _dafny.CodePoint('q')
                (globalState).f8 = default__.fm1((d_694_v1_).f17, d_925_v165_, globalState)
            if d_693_v0_:
                d_926_v166_: _dafny.Seq
                d_926_v166_ = _dafny.SeqWithoutIsStrInference([d_695_v2_, d_695_v2_])
                (globalState).f9 = (d_926_v166_)[default__.safeIndex(744, len(d_926_v166_))]
                default__.m0(globalState)
                d_927_v167_: _dafny.MultiSet
                d_927_v167_ = _dafny.MultiSet([d_695_v2_, d_695_v2_])
                d_910_v154_ = (d_910_v154_).set((d_694_v1_).f17, len((d_926_v166_).set(default__.safeIndex(((d_927_v167_)[d_695_v2_] if (d_695_v2_) in (d_927_v167_) else d_695_v2_), len(d_926_v166_)), d_695_v2_)))
                rhs157_ = d_695_v2_
                rhs158_ = ((d_696_v3_) + (d_696_v3_)) + (d_696_v3_)
                d_695_v2_ = rhs157_
                d_696_v3_ = rhs158_
                d_928_v168_: _dafny.Array
                nw164_ = _dafny.Array(_dafny.Set({}), 13)
                d_928_v168_ = nw164_
                nw165_ = _dafny.Array(_dafny.Set({}), 29)
                d_928_v168_ = nw165_
            elif True:
                d_929_v169_: _dafny.Set
                d_929_v169_ = _dafny.Set({d_695_v2_, d_695_v2_})
                d_930_v170_: _dafny.Array
                def lambda44_(d_931_v3_):
                    def lambda45_(d_932_i23_):
                        return (d_931_v3_) + (d_931_v3_)

                    return lambda45_

                init24_ = lambda44_(d_696_v3_)
                nw166_ = _dafny.Array(None, 5)
                for i0_24_ in range(nw166_.length(0)):
                    nw166_[i0_24_] = init24_(i0_24_)
                d_930_v170_ = nw166_
                d_933_v171_: _dafny.Seq
                d_933_v171_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rg"))
                d_934_v172_: D0
                d_934_v172_ = D0_DC0(d_695_v2_)
                d_935_v173_: str
                d_935_v173_ = _dafny.CodePoint('h')
                d_936_v174_: T2
                nw167_ = C1()
                nw167_.ctor__((d_694_v1_).f17, (d_694_v1_).f17, d_934_v172_, default__.fm1((d_694_v1_).f17, d_935_v173_, globalState))
                d_936_v174_ = nw167_
                d_937_v175_: _dafny.Set
                d_937_v175_ = _dafny.Set({d_936_v174_, d_936_v174_})
                d_938_v176_: _dafny.MultiSet
                d_938_v176_ = _dafny.MultiSet([len(d_937_v175_)])
                d_939_v177_: _dafny.Seq
                d_939_v177_ = _dafny.SeqWithoutIsStrInference([len(d_933_v171_)])
                d_940_v178_: _dafny.Seq
                d_940_v178_ = _dafny.SeqWithoutIsStrInference([len(d_939_v177_), d_695_v2_])
                d_941_v179_: _dafny.Array
                nw168_ = _dafny.Array(None, 8)
                nw168_[int(0)] = d_693_v0_
                nw168_[int(1)] = d_693_v0_
                nw168_[int(2)] = d_693_v0_
                nw168_[int(3)] = (d_694_v1_).f17
                nw168_[int(4)] = (d_694_v1_).f17
                nw168_[int(5)] = (d_938_v176_).issubset(_dafny.MultiSet(d_939_v177_))
                nw168_[int(6)] = d_693_v0_
                nw168_[int(7)] = (d_936_v174_).f13
                d_941_v179_ = nw168_
                rhs159_ = d_929_v169_
                rhs160_ = not(True)
                rhs161_ = d_930_v170_
                rhs162_ = d_933_v171_
                rhs163_ = d_941_v179_
                d_929_v169_ = rhs159_
                d_693_v0_ = rhs160_
                d_930_v170_ = rhs161_
                d_933_v171_ = rhs162_
                d_941_v179_ = rhs163_
                d_933_v171_ = d_933_v171_
                d_929_v169_ = d_929_v169_
                r2 = (d_694_v1_).f17
                index138_ = default__.safeIndex(103, (d_941_v179_).length(0))
                (d_941_v179_)[index138_] = d_693_v0_
                index139_ = default__.safeIndex(103, (d_941_v179_).length(0))
                rhs164_ = d_935_v173_
                rhs165_ = ((0) - (d_695_v2_)) >= (d_695_v2_)
                rhs166_ = (d_696_v3_)[default__.safeIndex(d_695_v2_, len(d_696_v3_))]
                lhs106_ = globalState
                lhs107_ = d_941_v179_
                lhs108_ = default__.safeIndex(103, (d_941_v179_).length(0))
                d_935_v173_ = rhs164_
                lhs106_.f8 = rhs165_
                lhs107_[lhs108_] = rhs166_
            d_942_v180_: _dafny.Map
            d_942_v180_ = _dafny.Map({d_695_v2_: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qrta"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "uqxf")))})
            d_942_v180_ = (d_942_v180_).set(d_695_v2_, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('a') for d_943_i24_ in range(default__.abs(826))]))
        d_944_v181_: str
        d_944_v181_ = _dafny.CodePoint('b')
        r0 = d_944_v181_
        r1 = d_693_v0_
        r2 = False
        return r0, r1, r2

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_945_v0_: _dafny.Seq
        d_945_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rewlpo"))
        d_946_v1_: int
        d_946_v1_ = 218
        d_947_v2_: D8
        d_947_v2_ = D8_DC19(d_945_v0_, d_946_v1_)
        d_948_v3_: _dafny.MultiSet
        d_948_v3_ = _dafny.MultiSet([d_947_v2_])
        d_949_v4_: _dafny.Map
        d_949_v4_ = _dafny.Map({559: d_946_v1_})
        d_950_v5_: _dafny.Map
        d_950_v5_ = _dafny.Map({d_946_v1_: ((d_949_v4_)[d_946_v1_] if (d_946_v1_) in (d_949_v4_) else len(d_945_v0_))})
        d_951_v6_: D0
        d_951_v6_ = D0_DC0(len(d_949_v4_))
        d_952_v7_: T2
        nw169_ = C1()
        nw169_.ctor__(p0, p0, d_951_v6_, p0)
        d_952_v7_ = nw169_
        d_953_v8_: _dafny.Map
        d_953_v8_ = _dafny.Map({(d_948_v3_).set(d_947_v2_, default__.abs(-943)): d_952_v7_})
        (globalState).f8 = ((d_953_v8_) | (_dafny.Map({d_948_v3_: d_952_v7_}))) != (d_953_v8_)
        (globalState).f8 = ((d_952_v7_).f13 if p0 else (d_952_v7_).f13)
        hi8_ = default__.safeModuloInt(d_946_v1_, (0) - (d_946_v1_))
        for d_954_i0_ in range((d_946_v1_) - (d_946_v1_), hi8_):
            d_955_v9_: str
            d_955_v9_ = _dafny.CodePoint('n')
            d_956_v10_: _dafny.Seq
            d_956_v10_ = _dafny.SeqWithoutIsStrInference([len(default__.fm37((d_952_v7_).f13, (0) - (d_954_i0_), d_954_i0_, d_946_v1_, globalState)), (default__.fm38((d_952_v7_).f13, (d_952_v7_).f13, d_955_v9_, d_954_i0_, globalState)).cardinality, d_946_v1_, d_954_i0_, len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('c'), d_955_v9_]))])
            d_946_v1_ = default__.safeModuloInt(d_954_i0_, (d_956_v10_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kcbsmgk"))), len(d_956_v10_))])
            d_957_v11_: D1
            d_957_v11_ = D1_DC4(d_955_v9_)
            d_958_v12_: _dafny.MultiSet
            d_958_v12_ = _dafny.MultiSet([(d_952_v7_).f13])
            d_959_v13_: _dafny.Set
            d_959_v13_ = _dafny.Set({(self).fm20((d_952_v7_).f13, ((d_958_v12_)[default__.fm1((d_952_v7_).f13, d_955_v9_, globalState)] if (default__.fm1((d_952_v7_).f13, d_955_v9_, globalState)) in (d_958_v12_) else (d_956_v10_)[default__.safeIndex(d_954_i0_, len(d_956_v10_))]), globalState), d_946_v1_})
            d_960_v14_: C0
            nw170_ = C0()
            nw170_.ctor__(((d_952_v7_).fm9(d_957_v11_, d_959_v13_, globalState)) == (d_946_v1_))
            d_960_v14_ = nw170_
            d_961_v15_: _dafny.Array
            nw171_ = _dafny.Array(int(0), 26)
            d_961_v15_ = nw171_
            d_962_v16_: D9
            d_962_v16_ = D9_DC21(d_961_v15_)
            d_963_v17_: _dafny.Array
            nw172_ = _dafny.Array(None, 24)
            nw172_[int(0)] = d_961_v15_
            nw172_[int(1)] = d_961_v15_
            nw172_[int(2)] = d_961_v15_
            nw172_[int(3)] = d_961_v15_
            nw172_[int(4)] = d_961_v15_
            nw172_[int(5)] = d_961_v15_
            nw172_[int(6)] = (d_962_v16_).cf39
            nw172_[int(7)] = d_961_v15_
            nw172_[int(8)] = d_961_v15_
            nw172_[int(9)] = d_961_v15_
            nw172_[int(10)] = d_961_v15_
            nw172_[int(11)] = d_961_v15_
            nw172_[int(12)] = d_961_v15_
            nw172_[int(13)] = d_961_v15_
            nw172_[int(14)] = d_961_v15_
            nw172_[int(15)] = d_961_v15_
            nw172_[int(16)] = d_961_v15_
            nw172_[int(17)] = d_961_v15_
            nw172_[int(18)] = d_961_v15_
            nw172_[int(19)] = (d_961_v15_ if p0 else d_961_v15_)
            nw172_[int(20)] = d_961_v15_
            nw172_[int(21)] = d_961_v15_
            nw172_[int(22)] = d_961_v15_
            nw172_[int(23)] = d_961_v15_
            d_963_v17_ = nw172_
            index140_ = default__.safeIndex(997, (d_963_v17_).length(0))
            (d_963_v17_)[index140_] = (d_961_v15_ if not((d_960_v14_).f17) else d_961_v15_)
            index141_ = default__.safeIndex(997, (d_963_v17_).length(0))
            (d_963_v17_)[index141_] = d_961_v15_
            d_964_v18_: D2
            d_964_v18_ = D2_DC8(d_946_v1_, default__.fm3(d_954_i0_, (d_952_v7_).f13, globalState))
            d_965_v19_: _dafny.Set
            d_965_v19_ = _dafny.Set({d_955_v9_, (d_964_v18_).cf16, d_955_v9_, d_955_v9_, d_955_v9_})
            d_966_v20_: _dafny.Seq
            d_966_v20_ = _dafny.SeqWithoutIsStrInference([(d_954_i0_) != (len(d_965_v19_))])
            d_967_v21_: _dafny.Array
            nw173_ = _dafny.Array(False, 27)
            d_967_v21_ = nw173_
            d_968_v22_: _dafny.Map
            d_968_v22_ = _dafny.Map({d_954_i0_: d_967_v21_})
            index142_ = default__.safeIndex(325, (d_967_v21_).length(0))
            (d_967_v21_)[index142_] = (d_960_v14_).f17
            d_969_v23_: _dafny.Map
            d_969_v23_ = _dafny.Map({(d_952_v7_).f13: (d_960_v14_).f17})
            index143_ = default__.safeIndex(325, (d_967_v21_).length(0))
            rhs167_ = (0) - ((d_946_v1_) + (((self).fm20((d_952_v7_).f13, (0) - (len((d_969_v23_).set(True, (d_952_v7_).f13))), globalState)) - (d_946_v1_)))
            rhs168_ = (_dafny.SeqWithoutIsStrInference([not(p0)])) + (d_966_v20_)
            rhs169_ = d_968_v22_
            rhs170_ = ((d_960_v14_).f17) and (not((d_966_v20_)[default__.safeIndex(d_954_i0_, len(d_966_v20_))]))
            rhs171_ = default__.fm1((_dafny.SeqWithoutIsStrInference([d_946_v1_])) == (_dafny.SeqWithoutIsStrInference([d_954_i0_])), d_955_v9_, globalState)
            lhs109_ = globalState
            lhs110_ = d_967_v21_
            lhs111_ = default__.safeIndex(325, (d_967_v21_).length(0))
            d_946_v1_ = rhs167_
            d_966_v20_ = rhs168_
            d_968_v22_ = rhs169_
            lhs109_.f8 = rhs170_
            lhs110_[lhs111_] = rhs171_
        d_970_v24_: _dafny.Array
        def lambda46_(d_971_v7_):
            def lambda47_(d_972_i1_):
                return (d_971_v7_).f13

            return lambda47_

        init25_ = lambda46_(d_952_v7_)
        nw174_ = _dafny.Array(None, 17)
        for i0_25_ in range(nw174_.length(0)):
            nw174_[i0_25_] = init25_(i0_25_)
        d_970_v24_ = nw174_
        d_973_v25_: _dafny.Array
        nw175_ = _dafny.Array(None, 2)
        nw175_[int(0)] = d_970_v24_
        nw175_[int(1)] = d_970_v24_
        d_973_v25_ = nw175_
        d_973_v25_ = d_973_v25_
        index144_ = default__.safeIndex(522, (d_970_v24_).length(0))
        (d_970_v24_)[index144_] = (d_952_v7_).f13
        d_974_v26_: _dafny.Map
        d_974_v26_ = _dafny.Map({p0: p0})
        index145_ = default__.safeIndex(522, (d_970_v24_).length(0))
        (d_970_v24_)[index145_] = not (p0) or (((d_974_v26_)[(d_952_v7_).f13] if ((d_952_v7_).f13) in (d_974_v26_) else (d_952_v7_).f13))
        d_975_v27_: str
        d_975_v27_ = _dafny.CodePoint('y')
        d_976_v28_: _dafny.Set
        d_976_v28_ = _dafny.Set({d_975_v27_})
        d_977_v29_: _dafny.Set
        d_977_v29_ = d_976_v28_
        source9_ = d_977_v29_
        d_978___mcc_h0_ = source9_
        d_979_cf27_ = d_978___mcc_h0_
        source10_ = (d_977_v29_ if (d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))] else d_977_v29_)
        d_980___mcc_h1_ = source10_
        d_981_cf27_ = d_980___mcc_h1_
        d_982_v30_: _dafny.MultiSet
        d_982_v30_ = _dafny.MultiSet([False])
        d_983_v31_: _dafny.Seq
        d_983_v31_ = _dafny.SeqWithoutIsStrInference([d_946_v1_])
        d_984_v32_: _dafny.Map
        d_984_v32_ = _dafny.Map({not((d_952_v7_).f13): ((0) - (d_946_v1_) if (d_952_v7_).f13 else (d_983_v31_)[default__.safeIndex(d_946_v1_, len(d_983_v31_))])})
        index146_ = default__.safeIndex(522, (d_970_v24_).length(0))
        rhs172_ = False
        rhs173_ = d_982_v30_
        rhs174_ = default__.safeDivisionInt(default__.safeDivisionInt(d_946_v1_, d_946_v1_), d_946_v1_)
        rhs175_ = ((d_946_v1_) * (len(default__.fm39(249, globalState)))) - (default__.safeDivisionInt(d_946_v1_, 15))
        rhs176_ = ((d_984_v32_) | (d_984_v32_) if (d_952_v7_).f13 else ((d_984_v32_).set(False, d_946_v1_)) | (d_984_v32_))
        lhs112_ = d_970_v24_
        lhs113_ = default__.safeIndex(522, (d_970_v24_).length(0))
        lhs114_ = globalState
        lhs115_ = globalState
        lhs112_[lhs113_] = rhs172_
        d_982_v30_ = rhs173_
        lhs114_.f9 = rhs174_
        lhs115_.f9 = rhs175_
        d_984_v32_ = rhs176_
        d_970_v24_ = (d_970_v24_ if ((d_952_v7_).f13) and ((d_952_v7_).f13) else d_970_v24_)
        d_985_v33_: _dafny.Array
        def lambda48_(d_986_i2_):
            return (d_986_i2_) * (len(_dafny.SeqWithoutIsStrInference([True])))

        init26_ = lambda48_
        nw176_ = _dafny.Array(None, 3)
        for i0_26_ in range(nw176_.length(0)):
            nw176_[i0_26_] = init26_(i0_26_)
        d_985_v33_ = nw176_
        index147_ = default__.safeIndex(137, (d_985_v33_).length(0))
        (d_985_v33_)[index147_] = d_946_v1_
        index148_ = default__.safeIndex(137, (d_985_v33_).length(0))
        (d_985_v33_)[index148_] = d_946_v1_
        index149_ = default__.safeIndex(522, (d_970_v24_).length(0))
        (d_970_v24_)[index149_] = True
        index150_ = default__.safeIndex(522, (d_970_v24_).length(0))
        (d_970_v24_)[index150_] = (d_952_v7_).f13
        d_987_v34_: _dafny.MultiSet
        d_987_v34_ = _dafny.MultiSet([d_946_v1_])
        d_988_v35_: _dafny.Map
        d_988_v35_ = _dafny.Map({192: d_987_v34_})
        d_989_v36_: _dafny.MultiSet
        d_989_v36_ = _dafny.MultiSet([(d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))], (((d_988_v35_)[d_946_v1_] if (d_946_v1_) in (d_988_v35_) else d_987_v34_)).ispropersubset(d_987_v34_), ((d_974_v26_)[(d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))]] if ((d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))]) in (d_974_v26_) else default__.fm1(p0, d_975_v27_, globalState)), (d_952_v7_).f13])
        d_989_v36_ = d_989_v36_
        if True:
            (globalState).f8 = not(p0)
            (globalState).f9 = d_946_v1_
            d_990_v37_: _dafny.Map
            d_990_v37_ = _dafny.Map({d_975_v27_: d_970_v24_})
            d_991_v38_: _dafny.Map
            d_991_v38_ = _dafny.Map({(d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))]: d_990_v37_})
            d_991_v38_ = (d_991_v38_).set((d_952_v7_).f13, d_990_v37_)
            d_992_v39_: _dafny.Array
            nw177_ = _dafny.Array(_dafny.Set({}), 3)
            d_992_v39_ = nw177_
            d_992_v39_ = (d_992_v39_ if (p0) == ((d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))]) else d_992_v39_)
            default__.m0(globalState)
        elif True:
            d_993_v40_: _dafny.Map
            d_993_v40_ = _dafny.Map({_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_946_v1_ for d_994_i3_ in range(default__.abs(577))])): _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_995_i4_ in range(default__.abs(952))])})
            d_996_v41_: _dafny.Map
            d_996_v41_ = _dafny.Map({d_993_v40_: (d_946_v1_) < (d_946_v1_)})
            d_997_v43_: _dafny.Seq
            d_997_v43_ = _dafny.SeqWithoutIsStrInference([True, not(not((d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))])), (d_952_v7_).f13])
            d_998_v44_: _dafny.Map
            d_998_v44_ = _dafny.Map({d_987_v34_: (d_997_v43_)[default__.safeIndex(d_946_v1_, len(d_997_v43_))]})
            d_999_v45_: _dafny.Seq
            d_999_v45_ = _dafny.SeqWithoutIsStrInference([575, d_946_v1_])
            d_1000_v46_: _dafny.Seq
            d_1000_v46_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_999_v45_)[default__.safeIndex(len(d_949_v4_), len(d_999_v45_))])])
            def iife69_():
                coll43_ = _dafny.Map()
                compr_43_: _dafny.MultiSet
                for compr_43_ in (d_998_v44_).keys.Elements:
                    d_1001_v42_: _dafny.MultiSet = compr_43_
                    if (d_1001_v42_) in (d_998_v44_):
                        coll43_[d_1001_v42_] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "j"))
                return _dafny.Map(coll43_)
            d_996_v41_ = (d_996_v41_).set(iife69_()
            , (len(d_999_v45_)) > ((0) - (d_946_v1_)))
            (globalState).f9 = default__.safeModuloInt(len((d_945_v0_) + (d_945_v0_)), d_946_v1_)
            d_1002_v47_: C1
            nw178_ = C1()
            nw178_.ctor__((d_952_v7_).f13, (d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))], d_951_v6_, (d_970_v24_)[default__.safeIndex(522, (d_970_v24_).length(0))])
            d_1002_v47_ = nw178_
            d_1003_v48_: _dafny.MultiSet
            d_1003_v48_ = _dafny.MultiSet([d_1002_v47_, d_1002_v47_, d_1002_v47_])
            d_1004_v49_: _dafny.Map
            d_1004_v49_ = _dafny.Map({((d_1003_v48_).set(d_1002_v47_, default__.abs(d_946_v1_))).issubset(d_1003_v48_): d_946_v1_})
            d_1004_v49_ = (default__.fm25(d_946_v1_, globalState) if True else d_1004_v49_)
            d_1005_v50_: _dafny.Array
            nw179_ = _dafny.Array(int(0), 26)
            d_1005_v50_ = nw179_
            rhs177_ = d_989_v36_
            rhs178_ = d_1005_v50_
            rhs179_ = d_1005_v50_
            rhs180_ = d_1005_v50_
            d_989_v36_ = rhs177_
            d_1005_v50_ = rhs178_
            d_1005_v50_ = rhs179_
            d_1005_v50_ = rhs180_
            d_1006_v52_: _dafny.Array
            def lambda49_(d_1007_v0_):
                def lambda50_(d_1008_i5_):
                    def iife70_():
                        coll44_ = _dafny.Set()
                        compr_44_: _dafny.Seq
                        for compr_44_ in (_dafny.Set({d_1007_v0_})).Elements:
                            d_1009_v51_: _dafny.Seq = compr_44_
                            if (d_1009_v51_) in (_dafny.Set({d_1007_v0_})):
                                coll44_ = coll44_.union(_dafny.Set([d_1009_v51_]))
                        return _dafny.Set(coll44_)
                    return (iife70_()
                    ) - (_dafny.Set({d_1007_v0_}))

                return lambda50_

            init27_ = lambda49_(d_945_v0_)
            nw180_ = _dafny.Array(None, 18)
            for i0_27_ in range(nw180_.length(0)):
                nw180_[i0_27_] = init27_(i0_27_)
            d_1006_v52_ = nw180_
            d_1010_v53_: _dafny.Set
            d_1010_v53_ = _dafny.Set({d_945_v0_, d_945_v0_, d_945_v0_, d_945_v0_})
            index151_ = default__.safeIndex(142, (d_1006_v52_).length(0))
            (d_1006_v52_)[index151_] = (d_1010_v53_).intersection(_dafny.Set({default__.fm2(d_946_v1_, d_987_v34_, d_949_v4_, globalState), d_945_v0_, (self).fm19(globalState), _dafny.SeqWithoutIsStrInference([d_975_v27_ for d_1011_i6_ in range(default__.abs(252))])}))
            d_1012_v54_: _dafny.Map
            d_1012_v54_ = _dafny.Map({(d_1010_v53_) | (d_1010_v53_): d_1010_v53_})
            index152_ = default__.safeIndex(142, (d_1006_v52_).length(0))
            (d_1006_v52_)[index152_] = ((d_1012_v54_)[d_1010_v53_] if (d_1010_v53_) in (d_1012_v54_) else d_1010_v53_)
        r0 = d_946_v1_
        d_1013_v55_: _dafny.Set
        d_1013_v55_ = _dafny.Set({p0, p0})
        d_1014_v56_: _dafny.Set
        d_1014_v56_ = _dafny.Set({(len(d_1013_v55_)) >= (d_946_v1_), default__.fm1(True, _dafny.CodePoint('w'), globalState)})
        r1 = d_1013_v55_
        return r0, r1

    def m8(self, p0, p1, globalState):
        r0: int = int(0)
        d_1015_v0_: _dafny.Array
        def lambda51_(d_1016_i0_):
            return True

        init28_ = lambda51_
        nw181_ = _dafny.Array(None, 18)
        for i0_28_ in range(nw181_.length(0)):
            nw181_[i0_28_] = init28_(i0_28_)
        d_1015_v0_ = nw181_
        d_1017_v1_: bool
        d_1017_v1_ = True
        index153_ = default__.safeIndex(285, (d_1015_v0_).length(0))
        (d_1015_v0_)[index153_] = d_1017_v1_
        index154_ = default__.safeIndex(285, (d_1015_v0_).length(0))
        (d_1015_v0_)[index154_] = d_1017_v1_
        d_1018_v2_: str
        d_1018_v2_ = _dafny.CodePoint('j')
        index155_ = default__.safeIndex(285, (d_1015_v0_).length(0))
        (d_1015_v0_)[index155_] = default__.fm1((d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))], d_1018_v2_, globalState)
        d_1019_i1_: int
        d_1019_i1_ = 0
        with _dafny.label("5"):
            while d_1017_v1_:
                with _dafny.c_label("5"):
                    if (d_1019_i1_) >= (100):
                        raise _dafny.Break("5")
                    d_1019_i1_ = (d_1019_i1_) + (1)
                    d_1020_v3_: _dafny.Seq
                    d_1020_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wlm"))
                    index156_ = default__.safeIndex(285, (d_1015_v0_).length(0))
                    (d_1015_v0_)[index156_] = (default__.fm1((d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))], d_1018_v2_, globalState)) and ((d_1020_v3_) <= (d_1020_v3_))
                    index157_ = default__.safeIndex(285, (d_1015_v0_).length(0))
                    (d_1015_v0_)[index157_] = False
                    r0 = p1
                    d_1021_v4_: _dafny.Array
                    def lambda52_(d_1022_p1_):
                        def lambda53_(d_1023_i2_):
                            return default__.safeDivisionInt(d_1023_i2_, d_1022_p1_)

                        return lambda53_

                    init29_ = lambda52_(p1)
                    nw182_ = _dafny.Array(None, 18)
                    for i0_29_ in range(nw182_.length(0)):
                        nw182_[i0_29_] = init29_(i0_29_)
                    d_1021_v4_ = nw182_
                    index158_ = default__.safeIndex(911, (d_1021_v4_).length(0))
                    (d_1021_v4_)[index158_] = (p1) - (p1)
                    index159_ = default__.safeIndex(911, (d_1021_v4_).length(0))
                    rhs181_ = default__.safeDivisionInt((p1) + (p1), p1)
                    rhs182_ = p1
                    lhs116_ = globalState
                    lhs117_ = d_1021_v4_
                    lhs118_ = default__.safeIndex(911, (d_1021_v4_).length(0))
                    lhs116_.f9 = rhs181_
                    lhs117_[lhs118_] = rhs182_
                    pass
            pass
        d_1024_v5_: D5
        d_1024_v5_ = D5_DC14((d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))])
        d_1025_v6_: D0
        d_1025_v6_ = D0_DC0(p1)
        d_1026_v7_: _dafny.Seq
        d_1026_v7_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1017_v1_])])
        d_1027_v8_: C1
        nw183_ = C1()
        nw183_.ctor__((d_1024_v5_).cf26, (d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))], d_1025_v6_, (len((d_1026_v7_)[default__.safeIndex(803, len(d_1026_v7_))])) == (222))
        d_1027_v8_ = nw183_
        if (d_1027_v8_).f20:
            if (default__.safeDivisionInt(p1, p1)) == (p1):
                d_1028_v9_: _dafny.MultiSet
                d_1028_v9_ = _dafny.MultiSet([(d_1027_v8_).f19])
                d_1029_v10_: D0
                d_1029_v10_ = D0_DC2(d_1028_v9_, (d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))], p1, (d_1027_v8_).f20)
                d_1030_v11_: _dafny.Map
                d_1030_v11_ = _dafny.Map({p1: d_1029_v10_})
                d_1030_v11_ = (d_1030_v11_).set(p1, d_1029_v10_)
                (globalState).f9 = (0) - (p1)
                (globalState).f8 = not(not((d_1027_v8_).f20))
                index160_ = default__.safeIndex(285, (d_1015_v0_).length(0))
                (d_1015_v0_)[index160_] = ((p1) * (p1)) > (p1)
                (globalState).f9 = 749
            elif True:
                d_1031_v12_: _dafny.Seq
                d_1031_v12_ = _dafny.SeqWithoutIsStrInference([not((d_1027_v8_).f20)])
                r0 = (len((d_1031_v12_) + (d_1031_v12_))) * ((0) - (p1))
                d_1032_v13_: _dafny.Seq
                d_1032_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fddf"))
                d_1032_v13_ = ((d_1032_v13_) + (_dafny.SeqWithoutIsStrInference([d_1018_v2_ for d_1033_i3_ in range(default__.abs(441))]))) + ((d_1032_v13_) + (d_1032_v13_))
                d_1032_v13_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "phdrqre")) if (d_1027_v8_).f20 else _dafny.SeqWithoutIsStrInference([d_1018_v2_ for d_1034_i4_ in range(default__.abs(-297))]))) + ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qlro"))) + (_dafny.SeqWithoutIsStrInference([d_1018_v2_ for d_1035_i5_ in range(default__.abs(412))])))
                d_1036_v14_: _dafny.MultiSet
                d_1036_v14_ = _dafny.MultiSet([d_1017_v1_, True])
                d_1037_v15_: C1
                nw184_ = C1()
                nw184_.ctor__((d_1027_v8_).f20, not((d_1036_v14_).issubset(d_1036_v14_)), default__.fm40(_dafny.CodePoint('u'), globalState), (d_1027_v8_).f19)
                d_1037_v15_ = nw184_
                rhs183_ = (0) - ((0) - ((len(d_1032_v13_)) + (p1)))
                rhs184_ = d_1015_v0_
                rhs185_ = d_1037_v15_
                rhs186_ = (d_1037_v15_).f19
                lhs119_ = globalState
                r0 = rhs183_
                d_1015_v0_ = rhs184_
                d_1037_v15_ = rhs185_
                lhs119_.f8 = rhs186_
                (globalState).f9 = ((0) - (default__.fm0((d_1027_v8_).f19, globalState))) - ((p1) + (860))
            rhs187_ = default__.fm1((d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))], d_1018_v2_, globalState)
            rhs188_ = d_1018_v2_
            rhs189_ = (d_1027_v8_).f20
            rhs190_ = default__.fm0((p1) == (-324), globalState)
            lhs120_ = globalState
            lhs121_ = globalState
            d_1017_v1_ = rhs187_
            d_1018_v2_ = rhs188_
            lhs120_.f8 = rhs189_
            lhs121_.f9 = rhs190_
            r0 = (0) - (p1)
            (globalState).f9 = p1
            (globalState).f9 = p1
        elif True:
            d_1026_v7_ = d_1026_v7_
            d_1038_v16_: _dafny.Array
            nw185_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
            d_1038_v16_ = nw185_
            d_1039_v17_: _dafny.Seq
            d_1039_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ncq"))
            index161_ = default__.safeIndex(539, (d_1038_v16_).length(0))
            (d_1038_v16_)[index161_] = d_1039_v17_
            d_1040_v18_: _dafny.Array
            nw186_ = _dafny.Array(_dafny.Array(None, 0), 22)
            d_1040_v18_ = nw186_
            d_1041_v19_: _dafny.Array
            nw187_ = _dafny.Array(int(0), 1)
            d_1041_v19_ = nw187_
            d_1042_v20_: _dafny.Array
            nw188_ = _dafny.Array(None, 10)
            nw188_[int(0)] = d_1041_v19_
            nw188_[int(1)] = d_1041_v19_
            nw188_[int(2)] = d_1041_v19_
            nw188_[int(3)] = d_1041_v19_
            nw188_[int(4)] = d_1041_v19_
            nw188_[int(5)] = d_1041_v19_
            nw188_[int(6)] = d_1041_v19_
            nw188_[int(7)] = d_1041_v19_
            nw188_[int(8)] = d_1041_v19_
            nw188_[int(9)] = d_1041_v19_
            d_1042_v20_ = nw188_
            index162_ = default__.safeIndex(422, (d_1040_v18_).length(0))
            (d_1040_v18_)[index162_] = d_1042_v20_
            index163_ = default__.safeIndex(539, (d_1038_v16_).length(0))
            index164_ = default__.safeIndex(422, (d_1040_v18_).length(0))
            rhs191_ = d_1039_v17_
            rhs192_ = d_1042_v20_
            lhs122_ = d_1038_v16_
            lhs123_ = default__.safeIndex(539, (d_1038_v16_).length(0))
            lhs124_ = d_1040_v18_
            lhs125_ = default__.safeIndex(422, (d_1040_v18_).length(0))
            lhs122_[lhs123_] = rhs191_
            lhs124_[lhs125_] = rhs192_
            d_1043_v21_: _dafny.MultiSet
            d_1043_v21_ = _dafny.MultiSet([True, True])
            d_1044_v22_: D0
            d_1044_v22_ = D0_DC2(d_1043_v21_, (d_1027_v8_).f19, p1, d_1017_v1_)
            index165_ = default__.safeIndex(784, (d_1041_v19_).length(0))
            (d_1041_v19_)[index165_] = (d_1044_v22_).cf6
            index166_ = default__.safeIndex(784, (d_1041_v19_).length(0))
            (d_1041_v19_)[index166_] = (p1) + ((0) - ((len(default__.fm23(False, globalState))) + (p1)))
            (globalState).f9 = p1
            d_1018_v2_ = d_1018_v2_
        rhs193_ = (d_1017_v1_ if not((d_1015_v0_)[default__.safeIndex(285, (d_1015_v0_).length(0))]) else (d_1027_v8_).f20)
        rhs194_ = d_1017_v1_
        rhs195_ = p1
        rhs196_ = p1
        d_1017_v1_ = rhs193_
        d_1017_v1_ = rhs194_
        r0 = rhs195_
        r0 = rhs196_
        r0 = ((0) - ((0) - ((p1) - (p1))) if (d_1027_v8_).f20 else p1)
        return r0


class C3(T0):
    def  __init__(self):
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self.f21: bool = False
        self._f22: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C3"
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f21, f22, f11):
        (self).f21 = f21
        (self)._f22 = f22
        (self)._f11 = f11

    def fm47(self, p0, globalState):
        return (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vedhv"))))

    def fm48(self, p0, p1, p2, globalState):
        def iife71_():
            coll45_ = _dafny.Map()
            compr_45_: _dafny.Seq
            for compr_45_ in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, (self).f22])])).Elements:
                d_1045_v0_: _dafny.Seq = compr_45_
                if (d_1045_v0_) in (_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([False, (self).f22])])):
                    coll45_[d_1045_v0_] = len(_dafny.Set({False, False}))
            return _dafny.Map(coll45_)
        return ((_dafny.Set({416})) - (_dafny.Set({len(iife71_()
        )}))).intersection((_dafny.Set({807, 539})) - (_dafny.Set({859})))

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        if p0:
            d_1046_v0_: _dafny.Array
            def lambda54_(d_1047_i0_):
                return self.f21

            init30_ = lambda54_
            nw189_ = _dafny.Array(None, 1)
            for i0_30_ in range(nw189_.length(0)):
                nw189_[i0_30_] = init30_(i0_30_)
            d_1046_v0_ = nw189_
            d_1046_v0_ = d_1046_v0_
            (globalState).f8 = self.f21
            d_1048_v1_: C2
            nw190_ = C2()
            nw190_.ctor__(((self).f11 if self.f21 else (self).f11))
            d_1048_v1_ = nw190_
            d_1049_v2_: int
            d_1049_v2_ = 708
            d_1050_v3_: _dafny.Seq
            d_1050_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vilgvrgs"))
            rhs197_ = default__.safeDivisionInt(d_1049_v2_, (873) - (d_1049_v2_))
            rhs198_ = not(not((default__.safeDivisionInt(len(d_1050_v3_), d_1049_v2_)) < (d_1049_v2_)))
            lhs126_ = globalState
            r0 = rhs197_
            lhs126_.f8 = rhs198_
            (globalState).f9 = (d_1049_v2_ if p0 else d_1049_v2_)
        elif True:
            (globalState).f8 = self.f21
            d_1051_v4_: C0
            nw191_ = C0()
            nw191_.ctor__(self.f21)
            d_1051_v4_ = nw191_
            d_1052_v5_: int
            d_1052_v5_ = 180
            d_1053_v6_: _dafny.Seq
            d_1053_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiuask"))
            d_1054_v7_: _dafny.Map
            d_1054_v7_ = _dafny.Map({(d_1051_v4_).f17: len(d_1053_v6_)})
            d_1055_v8_: D8
            d_1055_v8_ = D8_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "s")), d_1052_v5_)
            d_1056_v9_: C2
            nw192_ = C2()
            nw192_.ctor__((self).f11)
            d_1056_v9_ = nw192_
            d_1057_v10_: _dafny.Set
            d_1057_v10_ = _dafny.Set({d_1056_v9_})
            d_1058_v11_: _dafny.MultiSet
            d_1058_v11_ = _dafny.MultiSet([(self).f22])
            d_1059_v12_: _dafny.Array
            nw193_ = _dafny.Array(None, 26)
            nw193_[int(0)] = default__.fm0((self).f22, globalState)
            nw193_[int(1)] = d_1052_v5_
            nw193_[int(2)] = d_1052_v5_
            nw193_[int(3)] = d_1052_v5_
            nw193_[int(4)] = 470
            nw193_[int(5)] = d_1052_v5_
            nw193_[int(6)] = d_1052_v5_
            nw193_[int(7)] = -750
            nw193_[int(8)] = d_1052_v5_
            nw193_[int(9)] = (len(d_1053_v6_)) - (d_1052_v5_)
            nw193_[int(10)] = d_1052_v5_
            nw193_[int(11)] = d_1052_v5_
            nw193_[int(12)] = d_1052_v5_
            nw193_[int(13)] = 680
            nw193_[int(14)] = 991
            nw193_[int(15)] = ((d_1054_v7_)[self.f21] if (self.f21) in (d_1054_v7_) else d_1052_v5_)
            nw193_[int(16)] = ((D4_DC11(d_1052_v5_, (0) - ((d_1055_v8_).cf35), not((self).f22))).cf20) - (len(d_1057_v10_))
            nw193_[int(17)] = len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1060_i1_ in range(default__.abs(-340))]))
            nw193_[int(18)] = d_1052_v5_
            nw193_[int(19)] = d_1052_v5_
            nw193_[int(20)] = (len(_dafny.Set({d_1051_v4_})) if (self).f22 else (d_1058_v11_).cardinality)
            nw193_[int(21)] = d_1052_v5_
            nw193_[int(22)] = d_1052_v5_
            nw193_[int(23)] = d_1052_v5_
            nw193_[int(24)] = d_1052_v5_
            nw193_[int(25)] = d_1052_v5_
            d_1059_v12_ = nw193_
            index167_ = default__.safeIndex(598, (d_1059_v12_).length(0))
            (d_1059_v12_)[index167_] = (d_1056_v9_).fm20(not(self.f21), d_1052_v5_, globalState)
            index168_ = default__.safeIndex(598, (d_1059_v12_).length(0))
            rhs199_ = (self).f22
            rhs200_ = (0) - (default__.safeDivisionInt(d_1052_v5_, 0))
            lhs127_ = globalState
            lhs128_ = d_1059_v12_
            lhs129_ = default__.safeIndex(598, (d_1059_v12_).length(0))
            lhs127_.f8 = rhs199_
            lhs128_[lhs129_] = rhs200_
            d_1061_v13_: _dafny.Array
            def lambda55_(d_1062_p0_):
                def lambda56_(d_1063_i2_):
                    return d_1062_p0_

                return lambda56_

            init31_ = lambda55_(p0)
            nw194_ = _dafny.Array(None, 5)
            for i0_31_ in range(nw194_.length(0)):
                nw194_[i0_31_] = init31_(i0_31_)
            d_1061_v13_ = nw194_
            d_1064_v14_: _dafny.Array
            nw195_ = _dafny.Array(None, 25)
            nw195_[int(0)] = d_1061_v13_
            nw195_[int(1)] = d_1061_v13_
            nw195_[int(2)] = d_1061_v13_
            nw195_[int(3)] = d_1061_v13_
            nw195_[int(4)] = d_1061_v13_
            nw195_[int(5)] = d_1061_v13_
            nw195_[int(6)] = d_1061_v13_
            nw195_[int(7)] = d_1061_v13_
            nw195_[int(8)] = d_1061_v13_
            nw195_[int(9)] = d_1061_v13_
            nw195_[int(10)] = d_1061_v13_
            nw195_[int(11)] = d_1061_v13_
            nw195_[int(12)] = d_1061_v13_
            nw195_[int(13)] = d_1061_v13_
            nw195_[int(14)] = d_1061_v13_
            nw195_[int(15)] = d_1061_v13_
            nw195_[int(16)] = d_1061_v13_
            nw195_[int(17)] = d_1061_v13_
            nw195_[int(18)] = d_1061_v13_
            nw195_[int(19)] = d_1061_v13_
            nw195_[int(20)] = d_1061_v13_
            nw195_[int(21)] = d_1061_v13_
            nw195_[int(22)] = d_1061_v13_
            nw195_[int(23)] = d_1061_v13_
            nw195_[int(24)] = d_1061_v13_
            d_1064_v14_ = nw195_
            index169_ = default__.safeIndex(302, (d_1064_v14_).length(0))
            (d_1064_v14_)[index169_] = d_1061_v13_
            index170_ = default__.safeIndex(302, (d_1064_v14_).length(0))
            (d_1064_v14_)[index170_] = d_1061_v13_
            index171_ = default__.safeIndex(598, (d_1059_v12_).length(0))
            (d_1059_v12_)[index171_] = ((d_1059_v12_)[default__.safeIndex(598, (d_1059_v12_).length(0))]) * (((d_1059_v12_)[default__.safeIndex(598, (d_1059_v12_).length(0))]) * ((d_1059_v12_)[default__.safeIndex(598, (d_1059_v12_).length(0))]))
        d_1065_v15_: int
        d_1065_v15_ = 989
        d_1066_v16_: D0
        d_1066_v16_ = D0_DC0(d_1065_v15_)
        source11_ = d_1066_v16_
        if source11_.is_DC1:
            d_1067___mcc_h0_ = source11_.cf1
            d_1068___mcc_h1_ = source11_.cf2
            d_1069___mcc_h2_ = source11_.cf3
            d_1070_cf3_ = d_1069___mcc_h2_
            d_1071_cf2_ = d_1068___mcc_h1_
            d_1072_cf1_ = d_1067___mcc_h0_
            d_1073_v17_: _dafny.MultiSet
            d_1073_v17_ = _dafny.MultiSet([(self).f22, (self).f22, (self).f22])
            d_1073_v17_ = d_1073_v17_
            d_1065_v15_ = d_1070_cf3_
            d_1074_v18_: _dafny.Set
            d_1074_v18_ = _dafny.Set({p0})
            r1 = d_1074_v18_
            d_1075_v19_: str
            d_1075_v19_ = _dafny.CodePoint('y')
            d_1076_v20_: _dafny.Map
            d_1076_v20_ = _dafny.Map({(_dafny.MultiSet([d_1075_v19_])).cardinality: (0) - ((d_1065_v15_) + (d_1065_v15_))})
            d_1077_v22_: _dafny.Set
            d_1077_v22_ = _dafny.Set({_dafny.Map({_dafny.CodePoint('e'): d_1070_cf3_})})
            d_1078_v23_: _dafny.Map
            d_1078_v23_ = _dafny.Map({d_1075_v19_: d_1077_v22_})
            def iife72_():
                coll46_ = _dafny.Map()
                compr_46_: _dafny.Map
                for compr_46_ in (((d_1078_v23_)[d_1075_v19_] if (d_1075_v19_) in (d_1078_v23_) else d_1077_v22_)).Elements:
                    d_1079_v21_: _dafny.Map = compr_46_
                    if (d_1079_v21_) in (((d_1078_v23_)[d_1075_v19_] if (d_1075_v19_) in (d_1078_v23_) else d_1077_v22_)):
                        coll46_[d_1079_v21_] = p0
                return _dafny.Map(coll46_)
            def iife73_():
                coll47_ = _dafny.Map()
                compr_47_: _dafny.Map
                for compr_47_ in (((d_1078_v23_)[d_1075_v19_] if (d_1075_v19_) in (d_1078_v23_) else d_1077_v22_)).Elements:
                    d_1080_v21_: _dafny.Map = compr_47_
                    if (d_1080_v21_) in (((d_1078_v23_)[d_1075_v19_] if (d_1075_v19_) in (d_1078_v23_) else d_1077_v22_)):
                        coll47_[d_1080_v21_] = p0
                return _dafny.Map(coll47_)
            (globalState).f9 = ((d_1076_v20_)[(len(iife72_()
            )) * (d_1065_v15_)] if ((len(iife73_()
            )) * (d_1065_v15_)) in (d_1076_v20_) else d_1070_cf3_)
        elif source11_.is_DC2:
            d_1081___mcc_h3_ = source11_.cf4
            d_1082___mcc_h4_ = source11_.cf5
            d_1083___mcc_h5_ = source11_.cf6
            d_1084___mcc_h6_ = source11_.cf7
            d_1085_cf7_ = d_1084___mcc_h6_
            d_1086_cf6_ = d_1083___mcc_h5_
            d_1087_cf5_ = d_1082___mcc_h4_
            d_1088_cf4_ = d_1081___mcc_h3_
            d_1089_v24_: str
            d_1089_v24_ = _dafny.CodePoint('g')
            d_1090_v25_: bool
            d_1091_v26_: _dafny.Array
            out10_: bool
            out11_: _dafny.Array
            out10_, out11_ = (self).m13((d_1086_cf6_ if self.f21 else d_1065_v15_), (self).f22, d_1086_cf6_, default__.fm1(d_1085_cf7_, d_1089_v24_, globalState), globalState)
            d_1090_v25_ = out10_
            d_1091_v26_ = out11_
            if not((self).f22):
                d_1092_v27_: _dafny.Map
                d_1092_v27_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([d_1065_v15_ for d_1093_i3_ in range(default__.abs(175))])): d_1086_cf6_})
                d_1094_v28_: _dafny.Seq
                d_1094_v28_ = _dafny.SeqWithoutIsStrInference([len(d_1092_v27_)])
                d_1094_v28_ = d_1094_v28_
                d_1095_v29_: _dafny.MultiSet
                d_1095_v29_ = _dafny.MultiSet([d_1086_cf6_, d_1065_v15_])
                d_1096_v30_: _dafny.Map
                d_1096_v30_ = _dafny.Map({d_1088_cf4_: d_1095_v29_})
                d_1097_v31_: _dafny.Map
                d_1097_v31_ = _dafny.Map({d_1065_v15_: self.f21})
                d_1098_v32_: _dafny.Seq
                d_1098_v32_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bfiaytxu"))
                d_1099_v33_: _dafny.Seq
                d_1099_v33_ = _dafny.SeqWithoutIsStrInference([((d_1097_v31_)[len(d_1098_v32_)] if (len(d_1098_v32_)) in (d_1097_v31_) else d_1090_v25_)])
                d_1096_v30_ = ((d_1096_v30_).set(_dafny.MultiSet(d_1099_v33_), d_1095_v29_)) | (d_1096_v30_)
                d_1100_v34_: D2
                d_1100_v34_ = D2_DC5(d_1098_v32_)
                d_1101_v35_: _dafny.Array
                nw196_ = _dafny.Array(None, 23)
                nw196_[int(0)] = d_1100_v34_
                nw196_[int(1)] = d_1100_v34_
                nw196_[int(2)] = d_1100_v34_
                nw196_[int(3)] = default__.fm49(d_1086_cf6_, globalState)
                nw196_[int(4)] = d_1100_v34_
                nw196_[int(5)] = d_1100_v34_
                nw196_[int(6)] = d_1100_v34_
                nw196_[int(7)] = (D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))) if (self).f22 else d_1100_v34_)
                nw196_[int(8)] = D2_DC5(d_1098_v32_)
                nw196_[int(9)] = d_1100_v34_
                nw196_[int(10)] = d_1100_v34_
                nw196_[int(11)] = d_1100_v34_
                nw196_[int(12)] = d_1100_v34_
                nw196_[int(13)] = D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eywbyk")))
                nw196_[int(14)] = d_1100_v34_
                nw196_[int(15)] = d_1100_v34_
                nw196_[int(16)] = D2_DC5(d_1098_v32_)
                nw196_[int(17)] = default__.fm49(d_1065_v15_, globalState)
                nw196_[int(18)] = d_1100_v34_
                nw196_[int(19)] = d_1100_v34_
                nw196_[int(20)] = d_1100_v34_
                nw196_[int(21)] = d_1100_v34_
                nw196_[int(22)] = d_1100_v34_
                d_1101_v35_ = nw196_
                index172_ = default__.safeIndex(658, (d_1101_v35_).length(0))
                (d_1101_v35_)[index172_] = d_1100_v34_
                index173_ = default__.safeIndex(658, (d_1101_v35_).length(0))
                def iife74_(_pat_let13_0):
                    def iife75_(d_1102_dt__update__tmp_h0_):
                        def iife76_(_pat_let14_0):
                            def iife77_(d_1103_dt__update_hcf10_h0_):
                                return D2_DC5(d_1103_dt__update_hcf10_h0_)
                            return iife77_(_pat_let14_0)
                        return iife76_(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "shuvv")))
                    return iife75_(_pat_let13_0)
                (d_1101_v35_)[index173_] = iife74_(D2_DC5(d_1098_v32_))
                d_1104_v36_: _dafny.Array
                nw197_ = _dafny.Array(_dafny.Seq({}), 1)
                d_1104_v36_ = nw197_
                d_1105_v37_: _dafny.MultiSet
                d_1105_v37_ = _dafny.MultiSet([d_1104_v36_])
                d_1086_cf6_ = ((d_1105_v37_)[d_1104_v36_] if (d_1104_v36_) in (d_1105_v37_) else d_1065_v15_)
                d_1106_v38_: _dafny.Array
                nw198_ = _dafny.Array(int(0), 19)
                d_1106_v38_ = nw198_
                d_1106_v38_ = d_1106_v38_
            elif True:
                d_1107_v39_: _dafny.Array
                def lambda57_(d_1108_i4_):
                    return _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "dqxxcuign"))

                init32_ = lambda57_
                nw199_ = _dafny.Array(None, 9)
                for i0_32_ in range(nw199_.length(0)):
                    nw199_[i0_32_] = init32_(i0_32_)
                d_1107_v39_ = nw199_
                d_1109_v40_: _dafny.Array
                nw200_ = _dafny.Array(None, 12)
                nw200_[int(0)] = d_1107_v39_
                nw200_[int(1)] = d_1107_v39_
                nw200_[int(2)] = d_1107_v39_
                nw200_[int(3)] = d_1107_v39_
                nw200_[int(4)] = d_1107_v39_
                nw200_[int(5)] = d_1107_v39_
                nw200_[int(6)] = d_1107_v39_
                nw200_[int(7)] = d_1107_v39_
                nw200_[int(8)] = d_1107_v39_
                nw200_[int(9)] = d_1107_v39_
                nw200_[int(10)] = d_1107_v39_
                nw200_[int(11)] = d_1107_v39_
                d_1109_v40_ = nw200_
                index174_ = default__.safeIndex(347, (d_1109_v40_).length(0))
                (d_1109_v40_)[index174_] = d_1107_v39_
                index175_ = default__.safeIndex(347, (d_1109_v40_).length(0))
                (d_1109_v40_)[index175_] = d_1107_v39_
                d_1110_v41_: _dafny.Seq
                d_1110_v41_ = _dafny.SeqWithoutIsStrInference([d_1089_v24_])
                d_1111_v42_: _dafny.Seq
                d_1111_v42_ = _dafny.SeqWithoutIsStrInference([d_1065_v15_, (d_1065_v15_) - (default__.fm0(d_1085_cf7_, globalState)), d_1065_v15_, d_1065_v15_, default__.fm0(not(d_1090_v25_), globalState)])
                rhs201_ = d_1110_v41_
                rhs202_ = d_1111_v42_
                rhs203_ = self.f21
                lhs130_ = globalState
                d_1110_v41_ = rhs201_
                d_1111_v42_ = rhs202_
                lhs130_.f8 = rhs203_
                (self).f21 = not(default__.fm1((self).f22, d_1089_v24_, globalState))
                (globalState).f8 = not(not((self).f22))
                d_1112_v43_: _dafny.Set
                d_1112_v43_ = _dafny.Set({d_1089_v24_})
                (globalState).f8 = (len(d_1112_v43_)) < (d_1086_cf6_)
            d_1113_v44_: D2
            d_1113_v44_ = D2_DC8(d_1086_cf6_, d_1089_v24_)
            d_1065_v15_ = default__.safeModuloInt(((d_1113_v44_).cf15) * (-533), d_1065_v15_)
            d_1114_v45_: _dafny.Array
            nw201_ = _dafny.Array(False, 25)
            d_1114_v45_ = nw201_
            d_1115_v46_: _dafny.Seq
            d_1115_v46_ = _dafny.SeqWithoutIsStrInference([d_1114_v45_, d_1114_v45_, d_1114_v45_])
            d_1114_v45_ = (d_1115_v46_)[default__.safeIndex(d_1065_v15_, len(d_1115_v46_))]
        elif True:
            d_1116___mcc_h7_ = source11_.cf0
            d_1117_cf0_ = d_1116___mcc_h7_
            d_1118_v47_: _dafny.Set
            d_1118_v47_ = _dafny.Set({self.f21})
            d_1119_v48_: _dafny.Map
            d_1119_v48_ = _dafny.Map({p0: _dafny.Set({(self).f22})})
            d_1120_v49_: _dafny.MultiSet
            d_1120_v49_ = _dafny.MultiSet([d_1118_v47_, ((d_1119_v48_)[(self).f22] if ((self).f22) in (d_1119_v48_) else _dafny.Set({not((self).f22), p0}))])
            d_1120_v49_ = (d_1120_v49_) - (d_1120_v49_)
            if p0:
                (globalState).f9 = d_1117_cf0_
                d_1121_v50_: _dafny.Array
                def lambda58_(d_1122_p0_):
                    def lambda59_(d_1123_i5_):
                        return d_1122_p0_

                    return lambda59_

                init33_ = lambda58_(p0)
                nw202_ = _dafny.Array(None, 12)
                for i0_33_ in range(nw202_.length(0)):
                    nw202_[i0_33_] = init33_(i0_33_)
                d_1121_v50_ = nw202_
                index176_ = default__.safeIndex(187, (d_1121_v50_).length(0))
                (d_1121_v50_)[index176_] = (self).f22
                d_1124_v51_: D2
                d_1124_v51_ = D2_DC8((0) - (d_1065_v15_), _dafny.CodePoint('k'))
                d_1125_v52_: str
                d_1125_v52_ = _dafny.CodePoint('g')
                d_1126_v53_: _dafny.MultiSet
                d_1126_v53_ = _dafny.MultiSet([d_1117_cf0_, (0) - (len(_dafny.Set({default__.fm1(p0, d_1125_v52_, globalState)}))), (self).fm47(d_1120_v49_, globalState)])
                d_1127_v54_: _dafny.Seq
                d_1127_v54_ = _dafny.SeqWithoutIsStrInference([d_1065_v15_, d_1065_v15_, (d_1126_v53_).cardinality])
                index177_ = default__.safeIndex(187, (d_1121_v50_).length(0))
                (d_1121_v50_)[index177_] = ((d_1124_v51_).cf15) == ((d_1127_v54_)[default__.safeIndex(d_1117_cf0_, len(d_1127_v54_))])
                d_1065_v15_ = (0) - (d_1117_cf0_)
                d_1128_v55_: _dafny.MultiSet
                d_1128_v55_ = d_1126_v53_
                d_1129_v56_: C1
                nw203_ = C1()
                nw203_.ctor__(p0, not(self.f21), D0_DC0(d_1065_v15_), ((d_1128_v55_)).ispropersubset(d_1126_v53_))
                d_1129_v56_ = nw203_
                d_1065_v15_ = (d_1117_cf0_) + (default__.safeModuloInt(d_1065_v15_, d_1065_v15_))
            elif True:
                d_1130_v57_: _dafny.Array
                nw204_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 25)
                d_1130_v57_ = nw204_
                d_1131_v58_: _dafny.Map
                d_1131_v58_ = _dafny.Map({not ((self).f22) or (p0): d_1130_v57_})
                d_1131_v58_ = (d_1131_v58_).set(p0, d_1130_v57_)
                (self).f21 = (self).f22
                d_1132_v59_: _dafny.Array
                def lambda60_(d_1133_i6_):
                    return False

                init34_ = lambda60_
                nw205_ = _dafny.Array(None, 18)
                for i0_34_ in range(nw205_.length(0)):
                    nw205_[i0_34_] = init34_(i0_34_)
                d_1132_v59_ = nw205_
                d_1134_v60_: _dafny.Map
                d_1134_v60_ = _dafny.Map({self.f21: d_1132_v59_})
                d_1135_v61_: _dafny.Map
                d_1135_v61_ = _dafny.Map({((d_1134_v60_)[p0] if (p0) in (d_1134_v60_) else d_1132_v59_): 42})
                d_1136_v62_: _dafny.Seq
                d_1136_v62_ = _dafny.SeqWithoutIsStrInference([(0) - (d_1117_cf0_)])
                d_1135_v61_ = (d_1135_v61_).set((d_1132_v59_ if not((self).f22) else ((d_1134_v60_)[(self).f22] if ((self).f22) in (d_1134_v60_) else d_1132_v59_)), len(d_1136_v62_))
                d_1137_v63_: str
                d_1137_v63_ = _dafny.CodePoint('l')
                d_1138_v64_: _dafny.Seq
                d_1138_v64_ = _dafny.SeqWithoutIsStrInference([default__.fm1(self.f21, d_1137_v63_, globalState), self.f21])
                d_1139_v65_: _dafny.MultiSet
                d_1139_v65_ = _dafny.MultiSet([d_1065_v15_])
                d_1140_v66_: _dafny.Map
                d_1140_v66_ = _dafny.Map({216: d_1117_cf0_})
                d_1141_v67_: _dafny.Map
                d_1141_v67_ = _dafny.Map({d_1117_cf0_: default__.fm2(d_1117_cf0_, d_1139_v65_, d_1140_v66_, globalState)})
                d_1142_v68_: _dafny.Map
                d_1142_v68_ = _dafny.Map({d_1117_cf0_: d_1141_v67_})
                d_1143_v69_: _dafny.Seq
                d_1143_v69_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "higtskjfd"))
                d_1144_v70_: D9
                d_1144_v70_ = D9_DC22(len((d_1138_v64_) + (d_1138_v64_)), ((d_1142_v68_)[-276] if (-276) in (d_1142_v68_) else (d_1141_v67_).set(len(d_1138_v64_), d_1143_v69_)))
                d_1144_v70_ = d_1144_v70_
                d_1117_cf0_ = (0) - (d_1065_v15_)
            d_1145_v71_: _dafny.Seq
            d_1145_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bub"))
            if (len((d_1145_v71_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))))) == ((0) - ((d_1117_cf0_) + (default__.fm0((self).f22, globalState)))):
                d_1146_v72_: _dafny.Seq
                d_1146_v72_ = _dafny.SeqWithoutIsStrInference([d_1117_cf0_])
                d_1147_v73_: _dafny.Set
                d_1147_v73_ = _dafny.Set({d_1146_v72_})
                d_1148_v74_: C0
                nw206_ = C0()
                nw206_.ctor__(not((d_1147_v73_).issubset(d_1147_v73_)))
                d_1148_v74_ = nw206_
                d_1146_v72_ = (D5_DC13(_dafny.SeqWithoutIsStrInference([d_1117_cf0_]))).cf25
                d_1149_v75_: _dafny.Array
                def lambda61_(d_1150_cf0_):
                    def lambda62_(d_1151_i7_):
                        return (d_1151_i7_) + (d_1150_cf0_)

                    return lambda62_

                init35_ = lambda61_(d_1117_cf0_)
                nw207_ = _dafny.Array(None, 29)
                for i0_35_ in range(nw207_.length(0)):
                    nw207_[i0_35_] = init35_(i0_35_)
                d_1149_v75_ = nw207_
                index178_ = default__.safeIndex(761, (d_1149_v75_).length(0))
                (d_1149_v75_)[index178_] = (0) - (d_1065_v15_)
                d_1152_v76_: _dafny.Map
                d_1152_v76_ = _dafny.Map({d_1117_cf0_: d_1065_v15_})
                index179_ = default__.safeIndex(761, (d_1149_v75_).length(0))
                rhs204_ = default__.safeDivisionInt(len(d_1152_v76_), d_1065_v15_)
                rhs205_ = not((d_1148_v74_).f17)
                lhs131_ = d_1149_v75_
                lhs132_ = default__.safeIndex(761, (d_1149_v75_).length(0))
                lhs133_ = self
                lhs131_[lhs132_] = rhs204_
                lhs133_.f21 = rhs205_
                index180_ = default__.safeIndex(761, (d_1149_v75_).length(0))
                (d_1149_v75_)[index180_] = (d_1149_v75_)[default__.safeIndex(761, (d_1149_v75_).length(0))]
                d_1153_v77_: _dafny.Array
                def lambda63_(d_1154_v71_):
                    def lambda64_(d_1155_i8_):
                        return (d_1154_v71_) + (d_1154_v71_)

                    return lambda64_

                init36_ = lambda63_(d_1145_v71_)
                nw208_ = _dafny.Array(None, 4)
                for i0_36_ in range(nw208_.length(0)):
                    nw208_[i0_36_] = init36_(i0_36_)
                d_1153_v77_ = nw208_
                index181_ = default__.safeIndex(23, (d_1153_v77_).length(0))
                (d_1153_v77_)[index181_] = d_1145_v71_
                d_1156_v78_: _dafny.Seq
                d_1156_v78_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_1157_i9_ in range(default__.abs(224))])])
                index182_ = default__.safeIndex(23, (d_1153_v77_).length(0))
                (d_1153_v77_)[index182_] = ((d_1156_v78_)[default__.safeIndex(d_1117_cf0_, len(d_1156_v78_))] if (d_1117_cf0_) >= (d_1065_v15_) else d_1145_v71_)
            elif True:
                d_1158_v79_: _dafny.Map
                d_1158_v79_ = _dafny.Map({False: p0})
                d_1159_v80_: _dafny.Map
                d_1159_v80_ = _dafny.Map({d_1158_v79_: (self).f22})
                d_1160_v82_: _dafny.Array
                nw209_ = _dafny.Array(None, 8)
                nw209_[int(0)] = d_1159_v80_
                nw209_[int(1)] = d_1159_v80_
                nw209_[int(2)] = d_1159_v80_
                nw209_[int(3)] = d_1159_v80_
                nw209_[int(4)] = d_1159_v80_
                nw209_[int(5)] = d_1159_v80_
                nw209_[int(6)] = d_1159_v80_
                def iife78_():
                    coll48_ = _dafny.Map()
                    compr_48_: _dafny.Map
                    for compr_48_ in ((_dafny.Map({d_1158_v79_: p0})).set(d_1158_v79_, self.f21)).keys.Elements:
                        d_1161_v81_: _dafny.Map = compr_48_
                        if (d_1161_v81_) in ((_dafny.Map({d_1158_v79_: p0})).set(d_1158_v79_, self.f21)):
                            coll48_[d_1161_v81_] = not((self).f22)
                    return _dafny.Map(coll48_)
                nw209_[int(7)] = iife78_()
                
                d_1160_v82_ = nw209_
                d_1162_v84_: _dafny.Map
                d_1162_v84_ = _dafny.Map({d_1158_v79_: d_1117_cf0_})
                index183_ = default__.safeIndex(53, (d_1160_v82_).length(0))
                def iife79_():
                    coll49_ = _dafny.Map()
                    compr_49_: _dafny.Map
                    for compr_49_ in (d_1162_v84_).keys.Elements:
                        d_1163_v83_: _dafny.Map = compr_49_
                        if (d_1163_v83_) in (d_1162_v84_):
                            coll49_[d_1163_v83_] = p0
                    return _dafny.Map(coll49_)
                (d_1160_v82_)[index183_] = (d_1159_v80_) | (iife79_()
                )
                d_1164_v85_: _dafny.Array
                nw210_ = _dafny.Array(False, 19)
                d_1164_v85_ = nw210_
                d_1165_v86_: _dafny.Map
                d_1165_v86_ = _dafny.Map({d_1117_cf0_: self.f21})
                index184_ = default__.safeIndex(110, (d_1164_v85_).length(0))
                (d_1164_v85_)[index184_] = ((d_1165_v86_)[264] if (264) in (d_1165_v86_) else True)
                d_1166_v87_: str
                d_1166_v87_ = _dafny.CodePoint('o')
                index185_ = default__.safeIndex(158, (d_1164_v85_).length(0))
                (d_1164_v85_)[index185_] = default__.fm1(p0, d_1166_v87_, globalState)
                d_1167_v89_: _dafny.Array
                def lambda65_(d_1168_v79_):
                    def lambda66_(d_1169_i10_):
                        def iife80_():
                            coll50_ = _dafny.Map()
                            compr_50_: int
                            for compr_50_ in _dafny.IntegerRange(921, 224):
                                d_1170_v88_: int = compr_50_
                                if ((921) <= (d_1170_v88_)) and ((d_1170_v88_) < (224)):
                                    coll50_[(d_1170_v88_) * (42)] = d_1168_v79_
                            return _dafny.Map(coll50_)
                        return iife80_()
                        

                    return lambda66_

                init37_ = lambda65_(d_1158_v79_)
                nw211_ = _dafny.Array(None, 6)
                for i0_37_ in range(nw211_.length(0)):
                    nw211_[i0_37_] = init37_(i0_37_)
                d_1167_v89_ = nw211_
                d_1171_v90_: _dafny.Map
                d_1171_v90_ = _dafny.Map({(self).fm47(_dafny.MultiSet([d_1118_v47_]), globalState): d_1158_v79_})
                index186_ = default__.safeIndex(218, (d_1167_v89_).length(0))
                (d_1167_v89_)[index186_] = d_1171_v90_
                d_1172_v91_: _dafny.Array
                nw212_ = _dafny.Array(D2.default()(), 12)
                d_1172_v91_ = nw212_
                d_1173_v92_: _dafny.MultiSet
                d_1173_v92_ = _dafny.MultiSet([d_1172_v91_])
                d_1174_v93_: _dafny.Seq
                d_1174_v93_ = _dafny.SeqWithoutIsStrInference([(d_1173_v92_).cardinality])
                index187_ = default__.safeIndex(53, (d_1160_v82_).length(0))
                index188_ = default__.safeIndex(110, (d_1164_v85_).length(0))
                index189_ = default__.safeIndex(158, (d_1164_v85_).length(0))
                index190_ = default__.safeIndex(218, (d_1167_v89_).length(0))
                rhs206_ = _dafny.Map({_dafny.Map({p0: True}): (d_1065_v15_) <= (default__.fm0(self.f21, globalState))})
                rhs207_ = (d_1174_v93_) <= (d_1174_v93_)
                rhs208_ = (542) == (len(d_1145_v71_))
                rhs209_ = not(not(self.f21))
                rhs210_ = (_dafny.Map({d_1117_cf0_: _dafny.Map({p0: True})}) if False else d_1171_v90_)
                lhs134_ = d_1160_v82_
                lhs135_ = default__.safeIndex(53, (d_1160_v82_).length(0))
                lhs136_ = globalState
                lhs137_ = d_1164_v85_
                lhs138_ = default__.safeIndex(110, (d_1164_v85_).length(0))
                lhs139_ = d_1164_v85_
                lhs140_ = default__.safeIndex(158, (d_1164_v85_).length(0))
                lhs141_ = d_1167_v89_
                lhs142_ = default__.safeIndex(218, (d_1167_v89_).length(0))
                lhs134_[lhs135_] = rhs206_
                lhs136_.f8 = rhs207_
                lhs137_[lhs138_] = rhs208_
                lhs139_[lhs140_] = rhs209_
                lhs141_[lhs142_] = rhs210_
                d_1175_v94_: _dafny.Seq
                d_1175_v94_ = _dafny.SeqWithoutIsStrInference([d_1118_v47_, d_1118_v47_])
                d_1176_v95_: _dafny.MultiSet
                d_1176_v95_ = _dafny.MultiSet([(self).fm47(_dafny.MultiSet(d_1175_v94_), globalState)])
                d_1177_v97_: _dafny.Map
                def iife81_():
                    coll51_ = _dafny.Map()
                    compr_51_: int
                    for compr_51_ in _dafny.IntegerRange(559, 914):
                        d_1178_v96_: int = compr_51_
                        if ((559) <= (d_1178_v96_)) and ((d_1178_v96_) < (914)):
                            coll51_[(d_1178_v96_) * ((0) - (d_1065_v15_))] = d_1117_cf0_
                    return _dafny.Map(coll51_)
                d_1177_v97_ = _dafny.Map({d_1065_v15_: len(iife81_()
                )})
                d_1179_v98_: C0
                nw213_ = C0()
                nw213_.ctor__(self.f21)
                d_1179_v98_ = nw213_
                d_1180_v99_: _dafny.MultiSet
                d_1180_v99_ = _dafny.MultiSet([d_1179_v98_])
                d_1181_v100_: _dafny.Map
                d_1181_v100_ = _dafny.Map({p0: d_1180_v99_})
                d_1182_v101_: _dafny.Array
                nw214_ = _dafny.Array(None, 18)
                nw214_[int(0)] = _dafny.SeqWithoutIsStrInference([d_1166_v87_ for d_1183_i11_ in range(default__.abs(638))])
                nw214_[int(1)] = (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yygsnu"))) + (d_1145_v71_)
                nw214_[int(2)] = (d_1145_v71_) + (d_1145_v71_)
                nw214_[int(3)] = ((d_1145_v71_) + ((default__.fm2((0) - (d_1117_cf0_), d_1176_v95_, d_1177_v97_, globalState)).set(default__.safeIndex(d_1065_v15_, len(default__.fm2((0) - (d_1117_cf0_), d_1176_v95_, d_1177_v97_, globalState))), _dafny.CodePoint('g')))).set(default__.safeIndex(len(d_1181_v100_), len((d_1145_v71_) + ((default__.fm2((0) - (d_1117_cf0_), d_1176_v95_, d_1177_v97_, globalState)).set(default__.safeIndex(d_1065_v15_, len(default__.fm2((0) - (d_1117_cf0_), d_1176_v95_, d_1177_v97_, globalState))), _dafny.CodePoint('g'))))), d_1166_v87_)
                nw214_[int(4)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wquijjw"))
                nw214_[int(5)] = (d_1145_v71_) + (_dafny.SeqWithoutIsStrInference([d_1166_v87_ for d_1184_i12_ in range(default__.abs(-538))]))
                nw214_[int(6)] = _dafny.SeqWithoutIsStrInference([d_1166_v87_ for d_1185_i13_ in range(default__.abs(-476))])
                nw214_[int(7)] = (d_1145_v71_) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gmheovam")))
                nw214_[int(8)] = d_1145_v71_
                nw214_[int(9)] = default__.fm2(589, d_1176_v95_, d_1177_v97_, globalState)
                nw214_[int(10)] = (d_1145_v71_) + (d_1145_v71_)
                nw214_[int(11)] = d_1145_v71_
                nw214_[int(12)] = (_dafny.SeqWithoutIsStrInference([d_1166_v87_ for d_1186_i14_ in range(default__.abs(-550))])) + (d_1145_v71_)
                nw214_[int(13)] = (d_1145_v71_).set(default__.safeIndex(len(d_1145_v71_), len(d_1145_v71_)), (d_1145_v71_)[default__.safeIndex(d_1065_v15_, len(d_1145_v71_))])
                nw214_[int(14)] = d_1145_v71_
                nw214_[int(15)] = (d_1145_v71_) + ((d_1145_v71_).set(default__.safeIndex(934, len(d_1145_v71_)), (d_1145_v71_)[default__.safeIndex(d_1065_v15_, len(d_1145_v71_))]))
                nw214_[int(16)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hrfp"))
                nw214_[int(17)] = d_1145_v71_
                d_1182_v101_ = nw214_
                d_1187_v102_: _dafny.Seq
                d_1187_v102_ = _dafny.SeqWithoutIsStrInference([(d_1164_v85_)[default__.safeIndex(110, (d_1164_v85_).length(0))]])
                d_1188_v103_: D7
                d_1188_v103_ = D7_DC16(d_1187_v102_)
                d_1189_v104_: _dafny.Seq
                d_1189_v104_ = _dafny.SeqWithoutIsStrInference([d_1188_v103_])
                d_1190_v105_: _dafny.Map
                d_1190_v105_ = _dafny.Map({default__.safeModuloInt(d_1117_cf0_, len(d_1189_v104_)): d_1182_v101_})
                d_1182_v101_ = ((d_1190_v105_)[default__.safeModuloInt(d_1117_cf0_, d_1117_cf0_)] if (default__.safeModuloInt(d_1117_cf0_, d_1117_cf0_)) in (d_1190_v105_) else (d_1182_v101_ if p0 else d_1182_v101_))
                r0 = default__.fm0(not(False), globalState)
                d_1191_v106_: D2
                d_1191_v106_ = D2_DC7(p0, len(_dafny.SeqWithoutIsStrInference([d_1166_v87_ for d_1192_i15_ in range(default__.abs(320))])), not((self).f22), d_1145_v71_)
                d_1193_v107_: _dafny.Map
                d_1193_v107_ = _dafny.Map({(d_1177_v97_ if (d_1164_v85_)[default__.safeIndex(110, (d_1164_v85_).length(0))] else (d_1177_v97_).set(d_1065_v15_, len(d_1187_v102_))): (d_1191_v106_).cf12})
                def iife82_():
                    coll52_ = _dafny.Map()
                    compr_52_: _dafny.Map
                    for compr_52_ in (_dafny.Set({d_1177_v97_, _dafny.Map({d_1117_cf0_: d_1065_v15_})})).Elements:
                        d_1194_v108_: _dafny.Map = compr_52_
                        if (d_1194_v108_) in (_dafny.Set({d_1177_v97_, _dafny.Map({d_1117_cf0_: d_1065_v15_})})):
                            coll52_[d_1194_v108_] = (d_1176_v95_).cardinality
                    return _dafny.Map(coll52_)
                d_1193_v107_ = ((d_1193_v107_) | (iife82_()
                )) | (d_1193_v107_)
                d_1195_v109_: C1
                nw215_ = C1()
                nw215_.ctor__(self.f21, False, d_1066_v16_, self.f21)
                d_1195_v109_ = nw215_
            d_1196_v110_: _dafny.Map
            d_1196_v110_ = _dafny.Map({d_1117_cf0_: d_1065_v15_})
            r0 = (0) - (default__.safeModuloInt((d_1065_v15_ if False else len(d_1196_v110_)), (0) - (d_1117_cf0_)))
        d_1197_v111_: _dafny.Array
        def lambda67_(d_1198_v15_):
            def lambda68_(d_1199_i17_):
                return (d_1199_i17_) - (d_1198_v15_)

            return lambda68_

        init38_ = lambda67_(d_1065_v15_)
        nw216_ = _dafny.Array(None, 15)
        for i0_38_ in range(nw216_.length(0)):
            nw216_[i0_38_] = init38_(i0_38_)
        d_1197_v111_ = nw216_
        guard_loop_0_: int
        for guard_loop_0_ in _dafny.IntegerRange(0, (d_1197_v111_).length(0)):
            d_1200_i16_: int = guard_loop_0_
            if (True) and (((0) <= (d_1200_i16_)) and ((d_1200_i16_) < ((d_1197_v111_).length(0)))):
                (d_1197_v111_)[(d_1200_i16_)] = (d_1200_i16_) + ((0) - ((0) - ((591) - (-626))))
        d_1201_v112_: _dafny.Array
        nw217_ = _dafny.Array(_dafny.Map({}), 9)
        d_1201_v112_ = nw217_
        d_1202_v113_: D10
        d_1202_v113_ = D10_DC23(d_1201_v112_)
        d_1201_v112_ = (d_1202_v113_).cf42
        d_1203_v114_: _dafny.Map
        d_1203_v114_ = _dafny.Map({p0: False})
        d_1204_v115_: _dafny.Seq
        d_1204_v115_ = _dafny.SeqWithoutIsStrInference([d_1203_v114_])
        d_1205_v116_: str
        d_1205_v116_ = _dafny.CodePoint('y')
        hi9_ = len(default__.fm50(d_1205_v116_, True, (self).f22, globalState))
        for d_1206_i18_ in range(default__.safeModuloInt(len(d_1204_v115_), -112), hi9_):
            d_1207_v117_: _dafny.Set
            d_1207_v117_ = _dafny.Set({-535, -722})
            d_1208_v118_: D8
            d_1208_v118_ = D8_DC18((d_1207_v117_) - (d_1207_v117_))
            source12_ = d_1208_v118_
            if source12_.is_DC19:
                d_1209___mcc_h8_ = source12_.cf34
                d_1210___mcc_h9_ = source12_.cf35
                d_1211_cf35_ = d_1210___mcc_h9_
                d_1212_cf34_ = d_1209___mcc_h8_
                d_1211_cf35_ = len((_dafny.Map({d_1211_cf35_: _dafny.CodePoint('s')})).set(d_1206_i18_, d_1205_v116_))
                d_1213_v119_: _dafny.Map
                d_1213_v119_ = _dafny.Map({False: (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))).set(default__.safeIndex(d_1206_i18_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))), d_1205_v116_)})
                d_1213_v119_ = (d_1213_v119_).set(self.f21, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lxhumbbwu")))
                d_1214_v120_: _dafny.Map
                d_1214_v120_ = _dafny.Map({d_1066_v16_: _dafny.CodePoint('x')})
                d_1214_v120_ = ((default__.fm51(globalState)) | (d_1214_v120_)).set(d_1066_v16_, default__.fm3(d_1211_cf35_, (self).f22, globalState))
                (globalState).f9 = ((0) - ((0) - (d_1206_i18_))) + (d_1065_v15_)
            elif source12_.is_DC20:
                d_1215___mcc_h10_ = source12_.cf36
                d_1216___mcc_h11_ = source12_.cf37
                d_1217___mcc_h12_ = source12_.cf38
                d_1218_cf38_ = d_1217___mcc_h12_
                d_1219_cf37_ = d_1216___mcc_h11_
                d_1220_cf36_ = d_1215___mcc_h10_
                (globalState).f9 = d_1065_v15_
                r0 = d_1206_i18_
                d_1221_v121_: _dafny.MultiSet
                d_1221_v121_ = _dafny.MultiSet([d_1219_cf37_])
                (globalState).f8 = not(((d_1203_v114_)[(d_1221_v121_).issubset(d_1221_v121_)] if ((d_1221_v121_).issubset(d_1221_v121_)) in (d_1203_v114_) else (d_1065_v15_) >= (d_1206_i18_)))
                d_1222_v122_: D8
                d_1222_v122_ = D8_DC20(d_1220_cf36_, d_1219_cf37_, p0)
                rhs211_ = d_1222_v122_
                rhs212_ = d_1205_v116_
                d_1222_v122_ = rhs211_
                d_1205_v116_ = rhs212_
            elif True:
                d_1223___mcc_h13_ = source12_.cf33
                d_1224_cf33_ = d_1223___mcc_h13_
                default__.m0(globalState)
                d_1225_v123_: _dafny.Seq
                d_1225_v123_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ses"))
                (globalState).f9 = ((len(d_1225_v123_)) * (d_1065_v15_)) + (836)
                d_1226_v124_: _dafny.Set
                d_1226_v124_ = _dafny.Set({d_1225_v123_, d_1225_v123_, d_1225_v123_})
                (self).f21 = not(default__.fm1((d_1226_v124_).isdisjoint(_dafny.Set({d_1225_v123_})), d_1205_v116_, globalState))
                d_1227_v125_: _dafny.Map
                d_1227_v125_ = _dafny.Map({d_1065_v15_: (0) - ((0) - ((d_1206_i18_) - (d_1065_v15_)))})
                r0 = len(d_1227_v125_)
            (self).f21 = (self).f22
            (globalState).f9 = (default__.safeDivisionInt((0) - (d_1206_i18_), d_1065_v15_)) * (len(_dafny.SeqWithoutIsStrInference([d_1205_v116_ for d_1228_i19_ in range(default__.abs(779))])))
            (globalState).f8 = self.f21
        d_1229_v126_: _dafny.Seq
        d_1229_v126_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "taiis"))
        d_1230_v127_: D2
        d_1230_v127_ = D2_DC7(self.f21, d_1065_v15_, p0, d_1229_v126_)
        d_1231_v128_: _dafny.Set
        d_1231_v128_ = _dafny.Set({(d_1230_v127_).cf11})
        d_1232_v129_: _dafny.Map
        d_1232_v129_ = _dafny.Map({d_1231_v128_: d_1065_v15_})
        if (d_1231_v128_) in (d_1232_v129_):
            r0 = d_1065_v15_
            d_1233_v130_: _dafny.Set
            d_1233_v130_ = _dafny.Set({-109})
            d_1234_v131_: D8
            d_1234_v131_ = D8_DC18(d_1233_v130_)
            d_1235_v132_: _dafny.Map
            d_1235_v132_ = _dafny.Map({d_1234_v131_: (self).f22})
            d_1236_v133_: _dafny.Seq
            d_1236_v133_ = _dafny.SeqWithoutIsStrInference([d_1235_v132_, _dafny.Map({d_1234_v131_: (self).f22}), d_1235_v132_, d_1235_v132_, d_1235_v132_])
            d_1237_v134_: _dafny.Map
            d_1237_v134_ = _dafny.Map({not((self).f22): d_1065_v15_})
            d_1238_v135_: _dafny.MultiSet
            d_1238_v135_ = _dafny.MultiSet([d_1231_v128_, d_1231_v128_, d_1231_v128_])
            nw218_ = _dafny.Array(None, 17)
            nw218_[int(0)] = d_1065_v15_
            nw218_[int(1)] = len(d_1236_v133_)
            nw218_[int(2)] = default__.safeModuloInt(722, d_1065_v15_)
            nw218_[int(3)] = d_1065_v15_
            nw218_[int(4)] = len(d_1229_v126_)
            nw218_[int(5)] = d_1065_v15_
            nw218_[int(6)] = 241
            nw218_[int(7)] = default__.safeDivisionInt(len(d_1237_v134_), d_1065_v15_)
            nw218_[int(8)] = (d_1065_v15_) - (d_1065_v15_)
            nw218_[int(9)] = d_1065_v15_
            nw218_[int(10)] = (self).fm47(d_1238_v135_, globalState)
            nw218_[int(11)] = d_1065_v15_
            nw218_[int(12)] = d_1065_v15_
            nw218_[int(13)] = (d_1065_v15_ if p0 else d_1065_v15_)
            nw218_[int(14)] = d_1065_v15_
            nw218_[int(15)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "b")))
            nw218_[int(16)] = d_1065_v15_
            d_1197_v111_ = nw218_
            d_1239_v136_: _dafny.Map
            d_1239_v136_ = _dafny.Map({-248: d_1229_v126_})
            d_1240_v137_: D9
            d_1240_v137_ = D9_DC22(d_1065_v15_, d_1239_v136_)
            d_1241_v138_: _dafny.Map
            d_1241_v138_ = _dafny.Map({d_1065_v15_: d_1240_v137_})
            d_1241_v138_ = (d_1241_v138_).set(d_1065_v15_, d_1240_v137_)
            d_1242_v139_: _dafny.MultiSet
            d_1242_v139_ = _dafny.MultiSet([not(p0)])
            d_1243_v140_: _dafny.Map
            d_1243_v140_ = _dafny.Map({d_1065_v15_: (d_1242_v139_).isdisjoint(d_1242_v139_)})
            d_1243_v140_ = (d_1243_v140_).set((d_1065_v15_ if p0 else len(default__.fm52(globalState))), self.f21)
            d_1244_v141_: D4
            d_1244_v141_ = D4_DC12((self).f22, (self).f22, p0)
            d_1245_v142_: _dafny.Array
            nw219_ = _dafny.Array(None, 11)
            nw219_[int(0)] = (self).f22
            nw219_[int(1)] = p0
            nw219_[int(2)] = False
            nw219_[int(3)] = (self).f22
            nw219_[int(4)] = p0
            nw219_[int(5)] = self.f21
            nw219_[int(6)] = (d_1244_v141_).cf22
            nw219_[int(7)] = p0
            nw219_[int(8)] = (self).f22
            nw219_[int(9)] = (self).f22
            nw219_[int(10)] = p0
            d_1245_v142_ = nw219_
            d_1246_v143_: D7
            d_1246_v143_ = D7_DC17(d_1229_v126_, d_1245_v142_, p0, len(d_1231_v128_))
            d_1247_v144_: _dafny.Array
            nw220_ = _dafny.Array(None, 11)
            nw220_[int(0)] = (d_1246_v143_).cf30
            nw220_[int(1)] = d_1245_v142_
            nw220_[int(2)] = d_1245_v142_
            nw220_[int(3)] = d_1245_v142_
            nw220_[int(4)] = d_1245_v142_
            nw220_[int(5)] = d_1245_v142_
            nw220_[int(6)] = d_1245_v142_
            nw220_[int(7)] = d_1245_v142_
            nw220_[int(8)] = d_1245_v142_
            nw220_[int(9)] = d_1245_v142_
            nw220_[int(10)] = d_1245_v142_
            d_1247_v144_ = nw220_
            d_1248_v145_: _dafny.Map
            d_1248_v145_ = _dafny.Map({d_1247_v144_: default__.safeDivisionInt(d_1065_v15_, 512)})
            d_1248_v145_ = (d_1248_v145_).set(d_1247_v144_, default__.fm0((self).f22, globalState))
        elif True:
            (self).f21 = True
            d_1249_v146_: _dafny.MultiSet
            d_1249_v146_ = _dafny.MultiSet([self.f21])
            d_1250_v147_: _dafny.Seq
            d_1250_v147_ = _dafny.SeqWithoutIsStrInference([self.f21, not(p0), True])
            d_1251_v148_: _dafny.Map
            d_1251_v148_ = _dafny.Map({(self).f22: d_1065_v15_})
            if not(((d_1249_v146_).cardinality) == ((len(d_1250_v147_)) - (len(d_1251_v148_)))):
                (globalState).f9 = 367
                d_1252_v149_: _dafny.Map
                d_1252_v149_ = _dafny.Map({((d_1249_v146_).intersection(d_1249_v146_)).cardinality: (d_1065_v15_) * (default__.fm0((self).f22, globalState))})
                d_1253_v150_: _dafny.Seq
                d_1253_v150_ = _dafny.SeqWithoutIsStrInference([d_1197_v111_])
                d_1254_v151_: D9
                d_1254_v151_ = D9_DC21((d_1253_v150_)[default__.safeIndex(d_1065_v15_, len(d_1253_v150_))])
                d_1255_v152_: _dafny.Seq
                d_1255_v152_ = _dafny.SeqWithoutIsStrInference([d_1197_v111_, (d_1254_v151_).cf39, d_1197_v111_, d_1197_v111_, d_1197_v111_])
                d_1252_v149_ = (d_1252_v149_).set(len(d_1253_v150_), len((d_1203_v114_).set(False, self.f21)))
                d_1256_v153_: C0
                nw221_ = C0()
                nw221_.ctor__(self.f21)
                d_1256_v153_ = nw221_
                d_1257_v154_: _dafny.Array
                nw222_ = _dafny.Array(False, 12)
                d_1257_v154_ = nw222_
                index191_ = default__.safeIndex(63, (d_1257_v154_).length(0))
                (d_1257_v154_)[index191_] = p0
                index192_ = default__.safeIndex(63, (d_1257_v154_).length(0))
                rhs213_ = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nbmrqg")))
                rhs214_ = (d_1256_v153_).f17
                rhs215_ = (d_1257_v154_ if (self).f22 else d_1257_v154_)
                rhs216_ = (self).f22
                lhs143_ = globalState
                lhs144_ = d_1257_v154_
                lhs145_ = default__.safeIndex(63, (d_1257_v154_).length(0))
                lhs146_ = globalState
                lhs143_.f9 = rhs213_
                lhs144_[lhs145_] = rhs214_
                d_1257_v154_ = rhs215_
                lhs146_.f8 = rhs216_
                d_1258_v155_: D0
                d_1258_v155_ = D0_DC2(_dafny.MultiSet(d_1250_v147_), False, -678, (self).f22)
                d_1259_v156_: D4
                d_1259_v156_ = D4_DC12((d_1258_v155_).cf5, (self).f22, (self).f22)
                (globalState).f8 = (d_1259_v156_).cf24
            elif True:
                index193_ = default__.safeIndex(50, (d_1197_v111_).length(0))
                (d_1197_v111_)[index193_] = len((d_1203_v114_).set(not((self).f22), (self).f22))
                index194_ = default__.safeIndex(50, (d_1197_v111_).length(0))
                (d_1197_v111_)[index194_] = 711
                (globalState).f8 = (self).f22
                d_1260_v157_: C0
                nw223_ = C0()
                nw223_.ctor__(not(p0))
                d_1260_v157_ = nw223_
                d_1065_v15_ = d_1065_v15_
                (self).f21 = (d_1260_v157_).f17
            d_1261_v158_: C2
            nw224_ = C2()
            nw224_.ctor__((self).f11)
            d_1261_v158_ = nw224_
            r0 = d_1065_v15_
            d_1262_v159_: _dafny.Map
            d_1262_v159_ = _dafny.Map({d_1065_v15_: (d_1229_v126_) <= (d_1229_v126_)})
            d_1262_v159_ = d_1262_v159_
        r0 = (0) - (d_1065_v15_)
        r1 = default__.fm53(globalState)
        return r0, r1

    def m12(self, p0, p1, p2, globalState):
        r0: bool = False
        d_1263_v0_: _dafny.Seq
        d_1263_v0_ = _dafny.SeqWithoutIsStrInference([not(False)])
        d_1264_v1_: _dafny.Array
        nw225_ = _dafny.Array(None, 5)
        nw225_[int(0)] = (self).f22
        nw225_[int(1)] = (True) and (p1)
        nw225_[int(2)] = (self).f22
        nw225_[int(3)] = (_dafny.SeqWithoutIsStrInference([self.f21])) == (d_1263_v0_)
        nw225_[int(4)] = p0
        d_1264_v1_ = nw225_
        d_1265_v2_: _dafny.Seq
        d_1265_v2_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bhgv"))
        d_1266_v3_: int
        d_1266_v3_ = 485
        d_1267_v4_: str
        d_1267_v4_ = _dafny.CodePoint('l')
        index195_ = default__.safeIndex(729, (d_1264_v1_).length(0))
        (d_1264_v1_)[index195_] = ((d_1265_v2_).set(default__.safeIndex(d_1266_v3_, len(d_1265_v2_)), d_1267_v4_)) <= (d_1265_v2_)
        index196_ = default__.safeIndex(729, (d_1264_v1_).length(0))
        (d_1264_v1_)[index196_] = default__.fm1(p1, d_1267_v4_, globalState)
        d_1268_v5_: _dafny.Set
        d_1268_v5_ = _dafny.Set({p1, self.f21})
        d_1269_v6_: _dafny.Seq
        d_1269_v6_ = _dafny.SeqWithoutIsStrInference([len(d_1268_v5_)])
        hi10_ = d_1266_v3_
        for d_1270_i0_ in range((d_1269_v6_)[default__.safeIndex(d_1266_v3_, len(d_1269_v6_))], hi10_):
            d_1271_v7_: _dafny.Set
            d_1271_v7_ = _dafny.Set({d_1264_v1_})
            d_1271_v7_ = (d_1271_v7_).intersection(d_1271_v7_)
            if p1:
                d_1272_v8_: _dafny.Set
                d_1272_v8_ = _dafny.Set({d_1270_i0_})
                d_1273_v9_: _dafny.Map
                d_1273_v9_ = _dafny.Map({d_1272_v8_: self.f21})
                d_1266_v3_ = len((d_1273_v9_) | (d_1273_v9_))
                index197_ = default__.safeIndex(729, (d_1264_v1_).length(0))
                (d_1264_v1_)[index197_] = p0
                d_1274_v10_: _dafny.Array
                def lambda69_(d_1275_v3_):
                    def lambda70_(d_1276_i1_):
                        return (d_1276_i1_) - (d_1275_v3_)

                    return lambda70_

                init39_ = lambda69_(d_1266_v3_)
                nw226_ = _dafny.Array(None, 15)
                for i0_39_ in range(nw226_.length(0)):
                    nw226_[i0_39_] = init39_(i0_39_)
                d_1274_v10_ = nw226_
                d_1277_v11_: _dafny.MultiSet
                d_1277_v11_ = _dafny.MultiSet([d_1274_v10_])
                d_1278_v12_: _dafny.Array
                nw227_ = _dafny.Array(None, 18)
                nw227_[int(0)] = d_1270_i0_
                nw227_[int(1)] = d_1266_v3_
                nw227_[int(2)] = d_1270_i0_
                nw227_[int(3)] = -484
                nw227_[int(4)] = d_1266_v3_
                nw227_[int(5)] = d_1266_v3_
                nw227_[int(6)] = len(d_1269_v6_)
                nw227_[int(7)] = d_1270_i0_
                nw227_[int(8)] = len(_dafny.SeqWithoutIsStrInference([d_1266_v3_]))
                nw227_[int(9)] = d_1266_v3_
                nw227_[int(10)] = d_1270_i0_
                nw227_[int(11)] = d_1270_i0_
                nw227_[int(12)] = d_1270_i0_
                nw227_[int(13)] = d_1270_i0_
                nw227_[int(14)] = d_1266_v3_
                nw227_[int(15)] = (0) - (d_1270_i0_)
                nw227_[int(16)] = (D0_DC1(d_1277_v11_, d_1270_i0_, len(_dafny.SeqWithoutIsStrInference([d_1267_v4_ for d_1279_i2_ in range(default__.abs(939))])))).cf3
                nw227_[int(17)] = d_1270_i0_
                d_1278_v12_ = nw227_
                d_1280_v13_: _dafny.Map
                d_1280_v13_ = _dafny.Map({d_1266_v3_: d_1274_v10_})
                d_1280_v13_ = (d_1280_v13_).set(963, d_1278_v12_)
                d_1281_v14_: _dafny.Map
                d_1281_v14_ = _dafny.Map({d_1267_v4_: p1})
                d_1282_v15_: _dafny.Map
                d_1282_v15_ = _dafny.Map({865: (_dafny.Map({d_1267_v4_: p0})) | (d_1281_v14_)})
                d_1282_v15_ = (d_1282_v15_).set(d_1266_v3_, d_1281_v14_)
                (globalState).f9 = d_1266_v3_
            elif True:
                d_1283_v16_: C2
                nw228_ = C2()
                nw228_.ctor__((self).f11)
                d_1283_v16_ = nw228_
                index198_ = default__.safeIndex(729, (d_1264_v1_).length(0))
                (d_1264_v1_)[index198_] = p0
                d_1284_v17_: _dafny.Array
                nw229_ = _dafny.Array(int(0), 29)
                d_1284_v17_ = nw229_
                d_1285_v18_: _dafny.Array
                nw230_ = _dafny.Array(_dafny.CodePoint('D'), 1)
                d_1285_v18_ = nw230_
                rhs217_ = d_1284_v17_
                rhs218_ = default__.safeDivisionInt(d_1270_i0_, d_1270_i0_)
                rhs219_ = (d_1285_v18_ if default__.fm1(not(True), d_1267_v4_, globalState) else d_1285_v18_)
                lhs147_ = globalState
                d_1284_v17_ = rhs217_
                lhs147_.f9 = rhs218_
                d_1285_v18_ = rhs219_
                d_1286_v19_: _dafny.Array
                nw231_ = _dafny.Array(None, 9)
                nw231_[int(0)] = p2
                nw231_[int(1)] = p2
                nw231_[int(2)] = p2
                nw231_[int(3)] = p2
                nw231_[int(4)] = p2
                nw231_[int(5)] = (p2) | (p2)
                nw231_[int(6)] = _dafny.Map({self.f21: len(d_1269_v6_)})
                nw231_[int(7)] = _dafny.Map({(d_1264_v1_)[default__.safeIndex(729, (d_1264_v1_).length(0))]: d_1270_i0_})
                nw231_[int(8)] = _dafny.Map({p0: d_1266_v3_})
                d_1286_v19_ = nw231_
                index199_ = default__.safeIndex(729, (d_1264_v1_).length(0))
                rhs220_ = d_1267_v4_
                rhs221_ = (self).f22
                rhs222_ = d_1286_v19_
                rhs223_ = ((0) - (d_1266_v3_)) < (99)
                lhs148_ = globalState
                lhs149_ = d_1264_v1_
                lhs150_ = default__.safeIndex(729, (d_1264_v1_).length(0))
                d_1267_v4_ = rhs220_
                lhs148_.f8 = rhs221_
                d_1286_v19_ = rhs222_
                lhs149_[lhs150_] = rhs223_
                d_1287_v20_: D0
                d_1287_v20_ = D0_DC0(d_1270_i0_)
                d_1288_v21_: C1
                nw232_ = C1()
                nw232_.ctor__((d_1264_v1_)[default__.safeIndex(729, (d_1264_v1_).length(0))], self.f21, d_1287_v20_, not (self.f21) or (self.f21))
                d_1288_v21_ = nw232_
            default__.m0(globalState)
            (globalState).f9 = d_1266_v3_
        d_1289_v22_: _dafny.Map
        d_1289_v22_ = _dafny.Map({d_1266_v3_: self.f21})
        d_1290_v23_: _dafny.Array
        nw233_ = _dafny.Array(int(0), 10)
        d_1290_v23_ = nw233_
        d_1291_v24_: _dafny.Set
        d_1291_v24_ = _dafny.Set({d_1290_v23_, d_1290_v23_})
        (globalState).f9 = len(((d_1289_v22_ if False else _dafny.Map({960: p1}))).set(len((d_1291_v24_).intersection(d_1291_v24_)), p1))
        guard_loop_1_: int
        for guard_loop_1_ in _dafny.IntegerRange(0, (d_1264_v1_).length(0)):
            d_1292_i3_: int = guard_loop_1_
            if (True) and (((0) <= (d_1292_i3_)) and ((d_1292_i3_) < ((d_1264_v1_).length(0)))):
                (d_1264_v1_)[(d_1292_i3_)] = not(not((d_1266_v3_) < (d_1266_v3_)))
        d_1293_v25_: C0
        nw234_ = C0()
        nw234_.ctor__(True)
        d_1293_v25_ = nw234_
        d_1294_v26_: _dafny.Array
        def lambda71_(d_1295_v3_):
            def lambda72_(d_1296_i4_):
                return (_dafny.Set({d_1295_v3_}) if True else _dafny.Set({d_1295_v3_, 67, d_1295_v3_, d_1295_v3_}))

            return lambda72_

        init40_ = lambda71_(d_1266_v3_)
        nw235_ = _dafny.Array(None, 3)
        for i0_40_ in range(nw235_.length(0)):
            nw235_[i0_40_] = init40_(i0_40_)
        d_1294_v26_ = nw235_
        d_1297_v27_: _dafny.Set
        d_1297_v27_ = _dafny.Set({-662, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "d"))), d_1266_v3_, d_1266_v3_, (0) - (d_1266_v3_)})
        index200_ = default__.safeIndex(346, (d_1294_v26_).length(0))
        (d_1294_v26_)[index200_] = d_1297_v27_
        d_1298_v28_: _dafny.Array
        nw236_ = _dafny.Array(D2.default()(), 13)
        d_1298_v28_ = nw236_
        d_1299_v29_: D11
        d_1299_v29_ = D11_DC25(d_1298_v28_)
        d_1300_v30_: _dafny.Array
        nw237_ = _dafny.Array(None, 22)
        nw237_[int(0)] = d_1298_v28_
        nw237_[int(1)] = d_1298_v28_
        nw237_[int(2)] = d_1298_v28_
        nw237_[int(3)] = d_1298_v28_
        nw237_[int(4)] = d_1298_v28_
        nw237_[int(5)] = d_1298_v28_
        nw237_[int(6)] = d_1298_v28_
        nw237_[int(7)] = d_1298_v28_
        nw237_[int(8)] = d_1298_v28_
        nw237_[int(9)] = d_1298_v28_
        nw237_[int(10)] = d_1298_v28_
        nw237_[int(11)] = d_1298_v28_
        nw237_[int(12)] = d_1298_v28_
        nw237_[int(13)] = (d_1299_v29_).cf46
        nw237_[int(14)] = d_1298_v28_
        nw237_[int(15)] = d_1298_v28_
        nw237_[int(16)] = d_1298_v28_
        nw237_[int(17)] = d_1298_v28_
        nw237_[int(18)] = d_1298_v28_
        nw237_[int(19)] = d_1298_v28_
        nw237_[int(20)] = d_1298_v28_
        nw237_[int(21)] = d_1298_v28_
        d_1300_v30_ = nw237_
        d_1301_v31_: _dafny.Map
        d_1301_v31_ = _dafny.Map({default__.fm1(p0, d_1267_v4_, globalState): p1})
        d_1302_v32_: _dafny.Map
        d_1302_v32_ = _dafny.Map({len(d_1301_v31_): d_1298_v28_})
        d_1303_v33_: _dafny.Seq
        d_1303_v33_ = _dafny.SeqWithoutIsStrInference([((d_1302_v32_)[d_1266_v3_] if (d_1266_v3_) in (d_1302_v32_) else d_1298_v28_), d_1298_v28_, d_1298_v28_])
        d_1304_v34_: _dafny.Map
        d_1304_v34_ = _dafny.Map({d_1266_v3_: 862})
        index201_ = default__.safeIndex(330, (d_1300_v30_).length(0))
        (d_1300_v30_)[index201_] = (d_1303_v33_)[default__.safeIndex(((d_1304_v34_)[d_1266_v3_] if (d_1266_v3_) in (d_1304_v34_) else d_1266_v3_), len(d_1303_v33_))]
        index202_ = default__.safeIndex(346, (d_1294_v26_).length(0))
        index203_ = default__.safeIndex(330, (d_1300_v30_).length(0))
        rhs224_ = _dafny.Set({d_1266_v3_, d_1266_v3_, d_1266_v3_})
        rhs225_ = d_1298_v28_
        rhs226_ = d_1266_v3_
        rhs227_ = (d_1266_v3_) + (d_1266_v3_)
        lhs151_ = d_1294_v26_
        lhs152_ = default__.safeIndex(346, (d_1294_v26_).length(0))
        lhs153_ = d_1300_v30_
        lhs154_ = default__.safeIndex(330, (d_1300_v30_).length(0))
        lhs155_ = globalState
        lhs151_[lhs152_] = rhs224_
        lhs153_[lhs154_] = rhs225_
        lhs155_.f9 = rhs226_
        d_1266_v3_ = rhs227_
        r0 = (d_1293_v25_).f17
        return r0

    def m13(self, p0, p1, p2, p3, globalState):
        r0: bool = False
        r1: _dafny.Array = _dafny.Array(None, 0)
        d_1305_v0_: _dafny.Array
        nw238_ = _dafny.Array(_dafny.Array(None, 0), 21)
        d_1305_v0_ = nw238_
        d_1306_v1_: _dafny.Seq
        d_1306_v1_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gadrckqd"))
        d_1307_v3_: _dafny.Array
        nw239_ = _dafny.Array(None, 16)
        nw239_[int(0)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1308_i0_ in range(default__.abs(348))])
        nw239_[int(1)] = d_1306_v1_
        def iife83_():
            coll53_ = _dafny.Map()
            compr_53_: int
            for compr_53_ in _dafny.IntegerRange(-908, 962):
                d_1309_v2_: int = compr_53_
                if ((-908) <= (d_1309_v2_)) and ((d_1309_v2_) < (962)):
                    coll53_[(d_1309_v2_) - (116)] = _dafny.CodePoint('e')
            return _dafny.Map(coll53_)
        nw239_[int(2)] = _dafny.SeqWithoutIsStrInference([(D2_DC8(len(iife83_()
), _dafny.CodePoint('d'))).cf16 for d_1310_i1_ in range(default__.abs(-276))])
        nw239_[int(3)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "longt"))
        nw239_[int(4)] = d_1306_v1_
        nw239_[int(5)] = _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('v') for d_1311_i2_ in range(default__.abs(-677))])
        nw239_[int(6)] = d_1306_v1_
        nw239_[int(7)] = d_1306_v1_
        nw239_[int(8)] = d_1306_v1_
        nw239_[int(9)] = d_1306_v1_
        nw239_[int(10)] = d_1306_v1_
        nw239_[int(11)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "sp"))
        nw239_[int(12)] = d_1306_v1_
        nw239_[int(13)] = d_1306_v1_
        nw239_[int(14)] = d_1306_v1_
        nw239_[int(15)] = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nk"))
        d_1307_v3_ = nw239_
        index204_ = default__.safeIndex(689, (d_1305_v0_).length(0))
        (d_1305_v0_)[index204_] = d_1307_v3_
        index205_ = default__.safeIndex(689, (d_1305_v0_).length(0))
        (d_1305_v0_)[index205_] = d_1307_v3_
        d_1312_v4_: _dafny.Seq
        d_1312_v4_ = _dafny.SeqWithoutIsStrInference([p3])
        rhs228_ = default__.fm1((self).f22, _dafny.CodePoint('h'), globalState)
        rhs229_ = len(d_1306_v1_)
        rhs230_ = False
        rhs231_ = d_1312_v4_
        lhs156_ = globalState
        lhs157_ = globalState
        lhs158_ = globalState
        lhs156_.f8 = rhs228_
        lhs157_.f9 = rhs229_
        lhs158_.f8 = rhs230_
        d_1312_v4_ = rhs231_
        d_1313_v5_: _dafny.Array
        def lambda73_(d_1314_p0_, d_1315_v1_):
            def lambda74_(d_1316_i3_):
                return _dafny.Map({(self).f22: (_dafny.MultiSet([d_1314_p0_, len(d_1315_v1_), d_1314_p0_, d_1314_p0_, d_1314_p0_])).cardinality})

            return lambda74_

        init41_ = lambda73_(p0, d_1306_v1_)
        nw240_ = _dafny.Array(None, 7)
        for i0_41_ in range(nw240_.length(0)):
            nw240_[i0_41_] = init41_(i0_41_)
        d_1313_v5_ = nw240_
        d_1317_v6_: _dafny.Map
        d_1317_v6_ = _dafny.Map({p3: len(d_1306_v1_)})
        nw241_ = _dafny.Array(None, 1)
        nw241_[int(0)] = (d_1317_v6_) | (d_1317_v6_)
        d_1313_v5_ = nw241_
        (self).f21 = p1
        hi11_ = p0
        for d_1318_i4_ in range(len(d_1306_v1_), hi11_):
            (globalState).f8 = (self).f22
            d_1319_v7_: _dafny.MultiSet
            d_1319_v7_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xkcswj"))])
            d_1320_v8_: D5
            d_1320_v8_ = D5_DC13(default__.fm54(p1, not((self).f22), self.f21, globalState))
            source13_ = (d_1320_v8_ if (d_1319_v7_).ispropersubset(_dafny.MultiSet([d_1306_v1_])) else d_1320_v8_)
            if source13_.is_DC14:
                d_1321___mcc_h0_ = source13_.cf26
                d_1322_cf26_ = d_1321___mcc_h0_
                (globalState).f8 = (self).f22
                d_1323_v9_: _dafny.Array
                def lambda75_(d_1324_p3_):
                    def lambda76_(d_1325_i5_):
                        return d_1324_p3_

                    return lambda76_

                init42_ = lambda75_(p3)
                nw242_ = _dafny.Array(None, 6)
                for i0_42_ in range(nw242_.length(0)):
                    nw242_[i0_42_] = init42_(i0_42_)
                d_1323_v9_ = nw242_
                d_1326_v10_: _dafny.Seq
                d_1326_v10_ = _dafny.SeqWithoutIsStrInference([d_1323_v9_])
                d_1326_v10_ = (d_1326_v10_) + (d_1326_v10_)
                d_1327_v11_: _dafny.Map
                d_1327_v11_ = _dafny.Map({p1: _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('r') for d_1328_i6_ in range(default__.abs(642))])})
                (globalState).f9 = (0) - (default__.safeDivisionInt(d_1318_i4_, default__.safeDivisionInt(len(d_1327_v11_), p2)))
                d_1329_v12_: D0
                d_1329_v12_ = D0_DC0((0) - ((0) - (len(d_1312_v4_))))
                d_1330_v13_: _dafny.MultiSet
                d_1330_v13_ = _dafny.MultiSet([p2])
                d_1331_v14_: _dafny.Seq
                d_1331_v14_ = _dafny.SeqWithoutIsStrInference([p0])
                d_1332_v15_: _dafny.Set
                d_1332_v15_ = _dafny.Set({p2})
                d_1333_v16_: _dafny.Seq
                d_1333_v16_ = _dafny.SeqWithoutIsStrInference([d_1332_v15_])
                d_1334_v17_: _dafny.Map
                d_1334_v17_ = _dafny.Map({len((d_1306_v1_).set(default__.safeIndex(len(d_1331_v14_), len(d_1306_v1_)), _dafny.CodePoint('k'))): d_1323_v9_})
                d_1335_v18_: D8
                d_1335_v18_ = D8_DC19(d_1306_v1_, d_1318_i4_)
                d_1336_v19_: _dafny.Array
                nw243_ = _dafny.Array(None, 26)
                nw243_[int(0)] = d_1318_i4_
                nw243_[int(1)] = p2
                nw243_[int(2)] = (p0) * ((d_1329_v12_).cf0)
                nw243_[int(3)] = p2
                nw243_[int(4)] = p0
                nw243_[int(5)] = d_1318_i4_
                nw243_[int(6)] = ((_dafny.MultiSet([d_1306_v1_])).cardinality) - (p2)
                nw243_[int(7)] = d_1318_i4_
                nw243_[int(8)] = (p0 if self.f21 else d_1318_i4_)
                nw243_[int(9)] = p0
                nw243_[int(10)] = d_1318_i4_
                nw243_[int(11)] = (0) - ((d_1330_v13_).cardinality)
                nw243_[int(12)] = len(d_1331_v14_)
                nw243_[int(13)] = len((_dafny.SeqWithoutIsStrInference([_dafny.Set({d_1318_i4_}) for d_1337_i7_ in range(default__.abs(242))])) + (d_1333_v16_))
                nw243_[int(14)] = len((d_1334_v17_) | (d_1334_v17_))
                nw243_[int(15)] = default__.safeDivisionInt(d_1318_i4_, (d_1335_v18_).cf35)
                nw243_[int(16)] = (d_1330_v13_).cardinality
                nw243_[int(17)] = default__.safeDivisionInt(249, (0) - (p2))
                nw243_[int(18)] = (d_1318_i4_) - (len(d_1306_v1_))
                nw243_[int(19)] = p2
                nw243_[int(20)] = len(d_1306_v1_)
                nw243_[int(21)] = default__.safeDivisionInt(108, len(d_1312_v4_))
                nw243_[int(22)] = (0) - (p0)
                nw243_[int(23)] = p0
                nw243_[int(24)] = p2
                nw243_[int(25)] = default__.fm0((self).f22, globalState)
                d_1336_v19_ = nw243_
                index206_ = default__.safeIndex(90, (d_1336_v19_).length(0))
                (d_1336_v19_)[index206_] = p0
                index207_ = default__.safeIndex(90, (d_1336_v19_).length(0))
                (d_1336_v19_)[index207_] = p2
            elif True:
                d_1338___mcc_h1_ = source13_.cf25
                d_1339_cf25_ = d_1338___mcc_h1_
                d_1340_v20_: _dafny.Array
                nw244_ = _dafny.Array(False, 17)
                d_1340_v20_ = nw244_
                d_1341_v21_: _dafny.Set
                d_1341_v21_ = _dafny.Set({d_1340_v20_})
                rhs232_ = (0) - (p2)
                rhs233_ = (0) - (default__.safeDivisionInt((0) - (p2), len(d_1306_v1_)))
                rhs234_ = (d_1340_v20_) not in (d_1341_v21_)
                rhs235_ = p3
                lhs159_ = globalState
                lhs160_ = globalState
                lhs161_ = globalState
                lhs159_.f9 = rhs232_
                lhs160_.f9 = rhs233_
                lhs161_.f8 = rhs234_
                r0 = rhs235_
                d_1342_v23_: _dafny.MultiSet
                d_1342_v23_ = _dafny.MultiSet([self.f21])
                d_1343_v24_: _dafny.Map
                d_1343_v24_ = _dafny.Map({p3: (self).f22})
                d_1344_v25_: D8
                d_1344_v25_ = D8_DC19(d_1306_v1_, p2)
                d_1345_v26_: _dafny.MultiSet
                d_1345_v26_ = _dafny.MultiSet([p0])
                d_1346_v27_: bool
                out12_: bool
                def iife84_():
                    coll54_ = _dafny.Map()
                    compr_54_: int
                    for compr_54_ in _dafny.IntegerRange(373, -484):
                        d_1347_v22_: int = compr_54_
                        if ((373) <= (d_1347_v22_)) and ((d_1347_v22_) < (-484)):
                            coll54_[default__.safeDivisionInt(d_1347_v22_, p2)] = _dafny.Map({self.f21: p3})
                    return _dafny.Map(coll54_)
                out12_ = (self).m12(((iife84_()
                ).set((d_1342_v23_).cardinality, d_1343_v24_)) == (_dafny.Map({(d_1344_v25_).cf35: d_1343_v24_})), (d_1345_v26_).issubset(d_1345_v26_), d_1317_v6_, globalState)
                d_1346_v27_ = out12_
                (globalState).f8 = not(((self).f22) == (p1))
                d_1348_v28_: _dafny.Map
                d_1348_v28_ = _dafny.Map({((0) - (d_1318_i4_) if (self).f22 else d_1318_i4_): d_1340_v20_})
                d_1349_v29_: str
                d_1349_v29_ = _dafny.CodePoint('l')
                d_1350_v30_: _dafny.Map
                d_1350_v30_ = _dafny.Map({d_1349_v29_: len(d_1306_v1_)})
                d_1348_v28_ = (_dafny.Map({((d_1350_v30_)[d_1349_v29_] if (d_1349_v29_) in (d_1350_v30_) else p2): d_1340_v20_})) | (_dafny.Map({d_1318_i4_: d_1340_v20_}))
            (globalState).f8 = p3
            d_1351_v31_: D4
            d_1351_v31_ = D4_DC12(True, False, (self).f22)
            d_1352_v32_: _dafny.Seq
            d_1352_v32_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tey"))])
            d_1353_v33_: _dafny.Map
            d_1353_v33_ = _dafny.Map({d_1306_v1_: (d_1352_v32_)[default__.safeIndex(p2, len(d_1352_v32_))]})
            (globalState).f8 = ((d_1351_v31_).cf24) and ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pxt"))) <= (((d_1353_v33_)[d_1306_v1_] if (d_1306_v1_) in (d_1353_v33_) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "turdyhf")))))
        d_1354_i8_: int
        d_1354_i8_ = 0
        with _dafny.label("6"):
            while p1:
                with _dafny.c_label("6"):
                    if (d_1354_i8_) >= (100):
                        raise _dafny.Break("6")
                    d_1354_i8_ = (d_1354_i8_) + (1)
                    d_1306_v1_ = d_1306_v1_
                    d_1355_v34_: _dafny.Array
                    def lambda77_(d_1356_i9_):
                        return (self).f22

                    init43_ = lambda77_
                    nw245_ = _dafny.Array(None, 2)
                    for i0_43_ in range(nw245_.length(0)):
                        nw245_[i0_43_] = init43_(i0_43_)
                    d_1355_v34_ = nw245_
                    d_1357_v35_: D7
                    d_1357_v35_ = D7_DC17(d_1306_v1_, d_1355_v34_, (self).f22, 629)
                    d_1358_v36_: _dafny.Map
                    d_1358_v36_ = _dafny.Map({d_1357_v35_: d_1312_v4_})
                    d_1358_v36_ = (d_1358_v36_).set(d_1357_v35_, d_1312_v4_)
                    d_1359_v37_: str
                    d_1359_v37_ = _dafny.CodePoint('e')
                    pat_let_tv12_ = d_1359_v37_
                    def iife85_(_pat_let15_0):
                        def iife86_(d_1360_dt__update__tmp_h0_):
                            def iife87_(_pat_let16_0):
                                def iife88_(d_1361_dt__update_hcf9_h0_):
                                    return D1_DC4(d_1361_dt__update_hcf9_h0_)
                                return iife88_(_pat_let16_0)
                            return iife87_(pat_let_tv12_)
                        return iife86_(_pat_let15_0)
                    source14_ = iife85_(D1_DC4(d_1359_v37_))
                    if source14_.is_DC4:
                        d_1362___mcc_h2_ = source14_.cf9
                        d_1363_cf9_ = d_1362___mcc_h2_
                        d_1364_v38_: _dafny.Map
                        d_1364_v38_ = _dafny.Map({p0: self.f21})
                        d_1365_v39_: _dafny.Map
                        d_1365_v39_ = _dafny.Map({(self).f22: d_1364_v38_})
                        d_1366_v40_: _dafny.Array
                        def lambda78_(d_1367_p2_):
                            def lambda79_(d_1368_i10_):
                                return default__.safeModuloInt(d_1368_i10_, d_1367_p2_)

                            return lambda79_

                        init44_ = lambda78_(p2)
                        nw246_ = _dafny.Array(None, 13)
                        for i0_44_ in range(nw246_.length(0)):
                            nw246_[i0_44_] = init44_(i0_44_)
                        d_1366_v40_ = nw246_
                        d_1369_v41_: _dafny.MultiSet
                        d_1369_v41_ = _dafny.MultiSet([d_1366_v40_])
                        d_1370_v42_: D0
                        d_1370_v42_ = D0_DC1(d_1369_v41_, p2, 777)
                        d_1371_v43_: _dafny.Map
                        d_1371_v43_ = _dafny.Map({(0) - ((len(d_1365_v39_)) - ((d_1370_v42_).cf3)): d_1355_v34_})
                        d_1371_v43_ = (d_1371_v43_).set(default__.fm0(self.f21, globalState), d_1355_v34_)
                        d_1372_v44_: _dafny.Seq
                        d_1372_v44_ = _dafny.SeqWithoutIsStrInference([len(d_1306_v1_), p0, p0, (0) - (p2)])
                        (globalState).f9 = (d_1372_v44_)[default__.safeIndex(p2, len(d_1372_v44_))]
                        (globalState).f9 = p2
                        d_1373_v45_: T1
                        nw247_ = C2()
                        nw247_.ctor__((self).f11)
                        d_1373_v45_ = nw247_
                        index208_ = default__.safeIndex(354, (d_1366_v40_).length(0))
                        (d_1366_v40_)[index208_] = p2
                        index209_ = default__.safeIndex(163, (d_1355_v34_).length(0))
                        (d_1355_v34_)[index209_] = (self).f22
                        d_1374_v46_: _dafny.Set
                        d_1374_v46_ = _dafny.Set({(self).f22})
                        d_1375_v47_: _dafny.MultiSet
                        d_1375_v47_ = _dafny.MultiSet([self.f21])
                        index210_ = default__.safeIndex(354, (d_1366_v40_).length(0))
                        index211_ = default__.safeIndex(163, (d_1355_v34_).length(0))
                        rhs236_ = len((d_1306_v1_) + (d_1306_v1_))
                        rhs237_ = d_1373_v45_
                        rhs238_ = default__.safeModuloInt(((0) - (len(_dafny.SeqWithoutIsStrInference([d_1359_v37_ for d_1376_i11_ in range(default__.abs(108))]))) if True else p0), (0) - (len(((d_1306_v1_).set(default__.safeIndex(p0, len(d_1306_v1_)), d_1363_cf9_)) + (d_1306_v1_))))
                        rhs239_ = default__.fm0(p3, globalState)
                        rhs240_ = ((self.f21) not in (d_1312_v4_) if not((p2) < ((0) - (len(d_1374_v46_)))) else (d_1375_v47_).issubset(d_1375_v47_))
                        lhs162_ = globalState
                        lhs163_ = d_1366_v40_
                        lhs164_ = default__.safeIndex(354, (d_1366_v40_).length(0))
                        lhs165_ = globalState
                        lhs166_ = d_1355_v34_
                        lhs167_ = default__.safeIndex(163, (d_1355_v34_).length(0))
                        lhs162_.f9 = rhs236_
                        d_1373_v45_ = rhs237_
                        lhs163_[lhs164_] = rhs238_
                        lhs165_.f9 = rhs239_
                        lhs166_[lhs167_] = rhs240_
                    elif True:
                        d_1377___mcc_h3_ = source14_.cf8
                        d_1378_cf8_ = d_1377___mcc_h3_
                        d_1379_v48_: _dafny.Set
                        d_1379_v48_ = _dafny.Set({p1, p3, self.f21})
                        d_1380_v49_: _dafny.Map
                        d_1380_v49_ = _dafny.Map({d_1379_v48_: 822})
                        d_1380_v49_ = (d_1380_v49_).set(d_1379_v48_, p2)
                        d_1359_v37_ = d_1359_v37_
                        default__.m0(globalState)
                        d_1381_v50_: D5
                        d_1381_v50_ = D5_DC14(p1)
                        def iife89_():
                            coll55_ = _dafny.Map()
                            compr_55_: int
                            for compr_55_ in _dafny.IntegerRange(412, 572):
                                d_1382_v51_: int = compr_55_
                                if ((412) <= (d_1382_v51_)) and ((d_1382_v51_) < (572)):
                                    coll55_[(d_1382_v51_) - (p2)] = True
                            return _dafny.Map(coll55_)
                        d_1381_v50_ = (default__.fm55(p1, p1, -483, globalState) if (p0) in (iife89_()
                        ) else d_1381_v50_)
                    rhs241_ = self.f21
                    rhs242_ = (p0) - (-712)
                    rhs243_ = p2
                    lhs168_ = globalState
                    lhs169_ = globalState
                    lhs170_ = globalState
                    lhs168_.f8 = rhs241_
                    lhs169_.f9 = rhs242_
                    lhs170_.f9 = rhs243_
                    pass
            pass
        r0 = p1
        d_1383_v52_: _dafny.Array
        nw248_ = _dafny.Array(_dafny.Array(None, 0), 28)
        d_1383_v52_ = nw248_
        r1 = d_1383_v52_
        return r0, r1

    @property
    def f22(self):
        return self._f22

class C4(T0):
    def  __init__(self):
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C4"
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f11):
        (self)._f11 = f11

    def fm41(self, p0, p1, p2, p3, globalState):
        def iife90_():
            coll56_ = _dafny.Map()
            compr_56_: int
            for compr_56_ in _dafny.IntegerRange(674, 977):
                d_1384_v0_: int = compr_56_
                if ((674) <= (d_1384_v0_)) and ((d_1384_v0_) < (977)):
                    coll56_[(d_1384_v0_) + ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s')]))).cardinality)] = (0) - (len(_dafny.Set({(0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rhmhc"))))})))
            return _dafny.Map(coll56_)
        return (len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wmeejk"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "guqpxy"))))) - (((_dafny.MultiSet([(0) - (len(_dafny.Map({400: len(iife90_()
        )})))])).cardinality) - (-830))

    def fm42(self, p0, p1, globalState):
        source15_ = D2_DC7(True, 8, False, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qusyb")))
        if source15_.is_DC6:
            return True
        elif source15_.is_DC7:
            d_1385___mcc_h0_ = source15_.cf11
            d_1386___mcc_h1_ = source15_.cf12
            d_1387___mcc_h2_ = source15_.cf13
            d_1388___mcc_h3_ = source15_.cf14
            d_1389_cf14_ = d_1388___mcc_h3_
            d_1390_cf13_ = d_1387___mcc_h2_
            d_1391_cf12_ = d_1386___mcc_h1_
            d_1392_cf11_ = d_1385___mcc_h0_
            return True
        elif source15_.is_DC8:
            d_1393___mcc_h4_ = source15_.cf15
            d_1394___mcc_h5_ = source15_.cf16
            d_1395_cf16_ = d_1394___mcc_h5_
            d_1396_cf15_ = d_1393___mcc_h4_
            return not (False) or (False)
        elif True:
            d_1397___mcc_h6_ = source15_.cf10
            d_1398_cf10_ = d_1397___mcc_h6_
            if False:
                return False
            elif True:
                return True

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_1399_v0_: _dafny.Seq
        d_1399_v0_ = _dafny.SeqWithoutIsStrInference([p0])
        (globalState).f8 = (p0) in ((d_1399_v0_) + (d_1399_v0_))
        d_1400_v1_: int
        d_1400_v1_ = -850
        d_1401_v2_: _dafny.Set
        d_1401_v2_ = _dafny.Set({(d_1399_v0_)[default__.safeIndex(d_1400_v1_, len(d_1399_v0_))]})
        d_1402_v3_: _dafny.MultiSet
        d_1402_v3_ = _dafny.MultiSet([256])
        d_1403_v4_: _dafny.Map
        d_1403_v4_ = _dafny.Map({(d_1402_v3_).cardinality: len(_dafny.SeqWithoutIsStrInference([not(p0)]))})
        rhs244_ = default__.safeDivisionInt(len((d_1399_v0_) + (d_1399_v0_)), (d_1400_v1_) * (len(d_1401_v2_)))
        rhs245_ = (d_1400_v1_) + (len(d_1403_v4_))
        rhs246_ = d_1400_v1_
        rhs247_ = d_1400_v1_
        lhs171_ = globalState
        r0 = rhs244_
        r0 = rhs245_
        r0 = rhs246_
        lhs171_.f9 = rhs247_
        d_1404_v5_: D5
        d_1404_v5_ = D5_DC14(p0)
        d_1405_v6_: _dafny.Map
        d_1405_v6_ = _dafny.Map({d_1404_v5_: default__.fm0(p0, globalState)})
        (globalState).f9 = (((d_1405_v6_)[D5_DC14(p0)] if (D5_DC14(p0)) in (d_1405_v6_) else d_1400_v1_)) * (-842)
        d_1406_v7_: D0
        d_1406_v7_ = D0_DC0(d_1400_v1_)
        d_1407_v8_: C1
        nw249_ = C1()
        nw249_.ctor__(False, False, d_1406_v7_, p0)
        d_1407_v8_ = nw249_
        d_1408_v9_: _dafny.MultiSet
        d_1408_v9_ = _dafny.MultiSet([d_1407_v8_])
        pat_let_tv13_ = p0
        pat_let_tv14_ = d_1407_v8_
        pat_let_tv15_ = p0
        pat_let_tv16_ = p0
        pat_let_tv17_ = d_1407_v8_
        pat_let_tv18_ = d_1407_v8_
        pat_let_tv19_ = d_1400_v1_
        pat_let_tv20_ = d_1407_v8_
        pat_let_tv21_ = d_1400_v1_
        pat_let_tv22_ = d_1407_v8_
        def lambda80_(source17_):
            if source17_.is_DC11:
                d_1409___mcc_h7_ = source17_.cf19
                d_1410___mcc_h8_ = source17_.cf20
                d_1411___mcc_h9_ = source17_.cf21
                d_1412_cf21_ = d_1411___mcc_h9_
                d_1413_cf20_ = d_1410___mcc_h8_
                d_1414_cf19_ = d_1409___mcc_h7_
                if pat_let_tv13_:
                    return D2_DC7((pat_let_tv14_).f19, (0) - (d_1414_cf19_), pat_let_tv15_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pvjyrtkae")))
                elif True:
                    return D2_DC7(pat_let_tv16_, d_1414_cf19_, (pat_let_tv17_).f20, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aqxjcm")))
            elif source17_.is_DC12:
                d_1415___mcc_h10_ = source17_.cf22
                d_1416___mcc_h11_ = source17_.cf23
                d_1417___mcc_h12_ = source17_.cf24
                d_1418_cf24_ = d_1417___mcc_h12_
                d_1419_cf23_ = d_1416___mcc_h11_
                d_1420_cf22_ = d_1415___mcc_h10_
                return D2_DC7((pat_let_tv18_).f19, pat_let_tv19_, d_1419_cf23_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qx")))
            elif True:
                d_1421___mcc_h13_ = source17_.cf18
                d_1422_cf18_ = d_1421___mcc_h13_
                return D2_DC7((pat_let_tv20_).f19, pat_let_tv21_, (pat_let_tv22_).f19, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('i') for d_1423_i0_ in range(default__.abs(203))]))

        source16_ = lambda80_(default__.fm43(p0, d_1399_v0_, (d_1408_v9_).cardinality, -980, globalState))
        if source16_.is_DC6:
            d_1424_v10_: _dafny.Array
            def lambda81_(d_1425_v1_):
                def lambda82_(d_1426_i1_):
                    return default__.safeDivisionInt(d_1426_i1_, d_1425_v1_)

                return lambda82_

            init45_ = lambda81_(d_1400_v1_)
            nw250_ = _dafny.Array(None, 22)
            for i0_45_ in range(nw250_.length(0)):
                nw250_[i0_45_] = init45_(i0_45_)
            d_1424_v10_ = nw250_
            index212_ = default__.safeIndex(902, (d_1424_v10_).length(0))
            (d_1424_v10_)[index212_] = (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qaxhnnkb")))) + (d_1400_v1_)
            index213_ = default__.safeIndex(902, (d_1424_v10_).length(0))
            (d_1424_v10_)[index213_] = (0) - (d_1400_v1_)
            (globalState).f8 = not ((d_1407_v8_).f20) or ((d_1407_v8_).f19)
            if not((d_1407_v8_).f19):
                (globalState).f9 = (d_1400_v1_) + ((d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))])
                d_1427_v11_: _dafny.Set
                d_1427_v11_ = _dafny.Set({d_1400_v1_, d_1400_v1_})
                r0 = ((D2_DC7((d_1407_v8_).f19, (d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))], False, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('x') for d_1428_i2_ in range(default__.abs(-336))]))).cf12) + ((len(d_1427_v11_)) + (650))
                d_1429_v12_: D8
                d_1429_v12_ = D8_DC18(d_1427_v11_)
                def iife91_():
                    coll57_ = _dafny.Set()
                    compr_57_: int
                    for compr_57_ in ((d_1429_v12_).cf33).Elements:
                        d_1430_v13_: int = compr_57_
                        if (d_1430_v13_) in ((d_1429_v12_).cf33):
                            coll57_ = coll57_.union(_dafny.Set([(d_1430_v13_) - (len(_dafny.SeqWithoutIsStrInference([len(_dafny.Set({len(_dafny.SeqWithoutIsStrInference([False]))})), -435])))]))
                    return _dafny.Set(coll57_)
                d_1427_v11_ = (iife91_()
                ) - (d_1427_v11_)
                d_1431_v14_: _dafny.Array
                def lambda83_(d_1432_v8_):
                    def lambda84_(d_1433_i3_):
                        return (d_1432_v8_).f19

                    return lambda84_

                init46_ = lambda83_(d_1407_v8_)
                nw251_ = _dafny.Array(None, 10)
                for i0_46_ in range(nw251_.length(0)):
                    nw251_[i0_46_] = init46_(i0_46_)
                d_1431_v14_ = nw251_
                d_1434_v15_: _dafny.Seq
                d_1434_v15_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "e"))
                d_1431_v14_ = (D7_DC17(d_1434_v15_, d_1431_v14_, (d_1407_v8_).f19, d_1400_v1_)).cf30
                d_1435_v16_: str
                d_1435_v16_ = _dafny.CodePoint('f')
                d_1436_v17_: _dafny.Map
                d_1436_v17_ = _dafny.Map({(self).fm42((d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))], d_1400_v1_, globalState): d_1400_v1_})
                d_1437_v18_: _dafny.Seq
                d_1437_v18_ = _dafny.SeqWithoutIsStrInference([len(d_1436_v17_)])
                d_1438_v19_: _dafny.MultiSet
                d_1438_v19_ = _dafny.MultiSet([d_1431_v14_, d_1431_v14_])
                rhs248_ = default__.safeDivisionInt(((self).fm41(d_1434_v15_, len(d_1437_v18_), d_1400_v1_, (d_1407_v8_).f20, globalState) if not(not(False)) else (d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))]), (d_1438_v19_).cardinality)
                rhs249_ = d_1435_v16_
                lhs172_ = globalState
                lhs172_.f9 = rhs248_
                d_1435_v16_ = rhs249_
            elif True:
                d_1439_v20_: str
                d_1439_v20_ = _dafny.CodePoint('v')
                d_1439_v20_ = _dafny.CodePoint('l')
                d_1440_v21_: D2
                d_1440_v21_ = D2_DC5(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kqmqprqxw")))
                d_1441_v22_: _dafny.Map
                d_1441_v22_ = _dafny.Map({False: (d_1407_v8_).f20})
                d_1442_v23_: _dafny.Seq
                d_1442_v23_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gx"))
                d_1443_v24_: _dafny.Seq
                d_1443_v24_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1439_v20_ for d_1444_i4_ in range(default__.abs(624))]), d_1442_v23_])
                d_1445_v25_: D1
                d_1445_v25_ = D1_DC4(d_1439_v20_)
                d_1446_v26_: _dafny.Set
                d_1446_v26_ = _dafny.Set({len(_dafny.SeqWithoutIsStrInference([d_1439_v20_ for d_1447_i5_ in range(default__.abs(987))]))})
                d_1448_v27_: _dafny.Map
                d_1448_v27_ = _dafny.Map({(d_1443_v24_)[default__.safeIndex((d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))], len(d_1443_v24_))]: (d_1407_v8_).fm9(d_1445_v25_, d_1446_v26_, globalState)})
                d_1449_v28_: _dafny.Array
                nw252_ = _dafny.Array(None, 9)
                nw252_[int(0)] = (d_1407_v8_).f20
                nw252_[int(1)] = (d_1407_v8_).f19
                nw252_[int(2)] = (d_1407_v8_).f20
                nw252_[int(3)] = p0
                nw252_[int(4)] = not((d_1407_v8_).f19)
                nw252_[int(5)] = (d_1407_v8_).f20
                nw252_[int(6)] = (d_1407_v8_).f19
                nw252_[int(7)] = (d_1407_v8_).f20
                nw252_[int(8)] = ((0) - (len(d_1441_v22_))) < (((d_1448_v27_)[_dafny.SeqWithoutIsStrInference([d_1439_v20_ for d_1450_i6_ in range(default__.abs(46))])] if (_dafny.SeqWithoutIsStrInference([d_1439_v20_ for d_1451_i6_ in range(default__.abs(46))])) in (d_1448_v27_) else d_1400_v1_))
                d_1449_v28_ = nw252_
                d_1452_v29_: _dafny.Seq
                d_1452_v29_ = _dafny.SeqWithoutIsStrInference([(d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))]])
                index214_ = default__.safeIndex(932, (d_1449_v28_).length(0))
                (d_1449_v28_)[index214_] = ((d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))]) > (len((d_1452_v29_).set(default__.safeIndex(d_1400_v1_, len(d_1452_v29_)), len(d_1403_v4_))))
                index215_ = default__.safeIndex(932, (d_1449_v28_).length(0))
                rhs250_ = d_1440_v21_
                rhs251_ = ((d_1400_v1_ if False else (d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))])) > (default__.safeDivisionInt(d_1400_v1_, len(_dafny.SeqWithoutIsStrInference([d_1400_v1_]))))
                rhs252_ = d_1439_v20_
                rhs253_ = (d_1400_v1_ if (d_1407_v8_).f19 else 837)
                rhs254_ = (((d_1399_v0_) + (_dafny.SeqWithoutIsStrInference([p0, (d_1407_v8_).f20]))).set(default__.safeIndex(-896, len((d_1399_v0_) + (_dafny.SeqWithoutIsStrInference([p0, (d_1407_v8_).f20])))), (d_1407_v8_).f19)) == (d_1399_v0_)
                lhs173_ = d_1449_v28_
                lhs174_ = default__.safeIndex(932, (d_1449_v28_).length(0))
                lhs175_ = globalState
                d_1440_v21_ = rhs250_
                lhs173_[lhs174_] = rhs251_
                d_1439_v20_ = rhs252_
                r0 = rhs253_
                lhs175_.f8 = rhs254_
                d_1453_v30_: _dafny.Array
                def lambda85_(d_1454_v10_, d_1455_v29_):
                    def lambda86_(d_1456_i7_):
                        return (_dafny.SeqWithoutIsStrInference([(d_1454_v10_)[default__.safeIndex(902, (d_1454_v10_).length(0))] for d_1457_i8_ in range(default__.abs(298))])) + (d_1455_v29_)

                    return lambda86_

                init47_ = lambda85_(d_1424_v10_, d_1452_v29_)
                nw253_ = _dafny.Array(None, 16)
                for i0_47_ in range(nw253_.length(0)):
                    nw253_[i0_47_] = init47_(i0_47_)
                d_1453_v30_ = nw253_
                index216_ = default__.safeIndex(631, (d_1453_v30_).length(0))
                (d_1453_v30_)[index216_] = ((_dafny.SeqWithoutIsStrInference([d_1400_v1_])).set(default__.safeIndex(d_1400_v1_, len(_dafny.SeqWithoutIsStrInference([d_1400_v1_]))), (d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))])) + (_dafny.SeqWithoutIsStrInference([(d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))] for d_1458_i9_ in range(default__.abs(464))]))
                index217_ = default__.safeIndex(631, (d_1453_v30_).length(0))
                (d_1453_v30_)[index217_] = (d_1452_v29_) + (d_1452_v29_)
                d_1459_v31_: C2
                nw254_ = C2()
                nw254_.ctor__((self).f11)
                d_1459_v31_ = nw254_
                d_1460_v32_: int
                d_1461_v33_: _dafny.MultiSet
                out13_: int
                out14_: _dafny.MultiSet
                out13_, out14_ = (d_1407_v8_).m9(_dafny.MultiSet([(d_1453_v30_)[default__.safeIndex(631, (d_1453_v30_).length(0))]]), globalState)
                d_1460_v32_ = out13_
                d_1461_v33_ = out14_
            if not (not((d_1407_v8_).f20)) or (False):
                d_1462_v34_: str
                d_1462_v34_ = _dafny.CodePoint('v')
                d_1463_v35_: _dafny.Array
                nw255_ = _dafny.Array(None, 21)
                nw255_[int(0)] = d_1462_v34_
                nw255_[int(1)] = d_1462_v34_
                nw255_[int(2)] = d_1462_v34_
                nw255_[int(3)] = d_1462_v34_
                nw255_[int(4)] = d_1462_v34_
                nw255_[int(5)] = d_1462_v34_
                nw255_[int(6)] = d_1462_v34_
                nw255_[int(7)] = d_1462_v34_
                nw255_[int(8)] = d_1462_v34_
                nw255_[int(9)] = default__.fm3((d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))], (d_1407_v8_).f19, globalState)
                nw255_[int(10)] = d_1462_v34_
                nw255_[int(11)] = _dafny.CodePoint('i')
                nw255_[int(12)] = d_1462_v34_
                nw255_[int(13)] = d_1462_v34_
                nw255_[int(14)] = d_1462_v34_
                nw255_[int(15)] = d_1462_v34_
                nw255_[int(16)] = d_1462_v34_
                nw255_[int(17)] = d_1462_v34_
                nw255_[int(18)] = d_1462_v34_
                nw255_[int(19)] = default__.fm3((0) - (-858), (d_1407_v8_).f20, globalState)
                nw255_[int(20)] = d_1462_v34_
                d_1463_v35_ = nw255_
                index218_ = default__.safeIndex(84, (d_1463_v35_).length(0))
                (d_1463_v35_)[index218_] = d_1462_v34_
                d_1464_v36_: D1
                d_1464_v36_ = D1_DC4(d_1462_v34_)
                d_1465_v37_: _dafny.Set
                d_1465_v37_ = _dafny.Set({(d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))]})
                index219_ = default__.safeIndex(84, (d_1463_v35_).length(0))
                rhs255_ = d_1462_v34_
                rhs256_ = (d_1407_v8_).f20
                rhs257_ = (d_1407_v8_).fm9(d_1464_v36_, d_1465_v37_, globalState)
                rhs258_ = (d_1407_v8_).f19
                lhs176_ = d_1463_v35_
                lhs177_ = default__.safeIndex(84, (d_1463_v35_).length(0))
                lhs178_ = globalState
                lhs179_ = globalState
                lhs176_[lhs177_] = rhs255_
                lhs178_.f8 = rhs256_
                r0 = rhs257_
                lhs179_.f8 = rhs258_
                d_1462_v34_ = (d_1463_v35_)[default__.safeIndex(84, (d_1463_v35_).length(0))]
                (globalState).f9 = (d_1400_v1_) + (d_1400_v1_)
                d_1400_v1_ = (d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))]
                d_1466_v38_: _dafny.Set
                d_1466_v38_ = _dafny.Set({d_1424_v10_, d_1424_v10_, d_1424_v10_, d_1424_v10_, d_1424_v10_})
                index220_ = default__.safeIndex(902, (d_1424_v10_).length(0))
                (d_1424_v10_)[index220_] = len(d_1466_v38_)
            elif True:
                (globalState).f8 = p0
                d_1467_v39_: D8
                d_1467_v39_ = D8_DC20(D0_DC2(_dafny.MultiSet([False, (d_1407_v8_).f19]), (d_1407_v8_).f19, d_1400_v1_, (d_1407_v8_).f19), True, (d_1407_v8_).f19)
                d_1468_v40_: _dafny.Seq
                d_1468_v40_ = _dafny.SeqWithoutIsStrInference([d_1467_v39_])
                d_1469_v41_: _dafny.Seq
                d_1469_v41_ = _dafny.SeqWithoutIsStrInference([len(d_1468_v40_)])
                d_1469_v41_ = (d_1469_v41_) + (d_1469_v41_)
                r0 = default__.safeModuloInt(len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "psom"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jjhaifpp")))), d_1400_v1_)
                d_1470_v42_: _dafny.Seq
                d_1470_v42_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lfja"))
                d_1471_v43_: str
                d_1471_v43_ = _dafny.CodePoint('l')
                d_1472_v44_: _dafny.Map
                d_1472_v44_ = _dafny.Map({d_1470_v42_: d_1471_v43_})
                d_1473_v45_: _dafny.Map
                d_1473_v45_ = _dafny.Map({(d_1407_v8_).f20: _dafny.MultiSet(d_1469_v41_)})
                d_1474_v46_: _dafny.Map
                d_1474_v46_ = _dafny.Map({(d_1407_v8_).f19: default__.fm0((d_1407_v8_).f19, globalState)})
                def iife92_():
                    coll58_ = _dafny.Map()
                    compr_58_: int
                    for compr_58_ in (d_1403_v4_).keys.Elements:
                        d_1475_v47_: int = compr_58_
                        if (d_1475_v47_) in (d_1403_v4_):
                            coll58_[(d_1475_v47_) - ((d_1424_v10_)[default__.safeIndex(902, (d_1424_v10_).length(0))])] = d_1400_v1_
                    return _dafny.Map(coll58_)
                d_1472_v44_ = (d_1472_v44_).set(default__.fm2(d_1400_v1_, ((d_1473_v45_)[p0] if (p0) in (d_1473_v45_) else _dafny.MultiSet((d_1469_v41_).set(default__.safeIndex(((d_1474_v46_)[(d_1407_v8_).f19] if ((d_1407_v8_).f19) in (d_1474_v46_) else d_1400_v1_), len(d_1469_v41_)), d_1400_v1_))), iife92_()
                , globalState), d_1471_v43_)
                d_1476_v48_: D8
                d_1476_v48_ = D8_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")), 297)
                d_1477_v49_: D2
                d_1477_v49_ = D2_DC8(d_1400_v1_, (d_1470_v42_)[default__.safeIndex(((d_1403_v4_)[d_1400_v1_] if (d_1400_v1_) in (d_1403_v4_) else (d_1476_v48_).cf35), len(d_1470_v42_))])
                rhs259_ = (d_1400_v1_) + (d_1400_v1_)
                rhs260_ = (d_1477_v49_).cf15
                rhs261_ = ((_dafny.SeqWithoutIsStrInference([d_1399_v0_ for d_1478_i10_ in range(default__.abs(460))])) < (_dafny.SeqWithoutIsStrInference([d_1399_v0_])) if (d_1407_v8_).f19 else (d_1407_v8_).f20)
                lhs180_ = globalState
                lhs181_ = globalState
                lhs182_ = globalState
                lhs180_.f9 = rhs259_
                lhs181_.f9 = rhs260_
                lhs182_.f8 = rhs261_
        elif source16_.is_DC7:
            d_1479___mcc_h0_ = source16_.cf11
            d_1480___mcc_h1_ = source16_.cf12
            d_1481___mcc_h2_ = source16_.cf13
            d_1482___mcc_h3_ = source16_.cf14
            d_1483_cf14_ = d_1482___mcc_h3_
            d_1484_cf13_ = d_1481___mcc_h2_
            d_1485_cf12_ = d_1480___mcc_h1_
            d_1486_cf11_ = d_1479___mcc_h0_
            d_1487_v50_: _dafny.Array
            nw256_ = _dafny.Array(False, 7)
            d_1487_v50_ = nw256_
            d_1488_v51_: str
            d_1488_v51_ = _dafny.CodePoint('p')
            index221_ = default__.safeIndex(195, (d_1487_v50_).length(0))
            (d_1487_v50_)[index221_] = default__.fm1(d_1486_cf11_, d_1488_v51_, globalState)
            index222_ = default__.safeIndex(195, (d_1487_v50_).length(0))
            (d_1487_v50_)[index222_] = not ((d_1407_v8_).f19) or ((d_1407_v8_).f19)
            r0 = d_1400_v1_
            d_1489_v52_: _dafny.Seq
            d_1489_v52_ = _dafny.SeqWithoutIsStrInference([d_1485_cf12_])
            d_1489_v52_ = (d_1489_v52_) + (_dafny.SeqWithoutIsStrInference([(0) - ((0) - (d_1485_cf12_)) for d_1490_i11_ in range(default__.abs(335))]))
            d_1483_cf14_ = _dafny.SeqWithoutIsStrInference([d_1488_v51_ for d_1491_i12_ in range(default__.abs(729))])
        elif source16_.is_DC8:
            d_1492___mcc_h4_ = source16_.cf15
            d_1493___mcc_h5_ = source16_.cf16
            d_1494_cf16_ = d_1493___mcc_h5_
            d_1495_cf15_ = d_1492___mcc_h4_
            d_1496_v53_: _dafny.Array
            def lambda87_(d_1497_v8_, d_1498_cf15_):
                def lambda88_(d_1499_i13_):
                    return (D8_DC20(D0_DC2(_dafny.MultiSet([(d_1497_v8_).f20, (d_1497_v8_).f20]), (d_1497_v8_).f19, d_1498_cf15_, (d_1497_v8_).f20), (d_1497_v8_).f20, True)).cf37

                return lambda88_

            init48_ = lambda87_(d_1407_v8_, d_1495_cf15_)
            nw257_ = _dafny.Array(None, 1)
            for i0_48_ in range(nw257_.length(0)):
                nw257_[i0_48_] = init48_(i0_48_)
            d_1496_v53_ = nw257_
            d_1500_v54_: _dafny.Array
            nw258_ = _dafny.Array(int(0), 8)
            d_1500_v54_ = nw258_
            index223_ = default__.safeIndex(384, (d_1500_v54_).length(0))
            (d_1500_v54_)[index223_] = 455
            d_1501_v55_: D1
            d_1501_v55_ = D1_DC4(d_1494_cf16_)
            d_1502_v56_: _dafny.Set
            d_1502_v56_ = _dafny.Set({d_1495_cf15_})
            index224_ = default__.safeIndex(384, (d_1500_v54_).length(0))
            rhs262_ = (d_1496_v53_ if False else d_1496_v53_)
            rhs263_ = (d_1407_v8_).fm9(d_1501_v55_, d_1502_v56_, globalState)
            lhs183_ = d_1500_v54_
            lhs184_ = default__.safeIndex(384, (d_1500_v54_).length(0))
            d_1496_v53_ = rhs262_
            lhs183_[lhs184_] = rhs263_
            (globalState).f9 = (0) - (len((d_1401_v2_).intersection(d_1401_v2_)))
            d_1503_v57_: _dafny.Array
            def lambda89_(d_1504_v8_, d_1505_cf15_, d_1506_p0_, d_1507_v54_):
                def lambda90_(d_1508_i14_):
                    return (D2_DC7((d_1504_v8_).f19, d_1505_cf15_, d_1506_p0_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wpk"))) if (d_1504_v8_).f20 else D2_DC7((d_1504_v8_).f19, (d_1507_v54_)[default__.safeIndex(384, (d_1507_v54_).length(0))], True, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mynlae"))))

                return lambda90_

            init49_ = lambda89_(d_1407_v8_, d_1495_cf15_, p0, d_1500_v54_)
            nw259_ = _dafny.Array(None, 1)
            for i0_49_ in range(nw259_.length(0)):
                nw259_[i0_49_] = init49_(i0_49_)
            d_1503_v57_ = nw259_
            d_1509_v58_: D2
            d_1509_v58_ = D2_DC7((d_1407_v8_).f20, 773, (d_1407_v8_).f19, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hticdw")))
            index225_ = default__.safeIndex(265, (d_1503_v57_).length(0))
            (d_1503_v57_)[index225_] = d_1509_v58_
            d_1510_v59_: _dafny.Seq
            d_1510_v59_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qye"))
            pat_let_tv23_ = d_1500_v54_
            pat_let_tv24_ = d_1500_v54_
            pat_let_tv25_ = d_1407_v8_
            pat_let_tv26_ = d_1510_v59_
            index226_ = default__.safeIndex(265, (d_1503_v57_).length(0))
            def iife93_(_pat_let17_0):
                def iife94_(d_1512_dt__update__tmp_h0_):
                    def iife95_(_pat_let18_0):
                        def iife96_(d_1513_dt__update_hcf12_h0_):
                            def iife97_(_pat_let19_0):
                                def iife98_(d_1514_dt__update_hcf11_h0_):
                                    def iife99_(_pat_let20_0):
                                        def iife100_(d_1515_dt__update_hcf14_h0_):
                                            return D2_DC7(d_1514_dt__update_hcf11_h0_, d_1513_dt__update_hcf12_h0_, (d_1512_dt__update__tmp_h0_).cf13, d_1515_dt__update_hcf14_h0_)
                                        return iife100_(_pat_let20_0)
                                    return iife99_(pat_let_tv26_)
                                return iife98_(_pat_let19_0)
                            return iife97_((pat_let_tv25_).f20)
                        return iife96_(_pat_let18_0)
                    return iife95_((pat_let_tv24_)[default__.safeIndex(384, (pat_let_tv23_).length(0))])
                return iife94_(_pat_let17_0)
            (d_1503_v57_)[index226_] = iife93_((D2_DC7(not((d_1407_v8_).f19), (d_1500_v54_)[default__.safeIndex(384, (d_1500_v54_).length(0))], (d_1407_v8_).f20, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1511_i15_ in range(default__.abs(-908))])) if (d_1407_v8_).f19 else d_1509_v58_))
            d_1516_v60_: _dafny.MultiSet
            d_1516_v60_ = _dafny.MultiSet([(d_1407_v8_).f19, False])
            d_1517_v61_: D0
            d_1517_v61_ = D0_DC2(d_1516_v60_, True, d_1495_cf15_, p0)
            d_1518_v62_: _dafny.Seq
            d_1518_v62_ = _dafny.SeqWithoutIsStrInference([d_1400_v1_])
            d_1519_v63_: D8
            d_1519_v63_ = D8_DC20(d_1517_v61_, (d_1518_v62_) != (d_1518_v62_), (d_1407_v8_).f20)
            source18_ = d_1519_v63_
            if source18_.is_DC19:
                d_1520___mcc_h14_ = source18_.cf34
                d_1521___mcc_h15_ = source18_.cf35
                d_1522_cf35_ = d_1521___mcc_h15_
                d_1523_cf34_ = d_1520___mcc_h14_
                d_1495_cf15_ = (d_1495_cf15_) + (d_1495_cf15_)
                d_1524_v64_: _dafny.Map
                d_1524_v64_ = _dafny.Map({(d_1407_v8_).f20: not (False) or (True)})
                d_1524_v64_ = (d_1524_v64_).set(False, (d_1407_v8_).f20)
                d_1494_cf16_ = d_1494_cf16_
                d_1407_v8_ = d_1407_v8_
            elif source18_.is_DC20:
                d_1525___mcc_h16_ = source18_.cf36
                d_1526___mcc_h17_ = source18_.cf37
                d_1527___mcc_h18_ = source18_.cf38
                d_1528_cf38_ = d_1527___mcc_h18_
                d_1529_cf37_ = d_1526___mcc_h17_
                d_1530_cf36_ = d_1525___mcc_h16_
                index227_ = default__.safeIndex(495, (d_1496_v53_).length(0))
                (d_1496_v53_)[index227_] = (default__.fm0(not((d_1407_v8_).f19), globalState)) == (d_1400_v1_)
                index228_ = default__.safeIndex(495, (d_1496_v53_).length(0))
                (d_1496_v53_)[index228_] = ((0) - (d_1400_v1_)) not in (d_1518_v62_)
                d_1531_v65_: _dafny.Array
                nw260_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
                d_1531_v65_ = nw260_
                index229_ = default__.safeIndex(58, (d_1531_v65_).length(0))
                (d_1531_v65_)[index229_] = _dafny.SeqWithoutIsStrInference([d_1494_cf16_ for d_1532_i16_ in range(default__.abs(-863))])
                index230_ = default__.safeIndex(58, (d_1531_v65_).length(0))
                (d_1531_v65_)[index230_] = d_1510_v59_
                d_1400_v1_ = (d_1500_v54_)[default__.safeIndex(384, (d_1500_v54_).length(0))]
                d_1533_v66_: C1
                nw261_ = C1()
                nw261_.ctor__((d_1407_v8_).f20, (self).fm42((d_1500_v54_)[default__.safeIndex(384, (d_1500_v54_).length(0))], len((d_1531_v65_)[default__.safeIndex(58, (d_1531_v65_).length(0))]), globalState), D0_DC0(136), True)
                d_1533_v66_ = nw261_
            elif True:
                d_1534___mcc_h19_ = source18_.cf33
                d_1535_cf33_ = d_1534___mcc_h19_
                (globalState).f8 = (d_1407_v8_).f19
                (globalState).f8 = (d_1407_v8_).f19
                d_1536_v67_: _dafny.Array
                def lambda91_(d_1537_v8_, d_1538_v62_):
                    def lambda92_(d_1539_i17_):
                        return _dafny.Map({not((d_1537_v8_).f19): len(d_1538_v62_)})

                    return lambda92_

                init50_ = lambda91_(d_1407_v8_, d_1518_v62_)
                nw262_ = _dafny.Array(None, 24)
                for i0_50_ in range(nw262_.length(0)):
                    nw262_[i0_50_] = init50_(i0_50_)
                d_1536_v67_ = nw262_
                d_1536_v67_ = d_1536_v67_
                index231_ = default__.safeIndex(384, (d_1500_v54_).length(0))
                (d_1500_v54_)[index231_] = (d_1400_v1_) * (len(d_1518_v62_))
        elif True:
            d_1540___mcc_h6_ = source16_.cf10
            d_1541_cf10_ = d_1540___mcc_h6_
            d_1542_v68_: _dafny.Array
            nw263_ = _dafny.Array(False, 27)
            d_1542_v68_ = nw263_
            index232_ = default__.safeIndex(162, (d_1542_v68_).length(0))
            (d_1542_v68_)[index232_] = (d_1407_v8_).f20
            d_1543_v69_: _dafny.MultiSet
            d_1543_v69_ = _dafny.MultiSet([p0])
            d_1544_v70_: str
            d_1544_v70_ = _dafny.CodePoint('f')
            d_1545_v71_: _dafny.Map
            d_1545_v71_ = _dafny.Map({d_1400_v1_: d_1544_v70_})
            d_1546_v72_: D2
            d_1546_v72_ = D2_DC8(len(_dafny.Map({d_1400_v1_: d_1400_v1_})), d_1544_v70_)
            d_1547_v73_: _dafny.Map
            d_1547_v73_ = _dafny.Map({d_1400_v1_: (d_1407_v8_).f19})
            d_1548_v74_: _dafny.Array
            nw264_ = _dafny.Array(None, 11)
            nw264_[int(0)] = d_1400_v1_
            nw264_[int(1)] = (0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "abdf"))))
            nw264_[int(2)] = (0) - ((d_1546_v72_).cf15)
            nw264_[int(3)] = (48) + (d_1400_v1_)
            nw264_[int(4)] = (d_1543_v69_).cardinality
            nw264_[int(5)] = d_1400_v1_
            nw264_[int(6)] = (0) - (default__.safeDivisionInt(((d_1402_v3_)[d_1400_v1_] if (d_1400_v1_) in (d_1402_v3_) else d_1400_v1_), d_1400_v1_))
            nw264_[int(7)] = d_1400_v1_
            nw264_[int(8)] = len(_dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "h")), d_1541_cf10_]))
            nw264_[int(9)] = (d_1407_v8_).fm30(d_1547_v73_, globalState)
            nw264_[int(10)] = (d_1400_v1_) + (d_1400_v1_)
            d_1548_v74_ = nw264_
            index233_ = default__.safeIndex(793, (d_1548_v74_).length(0))
            (d_1548_v74_)[index233_] = ((0) - (d_1400_v1_)) * (d_1400_v1_)
            d_1549_v75_: _dafny.Seq
            d_1549_v75_ = _dafny.SeqWithoutIsStrInference([d_1541_cf10_, _dafny.SeqWithoutIsStrInference([d_1544_v70_ for d_1550_i18_ in range(default__.abs(425))])])
            index234_ = default__.safeIndex(162, (d_1542_v68_).length(0))
            index235_ = default__.safeIndex(793, (d_1548_v74_).length(0))
            rhs264_ = p0
            rhs265_ = d_1543_v69_
            rhs266_ = d_1542_v68_
            rhs267_ = d_1545_v71_
            rhs268_ = default__.safeDivisionInt(len(d_1549_v75_), d_1400_v1_)
            lhs185_ = d_1542_v68_
            lhs186_ = default__.safeIndex(162, (d_1542_v68_).length(0))
            lhs187_ = d_1548_v74_
            lhs188_ = default__.safeIndex(793, (d_1548_v74_).length(0))
            lhs185_[lhs186_] = rhs264_
            d_1543_v69_ = rhs265_
            d_1542_v68_ = rhs266_
            d_1545_v71_ = rhs267_
            lhs187_[lhs188_] = rhs268_
            index236_ = default__.safeIndex(793, (d_1548_v74_).length(0))
            (d_1548_v74_)[index236_] = (305 if (d_1407_v8_).f20 else (825) * (318))
            d_1551_v76_: _dafny.Map
            d_1551_v76_ = _dafny.Map({(d_1400_v1_) + (d_1400_v1_): d_1402_v3_})
            d_1552_v77_: _dafny.Array
            def lambda93_(d_1553_v8_):
                def lambda94_(d_1554_i19_):
                    return _dafny.SeqWithoutIsStrInference([True, (d_1553_v8_).f19, (d_1553_v8_).f19, (d_1553_v8_).f19])

                return lambda94_

            init51_ = lambda93_(d_1407_v8_)
            nw265_ = _dafny.Array(None, 22)
            for i0_51_ in range(nw265_.length(0)):
                nw265_[i0_51_] = init51_(i0_51_)
            d_1552_v77_ = nw265_
            d_1555_v78_: _dafny.Seq
            d_1555_v78_ = _dafny.SeqWithoutIsStrInference([d_1399_v0_, d_1399_v0_, _dafny.SeqWithoutIsStrInference([(d_1407_v8_).f20])])
            index237_ = default__.safeIndex(744, (d_1552_v77_).length(0))
            (d_1552_v77_)[index237_] = default__.fm44(d_1555_v78_, d_1544_v70_, globalState)
            d_1556_v79_: _dafny.Array
            nw266_ = _dafny.Array(_dafny.Array(None, 0), 6)
            d_1556_v79_ = nw266_
            d_1557_v80_: _dafny.Array
            def lambda95_(d_1558_v2_):
                def lambda96_(d_1559_i20_):
                    return d_1558_v2_

                return lambda96_

            init52_ = lambda95_(d_1401_v2_)
            nw267_ = _dafny.Array(None, 9)
            for i0_52_ in range(nw267_.length(0)):
                nw267_[i0_52_] = init52_(i0_52_)
            d_1557_v80_ = nw267_
            index238_ = default__.safeIndex(658, (d_1556_v79_).length(0))
            (d_1556_v79_)[index238_] = d_1557_v80_
            d_1560_v81_: _dafny.Map
            d_1560_v81_ = _dafny.Map({d_1548_v74_: (d_1407_v8_).f20})
            index239_ = default__.safeIndex(744, (d_1552_v77_).length(0))
            index240_ = default__.safeIndex(658, (d_1556_v79_).length(0))
            rhs269_ = d_1551_v76_
            rhs270_ = d_1399_v0_
            rhs271_ = ((d_1560_v81_)[d_1548_v74_] if (d_1548_v74_) in (d_1560_v81_) else (d_1407_v8_).f19)
            rhs272_ = d_1557_v80_
            lhs189_ = d_1552_v77_
            lhs190_ = default__.safeIndex(744, (d_1552_v77_).length(0))
            lhs191_ = globalState
            lhs192_ = d_1556_v79_
            lhs193_ = default__.safeIndex(658, (d_1556_v79_).length(0))
            d_1551_v76_ = rhs269_
            lhs189_[lhs190_] = rhs270_
            lhs191_.f8 = rhs271_
            lhs192_[lhs193_] = rhs272_
            if ((d_1552_v77_)[default__.safeIndex(744, (d_1552_v77_).length(0))])[default__.safeIndex(((d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))]) * ((d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))]), len((d_1552_v77_)[default__.safeIndex(744, (d_1552_v77_).length(0))]))]:
                d_1561_v82_: C0
                nw268_ = C0()
                nw268_.ctor__(((d_1407_v8_).f20 if (d_1407_v8_).f19 else (d_1407_v8_).f19))
                d_1561_v82_ = nw268_
                index241_ = default__.safeIndex(793, (d_1548_v74_).length(0))
                (d_1548_v74_)[index241_] = d_1400_v1_
                d_1562_v83_: _dafny.Array
                nw269_ = _dafny.Array(D0.default()(), 17)
                d_1562_v83_ = nw269_
                d_1563_v84_: D0
                d_1563_v84_ = D0_DC2(_dafny.MultiSet([(d_1407_v8_).f20, (d_1561_v82_).f17]), (d_1407_v8_).f20, (d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))], (d_1407_v8_).f20)
                index242_ = default__.safeIndex(313, (d_1562_v83_).length(0))
                (d_1562_v83_)[index242_] = (d_1563_v84_ if (d_1561_v82_).f17 else d_1563_v84_)
                index243_ = default__.safeIndex(313, (d_1562_v83_).length(0))
                (d_1562_v83_)[index243_] = d_1563_v84_
                (globalState).f8 = p0
                d_1564_v85_: _dafny.Map
                d_1564_v85_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qwbav"))): (d_1549_v75_)[default__.safeIndex((d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))], len(d_1549_v75_))]})
                d_1565_v86_: D9
                d_1565_v86_ = D9_DC22(d_1400_v1_, d_1564_v85_)
                (globalState).f9 = (d_1565_v86_).cf40
            elif True:
                pat_let_tv27_ = d_1400_v1_
                index244_ = default__.safeIndex(793, (d_1548_v74_).length(0))
                index245_ = default__.safeIndex(744, (d_1552_v77_).length(0))
                def iife101_(_pat_let21_0):
                    def iife102_(d_1567_dt__update__tmp_h1_):
                        def iife103_(_pat_let22_0):
                            def iife104_(d_1568_dt__update_hcf15_h0_):
                                return D2_DC8(d_1568_dt__update_hcf15_h0_, (d_1567_dt__update__tmp_h1_).cf16)
                            return iife104_(_pat_let22_0)
                        return iife103_(pat_let_tv27_)
                    return iife102_(_pat_let21_0)
                rhs273_ = len(_dafny.SeqWithoutIsStrInference([_dafny.Map({p0: (d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))]}) for d_1566_i21_ in range(default__.abs(816))]))
                rhs274_ = (len(d_1547_v73_)) + (606)
                rhs275_ = _dafny.SeqWithoutIsStrInference([False])
                rhs276_ = (iife101_(d_1546_v72_)).cf15
                lhs194_ = d_1548_v74_
                lhs195_ = default__.safeIndex(793, (d_1548_v74_).length(0))
                lhs196_ = globalState
                lhs197_ = d_1552_v77_
                lhs198_ = default__.safeIndex(744, (d_1552_v77_).length(0))
                lhs199_ = globalState
                lhs194_[lhs195_] = rhs273_
                lhs196_.f9 = rhs274_
                lhs197_[lhs198_] = rhs275_
                lhs199_.f9 = rhs276_
                d_1569_v87_: D4
                d_1569_v87_ = D4_DC12(True, (d_1407_v8_).f20, (d_1407_v8_).f20)
                d_1570_v88_: _dafny.MultiSet
                d_1570_v88_ = _dafny.MultiSet([d_1569_v87_, d_1569_v87_])
                d_1571_v89_: D0
                d_1571_v89_ = D0_DC2(_dafny.MultiSet((d_1552_v77_)[default__.safeIndex(744, (d_1552_v77_).length(0))]), (d_1542_v68_)[default__.safeIndex(162, (d_1542_v68_).length(0))], (d_1570_v88_).cardinality, p0)
                d_1572_v90_: _dafny.MultiSet
                d_1572_v90_ = _dafny.MultiSet([d_1571_v89_, d_1571_v89_, d_1571_v89_])
                (globalState).f9 = ((d_1572_v90_).cardinality) + (d_1400_v1_)
                d_1573_v91_: _dafny.Seq
                d_1573_v91_ = _dafny.SeqWithoutIsStrInference([d_1548_v74_])
                d_1574_v92_: _dafny.Array
                nw270_ = _dafny.Array(None, 11)
                nw270_[int(0)] = d_1548_v74_
                nw270_[int(1)] = d_1548_v74_
                nw270_[int(2)] = d_1548_v74_
                nw270_[int(3)] = d_1548_v74_
                nw270_[int(4)] = d_1548_v74_
                nw270_[int(5)] = (d_1573_v91_)[default__.safeIndex(d_1400_v1_, len(d_1573_v91_))]
                nw270_[int(6)] = d_1548_v74_
                nw270_[int(7)] = d_1548_v74_
                nw270_[int(8)] = d_1548_v74_
                nw270_[int(9)] = (d_1548_v74_ if (d_1407_v8_).f19 else d_1548_v74_)
                nw270_[int(10)] = d_1548_v74_
                d_1574_v92_ = nw270_
                d_1575_v93_: D9
                d_1575_v93_ = D9_DC21((d_1573_v91_)[default__.safeIndex((d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))], len(d_1573_v91_))])
                index246_ = default__.safeIndex(499, (d_1574_v92_).length(0))
                (d_1574_v92_)[index246_] = (d_1575_v93_).cf39
                index247_ = default__.safeIndex(499, (d_1574_v92_).length(0))
                (d_1574_v92_)[index247_] = d_1548_v74_
                (globalState).f8 = ((d_1407_v8_).f20) not in (d_1399_v0_)
                index248_ = default__.safeIndex(793, (d_1548_v74_).length(0))
                (d_1548_v74_)[index248_] = (0) - ((0) - ((d_1548_v74_)[default__.safeIndex(793, (d_1548_v74_).length(0))]))
        d_1576_v94_: _dafny.Map
        d_1576_v94_ = _dafny.Map({_dafny.CodePoint('j'): (d_1402_v3_) | (d_1402_v3_)})
        d_1577_v95_: str
        d_1577_v95_ = _dafny.CodePoint('o')
        d_1576_v94_ = (d_1576_v94_).set((d_1577_v95_ if (d_1407_v8_).f20 else d_1577_v95_), d_1402_v3_)
        d_1578_i22_: int
        d_1578_i22_ = 0
        with _dafny.label("7"):
            while (d_1407_v8_).f19:
                with _dafny.c_label("7"):
                    if (d_1578_i22_) >= (100):
                        raise _dafny.Break("7")
                    d_1578_i22_ = (d_1578_i22_) + (1)
                    d_1579_v96_: _dafny.Seq
                    d_1579_v96_ = _dafny.SeqWithoutIsStrInference([d_1400_v1_, d_1400_v1_, d_1400_v1_])
                    d_1403_v4_ = (d_1403_v4_).set((d_1400_v1_ if p0 else (0) - (d_1400_v1_)), len(d_1579_v96_))
                    d_1580_v97_: _dafny.Array
                    nw271_ = _dafny.Array(None, 6)
                    nw271_[int(0)] = False
                    nw271_[int(1)] = (d_1407_v8_).f19
                    nw271_[int(2)] = p0
                    nw271_[int(3)] = not(not((d_1407_v8_).f19))
                    nw271_[int(4)] = (d_1407_v8_).f20
                    nw271_[int(5)] = (d_1407_v8_).f19
                    d_1580_v97_ = nw271_
                    d_1581_v98_: _dafny.Seq
                    d_1581_v98_ = _dafny.SeqWithoutIsStrInference([d_1580_v97_])
                    d_1582_v99_: _dafny.Seq
                    d_1582_v99_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([d_1580_v97_, d_1580_v97_]), _dafny.SeqWithoutIsStrInference([d_1580_v97_, d_1580_v97_, d_1580_v97_])])
                    d_1583_v100_: _dafny.Seq
                    d_1583_v100_ = _dafny.SeqWithoutIsStrInference([d_1581_v98_, (d_1582_v99_)[default__.safeIndex(d_1400_v1_, len(d_1582_v99_))]])
                    d_1581_v98_ = ((d_1583_v100_)[default__.safeIndex(d_1400_v1_, len(d_1583_v100_))]) + (d_1581_v98_)
                    d_1584_v101_: _dafny.Map
                    d_1584_v101_ = _dafny.Map({d_1400_v1_: p0})
                    d_1584_v101_ = d_1584_v101_
                    d_1585_v102_: _dafny.Array
                    nw272_ = _dafny.Array(int(0), 28)
                    d_1585_v102_ = nw272_
                    index249_ = default__.safeIndex(193, (d_1585_v102_).length(0))
                    (d_1585_v102_)[index249_] = (d_1400_v1_) * (15)
                    d_1586_v103_: _dafny.MultiSet
                    d_1586_v103_ = _dafny.MultiSet([_dafny.MultiSet([765, (0) - (d_1400_v1_), len(d_1581_v98_), d_1400_v1_, d_1400_v1_])])
                    index250_ = default__.safeIndex(193, (d_1585_v102_).length(0))
                    (d_1585_v102_)[index250_] = ((0) - (d_1400_v1_) if not((self).fm42(d_1400_v1_, (d_1586_v103_).cardinality, globalState)) else 878)
                    pass
            pass
        d_1587_v104_: _dafny.Seq
        d_1587_v104_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yxm"))
        r0 = (0) - ((d_1400_v1_ if not (p0) or (p0) else (len(d_1587_v104_)) + (d_1400_v1_)))
        r1 = ((d_1401_v2_) | (d_1401_v2_)) - (d_1401_v2_)
        return r0, r1

    def m11(self, globalState):
        r0: _dafny.MultiSet = _dafny.MultiSet({})
        d_1588_v0_: int
        d_1588_v0_ = -315
        (globalState).f9 = default__.safeDivisionInt(d_1588_v0_, d_1588_v0_)
        d_1589_v1_: _dafny.Seq
        d_1589_v1_ = _dafny.SeqWithoutIsStrInference([len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('b') for d_1590_i0_ in range(default__.abs(59))])), d_1588_v0_])
        d_1591_v2_: D5
        d_1591_v2_ = D5_DC13(d_1589_v1_)
        d_1592_v3_: bool
        d_1592_v3_ = False
        d_1593_v4_: _dafny.Seq
        d_1593_v4_ = _dafny.SeqWithoutIsStrInference([d_1592_v3_])
        d_1594_v5_: str
        d_1594_v5_ = _dafny.CodePoint('r')
        d_1595_v6_: _dafny.Map
        d_1595_v6_ = _dafny.Map({d_1594_v5_: d_1588_v0_})
        d_1596_v7_: _dafny.Map
        d_1596_v7_ = _dafny.Map({not(True): len(d_1595_v6_)})
        d_1597_v8_: _dafny.Set
        d_1597_v8_ = _dafny.Set({d_1592_v3_, d_1592_v3_})
        d_1598_v9_: _dafny.Map
        d_1598_v9_ = _dafny.Map({916: (self).fm42(len(d_1593_v4_), len(d_1597_v8_), globalState)})
        d_1599_v10_: _dafny.Array
        nw273_ = _dafny.Array(None, 22)
        nw273_[int(0)] = d_1588_v0_
        nw273_[int(1)] = d_1588_v0_
        nw273_[int(2)] = d_1588_v0_
        nw273_[int(3)] = (default__.fm45(d_1588_v0_, (d_1591_v2_).cf25, globalState)).cardinality
        nw273_[int(4)] = len(d_1593_v4_)
        nw273_[int(5)] = d_1588_v0_
        nw273_[int(6)] = d_1588_v0_
        nw273_[int(7)] = d_1588_v0_
        nw273_[int(8)] = -209
        nw273_[int(9)] = ((d_1596_v7_)[d_1592_v3_] if (d_1592_v3_) in (d_1596_v7_) else d_1588_v0_)
        nw273_[int(10)] = d_1588_v0_
        nw273_[int(11)] = len(d_1598_v9_)
        nw273_[int(12)] = d_1588_v0_
        nw273_[int(13)] = d_1588_v0_
        nw273_[int(14)] = 118
        nw273_[int(15)] = d_1588_v0_
        nw273_[int(16)] = d_1588_v0_
        nw273_[int(17)] = len(d_1595_v6_)
        nw273_[int(18)] = len(d_1596_v7_)
        nw273_[int(19)] = d_1588_v0_
        nw273_[int(20)] = -547
        nw273_[int(21)] = d_1588_v0_
        d_1599_v10_ = nw273_
        pat_let_tv28_ = d_1599_v10_
        pat_let_tv29_ = d_1599_v10_
        def iife106_(_pat_let24_0):
            def iife107_(d_1600_dt__update__tmp_h0_):
                def iife108_(_pat_let25_0):
                    def iife109_(d_1601_dt__update_hcf39_h0_):
                        return D9_DC21(d_1601_dt__update_hcf39_h0_)
                    return iife109_(_pat_let25_0)
                return iife108_(pat_let_tv28_)
            return iife107_(_pat_let24_0)
        def iife105_(_pat_let23_0):
            def iife110_(d_1602_dt__update__tmp_h1_):
                def iife111_(_pat_let26_0):
                    def iife112_(d_1603_dt__update_hcf39_h1_):
                        return D9_DC21(d_1603_dt__update_hcf39_h1_)
                    return iife112_(_pat_let26_0)
                return iife111_(pat_let_tv29_)
            return iife110_(_pat_let23_0)
        source19_ = iife105_(iife106_(D9_DC21(d_1599_v10_)))
        if source19_.is_DC22:
            d_1604___mcc_h0_ = source19_.cf40
            d_1605___mcc_h1_ = source19_.cf41
            d_1606_cf41_ = d_1605___mcc_h1_
            d_1607_cf40_ = d_1604___mcc_h0_
            d_1608_v11_: _dafny.Array
            def lambda97_(d_1609_i1_):
                return True

            init53_ = lambda97_
            nw274_ = _dafny.Array(None, 22)
            for i0_53_ in range(nw274_.length(0)):
                nw274_[i0_53_] = init53_(i0_53_)
            d_1608_v11_ = nw274_
            d_1610_v12_: _dafny.Map
            d_1610_v12_ = _dafny.Map({(d_1589_v1_) + (d_1589_v1_): d_1608_v11_})
            rhs277_ = (302) >= (d_1588_v0_)
            rhs278_ = d_1610_v12_
            lhs200_ = globalState
            lhs200_.f8 = rhs277_
            d_1610_v12_ = rhs278_
            (globalState).f8 = not((self).fm42(((default__.fm46(globalState)).cardinality) * (18), d_1588_v0_, globalState))
            d_1611_v13_: _dafny.Seq
            d_1611_v13_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "orpv"))
            d_1607_cf40_ = ((len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "eadjqkvyy")))) * ((self).fm41(d_1611_v13_, d_1588_v0_, d_1607_cf40_, (d_1593_v4_)[default__.safeIndex(d_1607_cf40_, len(d_1593_v4_))], globalState))) - (default__.safeDivisionInt(d_1607_cf40_, d_1607_cf40_))
            d_1612_v14_: _dafny.MultiSet
            d_1612_v14_ = _dafny.MultiSet([False])
            d_1613_v15_: _dafny.Array
            nw275_ = _dafny.Array(None, 4)
            nw275_[int(0)] = d_1612_v14_
            nw275_[int(1)] = d_1612_v14_
            nw275_[int(2)] = d_1612_v14_
            nw275_[int(3)] = _dafny.MultiSet([not(d_1592_v3_), d_1592_v3_])
            d_1613_v15_ = nw275_
            d_1614_v16_: _dafny.Map
            d_1614_v16_ = _dafny.Map({d_1613_v15_: d_1592_v3_})
            d_1614_v16_ = (d_1614_v16_).set(d_1613_v15_, (d_1588_v0_) == ((0) - (d_1607_cf40_)))
        elif True:
            d_1615___mcc_h2_ = source19_.cf39
            d_1616_cf39_ = d_1615___mcc_h2_
            d_1617_v17_: _dafny.Array
            nw276_ = _dafny.Array(False, 22)
            d_1617_v17_ = nw276_
            index251_ = default__.safeIndex(585, (d_1617_v17_).length(0))
            (d_1617_v17_)[index251_] = d_1592_v3_
            d_1618_v18_: _dafny.Array
            nw277_ = _dafny.Array(_dafny.CodePoint('D'), 23)
            d_1618_v18_ = nw277_
            index252_ = default__.safeIndex(89, (d_1618_v18_).length(0))
            (d_1618_v18_)[index252_] = d_1594_v5_
            d_1619_v19_: _dafny.Seq
            d_1619_v19_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "u"))
            d_1620_v20_: _dafny.Array
            nw278_ = _dafny.Array(None, 14)
            nw278_[int(0)] = d_1589_v1_
            nw278_[int(1)] = d_1589_v1_
            nw278_[int(2)] = d_1589_v1_
            nw278_[int(3)] = d_1589_v1_
            nw278_[int(4)] = d_1589_v1_
            nw278_[int(5)] = d_1589_v1_
            nw278_[int(6)] = d_1589_v1_
            nw278_[int(7)] = _dafny.SeqWithoutIsStrInference([len(d_1589_v1_) for d_1621_i2_ in range(default__.abs(-140))])
            nw278_[int(8)] = d_1589_v1_
            nw278_[int(9)] = d_1589_v1_
            nw278_[int(10)] = d_1589_v1_
            nw278_[int(11)] = d_1589_v1_
            nw278_[int(12)] = d_1589_v1_
            nw278_[int(13)] = d_1589_v1_
            d_1620_v20_ = nw278_
            d_1622_v21_: _dafny.Map
            d_1622_v21_ = _dafny.Map({d_1620_v20_: len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erktiyai")))})
            index253_ = default__.safeIndex(585, (d_1617_v17_).length(0))
            index254_ = default__.safeIndex(89, (d_1618_v18_).length(0))
            rhs279_ = (d_1594_v5_) not in (d_1619_v19_)
            rhs280_ = 625
            rhs281_ = default__.fm1(d_1592_v3_, d_1594_v5_, globalState)
            rhs282_ = (((d_1622_v21_)[d_1620_v20_] if (d_1620_v20_) in (d_1622_v21_) else d_1588_v0_)) - (461)
            rhs283_ = d_1594_v5_
            lhs201_ = globalState
            lhs202_ = d_1617_v17_
            lhs203_ = default__.safeIndex(585, (d_1617_v17_).length(0))
            lhs204_ = globalState
            lhs205_ = d_1618_v18_
            lhs206_ = default__.safeIndex(89, (d_1618_v18_).length(0))
            d_1592_v3_ = rhs279_
            lhs201_.f9 = rhs280_
            lhs202_[lhs203_] = rhs281_
            lhs204_.f9 = rhs282_
            lhs205_[lhs206_] = rhs283_
            d_1623_v22_: _dafny.Set
            d_1623_v22_ = _dafny.Set({d_1588_v0_})
            d_1624_v23_: D0
            d_1624_v23_ = D0_DC0(d_1588_v0_)
            d_1625_v24_: D2
            d_1625_v24_ = D2_DC5(d_1619_v19_)
            d_1626_v25_: C1
            nw279_ = C1()
            nw279_.ctor__((d_1623_v22_).ispropersubset(d_1623_v22_), (d_1617_v17_)[default__.safeIndex(585, (d_1617_v17_).length(0))], d_1624_v23_, ((d_1625_v24_).cf10) <= (d_1619_v19_))
            d_1626_v25_ = nw279_
            (globalState).f9 = (d_1588_v0_) + (d_1588_v0_)
            (globalState).f8 = d_1592_v3_
        d_1627_v26_: T0
        nw280_ = C3()
        nw280_.ctor__((d_1588_v0_) != (len(d_1596_v7_)), (d_1592_v3_) == (d_1592_v3_), (self).f11)
        d_1627_v26_ = nw280_
        d_1627_v26_ = d_1627_v26_
        default__.m0(globalState)
        d_1599_v10_ = (d_1599_v10_ if d_1592_v3_ else d_1599_v10_)
        hi12_ = default__.fm0(d_1592_v3_, globalState)
        for d_1628_i3_ in range(len((_dafny.Map({d_1588_v0_: d_1592_v3_})) | (d_1598_v9_)), hi12_):
            d_1629_v27_: _dafny.Array
            def lambda98_(d_1630_v3_, d_1631_i3_):
                def lambda99_(d_1632_i4_):
                    return D2_DC7(d_1630_v3_, (0) - (d_1631_i3_), d_1630_v3_, _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ygpysjnm")))

                return lambda99_

            init54_ = lambda98_(d_1592_v3_, d_1628_i3_)
            nw281_ = _dafny.Array(None, 25)
            for i0_54_ in range(nw281_.length(0)):
                nw281_[i0_54_] = init54_(i0_54_)
            d_1629_v27_ = nw281_
            d_1633_v28_: D11
            d_1633_v28_ = D11_DC25(d_1629_v27_)
            source20_ = d_1633_v28_
            if source20_.is_DC26:
                d_1634___mcc_h3_ = source20_.cf47
                d_1635_cf47_ = d_1634___mcc_h3_
                d_1636_v29_: _dafny.Map
                d_1636_v29_ = _dafny.Map({d_1592_v3_: _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([d_1635_cf47_ for d_1637_i5_ in range(default__.abs(364))]))})
                d_1638_v30_: _dafny.MultiSet
                d_1638_v30_ = _dafny.MultiSet([d_1635_cf47_])
                d_1592_v3_ = ((d_1636_v29_).set(d_1592_v3_, d_1638_v30_)) == (d_1636_v29_)
                d_1639_v31_: _dafny.Array
                nw282_ = _dafny.Array(_dafny.Set({}), 13)
                d_1639_v31_ = nw282_
                index255_ = default__.safeIndex(81, (d_1639_v31_).length(0))
                (d_1639_v31_)[index255_] = _dafny.Set({d_1588_v0_, (0) - (d_1588_v0_)})
                index256_ = default__.safeIndex(81, (d_1639_v31_).length(0))
                def iife113_():
                    coll59_ = _dafny.Set()
                    compr_59_: int
                    for compr_59_ in _dafny.IntegerRange(428, 619):
                        d_1640_v32_: int = compr_59_
                        if ((428) <= (d_1640_v32_)) and ((d_1640_v32_) < (619)):
                            coll59_ = coll59_.union(_dafny.Set([(d_1640_v32_) * ((D9_DC22(d_1628_i3_, _dafny.Map({d_1628_i3_: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kamnjg"))}))).cf40)]))
                    return _dafny.Set(coll59_)
                (d_1639_v31_)[index256_] = iife113_()
                
                d_1641_v35_: _dafny.Map
                d_1641_v35_ = _dafny.Map({len(_dafny.SeqWithoutIsStrInference([len((d_1639_v31_)[default__.safeIndex(81, (d_1639_v31_).length(0))]), 834, d_1588_v0_])): ((d_1596_v7_)[d_1592_v3_] if (d_1592_v3_) in (d_1596_v7_) else d_1635_cf47_)})
                d_1642_v36_: _dafny.Seq
                d_1642_v36_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vusjkfsmm"))
                d_1643_v37_: D2
                d_1643_v37_ = D2_DC7(d_1592_v3_, d_1628_i3_, d_1592_v3_, d_1642_v36_)
                d_1644_v39_: _dafny.Array
                nw283_ = _dafny.Array(None, 20)
                nw283_[int(0)] = d_1592_v3_
                nw283_[int(1)] = d_1592_v3_
                nw283_[int(2)] = d_1592_v3_
                nw283_[int(3)] = d_1592_v3_
                nw283_[int(4)] = d_1592_v3_
                nw283_[int(5)] = False
                nw283_[int(6)] = d_1592_v3_
                nw283_[int(7)] = d_1592_v3_
                nw283_[int(8)] = d_1592_v3_
                nw283_[int(9)] = d_1592_v3_
                nw283_[int(10)] = d_1592_v3_
                nw283_[int(11)] = d_1592_v3_
                nw283_[int(12)] = d_1592_v3_
                nw283_[int(13)] = d_1592_v3_
                nw283_[int(14)] = d_1592_v3_
                nw283_[int(15)] = d_1592_v3_
                nw283_[int(16)] = d_1592_v3_
                nw283_[int(17)] = d_1592_v3_
                nw283_[int(18)] = d_1592_v3_
                nw283_[int(19)] = d_1592_v3_
                d_1644_v39_ = nw283_
                d_1645_v40_: _dafny.Map
                d_1645_v40_ = _dafny.Map({d_1644_v39_: d_1644_v39_})
                d_1646_v41_: _dafny.Array
                nw284_ = _dafny.Array(None, 15)
                def iife114_():
                    coll60_ = _dafny.Map()
                    compr_60_: int
                    for compr_60_ in (d_1598_v9_).keys.Elements:
                        d_1647_v33_: int = compr_60_
                        if (d_1647_v33_) in (d_1598_v9_):
                            coll60_[(d_1647_v33_) + (d_1628_i3_)] = d_1635_cf47_
                    return _dafny.Map(coll60_)
                nw284_[int(0)] = iife114_()
                
                def iife115_():
                    coll61_ = _dafny.Map()
                    compr_61_: int
                    for compr_61_ in (d_1598_v9_).keys.Elements:
                        d_1648_v34_: int = compr_61_
                        if (d_1648_v34_) in (d_1598_v9_):
                            coll61_[default__.safeModuloInt(d_1648_v34_, d_1635_cf47_)] = d_1628_i3_
                    return _dafny.Map(coll61_)
                nw284_[int(1)] = _dafny.Map({(((d_1636_v29_)[d_1592_v3_] if (d_1592_v3_) in (d_1636_v29_) else d_1638_v30_)).cardinality: len(iife115_()
                )})
                nw284_[int(2)] = d_1641_v35_
                nw284_[int(3)] = _dafny.Map({len(d_1642_v36_): len((d_1643_v37_).cf14)})
                nw284_[int(4)] = d_1641_v35_
                def iife116_():
                    coll62_ = _dafny.Map()
                    compr_62_: int
                    for compr_62_ in _dafny.IntegerRange(979, 110):
                        d_1649_v38_: int = compr_62_
                        if ((979) <= (d_1649_v38_)) and ((d_1649_v38_) < (110)):
                            coll62_[default__.safeModuloInt(d_1649_v38_, -347)] = d_1628_i3_
                    return _dafny.Map(coll62_)
                nw284_[int(5)] = iife116_()
                
                nw284_[int(6)] = _dafny.Map({558: d_1588_v0_})
                nw284_[int(7)] = (d_1641_v35_)
                nw284_[int(8)] = d_1641_v35_
                nw284_[int(9)] = d_1641_v35_
                nw284_[int(10)] = d_1641_v35_
                nw284_[int(11)] = _dafny.Map({d_1635_cf47_: d_1635_cf47_})
                nw284_[int(12)] = _dafny.Map({len(d_1645_v40_): d_1628_i3_})
                nw284_[int(13)] = d_1641_v35_
                nw284_[int(14)] = d_1641_v35_
                d_1646_v41_ = nw284_
                d_1650_v42_: D10
                d_1650_v42_ = D10_DC23(d_1646_v41_)
                d_1651_v43_: _dafny.Seq
                d_1651_v43_ = _dafny.SeqWithoutIsStrInference([d_1638_v30_])
                pat_let_tv30_ = d_1646_v41_
                def iife117_(_pat_let27_0):
                    def iife118_(d_1652_dt__update__tmp_h2_):
                        def iife119_(_pat_let28_0):
                            def iife120_(d_1653_dt__update_hcf42_h0_):
                                return D10_DC23(d_1653_dt__update_hcf42_h0_)
                            return iife120_(_pat_let28_0)
                        return iife119_(pat_let_tv30_)
                    return iife118_(_pat_let27_0)
                rhs284_ = ((d_1598_v9_)[d_1588_v0_] if (d_1588_v0_) in (d_1598_v9_) else True)
                rhs285_ = (d_1642_v36_)[default__.safeIndex(d_1628_i3_, len(d_1642_v36_))]
                rhs286_ = iife117_(d_1650_v42_)
                rhs287_ = (_dafny.Set({d_1638_v30_, (d_1651_v43_)[default__.safeIndex((0) - ((0) - (d_1628_i3_)), len(d_1651_v43_))]})).ispropersubset(_dafny.Set({((_dafny.MultiSet([len(d_1589_v1_)])).set(d_1588_v0_, default__.abs(d_1628_i3_))).set((0) - (len(d_1589_v1_)), default__.abs(len(d_1642_v36_)))}))
                lhs207_ = globalState
                lhs207_.f8 = rhs284_
                d_1594_v5_ = rhs285_
                d_1650_v42_ = rhs286_
                d_1592_v3_ = rhs287_
                d_1654_v44_: _dafny.Map
                d_1654_v44_ = _dafny.Map({d_1635_cf47_: d_1599_v10_})
                d_1655_v45_: _dafny.Seq
                d_1655_v45_ = _dafny.SeqWithoutIsStrInference([d_1654_v44_])
                d_1656_v46_: _dafny.Map
                d_1656_v46_ = _dafny.Map({d_1599_v10_: (d_1655_v45_)[default__.safeIndex(d_1635_cf47_, len(d_1655_v45_))]})
                d_1654_v44_ = (((d_1656_v46_)[d_1599_v10_] if (d_1599_v10_) in (d_1656_v46_) else d_1654_v44_)).set(d_1588_v0_, d_1599_v10_)
            elif source20_.is_DC25:
                d_1657___mcc_h4_ = source20_.cf46
                d_1658_cf46_ = d_1657___mcc_h4_
                index257_ = default__.safeIndex(771, (d_1599_v10_).length(0))
                (d_1599_v10_)[index257_] = (0) - (d_1628_i3_)
                index258_ = default__.safeIndex(771, (d_1599_v10_).length(0))
                (d_1599_v10_)[index258_] = d_1628_i3_
                d_1659_v47_: _dafny.Seq
                d_1659_v47_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))
                d_1660_v48_: _dafny.Seq
                d_1660_v48_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "wj")), (d_1659_v47_) + (d_1659_v47_)])
                d_1660_v48_ = d_1660_v48_
                d_1661_v49_: _dafny.MultiSet
                d_1661_v49_ = _dafny.MultiSet([d_1592_v3_])
                (globalState).f9 = (((d_1661_v49_)[d_1592_v3_] if (d_1592_v3_) in (d_1661_v49_) else d_1628_i3_) if d_1592_v3_ else d_1588_v0_)
                d_1662_v50_: _dafny.Array
                def lambda100_(d_1663_v47_):
                    def lambda101_(d_1664_i6_):
                        return default__.safeDivisionInt(d_1664_i6_, len(d_1663_v47_))

                    return lambda101_

                init55_ = lambda100_(d_1659_v47_)
                nw285_ = _dafny.Array(None, 19)
                for i0_55_ in range(nw285_.length(0)):
                    nw285_[i0_55_] = init55_(i0_55_)
                d_1662_v50_ = nw285_
                d_1665_v51_: _dafny.Map
                d_1665_v51_ = _dafny.Map({d_1662_v50_: d_1592_v3_})
                d_1666_v52_: _dafny.Map
                d_1666_v52_ = _dafny.Map({d_1665_v51_: d_1592_v3_})
                d_1666_v52_ = (d_1666_v52_).set((_dafny.Map({d_1662_v50_: d_1592_v3_})) | (_dafny.Map({d_1662_v50_: d_1592_v3_})), False)
            elif True:
                d_1667___mcc_h5_ = source20_.cf48
                d_1668_cf48_ = d_1667___mcc_h5_
                d_1669_v53_: _dafny.MultiSet
                d_1669_v53_ = _dafny.MultiSet([533, (d_1589_v1_)[default__.safeIndex((0) - (d_1588_v0_), len(d_1589_v1_))]])
                d_1670_v54_: _dafny.MultiSet
                d_1670_v54_ = d_1669_v53_
                d_1671_v55_: _dafny.Seq
                d_1671_v55_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "offr"))
                d_1672_v56_: _dafny.MultiSet
                d_1672_v56_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_1593_v4_)).cardinality for d_1673_i7_ in range(default__.abs(18))])])
                d_1674_v57_: _dafny.Seq
                d_1674_v57_ = _dafny.SeqWithoutIsStrInference([d_1669_v53_, d_1669_v53_])
                d_1675_v58_: _dafny.Seq
                d_1675_v58_ = _dafny.SeqWithoutIsStrInference([d_1669_v53_, d_1669_v53_, (d_1674_v57_)[default__.safeIndex(len(_dafny.SeqWithoutIsStrInference([d_1628_i3_, 698])), len(d_1674_v57_))], d_1669_v53_, _dafny.MultiSet([d_1628_i3_, d_1588_v0_])])
                d_1676_v59_: _dafny.Array
                nw286_ = _dafny.Array(None, 20)
                nw286_[int(0)] = (d_1669_v53_).intersection(d_1669_v53_)
                nw286_[int(1)] = d_1669_v53_
                nw286_[int(2)] = d_1669_v53_
                nw286_[int(3)] = d_1669_v53_
                nw286_[int(4)] = d_1669_v53_
                nw286_[int(5)] = d_1669_v53_
                nw286_[int(6)] = _dafny.MultiSet([256])
                nw286_[int(7)] = (d_1669_v53_) - (d_1669_v53_)
                nw286_[int(8)] = d_1669_v53_
                nw286_[int(9)] = (d_1670_v54_)
                nw286_[int(10)] = (d_1669_v53_).intersection(d_1669_v53_)
                nw286_[int(11)] = (d_1669_v53_) - (_dafny.MultiSet([len(d_1671_v55_), d_1628_i3_, d_1628_i3_, len(d_1596_v7_), (0) - (d_1628_i3_)]))
                nw286_[int(12)] = default__.fm56(d_1592_v3_, d_1628_i3_, globalState)
                nw286_[int(13)] = d_1669_v53_
                nw286_[int(14)] = (d_1669_v53_).set(default__.fm0(d_1592_v3_, globalState), default__.abs(default__.fm0(not(d_1592_v3_), globalState)))
                nw286_[int(15)] = (_dafny.MultiSet([d_1588_v0_, (d_1672_v56_).cardinality, d_1628_i3_]) if d_1592_v3_ else d_1669_v53_)
                nw286_[int(16)] = (d_1669_v53_) | ((d_1674_v57_)[default__.safeIndex(d_1588_v0_, len(d_1674_v57_))])
                nw286_[int(17)] = _dafny.MultiSet(default__.fm54(d_1592_v3_, d_1592_v3_, not(d_1592_v3_), globalState))
                nw286_[int(18)] = (d_1669_v53_).intersection(d_1669_v53_)
                nw286_[int(19)] = d_1669_v53_
                d_1676_v59_ = nw286_
                index259_ = default__.safeIndex(847, (d_1676_v59_).length(0))
                (d_1676_v59_)[index259_] = d_1669_v53_
                index260_ = default__.safeIndex(847, (d_1676_v59_).length(0))
                (d_1676_v59_)[index260_] = d_1669_v53_
                d_1677_v60_: _dafny.Array
                nw287_ = _dafny.Array(False, 9)
                d_1677_v60_ = nw287_
                d_1678_v61_: _dafny.Map
                d_1678_v61_ = _dafny.Map({d_1599_v10_: d_1592_v3_})
                index261_ = default__.safeIndex(136, (d_1677_v60_).length(0))
                (d_1677_v60_)[index261_] = ((d_1678_v61_)[d_1599_v10_] if (d_1599_v10_) in (d_1678_v61_) else d_1592_v3_)
                index262_ = default__.safeIndex(136, (d_1677_v60_).length(0))
                (d_1677_v60_)[index262_] = (d_1597_v8_).ispropersubset(d_1597_v8_)
                d_1592_v3_ = d_1592_v3_
                d_1679_v62_: _dafny.Map
                d_1679_v62_ = _dafny.Map({d_1628_i3_: d_1628_i3_})
                index263_ = default__.safeIndex(805, (d_1599_v10_).length(0))
                (d_1599_v10_)[index263_] = ((d_1679_v62_)[d_1628_i3_] if (d_1628_i3_) in (d_1679_v62_) else 407)
                d_1680_v63_: _dafny.Seq
                d_1680_v63_ = _dafny.SeqWithoutIsStrInference([d_1593_v4_])
                d_1681_v64_: _dafny.MultiSet
                d_1681_v64_ = _dafny.MultiSet([(d_1677_v60_)[default__.safeIndex(136, (d_1677_v60_).length(0))], (len(default__.fm2((0) - (d_1628_i3_), ((d_1676_v59_)[default__.safeIndex(847, (d_1676_v59_).length(0))]).set(d_1588_v0_, default__.abs(d_1628_i3_)), d_1679_v62_, globalState))) < (d_1628_i3_), (d_1677_v60_)[default__.safeIndex(136, (d_1677_v60_).length(0))]])
                d_1682_v65_: _dafny.Set
                d_1682_v65_ = _dafny.Set({len(d_1671_v55_)})
                index264_ = default__.safeIndex(805, (d_1599_v10_).length(0))
                rhs288_ = ((d_1593_v4_) + (default__.fm44(d_1680_v63_, _dafny.CodePoint('e'), globalState))) + (d_1593_v4_)
                rhs289_ = ((d_1681_v64_)[(d_1682_v65_).issubset(d_1682_v65_)] if ((d_1682_v65_).issubset(d_1682_v65_)) in (d_1681_v64_) else default__.safeDivisionInt(d_1588_v0_, (0) - (len(d_1671_v55_))))
                lhs208_ = d_1599_v10_
                lhs209_ = default__.safeIndex(805, (d_1599_v10_).length(0))
                d_1593_v4_ = rhs288_
                lhs208_[lhs209_] = rhs289_
            d_1592_v3_ = d_1592_v3_
            d_1683_v66_: _dafny.MultiSet
            d_1683_v66_ = _dafny.MultiSet([d_1628_i3_])
            d_1684_v67_: _dafny.MultiSet
            d_1684_v67_ = _dafny.MultiSet([d_1683_v66_, d_1683_v66_, (d_1683_v66_).set(d_1588_v0_, default__.abs(d_1628_i3_))])
            rhs290_ = d_1592_v3_
            rhs291_ = d_1684_v67_
            lhs210_ = globalState
            lhs210_.f8 = rhs290_
            d_1684_v67_ = rhs291_
            d_1685_v68_: _dafny.Set
            d_1685_v68_ = _dafny.Set({_dafny.CodePoint('r'), d_1594_v5_})
            d_1686_v69_: _dafny.Set
            d_1686_v69_ = d_1685_v68_
            d_1596_v7_ = (d_1596_v7_).set(d_1592_v3_, (0) - ((len(_dafny.Set({d_1686_v69_}))) - (504)))
        d_1687_v70_: _dafny.MultiSet
        d_1687_v70_ = _dafny.MultiSet([d_1592_v3_, not(d_1592_v3_), d_1592_v3_])
        r0 = d_1687_v70_
        return r0


class C5(T0, T1):
    def  __init__(self):
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C5"
    @property
    def f11(self):
        return self._f11
    def ctor__(self, f11):
        (self)._f11 = f11

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_1688_v0_: int
        d_1688_v0_ = -540
        d_1689_v1_: _dafny.MultiSet
        d_1689_v1_ = _dafny.MultiSet([d_1688_v0_, d_1688_v0_, default__.fm0(p0, globalState)])
        (globalState).f9 = default__.fm0((d_1689_v1_).issubset(d_1689_v1_), globalState)
        (globalState).f8 = p0
        (globalState).f8 = p0
        d_1690_v2_: str
        d_1690_v2_ = _dafny.CodePoint('y')
        d_1691_v3_: D2
        d_1691_v3_ = D2_DC8(d_1688_v0_, d_1690_v2_)
        d_1692_v4_: _dafny.Map
        d_1692_v4_ = _dafny.Map({817: (d_1691_v3_).cf16})
        d_1692_v4_ = (d_1692_v4_).set(d_1688_v0_, d_1690_v2_)
        default__.m0(globalState)
        d_1693_v5_: C0
        nw288_ = C0()
        nw288_.ctor__(p0)
        d_1693_v5_ = nw288_
        r0 = d_1688_v0_
        pat_let_tv31_ = d_1693_v5_
        pat_let_tv32_ = d_1693_v5_
        pat_let_tv33_ = d_1693_v5_
        pat_let_tv34_ = d_1693_v5_
        pat_let_tv35_ = d_1693_v5_
        pat_let_tv36_ = d_1693_v5_
        def lambda102_(source21_):
            if source21_.is_DC6:
                return (_dafny.Set({not((pat_let_tv31_).f17)})) | (_dafny.Set({(pat_let_tv32_).f17, False}))
            elif source21_.is_DC7:
                d_1694___mcc_h0_ = source21_.cf11
                d_1695___mcc_h1_ = source21_.cf12
                d_1696___mcc_h2_ = source21_.cf13
                d_1697___mcc_h3_ = source21_.cf14
                d_1698_cf14_ = d_1697___mcc_h3_
                d_1699_cf13_ = d_1696___mcc_h2_
                d_1700_cf12_ = d_1695___mcc_h1_
                d_1701_cf11_ = d_1694___mcc_h0_
                return _dafny.Set({True, True})
            elif source21_.is_DC8:
                d_1702___mcc_h4_ = source21_.cf15
                d_1703___mcc_h5_ = source21_.cf16
                d_1704_cf16_ = d_1703___mcc_h5_
                d_1705_cf15_ = d_1702___mcc_h4_
                return (_dafny.Set({(pat_let_tv33_).f17})).intersection(_dafny.Set({(pat_let_tv34_).f17, (pat_let_tv35_).f17}))
            elif True:
                d_1706___mcc_h6_ = source21_.cf10
                d_1707_cf10_ = d_1706___mcc_h6_
                return (_dafny.Set({(pat_let_tv36_).f17})).intersection(_dafny.Set({False}))

        r1 = lambda102_(d_1691_v3_)
        return r0, r1

    def m4(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: D2 = D2.default()()
        r3: bool = False
        d_1708_v0_: bool
        d_1708_v0_ = True
        if d_1708_v0_:
            d_1709_v2_: _dafny.Array
            def lambda103_(d_1710_v0_):
                def lambda104_(d_1711_i0_):
                    def iife121_():
                        coll63_ = _dafny.Set()
                        compr_63_: int
                        for compr_63_ in _dafny.IntegerRange(-866, 290):
                            d_1712_v1_: int = compr_63_
                            if ((-866) <= (d_1712_v1_)) and ((d_1712_v1_) < (290)):
                                coll63_ = coll63_.union(_dafny.Set([(d_1712_v1_) - (len(_dafny.Map({False: len(_dafny.Map({d_1710_v0_: d_1710_v0_}))})))]))
                        return _dafny.Set(coll63_)
                    return iife121_()
                    

                return lambda104_

            init56_ = lambda103_(d_1708_v0_)
            nw289_ = _dafny.Array(None, 10)
            for i0_56_ in range(nw289_.length(0)):
                nw289_[i0_56_] = init56_(i0_56_)
            d_1709_v2_ = nw289_
            index265_ = default__.safeIndex(253, (d_1709_v2_).length(0))
            (d_1709_v2_)[index265_] = _dafny.Set({p0, p0})
            d_1713_v3_: _dafny.Set
            d_1713_v3_ = _dafny.Set({p0})
            index266_ = default__.safeIndex(253, (d_1709_v2_).length(0))
            (d_1709_v2_)[index266_] = (d_1713_v3_) - ((d_1713_v3_) | (d_1713_v3_))
            d_1714_v4_: _dafny.Array
            def lambda105_(d_1715_v0_):
                def lambda106_(d_1716_i1_):
                    return d_1715_v0_

                return lambda106_

            init57_ = lambda105_(d_1708_v0_)
            nw290_ = _dafny.Array(None, 4)
            for i0_57_ in range(nw290_.length(0)):
                nw290_[i0_57_] = init57_(i0_57_)
            d_1714_v4_ = nw290_
            index267_ = default__.safeIndex(601, (d_1714_v4_).length(0))
            (d_1714_v4_)[index267_] = True
            index268_ = default__.safeIndex(601, (d_1714_v4_).length(0))
            (d_1714_v4_)[index268_] = (p0) > (p0)
            r1 = p0
            d_1717_v5_: str
            d_1717_v5_ = _dafny.CodePoint('h')
            r0 = d_1717_v5_
            d_1718_v6_: _dafny.Seq
            d_1718_v6_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jh"))
            (globalState).f9 = len((d_1718_v6_).set(default__.safeIndex(p0, len(d_1718_v6_)), d_1717_v5_))
        elif True:
            default__.m0(globalState)
            d_1719_v7_: _dafny.Seq
            d_1719_v7_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cyw"))
            if (d_1708_v0_ if (d_1719_v7_) != (d_1719_v7_) else d_1708_v0_):
                d_1720_v8_: _dafny.Map
                d_1720_v8_ = _dafny.Map({d_1708_v0_: p0})
                (globalState).f8 = ((p0) + (p0)) != (len((d_1720_v8_) | (d_1720_v8_)))
                r1 = -719
                (globalState).f8 = d_1708_v0_
                d_1721_v9_: D2
                d_1721_v9_ = D2_DC7(d_1708_v0_, p0, d_1708_v0_, d_1719_v7_)
                d_1722_v10_: _dafny.MultiSet
                d_1722_v10_ = _dafny.MultiSet([d_1708_v0_])
                rhs292_ = not ((d_1721_v9_).cf13) or (d_1708_v0_)
                rhs293_ = d_1708_v0_
                rhs294_ = ((0) - (p0)) * ((0) - ((p0) - ((0) - ((d_1722_v10_).cardinality))))
                rhs295_ = p0
                lhs211_ = globalState
                lhs212_ = globalState
                lhs213_ = globalState
                lhs211_.f8 = rhs292_
                lhs212_.f8 = rhs293_
                lhs213_.f9 = rhs294_
                r1 = rhs295_
                d_1708_v0_ = (_dafny.CodePoint('u')) in (d_1719_v7_)
            elif True:
                (globalState).f9 = p0
                r1 = p0
                d_1723_v11_: _dafny.Array
                nw291_ = _dafny.Array(_dafny.Array(None, 0), 9)
                d_1723_v11_ = nw291_
                index269_ = default__.safeIndex(974, (d_1723_v11_).length(0))
                (d_1723_v11_)[index269_] = p1
                d_1724_v12_: _dafny.Array
                nw292_ = _dafny.Array(False, 25)
                d_1724_v12_ = nw292_
                index270_ = default__.safeIndex(756, (d_1724_v12_).length(0))
                (d_1724_v12_)[index270_] = (default__.fm0(d_1708_v0_, globalState)) != ((_dafny.MultiSet([p0, p0, p0, p0, p0])).cardinality)
                index271_ = default__.safeIndex(974, (d_1723_v11_).length(0))
                index272_ = default__.safeIndex(756, (d_1724_v12_).length(0))
                rhs296_ = p1
                rhs297_ = d_1708_v0_
                lhs214_ = d_1723_v11_
                lhs215_ = default__.safeIndex(974, (d_1723_v11_).length(0))
                lhs216_ = d_1724_v12_
                lhs217_ = default__.safeIndex(756, (d_1724_v12_).length(0))
                lhs214_[lhs215_] = rhs296_
                lhs216_[lhs217_] = rhs297_
                d_1725_v13_: _dafny.Seq
                d_1725_v13_ = _dafny.SeqWithoutIsStrInference([p0, p0])
                (globalState).f9 = (d_1725_v13_)[default__.safeIndex(p0, len(d_1725_v13_))]
                d_1726_v14_: _dafny.Map
                d_1726_v14_ = _dafny.Map({d_1708_v0_: p0})
                d_1726_v14_ = (d_1726_v14_).set((d_1708_v0_) or ((d_1724_v12_)[default__.safeIndex(756, (d_1724_v12_).length(0))]), p0)
            (globalState).f9 = default__.safeModuloInt(p0, 879)
            d_1727_v15_: _dafny.MultiSet
            d_1727_v15_ = _dafny.MultiSet([d_1708_v0_, (False if d_1708_v0_ else d_1708_v0_), d_1708_v0_])
            d_1727_v15_ = d_1727_v15_
            d_1708_v0_ = d_1708_v0_
        hi13_ = default__.fm0(d_1708_v0_, globalState)
        for d_1728_i2_ in range(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "xpnbhaeoo"))), hi13_):
            source22_ = D4_DC10(default__.fm16(globalState))
            if source22_.is_DC11:
                d_1729___mcc_h0_ = source22_.cf19
                d_1730___mcc_h1_ = source22_.cf20
                d_1731___mcc_h2_ = source22_.cf21
                d_1732_cf21_ = d_1731___mcc_h2_
                d_1733_cf20_ = d_1730___mcc_h1_
                d_1734_cf19_ = d_1729___mcc_h0_
                (globalState).f9 = p0
                d_1735_v16_: _dafny.Seq
                d_1735_v16_ = _dafny.SeqWithoutIsStrInference([d_1732_cf21_])
                d_1736_v17_: _dafny.Seq
                d_1736_v17_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fsyimuwr"))
                d_1737_v18_: str
                d_1737_v18_ = _dafny.CodePoint('r')
                d_1738_v19_: D2
                d_1738_v19_ = D2_DC7(d_1708_v0_, d_1728_i2_, (d_1735_v16_)[default__.safeIndex(p0, len(d_1735_v16_))], ((d_1736_v17_).set(default__.safeIndex(114, len(d_1736_v17_)), d_1737_v18_)) + (d_1736_v17_))
                d_1738_v19_ = default__.fm17(globalState)
                index273_ = default__.safeIndex(953, (p1).length(0))
                (p1)[index273_] = d_1734_cf19_
                d_1739_v20_: _dafny.Array
                nw293_ = _dafny.Array(_dafny.CodePoint('D'), 15)
                d_1739_v20_ = nw293_
                d_1740_v21_: _dafny.MultiSet
                d_1740_v21_ = _dafny.MultiSet([d_1739_v20_, d_1739_v20_])
                index274_ = default__.safeIndex(953, (p1).length(0))
                (p1)[index274_] = (0) - ((d_1740_v21_).cardinality)
                (globalState).f9 = p0
            elif source22_.is_DC12:
                d_1741___mcc_h3_ = source22_.cf22
                d_1742___mcc_h4_ = source22_.cf23
                d_1743___mcc_h5_ = source22_.cf24
                d_1744_cf24_ = d_1743___mcc_h5_
                d_1745_cf23_ = d_1742___mcc_h4_
                d_1746_cf22_ = d_1741___mcc_h3_
                d_1747_v22_: str
                d_1747_v22_ = _dafny.CodePoint('u')
                d_1748_v23_: _dafny.Map
                d_1748_v23_ = _dafny.Map({D2_DC6(): default__.fm1(d_1708_v0_, d_1747_v22_, globalState)})
                d_1749_v24_: D2
                d_1749_v24_ = D2_DC6()
                d_1750_v25_: C0
                nw294_ = C0()
                nw294_.ctor__(d_1745_cf23_)
                d_1750_v25_ = nw294_
                d_1751_v26_: _dafny.Map
                d_1751_v26_ = _dafny.Map({d_1747_v22_: d_1750_v25_})
                d_1752_v27_: _dafny.Seq
                d_1752_v27_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "nhbbrfs"))
                d_1748_v23_ = (d_1748_v23_).set(d_1749_v24_, (len(d_1751_v26_)) < (len(d_1752_v27_)))
                d_1753_v28_: _dafny.Array
                nw295_ = _dafny.Array(None, 4)
                nw295_[int(0)] = p0
                nw295_[int(1)] = p0
                nw295_[int(2)] = 588
                nw295_[int(3)] = (20) * (633)
                d_1753_v28_ = nw295_
                d_1753_v28_ = d_1753_v28_
                (globalState).f9 = p0
                d_1754_v29_: _dafny.Map
                d_1754_v29_ = _dafny.Map({782: d_1728_i2_})
                d_1755_v30_: _dafny.MultiSet
                d_1755_v30_ = _dafny.MultiSet([len(d_1754_v29_)])
                d_1756_v31_: _dafny.Set
                d_1756_v31_ = _dafny.Set({d_1755_v30_})
                d_1757_v32_: _dafny.Map
                d_1757_v32_ = _dafny.Map({d_1728_i2_: False})
                d_1758_v33_: _dafny.Map
                d_1758_v33_ = _dafny.Map({d_1746_cf22_: ((d_1757_v32_)[d_1728_i2_] if (d_1728_i2_) in (d_1757_v32_) else d_1745_cf23_)})
                d_1759_v34_: _dafny.MultiSet
                d_1759_v34_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_1747_v22_ for d_1760_i3_ in range(default__.abs(335))])])
                d_1761_v35_: _dafny.Seq
                d_1761_v35_ = _dafny.SeqWithoutIsStrInference([d_1728_i2_])
                d_1762_v36_: _dafny.Seq
                d_1762_v36_ = _dafny.SeqWithoutIsStrInference([d_1746_cf22_, (d_1756_v31_).isdisjoint(default__.fm18(d_1758_v33_, d_1759_v34_, d_1747_v22_, (d_1761_v35_)[default__.safeIndex(d_1728_i2_, len(d_1761_v35_))], globalState)), (d_1755_v30_).ispropersubset(d_1755_v30_)])
                rhs298_ = d_1750_v25_
                rhs299_ = d_1708_v0_
                rhs300_ = default__.fm1((d_1745_cf23_) == (d_1744_cf24_), _dafny.CodePoint('a'), globalState)
                rhs301_ = d_1762_v36_
                rhs302_ = (d_1750_v25_).f17
                lhs218_ = globalState
                d_1750_v25_ = rhs298_
                lhs218_.f8 = rhs299_
                r3 = rhs300_
                d_1762_v36_ = rhs301_
                d_1744_cf24_ = rhs302_
            elif True:
                d_1763___mcc_h6_ = source22_.cf18
                d_1764_cf18_ = d_1763___mcc_h6_
                d_1765_v37_: _dafny.Map
                d_1765_v37_ = _dafny.Map({d_1708_v0_: d_1728_i2_})
                d_1765_v37_ = (d_1765_v37_).set(d_1708_v0_, d_1728_i2_)
                d_1766_v38_: _dafny.Seq
                d_1766_v38_ = _dafny.SeqWithoutIsStrInference([799, p0])
                d_1766_v38_ = (_dafny.SeqWithoutIsStrInference([d_1728_i2_ for d_1767_i4_ in range(default__.abs(-124))])).set(default__.safeIndex(default__.safeModuloInt(p0, p0), len(_dafny.SeqWithoutIsStrInference([d_1728_i2_ for d_1768_i4_ in range(default__.abs(-124))]))), (d_1766_v38_)[default__.safeIndex(655, len(d_1766_v38_))])
                (globalState).f9 = d_1728_i2_
                d_1769_v39_: C0
                nw296_ = C0()
                nw296_.ctor__(not((d_1766_v38_) < (d_1766_v38_)))
                d_1769_v39_ = nw296_
            d_1770_v40_: T1
            nw297_ = C2()
            nw297_.ctor__((self).f11)
            d_1770_v40_ = nw297_
            d_1771_v41_: _dafny.Map
            d_1771_v41_ = _dafny.Map({p0: d_1770_v40_})
            r1 = len((d_1771_v41_) | (d_1771_v41_))
            d_1772_v42_: _dafny.MultiSet
            d_1772_v42_ = _dafny.MultiSet([d_1708_v0_, d_1708_v0_])
            d_1773_v43_: _dafny.MultiSet
            d_1773_v43_ = _dafny.MultiSet([d_1772_v42_, d_1772_v42_])
            d_1774_v44_: _dafny.Seq
            d_1774_v44_ = _dafny.SeqWithoutIsStrInference([d_1708_v0_])
            d_1775_v45_: _dafny.Map
            d_1775_v45_ = _dafny.Map({p0: d_1708_v0_})
            d_1776_v46_: _dafny.Seq
            d_1776_v46_ = _dafny.SeqWithoutIsStrInference([len(d_1775_v45_), d_1728_i2_])
            d_1777_v47_: _dafny.Seq
            d_1777_v47_ = _dafny.SeqWithoutIsStrInference([d_1776_v46_])
            d_1778_v48_: _dafny.Set
            d_1778_v48_ = _dafny.Set({len((d_1777_v47_).set(default__.safeIndex(len(d_1774_v44_), len(d_1777_v47_)), d_1776_v46_)), d_1728_i2_})
            (globalState).f9 = ((d_1773_v43_)[_dafny.MultiSet((d_1774_v44_) + (d_1774_v44_))] if (_dafny.MultiSet((d_1774_v44_) + (d_1774_v44_))) in (d_1773_v43_) else len(d_1778_v48_))
            r3 = default__.fm1(d_1708_v0_, _dafny.CodePoint('h'), globalState)
        d_1779_v49_: _dafny.Seq
        d_1779_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iaxo"))
        d_1780_v50_: _dafny.MultiSet
        d_1780_v50_ = _dafny.MultiSet([len(d_1779_v49_)])
        if ((d_1780_v50_).intersection(d_1780_v50_)).ispropersubset(d_1780_v50_):
            d_1781_v51_: D0
            d_1781_v51_ = D0_DC0(p0)
            d_1782_v52_: C1
            nw298_ = C1()
            nw298_.ctor__(d_1708_v0_, not(not(d_1708_v0_)), d_1781_v51_, not(not (d_1708_v0_) or (d_1708_v0_)))
            d_1782_v52_ = nw298_
            d_1779_v49_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gkhk"))
            d_1783_v53_: _dafny.Array
            nw299_ = _dafny.Array(None, 10)
            nw299_[int(0)] = p1
            nw299_[int(1)] = p1
            nw299_[int(2)] = p1
            nw299_[int(3)] = p1
            nw299_[int(4)] = p1
            nw299_[int(5)] = p1
            nw299_[int(6)] = p1
            nw299_[int(7)] = p1
            nw299_[int(8)] = p1
            nw299_[int(9)] = p1
            d_1783_v53_ = nw299_
            index275_ = default__.safeIndex(816, (d_1783_v53_).length(0))
            (d_1783_v53_)[index275_] = p1
            d_1784_v54_: _dafny.Seq
            d_1784_v54_ = _dafny.SeqWithoutIsStrInference([(d_1782_v52_).f20, d_1708_v0_, (d_1782_v52_).f19, d_1708_v0_])
            index276_ = default__.safeIndex(44, (p1).length(0))
            (p1)[index276_] = len(d_1784_v54_)
            d_1785_v55_: _dafny.Map
            d_1785_v55_ = _dafny.Map({d_1708_v0_: p1})
            index277_ = default__.safeIndex(598, (d_1783_v53_).length(0))
            (d_1783_v53_)[index277_] = ((d_1785_v55_)[(d_1784_v54_)[default__.safeIndex(p0, len(d_1784_v54_))]] if ((d_1784_v54_)[default__.safeIndex(p0, len(d_1784_v54_))]) in (d_1785_v55_) else p1)
            d_1786_v56_: _dafny.MultiSet
            d_1786_v56_ = _dafny.MultiSet([d_1708_v0_, (d_1782_v52_).f19, not((d_1782_v52_).f19)])
            d_1787_v57_: D0
            d_1787_v57_ = D0_DC2(d_1786_v56_, (d_1782_v52_).f19, 325, False)
            d_1788_v58_: _dafny.Map
            d_1788_v58_ = _dafny.Map({p0: d_1787_v57_})
            index278_ = default__.safeIndex(816, (d_1783_v53_).length(0))
            index279_ = default__.safeIndex(44, (p1).length(0))
            index280_ = default__.safeIndex(598, (d_1783_v53_).length(0))
            rhs303_ = p1
            rhs304_ = (len(d_1788_v58_)) * (default__.fm0((d_1782_v52_).f19, globalState))
            rhs305_ = p1
            lhs219_ = d_1783_v53_
            lhs220_ = default__.safeIndex(816, (d_1783_v53_).length(0))
            lhs221_ = p1
            lhs222_ = default__.safeIndex(44, (p1).length(0))
            lhs223_ = d_1783_v53_
            lhs224_ = default__.safeIndex(598, (d_1783_v53_).length(0))
            lhs219_[lhs220_] = rhs303_
            lhs221_[lhs222_] = rhs304_
            lhs223_[lhs224_] = rhs305_
            d_1789_v59_: str
            d_1789_v59_ = _dafny.CodePoint('j')
            d_1708_v0_ = (d_1789_v59_) in (d_1779_v49_)
            (globalState).f9 = default__.safeDivisionInt(p0, 438)
        elif True:
            d_1790_v60_: _dafny.Map
            d_1790_v60_ = _dafny.Map({p1: d_1708_v0_})
            rhs306_ = ((d_1790_v60_)[p1] if (p1) in (d_1790_v60_) else d_1708_v0_)
            rhs307_ = default__.safeModuloInt(p0, (0) - (p0))
            lhs225_ = globalState
            r3 = rhs306_
            lhs225_.f9 = rhs307_
            d_1791_v61_: str
            d_1791_v61_ = _dafny.CodePoint('t')
            d_1708_v0_ = default__.fm1(d_1708_v0_, d_1791_v61_, globalState)
            d_1792_v62_: T0
            nw300_ = C3()
            nw300_.ctor__(d_1708_v0_, d_1708_v0_, (self).f11)
            d_1792_v62_ = nw300_
            d_1792_v62_ = d_1792_v62_
            d_1793_v63_: _dafny.Array
            nw301_ = _dafny.Array(False, 16)
            d_1793_v63_ = nw301_
            index281_ = default__.safeIndex(235, (d_1793_v63_).length(0))
            (d_1793_v63_)[index281_] = default__.fm1(default__.fm1(d_1708_v0_, d_1791_v61_, globalState), d_1791_v61_, globalState)
            d_1794_v64_: _dafny.Map
            d_1794_v64_ = _dafny.Map({d_1708_v0_: True})
            d_1795_v65_: _dafny.Map
            d_1795_v65_ = _dafny.Map({d_1794_v64_: d_1708_v0_})
            d_1796_v66_: _dafny.Map
            d_1796_v66_ = _dafny.Map({len(d_1795_v65_): p0})
            index282_ = default__.safeIndex(235, (d_1793_v63_).length(0))
            (d_1793_v63_)[index282_] = (p0) >= (((d_1796_v66_)[len(d_1796_v66_)] if (len(d_1796_v66_)) in (d_1796_v66_) else -893))
            d_1797_v67_: _dafny.MultiSet
            d_1797_v67_ = _dafny.MultiSet([p1, p1, p1])
            if (d_1797_v67_).issubset(d_1797_v67_):
                d_1798_v68_: _dafny.Seq
                d_1798_v68_ = _dafny.SeqWithoutIsStrInference([790, len(d_1779_v49_)])
                (globalState).f8 = (len(d_1798_v68_)) <= (len(d_1798_v68_))
                index283_ = default__.safeIndex(235, (d_1793_v63_).length(0))
                (d_1793_v63_)[index283_] = (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))]
                d_1799_v69_: _dafny.Array
                def lambda107_(d_1800_v64_):
                    def lambda108_(d_1801_i5_):
                        return d_1800_v64_

                    return lambda108_

                init58_ = lambda107_(d_1794_v64_)
                nw302_ = _dafny.Array(None, 24)
                for i0_58_ in range(nw302_.length(0)):
                    nw302_[i0_58_] = init58_(i0_58_)
                d_1799_v69_ = nw302_
                index284_ = default__.safeIndex(654, (p1).length(0))
                (p1)[index284_] = (0) - (len(d_1798_v68_))
                d_1802_v70_: _dafny.MultiSet
                d_1802_v70_ = _dafny.MultiSet([(d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], not(d_1708_v0_)])
                index285_ = default__.safeIndex(235, (d_1793_v63_).length(0))
                index286_ = default__.safeIndex(654, (p1).length(0))
                rhs308_ = d_1799_v69_
                rhs309_ = d_1708_v0_
                rhs310_ = ((d_1779_v49_) + (((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouk"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouk")))), d_1791_v61_)).set(default__.safeIndex(p0, len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouk"))).set(default__.safeIndex(p0, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ouk")))), d_1791_v61_))), d_1791_v61_)) if not(not ((d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))]) or (False)) else d_1779_v49_)
                rhs311_ = default__.fm1((d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], d_1791_v61_, globalState)
                rhs312_ = (((d_1802_v70_)[(d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))]] if ((d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))]) in (d_1802_v70_) else 448)) - (p0)
                lhs226_ = d_1793_v63_
                lhs227_ = default__.safeIndex(235, (d_1793_v63_).length(0))
                lhs228_ = globalState
                lhs229_ = p1
                lhs230_ = default__.safeIndex(654, (p1).length(0))
                d_1799_v69_ = rhs308_
                lhs226_[lhs227_] = rhs309_
                d_1779_v49_ = rhs310_
                lhs228_.f8 = rhs311_
                lhs229_[lhs230_] = rhs312_
                d_1794_v64_ = (d_1794_v64_).set((d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], d_1708_v0_)
                (globalState).f9 = p0
            elif True:
                d_1803_v71_: _dafny.Seq
                d_1803_v71_ = _dafny.SeqWithoutIsStrInference([default__.fm1(d_1708_v0_, d_1791_v61_, globalState)])
                d_1804_v72_: _dafny.Set
                d_1804_v72_ = _dafny.Set({d_1794_v64_, (d_1794_v64_) | (d_1794_v64_), (d_1794_v64_ if not(d_1708_v0_) else d_1794_v64_), (d_1794_v64_ if not(default__.fm1((d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], d_1791_v61_, globalState)) else default__.fm34(d_1708_v0_, (_dafny.MultiSet(d_1803_v71_)).set(d_1708_v0_, default__.abs(len(d_1779_v49_))), globalState)), d_1794_v64_})
                d_1805_v73_: _dafny.Seq
                d_1805_v73_ = _dafny.SeqWithoutIsStrInference([d_1804_v72_])
                d_1804_v72_ = (((d_1805_v73_)[default__.safeIndex(p0, len(d_1805_v73_))]) - (d_1804_v72_)).intersection(_dafny.Set({d_1794_v64_, d_1794_v64_, d_1794_v64_, _dafny.Map({(d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))]: (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))]})}))
                d_1806_v74_: _dafny.Set
                d_1806_v74_ = _dafny.Set({50})
                d_1806_v74_ = (_dafny.Set({806}) if (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))] else d_1806_v74_)
                d_1807_v75_: _dafny.Seq
                d_1807_v75_ = _dafny.SeqWithoutIsStrInference([d_1793_v63_, d_1793_v63_])
                d_1808_v76_: _dafny.Seq
                d_1808_v76_ = _dafny.SeqWithoutIsStrInference([((d_1807_v75_)[default__.safeIndex(p0, len(d_1807_v75_))] if False else d_1793_v63_), d_1793_v63_])
                d_1793_v63_ = (d_1807_v75_)[default__.safeIndex(default__.safeDivisionInt(588, len(d_1806_v74_)), len(d_1807_v75_))]
                d_1809_v77_: _dafny.Map
                d_1809_v77_ = _dafny.Map({d_1791_v61_: d_1779_v49_})
                d_1779_v49_ = (((d_1809_v77_)[default__.fm3(p0, (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], globalState)] if (default__.fm3(p0, (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], globalState)) in (d_1809_v77_) else (d_1779_v49_) + ((d_1779_v49_).set(default__.safeIndex(p0, len(d_1779_v49_)), d_1791_v61_)))).set(default__.safeIndex(default__.safeDivisionInt(p0, p0), len(((d_1809_v77_)[default__.fm3(p0, (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], globalState)] if (default__.fm3(p0, (d_1793_v63_)[default__.safeIndex(235, (d_1793_v63_).length(0))], globalState)) in (d_1809_v77_) else (d_1779_v49_) + ((d_1779_v49_).set(default__.safeIndex(p0, len(d_1779_v49_)), d_1791_v61_))))), d_1791_v61_)
                d_1803_v71_ = d_1803_v71_
        if (p0) > ((0) - (p0)):
            (globalState).f9 = (p0) - (p0)
            d_1810_v78_: _dafny.Map
            d_1810_v78_ = _dafny.Map({d_1708_v0_: d_1708_v0_})
            d_1811_v79_: str
            d_1811_v79_ = _dafny.CodePoint('m')
            if not (not (((d_1810_v78_)[d_1708_v0_] if (d_1708_v0_) in (d_1810_v78_) else default__.fm1(d_1708_v0_, d_1811_v79_, globalState))) or (d_1708_v0_)) or (d_1708_v0_):
                r3 = d_1708_v0_
                (globalState).f9 = p0
                d_1812_v80_: D0
                d_1812_v80_ = D0_DC0(len(default__.fm50(d_1811_v79_, d_1708_v0_, d_1708_v0_, globalState)))
                d_1813_v81_: _dafny.Set
                d_1813_v81_ = _dafny.Set({d_1708_v0_, d_1708_v0_, d_1708_v0_})
                d_1814_v82_: D4
                d_1814_v82_ = D4_DC10(d_1813_v81_)
                d_1815_v83_: C1
                nw303_ = C1()
                nw303_.ctor__(d_1708_v0_, d_1708_v0_, d_1812_v80_, ((d_1814_v82_).cf18).ispropersubset(d_1813_v81_))
                d_1815_v83_ = nw303_
                d_1816_v84_: _dafny.MultiSet
                d_1816_v84_ = _dafny.MultiSet([d_1708_v0_])
                d_1816_v84_ = d_1816_v84_
                d_1708_v0_ = False
            elif True:
                rhs313_ = p0
                rhs314_ = default__.safeDivisionInt(p0, default__.fm0(d_1708_v0_, globalState))
                rhs315_ = (d_1779_v49_) + ((d_1779_v49_) + (d_1779_v49_))
                lhs231_ = globalState
                lhs231_.f9 = rhs313_
                r1 = rhs314_
                d_1779_v49_ = rhs315_
                (globalState).f9 = (0) - (p0)
                d_1817_v85_: _dafny.Set
                d_1817_v85_ = _dafny.Set({d_1708_v0_})
                d_1818_v86_: _dafny.Map
                d_1818_v86_ = _dafny.Map({default__.safeDivisionInt(p0, -685): d_1708_v0_})
                d_1819_v87_: _dafny.Array
                nw304_ = _dafny.Array(False, 6)
                d_1819_v87_ = nw304_
                d_1820_v88_: _dafny.MultiSet
                d_1820_v88_ = _dafny.MultiSet([d_1819_v87_, d_1819_v87_, d_1819_v87_])
                d_1821_v89_: _dafny.Seq
                d_1821_v89_ = _dafny.SeqWithoutIsStrInference([p0, (0) - (p0)])
                rhs316_ = d_1708_v0_
                rhs317_ = d_1708_v0_
                rhs318_ = (d_1708_v0_) in ((d_1817_v85_ if default__.fm1(d_1708_v0_, d_1811_v79_, globalState) else d_1817_v85_))
                rhs319_ = ((d_1818_v86_)[len(d_1779_v49_)] if (len(d_1779_v49_)) in (d_1818_v86_) else (d_1820_v88_) == (d_1820_v88_))
                rhs320_ = (d_1821_v89_) <= (d_1821_v89_)
                lhs232_ = globalState
                lhs233_ = globalState
                lhs234_ = globalState
                lhs235_ = globalState
                lhs232_.f8 = rhs316_
                lhs233_.f8 = rhs317_
                lhs234_.f8 = rhs318_
                lhs235_.f8 = rhs319_
                r3 = rhs320_
                d_1810_v78_ = (d_1810_v78_).set(True, (p0) > (p0))
                (globalState).f9 = len(default__.fm39(p0, globalState))
            (globalState).f9 = p0
            r1 = p0
            default__.m0(globalState)
        elif True:
            index287_ = default__.safeIndex(193, (p1).length(0))
            (p1)[index287_] = 770
            d_1822_v90_: _dafny.Seq
            d_1822_v90_ = _dafny.SeqWithoutIsStrInference([-29])
            index288_ = default__.safeIndex(193, (p1).length(0))
            (p1)[index288_] = (p0 if d_1708_v0_ else len(d_1822_v90_))
            d_1823_v91_: D5
            d_1823_v91_ = D5_DC13(_dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(193, (p1).length(0))] for d_1824_i6_ in range(default__.abs(643))]))
            d_1825_v92_: _dafny.Seq
            d_1825_v92_ = _dafny.SeqWithoutIsStrInference([d_1708_v0_, False])
            d_1826_v93_: _dafny.Set
            d_1826_v93_ = _dafny.Set({d_1708_v0_, d_1708_v0_, d_1708_v0_})
            d_1827_v94_: _dafny.Array
            nw305_ = _dafny.Array(None, 6)
            nw305_[int(0)] = (p1)[default__.safeIndex(193, (p1).length(0))]
            nw305_[int(1)] = p0
            nw305_[int(2)] = (0) - ((p1)[default__.safeIndex(193, (p1).length(0))])
            nw305_[int(3)] = (len(d_1825_v92_)) + ((p1)[default__.safeIndex(193, (p1).length(0))])
            nw305_[int(4)] = (0) - (p0)
            nw305_[int(5)] = len(d_1826_v93_)
            d_1827_v94_ = nw305_
            rhs321_ = d_1823_v91_
            rhs322_ = d_1708_v0_
            rhs323_ = d_1827_v94_
            rhs324_ = ((0) - ((p1)[default__.safeIndex(193, (p1).length(0))])) + (default__.safeDivisionInt(418, (p1)[default__.safeIndex(193, (p1).length(0))]))
            lhs236_ = globalState
            d_1823_v91_ = rhs321_
            lhs236_.f8 = rhs322_
            d_1827_v94_ = rhs323_
            r1 = rhs324_
            d_1828_v95_: C2
            nw306_ = C2()
            nw306_.ctor__((self).f11)
            d_1828_v95_ = nw306_
            d_1829_v96_: _dafny.MultiSet
            d_1829_v96_ = _dafny.MultiSet([d_1708_v0_])
            (globalState).f9 = default__.safeDivisionInt((20) * (p0), ((d_1829_v96_) - (d_1829_v96_)).cardinality)
            d_1830_v97_: str
            d_1830_v97_ = _dafny.CodePoint('o')
            index289_ = default__.safeIndex(193, (p1).length(0))
            index290_ = default__.safeIndex(193, (p1).length(0))
            rhs325_ = d_1830_v97_
            rhs326_ = (p0) + (default__.safeModuloInt((p1)[default__.safeIndex(193, (p1).length(0))], (0) - (p0)))
            rhs327_ = (0) - (-409)
            lhs237_ = p1
            lhs238_ = default__.safeIndex(193, (p1).length(0))
            lhs239_ = p1
            lhs240_ = default__.safeIndex(193, (p1).length(0))
            r0 = rhs325_
            lhs237_[lhs238_] = rhs326_
            lhs239_[lhs240_] = rhs327_
        default__.m0(globalState)
        d_1831_v98_: str
        d_1831_v98_ = _dafny.CodePoint('l')
        d_1832_i7_: int
        d_1832_i7_ = 0
        with _dafny.label("8"):
            while default__.fm1((p0) < (p0), d_1831_v98_, globalState):
                with _dafny.c_label("8"):
                    if (d_1832_i7_) >= (100):
                        raise _dafny.Break("8")
                    d_1832_i7_ = (d_1832_i7_) + (1)
                    if d_1708_v0_:
                        d_1833_v99_: D5
                        d_1833_v99_ = D5_DC14((d_1779_v49_) != (d_1779_v49_))
                        pat_let_tv37_ = d_1708_v0_
                        def iife122_(_pat_let29_0):
                            def iife123_(d_1834_dt__update__tmp_h0_):
                                def iife124_(_pat_let30_0):
                                    def iife125_(d_1835_dt__update_hcf26_h0_):
                                        return D5_DC14(d_1835_dt__update_hcf26_h0_)
                                    return iife125_(_pat_let30_0)
                                return iife124_(not(not(pat_let_tv37_)))
                            return iife123_(_pat_let29_0)
                        d_1833_v99_ = iife122_(d_1833_v99_)
                        d_1836_v100_: _dafny.Array
                        nw307_ = _dafny.Array(None, 25)
                        nw307_[int(0)] = d_1831_v98_
                        nw307_[int(1)] = d_1831_v98_
                        nw307_[int(2)] = d_1831_v98_
                        nw307_[int(3)] = _dafny.CodePoint('i')
                        nw307_[int(4)] = _dafny.CodePoint('r')
                        nw307_[int(5)] = d_1831_v98_
                        nw307_[int(6)] = d_1831_v98_
                        nw307_[int(7)] = d_1831_v98_
                        nw307_[int(8)] = (d_1779_v49_)[default__.safeIndex(606, len(d_1779_v49_))]
                        nw307_[int(9)] = d_1831_v98_
                        nw307_[int(10)] = (d_1779_v49_)[default__.safeIndex(p0, len(d_1779_v49_))]
                        nw307_[int(11)] = d_1831_v98_
                        nw307_[int(12)] = _dafny.CodePoint('o')
                        nw307_[int(13)] = d_1831_v98_
                        nw307_[int(14)] = d_1831_v98_
                        nw307_[int(15)] = d_1831_v98_
                        nw307_[int(16)] = (d_1779_v49_)[default__.safeIndex(214, len(d_1779_v49_))]
                        nw307_[int(17)] = d_1831_v98_
                        nw307_[int(18)] = d_1831_v98_
                        nw307_[int(19)] = default__.fm3(p0, False, globalState)
                        nw307_[int(20)] = d_1831_v98_
                        nw307_[int(21)] = _dafny.CodePoint('l')
                        nw307_[int(22)] = d_1831_v98_
                        nw307_[int(23)] = d_1831_v98_
                        nw307_[int(24)] = d_1831_v98_
                        d_1836_v100_ = nw307_
                        d_1837_v101_: _dafny.Array
                        nw308_ = _dafny.Array(None, 2)
                        nw308_[int(0)] = d_1836_v100_
                        nw308_[int(1)] = d_1836_v100_
                        d_1837_v101_ = nw308_
                        d_1838_v102_: _dafny.MultiSet
                        d_1838_v102_ = _dafny.MultiSet([d_1831_v98_, d_1831_v98_, d_1831_v98_])
                        index291_ = default__.safeIndex(214, (p1).length(0))
                        (p1)[index291_] = p0
                        d_1839_v103_: _dafny.Seq
                        d_1839_v103_ = _dafny.SeqWithoutIsStrInference([d_1708_v0_, d_1708_v0_, d_1708_v0_, d_1708_v0_])
                        d_1840_v104_: _dafny.Set
                        d_1840_v104_ = _dafny.Set({p0, p0})
                        index292_ = default__.safeIndex(214, (p1).length(0))
                        rhs328_ = d_1837_v101_
                        rhs329_ = d_1838_v102_
                        rhs330_ = default__.fm0(((d_1839_v103_)[default__.safeIndex(p0, len(d_1839_v103_))]) or (d_1708_v0_), globalState)
                        rhs331_ = (p0) not in (d_1840_v104_)
                        lhs241_ = p1
                        lhs242_ = default__.safeIndex(214, (p1).length(0))
                        d_1837_v101_ = rhs328_
                        d_1838_v102_ = rhs329_
                        lhs241_[lhs242_] = rhs330_
                        r3 = rhs331_
                        d_1841_v105_: _dafny.Map
                        d_1841_v105_ = _dafny.Map({default__.fm0(d_1708_v0_, globalState): d_1708_v0_})
                        d_1842_v106_: D4
                        d_1842_v106_ = D4_DC11(len(d_1841_v105_), (p1)[default__.safeIndex(214, (p1).length(0))], d_1708_v0_)
                        d_1843_v107_: D0
                        d_1843_v107_ = D0_DC0((p1)[default__.safeIndex(214, (p1).length(0))])
                        d_1844_v108_: _dafny.Map
                        d_1844_v108_ = _dafny.Map({_dafny.CodePoint('p'): (p1)[default__.safeIndex(214, (p1).length(0))]})
                        d_1845_v109_: C1
                        nw309_ = C1()
                        nw309_.ctor__((d_1842_v106_).cf21, d_1708_v0_, d_1843_v107_, (len(d_1844_v108_)) != (p0))
                        d_1845_v109_ = nw309_
                        (globalState).f9 = p0
                        (globalState).f8 = (d_1845_v109_).f19
                    elif True:
                        d_1846_v111_: _dafny.Map
                        d_1846_v111_ = _dafny.Map({p0: p0})
                        d_1847_v112_: _dafny.Array
                        nw310_ = _dafny.Array(None, 15)
                        nw310_[int(0)] = d_1779_v49_
                        nw310_[int(1)] = (_dafny.SeqWithoutIsStrInference([d_1831_v98_ for d_1848_i8_ in range(default__.abs(-820))]) if True else d_1779_v49_)
                        nw310_[int(2)] = ((d_1779_v49_) + (d_1779_v49_)).set(default__.safeIndex(p0, len((d_1779_v49_) + (d_1779_v49_))), d_1831_v98_)
                        nw310_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1831_v98_, d_1831_v98_, d_1831_v98_])
                        nw310_[int(4)] = (d_1779_v49_) + (_dafny.SeqWithoutIsStrInference([d_1831_v98_, d_1831_v98_]))
                        nw310_[int(5)] = (d_1779_v49_ if d_1708_v0_ else d_1779_v49_)
                        nw310_[int(6)] = d_1779_v49_
                        nw310_[int(7)] = (d_1779_v49_) + (_dafny.SeqWithoutIsStrInference([d_1831_v98_]))
                        nw310_[int(8)] = _dafny.SeqWithoutIsStrInference([d_1831_v98_])
                        nw310_[int(9)] = d_1779_v49_
                        def iife126_():
                            coll64_ = _dafny.Map()
                            compr_64_: int
                            for compr_64_ in (d_1780_v50_).Elements:
                                d_1849_v110_: int = compr_64_
                                if (d_1849_v110_) in (d_1780_v50_):
                                    coll64_[(d_1849_v110_) + (p0)] = 885
                            return _dafny.Map(coll64_)
                        nw310_[int(10)] = (_dafny.SeqWithoutIsStrInference([d_1831_v98_, d_1831_v98_])) + (default__.fm2(p0, d_1780_v50_, iife126_()
                        , globalState))
                        nw310_[int(11)] = default__.fm2(p0, d_1780_v50_, d_1846_v111_, globalState)
                        nw310_[int(12)] = _dafny.SeqWithoutIsStrInference([d_1831_v98_ for d_1850_i9_ in range(default__.abs(681))])
                        nw310_[int(13)] = d_1779_v49_
                        nw310_[int(14)] = d_1779_v49_
                        d_1847_v112_ = nw310_
                        index293_ = default__.safeIndex(5, (d_1847_v112_).length(0))
                        (d_1847_v112_)[index293_] = d_1779_v49_
                        index294_ = default__.safeIndex(5, (d_1847_v112_).length(0))
                        (d_1847_v112_)[index294_] = (D2_DC5(d_1779_v49_)).cf10
                        (globalState).f8 = d_1708_v0_
                        d_1851_v113_: _dafny.Array
                        nw311_ = _dafny.Array(D2.default()(), 26)
                        d_1851_v113_ = nw311_
                        d_1851_v113_ = d_1851_v113_
                        d_1852_v114_: C0
                        nw312_ = C0()
                        nw312_.ctor__(True)
                        d_1852_v114_ = nw312_
                        (globalState).f9 = p0
                    d_1853_v115_: _dafny.Seq
                    d_1853_v115_ = _dafny.SeqWithoutIsStrInference([d_1708_v0_])
                    d_1831_v98_ = (d_1831_v98_ if (d_1853_v115_)[default__.safeIndex((d_1780_v50_).cardinality, len(d_1853_v115_))] else d_1831_v98_)
                    index295_ = default__.safeIndex(102, (p1).length(0))
                    (p1)[index295_] = p0
                    index296_ = default__.safeIndex(102, (p1).length(0))
                    (p1)[index296_] = p0
                    d_1854_v116_: C3
                    nw313_ = C3()
                    nw313_.ctor__(d_1708_v0_, default__.fm1(d_1708_v0_, d_1831_v98_, globalState), (self).f11)
                    d_1854_v116_ = nw313_
                    d_1855_v117_: _dafny.Map
                    d_1855_v117_ = _dafny.Map({d_1854_v116_: (d_1854_v116_).f22})
                    d_1856_v118_: _dafny.Seq
                    d_1856_v118_ = _dafny.SeqWithoutIsStrInference([(p1)[default__.safeIndex(102, (p1).length(0))], 526])
                    d_1857_v119_: _dafny.MultiSet
                    d_1857_v119_ = _dafny.MultiSet([d_1856_v118_])
                    d_1855_v117_ = (d_1855_v117_).set(d_1854_v116_, (_dafny.MultiSet([(D5_DC13(_dafny.SeqWithoutIsStrInference([421 for d_1858_i10_ in range(default__.abs(335))]))).cf25])).isdisjoint(d_1857_v119_))
                    pass
            pass
        r0 = d_1831_v98_
        r1 = default__.safeModuloInt((0) - (len((d_1779_v49_ if False else d_1779_v49_))), (p0) + (p0))
        d_1859_v120_: D2
        d_1859_v120_ = D2_DC8(p0, d_1831_v98_)
        r2 = d_1859_v120_
        r3 = d_1708_v0_
        return r0, r1, r2, r3

    def m5(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: bool = False
        r2: bool = False
        d_1860_v0_: bool
        d_1860_v0_ = True
        d_1861_i0_: int
        d_1861_i0_ = 0
        with _dafny.label("9"):
            while not (d_1860_v0_) or (not(False)):
                with _dafny.c_label("9"):
                    if (d_1861_i0_) >= (100):
                        raise _dafny.Break("9")
                    d_1861_i0_ = (d_1861_i0_) + (1)
                    d_1862_v1_: int
                    d_1862_v1_ = 631
                    d_1863_v2_: _dafny.MultiSet
                    d_1863_v2_ = _dafny.MultiSet([d_1860_v0_])
                    d_1864_v3_: _dafny.Seq
                    d_1864_v3_ = _dafny.SeqWithoutIsStrInference([-628])
                    d_1865_v4_: _dafny.Seq
                    d_1865_v4_ = _dafny.SeqWithoutIsStrInference([((d_1863_v2_)[d_1860_v0_] if (d_1860_v0_) in (d_1863_v2_) else d_1862_v1_), d_1862_v1_, (d_1864_v3_)[default__.safeIndex(d_1862_v1_, len(d_1864_v3_))]])
                    rhs332_ = d_1862_v1_
                    rhs333_ = len((_dafny.SeqWithoutIsStrInference([d_1862_v1_ for d_1866_i1_ in range(default__.abs(137))]) if d_1860_v0_ else d_1865_v4_))
                    lhs243_ = globalState
                    lhs244_ = globalState
                    lhs243_.f9 = rhs332_
                    lhs244_.f9 = rhs333_
                    d_1860_v0_ = True
                    d_1867_v5_: str
                    d_1867_v5_ = _dafny.CodePoint('c')
                    d_1868_v6_: _dafny.Map
                    d_1868_v6_ = _dafny.Map({d_1862_v1_: 921})
                    d_1869_v7_: _dafny.Map
                    d_1869_v7_ = _dafny.Map({d_1867_v5_: d_1868_v6_})
                    d_1870_v8_: D5
                    d_1870_v8_ = D5_DC13((((d_1864_v3_).set(default__.safeIndex(d_1862_v1_, len(d_1864_v3_)), d_1862_v1_)) + (d_1865_v4_)).set(default__.safeIndex(len(d_1865_v4_), len(((d_1864_v3_).set(default__.safeIndex(d_1862_v1_, len(d_1864_v3_)), d_1862_v1_)) + (d_1865_v4_))), len(((d_1869_v7_)[d_1867_v5_] if (d_1867_v5_) in (d_1869_v7_) else d_1868_v6_))))
                    d_1870_v8_ = d_1870_v8_
                    d_1871_v9_: _dafny.Array
                    nw314_ = _dafny.Array(int(0), 24)
                    d_1871_v9_ = nw314_
                    index297_ = default__.safeIndex(850, (d_1871_v9_).length(0))
                    (d_1871_v9_)[index297_] = d_1862_v1_
                    index298_ = default__.safeIndex(850, (d_1871_v9_).length(0))
                    (d_1871_v9_)[index298_] = (0) - (d_1862_v1_)
                    pass
            pass
        default__.m0(globalState)
        d_1872_v10_: _dafny.Array
        def lambda109_(d_1873_v0_):
            def lambda110_(d_1874_i3_):
                return (_dafny.SeqWithoutIsStrInference([d_1873_v0_, d_1873_v0_])) + (_dafny.SeqWithoutIsStrInference([d_1873_v0_, d_1873_v0_]))

            return lambda110_

        init59_ = lambda109_(d_1860_v0_)
        nw315_ = _dafny.Array(None, 29)
        for i0_59_ in range(nw315_.length(0)):
            nw315_[i0_59_] = init59_(i0_59_)
        d_1872_v10_ = nw315_
        guard_loop_2_: int
        for guard_loop_2_ in _dafny.IntegerRange(0, (d_1872_v10_).length(0)):
            d_1875_i2_: int = guard_loop_2_
            if (True) and (((0) <= (d_1875_i2_)) and ((d_1875_i2_) < ((d_1872_v10_).length(0)))):
                (d_1872_v10_)[(d_1875_i2_)] = (_dafny.SeqWithoutIsStrInference([d_1860_v0_, not(d_1860_v0_), d_1860_v0_])) + (_dafny.SeqWithoutIsStrInference([d_1860_v0_, d_1860_v0_]))
        d_1876_v11_: int
        d_1876_v11_ = 985
        (globalState).f9 = (d_1876_v11_) - (d_1876_v11_)
        if (default__.fm0(d_1860_v0_, globalState)) < ((355) * (d_1876_v11_)):
            d_1877_v12_: _dafny.Array
            def lambda111_(d_1878_v0_):
                def lambda112_(d_1879_i4_):
                    return d_1878_v0_

                return lambda112_

            init60_ = lambda111_(d_1860_v0_)
            nw316_ = _dafny.Array(None, 16)
            for i0_60_ in range(nw316_.length(0)):
                nw316_[i0_60_] = init60_(i0_60_)
            d_1877_v12_ = nw316_
            d_1880_v13_: D5
            d_1880_v13_ = D5_DC14(d_1860_v0_)
            index299_ = default__.safeIndex(192, (d_1877_v12_).length(0))
            (d_1877_v12_)[index299_] = (d_1880_v13_).cf26
            index300_ = default__.safeIndex(192, (d_1877_v12_).length(0))
            (d_1877_v12_)[index300_] = (d_1860_v0_ if d_1860_v0_ else d_1860_v0_)
            d_1881_v14_: _dafny.Seq
            d_1881_v14_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Set({d_1876_v11_, d_1876_v11_}))])
            d_1882_v15_: C0
            nw317_ = C0()
            nw317_.ctor__(((d_1881_v14_)[default__.safeIndex(d_1876_v11_, len(d_1881_v14_))]) >= (d_1876_v11_))
            d_1882_v15_ = nw317_
            d_1883_v16_: _dafny.Seq
            d_1883_v16_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "kba"))
            d_1884_v17_: _dafny.Set
            d_1884_v17_ = _dafny.Set({(d_1882_v15_).f17, (d_1882_v15_).f17, (d_1882_v15_).f17})
            d_1885_v18_: _dafny.Seq
            d_1885_v18_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({(d_1877_v12_)[default__.safeIndex(192, (d_1877_v12_).length(0))], True}), (d_1884_v17_).intersection(d_1884_v17_), d_1884_v17_])
            rhs334_ = ((d_1883_v16_) + (d_1883_v16_)) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "fns")))
            rhs335_ = (0) - (d_1876_v11_)
            rhs336_ = (d_1885_v18_)[default__.safeIndex(default__.fm0((d_1882_v15_).f17, globalState), len(d_1885_v18_))]
            rhs337_ = True
            lhs245_ = globalState
            lhs246_ = globalState
            d_1883_v16_ = rhs334_
            lhs245_.f9 = rhs335_
            d_1884_v17_ = rhs336_
            lhs246_.f8 = rhs337_
            d_1886_v19_: str
            d_1886_v19_ = _dafny.CodePoint('u')
            d_1887_v20_: _dafny.Map
            d_1887_v20_ = _dafny.Map({default__.fm0(default__.fm1((d_1882_v15_).f17, d_1886_v19_, globalState), globalState): default__.safeDivisionInt(d_1876_v11_, d_1876_v11_)})
            d_1887_v20_ = (d_1887_v20_).set(default__.fm0((d_1877_v12_)[default__.safeIndex(192, (d_1877_v12_).length(0))], globalState), 518)
            d_1888_v21_: D7
            d_1888_v21_ = D7_DC17(((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))).set(default__.safeIndex(d_1876_v11_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y")))), d_1886_v19_)) + ((d_1883_v16_).set(default__.safeIndex(d_1876_v11_, len(d_1883_v16_)), default__.fm3(d_1876_v11_, True, globalState))), d_1877_v12_, d_1860_v0_, d_1876_v11_)
            source23_ = d_1888_v21_
            if source23_.is_DC17:
                d_1889___mcc_h0_ = source23_.cf29
                d_1890___mcc_h1_ = source23_.cf30
                d_1891___mcc_h2_ = source23_.cf31
                d_1892___mcc_h3_ = source23_.cf32
                d_1893_cf32_ = d_1892___mcc_h3_
                d_1894_cf31_ = d_1891___mcc_h2_
                d_1895_cf30_ = d_1890___mcc_h1_
                d_1896_cf29_ = d_1889___mcc_h0_
                d_1897_v22_: _dafny.Array
                def lambda113_(d_1898_cf29_):
                    def lambda114_(d_1899_i5_):
                        return d_1898_cf29_

                    return lambda114_

                init61_ = lambda113_(d_1896_cf29_)
                nw318_ = _dafny.Array(None, 27)
                for i0_61_ in range(nw318_.length(0)):
                    nw318_[i0_61_] = init61_(i0_61_)
                d_1897_v22_ = nw318_
                index301_ = default__.safeIndex(604, (d_1897_v22_).length(0))
                (d_1897_v22_)[index301_] = d_1883_v16_
                index302_ = default__.safeIndex(604, (d_1897_v22_).length(0))
                rhs338_ = _dafny.CodePoint('a')
                rhs339_ = (_dafny.SeqWithoutIsStrInference([d_1886_v19_ for d_1900_i6_ in range(default__.abs(-245))])) + (_dafny.SeqWithoutIsStrInference([d_1886_v19_ for d_1901_i7_ in range(default__.abs(-320))]))
                lhs247_ = d_1897_v22_
                lhs248_ = default__.safeIndex(604, (d_1897_v22_).length(0))
                r0 = rhs338_
                lhs247_[lhs248_] = rhs339_
                (globalState).f9 = d_1893_cf32_
                index303_ = default__.safeIndex(604, (d_1897_v22_).length(0))
                (d_1897_v22_)[index303_] = d_1883_v16_
                (globalState).f9 = (((d_1881_v14_)[default__.safeIndex(d_1893_cf32_, len(d_1881_v14_))]) * ((0) - (d_1893_cf32_))) * (d_1876_v11_)
            elif True:
                d_1902___mcc_h4_ = source23_.cf28
                d_1903_cf28_ = d_1902___mcc_h4_
                d_1904_v23_: D0
                d_1904_v23_ = D0_DC0(d_1876_v11_)
                d_1905_v24_: C1
                nw319_ = C1()
                nw319_.ctor__((d_1882_v15_).f17, (d_1882_v15_).f17, d_1904_v23_, (d_1877_v12_)[default__.safeIndex(192, (d_1877_v12_).length(0))])
                d_1905_v24_ = nw319_
                d_1906_v25_: C1
                nw320_ = C1()
                nw320_.ctor__((d_1876_v11_) >= (len(d_1887_v20_)), (d_1882_v15_).f17, D0_DC0(len(_dafny.SeqWithoutIsStrInference([d_1905_v24_]))), (d_1905_v24_).f20)
                d_1906_v25_ = nw320_
                d_1907_v26_: _dafny.Map
                d_1907_v26_ = _dafny.Map({d_1876_v11_: (d_1906_v25_).f20})
                d_1908_v27_: _dafny.MultiSet
                d_1908_v27_ = _dafny.MultiSet([(d_1905_v24_).f20])
                d_1909_v28_: D0
                d_1909_v28_ = D0_DC2(d_1908_v27_, (d_1906_v25_).f19, d_1876_v11_, (d_1905_v24_).f19)
                d_1910_v29_: _dafny.MultiSet
                d_1910_v29_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "pmk")), d_1883_v16_])
                rhs340_ = (d_1907_v26_).set(277, (d_1882_v15_).f17)
                rhs341_ = default__.safeModuloInt((0) - (d_1876_v11_), (0) - (d_1876_v11_))
                rhs342_ = (d_1909_v28_).cf7
                rhs343_ = (d_1910_v29_).isdisjoint(d_1910_v29_)
                lhs249_ = globalState
                d_1907_v26_ = rhs340_
                d_1876_v11_ = rhs341_
                r1 = rhs342_
                lhs249_.f8 = rhs343_
                d_1911_v30_: _dafny.Map
                d_1911_v30_ = _dafny.Map({59: _dafny.CodePoint('s')})
                d_1912_v31_: _dafny.Set
                d_1912_v31_ = _dafny.Set({d_1876_v11_, len(d_1883_v16_)})
                d_1913_v33_: _dafny.Array
                nw321_ = _dafny.Array(None, 24)
                nw321_[int(0)] = len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "gsddkqjn")) if (d_1903_cf28_)[default__.safeIndex(len(d_1911_v30_), len(d_1903_cf28_))] else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qbyry"))))
                nw321_[int(1)] = len(d_1912_v31_)
                nw321_[int(2)] = d_1876_v11_
                nw321_[int(3)] = (d_1905_v24_).fm9(default__.fm57(globalState), _dafny.Set({d_1876_v11_}), globalState)
                nw321_[int(4)] = d_1876_v11_
                nw321_[int(5)] = (d_1906_v25_).fm30(_dafny.Map({d_1876_v11_: (d_1877_v12_)[default__.safeIndex(192, (d_1877_v12_).length(0))]}), globalState)
                nw321_[int(6)] = d_1876_v11_
                nw321_[int(7)] = len(_dafny.Set({d_1876_v11_, len(d_1881_v14_)}))
                nw321_[int(8)] = d_1876_v11_
                nw321_[int(9)] = d_1876_v11_
                nw321_[int(10)] = (d_1876_v11_ if (d_1882_v15_).f17 else d_1876_v11_)
                nw321_[int(11)] = (d_1881_v14_)[default__.safeIndex(d_1876_v11_, len(d_1881_v14_))]
                nw321_[int(12)] = len((d_1883_v16_).set(default__.safeIndex(d_1876_v11_, len(d_1883_v16_)), d_1886_v19_))
                def iife127_():
                    coll65_ = _dafny.Set()
                    compr_65_: int
                    for compr_65_ in _dafny.IntegerRange(372, -617):
                        d_1914_v32_: int = compr_65_
                        if ((372) <= (d_1914_v32_)) and ((d_1914_v32_) < (-617)):
                            coll65_ = coll65_.union(_dafny.Set([default__.safeModuloInt(d_1914_v32_, (0) - (d_1876_v11_))]))
                    return _dafny.Set(coll65_)
                nw321_[int(13)] = ((d_1906_v25_).fm29(d_1876_v11_, globalState)) * (len(iife127_()
                ))
                nw321_[int(14)] = 332
                nw321_[int(15)] = d_1876_v11_
                nw321_[int(16)] = (0) - (d_1876_v11_)
                nw321_[int(17)] = (len(_dafny.SeqWithoutIsStrInference([(d_1905_v24_).f20, False]))) * (d_1876_v11_)
                nw321_[int(18)] = d_1876_v11_
                nw321_[int(19)] = default__.safeModuloInt(d_1876_v11_, d_1876_v11_)
                nw321_[int(20)] = (d_1876_v11_) * (len(_dafny.Map({(d_1906_v25_).f19: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jrtnqutc"))})))
                nw321_[int(21)] = (d_1876_v11_) + (d_1876_v11_)
                nw321_[int(22)] = (d_1881_v14_)[default__.safeIndex(d_1876_v11_, len(d_1881_v14_))]
                nw321_[int(23)] = default__.safeModuloInt(d_1876_v11_, d_1876_v11_)
                d_1913_v33_ = nw321_
                def lambda115_(d_1915_v11_):
                    def lambda116_(d_1916_i8_):
                        return default__.safeModuloInt(d_1916_i8_, d_1915_v11_)

                    return lambda116_

                init62_ = lambda115_(d_1876_v11_)
                nw322_ = _dafny.Array(None, 26)
                for i0_62_ in range(nw322_.length(0)):
                    nw322_[i0_62_] = init62_(i0_62_)
                d_1913_v33_ = nw322_
                d_1917_v34_: C0
                nw323_ = C0()
                nw323_.ctor__((d_1877_v12_)[default__.safeIndex(192, (d_1877_v12_).length(0))])
                d_1917_v34_ = nw323_
        elif True:
            d_1918_v35_: _dafny.Seq
            d_1918_v35_ = _dafny.SeqWithoutIsStrInference([(self).f11])
            d_1919_v36_: C4
            nw324_ = C4()
            nw324_.ctor__((d_1918_v35_)[default__.safeIndex(d_1876_v11_, len(d_1918_v35_))])
            d_1919_v36_ = nw324_
            default__.m0(globalState)
            (globalState).f8 = not(d_1860_v0_)
            (globalState).f9 = 385
            (globalState).f9 = (0) - ((default__.fm0(False, globalState)) + (d_1876_v11_))
        d_1920_v37_: _dafny.MultiSet
        d_1920_v37_ = _dafny.MultiSet([d_1876_v11_])
        d_1921_v38_: _dafny.Seq
        d_1921_v38_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vmba"))
        d_1922_v39_: str
        d_1922_v39_ = _dafny.CodePoint('j')
        d_1923_v40_: _dafny.Set
        d_1923_v40_ = _dafny.Set({len((d_1921_v38_).set(default__.safeIndex(d_1876_v11_, len(d_1921_v38_)), d_1922_v39_))})
        d_1920_v37_ = (d_1920_v37_).set((len(d_1923_v40_)) * (d_1876_v11_), default__.abs((d_1876_v11_) * ((0) - (d_1876_v11_))))
        r0 = default__.fm3(((_dafny.MultiSet([d_1876_v11_])) - (_dafny.MultiSet([d_1876_v11_, -496]))).cardinality, (d_1860_v0_) or (not(d_1860_v0_)), globalState)
        d_1924_v41_: _dafny.Map
        d_1924_v41_ = _dafny.Map({d_1876_v11_: d_1860_v0_})
        r1 = not (((d_1924_v41_)[default__.fm0(not(not(d_1860_v0_)), globalState)] if (default__.fm0(not(not(d_1860_v0_)), globalState)) in (d_1924_v41_) else d_1860_v0_)) or ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "f"))) != (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnwss"))))
        r2 = d_1860_v0_
        return r0, r1, r2


class C6:
    def  __init__(self):
        self.f16: _dafny.Seq = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, ""))
        self._f15: int = int(0)
        pass

    def __dafnystr__(self) -> str:
        return "_module.C6"
    def ctor__(self, f15, f16):
        (self)._f15 = f15
        (self).f16 = f16

    def m6(self, p0, globalState):
        r0: str = _dafny.CodePoint('D')
        d_1925_v0_: _dafny.Map
        d_1925_v0_ = _dafny.Map({p0: (self).f15})
        hi14_ = ((d_1925_v0_)[False] if (False) in (d_1925_v0_) else len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('d') for d_1926_i1_ in range(default__.abs(647))])))
        for d_1927_i0_ in range((self).f15, hi14_):
            d_1928_v1_: _dafny.Seq
            d_1928_v1_ = _dafny.SeqWithoutIsStrInference([p0, p0])
            d_1929_v2_: _dafny.Array
            nw325_ = _dafny.Array(None, 3)
            nw325_[int(0)] = (d_1928_v1_)[default__.safeIndex(d_1927_i0_, len(d_1928_v1_))]
            nw325_[int(1)] = p0
            nw325_[int(2)] = p0
            d_1929_v2_ = nw325_
            d_1930_v3_: _dafny.Set
            d_1930_v3_ = _dafny.Set({d_1929_v2_})
            rhs344_ = (d_1930_v3_) - ((d_1930_v3_ if p0 else d_1930_v3_))
            rhs345_ = (self).f15
            rhs346_ = (0) - (default__.safeDivisionInt(d_1927_i0_, len(self.f16)))
            lhs250_ = globalState
            lhs251_ = globalState
            d_1930_v3_ = rhs344_
            lhs250_.f9 = rhs345_
            lhs251_.f9 = rhs346_
            d_1931_v4_: _dafny.Map
            d_1931_v4_ = _dafny.Map({p0: self.f16})
            d_1931_v4_ = (d_1931_v4_).set(p0, self.f16)
            d_1932_v5_: _dafny.MultiSet
            d_1932_v5_ = _dafny.MultiSet([p0])
            d_1933_v6_: D0
            d_1933_v6_ = D0_DC2(d_1932_v5_, p0, (self).f15, p0)
            d_1934_v7_: C0
            nw326_ = C0()
            nw326_.ctor__((d_1933_v6_).cf7)
            d_1934_v7_ = nw326_
            d_1935_v8_: str
            d_1935_v8_ = _dafny.CodePoint('u')
            d_1936_v9_: _dafny.Array
            nw327_ = _dafny.Array(None, 19)
            nw327_[int(0)] = d_1935_v8_
            nw327_[int(1)] = d_1935_v8_
            nw327_[int(2)] = d_1935_v8_
            nw327_[int(3)] = default__.fm3(d_1927_i0_, p0, globalState)
            nw327_[int(4)] = d_1935_v8_
            nw327_[int(5)] = (D1_DC4((self.f16)[default__.safeIndex(len(d_1928_v1_), len(self.f16))])).cf9
            nw327_[int(6)] = d_1935_v8_
            nw327_[int(7)] = d_1935_v8_
            nw327_[int(8)] = d_1935_v8_
            nw327_[int(9)] = d_1935_v8_
            nw327_[int(10)] = d_1935_v8_
            nw327_[int(11)] = d_1935_v8_
            nw327_[int(12)] = d_1935_v8_
            nw327_[int(13)] = d_1935_v8_
            nw327_[int(14)] = d_1935_v8_
            nw327_[int(15)] = d_1935_v8_
            nw327_[int(16)] = d_1935_v8_
            nw327_[int(17)] = d_1935_v8_
            nw327_[int(18)] = _dafny.CodePoint('e')
            d_1936_v9_ = nw327_
            d_1937_v10_: _dafny.MultiSet
            d_1937_v10_ = _dafny.MultiSet([d_1936_v9_, d_1936_v9_])
            d_1938_v11_: _dafny.Array
            nw328_ = _dafny.Array(None, 1)
            nw328_[int(0)] = d_1937_v10_
            d_1938_v11_ = nw328_
            index304_ = default__.safeIndex(722, (d_1938_v11_).length(0))
            (d_1938_v11_)[index304_] = _dafny.MultiSet([d_1936_v9_, d_1936_v9_])
            d_1939_v12_: _dafny.Seq
            d_1939_v12_ = _dafny.SeqWithoutIsStrInference([d_1937_v10_])
            index305_ = default__.safeIndex(722, (d_1938_v11_).length(0))
            (d_1938_v11_)[index305_] = ((d_1939_v12_)[default__.safeIndex(464, len(d_1939_v12_))]).intersection(d_1937_v10_)
        d_1940_v13_: _dafny.Set
        d_1940_v13_ = _dafny.Set({default__.fm0(p0, globalState)})
        d_1941_v14_: _dafny.MultiSet
        d_1941_v14_ = _dafny.MultiSet([not(p0)])
        d_1942_v15_: _dafny.Array
        nw329_ = _dafny.Array(None, 23)
        nw329_[int(0)] = True
        nw329_[int(1)] = (p0 if p0 else not(p0))
        nw329_[int(2)] = p0
        nw329_[int(3)] = p0
        nw329_[int(4)] = p0
        nw329_[int(5)] = ((self).f15) == (len(d_1925_v0_))
        nw329_[int(6)] = ((self).f15) > (default__.fm0(True, globalState))
        nw329_[int(7)] = p0
        nw329_[int(8)] = p0
        nw329_[int(9)] = p0
        nw329_[int(10)] = p0
        nw329_[int(11)] = (d_1940_v13_).ispropersubset(d_1940_v13_)
        nw329_[int(12)] = p0
        nw329_[int(13)] = p0
        nw329_[int(14)] = p0
        nw329_[int(15)] = p0
        nw329_[int(16)] = p0
        nw329_[int(17)] = p0
        nw329_[int(18)] = False
        nw329_[int(19)] = not(p0)
        nw329_[int(20)] = p0
        nw329_[int(21)] = p0
        nw329_[int(22)] = ((d_1941_v14_).set(not(p0), default__.abs(len(self.f16)))).issubset(d_1941_v14_)
        d_1942_v15_ = nw329_
        index306_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        (d_1942_v15_)[index306_] = p0
        d_1943_v16_: _dafny.Map
        d_1943_v16_ = _dafny.Map({p0: p0})
        index307_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        (d_1942_v15_)[index307_] = ((d_1943_v16_)[p0] if (p0) in (d_1943_v16_) else (_dafny.Set({(self).f15, (self).f15})).ispropersubset(d_1940_v13_))
        d_1944_v17_: _dafny.Array
        nw330_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 13)
        d_1944_v17_ = nw330_
        d_1945_v18_: _dafny.Map
        d_1945_v18_ = _dafny.Map({d_1944_v17_: len(self.f16)})
        d_1945_v18_ = (d_1945_v18_).set(d_1944_v17_, default__.safeDivisionInt(-627, -656))
        d_1946_v19_: _dafny.Array
        nw331_ = _dafny.Array(int(0), 20)
        d_1946_v19_ = nw331_
        index308_ = default__.safeIndex(133, (d_1946_v19_).length(0))
        (d_1946_v19_)[index308_] = ((self).f15) - ((self).f15)
        d_1947_v20_: _dafny.Seq
        d_1947_v20_ = _dafny.SeqWithoutIsStrInference([(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]])
        d_1948_v21_: _dafny.Map
        d_1948_v21_ = _dafny.Map({(self).f15: (d_1947_v20_)[default__.safeIndex((self).f15, len(d_1947_v20_))]})
        index309_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        index310_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        index311_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        index312_ = default__.safeIndex(133, (d_1946_v19_).length(0))
        rhs347_ = ((d_1948_v21_)[default__.fm0((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], globalState)] if (default__.fm0((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], globalState)) in (d_1948_v21_) else (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))])
        rhs348_ = (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]
        rhs349_ = (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]
        rhs350_ = (self).f15
        lhs252_ = d_1942_v15_
        lhs253_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        lhs254_ = d_1942_v15_
        lhs255_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        lhs256_ = d_1942_v15_
        lhs257_ = default__.safeIndex(921, (d_1942_v15_).length(0))
        lhs258_ = d_1946_v19_
        lhs259_ = default__.safeIndex(133, (d_1946_v19_).length(0))
        lhs252_[lhs253_] = rhs347_
        lhs254_[lhs255_] = rhs348_
        lhs256_[lhs257_] = rhs349_
        lhs258_[lhs259_] = rhs350_
        d_1949_v22_: _dafny.Set
        d_1949_v22_ = _dafny.Set({(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]})
        d_1949_v22_ = d_1949_v22_
        if not((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]):
            d_1950_v23_: _dafny.Map
            d_1950_v23_ = _dafny.Map({(d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]: d_1940_v13_})
            d_1951_v24_: _dafny.Map
            d_1951_v24_ = _dafny.Map({(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]: d_1940_v13_})
            d_1952_v25_: _dafny.MultiSet
            d_1952_v25_ = _dafny.MultiSet([((d_1950_v23_)[len(self.f16)] if (len(self.f16)) in (d_1950_v23_) else ((d_1951_v24_)[(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]] if ((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]) in (d_1951_v24_) else d_1940_v13_))])
            d_1953_v26_: _dafny.Map
            d_1953_v26_ = _dafny.Map({p0: d_1952_v25_})
            index313_ = default__.safeIndex(921, (d_1942_v15_).length(0))
            (d_1942_v15_)[index313_] = (((d_1953_v26_)[(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]] if ((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]) in (d_1953_v26_) else d_1952_v25_)) != ((d_1952_v25_ if not(p0) else default__.fm14(d_1947_v20_, (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], globalState)))
            d_1954_v27_: _dafny.Seq
            d_1954_v27_ = _dafny.SeqWithoutIsStrInference([d_1943_v16_])
            d_1955_v28_: _dafny.Map
            d_1955_v28_ = _dafny.Map({d_1946_v19_: d_1954_v27_})
            d_1956_v29_: _dafny.Array
            nw332_ = _dafny.Array(None, 13)
            nw332_[int(0)] = (d_1954_v27_ if (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))] else d_1954_v27_)
            nw332_[int(1)] = (d_1954_v27_) + (d_1954_v27_)
            nw332_[int(2)] = d_1954_v27_
            nw332_[int(3)] = _dafny.SeqWithoutIsStrInference([d_1943_v16_])
            nw332_[int(4)] = _dafny.SeqWithoutIsStrInference([d_1943_v16_ for d_1957_i2_ in range(default__.abs(-892))])
            nw332_[int(5)] = d_1954_v27_
            nw332_[int(6)] = ((d_1955_v28_)[d_1946_v19_] if (d_1946_v19_) in (d_1955_v28_) else d_1954_v27_)
            nw332_[int(7)] = d_1954_v27_
            nw332_[int(8)] = default__.fm15((d_1947_v20_)[default__.safeIndex(len(self.f16), len(d_1947_v20_))], globalState)
            nw332_[int(9)] = (d_1954_v27_) + (d_1954_v27_)
            nw332_[int(10)] = _dafny.SeqWithoutIsStrInference([d_1943_v16_, d_1943_v16_])
            nw332_[int(11)] = d_1954_v27_
            nw332_[int(12)] = (d_1954_v27_ if True else d_1954_v27_)
            d_1956_v29_ = nw332_
            index314_ = default__.safeIndex(313, (d_1956_v29_).length(0))
            (d_1956_v29_)[index314_] = _dafny.SeqWithoutIsStrInference([_dafny.Map({(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]: p0}) for d_1958_i3_ in range(default__.abs(-742))])
            index315_ = default__.safeIndex(313, (d_1956_v29_).length(0))
            (d_1956_v29_)[index315_] = (d_1954_v27_) + (_dafny.SeqWithoutIsStrInference([d_1943_v16_ for d_1959_i4_ in range(default__.abs(960))]))
            if p0:
                d_1960_v30_: D2
                d_1960_v30_ = D2_DC7(p0, (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], p0, _dafny.SeqWithoutIsStrInference([_dafny.CodePoint('f') for d_1961_i5_ in range(default__.abs(165))]))
                d_1962_v31_: str
                d_1962_v31_ = _dafny.CodePoint('u')
                index316_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                rhs351_ = ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))] if not((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]) else len(self.f16))
                rhs352_ = len(self.f16)
                rhs353_ = ((default__.fm1((d_1960_v30_).cf13, d_1962_v31_, globalState)) or (False) if not (True) or (p0) else p0)
                lhs260_ = globalState
                lhs261_ = globalState
                lhs262_ = d_1942_v15_
                lhs263_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                lhs260_.f9 = rhs351_
                lhs261_.f9 = rhs352_
                lhs262_[lhs263_] = rhs353_
                d_1963_v33_: _dafny.Map
                d_1963_v33_ = _dafny.Map({len(self.f16): default__.fm0((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], globalState)})
                d_1964_v34_: _dafny.Array
                nw333_ = _dafny.Array(None, 8)
                nw333_[int(0)] = d_1948_v21_
                nw333_[int(1)] = d_1948_v21_
                nw333_[int(2)] = d_1948_v21_
                nw333_[int(3)] = _dafny.Map({(self).f15: p0})
                nw333_[int(4)] = (d_1948_v21_).set(default__.fm0((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], globalState), (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))])
                nw333_[int(5)] = d_1948_v21_
                def iife128_():
                    coll66_ = _dafny.Map()
                    compr_66_: int
                    for compr_66_ in (d_1963_v33_).keys.Elements:
                        d_1965_v32_: int = compr_66_
                        if (d_1965_v32_) in (d_1963_v33_):
                            coll66_[default__.safeDivisionInt(d_1965_v32_, (self).f15)] = p0
                    return _dafny.Map(coll66_)
                nw333_[int(6)] = iife128_()
                
                nw333_[int(7)] = d_1948_v21_
                d_1964_v34_ = nw333_
                d_1966_v35_: T1
                nw334_ = C2()
                nw334_.ctor__(d_1964_v34_)
                d_1966_v35_ = nw334_
                d_1966_v35_ = d_1966_v35_
                index317_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                (d_1942_v15_)[index317_] = (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]
                index318_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                (d_1946_v19_)[index318_] = (0) - ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))])
                d_1967_v36_: _dafny.MultiSet
                d_1967_v36_ = _dafny.MultiSet([d_1942_v15_, d_1942_v15_])
                d_1949_v22_ = _dafny.Set({((d_1948_v21_)[(self).f15] if ((self).f15) in (d_1948_v21_) else not(p0)), p0, (d_1967_v36_).ispropersubset(d_1967_v36_)})
            elif True:
                d_1968_v37_: str
                d_1968_v37_ = _dafny.CodePoint('n')
                (globalState).f8 = default__.fm1((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], (d_1968_v37_ if (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))] else d_1968_v37_), globalState)
                d_1969_v38_: _dafny.MultiSet
                d_1969_v38_ = _dafny.MultiSet([(self).f15, (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]])
                d_1970_v39_: D2
                d_1970_v39_ = D2_DC7(p0, (d_1969_v38_).cardinality, False, self.f16)
                d_1971_v40_: _dafny.MultiSet
                d_1971_v40_ = _dafny.MultiSet([d_1968_v37_])
                d_1972_v41_: _dafny.Map
                d_1972_v41_ = _dafny.Map({(d_1971_v40_ if (d_1970_v39_).cf13 else d_1971_v40_): len(d_1940_v13_)})
                d_1972_v41_ = (d_1972_v41_).set((d_1971_v40_).set(d_1968_v37_, default__.abs(len(_dafny.SeqWithoutIsStrInference([d_1968_v37_ for d_1973_i6_ in range(default__.abs(120))])))), (0) - ((self).f15))
                index319_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                (d_1946_v19_)[index319_] = (self).f15
                d_1974_v42_: D0
                d_1974_v42_ = D0_DC2((d_1941_v14_).set((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], default__.abs((self).f15)), ((d_1948_v21_)[(self).f15] if ((self).f15) in (d_1948_v21_) else (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]), (self).f15, (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))])
                d_1975_v43_: _dafny.Seq
                d_1975_v43_ = _dafny.SeqWithoutIsStrInference([default__.fm33(d_1943_v16_, globalState), (d_1974_v42_).cf4, default__.fm33(d_1943_v16_, globalState)])
                d_1941_v14_ = (d_1975_v43_)[default__.safeIndex((self).f15, len(d_1975_v43_))]
                (self).f16 = self.f16
            if (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]:
                (globalState).f9 = len((self.f16 if (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))] else self.f16))
                d_1976_v44_: _dafny.Seq
                d_1976_v44_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([not((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]), (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]])])
                d_1977_v45_: _dafny.Seq
                d_1977_v45_ = _dafny.SeqWithoutIsStrInference([d_1976_v44_])
                d_1978_v46_: _dafny.Map
                d_1978_v46_ = _dafny.Map({p0: self.f16})
                d_1979_v47_: _dafny.Array
                nw335_ = _dafny.Array(None, 16)
                nw335_[int(0)] = d_1948_v21_
                nw335_[int(1)] = (d_1948_v21_).set((self).f15, (d_1947_v20_)[default__.safeIndex((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], len(d_1947_v20_))])
                nw335_[int(2)] = d_1948_v21_
                nw335_[int(3)] = d_1948_v21_
                nw335_[int(4)] = default__.fm27((d_1977_v45_)[default__.safeIndex((self).f15, len(d_1977_v45_))], globalState)
                nw335_[int(5)] = d_1948_v21_
                nw335_[int(6)] = d_1948_v21_
                nw335_[int(7)] = d_1948_v21_
                nw335_[int(8)] = _dafny.Map({default__.fm0(True, globalState): (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]})
                nw335_[int(9)] = d_1948_v21_
                nw335_[int(10)] = (d_1948_v21_).set(len(((d_1978_v46_)[(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]] if ((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]) in (d_1978_v46_) else self.f16)), not((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]))
                nw335_[int(11)] = d_1948_v21_
                nw335_[int(12)] = _dafny.Map({(self).f15: False})
                nw335_[int(13)] = d_1948_v21_
                nw335_[int(14)] = d_1948_v21_
                nw335_[int(15)] = default__.fm27(d_1976_v44_, globalState)
                d_1979_v47_ = nw335_
                d_1980_v48_: C4
                nw336_ = C4()
                nw336_.ctor__(d_1979_v47_)
                d_1980_v48_ = nw336_
                index320_ = default__.safeIndex(878, (d_1944_v17_).length(0))
                (d_1944_v17_)[index320_] = self.f16
                d_1981_v49_: str
                d_1981_v49_ = _dafny.CodePoint('c')
                d_1982_v50_: _dafny.Set
                d_1982_v50_ = _dafny.Set({self.f16})
                d_1983_v51_: _dafny.Map
                d_1983_v51_ = _dafny.Map({(self).f15: d_1982_v50_})
                index321_ = default__.safeIndex(878, (d_1944_v17_).length(0))
                rhs354_ = len((_dafny.Set({self.f16, (self.f16).set(default__.safeIndex((0) - ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]), len(self.f16)), d_1981_v49_), _dafny.SeqWithoutIsStrInference([d_1981_v49_ for d_1984_i7_ in range(default__.abs(720))])})).intersection(((d_1983_v51_)[(d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]] if ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]) in (d_1983_v51_) else _dafny.Set({self.f16, self.f16}))))
                rhs355_ = d_1980_v48_
                rhs356_ = self.f16
                lhs264_ = globalState
                lhs265_ = d_1944_v17_
                lhs266_ = default__.safeIndex(878, (d_1944_v17_).length(0))
                lhs264_.f9 = rhs354_
                d_1980_v48_ = rhs355_
                lhs265_[lhs266_] = rhs356_
                rhs357_ = default__.safeModuloInt(785, (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))])
                rhs358_ = default__.fm33(d_1943_v16_, globalState)
                lhs267_ = globalState
                lhs267_.f9 = rhs357_
                d_1941_v14_ = rhs358_
                d_1985_v52_: _dafny.Seq
                d_1985_v52_ = _dafny.SeqWithoutIsStrInference([(d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]])
                d_1986_v53_: D5
                d_1986_v53_ = D5_DC13(d_1985_v52_)
                d_1987_v54_: _dafny.Map
                d_1987_v54_ = _dafny.Map({(d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]: default__.fm0((d_1980_v48_).fm42((self).f15, len((d_1944_v17_)[default__.safeIndex(878, (d_1944_v17_).length(0))]), globalState), globalState)})
                index322_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                index323_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                index324_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                rhs359_ = d_1986_v53_
                rhs360_ = ((len(d_1948_v21_)) + ((self).f15)) != (len(d_1947_v20_))
                rhs361_ = (self).f15
                rhs362_ = len(((d_1987_v54_) | (d_1987_v54_)) | (d_1987_v54_))
                lhs268_ = d_1942_v15_
                lhs269_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                lhs270_ = d_1946_v19_
                lhs271_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                lhs272_ = d_1946_v19_
                lhs273_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                d_1986_v53_ = rhs359_
                lhs268_[lhs269_] = rhs360_
                lhs270_[lhs271_] = rhs361_
                lhs272_[lhs273_] = rhs362_
                d_1949_v22_ = d_1949_v22_
            elif True:
                index325_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                (d_1942_v15_)[index325_] = (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]
                index326_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                (d_1946_v19_)[index326_] = ((self).f15) - (len(_dafny.Set({d_1940_v13_})))
                index327_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                (d_1942_v15_)[index327_] = (d_1947_v20_)[default__.safeIndex(default__.fm0((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], globalState), len(d_1947_v20_))]
                d_1988_v55_: _dafny.Array
                def lambda117_(d_1989_v14_, d_1990_v20_):
                    def lambda118_(d_1991_i8_):
                        return (d_1989_v14_) | (_dafny.MultiSet(d_1990_v20_))

                    return lambda118_

                init63_ = lambda117_(d_1941_v14_, d_1947_v20_)
                nw337_ = _dafny.Array(None, 23)
                for i0_63_ in range(nw337_.length(0)):
                    nw337_[i0_63_] = init63_(i0_63_)
                d_1988_v55_ = nw337_
                d_1992_v56_: _dafny.Map
                d_1992_v56_ = _dafny.Map({p0: _dafny.MultiSet([(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]])})
                index328_ = default__.safeIndex(73, (d_1988_v55_).length(0))
                (d_1988_v55_)[index328_] = ((d_1992_v56_)[not(p0)] if (not(p0)) in (d_1992_v56_) else _dafny.MultiSet([p0]))
                index329_ = default__.safeIndex(73, (d_1988_v55_).length(0))
                index330_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                index331_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                rhs363_ = ((d_1941_v14_) - (d_1941_v14_)) | ((_dafny.MultiSet([p0])).intersection(d_1941_v14_))
                rhs364_ = (self).f15
                rhs365_ = (self).f15
                rhs366_ = len(d_1940_v13_)
                lhs274_ = d_1988_v55_
                lhs275_ = default__.safeIndex(73, (d_1988_v55_).length(0))
                lhs276_ = d_1946_v19_
                lhs277_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                lhs278_ = globalState
                lhs279_ = d_1946_v19_
                lhs280_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                lhs274_[lhs275_] = rhs363_
                lhs276_[lhs277_] = rhs364_
                lhs278_.f9 = rhs365_
                lhs279_[lhs280_] = rhs366_
                d_1993_v57_: _dafny.MultiSet
                d_1993_v57_ = _dafny.MultiSet([(self).f15])
                d_1994_v58_: D0
                d_1994_v58_ = D0_DC0((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))])
                d_1995_v59_: C1
                nw338_ = C1()
                nw338_.ctor__((d_1993_v57_).isdisjoint(_dafny.MultiSet([((d_1941_v14_)[(d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]] if ((d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]) in (d_1941_v14_) else (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))])])), (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))], d_1994_v58_, (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))])
                d_1995_v59_ = nw338_
            d_1996_v60_: _dafny.Array
            def lambda119_(d_1997_v21_):
                def lambda120_(d_1998_i9_):
                    return d_1997_v21_

                return lambda120_

            init64_ = lambda119_(d_1948_v21_)
            nw339_ = _dafny.Array(None, 16)
            for i0_64_ in range(nw339_.length(0)):
                nw339_[i0_64_] = init64_(i0_64_)
            d_1996_v60_ = nw339_
            d_1999_v61_: C4
            nw340_ = C4()
            nw340_.ctor__(d_1996_v60_)
            d_1999_v61_ = nw340_
            d_2000_v62_: _dafny.Map
            d_2000_v62_ = _dafny.Map({p0: self.f16})
            d_2001_v63_: _dafny.Map
            d_2001_v63_ = _dafny.Map({d_1999_v61_: len(d_2000_v62_)})
            (globalState).f9 = len((d_2001_v63_ if not(p0) else d_2001_v63_))
        elif True:
            d_2002_v64_: C0
            nw341_ = C0()
            nw341_.ctor__(((0) - (len(self.f16))) <= ((self).f15))
            d_2002_v64_ = nw341_
            d_2003_v65_: D0
            d_2003_v65_ = D0_DC0(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('s') for d_2004_i10_ in range(default__.abs(289))])))
            d_2005_v66_: C1
            nw342_ = C1()
            nw342_.ctor__(not (p0) or (False), p0, d_2003_v65_, ((d_1947_v20_)[default__.safeIndex((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], len(d_1947_v20_))]) == (not((d_2002_v64_).f17)))
            d_2005_v66_ = nw342_
            if (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]:
                (globalState).f8 = (d_2005_v66_).f19
                d_2006_v67_: str
                d_2006_v67_ = _dafny.CodePoint('b')
                d_2007_v68_: _dafny.Map
                d_2007_v68_ = _dafny.Map({d_2006_v67_: (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]})
                d_2008_v69_: _dafny.Map
                d_2008_v69_ = _dafny.Map({(_dafny.SeqWithoutIsStrInference([len(self.f16), (self).f15, len(default__.fm36((d_2005_v66_).f19, globalState)), (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]])).set(default__.safeIndex(((d_2007_v68_)[d_2006_v67_] if (d_2006_v67_) in (d_2007_v68_) else (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]), len(_dafny.SeqWithoutIsStrInference([len(self.f16), (self).f15, len(default__.fm36((d_2005_v66_).f19, globalState)), (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]]))), (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]): ((d_2005_v66_).f20) not in (d_1941_v14_)})
                d_2009_v70_: _dafny.Seq
                d_2009_v70_ = _dafny.SeqWithoutIsStrInference([(self).f15])
                rhs367_ = d_1942_v15_
                rhs368_ = False
                rhs369_ = (_dafny.Map({d_2009_v70_: True})) | (d_2008_v69_)
                lhs281_ = globalState
                d_1942_v15_ = rhs367_
                lhs281_.f8 = rhs368_
                d_2008_v69_ = rhs369_
                d_2010_v71_: _dafny.Map
                d_2010_v71_ = _dafny.Map({self.f16: (d_2005_v66_).f20})
                d_2010_v71_ = (d_2010_v71_).set(self.f16, (d_1941_v14_).issubset(d_1941_v14_))
                index332_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                (d_1946_v19_)[index332_] = (0) - (default__.safeModuloInt((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], (0) - (default__.fm0((d_2005_v66_).f19, globalState))))
                (globalState).f9 = (0) - ((self).f15)
            elif True:
                d_1941_v14_ = (d_1941_v14_) | (d_1941_v14_)
                d_2011_v72_: str
                d_2011_v72_ = _dafny.CodePoint('c')
                d_2012_v73_: D5
                d_2012_v73_ = D5_DC14((d_2005_v66_).f19)
                d_2013_v74_: _dafny.Map
                d_2013_v74_ = _dafny.Map({default__.fm58(d_2011_v72_, (d_2012_v73_).cf26, globalState): p0})
                d_2014_v75_: _dafny.MultiSet
                d_2014_v75_ = _dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f15 for d_2015_i11_ in range(default__.abs(-208))]))
                d_2016_v76_: _dafny.Seq
                d_2016_v76_ = _dafny.SeqWithoutIsStrInference([len(self.f16)])
                d_2013_v74_ = (d_2013_v74_).set(_dafny.MultiSet(d_2016_v76_), (d_2005_v66_).f19)
                d_2017_v77_: _dafny.Array
                def lambda121_(d_2018_v20_, d_2019_v19_, d_2020_v21_, d_2021_v64_, d_2022_v66_):
                    def lambda122_(d_2023_i12_):
                        return ((d_2018_v20_).set(default__.safeIndex((d_2019_v19_)[default__.safeIndex(133, (d_2019_v19_).length(0))], len(d_2018_v20_)), ((d_2020_v21_)[(self).f15] if ((self).f15) in (d_2020_v21_) else (d_2021_v64_).f17))) + (_dafny.SeqWithoutIsStrInference([(d_2022_v66_).f19, (d_2021_v64_).f17, (d_2021_v64_).f17]))

                    return lambda122_

                init65_ = lambda121_(d_1947_v20_, d_1946_v19_, d_1948_v21_, d_2002_v64_, d_2005_v66_)
                nw343_ = _dafny.Array(None, 7)
                for i0_65_ in range(nw343_.length(0)):
                    nw343_[i0_65_] = init65_(i0_65_)
                d_2017_v77_ = nw343_
                d_2024_v78_: _dafny.MultiSet
                d_2024_v78_ = _dafny.MultiSet([(self).f15])
                d_2025_v80_: _dafny.Map
                d_2025_v80_ = _dafny.Map({(d_1941_v14_).cardinality: d_2024_v78_})
                d_2026_v81_: _dafny.Map
                d_2026_v81_ = _dafny.Map({d_2011_v72_: d_2024_v78_})
                d_2027_v82_: _dafny.Array
                nw344_ = _dafny.Array(None, 18)
                nw344_[int(0)] = d_2024_v78_
                nw344_[int(1)] = (d_2024_v78_) - (d_2024_v78_)
                nw344_[int(2)] = d_2024_v78_
                nw344_[int(3)] = d_2024_v78_
                nw344_[int(4)] = d_2024_v78_
                nw344_[int(5)] = (d_2024_v78_).set(len(_dafny.SeqWithoutIsStrInference([(d_2002_v64_).f17, (d_2005_v66_).f19])), default__.abs(len(self.f16)))
                def iife129_():
                    coll67_ = _dafny.Map()
                    compr_67_: int
                    for compr_67_ in _dafny.IntegerRange(-154, 309):
                        d_2028_v79_: int = compr_67_
                        if ((-154) <= (d_2028_v79_)) and ((d_2028_v79_) < (309)):
                            coll67_[default__.safeModuloInt(d_2028_v79_, (self).f15)] = (d_1942_v15_)[default__.safeIndex(921, (d_1942_v15_).length(0))]
                    return _dafny.Map(coll67_)
                nw344_[int(6)] = (_dafny.MultiSet([(self).f15])) - (_dafny.MultiSet([(self).f15, (d_2005_v66_).fm30(iife129_()
                , globalState), (self).f15]))
                nw344_[int(7)] = default__.fm38((d_2005_v66_).f19, (d_2005_v66_).f19, d_2011_v72_, (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], globalState)
                nw344_[int(8)] = d_2024_v78_
                nw344_[int(9)] = ((d_2025_v80_)[979] if (979) in (d_2025_v80_) else d_2024_v78_)
                nw344_[int(10)] = ((d_2026_v81_)[d_2011_v72_] if (d_2011_v72_) in (d_2026_v81_) else _dafny.MultiSet([(self).f15, (self).f15]))
                nw344_[int(11)] = (_dafny.MultiSet(d_2016_v76_)).intersection(_dafny.MultiSet(d_2016_v76_))
                nw344_[int(12)] = d_2024_v78_
                nw344_[int(13)] = _dafny.MultiSet([(d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], len(d_1940_v13_)])
                nw344_[int(14)] = d_2024_v78_
                nw344_[int(15)] = d_2024_v78_
                nw344_[int(16)] = d_2024_v78_
                nw344_[int(17)] = d_2024_v78_
                d_2027_v82_ = nw344_
                index333_ = default__.safeIndex(977, (d_2027_v82_).length(0))
                (d_2027_v82_)[index333_] = d_2024_v78_
                index334_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                index335_ = default__.safeIndex(977, (d_2027_v82_).length(0))
                rhs370_ = self.f16
                rhs371_ = d_2017_v77_
                rhs372_ = ((self).f15) < ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))])
                rhs373_ = ((self).f15 if (d_2005_v66_).f20 else (self).f15)
                rhs374_ = ((_dafny.MultiSet([(self).f15, (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]])) | (d_2024_v78_)) | (default__.fm56((d_2005_v66_).f19, (self).f15, globalState))
                lhs282_ = self
                lhs283_ = d_1942_v15_
                lhs284_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                lhs285_ = globalState
                lhs286_ = d_2027_v82_
                lhs287_ = default__.safeIndex(977, (d_2027_v82_).length(0))
                lhs282_.f16 = rhs370_
                d_2017_v77_ = rhs371_
                lhs283_[lhs284_] = rhs372_
                lhs285_.f9 = rhs373_
                lhs286_[lhs287_] = rhs374_
                d_2029_v83_: _dafny.Array
                nw345_ = _dafny.Array(D2.default()(), 12)
                d_2029_v83_ = nw345_
                d_2030_v84_: _dafny.Map
                d_2030_v84_ = _dafny.Map({(d_2002_v64_).f17: (d_2027_v82_)[default__.safeIndex(977, (d_2027_v82_).length(0))]})
                index336_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                rhs375_ = not(((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]) < ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]))
                rhs376_ = d_2029_v83_
                rhs377_ = ((d_2030_v84_)[(d_2002_v64_).f17] if ((d_2002_v64_).f17) in (d_2030_v84_) else ((d_2027_v82_)[default__.safeIndex(977, (d_2027_v82_).length(0))]).set(719, default__.abs(len(self.f16))))
                lhs288_ = d_1942_v15_
                lhs289_ = default__.safeIndex(921, (d_1942_v15_).length(0))
                lhs288_[lhs289_] = rhs375_
                d_2029_v83_ = rhs376_
                d_2024_v78_ = rhs377_
                d_2031_v87_: _dafny.Array
                def lambda123_(d_2032_v76_, d_2033_v66_):
                    def lambda124_(d_2034_i13_):
                        def iife130_():
                            coll68_ = _dafny.Map()
                            compr_68_: int
                            for compr_68_ in (_dafny.MultiSet(d_2032_v76_)).Elements:
                                d_2035_v86_: int = compr_68_
                                if (d_2035_v86_) in (_dafny.MultiSet(d_2032_v76_)):
                                    coll68_[(d_2035_v86_) - ((self).f15)] = (d_2033_v66_).f19
                            return _dafny.Map(coll68_)
                        return iife130_()
                        

                    return lambda124_

                init66_ = lambda123_(d_2016_v76_, d_2005_v66_)
                nw346_ = _dafny.Array(None, 26)
                for i0_66_ in range(nw346_.length(0)):
                    nw346_[i0_66_] = init66_(i0_66_)
                d_2031_v87_ = nw346_
                d_2036_v88_: T1
                nw347_ = C5()
                nw347_.ctor__(d_2031_v87_)
                d_2036_v88_ = nw347_
                d_2037_v89_: _dafny.MultiSet
                d_2037_v89_ = _dafny.MultiSet([d_2036_v88_, d_2036_v88_])
                d_2038_v90_: _dafny.Array
                nw348_ = _dafny.Array(None, 22)
                nw348_[int(0)] = d_1948_v21_
                nw348_[int(1)] = _dafny.Map({(self).f15: (d_2005_v66_).f19})
                nw348_[int(2)] = d_1948_v21_
                def iife131_():
                    coll69_ = _dafny.Map()
                    compr_69_: int
                    for compr_69_ in _dafny.IntegerRange(330, 75):
                        d_2039_v85_: int = compr_69_
                        if ((330) <= (d_2039_v85_)) and ((d_2039_v85_) < (75)):
                            coll69_[default__.safeModuloInt(d_2039_v85_, len((D13_DC29(d_1948_v21_)).cf50))] = (d_2002_v64_).f17
                    return _dafny.Map(coll69_)
                nw348_[int(3)] = iife131_()
                
                nw348_[int(4)] = d_1948_v21_
                nw348_[int(5)] = d_1948_v21_
                nw348_[int(6)] = d_1948_v21_
                nw348_[int(7)] = d_1948_v21_
                nw348_[int(8)] = _dafny.Map({(d_2005_v66_).fm29((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))], globalState): (d_2005_v66_).f20})
                nw348_[int(9)] = d_1948_v21_
                nw348_[int(10)] = d_1948_v21_
                nw348_[int(11)] = d_1948_v21_
                nw348_[int(12)] = d_1948_v21_
                nw348_[int(13)] = d_1948_v21_
                nw348_[int(14)] = _dafny.Map({((d_2037_v89_)[d_2036_v88_] if (d_2036_v88_) in (d_2037_v89_) else (self).f15): (d_2005_v66_).f20})
                nw348_[int(15)] = (d_1948_v21_).set(-153, (d_2005_v66_).f20)
                nw348_[int(16)] = d_1948_v21_
                nw348_[int(17)] = _dafny.Map({(self).f15: (d_2002_v64_).f17})
                nw348_[int(18)] = d_1948_v21_
                nw348_[int(19)] = d_1948_v21_
                nw348_[int(20)] = d_1948_v21_
                nw348_[int(21)] = d_1948_v21_
                d_2038_v90_ = nw348_
                d_2040_v91_: C4
                nw349_ = C4()
                nw349_.ctor__((d_2036_v88_).f11)
                d_2040_v91_ = nw349_
            d_2041_v92_: _dafny.Array
            def lambda125_(d_2042_i14_):
                return _dafny.CodePoint('o')

            init67_ = lambda125_
            nw350_ = _dafny.Array(None, 18)
            for i0_67_ in range(nw350_.length(0)):
                nw350_[i0_67_] = init67_(i0_67_)
            d_2041_v92_ = nw350_
            d_2043_v93_: str
            d_2043_v93_ = _dafny.CodePoint('d')
            index337_ = default__.safeIndex(948, (d_2041_v92_).length(0))
            (d_2041_v92_)[index337_] = d_2043_v93_
            index338_ = default__.safeIndex(948, (d_2041_v92_).length(0))
            (d_2041_v92_)[index338_] = d_2043_v93_
            d_2044_v94_: D11
            d_2044_v94_ = D11_DC26((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))])
            source24_ = d_2044_v94_
            if source24_.is_DC26:
                d_2045___mcc_h0_ = source24_.cf47
                d_2046_cf47_ = d_2045___mcc_h0_
                d_2047_v95_: _dafny.Map
                d_2047_v95_ = _dafny.Map({self.f16: (-545) * (d_2046_cf47_)})
                d_2048_v96_: _dafny.MultiSet
                d_2048_v96_ = _dafny.MultiSet([(d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]])
                def iife132_():
                    coll70_ = _dafny.Map()
                    compr_70_: int
                    for compr_70_ in _dafny.IntegerRange(609, 344):
                        d_2049_v97_: int = compr_70_
                        if ((609) <= (d_2049_v97_)) and ((d_2049_v97_) < (344)):
                            coll70_[(d_2049_v97_) * (d_2046_cf47_)] = (d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]
                    return _dafny.Map(coll70_)
                d_2047_v95_ = (d_2047_v95_).set((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bj"))) + (default__.fm2(-834, d_2048_v96_, iife132_()
                , globalState)), d_2046_cf47_)
                (globalState).f9 = 601
                d_2050_v98_: _dafny.Map
                d_2050_v98_ = _dafny.Map({d_2043_v93_: p0})
                index339_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                rhs378_ = not ((d_2005_v66_).f20) or (((d_2050_v98_)[(d_2041_v92_)[default__.safeIndex(948, (d_2041_v92_).length(0))]] if ((d_2041_v92_)[default__.safeIndex(948, (d_2041_v92_).length(0))]) in (d_2050_v98_) else (d_2005_v66_).f20))
                rhs379_ = (0) - ((default__.safeDivisionInt((self).f15, d_2046_cf47_)) + ((d_1946_v19_)[default__.safeIndex(133, (d_1946_v19_).length(0))]))
                lhs290_ = globalState
                lhs291_ = d_1946_v19_
                lhs292_ = default__.safeIndex(133, (d_1946_v19_).length(0))
                lhs290_.f8 = rhs378_
                lhs291_[lhs292_] = rhs379_
                d_1943_v16_ = (d_1943_v16_).set((d_2005_v66_).f20, True)
            elif source24_.is_DC25:
                d_2051___mcc_h1_ = source24_.cf46
                d_2052_cf46_ = d_2051___mcc_h1_
                (globalState).f8 = default__.fm1(default__.fm1((d_2002_v64_).f17, d_2043_v93_, globalState), (d_2041_v92_)[default__.safeIndex(948, (d_2041_v92_).length(0))], globalState)
                d_2053_v99_: _dafny.Array
                def lambda126_(d_2054_v21_):
                    def lambda127_(d_2055_i15_):
                        return d_2054_v21_

                    return lambda127_

                init68_ = lambda126_(d_1948_v21_)
                nw351_ = _dafny.Array(None, 2)
                for i0_68_ in range(nw351_.length(0)):
                    nw351_[i0_68_] = init68_(i0_68_)
                d_2053_v99_ = nw351_
                d_2056_v100_: C2
                nw352_ = C2()
                nw352_.ctor__(d_2053_v99_)
                d_2056_v100_ = nw352_
                (globalState).f8 = (d_1941_v14_).isdisjoint(_dafny.MultiSet(d_1947_v20_))
                d_2057_v101_: C3
                nw353_ = C3()
                nw353_.ctor__(((self).f15) != ((_dafny.MultiSet(_dafny.SeqWithoutIsStrInference([(self).f15 for d_2058_i16_ in range(default__.abs(-388))]))).cardinality), p0, d_2053_v99_)
                d_2057_v101_ = nw353_
            elif True:
                d_2059___mcc_h2_ = source24_.cf48
                d_2060_cf48_ = d_2059___mcc_h2_
                (globalState).f8 = ((d_2005_v66_).f20 if (d_2002_v64_).f17 else (d_2005_v66_).f20)
                d_2061_v102_: D1
                d_2061_v102_ = D1_DC4(_dafny.CodePoint('f'))
                d_2061_v102_ = d_2061_v102_
                r0 = d_2043_v93_
                d_2062_v103_: _dafny.Array
                nw354_ = _dafny.Array(None, 1)
                d_2062_v103_ = nw354_
                d_2063_v104_: _dafny.Array
                def lambda128_(d_2064_v21_):
                    def lambda129_(d_2065_i17_):
                        return d_2064_v21_

                    return lambda129_

                init69_ = lambda128_(d_1948_v21_)
                nw355_ = _dafny.Array(None, 5)
                for i0_69_ in range(nw355_.length(0)):
                    nw355_[i0_69_] = init69_(i0_69_)
                d_2063_v104_ = nw355_
                d_2066_v105_: C5
                nw356_ = C5()
                nw356_.ctor__(d_2063_v104_)
                d_2066_v105_ = nw356_
                index340_ = default__.safeIndex(596, (d_2062_v103_).length(0))
                (d_2062_v103_)[index340_] = d_2066_v105_
                index341_ = default__.safeIndex(596, (d_2062_v103_).length(0))
                (d_2062_v103_)[index341_] = d_2066_v105_
        d_2067_v106_: str
        d_2067_v106_ = _dafny.CodePoint('g')
        r0 = d_2067_v106_
        return r0

    def m7(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: _dafny.Seq = _dafny.Seq({})
        r3: bool = False
        d_2068_v0_: _dafny.Seq
        d_2068_v0_ = _dafny.SeqWithoutIsStrInference([p1])
        d_2069_v1_: _dafny.Set
        d_2069_v1_ = _dafny.Set({p0, (self).f15, len(d_2068_v0_), 854})
        hi15_ = len(d_2069_v1_)
        for d_2070_i0_ in range(100, hi15_):
            d_2071_v3_: _dafny.Array
            nw357_ = _dafny.Array(int(0), 22)
            d_2071_v3_ = nw357_
            d_2072_v4_: _dafny.Map
            def iife133_():
                coll71_ = _dafny.Map()
                compr_71_: int
                for compr_71_ in _dafny.IntegerRange(692, 698):
                    d_2073_v2_: int = compr_71_
                    if ((692) <= (d_2073_v2_)) and ((d_2073_v2_) < (698)):
                        coll71_[(d_2073_v2_) - (len(_dafny.Map({p1: p1})))] = _dafny.Map({p1: p0})
                return _dafny.Map(coll71_)
            d_2072_v4_ = _dafny.Map({iife133_()
            : d_2071_v3_})
            d_2074_v5_: _dafny.Map
            d_2074_v5_ = _dafny.Map({p1: (self).f15})
            d_2072_v4_ = (d_2072_v4_).set(_dafny.Map({(self).f15: d_2074_v5_}), d_2071_v3_)
            d_2075_v6_: _dafny.Seq
            d_2075_v6_ = _dafny.SeqWithoutIsStrInference([d_2070_i0_, d_2070_i0_])
            d_2076_v7_: _dafny.Map
            d_2076_v7_ = _dafny.Map({len(d_2075_v6_): _dafny.SeqWithoutIsStrInference([p0, len(self.f16), 165])})
            d_2077_v8_: _dafny.MultiSet
            d_2077_v8_ = _dafny.MultiSet([-357])
            d_2078_v9_: _dafny.Map
            d_2078_v9_ = _dafny.Map({p1: p1})
            d_2079_v10_: _dafny.Set
            d_2079_v10_ = _dafny.Set({_dafny.Map({p1: p1}), d_2078_v9_, d_2078_v9_, d_2078_v9_})
            d_2080_v11_: _dafny.Set
            d_2080_v11_ = d_2079_v10_
            (globalState).f9 = ((0) - (len(((d_2076_v7_)[(d_2077_v8_).cardinality] if ((d_2077_v8_).cardinality) in (d_2076_v7_) else _dafny.SeqWithoutIsStrInference([(d_2077_v8_).cardinality, p0]))))) + (len((_dafny.Set({d_2078_v9_}))))
            default__.m0(globalState)
            d_2081_v12_: _dafny.Array
            nw358_ = _dafny.Array(_dafny.Array(None, 0), 24)
            d_2081_v12_ = nw358_
            index342_ = default__.safeIndex(130, (d_2081_v12_).length(0))
            (d_2081_v12_)[index342_] = d_2071_v3_
            index343_ = default__.safeIndex(130, (d_2081_v12_).length(0))
            rhs380_ = d_2070_i0_
            rhs381_ = self.f16
            rhs382_ = self.f16
            rhs383_ = default__.safeDivisionInt(p0, p0)
            rhs384_ = (d_2071_v3_ if (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vrgni")))) >= (p0) else d_2071_v3_)
            lhs293_ = self
            lhs294_ = self
            lhs295_ = globalState
            lhs296_ = d_2081_v12_
            lhs297_ = default__.safeIndex(130, (d_2081_v12_).length(0))
            r1 = rhs380_
            lhs293_.f16 = rhs381_
            lhs294_.f16 = rhs382_
            lhs295_.f9 = rhs383_
            lhs296_[lhs297_] = rhs384_
        r0 = (0) - (((self).f15) + ((self).f15))
        d_2082_i1_: int
        d_2082_i1_ = 0
        with _dafny.label("10"):
            while p1:
                with _dafny.c_label("10"):
                    if (d_2082_i1_) >= (100):
                        raise _dafny.Break("10")
                    d_2082_i1_ = (d_2082_i1_) + (1)
                    (globalState).f8 = p1
                    d_2083_v13_: str
                    d_2083_v13_ = _dafny.CodePoint('j')
                    d_2084_v14_: _dafny.MultiSet
                    d_2084_v14_ = _dafny.MultiSet([len(default__.fm39(p0, globalState)), (self).f15])
                    d_2085_v15_: _dafny.Seq
                    d_2085_v15_ = _dafny.SeqWithoutIsStrInference([p0, (self).f15, p0, p0, len(self.f16)])
                    d_2086_v16_: _dafny.Array
                    nw359_ = _dafny.Array(None, 26)
                    nw359_[int(0)] = p1
                    nw359_[int(1)] = p1
                    nw359_[int(2)] = p1
                    nw359_[int(3)] = p1
                    nw359_[int(4)] = p1
                    nw359_[int(5)] = default__.fm1(p1, _dafny.CodePoint('f'), globalState)
                    nw359_[int(6)] = p1
                    nw359_[int(7)] = (p0) > (len(self.f16))
                    nw359_[int(8)] = not((d_2069_v1_).isdisjoint(d_2069_v1_))
                    nw359_[int(9)] = p1
                    nw359_[int(10)] = (d_2068_v0_)[default__.safeIndex((self).f15, len(d_2068_v0_))]
                    nw359_[int(11)] = p1
                    nw359_[int(12)] = False
                    nw359_[int(13)] = p1
                    nw359_[int(14)] = p1
                    nw359_[int(15)] = p1
                    nw359_[int(16)] = p1
                    nw359_[int(17)] = default__.fm1(p1, d_2083_v13_, globalState)
                    nw359_[int(18)] = ((default__.fm38(p1, p1, (self.f16)[default__.safeIndex((self).f15, len(self.f16))], p0, globalState)).set((self).f15, default__.abs(p0))).ispropersubset(d_2084_v14_)
                    nw359_[int(19)] = p1
                    nw359_[int(20)] = (len(d_2085_v15_)) <= (p0)
                    nw359_[int(21)] = ((self.f16)[default__.safeIndex((d_2084_v14_).cardinality, len(self.f16))]) in (self.f16)
                    nw359_[int(22)] = p1
                    nw359_[int(23)] = p1
                    nw359_[int(24)] = p1
                    nw359_[int(25)] = p1
                    d_2086_v16_ = nw359_
                    index344_ = default__.safeIndex(265, (d_2086_v16_).length(0))
                    (d_2086_v16_)[index344_] = p1
                    index345_ = default__.safeIndex(265, (d_2086_v16_).length(0))
                    (d_2086_v16_)[index345_] = (p1 if p1 else p1)
                    index346_ = default__.safeIndex(265, (d_2086_v16_).length(0))
                    (d_2086_v16_)[index346_] = p1
                    d_2087_v17_: _dafny.Array
                    nw360_ = _dafny.Array(None, 6)
                    d_2087_v17_ = nw360_
                    d_2088_v18_: _dafny.Array
                    d_2088_v18_ = d_2087_v17_
                    d_2089_v19_: _dafny.Seq
                    d_2089_v19_ = _dafny.SeqWithoutIsStrInference([d_2087_v17_, d_2087_v17_])
                    d_2090_v20_: _dafny.Array
                    nw361_ = _dafny.Array(None, 29)
                    nw361_[int(0)] = d_2087_v17_
                    nw361_[int(1)] = (d_2088_v18_)
                    nw361_[int(2)] = d_2087_v17_
                    nw361_[int(3)] = d_2087_v17_
                    nw361_[int(4)] = d_2087_v17_
                    nw361_[int(5)] = d_2087_v17_
                    nw361_[int(6)] = (d_2089_v19_)[default__.safeIndex((self).f15, len(d_2089_v19_))]
                    nw361_[int(7)] = d_2087_v17_
                    nw361_[int(8)] = d_2087_v17_
                    nw361_[int(9)] = d_2087_v17_
                    nw361_[int(10)] = d_2087_v17_
                    nw361_[int(11)] = d_2087_v17_
                    nw361_[int(12)] = d_2087_v17_
                    nw361_[int(13)] = d_2087_v17_
                    nw361_[int(14)] = d_2087_v17_
                    nw361_[int(15)] = d_2087_v17_
                    nw361_[int(16)] = d_2087_v17_
                    nw361_[int(17)] = d_2087_v17_
                    nw361_[int(18)] = d_2087_v17_
                    nw361_[int(19)] = (d_2087_v17_ if p1 else d_2087_v17_)
                    nw361_[int(20)] = d_2087_v17_
                    nw361_[int(21)] = d_2087_v17_
                    nw361_[int(22)] = d_2087_v17_
                    nw361_[int(23)] = d_2087_v17_
                    nw361_[int(24)] = (d_2087_v17_ if False else d_2087_v17_)
                    nw361_[int(25)] = d_2087_v17_
                    nw361_[int(26)] = d_2087_v17_
                    nw361_[int(27)] = d_2087_v17_
                    nw361_[int(28)] = d_2087_v17_
                    d_2090_v20_ = nw361_
                    index347_ = default__.safeIndex(652, (d_2090_v20_).length(0))
                    (d_2090_v20_)[index347_] = d_2087_v17_
                    index348_ = default__.safeIndex(265, (d_2086_v16_).length(0))
                    index349_ = default__.safeIndex(652, (d_2090_v20_).length(0))
                    rhs385_ = (p0) < (p0)
                    rhs386_ = d_2087_v17_
                    rhs387_ = d_2086_v16_
                    rhs388_ = len(default__.fm59(globalState))
                    lhs298_ = d_2086_v16_
                    lhs299_ = default__.safeIndex(265, (d_2086_v16_).length(0))
                    lhs300_ = d_2090_v20_
                    lhs301_ = default__.safeIndex(652, (d_2090_v20_).length(0))
                    lhs298_[lhs299_] = rhs385_
                    lhs300_[lhs301_] = rhs386_
                    d_2086_v16_ = rhs387_
                    r1 = rhs388_
                    pass
            pass
        d_2091_v21_: _dafny.MultiSet
        d_2091_v21_ = _dafny.MultiSet([(self).f15])
        if ((self).f15) < (((d_2091_v21_)[p0] if (p0) in (d_2091_v21_) else (self).f15)):
            d_2092_v22_: _dafny.Array
            nw362_ = _dafny.Array(int(0), 6)
            d_2092_v22_ = nw362_
            index350_ = default__.safeIndex(827, (d_2092_v22_).length(0))
            (d_2092_v22_)[index350_] = p0
            d_2093_v23_: _dafny.Array
            nw363_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 12)
            d_2093_v23_ = nw363_
            index351_ = default__.safeIndex(827, (d_2092_v22_).length(0))
            rhs389_ = self.f16
            rhs390_ = (self).f15
            rhs391_ = (self).f15
            rhs392_ = d_2093_v23_
            lhs302_ = self
            lhs303_ = d_2092_v22_
            lhs304_ = default__.safeIndex(827, (d_2092_v22_).length(0))
            lhs302_.f16 = rhs389_
            lhs303_[lhs304_] = rhs390_
            r0 = rhs391_
            d_2093_v23_ = rhs392_
            d_2094_v24_: _dafny.Map
            d_2094_v24_ = _dafny.Map({p0: self.f16})
            d_2095_v25_: D9
            d_2095_v25_ = D9_DC22((self).f15, d_2094_v24_)
            r0 = (d_2095_v25_).cf40
            d_2096_v26_: _dafny.Map
            d_2096_v26_ = _dafny.Map({p1: p1})
            d_2097_v27_: _dafny.Seq
            d_2097_v27_ = _dafny.SeqWithoutIsStrInference([d_2096_v26_])
            d_2098_v28_: D0
            d_2098_v28_ = D0_DC1(_dafny.MultiSet([d_2092_v22_, d_2092_v22_]), default__.fm0(p1, globalState), len(d_2097_v27_))
            if not((p1 if (d_2068_v0_)[default__.safeIndex((d_2098_v28_).cf3, len(d_2068_v0_))] else p1)):
                d_2099_v29_: _dafny.Array
                def lambda130_(d_2100_p1_):
                    def lambda131_(d_2101_i2_):
                        return _dafny.Map({(self).f15: d_2100_p1_})

                    return lambda131_

                init70_ = lambda130_(p1)
                nw364_ = _dafny.Array(None, 21)
                for i0_70_ in range(nw364_.length(0)):
                    nw364_[i0_70_] = init70_(i0_70_)
                d_2099_v29_ = nw364_
                d_2102_v30_: T1
                nw365_ = C5()
                nw365_.ctor__(d_2099_v29_)
                d_2102_v30_ = nw365_
                d_2102_v30_ = d_2102_v30_
                r1 = ((d_2092_v22_)[default__.safeIndex(827, (d_2092_v22_).length(0))]) * (p0)
                d_2103_v31_: str
                d_2103_v31_ = _dafny.CodePoint('s')
                d_2104_v32_: D0
                d_2104_v32_ = D0_DC0(p0)
                d_2105_v33_: _dafny.Map
                d_2105_v33_ = _dafny.Map({(d_2104_v32_).cf0: d_2103_v31_})
                d_2103_v31_ = ((d_2105_v33_)[default__.fm0(True, globalState)] if (default__.fm0(True, globalState)) in (d_2105_v33_) else _dafny.CodePoint('p'))
                d_2106_v34_: _dafny.Array
                nw366_ = _dafny.Array(_dafny.Array(None, 0), 14)
                d_2106_v34_ = nw366_
                d_2107_v35_: _dafny.Array
                def lambda132_(d_2108_i3_):
                    return True

                init71_ = lambda132_
                nw367_ = _dafny.Array(None, 16)
                for i0_71_ in range(nw367_.length(0)):
                    nw367_[i0_71_] = init71_(i0_71_)
                d_2107_v35_ = nw367_
                index352_ = default__.safeIndex(775, (d_2106_v34_).length(0))
                (d_2106_v34_)[index352_] = d_2107_v35_
                d_2109_v36_: D10
                d_2109_v36_ = D10_DC24(d_2107_v35_, d_2092_v22_, self.f16)
                index353_ = default__.safeIndex(775, (d_2106_v34_).length(0))
                (d_2106_v34_)[index353_] = (d_2109_v36_).cf43
                d_2110_v37_: _dafny.Map
                d_2110_v37_ = d_2096_v26_
                d_2111_v38_: _dafny.Seq
                d_2111_v38_ = _dafny.SeqWithoutIsStrInference([(0) - ((self).f15)])
                d_2112_v39_: _dafny.Array
                nw368_ = _dafny.Array(None, 28)
                nw368_[int(0)] = d_2096_v26_
                nw368_[int(1)] = d_2096_v26_
                nw368_[int(2)] = (d_2110_v37_)
                nw368_[int(3)] = (d_2096_v26_) | (d_2096_v26_)
                nw368_[int(4)] = d_2096_v26_
                nw368_[int(5)] = d_2096_v26_
                nw368_[int(6)] = (d_2097_v27_)[default__.safeIndex(len(d_2111_v38_), len(d_2097_v27_))]
                nw368_[int(7)] = d_2096_v26_
                nw368_[int(8)] = d_2096_v26_
                nw368_[int(9)] = d_2096_v26_
                nw368_[int(10)] = (d_2096_v26_) | (d_2096_v26_)
                nw368_[int(11)] = d_2096_v26_
                nw368_[int(12)] = (d_2096_v26_).set(p1, p1)
                nw368_[int(13)] = d_2096_v26_
                nw368_[int(14)] = d_2096_v26_
                nw368_[int(15)] = (_dafny.Map({p1: p1}) if True else d_2096_v26_)
                nw368_[int(16)] = d_2096_v26_
                nw368_[int(17)] = (_dafny.Map({p1: default__.fm1(p1, d_2103_v31_, globalState)})) | (_dafny.Map({p1: p1}))
                nw368_[int(18)] = d_2096_v26_
                nw368_[int(19)] = ((d_2097_v27_)[default__.safeIndex((self).f15, len(d_2097_v27_))]) | ((d_2097_v27_)[default__.safeIndex(len(d_2068_v0_), len(d_2097_v27_))])
                nw368_[int(20)] = d_2096_v26_
                nw368_[int(21)] = (d_2096_v26_ if p1 else d_2096_v26_)
                nw368_[int(22)] = d_2096_v26_
                nw368_[int(23)] = (d_2096_v26_).set(p1, not(((d_2096_v26_)[p1] if (p1) in (d_2096_v26_) else p1)))
                nw368_[int(24)] = (d_2110_v37_)
                nw368_[int(25)] = d_2096_v26_
                nw368_[int(26)] = d_2096_v26_
                nw368_[int(27)] = d_2096_v26_
                d_2112_v39_ = nw368_
                index354_ = default__.safeIndex(362, (d_2112_v39_).length(0))
                (d_2112_v39_)[index354_] = d_2096_v26_
                index355_ = default__.safeIndex(362, (d_2112_v39_).length(0))
                (d_2112_v39_)[index355_] = (d_2096_v26_) | (_dafny.Map({p1: p1}))
            elif True:
                d_2113_v40_: _dafny.Set
                d_2113_v40_ = _dafny.Set({p1})
                (globalState).f8 = (d_2113_v40_).isdisjoint(d_2113_v40_)
                d_2114_v41_: _dafny.Map
                d_2114_v41_ = _dafny.Map({p1: p0})
                d_2114_v41_ = (d_2114_v41_).set((p1) and (p1), p0)
                d_2115_v42_: _dafny.Map
                d_2115_v42_ = _dafny.Map({(d_2092_v22_)[default__.safeIndex(827, (d_2092_v22_).length(0))]: p1})
                d_2116_v45_: _dafny.Map
                d_2116_v45_ = _dafny.Map({p0: (self).f15})
                d_2117_v46_: _dafny.Array
                nw369_ = _dafny.Array(None, 27)
                nw369_[int(0)] = d_2115_v42_
                nw369_[int(1)] = d_2115_v42_
                nw369_[int(2)] = d_2115_v42_
                nw369_[int(3)] = d_2115_v42_
                nw369_[int(4)] = d_2115_v42_
                nw369_[int(5)] = d_2115_v42_
                nw369_[int(6)] = d_2115_v42_
                nw369_[int(7)] = _dafny.Map({default__.fm0(((d_2096_v26_)[p1] if (p1) in (d_2096_v26_) else p1), globalState): p1})
                nw369_[int(8)] = d_2115_v42_
                nw369_[int(9)] = d_2115_v42_
                nw369_[int(10)] = d_2115_v42_
                def iife134_():
                    coll72_ = _dafny.Map()
                    compr_72_: int
                    for compr_72_ in _dafny.IntegerRange(194, 972):
                        d_2118_v43_: int = compr_72_
                        if ((194) <= (d_2118_v43_)) and ((d_2118_v43_) < (972)):
                            coll72_[default__.safeDivisionInt(d_2118_v43_, len(_dafny.Map({p1: p1})))] = p1
                    return _dafny.Map(coll72_)
                nw369_[int(11)] = iife134_()
                
                nw369_[int(12)] = d_2115_v42_
                nw369_[int(13)] = d_2115_v42_
                nw369_[int(14)] = d_2115_v42_
                nw369_[int(15)] = d_2115_v42_
                nw369_[int(16)] = d_2115_v42_
                nw369_[int(17)] = d_2115_v42_
                nw369_[int(18)] = d_2115_v42_
                def iife135_():
                    coll73_ = _dafny.Map()
                    compr_73_: int
                    for compr_73_ in (d_2116_v45_).keys.Elements:
                        d_2119_v44_: int = compr_73_
                        if (d_2119_v44_) in (d_2116_v45_):
                            coll73_[(d_2119_v44_) + ((self).f15)] = p1
                    return _dafny.Map(coll73_)
                nw369_[int(19)] = iife135_()
                
                nw369_[int(20)] = d_2115_v42_
                nw369_[int(21)] = d_2115_v42_
                nw369_[int(22)] = d_2115_v42_
                nw369_[int(23)] = d_2115_v42_
                nw369_[int(24)] = d_2115_v42_
                nw369_[int(25)] = d_2115_v42_
                nw369_[int(26)] = d_2115_v42_
                d_2117_v46_ = nw369_
                d_2120_v47_: C4
                nw370_ = C4()
                nw370_.ctor__((d_2117_v46_ if p1 else d_2117_v46_))
                d_2120_v47_ = nw370_
                d_2121_v48_: _dafny.Array
                nw371_ = _dafny.Array(False, 10)
                d_2121_v48_ = nw371_
                d_2122_v49_: D7
                d_2122_v49_ = D7_DC17(self.f16, d_2121_v48_, not(p1), -549)
                d_2123_v50_: _dafny.Set
                d_2123_v50_ = _dafny.Set({d_2092_v22_, d_2092_v22_})
                d_2124_v51_: D7
                d_2124_v51_ = D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "v")), (d_2122_v49_).cf30, p1, len(d_2123_v50_))
                d_2125_v52_: _dafny.Set
                d_2125_v52_ = _dafny.Set({(d_2122_v49_).cf30})
                d_2125_v52_ = d_2125_v52_
                d_2126_v53_: C5
                nw372_ = C5()
                nw372_.ctor__(d_2117_v46_)
                d_2126_v53_ = nw372_
            d_2127_v54_: _dafny.Array
            def lambda133_(d_2128_i4_):
                return False

            init72_ = lambda133_
            nw373_ = _dafny.Array(None, 4)
            for i0_72_ in range(nw373_.length(0)):
                nw373_[i0_72_] = init72_(i0_72_)
            d_2127_v54_ = nw373_
            d_2129_v55_: _dafny.Array
            nw374_ = _dafny.Array(None, 11)
            nw374_[int(0)] = d_2127_v54_
            nw374_[int(1)] = d_2127_v54_
            nw374_[int(2)] = d_2127_v54_
            nw374_[int(3)] = d_2127_v54_
            nw374_[int(4)] = d_2127_v54_
            nw374_[int(5)] = (d_2127_v54_ if p1 else d_2127_v54_)
            nw374_[int(6)] = d_2127_v54_
            nw374_[int(7)] = d_2127_v54_
            nw374_[int(8)] = d_2127_v54_
            nw374_[int(9)] = d_2127_v54_
            nw374_[int(10)] = d_2127_v54_
            d_2129_v55_ = nw374_
            index356_ = default__.safeIndex(491, (d_2129_v55_).length(0))
            (d_2129_v55_)[index356_] = d_2127_v54_
            d_2130_v56_: _dafny.Set
            d_2130_v56_ = _dafny.Set({p1, p1})
            d_2131_v57_: _dafny.Map
            d_2131_v57_ = _dafny.Map({d_2127_v54_: (d_2092_v22_)[default__.safeIndex(827, (d_2092_v22_).length(0))]})
            d_2132_v58_: str
            d_2132_v58_ = _dafny.CodePoint('w')
            d_2133_v59_: D2
            d_2133_v59_ = D2_DC8((0) - (len(d_2131_v57_)), d_2132_v58_)
            pat_let_tv38_ = d_2092_v22_
            pat_let_tv39_ = d_2092_v22_
            index357_ = default__.safeIndex(827, (d_2092_v22_).length(0))
            index358_ = default__.safeIndex(491, (d_2129_v55_).length(0))
            def iife136_(_pat_let31_0):
                def iife137_(d_2134_dt__update__tmp_h1_):
                    def iife138_(_pat_let32_0):
                        def iife139_(d_2135_dt__update_hcf15_h0_):
                            return D2_DC8(d_2135_dt__update_hcf15_h0_, (d_2134_dt__update__tmp_h1_).cf16)
                        return iife139_(_pat_let32_0)
                    return iife138_((pat_let_tv39_)[default__.safeIndex(827, (pat_let_tv38_).length(0))])
                return iife137_(_pat_let31_0)
            rhs393_ = (default__.fm0(p1, globalState)) * ((len(d_2130_v56_)) * ((self).f15))
            rhs394_ = d_2127_v54_
            rhs395_ = (iife136_(d_2133_v59_)).cf15
            lhs305_ = d_2092_v22_
            lhs306_ = default__.safeIndex(827, (d_2092_v22_).length(0))
            lhs307_ = d_2129_v55_
            lhs308_ = default__.safeIndex(491, (d_2129_v55_).length(0))
            lhs309_ = globalState
            lhs305_[lhs306_] = rhs393_
            lhs307_[lhs308_] = rhs394_
            lhs309_.f9 = rhs395_
            d_2136_v60_: _dafny.Seq
            d_2136_v60_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2137_v61_: _dafny.Seq
            d_2137_v61_ = _dafny.SeqWithoutIsStrInference([len(d_2130_v56_), (d_2092_v22_)[default__.safeIndex(827, (d_2092_v22_).length(0))], (0) - ((d_2136_v60_)[default__.safeIndex((d_2092_v22_)[default__.safeIndex(827, (d_2092_v22_).length(0))], len(d_2136_v60_))]), 967])
            d_2138_v62_: _dafny.Seq
            d_2138_v62_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([786]), d_2136_v60_, _dafny.SeqWithoutIsStrInference([-650])])
            (globalState).f9 = (0) - (len((d_2136_v60_) + ((d_2138_v62_)[default__.safeIndex((self).f15, len(d_2138_v62_))])))
        elif True:
            d_2139_v63_: str
            d_2139_v63_ = _dafny.CodePoint('k')
            (globalState).f8 = default__.fm1(p1, d_2139_v63_, globalState)
            d_2140_v64_: _dafny.Array
            def lambda134_(d_2141_p1_):
                def lambda135_(d_2142_i5_):
                    return _dafny.Map({(self).f15: d_2141_p1_})

                return lambda135_

            init73_ = lambda134_(p1)
            nw375_ = _dafny.Array(None, 22)
            for i0_73_ in range(nw375_.length(0)):
                nw375_[i0_73_] = init73_(i0_73_)
            d_2140_v64_ = nw375_
            d_2143_v65_: C4
            nw376_ = C4()
            nw376_.ctor__(d_2140_v64_)
            d_2143_v65_ = nw376_
            d_2144_v66_: _dafny.Seq
            d_2144_v66_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2145_v67_: _dafny.Map
            d_2145_v67_ = _dafny.Map({d_2143_v65_: len(d_2144_v66_)})
            d_2146_v68_: _dafny.MultiSet
            d_2146_v68_ = _dafny.MultiSet([d_2145_v67_])
            d_2147_v69_: _dafny.Seq
            d_2147_v69_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2143_v65_: (self).f15})])
            d_2146_v68_ = _dafny.MultiSet(d_2147_v69_)
            (globalState).f9 = p0
            d_2148_v70_: _dafny.Set
            d_2148_v70_ = _dafny.Set({p1})
            d_2149_v71_: D0
            d_2149_v71_ = D0_DC2(_dafny.MultiSet([not(p1)]), False, (0) - (len(d_2148_v70_)), p1)
            d_2150_v72_: _dafny.Map
            d_2150_v72_ = _dafny.Map({d_2149_v71_: p1})
            d_2151_v73_: _dafny.Map
            d_2151_v73_ = _dafny.Map({(d_2149_v71_) in (d_2150_v72_): p1})
            d_2151_v73_ = (d_2151_v73_).set(p1, p1)
            r3 = p1
        hi16_ = p0
        for d_2152_i6_ in range(len((d_2068_v0_) + ((d_2068_v0_).set(default__.safeIndex((self).f15, len(d_2068_v0_)), p1))), hi16_):
            default__.m0(globalState)
            d_2153_v74_: _dafny.Seq
            d_2153_v74_ = _dafny.SeqWithoutIsStrInference([p0])
            d_2154_v75_: D0
            d_2154_v75_ = D0_DC0(p0)
            d_2155_v76_: C1
            nw377_ = C1()
            nw377_.ctor__(p1, p1, d_2154_v75_, p1)
            d_2155_v76_ = nw377_
            d_2156_v77_: _dafny.Seq
            d_2156_v77_ = _dafny.SeqWithoutIsStrInference([d_2155_v76_])
            d_2157_v78_: _dafny.Seq
            d_2157_v78_ = _dafny.SeqWithoutIsStrInference([d_2152_i6_, ((self).f15) + ((d_2153_v74_)[default__.safeIndex(len(d_2156_v77_), len(d_2153_v74_))]), default__.safeModuloInt((self).f15, p0), p0])
            d_2158_v79_: _dafny.Array
            nw378_ = _dafny.Array(False, 3)
            d_2158_v79_ = nw378_
            index359_ = default__.safeIndex(524, (d_2158_v79_).length(0))
            (d_2158_v79_)[index359_] = (self.f16) <= (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "hapbpk")))
            index360_ = default__.safeIndex(524, (d_2158_v79_).length(0))
            rhs396_ = not(True)
            rhs397_ = (d_2155_v76_).f20
            rhs398_ = d_2153_v74_
            rhs399_ = ((d_2155_v76_).f20) or (p1)
            lhs310_ = globalState
            lhs311_ = d_2158_v79_
            lhs312_ = default__.safeIndex(524, (d_2158_v79_).length(0))
            lhs310_.f8 = rhs396_
            r3 = rhs397_
            d_2153_v74_ = rhs398_
            lhs311_[lhs312_] = rhs399_
            d_2159_v80_: _dafny.MultiSet
            d_2159_v80_ = _dafny.MultiSet([(d_2155_v76_).f20])
            d_2160_v81_: _dafny.Map
            d_2160_v81_ = _dafny.Map({d_2158_v79_: d_2159_v80_})
            d_2160_v81_ = (d_2160_v81_).set(d_2158_v79_, d_2159_v80_)
            r3 = (d_2155_v76_).f19
        hi17_ = (self).f15
        for d_2161_i7_ in range(p0, hi17_):
            d_2162_v82_: _dafny.MultiSet
            d_2162_v82_ = _dafny.MultiSet([not(p1), p1])
            d_2163_v83_: D0
            d_2163_v83_ = D0_DC2(d_2162_v82_, p1, p0, p1)
            d_2164_v84_: D8
            d_2164_v84_ = D8_DC20(d_2163_v83_, p1, p1)
            r3 = (d_2164_v84_).cf38
            (globalState).f8 = p1
            d_2165_v85_: _dafny.Map
            d_2165_v85_ = _dafny.Map({(self).f15: d_2161_i7_})
            d_2166_v86_: _dafny.Seq
            d_2166_v86_ = _dafny.SeqWithoutIsStrInference([32])
            (globalState).f8 = ((len(d_2165_v85_)) - ((0) - (len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('t') for d_2167_i8_ in range(default__.abs(240))]))))) >= ((len(d_2166_v86_)) + ((self).f15))
            d_2168_v87_: str
            d_2168_v87_ = _dafny.CodePoint('e')
            r3 = not (p1) or (default__.fm1(p1, d_2168_v87_, globalState))
        r0 = (self).f15
        r1 = p0
        d_2169_v88_: D2
        d_2169_v88_ = D2_DC5(_dafny.SeqWithoutIsStrInference([(self.f16)[default__.safeIndex((0) - (p0), len(self.f16))] for d_2170_i9_ in range(default__.abs(161))]))
        d_2171_v89_: _dafny.Seq
        def iife140_(_pat_let33_0):
            def iife141_(d_2172_dt__update__tmp_h2_):
                def iife142_(_pat_let34_0):
                    def iife143_(d_2173_dt__update_hcf10_h0_):
                        return D2_DC5(d_2173_dt__update_hcf10_h0_)
                    return iife143_(_pat_let34_0)
                return iife142_(self.f16)
            return iife141_(_pat_let33_0)
        def iife144_(_pat_let35_0):
            def iife145_(d_2174_dt__update__tmp_h3_):
                def iife146_(_pat_let36_0):
                    def iife147_(d_2175_dt__update_hcf10_h1_):
                        return D2_DC5(d_2175_dt__update_hcf10_h1_)
                    return iife147_(_pat_let36_0)
                return iife146_(self.f16)
            return iife145_(_pat_let35_0)
        d_2171_v89_ = _dafny.SeqWithoutIsStrInference([d_2169_v88_, d_2169_v88_, iife140_(d_2169_v88_), iife144_(D2_DC5(self.f16))])
        r2 = d_2171_v89_
        r3 = p1
        return r0, r1, r2, r3

    @property
    def f15(self):
        return self._f15

class C7(T1, T2, T0):
    def  __init__(self):
        self._f12: D0 = D0.default()()
        self._f11: _dafny.Array = _dafny.Array(None, 0)
        self._f13: bool = False
        self._f14: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C7"
    @property
    def f12(self):
        return self._f12
    @f12.setter
    def f12(self, value):
        self._f12 = value
    @property
    def f11(self):
        return self._f11
    @property
    def f13(self):
        return self._f13
    def ctor__(self, f14, f11, f12, f13):
        (self)._f14 = f14
        (self)._f11 = f11
        (self).f12 = f12
        (self)._f13 = f13

    def fm8(self, globalState):
        source25_ = D0_DC2(_dafny.MultiSet([not((self).f14), (self).f13]), (self).f14, (_dafny.MultiSet([(0) - (-158)])).cardinality, (self).f14)
        if source25_.is_DC1:
            d_2176___mcc_h0_ = source25_.cf1
            d_2177___mcc_h1_ = source25_.cf2
            d_2178___mcc_h2_ = source25_.cf3
            d_2179_cf3_ = d_2178___mcc_h2_
            d_2180_cf2_ = d_2177___mcc_h1_
            d_2181_cf1_ = d_2176___mcc_h0_
            return D2_DC6()
        elif source25_.is_DC2:
            d_2182___mcc_h3_ = source25_.cf4
            d_2183___mcc_h4_ = source25_.cf5
            d_2184___mcc_h5_ = source25_.cf6
            d_2185___mcc_h6_ = source25_.cf7
            d_2186_cf7_ = d_2185___mcc_h6_
            d_2187_cf6_ = d_2184___mcc_h5_
            d_2188_cf5_ = d_2183___mcc_h4_
            d_2189_cf4_ = d_2182___mcc_h3_
            return D2_DC6()
        elif True:
            d_2190___mcc_h7_ = source25_.cf0
            d_2191_cf0_ = d_2190___mcc_h7_
            return D2_DC6()

    def fm9(self, p0, p1, globalState):
        return default__.safeModuloInt(default__.safeDivisionInt(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "equcqbrac"))), (0) - ((_dafny.MultiSet([_dafny.CodePoint('w')])).cardinality)), len((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "bxaa"))) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "lgrqsywy")))))

    def fm10(self, globalState):
        return (self).f14

    def fm11(self, p0, p1, p2, p3, globalState):
        return default__.safeDivisionInt(len((_dafny.Map({(0) - (-656): _dafny.SeqWithoutIsStrInference([(self).f14])})) | (_dafny.Map({485: _dafny.SeqWithoutIsStrInference([(self).f14, (self).f13])}))), (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rvymo")))) * ((0) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "x"))))))

    def m4(self, p0, p1, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: int = int(0)
        r2: D2 = D2.default()()
        r3: bool = False
        d_2192_v0_: _dafny.Seq
        d_2192_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "owbxdur"))
        d_2193_v1_: D1
        d_2193_v1_ = D1_DC4((d_2192_v0_)[default__.safeIndex(p0, len(d_2192_v0_))])
        d_2194_v3_: str
        d_2194_v3_ = _dafny.CodePoint('v')
        d_2195_v4_: _dafny.MultiSet
        d_2195_v4_ = _dafny.MultiSet([(self).f14])
        d_2196_v5_: _dafny.Set
        d_2196_v5_ = _dafny.Set({(self).f13})
        d_2197_v6_: D0
        d_2197_v6_ = D0_DC2(d_2195_v4_, False, len(d_2196_v5_), (self).f14)
        d_2198_v7_: _dafny.Seq
        d_2198_v7_ = _dafny.SeqWithoutIsStrInference([len(d_2192_v0_)])
        d_2199_v8_: _dafny.Map
        d_2199_v8_ = _dafny.Map({d_2198_v7_: p0})
        d_2200_v9_: _dafny.MultiSet
        d_2200_v9_ = _dafny.MultiSet([p0])
        d_2201_v10_: _dafny.MultiSet
        d_2201_v10_ = _dafny.MultiSet([((d_2199_v8_)[d_2198_v7_] if (d_2198_v7_) in (d_2199_v8_) else ((d_2200_v9_)).cardinality)])
        d_2202_v11_: _dafny.Set
        d_2202_v11_ = _dafny.Set({(self).fm11(default__.fm1((self).f13, d_2194_v3_, globalState), False, p0, (d_2197_v6_).cf5, globalState), ((d_2201_v10_)[p0] if (p0) in (d_2201_v10_) else len(d_2192_v0_)), p0})
        d_2203_v12_: _dafny.Map
        def iife148_():
            coll74_ = _dafny.Set()
            compr_74_: int
            for compr_74_ in _dafny.IntegerRange(-48, 590):
                d_2204_v2_: int = compr_74_
                if ((-48) <= (d_2204_v2_)) and ((d_2204_v2_) < (590)):
                    coll74_ = coll74_.union(_dafny.Set([default__.safeModuloInt(d_2204_v2_, p0)]))
            return _dafny.Set(coll74_)
        d_2203_v12_ = _dafny.Map({(self).fm9(d_2193_v1_, iife148_()
        , globalState): (d_2202_v11_).issubset(d_2202_v11_)})
        d_2203_v12_ = (d_2203_v12_).set(p0, (self).f14)
        d_2205_v13_: _dafny.Array
        nw379_ = _dafny.Array(False, 25)
        d_2205_v13_ = nw379_
        index361_ = default__.safeIndex(795, (d_2205_v13_).length(0))
        (d_2205_v13_)[index361_] = (self).f13
        d_2206_v14_: _dafny.MultiSet
        d_2206_v14_ = _dafny.MultiSet([p1, p1, p1])
        index362_ = default__.safeIndex(795, (d_2205_v13_).length(0))
        (d_2205_v13_)[index362_] = (d_2206_v14_).ispropersubset((d_2206_v14_).set(p1, default__.abs(len(d_2192_v0_))))
        d_2207_v15_: _dafny.Array
        nw380_ = _dafny.Array(D1.default()(), 14)
        d_2207_v15_ = nw380_
        d_2207_v15_ = d_2207_v15_
        d_2208_v16_: D2
        d_2208_v16_ = D2_DC6()
        source26_ = d_2208_v16_
        if source26_.is_DC6:
            d_2209_v17_: D2
            d_2209_v17_ = D2_DC7((self).f14, p0, (self).f13, _dafny.SeqWithoutIsStrInference([d_2194_v3_ for d_2210_i0_ in range(default__.abs(272))]))
            rhs400_ = ((d_2195_v4_) - (_dafny.MultiSet([(self).f13, (self).f14]))) - ((d_2197_v6_).cf4)
            rhs401_ = default__.safeDivisionInt((0) - (default__.fm0(not((self).f13), globalState)), (d_2209_v17_).cf12)
            rhs402_ = default__.fm0((self).f13, globalState)
            lhs313_ = globalState
            lhs314_ = globalState
            d_2195_v4_ = rhs400_
            lhs313_.f9 = rhs401_
            lhs314_.f9 = rhs402_
            index363_ = default__.safeIndex(795, (d_2205_v13_).length(0))
            rhs403_ = (p0) * (p0)
            rhs404_ = default__.safeDivisionInt(p0, len(_dafny.Set({(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]})))
            rhs405_ = ((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))] if not ((self).f13) or (not((self).f14)) else (self).f14)
            rhs406_ = p0
            rhs407_ = True
            lhs315_ = globalState
            lhs316_ = d_2205_v13_
            lhs317_ = default__.safeIndex(795, (d_2205_v13_).length(0))
            lhs315_.f9 = rhs403_
            r1 = rhs404_
            r3 = rhs405_
            r1 = rhs406_
            lhs316_[lhs317_] = rhs407_
            d_2211_v18_: _dafny.MultiSet
            d_2211_v18_ = _dafny.MultiSet([d_2205_v13_, d_2205_v13_])
            d_2212_v19_: C1
            nw381_ = C1()
            nw381_.ctor__(((self).f14) and ((self).f13), (d_2211_v18_) != (_dafny.MultiSet([d_2205_v13_])), default__.fm40(d_2194_v3_, globalState), (self).f13)
            d_2212_v19_ = nw381_
            index364_ = default__.safeIndex(896, (p1).length(0))
            (p1)[index364_] = p0
            d_2213_v20_: _dafny.Seq
            d_2213_v20_ = _dafny.SeqWithoutIsStrInference([(self).f14, (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]])
            d_2214_v21_: _dafny.Map
            d_2214_v21_ = _dafny.Map({d_2213_v20_: default__.fm1((self).f14, d_2194_v3_, globalState)})
            index365_ = default__.safeIndex(896, (p1).length(0))
            rhs408_ = ((d_2192_v0_) <= (d_2192_v0_) if ((self).f14 if (d_2212_v19_).f19 else (self).f13) else (d_2213_v20_)[default__.safeIndex(567, len(d_2213_v20_))])
            rhs409_ = (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]
            rhs410_ = (p0) - (len(d_2214_v21_))
            rhs411_ = ((d_2202_v11_).intersection(_dafny.Set({p0}))).issubset(d_2202_v11_)
            rhs412_ = d_2194_v3_
            lhs318_ = globalState
            lhs319_ = p1
            lhs320_ = default__.safeIndex(896, (p1).length(0))
            lhs318_.f8 = rhs408_
            r3 = rhs409_
            lhs319_[lhs320_] = rhs410_
            r3 = rhs411_
            d_2194_v3_ = rhs412_
        elif source26_.is_DC7:
            d_2215___mcc_h0_ = source26_.cf11
            d_2216___mcc_h1_ = source26_.cf12
            d_2217___mcc_h2_ = source26_.cf13
            d_2218___mcc_h3_ = source26_.cf14
            d_2219_cf14_ = d_2218___mcc_h3_
            d_2220_cf13_ = d_2217___mcc_h2_
            d_2221_cf12_ = d_2216___mcc_h1_
            d_2222_cf11_ = d_2215___mcc_h0_
            d_2223_v22_: _dafny.Set
            d_2223_v22_ = _dafny.Set({d_2205_v13_})
            d_2224_v23_: _dafny.Map
            d_2224_v23_ = _dafny.Map({p1: d_2223_v22_})
            d_2224_v23_ = (d_2224_v23_).set(p1, (d_2223_v22_).intersection(d_2223_v22_))
            if (self).f13:
                (globalState).f9 = ((d_2221_cf12_) - (p0)) * (len(d_2192_v0_))
                d_2225_v24_: _dafny.Seq
                d_2225_v24_ = _dafny.SeqWithoutIsStrInference([False])
                d_2192_v0_ = (d_2219_cf14_) + ((default__.fm60(len(d_2225_v24_), not(d_2222_cf11_), globalState)).cf34)
                d_2226_v25_: _dafny.Array
                def lambda136_(d_2227_v3_):
                    def lambda137_(d_2228_i1_):
                        return d_2227_v3_

                    return lambda137_

                init74_ = lambda136_(d_2194_v3_)
                nw382_ = _dafny.Array(None, 29)
                for i0_74_ in range(nw382_.length(0)):
                    nw382_[i0_74_] = init74_(i0_74_)
                d_2226_v25_ = nw382_
                index366_ = default__.safeIndex(437, (d_2226_v25_).length(0))
                (d_2226_v25_)[index366_] = d_2194_v3_
                index367_ = default__.safeIndex(437, (d_2226_v25_).length(0))
                (d_2226_v25_)[index367_] = d_2194_v3_
                d_2225_v24_ = d_2225_v24_
                index368_ = default__.safeIndex(437, (d_2226_v25_).length(0))
                (d_2226_v25_)[index368_] = default__.fm3((0) - (d_2221_cf12_), d_2222_cf11_, globalState)
            elif True:
                d_2229_v26_: C0
                nw383_ = C0()
                nw383_.ctor__(d_2222_cf11_)
                d_2229_v26_ = nw383_
                d_2230_v27_: C5
                nw384_ = C5()
                nw384_.ctor__((self).f11)
                d_2230_v27_ = nw384_
                index369_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index369_] = ((d_2221_cf12_) >= (p0)) and ((_dafny.Set({(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], (d_2229_v26_).f17, d_2222_cf11_, default__.fm1((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], d_2194_v3_, globalState), True})).issubset(_dafny.Set({(self).f14})))
                d_2231_v28_: _dafny.Map
                d_2231_v28_ = _dafny.Map({(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]: len(d_2192_v0_)})
                d_2231_v28_ = (d_2231_v28_).set((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], p0)
                index370_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index370_] = ((p0) * (619)) > (p0)
            if default__.fm1(False, _dafny.CodePoint('w'), globalState):
                d_2232_v29_: _dafny.Seq
                d_2232_v29_ = _dafny.SeqWithoutIsStrInference([((d_2203_v12_)[(0) - (p0)] if ((0) - (p0)) in (d_2203_v12_) else (self).f14), False])
                d_2232_v29_ = _dafny.SeqWithoutIsStrInference([(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], False])
                r0 = d_2194_v3_
                d_2233_v30_: C6
                nw385_ = C6()
                nw385_.ctor__((p0) * (p0), (d_2192_v0_) + (d_2192_v0_))
                d_2233_v30_ = nw385_
                d_2232_v29_ = d_2232_v29_
                (globalState).f9 = d_2221_cf12_
            elif True:
                d_2193_v1_ = d_2193_v1_
                default__.m0(globalState)
                d_2234_v31_: _dafny.Seq
                d_2234_v31_ = _dafny.SeqWithoutIsStrInference([d_2192_v0_])
                index371_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index371_] = (d_2194_v3_) in ((d_2234_v31_)[default__.safeIndex(p0, len(d_2234_v31_))])
                d_2235_v32_: C6
                nw386_ = C6()
                nw386_.ctor__(p0, d_2192_v0_)
                d_2235_v32_ = nw386_
                d_2236_v33_: _dafny.Seq
                d_2236_v33_ = _dafny.SeqWithoutIsStrInference([d_2222_cf11_])
                d_2237_v34_: D2
                d_2237_v34_ = D2_DC7(d_2220_cf13_, len(((d_2234_v31_)[default__.safeIndex(len(d_2236_v33_), len(d_2234_v31_))]).set(default__.safeIndex((d_2235_v32_).f15, len((d_2234_v31_)[default__.safeIndex(len(d_2236_v33_), len(d_2234_v31_))])), d_2194_v3_)), (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], d_2235_v32_.f16)
                d_2238_v35_: _dafny.Map
                d_2238_v35_ = _dafny.Map({d_2200_v9_: not((self).f13)})
                rhs413_ = default__.fm26((d_2237_v34_).cf13, (0) - (len(d_2238_v35_)), (d_2236_v33_) != (d_2236_v33_), globalState)
                rhs414_ = default__.safeModuloInt(d_2221_cf12_, default__.safeModuloInt(d_2221_cf12_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ilvvkfxmt")))))
                rhs415_ = d_2235_v32_
                rhs416_ = ((self).f14 if d_2220_cf13_ else True)
                lhs321_ = globalState
                lhs322_ = globalState
                d_2198_v7_ = rhs413_
                lhs321_.f9 = rhs414_
                d_2235_v32_ = rhs415_
                lhs322_.f8 = rhs416_
                index372_ = default__.safeIndex(549, (p1).length(0))
                (p1)[index372_] = (p0) - (p0)
                index373_ = default__.safeIndex(549, (p1).length(0))
                index374_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                rhs417_ = (512) * (default__.safeModuloInt((d_2235_v32_).f15, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "aiwr")))))
                rhs418_ = ((0) - (default__.safeModuloInt(326, p0))) != (d_2221_cf12_)
                lhs323_ = p1
                lhs324_ = default__.safeIndex(549, (p1).length(0))
                lhs325_ = d_2205_v13_
                lhs326_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                lhs323_[lhs324_] = rhs417_
                lhs325_[lhs326_] = rhs418_
            index375_ = default__.safeIndex(845, ((self).f11).length(0))
            ((self).f11)[index375_] = _dafny.Map({p0: (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]})
            d_2239_v36_: _dafny.Seq
            d_2239_v36_ = _dafny.SeqWithoutIsStrInference([(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], (self).f13, (self).f13, (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], (self).f13])
            index376_ = default__.safeIndex(845, ((self).f11).length(0))
            rhs419_ = (d_2194_v3_) in (d_2192_v0_)
            rhs420_ = (d_2196_v5_).ispropersubset(d_2196_v5_)
            rhs421_ = d_2203_v12_
            rhs422_ = (0) - ((0) - (p0))
            rhs423_ = d_2239_v36_
            lhs327_ = (self).f11
            lhs328_ = default__.safeIndex(845, ((self).f11).length(0))
            lhs329_ = globalState
            r3 = rhs419_
            d_2220_cf13_ = rhs420_
            lhs327_[lhs328_] = rhs421_
            lhs329_.f9 = rhs422_
            d_2239_v36_ = rhs423_
        elif source26_.is_DC8:
            d_2240___mcc_h4_ = source26_.cf15
            d_2241___mcc_h5_ = source26_.cf16
            d_2242_cf16_ = d_2241___mcc_h5_
            d_2243_cf15_ = d_2240___mcc_h4_
            (globalState).f9 = d_2243_cf15_
            d_2244_v37_: C1
            nw387_ = C1()
            nw387_.ctor__((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], (self).f14, self.f12, (self).f14)
            d_2244_v37_ = nw387_
            def iife149_():
                coll75_ = _dafny.Set()
                compr_75_: int
                for compr_75_ in _dafny.IntegerRange(653, 65):
                    d_2245_v38_: int = compr_75_
                    if ((653) <= (d_2245_v38_)) and ((d_2245_v38_) < (65)):
                        coll75_ = coll75_.union(_dafny.Set([(d_2245_v38_) + (p0)]))
                return _dafny.Set(coll75_)
            d_2243_cf15_ = ((d_2243_cf15_) + ((d_2244_v37_).fm9(D1_DC4(d_2194_v3_), iife149_()
            , globalState))) + (((d_2201_v10_)[(0) - (p0)] if ((0) - (p0)) in (d_2201_v10_) else d_2243_cf15_))
            (globalState).f8 = default__.fm1((self).f13, d_2242_cf16_, globalState)
        elif True:
            d_2246___mcc_h6_ = source26_.cf10
            d_2247_cf10_ = d_2246___mcc_h6_
            rhs424_ = d_2205_v13_
            rhs425_ = (self).f14
            rhs426_ = default__.safeModuloInt(p0, len(d_2196_v5_))
            d_2205_v13_ = rhs424_
            r3 = rhs425_
            r1 = rhs426_
            (globalState).f9 = (0) - (p0)
            d_2248_v39_: D2
            d_2248_v39_ = D2_DC7((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], p0, (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], d_2247_cf10_)
            if not((d_2248_v39_).cf13):
                d_2249_v40_: _dafny.Map
                d_2249_v40_ = _dafny.Map({p0: p1})
                d_2249_v40_ = (d_2249_v40_).set((615 if (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))] else len(default__.fm53(globalState))), p1)
                d_2250_v41_: D4
                d_2250_v41_ = D4_DC11(p0, p0, (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))])
                d_2251_v42_: C0
                nw388_ = C0()
                nw388_.ctor__((d_2250_v41_).cf21)
                d_2251_v42_ = nw388_
                d_2252_v43_: _dafny.Map
                d_2252_v43_ = _dafny.Map({(d_2202_v11_).issubset(d_2202_v11_): d_2208_v16_})
                d_2252_v43_ = (d_2252_v43_).set((d_2251_v42_).f17, d_2208_v16_)
                d_2253_v44_: C4
                nw389_ = C4()
                nw389_.ctor__((self).f11)
                d_2253_v44_ = nw389_
                d_2253_v44_ = d_2253_v44_
                d_2254_v45_: C5
                nw390_ = C5()
                nw390_.ctor__((self).f11)
                d_2254_v45_ = nw390_
            elif True:
                index377_ = default__.safeIndex(42, (p1).length(0))
                (p1)[index377_] = p0
                d_2255_v47_: _dafny.Map
                def iife150_():
                    coll76_ = _dafny.Map()
                    compr_76_: int
                    for compr_76_ in (d_2198_v7_).Elements:
                        d_2256_v46_: int = compr_76_
                        if (d_2256_v46_) in (d_2198_v7_):
                            coll76_[(d_2256_v46_) * (p0)] = _dafny.CodePoint('i')
                    return _dafny.Map(coll76_)
                d_2255_v47_ = _dafny.Map({False: len(iife150_()
                )})
                d_2257_v48_: _dafny.MultiSet
                d_2257_v48_ = _dafny.MultiSet([(d_2255_v47_) | (d_2255_v47_), d_2255_v47_, d_2255_v47_, (default__.fm50(d_2194_v3_, (self).f14, (self).f13, globalState)) | (d_2255_v47_), d_2255_v47_])
                d_2258_v49_: D2
                d_2258_v49_ = D2_DC8(p0, d_2194_v3_)
                d_2259_v50_: _dafny.MultiSet
                d_2259_v50_ = _dafny.MultiSet([default__.fm37((self).f14, len(d_2192_v0_), p0, p0, globalState), d_2202_v11_, d_2202_v11_])
                index378_ = default__.safeIndex(42, (p1).length(0))
                rhs427_ = ((d_2257_v48_)[d_2255_v47_] if (d_2255_v47_) in (d_2257_v48_) else p0)
                rhs428_ = (d_2198_v7_)[default__.safeIndex(len(d_2192_v0_), len(d_2198_v7_))]
                rhs429_ = default__.fm1((self).f13, (d_2258_v49_).cf16, globalState)
                rhs430_ = default__.fm0(default__.fm1((self).f13, d_2194_v3_, globalState), globalState)
                rhs431_ = ((d_2259_v50_)[d_2202_v11_] if (d_2202_v11_) in (d_2259_v50_) else p0)
                lhs330_ = globalState
                lhs331_ = p1
                lhs332_ = default__.safeIndex(42, (p1).length(0))
                lhs333_ = globalState
                lhs334_ = globalState
                lhs335_ = globalState
                lhs330_.f9 = rhs427_
                lhs331_[lhs332_] = rhs428_
                lhs333_.f8 = rhs429_
                lhs334_.f9 = rhs430_
                lhs335_.f9 = rhs431_
                d_2260_v51_: _dafny.Array
                nw391_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 17)
                d_2260_v51_ = nw391_
                index379_ = default__.safeIndex(734, (d_2260_v51_).length(0))
                (d_2260_v51_)[index379_] = (_dafny.SeqWithoutIsStrInference([d_2194_v3_ for d_2261_i2_ in range(default__.abs(570))])).set(default__.safeIndex((p1)[default__.safeIndex(42, (p1).length(0))], len(_dafny.SeqWithoutIsStrInference([d_2194_v3_ for d_2262_i2_ in range(default__.abs(570))]))), _dafny.CodePoint('a'))
                index380_ = default__.safeIndex(734, (d_2260_v51_).length(0))
                (d_2260_v51_)[index380_] = (d_2192_v0_) + (d_2247_cf10_)
                d_2263_v52_: _dafny.Array
                nw392_ = _dafny.Array(None, 23)
                d_2263_v52_ = nw392_
                d_2264_v53_: C1
                nw393_ = C1()
                nw393_.ctor__((self).f14, False, self.f12, not(False))
                d_2264_v53_ = nw393_
                index381_ = default__.safeIndex(470, (d_2263_v52_).length(0))
                (d_2263_v52_)[index381_] = d_2264_v53_
                d_2265_v54_: _dafny.Array
                nw394_ = _dafny.Array(int(0), 6)
                d_2265_v54_ = nw394_
                d_2266_v55_: _dafny.Seq
                d_2266_v55_ = _dafny.SeqWithoutIsStrInference([(d_2264_v53_).f19])
                d_2267_v56_: _dafny.Seq
                d_2267_v56_ = _dafny.SeqWithoutIsStrInference([(_dafny.MultiSet(d_2266_v55_)).intersection(_dafny.MultiSet(d_2266_v55_)), d_2195_v4_])
                index382_ = default__.safeIndex(42, (p1).length(0))
                index383_ = default__.safeIndex(470, (d_2263_v52_).length(0))
                index384_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                rhs432_ = len(d_2267_v56_)
                rhs433_ = d_2264_v53_
                rhs434_ = d_2265_v54_
                rhs435_ = (self).f13
                rhs436_ = (default__.fm1((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], d_2194_v3_, globalState)) or ((d_2247_cf10_) == (d_2192_v0_))
                lhs336_ = p1
                lhs337_ = default__.safeIndex(42, (p1).length(0))
                lhs338_ = d_2263_v52_
                lhs339_ = default__.safeIndex(470, (d_2263_v52_).length(0))
                lhs340_ = d_2205_v13_
                lhs341_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                lhs342_ = globalState
                lhs336_[lhs337_] = rhs432_
                lhs338_[lhs339_] = rhs433_
                d_2265_v54_ = rhs434_
                lhs340_[lhs341_] = rhs435_
                lhs342_.f8 = rhs436_
                d_2268_v57_: _dafny.Array
                def lambda138_(d_2269_v39_):
                    def lambda139_(d_2270_i3_):
                        return d_2269_v39_

                    return lambda139_

                init75_ = lambda138_(d_2248_v39_)
                nw395_ = _dafny.Array(None, 23)
                for i0_75_ in range(nw395_.length(0)):
                    nw395_[i0_75_] = init75_(i0_75_)
                d_2268_v57_ = nw395_
                d_2271_v58_: D11
                d_2271_v58_ = D11_DC25(d_2268_v57_)
                d_2272_v59_: D11
                d_2272_v59_ = D11_DC27(d_2271_v58_)
                d_2273_v60_: D11
                d_2273_v60_ = D11_DC27(d_2272_v59_)
                pat_let_tv40_ = d_2272_v59_
                def iife151_(_pat_let37_0):
                    def iife152_(d_2274_dt__update__tmp_h0_):
                        def iife153_(_pat_let38_0):
                            def iife154_(d_2275_dt__update_hcf48_h0_):
                                return D11_DC27(d_2275_dt__update_hcf48_h0_)
                            return iife154_(_pat_let38_0)
                        return iife153_(D11_DC27(D11_DC27(pat_let_tv40_)))
                    return iife152_(_pat_let37_0)
                d_2273_v60_ = iife151_(d_2273_v60_)
                d_2276_v61_: C0
                nw396_ = C0()
                nw396_.ctor__(not((d_2248_v39_).cf13))
                d_2276_v61_ = nw396_
            r1 = p0
        d_2277_v62_: _dafny.Map
        d_2277_v62_ = _dafny.Map({p0: _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "erufxrt"))})
        d_2278_v63_: D9
        d_2278_v63_ = D9_DC22(len(_dafny.SeqWithoutIsStrInference([(d_2198_v7_)[default__.safeIndex((_dafny.MultiSet([d_2192_v0_])).cardinality, len(d_2198_v7_))] for d_2279_i4_ in range(default__.abs(735))])), d_2277_v62_)
        d_2280_v64_: _dafny.Seq
        d_2280_v64_ = _dafny.SeqWithoutIsStrInference([(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]])
        d_2281_v65_: _dafny.Seq
        d_2281_v65_ = _dafny.SeqWithoutIsStrInference([d_2278_v63_, D9_DC22(len(d_2280_v64_), d_2277_v62_), D9_DC22(p0, d_2277_v62_)])
        pat_let_tv41_ = d_2194_v3_
        pat_let_tv42_ = d_2194_v3_
        def lambda140_(source27_):
            if source27_.is_DC22:
                d_2282___mcc_h7_ = source27_.cf40
                d_2283___mcc_h8_ = source27_.cf41
                d_2284_cf41_ = d_2283___mcc_h8_
                d_2285_cf40_ = d_2282___mcc_h7_
                return pat_let_tv41_
            elif True:
                d_2286___mcc_h9_ = source27_.cf39
                d_2287_cf39_ = d_2286___mcc_h9_
                return pat_let_tv42_

        rhs437_ = ((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]) or ((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))])
        rhs438_ = d_2281_v65_
        rhs439_ = lambda140_(d_2278_v63_)
        rhs440_ = (d_2192_v0_) + (d_2192_v0_)
        rhs441_ = p0
        lhs343_ = globalState
        r3 = rhs437_
        d_2281_v65_ = rhs438_
        r0 = rhs439_
        d_2192_v0_ = rhs440_
        lhs343_.f9 = rhs441_
        if (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]:
            (globalState).f9 = p0
            index385_ = default__.safeIndex(795, (d_2205_v13_).length(0))
            (d_2205_v13_)[index385_] = (self).f14
            (globalState).f9 = (p0) + (len(d_2280_v64_))
            if (self).f13:
                index386_ = default__.safeIndex(853, (p1).length(0))
                (p1)[index386_] = default__.safeModuloInt(p0, p0)
                d_2288_v66_: _dafny.Set
                d_2288_v66_ = _dafny.Set({d_2192_v0_})
                index387_ = default__.safeIndex(853, (p1).length(0))
                (p1)[index387_] = len(d_2288_v66_)
                d_2289_v67_: _dafny.Array
                nw397_ = _dafny.Array(_dafny.CodePoint('D'), 2)
                d_2289_v67_ = nw397_
                d_2290_v68_: _dafny.Seq
                d_2290_v68_ = _dafny.SeqWithoutIsStrInference([d_2196_v5_, d_2196_v5_])
                d_2291_v69_: _dafny.Map
                d_2291_v69_ = _dafny.Map({len(_dafny.Map({d_2289_v67_: (p1)[default__.safeIndex(853, (p1).length(0))]})): len(d_2290_v68_)})
                d_2292_v70_: _dafny.Map
                d_2292_v70_ = _dafny.Map({p0: d_2200_v9_})
                d_2293_v71_: _dafny.Seq
                d_2293_v71_ = _dafny.SeqWithoutIsStrInference([d_2200_v9_])
                d_2291_v69_ = (d_2291_v69_).set((len(d_2198_v7_)) + ((((d_2292_v70_)[(d_2195_v4_).cardinality] if ((d_2195_v4_).cardinality) in (d_2292_v70_) else (d_2293_v71_)[default__.safeIndex((p1)[default__.safeIndex(853, (p1).length(0))], len(d_2293_v71_))])).cardinality), (p1)[default__.safeIndex(853, (p1).length(0))])
                index388_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index388_] = (self).f13
                (globalState).f8 = default__.fm1((d_2202_v11_).ispropersubset(d_2202_v11_), d_2194_v3_, globalState)
                d_2198_v7_ = d_2198_v7_
            elif True:
                rhs442_ = True
                rhs443_ = (d_2280_v64_)[default__.safeIndex(p0, len(d_2280_v64_))]
                rhs444_ = (True) in ((_dafny.MultiSet([True])).intersection(d_2195_v4_))
                rhs445_ = ((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))] if (self).f14 else (self).f13)
                lhs344_ = globalState
                lhs345_ = globalState
                lhs346_ = globalState
                lhs344_.f8 = rhs442_
                lhs345_.f8 = rhs443_
                r3 = rhs444_
                lhs346_.f8 = rhs445_
                index389_ = default__.safeIndex(611, (p1).length(0))
                (p1)[index389_] = len((d_2280_v64_ if (self).f14 else d_2280_v64_))
                d_2294_v72_: D2
                d_2294_v72_ = D2_DC7((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], ((d_2201_v10_)[len((d_2192_v0_).set(default__.safeIndex(p0, len(d_2192_v0_)), d_2194_v3_))] if (len((d_2192_v0_).set(default__.safeIndex(p0, len(d_2192_v0_)), d_2194_v3_))) in (d_2201_v10_) else p0), (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))], d_2192_v0_)
                index390_ = default__.safeIndex(611, (p1).length(0))
                (p1)[index390_] = (0) - (default__.safeDivisionInt(len((d_2294_v72_).cf14), (0) - ((0) - (p0))))
                index391_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index391_] = (self).f13
                d_2295_v73_: D7
                d_2295_v73_ = D7_DC17(d_2192_v0_, d_2205_v13_, (self).f13, (p1)[default__.safeIndex(611, (p1).length(0))])
                d_2295_v73_ = d_2295_v73_
                (globalState).f8 = not((d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))])
            (globalState).f9 = (0) - (len((default__.fm53(globalState)).intersection(d_2196_v5_)))
        elif True:
            if (self).f13:
                index392_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index392_] = (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]
                d_2195_v4_ = (d_2195_v4_) - (default__.fm46(globalState))
                d_2200_v9_ = d_2201_v10_
                index393_ = default__.safeIndex(434, (p1).length(0))
                (p1)[index393_] = p0
                index394_ = default__.safeIndex(434, (p1).length(0))
                (p1)[index394_] = -95
                d_2296_v74_: C6
                nw398_ = C6()
                nw398_.ctor__((d_2198_v7_)[default__.safeIndex(p0, len(d_2198_v7_))], ((d_2192_v0_) + (d_2192_v0_)).set(default__.safeIndex(p0, len((d_2192_v0_) + (d_2192_v0_))), d_2194_v3_))
                d_2296_v74_ = nw398_
            elif True:
                d_2297_v75_: _dafny.Array
                nw399_ = _dafny.Array(_dafny.MultiSet({}), 18)
                d_2297_v75_ = nw399_
                d_2297_v75_ = (d_2297_v75_ if False else d_2297_v75_)
                (globalState).f9 = -735
                (globalState).f9 = (self).fm11((self).f13, False, ((d_2195_v4_).intersection(d_2195_v4_)).cardinality, (self).f14, globalState)
                index395_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index395_] = (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]
                d_2298_v76_: _dafny.Array
                def lambda141_(d_2299_v0_):
                    def lambda142_(d_2300_i5_):
                        return (_dafny.Map({(self).f14: len(d_2299_v0_)})) | (_dafny.Map({(self).f13: len(d_2299_v0_)}))

                    return lambda142_

                init76_ = lambda141_(d_2192_v0_)
                nw400_ = _dafny.Array(None, 28)
                for i0_76_ in range(nw400_.length(0)):
                    nw400_[i0_76_] = init76_(i0_76_)
                d_2298_v76_ = nw400_
                d_2301_v77_: _dafny.Map
                d_2301_v77_ = _dafny.Map({(d_2280_v64_ if (self).f14 else d_2280_v64_): (d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]})
                index396_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                rhs446_ = d_2298_v76_
                rhs447_ = d_2205_v13_
                rhs448_ = _dafny.Map({(d_2280_v64_) + (d_2280_v64_): (self).f13})
                rhs449_ = (self).fm11((self).f14, True, (default__.fm0((self).f13, globalState)) + (p0), (self).f13, globalState)
                rhs450_ = (self).f14
                lhs347_ = globalState
                lhs348_ = d_2205_v13_
                lhs349_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                d_2298_v76_ = rhs446_
                d_2205_v13_ = rhs447_
                d_2301_v77_ = rhs448_
                lhs347_.f9 = rhs449_
                lhs348_[lhs349_] = rhs450_
            d_2302_v78_: C0
            nw401_ = C0()
            nw401_.ctor__((self).f14)
            d_2302_v78_ = nw401_
            index397_ = default__.safeIndex(79, (p1).length(0))
            (p1)[index397_] = p0
            index398_ = default__.safeIndex(79, (p1).length(0))
            (p1)[index398_] = p0
            d_2303_v79_: D8
            d_2303_v79_ = D8_DC20(d_2197_v6_, (self).f14, (self).f14)
            source28_ = d_2303_v79_
            if source28_.is_DC19:
                d_2304___mcc_h10_ = source28_.cf34
                d_2305___mcc_h11_ = source28_.cf35
                d_2306_cf35_ = d_2305___mcc_h11_
                d_2307_cf34_ = d_2304___mcc_h10_
                index399_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index399_] = (self).f13
                d_2308_v80_: C4
                nw402_ = C4()
                nw402_.ctor__((self).f11)
                d_2308_v80_ = nw402_
                index400_ = default__.safeIndex(79, (p1).length(0))
                (p1)[index400_] = len(((d_2198_v7_) + ((d_2198_v7_) + (d_2198_v7_))).set(default__.safeIndex(p0, len((d_2198_v7_) + ((d_2198_v7_) + (d_2198_v7_)))), p0))
                (globalState).f9 = (d_2278_v63_).cf40
            elif source28_.is_DC20:
                d_2309___mcc_h12_ = source28_.cf36
                d_2310___mcc_h13_ = source28_.cf37
                d_2311___mcc_h14_ = source28_.cf38
                d_2312_cf38_ = d_2311___mcc_h14_
                d_2313_cf37_ = d_2310___mcc_h13_
                d_2314_cf36_ = d_2309___mcc_h12_
                (globalState).f9 = (p1)[default__.safeIndex(79, (p1).length(0))]
                (globalState).f9 = 547
                d_2315_v81_: _dafny.Set
                d_2315_v81_ = _dafny.Set({d_2205_v13_})
                d_2316_v82_: _dafny.Array
                nw403_ = _dafny.Array(None, 3)
                nw403_[int(0)] = (p1)[default__.safeIndex(79, (p1).length(0))]
                nw403_[int(1)] = len(d_2315_v81_)
                nw403_[int(2)] = (p1)[default__.safeIndex(79, (p1).length(0))]
                d_2316_v82_ = nw403_
                d_2317_v83_: _dafny.Seq
                d_2317_v83_ = _dafny.SeqWithoutIsStrInference([d_2316_v82_])
                d_2317_v83_ = d_2317_v83_
                index401_ = default__.safeIndex(79, (p1).length(0))
                (p1)[index401_] = default__.safeModuloInt(p0, ((d_2195_v4_).cardinality) + ((p1)[default__.safeIndex(79, (p1).length(0))]))
            elif True:
                d_2318___mcc_h15_ = source28_.cf33
                d_2319_cf33_ = d_2318___mcc_h15_
                d_2320_v84_: _dafny.Map
                d_2320_v84_ = _dafny.Map({d_2198_v7_: True})
                d_2321_v85_: _dafny.Map
                d_2321_v85_ = _dafny.Map({(0) - ((p1)[default__.safeIndex(79, (p1).length(0))]): (default__.fm46(globalState)).cardinality})
                rhs451_ = (self).fm9(d_2193_v1_, d_2319_cf33_, globalState)
                rhs452_ = ((len((d_2320_v84_).set(d_2198_v7_, (self).f13))) * (p0)) not in (d_2321_v85_)
                lhs350_ = globalState
                lhs351_ = globalState
                lhs350_.f9 = rhs451_
                lhs351_.f8 = rhs452_
                (globalState).f8 = (len((d_2280_v64_) + (d_2280_v64_))) >= (default__.fm0((self).f13, globalState))
                d_2194_v3_ = d_2194_v3_
                index402_ = default__.safeIndex(795, (d_2205_v13_).length(0))
                (d_2205_v13_)[index402_] = (self).f14
            r3 = (self).f13
        r0 = d_2194_v3_
        d_2322_v86_: D5
        d_2322_v86_ = D5_DC14((self).f13)
        pat_let_tv43_ = d_2192_v0_
        pat_let_tv44_ = p0
        def lambda143_(source29_):
            if source29_.is_DC14:
                d_2323___mcc_h16_ = source29_.cf26
                d_2324_cf26_ = d_2323___mcc_h16_
                return len(pat_let_tv43_)
            elif True:
                d_2325___mcc_h17_ = source29_.cf25
                d_2326_cf25_ = d_2325___mcc_h17_
                return (0) - (pat_let_tv44_)

        def iife155_(_pat_let39_0):
            def iife156_(d_2327_dt__update__tmp_h1_):
                def iife157_(_pat_let40_0):
                    def iife158_(d_2328_dt__update_hcf26_h0_):
                        return D5_DC14(d_2328_dt__update_hcf26_h0_)
                    return iife158_(_pat_let40_0)
                return iife157_((self).f13)
            return iife156_(_pat_let39_0)
        r1 = lambda143_(iife155_(d_2322_v86_))
        d_2329_v87_: D4
        d_2329_v87_ = D4_DC11(p0, p0, (self).f14)
        d_2330_v88_: D2
        d_2330_v88_ = D2_DC8(p0, d_2194_v3_)
        r2 = (d_2330_v88_ if (d_2329_v87_).cf21 else d_2330_v88_)
        d_2331_v89_: _dafny.Map
        d_2331_v89_ = _dafny.Map({(d_2205_v13_)[default__.safeIndex(795, (d_2205_v13_).length(0))]: p0})
        r3 = (-251) == (((d_2331_v89_)[not((self).f13)] if (not((self).f13)) in (d_2331_v89_) else p0))
        return r0, r1, r2, r3

    def m5(self, globalState):
        r0: str = _dafny.CodePoint('D')
        r1: bool = False
        r2: bool = False
        if not (not((self).f13)) or (((self).f13) and ((self).f13)):
            default__.m0(globalState)
            d_2332_v0_: _dafny.Array
            nw404_ = _dafny.Array(int(0), 16)
            d_2332_v0_ = nw404_
            index403_ = default__.safeIndex(207, (d_2332_v0_).length(0))
            (d_2332_v0_)[index403_] = 740
            d_2333_v1_: int
            d_2333_v1_ = 686
            d_2334_v2_: _dafny.Map
            d_2334_v2_ = _dafny.Map({d_2333_v1_: _dafny.SeqWithoutIsStrInference([(self).f13, (self).f13, (self).f13])})
            d_2335_v3_: _dafny.Seq
            d_2335_v3_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "y"))
            d_2336_v4_: _dafny.MultiSet
            d_2336_v4_ = _dafny.MultiSet([d_2335_v3_, _dafny.SeqWithoutIsStrInference([(D13_DC30(_dafny.CodePoint('q'))).cf51 for d_2337_i0_ in range(default__.abs(998))])])
            d_2338_v5_: _dafny.Seq
            d_2338_v5_ = _dafny.SeqWithoutIsStrInference([(462) - ((0) - (d_2333_v1_))])
            d_2339_v6_: _dafny.Set
            d_2339_v6_ = _dafny.Set({(self).f13, (self).f14, (self).f13, (self).f14})
            index404_ = default__.safeIndex(207, (d_2332_v0_).length(0))
            rhs453_ = (d_2338_v5_)[default__.safeIndex(default__.safeModuloInt(d_2333_v1_, d_2333_v1_), len(d_2338_v5_))]
            rhs454_ = (self).f14
            rhs455_ = default__.fm61(globalState)
            rhs456_ = ((d_2339_v6_).intersection(d_2339_v6_)).issubset(d_2339_v6_)
            rhs457_ = d_2336_v4_
            lhs352_ = d_2332_v0_
            lhs353_ = default__.safeIndex(207, (d_2332_v0_).length(0))
            lhs352_[lhs353_] = rhs453_
            r1 = rhs454_
            d_2334_v2_ = rhs455_
            r2 = rhs456_
            d_2336_v4_ = rhs457_
            d_2340_v7_: _dafny.Array
            def lambda144_(d_2341_i1_):
                return not((self).f13)

            init77_ = lambda144_
            nw405_ = _dafny.Array(None, 25)
            for i0_77_ in range(nw405_.length(0)):
                nw405_[i0_77_] = init77_(i0_77_)
            d_2340_v7_ = nw405_
            d_2342_v8_: D7
            d_2342_v8_ = D7_DC17(d_2335_v3_, d_2340_v7_, (self).f13, d_2333_v1_)
            if (((d_2342_v8_).cf29) + (_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcasst")))) < (d_2335_v3_):
                (globalState).f9 = d_2333_v1_
                r1 = (self).f14
                index405_ = default__.safeIndex(207, (d_2332_v0_).length(0))
                (d_2332_v0_)[index405_] = default__.fm0((self).fm10(globalState), globalState)
                d_2343_v9_: _dafny.Map
                d_2343_v9_ = _dafny.Map({(self).f13: (self).f11})
                d_2344_v10_: C5
                nw406_ = C5()
                nw406_.ctor__(((d_2343_v9_)[(self).f14] if ((self).f14) in (d_2343_v9_) else (self).f11))
                d_2344_v10_ = nw406_
                d_2345_v11_: _dafny.MultiSet
                d_2345_v11_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([d_2333_v1_])])
                index406_ = default__.safeIndex(207, (d_2332_v0_).length(0))
                (d_2332_v0_)[index406_] = ((d_2345_v11_)[_dafny.SeqWithoutIsStrInference([(d_2332_v0_)[default__.safeIndex(207, (d_2332_v0_).length(0))], d_2333_v1_])] if (_dafny.SeqWithoutIsStrInference([(d_2332_v0_)[default__.safeIndex(207, (d_2332_v0_).length(0))], d_2333_v1_])) in (d_2345_v11_) else default__.fm0((self).f14, globalState))
            elif True:
                d_2346_v12_: str
                d_2346_v12_ = _dafny.CodePoint('s')
                r0 = d_2346_v12_
                d_2347_v13_: _dafny.Map
                d_2347_v13_ = _dafny.Map({(d_2338_v5_).set(default__.safeIndex((d_2332_v0_)[default__.safeIndex(207, (d_2332_v0_).length(0))], len(d_2338_v5_)), (d_2332_v0_)[default__.safeIndex(207, (d_2332_v0_).length(0))]): d_2333_v1_})
                d_2347_v13_ = (d_2347_v13_).set(d_2338_v5_, (d_2332_v0_)[default__.safeIndex(207, (d_2332_v0_).length(0))])
                r0 = d_2346_v12_
                rhs458_ = ((self).f13 if (self).f14 else False)
                rhs459_ = (self).f13
                rhs460_ = (self).f13
                rhs461_ = default__.safeModuloInt(464, d_2333_v1_)
                lhs354_ = globalState
                r1 = rhs458_
                r2 = rhs459_
                r2 = rhs460_
                lhs354_.f9 = rhs461_
                d_2348_v14_: _dafny.MultiSet
                d_2348_v14_ = _dafny.MultiSet([(self).f13, (self).f14, (self).f13, (self).f13, (self).f14])
                d_2349_v15_: _dafny.Set
                d_2349_v15_ = _dafny.Set({(0) - ((d_2348_v14_).cardinality)})
                d_2350_v16_: C3
                nw407_ = C3()
                nw407_.ctor__((self).f13, (d_2349_v15_).ispropersubset(d_2349_v15_), (self).f11)
                d_2350_v16_ = nw407_
            d_2351_v17_: _dafny.Map
            d_2351_v17_ = _dafny.Map({(self).f13: not(not((self).f13))})
            d_2352_v18_: str
            d_2352_v18_ = _dafny.CodePoint('t')
            d_2353_v19_: T2
            nw408_ = C1()
            nw408_.ctor__(((d_2351_v17_)[(self).f13] if ((self).f13) in (d_2351_v17_) else (self).f14), (self).f14, default__.fm40(d_2352_v18_, globalState), (self).f14)
            d_2353_v19_ = nw408_
            d_2353_v19_ = d_2353_v19_
            d_2354_v20_: _dafny.MultiSet
            d_2354_v20_ = _dafny.MultiSet([False, (self).f14, (d_2353_v19_).f13, (d_2353_v19_).f13, (self).f13])
            d_2355_v21_: _dafny.Seq
            d_2355_v21_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "yphid")), d_2335_v3_])
            d_2356_v22_: D0
            d_2356_v22_ = D0_DC2(d_2354_v20_, (self).f14, len((d_2355_v21_)[default__.safeIndex(821, len(d_2355_v21_))]), (self).f13)
            r1 = (d_2356_v22_).cf5
        elif True:
            d_2357_v23_: T1
            nw409_ = C5()
            nw409_.ctor__((self).f11)
            d_2357_v23_ = nw409_
            d_2358_v24_: int
            d_2358_v24_ = 938
            d_2359_v25_: _dafny.Seq
            d_2359_v25_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oleu"))
            d_2360_v26_: str
            d_2360_v26_ = _dafny.CodePoint('h')
            d_2361_v27_: D2
            d_2361_v27_ = D2_DC7(((_dafny.MultiSet([d_2357_v23_])).cardinality) <= (d_2358_v24_), (d_2358_v24_) + (len(_dafny.SeqWithoutIsStrInference([(self).f13]))), not((self).f14), (d_2359_v25_).set(default__.safeIndex((0) - (len(d_2359_v25_)), len(d_2359_v25_)), d_2360_v26_))
            source30_ = d_2361_v27_
            if source30_.is_DC6:
                (globalState).f9 = d_2358_v24_
                d_2362_v28_: _dafny.Array
                nw410_ = _dafny.Array(False, 16)
                d_2362_v28_ = nw410_
                d_2363_v29_: _dafny.Seq
                d_2363_v29_ = _dafny.SeqWithoutIsStrInference([d_2358_v24_])
                d_2364_v30_: _dafny.Map
                d_2364_v30_ = _dafny.Map({d_2362_v28_: (d_2363_v29_).set(default__.safeIndex(753, len(d_2363_v29_)), len(d_2363_v29_))})
                d_2365_v31_: _dafny.MultiSet
                d_2365_v31_ = _dafny.MultiSet([(self).f13, (self).f13])
                d_2366_v32_: _dafny.Map
                d_2366_v32_ = _dafny.Map({(self).f13: d_2358_v24_})
                d_2367_v33_: _dafny.Seq
                d_2367_v33_ = _dafny.SeqWithoutIsStrInference([d_2359_v25_])
                d_2368_v34_: _dafny.Array
                nw411_ = _dafny.Array(None, 27)
                nw411_[int(0)] = d_2358_v24_
                nw411_[int(1)] = default__.safeDivisionInt(d_2358_v24_, d_2358_v24_)
                nw411_[int(2)] = default__.safeModuloInt(d_2358_v24_, d_2358_v24_)
                nw411_[int(3)] = default__.safeModuloInt((d_2363_v29_)[default__.safeIndex(d_2358_v24_, len(d_2363_v29_))], d_2358_v24_)
                nw411_[int(4)] = d_2358_v24_
                nw411_[int(5)] = d_2358_v24_
                nw411_[int(6)] = (len(_dafny.SeqWithoutIsStrInference([d_2358_v24_, len(_dafny.Map({(self).f13: d_2358_v24_})), (d_2365_v31_).cardinality, 359, d_2358_v24_]))) * ((_dafny.MultiSet(d_2363_v29_)).cardinality)
                nw411_[int(7)] = len(d_2366_v32_)
                nw411_[int(8)] = (((d_2366_v32_)[(self).f14] if ((self).f14) in (d_2366_v32_) else d_2358_v24_)) - (d_2358_v24_)
                nw411_[int(9)] = (d_2358_v24_) - (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "rh"))))
                nw411_[int(10)] = 334
                nw411_[int(11)] = d_2358_v24_
                nw411_[int(12)] = (len(d_2363_v29_) if (self).f14 else d_2358_v24_)
                nw411_[int(13)] = d_2358_v24_
                nw411_[int(14)] = (d_2365_v31_).cardinality
                nw411_[int(15)] = d_2358_v24_
                nw411_[int(16)] = (d_2358_v24_ if (self).f14 else (0) - (d_2358_v24_))
                nw411_[int(17)] = d_2358_v24_
                nw411_[int(18)] = default__.fm0((self).f13, globalState)
                nw411_[int(19)] = d_2358_v24_
                nw411_[int(20)] = 514
                nw411_[int(21)] = 278
                nw411_[int(22)] = (d_2363_v29_)[default__.safeIndex((0) - (d_2358_v24_), len(d_2363_v29_))]
                nw411_[int(23)] = len(d_2367_v33_)
                nw411_[int(24)] = d_2358_v24_
                nw411_[int(25)] = d_2358_v24_
                nw411_[int(26)] = len(d_2363_v29_)
                d_2368_v34_ = nw411_
                index407_ = default__.safeIndex(259, (d_2368_v34_).length(0))
                (d_2368_v34_)[index407_] = 909
                d_2369_v35_: _dafny.Seq
                d_2369_v35_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f14: 962})])
                index408_ = default__.safeIndex(259, (d_2368_v34_).length(0))
                rhs462_ = d_2360_v26_
                rhs463_ = (d_2364_v30_) | (d_2364_v30_)
                rhs464_ = len(((d_2369_v35_).set(default__.safeIndex(d_2358_v24_, len(d_2369_v35_)), d_2366_v32_)) + (d_2369_v35_))
                rhs465_ = default__.fm1(not((self).f14), d_2360_v26_, globalState)
                lhs355_ = d_2368_v34_
                lhs356_ = default__.safeIndex(259, (d_2368_v34_).length(0))
                lhs357_ = globalState
                r0 = rhs462_
                d_2364_v30_ = rhs463_
                lhs355_[lhs356_] = rhs464_
                lhs357_.f8 = rhs465_
                d_2370_v36_: _dafny.Seq
                d_2370_v36_ = _dafny.SeqWithoutIsStrInference([not(default__.fm1((self).f13, d_2360_v26_, globalState))])
                d_2370_v36_ = (d_2370_v36_) + (d_2370_v36_)
                d_2371_v37_: C0
                nw412_ = C0()
                nw412_.ctor__((d_2359_v25_) != (d_2359_v25_))
                d_2371_v37_ = nw412_
            elif source30_.is_DC7:
                d_2372___mcc_h0_ = source30_.cf11
                d_2373___mcc_h1_ = source30_.cf12
                d_2374___mcc_h2_ = source30_.cf13
                d_2375___mcc_h3_ = source30_.cf14
                d_2376_cf14_ = d_2375___mcc_h3_
                d_2377_cf13_ = d_2374___mcc_h2_
                d_2378_cf12_ = d_2373___mcc_h1_
                d_2379_cf11_ = d_2372___mcc_h0_
                d_2380_v38_: _dafny.Array
                nw413_ = _dafny.Array(None, 24)
                nw413_[int(0)] = (0) - (d_2358_v24_)
                nw413_[int(1)] = d_2358_v24_
                nw413_[int(2)] = d_2358_v24_
                nw413_[int(3)] = d_2378_cf12_
                nw413_[int(4)] = d_2378_cf12_
                nw413_[int(5)] = d_2358_v24_
                nw413_[int(6)] = d_2358_v24_
                nw413_[int(7)] = d_2358_v24_
                nw413_[int(8)] = d_2378_cf12_
                nw413_[int(9)] = d_2378_cf12_
                nw413_[int(10)] = -98
                nw413_[int(11)] = d_2378_cf12_
                nw413_[int(12)] = 288
                nw413_[int(13)] = d_2358_v24_
                nw413_[int(14)] = d_2358_v24_
                nw413_[int(15)] = d_2358_v24_
                nw413_[int(16)] = len(_dafny.SeqWithoutIsStrInference([_dafny.Map({True: d_2378_cf12_}) for d_2381_i2_ in range(default__.abs(937))]))
                nw413_[int(17)] = d_2378_cf12_
                nw413_[int(18)] = (0) - (d_2378_cf12_)
                nw413_[int(19)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "tnrgo")))
                nw413_[int(20)] = (0) - (d_2378_cf12_)
                nw413_[int(21)] = default__.fm0((self).f14, globalState)
                nw413_[int(22)] = d_2378_cf12_
                nw413_[int(23)] = len(_dafny.SeqWithoutIsStrInference([d_2360_v26_ for d_2382_i3_ in range(default__.abs(425))]))
                d_2380_v38_ = nw413_
                d_2383_v39_: str
                d_2384_v40_: int
                d_2385_v41_: D2
                d_2386_v42_: bool
                out15_: str
                out16_: int
                out17_: D2
                out18_: bool
                out15_, out16_, out17_, out18_ = (d_2357_v23_).m4(d_2358_v24_, d_2380_v38_, globalState)
                d_2383_v39_ = out15_
                d_2384_v40_ = out16_
                d_2385_v41_ = out17_
                d_2386_v42_ = out18_
                d_2387_v43_: _dafny.Seq
                d_2387_v43_ = _dafny.SeqWithoutIsStrInference([True])
                d_2388_v44_: _dafny.Seq
                d_2388_v44_ = _dafny.SeqWithoutIsStrInference([default__.fm1(not((d_2387_v43_)[default__.safeIndex(d_2358_v24_, len(d_2387_v43_))]), d_2360_v26_, globalState)])
                d_2386_v42_ = ((self).f13) or ((d_2388_v44_) < (d_2388_v44_))
                d_2359_v25_ = d_2376_cf14_
                d_2379_cf11_ = ((0) - ((d_2378_cf12_) + (897))) > ((340) + (d_2384_v40_))
            elif source30_.is_DC8:
                d_2389___mcc_h4_ = source30_.cf15
                d_2390___mcc_h5_ = source30_.cf16
                d_2391_cf16_ = d_2390___mcc_h5_
                d_2392_cf15_ = d_2389___mcc_h4_
                d_2393_v45_: _dafny.Array
                nw414_ = _dafny.Array(int(0), 17)
                d_2393_v45_ = nw414_
                d_2394_v46_: _dafny.MultiSet
                d_2394_v46_ = _dafny.MultiSet([d_2392_cf15_, d_2392_cf15_])
                d_2395_v47_: _dafny.Map
                d_2395_v47_ = _dafny.Map({(self).f13: d_2394_v46_})
                index409_ = default__.safeIndex(203, (d_2393_v45_).length(0))
                (d_2393_v45_)[index409_] = (0) - (len(d_2395_v47_))
                d_2396_v48_: _dafny.Seq
                d_2396_v48_ = _dafny.SeqWithoutIsStrInference([(self).f14])
                d_2397_v49_: _dafny.Seq
                d_2397_v49_ = _dafny.SeqWithoutIsStrInference([_dafny.SeqWithoutIsStrInference([(self).f14]), (d_2396_v48_).set(default__.safeIndex(default__.fm0((self).f14, globalState), len(d_2396_v48_)), (self).f13)])
                index410_ = default__.safeIndex(203, (d_2393_v45_).length(0))
                (d_2393_v45_)[index410_] = len((default__.fm44(d_2397_v49_, d_2360_v26_, globalState)) + ((d_2396_v48_) + (d_2396_v48_)))
                (globalState).f8 = (self).f13
                d_2398_v50_: C0
                nw415_ = C0()
                nw415_.ctor__((self).f13)
                d_2398_v50_ = nw415_
                d_2399_v51_: _dafny.Array
                nw416_ = _dafny.Array(_dafny.Map({}), 1)
                d_2399_v51_ = nw416_
                index411_ = default__.safeIndex(395, (d_2399_v51_).length(0))
                (d_2399_v51_)[index411_] = default__.fm50(d_2391_cf16_, (self).f13, not((self).f13), globalState)
                d_2400_v52_: D4
                d_2400_v52_ = D4_DC11(d_2392_cf15_, d_2392_cf15_, default__.fm1((d_2398_v50_).f17, d_2360_v26_, globalState))
                d_2401_v53_: _dafny.Map
                d_2401_v53_ = _dafny.Map({(len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "onjpllpsl")))) < (614): (d_2400_v52_).cf19})
                index412_ = default__.safeIndex(395, (d_2399_v51_).length(0))
                (d_2399_v51_)[index412_] = d_2401_v53_
            elif True:
                d_2402___mcc_h6_ = source30_.cf10
                d_2403_cf10_ = d_2402___mcc_h6_
                (globalState).f9 = default__.safeDivisionInt(d_2358_v24_, d_2358_v24_)
                d_2404_v54_: _dafny.Array
                def lambda145_(d_2405_i4_):
                    return (self).f14

                init78_ = lambda145_
                nw417_ = _dafny.Array(None, 17)
                for i0_78_ in range(nw417_.length(0)):
                    nw417_[i0_78_] = init78_(i0_78_)
                d_2404_v54_ = nw417_
                index413_ = default__.safeIndex(833, (d_2404_v54_).length(0))
                (d_2404_v54_)[index413_] = (self).f13
                d_2406_v55_: _dafny.Array
                nw418_ = _dafny.Array(_dafny.Array(None, 0), 19)
                d_2406_v55_ = nw418_
                d_2407_v56_: _dafny.Map
                d_2407_v56_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(self).f14]): True})
                d_2408_v57_: D2
                d_2408_v57_ = D2_DC8(len(d_2407_v56_), d_2360_v26_)
                d_2409_v58_: _dafny.Array
                nw419_ = _dafny.Array(None, 7)
                nw419_[int(0)] = d_2360_v26_
                nw419_[int(1)] = (d_2408_v57_).cf16
                nw419_[int(2)] = d_2360_v26_
                nw419_[int(3)] = d_2360_v26_
                nw419_[int(4)] = d_2360_v26_
                nw419_[int(5)] = d_2360_v26_
                nw419_[int(6)] = d_2360_v26_
                d_2409_v58_ = nw419_
                index414_ = default__.safeIndex(509, (d_2406_v55_).length(0))
                (d_2406_v55_)[index414_] = d_2409_v58_
                index415_ = default__.safeIndex(833, (d_2404_v54_).length(0))
                index416_ = default__.safeIndex(509, (d_2406_v55_).length(0))
                rhs466_ = (self).fm10(globalState)
                rhs467_ = d_2409_v58_
                rhs468_ = (self).f13
                lhs358_ = d_2404_v54_
                lhs359_ = default__.safeIndex(833, (d_2404_v54_).length(0))
                lhs360_ = d_2406_v55_
                lhs361_ = default__.safeIndex(509, (d_2406_v55_).length(0))
                lhs362_ = globalState
                lhs358_[lhs359_] = rhs466_
                lhs360_[lhs361_] = rhs467_
                lhs362_.f8 = rhs468_
                d_2410_v59_: _dafny.Map
                d_2410_v59_ = _dafny.Map({d_2358_v24_: d_2358_v24_})
                (globalState).f9 = len(d_2410_v59_)
                (globalState).f9 = default__.safeModuloInt(d_2358_v24_, len(default__.fm62((self).f14, globalState)))
            d_2411_v60_: C5
            nw420_ = C5()
            nw420_.ctor__((self).f11)
            d_2411_v60_ = nw420_
            d_2412_v61_: int
            d_2413_v62_: _dafny.Set
            out19_: int
            out20_: _dafny.Set
            out19_, out20_ = (d_2357_v23_).m3((self).f14, globalState)
            d_2412_v61_ = out19_
            d_2413_v62_ = out20_
            d_2412_v61_ = d_2358_v24_
            d_2414_v63_: _dafny.Array
            nw421_ = _dafny.Array(False, 8)
            d_2414_v63_ = nw421_
            index417_ = default__.safeIndex(296, (d_2414_v63_).length(0))
            (d_2414_v63_)[index417_] = (self).f13
            index418_ = default__.safeIndex(296, (d_2414_v63_).length(0))
            (d_2414_v63_)[index418_] = default__.fm1((self).f14, d_2360_v26_, globalState)
        d_2415_v64_: int
        d_2415_v64_ = 475
        rhs469_ = d_2415_v64_
        rhs470_ = d_2415_v64_
        lhs363_ = globalState
        lhs364_ = globalState
        lhs363_.f9 = rhs469_
        lhs364_.f9 = rhs470_
        d_2416_v65_: _dafny.Seq
        d_2416_v65_ = _dafny.SeqWithoutIsStrInference([d_2415_v64_])
        d_2417_v66_: _dafny.Array
        def lambda146_(d_2418_i6_):
            return False

        init79_ = lambda146_
        nw422_ = _dafny.Array(None, 6)
        for i0_79_ in range(nw422_.length(0)):
            nw422_[i0_79_] = init79_(i0_79_)
        d_2417_v66_ = nw422_
        d_2419_v67_: _dafny.Map
        d_2419_v67_ = _dafny.Map({d_2416_v65_: d_2417_v66_})
        hi18_ = d_2415_v64_
        for d_2420_i5_ in range(len((d_2419_v67_).set(d_2416_v65_, d_2417_v66_)), hi18_):
            index419_ = default__.safeIndex(191, (d_2417_v66_).length(0))
            (d_2417_v66_)[index419_] = (d_2420_i5_) > (d_2415_v64_)
            d_2421_v68_: _dafny.Seq
            d_2421_v68_ = _dafny.SeqWithoutIsStrInference([(self).f14])
            d_2422_v69_: D0
            d_2422_v69_ = D0_DC2(_dafny.MultiSet(d_2421_v68_), (self).f13, d_2415_v64_, (self).f14)
            d_2423_v70_: _dafny.Set
            d_2423_v70_ = _dafny.Set({d_2415_v64_})
            d_2424_v71_: _dafny.Seq
            d_2424_v71_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "vbrdtah"))
            d_2425_v72_: _dafny.Array
            nw423_ = _dafny.Array(None, 23)
            nw423_[int(0)] = d_2415_v64_
            nw423_[int(1)] = d_2415_v64_
            nw423_[int(2)] = d_2415_v64_
            nw423_[int(3)] = default__.safeDivisionInt(d_2415_v64_, d_2420_i5_)
            nw423_[int(4)] = d_2415_v64_
            nw423_[int(5)] = (d_2420_i5_) * (d_2420_i5_)
            nw423_[int(6)] = d_2420_i5_
            nw423_[int(7)] = d_2420_i5_
            nw423_[int(8)] = d_2420_i5_
            nw423_[int(9)] = (208) - (d_2415_v64_)
            nw423_[int(10)] = d_2415_v64_
            def iife159_(_pat_let41_0):
                def iife160_(d_2426_dt__update__tmp_h0_):
                    def iife161_(_pat_let42_0):
                        def iife162_(d_2427_dt__update_hcf6_h0_):
                            return D0_DC2((d_2426_dt__update__tmp_h0_).cf4, (d_2426_dt__update__tmp_h0_).cf5, d_2427_dt__update_hcf6_h0_, (d_2426_dt__update__tmp_h0_).cf7)
                        return iife162_(_pat_let42_0)
                    return iife161_(d_2420_i5_)
                return iife160_(_pat_let41_0)
            nw423_[int(11)] = default__.safeDivisionInt((iife159_(d_2422_v69_)).cf6, 183)
            nw423_[int(12)] = 471
            nw423_[int(13)] = default__.safeDivisionInt(d_2420_i5_, d_2415_v64_)
            nw423_[int(14)] = d_2415_v64_
            nw423_[int(15)] = (d_2420_i5_) - (d_2415_v64_)
            nw423_[int(16)] = (d_2415_v64_ if (self).f14 else d_2420_i5_)
            nw423_[int(17)] = default__.safeModuloInt(d_2415_v64_, d_2420_i5_)
            nw423_[int(18)] = d_2415_v64_
            nw423_[int(19)] = default__.safeDivisionInt(946, len(d_2421_v68_))
            nw423_[int(20)] = d_2420_i5_
            nw423_[int(21)] = len(d_2423_v70_)
            nw423_[int(22)] = len(d_2424_v71_)
            d_2425_v72_ = nw423_
            index420_ = default__.safeIndex(615, (d_2425_v72_).length(0))
            (d_2425_v72_)[index420_] = d_2420_i5_
            d_2428_v73_: _dafny.MultiSet
            d_2428_v73_ = _dafny.MultiSet([_dafny.SeqWithoutIsStrInference([(d_2424_v71_)[default__.safeIndex(len(d_2421_v68_), len(d_2424_v71_))] for d_2429_i7_ in range(default__.abs(987))])])
            index421_ = default__.safeIndex(191, (d_2417_v66_).length(0))
            index422_ = default__.safeIndex(615, (d_2425_v72_).length(0))
            rhs471_ = (d_2415_v64_ if ((0) - (d_2420_i5_)) != (-11) else default__.fm0((self).f14, globalState))
            rhs472_ = (self).f13
            rhs473_ = (d_2420_i5_ if (d_2423_v70_).ispropersubset(d_2423_v70_) else ((d_2428_v73_)[d_2424_v71_] if (d_2424_v71_) in (d_2428_v73_) else default__.fm0((self).f13, globalState)))
            rhs474_ = (880) - (d_2415_v64_)
            lhs365_ = globalState
            lhs366_ = d_2417_v66_
            lhs367_ = default__.safeIndex(191, (d_2417_v66_).length(0))
            lhs368_ = globalState
            lhs369_ = d_2425_v72_
            lhs370_ = default__.safeIndex(615, (d_2425_v72_).length(0))
            lhs365_.f9 = rhs471_
            lhs366_[lhs367_] = rhs472_
            lhs368_.f9 = rhs473_
            lhs369_[lhs370_] = rhs474_
            (globalState).f9 = d_2420_i5_
            d_2430_v74_: C3
            nw424_ = C3()
            nw424_.ctor__((self).f13, (d_2417_v66_)[default__.safeIndex(191, (d_2417_v66_).length(0))], (self).f11)
            d_2430_v74_ = nw424_
            (globalState).f8 = (d_2430_v74_).f22
        (globalState).f8 = (self).f13
        d_2431_v75_: _dafny.Seq
        d_2431_v75_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "ucbqax"))
        d_2432_v76_: _dafny.Seq
        d_2432_v76_ = _dafny.SeqWithoutIsStrInference([d_2431_v75_])
        d_2433_v77_: _dafny.Map
        d_2433_v77_ = _dafny.Map({(d_2432_v76_)[default__.safeIndex(len(d_2431_v75_), len(d_2432_v76_))]: (self).f13})
        d_2434_v78_: C3
        nw425_ = C3()
        nw425_.ctor__((self).f14, not(((d_2433_v77_)[d_2431_v75_] if (d_2431_v75_) in (d_2433_v77_) else (self).f13)), (self).f11)
        d_2434_v78_ = nw425_
        d_2435_v79_: _dafny.Array
        nw426_ = _dafny.Array(int(0), 15)
        d_2435_v79_ = nw426_
        index423_ = default__.safeIndex(522, (d_2435_v79_).length(0))
        (d_2435_v79_)[index423_] = 813
        index424_ = default__.safeIndex(522, (d_2435_v79_).length(0))
        (d_2435_v79_)[index424_] = (d_2415_v64_) * (d_2415_v64_)
        r0 = _dafny.CodePoint('e')
        r1 = (self).f14
        r2 = (d_2415_v64_) == ((d_2415_v64_) * (d_2415_v64_))
        return r0, r1, r2

    def m3(self, p0, globalState):
        r0: int = int(0)
        r1: _dafny.Set = _dafny.Set({})
        d_2436_v0_: int
        d_2436_v0_ = 989
        d_2437_v1_: _dafny.MultiSet
        d_2437_v1_ = _dafny.MultiSet([137])
        d_2438_v2_: _dafny.MultiSet
        d_2438_v2_ = _dafny.MultiSet([(d_2437_v1_).cardinality])
        d_2439_v3_: _dafny.MultiSet
        d_2439_v3_ = _dafny.MultiSet([_dafny.MultiSet([d_2436_v0_]), d_2438_v2_, d_2437_v1_])
        d_2440_v4_: C3
        nw427_ = C3()
        nw427_.ctor__((d_2439_v3_).ispropersubset((d_2439_v3_).set(d_2437_v1_, default__.abs(d_2436_v0_))), p0, (self).f11)
        d_2440_v4_ = nw427_
        hi19_ = d_2436_v0_
        for d_2441_i0_ in range(default__.fm0((self).f13, globalState), hi19_):
            (globalState).f9 = ((0) - (d_2436_v0_) if False else d_2436_v0_)
            r1 = _dafny.Set({False})
            (globalState).f8 = d_2440_v4_.f21
            d_2442_v5_: C5
            nw428_ = C5()
            nw428_.ctor__((self).f11)
            d_2442_v5_ = nw428_
        d_2443_v6_: _dafny.Map
        d_2443_v6_ = _dafny.Map({(d_2440_v4_).f22: (d_2436_v0_) - (d_2436_v0_)})
        d_2443_v6_ = (d_2443_v6_).set((self).f13, 724)
        d_2444_v7_: C2
        nw429_ = C2()
        nw429_.ctor__(((self).f11 if (self).f13 else (self).f11))
        d_2444_v7_ = nw429_
        d_2445_i1_: int
        d_2445_i1_ = 0
        with _dafny.label("11"):
            while (self).f13:
                with _dafny.c_label("11"):
                    if (d_2445_i1_) >= (100):
                        raise _dafny.Break("11")
                    d_2445_i1_ = (d_2445_i1_) + (1)
                    d_2446_v8_: _dafny.Array
                    def lambda147_(d_2447_v0_):
                        def lambda148_(d_2448_i2_):
                            return (_dafny.SeqWithoutIsStrInference([d_2447_v0_])) < (_dafny.SeqWithoutIsStrInference([d_2447_v0_, d_2447_v0_, 255, d_2447_v0_, d_2447_v0_]))

                        return lambda148_

                    init80_ = lambda147_(d_2436_v0_)
                    nw430_ = _dafny.Array(None, 18)
                    for i0_80_ in range(nw430_.length(0)):
                        nw430_[i0_80_] = init80_(i0_80_)
                    d_2446_v8_ = nw430_
                    index425_ = default__.safeIndex(629, (d_2446_v8_).length(0))
                    (d_2446_v8_)[index425_] = not ((self).f13) or (d_2440_v4_.f21)
                    index426_ = default__.safeIndex(629, (d_2446_v8_).length(0))
                    (d_2446_v8_)[index426_] = (self).f14
                    d_2449_v9_: _dafny.Map
                    d_2449_v9_ = _dafny.Map({False: d_2440_v4_.f21})
                    d_2436_v0_ = (len((_dafny.Map({(d_2440_v4_).f22: True})) | (d_2449_v9_)) if (d_2436_v0_) == (d_2436_v0_) else d_2436_v0_)
                    d_2450_v10_: _dafny.Array
                    nw431_ = _dafny.Array(_dafny.Seq({}), 9)
                    d_2450_v10_ = nw431_
                    d_2450_v10_ = d_2450_v10_
                    d_2451_v11_: _dafny.Map
                    d_2451_v11_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([d_2436_v0_ for d_2452_i3_ in range(default__.abs(325))]): (d_2436_v0_) - (d_2436_v0_)})
                    d_2453_v12_: _dafny.Seq
                    d_2453_v12_ = _dafny.SeqWithoutIsStrInference([d_2436_v0_, d_2436_v0_])
                    d_2451_v11_ = (d_2451_v11_).set((d_2453_v12_) + (d_2453_v12_), d_2436_v0_)
                    pass
            pass
        d_2454_v13_: _dafny.Array
        nw432_ = _dafny.Array(_dafny.Seq({}), 10)
        d_2454_v13_ = nw432_
        nw433_ = _dafny.Array(_dafny.Seq({}), 8)
        d_2454_v13_ = nw433_
        d_2455_v14_: str
        d_2455_v14_ = _dafny.CodePoint('l')
        d_2456_v16_: _dafny.Set
        def iife163_():
            coll77_ = _dafny.Set()
            compr_77_: str
            for compr_77_ in ((_dafny.Map({_dafny.CodePoint('s'): p0})).set(d_2455_v14_, (self).f13)).keys.Elements:
                d_2457_v15_: str = compr_77_
                if (d_2457_v15_) in ((_dafny.Map({_dafny.CodePoint('s'): p0})).set(d_2455_v14_, (self).f13)):
                    coll77_ = coll77_.union(_dafny.Set([d_2457_v15_]))
            return _dafny.Set(coll77_)
        d_2456_v16_ = iife163_()
        
        pat_let_tv45_ = d_2436_v0_
        def lambda149_(source31_):
            d_2458___mcc_h0_ = source31_
            d_2459_cf27_ = d_2458___mcc_h0_
            return (D8_DC19(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "imiictkf")), pat_let_tv45_)).cf35

        r0 = lambda149_(d_2456_v16_)
        d_2460_v17_: _dafny.Set
        d_2460_v17_ = _dafny.Set({(self).f13, default__.fm1(d_2440_v4_.f21, d_2455_v14_, globalState)})
        r1 = d_2460_v17_
        return r0, r1

    @property
    def f14(self):
        return self._f14

class C8:
    def  __init__(self):
        self._f10: bool = False
        pass

    def __dafnystr__(self) -> str:
        return "_module.C8"
    def ctor__(self, f10):
        (self)._f10 = f10

    def fm5(self, p0, p1, p2, globalState):
        return D0_DC2(_dafny.MultiSet([False, (self).f10]), (self).f10, (len(_dafny.Map({(self).f10: (_dafny.MultiSet([False])).cardinality}))) - (514), (self).f10)

    def fm6(self, globalState):
        def iife164_():
            def iife165_():
                coll79_ = _dafny.Map()
                compr_79_: int
                for compr_79_ in (_dafny.Map({-384: (self).f10})).keys.Elements:
                    d_2462_v1_: int = compr_79_
                    if (d_2462_v1_) in (_dafny.Map({-384: (self).f10})):
                        coll79_[default__.safeDivisionInt(d_2462_v1_, len(_dafny.Map({len(_dafny.Map({881: (_dafny.MultiSet([(self).f10])).cardinality})): len(_dafny.SeqWithoutIsStrInference([len(_dafny.Map({(self).f10: _dafny.CodePoint('q')})), len(_dafny.SeqWithoutIsStrInference([-562, -528]))]))})))] = len(_dafny.SeqWithoutIsStrInference([(self).f10]))
                return _dafny.Map(coll79_)
            coll78_ = _dafny.Map()
            compr_78_: int
            for compr_78_ in (_dafny.Map({len(_dafny.Set({_dafny.SeqWithoutIsStrInference([(D0_DC0(268)).cf0, -311])})): len(_dafny.Map({(self).f10: len(_dafny.Map({False: 551}))}))})).keys.Elements:
                d_2461_v0_: int = compr_78_
                if (d_2461_v0_) in (_dafny.Map({len(_dafny.Set({_dafny.SeqWithoutIsStrInference([(D0_DC0(268)).cf0, -311])})): len(_dafny.Map({(self).f10: len(_dafny.Map({False: 551}))}))})):
                    coll78_[default__.safeDivisionInt(d_2461_v0_, len(iife165_()
                    ))] = 604
            return _dafny.Map(coll78_)
        return (_dafny.Set({D0_DC0(len(iife164_()
)), D0_DC0(-987), D0_DC0(-923), D0_DC0(186)})).ispropersubset(_dafny.Set({D0_DC0(801), D0_DC0(239)}))

    def m1(self, p0, p1, globalState):
        r0: int = int(0)
        r1: int = int(0)
        r2: bool = False
        (globalState).f9 = default__.safeDivisionInt((0) - ((0) - (p1)), p1)
        if (self).fm6(globalState):
            d_2463_v1_: _dafny.Seq
            def iife166_():
                coll80_ = _dafny.Set()
                compr_80_: int
                for compr_80_ in _dafny.IntegerRange(163, -320):
                    d_2464_v0_: int = compr_80_
                    if ((163) <= (d_2464_v0_)) and ((d_2464_v0_) < (-320)):
                        coll80_ = coll80_.union(_dafny.Set([(d_2464_v0_) * (len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "icn"))))]))
                return _dafny.Set(coll80_)
            d_2463_v1_ = _dafny.SeqWithoutIsStrInference([_dafny.Set({p1}), p0, iife166_()
            ])
            def iife167_():
                coll81_ = _dafny.Set()
                compr_81_: int
                for compr_81_ in _dafny.IntegerRange(844, 84):
                    d_2465_v2_: int = compr_81_
                    if ((844) <= (d_2465_v2_)) and ((d_2465_v2_) < (84)):
                        coll81_ = coll81_.union(_dafny.Set([default__.safeModuloInt(d_2465_v2_, 708)]))
                return _dafny.Set(coll81_)
            d_2463_v1_ = _dafny.SeqWithoutIsStrInference([p0, iife167_()
            , p0])
            d_2466_v3_: _dafny.Map
            d_2466_v3_ = _dafny.Map({not(not((self).f10)): (self).f10})
            r2 = not(((len(d_2466_v3_) if not((self).f10) else p1)) > (326))
            d_2467_v4_: _dafny.Array
            def lambda150_(d_2468_i0_):
                return (self).f10

            init81_ = lambda150_
            nw434_ = _dafny.Array(None, 19)
            for i0_81_ in range(nw434_.length(0)):
                nw434_[i0_81_] = init81_(i0_81_)
            d_2467_v4_ = nw434_
            d_2469_v5_: _dafny.Array
            nw435_ = _dafny.Array(None, 6)
            nw435_[int(0)] = p1
            nw435_[int(1)] = p1
            nw435_[int(2)] = p1
            nw435_[int(3)] = p1
            nw435_[int(4)] = p1
            nw435_[int(5)] = 192
            d_2469_v5_ = nw435_
            d_2470_v6_: D0
            d_2470_v6_ = D0_DC1(_dafny.MultiSet([d_2469_v5_]), p1, p1)
            d_2471_v7_: _dafny.Map
            d_2471_v7_ = _dafny.Map({d_2467_v4_: d_2470_v6_})
            d_2472_v8_: _dafny.Seq
            d_2472_v8_ = _dafny.SeqWithoutIsStrInference([d_2471_v7_])
            d_2473_v9_: _dafny.MultiSet
            d_2473_v9_ = _dafny.MultiSet([d_2469_v5_])
            d_2474_v10_: _dafny.Map
            d_2474_v10_ = _dafny.Map({p1: d_2471_v7_})
            d_2475_v11_: _dafny.Seq
            d_2475_v11_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jxkmnorw"))
            d_2476_v12_: _dafny.Map
            d_2476_v12_ = _dafny.Map({d_2475_v11_: d_2471_v7_})
            pat_let_tv46_ = d_2473_v9_
            d_2477_v13_: _dafny.Array
            nw436_ = _dafny.Array(None, 12)
            nw436_[int(0)] = d_2471_v7_
            nw436_[int(1)] = (d_2472_v8_)[default__.safeIndex(p1, len(d_2472_v8_))]
            nw436_[int(2)] = d_2471_v7_
            def iife168_(_pat_let43_0):
                def iife169_(d_2478_dt__update__tmp_h0_):
                    def iife170_(_pat_let44_0):
                        def iife171_(d_2479_dt__update_hcf1_h0_):
                            return D0_DC1(d_2479_dt__update_hcf1_h0_, (d_2478_dt__update__tmp_h0_).cf2, (d_2478_dt__update__tmp_h0_).cf3)
                        return iife171_(_pat_let44_0)
                    return iife170_(pat_let_tv46_)
                return iife169_(_pat_let43_0)
            nw436_[int(3)] = ((d_2471_v7_).set(d_2467_v4_, d_2470_v6_)) | (_dafny.Map({d_2467_v4_: iife168_(D0_DC1(_dafny.MultiSet([d_2469_v5_]), p1, p1))}))
            nw436_[int(4)] = (_dafny.Map({d_2467_v4_: d_2470_v6_})) | (d_2471_v7_)
            nw436_[int(5)] = _dafny.Map({d_2467_v4_: d_2470_v6_})
            nw436_[int(6)] = (_dafny.Map({d_2467_v4_: d_2470_v6_})) | ((d_2471_v7_).set(d_2467_v4_, d_2470_v6_))
            nw436_[int(7)] = ((d_2474_v10_)[p1] if (p1) in (d_2474_v10_) else d_2471_v7_)
            nw436_[int(8)] = d_2471_v7_
            nw436_[int(9)] = d_2471_v7_
            nw436_[int(10)] = (((d_2476_v12_)[d_2475_v11_] if (d_2475_v11_) in (d_2476_v12_) else _dafny.Map({d_2467_v4_: d_2470_v6_}))) | (d_2471_v7_)
            nw436_[int(11)] = _dafny.Map({d_2467_v4_: d_2470_v6_})
            d_2477_v13_ = nw436_
            d_2480_v14_: D1
            d_2480_v14_ = D1_DC3(_dafny.Map({d_2467_v4_: d_2470_v6_}))
            index427_ = default__.safeIndex(51, (d_2477_v13_).length(0))
            (d_2477_v13_)[index427_] = (d_2480_v14_).cf8
            index428_ = default__.safeIndex(51, (d_2477_v13_).length(0))
            (d_2477_v13_)[index428_] = d_2471_v7_
            (globalState).f8 = (self).f10
            d_2481_v15_: _dafny.Map
            d_2481_v15_ = _dafny.Map({False: (0) - (default__.safeDivisionInt(p1, p1))})
            d_2482_v16_: _dafny.Array
            nw437_ = _dafny.Array(_dafny.MultiSet({}), 4)
            d_2482_v16_ = nw437_
            d_2483_v17_: _dafny.MultiSet
            d_2483_v17_ = _dafny.MultiSet([d_2470_v6_])
            index429_ = default__.safeIndex(393, (d_2482_v16_).length(0))
            (d_2482_v16_)[index429_] = d_2483_v17_
            d_2484_v18_: _dafny.Seq
            d_2484_v18_ = _dafny.SeqWithoutIsStrInference([d_2475_v11_])
            d_2485_v19_: _dafny.Map
            d_2485_v19_ = _dafny.Map({(self).f10: (d_2481_v15_).set((self).f10, p1)})
            d_2486_v20_: _dafny.Map
            d_2486_v20_ = _dafny.Map({(self).f10: ((d_2485_v19_)[(self).f10] if ((self).f10) in (d_2485_v19_) else d_2481_v15_)})
            d_2487_v21_: _dafny.Seq
            d_2487_v21_ = _dafny.SeqWithoutIsStrInference([(d_2481_v15_) | (((d_2486_v20_)[(self).f10] if ((self).f10) in (d_2486_v20_) else d_2481_v15_))])
            index430_ = default__.safeIndex(393, (d_2482_v16_).length(0))
            rhs475_ = (d_2487_v21_)[default__.safeIndex(p1, len(d_2487_v21_))]
            rhs476_ = d_2483_v17_
            rhs477_ = (_dafny.SeqWithoutIsStrInference([d_2475_v11_])) + (d_2484_v18_)
            rhs478_ = (self).f10
            lhs371_ = d_2482_v16_
            lhs372_ = default__.safeIndex(393, (d_2482_v16_).length(0))
            lhs373_ = globalState
            d_2481_v15_ = rhs475_
            lhs371_[lhs372_] = rhs476_
            d_2484_v18_ = rhs477_
            lhs373_.f8 = rhs478_
        elif True:
            d_2488_v22_: _dafny.Seq
            d_2488_v22_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "i"))
            d_2489_v23_: str
            d_2489_v23_ = _dafny.CodePoint('u')
            d_2490_v24_: _dafny.Map
            d_2490_v24_ = _dafny.Map({p1: (self).f10})
            d_2491_v25_: _dafny.Seq
            d_2491_v25_ = _dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])
            d_2492_v26_: D0
            d_2492_v26_ = D0_DC2(_dafny.MultiSet([True]), False, p1, (d_2491_v25_)[default__.safeIndex(p1, len(d_2491_v25_))])
            d_2493_v27_: _dafny.MultiSet
            d_2493_v27_ = _dafny.MultiSet([(d_2492_v26_).cf6, p1])
            d_2494_v29_: D0
            d_2494_v29_ = D0_DC0(p1)
            d_2495_v30_: _dafny.Map
            d_2495_v30_ = _dafny.Map({(self).f10: d_2494_v29_})
            d_2496_v31_: _dafny.MultiSet
            d_2496_v31_ = _dafny.MultiSet([d_2495_v30_, _dafny.Map({(self).f10: d_2494_v29_}), d_2495_v30_])
            d_2497_v32_: _dafny.Map
            d_2497_v32_ = _dafny.Map({(d_2496_v31_).cardinality: d_2489_v23_})
            d_2498_v33_: _dafny.Array
            nw438_ = _dafny.Array(None, 19)
            nw438_[int(0)] = len(d_2490_v24_)
            nw438_[int(1)] = p1
            nw438_[int(2)] = 472
            nw438_[int(3)] = p1
            nw438_[int(4)] = p1
            nw438_[int(5)] = p1
            nw438_[int(6)] = len(d_2488_v22_)
            nw438_[int(7)] = p1
            nw438_[int(8)] = p1
            nw438_[int(9)] = (d_2493_v27_).cardinality
            nw438_[int(10)] = p1
            nw438_[int(11)] = len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cxhqskyew")))
            def iife172_():
                coll82_ = _dafny.Map()
                compr_82_: int
                for compr_82_ in _dafny.IntegerRange(943, -329):
                    d_2499_v28_: int = compr_82_
                    if ((943) <= (d_2499_v28_)) and ((d_2499_v28_) < (-329)):
                        coll82_[(d_2499_v28_) + (p1)] = not((self).f10)
                return _dafny.Map(coll82_)
            nw438_[int(12)] = len(iife172_()
            )
            nw438_[int(13)] = p1
            nw438_[int(14)] = p1
            nw438_[int(15)] = 250
            nw438_[int(16)] = len(d_2497_v32_)
            nw438_[int(17)] = (d_2493_v27_).cardinality
            nw438_[int(18)] = (0) - (p1)
            d_2498_v33_ = nw438_
            d_2500_v34_: _dafny.MultiSet
            d_2500_v34_ = _dafny.MultiSet([d_2498_v33_, d_2498_v33_])
            d_2501_v35_: D0
            d_2501_v35_ = D0_DC1(d_2500_v34_, len(d_2491_v25_), p1)
            d_2502_v36_: _dafny.Array
            nw439_ = _dafny.Array(None, 15)
            nw439_[int(0)] = (p1 if (self).f10 else p1)
            nw439_[int(1)] = default__.safeDivisionInt((0) - (p1), len(_dafny.SeqWithoutIsStrInference([(self).f10, (self).f10])))
            nw439_[int(2)] = len(p0)
            nw439_[int(3)] = p1
            nw439_[int(4)] = len(d_2488_v22_)
            nw439_[int(5)] = (p1) * (p1)
            nw439_[int(6)] = p1
            nw439_[int(7)] = len((_dafny.Map({d_2489_v23_: d_2498_v33_})) | (_dafny.Map({d_2489_v23_: d_2498_v33_})))
            nw439_[int(8)] = p1
            nw439_[int(9)] = default__.safeDivisionInt(p1, len(d_2488_v22_))
            nw439_[int(10)] = p1
            nw439_[int(11)] = p1
            nw439_[int(12)] = (d_2501_v35_).cf3
            nw439_[int(13)] = (p1) * (p1)
            nw439_[int(14)] = p1
            d_2502_v36_ = nw439_
            d_2502_v36_ = d_2502_v36_
            d_2503_v37_: D2
            d_2503_v37_ = D2_DC5(d_2488_v22_)
            d_2504_v38_: _dafny.Array
            nw440_ = _dafny.Array(None, 5)
            nw440_[int(0)] = d_2488_v22_
            nw440_[int(1)] = d_2488_v22_
            nw440_[int(2)] = default__.fm2(p1, d_2493_v27_, default__.fm7(globalState), globalState)
            nw440_[int(3)] = (d_2503_v37_).cf10
            nw440_[int(4)] = d_2488_v22_
            d_2504_v38_ = nw440_
            index431_ = default__.safeIndex(422, (d_2504_v38_).length(0))
            (d_2504_v38_)[index431_] = d_2488_v22_
            index432_ = default__.safeIndex(422, (d_2504_v38_).length(0))
            (d_2504_v38_)[index432_] = (d_2488_v22_) + ((d_2488_v22_).set(default__.safeIndex(p1, len(d_2488_v22_)), d_2489_v23_))
            d_2505_v39_: D13
            d_2505_v39_ = D13_DC29(d_2490_v24_)
            d_2506_v40_: _dafny.Array
            nw441_ = _dafny.Array(None, 6)
            nw441_[int(0)] = d_2490_v24_
            nw441_[int(1)] = _dafny.Map({p1: (self).f10})
            nw441_[int(2)] = d_2490_v24_
            nw441_[int(3)] = (d_2505_v39_).cf50
            nw441_[int(4)] = d_2490_v24_
            nw441_[int(5)] = (d_2505_v39_).cf50
            d_2506_v40_ = nw441_
            d_2507_v41_: C2
            nw442_ = C2()
            nw442_.ctor__(d_2506_v40_)
            d_2507_v41_ = nw442_
            d_2508_v42_: _dafny.Array
            nw443_ = _dafny.Array(_dafny.MultiSet({}), 19)
            d_2508_v42_ = nw443_
            d_2509_v43_: _dafny.Seq
            d_2509_v43_ = _dafny.SeqWithoutIsStrInference([len(_dafny.Map({_dafny.Map({(self).f10: p1}): p1}))])
            index433_ = default__.safeIndex(469, (d_2508_v42_).length(0))
            (d_2508_v42_)[index433_] = _dafny.MultiSet((d_2509_v43_) + (d_2509_v43_))
            d_2510_v44_: _dafny.Seq
            d_2510_v44_ = _dafny.SeqWithoutIsStrInference([d_2493_v27_, _dafny.MultiSet([p1])])
            index434_ = default__.safeIndex(469, (d_2508_v42_).length(0))
            (d_2508_v42_)[index434_] = ((d_2510_v44_)[default__.safeIndex(p1, len(d_2510_v44_))]) | (d_2493_v27_)
            d_2511_v45_: D4
            d_2511_v45_ = D4_DC10(_dafny.Set({(self).f10}))
            d_2511_v45_ = d_2511_v45_
        d_2512_v46_: _dafny.Seq
        d_2512_v46_ = _dafny.SeqWithoutIsStrInference([(self).f10, not((self).f10)])
        r1 = len(d_2512_v46_)
        d_2513_v47_: str
        d_2513_v47_ = _dafny.CodePoint('o')
        d_2514_v48_: _dafny.Seq
        d_2514_v48_ = _dafny.SeqWithoutIsStrInference([d_2513_v47_, d_2513_v47_])
        d_2515_v49_: D1
        d_2515_v49_ = D1_DC4((d_2514_v48_)[default__.safeIndex(p1, len(d_2514_v48_))])
        d_2516_v50_: _dafny.Array
        nw444_ = _dafny.Array(None, 8)
        nw444_[int(0)] = default__.fm3(p1, not(not((self).f10)), globalState)
        nw444_[int(1)] = d_2513_v47_
        nw444_[int(2)] = d_2513_v47_
        nw444_[int(3)] = (d_2515_v49_).cf9
        nw444_[int(4)] = d_2513_v47_
        nw444_[int(5)] = d_2513_v47_
        nw444_[int(6)] = d_2513_v47_
        nw444_[int(7)] = (d_2514_v48_)[default__.safeIndex(p1, len(d_2514_v48_))]
        d_2516_v50_ = nw444_
        index435_ = default__.safeIndex(35, (d_2516_v50_).length(0))
        (d_2516_v50_)[index435_] = d_2513_v47_
        index436_ = default__.safeIndex(35, (d_2516_v50_).length(0))
        (d_2516_v50_)[index436_] = d_2513_v47_
        if ((0) - (p1)) > ((p1) + (len(p0))):
            r2 = (self).f10
            default__.m0(globalState)
            (globalState).f9 = len(d_2514_v48_)
            d_2517_v51_: _dafny.Seq
            d_2517_v51_ = _dafny.SeqWithoutIsStrInference([-738, p1])
            d_2518_v52_: _dafny.MultiSet
            d_2518_v52_ = _dafny.MultiSet([(self).f10])
            d_2519_v53_: D0
            d_2519_v53_ = D0_DC2((d_2518_v52_).set((self).f10, default__.abs(p1)), (self).f10, p1, (self).f10)
            d_2520_v54_: _dafny.Map
            d_2520_v54_ = _dafny.Map({p1: p1})
            d_2521_v55_: _dafny.Set
            d_2521_v55_ = _dafny.Set({_dafny.SeqWithoutIsStrInference([p1]), d_2517_v51_, d_2517_v51_, d_2517_v51_, d_2517_v51_})
            d_2522_v56_: _dafny.Array
            nw445_ = _dafny.Array(int(0), 14)
            d_2522_v56_ = nw445_
            d_2523_v57_: _dafny.Set
            d_2523_v57_ = _dafny.Set({d_2522_v56_, d_2522_v56_, d_2522_v56_})
            d_2524_v58_: _dafny.Array
            nw446_ = _dafny.Array(None, 24)
            nw446_[int(0)] = default__.safeModuloInt(p1, len(d_2517_v51_))
            nw446_[int(1)] = (p1) - (p1)
            nw446_[int(2)] = p1
            nw446_[int(3)] = 309
            nw446_[int(4)] = (d_2519_v53_).cf6
            nw446_[int(5)] = len(d_2520_v54_)
            nw446_[int(6)] = p1
            nw446_[int(7)] = (p1) * (p1)
            nw446_[int(8)] = default__.safeDivisionInt(p1, p1)
            nw446_[int(9)] = default__.safeDivisionInt(len(_dafny.Set({_dafny.SeqWithoutIsStrInference([(self).f10]), d_2512_v46_})), len((d_2514_v48_).set(default__.safeIndex(p1, len(d_2514_v48_)), (d_2516_v50_)[default__.safeIndex(35, (d_2516_v50_).length(0))])))
            nw446_[int(10)] = default__.safeModuloInt(default__.fm0((self).f10, globalState), len((d_2514_v48_).set(default__.safeIndex(p1, len(d_2514_v48_)), _dafny.CodePoint('a'))))
            nw446_[int(11)] = len(d_2521_v55_)
            nw446_[int(12)] = p1
            nw446_[int(13)] = (0) - (p1)
            nw446_[int(14)] = p1
            nw446_[int(15)] = (default__.fm0(True, globalState)) + (p1)
            nw446_[int(16)] = p1
            nw446_[int(17)] = (0) - ((0) - ((len(d_2514_v48_)) - (p1)))
            nw446_[int(18)] = len(d_2523_v57_)
            nw446_[int(19)] = len(d_2517_v51_)
            nw446_[int(20)] = p1
            nw446_[int(21)] = p1
            nw446_[int(22)] = p1
            nw446_[int(23)] = -123
            d_2524_v58_ = nw446_
            index437_ = default__.safeIndex(219, (d_2524_v58_).length(0))
            (d_2524_v58_)[index437_] = p1
            index438_ = default__.safeIndex(219, (d_2524_v58_).length(0))
            (d_2524_v58_)[index438_] = (p1) * (p1)
            d_2525_v59_: _dafny.Map
            d_2525_v59_ = _dafny.Map({(d_2514_v48_) + (d_2514_v48_): (self).f10})
            if ((d_2525_v59_)[d_2514_v48_] if (d_2514_v48_) in (d_2525_v59_) else (self).f10):
                (globalState).f8 = (self).f10
                d_2514_v48_ = d_2514_v48_
                d_2526_v60_: _dafny.Array
                nw447_ = _dafny.Array(_dafny.Map({}), 21)
                d_2526_v60_ = nw447_
                d_2527_v61_: C3
                nw448_ = C3()
                nw448_.ctor__((self).f10, not((self).f10), d_2526_v60_)
                d_2527_v61_ = nw448_
                d_2527_v61_ = d_2527_v61_
                d_2528_v62_: _dafny.Array
                nw449_ = _dafny.Array(False, 13)
                d_2528_v62_ = nw449_
                index439_ = default__.safeIndex(354, (d_2528_v62_).length(0))
                (d_2528_v62_)[index439_] = (self).f10
                index440_ = default__.safeIndex(354, (d_2528_v62_).length(0))
                (d_2528_v62_)[index440_] = ((d_2512_v46_) < (d_2512_v46_)) == ((d_2512_v46_)[default__.safeIndex((d_2524_v58_)[default__.safeIndex(219, (d_2524_v58_).length(0))], len(d_2512_v46_))])
                d_2529_v63_: _dafny.Map
                d_2529_v63_ = _dafny.Map({(d_2528_v62_)[default__.safeIndex(354, (d_2528_v62_).length(0))]: (d_2524_v58_)[default__.safeIndex(219, (d_2524_v58_).length(0))]})
                d_2530_v64_: _dafny.Map
                d_2530_v64_ = _dafny.Map({(d_2529_v63_) | ((d_2529_v63_).set(not((d_2528_v62_)[default__.safeIndex(354, (d_2528_v62_).length(0))]), p1)): d_2512_v46_})
                d_2530_v64_ = (d_2530_v64_).set((d_2529_v63_) | (_dafny.Map({(self).f10: p1})), (_dafny.SeqWithoutIsStrInference([(d_2528_v62_)[default__.safeIndex(354, (d_2528_v62_).length(0))], (d_2527_v61_).f22])) + (_dafny.SeqWithoutIsStrInference([(d_2528_v62_)[default__.safeIndex(354, (d_2528_v62_).length(0))], d_2527_v61_.f21])))
            elif True:
                d_2531_v66_: _dafny.Array
                def lambda151_(d_2532_v58_, d_2533_p1_):
                    def lambda152_(d_2534_i1_):
                        def iife173_():
                            coll83_ = _dafny.Map()
                            compr_83_: int
                            for compr_83_ in _dafny.IntegerRange(609, 188):
                                d_2535_v65_: int = compr_83_
                                if ((609) <= (d_2535_v65_)) and ((d_2535_v65_) < (188)):
                                    coll83_[default__.safeDivisionInt(d_2535_v65_, (d_2532_v58_)[default__.safeIndex(219, (d_2532_v58_).length(0))])] = d_2533_p1_
                            return _dafny.Map(coll83_)
                        return iife173_()
                        

                    return lambda152_

                init82_ = lambda151_(d_2524_v58_, p1)
                nw450_ = _dafny.Array(None, 16)
                for i0_82_ in range(nw450_.length(0)):
                    nw450_[i0_82_] = init82_(i0_82_)
                d_2531_v66_ = nw450_
                d_2536_v67_: D10
                d_2536_v67_ = D10_DC23(d_2531_v66_)
                d_2537_v68_: _dafny.Seq
                d_2537_v68_ = _dafny.SeqWithoutIsStrInference([d_2536_v67_, D10_DC23(d_2531_v66_)])
                d_2537_v68_ = d_2537_v68_
                d_2538_v69_: _dafny.Map
                d_2538_v69_ = _dafny.Map({919: False})
                d_2539_v71_: D13
                d_2539_v71_ = D13_DC29(d_2538_v69_)
                d_2540_v73_: _dafny.Array
                nw451_ = _dafny.Array(None, 16)
                nw451_[int(0)] = d_2538_v69_
                nw451_[int(1)] = d_2538_v69_
                nw451_[int(2)] = d_2538_v69_
                nw451_[int(3)] = d_2538_v69_
                def iife174_():
                    coll84_ = _dafny.Map()
                    compr_84_: int
                    for compr_84_ in _dafny.IntegerRange(706, 184):
                        d_2541_v70_: int = compr_84_
                        if ((706) <= (d_2541_v70_)) and ((d_2541_v70_) < (184)):
                            coll84_[(d_2541_v70_) + (p1)] = (self).f10
                    return _dafny.Map(coll84_)
                nw451_[int(4)] = iife174_()
                
                nw451_[int(5)] = (d_2539_v71_).cf50
                nw451_[int(6)] = d_2538_v69_
                nw451_[int(7)] = d_2538_v69_
                def iife175_():
                    coll85_ = _dafny.Map()
                    compr_85_: int
                    for compr_85_ in _dafny.IntegerRange(956, 115):
                        d_2542_v72_: int = compr_85_
                        if ((956) <= (d_2542_v72_)) and ((d_2542_v72_) < (115)):
                            coll85_[(d_2542_v72_) * ((d_2524_v58_)[default__.safeIndex(219, (d_2524_v58_).length(0))])] = (self).f10
                    return _dafny.Map(coll85_)
                nw451_[int(8)] = iife175_()
                
                nw451_[int(9)] = d_2538_v69_
                nw451_[int(10)] = d_2538_v69_
                nw451_[int(11)] = d_2538_v69_
                nw451_[int(12)] = d_2538_v69_
                nw451_[int(13)] = d_2538_v69_
                nw451_[int(14)] = d_2538_v69_
                nw451_[int(15)] = _dafny.Map({len(_dafny.Set({(self).f10})): (self).f10})
                d_2540_v73_ = nw451_
                d_2543_v74_: C2
                nw452_ = C2()
                nw452_.ctor__(d_2540_v73_)
                d_2543_v74_ = nw452_
                d_2544_v75_: _dafny.Array
                nw453_ = _dafny.Array(False, 7)
                d_2544_v75_ = nw453_
                d_2544_v75_ = d_2544_v75_
                rhs479_ = (self).f10
                rhs480_ = p1
                rhs481_ = ((d_2514_v48_) + (d_2514_v48_) if (self).f10 else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "cydhyslov")))
                lhs374_ = globalState
                lhs375_ = globalState
                lhs374_.f8 = rhs479_
                lhs375_.f9 = rhs480_
                d_2514_v48_ = rhs481_
                d_2545_v76_: _dafny.Set
                d_2545_v76_ = _dafny.Set({(self).f10, not((self).f10)})
                d_2546_v77_: _dafny.Map
                d_2546_v77_ = _dafny.Map({len(d_2545_v76_): d_2516_v50_})
                d_2547_v78_: _dafny.Seq
                d_2547_v78_ = _dafny.SeqWithoutIsStrInference([d_2544_v75_, d_2544_v75_])
                d_2546_v77_ = (d_2546_v77_).set(len((d_2547_v78_) + (d_2547_v78_)), d_2516_v50_)
        elif True:
            d_2548_v79_: _dafny.Seq
            d_2548_v79_ = _dafny.SeqWithoutIsStrInference([p1, p1])
            d_2514_v48_ = (d_2514_v48_ if (len(_dafny.SeqWithoutIsStrInference([False]))) > (len(d_2548_v79_)) else _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "g")))
            d_2549_v80_: _dafny.MultiSet
            d_2549_v80_ = _dafny.MultiSet([not((self).f10), not(default__.fm1((self).f10, d_2513_v47_, globalState))])
            d_2550_v81_: _dafny.Map
            d_2550_v81_ = _dafny.Map({((d_2549_v80_)[(self).f10] if ((self).f10) in (d_2549_v80_) else p1): p1})
            d_2551_v82_: _dafny.Map
            d_2551_v82_ = _dafny.Map({default__.fm17(globalState): (0) - (len(d_2550_v81_))})
            d_2552_v83_: D2
            d_2552_v83_ = D2_DC7(False, p1, True, d_2514_v48_)
            d_2551_v82_ = (_dafny.Map({d_2552_v83_: p1})) | (d_2551_v82_)
            r0 = p1
            d_2553_v84_: _dafny.Array
            nw454_ = _dafny.Array(_dafny.Array(None, 0), 28)
            d_2553_v84_ = nw454_
            d_2554_v85_: _dafny.Array
            nw455_ = _dafny.Array(None, 16)
            d_2554_v85_ = nw455_
            d_2555_v86_: _dafny.Array
            def lambda153_(d_2556_p1_):
                def lambda154_(d_2557_i2_):
                    return _dafny.Map({d_2556_p1_: (self).f10})

                return lambda154_

            init83_ = lambda153_(p1)
            nw456_ = _dafny.Array(None, 17)
            for i0_83_ in range(nw456_.length(0)):
                nw456_[i0_83_] = init83_(i0_83_)
            d_2555_v86_ = nw456_
            d_2558_v87_: C4
            nw457_ = C4()
            nw457_.ctor__(d_2555_v86_)
            d_2558_v87_ = nw457_
            d_2559_v88_: C4
            d_2559_v88_ = d_2558_v87_
            index441_ = default__.safeIndex(596, (d_2554_v85_).length(0))
            (d_2554_v85_)[index441_] = (d_2559_v88_)
            d_2560_v89_: _dafny.Set
            d_2560_v89_ = _dafny.Set({(self).f10, (self).f10})
            d_2561_v90_: D4
            d_2561_v90_ = D4_DC10(_dafny.Set({not((self).f10)}))
            index442_ = default__.safeIndex(596, (d_2554_v85_).length(0))
            rhs482_ = (d_2560_v89_).isdisjoint(((d_2561_v90_).cf18).intersection((d_2561_v90_).cf18))
            rhs483_ = p1
            rhs484_ = (default__.fm0((self).f10, globalState)) + ((p1) + (p1))
            rhs485_ = d_2553_v84_
            rhs486_ = d_2558_v87_
            lhs376_ = globalState
            lhs377_ = globalState
            lhs378_ = d_2554_v85_
            lhs379_ = default__.safeIndex(596, (d_2554_v85_).length(0))
            lhs376_.f8 = rhs482_
            lhs377_.f9 = rhs483_
            r0 = rhs484_
            d_2553_v84_ = rhs485_
            lhs378_[lhs379_] = rhs486_
            (globalState).f8 = ((self).f10 if (self).f10 else (self).f10)
        d_2562_v91_: _dafny.MultiSet
        d_2562_v91_ = _dafny.MultiSet([(self).f10])
        d_2563_v92_: D0
        d_2563_v92_ = D0_DC2(d_2562_v91_, False, p1, False)
        d_2564_v93_: D8
        d_2564_v93_ = D8_DC20(d_2563_v92_, True, (self).f10)
        source32_ = d_2564_v93_
        if source32_.is_DC19:
            d_2565___mcc_h0_ = source32_.cf34
            d_2566___mcc_h1_ = source32_.cf35
            d_2567_cf35_ = d_2566___mcc_h1_
            d_2568_cf34_ = d_2565___mcc_h0_
            d_2569_v94_: _dafny.Map
            d_2569_v94_ = _dafny.Map({d_2567_cf35_: (self).f10})
            d_2570_v97_: _dafny.Seq
            d_2570_v97_ = _dafny.SeqWithoutIsStrInference([d_2512_v46_])
            d_2571_v99_: D13
            d_2571_v99_ = D13_DC29(d_2569_v94_)
            pat_let_tv47_ = d_2569_v94_
            d_2572_v100_: _dafny.Array
            nw458_ = _dafny.Array(None, 17)
            nw458_[int(0)] = d_2569_v94_
            nw458_[int(1)] = _dafny.Map({428: (self).f10})
            nw458_[int(2)] = d_2569_v94_
            nw458_[int(3)] = d_2569_v94_
            nw458_[int(4)] = d_2569_v94_
            def iife176_():
                coll86_ = _dafny.Map()
                compr_86_: int
                for compr_86_ in _dafny.IntegerRange(-281, 217):
                    d_2573_v95_: int = compr_86_
                    if ((-281) <= (d_2573_v95_)) and ((d_2573_v95_) < (217)):
                        coll86_[(d_2573_v95_) * (322)] = (self).f10
                return _dafny.Map(coll86_)
            nw458_[int(5)] = iife176_()
            
            def iife177_():
                coll87_ = _dafny.Map()
                compr_87_: int
                for compr_87_ in _dafny.IntegerRange(185, 173):
                    d_2574_v96_: int = compr_87_
                    if ((185) <= (d_2574_v96_)) and ((d_2574_v96_) < (173)):
                        coll87_[default__.safeModuloInt(d_2574_v96_, -70)] = (self).f10
                return _dafny.Map(coll87_)
            nw458_[int(6)] = iife177_()
            
            nw458_[int(7)] = default__.fm27(d_2570_v97_, globalState)
            def iife178_():
                coll88_ = _dafny.Map()
                compr_88_: int
                for compr_88_ in (p0).Elements:
                    d_2575_v98_: int = compr_88_
                    if (d_2575_v98_) in (p0):
                        coll88_[default__.safeDivisionInt(d_2575_v98_, p1)] = (self).f10
                return _dafny.Map(coll88_)
            nw458_[int(8)] = iife178_()
            
            nw458_[int(9)] = _dafny.Map({d_2567_cf35_: True})
            nw458_[int(10)] = d_2569_v94_
            def iife179_(_pat_let45_0):
                def iife180_(d_2576_dt__update__tmp_h1_):
                    def iife181_(_pat_let46_0):
                        def iife182_(d_2577_dt__update_hcf50_h0_):
                            return D13_DC29(d_2577_dt__update_hcf50_h0_)
                        return iife182_(_pat_let46_0)
                    return iife181_(pat_let_tv47_)
                return iife180_(_pat_let45_0)
            nw458_[int(11)] = (iife179_(d_2571_v99_)).cf50
            nw458_[int(12)] = d_2569_v94_
            nw458_[int(13)] = _dafny.Map({460: True})
            nw458_[int(14)] = d_2569_v94_
            nw458_[int(15)] = d_2569_v94_
            nw458_[int(16)] = d_2569_v94_
            d_2572_v100_ = nw458_
            d_2578_v101_: C7
            nw459_ = C7()
            nw459_.ctor__((d_2567_cf35_) != (default__.fm0((self).f10, globalState)), d_2572_v100_, D0_DC0(p1), not((d_2567_cf35_) == ((_dafny.MultiSet([d_2567_cf35_])).cardinality)))
            d_2578_v101_ = nw459_
            d_2513_v47_ = (d_2516_v50_)[default__.safeIndex(35, (d_2516_v50_).length(0))]
            d_2512_v46_ = (d_2512_v46_) + (d_2512_v46_)
            d_2579_v102_: _dafny.Map
            d_2579_v102_ = _dafny.Map({(d_2578_v101_).f14: (d_2578_v101_).f14})
            d_2580_v103_: _dafny.Seq
            d_2580_v103_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({(self).f10: (d_2578_v101_).f14}), d_2579_v102_])
            d_2581_v104_: _dafny.Seq
            d_2581_v104_ = _dafny.SeqWithoutIsStrInference([(0) - ((d_2578_v101_).fm9(d_2515_v49_, _dafny.Set({len((d_2580_v103_)[default__.safeIndex(p1, len(d_2580_v103_))]), p1}), globalState))])
            r1 = default__.safeModuloInt(len(d_2581_v104_), (d_2562_v91_).cardinality)
        elif source32_.is_DC20:
            d_2582___mcc_h2_ = source32_.cf36
            d_2583___mcc_h3_ = source32_.cf37
            d_2584___mcc_h4_ = source32_.cf38
            d_2585_cf38_ = d_2584___mcc_h4_
            d_2586_cf37_ = d_2583___mcc_h3_
            d_2587_cf36_ = d_2582___mcc_h2_
            d_2588_v105_: _dafny.Array
            def lambda155_(d_2589_i3_):
                return (self).f10

            init84_ = lambda155_
            nw460_ = _dafny.Array(None, 26)
            for i0_84_ in range(nw460_.length(0)):
                nw460_[i0_84_] = init84_(i0_84_)
            d_2588_v105_ = nw460_
            d_2590_v106_: _dafny.Seq
            d_2590_v106_ = _dafny.SeqWithoutIsStrInference([(0) - ((D8_DC19(d_2514_v48_, len(d_2514_v48_))).cf35), p1, p1])
            index443_ = default__.safeIndex(899, (d_2588_v105_).length(0))
            (d_2588_v105_)[index443_] = ((d_2590_v106_)[default__.safeIndex(p1, len(d_2590_v106_))]) <= (p1)
            d_2591_v107_: _dafny.Array
            nw461_ = _dafny.Array(int(0), 15)
            d_2591_v107_ = nw461_
            d_2592_v108_: _dafny.MultiSet
            d_2592_v108_ = _dafny.MultiSet([d_2591_v107_])
            d_2593_v109_: D0
            d_2593_v109_ = D0_DC1(d_2592_v108_, p1, 454)
            d_2594_v110_: _dafny.Map
            d_2594_v110_ = _dafny.Map({d_2593_v109_: d_2588_v105_})
            d_2595_v111_: _dafny.Seq
            d_2595_v111_ = _dafny.SeqWithoutIsStrInference([((d_2594_v110_)[d_2593_v109_] if (d_2593_v109_) in (d_2594_v110_) else d_2588_v105_)])
            index444_ = default__.safeIndex(899, (d_2588_v105_).length(0))
            rhs487_ = (d_2512_v46_)[default__.safeIndex(p1, len(d_2512_v46_))]
            rhs488_ = (d_2595_v111_)[default__.safeIndex(366, len(d_2595_v111_))]
            rhs489_ = (d_2516_v50_)[default__.safeIndex(35, (d_2516_v50_).length(0))]
            rhs490_ = (self).f10
            lhs380_ = d_2588_v105_
            lhs381_ = default__.safeIndex(899, (d_2588_v105_).length(0))
            lhs380_[lhs381_] = rhs487_
            d_2588_v105_ = rhs488_
            d_2513_v47_ = rhs489_
            r2 = rhs490_
            rhs491_ = ((p0) | (p0)) == ((p0).intersection(p0))
            rhs492_ = d_2586_cf37_
            rhs493_ = p1
            lhs382_ = globalState
            d_2585_cf38_ = rhs491_
            lhs382_.f8 = rhs492_
            r0 = rhs493_
            d_2596_v112_: C6
            nw462_ = C6()
            nw462_.ctor__(p1, d_2514_v48_)
            d_2596_v112_ = nw462_
            index445_ = default__.safeIndex(899, (d_2588_v105_).length(0))
            (d_2588_v105_)[index445_] = d_2586_cf37_
        elif True:
            d_2597___mcc_h5_ = source32_.cf33
            d_2598_cf33_ = d_2597___mcc_h5_
            (globalState).f8 = default__.fm1((self).f10, d_2513_v47_, globalState)
            d_2599_v113_: _dafny.Array
            nw463_ = _dafny.Array(_dafny.Map({}), 4)
            d_2599_v113_ = nw463_
            d_2600_v114_: _dafny.Seq
            d_2600_v114_ = _dafny.SeqWithoutIsStrInference([d_2599_v113_, d_2599_v113_])
            d_2601_v115_: C3
            nw464_ = C3()
            nw464_.ctor__((self).f10, (self).f10, (d_2600_v114_)[default__.safeIndex((0) - (p1), len(d_2600_v114_))])
            d_2601_v115_ = nw464_
            d_2512_v46_ = _dafny.SeqWithoutIsStrInference([(p1) >= (-322), d_2601_v115_.f21, d_2601_v115_.f21, not(d_2601_v115_.f21)])
            d_2602_v116_: _dafny.Array
            nw465_ = _dafny.Array(None, 15)
            nw465_[int(0)] = (d_2512_v46_)[default__.safeIndex(p1, len(d_2512_v46_))]
            nw465_[int(1)] = (d_2512_v46_)[default__.safeIndex(p1, len(d_2512_v46_))]
            nw465_[int(2)] = (d_2601_v115_).f22
            nw465_[int(3)] = (d_2601_v115_).f22
            nw465_[int(4)] = False
            nw465_[int(5)] = (d_2601_v115_).f22
            nw465_[int(6)] = (self).f10
            nw465_[int(7)] = False
            nw465_[int(8)] = d_2601_v115_.f21
            nw465_[int(9)] = not(d_2601_v115_.f21)
            nw465_[int(10)] = d_2601_v115_.f21
            nw465_[int(11)] = (d_2601_v115_).f22
            nw465_[int(12)] = d_2601_v115_.f21
            nw465_[int(13)] = (d_2601_v115_).f22
            nw465_[int(14)] = (self).f10
            d_2602_v116_ = nw465_
            d_2603_v117_: D7
            d_2603_v117_ = D7_DC17(d_2514_v48_, d_2602_v116_, (d_2601_v115_).f22, p1)
            d_2604_v118_: _dafny.Map
            d_2604_v118_ = _dafny.Map({(self).f10: d_2602_v116_})
            d_2605_v119_: _dafny.Map
            d_2605_v119_ = _dafny.Map({p1: (self).fm6(globalState)})
            d_2606_v120_: D13
            d_2606_v120_ = D13_DC29(d_2605_v119_)
            d_2607_v121_: _dafny.Seq
            d_2607_v121_ = _dafny.SeqWithoutIsStrInference([d_2606_v120_, d_2606_v120_, d_2606_v120_, d_2606_v120_, D13_DC29((d_2606_v120_).cf50)])
            def iife183_():
                coll89_ = _dafny.Set()
                compr_89_: D13
                for compr_89_ in (d_2607_v121_).Elements:
                    d_2608_v122_: D13 = compr_89_
                    if (d_2608_v122_) in (d_2607_v121_):
                        coll89_ = coll89_.union(_dafny.Set([d_2608_v122_]))
                return _dafny.Set(coll89_)
            d_2603_v117_ = D7_DC17(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "oggvmdojg")), ((d_2604_v118_)[d_2601_v115_.f21] if (d_2601_v115_.f21) in (d_2604_v118_) else d_2602_v116_), default__.fm1((self).f10, (d_2516_v50_)[default__.safeIndex(35, (d_2516_v50_).length(0))], globalState), len(iife183_()
))
        d_2609_v123_: _dafny.Array
        def lambda156_(d_2610_p1_):
            def lambda157_(d_2611_i4_):
                return _dafny.Map({d_2610_p1_: (self).f10})

            return lambda157_

        init85_ = lambda156_(p1)
        nw466_ = _dafny.Array(None, 2)
        for i0_85_ in range(nw466_.length(0)):
            nw466_[i0_85_] = init85_(i0_85_)
        d_2609_v123_ = nw466_
        d_2612_v124_: D0
        d_2612_v124_ = D0_DC0(p1)
        d_2613_v125_: T1
        nw467_ = C7()
        nw467_.ctor__((self).f10, d_2609_v123_, d_2612_v124_, (self).f10)
        d_2613_v125_ = nw467_
        d_2614_v126_: _dafny.Map
        d_2614_v126_ = _dafny.Map({(self).f10: d_2613_v125_})
        d_2615_v127_: _dafny.MultiSet
        d_2615_v127_ = _dafny.MultiSet([p1, p1])
        d_2616_v128_: _dafny.MultiSet
        d_2616_v128_ = _dafny.MultiSet([default__.safeDivisionInt(-533, len(default__.fm23((self).f10, globalState))), p1, len(_dafny.Map({p1: p1})), len(d_2614_v126_), (d_2615_v127_).cardinality])
        r0 = (d_2615_v127_).cardinality
        r1 = default__.safeDivisionInt(p1, (0) - (p1))
        r2 = (self).f10
        return r0, r1, r2

    def m2(self, p0, globalState):
        r0: bool = False
        d_2617_v0_: _dafny.Seq
        d_2617_v0_ = _dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "mdworlyct"))
        d_2618_v1_: _dafny.Seq
        d_2618_v1_ = _dafny.SeqWithoutIsStrInference([d_2617_v0_])
        d_2618_v1_ = d_2618_v1_
        d_2619_v2_: _dafny.Array
        nw468_ = _dafny.Array(False, 20)
        d_2619_v2_ = nw468_
        d_2620_v3_: _dafny.Map
        d_2620_v3_ = _dafny.Map({(self).f10: d_2619_v2_})
        d_2620_v3_ = d_2620_v3_
        d_2621_v4_: int
        d_2621_v4_ = -574
        hi20_ = default__.safeDivisionInt(456, d_2621_v4_)
        for d_2622_i0_ in range(len(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('n') for d_2676_i1_ in range(default__.abs(-574))])), hi20_):
            index446_ = default__.safeIndex(334, (d_2619_v2_).length(0))
            (d_2619_v2_)[index446_] = (self).f10
            index447_ = default__.safeIndex(334, (d_2619_v2_).length(0))
            (d_2619_v2_)[index447_] = (self).f10
            d_2623_v5_: str
            d_2623_v5_ = _dafny.CodePoint('i')
            d_2624_v6_: _dafny.Map
            d_2624_v6_ = _dafny.Map({d_2623_v5_: False})
            d_2625_v7_: _dafny.Map
            d_2625_v7_ = _dafny.Map({(self).f10: (len(d_2624_v6_)) <= ((0) - (d_2622_i0_))})
            d_2625_v7_ = (d_2625_v7_) | (d_2625_v7_)
            d_2626_v8_: D2
            d_2626_v8_ = D2_DC8(d_2622_i0_, d_2623_v5_)
            source33_ = d_2626_v8_
            if source33_.is_DC6:
                d_2627_v9_: _dafny.Seq
                d_2627_v9_ = _dafny.SeqWithoutIsStrInference([(self).f10])
                d_2627_v9_ = d_2627_v9_
                d_2628_v11_: _dafny.Map
                d_2628_v11_ = _dafny.Map({d_2622_i0_: d_2622_i0_})
                d_2629_v12_: _dafny.Map
                d_2629_v12_ = d_2628_v11_
                d_2630_v13_: _dafny.Map
                d_2630_v13_ = _dafny.Map({d_2629_v12_: d_2622_i0_})
                d_2631_v14_: _dafny.Map
                d_2631_v14_ = _dafny.Map({d_2629_v12_: (d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))]})
                d_2632_v15_: _dafny.Map
                def iife184_():
                    coll90_ = _dafny.Map()
                    compr_90_: _dafny.Map
                    for compr_90_ in (d_2630_v13_).keys.Elements:
                        d_2633_v10_: _dafny.Map = compr_90_
                        if (d_2633_v10_) in (d_2630_v13_):
                            coll90_[d_2633_v10_] = not((d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))])
                    return _dafny.Map(coll90_)
                d_2632_v15_ = _dafny.Map({(0) - (d_2621_v4_): (iife184_()
                ) | (d_2631_v14_)})
                d_2632_v15_ = (d_2632_v15_).set(default__.fm0((self).f10, globalState), d_2631_v14_)
                d_2634_v16_: _dafny.MultiSet
                d_2634_v16_ = _dafny.MultiSet([d_2621_v4_, d_2621_v4_, (_dafny.MultiSet((d_2618_v1_).set(default__.safeIndex(d_2621_v4_, len(d_2618_v1_)), d_2617_v0_))).cardinality])
                d_2635_v17_: _dafny.Map
                d_2635_v17_ = _dafny.Map({not((self).f10): d_2622_i0_})
                (globalState).f8 = (561) <= (((d_2634_v16_)[d_2621_v4_] if (d_2621_v4_) in (d_2634_v16_) else ((d_2635_v17_)[(self).f10] if ((self).f10) in (d_2635_v17_) else len(d_2628_v11_))))
                d_2617_v0_ = ((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "iqu"))) + (d_2617_v0_) if ((d_2625_v7_)[True] if (True) in (d_2625_v7_) else True) else d_2617_v0_)
            elif source33_.is_DC7:
                d_2636___mcc_h0_ = source33_.cf11
                d_2637___mcc_h1_ = source33_.cf12
                d_2638___mcc_h2_ = source33_.cf13
                d_2639___mcc_h3_ = source33_.cf14
                d_2640_cf14_ = d_2639___mcc_h3_
                d_2641_cf13_ = d_2638___mcc_h2_
                d_2642_cf12_ = d_2637___mcc_h1_
                d_2643_cf11_ = d_2636___mcc_h0_
                d_2644_v18_: _dafny.Map
                d_2644_v18_ = _dafny.Map({(638) + ((0) - (d_2622_i0_)): d_2617_v0_})
                d_2644_v18_ = (d_2644_v18_).set(default__.fm0(d_2643_cf11_, globalState), d_2640_cf14_)
                d_2645_v19_: _dafny.Array
                def lambda158_(d_2646_v4_, d_2647_cf13_):
                    def lambda159_(d_2648_i2_):
                        return _dafny.Map({d_2646_v4_: d_2647_cf13_})

                    return lambda159_

                init86_ = lambda158_(d_2621_v4_, d_2641_cf13_)
                nw469_ = _dafny.Array(None, 14)
                for i0_86_ in range(nw469_.length(0)):
                    nw469_[i0_86_] = init86_(i0_86_)
                d_2645_v19_ = nw469_
                d_2649_v20_: T0
                nw470_ = C3()
                nw470_.ctor__(d_2641_cf13_, not(d_2643_cf11_), d_2645_v19_)
                d_2649_v20_ = nw470_
                d_2650_v21_: _dafny.Seq
                d_2650_v21_ = _dafny.SeqWithoutIsStrInference([d_2622_i0_, 320])
                rhs494_ = default__.fm1((self).f10, d_2623_v5_, globalState)
                rhs495_ = d_2649_v20_
                rhs496_ = default__.safeDivisionInt(len((_dafny.Map({_dafny.SeqWithoutIsStrInference([d_2621_v4_]): d_2622_i0_})).set(d_2650_v21_, len(d_2640_cf14_))), len((d_2650_v21_) + (d_2650_v21_)))
                r0 = rhs494_
                d_2649_v20_ = rhs495_
                d_2642_cf12_ = rhs496_
                d_2643_cf11_ = not(d_2643_cf11_)
                index448_ = default__.safeIndex(334, (d_2619_v2_).length(0))
                index449_ = default__.safeIndex(334, (d_2619_v2_).length(0))
                rhs497_ = not(d_2641_cf13_)
                rhs498_ = not(not(d_2643_cf11_))
                lhs383_ = d_2619_v2_
                lhs384_ = default__.safeIndex(334, (d_2619_v2_).length(0))
                lhs385_ = d_2619_v2_
                lhs386_ = default__.safeIndex(334, (d_2619_v2_).length(0))
                lhs383_[lhs384_] = rhs497_
                lhs385_[lhs386_] = rhs498_
            elif source33_.is_DC8:
                d_2651___mcc_h4_ = source33_.cf15
                d_2652___mcc_h5_ = source33_.cf16
                d_2653_cf16_ = d_2652___mcc_h5_
                d_2654_cf15_ = d_2651___mcc_h4_
                d_2655_v22_: _dafny.Map
                d_2655_v22_ = _dafny.Map({_dafny.SeqWithoutIsStrInference([(d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))]]): not((d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))])})
                d_2656_v23_: _dafny.Seq
                d_2656_v23_ = _dafny.SeqWithoutIsStrInference([(d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))]])
                d_2657_v24_: _dafny.Map
                d_2657_v24_ = _dafny.Map({(d_2655_v22_) | (_dafny.Map({d_2656_v23_: (self).f10})): _dafny.SeqWithoutIsStrInference([d_2621_v4_, 622, len(d_2656_v23_), d_2621_v4_])})
                d_2621_v4_ = len((((d_2657_v24_)[d_2655_v22_] if (d_2655_v22_) in (d_2657_v24_) else _dafny.SeqWithoutIsStrInference([145 for d_2658_i3_ in range(default__.abs(706))]))).set(default__.safeIndex(d_2654_cf15_, len(((d_2657_v24_)[d_2655_v22_] if (d_2655_v22_) in (d_2657_v24_) else _dafny.SeqWithoutIsStrInference([145 for d_2659_i3_ in range(default__.abs(706))])))), d_2621_v4_))
                (globalState).f9 = d_2654_cf15_
                d_2660_v25_: _dafny.Array
                nw471_ = _dafny.Array(_dafny.Map({}), 15)
                d_2660_v25_ = nw471_
                d_2661_v26_: C5
                nw472_ = C5()
                nw472_.ctor__(d_2660_v25_)
                d_2661_v26_ = nw472_
                d_2662_v27_: D0
                d_2662_v27_ = D0_DC0(default__.fm0((self).f10, globalState))
                d_2663_v28_: T1
                nw473_ = C7()
                nw473_.ctor__((self).f10, d_2660_v25_, d_2662_v27_, (self).f10)
                d_2663_v28_ = nw473_
                d_2664_v29_: _dafny.Seq
                d_2664_v29_ = _dafny.SeqWithoutIsStrInference([d_2663_v28_])
                d_2665_v30_: _dafny.Map
                d_2665_v30_ = _dafny.Map({d_2664_v29_: False})
                index450_ = default__.safeIndex(334, (d_2619_v2_).length(0))
                (d_2619_v2_)[index450_] = not ((default__.fm1((d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))], d_2623_v5_, globalState) if (d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))] else ((d_2665_v30_)[(d_2664_v29_).set(default__.safeIndex(d_2654_cf15_, len(d_2664_v29_)), d_2663_v28_)] if ((d_2664_v29_).set(default__.safeIndex(d_2654_cf15_, len(d_2664_v29_)), d_2663_v28_)) in (d_2665_v30_) else (d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))]))) or (not((d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))]))
            elif True:
                d_2666___mcc_h6_ = source33_.cf10
                d_2667_cf10_ = d_2666___mcc_h6_
                d_2625_v7_ = (d_2625_v7_).set((self).f10, (self).f10)
                d_2668_v31_: C0
                nw474_ = C0()
                nw474_.ctor__(True)
                d_2668_v31_ = nw474_
                d_2669_v32_: _dafny.Array
                nw475_ = _dafny.Array(None, 4)
                nw475_[int(0)] = d_2668_v31_
                nw475_[int(1)] = d_2668_v31_
                nw475_[int(2)] = d_2668_v31_
                nw475_[int(3)] = d_2668_v31_
                d_2669_v32_ = nw475_
                index451_ = default__.safeIndex(795, (d_2669_v32_).length(0))
                (d_2669_v32_)[index451_] = d_2668_v31_
                index452_ = default__.safeIndex(795, (d_2669_v32_).length(0))
                (d_2669_v32_)[index452_] = d_2668_v31_
                d_2670_v33_: D5
                d_2670_v33_ = D5_DC14((not((d_2668_v31_).f17)) == ((self).f10))
                d_2670_v33_ = default__.fm55((d_2622_i0_) < (d_2621_v4_), True, (0) - (d_2622_i0_), globalState)
                d_2671_v34_: D8
                d_2671_v34_ = D8_DC20(default__.fm4(d_2621_v4_, d_2621_v4_, globalState), (d_2668_v31_).f17, (self).f10)
                d_2671_v34_ = d_2671_v34_
            d_2672_v35_: _dafny.Map
            d_2672_v35_ = _dafny.Map({d_2622_i0_: not((d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))])})
            d_2673_v36_: D13
            d_2673_v36_ = D13_DC29(d_2672_v35_)
            d_2674_v37_: _dafny.Array
            nw476_ = _dafny.Array(None, 1)
            nw476_[int(0)] = (d_2673_v36_).cf50
            d_2674_v37_ = nw476_
            index453_ = default__.safeIndex(125, (d_2674_v37_).length(0))
            (d_2674_v37_)[index453_] = d_2672_v35_
            index454_ = default__.safeIndex(125, (d_2674_v37_).length(0))
            def iife185_():
                coll91_ = _dafny.Map()
                compr_91_: int
                for compr_91_ in _dafny.IntegerRange(518, -893):
                    d_2675_v38_: int = compr_91_
                    if ((518) <= (d_2675_v38_)) and ((d_2675_v38_) < (-893)):
                        coll91_[(d_2675_v38_) + (d_2621_v4_)] = False
                return _dafny.Map(coll91_)
            (d_2674_v37_)[index454_] = (d_2672_v35_) | ((iife185_()
            ).set(173, (d_2619_v2_)[default__.safeIndex(334, (d_2619_v2_).length(0))]))
        r0 = (self).fm6(globalState)
        d_2677_v40_: _dafny.Seq
        d_2677_v40_ = _dafny.SeqWithoutIsStrInference([d_2621_v4_, d_2621_v4_])
        d_2678_v41_: _dafny.Map
        d_2678_v41_ = _dafny.Map({len(d_2677_v40_): d_2621_v4_})
        d_2679_v42_: _dafny.Array
        nw477_ = _dafny.Array(None, 4)
        def iife186_():
            coll92_ = _dafny.Map()
            compr_92_: int
            for compr_92_ in _dafny.IntegerRange(403, 421):
                d_2680_v39_: int = compr_92_
                if ((403) <= (d_2680_v39_)) and ((d_2680_v39_) < (421)):
                    coll92_[(d_2680_v39_) + (d_2621_v4_)] = d_2621_v4_
            return _dafny.Map(coll92_)
        nw477_[int(0)] = iife186_()
        
        nw477_[int(1)] = d_2678_v41_
        nw477_[int(2)] = d_2678_v41_
        nw477_[int(3)] = d_2678_v41_
        d_2679_v42_ = nw477_
        d_2681_v43_: D10
        d_2681_v43_ = D10_DC23(d_2679_v42_)
        pat_let_tv48_ = d_2679_v42_
        def iife187_(_pat_let47_0):
            def iife188_(d_2682_dt__update__tmp_h0_):
                def iife189_(_pat_let48_0):
                    def iife190_(d_2683_dt__update_hcf42_h0_):
                        return D10_DC23(d_2683_dt__update_hcf42_h0_)
                    return iife190_(_pat_let48_0)
                return iife189_(pat_let_tv48_)
            return iife188_(_pat_let47_0)
        source34_ = iife187_(d_2681_v43_)
        if source34_.is_DC24:
            d_2684___mcc_h7_ = source34_.cf43
            d_2685___mcc_h8_ = source34_.cf44
            d_2686___mcc_h9_ = source34_.cf45
            d_2687_cf45_ = d_2686___mcc_h9_
            d_2688_cf44_ = d_2685___mcc_h8_
            d_2689_cf43_ = d_2684___mcc_h7_
            index455_ = default__.safeIndex(733, (d_2688_cf44_).length(0))
            (d_2688_cf44_)[index455_] = -775
            index456_ = default__.safeIndex(733, (d_2688_cf44_).length(0))
            (d_2688_cf44_)[index456_] = default__.fm0((self).f10, globalState)
            d_2690_v44_: C6
            nw478_ = C6()
            nw478_.ctor__((d_2688_cf44_)[default__.safeIndex(733, (d_2688_cf44_).length(0))], d_2617_v0_)
            d_2690_v44_ = nw478_
            d_2691_v45_: _dafny.Array
            nw479_ = _dafny.Array(_dafny.Map({}), 17)
            d_2691_v45_ = nw479_
            d_2692_v46_: C5
            nw480_ = C5()
            nw480_.ctor__(d_2691_v45_)
            d_2692_v46_ = nw480_
            d_2693_v47_: str
            d_2693_v47_ = _dafny.CodePoint('j')
            d_2694_v48_: _dafny.MultiSet
            d_2694_v48_ = _dafny.MultiSet([d_2693_v47_, d_2693_v47_])
            (globalState).f8 = ((d_2694_v48_).intersection(d_2694_v48_)).issubset((d_2694_v48_).set(d_2693_v47_, default__.abs((d_2688_cf44_)[default__.safeIndex(733, (d_2688_cf44_).length(0))])))
        elif True:
            d_2695___mcc_h10_ = source34_.cf42
            d_2696_cf42_ = d_2695___mcc_h10_
            d_2697_v49_: D11
            d_2697_v49_ = D11_DC26(d_2621_v4_)
            d_2698_v50_: D4
            d_2698_v50_ = D4_DC11(602, d_2621_v4_, (self).f10)
            d_2699_v51_: _dafny.Map
            d_2699_v51_ = _dafny.Map({d_2698_v50_: d_2619_v2_})
            d_2700_v52_: _dafny.Map
            d_2700_v52_ = _dafny.Map({d_2697_v49_: ((d_2699_v51_)[d_2698_v50_] if (d_2698_v50_) in (d_2699_v51_) else d_2619_v2_)})
            d_2701_v53_: _dafny.Map
            d_2701_v53_ = _dafny.Map({d_2621_v4_: (d_2619_v2_ if (self).f10 else ((d_2700_v52_)[d_2697_v49_] if (d_2697_v49_) in (d_2700_v52_) else d_2619_v2_))})
            d_2702_v54_: _dafny.MultiSet
            d_2702_v54_ = _dafny.MultiSet([d_2621_v4_, len(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "jcvn"))), d_2621_v4_])
            d_2703_v55_: _dafny.Seq
            d_2703_v55_ = _dafny.SeqWithoutIsStrInference([d_2619_v2_, d_2619_v2_])
            d_2619_v2_ = ((d_2701_v53_)[(d_2702_v54_).cardinality] if ((d_2702_v54_).cardinality) in (d_2701_v53_) else (d_2703_v55_)[default__.safeIndex(921, len(d_2703_v55_))])
            d_2704_v56_: _dafny.Array
            nw481_ = _dafny.Array(int(0), 22)
            d_2704_v56_ = nw481_
            d_2705_v57_: _dafny.Map
            d_2705_v57_ = _dafny.Map({(d_2621_v4_) + (d_2621_v4_): d_2704_v56_})
            d_2705_v57_ = (d_2705_v57_).set((d_2621_v4_) * (d_2621_v4_), d_2704_v56_)
            d_2678_v41_ = (d_2678_v41_).set(343, d_2621_v4_)
            d_2706_v58_: _dafny.Map
            d_2706_v58_ = _dafny.Map({len((d_2617_v0_) + (d_2617_v0_)): (self).f10})
            r0 = ((d_2706_v58_)[(d_2621_v4_) - (d_2621_v4_)] if ((d_2621_v4_) - (d_2621_v4_)) in (d_2706_v58_) else (self).f10)
        d_2707_v59_: _dafny.Array
        nw482_ = _dafny.Array(_dafny.Map({}), 2)
        d_2707_v59_ = nw482_
        d_2708_v60_: D0
        d_2708_v60_ = D0_DC0(d_2621_v4_)
        d_2709_v61_: C7
        nw483_ = C7()
        nw483_.ctor__(True, d_2707_v59_, d_2708_v60_, (self).f10)
        d_2709_v61_ = nw483_
        d_2710_v62_: C1
        nw484_ = C1()
        nw484_.ctor__((d_2709_v61_).f14, (d_2709_v61_).fm10(globalState), d_2708_v60_, (self).f10)
        d_2710_v62_ = nw484_
        d_2711_v63_: _dafny.Map
        d_2711_v63_ = _dafny.Map({d_2709_v61_: d_2710_v62_})
        d_2712_v64_: _dafny.Seq
        d_2712_v64_ = _dafny.SeqWithoutIsStrInference([d_2711_v63_, d_2711_v63_])
        index457_ = default__.safeIndex(408, (d_2619_v2_).length(0))
        (d_2619_v2_)[index457_] = not ((d_2710_v62_).f20) or ((d_2710_v62_).f20)
        d_2713_v65_: _dafny.Array
        nw485_ = _dafny.Array(_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "")), 28)
        d_2713_v65_ = nw485_
        index458_ = default__.safeIndex(884, (d_2713_v65_).length(0))
        (d_2713_v65_)[index458_] = d_2617_v0_
        d_2714_v66_: D7
        d_2714_v66_ = D7_DC17(_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('j') for d_2715_i4_ in range(default__.abs(633))]), d_2619_v2_, (d_2710_v62_).f19, d_2621_v4_)
        d_2716_v67_: str
        d_2716_v67_ = _dafny.CodePoint('l')
        index459_ = default__.safeIndex(408, (d_2619_v2_).length(0))
        index460_ = default__.safeIndex(884, (d_2713_v65_).length(0))
        rhs499_ = _dafny.SeqWithoutIsStrInference([_dafny.Map({d_2709_v61_: d_2710_v62_}), (_dafny.Map({d_2709_v61_: d_2710_v62_})) | ((d_2711_v63_).set(d_2709_v61_, d_2710_v62_)), (_dafny.Map({d_2709_v61_: d_2710_v62_})) | (d_2711_v63_)])
        rhs500_ = not((_dafny.SeqWithoutIsStrInference(map(_dafny.CodePoint, "qldina"))) <= (_dafny.SeqWithoutIsStrInference([_dafny.CodePoint('p') for d_2717_i5_ in range(default__.abs(497))])))
        rhs501_ = (d_2617_v0_).set(default__.safeIndex(d_2621_v4_, len(d_2617_v0_)), d_2716_v67_)
        rhs502_ = d_2714_v66_
        lhs387_ = d_2619_v2_
        lhs388_ = default__.safeIndex(408, (d_2619_v2_).length(0))
        lhs389_ = d_2713_v65_
        lhs390_ = default__.safeIndex(884, (d_2713_v65_).length(0))
        d_2712_v64_ = rhs499_
        lhs387_[lhs388_] = rhs500_
        lhs389_[lhs390_] = rhs501_
        d_2714_v66_ = rhs502_
        d_2718_v68_: _dafny.Map
        d_2718_v68_ = _dafny.Map({(d_2710_v62_).f20: (d_2709_v61_).f14})
        d_2719_v69_: _dafny.MultiSet
        d_2719_v69_ = _dafny.MultiSet([(d_2710_v62_).f20, (d_2619_v2_)[default__.safeIndex(408, (d_2619_v2_).length(0))]])
        r0 = ((d_2718_v68_)[(not((d_2619_v2_)[default__.safeIndex(408, (d_2619_v2_).length(0))])) in (d_2719_v69_)] if ((not((d_2619_v2_)[default__.safeIndex(408, (d_2619_v2_).length(0))])) in (d_2719_v69_)) in (d_2718_v68_) else (d_2709_v61_).f14)
        return r0

    @property
    def f10(self):
        return self._f10
